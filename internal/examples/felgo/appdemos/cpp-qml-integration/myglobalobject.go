package main

import (
	"github.com/therecipe/qt/core"
)

type MyGlobalObject struct {
	core.QObject

	_ func() `constructor:"init"`

	_ int `property:"counter"` // this makes counter available as a QML property

	_ func(text string) `slot:"doSomething,auto"` // slots are public methods available in QML
}

func (o *MyGlobalObject) init() {
	o.SetCounter(0)
}

func (o *MyGlobalObject) doSomething(text string) {
	println("MyGlobalObject doSomething called with", text)
	o.SetCounter(o.Counter() + 1)
}
