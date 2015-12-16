package widgets

//#include "widgets.h"
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

type QTableWidgetItem_ITF interface {
	QTableWidgetItem_PTR() *QTableWidgetItem
}

func (p *QTableWidgetItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTableWidgetItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTableWidgetItem(ptr QTableWidgetItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidgetItem_PTR().Pointer()
	}
	return nil
}

func NewQTableWidgetItemFromPointer(ptr unsafe.Pointer) *QTableWidgetItem {
	var n = new(QTableWidgetItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QTableWidgetItem_") {
		n.SetObjectNameAbs("QTableWidgetItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QTableWidgetItem) QTableWidgetItem_PTR() *QTableWidgetItem {
	return ptr
}

//QTableWidgetItem::ItemType
type QTableWidgetItem__ItemType int64

const (
	QTableWidgetItem__Type     = QTableWidgetItem__ItemType(0)
	QTableWidgetItem__UserType = QTableWidgetItem__ItemType(1000)
)

func (ptr *QTableWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	defer qt.Recovering("QTableWidgetItem::setFlags")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func NewQTableWidgetItem3(icon gui.QIcon_ITF, text string, ty int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidgetItem::QTableWidgetItem")

	return NewQTableWidgetItemFromPointer(C.QTableWidgetItem_NewQTableWidgetItem3(gui.PointerFromQIcon(icon), C.CString(text), C.int(ty)))
}

func NewQTableWidgetItem2(text string, ty int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidgetItem::QTableWidgetItem")

	return NewQTableWidgetItemFromPointer(C.QTableWidgetItem_NewQTableWidgetItem2(C.CString(text), C.int(ty)))
}

func NewQTableWidgetItem4(other QTableWidgetItem_ITF) *QTableWidgetItem {
	defer qt.Recovering("QTableWidgetItem::QTableWidgetItem")

	return NewQTableWidgetItemFromPointer(C.QTableWidgetItem_NewQTableWidgetItem4(PointerFromQTableWidgetItem(other)))
}

func NewQTableWidgetItem(ty int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidgetItem::QTableWidgetItem")

	return NewQTableWidgetItemFromPointer(C.QTableWidgetItem_NewQTableWidgetItem(C.int(ty)))
}

func (ptr *QTableWidgetItem) Background() *gui.QBrush {
	defer qt.Recovering("QTableWidgetItem::background")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QTableWidgetItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidgetItem) CheckState() core.Qt__CheckState {
	defer qt.Recovering("QTableWidgetItem::checkState")

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QTableWidgetItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetItem) Clone() *QTableWidgetItem {
	defer qt.Recovering("QTableWidgetItem::clone")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidgetItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidgetItem) Column() int {
	defer qt.Recovering("QTableWidgetItem::column")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetItem) Data(role int) *core.QVariant {
	defer qt.Recovering("QTableWidgetItem::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTableWidgetItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QTableWidgetItem) Flags() core.Qt__ItemFlag {
	defer qt.Recovering("QTableWidgetItem::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QTableWidgetItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetItem) Foreground() *gui.QBrush {
	defer qt.Recovering("QTableWidgetItem::foreground")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QTableWidgetItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidgetItem) Icon() *gui.QIcon {
	defer qt.Recovering("QTableWidgetItem::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QTableWidgetItem_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidgetItem) IsSelected() bool {
	defer qt.Recovering("QTableWidgetItem::isSelected")

	if ptr.Pointer() != nil {
		return C.QTableWidgetItem_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableWidgetItem) Row() int {
	defer qt.Recovering("QTableWidgetItem::row")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetItem) SetBackground(brush gui.QBrush_ITF) {
	defer qt.Recovering("QTableWidgetItem::setBackground")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetBackground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QTableWidgetItem) SetCheckState(state core.Qt__CheckState) {
	defer qt.Recovering("QTableWidgetItem::setCheckState")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QTableWidgetItem) ConnectSetData(f func(role int, value *core.QVariant)) {
	defer qt.Recovering("connect QTableWidgetItem::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData", f)
	}
}

func (ptr *QTableWidgetItem) DisconnectSetData() {
	defer qt.Recovering("disconnect QTableWidgetItem::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData")
	}
}

//export callbackQTableWidgetItemSetData
func callbackQTableWidgetItemSetData(ptrName *C.char, role C.int, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidgetItem::setData")

	var signal = qt.GetSignal(C.GoString(ptrName), "setData")
	if signal != nil {
		defer signal.(func(int, *core.QVariant))(int(role), core.NewQVariantFromPointer(value))
		return true
	}
	return false

}

func (ptr *QTableWidgetItem) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QTableWidgetItem::setFont")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QTableWidgetItem) SetForeground(brush gui.QBrush_ITF) {
	defer qt.Recovering("QTableWidgetItem::setForeground")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetForeground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QTableWidgetItem) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QTableWidgetItem::setIcon")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTableWidgetItem) SetSelected(sele bool) {
	defer qt.Recovering("QTableWidgetItem::setSelected")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetSelected(ptr.Pointer(), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTableWidgetItem) SetSizeHint(size core.QSize_ITF) {
	defer qt.Recovering("QTableWidgetItem::setSizeHint")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QTableWidgetItem) SetStatusTip(statusTip string) {
	defer qt.Recovering("QTableWidgetItem::setStatusTip")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QTableWidgetItem) SetText(text string) {
	defer qt.Recovering("QTableWidgetItem::setText")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTableWidgetItem) SetTextAlignment(alignment int) {
	defer qt.Recovering("QTableWidgetItem::setTextAlignment")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QTableWidgetItem) SetToolTip(toolTip string) {
	defer qt.Recovering("QTableWidgetItem::setToolTip")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QTableWidgetItem) SetWhatsThis(whatsThis string) {
	defer qt.Recovering("QTableWidgetItem::setWhatsThis")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QTableWidgetItem) SizeHint() *core.QSize {
	defer qt.Recovering("QTableWidgetItem::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTableWidgetItem_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidgetItem) StatusTip() string {
	defer qt.Recovering("QTableWidgetItem::statusTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTableWidgetItem) TableWidget() *QTableWidget {
	defer qt.Recovering("QTableWidgetItem::tableWidget")

	if ptr.Pointer() != nil {
		return NewQTableWidgetFromPointer(C.QTableWidgetItem_TableWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidgetItem) Text() string {
	defer qt.Recovering("QTableWidgetItem::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTableWidgetItem) TextAlignment() int {
	defer qt.Recovering("QTableWidgetItem::textAlignment")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetItem) ToolTip() string {
	defer qt.Recovering("QTableWidgetItem::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTableWidgetItem) Type() int {
	defer qt.Recovering("QTableWidgetItem::type")

	if ptr.Pointer() != nil {
		return int(C.QTableWidgetItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidgetItem) WhatsThis() string {
	defer qt.Recovering("QTableWidgetItem::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTableWidgetItem) DestroyQTableWidgetItem() {
	defer qt.Recovering("QTableWidgetItem::~QTableWidgetItem")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_DestroyQTableWidgetItem(ptr.Pointer())
	}
}

func (ptr *QTableWidgetItem) ObjectNameAbs() string {
	defer qt.Recovering("QTableWidgetItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTableWidgetItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTableWidgetItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QTableWidgetItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QTableWidgetItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
