package widgets

//#include "qgraphicslineitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsLineItem struct {
	QGraphicsItem
}

type QGraphicsLineItemITF interface {
	QGraphicsItemITF
	QGraphicsLineItemPTR() *QGraphicsLineItem
}

func PointerFromQGraphicsLineItem(ptr QGraphicsLineItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLineItemPTR().Pointer()
	}
	return nil
}

func QGraphicsLineItemFromPointer(ptr unsafe.Pointer) *QGraphicsLineItem {
	var n = new(QGraphicsLineItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLineItem) QGraphicsLineItemPTR() *QGraphicsLineItem {
	return ptr
}

func NewQGraphicsLineItem(parent QGraphicsItemITF) *QGraphicsLineItem {
	return QGraphicsLineItemFromPointer(unsafe.Pointer(C.QGraphicsLineItem_NewQGraphicsLineItem(C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func NewQGraphicsLineItem2(line core.QLineFITF, parent QGraphicsItemITF) *QGraphicsLineItem {
	return QGraphicsLineItemFromPointer(unsafe.Pointer(C.QGraphicsLineItem_NewQGraphicsLineItem2(C.QtObjectPtr(core.PointerFromQLineF(line)), C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func (ptr *QGraphicsLineItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsLineItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsLineItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsLineItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsLineItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsLineItem) SetLine(line core.QLineFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetLine(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLineF(line)))
	}
}

func (ptr *QGraphicsLineItem) SetPen(pen gui.QPenITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetPen(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPen(pen)))
	}
}

func (ptr *QGraphicsLineItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsLineItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsLineItem) DestroyQGraphicsLineItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_DestroyQGraphicsLineItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
