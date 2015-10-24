package core

//#include "qrunnable.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRunnable struct {
	ptr unsafe.Pointer
}

type QRunnableITF interface {
	QRunnablePTR() *QRunnable
}

func (p *QRunnable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRunnable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRunnable(ptr QRunnableITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRunnablePTR().Pointer()
	}
	return nil
}

func QRunnableFromPointer(ptr unsafe.Pointer) *QRunnable {
	var n = new(QRunnable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRunnable) QRunnablePTR() *QRunnable {
	return ptr
}

func (ptr *QRunnable) AutoDelete() bool {
	if ptr.Pointer() != nil {
		return C.QRunnable_AutoDelete(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRunnable) Run() {
	if ptr.Pointer() != nil {
		C.QRunnable_Run(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRunnable) SetAutoDelete(autoDelete bool) {
	if ptr.Pointer() != nil {
		C.QRunnable_SetAutoDelete(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(autoDelete)))
	}
}

func (ptr *QRunnable) DestroyQRunnable() {
	if ptr.Pointer() != nil {
		C.QRunnable_DestroyQRunnable(C.QtObjectPtr(ptr.Pointer()))
	}
}
