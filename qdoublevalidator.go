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
	SetDecimals_Int(in int)
	SetNotation_Notation(No Notation)
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

func NewQDoubleValidator_QObject(parent QObject) QDoubleValidator {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdoublevalidator = new(qdoublevalidator)
	qdoublevalidator.SetPointer(C.QDoubleValidator_New_QObject(parentPtr))
	qdoublevalidator.SetObjectName_String("QDoubleValidator_" + randomIdentifier())
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
	} else {
		return int(C.QDoubleValidator_Decimals(p.Pointer()))
	}
}

func (p *qdoublevalidator) Notation() Notation {
	if p.Pointer() == nil {
		return 0
	} else {
		return Notation(C.QDoubleValidator_Notation(p.Pointer()))
	}
}

func (p *qdoublevalidator) SetDecimals_Int(in int) {
	if p.Pointer() != nil {
		C.QDoubleValidator_SetDecimals_Int(p.Pointer(), C.int(in))
	}
}

func (p *qdoublevalidator) SetNotation_Notation(No Notation) {
	if p.Pointer() != nil {
		C.QDoubleValidator_SetNotation_Notation(p.Pointer(), C.int(No))
	}
}
