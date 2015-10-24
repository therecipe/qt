package dbus

//#include "qdbuspendingreply.h"
import "C"
import (
	"unsafe"
)

type QDBusPendingReply struct {
	QDBusPendingCall
}

type QDBusPendingReplyITF interface {
	QDBusPendingCallITF
	QDBusPendingReplyPTR() *QDBusPendingReply
}

func PointerFromQDBusPendingReply(ptr QDBusPendingReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReplyPTR().Pointer()
	}
	return nil
}

func QDBusPendingReplyFromPointer(ptr unsafe.Pointer) *QDBusPendingReply {
	var n = new(QDBusPendingReply)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusPendingReply) QDBusPendingReplyPTR() *QDBusPendingReply {
	return ptr
}
