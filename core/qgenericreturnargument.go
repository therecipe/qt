package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QGenericReturnArgument struct {
	QGenericArgument
}

type QGenericReturnArgument_ITF interface {
	QGenericArgument_ITF
	QGenericReturnArgument_PTR() *QGenericReturnArgument
}

func PointerFromQGenericReturnArgument(ptr QGenericReturnArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericReturnArgument_PTR().Pointer()
	}
	return nil
}

func NewQGenericReturnArgumentFromPointer(ptr unsafe.Pointer) *QGenericReturnArgument {
	var n = new(QGenericReturnArgument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericReturnArgument) QGenericReturnArgument_PTR() *QGenericReturnArgument {
	return ptr
}
