package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSGNode struct {
	ptr unsafe.Pointer
}

type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (p *QSGNode) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGNode) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGNode(ptr QSGNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func NewQSGNodeFromPointer(ptr unsafe.Pointer) *QSGNode {
	var n = new(QSGNode)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSGNode_") {
		n.SetObjectNameAbs("QSGNode_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGNode) QSGNode_PTR() *QSGNode {
	return ptr
}

//QSGNode::DirtyStateBit
type QSGNode__DirtyStateBit int64

const (
	QSGNode__DirtySubtreeBlocked = QSGNode__DirtyStateBit(0x0080)
	QSGNode__DirtyMatrix         = QSGNode__DirtyStateBit(0x0100)
	QSGNode__DirtyNodeAdded      = QSGNode__DirtyStateBit(0x0400)
	QSGNode__DirtyNodeRemoved    = QSGNode__DirtyStateBit(0x0800)
	QSGNode__DirtyGeometry       = QSGNode__DirtyStateBit(0x1000)
	QSGNode__DirtyMaterial       = QSGNode__DirtyStateBit(0x2000)
	QSGNode__DirtyOpacity        = QSGNode__DirtyStateBit(0x4000)
)

//QSGNode::Flag
type QSGNode__Flag int64

const (
	QSGNode__OwnedByParent      = QSGNode__Flag(0x0001)
	QSGNode__UsePreprocess      = QSGNode__Flag(0x0002)
	QSGNode__OwnsGeometry       = QSGNode__Flag(0x00010000)
	QSGNode__OwnsMaterial       = QSGNode__Flag(0x00020000)
	QSGNode__OwnsOpaqueMaterial = QSGNode__Flag(0x00040000)
	QSGNode__InternalReserved   = QSGNode__Flag(0x01000000)
)

//QSGNode::NodeType
type QSGNode__NodeType int64

const (
	QSGNode__BasicNodeType     = QSGNode__NodeType(0)
	QSGNode__GeometryNodeType  = QSGNode__NodeType(1)
	QSGNode__TransformNodeType = QSGNode__NodeType(2)
	QSGNode__ClipNodeType      = QSGNode__NodeType(3)
	QSGNode__OpacityNodeType   = QSGNode__NodeType(4)
)

func (ptr *QSGNode) ChildAtIndex(i int) *QSGNode {
	defer qt.Recovering("QSGNode::childAtIndex")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_ChildAtIndex(ptr.Pointer(), C.int(i)))
	}
	return nil
}

func (ptr *QSGNode) ChildCount() int {
	defer qt.Recovering("QSGNode::childCount")

	if ptr.Pointer() != nil {
		return int(C.QSGNode_ChildCount(ptr.Pointer()))
	}
	return 0
}

func NewQSGNode() *QSGNode {
	defer qt.Recovering("QSGNode::QSGNode")

	return NewQSGNodeFromPointer(C.QSGNode_NewQSGNode())
}

func (ptr *QSGNode) AppendChildNode(node QSGNode_ITF) {
	defer qt.Recovering("QSGNode::appendChildNode")

	if ptr.Pointer() != nil {
		C.QSGNode_AppendChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) FirstChild() *QSGNode {
	defer qt.Recovering("QSGNode::firstChild")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_FirstChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Flags() QSGNode__Flag {
	defer qt.Recovering("QSGNode::flags")

	if ptr.Pointer() != nil {
		return QSGNode__Flag(C.QSGNode_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGNode) InsertChildNodeAfter(node QSGNode_ITF, after QSGNode_ITF) {
	defer qt.Recovering("QSGNode::insertChildNodeAfter")

	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeAfter(ptr.Pointer(), PointerFromQSGNode(node), PointerFromQSGNode(after))
	}
}

func (ptr *QSGNode) InsertChildNodeBefore(node QSGNode_ITF, before QSGNode_ITF) {
	defer qt.Recovering("QSGNode::insertChildNodeBefore")

	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeBefore(ptr.Pointer(), PointerFromQSGNode(node), PointerFromQSGNode(before))
	}
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGNode) LastChild() *QSGNode {
	defer qt.Recovering("QSGNode::lastChild")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_LastChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) MarkDirty(bits QSGNode__DirtyStateBit) {
	defer qt.Recovering("QSGNode::markDirty")

	if ptr.Pointer() != nil {
		C.QSGNode_MarkDirty(ptr.Pointer(), C.int(bits))
	}
}

func (ptr *QSGNode) NextSibling() *QSGNode {
	defer qt.Recovering("QSGNode::nextSibling")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_NextSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Parent() *QSGNode {
	defer qt.Recovering("QSGNode::parent")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) PrependChildNode(node QSGNode_ITF) {
	defer qt.Recovering("QSGNode::prependChildNode")

	if ptr.Pointer() != nil {
		C.QSGNode_PrependChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGNodePreprocess
func callbackQSGNodePreprocess(ptrName *C.char) bool {
	defer qt.Recovering("callback QSGNode::preprocess")

	var signal = qt.GetSignal(C.GoString(ptrName), "preprocess")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QSGNode) PreviousSibling() *QSGNode {
	defer qt.Recovering("QSGNode::previousSibling")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_PreviousSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) RemoveAllChildNodes() {
	defer qt.Recovering("QSGNode::removeAllChildNodes")

	if ptr.Pointer() != nil {
		C.QSGNode_RemoveAllChildNodes(ptr.Pointer())
	}
}

func (ptr *QSGNode) RemoveChildNode(node QSGNode_ITF) {
	defer qt.Recovering("QSGNode::removeChildNode")

	if ptr.Pointer() != nil {
		C.QSGNode_RemoveChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) SetFlag(f QSGNode__Flag, enabled bool) {
	defer qt.Recovering("QSGNode::setFlag")

	if ptr.Pointer() != nil {
		C.QSGNode_SetFlag(ptr.Pointer(), C.int(f), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSGNode) SetFlags(f QSGNode__Flag, enabled bool) {
	defer qt.Recovering("QSGNode::setFlags")

	if ptr.Pointer() != nil {
		C.QSGNode_SetFlags(ptr.Pointer(), C.int(f), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSGNode) Type() QSGNode__NodeType {
	defer qt.Recovering("QSGNode::type")

	if ptr.Pointer() != nil {
		return QSGNode__NodeType(C.QSGNode_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGNode) DestroyQSGNode() {
	defer qt.Recovering("QSGNode::~QSGNode")

	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNode(ptr.Pointer())
	}
}

func (ptr *QSGNode) ObjectNameAbs() string {
	defer qt.Recovering("QSGNode::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGNode_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGNode) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGNode::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGNode_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
