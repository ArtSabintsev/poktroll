Feature: Stake Gateway Namespaces

    Scenario: User can stake a Gateway
        Given the user has the pocketd binary installed
        And the user verifies the "gateway" for account "gateway2" is not staked
        And the account "gateway2" has a balance greater than "1000070" uPOKT
        When the user stakes a "gateway" with "1000070" uPOKT from the account "gateway2"
        Then the user should be able to see standard output containing "txhash:"
        And the user should be able to see standard output containing "code: 0"
        And the pocketd binary should exit without error
        And the user should wait for the "gateway" module "StakeGateway" message to be submitted
        And the "gateway" for account "gateway2" is staked with "1000070" uPOKT
        And the account balance of "gateway2" should be "1000070" uPOKT "less" than before

    Scenario: User can unstake a Gateway
        Given the user has the pocketd binary installed
        And the "gateway" for account "gateway2" is staked with "1000070" uPOKT
        And an account exists for "gateway2"
        When the user unstakes a "gateway" from the account "gateway2"
        Then the user should be able to see standard output containing "txhash:"
        And the user should be able to see standard output containing "code: 0"
        And the pocketd binary should exit without error
        And the user should wait for the "gateway" module "UnstakeGateway" message to be submitted
        And the gateway for account "gateway2" is unbonding
        And the user should wait for the "gateway" module "GatewayUnbondingBegin" tx event to be broadcast
        And a "gateway" module "GatewayUnbondingEnd" end block event is broadcast
        And the user verifies the "gateway" for account "gateway2" is not staked
        And the account balance of "gateway2" should be "1000070" uPOKT "more" than before