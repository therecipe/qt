//source: https://doc.qt.io/qt-5/qtqml-referenceexamples-adding-example.html

package main

import (
	"os"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	core.NewQCoreApplication(len(os.Args), os.Args)

	Person_QmlRegisterType2("People", 1, 0, "Person")

	engine := qml.NewQQmlEngine(nil)
	component := qml.NewQQmlComponent5(engine, core.NewQUrl3("qrc:example.qml", 0), nil)
	person := NewPersonFromPointer(component.Create(nil).Pointer())
	if person.Pointer() != nil {
		println("The person's name is", person.Name())
		println("They wear a", person.ShoeSize(), "sized shoe")
	} else {
		for _, e := range component.Errors() {
			println(e.Description())
		}
	}

	core.QCoreApplication_Exec()
}
