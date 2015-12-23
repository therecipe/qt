package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDoubleValidator struct {
	QValidator
}

type QDoubleValidator_ITF interface {
	QValidator_ITF
	QDoubleValidator_PTR() *QDoubleValidator
}

func PointerFromQDoubleValidator(ptr QDoubleValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDoubleValidator_PTR().Pointer()
	}
	return nil
}

func NewQDoubleValidatorFromPointer(ptr unsafe.Pointer) *QDoubleValidator {
	var n = new(QDoubleValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDoubleValidator_") {
		n.SetObjectName("QDoubleValidator_" + qt.Identifier())
	}
	return n
}

func (ptr *QDoubleValidator) QDoubleValidator_PTR() *QDoubleValidator {
	return ptr
}

//QDoubleValidator::Notation
type QDoubleValidator__Notation int64

const (
	QDoubleValidator__StandardNotation   = QDoubleValidator__Notation(0)
	QDoubleValidator__ScientificNotation = QDoubleValidator__Notation(1)
)

func (ptr *QDoubleValidator) Notation() QDoubleValidator__Notation {
	defer qt.Recovering("QDoubleValidator::notation")

	if ptr.Pointer() != nil {
		return QDoubleValidator__Notation(C.QDoubleValidator_Notation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleValidator) SetDecimals(v int) {
	defer qt.Recovering("QDoubleValidator::setDecimals")

	if ptr.Pointer() != nil {
		C.QDoubleValidator_SetDecimals(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QDoubleValidator) SetNotation(v QDoubleValidator__Notation) {
	defer qt.Recovering("QDoubleValidator::setNotation")

	if ptr.Pointer() != nil {
		C.QDoubleValidator_SetNotation(ptr.Pointer(), C.int(v))
	}
}

func NewQDoubleValidator(parent core.QObject_ITF) *QDoubleValidator {
	defer qt.Recovering("QDoubleValidator::QDoubleValidator")

	return NewQDoubleValidatorFromPointer(C.QDoubleValidator_NewQDoubleValidator(core.PointerFromQObject(parent)))
}

func (ptr *QDoubleValidator) Decimals() int {
	defer qt.Recovering("QDoubleValidator::decimals")

	if ptr.Pointer() != nil {
		return int(C.QDoubleValidator_Decimals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleValidator) DestroyQDoubleValidator() {
	defer qt.Recovering("QDoubleValidator::~QDoubleValidator")

	if ptr.Pointer() != nil {
		C.QDoubleValidator_DestroyQDoubleValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDoubleValidator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDoubleValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDoubleValidator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDoubleValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDoubleValidatorTimerEvent
func callbackQDoubleValidatorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleValidator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleValidator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDoubleValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDoubleValidator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDoubleValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDoubleValidatorChildEvent
func callbackQDoubleValidatorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleValidator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleValidator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDoubleValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDoubleValidator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDoubleValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDoubleValidatorCustomEvent
func callbackQDoubleValidatorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleValidator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
