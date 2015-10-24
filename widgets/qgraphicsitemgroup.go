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

type QGraphicsItemGroupITF interface {
	QGraphicsItemITF
	QGraphicsItemGroupPTR() *QGraphicsItemGroup
}

func PointerFromQGraphicsItemGroup(ptr QGraphicsItemGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItemGroupPTR().Pointer()
	}
	return nil
}

func QGraphicsItemGroupFromPointer(ptr unsafe.Pointer) *QGraphicsItemGroup {
	var n = new(QGraphicsItemGroup)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroupPTR() *QGraphicsItemGroup {
	return ptr
}

func NewQGraphicsItemGroup(parent QGraphicsItemITF) *QGraphicsItemGroup {
	return QGraphicsItemGroupFromPointer(unsafe.Pointer(C.QGraphicsItemGroup_NewQGraphicsItemGroup(C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func (ptr *QGraphicsItemGroup) AddToGroup(item QGraphicsItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_AddToGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)))
	}
}

func (ptr *QGraphicsItemGroup) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsItemGroup_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsItemGroup) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsItemGroup) RemoveFromGroup(item QGraphicsItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_RemoveFromGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)))
	}
}

func (ptr *QGraphicsItemGroup) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsItemGroup_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroup() {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroup(C.QtObjectPtr(ptr.Pointer()))
	}
}
