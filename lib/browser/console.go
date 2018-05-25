// +build js,wasm

package browser

import "syscall/js"

// Console provides access to the browsers debugging console.
type debug struct {
	console js.Value
}

// Clear clears the console.
func (d debug) Clear() {
	d.console.Call("clear")
}

// Count  logs the number of times this count has been called.
func (d debug) Count() int {
	return d.console.Call("count").Int()
}

// Error writes an error message to the console.
func (d debug) Error(m string) {
	d.console.Call("error", m)
}

// Info writes an information message to the console.
func (d debug) Info(m string) {
	d.console.Call("info", m)

}

// Warn writes an warning message to the console.
func (d debug) Warn(m string) {
	d.console.Call("warn", m)
}

// Trace writes the stack trace to the console.
func (d debug) Trace() {
	d.console.Call("trace")
}
