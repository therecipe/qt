package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextBlockGroup struct {
	QTextObject
}

type QTextBlockGroup_ITF interface {
	QTextObject_ITF
	QTextBlockGroup_PTR() *QTextBlockGroup
}

func PointerFromQTextBlockGroup(ptr QTextBlockGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockGroup_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockGroupFromPointer(ptr unsafe.Pointer) *QTextBlockGroup {
	var n = new(QTextBlockGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextBlockGroup_") {
		n.SetObjectName("QTextBlockGroup_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextBlockGroup) QTextBlockGroup_PTR() *QTextBlockGroup {
	return ptr
}

func (ptr *QTextBlockGroup) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextBlockGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextBlockGroup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextBlockGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextBlockGroupTimerEvent
func callbackQTextBlockGroupTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBlockGroup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextBlockGroupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextBlockGroup) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextBlockGroup::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextBlockGroup_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextBlockGroup) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextBlockGroup::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextBlockGroup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextBlockGroup) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextBlockGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextBlockGroup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextBlockGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextBlockGroupChildEvent
func callbackQTextBlockGroupChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBlockGroup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextBlockGroupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextBlockGroup) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextBlockGroup::childEvent")

	if ptr.Pointer() != nil {
		C.QTextBlockGroup_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextBlockGroup) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextBlockGroup::childEvent")

	if ptr.Pointer() != nil {
		C.QTextBlockGroup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextBlockGroup) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextBlockGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextBlockGroup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextBlockGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextBlockGroupCustomEvent
func callbackQTextBlockGroupCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBlockGroup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextBlockGroupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextBlockGroup) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBlockGroup::customEvent")

	if ptr.Pointer() != nil {
		C.QTextBlockGroup_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextBlockGroup) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBlockGroup::customEvent")

	if ptr.Pointer() != nil {
		C.QTextBlockGroup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
