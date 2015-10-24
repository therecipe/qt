package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {

	gui.NewQGuiApplication(0, "")

	var view = quick.NewQQuickView(nil)
	view.SetSource("qml/view.qml")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
