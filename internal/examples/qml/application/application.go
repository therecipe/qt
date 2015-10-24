package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {

	gui.NewQGuiApplication(0, "")

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load("qml/application.qml")

	gui.QGuiApplication_Exec()
}
