package quick

//#include "quick.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGGeometryNode::QSGGeometryNode")
		}
	}()

	return NewQSGGeometryNodeFromPointer(C.QSGGeometryNode_NewQSGGeometryNode())
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGGeometryNode::material")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_Material(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGGeometryNode::opaqueMaterial")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_OpaqueMaterial(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGGeometryNode::setMaterial")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGGeometryNode::setOpaqueMaterial")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetOpaqueMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGGeometryNode::~QSGGeometryNode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(ptr.Pointer())
	}
}
