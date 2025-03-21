package null

import (
	"context"
	"errors"

	abci "github.com/depinnetwork/por-consensus/abci/types"
	"github.com/depinnetwork/por-consensus/libs/log"
	"github.com/depinnetwork/por-consensus/libs/pubsub/query"
	"github.com/depinnetwork/por-consensus/state/txindex"
)

var _ txindex.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

func (*TxIndex) SetRetainHeight(_ int64) error {
	return nil
}

func (*TxIndex) GetRetainHeight() (int64, error) {
	return 0, nil
}

func (*TxIndex) Prune(_ int64) (numPruned, newRetainHeight int64, err error) {
	return 0, 0, nil
}

// Get on a TxIndex is disabled and panics when invoked.
func (*TxIndex) Get(_ []byte) (*abci.TxResult, error) {
	return nil, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

// AddBatch is a noop and always returns nil.
func (*TxIndex) AddBatch(_ *txindex.Batch) error {
	return nil
}

// Index is a noop and always returns nil.
func (*TxIndex) Index(_ *abci.TxResult) error {
	return nil
}

func (*TxIndex) Search(_ context.Context, _ *query.Query, _ txindex.Pagination) ([]*abci.TxResult, int, error) {
	return []*abci.TxResult{}, 0, nil
}

func (*TxIndex) SetLogger(log.Logger) {
}

func (*TxIndex) Close() error {
	return nil
}
