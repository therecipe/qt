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

type QNetworkDiskCacheITF interface {
	QAbstractNetworkCacheITF
	QNetworkDiskCachePTR() *QNetworkDiskCache
}

func PointerFromQNetworkDiskCache(ptr QNetworkDiskCacheITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDiskCachePTR().Pointer()
	}
	return nil
}

func QNetworkDiskCacheFromPointer(ptr unsafe.Pointer) *QNetworkDiskCache {
	var n = new(QNetworkDiskCache)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkDiskCache_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkDiskCache) QNetworkDiskCachePTR() *QNetworkDiskCache {
	return ptr
}

func NewQNetworkDiskCache(parent core.QObjectITF) *QNetworkDiskCache {
	return QNetworkDiskCacheFromPointer(unsafe.Pointer(C.QNetworkDiskCache_NewQNetworkDiskCache(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkDiskCache_CacheDirectory(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkDiskCache) Clear() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkDiskCache) Data(url string) *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QNetworkDiskCache_Data(C.QtObjectPtr(ptr.Pointer()), C.CString(url))))
	}
	return nil
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaDataITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QNetworkDiskCache_Prepare(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCacheMetaData(metaData)))))
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetCacheDirectory(C.QtObjectPtr(ptr.Pointer()), C.CString(cacheDir))
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaDataITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCacheMetaData(metaData)))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
