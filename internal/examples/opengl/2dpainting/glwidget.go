package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type GLWidget struct {
	widgets.QOpenGLWidget

	_ func() `constructor:"init"`

	_ func() `slot:"animate"`

	Helper  *Helper
	elapsed int
}

func (w *GLWidget) init() {
	w.SetFixedSize2(200, 200)
	w.SetAutoFillBackground(false)

	w.ConnectAnimate(w.animate)
	w.ConnectPaintEvent(w.paintEvent)
}

func (w *GLWidget) animate() {
	w.elapsed = (w.elapsed + 50) % 1000
	w.Update()
}

func (w *GLWidget) paintEvent(event *gui.QPaintEvent) {
	painter := gui.NewQPainter2(w)
	painter.SetRenderHint(gui.QPainter__Antialiasing, true)
	w.Helper.Paint(painter, event, w.elapsed)
	painter.DestroyQPainter()
}
