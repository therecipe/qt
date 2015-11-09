package core

//#include "qthreadpool.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QThreadPool struct {
	QObject
}

type QThreadPool_ITF interface {
	QObject_ITF
	QThreadPool_PTR() *QThreadPool
}

func PointerFromQThreadPool(ptr QThreadPool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThreadPool_PTR().Pointer()
	}
	return nil
}

func NewQThreadPoolFromPointer(ptr unsafe.Pointer) *QThreadPool {
	var n = new(QThreadPool)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QThreadPool_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QThreadPool) QThreadPool_PTR() *QThreadPool {
	return ptr
}

func (ptr *QThreadPool) ActiveThreadCount() int {
	if ptr.Pointer() != nil {
		return int(C.QThreadPool_ActiveThreadCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThreadPool) ExpiryTimeout() int {
	if ptr.Pointer() != nil {
		return int(C.QThreadPool_ExpiryTimeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThreadPool) MaxThreadCount() int {
	if ptr.Pointer() != nil {
		return int(C.QThreadPool_MaxThreadCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	if ptr.Pointer() != nil {
		C.QThreadPool_SetExpiryTimeout(ptr.Pointer(), C.int(expiryTimeout))
	}
}

func (ptr *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	if ptr.Pointer() != nil {
		C.QThreadPool_SetMaxThreadCount(ptr.Pointer(), C.int(maxThreadCount))
	}
}

func NewQThreadPool(parent QObject_ITF) *QThreadPool {
	return NewQThreadPoolFromPointer(C.QThreadPool_NewQThreadPool(PointerFromQObject(parent)))
}

func (ptr *QThreadPool) Cancel(runnable QRunnable_ITF) {
	if ptr.Pointer() != nil {
		C.QThreadPool_Cancel(ptr.Pointer(), PointerFromQRunnable(runnable))
	}
}

func (ptr *QThreadPool) Clear() {
	if ptr.Pointer() != nil {
		C.QThreadPool_Clear(ptr.Pointer())
	}
}

func QThreadPool_GlobalInstance() *QThreadPool {
	return NewQThreadPoolFromPointer(C.QThreadPool_QThreadPool_GlobalInstance())
}

func (ptr *QThreadPool) ReleaseThread() {
	if ptr.Pointer() != nil {
		C.QThreadPool_ReleaseThread(ptr.Pointer())
	}
}

func (ptr *QThreadPool) ReserveThread() {
	if ptr.Pointer() != nil {
		C.QThreadPool_ReserveThread(ptr.Pointer())
	}
}

func (ptr *QThreadPool) Start(runnable QRunnable_ITF, priority int) {
	if ptr.Pointer() != nil {
		C.QThreadPool_Start(ptr.Pointer(), PointerFromQRunnable(runnable), C.int(priority))
	}
}

func (ptr *QThreadPool) TryStart(runnable QRunnable_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QThreadPool_TryStart(ptr.Pointer(), PointerFromQRunnable(runnable)) != 0
	}
	return false
}

func (ptr *QThreadPool) WaitForDone(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QThreadPool_WaitForDone(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QThreadPool) DestroyQThreadPool() {
	if ptr.Pointer() != nil {
		C.QThreadPool_DestroyQThreadPool(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
