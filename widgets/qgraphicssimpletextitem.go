package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsSimpleTextItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsSimpleTextItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsSimpleTextItem_PTR() *QGraphicsSimpleTextItem
}

func PointerFromQGraphicsSimpleTextItem(ptr QGraphicsSimpleTextItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSimpleTextItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSimpleTextItemFromPointer(ptr unsafe.Pointer) *QGraphicsSimpleTextItem {
	var n = new(QGraphicsSimpleTextItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGraphicsSimpleTextItem_") {
		n.SetObjectNameAbs("QGraphicsSimpleTextItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsSimpleTextItem) QGraphicsSimpleTextItem_PTR() *QGraphicsSimpleTextItem {
	return ptr
}

func (ptr *QGraphicsSimpleTextItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsSimpleTextItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsSimpleTextItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsSimpleTextItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsSimpleTextItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsSimpleTextItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsSimpleTextItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsSimpleTextItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsSimpleTextItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsSimpleTextItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsSimpleTextItemPaint
func callbackQGraphicsSimpleTextItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSimpleTextItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsSimpleTextItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsSimpleTextItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSimpleTextItem) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsSimpleTextItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSimpleTextItem) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QGraphicsSimpleTextItem::setFont")

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsSimpleTextItem) SetText(text string) {
	defer qt.Recovering("QGraphicsSimpleTextItem::setText")

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGraphicsSimpleTextItem) Text() string {
	defer qt.Recovering("QGraphicsSimpleTextItem::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSimpleTextItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSimpleTextItem) Type() int {
	defer qt.Recovering("QGraphicsSimpleTextItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsSimpleTextItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSimpleTextItem) DestroyQGraphicsSimpleTextItem() {
	defer qt.Recovering("QGraphicsSimpleTextItem::~QGraphicsSimpleTextItem")

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsSimpleTextItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsSimpleTextItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSimpleTextItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSimpleTextItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsSimpleTextItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
