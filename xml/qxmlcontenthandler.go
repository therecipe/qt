package xml

//#include "qxmlcontenthandler.h"
import "C"
import (
	"unsafe"
)

type QXmlContentHandler struct {
	ptr unsafe.Pointer
}

type QXmlContentHandlerITF interface {
	QXmlContentHandlerPTR() *QXmlContentHandler
}

func (p *QXmlContentHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlContentHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlContentHandler(ptr QXmlContentHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandlerPTR().Pointer()
	}
	return nil
}

func QXmlContentHandlerFromPointer(ptr unsafe.Pointer) *QXmlContentHandler {
	var n = new(QXmlContentHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlContentHandler) QXmlContentHandlerPTR() *QXmlContentHandler {
	return ptr
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_Characters(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndPrefixMapping(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_IgnorableWhitespace(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_ProcessingInstruction(C.QtObjectPtr(ptr.Pointer()), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocatorITF) {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlLocator(locator)))
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_SkippedEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributesITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceURI), C.CString(localName), C.CString(qName), C.QtObjectPtr(PointerFromQXmlAttributes(atts))) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartPrefixMapping(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandler(C.QtObjectPtr(ptr.Pointer()))
	}
}
