package xml

//#include "qdomentityreference.h"
import "C"
import (
	"unsafe"
)

type QDomEntityReference struct {
	QDomNode
}

type QDomEntityReferenceITF interface {
	QDomNodeITF
	QDomEntityReferencePTR() *QDomEntityReference
}

func PointerFromQDomEntityReference(ptr QDomEntityReferenceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntityReferencePTR().Pointer()
	}
	return nil
}

func QDomEntityReferenceFromPointer(ptr unsafe.Pointer) *QDomEntityReference {
	var n = new(QDomEntityReference)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomEntityReference) QDomEntityReferencePTR() *QDomEntityReference {
	return ptr
}

func NewQDomEntityReference() *QDomEntityReference {
	return QDomEntityReferenceFromPointer(unsafe.Pointer(C.QDomEntityReference_NewQDomEntityReference()))
}

func NewQDomEntityReference2(x QDomEntityReferenceITF) *QDomEntityReference {
	return QDomEntityReferenceFromPointer(unsafe.Pointer(C.QDomEntityReference_NewQDomEntityReference2(C.QtObjectPtr(PointerFromQDomEntityReference(x)))))
}

func (ptr *QDomEntityReference) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntityReference_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
