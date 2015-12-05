package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDomEntityReference struct {
	QDomNode
}

type QDomEntityReference_ITF interface {
	QDomNode_ITF
	QDomEntityReference_PTR() *QDomEntityReference
}

func PointerFromQDomEntityReference(ptr QDomEntityReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntityReference_PTR().Pointer()
	}
	return nil
}

func NewQDomEntityReferenceFromPointer(ptr unsafe.Pointer) *QDomEntityReference {
	var n = new(QDomEntityReference)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomEntityReference) QDomEntityReference_PTR() *QDomEntityReference {
	return ptr
}

func NewQDomEntityReference() *QDomEntityReference {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomEntityReference::QDomEntityReference")
		}
	}()

	return NewQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference())
}

func NewQDomEntityReference2(x QDomEntityReference_ITF) *QDomEntityReference {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomEntityReference::QDomEntityReference")
		}
	}()

	return NewQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference2(PointerFromQDomEntityReference(x)))
}

func (ptr *QDomEntityReference) NodeType() QDomNode__NodeType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomEntityReference::nodeType")
		}
	}()

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntityReference_NodeType(ptr.Pointer()))
	}
	return 0
}
