//source: http://doc.qt.io/qt-5/qtgui-rasterwindow-example.html

package main

import (
	"os"

	"github.com/therecipe/qt/gui"
)

func main() {
	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var window = NewRasterWindow(nil)
	window.Show()

	app.Exec()
}
