package gui

//#include "qtextlistformat.h"
import "C"
import (
	"unsafe"
)

type QTextListFormat struct {
	QTextFormat
}

type QTextListFormatITF interface {
	QTextFormatITF
	QTextListFormatPTR() *QTextListFormat
}

func PointerFromQTextListFormat(ptr QTextListFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextListFormatPTR().Pointer()
	}
	return nil
}

func QTextListFormatFromPointer(ptr unsafe.Pointer) *QTextListFormat {
	var n = new(QTextListFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextListFormat) QTextListFormatPTR() *QTextListFormat {
	return ptr
}

//QTextListFormat::Style
type QTextListFormat__Style int

var (
	QTextListFormat__ListDisc           = QTextListFormat__Style(-1)
	QTextListFormat__ListCircle         = QTextListFormat__Style(-2)
	QTextListFormat__ListSquare         = QTextListFormat__Style(-3)
	QTextListFormat__ListDecimal        = QTextListFormat__Style(-4)
	QTextListFormat__ListLowerAlpha     = QTextListFormat__Style(-5)
	QTextListFormat__ListUpperAlpha     = QTextListFormat__Style(-6)
	QTextListFormat__ListLowerRoman     = QTextListFormat__Style(-7)
	QTextListFormat__ListUpperRoman     = QTextListFormat__Style(-8)
	QTextListFormat__ListStyleUndefined = QTextListFormat__Style(0)
)

func NewQTextListFormat() *QTextListFormat {
	return QTextListFormatFromPointer(unsafe.Pointer(C.QTextListFormat_NewQTextListFormat()))
}

func (ptr *QTextListFormat) Indent() int {
	if ptr.Pointer() != nil {
		return int(C.QTextListFormat_Indent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextListFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextListFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextListFormat) NumberPrefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextListFormat_NumberPrefix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextListFormat) NumberSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextListFormat_NumberSuffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextListFormat) SetIndent(indentation int) {
	if ptr.Pointer() != nil {
		C.QTextListFormat_SetIndent(C.QtObjectPtr(ptr.Pointer()), C.int(indentation))
	}
}

func (ptr *QTextListFormat) SetNumberPrefix(numberPrefix string) {
	if ptr.Pointer() != nil {
		C.QTextListFormat_SetNumberPrefix(C.QtObjectPtr(ptr.Pointer()), C.CString(numberPrefix))
	}
}

func (ptr *QTextListFormat) SetNumberSuffix(numberSuffix string) {
	if ptr.Pointer() != nil {
		C.QTextListFormat_SetNumberSuffix(C.QtObjectPtr(ptr.Pointer()), C.CString(numberSuffix))
	}
}

func (ptr *QTextListFormat) SetStyle(style QTextListFormat__Style) {
	if ptr.Pointer() != nil {
		C.QTextListFormat_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QTextListFormat) Style() QTextListFormat__Style {
	if ptr.Pointer() != nil {
		return QTextListFormat__Style(C.QTextListFormat_Style(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
