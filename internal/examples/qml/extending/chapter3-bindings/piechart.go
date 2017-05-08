package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

type PieChart struct {
	quick.QQuickPaintedItem

	_ string      `property:"name"`
	_ *gui.QColor `property:"color"`
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

type PieChartFactory struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(p *PieChart) `slot:"create"`
}

func (p *PieChartFactory) init() {
	p.ConnectCreate(p.create)
}

func (p *PieChartFactory) create(pie *PieChart) {
	pie.ConnectPaint(pie.paint)
	pie.ConnectColorChanged(pie.colorChanged)
}
