package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsRectItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsRectItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsRectItem_PTR() *QGraphicsRectItem
}

func PointerFromQGraphicsRectItem(ptr QGraphicsRectItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsRectItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsRectItemFromPointer(ptr unsafe.Pointer) *QGraphicsRectItem {
	var n = new(QGraphicsRectItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGraphicsRectItem_") {
		n.SetObjectNameAbs("QGraphicsRectItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsRectItem) QGraphicsRectItem_PTR() *QGraphicsRectItem {
	return ptr
}

func (ptr *QGraphicsRectItem) SetRect(rectangle core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsRectItem::setRect")

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_SetRect(ptr.Pointer(), core.PointerFromQRectF(rectangle))
	}
}

func (ptr *QGraphicsRectItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsRectItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsRectItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsRectItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsRectItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsRectItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsRectItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsRectItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsRectItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsRectItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsRectItemPaint
func callbackQGraphicsRectItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsRectItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsRectItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsRectItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsRectItem) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsRectItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsRectItem) SetRect2(x float64, y float64, width float64, height float64) {
	defer qt.Recovering("QGraphicsRectItem::setRect")

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QGraphicsRectItem) Type() int {
	defer qt.Recovering("QGraphicsRectItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsRectItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsRectItem) DestroyQGraphicsRectItem() {
	defer qt.Recovering("QGraphicsRectItem::~QGraphicsRectItem")

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_DestroyQGraphicsRectItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsRectItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsRectItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsRectItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsRectItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsRectItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
