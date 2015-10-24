package gui

//#include "qaccessiblevalueinterface.h"
import "C"
import (
	"unsafe"
)

type QAccessibleValueInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleValueInterfaceITF interface {
	QAccessibleValueInterfacePTR() *QAccessibleValueInterface
}

func (p *QAccessibleValueInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleValueInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleValueInterface(ptr QAccessibleValueInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleValueInterfacePTR().Pointer()
	}
	return nil
}

func QAccessibleValueInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleValueInterface {
	var n = new(QAccessibleValueInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleValueInterface) QAccessibleValueInterfacePTR() *QAccessibleValueInterface {
	return ptr
}

func (ptr *QAccessibleValueInterface) CurrentValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleValueInterface_CurrentValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAccessibleValueInterface) MaximumValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleValueInterface_MaximumValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAccessibleValueInterface) MinimumStepSize() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleValueInterface_MinimumStepSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAccessibleValueInterface) MinimumValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleValueInterface_MinimumValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAccessibleValueInterface) SetCurrentValue(value string) {
	if ptr.Pointer() != nil {
		C.QAccessibleValueInterface_SetCurrentValue(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QAccessibleValueInterface) DestroyQAccessibleValueInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleValueInterface_DestroyQAccessibleValueInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
