package gui

//#include "qvalidator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QValidator struct {
	core.QObject
}

type QValidatorITF interface {
	core.QObjectITF
	QValidatorPTR() *QValidator
}

func PointerFromQValidator(ptr QValidatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValidatorPTR().Pointer()
	}
	return nil
}

func QValidatorFromPointer(ptr unsafe.Pointer) *QValidator {
	var n = new(QValidator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QValidator) QValidatorPTR() *QValidator {
	return ptr
}

//QValidator::State
type QValidator__State int

var (
	QValidator__Invalid      = QValidator__State(0)
	QValidator__Intermediate = QValidator__State(1)
	QValidator__Acceptable   = QValidator__State(2)
)

func (ptr *QValidator) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QValidator_ConnectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QValidator) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QValidator_DisconnectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQValidatorChanged
func callbackQValidatorChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "changed").(func())()
}

func (ptr *QValidator) SetLocale(locale core.QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QValidator_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLocale(locale)))
	}
}

func (ptr *QValidator) DestroyQValidator() {
	if ptr.Pointer() != nil {
		C.QValidator_DestroyQValidator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
