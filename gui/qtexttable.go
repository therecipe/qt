package gui

//#include "qtexttable.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextTable struct {
	QTextFrame
}

type QTextTableITF interface {
	QTextFrameITF
	QTextTablePTR() *QTextTable
}

func PointerFromQTextTable(ptr QTextTableITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTablePTR().Pointer()
	}
	return nil
}

func QTextTableFromPointer(ptr unsafe.Pointer) *QTextTable {
	var n = new(QTextTable)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextTable_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextTable) QTextTablePTR() *QTextTable {
	return ptr
}

func (ptr *QTextTable) InsertColumns(index int, columns int) {
	if ptr.Pointer() != nil {
		C.QTextTable_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(columns))
	}
}

func (ptr *QTextTable) InsertRows(index int, rows int) {
	if ptr.Pointer() != nil {
		C.QTextTable_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(rows))
	}
}

func (ptr *QTextTable) RemoveColumns(index int, columns int) {
	if ptr.Pointer() != nil {
		C.QTextTable_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(columns))
	}
}

func (ptr *QTextTable) RemoveRows(index int, rows int) {
	if ptr.Pointer() != nil {
		C.QTextTable_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(rows))
	}
}

func (ptr *QTextTable) Resize(rows int, columns int) {
	if ptr.Pointer() != nil {
		C.QTextTable_Resize(C.QtObjectPtr(ptr.Pointer()), C.int(rows), C.int(columns))
	}
}

func (ptr *QTextTable) SetFormat(format QTextTableFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextTable_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextTableFormat(format)))
	}
}

func (ptr *QTextTable) AppendColumns(count int) {
	if ptr.Pointer() != nil {
		C.QTextTable_AppendColumns(C.QtObjectPtr(ptr.Pointer()), C.int(count))
	}
}

func (ptr *QTextTable) AppendRows(count int) {
	if ptr.Pointer() != nil {
		C.QTextTable_AppendRows(C.QtObjectPtr(ptr.Pointer()), C.int(count))
	}
}

func (ptr *QTextTable) Columns() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTable_Columns(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTable) MergeCells2(cursor QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QTextTable_MergeCells2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCursor(cursor)))
	}
}

func (ptr *QTextTable) MergeCells(row int, column int, numRows int, numCols int) {
	if ptr.Pointer() != nil {
		C.QTextTable_MergeCells(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.int(numRows), C.int(numCols))
	}
}

func (ptr *QTextTable) Rows() int {
	if ptr.Pointer() != nil {
		return int(C.QTextTable_Rows(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTable) SplitCell(row int, column int, numRows int, numCols int) {
	if ptr.Pointer() != nil {
		C.QTextTable_SplitCell(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.int(numRows), C.int(numCols))
	}
}
