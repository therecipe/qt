// +build !minimal

package qml

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/internal"
	"github.com/dev-drprasad/qt/network"
	"unsafe"
)

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func (ptr *QJSEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QJSEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func (n *QJSEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QJSEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) (n *QJSEngine) {
	n = new(QJSEngine)
	n.InitFromInternal(uintptr(ptr), "qml.QJSEngine")
	return
}

//go:generate stringer -type=QJSEngine__Extension
//QJSEngine::Extension
type QJSEngine__Extension int64

const (
	QJSEngine__TranslationExtension       QJSEngine__Extension = QJSEngine__Extension(0x1)
	QJSEngine__ConsoleExtension           QJSEngine__Extension = QJSEngine__Extension(0x2)
	QJSEngine__GarbageCollectionExtension QJSEngine__Extension = QJSEngine__Extension(0x4)
	QJSEngine__AllExtensions              QJSEngine__Extension = QJSEngine__Extension(0xffffffff)
)

func NewQJSEngine() *QJSEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSEngine", ""}).(*QJSEngine)
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSEngine2", "", parent}).(*QJSEngine)
}

func (ptr *QJSEngine) CollectGarbage() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollectGarbage"})
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Evaluate", program, fileName, lineNumber}).(*QJSValue)
}

func (ptr *QJSEngine) FromScriptValue(value QJSValue_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FromScriptValue", value}).(*core.QVariant)
}

func (ptr *QJSEngine) GlobalObject() *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GlobalObject"}).(*QJSValue)
}

func (ptr *QJSEngine) ImportModule(fileName string) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImportModule", fileName}).(*QJSValue)
}

func (ptr *QJSEngine) InstallExtensions(extensions QJSEngine__Extension, object QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstallExtensions", extensions, object})
}

func (ptr *QJSEngine) NewArray(length uint) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewArray", length}).(*QJSValue)
}

func (ptr *QJSEngine) NewErrorObject(errorType QJSValue__ErrorType, message string) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewErrorObject", errorType, message}).(*QJSValue)
}

func (ptr *QJSEngine) NewObject() *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewObject"}).(*QJSValue)
}

func (ptr *QJSEngine) NewQMetaObject(metaObject core.QMetaObject_ITF) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewQMetaObject", metaObject}).(*QJSValue)
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewQObject", object}).(*QJSValue)
}

func QJSEngine_qjsEngine(object core.QObject_ITF) *QJSEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QJSEngine_qjsEngine", "", object}).(*QJSEngine)
}

func (ptr *QJSEngine) qjsEngine(object core.QObject_ITF) *QJSEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QJSEngine_qjsEngine", "", object}).(*QJSEngine)
}

func (ptr *QJSEngine) ThrowError(message string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThrowError", message})
}

func (ptr *QJSEngine) ThrowError2(errorType QJSValue__ErrorType, message string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThrowError2", errorType, message})
}

func (ptr *QJSEngine) ToScriptValue(value core.QVariant_ITF) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToScriptValue", value}).(*QJSValue)
}

func (ptr *QJSEngine) ConnectDestroyQJSEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQJSEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QJSEngine) DisconnectDestroyQJSEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQJSEngine"})
}

func (ptr *QJSEngine) DestroyQJSEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQJSEngine"})
}

func (ptr *QJSEngine) DestroyQJSEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQJSEngineDefault"})
}

func (ptr *QJSEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QJSEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QJSEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QJSEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QJSEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QJSEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QJSEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QJSEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QJSEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QJSEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QJSEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QJSEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QJSEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QJSEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QJSEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QJSEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QJSEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QJSEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QJSValue struct {
	internal.Internal
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue {
	return ptr
}

func (ptr *QJSValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QJSValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQJSValue(ptr QJSValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValue_PTR().Pointer()
	}
	return nil
}

func (n *QJSValue) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQJSValueFromPointer(ptr unsafe.Pointer) (n *QJSValue) {
	n = new(QJSValue)
	n.InitFromInternal(uintptr(ptr), "qml.QJSValue")
	return
}

//go:generate stringer -type=QJSValue__SpecialValue
//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      QJSValue__SpecialValue = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue QJSValue__SpecialValue = QJSValue__SpecialValue(1)
)

//go:generate stringer -type=QJSValue__ErrorType
//QJSValue::ErrorType
type QJSValue__ErrorType int64

const (
	QJSValue__NoError        QJSValue__ErrorType = QJSValue__ErrorType(0)
	QJSValue__GenericError   QJSValue__ErrorType = QJSValue__ErrorType(1)
	QJSValue__EvalError      QJSValue__ErrorType = QJSValue__ErrorType(2)
	QJSValue__RangeError     QJSValue__ErrorType = QJSValue__ErrorType(3)
	QJSValue__ReferenceError QJSValue__ErrorType = QJSValue__ErrorType(4)
	QJSValue__SyntaxError    QJSValue__ErrorType = QJSValue__ErrorType(5)
	QJSValue__TypeError      QJSValue__ErrorType = QJSValue__ErrorType(6)
	QJSValue__URIError       QJSValue__ErrorType = QJSValue__ErrorType(7)
)

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue", "", value}).(*QJSValue)
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue2", "", other}).(*QJSValue)
}

func NewQJSValue3(other QJSValue_ITF) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue3", "", other}).(*QJSValue)
}

func NewQJSValue4(value bool) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue4", "", value}).(*QJSValue)
}

func NewQJSValue5(value int) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue5", "", value}).(*QJSValue)
}

func NewQJSValue6(value uint) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue6", "", value}).(*QJSValue)
}

func NewQJSValue7(value float64) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue7", "", value}).(*QJSValue)
}

func NewQJSValue8(value string) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue8", "", value}).(*QJSValue)
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue9", "", value}).(*QJSValue)
}

func NewQJSValue10(value string) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValue10", "", value}).(*QJSValue)
}

func (ptr *QJSValue) Call(args []*QJSValue) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Call", args}).(*QJSValue)
}

func (ptr *QJSValue) CallAsConstructor(args []*QJSValue) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallAsConstructor", args}).(*QJSValue)
}

func (ptr *QJSValue) CallWithInstance(instance QJSValue_ITF, args []*QJSValue) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CallWithInstance", instance, args}).(*QJSValue)
}

func (ptr *QJSValue) DeleteProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteProperty", name}).(bool)
}

func (ptr *QJSValue) Equals(other QJSValue_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Equals", other}).(bool)
}

func (ptr *QJSValue) ErrorType() QJSValue__ErrorType {

	return QJSValue__ErrorType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorType"}).(float64))
}

func (ptr *QJSValue) HasOwnProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasOwnProperty", name}).(bool)
}

func (ptr *QJSValue) HasProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasProperty", name}).(bool)
}

func (ptr *QJSValue) IsArray() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsArray"}).(bool)
}

func (ptr *QJSValue) IsBool() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBool"}).(bool)
}

func (ptr *QJSValue) IsCallable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCallable"}).(bool)
}

func (ptr *QJSValue) IsDate() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDate"}).(bool)
}

func (ptr *QJSValue) IsError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsError"}).(bool)
}

func (ptr *QJSValue) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QJSValue) IsNumber() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNumber"}).(bool)
}

func (ptr *QJSValue) IsObject() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsObject"}).(bool)
}

func (ptr *QJSValue) IsQMetaObject() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsQMetaObject"}).(bool)
}

func (ptr *QJSValue) IsQObject() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsQObject"}).(bool)
}

func (ptr *QJSValue) IsRegExp() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRegExp"}).(bool)
}

func (ptr *QJSValue) IsString() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsString"}).(bool)
}

func (ptr *QJSValue) IsUndefined() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUndefined"}).(bool)
}

func (ptr *QJSValue) IsVariant() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVariant"}).(bool)
}

func (ptr *QJSValue) Property(name string) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", name}).(*QJSValue)
}

func (ptr *QJSValue) Property2(arrayIndex uint) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property2", arrayIndex}).(*QJSValue)
}

func (ptr *QJSValue) Prototype() *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prototype"}).(*QJSValue)
}

func (ptr *QJSValue) SetProperty(name string, value QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", name, value})
}

func (ptr *QJSValue) SetProperty2(arrayIndex uint, value QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty2", arrayIndex, value})
}

func (ptr *QJSValue) SetPrototype(prototype QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrototype", prototype})
}

func (ptr *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StrictlyEquals", other}).(bool)
}

func (ptr *QJSValue) ToBool() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToBool"}).(bool)
}

func (ptr *QJSValue) ToDateTime() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDateTime"}).(*core.QDateTime)
}

func (ptr *QJSValue) ToInt() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToInt"}).(float64))
}

func (ptr *QJSValue) ToNumber() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToNumber"}).(float64)
}

func (ptr *QJSValue) ToQMetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToQMetaObject"}).(*core.QMetaObject)
}

func (ptr *QJSValue) ToQObject() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToQObject"}).(*core.QObject)
}

func (ptr *QJSValue) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QJSValue) ToUInt() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToUInt"}).(float64))
}

func (ptr *QJSValue) ToVariant() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToVariant"}).(*core.QVariant)
}

func (ptr *QJSValue) DestroyQJSValue() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQJSValue"})
}

func (ptr *QJSValue) __call_args_atList(i int) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__call_args_atList", i}).(*QJSValue)
}

func (ptr *QJSValue) __call_args_setList(i QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__call_args_setList", i})
}

func (ptr *QJSValue) __call_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__call_args_newList"}).(unsafe.Pointer)
}

func (ptr *QJSValue) __callAsConstructor_args_atList(i int) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callAsConstructor_args_atList", i}).(*QJSValue)
}

func (ptr *QJSValue) __callAsConstructor_args_setList(i QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callAsConstructor_args_setList", i})
}

func (ptr *QJSValue) __callAsConstructor_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callAsConstructor_args_newList"}).(unsafe.Pointer)
}

func (ptr *QJSValue) __callWithInstance_args_atList(i int) *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithInstance_args_atList", i}).(*QJSValue)
}

func (ptr *QJSValue) __callWithInstance_args_setList(i QJSValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithInstance_args_setList", i})
}

func (ptr *QJSValue) __callWithInstance_args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__callWithInstance_args_newList"}).(unsafe.Pointer)
}

type QJSValueIterator struct {
	internal.Internal
}

type QJSValueIterator_ITF interface {
	QJSValueIterator_PTR() *QJSValueIterator
}

func (ptr *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator {
	return ptr
}

func (ptr *QJSValueIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QJSValueIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQJSValueIterator(ptr QJSValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValueIterator_PTR().Pointer()
	}
	return nil
}

func (n *QJSValueIterator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQJSValueIteratorFromPointer(ptr unsafe.Pointer) (n *QJSValueIterator) {
	n = new(QJSValueIterator)
	n.InitFromInternal(uintptr(ptr), "qml.QJSValueIterator")
	return
}
func NewQJSValueIterator(object QJSValue_ITF) *QJSValueIterator {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQJSValueIterator", "", object}).(*QJSValueIterator)
}

func (ptr *QJSValueIterator) HasNext() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasNext"}).(bool)
}

func (ptr *QJSValueIterator) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QJSValueIterator) Next() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Next"}).(bool)
}

func (ptr *QJSValueIterator) Value() *QJSValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*QJSValue)
}

func (ptr *QJSValueIterator) DestroyQJSValueIterator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQJSValueIterator"})
}

type QQmlAbstractUrlInterceptor struct {
	internal.Internal
}

type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

func (ptr *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlAbstractUrlInterceptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlAbstractUrlInterceptor(ptr QQmlAbstractUrlInterceptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlAbstractUrlInterceptor_PTR().Pointer()
	}
	return nil
}

func (n *QQmlAbstractUrlInterceptor) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) (n *QQmlAbstractUrlInterceptor) {
	n = new(QQmlAbstractUrlInterceptor)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlAbstractUrlInterceptor")
	return
}

//go:generate stringer -type=QQmlAbstractUrlInterceptor__DataType
//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

func NewQQmlAbstractUrlInterceptor() *QQmlAbstractUrlInterceptor {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlAbstractUrlInterceptor", ""}).(*QQmlAbstractUrlInterceptor)
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectIntercept(f func(url *core.QUrl, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIntercept", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectIntercept() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIntercept"})
}

func (ptr *QQmlAbstractUrlInterceptor) Intercept(url core.QUrl_ITF, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Intercept", url, ty}).(*core.QUrl)
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectDestroyQQmlAbstractUrlInterceptor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlAbstractUrlInterceptor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectDestroyQQmlAbstractUrlInterceptor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlAbstractUrlInterceptor"})
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlAbstractUrlInterceptor"})
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlAbstractUrlInterceptorDefault"})
}

type QQmlApplicationEngine struct {
	QQmlEngine
}

type QQmlApplicationEngine_ITF interface {
	QQmlEngine_ITF
	QQmlApplicationEngine_PTR() *QQmlApplicationEngine
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine {
	return ptr
}

func (ptr *QQmlApplicationEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlApplicationEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQmlEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlApplicationEngine(ptr QQmlApplicationEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlApplicationEngine_PTR().Pointer()
	}
	return nil
}

func (n *QQmlApplicationEngine) InitFromInternal(ptr uintptr, name string) {
	n.QQmlEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlApplicationEngine) ClassNameInternalF() string {
	return n.QQmlEngine_PTR().ClassNameInternalF()
}

func NewQQmlApplicationEngineFromPointer(ptr unsafe.Pointer) (n *QQmlApplicationEngine) {
	n = new(QQmlApplicationEngine)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlApplicationEngine")
	return
}
func NewQQmlApplicationEngine(parent core.QObject_ITF) *QQmlApplicationEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlApplicationEngine", "", parent}).(*QQmlApplicationEngine)
}

func NewQQmlApplicationEngine2(url core.QUrl_ITF, parent core.QObject_ITF) *QQmlApplicationEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlApplicationEngine2", "", url, parent}).(*QQmlApplicationEngine)
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObject_ITF) *QQmlApplicationEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlApplicationEngine3", "", filePath, parent}).(*QQmlApplicationEngine)
}

func (ptr *QQmlApplicationEngine) ConnectLoad(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlApplicationEngine) DisconnectLoad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad"})
}

func (ptr *QQmlApplicationEngine) Load(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", url})
}

func (ptr *QQmlApplicationEngine) LoadDefault(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadDefault", url})
}

func (ptr *QQmlApplicationEngine) ConnectLoad2(f func(filePath string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlApplicationEngine) DisconnectLoad2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad2"})
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2", filePath})
}

func (ptr *QQmlApplicationEngine) Load2Default(filePath string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2Default", filePath})
}

func (ptr *QQmlApplicationEngine) ConnectLoadData(f func(data *core.QByteArray, url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlApplicationEngine) DisconnectLoadData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadData"})
}

func (ptr *QQmlApplicationEngine) LoadData(data core.QByteArray_ITF, url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadData", data, url})
}

func (ptr *QQmlApplicationEngine) LoadDataDefault(data core.QByteArray_ITF, url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadDataDefault", data, url})
}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object *core.QObject, url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectObjectCreated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectObjectCreated"})
}

func (ptr *QQmlApplicationEngine) ObjectCreated(object core.QObject_ITF, url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ObjectCreated", object, url})
}

func (ptr *QQmlApplicationEngine) RootObjects() []*core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootObjects"}).([]*core.QObject)
}

func (ptr *QQmlApplicationEngine) ConnectDestroyQQmlApplicationEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlApplicationEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlApplicationEngine) DisconnectDestroyQQmlApplicationEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlApplicationEngine"})
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlApplicationEngine"})
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlApplicationEngineDefault"})
}

func (ptr *QQmlApplicationEngine) __rootObjects_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rootObjects_atList", i}).(*core.QObject)
}

func (ptr *QQmlApplicationEngine) __rootObjects_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rootObjects_setList", i})
}

func (ptr *QQmlApplicationEngine) __rootObjects_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rootObjects_newList"}).(unsafe.Pointer)
}

type QQmlComponent struct {
	core.QObject
}

type QQmlComponent_ITF interface {
	core.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func (ptr *QQmlComponent) QQmlComponent_PTR() *QQmlComponent {
	return ptr
}

func (ptr *QQmlComponent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlComponent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlComponent(ptr QQmlComponent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlComponent_PTR().Pointer()
	}
	return nil
}

func (n *QQmlComponent) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlComponent) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlComponentFromPointer(ptr unsafe.Pointer) (n *QQmlComponent) {
	n = new(QQmlComponent)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlComponent")
	return
}

//go:generate stringer -type=QQmlComponent__CompilationMode
//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(1)
)

//go:generate stringer -type=QQmlComponent__Status
//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    QQmlComponent__Status = QQmlComponent__Status(0)
	QQmlComponent__Ready   QQmlComponent__Status = QQmlComponent__Status(1)
	QQmlComponent__Loading QQmlComponent__Status = QQmlComponent__Status(2)
	QQmlComponent__Error   QQmlComponent__Status = QQmlComponent__Status(3)
)

func NewQQmlComponent2(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlComponent2", "", engine, parent}).(*QQmlComponent)
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlComponent3", "", engine, fileName, parent}).(*QQmlComponent)
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlComponent4", "", engine, fileName, mode, parent}).(*QQmlComponent)
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlComponent5", "", engine, url, parent}).(*QQmlComponent)
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlComponent6", "", engine, url, mode, parent}).(*QQmlComponent)
}

func (ptr *QQmlComponent) ConnectBeginCreate(f func(publicContext *QQmlContext) *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeginCreate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectBeginCreate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeginCreate"})
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginCreate", publicContext}).(*core.QObject)
}

func (ptr *QQmlComponent) BeginCreateDefault(publicContext QQmlContext_ITF) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginCreateDefault", publicContext}).(*core.QObject)
}

func (ptr *QQmlComponent) ConnectCompleteCreate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCompleteCreate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectCompleteCreate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCompleteCreate"})
}

func (ptr *QQmlComponent) CompleteCreate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompleteCreate"})
}

func (ptr *QQmlComponent) CompleteCreateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompleteCreateDefault"})
}

func (ptr *QQmlComponent) ConnectCreate(f func(context *QQmlContext) *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectCreate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreate"})
}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Create", context}).(*core.QObject)
}

func (ptr *QQmlComponent) CreateDefault(context QQmlContext_ITF) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDefault", context}).(*core.QObject)
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubator_ITF, context QQmlContext_ITF, forContext QQmlContext_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Create2", incubator, context, forContext})
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreationContext"}).(*QQmlContext)
}

func (ptr *QQmlComponent) Engine() *QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QQmlEngine)
}

func (ptr *QQmlComponent) Errors() []*QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Errors"}).([]*QQmlError)
}

func (ptr *QQmlComponent) IsError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsError"}).(bool)
}

func (ptr *QQmlComponent) IsLoading() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLoading"}).(bool)
}

func (ptr *QQmlComponent) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QQmlComponent) IsReady() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReady"}).(bool)
}

func (ptr *QQmlComponent) ConnectLoadUrl(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectLoadUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadUrl"})
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadUrl", url})
}

func (ptr *QQmlComponent) LoadUrlDefault(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadUrlDefault", url})
}

func (ptr *QQmlComponent) ConnectLoadUrl2(f func(url *core.QUrl, mode QQmlComponent__CompilationMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadUrl2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectLoadUrl2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadUrl2"})
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadUrl2", url, mode})
}

func (ptr *QQmlComponent) LoadUrl2Default(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadUrl2Default", url, mode})
}

func (ptr *QQmlComponent) Progress() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Progress"}).(float64)
}

func (ptr *QQmlComponent) ConnectProgressChanged(f func(progress float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProgressChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectProgressChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProgressChanged"})
}

func (ptr *QQmlComponent) ProgressChanged(progress float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProgressChanged", progress})
}

func (ptr *QQmlComponent) ConnectSetData(f func(data *core.QByteArray, url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectSetData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetData"})
}

func (ptr *QQmlComponent) SetData(data core.QByteArray_ITF, url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", data, url})
}

func (ptr *QQmlComponent) SetDataDefault(data core.QByteArray_ITF, url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", data, url})
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {

	return QQmlComponent__Status(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Status"}).(float64))
}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStatusChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStatusChanged"})
}

func (ptr *QQmlComponent) StatusChanged(status QQmlComponent__Status) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StatusChanged", status})
}

func (ptr *QQmlComponent) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QQmlComponent) ConnectDestroyQQmlComponent(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlComponent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlComponent) DisconnectDestroyQQmlComponent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlComponent"})
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlComponent"})
}

func (ptr *QQmlComponent) DestroyQQmlComponentDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlComponentDefault"})
}

func (ptr *QQmlComponent) __errors_atList(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_atList", i}).(*QQmlError)
}

func (ptr *QQmlComponent) __errors_setList(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_setList", i})
}

func (ptr *QQmlComponent) __errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlComponent) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlComponent) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlComponent) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlComponent) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlComponent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlComponent) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlComponent) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlComponent) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlComponent) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlComponent) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlComponent) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlComponent) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlComponent) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlComponent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlComponent) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlComponent) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlComponent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlComponent) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlComponent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlComponent) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlComponent) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlContext struct {
	core.QObject
}

type QQmlContext_ITF interface {
	core.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return ptr
}

func (ptr *QQmlContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlContext(ptr QQmlContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContext_PTR().Pointer()
	}
	return nil
}

func (n *QQmlContext) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlContext) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlContextFromPointer(ptr unsafe.Pointer) (n *QQmlContext) {
	n = new(QQmlContext)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlContext")
	return
}
func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlContext", "", engine, parent}).(*QQmlContext)
}

func NewQQmlContext2(parentContext QQmlContext_ITF, parent core.QObject_ITF) *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlContext2", "", parentContext, parent}).(*QQmlContext)
}

func (ptr *QQmlContext) BaseUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseUrl"}).(*core.QUrl)
}

func (ptr *QQmlContext) ContextObject() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextObject"}).(*core.QObject)
}

func (ptr *QQmlContext) ContextProperty(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextProperty", name}).(*core.QVariant)
}

func (ptr *QQmlContext) Engine() *QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QQmlEngine)
}

func (ptr *QQmlContext) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameForObject", object}).(string)
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentContext"}).(*QQmlContext)
}

func (ptr *QQmlContext) ResolvedUrl(src core.QUrl_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResolvedUrl", src}).(*core.QUrl)
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaseUrl", baseUrl})
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContextObject", object})
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContextProperty", name, value})
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContextProperty2", name, value})
}

func (ptr *QQmlContext) ConnectDestroyQQmlContext(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlContext", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlContext) DisconnectDestroyQQmlContext() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlContext"})
}

func (ptr *QQmlContext) DestroyQQmlContext() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlContext"})
}

func (ptr *QQmlContext) DestroyQQmlContextDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlContextDefault"})
}

func (ptr *QQmlContext) __setContextProperties_properties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setContextProperties_properties_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlContext) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlContext) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlContext) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlContext) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlContext) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlContext) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlContext) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlContext) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlContext) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlContext) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlContext) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlContext) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlContext) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlContext) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlContext) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlContext) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlContext) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlContext) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlContext) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlContext) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlContext) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlDebuggingEnabler struct {
	internal.Internal
}

type QQmlDebuggingEnabler_ITF interface {
	QQmlDebuggingEnabler_PTR() *QQmlDebuggingEnabler
}

func (ptr *QQmlDebuggingEnabler) QQmlDebuggingEnabler_PTR() *QQmlDebuggingEnabler {
	return ptr
}

func (ptr *QQmlDebuggingEnabler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlDebuggingEnabler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlDebuggingEnabler(ptr QQmlDebuggingEnabler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlDebuggingEnabler_PTR().Pointer()
	}
	return nil
}

func (n *QQmlDebuggingEnabler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlDebuggingEnablerFromPointer(ptr unsafe.Pointer) (n *QQmlDebuggingEnabler) {
	n = new(QQmlDebuggingEnabler)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlDebuggingEnabler")
	return
}

func (ptr *QQmlDebuggingEnabler) DestroyQQmlDebuggingEnabler() {
}

//go:generate stringer -type=QQmlDebuggingEnabler__StartMode
//QQmlDebuggingEnabler::StartMode
type QQmlDebuggingEnabler__StartMode int64

const (
	QQmlDebuggingEnabler__DoNotWaitForClient QQmlDebuggingEnabler__StartMode = QQmlDebuggingEnabler__StartMode(0)
	QQmlDebuggingEnabler__WaitForClient      QQmlDebuggingEnabler__StartMode = QQmlDebuggingEnabler__StartMode(1)
)

func QQmlDebuggingEnabler_ConnectToLocalDebugger(socketFileName string, mode QQmlDebuggingEnabler__StartMode) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_ConnectToLocalDebugger", "", socketFileName, mode}).(bool)
}

func (ptr *QQmlDebuggingEnabler) ConnectToLocalDebugger(socketFileName string, mode QQmlDebuggingEnabler__StartMode) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_ConnectToLocalDebugger", "", socketFileName, mode}).(bool)
}

func QQmlDebuggingEnabler_DebuggerServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_DebuggerServices", ""}).([]string)
}

func (ptr *QQmlDebuggingEnabler) DebuggerServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_DebuggerServices", ""}).([]string)
}

func QQmlDebuggingEnabler_InspectorServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_InspectorServices", ""}).([]string)
}

func (ptr *QQmlDebuggingEnabler) InspectorServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_InspectorServices", ""}).([]string)
}

func QQmlDebuggingEnabler_NativeDebuggerServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_NativeDebuggerServices", ""}).([]string)
}

func (ptr *QQmlDebuggingEnabler) NativeDebuggerServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_NativeDebuggerServices", ""}).([]string)
}

func QQmlDebuggingEnabler_ProfilerServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_ProfilerServices", ""}).([]string)
}

func (ptr *QQmlDebuggingEnabler) ProfilerServices() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_ProfilerServices", ""}).([]string)
}

func QQmlDebuggingEnabler_SetServices(services []string) {

	internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_SetServices", "", services})
}

func (ptr *QQmlDebuggingEnabler) SetServices(services []string) {

	internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_SetServices", "", services})
}

func QQmlDebuggingEnabler_StartDebugConnector(pluginName string, configuration map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_StartDebugConnector", "", pluginName, configuration}).(bool)
}

func (ptr *QQmlDebuggingEnabler) StartDebugConnector(pluginName string, configuration map[string]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_StartDebugConnector", "", pluginName, configuration}).(bool)
}

func QQmlDebuggingEnabler_StartTcpDebugServer(port int, mode QQmlDebuggingEnabler__StartMode, hostName string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_StartTcpDebugServer", "", port, mode, hostName}).(bool)
}

func (ptr *QQmlDebuggingEnabler) StartTcpDebugServer(port int, mode QQmlDebuggingEnabler__StartMode, hostName string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlDebuggingEnabler_StartTcpDebugServer", "", port, mode, hostName}).(bool)
}

func (ptr *QQmlDebuggingEnabler) __startDebugConnector_configuration_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__startDebugConnector_configuration_atList", v, i}).(*core.QVariant)
}

func (ptr *QQmlDebuggingEnabler) __startDebugConnector_configuration_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__startDebugConnector_configuration_setList", key, i})
}

func (ptr *QQmlDebuggingEnabler) __startDebugConnector_configuration_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__startDebugConnector_configuration_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlDebuggingEnabler) __startDebugConnector_configuration_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__startDebugConnector_configuration_keyList"}).([]string)
}

func (ptr *QQmlDebuggingEnabler) ____startDebugConnector_configuration_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____startDebugConnector_configuration_keyList_atList", i}).(string)
}

func (ptr *QQmlDebuggingEnabler) ____startDebugConnector_configuration_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____startDebugConnector_configuration_keyList_setList", i})
}

func (ptr *QQmlDebuggingEnabler) ____startDebugConnector_configuration_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____startDebugConnector_configuration_keyList_newList"}).(unsafe.Pointer)
}

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

func (ptr *QQmlEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QJSEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func (n *QQmlEngine) InitFromInternal(ptr uintptr, name string) {
	n.QJSEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlEngine) ClassNameInternalF() string {
	return n.QJSEngine_PTR().ClassNameInternalF()
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) (n *QQmlEngine) {
	n = new(QQmlEngine)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlEngine")
	return
}

//go:generate stringer -type=QQmlEngine__ObjectOwnership
//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(1)
)

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlEngine", "", parent}).(*QQmlEngine)
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddImageProvider", providerId, provider})
}

func (ptr *QQmlEngine) AddImportPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddImportPath", path})
}

func (ptr *QQmlEngine) AddPluginPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddPluginPath", path})
}

func (ptr *QQmlEngine) BaseUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseUrl"}).(*core.QUrl)
}

func (ptr *QQmlEngine) ClearComponentCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearComponentCache"})
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_ContextForObject", "", object}).(*QQmlContext)
}

func (ptr *QQmlEngine) ContextForObject(object core.QObject_ITF) *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_ContextForObject", "", object}).(*QQmlContext)
}

func (ptr *QQmlEngine) ConnectExit(f func(retCode int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExit", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlEngine) DisconnectExit() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExit"})
}

func (ptr *QQmlEngine) Exit(retCode int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exit", retCode})
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImageProvider", providerId}).(*QQmlImageProviderBase)
}

func (ptr *QQmlEngine) ImportPathList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImportPathList"}).([]string)
}

func (ptr *QQmlEngine) ImportPlugin(filePath string, uri string, errors []*QQmlError) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImportPlugin", filePath, uri, errors}).(bool)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubationController"}).(*QQmlIncubationController)
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessManager"}).(*network.QNetworkAccessManager)
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessManagerFactory"}).(*QQmlNetworkAccessManagerFactory)
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {

	return QQmlEngine__ObjectOwnership(internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_ObjectOwnership", "", object}).(float64))
}

func (ptr *QQmlEngine) ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {

	return QQmlEngine__ObjectOwnership(internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_ObjectOwnership", "", object}).(float64))
}

func (ptr *QQmlEngine) OfflineStorageDatabaseFilePath(databaseName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OfflineStorageDatabaseFilePath", databaseName}).(string)
}

func (ptr *QQmlEngine) OfflineStoragePath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OfflineStoragePath"}).(string)
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OutputWarningsToStandardError"}).(bool)
}

func (ptr *QQmlEngine) PluginPathList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PluginPathList"}).([]string)
}

func (ptr *QQmlEngine) ConnectQuit(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectQuit", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlEngine) DisconnectQuit() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectQuit"})
}

func (ptr *QQmlEngine) Quit() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Quit"})
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveImageProvider", providerId})
}

func (ptr *QQmlEngine) ConnectRetranslate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRetranslate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlEngine) DisconnectRetranslate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRetranslate"})
}

func (ptr *QQmlEngine) Retranslate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Retranslate"})
}

func (ptr *QQmlEngine) RetranslateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RetranslateDefault"})
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootContext"}).(*QQmlContext)
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaseUrl", url})
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_SetContextForObject", "", object, context})
}

func (ptr *QQmlEngine) SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_SetContextForObject", "", object, context})
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetImportPathList", paths})
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIncubationController", controller})
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNetworkAccessManagerFactory", factory})
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {

	internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_SetObjectOwnership", "", object, ownership})
}

func (ptr *QQmlEngine) SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {

	internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_SetObjectOwnership", "", object, ownership})
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOfflineStoragePath", dir})
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOutputWarningsToStandardError", enabled})
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPluginPathList", paths})
}

func (ptr *QQmlEngine) TrimComponentCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TrimComponentCache"})
}

func (ptr *QQmlEngine) ConnectWarnings(f func(warnings []*QQmlError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWarnings", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlEngine) DisconnectWarnings() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWarnings"})
}

func (ptr *QQmlEngine) Warnings(warnings []*QQmlError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Warnings", warnings})
}

func (ptr *QQmlEngine) ConnectDestroyQQmlEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlEngine) DisconnectDestroyQQmlEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlEngine"})
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlEngine"})
}

func (ptr *QQmlEngine) DestroyQQmlEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlEngineDefault"})
}

func QQmlEngine_QmlRegisterSingletonType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_QmlRegisterSingletonType", "", url, uri, versionMajor, versionMinor, qmlName}).(float64))
}

func (ptr *QQmlEngine) QmlRegisterSingletonType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_QmlRegisterSingletonType", "", url, uri, versionMajor, versionMinor, qmlName}).(float64))
}

func QQmlEngine_QmlRegisterType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_QmlRegisterType", "", url, uri, versionMajor, versionMinor, qmlName}).(float64))
}

func (ptr *QQmlEngine) QmlRegisterType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlEngine_QmlRegisterType", "", url, uri, versionMajor, versionMinor, qmlName}).(float64))
}

func (ptr *QQmlEngine) __importPlugin_errors_atList(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__importPlugin_errors_atList", i}).(*QQmlError)
}

func (ptr *QQmlEngine) __importPlugin_errors_setList(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__importPlugin_errors_setList", i})
}

func (ptr *QQmlEngine) __importPlugin_errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__importPlugin_errors_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlEngine) __qmlDebug_errors_atList3(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlDebug_errors_atList3", i}).(*QQmlError)
}

func (ptr *QQmlEngine) __qmlDebug_errors_setList3(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlDebug_errors_setList3", i})
}

func (ptr *QQmlEngine) __qmlDebug_errors_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlDebug_errors_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlEngine) __qmlInfo_errors_atList3(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlInfo_errors_atList3", i}).(*QQmlError)
}

func (ptr *QQmlEngine) __qmlInfo_errors_setList3(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlInfo_errors_setList3", i})
}

func (ptr *QQmlEngine) __qmlInfo_errors_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlInfo_errors_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlEngine) __qmlWarning_errors_atList3(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlWarning_errors_atList3", i}).(*QQmlError)
}

func (ptr *QQmlEngine) __qmlWarning_errors_setList3(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlWarning_errors_setList3", i})
}

func (ptr *QQmlEngine) __qmlWarning_errors_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__qmlWarning_errors_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlEngine) __warnings_warnings_atList(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__warnings_warnings_atList", i}).(*QQmlError)
}

func (ptr *QQmlEngine) __warnings_warnings_setList(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__warnings_warnings_setList", i})
}

func (ptr *QQmlEngine) __warnings_warnings_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__warnings_warnings_newList"}).(unsafe.Pointer)
}

type QQmlError struct {
	internal.Internal
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func (ptr *QQmlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func (n *QQmlError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) (n *QQmlError) {
	n = new(QQmlError)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlError")
	return
}

func (ptr *QQmlError) DestroyQQmlError() {
}

func NewQQmlError() *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlError", ""}).(*QQmlError)
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlError2", "", other}).(*QQmlError)
}

func (ptr *QQmlError) Column() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Column"}).(float64))
}

func (ptr *QQmlError) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QQmlError) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QQmlError) Line() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Line"}).(float64))
}

func (ptr *QQmlError) Object() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Object"}).(*core.QObject)
}

func (ptr *QQmlError) SetColumn(column int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumn", column})
}

func (ptr *QQmlError) SetDescription(description string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDescription", description})
}

func (ptr *QQmlError) SetLine(line int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLine", line})
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetObject", object})
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QQmlError) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QQmlError) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

type QQmlExpression struct {
	core.QObject
}

type QQmlExpression_ITF interface {
	core.QObject_ITF
	QQmlExpression_PTR() *QQmlExpression
}

func (ptr *QQmlExpression) QQmlExpression_PTR() *QQmlExpression {
	return ptr
}

func (ptr *QQmlExpression) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlExpression) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlExpression(ptr QQmlExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExpression_PTR().Pointer()
	}
	return nil
}

func (n *QQmlExpression) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlExpression) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlExpressionFromPointer(ptr unsafe.Pointer) (n *QQmlExpression) {
	n = new(QQmlExpression)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlExpression")
	return
}
func NewQQmlExpression() *QQmlExpression {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlExpression", ""}).(*QQmlExpression)
}

func NewQQmlExpression2(ctxt QQmlContext_ITF, scope core.QObject_ITF, expression string, parent core.QObject_ITF) *QQmlExpression {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlExpression2", "", ctxt, scope, expression, parent}).(*QQmlExpression)
}

func NewQQmlExpression3(scri QQmlScriptString_ITF, ctxt QQmlContext_ITF, scope core.QObject_ITF, parent core.QObject_ITF) *QQmlExpression {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlExpression3", "", scri, ctxt, scope, parent}).(*QQmlExpression)
}

func (ptr *QQmlExpression) ClearError() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearError"})
}

func (ptr *QQmlExpression) ColumnNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnNumber"}).(float64))
}

func (ptr *QQmlExpression) Context() *QQmlContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Context"}).(*QQmlContext)
}

func (ptr *QQmlExpression) Engine() *QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QQmlEngine)
}

func (ptr *QQmlExpression) Error() *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(*QQmlError)
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined *bool) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Evaluate", valueIsUndefined}).(*core.QVariant)
}

func (ptr *QQmlExpression) Expression() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Expression"}).(string)
}

func (ptr *QQmlExpression) HasError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasError"}).(bool)
}

func (ptr *QQmlExpression) LineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineNumber"}).(float64))
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NotifyOnValueChanged"}).(bool)
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScopeObject"}).(*core.QObject)
}

func (ptr *QQmlExpression) SetExpression(expression string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExpression", expression})
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNotifyOnValueChanged", notifyOnChange})
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceLocation", url, line, column})
}

func (ptr *QQmlExpression) SourceFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceFile"}).(string)
}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlExpression) DisconnectValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueChanged"})
}

func (ptr *QQmlExpression) ValueChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueChanged"})
}

func (ptr *QQmlExpression) ConnectDestroyQQmlExpression(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlExpression", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlExpression) DisconnectDestroyQQmlExpression() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlExpression"})
}

func (ptr *QQmlExpression) DestroyQQmlExpression() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlExpression"})
}

func (ptr *QQmlExpression) DestroyQQmlExpressionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlExpressionDefault"})
}

func (ptr *QQmlExpression) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlExpression) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlExpression) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlExpression) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlExpression) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlExpression) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlExpression) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlExpression) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlExpression) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlExpression) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlExpression) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlExpression) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlExpression) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlExpression) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlExpression) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlExpression) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlExpression) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlExpression) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlExpression) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlExpression) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlExpression) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlExtensionPlugin struct {
	core.QObject
}

type QQmlExtensionPlugin_ITF interface {
	core.QObject_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func (n *QQmlExtensionPlugin) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlExtensionPlugin) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) (n *QQmlExtensionPlugin) {
	n = new(QQmlExtensionPlugin)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlExtensionPlugin")
	return
}

func (ptr *QQmlExtensionPlugin) DestroyQQmlExtensionPlugin() {
}

func NewQQmlExtensionPlugin(parent core.QObject_ITF) *QQmlExtensionPlugin {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlExtensionPlugin", "", parent}).(*QQmlExtensionPlugin)
}

func (ptr *QQmlExtensionPlugin) BaseUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseUrl"}).(*core.QUrl)
}

func (ptr *QQmlExtensionPlugin) ConnectInitializeEngine(f func(engine *QQmlEngine, uri string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInitializeEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlExtensionPlugin) DisconnectInitializeEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInitializeEngine"})
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF, uri string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitializeEngine", engine, uri})
}

func (ptr *QQmlExtensionPlugin) InitializeEngineDefault(engine QQmlEngine_ITF, uri string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitializeEngineDefault", engine, uri})
}

func (ptr *QQmlExtensionPlugin) ConnectRegisterTypes(f func(uri string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRegisterTypes", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlExtensionPlugin) DisconnectRegisterTypes() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRegisterTypes"})
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterTypes", uri})
}

func (ptr *QQmlExtensionPlugin) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlExtensionPlugin) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlExtensionPlugin) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlExtensionPlugin) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlExtensionPlugin) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelector_ITF interface {
	core.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return ptr
}

func (ptr *QQmlFileSelector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlFileSelector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelector_PTR().Pointer()
	}
	return nil
}

func (n *QQmlFileSelector) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlFileSelector) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlFileSelectorFromPointer(ptr unsafe.Pointer) (n *QQmlFileSelector) {
	n = new(QQmlFileSelector)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlFileSelector")
	return
}
func NewQQmlFileSelector(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlFileSelector {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlFileSelector", "", engine, parent}).(*QQmlFileSelector)
}

func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlFileSelector_Get", "", engine}).(*QQmlFileSelector)
}

func (ptr *QQmlFileSelector) Get(engine QQmlEngine_ITF) *QQmlFileSelector {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlFileSelector_Get", "", engine}).(*QQmlFileSelector)
}

func (ptr *QQmlFileSelector) Selector() *core.QFileSelector {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Selector"}).(*core.QFileSelector)
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExtraSelectors", strin})
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExtraSelectors2", strin})
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelector", selector})
}

func (ptr *QQmlFileSelector) ConnectDestroyQQmlFileSelector(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlFileSelector", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlFileSelector) DisconnectDestroyQQmlFileSelector() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlFileSelector"})
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlFileSelector"})
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelectorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlFileSelectorDefault"})
}

func (ptr *QQmlFileSelector) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlFileSelector) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlFileSelector) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlFileSelector) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlFileSelector) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlFileSelector) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlFileSelector) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlFileSelector) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlFileSelector) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlFileSelector) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlFileSelector) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlFileSelector) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlFileSelector) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlFileSelector) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlFileSelector) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlFileSelector) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlFileSelector) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlFileSelector) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlImageProviderBase struct {
	internal.Internal
}

type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase {
	return ptr
}

func (ptr *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlImageProviderBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlImageProviderBase(ptr QQmlImageProviderBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func (n *QQmlImageProviderBase) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) (n *QQmlImageProviderBase) {
	n = new(QQmlImageProviderBase)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlImageProviderBase")
	return
}

func (ptr *QQmlImageProviderBase) DestroyQQmlImageProviderBase() {
}

//go:generate stringer -type=QQmlImageProviderBase__ImageType
//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image         QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap        QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(3)
	QQmlImageProviderBase__ImageResponse QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(4)
)

//go:generate stringer -type=QQmlImageProviderBase__Flag
//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading QQmlImageProviderBase__Flag = QQmlImageProviderBase__Flag(0x01)
)

func (ptr *QQmlImageProviderBase) ConnectFlags(f func() QQmlImageProviderBase__Flag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlags", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlImageProviderBase) DisconnectFlags() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlags"})
}

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {

	return QQmlImageProviderBase__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QQmlImageProviderBase) ConnectImageType(f func() QQmlImageProviderBase__ImageType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectImageType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlImageProviderBase) DisconnectImageType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectImageType"})
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {

	return QQmlImageProviderBase__ImageType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImageType"}).(float64))
}

type QQmlIncubationController struct {
	internal.Internal
}

type QQmlIncubationController_ITF interface {
	QQmlIncubationController_PTR() *QQmlIncubationController
}

func (ptr *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return ptr
}

func (ptr *QQmlIncubationController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlIncubationController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlIncubationController(ptr QQmlIncubationController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubationController_PTR().Pointer()
	}
	return nil
}

func (n *QQmlIncubationController) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlIncubationControllerFromPointer(ptr unsafe.Pointer) (n *QQmlIncubationController) {
	n = new(QQmlIncubationController)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlIncubationController")
	return
}

func (ptr *QQmlIncubationController) DestroyQQmlIncubationController() {
}

func NewQQmlIncubationController2() *QQmlIncubationController {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlIncubationController2", ""}).(*QQmlIncubationController)
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QQmlEngine)
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubateFor", msecs})
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubatingObjectCount"}).(float64))
}

func (ptr *QQmlIncubationController) ConnectIncubatingObjectCountChanged(f func(incubatingObjectCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIncubatingObjectCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlIncubationController) DisconnectIncubatingObjectCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIncubatingObjectCountChanged"})
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChanged(incubatingObjectCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubatingObjectCountChanged", incubatingObjectCount})
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChangedDefault(incubatingObjectCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubatingObjectCountChangedDefault", incubatingObjectCount})
}

type QQmlIncubator struct {
	internal.Internal
}

type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (ptr *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator {
	return ptr
}

func (ptr *QQmlIncubator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlIncubator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlIncubator(ptr QQmlIncubator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubator_PTR().Pointer()
	}
	return nil
}

func (n *QQmlIncubator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlIncubatorFromPointer(ptr unsafe.Pointer) (n *QQmlIncubator) {
	n = new(QQmlIncubator)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlIncubator")
	return
}

func (ptr *QQmlIncubator) DestroyQQmlIncubator() {
}

//go:generate stringer -type=QQmlIncubator__IncubationMode
//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(2)
)

//go:generate stringer -type=QQmlIncubator__Status
//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    QQmlIncubator__Status = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   QQmlIncubator__Status = QQmlIncubator__Status(1)
	QQmlIncubator__Loading QQmlIncubator__Status = QQmlIncubator__Status(2)
	QQmlIncubator__Error   QQmlIncubator__Status = QQmlIncubator__Status(3)
)

func NewQQmlIncubator2(mode QQmlIncubator__IncubationMode) *QQmlIncubator {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlIncubator2", "", mode}).(*QQmlIncubator)
}

func (ptr *QQmlIncubator) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QQmlIncubator) Errors() []*QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Errors"}).([]*QQmlError)
}

func (ptr *QQmlIncubator) ForceCompletion() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ForceCompletion"})
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {

	return QQmlIncubator__IncubationMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncubationMode"}).(float64))
}

func (ptr *QQmlIncubator) IsError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsError"}).(bool)
}

func (ptr *QQmlIncubator) IsLoading() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLoading"}).(bool)
}

func (ptr *QQmlIncubator) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QQmlIncubator) IsReady() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReady"}).(bool)
}

func (ptr *QQmlIncubator) Object() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Object"}).(*core.QObject)
}

func (ptr *QQmlIncubator) ConnectSetInitialState(f func(object *core.QObject)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetInitialState", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlIncubator) DisconnectSetInitialState() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetInitialState"})
}

func (ptr *QQmlIncubator) SetInitialState(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInitialState", object})
}

func (ptr *QQmlIncubator) SetInitialStateDefault(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInitialStateDefault", object})
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {

	return QQmlIncubator__Status(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Status"}).(float64))
}

func (ptr *QQmlIncubator) ConnectStatusChanged(f func(status QQmlIncubator__Status)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStatusChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlIncubator) DisconnectStatusChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStatusChanged"})
}

func (ptr *QQmlIncubator) StatusChanged(status QQmlIncubator__Status) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StatusChanged", status})
}

func (ptr *QQmlIncubator) StatusChangedDefault(status QQmlIncubator__Status) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StatusChangedDefault", status})
}

func (ptr *QQmlIncubator) __errors_atList(i int) *QQmlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_atList", i}).(*QQmlError)
}

func (ptr *QQmlIncubator) __errors_setList(i QQmlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_setList", i})
}

func (ptr *QQmlIncubator) __errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__errors_newList"}).(unsafe.Pointer)
}

type QQmlInfo struct {
	core.QDebug
}

type QQmlInfo_ITF interface {
	core.QDebug_ITF
	QQmlInfo_PTR() *QQmlInfo
}

func (ptr *QQmlInfo) QQmlInfo_PTR() *QQmlInfo {
	return ptr
}

func (ptr *QQmlInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebug_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDebug_PTR().SetPointer(p)
	}
}

func PointerFromQQmlInfo(ptr QQmlInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlInfo_PTR().Pointer()
	}
	return nil
}

func (n *QQmlInfo) InitFromInternal(ptr uintptr, name string) {
	n.QDebug_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlInfo) ClassNameInternalF() string {
	return n.QDebug_PTR().ClassNameInternalF()
}

func NewQQmlInfoFromPointer(ptr unsafe.Pointer) (n *QQmlInfo) {
	n = new(QQmlInfo)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlInfo")
	return
}

func (ptr *QQmlInfo) DestroyQQmlInfo() {
}

type QQmlListProperty struct {
	internal.Internal
}

type QQmlListProperty_ITF interface {
	QQmlListProperty_PTR() *QQmlListProperty
}

func (ptr *QQmlListProperty) QQmlListProperty_PTR() *QQmlListProperty {
	return ptr
}

func (ptr *QQmlListProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlListProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlListProperty(ptr QQmlListProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListProperty_PTR().Pointer()
	}
	return nil
}

func (n *QQmlListProperty) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlListPropertyFromPointer(ptr unsafe.Pointer) (n *QQmlListProperty) {
	n = new(QQmlListProperty)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlListProperty")
	return
}

func (ptr *QQmlListProperty) DestroyQQmlListProperty() {
}

type QQmlListReference struct {
	internal.Internal
}

type QQmlListReference_ITF interface {
	QQmlListReference_PTR() *QQmlListReference
}

func (ptr *QQmlListReference) QQmlListReference_PTR() *QQmlListReference {
	return ptr
}

func (ptr *QQmlListReference) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlListReference) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlListReference(ptr QQmlListReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListReference_PTR().Pointer()
	}
	return nil
}

func (n *QQmlListReference) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlListReferenceFromPointer(ptr unsafe.Pointer) (n *QQmlListReference) {
	n = new(QQmlListReference)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlListReference")
	return
}

func (ptr *QQmlListReference) DestroyQQmlListReference() {
}

func NewQQmlListReference() *QQmlListReference {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlListReference", ""}).(*QQmlListReference)
}

func NewQQmlListReference2(object core.QObject_ITF, property string, engine QQmlEngine_ITF) *QQmlListReference {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlListReference2", "", object, property, engine}).(*QQmlListReference)
}

func (ptr *QQmlListReference) Append(object core.QObject_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", object}).(bool)
}

func (ptr *QQmlListReference) At(index int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At", index}).(*core.QObject)
}

func (ptr *QQmlListReference) CanAppend() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanAppend"}).(bool)
}

func (ptr *QQmlListReference) CanAt() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanAt"}).(bool)
}

func (ptr *QQmlListReference) CanClear() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanClear"}).(bool)
}

func (ptr *QQmlListReference) CanCount() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanCount"}).(bool)
}

func (ptr *QQmlListReference) Clear() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"}).(bool)
}

func (ptr *QQmlListReference) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QQmlListReference) IsManipulable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsManipulable"}).(bool)
}

func (ptr *QQmlListReference) IsReadable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReadable"}).(bool)
}

func (ptr *QQmlListReference) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ListElementType"}).(*core.QMetaObject)
}

func (ptr *QQmlListReference) Object() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Object"}).(*core.QObject)
}

type QQmlNetworkAccessManagerFactory struct {
	internal.Internal
}

type QQmlNetworkAccessManagerFactory_ITF interface {
	QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlNetworkAccessManagerFactory(ptr QQmlNetworkAccessManagerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNetworkAccessManagerFactory_PTR().Pointer()
	}
	return nil
}

func (n *QQmlNetworkAccessManagerFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlNetworkAccessManagerFactory) {
	n = new(QQmlNetworkAccessManagerFactory)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlNetworkAccessManagerFactory")
	return
}
func (ptr *QQmlNetworkAccessManagerFactory) ConnectCreate(f func(parent *core.QObject) *network.QNetworkAccessManager) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectCreate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreate"})
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Create", parent}).(*network.QNetworkAccessManager)
}

func (ptr *QQmlNetworkAccessManagerFactory) ConnectDestroyQQmlNetworkAccessManagerFactory(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlNetworkAccessManagerFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectDestroyQQmlNetworkAccessManagerFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlNetworkAccessManagerFactory"})
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlNetworkAccessManagerFactory"})
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactoryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlNetworkAccessManagerFactoryDefault"})
}

func NewQQmlNetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlNetworkAccessManagerFactory", ""}).(*QQmlNetworkAccessManagerFactory)
}

type QQmlParserStatus struct {
	internal.Internal
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlParserStatus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatus_PTR().Pointer()
	}
	return nil
}

func (n *QQmlParserStatus) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlParserStatusFromPointer(ptr unsafe.Pointer) (n *QQmlParserStatus) {
	n = new(QQmlParserStatus)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlParserStatus")
	return
}

func (ptr *QQmlParserStatus) DestroyQQmlParserStatus() {
}

func (ptr *QQmlParserStatus) ConnectClassBegin(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClassBegin", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlParserStatus) DisconnectClassBegin() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClassBegin"})
}

func (ptr *QQmlParserStatus) ClassBegin() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClassBegin"})
}

func (ptr *QQmlParserStatus) ConnectComponentComplete(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectComponentComplete", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlParserStatus) DisconnectComponentComplete() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectComponentComplete"})
}

func (ptr *QQmlParserStatus) ComponentComplete() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ComponentComplete"})
}

type QQmlProperty struct {
	internal.Internal
}

type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (ptr *QQmlProperty) QQmlProperty_PTR() *QQmlProperty {
	return ptr
}

func (ptr *QQmlProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlProperty(ptr QQmlProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProperty_PTR().Pointer()
	}
	return nil
}

func (n *QQmlProperty) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPropertyFromPointer(ptr unsafe.Pointer) (n *QQmlProperty) {
	n = new(QQmlProperty)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlProperty")
	return
}

func (ptr *QQmlProperty) DestroyQQmlProperty() {
}

//go:generate stringer -type=QQmlProperty__PropertyTypeCategory
//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(3)
)

//go:generate stringer -type=QQmlProperty__Type
//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        QQmlProperty__Type = QQmlProperty__Type(0)
	QQmlProperty__Property       QQmlProperty__Type = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty QQmlProperty__Type = QQmlProperty__Type(2)
)

func NewQQmlProperty() *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty", ""}).(*QQmlProperty)
}

func NewQQmlProperty2(obj core.QObject_ITF) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty2", "", obj}).(*QQmlProperty)
}

func NewQQmlProperty3(obj core.QObject_ITF, ctxt QQmlContext_ITF) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty3", "", obj, ctxt}).(*QQmlProperty)
}

func NewQQmlProperty4(obj core.QObject_ITF, engine QQmlEngine_ITF) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty4", "", obj, engine}).(*QQmlProperty)
}

func NewQQmlProperty5(obj core.QObject_ITF, name string) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty5", "", obj, name}).(*QQmlProperty)
}

func NewQQmlProperty6(obj core.QObject_ITF, name string, ctxt QQmlContext_ITF) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty6", "", obj, name, ctxt}).(*QQmlProperty)
}

func NewQQmlProperty7(obj core.QObject_ITF, name string, engine QQmlEngine_ITF) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty7", "", obj, name, engine}).(*QQmlProperty)
}

func NewQQmlProperty8(other QQmlProperty_ITF) *QQmlProperty {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlProperty8", "", other}).(*QQmlProperty)
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObject_ITF, slot string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifySignal", dest, slot}).(bool)
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObject_ITF, method int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifySignal2", dest, method}).(bool)
}

func (ptr *QQmlProperty) HasNotifySignal() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasNotifySignal"}).(bool)
}

func (ptr *QQmlProperty) Index() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index"}).(float64))
}

func (ptr *QQmlProperty) IsDesignable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDesignable"}).(bool)
}

func (ptr *QQmlProperty) IsProperty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsProperty"}).(bool)
}

func (ptr *QQmlProperty) IsResettable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsResettable"}).(bool)
}

func (ptr *QQmlProperty) IsSignalProperty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSignalProperty"}).(bool)
}

func (ptr *QQmlProperty) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QQmlProperty) IsWritable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsWritable"}).(bool)
}

func (ptr *QQmlProperty) Method() *core.QMetaMethod {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Method"}).(*core.QMetaMethod)
}

func (ptr *QQmlProperty) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NeedsNotifySignal"}).(bool)
}

func (ptr *QQmlProperty) Object() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Object"}).(*core.QObject)
}

func (ptr *QQmlProperty) PropertyType() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyType"}).(float64))
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {

	return QQmlProperty__PropertyTypeCategory(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyTypeCategory"}).(float64))
}

func (ptr *QQmlProperty) PropertyTypeName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyTypeName"}).(string)
}

func (ptr *QQmlProperty) Read() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Read"}).(*core.QVariant)
}

func QQmlProperty_Read2(object core.QObject_ITF, name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Read2", "", object, name}).(*core.QVariant)
}

func (ptr *QQmlProperty) Read2(object core.QObject_ITF, name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Read2", "", object, name}).(*core.QVariant)
}

func QQmlProperty_Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Read3", "", object, name, ctxt}).(*core.QVariant)
}

func (ptr *QQmlProperty) Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Read3", "", object, name, ctxt}).(*core.QVariant)
}

func QQmlProperty_Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Read4", "", object, name, engine}).(*core.QVariant)
}

func (ptr *QQmlProperty) Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Read4", "", object, name, engine}).(*core.QVariant)
}

func (ptr *QQmlProperty) Reset() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reset"}).(bool)
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {

	return QQmlProperty__Type(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QQmlProperty) Write(value core.QVariant_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Write", value}).(bool)
}

func QQmlProperty_Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Write2", "", object, name, value}).(bool)
}

func (ptr *QQmlProperty) Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Write2", "", object, name, value}).(bool)
}

func QQmlProperty_Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Write3", "", object, name, value, ctxt}).(bool)
}

func (ptr *QQmlProperty) Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Write3", "", object, name, value, ctxt}).(bool)
}

func QQmlProperty_Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Write4", "", object, name, value, engine}).(bool)
}

func (ptr *QQmlProperty) Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.QQmlProperty_Write4", "", object, name, value, engine}).(bool)
}

type QQmlPropertyMap struct {
	core.QObject
}

type QQmlPropertyMap_ITF interface {
	core.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return ptr
}

func (ptr *QQmlPropertyMap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlPropertyMap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMap_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPropertyMap) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlPropertyMap) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlPropertyMapFromPointer(ptr unsafe.Pointer) (n *QQmlPropertyMap) {
	n = new(QQmlPropertyMap)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlPropertyMap")
	return
}
func NewQQmlPropertyMap(parent core.QObject_ITF) *QQmlPropertyMap {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlPropertyMap", "", parent}).(*QQmlPropertyMap)
}

func (ptr *QQmlPropertyMap) Clear(key string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear", key})
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", key}).(bool)
}

func (ptr *QQmlPropertyMap) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QQmlPropertyMap) Insert(key string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", key, value})
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QQmlPropertyMap) Keys() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Keys"}).([]string)
}

func (ptr *QQmlPropertyMap) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QQmlPropertyMap) ConnectUpdateValue(f func(key string, input *core.QVariant) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlPropertyMap) DisconnectUpdateValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateValue"})
}

func (ptr *QQmlPropertyMap) UpdateValue(key string, input core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateValue", key, input}).(*core.QVariant)
}

func (ptr *QQmlPropertyMap) UpdateValueDefault(key string, input core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateValueDefault", key, input}).(*core.QVariant)
}

func (ptr *QQmlPropertyMap) Value(key string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", key}).(*core.QVariant)
}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueChanged"})
}

func (ptr *QQmlPropertyMap) ValueChanged(key string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueChanged", key, value})
}

func (ptr *QQmlPropertyMap) ConnectDestroyQQmlPropertyMap(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlPropertyMap", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlPropertyMap) DisconnectDestroyQQmlPropertyMap() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlPropertyMap"})
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlPropertyMap"})
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMapDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlPropertyMapDefault"})
}

func (ptr *QQmlPropertyMap) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQmlPropertyMap) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQmlPropertyMap) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlPropertyMap) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQmlPropertyMap) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQmlPropertyMap) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQmlPropertyMap) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQmlPropertyMap) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQmlPropertyMap) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQmlPropertyMap) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQmlPropertyMap) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQmlPropertyMap) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQmlPropertyMap) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQmlPropertyMap) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQmlPropertyMap) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQmlPropertyMap) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQmlPropertyMap) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQmlPropertyMap) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlPropertyValueSource struct {
	internal.Internal
}

type QQmlPropertyValueSource_ITF interface {
	QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlPropertyValueSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlPropertyValueSource(ptr QQmlPropertyValueSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyValueSource_PTR().Pointer()
	}
	return nil
}

func (n *QQmlPropertyValueSource) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) (n *QQmlPropertyValueSource) {
	n = new(QQmlPropertyValueSource)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlPropertyValueSource")
	return
}
func NewQQmlPropertyValueSource() *QQmlPropertyValueSource {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlPropertyValueSource", ""}).(*QQmlPropertyValueSource)
}

func (ptr *QQmlPropertyValueSource) ConnectSetTarget(f func(property *QQmlProperty)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetTarget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlPropertyValueSource) DisconnectSetTarget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetTarget"})
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTarget", property})
}

func (ptr *QQmlPropertyValueSource) ConnectDestroyQQmlPropertyValueSource(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQQmlPropertyValueSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQmlPropertyValueSource) DisconnectDestroyQQmlPropertyValueSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQQmlPropertyValueSource"})
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlPropertyValueSource"})
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSourceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQQmlPropertyValueSourceDefault"})
}

type QQmlScriptString struct {
	internal.Internal
}

type QQmlScriptString_ITF interface {
	QQmlScriptString_PTR() *QQmlScriptString
}

func (ptr *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString {
	return ptr
}

func (ptr *QQmlScriptString) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQmlScriptString) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQmlScriptString(ptr QQmlScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlScriptString_PTR().Pointer()
	}
	return nil
}

func (n *QQmlScriptString) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQmlScriptStringFromPointer(ptr unsafe.Pointer) (n *QQmlScriptString) {
	n = new(QQmlScriptString)
	n.InitFromInternal(uintptr(ptr), "qml.QQmlScriptString")
	return
}

func (ptr *QQmlScriptString) DestroyQQmlScriptString() {
}

func NewQQmlScriptString() *QQmlScriptString {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlScriptString", ""}).(*QQmlScriptString)
}

func NewQQmlScriptString2(other QQmlScriptString_ITF) *QQmlScriptString {

	return internal.CallLocalFunction([]interface{}{"", "", "qml.NewQQmlScriptString2", "", other}).(*QQmlScriptString)
}

func (ptr *QQmlScriptString) BooleanLiteral(ok *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BooleanLiteral", ok}).(bool)
}

func (ptr *QQmlScriptString) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNullLiteral"}).(bool)
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUndefinedLiteral"}).(bool)
}

func (ptr *QQmlScriptString) NumberLiteral(ok *bool) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumberLiteral", ok}).(float64)
}

func (ptr *QQmlScriptString) StringLiteral() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StringLiteral"}).(string)
}

func init() {
	internal.ConstructorTable["qml.QJSEngine"] = NewQJSEngineFromPointer
	internal.ConstructorTable["qml.QJSValue"] = NewQJSValueFromPointer
	internal.ConstructorTable["qml.QJSValueIterator"] = NewQJSValueIteratorFromPointer
	internal.ConstructorTable["qml.QQmlAbstractUrlInterceptor"] = NewQQmlAbstractUrlInterceptorFromPointer
	internal.ConstructorTable["qml.QQmlApplicationEngine"] = NewQQmlApplicationEngineFromPointer
	internal.ConstructorTable["qml.QQmlComponent"] = NewQQmlComponentFromPointer
	internal.ConstructorTable["qml.QQmlContext"] = NewQQmlContextFromPointer
	internal.ConstructorTable["qml.QQmlDebuggingEnabler"] = NewQQmlDebuggingEnablerFromPointer
	internal.ConstructorTable["qml.QQmlEngine"] = NewQQmlEngineFromPointer
	internal.ConstructorTable["qml.QQmlError"] = NewQQmlErrorFromPointer
	internal.ConstructorTable["qml.QQmlExpression"] = NewQQmlExpressionFromPointer
	internal.ConstructorTable["qml.QQmlExtensionPlugin"] = NewQQmlExtensionPluginFromPointer
	internal.ConstructorTable["qml.QQmlFileSelector"] = NewQQmlFileSelectorFromPointer
	internal.ConstructorTable["qml.QQmlImageProviderBase"] = NewQQmlImageProviderBaseFromPointer
	internal.ConstructorTable["qml.QQmlIncubationController"] = NewQQmlIncubationControllerFromPointer
	internal.ConstructorTable["qml.QQmlIncubator"] = NewQQmlIncubatorFromPointer
	internal.ConstructorTable["qml.QQmlInfo"] = NewQQmlInfoFromPointer
	internal.ConstructorTable["qml.QQmlListReference"] = NewQQmlListReferenceFromPointer
	internal.ConstructorTable["qml.QQmlNetworkAccessManagerFactory"] = NewQQmlNetworkAccessManagerFactoryFromPointer
	internal.ConstructorTable["qml.QQmlParserStatus"] = NewQQmlParserStatusFromPointer
	internal.ConstructorTable["qml.QQmlProperty"] = NewQQmlPropertyFromPointer
	internal.ConstructorTable["qml.QQmlPropertyMap"] = NewQQmlPropertyMapFromPointer
	internal.ConstructorTable["qml.QQmlPropertyValueSource"] = NewQQmlPropertyValueSourceFromPointer
	internal.ConstructorTable["qml.QQmlScriptString"] = NewQQmlScriptStringFromPointer
}
