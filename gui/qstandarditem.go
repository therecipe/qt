package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::QStandardItem")
		}
	}()

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem())
}

func NewQStandardItem3(icon QIcon_ITF, text string) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::QStandardItem")
		}
	}()

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem3(PointerFromQIcon(icon), C.CString(text)))
}

func NewQStandardItem2(text string) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::QStandardItem")
		}
	}()

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem2(C.CString(text)))
}

func NewQStandardItem4(rows int, columns int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::QStandardItem")
		}
	}()

	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem4(C.int(rows), C.int(columns)))
}

func (ptr *QStandardItem) AccessibleDescription() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::accessibleDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AccessibleText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::accessibleText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_AccessibleText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AppendRow2(item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::appendRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) Background() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::background")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QStandardItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) CheckState() core.Qt__CheckState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::checkState")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QStandardItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Child(row int, column int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::child")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Child(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItem) Clone() *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::clone")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Column() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::column")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Data(role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItem) Flags() core.Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) Foreground() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::foreground")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QStandardItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) HasChildren() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::hasChildren")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_HasChildren(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) Index() *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::index")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItem_Index(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) InsertColumns(column int, count int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::insertColumns")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_InsertColumns(ptr.Pointer(), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) InsertRow2(row int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRow2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) InsertRows2(row int, count int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::insertRows")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRows2(ptr.Pointer(), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) IsCheckable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isCheckable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDragEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isDragEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsDropEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isDropEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsDropEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEditable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isEditable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsSelectable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isSelectable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsSelectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) IsTristate() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::isTristate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItem_IsTristate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStandardItem) Model() *QStandardItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::model")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemModelFromPointer(C.QStandardItem_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Parent() *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItem) Read(in core.QDataStream_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::read")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_Read(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QStandardItem) RemoveColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::removeColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QStandardItem) RemoveColumns(column int, count int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::removeColumns")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count))
	}
}

func (ptr *QStandardItem) RemoveRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::removeRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QStandardItem) RemoveRows(row int, count int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::removeRows")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRows(ptr.Pointer(), C.int(row), C.int(count))
	}
}

func (ptr *QStandardItem) Row() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::row")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) RowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setAccessibleDescription")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleDescription(ptr.Pointer(), C.CString(accessibleDescription))
	}
}

func (ptr *QStandardItem) SetAccessibleText(accessibleText string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setAccessibleText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetAccessibleText(ptr.Pointer(), C.CString(accessibleText))
	}
}

func (ptr *QStandardItem) SetBackground(brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetCheckState(state core.Qt__CheckState) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setCheckState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QStandardItem) SetCheckable(checkable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setCheckable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QStandardItem) SetChild2(row int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setChild")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetChild(row int, column int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setChild")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild(ptr.Pointer(), C.int(row), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetColumnCount(columns int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setColumnCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QStandardItem) SetData(value core.QVariant_ITF, role int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetData(ptr.Pointer(), core.PointerFromQVariant(value), C.int(role))
	}
}

func (ptr *QStandardItem) SetDragEnabled(dragEnabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setDragEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(dragEnabled)))
	}
}

func (ptr *QStandardItem) SetDropEnabled(dropEnabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setDropEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetDropEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(dropEnabled)))
	}
}

func (ptr *QStandardItem) SetEditable(editable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setEditable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QStandardItem) SetEnabled(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStandardItem) SetFlags(flags core.Qt__ItemFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QStandardItem) SetFont(font QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QStandardItem) SetForeground(brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setForeground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetForeground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetIcon(icon QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

func (ptr *QStandardItem) SetRowCount(rows int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setRowCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QStandardItem) SetSelectable(selectable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setSelectable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetSelectable(ptr.Pointer(), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QStandardItem) SetSizeHint(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setSizeHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QStandardItem) SetStatusTip(statusTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setStatusTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QStandardItem) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QStandardItem) SetTextAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setTextAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QStandardItem) SetToolTip(toolTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QStandardItem) SetTristate(tristate bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setTristate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetTristate(ptr.Pointer(), C.int(qt.GoBoolToInt(tristate)))
	}
}

func (ptr *QStandardItem) SetWhatsThis(whatsThis string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::setWhatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QStandardItem) SortChildren(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::sortChildren")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_SortChildren(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QStandardItem) StatusTip() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::statusTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) TakeChild(row int, column int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::takeChild")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_TakeChild(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItem) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) TextAlignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::textAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QStandardItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) ToolTip() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::toolTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItem) WhatsThis() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::whatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) Write(out core.QDataStream_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::write")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_Write(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

func (ptr *QStandardItem) DestroyQStandardItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItem::~QStandardItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItem_DestroyQStandardItem(ptr.Pointer())
	}
}
