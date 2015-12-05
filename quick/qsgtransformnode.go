package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func PointerFromQSGTransformNode(ptr QSGTransformNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTransformNode_PTR().Pointer()
	}
	return nil
}

func NewQSGTransformNodeFromPointer(ptr unsafe.Pointer) *QSGTransformNode {
	var n = new(QSGTransformNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode {
	return ptr
}

func NewQSGTransformNode() *QSGTransformNode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTransformNode::QSGTransformNode")
		}
	}()

	return NewQSGTransformNodeFromPointer(C.QSGTransformNode_NewQSGTransformNode())
}

func (ptr *QSGTransformNode) SetMatrix(matrix gui.QMatrix4x4_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTransformNode::setMatrix")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGTransformNode_SetMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QSGTransformNode) DestroyQSGTransformNode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTransformNode::~QSGTransformNode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGTransformNode_DestroyQSGTransformNode(ptr.Pointer())
	}
}
