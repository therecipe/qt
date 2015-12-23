package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QValidator struct {
	core.QObject
}

type QValidator_ITF interface {
	core.QObject_ITF
	QValidator_PTR() *QValidator
}

func PointerFromQValidator(ptr QValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValidator_PTR().Pointer()
	}
	return nil
}

func NewQValidatorFromPointer(ptr unsafe.Pointer) *QValidator {
	var n = new(QValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QValidator_") {
		n.SetObjectName("QValidator_" + qt.Identifier())
	}
	return n
}

func (ptr *QValidator) QValidator_PTR() *QValidator {
	return ptr
}

//QValidator::State
type QValidator__State int64

const (
	QValidator__Invalid      = QValidator__State(0)
	QValidator__Intermediate = QValidator__State(1)
	QValidator__Acceptable   = QValidator__State(2)
)

func (ptr *QValidator) ConnectChanged(f func()) {
	defer qt.Recovering("connect QValidator::changed")

	if ptr.Pointer() != nil {
		C.QValidator_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QValidator) DisconnectChanged() {
	defer qt.Recovering("disconnect QValidator::changed")

	if ptr.Pointer() != nil {
		C.QValidator_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQValidatorChanged
func callbackQValidatorChanged(ptrName *C.char) {
	defer qt.Recovering("callback QValidator::changed")

	if signal := qt.GetSignal(C.GoString(ptrName), "changed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QValidator) SetLocale(locale core.QLocale_ITF) {
	defer qt.Recovering("QValidator::setLocale")

	if ptr.Pointer() != nil {
		C.QValidator_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QValidator) DestroyQValidator() {
	defer qt.Recovering("QValidator::~QValidator")

	if ptr.Pointer() != nil {
		C.QValidator_DestroyQValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QValidator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QValidator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQValidatorTimerEvent
func callbackQValidatorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QValidator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QValidator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QValidator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQValidatorChildEvent
func callbackQValidatorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QValidator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QValidator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QValidator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQValidatorCustomEvent
func callbackQValidatorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QValidator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
