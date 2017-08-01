package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Widget struct {
	widgets.QWidget

	_ func() `constructor:"init"`

	_ func() `slot:"animate"`

	Helper  *Helper
	elapsed int
}

func (w *Widget) init() {
	w.SetFixedSize2(200, 200)

	w.ConnectAnimate(w.animate)
	w.ConnectPaintEvent(w.paintEvent)
}

func (w *Widget) animate() {
	w.elapsed = (w.elapsed + 50) % 1000
	w.Update()
}

func (w *Widget) paintEvent(event *gui.QPaintEvent) {
	painter := gui.NewQPainter2(w)
	painter.SetRenderHint(gui.QPainter__Antialiasing, true)
	w.Helper.Paint(painter, event, w.elapsed)
	painter.DestroyQPainter()
}
