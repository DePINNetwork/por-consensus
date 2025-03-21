package block

import (
	"errors"
	"fmt"

	"github.com/depinnetwork/por-consensus/config"
	cmtdb "github.com/depinnetwork/por-consensus/db"
	"github.com/depinnetwork/por-consensus/state/indexer"
	blockidxkv "github.com/depinnetwork/por-consensus/state/indexer/block/kv"
	blockidxnull "github.com/depinnetwork/por-consensus/state/indexer/block/null"
	"github.com/depinnetwork/por-consensus/state/indexer/sink/psql"
	"github.com/depinnetwork/por-consensus/state/txindex"
	"github.com/depinnetwork/por-consensus/state/txindex/kv"
	"github.com/depinnetwork/por-consensus/state/txindex/null"
)

// IndexerFromConfig constructs a slice of indexer.EventSink using the provided
// configuration.
func IndexerFromConfig(cfg *config.Config, dbProvider config.DBProvider, chainID string) (
	txIdx txindex.TxIndexer, blockIdx indexer.BlockIndexer, allIndexersDisabled bool, err error,
) {
	switch cfg.TxIndex.Indexer {
	case "kv":
		store, err := dbProvider(&config.DBContext{ID: "tx_index", Config: cfg})
		if err != nil {
			return nil, nil, false, err
		}

		prefixDB, err := cmtdb.NewWithPrefix(store, []byte("block_events"))
		if err != nil {
			return nil, nil, false, fmt.Errorf("creating indexer: %w", err)
		}
		return kv.NewTxIndex(store),
			blockidxkv.New(prefixDB,
				blockidxkv.WithCompaction(cfg.Storage.Compact, cfg.Storage.CompactionInterval)),
			false,
			nil

	case "psql":
		conn := cfg.TxIndex.PsqlConn
		if conn == "" {
			return nil, nil, false, errors.New("the psql connection settings cannot be empty")
		}
		opts := []psql.EventSinkOption{}

		txIndexCfg := cfg.TxIndex
		if txIndexCfg.TableBlocks != "" {
			opts = append(opts, psql.WithTableBlocks(txIndexCfg.TableBlocks))
		}

		if txIndexCfg.TableTxResults != "" {
			opts = append(opts, psql.WithTableTxResults(txIndexCfg.TableTxResults))
		}

		if txIndexCfg.TableEvents != "" {
			opts = append(opts, psql.WithTableEvents(txIndexCfg.TableEvents))
		}

		if txIndexCfg.TableAttributes != "" {
			opts = append(opts, psql.WithTableAttributes(txIndexCfg.TableAttributes))
		}

		es, err := psql.NewEventSink(cfg.TxIndex.PsqlConn, chainID, opts...)
		if err != nil {
			return nil, nil, false, fmt.Errorf("creating psql indexer: %w", err)
		}
		return es.TxIndexer(), es.BlockIndexer(), false, nil

	default:
		return &null.TxIndex{}, &blockidxnull.BlockerIndexer{}, true, nil
	}
}
