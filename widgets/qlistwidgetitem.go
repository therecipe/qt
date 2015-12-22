package widgets

//#include "widgets.h"
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
	for len(n.ObjectNameAbs()) < len("QListWidgetItem_") {
		n.SetObjectNameAbs("QListWidgetItem_" + qt.Identifier())
	}
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
	defer qt.Recovering("QListWidgetItem::QListWidgetItem")

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem(PointerFromQListWidget(parent), C.int(ty)))
}

func NewQListWidgetItem3(icon gui.QIcon_ITF, text string, parent QListWidget_ITF, ty int) *QListWidgetItem {
	defer qt.Recovering("QListWidgetItem::QListWidgetItem")

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem3(gui.PointerFromQIcon(icon), C.CString(text), PointerFromQListWidget(parent), C.int(ty)))
}

func NewQListWidgetItem2(text string, parent QListWidget_ITF, ty int) *QListWidgetItem {
	defer qt.Recovering("QListWidgetItem::QListWidgetItem")

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem2(C.CString(text), PointerFromQListWidget(parent), C.int(ty)))
}

func (ptr *QListWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	defer qt.Recovering("QListWidgetItem::setFlags")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func NewQListWidgetItem4(other QListWidgetItem_ITF) *QListWidgetItem {
	defer qt.Recovering("QListWidgetItem::QListWidgetItem")

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem4(PointerFromQListWidgetItem(other)))
}

func (ptr *QListWidgetItem) Background() *gui.QBrush {
	defer qt.Recovering("QListWidgetItem::background")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QListWidgetItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) CheckState() core.Qt__CheckState {
	defer qt.Recovering("QListWidgetItem::checkState")

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QListWidgetItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) Clone() *QListWidgetItem {
	defer qt.Recovering("QListWidgetItem::clone")

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidgetItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) Data(role int) *core.QVariant {
	defer qt.Recovering("QListWidgetItem::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QListWidgetItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QListWidgetItem) Flags() core.Qt__ItemFlag {
	defer qt.Recovering("QListWidgetItem::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QListWidgetItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) Foreground() *gui.QBrush {
	defer qt.Recovering("QListWidgetItem::foreground")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QListWidgetItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) Icon() *gui.QIcon {
	defer qt.Recovering("QListWidgetItem::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QListWidgetItem_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) IsHidden() bool {
	defer qt.Recovering("QListWidgetItem::isHidden")

	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidgetItem) IsSelected() bool {
	defer qt.Recovering("QListWidgetItem::isSelected")

	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidgetItem) ListWidget() *QListWidget {
	defer qt.Recovering("QListWidgetItem::listWidget")

	if ptr.Pointer() != nil {
		return NewQListWidgetFromPointer(C.QListWidgetItem_ListWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) SetBackground(brush gui.QBrush_ITF) {
	defer qt.Recovering("QListWidgetItem::setBackground")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetBackground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QListWidgetItem) SetCheckState(state core.Qt__CheckState) {
	defer qt.Recovering("QListWidgetItem::setCheckState")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QListWidgetItem) ConnectSetData(f func(role int, value *core.QVariant)) {
	defer qt.Recovering("connect QListWidgetItem::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData", f)
	}
}

func (ptr *QListWidgetItem) DisconnectSetData() {
	defer qt.Recovering("disconnect QListWidgetItem::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData")
	}
}

//export callbackQListWidgetItemSetData
func callbackQListWidgetItemSetData(ptrName *C.char, role C.int, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QListWidgetItem::setData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setData"); signal != nil {
		signal.(func(int, *core.QVariant))(int(role), core.NewQVariantFromPointer(value))
		return true
	}
	return false

}

func (ptr *QListWidgetItem) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QListWidgetItem::setFont")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QListWidgetItem) SetForeground(brush gui.QBrush_ITF) {
	defer qt.Recovering("QListWidgetItem::setForeground")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetForeground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QListWidgetItem) SetHidden(hide bool) {
	defer qt.Recovering("QListWidgetItem::setHidden")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QListWidgetItem) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QListWidgetItem::setIcon")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QListWidgetItem) SetSelected(sele bool) {
	defer qt.Recovering("QListWidgetItem::setSelected")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSelected(ptr.Pointer(), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QListWidgetItem) SetSizeHint(size core.QSize_ITF) {
	defer qt.Recovering("QListWidgetItem::setSizeHint")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QListWidgetItem) SetStatusTip(statusTip string) {
	defer qt.Recovering("QListWidgetItem::setStatusTip")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QListWidgetItem) SetText(text string) {
	defer qt.Recovering("QListWidgetItem::setText")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QListWidgetItem) SetTextAlignment(alignment int) {
	defer qt.Recovering("QListWidgetItem::setTextAlignment")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QListWidgetItem) SetToolTip(toolTip string) {
	defer qt.Recovering("QListWidgetItem::setToolTip")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QListWidgetItem) SetWhatsThis(whatsThis string) {
	defer qt.Recovering("QListWidgetItem::setWhatsThis")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QListWidgetItem) SizeHint() *core.QSize {
	defer qt.Recovering("QListWidgetItem::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QListWidgetItem_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) StatusTip() string {
	defer qt.Recovering("QListWidgetItem::statusTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Text() string {
	defer qt.Recovering("QListWidgetItem::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) TextAlignment() int {
	defer qt.Recovering("QListWidgetItem::textAlignment")

	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) ToolTip() string {
	defer qt.Recovering("QListWidgetItem::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Type() int {
	defer qt.Recovering("QListWidgetItem::type")

	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) WhatsThis() string {
	defer qt.Recovering("QListWidgetItem::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) DestroyQListWidgetItem() {
	defer qt.Recovering("QListWidgetItem::~QListWidgetItem")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_DestroyQListWidgetItem(ptr.Pointer())
	}
}

func (ptr *QListWidgetItem) ObjectNameAbs() string {
	defer qt.Recovering("QListWidgetItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QListWidgetItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
