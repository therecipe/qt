package core

//#include "qcache.h"
import "C"
import (
	"unsafe"
)

type QCache struct {
	ptr unsafe.Pointer
}

type QCache_ITF interface {
	QCache_PTR() *QCache
}

func (p *QCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCache(ptr QCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCache_PTR().Pointer()
	}
	return nil
}

func NewQCacheFromPointer(ptr unsafe.Pointer) *QCache {
	var n = new(QCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCache) QCache_PTR() *QCache {
	return ptr
}
