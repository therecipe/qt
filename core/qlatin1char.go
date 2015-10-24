package core

//#include "qlatin1char.h"
import "C"
import (
	"unsafe"
)

type QLatin1Char struct {
	ptr unsafe.Pointer
}

type QLatin1CharITF interface {
	QLatin1CharPTR() *QLatin1Char
}

func (p *QLatin1Char) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLatin1Char) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLatin1Char(ptr QLatin1CharITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLatin1CharPTR().Pointer()
	}
	return nil
}

func QLatin1CharFromPointer(ptr unsafe.Pointer) *QLatin1Char {
	var n = new(QLatin1Char)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLatin1Char) QLatin1CharPTR() *QLatin1Char {
	return ptr
}

func NewQLatin1Char(c string) *QLatin1Char {
	return QLatin1CharFromPointer(unsafe.Pointer(C.QLatin1Char_NewQLatin1Char(C.CString(c))))
}
