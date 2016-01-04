package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsPathItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsPathItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsPathItem_PTR() *QGraphicsPathItem
}

func PointerFromQGraphicsPathItem(ptr QGraphicsPathItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsPathItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsPathItemFromPointer(ptr unsafe.Pointer) *QGraphicsPathItem {
	var n = new(QGraphicsPathItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGraphicsPathItem_") {
		n.SetObjectNameAbs("QGraphicsPathItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsPathItem) QGraphicsPathItem_PTR() *QGraphicsPathItem {
	return ptr
}

func (ptr *QGraphicsPathItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsPathItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsPathItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsPathItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsPathItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsPathItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsPathItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsPathItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsPathItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsPathItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsPathItemPaint
func callbackQGraphicsPathItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsPathItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsPathItemFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsPathItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsPathItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsPathItem) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsPathItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsPathItem) SetPath(path gui.QPainterPath_ITF) {
	defer qt.Recovering("QGraphicsPathItem::setPath")

	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_SetPath(ptr.Pointer(), gui.PointerFromQPainterPath(path))
	}
}

func (ptr *QGraphicsPathItem) Type() int {
	defer qt.Recovering("QGraphicsPathItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsPathItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPathItem) DestroyQGraphicsPathItem() {
	defer qt.Recovering("QGraphicsPathItem::~QGraphicsPathItem")

	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_DestroyQGraphicsPathItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsPathItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsPathItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsPathItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsPathItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsPathItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsPathItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
