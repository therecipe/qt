package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QDebugStateSaver::QDebugStateSaver")

	return NewQDebugStateSaverFromPointer(C.QDebugStateSaver_NewQDebugStateSaver(PointerFromQDebug(dbg)))
}

func (ptr *QDebugStateSaver) DestroyQDebugStateSaver() {
	defer qt.Recovering("QDebugStateSaver::~QDebugStateSaver")

	if ptr.Pointer() != nil {
		C.QDebugStateSaver_DestroyQDebugStateSaver(ptr.Pointer())
	}
}
