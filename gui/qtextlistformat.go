package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextListFormat struct {
	QTextFormat
}

type QTextListFormat_ITF interface {
	QTextFormat_ITF
	QTextListFormat_PTR() *QTextListFormat
}

func PointerFromQTextListFormat(ptr QTextListFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextListFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextListFormatFromPointer(ptr unsafe.Pointer) *QTextListFormat {
	var n = new(QTextListFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextListFormat) QTextListFormat_PTR() *QTextListFormat {
	return ptr
}

//QTextListFormat::Style
type QTextListFormat__Style int64

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
	defer qt.Recovering("QTextListFormat::QTextListFormat")

	return NewQTextListFormatFromPointer(C.QTextListFormat_NewQTextListFormat())
}

func (ptr *QTextListFormat) Indent() int {
	defer qt.Recovering("QTextListFormat::indent")

	if ptr.Pointer() != nil {
		return int(C.QTextListFormat_Indent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextListFormat) IsValid() bool {
	defer qt.Recovering("QTextListFormat::isValid")

	if ptr.Pointer() != nil {
		return C.QTextListFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextListFormat) NumberPrefix() string {
	defer qt.Recovering("QTextListFormat::numberPrefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextListFormat_NumberPrefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextListFormat) NumberSuffix() string {
	defer qt.Recovering("QTextListFormat::numberSuffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextListFormat_NumberSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextListFormat) SetIndent(indentation int) {
	defer qt.Recovering("QTextListFormat::setIndent")

	if ptr.Pointer() != nil {
		C.QTextListFormat_SetIndent(ptr.Pointer(), C.int(indentation))
	}
}

func (ptr *QTextListFormat) SetNumberPrefix(numberPrefix string) {
	defer qt.Recovering("QTextListFormat::setNumberPrefix")

	if ptr.Pointer() != nil {
		C.QTextListFormat_SetNumberPrefix(ptr.Pointer(), C.CString(numberPrefix))
	}
}

func (ptr *QTextListFormat) SetNumberSuffix(numberSuffix string) {
	defer qt.Recovering("QTextListFormat::setNumberSuffix")

	if ptr.Pointer() != nil {
		C.QTextListFormat_SetNumberSuffix(ptr.Pointer(), C.CString(numberSuffix))
	}
}

func (ptr *QTextListFormat) SetStyle(style QTextListFormat__Style) {
	defer qt.Recovering("QTextListFormat::setStyle")

	if ptr.Pointer() != nil {
		C.QTextListFormat_SetStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTextListFormat) Style() QTextListFormat__Style {
	defer qt.Recovering("QTextListFormat::style")

	if ptr.Pointer() != nil {
		return QTextListFormat__Style(C.QTextListFormat_Style(ptr.Pointer()))
	}
	return 0
}
