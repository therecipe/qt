package gui

//#include "qtextimageformat.h"
import "C"
import (
	"unsafe"
)

type QTextImageFormat struct {
	QTextCharFormat
}

type QTextImageFormat_ITF interface {
	QTextCharFormat_ITF
	QTextImageFormat_PTR() *QTextImageFormat
}

func PointerFromQTextImageFormat(ptr QTextImageFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextImageFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextImageFormatFromPointer(ptr unsafe.Pointer) *QTextImageFormat {
	var n = new(QTextImageFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextImageFormat) QTextImageFormat_PTR() *QTextImageFormat {
	return ptr
}

func NewQTextImageFormat() *QTextImageFormat {
	return NewQTextImageFormatFromPointer(C.QTextImageFormat_NewQTextImageFormat())
}

func (ptr *QTextImageFormat) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextImageFormat_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextImageFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextImageFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextImageFormat) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextImageFormat_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextImageFormat) SetHeight(height float64) {
	if ptr.Pointer() != nil {
		C.QTextImageFormat_SetHeight(ptr.Pointer(), C.double(height))
	}
}

func (ptr *QTextImageFormat) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QTextImageFormat_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTextImageFormat) SetWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QTextImageFormat_SetWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextImageFormat) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextImageFormat_Width(ptr.Pointer()))
	}
	return 0
}
