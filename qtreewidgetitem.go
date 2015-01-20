package qt

//#include "qtreewidgetitem.h"
import "C"

type qtreewidgetitem struct {
	ptr C.QtObjectPtr
}

type QTreeWidgetItem interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Destroy()
	AddChild(child QTreeWidgetItem)
	CheckState(column int) CheckState
	Child(index int) QTreeWidgetItem
	ChildCount() int
	Clone() QTreeWidgetItem
	ColumnCount() int
	Flags() ItemFlag
	IndexOfChild(child QTreeWidgetItem) int
	InsertChild(index int, child QTreeWidgetItem)
	IsDisabled() bool
	IsExpanded() bool
	IsFirstColumnSpanned() bool
	IsHidden() bool
	IsSelected() bool
	Parent() QTreeWidgetItem
	RemoveChild(child QTreeWidgetItem)
	SetBackground(column int, brush QBrush)
	SetCheckState(column int, state CheckState)
	SetDisabled(disabled bool)
	SetExpanded(expand bool)
	SetFirstColumnSpanned(span bool)
	SetFlags(flags ItemFlag)
	SetForeground(column int, brush QBrush)
	SetHidden(hide bool)
	SetSelected(selected bool)
	SetStatusTip(column int, statusTip string)
	SetText(column int, text string)
	SetTextAlignment(column int, alignment int)
	SetToolTip(column int, toolTip string)
	SetWhatsThis(column int, whatsThis string)
	SortChildren(column int, order SortOrder)
	StatusTip(column int) string
	TakeChild(index int) QTreeWidgetItem
	Text(column int) string
	TextAlignment(column int) int
	ToolTip(column int) string
	TreeWidget() QTreeWidget
	Type() int
	WhatsThis(column int) string
}

func (p *qtreewidgetitem) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtreewidgetitem) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTreeWidgetItem1(typ int) QTreeWidgetItem {
	var qtreewidgetitem = new(qtreewidgetitem)
	qtreewidgetitem.SetPointer(C.QTreeWidgetItem_New_Int(C.int(typ)))
	return qtreewidgetitem
}

func NewQTreeWidgetItem2(parent QTreeWidget, typ int) QTreeWidgetItem {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtreewidgetitem = new(qtreewidgetitem)
	qtreewidgetitem.SetPointer(C.QTreeWidgetItem_New_QTreeWidget_Int(parentPtr, C.int(typ)))
	return qtreewidgetitem
}

func NewQTreeWidgetItem3(parent QTreeWidget, preceding QTreeWidgetItem, typ int) QTreeWidgetItem {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var precedingPtr C.QtObjectPtr
	if preceding != nil {
		precedingPtr = preceding.Pointer()
	}
	var qtreewidgetitem = new(qtreewidgetitem)
	qtreewidgetitem.SetPointer(C.QTreeWidgetItem_New_QTreeWidget_QTreeWidgetItem_Int(parentPtr, precedingPtr, C.int(typ)))
	return qtreewidgetitem
}

func NewQTreeWidgetItem4(parent QTreeWidgetItem, typ int) QTreeWidgetItem {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtreewidgetitem = new(qtreewidgetitem)
	qtreewidgetitem.SetPointer(C.QTreeWidgetItem_New_QTreeWidgetItem_Int(parentPtr, C.int(typ)))
	return qtreewidgetitem
}

func NewQTreeWidgetItem5(parent QTreeWidgetItem, preceding QTreeWidgetItem, typ int) QTreeWidgetItem {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var precedingPtr C.QtObjectPtr
	if preceding != nil {
		precedingPtr = preceding.Pointer()
	}
	var qtreewidgetitem = new(qtreewidgetitem)
	qtreewidgetitem.SetPointer(C.QTreeWidgetItem_New_QTreeWidgetItem_QTreeWidgetItem_Int(parentPtr, precedingPtr, C.int(typ)))
	return qtreewidgetitem
}

func NewQTreeWidgetItem6(other QTreeWidgetItem) QTreeWidgetItem {
	var otherPtr C.QtObjectPtr
	if other != nil {
		otherPtr = other.Pointer()
	}
	var qtreewidgetitem = new(qtreewidgetitem)
	qtreewidgetitem.SetPointer(C.QTreeWidgetItem_New_QTreeWidgetItem(otherPtr))
	return qtreewidgetitem
}

func (p *qtreewidgetitem) Destroy() {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtreewidgetitem) AddChild(child QTreeWidgetItem) {
	if p.Pointer() != nil {
		var childPtr C.QtObjectPtr
		if child != nil {
			childPtr = child.Pointer()
		}
		C.QTreeWidgetItem_AddChild_QTreeWidgetItem(p.Pointer(), childPtr)
	}
}

func (p *qtreewidgetitem) CheckState(column int) CheckState {
	if p.Pointer() == nil {
		return 0
	}
	return CheckState(C.QTreeWidgetItem_CheckState_Int(p.Pointer(), C.int(column)))
}

func (p *qtreewidgetitem) Child(index int) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidgetItem_Child_Int(p.Pointer(), C.int(index)))
		return qtreewidgetitem
	}
}

func (p *qtreewidgetitem) ChildCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTreeWidgetItem_ChildCount(p.Pointer()))
}

func (p *qtreewidgetitem) Clone() QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidgetItem_Clone(p.Pointer()))
		return qtreewidgetitem
	}
}

func (p *qtreewidgetitem) ColumnCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTreeWidgetItem_ColumnCount(p.Pointer()))
}

func (p *qtreewidgetitem) Flags() ItemFlag {
	if p.Pointer() == nil {
		return 0
	}
	return ItemFlag(C.QTreeWidgetItem_Flags(p.Pointer()))
}

func (p *qtreewidgetitem) IndexOfChild(child QTreeWidgetItem) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var childPtr C.QtObjectPtr
		if child != nil {
			childPtr = child.Pointer()
		}
		return int(C.QTreeWidgetItem_IndexOfChild_QTreeWidgetItem(p.Pointer(), childPtr))
	}
}

func (p *qtreewidgetitem) InsertChild(index int, child QTreeWidgetItem) {
	if p.Pointer() != nil {
		var childPtr C.QtObjectPtr
		if child != nil {
			childPtr = child.Pointer()
		}
		C.QTreeWidgetItem_InsertChild_Int_QTreeWidgetItem(p.Pointer(), C.int(index), childPtr)
	}
}

func (p *qtreewidgetitem) IsDisabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTreeWidgetItem_IsDisabled(p.Pointer()) != 0
}

func (p *qtreewidgetitem) IsExpanded() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTreeWidgetItem_IsExpanded(p.Pointer()) != 0
}

func (p *qtreewidgetitem) IsFirstColumnSpanned() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTreeWidgetItem_IsFirstColumnSpanned(p.Pointer()) != 0
}

func (p *qtreewidgetitem) IsHidden() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTreeWidgetItem_IsHidden(p.Pointer()) != 0
}

func (p *qtreewidgetitem) IsSelected() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTreeWidgetItem_IsSelected(p.Pointer()) != 0
}

func (p *qtreewidgetitem) Parent() QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidgetItem_Parent(p.Pointer()))
		return qtreewidgetitem
	}
}

func (p *qtreewidgetitem) RemoveChild(child QTreeWidgetItem) {
	if p.Pointer() != nil {
		var childPtr C.QtObjectPtr
		if child != nil {
			childPtr = child.Pointer()
		}
		C.QTreeWidgetItem_RemoveChild_QTreeWidgetItem(p.Pointer(), childPtr)
	}
}

func (p *qtreewidgetitem) SetBackground(column int, brush QBrush) {
	if p.Pointer() != nil {
		var brushPtr C.QtObjectPtr
		if brush != nil {
			brushPtr = brush.Pointer()
		}
		C.QTreeWidgetItem_SetBackground_Int_QBrush(p.Pointer(), C.int(column), brushPtr)
	}
}

func (p *qtreewidgetitem) SetCheckState(column int, state CheckState) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetCheckState_Int_CheckState(p.Pointer(), C.int(column), C.int(state))
	}
}

func (p *qtreewidgetitem) SetDisabled(disabled bool) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetDisabled_Bool(p.Pointer(), goBoolToCInt(disabled))
	}
}

func (p *qtreewidgetitem) SetExpanded(expand bool) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetExpanded_Bool(p.Pointer(), goBoolToCInt(expand))
	}
}

func (p *qtreewidgetitem) SetFirstColumnSpanned(span bool) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetFirstColumnSpanned_Bool(p.Pointer(), goBoolToCInt(span))
	}
}

func (p *qtreewidgetitem) SetFlags(flags ItemFlag) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetFlags_ItemFlag(p.Pointer(), C.int(flags))
	}
}

func (p *qtreewidgetitem) SetForeground(column int, brush QBrush) {
	if p.Pointer() != nil {
		var brushPtr C.QtObjectPtr
		if brush != nil {
			brushPtr = brush.Pointer()
		}
		C.QTreeWidgetItem_SetForeground_Int_QBrush(p.Pointer(), C.int(column), brushPtr)
	}
}

func (p *qtreewidgetitem) SetHidden(hide bool) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetHidden_Bool(p.Pointer(), goBoolToCInt(hide))
	}
}

func (p *qtreewidgetitem) SetSelected(selected bool) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetSelected_Bool(p.Pointer(), goBoolToCInt(selected))
	}
}

func (p *qtreewidgetitem) SetStatusTip(column int, statusTip string) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetStatusTip_Int_String(p.Pointer(), C.int(column), C.CString(statusTip))
	}
}

func (p *qtreewidgetitem) SetText(column int, text string) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetText_Int_String(p.Pointer(), C.int(column), C.CString(text))
	}
}

func (p *qtreewidgetitem) SetTextAlignment(column int, alignment int) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetTextAlignment_Int_Int(p.Pointer(), C.int(column), C.int(alignment))
	}
}

func (p *qtreewidgetitem) SetToolTip(column int, toolTip string) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetToolTip_Int_String(p.Pointer(), C.int(column), C.CString(toolTip))
	}
}

func (p *qtreewidgetitem) SetWhatsThis(column int, whatsThis string) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SetWhatsThis_Int_String(p.Pointer(), C.int(column), C.CString(whatsThis))
	}
}

func (p *qtreewidgetitem) SortChildren(column int, order SortOrder) {
	if p.Pointer() != nil {
		C.QTreeWidgetItem_SortChildren_Int_SortOrder(p.Pointer(), C.int(column), C.int(order))
	}
}

func (p *qtreewidgetitem) StatusTip(column int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTreeWidgetItem_StatusTip_Int(p.Pointer(), C.int(column)))
}

func (p *qtreewidgetitem) TakeChild(index int) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidgetItem_TakeChild_Int(p.Pointer(), C.int(index)))
		return qtreewidgetitem
	}
}

func (p *qtreewidgetitem) Text(column int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTreeWidgetItem_Text_Int(p.Pointer(), C.int(column)))
}

func (p *qtreewidgetitem) TextAlignment(column int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTreeWidgetItem_TextAlignment_Int(p.Pointer(), C.int(column)))
}

func (p *qtreewidgetitem) ToolTip(column int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTreeWidgetItem_ToolTip_Int(p.Pointer(), C.int(column)))
}

func (p *qtreewidgetitem) TreeWidget() QTreeWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidget = new(qtreewidget)
		qtreewidget.SetPointer(C.QTreeWidgetItem_TreeWidget(p.Pointer()))
		if qtreewidget.ObjectName() == "" {
			qtreewidget.SetObjectName("QTreeWidget_" + randomIdentifier())
		}
		return qtreewidget
	}
}

func (p *qtreewidgetitem) Type() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTreeWidgetItem_Type(p.Pointer()))
}

func (p *qtreewidgetitem) WhatsThis(column int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTreeWidgetItem_WhatsThis_Int(p.Pointer(), C.int(column)))
}
