package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QGraphicsPolygonItem_") {
		n.SetObjectNameAbs("QGraphicsPolygonItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsPolygonItem) QGraphicsPolygonItem_PTR() *QGraphicsPolygonItem {
	return ptr
}

func (ptr *QGraphicsPolygonItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsPolygonItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsPolygonItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsPolygonItem) FillRule() core.Qt__FillRule {
	defer qt.Recovering("QGraphicsPolygonItem::fillRule")

	if ptr.Pointer() != nil {
		return core.Qt__FillRule(C.QGraphicsPolygonItem_FillRule(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPolygonItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsPolygonItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsPolygonItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsPolygonItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsPolygonItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsPolygonItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsPolygonItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsPolygonItemPaint
func callbackQGraphicsPolygonItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsPolygonItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsPolygonItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsPolygonItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsPolygonItem) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsPolygonItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsPolygonItem) SetFillRule(rule core.Qt__FillRule) {
	defer qt.Recovering("QGraphicsPolygonItem::setFillRule")

	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetFillRule(ptr.Pointer(), C.int(rule))
	}
}

func (ptr *QGraphicsPolygonItem) SetPolygon(polygon gui.QPolygonF_ITF) {
	defer qt.Recovering("QGraphicsPolygonItem::setPolygon")

	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetPolygon(ptr.Pointer(), gui.PointerFromQPolygonF(polygon))
	}
}

func (ptr *QGraphicsPolygonItem) Type() int {
	defer qt.Recovering("QGraphicsPolygonItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsPolygonItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPolygonItem) DestroyQGraphicsPolygonItem() {
	defer qt.Recovering("QGraphicsPolygonItem::~QGraphicsPolygonItem")

	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsPolygonItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsPolygonItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsPolygonItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsPolygonItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsPolygonItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsPolygonItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
