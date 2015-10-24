package core

//#include "qmarginsf.h"
import "C"
import (
	"unsafe"
)

type QMarginsF struct {
	ptr unsafe.Pointer
}

type QMarginsFITF interface {
	QMarginsFPTR() *QMarginsF
}

func (p *QMarginsF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMarginsF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMarginsF(ptr QMarginsFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMarginsFPTR().Pointer()
	}
	return nil
}

func QMarginsFFromPointer(ptr unsafe.Pointer) *QMarginsF {
	var n = new(QMarginsF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMarginsF) QMarginsFPTR() *QMarginsF {
	return ptr
}

func NewQMarginsF() *QMarginsF {
	return QMarginsFFromPointer(unsafe.Pointer(C.QMarginsF_NewQMarginsF()))
}

func NewQMarginsF3(margins QMarginsITF) *QMarginsF {
	return QMarginsFFromPointer(unsafe.Pointer(C.QMarginsF_NewQMarginsF3(C.QtObjectPtr(PointerFromQMargins(margins)))))
}

func (ptr *QMarginsF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QMarginsF_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
