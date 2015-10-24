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

type QRasterWindowITF interface {
	QPaintDeviceWindowITF
	QRasterWindowPTR() *QRasterWindow
}

func PointerFromQRasterWindow(ptr QRasterWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRasterWindowPTR().Pointer()
	}
	return nil
}

func QRasterWindowFromPointer(ptr unsafe.Pointer) *QRasterWindow {
	var n = new(QRasterWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRasterWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRasterWindow) QRasterWindowPTR() *QRasterWindow {
	return ptr
}

func NewQRasterWindow(parent QWindowITF) *QRasterWindow {
	return QRasterWindowFromPointer(unsafe.Pointer(C.QRasterWindow_NewQRasterWindow(C.QtObjectPtr(PointerFromQWindow(parent)))))
}
