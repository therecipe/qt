package core

//#include "qsignalblocker.h"
import "C"
import (
	"unsafe"
)

type QSignalBlocker struct {
	ptr unsafe.Pointer
}

type QSignalBlocker_ITF interface {
	QSignalBlocker_PTR() *QSignalBlocker
}

func (p *QSignalBlocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSignalBlocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSignalBlocker(ptr QSignalBlocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalBlocker_PTR().Pointer()
	}
	return nil
}

func NewQSignalBlockerFromPointer(ptr unsafe.Pointer) *QSignalBlocker {
	var n = new(QSignalBlocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSignalBlocker) QSignalBlocker_PTR() *QSignalBlocker {
	return ptr
}

func (ptr *QSignalBlocker) Reblock() {
	if ptr.Pointer() != nil {
		C.QSignalBlocker_Reblock(ptr.Pointer())
	}
}

func (ptr *QSignalBlocker) Unblock() {
	if ptr.Pointer() != nil {
		C.QSignalBlocker_Unblock(ptr.Pointer())
	}
}

func (ptr *QSignalBlocker) DestroyQSignalBlocker() {
	if ptr.Pointer() != nil {
		C.QSignalBlocker_DestroyQSignalBlocker(ptr.Pointer())
	}
}
