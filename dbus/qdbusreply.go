package dbus

//#include "qdbusreply.h"
import "C"
import (
	"unsafe"
)

type QDBusReply struct {
	ptr unsafe.Pointer
}

type QDBusReplyITF interface {
	QDBusReplyPTR() *QDBusReply
}

func (p *QDBusReply) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusReply) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusReply(ptr QDBusReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusReplyPTR().Pointer()
	}
	return nil
}

func QDBusReplyFromPointer(ptr unsafe.Pointer) *QDBusReply {
	var n = new(QDBusReply)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusReply) QDBusReplyPTR() *QDBusReply {
	return ptr
}
