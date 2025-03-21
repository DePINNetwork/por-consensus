//go:build gofuzz || go1.20

package tests

import (
	"context"
	"testing"

	abciclient "github.com/depinnetwork/por-consensus/abci/client"
	"github.com/depinnetwork/por-consensus/abci/example/kvstore"
	"github.com/depinnetwork/por-consensus/config"
	cmtsync "github.com/depinnetwork/por-consensus/libs/sync"
	mempl "github.com/depinnetwork/por-consensus/mempool"
	"github.com/depinnetwork/por-consensus/proxy"
)

func FuzzMempool(f *testing.F) {
	app := kvstore.NewInMemoryApplication()
	mtx := new(cmtsync.Mutex)
	conn := abciclient.NewLocalClient(mtx, app)
	err := conn.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	resp, err := app.Info(context.Background(), proxy.InfoRequest)
	if err != nil {
		panic(err)
	}
	lanesInfo, err := mempl.BuildLanesInfo(resp.LanePriorities, resp.DefaultLane)
	if err != nil {
		panic(err)
	}
	mp := mempl.NewCListMempool(cfg, conn, lanesInfo, 0)

	f.Fuzz(func(_ *testing.T, data []byte) {
		_, _ = mp.CheckTx(data, "")
	})
}
