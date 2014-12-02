package qt

//#include "qbuttongroup.h"
import "C"

type qbuttongroup struct {
	qobject
}

type QButtonGroup interface {
	QObject
	AddButton_QAbstractButton_Int(button QAbstractButton, id int)
	Button_Int(id int) QAbstractButton
	CheckedButton() QAbstractButton
	CheckedId() int
	Exclusive() bool
	Id_QAbstractButton(button QAbstractButton) int
	RemoveButton_QAbstractButton(button QAbstractButton)
	SetExclusive_Bool(exclusive bool)
	SetId_QAbstractButton_Int(button QAbstractButton, id int)
}

func (p *qbuttongroup) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qbuttongroup) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQButtonGroup_QObject(parent QObject) QButtonGroup {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qbuttongroup = new(qbuttongroup)
	qbuttongroup.SetPointer(C.QButtonGroup_New_QObject(parentPtr))
	qbuttongroup.SetObjectName_String("QButtonGroup_" + randomIdentifier())
	return qbuttongroup
}

func (p *qbuttongroup) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QButtonGroup_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qbuttongroup) AddButton_QAbstractButton_Int(button QAbstractButton, id int) {
	if p.Pointer() == nil {
	} else {
		var buttonPtr C.QtObjectPtr = nil
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QButtonGroup_AddButton_QAbstractButton_Int(p.Pointer(), buttonPtr, C.int(id))
	}
}

func (p *qbuttongroup) Button_Int(id int) QAbstractButton {
	if p.Pointer() == nil {
		return nil
	} else {
		var qabstractbutton = new(qabstractbutton)
		qabstractbutton.SetPointer(C.QButtonGroup_Button_Int(p.Pointer(), C.int(id)))
		if qabstractbutton.ObjectName() == "" {
			qabstractbutton.SetObjectName_String("QAbstractButton_" + randomIdentifier())
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
			qabstractbutton.SetObjectName_String("QAbstractButton_" + randomIdentifier())
		}
		return qabstractbutton
	}
}

func (p *qbuttongroup) CheckedId() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QButtonGroup_CheckedId(p.Pointer()))
	}
}

func (p *qbuttongroup) Exclusive() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QButtonGroup_Exclusive(p.Pointer()) != 0
	}
}

func (p *qbuttongroup) Id_QAbstractButton(button QAbstractButton) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var buttonPtr C.QtObjectPtr = nil
		if button != nil {
			buttonPtr = button.Pointer()
		}
		return int(C.QButtonGroup_Id_QAbstractButton(p.Pointer(), buttonPtr))
	}
}

func (p *qbuttongroup) RemoveButton_QAbstractButton(button QAbstractButton) {
	if p.Pointer() == nil {
	} else {
		var buttonPtr C.QtObjectPtr = nil
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QButtonGroup_RemoveButton_QAbstractButton(p.Pointer(), buttonPtr)
	}
}

func (p *qbuttongroup) SetExclusive_Bool(exclusive bool) {
	if p.Pointer() != nil {
		C.QButtonGroup_SetExclusive_Bool(p.Pointer(), goBoolToCInt(exclusive))
	}
}

func (p *qbuttongroup) SetId_QAbstractButton_Int(button QAbstractButton, id int) {
	if p.Pointer() == nil {
	} else {
		var buttonPtr C.QtObjectPtr = nil
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QButtonGroup_SetId_QAbstractButton_Int(p.Pointer(), buttonPtr, C.int(id))
	}
}
