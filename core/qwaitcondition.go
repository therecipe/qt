package core

//#include "qwaitcondition.h"
import "C"
import (
	"unsafe"
)

type QWaitCondition struct {
	ptr unsafe.Pointer
}

type QWaitConditionITF interface {
	QWaitConditionPTR() *QWaitCondition
}

func (p *QWaitCondition) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWaitCondition) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWaitCondition(ptr QWaitConditionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWaitConditionPTR().Pointer()
	}
	return nil
}

func QWaitConditionFromPointer(ptr unsafe.Pointer) *QWaitCondition {
	var n = new(QWaitCondition)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWaitCondition) QWaitConditionPTR() *QWaitCondition {
	return ptr
}

func NewQWaitCondition() *QWaitCondition {
	return QWaitConditionFromPointer(unsafe.Pointer(C.QWaitCondition_NewQWaitCondition()))
}

func (ptr *QWaitCondition) WakeAll() {
	if ptr.Pointer() != nil {
		C.QWaitCondition_WakeAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWaitCondition) WakeOne() {
	if ptr.Pointer() != nil {
		C.QWaitCondition_WakeOne(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWaitCondition) DestroyQWaitCondition() {
	if ptr.Pointer() != nil {
		C.QWaitCondition_DestroyQWaitCondition(C.QtObjectPtr(ptr.Pointer()))
	}
}
