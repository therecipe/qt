package xml

//#include "qdomentity.h"
import "C"
import (
	"unsafe"
)

type QDomEntity struct {
	QDomNode
}

type QDomEntityITF interface {
	QDomNodeITF
	QDomEntityPTR() *QDomEntity
}

func PointerFromQDomEntity(ptr QDomEntityITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntityPTR().Pointer()
	}
	return nil
}

func QDomEntityFromPointer(ptr unsafe.Pointer) *QDomEntity {
	var n = new(QDomEntity)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomEntity) QDomEntityPTR() *QDomEntity {
	return ptr
}

func NewQDomEntity() *QDomEntity {
	return QDomEntityFromPointer(unsafe.Pointer(C.QDomEntity_NewQDomEntity()))
}

func NewQDomEntity2(x QDomEntityITF) *QDomEntity {
	return QDomEntityFromPointer(unsafe.Pointer(C.QDomEntity_NewQDomEntity2(C.QtObjectPtr(PointerFromQDomEntity(x)))))
}

func (ptr *QDomEntity) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntity_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomEntity) NotationName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_NotationName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomEntity) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_PublicId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomEntity) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_SystemId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
