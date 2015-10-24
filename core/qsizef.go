package core

//#include "qsizef.h"
import "C"
import (
	"unsafe"
)

type QSizeF struct {
	ptr unsafe.Pointer
}

type QSizeFITF interface {
	QSizeFPTR() *QSizeF
}

func (p *QSizeF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSizeF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSizeF(ptr QSizeFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeFPTR().Pointer()
	}
	return nil
}

func QSizeFFromPointer(ptr unsafe.Pointer) *QSizeF {
	var n = new(QSizeF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSizeF) QSizeFPTR() *QSizeF {
	return ptr
}

func NewQSizeF() *QSizeF {
	return QSizeFFromPointer(unsafe.Pointer(C.QSizeF_NewQSizeF()))
}

func NewQSizeF2(size QSizeITF) *QSizeF {
	return QSizeFFromPointer(unsafe.Pointer(C.QSizeF_NewQSizeF2(C.QtObjectPtr(PointerFromQSize(size)))))
}

func (ptr *QSizeF) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSizeF_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSizeF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSizeF_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSizeF) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSizeF_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSizeF) Scale2(size QSizeFITF, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSizeF_Scale2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSizeF(size)), C.int(mode))
	}
}

func (ptr *QSizeF) Transpose() {
	if ptr.Pointer() != nil {
		C.QSizeF_Transpose(C.QtObjectPtr(ptr.Pointer()))
	}
}
