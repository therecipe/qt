package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QGraphicsLinearLayout_") {
		n.SetObjectNameAbs("QGraphicsLinearLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsLinearLayout) QGraphicsLinearLayout_PTR() *QGraphicsLinearLayout {
	return ptr
}

func NewQGraphicsLinearLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsLinearLayout {
	defer qt.Recovering("QGraphicsLinearLayout::QGraphicsLinearLayout")

	return NewQGraphicsLinearLayoutFromPointer(C.QGraphicsLinearLayout_NewQGraphicsLinearLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func NewQGraphicsLinearLayout2(orientation core.Qt__Orientation, parent QGraphicsLayoutItem_ITF) *QGraphicsLinearLayout {
	defer qt.Recovering("QGraphicsLinearLayout::QGraphicsLinearLayout")

	return NewQGraphicsLinearLayoutFromPointer(C.QGraphicsLinearLayout_NewQGraphicsLinearLayout2(C.int(orientation), PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsLinearLayout) AddItem(item QGraphicsLayoutItem_ITF) {
	defer qt.Recovering("QGraphicsLinearLayout::addItem")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_AddItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsLinearLayout) AddStretch(stretch int) {
	defer qt.Recovering("QGraphicsLinearLayout::addStretch")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_AddStretch(ptr.Pointer(), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Alignment(item QGraphicsLayoutItem_ITF) core.Qt__AlignmentFlag {
	defer qt.Recovering("QGraphicsLinearLayout::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsLinearLayout_Alignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item)))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) Count() int {
	defer qt.Recovering("QGraphicsLinearLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsLinearLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) InsertItem(index int, item QGraphicsLayoutItem_ITF) {
	defer qt.Recovering("QGraphicsLinearLayout::insertItem")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InsertItem(ptr.Pointer(), C.int(index), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	defer qt.Recovering("QGraphicsLinearLayout::insertStretch")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InsertStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QGraphicsLinearLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "invalidate", f)
	}
}

func (ptr *QGraphicsLinearLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QGraphicsLinearLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "invalidate")
	}
}

//export callbackQGraphicsLinearLayoutInvalidate
func callbackQGraphicsLinearLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsLinearLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsLinearLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QGraphicsLinearLayout) Invalidate() {
	defer qt.Recovering("QGraphicsLinearLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsLinearLayout) InvalidateDefault() {
	defer qt.Recovering("QGraphicsLinearLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem {
	defer qt.Recovering("QGraphicsLinearLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsLinearLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	defer qt.Recovering("QGraphicsLinearLayout::itemSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLinearLayout_ItemSpacing(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QGraphicsLinearLayout::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGraphicsLinearLayout_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) ConnectRemoveAt(f func(index int)) {
	defer qt.Recovering("connect QGraphicsLinearLayout::removeAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "removeAt", f)
	}
}

func (ptr *QGraphicsLinearLayout) DisconnectRemoveAt() {
	defer qt.Recovering("disconnect QGraphicsLinearLayout::removeAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "removeAt")
	}
}

//export callbackQGraphicsLinearLayoutRemoveAt
func callbackQGraphicsLinearLayoutRemoveAt(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QGraphicsLinearLayout::removeAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "removeAt"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QGraphicsLinearLayout) RemoveAt(index int) {
	defer qt.Recovering("QGraphicsLinearLayout::removeAt")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsLinearLayout) RemoveAtDefault(index int) {
	defer qt.Recovering("QGraphicsLinearLayout::removeAt")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveAtDefault(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsLinearLayout) RemoveItem(item QGraphicsLayoutItem_ITF) {
	defer qt.Recovering("QGraphicsLinearLayout::removeItem")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_RemoveItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsLinearLayout) SetAlignment(item QGraphicsLayoutItem_ITF, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsLinearLayout::setAlignment")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetAlignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(alignment))
	}
}

func (ptr *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	defer qt.Recovering("QGraphicsLinearLayout::setItemSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetItemSpacing(ptr.Pointer(), C.int(index), C.double(spacing))
	}
}

func (ptr *QGraphicsLinearLayout) SetOrientation(orientation core.Qt__Orientation) {
	defer qt.Recovering("QGraphicsLinearLayout::setOrientation")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsLinearLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsLinearLayout) SetStretchFactor(item QGraphicsLayoutItem_ITF, stretch int) {
	defer qt.Recovering("QGraphicsLinearLayout::setStretchFactor")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetStretchFactor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(stretch))
	}
}

func (ptr *QGraphicsLinearLayout) Spacing() float64 {
	defer qt.Recovering("QGraphicsLinearLayout::spacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLinearLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) StretchFactor(item QGraphicsLayoutItem_ITF) int {
	defer qt.Recovering("QGraphicsLinearLayout::stretchFactor")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsLinearLayout_StretchFactor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item)))
	}
	return 0
}

func (ptr *QGraphicsLinearLayout) DestroyQGraphicsLinearLayout() {
	defer qt.Recovering("QGraphicsLinearLayout::~QGraphicsLinearLayout")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(ptr.Pointer())
	}
}

func (ptr *QGraphicsLinearLayout) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsLinearLayout::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsLinearLayout_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsLinearLayout) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsLinearLayout::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsLinearLayout_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
