package widgets

//#include "qboxlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBoxLayout struct {
	QLayout
}

type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func PointerFromQBoxLayout(ptr QBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQBoxLayoutFromPointer(ptr unsafe.Pointer) *QBoxLayout {
	var n = new(QBoxLayout)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout {
	return ptr
}

//QBoxLayout::Direction
type QBoxLayout__Direction int64

const (
	QBoxLayout__LeftToRight = QBoxLayout__Direction(0)
	QBoxLayout__RightToLeft = QBoxLayout__Direction(1)
	QBoxLayout__TopToBottom = QBoxLayout__Direction(2)
	QBoxLayout__BottomToTop = QBoxLayout__Direction(3)
	QBoxLayout__Down        = QBoxLayout__Direction(QBoxLayout__TopToBottom)
	QBoxLayout__Up          = QBoxLayout__Direction(QBoxLayout__BottomToTop)
)

func (ptr *QBoxLayout) Direction() QBoxLayout__Direction {
	if ptr.Pointer() != nil {
		return QBoxLayout__Direction(C.QBoxLayout_Direction(ptr.Pointer()))
	}
	return 0
}

func NewQBoxLayout(dir QBoxLayout__Direction, parent QWidget_ITF) *QBoxLayout {
	return NewQBoxLayoutFromPointer(C.QBoxLayout_NewQBoxLayout(C.int(dir), PointerFromQWidget(parent)))
}

func (ptr *QBoxLayout) AddItem(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddLayout(layout QLayout_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddLayout(ptr.Pointer(), PointerFromQLayout(layout), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddSpacerItem(spacerItem QSpacerItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacerItem(ptr.Pointer(), PointerFromQSpacerItem(spacerItem))
	}
}

func (ptr *QBoxLayout) AddSpacing(size int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacing(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBoxLayout) AddStretch(stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStretch(ptr.Pointer(), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddStrut(size int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStrut(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBoxLayout) AddWidget(widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QBoxLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QBoxLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBoxLayout) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) InsertItem(index int, item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertItem(ptr.Pointer(), C.int(index), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) InsertLayout(index int, layout QLayout_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertLayout(ptr.Pointer(), C.int(index), PointerFromQLayout(layout), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertSpacerItem(index int, spacerItem QSpacerItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacerItem(ptr.Pointer(), C.int(index), PointerFromQSpacerItem(spacerItem))
	}
}

func (ptr *QBoxLayout) InsertSpacing(index int, size int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacing(ptr.Pointer(), C.int(index), C.int(size))
	}
}

func (ptr *QBoxLayout) InsertStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertWidget(index int, widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QBoxLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QBoxLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QBoxLayout) MinimumHeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) SetDirection(direction QBoxLayout__Direction) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QBoxLayout) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QBoxLayout) SetSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QBoxLayout) SetStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) SetStretchFactor2(layout QLayout_ITF, stretch int) bool {
	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor2(ptr.Pointer(), PointerFromQLayout(layout), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) SetStretchFactor(widget QWidget_ITF, stretch int) bool {
	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) Spacing() int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) Stretch(index int) int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Stretch(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QBoxLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QBoxLayout) DestroyQBoxLayout() {
	if ptr.Pointer() != nil {
		C.QBoxLayout_DestroyQBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
