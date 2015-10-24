package bluetooth

//#include "qlowenergycharacteristic.h"
import "C"
import (
	"unsafe"
)

type QLowEnergyCharacteristic struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristicITF interface {
	QLowEnergyCharacteristicPTR() *QLowEnergyCharacteristic
}

func (p *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLowEnergyCharacteristic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLowEnergyCharacteristic(ptr QLowEnergyCharacteristicITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristicPTR().Pointer()
	}
	return nil
}

func QLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = new(QLowEnergyCharacteristic)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLowEnergyCharacteristic) QLowEnergyCharacteristicPTR() *QLowEnergyCharacteristic {
	return ptr
}

//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int

var (
	QLowEnergyCharacteristic__Unknown          = QLowEnergyCharacteristic__PropertyType(0x00)
	QLowEnergyCharacteristic__Broadcasting     = QLowEnergyCharacteristic__PropertyType(0x01)
	QLowEnergyCharacteristic__Read             = QLowEnergyCharacteristic__PropertyType(0x02)
	QLowEnergyCharacteristic__WriteNoResponse  = QLowEnergyCharacteristic__PropertyType(0x04)
	QLowEnergyCharacteristic__Write            = QLowEnergyCharacteristic__PropertyType(0x08)
	QLowEnergyCharacteristic__Notify           = QLowEnergyCharacteristic__PropertyType(0x10)
	QLowEnergyCharacteristic__Indicate         = QLowEnergyCharacteristic__PropertyType(0x20)
	QLowEnergyCharacteristic__WriteSigned      = QLowEnergyCharacteristic__PropertyType(0x40)
	QLowEnergyCharacteristic__ExtendedProperty = QLowEnergyCharacteristic__PropertyType(0x80)
)

func NewQLowEnergyCharacteristic() *QLowEnergyCharacteristic {
	return QLowEnergyCharacteristicFromPointer(unsafe.Pointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic()))
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristicITF) *QLowEnergyCharacteristic {
	return QLowEnergyCharacteristicFromPointer(unsafe.Pointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(C.QtObjectPtr(PointerFromQLowEnergyCharacteristic(other)))))
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristic_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyCharacteristic_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(C.QtObjectPtr(ptr.Pointer()))
	}
}
