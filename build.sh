#!/bin/sh


if [ -x test.wasm ]; then
    rm test.wasm
fi

GOROOT=$HOME/go-wasm
GOOS=js GOARCH=wasm $HOME/go-wasm/bin/go build -o test.wasm test.go
