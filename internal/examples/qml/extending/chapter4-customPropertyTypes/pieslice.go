package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func init() {
	PieSlice_QmlRegisterType2("Charts", 1, 0, "PieSlice")
}

type PieSlice struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`

	_ *gui.QColor `property:"color"`
}

func (p *PieSlice) init() {
	p.ConnectPaint(p.paint)
}

func (p *PieSlice) paint(painter *gui.QPainter) {
	pen := gui.NewQPen3(p.Color())
	pen.SetWidth(2)
	painter.SetPen(pen)
	painter.SetRenderHints(gui.QPainter__Antialiasing, true)
	painter.DrawPie3(core.NewQRect4(0, 0, int(p.Width()), int(p.Height())).Adjusted(1, 1, -1, -1), 90*16, 290*16)
}
