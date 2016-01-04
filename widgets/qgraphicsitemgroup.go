package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsItemGroup struct {
	QGraphicsItem
}

type QGraphicsItemGroup_ITF interface {
	QGraphicsItem_ITF
	QGraphicsItemGroup_PTR() *QGraphicsItemGroup
}

func PointerFromQGraphicsItemGroup(ptr QGraphicsItemGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItemGroup_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemGroupFromPointer(ptr unsafe.Pointer) *QGraphicsItemGroup {
	var n = new(QGraphicsItemGroup)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGraphicsItemGroup_") {
		n.SetObjectNameAbs("QGraphicsItemGroup_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroup_PTR() *QGraphicsItemGroup {
	return ptr
}

func NewQGraphicsItemGroup(parent QGraphicsItem_ITF) *QGraphicsItemGroup {
	defer qt.Recovering("QGraphicsItemGroup::QGraphicsItemGroup")

	return NewQGraphicsItemGroupFromPointer(C.QGraphicsItemGroup_NewQGraphicsItemGroup(PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsItemGroup) AddToGroup(item QGraphicsItem_ITF) {
	defer qt.Recovering("QGraphicsItemGroup::addToGroup")

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_AddToGroup(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsItemGroup) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsItemGroup::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsItemGroup_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsItemGroup) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsItemGroup::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsItemGroup) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsItemGroup::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsItemGroupPaint
func callbackQGraphicsItemGroupPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsItemGroup::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsItemGroupFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsItemGroup) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsItemGroup::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItemGroup) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsItemGroup::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItemGroup) RemoveFromGroup(item QGraphicsItem_ITF) {
	defer qt.Recovering("QGraphicsItemGroup::removeFromGroup")

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_RemoveFromGroup(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsItemGroup) Type() int {
	defer qt.Recovering("QGraphicsItemGroup::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsItemGroup_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroup() {
	defer qt.Recovering("QGraphicsItemGroup::~QGraphicsItemGroup")

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroup(ptr.Pointer())
	}
}

func (ptr *QGraphicsItemGroup) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsItemGroup::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsItemGroup_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsItemGroup) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsItemGroup::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
