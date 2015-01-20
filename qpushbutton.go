package qt

//#include "qpushbutton.h"
import "C"

type qpushbutton struct {
	qabstractbutton
}

type QPushButton interface {
	QAbstractButton
	AutoDefault() bool
	IsDefault() bool
	IsFlat() bool
	Menu() QMenu
	SetAutoDefault(autoDefault bool)
	SetDefault(defaul bool)
	SetFlat(flat bool)
	SetMenu(menu QMenu)
	ConnectSlotShowMenu()
	DisconnectSlotShowMenu()
	SlotShowMenu()
}

func (p *qpushbutton) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qpushbutton) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQPushButton1(parent QWidget) QPushButton {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qpushbutton = new(qpushbutton)
	qpushbutton.SetPointer(C.QPushButton_New_QWidget(parentPtr))
	qpushbutton.SetObjectName("QPushButton_" + randomIdentifier())
	return qpushbutton
}

func NewQPushButton2(text string, parent QWidget) QPushButton {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qpushbutton = new(qpushbutton)
	qpushbutton.SetPointer(C.QPushButton_New_String_QWidget(C.CString(text), parentPtr))
	qpushbutton.SetObjectName("QPushButton_" + randomIdentifier())
	return qpushbutton
}

func (p *qpushbutton) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QPushButton_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qpushbutton) AutoDefault() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QPushButton_AutoDefault(p.Pointer()) != 0
}

func (p *qpushbutton) IsDefault() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QPushButton_IsDefault(p.Pointer()) != 0
}

func (p *qpushbutton) IsFlat() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QPushButton_IsFlat(p.Pointer()) != 0
}

func (p *qpushbutton) Menu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QPushButton_Menu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qpushbutton) SetAutoDefault(autoDefault bool) {
	if p.Pointer() != nil {
		C.QPushButton_SetAutoDefault_Bool(p.Pointer(), goBoolToCInt(autoDefault))
	}
}

func (p *qpushbutton) SetDefault(defaul bool) {
	if p.Pointer() != nil {
		C.QPushButton_SetDefault_Bool(p.Pointer(), goBoolToCInt(defaul))
	}
}

func (p *qpushbutton) SetFlat(flat bool) {
	if p.Pointer() != nil {
		C.QPushButton_SetFlat_Bool(p.Pointer(), goBoolToCInt(flat))
	}
}

func (p *qpushbutton) SetMenu(menu QMenu) {
	if p.Pointer() != nil {
		var menuPtr C.QtObjectPtr
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		C.QPushButton_SetMenu_QMenu(p.Pointer(), menuPtr)
	}
}

func (p *qpushbutton) ConnectSlotShowMenu() {
	C.QPushButton_ConnectSlotShowMenu(p.Pointer())
}

func (p *qpushbutton) DisconnectSlotShowMenu() {
	C.QPushButton_DisconnectSlotShowMenu(p.Pointer())
}

func (p *qpushbutton) SlotShowMenu() {
	if p.Pointer() != nil {
		C.QPushButton_ShowMenu(p.Pointer())
	}
}
