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

type QValidator_ITF interface {
	core.QObject_ITF
	QValidator_PTR() *QValidator
}

func PointerFromQValidator(ptr QValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValidator_PTR().Pointer()
	}
	return nil
}

func NewQValidatorFromPointer(ptr unsafe.Pointer) *QValidator {
	var n = new(QValidator)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QValidator) QValidator_PTR() *QValidator {
	return ptr
}

//QValidator::State
type QValidator__State int64

const (
	QValidator__Invalid      = QValidator__State(0)
	QValidator__Intermediate = QValidator__State(1)
	QValidator__Acceptable   = QValidator__State(2)
)

func (ptr *QValidator) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QValidator_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QValidator) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QValidator_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQValidatorChanged
func callbackQValidatorChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "changed").(func())()
}

func (ptr *QValidator) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QValidator) DestroyQValidator() {
	if ptr.Pointer() != nil {
		C.QValidator_DestroyQValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
