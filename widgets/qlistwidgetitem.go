package widgets

//#include "qlistwidgetitem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QListWidgetItem struct {
	ptr unsafe.Pointer
}

type QListWidgetItem_ITF interface {
	QListWidgetItem_PTR() *QListWidgetItem
}

func (p *QListWidgetItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QListWidgetItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQListWidgetItem(ptr QListWidgetItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListWidgetItem_PTR().Pointer()
	}
	return nil
}

func NewQListWidgetItemFromPointer(ptr unsafe.Pointer) *QListWidgetItem {
	var n = new(QListWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QListWidgetItem) QListWidgetItem_PTR() *QListWidgetItem {
	return ptr
}

//QListWidgetItem::ItemType
type QListWidgetItem__ItemType int64

const (
	QListWidgetItem__Type     = QListWidgetItem__ItemType(0)
	QListWidgetItem__UserType = QListWidgetItem__ItemType(1000)
)

func NewQListWidgetItem(parent QListWidget_ITF, ty int) *QListWidgetItem {
	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem(PointerFromQListWidget(parent), C.int(ty)))
}

func NewQListWidgetItem3(icon gui.QIcon_ITF, text string, parent QListWidget_ITF, ty int) *QListWidgetItem {
	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem3(gui.PointerFromQIcon(icon), C.CString(text), PointerFromQListWidget(parent), C.int(ty)))
}

func NewQListWidgetItem2(text string, parent QListWidget_ITF, ty int) *QListWidgetItem {
	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem2(C.CString(text), PointerFromQListWidget(parent), C.int(ty)))
}

func (ptr *QListWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func NewQListWidgetItem4(other QListWidgetItem_ITF) *QListWidgetItem {
	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem4(PointerFromQListWidgetItem(other)))
}

func (ptr *QListWidgetItem) Background() *gui.QBrush {
	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QListWidgetItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QListWidgetItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) Clone() *QListWidgetItem {
	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidgetItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) Data(role int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QListWidgetItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QListWidgetItem) Flags() core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QListWidgetItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) Foreground() *gui.QBrush {
	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QListWidgetItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidgetItem) IsSelected() bool {
	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidgetItem) ListWidget() *QListWidget {
	if ptr.Pointer() != nil {
		return NewQListWidgetFromPointer(C.QListWidgetItem_ListWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) Read(in core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_Read(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QListWidgetItem) SetBackground(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetBackground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QListWidgetItem) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QListWidgetItem) SetData(role int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetData(ptr.Pointer(), C.int(role), core.PointerFromQVariant(value))
	}
}

func (ptr *QListWidgetItem) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QListWidgetItem) SetForeground(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetForeground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QListWidgetItem) SetHidden(hide bool) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QListWidgetItem) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QListWidgetItem) SetSelected(sele bool) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSelected(ptr.Pointer(), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QListWidgetItem) SetSizeHint(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QListWidgetItem) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QListWidgetItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QListWidgetItem) SetTextAlignment(alignment int) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QListWidgetItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QListWidgetItem) SetWhatsThis(whatsThis string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QListWidgetItem) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) TextAlignment() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Write(out core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_Write(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

func (ptr *QListWidgetItem) DestroyQListWidgetItem() {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_DestroyQListWidgetItem(ptr.Pointer())
	}
}
