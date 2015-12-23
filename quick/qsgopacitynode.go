package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNode_ITF interface {
	QSGNode_ITF
	QSGOpacityNode_PTR() *QSGOpacityNode
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNode_PTR().Pointer()
	}
	return nil
}

func NewQSGOpacityNodeFromPointer(ptr unsafe.Pointer) *QSGOpacityNode {
	var n = new(QSGOpacityNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return ptr
}

func NewQSGOpacityNode() *QSGOpacityNode {
	defer qt.Recovering("QSGOpacityNode::QSGOpacityNode")

	return NewQSGOpacityNodeFromPointer(C.QSGOpacityNode_NewQSGOpacityNode())
}

func (ptr *QSGOpacityNode) Opacity() float64 {
	defer qt.Recovering("QSGOpacityNode::opacity")

	if ptr.Pointer() != nil {
		return float64(C.QSGOpacityNode_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpacityNode) SetOpacity(opacity float64) {
	defer qt.Recovering("QSGOpacityNode::setOpacity")

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {
	defer qt.Recovering("QSGOpacityNode::~QSGOpacityNode")

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNode(ptr.Pointer())
	}
}

func (ptr *QSGOpacityNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGOpacityNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGOpacityNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGOpacityNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGOpacityNodePreprocess
func callbackQSGOpacityNodePreprocess(ptrName *C.char) bool {
	defer qt.Recovering("callback QSGOpacityNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}
