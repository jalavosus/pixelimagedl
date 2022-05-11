GO=$(shell which go)
CMD=./cmd/pixelimagedl
BIN=./bin

OUT=$(BIN)/main

.PHONY: build clean

build : clean
	$(GO) build -o $(OUT) $(CMD)

clean :
	rm -rf $(OUT)


.PHONY: lint fmt

lint :
	golangci-lint run --config=.golangci.yml

fmt :
	gofmt -s -w ./

.PHONY: mod-download mod-tidy

mod-download :
	$(GO) mod download

mod-tidy :
	$(GO) mod tidy