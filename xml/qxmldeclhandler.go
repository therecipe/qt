package xml

//#include "qxmldeclhandler.h"
import "C"
import (
	"unsafe"
)

type QXmlDeclHandler struct {
	ptr unsafe.Pointer
}

type QXmlDeclHandlerITF interface {
	QXmlDeclHandlerPTR() *QXmlDeclHandler
}

func (p *QXmlDeclHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlDeclHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlDeclHandler(ptr QXmlDeclHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDeclHandlerPTR().Pointer()
	}
	return nil
}

func QXmlDeclHandlerFromPointer(ptr unsafe.Pointer) *QXmlDeclHandler {
	var n = new(QXmlDeclHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlDeclHandler) QXmlDeclHandlerPTR() *QXmlDeclHandler {
	return ptr
}

func (ptr *QXmlDeclHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_AttributeDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlDeclHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_ExternalEntityDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) InternalEntityDecl(name string, value string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_InternalEntityDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_DestroyQXmlDeclHandler(C.QtObjectPtr(ptr.Pointer()))
	}
}
