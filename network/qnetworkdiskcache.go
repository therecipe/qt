package network

//#include "qnetworkdiskcache.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QNetworkDiskCache_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return ptr
}

func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	return NewQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkDiskCache_CacheDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkDiskCache) Clear() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), C.CString(cacheDir))
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
