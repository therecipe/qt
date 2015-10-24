package xml

//#include "qxmldefaulthandler.h"
import "C"
import (
	"unsafe"
)

type QXmlDefaultHandler struct {
	QXmlDTDHandler
	QXmlErrorHandler
	QXmlLexicalHandler
	QXmlEntityResolver
	QXmlContentHandler
	QXmlDeclHandler
}

type QXmlDefaultHandlerITF interface {
	QXmlDTDHandlerITF
	QXmlErrorHandlerITF
	QXmlLexicalHandlerITF
	QXmlEntityResolverITF
	QXmlContentHandlerITF
	QXmlDeclHandlerITF
	QXmlDefaultHandlerPTR() *QXmlDefaultHandler
}

func (p *QXmlDefaultHandler) Pointer() unsafe.Pointer {
	return p.QXmlDTDHandlerPTR().Pointer()
}

func (p *QXmlDefaultHandler) SetPointer(ptr unsafe.Pointer) {
	p.QXmlDTDHandlerPTR().SetPointer(ptr)
	p.QXmlErrorHandlerPTR().SetPointer(ptr)
	p.QXmlLexicalHandlerPTR().SetPointer(ptr)
	p.QXmlEntityResolverPTR().SetPointer(ptr)
	p.QXmlContentHandlerPTR().SetPointer(ptr)
	p.QXmlDeclHandlerPTR().SetPointer(ptr)
}

func PointerFromQXmlDefaultHandler(ptr QXmlDefaultHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDefaultHandlerPTR().Pointer()
	}
	return nil
}

func QXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) *QXmlDefaultHandler {
	var n = new(QXmlDefaultHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlDefaultHandler) QXmlDefaultHandlerPTR() *QXmlDefaultHandler {
	return ptr
}

func NewQXmlDefaultHandler() *QXmlDefaultHandler {
	return QXmlDefaultHandlerFromPointer(unsafe.Pointer(C.QXmlDefaultHandler_NewQXmlDefaultHandler()))
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_DestroyQXmlDefaultHandler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlDefaultHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_AttributeDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Characters(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Characters(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Comment(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Comment(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndCDATA() bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndCDATA(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDTD() bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDTD(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndPrefixMapping(prefix string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndPrefixMapping(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Error(exception QXmlParseExceptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Error(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDefaultHandler_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlDefaultHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ExternalEntityDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) FatalError(exception QXmlParseExceptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_FatalError(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespace(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_IgnorableWhitespace(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) InternalEntityDecl(name string, value string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_InternalEntityDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) NotationDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_NotationDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ProcessingInstruction(target string, data string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ProcessingInstruction(C.QtObjectPtr(ptr.Pointer()), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) SetDocumentLocator(locator QXmlLocatorITF) {
	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlLocator(locator)))
	}
}

func (ptr *QXmlDefaultHandler) SkippedEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_SkippedEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartCDATA() bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartCDATA(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDTD(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDTD(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributesITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceURI), C.CString(localName), C.CString(qName), C.QtObjectPtr(PointerFromQXmlAttributes(atts))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartEntity(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartPrefixMapping(prefix string, uri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartPrefixMapping(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_UnparsedEntityDecl(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Warning(exception QXmlParseExceptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Warning(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlParseException(exception))) != 0
	}
	return false
}
