// +build js,wasm

// Package browser wraps syscall/js.
package browser

import (
	"syscall/js"
)

var global = js.Global

type window struct {
	Console debug
}

// GetWindow returns the main browser window object.
func Window() window {
	return window{
		Console: debug{console: global.Get("console")},
	}
}

// Alert shows an alert box in the browser with an OK button.
func (window) Alert(m string) {
	global.Call("alert", m)
}

func OnClick(id string, f func(js.Value)) func() {
	cb := js.NewEventCallback(false, false, false, f)
	js.Global.Get("document").Call("getElementById", id).Call("addEventListener", "click", js.ValueOf(cb))
	return cb.Close
}

// Set sets a property of a DOM element to the given value.
func Set(element, property, value string) {
	js.Global.Get("document").Call("getElementById", element).Set(property, js.ValueOf(value))
}

func Sometime(f func([]js.Value)) func() {
	c := js.NewCallback(f)
	js.ValueOf(c).Invoke()
	return c.Close
}

func ServeForever() {
	select {}
}
