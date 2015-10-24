package widgets

//#include "qgridlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGridLayout struct {
	QLayout
}

type QGridLayoutITF interface {
	QLayoutITF
	QGridLayoutPTR() *QGridLayout
}

func PointerFromQGridLayout(ptr QGridLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGridLayoutPTR().Pointer()
	}
	return nil
}

func QGridLayoutFromPointer(ptr unsafe.Pointer) *QGridLayout {
	var n = new(QGridLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGridLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGridLayout) QGridLayoutPTR() *QGridLayout {
	return ptr
}

func (ptr *QGridLayout) HorizontalSpacing() int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_HorizontalSpacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) SetHorizontalSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetHorizontalSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QGridLayout) SetVerticalSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetVerticalSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QGridLayout) VerticalSpacing() int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_VerticalSpacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQGridLayout2() *QGridLayout {
	return QGridLayoutFromPointer(unsafe.Pointer(C.QGridLayout_NewQGridLayout2()))
}

func NewQGridLayout(parent QWidgetITF) *QGridLayout {
	return QGridLayoutFromPointer(unsafe.Pointer(C.QGridLayout_NewQGridLayout(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QGridLayout) AddItem(item QLayoutItemITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayoutItem(item)), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddLayout(layout QLayoutITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(layout)), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddLayout2(layout QLayoutITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddLayout2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(layout)), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddWidget2(widget QWidgetITF, fromRow int, fromColumn int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddWidget2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(fromRow), C.int(fromColumn), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGridLayout) AddWidget(widget QWidgetITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGridLayout) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) ColumnMinimumWidth(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_ColumnMinimumWidth(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QGridLayout) ColumnStretch(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_ColumnStretch(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QGridLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGridLayout_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) GetItemPosition(index int, row int, column int, rowSpan int, columnSpan int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_GetItemPosition(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan))
	}
}

func (ptr *QGridLayout) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QGridLayout_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGridLayout) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QGridLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QGridLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGridLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QGridLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QGridLayout) ItemAtPosition(row int, column int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QGridLayout_ItemAtPosition(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QGridLayout) MinimumHeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_MinimumHeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QGridLayout) OriginCorner() core.Qt__Corner {
	if ptr.Pointer() != nil {
		return core.Qt__Corner(C.QGridLayout_OriginCorner(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) RowMinimumHeight(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_RowMinimumHeight(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QGridLayout) RowStretch(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_RowStretch(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QGridLayout) SetColumnMinimumWidth(column int, minSize int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetColumnMinimumWidth(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(minSize))
	}
}

func (ptr *QGridLayout) SetColumnStretch(column int, stretch int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetColumnStretch(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(stretch))
	}
}

func (ptr *QGridLayout) SetGeometry(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QGridLayout) SetOriginCorner(corner core.Qt__Corner) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetOriginCorner(C.QtObjectPtr(ptr.Pointer()), C.int(corner))
	}
}

func (ptr *QGridLayout) SetRowMinimumHeight(row int, minSize int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetRowMinimumHeight(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(minSize))
	}
}

func (ptr *QGridLayout) SetRowStretch(row int, stretch int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetRowStretch(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(stretch))
	}
}

func (ptr *QGridLayout) SetSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QGridLayout_SetSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QGridLayout) Spacing() int {
	if ptr.Pointer() != nil {
		return int(C.QGridLayout_Spacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QGridLayout_TakeAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QGridLayout) DestroyQGridLayout() {
	if ptr.Pointer() != nil {
		C.QGridLayout_DestroyQGridLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
