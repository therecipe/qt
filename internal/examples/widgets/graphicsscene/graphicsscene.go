//author: https://github.com/5k3105

package main

import (
	"os"
	"runtime"

	"github.com/StarAurryon/qt/gui"
	"github.com/StarAurryon/qt/widgets"
)

var (
	canvas *widgets.QWidget
	scene  *widgets.QGraphicsScene
	view   *widgets.QGraphicsView
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	NewCanvas().Show()

	widgets.QApplication_Exec()
}

func NewCanvas() *widgets.QWidget {
	canvas = widgets.NewQWidget(nil, 0)
	scene = widgets.NewQGraphicsScene(nil)
	view = widgets.NewQGraphicsView(nil)

	var font = gui.NewQFont2("Meiryo", 20, 2, false)
	if runtime.GOARCH == "js" || runtime.GOARCH == "wasm" {
		scene.AddText("Hello World", font)
	} else {
		scene.AddText("Hello 世界", font)
	}

	var color = gui.NewQColor3(255, 0, 0, 255)
	var pen = gui.NewQPen3(color)

	scene.AddLine2(0, scene.Height(), scene.Width(), scene.Height(), pen)

	view.SetScene(scene)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(view, 0, 0)

	canvas.SetLayout(layout)

	return canvas
}
