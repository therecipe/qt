package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type RasterWindow struct {
	gui.QWindow

	_ func() `constructor:"init"`

	_ func() `slot:"renderLater"`
	_ func() `slot:"renderNow"`

	render func(painter *gui.QPainter)

	m_backingStore *gui.QBackingStore
}

func (w *RasterWindow) init() {
	w.Create()
	w.m_backingStore = gui.NewQBackingStore(w)

	w.SetGeometry(100, 100, 300, 200)

	w.ConnectEvent(w.event)
	w.ConnectRenderLater(w.renderLater)
	w.ConnectResizeEvent(w.resizeEvent)
	w.ConnectExposeEvent(w.exposeEvent)
	w.ConnectRenderNow(w.renderNow)

	w.render = w.renderFunc
}

func (w *RasterWindow) event(ev *core.QEvent) bool {
	if ev.Type() == core.QEvent__UpdateRequest {
		w.RenderNow()
		return true
	}
	return w.EventDefault(ev)
}

func (w *RasterWindow) renderLater() {
	w.RequestUpdate()
}

func (w *RasterWindow) resizeEvent(ev *gui.QResizeEvent) {
	w.m_backingStore.Resize(ev.Size())
	if w.IsExposed() {
		w.RenderNow()
	}
}

func (w *RasterWindow) exposeEvent(ev *gui.QExposeEvent) {
	if w.IsExposed() {
		w.RenderNow()
	}
}

func (w *RasterWindow) renderNow() {
	if !w.IsExposed() {
		return
	}

	var rect = core.NewQRect4(0, 0, w.Width(), w.Height())
	w.m_backingStore.BeginPaint(gui.NewQRegion3(rect, 0))

	var device = w.m_backingStore.PaintDevice()
	var painter = gui.NewQPainter2(device)
	defer painter.DestroyQPainter()

	painter.FillRect7(0, 0, w.Width(), w.Height(), core.Qt__white)
	w.render(painter)

	w.m_backingStore.EndPaint()
	w.m_backingStore.Flush(gui.NewQRegion3(rect, 0), w, core.NewQPoint())
}

func (w *RasterWindow) renderFunc(painter *gui.QPainter) {
	painter.DrawText5(core.NewQRectF4(0, 0, float64(w.Width()), float64(w.Height())), int(core.Qt__AlignCenter), "QWindow", core.NewQRectF())
}
