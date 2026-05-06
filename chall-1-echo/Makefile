BINARY := bin/maelstrom-echo
MAELSTROM ?= ../maelstrom/maelstrom

.PHONY: build test-maelstrom clean

build:
	go build -o $(BINARY) .

test-maelstrom: build
	@test -x "$(MAELSTROM)" || (echo "Set MAELSTROM=/path/to/maelstrom" && exit 1)
	$(MAELSTROM) test -w echo --bin ./$(BINARY) --node-count 1 --time-limit 10

clean:
	rm -rf bin store
