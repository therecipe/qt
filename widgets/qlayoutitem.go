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

type QLayoutItemITF interface {
	QLayoutItemPTR() *QLayoutItem
}

func (p *QLayoutItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLayoutItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLayoutItem(ptr QLayoutItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutItemPTR().Pointer()
	}
	return nil
}

func QLayoutItemFromPointer(ptr unsafe.Pointer) *QLayoutItem {
	var n = new(QLayoutItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLayoutItem) QLayoutItemPTR() *QLayoutItem {
	return ptr
}

func (ptr *QLayoutItem) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLayoutItem_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayoutItem) ControlTypes() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayoutItem_ControlTypes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayoutItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayoutItem_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayoutItem) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QLayoutItem_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayoutItem) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) Invalidate() {
	if ptr.Pointer() != nil {
		C.QLayoutItem_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLayoutItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QLayoutItem_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayoutItem) Layout() *QLayout {
	if ptr.Pointer() != nil {
		return QLayoutFromPointer(unsafe.Pointer(C.QLayoutItem_Layout(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLayoutItem) MinimumHeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_MinimumHeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QLayoutItem) SetGeometry(r core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(r)))
	}
}

func (ptr *QLayoutItem) SpacerItem() *QSpacerItem {
	if ptr.Pointer() != nil {
		return QSpacerItemFromPointer(unsafe.Pointer(C.QLayoutItem_SpacerItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLayoutItem) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QLayoutItem_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLayoutItem) DestroyQLayoutItem() {
	if ptr.Pointer() != nil {
		C.QLayoutItem_DestroyQLayoutItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
