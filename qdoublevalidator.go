package qt

//#include "qdoublevalidator.h"
import "C"

type qdoublevalidator struct {
	qvalidator
}

type QDoubleValidator interface {
	QValidator
	Decimals() int
	Notation() Notation
	SetDecimals(in int)
	SetNotation(No Notation)
}

func (p *qdoublevalidator) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qdoublevalidator) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//Notation
type Notation int

var (
	STANDARDNOTATION   = Notation(C.QDoubleValidator_StandardNotation())
	SCIENTIFICNOTATION = Notation(C.QDoubleValidator_ScientificNotation())
)

func NewQDoubleValidator(parent QObject) QDoubleValidator {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdoublevalidator = new(qdoublevalidator)
	qdoublevalidator.SetPointer(C.QDoubleValidator_New_QObject(parentPtr))
	qdoublevalidator.SetObjectName("QDoubleValidator_" + randomIdentifier())
	return qdoublevalidator
}

func (p *qdoublevalidator) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QDoubleValidator_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qdoublevalidator) Decimals() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QDoubleValidator_Decimals(p.Pointer()))
}

func (p *qdoublevalidator) Notation() Notation {
	if p.Pointer() == nil {
		return 0
	}
	return Notation(C.QDoubleValidator_Notation(p.Pointer()))
}

func (p *qdoublevalidator) SetDecimals(in int) {
	if p.Pointer() != nil {
		C.QDoubleValidator_SetDecimals_Int(p.Pointer(), C.int(in))
	}
}

func (p *qdoublevalidator) SetNotation(No Notation) {
	if p.Pointer() != nil {
		C.QDoubleValidator_SetNotation_Notation(p.Pointer(), C.int(No))
	}
}
