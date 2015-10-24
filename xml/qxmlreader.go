package xml

//#include "qxmlreader.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlReader struct {
	ptr unsafe.Pointer
}

type QXmlReaderITF interface {
	QXmlReaderPTR() *QXmlReader
}

func (p *QXmlReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlReader(ptr QXmlReaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReaderPTR().Pointer()
	}
	return nil
}

func QXmlReaderFromPointer(ptr unsafe.Pointer) *QXmlReader {
	var n = new(QXmlReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlReader) QXmlReaderPTR() *QXmlReader {
	return ptr
}

func (ptr *QXmlReader) DTDHandler() *QXmlDTDHandler {
	if ptr.Pointer() != nil {
		return QXmlDTDHandlerFromPointer(unsafe.Pointer(C.QXmlReader_DTDHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlReader) ContentHandler() *QXmlContentHandler {
	if ptr.Pointer() != nil {
		return QXmlContentHandlerFromPointer(unsafe.Pointer(C.QXmlReader_ContentHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlReader) DeclHandler() *QXmlDeclHandler {
	if ptr.Pointer() != nil {
		return QXmlDeclHandlerFromPointer(unsafe.Pointer(C.QXmlReader_DeclHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlReader) EntityResolver() *QXmlEntityResolver {
	if ptr.Pointer() != nil {
		return QXmlEntityResolverFromPointer(unsafe.Pointer(C.QXmlReader_EntityResolver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlReader) ErrorHandler() *QXmlErrorHandler {
	if ptr.Pointer() != nil {
		return QXmlErrorHandlerFromPointer(unsafe.Pointer(C.QXmlReader_ErrorHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlReader) Feature(name string, ok bool) bool {
	if ptr.Pointer() != nil {
		return C.QXmlReader_Feature(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlReader) HasFeature(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlReader_HasFeature(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlReader) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlReader_HasProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlReader) LexicalHandler() *QXmlLexicalHandler {
	if ptr.Pointer() != nil {
		return QXmlLexicalHandlerFromPointer(unsafe.Pointer(C.QXmlReader_LexicalHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlReader) Parse(input QXmlInputSourceITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlReader_Parse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

func (ptr *QXmlReader) Property(name string, ok bool) {
	if ptr.Pointer() != nil {
		C.QXmlReader_Property(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(ok)))
	}
}

func (ptr *QXmlReader) SetContentHandler(handler QXmlContentHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetContentHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlContentHandler(handler)))
	}
}

func (ptr *QXmlReader) SetDTDHandler(handler QXmlDTDHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetDTDHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlDTDHandler(handler)))
	}
}

func (ptr *QXmlReader) SetDeclHandler(handler QXmlDeclHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetDeclHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlDeclHandler(handler)))
	}
}

func (ptr *QXmlReader) SetEntityResolver(handler QXmlEntityResolverITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetEntityResolver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlEntityResolver(handler)))
	}
}

func (ptr *QXmlReader) SetErrorHandler(handler QXmlErrorHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetErrorHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlErrorHandler(handler)))
	}
}

func (ptr *QXmlReader) SetFeature(name string, value bool) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetFeature(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QXmlReader) SetLexicalHandler(handler QXmlLexicalHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetLexicalHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlLexicalHandler(handler)))
	}
}

func (ptr *QXmlReader) DestroyQXmlReader() {
	if ptr.Pointer() != nil {
		C.QXmlReader_DestroyQXmlReader(C.QtObjectPtr(ptr.Pointer()))
	}
}
