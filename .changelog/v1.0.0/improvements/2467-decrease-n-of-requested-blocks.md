- `[blocksync]` make the max number of downloaded blocks dynamic.
  Previously it was a const 600. Now it's `peersCount * maxPendingRequestsPerPeer (20)`
  ([\#2467](https://github.com/depinnetwork/por-consensus/pull/2467))
