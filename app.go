package main

import (
	"syscall/js"

	"lazyhackergo.com/browser"
)

var signal = make(chan int)

func cb(args []js.Value) {
	println("callback")
}

func quitCallback(e js.Value) {
	window := browser.Window()
	window.Alert("Imma quiting!")
	println("got event callback!")
	signal <- 0
}

func keepalive() {
	for {
		m := <-signal
		if m == 0 {
			println("quit signal received")
			break
		}
	}
}

func main() {
	q := js.NewEventCallback(false, false, false, quitCallback)
	defer q.Close()

	c := js.NewCallback(cb)
	defer c.Close()
	js.ValueOf(c).Invoke()

	js.Global.Get("document").Call("getElementById", "quit").Call("addEventListener", "click", js.ValueOf(q))

	window := browser.Window()

	window.Alert("hello, browser")
	window.Console.Info("hello, browser console")

	keepalive()
}
