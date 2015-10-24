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

type QXmlFormatterITF interface {
	QXmlSerializerITF
	QXmlFormatterPTR() *QXmlFormatter
}

func PointerFromQXmlFormatter(ptr QXmlFormatterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlFormatterPTR().Pointer()
	}
	return nil
}

func QXmlFormatterFromPointer(ptr unsafe.Pointer) *QXmlFormatter {
	var n = new(QXmlFormatter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlFormatter) QXmlFormatterPTR() *QXmlFormatter {
	return ptr
}

func NewQXmlFormatter(query QXmlQueryITF, outputDevice core.QIODeviceITF) *QXmlFormatter {
	return QXmlFormatterFromPointer(unsafe.Pointer(C.QXmlFormatter_NewQXmlFormatter(C.QtObjectPtr(PointerFromQXmlQuery(query)), C.QtObjectPtr(core.PointerFromQIODevice(outputDevice)))))
}

func (ptr *QXmlFormatter) Attribute(name QXmlNameITF, value core.QStringRefITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Attribute(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.QtObjectPtr(core.PointerFromQStringRef(value)))
	}
}

func (ptr *QXmlFormatter) Characters(value core.QStringRefITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Characters(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQStringRef(value)))
	}
}

func (ptr *QXmlFormatter) Comment(value string) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Comment(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QXmlFormatter) EndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlFormatter) EndElement() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElement(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlFormatter) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequence(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlFormatter_IndentationDepth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlFormatter) ProcessingInstruction(name QXmlNameITF, value string) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_ProcessingInstruction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.CString(value))
	}
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(C.QtObjectPtr(ptr.Pointer()), C.int(depth))
	}
}

func (ptr *QXmlFormatter) StartDocument() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlFormatter) StartElement(name QXmlNameITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElement(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)))
	}
}

func (ptr *QXmlFormatter) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequence(C.QtObjectPtr(ptr.Pointer()))
	}
}
