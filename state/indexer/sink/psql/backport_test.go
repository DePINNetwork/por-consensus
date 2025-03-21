package psql

import (
	"github.com/depinnetwork/por-consensus/state/indexer"
	"github.com/depinnetwork/por-consensus/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
