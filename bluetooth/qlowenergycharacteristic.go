package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLowEnergyCharacteristic struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristic_ITF interface {
	QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic
}

func (p *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLowEnergyCharacteristic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLowEnergyCharacteristic(ptr QLowEnergyCharacteristic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristic_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = new(QLowEnergyCharacteristic)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLowEnergyCharacteristic) QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic {
	return ptr
}

//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int64

const (
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
	defer qt.Recovering("QLowEnergyCharacteristic::QLowEnergyCharacteristic")

	return NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic())
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {
	defer qt.Recovering("QLowEnergyCharacteristic::QLowEnergyCharacteristic")

	return NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(PointerFromQLowEnergyCharacteristic(other)))
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	defer qt.Recovering("QLowEnergyCharacteristic::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristic_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	defer qt.Recovering("QLowEnergyCharacteristic::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyCharacteristic_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	defer qt.Recovering("QLowEnergyCharacteristic::properties")

	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) Value() *core.QByteArray {
	defer qt.Recovering("QLowEnergyCharacteristic::value")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristic_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	defer qt.Recovering("QLowEnergyCharacteristic::~QLowEnergyCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(ptr.Pointer())
	}
}
