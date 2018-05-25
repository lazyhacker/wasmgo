package main

import (
	"fmt"
	"syscall/js"
	"time"

	"lazyhackergo.com/wasmgo/lib/browser"
)

var signal = make(chan int)
var counter int

func main() {
	window := browser.Window()
	window.Console.Info("go main() started")

	done := browser.Sometime(func(_ []js.Value) {
		println("a callback")
	})
	defer done()

	// Increment button
	done = browser.OnClick("increment", func(e js.Value) {
		counter++
		browser.Set("counter", "textContent", fmt.Sprint(counter))
	})
	defer done()

	done = browser.OnClick("alert", func(e js.Value) {
		window.Alert("alert button pressed")
	})
	defer done()

	window.Console.Info("starting ping goroutine")
	go func() {
		clock := 0
		for {
			time.Sleep(1 * time.Second)
			browser.Set("pings", "textContent", fmt.Sprintf("Goroutine clock has run for %d seconds", clock))
			clock++
		}
	}()
	window.Console.Info("ping goroutine started")

	// Keep-alive (when all gorountines are blocked, Go yields back to JS)
	browser.ServeForever()

	window.Console.Info("go main() terminating")
}
