package qt

//#include "qtabwidget.h"
import "C"

type qtabwidget struct {
	qwidget
}

type QTabWidget interface {
	QWidget
	AddTab(page QWidget, label string) int
	Clear()
	CornerWidget(corner Corner) QWidget
	Count() int
	CurrentIndex() int
	CurrentWidget() QWidget
	DocumentMode() bool
	ElideMode() TextElideMode
	IndexOf(w QWidget) int
	InsertTab(index int, page QWidget, label string) int
	IsMovable() bool
	IsTabEnabled(index int) bool
	RemoveTab(index int)
	SetCornerWidget(widget QWidget, corner Corner)
	SetDocumentMode(set bool)
	SetElideMode(TextElideMode TextElideMode)
	SetMovable(movable bool)
	SetTabEnabled(index int, enable bool)
	SetTabText(index int, label string)
	SetTabToolTip(index int, tip string)
	SetTabWhatsThis(index int, text string)
	SetTabsClosable(closeable bool)
	SetUsesScrollButtons(useButtons bool)
	TabBar() QTabBar
	TabText(index int) string
	TabToolTip(index int) string
	TabWhatsThis(index int) string
	TabsClosable() bool
	UsesScrollButtons() bool
	Widget(index int) QWidget
	ConnectSlotSetCurrentIndex()
	DisconnectSlotSetCurrentIndex()
	SlotSetCurrentIndex(index int)
	ConnectSignalCurrentChanged(f func())
	DisconnectSignalCurrentChanged()
	SignalCurrentChanged() func()
	ConnectSignalTabBarClicked(f func())
	DisconnectSignalTabBarClicked()
	SignalTabBarClicked() func()
	ConnectSignalTabBarDoubleClicked(f func())
	DisconnectSignalTabBarDoubleClicked()
	SignalTabBarDoubleClicked() func()
	ConnectSignalTabCloseRequested(f func())
	DisconnectSignalTabCloseRequested()
	SignalTabCloseRequested() func()
}

func (p *qtabwidget) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtabwidget) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTabWidget(parent QWidget) QTabWidget {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtabwidget = new(qtabwidget)
	qtabwidget.SetPointer(C.QTabWidget_New_QWidget(parentPtr))
	qtabwidget.SetObjectName("QTabWidget_" + randomIdentifier())
	return qtabwidget
}

func (p *qtabwidget) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTabWidget_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtabwidget) AddTab(page QWidget, label string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var pagePtr C.QtObjectPtr
		if page != nil {
			pagePtr = page.Pointer()
		}
		return int(C.QTabWidget_AddTab_QWidget_String(p.Pointer(), pagePtr, C.CString(label)))
	}
}

func (p *qtabwidget) Clear() {
	if p.Pointer() != nil {
		C.QTabWidget_Clear(p.Pointer())
	}
}

func (p *qtabwidget) CornerWidget(corner Corner) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTabWidget_CornerWidget_Corner(p.Pointer(), C.int(corner)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtabwidget) Count() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTabWidget_Count(p.Pointer()))
}

func (p *qtabwidget) CurrentIndex() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTabWidget_CurrentIndex(p.Pointer()))
}

func (p *qtabwidget) CurrentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTabWidget_CurrentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtabwidget) DocumentMode() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTabWidget_DocumentMode(p.Pointer()) != 0
}

func (p *qtabwidget) ElideMode() TextElideMode {
	if p.Pointer() == nil {
		return 0
	}
	return TextElideMode(C.QTabWidget_ElideMode(p.Pointer()))
}

func (p *qtabwidget) IndexOf(w QWidget) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var wPtr C.QtObjectPtr
		if w != nil {
			wPtr = w.Pointer()
		}
		return int(C.QTabWidget_IndexOf_QWidget(p.Pointer(), wPtr))
	}
}

func (p *qtabwidget) InsertTab(index int, page QWidget, label string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var pagePtr C.QtObjectPtr
		if page != nil {
			pagePtr = page.Pointer()
		}
		return int(C.QTabWidget_InsertTab_Int_QWidget_String(p.Pointer(), C.int(index), pagePtr, C.CString(label)))
	}
}

func (p *qtabwidget) IsMovable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTabWidget_IsMovable(p.Pointer()) != 0
}

func (p *qtabwidget) IsTabEnabled(index int) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTabWidget_IsTabEnabled_Int(p.Pointer(), C.int(index)) != 0
}

func (p *qtabwidget) RemoveTab(index int) {
	if p.Pointer() != nil {
		C.QTabWidget_RemoveTab_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtabwidget) SetCornerWidget(widget QWidget, corner Corner) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QTabWidget_SetCornerWidget_QWidget_Corner(p.Pointer(), widgetPtr, C.int(corner))
	}
}

func (p *qtabwidget) SetDocumentMode(set bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetDocumentMode_Bool(p.Pointer(), goBoolToCInt(set))
	}
}

func (p *qtabwidget) SetElideMode(TextElideMode TextElideMode) {
	if p.Pointer() != nil {
		C.QTabWidget_SetElideMode_TextElideMode(p.Pointer(), C.int(TextElideMode))
	}
}

func (p *qtabwidget) SetMovable(movable bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetMovable_Bool(p.Pointer(), goBoolToCInt(movable))
	}
}

func (p *qtabwidget) SetTabEnabled(index int, enable bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabEnabled_Int_Bool(p.Pointer(), C.int(index), goBoolToCInt(enable))
	}
}

func (p *qtabwidget) SetTabText(index int, label string) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabText_Int_String(p.Pointer(), C.int(index), C.CString(label))
	}
}

func (p *qtabwidget) SetTabToolTip(index int, tip string) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabToolTip_Int_String(p.Pointer(), C.int(index), C.CString(tip))
	}
}

func (p *qtabwidget) SetTabWhatsThis(index int, text string) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabWhatsThis_Int_String(p.Pointer(), C.int(index), C.CString(text))
	}
}

func (p *qtabwidget) SetTabsClosable(closeable bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabsClosable_Bool(p.Pointer(), goBoolToCInt(closeable))
	}
}

func (p *qtabwidget) SetUsesScrollButtons(useButtons bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetUsesScrollButtons_Bool(p.Pointer(), goBoolToCInt(useButtons))
	}
}

func (p *qtabwidget) TabBar() QTabBar {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtabbar = new(qtabbar)
		qtabbar.SetPointer(C.QTabWidget_TabBar(p.Pointer()))
		if qtabbar.ObjectName() == "" {
			qtabbar.SetObjectName("QTabBar_" + randomIdentifier())
		}
		return qtabbar
	}
}

func (p *qtabwidget) TabText(index int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTabWidget_TabText_Int(p.Pointer(), C.int(index)))
}

func (p *qtabwidget) TabToolTip(index int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTabWidget_TabToolTip_Int(p.Pointer(), C.int(index)))
}

func (p *qtabwidget) TabWhatsThis(index int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QTabWidget_TabWhatsThis_Int(p.Pointer(), C.int(index)))
}

func (p *qtabwidget) TabsClosable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTabWidget_TabsClosable(p.Pointer()) != 0
}

func (p *qtabwidget) UsesScrollButtons() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTabWidget_UsesScrollButtons(p.Pointer()) != 0
}

func (p *qtabwidget) Widget(index int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTabWidget_Widget_Int(p.Pointer(), C.int(index)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtabwidget) ConnectSlotSetCurrentIndex() {
	C.QTabWidget_ConnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qtabwidget) DisconnectSlotSetCurrentIndex() {
	C.QTabWidget_DisconnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qtabwidget) SlotSetCurrentIndex(index int) {
	if p.Pointer() != nil {
		C.QTabWidget_SetCurrentIndex_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtabwidget) ConnectSignalCurrentChanged(f func()) {
	C.QTabWidget_ConnectSignalCurrentChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentChanged", f)
}

func (p *qtabwidget) DisconnectSignalCurrentChanged() {
	C.QTabWidget_DisconnectSignalCurrentChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentChanged")
}

func (p *qtabwidget) SignalCurrentChanged() func() {
	return getSignal(p.ObjectName(), "currentChanged")
}

func (p *qtabwidget) ConnectSignalTabBarClicked(f func()) {
	C.QTabWidget_ConnectSignalTabBarClicked(p.Pointer())
	connectSignal(p.ObjectName(), "tabBarClicked", f)
}

func (p *qtabwidget) DisconnectSignalTabBarClicked() {
	C.QTabWidget_DisconnectSignalTabBarClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabBarClicked")
}

func (p *qtabwidget) SignalTabBarClicked() func() {
	return getSignal(p.ObjectName(), "tabBarClicked")
}

func (p *qtabwidget) ConnectSignalTabBarDoubleClicked(f func()) {
	C.QTabWidget_ConnectSignalTabBarDoubleClicked(p.Pointer())
	connectSignal(p.ObjectName(), "tabBarDoubleClicked", f)
}

func (p *qtabwidget) DisconnectSignalTabBarDoubleClicked() {
	C.QTabWidget_DisconnectSignalTabBarDoubleClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabBarDoubleClicked")
}

func (p *qtabwidget) SignalTabBarDoubleClicked() func() {
	return getSignal(p.ObjectName(), "tabBarDoubleClicked")
}

func (p *qtabwidget) ConnectSignalTabCloseRequested(f func()) {
	C.QTabWidget_ConnectSignalTabCloseRequested(p.Pointer())
	connectSignal(p.ObjectName(), "tabCloseRequested", f)
}

func (p *qtabwidget) DisconnectSignalTabCloseRequested() {
	C.QTabWidget_DisconnectSignalTabCloseRequested(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabCloseRequested")
}

func (p *qtabwidget) SignalTabCloseRequested() func() {
	return getSignal(p.ObjectName(), "tabCloseRequested")
}

//export callbackQTabWidget
func callbackQTabWidget(ptr C.QtObjectPtr, msg *C.char) {
	var qtabwidget = new(qtabwidget)
	qtabwidget.SetPointer(ptr)
	getSignal(qtabwidget.ObjectName(), C.GoString(msg))()
}
