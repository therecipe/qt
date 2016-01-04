package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextObject struct {
	core.QObject
}

type QTextObject_ITF interface {
	core.QObject_ITF
	QTextObject_PTR() *QTextObject
}

func PointerFromQTextObject(ptr QTextObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObject_PTR().Pointer()
	}
	return nil
}

func NewQTextObjectFromPointer(ptr unsafe.Pointer) *QTextObject {
	var n = new(QTextObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextObject_") {
		n.SetObjectName("QTextObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextObject) QTextObject_PTR() *QTextObject {
	return ptr
}

func (ptr *QTextObject) Document() *QTextDocument {
	defer qt.Recovering("QTextObject::document")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QTextObject_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextObject) FormatIndex() int {
	defer qt.Recovering("QTextObject::formatIndex")

	if ptr.Pointer() != nil {
		return int(C.QTextObject_FormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextObject) ObjectIndex() int {
	defer qt.Recovering("QTextObject::objectIndex")

	if ptr.Pointer() != nil {
		return int(C.QTextObject_ObjectIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextObjectTimerEvent
func callbackQTextObjectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextObject) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextObjectChildEvent
func callbackQTextObjectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextObject) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextObject::childEvent")

	if ptr.Pointer() != nil {
		C.QTextObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextObject) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextObject::childEvent")

	if ptr.Pointer() != nil {
		C.QTextObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextObjectCustomEvent
func callbackQTextObjectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextObject) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextObject::customEvent")

	if ptr.Pointer() != nil {
		C.QTextObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextObject) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextObject::customEvent")

	if ptr.Pointer() != nil {
		C.QTextObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
