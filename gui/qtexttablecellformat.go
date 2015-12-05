package gui

//#include "gui.h"
import "C"
import (
	"log"
	"unsafe"
)

type QTextTableCellFormat struct {
	QTextCharFormat
}

type QTextTableCellFormat_ITF interface {
	QTextCharFormat_ITF
	QTextTableCellFormat_PTR() *QTextTableCellFormat
}

func PointerFromQTextTableCellFormat(ptr QTextTableCellFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableCellFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextTableCellFormatFromPointer(ptr unsafe.Pointer) *QTextTableCellFormat {
	var n = new(QTextTableCellFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextTableCellFormat) QTextTableCellFormat_PTR() *QTextTableCellFormat {
	return ptr
}

func NewQTextTableCellFormat() *QTextTableCellFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::QTextTableCellFormat")
		}
	}()

	return NewQTextTableCellFormatFromPointer(C.QTextTableCellFormat_NewQTextTableCellFormat())
}

func (ptr *QTextTableCellFormat) BottomPadding() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::bottomPadding")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextTableCellFormat_BottomPadding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCellFormat) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextTableCellFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextTableCellFormat) LeftPadding() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::leftPadding")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextTableCellFormat_LeftPadding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCellFormat) RightPadding() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::rightPadding")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextTableCellFormat_RightPadding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableCellFormat) SetBottomPadding(padding float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::setBottomPadding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextTableCellFormat_SetBottomPadding(ptr.Pointer(), C.double(padding))
	}
}

func (ptr *QTextTableCellFormat) SetLeftPadding(padding float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::setLeftPadding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextTableCellFormat_SetLeftPadding(ptr.Pointer(), C.double(padding))
	}
}

func (ptr *QTextTableCellFormat) SetPadding(padding float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::setPadding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextTableCellFormat_SetPadding(ptr.Pointer(), C.double(padding))
	}
}

func (ptr *QTextTableCellFormat) SetRightPadding(padding float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::setRightPadding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextTableCellFormat_SetRightPadding(ptr.Pointer(), C.double(padding))
	}
}

func (ptr *QTextTableCellFormat) SetTopPadding(padding float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::setTopPadding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextTableCellFormat_SetTopPadding(ptr.Pointer(), C.double(padding))
	}
}

func (ptr *QTextTableCellFormat) TopPadding() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextTableCellFormat::topPadding")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextTableCellFormat_TopPadding(ptr.Pointer()))
	}
	return 0
}
