This is an experimental app written in Go on the server and client.

# Setup

These instructions install an experimental Go compiler in `$HOME/go-wasm`. You will need a [Go compiler](https://golang.org/) installed on your system.

```bash
$ git clone --branch wasm-wip https://github.com/neelance/go.git $HOME/go-wasm
$ cd $HOME/go-wasm/src && ./all.bash
```

To use the WASM-capable compiler for the duration of a terminal session:

```bash
$ GOROOT="$HOME/go-wasm"
$ alias go="$HOME/go-wasm/bin/go"
```

IDEs will need to be configured to use this compiler. For VS Code, add the following to your workspace settings (you'll need to expand out $HOME yourself):

```json
"go.goroot": "$HOME/go-wasm"
```

See https://blog.lazyhacker.com/2018/05/webassembly-wasm-with-go.html for details on setting up the Go WASM compiler.

# Start up

Start the server from the repo directory. By default the server listens on port 8080.

```bash
$ go run server.go
```

Connect to the app by navigating to [localhost:8080](http://localhost:8080/).

