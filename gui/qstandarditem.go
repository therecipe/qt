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

type QStandardItem_ITF interface {
	QStandardItem_PTR() *QStandardItem
}

func (p *QStandardItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStandardItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStandardItem(ptr QStandardItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItem_PTR().Pointer()
	}
	return nil
}

func NewQStandardItemFromPointer(ptr unsafe.Pointer) *QStandardItem {
	var n = new(QStandardItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStandardItem) QStandardItem_PTR() *QStandardItem {
	return ptr
}

//QStandardItem::ItemType
type QStandardItem__ItemType int64

const (
	QStandardItem__Type     = QStandardItem__ItemType(0)
	QStandardItem__UserType = QStandardItem__ItemType(1000)
)

func NewQStandardItem() *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem())
}

func NewQStandardItem3(icon QIcon_ITF, text string) *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem3(PointerFromQIcon(icon), C.CString(text)))
}

func NewQStandardItem2(text string) *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem2(C.CString(text)))
}

func NewQStandardItem4(rows int, columns int) *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem4(C.int(rows), C.int(columns)))
}

func (ptr *QStandardItem) AccessibleDescription() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AccessibleText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AppendRow2(item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) Background() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QStandardItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QStandardItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Child(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Child(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItem) Clone() *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Data(role int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItem) Flags() core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Foreground() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QStandardItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) HasChildren() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_HasChildren(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) Index() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItem_Index(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) InsertColumns(column int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertColumns(ptr.Pointer(), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) InsertRow2(row int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRow2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) InsertRows2(row int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRows2(ptr.Pointer(), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDragEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDropEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDropEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEditable() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsSelectable() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsSelectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsTristate() bool {
	if ptr.Pointer() != nil {
		return C.QStandardItem_IsTristate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) Model() *QStandardItemModel {
	if ptr.Pointer() != nil {
		return NewQStandardItemModelFromPointer(C.QStandardItem_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Parent() *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Read(in core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_Read(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QStandardItem) RemoveColumn(column int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QStandardItem) RemoveColumns(column int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) RemoveRow(row int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QStandardItem) RemoveRows(row int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRows(ptr.Pointer(), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleDescription(ptr.Pointer(), C.CString(accessibleDescription))
	}
}

func (ptr *QStandardItem) SetAccessibleText(accessibleText string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleText(ptr.Pointer(), C.CString(accessibleText))
	}
}

func (ptr *QStandardItem) SetBackground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QStandardItem) SetCheckable(checkable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QStandardItem) SetChild2(row int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetChild(row int, column int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild(ptr.Pointer(), C.int(row), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QStandardItem) SetData(value core.QVariant_ITF, role int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetData(ptr.Pointer(), core.PointerFromQVariant(value), C.int(role))
	}
}

func (ptr *QStandardItem) SetDragEnabled(dragEnabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(dragEnabled)))
	}
}

func (ptr *QStandardItem) SetDropEnabled(dropEnabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetDropEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(dropEnabled)))
	}
}

func (ptr *QStandardItem) SetEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QStandardItem) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStandardItem) SetFlags(flags core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QStandardItem) SetFont(font QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QStandardItem) SetForeground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetForeground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetIcon(icon QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

func (ptr *QStandardItem) SetRowCount(rows int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QStandardItem) SetSelectable(selectable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetSelectable(ptr.Pointer(), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QStandardItem) SetSizeHint(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QStandardItem) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QStandardItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QStandardItem) SetTextAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QStandardItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QStandardItem) SetTristate(tristate bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetTristate(ptr.Pointer(), C.int(qt.GoBoolToInt(tristate)))
	}
}

func (ptr *QStandardItem) SetWhatsThis(whatsThis string) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QStandardItem) SortChildren(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SortChildren(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QStandardItem) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) TakeChild(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_TakeChild(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) TextAlignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QStandardItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) Write(out core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_Write(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

func (ptr *QStandardItem) DestroyQStandardItem() {
	if ptr.Pointer() != nil {
		C.QStandardItem_DestroyQStandardItem(ptr.Pointer())
	}
}
