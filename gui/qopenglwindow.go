package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOpenGLWindow struct {
	QPaintDeviceWindow
}

type QOpenGLWindow_ITF interface {
	QPaintDeviceWindow_ITF
	QOpenGLWindow_PTR() *QOpenGLWindow
}

func PointerFromQOpenGLWindow(ptr QOpenGLWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLWindow_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLWindowFromPointer(ptr unsafe.Pointer) *QOpenGLWindow {
	var n = new(QOpenGLWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOpenGLWindow_") {
		n.SetObjectName("QOpenGLWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QOpenGLWindow) QOpenGLWindow_PTR() *QOpenGLWindow {
	return ptr
}

//QOpenGLWindow::UpdateBehavior
type QOpenGLWindow__UpdateBehavior int64

const (
	QOpenGLWindow__NoPartialUpdate    = QOpenGLWindow__UpdateBehavior(0)
	QOpenGLWindow__PartialUpdateBlit  = QOpenGLWindow__UpdateBehavior(1)
	QOpenGLWindow__PartialUpdateBlend = QOpenGLWindow__UpdateBehavior(2)
)
