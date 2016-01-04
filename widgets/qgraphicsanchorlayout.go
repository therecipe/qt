package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsAnchorLayout struct {
	QGraphicsLayout
}

type QGraphicsAnchorLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsAnchorLayout_PTR() *QGraphicsAnchorLayout
}

func PointerFromQGraphicsAnchorLayout(ptr QGraphicsAnchorLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsAnchorLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsAnchorLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsAnchorLayout {
	var n = new(QGraphicsAnchorLayout)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGraphicsAnchorLayout_") {
		n.SetObjectNameAbs("QGraphicsAnchorLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsAnchorLayout) QGraphicsAnchorLayout_PTR() *QGraphicsAnchorLayout {
	return ptr
}

func NewQGraphicsAnchorLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsAnchorLayout {
	defer qt.Recovering("QGraphicsAnchorLayout::QGraphicsAnchorLayout")

	return NewQGraphicsAnchorLayoutFromPointer(C.QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsAnchorLayout) AddAnchor(firstItem QGraphicsLayoutItem_ITF, firstEdge core.Qt__AnchorPoint, secondItem QGraphicsLayoutItem_ITF, secondEdge core.Qt__AnchorPoint) *QGraphicsAnchor {
	defer qt.Recovering("QGraphicsAnchorLayout::addAnchor")

	if ptr.Pointer() != nil {
		return NewQGraphicsAnchorFromPointer(C.QGraphicsAnchorLayout_AddAnchor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), C.int(firstEdge), PointerFromQGraphicsLayoutItem(secondItem), C.int(secondEdge)))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) AddAnchors(firstItem QGraphicsLayoutItem_ITF, secondItem QGraphicsLayoutItem_ITF, orientations core.Qt__Orientation) {
	defer qt.Recovering("QGraphicsAnchorLayout::addAnchors")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_AddAnchors(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), PointerFromQGraphicsLayoutItem(secondItem), C.int(orientations))
	}
}

func (ptr *QGraphicsAnchorLayout) AddCornerAnchors(firstItem QGraphicsLayoutItem_ITF, firstCorner core.Qt__Corner, secondItem QGraphicsLayoutItem_ITF, secondCorner core.Qt__Corner) {
	defer qt.Recovering("QGraphicsAnchorLayout::addCornerAnchors")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_AddCornerAnchors(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), C.int(firstCorner), PointerFromQGraphicsLayoutItem(secondItem), C.int(secondCorner))
	}
}

func (ptr *QGraphicsAnchorLayout) Anchor(firstItem QGraphicsLayoutItem_ITF, firstEdge core.Qt__AnchorPoint, secondItem QGraphicsLayoutItem_ITF, secondEdge core.Qt__AnchorPoint) *QGraphicsAnchor {
	defer qt.Recovering("QGraphicsAnchorLayout::anchor")

	if ptr.Pointer() != nil {
		return NewQGraphicsAnchorFromPointer(C.QGraphicsAnchorLayout_Anchor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), C.int(firstEdge), PointerFromQGraphicsLayoutItem(secondItem), C.int(secondEdge)))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) Count() int {
	defer qt.Recovering("QGraphicsAnchorLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsAnchorLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	defer qt.Recovering("QGraphicsAnchorLayout::horizontalSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsAnchorLayout_HorizontalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QGraphicsAnchorLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "invalidate", f)
	}
}

func (ptr *QGraphicsAnchorLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QGraphicsAnchorLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "invalidate")
	}
}

//export callbackQGraphicsAnchorLayoutInvalidate
func callbackQGraphicsAnchorLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsAnchorLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsAnchorLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QGraphicsAnchorLayout) Invalidate() {
	defer qt.Recovering("QGraphicsAnchorLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsAnchorLayout) InvalidateDefault() {
	defer qt.Recovering("QGraphicsAnchorLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem {
	defer qt.Recovering("QGraphicsAnchorLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsAnchorLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) ConnectRemoveAt(f func(index int)) {
	defer qt.Recovering("connect QGraphicsAnchorLayout::removeAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "removeAt", f)
	}
}

func (ptr *QGraphicsAnchorLayout) DisconnectRemoveAt() {
	defer qt.Recovering("disconnect QGraphicsAnchorLayout::removeAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "removeAt")
	}
}

//export callbackQGraphicsAnchorLayoutRemoveAt
func callbackQGraphicsAnchorLayoutRemoveAt(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QGraphicsAnchorLayout::removeAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "removeAt"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QGraphicsAnchorLayout) RemoveAt(index int) {
	defer qt.Recovering("QGraphicsAnchorLayout::removeAt")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_RemoveAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsAnchorLayout) RemoveAtDefault(index int) {
	defer qt.Recovering("QGraphicsAnchorLayout::removeAt")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_RemoveAtDefault(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsAnchorLayout::setHorizontalSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetHorizontalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsAnchorLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsAnchorLayout::setVerticalSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetVerticalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	defer qt.Recovering("QGraphicsAnchorLayout::verticalSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsAnchorLayout_VerticalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) DestroyQGraphicsAnchorLayout() {
	defer qt.Recovering("QGraphicsAnchorLayout::~QGraphicsAnchorLayout")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(ptr.Pointer())
	}
}

func (ptr *QGraphicsAnchorLayout) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsAnchorLayout::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsAnchorLayout_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsAnchorLayout) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsAnchorLayout::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
