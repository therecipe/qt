package qt

//#include "qtreewidget.h"
import "C"

import "strings"

type qtreewidget struct {
	qtreeview
}

type QTreeWidget interface {
	QTreeView
	AddTopLevelItem_QTreeWidgetItem(item QTreeWidgetItem)
	ClosePersistentEditor_QTreeWidgetItem_Int(item QTreeWidgetItem, column int)
	ColumnCount() int
	CurrentColumn() int
	CurrentItem() QTreeWidgetItem
	EditItem_QTreeWidgetItem_Int(item QTreeWidgetItem, column int)
	HeaderItem() QTreeWidgetItem
	IndexOfTopLevelItem_QTreeWidgetItem(item QTreeWidgetItem) int
	InsertTopLevelItem_Int_QTreeWidgetItem(index int, item QTreeWidgetItem)
	InvisibleRootItem() QTreeWidgetItem
	IsFirstItemColumnSpanned_QTreeWidgetItem(item QTreeWidgetItem) bool
	ItemAbove_QTreeWidgetItem(item QTreeWidgetItem) QTreeWidgetItem
	ItemAt_Int_Int(x int, y int) QTreeWidgetItem
	ItemBelow_QTreeWidgetItem(item QTreeWidgetItem) QTreeWidgetItem
	ItemWidget_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) QWidget
	OpenPersistentEditor_QTreeWidgetItem_Int(item QTreeWidgetItem, column int)
	RemoveItemWidget_QTreeWidgetItem_Int(item QTreeWidgetItem, column int)
	SetColumnCount_Int(columns int)
	SetCurrentItem_QTreeWidgetItem(item QTreeWidgetItem)
	SetCurrentItem_QTreeWidgetItem_Int(item QTreeWidgetItem, column int)
	SetFirstItemColumnSpanned_QTreeWidgetItem_Bool(item QTreeWidgetItem, span bool)
	SetHeaderItem_QTreeWidgetItem(item QTreeWidgetItem)
	SetHeaderLabel_String(label string)
	SetHeaderLabels_QStringList(labels []string)
	SetItemWidget_QTreeWidgetItem_Int_QWidget(item QTreeWidgetItem, column int, widget QWidget)
	SortColumn() int
	SortItems_Int_SortOrder(column int, order SortOrder)
	TakeTopLevelItem_Int(index int) QTreeWidgetItem
	TopLevelItem_Int(index int) QTreeWidgetItem
	TopLevelItemCount() int
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSignalCurrentItemChanged(f func())
	DisconnectSignalCurrentItemChanged()
	SignalCurrentItemChanged() func()
	ConnectSignalItemActivated(f func())
	DisconnectSignalItemActivated()
	SignalItemActivated() func()
	ConnectSignalItemChanged(f func())
	DisconnectSignalItemChanged()
	SignalItemChanged() func()
	ConnectSignalItemClicked(f func())
	DisconnectSignalItemClicked()
	SignalItemClicked() func()
	ConnectSignalItemCollapsed(f func())
	DisconnectSignalItemCollapsed()
	SignalItemCollapsed() func()
	ConnectSignalItemDoubleClicked(f func())
	DisconnectSignalItemDoubleClicked()
	SignalItemDoubleClicked() func()
	ConnectSignalItemEntered(f func())
	DisconnectSignalItemEntered()
	SignalItemEntered() func()
	ConnectSignalItemExpanded(f func())
	DisconnectSignalItemExpanded()
	SignalItemExpanded() func()
	ConnectSignalItemPressed(f func())
	DisconnectSignalItemPressed()
	SignalItemPressed() func()
	ConnectSignalItemSelectionChanged(f func())
	DisconnectSignalItemSelectionChanged()
	SignalItemSelectionChanged() func()
}

func (p *qtreewidget) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtreewidget) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTreeWidget_QWidget(parent QWidget) QTreeWidget {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtreewidget = new(qtreewidget)
	qtreewidget.SetPointer(C.QTreeWidget_New_QWidget(parentPtr))
	qtreewidget.SetObjectName_String("QTreeWidget_" + randomIdentifier())
	return qtreewidget
}

func (p *qtreewidget) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTreeWidget_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtreewidget) AddTopLevelItem_QTreeWidgetItem(item QTreeWidgetItem) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_AddTopLevelItem_QTreeWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qtreewidget) ClosePersistentEditor_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_ClosePersistentEditor_QTreeWidgetItem_Int(p.Pointer(), itemPtr, C.int(column))
	}
}

func (p *qtreewidget) ColumnCount() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeWidget_ColumnCount(p.Pointer()))
	}
}

func (p *qtreewidget) CurrentColumn() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeWidget_CurrentColumn(p.Pointer()))
	}
}

func (p *qtreewidget) CurrentItem() QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_CurrentItem(p.Pointer()))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) EditItem_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_EditItem_QTreeWidgetItem_Int(p.Pointer(), itemPtr, C.int(column))
	}
}

func (p *qtreewidget) HeaderItem() QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_HeaderItem(p.Pointer()))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) IndexOfTopLevelItem_QTreeWidgetItem(item QTreeWidgetItem) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		return int(C.QTreeWidget_IndexOfTopLevelItem_QTreeWidgetItem(p.Pointer(), itemPtr))
	}
}

func (p *qtreewidget) InsertTopLevelItem_Int_QTreeWidgetItem(index int, item QTreeWidgetItem) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_InsertTopLevelItem_Int_QTreeWidgetItem(p.Pointer(), C.int(index), itemPtr)
	}
}

func (p *qtreewidget) InvisibleRootItem() QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_InvisibleRootItem(p.Pointer()))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) IsFirstItemColumnSpanned_QTreeWidgetItem(item QTreeWidgetItem) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		return C.QTreeWidget_IsFirstItemColumnSpanned_QTreeWidgetItem(p.Pointer(), itemPtr) != 0
	}
}

func (p *qtreewidget) ItemAbove_QTreeWidgetItem(item QTreeWidgetItem) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_ItemAbove_QTreeWidgetItem(p.Pointer(), itemPtr))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) ItemAt_Int_Int(x int, y int) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_ItemAt_Int_Int(p.Pointer(), C.int(x), C.int(y)))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) ItemBelow_QTreeWidgetItem(item QTreeWidgetItem) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_ItemBelow_QTreeWidgetItem(p.Pointer(), itemPtr))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) ItemWidget_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTreeWidget_ItemWidget_QTreeWidgetItem_Int(p.Pointer(), itemPtr, C.int(column)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtreewidget) OpenPersistentEditor_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_OpenPersistentEditor_QTreeWidgetItem_Int(p.Pointer(), itemPtr, C.int(column))
	}
}

func (p *qtreewidget) RemoveItemWidget_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_RemoveItemWidget_QTreeWidgetItem_Int(p.Pointer(), itemPtr, C.int(column))
	}
}

func (p *qtreewidget) SetColumnCount_Int(columns int) {
	if p.Pointer() != nil {
		C.QTreeWidget_SetColumnCount_Int(p.Pointer(), C.int(columns))
	}
}

func (p *qtreewidget) SetCurrentItem_QTreeWidgetItem(item QTreeWidgetItem) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_SetCurrentItem_QTreeWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qtreewidget) SetCurrentItem_QTreeWidgetItem_Int(item QTreeWidgetItem, column int) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_SetCurrentItem_QTreeWidgetItem_Int(p.Pointer(), itemPtr, C.int(column))
	}
}

func (p *qtreewidget) SetFirstItemColumnSpanned_QTreeWidgetItem_Bool(item QTreeWidgetItem, span bool) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_SetFirstItemColumnSpanned_QTreeWidgetItem_Bool(p.Pointer(), itemPtr, goBoolToCInt(span))
	}
}

func (p *qtreewidget) SetHeaderItem_QTreeWidgetItem(item QTreeWidgetItem) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QTreeWidget_SetHeaderItem_QTreeWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qtreewidget) SetHeaderLabel_String(label string) {
	if p.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabel_String(p.Pointer(), C.CString(label))
	}
}

func (p *qtreewidget) SetHeaderLabels_QStringList(labels []string) {
	if p.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabels_QStringList(p.Pointer(), C.CString(strings.Join(labels, "|")))
	}
}

func (p *qtreewidget) SetItemWidget_QTreeWidgetItem_Int_QWidget(item QTreeWidgetItem, column int, widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QTreeWidget_SetItemWidget_QTreeWidgetItem_Int_QWidget(p.Pointer(), itemPtr, C.int(column), widgetPtr)
	}
}

func (p *qtreewidget) SortColumn() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeWidget_SortColumn(p.Pointer()))
	}
}

func (p *qtreewidget) SortItems_Int_SortOrder(column int, order SortOrder) {
	if p.Pointer() != nil {
		C.QTreeWidget_SortItems_Int_SortOrder(p.Pointer(), C.int(column), C.int(order))
	}
}

func (p *qtreewidget) TakeTopLevelItem_Int(index int) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_TakeTopLevelItem_Int(p.Pointer(), C.int(index)))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) TopLevelItem_Int(index int) QTreeWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtreewidgetitem = new(qtreewidgetitem)
		qtreewidgetitem.SetPointer(C.QTreeWidget_TopLevelItem_Int(p.Pointer(), C.int(index)))
		return qtreewidgetitem
	}
}

func (p *qtreewidget) TopLevelItemCount() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeWidget_TopLevelItemCount(p.Pointer()))
	}
}

func (p *qtreewidget) ConnectSlotClear() {
	C.QTreeWidget_ConnectSlotClear(p.Pointer())
}

func (p *qtreewidget) DisconnectSlotClear() {
	C.QTreeWidget_DisconnectSlotClear(p.Pointer())
}

func (p *qtreewidget) SlotClear() {
	if p.Pointer() != nil {
		C.QTreeWidget_Clear(p.Pointer())
	}
}

func (p *qtreewidget) ConnectSignalCurrentItemChanged(f func()) {
	C.QTreeWidget_ConnectSignalCurrentItemChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentItemChanged", f)
}

func (p *qtreewidget) DisconnectSignalCurrentItemChanged() {
	C.QTreeWidget_DisconnectSignalCurrentItemChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentItemChanged")
}

func (p *qtreewidget) SignalCurrentItemChanged() func() {
	return getSignal(p.ObjectName(), "currentItemChanged")
}

func (p *qtreewidget) ConnectSignalItemActivated(f func()) {
	C.QTreeWidget_ConnectSignalItemActivated(p.Pointer())
	connectSignal(p.ObjectName(), "itemActivated", f)
}

func (p *qtreewidget) DisconnectSignalItemActivated() {
	C.QTreeWidget_DisconnectSignalItemActivated(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemActivated")
}

func (p *qtreewidget) SignalItemActivated() func() {
	return getSignal(p.ObjectName(), "itemActivated")
}

func (p *qtreewidget) ConnectSignalItemChanged(f func()) {
	C.QTreeWidget_ConnectSignalItemChanged(p.Pointer())
	connectSignal(p.ObjectName(), "itemChanged", f)
}

func (p *qtreewidget) DisconnectSignalItemChanged() {
	C.QTreeWidget_DisconnectSignalItemChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemChanged")
}

func (p *qtreewidget) SignalItemChanged() func() {
	return getSignal(p.ObjectName(), "itemChanged")
}

func (p *qtreewidget) ConnectSignalItemClicked(f func()) {
	C.QTreeWidget_ConnectSignalItemClicked(p.Pointer())
	connectSignal(p.ObjectName(), "itemClicked", f)
}

func (p *qtreewidget) DisconnectSignalItemClicked() {
	C.QTreeWidget_DisconnectSignalItemClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemClicked")
}

func (p *qtreewidget) SignalItemClicked() func() {
	return getSignal(p.ObjectName(), "itemClicked")
}

func (p *qtreewidget) ConnectSignalItemCollapsed(f func()) {
	C.QTreeWidget_ConnectSignalItemCollapsed(p.Pointer())
	connectSignal(p.ObjectName(), "itemCollapsed", f)
}

func (p *qtreewidget) DisconnectSignalItemCollapsed() {
	C.QTreeWidget_DisconnectSignalItemCollapsed(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemCollapsed")
}

func (p *qtreewidget) SignalItemCollapsed() func() {
	return getSignal(p.ObjectName(), "itemCollapsed")
}

func (p *qtreewidget) ConnectSignalItemDoubleClicked(f func()) {
	C.QTreeWidget_ConnectSignalItemDoubleClicked(p.Pointer())
	connectSignal(p.ObjectName(), "itemDoubleClicked", f)
}

func (p *qtreewidget) DisconnectSignalItemDoubleClicked() {
	C.QTreeWidget_DisconnectSignalItemDoubleClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemDoubleClicked")
}

func (p *qtreewidget) SignalItemDoubleClicked() func() {
	return getSignal(p.ObjectName(), "itemDoubleClicked")
}

func (p *qtreewidget) ConnectSignalItemEntered(f func()) {
	C.QTreeWidget_ConnectSignalItemEntered(p.Pointer())
	connectSignal(p.ObjectName(), "itemEntered", f)
}

func (p *qtreewidget) DisconnectSignalItemEntered() {
	C.QTreeWidget_DisconnectSignalItemEntered(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemEntered")
}

func (p *qtreewidget) SignalItemEntered() func() {
	return getSignal(p.ObjectName(), "itemEntered")
}

func (p *qtreewidget) ConnectSignalItemExpanded(f func()) {
	C.QTreeWidget_ConnectSignalItemExpanded(p.Pointer())
	connectSignal(p.ObjectName(), "itemExpanded", f)
}

func (p *qtreewidget) DisconnectSignalItemExpanded() {
	C.QTreeWidget_DisconnectSignalItemExpanded(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemExpanded")
}

func (p *qtreewidget) SignalItemExpanded() func() {
	return getSignal(p.ObjectName(), "itemExpanded")
}

func (p *qtreewidget) ConnectSignalItemPressed(f func()) {
	C.QTreeWidget_ConnectSignalItemPressed(p.Pointer())
	connectSignal(p.ObjectName(), "itemPressed", f)
}

func (p *qtreewidget) DisconnectSignalItemPressed() {
	C.QTreeWidget_DisconnectSignalItemPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemPressed")
}

func (p *qtreewidget) SignalItemPressed() func() {
	return getSignal(p.ObjectName(), "itemPressed")
}

func (p *qtreewidget) ConnectSignalItemSelectionChanged(f func()) {
	C.QTreeWidget_ConnectSignalItemSelectionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "itemSelectionChanged", f)
}

func (p *qtreewidget) DisconnectSignalItemSelectionChanged() {
	C.QTreeWidget_DisconnectSignalItemSelectionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemSelectionChanged")
}

func (p *qtreewidget) SignalItemSelectionChanged() func() {
	return getSignal(p.ObjectName(), "itemSelectionChanged")
}

//export callbackQTreeWidget
func callbackQTreeWidget(ptr C.QtObjectPtr, msg *C.char) {
	var qtreewidget = new(qtreewidget)
	qtreewidget.SetPointer(ptr)
	getSignal(qtreewidget.ObjectName(), C.GoString(msg))()
}
