package quick

//#include "qsggeometrynode.h"
import "C"
import (
	"unsafe"
)

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNodeITF interface {
	QSGBasicGeometryNodeITF
	QSGGeometryNodePTR() *QSGGeometryNode
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNodePTR().Pointer()
	}
	return nil
}

func QSGGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGGeometryNode {
	var n = new(QSGGeometryNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGGeometryNode) QSGGeometryNodePTR() *QSGGeometryNode {
	return ptr
}

func NewQSGGeometryNode() *QSGGeometryNode {
	return QSGGeometryNodeFromPointer(unsafe.Pointer(C.QSGGeometryNode_NewQSGGeometryNode()))
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {
	if ptr.Pointer() != nil {
		return QSGMaterialFromPointer(unsafe.Pointer(C.QSGGeometryNode_Material(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {
	if ptr.Pointer() != nil {
		return QSGMaterialFromPointer(unsafe.Pointer(C.QSGGeometryNode_OpaqueMaterial(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterialITF) {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetMaterial(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGMaterial(material)))
	}
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterialITF) {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetOpaqueMaterial(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGMaterial(material)))
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
