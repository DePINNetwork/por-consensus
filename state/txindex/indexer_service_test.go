package txindex_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/depinnetwork/por-consensus/abci/types"
	cmtdb "github.com/depinnetwork/por-consensus/db"
	"github.com/depinnetwork/por-consensus/libs/log"
	"github.com/depinnetwork/por-consensus/state/indexer"
	blockidxkv "github.com/depinnetwork/por-consensus/state/indexer/block/kv"
	"github.com/depinnetwork/por-consensus/state/txindex"
	"github.com/depinnetwork/por-consensus/state/txindex/kv"
	"github.com/depinnetwork/por-consensus/types"
)

func TestIndexerServiceIndexesBlocks(t *testing.T) {
	_, txIndexer, blockIndexer, eventBus := createTestSetup(t)

	height := int64(1)

	events, txResult1, txResult2 := getEventsAndResults(height)
	// publish block with events
	err := eventBus.PublishEventNewBlockEvents(events)
	require.NoError(t, err)

	err = eventBus.PublishEventTx(types.EventDataTx{TxResult: *txResult1})
	require.NoError(t, err)

	err = eventBus.PublishEventTx(types.EventDataTx{TxResult: *txResult2})
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	res, err := txIndexer.Get(types.Tx(fmt.Sprintf("foo%d", height)).Hash())
	require.NoError(t, err)
	require.Equal(t, txResult1, res)

	ok, err := blockIndexer.Has(height)
	require.NoError(t, err)
	require.True(t, ok)

	res, err = txIndexer.Get(types.Tx(fmt.Sprintf("bar%d", height)).Hash())
	require.NoError(t, err)
	require.Equal(t, txResult2, res)
}

func createTestSetup(t *testing.T) (*txindex.IndexerService, *kv.TxIndex, indexer.BlockIndexer, *types.EventBus) {
	t.Helper()
	// event bus
	eventBus := types.NewEventBus()
	eventBus.SetLogger(log.TestingLogger())
	err := eventBus.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Stop(); err != nil {
			t.Error(err)
		}
	})

	// tx indexer
	store, err := cmtdb.NewInMem()
	require.NoError(t, err)
	txIndexer := kv.NewTxIndex(store)

	prefixDB, err := cmtdb.NewWithPrefix(store, []byte("block_events"))
	require.NoError(t, err)
	blockIndexer := blockidxkv.New(prefixDB)

	service := txindex.NewIndexerService(txIndexer, blockIndexer, eventBus, false)
	service.SetLogger(log.TestingLogger())
	err = service.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := service.Stop(); err != nil {
			t.Error(err)
		}
	})
	return service, txIndexer, blockIndexer, eventBus
}

func getEventsAndResults(height int64) (types.EventDataNewBlockEvents, *abci.TxResult, *abci.TxResult) {
	events := types.EventDataNewBlockEvents{
		Height: height,
		Events: []abci.Event{
			{
				Type: "begin_event",
				Attributes: []abci.EventAttribute{
					{
						Key:   "proposer",
						Value: "FCAA001",
						Index: true,
					},
				},
			},
		},
		NumTxs: int64(2),
	}
	txResult1 := &abci.TxResult{
		Height: height,
		Index:  uint32(0),
		Tx:     types.Tx(fmt.Sprintf("foo%d", height)),
		Result: abci.ExecTxResult{Code: 0},
	}
	txResult2 := &abci.TxResult{
		Height: height,
		Index:  uint32(1),
		Tx:     types.Tx(fmt.Sprintf("bar%d", height)),
		Result: abci.ExecTxResult{Code: 0},
	}
	return events, txResult1, txResult2
}
