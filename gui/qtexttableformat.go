package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextTableFormat struct {
	QTextFrameFormat
}

type QTextTableFormat_ITF interface {
	QTextFrameFormat_ITF
	QTextTableFormat_PTR() *QTextTableFormat
}

func PointerFromQTextTableFormat(ptr QTextTableFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextTableFormatFromPointer(ptr unsafe.Pointer) *QTextTableFormat {
	var n = new(QTextTableFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextTableFormat) QTextTableFormat_PTR() *QTextTableFormat {
	return ptr
}

func NewQTextTableFormat() *QTextTableFormat {
	defer qt.Recovering("QTextTableFormat::QTextTableFormat")

	return NewQTextTableFormatFromPointer(C.QTextTableFormat_NewQTextTableFormat())
}

func (ptr *QTextTableFormat) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QTextTableFormat::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextTableFormat_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableFormat) CellPadding() float64 {
	defer qt.Recovering("QTextTableFormat::cellPadding")

	if ptr.Pointer() != nil {
		return float64(C.QTextTableFormat_CellPadding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableFormat) CellSpacing() float64 {
	defer qt.Recovering("QTextTableFormat::cellSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QTextTableFormat_CellSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableFormat) ClearColumnWidthConstraints() {
	defer qt.Recovering("QTextTableFormat::clearColumnWidthConstraints")

	if ptr.Pointer() != nil {
		C.QTextTableFormat_ClearColumnWidthConstraints(ptr.Pointer())
	}
}

func (ptr *QTextTableFormat) Columns() int {
	defer qt.Recovering("QTextTableFormat::columns")

	if ptr.Pointer() != nil {
		return int(C.QTextTableFormat_Columns(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableFormat) HeaderRowCount() int {
	defer qt.Recovering("QTextTableFormat::headerRowCount")

	if ptr.Pointer() != nil {
		return int(C.QTextTableFormat_HeaderRowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableFormat) IsValid() bool {
	defer qt.Recovering("QTextTableFormat::isValid")

	if ptr.Pointer() != nil {
		return C.QTextTableFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextTableFormat) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QTextTableFormat::setAlignment")

	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QTextTableFormat) SetCellPadding(padding float64) {
	defer qt.Recovering("QTextTableFormat::setCellPadding")

	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetCellPadding(ptr.Pointer(), C.double(padding))
	}
}

func (ptr *QTextTableFormat) SetCellSpacing(spacing float64) {
	defer qt.Recovering("QTextTableFormat::setCellSpacing")

	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetCellSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QTextTableFormat) SetHeaderRowCount(count int) {
	defer qt.Recovering("QTextTableFormat::setHeaderRowCount")

	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetHeaderRowCount(ptr.Pointer(), C.int(count))
	}
}
