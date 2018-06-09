// +build js

package qt

import "github.com/gopherjs/gopherjs/js"

var WASM = js.Global.Call("eval", "Module")
