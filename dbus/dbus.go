// +build !minimal

package dbus

//#include <stdint.h>
//#include <stdlib.h>
//#include "dbus.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

type QDBusAbstractAdaptor struct {
	core.QObject
}

type QDBusAbstractAdaptor_ITF interface {
	core.QObject_ITF
	QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor
}

func (p *QDBusAbstractAdaptor) QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor {
	return p
}

func (p *QDBusAbstractAdaptor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDBusAbstractAdaptor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusAbstractAdaptor(ptr QDBusAbstractAdaptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractAdaptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) *QDBusAbstractAdaptor {
	var n = new(QDBusAbstractAdaptor)
	n.SetPointer(ptr)
	return n
}
func NewQDBusAbstractAdaptor(obj core.QObject_ITF) *QDBusAbstractAdaptor {
	defer qt.Recovering("QDBusAbstractAdaptor::QDBusAbstractAdaptor")

	return NewQDBusAbstractAdaptorFromPointer(C.QDBusAbstractAdaptor_NewQDBusAbstractAdaptor(core.PointerFromQObject(obj)))
}

func (ptr *QDBusAbstractAdaptor) AutoRelaySignals() bool {
	defer qt.Recovering("QDBusAbstractAdaptor::autoRelaySignals")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_AutoRelaySignals(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractAdaptor) SetAutoRelaySignals(enable bool) {
	defer qt.Recovering("QDBusAbstractAdaptor::setAutoRelaySignals")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_SetAutoRelaySignals(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {
	defer qt.Recovering("QDBusAbstractAdaptor::~QDBusAbstractAdaptor")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractAdaptor_TimerEvent
func callbackQDBusAbstractAdaptor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_ChildEvent
func callbackQDBusAbstractAdaptor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_ConnectNotify
func callbackQDBusAbstractAdaptor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractAdaptor_CustomEvent
func callbackQDBusAbstractAdaptor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_DeleteLater
func callbackQDBusAbstractAdaptor_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusAbstractAdaptor) DeleteLater() {
	defer qt.Recovering("QDBusAbstractAdaptor::deleteLater")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractAdaptor) DeleteLaterDefault() {
	defer qt.Recovering("QDBusAbstractAdaptor::deleteLater")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractAdaptor_DisconnectNotify
func callbackQDBusAbstractAdaptor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractAdaptor_Event
func callbackQDBusAbstractAdaptor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusAbstractAdaptor::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractAdaptorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusAbstractAdaptor) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusAbstractAdaptor) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractAdaptor::event")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusAbstractAdaptor) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractAdaptor::event")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_EventFilter
func callbackQDBusAbstractAdaptor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusAbstractAdaptor::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractAdaptorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusAbstractAdaptor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusAbstractAdaptor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractAdaptor::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusAbstractAdaptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractAdaptor::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_MetaObject
func callbackQDBusAbstractAdaptor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusAbstractAdaptor::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusAbstractAdaptorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusAbstractAdaptor) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractAdaptor(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusAbstractAdaptor) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusAbstractAdaptor::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractAdaptor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusAbstractAdaptor::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractAdaptor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDBusAbstractInterface struct {
	core.QObject
}

type QDBusAbstractInterface_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterface_PTR() *QDBusAbstractInterface
}

func (p *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return p
}

func (p *QDBusAbstractInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDBusAbstractInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) *QDBusAbstractInterface {
	var n = new(QDBusAbstractInterface)
	n.SetPointer(ptr)
	return n
}
func (ptr *QDBusAbstractInterface) AsyncCall(method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusAbstractInterface::asyncCall")

	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusAbstractInterface_AsyncCall(ptr.Pointer(), methodC, core.PointerFromQVariant(arg1), core.PointerFromQVariant(arg2), core.PointerFromQVariant(arg3), core.PointerFromQVariant(arg4), core.PointerFromQVariant(arg5), core.PointerFromQVariant(arg6), core.PointerFromQVariant(arg7), core.PointerFromQVariant(arg8)))
		runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Call(method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusAbstractInterface::call")

	if ptr.Pointer() != nil {
		var methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusAbstractInterface_Call(ptr.Pointer(), methodC, core.PointerFromQVariant(arg1), core.PointerFromQVariant(arg2), core.PointerFromQVariant(arg3), core.PointerFromQVariant(arg4), core.PointerFromQVariant(arg5), core.PointerFromQVariant(arg6), core.PointerFromQVariant(arg7), core.PointerFromQVariant(arg8)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Connection() *QDBusConnection {
	defer qt.Recovering("QDBusAbstractInterface::connection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionFromPointer(C.QDBusAbstractInterface_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Interface() string {
	defer qt.Recovering("QDBusAbstractInterface::interface")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	defer qt.Recovering("QDBusAbstractInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) LastError() *QDBusError {
	defer qt.Recovering("QDBusAbstractInterface::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusErrorFromPointer(C.QDBusAbstractInterface_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Path() string {
	defer qt.Recovering("QDBusAbstractInterface::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	defer qt.Recovering("QDBusAbstractInterface::service")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	defer qt.Recovering("QDBusAbstractInterface::setTimeout")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(int32(timeout)))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	defer qt.Recovering("QDBusAbstractInterface::timeout")

	if ptr.Pointer() != nil {
		return int(int32(C.QDBusAbstractInterface_Timeout(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	defer qt.Recovering("QDBusAbstractInterface::~QDBusAbstractInterface")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()))
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractInterface_TimerEvent
func callbackQDBusAbstractInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusAbstractInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusAbstractInterface_ChildEvent
func callbackQDBusAbstractInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusAbstractInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusAbstractInterface_ConnectNotify
func callbackQDBusAbstractInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusAbstractInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusAbstractInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractInterface_CustomEvent
func callbackQDBusAbstractInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusAbstractInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusAbstractInterface_DeleteLater
func callbackQDBusAbstractInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusAbstractInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusAbstractInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusAbstractInterface) DeleteLater() {
	defer qt.Recovering("QDBusAbstractInterface::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()))
		C.QDBusAbstractInterface_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDBusAbstractInterface::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()))
		C.QDBusAbstractInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractInterface_DisconnectNotify
func callbackQDBusAbstractInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusAbstractInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusAbstractInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractInterface_Event
func callbackQDBusAbstractInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusAbstractInterface::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusAbstractInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusAbstractInterface::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusAbstractInterface) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractInterface::event")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractInterface::event")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_EventFilter
func callbackQDBusAbstractInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusAbstractInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusAbstractInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusAbstractInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusAbstractInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusAbstractInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_MetaObject
func callbackQDBusAbstractInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusAbstractInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusAbstractInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusAbstractInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusAbstractInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusAbstractInterface(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusAbstractInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusAbstractInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusAbstractInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusAbstractInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QDBusArgument::ElementType
type QDBusArgument__ElementType int64

const (
	QDBusArgument__BasicType     = QDBusArgument__ElementType(0)
	QDBusArgument__VariantType   = QDBusArgument__ElementType(1)
	QDBusArgument__ArrayType     = QDBusArgument__ElementType(2)
	QDBusArgument__StructureType = QDBusArgument__ElementType(3)
	QDBusArgument__MapType       = QDBusArgument__ElementType(4)
	QDBusArgument__MapEntryType  = QDBusArgument__ElementType(5)
	QDBusArgument__UnknownType   = QDBusArgument__ElementType(-1)
)

type QDBusArgument struct {
	ptr unsafe.Pointer
}

type QDBusArgument_ITF interface {
	QDBusArgument_PTR() *QDBusArgument
}

func (p *QDBusArgument) QDBusArgument_PTR() *QDBusArgument {
	return p
}

func (p *QDBusArgument) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusArgument) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusArgument(ptr QDBusArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusArgument_PTR().Pointer()
	}
	return nil
}

func NewQDBusArgumentFromPointer(ptr unsafe.Pointer) *QDBusArgument {
	var n = new(QDBusArgument)
	n.SetPointer(ptr)
	return n
}
func NewQDBusArgument() *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	var tmpValue = NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument())
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func NewQDBusArgument3(other QDBusArgument_ITF) *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	var tmpValue = NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument3(PointerFromQDBusArgument(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func NewQDBusArgument2(other QDBusArgument_ITF) *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	var tmpValue = NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument2(PointerFromQDBusArgument(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func (ptr *QDBusArgument) AsVariant() *core.QVariant {
	defer qt.Recovering("QDBusArgument::asVariant")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDBusArgument_AsVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusArgument) AtEnd() bool {
	defer qt.Recovering("QDBusArgument::atEnd")

	if ptr.Pointer() != nil {
		return C.QDBusArgument_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusArgument) BeginArray(id int) {
	defer qt.Recovering("QDBusArgument::beginArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QDBusArgument) BeginArray2() {
	defer qt.Recovering("QDBusArgument::beginArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {
	defer qt.Recovering("QDBusArgument::beginMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap(ptr.Pointer(), C.int(int32(kid)), C.int(int32(vid)))
	}
}

func (ptr *QDBusArgument) BeginMap2() {
	defer qt.Recovering("QDBusArgument::beginMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry() {
	defer qt.Recovering("QDBusArgument::beginMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry2() {
	defer qt.Recovering("QDBusArgument::beginMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure() {
	defer qt.Recovering("QDBusArgument::beginStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure2() {
	defer qt.Recovering("QDBusArgument::beginStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {
	defer qt.Recovering("QDBusArgument::currentType")

	if ptr.Pointer() != nil {
		return QDBusArgument__ElementType(C.QDBusArgument_CurrentType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusArgument) EndArray() {
	defer qt.Recovering("QDBusArgument::endArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndArray2() {
	defer qt.Recovering("QDBusArgument::endArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap() {
	defer qt.Recovering("QDBusArgument::endMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap2() {
	defer qt.Recovering("QDBusArgument::endMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry() {
	defer qt.Recovering("QDBusArgument::endMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry2() {
	defer qt.Recovering("QDBusArgument::endMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure() {
	defer qt.Recovering("QDBusArgument::endStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure2() {
	defer qt.Recovering("QDBusArgument::endStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) Swap(other QDBusArgument_ITF) {
	defer qt.Recovering("QDBusArgument::swap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_Swap(ptr.Pointer(), PointerFromQDBusArgument(other))
	}
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {
	defer qt.Recovering("QDBusArgument::~QDBusArgument")

	if ptr.Pointer() != nil {
		C.QDBusArgument_DestroyQDBusArgument(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QDBusConnection::BusType
type QDBusConnection__BusType int64

const (
	QDBusConnection__SessionBus    = QDBusConnection__BusType(0)
	QDBusConnection__SystemBus     = QDBusConnection__BusType(1)
	QDBusConnection__ActivationBus = QDBusConnection__BusType(2)
)

//QDBusConnection::ConnectionCapability
type QDBusConnection__ConnectionCapability int64

const (
	QDBusConnection__UnixFileDescriptorPassing = QDBusConnection__ConnectionCapability(0x0001)
)

//QDBusConnection::RegisterOption
type QDBusConnection__RegisterOption int64

const (
	QDBusConnection__ExportAdaptors                = QDBusConnection__RegisterOption(0x01)
	QDBusConnection__ExportScriptableSlots         = QDBusConnection__RegisterOption(0x10)
	QDBusConnection__ExportScriptableSignals       = QDBusConnection__RegisterOption(0x20)
	QDBusConnection__ExportScriptableProperties    = QDBusConnection__RegisterOption(0x40)
	QDBusConnection__ExportScriptableInvokables    = QDBusConnection__RegisterOption(0x80)
	QDBusConnection__ExportScriptableContents      = QDBusConnection__RegisterOption(0xf0)
	QDBusConnection__ExportNonScriptableSlots      = QDBusConnection__RegisterOption(0x100)
	QDBusConnection__ExportNonScriptableSignals    = QDBusConnection__RegisterOption(0x200)
	QDBusConnection__ExportNonScriptableProperties = QDBusConnection__RegisterOption(0x400)
	QDBusConnection__ExportNonScriptableInvokables = QDBusConnection__RegisterOption(0x800)
	QDBusConnection__ExportNonScriptableContents   = QDBusConnection__RegisterOption(0xf00)
	QDBusConnection__ExportAllSlots                = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSlots | QDBusConnection__ExportNonScriptableSlots)
	QDBusConnection__ExportAllSignals              = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSignals | QDBusConnection__ExportNonScriptableSignals)
	QDBusConnection__ExportAllProperties           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableProperties | QDBusConnection__ExportNonScriptableProperties)
	QDBusConnection__ExportAllInvokables           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableInvokables | QDBusConnection__ExportNonScriptableInvokables)
	QDBusConnection__ExportAllContents             = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableContents | QDBusConnection__ExportNonScriptableContents)
	QDBusConnection__ExportChildObjects            = QDBusConnection__RegisterOption(0x1000)
)

//QDBusConnection::UnregisterMode
type QDBusConnection__UnregisterMode int64

const (
	QDBusConnection__UnregisterNode = QDBusConnection__UnregisterMode(0)
	QDBusConnection__UnregisterTree = QDBusConnection__UnregisterMode(1)
)

type QDBusConnection struct {
	ptr unsafe.Pointer
}

type QDBusConnection_ITF interface {
	QDBusConnection_PTR() *QDBusConnection
}

func (p *QDBusConnection) QDBusConnection_PTR() *QDBusConnection {
	return p
}

func (p *QDBusConnection) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusConnection) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusConnection(ptr QDBusConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnection_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionFromPointer(ptr unsafe.Pointer) *QDBusConnection {
	var n = new(QDBusConnection)
	n.SetPointer(ptr)
	return n
}
func QDBusConnection_SessionBus() *QDBusConnection {
	defer qt.Recovering("QDBusConnection::sessionBus")

	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SessionBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) SessionBus() *QDBusConnection {
	defer qt.Recovering("QDBusConnection::sessionBus")

	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SessionBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_SystemBus() *QDBusConnection {
	defer qt.Recovering("QDBusConnection::systemBus")

	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SystemBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) SystemBus() *QDBusConnection {
	defer qt.Recovering("QDBusConnection::systemBus")

	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SystemBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection3(other QDBusConnection_ITF) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection3(PointerFromQDBusConnection(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection2(other QDBusConnection_ITF) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection2(PointerFromQDBusConnection(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection(name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection(nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) BaseService() string {
	defer qt.Recovering("QDBusConnection::baseService")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_BaseService(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) AsyncCall(message QDBusMessage_ITF, timeout int) *QDBusPendingCall {
	defer qt.Recovering("QDBusConnection::asyncCall")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusConnection_AsyncCall(ptr.Pointer(), PointerFromQDBusMessage(message), C.int(int32(timeout))))
		runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessage_ITF, receiver core.QObject_ITF, returnMethod string, errorMethod string, timeout int) bool {
	defer qt.Recovering("QDBusConnection::callWithCallback")

	if ptr.Pointer() != nil {
		var returnMethodC = C.CString(returnMethod)
		defer C.free(unsafe.Pointer(returnMethodC))
		var errorMethodC = C.CString(errorMethod)
		defer C.free(unsafe.Pointer(errorMethodC))
		return C.QDBusConnection_CallWithCallback(ptr.Pointer(), PointerFromQDBusMessage(message), core.PointerFromQObject(receiver), returnMethodC, errorMethodC, C.int(int32(timeout))) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QDBusConnection_Connect(ptr.Pointer(), serviceC, pathC, interfaC, nameC, core.PointerFromQObject(receiver), slotC) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QDBusConnection_Connect2(ptr.Pointer(), serviceC, pathC, interfaC, nameC, signatureC, core.PointerFromQObject(receiver), slotC) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var argumentMatchC = C.CString(strings.Join(argumentMatch, "|"))
		defer C.free(unsafe.Pointer(argumentMatchC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QDBusConnection_Connect3(ptr.Pointer(), serviceC, pathC, interfaC, nameC, argumentMatchC, signatureC, core.PointerFromQObject(receiver), slotC) != 0
	}
	return false
}

func QDBusConnection_ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::connectToBus")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus(C.longlong(ty), nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::connectToBus")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus(C.longlong(ty), nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_ConnectToBus2(address string, name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::connectToBus")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus2(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToBus2(address string, name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::connectToBus")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus2(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_ConnectToPeer(address string, name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::connectToPeer")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToPeer(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToPeer(address string, name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::connectToPeer")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToPeer(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {
	defer qt.Recovering("QDBusConnection::connectionCapabilities")

	if ptr.Pointer() != nil {
		return QDBusConnection__ConnectionCapability(C.QDBusConnection_ConnectionCapabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusConnection) Disconnect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::disconnect")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QDBusConnection_Disconnect(ptr.Pointer(), serviceC, pathC, interfaC, nameC, core.PointerFromQObject(receiver), slotC) != 0
	}
	return false
}

func (ptr *QDBusConnection) Disconnect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::disconnect")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QDBusConnection_Disconnect2(ptr.Pointer(), serviceC, pathC, interfaC, nameC, signatureC, core.PointerFromQObject(receiver), slotC) != 0
	}
	return false
}

func (ptr *QDBusConnection) Disconnect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::disconnect")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var argumentMatchC = C.CString(strings.Join(argumentMatch, "|"))
		defer C.free(unsafe.Pointer(argumentMatchC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QDBusConnection_Disconnect3(ptr.Pointer(), serviceC, pathC, interfaC, nameC, argumentMatchC, signatureC, core.PointerFromQObject(receiver), slotC) != 0
	}
	return false
}

func QDBusConnection_DisconnectFromBus(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromBus")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(nameC)
}

func (ptr *QDBusConnection) DisconnectFromBus(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromBus")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(nameC)
}

func QDBusConnection_DisconnectFromPeer(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromPeer")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(nameC)
}

func (ptr *QDBusConnection) DisconnectFromPeer(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromPeer")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(nameC)
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {
	defer qt.Recovering("QDBusConnection::interface")

	if ptr.Pointer() != nil {
		return NewQDBusConnectionInterfaceFromPointer(C.QDBusConnection_Interface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusConnection) IsConnected() bool {
	defer qt.Recovering("QDBusConnection::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusConnection) LastError() *QDBusError {
	defer qt.Recovering("QDBusConnection::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusErrorFromPointer(C.QDBusConnection_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func QDBusConnection_LocalMachineId() string {
	defer qt.Recovering("QDBusConnection::localMachineId")

	return qt.HexDecodeToString(C.GoString(C.QDBusConnection_QDBusConnection_LocalMachineId()))
}

func (ptr *QDBusConnection) LocalMachineId() string {
	defer qt.Recovering("QDBusConnection::localMachineId")

	return qt.HexDecodeToString(C.GoString(C.QDBusConnection_QDBusConnection_LocalMachineId()))
}

func (ptr *QDBusConnection) Name() string {
	defer qt.Recovering("QDBusConnection::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {
	defer qt.Recovering("QDBusConnection::objectRegisteredAt")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		return core.NewQObjectFromPointer(C.QDBusConnection_ObjectRegisteredAt(ptr.Pointer(), pathC))
	}
	return nil
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	defer qt.Recovering("QDBusConnection::registerObject")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		return C.QDBusConnection_RegisterObject(ptr.Pointer(), pathC, core.PointerFromQObject(object), C.longlong(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	defer qt.Recovering("QDBusConnection::registerObject")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
		return C.QDBusConnection_RegisterObject2(ptr.Pointer(), pathC, interfaC, core.PointerFromQObject(object), C.longlong(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterService(serviceName string) bool {
	defer qt.Recovering("QDBusConnection::registerService")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		return C.QDBusConnection_RegisterService(ptr.Pointer(), serviceNameC) != 0
	}
	return false
}

func (ptr *QDBusConnection) Send(message QDBusMessage_ITF) bool {
	defer qt.Recovering("QDBusConnection::send")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Send(ptr.Pointer(), PointerFromQDBusMessage(message)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Swap(other QDBusConnection_ITF) {
	defer qt.Recovering("QDBusConnection::swap")

	if ptr.Pointer() != nil {
		C.QDBusConnection_Swap(ptr.Pointer(), PointerFromQDBusConnection(other))
	}
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {
	defer qt.Recovering("QDBusConnection::unregisterObject")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QDBusConnection_UnregisterObject(ptr.Pointer(), pathC, C.longlong(mode))
	}
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {
	defer qt.Recovering("QDBusConnection::unregisterService")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		return C.QDBusConnection_UnregisterService(ptr.Pointer(), serviceNameC) != 0
	}
	return false
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {
	defer qt.Recovering("QDBusConnection::~QDBusConnection")

	if ptr.Pointer() != nil {
		C.QDBusConnection_DestroyQDBusConnection(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QDBusConnectionInterface::RegisterServiceReply
type QDBusConnectionInterface__RegisterServiceReply int64

const (
	QDBusConnectionInterface__ServiceNotRegistered = QDBusConnectionInterface__RegisterServiceReply(0)
	QDBusConnectionInterface__ServiceRegistered    = QDBusConnectionInterface__RegisterServiceReply(1)
	QDBusConnectionInterface__ServiceQueued        = QDBusConnectionInterface__RegisterServiceReply(2)
)

//QDBusConnectionInterface::ServiceQueueOptions
type QDBusConnectionInterface__ServiceQueueOptions int64

const (
	QDBusConnectionInterface__DontQueueService       = QDBusConnectionInterface__ServiceQueueOptions(0)
	QDBusConnectionInterface__QueueService           = QDBusConnectionInterface__ServiceQueueOptions(1)
	QDBusConnectionInterface__ReplaceExistingService = QDBusConnectionInterface__ServiceQueueOptions(2)
)

//QDBusConnectionInterface::ServiceReplacementOptions
type QDBusConnectionInterface__ServiceReplacementOptions int64

const (
	QDBusConnectionInterface__DontAllowReplacement = QDBusConnectionInterface__ServiceReplacementOptions(0)
	QDBusConnectionInterface__AllowReplacement     = QDBusConnectionInterface__ServiceReplacementOptions(1)
)

type QDBusConnectionInterface struct {
	QDBusAbstractInterface
}

type QDBusConnectionInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusConnectionInterface_PTR() *QDBusConnectionInterface
}

func (p *QDBusConnectionInterface) QDBusConnectionInterface_PTR() *QDBusConnectionInterface {
	return p
}

func (p *QDBusConnectionInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (p *QDBusConnectionInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDBusAbstractInterface_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusConnectionInterface(ptr QDBusConnectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) *QDBusConnectionInterface {
	var n = new(QDBusConnectionInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusConnectionInterface) DestroyQDBusConnectionInterface() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

//export callbackQDBusConnectionInterface_CallWithCallbackFailed
func callbackQDBusConnectionInterface_CallWithCallbackFailed(ptr unsafe.Pointer, error unsafe.Pointer, call unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::callWithCallbackFailed")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "callWithCallbackFailed"); signal != nil {
		signal.(func(*QDBusError, *QDBusMessage))(NewQDBusErrorFromPointer(error), NewQDBusMessageFromPointer(call))
	}

}

func (ptr *QDBusConnectionInterface) ConnectCallWithCallbackFailed(f func(error *QDBusError, call *QDBusMessage)) {
	defer qt.Recovering("connect QDBusConnectionInterface::callWithCallbackFailed")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectCallWithCallbackFailed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "callWithCallbackFailed", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCallWithCallbackFailed() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::callWithCallbackFailed")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectCallWithCallbackFailed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "callWithCallbackFailed")
	}
}

func (ptr *QDBusConnectionInterface) CallWithCallbackFailed(error QDBusError_ITF, call QDBusMessage_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::callWithCallbackFailed")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CallWithCallbackFailed(ptr.Pointer(), PointerFromQDBusError(error), PointerFromQDBusMessage(call))
	}
}

//export callbackQDBusConnectionInterface_ServiceRegistered
func callbackQDBusConnectionInterface_ServiceRegistered(ptr unsafe.Pointer, serviceName *C.char) {
	defer qt.Recovering("callback QDBusConnectionInterface::serviceRegistered")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "serviceRegistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "serviceRegistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "serviceRegistered")
	}
}

func (ptr *QDBusConnectionInterface) ServiceRegistered(serviceName string) {
	defer qt.Recovering("QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusConnectionInterface_ServiceRegistered(ptr.Pointer(), serviceNameC)
	}
}

//export callbackQDBusConnectionInterface_ServiceUnregistered
func callbackQDBusConnectionInterface_ServiceUnregistered(ptr unsafe.Pointer, serviceName *C.char) {
	defer qt.Recovering("callback QDBusConnectionInterface::serviceUnregistered")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "serviceUnregistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "serviceUnregistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "serviceUnregistered")
	}
}

func (ptr *QDBusConnectionInterface) ServiceUnregistered(serviceName string) {
	defer qt.Recovering("QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusConnectionInterface_ServiceUnregistered(ptr.Pointer(), serviceNameC)
	}
}

//export callbackQDBusConnectionInterface_TimerEvent
func callbackQDBusConnectionInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusConnectionInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusConnectionInterface_ChildEvent
func callbackQDBusConnectionInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusConnectionInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusConnectionInterface_ConnectNotify
func callbackQDBusConnectionInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusConnectionInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusConnectionInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusConnectionInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusConnectionInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusConnectionInterface_CustomEvent
func callbackQDBusConnectionInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusConnectionInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusConnectionInterface_DeleteLater
func callbackQDBusConnectionInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusConnectionInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusConnectionInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusConnectionInterface) DeleteLater() {
	defer qt.Recovering("QDBusConnectionInterface::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()))
		C.QDBusConnectionInterface_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusConnectionInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDBusConnectionInterface::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()))
		C.QDBusConnectionInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusConnectionInterface_DisconnectNotify
func callbackQDBusConnectionInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusConnectionInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusConnectionInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusConnectionInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusConnectionInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusConnectionInterface_Event
func callbackQDBusConnectionInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusConnectionInterface::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusConnectionInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusConnectionInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusConnectionInterface::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusConnectionInterface) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusConnectionInterface::event")

	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusConnectionInterface) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusConnectionInterface::event")

	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusConnectionInterface_EventFilter
func callbackQDBusConnectionInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusConnectionInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusConnectionInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusConnectionInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusConnectionInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusConnectionInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusConnectionInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusConnectionInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusConnectionInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusConnectionInterface_MetaObject
func callbackQDBusConnectionInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusConnectionInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusConnectionInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusConnectionInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusConnectionInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusConnectionInterface(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusConnectionInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusConnectionInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusConnectionInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusConnectionInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusConnectionInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusConnectionInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDBusContext struct {
	ptr unsafe.Pointer
}

type QDBusContext_ITF interface {
	QDBusContext_PTR() *QDBusContext
}

func (p *QDBusContext) QDBusContext_PTR() *QDBusContext {
	return p
}

func (p *QDBusContext) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusContext) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusContext(ptr QDBusContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusContext_PTR().Pointer()
	}
	return nil
}

func NewQDBusContextFromPointer(ptr unsafe.Pointer) *QDBusContext {
	var n = new(QDBusContext)
	n.SetPointer(ptr)
	return n
}
func NewQDBusContext() *QDBusContext {
	defer qt.Recovering("QDBusContext::QDBusContext")

	var tmpValue = NewQDBusContextFromPointer(C.QDBusContext_NewQDBusContext())
	runtime.SetFinalizer(tmpValue, (*QDBusContext).DestroyQDBusContext)
	return tmpValue
}

func (ptr *QDBusContext) CalledFromDBus() bool {
	defer qt.Recovering("QDBusContext::calledFromDBus")

	if ptr.Pointer() != nil {
		return C.QDBusContext_CalledFromDBus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) Connection() *QDBusConnection {
	defer qt.Recovering("QDBusContext::connection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionFromPointer(C.QDBusContext_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusContext) IsDelayedReply() bool {
	defer qt.Recovering("QDBusContext::isDelayedReply")

	if ptr.Pointer() != nil {
		return C.QDBusContext_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) Message() *QDBusMessage {
	defer qt.Recovering("QDBusContext::message")

	if ptr.Pointer() != nil {
		return NewQDBusMessageFromPointer(C.QDBusContext_Message(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {
	defer qt.Recovering("QDBusContext::sendErrorReply")

	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QDBusContext_SendErrorReply2(ptr.Pointer(), C.longlong(ty), msgC)
	}
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {
	defer qt.Recovering("QDBusContext::sendErrorReply")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QDBusContext_SendErrorReply(ptr.Pointer(), nameC, msgC)
	}
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {
	defer qt.Recovering("QDBusContext::setDelayedReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SetDelayedReply(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusContext) DestroyQDBusContext() {
	defer qt.Recovering("QDBusContext::~QDBusContext")

	if ptr.Pointer() != nil {
		C.QDBusContext_DestroyQDBusContext(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QDBusError::ErrorType
type QDBusError__ErrorType int64

const (
	QDBusError__NoError           = QDBusError__ErrorType(0)
	QDBusError__Other             = QDBusError__ErrorType(1)
	QDBusError__Failed            = QDBusError__ErrorType(2)
	QDBusError__NoMemory          = QDBusError__ErrorType(3)
	QDBusError__ServiceUnknown    = QDBusError__ErrorType(4)
	QDBusError__NoReply           = QDBusError__ErrorType(5)
	QDBusError__BadAddress        = QDBusError__ErrorType(6)
	QDBusError__NotSupported      = QDBusError__ErrorType(7)
	QDBusError__LimitsExceeded    = QDBusError__ErrorType(8)
	QDBusError__AccessDenied      = QDBusError__ErrorType(9)
	QDBusError__NoServer          = QDBusError__ErrorType(10)
	QDBusError__Timeout           = QDBusError__ErrorType(11)
	QDBusError__NoNetwork         = QDBusError__ErrorType(12)
	QDBusError__AddressInUse      = QDBusError__ErrorType(13)
	QDBusError__Disconnected      = QDBusError__ErrorType(14)
	QDBusError__InvalidArgs       = QDBusError__ErrorType(15)
	QDBusError__UnknownMethod     = QDBusError__ErrorType(16)
	QDBusError__TimedOut          = QDBusError__ErrorType(17)
	QDBusError__InvalidSignature  = QDBusError__ErrorType(18)
	QDBusError__UnknownInterface  = QDBusError__ErrorType(19)
	QDBusError__UnknownObject     = QDBusError__ErrorType(20)
	QDBusError__UnknownProperty   = QDBusError__ErrorType(21)
	QDBusError__PropertyReadOnly  = QDBusError__ErrorType(22)
	QDBusError__InternalError     = QDBusError__ErrorType(23)
	QDBusError__InvalidService    = QDBusError__ErrorType(24)
	QDBusError__InvalidObjectPath = QDBusError__ErrorType(25)
	QDBusError__InvalidInterface  = QDBusError__ErrorType(26)
	QDBusError__InvalidMember     = QDBusError__ErrorType(27)
)

type QDBusError struct {
	ptr unsafe.Pointer
}

type QDBusError_ITF interface {
	QDBusError_PTR() *QDBusError
}

func (p *QDBusError) QDBusError_PTR() *QDBusError {
	return p
}

func (p *QDBusError) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusError) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusError(ptr QDBusError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusError_PTR().Pointer()
	}
	return nil
}

func NewQDBusErrorFromPointer(ptr unsafe.Pointer) *QDBusError {
	var n = new(QDBusError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusError) DestroyQDBusError() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQDBusError(other QDBusError_ITF) *QDBusError {
	defer qt.Recovering("QDBusError::QDBusError")

	var tmpValue = NewQDBusErrorFromPointer(C.QDBusError_NewQDBusError(PointerFromQDBusError(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
	return tmpValue
}

func QDBusError_ErrorString(error QDBusError__ErrorType) string {
	defer qt.Recovering("QDBusError::errorString")

	return C.GoString(C.QDBusError_QDBusError_ErrorString(C.longlong(error)))
}

func (ptr *QDBusError) ErrorString(error QDBusError__ErrorType) string {
	defer qt.Recovering("QDBusError::errorString")

	return C.GoString(C.QDBusError_QDBusError_ErrorString(C.longlong(error)))
}

func (ptr *QDBusError) IsValid() bool {
	defer qt.Recovering("QDBusError::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusError) Message() string {
	defer qt.Recovering("QDBusError::message")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusError_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Name() string {
	defer qt.Recovering("QDBusError::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusError_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Swap(other QDBusError_ITF) {
	defer qt.Recovering("QDBusError::swap")

	if ptr.Pointer() != nil {
		C.QDBusError_Swap(ptr.Pointer(), PointerFromQDBusError(other))
	}
}

func (ptr *QDBusError) Type() QDBusError__ErrorType {
	defer qt.Recovering("QDBusError::type")

	if ptr.Pointer() != nil {
		return QDBusError__ErrorType(C.QDBusError_Type(ptr.Pointer()))
	}
	return 0
}

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusInterface_PTR() *QDBusInterface
}

func (p *QDBusInterface) QDBusInterface_PTR() *QDBusInterface {
	return p
}

func (p *QDBusInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (p *QDBusInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDBusAbstractInterface_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusInterface(ptr QDBusInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusInterfaceFromPointer(ptr unsafe.Pointer) *QDBusInterface {
	var n = new(QDBusInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDBusInterface(service string, path string, interfa string, connection QDBusConnection_ITF, parent core.QObject_ITF) *QDBusInterface {
	defer qt.Recovering("QDBusInterface::QDBusInterface")

	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	return NewQDBusInterfaceFromPointer(C.QDBusInterface_NewQDBusInterface(serviceC, pathC, interfaC, PointerFromQDBusConnection(connection), core.PointerFromQObject(parent)))
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	defer qt.Recovering("QDBusInterface::~QDBusInterface")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusInterface_TimerEvent
func callbackQDBusInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusInterface_ChildEvent
func callbackQDBusInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusInterface_ConnectNotify
func callbackQDBusInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusInterface_CustomEvent
func callbackQDBusInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusInterface_DeleteLater
func callbackQDBusInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusInterface) DeleteLater() {
	defer qt.Recovering("QDBusInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDBusInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusInterface_DisconnectNotify
func callbackQDBusInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusInterface_Event
func callbackQDBusInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusInterface::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusInterface::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusInterface) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusInterface::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusInterface) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusInterface::event")

	if ptr.Pointer() != nil {
		return C.QDBusInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusInterface) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusInterface::event")

	if ptr.Pointer() != nil {
		return C.QDBusInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusInterface_EventFilter
func callbackQDBusInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusInterface_MetaObject
func callbackQDBusInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusInterface(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusInterface(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QDBusMessage::MessageType
type QDBusMessage__MessageType int64

const (
	QDBusMessage__InvalidMessage    = QDBusMessage__MessageType(0)
	QDBusMessage__MethodCallMessage = QDBusMessage__MessageType(1)
	QDBusMessage__ReplyMessage      = QDBusMessage__MessageType(2)
	QDBusMessage__ErrorMessage      = QDBusMessage__MessageType(3)
	QDBusMessage__SignalMessage     = QDBusMessage__MessageType(4)
)

type QDBusMessage struct {
	ptr unsafe.Pointer
}

type QDBusMessage_ITF interface {
	QDBusMessage_PTR() *QDBusMessage
}

func (p *QDBusMessage) QDBusMessage_PTR() *QDBusMessage {
	return p
}

func (p *QDBusMessage) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusMessage) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusMessage(ptr QDBusMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusMessage_PTR().Pointer()
	}
	return nil
}

func NewQDBusMessageFromPointer(ptr unsafe.Pointer) *QDBusMessage {
	var n = new(QDBusMessage)
	n.SetPointer(ptr)
	return n
}
func (ptr *QDBusMessage) CreateErrorReply3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createErrorReply")

	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply3(ptr.Pointer(), C.longlong(ty), msgC))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func NewQDBusMessage() *QDBusMessage {
	defer qt.Recovering("QDBusMessage::QDBusMessage")

	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage())
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func NewQDBusMessage2(other QDBusMessage_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::QDBusMessage")

	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage2(PointerFromQDBusMessage(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) AutoStartService() bool {
	defer qt.Recovering("QDBusMessage::autoStartService")

	if ptr.Pointer() != nil {
		return C.QDBusMessage_AutoStartService(ptr.Pointer()) != 0
	}
	return false
}

func QDBusMessage_CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createError")

	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError3(C.longlong(ty), msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createError")

	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError3(C.longlong(ty), msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateError2(error QDBusError_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createError")

	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError2(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError2(error QDBusError_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createError")

	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError2(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateError(name string, msg string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createError")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError(nameC, msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError(name string, msg string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createError")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError(nameC, msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateErrorReply2(error QDBusError_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createErrorReply")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply2(ptr.Pointer(), PointerFromQDBusError(error)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) CreateErrorReply(name string, msg string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createErrorReply")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply(ptr.Pointer(), nameC, msgC))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func QDBusMessage_CreateMethodCall(service string, path string, interfa string, method string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createMethodCall")

	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var methodC = C.CString(method)
	defer C.free(unsafe.Pointer(methodC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateMethodCall(serviceC, pathC, interfaC, methodC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateMethodCall(service string, path string, interfa string, method string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createMethodCall")

	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var methodC = C.CString(method)
	defer C.free(unsafe.Pointer(methodC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateMethodCall(serviceC, pathC, interfaC, methodC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateReply2(argument core.QVariant_ITF) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createReply")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_CreateReply2(ptr.Pointer(), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func QDBusMessage_CreateSignal(path string, interfa string, name string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createSignal")

	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateSignal(pathC, interfaC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateSignal(path string, interfa string, name string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createSignal")

	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateSignal(pathC, interfaC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateTargetedSignal(service string, path string, interfa string, name string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createTargetedSignal")

	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateTargetedSignal(serviceC, pathC, interfaC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateTargetedSignal(service string, path string, interfa string, name string) *QDBusMessage {
	defer qt.Recovering("QDBusMessage::createTargetedSignal")

	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateTargetedSignal(serviceC, pathC, interfaC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) ErrorMessage() string {
	defer qt.Recovering("QDBusMessage::errorMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) ErrorName() string {
	defer qt.Recovering("QDBusMessage::errorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Interface() string {
	defer qt.Recovering("QDBusMessage::interface")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) IsDelayedReply() bool {
	defer qt.Recovering("QDBusMessage::isDelayedReply")

	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsReplyRequired() bool {
	defer qt.Recovering("QDBusMessage::isReplyRequired")

	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsReplyRequired(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) Member() string {
	defer qt.Recovering("QDBusMessage::member")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Member(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Path() string {
	defer qt.Recovering("QDBusMessage::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Service() string {
	defer qt.Recovering("QDBusMessage::service")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {
	defer qt.Recovering("QDBusMessage::setAutoStartService")

	if ptr.Pointer() != nil {
		C.QDBusMessage_SetAutoStartService(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {
	defer qt.Recovering("QDBusMessage::setDelayedReply")

	if ptr.Pointer() != nil {
		C.QDBusMessage_SetDelayedReply(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) Signature() string {
	defer qt.Recovering("QDBusMessage::signature")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Swap(other QDBusMessage_ITF) {
	defer qt.Recovering("QDBusMessage::swap")

	if ptr.Pointer() != nil {
		C.QDBusMessage_Swap(ptr.Pointer(), PointerFromQDBusMessage(other))
	}
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {
	defer qt.Recovering("QDBusMessage::type")

	if ptr.Pointer() != nil {
		return QDBusMessage__MessageType(C.QDBusMessage_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {
	defer qt.Recovering("QDBusMessage::~QDBusMessage")

	if ptr.Pointer() != nil {
		C.QDBusMessage_DestroyQDBusMessage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDBusObjectPath struct {
	ptr unsafe.Pointer
}

type QDBusObjectPath_ITF interface {
	QDBusObjectPath_PTR() *QDBusObjectPath
}

func (p *QDBusObjectPath) QDBusObjectPath_PTR() *QDBusObjectPath {
	return p
}

func (p *QDBusObjectPath) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusObjectPath) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusObjectPath(ptr QDBusObjectPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusObjectPath_PTR().Pointer()
	}
	return nil
}

func NewQDBusObjectPathFromPointer(ptr unsafe.Pointer) *QDBusObjectPath {
	var n = new(QDBusObjectPath)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusObjectPath) DestroyQDBusObjectPath() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQDBusObjectPath() *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath())
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath3(path core.QLatin1String_ITF) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath3(core.PointerFromQLatin1String(path)))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath5(p string) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	var pC = C.CString(p)
	defer C.free(unsafe.Pointer(pC))
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath5(pC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath4(pathC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {
	defer qt.Recovering("QDBusObjectPath::QDBusObjectPath")

	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath2(pathC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func (ptr *QDBusObjectPath) Path() string {
	defer qt.Recovering("QDBusObjectPath::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusObjectPath_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusObjectPath) SetPath(path string) {
	defer qt.Recovering("QDBusObjectPath::setPath")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QDBusObjectPath_SetPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QDBusObjectPath) Swap(other QDBusObjectPath_ITF) {
	defer qt.Recovering("QDBusObjectPath::swap")

	if ptr.Pointer() != nil {
		C.QDBusObjectPath_Swap(ptr.Pointer(), PointerFromQDBusObjectPath(other))
	}
}

type QDBusPendingCall struct {
	ptr unsafe.Pointer
}

type QDBusPendingCall_ITF interface {
	QDBusPendingCall_PTR() *QDBusPendingCall
}

func (p *QDBusPendingCall) QDBusPendingCall_PTR() *QDBusPendingCall {
	return p
}

func (p *QDBusPendingCall) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusPendingCall) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusPendingCall(ptr QDBusPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallFromPointer(ptr unsafe.Pointer) *QDBusPendingCall {
	var n = new(QDBusPendingCall)
	n.SetPointer(ptr)
	return n
}
func NewQDBusPendingCall(other QDBusPendingCall_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusPendingCall::QDBusPendingCall")

	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_NewQDBusPendingCall(PointerFromQDBusPendingCall(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func QDBusPendingCall_FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusPendingCall::fromCompletedCall")

	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromCompletedCall(PointerFromQDBusMessage(msg)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusPendingCall::fromCompletedCall")

	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromCompletedCall(PointerFromQDBusMessage(msg)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func QDBusPendingCall_FromError(error QDBusError_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusPendingCall::fromError")

	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromError(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) FromError(error QDBusError_ITF) *QDBusPendingCall {
	defer qt.Recovering("QDBusPendingCall::fromError")

	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromError(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCall_ITF) {
	defer qt.Recovering("QDBusPendingCall::swap")

	if ptr.Pointer() != nil {
		C.QDBusPendingCall_Swap(ptr.Pointer(), PointerFromQDBusPendingCall(other))
	}
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {
	defer qt.Recovering("QDBusPendingCall::~QDBusPendingCall")

	if ptr.Pointer() != nil {
		C.QDBusPendingCall_DestroyQDBusPendingCall(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDBusPendingCallWatcher struct {
	core.QObject
	QDBusPendingCall
}

type QDBusPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QDBusPendingCall_ITF
	QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher
}

func (p *QDBusPendingCallWatcher) QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher {
	return p
}

func (p *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDBusPendingCallWatcher) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QDBusPendingCall_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) *QDBusPendingCallWatcher {
	var n = new(QDBusPendingCallWatcher)
	n.SetPointer(ptr)
	return n
}
func (ptr *QDBusPendingCallWatcher) WaitForFinished() {
	defer qt.Recovering("QDBusPendingCallWatcher::waitForFinished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {
	defer qt.Recovering("QDBusPendingCallWatcher::QDBusPendingCallWatcher")

	return NewQDBusPendingCallWatcherFromPointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(PointerFromQDBusPendingCall(call), core.PointerFromQObject(parent)))
}

//export callbackQDBusPendingCallWatcher_Finished
func callbackQDBusPendingCallWatcher_Finished(ptr unsafe.Pointer, self unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::finished")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "finished"); signal != nil {
		signal.(func(*QDBusPendingCallWatcher))(NewQDBusPendingCallWatcherFromPointer(self))
	}

}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self *QDBusPendingCallWatcher)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "finished", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "finished")
	}
}

func (ptr *QDBusPendingCallWatcher) Finished(self QDBusPendingCallWatcher_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_Finished(ptr.Pointer(), PointerFromQDBusPendingCallWatcher(self))
	}
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	defer qt.Recovering("QDBusPendingCallWatcher::isFinished")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	defer qt.Recovering("QDBusPendingCallWatcher::~QDBusPendingCallWatcher")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()))
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusPendingCallWatcher_TimerEvent
func callbackQDBusPendingCallWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_ChildEvent
func callbackQDBusPendingCallWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_ConnectNotify
func callbackQDBusPendingCallWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusPendingCallWatcher_CustomEvent
func callbackQDBusPendingCallWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_DeleteLater
func callbackQDBusPendingCallWatcher_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusPendingCallWatcher) DeleteLater() {
	defer qt.Recovering("QDBusPendingCallWatcher::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()))
		C.QDBusPendingCallWatcher_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusPendingCallWatcher) DeleteLaterDefault() {
	defer qt.Recovering("QDBusPendingCallWatcher::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()))
		C.QDBusPendingCallWatcher_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusPendingCallWatcher_DisconnectNotify
func callbackQDBusPendingCallWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusPendingCallWatcher_Event
func callbackQDBusPendingCallWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusPendingCallWatcher::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusPendingCallWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusPendingCallWatcher) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusPendingCallWatcher) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusPendingCallWatcher::event")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusPendingCallWatcher::event")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_EventFilter
func callbackQDBusPendingCallWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusPendingCallWatcher::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusPendingCallWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusPendingCallWatcher) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusPendingCallWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusPendingCallWatcher::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusPendingCallWatcher::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_MetaObject
func callbackQDBusPendingCallWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusPendingCallWatcher::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusPendingCallWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusPendingCallWatcher) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusPendingCallWatcher(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusPendingCallWatcher) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusPendingCallWatcher::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusPendingCallWatcher_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusPendingCallWatcher::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusPendingCallWatcher_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDBusPendingReply struct {
	QDBusPendingCall
}

type QDBusPendingReply_ITF interface {
	QDBusPendingCall_ITF
	QDBusPendingReply_PTR() *QDBusPendingReply
}

func (p *QDBusPendingReply) QDBusPendingReply_PTR() *QDBusPendingReply {
	return p
}

func (p *QDBusPendingReply) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func (p *QDBusPendingReply) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QDBusPendingCall_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusPendingReply(ptr QDBusPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingReplyFromPointer(ptr unsafe.Pointer) *QDBusPendingReply {
	var n = new(QDBusPendingReply)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusPendingReply) DestroyQDBusPendingReply() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

type QDBusReply struct {
	ptr unsafe.Pointer
}

type QDBusReply_ITF interface {
	QDBusReply_PTR() *QDBusReply
}

func (p *QDBusReply) QDBusReply_PTR() *QDBusReply {
	return p
}

func (p *QDBusReply) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusReply) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusReply(ptr QDBusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusReplyFromPointer(ptr unsafe.Pointer) *QDBusReply {
	var n = new(QDBusReply)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusReply) DestroyQDBusReply() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

type QDBusServer struct {
	core.QObject
}

type QDBusServer_ITF interface {
	core.QObject_ITF
	QDBusServer_PTR() *QDBusServer
}

func (p *QDBusServer) QDBusServer_PTR() *QDBusServer {
	return p
}

func (p *QDBusServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDBusServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusServer(ptr QDBusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServer_PTR().Pointer()
	}
	return nil
}

func NewQDBusServerFromPointer(ptr unsafe.Pointer) *QDBusServer {
	var n = new(QDBusServer)
	n.SetPointer(ptr)
	return n
}
func NewQDBusServer2(parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer2(core.PointerFromQObject(parent)))
}

func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	return NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer(addressC, core.PointerFromQObject(parent)))
}

func (ptr *QDBusServer) Address() string {
	defer qt.Recovering("QDBusServer::address")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusServer_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	defer qt.Recovering("QDBusServer::isAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsAnonymousAuthenticationAllowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	defer qt.Recovering("QDBusServer::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) LastError() *QDBusError {
	defer qt.Recovering("QDBusServer::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusErrorFromPointer(C.QDBusServer_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

//export callbackQDBusServer_NewConnection
func callbackQDBusServer_NewConnection(ptr unsafe.Pointer, connection unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::newConnection")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "newConnection"); signal != nil {
		signal.(func(*QDBusConnection))(NewQDBusConnectionFromPointer(connection))
	}

}

func (ptr *QDBusServer) ConnectNewConnection(f func(connection *QDBusConnection)) {
	defer qt.Recovering("connect QDBusServer::newConnection")

	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "newConnection", f)
	}
}

func (ptr *QDBusServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QDBusServer::newConnection")

	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "newConnection")
	}
}

func (ptr *QDBusServer) NewConnection(connection QDBusConnection_ITF) {
	defer qt.Recovering("QDBusServer::newConnection")

	if ptr.Pointer() != nil {
		C.QDBusServer_NewConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	defer qt.Recovering("QDBusServer::setAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	defer qt.Recovering("QDBusServer::~QDBusServer")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()))
		C.QDBusServer_DestroyQDBusServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServer_TimerEvent
func callbackQDBusServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusServer_ChildEvent
func callbackQDBusServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusServer_ConnectNotify
func callbackQDBusServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServer_CustomEvent
func callbackQDBusServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusServer_DeleteLater
func callbackQDBusServer_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusServer) DeleteLater() {
	defer qt.Recovering("QDBusServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()))
		C.QDBusServer_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServer) DeleteLaterDefault() {
	defer qt.Recovering("QDBusServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()))
		C.QDBusServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServer_DisconnectNotify
func callbackQDBusServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServer_Event
func callbackQDBusServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusServer::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServer::event")

	if ptr.Pointer() != nil {
		return C.QDBusServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServer::event")

	if ptr.Pointer() != nil {
		return C.QDBusServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusServer_EventFilter
func callbackQDBusServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusServer::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusServer_MetaObject
func callbackQDBusServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusServer::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServer(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServer(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int64

const (
	QDBusServiceWatcher__WatchForRegistration   = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    = QDBusServiceWatcher__WatchModeFlag(0x03)
)

type QDBusServiceWatcher struct {
	core.QObject
}

type QDBusServiceWatcher_ITF interface {
	core.QObject_ITF
	QDBusServiceWatcher_PTR() *QDBusServiceWatcher
}

func (p *QDBusServiceWatcher) QDBusServiceWatcher_PTR() *QDBusServiceWatcher {
	return p
}

func (p *QDBusServiceWatcher) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDBusServiceWatcher) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) *QDBusServiceWatcher {
	var n = new(QDBusServiceWatcher)
	n.SetPointer(ptr)
	return n
}
func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {
	defer qt.Recovering("QDBusServiceWatcher::setWatchMode")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	defer qt.Recovering("QDBusServiceWatcher::watchMode")

	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(ptr.Pointer()))
	}
	return 0
}

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {
	defer qt.Recovering("QDBusServiceWatcher::QDBusServiceWatcher")

	return NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(core.PointerFromQObject(parent)))
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {
	defer qt.Recovering("QDBusServiceWatcher::QDBusServiceWatcher")

	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	return NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(serviceC, PointerFromQDBusConnection(connection), C.longlong(watchMode), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	defer qt.Recovering("QDBusServiceWatcher::addWatchedService")

	if ptr.Pointer() != nil {
		var newServiceC = C.CString(newService)
		defer C.free(unsafe.Pointer(newServiceC))
		C.QDBusServiceWatcher_AddWatchedService(ptr.Pointer(), newServiceC)
	}
}

func (ptr *QDBusServiceWatcher) Connection() *QDBusConnection {
	defer qt.Recovering("QDBusServiceWatcher::connection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionFromPointer(C.QDBusServiceWatcher_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	defer qt.Recovering("QDBusServiceWatcher::removeWatchedService")

	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		return C.QDBusServiceWatcher_RemoveWatchedService(ptr.Pointer(), serviceC) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_ServiceOwnerChanged
func callbackQDBusServiceWatcher_ServiceOwnerChanged(ptr unsafe.Pointer, serviceName *C.char, oldOwner *C.char, newOwner *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceOwnerChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "serviceOwnerChanged"); signal != nil {
		signal.(func(string, string, string))(C.GoString(serviceName), C.GoString(oldOwner), C.GoString(newOwner))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceOwnerChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "serviceOwnerChanged", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "serviceOwnerChanged")
	}
}

func (ptr *QDBusServiceWatcher) ServiceOwnerChanged(serviceName string, oldOwner string, newOwner string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		var oldOwnerC = C.CString(oldOwner)
		defer C.free(unsafe.Pointer(oldOwnerC))
		var newOwnerC = C.CString(newOwner)
		defer C.free(unsafe.Pointer(newOwnerC))
		C.QDBusServiceWatcher_ServiceOwnerChanged(ptr.Pointer(), serviceNameC, oldOwnerC, newOwnerC)
	}
}

//export callbackQDBusServiceWatcher_ServiceRegistered
func callbackQDBusServiceWatcher_ServiceRegistered(ptr unsafe.Pointer, serviceName *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceRegistered")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "serviceRegistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "serviceRegistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "serviceRegistered")
	}
}

func (ptr *QDBusServiceWatcher) ServiceRegistered(serviceName string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusServiceWatcher_ServiceRegistered(ptr.Pointer(), serviceNameC)
	}
}

//export callbackQDBusServiceWatcher_ServiceUnregistered
func callbackQDBusServiceWatcher_ServiceUnregistered(ptr unsafe.Pointer, serviceName *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceUnregistered")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "serviceUnregistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "serviceUnregistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "serviceUnregistered")
	}
}

func (ptr *QDBusServiceWatcher) ServiceUnregistered(serviceName string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusServiceWatcher_ServiceUnregistered(ptr.Pointer(), serviceNameC)
	}
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::setConnection")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	defer qt.Recovering("QDBusServiceWatcher::setWatchedServices")

	if ptr.Pointer() != nil {
		var servicesC = C.CString(strings.Join(services, "|"))
		defer C.free(unsafe.Pointer(servicesC))
		C.QDBusServiceWatcher_SetWatchedServices(ptr.Pointer(), servicesC)
	}
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	defer qt.Recovering("QDBusServiceWatcher::watchedServices")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDBusServiceWatcher_WatchedServices(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	defer qt.Recovering("QDBusServiceWatcher::~QDBusServiceWatcher")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()))
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServiceWatcher_TimerEvent
func callbackQDBusServiceWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusServiceWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusServiceWatcher_ChildEvent
func callbackQDBusServiceWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusServiceWatcher) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusServiceWatcher_ConnectNotify
func callbackQDBusServiceWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusServiceWatcher::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusServiceWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServiceWatcher_CustomEvent
func callbackQDBusServiceWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusServiceWatcher) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusServiceWatcher_DeleteLater
func callbackQDBusServiceWatcher_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusServiceWatcher) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusServiceWatcher::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusServiceWatcher) DeleteLater() {
	defer qt.Recovering("QDBusServiceWatcher::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()))
		C.QDBusServiceWatcher_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServiceWatcher) DeleteLaterDefault() {
	defer qt.Recovering("QDBusServiceWatcher::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()))
		C.QDBusServiceWatcher_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServiceWatcher_DisconnectNotify
func callbackQDBusServiceWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusServiceWatcher::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusServiceWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServiceWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServiceWatcher_Event
func callbackQDBusServiceWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusServiceWatcher::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServiceWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusServiceWatcher) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusServiceWatcher::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusServiceWatcher) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServiceWatcher::event")

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServiceWatcher::event")

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_EventFilter
func callbackQDBusServiceWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusServiceWatcher::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServiceWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusServiceWatcher) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusServiceWatcher::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusServiceWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServiceWatcher::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusServiceWatcher::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_MetaObject
func callbackQDBusServiceWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusServiceWatcher::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusServiceWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusServiceWatcher) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusServiceWatcher::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusServiceWatcher(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusServiceWatcher) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusServiceWatcher::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServiceWatcher_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusServiceWatcher) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusServiceWatcher::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServiceWatcher_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDBusSignature struct {
	ptr unsafe.Pointer
}

type QDBusSignature_ITF interface {
	QDBusSignature_PTR() *QDBusSignature
}

func (p *QDBusSignature) QDBusSignature_PTR() *QDBusSignature {
	return p
}

func (p *QDBusSignature) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusSignature) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusSignature(ptr QDBusSignature_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusSignature_PTR().Pointer()
	}
	return nil
}

func NewQDBusSignatureFromPointer(ptr unsafe.Pointer) *QDBusSignature {
	var n = new(QDBusSignature)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusSignature) DestroyQDBusSignature() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQDBusSignature() *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature())
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature3(signature core.QLatin1String_ITF) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature3(core.PointerFromQLatin1String(signature)))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature5(sig string) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	var sigC = C.CString(sig)
	defer C.free(unsafe.Pointer(sigC))
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature5(sigC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature4(signature string) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature4(signatureC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature2(signature string) *QDBusSignature {
	defer qt.Recovering("QDBusSignature::QDBusSignature")

	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature2(signatureC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func (ptr *QDBusSignature) SetSignature(signature string) {
	defer qt.Recovering("QDBusSignature::setSignature")

	if ptr.Pointer() != nil {
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		C.QDBusSignature_SetSignature(ptr.Pointer(), signatureC)
	}
}

func (ptr *QDBusSignature) Signature() string {
	defer qt.Recovering("QDBusSignature::signature")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusSignature_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusSignature) Swap(other QDBusSignature_ITF) {
	defer qt.Recovering("QDBusSignature::swap")

	if ptr.Pointer() != nil {
		C.QDBusSignature_Swap(ptr.Pointer(), PointerFromQDBusSignature(other))
	}
}

type QDBusUnixFileDescriptor struct {
	ptr unsafe.Pointer
}

type QDBusUnixFileDescriptor_ITF interface {
	QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor
}

func (p *QDBusUnixFileDescriptor) QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor {
	return p
}

func (p *QDBusUnixFileDescriptor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusUnixFileDescriptor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusUnixFileDescriptor(ptr QDBusUnixFileDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusUnixFileDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) *QDBusUnixFileDescriptor {
	var n = new(QDBusUnixFileDescriptor)
	n.SetPointer(ptr)
	return n
}
func NewQDBusUnixFileDescriptor() *QDBusUnixFileDescriptor {
	defer qt.Recovering("QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")

	var tmpValue = NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor())
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {
	defer qt.Recovering("QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")

	var tmpValue = NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(PointerFromQDBusUnixFileDescriptor(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	defer qt.Recovering("QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")

	var tmpValue = NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(int32(fileDescriptor))))
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	defer qt.Recovering("QDBusUnixFileDescriptor::fileDescriptor")

	if ptr.Pointer() != nil {
		return int(int32(C.QDBusUnixFileDescriptor_FileDescriptor(ptr.Pointer())))
	}
	return 0
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	defer qt.Recovering("QDBusUnixFileDescriptor::isSupported")

	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsSupported() bool {
	defer qt.Recovering("QDBusUnixFileDescriptor::isSupported")

	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	defer qt.Recovering("QDBusUnixFileDescriptor::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusUnixFileDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	defer qt.Recovering("QDBusUnixFileDescriptor::setFileDescriptor")

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(ptr.Pointer(), C.int(int32(fileDescriptor)))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {
	defer qt.Recovering("QDBusUnixFileDescriptor::swap")

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(ptr.Pointer(), PointerFromQDBusUnixFileDescriptor(other))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
	defer qt.Recovering("QDBusUnixFileDescriptor::~QDBusUnixFileDescriptor")

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDBusVariant struct {
	ptr unsafe.Pointer
}

type QDBusVariant_ITF interface {
	QDBusVariant_PTR() *QDBusVariant
}

func (p *QDBusVariant) QDBusVariant_PTR() *QDBusVariant {
	return p
}

func (p *QDBusVariant) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDBusVariant) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDBusVariant(ptr QDBusVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVariant_PTR().Pointer()
	}
	return nil
}

func NewQDBusVariantFromPointer(ptr unsafe.Pointer) *QDBusVariant {
	var n = new(QDBusVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusVariant) DestroyQDBusVariant() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQDBusVariant() *QDBusVariant {
	defer qt.Recovering("QDBusVariant::QDBusVariant")

	var tmpValue = NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant())
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func NewQDBusVariant3(v core.QVariant_ITF) *QDBusVariant {
	defer qt.Recovering("QDBusVariant::QDBusVariant")

	var tmpValue = NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant3(core.PointerFromQVariant(v)))
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func NewQDBusVariant2(variant core.QVariant_ITF) *QDBusVariant {
	defer qt.Recovering("QDBusVariant::QDBusVariant")

	var tmpValue = NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant2(core.PointerFromQVariant(variant)))
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func (ptr *QDBusVariant) SetVariant(variant core.QVariant_ITF) {
	defer qt.Recovering("QDBusVariant::setVariant")

	if ptr.Pointer() != nil {
		C.QDBusVariant_SetVariant(ptr.Pointer(), core.PointerFromQVariant(variant))
	}
}

func (ptr *QDBusVariant) Swap(other QDBusVariant_ITF) {
	defer qt.Recovering("QDBusVariant::swap")

	if ptr.Pointer() != nil {
		C.QDBusVariant_Swap(ptr.Pointer(), PointerFromQDBusVariant(other))
	}
}

func (ptr *QDBusVariant) Variant() *core.QVariant {
	defer qt.Recovering("QDBusVariant::variant")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDBusVariant_Variant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObject_ITF interface {
	core.QObject_ITF
	QDBusVirtualObject_PTR() *QDBusVirtualObject
}

func (p *QDBusVirtualObject) QDBusVirtualObject_PTR() *QDBusVirtualObject {
	return p
}

func (p *QDBusVirtualObject) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDBusVirtualObject) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObject_PTR().Pointer()
	}
	return nil
}

func NewQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) *QDBusVirtualObject {
	var n = new(QDBusVirtualObject)
	n.SetPointer(ptr)
	return n
}
func NewQDBusVirtualObject(parent core.QObject_ITF) *QDBusVirtualObject {
	defer qt.Recovering("QDBusVirtualObject::QDBusVirtualObject")

	return NewQDBusVirtualObjectFromPointer(C.QDBusVirtualObject_NewQDBusVirtualObject(core.PointerFromQObject(parent)))
}

//export callbackQDBusVirtualObject_HandleMessage
func callbackQDBusVirtualObject_HandleMessage(ptr unsafe.Pointer, message unsafe.Pointer, connection unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusVirtualObject::handleMessage")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "handleMessage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QDBusMessage, *QDBusConnection) bool)(NewQDBusMessageFromPointer(message), NewQDBusConnectionFromPointer(connection)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDBusVirtualObject) ConnectHandleMessage(f func(message *QDBusMessage, connection *QDBusConnection) bool) {
	defer qt.Recovering("connect QDBusVirtualObject::handleMessage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "handleMessage", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectHandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) {
	defer qt.Recovering("disconnect QDBusVirtualObject::handleMessage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "handleMessage")
	}
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::handleMessage")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_HandleMessage(ptr.Pointer(), PointerFromQDBusMessage(message), PointerFromQDBusConnection(connection)) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_Introspect
func callbackQDBusVirtualObject_Introspect(ptr unsafe.Pointer, path *C.char) *C.char {
	defer qt.Recovering("callback QDBusVirtualObject::introspect")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "introspect"); signal != nil {
		return C.CString(signal.(func(string) string)(C.GoString(path)))
	}

	return C.CString("")
}

func (ptr *QDBusVirtualObject) ConnectIntrospect(f func(path string) string) {
	defer qt.Recovering("connect QDBusVirtualObject::introspect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "introspect", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectIntrospect(path string) {
	defer qt.Recovering("disconnect QDBusVirtualObject::introspect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "introspect")
	}
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	defer qt.Recovering("QDBusVirtualObject::introspect")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		return C.GoString(C.QDBusVirtualObject_Introspect(ptr.Pointer(), pathC))
	}
	return ""
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	defer qt.Recovering("QDBusVirtualObject::~QDBusVirtualObject")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()))
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusVirtualObject_TimerEvent
func callbackQDBusVirtualObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QDBusVirtualObject) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusVirtualObject_ChildEvent
func callbackQDBusVirtualObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QDBusVirtualObject) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusVirtualObject_ConnectNotify
func callbackQDBusVirtualObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusVirtualObject::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDBusVirtualObject::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QDBusVirtualObject) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusVirtualObject::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusVirtualObject::connectNotify")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusVirtualObject_CustomEvent
func callbackQDBusVirtualObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QDBusVirtualObject) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusVirtualObject_DeleteLater
func callbackQDBusVirtualObject_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusVirtualObject) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDBusVirtualObject::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDBusVirtualObject::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QDBusVirtualObject) DeleteLater() {
	defer qt.Recovering("QDBusVirtualObject::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()))
		C.QDBusVirtualObject_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusVirtualObject) DeleteLaterDefault() {
	defer qt.Recovering("QDBusVirtualObject::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()))
		C.QDBusVirtualObject_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusVirtualObject_DisconnectNotify
func callbackQDBusVirtualObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDBusVirtualObject::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDBusVirtualObject::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QDBusVirtualObject) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusVirtualObject::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusVirtualObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDBusVirtualObject::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusVirtualObject_Event
func callbackQDBusVirtualObject_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusVirtualObject::event")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusVirtualObjectFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusVirtualObject) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusVirtualObject::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QDBusVirtualObject) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::event")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::event")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_EventFilter
func callbackQDBusVirtualObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDBusVirtualObject::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusVirtualObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusVirtualObject) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDBusVirtualObject::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDBusVirtualObject::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QDBusVirtualObject) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_MetaObject
func callbackQDBusVirtualObject_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDBusVirtualObject::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusVirtualObjectFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusVirtualObject) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDBusVirtualObject::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDBusVirtualObject::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QDBusVirtualObject(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QDBusVirtualObject) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDBusVirtualObject::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusVirtualObject_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusVirtualObject) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDBusVirtualObject::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusVirtualObject_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
