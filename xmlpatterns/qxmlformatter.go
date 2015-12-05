package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::QXmlFormatter")
		}
	}()

	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlFormatter) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::attribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) Characters(value core.QStringRef_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::characters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) Comment(value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::comment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Comment(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QXmlFormatter) EndDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::endDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndElement() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::endElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndOfSequence() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::endOfSequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::indentationDepth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlFormatter_IndentationDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlFormatter) ProcessingInstruction(name QXmlName_ITF, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::processingInstruction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), C.CString(value))
	}
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::setIndentationDepth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QXmlFormatter) StartDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::startDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartElement(name QXmlName_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::startElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlFormatter) StartOfSequence() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlFormatter::startOfSequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequence(ptr.Pointer())
	}
}
