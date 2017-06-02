//source: https://github.com/KDE/kirigami

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	if quickcontrols2.QQuickStyle_Name() == "" {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	}

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/contents/ui/ExampleApp.qml", 0))

	gui.QGuiApplication_Exec()
}
