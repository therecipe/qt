// +build !minimal

package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

//QScriptClass::Extension
type QScriptClass__Extension int64

const (
	QScriptClass__Callable    = QScriptClass__Extension(0)
	QScriptClass__HasInstance = QScriptClass__Extension(1)
)

//QScriptClass::QueryFlag
type QScriptClass__QueryFlag int64

const (
	QScriptClass__HandlesReadAccess  = QScriptClass__QueryFlag(0x01)
	QScriptClass__HandlesWriteAccess = QScriptClass__QueryFlag(0x02)
)

type QScriptClass struct {
	ptr unsafe.Pointer
}

type QScriptClass_ITF interface {
	QScriptClass_PTR() *QScriptClass
}

func (p *QScriptClass) QScriptClass_PTR() *QScriptClass {
	return p
}

func (p *QScriptClass) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptClass) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptClass(ptr QScriptClass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClass_PTR().Pointer()
	}
	return nil
}

func NewQScriptClassFromPointer(ptr unsafe.Pointer) *QScriptClass {
	var n = new(QScriptClass)
	n.SetPointer(ptr)
	return n
}

func newQScriptClassFromPointer(ptr unsafe.Pointer) *QScriptClass {
	var n = NewQScriptClassFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QScriptClass_") {
		n.SetObjectNameAbs("QScriptClass_" + qt.Identifier())
	}
	return n
}

func NewQScriptClass(engine QScriptEngine_ITF) *QScriptClass {
	defer qt.Recovering("QScriptClass::QScriptClass")

	return newQScriptClassFromPointer(C.QScriptClass_NewQScriptClass(PointerFromQScriptEngine(engine)))
}

func (ptr *QScriptClass) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptClass::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptClass_Engine(ptr.Pointer()))
	}
	return nil
}

//export callbackQScriptClass_Extension
func callbackQScriptClass_Extension(ptr unsafe.Pointer, ptrName *C.char, extension C.int, argument unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QScriptClass::extension")

	if signal := qt.GetSignal(C.GoString(ptrName), "extension"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QScriptClass__Extension, *core.QVariant) *core.QVariant)(QScriptClass__Extension(extension), core.NewQVariantFromPointer(argument)))
	}

	return core.PointerFromQVariant(NewQScriptClassFromPointer(ptr).ExtensionDefault(QScriptClass__Extension(extension), core.NewQVariantFromPointer(argument)))
}

func (ptr *QScriptClass) ConnectExtension(f func(extension QScriptClass__Extension, argument *core.QVariant) *core.QVariant) {
	defer qt.Recovering("connect QScriptClass::extension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "extension", f)
	}
}

func (ptr *QScriptClass) DisconnectExtension() {
	defer qt.Recovering("disconnect QScriptClass::extension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "extension")
	}
}

func (ptr *QScriptClass) Extension(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptClass::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptClass_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

func (ptr *QScriptClass) ExtensionDefault(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptClass::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptClass_ExtensionDefault(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

//export callbackQScriptClass_Name
func callbackQScriptClass_Name(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QScriptClass::name")

	if signal := qt.GetSignal(C.GoString(ptrName), "name"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQScriptClassFromPointer(ptr).NameDefault())
}

func (ptr *QScriptClass) ConnectName(f func() string) {
	defer qt.Recovering("connect QScriptClass::name")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "name", f)
	}
}

func (ptr *QScriptClass) DisconnectName() {
	defer qt.Recovering("disconnect QScriptClass::name")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "name")
	}
}

func (ptr *QScriptClass) Name() string {
	defer qt.Recovering("QScriptClass::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptClass) NameDefault() string {
	defer qt.Recovering("QScriptClass::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_NameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQScriptClass_NewIterator
func callbackQScriptClass_NewIterator(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QScriptClass::newIterator")

	if signal := qt.GetSignal(C.GoString(ptrName), "newIterator"); signal != nil {
		return PointerFromQScriptClassPropertyIterator(signal.(func(*QScriptValue) *QScriptClassPropertyIterator)(NewQScriptValueFromPointer(object)))
	}

	return PointerFromQScriptClassPropertyIterator(NewQScriptClassFromPointer(ptr).NewIteratorDefault(NewQScriptValueFromPointer(object)))
}

func (ptr *QScriptClass) ConnectNewIterator(f func(object *QScriptValue) *QScriptClassPropertyIterator) {
	defer qt.Recovering("connect QScriptClass::newIterator")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "newIterator", f)
	}
}

func (ptr *QScriptClass) DisconnectNewIterator() {
	defer qt.Recovering("disconnect QScriptClass::newIterator")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "newIterator")
	}
}

func (ptr *QScriptClass) NewIterator(object QScriptValue_ITF) *QScriptClassPropertyIterator {
	defer qt.Recovering("QScriptClass::newIterator")

	if ptr.Pointer() != nil {
		return NewQScriptClassPropertyIteratorFromPointer(C.QScriptClass_NewIterator(ptr.Pointer(), PointerFromQScriptValue(object)))
	}
	return nil
}

func (ptr *QScriptClass) NewIteratorDefault(object QScriptValue_ITF) *QScriptClassPropertyIterator {
	defer qt.Recovering("QScriptClass::newIterator")

	if ptr.Pointer() != nil {
		return NewQScriptClassPropertyIteratorFromPointer(C.QScriptClass_NewIteratorDefault(ptr.Pointer(), PointerFromQScriptValue(object)))
	}
	return nil
}

//export callbackQScriptClass_Prototype
func callbackQScriptClass_Prototype(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScriptClass::prototype")

	if signal := qt.GetSignal(C.GoString(ptrName), "prototype"); signal != nil {
		return PointerFromQScriptValue(signal.(func() *QScriptValue)())
	}

	return PointerFromQScriptValue(NewQScriptClassFromPointer(ptr).PrototypeDefault())
}

func (ptr *QScriptClass) ConnectPrototype(f func() *QScriptValue) {
	defer qt.Recovering("connect QScriptClass::prototype")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "prototype", f)
	}
}

func (ptr *QScriptClass) DisconnectPrototype() {
	defer qt.Recovering("disconnect QScriptClass::prototype")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "prototype")
	}
}

func (ptr *QScriptClass) Prototype() *QScriptValue {
	defer qt.Recovering("QScriptClass::prototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptClass_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptClass) PrototypeDefault() *QScriptValue {
	defer qt.Recovering("QScriptClass::prototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptClass_PrototypeDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQScriptClass_SupportsExtension
func callbackQScriptClass_SupportsExtension(ptr unsafe.Pointer, ptrName *C.char, extension C.int) C.int {
	defer qt.Recovering("callback QScriptClass::supportsExtension")

	if signal := qt.GetSignal(C.GoString(ptrName), "supportsExtension"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(QScriptClass__Extension) bool)(QScriptClass__Extension(extension))))
	}

	return C.int(qt.GoBoolToInt(NewQScriptClassFromPointer(ptr).SupportsExtensionDefault(QScriptClass__Extension(extension))))
}

func (ptr *QScriptClass) ConnectSupportsExtension(f func(extension QScriptClass__Extension) bool) {
	defer qt.Recovering("connect QScriptClass::supportsExtension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "supportsExtension", f)
	}
}

func (ptr *QScriptClass) DisconnectSupportsExtension() {
	defer qt.Recovering("disconnect QScriptClass::supportsExtension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "supportsExtension")
	}
}

func (ptr *QScriptClass) SupportsExtension(extension QScriptClass__Extension) bool {
	defer qt.Recovering("QScriptClass::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptClass_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptClass) SupportsExtensionDefault(extension QScriptClass__Extension) bool {
	defer qt.Recovering("QScriptClass::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptClass_SupportsExtensionDefault(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptClass) DestroyQScriptClass() {
	defer qt.Recovering("QScriptClass::~QScriptClass")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QScriptClass_DestroyQScriptClass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptClass) ObjectNameAbs() string {
	defer qt.Recovering("QScriptClass::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptClass) SetObjectNameAbs(name string) {
	defer qt.Recovering("QScriptClass::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QScriptClass_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QScriptClassPropertyIterator struct {
	ptr unsafe.Pointer
}

type QScriptClassPropertyIterator_ITF interface {
	QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator
}

func (p *QScriptClassPropertyIterator) QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator {
	return p
}

func (p *QScriptClassPropertyIterator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptClassPropertyIterator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptClassPropertyIterator(ptr QScriptClassPropertyIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClassPropertyIterator_PTR().Pointer()
	}
	return nil
}

func NewQScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) *QScriptClassPropertyIterator {
	var n = new(QScriptClassPropertyIterator)
	n.SetPointer(ptr)
	return n
}

func newQScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) *QScriptClassPropertyIterator {
	var n = NewQScriptClassPropertyIteratorFromPointer(ptr)
	return n
}

//QScriptContext::Error
type QScriptContext__Error int64

const (
	QScriptContext__UnknownError   = QScriptContext__Error(0)
	QScriptContext__ReferenceError = QScriptContext__Error(1)
	QScriptContext__SyntaxError    = QScriptContext__Error(2)
	QScriptContext__TypeError      = QScriptContext__Error(3)
	QScriptContext__RangeError     = QScriptContext__Error(4)
	QScriptContext__URIError       = QScriptContext__Error(5)
)

//QScriptContext::ExecutionState
type QScriptContext__ExecutionState int64

const (
	QScriptContext__NormalState    = QScriptContext__ExecutionState(0)
	QScriptContext__ExceptionState = QScriptContext__ExecutionState(1)
)

type QScriptContext struct {
	ptr unsafe.Pointer
}

type QScriptContext_ITF interface {
	QScriptContext_PTR() *QScriptContext
}

func (p *QScriptContext) QScriptContext_PTR() *QScriptContext {
	return p
}

func (p *QScriptContext) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptContext) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptContext(ptr QScriptContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContext_PTR().Pointer()
	}
	return nil
}

func NewQScriptContextFromPointer(ptr unsafe.Pointer) *QScriptContext {
	var n = new(QScriptContext)
	n.SetPointer(ptr)
	return n
}

func newQScriptContextFromPointer(ptr unsafe.Pointer) *QScriptContext {
	var n = NewQScriptContextFromPointer(ptr)
	return n
}

func (ptr *QScriptContext) ActivationObject() *QScriptValue {
	defer qt.Recovering("QScriptContext::activationObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_ActivationObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) Argument(index int) *QScriptValue {
	defer qt.Recovering("QScriptContext::argument")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_Argument(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QScriptContext) ArgumentCount() int {
	defer qt.Recovering("QScriptContext::argumentCount")

	if ptr.Pointer() != nil {
		return int(C.QScriptContext_ArgumentCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContext) ArgumentsObject() *QScriptValue {
	defer qt.Recovering("QScriptContext::argumentsObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_ArgumentsObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) Backtrace() []string {
	defer qt.Recovering("QScriptContext::backtrace")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptContext_Backtrace(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptContext) Callee() *QScriptValue {
	defer qt.Recovering("QScriptContext::callee")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_Callee(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptContext::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptContext_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) IsCalledAsConstructor() bool {
	defer qt.Recovering("QScriptContext::isCalledAsConstructor")

	if ptr.Pointer() != nil {
		return C.QScriptContext_IsCalledAsConstructor(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptContext) ParentContext() *QScriptContext {
	defer qt.Recovering("QScriptContext::parentContext")

	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptContext_ParentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) SetActivationObject(activation QScriptValue_ITF) {
	defer qt.Recovering("QScriptContext::setActivationObject")

	if ptr.Pointer() != nil {
		C.QScriptContext_SetActivationObject(ptr.Pointer(), PointerFromQScriptValue(activation))
	}
}

func (ptr *QScriptContext) SetThisObject(thisObject QScriptValue_ITF) {
	defer qt.Recovering("QScriptContext::setThisObject")

	if ptr.Pointer() != nil {
		C.QScriptContext_SetThisObject(ptr.Pointer(), PointerFromQScriptValue(thisObject))
	}
}

func (ptr *QScriptContext) State() QScriptContext__ExecutionState {
	defer qt.Recovering("QScriptContext::state")

	if ptr.Pointer() != nil {
		return QScriptContext__ExecutionState(C.QScriptContext_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContext) ThisObject() *QScriptValue {
	defer qt.Recovering("QScriptContext::thisObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_ThisObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) ThrowError(error QScriptContext__Error, text string) *QScriptValue {
	defer qt.Recovering("QScriptContext::throwError")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_ThrowError(ptr.Pointer(), C.int(error), C.CString(text)))
	}
	return nil
}

func (ptr *QScriptContext) ThrowError2(text string) *QScriptValue {
	defer qt.Recovering("QScriptContext::throwError")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_ThrowError2(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QScriptContext) ThrowValue(value QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptContext::throwValue")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptContext_ThrowValue(ptr.Pointer(), PointerFromQScriptValue(value)))
	}
	return nil
}

func (ptr *QScriptContext) ToString() string {
	defer qt.Recovering("QScriptContext::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContext_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContext) DestroyQScriptContext() {
	defer qt.Recovering("QScriptContext::~QScriptContext")

	if ptr.Pointer() != nil {
		C.QScriptContext_DestroyQScriptContext(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QScriptContextInfo::FunctionType
type QScriptContextInfo__FunctionType int64

const (
	QScriptContextInfo__ScriptFunction     = QScriptContextInfo__FunctionType(0)
	QScriptContextInfo__QtFunction         = QScriptContextInfo__FunctionType(1)
	QScriptContextInfo__QtPropertyFunction = QScriptContextInfo__FunctionType(2)
	QScriptContextInfo__NativeFunction     = QScriptContextInfo__FunctionType(3)
)

type QScriptContextInfo struct {
	ptr unsafe.Pointer
}

type QScriptContextInfo_ITF interface {
	QScriptContextInfo_PTR() *QScriptContextInfo
}

func (p *QScriptContextInfo) QScriptContextInfo_PTR() *QScriptContextInfo {
	return p
}

func (p *QScriptContextInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptContextInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptContextInfo(ptr QScriptContextInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContextInfo_PTR().Pointer()
	}
	return nil
}

func NewQScriptContextInfoFromPointer(ptr unsafe.Pointer) *QScriptContextInfo {
	var n = new(QScriptContextInfo)
	n.SetPointer(ptr)
	return n
}

func newQScriptContextInfoFromPointer(ptr unsafe.Pointer) *QScriptContextInfo {
	var n = NewQScriptContextInfoFromPointer(ptr)
	return n
}

func NewQScriptContextInfo3() *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return newQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo3())
}

func NewQScriptContextInfo(context QScriptContext_ITF) *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return newQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo(PointerFromQScriptContext(context)))
}

func NewQScriptContextInfo2(other QScriptContextInfo_ITF) *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return newQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo2(PointerFromQScriptContextInfo(other)))
}

func (ptr *QScriptContextInfo) FileName() string {
	defer qt.Recovering("QScriptContextInfo::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContextInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionEndLineNumber() int {
	defer qt.Recovering("QScriptContextInfo::functionEndLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionEndLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionMetaIndex() int {
	defer qt.Recovering("QScriptContextInfo::functionMetaIndex")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionMetaIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionName() string {
	defer qt.Recovering("QScriptContextInfo::functionName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContextInfo_FunctionName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionParameterNames() []string {
	defer qt.Recovering("QScriptContextInfo::functionParameterNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptContextInfo_FunctionParameterNames(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptContextInfo) FunctionStartLineNumber() int {
	defer qt.Recovering("QScriptContextInfo::functionStartLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionStartLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionType() QScriptContextInfo__FunctionType {
	defer qt.Recovering("QScriptContextInfo::functionType")

	if ptr.Pointer() != nil {
		return QScriptContextInfo__FunctionType(C.QScriptContextInfo_FunctionType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) IsNull() bool {
	defer qt.Recovering("QScriptContextInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QScriptContextInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptContextInfo) LineNumber() int {
	defer qt.Recovering("QScriptContextInfo::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) ScriptId() int64 {
	defer qt.Recovering("QScriptContextInfo::scriptId")

	if ptr.Pointer() != nil {
		return int64(C.QScriptContextInfo_ScriptId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) DestroyQScriptContextInfo() {
	defer qt.Recovering("QScriptContextInfo::~QScriptContextInfo")

	if ptr.Pointer() != nil {
		C.QScriptContextInfo_DestroyQScriptContextInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QScriptEngine::QObjectWrapOption
type QScriptEngine__QObjectWrapOption int64

const (
	QScriptEngine__ExcludeChildObjects         = QScriptEngine__QObjectWrapOption(0x0001)
	QScriptEngine__ExcludeSuperClassMethods    = QScriptEngine__QObjectWrapOption(0x0002)
	QScriptEngine__ExcludeSuperClassProperties = QScriptEngine__QObjectWrapOption(0x0004)
	QScriptEngine__ExcludeSuperClassContents   = QScriptEngine__QObjectWrapOption(0x0006)
	QScriptEngine__SkipMethodsInEnumeration    = QScriptEngine__QObjectWrapOption(0x0008)
	QScriptEngine__ExcludeDeleteLater          = QScriptEngine__QObjectWrapOption(0x0010)
	QScriptEngine__ExcludeSlots                = QScriptEngine__QObjectWrapOption(0x0020)
	QScriptEngine__AutoCreateDynamicProperties = QScriptEngine__QObjectWrapOption(0x0100)
	QScriptEngine__PreferExistingWrapperObject = QScriptEngine__QObjectWrapOption(0x0200)
)

//QScriptEngine::ValueOwnership
type QScriptEngine__ValueOwnership int64

const (
	QScriptEngine__QtOwnership     = QScriptEngine__ValueOwnership(0)
	QScriptEngine__ScriptOwnership = QScriptEngine__ValueOwnership(1)
	QScriptEngine__AutoOwnership   = QScriptEngine__ValueOwnership(2)
)

type QScriptEngine struct {
	core.QObject
}

type QScriptEngine_ITF interface {
	core.QObject_ITF
	QScriptEngine_PTR() *QScriptEngine
}

func (p *QScriptEngine) QScriptEngine_PTR() *QScriptEngine {
	return p
}

func (p *QScriptEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QScriptEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQScriptEngine(ptr QScriptEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngine_PTR().Pointer()
	}
	return nil
}

func NewQScriptEngineFromPointer(ptr unsafe.Pointer) *QScriptEngine {
	var n = new(QScriptEngine)
	n.SetPointer(ptr)
	return n
}

func newQScriptEngineFromPointer(ptr unsafe.Pointer) *QScriptEngine {
	var n = NewQScriptEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QScriptEngine_") {
		n.SetObjectName("QScriptEngine_" + qt.Identifier())
	}
	return n
}

func NewQScriptEngine() *QScriptEngine {
	defer qt.Recovering("QScriptEngine::QScriptEngine")

	return newQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine())
}

func NewQScriptEngine2(parent core.QObject_ITF) *QScriptEngine {
	defer qt.Recovering("QScriptEngine::QScriptEngine")

	return newQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine2(core.PointerFromQObject(parent)))
}

func (ptr *QScriptEngine) AbortEvaluation(result QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngine::abortEvaluation")

	if ptr.Pointer() != nil {
		C.QScriptEngine_AbortEvaluation(ptr.Pointer(), PointerFromQScriptValue(result))
	}
}

func (ptr *QScriptEngine) Agent() *QScriptEngineAgent {
	defer qt.Recovering("QScriptEngine::agent")

	if ptr.Pointer() != nil {
		return NewQScriptEngineAgentFromPointer(C.QScriptEngine_Agent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) AvailableExtensions() []string {
	defer qt.Recovering("QScriptEngine::availableExtensions")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_AvailableExtensions(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func QScriptEngine_CheckSyntax(program string) *QScriptSyntaxCheckResult {
	defer qt.Recovering("QScriptEngine::checkSyntax")

	return NewQScriptSyntaxCheckResultFromPointer(C.QScriptEngine_QScriptEngine_CheckSyntax(C.CString(program)))
}

func (ptr *QScriptEngine) CheckSyntax(program string) *QScriptSyntaxCheckResult {
	defer qt.Recovering("QScriptEngine::checkSyntax")

	return NewQScriptSyntaxCheckResultFromPointer(C.QScriptEngine_QScriptEngine_CheckSyntax(C.CString(program)))
}

func (ptr *QScriptEngine) ClearExceptions() {
	defer qt.Recovering("QScriptEngine::clearExceptions")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ClearExceptions(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) CollectGarbage() {
	defer qt.Recovering("QScriptEngine::collectGarbage")

	if ptr.Pointer() != nil {
		C.QScriptEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) CurrentContext() *QScriptContext {
	defer qt.Recovering("QScriptEngine::currentContext")

	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_CurrentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) DefaultPrototype(metaTypeId int) *QScriptValue {
	defer qt.Recovering("QScriptEngine::defaultPrototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_DefaultPrototype(ptr.Pointer(), C.int(metaTypeId)))
	}
	return nil
}

func (ptr *QScriptEngine) Evaluate2(program QScriptProgram_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::evaluate")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_Evaluate2(ptr.Pointer(), PointerFromQScriptProgram(program)))
	}
	return nil
}

func (ptr *QScriptEngine) Evaluate(program string, fileName string, lineNumber int) *QScriptValue {
	defer qt.Recovering("QScriptEngine::evaluate")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_Evaluate(ptr.Pointer(), C.CString(program), C.CString(fileName), C.int(lineNumber)))
	}
	return nil
}

func (ptr *QScriptEngine) GlobalObject() *QScriptValue {
	defer qt.Recovering("QScriptEngine::globalObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_GlobalObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) HasUncaughtException() bool {
	defer qt.Recovering("QScriptEngine::hasUncaughtException")

	if ptr.Pointer() != nil {
		return C.QScriptEngine_HasUncaughtException(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngine) ImportExtension(extension string) *QScriptValue {
	defer qt.Recovering("QScriptEngine::importExtension")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_ImportExtension(ptr.Pointer(), C.CString(extension)))
	}
	return nil
}

func (ptr *QScriptEngine) ImportedExtensions() []string {
	defer qt.Recovering("QScriptEngine::importedExtensions")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_ImportedExtensions(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) InstallTranslatorFunctions(object QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngine::installTranslatorFunctions")

	if ptr.Pointer() != nil {
		C.QScriptEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) IsEvaluating() bool {
	defer qt.Recovering("QScriptEngine::isEvaluating")

	if ptr.Pointer() != nil {
		return C.QScriptEngine_IsEvaluating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngine) NewDate2(value core.QDateTime_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newDate")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewDate2(ptr.Pointer(), core.PointerFromQDateTime(value)))
	}
	return nil
}

func (ptr *QScriptEngine) NewObject() *QScriptValue {
	defer qt.Recovering("QScriptEngine::newObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) NewObject2(scriptClass QScriptClass_ITF, data QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewObject2(ptr.Pointer(), PointerFromQScriptClass(scriptClass), PointerFromQScriptValue(data)))
	}
	return nil
}

func (ptr *QScriptEngine) NewQMetaObject(metaObject core.QMetaObject_ITF, ctor QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newQMetaObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewQMetaObject(ptr.Pointer(), core.PointerFromQMetaObject(metaObject), PointerFromQScriptValue(ctor)))
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject(object core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newQObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object), C.int(ownership), C.int(options)))
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject2(scriptObject QScriptValue_ITF, qtObject core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newQObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewQObject2(ptr.Pointer(), PointerFromQScriptValue(scriptObject), core.PointerFromQObject(qtObject), C.int(ownership), C.int(options)))
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp(regexp core.QRegExp_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newRegExp")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp(ptr.Pointer(), core.PointerFromQRegExp(regexp)))
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp2(pattern string, flags string) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newRegExp")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp2(ptr.Pointer(), C.CString(pattern), C.CString(flags)))
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant2(object QScriptValue_ITF, value core.QVariant_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newVariant")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewVariant2(ptr.Pointer(), PointerFromQScriptValue(object), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant(value core.QVariant_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::newVariant")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewVariant(ptr.Pointer(), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QScriptEngine) NullValue() *QScriptValue {
	defer qt.Recovering("QScriptEngine::nullValue")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NullValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) PopContext() {
	defer qt.Recovering("QScriptEngine::popContext")

	if ptr.Pointer() != nil {
		C.QScriptEngine_PopContext(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) ProcessEventsInterval() int {
	defer qt.Recovering("QScriptEngine::processEventsInterval")

	if ptr.Pointer() != nil {
		return int(C.QScriptEngine_ProcessEventsInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptEngine) PushContext() *QScriptContext {
	defer qt.Recovering("QScriptEngine::pushContext")

	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_PushContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) ReportAdditionalMemoryCost(size int) {
	defer qt.Recovering("QScriptEngine::reportAdditionalMemoryCost")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ReportAdditionalMemoryCost(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QScriptEngine) SetAgent(agent QScriptEngineAgent_ITF) {
	defer qt.Recovering("QScriptEngine::setAgent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_SetAgent(ptr.Pointer(), PointerFromQScriptEngineAgent(agent))
	}
}

func (ptr *QScriptEngine) SetDefaultPrototype(metaTypeId int, prototype QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngine::setDefaultPrototype")

	if ptr.Pointer() != nil {
		C.QScriptEngine_SetDefaultPrototype(ptr.Pointer(), C.int(metaTypeId), PointerFromQScriptValue(prototype))
	}
}

func (ptr *QScriptEngine) SetGlobalObject(object QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngine::setGlobalObject")

	if ptr.Pointer() != nil {
		C.QScriptEngine_SetGlobalObject(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) SetProcessEventsInterval(interval int) {
	defer qt.Recovering("QScriptEngine::setProcessEventsInterval")

	if ptr.Pointer() != nil {
		C.QScriptEngine_SetProcessEventsInterval(ptr.Pointer(), C.int(interval))
	}
}

//export callbackQScriptEngine_SignalHandlerException
func callbackQScriptEngine_SignalHandlerException(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::signalHandlerException")

	if signal := qt.GetSignal(C.GoString(ptrName), "signalHandlerException"); signal != nil {
		signal.(func(*QScriptValue))(NewQScriptValueFromPointer(exception))
	}

}

func (ptr *QScriptEngine) ConnectSignalHandlerException(f func(exception *QScriptValue)) {
	defer qt.Recovering("connect QScriptEngine::signalHandlerException")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ConnectSignalHandlerException(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalHandlerException", f)
	}
}

func (ptr *QScriptEngine) DisconnectSignalHandlerException() {
	defer qt.Recovering("disconnect QScriptEngine::signalHandlerException")

	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectSignalHandlerException(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalHandlerException")
	}
}

func (ptr *QScriptEngine) SignalHandlerException(exception QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngine::signalHandlerException")

	if ptr.Pointer() != nil {
		C.QScriptEngine_SignalHandlerException(ptr.Pointer(), PointerFromQScriptValue(exception))
	}
}

func (ptr *QScriptEngine) ToObject(value QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptEngine::toObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_ToObject(ptr.Pointer(), PointerFromQScriptValue(value)))
	}
	return nil
}

func (ptr *QScriptEngine) ToStringHandle(str string) *QScriptString {
	defer qt.Recovering("QScriptEngine::toStringHandle")

	if ptr.Pointer() != nil {
		return NewQScriptStringFromPointer(C.QScriptEngine_ToStringHandle(ptr.Pointer(), C.CString(str)))
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtException() *QScriptValue {
	defer qt.Recovering("QScriptEngine::uncaughtException")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_UncaughtException(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtExceptionBacktrace() []string {
	defer qt.Recovering("QScriptEngine::uncaughtExceptionBacktrace")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_UncaughtExceptionBacktrace(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) UncaughtExceptionLineNumber() int {
	defer qt.Recovering("QScriptEngine::uncaughtExceptionLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptEngine_UncaughtExceptionLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptEngine) UndefinedValue() *QScriptValue {
	defer qt.Recovering("QScriptEngine::undefinedValue")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_UndefinedValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) DestroyQScriptEngine() {
	defer qt.Recovering("QScriptEngine::~QScriptEngine")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScriptEngine_DestroyQScriptEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScriptEngine_TimerEvent
func callbackQScriptEngine_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScriptEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScriptEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScriptEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScriptEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScriptEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScriptEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScriptEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScriptEngine_ChildEvent
func callbackQScriptEngine_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScriptEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScriptEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScriptEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScriptEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScriptEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScriptEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScriptEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScriptEngine_ConnectNotify
func callbackQScriptEngine_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScriptEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScriptEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScriptEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScriptEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngine_CustomEvent
func callbackQScriptEngine_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScriptEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScriptEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScriptEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScriptEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScriptEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScriptEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScriptEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScriptEngine_DeleteLater
func callbackQScriptEngine_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngine::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScriptEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScriptEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScriptEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScriptEngine) DeleteLater() {
	defer qt.Recovering("QScriptEngine::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScriptEngine_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngine) DeleteLaterDefault() {
	defer qt.Recovering("QScriptEngine::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScriptEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScriptEngine_DisconnectNotify
func callbackQScriptEngine_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScriptEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScriptEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScriptEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScriptEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngine_Event
func callbackQScriptEngine_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScriptEngine::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScriptEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScriptEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScriptEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScriptEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QScriptEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScriptEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngine::event")

	if ptr.Pointer() != nil {
		return C.QScriptEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScriptEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngine::event")

	if ptr.Pointer() != nil {
		return C.QScriptEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScriptEngine_EventFilter
func callbackQScriptEngine_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScriptEngine::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScriptEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScriptEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScriptEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScriptEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScriptEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScriptEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScriptEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScriptEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScriptEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScriptEngine_MetaObject
func callbackQScriptEngine_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScriptEngine::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScriptEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScriptEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScriptEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScriptEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScriptEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScriptEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScriptEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QScriptEngineAgent::Extension
type QScriptEngineAgent__Extension int64

const (
	QScriptEngineAgent__DebuggerInvocationRequest = QScriptEngineAgent__Extension(0)
)

type QScriptEngineAgent struct {
	ptr unsafe.Pointer
}

type QScriptEngineAgent_ITF interface {
	QScriptEngineAgent_PTR() *QScriptEngineAgent
}

func (p *QScriptEngineAgent) QScriptEngineAgent_PTR() *QScriptEngineAgent {
	return p
}

func (p *QScriptEngineAgent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptEngineAgent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptEngineAgent(ptr QScriptEngineAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngineAgent_PTR().Pointer()
	}
	return nil
}

func NewQScriptEngineAgentFromPointer(ptr unsafe.Pointer) *QScriptEngineAgent {
	var n = new(QScriptEngineAgent)
	n.SetPointer(ptr)
	return n
}

func newQScriptEngineAgentFromPointer(ptr unsafe.Pointer) *QScriptEngineAgent {
	var n = NewQScriptEngineAgentFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QScriptEngineAgent_") {
		n.SetObjectNameAbs("QScriptEngineAgent_" + qt.Identifier())
	}
	return n
}

func NewQScriptEngineAgent(engine QScriptEngine_ITF) *QScriptEngineAgent {
	defer qt.Recovering("QScriptEngineAgent::QScriptEngineAgent")

	return newQScriptEngineAgentFromPointer(C.QScriptEngineAgent_NewQScriptEngineAgent(PointerFromQScriptEngine(engine)))
}

//export callbackQScriptEngineAgent_ContextPop
func callbackQScriptEngineAgent_ContextPop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngineAgent::contextPop")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextPop"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPopDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectContextPop(f func()) {
	defer qt.Recovering("connect QScriptEngineAgent::contextPop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "contextPop", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectContextPop() {
	defer qt.Recovering("disconnect QScriptEngineAgent::contextPop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "contextPop")
	}
}

func (ptr *QScriptEngineAgent) ContextPop() {
	defer qt.Recovering("QScriptEngineAgent::contextPop")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPop(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) ContextPopDefault() {
	defer qt.Recovering("QScriptEngineAgent::contextPop")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPopDefault(ptr.Pointer())
	}
}

//export callbackQScriptEngineAgent_ContextPush
func callbackQScriptEngineAgent_ContextPush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngineAgent::contextPush")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextPush"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPushDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectContextPush(f func()) {
	defer qt.Recovering("connect QScriptEngineAgent::contextPush")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "contextPush", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectContextPush() {
	defer qt.Recovering("disconnect QScriptEngineAgent::contextPush")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "contextPush")
	}
}

func (ptr *QScriptEngineAgent) ContextPush() {
	defer qt.Recovering("QScriptEngineAgent::contextPush")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPush(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) ContextPushDefault() {
	defer qt.Recovering("QScriptEngineAgent::contextPush")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPushDefault(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptEngineAgent::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptEngineAgent_Engine(ptr.Pointer()))
	}
	return nil
}

//export callbackQScriptEngineAgent_ExceptionCatch
func callbackQScriptEngineAgent_ExceptionCatch(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, exception unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineAgent::exceptionCatch")

	if signal := qt.GetSignal(C.GoString(ptrName), "exceptionCatch"); signal != nil {
		signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(exception))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionCatchDefault(int64(scriptId), NewQScriptValueFromPointer(exception))
	}
}

func (ptr *QScriptEngineAgent) ConnectExceptionCatch(f func(scriptId int64, exception *QScriptValue)) {
	defer qt.Recovering("connect QScriptEngineAgent::exceptionCatch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "exceptionCatch", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectExceptionCatch() {
	defer qt.Recovering("disconnect QScriptEngineAgent::exceptionCatch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "exceptionCatch")
	}
}

func (ptr *QScriptEngineAgent) ExceptionCatch(scriptId int64, exception QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngineAgent::exceptionCatch")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionCatch(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception))
	}
}

func (ptr *QScriptEngineAgent) ExceptionCatchDefault(scriptId int64, exception QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngineAgent::exceptionCatch")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionCatchDefault(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception))
	}
}

//export callbackQScriptEngineAgent_ExceptionThrow
func callbackQScriptEngineAgent_ExceptionThrow(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, exception unsafe.Pointer, hasHandler C.int) {
	defer qt.Recovering("callback QScriptEngineAgent::exceptionThrow")

	if signal := qt.GetSignal(C.GoString(ptrName), "exceptionThrow"); signal != nil {
		signal.(func(int64, *QScriptValue, bool))(int64(scriptId), NewQScriptValueFromPointer(exception), int(hasHandler) != 0)
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionThrowDefault(int64(scriptId), NewQScriptValueFromPointer(exception), int(hasHandler) != 0)
	}
}

func (ptr *QScriptEngineAgent) ConnectExceptionThrow(f func(scriptId int64, exception *QScriptValue, hasHandler bool)) {
	defer qt.Recovering("connect QScriptEngineAgent::exceptionThrow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "exceptionThrow", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectExceptionThrow() {
	defer qt.Recovering("disconnect QScriptEngineAgent::exceptionThrow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "exceptionThrow")
	}
}

func (ptr *QScriptEngineAgent) ExceptionThrow(scriptId int64, exception QScriptValue_ITF, hasHandler bool) {
	defer qt.Recovering("QScriptEngineAgent::exceptionThrow")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionThrow(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception), C.int(qt.GoBoolToInt(hasHandler)))
	}
}

func (ptr *QScriptEngineAgent) ExceptionThrowDefault(scriptId int64, exception QScriptValue_ITF, hasHandler bool) {
	defer qt.Recovering("QScriptEngineAgent::exceptionThrow")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionThrowDefault(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception), C.int(qt.GoBoolToInt(hasHandler)))
	}
}

//export callbackQScriptEngineAgent_Extension
func callbackQScriptEngineAgent_Extension(ptr unsafe.Pointer, ptrName *C.char, extension C.int, argument unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QScriptEngineAgent::extension")

	if signal := qt.GetSignal(C.GoString(ptrName), "extension"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QScriptEngineAgent__Extension, *core.QVariant) *core.QVariant)(QScriptEngineAgent__Extension(extension), core.NewQVariantFromPointer(argument)))
	}

	return core.PointerFromQVariant(NewQScriptEngineAgentFromPointer(ptr).ExtensionDefault(QScriptEngineAgent__Extension(extension), core.NewQVariantFromPointer(argument)))
}

func (ptr *QScriptEngineAgent) ConnectExtension(f func(extension QScriptEngineAgent__Extension, argument *core.QVariant) *core.QVariant) {
	defer qt.Recovering("connect QScriptEngineAgent::extension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "extension", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectExtension() {
	defer qt.Recovering("disconnect QScriptEngineAgent::extension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "extension")
	}
}

func (ptr *QScriptEngineAgent) Extension(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptEngineAgent::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptEngineAgent_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

func (ptr *QScriptEngineAgent) ExtensionDefault(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptEngineAgent::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptEngineAgent_ExtensionDefault(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

//export callbackQScriptEngineAgent_FunctionEntry
func callbackQScriptEngineAgent_FunctionEntry(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong) {
	defer qt.Recovering("callback QScriptEngineAgent::functionEntry")

	if signal := qt.GetSignal(C.GoString(ptrName), "functionEntry"); signal != nil {
		signal.(func(int64))(int64(scriptId))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionEntryDefault(int64(scriptId))
	}
}

func (ptr *QScriptEngineAgent) ConnectFunctionEntry(f func(scriptId int64)) {
	defer qt.Recovering("connect QScriptEngineAgent::functionEntry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "functionEntry", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectFunctionEntry() {
	defer qt.Recovering("disconnect QScriptEngineAgent::functionEntry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "functionEntry")
	}
}

func (ptr *QScriptEngineAgent) FunctionEntry(scriptId int64) {
	defer qt.Recovering("QScriptEngineAgent::functionEntry")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionEntry(ptr.Pointer(), C.longlong(scriptId))
	}
}

func (ptr *QScriptEngineAgent) FunctionEntryDefault(scriptId int64) {
	defer qt.Recovering("QScriptEngineAgent::functionEntry")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionEntryDefault(ptr.Pointer(), C.longlong(scriptId))
	}
}

//export callbackQScriptEngineAgent_FunctionExit
func callbackQScriptEngineAgent_FunctionExit(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, returnValue unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineAgent::functionExit")

	if signal := qt.GetSignal(C.GoString(ptrName), "functionExit"); signal != nil {
		signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionExitDefault(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	}
}

func (ptr *QScriptEngineAgent) ConnectFunctionExit(f func(scriptId int64, returnValue *QScriptValue)) {
	defer qt.Recovering("connect QScriptEngineAgent::functionExit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "functionExit", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectFunctionExit() {
	defer qt.Recovering("disconnect QScriptEngineAgent::functionExit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "functionExit")
	}
}

func (ptr *QScriptEngineAgent) FunctionExit(scriptId int64, returnValue QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngineAgent::functionExit")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionExit(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(returnValue))
	}
}

func (ptr *QScriptEngineAgent) FunctionExitDefault(scriptId int64, returnValue QScriptValue_ITF) {
	defer qt.Recovering("QScriptEngineAgent::functionExit")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionExitDefault(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(returnValue))
	}
}

//export callbackQScriptEngineAgent_PositionChange
func callbackQScriptEngineAgent_PositionChange(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, lineNumber C.int, columnNumber C.int) {
	defer qt.Recovering("callback QScriptEngineAgent::positionChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChange"); signal != nil {
		signal.(func(int64, int, int))(int64(scriptId), int(lineNumber), int(columnNumber))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).PositionChangeDefault(int64(scriptId), int(lineNumber), int(columnNumber))
	}
}

func (ptr *QScriptEngineAgent) ConnectPositionChange(f func(scriptId int64, lineNumber int, columnNumber int)) {
	defer qt.Recovering("connect QScriptEngineAgent::positionChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "positionChange", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectPositionChange() {
	defer qt.Recovering("disconnect QScriptEngineAgent::positionChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "positionChange")
	}
}

func (ptr *QScriptEngineAgent) PositionChange(scriptId int64, lineNumber int, columnNumber int) {
	defer qt.Recovering("QScriptEngineAgent::positionChange")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_PositionChange(ptr.Pointer(), C.longlong(scriptId), C.int(lineNumber), C.int(columnNumber))
	}
}

func (ptr *QScriptEngineAgent) PositionChangeDefault(scriptId int64, lineNumber int, columnNumber int) {
	defer qt.Recovering("QScriptEngineAgent::positionChange")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_PositionChangeDefault(ptr.Pointer(), C.longlong(scriptId), C.int(lineNumber), C.int(columnNumber))
	}
}

//export callbackQScriptEngineAgent_ScriptLoad
func callbackQScriptEngineAgent_ScriptLoad(ptr unsafe.Pointer, ptrName *C.char, id C.longlong, program *C.char, fileName *C.char, baseLineNumber C.int) {
	defer qt.Recovering("callback QScriptEngineAgent::scriptLoad")

	if signal := qt.GetSignal(C.GoString(ptrName), "scriptLoad"); signal != nil {
		signal.(func(int64, string, string, int))(int64(id), C.GoString(program), C.GoString(fileName), int(baseLineNumber))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptLoadDefault(int64(id), C.GoString(program), C.GoString(fileName), int(baseLineNumber))
	}
}

func (ptr *QScriptEngineAgent) ConnectScriptLoad(f func(id int64, program string, fileName string, baseLineNumber int)) {
	defer qt.Recovering("connect QScriptEngineAgent::scriptLoad")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "scriptLoad", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectScriptLoad() {
	defer qt.Recovering("disconnect QScriptEngineAgent::scriptLoad")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "scriptLoad")
	}
}

func (ptr *QScriptEngineAgent) ScriptLoad(id int64, program string, fileName string, baseLineNumber int) {
	defer qt.Recovering("QScriptEngineAgent::scriptLoad")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ScriptLoad(ptr.Pointer(), C.longlong(id), C.CString(program), C.CString(fileName), C.int(baseLineNumber))
	}
}

func (ptr *QScriptEngineAgent) ScriptLoadDefault(id int64, program string, fileName string, baseLineNumber int) {
	defer qt.Recovering("QScriptEngineAgent::scriptLoad")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ScriptLoadDefault(ptr.Pointer(), C.longlong(id), C.CString(program), C.CString(fileName), C.int(baseLineNumber))
	}
}

//export callbackQScriptEngineAgent_ScriptUnload
func callbackQScriptEngineAgent_ScriptUnload(ptr unsafe.Pointer, ptrName *C.char, id C.longlong) {
	defer qt.Recovering("callback QScriptEngineAgent::scriptUnload")

	if signal := qt.GetSignal(C.GoString(ptrName), "scriptUnload"); signal != nil {
		signal.(func(int64))(int64(id))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptUnloadDefault(int64(id))
	}
}

func (ptr *QScriptEngineAgent) ConnectScriptUnload(f func(id int64)) {
	defer qt.Recovering("connect QScriptEngineAgent::scriptUnload")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "scriptUnload", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectScriptUnload() {
	defer qt.Recovering("disconnect QScriptEngineAgent::scriptUnload")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "scriptUnload")
	}
}

func (ptr *QScriptEngineAgent) ScriptUnload(id int64) {
	defer qt.Recovering("QScriptEngineAgent::scriptUnload")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ScriptUnload(ptr.Pointer(), C.longlong(id))
	}
}

func (ptr *QScriptEngineAgent) ScriptUnloadDefault(id int64) {
	defer qt.Recovering("QScriptEngineAgent::scriptUnload")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ScriptUnloadDefault(ptr.Pointer(), C.longlong(id))
	}
}

//export callbackQScriptEngineAgent_SupportsExtension
func callbackQScriptEngineAgent_SupportsExtension(ptr unsafe.Pointer, ptrName *C.char, extension C.int) C.int {
	defer qt.Recovering("callback QScriptEngineAgent::supportsExtension")

	if signal := qt.GetSignal(C.GoString(ptrName), "supportsExtension"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(QScriptEngineAgent__Extension) bool)(QScriptEngineAgent__Extension(extension))))
	}

	return C.int(qt.GoBoolToInt(NewQScriptEngineAgentFromPointer(ptr).SupportsExtensionDefault(QScriptEngineAgent__Extension(extension))))
}

func (ptr *QScriptEngineAgent) ConnectSupportsExtension(f func(extension QScriptEngineAgent__Extension) bool) {
	defer qt.Recovering("connect QScriptEngineAgent::supportsExtension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "supportsExtension", f)
	}
}

func (ptr *QScriptEngineAgent) DisconnectSupportsExtension() {
	defer qt.Recovering("disconnect QScriptEngineAgent::supportsExtension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "supportsExtension")
	}
}

func (ptr *QScriptEngineAgent) SupportsExtension(extension QScriptEngineAgent__Extension) bool {
	defer qt.Recovering("QScriptEngineAgent::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptEngineAgent_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) SupportsExtensionDefault(extension QScriptEngineAgent__Extension) bool {
	defer qt.Recovering("QScriptEngineAgent::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptEngineAgent_SupportsExtensionDefault(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgent() {
	defer qt.Recovering("QScriptEngineAgent::~QScriptEngineAgent")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QScriptEngineAgent_DestroyQScriptEngineAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngineAgent) ObjectNameAbs() string {
	defer qt.Recovering("QScriptEngineAgent::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptEngineAgent_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptEngineAgent) SetObjectNameAbs(name string) {
	defer qt.Recovering("QScriptEngineAgent::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QScriptExtensionPlugin struct {
	core.QObject
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
	QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin
}

func (p *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return p
}

func (p *QScriptExtensionPlugin) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QScriptExtensionPlugin) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	return n
}

func newQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = NewQScriptExtensionPluginFromPointer(ptr)
	for len(n.ObjectName()) < len("QScriptExtensionPlugin_") {
		n.SetObjectName("QScriptExtensionPlugin_" + qt.Identifier())
	}
	return n
}

func NewQScriptExtensionPlugin(parent core.QObject_ITF) *QScriptExtensionPlugin {
	defer qt.Recovering("QScriptExtensionPlugin::QScriptExtensionPlugin")

	return newQScriptExtensionPluginFromPointer(C.QScriptExtensionPlugin_NewQScriptExtensionPlugin(core.PointerFromQObject(parent)))
}

//export callbackQScriptExtensionPlugin_Initialize
func callbackQScriptExtensionPlugin_Initialize(ptr unsafe.Pointer, ptrName *C.char, key *C.char, engine unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::initialize")

	if signal := qt.GetSignal(C.GoString(ptrName), "initialize"); signal != nil {
		signal.(func(string, *QScriptEngine))(C.GoString(key), NewQScriptEngineFromPointer(engine))
	}

}

func (ptr *QScriptExtensionPlugin) ConnectInitialize(f func(key string, engine *QScriptEngine)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::initialize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initialize", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectInitialize(key string, engine QScriptEngine_ITF) {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::initialize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initialize")
	}
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::initialize")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_Initialize(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine))
	}
}

//export callbackQScriptExtensionPlugin_Keys
func callbackQScriptExtensionPlugin_Keys(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QScriptExtensionPlugin::keys")

	if signal := qt.GetSignal(C.GoString(ptrName), "keys"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(make([]string, 0), "|"))
}

func (ptr *QScriptExtensionPlugin) ConnectKeys(f func() []string) {
	defer qt.Recovering("connect QScriptExtensionPlugin::keys")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keys", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectKeys() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::keys")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keys")
	}
}

func (ptr *QScriptExtensionPlugin) Keys() []string {
	defer qt.Recovering("QScriptExtensionPlugin::keys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptExtensionPlugin_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {
	defer qt.Recovering("QScriptExtensionPlugin::setupPackage")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptExtensionPlugin_SetupPackage(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine)))
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	defer qt.Recovering("QScriptExtensionPlugin::~QScriptExtensionPlugin")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScriptExtensionPlugin_TimerEvent
func callbackQScriptExtensionPlugin_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScriptExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScriptExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScriptExtensionPlugin_ChildEvent
func callbackQScriptExtensionPlugin_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScriptExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScriptExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScriptExtensionPlugin_ConnectNotify
func callbackQScriptExtensionPlugin_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScriptExtensionPlugin) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptExtensionPlugin_CustomEvent
func callbackQScriptExtensionPlugin_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScriptExtensionPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScriptExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScriptExtensionPlugin_DeleteLater
func callbackQScriptExtensionPlugin_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScriptExtensionPlugin::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptExtensionPlugin) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScriptExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScriptExtensionPlugin) DeleteLater() {
	defer qt.Recovering("QScriptExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScriptExtensionPlugin_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptExtensionPlugin) DeleteLaterDefault() {
	defer qt.Recovering("QScriptExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScriptExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScriptExtensionPlugin_DisconnectNotify
func callbackQScriptExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptExtensionPlugin_Event
func callbackQScriptExtensionPlugin_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScriptExtensionPlugin::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScriptExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScriptExtensionPlugin) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScriptExtensionPlugin::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScriptExtensionPlugin) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptExtensionPlugin::event")

	if ptr.Pointer() != nil {
		return C.QScriptExtensionPlugin_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScriptExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptExtensionPlugin::event")

	if ptr.Pointer() != nil {
		return C.QScriptExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScriptExtensionPlugin_EventFilter
func callbackQScriptExtensionPlugin_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScriptExtensionPlugin::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScriptExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScriptExtensionPlugin) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScriptExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScriptExtensionPlugin) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScriptExtensionPlugin_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScriptExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScriptExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScriptExtensionPlugin_MetaObject
func callbackQScriptExtensionPlugin_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScriptExtensionPlugin::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScriptExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptExtensionPlugin) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScriptExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScriptExtensionPlugin) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScriptExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptExtensionPlugin_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScriptExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptExtensionPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QScriptProgram struct {
	ptr unsafe.Pointer
}

type QScriptProgram_ITF interface {
	QScriptProgram_PTR() *QScriptProgram
}

func (p *QScriptProgram) QScriptProgram_PTR() *QScriptProgram {
	return p
}

func (p *QScriptProgram) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptProgram) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptProgram(ptr QScriptProgram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptProgram_PTR().Pointer()
	}
	return nil
}

func NewQScriptProgramFromPointer(ptr unsafe.Pointer) *QScriptProgram {
	var n = new(QScriptProgram)
	n.SetPointer(ptr)
	return n
}

func newQScriptProgramFromPointer(ptr unsafe.Pointer) *QScriptProgram {
	var n = NewQScriptProgramFromPointer(ptr)
	return n
}

func NewQScriptProgram() *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return newQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram())
}

func NewQScriptProgram3(other QScriptProgram_ITF) *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return newQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram3(PointerFromQScriptProgram(other)))
}

func NewQScriptProgram2(sourceCode string, fileName string, firstLineNumber int) *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return newQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram2(C.CString(sourceCode), C.CString(fileName), C.int(firstLineNumber)))
}

func (ptr *QScriptProgram) FileName() string {
	defer qt.Recovering("QScriptProgram::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptProgram_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) FirstLineNumber() int {
	defer qt.Recovering("QScriptProgram::firstLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptProgram_FirstLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptProgram) IsNull() bool {
	defer qt.Recovering("QScriptProgram::isNull")

	if ptr.Pointer() != nil {
		return C.QScriptProgram_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptProgram) SourceCode() string {
	defer qt.Recovering("QScriptProgram::sourceCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptProgram_SourceCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) DestroyQScriptProgram() {
	defer qt.Recovering("QScriptProgram::~QScriptProgram")

	if ptr.Pointer() != nil {
		C.QScriptProgram_DestroyQScriptProgram(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptString struct {
	ptr unsafe.Pointer
}

type QScriptString_ITF interface {
	QScriptString_PTR() *QScriptString
}

func (p *QScriptString) QScriptString_PTR() *QScriptString {
	return p
}

func (p *QScriptString) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptString) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptString(ptr QScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptString_PTR().Pointer()
	}
	return nil
}

func NewQScriptStringFromPointer(ptr unsafe.Pointer) *QScriptString {
	var n = new(QScriptString)
	n.SetPointer(ptr)
	return n
}

func newQScriptStringFromPointer(ptr unsafe.Pointer) *QScriptString {
	var n = NewQScriptStringFromPointer(ptr)
	return n
}

func NewQScriptString() *QScriptString {
	defer qt.Recovering("QScriptString::QScriptString")

	return newQScriptStringFromPointer(C.QScriptString_NewQScriptString())
}

func NewQScriptString2(other QScriptString_ITF) *QScriptString {
	defer qt.Recovering("QScriptString::QScriptString")

	return newQScriptStringFromPointer(C.QScriptString_NewQScriptString2(PointerFromQScriptString(other)))
}

func (ptr *QScriptString) IsValid() bool {
	defer qt.Recovering("QScriptString::isValid")

	if ptr.Pointer() != nil {
		return C.QScriptString_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptString) ToString() string {
	defer qt.Recovering("QScriptString::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptString_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptString) DestroyQScriptString() {
	defer qt.Recovering("QScriptString::~QScriptString")

	if ptr.Pointer() != nil {
		C.QScriptString_DestroyQScriptString(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QScriptSyntaxCheckResult::State
type QScriptSyntaxCheckResult__State int64

const (
	QScriptSyntaxCheckResult__Error        = QScriptSyntaxCheckResult__State(0)
	QScriptSyntaxCheckResult__Intermediate = QScriptSyntaxCheckResult__State(1)
	QScriptSyntaxCheckResult__Valid        = QScriptSyntaxCheckResult__State(2)
)

type QScriptSyntaxCheckResult struct {
	ptr unsafe.Pointer
}

type QScriptSyntaxCheckResult_ITF interface {
	QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult
}

func (p *QScriptSyntaxCheckResult) QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult {
	return p
}

func (p *QScriptSyntaxCheckResult) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptSyntaxCheckResult) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptSyntaxCheckResult(ptr QScriptSyntaxCheckResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptSyntaxCheckResult_PTR().Pointer()
	}
	return nil
}

func NewQScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) *QScriptSyntaxCheckResult {
	var n = new(QScriptSyntaxCheckResult)
	n.SetPointer(ptr)
	return n
}

func newQScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) *QScriptSyntaxCheckResult {
	var n = NewQScriptSyntaxCheckResultFromPointer(ptr)
	return n
}

func NewQScriptSyntaxCheckResult(other QScriptSyntaxCheckResult_ITF) *QScriptSyntaxCheckResult {
	defer qt.Recovering("QScriptSyntaxCheckResult::QScriptSyntaxCheckResult")

	return newQScriptSyntaxCheckResultFromPointer(C.QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(PointerFromQScriptSyntaxCheckResult(other)))
}

func (ptr *QScriptSyntaxCheckResult) ErrorColumnNumber() int {
	defer qt.Recovering("QScriptSyntaxCheckResult::errorColumnNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptSyntaxCheckResult_ErrorColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorLineNumber() int {
	defer qt.Recovering("QScriptSyntaxCheckResult::errorLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptSyntaxCheckResult_ErrorLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorMessage() string {
	defer qt.Recovering("QScriptSyntaxCheckResult::errorMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptSyntaxCheckResult_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptSyntaxCheckResult) State() QScriptSyntaxCheckResult__State {
	defer qt.Recovering("QScriptSyntaxCheckResult::state")

	if ptr.Pointer() != nil {
		return QScriptSyntaxCheckResult__State(C.QScriptSyntaxCheckResult_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) DestroyQScriptSyntaxCheckResult() {
	defer qt.Recovering("QScriptSyntaxCheckResult::~QScriptSyntaxCheckResult")

	if ptr.Pointer() != nil {
		C.QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QScriptValue::PropertyFlag
type QScriptValue__PropertyFlag int64

const (
	QScriptValue__ReadOnly          = QScriptValue__PropertyFlag(0x00000001)
	QScriptValue__Undeletable       = QScriptValue__PropertyFlag(0x00000002)
	QScriptValue__SkipInEnumeration = QScriptValue__PropertyFlag(0x00000004)
	QScriptValue__PropertyGetter    = QScriptValue__PropertyFlag(0x00000008)
	QScriptValue__PropertySetter    = QScriptValue__PropertyFlag(0x00000010)
	QScriptValue__QObjectMember     = QScriptValue__PropertyFlag(0x00000020)
	QScriptValue__KeepExistingFlags = QScriptValue__PropertyFlag(0x00000800)
	QScriptValue__UserRange         = QScriptValue__PropertyFlag(0xff000000)
)

//QScriptValue::ResolveFlag
type QScriptValue__ResolveFlag int64

const (
	QScriptValue__ResolveLocal     = QScriptValue__ResolveFlag(0x00)
	QScriptValue__ResolvePrototype = QScriptValue__ResolveFlag(0x01)
	QScriptValue__ResolveScope     = QScriptValue__ResolveFlag(0x02)
	QScriptValue__ResolveFull      = QScriptValue__ResolveFlag(QScriptValue__ResolvePrototype | QScriptValue__ResolveScope)
)

//QScriptValue::SpecialValue
type QScriptValue__SpecialValue int64

const (
	QScriptValue__NullValue      = QScriptValue__SpecialValue(0)
	QScriptValue__UndefinedValue = QScriptValue__SpecialValue(1)
)

type QScriptValue struct {
	ptr unsafe.Pointer
}

type QScriptValue_ITF interface {
	QScriptValue_PTR() *QScriptValue
}

func (p *QScriptValue) QScriptValue_PTR() *QScriptValue {
	return p
}

func (p *QScriptValue) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptValue) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptValue(ptr QScriptValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValue_PTR().Pointer()
	}
	return nil
}

func NewQScriptValueFromPointer(ptr unsafe.Pointer) *QScriptValue {
	var n = new(QScriptValue)
	n.SetPointer(ptr)
	return n
}

func newQScriptValueFromPointer(ptr unsafe.Pointer) *QScriptValue {
	var n = NewQScriptValueFromPointer(ptr)
	return n
}

func NewQScriptValue() *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue())
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue10(C.int(value)))
}

func NewQScriptValue11(value bool) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue11(C.int(qt.GoBoolToInt(value))))
}

func NewQScriptValue16(value core.QLatin1String_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue16(core.PointerFromQLatin1String(value)))
}

func NewQScriptValue2(other QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue2(PointerFromQScriptValue(other)))
}

func NewQScriptValue15(value string) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue15(C.CString(value)))
}

func NewQScriptValue17(value string) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue17(C.CString(value)))
}

func NewQScriptValue12(value int) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return newQScriptValueFromPointer(C.QScriptValue_NewQScriptValue12(C.int(value)))
}

func (ptr *QScriptValue) Call2(thisObject QScriptValue_ITF, arguments QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::call")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Call2(ptr.Pointer(), PointerFromQScriptValue(thisObject), PointerFromQScriptValue(arguments)))
	}
	return nil
}

func (ptr *QScriptValue) Construct2(arguments QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::construct")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Construct2(ptr.Pointer(), PointerFromQScriptValue(arguments)))
	}
	return nil
}

func (ptr *QScriptValue) Data() *QScriptValue {
	defer qt.Recovering("QScriptValue::data")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptValue::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptValue_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) Equals(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::equals")

	if ptr.Pointer() != nil {
		return C.QScriptValue_Equals(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) InstanceOf(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::instanceOf")

	if ptr.Pointer() != nil {
		return C.QScriptValue_InstanceOf(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) IsArray() bool {
	defer qt.Recovering("QScriptValue::isArray")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsBool() bool {
	defer qt.Recovering("QScriptValue::isBool")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsDate() bool {
	defer qt.Recovering("QScriptValue::isDate")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsDate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsError() bool {
	defer qt.Recovering("QScriptValue::isError")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsFunction() bool {
	defer qt.Recovering("QScriptValue::isFunction")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsFunction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNull() bool {
	defer qt.Recovering("QScriptValue::isNull")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNumber() bool {
	defer qt.Recovering("QScriptValue::isNumber")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsObject() bool {
	defer qt.Recovering("QScriptValue::isObject")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQMetaObject() bool {
	defer qt.Recovering("QScriptValue::isQMetaObject")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQMetaObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQObject() bool {
	defer qt.Recovering("QScriptValue::isQObject")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsRegExp() bool {
	defer qt.Recovering("QScriptValue::isRegExp")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsRegExp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsString() bool {
	defer qt.Recovering("QScriptValue::isString")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsUndefined() bool {
	defer qt.Recovering("QScriptValue::isUndefined")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsValid() bool {
	defer qt.Recovering("QScriptValue::isValid")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsVariant() bool {
	defer qt.Recovering("QScriptValue::isVariant")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsVariant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) LessThan(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::lessThan")

	if ptr.Pointer() != nil {
		return C.QScriptValue_LessThan(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) Property3(name QScriptString_ITF, mode QScriptValue__ResolveFlag) *QScriptValue {
	defer qt.Recovering("QScriptValue::property")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Property3(ptr.Pointer(), PointerFromQScriptString(name), C.int(mode)))
	}
	return nil
}

func (ptr *QScriptValue) Property(name string, mode QScriptValue__ResolveFlag) *QScriptValue {
	defer qt.Recovering("QScriptValue::property")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Property(ptr.Pointer(), C.CString(name), C.int(mode)))
	}
	return nil
}

func (ptr *QScriptValue) PropertyFlags2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	defer qt.Recovering("QScriptValue::propertyFlags")

	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags2(ptr.Pointer(), PointerFromQScriptString(name), C.int(mode)))
	}
	return 0
}

func (ptr *QScriptValue) PropertyFlags(name string, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	defer qt.Recovering("QScriptValue::propertyFlags")

	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags(ptr.Pointer(), C.CString(name), C.int(mode)))
	}
	return 0
}

func (ptr *QScriptValue) Prototype() *QScriptValue {
	defer qt.Recovering("QScriptValue::prototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ScriptClass() *QScriptClass {
	defer qt.Recovering("QScriptValue::scriptClass")

	if ptr.Pointer() != nil {
		return NewQScriptClassFromPointer(C.QScriptValue_ScriptClass(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) SetData(data QScriptValue_ITF) {
	defer qt.Recovering("QScriptValue::setData")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetData(ptr.Pointer(), PointerFromQScriptValue(data))
	}
}

func (ptr *QScriptValue) SetProperty3(name QScriptString_ITF, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	defer qt.Recovering("QScriptValue::setProperty")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty3(ptr.Pointer(), PointerFromQScriptString(name), PointerFromQScriptValue(value), C.int(flags))
	}
}

func (ptr *QScriptValue) SetProperty(name string, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	defer qt.Recovering("QScriptValue::setProperty")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty(ptr.Pointer(), C.CString(name), PointerFromQScriptValue(value), C.int(flags))
	}
}

func (ptr *QScriptValue) SetPrototype(prototype QScriptValue_ITF) {
	defer qt.Recovering("QScriptValue::setPrototype")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetPrototype(ptr.Pointer(), PointerFromQScriptValue(prototype))
	}
}

func (ptr *QScriptValue) SetScriptClass(scriptClass QScriptClass_ITF) {
	defer qt.Recovering("QScriptValue::setScriptClass")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetScriptClass(ptr.Pointer(), PointerFromQScriptClass(scriptClass))
	}
}

func (ptr *QScriptValue) StrictlyEquals(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::strictlyEquals")

	if ptr.Pointer() != nil {
		return C.QScriptValue_StrictlyEquals(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) ToBool() bool {
	defer qt.Recovering("QScriptValue::toBool")

	if ptr.Pointer() != nil {
		return C.QScriptValue_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) ToDateTime() *core.QDateTime {
	defer qt.Recovering("QScriptValue::toDateTime")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QScriptValue_ToDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToQMetaObject() *core.QMetaObject {
	defer qt.Recovering("QScriptValue::toQMetaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptValue_ToQMetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToQObject() *core.QObject {
	defer qt.Recovering("QScriptValue::toQObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QScriptValue_ToQObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToRegExp() *core.QRegExp {
	defer qt.Recovering("QScriptValue::toRegExp")

	if ptr.Pointer() != nil {
		return core.NewQRegExpFromPointer(C.QScriptValue_ToRegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToString() string {
	defer qt.Recovering("QScriptValue::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptValue) ToVariant() *core.QVariant {
	defer qt.Recovering("QScriptValue::toVariant")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptValue_ToVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) DestroyQScriptValue() {
	defer qt.Recovering("QScriptValue::~QScriptValue")

	if ptr.Pointer() != nil {
		C.QScriptValue_DestroyQScriptValue(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptValueIterator struct {
	ptr unsafe.Pointer
}

type QScriptValueIterator_ITF interface {
	QScriptValueIterator_PTR() *QScriptValueIterator
}

func (p *QScriptValueIterator) QScriptValueIterator_PTR() *QScriptValueIterator {
	return p
}

func (p *QScriptValueIterator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptValueIterator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptValueIterator(ptr QScriptValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQScriptValueIteratorFromPointer(ptr unsafe.Pointer) *QScriptValueIterator {
	var n = new(QScriptValueIterator)
	n.SetPointer(ptr)
	return n
}

func newQScriptValueIteratorFromPointer(ptr unsafe.Pointer) *QScriptValueIterator {
	var n = NewQScriptValueIteratorFromPointer(ptr)
	return n
}

type QScriptable struct {
	ptr unsafe.Pointer
}

type QScriptable_ITF interface {
	QScriptable_PTR() *QScriptable
}

func (p *QScriptable) QScriptable_PTR() *QScriptable {
	return p
}

func (p *QScriptable) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScriptable) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQScriptable(ptr QScriptable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptable_PTR().Pointer()
	}
	return nil
}

func NewQScriptableFromPointer(ptr unsafe.Pointer) *QScriptable {
	var n = new(QScriptable)
	n.SetPointer(ptr)
	return n
}

func newQScriptableFromPointer(ptr unsafe.Pointer) *QScriptable {
	var n = NewQScriptableFromPointer(ptr)
	return n
}

func (ptr *QScriptable) Argument(index int) *QScriptValue {
	defer qt.Recovering("QScriptable::argument")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptable_Argument(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QScriptable) ArgumentCount() int {
	defer qt.Recovering("QScriptable::argumentCount")

	if ptr.Pointer() != nil {
		return int(C.QScriptable_ArgumentCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptable) Context() *QScriptContext {
	defer qt.Recovering("QScriptable::context")

	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptable_Context(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptable) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptable::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptable_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptable) ThisObject() *QScriptValue {
	defer qt.Recovering("QScriptable::thisObject")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptable_ThisObject(ptr.Pointer()))
	}
	return nil
}
