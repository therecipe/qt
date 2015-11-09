package core

//#include "qeventlooplocker.h"
import "C"
import (
	"unsafe"
)

type QEventLoopLocker struct {
	ptr unsafe.Pointer
}

type QEventLoopLocker_ITF interface {
	QEventLoopLocker_PTR() *QEventLoopLocker
}

func (p *QEventLoopLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QEventLoopLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQEventLoopLocker(ptr QEventLoopLocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventLoopLocker_PTR().Pointer()
	}
	return nil
}

func NewQEventLoopLockerFromPointer(ptr unsafe.Pointer) *QEventLoopLocker {
	var n = new(QEventLoopLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEventLoopLocker) QEventLoopLocker_PTR() *QEventLoopLocker {
	return ptr
}

func NewQEventLoopLocker() *QEventLoopLocker {
	return NewQEventLoopLockerFromPointer(C.QEventLoopLocker_NewQEventLoopLocker())
}

func NewQEventLoopLocker2(loop QEventLoop_ITF) *QEventLoopLocker {
	return NewQEventLoopLockerFromPointer(C.QEventLoopLocker_NewQEventLoopLocker2(PointerFromQEventLoop(loop)))
}

func NewQEventLoopLocker3(thread QThread_ITF) *QEventLoopLocker {
	return NewQEventLoopLockerFromPointer(C.QEventLoopLocker_NewQEventLoopLocker3(PointerFromQThread(thread)))
}

func (ptr *QEventLoopLocker) DestroyQEventLoopLocker() {
	if ptr.Pointer() != nil {
		C.QEventLoopLocker_DestroyQEventLoopLocker(ptr.Pointer())
	}
}
