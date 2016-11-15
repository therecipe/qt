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

func cGoUnpackString(s C.struct_QtDBus_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

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
	var tmpValue = NewQDBusAbstractAdaptorFromPointer(C.QDBusAbstractAdaptor_NewQDBusAbstractAdaptor(core.PointerFromQObject(obj)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDBusAbstractAdaptor) AutoRelaySignals() bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_AutoRelaySignals(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractAdaptor) SetAutoRelaySignals(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_SetAutoRelaySignals(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractAdaptor_TimerEvent
func callbackQDBusAbstractAdaptor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::timerEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::timerEvent")
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_ChildEvent
func callbackQDBusAbstractAdaptor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::childEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::childEvent")
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_ConnectNotify
func callbackQDBusAbstractAdaptor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::connectNotify", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::connectNotify")
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractAdaptor_CustomEvent
func callbackQDBusAbstractAdaptor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::customEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::customEvent")
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_DeleteLater
func callbackQDBusAbstractAdaptor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::deleteLater", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::deleteLater")
	}
}

func (ptr *QDBusAbstractAdaptor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractAdaptor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractAdaptor_DisconnectNotify
func callbackQDBusAbstractAdaptor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::disconnectNotify", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::disconnectNotify")
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractAdaptor_Event
func callbackQDBusAbstractAdaptor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractAdaptorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusAbstractAdaptor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::event", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::event")
	}
}

func (ptr *QDBusAbstractAdaptor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusAbstractAdaptor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_EventFilter
func callbackQDBusAbstractAdaptor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractAdaptorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusAbstractAdaptor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::eventFilter", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::eventFilter")
	}
}

func (ptr *QDBusAbstractAdaptor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusAbstractAdaptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractAdaptor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_MetaObject
func callbackQDBusAbstractAdaptor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractAdaptor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusAbstractAdaptorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusAbstractAdaptor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::metaObject", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractAdaptor::metaObject")
	}
}

func (ptr *QDBusAbstractAdaptor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractAdaptor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) MetaObjectDefault() *core.QMetaObject {
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
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionFromPointer(C.QDBusAbstractInterface_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Interface() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) LastError() *QDBusError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusErrorFromPointer(C.QDBusAbstractInterface_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(int32(timeout)))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDBusAbstractInterface_Timeout(ptr.Pointer())))
	}
	return 0
}

//export callbackQDBusAbstractInterface_DestroyQDBusAbstractInterface
func callbackQDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::~QDBusAbstractInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DestroyQDBusAbstractInterfaceDefault()
	}
}

func (ptr *QDBusAbstractInterface) ConnectDestroyQDBusAbstractInterface(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::~QDBusAbstractInterface", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectDestroyQDBusAbstractInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::~QDBusAbstractInterface")
	}
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterfaceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractInterface_TimerEvent
func callbackQDBusAbstractInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::timerEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::timerEvent")
	}
}

func (ptr *QDBusAbstractInterface) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusAbstractInterface_ChildEvent
func callbackQDBusAbstractInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::childEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::childEvent")
	}
}

func (ptr *QDBusAbstractInterface) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusAbstractInterface_ConnectNotify
func callbackQDBusAbstractInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::connectNotify", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::connectNotify")
	}
}

func (ptr *QDBusAbstractInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractInterface_CustomEvent
func callbackQDBusAbstractInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::customEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::customEvent")
	}
}

func (ptr *QDBusAbstractInterface) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusAbstractInterface_DeleteLater
func callbackQDBusAbstractInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusAbstractInterface) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::deleteLater", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::deleteLater")
	}
}

func (ptr *QDBusAbstractInterface) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusAbstractInterface_DisconnectNotify
func callbackQDBusAbstractInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::disconnectNotify", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::disconnectNotify")
	}
}

func (ptr *QDBusAbstractInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusAbstractInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractInterface_Event
func callbackQDBusAbstractInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusAbstractInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::event", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::event")
	}
}

func (ptr *QDBusAbstractInterface) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_EventFilter
func callbackQDBusAbstractInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusAbstractInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::eventFilter", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::eventFilter")
	}
}

func (ptr *QDBusAbstractInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_MetaObject
func callbackQDBusAbstractInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusAbstractInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusAbstractInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusAbstractInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::metaObject", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusAbstractInterface::metaObject")
	}
}

func (ptr *QDBusAbstractInterface) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusAbstractInterface) MetaObjectDefault() *core.QMetaObject {
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
	var tmpValue = NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument())
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func NewQDBusArgument3(other QDBusArgument_ITF) *QDBusArgument {
	var tmpValue = NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument3(PointerFromQDBusArgument(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func NewQDBusArgument2(other QDBusArgument_ITF) *QDBusArgument {
	var tmpValue = NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument2(PointerFromQDBusArgument(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func (ptr *QDBusArgument) AsVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDBusArgument_AsVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusArgument) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QDBusArgument_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusArgument) BeginArray(id int) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QDBusArgument) BeginArray2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap(ptr.Pointer(), C.int(int32(kid)), C.int(int32(vid)))
	}
}

func (ptr *QDBusArgument) BeginMap2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {
	if ptr.Pointer() != nil {
		return QDBusArgument__ElementType(C.QDBusArgument_CurrentType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusArgument) EndArray() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndArray2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) Swap(other QDBusArgument_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_Swap(ptr.Pointer(), PointerFromQDBusArgument(other))
	}
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {
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
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SessionBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) SessionBus() *QDBusConnection {
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SessionBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_SystemBus() *QDBusConnection {
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SystemBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) SystemBus() *QDBusConnection {
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SystemBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection3(other QDBusConnection_ITF) *QDBusConnection {
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection3(PointerFromQDBusConnection(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection2(other QDBusConnection_ITF) *QDBusConnection {
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection2(PointerFromQDBusConnection(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection(name string) *QDBusConnection {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection(nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) BaseService() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusConnection_BaseService(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) AsyncCall(message QDBusMessage_ITF, timeout int) *QDBusPendingCall {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusConnection_AsyncCall(ptr.Pointer(), PointerFromQDBusMessage(message), C.int(int32(timeout))))
		runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessage_ITF, receiver core.QObject_ITF, returnMethod string, errorMethod string, timeout int) bool {
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
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus(C.longlong(ty), nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus(C.longlong(ty), nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_ConnectToBus2(address string, name string) *QDBusConnection {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus2(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToBus2(address string, name string) *QDBusConnection {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus2(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_ConnectToPeer(address string, name string) *QDBusConnection {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToPeer(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToPeer(address string, name string) *QDBusConnection {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToPeer(addressC, nameC))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {
	if ptr.Pointer() != nil {
		return QDBusConnection__ConnectionCapability(C.QDBusConnection_ConnectionCapabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusConnection) Disconnect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
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
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(nameC)
}

func (ptr *QDBusConnection) DisconnectFromBus(name string) {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(nameC)
}

func QDBusConnection_DisconnectFromPeer(name string) {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(nameC)
}

func (ptr *QDBusConnection) DisconnectFromPeer(name string) {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(nameC)
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionInterfaceFromPointer(C.QDBusConnection_Interface(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) IsConnected() bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusConnection) LastError() *QDBusError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusErrorFromPointer(C.QDBusConnection_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func QDBusConnection_LocalMachineId() *core.QByteArray {
	var tmpValue = core.NewQByteArrayFromPointer(C.QDBusConnection_QDBusConnection_LocalMachineId())
	runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QDBusConnection) LocalMachineId() *core.QByteArray {
	var tmpValue = core.NewQByteArrayFromPointer(C.QDBusConnection_QDBusConnection_LocalMachineId())
	runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QDBusConnection) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusConnection_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		var tmpValue = core.NewQObjectFromPointer(C.QDBusConnection_ObjectRegisteredAt(ptr.Pointer(), pathC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		return C.QDBusConnection_RegisterObject(ptr.Pointer(), pathC, core.PointerFromQObject(object), C.longlong(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
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
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		return C.QDBusConnection_RegisterService(ptr.Pointer(), serviceNameC) != 0
	}
	return false
}

func (ptr *QDBusConnection) Send(message QDBusMessage_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_Send(ptr.Pointer(), PointerFromQDBusMessage(message)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Swap(other QDBusConnection_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnection_Swap(ptr.Pointer(), PointerFromQDBusConnection(other))
	}
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QDBusConnection_UnregisterObject(ptr.Pointer(), pathC, C.longlong(mode))
	}
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		return C.QDBusConnection_UnregisterService(ptr.Pointer(), serviceNameC) != 0
	}
	return false
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {
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
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQDBusConnectionInterface_CallWithCallbackFailed
func callbackQDBusConnectionInterface_CallWithCallbackFailed(ptr unsafe.Pointer, error unsafe.Pointer, call unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::callWithCallbackFailed"); signal != nil {
		signal.(func(*QDBusError, *QDBusMessage))(NewQDBusErrorFromPointer(error), NewQDBusMessageFromPointer(call))
	}

}

func (ptr *QDBusConnectionInterface) ConnectCallWithCallbackFailed(f func(error *QDBusError, call *QDBusMessage)) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectCallWithCallbackFailed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::callWithCallbackFailed", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCallWithCallbackFailed() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectCallWithCallbackFailed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::callWithCallbackFailed")
	}
}

func (ptr *QDBusConnectionInterface) CallWithCallbackFailed(error QDBusError_ITF, call QDBusMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CallWithCallbackFailed(ptr.Pointer(), PointerFromQDBusError(error), PointerFromQDBusMessage(call))
	}
}

//export callbackQDBusConnectionInterface_ServiceRegistered
func callbackQDBusConnectionInterface_ServiceRegistered(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::serviceRegistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::serviceRegistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::serviceRegistered")
	}
}

func (ptr *QDBusConnectionInterface) ServiceRegistered(serviceName string) {
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusConnectionInterface_ServiceRegistered(ptr.Pointer(), serviceNameC)
	}
}

//export callbackQDBusConnectionInterface_ServiceUnregistered
func callbackQDBusConnectionInterface_ServiceUnregistered(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::serviceUnregistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::serviceUnregistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::serviceUnregistered")
	}
}

func (ptr *QDBusConnectionInterface) ServiceUnregistered(serviceName string) {
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusConnectionInterface_ServiceUnregistered(ptr.Pointer(), serviceNameC)
	}
}

//export callbackQDBusConnectionInterface_TimerEvent
func callbackQDBusConnectionInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::timerEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::timerEvent")
	}
}

func (ptr *QDBusConnectionInterface) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusConnectionInterface_ChildEvent
func callbackQDBusConnectionInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::childEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::childEvent")
	}
}

func (ptr *QDBusConnectionInterface) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusConnectionInterface_ConnectNotify
func callbackQDBusConnectionInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusConnectionInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::connectNotify", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::connectNotify")
	}
}

func (ptr *QDBusConnectionInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusConnectionInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusConnectionInterface_CustomEvent
func callbackQDBusConnectionInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::customEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::customEvent")
	}
}

func (ptr *QDBusConnectionInterface) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusConnectionInterface_DeleteLater
func callbackQDBusConnectionInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusConnectionInterface) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::deleteLater", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::deleteLater")
	}
}

func (ptr *QDBusConnectionInterface) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusConnectionInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusConnectionInterface_DisconnectNotify
func callbackQDBusConnectionInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusConnectionInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::disconnectNotify", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::disconnectNotify")
	}
}

func (ptr *QDBusConnectionInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusConnectionInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusConnectionInterface_Event
func callbackQDBusConnectionInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusConnectionInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusConnectionInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::event", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::event")
	}
}

func (ptr *QDBusConnectionInterface) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusConnectionInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusConnectionInterface_EventFilter
func callbackQDBusConnectionInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusConnectionInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusConnectionInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::eventFilter", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::eventFilter")
	}
}

func (ptr *QDBusConnectionInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusConnectionInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnectionInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusConnectionInterface_MetaObject
func callbackQDBusConnectionInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusConnectionInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusConnectionInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusConnectionInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::metaObject", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusConnectionInterface::metaObject")
	}
}

func (ptr *QDBusConnectionInterface) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusConnectionInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusConnectionInterface) MetaObjectDefault() *core.QMetaObject {
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
	var tmpValue = NewQDBusContextFromPointer(C.QDBusContext_NewQDBusContext())
	runtime.SetFinalizer(tmpValue, (*QDBusContext).DestroyQDBusContext)
	return tmpValue
}

func (ptr *QDBusContext) CalledFromDBus() bool {
	if ptr.Pointer() != nil {
		return C.QDBusContext_CalledFromDBus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) Connection() *QDBusConnection {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionFromPointer(C.QDBusContext_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusContext) IsDelayedReply() bool {
	if ptr.Pointer() != nil {
		return C.QDBusContext_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) Message() *QDBusMessage {
	if ptr.Pointer() != nil {
		return NewQDBusMessageFromPointer(C.QDBusContext_Message(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QDBusContext_SendErrorReply2(ptr.Pointer(), C.longlong(ty), msgC)
	}
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QDBusContext_SendErrorReply(ptr.Pointer(), nameC, msgC)
	}
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusContext_SetDelayedReply(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusContext) DestroyQDBusContext() {
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
	var tmpValue = NewQDBusErrorFromPointer(C.QDBusError_NewQDBusError(PointerFromQDBusError(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
	return tmpValue
}

func QDBusError_ErrorString(error QDBusError__ErrorType) string {
	return cGoUnpackString(C.QDBusError_QDBusError_ErrorString(C.longlong(error)))
}

func (ptr *QDBusError) ErrorString(error QDBusError__ErrorType) string {
	return cGoUnpackString(C.QDBusError_QDBusError_ErrorString(C.longlong(error)))
}

func (ptr *QDBusError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusError) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusError_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusError_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Swap(other QDBusError_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusError_Swap(ptr.Pointer(), PointerFromQDBusError(other))
	}
}

func (ptr *QDBusError) Type() QDBusError__ErrorType {
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
	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var interfaC = C.CString(interfa)
	defer C.free(unsafe.Pointer(interfaC))
	var tmpValue = NewQDBusInterfaceFromPointer(C.QDBusInterface_NewQDBusInterface(serviceC, pathC, interfaC, PointerFromQDBusConnection(connection), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusInterface_TimerEvent
func callbackQDBusInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::timerEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::timerEvent")
	}
}

func (ptr *QDBusInterface) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusInterface_ChildEvent
func callbackQDBusInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::childEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::childEvent")
	}
}

func (ptr *QDBusInterface) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusInterface_ConnectNotify
func callbackQDBusInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::connectNotify", f)
	}
}

func (ptr *QDBusInterface) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::connectNotify")
	}
}

func (ptr *QDBusInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusInterface_CustomEvent
func callbackQDBusInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::customEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::customEvent")
	}
}

func (ptr *QDBusInterface) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusInterface_DeleteLater
func callbackQDBusInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusInterface) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::deleteLater", f)
	}
}

func (ptr *QDBusInterface) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::deleteLater")
	}
}

func (ptr *QDBusInterface) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusInterface_DisconnectNotify
func callbackQDBusInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::disconnectNotify", f)
	}
}

func (ptr *QDBusInterface) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::disconnectNotify")
	}
}

func (ptr *QDBusInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusInterface_Event
func callbackQDBusInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::event", f)
	}
}

func (ptr *QDBusInterface) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::event")
	}
}

func (ptr *QDBusInterface) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusInterface_EventFilter
func callbackQDBusInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::eventFilter", f)
	}
}

func (ptr *QDBusInterface) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::eventFilter")
	}
}

func (ptr *QDBusInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusInterface_MetaObject
func callbackQDBusInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::metaObject", f)
	}
}

func (ptr *QDBusInterface) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusInterface::metaObject")
	}
}

func (ptr *QDBusInterface) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusInterface) MetaObjectDefault() *core.QMetaObject {
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
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage())
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func NewQDBusMessage2(other QDBusMessage_ITF) *QDBusMessage {
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage2(PointerFromQDBusMessage(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) Arguments() []*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDBus_PackedList) []*core.QVariant {
			var out = make([]*core.QVariant, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDBusMessageFromPointer(l.data).arguments_atList(i)
			}
			return out
		}(C.QDBusMessage_Arguments(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusMessage) AutoStartService() bool {
	if ptr.Pointer() != nil {
		return C.QDBusMessage_AutoStartService(ptr.Pointer()) != 0
	}
	return false
}

func QDBusMessage_CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError3(C.longlong(ty), msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError3(C.longlong(ty), msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateError2(error QDBusError_ITF) *QDBusMessage {
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError2(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError2(error QDBusError_ITF) *QDBusMessage {
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError2(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateError(name string, msg string) *QDBusMessage {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError(nameC, msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError(name string, msg string) *QDBusMessage {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var msgC = C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError(nameC, msgC))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateErrorReply2(error QDBusError_ITF) *QDBusMessage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply2(ptr.Pointer(), PointerFromQDBusError(error)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) CreateErrorReply(name string, msg string) *QDBusMessage {
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
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusMessageFromPointer(C.QDBusMessage_CreateReply2(ptr.Pointer(), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func QDBusMessage_CreateSignal(path string, interfa string, name string) *QDBusMessage {
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
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) ErrorName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_ErrorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Interface() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) IsDelayedReply() bool {
	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsReplyRequired() bool {
	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsReplyRequired(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) Member() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Member(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Service() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetAutoStartService(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetDelayedReply(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) Signature() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Swap(other QDBusMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_Swap(ptr.Pointer(), PointerFromQDBusMessage(other))
	}
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {
	if ptr.Pointer() != nil {
		return QDBusMessage__MessageType(C.QDBusMessage_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {
	if ptr.Pointer() != nil {
		C.QDBusMessage_DestroyQDBusMessage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusMessage) arguments_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDBusMessage_arguments_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
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
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath())
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath3(path core.QLatin1String_ITF) *QDBusObjectPath {
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath3(core.PointerFromQLatin1String(path)))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath5(p string) *QDBusObjectPath {
	var pC = C.CString(p)
	defer C.free(unsafe.Pointer(pC))
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath5(pC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath4(pathC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	var tmpValue = NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath2(pathC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func (ptr *QDBusObjectPath) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusObjectPath_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusObjectPath) SetPath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QDBusObjectPath_SetPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QDBusObjectPath) Swap(other QDBusObjectPath_ITF) {
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
	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_NewQDBusPendingCall(PointerFromQDBusPendingCall(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func QDBusPendingCall_FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {
	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromCompletedCall(PointerFromQDBusMessage(msg)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {
	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromCompletedCall(PointerFromQDBusMessage(msg)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func QDBusPendingCall_FromError(error QDBusError_ITF) *QDBusPendingCall {
	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromError(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) FromError(error QDBusError_ITF) *QDBusPendingCall {
	var tmpValue = NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromError(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCall_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCall_Swap(ptr.Pointer(), PointerFromQDBusPendingCall(other))
	}
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {
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
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {
	var tmpValue = NewQDBusPendingCallWatcherFromPointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(PointerFromQDBusPendingCall(call), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDBusPendingCallWatcher_Finished
func callbackQDBusPendingCallWatcher_Finished(ptr unsafe.Pointer, self unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::finished"); signal != nil {
		signal.(func(*QDBusPendingCallWatcher))(NewQDBusPendingCallWatcherFromPointer(self))
	}

}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self *QDBusPendingCallWatcher)) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::finished", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::finished")
	}
}

func (ptr *QDBusPendingCallWatcher) Finished(self QDBusPendingCallWatcher_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_Finished(ptr.Pointer(), PointerFromQDBusPendingCallWatcher(self))
	}
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusPendingCallWatcher_TimerEvent
func callbackQDBusPendingCallWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::timerEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::timerEvent")
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_ChildEvent
func callbackQDBusPendingCallWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::childEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::childEvent")
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_ConnectNotify
func callbackQDBusPendingCallWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::connectNotify", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::connectNotify")
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusPendingCallWatcher_CustomEvent
func callbackQDBusPendingCallWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::customEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::customEvent")
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_DeleteLater
func callbackQDBusPendingCallWatcher_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::deleteLater", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::deleteLater")
	}
}

func (ptr *QDBusPendingCallWatcher) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusPendingCallWatcher) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusPendingCallWatcher_DisconnectNotify
func callbackQDBusPendingCallWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::disconnectNotify", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::disconnectNotify")
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusPendingCallWatcher_Event
func callbackQDBusPendingCallWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusPendingCallWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusPendingCallWatcher) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::event", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::event")
	}
}

func (ptr *QDBusPendingCallWatcher) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_EventFilter
func callbackQDBusPendingCallWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusPendingCallWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusPendingCallWatcher) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::eventFilter", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::eventFilter")
	}
}

func (ptr *QDBusPendingCallWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_MetaObject
func callbackQDBusPendingCallWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusPendingCallWatcher::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusPendingCallWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusPendingCallWatcher) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::metaObject", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusPendingCallWatcher::metaObject")
	}
}

func (ptr *QDBusPendingCallWatcher) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusPendingCallWatcher_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) MetaObjectDefault() *core.QMetaObject {
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
	var tmpValue = NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var tmpValue = NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer(addressC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDBusServer) Address() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusServer_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_IsAnonymousAuthenticationAllowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) LastError() *QDBusError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusErrorFromPointer(C.QDBusServer_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

//export callbackQDBusServer_NewConnection
func callbackQDBusServer_NewConnection(ptr unsafe.Pointer, connection unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::newConnection"); signal != nil {
		signal.(func(*QDBusConnection))(NewQDBusConnectionFromPointer(connection))
	}

}

func (ptr *QDBusServer) ConnectNewConnection(f func(connection *QDBusConnection)) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::newConnection", f)
	}
}

func (ptr *QDBusServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::newConnection")
	}
}

func (ptr *QDBusServer) NewConnection(connection QDBusConnection_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_NewConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQDBusServer_DestroyQDBusServer
func callbackQDBusServer_DestroyQDBusServer(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::~QDBusServer"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServerFromPointer(ptr).DestroyQDBusServerDefault()
	}
}

func (ptr *QDBusServer) ConnectDestroyQDBusServer(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::~QDBusServer", f)
	}
}

func (ptr *QDBusServer) DisconnectDestroyQDBusServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::~QDBusServer")
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServer) DestroyQDBusServerDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServer_TimerEvent
func callbackQDBusServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::timerEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::timerEvent")
	}
}

func (ptr *QDBusServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusServer_ChildEvent
func callbackQDBusServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::childEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::childEvent")
	}
}

func (ptr *QDBusServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusServer_ConnectNotify
func callbackQDBusServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::connectNotify", f)
	}
}

func (ptr *QDBusServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::connectNotify")
	}
}

func (ptr *QDBusServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServer_CustomEvent
func callbackQDBusServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::customEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::customEvent")
	}
}

func (ptr *QDBusServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusServer_DeleteLater
func callbackQDBusServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::deleteLater", f)
	}
}

func (ptr *QDBusServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::deleteLater")
	}
}

func (ptr *QDBusServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServer_DisconnectNotify
func callbackQDBusServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::disconnectNotify", f)
	}
}

func (ptr *QDBusServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::disconnectNotify")
	}
}

func (ptr *QDBusServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServer_Event
func callbackQDBusServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::event", f)
	}
}

func (ptr *QDBusServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::event")
	}
}

func (ptr *QDBusServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusServer_EventFilter
func callbackQDBusServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::eventFilter", f)
	}
}

func (ptr *QDBusServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::eventFilter")
	}
}

func (ptr *QDBusServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusServer_MetaObject
func callbackQDBusServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::metaObject", f)
	}
}

func (ptr *QDBusServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServer::metaObject")
	}
}

func (ptr *QDBusServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusServer) MetaObjectDefault() *core.QMetaObject {
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
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(ptr.Pointer()))
	}
	return 0
}

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {
	var tmpValue = NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {
	var serviceC = C.CString(service)
	defer C.free(unsafe.Pointer(serviceC))
	var tmpValue = NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(serviceC, PointerFromQDBusConnection(connection), C.longlong(watchMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	if ptr.Pointer() != nil {
		var newServiceC = C.CString(newService)
		defer C.free(unsafe.Pointer(newServiceC))
		C.QDBusServiceWatcher_AddWatchedService(ptr.Pointer(), newServiceC)
	}
}

func (ptr *QDBusServiceWatcher) Connection() *QDBusConnection {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDBusConnectionFromPointer(C.QDBusServiceWatcher_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	if ptr.Pointer() != nil {
		var serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
		return C.QDBusServiceWatcher_RemoveWatchedService(ptr.Pointer(), serviceC) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_ServiceOwnerChanged
func callbackQDBusServiceWatcher_ServiceOwnerChanged(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString, oldOwner C.struct_QtDBus_PackedString, newOwner C.struct_QtDBus_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::serviceOwnerChanged"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(serviceName), cGoUnpackString(oldOwner), cGoUnpackString(newOwner))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceOwnerChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::serviceOwnerChanged", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::serviceOwnerChanged")
	}
}

func (ptr *QDBusServiceWatcher) ServiceOwnerChanged(serviceName string, oldOwner string, newOwner string) {
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
func callbackQDBusServiceWatcher_ServiceRegistered(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::serviceRegistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::serviceRegistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::serviceRegistered")
	}
}

func (ptr *QDBusServiceWatcher) ServiceRegistered(serviceName string) {
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusServiceWatcher_ServiceRegistered(ptr.Pointer(), serviceNameC)
	}
}

//export callbackQDBusServiceWatcher_ServiceUnregistered
func callbackQDBusServiceWatcher_ServiceUnregistered(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::serviceUnregistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::serviceUnregistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::serviceUnregistered")
	}
}

func (ptr *QDBusServiceWatcher) ServiceUnregistered(serviceName string) {
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		C.QDBusServiceWatcher_ServiceUnregistered(ptr.Pointer(), serviceNameC)
	}
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	if ptr.Pointer() != nil {
		var servicesC = C.CString(strings.Join(services, "|"))
		defer C.free(unsafe.Pointer(servicesC))
		C.QDBusServiceWatcher_SetWatchedServices(ptr.Pointer(), servicesC)
	}
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QDBusServiceWatcher_WatchedServices(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServiceWatcher_TimerEvent
func callbackQDBusServiceWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::timerEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::timerEvent")
	}
}

func (ptr *QDBusServiceWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusServiceWatcher_ChildEvent
func callbackQDBusServiceWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::childEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::childEvent")
	}
}

func (ptr *QDBusServiceWatcher) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusServiceWatcher_ConnectNotify
func callbackQDBusServiceWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::connectNotify", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::connectNotify")
	}
}

func (ptr *QDBusServiceWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServiceWatcher_CustomEvent
func callbackQDBusServiceWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::customEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::customEvent")
	}
}

func (ptr *QDBusServiceWatcher) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusServiceWatcher_DeleteLater
func callbackQDBusServiceWatcher_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusServiceWatcher) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::deleteLater", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::deleteLater")
	}
}

func (ptr *QDBusServiceWatcher) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServiceWatcher) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusServiceWatcher_DisconnectNotify
func callbackQDBusServiceWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::disconnectNotify", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::disconnectNotify")
	}
}

func (ptr *QDBusServiceWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusServiceWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServiceWatcher_Event
func callbackQDBusServiceWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServiceWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusServiceWatcher) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::event", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::event")
	}
}

func (ptr *QDBusServiceWatcher) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_EventFilter
func callbackQDBusServiceWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServiceWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusServiceWatcher) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::eventFilter", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::eventFilter")
	}
}

func (ptr *QDBusServiceWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_MetaObject
func callbackQDBusServiceWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusServiceWatcher::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusServiceWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusServiceWatcher) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::metaObject", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusServiceWatcher::metaObject")
	}
}

func (ptr *QDBusServiceWatcher) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServiceWatcher_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusServiceWatcher) MetaObjectDefault() *core.QMetaObject {
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
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature())
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature3(signature core.QLatin1String_ITF) *QDBusSignature {
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature3(core.PointerFromQLatin1String(signature)))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature5(sig string) *QDBusSignature {
	var sigC = C.CString(sig)
	defer C.free(unsafe.Pointer(sigC))
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature5(sigC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature4(signature string) *QDBusSignature {
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature4(signatureC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature2(signature string) *QDBusSignature {
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature2(signatureC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func (ptr *QDBusSignature) SetSignature(signature string) {
	if ptr.Pointer() != nil {
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		C.QDBusSignature_SetSignature(ptr.Pointer(), signatureC)
	}
}

func (ptr *QDBusSignature) Signature() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusSignature_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusSignature) Swap(other QDBusSignature_ITF) {
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
	var tmpValue = NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor())
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {
	var tmpValue = NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(PointerFromQDBusUnixFileDescriptor(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	var tmpValue = NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(int32(fileDescriptor))))
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDBusUnixFileDescriptor_FileDescriptor(ptr.Pointer())))
	}
	return 0
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsSupported() bool {
	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusUnixFileDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(ptr.Pointer(), C.int(int32(fileDescriptor)))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(ptr.Pointer(), PointerFromQDBusUnixFileDescriptor(other))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
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
	var tmpValue = NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant())
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func NewQDBusVariant3(v core.QVariant_ITF) *QDBusVariant {
	var tmpValue = NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant3(core.PointerFromQVariant(v)))
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func NewQDBusVariant2(variant core.QVariant_ITF) *QDBusVariant {
	var tmpValue = NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant2(core.PointerFromQVariant(variant)))
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func (ptr *QDBusVariant) SetVariant(variant core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVariant_SetVariant(ptr.Pointer(), core.PointerFromQVariant(variant))
	}
}

func (ptr *QDBusVariant) Swap(other QDBusVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVariant_Swap(ptr.Pointer(), PointerFromQDBusVariant(other))
	}
}

func (ptr *QDBusVariant) Variant() *core.QVariant {
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
	var tmpValue = NewQDBusVirtualObjectFromPointer(C.QDBusVirtualObject_NewQDBusVirtualObject(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDBusVirtualObject_HandleMessage
func callbackQDBusVirtualObject_HandleMessage(ptr unsafe.Pointer, message unsafe.Pointer, connection unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::handleMessage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QDBusMessage, *QDBusConnection) bool)(NewQDBusMessageFromPointer(message), NewQDBusConnectionFromPointer(connection)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDBusVirtualObject) ConnectHandleMessage(f func(message *QDBusMessage, connection *QDBusConnection) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::handleMessage", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectHandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::handleMessage")
	}
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_HandleMessage(ptr.Pointer(), PointerFromQDBusMessage(message), PointerFromQDBusConnection(connection)) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_Introspect
func callbackQDBusVirtualObject_Introspect(ptr unsafe.Pointer, path C.struct_QtDBus_PackedString) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::introspect"); signal != nil {
		return C.CString(signal.(func(string) string)(cGoUnpackString(path)))
	}

	return C.CString("")
}

func (ptr *QDBusVirtualObject) ConnectIntrospect(f func(path string) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::introspect", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectIntrospect(path string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::introspect")
	}
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		return cGoUnpackString(C.QDBusVirtualObject_Introspect(ptr.Pointer(), pathC))
	}
	return ""
}

//export callbackQDBusVirtualObject_DestroyQDBusVirtualObject
func callbackQDBusVirtualObject_DestroyQDBusVirtualObject(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::~QDBusVirtualObject"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DestroyQDBusVirtualObjectDefault()
	}
}

func (ptr *QDBusVirtualObject) ConnectDestroyQDBusVirtualObject(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::~QDBusVirtualObject", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectDestroyQDBusVirtualObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::~QDBusVirtualObject")
	}
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObjectDefault() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObjectDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusVirtualObject_TimerEvent
func callbackQDBusVirtualObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::timerEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::timerEvent")
	}
}

func (ptr *QDBusVirtualObject) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDBusVirtualObject_ChildEvent
func callbackQDBusVirtualObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::childEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::childEvent")
	}
}

func (ptr *QDBusVirtualObject) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusVirtualObject_ConnectNotify
func callbackQDBusVirtualObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::connectNotify", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::connectNotify")
	}
}

func (ptr *QDBusVirtualObject) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusVirtualObject_CustomEvent
func callbackQDBusVirtualObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::customEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::customEvent")
	}
}

func (ptr *QDBusVirtualObject) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusVirtualObject_DeleteLater
func callbackQDBusVirtualObject_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusVirtualObject) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::deleteLater", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::deleteLater")
	}
}

func (ptr *QDBusVirtualObject) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusVirtualObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDBusVirtualObject_DisconnectNotify
func callbackQDBusVirtualObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::disconnectNotify", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::disconnectNotify")
	}
}

func (ptr *QDBusVirtualObject) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusVirtualObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusVirtualObject_Event
func callbackQDBusVirtualObject_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusVirtualObjectFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusVirtualObject) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::event", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::event")
	}
}

func (ptr *QDBusVirtualObject) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_EventFilter
func callbackQDBusVirtualObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusVirtualObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusVirtualObject) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::eventFilter", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::eventFilter")
	}
}

func (ptr *QDBusVirtualObject) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_MetaObject
func callbackQDBusVirtualObject_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDBusVirtualObject::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusVirtualObjectFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusVirtualObject) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::metaObject", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDBusVirtualObject::metaObject")
	}
}

func (ptr *QDBusVirtualObject) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusVirtualObject_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusVirtualObject) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusVirtualObject_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
