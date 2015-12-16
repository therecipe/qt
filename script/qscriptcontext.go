package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

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
		return strings.Split(C.GoString(C.QScriptContext_Backtrace(ptr.Pointer())), ",,,")
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
