package widgets

//#include "qlayoutitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLayoutItem struct {
	ptr unsafe.Pointer
}

type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (p *QLayoutItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLayoutItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLayoutItem(ptr QLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQLayoutItemFromPointer(ptr unsafe.Pointer) *QLayoutItem {
	var n = new(QLayoutItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem {
	return ptr
}

func (ptr *QLayoutItem) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLayoutItem_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ControlTypes() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayoutItem_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayoutItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QLayoutItem_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayoutItem) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) Invalidate() {
	if ptr.Pointer() != nil {
		C.QLayoutItem_Invalidate(ptr.Pointer())
	}
}

func (ptr *QLayoutItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QLayoutItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayoutItem) Layout() *QLayout {
	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QLayoutItem_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) MinimumHeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QLayoutItem) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayoutItem) SpacerItem() *QSpacerItem {
	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QLayoutItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayoutItem_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) DestroyQLayoutItem() {
	if ptr.Pointer() != nil {
		C.QLayoutItem_DestroyQLayoutItem(ptr.Pointer())
	}
}
