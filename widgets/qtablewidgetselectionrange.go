package widgets

//#include "qtablewidgetselectionrange.h"
import "C"
import (
	"unsafe"
)

type QTableWidgetSelectionRange struct {
	ptr unsafe.Pointer
}

type QTableWidgetSelectionRangeITF interface {
	QTableWidgetSelectionRangePTR() *QTableWidgetSelectionRange
}

func (p *QTableWidgetSelectionRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTableWidgetSelectionRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTableWidgetSelectionRange(ptr QTableWidgetSelectionRangeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidgetSelectionRangePTR().Pointer()
	}
	return nil
}

func QTableWidgetSelectionRangeFromPointer(ptr unsafe.Pointer) *QTableWidgetSelectionRange {
	var n = new(QTableWidgetSelectionRange)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTableWidgetSelectionRange) QTableWidgetSelectionRangePTR() *QTableWidgetSelectionRange {
	return ptr
}

func NewQTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	return QTableWidgetSelectionRangeFromPointer(unsafe.Pointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange()))
}

func NewQTableWidgetSelectionRange3(other QTableWidgetSelectionRangeITF) *QTableWidgetSelectionRange {
	return QTableWidgetSelectionRangeFromPointer(unsafe.Pointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(C.QtObjectPtr(PointerFromQTableWidgetSelectionRange(other)))))
}

func NewQTableWidgetSelectionRange2(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	return QTableWidgetSelectionRangeFromPointer(unsafe.Pointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(C.int(top), C.int(left), C.int(bottom), C.int(right))))
}

func (ptr *QTableWidgetSelectionRange) BottomRow() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_BottomRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) LeftColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_LeftColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) RightColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_RightColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) TopRow() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_TopRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) DestroyQTableWidgetSelectionRange() {
	if ptr.Pointer() != nil {
		C.QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(C.QtObjectPtr(ptr.Pointer()))
	}
}
