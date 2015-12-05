package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QContiguousCache struct {
	ptr unsafe.Pointer
}

type QContiguousCache_ITF interface {
	QContiguousCache_PTR() *QContiguousCache
}

func (p *QContiguousCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QContiguousCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQContiguousCache(ptr QContiguousCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QContiguousCache_PTR().Pointer()
	}
	return nil
}

func NewQContiguousCacheFromPointer(ptr unsafe.Pointer) *QContiguousCache {
	var n = new(QContiguousCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QContiguousCache) QContiguousCache_PTR() *QContiguousCache {
	return ptr
}
