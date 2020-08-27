// +build !minimal

package dbus

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"unsafe"
)

type QDBus struct {
	internal.Internal
}

type QDBus_ITF interface {
	QDBus_PTR() *QDBus
}

func (ptr *QDBus) QDBus_PTR() *QDBus {
	return ptr
}

func (ptr *QDBus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBus(ptr QDBus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBus_PTR().Pointer()
	}
	return nil
}

func (n *QDBus) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusFromPointer(ptr unsafe.Pointer) (n *QDBus) {
	n = new(QDBus)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBus")
	return
}

func (ptr *QDBus) DestroyQDBus() {
}

//go:generate stringer -type=QDBus__CallMode
//QDBus::CallMode
type QDBus__CallMode int64

const (
	QDBus__NoBlock      QDBus__CallMode = QDBus__CallMode(0)
	QDBus__Block        QDBus__CallMode = QDBus__CallMode(1)
	QDBus__BlockWithGui QDBus__CallMode = QDBus__CallMode(2)
	QDBus__AutoDetect   QDBus__CallMode = QDBus__CallMode(3)
)

type QDBusAbstractAdaptor struct {
	core.QObject
}

type QDBusAbstractAdaptor_ITF interface {
	core.QObject_ITF
	QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor
}

func (ptr *QDBusAbstractAdaptor) QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor {
	return ptr
}

func (ptr *QDBusAbstractAdaptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusAbstractAdaptor(ptr QDBusAbstractAdaptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractAdaptor_PTR().Pointer()
	}
	return nil
}

func (n *QDBusAbstractAdaptor) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusAbstractAdaptor) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) (n *QDBusAbstractAdaptor) {
	n = new(QDBusAbstractAdaptor)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusAbstractAdaptor")
	return
}
func NewQDBusAbstractAdaptor(obj core.QObject_ITF) *QDBusAbstractAdaptor {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusAbstractAdaptor", "", obj}).(*QDBusAbstractAdaptor)
}

func (ptr *QDBusAbstractAdaptor) AutoRelaySignals() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoRelaySignals"}).(bool)
}

func (ptr *QDBusAbstractAdaptor) SetAutoRelaySignals(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoRelaySignals", enable})
}

func (ptr *QDBusAbstractAdaptor) ConnectDestroyQDBusAbstractAdaptor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDBusAbstractAdaptor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusAbstractAdaptor) DisconnectDestroyQDBusAbstractAdaptor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDBusAbstractAdaptor"})
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusAbstractAdaptor"})
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusAbstractAdaptorDefault"})
}

func (ptr *QDBusAbstractAdaptor) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusAbstractAdaptor) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusAbstractAdaptor) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractAdaptor) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusAbstractAdaptor) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusAbstractAdaptor) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractAdaptor) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusAbstractAdaptor) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusAbstractAdaptor) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractAdaptor) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusAbstractAdaptor) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusAbstractAdaptor) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractAdaptor) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusAbstractAdaptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusAbstractAdaptor) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusAbstractAdaptor) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusAbstractAdaptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusAbstractAdaptor) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusAbstractAdaptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusAbstractAdaptor) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusAbstractAdaptor) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDBusAbstractInterface struct {
	core.QObject
}

type QDBusAbstractInterface_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterface_PTR() *QDBusAbstractInterface
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusAbstractInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDBusAbstractInterface) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusAbstractInterface) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) (n *QDBusAbstractInterface) {
	n = new(QDBusAbstractInterface)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusAbstractInterface")
	return
}
func (ptr *QDBusAbstractInterface) AsyncCall(method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AsyncCall", method, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}).(*QDBusPendingCall)
}

func (ptr *QDBusAbstractInterface) AsyncCallWithArgumentList(method string, args []*core.QVariant) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AsyncCallWithArgumentList", method, args}).(*QDBusPendingCall)
}

func (ptr *QDBusAbstractInterface) Call(method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Call", method, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}).(*QDBusMessage)
}

func (ptr *QDBusAbstractInterface) Call2(mode QDBus__CallMode, method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Call2", mode, method, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}).(*QDBusMessage)
}

func (ptr *QDBusAbstractInterface) CallWithArgumentList(mode QDBus__CallMode, method string, args []*core.QVariant) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallWithArgumentList", mode, method, args}).(*QDBusMessage)
}

func (ptr *QDBusAbstractInterface) CallWithCallback(method string, args []*core.QVariant, receiver core.QObject_ITF, returnMethod string, errorMethod string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallWithCallback", method, args, receiver, returnMethod, errorMethod}).(bool)
}

func (ptr *QDBusAbstractInterface) CallWithCallback2(method string, args []*core.QVariant, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallWithCallback2", method, args, receiver, slot}).(bool)
}

func (ptr *QDBusAbstractInterface) Connection() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connection"}).(*QDBusConnection)
}

func (ptr *QDBusAbstractInterface) Interface() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Interface"}).(string)
}

func (ptr *QDBusAbstractInterface) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QDBusAbstractInterface) Path() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).(string)
}

func (ptr *QDBusAbstractInterface) Service() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Service"}).(string)
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimeout", timeout})
}

func (ptr *QDBusAbstractInterface) Timeout() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timeout"}).(float64))
}

func (ptr *QDBusAbstractInterface) ConnectDestroyQDBusAbstractInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDBusAbstractInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusAbstractInterface) DisconnectDestroyQDBusAbstractInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDBusAbstractInterface"})
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusAbstractInterface"})
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusAbstractInterfaceDefault"})
}

func (ptr *QDBusAbstractInterface) __asyncCallWithArgumentList_args_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__asyncCallWithArgumentList_args_atList", i}).(*core.QVariant)
}

func (ptr *QDBusAbstractInterface) __asyncCallWithArgumentList_args_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__asyncCallWithArgumentList_args_setList", i})
}

func (ptr *QDBusAbstractInterface) __asyncCallWithArgumentList_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__asyncCallWithArgumentList_args_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __callWithArgumentList_args_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithArgumentList_args_atList", i}).(*core.QVariant)
}

func (ptr *QDBusAbstractInterface) __callWithArgumentList_args_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithArgumentList_args_setList", i})
}

func (ptr *QDBusAbstractInterface) __callWithArgumentList_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithArgumentList_args_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithCallback_args_atList", i}).(*core.QVariant)
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithCallback_args_setList", i})
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithCallback_args_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_atList2(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithCallback_args_atList2", i}).(*core.QVariant)
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_setList2(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithCallback_args_setList2", i})
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithCallback_args_newList2"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __internalConstCall_args_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__internalConstCall_args_atList", i}).(*core.QVariant)
}

func (ptr *QDBusAbstractInterface) __internalConstCall_args_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__internalConstCall_args_setList", i})
}

func (ptr *QDBusAbstractInterface) __internalConstCall_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__internalConstCall_args_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusAbstractInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusAbstractInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusAbstractInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusAbstractInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusAbstractInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusAbstractInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusAbstractInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusAbstractInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusAbstractInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusAbstractInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusAbstractInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusAbstractInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusAbstractInterface) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusAbstractInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusAbstractInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusAbstractInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDBusAbstractInterfaceBase struct {
	core.QObject
}

type QDBusAbstractInterfaceBase_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterfaceBase_PTR() *QDBusAbstractInterfaceBase
}

func (ptr *QDBusAbstractInterfaceBase) QDBusAbstractInterfaceBase_PTR() *QDBusAbstractInterfaceBase {
	return ptr
}

func (ptr *QDBusAbstractInterfaceBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusAbstractInterfaceBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusAbstractInterfaceBase(ptr QDBusAbstractInterfaceBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterfaceBase_PTR().Pointer()
	}
	return nil
}

func (n *QDBusAbstractInterfaceBase) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusAbstractInterfaceBase) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusAbstractInterfaceBaseFromPointer(ptr unsafe.Pointer) (n *QDBusAbstractInterfaceBase) {
	n = new(QDBusAbstractInterfaceBase)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusAbstractInterfaceBase")
	return
}

func (ptr *QDBusAbstractInterfaceBase) DestroyQDBusAbstractInterfaceBase() {
}

func (ptr *QDBusAbstractInterfaceBase) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusAbstractInterfaceBase) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusAbstractInterfaceBase) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterfaceBase) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusAbstractInterfaceBase) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusAbstractInterfaceBase) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterfaceBase) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusAbstractInterfaceBase) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusAbstractInterfaceBase) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterfaceBase) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusAbstractInterfaceBase) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusAbstractInterfaceBase) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusAbstractInterfaceBase) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusAbstractInterfaceBase) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusAbstractInterfaceBase) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusAbstractInterfaceBase) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusAbstractInterfaceBase) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusAbstractInterfaceBase) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusAbstractInterfaceBase) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusAbstractInterfaceBase) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusAbstractInterfaceBase) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDBusArgument struct {
	internal.Internal
}

type QDBusArgument_ITF interface {
	QDBusArgument_PTR() *QDBusArgument
}

func (ptr *QDBusArgument) QDBusArgument_PTR() *QDBusArgument {
	return ptr
}

func (ptr *QDBusArgument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusArgument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusArgument(ptr QDBusArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusArgument_PTR().Pointer()
	}
	return nil
}

func (n *QDBusArgument) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusArgumentFromPointer(ptr unsafe.Pointer) (n *QDBusArgument) {
	n = new(QDBusArgument)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusArgument")
	return
}

//go:generate stringer -type=QDBusArgument__ElementType
//QDBusArgument::ElementType
type QDBusArgument__ElementType int64

const (
	QDBusArgument__BasicType     QDBusArgument__ElementType = QDBusArgument__ElementType(0)
	QDBusArgument__VariantType   QDBusArgument__ElementType = QDBusArgument__ElementType(1)
	QDBusArgument__ArrayType     QDBusArgument__ElementType = QDBusArgument__ElementType(2)
	QDBusArgument__StructureType QDBusArgument__ElementType = QDBusArgument__ElementType(3)
	QDBusArgument__MapType       QDBusArgument__ElementType = QDBusArgument__ElementType(4)
	QDBusArgument__MapEntryType  QDBusArgument__ElementType = QDBusArgument__ElementType(5)
	QDBusArgument__UnknownType   QDBusArgument__ElementType = QDBusArgument__ElementType(-1)
)

func NewQDBusArgument() *QDBusArgument {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusArgument", ""}).(*QDBusArgument)
}

func NewQDBusArgument2(other QDBusArgument_ITF) *QDBusArgument {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusArgument2", "", other}).(*QDBusArgument)
}

func (ptr *QDBusArgument) AsVariant() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AsVariant"}).(*core.QVariant)
}

func (ptr *QDBusArgument) AtEnd() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtEnd"}).(bool)
}

func (ptr *QDBusArgument) BeginArray(id int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginArray", id})
}

func (ptr *QDBusArgument) BeginArray2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginArray2"})
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginMap", kid, vid})
}

func (ptr *QDBusArgument) BeginMap2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginMap2"})
}

func (ptr *QDBusArgument) BeginMapEntry() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginMapEntry"})
}

func (ptr *QDBusArgument) BeginMapEntry2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginMapEntry2"})
}

func (ptr *QDBusArgument) BeginStructure() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginStructure"})
}

func (ptr *QDBusArgument) BeginStructure2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginStructure2"})
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {

	return QDBusArgument__ElementType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentType"}).(float64))
}

func (ptr *QDBusArgument) EndArray() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndArray"})
}

func (ptr *QDBusArgument) EndArray2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndArray2"})
}

func (ptr *QDBusArgument) EndMap() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndMap"})
}

func (ptr *QDBusArgument) EndMap2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndMap2"})
}

func (ptr *QDBusArgument) EndMapEntry() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndMapEntry"})
}

func (ptr *QDBusArgument) EndMapEntry2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndMapEntry2"})
}

func (ptr *QDBusArgument) EndStructure() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndStructure"})
}

func (ptr *QDBusArgument) EndStructure2() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndStructure2"})
}

func (ptr *QDBusArgument) Swap(other QDBusArgument_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusArgument"})
}

type QDBusConnection struct {
	internal.Internal
}

type QDBusConnection_ITF interface {
	QDBusConnection_PTR() *QDBusConnection
}

func (ptr *QDBusConnection) QDBusConnection_PTR() *QDBusConnection {
	return ptr
}

func (ptr *QDBusConnection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusConnection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusConnection(ptr QDBusConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnection_PTR().Pointer()
	}
	return nil
}

func (n *QDBusConnection) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusConnectionFromPointer(ptr unsafe.Pointer) (n *QDBusConnection) {
	n = new(QDBusConnection)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusConnection")
	return
}

//go:generate stringer -type=QDBusConnection__BusType
//QDBusConnection::BusType
type QDBusConnection__BusType int64

const (
	QDBusConnection__SessionBus    QDBusConnection__BusType = QDBusConnection__BusType(0)
	QDBusConnection__SystemBus     QDBusConnection__BusType = QDBusConnection__BusType(1)
	QDBusConnection__ActivationBus QDBusConnection__BusType = QDBusConnection__BusType(2)
)

//go:generate stringer -type=QDBusConnection__RegisterOption
//QDBusConnection::RegisterOption
type QDBusConnection__RegisterOption int64

const (
	QDBusConnection__ExportAdaptors                QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x01)
	QDBusConnection__ExportScriptableSlots         QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x10)
	QDBusConnection__ExportScriptableSignals       QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x20)
	QDBusConnection__ExportScriptableProperties    QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x40)
	QDBusConnection__ExportScriptableInvokables    QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x80)
	QDBusConnection__ExportScriptableContents      QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0xf0)
	QDBusConnection__ExportNonScriptableSlots      QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x100)
	QDBusConnection__ExportNonScriptableSignals    QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x200)
	QDBusConnection__ExportNonScriptableProperties QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x400)
	QDBusConnection__ExportNonScriptableInvokables QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x800)
	QDBusConnection__ExportNonScriptableContents   QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0xf00)
	QDBusConnection__ExportChildObjects            QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x1000)
)

//go:generate stringer -type=QDBusConnection__UnregisterMode
//QDBusConnection::UnregisterMode
type QDBusConnection__UnregisterMode int64

const (
	QDBusConnection__UnregisterNode QDBusConnection__UnregisterMode = QDBusConnection__UnregisterMode(0)
	QDBusConnection__UnregisterTree QDBusConnection__UnregisterMode = QDBusConnection__UnregisterMode(1)
)

//go:generate stringer -type=QDBusConnection__ConnectionCapability
//QDBusConnection::ConnectionCapability
type QDBusConnection__ConnectionCapability int64

const (
	QDBusConnection__UnixFileDescriptorPassing QDBusConnection__ConnectionCapability = QDBusConnection__ConnectionCapability(0x0001)
)

func NewQDBusConnection(name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusConnection", "", name}).(*QDBusConnection)
}

func NewQDBusConnection2(other QDBusConnection_ITF) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusConnection2", "", other}).(*QDBusConnection)
}

func (ptr *QDBusConnection) AsyncCall(message QDBusMessage_ITF, timeout int) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AsyncCall", message, timeout}).(*QDBusPendingCall)
}

func (ptr *QDBusConnection) BaseService() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseService"}).(string)
}

func (ptr *QDBusConnection) Call(message QDBusMessage_ITF, mode QDBus__CallMode, timeout int) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Call", message, mode, timeout}).(*QDBusMessage)
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessage_ITF, receiver core.QObject_ITF, returnMethod string, errorMethod string, timeout int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallWithCallback", message, receiver, returnMethod, errorMethod, timeout}).(bool)
}

func (ptr *QDBusConnection) Connect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connect", service, path, interfa, name, receiver, slot}).(bool)
}

func (ptr *QDBusConnection) Connect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connect2", service, path, interfa, name, signature, receiver, slot}).(bool)
}

func (ptr *QDBusConnection) Connect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connect3", service, path, interfa, name, argumentMatch, signature, receiver, slot}).(bool)
}

func QDBusConnection_ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_ConnectToBus", "", ty, name}).(*QDBusConnection)
}

func (ptr *QDBusConnection) ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_ConnectToBus", "", ty, name}).(*QDBusConnection)
}

func QDBusConnection_ConnectToBus2(address string, name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_ConnectToBus2", "", address, name}).(*QDBusConnection)
}

func (ptr *QDBusConnection) ConnectToBus2(address string, name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_ConnectToBus2", "", address, name}).(*QDBusConnection)
}

func QDBusConnection_ConnectToPeer(address string, name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_ConnectToPeer", "", address, name}).(*QDBusConnection)
}

func (ptr *QDBusConnection) ConnectToPeer(address string, name string) *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_ConnectToPeer", "", address, name}).(*QDBusConnection)
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {

	return QDBusConnection__ConnectionCapability(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectionCapabilities"}).(float64))
}

func (ptr *QDBusConnection) Disconnect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnect", service, path, interfa, name, receiver, slot}).(bool)
}

func (ptr *QDBusConnection) Disconnect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnect2", service, path, interfa, name, signature, receiver, slot}).(bool)
}

func (ptr *QDBusConnection) Disconnect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnect3", service, path, interfa, name, argumentMatch, signature, receiver, slot}).(bool)
}

func QDBusConnection_DisconnectFromBus(name string) {

	internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_DisconnectFromBus", "", name})
}

func (ptr *QDBusConnection) DisconnectFromBus(name string) {

	internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_DisconnectFromBus", "", name})
}

func QDBusConnection_DisconnectFromPeer(name string) {

	internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_DisconnectFromPeer", "", name})
}

func (ptr *QDBusConnection) DisconnectFromPeer(name string) {

	internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_DisconnectFromPeer", "", name})
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Interface"}).(*QDBusConnectionInterface)
}

func (ptr *QDBusConnection) IsConnected() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsConnected"}).(bool)
}

func QDBusConnection_LocalMachineId() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_LocalMachineId", ""}).(*core.QByteArray)
}

func (ptr *QDBusConnection) LocalMachineId() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_LocalMachineId", ""}).(*core.QByteArray)
}

func (ptr *QDBusConnection) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ObjectRegisteredAt", path}).(*core.QObject)
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterObject", path, object, options}).(bool)
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterObject2", path, interfa, object, options}).(bool)
}

func (ptr *QDBusConnection) RegisterService(serviceName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterService", serviceName}).(bool)
}

func (ptr *QDBusConnection) Send(message QDBusMessage_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Send", message}).(bool)
}

func QDBusConnection_SessionBus() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_SessionBus", ""}).(*QDBusConnection)
}

func (ptr *QDBusConnection) SessionBus() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_SessionBus", ""}).(*QDBusConnection)
}

func (ptr *QDBusConnection) Swap(other QDBusConnection_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func QDBusConnection_SystemBus() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_SystemBus", ""}).(*QDBusConnection)
}

func (ptr *QDBusConnection) SystemBus() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusConnection_SystemBus", ""}).(*QDBusConnection)
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnregisterObject", path, mode})
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnregisterService", serviceName}).(bool)
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusConnection"})
}

type QDBusConnectionInterface struct {
	QDBusAbstractInterface
}

type QDBusConnectionInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusConnectionInterface_PTR() *QDBusConnectionInterface
}

func (ptr *QDBusConnectionInterface) QDBusConnectionInterface_PTR() *QDBusConnectionInterface {
	return ptr
}

func (ptr *QDBusConnectionInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusConnectionInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusAbstractInterface_PTR().SetPointer(p)
	}
}

func PointerFromQDBusConnectionInterface(ptr QDBusConnectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDBusConnectionInterface) InitFromInternal(ptr uintptr, name string) {
	n.QDBusAbstractInterface_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusConnectionInterface) ClassNameInternalF() string {
	return n.QDBusAbstractInterface_PTR().ClassNameInternalF()
}

func NewQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) (n *QDBusConnectionInterface) {
	n = new(QDBusConnectionInterface)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusConnectionInterface")
	return
}

func (ptr *QDBusConnectionInterface) DestroyQDBusConnectionInterface() {
}

//go:generate stringer -type=QDBusConnectionInterface__ServiceQueueOptions
//QDBusConnectionInterface::ServiceQueueOptions
type QDBusConnectionInterface__ServiceQueueOptions int64

const (
	QDBusConnectionInterface__DontQueueService       QDBusConnectionInterface__ServiceQueueOptions = QDBusConnectionInterface__ServiceQueueOptions(0)
	QDBusConnectionInterface__QueueService           QDBusConnectionInterface__ServiceQueueOptions = QDBusConnectionInterface__ServiceQueueOptions(1)
	QDBusConnectionInterface__ReplaceExistingService QDBusConnectionInterface__ServiceQueueOptions = QDBusConnectionInterface__ServiceQueueOptions(2)
)

//go:generate stringer -type=QDBusConnectionInterface__ServiceReplacementOptions
//QDBusConnectionInterface::ServiceReplacementOptions
type QDBusConnectionInterface__ServiceReplacementOptions int64

const (
	QDBusConnectionInterface__DontAllowReplacement QDBusConnectionInterface__ServiceReplacementOptions = QDBusConnectionInterface__ServiceReplacementOptions(0)
	QDBusConnectionInterface__AllowReplacement     QDBusConnectionInterface__ServiceReplacementOptions = QDBusConnectionInterface__ServiceReplacementOptions(1)
)

//go:generate stringer -type=QDBusConnectionInterface__RegisterServiceReply
//QDBusConnectionInterface::RegisterServiceReply
type QDBusConnectionInterface__RegisterServiceReply int64

const (
	QDBusConnectionInterface__ServiceNotRegistered QDBusConnectionInterface__RegisterServiceReply = QDBusConnectionInterface__RegisterServiceReply(0)
	QDBusConnectionInterface__ServiceRegistered    QDBusConnectionInterface__RegisterServiceReply = QDBusConnectionInterface__RegisterServiceReply(1)
	QDBusConnectionInterface__ServiceQueued        QDBusConnectionInterface__RegisterServiceReply = QDBusConnectionInterface__RegisterServiceReply(2)
)

func (ptr *QDBusConnectionInterface) ConnectCallWithCallbackFailed(f func(error *QDBusError, call *QDBusMessage)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCallWithCallbackFailed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusConnectionInterface) DisconnectCallWithCallbackFailed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCallWithCallbackFailed"})
}

func (ptr *QDBusConnectionInterface) CallWithCallbackFailed(error QDBusError_ITF, call QDBusMessage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallWithCallbackFailed", error, call})
}

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(service string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceRegistered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceRegistered"})
}

func (ptr *QDBusConnectionInterface) ServiceRegistered(service string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceRegistered", service})
}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(service string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceUnregistered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceUnregistered"})
}

func (ptr *QDBusConnectionInterface) ServiceUnregistered(service string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceUnregistered", service})
}

type QDBusContext struct {
	internal.Internal
}

type QDBusContext_ITF interface {
	QDBusContext_PTR() *QDBusContext
}

func (ptr *QDBusContext) QDBusContext_PTR() *QDBusContext {
	return ptr
}

func (ptr *QDBusContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusContext(ptr QDBusContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusContext_PTR().Pointer()
	}
	return nil
}

func (n *QDBusContext) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusContextFromPointer(ptr unsafe.Pointer) (n *QDBusContext) {
	n = new(QDBusContext)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusContext")
	return
}
func NewQDBusContext() *QDBusContext {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusContext", ""}).(*QDBusContext)
}

func (ptr *QDBusContext) CalledFromDBus() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CalledFromDBus"}).(bool)
}

func (ptr *QDBusContext) Connection() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connection"}).(*QDBusConnection)
}

func (ptr *QDBusContext) IsDelayedReply() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDelayedReply"}).(bool)
}

func (ptr *QDBusContext) Message() *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Message"}).(*QDBusMessage)
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendErrorReply", name, msg})
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendErrorReply2", ty, msg})
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDelayedReply", enable})
}

func (ptr *QDBusContext) DestroyQDBusContext() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusContext"})
}

type QDBusError struct {
	internal.Internal
}

type QDBusError_ITF interface {
	QDBusError_PTR() *QDBusError
}

func (ptr *QDBusError) QDBusError_PTR() *QDBusError {
	return ptr
}

func (ptr *QDBusError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusError(ptr QDBusError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusError_PTR().Pointer()
	}
	return nil
}

func (n *QDBusError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusErrorFromPointer(ptr unsafe.Pointer) (n *QDBusError) {
	n = new(QDBusError)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusError")
	return
}

func (ptr *QDBusError) DestroyQDBusError() {
}

//go:generate stringer -type=QDBusError__ErrorType
//QDBusError::ErrorType
type QDBusError__ErrorType int64

const (
	QDBusError__NoError           QDBusError__ErrorType = QDBusError__ErrorType(0)
	QDBusError__Other             QDBusError__ErrorType = QDBusError__ErrorType(1)
	QDBusError__Failed            QDBusError__ErrorType = QDBusError__ErrorType(2)
	QDBusError__NoMemory          QDBusError__ErrorType = QDBusError__ErrorType(3)
	QDBusError__ServiceUnknown    QDBusError__ErrorType = QDBusError__ErrorType(4)
	QDBusError__NoReply           QDBusError__ErrorType = QDBusError__ErrorType(5)
	QDBusError__BadAddress        QDBusError__ErrorType = QDBusError__ErrorType(6)
	QDBusError__NotSupported      QDBusError__ErrorType = QDBusError__ErrorType(7)
	QDBusError__LimitsExceeded    QDBusError__ErrorType = QDBusError__ErrorType(8)
	QDBusError__AccessDenied      QDBusError__ErrorType = QDBusError__ErrorType(9)
	QDBusError__NoServer          QDBusError__ErrorType = QDBusError__ErrorType(10)
	QDBusError__Timeout           QDBusError__ErrorType = QDBusError__ErrorType(11)
	QDBusError__NoNetwork         QDBusError__ErrorType = QDBusError__ErrorType(12)
	QDBusError__AddressInUse      QDBusError__ErrorType = QDBusError__ErrorType(13)
	QDBusError__Disconnected      QDBusError__ErrorType = QDBusError__ErrorType(14)
	QDBusError__InvalidArgs       QDBusError__ErrorType = QDBusError__ErrorType(15)
	QDBusError__UnknownMethod     QDBusError__ErrorType = QDBusError__ErrorType(16)
	QDBusError__TimedOut          QDBusError__ErrorType = QDBusError__ErrorType(17)
	QDBusError__InvalidSignature  QDBusError__ErrorType = QDBusError__ErrorType(18)
	QDBusError__UnknownInterface  QDBusError__ErrorType = QDBusError__ErrorType(19)
	QDBusError__UnknownObject     QDBusError__ErrorType = QDBusError__ErrorType(20)
	QDBusError__UnknownProperty   QDBusError__ErrorType = QDBusError__ErrorType(21)
	QDBusError__PropertyReadOnly  QDBusError__ErrorType = QDBusError__ErrorType(22)
	QDBusError__InternalError     QDBusError__ErrorType = QDBusError__ErrorType(23)
	QDBusError__InvalidService    QDBusError__ErrorType = QDBusError__ErrorType(24)
	QDBusError__InvalidObjectPath QDBusError__ErrorType = QDBusError__ErrorType(25)
	QDBusError__InvalidInterface  QDBusError__ErrorType = QDBusError__ErrorType(26)
	QDBusError__InvalidMember     QDBusError__ErrorType = QDBusError__ErrorType(27)
)

func QDBusError_ErrorString(error QDBusError__ErrorType) string {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusError_ErrorString", "", error}).(string)
}

func (ptr *QDBusError) ErrorString(error QDBusError__ErrorType) string {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusError_ErrorString", "", error}).(string)
}

func (ptr *QDBusError) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QDBusError) Message() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Message"}).(string)
}

func (ptr *QDBusError) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDBusError) Swap(other QDBusError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDBusError) Type() QDBusError__ErrorType {

	return QDBusError__ErrorType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusInterface_PTR() *QDBusInterface
}

func (ptr *QDBusInterface) QDBusInterface_PTR() *QDBusInterface {
	return ptr
}

func (ptr *QDBusInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusAbstractInterface_PTR().SetPointer(p)
	}
}

func PointerFromQDBusInterface(ptr QDBusInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDBusInterface) InitFromInternal(ptr uintptr, name string) {
	n.QDBusAbstractInterface_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusInterface) ClassNameInternalF() string {
	return n.QDBusAbstractInterface_PTR().ClassNameInternalF()
}

func NewQDBusInterfaceFromPointer(ptr unsafe.Pointer) (n *QDBusInterface) {
	n = new(QDBusInterface)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusInterface")
	return
}
func NewQDBusInterface2(service string, path string, interfa string, connection QDBusConnection_ITF, parent core.QObject_ITF) *QDBusInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusInterface2", "", service, path, interfa, connection, parent}).(*QDBusInterface)
}

func (ptr *QDBusInterface) ConnectDestroyQDBusInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDBusInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusInterface) DisconnectDestroyQDBusInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDBusInterface"})
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusInterface"})
}

func (ptr *QDBusInterface) DestroyQDBusInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusInterfaceDefault"})
}

type QDBusMessage struct {
	internal.Internal
}

type QDBusMessage_ITF interface {
	QDBusMessage_PTR() *QDBusMessage
}

func (ptr *QDBusMessage) QDBusMessage_PTR() *QDBusMessage {
	return ptr
}

func (ptr *QDBusMessage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusMessage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusMessage(ptr QDBusMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusMessage_PTR().Pointer()
	}
	return nil
}

func (n *QDBusMessage) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusMessageFromPointer(ptr unsafe.Pointer) (n *QDBusMessage) {
	n = new(QDBusMessage)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusMessage")
	return
}

//go:generate stringer -type=QDBusMessage__MessageType
//QDBusMessage::MessageType
type QDBusMessage__MessageType int64

const (
	QDBusMessage__InvalidMessage    QDBusMessage__MessageType = QDBusMessage__MessageType(0)
	QDBusMessage__MethodCallMessage QDBusMessage__MessageType = QDBusMessage__MessageType(1)
	QDBusMessage__ReplyMessage      QDBusMessage__MessageType = QDBusMessage__MessageType(2)
	QDBusMessage__ErrorMessage      QDBusMessage__MessageType = QDBusMessage__MessageType(3)
	QDBusMessage__SignalMessage     QDBusMessage__MessageType = QDBusMessage__MessageType(4)
)

func NewQDBusMessage() *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusMessage", ""}).(*QDBusMessage)
}

func NewQDBusMessage2(other QDBusMessage_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusMessage2", "", other}).(*QDBusMessage)
}

func (ptr *QDBusMessage) Arguments() []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Arguments"}).([]*core.QVariant)
}

func (ptr *QDBusMessage) AutoStartService() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoStartService"}).(bool)
}

func QDBusMessage_CreateError(name string, msg string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateError", "", name, msg}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateError(name string, msg string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateError", "", name, msg}).(*QDBusMessage)
}

func QDBusMessage_CreateError2(error QDBusError_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateError2", "", error}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateError2(error QDBusError_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateError2", "", error}).(*QDBusMessage)
}

func QDBusMessage_CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateError3", "", ty, msg}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateError3", "", ty, msg}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateErrorReply(name string, msg string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateErrorReply", name, msg}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateErrorReply2(error QDBusError_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateErrorReply2", error}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateErrorReply3(ty QDBusError__ErrorType, msg string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateErrorReply3", ty, msg}).(*QDBusMessage)
}

func QDBusMessage_CreateMethodCall(service string, path string, interfa string, method string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateMethodCall", "", service, path, interfa, method}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateMethodCall(service string, path string, interfa string, method string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateMethodCall", "", service, path, interfa, method}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateReply(arguments []*core.QVariant) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateReply", arguments}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateReply2(argument core.QVariant_ITF) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateReply2", argument}).(*QDBusMessage)
}

func QDBusMessage_CreateSignal(path string, interfa string, name string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateSignal", "", path, interfa, name}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateSignal(path string, interfa string, name string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateSignal", "", path, interfa, name}).(*QDBusMessage)
}

func QDBusMessage_CreateTargetedSignal(service string, path string, interfa string, name string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateTargetedSignal", "", service, path, interfa, name}).(*QDBusMessage)
}

func (ptr *QDBusMessage) CreateTargetedSignal(service string, path string, interfa string, name string) *QDBusMessage {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusMessage_CreateTargetedSignal", "", service, path, interfa, name}).(*QDBusMessage)
}

func (ptr *QDBusMessage) ErrorMessage() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorMessage"}).(string)
}

func (ptr *QDBusMessage) ErrorName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorName"}).(string)
}

func (ptr *QDBusMessage) Interface() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Interface"}).(string)
}

func (ptr *QDBusMessage) IsDelayedReply() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDelayedReply"}).(bool)
}

func (ptr *QDBusMessage) IsInteractiveAuthorizationAllowed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInteractiveAuthorizationAllowed"}).(bool)
}

func (ptr *QDBusMessage) IsReplyRequired() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReplyRequired"}).(bool)
}

func (ptr *QDBusMessage) Member() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Member"}).(string)
}

func (ptr *QDBusMessage) Path() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).(string)
}

func (ptr *QDBusMessage) Service() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Service"}).(string)
}

func (ptr *QDBusMessage) SetArguments(arguments []*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetArguments", arguments})
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoStartService", enable})
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDelayedReply", enable})
}

func (ptr *QDBusMessage) SetInteractiveAuthorizationAllowed(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInteractiveAuthorizationAllowed", enable})
}

func (ptr *QDBusMessage) Signature() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Signature"}).(string)
}

func (ptr *QDBusMessage) Swap(other QDBusMessage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {

	return QDBusMessage__MessageType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusMessage"})
}

func (ptr *QDBusMessage) __arguments_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__arguments_atList", i}).(*core.QVariant)
}

func (ptr *QDBusMessage) __arguments_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__arguments_setList", i})
}

func (ptr *QDBusMessage) __arguments_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__arguments_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusMessage) __createReply_arguments_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createReply_arguments_atList", i}).(*core.QVariant)
}

func (ptr *QDBusMessage) __createReply_arguments_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createReply_arguments_setList", i})
}

func (ptr *QDBusMessage) __createReply_arguments_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createReply_arguments_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusMessage) __setArguments_arguments_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setArguments_arguments_atList", i}).(*core.QVariant)
}

func (ptr *QDBusMessage) __setArguments_arguments_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setArguments_arguments_setList", i})
}

func (ptr *QDBusMessage) __setArguments_arguments_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setArguments_arguments_newList"}).(unsafe.Pointer)
}

type QDBusObjectPath struct {
	internal.Internal
}

type QDBusObjectPath_ITF interface {
	QDBusObjectPath_PTR() *QDBusObjectPath
}

func (ptr *QDBusObjectPath) QDBusObjectPath_PTR() *QDBusObjectPath {
	return ptr
}

func (ptr *QDBusObjectPath) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusObjectPath) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusObjectPath(ptr QDBusObjectPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusObjectPath_PTR().Pointer()
	}
	return nil
}

func (n *QDBusObjectPath) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusObjectPathFromPointer(ptr unsafe.Pointer) (n *QDBusObjectPath) {
	n = new(QDBusObjectPath)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusObjectPath")
	return
}

func (ptr *QDBusObjectPath) DestroyQDBusObjectPath() {
}

func NewQDBusObjectPath() *QDBusObjectPath {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusObjectPath", ""}).(*QDBusObjectPath)
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusObjectPath2", "", path}).(*QDBusObjectPath)
}

func NewQDBusObjectPath3(path core.QLatin1String_ITF) *QDBusObjectPath {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusObjectPath3", "", path}).(*QDBusObjectPath)
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusObjectPath4", "", path}).(*QDBusObjectPath)
}

func (ptr *QDBusObjectPath) Path() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).(string)
}

func (ptr *QDBusObjectPath) SetPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPath", path})
}

func (ptr *QDBusObjectPath) Swap(other QDBusObjectPath_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

type QDBusPendingCall struct {
	internal.Internal
}

type QDBusPendingCall_ITF interface {
	QDBusPendingCall_PTR() *QDBusPendingCall
}

func (ptr *QDBusPendingCall) QDBusPendingCall_PTR() *QDBusPendingCall {
	return ptr
}

func (ptr *QDBusPendingCall) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusPendingCall) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusPendingCall(ptr QDBusPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func (n *QDBusPendingCall) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusPendingCallFromPointer(ptr unsafe.Pointer) (n *QDBusPendingCall) {
	n = new(QDBusPendingCall)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusPendingCall")
	return
}
func NewQDBusPendingCall(other QDBusPendingCall_ITF) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusPendingCall", "", other}).(*QDBusPendingCall)
}

func QDBusPendingCall_FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusPendingCall_FromCompletedCall", "", msg}).(*QDBusPendingCall)
}

func (ptr *QDBusPendingCall) FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusPendingCall_FromCompletedCall", "", msg}).(*QDBusPendingCall)
}

func QDBusPendingCall_FromError(error QDBusError_ITF) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusPendingCall_FromError", "", error}).(*QDBusPendingCall)
}

func (ptr *QDBusPendingCall) FromError(error QDBusError_ITF) *QDBusPendingCall {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusPendingCall_FromError", "", error}).(*QDBusPendingCall)
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCall_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusPendingCall"})
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

func (ptr *QDBusPendingCallWatcher) QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher {
	return ptr
}

func (ptr *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QDBusPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func (n *QDBusPendingCallWatcher) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QDBusPendingCall_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusPendingCallWatcher) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) (n *QDBusPendingCallWatcher) {
	n = new(QDBusPendingCallWatcher)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusPendingCallWatcher")
	return
}
func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusPendingCallWatcher", "", call, parent}).(*QDBusPendingCallWatcher)
}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(se *QDBusPendingCallWatcher)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QDBusPendingCallWatcher) Finished(se QDBusPendingCallWatcher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Finished", se})
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "IsFinished"}).(bool)
}

func (ptr *QDBusPendingCallWatcher) WaitForFinished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "WaitForFinished"})
}

func (ptr *QDBusPendingCallWatcher) ConnectDestroyQDBusPendingCallWatcher(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDestroyQDBusPendingCallWatcher", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusPendingCallWatcher) DisconnectDestroyQDBusPendingCallWatcher() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDestroyQDBusPendingCallWatcher"})
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQDBusPendingCallWatcher"})
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcherDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQDBusPendingCallWatcherDefault"})
}

func (ptr *QDBusPendingCallWatcher) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusPendingCallWatcher) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusPendingCallWatcher) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusPendingCallWatcher) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusPendingCallWatcher) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusPendingCallWatcher) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusPendingCallWatcher) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusPendingCallWatcher) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusPendingCallWatcher) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusPendingCallWatcher) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusPendingCallWatcher) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusPendingCallWatcher) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEvent", event})
}

func (ptr *QDBusPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusPendingCallWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotify", sign})
}

func (ptr *QDBusPendingCallWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEvent", event})
}

func (ptr *QDBusPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusPendingCallWatcher) DeleteLater() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLater"})
}

func (ptr *QDBusPendingCallWatcher) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotify", sign})
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusPendingCallWatcher) Event(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Event", e}).(bool)
}

func (ptr *QDBusPendingCallWatcher) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusPendingCallWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilter", watched, event}).(bool)
}

func (ptr *QDBusPendingCallWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusPendingCallWatcher) MetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObject"}).(*core.QMetaObject)
}

func (ptr *QDBusPendingCallWatcher) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEvent", event})
}

func (ptr *QDBusPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEventDefault", event})
}

type QDBusPendingReply struct {
	QDBusPendingCall
}

type QDBusPendingReply_ITF interface {
	QDBusPendingCall_ITF
	QDBusPendingReply_PTR() *QDBusPendingReply
}

func (ptr *QDBusPendingReply) QDBusPendingReply_PTR() *QDBusPendingReply {
	return ptr
}

func (ptr *QDBusPendingReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusPendingReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQDBusPendingReply(ptr QDBusPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReply_PTR().Pointer()
	}
	return nil
}

func (n *QDBusPendingReply) InitFromInternal(ptr uintptr, name string) {
	n.QDBusPendingCall_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusPendingReply) ClassNameInternalF() string {
	return n.QDBusPendingCall_PTR().ClassNameInternalF()
}

func NewQDBusPendingReplyFromPointer(ptr unsafe.Pointer) (n *QDBusPendingReply) {
	n = new(QDBusPendingReply)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusPendingReply")
	return
}

func (ptr *QDBusPendingReply) DestroyQDBusPendingReply() {
}

type QDBusPendingReplyData struct {
	QDBusPendingCall
}

type QDBusPendingReplyData_ITF interface {
	QDBusPendingCall_ITF
	QDBusPendingReplyData_PTR() *QDBusPendingReplyData
}

func (ptr *QDBusPendingReplyData) QDBusPendingReplyData_PTR() *QDBusPendingReplyData {
	return ptr
}

func (ptr *QDBusPendingReplyData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusPendingReplyData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQDBusPendingReplyData(ptr QDBusPendingReplyData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReplyData_PTR().Pointer()
	}
	return nil
}

func (n *QDBusPendingReplyData) InitFromInternal(ptr uintptr, name string) {
	n.QDBusPendingCall_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusPendingReplyData) ClassNameInternalF() string {
	return n.QDBusPendingCall_PTR().ClassNameInternalF()
}

func NewQDBusPendingReplyDataFromPointer(ptr unsafe.Pointer) (n *QDBusPendingReplyData) {
	n = new(QDBusPendingReplyData)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusPendingReplyData")
	return
}

func (ptr *QDBusPendingReplyData) DestroyQDBusPendingReplyData() {
}

type QDBusReply struct {
	internal.Internal
}

type QDBusReply_ITF interface {
	QDBusReply_PTR() *QDBusReply
}

func (ptr *QDBusReply) QDBusReply_PTR() *QDBusReply {
	return ptr
}

func (ptr *QDBusReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusReply(ptr QDBusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusReply_PTR().Pointer()
	}
	return nil
}

func (n *QDBusReply) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusReplyFromPointer(ptr unsafe.Pointer) (n *QDBusReply) {
	n = new(QDBusReply)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusReply")
	return
}

func (ptr *QDBusReply) DestroyQDBusReply() {
}

type QDBusServer struct {
	core.QObject
}

type QDBusServer_ITF interface {
	core.QObject_ITF
	QDBusServer_PTR() *QDBusServer
}

func (ptr *QDBusServer) QDBusServer_PTR() *QDBusServer {
	return ptr
}

func (ptr *QDBusServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusServer(ptr QDBusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServer_PTR().Pointer()
	}
	return nil
}

func (n *QDBusServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusServerFromPointer(ptr unsafe.Pointer) (n *QDBusServer) {
	n = new(QDBusServer)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusServer")
	return
}
func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusServer", "", address, parent}).(*QDBusServer)
}

func NewQDBusServer2(parent core.QObject_ITF) *QDBusServer {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusServer2", "", parent}).(*QDBusServer)
}

func (ptr *QDBusServer) Address() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Address"}).(string)
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAnonymousAuthenticationAllowed"}).(bool)
}

func (ptr *QDBusServer) IsConnected() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsConnected"}).(bool)
}

func (ptr *QDBusServer) ConnectNewConnection(f func(connection *QDBusConnection)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusServer) DisconnectNewConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewConnection"})
}

func (ptr *QDBusServer) NewConnection(connection QDBusConnection_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewConnection", connection})
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnonymousAuthenticationAllowed", value})
}

func (ptr *QDBusServer) ConnectDestroyQDBusServer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDBusServer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusServer) DisconnectDestroyQDBusServer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDBusServer"})
}

func (ptr *QDBusServer) DestroyQDBusServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusServer"})
}

func (ptr *QDBusServer) DestroyQDBusServerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusServerDefault"})
}

func (ptr *QDBusServer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusServer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusServer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusServer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusServer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusServer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusServer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusServer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusServer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusServer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusServer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusServer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusServer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusServer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusServer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusServer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDBusServiceWatcher struct {
	core.QObject
}

type QDBusServiceWatcher_ITF interface {
	core.QObject_ITF
	QDBusServiceWatcher_PTR() *QDBusServiceWatcher
}

func (ptr *QDBusServiceWatcher) QDBusServiceWatcher_PTR() *QDBusServiceWatcher {
	return ptr
}

func (ptr *QDBusServiceWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusServiceWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcher_PTR().Pointer()
	}
	return nil
}

func (n *QDBusServiceWatcher) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusServiceWatcher) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) (n *QDBusServiceWatcher) {
	n = new(QDBusServiceWatcher)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusServiceWatcher")
	return
}

//go:generate stringer -type=QDBusServiceWatcher__WatchModeFlag
//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int64

const (
	QDBusServiceWatcher__WatchForRegistration   QDBusServiceWatcher__WatchModeFlag = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration QDBusServiceWatcher__WatchModeFlag = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    QDBusServiceWatcher__WatchModeFlag = QDBusServiceWatcher__WatchModeFlag(0x03)
)

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusServiceWatcher", "", parent}).(*QDBusServiceWatcher)
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusServiceWatcher2", "", service, connection, watchMode, parent}).(*QDBusServiceWatcher)
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddWatchedService", newService})
}

func (ptr *QDBusServiceWatcher) Connection() *QDBusConnection {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connection"}).(*QDBusConnection)
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveWatchedService", service}).(bool)
}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceOwnerChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceOwnerChanged"})
}

func (ptr *QDBusServiceWatcher) ServiceOwnerChanged(serviceName string, oldOwner string, newOwner string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceOwnerChanged", serviceName, oldOwner, newOwner})
}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceRegistered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceRegistered"})
}

func (ptr *QDBusServiceWatcher) ServiceRegistered(serviceName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceRegistered", serviceName})
}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceUnregistered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceUnregistered"})
}

func (ptr *QDBusServiceWatcher) ServiceUnregistered(serviceName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceUnregistered", serviceName})
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConnection", connection})
}

func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWatchMode", mode})
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWatchedServices", services})
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {

	return QDBusServiceWatcher__WatchModeFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WatchMode"}).(float64))
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WatchedServices"}).([]string)
}

func (ptr *QDBusServiceWatcher) ConnectDestroyQDBusServiceWatcher(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDBusServiceWatcher", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusServiceWatcher) DisconnectDestroyQDBusServiceWatcher() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDBusServiceWatcher"})
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusServiceWatcher"})
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcherDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusServiceWatcherDefault"})
}

func (ptr *QDBusServiceWatcher) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusServiceWatcher) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusServiceWatcher) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusServiceWatcher) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusServiceWatcher) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusServiceWatcher) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusServiceWatcher) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusServiceWatcher) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusServiceWatcher) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusServiceWatcher) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusServiceWatcher) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusServiceWatcher) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusServiceWatcher) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusServiceWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusServiceWatcher) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusServiceWatcher) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusServiceWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusServiceWatcher) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusServiceWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusServiceWatcher) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusServiceWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDBusSignature struct {
	internal.Internal
}

type QDBusSignature_ITF interface {
	QDBusSignature_PTR() *QDBusSignature
}

func (ptr *QDBusSignature) QDBusSignature_PTR() *QDBusSignature {
	return ptr
}

func (ptr *QDBusSignature) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusSignature) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusSignature(ptr QDBusSignature_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusSignature_PTR().Pointer()
	}
	return nil
}

func (n *QDBusSignature) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusSignatureFromPointer(ptr unsafe.Pointer) (n *QDBusSignature) {
	n = new(QDBusSignature)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusSignature")
	return
}

func (ptr *QDBusSignature) DestroyQDBusSignature() {
}

func NewQDBusSignature() *QDBusSignature {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusSignature", ""}).(*QDBusSignature)
}

func NewQDBusSignature2(signature string) *QDBusSignature {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusSignature2", "", signature}).(*QDBusSignature)
}

func NewQDBusSignature3(signature core.QLatin1String_ITF) *QDBusSignature {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusSignature3", "", signature}).(*QDBusSignature)
}

func NewQDBusSignature4(signature string) *QDBusSignature {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusSignature4", "", signature}).(*QDBusSignature)
}

func (ptr *QDBusSignature) SetSignature(signature string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSignature", signature})
}

func (ptr *QDBusSignature) Signature() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Signature"}).(string)
}

func (ptr *QDBusSignature) Swap(other QDBusSignature_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

type QDBusUnixFileDescriptor struct {
	internal.Internal
}

type QDBusUnixFileDescriptor_ITF interface {
	QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor
}

func (ptr *QDBusUnixFileDescriptor) QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor {
	return ptr
}

func (ptr *QDBusUnixFileDescriptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusUnixFileDescriptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusUnixFileDescriptor(ptr QDBusUnixFileDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusUnixFileDescriptor_PTR().Pointer()
	}
	return nil
}

func (n *QDBusUnixFileDescriptor) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) (n *QDBusUnixFileDescriptor) {
	n = new(QDBusUnixFileDescriptor)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusUnixFileDescriptor")
	return
}
func NewQDBusUnixFileDescriptor() *QDBusUnixFileDescriptor {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusUnixFileDescriptor", ""}).(*QDBusUnixFileDescriptor)
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusUnixFileDescriptor2", "", fileDescriptor}).(*QDBusUnixFileDescriptor)
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusUnixFileDescriptor3", "", other}).(*QDBusUnixFileDescriptor)
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileDescriptor"}).(float64))
}

func QDBusUnixFileDescriptor_IsSupported() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusUnixFileDescriptor_IsSupported", ""}).(bool)
}

func (ptr *QDBusUnixFileDescriptor) IsSupported() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.QDBusUnixFileDescriptor_IsSupported", ""}).(bool)
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFileDescriptor", fileDescriptor})
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusUnixFileDescriptor"})
}

type QDBusVariant struct {
	internal.Internal
}

type QDBusVariant_ITF interface {
	QDBusVariant_PTR() *QDBusVariant
}

func (ptr *QDBusVariant) QDBusVariant_PTR() *QDBusVariant {
	return ptr
}

func (ptr *QDBusVariant) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDBusVariant) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDBusVariant(ptr QDBusVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVariant_PTR().Pointer()
	}
	return nil
}

func (n *QDBusVariant) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDBusVariantFromPointer(ptr unsafe.Pointer) (n *QDBusVariant) {
	n = new(QDBusVariant)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusVariant")
	return
}

func (ptr *QDBusVariant) DestroyQDBusVariant() {
}

func NewQDBusVariant() *QDBusVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusVariant", ""}).(*QDBusVariant)
}

func NewQDBusVariant2(variant core.QVariant_ITF) *QDBusVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusVariant2", "", variant}).(*QDBusVariant)
}

func (ptr *QDBusVariant) SetVariant(variant core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVariant", variant})
}

func (ptr *QDBusVariant) Swap(other QDBusVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDBusVariant) Variant() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Variant"}).(*core.QVariant)
}

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObject_ITF interface {
	core.QObject_ITF
	QDBusVirtualObject_PTR() *QDBusVirtualObject
}

func (ptr *QDBusVirtualObject) QDBusVirtualObject_PTR() *QDBusVirtualObject {
	return ptr
}

func (ptr *QDBusVirtualObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusVirtualObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObject_PTR().Pointer()
	}
	return nil
}

func (n *QDBusVirtualObject) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDBusVirtualObject) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) (n *QDBusVirtualObject) {
	n = new(QDBusVirtualObject)
	n.InitFromInternal(uintptr(ptr), "dbus.QDBusVirtualObject")
	return
}
func NewQDBusVirtualObject(parent core.QObject_ITF) *QDBusVirtualObject {

	return internal.CallLocalFunction([]interface{}{"", "", "dbus.NewQDBusVirtualObject", "", parent}).(*QDBusVirtualObject)
}

func (ptr *QDBusVirtualObject) ConnectHandleMessage(f func(message *QDBusMessage, connection *QDBusConnection) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHandleMessage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusVirtualObject) DisconnectHandleMessage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHandleMessage"})
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HandleMessage", message, connection}).(bool)
}

func (ptr *QDBusVirtualObject) ConnectIntrospect(f func(path string) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIntrospect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusVirtualObject) DisconnectIntrospect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIntrospect"})
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Introspect", path}).(string)
}

func (ptr *QDBusVirtualObject) ConnectDestroyQDBusVirtualObject(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDBusVirtualObject", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDBusVirtualObject) DisconnectDestroyQDBusVirtualObject() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDBusVirtualObject"})
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusVirtualObject"})
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObjectDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDBusVirtualObjectDefault"})
}

func (ptr *QDBusVirtualObject) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDBusVirtualObject) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDBusVirtualObject) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusVirtualObject) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDBusVirtualObject) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDBusVirtualObject) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusVirtualObject) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDBusVirtualObject) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDBusVirtualObject) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDBusVirtualObject) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDBusVirtualObject) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDBusVirtualObject) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDBusVirtualObject) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDBusVirtualObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDBusVirtualObject) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDBusVirtualObject) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDBusVirtualObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDBusVirtualObject) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDBusVirtualObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDBusVirtualObject) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDBusVirtualObject) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QMetaTypeId2 struct {
	internal.Internal
}

type QMetaTypeId2_ITF interface {
	QMetaTypeId2_PTR() *QMetaTypeId2
}

func (ptr *QMetaTypeId2) QMetaTypeId2_PTR() *QMetaTypeId2 {
	return ptr
}

func (ptr *QMetaTypeId2) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QMetaTypeId2) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQMetaTypeId2(ptr QMetaTypeId2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaTypeId2_PTR().Pointer()
	}
	return nil
}

func (n *QMetaTypeId2) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQMetaTypeId2FromPointer(ptr unsafe.Pointer) (n *QMetaTypeId2) {
	n = new(QMetaTypeId2)
	n.InitFromInternal(uintptr(ptr), "dbus.QMetaTypeId2")
	return
}

func (ptr *QMetaTypeId2) DestroyQMetaTypeId2() {
}

func init() {
	internal.ConstructorTable["dbus.QDBus"] = NewQDBusFromPointer
	internal.ConstructorTable["dbus.QDBusAbstractAdaptor"] = NewQDBusAbstractAdaptorFromPointer
	internal.ConstructorTable["dbus.QDBusAbstractInterface"] = NewQDBusAbstractInterfaceFromPointer
	internal.ConstructorTable["dbus.QDBusAbstractInterfaceBase"] = NewQDBusAbstractInterfaceBaseFromPointer
	internal.ConstructorTable["dbus.QDBusArgument"] = NewQDBusArgumentFromPointer
	internal.ConstructorTable["dbus.QDBusConnection"] = NewQDBusConnectionFromPointer
	internal.ConstructorTable["dbus.QDBusConnectionInterface"] = NewQDBusConnectionInterfaceFromPointer
	internal.ConstructorTable["dbus.QDBusContext"] = NewQDBusContextFromPointer
	internal.ConstructorTable["dbus.QDBusError"] = NewQDBusErrorFromPointer
	internal.ConstructorTable["dbus.QDBusInterface"] = NewQDBusInterfaceFromPointer
	internal.ConstructorTable["dbus.QDBusMessage"] = NewQDBusMessageFromPointer
	internal.ConstructorTable["dbus.QDBusObjectPath"] = NewQDBusObjectPathFromPointer
	internal.ConstructorTable["dbus.QDBusPendingCall"] = NewQDBusPendingCallFromPointer
	internal.ConstructorTable["dbus.QDBusPendingCallWatcher"] = NewQDBusPendingCallWatcherFromPointer
	internal.ConstructorTable["dbus.QDBusPendingReplyData"] = NewQDBusPendingReplyDataFromPointer
	internal.ConstructorTable["dbus.QDBusServer"] = NewQDBusServerFromPointer
	internal.ConstructorTable["dbus.QDBusServiceWatcher"] = NewQDBusServiceWatcherFromPointer
	internal.ConstructorTable["dbus.QDBusSignature"] = NewQDBusSignatureFromPointer
	internal.ConstructorTable["dbus.QDBusUnixFileDescriptor"] = NewQDBusUnixFileDescriptorFromPointer
	internal.ConstructorTable["dbus.QDBusVariant"] = NewQDBusVariantFromPointer
	internal.ConstructorTable["dbus.QDBusVirtualObject"] = NewQDBusVirtualObjectFromPointer
	internal.ConstructorTable["dbus.QMetaTypeId2"] = NewQMetaTypeId2FromPointer
}
