package gui

//#include "qdoublevalidator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDoubleValidator struct {
	QValidator
}

type QDoubleValidatorITF interface {
	QValidatorITF
	QDoubleValidatorPTR() *QDoubleValidator
}

func PointerFromQDoubleValidator(ptr QDoubleValidatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDoubleValidatorPTR().Pointer()
	}
	return nil
}

func QDoubleValidatorFromPointer(ptr unsafe.Pointer) *QDoubleValidator {
	var n = new(QDoubleValidator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDoubleValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDoubleValidator) QDoubleValidatorPTR() *QDoubleValidator {
	return ptr
}

//QDoubleValidator::Notation
type QDoubleValidator__Notation int

var (
	QDoubleValidator__StandardNotation   = QDoubleValidator__Notation(0)
	QDoubleValidator__ScientificNotation = QDoubleValidator__Notation(1)
)

func (ptr *QDoubleValidator) Notation() QDoubleValidator__Notation {
	if ptr.Pointer() != nil {
		return QDoubleValidator__Notation(C.QDoubleValidator_Notation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDoubleValidator) SetDecimals(v int) {
	if ptr.Pointer() != nil {
		C.QDoubleValidator_SetDecimals(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QDoubleValidator) SetNotation(v QDoubleValidator__Notation) {
	if ptr.Pointer() != nil {
		C.QDoubleValidator_SetNotation(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func NewQDoubleValidator(parent core.QObjectITF) *QDoubleValidator {
	return QDoubleValidatorFromPointer(unsafe.Pointer(C.QDoubleValidator_NewQDoubleValidator(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDoubleValidator) Decimals() int {
	if ptr.Pointer() != nil {
		return int(C.QDoubleValidator_Decimals(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDoubleValidator) DestroyQDoubleValidator() {
	if ptr.Pointer() != nil {
		C.QDoubleValidator_DestroyQDoubleValidator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
