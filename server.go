package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
)

var port = flag.Int("port", 8080, "server port")

var fileserver = http.FileServer(http.Dir("."))

func main() {
	mime.AddExtensionType(".wasm", "application/wasm")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), fileserver))
}
