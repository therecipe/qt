package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAbstractXmlReceiver struct {
	ptr unsafe.Pointer
}

type QAbstractXmlReceiver_ITF interface {
	QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver
}

func (p *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractXmlReceiver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractXmlReceiver(ptr QAbstractXmlReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) *QAbstractXmlReceiver {
	var n = new(QAbstractXmlReceiver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractXmlReceiver) QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver {
	return ptr
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::attribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRef_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::characters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::comment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QAbstractXmlReceiver) EndDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::endDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndDocument(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) EndElement() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::endElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndElement(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::endOfSequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlName_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::namespaceBinding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlName_ITF, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::processingInstruction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(target), C.CString(value))
	}
}

func (ptr *QAbstractXmlReceiver) StartDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::startDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartDocument(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlName_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::startElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::startOfSequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlReceiver::~QAbstractXmlReceiver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr.Pointer())
	}
}
