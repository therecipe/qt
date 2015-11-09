package gui

//#include "qsupportedwritingsystems.h"
import "C"
import (
	"unsafe"
)

type QSupportedWritingSystems struct {
	ptr unsafe.Pointer
}

type QSupportedWritingSystems_ITF interface {
	QSupportedWritingSystems_PTR() *QSupportedWritingSystems
}

func (p *QSupportedWritingSystems) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSupportedWritingSystems) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSupportedWritingSystems(ptr QSupportedWritingSystems_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSupportedWritingSystems_PTR().Pointer()
	}
	return nil
}

func NewQSupportedWritingSystemsFromPointer(ptr unsafe.Pointer) *QSupportedWritingSystems {
	var n = new(QSupportedWritingSystems)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSupportedWritingSystems) QSupportedWritingSystems_PTR() *QSupportedWritingSystems {
	return ptr
}
