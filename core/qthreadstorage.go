package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QThreadStorage struct {
	ptr unsafe.Pointer
}

type QThreadStorage_ITF interface {
	QThreadStorage_PTR() *QThreadStorage
}

func (p *QThreadStorage) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QThreadStorage) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQThreadStorage(ptr QThreadStorage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThreadStorage_PTR().Pointer()
	}
	return nil
}

func NewQThreadStorageFromPointer(ptr unsafe.Pointer) *QThreadStorage {
	var n = new(QThreadStorage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QThreadStorage) QThreadStorage_PTR() *QThreadStorage {
	return ptr
}
