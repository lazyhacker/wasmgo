This is some experimental code for working with Go and WebAssembly.

## Prerequisite

As of this writing, the Go release doesn't support WebAssembly so the version of
Go from neelance is needed.

See https://blog.lazyhacker.com/2018/05/webassembly-wasm-with-go.html for how to
build Go with WASM.

## Install

```
go get github.com/lazyhacker/wasmgo
cd $GOPATH/src/github.com/lazyhacker/wasmgo
./build.sh
```

Note: build.sh assumes that the 'go' command is located in $HOME/go-wasm/bin.
Make sure GOROOT is pointing to the 'go' command that has WASM support.

## Running

The code needs to be served through a web server.

```
go install lazyhackergo/simplehttpserver
$GOPATH/bin/simplehttpserver
```
