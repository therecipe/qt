package script

//#include "script.h"
import "C"
import (
	"log"
	"unsafe"
)

type QScriptable struct {
	ptr unsafe.Pointer
}

type QScriptable_ITF interface {
	QScriptable_PTR() *QScriptable
}

func (p *QScriptable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptable(ptr QScriptable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptable_PTR().Pointer()
	}
	return nil
}

func NewQScriptableFromPointer(ptr unsafe.Pointer) *QScriptable {
	var n = new(QScriptable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptable) QScriptable_PTR() *QScriptable {
	return ptr
}

func (ptr *QScriptable) Argument(index int) *QScriptValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptable::argument")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptable_Argument(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QScriptable) ArgumentCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptable::argumentCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QScriptable_ArgumentCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptable) Context() *QScriptContext {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptable::context")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptable_Context(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptable) Engine() *QScriptEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptable::engine")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptable_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptable) ThisObject() *QScriptValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptable::thisObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptable_ThisObject(ptr.Pointer()))
	}
	return nil
}
