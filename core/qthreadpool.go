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

type QThreadPoolITF interface {
	QObjectITF
	QThreadPoolPTR() *QThreadPool
}

func PointerFromQThreadPool(ptr QThreadPoolITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThreadPoolPTR().Pointer()
	}
	return nil
}

func QThreadPoolFromPointer(ptr unsafe.Pointer) *QThreadPool {
	var n = new(QThreadPool)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QThreadPool_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QThreadPool) QThreadPoolPTR() *QThreadPool {
	return ptr
}

func (ptr *QThreadPool) ActiveThreadCount() int {
	if ptr.Pointer() != nil {
		return int(C.QThreadPool_ActiveThreadCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QThreadPool) ExpiryTimeout() int {
	if ptr.Pointer() != nil {
		return int(C.QThreadPool_ExpiryTimeout(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QThreadPool) MaxThreadCount() int {
	if ptr.Pointer() != nil {
		return int(C.QThreadPool_MaxThreadCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	if ptr.Pointer() != nil {
		C.QThreadPool_SetExpiryTimeout(C.QtObjectPtr(ptr.Pointer()), C.int(expiryTimeout))
	}
}

func (ptr *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	if ptr.Pointer() != nil {
		C.QThreadPool_SetMaxThreadCount(C.QtObjectPtr(ptr.Pointer()), C.int(maxThreadCount))
	}
}

func NewQThreadPool(parent QObjectITF) *QThreadPool {
	return QThreadPoolFromPointer(unsafe.Pointer(C.QThreadPool_NewQThreadPool(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QThreadPool) Cancel(runnable QRunnableITF) {
	if ptr.Pointer() != nil {
		C.QThreadPool_Cancel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRunnable(runnable)))
	}
}

func (ptr *QThreadPool) Clear() {
	if ptr.Pointer() != nil {
		C.QThreadPool_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QThreadPool_GlobalInstance() *QThreadPool {
	return QThreadPoolFromPointer(unsafe.Pointer(C.QThreadPool_QThreadPool_GlobalInstance()))
}

func (ptr *QThreadPool) ReleaseThread() {
	if ptr.Pointer() != nil {
		C.QThreadPool_ReleaseThread(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QThreadPool) ReserveThread() {
	if ptr.Pointer() != nil {
		C.QThreadPool_ReserveThread(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QThreadPool) Start(runnable QRunnableITF, priority int) {
	if ptr.Pointer() != nil {
		C.QThreadPool_Start(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRunnable(runnable)), C.int(priority))
	}
}

func (ptr *QThreadPool) TryStart(runnable QRunnableITF) bool {
	if ptr.Pointer() != nil {
		return C.QThreadPool_TryStart(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRunnable(runnable))) != 0
	}
	return false
}

func (ptr *QThreadPool) WaitForDone(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QThreadPool_WaitForDone(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QThreadPool) DestroyQThreadPool() {
	if ptr.Pointer() != nil {
		C.QThreadPool_DestroyQThreadPool(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
