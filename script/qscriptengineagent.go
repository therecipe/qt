package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

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
