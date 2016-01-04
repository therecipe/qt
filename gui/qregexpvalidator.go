package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRegExpValidator struct {
	QValidator
}

type QRegExpValidator_ITF interface {
	QValidator_ITF
	QRegExpValidator_PTR() *QRegExpValidator
}

func PointerFromQRegExpValidator(ptr QRegExpValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegExpValidator_PTR().Pointer()
	}
	return nil
}

func NewQRegExpValidatorFromPointer(ptr unsafe.Pointer) *QRegExpValidator {
	var n = new(QRegExpValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRegExpValidator_") {
		n.SetObjectName("QRegExpValidator_" + qt.Identifier())
	}
	return n
}

func (ptr *QRegExpValidator) QRegExpValidator_PTR() *QRegExpValidator {
	return ptr
}

func (ptr *QRegExpValidator) SetRegExp(rx core.QRegExp_ITF) {
	defer qt.Recovering("QRegExpValidator::setRegExp")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_SetRegExp(ptr.Pointer(), core.PointerFromQRegExp(rx))
	}
}

func NewQRegExpValidator(parent core.QObject_ITF) *QRegExpValidator {
	defer qt.Recovering("QRegExpValidator::QRegExpValidator")

	return NewQRegExpValidatorFromPointer(C.QRegExpValidator_NewQRegExpValidator(core.PointerFromQObject(parent)))
}

func NewQRegExpValidator2(rx core.QRegExp_ITF, parent core.QObject_ITF) *QRegExpValidator {
	defer qt.Recovering("QRegExpValidator::QRegExpValidator")

	return NewQRegExpValidatorFromPointer(C.QRegExpValidator_NewQRegExpValidator2(core.PointerFromQRegExp(rx), core.PointerFromQObject(parent)))
}

func (ptr *QRegExpValidator) RegExp() *core.QRegExp {
	defer qt.Recovering("QRegExpValidator::regExp")

	if ptr.Pointer() != nil {
		return core.NewQRegExpFromPointer(C.QRegExpValidator_RegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRegExpValidator) DestroyQRegExpValidator() {
	defer qt.Recovering("QRegExpValidator::~QRegExpValidator")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_DestroyQRegExpValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRegExpValidator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRegExpValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRegExpValidator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRegExpValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRegExpValidatorTimerEvent
func callbackQRegExpValidatorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRegExpValidator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRegExpValidatorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRegExpValidator) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRegExpValidator::timerEvent")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRegExpValidator) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRegExpValidator::timerEvent")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRegExpValidator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRegExpValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRegExpValidator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRegExpValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRegExpValidatorChildEvent
func callbackQRegExpValidatorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRegExpValidator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRegExpValidatorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRegExpValidator) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRegExpValidator::childEvent")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRegExpValidator) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRegExpValidator::childEvent")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRegExpValidator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRegExpValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRegExpValidator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRegExpValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRegExpValidatorCustomEvent
func callbackQRegExpValidatorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRegExpValidator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRegExpValidatorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRegExpValidator) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRegExpValidator::customEvent")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRegExpValidator) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRegExpValidator::customEvent")

	if ptr.Pointer() != nil {
		C.QRegExpValidator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
