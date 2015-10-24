package xml

//#include "qxmllexicalhandler.h"
import "C"
import (
	"unsafe"
)

type QXmlLexicalHandler struct {
	ptr unsafe.Pointer
}

type QXmlLexicalHandlerITF interface {
	QXmlLexicalHandlerPTR() *QXmlLexicalHandler
}

func (p *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLexicalHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandlerPTR().Pointer()
	}
	return nil
}

func QXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = new(QXmlLexicalHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlLexicalHandler) QXmlLexicalHandlerPTR() *QXmlLexicalHandler {
	return ptr
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_Comment(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndCDATA(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndDTD(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartCDATA(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartDTD(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {
	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandler(C.QtObjectPtr(ptr.Pointer()))
	}
}
