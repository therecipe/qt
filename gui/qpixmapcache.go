package gui

//#include "gui.h"
import "C"
import (
	"log"
	"unsafe"
)

type QPixmapCache struct {
	ptr unsafe.Pointer
}

type QPixmapCache_ITF interface {
	QPixmapCache_PTR() *QPixmapCache
}

func (p *QPixmapCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPixmapCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPixmapCache(ptr QPixmapCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmapCache_PTR().Pointer()
	}
	return nil
}

func NewQPixmapCacheFromPointer(ptr unsafe.Pointer) *QPixmapCache {
	var n = new(QPixmapCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPixmapCache) QPixmapCache_PTR() *QPixmapCache {
	return ptr
}

func QPixmapCache_CacheLimit() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixmapCache::cacheLimit")
		}
	}()

	return int(C.QPixmapCache_QPixmapCache_CacheLimit())
}

func QPixmapCache_Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixmapCache::clear")
		}
	}()

	C.QPixmapCache_QPixmapCache_Clear()
}

func QPixmapCache_SetCacheLimit(n int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixmapCache::setCacheLimit")
		}
	}()

	C.QPixmapCache_QPixmapCache_SetCacheLimit(C.int(n))
}
