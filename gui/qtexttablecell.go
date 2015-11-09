package gui

//#include "qtexttablecell.h"
import "C"
import (
	"unsafe"
)

type QTextTableCell struct {
	ptr unsafe.Pointer
}

type QTextTableCell_ITF interface {
	QTextTableCell_PTR() *QTextTableCell
}

func (p *QTextTableCell) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextTableCell) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextTableCell(ptr QTextTableCell_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableCell_PTR().Pointer()
	}
	return nil
}

func NewQTextTableCellFromPointer(ptr unsafe.Pointer) *QTextTableCell {
	var n = new(QTextTableCell)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextTableCell) QTextTableCell_PTR() *QTextTableCell {
	return ptr
}

func NewQTextTableCell() *QTextTableCell {
	return NewQTextTableCellFromPointer(C.QTextTableCell_NewQTextTableCell())
}

func NewQTextTableCell2(other QTextTableCell_ITF) *QTextTableCell {
	return NewQTextTableCellFromPointer(C.QTextTableCell_NewQTextTableCell2(PointerFromQTextTableCell(other)))
}

func (ptr *QTextTableCell) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCell) ColumnSpan() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_ColumnSpan(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCell) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextTableCell_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextTableCell) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCell) RowSpan() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_RowSpan(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCell) SetFormat(format QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextTableCell_SetFormat(ptr.Pointer(), PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextTableCell) TableCellFormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTableCell_TableCellFormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCell) DestroyQTextTableCell() {
	if ptr.Pointer() != nil {
		C.QTextTableCell_DestroyQTextTableCell(ptr.Pointer())
	}
}
