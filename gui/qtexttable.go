package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextTable struct {
	QTextFrame
}

type QTextTable_ITF interface {
	QTextFrame_ITF
	QTextTable_PTR() *QTextTable
}

func PointerFromQTextTable(ptr QTextTable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTable_PTR().Pointer()
	}
	return nil
}

func NewQTextTableFromPointer(ptr unsafe.Pointer) *QTextTable {
	var n = new(QTextTable)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextTable_") {
		n.SetObjectName("QTextTable_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextTable) QTextTable_PTR() *QTextTable {
	return ptr
}

func (ptr *QTextTable) InsertColumns(index int, columns int) {
	defer qt.Recovering("QTextTable::insertColumns")

	if ptr.Pointer() != nil {
		C.QTextTable_InsertColumns(ptr.Pointer(), C.int(index), C.int(columns))
	}
}

func (ptr *QTextTable) InsertRows(index int, rows int) {
	defer qt.Recovering("QTextTable::insertRows")

	if ptr.Pointer() != nil {
		C.QTextTable_InsertRows(ptr.Pointer(), C.int(index), C.int(rows))
	}
}

func (ptr *QTextTable) RemoveColumns(index int, columns int) {
	defer qt.Recovering("QTextTable::removeColumns")

	if ptr.Pointer() != nil {
		C.QTextTable_RemoveColumns(ptr.Pointer(), C.int(index), C.int(columns))
	}
}

func (ptr *QTextTable) RemoveRows(index int, rows int) {
	defer qt.Recovering("QTextTable::removeRows")

	if ptr.Pointer() != nil {
		C.QTextTable_RemoveRows(ptr.Pointer(), C.int(index), C.int(rows))
	}
}

func (ptr *QTextTable) Resize(rows int, columns int) {
	defer qt.Recovering("QTextTable::resize")

	if ptr.Pointer() != nil {
		C.QTextTable_Resize(ptr.Pointer(), C.int(rows), C.int(columns))
	}
}

func (ptr *QTextTable) SetFormat(format QTextTableFormat_ITF) {
	defer qt.Recovering("QTextTable::setFormat")

	if ptr.Pointer() != nil {
		C.QTextTable_SetFormat(ptr.Pointer(), PointerFromQTextTableFormat(format))
	}
}

func (ptr *QTextTable) AppendColumns(count int) {
	defer qt.Recovering("QTextTable::appendColumns")

	if ptr.Pointer() != nil {
		C.QTextTable_AppendColumns(ptr.Pointer(), C.int(count))
	}
}

func (ptr *QTextTable) AppendRows(count int) {
	defer qt.Recovering("QTextTable::appendRows")

	if ptr.Pointer() != nil {
		C.QTextTable_AppendRows(ptr.Pointer(), C.int(count))
	}
}

func (ptr *QTextTable) Columns() int {
	defer qt.Recovering("QTextTable::columns")

	if ptr.Pointer() != nil {
		return int(C.QTextTable_Columns(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTable) MergeCells2(cursor QTextCursor_ITF) {
	defer qt.Recovering("QTextTable::mergeCells")

	if ptr.Pointer() != nil {
		C.QTextTable_MergeCells2(ptr.Pointer(), PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextTable) MergeCells(row int, column int, numRows int, numCols int) {
	defer qt.Recovering("QTextTable::mergeCells")

	if ptr.Pointer() != nil {
		C.QTextTable_MergeCells(ptr.Pointer(), C.int(row), C.int(column), C.int(numRows), C.int(numCols))
	}
}

func (ptr *QTextTable) Rows() int {
	defer qt.Recovering("QTextTable::rows")

	if ptr.Pointer() != nil {
		return int(C.QTextTable_Rows(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTable) SplitCell(row int, column int, numRows int, numCols int) {
	defer qt.Recovering("QTextTable::splitCell")

	if ptr.Pointer() != nil {
		C.QTextTable_SplitCell(ptr.Pointer(), C.int(row), C.int(column), C.int(numRows), C.int(numCols))
	}
}
