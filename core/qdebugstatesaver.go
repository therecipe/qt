package core

//#include "qdebugstatesaver.h"
import "C"
import (
	"unsafe"
)

type QDebugStateSaver struct {
	ptr unsafe.Pointer
}

type QDebugStateSaver_ITF interface {
	QDebugStateSaver_PTR() *QDebugStateSaver
}

func (p *QDebugStateSaver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDebugStateSaver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDebugStateSaver(ptr QDebugStateSaver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebugStateSaver_PTR().Pointer()
	}
	return nil
}

func NewQDebugStateSaverFromPointer(ptr unsafe.Pointer) *QDebugStateSaver {
	var n = new(QDebugStateSaver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDebugStateSaver) QDebugStateSaver_PTR() *QDebugStateSaver {
	return ptr
}

func NewQDebugStateSaver(dbg QDebug_ITF) *QDebugStateSaver {
	return NewQDebugStateSaverFromPointer(C.QDebugStateSaver_NewQDebugStateSaver(PointerFromQDebug(dbg)))
}

func (ptr *QDebugStateSaver) DestroyQDebugStateSaver() {
	if ptr.Pointer() != nil {
		C.QDebugStateSaver_DestroyQDebugStateSaver(ptr.Pointer())
	}
}
