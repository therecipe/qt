package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QXmlSimpleReader struct {
	QXmlReader
}

type QXmlSimpleReader_ITF interface {
	QXmlReader_ITF
	QXmlSimpleReader_PTR() *QXmlSimpleReader
}

func PointerFromQXmlSimpleReader(ptr QXmlSimpleReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSimpleReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) *QXmlSimpleReader {
	var n = new(QXmlSimpleReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSimpleReader) QXmlSimpleReader_PTR() *QXmlSimpleReader {
	return ptr
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::DTDHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func NewQXmlSimpleReader() *QXmlSimpleReader {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::QXmlSimpleReader")
		}
	}()

	return NewQXmlSimpleReaderFromPointer(C.QXmlSimpleReader_NewQXmlSimpleReader())
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::contentHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::declHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::entityResolver")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::errorHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Feature(name string, ok bool) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::feature")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::hasFeature")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::hasProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::lexicalHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Parse(input QXmlInputSource_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::parse")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse2(input QXmlInputSource_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::parse")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse2(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse3(input QXmlInputSource_ITF, incremental bool) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::parse")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse3(ptr.Pointer(), PointerFromQXmlInputSource(input), C.int(qt.GoBoolToInt(incremental))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::parseContinue")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_ParseContinue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Property(name string, ok bool) unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::property")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlSimpleReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setContentHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setDTDHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setDeclHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setEntityResolver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setErrorHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setFeature")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::setLexicalHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSimpleReader::~QXmlSimpleReader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_DestroyQXmlSimpleReader(ptr.Pointer())
	}
}
