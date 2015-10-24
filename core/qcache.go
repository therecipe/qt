package core

//#include "qcache.h"
import "C"
import (
	"unsafe"
)

type QCache struct {
	ptr unsafe.Pointer
}

type QCacheITF interface {
	QCachePTR() *QCache
}

func (p *QCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCache(ptr QCacheITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCachePTR().Pointer()
	}
	return nil
}

func QCacheFromPointer(ptr unsafe.Pointer) *QCache {
	var n = new(QCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCache) QCachePTR() *QCache {
	return ptr
}
