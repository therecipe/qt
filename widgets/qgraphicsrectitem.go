package widgets

//#include "qgraphicsrectitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsRectItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsRectItemITF interface {
	QAbstractGraphicsShapeItemITF
	QGraphicsRectItemPTR() *QGraphicsRectItem
}

func PointerFromQGraphicsRectItem(ptr QGraphicsRectItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsRectItemPTR().Pointer()
	}
	return nil
}

func QGraphicsRectItemFromPointer(ptr unsafe.Pointer) *QGraphicsRectItem {
	var n = new(QGraphicsRectItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsRectItem) QGraphicsRectItemPTR() *QGraphicsRectItem {
	return ptr
}

func (ptr *QGraphicsRectItem) SetRect(rectangle core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_SetRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)))
	}
}

func (ptr *QGraphicsRectItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsRectItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsRectItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsRectItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsRectItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsRectItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsRectItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsRectItem) DestroyQGraphicsRectItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_DestroyQGraphicsRectItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
