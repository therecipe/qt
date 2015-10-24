package script

//#include "qscriptengineagent.h"
import "C"
import (
	"unsafe"
)

type QScriptEngineAgent struct {
	ptr unsafe.Pointer
}

type QScriptEngineAgentITF interface {
	QScriptEngineAgentPTR() *QScriptEngineAgent
}

func (p *QScriptEngineAgent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptEngineAgent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptEngineAgent(ptr QScriptEngineAgentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngineAgentPTR().Pointer()
	}
	return nil
}

func QScriptEngineAgentFromPointer(ptr unsafe.Pointer) *QScriptEngineAgent {
	var n = new(QScriptEngineAgent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptEngineAgent) QScriptEngineAgentPTR() *QScriptEngineAgent {
	return ptr
}

//QScriptEngineAgent::Extension
type QScriptEngineAgent__Extension int

var (
	QScriptEngineAgent__DebuggerInvocationRequest = QScriptEngineAgent__Extension(0)
)

func NewQScriptEngineAgent(engine QScriptEngineITF) *QScriptEngineAgent {
	return QScriptEngineAgentFromPointer(unsafe.Pointer(C.QScriptEngineAgent_NewQScriptEngineAgent(C.QtObjectPtr(PointerFromQScriptEngine(engine)))))
}

func (ptr *QScriptEngineAgent) ContextPop() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScriptEngineAgent) ContextPush() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScriptEngineAgent) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptEngineAgent_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptEngineAgent) Extension(extension QScriptEngineAgent__Extension, argument string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptEngineAgent_Extension(C.QtObjectPtr(ptr.Pointer()), C.int(extension), C.CString(argument)))
	}
	return ""
}

func (ptr *QScriptEngineAgent) SupportsExtension(extension QScriptEngineAgent__Extension) bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngineAgent_SupportsExtension(C.QtObjectPtr(ptr.Pointer()), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgent() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_DestroyQScriptEngineAgent(C.QtObjectPtr(ptr.Pointer()))
	}
}
