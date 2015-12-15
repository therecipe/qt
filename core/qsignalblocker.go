package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSignalBlocker struct {
	ptr unsafe.Pointer
}

type QSignalBlocker_ITF interface {
	QSignalBlocker_PTR() *QSignalBlocker
}

func (p *QSignalBlocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSignalBlocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSignalBlocker(ptr QSignalBlocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalBlocker_PTR().Pointer()
	}
	return nil
}

func NewQSignalBlockerFromPointer(ptr unsafe.Pointer) *QSignalBlocker {
	var n = new(QSignalBlocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSignalBlocker) QSignalBlocker_PTR() *QSignalBlocker {
	return ptr
}

func (ptr *QSignalBlocker) Reblock() {
	defer qt.Recovering("QSignalBlocker::reblock")

	if ptr.Pointer() != nil {
		C.QSignalBlocker_Reblock(ptr.Pointer())
	}
}

func (ptr *QSignalBlocker) Unblock() {
	defer qt.Recovering("QSignalBlocker::unblock")

	if ptr.Pointer() != nil {
		C.QSignalBlocker_Unblock(ptr.Pointer())
	}
}

func (ptr *QSignalBlocker) DestroyQSignalBlocker() {
	defer qt.Recovering("QSignalBlocker::~QSignalBlocker")

	if ptr.Pointer() != nil {
		C.QSignalBlocker_DestroyQSignalBlocker(ptr.Pointer())
	}
}
