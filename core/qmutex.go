package core

//#include "qmutex.h"
import "C"
import (
	"unsafe"
)

type QMutex struct {
	ptr unsafe.Pointer
}

type QMutexITF interface {
	QMutexPTR() *QMutex
}

func (p *QMutex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutex(ptr QMutexITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutexPTR().Pointer()
	}
	return nil
}

func QMutexFromPointer(ptr unsafe.Pointer) *QMutex {
	var n = new(QMutex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutex) QMutexPTR() *QMutex {
	return ptr
}

//QMutex::RecursionMode
type QMutex__RecursionMode int

var (
	QMutex__NonRecursive = QMutex__RecursionMode(0)
	QMutex__Recursive    = QMutex__RecursionMode(1)
)

func (ptr *QMutex) Lock() {
	if ptr.Pointer() != nil {
		C.QMutex_Lock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMutex) TryLock(timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QMutex_TryLock(C.QtObjectPtr(ptr.Pointer()), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QMutex) Unlock() {
	if ptr.Pointer() != nil {
		C.QMutex_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func NewQMutex(mode QMutex__RecursionMode) *QMutex {
	return QMutexFromPointer(unsafe.Pointer(C.QMutex_NewQMutex(C.int(mode))))
}

func (ptr *QMutex) IsRecursive() bool {
	if ptr.Pointer() != nil {
		return C.QMutex_IsRecursive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
