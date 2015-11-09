package widgets

//#include "qgraphicsellipseitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsEllipseItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsEllipseItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsEllipseItem_PTR() *QGraphicsEllipseItem
}

func PointerFromQGraphicsEllipseItem(ptr QGraphicsEllipseItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsEllipseItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsEllipseItemFromPointer(ptr unsafe.Pointer) *QGraphicsEllipseItem {
	var n = new(QGraphicsEllipseItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsEllipseItem) QGraphicsEllipseItem_PTR() *QGraphicsEllipseItem {
	return ptr
}

func (ptr *QGraphicsEllipseItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsEllipseItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsEllipseItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsEllipseItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsEllipseItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsEllipseItem) SetRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsEllipseItem) SetRect2(x float64, y float64, width float64, height float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetSpanAngle(ptr.Pointer(), C.int(angle))
	}
}

func (ptr *QGraphicsEllipseItem) SetStartAngle(angle int) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetStartAngle(ptr.Pointer(), C.int(angle))
	}
}

func (ptr *QGraphicsEllipseItem) SpanAngle() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_SpanAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) StartAngle() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_StartAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) DestroyQGraphicsEllipseItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(ptr.Pointer())
	}
}
