package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QXmlSimpleReader_") {
		n.SetObjectNameAbs("QXmlSimpleReader_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlSimpleReader) QXmlSimpleReader_PTR() *QXmlSimpleReader {
	return ptr
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {
	defer qt.Recovering("QXmlSimpleReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func NewQXmlSimpleReader() *QXmlSimpleReader {
	defer qt.Recovering("QXmlSimpleReader::QXmlSimpleReader")

	return NewQXmlSimpleReaderFromPointer(C.QXmlSimpleReader_NewQXmlSimpleReader())
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {
	defer qt.Recovering("QXmlSimpleReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {
	defer qt.Recovering("QXmlSimpleReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {
	defer qt.Recovering("QXmlSimpleReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {
	defer qt.Recovering("QXmlSimpleReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Feature(name string, ok bool) bool {
	defer qt.Recovering("QXmlSimpleReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlSimpleReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Parse(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {
	defer qt.Recovering("QXmlSimpleReader::parseContinue")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_ParseContinue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Property(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlSimpleReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlSimpleReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setContentHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetContentHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setContentHandler")
	}
}

//export callbackQXmlSimpleReaderSetContentHandler
func callbackQXmlSimpleReaderSetContentHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setContentHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContentHandler"); signal != nil {
		signal.(func(*QXmlContentHandler))(NewQXmlContentHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetContentHandlerDefault(NewQXmlContentHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetContentHandlerDefault(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandlerDefault(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDTDHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDTDHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDTDHandler")
	}
}

//export callbackQXmlSimpleReaderSetDTDHandler
func callbackQXmlSimpleReaderSetDTDHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setDTDHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDTDHandler"); signal != nil {
		signal.(func(*QXmlDTDHandler))(NewQXmlDTDHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetDTDHandlerDefault(NewQXmlDTDHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandlerDefault(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandlerDefault(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDeclHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDeclHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDeclHandler")
	}
}

//export callbackQXmlSimpleReaderSetDeclHandler
func callbackQXmlSimpleReaderSetDeclHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setDeclHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDeclHandler"); signal != nil {
		signal.(func(*QXmlDeclHandler))(NewQXmlDeclHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetDeclHandlerDefault(NewQXmlDeclHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandlerDefault(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandlerDefault(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {
	defer qt.Recovering("connect QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setEntityResolver", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetEntityResolver() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setEntityResolver")
	}
}

//export callbackQXmlSimpleReaderSetEntityResolver
func callbackQXmlSimpleReaderSetEntityResolver(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setEntityResolver")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEntityResolver"); signal != nil {
		signal.(func(*QXmlEntityResolver))(NewQXmlEntityResolverFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetEntityResolverDefault(NewQXmlEntityResolverFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolverDefault(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolverDefault(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setErrorHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetErrorHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setErrorHandler")
	}
}

//export callbackQXmlSimpleReaderSetErrorHandler
func callbackQXmlSimpleReaderSetErrorHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setErrorHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setErrorHandler"); signal != nil {
		signal.(func(*QXmlErrorHandler))(NewQXmlErrorHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetErrorHandlerDefault(NewQXmlErrorHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandlerDefault(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandlerDefault(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetFeature(f func(name string, enable bool)) {
	defer qt.Recovering("connect QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFeature", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetFeature() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFeature")
	}
}

//export callbackQXmlSimpleReaderSetFeature
func callbackQXmlSimpleReaderSetFeature(ptr unsafe.Pointer, ptrName *C.char, name *C.char, enable C.int) {
	defer qt.Recovering("callback QXmlSimpleReader::setFeature")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFeature"); signal != nil {
		signal.(func(string, bool))(C.GoString(name), int(enable) != 0)
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetFeatureDefault(C.GoString(name), int(enable) != 0)
	}
}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {
	defer qt.Recovering("QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) SetFeatureDefault(name string, enable bool) {
	defer qt.Recovering("QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeatureDefault(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetLexicalHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler")
	}
}

//export callbackQXmlSimpleReaderSetLexicalHandler
func callbackQXmlSimpleReaderSetLexicalHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setLexicalHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setLexicalHandler"); signal != nil {
		signal.(func(*QXmlLexicalHandler))(NewQXmlLexicalHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetLexicalHandlerDefault(NewQXmlLexicalHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandlerDefault(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandlerDefault(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {
	defer qt.Recovering("QXmlSimpleReader::~QXmlSimpleReader")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_DestroyQXmlSimpleReader(ptr.Pointer())
	}
}

func (ptr *QXmlSimpleReader) ObjectNameAbs() string {
	defer qt.Recovering("QXmlSimpleReader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSimpleReader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlSimpleReader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlSimpleReader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
