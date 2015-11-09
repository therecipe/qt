package widgets

//#include "qgraphicsitemgroup.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsItemGroup struct {
	QGraphicsItem
}

type QGraphicsItemGroup_ITF interface {
	QGraphicsItem_ITF
	QGraphicsItemGroup_PTR() *QGraphicsItemGroup
}

func PointerFromQGraphicsItemGroup(ptr QGraphicsItemGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItemGroup_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemGroupFromPointer(ptr unsafe.Pointer) *QGraphicsItemGroup {
	var n = new(QGraphicsItemGroup)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroup_PTR() *QGraphicsItemGroup {
	return ptr
}

func NewQGraphicsItemGroup(parent QGraphicsItem_ITF) *QGraphicsItemGroup {
	return NewQGraphicsItemGroupFromPointer(C.QGraphicsItemGroup_NewQGraphicsItemGroup(PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsItemGroup) AddToGroup(item QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_AddToGroup(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsItemGroup) IsObscuredBy(item QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsItemGroup_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsItemGroup) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItemGroup) RemoveFromGroup(item QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_RemoveFromGroup(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsItemGroup) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsItemGroup_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroup() {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroup(ptr.Pointer())
	}
}
