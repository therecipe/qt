package core

//#include "qdebugstatesaver.h"
import "C"
import (
	"unsafe"
)

type QDebugStateSaver struct {
	ptr unsafe.Pointer
}

type QDebugStateSaverITF interface {
	QDebugStateSaverPTR() *QDebugStateSaver
}

func (p *QDebugStateSaver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDebugStateSaver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDebugStateSaver(ptr QDebugStateSaverITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebugStateSaverPTR().Pointer()
	}
	return nil
}

func QDebugStateSaverFromPointer(ptr unsafe.Pointer) *QDebugStateSaver {
	var n = new(QDebugStateSaver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDebugStateSaver) QDebugStateSaverPTR() *QDebugStateSaver {
	return ptr
}

func NewQDebugStateSaver(dbg QDebugITF) *QDebugStateSaver {
	return QDebugStateSaverFromPointer(unsafe.Pointer(C.QDebugStateSaver_NewQDebugStateSaver(C.QtObjectPtr(PointerFromQDebug(dbg)))))
}

func (ptr *QDebugStateSaver) DestroyQDebugStateSaver() {
	if ptr.Pointer() != nil {
		C.QDebugStateSaver_DestroyQDebugStateSaver(C.QtObjectPtr(ptr.Pointer()))
	}
}
