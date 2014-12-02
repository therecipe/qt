package qt

//#include "qtabwidget.h"
import "C"

type qtabwidget struct {
	qwidget
}

type QTabWidget interface {
	QWidget
	AddTab_QWidget_String(page QWidget, label string) int
	Clear()
	CornerWidget_Corner(corner Corner) QWidget
	Count() int
	CurrentIndex() int
	CurrentWidget() QWidget
	DocumentMode() bool
	ElideMode() TextElideMode
	IndexOf_QWidget(w QWidget) int
	InsertTab_Int_QWidget_String(index int, page QWidget, label string) int
	IsMovable() bool
	IsTabEnabled_Int(index int) bool
	RemoveTab_Int(index int)
	SetCornerWidget_QWidget_Corner(widget QWidget, corner Corner)
	SetDocumentMode_Bool(set bool)
	SetElideMode_TextElideMode(TextElideMode TextElideMode)
	SetMovable_Bool(movable bool)
	SetTabEnabled_Int_Bool(index int, enable bool)
	SetTabText_Int_String(index int, label string)
	SetTabToolTip_Int_String(index int, tip string)
	SetTabWhatsThis_Int_String(index int, text string)
	SetTabsClosable_Bool(closeable bool)
	SetUsesScrollButtons_Bool(useButtons bool)
	TabBar() QTabBar
	TabText_Int(index int) string
	TabToolTip_Int(index int) string
	TabWhatsThis_Int(index int) string
	TabsClosable() bool
	UsesScrollButtons() bool
	Widget_Int(index int) QWidget
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
}

func (p *qtabwidget) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtabwidget) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTabWidget_QWidget(parent QWidget) QTabWidget {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtabwidget = new(qtabwidget)
	qtabwidget.SetPointer(C.QTabWidget_New_QWidget(parentPtr))
	qtabwidget.SetObjectName_String("QTabWidget_" + randomIdentifier())
	return qtabwidget
}

func (p *qtabwidget) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTabWidget_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtabwidget) AddTab_QWidget_String(page QWidget, label string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var pagePtr C.QtObjectPtr = nil
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

func (p *qtabwidget) CornerWidget_Corner(corner Corner) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTabWidget_CornerWidget_Corner(p.Pointer(), C.int(corner)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtabwidget) Count() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTabWidget_Count(p.Pointer()))
	}
}

func (p *qtabwidget) CurrentIndex() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTabWidget_CurrentIndex(p.Pointer()))
	}
}

func (p *qtabwidget) CurrentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTabWidget_CurrentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtabwidget) DocumentMode() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabWidget_DocumentMode(p.Pointer()) != 0
	}
}

func (p *qtabwidget) ElideMode() TextElideMode {
	if p.Pointer() == nil {
		return 0
	} else {
		return TextElideMode(C.QTabWidget_ElideMode(p.Pointer()))
	}
}

func (p *qtabwidget) IndexOf_QWidget(w QWidget) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var wPtr C.QtObjectPtr = nil
		if w != nil {
			wPtr = w.Pointer()
		}
		return int(C.QTabWidget_IndexOf_QWidget(p.Pointer(), wPtr))
	}
}

func (p *qtabwidget) InsertTab_Int_QWidget_String(index int, page QWidget, label string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var pagePtr C.QtObjectPtr = nil
		if page != nil {
			pagePtr = page.Pointer()
		}
		return int(C.QTabWidget_InsertTab_Int_QWidget_String(p.Pointer(), C.int(index), pagePtr, C.CString(label)))
	}
}

func (p *qtabwidget) IsMovable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabWidget_IsMovable(p.Pointer()) != 0
	}
}

func (p *qtabwidget) IsTabEnabled_Int(index int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabWidget_IsTabEnabled_Int(p.Pointer(), C.int(index)) != 0
	}
}

func (p *qtabwidget) RemoveTab_Int(index int) {
	if p.Pointer() != nil {
		C.QTabWidget_RemoveTab_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtabwidget) SetCornerWidget_QWidget_Corner(widget QWidget, corner Corner) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QTabWidget_SetCornerWidget_QWidget_Corner(p.Pointer(), widgetPtr, C.int(corner))
	}
}

func (p *qtabwidget) SetDocumentMode_Bool(set bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetDocumentMode_Bool(p.Pointer(), goBoolToCInt(set))
	}
}

func (p *qtabwidget) SetElideMode_TextElideMode(TextElideMode TextElideMode) {
	if p.Pointer() != nil {
		C.QTabWidget_SetElideMode_TextElideMode(p.Pointer(), C.int(TextElideMode))
	}
}

func (p *qtabwidget) SetMovable_Bool(movable bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetMovable_Bool(p.Pointer(), goBoolToCInt(movable))
	}
}

func (p *qtabwidget) SetTabEnabled_Int_Bool(index int, enable bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabEnabled_Int_Bool(p.Pointer(), C.int(index), goBoolToCInt(enable))
	}
}

func (p *qtabwidget) SetTabText_Int_String(index int, label string) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabText_Int_String(p.Pointer(), C.int(index), C.CString(label))
	}
}

func (p *qtabwidget) SetTabToolTip_Int_String(index int, tip string) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabToolTip_Int_String(p.Pointer(), C.int(index), C.CString(tip))
	}
}

func (p *qtabwidget) SetTabWhatsThis_Int_String(index int, text string) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabWhatsThis_Int_String(p.Pointer(), C.int(index), C.CString(text))
	}
}

func (p *qtabwidget) SetTabsClosable_Bool(closeable bool) {
	if p.Pointer() != nil {
		C.QTabWidget_SetTabsClosable_Bool(p.Pointer(), goBoolToCInt(closeable))
	}
}

func (p *qtabwidget) SetUsesScrollButtons_Bool(useButtons bool) {
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
			qtabbar.SetObjectName_String("QTabBar_" + randomIdentifier())
		}
		return qtabbar
	}
}

func (p *qtabwidget) TabText_Int(index int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTabWidget_TabText_Int(p.Pointer(), C.int(index)))
	}
}

func (p *qtabwidget) TabToolTip_Int(index int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTabWidget_TabToolTip_Int(p.Pointer(), C.int(index)))
	}
}

func (p *qtabwidget) TabWhatsThis_Int(index int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTabWidget_TabWhatsThis_Int(p.Pointer(), C.int(index)))
	}
}

func (p *qtabwidget) TabsClosable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabWidget_TabsClosable(p.Pointer()) != 0
	}
}

func (p *qtabwidget) UsesScrollButtons() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTabWidget_UsesScrollButtons(p.Pointer()) != 0
	}
}

func (p *qtabwidget) Widget_Int(index int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTabWidget_Widget_Int(p.Pointer(), C.int(index)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
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

func (p *qtabwidget) SlotSetCurrentIndex_Int(index int) {
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
