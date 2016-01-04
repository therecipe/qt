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
func callbackQNetworkDiskCacheClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkDiskCache::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkDiskCache) Clear() {
	defer qt.Recovering("QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) ClearDefault() {
	defer qt.Recovering("QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::data")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) Expire() int64 {
	defer qt.Recovering("QNetworkDiskCache::expire")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_Expire(ptr.Pointer()))
	}
	return 0
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

func (ptr *QNetworkDiskCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkDiskCacheTimerEvent
func callbackQNetworkDiskCacheTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkDiskCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkDiskCacheChildEvent
func callbackQNetworkDiskCacheChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkDiskCacheCustomEvent
func callbackQNetworkDiskCacheCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkDiskCache) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
