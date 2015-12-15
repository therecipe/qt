package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGClipNode struct {
	QSGBasicGeometryNode
}

type QSGClipNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGClipNode_PTR() *QSGClipNode
}

func PointerFromQSGClipNode(ptr QSGClipNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGClipNode_PTR().Pointer()
	}
	return nil
}

func NewQSGClipNodeFromPointer(ptr unsafe.Pointer) *QSGClipNode {
	var n = new(QSGClipNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGClipNode) QSGClipNode_PTR() *QSGClipNode {
	return ptr
}

func NewQSGClipNode() *QSGClipNode {
	defer qt.Recovering("QSGClipNode::QSGClipNode")

	return NewQSGClipNodeFromPointer(C.QSGClipNode_NewQSGClipNode())
}

func (ptr *QSGClipNode) IsRectangular() bool {
	defer qt.Recovering("QSGClipNode::isRectangular")

	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsRectangular(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGClipNode) SetClipRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QSGClipNode::setClipRect")

	if ptr.Pointer() != nil {
		C.QSGClipNode_SetClipRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGClipNode) SetIsRectangular(rectHint bool) {
	defer qt.Recovering("QSGClipNode::setIsRectangular")

	if ptr.Pointer() != nil {
		C.QSGClipNode_SetIsRectangular(ptr.Pointer(), C.int(qt.GoBoolToInt(rectHint)))
	}
}

func (ptr *QSGClipNode) DestroyQSGClipNode() {
	defer qt.Recovering("QSGClipNode::~QSGClipNode")

	if ptr.Pointer() != nil {
		C.QSGClipNode_DestroyQSGClipNode(ptr.Pointer())
	}
}
