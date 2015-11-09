package widgets

//#include "qgraphicspolygonitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsPolygonItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsPolygonItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsPolygonItem_PTR() *QGraphicsPolygonItem
}

func PointerFromQGraphicsPolygonItem(ptr QGraphicsPolygonItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsPolygonItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsPolygonItemFromPointer(ptr unsafe.Pointer) *QGraphicsPolygonItem {
	var n = new(QGraphicsPolygonItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsPolygonItem) QGraphicsPolygonItem_PTR() *QGraphicsPolygonItem {
	return ptr
}

func (ptr *QGraphicsPolygonItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPolygonItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsPolygonItem) FillRule() core.Qt__FillRule {
	if ptr.Pointer() != nil {
		return core.Qt__FillRule(C.QGraphicsPolygonItem_FillRule(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPolygonItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPolygonItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsPolygonItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsPolygonItem) SetFillRule(rule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetFillRule(ptr.Pointer(), C.int(rule))
	}
}

func (ptr *QGraphicsPolygonItem) SetPolygon(polygon gui.QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetPolygon(ptr.Pointer(), gui.PointerFromQPolygonF(polygon))
	}
}

func (ptr *QGraphicsPolygonItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsPolygonItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPolygonItem) DestroyQGraphicsPolygonItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(ptr.Pointer())
	}
}
