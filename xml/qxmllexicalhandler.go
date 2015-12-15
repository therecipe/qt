package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlLexicalHandler struct {
	ptr unsafe.Pointer
}

type QXmlLexicalHandler_ITF interface {
	QXmlLexicalHandler_PTR() *QXmlLexicalHandler
}

func (p *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLexicalHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = new(QXmlLexicalHandler)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlLexicalHandler_") {
		n.SetObjectNameAbs("QXmlLexicalHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlLexicalHandler) QXmlLexicalHandler_PTR() *QXmlLexicalHandler {
	return ptr
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {
	defer qt.Recovering("QXmlLexicalHandler::comment")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_Comment(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {
	defer qt.Recovering("QXmlLexicalHandler::endCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {
	defer qt.Recovering("QXmlLexicalHandler::endDTD")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndDTD(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {
	defer qt.Recovering("QXmlLexicalHandler::endEntity")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) ErrorString() string {
	defer qt.Recovering("QXmlLexicalHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {
	defer qt.Recovering("QXmlLexicalHandler::startCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlLexicalHandler::startDTD")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartDTD(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {
	defer qt.Recovering("QXmlLexicalHandler::startEntity")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {
	defer qt.Recovering("QXmlLexicalHandler::~QXmlLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandler(ptr.Pointer())
	}
}

func (ptr *QXmlLexicalHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlLexicalHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlLexicalHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
