package gui

//#include "qregexpvalidator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRegExpValidator struct {
	QValidator
}

type QRegExpValidatorITF interface {
	QValidatorITF
	QRegExpValidatorPTR() *QRegExpValidator
}

func PointerFromQRegExpValidator(ptr QRegExpValidatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegExpValidatorPTR().Pointer()
	}
	return nil
}

func QRegExpValidatorFromPointer(ptr unsafe.Pointer) *QRegExpValidator {
	var n = new(QRegExpValidator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRegExpValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRegExpValidator) QRegExpValidatorPTR() *QRegExpValidator {
	return ptr
}

func (ptr *QRegExpValidator) SetRegExp(rx core.QRegExpITF) {
	if ptr.Pointer() != nil {
		C.QRegExpValidator_SetRegExp(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRegExp(rx)))
	}
}

func NewQRegExpValidator(parent core.QObjectITF) *QRegExpValidator {
	return QRegExpValidatorFromPointer(unsafe.Pointer(C.QRegExpValidator_NewQRegExpValidator(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQRegExpValidator2(rx core.QRegExpITF, parent core.QObjectITF) *QRegExpValidator {
	return QRegExpValidatorFromPointer(unsafe.Pointer(C.QRegExpValidator_NewQRegExpValidator2(C.QtObjectPtr(core.PointerFromQRegExp(rx)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QRegExpValidator) DestroyQRegExpValidator() {
	if ptr.Pointer() != nil {
		C.QRegExpValidator_DestroyQRegExpValidator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
