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

type QGraphicsPolygonItemITF interface {
	QAbstractGraphicsShapeItemITF
	QGraphicsPolygonItemPTR() *QGraphicsPolygonItem
}

func PointerFromQGraphicsPolygonItem(ptr QGraphicsPolygonItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsPolygonItemPTR().Pointer()
	}
	return nil
}

func QGraphicsPolygonItemFromPointer(ptr unsafe.Pointer) *QGraphicsPolygonItem {
	var n = new(QGraphicsPolygonItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsPolygonItem) QGraphicsPolygonItemPTR() *QGraphicsPolygonItem {
	return ptr
}

func (ptr *QGraphicsPolygonItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPolygonItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsPolygonItem) FillRule() core.Qt__FillRule {
	if ptr.Pointer() != nil {
		return core.Qt__FillRule(C.QGraphicsPolygonItem_FillRule(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsPolygonItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPolygonItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsPolygonItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsPolygonItem) SetFillRule(rule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetFillRule(C.QtObjectPtr(ptr.Pointer()), C.int(rule))
	}
}

func (ptr *QGraphicsPolygonItem) SetPolygon(polygon gui.QPolygonFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetPolygon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPolygonF(polygon)))
	}
}

func (ptr *QGraphicsPolygonItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsPolygonItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsPolygonItem) DestroyQGraphicsPolygonItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
