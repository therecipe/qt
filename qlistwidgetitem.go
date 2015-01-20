package qt

//#include "qlistwidgetitem.h"
import "C"

type qlistwidgetitem struct {
	ptr C.QtObjectPtr
}

type QListWidgetItem interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Destroy()
	CheckState() CheckState
	Clone() QListWidgetItem
	Flags() ItemFlag
	IsHidden() bool
	IsSelected() bool
	ListWidget() QListWidget
	SetCheckState(state CheckState)
	SetFlags(flags ItemFlag)
	SetHidden(hide bool)
	SetSelected(selected bool)
	SetStatusTip(statusTip string)
	SetText(text string)
	SetTextAlignment(alignment int)
	SetToolTip(toolTip string)
	SetWhatsThis(whatsThis string)
	StatusTip() string
	Text() string
	TextAlignment() int
	ToolTip() string
	Type() int
	WhatsThis() string
}

func (p *qlistwidgetitem) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlistwidgetitem) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQListWidgetItem1(parent QListWidget, typ int) QListWidgetItem {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlistwidgetitem = new(qlistwidgetitem)
	qlistwidgetitem.SetPointer(C.QListWidgetItem_New_QListWidget_Int(parentPtr, C.int(typ)))
	return qlistwidgetitem
}

func NewQListWidgetItem2(text string, parent QListWidget, typ int) QListWidgetItem {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlistwidgetitem = new(qlistwidgetitem)
	qlistwidgetitem.SetPointer(C.QListWidgetItem_New_String_QListWidget_Int(C.CString(text), parentPtr, C.int(typ)))
	return qlistwidgetitem
}

func (p *qlistwidgetitem) Destroy() {
	if p.Pointer() != nil {
		C.QListWidgetItem_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qlistwidgetitem) CheckState() CheckState {
	if p.Pointer() == nil {
		return 0
	}
	return CheckState(C.QListWidgetItem_CheckState(p.Pointer()))
}

func (p *qlistwidgetitem) Clone() QListWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlistwidgetitem = new(qlistwidgetitem)
		qlistwidgetitem.SetPointer(C.QListWidgetItem_Clone(p.Pointer()))
		return qlistwidgetitem
	}
}

func (p *qlistwidgetitem) Flags() ItemFlag {
	if p.Pointer() == nil {
		return 0
	}
	return ItemFlag(C.QListWidgetItem_Flags(p.Pointer()))
}

func (p *qlistwidgetitem) IsHidden() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListWidgetItem_IsHidden(p.Pointer()) != 0
}

func (p *qlistwidgetitem) IsSelected() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListWidgetItem_IsSelected(p.Pointer()) != 0
}

func (p *qlistwidgetitem) ListWidget() QListWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlistwidget = new(qlistwidget)
		qlistwidget.SetPointer(C.QListWidgetItem_ListWidget(p.Pointer()))
		if qlistwidget.ObjectName() == "" {
			qlistwidget.SetObjectName("QListWidget_" + randomIdentifier())
		}
		return qlistwidget
	}
}

func (p *qlistwidgetitem) SetCheckState(state CheckState) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetCheckState_CheckState(p.Pointer(), C.int(state))
	}
}

func (p *qlistwidgetitem) SetFlags(flags ItemFlag) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetFlags_ItemFlag(p.Pointer(), C.int(flags))
	}
}

func (p *qlistwidgetitem) SetHidden(hide bool) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetHidden_Bool(p.Pointer(), goBoolToCInt(hide))
	}
}

func (p *qlistwidgetitem) SetSelected(selected bool) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetSelected_Bool(p.Pointer(), goBoolToCInt(selected))
	}
}

func (p *qlistwidgetitem) SetStatusTip(statusTip string) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetStatusTip_String(p.Pointer(), C.CString(statusTip))
	}
}

func (p *qlistwidgetitem) SetText(text string) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetText_String(p.Pointer(), C.CString(text))
	}
}

func (p *qlistwidgetitem) SetTextAlignment(alignment int) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetTextAlignment_Int(p.Pointer(), C.int(alignment))
	}
}

func (p *qlistwidgetitem) SetToolTip(toolTip string) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetToolTip_String(p.Pointer(), C.CString(toolTip))
	}
}

func (p *qlistwidgetitem) SetWhatsThis(whatsThis string) {
	if p.Pointer() != nil {
		C.QListWidgetItem_SetWhatsThis_String(p.Pointer(), C.CString(whatsThis))
	}
}

func (p *qlistwidgetitem) StatusTip() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QListWidgetItem_StatusTip(p.Pointer()))
}

func (p *qlistwidgetitem) Text() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QListWidgetItem_Text(p.Pointer()))
}

func (p *qlistwidgetitem) TextAlignment() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListWidgetItem_TextAlignment(p.Pointer()))
}

func (p *qlistwidgetitem) ToolTip() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QListWidgetItem_ToolTip(p.Pointer()))
}

func (p *qlistwidgetitem) Type() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListWidgetItem_Type(p.Pointer()))
}

func (p *qlistwidgetitem) WhatsThis() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QListWidgetItem_WhatsThis(p.Pointer()))
}
