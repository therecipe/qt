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

type QGraphicsEllipseItemITF interface {
	QAbstractGraphicsShapeItemITF
	QGraphicsEllipseItemPTR() *QGraphicsEllipseItem
}

func PointerFromQGraphicsEllipseItem(ptr QGraphicsEllipseItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsEllipseItemPTR().Pointer()
	}
	return nil
}

func QGraphicsEllipseItemFromPointer(ptr unsafe.Pointer) *QGraphicsEllipseItem {
	var n = new(QGraphicsEllipseItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsEllipseItem) QGraphicsEllipseItemPTR() *QGraphicsEllipseItem {
	return ptr
}

func (ptr *QGraphicsEllipseItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsEllipseItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsEllipseItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsEllipseItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsEllipseItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsEllipseItem) SetRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetSpanAngle(C.QtObjectPtr(ptr.Pointer()), C.int(angle))
	}
}

func (ptr *QGraphicsEllipseItem) SetStartAngle(angle int) {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetStartAngle(C.QtObjectPtr(ptr.Pointer()), C.int(angle))
	}
}

func (ptr *QGraphicsEllipseItem) SpanAngle() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_SpanAngle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) StartAngle() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_StartAngle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) DestroyQGraphicsEllipseItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
