package quick

//#include "qsgclipnode.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGClipNode struct {
	QSGBasicGeometryNode
}

type QSGClipNodeITF interface {
	QSGBasicGeometryNodeITF
	QSGClipNodePTR() *QSGClipNode
}

func PointerFromQSGClipNode(ptr QSGClipNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGClipNodePTR().Pointer()
	}
	return nil
}

func QSGClipNodeFromPointer(ptr unsafe.Pointer) *QSGClipNode {
	var n = new(QSGClipNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGClipNode) QSGClipNodePTR() *QSGClipNode {
	return ptr
}

func NewQSGClipNode() *QSGClipNode {
	return QSGClipNodeFromPointer(unsafe.Pointer(C.QSGClipNode_NewQSGClipNode()))
}

func (ptr *QSGClipNode) IsRectangular() bool {
	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsRectangular(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGClipNode) SetClipRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSGClipNode_SetClipRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QSGClipNode) SetIsRectangular(rectHint bool) {
	if ptr.Pointer() != nil {
		C.QSGClipNode_SetIsRectangular(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(rectHint)))
	}
}

func (ptr *QSGClipNode) DestroyQSGClipNode() {
	if ptr.Pointer() != nil {
		C.QSGClipNode_DestroyQSGClipNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
