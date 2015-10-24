package widgets

//#include "qformlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFormLayout struct {
	QLayout
}

type QFormLayoutITF interface {
	QLayoutITF
	QFormLayoutPTR() *QFormLayout
}

func PointerFromQFormLayout(ptr QFormLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFormLayoutPTR().Pointer()
	}
	return nil
}

func QFormLayoutFromPointer(ptr unsafe.Pointer) *QFormLayout {
	var n = new(QFormLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFormLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFormLayout) QFormLayoutPTR() *QFormLayout {
	return ptr
}

//QFormLayout::FieldGrowthPolicy
type QFormLayout__FieldGrowthPolicy int

var (
	QFormLayout__FieldsStayAtSizeHint  = QFormLayout__FieldGrowthPolicy(0)
	QFormLayout__ExpandingFieldsGrow   = QFormLayout__FieldGrowthPolicy(1)
	QFormLayout__AllNonFixedFieldsGrow = QFormLayout__FieldGrowthPolicy(2)
)

//QFormLayout::ItemRole
type QFormLayout__ItemRole int

var (
	QFormLayout__LabelRole    = QFormLayout__ItemRole(0)
	QFormLayout__FieldRole    = QFormLayout__ItemRole(1)
	QFormLayout__SpanningRole = QFormLayout__ItemRole(2)
)

//QFormLayout::RowWrapPolicy
type QFormLayout__RowWrapPolicy int

var (
	QFormLayout__DontWrapRows = QFormLayout__RowWrapPolicy(0)
	QFormLayout__WrapLongRows = QFormLayout__RowWrapPolicy(1)
	QFormLayout__WrapAllRows  = QFormLayout__RowWrapPolicy(2)
)

func (ptr *QFormLayout) FieldGrowthPolicy() QFormLayout__FieldGrowthPolicy {
	if ptr.Pointer() != nil {
		return QFormLayout__FieldGrowthPolicy(C.QFormLayout_FieldGrowthPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) FormAlignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QFormLayout_FormAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) HorizontalSpacing() int {
	if ptr.Pointer() != nil {
		return int(C.QFormLayout_HorizontalSpacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) LabelAlignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QFormLayout_LabelAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) RowWrapPolicy() QFormLayout__RowWrapPolicy {
	if ptr.Pointer() != nil {
		return QFormLayout__RowWrapPolicy(C.QFormLayout_RowWrapPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) SetFieldGrowthPolicy(policy QFormLayout__FieldGrowthPolicy) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetFieldGrowthPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QFormLayout) SetFormAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetFormAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QFormLayout) SetHorizontalSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetHorizontalSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QFormLayout) SetLabelAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetLabelAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QFormLayout) SetRowWrapPolicy(policy QFormLayout__RowWrapPolicy) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetRowWrapPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QFormLayout) SetVerticalSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetVerticalSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QFormLayout) VerticalSpacing() int {
	if ptr.Pointer() != nil {
		return int(C.QFormLayout_VerticalSpacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQFormLayout(parent QWidgetITF) *QFormLayout {
	return QFormLayoutFromPointer(unsafe.Pointer(C.QFormLayout_NewQFormLayout(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QFormLayout) AddItem(item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QFormLayout) AddRow6(layout QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow6(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(layout)))
	}
}

func (ptr *QFormLayout) AddRow2(label QWidgetITF, field QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(label)), C.QtObjectPtr(PointerFromQLayout(field)))
	}
}

func (ptr *QFormLayout) AddRow(label QWidgetITF, field QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(label)), C.QtObjectPtr(PointerFromQWidget(field)))
	}
}

func (ptr *QFormLayout) AddRow5(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QFormLayout) AddRow4(labelText string, field QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow4(C.QtObjectPtr(ptr.Pointer()), C.CString(labelText), C.QtObjectPtr(PointerFromQLayout(field)))
	}
}

func (ptr *QFormLayout) AddRow3(labelText string, field QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow3(C.QtObjectPtr(ptr.Pointer()), C.CString(labelText), C.QtObjectPtr(PointerFromQWidget(field)))
	}
}

func (ptr *QFormLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QFormLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QFormLayout_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QFormLayout_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFormLayout) HeightForWidth(width int) int {
	if ptr.Pointer() != nil {
		return int(C.QFormLayout_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width)))
	}
	return 0
}

func (ptr *QFormLayout) InsertRow6(row int, layout QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow6(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQLayout(layout)))
	}
}

func (ptr *QFormLayout) InsertRow2(row int, label QWidgetITF, field QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQWidget(label)), C.QtObjectPtr(PointerFromQLayout(field)))
	}
}

func (ptr *QFormLayout) InsertRow(row int, label QWidgetITF, field QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQWidget(label)), C.QtObjectPtr(PointerFromQWidget(field)))
	}
}

func (ptr *QFormLayout) InsertRow5(row int, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow5(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QFormLayout) InsertRow4(row int, labelText string, field QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow4(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.CString(labelText), C.QtObjectPtr(PointerFromQLayout(field)))
	}
}

func (ptr *QFormLayout) InsertRow3(row int, labelText string, field QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow3(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.CString(labelText), C.QtObjectPtr(PointerFromQWidget(field)))
	}
}

func (ptr *QFormLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QFormLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QFormLayout) ItemAt2(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QFormLayout_ItemAt2(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QFormLayout) ItemAt(row int, role QFormLayout__ItemRole) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QFormLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(role))))
	}
	return nil
}

func (ptr *QFormLayout) LabelForField2(field QLayoutITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QFormLayout_LabelForField2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(field)))))
	}
	return nil
}

func (ptr *QFormLayout) LabelForField(field QWidgetITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QFormLayout_LabelForField(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(field)))))
	}
	return nil
}

func (ptr *QFormLayout) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QFormLayout_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) SetGeometry(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QFormLayout) SetItem(row int, role QFormLayout__ItemRole, item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(role), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QFormLayout) SetLayout(row int, role QFormLayout__ItemRole, layout QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetLayout(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(role), C.QtObjectPtr(PointerFromQLayout(layout)))
	}
}

func (ptr *QFormLayout) SetSpacing(spacing int) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(spacing))
	}
}

func (ptr *QFormLayout) SetWidget(row int, role QFormLayout__ItemRole, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFormLayout_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(role), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QFormLayout) Spacing() int {
	if ptr.Pointer() != nil {
		return int(C.QFormLayout_Spacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFormLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QFormLayout_TakeAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QFormLayout) DestroyQFormLayout() {
	if ptr.Pointer() != nil {
		C.QFormLayout_DestroyQFormLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
