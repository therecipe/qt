// +build !minimal

package xmlpatterns

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "xmlpatterns.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtXmlPatterns_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandler_ITF interface {
	core.QObject_ITF
	QAbstractMessageHandler_PTR() *QAbstractMessageHandler
}

func (ptr *QAbstractMessageHandler) QAbstractMessageHandler_PTR() *QAbstractMessageHandler {
	return ptr
}

func (ptr *QAbstractMessageHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractMessageHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//export callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler
func callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::~QAbstractMessageHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DestroyQAbstractMessageHandlerDefault()
	}
}

func (ptr *QAbstractMessageHandler) ConnectDestroyQAbstractMessageHandler(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::~QAbstractMessageHandler", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectDestroyQAbstractMessageHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::~QAbstractMessageHandler")
	}
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandlerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractMessageHandler_TimerEvent
func callbackQAbstractMessageHandler_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::timerEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::timerEvent")
	}
}

func (ptr *QAbstractMessageHandler) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractMessageHandler_ChildEvent
func callbackQAbstractMessageHandler_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::childEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::childEvent")
	}
}

func (ptr *QAbstractMessageHandler) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractMessageHandler_ConnectNotify
func callbackQAbstractMessageHandler_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::connectNotify", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::connectNotify")
	}
}

func (ptr *QAbstractMessageHandler) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_CustomEvent
func callbackQAbstractMessageHandler_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::customEvent", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::customEvent")
	}
}

func (ptr *QAbstractMessageHandler) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractMessageHandler) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractMessageHandler_DeleteLater
func callbackQAbstractMessageHandler_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractMessageHandler) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::deleteLater", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::deleteLater")
	}
}

func (ptr *QAbstractMessageHandler) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractMessageHandler) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractMessageHandler_DisconnectNotify
func callbackQAbstractMessageHandler_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::disconnectNotify", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::disconnectNotify")
	}
}

func (ptr *QAbstractMessageHandler) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractMessageHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_Event
func callbackQAbstractMessageHandler_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractMessageHandler) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::event", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::event")
	}
}

func (ptr *QAbstractMessageHandler) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractMessageHandler) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_EventFilter
func callbackQAbstractMessageHandler_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractMessageHandler) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::eventFilter", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::eventFilter")
	}
}

func (ptr *QAbstractMessageHandler) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractMessageHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractMessageHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_MetaObject
func callbackQAbstractMessageHandler_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractMessageHandler::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractMessageHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractMessageHandler) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::metaObject", f)
	}
}

func (ptr *QAbstractMessageHandler) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractMessageHandler::metaObject")
	}
}

func (ptr *QAbstractMessageHandler) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractMessageHandler_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractMessageHandler) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractUriResolver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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
func NewQAbstractUriResolver(parent core.QObject_ITF) *QAbstractUriResolver {
	var tmpValue = NewQAbstractUriResolverFromPointer(C.QAbstractUriResolver_NewQAbstractUriResolver(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractUriResolver_Resolve
func callbackQAbstractUriResolver_Resolve(ptr unsafe.Pointer, relative unsafe.Pointer, baseURI unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::resolve"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*core.QUrl, *core.QUrl) *core.QUrl)(core.NewQUrlFromPointer(relative), core.NewQUrlFromPointer(baseURI)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractUriResolver) ConnectResolve(f func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::resolve", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectResolve() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::resolve")
	}
}

func (ptr *QAbstractUriResolver) Resolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QAbstractUriResolver_Resolve(ptr.Pointer(), core.PointerFromQUrl(relative), core.PointerFromQUrl(baseURI)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractUriResolver_DestroyQAbstractUriResolver
func callbackQAbstractUriResolver_DestroyQAbstractUriResolver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::~QAbstractUriResolver"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DestroyQAbstractUriResolverDefault()
	}
}

func (ptr *QAbstractUriResolver) ConnectDestroyQAbstractUriResolver(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::~QAbstractUriResolver", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectDestroyQAbstractUriResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::~QAbstractUriResolver")
	}
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolverDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolverDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractUriResolver_TimerEvent
func callbackQAbstractUriResolver_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::timerEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::timerEvent")
	}
}

func (ptr *QAbstractUriResolver) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractUriResolver_ChildEvent
func callbackQAbstractUriResolver_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::childEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::childEvent")
	}
}

func (ptr *QAbstractUriResolver) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractUriResolver_ConnectNotify
func callbackQAbstractUriResolver_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::connectNotify", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::connectNotify")
	}
}

func (ptr *QAbstractUriResolver) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_CustomEvent
func callbackQAbstractUriResolver_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::customEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::customEvent")
	}
}

func (ptr *QAbstractUriResolver) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractUriResolver_DeleteLater
func callbackQAbstractUriResolver_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractUriResolver) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::deleteLater", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::deleteLater")
	}
}

func (ptr *QAbstractUriResolver) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractUriResolver_DisconnectNotify
func callbackQAbstractUriResolver_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::disconnectNotify", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::disconnectNotify")
	}
}

func (ptr *QAbstractUriResolver) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractUriResolver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_Event
func callbackQAbstractUriResolver_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractUriResolver) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::event", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::event")
	}
}

func (ptr *QAbstractUriResolver) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractUriResolver) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_EventFilter
func callbackQAbstractUriResolver_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractUriResolver) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::eventFilter", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::eventFilter")
	}
}

func (ptr *QAbstractUriResolver) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractUriResolver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractUriResolver_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_MetaObject
func callbackQAbstractUriResolver_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractUriResolver::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractUriResolverFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractUriResolver) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::metaObject", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractUriResolver::metaObject")
	}
}

func (ptr *QAbstractUriResolver) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractUriResolver_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractUriResolver) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractUriResolver_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModel_ITF interface {
	core.QSharedData_ITF
	QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel
}

func (ptr *QAbstractXmlNodeModel) QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel {
	return ptr
}

func (ptr *QAbstractXmlNodeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedData_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSharedData_PTR().SetPointer(p)
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

//go:generate stringer -type=QAbstractXmlNodeModel__SimpleAxis
//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int64

const (
	QAbstractXmlNodeModel__Parent          QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(3)
)

func (ptr *QAbstractXmlNodeModel) ConnectAttributes(f func(element *QXmlNodeModelIndex) []*QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::attributes", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectAttributes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::attributes")
	}
}

func (ptr *QAbstractXmlNodeModel) Attributes(element QXmlNodeModelIndex_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			var out = make([]*QXmlNodeModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractXmlNodeModelFromPointer(l.data).attributes_atList(i)
			}
			return out
		}(C.QAbstractXmlNodeModel_Attributes(ptr.Pointer(), PointerFromQXmlNodeModelIndex(element)))
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_BaseUri
func callbackQAbstractXmlNodeModel_BaseUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::baseUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractXmlNodeModel) ConnectBaseUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::baseUri", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectBaseUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::baseUri")
	}
}

func (ptr *QAbstractXmlNodeModel) BaseUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_CompareOrder
func callbackQAbstractXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::compareOrder"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.longlong(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::compareOrder", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectCompareOrder() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::compareOrder")
	}
}

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QAbstractXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) CreateIndex(data int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex(ptr.Pointer(), C.longlong(data)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex3(data int64, additionalData int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex3(ptr.Pointer(), C.longlong(data), C.longlong(additionalData)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex2(pointer unsafe.Pointer, additionalData int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex2(ptr.Pointer(), pointer, C.longlong(additionalData)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_DocumentUri
func callbackQAbstractXmlNodeModel_DocumentUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::documentUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::documentUri", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectDocumentUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::documentUri")
	}
}

func (ptr *QAbstractXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_ElementById
func callbackQAbstractXmlNodeModel_ElementById(ptr unsafe.Pointer, id unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlName) *QXmlNodeModelIndex)(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::elementById", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectElementById() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::elementById")
	}
}

func (ptr *QAbstractXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_ElementById(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Kind
func callbackQAbstractXmlNodeModel_Kind(ptr unsafe.Pointer, ni unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::kind"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.longlong(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::kind", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectKind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::kind")
	}
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QAbstractXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) ConnectNamespaceBindings(f func(n *QXmlNodeModelIndex) []*QXmlName) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::namespaceBindings", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNamespaceBindings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::namespaceBindings")
	}
}

func (ptr *QAbstractXmlNodeModel) NamespaceBindings(n QXmlNodeModelIndex_ITF) []*QXmlName {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlName {
			var out = make([]*QXmlName, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractXmlNodeModelFromPointer(l.data).namespaceBindings_atList(i)
			}
			return out
		}(C.QAbstractXmlNodeModel_NamespaceBindings(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Name
func callbackQAbstractXmlNodeModel_Name(ptr unsafe.Pointer, ni unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::name"); signal != nil {
		return PointerFromQXmlName(signal.(func(*QXmlNodeModelIndex) *QXmlName)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQXmlName())
}

func (ptr *QAbstractXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::name", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::name")
	}
}

func (ptr *QAbstractXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNameFromPointer(C.QAbstractXmlNodeModel_Name(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_NextFromSimpleAxis
func callbackQAbstractXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, axis C.longlong, origin unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::nextFromSimpleAxis", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNextFromSimpleAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::nextFromSimpleAxis")
	}
}

func (ptr *QAbstractXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) ConnectNodesByIdref(f func(idref *QXmlName) []*QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::nodesByIdref", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNodesByIdref() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::nodesByIdref")
	}
}

func (ptr *QAbstractXmlNodeModel) NodesByIdref(idref QXmlName_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			var out = make([]*QXmlNodeModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractXmlNodeModelFromPointer(l.data).nodesByIdref_atList(i)
			}
			return out
		}(C.QAbstractXmlNodeModel_NodesByIdref(ptr.Pointer(), PointerFromQXmlName(idref)))
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Root
func callbackQAbstractXmlNodeModel_Root(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::root"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::root", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectRoot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::root")
	}
}

func (ptr *QAbstractXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) SourceLocation(index QXmlNodeModelIndex_ITF) *QSourceLocation {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSourceLocationFromPointer(C.QAbstractXmlNodeModel_SourceLocation(ptr.Pointer(), PointerFromQXmlNodeModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_StringValue
func callbackQAbstractXmlNodeModel_StringValue(ptr unsafe.Pointer, n unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::stringValue"); signal != nil {
		return C.CString(signal.(func(*QXmlNodeModelIndex) string)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return C.CString("")
}

func (ptr *QAbstractXmlNodeModel) ConnectStringValue(f func(n *QXmlNodeModelIndex) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::stringValue", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectStringValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::stringValue")
	}
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndex_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return ""
}

//export callbackQAbstractXmlNodeModel_TypedValue
func callbackQAbstractXmlNodeModel_TypedValue(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::typedValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*QXmlNodeModelIndex) *core.QVariant)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QAbstractXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::typedValue", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectTypedValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::typedValue")
	}
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel
func callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlNodeModel::~QAbstractXmlNodeModel"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractXmlNodeModelFromPointer(ptr).DestroyQAbstractXmlNodeModelDefault()
	}
}

func (ptr *QAbstractXmlNodeModel) ConnectDestroyQAbstractXmlNodeModel(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::~QAbstractXmlNodeModel", f)
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectDestroyQAbstractXmlNodeModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlNodeModel::~QAbstractXmlNodeModel")
	}
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModelDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModelDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlNodeModel) attributes_atList(i int) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_attributes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNameFromPointer(C.QAbstractXmlNodeModel_namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) nodesByIdref_atList(i int) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_nodesByIdref_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

type QAbstractXmlReceiver struct {
	ptr unsafe.Pointer
}

type QAbstractXmlReceiver_ITF interface {
	QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver
}

func (ptr *QAbstractXmlReceiver) QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver {
	return ptr
}

func (ptr *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractXmlReceiver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQAbstractXmlReceiver() *QAbstractXmlReceiver {
	return NewQAbstractXmlReceiverFromPointer(C.QAbstractXmlReceiver_NewQAbstractXmlReceiver())
}

//export callbackQAbstractXmlReceiver_AtomicValue
func callbackQAbstractXmlReceiver_AtomicValue(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectAtomicValue(f func(value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::atomicValue", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectAtomicValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::atomicValue")
	}
}

func (ptr *QAbstractXmlReceiver) AtomicValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQAbstractXmlReceiver_Attribute
func callbackQAbstractXmlReceiver_Attribute(ptr unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::attribute"); signal != nil {
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::attribute", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::attribute")
	}
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQAbstractXmlReceiver_Characters
func callbackQAbstractXmlReceiver_Characters(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectCharacters(f func(value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::characters", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::characters")
	}
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQAbstractXmlReceiver_Comment
func callbackQAbstractXmlReceiver_Comment(ptr unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::comment"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectComment(f func(value string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::comment", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::comment")
	}
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QAbstractXmlReceiver_Comment(ptr.Pointer(), valueC)
	}
}

//export callbackQAbstractXmlReceiver_EndDocument
func callbackQAbstractXmlReceiver_EndDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::endDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndDocument(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::endDocument", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::endDocument")
	}
}

func (ptr *QAbstractXmlReceiver) EndDocument() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndDocument(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_EndElement
func callbackQAbstractXmlReceiver_EndElement(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::endElement"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndElement(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::endElement", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::endElement")
	}
}

func (ptr *QAbstractXmlReceiver) EndElement() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndElement(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_EndOfSequence
func callbackQAbstractXmlReceiver_EndOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::endOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndOfSequence(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::endOfSequence", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::endOfSequence")
	}
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndOfSequence(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_NamespaceBinding
func callbackQAbstractXmlReceiver_NamespaceBinding(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::namespaceBinding"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectNamespaceBinding(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::namespaceBinding", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectNamespaceBinding() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::namespaceBinding")
	}
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQAbstractXmlReceiver_ProcessingInstruction
func callbackQAbstractXmlReceiver_ProcessingInstruction(ptr unsafe.Pointer, target unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(target), cGoUnpackString(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectProcessingInstruction(f func(target *QXmlName, value string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::processingInstruction", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::processingInstruction")
	}
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QAbstractXmlReceiver_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(target), valueC)
	}
}

//export callbackQAbstractXmlReceiver_StartDocument
func callbackQAbstractXmlReceiver_StartDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::startDocument"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartDocument(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::startDocument", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::startDocument")
	}
}

func (ptr *QAbstractXmlReceiver) StartDocument() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartDocument(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_StartElement
func callbackQAbstractXmlReceiver_StartElement(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::startElement"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartElement(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::startElement", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::startElement")
	}
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQAbstractXmlReceiver_StartOfSequence
func callbackQAbstractXmlReceiver_StartOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::startOfSequence"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartOfSequence(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::startOfSequence", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::startOfSequence")
	}
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartOfSequence(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_DestroyQAbstractXmlReceiver
func callbackQAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractXmlReceiver::~QAbstractXmlReceiver"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractXmlReceiverFromPointer(ptr).DestroyQAbstractXmlReceiverDefault()
	}
}

func (ptr *QAbstractXmlReceiver) ConnectDestroyQAbstractXmlReceiver(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::~QAbstractXmlReceiver", f)
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectDestroyQAbstractXmlReceiver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractXmlReceiver::~QAbstractXmlReceiver")
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiverDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiverDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QPatternist struct {
	ptr unsafe.Pointer
}

type QPatternist_ITF interface {
	QPatternist_PTR() *QPatternist
}

func (ptr *QPatternist) QPatternist_PTR() *QPatternist {
	return ptr
}

func (ptr *QPatternist) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPatternist) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPatternist(ptr QPatternist_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPatternist_PTR().Pointer()
	}
	return nil
}

func NewQPatternistFromPointer(ptr unsafe.Pointer) *QPatternist {
	var n = new(QPatternist)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPatternist) DestroyQPatternist() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QPatternist__DerivedIntegerLimitsUsage
//QPatternist::DerivedIntegerLimitsUsage
type QPatternist__DerivedIntegerLimitsUsage int64

const (
	QPatternist__None           QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(1)
	QPatternist__LimitUpwards   QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(2)
	QPatternist__LimitDownwards QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(4)
	QPatternist__LimitBoth      QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(QPatternist__LimitUpwards | QPatternist__LimitDownwards)
)

//go:generate stringer -type=QPatternist__TypeOfDerivedInteger
//QPatternist::TypeOfDerivedInteger
type QPatternist__TypeOfDerivedInteger int64

const (
	QPatternist__TypeByte               QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(0)
	QPatternist__TypeInt                QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(1)
	QPatternist__TypeLong               QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(2)
	QPatternist__TypeNegativeInteger    QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(3)
	QPatternist__TypeNonNegativeInteger QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(4)
	QPatternist__TypeNonPositiveInteger QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(5)
	QPatternist__TypePositiveInteger    QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(6)
	QPatternist__TypeShort              QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(7)
	QPatternist__TypeUnsignedByte       QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(8)
	QPatternist__TypeUnsignedInt        QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(9)
	QPatternist__TypeUnsignedLong       QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(10)
	QPatternist__TypeUnsignedShort      QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(11)
)

//go:generate stringer -type=QPatternist__TypeOfDerivedString
//QPatternist::TypeOfDerivedString
type QPatternist__TypeOfDerivedString int64

const (
	QPatternist__TypeString           QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(0)
	QPatternist__TypeNormalizedString QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(1)
	QPatternist__TypeToken            QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(2)
	QPatternist__TypeLanguage         QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(3)
	QPatternist__TypeNMTOKEN          QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(4)
	QPatternist__TypeName             QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(5)
	QPatternist__TypeNCName           QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(6)
	QPatternist__TypeID               QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(7)
	QPatternist__TypeIDREF            QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(8)
	QPatternist__TypeENTITY           QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(9)
)

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModel_ITF interface {
	QAbstractXmlNodeModel_ITF
	QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractXmlNodeModel_PTR().SetPointer(p)
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

//export callbackQSimpleXmlNodeModel_BaseUri
func callbackQSimpleXmlNodeModel_BaseUri(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::baseUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQUrl(NewQSimpleXmlNodeModelFromPointer(ptr).BaseUriDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) ConnectBaseUri(f func(node *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::baseUri", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectBaseUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::baseUri")
	}
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) BaseUriDefault(node QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUriDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_ElementById
func callbackQSimpleXmlNodeModel_ElementById(ptr unsafe.Pointer, id unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlName) *QXmlNodeModelIndex)(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).ElementByIdDefault(NewQXmlNameFromPointer(id)))
}

func (ptr *QSimpleXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::elementById", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectElementById() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::elementById")
	}
}

func (ptr *QSimpleXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementById(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ElementByIdDefault(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementByIdDefault(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QSimpleXmlNodeModel_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ConnectNamespaceBindings(f func(node *QXmlNodeModelIndex) []*QXmlName) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::namespaceBindings", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectNamespaceBindings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::namespaceBindings")
	}
}

func (ptr *QSimpleXmlNodeModel) NamespaceBindings(node QXmlNodeModelIndex_ITF) []*QXmlName {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlName {
			var out = make([]*QXmlName, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSimpleXmlNodeModelFromPointer(l.data).namespaceBindings_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NamespaceBindings(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NamespaceBindingsDefault(node QXmlNodeModelIndex_ITF) []*QXmlName {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlName {
			var out = make([]*QXmlName, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSimpleXmlNodeModelFromPointer(l.data).namespaceBindings_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NamespaceBindingsDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ConnectNodesByIdref(f func(idref *QXmlName) []*QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::nodesByIdref", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectNodesByIdref() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::nodesByIdref")
	}
}

func (ptr *QSimpleXmlNodeModel) NodesByIdref(idref QXmlName_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			var out = make([]*QXmlNodeModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSimpleXmlNodeModelFromPointer(l.data).nodesByIdref_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NodesByIdref(ptr.Pointer(), PointerFromQXmlName(idref)))
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NodesByIdrefDefault(idref QXmlName_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			var out = make([]*QXmlNodeModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSimpleXmlNodeModelFromPointer(l.data).nodesByIdref_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NodesByIdrefDefault(ptr.Pointer(), PointerFromQXmlName(idref)))
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_StringValue
func callbackQSimpleXmlNodeModel_StringValue(ptr unsafe.Pointer, node unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::stringValue"); signal != nil {
		return C.CString(signal.(func(*QXmlNodeModelIndex) string)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return C.CString(NewQSimpleXmlNodeModelFromPointer(ptr).StringValueDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) ConnectStringValue(f func(node *QXmlNodeModelIndex) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::stringValue", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectStringValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::stringValue")
	}
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSimpleXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) StringValueDefault(node QXmlNodeModelIndex_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSimpleXmlNodeModel_StringValueDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

//export callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel
func callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::~QSimpleXmlNodeModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSimpleXmlNodeModelFromPointer(ptr).DestroyQSimpleXmlNodeModelDefault()
	}
}

func (ptr *QSimpleXmlNodeModel) ConnectDestroyQSimpleXmlNodeModel(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::~QSimpleXmlNodeModel", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectDestroyQSimpleXmlNodeModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::~QSimpleXmlNodeModel")
	}
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModelDefault() {
	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModelDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSimpleXmlNodeModel) namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNameFromPointer(C.QSimpleXmlNodeModel_namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) nodesByIdref_atList(i int) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_nodesByIdref_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ConnectAttributes(f func(element *QXmlNodeModelIndex) []*QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::attributes", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectAttributes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::attributes")
	}
}

func (ptr *QSimpleXmlNodeModel) Attributes(element QXmlNodeModelIndex_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			var out = make([]*QXmlNodeModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSimpleXmlNodeModelFromPointer(l.data).attributes_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_Attributes(ptr.Pointer(), PointerFromQXmlNodeModelIndex(element)))
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_CompareOrder
func callbackQSimpleXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::compareOrder"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.longlong(0)
}

func (ptr *QSimpleXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::compareOrder", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectCompareOrder() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::compareOrder")
	}
}

func (ptr *QSimpleXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QSimpleXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

//export callbackQSimpleXmlNodeModel_DocumentUri
func callbackQSimpleXmlNodeModel_DocumentUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::documentUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QSimpleXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::documentUri", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectDocumentUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::documentUri")
	}
}

func (ptr *QSimpleXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Kind
func callbackQSimpleXmlNodeModel_Kind(ptr unsafe.Pointer, ni unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::kind"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.longlong(0)
}

func (ptr *QSimpleXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::kind", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectKind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::kind")
	}
}

func (ptr *QSimpleXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QSimpleXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

//export callbackQSimpleXmlNodeModel_Name
func callbackQSimpleXmlNodeModel_Name(ptr unsafe.Pointer, ni unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::name"); signal != nil {
		return PointerFromQXmlName(signal.(func(*QXmlNodeModelIndex) *QXmlName)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQXmlName())
}

func (ptr *QSimpleXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::name", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::name")
	}
}

func (ptr *QSimpleXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNameFromPointer(C.QSimpleXmlNodeModel_Name(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_NextFromSimpleAxis
func callbackQSimpleXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, axis C.longlong, origin unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QSimpleXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::nextFromSimpleAxis", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectNextFromSimpleAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::nextFromSimpleAxis")
	}
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Root
func callbackQSimpleXmlNodeModel_Root(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::root"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QSimpleXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::root", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectRoot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::root")
	}
}

func (ptr *QSimpleXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_TypedValue
func callbackQSimpleXmlNodeModel_TypedValue(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSimpleXmlNodeModel::typedValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*QXmlNodeModelIndex) *core.QVariant)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QSimpleXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::typedValue", f)
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectTypedValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSimpleXmlNodeModel::typedValue")
	}
}

func (ptr *QSimpleXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
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

func (ptr *QSourceLocation) QSourceLocation_PTR() *QSourceLocation {
	return ptr
}

func (ptr *QSourceLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSourceLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQSourceLocation() *QSourceLocation {
	var tmpValue = NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation())
	runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {
	var tmpValue = NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation2(PointerFromQSourceLocation(other)))
	runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {
	var tmpValue = NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation3(core.PointerFromQUrl(u), C.int(int32(l)), C.int(int32(c))))
	runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func (ptr *QSourceLocation) Column() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSourceLocation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSourceLocation) Line() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) SetColumn(newColumn int64) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetColumn(ptr.Pointer(), C.longlong(newColumn))
	}
}

func (ptr *QSourceLocation) SetLine(newLine int64) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetLine(ptr.Pointer(), C.longlong(newLine))
	}
}

func (ptr *QSourceLocation) SetUri(newUri core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetUri(ptr.Pointer(), core.PointerFromQUrl(newUri))
	}
}

func (ptr *QSourceLocation) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QSourceLocation_Uri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {
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

func (ptr *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return ptr
}

func (ptr *QXmlFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXmlSerializer_PTR().SetPointer(p)
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

func (ptr *QXmlFormatter) DestroyQXmlFormatter() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

//export callbackQXmlFormatter_AtomicValue
func callbackQXmlFormatter_AtomicValue(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).AtomicValueDefault(core.NewQVariantFromPointer(value))
	}
}

func (ptr *QXmlFormatter) ConnectAtomicValue(f func(value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::atomicValue", f)
	}
}

func (ptr *QXmlFormatter) DisconnectAtomicValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::atomicValue")
	}
}

func (ptr *QXmlFormatter) AtomicValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QXmlFormatter) AtomicValueDefault(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_AtomicValueDefault(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQXmlFormatter_Attribute
func callbackQXmlFormatter_Attribute(ptr unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::attribute"); signal != nil {
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).AttributeDefault(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlFormatter) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::attribute", f)
	}
}

func (ptr *QXmlFormatter) DisconnectAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::attribute")
	}
}

func (ptr *QXmlFormatter) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) AttributeDefault(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_AttributeDefault(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlFormatter_Characters
func callbackQXmlFormatter_Characters(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlFormatter) ConnectCharacters(f func(value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::characters", f)
	}
}

func (ptr *QXmlFormatter) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::characters")
	}
}

func (ptr *QXmlFormatter) Characters(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlFormatter) CharactersDefault(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlFormatter_Comment
func callbackQXmlFormatter_Comment(ptr unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::comment"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).CommentDefault(cGoUnpackString(value))
	}
}

func (ptr *QXmlFormatter) ConnectComment(f func(value string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::comment", f)
	}
}

func (ptr *QXmlFormatter) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::comment")
	}
}

func (ptr *QXmlFormatter) Comment(value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_Comment(ptr.Pointer(), valueC)
	}
}

func (ptr *QXmlFormatter) CommentDefault(value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_CommentDefault(ptr.Pointer(), valueC)
	}
}

//export callbackQXmlFormatter_EndDocument
func callbackQXmlFormatter_EndDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::endDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlFormatter) ConnectEndDocument(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::endDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::endDocument")
	}
}

func (ptr *QXmlFormatter) EndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndDocumentDefault() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_EndElement
func callbackQXmlFormatter_EndElement(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::endElement"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlFormatter) ConnectEndElement(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::endElement", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::endElement")
	}
}

func (ptr *QXmlFormatter) EndElement() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndElementDefault() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndElementDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_EndOfSequence
func callbackQXmlFormatter_EndOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::endOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlFormatter) ConnectEndOfSequence(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::endOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectEndOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::endOfSequence")
	}
}

func (ptr *QXmlFormatter) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) EndOfSequenceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_EndOfSequenceDefault(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlFormatter_IndentationDepth(ptr.Pointer())))
	}
	return 0
}

//export callbackQXmlFormatter_ProcessingInstruction
func callbackQXmlFormatter_ProcessingInstruction(ptr unsafe.Pointer, name unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	} else {
		NewQXmlFormatterFromPointer(ptr).ProcessingInstructionDefault(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	}
}

func (ptr *QXmlFormatter) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::processingInstruction", f)
	}
}

func (ptr *QXmlFormatter) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::processingInstruction")
	}
}

func (ptr *QXmlFormatter) ProcessingInstruction(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlFormatter) ProcessingInstructionDefault(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlFormatter_ProcessingInstructionDefault(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

//export callbackQXmlFormatter_StartDocument
func callbackQXmlFormatter_StartDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::startDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlFormatter) ConnectStartDocument(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::startDocument", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::startDocument")
	}
}

func (ptr *QXmlFormatter) StartDocument() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartDocumentDefault() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_StartElement
func callbackQXmlFormatter_StartElement(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::startElement"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	} else {
		NewQXmlFormatterFromPointer(ptr).StartElementDefault(NewQXmlNameFromPointer(name))
	}
}

func (ptr *QXmlFormatter) ConnectStartElement(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::startElement", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::startElement")
	}
}

func (ptr *QXmlFormatter) StartElement(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlFormatter) StartElementDefault(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartElementDefault(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQXmlFormatter_StartOfSequence
func callbackQXmlFormatter_StartOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::startOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlFormatterFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlFormatter) ConnectStartOfSequence(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::startOfSequence", f)
	}
}

func (ptr *QXmlFormatter) DisconnectStartOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::startOfSequence")
	}
}

func (ptr *QXmlFormatter) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlFormatter) StartOfSequenceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_StartOfSequenceDefault(ptr.Pointer())
	}
}

//export callbackQXmlFormatter_NamespaceBinding
func callbackQXmlFormatter_NamespaceBinding(ptr unsafe.Pointer, nb unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlFormatter::namespaceBinding"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(nb))
	} else {
		NewQXmlFormatterFromPointer(ptr).NamespaceBindingDefault(NewQXmlNameFromPointer(nb))
	}
}

func (ptr *QXmlFormatter) ConnectNamespaceBinding(f func(nb *QXmlName)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::namespaceBinding", f)
	}
}

func (ptr *QXmlFormatter) DisconnectNamespaceBinding() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlFormatter::namespaceBinding")
	}
}

func (ptr *QXmlFormatter) NamespaceBinding(nb QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlFormatter) NamespaceBindingDefault(nb QXmlName_ITF) {
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

func (ptr *QXmlItem) QXmlItem_PTR() *QXmlItem {
	return ptr
}

func (ptr *QXmlItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQXmlItem() *QXmlItem {
	var tmpValue = NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem())
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	var tmpValue = NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	var tmpValue = NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	var tmpValue = NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func (ptr *QXmlItem) IsAtomicValue() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsAtomicValue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNode() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) ToAtomicValue() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QXmlItem_ToAtomicValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) ToNodeModelIndex() *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QXmlItem_ToNodeModelIndex(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) DestroyQXmlItem() {
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

func (ptr *QXmlName) QXmlName_PTR() *QXmlName {
	return ptr
}

func (ptr *QXmlName) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlName) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QXmlName) DestroyQXmlName() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQXmlName() *QXmlName {
	var tmpValue = NewQXmlNameFromPointer(C.QXmlName_NewQXmlName())
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func NewQXmlName2(namePool QXmlNamePool_ITF, localName string, namespaceURI string, prefix string) *QXmlName {
	var localNameC = C.CString(localName)
	defer C.free(unsafe.Pointer(localNameC))
	var namespaceURIC = C.CString(namespaceURI)
	defer C.free(unsafe.Pointer(namespaceURIC))
	var prefixC = C.CString(prefix)
	defer C.free(unsafe.Pointer(prefixC))
	var tmpValue = NewQXmlNameFromPointer(C.QXmlName_NewQXmlName2(PointerFromQXmlNamePool(namePool), localNameC, namespaceURIC, prefixC))
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func QXmlName_FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {
	var clarkNameC = C.CString(clarkName)
	defer C.free(unsafe.Pointer(clarkNameC))
	var tmpValue = NewQXmlNameFromPointer(C.QXmlName_QXmlName_FromClarkName(clarkNameC, PointerFromQXmlNamePool(namePool)))
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func (ptr *QXmlName) FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {
	var clarkNameC = C.CString(clarkName)
	defer C.free(unsafe.Pointer(clarkNameC))
	var tmpValue = NewQXmlNameFromPointer(C.QXmlName_QXmlName_FromClarkName(clarkNameC, PointerFromQXmlNamePool(namePool)))
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func QXmlName_IsNCName(candidate string) bool {
	var candidateC = C.CString(candidate)
	defer C.free(unsafe.Pointer(candidateC))
	return C.QXmlName_QXmlName_IsNCName(candidateC) != 0
}

func (ptr *QXmlName) IsNCName(candidate string) bool {
	var candidateC = C.CString(candidate)
	defer C.free(unsafe.Pointer(candidateC))
	return C.QXmlName_QXmlName_IsNCName(candidateC) != 0
}

func (ptr *QXmlName) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlName_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_LocalName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_NamespaceUri(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_Prefix(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_ToClarkName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

type QXmlNamePool struct {
	ptr unsafe.Pointer
}

type QXmlNamePool_ITF interface {
	QXmlNamePool_PTR() *QXmlNamePool
}

func (ptr *QXmlNamePool) QXmlNamePool_PTR() *QXmlNamePool {
	return ptr
}

func (ptr *QXmlNamePool) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlNamePool) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQXmlNamePool() *QXmlNamePool {
	var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool())
	runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
	return tmpValue
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {
	var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool2(PointerFromQXmlNamePool(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
	return tmpValue
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	if ptr.Pointer() != nil {
		C.QXmlNamePool_DestroyQXmlNamePool(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlNodeModelIndex struct {
	ptr unsafe.Pointer
}

type QXmlNodeModelIndex_ITF interface {
	QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex
}

func (ptr *QXmlNodeModelIndex) QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex {
	return ptr
}

func (ptr *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QXmlNodeModelIndex) DestroyQXmlNodeModelIndex() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QXmlNodeModelIndex__DocumentOrder
//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(1)
)

//go:generate stringer -type=QXmlNodeModelIndex__NodeKind
//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int64

const (
	QXmlNodeModelIndex__Attribute             QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(64)
)

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex())
	runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
	return tmpValue
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	var tmpValue = NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(PointerFromQXmlNodeModelIndex(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
	return tmpValue
}

func (ptr *QXmlNodeModelIndex) AdditionalData() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_AdditionalData(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) Data() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_Data(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) InternalPointer() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlNodeModelIndex_InternalPointer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlNodeModelIndex_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {
	if ptr.Pointer() != nil {
		return NewQAbstractXmlNodeModelFromPointer(C.QXmlNodeModelIndex_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNameFromPointer(C.QXmlNodeModelIndex_namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

type QXmlQuery struct {
	ptr unsafe.Pointer
}

type QXmlQuery_ITF interface {
	QXmlQuery_PTR() *QXmlQuery
}

func (ptr *QXmlQuery) QXmlQuery_PTR() *QXmlQuery {
	return ptr
}

func (ptr *QXmlQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QXmlQuery__QueryLanguage
//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int64

const (
	QXmlQuery__XQuery10                              QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(4096)
)

func NewQXmlQuery() *QXmlQuery {
	var tmpValue = NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery())
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {
	var tmpValue = NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery4(C.longlong(queryLanguage), PointerFromQXmlNamePool(np)))
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {
	var tmpValue = NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery3(PointerFromQXmlNamePool(np)))
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {
	var tmpValue = NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery2(PointerFromQXmlQuery(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func (ptr *QXmlQuery) BindVariable4(localName string, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_BindVariable4(ptr.Pointer(), localNameC, core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable2(localName string, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_BindVariable2(ptr.Pointer(), localNameC, PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_BindVariable6(ptr.Pointer(), localNameC, PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable5(name QXmlName_ITF, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) EvaluateTo2(callback QAbstractXmlReceiver_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_EvaluateTo2(ptr.Pointer(), PointerFromQAbstractXmlReceiver(callback)) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo4(target core.QIODevice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_EvaluateTo4(ptr.Pointer(), core.PointerFromQIODevice(target)) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo5(output string) bool {
	if ptr.Pointer() != nil {
		var outputC = C.CString(output)
		defer C.free(unsafe.Pointer(outputC))
		return C.QXmlQuery_EvaluateTo5(ptr.Pointer(), outputC) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo3(target []string) bool {
	if ptr.Pointer() != nil {
		var targetC = C.CString(strings.Join(target, "|"))
		defer C.free(unsafe.Pointer(targetC))
		return C.QXmlQuery_EvaluateTo3(ptr.Pointer(), targetC) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo(result QXmlResultItems_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_EvaluateTo(ptr.Pointer(), PointerFromQXmlResultItems(result))
	}
}

func (ptr *QXmlQuery) InitialTemplateName() *QXmlName {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNameFromPointer(C.QXmlQuery_InitialTemplateName(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractMessageHandlerFromPointer(C.QXmlQuery_MessageHandler(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlQuery_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QXmlQuery_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {
	if ptr.Pointer() != nil {
		return QXmlQuery__QueryLanguage(C.QXmlQuery_QueryLanguage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODevice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus3(ptr.Pointer(), core.PointerFromQIODevice(document)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {
	if ptr.Pointer() != nil {
		var focusC = C.CString(focus)
		defer C.free(unsafe.Pointer(focusC))
		return C.QXmlQuery_SetFocus4(ptr.Pointer(), focusC) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus2(ptr.Pointer(), core.PointerFromQUrl(documentURI)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(ptr.Pointer(), PointerFromQXmlItem(item))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {
	if ptr.Pointer() != nil {
		var localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
		C.QXmlQuery_SetInitialTemplateName2(ptr.Pointer(), localNameC)
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(aMessageHandler))
	}
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(newManager))
	}
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODevice_ITF, documentURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery(ptr.Pointer(), core.PointerFromQIODevice(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery2(sourceCode string, documentURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var sourceCodeC = C.CString(sourceCode)
		defer C.free(unsafe.Pointer(sourceCodeC))
		C.QXmlQuery_SetQuery2(ptr.Pointer(), sourceCodeC, core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery3(queryURI core.QUrl_ITF, baseURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery3(ptr.Pointer(), core.PointerFromQUrl(queryURI), core.PointerFromQUrl(baseURI))
	}
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractUriResolverFromPointer(C.QXmlQuery_UriResolver(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {
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

func (ptr *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return ptr
}

func (ptr *QXmlResultItems) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlResultItems) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQXmlResultItems() *QXmlResultItems {
	return NewQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
}

func (ptr *QXmlResultItems) Current() *QXmlItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlItemFromPointer(C.QXmlResultItems_Current(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlResultItems) HasError() bool {
	if ptr.Pointer() != nil {
		return C.QXmlResultItems_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlResultItems) Next() *QXmlItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlItemFromPointer(C.QXmlResultItems_Next(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

//export callbackQXmlResultItems_DestroyQXmlResultItems
func callbackQXmlResultItems_DestroyQXmlResultItems(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlResultItems::~QXmlResultItems"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlResultItemsFromPointer(ptr).DestroyQXmlResultItemsDefault()
	}
}

func (ptr *QXmlResultItems) ConnectDestroyQXmlResultItems(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlResultItems::~QXmlResultItems", f)
	}
}

func (ptr *QXmlResultItems) DisconnectDestroyQXmlResultItems() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlResultItems::~QXmlResultItems")
	}
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlResultItems) DestroyQXmlResultItemsDefault() {
	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItemsDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QXmlSchema struct {
	ptr unsafe.Pointer
}

type QXmlSchema_ITF interface {
	QXmlSchema_PTR() *QXmlSchema
}

func (ptr *QXmlSchema) QXmlSchema_PTR() *QXmlSchema {
	return ptr
}

func (ptr *QXmlSchema) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlSchema) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQXmlSchema() *QXmlSchema {
	var tmpValue = NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema())
	runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
	return tmpValue
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {
	var tmpValue = NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema2(PointerFromQXmlSchema(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
	return tmpValue
}

func (ptr *QXmlSchema) DocumentUri() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QXmlSchema_DocumentUri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractMessageHandlerFromPointer(C.QXmlSchema_MessageHandler(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlSchema_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QXmlSchema_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractUriResolverFromPointer(C.QXmlSchema_UriResolver(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {
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

func (ptr *QXmlSchemaValidator) QXmlSchemaValidator_PTR() *QXmlSchemaValidator {
	return ptr
}

func (ptr *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQXmlSchemaValidator() *QXmlSchemaValidator {
	var tmpValue = NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator())
	runtime.SetFinalizer(tmpValue, (*QXmlSchemaValidator).DestroyQXmlSchemaValidator)
	return tmpValue
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {
	var tmpValue = NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(PointerFromQXmlSchema(schema)))
	runtime.SetFinalizer(tmpValue, (*QXmlSchemaValidator).DestroyQXmlSchemaValidator)
	return tmpValue
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractMessageHandlerFromPointer(C.QXmlSchemaValidator_MessageHandler(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlNamePoolFromPointer(C.QXmlSchemaValidator_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QXmlSchemaValidator_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Schema() *QXmlSchema {
	if ptr.Pointer() != nil {
		var tmpValue = NewQXmlSchemaFromPointer(C.QXmlSchemaValidator_Schema(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchema_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetSchema(ptr.Pointer(), PointerFromQXmlSchema(schema))
	}
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractUriResolverFromPointer(C.QXmlSchemaValidator_UriResolver(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
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

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func (ptr *QXmlSerializer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlSerializer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractXmlReceiver_PTR().SetPointer(p)
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

func (ptr *QXmlSerializer) DestroyQXmlSerializer() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	return NewQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

//export callbackQXmlSerializer_AtomicValue
func callbackQXmlSerializer_AtomicValue(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AtomicValueDefault(core.NewQVariantFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAtomicValue(f func(value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::atomicValue", f)
	}
}

func (ptr *QXmlSerializer) DisconnectAtomicValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::atomicValue")
	}
}

func (ptr *QXmlSerializer) AtomicValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QXmlSerializer) AtomicValueDefault(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_AtomicValueDefault(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQXmlSerializer_Attribute
func callbackQXmlSerializer_Attribute(ptr unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::attribute"); signal != nil {
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AttributeDefault(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::attribute", f)
	}
}

func (ptr *QXmlSerializer) DisconnectAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::attribute")
	}
}

func (ptr *QXmlSerializer) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) AttributeDefault(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_AttributeDefault(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlSerializer_Characters
func callbackQXmlSerializer_Characters(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::characters"); signal != nil {
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::characters", f)
	}
}

func (ptr *QXmlSerializer) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::characters")
	}
}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) CharactersDefault(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlSerializer_Comment
func callbackQXmlSerializer_Comment(ptr unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::comment"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CommentDefault(cGoUnpackString(value))
	}
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::comment", f)
	}
}

func (ptr *QXmlSerializer) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::comment")
	}
}

func (ptr *QXmlSerializer) Comment(value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_Comment(ptr.Pointer(), valueC)
	}
}

func (ptr *QXmlSerializer) CommentDefault(value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_CommentDefault(ptr.Pointer(), valueC)
	}
}

//export callbackQXmlSerializer_EndDocument
func callbackQXmlSerializer_EndDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::endDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::endDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::endDocument")
	}
}

func (ptr *QXmlSerializer) EndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndDocumentDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_EndElement
func callbackQXmlSerializer_EndElement(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::endElement"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::endElement", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::endElement")
	}
}

func (ptr *QXmlSerializer) EndElement() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndElementDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElementDefault(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSerializer_EndOfSequence
func callbackQXmlSerializer_EndOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::endOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::endOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectEndOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::endOfSequence")
	}
}

func (ptr *QXmlSerializer) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndOfSequenceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequenceDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_NamespaceBinding
func callbackQXmlSerializer_NamespaceBinding(ptr unsafe.Pointer, nb unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::namespaceBinding"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(nb))
	} else {
		NewQXmlSerializerFromPointer(ptr).NamespaceBindingDefault(NewQXmlNameFromPointer(nb))
	}
}

func (ptr *QXmlSerializer) ConnectNamespaceBinding(f func(nb *QXmlName)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::namespaceBinding", f)
	}
}

func (ptr *QXmlSerializer) DisconnectNamespaceBinding() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::namespaceBinding")
	}
}

func (ptr *QXmlSerializer) NamespaceBinding(nb QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) NamespaceBindingDefault(nb QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBindingDefault(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QXmlSerializer_OutputDevice(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQXmlSerializer_ProcessingInstruction
func callbackQXmlSerializer_ProcessingInstruction(ptr unsafe.Pointer, name unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).ProcessingInstructionDefault(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	}
}

func (ptr *QXmlSerializer) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::processingInstruction", f)
	}
}

func (ptr *QXmlSerializer) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::processingInstruction")
	}
}

func (ptr *QXmlSerializer) ProcessingInstruction(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlSerializer) ProcessingInstructionDefault(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QXmlSerializer_ProcessingInstructionDefault(ptr.Pointer(), PointerFromQXmlName(name), valueC)
	}
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(ptr.Pointer(), core.PointerFromQTextCodec(outputCodec))
	}
}

//export callbackQXmlSerializer_StartDocument
func callbackQXmlSerializer_StartDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::startDocument"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::startDocument", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::startDocument")
	}
}

func (ptr *QXmlSerializer) StartDocument() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartDocumentDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_StartElement
func callbackQXmlSerializer_StartElement(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::startElement"); signal != nil {
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	} else {
		NewQXmlSerializerFromPointer(ptr).StartElementDefault(NewQXmlNameFromPointer(name))
	}
}

func (ptr *QXmlSerializer) ConnectStartElement(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::startElement", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::startElement")
	}
}

func (ptr *QXmlSerializer) StartElement(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlSerializer) StartElementDefault(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElementDefault(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQXmlSerializer_StartOfSequence
func callbackQXmlSerializer_StartOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXmlSerializer::startOfSequence"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::startOfSequence", f)
	}
}

func (ptr *QXmlSerializer) DisconnectStartOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXmlSerializer::startOfSequence")
	}
}

func (ptr *QXmlSerializer) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartOfSequenceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequenceDefault(ptr.Pointer())
	}
}
