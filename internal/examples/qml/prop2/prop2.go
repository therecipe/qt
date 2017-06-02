package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

type QmlBride struct {
	core.QObject

	_ bool          `property:"boolProp"`
	_ int           `property:"intProp"`
	_ string        `property:"stringProp"`
	_ []string      `property:"stringListProp"`
	_ *core.QObject `property:"objectProp"`

	_ func() `constructor:"init"`
}

func (bridge *QmlBride) init() {
	bridge.ConnectBoolPropChanged(func(boolProp bool) {
		fmt.Println(" go: changed bool ->", boolProp)
	})
	bridge.ConnectIntPropChanged(func(intProp int) {
		fmt.Println(" go: changed int ->", intProp)
	})
	bridge.ConnectStringPropChanged(func(stringProp string) {
		fmt.Println(" go: changed string ->", stringProp)
	})
	bridge.ConnectStringListPropChanged(func(stringListProp []string) {
		fmt.Println(" go: changed stringlist ->", stringListProp)
	})
	bridge.ConnectObjectPropChanged(func(objectProp *core.QObject) {
		fmt.Println(" go: changed object ->", objectProp.ObjectName())
	})

	bridge.SetBoolProp(true)
	bridge.SetIntProp(123)
	bridge.SetStringProp("hello")
	bridge.SetStringListProp([]string{"hello", "world"})
	var obj = core.NewQObject(nil)
	obj.SetObjectName("objectName")
	bridge.SetObjectProp(obj)
}

func main() {

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var bridge = NewQmlBride(nil)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.RootContext().SetContextProperty("bridge", bridge)

	app.Load(core.NewQUrl3("qrc:/qml/prop2.qml", 0))

	gui.QGuiApplication_Exec()
}
