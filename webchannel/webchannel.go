//go:build !minimal
// +build !minimal

package webchannel

import (
	"unsafe"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/internal"
)

type QQmlWebChannel struct {
	QWebChannel
}

type QQmlWebChannel_ITF interface {
	QWebChannel_ITF
	QQmlWebChannel_PTR() *QQmlWebChannel
}

func (ptr *QQmlWebChannel) QQmlWebChannel_PTR() *QQmlWebChannel {
	return ptr
}

func (ptr *QQmlWebChannel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannel_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlWebChannel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWebChannel_PTR().SetPointer(p)
	}
}

func PointerFromQQmlWebChannel(ptr QQmlWebChannel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlWebChannel_PTR().Pointer()
	}
	return nil
}

func (n *QQmlWebChannel) InitFromInternal(ptr uintptr, name string) {
	n.QWebChannel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlWebChannel) ClassNameInternalF() string {
	return n.QWebChannel_PTR().ClassNameInternalF()
}

func NewQQmlWebChannelFromPointer(ptr unsafe.Pointer) (n *QQmlWebChannel) {
	n = new(QQmlWebChannel)
	n.InitFromInternal(uintptr(ptr), "webchannel.QQmlWebChannel")
	return
}

func (ptr *QQmlWebChannel) DestroyQQmlWebChannel() {
}

func (ptr *QQmlWebChannel) __registerObjects_objects_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registerObjects_objects_atList", v, i}).(*core.QVariant)
}

func (ptr *QQmlWebChannel) __registerObjects_objects_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registerObjects_objects_setList", key, i})
}

type QWebChannel struct {
	core.QObject
}

type QWebChannel_ITF interface {
	core.QObject_ITF
	QWebChannel_PTR() *QWebChannel
}

func (ptr *QWebChannel) QWebChannel_PTR() *QWebChannel {
	return ptr
}

func (ptr *QWebChannel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebChannel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebChannel(ptr QWebChannel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannel_PTR().Pointer()
	}
	return nil
}

func (n *QWebChannel) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebChannel) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebChannelFromPointer(ptr unsafe.Pointer) (n *QWebChannel) {
	n = new(QWebChannel)
	n.InitFromInternal(uintptr(ptr), "webchannel.QWebChannel")
	return
}
func NewQWebChannel(parent core.QObject_ITF) *QWebChannel {

	return internal.CallLocalFunction([]interface{}{"", "", "webchannel.NewQWebChannel", "", parent}).(*QWebChannel)
}

func (ptr *QWebChannel) BlockUpdates() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BlockUpdates"}).(bool)
}

func (ptr *QWebChannel) ConnectBlockUpdatesChanged(f func(block bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBlockUpdatesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannel) DisconnectBlockUpdatesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBlockUpdatesChanged"})
}

func (ptr *QWebChannel) BlockUpdatesChanged(block bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BlockUpdatesChanged", block})
}

func (ptr *QWebChannel) ConnectConnectTo(f func(transport *QWebChannelAbstractTransport)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnectTo", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannel) DisconnectConnectTo() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnectTo"})
}

func (ptr *QWebChannel) ConnectTo(transport QWebChannelAbstractTransport_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTo", transport})
}

func (ptr *QWebChannel) ConnectToDefault(transport QWebChannelAbstractTransport_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToDefault", transport})
}

func (ptr *QWebChannel) DeregisterObject(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeregisterObject", object})
}

func (ptr *QWebChannel) ConnectDisconnectFrom(f func(transport *QWebChannelAbstractTransport)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnectFrom", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannel) DisconnectDisconnectFrom() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnectFrom"})
}

func (ptr *QWebChannel) DisconnectFrom(transport QWebChannelAbstractTransport_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFrom", transport})
}

func (ptr *QWebChannel) DisconnectFromDefault(transport QWebChannelAbstractTransport_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromDefault", transport})
}

func (ptr *QWebChannel) RegisterObject(id string, object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterObject", id, object})
}

func (ptr *QWebChannel) RegisterObjects(objects map[string]*core.QObject) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterObjects", objects})
}

func (ptr *QWebChannel) RegisteredObjects() map[string]*core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisteredObjects"}).(map[string]*core.QObject)
}

func (ptr *QWebChannel) SetBlockUpdates(block bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBlockUpdates", block})
}

func (ptr *QWebChannel) ConnectDestroyQWebChannel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebChannel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannel) DisconnectDestroyQWebChannel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebChannel"})
}

func (ptr *QWebChannel) DestroyQWebChannel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebChannel"})
}

func (ptr *QWebChannel) DestroyQWebChannelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebChannelDefault"})
}

func (ptr *QWebChannel) __registerObjects_objects_atList(v string, i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registerObjects_objects_atList", v, i}).(*core.QObject)
}

func (ptr *QWebChannel) __registerObjects_objects_setList(key string, i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registerObjects_objects_setList", key, i})
}

func (ptr *QWebChannel) __registerObjects_objects_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registerObjects_objects_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) __registerObjects_objects_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registerObjects_objects_keyList"}).([]string)
}

func (ptr *QWebChannel) __registeredObjects_atList(v string, i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registeredObjects_atList", v, i}).(*core.QObject)
}

func (ptr *QWebChannel) __registeredObjects_setList(key string, i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registeredObjects_setList", key, i})
}

func (ptr *QWebChannel) __registeredObjects_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registeredObjects_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) __registeredObjects_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__registeredObjects_keyList"}).([]string)
}

func (ptr *QWebChannel) ____registerObjects_objects_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____registerObjects_objects_keyList_atList", i}).(string)
}

func (ptr *QWebChannel) ____registerObjects_objects_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____registerObjects_objects_keyList_setList", i})
}

func (ptr *QWebChannel) ____registerObjects_objects_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____registerObjects_objects_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) ____registeredObjects_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____registeredObjects_keyList_atList", i}).(string)
}

func (ptr *QWebChannel) ____registeredObjects_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____registeredObjects_keyList_setList", i})
}

func (ptr *QWebChannel) ____registeredObjects_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____registeredObjects_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebChannel) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebChannel) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebChannel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebChannel) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebChannel) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebChannel) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebChannel) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebChannel) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebChannel) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebChannel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebChannel) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebChannel) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebChannel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebChannel) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebChannel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebChannel) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebChannel) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransport_ITF interface {
	core.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebChannelAbstractTransport(ptr QWebChannelAbstractTransport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelAbstractTransport_PTR().Pointer()
	}
	return nil
}

func (n *QWebChannelAbstractTransport) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebChannelAbstractTransport) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) (n *QWebChannelAbstractTransport) {
	n = new(QWebChannelAbstractTransport)
	n.InitFromInternal(uintptr(ptr), "webchannel.QWebChannelAbstractTransport")
	return
}
func NewQWebChannelAbstractTransport(parent core.QObject_ITF) *QWebChannelAbstractTransport {

	return internal.CallLocalFunction([]interface{}{"", "", "webchannel.NewQWebChannelAbstractTransport", "", parent}).(*QWebChannelAbstractTransport)
}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMessageReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMessageReceived"})
}

func (ptr *QWebChannelAbstractTransport) MessageReceived(message core.QJsonObject_ITF, transport QWebChannelAbstractTransport_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MessageReceived", message, transport})
}

func (ptr *QWebChannelAbstractTransport) ConnectSendMessage(f func(message *core.QJsonObject)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSendMessage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannelAbstractTransport) DisconnectSendMessage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSendMessage"})
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendMessage", message})
}

func (ptr *QWebChannelAbstractTransport) ConnectDestroyQWebChannelAbstractTransport(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebChannelAbstractTransport", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebChannelAbstractTransport) DisconnectDestroyQWebChannelAbstractTransport() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebChannelAbstractTransport"})
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebChannelAbstractTransport"})
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransportDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebChannelAbstractTransportDefault"})
}

func (ptr *QWebChannelAbstractTransport) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebChannelAbstractTransport) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebChannelAbstractTransport) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannelAbstractTransport) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebChannelAbstractTransport) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebChannelAbstractTransport) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannelAbstractTransport) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebChannelAbstractTransport) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebChannelAbstractTransport) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebChannelAbstractTransport) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebChannelAbstractTransport) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebChannelAbstractTransport) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebChannelAbstractTransport) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebChannelAbstractTransport) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebChannelAbstractTransport) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebChannelAbstractTransport) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebChannelAbstractTransport) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebChannelAbstractTransport) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebChannelAbstractTransport) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebChannelAbstractTransport) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebChannelAbstractTransport) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["webchannel.QQmlWebChannel"] = NewQQmlWebChannelFromPointer
	internal.ConstructorTable["webchannel.QWebChannel"] = NewQWebChannelFromPointer
	internal.ConstructorTable["webchannel.QWebChannelAbstractTransport"] = NewQWebChannelAbstractTransportFromPointer
}
