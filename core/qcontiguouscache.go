package core

//#include "qcontiguouscache.h"
import "C"
import (
	"unsafe"
)

type QContiguousCache struct {
	ptr unsafe.Pointer
}

type QContiguousCacheITF interface {
	QContiguousCachePTR() *QContiguousCache
}

func (p *QContiguousCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QContiguousCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQContiguousCache(ptr QContiguousCacheITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QContiguousCachePTR().Pointer()
	}
	return nil
}

func QContiguousCacheFromPointer(ptr unsafe.Pointer) *QContiguousCache {
	var n = new(QContiguousCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QContiguousCache) QContiguousCachePTR() *QContiguousCache {
	return ptr
}
