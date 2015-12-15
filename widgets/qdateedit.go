package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDateEdit struct {
	QDateTimeEdit
}

type QDateEdit_ITF interface {
	QDateTimeEdit_ITF
	QDateEdit_PTR() *QDateEdit
}

func PointerFromQDateEdit(ptr QDateEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateEdit_PTR().Pointer()
	}
	return nil
}

func NewQDateEditFromPointer(ptr unsafe.Pointer) *QDateEdit {
	var n = new(QDateEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDateEdit_") {
		n.SetObjectName("QDateEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QDateEdit) QDateEdit_PTR() *QDateEdit {
	return ptr
}

func NewQDateEdit(parent QWidget_ITF) *QDateEdit {
	defer qt.Recovering("QDateEdit::QDateEdit")

	return NewQDateEditFromPointer(C.QDateEdit_NewQDateEdit(PointerFromQWidget(parent)))
}

func NewQDateEdit2(date core.QDate_ITF, parent QWidget_ITF) *QDateEdit {
	defer qt.Recovering("QDateEdit::QDateEdit")

	return NewQDateEditFromPointer(C.QDateEdit_NewQDateEdit2(core.PointerFromQDate(date), PointerFromQWidget(parent)))
}

func (ptr *QDateEdit) DestroyQDateEdit() {
	defer qt.Recovering("QDateEdit::~QDateEdit")

	if ptr.Pointer() != nil {
		C.QDateEdit_DestroyQDateEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
