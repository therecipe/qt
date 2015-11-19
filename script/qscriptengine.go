package script

//#include "qscriptengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

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
	if n.ObjectName() == "" {
		n.SetObjectName("QScriptEngine_" + qt.RandomIdentifier())
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
	return NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine())
}

func NewQScriptEngine2(parent core.QObject_ITF) *QScriptEngine {
	return NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine2(core.PointerFromQObject(parent)))
}

func (ptr *QScriptEngine) AbortEvaluation(result QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_AbortEvaluation(ptr.Pointer(), PointerFromQScriptValue(result))
	}
}

func (ptr *QScriptEngine) Agent() *QScriptEngineAgent {
	if ptr.Pointer() != nil {
		return NewQScriptEngineAgentFromPointer(C.QScriptEngine_Agent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) AvailableExtensions() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_AvailableExtensions(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) ClearExceptions() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ClearExceptions(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) CollectGarbage() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) CurrentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_CurrentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) DefaultPrototype(metaTypeId int) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_DefaultPrototype(ptr.Pointer(), C.int(metaTypeId)))
	}
	return nil
}

func (ptr *QScriptEngine) Evaluate2(program QScriptProgram_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_Evaluate2(ptr.Pointer(), PointerFromQScriptProgram(program)))
	}
	return nil
}

func (ptr *QScriptEngine) Evaluate(program string, fileName string, lineNumber int) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_Evaluate(ptr.Pointer(), C.CString(program), C.CString(fileName), C.int(lineNumber)))
	}
	return nil
}

func (ptr *QScriptEngine) GlobalObject() *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_GlobalObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) HasUncaughtException() bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_HasUncaughtException(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngine) ImportExtension(extension string) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_ImportExtension(ptr.Pointer(), C.CString(extension)))
	}
	return nil
}

func (ptr *QScriptEngine) ImportedExtensions() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_ImportedExtensions(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) InstallTranslatorFunctions(object QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) IsEvaluating() bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_IsEvaluating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngine) NewDate2(value core.QDateTime_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewDate2(ptr.Pointer(), core.PointerFromQDateTime(value)))
	}
	return nil
}

func (ptr *QScriptEngine) NewObject() *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) NewObject2(scriptClass QScriptClass_ITF, data QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewObject2(ptr.Pointer(), PointerFromQScriptClass(scriptClass), PointerFromQScriptValue(data)))
	}
	return nil
}

func (ptr *QScriptEngine) NewQMetaObject(metaObject core.QMetaObject_ITF, ctor QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewQMetaObject(ptr.Pointer(), core.PointerFromQMetaObject(metaObject), PointerFromQScriptValue(ctor)))
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject(object core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object), C.int(ownership), C.int(options)))
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject2(scriptObject QScriptValue_ITF, qtObject core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewQObject2(ptr.Pointer(), PointerFromQScriptValue(scriptObject), core.PointerFromQObject(qtObject), C.int(ownership), C.int(options)))
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp(regexp core.QRegExp_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp(ptr.Pointer(), core.PointerFromQRegExp(regexp)))
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp2(pattern string, flags string) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp2(ptr.Pointer(), C.CString(pattern), C.CString(flags)))
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant2(object QScriptValue_ITF, value core.QVariant_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewVariant2(ptr.Pointer(), PointerFromQScriptValue(object), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant(value core.QVariant_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NewVariant(ptr.Pointer(), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QScriptEngine) NullValue() *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_NullValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) PopContext() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_PopContext(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) ProcessEventsInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptEngine_ProcessEventsInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptEngine) PushContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_PushContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) ReportAdditionalMemoryCost(size int) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ReportAdditionalMemoryCost(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QScriptEngine) SetAgent(agent QScriptEngineAgent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetAgent(ptr.Pointer(), PointerFromQScriptEngineAgent(agent))
	}
}

func (ptr *QScriptEngine) SetDefaultPrototype(metaTypeId int, prototype QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetDefaultPrototype(ptr.Pointer(), C.int(metaTypeId), PointerFromQScriptValue(prototype))
	}
}

func (ptr *QScriptEngine) SetGlobalObject(object QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetGlobalObject(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) SetProcessEventsInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetProcessEventsInterval(ptr.Pointer(), C.int(interval))
	}
}

func (ptr *QScriptEngine) ConnectSignalHandlerException(f func(exception *QScriptValue)) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ConnectSignalHandlerException(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalHandlerException", f)
	}
}

func (ptr *QScriptEngine) DisconnectSignalHandlerException() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectSignalHandlerException(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalHandlerException")
	}
}

//export callbackQScriptEngineSignalHandlerException
func callbackQScriptEngineSignalHandlerException(ptrName *C.char, exception unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "signalHandlerException").(func(*QScriptValue))(NewQScriptValueFromPointer(exception))
}

func (ptr *QScriptEngine) ToObject(value QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_ToObject(ptr.Pointer(), PointerFromQScriptValue(value)))
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtException() *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_UncaughtException(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtExceptionBacktrace() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_UncaughtExceptionBacktrace(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) UncaughtExceptionLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptEngine_UncaughtExceptionLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptEngine) UndefinedValue() *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptEngine_UndefinedValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) DestroyQScriptEngine() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DestroyQScriptEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
