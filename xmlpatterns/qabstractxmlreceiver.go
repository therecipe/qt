package xmlpatterns

//#include "qabstractxmlreceiver.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractXmlReceiver struct {
	ptr unsafe.Pointer
}

type QAbstractXmlReceiverITF interface {
	QAbstractXmlReceiverPTR() *QAbstractXmlReceiver
}

func (p *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractXmlReceiver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractXmlReceiver(ptr QAbstractXmlReceiverITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiverPTR().Pointer()
	}
	return nil
}

func QAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) *QAbstractXmlReceiver {
	var n = new(QAbstractXmlReceiver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractXmlReceiver) QAbstractXmlReceiverPTR() *QAbstractXmlReceiver {
	return ptr
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlNameITF, value core.QStringRefITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Attribute(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.QtObjectPtr(core.PointerFromQStringRef(value)))
	}
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRefITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Characters(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQStringRef(value)))
	}
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Comment(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QAbstractXmlReceiver) EndDocument() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractXmlReceiver) EndElement() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndElement(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndOfSequence(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlNameITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_NamespaceBinding(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)))
	}
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlNameITF, value string) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_ProcessingInstruction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(target)), C.CString(value))
	}
}

func (ptr *QAbstractXmlReceiver) StartDocument() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlNameITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartElement(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)))
	}
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartOfSequence(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(C.QtObjectPtr(ptr.Pointer()))
	}
}
