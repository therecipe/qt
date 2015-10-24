package dbus

//#include "qdbuspendingcall.h"
import "C"
import (
	"unsafe"
)

type QDBusPendingCall struct {
	ptr unsafe.Pointer
}

type QDBusPendingCallITF interface {
	QDBusPendingCallPTR() *QDBusPendingCall
}

func (p *QDBusPendingCall) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusPendingCall) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusPendingCall(ptr QDBusPendingCallITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallPTR().Pointer()
	}
	return nil
}

func QDBusPendingCallFromPointer(ptr unsafe.Pointer) *QDBusPendingCall {
	var n = new(QDBusPendingCall)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusPendingCall) QDBusPendingCallPTR() *QDBusPendingCall {
	return ptr
}

func NewQDBusPendingCall(other QDBusPendingCallITF) *QDBusPendingCall {
	return QDBusPendingCallFromPointer(unsafe.Pointer(C.QDBusPendingCall_NewQDBusPendingCall(C.QtObjectPtr(PointerFromQDBusPendingCall(other)))))
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCallITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCall_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDBusPendingCall(other)))
	}
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCall_DestroyQDBusPendingCall(C.QtObjectPtr(ptr.Pointer()))
	}
}
