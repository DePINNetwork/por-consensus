package imports

// This file imports all API packages to ensure they're built correctly

import (
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1beta1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1beta2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1beta3"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/crypto/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/libs/bits/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/mempool/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/blocksync/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/consensus/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/p2p/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/privval/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/services/block/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/services/block_results/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/services/pruning/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/services/version/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/state/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/state/v1beta2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/state/v1beta3"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/state/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/statesync/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/store/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/types/v1beta1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/types/v1beta2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/types/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/version/v1"
)

// UseAllImports is a placeholder function to avoid unused import errors
func UseAllImports() {}
