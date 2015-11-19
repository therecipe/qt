package gui

//#include "qintvalidator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIntValidator struct {
	QValidator
}

type QIntValidator_ITF interface {
	QValidator_ITF
	QIntValidator_PTR() *QIntValidator
}

func PointerFromQIntValidator(ptr QIntValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIntValidator_PTR().Pointer()
	}
	return nil
}

func NewQIntValidatorFromPointer(ptr unsafe.Pointer) *QIntValidator {
	var n = new(QIntValidator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIntValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIntValidator) QIntValidator_PTR() *QIntValidator {
	return ptr
}

func (ptr *QIntValidator) SetBottom(v int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetBottom(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QIntValidator) SetTop(v int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetTop(ptr.Pointer(), C.int(v))
	}
}

func NewQIntValidator(parent core.QObject_ITF) *QIntValidator {
	return NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator(core.PointerFromQObject(parent)))
}

func NewQIntValidator2(minimum int, maximum int, parent core.QObject_ITF) *QIntValidator {
	return NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator2(C.int(minimum), C.int(maximum), core.PointerFromQObject(parent)))
}

func (ptr *QIntValidator) Bottom() int {
	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIntValidator) SetRange(bottom int, top int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetRange(ptr.Pointer(), C.int(bottom), C.int(top))
	}
}

func (ptr *QIntValidator) Top() int {
	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIntValidator) DestroyQIntValidator() {
	if ptr.Pointer() != nil {
		C.QIntValidator_DestroyQIntValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
