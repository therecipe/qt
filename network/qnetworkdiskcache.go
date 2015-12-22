package network

//#include "network.h"
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
	for len(n.ObjectName()) < len("QNetworkDiskCache_") {
		n.SetObjectName("QNetworkDiskCache_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return ptr
}

func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	defer qt.Recovering("QNetworkDiskCache::QNetworkDiskCache")

	return NewQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	defer qt.Recovering("QNetworkDiskCache::cacheDirectory")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkDiskCache_CacheDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkDiskCache) CacheSize() int64 {
	defer qt.Recovering("QNetworkDiskCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) ConnectClear(f func()) {
	defer qt.Recovering("connect QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectClear() {
	defer qt.Recovering("disconnect QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQNetworkDiskCacheClear
func callbackQNetworkDiskCacheClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QNetworkDiskCache::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::data")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) MaximumCacheSize() int64 {
	defer qt.Recovering("QNetworkDiskCache::maximumCacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_MaximumCacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::prepare")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	defer qt.Recovering("QNetworkDiskCache::setCacheDirectory")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), C.CString(cacheDir))
	}
}

func (ptr *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	defer qt.Recovering("QNetworkDiskCache::setMaximumCacheSize")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetMaximumCacheSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	defer qt.Recovering("QNetworkDiskCache::~QNetworkDiskCache")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
