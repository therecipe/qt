package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNetworkDiskCache struct {
	QAbstractNetworkCache
}

type QNetworkDiskCache_ITF interface {
	QAbstractNetworkCache_ITF
	QNetworkDiskCache_PTR() *QNetworkDiskCache
}

func PointerFromQNetworkDiskCache(ptr QNetworkDiskCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDiskCache_PTR().Pointer()
	}
	return nil
}

func NewQNetworkDiskCacheFromPointer(ptr unsafe.Pointer) *QNetworkDiskCache {
	var n = new(QNetworkDiskCache)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkDiskCache_") {
		n.SetObjectName("QNetworkDiskCache_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return ptr
}

func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::QNetworkDiskCache")
		}
	}()

	return NewQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::cacheDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkDiskCache_CacheDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkDiskCache) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::prepare")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::setCacheDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), C.CString(cacheDir))
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::updateMetaData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkDiskCache::~QNetworkDiskCache")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
