# Fly.io Distributed Systems Challenges

This repository contains my solutions and learning notes for the [Fly.io Gossip Glomers](https://fly.io/dist-sys/) distributed systems challenges.

The challenges are built on top of [Maelstrom](https://github.com/jepsen-io/maelstrom), a testing framework from the Jepsen ecosystem. Maelstrom runs node processes, sends JSON messages over standard input/output, injects network behavior such as partitions and latency, and verifies whether the system satisfies the required consistency and availability properties.

The goal of this repository is not only to pass the tests, but to document the reasoning behind each design decision. Each challenge README explains what the challenge teaches, the approach used, the tradeoffs involved, and what could be improved.

## What This Project Demonstrates

- Building distributed nodes that communicate through message passing.
- Designing systems that remain available during network partitions.
- Avoiding unnecessary coordination when local decisions are enough.
- Using idempotency and deduplication to safely handle repeated messages.
- Understanding eventual consistency and convergence.
- Managing shared state safely in concurrent Go handlers.
- Evaluating tradeoffs between correctness, latency, and message volume.

## Repository Structure

```text
chall-1-echo/
  Basic Maelstrom request/reply node.

chall-2-uid-generation/
  Totally available unique ID generation.

chall-3-broadcast/
  Broadcast and gossip-based message propagation.

maelstrom/
  Local Maelstrom checkout used to run the workloads.
```

## Challenge Status

| Challenge | Status | Main Concepts |
| --- | --- | --- |
| 1: Echo | Complete | Maelstrom protocol, handlers, replies |
| 2: Unique ID Generation | Complete | Total availability, uniqueness without coordination |
| 3a: Single-Node Broadcast | Complete | Local state, deduplication, idempotency |
| 3b: Multi-Node Broadcast | Complete | Gossip, replication, convergence |

## Running Challenges

Each challenge is a separate Go module with its own Makefile. The Makefiles build local binaries with `go build`; they do not use `go install`.

Example:

```sh
cd chall-1-echo
make build
```

To run a Maelstrom workload, pass the Maelstrom binary path explicitly:

```sh
make test-maelstrom MAELSTROM=/path/to/maelstrom
```

The challenge-specific README files include the Makefile commands for each workload.

## Learning Approach

For each challenge, I try to answer three questions:

1. What correctness property is Maelstrom checking?
2. What is the simplest design that satisfies that property?
3. What tradeoff would matter if this were a real production system?

This keeps the project focused on distributed-systems reasoning instead of only implementing code that passes a test.
