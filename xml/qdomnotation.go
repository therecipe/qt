package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDomNotation struct {
	QDomNode
}

type QDomNotation_ITF interface {
	QDomNode_ITF
	QDomNotation_PTR() *QDomNotation
}

func PointerFromQDomNotation(ptr QDomNotation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNotation_PTR().Pointer()
	}
	return nil
}

func NewQDomNotationFromPointer(ptr unsafe.Pointer) *QDomNotation {
	var n = new(QDomNotation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNotation) QDomNotation_PTR() *QDomNotation {
	return ptr
}

func NewQDomNotation() *QDomNotation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNotation::QDomNotation")
		}
	}()

	return NewQDomNotationFromPointer(C.QDomNotation_NewQDomNotation())
}

func NewQDomNotation2(x QDomNotation_ITF) *QDomNotation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNotation::QDomNotation")
		}
	}()

	return NewQDomNotationFromPointer(C.QDomNotation_NewQDomNotation2(PointerFromQDomNotation(x)))
}

func (ptr *QDomNotation) NodeType() QDomNode__NodeType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNotation::nodeType")
		}
	}()

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNotation_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNotation) PublicId() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNotation::publicId")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNotation) SystemId() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNotation::systemId")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_SystemId(ptr.Pointer()))
	}
	return ""
}
