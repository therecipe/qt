package qt

//#include "qintvalidator.h"
import "C"

type qintvalidator struct {
	qvalidator
}

type QIntValidator interface {
	QValidator
	Bottom() int
	SetBottom_Int(in int)
	SetRange_Int_Int(bottom int, top int)
	SetTop_Int(in int)
	Top() int
}

func (p *qintvalidator) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qintvalidator) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQIntValidator_QObject(parent QObject) QIntValidator {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qintvalidator = new(qintvalidator)
	qintvalidator.SetPointer(C.QIntValidator_New_QObject(parentPtr))
	qintvalidator.SetObjectName_String("QIntValidator_" + randomIdentifier())
	return qintvalidator
}

func NewQIntValidator_Int_Int_QObject(minimum int, maximum int, parent QObject) QIntValidator {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qintvalidator = new(qintvalidator)
	qintvalidator.SetPointer(C.QIntValidator_New_Int_Int_QObject(C.int(minimum), C.int(maximum), parentPtr))
	qintvalidator.SetObjectName_String("QIntValidator_" + randomIdentifier())
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
	} else {
		return int(C.QIntValidator_Bottom(p.Pointer()))
	}
}

func (p *qintvalidator) SetBottom_Int(in int) {
	if p.Pointer() != nil {
		C.QIntValidator_SetBottom_Int(p.Pointer(), C.int(in))
	}
}

func (p *qintvalidator) SetRange_Int_Int(bottom int, top int) {
	if p.Pointer() != nil {
		C.QIntValidator_SetRange_Int_Int(p.Pointer(), C.int(bottom), C.int(top))
	}
}

func (p *qintvalidator) SetTop_Int(in int) {
	if p.Pointer() != nil {
		C.QIntValidator_SetTop_Int(p.Pointer(), C.int(in))
	}
}

func (p *qintvalidator) Top() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QIntValidator_Top(p.Pointer()))
	}
}
