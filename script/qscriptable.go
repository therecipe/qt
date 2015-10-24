package script

//#include "qscriptable.h"
import "C"
import (
	"unsafe"
)

type QScriptable struct {
	ptr unsafe.Pointer
}

type QScriptableITF interface {
	QScriptablePTR() *QScriptable
}

func (p *QScriptable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptable(ptr QScriptableITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptablePTR().Pointer()
	}
	return nil
}

func QScriptableFromPointer(ptr unsafe.Pointer) *QScriptable {
	var n = new(QScriptable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptable) QScriptablePTR() *QScriptable {
	return ptr
}

func (ptr *QScriptable) ArgumentCount() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptable_ArgumentCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptable) Context() *QScriptContext {
	if ptr.Pointer() != nil {
		return QScriptContextFromPointer(unsafe.Pointer(C.QScriptable_Context(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptable) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptable_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
