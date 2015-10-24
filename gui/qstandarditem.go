package gui

//#include "qstandarditem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStandardItem struct {
	ptr unsafe.Pointer
}

type QStandardItemITF interface {
	QStandardItemPTR() *QStandardItem
}

func (p *QStandardItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStandardItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStandardItem(ptr QStandardItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemPTR().Pointer()
	}
	return nil
}

func QStandardItemFromPointer(ptr unsafe.Pointer) *QStandardItem {
	var n = new(QStandardItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStandardItem) QStandardItemPTR() *QStandardItem {
	return ptr
}

//QStandardItem::ItemType
type QStandardItem__ItemType int

var (
	QStandardItem__Type     = QStandardItem__ItemType(0)
	QStandardItem__UserType = QStandardItem__ItemType(1000)
)

func NewQStandardItem() *QStandardItem {
	return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_NewQStandardItem()))
}

func NewQStandardItem3(icon QIconITF, text string) *QStandardItem {
	return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_NewQStandardItem3(C.QtObjectPtr(PointerFromQIcon(icon)), C.CString(text))))
}

func NewQStandardItem2(text string) *QStandardItem {
	return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_NewQStandardItem2(C.CString(text))))
}

func NewQStandardItem4(rows int, columns int) *QStandardItem {
	return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_NewQStandardItem4(C.int(rows), C.int(columns))))
}

func (ptr *QStandardItem) AccessibleDescription() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleDescription(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStandardItem) AccessibleText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStandardItem) AppendRow2(item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRow2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItem) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QStandardItem_CheckState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) Child(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_Child(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QStandardItem) Clone() *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_Clone(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStandardItem) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Column(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) Data(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_Data(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QStandardItem) Flags() core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItem_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) HasChildren() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_HasChildren(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) Index() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QStandardItem_Index(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStandardItem) InsertColumns(column int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) InsertRow2(row int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRow2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItem) InsertRows2(row int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRows2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsCheckable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDragEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDragEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDropEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDropEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEditable() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEditable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) IsSelectable() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsSelectable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) IsTristate() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsTristate(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) Model() *QStandardItemModel {
	if ptr.Pointer() != nil {
		return QStandardItemModelFromPointer(unsafe.Pointer(C.QStandardItem_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStandardItem) Parent() *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStandardItem) Read(in core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_Read(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(in)))
	}
}

func (ptr *QStandardItem) RemoveColumn(column int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QStandardItem) RemoveColumns(column int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) RemoveRow(row int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QStandardItem) RemoveRows(row int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Row(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(accessibleDescription))
	}
}

func (ptr *QStandardItem) SetAccessibleText(accessibleText string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleText(C.QtObjectPtr(ptr.Pointer()), C.CString(accessibleText))
	}
}

func (ptr *QStandardItem) SetBackground(brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetBackground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QStandardItem) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QStandardItem) SetCheckable(checkable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QStandardItem) SetChild2(row int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItem) SetChild(row int, column int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItem) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetColumnCount(C.QtObjectPtr(ptr.Pointer()), C.int(columns))
	}
}

func (ptr *QStandardItem) SetData(value string, role int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetData(C.QtObjectPtr(ptr.Pointer()), C.CString(value), C.int(role))
	}
}

func (ptr *QStandardItem) SetDragEnabled(dragEnabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetDragEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(dragEnabled)))
	}
}

func (ptr *QStandardItem) SetDropEnabled(dropEnabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetDropEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(dropEnabled)))
	}
}

func (ptr *QStandardItem) SetEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetEditable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QStandardItem) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStandardItem) SetFlags(flags core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QStandardItem) SetFont(font QFontITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font)))
	}
}

func (ptr *QStandardItem) SetForeground(brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetForeground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QStandardItem) SetIcon(icon QIconITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIcon(icon)))
	}
}

func (ptr *QStandardItem) SetRowCount(rows int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetRowCount(C.QtObjectPtr(ptr.Pointer()), C.int(rows))
	}
}

func (ptr *QStandardItem) SetSelectable(selectable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetSelectable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QStandardItem) SetSizeHint(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetSizeHint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QStandardItem) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetStatusTip(C.QtObjectPtr(ptr.Pointer()), C.CString(statusTip))
	}
}

func (ptr *QStandardItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QStandardItem) SetTextAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetTextAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QStandardItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(toolTip))
	}
}

func (ptr *QStandardItem) SetTristate(tristate bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetTristate(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(tristate)))
	}
}

func (ptr *QStandardItem) SetWhatsThis(whatsThis string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.CString(whatsThis))
	}
}

func (ptr *QStandardItem) SortChildren(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SortChildren(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QStandardItem) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_StatusTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStandardItem) TakeChild(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItem_TakeChild(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QStandardItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStandardItem) TextAlignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QStandardItem_TextAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStandardItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_WhatsThis(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStandardItem) Write(out core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(out)))
	}
}

func (ptr *QStandardItem) DestroyQStandardItem() {
	if ptr.Pointer() != nil {
		C.QStandardItem_DestroyQStandardItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
