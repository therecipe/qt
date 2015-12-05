package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDesktopServices struct {
	ptr unsafe.Pointer
}

type QDesktopServices_ITF interface {
	QDesktopServices_PTR() *QDesktopServices
}

func (p *QDesktopServices) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDesktopServices) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDesktopServices(ptr QDesktopServices_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesktopServices_PTR().Pointer()
	}
	return nil
}

func NewQDesktopServicesFromPointer(ptr unsafe.Pointer) *QDesktopServices {
	var n = new(QDesktopServices)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDesktopServices) QDesktopServices_PTR() *QDesktopServices {
	return ptr
}

func QDesktopServices_OpenUrl(url core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopServices::openUrl")
		}
	}()

	return C.QDesktopServices_QDesktopServices_OpenUrl(core.PointerFromQUrl(url)) != 0
}

func QDesktopServices_SetUrlHandler(scheme string, receiver core.QObject_ITF, method string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopServices::setUrlHandler")
		}
	}()

	C.QDesktopServices_QDesktopServices_SetUrlHandler(C.CString(scheme), core.PointerFromQObject(receiver), C.CString(method))
}

func QDesktopServices_UnsetUrlHandler(scheme string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopServices::unsetUrlHandler")
		}
	}()

	C.QDesktopServices_QDesktopServices_UnsetUrlHandler(C.CString(scheme))
}
