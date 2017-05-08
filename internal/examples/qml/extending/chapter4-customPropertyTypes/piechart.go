package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type PieChart struct {
	quick.QQuickPaintedItem

	_ *PieSlice `property:"pieSlice"`
}

type PieFactory struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(*PieSlice, *PieChart) `slot:"createSliceInChart"`
}

func (p *PieFactory) init() {
	p.ConnectCreateSliceInChart(p.createSliceInChart)
}

func (p *PieFactory) createSliceInChart(slice *PieSlice, chart *PieChart) {
	slice.SetParentItem(chart)
	slice.ConnectPaint(slice.paint)
}
