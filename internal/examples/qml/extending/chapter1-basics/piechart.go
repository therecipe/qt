package main

import (
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/quick"
)

func init() {
	PieChart_QmlRegisterType2("Charts", 1, 0, "PieChart")
}

type PieChart struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`

	_ string      `property:"name"`
	_ *gui.QColor `property:"color"`
}

func (p *PieChart) init() {
	p.ConnectPaint(p.paint)
}

func (p *PieChart) paint(painter *gui.QPainter) {
	pen := gui.NewQPen3(p.Color())
	pen.SetWidth(2)
	painter.SetPen(pen)
	painter.SetRenderHints(gui.QPainter__Antialiasing, true)
	painter.DrawPie3(core.NewQRect4(0, 0, int(p.Width()), int(p.Height())).Adjusted(1, 1, -1, -1), 90*16, 290*16)
}
