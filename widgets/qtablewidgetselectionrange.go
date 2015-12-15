package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTableWidgetSelectionRange struct {
	ptr unsafe.Pointer
}

type QTableWidgetSelectionRange_ITF interface {
	QTableWidgetSelectionRange_PTR() *QTableWidgetSelectionRange
}

func (p *QTableWidgetSelectionRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTableWidgetSelectionRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTableWidgetSelectionRange(ptr QTableWidgetSelectionRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidgetSelectionRange_PTR().Pointer()
	}
	return nil
}

func NewQTableWidgetSelectionRangeFromPointer(ptr unsafe.Pointer) *QTableWidgetSelectionRange {
	var n = new(QTableWidgetSelectionRange)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTableWidgetSelectionRange) QTableWidgetSelectionRange_PTR() *QTableWidgetSelectionRange {
	return ptr
}

func NewQTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	defer qt.Recovering("QTableWidgetSelectionRange::QTableWidgetSelectionRange")

	return NewQTableWidgetSelectionRangeFromPointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange())
}

func NewQTableWidgetSelectionRange3(other QTableWidgetSelectionRange_ITF) *QTableWidgetSelectionRange {
	defer qt.Recovering("QTableWidgetSelectionRange::QTableWidgetSelectionRange")

	return NewQTableWidgetSelectionRangeFromPointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(PointerFromQTableWidgetSelectionRange(other)))
}

func NewQTableWidgetSelectionRange2(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	defer qt.Recovering("QTableWidgetSelectionRange::QTableWidgetSelectionRange")

	return NewQTableWidgetSelectionRangeFromPointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(C.int(top), C.int(left), C.int(bottom), C.int(right)))
}

func (ptr *QTableWidgetSelectionRange) BottomRow() int {
	defer qt.Recovering("QTableWidgetSelectionRange::bottomRow")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_BottomRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) ColumnCount() int {
	defer qt.Recovering("QTableWidgetSelectionRange::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) LeftColumn() int {
	defer qt.Recovering("QTableWidgetSelectionRange::leftColumn")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_LeftColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) RightColumn() int {
	defer qt.Recovering("QTableWidgetSelectionRange::rightColumn")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_RightColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) RowCount() int {
	defer qt.Recovering("QTableWidgetSelectionRange::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) TopRow() int {
	defer qt.Recovering("QTableWidgetSelectionRange::topRow")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_TopRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) DestroyQTableWidgetSelectionRange() {
	defer qt.Recovering("QTableWidgetSelectionRange::~QTableWidgetSelectionRange")

	if ptr.Pointer() != nil {
		C.QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(ptr.Pointer())
	}
}
