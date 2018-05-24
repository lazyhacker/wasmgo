#!/bin/sh


if [ -x test.wasm ]; then
    rm test.wasm
fi

GOOS=js GOARCH=wasm $HOME/go-wasm/bin/go build -o test.wasm test.go
