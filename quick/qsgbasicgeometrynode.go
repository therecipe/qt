package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QSGBasicGeometryNode::geometry")

	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {
	defer qt.Recovering("QSGBasicGeometryNode::geometry")

	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF) {
	defer qt.Recovering("QSGBasicGeometryNode::setGeometry")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_SetGeometry(ptr.Pointer(), PointerFromQSGGeometry(geometry))
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {
	defer qt.Recovering("QSGBasicGeometryNode::~QSGBasicGeometryNode")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(ptr.Pointer())
	}
}

func (ptr *QSGBasicGeometryNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGBasicGeometryNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGBasicGeometryNodePreprocess
func callbackQSGBasicGeometryNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGBasicGeometryNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGBasicGeometryNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGBasicGeometryNode) Preprocess() {
	defer qt.Recovering("QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGBasicGeometryNode) PreprocessDefault() {
	defer qt.Recovering("QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_PreprocessDefault(ptr.Pointer())
	}
}
