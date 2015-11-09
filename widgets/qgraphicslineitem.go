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

type QGraphicsLineItem_ITF interface {
	QGraphicsItem_ITF
	QGraphicsLineItem_PTR() *QGraphicsLineItem
}

func PointerFromQGraphicsLineItem(ptr QGraphicsLineItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLineItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLineItemFromPointer(ptr unsafe.Pointer) *QGraphicsLineItem {
	var n = new(QGraphicsLineItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLineItem) QGraphicsLineItem_PTR() *QGraphicsLineItem {
	return ptr
}

func NewQGraphicsLineItem(parent QGraphicsItem_ITF) *QGraphicsLineItem {
	return NewQGraphicsLineItemFromPointer(C.QGraphicsLineItem_NewQGraphicsLineItem(PointerFromQGraphicsItem(parent)))
}

func NewQGraphicsLineItem2(line core.QLineF_ITF, parent QGraphicsItem_ITF) *QGraphicsLineItem {
	return NewQGraphicsLineItemFromPointer(C.QGraphicsLineItem_NewQGraphicsLineItem2(core.PointerFromQLineF(line), PointerFromQGraphicsItem(parent)))
}

func NewQGraphicsLineItem3(x1 float64, y1 float64, x2 float64, y2 float64, parent QGraphicsItem_ITF) *QGraphicsLineItem {
	return NewQGraphicsLineItemFromPointer(C.QGraphicsLineItem_NewQGraphicsLineItem3(C.double(x1), C.double(y1), C.double(x2), C.double(y2), PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsLineItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsLineItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsLineItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsLineItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsLineItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsLineItem) SetLine(line core.QLineF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetLine(ptr.Pointer(), core.PointerFromQLineF(line))
	}
}

func (ptr *QGraphicsLineItem) SetLine2(x1 float64, y1 float64, x2 float64, y2 float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetLine2(ptr.Pointer(), C.double(x1), C.double(y1), C.double(x2), C.double(y2))
	}
}

func (ptr *QGraphicsLineItem) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QGraphicsLineItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsLineItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLineItem) DestroyQGraphicsLineItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_DestroyQGraphicsLineItem(ptr.Pointer())
	}
}
