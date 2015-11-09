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

type QWidgetItem_ITF interface {
	QLayoutItem_ITF
	QWidgetItem_PTR() *QWidgetItem
}

func PointerFromQWidgetItem(ptr QWidgetItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetItem_PTR().Pointer()
	}
	return nil
}

func NewQWidgetItemFromPointer(ptr unsafe.Pointer) *QWidgetItem {
	var n = new(QWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWidgetItem) QWidgetItem_PTR() *QWidgetItem {
	return ptr
}

func NewQWidgetItem(widget QWidget_ITF) *QWidgetItem {
	return NewQWidgetItemFromPointer(C.QWidgetItem_NewQWidgetItem(PointerFromQWidget(widget)))
}

func (ptr *QWidgetItem) ControlTypes() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QWidgetItem_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidgetItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QWidgetItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidgetItem) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QWidgetItem_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidgetItem) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QWidgetItem_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWidgetItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QWidgetItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidgetItem) SetGeometry(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidgetItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWidgetItem) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetItem_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidgetItem) DestroyQWidgetItem() {
	if ptr.Pointer() != nil {
		C.QWidgetItem_DestroyQWidgetItem(ptr.Pointer())
	}
}
