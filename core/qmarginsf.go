package core

//#include "qmarginsf.h"
import "C"
import (
	"unsafe"
)

type QMarginsF struct {
	ptr unsafe.Pointer
}

type QMarginsF_ITF interface {
	QMarginsF_PTR() *QMarginsF
}

func (p *QMarginsF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMarginsF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMarginsF(ptr QMarginsF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMarginsF_PTR().Pointer()
	}
	return nil
}

func NewQMarginsFFromPointer(ptr unsafe.Pointer) *QMarginsF {
	var n = new(QMarginsF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMarginsF) QMarginsF_PTR() *QMarginsF {
	return ptr
}

func NewQMarginsF() *QMarginsF {
	return NewQMarginsFFromPointer(C.QMarginsF_NewQMarginsF())
}

func NewQMarginsF3(margins QMargins_ITF) *QMarginsF {
	return NewQMarginsFFromPointer(C.QMarginsF_NewQMarginsF3(PointerFromQMargins(margins)))
}

func NewQMarginsF2(left float64, top float64, right float64, bottom float64) *QMarginsF {
	return NewQMarginsFFromPointer(C.QMarginsF_NewQMarginsF2(C.double(left), C.double(top), C.double(right), C.double(bottom)))
}

func (ptr *QMarginsF) Bottom() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMarginsF_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMarginsF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QMarginsF_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMarginsF) Left() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMarginsF_Left(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMarginsF) Right() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMarginsF_Right(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMarginsF) SetBottom(bottom float64) {
	if ptr.Pointer() != nil {
		C.QMarginsF_SetBottom(ptr.Pointer(), C.double(bottom))
	}
}

func (ptr *QMarginsF) SetLeft(left float64) {
	if ptr.Pointer() != nil {
		C.QMarginsF_SetLeft(ptr.Pointer(), C.double(left))
	}
}

func (ptr *QMarginsF) SetRight(right float64) {
	if ptr.Pointer() != nil {
		C.QMarginsF_SetRight(ptr.Pointer(), C.double(right))
	}
}

func (ptr *QMarginsF) SetTop(Top float64) {
	if ptr.Pointer() != nil {
		C.QMarginsF_SetTop(ptr.Pointer(), C.double(Top))
	}
}

func (ptr *QMarginsF) Top() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMarginsF_Top(ptr.Pointer()))
	}
	return 0
}
