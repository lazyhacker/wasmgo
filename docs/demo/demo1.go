package main // import "lazyhackergo.com/wasmgo"

import (
	"math"
	"strconv"
	"syscall/js"
	"time"

	"lazyhackergo.com/browser"
)

var signal = make(chan int)

// Example of a callback function that can be registered.
func cb(args []js.Value) {
	println("callback")
}

// START DEMO1_2 OMIT
func cbQuit(e js.Value) {
	window := browser.GetWindow()
	window.Document.GetElementById("runButton").SetProperty("disabled", false)
	window.Document.GetElementById("quit").SetAttribute("disabled", true)
	signal <- 0
}

//END DEMO1_2 OMIT

// keepalive waits for a specific value as a signal to quit.  Browsers still run
// javascript and wasm in a single thread so until the wasm module releases
// control, the entire browser window is blocked waiting for the module to
// finish.  It looks like that while waiting for the blocked channel, the
// browser window gets control back and can continue its event loop.
// START DEMO1_4 OMIT
func keepalive() {
	for {
		m := <-signal
		if m == 0 {
			println("quit signal received")
			break
		}
	}
	// or select {} if never meant to end module.
}

// END DEMO1_4 OMIT

// START DEMO1_1 OMIT
func main() {
	q := js.NewEventCallback(false, false, false, cbQuit) // HL
	defer q.Close()

	window := browser.GetWindow()
	runButton := Window.Document.GetElementById("runButton")
	runButton.SetAttribute("disabled", true)

	quitButton := window.Document.GetElementById("quit")
	quitButton.AddEventListener(browser.EventClick, q) // HL
	quitButton.SetProperty("disabled", false)

	window.Alert("Triggered from Go WASM module")
	window.Console.Info("hello, browser console")
	// END DEMO1_1 OMIT

	// START DEMO1_3 OMIT
	canvas, _ := window.Document.GetElementById("testcanvas").ToCanvas() // HL

	canvas.Clear()
	canvas.BeginPath()
	canvas.Arc(100, 75, 50, 0, 2*math.Pi, false)
	canvas.Stroke()

	canvas.Font("30px Arial")
	go func() { // HL
		for i := 0; i < 10; i++ {
			canvas.Clear()
			canvas.FillText(strconv.Itoa(i), 10, 50)
			time.Sleep(1 * time.Second) // HL
		}
		canvas.Clear()
		canvas.FillText("Stop counting!", 10, 50)
	}()
	window.Console.Info("Go routine running while this message is printed to the console.")
	// END DEMO1_3 OMIT

	keepalive()
}
