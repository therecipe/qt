package xmlpatterns

//#include "qxmlserializer.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QXmlSerializer struct {
	QAbstractXmlReceiver
}

type QXmlSerializerITF interface {
	QAbstractXmlReceiverITF
	QXmlSerializerPTR() *QXmlSerializer
}

func PointerFromQXmlSerializer(ptr QXmlSerializerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializerPTR().Pointer()
	}
	return nil
}

func QXmlSerializerFromPointer(ptr unsafe.Pointer) *QXmlSerializer {
	var n = new(QXmlSerializer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSerializer) QXmlSerializerPTR() *QXmlSerializer {
	return ptr
}

func NewQXmlSerializer(query QXmlQueryITF, outputDevice core.QIODeviceITF) *QXmlSerializer {
	return QXmlSerializerFromPointer(unsafe.Pointer(C.QXmlSerializer_NewQXmlSerializer(C.QtObjectPtr(PointerFromQXmlQuery(query)), C.QtObjectPtr(core.PointerFromQIODevice(outputDevice)))))
}

func (ptr *QXmlSerializer) Attribute(name QXmlNameITF, value core.QStringRefITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Attribute(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.QtObjectPtr(core.PointerFromQStringRef(value)))
	}
}

func (ptr *QXmlSerializer) Characters(value core.QStringRefITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQStringRef(value)))
	}
}

func (ptr *QXmlSerializer) Comment(value string) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Comment(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QXmlSerializer) EndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlSerializer) EndElement() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	if ptr.Pointer() != nil {
		return core.QTextCodecFromPointer(unsafe.Pointer(C.QXmlSerializer_Codec(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSerializer) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlSerializer) NamespaceBinding(nb QXmlNameITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBinding(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(nb)))
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QXmlSerializer_OutputDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSerializer) ProcessingInstruction(name QXmlNameITF, value string) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_ProcessingInstruction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.CString(value))
	}
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodecITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQTextCodec(outputCodec)))
	}
}

func (ptr *QXmlSerializer) StartDocument() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlSerializer) StartElement(name QXmlNameITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElement(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)))
	}
}

func (ptr *QXmlSerializer) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(C.QtObjectPtr(ptr.Pointer()))
	}
}
