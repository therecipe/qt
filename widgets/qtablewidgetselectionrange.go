package widgets

//#include "widgets.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::QTableWidgetSelectionRange")
		}
	}()

	return NewQTableWidgetSelectionRangeFromPointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange())
}

func NewQTableWidgetSelectionRange3(other QTableWidgetSelectionRange_ITF) *QTableWidgetSelectionRange {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::QTableWidgetSelectionRange")
		}
	}()

	return NewQTableWidgetSelectionRangeFromPointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(PointerFromQTableWidgetSelectionRange(other)))
}

func NewQTableWidgetSelectionRange2(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::QTableWidgetSelectionRange")
		}
	}()

	return NewQTableWidgetSelectionRangeFromPointer(C.QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(C.int(top), C.int(left), C.int(bottom), C.int(right)))
}

func (ptr *QTableWidgetSelectionRange) BottomRow() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::bottomRow")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_BottomRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) LeftColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::leftColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_LeftColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) RightColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::rightColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_RightColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) RowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) TopRow() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::topRow")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetSelectionRange_TopRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetSelectionRange) DestroyQTableWidgetSelectionRange() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidgetSelectionRange::~QTableWidgetSelectionRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(ptr.Pointer())
	}
}
