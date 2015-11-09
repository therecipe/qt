package gui

//#include "qpaintdevicewindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPaintDeviceWindow struct {
	QWindow
	QPaintDevice
}

type QPaintDeviceWindow_ITF interface {
	QWindow_ITF
	QPaintDevice_ITF
	QPaintDeviceWindow_PTR() *QPaintDeviceWindow
}

func (p *QPaintDeviceWindow) Pointer() unsafe.Pointer {
	return p.QWindow_PTR().Pointer()
}

func (p *QPaintDeviceWindow) SetPointer(ptr unsafe.Pointer) {
	p.QWindow_PTR().SetPointer(ptr)
	p.QPaintDevice_PTR().SetPointer(ptr)
}

func PointerFromQPaintDeviceWindow(ptr QPaintDeviceWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDeviceWindow_PTR().Pointer()
	}
	return nil
}

func NewQPaintDeviceWindowFromPointer(ptr unsafe.Pointer) *QPaintDeviceWindow {
	var n = new(QPaintDeviceWindow)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPaintDeviceWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPaintDeviceWindow) QPaintDeviceWindow_PTR() *QPaintDeviceWindow {
	return ptr
}

func (ptr *QPaintDeviceWindow) Update3() {
	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update3(ptr.Pointer())
	}
}

func (ptr *QPaintDeviceWindow) Update(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QPaintDeviceWindow) Update2(region QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update2(ptr.Pointer(), PointerFromQRegion(region))
	}
}
