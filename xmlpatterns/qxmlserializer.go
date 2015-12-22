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
func callbackQXmlSerializerCharacters(ptrName *C.char, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QXmlSerializer::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
		return true
	}
	return false

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
func callbackQXmlSerializerComment(ptrName *C.char, value *C.char) bool {
	defer qt.Recovering("callback QXmlSerializer::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
		return true
	}
	return false

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
func callbackQXmlSerializerEndDocument(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlSerializer::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQXmlSerializerEndElement(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlSerializer::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQXmlSerializerEndOfSequence(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlSerializer::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQXmlSerializerStartDocument(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlSerializer::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQXmlSerializerStartOfSequence(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlSerializer::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
