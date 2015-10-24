package script

//#include "qscriptcontext.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QScriptContext struct {
	ptr unsafe.Pointer
}

type QScriptContextITF interface {
	QScriptContextPTR() *QScriptContext
}

func (p *QScriptContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptContext(ptr QScriptContextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContextPTR().Pointer()
	}
	return nil
}

func QScriptContextFromPointer(ptr unsafe.Pointer) *QScriptContext {
	var n = new(QScriptContext)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptContext) QScriptContextPTR() *QScriptContext {
	return ptr
}

//QScriptContext::Error
type QScriptContext__Error int

var (
	QScriptContext__UnknownError   = QScriptContext__Error(0)
	QScriptContext__ReferenceError = QScriptContext__Error(1)
	QScriptContext__SyntaxError    = QScriptContext__Error(2)
	QScriptContext__TypeError      = QScriptContext__Error(3)
	QScriptContext__RangeError     = QScriptContext__Error(4)
	QScriptContext__URIError       = QScriptContext__Error(5)
)

//QScriptContext::ExecutionState
type QScriptContext__ExecutionState int

var (
	QScriptContext__NormalState    = QScriptContext__ExecutionState(0)
	QScriptContext__ExceptionState = QScriptContext__ExecutionState(1)
)

func (ptr *QScriptContext) ArgumentCount() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptContext_ArgumentCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContext) Backtrace() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptContext_Backtrace(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptContext) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptContext_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptContext) IsCalledAsConstructor() bool {
	if ptr.Pointer() != nil {
		return C.QScriptContext_IsCalledAsConstructor(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptContext) ParentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return QScriptContextFromPointer(unsafe.Pointer(C.QScriptContext_ParentContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptContext) SetActivationObject(activation QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptContext_SetActivationObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(activation)))
	}
}

func (ptr *QScriptContext) SetThisObject(thisObject QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptContext_SetThisObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(thisObject)))
	}
}

func (ptr *QScriptContext) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContext_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptContext) DestroyQScriptContext() {
	if ptr.Pointer() != nil {
		C.QScriptContext_DestroyQScriptContext(C.QtObjectPtr(ptr.Pointer()))
	}
}
