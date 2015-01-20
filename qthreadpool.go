package qt

//#include "qthreadpool.h"
import "C"

type qthreadpool struct {
	qobject
}

type QThreadPool interface {
	QObject
	ActiveThreadCount() int
	Clear()
	ExpiryTimeout() int
	MaxThreadCount() int
	ReleaseThread()
	ReserveThread()
	SetExpiryTimeout(expiryTimeout int)
	SetMaxThreadCount(maxThreadCount int)
	WaitForDone(msecs int) bool
}

func (p *qthreadpool) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qthreadpool) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQThreadPool(parent QObject) QThreadPool {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qthreadpool = new(qthreadpool)
	qthreadpool.SetPointer(C.QThreadPool_New_QObject(parentPtr))
	qthreadpool.SetObjectName("QThreadPool_" + randomIdentifier())
	return qthreadpool
}

func (p *qthreadpool) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QThreadPool_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qthreadpool) ActiveThreadCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QThreadPool_ActiveThreadCount(p.Pointer()))
}

func (p *qthreadpool) Clear() {
	if p.Pointer() != nil {
		C.QThreadPool_Clear(p.Pointer())
	}
}

func (p *qthreadpool) ExpiryTimeout() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QThreadPool_ExpiryTimeout(p.Pointer()))
}

func (p *qthreadpool) MaxThreadCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QThreadPool_MaxThreadCount(p.Pointer()))
}

func (p *qthreadpool) ReleaseThread() {
	if p.Pointer() != nil {
		C.QThreadPool_ReleaseThread(p.Pointer())
	}
}

func (p *qthreadpool) ReserveThread() {
	if p.Pointer() != nil {
		C.QThreadPool_ReserveThread(p.Pointer())
	}
}

func (p *qthreadpool) SetExpiryTimeout(expiryTimeout int) {
	if p.Pointer() != nil {
		C.QThreadPool_SetExpiryTimeout_Int(p.Pointer(), C.int(expiryTimeout))
	}
}

func (p *qthreadpool) SetMaxThreadCount(maxThreadCount int) {
	if p.Pointer() != nil {
		C.QThreadPool_SetMaxThreadCount_Int(p.Pointer(), C.int(maxThreadCount))
	}
}

func (p *qthreadpool) WaitForDone(msecs int) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QThreadPool_WaitForDone_Int(p.Pointer(), C.int(msecs)) != 0
}

func QThreadPool_GlobalInstance() QThreadPool {
	var qthreadpool = new(qthreadpool)
	qthreadpool.SetPointer(C.QThreadPool_GlobalInstance())
	if qthreadpool.ObjectName() == "" {
		qthreadpool.SetObjectName("QThreadPool_" + randomIdentifier())
	}
	return qthreadpool
}
