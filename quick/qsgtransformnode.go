package quick

//#include "qsgtransformnode.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNodeITF interface {
	QSGNodeITF
	QSGTransformNodePTR() *QSGTransformNode
}

func PointerFromQSGTransformNode(ptr QSGTransformNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTransformNodePTR().Pointer()
	}
	return nil
}

func QSGTransformNodeFromPointer(ptr unsafe.Pointer) *QSGTransformNode {
	var n = new(QSGTransformNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGTransformNode) QSGTransformNodePTR() *QSGTransformNode {
	return ptr
}

func NewQSGTransformNode() *QSGTransformNode {
	return QSGTransformNodeFromPointer(unsafe.Pointer(C.QSGTransformNode_NewQSGTransformNode()))
}

func (ptr *QSGTransformNode) SetMatrix(matrix gui.QMatrix4x4ITF) {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_SetMatrix(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQMatrix4x4(matrix)))
	}
}

func (ptr *QSGTransformNode) DestroyQSGTransformNode() {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_DestroyQSGTransformNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
