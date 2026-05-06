BINARY := bin/maelstrom-broadcast
MAELSTROM ?= ../../maelstrom/maelstrom

.PHONY: build test-maelstrom clean

build:
	go build -o $(BINARY) .

test-maelstrom: build
	@test -x "$(MAELSTROM)" || (echo "Set MAELSTROM=/path/to/maelstrom" && exit 1)
	$(MAELSTROM) test -w broadcast --bin ./$(BINARY) --node-count 5 --time-limit 20 --rate 10

clean:
	rm -rf bin store
