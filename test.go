package main

import (
	"math"
	"syscall/js"

	"lazyhackergo.com/browser"
)

var signal = make(chan int)

func cb(args []js.Value) {
	println("callback")
}

func cbQuit(e js.Value) {
	println("got Quit event callback!")
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
	q := js.NewEventCallback(false, false, false, cbQuit)
	defer q.Close()

	c := js.NewCallback(cb)
	defer c.Close()
	browser.Invoke(c)
	//js.ValueOf(c).Invoke()

	window := browser.Window()

	window.Document.GetElementById("quit").AddEventListener(browser.EventClick, q)
	//js.Global.Get("document").Call("getElementById", "quit").Call("addEventListener", "click", js.ValueOf(q))

	window.Alert("hello, browser")
	window.Console.Info("hello, browser console")

	canvas, err := window.Document.GetElementById("testcanvas").ToCanvas()
	if err != nil {
		window.Console.Warn(err.Error())
	}

	canvas.BeginPath()
	canvas.Arc(100, 75, 50, 0, 2*math.Pi, false)
	canvas.Stroke()
	keepalive()
}
