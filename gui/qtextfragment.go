package gui

//#include "qtextfragment.h"
import "C"
import (
	"unsafe"
)

type QTextFragment struct {
	ptr unsafe.Pointer
}

type QTextFragmentITF interface {
	QTextFragmentPTR() *QTextFragment
}

func (p *QTextFragment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextFragment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextFragment(ptr QTextFragmentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFragmentPTR().Pointer()
	}
	return nil
}

func QTextFragmentFromPointer(ptr unsafe.Pointer) *QTextFragment {
	var n = new(QTextFragment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextFragment) QTextFragmentPTR() *QTextFragment {
	return ptr
}

func NewQTextFragment() *QTextFragment {
	return QTextFragmentFromPointer(unsafe.Pointer(C.QTextFragment_NewQTextFragment()))
}

func NewQTextFragment3(other QTextFragmentITF) *QTextFragment {
	return QTextFragmentFromPointer(unsafe.Pointer(C.QTextFragment_NewQTextFragment3(C.QtObjectPtr(PointerFromQTextFragment(other)))))
}

func (ptr *QTextFragment) CharFormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFragment_CharFormatIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFragment) Contains(position int) bool {
	if ptr.Pointer() != nil {
		return C.QTextFragment_Contains(C.QtObjectPtr(ptr.Pointer()), C.int(position)) != 0
	}
	return false
}

func (ptr *QTextFragment) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextFragment_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextFragment) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFragment_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFragment) Position() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFragment_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFragment) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextFragment_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
