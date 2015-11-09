package xmlpatterns

//#include "qxmlformatter.h"
import "C"
import (
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
	return n
}

func (ptr *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return ptr
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlFormatter) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) Characters(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) Comment(value string) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlFormatter) EndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndElement() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlFormatter_IndentationDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlFormatter) ProcessingInstruction(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), C.CString(value))
	}
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QXmlFormatter) StartDocument() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartElement(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlFormatter) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequence(ptr.Pointer())
	}
}
