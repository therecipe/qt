package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGGeometryNode {
	var n = new(QSGGeometryNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return ptr
}

func NewQSGGeometryNode() *QSGGeometryNode {
	defer qt.Recovering("QSGGeometryNode::QSGGeometryNode")

	return NewQSGGeometryNodeFromPointer(C.QSGGeometryNode_NewQSGGeometryNode())
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {
	defer qt.Recovering("QSGGeometryNode::material")

	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_Material(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {
	defer qt.Recovering("QSGGeometryNode::opaqueMaterial")

	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_OpaqueMaterial(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF) {
	defer qt.Recovering("QSGGeometryNode::setMaterial")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF) {
	defer qt.Recovering("QSGGeometryNode::setOpaqueMaterial")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetOpaqueMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	defer qt.Recovering("QSGGeometryNode::~QSGGeometryNode")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(ptr.Pointer())
	}
}

func (ptr *QSGGeometryNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGGeometryNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGGeometryNodePreprocess
func callbackQSGGeometryNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGGeometryNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGGeometryNode) Preprocess() {
	defer qt.Recovering("QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGGeometryNode) PreprocessDefault() {
	defer qt.Recovering("QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_PreprocessDefault(ptr.Pointer())
	}
}
