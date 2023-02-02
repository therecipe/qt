package main

import (
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/quick"
)

func init() {
	SomeComponentInternal_QmlRegisterType2("CustomComponentsInternal", 1, 0, "SomeComponent")
}

type SomeComponentInternal struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`

	_ *gui.QColor `property:"color"`
}

func (p *SomeComponentInternal) init() {
	p.SetWidth(200)
	p.SetHeight(200)
	p.SetColor(gui.NewQColor2(core.Qt__red))
	p.ConnectPaint(p.paint)
}

func (p *SomeComponentInternal) paint(painter *gui.QPainter) {
	pen := gui.NewQPen3(p.Color())
	pen.SetWidth(2)
	painter.SetPen(pen)
	painter.SetRenderHints(gui.QPainter__Antialiasing, true)
	painter.FillRect5(0, 0, int(p.Width()), int(p.Height()), p.Color())
}
