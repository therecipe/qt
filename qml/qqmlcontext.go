package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlContext struct {
	core.QObject
}

type QQmlContext_ITF interface {
	core.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func PointerFromQQmlContext(ptr QQmlContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContext_PTR().Pointer()
	}
	return nil
}

func NewQQmlContextFromPointer(ptr unsafe.Pointer) *QQmlContext {
	var n = new(QQmlContext)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlContext_") {
		n.SetObjectName("QQmlContext_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return ptr
}

func NewQQmlContext2(parentContext QQmlContext_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	return NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext2(PointerFromQQmlContext(parentContext), core.PointerFromQObject(parent)))
}

func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	return NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func (ptr *QQmlContext) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlContext::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlContext_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	defer qt.Recovering("QQmlContext::contextObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlContext_ContextObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ContextProperty(name string) *core.QVariant {
	defer qt.Recovering("QQmlContext::contextProperty")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlContext_ContextProperty(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QQmlContext) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlContext::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlContext_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) IsValid() bool {
	defer qt.Recovering("QQmlContext::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlContext_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {
	defer qt.Recovering("QQmlContext::nameForObject")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_NameForObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return ""
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	defer qt.Recovering("QQmlContext::parentContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlContext_ParentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ResolvedUrl(src core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QQmlContext::resolvedUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlContext_ResolvedUrl(ptr.Pointer(), core.PointerFromQUrl(src)))
	}
	return nil
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QQmlContext::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextObject")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty(ptr.Pointer(), C.CString(name), core.PointerFromQObject(value))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty2(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	defer qt.Recovering("QQmlContext::~QQmlContext")

	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlContext) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlContext::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlContext::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlContextTimerEvent
func callbackQQmlContextTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlContext::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlContext) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlContext::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlContext::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlContextChildEvent
func callbackQQmlContextChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlContext::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlContext) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlContext::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlContext::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlContextCustomEvent
func callbackQQmlContextCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlContext::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
