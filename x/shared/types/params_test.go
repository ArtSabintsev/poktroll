package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams_ValidateNumBlocksPerSession(t *testing.T) {
	tests := []struct {
		desc                string
		numBlocksPerSession any
		err                 error
	}{
		{
			desc:                "invalid type",
			numBlocksPerSession: "invalid",
			err:                 ErrSharedParamInvalid.Wrapf("invalid parameter type: %T", "invalid"),
		},
		{
			desc:                "zero NumBlocksPerSession",
			numBlocksPerSession: uint64(0),
			err:                 ErrSharedParamInvalid.Wrapf("invalid NumBlocksPerSession: (%v)", uint64(0)),
		},
		{
			desc:                "valid NumBlocksPerSession",
			numBlocksPerSession: uint64(4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := ValidateNumBlocksPerSession(tt.numBlocksPerSession)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParams_ValidateClaimWindowOpenOffsetBlocks(t *testing.T) {
	tests := []struct {
		desc                        string
		claimWindowOpenOffsetBlocks any
		err                         error
	}{
		{
			desc:                        "invalid type",
			claimWindowOpenOffsetBlocks: "invalid",
			err:                         ErrSharedParamInvalid.Wrapf("invalid parameter type: %T", "invalid"),
		},
		{
			desc:                        "valid ClaimWindowOpenOffsetBlocks",
			claimWindowOpenOffsetBlocks: uint64(4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := ValidateClaimWindowOpenOffsetBlocks(tt.claimWindowOpenOffsetBlocks)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
