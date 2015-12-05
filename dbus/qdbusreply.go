package dbus

//#include "dbus.h"
import "C"
import (
	"unsafe"
)

type QDBusReply struct {
	ptr unsafe.Pointer
}

type QDBusReply_ITF interface {
	QDBusReply_PTR() *QDBusReply
}

func (p *QDBusReply) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusReply) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusReply(ptr QDBusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusReplyFromPointer(ptr unsafe.Pointer) *QDBusReply {
	var n = new(QDBusReply)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusReply) QDBusReply_PTR() *QDBusReply {
	return ptr
}
