SRC := main.go
PKG := ./cmd ./build ./util

BUILD_DIR := build

CONTRACT := ./contracts/SimpleCounter.fc

GO := go

.PHONY: all build clean run compile

all: build

build:
	$(GO) build -o tonc $(SRC)

run:
	./tonc build --contract $(CONTRACT) --boc --json --hex --verbose

clean:
	rm -rf $(BUILD_DIR) tonc

compile: build
	./tonc build --contract $(CONTRACT) --boc --json --hex --verbose

compile-all: build
	./tonc build --src ./contracts --boc --json --hex --verbose
