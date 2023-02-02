// +build !minimal

package scxml

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"strings"
	"unsafe"
)

type QScxmlCompiler struct {
	internal.Internal
}

type QScxmlCompiler_ITF interface {
	QScxmlCompiler_PTR() *QScxmlCompiler
}

func (ptr *QScxmlCompiler) QScxmlCompiler_PTR() *QScxmlCompiler {
	return ptr
}

func (ptr *QScxmlCompiler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScxmlCompiler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScxmlCompiler(ptr QScxmlCompiler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlCompiler_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlCompiler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScxmlCompilerFromPointer(ptr unsafe.Pointer) (n *QScxmlCompiler) {
	n = new(QScxmlCompiler)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlCompiler")
	return
}
func NewQScxmlCompiler(reader core.QXmlStreamReader_ITF) *QScxmlCompiler {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlCompiler", "", reader}).(*QScxmlCompiler)
}

func (ptr *QScxmlCompiler) Compile() *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Compile"}).(*QScxmlStateMachine)
}

func (ptr *QScxmlCompiler) Errors() []*QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Errors"}).([]*QScxmlError)
}

func (ptr *QScxmlCompiler) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QScxmlCompiler) SetFileName(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFileName", fileName})
}

func (ptr *QScxmlCompiler) DestroyQScxmlCompiler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlCompiler"})
}

func (ptr *QScxmlCompiler) __errors_atList(i int) *QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_atList", i}).(*QScxmlError)
}

func (ptr *QScxmlCompiler) __errors_setList(i QScxmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_setList", i})
}

func (ptr *QScxmlCompiler) __errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_newList"}).(unsafe.Pointer)
}

type QScxmlCppDataModel struct {
	QScxmlDataModel
}

type QScxmlCppDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlCppDataModel_PTR() *QScxmlCppDataModel
}

func (ptr *QScxmlCppDataModel) QScxmlCppDataModel_PTR() *QScxmlCppDataModel {
	return ptr
}

func (ptr *QScxmlCppDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlCppDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlDataModel_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlCppDataModel(ptr QScxmlCppDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlCppDataModel_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlCppDataModel) InitFromInternal(ptr uintptr, name string) {
	n.QScxmlDataModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlCppDataModel) ClassNameInternalF() string {
	return n.QScxmlDataModel_PTR().ClassNameInternalF()
}

func NewQScxmlCppDataModelFromPointer(ptr unsafe.Pointer) (n *QScxmlCppDataModel) {
	n = new(QScxmlCppDataModel)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlCppDataModel")
	return
}

func (ptr *QScxmlCppDataModel) DestroyQScxmlCppDataModel() {
}

func (ptr *QScxmlCppDataModel) ConnectHasScxmlProperty(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlCppDataModel) DisconnectHasScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasScxmlProperty"})
}

func (ptr *QScxmlCppDataModel) HasScxmlProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlProperty", name}).(bool)
}

func (ptr *QScxmlCppDataModel) HasScxmlPropertyDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlPropertyDefault", name}).(bool)
}

func (ptr *QScxmlCppDataModel) InState(stateName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InState", stateName}).(bool)
}

func (ptr *QScxmlCppDataModel) ScxmlEvent() *QScxmlEvent {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlEvent"}).(*QScxmlEvent)
}

func (ptr *QScxmlCppDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlCppDataModel) DisconnectScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScxmlProperty"})
}

func (ptr *QScxmlCppDataModel) ScxmlProperty(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlProperty", name}).(*core.QVariant)
}

func (ptr *QScxmlCppDataModel) ScxmlPropertyDefault(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlPropertyDefault", name}).(*core.QVariant)
}

func (ptr *QScxmlCppDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlEvent", event})
}

func (ptr *QScxmlCppDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlCppDataModel) DisconnectSetScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetScxmlProperty"})
}

func (ptr *QScxmlCppDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlProperty", name, value, context}).(bool)
}

func (ptr *QScxmlCppDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlPropertyDefault", name, value, context}).(bool)
}

func (ptr *QScxmlCppDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlCppDataModel) DisconnectSetup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetup"})
}

func (ptr *QScxmlCppDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Setup", initialDataValues}).(bool)
}

func (ptr *QScxmlCppDataModel) SetupDefault(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupDefault", initialDataValues}).(bool)
}

type QScxmlDataModel struct {
	core.QObject
}

type QScxmlDataModel_ITF interface {
	core.QObject_ITF
	QScxmlDataModel_PTR() *QScxmlDataModel
}

func (ptr *QScxmlDataModel) QScxmlDataModel_PTR() *QScxmlDataModel {
	return ptr
}

func (ptr *QScxmlDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlDataModel(ptr QScxmlDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlDataModel) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlDataModel) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScxmlDataModelFromPointer(ptr unsafe.Pointer) (n *QScxmlDataModel) {
	n = new(QScxmlDataModel)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlDataModel")
	return
}

func (ptr *QScxmlDataModel) DestroyQScxmlDataModel() {
}

func (ptr *QScxmlDataModel) ConnectHasScxmlProperty(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlDataModel) DisconnectHasScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasScxmlProperty"})
}

func (ptr *QScxmlDataModel) HasScxmlProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlProperty", name}).(bool)
}

func (ptr *QScxmlDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlDataModel) DisconnectScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScxmlProperty"})
}

func (ptr *QScxmlDataModel) ScxmlProperty(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlProperty", name}).(*core.QVariant)
}

func (ptr *QScxmlDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlDataModel) DisconnectSetScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetScxmlProperty"})
}

func (ptr *QScxmlDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlProperty", name, value, context}).(bool)
}

func (ptr *QScxmlDataModel) SetStateMachine(stateMachine QScxmlStateMachine_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStateMachine", stateMachine})
}

func (ptr *QScxmlDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlDataModel) DisconnectSetup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetup"})
}

func (ptr *QScxmlDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Setup", initialDataValues}).(bool)
}

func (ptr *QScxmlDataModel) StateMachine() *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateMachine"}).(*QScxmlStateMachine)
}

func (ptr *QScxmlDataModel) ConnectStateMachineChanged(f func(stateMachine *QScxmlStateMachine)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateMachineChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlDataModel) DisconnectStateMachineChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateMachineChanged"})
}

func (ptr *QScxmlDataModel) StateMachineChanged(stateMachine QScxmlStateMachine_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateMachineChanged", stateMachine})
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setup_initialDataValues_atList", v, i}).(*core.QVariant)
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setup_initialDataValues_setList", key, i})
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setup_initialDataValues_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setup_initialDataValues_keyList"}).([]string)
}

func (ptr *QScxmlDataModel) ____setup_initialDataValues_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setup_initialDataValues_keyList_atList", i}).(string)
}

func (ptr *QScxmlDataModel) ____setup_initialDataValues_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setup_initialDataValues_keyList_setList", i})
}

func (ptr *QScxmlDataModel) ____setup_initialDataValues_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setup_initialDataValues_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlDataModel) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScxmlDataModel) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScxmlDataModel) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlDataModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScxmlDataModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScxmlDataModel) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlDataModel) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScxmlDataModel) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScxmlDataModel) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlDataModel) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScxmlDataModel) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScxmlDataModel) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScxmlDataModel) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScxmlDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScxmlDataModel) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScxmlDataModel) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScxmlDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScxmlDataModel) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScxmlDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScxmlDataModel) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScxmlDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QScxmlDynamicScxmlServiceFactory struct {
	QScxmlInvokableServiceFactory
}

type QScxmlDynamicScxmlServiceFactory_ITF interface {
	QScxmlInvokableServiceFactory_ITF
	QScxmlDynamicScxmlServiceFactory_PTR() *QScxmlDynamicScxmlServiceFactory
}

func (ptr *QScxmlDynamicScxmlServiceFactory) QScxmlDynamicScxmlServiceFactory_PTR() *QScxmlDynamicScxmlServiceFactory {
	return ptr
}

func (ptr *QScxmlDynamicScxmlServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableServiceFactory_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlInvokableServiceFactory_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlDynamicScxmlServiceFactory(ptr QScxmlDynamicScxmlServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDynamicScxmlServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlDynamicScxmlServiceFactory) InitFromInternal(ptr uintptr, name string) {
	n.QScxmlInvokableServiceFactory_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlDynamicScxmlServiceFactory) ClassNameInternalF() string {
	return n.QScxmlInvokableServiceFactory_PTR().ClassNameInternalF()
}

func NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QScxmlDynamicScxmlServiceFactory) {
	n = new(QScxmlDynamicScxmlServiceFactory)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlDynamicScxmlServiceFactory")
	return
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DestroyQScxmlDynamicScxmlServiceFactory() {
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInvoke", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectInvoke() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInvoke"})
}

func (ptr *QScxmlDynamicScxmlServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Invoke", parentStateMachine}).(*QScxmlInvokableService)
}

func (ptr *QScxmlDynamicScxmlServiceFactory) InvokeDefault(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvokeDefault", parentStateMachine}).(*QScxmlInvokableService)
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlDynamicScxmlServiceFactory_names_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QScxmlDynamicScxmlServiceFactory_names_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlDynamicScxmlServiceFactory_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QScxmlDynamicScxmlServiceFactory_parameters_newList"}).(unsafe.Pointer)
}

type QScxmlEcmaScriptDataModel struct {
	QScxmlDataModel
}

type QScxmlEcmaScriptDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlEcmaScriptDataModel_PTR() *QScxmlEcmaScriptDataModel
}

func (ptr *QScxmlEcmaScriptDataModel) QScxmlEcmaScriptDataModel_PTR() *QScxmlEcmaScriptDataModel {
	return ptr
}

func (ptr *QScxmlEcmaScriptDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlDataModel_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlEcmaScriptDataModel(ptr QScxmlEcmaScriptDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlEcmaScriptDataModel_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlEcmaScriptDataModel) InitFromInternal(ptr uintptr, name string) {
	n.QScxmlDataModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlEcmaScriptDataModel) ClassNameInternalF() string {
	return n.QScxmlDataModel_PTR().ClassNameInternalF()
}

func NewQScxmlEcmaScriptDataModelFromPointer(ptr unsafe.Pointer) (n *QScxmlEcmaScriptDataModel) {
	n = new(QScxmlEcmaScriptDataModel)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlEcmaScriptDataModel")
	return
}

func (ptr *QScxmlEcmaScriptDataModel) DestroyQScxmlEcmaScriptDataModel() {
}

func NewQScxmlEcmaScriptDataModel(parent core.QObject_ITF) *QScxmlEcmaScriptDataModel {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlEcmaScriptDataModel", "", parent}).(*QScxmlEcmaScriptDataModel)
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectHasScxmlProperty(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectHasScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasScxmlProperty"})
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlProperty", name}).(bool)
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlPropertyDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlPropertyDefault", name}).(bool)
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScxmlProperty"})
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlProperty(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlProperty", name}).(*core.QVariant)
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlPropertyDefault(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlPropertyDefault", name}).(*core.QVariant)
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetScxmlEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetScxmlEvent"})
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlEvent", event})
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlEventDefault(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlEventDefault", event})
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetScxmlProperty"})
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlProperty", name, value, context}).(bool)
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlPropertyDefault", name, value, context}).(bool)
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetup"})
}

func (ptr *QScxmlEcmaScriptDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Setup", initialDataValues}).(bool)
}

func (ptr *QScxmlEcmaScriptDataModel) SetupDefault(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupDefault", initialDataValues}).(bool)
}

type QScxmlError struct {
	internal.Internal
}

type QScxmlError_ITF interface {
	QScxmlError_PTR() *QScxmlError
}

func (ptr *QScxmlError) QScxmlError_PTR() *QScxmlError {
	return ptr
}

func (ptr *QScxmlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScxmlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScxmlError(ptr QScxmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlError_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScxmlErrorFromPointer(ptr unsafe.Pointer) (n *QScxmlError) {
	n = new(QScxmlError)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlError")
	return
}
func NewQScxmlError() *QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlError", ""}).(*QScxmlError)
}

func NewQScxmlError2(fileName string, line int, column int, description string) *QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlError2", "", fileName, line, column, description}).(*QScxmlError)
}

func NewQScxmlError3(other QScxmlError_ITF) *QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlError3", "", other}).(*QScxmlError)
}

func (ptr *QScxmlError) Column() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Column"}).(float64))
}

func (ptr *QScxmlError) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QScxmlError) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QScxmlError) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QScxmlError) Line() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Line"}).(float64))
}

func (ptr *QScxmlError) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QScxmlError) DestroyQScxmlError() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlError"})
}

type QScxmlEvent struct {
	internal.Internal
}

type QScxmlEvent_ITF interface {
	QScxmlEvent_PTR() *QScxmlEvent
}

func (ptr *QScxmlEvent) QScxmlEvent_PTR() *QScxmlEvent {
	return ptr
}

func (ptr *QScxmlEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScxmlEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScxmlEvent(ptr QScxmlEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlEvent_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlEvent) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScxmlEventFromPointer(ptr unsafe.Pointer) (n *QScxmlEvent) {
	n = new(QScxmlEvent)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlEvent")
	return
}

//go:generate stringer -type=QScxmlEvent__EventType
//QScxmlEvent::EventType
type QScxmlEvent__EventType int64

const (
	QScxmlEvent__PlatformEvent QScxmlEvent__EventType = QScxmlEvent__EventType(0)
	QScxmlEvent__InternalEvent QScxmlEvent__EventType = QScxmlEvent__EventType(1)
	QScxmlEvent__ExternalEvent QScxmlEvent__EventType = QScxmlEvent__EventType(2)
)

func NewQScxmlEvent() *QScxmlEvent {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlEvent", ""}).(*QScxmlEvent)
}

func NewQScxmlEvent2(other QScxmlEvent_ITF) *QScxmlEvent {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlEvent2", "", other}).(*QScxmlEvent)
}

func (ptr *QScxmlEvent) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QScxmlEvent) Data() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(*core.QVariant)
}

func (ptr *QScxmlEvent) Delay() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Delay"}).(float64))
}

func (ptr *QScxmlEvent) ErrorMessage() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorMessage"}).(string)
}

func (ptr *QScxmlEvent) EventType() QScxmlEvent__EventType {

	return QScxmlEvent__EventType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventType"}).(float64))
}

func (ptr *QScxmlEvent) InvokeId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvokeId"}).(string)
}

func (ptr *QScxmlEvent) IsErrorEvent() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsErrorEvent"}).(bool)
}

func (ptr *QScxmlEvent) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QScxmlEvent) Origin() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Origin"}).(string)
}

func (ptr *QScxmlEvent) OriginType() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OriginType"}).(string)
}

func (ptr *QScxmlEvent) ScxmlType() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlType"}).(string)
}

func (ptr *QScxmlEvent) SendId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendId"}).(string)
}

func (ptr *QScxmlEvent) SetData(data core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", data})
}

func (ptr *QScxmlEvent) SetDelay(delayInMiliSecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDelay", delayInMiliSecs})
}

func (ptr *QScxmlEvent) SetErrorMessage(message string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetErrorMessage", message})
}

func (ptr *QScxmlEvent) SetEventType(ty QScxmlEvent__EventType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEventType", ty})
}

func (ptr *QScxmlEvent) SetInvokeId(invokeid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInvokeId", invokeid})
}

func (ptr *QScxmlEvent) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QScxmlEvent) SetOrigin(origin string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrigin", origin})
}

func (ptr *QScxmlEvent) SetOriginType(origintype string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOriginType", origintype})
}

func (ptr *QScxmlEvent) SetSendId(sendid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSendId", sendid})
}

func (ptr *QScxmlEvent) DestroyQScxmlEvent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlEvent"})
}

type QScxmlInvokableService struct {
	core.QObject
}

type QScxmlInvokableService_ITF interface {
	core.QObject_ITF
	QScxmlInvokableService_PTR() *QScxmlInvokableService
}

func (ptr *QScxmlInvokableService) QScxmlInvokableService_PTR() *QScxmlInvokableService {
	return ptr
}

func (ptr *QScxmlInvokableService) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlInvokableService) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlInvokableService(ptr QScxmlInvokableService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableService_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlInvokableService) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlInvokableService) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScxmlInvokableServiceFromPointer(ptr unsafe.Pointer) (n *QScxmlInvokableService) {
	n = new(QScxmlInvokableService)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlInvokableService")
	return
}

func (ptr *QScxmlInvokableService) DestroyQScxmlInvokableService() {
}

func (ptr *QScxmlInvokableService) ConnectId(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectId", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlInvokableService) DisconnectId() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectId"})
}

func (ptr *QScxmlInvokableService) Id() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Id"}).(string)
}

func (ptr *QScxmlInvokableService) ConnectName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlInvokableService) DisconnectName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectName"})
}

func (ptr *QScxmlInvokableService) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QScxmlInvokableService) ParentStateMachine() *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentStateMachine"}).(*QScxmlStateMachine)
}

func (ptr *QScxmlInvokableService) ConnectPostEvent(f func(event *QScxmlEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPostEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlInvokableService) DisconnectPostEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPostEvent"})
}

func (ptr *QScxmlInvokableService) PostEvent(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PostEvent", event})
}

func (ptr *QScxmlInvokableService) ConnectStart(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlInvokableService) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QScxmlInvokableService) Start() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start"}).(bool)
}

func (ptr *QScxmlInvokableService) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScxmlInvokableService) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScxmlInvokableService) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableService) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScxmlInvokableService) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScxmlInvokableService) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableService) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScxmlInvokableService) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScxmlInvokableService) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableService) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScxmlInvokableService) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScxmlInvokableService) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableService) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScxmlInvokableService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScxmlInvokableService) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScxmlInvokableService) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScxmlInvokableService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScxmlInvokableService) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScxmlInvokableService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScxmlInvokableService) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScxmlInvokableService) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QScxmlInvokableServiceFactory struct {
	core.QObject
}

type QScxmlInvokableServiceFactory_ITF interface {
	core.QObject_ITF
	QScxmlInvokableServiceFactory_PTR() *QScxmlInvokableServiceFactory
}

func (ptr *QScxmlInvokableServiceFactory) QScxmlInvokableServiceFactory_PTR() *QScxmlInvokableServiceFactory {
	return ptr
}

func (ptr *QScxmlInvokableServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlInvokableServiceFactory(ptr QScxmlInvokableServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlInvokableServiceFactory) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlInvokableServiceFactory) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScxmlInvokableServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QScxmlInvokableServiceFactory) {
	n = new(QScxmlInvokableServiceFactory)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlInvokableServiceFactory")
	return
}

func (ptr *QScxmlInvokableServiceFactory) DestroyQScxmlInvokableServiceFactory() {
}

func (ptr *QScxmlInvokableServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInvoke", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectInvoke() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInvoke"})
}

func (ptr *QScxmlInvokableServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Invoke", parentStateMachine}).(*QScxmlInvokableService)
}

func (ptr *QScxmlInvokableServiceFactory) __QScxmlInvokableServiceFactory_names_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QScxmlInvokableServiceFactory_names_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __QScxmlInvokableServiceFactory_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QScxmlInvokableServiceFactory_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __names_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__names_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScxmlInvokableServiceFactory) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScxmlInvokableServiceFactory) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScxmlInvokableServiceFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScxmlInvokableServiceFactory) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScxmlInvokableServiceFactory) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScxmlInvokableServiceFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScxmlInvokableServiceFactory) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScxmlInvokableServiceFactory) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScxmlInvokableServiceFactory) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScxmlInvokableServiceFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScxmlInvokableServiceFactory) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScxmlInvokableServiceFactory) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QScxmlNullDataModel struct {
	QScxmlDataModel
}

type QScxmlNullDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlNullDataModel_PTR() *QScxmlNullDataModel
}

func (ptr *QScxmlNullDataModel) QScxmlNullDataModel_PTR() *QScxmlNullDataModel {
	return ptr
}

func (ptr *QScxmlNullDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlNullDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlDataModel_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlNullDataModel(ptr QScxmlNullDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlNullDataModel_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlNullDataModel) InitFromInternal(ptr uintptr, name string) {
	n.QScxmlDataModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlNullDataModel) ClassNameInternalF() string {
	return n.QScxmlDataModel_PTR().ClassNameInternalF()
}

func NewQScxmlNullDataModelFromPointer(ptr unsafe.Pointer) (n *QScxmlNullDataModel) {
	n = new(QScxmlNullDataModel)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlNullDataModel")
	return
}
func NewQScxmlNullDataModel(parent core.QObject_ITF) *QScxmlNullDataModel {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.NewQScxmlNullDataModel", "", parent}).(*QScxmlNullDataModel)
}

func (ptr *QScxmlNullDataModel) ConnectHasScxmlProperty(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlNullDataModel) DisconnectHasScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasScxmlProperty"})
}

func (ptr *QScxmlNullDataModel) HasScxmlProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlProperty", name}).(bool)
}

func (ptr *QScxmlNullDataModel) HasScxmlPropertyDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasScxmlPropertyDefault", name}).(bool)
}

func (ptr *QScxmlNullDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlNullDataModel) DisconnectScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScxmlProperty"})
}

func (ptr *QScxmlNullDataModel) ScxmlProperty(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlProperty", name}).(*core.QVariant)
}

func (ptr *QScxmlNullDataModel) ScxmlPropertyDefault(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScxmlPropertyDefault", name}).(*core.QVariant)
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetScxmlEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetScxmlEvent"})
}

func (ptr *QScxmlNullDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlEvent", event})
}

func (ptr *QScxmlNullDataModel) SetScxmlEventDefault(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlEventDefault", event})
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetScxmlProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetScxmlProperty"})
}

func (ptr *QScxmlNullDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlProperty", name, value, context}).(bool)
}

func (ptr *QScxmlNullDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScxmlPropertyDefault", name, value, context}).(bool)
}

func (ptr *QScxmlNullDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlNullDataModel) DisconnectSetup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetup"})
}

func (ptr *QScxmlNullDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Setup", initialDataValues}).(bool)
}

func (ptr *QScxmlNullDataModel) SetupDefault(initialDataValues map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupDefault", initialDataValues}).(bool)
}

func (ptr *QScxmlNullDataModel) ConnectDestroyQScxmlNullDataModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScxmlNullDataModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlNullDataModel) DisconnectDestroyQScxmlNullDataModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScxmlNullDataModel"})
}

func (ptr *QScxmlNullDataModel) DestroyQScxmlNullDataModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlNullDataModel"})
}

func (ptr *QScxmlNullDataModel) DestroyQScxmlNullDataModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlNullDataModelDefault"})
}

type QScxmlStateMachine struct {
	core.QObject
}

type QScxmlStateMachine_ITF interface {
	core.QObject_ITF
	QScxmlStateMachine_PTR() *QScxmlStateMachine
}

func (ptr *QScxmlStateMachine) QScxmlStateMachine_PTR() *QScxmlStateMachine {
	return ptr
}

func (ptr *QScxmlStateMachine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlStateMachine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlStateMachine(ptr QScxmlStateMachine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlStateMachine_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlStateMachine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlStateMachine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScxmlStateMachineFromPointer(ptr unsafe.Pointer) (n *QScxmlStateMachine) {
	n = new(QScxmlStateMachine)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlStateMachine")
	return
}

func (ptr *QScxmlStateMachine) DestroyQScxmlStateMachine() {
}

func (ptr *QScxmlStateMachine) ActiveStateNames(compress bool) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveStateNames", compress}).([]string)
}

func (ptr *QScxmlStateMachine) CancelDelayedEvent(sendId string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CancelDelayedEvent", sendId})
}

func (ptr *QScxmlStateMachine) DataModel() *QScxmlDataModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataModel"}).(*QScxmlDataModel)
}

func (ptr *QScxmlStateMachine) ConnectDataModelChanged(f func(model *QScxmlDataModel)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataModelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectDataModelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataModelChanged"})
}

func (ptr *QScxmlStateMachine) DataModelChanged(model QScxmlDataModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataModelChanged", model})
}

func (ptr *QScxmlStateMachine) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QScxmlStateMachine) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func QScxmlStateMachine_FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.QScxmlStateMachine_FromData", "", data, fileName}).(*QScxmlStateMachine)
}

func (ptr *QScxmlStateMachine) FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.QScxmlStateMachine_FromData", "", data, fileName}).(*QScxmlStateMachine)
}

func QScxmlStateMachine_FromFile(fileName string) *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.QScxmlStateMachine_FromFile", "", fileName}).(*QScxmlStateMachine)
}

func (ptr *QScxmlStateMachine) FromFile(fileName string) *QScxmlStateMachine {

	return internal.CallLocalFunction([]interface{}{"", "", "scxml.QScxmlStateMachine_FromFile", "", fileName}).(*QScxmlStateMachine)
}

func (ptr *QScxmlStateMachine) ConnectInit(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInit", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectInit() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInit"})
}

func (ptr *QScxmlStateMachine) Init() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Init"}).(bool)
}

func (ptr *QScxmlStateMachine) InitDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitDefault"}).(bool)
}

func (ptr *QScxmlStateMachine) InitialValues() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitialValues"}).(map[string]*core.QVariant)
}

func (ptr *QScxmlStateMachine) ConnectInitializedChanged(f func(initialized bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInitializedChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectInitializedChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInitializedChanged"})
}

func (ptr *QScxmlStateMachine) InitializedChanged(initialized bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitializedChanged", initialized})
}

func (ptr *QScxmlStateMachine) InvokedServices() []*QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvokedServices"}).([]*QScxmlInvokableService)
}

func (ptr *QScxmlStateMachine) ConnectInvokedServicesChanged(f func(invokedServices []*QScxmlInvokableService)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInvokedServicesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectInvokedServicesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInvokedServicesChanged"})
}

func (ptr *QScxmlStateMachine) InvokedServicesChanged(invokedServices []*QScxmlInvokableService) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvokedServicesChanged", invokedServices})
}

func (ptr *QScxmlStateMachine) IsActive(scxmlStateName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive", scxmlStateName}).(bool)
}

func (ptr *QScxmlStateMachine) IsActive2(stateIndex int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive2", stateIndex}).(bool)
}

func (ptr *QScxmlStateMachine) IsDispatchableTarget(target string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDispatchableTarget", target}).(bool)
}

func (ptr *QScxmlStateMachine) IsInitialized() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInitialized"}).(bool)
}

func (ptr *QScxmlStateMachine) IsInvoked() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInvoked"}).(bool)
}

func (ptr *QScxmlStateMachine) IsRunning() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRunning"}).(bool)
}

func (ptr *QScxmlStateMachine) DisconnectLoaderChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoaderChanged"})
}

func (ptr *QScxmlStateMachine) ConnectLog(f func(label string, msg string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLog", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectLog() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLog"})
}

func (ptr *QScxmlStateMachine) Log(label string, msg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Log", label, msg})
}

func (ptr *QScxmlStateMachine) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QScxmlStateMachine) ParseErrors() []*QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParseErrors"}).([]*QScxmlError)
}

func (ptr *QScxmlStateMachine) ConnectReachedStableState(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReachedStableState", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectReachedStableState() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReachedStableState"})
}

func (ptr *QScxmlStateMachine) ReachedStableState() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReachedStableState"})
}

func (ptr *QScxmlStateMachine) ConnectRunningChanged(f func(running bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRunningChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectRunningChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRunningChanged"})
}

func (ptr *QScxmlStateMachine) RunningChanged(running bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RunningChanged", running})
}

func (ptr *QScxmlStateMachine) SessionId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionId"}).(string)
}

func (ptr *QScxmlStateMachine) SetDataModel(model QScxmlDataModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataModel", model})
}

func (ptr *QScxmlStateMachine) SetInitialValues(initialValues map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInitialValues", initialValues})
}

func (ptr *QScxmlStateMachine) SetRunning(running bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRunning", running})
}

func (ptr *QScxmlStateMachine) SetTableData(tableData QScxmlTableData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTableData", tableData})
}

func (ptr *QScxmlStateMachine) ConnectStart(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QScxmlStateMachine) Start() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start"})
}

func (ptr *QScxmlStateMachine) StartDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDefault"})
}

func (ptr *QScxmlStateMachine) StateNames(compress bool) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateNames", compress}).([]string)
}

func (ptr *QScxmlStateMachine) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QScxmlStateMachine) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QScxmlStateMachine) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func (ptr *QScxmlStateMachine) SubmitEvent(event QScxmlEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitEvent", event})
}

func (ptr *QScxmlStateMachine) SubmitEvent2(eventName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitEvent2", eventName})
}

func (ptr *QScxmlStateMachine) SubmitEvent3(eventName string, data core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitEvent3", eventName, data})
}

func (ptr *QScxmlStateMachine) TableData() *QScxmlTableData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TableData"}).(*QScxmlTableData)
}

func (ptr *QScxmlStateMachine) ConnectTableDataChanged(f func(tableData *QScxmlTableData)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTableDataChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStateMachine) DisconnectTableDataChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTableDataChanged"})
}

func (ptr *QScxmlStateMachine) TableDataChanged(tableData QScxmlTableData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TableDataChanged", tableData})
}

func (ptr *QScxmlStateMachine) __initialValues_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValues_atList", v, i}).(*core.QVariant)
}

func (ptr *QScxmlStateMachine) __initialValues_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValues_setList", key, i})
}

func (ptr *QScxmlStateMachine) __initialValues_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValues_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __initialValues_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValues_keyList"}).([]string)
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValuesChanged_initialValues_atList", v, i}).(*core.QVariant)
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValuesChanged_initialValues_setList", key, i})
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValuesChanged_initialValues_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__initialValuesChanged_initialValues_keyList"}).([]string)
}

func (ptr *QScxmlStateMachine) __invokedServices_atList(i int) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__invokedServices_atList", i}).(*QScxmlInvokableService)
}

func (ptr *QScxmlStateMachine) __invokedServices_setList(i QScxmlInvokableService_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__invokedServices_setList", i})
}

func (ptr *QScxmlStateMachine) __invokedServices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__invokedServices_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __invokedServicesChanged_invokedServices_atList(i int) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__invokedServicesChanged_invokedServices_atList", i}).(*QScxmlInvokableService)
}

func (ptr *QScxmlStateMachine) __invokedServicesChanged_invokedServices_setList(i QScxmlInvokableService_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__invokedServicesChanged_invokedServices_setList", i})
}

func (ptr *QScxmlStateMachine) __invokedServicesChanged_invokedServices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__invokedServicesChanged_invokedServices_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __parseErrors_atList(i int) *QScxmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parseErrors_atList", i}).(*QScxmlError)
}

func (ptr *QScxmlStateMachine) __parseErrors_setList(i QScxmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parseErrors_setList", i})
}

func (ptr *QScxmlStateMachine) __parseErrors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parseErrors_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setInitialValues_initialValues_atList", v, i}).(*core.QVariant)
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setInitialValues_initialValues_setList", key, i})
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setInitialValues_initialValues_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setInitialValues_initialValues_keyList"}).([]string)
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____initialValues_keyList_atList", i}).(string)
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____initialValues_keyList_setList", i})
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____initialValues_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_initialValues_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____initialValuesChanged_initialValues_keyList_atList", i}).(string)
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_initialValues_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____initialValuesChanged_initialValues_keyList_setList", i})
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_initialValues_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____initialValuesChanged_initialValues_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) ____setInitialValues_initialValues_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setInitialValues_initialValues_keyList_atList", i}).(string)
}

func (ptr *QScxmlStateMachine) ____setInitialValues_initialValues_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setInitialValues_initialValues_keyList_setList", i})
}

func (ptr *QScxmlStateMachine) ____setInitialValues_initialValues_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setInitialValues_initialValues_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScxmlStateMachine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScxmlStateMachine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScxmlStateMachine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScxmlStateMachine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScxmlStateMachine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScxmlStateMachine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScxmlStateMachine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScxmlStateMachine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScxmlStateMachine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScxmlStateMachine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScxmlStateMachine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScxmlStateMachine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScxmlStateMachine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScxmlStateMachine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScxmlStateMachine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScxmlStateMachine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScxmlStateMachine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QScxmlStaticScxmlServiceFactory struct {
	QScxmlInvokableServiceFactory
}

type QScxmlStaticScxmlServiceFactory_ITF interface {
	QScxmlInvokableServiceFactory_ITF
	QScxmlStaticScxmlServiceFactory_PTR() *QScxmlStaticScxmlServiceFactory
}

func (ptr *QScxmlStaticScxmlServiceFactory) QScxmlStaticScxmlServiceFactory_PTR() *QScxmlStaticScxmlServiceFactory {
	return ptr
}

func (ptr *QScxmlStaticScxmlServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableServiceFactory_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlInvokableServiceFactory_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlStaticScxmlServiceFactory(ptr QScxmlStaticScxmlServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlStaticScxmlServiceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlStaticScxmlServiceFactory) InitFromInternal(ptr uintptr, name string) {
	n.QScxmlInvokableServiceFactory_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScxmlStaticScxmlServiceFactory) ClassNameInternalF() string {
	return n.QScxmlInvokableServiceFactory_PTR().ClassNameInternalF()
}

func NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr unsafe.Pointer) (n *QScxmlStaticScxmlServiceFactory) {
	n = new(QScxmlStaticScxmlServiceFactory)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlStaticScxmlServiceFactory")
	return
}

func (ptr *QScxmlStaticScxmlServiceFactory) DestroyQScxmlStaticScxmlServiceFactory() {
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInvoke", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectInvoke() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInvoke"})
}

func (ptr *QScxmlStaticScxmlServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Invoke", parentStateMachine}).(*QScxmlInvokableService)
}

func (ptr *QScxmlStaticScxmlServiceFactory) InvokeDefault(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvokeDefault", parentStateMachine}).(*QScxmlInvokableService)
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlStaticScxmlServiceFactory_nameList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QScxmlStaticScxmlServiceFactory_nameList_newList"}).(unsafe.Pointer)
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlStaticScxmlServiceFactory_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QScxmlStaticScxmlServiceFactory_parameters_newList"}).(unsafe.Pointer)
}

type QScxmlTableData struct {
	internal.Internal
}

type QScxmlTableData_ITF interface {
	QScxmlTableData_PTR() *QScxmlTableData
}

func (ptr *QScxmlTableData) QScxmlTableData_PTR() *QScxmlTableData {
	return ptr
}

func (ptr *QScxmlTableData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScxmlTableData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScxmlTableData(ptr QScxmlTableData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlTableData_PTR().Pointer()
	}
	return nil
}

func (n *QScxmlTableData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScxmlTableDataFromPointer(ptr unsafe.Pointer) (n *QScxmlTableData) {
	n = new(QScxmlTableData)
	n.InitFromInternal(uintptr(ptr), "scxml.QScxmlTableData")
	return
}
func (ptr *QScxmlTableData) ConnectName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlTableData) DisconnectName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectName"})
}

func (ptr *QScxmlTableData) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QScxmlTableData) ConnectServiceFactory(f func(id int) *QScxmlInvokableServiceFactory) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlTableData) DisconnectServiceFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceFactory"})
}

func (ptr *QScxmlTableData) ServiceFactory(id int) *QScxmlInvokableServiceFactory {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceFactory", id}).(*QScxmlInvokableServiceFactory)
}

func (ptr *QScxmlTableData) ConnectDestroyQScxmlTableData(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScxmlTableData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScxmlTableData) DisconnectDestroyQScxmlTableData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScxmlTableData"})
}

func (ptr *QScxmlTableData) DestroyQScxmlTableData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlTableData"})
}

func (ptr *QScxmlTableData) DestroyQScxmlTableDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScxmlTableDataDefault"})
}

func init() {
	internal.ConstructorTable["scxml.QScxmlCompiler"] = NewQScxmlCompilerFromPointer
	internal.ConstructorTable["scxml.QScxmlCppDataModel"] = NewQScxmlCppDataModelFromPointer
	internal.ConstructorTable["scxml.QScxmlDataModel"] = NewQScxmlDataModelFromPointer
	internal.ConstructorTable["scxml.QScxmlDynamicScxmlServiceFactory"] = NewQScxmlDynamicScxmlServiceFactoryFromPointer
	internal.ConstructorTable["scxml.QScxmlEcmaScriptDataModel"] = NewQScxmlEcmaScriptDataModelFromPointer
	internal.ConstructorTable["scxml.QScxmlError"] = NewQScxmlErrorFromPointer
	internal.ConstructorTable["scxml.QScxmlEvent"] = NewQScxmlEventFromPointer
	internal.ConstructorTable["scxml.QScxmlInvokableService"] = NewQScxmlInvokableServiceFromPointer
	internal.ConstructorTable["scxml.QScxmlInvokableServiceFactory"] = NewQScxmlInvokableServiceFactoryFromPointer
	internal.ConstructorTable["scxml.QScxmlNullDataModel"] = NewQScxmlNullDataModelFromPointer
	internal.ConstructorTable["scxml.QScxmlStateMachine"] = NewQScxmlStateMachineFromPointer
	internal.ConstructorTable["scxml.QScxmlStaticScxmlServiceFactory"] = NewQScxmlStaticScxmlServiceFactoryFromPointer
	internal.ConstructorTable["scxml.QScxmlTableData"] = NewQScxmlTableDataFromPointer
}
