package main

import (
	"math"
	"strconv"
	"syscall/js"
	"time"

	"lazyhackergo.com/browser"
)

var signal = make(chan int)

func cb(args []js.Value) {
	println("callback")
}

func cbQuit(e js.Value) {
	println("got Quit event callback!")
	window := browser.GetWindow()
	window.Document.GetElementById("runButton").SetProperty("disabled", false)
	window.Document.GetElementById("quit").SetAttribute("disabled", true)
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

	window := browser.GetWindow()

	window.Document.GetElementById("runButton").SetAttribute("disabled", true)

	window.Document.GetElementById("quit").AddEventListener(browser.EventClick, q)
	window.Document.GetElementById("quit").SetProperty("disabled", false)
	//js.Global.Get("document").Call("getElementById", "quit").Call("addEventListener", "click", js.ValueOf(q))

	window.Alert("Triggered from Go WASM module")
	window.Console.Info("hello, browser console")

	canvas, err := window.Document.GetElementById("testcanvas").ToCanvas()
	if err != nil {
		window.Console.Warn(err.Error())
	}

	canvas.Clear()
	canvas.BeginPath()
	canvas.Arc(100, 75, 50, 0, 2*math.Pi, false)
	canvas.Stroke()
	canvas.Font("30px Arial")

	time.Sleep(5 * time.Second)
	go func() {
		for i := 0; i < 10; i++ {

			canvas.Clear()
			canvas.FillText(strconv.Itoa(i), 10, 50)
			time.Sleep(1 * time.Second) // sleep allows the browser to take control otherwise the whole UI gets frozen.
		}
		canvas.Clear()
		canvas.FillText("Stop counting!", 10, 50)

	}()

	keepalive()
}
