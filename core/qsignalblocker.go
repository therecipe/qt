package core

//#include "qsignalblocker.h"
import "C"
import (
	"unsafe"
)

type QSignalBlocker struct {
	ptr unsafe.Pointer
}

type QSignalBlockerITF interface {
	QSignalBlockerPTR() *QSignalBlocker
}

func (p *QSignalBlocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSignalBlocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSignalBlocker(ptr QSignalBlockerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalBlockerPTR().Pointer()
	}
	return nil
}

func QSignalBlockerFromPointer(ptr unsafe.Pointer) *QSignalBlocker {
	var n = new(QSignalBlocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSignalBlocker) QSignalBlockerPTR() *QSignalBlocker {
	return ptr
}

func (ptr *QSignalBlocker) Reblock() {
	if ptr.Pointer() != nil {
		C.QSignalBlocker_Reblock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSignalBlocker) Unblock() {
	if ptr.Pointer() != nil {
		C.QSignalBlocker_Unblock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSignalBlocker) DestroyQSignalBlocker() {
	if ptr.Pointer() != nil {
		C.QSignalBlocker_DestroyQSignalBlocker(C.QtObjectPtr(ptr.Pointer()))
	}
}
