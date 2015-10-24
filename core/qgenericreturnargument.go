package core

//#include "qgenericreturnargument.h"
import "C"
import (
	"unsafe"
)

type QGenericReturnArgument struct {
	QGenericArgument
}

type QGenericReturnArgumentITF interface {
	QGenericArgumentITF
	QGenericReturnArgumentPTR() *QGenericReturnArgument
}

func PointerFromQGenericReturnArgument(ptr QGenericReturnArgumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericReturnArgumentPTR().Pointer()
	}
	return nil
}

func QGenericReturnArgumentFromPointer(ptr unsafe.Pointer) *QGenericReturnArgument {
	var n = new(QGenericReturnArgument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericReturnArgument) QGenericReturnArgumentPTR() *QGenericReturnArgument {
	return ptr
}
