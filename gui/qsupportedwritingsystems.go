package gui

//#include "qsupportedwritingsystems.h"
import "C"
import (
	"unsafe"
)

type QSupportedWritingSystems struct {
	ptr unsafe.Pointer
}

type QSupportedWritingSystemsITF interface {
	QSupportedWritingSystemsPTR() *QSupportedWritingSystems
}

func (p *QSupportedWritingSystems) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSupportedWritingSystems) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSupportedWritingSystems(ptr QSupportedWritingSystemsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSupportedWritingSystemsPTR().Pointer()
	}
	return nil
}

func QSupportedWritingSystemsFromPointer(ptr unsafe.Pointer) *QSupportedWritingSystems {
	var n = new(QSupportedWritingSystems)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSupportedWritingSystems) QSupportedWritingSystemsPTR() *QSupportedWritingSystems {
	return ptr
}
