package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsGridLayout struct {
	QGraphicsLayout
}

type QGraphicsGridLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsGridLayout_PTR() *QGraphicsGridLayout
}

func PointerFromQGraphicsGridLayout(ptr QGraphicsGridLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsGridLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsGridLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsGridLayout {
	var n = new(QGraphicsGridLayout)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGraphicsGridLayout_") {
		n.SetObjectNameAbs("QGraphicsGridLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsGridLayout) QGraphicsGridLayout_PTR() *QGraphicsGridLayout {
	return ptr
}

func NewQGraphicsGridLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsGridLayout {
	defer qt.Recovering("QGraphicsGridLayout::QGraphicsGridLayout")

	return NewQGraphicsGridLayoutFromPointer(C.QGraphicsGridLayout_NewQGraphicsGridLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsGridLayout) AddItem2(item QGraphicsLayoutItem_ITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsGridLayout::addItem")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_AddItem2(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) AddItem(item QGraphicsLayoutItem_ITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsGridLayout::addItem")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_AddItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) Alignment(item QGraphicsLayoutItem_ITF) core.Qt__AlignmentFlag {
	defer qt.Recovering("QGraphicsGridLayout::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_Alignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnAlignment(column int) core.Qt__AlignmentFlag {
	defer qt.Recovering("QGraphicsGridLayout::columnAlignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_ColumnAlignment(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnCount() int {
	defer qt.Recovering("QGraphicsGridLayout::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::columnMaximumWidth")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnMaximumWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::columnMinimumWidth")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnMinimumWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::columnPreferredWidth")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnPreferredWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::columnSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnSpacing(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	defer qt.Recovering("QGraphicsGridLayout::columnStretchFactor")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_ColumnStretchFactor(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) Count() int {
	defer qt.Recovering("QGraphicsGridLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) HorizontalSpacing() float64 {
	defer qt.Recovering("QGraphicsGridLayout::horizontalSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_HorizontalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QGraphicsGridLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "invalidate", f)
	}
}

func (ptr *QGraphicsGridLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QGraphicsGridLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "invalidate")
	}
}

//export callbackQGraphicsGridLayoutInvalidate
func callbackQGraphicsGridLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsGridLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsGridLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QGraphicsGridLayout) Invalidate() {
	defer qt.Recovering("QGraphicsGridLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsGridLayout) InvalidateDefault() {
	defer qt.Recovering("QGraphicsGridLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem {
	defer qt.Recovering("QGraphicsGridLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsGridLayout_ItemAt(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QGraphicsGridLayout) ConnectRemoveAt(f func(index int)) {
	defer qt.Recovering("connect QGraphicsGridLayout::removeAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "removeAt", f)
	}
}

func (ptr *QGraphicsGridLayout) DisconnectRemoveAt() {
	defer qt.Recovering("disconnect QGraphicsGridLayout::removeAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "removeAt")
	}
}

//export callbackQGraphicsGridLayoutRemoveAt
func callbackQGraphicsGridLayoutRemoveAt(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QGraphicsGridLayout::removeAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "removeAt"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QGraphicsGridLayout) RemoveAt(index int) {
	defer qt.Recovering("QGraphicsGridLayout::removeAt")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsGridLayout) RemoveAtDefault(index int) {
	defer qt.Recovering("QGraphicsGridLayout::removeAt")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveAtDefault(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsGridLayout) RemoveItem(item QGraphicsLayoutItem_ITF) {
	defer qt.Recovering("QGraphicsGridLayout::removeItem")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsGridLayout) RowAlignment(row int) core.Qt__AlignmentFlag {
	defer qt.Recovering("QGraphicsGridLayout::rowAlignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_RowAlignment(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowCount() int {
	defer qt.Recovering("QGraphicsGridLayout::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::rowMaximumHeight")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowMaximumHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::rowMinimumHeight")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowMinimumHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::rowPreferredHeight")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowPreferredHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowSpacing(row int) float64 {
	defer qt.Recovering("QGraphicsGridLayout::rowSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowSpacing(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowStretchFactor(row int) int {
	defer qt.Recovering("QGraphicsGridLayout::rowStretchFactor")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_RowStretchFactor(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) SetAlignment(item QGraphicsLayoutItem_ITF, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsGridLayout::setAlignment")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetAlignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnAlignment(column int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnAlignment")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnAlignment(ptr.Pointer(), C.int(column), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnFixedWidth")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnFixedWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnMaximumWidth")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnMaximumWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnMinimumWidth")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnMinimumWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnPreferredWidth")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnPreferredWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnSpacing(ptr.Pointer(), C.int(column), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	defer qt.Recovering("QGraphicsGridLayout::setColumnStretchFactor")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnStretchFactor(ptr.Pointer(), C.int(column), C.int(stretch))
	}
}

func (ptr *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsGridLayout::setHorizontalSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetHorizontalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetRowAlignment(row int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsGridLayout::setRowAlignment")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowAlignment(ptr.Pointer(), C.int(row), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	defer qt.Recovering("QGraphicsGridLayout::setRowFixedHeight")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowFixedHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	defer qt.Recovering("QGraphicsGridLayout::setRowMaximumHeight")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowMaximumHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	defer qt.Recovering("QGraphicsGridLayout::setRowMinimumHeight")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowMinimumHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	defer qt.Recovering("QGraphicsGridLayout::setRowPreferredHeight")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowPreferredHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	defer qt.Recovering("QGraphicsGridLayout::setRowSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowSpacing(ptr.Pointer(), C.int(row), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	defer qt.Recovering("QGraphicsGridLayout::setRowStretchFactor")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowStretchFactor(ptr.Pointer(), C.int(row), C.int(stretch))
	}
}

func (ptr *QGraphicsGridLayout) SetSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsGridLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	defer qt.Recovering("QGraphicsGridLayout::setVerticalSpacing")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetVerticalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) VerticalSpacing() float64 {
	defer qt.Recovering("QGraphicsGridLayout::verticalSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_VerticalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) DestroyQGraphicsGridLayout() {
	defer qt.Recovering("QGraphicsGridLayout::~QGraphicsGridLayout")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_DestroyQGraphicsGridLayout(ptr.Pointer())
	}
}

func (ptr *QGraphicsGridLayout) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsGridLayout::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsGridLayout_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsGridLayout) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsGridLayout::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
