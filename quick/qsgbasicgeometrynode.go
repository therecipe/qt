package quick

//#include "qsgbasicgeometrynode.h"
import "C"
import (
	"unsafe"
)

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNodeITF interface {
	QSGNodeITF
	QSGBasicGeometryNodePTR() *QSGBasicGeometryNode
}

func PointerFromQSGBasicGeometryNode(ptr QSGBasicGeometryNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNodePTR().Pointer()
	}
	return nil
}

func QSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGBasicGeometryNode {
	var n = new(QSGBasicGeometryNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGBasicGeometryNode) QSGBasicGeometryNodePTR() *QSGBasicGeometryNode {
	return ptr
}

func (ptr *QSGBasicGeometryNode) Geometry2() *QSGGeometry {
	if ptr.Pointer() != nil {
		return QSGGeometryFromPointer(unsafe.Pointer(C.QSGBasicGeometryNode_Geometry2(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {
	if ptr.Pointer() != nil {
		return QSGGeometryFromPointer(unsafe.Pointer(C.QSGBasicGeometryNode_Geometry(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometryITF) {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGGeometry(geometry)))
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
