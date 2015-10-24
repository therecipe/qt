package core

//#include "qsize.h"
import "C"
import (
	"unsafe"
)

type QSize struct {
	ptr unsafe.Pointer
}

type QSizeITF interface {
	QSizePTR() *QSize
}

func (p *QSize) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSize) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSize(ptr QSizeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizePTR().Pointer()
	}
	return nil
}

func QSizeFromPointer(ptr unsafe.Pointer) *QSize {
	var n = new(QSize)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSize) QSizePTR() *QSize {
	return ptr
}

func NewQSize() *QSize {
	return QSizeFromPointer(unsafe.Pointer(C.QSize_NewQSize()))
}

func NewQSize2(width int, height int) *QSize {
	return QSizeFromPointer(unsafe.Pointer(C.QSize_NewQSize2(C.int(width), C.int(height))))
}

func (ptr *QSize) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSize) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSize_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSize) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSize_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSize) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSize_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSize) Rheight() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Rheight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSize) Rwidth() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Rwidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSize) Scale2(size QSizeITF, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSize_Scale2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSize(size)), C.int(mode))
	}
}

func (ptr *QSize) Scale(width int, height int, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSize_Scale(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height), C.int(mode))
	}
}

func (ptr *QSize) SetHeight(height int) {
	if ptr.Pointer() != nil {
		C.QSize_SetHeight(C.QtObjectPtr(ptr.Pointer()), C.int(height))
	}
}

func (ptr *QSize) SetWidth(width int) {
	if ptr.Pointer() != nil {
		C.QSize_SetWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QSize) Transpose() {
	if ptr.Pointer() != nil {
		C.QSize_Transpose(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSize) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QSize_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
