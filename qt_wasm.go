// +build js,wasm

package qt

import "syscall/js"

var WASM = js.Global().Call("eval", "Module")
