package xml

//#include "qxmlsimplereader.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlSimpleReader struct {
	QXmlReader
}

type QXmlSimpleReaderITF interface {
	QXmlReaderITF
	QXmlSimpleReaderPTR() *QXmlSimpleReader
}

func PointerFromQXmlSimpleReader(ptr QXmlSimpleReaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSimpleReaderPTR().Pointer()
	}
	return nil
}

func QXmlSimpleReaderFromPointer(ptr unsafe.Pointer) *QXmlSimpleReader {
	var n = new(QXmlSimpleReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSimpleReader) QXmlSimpleReaderPTR() *QXmlSimpleReader {
	return ptr
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {
	if ptr.Pointer() != nil {
		return QXmlDTDHandlerFromPointer(unsafe.Pointer(C.QXmlSimpleReader_DTDHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQXmlSimpleReader() *QXmlSimpleReader {
	return QXmlSimpleReaderFromPointer(unsafe.Pointer(C.QXmlSimpleReader_NewQXmlSimpleReader()))
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {
	if ptr.Pointer() != nil {
		return QXmlContentHandlerFromPointer(unsafe.Pointer(C.QXmlSimpleReader_ContentHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {
	if ptr.Pointer() != nil {
		return QXmlDeclHandlerFromPointer(unsafe.Pointer(C.QXmlSimpleReader_DeclHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {
	if ptr.Pointer() != nil {
		return QXmlEntityResolverFromPointer(unsafe.Pointer(C.QXmlSimpleReader_EntityResolver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {
	if ptr.Pointer() != nil {
		return QXmlErrorHandlerFromPointer(unsafe.Pointer(C.QXmlSimpleReader_ErrorHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Feature(name string, ok bool) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Feature(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasFeature(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {
	if ptr.Pointer() != nil {
		return QXmlLexicalHandlerFromPointer(unsafe.Pointer(C.QXmlSimpleReader_LexicalHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Parse(input QXmlInputSourceITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse2(input QXmlInputSourceITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse3(input QXmlInputSourceITF, incremental bool) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlInputSource(input)), C.int(qt.GoBoolToInt(incremental))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {
	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_ParseContinue(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Property(name string, ok bool) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_Property(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(ok)))
	}
}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlContentHandler(handler)))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlDTDHandler(handler)))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlDeclHandler(handler)))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolverITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlEntityResolver(handler)))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlErrorHandler(handler)))
	}
}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeature(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlLexicalHandler(handler)))
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_DestroyQXmlSimpleReader(C.QtObjectPtr(ptr.Pointer()))
	}
}
