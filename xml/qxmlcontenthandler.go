package xml

//#include "qxmlcontenthandler.h"
import "C"
import (
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
	return n
}

func (ptr *QXmlContentHandler) QXmlContentHandler_PTR() *QXmlContentHandler {
	return ptr
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName), PointerFromQXmlAttributes(atts)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandler(ptr.Pointer())
	}
}
