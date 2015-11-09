package xml

//#include "qxmlerrorhandler.h"
import "C"
import (
	"unsafe"
)

type QXmlErrorHandler struct {
	ptr unsafe.Pointer
}

type QXmlErrorHandler_ITF interface {
	QXmlErrorHandler_PTR() *QXmlErrorHandler
}

func (p *QXmlErrorHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlErrorHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlErrorHandler(ptr QXmlErrorHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlErrorHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) *QXmlErrorHandler {
	var n = new(QXmlErrorHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlErrorHandler) QXmlErrorHandler_PTR() *QXmlErrorHandler {
	return ptr
}

func (ptr *QXmlErrorHandler) Error(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlErrorHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlErrorHandler) FatalError(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) Warning(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandler() {
	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_DestroyQXmlErrorHandler(ptr.Pointer())
	}
}
