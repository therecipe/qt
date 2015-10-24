package core

//#include "qeventlooplocker.h"
import "C"
import (
	"unsafe"
)

type QEventLoopLocker struct {
	ptr unsafe.Pointer
}

type QEventLoopLockerITF interface {
	QEventLoopLockerPTR() *QEventLoopLocker
}

func (p *QEventLoopLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QEventLoopLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQEventLoopLocker(ptr QEventLoopLockerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventLoopLockerPTR().Pointer()
	}
	return nil
}

func QEventLoopLockerFromPointer(ptr unsafe.Pointer) *QEventLoopLocker {
	var n = new(QEventLoopLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEventLoopLocker) QEventLoopLockerPTR() *QEventLoopLocker {
	return ptr
}

func NewQEventLoopLocker() *QEventLoopLocker {
	return QEventLoopLockerFromPointer(unsafe.Pointer(C.QEventLoopLocker_NewQEventLoopLocker()))
}

func NewQEventLoopLocker2(loop QEventLoopITF) *QEventLoopLocker {
	return QEventLoopLockerFromPointer(unsafe.Pointer(C.QEventLoopLocker_NewQEventLoopLocker2(C.QtObjectPtr(PointerFromQEventLoop(loop)))))
}

func NewQEventLoopLocker3(thread QThreadITF) *QEventLoopLocker {
	return QEventLoopLockerFromPointer(unsafe.Pointer(C.QEventLoopLocker_NewQEventLoopLocker3(C.QtObjectPtr(PointerFromQThread(thread)))))
}

func (ptr *QEventLoopLocker) DestroyQEventLoopLocker() {
	if ptr.Pointer() != nil {
		C.QEventLoopLocker_DestroyQEventLoopLocker(C.QtObjectPtr(ptr.Pointer()))
	}
}
