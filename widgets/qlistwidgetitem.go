package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::QListWidgetItem")
		}
	}()

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem(PointerFromQListWidget(parent), C.int(ty)))
}

func NewQListWidgetItem3(icon gui.QIcon_ITF, text string, parent QListWidget_ITF, ty int) *QListWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::QListWidgetItem")
		}
	}()

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem3(gui.PointerFromQIcon(icon), C.CString(text), PointerFromQListWidget(parent), C.int(ty)))
}

func NewQListWidgetItem2(text string, parent QListWidget_ITF, ty int) *QListWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::QListWidgetItem")
		}
	}()

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem2(C.CString(text), PointerFromQListWidget(parent), C.int(ty)))
}

func (ptr *QListWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func NewQListWidgetItem4(other QListWidgetItem_ITF) *QListWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::QListWidgetItem")
		}
	}()

	return NewQListWidgetItemFromPointer(C.QListWidgetItem_NewQListWidgetItem4(PointerFromQListWidgetItem(other)))
}

func (ptr *QListWidgetItem) Background() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::background")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QListWidgetItem_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) CheckState() core.Qt__CheckState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::checkState")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QListWidgetItem_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) Clone() *QListWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::clone")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidgetItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) Data(role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QListWidgetItem_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QListWidgetItem) Flags() core.Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QListWidgetItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) Foreground() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::foreground")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QListWidgetItem_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) IsHidden() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::isHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidgetItem) IsSelected() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::isSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListWidgetItem_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidgetItem) ListWidget() *QListWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::listWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQListWidgetFromPointer(C.QListWidgetItem_ListWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidgetItem) Read(in core.QDataStream_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::read")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_Read(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QListWidgetItem) SetBackground(brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetBackground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QListWidgetItem) SetCheckState(state core.Qt__CheckState) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setCheckState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QListWidgetItem) SetData(role int, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetData(ptr.Pointer(), C.int(role), core.PointerFromQVariant(value))
	}
}

func (ptr *QListWidgetItem) SetFont(font gui.QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QListWidgetItem) SetForeground(brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setForeground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetForeground(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QListWidgetItem) SetHidden(hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QListWidgetItem) SetIcon(icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QListWidgetItem) SetSelected(sele bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSelected(ptr.Pointer(), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QListWidgetItem) SetSizeHint(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setSizeHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetSizeHint(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QListWidgetItem) SetStatusTip(statusTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setStatusTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QListWidgetItem) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QListWidgetItem) SetTextAlignment(alignment int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setTextAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetTextAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QListWidgetItem) SetToolTip(toolTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetToolTip(ptr.Pointer(), C.CString(toolTip))
	}
}

func (ptr *QListWidgetItem) SetWhatsThis(whatsThis string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::setWhatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_SetWhatsThis(ptr.Pointer(), C.CString(whatsThis))
	}
}

func (ptr *QListWidgetItem) StatusTip() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::statusTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) TextAlignment() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::textAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_TextAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) ToolTip() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::toolTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QListWidgetItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidgetItem) WhatsThis() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::whatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QListWidgetItem_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QListWidgetItem) Write(out core.QDataStream_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::write")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_Write(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

func (ptr *QListWidgetItem) DestroyQListWidgetItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListWidgetItem::~QListWidgetItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListWidgetItem_DestroyQListWidgetItem(ptr.Pointer())
	}
}
