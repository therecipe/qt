package widgets

//#include "qdateedit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDateEdit struct {
	QDateTimeEdit
}

type QDateEditITF interface {
	QDateTimeEditITF
	QDateEditPTR() *QDateEdit
}

func PointerFromQDateEdit(ptr QDateEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateEditPTR().Pointer()
	}
	return nil
}

func QDateEditFromPointer(ptr unsafe.Pointer) *QDateEdit {
	var n = new(QDateEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDateEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDateEdit) QDateEditPTR() *QDateEdit {
	return ptr
}

func NewQDateEdit(parent QWidgetITF) *QDateEdit {
	return QDateEditFromPointer(unsafe.Pointer(C.QDateEdit_NewQDateEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQDateEdit2(date core.QDateITF, parent QWidgetITF) *QDateEdit {
	return QDateEditFromPointer(unsafe.Pointer(C.QDateEdit_NewQDateEdit2(C.QtObjectPtr(core.PointerFromQDate(date)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QDateEdit) DestroyQDateEdit() {
	if ptr.Pointer() != nil {
		C.QDateEdit_DestroyQDateEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
