package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGridLayout struct {
	QLayout
}

type QGridLayout_ITF interface {
	QLayout_ITF
	QGridLayout_PTR() *QGridLayout
}

func PointerFromQGridLayout(ptr QGridLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGridLayout_PTR().Pointer()
	}
	return nil
}

func NewQGridLayoutFromPointer(ptr unsafe.Pointer) *QGridLayout {
	var n = new(QGridLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGridLayout_") {
		n.SetObjectName("QGridLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QGridLayout) QGridLayout_PTR() *QGridLayout {
	return ptr
}

func (ptr *QGridLayout) HorizontalSpacing() int {
	defer qt.Recovering("QGridLayout::horizontalSpacing")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_HorizontalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) SetHorizontalSpacing(spacing int) {
	defer qt.Recovering("QGridLayout::setHorizontalSpacing")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetHorizontalSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QGridLayout) SetVerticalSpacing(spacing int) {
	defer qt.Recovering("QGridLayout::setVerticalSpacing")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetVerticalSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QGridLayout) VerticalSpacing() int {
	defer qt.Recovering("QGridLayout::verticalSpacing")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_VerticalSpacing(ptr.Pointer()))
	}
	return 0
}

func NewQGridLayout2() *QGridLayout {
	defer qt.Recovering("QGridLayout::QGridLayout")

	return NewQGridLayoutFromPointer(C.QGridLayout_NewQGridLayout2())
}

func NewQGridLayout(parent QWidget_ITF) *QGridLayout {
	defer qt.Recovering("QGridLayout::QGridLayout")

	return NewQGridLayoutFromPointer(C.QGridLayout_NewQGridLayout(PointerFromQWidget(parent)))
}

func (ptr *QGridLayout) AddItem(item QLayoutItem_ITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGridLayout::addItem")

	if ptr.Pointer() != nil {
		C.QGridLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddLayout(layout QLayout_ITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGridLayout::addLayout")

	if ptr.Pointer() != nil {
		C.QGridLayout_AddLayout(ptr.Pointer(), PointerFromQLayout(layout), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddLayout2(layout QLayout_ITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGridLayout::addLayout")

	if ptr.Pointer() != nil {
		C.QGridLayout_AddLayout2(ptr.Pointer(), PointerFromQLayout(layout), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddWidget2(widget QWidget_ITF, fromRow int, fromColumn int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGridLayout::addWidget")

	if ptr.Pointer() != nil {
		C.QGridLayout_AddWidget2(ptr.Pointer(), PointerFromQWidget(widget), C.int(fromRow), C.int(fromColumn), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddWidget(widget QWidget_ITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGridLayout::addWidget")

	if ptr.Pointer() != nil {
		C.QGridLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGridLayout) ColumnCount() int {
	defer qt.Recovering("QGridLayout::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) ColumnMinimumWidth(column int) int {
	defer qt.Recovering("QGridLayout::columnMinimumWidth")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_ColumnMinimumWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGridLayout) ColumnStretch(column int) int {
	defer qt.Recovering("QGridLayout::columnStretch")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_ColumnStretch(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGridLayout) Count() int {
	defer qt.Recovering("QGridLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QGridLayout::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGridLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) GetItemPosition(index int, row int, column int, rowSpan int, columnSpan int) {
	defer qt.Recovering("QGridLayout::getItemPosition")

	if ptr.Pointer() != nil {
		C.QGridLayout_GetItemPosition(ptr.Pointer(), C.int(index), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan))
	}
}

func (ptr *QGridLayout) HasHeightForWidth() bool {
	defer qt.Recovering("QGridLayout::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QGridLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGridLayout) HeightForWidth(w int) int {
	defer qt.Recovering("QGridLayout::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QGridLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QGridLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QGridLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QGridLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQGridLayoutInvalidate
func callbackQGridLayoutInvalidate(ptrName *C.char) bool {
	defer qt.Recovering("callback QGridLayout::invalidate")

	var signal = qt.GetSignal(C.GoString(ptrName), "invalidate")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QGridLayout) ItemAt(index int) *QLayoutItem {
	defer qt.Recovering("QGridLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGridLayout) ItemAtPosition(row int, column int) *QLayoutItem {
	defer qt.Recovering("QGridLayout::itemAtPosition")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_ItemAtPosition(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QGridLayout) MinimumHeightForWidth(w int) int {
	defer qt.Recovering("QGridLayout::minimumHeightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QGridLayout) OriginCorner() core.Qt__Corner {
	defer qt.Recovering("QGridLayout::originCorner")

	if ptr.Pointer() != nil {
		return core.Qt__Corner(C.QGridLayout_OriginCorner(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) RowCount() int {
	defer qt.Recovering("QGridLayout::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) RowMinimumHeight(row int) int {
	defer qt.Recovering("QGridLayout::rowMinimumHeight")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_RowMinimumHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGridLayout) RowStretch(row int) int {
	defer qt.Recovering("QGridLayout::rowStretch")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_RowStretch(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGridLayout) SetColumnMinimumWidth(column int, minSize int) {
	defer qt.Recovering("QGridLayout::setColumnMinimumWidth")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetColumnMinimumWidth(ptr.Pointer(), C.int(column), C.int(minSize))
	}
}

func (ptr *QGridLayout) SetColumnStretch(column int, stretch int) {
	defer qt.Recovering("QGridLayout::setColumnStretch")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetColumnStretch(ptr.Pointer(), C.int(column), C.int(stretch))
	}
}

func (ptr *QGridLayout) SetOriginCorner(corner core.Qt__Corner) {
	defer qt.Recovering("QGridLayout::setOriginCorner")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetOriginCorner(ptr.Pointer(), C.int(corner))
	}
}

func (ptr *QGridLayout) SetRowMinimumHeight(row int, minSize int) {
	defer qt.Recovering("QGridLayout::setRowMinimumHeight")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetRowMinimumHeight(ptr.Pointer(), C.int(row), C.int(minSize))
	}
}

func (ptr *QGridLayout) SetRowStretch(row int, stretch int) {
	defer qt.Recovering("QGridLayout::setRowStretch")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetRowStretch(ptr.Pointer(), C.int(row), C.int(stretch))
	}
}

func (ptr *QGridLayout) SetSpacing(spacing int) {
	defer qt.Recovering("QGridLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QGridLayout_SetSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QGridLayout) Spacing() int {
	defer qt.Recovering("QGridLayout::spacing")

	if ptr.Pointer() != nil {
		return int(C.QGridLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGridLayout) TakeAt(index int) *QLayoutItem {
	defer qt.Recovering("QGridLayout::takeAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGridLayout) DestroyQGridLayout() {
	defer qt.Recovering("QGridLayout::~QGridLayout")

	if ptr.Pointer() != nil {
		C.QGridLayout_DestroyQGridLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
