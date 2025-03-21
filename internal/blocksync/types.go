package blocksync

import (
	cmtbs "github.com/depinnetwork/por-consensus/api/cometbft/blocksync/v2"
	"github.com/depinnetwork/por-consensus/types"
)

var (
	_ types.Wrapper = &cmtbs.StatusRequest{}
	_ types.Wrapper = &cmtbs.StatusResponse{}
	_ types.Wrapper = &cmtbs.NoBlockResponse{}
	_ types.Wrapper = &cmtbs.BlockResponse{}
	_ types.Wrapper = &cmtbs.BlockRequest{}
)
