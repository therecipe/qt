package main

import (
	"github.com/therecipe/qt/core"
)

func init() { MyQMLType_QmlRegisterType2("com.yourcompany.xyz", 1, 0, "MyQMLType") }

type MyQMLType struct {
	core.QObject

	_ string `property:"message"` // this makes message available as a QML property

	_ func(value int) int `slot:"increment,auto"`
	_ func()              `slot:"startCppTask,auto"` // starts internal calculations of doCppTask()

	_ func() `signal:"cppTaskFinished"` // triggered after calculations in doCppTask()
}

func (o *MyQMLType) init() {
	o.SetMessage("")
}

func (o *MyQMLType) increment(value int) int {
	return value + 1
}

func (o *MyQMLType) startCppTask() {
	// NOTE: you can do calculations here in another thread, this may be used to perform
	// cpu-intense operations for e.g. AI (artificial itelligence), Machine Learning or similar purposes
	// When the work is done, we can trigger the cppTaskFinished signal and react anywhere in C++ or QML
	o.doCppTask()
}

func (o *MyQMLType) doCppTask() {
	o.CppTaskFinished()
}
