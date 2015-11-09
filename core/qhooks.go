package core

//#include "qhooks.h"
import "C"
import (
	"unsafe"
)

type QHooks struct {
	ptr unsafe.Pointer
}

type QHooks_ITF interface {
	QHooks_PTR() *QHooks
}

func (p *QHooks) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHooks) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHooks(ptr QHooks_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHooks_PTR().Pointer()
	}
	return nil
}

func NewQHooksFromPointer(ptr unsafe.Pointer) *QHooks {
	var n = new(QHooks)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHooks) QHooks_PTR() *QHooks {
	return ptr
}

//QHooks::HookIndex
type QHooks__HookIndex int64

const (
	QHooks__HookDataVersion = QHooks__HookIndex(0)
	QHooks__HookDataSize    = QHooks__HookIndex(1)
	QHooks__QtVersion       = QHooks__HookIndex(2)
	QHooks__AddQObject      = QHooks__HookIndex(3)
	QHooks__RemoveQObject   = QHooks__HookIndex(4)
	QHooks__Startup         = QHooks__HookIndex(5)
	QHooks__LastHookIndex   = QHooks__HookIndex(6)
)
