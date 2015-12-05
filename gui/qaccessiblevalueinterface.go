package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAccessibleValueInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleValueInterface_ITF interface {
	QAccessibleValueInterface_PTR() *QAccessibleValueInterface
}

func (p *QAccessibleValueInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleValueInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleValueInterface(ptr QAccessibleValueInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleValueInterface_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleValueInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleValueInterface {
	var n = new(QAccessibleValueInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleValueInterface) QAccessibleValueInterface_PTR() *QAccessibleValueInterface {
	return ptr
}

func (ptr *QAccessibleValueInterface) CurrentValue() *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleValueInterface::currentValue")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAccessibleValueInterface_CurrentValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleValueInterface) MaximumValue() *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleValueInterface::maximumValue")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAccessibleValueInterface_MaximumValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleValueInterface) MinimumStepSize() *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleValueInterface::minimumStepSize")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAccessibleValueInterface_MinimumStepSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleValueInterface) MinimumValue() *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleValueInterface::minimumValue")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAccessibleValueInterface_MinimumValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleValueInterface) SetCurrentValue(value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleValueInterface::setCurrentValue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleValueInterface_SetCurrentValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QAccessibleValueInterface) DestroyQAccessibleValueInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleValueInterface::~QAccessibleValueInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleValueInterface_DestroyQAccessibleValueInterface(ptr.Pointer())
	}
}
