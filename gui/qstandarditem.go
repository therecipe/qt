package gui

//#include "gui.h"
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
	for len(n.ObjectNameAbs()) < len("QStandardItem_") {
		n.SetObjectNameAbs("QStandardItem_" + qt.Identifier())
	}
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
	defer qt.Recovering("QStandardItem::QStandardItem")

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem())
}

func NewQStandardItem3(icon QIcon_ITF, text string) *QStandardItem {
	defer qt.Recovering("QStandardItem::QStandardItem")

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem3(PointerFromQIcon(icon), C.CString(text)))
}

func NewQStandardItem2(text string) *QStandardItem {
	defer qt.Recovering("QStandardItem::QStandardItem")

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem2(C.CString(text)))
}

func NewQStandardItem4(rows int, columns int) *QStandardItem {
	defer qt.Recovering("QStandardItem::QStandardItem")

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem4(C.int(rows), C.int(columns)))
}

func (ptr *QStandardItem) AccessibleDescription() string {
	defer qt.Recovering("QStandardItem::accessibleDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AccessibleText() string {
	defer qt.Recovering("QStandardItem::accessibleText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AppendRow2(item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItem::appendRow")

	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) Background() *QBrush {
	defer qt.Recovering("QStandardItem::background")

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QStandardItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) CheckState() core.Qt__CheckState {
	defer qt.Recovering("QStandardItem::checkState")

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QStandardItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Child(row int, column int) *QStandardItem {
	defer qt.Recovering("QStandardItem::child")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Child(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItem) Clone() *QStandardItem {
	defer qt.Recovering("QStandardItem::clone")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Column() int {
	defer qt.Recovering("QStandardItem::column")

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) ColumnCount() int {
	defer qt.Recovering("QStandardItem::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Data(role int) *core.QVariant {
	defer qt.Recovering("QStandardItem::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItem) Flags() core.Qt__ItemFlag {
	defer qt.Recovering("QStandardItem::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Foreground() *QBrush {
	defer qt.Recovering("QStandardItem::foreground")

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QStandardItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) HasChildren() bool {
	defer qt.Recovering("QStandardItem::hasChildren")

	if ptr.Pointer() != nil {
		return C.QStandardItem_HasChildren(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) Icon() *QIcon {
	defer qt.Recovering("QStandardItem::icon")

	if ptr.Pointer() != nil {
		return NewQIconFromPointer(C.QStandardItem_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Index() *core.QModelIndex {
	defer qt.Recovering("QStandardItem::index")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItem_Index(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) InsertColumns(column int, count int) {
	defer qt.Recovering("QStandardItem::insertColumns")

	if ptr.Pointer() != nil {
		C.QStandardItem_InsertColumns(ptr.Pointer(), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) InsertRow2(row int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItem::insertRow")

	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRow2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) InsertRows2(row int, count int) {
	defer qt.Recovering("QStandardItem::insertRows")

	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRows2(ptr.Pointer(), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) IsCheckable() bool {
	defer qt.Recovering("QStandardItem::isCheckable")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDragEnabled() bool {
	defer qt.Recovering("QStandardItem::isDragEnabled")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDropEnabled() bool {
	defer qt.Recovering("QStandardItem::isDropEnabled")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDropEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEditable() bool {
	defer qt.Recovering("QStandardItem::isEditable")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEnabled() bool {
	defer qt.Recovering("QStandardItem::isEnabled")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsSelectable() bool {
	defer qt.Recovering("QStandardItem::isSelectable")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsSelectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsTristate() bool {
	defer qt.Recovering("QStandardItem::isTristate")

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsTristate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) Model() *QStandardItemModel {
	defer qt.Recovering("QStandardItem::model")

	if ptr.Pointer() != nil {
		return NewQStandardItemModelFromPointer(C.QStandardItem_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Parent() *QStandardItem {
	defer qt.Recovering("QStandardItem::parent")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) RemoveColumn(column int) {
	defer qt.Recovering("QStandardItem::removeColumn")

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QStandardItem) RemoveColumns(column int, count int) {
	defer qt.Recovering("QStandardItem::removeColumns")

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) RemoveRow(row int) {
	defer qt.Recovering("QStandardItem::removeRow")

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QStandardItem) RemoveRows(row int, count int) {
	defer qt.Recovering("QStandardItem::removeRows")

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRows(ptr.Pointer(), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) Row() int {
	defer qt.Recovering("QStandardItem::row")

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) RowCount() int {
	defer qt.Recovering("QStandardItem::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	defer qt.Recovering("QStandardItem::setAccessibleDescription")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleDescription(ptr.Pointer(), C.CString(accessibleDescription))
	}
}

func (ptr *QStandardItem) SetAccessibleText(accessibleText string) {
	defer qt.Recovering("QStandardItem::setAccessibleText")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleText(ptr.Pointer(), C.CString(accessibleText))
	}
}

func (ptr *QStandardItem) SetBackground(brush QBrush_ITF) {
	defer qt.Recovering("QStandardItem::setBackground")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetCheckState(state core.Qt__CheckState) {
	defer qt.Recovering("QStandardItem::setCheckState")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QStandardItem) SetCheckable(checkable bool) {
	defer qt.Recovering("QStandardItem::setCheckable")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QStandardItem) SetChild2(row int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItem::setChild")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetChild(row int, column int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItem::setChild")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild(ptr.Pointer(), C.int(row), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetColumnCount(columns int) {
	defer qt.Recovering("QStandardItem::setColumnCount")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QStandardItem) ConnectSetData(f func(value *core.QVariant, role int)) {
	defer qt.Recovering("connect QStandardItem::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData", f)
	}
}

func (ptr *QStandardItem) DisconnectSetData() {
	defer qt.Recovering("disconnect QStandardItem::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData")
	}
}

//export callbackQStandardItemSetData
func callbackQStandardItemSetData(ptrName *C.char, value unsafe.Pointer, role C.int) bool {
	defer qt.Recovering("callback QStandardItem::setData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setData"); signal != nil {
		signal.(func(*core.QVariant, int))(core.NewQVariantFromPointer(value), int(role))
		return true
	}
	return false

}

func (ptr *QStandardItem) SetDragEnabled(dragEnabled bool) {
	defer qt.Recovering("QStandardItem::setDragEnabled")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(dragEnabled)))
	}
}

func (ptr *QStandardItem) SetDropEnabled(dropEnabled bool) {
	defer qt.Recovering("QStandardItem::setDropEnabled")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetDropEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(dropEnabled)))
	}
}

func (ptr *QStandardItem) SetEditable(editable bool) {
	defer qt.Recovering("QStandardItem::setEditable")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QStandardItem) SetEnabled(enabled bool) {
	defer qt.Recovering("QStandardItem::setEnabled")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStandardItem) SetFlags(flags core.Qt__ItemFlag) {
	defer qt.Recovering("QStandardItem::setFlags")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QStandardItem) SetFont(font QFont_ITF) {
	defer qt.Recovering("QStandardItem::setFont")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QStandardItem) SetForeground(brush QBrush_ITF) {
	defer qt.Recovering("QStandardItem::setForeground")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetForeground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetIcon(icon QIcon_ITF) {
	defer qt.Recovering("QStandardItem::setIcon")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

func (ptr *QStandardItem) SetRowCount(rows int) {
	defer qt.Recovering("QStandardItem::setRowCount")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QStandardItem) SetSelectable(selectable bool) {
	defer qt.Recovering("QStandardItem::setSelectable")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetSelectable(ptr.Pointer(), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QStandardItem) SetSizeHint(size core.QSize_ITF) {
	defer qt.Recovering("QStandardItem::setSizeHint")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QStandardItem) SetStatusTip(statusTip string) {
	defer qt.Recovering("QStandardItem::setStatusTip")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QStandardItem) SetText(text string) {
	defer qt.Recovering("QStandardItem::setText")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QStandardItem) SetTextAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QStandardItem::setTextAlignment")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QStandardItem) SetToolTip(toolTip string) {
	defer qt.Recovering("QStandardItem::setToolTip")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QStandardItem) SetTristate(tristate bool) {
	defer qt.Recovering("QStandardItem::setTristate")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetTristate(ptr.Pointer(), C.int(qt.GoBoolToInt(tristate)))
	}
}

func (ptr *QStandardItem) SetWhatsThis(whatsThis string) {
	defer qt.Recovering("QStandardItem::setWhatsThis")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QStandardItem) SizeHint() *core.QSize {
	defer qt.Recovering("QStandardItem::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QStandardItem_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) SortChildren(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QStandardItem::sortChildren")

	if ptr.Pointer() != nil {
		C.QStandardItem_SortChildren(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QStandardItem) StatusTip() string {
	defer qt.Recovering("QStandardItem::statusTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) TakeChild(row int, column int) *QStandardItem {
	defer qt.Recovering("QStandardItem::takeChild")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_TakeChild(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItem) Text() string {
	defer qt.Recovering("QStandardItem::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) TextAlignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QStandardItem::textAlignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QStandardItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) ToolTip() string {
	defer qt.Recovering("QStandardItem::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) Type() int {
	defer qt.Recovering("QStandardItem::type")

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) WhatsThis() string {
	defer qt.Recovering("QStandardItem::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) DestroyQStandardItem() {
	defer qt.Recovering("QStandardItem::~QStandardItem")

	if ptr.Pointer() != nil {
		C.QStandardItem_DestroyQStandardItem(ptr.Pointer())
	}
}

func (ptr *QStandardItem) ObjectNameAbs() string {
	defer qt.Recovering("QStandardItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QStandardItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QStandardItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
