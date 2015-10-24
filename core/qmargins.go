package core

//#include "qmargins.h"
import "C"
import (
	"unsafe"
)

type QMargins struct {
	ptr unsafe.Pointer
}

type QMarginsITF interface {
	QMarginsPTR() *QMargins
}

func (p *QMargins) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMargins) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMargins(ptr QMarginsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMarginsPTR().Pointer()
	}
	return nil
}

func QMarginsFromPointer(ptr unsafe.Pointer) *QMargins {
	var n = new(QMargins)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMargins) QMarginsPTR() *QMargins {
	return ptr
}

func NewQMargins() *QMargins {
	return QMarginsFromPointer(unsafe.Pointer(C.QMargins_NewQMargins()))
}

func NewQMargins2(left int, top int, right int, bottom int) *QMargins {
	return QMarginsFromPointer(unsafe.Pointer(C.QMargins_NewQMargins2(C.int(left), C.int(top), C.int(right), C.int(bottom))))
}

func (ptr *QMargins) Bottom() int {
	if ptr.Pointer() != nil {
		return int(C.QMargins_Bottom(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMargins) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QMargins_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMargins) Left() int {
	if ptr.Pointer() != nil {
		return int(C.QMargins_Left(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMargins) Right() int {
	if ptr.Pointer() != nil {
		return int(C.QMargins_Right(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMargins) SetBottom(bottom int) {
	if ptr.Pointer() != nil {
		C.QMargins_SetBottom(C.QtObjectPtr(ptr.Pointer()), C.int(bottom))
	}
}

func (ptr *QMargins) SetLeft(left int) {
	if ptr.Pointer() != nil {
		C.QMargins_SetLeft(C.QtObjectPtr(ptr.Pointer()), C.int(left))
	}
}

func (ptr *QMargins) SetRight(right int) {
	if ptr.Pointer() != nil {
		C.QMargins_SetRight(C.QtObjectPtr(ptr.Pointer()), C.int(right))
	}
}

func (ptr *QMargins) SetTop(Top int) {
	if ptr.Pointer() != nil {
		C.QMargins_SetTop(C.QtObjectPtr(ptr.Pointer()), C.int(Top))
	}
}

func (ptr *QMargins) Top() int {
	if ptr.Pointer() != nil {
		return int(C.QMargins_Top(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
