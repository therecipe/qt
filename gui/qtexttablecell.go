package gui

//#include "qtexttablecell.h"
import "C"
import (
	"unsafe"
)

type QTextTableCell struct {
	ptr unsafe.Pointer
}

type QTextTableCellITF interface {
	QTextTableCellPTR() *QTextTableCell
}

func (p *QTextTableCell) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextTableCell) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextTableCell(ptr QTextTableCellITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableCellPTR().Pointer()
	}
	return nil
}

func QTextTableCellFromPointer(ptr unsafe.Pointer) *QTextTableCell {
	var n = new(QTextTableCell)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextTableCell) QTextTableCellPTR() *QTextTableCell {
	return ptr
}

func NewQTextTableCell() *QTextTableCell {
	return QTextTableCellFromPointer(unsafe.Pointer(C.QTextTableCell_NewQTextTableCell()))
}

func NewQTextTableCell2(other QTextTableCellITF) *QTextTableCell {
	return QTextTableCellFromPointer(unsafe.Pointer(C.QTextTableCell_NewQTextTableCell2(C.QtObjectPtr(PointerFromQTextTableCell(other)))))
}

func (ptr *QTextTableCell) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_Column(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) ColumnSpan() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_ColumnSpan(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextTableCell_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextTableCell) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_Row(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) RowSpan() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_RowSpan(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) SetFormat(format QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextTableCell_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QTextTableCell) TableCellFormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_TableCellFormatIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) DestroyQTextTableCell() {
	if ptr.Pointer() != nil {
		C.QTextTableCell_DestroyQTextTableCell(C.QtObjectPtr(ptr.Pointer()))
	}
}
