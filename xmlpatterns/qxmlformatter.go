package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QXmlFormatter struct {
	QXmlSerializer
}

type QXmlFormatter_ITF interface {
	QXmlSerializer_ITF
	QXmlFormatter_PTR() *QXmlFormatter
}

func PointerFromQXmlFormatter(ptr QXmlFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlFormatter_PTR().Pointer()
	}
	return nil
}

func NewQXmlFormatterFromPointer(ptr unsafe.Pointer) *QXmlFormatter {
	var n = new(QXmlFormatter)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlFormatter_") {
		n.SetObjectNameAbs("QXmlFormatter_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return ptr
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	defer qt.Recovering("QXmlFormatter::QXmlFormatter")

	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlFormatter) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlFormatter::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlFormatter) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlFormatter::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

//export callbackQXmlFormatterCharacters
func callbackQXmlFormatterCharacters(ptrName *C.char, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QXmlFormatter::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
		return true
	}
	return false

}

func (ptr *QXmlFormatter) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QXmlFormatter::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlFormatter) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlFormatter::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

//export callbackQXmlFormatterComment
func callbackQXmlFormatterComment(ptrName *C.char, value *C.char) bool {
	defer qt.Recovering("callback QXmlFormatter::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
		return true
	}
	return false

}

func (ptr *QXmlFormatter) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

//export callbackQXmlFormatterEndDocument
func callbackQXmlFormatterEndDocument(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlFormatter::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QXmlFormatter) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlFormatter::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

//export callbackQXmlFormatterEndElement
func callbackQXmlFormatterEndElement(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlFormatter::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QXmlFormatter) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

//export callbackQXmlFormatterEndOfSequence
func callbackQXmlFormatterEndOfSequence(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlFormatter::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QXmlFormatter) IndentationDepth() int {
	defer qt.Recovering("QXmlFormatter::indentationDepth")

	if ptr.Pointer() != nil {
		return int(C.QXmlFormatter_IndentationDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	defer qt.Recovering("QXmlFormatter::setIndentationDepth")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QXmlFormatter) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

//export callbackQXmlFormatterStartDocument
func callbackQXmlFormatterStartDocument(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlFormatter::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QXmlFormatter) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

//export callbackQXmlFormatterStartOfSequence
func callbackQXmlFormatterStartOfSequence(ptrName *C.char) bool {
	defer qt.Recovering("callback QXmlFormatter::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QXmlFormatter) ObjectNameAbs() string {
	defer qt.Recovering("QXmlFormatter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlFormatter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlFormatter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlFormatter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
