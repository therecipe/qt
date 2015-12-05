package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QLatin1Char struct {
	ptr unsafe.Pointer
}

type QLatin1Char_ITF interface {
	QLatin1Char_PTR() *QLatin1Char
}

func (p *QLatin1Char) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLatin1Char) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLatin1Char(ptr QLatin1Char_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLatin1Char_PTR().Pointer()
	}
	return nil
}

func NewQLatin1CharFromPointer(ptr unsafe.Pointer) *QLatin1Char {
	var n = new(QLatin1Char)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLatin1Char) QLatin1Char_PTR() *QLatin1Char {
	return ptr
}

func NewQLatin1Char(c string) *QLatin1Char {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLatin1Char::QLatin1Char")
		}
	}()

	return NewQLatin1CharFromPointer(C.QLatin1Char_NewQLatin1Char(C.CString(c)))
}
