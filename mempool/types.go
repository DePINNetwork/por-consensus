package mempool

import (
	memprotos "github.com/depinnetwork/por-consensus/api/cometbft/mempool/v2"
	"github.com/depinnetwork/por-consensus/types"
)

var (
	_ types.Wrapper   = &memprotos.Txs{}
	_ types.Wrapper   = &memprotos.HaveTx{}
	_ types.Wrapper   = &memprotos.ResetRoute{}
	_ types.Unwrapper = &memprotos.Message{}
)
