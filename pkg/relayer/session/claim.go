package session

import (
	"context"
	"fmt"

	"github.com/pokt-network/poktroll/pkg/either"
	"github.com/pokt-network/poktroll/pkg/observable"
	"github.com/pokt-network/poktroll/pkg/observable/channel"
	"github.com/pokt-network/poktroll/pkg/observable/filter"
	"github.com/pokt-network/poktroll/pkg/observable/logging"
	"github.com/pokt-network/poktroll/pkg/relayer"
	"github.com/pokt-network/poktroll/pkg/relayer/protocol"
	sessionkeeper "github.com/pokt-network/poktroll/x/session/keeper"
)

// createClaims maps over the sessionsToClaimObs observable. For each claim batch, it:
// 1. Calculates the earliest block height at which it is safe to CreateClaims
// 2. Waits for said block and creates the claims on-chain
// 3. Maps errors to a new observable and logs them
// 4. Returns an observable of the successfully claimed sessions
// It DOES NOT BLOCK as map operations run in their own goroutines.
func (rs *relayerSessionsManager) createClaims(
	ctx context.Context,
) observable.Observable[[]relayer.SessionTree] {
	failedCreateClaimSessionsObs, failedCreateClaimSessionsPublishCh :=
		channel.NewObservable[[]relayer.SessionTree]()

	// Map sessionsToClaimObs to a new observable of the same type which is notified
	// when the session is eligible to be claimed.
	sessionsWithOpenClaimWindowObs := channel.Map(
		ctx, rs.sessionsToClaimObs,
		rs.mapWaitForEarliestCreateClaimsHeight(failedCreateClaimSessionsPublishCh),
	)

	// Map sessionsWithOpenClaimWindowObs to a new observable of an either type,
	// populated with the session or an error, which is notified after the session
	// claims have been created or an error has been encountered, respectively.
	eitherClaimedSessionsObs := channel.Map(
		ctx, sessionsWithOpenClaimWindowObs,
		rs.newMapClaimSessionsFn(failedCreateClaimSessionsPublishCh),
	)

	// TODO_TECHDEBT: pass failed create claim sessions to some retry mechanism.
	_ = failedCreateClaimSessionsObs
	logging.LogErrors(ctx, filter.EitherError(ctx, eitherClaimedSessionsObs))

	// Map eitherClaimedSessions to a new observable of []relayer.SessionTree
	// which is notified when the corresponding claims creation succeeded.
	return filter.EitherSuccess(ctx, eitherClaimedSessionsObs)
}

// mapWaitForEarliestCreateClaimsHeight returns a new MapFn that adds a delay
// between being notified and notifying.
// It calculates and waits for the earliest block height, allowed by the protocol,
// at which claims can be created for the given session number, then emits the
// session **at that moment**.
func (rs *relayerSessionsManager) mapWaitForEarliestCreateClaimsHeight(
	failedCreateClaimsSessionsPublishCh chan<- []relayer.SessionTree,
) channel.MapFn[[]relayer.SessionTree, []relayer.SessionTree] {
	return func(
		ctx context.Context,
		sessionTrees []relayer.SessionTree,
	) (_ []relayer.SessionTree, skip bool) {
		return rs.waitForEarliestCreateClaimsHeight(
			ctx, sessionTrees, failedCreateClaimsSessionsPublishCh,
		), false
	}
}

// waitForEarliestCreateClaimsHeight calculates and waits for (blocking until) the
// earliest block height, allowed by the protocol, at which claims can be created
// for a session with the given sessionEndHeight. It is calculated relative to
// sessionEndHeight using on-chain governance parameters and randomized input.
// It IS A BLOCKING function.
func (rs *relayerSessionsManager) waitForEarliestCreateClaimsHeight(
	ctx context.Context,
	sessionTrees []relayer.SessionTree,
	failSubmitProofsSessionsCh chan<- []relayer.SessionTree,
) []relayer.SessionTree {
	// Given the sessionTrees are grouped by their sessionEndHeight, we can use the
	// first one from the group to calculate the earliest height for claim creation.
	sessionEndHeight := sessionTrees[0].GetSessionHeader().GetSessionEndBlockHeight()

	// TODO_TECHDEBT(@red-0ne): Centralize the business logic that involves taking
	// into account the heights, windows and grace periods into helper functions.
	// An additional block is added to permit to relays arriving at the last block
	// of the session to be included in the claim before the smt is closed.
	createClaimsWindowStartHeight := sessionEndHeight + sessionkeeper.GetSessionGracePeriodBlockCount() + 1

	// TODO_BLOCKER: query the on-chain governance parameter once available.
	// + claimproofparams.GovCreateClaimWindowStartHeightOffset

	// we wait for createClaimsWindowStartHeight to be received before proceeding since we need its hash
	// to know where this servicer's claim submission window starts.
	rs.logger.Info().
		Int64("create_claim_window_start_height", createClaimsWindowStartHeight).
		Msg("waiting & blocking for global earliest claim submission height")

	// TODO_BLOCKER(@bryanchriswhite): The block that'll be used as a source of entropy for
	// which branch(es) to prove should be deterministic and use on-chain governance params.
	createClaimsWindowStartBlock := rs.waitForBlock(ctx, createClaimsWindowStartHeight)

	claimsFlushedCh := make(chan []relayer.SessionTree)
	defer close(claimsFlushedCh)
	go rs.goCreateClaimRoots(
		ctx,
		sessionTrees,
		failSubmitProofsSessionsCh,
		claimsFlushedCh,
	)

	rs.logger.Info().
		Int64("create_claim_window_start_height", createClaimsWindowStartBlock.Height()).
		Str("hash", fmt.Sprintf("%x", createClaimsWindowStartBlock.Hash())).
		Msg("received global earliest claim submission height")

	earliestCreateClaimHeight := protocol.GetEarliestCreateClaimHeight(
		ctx,
		createClaimsWindowStartBlock,
	)

	rs.logger.Info().
		Int64("earliest_create_claim_height", earliestCreateClaimHeight).
		Str("hash", fmt.Sprintf("%x", createClaimsWindowStartBlock.Hash())).
		Msg("waiting & blocking for earliest claim creation height for this supplier")

	_ = rs.waitForBlock(ctx, earliestCreateClaimHeight)

	return <-claimsFlushedCh
}

// newMapClaimSessionsFn returns a new MapFn that creates claims for the given
// session number. Any session which encounters an error while creating a claim
// is sent on the failedCreateClaimSessions channel.
func (rs *relayerSessionsManager) newMapClaimSessionsFn(
	failedCreateClaimsSessionsPublishCh chan<- []relayer.SessionTree,
) channel.MapFn[[]relayer.SessionTree, either.SessionTrees] {
	return func(
		ctx context.Context,
		sessionTrees []relayer.SessionTree,
	) (_ either.SessionTrees, skip bool) {
		rs.pendingTxMu.Lock()
		defer rs.pendingTxMu.Unlock()

		if len(sessionTrees) == 0 {
			return either.Success(sessionTrees), false
		}

		sessionClaims := []*relayer.SessionClaim{}
		for _, session := range sessionTrees {
			sessionClaims = append(sessionClaims, &relayer.SessionClaim{
				RootHash:      session.GetClaimRoot(),
				SessionHeader: session.GetSessionHeader(),
			})
		}

		if err := rs.supplierClient.CreateClaims(ctx, sessionClaims); err != nil {
			failedCreateClaimsSessionsPublishCh <- sessionTrees
			return either.Error[[]relayer.SessionTree](err), false
		}

		return either.Success(sessionTrees), false
	}
}

// goCreateClaimRoots creates the claim roots corresponding to the given sessionTrees,
// then sends the successful and failed claims to their respective channels.
// This function MUST to be run as a goroutine.
func (rs *relayerSessionsManager) goCreateClaimRoots(
	ctx context.Context,
	sessionTrees []relayer.SessionTree,
	failSubmitProofsSessionsCh chan<- []relayer.SessionTree,
	claimsFlushedCh chan<- []relayer.SessionTree,
) {
	failedClaims := []relayer.SessionTree{}
	flushedClaims := []relayer.SessionTree{}
	for _, sessionTree := range sessionTrees {
		select {
		case <-ctx.Done():
			return
		default:
		}
		// This session should no longer be updated
		if _, err := sessionTree.Flush(); err != nil {
			failedClaims = append(failedClaims, sessionTree)
			continue
		}

		flushedClaims = append(flushedClaims, sessionTree)
	}

	failSubmitProofsSessionsCh <- failedClaims
	claimsFlushedCh <- flushedClaims
}
