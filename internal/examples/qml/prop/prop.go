package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

func main() {

	core.NewQCoreApplication(len(os.Args), os.Args)

	var (
		component = qml.NewQQmlComponent5(qml.NewQQmlEngine(nil), core.NewQUrl3("qrc:/qml/prop.qml", 0), nil)
		object    = component.Create(nil)
	)

	var someBool bool
	fmt.Println("Property value:", object.Property("someNumber").ToInt(&someBool), someBool)
	object.SetProperty("someNumber", core.NewQVariant7(5000))
	fmt.Println("Property value:", object.Property("someNumber").ToInt(&someBool), someBool)
	object.SetProperty("someNumber", core.NewQVariant7(200))
	fmt.Println("Property value:", object.Property("someNumber").ToInt(&someBool), someBool)
}
