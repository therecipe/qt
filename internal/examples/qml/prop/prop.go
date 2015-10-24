package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

func main() {

	core.NewQCoreApplication(0, "")

	var (
		component = qml.NewQQmlComponent3(qml.NewQQmlEngine(nil), "qml/prop.qml", nil)
		object    = component.Create(nil)
	)

	fmt.Println("Property value:", object.Property("someNumber"))
	object.SetProperty("someNumber", "5000")
	fmt.Println("Property value:", object.Property("someNumber"))
	object.SetProperty("someNumber", "100")
}
