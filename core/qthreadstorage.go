package core

//#include "qthreadstorage.h"
import "C"
import (
	"unsafe"
)

type QThreadStorage struct {
	ptr unsafe.Pointer
}

type QThreadStorageITF interface {
	QThreadStoragePTR() *QThreadStorage
}

func (p *QThreadStorage) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QThreadStorage) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQThreadStorage(ptr QThreadStorageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThreadStoragePTR().Pointer()
	}
	return nil
}

func QThreadStorageFromPointer(ptr unsafe.Pointer) *QThreadStorage {
	var n = new(QThreadStorage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QThreadStorage) QThreadStoragePTR() *QThreadStorage {
	return ptr
}
