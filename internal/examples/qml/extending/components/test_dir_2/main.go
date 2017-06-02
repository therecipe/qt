package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"

	_ "github.com/therecipe/qt/internal/examples/qml/extending/components/test_dir_2/component"
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
