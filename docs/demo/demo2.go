// Ivy is a APL-like calculator written by Rob Pike.  This brings it over to
// webassembly.
package main

import (
	"syscall/js"

	"lazyhackergo.com/browser"
	"robpike.io/ivy/config"
	"robpike.io/ivy/mobile"
	"robpike.io/ivy/value"
)

var signal = make(chan int)

var (
	conf    config.Config
	context value.Context
)

// START DEMO2_1 OMIT
func main() {

	ivy := js.NewEventCallback(false, false, false, cbRunIvy)
	defer ivy.Close()

	window := browser.GetWindow()

	express := window.Document.GetElementById("expression")
	express.AddEventListener(browser.EventKeyUp, ivy)
	window.Document.GetElementById("loadspinner").SetAttribute("class", "")
	express.Focus()

	keepalive()
}

// END DEMO2_1 OMIT

func cbClear(e js.Value) {

	window := browser.GetWindow()
	element := window.Document.GetElementById("ivy-out")
	element.SetInnerHTML("")
	express := window.Document.GetElementById("expression")
	express.Focus()

}

//START DEMO2_2 OMIT
func cbRunIvy(e js.Value) {

	window := browser.GetWindow()
	if e.Get("keyCode").Int() == 13 {
		express := window.Document.GetElementById("expression")

		expr := express.Value()
		res, err := mobile.Eval(expr) // HL
		if err != nil {
			window.Console.Warn(err.Error())
			return
		}

		element := window.Document.GetElementById("ivy-out")
		content := element.InnerHTML()
		element.SetInnerHTML(content + "> " + expr + "<br/>" + res + "<br/>")
		express.SetValue("")

		window.ScrollTo(0, window.InnerHeight())
	}
}

//END DEMO2_2 OMIT

func keepalive() {
	for {
		m := <-signal
		if m == 0 {
			println("quit signal received")
			break
		}
	}
}
