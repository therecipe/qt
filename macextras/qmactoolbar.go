package macextras

//#include "macextras.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacToolBar struct {
	core.QObject
}

type QMacToolBar_ITF interface {
	core.QObject_ITF
	QMacToolBar_PTR() *QMacToolBar
}

func PointerFromQMacToolBar(ptr QMacToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBar_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarFromPointer(ptr unsafe.Pointer) *QMacToolBar {
	var n = new(QMacToolBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacToolBar_") {
		n.SetObjectName("QMacToolBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacToolBar) QMacToolBar_PTR() *QMacToolBar {
	return ptr
}

func NewQMacToolBar(parent core.QObject_ITF) *QMacToolBar {
	defer qt.Recovering("QMacToolBar::QMacToolBar")

	return NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar(core.PointerFromQObject(parent)))
}

func NewQMacToolBar2(identifier string, parent core.QObject_ITF) *QMacToolBar {
	defer qt.Recovering("QMacToolBar::QMacToolBar")

	return NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar2(C.CString(identifier), core.PointerFromQObject(parent)))
}

func (ptr *QMacToolBar) AddAllowedItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBar::addAllowedItem")

	if ptr.Pointer() != nil {
		return NewQMacToolBarItemFromPointer(C.QMacToolBar_AddAllowedItem(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMacToolBar) AddItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBar::addItem")

	if ptr.Pointer() != nil {
		return NewQMacToolBarItemFromPointer(C.QMacToolBar_AddItem(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMacToolBar) AddSeparator() {
	defer qt.Recovering("QMacToolBar::addSeparator")

	if ptr.Pointer() != nil {
		C.QMacToolBar_AddSeparator(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) AttachToWindow(window gui.QWindow_ITF) {
	defer qt.Recovering("QMacToolBar::attachToWindow")

	if ptr.Pointer() != nil {
		C.QMacToolBar_AttachToWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QMacToolBar) DetachFromWindow() {
	defer qt.Recovering("QMacToolBar::detachFromWindow")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DetachFromWindow(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
	defer qt.Recovering("QMacToolBar::~QMacToolBar")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMacToolBarTimerEvent
func callbackQMacToolBarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacToolBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacToolBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMacToolBarChildEvent
func callbackQMacToolBarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBar::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBar::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacToolBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacToolBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMacToolBarCustomEvent
func callbackQMacToolBarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBar::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacToolBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBar::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
