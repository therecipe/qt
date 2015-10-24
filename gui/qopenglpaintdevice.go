package gui

//#include "qopenglpaintdevice.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLPaintDevice struct {
	QPaintDevice
}

type QOpenGLPaintDeviceITF interface {
	QPaintDeviceITF
	QOpenGLPaintDevicePTR() *QOpenGLPaintDevice
}

func PointerFromQOpenGLPaintDevice(ptr QOpenGLPaintDeviceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLPaintDevicePTR().Pointer()
	}
	return nil
}

func QOpenGLPaintDeviceFromPointer(ptr unsafe.Pointer) *QOpenGLPaintDevice {
	var n = new(QOpenGLPaintDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLPaintDevice) QOpenGLPaintDevicePTR() *QOpenGLPaintDevice {
	return ptr
}

func NewQOpenGLPaintDevice() *QOpenGLPaintDevice {
	return QOpenGLPaintDeviceFromPointer(unsafe.Pointer(C.QOpenGLPaintDevice_NewQOpenGLPaintDevice()))
}

func NewQOpenGLPaintDevice2(size core.QSizeITF) *QOpenGLPaintDevice {
	return QOpenGLPaintDeviceFromPointer(unsafe.Pointer(C.QOpenGLPaintDevice_NewQOpenGLPaintDevice2(C.QtObjectPtr(core.PointerFromQSize(size)))))
}

func NewQOpenGLPaintDevice3(width int, height int) *QOpenGLPaintDevice {
	return QOpenGLPaintDeviceFromPointer(unsafe.Pointer(C.QOpenGLPaintDevice_NewQOpenGLPaintDevice3(C.int(width), C.int(height))))
}

func (ptr *QOpenGLPaintDevice) Context() *QOpenGLContext {
	if ptr.Pointer() != nil {
		return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLPaintDevice_Context(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLPaintDevice) EnsureActiveTarget() {
	if ptr.Pointer() != nil {
		C.QOpenGLPaintDevice_EnsureActiveTarget(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return QPaintEngineFromPointer(unsafe.Pointer(C.QOpenGLPaintDevice_PaintEngine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLPaintDevice) PaintFlipped() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLPaintDevice_PaintFlipped(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLPaintDevice) SetPaintFlipped(flipped bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLPaintDevice_SetPaintFlipped(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(flipped)))
	}
}

func (ptr *QOpenGLPaintDevice) SetSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLPaintDevice_SetSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QOpenGLPaintDevice) DestroyQOpenGLPaintDevice() {
	if ptr.Pointer() != nil {
		C.QOpenGLPaintDevice_DestroyQOpenGLPaintDevice(C.QtObjectPtr(ptr.Pointer()))
	}
}
