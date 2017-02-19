//source: http://doc.qt.io/qt-5/qtgui-analogclock-example.html

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type AnalogClockWindow struct {
	RasterWindow

	_ func() `constructor:"init"`

	m_timerId int
}

func (w *AnalogClockWindow) init() {
	w.RasterWindow.init()

	w.SetTitle("Analog Clock")
	w.Resize2(200, 200)

	w.m_timerId = w.StartTimer(1000, 0)

	w.ConnectTimerEvent(w.timerEvent)

	w.render = w.renderFunc
}

func (w *AnalogClockWindow) timerEvent(event *core.QTimerEvent) {
	if event.TimerId() == w.m_timerId {
		w.renderLater()
	}
}

func (w *AnalogClockWindow) renderFunc(p *gui.QPainter) {

	var hourHand = gui.NewQPolygon3([]*core.QPoint{
		core.NewQPoint2(7, 8),
		core.NewQPoint2(-7, 8),
		core.NewQPoint2(0, -40),
	})

	var minuteHand = gui.NewQPolygon4([]*core.QPoint{
		core.NewQPoint2(7, 8),
		core.NewQPoint2(-7, 8),
		core.NewQPoint2(0, -70),
	})

	var hourColor = gui.NewQColor3(127, 0, 127, 255)
	var minuteColor = gui.NewQColor3(0, 127, 127, 191)

	p.SetRenderHint(gui.QPainter__Antialiasing, true)

	p.Translate3(float64(w.Width())/2, float64(w.Height())/2)

	var side = func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}(w.Width(), w.Height())
	p.Scale(float64(side)/200.0, float64(side)/200.0)

	p.SetPen3(core.Qt__NoPen)
	p.SetBrush(gui.NewQBrush3(hourColor, 1))

	var time = core.QTime_CurrentTime()

	p.Save()
	p.Rotate(float64(30.0 * (time.Hour() + time.Minute()/60.0)))
	p.DrawConvexPolygon4(hourHand)
	p.Restore()

	p.SetPen2(hourColor)

	for i := 0; i < 12; i++ {
		p.DrawLine3(88, 0, 96, 0)
		p.Rotate(30.0)
	}

	p.SetPen3(core.Qt__NoPen)
	p.SetBrush(gui.NewQBrush3(minuteColor, 1))

	p.Save()
	p.Rotate(float64(6.0 * (time.Minute() + time.Second()/60.0)))
	p.DrawConvexPolygon4(minuteHand)
	p.Restore()

	p.SetPen2(minuteColor)

	for j := 0; j < 60; j++ {
		if (j % 5) != 0 {
			p.DrawLine3(92, 0, 96, 0)
		}
		p.Rotate(6.0)
	}
}

func main() {
	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var clock = NewAnalogClockWindow(nil)
	clock.Show()

	app.Exec()
}
