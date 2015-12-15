package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QGraphicsLineItem_") {
		n.SetObjectNameAbs("QGraphicsLineItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsLineItem) QGraphicsLineItem_PTR() *QGraphicsLineItem {
	return ptr
}

func NewQGraphicsLineItem(parent QGraphicsItem_ITF) *QGraphicsLineItem {
	defer qt.Recovering("QGraphicsLineItem::QGraphicsLineItem")

	return NewQGraphicsLineItemFromPointer(C.QGraphicsLineItem_NewQGraphicsLineItem(PointerFromQGraphicsItem(parent)))
}

func NewQGraphicsLineItem2(line core.QLineF_ITF, parent QGraphicsItem_ITF) *QGraphicsLineItem {
	defer qt.Recovering("QGraphicsLineItem::QGraphicsLineItem")

	return NewQGraphicsLineItemFromPointer(C.QGraphicsLineItem_NewQGraphicsLineItem2(core.PointerFromQLineF(line), PointerFromQGraphicsItem(parent)))
}

func NewQGraphicsLineItem3(x1 float64, y1 float64, x2 float64, y2 float64, parent QGraphicsItem_ITF) *QGraphicsLineItem {
	defer qt.Recovering("QGraphicsLineItem::QGraphicsLineItem")

	return NewQGraphicsLineItemFromPointer(C.QGraphicsLineItem_NewQGraphicsLineItem3(C.double(x1), C.double(y1), C.double(x2), C.double(y2), PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsLineItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsLineItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsLineItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsLineItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsLineItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsLineItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsLineItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsLineItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsLineItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsLineItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsLineItemPaint
func callbackQGraphicsLineItemPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsLineItem::paint")

	var signal = qt.GetSignal(C.GoString(ptrName), "paint")
	if signal != nil {
		defer signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsLineItem) SetLine(line core.QLineF_ITF) {
	defer qt.Recovering("QGraphicsLineItem::setLine")

	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetLine(ptr.Pointer(), core.PointerFromQLineF(line))
	}
}

func (ptr *QGraphicsLineItem) SetLine2(x1 float64, y1 float64, x2 float64, y2 float64) {
	defer qt.Recovering("QGraphicsLineItem::setLine")

	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetLine2(ptr.Pointer(), C.double(x1), C.double(y1), C.double(x2), C.double(y2))
	}
}

func (ptr *QGraphicsLineItem) SetPen(pen gui.QPen_ITF) {
	defer qt.Recovering("QGraphicsLineItem::setPen")

	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QGraphicsLineItem) Type() int {
	defer qt.Recovering("QGraphicsLineItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsLineItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLineItem) DestroyQGraphicsLineItem() {
	defer qt.Recovering("QGraphicsLineItem::~QGraphicsLineItem")

	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_DestroyQGraphicsLineItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsLineItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsLineItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsLineItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsLineItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsLineItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsLineItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
