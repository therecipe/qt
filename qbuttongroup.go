package qt

//#include "qbuttongroup.h"
import "C"

type qbuttongroup struct {
	qobject
}

type QButtonGroup interface {
	QObject
	AddButton(button QAbstractButton, id int)
	Button(id int) QAbstractButton
	CheckedButton() QAbstractButton
	CheckedId() int
	Exclusive() bool
	Id(button QAbstractButton) int
	RemoveButton(button QAbstractButton)
	SetExclusive(exclusive bool)
	SetId(button QAbstractButton, id int)
}

func (p *qbuttongroup) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qbuttongroup) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQButtonGroup(parent QObject) QButtonGroup {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qbuttongroup = new(qbuttongroup)
	qbuttongroup.SetPointer(C.QButtonGroup_New_QObject(parentPtr))
	qbuttongroup.SetObjectName("QButtonGroup_" + randomIdentifier())
	return qbuttongroup
}

func (p *qbuttongroup) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QButtonGroup_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qbuttongroup) AddButton(button QAbstractButton, id int) {
	if p.Pointer() != nil {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QButtonGroup_AddButton_QAbstractButton_Int(p.Pointer(), buttonPtr, C.int(id))
	}
}

func (p *qbuttongroup) Button(id int) QAbstractButton {
	if p.Pointer() == nil {
		return nil
	} else {
		var qabstractbutton = new(qabstractbutton)
		qabstractbutton.SetPointer(C.QButtonGroup_Button_Int(p.Pointer(), C.int(id)))
		if qabstractbutton.ObjectName() == "" {
			qabstractbutton.SetObjectName("QAbstractButton_" + randomIdentifier())
		}
		return qabstractbutton
	}
}

func (p *qbuttongroup) CheckedButton() QAbstractButton {
	if p.Pointer() == nil {
		return nil
	} else {
		var qabstractbutton = new(qabstractbutton)
		qabstractbutton.SetPointer(C.QButtonGroup_CheckedButton(p.Pointer()))
		if qabstractbutton.ObjectName() == "" {
			qabstractbutton.SetObjectName("QAbstractButton_" + randomIdentifier())
		}
		return qabstractbutton
	}
}

func (p *qbuttongroup) CheckedId() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QButtonGroup_CheckedId(p.Pointer()))
}

func (p *qbuttongroup) Exclusive() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QButtonGroup_Exclusive(p.Pointer()) != 0
}

func (p *qbuttongroup) Id(button QAbstractButton) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		return int(C.QButtonGroup_Id_QAbstractButton(p.Pointer(), buttonPtr))
	}
}

func (p *qbuttongroup) RemoveButton(button QAbstractButton) {
	if p.Pointer() != nil {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QButtonGroup_RemoveButton_QAbstractButton(p.Pointer(), buttonPtr)
	}
}

func (p *qbuttongroup) SetExclusive(exclusive bool) {
	if p.Pointer() != nil {
		C.QButtonGroup_SetExclusive_Bool(p.Pointer(), goBoolToCInt(exclusive))
	}
}

func (p *qbuttongroup) SetId(button QAbstractButton, id int) {
	if p.Pointer() != nil {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QButtonGroup_SetId_QAbstractButton_Int(p.Pointer(), buttonPtr, C.int(id))
	}
}
