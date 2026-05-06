# Challenge 1: Echo

## Goal

Implement a basic Maelstrom node that receives an `echo` message and replies with the same body, changing only the message type to `echo_ok`.

This challenge is mainly about learning how Maelstrom communicates with a node process. Maelstrom sends JSON messages through standard input and expects JSON responses through standard output.

## What I Learned

- How Maelstrom models a distributed node as a standalone process.
- How the Go Maelstrom library registers message handlers with `Node.Handle`.
- How `Node.Reply` automatically fills in response metadata such as `src`, `dest`, and `in_reply_to`.
- How challenge workloads are verified by Maelstrom instead of traditional unit tests.

## Approach

The solution keeps the request body as a generic `map[string]any`, changes the `type` field from `echo` to `echo_ok`, and sends the body back with `Node.Reply`.

This is the simplest correct approach because the challenge asks us to preserve the original body and only modify the response type.

## Tradeoffs

Using a generic map is flexible and fits this challenge well. For larger challenges, typed request and response structs are usually better because they make the protocol clearer and catch mistakes earlier.

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

- Use typed structs in later challenges where the message shape is more meaningful.
- Keep this challenge minimal because its value is understanding the Maelstrom protocol.
