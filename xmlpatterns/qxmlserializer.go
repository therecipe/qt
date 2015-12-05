package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	return n
}

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::QXmlSerializer")
		}
	}()

	return NewQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlSerializer) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::attribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::characters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) Comment(value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::comment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlSerializer) EndDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::endDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndElement() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::endElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::codec")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSerializer) EndOfSequence() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::endOfSequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) NamespaceBinding(nb QXmlName_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::namespaceBinding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::outputDevice")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QXmlSerializer_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSerializer) ProcessingInstruction(name QXmlName_ITF, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::processingInstruction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), C.CString(value))
	}
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::setCodec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(ptr.Pointer(), core.PointerFromQTextCodec(outputCodec))
	}
}

func (ptr *QXmlSerializer) StartDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::startDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartElement(name QXmlName_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::startElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlSerializer) StartOfSequence() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSerializer::startOfSequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(ptr.Pointer())
	}
}
