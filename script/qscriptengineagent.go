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
func callbackQScriptEngineAgentContextPop(ptrName *C.char) bool {
	defer qt.Recovering("callback QScriptEngineAgent::contextPop")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextPop")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

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
func callbackQScriptEngineAgentContextPush(ptrName *C.char) bool {
	defer qt.Recovering("callback QScriptEngineAgent::contextPush")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextPush")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

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
func callbackQScriptEngineAgentExceptionCatch(ptrName *C.char, scriptId C.longlong, exception unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngineAgent::exceptionCatch")

	var signal = qt.GetSignal(C.GoString(ptrName), "exceptionCatch")
	if signal != nil {
		defer signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(exception))
		return true
	}
	return false

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
func callbackQScriptEngineAgentExceptionThrow(ptrName *C.char, scriptId C.longlong, exception unsafe.Pointer, hasHandler C.int) bool {
	defer qt.Recovering("callback QScriptEngineAgent::exceptionThrow")

	var signal = qt.GetSignal(C.GoString(ptrName), "exceptionThrow")
	if signal != nil {
		defer signal.(func(int64, *QScriptValue, bool))(int64(scriptId), NewQScriptValueFromPointer(exception), int(hasHandler) != 0)
		return true
	}
	return false

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
func callbackQScriptEngineAgentFunctionEntry(ptrName *C.char, scriptId C.longlong) bool {
	defer qt.Recovering("callback QScriptEngineAgent::functionEntry")

	var signal = qt.GetSignal(C.GoString(ptrName), "functionEntry")
	if signal != nil {
		defer signal.(func(int64))(int64(scriptId))
		return true
	}
	return false

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
func callbackQScriptEngineAgentFunctionExit(ptrName *C.char, scriptId C.longlong, returnValue unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngineAgent::functionExit")

	var signal = qt.GetSignal(C.GoString(ptrName), "functionExit")
	if signal != nil {
		defer signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(returnValue))
		return true
	}
	return false

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
func callbackQScriptEngineAgentPositionChange(ptrName *C.char, scriptId C.longlong, lineNumber C.int, columnNumber C.int) bool {
	defer qt.Recovering("callback QScriptEngineAgent::positionChange")

	var signal = qt.GetSignal(C.GoString(ptrName), "positionChange")
	if signal != nil {
		defer signal.(func(int64, int, int))(int64(scriptId), int(lineNumber), int(columnNumber))
		return true
	}
	return false

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
func callbackQScriptEngineAgentScriptLoad(ptrName *C.char, id C.longlong, program *C.char, fileName *C.char, baseLineNumber C.int) bool {
	defer qt.Recovering("callback QScriptEngineAgent::scriptLoad")

	var signal = qt.GetSignal(C.GoString(ptrName), "scriptLoad")
	if signal != nil {
		defer signal.(func(int64, string, string, int))(int64(id), C.GoString(program), C.GoString(fileName), int(baseLineNumber))
		return true
	}
	return false

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
func callbackQScriptEngineAgentScriptUnload(ptrName *C.char, id C.longlong) bool {
	defer qt.Recovering("callback QScriptEngineAgent::scriptUnload")

	var signal = qt.GetSignal(C.GoString(ptrName), "scriptUnload")
	if signal != nil {
		defer signal.(func(int64))(int64(id))
		return true
	}
	return false

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
