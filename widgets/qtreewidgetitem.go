package widgets

//#include "qtreewidgetitem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

type QTreeWidgetItem struct {
	ptr unsafe.Pointer
}

type QTreeWidgetItemITF interface {
	QTreeWidgetItemPTR() *QTreeWidgetItem
}

func (p *QTreeWidgetItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTreeWidgetItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTreeWidgetItem(ptr QTreeWidgetItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeWidgetItemPTR().Pointer()
	}
	return nil
}

func QTreeWidgetItemFromPointer(ptr unsafe.Pointer) *QTreeWidgetItem {
	var n = new(QTreeWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTreeWidgetItem) QTreeWidgetItemPTR() *QTreeWidgetItem {
	return ptr
}

//QTreeWidgetItem::ChildIndicatorPolicy
type QTreeWidgetItem__ChildIndicatorPolicy int

var (
	QTreeWidgetItem__ShowIndicator                  = QTreeWidgetItem__ChildIndicatorPolicy(0)
	QTreeWidgetItem__DontShowIndicator              = QTreeWidgetItem__ChildIndicatorPolicy(1)
	QTreeWidgetItem__DontShowIndicatorWhenChildless = QTreeWidgetItem__ChildIndicatorPolicy(2)
)

//QTreeWidgetItem::ItemType
type QTreeWidgetItem__ItemType int

var (
	QTreeWidgetItem__Type     = QTreeWidgetItem__ItemType(0)
	QTreeWidgetItem__UserType = QTreeWidgetItem__ItemType(1000)
)

func NewQTreeWidgetItem5(parent QTreeWidgetITF, preceding QTreeWidgetItemITF, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem5(C.QtObjectPtr(PointerFromQTreeWidget(parent)), C.QtObjectPtr(PointerFromQTreeWidgetItem(preceding)), C.int(ty))))
}

func NewQTreeWidgetItem4(parent QTreeWidgetITF, strin []string, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem4(C.QtObjectPtr(PointerFromQTreeWidget(parent)), C.CString(strings.Join(strin, "|")), C.int(ty))))
}

func NewQTreeWidgetItem3(parent QTreeWidgetITF, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem3(C.QtObjectPtr(PointerFromQTreeWidget(parent)), C.int(ty))))
}

func NewQTreeWidgetItem8(parent QTreeWidgetItemITF, preceding QTreeWidgetItemITF, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem8(C.QtObjectPtr(PointerFromQTreeWidgetItem(parent)), C.QtObjectPtr(PointerFromQTreeWidgetItem(preceding)), C.int(ty))))
}

func (ptr *QTreeWidgetItem) Flags() core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QTreeWidgetItem_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func NewQTreeWidgetItem7(parent QTreeWidgetItemITF, strin []string, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem7(C.QtObjectPtr(PointerFromQTreeWidgetItem(parent)), C.CString(strings.Join(strin, "|")), C.int(ty))))
}

func NewQTreeWidgetItem6(parent QTreeWidgetItemITF, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem6(C.QtObjectPtr(PointerFromQTreeWidgetItem(parent)), C.int(ty))))
}

func NewQTreeWidgetItem2(strin []string, ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem2(C.CString(strings.Join(strin, "|")), C.int(ty))))
}

func NewQTreeWidgetItem9(other QTreeWidgetItemITF) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem9(C.QtObjectPtr(PointerFromQTreeWidgetItem(other)))))
}

func NewQTreeWidgetItem(ty int) *QTreeWidgetItem {
	return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_NewQTreeWidgetItem(C.int(ty))))
}

func (ptr *QTreeWidgetItem) AddChild(child QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_AddChild(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(child)))
	}
}

func (ptr *QTreeWidgetItem) CheckState(column int) core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QTreeWidgetItem_CheckState(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QTreeWidgetItem) Child(index int) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_Child(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QTreeWidgetItem) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_ChildCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidgetItem) ChildIndicatorPolicy() QTreeWidgetItem__ChildIndicatorPolicy {
	if ptr.Pointer() != nil {
		return QTreeWidgetItem__ChildIndicatorPolicy(C.QTreeWidgetItem_ChildIndicatorPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidgetItem) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidgetItem) Data(column int, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_Data(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(role)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) Clone() *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_Clone(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeWidgetItem) IndexOfChild(child QTreeWidgetItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_IndexOfChild(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(child))))
	}
	return 0
}

func (ptr *QTreeWidgetItem) InsertChild(index int, child QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_InsertChild(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQTreeWidgetItem(child)))
	}
}

func (ptr *QTreeWidgetItem) IsDisabled() bool {
	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsDisabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsExpanded() bool {
	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsExpanded(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsFirstColumnSpanned() bool {
	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsFirstColumnSpanned(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsHidden(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsSelected() bool {
	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsSelected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) Parent() *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Read(in core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_Read(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(in)))
	}
}

func (ptr *QTreeWidgetItem) RemoveChild(child QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_RemoveChild(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(child)))
	}
}

func (ptr *QTreeWidgetItem) SetBackground(column int, brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetBackground(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QTreeWidgetItem) SetCheckState(column int, state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetCheckState(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(state))
	}
}

func (ptr *QTreeWidgetItem) SetChildIndicatorPolicy(policy QTreeWidgetItem__ChildIndicatorPolicy) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetChildIndicatorPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QTreeWidgetItem) SetData(column int, role int, value string) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetData(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(role), C.CString(value))
	}
}

func (ptr *QTreeWidgetItem) SetDisabled(disabled bool) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetDisabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(disabled)))
	}
}

func (ptr *QTreeWidgetItem) SetExpanded(expand bool) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetExpanded(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(expand)))
	}
}

func (ptr *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetFirstColumnSpanned(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeWidgetItem) SetFont(column int, font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetFont(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QTreeWidgetItem) SetForeground(column int, brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetForeground(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QTreeWidgetItem) SetHidden(hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetHidden(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeWidgetItem) SetIcon(column int, icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QTreeWidgetItem) SetSelected(sele bool) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetSelected(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTreeWidgetItem) SetSizeHint(column int, size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetSizeHint(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QTreeWidgetItem) SetStatusTip(column int, statusTip string) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetStatusTip(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.CString(statusTip))
	}
}

func (ptr *QTreeWidgetItem) SetText(column int, text string) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetText(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.CString(text))
	}
}

func (ptr *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetTextAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(alignment))
	}
}

func (ptr *QTreeWidgetItem) SetToolTip(column int, toolTip string) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.CString(toolTip))
	}
}

func (ptr *QTreeWidgetItem) SetWhatsThis(column int, whatsThis string) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.CString(whatsThis))
	}
}

func (ptr *QTreeWidgetItem) SortChildren(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SortChildren(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QTreeWidgetItem) StatusTip(column int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_StatusTip(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) TakeChild(index int) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidgetItem_TakeChild(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Text(column int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_Text(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) TextAlignment(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_TextAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QTreeWidgetItem) ToolTip(column int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_ToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) TreeWidget() *QTreeWidget {
	if ptr.Pointer() != nil {
		return QTreeWidgetFromPointer(unsafe.Pointer(C.QTreeWidgetItem_TreeWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidgetItem) WhatsThis(column int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_WhatsThis(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) Write(out core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(out)))
	}
}

func (ptr *QTreeWidgetItem) DestroyQTreeWidgetItem() {
	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_DestroyQTreeWidgetItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
