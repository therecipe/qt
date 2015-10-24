package xml

//#include "qxmlerrorhandler.h"
import "C"
import (
	"unsafe"
)

type QXmlErrorHandler struct {
	ptr unsafe.Pointer
}

type QXmlErrorHandlerITF interface {
	QXmlErrorHandlerPTR() *QXmlErrorHandler
}

func (p *QXmlErrorHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlErrorHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlErrorHandler(ptr QXmlErrorHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlErrorHandlerPTR().Pointer()
	}
	return nil
}

func QXmlErrorHandlerFromPointer(ptr unsafe.Pointer) *QXmlErrorHandler {
	var n = new(QXmlErrorHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlErrorHandler) QXmlErrorHandlerPTR() *QXmlErrorHandler {
	return ptr
}

func (ptr *QXmlErrorHandler) Error(exception QXmlParseExceptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Error(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlErrorHandler_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlErrorHandler) FatalError(exception QXmlParseExceptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_FatalError(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) Warning(exception QXmlParseExceptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Warning(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandler() {
	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_DestroyQXmlErrorHandler(C.QtObjectPtr(ptr.Pointer()))
	}
}
