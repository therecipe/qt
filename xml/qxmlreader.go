package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlReader struct {
	ptr unsafe.Pointer
}

type QXmlReader_ITF interface {
	QXmlReader_PTR() *QXmlReader
}

func (p *QXmlReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlReader(ptr QXmlReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlReaderFromPointer(ptr unsafe.Pointer) *QXmlReader {
	var n = new(QXmlReader)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlReader_") {
		n.SetObjectNameAbs("QXmlReader_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlReader) QXmlReader_PTR() *QXmlReader {
	return ptr
}

func (ptr *QXmlReader) DTDHandler() *QXmlDTDHandler {
	defer qt.Recovering("QXmlReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) ContentHandler() *QXmlContentHandler {
	defer qt.Recovering("QXmlReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) DeclHandler() *QXmlDeclHandler {
	defer qt.Recovering("QXmlReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) EntityResolver() *QXmlEntityResolver {
	defer qt.Recovering("QXmlReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) ErrorHandler() *QXmlErrorHandler {
	defer qt.Recovering("QXmlReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) Feature(name string, ok bool) bool {
	defer qt.Recovering("QXmlReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlReader) HasFeature(name string) bool {
	defer qt.Recovering("QXmlReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlReader) HasProperty(name string) bool {
	defer qt.Recovering("QXmlReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlReader) LexicalHandler() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) Parse(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlReader) Property(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

func (ptr *QXmlReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlReader) SetFeature(name string, value bool) {
	defer qt.Recovering("QXmlReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QXmlReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlReader) DestroyQXmlReader() {
	defer qt.Recovering("QXmlReader::~QXmlReader")

	if ptr.Pointer() != nil {
		C.QXmlReader_DestroyQXmlReader(ptr.Pointer())
	}
}

func (ptr *QXmlReader) ObjectNameAbs() string {
	defer qt.Recovering("QXmlReader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlReader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlReader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlReader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
