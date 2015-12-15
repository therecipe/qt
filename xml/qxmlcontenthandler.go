package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlContentHandler struct {
	ptr unsafe.Pointer
}

type QXmlContentHandler_ITF interface {
	QXmlContentHandler_PTR() *QXmlContentHandler
}

func (p *QXmlContentHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlContentHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlContentHandler(ptr QXmlContentHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlContentHandlerFromPointer(ptr unsafe.Pointer) *QXmlContentHandler {
	var n = new(QXmlContentHandler)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlContentHandler_") {
		n.SetObjectNameAbs("QXmlContentHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlContentHandler) QXmlContentHandler_PTR() *QXmlContentHandler {
	return ptr
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {
	defer qt.Recovering("QXmlContentHandler::characters")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	defer qt.Recovering("QXmlContentHandler::endDocument")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	defer qt.Recovering("QXmlContentHandler::endElement")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	defer qt.Recovering("QXmlContentHandler::endPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ErrorString() string {
	defer qt.Recovering("QXmlContentHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	defer qt.Recovering("QXmlContentHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	defer qt.Recovering("QXmlContentHandler::processingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlContentHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	defer qt.Recovering("QXmlContentHandler::skippedEntity")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	defer qt.Recovering("QXmlContentHandler::startDocument")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	defer qt.Recovering("QXmlContentHandler::startElement")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName), PointerFromQXmlAttributes(atts)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	defer qt.Recovering("QXmlContentHandler::startPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	defer qt.Recovering("QXmlContentHandler::~QXmlContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandler(ptr.Pointer())
	}
}

func (ptr *QXmlContentHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlContentHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlContentHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
