package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) *QJSEngine {
	var n = new(QJSEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QJSEngine_") {
		n.SetObjectName("QJSEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func NewQJSEngine() *QJSEngine {
	defer qt.Recovering("QJSEngine::QJSEngine")

	return NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	defer qt.Recovering("QJSEngine::QJSEngine")

	return NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
}

func (ptr *QJSEngine) CollectGarbage() {
	defer qt.Recovering("QJSEngine::collectGarbage")

	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {
	defer qt.Recovering("QJSEngine::evaluate")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_Evaluate(ptr.Pointer(), C.CString(program), C.CString(fileName), C.int(lineNumber)))
	}
	return nil
}

func (ptr *QJSEngine) GlobalObject() *QJSValue {
	defer qt.Recovering("QJSEngine::globalObject")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_GlobalObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) InstallTranslatorFunctions(object QJSValue_ITF) {
	defer qt.Recovering("QJSEngine::installTranslatorFunctions")

	if ptr.Pointer() != nil {
		C.QJSEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQJSValue(object))
	}
}

func (ptr *QJSEngine) NewObject() *QJSValue {
	defer qt.Recovering("QJSEngine::newObject")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_NewObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	defer qt.Recovering("QJSEngine::newQObject")

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	defer qt.Recovering("QJSEngine::~QJSEngine")

	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QJSEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QJSEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QJSEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQJSEngineTimerEvent
func callbackQJSEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QJSEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QJSEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QJSEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QJSEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QJSEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QJSEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQJSEngineChildEvent
func callbackQJSEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QJSEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QJSEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QJSEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QJSEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QJSEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQJSEngineCustomEvent
func callbackQJSEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QJSEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QJSEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QJSEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
