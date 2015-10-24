package network

//#include "qabstractnetworkcache.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCacheITF interface {
	core.QObjectITF
	QAbstractNetworkCachePTR() *QAbstractNetworkCache
}

func PointerFromQAbstractNetworkCache(ptr QAbstractNetworkCacheITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCachePTR().Pointer()
	}
	return nil
}

func QAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) *QAbstractNetworkCache {
	var n = new(QAbstractNetworkCache)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractNetworkCache_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCachePTR() *QAbstractNetworkCache {
	return ptr
}

func (ptr *QAbstractNetworkCache) Clear() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractNetworkCache) Data(url string) *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QAbstractNetworkCache_Data(C.QtObjectPtr(ptr.Pointer()), C.CString(url))))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaDataITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QAbstractNetworkCache_Prepare(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCacheMetaData(metaData)))))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaDataITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCacheMetaData(metaData)))
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
