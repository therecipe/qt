package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCache_ITF interface {
	core.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func PointerFromQAbstractNetworkCache(ptr QAbstractNetworkCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func NewQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) *QAbstractNetworkCache {
	var n = new(QAbstractNetworkCache)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractNetworkCache_") {
		n.SetObjectName("QAbstractNetworkCache_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache {
	return ptr
}

func (ptr *QAbstractNetworkCache) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractNetworkCache::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(ptr.Pointer())
	}
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractNetworkCache::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractNetworkCache::prepare")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractNetworkCache::updateMetaData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractNetworkCache::~QAbstractNetworkCache")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
