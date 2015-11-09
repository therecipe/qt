package dbus

//#include "qdbuspendingreply.h"
import "C"
import (
	"unsafe"
)

type QDBusPendingReply struct {
	QDBusPendingCall
}

type QDBusPendingReply_ITF interface {
	QDBusPendingCall_ITF
	QDBusPendingReply_PTR() *QDBusPendingReply
}

func PointerFromQDBusPendingReply(ptr QDBusPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingReplyFromPointer(ptr unsafe.Pointer) *QDBusPendingReply {
	var n = new(QDBusPendingReply)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusPendingReply) QDBusPendingReply_PTR() *QDBusPendingReply {
	return ptr
}
