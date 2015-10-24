package gui

//#include "qtexttableformat.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextTableFormat struct {
	QTextFrameFormat
}

type QTextTableFormatITF interface {
	QTextFrameFormatITF
	QTextTableFormatPTR() *QTextTableFormat
}

func PointerFromQTextTableFormat(ptr QTextTableFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableFormatPTR().Pointer()
	}
	return nil
}

func QTextTableFormatFromPointer(ptr unsafe.Pointer) *QTextTableFormat {
	var n = new(QTextTableFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextTableFormat) QTextTableFormatPTR() *QTextTableFormat {
	return ptr
}

func NewQTextTableFormat() *QTextTableFormat {
	return QTextTableFormatFromPointer(unsafe.Pointer(C.QTextTableFormat_NewQTextTableFormat()))
}

func (ptr *QTextTableFormat) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextTableFormat_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableFormat) ClearColumnWidthConstraints() {
	if ptr.Pointer() != nil {
		C.QTextTableFormat_ClearColumnWidthConstraints(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextTableFormat) Columns() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableFormat_Columns(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableFormat) HeaderRowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableFormat_HeaderRowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextTableFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextTableFormat) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QTextTableFormat) SetHeaderRowCount(count int) {
	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetHeaderRowCount(C.QtObjectPtr(ptr.Pointer()), C.int(count))
	}
}
