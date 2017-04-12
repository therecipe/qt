//author: https://github.com/5k3105

package main

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	Scene     *widgets.QGraphicsScene
	View      *widgets.QGraphicsView
	Item      *widgets.QGraphicsPixmapItem
	r, g, b   int
	statusbar *widgets.QStatusBar
	mp        bool
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Main Window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Sprite Editor")
	window.SetMinimumSize2(360, 520)

	// Statusbar
	statusbar = widgets.NewQStatusBar(window)
	window.SetStatusBar(statusbar)

	Scene = widgets.NewQGraphicsScene(nil)
	View = widgets.NewQGraphicsView(nil)

	Scene.ConnectKeyPressEvent(keyPressEvent)
	Scene.ConnectWheelEvent(wheelEvent)
	View.ConnectResizeEvent(resizeEvent)

	dx, dy := 16, 32

	img := gui.NewQImage3(dx, dy, gui.QImage__Format_ARGB32)

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			img.SetPixelColor2(i, j, gui.NewQColor3(i*2, j*8, i*2, 255))

		}
	}

	//img = img.Scaled2(dx*2,dy,core.Qt__IgnoreAspectRatio, core.Qt__FastTransformation)

	Item = widgets.NewQGraphicsPixmapItem2(gui.NewQPixmap().FromImage(img, 0), nil)

	Item.ConnectMouseMoveEvent(ItemMouseMoveEvent)
	Item.ConnectMousePressEvent(ItemMousePressEvent)
	Item.ConnectMouseReleaseEvent(ItemMouseReleaseEvent)

	Item.SetAcceptHoverEvents(true)
	Item.ConnectHoverMoveEvent(ItemHoverMoveEvent)

	Scene.AddItem(Item)

	View.SetScene(Scene)
	View.Show()

	statusbar.ShowMessage(core.QCoreApplication_ApplicationDirPath(), 0)

	// Set Central Widget
	window.SetCentralWidget(View)

	// Run App
	widgets.QApplication_SetStyle2("fusion")
	window.Show()
	widgets.QApplication_Exec()
}

func ItemMousePressEvent(event *widgets.QGraphicsSceneMouseEvent) {
	mp = true
	mousePosition := event.Pos()
	x, y := int(mousePosition.X()), int(mousePosition.Y())
	drawpixel(x, y)

}

func ItemMouseReleaseEvent(event *widgets.QGraphicsSceneMouseEvent) {
	mp = false

	Item.MousePressEventDefault(event) // absofukinlutely necessary for drag & draw !!

	//Item.MouseReleaseEventDefault(event) // worthless
}

func ItemMouseMoveEvent(event *widgets.QGraphicsSceneMouseEvent) {
	mousePosition := event.Pos()
	x, y := int(mousePosition.X()), int(mousePosition.Y())

	drawpixel(x, y)

}

func ItemHoverMoveEvent(event *widgets.QGraphicsSceneHoverEvent) {
	mousePosition := event.Pos()
	x, y := int(mousePosition.X()), int(mousePosition.Y())

	rgbValue := Item.Pixmap().ToImage().PixelColor2(x, y)
	r, g, b := rgbValue.Red(), rgbValue.Green(), rgbValue.Blue()
	statusbar.ShowMessage("x: "+strconv.Itoa(x)+" y: "+strconv.Itoa(y)+" r: "+strconv.Itoa(r)+" g: "+strconv.Itoa(g)+" b: "+strconv.Itoa(b), 0)

}

func drawpixel(x, y int) {

	if mp {
		img := Item.Pixmap().ToImage()
		img.SetPixelColor2(x, y, gui.NewQColor3(255, 255, 255, 255))
		Item.SetPixmap(gui.NewQPixmap().FromImage(img, 0))
	}

}

func keyPressEvent(e *gui.QKeyEvent) {

	if e.Modifiers() == core.Qt__ControlModifier {
		switch int32(e.Key()) {
		case int32(core.Qt__Key_Equal):
			View.Scale(1.25, 1.25)

		case int32(core.Qt__Key_Minus):
			View.Scale(0.8, 0.8)
		}
	}

}

func wheelEvent(e *widgets.QGraphicsSceneWheelEvent) {
	if e.Modifiers() == core.Qt__ControlModifier {
		if e.Delta() > 0 {
			View.Scale(1.25, 1.25)
		} else {
			View.Scale(0.8, 0.8)
		}
	}
}

func resizeEvent(e *gui.QResizeEvent) {

	View.FitInView(Scene.ItemsBoundingRect(), core.Qt__KeepAspectRatio)

}
