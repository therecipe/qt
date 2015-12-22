package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QGraphicsEllipseItem_") {
		n.SetObjectNameAbs("QGraphicsEllipseItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsEllipseItem) QGraphicsEllipseItem_PTR() *QGraphicsEllipseItem {
	return ptr
}

func (ptr *QGraphicsEllipseItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsEllipseItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsEllipseItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsEllipseItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsEllipseItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsEllipseItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsEllipseItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsEllipseItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsEllipseItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsEllipseItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsEllipseItemPaint
func callbackQGraphicsEllipseItemPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsEllipseItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsEllipseItem) SetRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsEllipseItem::setRect")

	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsEllipseItem) SetRect2(x float64, y float64, width float64, height float64) {
	defer qt.Recovering("QGraphicsEllipseItem::setRect")

	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	defer qt.Recovering("QGraphicsEllipseItem::setSpanAngle")

	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetSpanAngle(ptr.Pointer(), C.int(angle))
	}
}

func (ptr *QGraphicsEllipseItem) SetStartAngle(angle int) {
	defer qt.Recovering("QGraphicsEllipseItem::setStartAngle")

	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetStartAngle(ptr.Pointer(), C.int(angle))
	}
}

func (ptr *QGraphicsEllipseItem) SpanAngle() int {
	defer qt.Recovering("QGraphicsEllipseItem::spanAngle")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_SpanAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) StartAngle() int {
	defer qt.Recovering("QGraphicsEllipseItem::startAngle")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_StartAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) Type() int {
	defer qt.Recovering("QGraphicsEllipseItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsEllipseItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsEllipseItem) DestroyQGraphicsEllipseItem() {
	defer qt.Recovering("QGraphicsEllipseItem::~QGraphicsEllipseItem")

	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsEllipseItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsEllipseItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsEllipseItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsEllipseItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsEllipseItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsEllipseItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
