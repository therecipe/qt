package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHBoxLayout struct {
	QBoxLayout
}

type QHBoxLayout_ITF interface {
	QBoxLayout_ITF
	QHBoxLayout_PTR() *QHBoxLayout
}

func PointerFromQHBoxLayout(ptr QHBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQHBoxLayoutFromPointer(ptr unsafe.Pointer) *QHBoxLayout {
	var n = new(QHBoxLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHBoxLayout_") {
		n.SetObjectName("QHBoxLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout {
	return ptr
}

func NewQHBoxLayout() *QHBoxLayout {
	defer qt.Recovering("QHBoxLayout::QHBoxLayout")

	return NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout())
}

func NewQHBoxLayout2(parent QWidget_ITF) *QHBoxLayout {
	defer qt.Recovering("QHBoxLayout::QHBoxLayout")

	return NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout2(PointerFromQWidget(parent)))
}

func (ptr *QHBoxLayout) DestroyQHBoxLayout() {
	defer qt.Recovering("QHBoxLayout::~QHBoxLayout")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_DestroyQHBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHBoxLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	defer qt.Recovering("connect QHBoxLayout::addItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "addItem", f)
	}
}

func (ptr *QHBoxLayout) DisconnectAddItem() {
	defer qt.Recovering("disconnect QHBoxLayout::addItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "addItem")
	}
}

//export callbackQHBoxLayoutAddItem
func callbackQHBoxLayoutAddItem(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QHBoxLayout::addItem")

	if signal := qt.GetSignal(C.GoString(ptrName), "addItem"); signal != nil {
		signal.(func(*QLayoutItem))(NewQLayoutItemFromPointer(item))
	} else {
		NewQHBoxLayoutFromPointer(ptr).AddItemDefault(NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QHBoxLayout) AddItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QHBoxLayout::addItem")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QHBoxLayout) AddItemDefault(item QLayoutItem_ITF) {
	defer qt.Recovering("QHBoxLayout::addItem")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_AddItemDefault(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QHBoxLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QHBoxLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QHBoxLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QHBoxLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQHBoxLayoutInvalidate
func callbackQHBoxLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHBoxLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQHBoxLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QHBoxLayout) Invalidate() {
	defer qt.Recovering("QHBoxLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QHBoxLayout) InvalidateDefault() {
	defer qt.Recovering("QHBoxLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QHBoxLayout) ConnectSetGeometry(f func(r *core.QRect)) {
	defer qt.Recovering("connect QHBoxLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setGeometry", f)
	}
}

func (ptr *QHBoxLayout) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QHBoxLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setGeometry")
	}
}

//export callbackQHBoxLayoutSetGeometry
func callbackQHBoxLayoutSetGeometry(ptr unsafe.Pointer, ptrName *C.char, r unsafe.Pointer) {
	defer qt.Recovering("callback QHBoxLayout::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(r))
	} else {
		NewQHBoxLayoutFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QHBoxLayout) SetGeometry(r core.QRect_ITF) {
	defer qt.Recovering("QHBoxLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QHBoxLayout) SetGeometryDefault(r core.QRect_ITF) {
	defer qt.Recovering("QHBoxLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QHBoxLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	defer qt.Recovering("connect QHBoxLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHBoxLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHBoxLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHBoxLayoutChildEvent
func callbackQHBoxLayoutChildEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHBoxLayout::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(e))
	} else {
		NewQHBoxLayoutFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QHBoxLayout) ChildEvent(e core.QChildEvent_ITF) {
	defer qt.Recovering("QHBoxLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QHBoxLayout) ChildEventDefault(e core.QChildEvent_ITF) {
	defer qt.Recovering("QHBoxLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QHBoxLayout) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHBoxLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHBoxLayout) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHBoxLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHBoxLayoutTimerEvent
func callbackQHBoxLayoutTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHBoxLayout::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHBoxLayoutFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHBoxLayout) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHBoxLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHBoxLayout) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHBoxLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHBoxLayout) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHBoxLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHBoxLayout) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHBoxLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHBoxLayoutCustomEvent
func callbackQHBoxLayoutCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHBoxLayout::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHBoxLayoutFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHBoxLayout) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHBoxLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHBoxLayout) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHBoxLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QHBoxLayout_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
