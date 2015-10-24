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

type QBoxLayoutITF interface {
	QLayoutITF
	QBoxLayoutPTR() *QBoxLayout
}

func PointerFromQBoxLayout(ptr QBoxLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayoutPTR().Pointer()
	}
	return nil
}

func QBoxLayoutFromPointer(ptr unsafe.Pointer) *QBoxLayout {
	var n = new(QBoxLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBoxLayout) QBoxLayoutPTR() *QBoxLayout {
	return ptr
}

//QBoxLayout::Direction
type QBoxLayout__Direction int

var (
	QBoxLayout__LeftToRight = QBoxLayout__Direction(0)
	QBoxLayout__RightToLeft = QBoxLayout__Direction(1)
	QBoxLayout__TopToBottom = QBoxLayout__Direction(2)
	QBoxLayout__BottomToTop = QBoxLayout__Direction(3)
	QBoxLayout__Down        = QBoxLayout__Direction(QBoxLayout__TopToBottom)
	QBoxLayout__Up          = QBoxLayout__Direction(QBoxLayout__BottomToTop)
)

func (ptr *QBoxLayout) Direction() QBoxLayout__Direction {
	if ptr.Pointer() != nil {
		return QBoxLayout__Direction(C.QBoxLayout_Direction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQBoxLayout(dir QBoxLayout__Direction, parent QWidgetITF) *QBoxLayout {
	return QBoxLayoutFromPointer(unsafe.Pointer(C.QBoxLayout_NewQBoxLayout(C.int(dir), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QBoxLayout) AddItem(item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QBoxLayout) AddLayout(layout QLayoutITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(layout)), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddSpacerItem(spacerItem QSpacerItemITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacerItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSpacerItem(spacerItem)))
	}
}

func (ptr *QBoxLayout) AddSpacing(size int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QBoxLayout) AddStretch(stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStretch(C.QtObjectPtr(ptr.Pointer()), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddStrut(size int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStrut(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QBoxLayout) AddWidget(widget QWidgetITF, stretch int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QBoxLayout_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxLayout) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QBoxLayout_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBoxLayout) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) InsertItem(index int, item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertItem(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QBoxLayout) InsertLayout(index int, layout QLayoutITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertLayout(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQLayout(layout)), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertSpacerItem(index int, spacerItem QSpacerItemITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacerItem(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQSpacerItem(spacerItem)))
	}
}

func (ptr *QBoxLayout) InsertSpacing(index int, size int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(size))
	}
}

func (ptr *QBoxLayout) InsertStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertStretch(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertWidget(index int, widget QWidgetITF, stretch int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertWidget(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QBoxLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBoxLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QBoxLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QBoxLayout) MinimumHeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_MinimumHeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) SetDirection(direction QBoxLayout__Direction) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QBoxLayout) SetGeometry(r core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(r)))
	}
}

func (ptr *QBoxLayout) SetSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QBoxLayout) SetStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetStretch(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) SetStretchFactor2(layout QLayoutITF, stretch int) bool {
	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(layout)), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) SetStretchFactor(widget QWidgetITF, stretch int) bool {
	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) Spacing() int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Spacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxLayout) Stretch(index int) int {
	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Stretch(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return 0
}

func (ptr *QBoxLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QBoxLayout_TakeAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QBoxLayout) DestroyQBoxLayout() {
	if ptr.Pointer() != nil {
		C.QBoxLayout_DestroyQBoxLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
