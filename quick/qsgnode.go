package quick

//#include "qsgnode.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSGNode struct {
	ptr unsafe.Pointer
}

type QSGNodeITF interface {
	QSGNodePTR() *QSGNode
}

func (p *QSGNode) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGNode) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGNode(ptr QSGNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNodePTR().Pointer()
	}
	return nil
}

func QSGNodeFromPointer(ptr unsafe.Pointer) *QSGNode {
	var n = new(QSGNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGNode) QSGNodePTR() *QSGNode {
	return ptr
}

//QSGNode::DirtyStateBit
type QSGNode__DirtyStateBit int

var (
	QSGNode__DirtySubtreeBlocked = QSGNode__DirtyStateBit(0x0080)
	QSGNode__DirtyMatrix         = QSGNode__DirtyStateBit(0x0100)
	QSGNode__DirtyNodeAdded      = QSGNode__DirtyStateBit(0x0400)
	QSGNode__DirtyNodeRemoved    = QSGNode__DirtyStateBit(0x0800)
	QSGNode__DirtyGeometry       = QSGNode__DirtyStateBit(0x1000)
	QSGNode__DirtyMaterial       = QSGNode__DirtyStateBit(0x2000)
	QSGNode__DirtyOpacity        = QSGNode__DirtyStateBit(0x4000)
)

//QSGNode::Flag
type QSGNode__Flag int

var (
	QSGNode__OwnedByParent      = QSGNode__Flag(0x0001)
	QSGNode__UsePreprocess      = QSGNode__Flag(0x0002)
	QSGNode__OwnsGeometry       = QSGNode__Flag(0x00010000)
	QSGNode__OwnsMaterial       = QSGNode__Flag(0x00020000)
	QSGNode__OwnsOpaqueMaterial = QSGNode__Flag(0x00040000)
	QSGNode__InternalReserved   = QSGNode__Flag(0x01000000)
)

//QSGNode::NodeType
type QSGNode__NodeType int

var (
	QSGNode__BasicNodeType     = QSGNode__NodeType(0)
	QSGNode__GeometryNodeType  = QSGNode__NodeType(1)
	QSGNode__TransformNodeType = QSGNode__NodeType(2)
	QSGNode__ClipNodeType      = QSGNode__NodeType(3)
	QSGNode__OpacityNodeType   = QSGNode__NodeType(4)
)

func (ptr *QSGNode) ChildAtIndex(i int) *QSGNode {
	if ptr.Pointer() != nil {
		return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_ChildAtIndex(C.QtObjectPtr(ptr.Pointer()), C.int(i))))
	}
	return nil
}

func (ptr *QSGNode) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGNode_ChildCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQSGNode() *QSGNode {
	return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_NewQSGNode()))
}

func (ptr *QSGNode) AppendChildNode(node QSGNodeITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_AppendChildNode(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGNode(node)))
	}
}

func (ptr *QSGNode) FirstChild() *QSGNode {
	if ptr.Pointer() != nil {
		return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_FirstChild(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGNode) Flags() QSGNode__Flag {
	if ptr.Pointer() != nil {
		return QSGNode__Flag(C.QSGNode_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGNode) InsertChildNodeAfter(node QSGNodeITF, after QSGNodeITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeAfter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGNode(node)), C.QtObjectPtr(PointerFromQSGNode(after)))
	}
}

func (ptr *QSGNode) InsertChildNodeBefore(node QSGNodeITF, before QSGNodeITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeBefore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGNode(node)), C.QtObjectPtr(PointerFromQSGNode(before)))
	}
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlocked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGNode) LastChild() *QSGNode {
	if ptr.Pointer() != nil {
		return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_LastChild(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGNode) MarkDirty(bits QSGNode__DirtyStateBit) {
	if ptr.Pointer() != nil {
		C.QSGNode_MarkDirty(C.QtObjectPtr(ptr.Pointer()), C.int(bits))
	}
}

func (ptr *QSGNode) NextSibling() *QSGNode {
	if ptr.Pointer() != nil {
		return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_NextSibling(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGNode) Parent() *QSGNode {
	if ptr.Pointer() != nil {
		return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGNode) PrependChildNode(node QSGNodeITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_PrependChildNode(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGNode(node)))
	}
}

func (ptr *QSGNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGNode_Preprocess(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGNode) PreviousSibling() *QSGNode {
	if ptr.Pointer() != nil {
		return QSGNodeFromPointer(unsafe.Pointer(C.QSGNode_PreviousSibling(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGNode) RemoveAllChildNodes() {
	if ptr.Pointer() != nil {
		C.QSGNode_RemoveAllChildNodes(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGNode) RemoveChildNode(node QSGNodeITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_RemoveChildNode(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGNode(node)))
	}
}

func (ptr *QSGNode) SetFlag(f QSGNode__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QSGNode_SetFlag(C.QtObjectPtr(ptr.Pointer()), C.int(f), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSGNode) SetFlags(f QSGNode__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QSGNode_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(f), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSGNode) Type() QSGNode__NodeType {
	if ptr.Pointer() != nil {
		return QSGNode__NodeType(C.QSGNode_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGNode) DestroyQSGNode() {
	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
