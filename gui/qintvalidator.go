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

type QIntValidatorITF interface {
	QValidatorITF
	QIntValidatorPTR() *QIntValidator
}

func PointerFromQIntValidator(ptr QIntValidatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIntValidatorPTR().Pointer()
	}
	return nil
}

func QIntValidatorFromPointer(ptr unsafe.Pointer) *QIntValidator {
	var n = new(QIntValidator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIntValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIntValidator) QIntValidatorPTR() *QIntValidator {
	return ptr
}

func (ptr *QIntValidator) SetBottom(v int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetBottom(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QIntValidator) SetTop(v int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetTop(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func NewQIntValidator(parent core.QObjectITF) *QIntValidator {
	return QIntValidatorFromPointer(unsafe.Pointer(C.QIntValidator_NewQIntValidator(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQIntValidator2(minimum int, maximum int, parent core.QObjectITF) *QIntValidator {
	return QIntValidatorFromPointer(unsafe.Pointer(C.QIntValidator_NewQIntValidator2(C.int(minimum), C.int(maximum), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QIntValidator) Bottom() int {
	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Bottom(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QIntValidator) SetRange(bottom int, top int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetRange(C.QtObjectPtr(ptr.Pointer()), C.int(bottom), C.int(top))
	}
}

func (ptr *QIntValidator) Top() int {
	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Top(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QIntValidator) DestroyQIntValidator() {
	if ptr.Pointer() != nil {
		C.QIntValidator_DestroyQIntValidator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
