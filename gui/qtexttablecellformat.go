package gui

//#include "qtexttablecellformat.h"
import "C"
import (
	"unsafe"
)

type QTextTableCellFormat struct {
	QTextCharFormat
}

type QTextTableCellFormatITF interface {
	QTextCharFormatITF
	QTextTableCellFormatPTR() *QTextTableCellFormat
}

func PointerFromQTextTableCellFormat(ptr QTextTableCellFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableCellFormatPTR().Pointer()
	}
	return nil
}

func QTextTableCellFormatFromPointer(ptr unsafe.Pointer) *QTextTableCellFormat {
	var n = new(QTextTableCellFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextTableCellFormat) QTextTableCellFormatPTR() *QTextTableCellFormat {
	return ptr
}

func NewQTextTableCellFormat() *QTextTableCellFormat {
	return QTextTableCellFormatFromPointer(unsafe.Pointer(C.QTextTableCellFormat_NewQTextTableCellFormat()))
}

func (ptr *QTextTableCellFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextTableCellFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
