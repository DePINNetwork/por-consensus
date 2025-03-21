- `[mempool]` Change the signature of `CheckTx` in the `Mempool` interface to
`CheckTx(tx types.Tx, sender p2p.ID) (*abcicli.ReqRes, error)`.
([\#1010](https://github.com/depinnetwork/por-consensus/issues/1010), [\#3084](https://github.com/depinnetwork/por-consensus/issues/3084))
