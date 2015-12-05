package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QRegExpValidator struct {
	QValidator
}

type QRegExpValidator_ITF interface {
	QValidator_ITF
	QRegExpValidator_PTR() *QRegExpValidator
}

func PointerFromQRegExpValidator(ptr QRegExpValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegExpValidator_PTR().Pointer()
	}
	return nil
}

func NewQRegExpValidatorFromPointer(ptr unsafe.Pointer) *QRegExpValidator {
	var n = new(QRegExpValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRegExpValidator_") {
		n.SetObjectName("QRegExpValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRegExpValidator) QRegExpValidator_PTR() *QRegExpValidator {
	return ptr
}

func (ptr *QRegExpValidator) SetRegExp(rx core.QRegExp_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExpValidator::setRegExp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExpValidator_SetRegExp(ptr.Pointer(), core.PointerFromQRegExp(rx))
	}
}

func NewQRegExpValidator(parent core.QObject_ITF) *QRegExpValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExpValidator::QRegExpValidator")
		}
	}()

	return NewQRegExpValidatorFromPointer(C.QRegExpValidator_NewQRegExpValidator(core.PointerFromQObject(parent)))
}

func NewQRegExpValidator2(rx core.QRegExp_ITF, parent core.QObject_ITF) *QRegExpValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExpValidator::QRegExpValidator")
		}
	}()

	return NewQRegExpValidatorFromPointer(C.QRegExpValidator_NewQRegExpValidator2(core.PointerFromQRegExp(rx), core.PointerFromQObject(parent)))
}

func (ptr *QRegExpValidator) RegExp() *core.QRegExp {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExpValidator::regExp")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQRegExpFromPointer(C.QRegExpValidator_RegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRegExpValidator) DestroyQRegExpValidator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExpValidator::~QRegExpValidator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExpValidator_DestroyQRegExpValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
