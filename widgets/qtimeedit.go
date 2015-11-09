package widgets

//#include "qtimeedit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTimeEdit struct {
	QDateTimeEdit
}

type QTimeEdit_ITF interface {
	QDateTimeEdit_ITF
	QTimeEdit_PTR() *QTimeEdit
}

func PointerFromQTimeEdit(ptr QTimeEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeEdit_PTR().Pointer()
	}
	return nil
}

func NewQTimeEditFromPointer(ptr unsafe.Pointer) *QTimeEdit {
	var n = new(QTimeEdit)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTimeEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTimeEdit) QTimeEdit_PTR() *QTimeEdit {
	return ptr
}

func NewQTimeEdit(parent QWidget_ITF) *QTimeEdit {
	return NewQTimeEditFromPointer(C.QTimeEdit_NewQTimeEdit(PointerFromQWidget(parent)))
}

func NewQTimeEdit2(time core.QTime_ITF, parent QWidget_ITF) *QTimeEdit {
	return NewQTimeEditFromPointer(C.QTimeEdit_NewQTimeEdit2(core.PointerFromQTime(time), PointerFromQWidget(parent)))
}

func (ptr *QTimeEdit) DestroyQTimeEdit() {
	if ptr.Pointer() != nil {
		C.QTimeEdit_DestroyQTimeEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
