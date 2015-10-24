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

type QPaintDeviceWindowITF interface {
	QWindowITF
	QPaintDeviceITF
	QPaintDeviceWindowPTR() *QPaintDeviceWindow
}

func (p *QPaintDeviceWindow) Pointer() unsafe.Pointer {
	return p.QWindowPTR().Pointer()
}

func (p *QPaintDeviceWindow) SetPointer(ptr unsafe.Pointer) {
	p.QWindowPTR().SetPointer(ptr)
	p.QPaintDevicePTR().SetPointer(ptr)
}

func PointerFromQPaintDeviceWindow(ptr QPaintDeviceWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDeviceWindowPTR().Pointer()
	}
	return nil
}

func QPaintDeviceWindowFromPointer(ptr unsafe.Pointer) *QPaintDeviceWindow {
	var n = new(QPaintDeviceWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPaintDeviceWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPaintDeviceWindow) QPaintDeviceWindowPTR() *QPaintDeviceWindow {
	return ptr
}

func (ptr *QPaintDeviceWindow) Update3() {
	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update3(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPaintDeviceWindow) Update(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QPaintDeviceWindow) Update2(region QRegionITF) {
	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)))
	}
}
