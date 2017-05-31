package main

import (
	"github.com/therecipe/qt/quick"
)

func init() {
	PieChart_QmlRegisterType2("Charts", 1, 0, "PieChart")
}

type PieChart struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`

	_ *PieSlice `property:"pieSlice"`
}

func (p *PieChart) init() {
	p.ConnectPieSliceChanged(p.setParentForSlice)
}

func (p *PieChart) setParentForSlice(slice *PieSlice) {
	slice.SetParentItem(p)
}
