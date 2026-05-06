# Challenge 3: Broadcast

## Goal

Build a broadcast system where nodes receive integer messages, store them locally, and eventually propagate them through the cluster.

The broadcast challenge is split into several parts. It starts with a single-node version and gradually adds multi-node replication, network partitions, and performance requirements.

## What I Learned

- Broadcast protocols need deduplication because the same message can arrive more than once.
- Idempotent handlers are safer in distributed systems because retries and duplicate delivery are normal.
- Local state must be protected because Maelstrom's Go node runs handlers concurrently.
- Multi-node broadcast requires propagation, not just local storage.
- Gossip protocols trade stronger immediate consistency for simpler, eventually consistent convergence.

## Challenge Parts

| Part | Status | Focus |
| --- | --- | --- |
| 3a: Single-Node Broadcast | Complete | Store and read local messages |
| 3b: Multi-Node Broadcast | Complete | Propagate messages between nodes |
| 3c: Fault-Tolerant Broadcast | Not started | Recover from partitions |
| 3d/3e: Efficient Broadcast | Not started | Reduce message volume and latency |

## Design Direction

The natural progression for this challenge is:

1. Store every received message locally.
2. Deduplicate messages with a `seen` set.
3. Forward newly seen messages to peers.
4. Retry or periodically gossip messages so partitions can heal.
5. Optimize topology and batching to reduce message count.

The key idea is convergence: every node should eventually learn the same set of messages, even if messages arrive in different orders or more than once.

Part 3b implements the first multi-node version with fire-and-forget gossip. Fault tolerance and efficiency are intentionally left for the later parts of the challenge.
