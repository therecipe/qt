package bluetooth

//#include "qlowenergydescriptor.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLowEnergyDescriptor struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptor_ITF interface {
	QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor
}

func (p *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLowEnergyDescriptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLowEnergyDescriptor(ptr QLowEnergyDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = new(QLowEnergyDescriptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLowEnergyDescriptor) QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor {
	return ptr
}

func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	return NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor())
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {
	return NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(PointerFromQLowEnergyDescriptor(other)))
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptor) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyDescriptor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QLowEnergyDescriptor_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(ptr.Pointer())
	}
}
