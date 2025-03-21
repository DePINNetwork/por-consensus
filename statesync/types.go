package statesync

import (
	ssproto "github.com/depinnetwork/por-consensus/api/cometbft/statesync/v1"
	"github.com/depinnetwork/por-consensus/types"
)

var (
	_ types.Wrapper = &ssproto.ChunkRequest{}
	_ types.Wrapper = &ssproto.ChunkResponse{}
	_ types.Wrapper = &ssproto.SnapshotsRequest{}
	_ types.Wrapper = &ssproto.SnapshotsResponse{}
)
