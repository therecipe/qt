package qt

//#include "qtoolbutton.h"
import "C"

type qtoolbutton struct {
	qabstractbutton
}

type QToolButton interface {
	QAbstractButton
	ArrowType() ArrowType
	AutoRaise() bool
	Menu() QMenu
	SetArrowType_ArrowType(typ ArrowType)
	SetAutoRaise_Bool(enable bool)
	SetMenu_QMenu(menu QMenu)
	ToolButtonStyle() ToolButtonStyle
	ConnectSlotShowMenu()
	DisconnectSlotShowMenu()
	SlotShowMenu()
	ConnectSignalTriggered(f func())
	DisconnectSignalTriggered()
	SignalTriggered() func()
}

func (p *qtoolbutton) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtoolbutton) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQToolButton_QWidget(parent QWidget) QToolButton {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtoolbutton = new(qtoolbutton)
	qtoolbutton.SetPointer(C.QToolButton_New_QWidget(parentPtr))
	qtoolbutton.SetObjectName_String("QToolButton_" + randomIdentifier())
	return qtoolbutton
}

func (p *qtoolbutton) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QToolButton_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtoolbutton) ArrowType() ArrowType {
	if p.Pointer() == nil {
		return 0
	} else {
		return ArrowType(C.QToolButton_ArrowType(p.Pointer()))
	}
}

func (p *qtoolbutton) AutoRaise() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QToolButton_AutoRaise(p.Pointer()) != 0
	}
}

func (p *qtoolbutton) Menu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QToolButton_Menu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qtoolbutton) SetArrowType_ArrowType(typ ArrowType) {
	if p.Pointer() != nil {
		C.QToolButton_SetArrowType_ArrowType(p.Pointer(), C.int(typ))
	}
}

func (p *qtoolbutton) SetAutoRaise_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QToolButton_SetAutoRaise_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtoolbutton) SetMenu_QMenu(menu QMenu) {
	if p.Pointer() == nil {
	} else {
		var menuPtr C.QtObjectPtr = nil
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		C.QToolButton_SetMenu_QMenu(p.Pointer(), menuPtr)
	}
}

func (p *qtoolbutton) ToolButtonStyle() ToolButtonStyle {
	if p.Pointer() == nil {
		return 0
	} else {
		return ToolButtonStyle(C.QToolButton_ToolButtonStyle(p.Pointer()))
	}
}

func (p *qtoolbutton) ConnectSlotShowMenu() {
	C.QToolButton_ConnectSlotShowMenu(p.Pointer())
}

func (p *qtoolbutton) DisconnectSlotShowMenu() {
	C.QToolButton_DisconnectSlotShowMenu(p.Pointer())
}

func (p *qtoolbutton) SlotShowMenu() {
	if p.Pointer() != nil {
		C.QToolButton_ShowMenu(p.Pointer())
	}
}

func (p *qtoolbutton) ConnectSignalTriggered(f func()) {
	C.QToolButton_ConnectSignalTriggered(p.Pointer())
	connectSignal(p.ObjectName(), "triggered", f)
}

func (p *qtoolbutton) DisconnectSignalTriggered() {
	C.QToolButton_DisconnectSignalTriggered(p.Pointer())
	disconnectSignal(p.ObjectName(), "triggered")
}

func (p *qtoolbutton) SignalTriggered() func() {
	return getSignal(p.ObjectName(), "triggered")
}

//export callbackQToolButton
func callbackQToolButton(ptr C.QtObjectPtr, msg *C.char) {
	var qtoolbutton = new(qtoolbutton)
	qtoolbutton.SetPointer(ptr)
	getSignal(qtoolbutton.ObjectName(), C.GoString(msg))()
}
