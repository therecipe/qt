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

func (ptr *QScriptEngineAgent) Extension(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptEngineAgent::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptEngineAgent_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
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
