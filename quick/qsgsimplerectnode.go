package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleRectNode_PTR() *QSGSimpleRectNode
}

func PointerFromQSGSimpleRectNode(ptr QSGSimpleRectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleRectNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleRectNode {
	var n = new(QSGSimpleRectNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleRectNode) QSGSimpleRectNode_PTR() *QSGSimpleRectNode {
	return ptr
}

func NewQSGSimpleRectNode2() *QSGSimpleRectNode {
	defer qt.Recovering("QSGSimpleRectNode::QSGSimpleRectNode")

	return NewQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode2())
}

func NewQSGSimpleRectNode(rect core.QRectF_ITF, color gui.QColor_ITF) *QSGSimpleRectNode {
	defer qt.Recovering("QSGSimpleRectNode::QSGSimpleRectNode")

	return NewQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode(core.PointerFromQRectF(rect), gui.PointerFromQColor(color)))
}

func (ptr *QSGSimpleRectNode) Color() *gui.QColor {
	defer qt.Recovering("QSGSimpleRectNode::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGSimpleRectNode_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleRectNode) SetColor(color gui.QColor_ITF) {
	defer qt.Recovering("QSGSimpleRectNode::setColor")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGSimpleRectNode) SetRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QSGSimpleRectNode::setRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGSimpleRectNode) SetRect2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QSGSimpleRectNode::setRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleRectNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGSimpleRectNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGSimpleRectNodePreprocess
func callbackQSGSimpleRectNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGSimpleRectNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleRectNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGSimpleRectNode) Preprocess() {
	defer qt.Recovering("QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGSimpleRectNode) PreprocessDefault() {
	defer qt.Recovering("QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_PreprocessDefault(ptr.Pointer())
	}
}
