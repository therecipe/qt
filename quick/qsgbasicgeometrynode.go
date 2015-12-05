package quick

//#include "quick.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func PointerFromQSGBasicGeometryNode(ptr QSGBasicGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGBasicGeometryNode {
	var n = new(QSGBasicGeometryNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode {
	return ptr
}

func (ptr *QSGBasicGeometryNode) Geometry2() *QSGGeometry {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGBasicGeometryNode::geometry")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGBasicGeometryNode::geometry")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGBasicGeometryNode::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_SetGeometry(ptr.Pointer(), PointerFromQSGGeometry(geometry))
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGBasicGeometryNode::~QSGBasicGeometryNode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(ptr.Pointer())
	}
}
