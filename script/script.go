// +build !minimal

package script

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/internal"
	"unsafe"
)

type QScriptClass struct {
	internal.Internal
}

type QScriptClass_ITF interface {
	QScriptClass_PTR() *QScriptClass
}

func (ptr *QScriptClass) QScriptClass_PTR() *QScriptClass {
	return ptr
}

func (ptr *QScriptClass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptClass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptClass(ptr QScriptClass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClass_PTR().Pointer()
	}
	return nil
}

func (n *QScriptClass) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptClassFromPointer(ptr unsafe.Pointer) (n *QScriptClass) {
	n = new(QScriptClass)
	n.InitFromInternal(uintptr(ptr), "script.QScriptClass")
	return
}

//go:generate stringer -type=QScriptClass__QueryFlag
//QScriptClass::QueryFlag
type QScriptClass__QueryFlag int64

const (
	QScriptClass__HandlesReadAccess  QScriptClass__QueryFlag = QScriptClass__QueryFlag(0x01)
	QScriptClass__HandlesWriteAccess QScriptClass__QueryFlag = QScriptClass__QueryFlag(0x02)
)

//go:generate stringer -type=QScriptClass__Extension
//QScriptClass::Extension
type QScriptClass__Extension int64

const (
	QScriptClass__Callable    QScriptClass__Extension = QScriptClass__Extension(0)
	QScriptClass__HasInstance QScriptClass__Extension = QScriptClass__Extension(1)
)

func NewQScriptClass(engine QScriptEngine_ITF) *QScriptClass {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptClass", "", engine}).(*QScriptClass)
}

func (ptr *QScriptClass) Engine() *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QScriptEngine)
}

func (ptr *QScriptClass) ConnectExtension(f func(extensi QScriptClass__Extension, argument *core.QVariant) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExtension"})
}

func (ptr *QScriptClass) Extension(extensi QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Extension", extensi, argument}).(*core.QVariant)
}

func (ptr *QScriptClass) ExtensionDefault(extensi QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtensionDefault", extensi, argument}).(*core.QVariant)
}

func (ptr *QScriptClass) ConnectName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectName"})
}

func (ptr *QScriptClass) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QScriptClass) NameDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameDefault"}).(string)
}

func (ptr *QScriptClass) ConnectNewIterator(f func(object *QScriptValue) *QScriptClassPropertyIterator) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewIterator", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectNewIterator() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewIterator"})
}

func (ptr *QScriptClass) NewIterator(object QScriptValue_ITF) *QScriptClassPropertyIterator {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewIterator", object}).(*QScriptClassPropertyIterator)
}

func (ptr *QScriptClass) NewIteratorDefault(object QScriptValue_ITF) *QScriptClassPropertyIterator {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewIteratorDefault", object}).(*QScriptClassPropertyIterator)
}

func (ptr *QScriptClass) ConnectProperty(f func(object *QScriptValue, name *QScriptString, id uint) *QScriptValue) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProperty"})
}

func (ptr *QScriptClass) Property(object QScriptValue_ITF, name QScriptString_ITF, id uint) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", object, name, id}).(*QScriptValue)
}

func (ptr *QScriptClass) PropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyDefault", object, name, id}).(*QScriptValue)
}

func (ptr *QScriptClass) ConnectPropertyFlags(f func(object *QScriptValue, name *QScriptString, id uint) QScriptValue__PropertyFlag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPropertyFlags", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectPropertyFlags() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPropertyFlags"})
}

func (ptr *QScriptClass) PropertyFlags(object QScriptValue_ITF, name QScriptString_ITF, id uint) QScriptValue__PropertyFlag {

	return QScriptValue__PropertyFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyFlags", object, name, id}).(float64))
}

func (ptr *QScriptClass) PropertyFlagsDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint) QScriptValue__PropertyFlag {

	return QScriptValue__PropertyFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyFlagsDefault", object, name, id}).(float64))
}

func (ptr *QScriptClass) ConnectPrototype(f func() *QScriptValue) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrototype", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectPrototype() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrototype"})
}

func (ptr *QScriptClass) Prototype() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prototype"}).(*QScriptValue)
}

func (ptr *QScriptClass) PrototypeDefault() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrototypeDefault"}).(*QScriptValue)
}

func (ptr *QScriptClass) ConnectQueryProperty(f func(object *QScriptValue, name *QScriptString, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectQueryProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectQueryProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectQueryProperty"})
}

func (ptr *QScriptClass) QueryProperty(object QScriptValue_ITF, name QScriptString_ITF, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag {

	return QScriptClass__QueryFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryProperty", object, name, flags, id}).(float64))
}

func (ptr *QScriptClass) QueryPropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag {

	return QScriptClass__QueryFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryPropertyDefault", object, name, flags, id}).(float64))
}

func (ptr *QScriptClass) ConnectSetProperty(f func(object *QScriptValue, name *QScriptString, id uint, value *QScriptValue)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectSetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetProperty"})
}

func (ptr *QScriptClass) SetProperty(object QScriptValue_ITF, name QScriptString_ITF, id uint, value QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", object, name, id, value})
}

func (ptr *QScriptClass) SetPropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint, value QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPropertyDefault", object, name, id, value})
}

func (ptr *QScriptClass) ConnectSupportsExtension(f func(extensi QScriptClass__Extension) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportsExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectSupportsExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportsExtension"})
}

func (ptr *QScriptClass) SupportsExtension(extensi QScriptClass__Extension) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportsExtension", extensi}).(bool)
}

func (ptr *QScriptClass) SupportsExtensionDefault(extensi QScriptClass__Extension) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportsExtensionDefault", extensi}).(bool)
}

func (ptr *QScriptClass) ConnectDestroyQScriptClass(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScriptClass", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptClass) DisconnectDestroyQScriptClass() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScriptClass"})
}

func (ptr *QScriptClass) DestroyQScriptClass() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptClass"})
}

func (ptr *QScriptClass) DestroyQScriptClassDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptClassDefault"})
}

type QScriptClassPropertyIterator struct {
	internal.Internal
}

type QScriptClassPropertyIterator_ITF interface {
	QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator
}

func (ptr *QScriptClassPropertyIterator) QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator {
	return ptr
}

func (ptr *QScriptClassPropertyIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptClassPropertyIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptClassPropertyIterator(ptr QScriptClassPropertyIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClassPropertyIterator_PTR().Pointer()
	}
	return nil
}

func (n *QScriptClassPropertyIterator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) (n *QScriptClassPropertyIterator) {
	n = new(QScriptClassPropertyIterator)
	n.InitFromInternal(uintptr(ptr), "script.QScriptClassPropertyIterator")
	return
}

type QScriptContext struct {
	internal.Internal
}

type QScriptContext_ITF interface {
	QScriptContext_PTR() *QScriptContext
}

func (ptr *QScriptContext) QScriptContext_PTR() *QScriptContext {
	return ptr
}

func (ptr *QScriptContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptContext(ptr QScriptContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContext_PTR().Pointer()
	}
	return nil
}

func (n *QScriptContext) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptContextFromPointer(ptr unsafe.Pointer) (n *QScriptContext) {
	n = new(QScriptContext)
	n.InitFromInternal(uintptr(ptr), "script.QScriptContext")
	return
}

//go:generate stringer -type=QScriptContext__ExecutionState
//QScriptContext::ExecutionState
type QScriptContext__ExecutionState int64

const (
	QScriptContext__NormalState    QScriptContext__ExecutionState = QScriptContext__ExecutionState(0)
	QScriptContext__ExceptionState QScriptContext__ExecutionState = QScriptContext__ExecutionState(1)
)

//go:generate stringer -type=QScriptContext__Error
//QScriptContext::Error
type QScriptContext__Error int64

const (
	QScriptContext__UnknownError   QScriptContext__Error = QScriptContext__Error(0)
	QScriptContext__ReferenceError QScriptContext__Error = QScriptContext__Error(1)
	QScriptContext__SyntaxError    QScriptContext__Error = QScriptContext__Error(2)
	QScriptContext__TypeError      QScriptContext__Error = QScriptContext__Error(3)
	QScriptContext__RangeError     QScriptContext__Error = QScriptContext__Error(4)
	QScriptContext__URIError       QScriptContext__Error = QScriptContext__Error(5)
)

func (ptr *QScriptContext) ActivationObject() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActivationObject"}).(*QScriptValue)
}

func (ptr *QScriptContext) Argument(index int) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Argument", index}).(*QScriptValue)
}

func (ptr *QScriptContext) ArgumentCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ArgumentCount"}).(float64))
}

func (ptr *QScriptContext) ArgumentsObject() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ArgumentsObject"}).(*QScriptValue)
}

func (ptr *QScriptContext) Backtrace() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Backtrace"}).([]string)
}

func (ptr *QScriptContext) Callee() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Callee"}).(*QScriptValue)
}

func (ptr *QScriptContext) Engine() *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QScriptEngine)
}

func (ptr *QScriptContext) IsCalledAsConstructor() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCalledAsConstructor"}).(bool)
}

func (ptr *QScriptContext) ParentContext() *QScriptContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentContext"}).(*QScriptContext)
}

func (ptr *QScriptContext) SetActivationObject(activation QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActivationObject", activation})
}

func (ptr *QScriptContext) SetThisObject(thisObject QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetThisObject", thisObject})
}

func (ptr *QScriptContext) State() QScriptContext__ExecutionState {

	return QScriptContext__ExecutionState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QScriptContext) ThisObject() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThisObject"}).(*QScriptValue)
}

func (ptr *QScriptContext) ThrowError(error QScriptContext__Error, text string) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThrowError", error, text}).(*QScriptValue)
}

func (ptr *QScriptContext) ThrowError2(text string) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThrowError2", text}).(*QScriptValue)
}

func (ptr *QScriptContext) ThrowValue(value QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThrowValue", value}).(*QScriptValue)
}

func (ptr *QScriptContext) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QScriptContext) DestroyQScriptContext() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptContext"})
}

type QScriptContextInfo struct {
	internal.Internal
}

type QScriptContextInfo_ITF interface {
	QScriptContextInfo_PTR() *QScriptContextInfo
}

func (ptr *QScriptContextInfo) QScriptContextInfo_PTR() *QScriptContextInfo {
	return ptr
}

func (ptr *QScriptContextInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptContextInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptContextInfo(ptr QScriptContextInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContextInfo_PTR().Pointer()
	}
	return nil
}

func (n *QScriptContextInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptContextInfoFromPointer(ptr unsafe.Pointer) (n *QScriptContextInfo) {
	n = new(QScriptContextInfo)
	n.InitFromInternal(uintptr(ptr), "script.QScriptContextInfo")
	return
}

//go:generate stringer -type=QScriptContextInfo__FunctionType
//QScriptContextInfo::FunctionType
type QScriptContextInfo__FunctionType int64

const (
	QScriptContextInfo__ScriptFunction     QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(0)
	QScriptContextInfo__QtFunction         QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(1)
	QScriptContextInfo__QtPropertyFunction QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(2)
	QScriptContextInfo__NativeFunction     QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(3)
)

func NewQScriptContextInfo(context QScriptContext_ITF) *QScriptContextInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptContextInfo", "", context}).(*QScriptContextInfo)
}

func NewQScriptContextInfo2(other QScriptContextInfo_ITF) *QScriptContextInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptContextInfo2", "", other}).(*QScriptContextInfo)
}

func NewQScriptContextInfo3() *QScriptContextInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptContextInfo3", ""}).(*QScriptContextInfo)
}

func (ptr *QScriptContextInfo) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QScriptContextInfo) FunctionEndLineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionEndLineNumber"}).(float64))
}

func (ptr *QScriptContextInfo) FunctionMetaIndex() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionMetaIndex"}).(float64))
}

func (ptr *QScriptContextInfo) FunctionName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionName"}).(string)
}

func (ptr *QScriptContextInfo) FunctionParameterNames() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionParameterNames"}).([]string)
}

func (ptr *QScriptContextInfo) FunctionStartLineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionStartLineNumber"}).(float64))
}

func (ptr *QScriptContextInfo) FunctionType() QScriptContextInfo__FunctionType {

	return QScriptContextInfo__FunctionType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionType"}).(float64))
}

func (ptr *QScriptContextInfo) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QScriptContextInfo) LineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineNumber"}).(float64))
}

func (ptr *QScriptContextInfo) ScriptId() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScriptId"}).(float64))
}

func (ptr *QScriptContextInfo) DestroyQScriptContextInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptContextInfo"})
}

type QScriptEngine struct {
	core.QObject
}

type QScriptEngine_ITF interface {
	core.QObject_ITF
	QScriptEngine_PTR() *QScriptEngine
}

func (ptr *QScriptEngine) QScriptEngine_PTR() *QScriptEngine {
	return ptr
}

func (ptr *QScriptEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScriptEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScriptEngine(ptr QScriptEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngine_PTR().Pointer()
	}
	return nil
}

func (n *QScriptEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScriptEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScriptEngineFromPointer(ptr unsafe.Pointer) (n *QScriptEngine) {
	n = new(QScriptEngine)
	n.InitFromInternal(uintptr(ptr), "script.QScriptEngine")
	return
}

//go:generate stringer -type=QScriptEngine__ValueOwnership
//QScriptEngine::ValueOwnership
type QScriptEngine__ValueOwnership int64

const (
	QScriptEngine__QtOwnership     QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(0)
	QScriptEngine__ScriptOwnership QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(1)
	QScriptEngine__AutoOwnership   QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(2)
)

//go:generate stringer -type=QScriptEngine__QObjectWrapOption
//QScriptEngine::QObjectWrapOption
type QScriptEngine__QObjectWrapOption int64

const (
	QScriptEngine__ExcludeChildObjects         QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0001)
	QScriptEngine__ExcludeSuperClassMethods    QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0002)
	QScriptEngine__ExcludeSuperClassProperties QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0004)
	QScriptEngine__ExcludeSuperClassContents   QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0006)
	QScriptEngine__SkipMethodsInEnumeration    QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0008)
	QScriptEngine__ExcludeDeleteLater          QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0010)
	QScriptEngine__ExcludeSlots                QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0020)
	QScriptEngine__AutoCreateDynamicProperties QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0100)
	QScriptEngine__PreferExistingWrapperObject QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0200)
)

func NewQScriptEngine() *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptEngine", ""}).(*QScriptEngine)
}

func NewQScriptEngine2(parent core.QObject_ITF) *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptEngine2", "", parent}).(*QScriptEngine)
}

func (ptr *QScriptEngine) AbortEvaluation(result QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AbortEvaluation", result})
}

func (ptr *QScriptEngine) Agent() *QScriptEngineAgent {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Agent"}).(*QScriptEngineAgent)
}

func (ptr *QScriptEngine) AvailableExtensions() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableExtensions"}).([]string)
}

func QScriptEngine_CheckSyntax(program string) *QScriptSyntaxCheckResult {

	return internal.CallLocalFunction([]interface{}{"", "", "script.QScriptEngine_CheckSyntax", "", program}).(*QScriptSyntaxCheckResult)
}

func (ptr *QScriptEngine) CheckSyntax(program string) *QScriptSyntaxCheckResult {

	return internal.CallLocalFunction([]interface{}{"", "", "script.QScriptEngine_CheckSyntax", "", program}).(*QScriptSyntaxCheckResult)
}

func (ptr *QScriptEngine) ClearExceptions() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearExceptions"})
}

func (ptr *QScriptEngine) CollectGarbage() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollectGarbage"})
}

func (ptr *QScriptEngine) CurrentContext() *QScriptContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentContext"}).(*QScriptContext)
}

func (ptr *QScriptEngine) DefaultPrototype(metaTypeId int) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultPrototype", metaTypeId}).(*QScriptValue)
}

func (ptr *QScriptEngine) Evaluate(program string, fileName string, lineNumber int) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Evaluate", program, fileName, lineNumber}).(*QScriptValue)
}

func (ptr *QScriptEngine) Evaluate2(program QScriptProgram_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Evaluate2", program}).(*QScriptValue)
}

func (ptr *QScriptEngine) GlobalObject() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GlobalObject"}).(*QScriptValue)
}

func (ptr *QScriptEngine) HasUncaughtException() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasUncaughtException"}).(bool)
}

func (ptr *QScriptEngine) ImportExtension(extensi string) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImportExtension", extensi}).(*QScriptValue)
}

func (ptr *QScriptEngine) ImportedExtensions() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImportedExtensions"}).([]string)
}

func (ptr *QScriptEngine) InstallTranslatorFunctions(object QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstallTranslatorFunctions", object})
}

func (ptr *QScriptEngine) IsEvaluating() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEvaluating"}).(bool)
}

func (ptr *QScriptEngine) NewArray(length uint) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewArray", length}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewDate2(value core.QDateTime_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewDate2", value}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewObject() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewObject"}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewObject2(scriptClass QScriptClass_ITF, data QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewObject2", scriptClass, data}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewQMetaObject(metaObject core.QMetaObject_ITF, ctor QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewQMetaObject", metaObject, ctor}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewQObject(object core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewQObject", object, ownership, options}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewQObject2(scriptObject QScriptValue_ITF, qtObject core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewQObject2", scriptObject, qtObject, ownership, options}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewRegExp(regexp core.QRegExp_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewRegExp", regexp}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewRegExp2(pattern string, flags string) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewRegExp2", pattern, flags}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewVariant(value core.QVariant_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewVariant", value}).(*QScriptValue)
}

func (ptr *QScriptEngine) NewVariant2(object QScriptValue_ITF, value core.QVariant_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewVariant2", object, value}).(*QScriptValue)
}

func (ptr *QScriptEngine) NullValue() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NullValue"}).(*QScriptValue)
}

func (ptr *QScriptEngine) PopContext() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PopContext"})
}

func (ptr *QScriptEngine) ProcessEventsInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessEventsInterval"}).(float64))
}

func (ptr *QScriptEngine) PushContext() *QScriptContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PushContext"}).(*QScriptContext)
}

func (ptr *QScriptEngine) ReportAdditionalMemoryCost(size int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReportAdditionalMemoryCost", size})
}

func (ptr *QScriptEngine) SetAgent(agent QScriptEngineAgent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAgent", agent})
}

func (ptr *QScriptEngine) SetDefaultPrototype(metaTypeId int, prototype QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDefaultPrototype", metaTypeId, prototype})
}

func (ptr *QScriptEngine) SetGlobalObject(object QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGlobalObject", object})
}

func (ptr *QScriptEngine) SetProcessEventsInterval(interval int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProcessEventsInterval", interval})
}

func (ptr *QScriptEngine) ConnectSignalHandlerException(f func(exception *QScriptValue)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSignalHandlerException", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngine) DisconnectSignalHandlerException() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSignalHandlerException"})
}

func (ptr *QScriptEngine) SignalHandlerException(exception QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SignalHandlerException", exception})
}

func (ptr *QScriptEngine) ToObject(value QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToObject", value}).(*QScriptValue)
}

func (ptr *QScriptEngine) ToStringHandle(str string) *QScriptString {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToStringHandle", str}).(*QScriptString)
}

func (ptr *QScriptEngine) UncaughtException() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UncaughtException"}).(*QScriptValue)
}

func (ptr *QScriptEngine) UncaughtExceptionBacktrace() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UncaughtExceptionBacktrace"}).([]string)
}

func (ptr *QScriptEngine) UncaughtExceptionLineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UncaughtExceptionLineNumber"}).(float64))
}

func (ptr *QScriptEngine) UndefinedValue() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UndefinedValue"}).(*QScriptValue)
}

func (ptr *QScriptEngine) ConnectDestroyQScriptEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScriptEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngine) DisconnectDestroyQScriptEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScriptEngine"})
}

func (ptr *QScriptEngine) DestroyQScriptEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptEngine"})
}

func (ptr *QScriptEngine) DestroyQScriptEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptEngineDefault"})
}

func (ptr *QScriptEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScriptEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScriptEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScriptEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScriptEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScriptEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScriptEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScriptEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScriptEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScriptEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScriptEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScriptEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScriptEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScriptEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScriptEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScriptEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScriptEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScriptEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QScriptEngineAgent struct {
	internal.Internal
}

type QScriptEngineAgent_ITF interface {
	QScriptEngineAgent_PTR() *QScriptEngineAgent
}

func (ptr *QScriptEngineAgent) QScriptEngineAgent_PTR() *QScriptEngineAgent {
	return ptr
}

func (ptr *QScriptEngineAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptEngineAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptEngineAgent(ptr QScriptEngineAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngineAgent_PTR().Pointer()
	}
	return nil
}

func (n *QScriptEngineAgent) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptEngineAgentFromPointer(ptr unsafe.Pointer) (n *QScriptEngineAgent) {
	n = new(QScriptEngineAgent)
	n.InitFromInternal(uintptr(ptr), "script.QScriptEngineAgent")
	return
}

//go:generate stringer -type=QScriptEngineAgent__Extension
//QScriptEngineAgent::Extension
type QScriptEngineAgent__Extension int64

const (
	QScriptEngineAgent__DebuggerInvocationRequest QScriptEngineAgent__Extension = QScriptEngineAgent__Extension(0)
)

func NewQScriptEngineAgent(engine QScriptEngine_ITF) *QScriptEngineAgent {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptEngineAgent", "", engine}).(*QScriptEngineAgent)
}

func (ptr *QScriptEngineAgent) ConnectContextPop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContextPop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectContextPop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContextPop"})
}

func (ptr *QScriptEngineAgent) ContextPop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextPop"})
}

func (ptr *QScriptEngineAgent) ContextPopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextPopDefault"})
}

func (ptr *QScriptEngineAgent) ConnectContextPush(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContextPush", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectContextPush() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContextPush"})
}

func (ptr *QScriptEngineAgent) ContextPush() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextPush"})
}

func (ptr *QScriptEngineAgent) ContextPushDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextPushDefault"})
}

func (ptr *QScriptEngineAgent) Engine() *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QScriptEngine)
}

func (ptr *QScriptEngineAgent) ConnectExceptionCatch(f func(scriptId int64, exception *QScriptValue)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExceptionCatch", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectExceptionCatch() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExceptionCatch"})
}

func (ptr *QScriptEngineAgent) ExceptionCatch(scriptId int64, exception QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExceptionCatch", scriptId, exception})
}

func (ptr *QScriptEngineAgent) ExceptionCatchDefault(scriptId int64, exception QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExceptionCatchDefault", scriptId, exception})
}

func (ptr *QScriptEngineAgent) ConnectExceptionThrow(f func(scriptId int64, exception *QScriptValue, hasHandler bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExceptionThrow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectExceptionThrow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExceptionThrow"})
}

func (ptr *QScriptEngineAgent) ExceptionThrow(scriptId int64, exception QScriptValue_ITF, hasHandler bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExceptionThrow", scriptId, exception, hasHandler})
}

func (ptr *QScriptEngineAgent) ExceptionThrowDefault(scriptId int64, exception QScriptValue_ITF, hasHandler bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExceptionThrowDefault", scriptId, exception, hasHandler})
}

func (ptr *QScriptEngineAgent) ConnectExtension(f func(extensi QScriptEngineAgent__Extension, argument *core.QVariant) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExtension"})
}

func (ptr *QScriptEngineAgent) Extension(extensi QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Extension", extensi, argument}).(*core.QVariant)
}

func (ptr *QScriptEngineAgent) ExtensionDefault(extensi QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtensionDefault", extensi, argument}).(*core.QVariant)
}

func (ptr *QScriptEngineAgent) ConnectFunctionEntry(f func(scriptId int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFunctionEntry", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectFunctionEntry() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFunctionEntry"})
}

func (ptr *QScriptEngineAgent) FunctionEntry(scriptId int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionEntry", scriptId})
}

func (ptr *QScriptEngineAgent) FunctionEntryDefault(scriptId int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionEntryDefault", scriptId})
}

func (ptr *QScriptEngineAgent) ConnectFunctionExit(f func(scriptId int64, returnValue *QScriptValue)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFunctionExit", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectFunctionExit() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFunctionExit"})
}

func (ptr *QScriptEngineAgent) FunctionExit(scriptId int64, returnValue QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionExit", scriptId, returnValue})
}

func (ptr *QScriptEngineAgent) FunctionExitDefault(scriptId int64, returnValue QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionExitDefault", scriptId, returnValue})
}

func (ptr *QScriptEngineAgent) ConnectPositionChange(f func(scriptId int64, lineNumber int, columnNumber int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionChange", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectPositionChange() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionChange"})
}

func (ptr *QScriptEngineAgent) PositionChange(scriptId int64, lineNumber int, columnNumber int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionChange", scriptId, lineNumber, columnNumber})
}

func (ptr *QScriptEngineAgent) PositionChangeDefault(scriptId int64, lineNumber int, columnNumber int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionChangeDefault", scriptId, lineNumber, columnNumber})
}

func (ptr *QScriptEngineAgent) ConnectScriptLoad(f func(id int64, program string, fileName string, baseLineNumber int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScriptLoad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectScriptLoad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScriptLoad"})
}

func (ptr *QScriptEngineAgent) ScriptLoad(id int64, program string, fileName string, baseLineNumber int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScriptLoad", id, program, fileName, baseLineNumber})
}

func (ptr *QScriptEngineAgent) ScriptLoadDefault(id int64, program string, fileName string, baseLineNumber int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScriptLoadDefault", id, program, fileName, baseLineNumber})
}

func (ptr *QScriptEngineAgent) ConnectScriptUnload(f func(id int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScriptUnload", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectScriptUnload() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScriptUnload"})
}

func (ptr *QScriptEngineAgent) ScriptUnload(id int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScriptUnload", id})
}

func (ptr *QScriptEngineAgent) ScriptUnloadDefault(id int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScriptUnloadDefault", id})
}

func (ptr *QScriptEngineAgent) ConnectSupportsExtension(f func(extensi QScriptEngineAgent__Extension) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportsExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectSupportsExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportsExtension"})
}

func (ptr *QScriptEngineAgent) SupportsExtension(extensi QScriptEngineAgent__Extension) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportsExtension", extensi}).(bool)
}

func (ptr *QScriptEngineAgent) SupportsExtensionDefault(extensi QScriptEngineAgent__Extension) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportsExtensionDefault", extensi}).(bool)
}

func (ptr *QScriptEngineAgent) ConnectDestroyQScriptEngineAgent(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScriptEngineAgent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineAgent) DisconnectDestroyQScriptEngineAgent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScriptEngineAgent"})
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptEngineAgent"})
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgentDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptEngineAgentDefault"})
}

type QScriptExtensionInterface struct {
	internal.Internal
}

type QScriptExtensionInterface_ITF interface {
	QScriptExtensionInterface_PTR() *QScriptExtensionInterface
}

func (ptr *QScriptExtensionInterface) QScriptExtensionInterface_PTR() *QScriptExtensionInterface {
	return ptr
}

func (ptr *QScriptExtensionInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptExtensionInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptExtensionInterface(ptr QScriptExtensionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionInterface_PTR().Pointer()
	}
	return nil
}

func (n *QScriptExtensionInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptExtensionInterfaceFromPointer(ptr unsafe.Pointer) (n *QScriptExtensionInterface) {
	n = new(QScriptExtensionInterface)
	n.InitFromInternal(uintptr(ptr), "script.QScriptExtensionInterface")
	return
}

func (ptr *QScriptExtensionInterface) DestroyQScriptExtensionInterface() {
}

type QScriptExtensionPlugin struct {
	core.QObject
	QScriptExtensionInterface
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
	QScriptExtensionInterface_ITF
	QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QScriptExtensionInterface_PTR().SetPointer(p)
	}
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func (n *QScriptExtensionPlugin) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QScriptExtensionInterface_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScriptExtensionPlugin) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) (n *QScriptExtensionPlugin) {
	n = new(QScriptExtensionPlugin)
	n.InitFromInternal(uintptr(ptr), "script.QScriptExtensionPlugin")
	return
}
func NewQScriptExtensionPlugin(parent core.QObject_ITF) *QScriptExtensionPlugin {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptExtensionPlugin", "", parent}).(*QScriptExtensionPlugin)
}

func (ptr *QScriptExtensionPlugin) ConnectInitialize(f func(key string, engine *QScriptEngine)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectInitialize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptExtensionPlugin) DisconnectInitialize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectInitialize"})
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Initialize", key, engine})
}

func (ptr *QScriptExtensionPlugin) ConnectKeys(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectKeys", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptExtensionPlugin) DisconnectKeys() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectKeys"})
}

func (ptr *QScriptExtensionPlugin) Keys() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Keys"}).([]string)
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "SetupPackage", key, engine}).(*QScriptValue)
}

func (ptr *QScriptExtensionPlugin) ConnectDestroyQScriptExtensionPlugin(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDestroyQScriptExtensionPlugin", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptExtensionPlugin) DisconnectDestroyQScriptExtensionPlugin() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDestroyQScriptExtensionPlugin"})
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQScriptExtensionPlugin"})
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPluginDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQScriptExtensionPluginDefault"})
}

func (ptr *QScriptExtensionPlugin) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScriptExtensionPlugin) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScriptExtensionPlugin) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScriptExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScriptExtensionPlugin) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScriptExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScriptExtensionPlugin) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScriptExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEvent", event})
}

func (ptr *QScriptExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScriptExtensionPlugin) ConnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotify", sign})
}

func (ptr *QScriptExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScriptExtensionPlugin) CustomEvent(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEvent", event})
}

func (ptr *QScriptExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScriptExtensionPlugin) DeleteLater() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLater"})
}

func (ptr *QScriptExtensionPlugin) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScriptExtensionPlugin) DisconnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotify", sign})
}

func (ptr *QScriptExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScriptExtensionPlugin) Event(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Event", e}).(bool)
}

func (ptr *QScriptExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScriptExtensionPlugin) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilter", watched, event}).(bool)
}

func (ptr *QScriptExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScriptExtensionPlugin) MetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObject"}).(*core.QMetaObject)
}

func (ptr *QScriptExtensionPlugin) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScriptExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEvent", event})
}

func (ptr *QScriptExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEventDefault", event})
}

type QScriptProgram struct {
	internal.Internal
}

type QScriptProgram_ITF interface {
	QScriptProgram_PTR() *QScriptProgram
}

func (ptr *QScriptProgram) QScriptProgram_PTR() *QScriptProgram {
	return ptr
}

func (ptr *QScriptProgram) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptProgram) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptProgram(ptr QScriptProgram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptProgram_PTR().Pointer()
	}
	return nil
}

func (n *QScriptProgram) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptProgramFromPointer(ptr unsafe.Pointer) (n *QScriptProgram) {
	n = new(QScriptProgram)
	n.InitFromInternal(uintptr(ptr), "script.QScriptProgram")
	return
}
func NewQScriptProgram() *QScriptProgram {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptProgram", ""}).(*QScriptProgram)
}

func NewQScriptProgram2(sourceCode string, fileName string, firstLineNumber int) *QScriptProgram {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptProgram2", "", sourceCode, fileName, firstLineNumber}).(*QScriptProgram)
}

func NewQScriptProgram3(other QScriptProgram_ITF) *QScriptProgram {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptProgram3", "", other}).(*QScriptProgram)
}

func (ptr *QScriptProgram) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QScriptProgram) FirstLineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstLineNumber"}).(float64))
}

func (ptr *QScriptProgram) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QScriptProgram) SourceCode() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceCode"}).(string)
}

func (ptr *QScriptProgram) DestroyQScriptProgram() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptProgram"})
}

type QScriptString struct {
	internal.Internal
}

type QScriptString_ITF interface {
	QScriptString_PTR() *QScriptString
}

func (ptr *QScriptString) QScriptString_PTR() *QScriptString {
	return ptr
}

func (ptr *QScriptString) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptString) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptString(ptr QScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptString_PTR().Pointer()
	}
	return nil
}

func (n *QScriptString) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptStringFromPointer(ptr unsafe.Pointer) (n *QScriptString) {
	n = new(QScriptString)
	n.InitFromInternal(uintptr(ptr), "script.QScriptString")
	return
}
func NewQScriptString() *QScriptString {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptString", ""}).(*QScriptString)
}

func NewQScriptString2(other QScriptString_ITF) *QScriptString {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptString2", "", other}).(*QScriptString)
}

func (ptr *QScriptString) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QScriptString) ToArrayIndex(ok *bool) uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToArrayIndex", ok}).(float64))
}

func (ptr *QScriptString) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QScriptString) DestroyQScriptString() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptString"})
}

type QScriptSyntaxCheckResult struct {
	internal.Internal
}

type QScriptSyntaxCheckResult_ITF interface {
	QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult
}

func (ptr *QScriptSyntaxCheckResult) QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult {
	return ptr
}

func (ptr *QScriptSyntaxCheckResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptSyntaxCheckResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptSyntaxCheckResult(ptr QScriptSyntaxCheckResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptSyntaxCheckResult_PTR().Pointer()
	}
	return nil
}

func (n *QScriptSyntaxCheckResult) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) (n *QScriptSyntaxCheckResult) {
	n = new(QScriptSyntaxCheckResult)
	n.InitFromInternal(uintptr(ptr), "script.QScriptSyntaxCheckResult")
	return
}

//go:generate stringer -type=QScriptSyntaxCheckResult__State
//QScriptSyntaxCheckResult::State
type QScriptSyntaxCheckResult__State int64

const (
	QScriptSyntaxCheckResult__Error        QScriptSyntaxCheckResult__State = QScriptSyntaxCheckResult__State(0)
	QScriptSyntaxCheckResult__Intermediate QScriptSyntaxCheckResult__State = QScriptSyntaxCheckResult__State(1)
	QScriptSyntaxCheckResult__Valid        QScriptSyntaxCheckResult__State = QScriptSyntaxCheckResult__State(2)
)

func NewQScriptSyntaxCheckResult(other QScriptSyntaxCheckResult_ITF) *QScriptSyntaxCheckResult {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptSyntaxCheckResult", "", other}).(*QScriptSyntaxCheckResult)
}

func (ptr *QScriptSyntaxCheckResult) ErrorColumnNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorColumnNumber"}).(float64))
}

func (ptr *QScriptSyntaxCheckResult) ErrorLineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorLineNumber"}).(float64))
}

func (ptr *QScriptSyntaxCheckResult) ErrorMessage() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorMessage"}).(string)
}

func (ptr *QScriptSyntaxCheckResult) State() QScriptSyntaxCheckResult__State {

	return QScriptSyntaxCheckResult__State(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QScriptSyntaxCheckResult) DestroyQScriptSyntaxCheckResult() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptSyntaxCheckResult"})
}

type QScriptValue struct {
	internal.Internal
}

type QScriptValue_ITF interface {
	QScriptValue_PTR() *QScriptValue
}

func (ptr *QScriptValue) QScriptValue_PTR() *QScriptValue {
	return ptr
}

func (ptr *QScriptValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptValue(ptr QScriptValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValue_PTR().Pointer()
	}
	return nil
}

func (n *QScriptValue) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptValueFromPointer(ptr unsafe.Pointer) (n *QScriptValue) {
	n = new(QScriptValue)
	n.InitFromInternal(uintptr(ptr), "script.QScriptValue")
	return
}

//go:generate stringer -type=QScriptValue__ResolveFlag
//QScriptValue::ResolveFlag
type QScriptValue__ResolveFlag int64

const (
	QScriptValue__ResolveLocal     QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(0x00)
	QScriptValue__ResolvePrototype QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(0x01)
	QScriptValue__ResolveScope     QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(0x02)
	QScriptValue__ResolveFull      QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(QScriptValue__ResolvePrototype | QScriptValue__ResolveScope)
)

//go:generate stringer -type=QScriptValue__SpecialValue
//QScriptValue::SpecialValue
type QScriptValue__SpecialValue int64

const (
	QScriptValue__NullValue      QScriptValue__SpecialValue = QScriptValue__SpecialValue(0)
	QScriptValue__UndefinedValue QScriptValue__SpecialValue = QScriptValue__SpecialValue(1)
)

//go:generate stringer -type=QScriptValue__PropertyFlag
//QScriptValue::PropertyFlag
type QScriptValue__PropertyFlag int64

const (
	QScriptValue__ReadOnly          QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000001)
	QScriptValue__Undeletable       QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000002)
	QScriptValue__SkipInEnumeration QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000004)
	QScriptValue__PropertyGetter    QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000008)
	QScriptValue__PropertySetter    QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000010)
	QScriptValue__QObjectMember     QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000020)
	QScriptValue__KeepExistingFlags QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000800)
	QScriptValue__UserRange         QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0xff000000)
)

func NewQScriptValue() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue", ""}).(*QScriptValue)
}

func NewQScriptValue2(other QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue2", "", other}).(*QScriptValue)
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue10", "", value}).(*QScriptValue)
}

func NewQScriptValue11(value bool) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue11", "", value}).(*QScriptValue)
}

func NewQScriptValue12(value int) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue12", "", value}).(*QScriptValue)
}

func NewQScriptValue13(value uint) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue13", "", value}).(*QScriptValue)
}

func NewQScriptValue15(value string) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue15", "", value}).(*QScriptValue)
}

func NewQScriptValue16(value core.QLatin1String_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue16", "", value}).(*QScriptValue)
}

func NewQScriptValue17(value string) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", "", "script.NewQScriptValue17", "", value}).(*QScriptValue)
}

func (ptr *QScriptValue) Call2(thisObject QScriptValue_ITF, arguments QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Call2", thisObject, arguments}).(*QScriptValue)
}

func (ptr *QScriptValue) Construct2(arguments QScriptValue_ITF) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Construct2", arguments}).(*QScriptValue)
}

func (ptr *QScriptValue) Data() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(*QScriptValue)
}

func (ptr *QScriptValue) Engine() *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QScriptEngine)
}

func (ptr *QScriptValue) Equals(other QScriptValue_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Equals", other}).(bool)
}

func (ptr *QScriptValue) InstanceOf(other QScriptValue_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstanceOf", other}).(bool)
}

func (ptr *QScriptValue) IsArray() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsArray"}).(bool)
}

func (ptr *QScriptValue) IsBool() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBool"}).(bool)
}

func (ptr *QScriptValue) IsDate() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDate"}).(bool)
}

func (ptr *QScriptValue) IsError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsError"}).(bool)
}

func (ptr *QScriptValue) IsFunction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFunction"}).(bool)
}

func (ptr *QScriptValue) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QScriptValue) IsNumber() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNumber"}).(bool)
}

func (ptr *QScriptValue) IsObject() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsObject"}).(bool)
}

func (ptr *QScriptValue) IsQMetaObject() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsQMetaObject"}).(bool)
}

func (ptr *QScriptValue) IsQObject() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsQObject"}).(bool)
}

func (ptr *QScriptValue) IsRegExp() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRegExp"}).(bool)
}

func (ptr *QScriptValue) IsString() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsString"}).(bool)
}

func (ptr *QScriptValue) IsUndefined() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUndefined"}).(bool)
}

func (ptr *QScriptValue) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QScriptValue) IsVariant() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVariant"}).(bool)
}

func (ptr *QScriptValue) LessThan(other QScriptValue_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LessThan", other}).(bool)
}

func (ptr *QScriptValue) Property(name string, mode QScriptValue__ResolveFlag) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", name, mode}).(*QScriptValue)
}

func (ptr *QScriptValue) Property2(arrayIndex uint, mode QScriptValue__ResolveFlag) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property2", arrayIndex, mode}).(*QScriptValue)
}

func (ptr *QScriptValue) Property3(name QScriptString_ITF, mode QScriptValue__ResolveFlag) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property3", name, mode}).(*QScriptValue)
}

func (ptr *QScriptValue) PropertyFlags(name string, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {

	return QScriptValue__PropertyFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyFlags", name, mode}).(float64))
}

func (ptr *QScriptValue) PropertyFlags2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {

	return QScriptValue__PropertyFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyFlags2", name, mode}).(float64))
}

func (ptr *QScriptValue) Prototype() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prototype"}).(*QScriptValue)
}

func (ptr *QScriptValue) ScriptClass() *QScriptClass {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScriptClass"}).(*QScriptClass)
}

func (ptr *QScriptValue) SetData(data QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", data})
}

func (ptr *QScriptValue) SetProperty(name string, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", name, value, flags})
}

func (ptr *QScriptValue) SetProperty2(arrayIndex uint, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty2", arrayIndex, value, flags})
}

func (ptr *QScriptValue) SetProperty3(name QScriptString_ITF, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty3", name, value, flags})
}

func (ptr *QScriptValue) SetPrototype(prototype QScriptValue_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrototype", prototype})
}

func (ptr *QScriptValue) SetScriptClass(scriptClass QScriptClass_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScriptClass", scriptClass})
}

func (ptr *QScriptValue) StrictlyEquals(other QScriptValue_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StrictlyEquals", other}).(bool)
}

func (ptr *QScriptValue) ToBool() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToBool"}).(bool)
}

func (ptr *QScriptValue) ToDateTime() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDateTime"}).(*core.QDateTime)
}

func (ptr *QScriptValue) ToInt32() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToInt32"}).(float64))
}

func (ptr *QScriptValue) ToQMetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToQMetaObject"}).(*core.QMetaObject)
}

func (ptr *QScriptValue) ToQObject() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToQObject"}).(*core.QObject)
}

func (ptr *QScriptValue) ToRegExp() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToRegExp"}).(*core.QRegExp)
}

func (ptr *QScriptValue) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QScriptValue) ToUInt16() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToUInt16"}).(float64))
}

func (ptr *QScriptValue) ToUInt32() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToUInt32"}).(float64))
}

func (ptr *QScriptValue) ToVariant() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToVariant"}).(*core.QVariant)
}

func (ptr *QScriptValue) DestroyQScriptValue() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptValue"})
}

type QScriptValueIterator struct {
	internal.Internal
}

type QScriptValueIterator_ITF interface {
	QScriptValueIterator_PTR() *QScriptValueIterator
}

func (ptr *QScriptValueIterator) QScriptValueIterator_PTR() *QScriptValueIterator {
	return ptr
}

func (ptr *QScriptValueIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptValueIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptValueIterator(ptr QScriptValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValueIterator_PTR().Pointer()
	}
	return nil
}

func (n *QScriptValueIterator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptValueIteratorFromPointer(ptr unsafe.Pointer) (n *QScriptValueIterator) {
	n = new(QScriptValueIterator)
	n.InitFromInternal(uintptr(ptr), "script.QScriptValueIterator")
	return
}

type QScriptable struct {
	internal.Internal
}

type QScriptable_ITF interface {
	QScriptable_PTR() *QScriptable
}

func (ptr *QScriptable) QScriptable_PTR() *QScriptable {
	return ptr
}

func (ptr *QScriptable) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScriptable) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScriptable(ptr QScriptable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptable_PTR().Pointer()
	}
	return nil
}

func (n *QScriptable) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScriptableFromPointer(ptr unsafe.Pointer) (n *QScriptable) {
	n = new(QScriptable)
	n.InitFromInternal(uintptr(ptr), "script.QScriptable")
	return
}

func (ptr *QScriptable) DestroyQScriptable() {
}

func (ptr *QScriptable) Argument(index int) *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Argument", index}).(*QScriptValue)
}

func (ptr *QScriptable) ArgumentCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ArgumentCount"}).(float64))
}

func (ptr *QScriptable) Context() *QScriptContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Context"}).(*QScriptContext)
}

func (ptr *QScriptable) Engine() *QScriptEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Engine"}).(*QScriptEngine)
}

func (ptr *QScriptable) ThisObject() *QScriptValue {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ThisObject"}).(*QScriptValue)
}

func init() {
	internal.ConstructorTable["script.QScriptClass"] = NewQScriptClassFromPointer
	internal.ConstructorTable["script.QScriptContext"] = NewQScriptContextFromPointer
	internal.ConstructorTable["script.QScriptContextInfo"] = NewQScriptContextInfoFromPointer
	internal.ConstructorTable["script.QScriptEngine"] = NewQScriptEngineFromPointer
	internal.ConstructorTable["script.QScriptEngineAgent"] = NewQScriptEngineAgentFromPointer
	internal.ConstructorTable["script.QScriptExtensionInterface"] = NewQScriptExtensionInterfaceFromPointer
	internal.ConstructorTable["script.QScriptExtensionPlugin"] = NewQScriptExtensionPluginFromPointer
	internal.ConstructorTable["script.QScriptProgram"] = NewQScriptProgramFromPointer
	internal.ConstructorTable["script.QScriptString"] = NewQScriptStringFromPointer
	internal.ConstructorTable["script.QScriptSyntaxCheckResult"] = NewQScriptSyntaxCheckResultFromPointer
	internal.ConstructorTable["script.QScriptValue"] = NewQScriptValueFromPointer
	internal.ConstructorTable["script.QScriptable"] = NewQScriptableFromPointer
}
