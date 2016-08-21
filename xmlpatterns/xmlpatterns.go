// +build !minimal

package xmlpatterns

//#include <stdlib.h>
//#include "xmlpatterns.h"
import "C"
import (
	"encoding/hex"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandler_ITF interface {
	core.QObject_ITF
	QAbstractMessageHandler_PTR() *QAbstractMessageHandler
}

func (p *QAbstractMessageHandler) QAbstractMessageHandler_PTR() *QAbstractMessageHandler {
	return p
}

func (p *QAbstractMessageHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QAbstractMessageHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQAbstractMessageHandler(ptr QAbstractMessageHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractMessageHandler_PTR().Pointer()
	}
	return nil
}

func NewQAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) *QAbstractMessageHandler {
	var n = new(QAbstractMessageHandler)
	n.SetPointer(ptr)
	return n
}

func newQAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) *QAbstractMessageHandler {
	var n = NewQAbstractMessageHandlerFromPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractMessageHandler_") {
		n.SetObjectName("QAbstractMessageHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {
	defer qt.Recovering("QAbstractMessageHandler::~QAbstractMessageHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractMessageHandler_TimerEvent
func callbackQAbstractMessageHandler_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QAbstractMessageHandler) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractMessageHandler_ChildEvent
func callbackQAbstractMessageHandler_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QAbstractMessageHandler) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractMessageHandler_ConnectNotify
func callbackQAbstractMessageHandler_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractMessageHandler::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QAbstractMessageHandler) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_CustomEvent
func callbackQAbstractMessageHandler_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QAbstractMessageHandler) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractMessageHandler_DeleteLater
func callbackQAbstractMessageHandler_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractMessageHandler::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractMessageHandler) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QAbstractMessageHandler::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QAbstractMessageHandler) DeleteLater() {
	defer qt.Recovering("QAbstractMessageHandler::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QAbstractMessageHandler_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractMessageHandler) DeleteLaterDefault() {
	defer qt.Recovering("QAbstractMessageHandler::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QAbstractMessageHandler_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractMessageHandler_DisconnectNotify
func callbackQAbstractMessageHandler_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractMessageHandler::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractMessageHandler::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QAbstractMessageHandler) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractMessageHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractMessageHandler::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_Event
func callbackQAbstractMessageHandler_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractMessageHandler::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QAbstractMessageHandler) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractMessageHandler::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectEvent() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QAbstractMessageHandler) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractMessageHandler::event")

	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractMessageHandler) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractMessageHandler::event")

	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_EventFilter
func callbackQAbstractMessageHandler_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractMessageHandler::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QAbstractMessageHandler) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractMessageHandler::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QAbstractMessageHandler) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractMessageHandler::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractMessageHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractMessageHandler::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_MetaObject
func callbackQAbstractMessageHandler_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractMessageHandler::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractMessageHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractMessageHandler) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QAbstractMessageHandler::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QAbstractMessageHandler::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QAbstractMessageHandler) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QAbstractMessageHandler::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractMessageHandler_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractMessageHandler) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QAbstractMessageHandler::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractMessageHandler_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolver_ITF interface {
	core.QObject_ITF
	QAbstractUriResolver_PTR() *QAbstractUriResolver
}

func (p *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return p
}

func (p *QAbstractUriResolver) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QAbstractUriResolver) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractUriResolverFromPointer(ptr unsafe.Pointer) *QAbstractUriResolver {
	var n = new(QAbstractUriResolver)
	n.SetPointer(ptr)
	return n
}

func newQAbstractUriResolverFromPointer(ptr unsafe.Pointer) *QAbstractUriResolver {
	var n = NewQAbstractUriResolverFromPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractUriResolver_") {
		n.SetObjectName("QAbstractUriResolver_" + qt.Identifier())
	}
	return n
}

func NewQAbstractUriResolver(parent core.QObject_ITF) *QAbstractUriResolver {
	defer qt.Recovering("QAbstractUriResolver::QAbstractUriResolver")

	return newQAbstractUriResolverFromPointer(C.QAbstractUriResolver_NewQAbstractUriResolver(core.PointerFromQObject(parent)))
}

//export callbackQAbstractUriResolver_Resolve
func callbackQAbstractUriResolver_Resolve(ptr unsafe.Pointer, ptrName *C.char, relative unsafe.Pointer, baseURI unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractUriResolver::resolve")

	if signal := qt.GetSignal(C.GoString(ptrName), "resolve"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*core.QUrl, *core.QUrl) *core.QUrl)(core.NewQUrlFromPointer(relative), core.NewQUrlFromPointer(baseURI)))
	}

	return core.PointerFromQUrl(nil)
}

func (ptr *QAbstractUriResolver) ConnectResolve(f func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl) {
	defer qt.Recovering("connect QAbstractUriResolver::resolve")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resolve", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectResolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) {
	defer qt.Recovering("disconnect QAbstractUriResolver::resolve")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resolve")
	}
}

func (ptr *QAbstractUriResolver) Resolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractUriResolver::resolve")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QAbstractUriResolver_Resolve(ptr.Pointer(), core.PointerFromQUrl(relative), core.PointerFromQUrl(baseURI)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	defer qt.Recovering("QAbstractUriResolver::~QAbstractUriResolver")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractUriResolver_TimerEvent
func callbackQAbstractUriResolver_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QAbstractUriResolver) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractUriResolver_ChildEvent
func callbackQAbstractUriResolver_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QAbstractUriResolver) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractUriResolver_ConnectNotify
func callbackQAbstractUriResolver_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractUriResolver::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QAbstractUriResolver::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QAbstractUriResolver) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractUriResolver::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractUriResolver::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_CustomEvent
func callbackQAbstractUriResolver_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QAbstractUriResolver) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractUriResolver_DeleteLater
func callbackQAbstractUriResolver_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractUriResolver::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractUriResolver) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QAbstractUriResolver::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QAbstractUriResolver::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QAbstractUriResolver) DeleteLater() {
	defer qt.Recovering("QAbstractUriResolver::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QAbstractUriResolver_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) DeleteLaterDefault() {
	defer qt.Recovering("QAbstractUriResolver::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QAbstractUriResolver_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractUriResolver_DisconnectNotify
func callbackQAbstractUriResolver_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractUriResolver::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QAbstractUriResolver::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QAbstractUriResolver) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractUriResolver::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractUriResolver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractUriResolver::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_Event
func callbackQAbstractUriResolver_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractUriResolver::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QAbstractUriResolver) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractUriResolver::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QAbstractUriResolver) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractUriResolver::event")

	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractUriResolver) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractUriResolver::event")

	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_EventFilter
func callbackQAbstractUriResolver_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractUriResolver::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QAbstractUriResolver) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractUriResolver::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QAbstractUriResolver::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QAbstractUriResolver) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractUriResolver::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractUriResolver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractUriResolver::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_MetaObject
func callbackQAbstractUriResolver_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractUriResolver::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractUriResolverFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractUriResolver) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QAbstractUriResolver::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QAbstractUriResolver::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QAbstractUriResolver) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QAbstractUriResolver::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractUriResolver_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractUriResolver) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QAbstractUriResolver::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractUriResolver_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int64

const (
	QAbstractXmlNodeModel__Parent          = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     = QAbstractXmlNodeModel__SimpleAxis(3)
)

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModel_ITF interface {
	core.QSharedData_ITF
	QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel
}

func (p *QAbstractXmlNodeModel) QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel {
	return p
}

func (p *QAbstractXmlNodeModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSharedData_PTR().Pointer()
	}
	return nil
}

func (p *QAbstractXmlNodeModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSharedData_PTR().SetPointer(ptr)
	}
}

func PointerFromQAbstractXmlNodeModel(ptr QAbstractXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) *QAbstractXmlNodeModel {
	var n = new(QAbstractXmlNodeModel)
	n.SetPointer(ptr)
	return n
}

func newQAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) *QAbstractXmlNodeModel {
	var n = NewQAbstractXmlNodeModelFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractXmlNodeModel_") {
		n.SetObjectNameAbs("QAbstractXmlNodeModel_" + qt.Identifier())
	}
	return n
}

//export callbackQAbstractXmlNodeModel_BaseUri
func callbackQAbstractXmlNodeModel_BaseUri(ptr unsafe.Pointer, ptrName *C.char, n unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::baseUri")

	if signal := qt.GetSignal(C.GoString(ptrName), "baseUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectBaseUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "baseUri", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectBaseUri(n QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "baseUri")
	}
}

func (ptr *QAbstractXmlNodeModel) BaseUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_CompareOrder
func callbackQAbstractXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ptrName *C.char, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractXmlNodeModel::compareOrder")

	if signal := qt.GetSignal(C.GoString(ptrName), "compareOrder"); signal != nil {
		return C.int(signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.int(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compareOrder", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectCompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compareOrder")
	}
}

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	defer qt.Recovering("QAbstractXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QAbstractXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) CreateIndex(data int64) *QXmlNodeModelIndex {
	defer qt.Recovering("QAbstractXmlNodeModel::createIndex")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex(ptr.Pointer(), C.longlong(data)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex3(data int64, additionalData int64) *QXmlNodeModelIndex {
	defer qt.Recovering("QAbstractXmlNodeModel::createIndex")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex3(ptr.Pointer(), C.longlong(data), C.longlong(additionalData)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex2(pointer unsafe.Pointer, additionalData int64) *QXmlNodeModelIndex {
	defer qt.Recovering("QAbstractXmlNodeModel::createIndex")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex2(ptr.Pointer(), pointer, C.longlong(additionalData)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_DocumentUri
func callbackQAbstractXmlNodeModel_DocumentUri(ptr unsafe.Pointer, ptrName *C.char, n unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::documentUri")

	if signal := qt.GetSignal(C.GoString(ptrName), "documentUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "documentUri", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectDocumentUri(n QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "documentUri")
	}
}

func (ptr *QAbstractXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_ElementById
func callbackQAbstractXmlNodeModel_ElementById(ptr unsafe.Pointer, ptrName *C.char, id unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::elementById")

	if signal := qt.GetSignal(C.GoString(ptrName), "elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlName) *QXmlNodeModelIndex)(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::elementById")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "elementById", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectElementById(id QXmlName_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::elementById")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "elementById")
	}
}

func (ptr *QAbstractXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QAbstractXmlNodeModel::elementById")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_ElementById(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Kind
func callbackQAbstractXmlNodeModel_Kind(ptr unsafe.Pointer, ptrName *C.char, ni unsafe.Pointer) C.int {
	defer qt.Recovering("callback QAbstractXmlNodeModel::kind")

	if signal := qt.GetSignal(C.GoString(ptrName), "kind"); signal != nil {
		return C.int(signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.int(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::kind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "kind", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectKind(ni QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::kind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "kind")
	}
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	defer qt.Recovering("QAbstractXmlNodeModel::kind")

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QAbstractXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

//export callbackQAbstractXmlNodeModel_Name
func callbackQAbstractXmlNodeModel_Name(ptr unsafe.Pointer, ptrName *C.char, ni unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::name")

	if signal := qt.GetSignal(C.GoString(ptrName), "name"); signal != nil {
		return PointerFromQXmlName(signal.(func(*QXmlNodeModelIndex) *QXmlName)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::name")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "name", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectName(ni QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::name")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "name")
	}
}

func (ptr *QAbstractXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	defer qt.Recovering("QAbstractXmlNodeModel::name")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQAbstractXmlNodeModel_NextFromSimpleAxis
func callbackQAbstractXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, ptrName *C.char, axis C.int, origin unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::nextFromSimpleAxis")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::nextFromSimpleAxis")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "nextFromSimpleAxis", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::nextFromSimpleAxis")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "nextFromSimpleAxis")
	}
}

func (ptr *QAbstractXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QAbstractXmlNodeModel::nextFromSimpleAxis")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.int(axis), PointerFromQXmlNodeModelIndex(origin)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Root
func callbackQAbstractXmlNodeModel_Root(ptr unsafe.Pointer, ptrName *C.char, n unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::root")

	if signal := qt.GetSignal(C.GoString(ptrName), "root"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::root")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "root", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectRoot(n QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::root")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "root")
	}
}

func (ptr *QAbstractXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QAbstractXmlNodeModel::root")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) SourceLocation(index QXmlNodeModelIndex_ITF) *QSourceLocation {
	defer qt.Recovering("QAbstractXmlNodeModel::sourceLocation")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSourceLocationFromPointer(C.QAbstractXmlNodeModel_SourceLocation(ptr.Pointer(), PointerFromQXmlNodeModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_StringValue
func callbackQAbstractXmlNodeModel_StringValue(ptr unsafe.Pointer, ptrName *C.char, n unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QAbstractXmlNodeModel::stringValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "stringValue"); signal != nil {
		return C.CString(signal.(func(*QXmlNodeModelIndex) string)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return C.CString("")
}

func (ptr *QAbstractXmlNodeModel) ConnectStringValue(f func(n *QXmlNodeModelIndex) string) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "stringValue", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectStringValue(n QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "stringValue")
	}
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndex_ITF) string {
	defer qt.Recovering("QAbstractXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return ""
}

//export callbackQAbstractXmlNodeModel_TypedValue
func callbackQAbstractXmlNodeModel_TypedValue(ptr unsafe.Pointer, ptrName *C.char, node unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractXmlNodeModel::typedValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "typedValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*QXmlNodeModelIndex) *core.QVariant)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QAbstractXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {
	defer qt.Recovering("connect QAbstractXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "typedValue", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectTypedValue(node QXmlNodeModelIndex_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "typedValue")
	}
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	defer qt.Recovering("QAbstractXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {
	defer qt.Recovering("QAbstractXmlNodeModel::~QAbstractXmlNodeModel")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlNodeModel) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractXmlNodeModel::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractXmlNodeModel::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QAbstractXmlNodeModel_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

type QAbstractXmlReceiver struct {
	ptr unsafe.Pointer
}

type QAbstractXmlReceiver_ITF interface {
	QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver
}

func (p *QAbstractXmlReceiver) QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver {
	return p
}

func (p *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAbstractXmlReceiver) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQAbstractXmlReceiver(ptr QAbstractXmlReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) *QAbstractXmlReceiver {
	var n = new(QAbstractXmlReceiver)
	n.SetPointer(ptr)
	return n
}

func newQAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) *QAbstractXmlReceiver {
	var n = NewQAbstractXmlReceiverFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractXmlReceiver_") {
		n.SetObjectNameAbs("QAbstractXmlReceiver_" + qt.Identifier())
	}
	return n
}

func NewQAbstractXmlReceiver() *QAbstractXmlReceiver {
	defer qt.Recovering("QAbstractXmlReceiver::QAbstractXmlReceiver")

	return newQAbstractXmlReceiverFromPointer(C.QAbstractXmlReceiver_NewQAbstractXmlReceiver())
}

//export callbackQAbstractXmlReceiver_AtomicValue
func callbackQAbstractXmlReceiver_AtomicValue(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractXmlReceiver::atomicValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectAtomicValue(f func(value *core.QVariant)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::atomicValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "atomicValue", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectAtomicValue(value core.QVariant_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::atomicValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "atomicValue")
	}
}

func (ptr *QAbstractXmlReceiver) AtomicValue(value core.QVariant_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::atomicValue")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQAbstractXmlReceiver_Attribute
func callbackQAbstractXmlReceiver_Attribute(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer, value unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractXmlReceiver::attribute")

	if signal := qt.GetSignal(C.GoString(ptrName), "attribute"); signal != nil {
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::attribute")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "attribute", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectAttribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::attribute")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "attribute")
	}
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::attribute")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQAbstractXmlReceiver_Characters
func callbackQAbstractXmlReceiver_Characters(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractXmlReceiver::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectCharacters(value core.QStringRef_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::characters")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQAbstractXmlReceiver_Comment
func callbackQAbstractXmlReceiver_Comment(ptr unsafe.Pointer, ptrName *C.char, value *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectComment(value string) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {
	defer qt.Recovering("QAbstractXmlReceiver::comment")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QAbstractXmlReceiver_Comment(ptr.Pointer(), valueC)
	}
}

//export callbackQAbstractXmlReceiver_EndDocument
func callbackQAbstractXmlReceiver_EndDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QAbstractXmlReceiver::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

func (ptr *QAbstractXmlReceiver) EndDocument() {
	defer qt.Recovering("QAbstractXmlReceiver::endDocument")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndDocument(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_EndElement
func callbackQAbstractXmlReceiver_EndElement(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QAbstractXmlReceiver::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndElement() {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

func (ptr *QAbstractXmlReceiver) EndElement() {
	defer qt.Recovering("QAbstractXmlReceiver::endElement")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndElement(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_EndOfSequence
func callbackQAbstractXmlReceiver_EndOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QAbstractXmlReceiver::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {
	defer qt.Recovering("QAbstractXmlReceiver::endOfSequence")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndOfSequence(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_NamespaceBinding
func callbackQAbstractXmlReceiver_NamespaceBinding(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractXmlReceiver::namespaceBinding")

	if signal := qt.GetSignal(C.GoString(ptrName), "namespaceBinding"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectNamespaceBinding(f func(name *QXmlName)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::namespaceBinding")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "namespaceBinding", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectNamespaceBinding(name QXmlName_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::namespaceBinding")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "namespaceBinding")
	}
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlName_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::namespaceBinding")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQAbstractXmlReceiver_ProcessingInstruction
func callbackQAbstractXmlReceiver_ProcessingInstruction(ptr unsafe.Pointer, ptrName *C.char, target unsafe.Pointer, value *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::processingInstruction")

	if signal := qt.GetSignal(C.GoString(ptrName), "processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(target), C.GoString(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectProcessingInstruction(f func(target *QXmlName, value string)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::processingInstruction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "processingInstruction", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectProcessingInstruction(target QXmlName_ITF, value string) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::processingInstruction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "processingInstruction")
	}
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlName_ITF, value string) {
	defer qt.Recovering("QAbstractXmlReceiver::processingInstruction")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QAbstractXmlReceiver_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(target), valueC)
	}
}

//export callbackQAbstractXmlReceiver_StartDocument
func callbackQAbstractXmlReceiver_StartDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QAbstractXmlReceiver::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

func (ptr *QAbstractXmlReceiver) StartDocument() {
	defer qt.Recovering("QAbstractXmlReceiver::startDocument")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartDocument(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_StartElement
func callbackQAbstractXmlReceiver_StartElement(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractXmlReceiver::startElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "startElement"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartElement(f func(name *QXmlName)) {
	defer qt.Recovering("connect QAbstractXmlReceiver::startElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startElement", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartElement(name QXmlName_ITF) {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::startElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startElement")
	}
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlName_ITF) {
	defer qt.Recovering("QAbstractXmlReceiver::startElement")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQAbstractXmlReceiver_StartOfSequence
func callbackQAbstractXmlReceiver_StartOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractXmlReceiver::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QAbstractXmlReceiver::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QAbstractXmlReceiver::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {
	defer qt.Recovering("QAbstractXmlReceiver::startOfSequence")

	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {
	defer qt.Recovering("QAbstractXmlReceiver::~QAbstractXmlReceiver")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlReceiver) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractXmlReceiver::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlReceiver_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractXmlReceiver) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractXmlReceiver::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QAbstractXmlReceiver_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModel_ITF interface {
	QAbstractXmlNodeModel_ITF
	QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel
}

func (p *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return p
}

func (p *QSimpleXmlNodeModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func (p *QSimpleXmlNodeModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractXmlNodeModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQSimpleXmlNodeModel(ptr QSimpleXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSimpleXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) *QSimpleXmlNodeModel {
	var n = new(QSimpleXmlNodeModel)
	n.SetPointer(ptr)
	return n
}

func newQSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) *QSimpleXmlNodeModel {
	var n = NewQSimpleXmlNodeModelFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSimpleXmlNodeModel_") {
		n.SetObjectNameAbs("QSimpleXmlNodeModel_" + qt.Identifier())
	}
	return n
}

//export callbackQSimpleXmlNodeModel_BaseUri
func callbackQSimpleXmlNodeModel_BaseUri(ptr unsafe.Pointer, ptrName *C.char, node unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::baseUri")

	if signal := qt.GetSignal(C.GoString(ptrName), "baseUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQUrl(NewQSimpleXmlNodeModelFromPointer(ptr).BaseUriDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) ConnectBaseUri(f func(node *QXmlNodeModelIndex) *core.QUrl) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "baseUri", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectBaseUri() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "baseUri")
	}
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QSimpleXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) BaseUriDefault(node QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QSimpleXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUriDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_ElementById
func callbackQSimpleXmlNodeModel_ElementById(ptr unsafe.Pointer, ptrName *C.char, id unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::elementById")

	if signal := qt.GetSignal(C.GoString(ptrName), "elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlName) *QXmlNodeModelIndex)(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).ElementByIdDefault(NewQXmlNameFromPointer(id)))
}

func (ptr *QSimpleXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::elementById")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "elementById", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectElementById() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::elementById")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "elementById")
	}
}

func (ptr *QSimpleXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QSimpleXmlNodeModel::elementById")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementById(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ElementByIdDefault(id QXmlName_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QSimpleXmlNodeModel::elementById")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementByIdDefault(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NamePool() *QXmlNamePool {
	defer qt.Recovering("QSimpleXmlNodeModel::namePool")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QSimpleXmlNodeModel_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_StringValue
func callbackQSimpleXmlNodeModel_StringValue(ptr unsafe.Pointer, ptrName *C.char, node unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSimpleXmlNodeModel::stringValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "stringValue"); signal != nil {
		return C.CString(signal.(func(*QXmlNodeModelIndex) string)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return C.CString(NewQSimpleXmlNodeModelFromPointer(ptr).StringValueDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) ConnectStringValue(f func(node *QXmlNodeModelIndex) string) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "stringValue", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectStringValue() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "stringValue")
	}
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {
	defer qt.Recovering("QSimpleXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) StringValueDefault(node QXmlNodeModelIndex_ITF) string {
	defer qt.Recovering("QSimpleXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_StringValueDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	defer qt.Recovering("QSimpleXmlNodeModel::~QSimpleXmlNodeModel")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSimpleXmlNodeModel) ObjectNameAbs() string {
	defer qt.Recovering("QSimpleXmlNodeModel::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSimpleXmlNodeModel::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSimpleXmlNodeModel_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

//export callbackQSimpleXmlNodeModel_CompareOrder
func callbackQSimpleXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ptrName *C.char, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSimpleXmlNodeModel::compareOrder")

	if signal := qt.GetSignal(C.GoString(ptrName), "compareOrder"); signal != nil {
		return C.int(signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.int(0)
}

func (ptr *QSimpleXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compareOrder", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectCompareOrder() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compareOrder")
	}
}

func (ptr *QSimpleXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	defer qt.Recovering("QSimpleXmlNodeModel::compareOrder")

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QSimpleXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

//export callbackQSimpleXmlNodeModel_DocumentUri
func callbackQSimpleXmlNodeModel_DocumentUri(ptr unsafe.Pointer, ptrName *C.char, n unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::documentUri")

	if signal := qt.GetSignal(C.GoString(ptrName), "documentUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(nil)
}

func (ptr *QSimpleXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "documentUri", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectDocumentUri() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "documentUri")
	}
}

func (ptr *QSimpleXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QSimpleXmlNodeModel::documentUri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Kind
func callbackQSimpleXmlNodeModel_Kind(ptr unsafe.Pointer, ptrName *C.char, ni unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSimpleXmlNodeModel::kind")

	if signal := qt.GetSignal(C.GoString(ptrName), "kind"); signal != nil {
		return C.int(signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.int(0)
}

func (ptr *QSimpleXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::kind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "kind", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectKind() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::kind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "kind")
	}
}

func (ptr *QSimpleXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	defer qt.Recovering("QSimpleXmlNodeModel::kind")

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QSimpleXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

//export callbackQSimpleXmlNodeModel_Name
func callbackQSimpleXmlNodeModel_Name(ptr unsafe.Pointer, ptrName *C.char, ni unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::name")

	if signal := qt.GetSignal(C.GoString(ptrName), "name"); signal != nil {
		return PointerFromQXmlName(signal.(func(*QXmlNodeModelIndex) *QXmlName)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQSimpleXmlNodeModelFromPointer(ptr).NameDefault(NewQXmlNodeModelIndexFromPointer(ni)))
}

func (ptr *QSimpleXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::name")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "name", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectName() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::name")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "name")
	}
}

func (ptr *QSimpleXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	defer qt.Recovering("QSimpleXmlNodeModel::name")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NameDefault(ni QXmlNodeModelIndex_ITF) *QXmlName {
	defer qt.Recovering("QSimpleXmlNodeModel::name")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQSimpleXmlNodeModel_NextFromSimpleAxis
func callbackQSimpleXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, ptrName *C.char, axis C.int, origin unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::nextFromSimpleAxis")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(nil)
}

func (ptr *QSimpleXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::nextFromSimpleAxis")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "nextFromSimpleAxis", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectNextFromSimpleAxis() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::nextFromSimpleAxis")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "nextFromSimpleAxis")
	}
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QSimpleXmlNodeModel::nextFromSimpleAxis")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.int(axis), PointerFromQXmlNodeModelIndex(origin)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Root
func callbackQSimpleXmlNodeModel_Root(ptr unsafe.Pointer, ptrName *C.char, n unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::root")

	if signal := qt.GetSignal(C.GoString(ptrName), "root"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(nil)
}

func (ptr *QSimpleXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::root")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "root", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectRoot() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::root")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "root")
	}
}

func (ptr *QSimpleXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QSimpleXmlNodeModel::root")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_TypedValue
func callbackQSimpleXmlNodeModel_TypedValue(ptr unsafe.Pointer, ptrName *C.char, node unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSimpleXmlNodeModel::typedValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "typedValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*QXmlNodeModelIndex) *core.QVariant)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QSimpleXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {
	defer qt.Recovering("connect QSimpleXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "typedValue", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectTypedValue() {
	defer qt.Recovering("disconnect QSimpleXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "typedValue")
	}
}

func (ptr *QSimpleXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	defer qt.Recovering("QSimpleXmlNodeModel::typedValue")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSimpleXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QSourceLocation struct {
	ptr unsafe.Pointer
}

type QSourceLocation_ITF interface {
	QSourceLocation_PTR() *QSourceLocation
}

func (p *QSourceLocation) QSourceLocation_PTR() *QSourceLocation {
	return p
}

func (p *QSourceLocation) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSourceLocation) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSourceLocation(ptr QSourceLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSourceLocation_PTR().Pointer()
	}
	return nil
}

func NewQSourceLocationFromPointer(ptr unsafe.Pointer) *QSourceLocation {
	var n = new(QSourceLocation)
	n.SetPointer(ptr)
	return n
}

func newQSourceLocationFromPointer(ptr unsafe.Pointer) *QSourceLocation {
	var n = NewQSourceLocationFromPointer(ptr)
	return n
}

func NewQSourceLocation() *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return newQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation())
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return newQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation2(PointerFromQSourceLocation(other)))
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return newQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation3(core.PointerFromQUrl(u), C.int(l), C.int(c)))
}

func (ptr *QSourceLocation) Column() int64 {
	defer qt.Recovering("QSourceLocation::column")

	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) IsNull() bool {
	defer qt.Recovering("QSourceLocation::isNull")

	if ptr.Pointer() != nil {
		return C.QSourceLocation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSourceLocation) Line() int64 {
	defer qt.Recovering("QSourceLocation::line")

	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) SetColumn(newColumn int64) {
	defer qt.Recovering("QSourceLocation::setColumn")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetColumn(ptr.Pointer(), C.longlong(newColumn))
	}
}

func (ptr *QSourceLocation) SetLine(newLine int64) {
	defer qt.Recovering("QSourceLocation::setLine")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetLine(ptr.Pointer(), C.longlong(newLine))
	}
}

func (ptr *QSourceLocation) SetUri(newUri core.QUrl_ITF) {
	defer qt.Recovering("QSourceLocation::setUri")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetUri(ptr.Pointer(), core.PointerFromQUrl(newUri))
	}
}

func (ptr *QSourceLocation) Uri() *core.QUrl {
	defer qt.Recovering("QSourceLocation::uri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSourceLocation_Uri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {
	defer qt.Recovering("QSourceLocation::~QSourceLocation")

	if ptr.Pointer() != nil {
		C.QSourceLocation_DestroyQSourceLocation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlFormatter struct {
	QXmlSerializer
}

type QXmlFormatter_ITF interface {
	QXmlSerializer_ITF
	QXmlFormatter_PTR() *QXmlFormatter
}

func (p *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return p
}

func (p *QXmlFormatter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func (p *QXmlFormatter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QXmlSerializer_PTR().SetPointer(ptr)
	}
}

func PointerFromQXmlFormatter(ptr QXmlFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlFormatter_PTR().Pointer()
	}
	return nil
}

func NewQXmlFormatterFromPointer(ptr unsafe.Pointer) *QXmlFormatter {
	var n = new(QXmlFormatter)
	n.SetPointer(ptr)
	return n
}

func newQXmlFormatterFromPointer(ptr unsafe.Pointer) *QXmlFormatter {
	var n = NewQXmlFormatterFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlFormatter_") {
		n.SetObjectNameAbs("QXmlFormatter_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlFormatter) DestroyQXmlFormatter() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	defer qt.Recovering("QXmlFormatter::QXmlFormatter")

	return newQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

//export callbackQXmlFormatter_AtomicValue
func callbackQXmlFormatter_AtomicValue(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlFormatter::atomicValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).AtomicValueDefault(core.NewQVariantFromPointer(value))
	}
}

func (ptr *QXmlFormatter) ConnectAtomicValue(f func(value *core.QVariant)) {
	defer qt.Recovering("connect QXmlFormatter::atomicValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "atomicValue", f)
	}
}

func (ptr *QXmlFormatter) DisconnectAtomicValue() {
	defer qt.Recovering("disconnect QXmlFormatter::atomicValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "atomicValue")
	}
}

func (ptr *QXmlFormatter) AtomicValue(value core.QVariant_ITF) {
	defer qt.Recovering("QXmlFormatter::atomicValue")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QXmlFormatter) AtomicValueDefault(value core.QVariant_ITF) {
	defer qt.Recovering("QXmlFormatter::atomicValue")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_AtomicValueDefault(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQXmlFormatter_Attribute
func callbackQXmlFormatter_Attribute(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlFormatter::attribute")

	if signal := qt.GetSignal(C.GoString(ptrName), "attribute"); signal != nil {
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).AttributeDefault(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlFormatter) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlFormatter::attribute")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "attribute", f)
	}
}

func (ptr *QXmlFormatter) DisconnectAttribute() {
	defer qt.Recovering("disconnect QXmlFormatter::attribute")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "attribute")
	}
}

func (ptr *QXmlFormatter) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlFormatter::attribute")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) AttributeDefault(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlFormatter::attribute")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_AttributeDefault(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlFormatter_Characters
func callbackQXmlFormatter_Characters(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlFormatter::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlFormatter) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlFormatter::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlFormatter) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlFormatter::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

func (ptr *QXmlFormatter) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlFormatter::characters")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) CharactersDefault(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlFormatter::characters")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlFormatter_Comment
func callbackQXmlFormatter_Comment(ptr unsafe.Pointer, ptrName *C.char, value *C.char) {
	defer qt.Recovering("callback QXmlFormatter::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).CommentDefault(C.GoString(value))
	}
}

func (ptr *QXmlFormatter) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QXmlFormatter::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlFormatter) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlFormatter::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

func (ptr *QXmlFormatter) Comment(value string) {
	defer qt.Recovering("QXmlFormatter::comment")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_Comment(ptr.Pointer(), valueC)
	}
}

func (ptr *QXmlFormatter) CommentDefault(value string) {
	defer qt.Recovering("QXmlFormatter::comment")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_CommentDefault(ptr.Pointer(), valueC)
	}
}

//export callbackQXmlFormatter_EndDocument
func callbackQXmlFormatter_EndDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlFormatter) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

func (ptr *QXmlFormatter) EndDocument() {
	defer qt.Recovering("QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndDocumentDefault() {
	defer qt.Recovering("QXmlFormatter::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_EndElement
func callbackQXmlFormatter_EndElement(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlFormatter) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlFormatter::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

func (ptr *QXmlFormatter) EndElement() {
	defer qt.Recovering("QXmlFormatter::endElement")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndElementDefault() {
	defer qt.Recovering("QXmlFormatter::endElement")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElementDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_EndOfSequence
func callbackQXmlFormatter_EndOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlFormatter) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

func (ptr *QXmlFormatter) EndOfSequence() {
	defer qt.Recovering("QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndOfSequenceDefault() {
	defer qt.Recovering("QXmlFormatter::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	defer qt.Recovering("QXmlFormatter::indentationDepth")

	if ptr.Pointer() != nil {
		return int(C.QXmlFormatter_IndentationDepth(ptr.Pointer()))
	}
	return 0
}

//export callbackQXmlFormatter_ProcessingInstruction
func callbackQXmlFormatter_ProcessingInstruction(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer, value *C.char) {
	defer qt.Recovering("callback QXmlFormatter::processingInstruction")

	if signal := qt.GetSignal(C.GoString(ptrName), "processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(name), C.GoString(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).ProcessingInstructionDefault(NewQXmlNameFromPointer(name), C.GoString(value))
	}
}

func (ptr *QXmlFormatter) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {
	defer qt.Recovering("connect QXmlFormatter::processingInstruction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "processingInstruction", f)
	}
}

func (ptr *QXmlFormatter) DisconnectProcessingInstruction() {
	defer qt.Recovering("disconnect QXmlFormatter::processingInstruction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "processingInstruction")
	}
}

func (ptr *QXmlFormatter) ProcessingInstruction(name QXmlName_ITF, value string) {
	defer qt.Recovering("QXmlFormatter::processingInstruction")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlFormatter) ProcessingInstructionDefault(name QXmlName_ITF, value string) {
	defer qt.Recovering("QXmlFormatter::processingInstruction")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_ProcessingInstructionDefault(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	defer qt.Recovering("QXmlFormatter::setIndentationDepth")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(depth))
	}
}

//export callbackQXmlFormatter_StartDocument
func callbackQXmlFormatter_StartDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlFormatter) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

func (ptr *QXmlFormatter) StartDocument() {
	defer qt.Recovering("QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartDocumentDefault() {
	defer qt.Recovering("QXmlFormatter::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_StartElement
func callbackQXmlFormatter_StartElement(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer) {
	defer qt.Recovering("callback QXmlFormatter::startElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "startElement"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	} else {
		NewQXmlFormatterFromPointer(ptr).StartElementDefault(NewQXmlNameFromPointer(name))
	}
}

func (ptr *QXmlFormatter) ConnectStartElement(f func(name *QXmlName)) {
	defer qt.Recovering("connect QXmlFormatter::startElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startElement", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartElement() {
	defer qt.Recovering("disconnect QXmlFormatter::startElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startElement")
	}
}

func (ptr *QXmlFormatter) StartElement(name QXmlName_ITF) {
	defer qt.Recovering("QXmlFormatter::startElement")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlFormatter) StartElementDefault(name QXmlName_ITF) {
	defer qt.Recovering("QXmlFormatter::startElement")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElementDefault(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQXmlFormatter_StartOfSequence
func callbackQXmlFormatter_StartOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlFormatter::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlFormatter) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

func (ptr *QXmlFormatter) StartOfSequence() {
	defer qt.Recovering("QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartOfSequenceDefault() {
	defer qt.Recovering("QXmlFormatter::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) ObjectNameAbs() string {
	defer qt.Recovering("QXmlFormatter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlFormatter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlFormatter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlFormatter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QXmlFormatter_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

//export callbackQXmlFormatter_NamespaceBinding
func callbackQXmlFormatter_NamespaceBinding(ptr unsafe.Pointer, ptrName *C.char, nb unsafe.Pointer) {
	defer qt.Recovering("callback QXmlFormatter::namespaceBinding")

	if signal := qt.GetSignal(C.GoString(ptrName), "namespaceBinding"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(nb))
	} else {
		NewQXmlFormatterFromPointer(ptr).NamespaceBindingDefault(NewQXmlNameFromPointer(nb))
	}
}

func (ptr *QXmlFormatter) ConnectNamespaceBinding(f func(nb *QXmlName)) {
	defer qt.Recovering("connect QXmlFormatter::namespaceBinding")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "namespaceBinding", f)
	}
}

func (ptr *QXmlFormatter) DisconnectNamespaceBinding() {
	defer qt.Recovering("disconnect QXmlFormatter::namespaceBinding")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "namespaceBinding")
	}
}

func (ptr *QXmlFormatter) NamespaceBinding(nb QXmlName_ITF) {
	defer qt.Recovering("QXmlFormatter::namespaceBinding")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlFormatter) NamespaceBindingDefault(nb QXmlName_ITF) {
	defer qt.Recovering("QXmlFormatter::namespaceBinding")

	if ptr.Pointer() != nil {
		C.QXmlFormatter_NamespaceBindingDefault(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

type QXmlItem struct {
	ptr unsafe.Pointer
}

type QXmlItem_ITF interface {
	QXmlItem_PTR() *QXmlItem
}

func (p *QXmlItem) QXmlItem_PTR() *QXmlItem {
	return p
}

func (p *QXmlItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlItem(ptr QXmlItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlItem_PTR().Pointer()
	}
	return nil
}

func NewQXmlItemFromPointer(ptr unsafe.Pointer) *QXmlItem {
	var n = new(QXmlItem)
	n.SetPointer(ptr)
	return n
}

func newQXmlItemFromPointer(ptr unsafe.Pointer) *QXmlItem {
	var n = NewQXmlItemFromPointer(ptr)
	return n
}

func NewQXmlItem() *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return newQXmlItemFromPointer(C.QXmlItem_NewQXmlItem())
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return newQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return newQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return newQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
}

func (ptr *QXmlItem) IsAtomicValue() bool {
	defer qt.Recovering("QXmlItem::isAtomicValue")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsAtomicValue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNode() bool {
	defer qt.Recovering("QXmlItem::isNode")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	defer qt.Recovering("QXmlItem::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) ToAtomicValue() *core.QVariant {
	defer qt.Recovering("QXmlItem::toAtomicValue")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QXmlItem_ToAtomicValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) ToNodeModelIndex() *QXmlNodeModelIndex {
	defer qt.Recovering("QXmlItem::toNodeModelIndex")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QXmlItem_ToNodeModelIndex(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	defer qt.Recovering("QXmlItem::~QXmlItem")

	if ptr.Pointer() != nil {
		C.QXmlItem_DestroyQXmlItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlName struct {
	ptr unsafe.Pointer
}

type QXmlName_ITF interface {
	QXmlName_PTR() *QXmlName
}

func (p *QXmlName) QXmlName_PTR() *QXmlName {
	return p
}

func (p *QXmlName) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlName) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlName(ptr QXmlName_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlName_PTR().Pointer()
	}
	return nil
}

func NewQXmlNameFromPointer(ptr unsafe.Pointer) *QXmlName {
	var n = new(QXmlName)
	n.SetPointer(ptr)
	return n
}

func newQXmlNameFromPointer(ptr unsafe.Pointer) *QXmlName {
	var n = NewQXmlNameFromPointer(ptr)
	return n
}

func (ptr *QXmlName) DestroyQXmlName() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQXmlName() *QXmlName {
	defer qt.Recovering("QXmlName::QXmlName")

	return newQXmlNameFromPointer(C.QXmlName_NewQXmlName())
}

func NewQXmlName2(namePool QXmlNamePool_ITF, localName string, namespaceURI string, prefix string) *QXmlName {
	defer qt.Recovering("QXmlName::QXmlName")

	var localNameC = C.CString(localName)
	defer C.free(unsafe.Pointer(localNameC))
	var namespaceURIC = C.CString(namespaceURI)
	defer C.free(unsafe.Pointer(namespaceURIC))
	var prefixC = C.CString(prefix)
	defer C.free(unsafe.Pointer(prefixC))
	return newQXmlNameFromPointer(C.QXmlName_NewQXmlName2(PointerFromQXmlNamePool(namePool), localNameC, namespaceURIC, prefixC))
}

func QXmlName_IsNCName(candidate string) bool {
	defer qt.Recovering("QXmlName::isNCName")

	var candidateC = C.CString(candidate)
	defer C.free(unsafe.Pointer(candidateC))
	return C.QXmlName_QXmlName_IsNCName(candidateC) != 0
}

func (ptr *QXmlName) IsNCName(candidate string) bool {
	defer qt.Recovering("QXmlName::isNCName")

	var candidateC = C.CString(candidate)
	defer C.free(unsafe.Pointer(candidateC))
	return C.QXmlName_QXmlName_IsNCName(candidateC) != 0
}

func (ptr *QXmlName) IsNull() bool {
	defer qt.Recovering("QXmlName::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlName_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_LocalName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::namespaceUri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_NamespaceUri(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_Prefix(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::toClarkName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_ToClarkName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

type QXmlNamePool struct {
	ptr unsafe.Pointer
}

type QXmlNamePool_ITF interface {
	QXmlNamePool_PTR() *QXmlNamePool
}

func (p *QXmlNamePool) QXmlNamePool_PTR() *QXmlNamePool {
	return p
}

func (p *QXmlNamePool) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlNamePool) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlNamePool(ptr QXmlNamePool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePool_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamePoolFromPointer(ptr unsafe.Pointer) *QXmlNamePool {
	var n = new(QXmlNamePool)
	n.SetPointer(ptr)
	return n
}

func newQXmlNamePoolFromPointer(ptr unsafe.Pointer) *QXmlNamePool {
	var n = NewQXmlNamePoolFromPointer(ptr)
	return n
}

func NewQXmlNamePool() *QXmlNamePool {
	defer qt.Recovering("QXmlNamePool::QXmlNamePool")

	return newQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool())
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {
	defer qt.Recovering("QXmlNamePool::QXmlNamePool")

	return newQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool2(PointerFromQXmlNamePool(other)))
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	defer qt.Recovering("QXmlNamePool::~QXmlNamePool")

	if ptr.Pointer() != nil {
		C.QXmlNamePool_DestroyQXmlNamePool(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  = QXmlNodeModelIndex__DocumentOrder(1)
)

//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int64

const (
	QXmlNodeModelIndex__Attribute             = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  = QXmlNodeModelIndex__NodeKind(64)
)

type QXmlNodeModelIndex struct {
	ptr unsafe.Pointer
}

type QXmlNodeModelIndex_ITF interface {
	QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex
}

func (p *QXmlNodeModelIndex) QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex {
	return p
}

func (p *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlNodeModelIndex) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlNodeModelIndex(ptr QXmlNodeModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNodeModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) *QXmlNodeModelIndex {
	var n = new(QXmlNodeModelIndex)
	n.SetPointer(ptr)
	return n
}

func newQXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) *QXmlNodeModelIndex {
	var n = NewQXmlNodeModelIndexFromPointer(ptr)
	return n
}

func (ptr *QXmlNodeModelIndex) DestroyQXmlNodeModelIndex() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	defer qt.Recovering("QXmlNodeModelIndex::QXmlNodeModelIndex")

	return newQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex())
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer qt.Recovering("QXmlNodeModelIndex::QXmlNodeModelIndex")

	return newQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(PointerFromQXmlNodeModelIndex(other)))
}

func (ptr *QXmlNodeModelIndex) AdditionalData() int64 {
	defer qt.Recovering("QXmlNodeModelIndex::additionalData")

	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_AdditionalData(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) Data() int64 {
	defer qt.Recovering("QXmlNodeModelIndex::data")

	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_Data(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) InternalPointer() unsafe.Pointer {
	defer qt.Recovering("QXmlNodeModelIndex::internalPointer")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlNodeModelIndex_InternalPointer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {
	defer qt.Recovering("QXmlNodeModelIndex::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlNodeModelIndex_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {
	defer qt.Recovering("QXmlNodeModelIndex::model")

	if ptr.Pointer() != nil {
		return NewQAbstractXmlNodeModelFromPointer(C.QXmlNodeModelIndex_Model(ptr.Pointer()))
	}
	return nil
}

//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int64

const (
	QXmlQuery__XQuery10                              = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               = QXmlQuery__QueryLanguage(4096)
)

type QXmlQuery struct {
	ptr unsafe.Pointer
}

type QXmlQuery_ITF interface {
	QXmlQuery_PTR() *QXmlQuery
}

func (p *QXmlQuery) QXmlQuery_PTR() *QXmlQuery {
	return p
}

func (p *QXmlQuery) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlQuery) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlQuery(ptr QXmlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlQuery_PTR().Pointer()
	}
	return nil
}

func NewQXmlQueryFromPointer(ptr unsafe.Pointer) *QXmlQuery {
	var n = new(QXmlQuery)
	n.SetPointer(ptr)
	return n
}

func newQXmlQueryFromPointer(ptr unsafe.Pointer) *QXmlQuery {
	var n = NewQXmlQueryFromPointer(ptr)
	return n
}

func NewQXmlQuery() *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return newQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery())
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return newQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery4(C.int(queryLanguage), PointerFromQXmlNamePool(np)))
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return newQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery3(PointerFromQXmlNamePool(np)))
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {
	defer qt.Recovering("QXmlQuery::QXmlQuery")

	return newQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery2(PointerFromQXmlQuery(other)))
}

func (ptr *QXmlQuery) BindVariable4(localName string, device core.QIODevice_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_BindVariable4(ptr.Pointer(), localNameC, core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable2(localName string, value QXmlItem_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_BindVariable2(ptr.Pointer(), localNameC, PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQuery_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_BindVariable6(ptr.Pointer(), localNameC, PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, device core.QIODevice_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable5(name QXmlName_ITF, query QXmlQuery_ITF) {
	defer qt.Recovering("QXmlQuery::bindVariable")

	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) EvaluateTo2(callback QAbstractXmlReceiver_ITF) bool {
	defer qt.Recovering("QXmlQuery::evaluateTo")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_EvaluateTo2(ptr.Pointer(), PointerFromQAbstractXmlReceiver(callback)) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo4(target core.QIODevice_ITF) bool {
	defer qt.Recovering("QXmlQuery::evaluateTo")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_EvaluateTo4(ptr.Pointer(), core.PointerFromQIODevice(target)) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo5(output string) bool {
	defer qt.Recovering("QXmlQuery::evaluateTo")

	if ptr.Pointer() != nil {
		var outputC = C.CString(output)
		defer C.free(unsafe.Pointer(outputC))
		return C.QXmlQuery_EvaluateTo5(ptr.Pointer(), outputC) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo3(target []string) bool {
	defer qt.Recovering("QXmlQuery::evaluateTo")

	if ptr.Pointer() != nil {
		var targetC = C.CString(strings.Join(target, "|"))
		defer C.free(unsafe.Pointer(targetC))
		return C.QXmlQuery_EvaluateTo3(ptr.Pointer(), targetC) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo(result QXmlResultItems_ITF) {
	defer qt.Recovering("QXmlQuery::evaluateTo")

	if ptr.Pointer() != nil {
		C.QXmlQuery_EvaluateTo(ptr.Pointer(), PointerFromQXmlResultItems(result))
	}
}

func (ptr *QXmlQuery) IsValid() bool {
	defer qt.Recovering("QXmlQuery::isValid")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {
	defer qt.Recovering("QXmlQuery::messageHandler")

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlQuery_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) NamePool() *QXmlNamePool {
	defer qt.Recovering("QXmlQuery::namePool")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlQuery_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QXmlQuery::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlQuery_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {
	defer qt.Recovering("QXmlQuery::queryLanguage")

	if ptr.Pointer() != nil {
		return QXmlQuery__QueryLanguage(C.QXmlQuery_QueryLanguage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODevice_ITF) bool {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus3(ptr.Pointer(), core.PointerFromQIODevice(document)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		var focusC = C.CString(focus)
		defer C.free(unsafe.Pointer(focusC))
		return C.QXmlQuery_SetFocus4(ptr.Pointer(), focusC) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus2(ptr.Pointer(), core.PointerFromQUrl(documentURI)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {
	defer qt.Recovering("QXmlQuery::setFocus")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(ptr.Pointer(), PointerFromQXmlItem(item))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {
	defer qt.Recovering("QXmlQuery::setInitialTemplateName")

	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_SetInitialTemplateName2(ptr.Pointer(), localNameC)
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {
	defer qt.Recovering("QXmlQuery::setInitialTemplateName")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandler_ITF) {
	defer qt.Recovering("QXmlQuery::setMessageHandler")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(aMessageHandler))
	}
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QXmlQuery::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(newManager))
	}
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODevice_ITF, documentURI core.QUrl_ITF) {
	defer qt.Recovering("QXmlQuery::setQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery(ptr.Pointer(), core.PointerFromQIODevice(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery2(sourceCode string, documentURI core.QUrl_ITF) {
	defer qt.Recovering("QXmlQuery::setQuery")

	if ptr.Pointer() != nil {
		var sourceCodeC = C.CString(sourceCode)
		defer C.free(unsafe.Pointer(sourceCodeC))
		C.QXmlQuery_SetQuery2(ptr.Pointer(), sourceCodeC, core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery3(queryURI core.QUrl_ITF, baseURI core.QUrl_ITF) {
	defer qt.Recovering("QXmlQuery::setQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery3(ptr.Pointer(), core.PointerFromQUrl(queryURI), core.PointerFromQUrl(baseURI))
	}
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer qt.Recovering("QXmlQuery::setUriResolver")

	if ptr.Pointer() != nil {
		C.QXmlQuery_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {
	defer qt.Recovering("QXmlQuery::uriResolver")

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlQuery_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {
	defer qt.Recovering("QXmlQuery::~QXmlQuery")

	if ptr.Pointer() != nil {
		C.QXmlQuery_DestroyQXmlQuery(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlResultItems struct {
	ptr unsafe.Pointer
}

type QXmlResultItems_ITF interface {
	QXmlResultItems_PTR() *QXmlResultItems
}

func (p *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return p
}

func (p *QXmlResultItems) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlResultItems) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlResultItems(ptr QXmlResultItems_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlResultItems_PTR().Pointer()
	}
	return nil
}

func NewQXmlResultItemsFromPointer(ptr unsafe.Pointer) *QXmlResultItems {
	var n = new(QXmlResultItems)
	n.SetPointer(ptr)
	return n
}

func newQXmlResultItemsFromPointer(ptr unsafe.Pointer) *QXmlResultItems {
	var n = NewQXmlResultItemsFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlResultItems_") {
		n.SetObjectNameAbs("QXmlResultItems_" + qt.Identifier())
	}
	return n
}

func NewQXmlResultItems() *QXmlResultItems {
	defer qt.Recovering("QXmlResultItems::QXmlResultItems")

	return newQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
}

func (ptr *QXmlResultItems) Current() *QXmlItem {
	defer qt.Recovering("QXmlResultItems::current")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlItemFromPointer(C.QXmlResultItems_Current(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlResultItems) HasError() bool {
	defer qt.Recovering("QXmlResultItems::hasError")

	if ptr.Pointer() != nil {
		return C.QXmlResultItems_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlResultItems) Next() *QXmlItem {
	defer qt.Recovering("QXmlResultItems::next")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlItemFromPointer(C.QXmlResultItems_Next(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	defer qt.Recovering("QXmlResultItems::~QXmlResultItems")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlResultItems) ObjectNameAbs() string {
	defer qt.Recovering("QXmlResultItems::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlResultItems_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlResultItems) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlResultItems::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QXmlResultItems_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

type QXmlSchema struct {
	ptr unsafe.Pointer
}

type QXmlSchema_ITF interface {
	QXmlSchema_PTR() *QXmlSchema
}

func (p *QXmlSchema) QXmlSchema_PTR() *QXmlSchema {
	return p
}

func (p *QXmlSchema) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlSchema) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlSchema(ptr QXmlSchema_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchema_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaFromPointer(ptr unsafe.Pointer) *QXmlSchema {
	var n = new(QXmlSchema)
	n.SetPointer(ptr)
	return n
}

func newQXmlSchemaFromPointer(ptr unsafe.Pointer) *QXmlSchema {
	var n = NewQXmlSchemaFromPointer(ptr)
	return n
}

func NewQXmlSchema() *QXmlSchema {
	defer qt.Recovering("QXmlSchema::QXmlSchema")

	return newQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema())
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {
	defer qt.Recovering("QXmlSchema::QXmlSchema")

	return newQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema2(PointerFromQXmlSchema(other)))
}

func (ptr *QXmlSchema) DocumentUri() *core.QUrl {
	defer qt.Recovering("QXmlSchema::documentUri")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QXmlSchema_DocumentUri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) IsValid() bool {
	defer qt.Recovering("QXmlSchema::isValid")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchema::load")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load3(data string, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchema::load")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		return C.QXmlSchema_Load3(ptr.Pointer(), dataC, core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchema::load")

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {
	defer qt.Recovering("QXmlSchema::messageHandler")

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlSchema_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) NamePool() *QXmlNamePool {
	defer qt.Recovering("QXmlSchema::namePool")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlSchema_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QXmlSchema::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlSchema_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	defer qt.Recovering("QXmlSchema::setMessageHandler")

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QXmlSchema::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer qt.Recovering("QXmlSchema::setUriResolver")

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {
	defer qt.Recovering("QXmlSchema::uriResolver")

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlSchema_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {
	defer qt.Recovering("QXmlSchema::~QXmlSchema")

	if ptr.Pointer() != nil {
		C.QXmlSchema_DestroyQXmlSchema(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlSchemaValidator struct {
	ptr unsafe.Pointer
}

type QXmlSchemaValidator_ITF interface {
	QXmlSchemaValidator_PTR() *QXmlSchemaValidator
}

func (p *QXmlSchemaValidator) QXmlSchemaValidator_PTR() *QXmlSchemaValidator {
	return p
}

func (p *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QXmlSchemaValidator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQXmlSchemaValidator(ptr QXmlSchemaValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaValidator_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) *QXmlSchemaValidator {
	var n = new(QXmlSchemaValidator)
	n.SetPointer(ptr)
	return n
}

func newQXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) *QXmlSchemaValidator {
	var n = NewQXmlSchemaValidatorFromPointer(ptr)
	return n
}

func NewQXmlSchemaValidator() *QXmlSchemaValidator {
	defer qt.Recovering("QXmlSchemaValidator::QXmlSchemaValidator")

	return newQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator())
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {
	defer qt.Recovering("QXmlSchemaValidator::QXmlSchemaValidator")

	return newQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(PointerFromQXmlSchema(schema)))
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {
	defer qt.Recovering("QXmlSchemaValidator::messageHandler")

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlSchemaValidator_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NamePool() *QXmlNamePool {
	defer qt.Recovering("QXmlSchemaValidator::namePool")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlSchemaValidator_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QXmlSchemaValidator::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlSchemaValidator_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Schema() *QXmlSchema {
	defer qt.Recovering("QXmlSchemaValidator::schema")

	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlSchemaFromPointer(C.QXmlSchemaValidator_Schema(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setMessageHandler")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchema_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setSchema")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetSchema(ptr.Pointer(), PointerFromQXmlSchema(schema))
	}
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer qt.Recovering("QXmlSchemaValidator::setUriResolver")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {
	defer qt.Recovering("QXmlSchemaValidator::uriResolver")

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlSchemaValidator_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchemaValidator::validate")

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate3(data string, documentUri core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchemaValidator::validate")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		return C.QXmlSchemaValidator_Validate3(ptr.Pointer(), dataC, core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {
	defer qt.Recovering("QXmlSchemaValidator::validate")

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
	defer qt.Recovering("QXmlSchemaValidator::~QXmlSchemaValidator")

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_DestroyQXmlSchemaValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlSerializer struct {
	QAbstractXmlReceiver
}

type QXmlSerializer_ITF interface {
	QAbstractXmlReceiver_ITF
	QXmlSerializer_PTR() *QXmlSerializer
}

func (p *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return p
}

func (p *QXmlSerializer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func (p *QXmlSerializer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractXmlReceiver_PTR().SetPointer(ptr)
	}
}

func PointerFromQXmlSerializer(ptr QXmlSerializer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func NewQXmlSerializerFromPointer(ptr unsafe.Pointer) *QXmlSerializer {
	var n = new(QXmlSerializer)
	n.SetPointer(ptr)
	return n
}

func newQXmlSerializerFromPointer(ptr unsafe.Pointer) *QXmlSerializer {
	var n = NewQXmlSerializerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlSerializer_") {
		n.SetObjectNameAbs("QXmlSerializer_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlSerializer) DestroyQXmlSerializer() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	defer qt.Recovering("QXmlSerializer::QXmlSerializer")

	return newQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

//export callbackQXmlSerializer_AtomicValue
func callbackQXmlSerializer_AtomicValue(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::atomicValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AtomicValueDefault(core.NewQVariantFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAtomicValue(f func(value *core.QVariant)) {
	defer qt.Recovering("connect QXmlSerializer::atomicValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "atomicValue", f)
	}
}

func (ptr *QXmlSerializer) DisconnectAtomicValue() {
	defer qt.Recovering("disconnect QXmlSerializer::atomicValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "atomicValue")
	}
}

func (ptr *QXmlSerializer) AtomicValue(value core.QVariant_ITF) {
	defer qt.Recovering("QXmlSerializer::atomicValue")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QXmlSerializer) AtomicValueDefault(value core.QVariant_ITF) {
	defer qt.Recovering("QXmlSerializer::atomicValue")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_AtomicValueDefault(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQXmlSerializer_Attribute
func callbackQXmlSerializer_Attribute(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::attribute")

	if signal := qt.GetSignal(C.GoString(ptrName), "attribute"); signal != nil {
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AttributeDefault(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlSerializer::attribute")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "attribute", f)
	}
}

func (ptr *QXmlSerializer) DisconnectAttribute() {
	defer qt.Recovering("disconnect QXmlSerializer::attribute")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "attribute")
	}
}

func (ptr *QXmlSerializer) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::attribute")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) AttributeDefault(name QXmlName_ITF, value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::attribute")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_AttributeDefault(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlSerializer_Characters
func callbackQXmlSerializer_Characters(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::characters")

	if signal := qt.GetSignal(C.GoString(ptrName), "characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {
	defer qt.Recovering("connect QXmlSerializer::characters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "characters", f)
	}
}

func (ptr *QXmlSerializer) DisconnectCharacters() {
	defer qt.Recovering("disconnect QXmlSerializer::characters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "characters")
	}
}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::characters")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) CharactersDefault(value core.QStringRef_ITF) {
	defer qt.Recovering("QXmlSerializer::characters")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlSerializer_Comment
func callbackQXmlSerializer_Comment(ptr unsafe.Pointer, ptrName *C.char, value *C.char) {
	defer qt.Recovering("callback QXmlSerializer::comment")

	if signal := qt.GetSignal(C.GoString(ptrName), "comment"); signal != nil {
		signal.(func(string))(C.GoString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CommentDefault(C.GoString(value))
	}
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {
	defer qt.Recovering("connect QXmlSerializer::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "comment", f)
	}
}

func (ptr *QXmlSerializer) DisconnectComment() {
	defer qt.Recovering("disconnect QXmlSerializer::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "comment")
	}
}

func (ptr *QXmlSerializer) Comment(value string) {
	defer qt.Recovering("QXmlSerializer::comment")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_Comment(ptr.Pointer(), valueC)
	}
}

func (ptr *QXmlSerializer) CommentDefault(value string) {
	defer qt.Recovering("QXmlSerializer::comment")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_CommentDefault(ptr.Pointer(), valueC)
	}
}

//export callbackQXmlSerializer_EndDocument
func callbackQXmlSerializer_EndDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "endDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndDocument() {
	defer qt.Recovering("disconnect QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endDocument")
	}
}

func (ptr *QXmlSerializer) EndDocument() {
	defer qt.Recovering("QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndDocumentDefault() {
	defer qt.Recovering("QXmlSerializer::endDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_EndElement
func callbackQXmlSerializer_EndElement(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "endElement"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endElement", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndElement() {
	defer qt.Recovering("disconnect QXmlSerializer::endElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endElement")
	}
}

func (ptr *QXmlSerializer) EndElement() {
	defer qt.Recovering("QXmlSerializer::endElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndElementDefault() {
	defer qt.Recovering("QXmlSerializer::endElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElementDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	defer qt.Recovering("QXmlSerializer::codec")

	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSerializer_EndOfSequence
func callbackQXmlSerializer_EndOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::endOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "endOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {
	defer qt.Recovering("connect QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "endOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndOfSequence() {
	defer qt.Recovering("disconnect QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "endOfSequence")
	}
}

func (ptr *QXmlSerializer) EndOfSequence() {
	defer qt.Recovering("QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndOfSequenceDefault() {
	defer qt.Recovering("QXmlSerializer::endOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequenceDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_NamespaceBinding
func callbackQXmlSerializer_NamespaceBinding(ptr unsafe.Pointer, ptrName *C.char, nb unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::namespaceBinding")

	if signal := qt.GetSignal(C.GoString(ptrName), "namespaceBinding"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(nb))
	} else {
		NewQXmlSerializerFromPointer(ptr).NamespaceBindingDefault(NewQXmlNameFromPointer(nb))
	}
}

func (ptr *QXmlSerializer) ConnectNamespaceBinding(f func(nb *QXmlName)) {
	defer qt.Recovering("connect QXmlSerializer::namespaceBinding")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "namespaceBinding", f)
	}
}

func (ptr *QXmlSerializer) DisconnectNamespaceBinding() {
	defer qt.Recovering("disconnect QXmlSerializer::namespaceBinding")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "namespaceBinding")
	}
}

func (ptr *QXmlSerializer) NamespaceBinding(nb QXmlName_ITF) {
	defer qt.Recovering("QXmlSerializer::namespaceBinding")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) NamespaceBindingDefault(nb QXmlName_ITF) {
	defer qt.Recovering("QXmlSerializer::namespaceBinding")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBindingDefault(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	defer qt.Recovering("QXmlSerializer::outputDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QXmlSerializer_OutputDevice(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSerializer_ProcessingInstruction
func callbackQXmlSerializer_ProcessingInstruction(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer, value *C.char) {
	defer qt.Recovering("callback QXmlSerializer::processingInstruction")

	if signal := qt.GetSignal(C.GoString(ptrName), "processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(name), C.GoString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).ProcessingInstructionDefault(NewQXmlNameFromPointer(name), C.GoString(value))
	}
}

func (ptr *QXmlSerializer) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {
	defer qt.Recovering("connect QXmlSerializer::processingInstruction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "processingInstruction", f)
	}
}

func (ptr *QXmlSerializer) DisconnectProcessingInstruction() {
	defer qt.Recovering("disconnect QXmlSerializer::processingInstruction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "processingInstruction")
	}
}

func (ptr *QXmlSerializer) ProcessingInstruction(name QXmlName_ITF, value string) {
	defer qt.Recovering("QXmlSerializer::processingInstruction")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlSerializer) ProcessingInstructionDefault(name QXmlName_ITF, value string) {
	defer qt.Recovering("QXmlSerializer::processingInstruction")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_ProcessingInstructionDefault(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {
	defer qt.Recovering("QXmlSerializer::setCodec")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(ptr.Pointer(), core.PointerFromQTextCodec(outputCodec))
	}
}

//export callbackQXmlSerializer_StartDocument
func callbackQXmlSerializer_StartDocument(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::startDocument")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {
	defer qt.Recovering("connect QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartDocument() {
	defer qt.Recovering("disconnect QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startDocument")
	}
}

func (ptr *QXmlSerializer) StartDocument() {
	defer qt.Recovering("QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartDocumentDefault() {
	defer qt.Recovering("QXmlSerializer::startDocument")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_StartElement
func callbackQXmlSerializer_StartElement(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSerializer::startElement")

	if signal := qt.GetSignal(C.GoString(ptrName), "startElement"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	} else {
		NewQXmlSerializerFromPointer(ptr).StartElementDefault(NewQXmlNameFromPointer(name))
	}
}

func (ptr *QXmlSerializer) ConnectStartElement(f func(name *QXmlName)) {
	defer qt.Recovering("connect QXmlSerializer::startElement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startElement", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartElement() {
	defer qt.Recovering("disconnect QXmlSerializer::startElement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startElement")
	}
}

func (ptr *QXmlSerializer) StartElement(name QXmlName_ITF) {
	defer qt.Recovering("QXmlSerializer::startElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlSerializer) StartElementDefault(name QXmlName_ITF) {
	defer qt.Recovering("QXmlSerializer::startElement")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElementDefault(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQXmlSerializer_StartOfSequence
func callbackQXmlSerializer_StartOfSequence(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlSerializer::startOfSequence")

	if signal := qt.GetSignal(C.GoString(ptrName), "startOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {
	defer qt.Recovering("connect QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "startOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartOfSequence() {
	defer qt.Recovering("disconnect QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "startOfSequence")
	}
}

func (ptr *QXmlSerializer) StartOfSequence() {
	defer qt.Recovering("QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartOfSequenceDefault() {
	defer qt.Recovering("QXmlSerializer::startOfSequence")

	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) ObjectNameAbs() string {
	defer qt.Recovering("QXmlSerializer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSerializer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlSerializer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlSerializer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QXmlSerializer_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}
