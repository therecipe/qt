package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/application.qml", 0))

	gui.QGuiApplication_Exec()
}
