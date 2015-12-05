package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLPaintDevice struct {
	QPaintDevice
}

type QOpenGLPaintDevice_ITF interface {
	QPaintDevice_ITF
	QOpenGLPaintDevice_PTR() *QOpenGLPaintDevice
}

func PointerFromQOpenGLPaintDevice(ptr QOpenGLPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLPaintDeviceFromPointer(ptr unsafe.Pointer) *QOpenGLPaintDevice {
	var n = new(QOpenGLPaintDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLPaintDevice) QOpenGLPaintDevice_PTR() *QOpenGLPaintDevice {
	return ptr
}
