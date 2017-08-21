package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type Helper struct {
	core.QObject

	_ func() `constructor:"init"`

	background  *gui.QBrush
	circleBrush *gui.QBrush
	textFont    *gui.QFont
	circlePen   *gui.QPen
	textPen     *gui.QPen
}

func (h *Helper) init() {
	gradient := gui.NewQLinearGradient2(core.NewQPointF3(50, -20), core.NewQPointF3(80, 20))
	gradient.SetColorAt(0, gui.NewQColor2(core.Qt__white))
	gradient.SetColorAt(1, gui.NewQColor3(0xa6, 0xce, 0x39, 255))

	h.background = gui.NewQBrush3(gui.NewQColor3(64, 32, 64, 255), 1)
	h.circleBrush = gui.NewQBrush10(gradient)
	h.circlePen = gui.NewQPen3(gui.NewQColor2(core.Qt__black))
	h.circlePen.SetWidth(1)
	h.textPen = gui.NewQPen3(gui.NewQColor2(core.Qt__white))
	h.textFont = gui.NewQFont()
	h.textFont.SetPixelSize(50)
}

func (h *Helper) Paint(painter *gui.QPainter, event *gui.QPaintEvent, elapsed int) {
	painter.FillRect3(event.Rect(), h.background)
	painter.Translate3(100, 100)

	painter.Save()
	painter.SetBrush(h.circleBrush)
	painter.SetPen(h.circlePen)
	painter.Rotate(float64(elapsed) * float64(0.03))

	r := float64(elapsed) / 1000
	n := float64(30)
	for i := 0; i < int(n); {
		i++
		painter.Rotate(30)
		factor := (float64(i) + r) / n
		radius := 0 + 120*factor
		circleRadius := 1 + factor*20
		painter.DrawEllipse(core.NewQRectF4(radius, -circleRadius, circleRadius*2, circleRadius*2))
	}
	painter.Restore()

	painter.SetPen(h.textPen)
	painter.SetFont(h.textFont)
	painter.DrawText6(core.NewQRect4(-50, -50, 100, 100), int(core.Qt__AlignCenter), "Qt", core.NewQRect())
}
