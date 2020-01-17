package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	widgets.QWidget

	_ func() `constructor:"init"`

	helper *Helper
}

func (w *Window) init() {
	w.helper = NewHelper(w)

	w.SetWindowTitle("2D Painting on Native and OpenGL Widgets")

	native := NewWidget(nil, 0)
	native.Helper = w.helper

	openGL := NewGLWidget(w, 0)
	openGL.Helper = w.helper

	nativeLabel := widgets.NewQLabel2("Native", w, 0)
	nativeLabel.SetAlignment(core.Qt__AlignHCenter)
	openGLLabel := widgets.NewQLabel2("OpenGL", w, 0)
	openGLLabel.SetAlignment(core.Qt__AlignHCenter)

	layout := widgets.NewQGridLayout(nil)
	layout.AddWidget2(native, 0, 0, 0)
	layout.AddWidget2(openGL, 0, 1, 0)
	layout.AddWidget2(nativeLabel, 1, 0, 0)
	layout.AddWidget2(openGLLabel, 1, 1, 0)
	w.SetLayout(layout)

	timer := core.NewQTimer(nil)
	timer.ConnectTimeout(native.Animate)
	timer.ConnectTimeout(openGL.Animate)
	timer.Start(50)
}
