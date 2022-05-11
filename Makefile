.PHONY: build clean

GO=$(shell which go)
CMD=./cmd
BIN=./bin

OUT=$(BIN)/main

build : clean
	$(GO) build -o $(OUT) $(CMD)

clean :
	rm -rf $(OUT)


.PHONY: mod-download mod-tidy

mod-download :
	$(GO) mod download

mod-tidy :
	$(GO) mod tidy