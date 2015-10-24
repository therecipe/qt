package core

//#include "qgenericargument.h"
import "C"
import (
	"unsafe"
)

type QGenericArgument struct {
	ptr unsafe.Pointer
}

type QGenericArgumentITF interface {
	QGenericArgumentPTR() *QGenericArgument
}

func (p *QGenericArgument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGenericArgument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGenericArgument(ptr QGenericArgumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericArgumentPTR().Pointer()
	}
	return nil
}

func QGenericArgumentFromPointer(ptr unsafe.Pointer) *QGenericArgument {
	var n = new(QGenericArgument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericArgument) QGenericArgumentPTR() *QGenericArgument {
	return ptr
}

func (ptr *QGenericArgument) Data() {
	if ptr.Pointer() != nil {
		C.QGenericArgument_Data(C.QtObjectPtr(ptr.Pointer()))
	}
}
