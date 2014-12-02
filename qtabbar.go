package qt

//#include "qtabbar.h"
import "C"

type qtabbar struct {
	qwidget
}

type QTabBar interface {
	QWidget
	AddTab_String(text string) int
	Count() int
	CurrentIndex() int
	DocumentMode() bool
	DrawBase() bool
	ElideMode() TextElideMode
	Expanding() bool
	InsertTab_Int_String(index int, text string) int
	IsMovable() bool
	IsTabEnabled_Int(index int) bool
	MoveTab_Int_Int(from int, to int)
	RemoveTab_Int(index int)
	SetDocumentMode_Bool(set bool)
	SetDrawBase_Bool(drawTheBase bool)
	SetElideMode_TextElideMode(TextElideMode TextElideMode)
	SetExpanding_Bool(enabled bool)
	SetMovable_Bool(movable bool)
	SetTabEnabled_Int_Bool(index int, enabled bool)
	SetTabText_Int_String(index int, text string)
	SetTabToolTip_Int_String(index int, tip string)
	SetTabWhatsThis_Int_String(index int, text string)
	SetTabsClosable_Bool(closable bool)
	SetUsesScrollButtons_Bool(useButtons bool)
	TabText_Int(index int) string
	TabToolTip_Int(index int) string
	TabWhatsThis_Int(index int) string
	TabsClosable() bool
	UsesScrollButtons() bool
	ConnectSlotSetCurrentIndex()
	DisconnectSlotSetCurrentIndex()
	SlotSetCurrentIndex_Int(index int)
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
	ConnectSignalTabMoved(f func())
	DisconnectSignalTabMoved()
	SignalTabMoved() func()
}

func (p *qtabbar) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtabbar) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTabBar_QWidget(parent QWidget) QTabBar {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtabbar = new(qtabbar)
	qtabbar.SetPointer(C.QTabBar_New_QWidget(parentPtr))
	qtabbar.SetObjectName_String("QTabBar_" + randomIdentifier())
	return qtabbar
}

func (p *qtabbar) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTabBar_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtabbar) AddTab_String(text string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTabBar_AddTab_String(p.Pointer(), C.CString(text)))
	}
}

func (p *qtabbar) Count() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTabBar_Count(p.Pointer()))
	}
}

func (p *qtabbar) CurrentIndex() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTabBar_CurrentIndex(p.Pointer()))
	}
}

func (p *qtabbar) DocumentMode() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_DocumentMode(p.Pointer()) != 0
	}
}

func (p *qtabbar) DrawBase() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_DrawBase(p.Pointer()) != 0
	}
}

func (p *qtabbar) ElideMode() TextElideMode {
	if p.Pointer() == nil {
		return 0
	} else {
		return TextElideMode(C.QTabBar_ElideMode(p.Pointer()))
	}
}

func (p *qtabbar) Expanding() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_Expanding(p.Pointer()) != 0
	}
}

func (p *qtabbar) InsertTab_Int_String(index int, text string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTabBar_InsertTab_Int_String(p.Pointer(), C.int(index), C.CString(text)))
	}
}

func (p *qtabbar) IsMovable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_IsMovable(p.Pointer()) != 0
	}
}

func (p *qtabbar) IsTabEnabled_Int(index int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_IsTabEnabled_Int(p.Pointer(), C.int(index)) != 0
	}
}

func (p *qtabbar) MoveTab_Int_Int(from int, to int) {
	if p.Pointer() != nil {
		C.QTabBar_MoveTab_Int_Int(p.Pointer(), C.int(from), C.int(to))
	}
}

func (p *qtabbar) RemoveTab_Int(index int) {
	if p.Pointer() != nil {
		C.QTabBar_RemoveTab_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtabbar) SetDocumentMode_Bool(set bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetDocumentMode_Bool(p.Pointer(), goBoolToCInt(set))
	}
}

func (p *qtabbar) SetDrawBase_Bool(drawTheBase bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetDrawBase_Bool(p.Pointer(), goBoolToCInt(drawTheBase))
	}
}

func (p *qtabbar) SetElideMode_TextElideMode(TextElideMode TextElideMode) {
	if p.Pointer() != nil {
		C.QTabBar_SetElideMode_TextElideMode(p.Pointer(), C.int(TextElideMode))
	}
}

func (p *qtabbar) SetExpanding_Bool(enabled bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetExpanding_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qtabbar) SetMovable_Bool(movable bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetMovable_Bool(p.Pointer(), goBoolToCInt(movable))
	}
}

func (p *qtabbar) SetTabEnabled_Int_Bool(index int, enabled bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetTabEnabled_Int_Bool(p.Pointer(), C.int(index), goBoolToCInt(enabled))
	}
}

func (p *qtabbar) SetTabText_Int_String(index int, text string) {
	if p.Pointer() != nil {
		C.QTabBar_SetTabText_Int_String(p.Pointer(), C.int(index), C.CString(text))
	}
}

func (p *qtabbar) SetTabToolTip_Int_String(index int, tip string) {
	if p.Pointer() != nil {
		C.QTabBar_SetTabToolTip_Int_String(p.Pointer(), C.int(index), C.CString(tip))
	}
}

func (p *qtabbar) SetTabWhatsThis_Int_String(index int, text string) {
	if p.Pointer() != nil {
		C.QTabBar_SetTabWhatsThis_Int_String(p.Pointer(), C.int(index), C.CString(text))
	}
}

func (p *qtabbar) SetTabsClosable_Bool(closable bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetTabsClosable_Bool(p.Pointer(), goBoolToCInt(closable))
	}
}

func (p *qtabbar) SetUsesScrollButtons_Bool(useButtons bool) {
	if p.Pointer() != nil {
		C.QTabBar_SetUsesScrollButtons_Bool(p.Pointer(), goBoolToCInt(useButtons))
	}
}

func (p *qtabbar) TabText_Int(index int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTabBar_TabText_Int(p.Pointer(), C.int(index)))
	}
}

func (p *qtabbar) TabToolTip_Int(index int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTabBar_TabToolTip_Int(p.Pointer(), C.int(index)))
	}
}

func (p *qtabbar) TabWhatsThis_Int(index int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTabBar_TabWhatsThis_Int(p.Pointer(), C.int(index)))
	}
}

func (p *qtabbar) TabsClosable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_TabsClosable(p.Pointer()) != 0
	}
}

func (p *qtabbar) UsesScrollButtons() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabBar_UsesScrollButtons(p.Pointer()) != 0
	}
}

func (p *qtabbar) ConnectSlotSetCurrentIndex() {
	C.QTabBar_ConnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qtabbar) DisconnectSlotSetCurrentIndex() {
	C.QTabBar_DisconnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qtabbar) SlotSetCurrentIndex_Int(index int) {
	if p.Pointer() != nil {
		C.QTabBar_SetCurrentIndex_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtabbar) ConnectSignalCurrentChanged(f func()) {
	C.QTabBar_ConnectSignalCurrentChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentChanged", f)
}

func (p *qtabbar) DisconnectSignalCurrentChanged() {
	C.QTabBar_DisconnectSignalCurrentChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentChanged")
}

func (p *qtabbar) SignalCurrentChanged() func() {
	return getSignal(p.ObjectName(), "currentChanged")
}

func (p *qtabbar) ConnectSignalTabBarClicked(f func()) {
	C.QTabBar_ConnectSignalTabBarClicked(p.Pointer())
	connectSignal(p.ObjectName(), "tabBarClicked", f)
}

func (p *qtabbar) DisconnectSignalTabBarClicked() {
	C.QTabBar_DisconnectSignalTabBarClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabBarClicked")
}

func (p *qtabbar) SignalTabBarClicked() func() {
	return getSignal(p.ObjectName(), "tabBarClicked")
}

func (p *qtabbar) ConnectSignalTabBarDoubleClicked(f func()) {
	C.QTabBar_ConnectSignalTabBarDoubleClicked(p.Pointer())
	connectSignal(p.ObjectName(), "tabBarDoubleClicked", f)
}

func (p *qtabbar) DisconnectSignalTabBarDoubleClicked() {
	C.QTabBar_DisconnectSignalTabBarDoubleClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabBarDoubleClicked")
}

func (p *qtabbar) SignalTabBarDoubleClicked() func() {
	return getSignal(p.ObjectName(), "tabBarDoubleClicked")
}

func (p *qtabbar) ConnectSignalTabCloseRequested(f func()) {
	C.QTabBar_ConnectSignalTabCloseRequested(p.Pointer())
	connectSignal(p.ObjectName(), "tabCloseRequested", f)
}

func (p *qtabbar) DisconnectSignalTabCloseRequested() {
	C.QTabBar_DisconnectSignalTabCloseRequested(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabCloseRequested")
}

func (p *qtabbar) SignalTabCloseRequested() func() {
	return getSignal(p.ObjectName(), "tabCloseRequested")
}

func (p *qtabbar) ConnectSignalTabMoved(f func()) {
	C.QTabBar_ConnectSignalTabMoved(p.Pointer())
	connectSignal(p.ObjectName(), "tabMoved", f)
}

func (p *qtabbar) DisconnectSignalTabMoved() {
	C.QTabBar_DisconnectSignalTabMoved(p.Pointer())
	disconnectSignal(p.ObjectName(), "tabMoved")
}

func (p *qtabbar) SignalTabMoved() func() {
	return getSignal(p.ObjectName(), "tabMoved")
}

//export callbackQTabBar
func callbackQTabBar(ptr C.QtObjectPtr, msg *C.char) {
	var qtabbar = new(qtabbar)
	qtabbar.SetPointer(ptr)
	getSignal(qtabbar.ObjectName(), C.GoString(msg))()
}
