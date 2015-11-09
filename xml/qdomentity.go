package xml

//#include "qdomentity.h"
import "C"
import (
	"unsafe"
)

type QDomEntity struct {
	QDomNode
}

type QDomEntity_ITF interface {
	QDomNode_ITF
	QDomEntity_PTR() *QDomEntity
}

func PointerFromQDomEntity(ptr QDomEntity_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntity_PTR().Pointer()
	}
	return nil
}

func NewQDomEntityFromPointer(ptr unsafe.Pointer) *QDomEntity {
	var n = new(QDomEntity)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomEntity) QDomEntity_PTR() *QDomEntity {
	return ptr
}

func NewQDomEntity() *QDomEntity {
	return NewQDomEntityFromPointer(C.QDomEntity_NewQDomEntity())
}

func NewQDomEntity2(x QDomEntity_ITF) *QDomEntity {
	return NewQDomEntityFromPointer(C.QDomEntity_NewQDomEntity2(PointerFromQDomEntity(x)))
}

func (ptr *QDomEntity) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntity_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomEntity) NotationName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_NotationName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_SystemId(ptr.Pointer()))
	}
	return ""
}
