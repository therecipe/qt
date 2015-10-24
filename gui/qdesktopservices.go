package gui

//#include "qdesktopservices.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDesktopServices struct {
	ptr unsafe.Pointer
}

type QDesktopServicesITF interface {
	QDesktopServicesPTR() *QDesktopServices
}

func (p *QDesktopServices) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDesktopServices) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDesktopServices(ptr QDesktopServicesITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesktopServicesPTR().Pointer()
	}
	return nil
}

func QDesktopServicesFromPointer(ptr unsafe.Pointer) *QDesktopServices {
	var n = new(QDesktopServices)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDesktopServices) QDesktopServicesPTR() *QDesktopServices {
	return ptr
}

func QDesktopServices_OpenUrl(url string) bool {
	return C.QDesktopServices_QDesktopServices_OpenUrl(C.CString(url)) != 0
}

func QDesktopServices_SetUrlHandler(scheme string, receiver core.QObjectITF, method string) {
	C.QDesktopServices_QDesktopServices_SetUrlHandler(C.CString(scheme), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(method))
}

func QDesktopServices_UnsetUrlHandler(scheme string) {
	C.QDesktopServices_QDesktopServices_UnsetUrlHandler(C.CString(scheme))
}
