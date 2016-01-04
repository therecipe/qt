package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QXmlSerializer struct {
	QAbstractXmlReceiver
}

type QXmlSerializer_ITF interface {
	QAbstractXmlReceiver_ITF
	QXmlSerializer_PTR() *QXmlSerializer
}

func PointerFromQXmlSerializer(ptr QXmlSerializer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func NewQXmlSerializerFromPointer(ptr unsafe.Pointer) *QXmlSerializer {
	var n = new(QXmlSerializer)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlSerializer_") {
		n.SetObjectNameAbs("QXmlSerializer_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	defer qt.Recovering("QXmlSerializer::QXmlSerializer")

	return NewQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlSerializer::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlSerializer) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlSerializer::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

//export callbackQXmlSerializerCharacters
func callbackQXmlSerializerCharacters(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::characters")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) CharactersDefault(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::characters")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QXmlSerializer::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlSerializer) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlSerializer::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

//export callbackQXmlSerializerComment
func callbackQXmlSerializerComment(ptr unsafe.Pointer, ptrName *C.char, value *C.char) {
	defer qt.Recovering("callback QXmlSerializer::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
	}

}

func (ptr *QXmlSerializer) Comment(value string) {
	defer qt.Recovering("QXmlSerializer::comment")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlSerializer) CommentDefault(value string) {
	defer qt.Recovering("QXmlSerializer::comment")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_CommentDefault(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

//export callbackQXmlSerializerEndDocument
func callbackQXmlSerializerEndDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) EndDocument() {
	defer qt.Recovering("QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndDocumentDefault() {
	defer qt.Recovering("QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocumentDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlSerializer::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

//export callbackQXmlSerializerEndElement
func callbackQXmlSerializerEndElement(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) EndElement() {
	defer qt.Recovering("QXmlSerializer::endElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndElementDefault() {
	defer qt.Recovering("QXmlSerializer::endElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElementDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	defer qt.Recovering("QXmlSerializer::codec")

	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

//export callbackQXmlSerializerEndOfSequence
func callbackQXmlSerializerEndOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) EndOfSequence() {
	defer qt.Recovering("QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndOfSequenceDefault() {
	defer qt.Recovering("QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	defer qt.Recovering("QXmlSerializer::outputDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QXmlSerializer_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {
	defer qt.Recovering("QXmlSerializer::setCodec")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(ptr.Pointer(), core.PointerFromQTextCodec(outputCodec))
	}
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

//export callbackQXmlSerializerStartDocument
func callbackQXmlSerializerStartDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) StartDocument() {
	defer qt.Recovering("QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartDocumentDefault() {
	defer qt.Recovering("QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocumentDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

//export callbackQXmlSerializerStartOfSequence
func callbackQXmlSerializerStartOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXmlSerializer) StartOfSequence() {
	defer qt.Recovering("QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartOfSequenceDefault() {
	defer qt.Recovering("QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ObjectNameAbs() string {
	defer qt.Recovering("QXmlSerializer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSerializer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlSerializer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlSerializer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
