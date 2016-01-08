package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QScriptClass struct {
	ptr unsafe.Pointer
}

type QScriptClass_ITF interface {
	QScriptClass_PTR() *QScriptClass
}

func (p *QScriptClass) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptClass) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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
	for len(n.ObjectNameAbs()) < len("QScriptClass_") {
		n.SetObjectNameAbs("QScriptClass_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptClass) QScriptClass_PTR() *QScriptClass {
	return ptr
}

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

func NewQScriptClass(engine QScriptEngine_ITF) *QScriptClass {
	defer qt.Recovering("QScriptClass::QScriptClass")

	return NewQScriptClassFromPointer(C.QScriptClass_NewQScriptClass(PointerFromQScriptEngine(engine)))
}

func (ptr *QScriptClass) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptClass::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptClass_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptClass) Extension(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptClass::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptClass_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

func (ptr *QScriptClass) Name() string {
	defer qt.Recovering("QScriptClass::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptClass) NewIterator(object QScriptValue_ITF) *QScriptClassPropertyIterator {
	defer qt.Recovering("QScriptClass::newIterator")

	if ptr.Pointer() != nil {
		return NewQScriptClassPropertyIteratorFromPointer(C.QScriptClass_NewIterator(ptr.Pointer(), PointerFromQScriptValue(object)))
	}
	return nil
}

func (ptr *QScriptClass) Prototype() *QScriptValue {
	defer qt.Recovering("QScriptClass::prototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptClass_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptClass) SupportsExtension(extension QScriptClass__Extension) bool {
	defer qt.Recovering("QScriptClass::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptClass_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptClass) DestroyQScriptClass() {
	defer qt.Recovering("QScriptClass::~QScriptClass")

	if ptr.Pointer() != nil {
		C.QScriptClass_DestroyQScriptClass(ptr.Pointer())
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

func (p *QScriptClassPropertyIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptClassPropertyIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptClassPropertyIterator) QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator {
	return ptr
}

type QScriptContext struct {
	ptr unsafe.Pointer
}

type QScriptContext_ITF interface {
	QScriptContext_PTR() *QScriptContext
}

func (p *QScriptContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptContext) QScriptContext_PTR() *QScriptContext {
	return ptr
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
	}
}

type QScriptContextInfo struct {
	ptr unsafe.Pointer
}

type QScriptContextInfo_ITF interface {
	QScriptContextInfo_PTR() *QScriptContextInfo
}

func (p *QScriptContextInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptContextInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptContextInfo) QScriptContextInfo_PTR() *QScriptContextInfo {
	return ptr
}

//QScriptContextInfo::FunctionType
type QScriptContextInfo__FunctionType int64

const (
	QScriptContextInfo__ScriptFunction     = QScriptContextInfo__FunctionType(0)
	QScriptContextInfo__QtFunction         = QScriptContextInfo__FunctionType(1)
	QScriptContextInfo__QtPropertyFunction = QScriptContextInfo__FunctionType(2)
	QScriptContextInfo__NativeFunction     = QScriptContextInfo__FunctionType(3)
)

func NewQScriptContextInfo3() *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo3())
}

func NewQScriptContextInfo(context QScriptContext_ITF) *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo(PointerFromQScriptContext(context)))
}

func NewQScriptContextInfo2(other QScriptContextInfo_ITF) *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo2(PointerFromQScriptContextInfo(other)))
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
	}
}

type QScriptEngine struct {
	core.QObject
}

type QScriptEngine_ITF interface {
	core.QObject_ITF
	QScriptEngine_PTR() *QScriptEngine
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
	for len(n.ObjectName()) < len("QScriptEngine_") {
		n.SetObjectName("QScriptEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptEngine) QScriptEngine_PTR() *QScriptEngine {
	return ptr
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

func NewQScriptEngine() *QScriptEngine {
	defer qt.Recovering("QScriptEngine::QScriptEngine")

	return NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine())
}

func NewQScriptEngine2(parent core.QObject_ITF) *QScriptEngine {
	defer qt.Recovering("QScriptEngine::QScriptEngine")

	return NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine2(core.PointerFromQObject(parent)))
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

//export callbackQScriptEngineSignalHandlerException
func callbackQScriptEngineSignalHandlerException(ptr unsafe.Pointer, ptrName *C.char, exception unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::signalHandlerException")

	if signal := qt.GetSignal(C.GoString(ptrName), "signalHandlerException"); signal != nil {
		signal.(func(*QScriptValue))(NewQScriptValueFromPointer(exception))
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
		C.QScriptEngine_DestroyQScriptEngine(ptr.Pointer())
		ptr.SetPointer(nil)
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

//export callbackQScriptEngineTimerEvent
func callbackQScriptEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQScriptEngineChildEvent
func callbackQScriptEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQScriptEngineCustomEvent
func callbackQScriptEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

type QScriptEngineAgent struct {
	ptr unsafe.Pointer
}

type QScriptEngineAgent_ITF interface {
	QScriptEngineAgent_PTR() *QScriptEngineAgent
}

func (p *QScriptEngineAgent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptEngineAgent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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
	for len(n.ObjectNameAbs()) < len("QScriptEngineAgent_") {
		n.SetObjectNameAbs("QScriptEngineAgent_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptEngineAgent) QScriptEngineAgent_PTR() *QScriptEngineAgent {
	return ptr
}

//QScriptEngineAgent::Extension
type QScriptEngineAgent__Extension int64

const (
	QScriptEngineAgent__DebuggerInvocationRequest = QScriptEngineAgent__Extension(0)
)

func NewQScriptEngineAgent(engine QScriptEngine_ITF) *QScriptEngineAgent {
	defer qt.Recovering("QScriptEngineAgent::QScriptEngineAgent")

	return NewQScriptEngineAgentFromPointer(C.QScriptEngineAgent_NewQScriptEngineAgent(PointerFromQScriptEngine(engine)))
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

//export callbackQScriptEngineAgentContextPop
func callbackQScriptEngineAgentContextPop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngineAgent::contextPop")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextPop"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPopDefault()
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

//export callbackQScriptEngineAgentContextPush
func callbackQScriptEngineAgentContextPush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngineAgent::contextPush")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextPush"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPushDefault()
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

//export callbackQScriptEngineAgentExceptionCatch
func callbackQScriptEngineAgentExceptionCatch(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, exception unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineAgent::exceptionCatch")

	if signal := qt.GetSignal(C.GoString(ptrName), "exceptionCatch"); signal != nil {
		signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(exception))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionCatchDefault(int64(scriptId), NewQScriptValueFromPointer(exception))
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

//export callbackQScriptEngineAgentExceptionThrow
func callbackQScriptEngineAgentExceptionThrow(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, exception unsafe.Pointer, hasHandler C.int) {
	defer qt.Recovering("callback QScriptEngineAgent::exceptionThrow")

	if signal := qt.GetSignal(C.GoString(ptrName), "exceptionThrow"); signal != nil {
		signal.(func(int64, *QScriptValue, bool))(int64(scriptId), NewQScriptValueFromPointer(exception), int(hasHandler) != 0)
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionThrowDefault(int64(scriptId), NewQScriptValueFromPointer(exception), int(hasHandler) != 0)
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

func (ptr *QScriptEngineAgent) Extension(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptEngineAgent::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptEngineAgent_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
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

//export callbackQScriptEngineAgentFunctionEntry
func callbackQScriptEngineAgentFunctionEntry(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong) {
	defer qt.Recovering("callback QScriptEngineAgent::functionEntry")

	if signal := qt.GetSignal(C.GoString(ptrName), "functionEntry"); signal != nil {
		signal.(func(int64))(int64(scriptId))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionEntryDefault(int64(scriptId))
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

//export callbackQScriptEngineAgentFunctionExit
func callbackQScriptEngineAgentFunctionExit(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, returnValue unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineAgent::functionExit")

	if signal := qt.GetSignal(C.GoString(ptrName), "functionExit"); signal != nil {
		signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionExitDefault(int64(scriptId), NewQScriptValueFromPointer(returnValue))
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

//export callbackQScriptEngineAgentPositionChange
func callbackQScriptEngineAgentPositionChange(ptr unsafe.Pointer, ptrName *C.char, scriptId C.longlong, lineNumber C.int, columnNumber C.int) {
	defer qt.Recovering("callback QScriptEngineAgent::positionChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChange"); signal != nil {
		signal.(func(int64, int, int))(int64(scriptId), int(lineNumber), int(columnNumber))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).PositionChangeDefault(int64(scriptId), int(lineNumber), int(columnNumber))
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

//export callbackQScriptEngineAgentScriptLoad
func callbackQScriptEngineAgentScriptLoad(ptr unsafe.Pointer, ptrName *C.char, id C.longlong, program *C.char, fileName *C.char, baseLineNumber C.int) {
	defer qt.Recovering("callback QScriptEngineAgent::scriptLoad")

	if signal := qt.GetSignal(C.GoString(ptrName), "scriptLoad"); signal != nil {
		signal.(func(int64, string, string, int))(int64(id), C.GoString(program), C.GoString(fileName), int(baseLineNumber))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptLoadDefault(int64(id), C.GoString(program), C.GoString(fileName), int(baseLineNumber))
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

//export callbackQScriptEngineAgentScriptUnload
func callbackQScriptEngineAgentScriptUnload(ptr unsafe.Pointer, ptrName *C.char, id C.longlong) {
	defer qt.Recovering("callback QScriptEngineAgent::scriptUnload")

	if signal := qt.GetSignal(C.GoString(ptrName), "scriptUnload"); signal != nil {
		signal.(func(int64))(int64(id))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptUnloadDefault(int64(id))
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

func (ptr *QScriptEngineAgent) SupportsExtension(extension QScriptEngineAgent__Extension) bool {
	defer qt.Recovering("QScriptEngineAgent::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptEngineAgent_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgent() {
	defer qt.Recovering("QScriptEngineAgent::~QScriptEngineAgent")

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_DestroyQScriptEngineAgent(ptr.Pointer())
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

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScriptExtensionPlugin_") {
		n.SetObjectName("QScriptExtensionPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::initialize")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_Initialize(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine))
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
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
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

//export callbackQScriptExtensionPluginTimerEvent
func callbackQScriptExtensionPluginTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQScriptExtensionPluginChildEvent
func callbackQScriptExtensionPluginChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQScriptExtensionPluginCustomEvent
func callbackQScriptExtensionPluginCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptExtensionPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

type QScriptProgram struct {
	ptr unsafe.Pointer
}

type QScriptProgram_ITF interface {
	QScriptProgram_PTR() *QScriptProgram
}

func (p *QScriptProgram) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptProgram) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptProgram) QScriptProgram_PTR() *QScriptProgram {
	return ptr
}

func NewQScriptProgram() *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram())
}

func NewQScriptProgram3(other QScriptProgram_ITF) *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram3(PointerFromQScriptProgram(other)))
}

func NewQScriptProgram2(sourceCode string, fileName string, firstLineNumber int) *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram2(C.CString(sourceCode), C.CString(fileName), C.int(firstLineNumber)))
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
	}
}

type QScriptString struct {
	ptr unsafe.Pointer
}

type QScriptString_ITF interface {
	QScriptString_PTR() *QScriptString
}

func (p *QScriptString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptString) QScriptString_PTR() *QScriptString {
	return ptr
}

func NewQScriptString() *QScriptString {
	defer qt.Recovering("QScriptString::QScriptString")

	return NewQScriptStringFromPointer(C.QScriptString_NewQScriptString())
}

func NewQScriptString2(other QScriptString_ITF) *QScriptString {
	defer qt.Recovering("QScriptString::QScriptString")

	return NewQScriptStringFromPointer(C.QScriptString_NewQScriptString2(PointerFromQScriptString(other)))
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
	}
}

type QScriptSyntaxCheckResult struct {
	ptr unsafe.Pointer
}

type QScriptSyntaxCheckResult_ITF interface {
	QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult
}

func (p *QScriptSyntaxCheckResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptSyntaxCheckResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptSyntaxCheckResult) QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult {
	return ptr
}

//QScriptSyntaxCheckResult::State
type QScriptSyntaxCheckResult__State int64

const (
	QScriptSyntaxCheckResult__Error        = QScriptSyntaxCheckResult__State(0)
	QScriptSyntaxCheckResult__Intermediate = QScriptSyntaxCheckResult__State(1)
	QScriptSyntaxCheckResult__Valid        = QScriptSyntaxCheckResult__State(2)
)

func NewQScriptSyntaxCheckResult(other QScriptSyntaxCheckResult_ITF) *QScriptSyntaxCheckResult {
	defer qt.Recovering("QScriptSyntaxCheckResult::QScriptSyntaxCheckResult")

	return NewQScriptSyntaxCheckResultFromPointer(C.QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(PointerFromQScriptSyntaxCheckResult(other)))
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
	}
}

type QScriptValue struct {
	ptr unsafe.Pointer
}

type QScriptValue_ITF interface {
	QScriptValue_PTR() *QScriptValue
}

func (p *QScriptValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptValue) QScriptValue_PTR() *QScriptValue {
	return ptr
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

func NewQScriptValue() *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue())
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue10(C.int(value)))
}

func NewQScriptValue11(value bool) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue11(C.int(qt.GoBoolToInt(value))))
}

func NewQScriptValue16(value core.QLatin1String_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue16(core.PointerFromQLatin1String(value)))
}

func NewQScriptValue2(other QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue2(PointerFromQScriptValue(other)))
}

func NewQScriptValue15(value string) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue15(C.CString(value)))
}

func NewQScriptValue17(value string) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue17(C.CString(value)))
}

func NewQScriptValue12(value int) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue12(C.int(value)))
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

func (ptr *QScriptValue) Property2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) *QScriptValue {
	defer qt.Recovering("QScriptValue::property")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Property2(ptr.Pointer(), PointerFromQScriptString(name), C.int(mode)))
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

func (ptr *QScriptValue) SetProperty2(name QScriptString_ITF, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	defer qt.Recovering("QScriptValue::setProperty")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty2(ptr.Pointer(), PointerFromQScriptString(name), PointerFromQScriptValue(value), C.int(flags))
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
	}
}

type QScriptValueIterator struct {
	ptr unsafe.Pointer
}

type QScriptValueIterator_ITF interface {
	QScriptValueIterator_PTR() *QScriptValueIterator
}

func (p *QScriptValueIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptValueIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptValueIterator) QScriptValueIterator_PTR() *QScriptValueIterator {
	return ptr
}

type QScriptable struct {
	ptr unsafe.Pointer
}

type QScriptable_ITF interface {
	QScriptable_PTR() *QScriptable
}

func (p *QScriptable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QScriptable) QScriptable_PTR() *QScriptable {
	return ptr
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
