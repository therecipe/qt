package xml

//#include "qdomnotation.h"
import "C"
import (
	"unsafe"
)

type QDomNotation struct {
	QDomNode
}

type QDomNotationITF interface {
	QDomNodeITF
	QDomNotationPTR() *QDomNotation
}

func PointerFromQDomNotation(ptr QDomNotationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNotationPTR().Pointer()
	}
	return nil
}

func QDomNotationFromPointer(ptr unsafe.Pointer) *QDomNotation {
	var n = new(QDomNotation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNotation) QDomNotationPTR() *QDomNotation {
	return ptr
}

func NewQDomNotation() *QDomNotation {
	return QDomNotationFromPointer(unsafe.Pointer(C.QDomNotation_NewQDomNotation()))
}

func NewQDomNotation2(x QDomNotationITF) *QDomNotation {
	return QDomNotationFromPointer(unsafe.Pointer(C.QDomNotation_NewQDomNotation2(C.QtObjectPtr(PointerFromQDomNotation(x)))))
}

func (ptr *QDomNotation) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNotation_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNotation) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_PublicId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomNotation) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_SystemId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
