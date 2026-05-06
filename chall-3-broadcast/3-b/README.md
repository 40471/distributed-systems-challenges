# Challenge 3b: Multi-Node Broadcast

## Goal

Extend the broadcast system so messages received by one node are propagated to every other node in the cluster.

This challenge runs multiple nodes without network partitions. Every value broadcast to any node should appear in reads from all nodes within a few seconds.

## Current Status

This solution is complete for challenge 3b.

It stores messages locally, protects shared state with a `sync.RWMutex`, and gossips newly seen messages to the other nodes in the cluster.

## What I Learned

- Thread-safe local state is necessary but not sufficient for distributed replication.
- A node must distinguish between newly seen messages and duplicate messages.
- Only newly seen messages should be forwarded, otherwise nodes can create infinite message loops.
- Multi-node broadcast is about eventual convergence: all nodes should eventually have the same set of messages.

## Approach

The implementation uses a simple fire-and-forget gossip strategy:

1. Store a message if it has not been seen before.
2. Reply to the client with `broadcast_ok`.
3. Forward the new message to peer nodes.
4. When a node receives a forwarded message, apply the same deduplication logic.
5. Do not forward messages that were already seen.

Node-to-node propagation uses an internal `gossip` message type. Client `broadcast` requests still receive `broadcast_ok`, while internal gossip messages are handled without sending a reply.

This approach works because each unique message spreads through the cluster while duplicates are safely ignored. The source node is excluded when forwarding a gossiped message to reduce unnecessary backtracking.

## Tradeoffs

A naive full-mesh gossip approach is easy to reason about, but it can send many duplicate messages. That is acceptable for learning the basic multi-node behavior, but later parts of the challenge require reducing message volume and latency.

This implementation assumes the 3b environment, which does not introduce network partitions. Future fault-tolerant versions will need retry or periodic gossip. Without retries, a dropped message or partition can prevent convergence.

## How To Run

From this directory:

```sh
make build
```

To run the Maelstrom workload, provide the path to your local Maelstrom binary:

```sh
make test-maelstrom MAELSTROM=/path/to/maelstrom
```

## Future Improvements

- Add retry or periodic gossip before attempting the fault-tolerant broadcast challenge.
- Batch multiple messages in a single gossip payload to reduce message volume.
- Use a more efficient topology for the performance-focused broadcast challenges.
