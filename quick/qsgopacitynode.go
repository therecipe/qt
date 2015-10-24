package quick

//#include "qsgopacitynode.h"
import "C"
import (
	"unsafe"
)

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNodeITF interface {
	QSGNodeITF
	QSGOpacityNodePTR() *QSGOpacityNode
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNodePTR().Pointer()
	}
	return nil
}

func QSGOpacityNodeFromPointer(ptr unsafe.Pointer) *QSGOpacityNode {
	var n = new(QSGOpacityNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGOpacityNode) QSGOpacityNodePTR() *QSGOpacityNode {
	return ptr
}

func NewQSGOpacityNode() *QSGOpacityNode {
	return QSGOpacityNodeFromPointer(unsafe.Pointer(C.QSGOpacityNode_NewQSGOpacityNode()))
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
