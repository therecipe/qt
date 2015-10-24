package gui

//#include "qpixmapcache.h"
import "C"
import (
	"unsafe"
)

type QPixmapCache struct {
	ptr unsafe.Pointer
}

type QPixmapCacheITF interface {
	QPixmapCachePTR() *QPixmapCache
}

func (p *QPixmapCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPixmapCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPixmapCache(ptr QPixmapCacheITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmapCachePTR().Pointer()
	}
	return nil
}

func QPixmapCacheFromPointer(ptr unsafe.Pointer) *QPixmapCache {
	var n = new(QPixmapCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPixmapCache) QPixmapCachePTR() *QPixmapCache {
	return ptr
}

func QPixmapCache_CacheLimit() int {
	return int(C.QPixmapCache_QPixmapCache_CacheLimit())
}

func QPixmapCache_Clear() {
	C.QPixmapCache_QPixmapCache_Clear()
}

func QPixmapCache_SetCacheLimit(n int) {
	C.QPixmapCache_QPixmapCache_SetCacheLimit(C.int(n))
}
