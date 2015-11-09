package xml

//#include "qxmldtdhandler.h"
import "C"
import (
	"unsafe"
)

type QXmlDTDHandler struct {
	ptr unsafe.Pointer
}

type QXmlDTDHandler_ITF interface {
	QXmlDTDHandler_PTR() *QXmlDTDHandler
}

func (p *QXmlDTDHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlDTDHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlDTDHandler(ptr QXmlDTDHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDTDHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) *QXmlDTDHandler {
	var n = new(QXmlDTDHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlDTDHandler) QXmlDTDHandler_PTR() *QXmlDTDHandler {
	return ptr
}

func (ptr *QXmlDTDHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_NotationDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_UnparsedEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_DestroyQXmlDTDHandler(ptr.Pointer())
	}
}
