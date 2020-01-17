package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func init() {
	SomeComponent_QmlRegisterType2("CustomComponents", 1, 0, "SomeComponent")
}

type SomeComponent struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`

	_ *gui.QColor `property:"color"`
}

func (p *SomeComponent) init() {
	p.SetWidth(200)
	p.SetHeight(200)
	p.SetColor(gui.NewQColor2(core.Qt__red))
	p.ConnectPaint(p.paint)
}

func (p *SomeComponent) paint(painter *gui.QPainter) {
	pen := gui.NewQPen3(p.Color())
	pen.SetWidth(2)
	painter.SetPen(pen)
	painter.SetRenderHints(gui.QPainter__Antialiasing, true)
	painter.FillRect5(0, 0, int(p.Width()), int(p.Height()), p.Color())
}
