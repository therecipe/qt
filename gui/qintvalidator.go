package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QIntValidator_") {
		n.SetObjectName("QIntValidator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIntValidator) QIntValidator_PTR() *QIntValidator {
	return ptr
}

func (ptr *QIntValidator) SetBottom(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::setBottom")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIntValidator_SetBottom(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QIntValidator) SetTop(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::setTop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIntValidator_SetTop(ptr.Pointer(), C.int(v))
	}
}

func NewQIntValidator(parent core.QObject_ITF) *QIntValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::QIntValidator")
		}
	}()

	return NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator(core.PointerFromQObject(parent)))
}

func NewQIntValidator2(minimum int, maximum int, parent core.QObject_ITF) *QIntValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::QIntValidator")
		}
	}()

	return NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator2(C.int(minimum), C.int(maximum), core.PointerFromQObject(parent)))
}

func (ptr *QIntValidator) Bottom() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::bottom")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIntValidator) SetRange(bottom int, top int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::setRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIntValidator_SetRange(ptr.Pointer(), C.int(bottom), C.int(top))
	}
}

func (ptr *QIntValidator) Top() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::top")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QIntValidator_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIntValidator) DestroyQIntValidator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIntValidator::~QIntValidator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIntValidator_DestroyQIntValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
