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

type QGraphicsLinearLayoutITF interface {
	QGraphicsLayoutITF
	QGraphicsLinearLayoutPTR() *QGraphicsLinearLayout
}

func PointerFromQGraphicsLinearLayout(ptr QGraphicsLinearLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLinearLayoutPTR().Pointer()
	}
	return nil
}

func QGraphicsLinearLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsLinearLayout {
	var n = new(QGraphicsLinearLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLinearLayout) QGraphicsLinearLayoutPTR() *QGraphicsLinearLayout {
	return ptr
}

func NewQGraphicsLinearLayout(parent QGraphicsLayoutItemITF) *QGraphicsLinearLayout {
	return QGraphicsLinearLayoutFromPointer(unsafe.Pointer(C.QGraphicsLinearLayout_NewQGraphicsLinearLayout(C.QtObjectPtr(PointerFromQGraphicsLayoutItem(parent)))))
}

func NewQGraphicsLinearLayout2(orientation core.Qt__Orientation, parent QGraphicsLayoutItemITF) *QGraphicsLinearLayout {
	return QGraphicsLinearLayoutFromPointer(unsafe.Pointer(C.QGraphicsLinearLayout_NewQGraphicsLinearLayout2(C.int(orientation), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(parent)))))
}

func (ptr *QGraphicsLinearLayout) AddItem(item QGraphicsLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)))
	}
}

func (ptr *QGraphicsLinearLayout) AddStretch(stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_AddStretch(C.QtObjectPtr(ptr.Pointer()), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Alignment(item QGraphicsLayoutItemITF) core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsLinearLayout_Alignment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item))))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsLinearLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) InsertItem(index int, item QGraphicsLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InsertItem(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)))
	}
}

func (ptr *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InsertStretch(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return QGraphicsLayoutItemFromPointer(unsafe.Pointer(C.QGraphicsLinearLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QGraphicsLinearLayout) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGraphicsLinearLayout_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) RemoveAt(index int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QGraphicsLinearLayout) RemoveItem(item QGraphicsLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)))
	}
}

func (ptr *QGraphicsLinearLayout) SetAlignment(item QGraphicsLayoutItemITF, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)), C.int(alignment))
	}
}

func (ptr *QGraphicsLinearLayout) SetGeometry(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsLinearLayout) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}

func (ptr *QGraphicsLinearLayout) SetStretchFactor(item QGraphicsLayoutItemITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) StretchFactor(item QGraphicsLayoutItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsLinearLayout_StretchFactor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item))))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) DestroyQGraphicsLinearLayout() {
	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}
