package core

//#include "qsharedpointer.h"
import "C"
import (
	"unsafe"
)

type QSharedPointer struct {
	ptr unsafe.Pointer
}

type QSharedPointerITF interface {
	QSharedPointerPTR() *QSharedPointer
}

func (p *QSharedPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSharedPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSharedPointer(ptr QSharedPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedPointerPTR().Pointer()
	}
	return nil
}

func QSharedPointerFromPointer(ptr unsafe.Pointer) *QSharedPointer {
	var n = new(QSharedPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSharedPointer) QSharedPointerPTR() *QSharedPointer {
	return ptr
}
