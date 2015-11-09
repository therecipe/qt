package core

//#include "qsize.h"
import "C"
import (
	"unsafe"
)

type QSize struct {
	ptr unsafe.Pointer
}

type QSize_ITF interface {
	QSize_PTR() *QSize
}

func (p *QSize) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSize) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSize(ptr QSize_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSize_PTR().Pointer()
	}
	return nil
}

func NewQSizeFromPointer(ptr unsafe.Pointer) *QSize {
	var n = new(QSize)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSize) QSize_PTR() *QSize {
	return ptr
}

func NewQSize() *QSize {
	return NewQSizeFromPointer(C.QSize_NewQSize())
}

func NewQSize2(width int, height int) *QSize {
	return NewQSizeFromPointer(C.QSize_NewQSize2(C.int(width), C.int(height)))
}

func (ptr *QSize) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSize) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSize_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSize) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSize_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSize) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSize_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSize) Rheight() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Rheight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSize) Rwidth() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Rwidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSize) Scale2(size QSize_ITF, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSize_Scale2(ptr.Pointer(), PointerFromQSize(size), C.int(mode))
	}
}

func (ptr *QSize) Scale(width int, height int, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSize_Scale(ptr.Pointer(), C.int(width), C.int(height), C.int(mode))
	}
}

func (ptr *QSize) SetHeight(height int) {
	if ptr.Pointer() != nil {
		C.QSize_SetHeight(ptr.Pointer(), C.int(height))
	}
}

func (ptr *QSize) SetWidth(width int) {
	if ptr.Pointer() != nil {
		C.QSize_SetWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QSize) Transpose() {
	if ptr.Pointer() != nil {
		C.QSize_Transpose(ptr.Pointer())
	}
}

func (ptr *QSize) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Width(ptr.Pointer()))
	}
	return 0
}
