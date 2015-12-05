package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QLatin1String struct {
	ptr unsafe.Pointer
}

type QLatin1String_ITF interface {
	QLatin1String_PTR() *QLatin1String
}

func (p *QLatin1String) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLatin1String) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLatin1String(ptr QLatin1String_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLatin1String_PTR().Pointer()
	}
	return nil
}

func NewQLatin1StringFromPointer(ptr unsafe.Pointer) *QLatin1String {
	var n = new(QLatin1String)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLatin1String) QLatin1String_PTR() *QLatin1String {
	return ptr
}

func NewQLatin1String3(str QByteArray_ITF) *QLatin1String {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLatin1String::QLatin1String")
		}
	}()

	return NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String3(PointerFromQByteArray(str)))
}

func NewQLatin1String(str string) *QLatin1String {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLatin1String::QLatin1String")
		}
	}()

	return NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String(C.CString(str)))
}

func NewQLatin1String2(str string, size int) *QLatin1String {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLatin1String::QLatin1String")
		}
	}()

	return NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String2(C.CString(str), C.int(size)))
}

func (ptr *QLatin1String) Size() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLatin1String::size")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLatin1String_Size(ptr.Pointer()))
	}
	return 0
}
