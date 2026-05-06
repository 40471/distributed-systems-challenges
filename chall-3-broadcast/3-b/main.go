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

	store := func(message int) bool {
		mu.Lock()
		defer mu.Unlock()

		if _, exists := seen[message]; exists {
			return false
		}

		seen[message] = struct{}{}
		msgs = append(msgs, message)
		return true
	}

	gossip := func(message int, exclude string) error {
		for _, nodeID := range n.NodeIDs() {
			if nodeID == n.ID() || nodeID == exclude {
				continue
			}

			if err := n.Send(nodeID, map[string]any{
				"type":    "gossip",
				"message": message,
			}); err != nil {
				return err
			}
		}

		return nil
	}

	//broadcast
	n.Handle("broadcast", func(msg maelstrom.Message) error {

		var rq BroadcastReq
		if err := json.Unmarshal(msg.Body, &rq); err != nil {
			return err
		}

		if store(rq.Message) {
			if err := gossip(rq.Message, ""); err != nil {
				return err
			}
		}

		return n.Reply(msg, map[string]string{
			"type": "broadcast_ok",
		})
	})

	//gossip
	n.Handle("gossip", func(msg maelstrom.Message) error {
		var rq BroadcastReq
		if err := json.Unmarshal(msg.Body, &rq); err != nil {
			return err
		}

		if store(rq.Message) {
			return gossip(rq.Message, msg.Src)
		}

		return nil
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
