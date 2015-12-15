package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QXmlDTDHandler_") {
		n.SetObjectNameAbs("QXmlDTDHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlDTDHandler) QXmlDTDHandler_PTR() *QXmlDTDHandler {
	return ptr
}

func (ptr *QXmlDTDHandler) ErrorString() string {
	defer qt.Recovering("QXmlDTDHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDTDHandler::notationDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_NotationDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	defer qt.Recovering("QXmlDTDHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_UnparsedEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {
	defer qt.Recovering("QXmlDTDHandler::~QXmlDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_DestroyQXmlDTDHandler(ptr.Pointer())
	}
}

func (ptr *QXmlDTDHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDTDHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDTDHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDTDHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
