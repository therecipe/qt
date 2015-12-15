package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDoubleValidator struct {
	QValidator
}

type QDoubleValidator_ITF interface {
	QValidator_ITF
	QDoubleValidator_PTR() *QDoubleValidator
}

func PointerFromQDoubleValidator(ptr QDoubleValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDoubleValidator_PTR().Pointer()
	}
	return nil
}

func NewQDoubleValidatorFromPointer(ptr unsafe.Pointer) *QDoubleValidator {
	var n = new(QDoubleValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDoubleValidator_") {
		n.SetObjectName("QDoubleValidator_" + qt.Identifier())
	}
	return n
}

func (ptr *QDoubleValidator) QDoubleValidator_PTR() *QDoubleValidator {
	return ptr
}

//QDoubleValidator::Notation
type QDoubleValidator__Notation int64

const (
	QDoubleValidator__StandardNotation   = QDoubleValidator__Notation(0)
	QDoubleValidator__ScientificNotation = QDoubleValidator__Notation(1)
)

func (ptr *QDoubleValidator) Notation() QDoubleValidator__Notation {
	defer qt.Recovering("QDoubleValidator::notation")

	if ptr.Pointer() != nil {
		return QDoubleValidator__Notation(C.QDoubleValidator_Notation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleValidator) SetDecimals(v int) {
	defer qt.Recovering("QDoubleValidator::setDecimals")

	if ptr.Pointer() != nil {
		C.QDoubleValidator_SetDecimals(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QDoubleValidator) SetNotation(v QDoubleValidator__Notation) {
	defer qt.Recovering("QDoubleValidator::setNotation")

	if ptr.Pointer() != nil {
		C.QDoubleValidator_SetNotation(ptr.Pointer(), C.int(v))
	}
}

func NewQDoubleValidator(parent core.QObject_ITF) *QDoubleValidator {
	defer qt.Recovering("QDoubleValidator::QDoubleValidator")

	return NewQDoubleValidatorFromPointer(C.QDoubleValidator_NewQDoubleValidator(core.PointerFromQObject(parent)))
}

func (ptr *QDoubleValidator) Decimals() int {
	defer qt.Recovering("QDoubleValidator::decimals")

	if ptr.Pointer() != nil {
		return int(C.QDoubleValidator_Decimals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleValidator) DestroyQDoubleValidator() {
	defer qt.Recovering("QDoubleValidator::~QDoubleValidator")

	if ptr.Pointer() != nil {
		C.QDoubleValidator_DestroyQDoubleValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
