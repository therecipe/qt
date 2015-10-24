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

type QScriptEngineITF interface {
	core.QObjectITF
	QScriptEnginePTR() *QScriptEngine
}

func PointerFromQScriptEngine(ptr QScriptEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEnginePTR().Pointer()
	}
	return nil
}

func QScriptEngineFromPointer(ptr unsafe.Pointer) *QScriptEngine {
	var n = new(QScriptEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScriptEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScriptEngine) QScriptEnginePTR() *QScriptEngine {
	return ptr
}

//QScriptEngine::QObjectWrapOption
type QScriptEngine__QObjectWrapOption int

var (
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
type QScriptEngine__ValueOwnership int

var (
	QScriptEngine__QtOwnership     = QScriptEngine__ValueOwnership(0)
	QScriptEngine__ScriptOwnership = QScriptEngine__ValueOwnership(1)
	QScriptEngine__AutoOwnership   = QScriptEngine__ValueOwnership(2)
)

func NewQScriptEngine() *QScriptEngine {
	return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptEngine_NewQScriptEngine()))
}

func NewQScriptEngine2(parent core.QObjectITF) *QScriptEngine {
	return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptEngine_NewQScriptEngine2(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QScriptEngine) AbortEvaluation(result QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_AbortEvaluation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(result)))
	}
}

func (ptr *QScriptEngine) Agent() *QScriptEngineAgent {
	if ptr.Pointer() != nil {
		return QScriptEngineAgentFromPointer(unsafe.Pointer(C.QScriptEngine_Agent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptEngine) AvailableExtensions() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_AvailableExtensions(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) ClearExceptions() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ClearExceptions(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScriptEngine) CollectGarbage() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_CollectGarbage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScriptEngine) CurrentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return QScriptContextFromPointer(unsafe.Pointer(C.QScriptEngine_CurrentContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptEngine) HasUncaughtException() bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_HasUncaughtException(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptEngine) ImportedExtensions() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_ImportedExtensions(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) InstallTranslatorFunctions(object QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_InstallTranslatorFunctions(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(object)))
	}
}

func (ptr *QScriptEngine) IsEvaluating() bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_IsEvaluating(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptEngine) PopContext() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_PopContext(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScriptEngine) ProcessEventsInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptEngine_ProcessEventsInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptEngine) PushContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return QScriptContextFromPointer(unsafe.Pointer(C.QScriptEngine_PushContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptEngine) ReportAdditionalMemoryCost(size int) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ReportAdditionalMemoryCost(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QScriptEngine) SetAgent(agent QScriptEngineAgentITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetAgent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptEngineAgent(agent)))
	}
}

func (ptr *QScriptEngine) SetDefaultPrototype(metaTypeId int, prototype QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetDefaultPrototype(C.QtObjectPtr(ptr.Pointer()), C.int(metaTypeId), C.QtObjectPtr(PointerFromQScriptValue(prototype)))
	}
}

func (ptr *QScriptEngine) SetGlobalObject(object QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetGlobalObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(object)))
	}
}

func (ptr *QScriptEngine) SetProcessEventsInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetProcessEventsInterval(C.QtObjectPtr(ptr.Pointer()), C.int(interval))
	}
}

func (ptr *QScriptEngine) UncaughtExceptionBacktrace() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptEngine_UncaughtExceptionBacktrace(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) UncaughtExceptionLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptEngine_UncaughtExceptionLineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptEngine) DestroyQScriptEngine() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DestroyQScriptEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
