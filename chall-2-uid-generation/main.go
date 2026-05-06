package main

import (
	"fmt"
	"log"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	var counter atomic.Uint64

	n.Handle("generate", func(msg maelstrom.Message) error {
		id := counter.Add(1)

		return n.Reply(msg, map[string]any{
			"type": "generate_ok",
			"id":   fmt.Sprintf("%s-%d", n.ID(), id),
		})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
