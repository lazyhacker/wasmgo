#!/bin/sh


if [ -x app.wasm ]; then
    rm app.wasm
fi

GOROOT=$HOME/go-wasm
GOOS=js GOARCH=wasm $HOME/go-wasm/bin/go build -o app.wasm app.go
