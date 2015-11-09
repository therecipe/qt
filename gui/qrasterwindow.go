package gui

//#include "qrasterwindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRasterWindow struct {
	QPaintDeviceWindow
}

type QRasterWindow_ITF interface {
	QPaintDeviceWindow_ITF
	QRasterWindow_PTR() *QRasterWindow
}

func PointerFromQRasterWindow(ptr QRasterWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRasterWindow_PTR().Pointer()
	}
	return nil
}

func NewQRasterWindowFromPointer(ptr unsafe.Pointer) *QRasterWindow {
	var n = new(QRasterWindow)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QRasterWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRasterWindow) QRasterWindow_PTR() *QRasterWindow {
	return ptr
}

func NewQRasterWindow(parent QWindow_ITF) *QRasterWindow {
	return NewQRasterWindowFromPointer(C.QRasterWindow_NewQRasterWindow(PointerFromQWindow(parent)))
}
