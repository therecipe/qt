package gui

//#include "qtextimageformat.h"
import "C"
import (
	"unsafe"
)

type QTextImageFormat struct {
	QTextCharFormat
}

type QTextImageFormatITF interface {
	QTextCharFormatITF
	QTextImageFormatPTR() *QTextImageFormat
}

func PointerFromQTextImageFormat(ptr QTextImageFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextImageFormatPTR().Pointer()
	}
	return nil
}

func QTextImageFormatFromPointer(ptr unsafe.Pointer) *QTextImageFormat {
	var n = new(QTextImageFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextImageFormat) QTextImageFormatPTR() *QTextImageFormat {
	return ptr
}

func NewQTextImageFormat() *QTextImageFormat {
	return QTextImageFormatFromPointer(unsafe.Pointer(C.QTextImageFormat_NewQTextImageFormat()))
}

func (ptr *QTextImageFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextImageFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextImageFormat) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextImageFormat_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextImageFormat) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QTextImageFormat_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}
