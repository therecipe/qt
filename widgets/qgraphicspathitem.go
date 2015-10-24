package widgets

//#include "qgraphicspathitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsPathItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsPathItemITF interface {
	QAbstractGraphicsShapeItemITF
	QGraphicsPathItemPTR() *QGraphicsPathItem
}

func PointerFromQGraphicsPathItem(ptr QGraphicsPathItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsPathItemPTR().Pointer()
	}
	return nil
}

func QGraphicsPathItemFromPointer(ptr unsafe.Pointer) *QGraphicsPathItem {
	var n = new(QGraphicsPathItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsPathItem) QGraphicsPathItemPTR() *QGraphicsPathItem {
	return ptr
}

func (ptr *QGraphicsPathItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPathItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsPathItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPathItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsPathItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsPathItem) SetPath(path gui.QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_SetPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainterPath(path)))
	}
}

func (ptr *QGraphicsPathItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsPathItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsPathItem) DestroyQGraphicsPathItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_DestroyQGraphicsPathItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
