package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGenericArgument struct {
	ptr unsafe.Pointer
}

type QGenericArgument_ITF interface {
	QGenericArgument_PTR() *QGenericArgument
}

func (p *QGenericArgument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGenericArgument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGenericArgument(ptr QGenericArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericArgument_PTR().Pointer()
	}
	return nil
}

func NewQGenericArgumentFromPointer(ptr unsafe.Pointer) *QGenericArgument {
	var n = new(QGenericArgument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericArgument) QGenericArgument_PTR() *QGenericArgument {
	return ptr
}

func (ptr *QGenericArgument) Data() unsafe.Pointer {
	defer qt.Recovering("QGenericArgument::data")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QGenericArgument_Data(ptr.Pointer()))
	}
	return nil
}
