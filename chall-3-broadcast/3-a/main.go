package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReq struct {
	Type    string `json:"type"`
	Message int    `json:"message"`
}

func main() {

	seen := make(map[int]struct{})
	msgs := make([]int, 0)
	var mu sync.RWMutex

	n := maelstrom.NewNode()

	//broadcast
	n.Handle("broadcast", func(msg maelstrom.Message) error {

		var rq BroadcastReq
		if err := json.Unmarshal(msg.Body, &rq); err != nil {
			return err
		}

		mu.Lock()
		if _, exists := seen[rq.Message]; !exists {
			seen[rq.Message] = struct{}{}
			msgs = append(msgs, rq.Message)
		}
		mu.Unlock()

		return n.Reply(msg, map[string]string{
			"type": "broadcast_ok",
		})
	})

	//read
	n.Handle("read", func(msg maelstrom.Message) error {
		mu.RLock()
		result := make([]int, len(msgs))
		copy(result, msgs)
		mu.RUnlock()

		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": result,
		})
	})

	//topology
	n.Handle("topology", func(msg maelstrom.Message) error {
		return n.Reply(msg, map[string]string{
			"type": "topology_ok",
		})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
