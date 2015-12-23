package script

//#include "script.h"
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
		return strings.Split(C.GoString(C.QScriptEngine_AvailableExtensions(ptr.Pointer())), ",,,")
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
		return strings.Split(C.GoString(C.QScriptEngine_ImportedExtensions(ptr.Pointer())), ",,,")
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
func callbackQScriptEngineSignalHandlerException(ptrName *C.char, exception unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngine::signalHandlerException")

	if signal := qt.GetSignal(C.GoString(ptrName), "signalHandlerException"); signal != nil {
		signal.(func(*QScriptValue))(NewQScriptValueFromPointer(exception))
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
		return strings.Split(C.GoString(C.QScriptEngine_UncaughtExceptionBacktrace(ptr.Pointer())), ",,,")
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
func callbackQScriptEngineTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScriptEngineChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScriptEngineCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
