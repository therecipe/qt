package quick

//#include "qsggeometrynode.h"
import "C"
import (
	"unsafe"
)

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGGeometryNode {
	var n = new(QSGGeometryNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return ptr
}

func NewQSGGeometryNode() *QSGGeometryNode {
	return NewQSGGeometryNodeFromPointer(C.QSGGeometryNode_NewQSGGeometryNode())
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {
	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_Material(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {
	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_OpaqueMaterial(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF) {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF) {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetOpaqueMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(ptr.Pointer())
	}
}
