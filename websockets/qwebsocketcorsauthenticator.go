package websockets

//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QWebSocketCorsAuthenticator struct {
	ptr unsafe.Pointer
}

type QWebSocketCorsAuthenticator_ITF interface {
	QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator
}

func (p *QWebSocketCorsAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWebSocketCorsAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWebSocketCorsAuthenticator(ptr QWebSocketCorsAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketCorsAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = new(QWebSocketCorsAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebSocketCorsAuthenticator) QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator {
	return ptr
}

func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(PointerFromQWebSocketCorsAuthenticator(other)))
}

func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(C.CString(origin)))
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(PointerFromQWebSocketCorsAuthenticator(other)))
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	defer qt.Recovering("QWebSocketCorsAuthenticator::allowed")

	if ptr.Pointer() != nil {
		return C.QWebSocketCorsAuthenticator_Allowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	defer qt.Recovering("QWebSocketCorsAuthenticator::origin")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketCorsAuthenticator_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {
	defer qt.Recovering("QWebSocketCorsAuthenticator::setAllowed")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_SetAllowed(ptr.Pointer(), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticator_ITF) {
	defer qt.Recovering("QWebSocketCorsAuthenticator::swap")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_Swap(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(other))
	}
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {
	defer qt.Recovering("QWebSocketCorsAuthenticator::~QWebSocketCorsAuthenticator")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(ptr.Pointer())
	}
}
