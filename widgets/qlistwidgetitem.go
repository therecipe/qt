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

type QListWidgetItemITF interface {
	QListWidgetItemPTR() *QListWidgetItem
}

func (p *QListWidgetItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QListWidgetItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQListWidgetItem(ptr QListWidgetItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListWidgetItemPTR().Pointer()
	}
	return nil
}

func QListWidgetItemFromPointer(ptr unsafe.Pointer) *QListWidgetItem {
	var n = new(QListWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QListWidgetItem) QListWidgetItemPTR() *QListWidgetItem {
	return ptr
}

//QListWidgetItem::ItemType
type QListWidgetItem__ItemType int

var (
	QListWidgetItem__Type     = QListWidgetItem__ItemType(0)
	QListWidgetItem__UserType = QListWidgetItem__ItemType(1000)
)

func NewQListWidgetItem(parent QListWidgetITF, ty int) *QListWidgetItem {
	return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidgetItem_NewQListWidgetItem(C.QtObjectPtr(PointerFromQListWidget(parent)), C.int(ty))))
}

func NewQListWidgetItem3(icon gui.QIconITF, text string, parent QListWidgetITF, ty int) *QListWidgetItem {
	return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidgetItem_NewQListWidgetItem3(C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.QtObjectPtr(PointerFromQListWidget(parent)), C.int(ty))))
}

func NewQListWidgetItem2(text string, parent QListWidgetITF, ty int) *QListWidgetItem {
	return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidgetItem_NewQListWidgetItem2(C.CString(text), C.QtObjectPtr(PointerFromQListWidget(parent)), C.int(ty))))
}

func (ptr *QListWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func NewQListWidgetItem4(other QListWidgetItemITF) *QListWidgetItem {
	return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidgetItem_NewQListWidgetItem4(C.QtObjectPtr(PointerFromQListWidgetItem(other)))))
}

func (ptr *QListWidgetItem) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QListWidgetItem_CheckState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QListWidgetItem) Clone() *QListWidgetItem {
	if ptr.Pointer() != nil {
		return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidgetItem_Clone(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QListWidgetItem) Data(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_Data(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QListWidgetItem) Flags() core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QListWidgetItem_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QListWidgetItem) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsHidden(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QListWidgetItem) IsSelected() bool {
	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsSelected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QListWidgetItem) ListWidget() *QListWidget {
	if ptr.Pointer() != nil {
		return QListWidgetFromPointer(unsafe.Pointer(C.QListWidgetItem_ListWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QListWidgetItem) Read(in core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_Read(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(in)))
	}
}

func (ptr *QListWidgetItem) SetBackground(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetBackground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QListWidgetItem) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetCheckState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QListWidgetItem) SetData(role int, value string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetData(C.QtObjectPtr(ptr.Pointer()), C.int(role), C.CString(value))
	}
}

func (ptr *QListWidgetItem) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QListWidgetItem) SetForeground(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetForeground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QListWidgetItem) SetHidden(hide bool) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetHidden(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QListWidgetItem) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QListWidgetItem) SetSelected(sele bool) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSelected(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QListWidgetItem) SetSizeHint(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSizeHint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QListWidgetItem) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetStatusTip(C.QtObjectPtr(ptr.Pointer()), C.CString(statusTip))
	}
}

func (ptr *QListWidgetItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QListWidgetItem) SetTextAlignment(alignment int) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetTextAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QListWidgetItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(toolTip))
	}
}

func (ptr *QListWidgetItem) SetWhatsThis(whatsThis string) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.CString(whatsThis))
	}
}

func (ptr *QListWidgetItem) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_StatusTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QListWidgetItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QListWidgetItem) TextAlignment() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_TextAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QListWidgetItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QListWidgetItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QListWidgetItem) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_WhatsThis(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QListWidgetItem) Write(out core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(out)))
	}
}

func (ptr *QListWidgetItem) DestroyQListWidgetItem() {
	if ptr.Pointer() != nil {
		C.QListWidgetItem_DestroyQListWidgetItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
