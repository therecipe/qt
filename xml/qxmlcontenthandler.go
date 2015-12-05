package xml

//#include "xml.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::characters")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::endDocument")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::endElement")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::endPrefixMapping")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::ignorableWhitespace")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::processingInstruction")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::setDocumentLocator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::skippedEntity")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::startDocument")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::startElement")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName), PointerFromQXmlAttributes(atts)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::startPrefixMapping")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlContentHandler::~QXmlContentHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandler(ptr.Pointer())
	}
}
