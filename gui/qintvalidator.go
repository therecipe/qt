package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIntValidator struct {
	QValidator
}

type QIntValidator_ITF interface {
	QValidator_ITF
	QIntValidator_PTR() *QIntValidator
}

func PointerFromQIntValidator(ptr QIntValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIntValidator_PTR().Pointer()
	}
	return nil
}

func NewQIntValidatorFromPointer(ptr unsafe.Pointer) *QIntValidator {
	var n = new(QIntValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIntValidator_") {
		n.SetObjectName("QIntValidator_" + qt.Identifier())
	}
	return n
}

func (ptr *QIntValidator) QIntValidator_PTR() *QIntValidator {
	return ptr
}

func (ptr *QIntValidator) SetBottom(v int) {
	defer qt.Recovering("QIntValidator::setBottom")

	if ptr.Pointer() != nil {
		C.QIntValidator_SetBottom(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QIntValidator) SetTop(v int) {
	defer qt.Recovering("QIntValidator::setTop")

	if ptr.Pointer() != nil {
		C.QIntValidator_SetTop(ptr.Pointer(), C.int(v))
	}
}

func NewQIntValidator(parent core.QObject_ITF) *QIntValidator {
	defer qt.Recovering("QIntValidator::QIntValidator")

	return NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator(core.PointerFromQObject(parent)))
}

func NewQIntValidator2(minimum int, maximum int, parent core.QObject_ITF) *QIntValidator {
	defer qt.Recovering("QIntValidator::QIntValidator")

	return NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator2(C.int(minimum), C.int(maximum), core.PointerFromQObject(parent)))
}

func (ptr *QIntValidator) Bottom() int {
	defer qt.Recovering("QIntValidator::bottom")

	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIntValidator) ConnectSetRange(f func(bottom int, top int)) {
	defer qt.Recovering("connect QIntValidator::setRange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRange", f)
	}
}

func (ptr *QIntValidator) DisconnectSetRange() {
	defer qt.Recovering("disconnect QIntValidator::setRange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRange")
	}
}

//export callbackQIntValidatorSetRange
func callbackQIntValidatorSetRange(ptr unsafe.Pointer, ptrName *C.char, bottom C.int, top C.int) {
	defer qt.Recovering("callback QIntValidator::setRange")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRange"); signal != nil {
		signal.(func(int, int))(int(bottom), int(top))
	} else {
		NewQIntValidatorFromPointer(ptr).SetRangeDefault(int(bottom), int(top))
	}
}

func (ptr *QIntValidator) SetRange(bottom int, top int) {
	defer qt.Recovering("QIntValidator::setRange")

	if ptr.Pointer() != nil {
		C.QIntValidator_SetRange(ptr.Pointer(), C.int(bottom), C.int(top))
	}
}

func (ptr *QIntValidator) SetRangeDefault(bottom int, top int) {
	defer qt.Recovering("QIntValidator::setRange")

	if ptr.Pointer() != nil {
		C.QIntValidator_SetRangeDefault(ptr.Pointer(), C.int(bottom), C.int(top))
	}
}

func (ptr *QIntValidator) Top() int {
	defer qt.Recovering("QIntValidator::top")

	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIntValidator) DestroyQIntValidator() {
	defer qt.Recovering("QIntValidator::~QIntValidator")

	if ptr.Pointer() != nil {
		C.QIntValidator_DestroyQIntValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIntValidator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QIntValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QIntValidator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QIntValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQIntValidatorTimerEvent
func callbackQIntValidatorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIntValidator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIntValidatorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIntValidator) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIntValidator::timerEvent")

	if ptr.Pointer() != nil {
		C.QIntValidator_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIntValidator) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIntValidator::timerEvent")

	if ptr.Pointer() != nil {
		C.QIntValidator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIntValidator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QIntValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QIntValidator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QIntValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQIntValidatorChildEvent
func callbackQIntValidatorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIntValidator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIntValidatorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIntValidator) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIntValidator::childEvent")

	if ptr.Pointer() != nil {
		C.QIntValidator_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIntValidator) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIntValidator::childEvent")

	if ptr.Pointer() != nil {
		C.QIntValidator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIntValidator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QIntValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QIntValidator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QIntValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQIntValidatorCustomEvent
func callbackQIntValidatorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIntValidator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIntValidatorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIntValidator) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QIntValidator::customEvent")

	if ptr.Pointer() != nil {
		C.QIntValidator_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIntValidator) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QIntValidator::customEvent")

	if ptr.Pointer() != nil {
		C.QIntValidator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
