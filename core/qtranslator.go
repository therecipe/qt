package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTranslator struct {
	QObject
}

type QTranslator_ITF interface {
	QObject_ITF
	QTranslator_PTR() *QTranslator
}

func PointerFromQTranslator(ptr QTranslator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTranslator_PTR().Pointer()
	}
	return nil
}

func NewQTranslatorFromPointer(ptr unsafe.Pointer) *QTranslator {
	var n = new(QTranslator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTranslator_") {
		n.SetObjectName("QTranslator_" + qt.Identifier())
	}
	return n
}

func (ptr *QTranslator) QTranslator_PTR() *QTranslator {
	return ptr
}

func NewQTranslator(parent QObject_ITF) *QTranslator {
	defer qt.Recovering("QTranslator::QTranslator")

	return NewQTranslatorFromPointer(C.QTranslator_NewQTranslator(PointerFromQObject(parent)))
}

func (ptr *QTranslator) IsEmpty() bool {
	defer qt.Recovering("QTranslator::isEmpty")

	if ptr.Pointer() != nil {
		return C.QTranslator_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTranslator) Load2(locale QLocale_ITF, filename string, prefix string, directory string, suffix string) bool {
	defer qt.Recovering("QTranslator::load")

	if ptr.Pointer() != nil {
		return C.QTranslator_Load2(ptr.Pointer(), PointerFromQLocale(locale), C.CString(filename), C.CString(prefix), C.CString(directory), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Load(filename string, directory string, search_delimiters string, suffix string) bool {
	defer qt.Recovering("QTranslator::load")

	if ptr.Pointer() != nil {
		return C.QTranslator_Load(ptr.Pointer(), C.CString(filename), C.CString(directory), C.CString(search_delimiters), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	defer qt.Recovering("QTranslator::translate")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTranslator_Translate(ptr.Pointer(), C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
	}
	return ""
}

func (ptr *QTranslator) DestroyQTranslator() {
	defer qt.Recovering("QTranslator::~QTranslator")

	if ptr.Pointer() != nil {
		C.QTranslator_DestroyQTranslator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTranslator) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QTranslator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTranslator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTranslator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTranslatorTimerEvent
func callbackQTranslatorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTranslator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQTranslatorFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTranslator) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QTranslator::timerEvent")

	if ptr.Pointer() != nil {
		C.QTranslator_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QTranslator) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QTranslator::timerEvent")

	if ptr.Pointer() != nil {
		C.QTranslator_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QTranslator) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QTranslator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTranslator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTranslator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTranslatorChildEvent
func callbackQTranslatorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTranslator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQTranslatorFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QTranslator) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QTranslator::childEvent")

	if ptr.Pointer() != nil {
		C.QTranslator_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QTranslator) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QTranslator::childEvent")

	if ptr.Pointer() != nil {
		C.QTranslator_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QTranslator) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QTranslator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTranslator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTranslator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTranslatorCustomEvent
func callbackQTranslatorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTranslator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQTranslatorFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QTranslator) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QTranslator::customEvent")

	if ptr.Pointer() != nil {
		C.QTranslator_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QTranslator) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QTranslator::customEvent")

	if ptr.Pointer() != nil {
		C.QTranslator_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
