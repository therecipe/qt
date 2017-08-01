package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

const (
	GL_COLOR_BUFFER_BIT   uint = 0x00004000
	GL_DEPTH_BUFFER_BIT   uint = 0x00000100
	GL_STENCIL_BUFFER_BIT uint = 0x00000400
	GL_FLOAT              uint = 0x1406
	GL_FALSE              bool = false //0
	GL_TRIANGLES          uint = 0x0004
)

type OpenGLWindow struct {
	gui.QWindow
	gui.QOpenGLFunctions

	_ func() `constructor:"init"`

	m_animating bool
	m_context   *gui.QOpenGLContext
	m_device    *gui.QOpenGLPaintDevice

	initializeLazy func()
	renderLazy     func(*gui.QPainter)
}

func (w *OpenGLWindow) init() {
	w.QOpenGLFunctions = *gui.NewQOpenGLFunctions()
	w.SetSurfaceType(gui.QSurface__OpenGLSurface)
}

func (w *OpenGLWindow) render() {
	if w.m_device == nil {
		w.m_device = gui.NewQOpenGLPaintDevice()
	}

	w.GlClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT | GL_STENCIL_BUFFER_BIT)

	w.m_device.SetSize(w.Size())

	painter := gui.NewQPainter2(w.m_device)
	w.renderLazy(painter)
	painter.DestroyQPainter()
}

func (w *OpenGLWindow) renderLater() {
	w.RequestUpdate()
}

func (w *OpenGLWindow) event(event *core.QEvent) bool {
	switch event.Type() {
	case core.QEvent__UpdateRequest:
		w.renderNow()
		return true
	default:
		return w.EventDefault(event)
	}
}

func (w *OpenGLWindow) exposeEvent(*gui.QExposeEvent) {
	if w.IsExposed() {
		w.renderNow()
	}
}

func (w *OpenGLWindow) renderNow() {
	if !w.IsExposed() {
		return
	}

	var needsInitialize bool

	if w.m_context == nil {
		w.m_context = gui.NewQOpenGLContext(w)
		w.m_context.SetFormat(w.RequestedFormat())
		w.m_context.Create()

		needsInitialize = true
	}

	w.m_context.MakeCurrent(w)

	if needsInitialize {
		w.InitializeOpenGLFunctions()
		w.initializeLazy()
	}

	w.render()

	w.m_context.SwapBuffers(w)

	if w.m_animating {
		w.renderLater()
	}
}

func (w *OpenGLWindow) setAnimating(animating bool) {
	w.m_animating = animating

	if animating {
		w.renderLater()
	}
}
