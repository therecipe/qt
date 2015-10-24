package gui

//#include "qregularexpressionvalidator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRegularExpressionValidator struct {
	QValidator
}

type QRegularExpressionValidatorITF interface {
	QValidatorITF
	QRegularExpressionValidatorPTR() *QRegularExpressionValidator
}

func PointerFromQRegularExpressionValidator(ptr QRegularExpressionValidatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionValidatorPTR().Pointer()
	}
	return nil
}

func QRegularExpressionValidatorFromPointer(ptr unsafe.Pointer) *QRegularExpressionValidator {
	var n = new(QRegularExpressionValidator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRegularExpressionValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRegularExpressionValidator) QRegularExpressionValidatorPTR() *QRegularExpressionValidator {
	return ptr
}

func (ptr *QRegularExpressionValidator) SetRegularExpression(re core.QRegularExpressionITF) {
	if ptr.Pointer() != nil {
		C.QRegularExpressionValidator_SetRegularExpression(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRegularExpression(re)))
	}
}

func NewQRegularExpressionValidator(parent core.QObjectITF) *QRegularExpressionValidator {
	return QRegularExpressionValidatorFromPointer(unsafe.Pointer(C.QRegularExpressionValidator_NewQRegularExpressionValidator(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQRegularExpressionValidator2(re core.QRegularExpressionITF, parent core.QObjectITF) *QRegularExpressionValidator {
	return QRegularExpressionValidatorFromPointer(unsafe.Pointer(C.QRegularExpressionValidator_NewQRegularExpressionValidator2(C.QtObjectPtr(core.PointerFromQRegularExpression(re)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QRegularExpressionValidator) DestroyQRegularExpressionValidator() {
	if ptr.Pointer() != nil {
		C.QRegularExpressionValidator_DestroyQRegularExpressionValidator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
