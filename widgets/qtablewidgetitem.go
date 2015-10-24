package widgets

//#include "qtablewidgetitem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QTableWidgetItem struct {
	ptr unsafe.Pointer
}

type QTableWidgetItemITF interface {
	QTableWidgetItemPTR() *QTableWidgetItem
}

func (p *QTableWidgetItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTableWidgetItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTableWidgetItem(ptr QTableWidgetItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidgetItemPTR().Pointer()
	}
	return nil
}

func QTableWidgetItemFromPointer(ptr unsafe.Pointer) *QTableWidgetItem {
	var n = new(QTableWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTableWidgetItem) QTableWidgetItemPTR() *QTableWidgetItem {
	return ptr
}

//QTableWidgetItem::ItemType
type QTableWidgetItem__ItemType int

var (
	QTableWidgetItem__Type     = QTableWidgetItem__ItemType(0)
	QTableWidgetItem__UserType = QTableWidgetItem__ItemType(1000)
)

func (ptr *QTableWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func NewQTableWidgetItem3(icon gui.QIconITF, text string, ty int) *QTableWidgetItem {
	return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidgetItem_NewQTableWidgetItem3(C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.int(ty))))
}

func NewQTableWidgetItem2(text string, ty int) *QTableWidgetItem {
	return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidgetItem_NewQTableWidgetItem2(C.CString(text), C.int(ty))))
}

func NewQTableWidgetItem4(other QTableWidgetItemITF) *QTableWidgetItem {
	return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidgetItem_NewQTableWidgetItem4(C.QtObjectPtr(PointerFromQTableWidgetItem(other)))))
}

func NewQTableWidgetItem(ty int) *QTableWidgetItem {
	return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidgetItem_NewQTableWidgetItem(C.int(ty))))
}

func (ptr *QTableWidgetItem) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QTableWidgetItem_CheckState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetItem) Clone() *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidgetItem_Clone(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTableWidgetItem) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_Column(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetItem) Data(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_Data(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QTableWidgetItem) Flags() core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QTableWidgetItem_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetItem) IsSelected() bool {
	if ptr.Pointer() != nil {
		return C.QTableWidgetItem_IsSelected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTableWidgetItem) Read(in core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_Read(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(in)))
	}
}

func (ptr *QTableWidgetItem) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_Row(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetItem) SetBackground(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetBackground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QTableWidgetItem) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetCheckState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QTableWidgetItem) SetData(role int, value string) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetData(C.QtObjectPtr(ptr.Pointer()), C.int(role), C.CString(value))
	}
}

func (ptr *QTableWidgetItem) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QTableWidgetItem) SetForeground(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetForeground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QTableWidgetItem) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QTableWidgetItem) SetSelected(sele bool) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetSelected(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTableWidgetItem) SetSizeHint(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetSizeHint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QTableWidgetItem) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetStatusTip(C.QtObjectPtr(ptr.Pointer()), C.CString(statusTip))
	}
}

func (ptr *QTableWidgetItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTableWidgetItem) SetTextAlignment(alignment int) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetTextAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QTableWidgetItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(toolTip))
	}
}

func (ptr *QTableWidgetItem) SetWhatsThis(whatsThis string) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.CString(whatsThis))
	}
}

func (ptr *QTableWidgetItem) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_StatusTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTableWidgetItem) TableWidget() *QTableWidget {
	if ptr.Pointer() != nil {
		return QTableWidgetFromPointer(unsafe.Pointer(C.QTableWidgetItem_TableWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTableWidgetItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTableWidgetItem) TextAlignment() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_TextAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTableWidgetItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidgetItem) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_WhatsThis(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTableWidgetItem) Write(out core.QDataStreamITF) {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(out)))
	}
}

func (ptr *QTableWidgetItem) DestroyQTableWidgetItem() {
	if ptr.Pointer() != nil {
		C.QTableWidgetItem_DestroyQTableWidgetItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
