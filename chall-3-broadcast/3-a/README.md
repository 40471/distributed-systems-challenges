# Challenge 3a: Single-Node Broadcast

## Goal

Implement the broadcast workload for a single node. The node must support three message types:

- `broadcast`: store a message locally.
- `read`: return all locally stored messages.
- `topology`: acknowledge the topology message.

Because this part runs with only one node, no network propagation is required yet.

## What I Learned

- Broadcast starts with local durability of received messages.
- Deduplication is important even before gossip is introduced.
- The order of returned messages does not matter for this workload.
- Maelstrom handlers can run concurrently, so shared state needs synchronization.

## Approach

The solution stores messages in two structures:

- A `seen` map for O(1) deduplication.
- A `msgs` slice for returning all stored values on `read`.

When a `broadcast` request arrives, the node checks whether the message has already been seen. If it is new, the node records it and appends it to the list returned by future reads.

Both structures are protected by a `sync.RWMutex`. Broadcast requests take the write lock because they may mutate state. Read requests take the read lock and copy the slice before replying, so the response does not share memory with state that another handler may modify later.

## Tradeoffs

This approach is simple and efficient for a single-node workload. It does not solve multi-node propagation yet, but it gives the later gossip implementation a safe local storage foundation.

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

- Use typed response structs for clearer protocol documentation.
- Reuse the same safe local-state pattern when implementing multi-node gossip.
