package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QEnableSharedFromThis struct {
	ptr unsafe.Pointer
}

type QEnableSharedFromThis_ITF interface {
	QEnableSharedFromThis_PTR() *QEnableSharedFromThis
}

func (p *QEnableSharedFromThis) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QEnableSharedFromThis) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQEnableSharedFromThis(ptr QEnableSharedFromThis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEnableSharedFromThis_PTR().Pointer()
	}
	return nil
}

func NewQEnableSharedFromThisFromPointer(ptr unsafe.Pointer) *QEnableSharedFromThis {
	var n = new(QEnableSharedFromThis)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEnableSharedFromThis) QEnableSharedFromThis_PTR() *QEnableSharedFromThis {
	return ptr
}
