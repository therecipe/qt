package widgets

//#include "qgraphicslinearlayout.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsLinearLayout struct {
	QGraphicsLayout
}

type QGraphicsLinearLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsLinearLayout_PTR() *QGraphicsLinearLayout
}

func PointerFromQGraphicsLinearLayout(ptr QGraphicsLinearLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLinearLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLinearLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsLinearLayout {
	var n = new(QGraphicsLinearLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLinearLayout) QGraphicsLinearLayout_PTR() *QGraphicsLinearLayout {
	return ptr
}

func NewQGraphicsLinearLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsLinearLayout {
	return NewQGraphicsLinearLayoutFromPointer(C.QGraphicsLinearLayout_NewQGraphicsLinearLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func NewQGraphicsLinearLayout2(orientation core.Qt__Orientation, parent QGraphicsLayoutItem_ITF) *QGraphicsLinearLayout {
	return NewQGraphicsLinearLayoutFromPointer(C.QGraphicsLinearLayout_NewQGraphicsLinearLayout2(C.int(orientation), PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsLinearLayout) AddItem(item QGraphicsLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_AddItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsLinearLayout) AddStretch(stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_AddStretch(ptr.Pointer(), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Alignment(item QGraphicsLayoutItem_ITF) core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsLinearLayout_Alignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item)))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsLinearLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) InsertItem(index int, item QGraphicsLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InsertItem(ptr.Pointer(), C.int(index), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InsertStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsLinearLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLinearLayout_ItemSpacing(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGraphicsLinearLayout_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) RemoveAt(index int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsLinearLayout) RemoveItem(item QGraphicsLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsLinearLayout) SetAlignment(item QGraphicsLayoutItem_ITF, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetAlignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(alignment))
	}
}

func (ptr *QGraphicsLinearLayout) SetGeometry(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetItemSpacing(ptr.Pointer(), C.int(index), C.double(spacing))
	}
}

func (ptr *QGraphicsLinearLayout) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsLinearLayout) SetStretchFactor(item QGraphicsLayoutItem_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetStretchFactor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Spacing() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLinearLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) StretchFactor(item QGraphicsLayoutItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsLinearLayout_StretchFactor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item)))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) DestroyQGraphicsLinearLayout() {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(ptr.Pointer())
	}
}
