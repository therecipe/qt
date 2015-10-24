package core

//#include "qenablesharedfromthis.h"
import "C"
import (
	"unsafe"
)

type QEnableSharedFromThis struct {
	ptr unsafe.Pointer
}

type QEnableSharedFromThisITF interface {
	QEnableSharedFromThisPTR() *QEnableSharedFromThis
}

func (p *QEnableSharedFromThis) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QEnableSharedFromThis) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQEnableSharedFromThis(ptr QEnableSharedFromThisITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEnableSharedFromThisPTR().Pointer()
	}
	return nil
}

func QEnableSharedFromThisFromPointer(ptr unsafe.Pointer) *QEnableSharedFromThis {
	var n = new(QEnableSharedFromThis)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEnableSharedFromThis) QEnableSharedFromThisPTR() *QEnableSharedFromThis {
	return ptr
}
