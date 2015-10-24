package widgets

//#include "qwidgetitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWidgetItem struct {
	QLayoutItem
}

type QWidgetItemITF interface {
	QLayoutItemITF
	QWidgetItemPTR() *QWidgetItem
}

func PointerFromQWidgetItem(ptr QWidgetItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetItemPTR().Pointer()
	}
	return nil
}

func QWidgetItemFromPointer(ptr unsafe.Pointer) *QWidgetItem {
	var n = new(QWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWidgetItem) QWidgetItemPTR() *QWidgetItem {
	return ptr
}

func NewQWidgetItem(widget QWidgetITF) *QWidgetItem {
	return QWidgetItemFromPointer(unsafe.Pointer(C.QWidgetItem_NewQWidgetItem(C.QtObjectPtr(PointerFromQWidget(widget)))))
}

func (ptr *QWidgetItem) ControlTypes() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QWidgetItem_ControlTypes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidgetItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QWidgetItem_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidgetItem) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QWidgetItem_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidgetItem) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QWidgetItem_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QWidgetItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QWidgetItem_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidgetItem) SetGeometry(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QWidgetItem_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QWidgetItem) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidgetItem_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidgetItem) DestroyQWidgetItem() {
	if ptr.Pointer() != nil {
		C.QWidgetItem_DestroyQWidgetItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
