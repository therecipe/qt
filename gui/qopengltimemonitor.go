package gui

//#include "qopengltimemonitor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLTimeMonitor struct {
	core.QObject
}

type QOpenGLTimeMonitor_ITF interface {
	core.QObject_ITF
	QOpenGLTimeMonitor_PTR() *QOpenGLTimeMonitor
}

func PointerFromQOpenGLTimeMonitor(ptr QOpenGLTimeMonitor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLTimeMonitor_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLTimeMonitorFromPointer(ptr unsafe.Pointer) *QOpenGLTimeMonitor {
	var n = new(QOpenGLTimeMonitor)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QOpenGLTimeMonitor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLTimeMonitor) QOpenGLTimeMonitor_PTR() *QOpenGLTimeMonitor {
	return ptr
}
