package quick

//#include "qsgsimplerectnode.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNodeITF interface {
	QSGGeometryNodeITF
	QSGSimpleRectNodePTR() *QSGSimpleRectNode
}

func PointerFromQSGSimpleRectNode(ptr QSGSimpleRectNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleRectNodePTR().Pointer()
	}
	return nil
}

func QSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleRectNode {
	var n = new(QSGSimpleRectNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleRectNode) QSGSimpleRectNodePTR() *QSGSimpleRectNode {
	return ptr
}

func NewQSGSimpleRectNode2() *QSGSimpleRectNode {
	return QSGSimpleRectNodeFromPointer(unsafe.Pointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode2()))
}

func NewQSGSimpleRectNode(rect core.QRectFITF, color gui.QColorITF) *QSGSimpleRectNode {
	return QSGSimpleRectNodeFromPointer(unsafe.Pointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode(C.QtObjectPtr(core.PointerFromQRectF(rect)), C.QtObjectPtr(gui.PointerFromQColor(color)))))
}

func (ptr *QSGSimpleRectNode) SetColor(color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func (ptr *QSGSimpleRectNode) SetRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}
