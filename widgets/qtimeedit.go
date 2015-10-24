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

type QTimeEditITF interface {
	QDateTimeEditITF
	QTimeEditPTR() *QTimeEdit
}

func PointerFromQTimeEdit(ptr QTimeEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeEditPTR().Pointer()
	}
	return nil
}

func QTimeEditFromPointer(ptr unsafe.Pointer) *QTimeEdit {
	var n = new(QTimeEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTimeEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTimeEdit) QTimeEditPTR() *QTimeEdit {
	return ptr
}

func NewQTimeEdit(parent QWidgetITF) *QTimeEdit {
	return QTimeEditFromPointer(unsafe.Pointer(C.QTimeEdit_NewQTimeEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQTimeEdit2(time core.QTimeITF, parent QWidgetITF) *QTimeEdit {
	return QTimeEditFromPointer(unsafe.Pointer(C.QTimeEdit_NewQTimeEdit2(C.QtObjectPtr(core.PointerFromQTime(time)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTimeEdit) DestroyQTimeEdit() {
	if ptr.Pointer() != nil {
		C.QTimeEdit_DestroyQTimeEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
