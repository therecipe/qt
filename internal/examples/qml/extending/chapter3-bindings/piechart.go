package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
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
	p.ConnectColorChanged(p.colorChanged)
}

func (p *PieChart) paint(painter *gui.QPainter) {
	pen := gui.NewQPen3(p.Color())
	pen.SetWidth(2)
	painter.SetPen(pen)
	painter.SetRenderHints(gui.QPainter__Antialiasing, true)
	painter.DrawPie3(core.NewQRect4(0, 0, int(p.Width()), int(p.Height())).Adjusted(1, 1, -1, -1), 90*16, 290*16)
}

func (p *PieChart) colorChanged(color *gui.QColor) {
	p.UpdateDefault() //TODO: why default?
}
