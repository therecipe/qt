package xml

//#include "qxmldtdhandler.h"
import "C"
import (
	"unsafe"
)

type QXmlDTDHandler struct {
	ptr unsafe.Pointer
}

type QXmlDTDHandlerITF interface {
	QXmlDTDHandlerPTR() *QXmlDTDHandler
}

func (p *QXmlDTDHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlDTDHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlDTDHandler(ptr QXmlDTDHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDTDHandlerPTR().Pointer()
	}
	return nil
}

func QXmlDTDHandlerFromPointer(ptr unsafe.Pointer) *QXmlDTDHandler {
	var n = new(QXmlDTDHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlDTDHandler) QXmlDTDHandlerPTR() *QXmlDTDHandler {
	return ptr
}

func (ptr *QXmlDTDHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_NotationDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_UnparsedEntityDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_DestroyQXmlDTDHandler(C.QtObjectPtr(ptr.Pointer()))
	}
}
