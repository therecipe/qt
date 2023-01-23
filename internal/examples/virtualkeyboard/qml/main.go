package main

import (
	"os"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/qml"
	"github.com/bluszcz/cutego/quickcontrols2"
)

func main() {
	os.Setenv("QT_IM_MODULE", "qtvirtualkeyboard")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("Material")
	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
