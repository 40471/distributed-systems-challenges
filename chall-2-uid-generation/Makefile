BINARY := bin/maelstrom-unique-ids
MAELSTROM ?= ../maelstrom/maelstrom

.PHONY: build test test-maelstrom clean

build:
	go build -o $(BINARY) .

test:
	go test ./...

test-maelstrom: build
	@test -x "$(MAELSTROM)" || (echo "Set MAELSTROM=/path/to/maelstrom" && exit 1)
	$(MAELSTROM) test -w unique-ids --bin ./$(BINARY) --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

clean:
	rm -rf bin store
