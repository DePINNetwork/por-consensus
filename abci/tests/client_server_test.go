package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	abciclient "github.com/depinnetwork/por-consensus/abci/client"
	"github.com/depinnetwork/por-consensus/abci/example/kvstore"
	abciserver "github.com/depinnetwork/por-consensus/abci/server"
)

func TestClientServerNoAddrPrefix(t *testing.T) {
	t.Helper()

	addr := "localhost:26658"
	transport := "socket"
	app := kvstore.NewInMemoryApplication()

	server, err := abciserver.NewServer(addr, transport, app)
	require.NoError(t, err)
	err = server.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := server.Stop(); err != nil {
			t.Error(err)
		}
	})

	client, err := abciclient.NewClient(addr, transport, true)
	require.NoError(t, err)
	err = client.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := client.Stop(); err != nil {
			t.Error(err)
		}
	})
}
