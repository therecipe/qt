package core

//#include "qlatin1string.h"
import "C"
import (
	"unsafe"
)

type QLatin1String struct {
	ptr unsafe.Pointer
}

type QLatin1StringITF interface {
	QLatin1StringPTR() *QLatin1String
}

func (p *QLatin1String) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLatin1String) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLatin1String(ptr QLatin1StringITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLatin1StringPTR().Pointer()
	}
	return nil
}

func QLatin1StringFromPointer(ptr unsafe.Pointer) *QLatin1String {
	var n = new(QLatin1String)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLatin1String) QLatin1StringPTR() *QLatin1String {
	return ptr
}

func NewQLatin1String3(str QByteArrayITF) *QLatin1String {
	return QLatin1StringFromPointer(unsafe.Pointer(C.QLatin1String_NewQLatin1String3(C.QtObjectPtr(PointerFromQByteArray(str)))))
}

func NewQLatin1String(str string) *QLatin1String {
	return QLatin1StringFromPointer(unsafe.Pointer(C.QLatin1String_NewQLatin1String(C.CString(str))))
}

func NewQLatin1String2(str string, size int) *QLatin1String {
	return QLatin1StringFromPointer(unsafe.Pointer(C.QLatin1String_NewQLatin1String2(C.CString(str), C.int(size))))
}

func (ptr *QLatin1String) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QLatin1String_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
