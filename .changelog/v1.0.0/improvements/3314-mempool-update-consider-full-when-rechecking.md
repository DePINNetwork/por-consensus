- `[mempool]` Before updating the mempool, consider it as full if rechecking is still in progress.
  This will stop accepting transactions in the mempool if the node can't keep up with re-CheckTx.
  ([\#3314](https://github.com/depinnetwork/por-consensus/pull/3314))
