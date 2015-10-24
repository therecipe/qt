package websockets

//#include "qwebsocketcorsauthenticator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QWebSocketCorsAuthenticator struct {
	ptr unsafe.Pointer
}

type QWebSocketCorsAuthenticatorITF interface {
	QWebSocketCorsAuthenticatorPTR() *QWebSocketCorsAuthenticator
}

func (p *QWebSocketCorsAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWebSocketCorsAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWebSocketCorsAuthenticator(ptr QWebSocketCorsAuthenticatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketCorsAuthenticatorPTR().Pointer()
	}
	return nil
}

func QWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = new(QWebSocketCorsAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebSocketCorsAuthenticator) QWebSocketCorsAuthenticatorPTR() *QWebSocketCorsAuthenticator {
	return ptr
}

func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticatorITF) *QWebSocketCorsAuthenticator {
	return QWebSocketCorsAuthenticatorFromPointer(unsafe.Pointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(C.QtObjectPtr(PointerFromQWebSocketCorsAuthenticator(other)))))
}

func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	return QWebSocketCorsAuthenticatorFromPointer(unsafe.Pointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(C.CString(origin))))
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticatorITF) *QWebSocketCorsAuthenticator {
	return QWebSocketCorsAuthenticatorFromPointer(unsafe.Pointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(C.QtObjectPtr(PointerFromQWebSocketCorsAuthenticator(other)))))
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketCorsAuthenticator_Allowed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketCorsAuthenticator_Origin(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_SetAllowed(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticatorITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWebSocketCorsAuthenticator(other)))
	}
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(C.QtObjectPtr(ptr.Pointer()))
	}
}
