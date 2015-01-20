package qt

//#include "qintvalidator.h"
import "C"

type qintvalidator struct {
	qvalidator
}

type QIntValidator interface {
	QValidator
	Bottom() int
	SetBottom(in int)
	SetRange(bottom int, top int)
	SetTop(in int)
	Top() int
}

func (p *qintvalidator) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qintvalidator) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQIntValidator1(parent QObject) QIntValidator {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qintvalidator = new(qintvalidator)
	qintvalidator.SetPointer(C.QIntValidator_New_QObject(parentPtr))
	qintvalidator.SetObjectName("QIntValidator_" + randomIdentifier())
	return qintvalidator
}

func NewQIntValidator2(minimum int, maximum int, parent QObject) QIntValidator {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qintvalidator = new(qintvalidator)
	qintvalidator.SetPointer(C.QIntValidator_New_Int_Int_QObject(C.int(minimum), C.int(maximum), parentPtr))
	qintvalidator.SetObjectName("QIntValidator_" + randomIdentifier())
	return qintvalidator
}

func (p *qintvalidator) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QIntValidator_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qintvalidator) Bottom() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QIntValidator_Bottom(p.Pointer()))
}

func (p *qintvalidator) SetBottom(in int) {
	if p.Pointer() != nil {
		C.QIntValidator_SetBottom_Int(p.Pointer(), C.int(in))
	}
}

func (p *qintvalidator) SetRange(bottom int, top int) {
	if p.Pointer() != nil {
		C.QIntValidator_SetRange_Int_Int(p.Pointer(), C.int(bottom), C.int(top))
	}
}

func (p *qintvalidator) SetTop(in int) {
	if p.Pointer() != nil {
		C.QIntValidator_SetTop_Int(p.Pointer(), C.int(in))
	}
}

func (p *qintvalidator) Top() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QIntValidator_Top(p.Pointer()))
}
