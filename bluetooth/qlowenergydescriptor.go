package bluetooth

//#include "qlowenergydescriptor.h"
import "C"
import (
	"unsafe"
)

type QLowEnergyDescriptor struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptorITF interface {
	QLowEnergyDescriptorPTR() *QLowEnergyDescriptor
}

func (p *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLowEnergyDescriptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLowEnergyDescriptor(ptr QLowEnergyDescriptorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptorPTR().Pointer()
	}
	return nil
}

func QLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = new(QLowEnergyDescriptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLowEnergyDescriptor) QLowEnergyDescriptorPTR() *QLowEnergyDescriptor {
	return ptr
}

func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	return QLowEnergyDescriptorFromPointer(unsafe.Pointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor()))
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptorITF) *QLowEnergyDescriptor {
	return QLowEnergyDescriptorFromPointer(unsafe.Pointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(C.QtObjectPtr(PointerFromQLowEnergyDescriptor(other)))))
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptor_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptor) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyDescriptor_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(C.QtObjectPtr(ptr.Pointer()))
	}
}
