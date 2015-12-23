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
func callbackQTextBlockGroupTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBlockGroup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTextBlockGroupChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBlockGroup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTextBlockGroupCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBlockGroup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
