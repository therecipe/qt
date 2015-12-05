package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::QScriptEngineAgent")
		}
	}()

	return NewQScriptEngineAgentFromPointer(C.QScriptEngineAgent_NewQScriptEngineAgent(PointerFromQScriptEngine(engine)))
}

func (ptr *QScriptEngineAgent) ContextPop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::contextPop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPop(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) ContextPush() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::contextPush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPush(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) Engine() *QScriptEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::engine")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptEngineAgent_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngineAgent) Extension(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::extension")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptEngineAgent_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

func (ptr *QScriptEngineAgent) SupportsExtension(extension QScriptEngineAgent__Extension) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::supportsExtension")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QScriptEngineAgent_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptEngineAgent::~QScriptEngineAgent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_DestroyQScriptEngineAgent(ptr.Pointer())
	}
}
