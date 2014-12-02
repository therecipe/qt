package qt

//#include "qradiobutton.h"
import "C"

type qradiobutton struct {
	qabstractbutton
}

type QRadioButton interface {
	QAbstractButton
}

func (p *qradiobutton) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qradiobutton) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQRadioButton_QWidget(parent QWidget) QRadioButton {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qradiobutton = new(qradiobutton)
	qradiobutton.SetPointer(C.QRadioButton_New_QWidget(parentPtr))
	qradiobutton.SetObjectName_String("QRadioButton_" + randomIdentifier())
	return qradiobutton
}

func NewQRadioButton_String_QWidget(text string, parent QWidget) QRadioButton {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qradiobutton = new(qradiobutton)
	qradiobutton.SetPointer(C.QRadioButton_New_String_QWidget(C.CString(text), parentPtr))
	qradiobutton.SetObjectName_String("QRadioButton_" + randomIdentifier())
	return qradiobutton
}

func (p *qradiobutton) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QRadioButton_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}
