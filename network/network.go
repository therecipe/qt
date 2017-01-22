// +build !minimal

package network

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "network.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtNetwork_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCache_ITF interface {
	core.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache {
	return ptr
}

func (ptr *QAbstractNetworkCache) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractNetworkCache) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
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
	return n
}
func NewQAbstractNetworkCache(parent core.QObject_ITF) *QAbstractNetworkCache {
	var tmpValue = NewQAbstractNetworkCacheFromPointer(C.QAbstractNetworkCache_NewQAbstractNetworkCache(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractNetworkCache_CacheSize
func callbackQAbstractNetworkCache_CacheSize(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::cacheSize"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(0)
}

func (ptr *QAbstractNetworkCache) ConnectCacheSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::cacheSize", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCacheSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::cacheSize")
	}
}

func (ptr *QAbstractNetworkCache) CacheSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractNetworkCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractNetworkCache_Clear
func callbackQAbstractNetworkCache_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractNetworkCache) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::clear", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::clear")
	}
}

func (ptr *QAbstractNetworkCache) Clear() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(ptr.Pointer())
	}
}

//export callbackQAbstractNetworkCache_Data
func callbackQAbstractNetworkCache_Data(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::data"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*core.QUrl) *core.QIODevice)(core.NewQUrlFromPointer(url)))
	}

	return core.PointerFromQIODevice(core.NewQIODevice())
}

func (ptr *QAbstractNetworkCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::data", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::data")
	}
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Insert
func callbackQAbstractNetworkCache_Insert(ptr unsafe.Pointer, device unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::insert"); signal != nil {
		signal.(func(*core.QIODevice))(core.NewQIODeviceFromPointer(device))
	}

}

func (ptr *QAbstractNetworkCache) ConnectInsert(f func(device *core.QIODevice)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::insert", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectInsert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::insert")
	}
}

func (ptr *QAbstractNetworkCache) Insert(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Insert(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

//export callbackQAbstractNetworkCache_MetaData
func callbackQAbstractNetworkCache_MetaData(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::metaData"); signal != nil {
		return PointerFromQNetworkCacheMetaData(signal.(func(*core.QUrl) *QNetworkCacheMetaData)(core.NewQUrlFromPointer(url)))
	}

	return PointerFromQNetworkCacheMetaData(NewQNetworkCacheMetaData())
}

func (ptr *QAbstractNetworkCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaData", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaData")
	}
}

func (ptr *QAbstractNetworkCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QAbstractNetworkCache_MetaData(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Prepare
func callbackQAbstractNetworkCache_Prepare(ptr unsafe.Pointer, metaData unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::prepare"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*QNetworkCacheMetaData) *core.QIODevice)(NewQNetworkCacheMetaDataFromPointer(metaData)))
	}

	return core.PointerFromQIODevice(core.NewQIODevice())
}

func (ptr *QAbstractNetworkCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::prepare", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectPrepare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::prepare")
	}
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Remove
func callbackQAbstractNetworkCache_Remove(ptr unsafe.Pointer, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::remove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl) bool)(core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAbstractNetworkCache) ConnectRemove(f func(url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::remove", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectRemove() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::remove")
	}
}

func (ptr *QAbstractNetworkCache) Remove(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_Remove(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_UpdateMetaData
func callbackQAbstractNetworkCache_UpdateMetaData(ptr unsafe.Pointer, metaData unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::updateMetaData"); signal != nil {
		signal.(func(*QNetworkCacheMetaData))(NewQNetworkCacheMetaDataFromPointer(metaData))
	}

}

func (ptr *QAbstractNetworkCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::updateMetaData", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectUpdateMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::updateMetaData")
	}
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

//export callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache
func callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::~QAbstractNetworkCache"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DestroyQAbstractNetworkCacheDefault()
	}
}

func (ptr *QAbstractNetworkCache) ConnectDestroyQAbstractNetworkCache(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::~QAbstractNetworkCache", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDestroyQAbstractNetworkCache() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::~QAbstractNetworkCache")
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCacheDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCacheDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractNetworkCache_TimerEvent
func callbackQAbstractNetworkCache_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::timerEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::timerEvent")
	}
}

func (ptr *QAbstractNetworkCache) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractNetworkCache_ChildEvent
func callbackQAbstractNetworkCache_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::childEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::childEvent")
	}
}

func (ptr *QAbstractNetworkCache) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractNetworkCache_ConnectNotify
func callbackQAbstractNetworkCache_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::connectNotify", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::connectNotify")
	}
}

func (ptr *QAbstractNetworkCache) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractNetworkCache_CustomEvent
func callbackQAbstractNetworkCache_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::customEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::customEvent")
	}
}

func (ptr *QAbstractNetworkCache) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractNetworkCache_DeleteLater
func callbackQAbstractNetworkCache_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractNetworkCache) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::deleteLater", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::deleteLater")
	}
}

func (ptr *QAbstractNetworkCache) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractNetworkCache_DisconnectNotify
func callbackQAbstractNetworkCache_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::disconnectNotify", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::disconnectNotify")
	}
}

func (ptr *QAbstractNetworkCache) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractNetworkCache) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractNetworkCache_Event
func callbackQAbstractNetworkCache_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractNetworkCacheFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractNetworkCache) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::event", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::event")
	}
}

func (ptr *QAbstractNetworkCache) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractNetworkCache) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_EventFilter
func callbackQAbstractNetworkCache_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractNetworkCacheFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractNetworkCache) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::eventFilter", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::eventFilter")
	}
}

func (ptr *QAbstractNetworkCache) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractNetworkCache) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_MetaObject
func callbackQAbstractNetworkCache_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractNetworkCacheFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractNetworkCache) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaObject", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaObject")
	}
}

func (ptr *QAbstractNetworkCache) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractNetworkCache_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractNetworkCache_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAbstractSocket struct {
	core.QIODevice
}

type QAbstractSocket_ITF interface {
	core.QIODevice_ITF
	QAbstractSocket_PTR() *QAbstractSocket
}

func (ptr *QAbstractSocket) QAbstractSocket_PTR() *QAbstractSocket {
	return ptr
}

func (ptr *QAbstractSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractSocket(ptr QAbstractSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSocketFromPointer(ptr unsafe.Pointer) *QAbstractSocket {
	var n = new(QAbstractSocket)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QAbstractSocket__BindFlag
//QAbstractSocket::BindFlag
type QAbstractSocket__BindFlag int64

const (
	QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x0)
	QAbstractSocket__ShareAddress       QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x1)
	QAbstractSocket__DontShareAddress   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x2)
	QAbstractSocket__ReuseAddressHint   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x4)
)

//go:generate stringer -type=QAbstractSocket__NetworkLayerProtocol
//QAbstractSocket::NetworkLayerProtocol
type QAbstractSocket__NetworkLayerProtocol int64

const (
	QAbstractSocket__IPv4Protocol                QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(0)
	QAbstractSocket__IPv6Protocol                QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(1)
	QAbstractSocket__AnyIPProtocol               QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(2)
	QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(-1)
)

//go:generate stringer -type=QAbstractSocket__PauseMode
//QAbstractSocket::PauseMode
type QAbstractSocket__PauseMode int64

const (
	QAbstractSocket__PauseNever       QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x0)
	QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x1)
)

//go:generate stringer -type=QAbstractSocket__SocketError
//QAbstractSocket::SocketError
type QAbstractSocket__SocketError int64

const (
	QAbstractSocket__ConnectionRefusedError           QAbstractSocket__SocketError = QAbstractSocket__SocketError(0)
	QAbstractSocket__RemoteHostClosedError            QAbstractSocket__SocketError = QAbstractSocket__SocketError(1)
	QAbstractSocket__HostNotFoundError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(2)
	QAbstractSocket__SocketAccessError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(3)
	QAbstractSocket__SocketResourceError              QAbstractSocket__SocketError = QAbstractSocket__SocketError(4)
	QAbstractSocket__SocketTimeoutError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(5)
	QAbstractSocket__DatagramTooLargeError            QAbstractSocket__SocketError = QAbstractSocket__SocketError(6)
	QAbstractSocket__NetworkError                     QAbstractSocket__SocketError = QAbstractSocket__SocketError(7)
	QAbstractSocket__AddressInUseError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(8)
	QAbstractSocket__SocketAddressNotAvailableError   QAbstractSocket__SocketError = QAbstractSocket__SocketError(9)
	QAbstractSocket__UnsupportedSocketOperationError  QAbstractSocket__SocketError = QAbstractSocket__SocketError(10)
	QAbstractSocket__UnfinishedSocketOperationError   QAbstractSocket__SocketError = QAbstractSocket__SocketError(11)
	QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = QAbstractSocket__SocketError(12)
	QAbstractSocket__SslHandshakeFailedError          QAbstractSocket__SocketError = QAbstractSocket__SocketError(13)
	QAbstractSocket__ProxyConnectionRefusedError      QAbstractSocket__SocketError = QAbstractSocket__SocketError(14)
	QAbstractSocket__ProxyConnectionClosedError       QAbstractSocket__SocketError = QAbstractSocket__SocketError(15)
	QAbstractSocket__ProxyConnectionTimeoutError      QAbstractSocket__SocketError = QAbstractSocket__SocketError(16)
	QAbstractSocket__ProxyNotFoundError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(17)
	QAbstractSocket__ProxyProtocolError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(18)
	QAbstractSocket__OperationError                   QAbstractSocket__SocketError = QAbstractSocket__SocketError(19)
	QAbstractSocket__SslInternalError                 QAbstractSocket__SocketError = QAbstractSocket__SocketError(20)
	QAbstractSocket__SslInvalidUserDataError          QAbstractSocket__SocketError = QAbstractSocket__SocketError(21)
	QAbstractSocket__TemporaryError                   QAbstractSocket__SocketError = QAbstractSocket__SocketError(22)
	QAbstractSocket__UnknownSocketError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(-1)
)

//go:generate stringer -type=QAbstractSocket__SocketOption
//QAbstractSocket::SocketOption
type QAbstractSocket__SocketOption int64

const (
	QAbstractSocket__LowDelayOption                QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(0)
	QAbstractSocket__KeepAliveOption               QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(1)
	QAbstractSocket__MulticastTtlOption            QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(2)
	QAbstractSocket__MulticastLoopbackOption       QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(3)
	QAbstractSocket__TypeOfServiceOption           QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(4)
	QAbstractSocket__SendBufferSizeSocketOption    QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(5)
	QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(6)
)

//go:generate stringer -type=QAbstractSocket__SocketState
//QAbstractSocket::SocketState
type QAbstractSocket__SocketState int64

const (
	QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = QAbstractSocket__SocketState(0)
	QAbstractSocket__HostLookupState  QAbstractSocket__SocketState = QAbstractSocket__SocketState(1)
	QAbstractSocket__ConnectingState  QAbstractSocket__SocketState = QAbstractSocket__SocketState(2)
	QAbstractSocket__ConnectedState   QAbstractSocket__SocketState = QAbstractSocket__SocketState(3)
	QAbstractSocket__BoundState       QAbstractSocket__SocketState = QAbstractSocket__SocketState(4)
	QAbstractSocket__ListeningState   QAbstractSocket__SocketState = QAbstractSocket__SocketState(5)
	QAbstractSocket__ClosingState     QAbstractSocket__SocketState = QAbstractSocket__SocketState(6)
)

//go:generate stringer -type=QAbstractSocket__SocketType
//QAbstractSocket::SocketType
type QAbstractSocket__SocketType int64

const (
	QAbstractSocket__TcpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(0)
	QAbstractSocket__UdpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(1)
	QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = QAbstractSocket__SocketType(-1)
)

func NewQAbstractSocket(socketType QAbstractSocket__SocketType, parent core.QObject_ITF) *QAbstractSocket {
	var tmpValue = NewQAbstractSocketFromPointer(C.QAbstractSocket_NewQAbstractSocket(C.longlong(socketType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAbstractSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Abort(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_AtEnd
func callbackQAbstractSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QAbstractSocket) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::atEnd", f)
	}
}

func (ptr *QAbstractSocket) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::atEnd")
	}
}

func (ptr *QAbstractSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Bind(address QHostAddress_ITF, port uint16, mode QAbstractSocket__BindFlag) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Bind(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Bind2(port uint16, mode QAbstractSocket__BindFlag) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Bind2(ptr.Pointer(), C.ushort(port), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQAbstractSocket_BytesAvailable
func callbackQAbstractSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QAbstractSocket) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::bytesAvailable", f)
	}
}

func (ptr *QAbstractSocket) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::bytesAvailable")
	}
}

func (ptr *QAbstractSocket) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_BytesToWrite
func callbackQAbstractSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QAbstractSocket) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::bytesToWrite", f)
	}
}

func (ptr *QAbstractSocket) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::bytesToWrite")
	}
}

func (ptr *QAbstractSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_CanReadLine
func callbackQAbstractSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QAbstractSocket) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::canReadLine", f)
	}
}

func (ptr *QAbstractSocket) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::canReadLine")
	}
}

func (ptr *QAbstractSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractSocket_Close
func callbackQAbstractSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QAbstractSocket) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::close", f)
	}
}

func (ptr *QAbstractSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::close")
	}
}

func (ptr *QAbstractSocket) Close() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Close(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_CloseDefault(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_ConnectToHost2
func callbackQAbstractSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QAbstractSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost2", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnectToHost2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost2")
	}
}

func (ptr *QAbstractSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QAbstractSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQAbstractSocket_ConnectToHost
func callbackQAbstractSocket_ConnectToHost(ptr unsafe.Pointer, hostName C.struct_QtNetwork_PackedString, port C.ushort, openMode C.longlong, protocol C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectToHostDefault(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QAbstractSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnectToHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost")
	}
}

func (ptr *QAbstractSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QAbstractSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QAbstractSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QAbstractSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQAbstractSocket_Connected
func callbackQAbstractSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connected")
	}
}

func (ptr *QAbstractSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Connected(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_DisconnectFromHost
func callbackQAbstractSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDisconnectFromHost(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectFromHost", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnectFromHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectFromHost")
	}
}

func (ptr *QAbstractSocket) DisconnectFromHost() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) DisconnectFromHostDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_Disconnected
func callbackQAbstractSocket_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnected")
	}
}

func (ptr *QAbstractSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_Error2
func callbackQAbstractSocket_Error2(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::error2"); signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QAbstractSocket) ConnectError2(f func(socketError QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::error2", f)
	}
}

func (ptr *QAbstractSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::error2")
	}
}

func (ptr *QAbstractSocket) Error2(socketError QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Error2(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QAbstractSocket) Error() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QAbstractSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractSocket_HostFound
func callbackQAbstractSocket_HostFound(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::hostFound"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) ConnectHostFound(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectHostFound(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::hostFound", f)
	}
}

func (ptr *QAbstractSocket) DisconnectHostFound() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectHostFound(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::hostFound")
	}
}

func (ptr *QAbstractSocket) HostFound() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_HostFound(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_IsSequential
func callbackQAbstractSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QAbstractSocket) ConnectIsSequential(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::isSequential", f)
	}
}

func (ptr *QAbstractSocket) DisconnectIsSequential() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::isSequential")
	}
}

func (ptr *QAbstractSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) LocalAddress() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QAbstractSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QAbstractSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PauseMode() QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return QAbstractSocket__PauseMode(C.QAbstractSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PeerAddress() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QAbstractSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QAbstractSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) Proxy() *QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QAbstractSocket_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractSocket_ProxyAuthenticationRequired
func callbackQAbstractSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::proxyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkProxy, *QAuthenticator))(NewQNetworkProxyFromPointer(proxy), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QAbstractSocket) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::proxyAuthenticationRequired", f)
	}
}

func (ptr *QAbstractSocket) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::proxyAuthenticationRequired")
	}
}

func (ptr *QAbstractSocket) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ProxyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkProxy(proxy), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QAbstractSocket) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_ReadData
func callbackQAbstractSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQAbstractSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QAbstractSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::readData", f)
	}
}

func (ptr *QAbstractSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::readData")
	}
}

func (ptr *QAbstractSocket) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QAbstractSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QAbstractSocket) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QAbstractSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQAbstractSocket_ReadLineData
func callbackQAbstractSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxlen C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxlen)))
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxlen)))
}

func (ptr *QAbstractSocket) ConnectReadLineData(f func(data string, maxlen int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::readLineData", f)
	}
}

func (ptr *QAbstractSocket) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::readLineData")
	}
}

func (ptr *QAbstractSocket) ReadLineData(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QAbstractSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QAbstractSocket) ReadLineDataDefault(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QAbstractSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

//export callbackQAbstractSocket_Resume
func callbackQAbstractSocket_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QAbstractSocket) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::resume", f)
	}
}

func (ptr *QAbstractSocket) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::resume")
	}
}

func (ptr *QAbstractSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ResumeDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) SetLocalAddress(address QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetLocalAddress(ptr.Pointer(), PointerFromQHostAddress(address))
	}
}

func (ptr *QAbstractSocket) SetLocalPort(port uint16) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetLocalPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QAbstractSocket) SetPauseMode(pauseMode QAbstractSocket__PauseMode) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPauseMode(ptr.Pointer(), C.longlong(pauseMode))
	}
}

func (ptr *QAbstractSocket) SetPeerAddress(address QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPeerAddress(ptr.Pointer(), PointerFromQHostAddress(address))
	}
}

func (ptr *QAbstractSocket) SetPeerName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QAbstractSocket_SetPeerName(ptr.Pointer(), nameC)
	}
}

func (ptr *QAbstractSocket) SetPeerPort(port uint16) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPeerPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

//export callbackQAbstractSocket_SetReadBufferSize
func callbackQAbstractSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QAbstractSocket) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setReadBufferSize", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setReadBufferSize")
	}
}

func (ptr *QAbstractSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QAbstractSocket) SetReadBufferSizeDefault(size int64) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QAbstractSocket) SetSocketError(socketError QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketError(ptr.Pointer(), C.longlong(socketError))
	}
}

//export callbackQAbstractSocket_SetSocketOption
func callbackQAbstractSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QAbstractSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setSocketOption", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSetSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setSocketOption")
	}
}

func (ptr *QAbstractSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SetSocketState(state QAbstractSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketState(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQAbstractSocket_SocketOption
func callbackQAbstractSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQAbstractSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QAbstractSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::socketOption", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::socketOption")
	}
}

func (ptr *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) SocketType() QAbstractSocket__SocketType {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketType(C.QAbstractSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) State() QAbstractSocket__SocketState {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketState(C.QAbstractSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_StateChanged
func callbackQAbstractSocket_StateChanged(ptr unsafe.Pointer, socketState C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::stateChanged"); signal != nil {
		signal.(func(QAbstractSocket__SocketState))(QAbstractSocket__SocketState(socketState))
	}

}

func (ptr *QAbstractSocket) ConnectStateChanged(f func(socketState QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::stateChanged", f)
	}
}

func (ptr *QAbstractSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::stateChanged")
	}
}

func (ptr *QAbstractSocket) StateChanged(socketState QAbstractSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_StateChanged(ptr.Pointer(), C.longlong(socketState))
	}
}

//export callbackQAbstractSocket_WaitForBytesWritten
func callbackQAbstractSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForBytesWritten", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForBytesWritten")
	}
}

func (ptr *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForConnected
func callbackQAbstractSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForConnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForConnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForConnected")
	}
}

func (ptr *QAbstractSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForConnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForDisconnected
func callbackQAbstractSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForDisconnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForDisconnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForDisconnected")
	}
}

func (ptr *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForDisconnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForReadyRead
func callbackQAbstractSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForReadyRead", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForReadyRead")
	}
}

func (ptr *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WriteData
func callbackQAbstractSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, size C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(size)))
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(size)))
}

func (ptr *QAbstractSocket) ConnectWriteData(f func(data string, size int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::writeData", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::writeData")
	}
}

func (ptr *QAbstractSocket) WriteData(data string, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QAbstractSocket_WriteData(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

func (ptr *QAbstractSocket) WriteDataDefault(data string, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QAbstractSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

//export callbackQAbstractSocket_DestroyQAbstractSocket
func callbackQAbstractSocket_DestroyQAbstractSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::~QAbstractSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DestroyQAbstractSocketDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDestroyQAbstractSocket(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::~QAbstractSocket", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDestroyQAbstractSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::~QAbstractSocket")
	}
}

func (ptr *QAbstractSocket) DestroyQAbstractSocket() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DestroyQAbstractSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) DestroyQAbstractSocketDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DestroyQAbstractSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractSocket_Open
func callbackQAbstractSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QAbstractSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::open", f)
	}
}

func (ptr *QAbstractSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::open")
	}
}

func (ptr *QAbstractSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQAbstractSocket_Pos
func callbackQAbstractSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).PosDefault())
}

func (ptr *QAbstractSocket) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::pos", f)
	}
}

func (ptr *QAbstractSocket) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::pos")
	}
}

func (ptr *QAbstractSocket) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_Reset
func callbackQAbstractSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QAbstractSocket) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::reset", f)
	}
}

func (ptr *QAbstractSocket) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::reset")
	}
}

func (ptr *QAbstractSocket) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractSocket_Seek
func callbackQAbstractSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QAbstractSocket) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::seek", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::seek")
	}
}

func (ptr *QAbstractSocket) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQAbstractSocket_Size
func callbackQAbstractSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QAbstractSocket) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::size", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::size")
	}
}

func (ptr *QAbstractSocket) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_TimerEvent
func callbackQAbstractSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::timerEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::timerEvent")
	}
}

func (ptr *QAbstractSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractSocket_ChildEvent
func callbackQAbstractSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::childEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::childEvent")
	}
}

func (ptr *QAbstractSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractSocket_ConnectNotify
func callbackQAbstractSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectNotify", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectNotify")
	}
}

func (ptr *QAbstractSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSocket_CustomEvent
func callbackQAbstractSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::customEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::customEvent")
	}
}

func (ptr *QAbstractSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractSocket_DeleteLater
func callbackQAbstractSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::deleteLater", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::deleteLater")
	}
}

func (ptr *QAbstractSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractSocket_DisconnectNotify
func callbackQAbstractSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectNotify", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectNotify")
	}
}

func (ptr *QAbstractSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSocket_Event
func callbackQAbstractSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::event", f)
	}
}

func (ptr *QAbstractSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::event")
	}
}

func (ptr *QAbstractSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractSocket_EventFilter
func callbackQAbstractSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::eventFilter", f)
	}
}

func (ptr *QAbstractSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::eventFilter")
	}
}

func (ptr *QAbstractSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractSocket_MetaObject
func callbackQAbstractSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::metaObject", f)
	}
}

func (ptr *QAbstractSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::metaObject")
	}
}

func (ptr *QAbstractSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAuthenticator struct {
	ptr unsafe.Pointer
}

type QAuthenticator_ITF interface {
	QAuthenticator_PTR() *QAuthenticator
}

func (ptr *QAuthenticator) QAuthenticator_PTR() *QAuthenticator {
	return ptr
}

func (ptr *QAuthenticator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAuthenticator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAuthenticator(ptr QAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = new(QAuthenticator)
	n.SetPointer(ptr)
	return n
}
func NewQAuthenticator() *QAuthenticator {
	var tmpValue = NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator())
	runtime.SetFinalizer(tmpValue, (*QAuthenticator).DestroyQAuthenticator)
	return tmpValue
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {
	var tmpValue = NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator2(PointerFromQAuthenticator(other)))
	runtime.SetFinalizer(tmpValue, (*QAuthenticator).DestroyQAuthenticator)
	return tmpValue
}

func (ptr *QAuthenticator) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QAuthenticator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {
	if ptr.Pointer() != nil {
		var optC = C.CString(opt)
		defer C.free(unsafe.Pointer(optC))
		var tmpValue = core.NewQVariantFromPointer(C.QAuthenticator_Option(ptr.Pointer(), optC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAuthenticator) Password() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAuthenticator_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) Realm() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAuthenticator_Realm(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) SetOption(opt string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var optC = C.CString(opt)
		defer C.free(unsafe.Pointer(optC))
		C.QAuthenticator_SetOption(ptr.Pointer(), optC, core.PointerFromQVariant(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		C.QAuthenticator_SetPassword(ptr.Pointer(), passwordC)
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	if ptr.Pointer() != nil {
		var userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
		C.QAuthenticator_SetUser(ptr.Pointer(), userC)
	}
}

func (ptr *QAuthenticator) User() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAuthenticator_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {
	if ptr.Pointer() != nil {
		C.QAuthenticator_DestroyQAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsDomainNameRecord struct {
	ptr unsafe.Pointer
}

type QDnsDomainNameRecord_ITF interface {
	QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord
}

func (ptr *QDnsDomainNameRecord) QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord {
	return ptr
}

func (ptr *QDnsDomainNameRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDnsDomainNameRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDnsDomainNameRecord(ptr QDnsDomainNameRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsDomainNameRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) *QDnsDomainNameRecord {
	var n = new(QDnsDomainNameRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
	return tmpValue
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecord_ITF) *QDnsDomainNameRecord {
	var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord2(PointerFromQDnsDomainNameRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
	return tmpValue
}

func (ptr *QDnsDomainNameRecord) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsDomainNameRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) Swap(other QDnsDomainNameRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_Swap(ptr.Pointer(), PointerFromQDnsDomainNameRecord(other))
	}
}

func (ptr *QDnsDomainNameRecord) TimeToLive() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsDomainNameRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsDomainNameRecord) Value() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsDomainNameRecord_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) DestroyQDnsDomainNameRecord() {
	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsHostAddressRecord struct {
	ptr unsafe.Pointer
}

type QDnsHostAddressRecord_ITF interface {
	QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord
}

func (ptr *QDnsHostAddressRecord) QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord {
	return ptr
}

func (ptr *QDnsHostAddressRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDnsHostAddressRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDnsHostAddressRecord(ptr QDnsHostAddressRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsHostAddressRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) *QDnsHostAddressRecord {
	var n = new(QDnsHostAddressRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	var tmpValue = NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
	return tmpValue
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecord_ITF) *QDnsHostAddressRecord {
	var tmpValue = NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord2(PointerFromQDnsHostAddressRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
	return tmpValue
}

func (ptr *QDnsHostAddressRecord) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsHostAddressRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_Swap(ptr.Pointer(), PointerFromQDnsHostAddressRecord(other))
	}
}

func (ptr *QDnsHostAddressRecord) TimeToLive() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsHostAddressRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsHostAddressRecord) Value() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QDnsHostAddressRecord_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {
	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsLookup struct {
	core.QObject
}

type QDnsLookup_ITF interface {
	core.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func (ptr *QDnsLookup) QDnsLookup_PTR() *QDnsLookup {
	return ptr
}

func (ptr *QDnsLookup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDnsLookup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDnsLookup(ptr QDnsLookup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookup_PTR().Pointer()
	}
	return nil
}

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = new(QDnsLookup)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QDnsLookup__Error
//QDnsLookup::Error
type QDnsLookup__Error int64

const (
	QDnsLookup__NoError                 QDnsLookup__Error = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           QDnsLookup__Error = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError QDnsLookup__Error = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     QDnsLookup__Error = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       QDnsLookup__Error = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      QDnsLookup__Error = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      QDnsLookup__Error = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           QDnsLookup__Error = QDnsLookup__Error(7)
)

//go:generate stringer -type=QDnsLookup__Type
//QDnsLookup::Type
type QDnsLookup__Type int64

const (
	QDnsLookup__A     QDnsLookup__Type = QDnsLookup__Type(1)
	QDnsLookup__AAAA  QDnsLookup__Type = QDnsLookup__Type(28)
	QDnsLookup__ANY   QDnsLookup__Type = QDnsLookup__Type(255)
	QDnsLookup__CNAME QDnsLookup__Type = QDnsLookup__Type(5)
	QDnsLookup__MX    QDnsLookup__Type = QDnsLookup__Type(15)
	QDnsLookup__NS    QDnsLookup__Type = QDnsLookup__Type(2)
	QDnsLookup__PTR   QDnsLookup__Type = QDnsLookup__Type(12)
	QDnsLookup__SRV   QDnsLookup__Type = QDnsLookup__Type(33)
	QDnsLookup__TXT   QDnsLookup__Type = QDnsLookup__Type(16)
)

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup3(C.longlong(ty), nameC, PointerFromQHostAddress(nameserver), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {
	if ptr.Pointer() != nil {
		return QDnsLookup__Error(C.QDnsLookup_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsLookup) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsLookup_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsLookup_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Nameserver() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QDnsLookup_Nameserver(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDnsLookup_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_SetNameserver(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

func (ptr *QDnsLookup) SetType(vqd QDnsLookup__Type) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_SetType(ptr.Pointer(), C.longlong(vqd))
	}
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(ptr.Pointer()))
	}
	return 0
}

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {
	var tmpValue = NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup2(C.longlong(ty), nameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDnsLookup_Abort
func callbackQDnsLookup_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::abort"); signal != nil {
		signal.(func())()
	} else {
		NewQDnsLookupFromPointer(ptr).AbortDefault()
	}
}

func (ptr *QDnsLookup) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::abort", f)
	}
}

func (ptr *QDnsLookup) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::abort")
	}
}

func (ptr *QDnsLookup) Abort() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_Abort(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) AbortDefault() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_AbortDefault(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) CanonicalNameRecords() []*QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsDomainNameRecord {
			var out = make([]*QDnsDomainNameRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).canonicalNameRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_CanonicalNameRecords(ptr.Pointer()))
	}
	return nil
}

//export callbackQDnsLookup_Finished
func callbackQDnsLookup_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::finished", f)
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::finished")
	}
}

func (ptr *QDnsLookup) Finished() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_Finished(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) HostAddressRecords() []*QDnsHostAddressRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsHostAddressRecord {
			var out = make([]*QDnsHostAddressRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).hostAddressRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_HostAddressRecords(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDnsLookup) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QDnsLookup_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDnsLookup_Lookup
func callbackQDnsLookup_Lookup(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::lookup"); signal != nil {
		signal.(func())()
	} else {
		NewQDnsLookupFromPointer(ptr).LookupDefault()
	}
}

func (ptr *QDnsLookup) ConnectLookup(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::lookup", f)
	}
}

func (ptr *QDnsLookup) DisconnectLookup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::lookup")
	}
}

func (ptr *QDnsLookup) Lookup() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_Lookup(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) LookupDefault() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_LookupDefault(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) MailExchangeRecords() []*QDnsMailExchangeRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsMailExchangeRecord {
			var out = make([]*QDnsMailExchangeRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).mailExchangeRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_MailExchangeRecords(ptr.Pointer()))
	}
	return nil
}

//export callbackQDnsLookup_NameChanged
func callbackQDnsLookup_NameChanged(ptr unsafe.Pointer, name C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::nameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(name))
	}

}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameChanged")
	}
}

func (ptr *QDnsLookup) NameChanged(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDnsLookup_NameChanged(ptr.Pointer(), nameC)
	}
}

func (ptr *QDnsLookup) NameServerRecords() []*QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsDomainNameRecord {
			var out = make([]*QDnsDomainNameRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).nameServerRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_NameServerRecords(ptr.Pointer()))
	}
	return nil
}

//export callbackQDnsLookup_NameserverChanged
func callbackQDnsLookup_NameserverChanged(ptr unsafe.Pointer, nameserver unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::nameserverChanged"); signal != nil {
		signal.(func(*QHostAddress))(NewQHostAddressFromPointer(nameserver))
	}

}

func (ptr *QDnsLookup) ConnectNameserverChanged(f func(nameserver *QHostAddress)) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameserverChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameserverChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameserverChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameserverChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameserverChanged")
	}
}

func (ptr *QDnsLookup) NameserverChanged(nameserver QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_NameserverChanged(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

func (ptr *QDnsLookup) PointerRecords() []*QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsDomainNameRecord {
			var out = make([]*QDnsDomainNameRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).pointerRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_PointerRecords(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDnsLookup) ServiceRecords() []*QDnsServiceRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsServiceRecord {
			var out = make([]*QDnsServiceRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).serviceRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_ServiceRecords(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDnsLookup) TextRecords() []*QDnsTextRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsTextRecord {
			var out = make([]*QDnsTextRecord, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsLookupFromPointer(l.data).textRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_TextRecords(ptr.Pointer()))
	}
	return nil
}

//export callbackQDnsLookup_TypeChanged
func callbackQDnsLookup_TypeChanged(ptr unsafe.Pointer, ty C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::typeChanged"); signal != nil {
		signal.(func(QDnsLookup__Type))(QDnsLookup__Type(ty))
	}

}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::typeChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::typeChanged")
	}
}

func (ptr *QDnsLookup) TypeChanged(ty QDnsLookup__Type) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_TypeChanged(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DestroyQDnsLookup(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) canonicalNameRecords_atList(i int) *QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsLookup_canonicalNameRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) hostAddressRecords_atList(i int) *QDnsHostAddressRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsHostAddressRecordFromPointer(C.QDnsLookup_hostAddressRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) mailExchangeRecords_atList(i int) *QDnsMailExchangeRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsMailExchangeRecordFromPointer(C.QDnsLookup_mailExchangeRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) nameServerRecords_atList(i int) *QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsLookup_nameServerRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) pointerRecords_atList(i int) *QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsLookup_pointerRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) serviceRecords_atList(i int) *QDnsServiceRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsServiceRecordFromPointer(C.QDnsLookup_serviceRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) textRecords_atList(i int) *QDnsTextRecord {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDnsTextRecordFromPointer(C.QDnsLookup_textRecords_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
		return tmpValue
	}
	return nil
}

//export callbackQDnsLookup_TimerEvent
func callbackQDnsLookup_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::timerEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::timerEvent")
	}
}

func (ptr *QDnsLookup) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDnsLookup) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDnsLookup_ChildEvent
func callbackQDnsLookup_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::childEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::childEvent")
	}
}

func (ptr *QDnsLookup) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDnsLookup) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDnsLookup_ConnectNotify
func callbackQDnsLookup_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDnsLookupFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDnsLookup) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::connectNotify", f)
	}
}

func (ptr *QDnsLookup) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::connectNotify")
	}
}

func (ptr *QDnsLookup) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDnsLookup) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDnsLookup_CustomEvent
func callbackQDnsLookup_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::customEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::customEvent")
	}
}

func (ptr *QDnsLookup) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDnsLookup) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDnsLookup_DeleteLater
func callbackQDnsLookup_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDnsLookupFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDnsLookup) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::deleteLater", f)
	}
}

func (ptr *QDnsLookup) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::deleteLater")
	}
}

func (ptr *QDnsLookup) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDnsLookup_DisconnectNotify
func callbackQDnsLookup_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDnsLookupFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDnsLookup) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::disconnectNotify", f)
	}
}

func (ptr *QDnsLookup) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::disconnectNotify")
	}
}

func (ptr *QDnsLookup) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDnsLookup) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDnsLookup_Event
func callbackQDnsLookup_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDnsLookupFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDnsLookup) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::event", f)
	}
}

func (ptr *QDnsLookup) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::event")
	}
}

func (ptr *QDnsLookup) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDnsLookup_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDnsLookup) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDnsLookup_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDnsLookup_EventFilter
func callbackQDnsLookup_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDnsLookupFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDnsLookup) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::eventFilter", f)
	}
}

func (ptr *QDnsLookup) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::eventFilter")
	}
}

func (ptr *QDnsLookup) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDnsLookup_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDnsLookup) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDnsLookup_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDnsLookup_MetaObject
func callbackQDnsLookup_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDnsLookupFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDnsLookup) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::metaObject", f)
	}
}

func (ptr *QDnsLookup) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::metaObject")
	}
}

func (ptr *QDnsLookup) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDnsLookup_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDnsLookup) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDnsLookup_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDnsMailExchangeRecord struct {
	ptr unsafe.Pointer
}

type QDnsMailExchangeRecord_ITF interface {
	QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord
}

func (ptr *QDnsMailExchangeRecord) QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord {
	return ptr
}

func (ptr *QDnsMailExchangeRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDnsMailExchangeRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDnsMailExchangeRecord(ptr QDnsMailExchangeRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsMailExchangeRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) *QDnsMailExchangeRecord {
	var n = new(QDnsMailExchangeRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	var tmpValue = NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
	return tmpValue
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecord_ITF) *QDnsMailExchangeRecord {
	var tmpValue = NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(PointerFromQDnsMailExchangeRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
	return tmpValue
}

func (ptr *QDnsMailExchangeRecord) Exchange() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsMailExchangeRecord_Exchange(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsMailExchangeRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Preference() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QDnsMailExchangeRecord_Preference(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsMailExchangeRecord) Swap(other QDnsMailExchangeRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_Swap(ptr.Pointer(), PointerFromQDnsMailExchangeRecord(other))
	}
}

func (ptr *QDnsMailExchangeRecord) TimeToLive() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsMailExchangeRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsMailExchangeRecord) DestroyQDnsMailExchangeRecord() {
	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsServiceRecord struct {
	ptr unsafe.Pointer
}

type QDnsServiceRecord_ITF interface {
	QDnsServiceRecord_PTR() *QDnsServiceRecord
}

func (ptr *QDnsServiceRecord) QDnsServiceRecord_PTR() *QDnsServiceRecord {
	return ptr
}

func (ptr *QDnsServiceRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDnsServiceRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDnsServiceRecord(ptr QDnsServiceRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsServiceRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsServiceRecordFromPointer(ptr unsafe.Pointer) *QDnsServiceRecord {
	var n = new(QDnsServiceRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsServiceRecord() *QDnsServiceRecord {
	var tmpValue = NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
	return tmpValue
}

func NewQDnsServiceRecord2(other QDnsServiceRecord_ITF) *QDnsServiceRecord {
	var tmpValue = NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord2(PointerFromQDnsServiceRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
	return tmpValue
}

func (ptr *QDnsServiceRecord) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsServiceRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) Port() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QDnsServiceRecord_Port(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsServiceRecord) Priority() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QDnsServiceRecord_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsServiceRecord) Swap(other QDnsServiceRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_Swap(ptr.Pointer(), PointerFromQDnsServiceRecord(other))
	}
}

func (ptr *QDnsServiceRecord) Target() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsServiceRecord_Target(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) TimeToLive() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsServiceRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsServiceRecord) Weight() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QDnsServiceRecord_Weight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsServiceRecord) DestroyQDnsServiceRecord() {
	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_DestroyQDnsServiceRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsTextRecord struct {
	ptr unsafe.Pointer
}

type QDnsTextRecord_ITF interface {
	QDnsTextRecord_PTR() *QDnsTextRecord
}

func (ptr *QDnsTextRecord) QDnsTextRecord_PTR() *QDnsTextRecord {
	return ptr
}

func (ptr *QDnsTextRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDnsTextRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDnsTextRecord(ptr QDnsTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsTextRecordFromPointer(ptr unsafe.Pointer) *QDnsTextRecord {
	var n = new(QDnsTextRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsTextRecord() *QDnsTextRecord {
	var tmpValue = NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
	return tmpValue
}

func NewQDnsTextRecord2(other QDnsTextRecord_ITF) *QDnsTextRecord {
	var tmpValue = NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord2(PointerFromQDnsTextRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
	return tmpValue
}

func (ptr *QDnsTextRecord) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsTextRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsTextRecord) Swap(other QDnsTextRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsTextRecord_Swap(ptr.Pointer(), PointerFromQDnsTextRecord(other))
	}
}

func (ptr *QDnsTextRecord) TimeToLive() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsTextRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsTextRecord) Values() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDnsTextRecordFromPointer(l.data).values_atList(i)
			}
			return out
		}(C.QDnsTextRecord_Values(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {
	if ptr.Pointer() != nil {
		C.QDnsTextRecord_DestroyQDnsTextRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsTextRecord) values_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDnsTextRecord_values_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

type QHostAddress struct {
	ptr unsafe.Pointer
}

type QHostAddress_ITF interface {
	QHostAddress_PTR() *QHostAddress
}

func (ptr *QHostAddress) QHostAddress_PTR() *QHostAddress {
	return ptr
}

func (ptr *QHostAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHostAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHostAddress(ptr QHostAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostAddress_PTR().Pointer()
	}
	return nil
}

func NewQHostAddressFromPointer(ptr unsafe.Pointer) *QHostAddress {
	var n = new(QHostAddress)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QHostAddress__SpecialAddress
//QHostAddress::SpecialAddress
type QHostAddress__SpecialAddress int64

const (
	QHostAddress__Null          QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(0)
	QHostAddress__Broadcast     QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(1)
	QHostAddress__LocalHost     QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(2)
	QHostAddress__LocalHostIPv6 QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(3)
	QHostAddress__Any           QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(4)
	QHostAddress__AnyIPv6       QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(5)
	QHostAddress__AnyIPv4       QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(6)
)

func NewQHostAddress() *QHostAddress {
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress())
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress9(address QHostAddress__SpecialAddress) *QHostAddress {
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress9(C.longlong(address)))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress8(address QHostAddress_ITF) *QHostAddress {
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress8(PointerFromQHostAddress(address)))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress7(address string) *QHostAddress {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress7(addressC))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress4(ip6Addr string) *QHostAddress {
	var ip6AddrC = C.CString(ip6Addr)
	defer C.free(unsafe.Pointer(ip6AddrC))
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress4(ip6AddrC))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress2(ip4Addr uint) *QHostAddress {
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress2(C.uint(uint32(ip4Addr))))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress3(ip6Addr string) *QHostAddress {
	var ip6AddrC = C.CString(ip6Addr)
	defer C.free(unsafe.Pointer(ip6AddrC))
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress3(ip6AddrC))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func (ptr *QHostAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QHostAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {
	if ptr.Pointer() != nil {
		return C.QHostAddress_IsInSubnet(ptr.Pointer(), PointerFromQHostAddress(subnet), C.int(int32(netmask))) != 0
	}
	return false
}

func (ptr *QHostAddress) IsLoopback() bool {
	if ptr.Pointer() != nil {
		return C.QHostAddress_IsLoopback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) IsMulticast() bool {
	if ptr.Pointer() != nil {
		return C.QHostAddress_IsMulticast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QHostAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) Protocol() QAbstractSocket__NetworkLayerProtocol {
	if ptr.Pointer() != nil {
		return QAbstractSocket__NetworkLayerProtocol(C.QHostAddress_Protocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostAddress) ScopeId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHostAddress_ScopeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) SetAddress6(address string) bool {
	if ptr.Pointer() != nil {
		var addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
		return C.QHostAddress_SetAddress6(ptr.Pointer(), addressC) != 0
	}
	return false
}

func (ptr *QHostAddress) SetAddress3(ip6Addr string) {
	if ptr.Pointer() != nil {
		var ip6AddrC = C.CString(ip6Addr)
		defer C.free(unsafe.Pointer(ip6AddrC))
		C.QHostAddress_SetAddress3(ptr.Pointer(), ip6AddrC)
	}
}

func (ptr *QHostAddress) SetAddress(ip4Addr uint) {
	if ptr.Pointer() != nil {
		C.QHostAddress_SetAddress(ptr.Pointer(), C.uint(uint32(ip4Addr)))
	}
}

func (ptr *QHostAddress) SetAddress2(ip6Addr string) {
	if ptr.Pointer() != nil {
		var ip6AddrC = C.CString(ip6Addr)
		defer C.free(unsafe.Pointer(ip6AddrC))
		C.QHostAddress_SetAddress2(ptr.Pointer(), ip6AddrC)
	}
}

func (ptr *QHostAddress) SetScopeId(id string) {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		C.QHostAddress_SetScopeId(ptr.Pointer(), idC)
	}
}

func (ptr *QHostAddress) Swap(other QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QHostAddress_Swap(ptr.Pointer(), PointerFromQHostAddress(other))
	}
}

func (ptr *QHostAddress) ToIPv4Address() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QHostAddress_ToIPv4Address(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostAddress) ToIPv4Address2(ok bool) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QHostAddress_ToIPv4Address2(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))))
	}
	return 0
}

func (ptr *QHostAddress) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHostAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) DestroyQHostAddress() {
	if ptr.Pointer() != nil {
		C.QHostAddress_DestroyQHostAddress(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHostInfo struct {
	ptr unsafe.Pointer
}

type QHostInfo_ITF interface {
	QHostInfo_PTR() *QHostInfo
}

func (ptr *QHostInfo) QHostInfo_PTR() *QHostInfo {
	return ptr
}

func (ptr *QHostInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHostInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHostInfo(ptr QHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQHostInfoFromPointer(ptr unsafe.Pointer) *QHostInfo {
	var n = new(QHostInfo)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QHostInfo__HostInfoError
//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      QHostInfo__HostInfoError = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound QHostInfo__HostInfoError = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError QHostInfo__HostInfoError = QHostInfo__HostInfoError(2)
)

func QHostInfo_LocalHostName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalHostName())
}

func (ptr *QHostInfo) LocalHostName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalHostName())
}

func NewQHostInfo2(other QHostInfo_ITF) *QHostInfo {
	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo2(PointerFromQHostInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func NewQHostInfo(id int) *QHostInfo {
	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo(C.int(int32(id))))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func QHostInfo_AbortHostLookup(id int) {
	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(int32(id)))
}

func (ptr *QHostInfo) AbortHostLookup(id int) {
	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(int32(id)))
}

func (ptr *QHostInfo) Addresses() []*QHostAddress {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QHostAddress {
			var out = make([]*QHostAddress, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQHostInfoFromPointer(l.data).addresses_atList(i)
			}
			return out
		}(C.QHostInfo_Addresses(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {
	if ptr.Pointer() != nil {
		return QHostInfo__HostInfoError(C.QHostInfo_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHostInfo_ErrorString(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_FromName(name string) *QHostInfo {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_QHostInfo_FromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func (ptr *QHostInfo) FromName(name string) *QHostInfo {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_QHostInfo_FromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func (ptr *QHostInfo) HostName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHostInfo_HostName(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var memberC = C.CString(member)
	defer C.free(unsafe.Pointer(memberC))
	return int(int32(C.QHostInfo_QHostInfo_LookupHost(nameC, core.PointerFromQObject(receiver), memberC)))
}

func (ptr *QHostInfo) LookupHost(name string, receiver core.QObject_ITF, member string) int {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var memberC = C.CString(member)
	defer C.free(unsafe.Pointer(memberC))
	return int(int32(C.QHostInfo_QHostInfo_LookupHost(nameC, core.PointerFromQObject(receiver), memberC)))
}

func (ptr *QHostInfo) LookupId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHostInfo_LookupId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	if ptr.Pointer() != nil {
		var strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
		C.QHostInfo_SetErrorString(ptr.Pointer(), strC)
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QHostInfo_SetHostName(ptr.Pointer(), hostNameC)
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	if ptr.Pointer() != nil {
		C.QHostInfo_DestroyQHostInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QHostInfo_LocalDomainName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalDomainName())
}

func (ptr *QHostInfo) LocalDomainName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalDomainName())
}

func (ptr *QHostInfo) addresses_atList(i int) *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QHostInfo_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return ptr
}

func (ptr *QHttpMultiPart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHttpMultiPart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QHttpMultiPart__ContentType
//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	var tmpValue = NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.longlong(contentType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	var tmpValue = NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(ptr.Pointer(), PointerFromQHttpPart(httpPart))
	}
}

func (ptr *QHttpMultiPart) Boundary() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QHttpMultiPart_Boundary(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetBoundary(ptr.Pointer(), core.PointerFromQByteArray(boundary))
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(ptr.Pointer(), C.longlong(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHttpMultiPart_TimerEvent
func callbackQHttpMultiPart_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::timerEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::timerEvent")
	}
}

func (ptr *QHttpMultiPart) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHttpMultiPart) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHttpMultiPart_ChildEvent
func callbackQHttpMultiPart_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::childEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::childEvent")
	}
}

func (ptr *QHttpMultiPart) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHttpMultiPart) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHttpMultiPart_ConnectNotify
func callbackQHttpMultiPart_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::connectNotify", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::connectNotify")
	}
}

func (ptr *QHttpMultiPart) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHttpMultiPart_CustomEvent
func callbackQHttpMultiPart_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::customEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::customEvent")
	}
}

func (ptr *QHttpMultiPart) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHttpMultiPart) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHttpMultiPart_DeleteLater
func callbackQHttpMultiPart_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHttpMultiPartFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHttpMultiPart) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::deleteLater", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::deleteLater")
	}
}

func (ptr *QHttpMultiPart) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHttpMultiPart) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHttpMultiPart_DisconnectNotify
func callbackQHttpMultiPart_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHttpMultiPartFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::disconnectNotify", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::disconnectNotify")
	}
}

func (ptr *QHttpMultiPart) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHttpMultiPart) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHttpMultiPart_Event
func callbackQHttpMultiPart_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHttpMultiPartFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHttpMultiPart) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::event", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::event")
	}
}

func (ptr *QHttpMultiPart) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHttpMultiPart) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHttpMultiPart_EventFilter
func callbackQHttpMultiPart_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHttpMultiPartFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHttpMultiPart) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::eventFilter", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::eventFilter")
	}
}

func (ptr *QHttpMultiPart) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHttpMultiPart) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHttpMultiPart_MetaObject
func callbackQHttpMultiPart_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHttpMultiPartFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHttpMultiPart) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::metaObject", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::metaObject")
	}
}

func (ptr *QHttpMultiPart) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHttpMultiPart_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHttpMultiPart) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHttpMultiPart_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHttpPart struct {
	ptr unsafe.Pointer
}

type QHttpPart_ITF interface {
	QHttpPart_PTR() *QHttpPart
}

func (ptr *QHttpPart) QHttpPart_PTR() *QHttpPart {
	return ptr
}

func (ptr *QHttpPart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHttpPart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHttpPart(ptr QHttpPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpPartFromPointer(ptr unsafe.Pointer) *QHttpPart {
	var n = new(QHttpPart)
	n.SetPointer(ptr)
	return n
}
func NewQHttpPart() *QHttpPart {
	var tmpValue = NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart())
	runtime.SetFinalizer(tmpValue, (*QHttpPart).DestroyQHttpPart)
	return tmpValue
}

func NewQHttpPart2(other QHttpPart_ITF) *QHttpPart {
	var tmpValue = NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart2(PointerFromQHttpPart(other)))
	runtime.SetFinalizer(tmpValue, (*QHttpPart).DestroyQHttpPart)
	return tmpValue
}

func (ptr *QHttpPart) SetBody(body core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetBody(ptr.Pointer(), core.PointerFromQByteArray(body))
	}
}

func (ptr *QHttpPart) SetBodyDevice(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetBodyDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QHttpPart) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QHttpPart) Swap(other QHttpPart_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_Swap(ptr.Pointer(), PointerFromQHttpPart(other))
	}
}

func (ptr *QHttpPart) DestroyQHttpPart() {
	if ptr.Pointer() != nil {
		C.QHttpPart_DestroyQHttpPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLocalServer struct {
	core.QObject
}

type QLocalServer_ITF interface {
	core.QObject_ITF
	QLocalServer_PTR() *QLocalServer
}

func (ptr *QLocalServer) QLocalServer_PTR() *QLocalServer {
	return ptr
}

func (ptr *QLocalServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLocalServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQLocalServer(ptr QLocalServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalServer_PTR().Pointer()
	}
	return nil
}

func NewQLocalServerFromPointer(ptr unsafe.Pointer) *QLocalServer {
	var n = new(QLocalServer)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QLocalServer__SocketOption
//QLocalServer::SocketOption
type QLocalServer__SocketOption int64

const (
	QLocalServer__NoOptions         QLocalServer__SocketOption = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  QLocalServer__SocketOption = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x7)
)

func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {
	if ptr.Pointer() != nil {
		C.QLocalServer_SetSocketOptions(ptr.Pointer(), C.longlong(options))
	}
}

func NewQLocalServer(parent core.QObject_ITF) *QLocalServer {
	var tmpValue = NewQLocalServerFromPointer(C.QLocalServer_NewQLocalServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLocalServer) Close() {
	if ptr.Pointer() != nil {
		C.QLocalServer_Close(ptr.Pointer())
	}
}

func (ptr *QLocalServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) FullServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalServer_FullServerName(ptr.Pointer()))
	}
	return ""
}

//export callbackQLocalServer_HasPendingConnections
func callbackQLocalServer_HasPendingConnections(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::hasPendingConnections"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).HasPendingConnectionsDefault())))
}

func (ptr *QLocalServer) ConnectHasPendingConnections(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::hasPendingConnections", f)
	}
}

func (ptr *QLocalServer) DisconnectHasPendingConnections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::hasPendingConnections")
	}
}

func (ptr *QLocalServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) HasPendingConnectionsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnectionsDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalServer_IncomingConnection
func callbackQLocalServer_IncomingConnection(ptr unsafe.Pointer, socketDescriptor C.uintptr_t) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::incomingConnection"); signal != nil {
		signal.(func(uintptr))(uintptr(socketDescriptor))
	} else {
		NewQLocalServerFromPointer(ptr).IncomingConnectionDefault(uintptr(socketDescriptor))
	}
}

func (ptr *QLocalServer) ConnectIncomingConnection(f func(socketDescriptor uintptr)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::incomingConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectIncomingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::incomingConnection")
	}
}

func (ptr *QLocalServer) IncomingConnection(socketDescriptor uintptr) {
	if ptr.Pointer() != nil {
		C.QLocalServer_IncomingConnection(ptr.Pointer(), C.uintptr_t(socketDescriptor))
	}
}

func (ptr *QLocalServer) IncomingConnectionDefault(socketDescriptor uintptr) {
	if ptr.Pointer() != nil {
		C.QLocalServer_IncomingConnectionDefault(ptr.Pointer(), C.uintptr_t(socketDescriptor))
	}
}

func (ptr *QLocalServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) Listen(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QLocalServer_Listen(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QLocalServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLocalServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQLocalServer_NewConnection
func callbackQLocalServer_NewConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::newConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::newConnection")
	}
}

func (ptr *QLocalServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QLocalServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQLocalServer_NextPendingConnection
func callbackQLocalServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::nextPendingConnection"); signal != nil {
		return PointerFromQLocalSocket(signal.(func() *QLocalSocket)())
	}

	return PointerFromQLocalSocket(NewQLocalServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QLocalServer) ConnectNextPendingConnection(f func() *QLocalSocket) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::nextPendingConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNextPendingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::nextPendingConnection")
	}
}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) NextPendingConnectionDefault() *QLocalSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QLocalServer_RemoveServer(name string) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QLocalServer_QLocalServer_RemoveServer(nameC) != 0
}

func (ptr *QLocalServer) RemoveServer(name string) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QLocalServer_QLocalServer_RemoveServer(nameC) != 0
}

func (ptr *QLocalServer) ServerError() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QLocalServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) ServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QLocalServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {
	if ptr.Pointer() != nil {
		return QLocalServer__SocketOption(C.QLocalServer_SocketOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut bool) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_WaitForNewConnection(ptr.Pointer(), C.int(int32(msec)), C.char(int8(qt.GoBoolToInt(timedOut)))) != 0
	}
	return false
}

func (ptr *QLocalServer) DestroyQLocalServer() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DestroyQLocalServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLocalServer_TimerEvent
func callbackQLocalServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::timerEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::timerEvent")
	}
}

func (ptr *QLocalServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLocalServer_ChildEvent
func callbackQLocalServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::childEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::childEvent")
	}
}

func (ptr *QLocalServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLocalServer_ConnectNotify
func callbackQLocalServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::connectNotify", f)
	}
}

func (ptr *QLocalServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::connectNotify")
	}
}

func (ptr *QLocalServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalServer_CustomEvent
func callbackQLocalServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::customEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::customEvent")
	}
}

func (ptr *QLocalServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLocalServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLocalServer_DeleteLater
func callbackQLocalServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLocalServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::deleteLater", f)
	}
}

func (ptr *QLocalServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::deleteLater")
	}
}

func (ptr *QLocalServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLocalServer_DisconnectNotify
func callbackQLocalServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::disconnectNotify", f)
	}
}

func (ptr *QLocalServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::disconnectNotify")
	}
}

func (ptr *QLocalServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalServer_Event
func callbackQLocalServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLocalServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::event", f)
	}
}

func (ptr *QLocalServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::event")
	}
}

func (ptr *QLocalServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLocalServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLocalServer_EventFilter
func callbackQLocalServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLocalServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::eventFilter", f)
	}
}

func (ptr *QLocalServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::eventFilter")
	}
}

func (ptr *QLocalServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLocalServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLocalServer_MetaObject
func callbackQLocalServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLocalServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLocalServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::metaObject", f)
	}
}

func (ptr *QLocalServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::metaObject")
	}
}

func (ptr *QLocalServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLocalServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QLocalSocket struct {
	core.QIODevice
}

type QLocalSocket_ITF interface {
	core.QIODevice_ITF
	QLocalSocket_PTR() *QLocalSocket
}

func (ptr *QLocalSocket) QLocalSocket_PTR() *QLocalSocket {
	return ptr
}

func (ptr *QLocalSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QLocalSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQLocalSocket(ptr QLocalSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalSocket_PTR().Pointer()
	}
	return nil
}

func NewQLocalSocketFromPointer(ptr unsafe.Pointer) *QLocalSocket {
	var n = new(QLocalSocket)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QLocalSocket__LocalSocketError
//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int64

const (
	QLocalSocket__ConnectionRefusedError          QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__ConnectionRefusedError)
	QLocalSocket__PeerClosedError                 QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__RemoteHostClosedError)
	QLocalSocket__ServerNotFoundError             QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__HostNotFoundError)
	QLocalSocket__SocketAccessError               QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketAccessError)
	QLocalSocket__SocketResourceError             QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketResourceError)
	QLocalSocket__SocketTimeoutError              QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketTimeoutError)
	QLocalSocket__DatagramTooLargeError           QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__DatagramTooLargeError)
	QLocalSocket__ConnectionError                 QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__NetworkError)
	QLocalSocket__UnsupportedSocketOperationError QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__UnsupportedSocketOperationError)
	QLocalSocket__UnknownSocketError              QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__UnknownSocketError)
	QLocalSocket__OperationError                  QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__OperationError)
)

//go:generate stringer -type=QLocalSocket__LocalSocketState
//QLocalSocket::LocalSocketState
type QLocalSocket__LocalSocketState int64

const (
	QLocalSocket__UnconnectedState QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

//export callbackQLocalSocket_Open
func callbackQLocalSocket_Open(ptr unsafe.Pointer, openMode C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(openMode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(openMode)))))
}

func (ptr *QLocalSocket) ConnectOpen(f func(openMode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::open", f)
	}
}

func (ptr *QLocalSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::open")
	}
}

func (ptr *QLocalSocket) Open(openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(ptr.Pointer(), C.longlong(openMode)) != 0
	}
	return false
}

func (ptr *QLocalSocket) OpenDefault(openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_OpenDefault(ptr.Pointer(), C.longlong(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	var tmpValue = NewQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), nameC, C.longlong(openMode))
	}
}

//export callbackQLocalSocket_Connected
func callbackQLocalSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connected")
	}
}

func (ptr *QLocalSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Connected(ptr.Pointer())
	}
}

//export callbackQLocalSocket_Disconnected
func callbackQLocalSocket_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnected")
	}
}

func (ptr *QLocalSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQLocalSocket_Error2
func callbackQLocalSocket_Error2(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::error2"); signal != nil {
		signal.(func(QLocalSocket__LocalSocketError))(QLocalSocket__LocalSocketError(socketError))
	}

}

func (ptr *QLocalSocket) ConnectError2(f func(socketError QLocalSocket__LocalSocketError)) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::error2", f)
	}
}

func (ptr *QLocalSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::error2")
	}
}

func (ptr *QLocalSocket) Error2(socketError QLocalSocket__LocalSocketError) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Error2(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QLocalSocket) FullServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

//export callbackQLocalSocket_IsSequential
func callbackQLocalSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QLocalSocket) ConnectIsSequential(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::isSequential", f)
	}
}

func (ptr *QLocalSocket) DisconnectIsSequential() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::isSequential")
	}
}

func (ptr *QLocalSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QLocalSocket_SetServerName(ptr.Pointer(), nameC)
	}
}

func (ptr *QLocalSocket) State() QLocalSocket__LocalSocketState {
	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketState(C.QLocalSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_StateChanged
func callbackQLocalSocket_StateChanged(ptr unsafe.Pointer, socketState C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::stateChanged"); signal != nil {
		signal.(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
	}

}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::stateChanged")
	}
}

func (ptr *QLocalSocket) StateChanged(socketState QLocalSocket__LocalSocketState) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_StateChanged(ptr.Pointer(), C.longlong(socketState))
	}
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

//export callbackQLocalSocket_BytesAvailable
func callbackQLocalSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QLocalSocket) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesAvailable", f)
	}
}

func (ptr *QLocalSocket) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesAvailable")
	}
}

func (ptr *QLocalSocket) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_BytesToWrite
func callbackQLocalSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QLocalSocket) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesToWrite", f)
	}
}

func (ptr *QLocalSocket) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesToWrite")
	}
}

func (ptr *QLocalSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_CanReadLine
func callbackQLocalSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QLocalSocket) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::canReadLine", f)
	}
}

func (ptr *QLocalSocket) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::canReadLine")
	}
}

func (ptr *QLocalSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalSocket_Close
func callbackQLocalSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QLocalSocket) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::close", f)
	}
}

func (ptr *QLocalSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::close")
	}
}

func (ptr *QLocalSocket) Close() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Close(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(ptr.Pointer(), C.longlong(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_ReadData
func callbackQLocalSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, c C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(c)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQLocalSocketFromPointer(ptr).ReadDataDefault(&retS, int64(c)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QLocalSocket) ConnectReadData(f func(data *string, c int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::readData", f)
	}
}

func (ptr *QLocalSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::readData")
	}
}

func (ptr *QLocalSocket) ReadData(data *string, c int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(c)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QLocalSocket_ReadData(ptr.Pointer(), dataC, C.longlong(c)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QLocalSocket) ReadDataDefault(data *string, c int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(c)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QLocalSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(c)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QLocalSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQLocalSocket_WaitForBytesWritten
func callbackQLocalSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QLocalSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::waitForBytesWritten", f)
	}
}

func (ptr *QLocalSocket) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::waitForBytesWritten")
	}
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQLocalSocket_WaitForReadyRead
func callbackQLocalSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QLocalSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::waitForReadyRead", f)
	}
}

func (ptr *QLocalSocket) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::waitForReadyRead")
	}
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQLocalSocket_WriteData
func callbackQLocalSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, c C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(c)))
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(c)))
}

func (ptr *QLocalSocket) ConnectWriteData(f func(data string, c int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::writeData", f)
	}
}

func (ptr *QLocalSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::writeData")
	}
}

func (ptr *QLocalSocket) WriteData(data string, c int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_WriteData(ptr.Pointer(), dataC, C.longlong(c)))
	}
	return 0
}

func (ptr *QLocalSocket) WriteDataDefault(data string, c int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(c)))
	}
	return 0
}

//export callbackQLocalSocket_AtEnd
func callbackQLocalSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QLocalSocket) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::atEnd", f)
	}
}

func (ptr *QLocalSocket) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::atEnd")
	}
}

func (ptr *QLocalSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalSocket_Pos
func callbackQLocalSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).PosDefault())
}

func (ptr *QLocalSocket) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::pos", f)
	}
}

func (ptr *QLocalSocket) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::pos")
	}
}

func (ptr *QLocalSocket) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_ReadLineData
func callbackQLocalSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QLocalSocket) ConnectReadLineData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::readLineData", f)
	}
}

func (ptr *QLocalSocket) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::readLineData")
	}
}

func (ptr *QLocalSocket) ReadLineData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QLocalSocket) ReadLineDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQLocalSocket_Reset
func callbackQLocalSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QLocalSocket) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::reset", f)
	}
}

func (ptr *QLocalSocket) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::reset")
	}
}

func (ptr *QLocalSocket) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalSocket_Seek
func callbackQLocalSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QLocalSocket) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::seek", f)
	}
}

func (ptr *QLocalSocket) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::seek")
	}
}

func (ptr *QLocalSocket) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QLocalSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQLocalSocket_Size
func callbackQLocalSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QLocalSocket) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::size", f)
	}
}

func (ptr *QLocalSocket) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::size")
	}
}

func (ptr *QLocalSocket) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_TimerEvent
func callbackQLocalSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::timerEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::timerEvent")
	}
}

func (ptr *QLocalSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLocalSocket_ChildEvent
func callbackQLocalSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::childEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::childEvent")
	}
}

func (ptr *QLocalSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLocalSocket_ConnectNotify
func callbackQLocalSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connectNotify", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connectNotify")
	}
}

func (ptr *QLocalSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalSocket_CustomEvent
func callbackQLocalSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::customEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::customEvent")
	}
}

func (ptr *QLocalSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLocalSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLocalSocket_DeleteLater
func callbackQLocalSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLocalSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::deleteLater", f)
	}
}

func (ptr *QLocalSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::deleteLater")
	}
}

func (ptr *QLocalSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLocalSocket_DisconnectNotify
func callbackQLocalSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnectNotify", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnectNotify")
	}
}

func (ptr *QLocalSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalSocket_Event
func callbackQLocalSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLocalSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::event", f)
	}
}

func (ptr *QLocalSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::event")
	}
}

func (ptr *QLocalSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLocalSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLocalSocket_EventFilter
func callbackQLocalSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLocalSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::eventFilter", f)
	}
}

func (ptr *QLocalSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::eventFilter")
	}
}

func (ptr *QLocalSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLocalSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLocalSocket_MetaObject
func callbackQLocalSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLocalSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLocalSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::metaObject", f)
	}
}

func (ptr *QLocalSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::metaObject")
	}
}

func (ptr *QLocalSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLocalSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkAccessManager struct {
	core.QObject
}

type QNetworkAccessManager_ITF interface {
	core.QObject_ITF
	QNetworkAccessManager_PTR() *QNetworkAccessManager
}

func (ptr *QNetworkAccessManager) QNetworkAccessManager_PTR() *QNetworkAccessManager {
	return ptr
}

func (ptr *QNetworkAccessManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkAccessManager(ptr QNetworkAccessManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAccessManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAccessManagerFromPointer(ptr unsafe.Pointer) *QNetworkAccessManager {
	var n = new(QNetworkAccessManager)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkAccessManager__NetworkAccessibility
//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(1)
)

//go:generate stringer -type=QNetworkAccessManager__Operation
//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int64

const (
	QNetworkAccessManager__HeadOperation    QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(1)
	QNetworkAccessManager__GetOperation     QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(2)
	QNetworkAccessManager__PutOperation     QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(3)
	QNetworkAccessManager__PostOperation    QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(4)
	QNetworkAccessManager__DeleteOperation  QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(5)
	QNetworkAccessManager__CustomOperation  QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(6)
	QNetworkAccessManager__UnknownOperation QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(0)
)

func (ptr *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {
	if ptr.Pointer() != nil {
		return NewQNetworkProxyFactoryFromPointer(C.QNetworkAccessManager_ProxyFactory(ptr.Pointer()))
	}
	return nil
}

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {
	var tmpValue = NewQNetworkAccessManagerFromPointer(C.QNetworkAccessManager_NewQNetworkAccessManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkAccessManager) ActiveConfiguration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkAccessManager_ActiveConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_AuthenticationRequired
func callbackQNetworkAccessManager_AuthenticationRequired(ptr unsafe.Pointer, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::authenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::authenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::authenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF, authenticator QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_AuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractNetworkCacheFromPointer(C.QNetworkAccessManager_Cache(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearAccessCache(ptr.Pointer())
	}
}

func (ptr *QNetworkAccessManager) Configuration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkAccessManager_Configuration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectToHost(hostName string, port uint16) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QNetworkAccessManager_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port))
	}
}

func (ptr *QNetworkAccessManager) ConnectToHostEncrypted(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QNetworkAccessManager_ConnectToHostEncrypted(ptr.Pointer(), hostNameC, C.ushort(port), PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCookieJarFromPointer(C.QNetworkAccessManager_CookieJar(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_CreateRequest
func callbackQNetworkAccessManager_CreateRequest(ptr unsafe.Pointer, op C.longlong, req unsafe.Pointer, outgoingData unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::createRequest"); signal != nil {
		return PointerFromQNetworkReply(signal.(func(QNetworkAccessManager__Operation, *QNetworkRequest, *core.QIODevice) *QNetworkReply)(QNetworkAccessManager__Operation(op), NewQNetworkRequestFromPointer(req), core.NewQIODeviceFromPointer(outgoingData)))
	}

	return PointerFromQNetworkReply(NewQNetworkAccessManagerFromPointer(ptr).CreateRequestDefault(QNetworkAccessManager__Operation(op), NewQNetworkRequestFromPointer(req), core.NewQIODeviceFromPointer(outgoingData)))
}

func (ptr *QNetworkAccessManager) ConnectCreateRequest(f func(op QNetworkAccessManager__Operation, req *QNetworkRequest, outgoingData *core.QIODevice) *QNetworkReply) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::createRequest", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectCreateRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::createRequest")
	}
}

func (ptr *QNetworkAccessManager) CreateRequest(op QNetworkAccessManager__Operation, req QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequest(ptr.Pointer(), C.longlong(op), PointerFromQNetworkRequest(req), core.PointerFromQIODevice(outgoingData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) CreateRequestDefault(op QNetworkAccessManager__Operation, req QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequestDefault(ptr.Pointer(), C.longlong(op), PointerFromQNetworkRequest(req), core.PointerFromQIODevice(outgoingData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_DeleteResource(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_Encrypted
func callbackQNetworkAccessManager_Encrypted(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::encrypted"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::encrypted", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::encrypted")
	}
}

func (ptr *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Encrypted(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

//export callbackQNetworkAccessManager_Finished
func callbackQNetworkAccessManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::finished"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::finished", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::finished")
	}
}

func (ptr *QNetworkAccessManager) Finished(reply QNetworkReply_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Finished(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Get(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Head(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkAccessManager_NetworkAccessibleChanged
func callbackQNetworkAccessManager_NetworkAccessibleChanged(ptr unsafe.Pointer, accessible C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::networkAccessibleChanged"); signal != nil {
		signal.(func(QNetworkAccessManager__NetworkAccessibility))(QNetworkAccessManager__NetworkAccessibility(accessible))
	}

}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::networkAccessibleChanged", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::networkAccessibleChanged")
	}
}

func (ptr *QNetworkAccessManager) NetworkAccessibleChanged(accessible QNetworkAccessManager__NetworkAccessibility) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_NetworkAccessibleChanged(ptr.Pointer(), C.longlong(accessible))
	}
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF, authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Proxy() *QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkAccessManager_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_ProxyAuthenticationRequired
func callbackQNetworkAccessManager_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::proxyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkProxy, *QAuthenticator))(NewQNetworkProxyFromPointer(proxy), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::proxyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::proxyAuthenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ProxyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkProxy(proxy), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCache(ptr.Pointer(), PointerFromQAbstractNetworkCache(cache))
	}
}

func (ptr *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCookieJar(ptr.Pointer(), PointerFromQNetworkCookieJar(cookieJar))
	}
}

func (ptr *QNetworkAccessManager) SetNetworkAccessible(accessible QNetworkAccessManager__NetworkAccessibility) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetNetworkAccessible(ptr.Pointer(), C.longlong(accessible))
	}
}

func (ptr *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(proxy))
	}
}

func (ptr *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxyFactory(ptr.Pointer(), PointerFromQNetworkProxyFactory(factory))
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QNetworkAccessManager_SupportedSchemes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQNetworkAccessManager_SupportedSchemesImplementation
func callbackQNetworkAccessManager_SupportedSchemesImplementation(ptr unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::supportedSchemesImplementation"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQNetworkAccessManagerFromPointer(ptr).SupportedSchemesImplementationDefault(), "|"))
}

func (ptr *QNetworkAccessManager) ConnectSupportedSchemesImplementation(f func() []string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::supportedSchemesImplementation", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectSupportedSchemesImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::supportedSchemesImplementation")
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementation() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QNetworkAccessManager_SupportedSchemesImplementation(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementationDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QNetworkAccessManager_SupportedSchemesImplementationDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkAccessManager_TimerEvent
func callbackQNetworkAccessManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::timerEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::timerEvent")
	}
}

func (ptr *QNetworkAccessManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkAccessManager_ChildEvent
func callbackQNetworkAccessManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::childEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::childEvent")
	}
}

func (ptr *QNetworkAccessManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkAccessManager_ConnectNotify
func callbackQNetworkAccessManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::connectNotify", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::connectNotify")
	}
}

func (ptr *QNetworkAccessManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkAccessManager_CustomEvent
func callbackQNetworkAccessManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::customEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::customEvent")
	}
}

func (ptr *QNetworkAccessManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkAccessManager_DeleteLater
func callbackQNetworkAccessManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkAccessManager) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::deleteLater", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::deleteLater")
	}
}

func (ptr *QNetworkAccessManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkAccessManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkAccessManager_DisconnectNotify
func callbackQNetworkAccessManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::disconnectNotify", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::disconnectNotify")
	}
}

func (ptr *QNetworkAccessManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkAccessManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkAccessManager_Event
func callbackQNetworkAccessManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkAccessManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkAccessManager) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::event", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::event")
	}
}

func (ptr *QNetworkAccessManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkAccessManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkAccessManager_EventFilter
func callbackQNetworkAccessManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkAccessManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkAccessManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::eventFilter", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::eventFilter")
	}
}

func (ptr *QNetworkAccessManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkAccessManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkAccessManager_MetaObject
func callbackQNetworkAccessManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkAccessManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkAccessManager) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::metaObject", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::metaObject")
	}
}

func (ptr *QNetworkAccessManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkAccessManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkAccessManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkAddressEntry struct {
	ptr unsafe.Pointer
}

type QNetworkAddressEntry_ITF interface {
	QNetworkAddressEntry_PTR() *QNetworkAddressEntry
}

func (ptr *QNetworkAddressEntry) QNetworkAddressEntry_PTR() *QNetworkAddressEntry {
	return ptr
}

func (ptr *QNetworkAddressEntry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkAddressEntry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkAddressEntry(ptr QNetworkAddressEntry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAddressEntry_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) *QNetworkAddressEntry {
	var n = new(QNetworkAddressEntry)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	var tmpValue = NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry())
	runtime.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
	return tmpValue
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {
	var tmpValue = NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry2(PointerFromQNetworkAddressEntry(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
	return tmpValue
}

func (ptr *QNetworkAddressEntry) Broadcast() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkAddressEntry_Broadcast(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) Ip() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkAddressEntry_Ip(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) Netmask() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkAddressEntry_Netmask(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) PrefixLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkAddressEntry_PrefixLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetBroadcast(ptr.Pointer(), PointerFromQHostAddress(newBroadcast))
	}
}

func (ptr *QNetworkAddressEntry) SetIp(newIp QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetIp(ptr.Pointer(), PointerFromQHostAddress(newIp))
	}
}

func (ptr *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetNetmask(ptr.Pointer(), PointerFromQHostAddress(newNetmask))
	}
}

func (ptr *QNetworkAddressEntry) SetPrefixLength(length int) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetPrefixLength(ptr.Pointer(), C.int(int32(length)))
	}
}

func (ptr *QNetworkAddressEntry) Swap(other QNetworkAddressEntry_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_Swap(ptr.Pointer(), PointerFromQNetworkAddressEntry(other))
	}
}

func (ptr *QNetworkAddressEntry) DestroyQNetworkAddressEntry() {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_DestroyQNetworkAddressEntry(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkCacheMetaData struct {
	ptr unsafe.Pointer
}

type QNetworkCacheMetaData_ITF interface {
	QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData
}

func (ptr *QNetworkCacheMetaData) QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData {
	return ptr
}

func (ptr *QNetworkCacheMetaData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkCacheMetaData(ptr QNetworkCacheMetaData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) *QNetworkCacheMetaData {
	var n = new(QNetworkCacheMetaData)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData())
	runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
	return tmpValue
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {
	var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(PointerFromQNetworkCacheMetaData(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
	return tmpValue
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_ExpirationDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_LastModified(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_SaveToDisk(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetLastModified(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetSaveToDisk(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allow))))
	}
}

func (ptr *QNetworkCacheMetaData) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_Swap(ptr.Pointer(), PointerFromQNetworkCacheMetaData(other))
	}
}

func (ptr *QNetworkCacheMetaData) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkCacheMetaData_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkConfiguration struct {
	ptr unsafe.Pointer
}

type QNetworkConfiguration_ITF interface {
	QNetworkConfiguration_PTR() *QNetworkConfiguration
}

func (ptr *QNetworkConfiguration) QNetworkConfiguration_PTR() *QNetworkConfiguration {
	return ptr
}

func (ptr *QNetworkConfiguration) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkConfiguration) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkConfiguration(ptr QNetworkConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationFromPointer(ptr unsafe.Pointer) *QNetworkConfiguration {
	var n = new(QNetworkConfiguration)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkConfiguration__BearerType
//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int64

const (
	QNetworkConfiguration__BearerUnknown   QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(0)
	QNetworkConfiguration__BearerEthernet  QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(1)
	QNetworkConfiguration__BearerWLAN      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(2)
	QNetworkConfiguration__Bearer2G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(3)
	QNetworkConfiguration__BearerCDMA2000  QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(4)
	QNetworkConfiguration__BearerWCDMA     QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(5)
	QNetworkConfiguration__BearerHSPA      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(6)
	QNetworkConfiguration__BearerBluetooth QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(7)
	QNetworkConfiguration__BearerWiMAX     QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(8)
	QNetworkConfiguration__BearerEVDO      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(9)
	QNetworkConfiguration__BearerLTE       QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(10)
	QNetworkConfiguration__Bearer3G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(11)
	QNetworkConfiguration__Bearer4G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(12)
)

//go:generate stringer -type=QNetworkConfiguration__Purpose
//QNetworkConfiguration::Purpose
type QNetworkConfiguration__Purpose int64

const (
	QNetworkConfiguration__UnknownPurpose         QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(3)
)

//go:generate stringer -type=QNetworkConfiguration__StateFlag
//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int64

const (
	QNetworkConfiguration__Undefined  QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x000000e)
)

//go:generate stringer -type=QNetworkConfiguration__Type
//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint QNetworkConfiguration__Type = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      QNetworkConfiguration__Type = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          QNetworkConfiguration__Type = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             QNetworkConfiguration__Type = QNetworkConfiguration__Type(3)
)

func NewQNetworkConfiguration() *QNetworkConfiguration {
	var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration())
	runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
	return tmpValue
}

func NewQNetworkConfiguration2(other QNetworkConfiguration_ITF) *QNetworkConfiguration {
	var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration2(PointerFromQNetworkConfiguration(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
	return tmpValue
}

func (ptr *QNetworkConfiguration) BearerType() QNetworkConfiguration__BearerType {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeFamily() QNetworkConfiguration__BearerType {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerTypeFamily(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkConfiguration_BearerTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Children() []*QNetworkConfiguration {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkConfiguration {
			var out = make([]*QNetworkConfiguration, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkConfigurationFromPointer(l.data).children_atList(i)
			}
			return out
		}(C.QNetworkConfiguration_Children(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkConfiguration) Identifier() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkConfiguration_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsRoamingAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkConfiguration_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Purpose() QNetworkConfiguration__Purpose {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Purpose(C.QNetworkConfiguration_Purpose(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) State() QNetworkConfiguration__StateFlag {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__StateFlag(C.QNetworkConfiguration_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) Swap(other QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_Swap(ptr.Pointer(), PointerFromQNetworkConfiguration(other))
	}
}

func (ptr *QNetworkConfiguration) Type() QNetworkConfiguration__Type {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Type(C.QNetworkConfiguration_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) DestroyQNetworkConfiguration() {
	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_DestroyQNetworkConfiguration(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfiguration) children_atList(i int) *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_children_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

type QNetworkConfigurationManager struct {
	core.QObject
}

type QNetworkConfigurationManager_ITF interface {
	core.QObject_ITF
	QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager
}

func (ptr *QNetworkConfigurationManager) QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager {
	return ptr
}

func (ptr *QNetworkConfigurationManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkConfigurationManager(ptr QNetworkConfigurationManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) *QNetworkConfigurationManager {
	var n = new(QNetworkConfigurationManager)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkConfigurationManager__Capability
//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int64

const (
	QNetworkConfigurationManager__CanStartAndStopInterfaces QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000040)
)

func NewQNetworkConfigurationManager(parent core.QObject_ITF) *QNetworkConfigurationManager {
	var tmpValue = NewQNetworkConfigurationManagerFromPointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkConfigurationManager) AllConfigurations(filter QNetworkConfiguration__StateFlag) []*QNetworkConfiguration {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkConfiguration {
			var out = make([]*QNetworkConfiguration, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkConfigurationManagerFromPointer(l.data).allConfigurations_atList(i)
			}
			return out
		}(C.QNetworkConfigurationManager_AllConfigurations(ptr.Pointer(), C.longlong(filter)))
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkConfigurationManager_ConfigurationAdded
func callbackQNetworkConfigurationManager_ConfigurationAdded(ptr unsafe.Pointer, config unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::configurationAdded"); signal != nil {
		signal.(func(*QNetworkConfiguration))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationAdded(f func(config *QNetworkConfiguration)) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectConfigurationAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationAdded", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationAdded() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationAdded")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationAdded(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationAdded(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

//export callbackQNetworkConfigurationManager_ConfigurationChanged
func callbackQNetworkConfigurationManager_ConfigurationChanged(ptr unsafe.Pointer, config unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::configurationChanged"); signal != nil {
		signal.(func(*QNetworkConfiguration))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationChanged(f func(config *QNetworkConfiguration)) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectConfigurationChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationChanged")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationChanged(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationChanged(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationFromIdentifier(identifier string) *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_ConfigurationFromIdentifier(ptr.Pointer(), identifierC))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkConfigurationManager_ConfigurationRemoved
func callbackQNetworkConfigurationManager_ConfigurationRemoved(ptr unsafe.Pointer, config unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::configurationRemoved"); signal != nil {
		signal.(func(*QNetworkConfiguration))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationRemoved(f func(config *QNetworkConfiguration)) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectConfigurationRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationRemoved", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationRemoved() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationRemoved")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationRemoved(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationRemoved(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkConfigurationManager) DefaultConfiguration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_DefaultConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_IsOnline(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_OnlineStateChanged
func callbackQNetworkConfigurationManager_OnlineStateChanged(ptr unsafe.Pointer, isOnline C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::onlineStateChanged"); signal != nil {
		signal.(func(bool))(int8(isOnline) != 0)
	}

}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectOnlineStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::onlineStateChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::onlineStateChanged")
	}
}

func (ptr *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_OnlineStateChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isOnline))))
	}
}

//export callbackQNetworkConfigurationManager_UpdateCompleted
func callbackQNetworkConfigurationManager_UpdateCompleted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::updateCompleted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectUpdateCompleted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateCompleted", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateCompleted")
	}
}

func (ptr *QNetworkConfigurationManager) UpdateCompleted() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateCompleted(ptr.Pointer())
	}
}

//export callbackQNetworkConfigurationManager_UpdateConfigurations
func callbackQNetworkConfigurationManager_UpdateConfigurations(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::updateConfigurations"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).UpdateConfigurationsDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateConfigurations(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateConfigurations", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateConfigurations() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateConfigurations")
	}
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurations(ptr.Pointer())
	}
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurationsDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurationsDefault(ptr.Pointer())
	}
}

//export callbackQNetworkConfigurationManager_DestroyQNetworkConfigurationManager
func callbackQNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::~QNetworkConfigurationManager"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DestroyQNetworkConfigurationManagerDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDestroyQNetworkConfigurationManager(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::~QNetworkConfigurationManager", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDestroyQNetworkConfigurationManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::~QNetworkConfigurationManager")
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManagerDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManagerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) allConfigurations_atList(i int) *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_allConfigurations_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkConfigurationManager_TimerEvent
func callbackQNetworkConfigurationManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::timerEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::timerEvent")
	}
}

func (ptr *QNetworkConfigurationManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_ChildEvent
func callbackQNetworkConfigurationManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::childEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::childEvent")
	}
}

func (ptr *QNetworkConfigurationManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_ConnectNotify
func callbackQNetworkConfigurationManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::connectNotify", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::connectNotify")
	}
}

func (ptr *QNetworkConfigurationManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkConfigurationManager_CustomEvent
func callbackQNetworkConfigurationManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::customEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::customEvent")
	}
}

func (ptr *QNetworkConfigurationManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_DeleteLater
func callbackQNetworkConfigurationManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::deleteLater", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::deleteLater")
	}
}

func (ptr *QNetworkConfigurationManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkConfigurationManager_DisconnectNotify
func callbackQNetworkConfigurationManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::disconnectNotify", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::disconnectNotify")
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkConfigurationManager_Event
func callbackQNetworkConfigurationManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkConfigurationManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkConfigurationManager) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::event", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::event")
	}
}

func (ptr *QNetworkConfigurationManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_EventFilter
func callbackQNetworkConfigurationManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkConfigurationManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkConfigurationManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::eventFilter", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::eventFilter")
	}
}

func (ptr *QNetworkConfigurationManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_MetaObject
func callbackQNetworkConfigurationManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkConfigurationManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkConfigurationManager) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::metaObject", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::metaObject")
	}
}

func (ptr *QNetworkConfigurationManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkConfigurationManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkConfigurationManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkCookie struct {
	ptr unsafe.Pointer
}

type QNetworkCookie_ITF interface {
	QNetworkCookie_PTR() *QNetworkCookie
}

func (ptr *QNetworkCookie) QNetworkCookie_PTR() *QNetworkCookie {
	return ptr
}

func (ptr *QNetworkCookie) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkCookie) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkCookie(ptr QNetworkCookie_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookie_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieFromPointer(ptr unsafe.Pointer) *QNetworkCookie {
	var n = new(QNetworkCookie)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkCookie__RawForm
//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             QNetworkCookie__RawForm = QNetworkCookie__RawForm(1)
)

func NewQNetworkCookie(name core.QByteArray_ITF, value core.QByteArray_ITF) *QNetworkCookie {
	var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie(core.PointerFromQByteArray(name), core.PointerFromQByteArray(value)))
	runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
	return tmpValue
}

func NewQNetworkCookie2(other QNetworkCookie_ITF) *QNetworkCookie {
	var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie2(PointerFromQNetworkCookie(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
	return tmpValue
}

func (ptr *QNetworkCookie) Domain() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkCookie_Domain(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) ExpirationDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QNetworkCookie_ExpirationDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_HasSameIdentifier(ptr.Pointer(), PointerFromQNetworkCookie(other)) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsHttpOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSecure() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSecure(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSessionCookie(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) Name() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkCookie_Name(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) Normalize(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_Normalize(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QNetworkCookie_ParseCookies(cookieString core.QByteArray_ITF) []*QNetworkCookie {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
		var out = make([]*QNetworkCookie, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkCookieFromPointer(l.data).parseCookies_atList(i)
		}
		return out
	}(C.QNetworkCookie_QNetworkCookie_ParseCookies(core.PointerFromQByteArray(cookieString)))
}

func (ptr *QNetworkCookie) ParseCookies(cookieString core.QByteArray_ITF) []*QNetworkCookie {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
		var out = make([]*QNetworkCookie, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkCookieFromPointer(l.data).parseCookies_atList(i)
		}
		return out
	}(C.QNetworkCookie_QNetworkCookie_ParseCookies(core.PointerFromQByteArray(cookieString)))
}

func (ptr *QNetworkCookie) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkCookie_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) SetDomain(domain string) {
	if ptr.Pointer() != nil {
		var domainC = C.CString(domain)
		defer C.free(unsafe.Pointer(domainC))
		C.QNetworkCookie_SetDomain(ptr.Pointer(), domainC)
	}
}

func (ptr *QNetworkCookie) SetExpirationDate(date core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(date))
	}
}

func (ptr *QNetworkCookie) SetHttpOnly(enable bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetHttpOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QNetworkCookie) SetName(cookieName core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetName(ptr.Pointer(), core.PointerFromQByteArray(cookieName))
	}
}

func (ptr *QNetworkCookie) SetPath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QNetworkCookie_SetPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QNetworkCookie) SetSecure(enable bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetSecure(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QNetworkCookie) SetValue(value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QNetworkCookie) Swap(other QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_Swap(ptr.Pointer(), PointerFromQNetworkCookie(other))
	}
}

func (ptr *QNetworkCookie) ToRawForm(form QNetworkCookie__RawForm) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkCookie_ToRawForm(ptr.Pointer(), C.longlong(form)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkCookie_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_DestroyQNetworkCookie(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookie) parseCookies_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookie_parseCookies_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJar_ITF interface {
	core.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func (ptr *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar {
	return ptr
}

func (ptr *QNetworkCookieJar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkCookieJar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJar_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {
	var tmpValue = NewQNetworkCookieJarFromPointer(C.QNetworkCookieJar_NewQNetworkCookieJar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkCookieJar) AllCookies() []*QNetworkCookie {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
			var out = make([]*QNetworkCookie, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkCookieJarFromPointer(l.data).allCookies_atList(i)
			}
			return out
		}(C.QNetworkCookieJar_AllCookies(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookieJar) CookiesForUrl(url core.QUrl_ITF) []*QNetworkCookie {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
			var out = make([]*QNetworkCookie, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkCookieJarFromPointer(l.data).cookiesForUrl_atList(i)
			}
			return out
		}(C.QNetworkCookieJar_CookiesForUrl(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

//export callbackQNetworkCookieJar_DeleteCookie
func callbackQNetworkCookieJar_DeleteCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::deleteCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie) bool)(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).DeleteCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectDeleteCookie(f func(cookie *QNetworkCookie) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDeleteCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteCookie")
	}
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DeleteCookieDefault(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_InsertCookie
func callbackQNetworkCookieJar_InsertCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::insertCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie) bool)(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).InsertCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectInsertCookie(f func(cookie *QNetworkCookie) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::insertCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectInsertCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::insertCookie")
	}
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookieDefault(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_UpdateCookie
func callbackQNetworkCookieJar_UpdateCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::updateCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie) bool)(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).UpdateCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectUpdateCookie(f func(cookie *QNetworkCookie) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::updateCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectUpdateCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::updateCookie")
	}
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookieDefault(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_ValidateCookie
func callbackQNetworkCookieJar_ValidateCookie(ptr unsafe.Pointer, cookie unsafe.Pointer, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::validateCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie, *core.QUrl) bool)(NewQNetworkCookieFromPointer(cookie), core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).ValidateCookieDefault(NewQNetworkCookieFromPointer(cookie), core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkCookieJar) ConnectValidateCookie(f func(cookie *QNetworkCookie, url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::validateCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectValidateCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::validateCookie")
	}
}

func (ptr *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_ValidateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) ValidateCookieDefault(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_ValidateCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_DestroyQNetworkCookieJar
func callbackQNetworkCookieJar_DestroyQNetworkCookieJar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::~QNetworkCookieJar"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DestroyQNetworkCookieJarDefault()
	}
}

func (ptr *QNetworkCookieJar) ConnectDestroyQNetworkCookieJar(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::~QNetworkCookieJar", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDestroyQNetworkCookieJar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::~QNetworkCookieJar")
	}
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJarDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJarDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) allCookies_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookieJar_allCookies_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) cookiesForUrl_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookieJar_cookiesForUrl_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkCookieJar_TimerEvent
func callbackQNetworkCookieJar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::timerEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::timerEvent")
	}
}

func (ptr *QNetworkCookieJar) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkCookieJar_ChildEvent
func callbackQNetworkCookieJar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::childEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::childEvent")
	}
}

func (ptr *QNetworkCookieJar) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkCookieJar_ConnectNotify
func callbackQNetworkCookieJar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::connectNotify", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::connectNotify")
	}
}

func (ptr *QNetworkCookieJar) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkCookieJar_CustomEvent
func callbackQNetworkCookieJar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::customEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::customEvent")
	}
}

func (ptr *QNetworkCookieJar) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkCookieJar_DeleteLater
func callbackQNetworkCookieJar_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkCookieJar) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteLater", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteLater")
	}
}

func (ptr *QNetworkCookieJar) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkCookieJar_DisconnectNotify
func callbackQNetworkCookieJar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::disconnectNotify", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::disconnectNotify")
	}
}

func (ptr *QNetworkCookieJar) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkCookieJar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkCookieJar_Event
func callbackQNetworkCookieJar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkCookieJar) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::event", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::event")
	}
}

func (ptr *QNetworkCookieJar) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_EventFilter
func callbackQNetworkCookieJar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkCookieJar) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::eventFilter", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::eventFilter")
	}
}

func (ptr *QNetworkCookieJar) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_MetaObject
func callbackQNetworkCookieJar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkCookieJarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkCookieJar) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::metaObject", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::metaObject")
	}
}

func (ptr *QNetworkCookieJar) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkCookieJar_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookieJar) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkCookieJar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkDiskCache struct {
	QAbstractNetworkCache
}

type QNetworkDiskCache_ITF interface {
	QAbstractNetworkCache_ITF
	QNetworkDiskCache_PTR() *QNetworkDiskCache
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return ptr
}

func (ptr *QNetworkDiskCache) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractNetworkCache_PTR().SetPointer(p)
	}
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
	return n
}
func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	var tmpValue = NewQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkDiskCache_CacheDirectory(ptr.Pointer()))
	}
	return ""
}

//export callbackQNetworkDiskCache_CacheSize
func callbackQNetworkDiskCache_CacheSize(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::cacheSize"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkDiskCacheFromPointer(ptr).CacheSizeDefault())
}

func (ptr *QNetworkDiskCache) ConnectCacheSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::cacheSize", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectCacheSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::cacheSize")
	}
}

func (ptr *QNetworkDiskCache) CacheSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) CacheSizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_CacheSizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkDiskCache_Clear
func callbackQNetworkDiskCache_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::clear"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QNetworkDiskCache) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::clear", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::clear")
	}
}

func (ptr *QNetworkDiskCache) Clear() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ClearDefault(ptr.Pointer())
	}
}

//export callbackQNetworkDiskCache_Data
func callbackQNetworkDiskCache_Data(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::data"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*core.QUrl) *core.QIODevice)(core.NewQUrlFromPointer(url)))
	}

	return core.PointerFromQIODevice(NewQNetworkDiskCacheFromPointer(ptr).DataDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QNetworkDiskCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::data", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::data")
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) DataDefault(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_DataDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Expire
func callbackQNetworkDiskCache_Expire(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::expire"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkDiskCacheFromPointer(ptr).ExpireDefault())
}

func (ptr *QNetworkDiskCache) ConnectExpire(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::expire", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectExpire() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::expire")
	}
}

func (ptr *QNetworkDiskCache) Expire() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_Expire(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) ExpireDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_ExpireDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) FileMetaData(fileName string) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_FileMetaData(ptr.Pointer(), fileNameC))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Insert
func callbackQNetworkDiskCache_Insert(ptr unsafe.Pointer, device unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::insert"); signal != nil {
		signal.(func(*core.QIODevice))(core.NewQIODeviceFromPointer(device))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).InsertDefault(core.NewQIODeviceFromPointer(device))
	}
}

func (ptr *QNetworkDiskCache) ConnectInsert(f func(device *core.QIODevice)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::insert", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectInsert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::insert")
	}
}

func (ptr *QNetworkDiskCache) Insert(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Insert(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNetworkDiskCache) InsertDefault(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_InsertDefault(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNetworkDiskCache) MaximumCacheSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_MaximumCacheSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkDiskCache_MetaData
func callbackQNetworkDiskCache_MetaData(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::metaData"); signal != nil {
		return PointerFromQNetworkCacheMetaData(signal.(func(*core.QUrl) *QNetworkCacheMetaData)(core.NewQUrlFromPointer(url)))
	}

	return PointerFromQNetworkCacheMetaData(NewQNetworkDiskCacheFromPointer(ptr).MetaDataDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QNetworkDiskCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaData", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaData")
	}
}

func (ptr *QNetworkDiskCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_MetaData(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) MetaDataDefault(url core.QUrl_ITF) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_MetaDataDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Prepare
func callbackQNetworkDiskCache_Prepare(ptr unsafe.Pointer, metaData unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::prepare"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*QNetworkCacheMetaData) *core.QIODevice)(NewQNetworkCacheMetaDataFromPointer(metaData)))
	}

	return core.PointerFromQIODevice(NewQNetworkDiskCacheFromPointer(ptr).PrepareDefault(NewQNetworkCacheMetaDataFromPointer(metaData)))
}

func (ptr *QNetworkDiskCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::prepare", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectPrepare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::prepare")
	}
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) PrepareDefault(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_PrepareDefault(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Remove
func callbackQNetworkDiskCache_Remove(ptr unsafe.Pointer, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::remove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl) bool)(core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).RemoveDefault(core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkDiskCache) ConnectRemove(f func(url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::remove", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectRemove() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::remove")
	}
}

func (ptr *QNetworkDiskCache) Remove(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_Remove(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) RemoveDefault(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_RemoveDefault(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	if ptr.Pointer() != nil {
		var cacheDirC = C.CString(cacheDir)
		defer C.free(unsafe.Pointer(cacheDirC))
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), cacheDirC)
	}
}

func (ptr *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetMaximumCacheSize(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQNetworkDiskCache_UpdateMetaData
func callbackQNetworkDiskCache_UpdateMetaData(ptr unsafe.Pointer, metaData unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::updateMetaData"); signal != nil {
		signal.(func(*QNetworkCacheMetaData))(NewQNetworkCacheMetaDataFromPointer(metaData))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).UpdateMetaDataDefault(NewQNetworkCacheMetaDataFromPointer(metaData))
	}
}

func (ptr *QNetworkDiskCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::updateMetaData", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectUpdateMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::updateMetaData")
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaDataDefault(metaData QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaDataDefault(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkDiskCache_TimerEvent
func callbackQNetworkDiskCache_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::timerEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::timerEvent")
	}
}

func (ptr *QNetworkDiskCache) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkDiskCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkDiskCache_ChildEvent
func callbackQNetworkDiskCache_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::childEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::childEvent")
	}
}

func (ptr *QNetworkDiskCache) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkDiskCache_ConnectNotify
func callbackQNetworkDiskCache_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkDiskCache) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::connectNotify", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::connectNotify")
	}
}

func (ptr *QNetworkDiskCache) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkDiskCache) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkDiskCache_CustomEvent
func callbackQNetworkDiskCache_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::customEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::customEvent")
	}
}

func (ptr *QNetworkDiskCache) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkDiskCache) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkDiskCache_DeleteLater
func callbackQNetworkDiskCache_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkDiskCache) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::deleteLater", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::deleteLater")
	}
}

func (ptr *QNetworkDiskCache) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkDiskCache) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkDiskCache_DisconnectNotify
func callbackQNetworkDiskCache_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkDiskCache) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::disconnectNotify", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::disconnectNotify")
	}
}

func (ptr *QNetworkDiskCache) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkDiskCache) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkDiskCache_Event
func callbackQNetworkDiskCache_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkDiskCache) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::event", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::event")
	}
}

func (ptr *QNetworkDiskCache) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkDiskCache_EventFilter
func callbackQNetworkDiskCache_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkDiskCache) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::eventFilter", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::eventFilter")
	}
}

func (ptr *QNetworkDiskCache) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkDiskCache_MetaObject
func callbackQNetworkDiskCache_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkDiskCacheFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkDiskCache) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaObject", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaObject")
	}
}

func (ptr *QNetworkDiskCache) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkDiskCache_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkDiskCache) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkDiskCache_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkInterface struct {
	ptr unsafe.Pointer
}

type QNetworkInterface_ITF interface {
	QNetworkInterface_PTR() *QNetworkInterface
}

func (ptr *QNetworkInterface) QNetworkInterface_PTR() *QNetworkInterface {
	return ptr
}

func (ptr *QNetworkInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkInterface(ptr QNetworkInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkInterface_PTR().Pointer()
	}
	return nil
}

func NewQNetworkInterfaceFromPointer(ptr unsafe.Pointer) *QNetworkInterface {
	var n = new(QNetworkInterface)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkInterface__InterfaceFlag
//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int64

const (
	QNetworkInterface__IsUp           QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x20)
)

func NewQNetworkInterface() *QNetworkInterface {
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface())
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface2(PointerFromQNetworkInterface(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) AddressEntries() []*QNetworkAddressEntry {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkAddressEntry {
			var out = make([]*QNetworkAddressEntry, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkInterfaceFromPointer(l.data).addressEntries_atList(i)
			}
			return out
		}(C.QNetworkInterface_AddressEntries(ptr.Pointer()))
	}
	return nil
}

func QNetworkInterface_AllAddresses() []*QHostAddress {
	return func(l C.struct_QtNetwork_PackedList) []*QHostAddress {
		var out = make([]*QHostAddress, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkInterfaceFromPointer(l.data).allAddresses_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllAddresses())
}

func (ptr *QNetworkInterface) AllAddresses() []*QHostAddress {
	return func(l C.struct_QtNetwork_PackedList) []*QHostAddress {
		var out = make([]*QHostAddress, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkInterfaceFromPointer(l.data).allAddresses_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllAddresses())
}

func QNetworkInterface_AllInterfaces() []*QNetworkInterface {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkInterface {
		var out = make([]*QNetworkInterface, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkInterfaceFromPointer(l.data).allInterfaces_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllInterfaces())
}

func (ptr *QNetworkInterface) AllInterfaces() []*QNetworkInterface {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkInterface {
		var out = make([]*QNetworkInterface, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkInterfaceFromPointer(l.data).allInterfaces_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllInterfaces())
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {
	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceFlag(C.QNetworkInterface_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) HardwareAddress() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkInterface_HardwareAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) HumanReadableName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkInterface_HumanReadableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Index() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkInterface_Index(ptr.Pointer())))
	}
	return 0
}

func QNetworkInterface_InterfaceFromIndex(index int) *QNetworkInterface {
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromIndex(C.int(int32(index))))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) InterfaceFromIndex(index int) *QNetworkInterface {
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromIndex(C.int(int32(index))))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func QNetworkInterface_InterfaceFromName(name string) *QNetworkInterface {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) InterfaceFromName(name string) *QNetworkInterface {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func QNetworkInterface_InterfaceIndexFromName(name string) int {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return int(int32(C.QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(nameC)))
}

func (ptr *QNetworkInterface) InterfaceIndexFromName(name string) int {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return int(int32(C.QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(nameC)))
}

func QNetworkInterface_InterfaceNameFromIndex(index int) string {
	return cGoUnpackString(C.QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(C.int(int32(index))))
}

func (ptr *QNetworkInterface) InterfaceNameFromIndex(index int) string {
	return cGoUnpackString(C.QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(C.int(int32(index))))
}

func (ptr *QNetworkInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkInterface) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkInterface_Swap(ptr.Pointer(), PointerFromQNetworkInterface(other))
	}
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	if ptr.Pointer() != nil {
		C.QNetworkInterface_DestroyQNetworkInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkInterface) addressEntries_atList(i int) *QNetworkAddressEntry {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkAddressEntryFromPointer(C.QNetworkInterface_addressEntries_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkInterface) allAddresses_atList(i int) *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkInterface_allAddresses_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkInterface) allInterfaces_atList(i int) *QNetworkInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_allInterfaces_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

type QNetworkProxy struct {
	ptr unsafe.Pointer
}

type QNetworkProxy_ITF interface {
	QNetworkProxy_PTR() *QNetworkProxy
}

func (ptr *QNetworkProxy) QNetworkProxy_PTR() *QNetworkProxy {
	return ptr
}

func (ptr *QNetworkProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkProxy(ptr QNetworkProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxy_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFromPointer(ptr unsafe.Pointer) *QNetworkProxy {
	var n = new(QNetworkProxy)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkProxy__Capability
//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   QNetworkProxy__Capability = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        QNetworkProxy__Capability = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = QNetworkProxy__Capability(0x0010)
)

//go:generate stringer -type=QNetworkProxy__ProxyType
//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int64

const (
	QNetworkProxy__DefaultProxy     QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(5)
)

func NewQNetworkProxy() *QNetworkProxy {
	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func NewQNetworkProxy2(ty QNetworkProxy__ProxyType, hostName string, port uint16, user string, password string) *QNetworkProxy {
	var hostNameC = C.CString(hostName)
	defer C.free(unsafe.Pointer(hostNameC))
	var userC = C.CString(user)
	defer C.free(unsafe.Pointer(userC))
	var passwordC = C.CString(password)
	defer C.free(unsafe.Pointer(passwordC))
	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy2(C.longlong(ty), hostNameC, C.ushort(port), userC, passwordC))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func NewQNetworkProxy3(other QNetworkProxy_ITF) *QNetworkProxy {
	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy3(PointerFromQNetworkProxy(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func QNetworkProxy_ApplicationProxy() *QNetworkProxy {
	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_QNetworkProxy_ApplicationProxy())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func (ptr *QNetworkProxy) ApplicationProxy() *QNetworkProxy {
	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_QNetworkProxy_ApplicationProxy())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func (ptr *QNetworkProxy) Capabilities() QNetworkProxy__Capability {
	if ptr.Pointer() != nil {
		return QNetworkProxy__Capability(C.QNetworkProxy_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) HasRawHeader(headerName core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkProxy_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkProxy_Header(ptr.Pointer(), C.longlong(header)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxy) HostName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkProxy_HostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) IsCachingProxy() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsCachingProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsTransparentProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Password() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkProxy_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) Port() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QNetworkProxy_Port(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkProxy_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxy) RawHeaderList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkProxyFromPointer(l.data).rawHeaderList_atList(i)
			}
			return out
		}(C.QNetworkProxy_RawHeaderList(ptr.Pointer()))
	}
	return nil
}

func QNetworkProxy_SetApplicationProxy(networkProxy QNetworkProxy_ITF) {
	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(PointerFromQNetworkProxy(networkProxy))
}

func (ptr *QNetworkProxy) SetApplicationProxy(networkProxy QNetworkProxy_ITF) {
	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(PointerFromQNetworkProxy(networkProxy))
}

func (ptr *QNetworkProxy) SetCapabilities(capabilities QNetworkProxy__Capability) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetCapabilities(ptr.Pointer(), C.longlong(capabilities))
	}
}

func (ptr *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkProxy) SetHostName(hostName string) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QNetworkProxy_SetHostName(ptr.Pointer(), hostNameC)
	}
}

func (ptr *QNetworkProxy) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		C.QNetworkProxy_SetPassword(ptr.Pointer(), passwordC)
	}
}

func (ptr *QNetworkProxy) SetPort(port uint16) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QNetworkProxy) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QNetworkProxy) SetType(ty QNetworkProxy__ProxyType) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QNetworkProxy) SetUser(user string) {
	if ptr.Pointer() != nil {
		var userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
		C.QNetworkProxy_SetUser(ptr.Pointer(), userC)
	}
}

func (ptr *QNetworkProxy) Swap(other QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_Swap(ptr.Pointer(), PointerFromQNetworkProxy(other))
	}
}

func (ptr *QNetworkProxy) Type() QNetworkProxy__ProxyType {
	if ptr.Pointer() != nil {
		return QNetworkProxy__ProxyType(C.QNetworkProxy_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) User() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkProxy_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) DestroyQNetworkProxy() {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_DestroyQNetworkProxy(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkProxy) rawHeaderList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkProxy_rawHeaderList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

type QNetworkProxyFactory struct {
	ptr unsafe.Pointer
}

type QNetworkProxyFactory_ITF interface {
	QNetworkProxyFactory_PTR() *QNetworkProxyFactory
}

func (ptr *QNetworkProxyFactory) QNetworkProxyFactory_PTR() *QNetworkProxyFactory {
	return ptr
}

func (ptr *QNetworkProxyFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkProxyFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkProxyFactory(ptr QNetworkProxyFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyFactory_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) *QNetworkProxyFactory {
	var n = new(QNetworkProxyFactory)
	n.SetPointer(ptr)
	return n
}
func QNetworkProxyFactory_ProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		var out = make([]*QNetworkProxy, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkProxyFactoryFromPointer(l.data).proxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_ProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func (ptr *QNetworkProxyFactory) ProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		var out = make([]*QNetworkProxy, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkProxyFactoryFromPointer(l.data).proxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_ProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func (ptr *QNetworkProxyFactory) ConnectQueryProxy(f func(query *QNetworkProxyQuery) []*QNetworkProxy) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkProxyFactory::queryProxy", f)
	}
}

func (ptr *QNetworkProxyFactory) DisconnectQueryProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkProxyFactory::queryProxy")
	}
}

func (ptr *QNetworkProxyFactory) QueryProxy(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
			var out = make([]*QNetworkProxy, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkProxyFactoryFromPointer(l.data).queryProxy_atList(i)
			}
			return out
		}(C.QNetworkProxyFactory_QueryProxy(ptr.Pointer(), PointerFromQNetworkProxyQuery(query)))
	}
	return nil
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func (ptr *QNetworkProxyFactory) SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.char(int8(qt.GoBoolToInt(enable))))
}

func (ptr *QNetworkProxyFactory) SetUseSystemConfiguration(enable bool) {
	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.char(int8(qt.GoBoolToInt(enable))))
}

//export callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory
func callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkProxyFactory::~QNetworkProxyFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkProxyFactoryFromPointer(ptr).DestroyQNetworkProxyFactoryDefault()
	}
}

func (ptr *QNetworkProxyFactory) ConnectDestroyQNetworkProxyFactory(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkProxyFactory::~QNetworkProxyFactory", f)
	}
}

func (ptr *QNetworkProxyFactory) DisconnectDestroyQNetworkProxyFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkProxyFactory::~QNetworkProxyFactory")
	}
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactoryDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QNetworkProxyFactory_SystemProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		var out = make([]*QNetworkProxy, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkProxyFactoryFromPointer(l.data).systemProxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_SystemProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func (ptr *QNetworkProxyFactory) SystemProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		var out = make([]*QNetworkProxy, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQNetworkProxyFactoryFromPointer(l.data).systemProxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_SystemProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func (ptr *QNetworkProxyFactory) proxyForQuery_atList(i int) *QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxyFactory_proxyForQuery_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyFactory) queryProxy_atList(i int) *QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxyFactory_queryProxy_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyFactory) systemProxyForQuery_atList(i int) *QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxyFactory_systemProxyForQuery_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

type QNetworkProxyQuery struct {
	ptr unsafe.Pointer
}

type QNetworkProxyQuery_ITF interface {
	QNetworkProxyQuery_PTR() *QNetworkProxyQuery
}

func (ptr *QNetworkProxyQuery) QNetworkProxyQuery_PTR() *QNetworkProxyQuery {
	return ptr
}

func (ptr *QNetworkProxyQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkProxyQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkProxyQuery(ptr QNetworkProxyQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyQuery_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) *QNetworkProxyQuery {
	var n = new(QNetworkProxyQuery)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkProxyQuery__QueryType
//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__TcpServer  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(101)
)

func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery6(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var hostnameC = C.CString(hostname)
	defer C.free(unsafe.Pointer(hostnameC))
	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery6(PointerFromQNetworkConfiguration(networkConfiguration), hostnameC, C.int(int32(port)), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery5(networkConfiguration QNetworkConfiguration_ITF, requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery5(PointerFromQNetworkConfiguration(networkConfiguration), core.PointerFromQUrl(requestUrl), C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery7(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery7(PointerFromQNetworkConfiguration(networkConfiguration), C.ushort(bindPort), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery8(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery8(PointerFromQNetworkProxyQuery(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var hostnameC = C.CString(hostname)
	defer C.free(unsafe.Pointer(hostnameC))
	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery3(hostnameC, C.int(int32(port)), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery2(requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery2(core.PointerFromQUrl(requestUrl), C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery4(bindPort uint16, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery4(C.ushort(bindPort), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func (ptr *QNetworkProxyQuery) LocalPort() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkProxyQuery_LocalPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) NetworkConfiguration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkProxyQuery_NetworkConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyQuery) PeerHostName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkProxyQuery_PeerHostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) PeerPort() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkProxyQuery_PeerPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) ProtocolTag() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkProxyQuery_ProtocolTag(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) QueryType() QNetworkProxyQuery__QueryType {
	if ptr.Pointer() != nil {
		return QNetworkProxyQuery__QueryType(C.QNetworkProxyQuery_QueryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) SetLocalPort(port int) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetLocalPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetNetworkConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(networkConfiguration))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	if ptr.Pointer() != nil {
		var hostnameC = C.CString(hostname)
		defer C.free(unsafe.Pointer(hostnameC))
		C.QNetworkProxyQuery_SetPeerHostName(ptr.Pointer(), hostnameC)
	}
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	if ptr.Pointer() != nil {
		var protocolTagC = C.CString(protocolTag)
		defer C.free(unsafe.Pointer(protocolTagC))
		C.QNetworkProxyQuery_SetProtocolTag(ptr.Pointer(), protocolTagC)
	}
}

func (ptr *QNetworkProxyQuery) SetQueryType(ty QNetworkProxyQuery__QueryType) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetQueryType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QNetworkProxyQuery) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkProxyQuery) Swap(other QNetworkProxyQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_Swap(ptr.Pointer(), PointerFromQNetworkProxyQuery(other))
	}
}

func (ptr *QNetworkProxyQuery) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkProxyQuery_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_DestroyQNetworkProxyQuery(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkReply struct {
	core.QIODevice
}

type QNetworkReply_ITF interface {
	core.QIODevice_ITF
	QNetworkReply_PTR() *QNetworkReply
}

func (ptr *QNetworkReply) QNetworkReply_PTR() *QNetworkReply {
	return ptr
}

func (ptr *QNetworkReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkReply(ptr QNetworkReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkReply_PTR().Pointer()
	}
	return nil
}

func NewQNetworkReplyFromPointer(ptr unsafe.Pointer) *QNetworkReply {
	var n = new(QNetworkReply)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkReply__NetworkError
//QNetworkReply::NetworkError
type QNetworkReply__NetworkError int64

const (
	QNetworkReply__NoError                           QNetworkReply__NetworkError = QNetworkReply__NetworkError(0)
	QNetworkReply__ConnectionRefusedError            QNetworkReply__NetworkError = QNetworkReply__NetworkError(1)
	QNetworkReply__RemoteHostClosedError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(2)
	QNetworkReply__HostNotFoundError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(3)
	QNetworkReply__TimeoutError                      QNetworkReply__NetworkError = QNetworkReply__NetworkError(4)
	QNetworkReply__OperationCanceledError            QNetworkReply__NetworkError = QNetworkReply__NetworkError(5)
	QNetworkReply__SslHandshakeFailedError           QNetworkReply__NetworkError = QNetworkReply__NetworkError(6)
	QNetworkReply__TemporaryNetworkFailureError      QNetworkReply__NetworkError = QNetworkReply__NetworkError(7)
	QNetworkReply__NetworkSessionFailedError         QNetworkReply__NetworkError = QNetworkReply__NetworkError(8)
	QNetworkReply__BackgroundRequestNotAllowedError  QNetworkReply__NetworkError = QNetworkReply__NetworkError(9)
	QNetworkReply__TooManyRedirectsError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(10)
	QNetworkReply__InsecureRedirectError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(11)
	QNetworkReply__UnknownNetworkError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(99)
	QNetworkReply__ProxyConnectionRefusedError       QNetworkReply__NetworkError = QNetworkReply__NetworkError(101)
	QNetworkReply__ProxyConnectionClosedError        QNetworkReply__NetworkError = QNetworkReply__NetworkError(102)
	QNetworkReply__ProxyNotFoundError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(103)
	QNetworkReply__ProxyTimeoutError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(104)
	QNetworkReply__ProxyAuthenticationRequiredError  QNetworkReply__NetworkError = QNetworkReply__NetworkError(105)
	QNetworkReply__UnknownProxyError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(199)
	QNetworkReply__ContentAccessDenied               QNetworkReply__NetworkError = QNetworkReply__NetworkError(201)
	QNetworkReply__ContentOperationNotPermittedError QNetworkReply__NetworkError = QNetworkReply__NetworkError(202)
	QNetworkReply__ContentNotFoundError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(203)
	QNetworkReply__AuthenticationRequiredError       QNetworkReply__NetworkError = QNetworkReply__NetworkError(204)
	QNetworkReply__ContentReSendError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(205)
	QNetworkReply__ContentConflictError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(206)
	QNetworkReply__ContentGoneError                  QNetworkReply__NetworkError = QNetworkReply__NetworkError(207)
	QNetworkReply__UnknownContentError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(299)
	QNetworkReply__ProtocolUnknownError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(301)
	QNetworkReply__ProtocolInvalidOperationError     QNetworkReply__NetworkError = QNetworkReply__NetworkError(302)
	QNetworkReply__ProtocolFailure                   QNetworkReply__NetworkError = QNetworkReply__NetworkError(399)
	QNetworkReply__InternalServerError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(401)
	QNetworkReply__OperationNotImplementedError      QNetworkReply__NetworkError = QNetworkReply__NetworkError(402)
	QNetworkReply__ServiceUnavailableError           QNetworkReply__NetworkError = QNetworkReply__NetworkError(403)
	QNetworkReply__UnknownServerError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(499)
)

//export callbackQNetworkReply_SetSslConfigurationImplementation
func callbackQNetworkReply_SetSslConfigurationImplementation(ptr unsafe.Pointer, configuration unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::setSslConfigurationImplementation"); signal != nil {
		signal.(func(*QSslConfiguration))(NewQSslConfigurationFromPointer(configuration))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetSslConfigurationImplementationDefault(NewQSslConfigurationFromPointer(configuration))
	}
}

func (ptr *QNetworkReply) ConnectSetSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setSslConfigurationImplementation", f)
	}
}

func (ptr *QNetworkReply) DisconnectSetSslConfigurationImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setSslConfigurationImplementation")
	}
}

func (ptr *QNetworkReply) SetSslConfigurationImplementation(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfigurationImplementation(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QNetworkReply) SetSslConfigurationImplementationDefault(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfigurationImplementationDefault(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

//export callbackQNetworkReply_SslConfigurationImplementation
func callbackQNetworkReply_SslConfigurationImplementation(ptr unsafe.Pointer, configuration unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::sslConfigurationImplementation"); signal != nil {
		signal.(func(*QSslConfiguration))(NewQSslConfigurationFromPointer(configuration))
	} else {
		NewQNetworkReplyFromPointer(ptr).SslConfigurationImplementationDefault(NewQSslConfigurationFromPointer(configuration))
	}
}

func (ptr *QNetworkReply) ConnectSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::sslConfigurationImplementation", f)
	}
}

func (ptr *QNetworkReply) DisconnectSslConfigurationImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::sslConfigurationImplementation")
	}
}

func (ptr *QNetworkReply) SslConfigurationImplementation(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SslConfigurationImplementation(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QNetworkReply) SslConfigurationImplementationDefault(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SslConfigurationImplementationDefault(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func NewQNetworkReply(parent core.QObject_ITF) *QNetworkReply {
	var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkReply_NewQNetworkReply(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNetworkReply_Abort
func callbackQNetworkReply_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::abort"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::abort", f)
	}
}

func (ptr *QNetworkReply) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::abort")
	}
}

func (ptr *QNetworkReply) Abort() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Abort(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Attribute(code QNetworkRequest__Attribute) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkReply_Attribute(ptr.Pointer(), C.longlong(code)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_Close
func callbackQNetworkReply_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::close"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QNetworkReply) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::close", f)
	}
}

func (ptr *QNetworkReply) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::close")
	}
}

func (ptr *QNetworkReply) Close() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_CloseDefault(ptr.Pointer())
	}
}

//export callbackQNetworkReply_DownloadProgress
func callbackQNetworkReply_DownloadProgress(ptr unsafe.Pointer, bytesReceived C.longlong, bytesTotal C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::downloadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesReceived), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectDownloadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::downloadProgress", f)
	}
}

func (ptr *QNetworkReply) DisconnectDownloadProgress() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectDownloadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::downloadProgress")
	}
}

func (ptr *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DownloadProgress(ptr.Pointer(), C.longlong(bytesReceived), C.longlong(bytesTotal))
	}
}

//export callbackQNetworkReply_Encrypted
func callbackQNetworkReply_Encrypted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::encrypted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::encrypted", f)
	}
}

func (ptr *QNetworkReply) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::encrypted")
	}
}

func (ptr *QNetworkReply) Encrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Encrypted(ptr.Pointer())
	}
}

//export callbackQNetworkReply_Error2
func callbackQNetworkReply_Error2(ptr unsafe.Pointer, code C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::error2"); signal != nil {
		signal.(func(QNetworkReply__NetworkError))(QNetworkReply__NetworkError(code))
	}

}

func (ptr *QNetworkReply) ConnectError2(f func(code QNetworkReply__NetworkError)) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::error2", f)
	}
}

func (ptr *QNetworkReply) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::error2")
	}
}

func (ptr *QNetworkReply) Error2(code QNetworkReply__NetworkError) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Error2(ptr.Pointer(), C.longlong(code))
	}
}

func (ptr *QNetworkReply) Error() QNetworkReply__NetworkError {
	if ptr.Pointer() != nil {
		return QNetworkReply__NetworkError(C.QNetworkReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_Finished
func callbackQNetworkReply_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::finished", f)
	}
}

func (ptr *QNetworkReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::finished")
	}
}

func (ptr *QNetworkReply) Finished() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Finished(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) HasRawHeader(headerName core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkReply_Header(ptr.Pointer(), C.longlong(header)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_IgnoreSslErrors
func callbackQNetworkReply_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::ignoreSslErrors"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::ignoreSslErrors", f)
	}
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::ignoreSslErrors")
	}
}

func (ptr *QNetworkReply) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) IgnoreSslErrorsDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrorsDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) Manager() *QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkAccessManagerFromPointer(C.QNetworkReply_Manager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_MetaDataChanged
func callbackQNetworkReply_MetaDataChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaDataChanged", f)
	}
}

func (ptr *QNetworkReply) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaDataChanged")
	}
}

func (ptr *QNetworkReply) MetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Operation() QNetworkAccessManager__Operation {
	if ptr.Pointer() != nil {
		return QNetworkAccessManager__Operation(C.QNetworkReply_Operation(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_PreSharedKeyAuthenticationRequired
func callbackQNetworkReply_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkReply) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkReply) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkReply) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkReply_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) RawHeaderList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkReplyFromPointer(l.data).rawHeaderList_atList(i)
			}
			return out
		}(C.QNetworkReply_RawHeaderList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkReply) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_Redirected
func callbackQNetworkReply_Redirected(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::redirected"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QNetworkReply) ConnectRedirected(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectRedirected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::redirected", f)
	}
}

func (ptr *QNetworkReply) DisconnectRedirected() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectRedirected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::redirected")
	}
}

func (ptr *QNetworkReply) Redirected(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Redirected(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkReply) Request() *QNetworkRequest {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkRequestFromPointer(C.QNetworkReply_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkReply) SetError(errorCode QNetworkReply__NetworkError, errorString string) {
	if ptr.Pointer() != nil {
		var errorStringC = C.CString(errorString)
		defer C.free(unsafe.Pointer(errorStringC))
		C.QNetworkReply_SetError(ptr.Pointer(), C.longlong(errorCode), errorStringC)
	}
}

func (ptr *QNetworkReply) SetFinished(finished bool) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(finished))))
	}
}

func (ptr *QNetworkReply) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkReply) SetOperation(operation QNetworkAccessManager__Operation) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetOperation(ptr.Pointer(), C.longlong(operation))
	}
}

func (ptr *QNetworkReply) SetRawHeader(headerName core.QByteArray_ITF, value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(value))
	}
}

//export callbackQNetworkReply_SetReadBufferSize
func callbackQNetworkReply_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QNetworkReply) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setReadBufferSize", f)
	}
}

func (ptr *QNetworkReply) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setReadBufferSize")
	}
}

func (ptr *QNetworkReply) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkReply) SetReadBufferSizeDefault(size int64) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkReply) SetRequest(request QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetRequest(ptr.Pointer(), PointerFromQNetworkRequest(request))
	}
}

func (ptr *QNetworkReply) SetSslConfiguration(config QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(config))
	}
}

func (ptr *QNetworkReply) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkReply) SslConfiguration() *QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslConfigurationFromPointer(C.QNetworkReply_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_UploadProgress
func callbackQNetworkReply_UploadProgress(ptr unsafe.Pointer, bytesSent C.longlong, bytesTotal C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::uploadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesSent), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) ConnectUploadProgress(f func(bytesSent int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectUploadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::uploadProgress", f)
	}
}

func (ptr *QNetworkReply) DisconnectUploadProgress() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectUploadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::uploadProgress")
	}
}

func (ptr *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_UploadProgress(ptr.Pointer(), C.longlong(bytesSent), C.longlong(bytesTotal))
	}
}

func (ptr *QNetworkReply) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkReply_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) DestroyQNetworkReply() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DestroyQNetworkReply(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkReply) rawHeaderList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkReply_rawHeaderList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_AtEnd
func callbackQNetworkReply_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).AtEndDefault())))
}

func (ptr *QNetworkReply) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::atEnd", f)
	}
}

func (ptr *QNetworkReply) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::atEnd")
	}
}

func (ptr *QNetworkReply) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_BytesAvailable
func callbackQNetworkReply_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QNetworkReply) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesAvailable", f)
	}
}

func (ptr *QNetworkReply) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesAvailable")
	}
}

func (ptr *QNetworkReply) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_BytesToWrite
func callbackQNetworkReply_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QNetworkReply) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesToWrite", f)
	}
}

func (ptr *QNetworkReply) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesToWrite")
	}
}

func (ptr *QNetworkReply) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_CanReadLine
func callbackQNetworkReply_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QNetworkReply) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::canReadLine", f)
	}
}

func (ptr *QNetworkReply) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::canReadLine")
	}
}

func (ptr *QNetworkReply) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_IsSequential
func callbackQNetworkReply_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QNetworkReply) ConnectIsSequential(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::isSequential", f)
	}
}

func (ptr *QNetworkReply) DisconnectIsSequential() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::isSequential")
	}
}

func (ptr *QNetworkReply) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_Open
func callbackQNetworkReply_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QNetworkReply) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::open", f)
	}
}

func (ptr *QNetworkReply) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::open")
	}
}

func (ptr *QNetworkReply) Open(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QNetworkReply) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQNetworkReply_Pos
func callbackQNetworkReply_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).PosDefault())
}

func (ptr *QNetworkReply) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::pos", f)
	}
}

func (ptr *QNetworkReply) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::pos")
	}
}

func (ptr *QNetworkReply) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_ReadData
func callbackQNetworkReply_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}

	return C.longlong(0)
}

func (ptr *QNetworkReply) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::readData", f)
	}
}

func (ptr *QNetworkReply) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::readData")
	}
}

func (ptr *QNetworkReply) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QNetworkReply_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQNetworkReply_ReadLineData
func callbackQNetworkReply_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QNetworkReply) ConnectReadLineData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::readLineData", f)
	}
}

func (ptr *QNetworkReply) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::readLineData")
	}
}

func (ptr *QNetworkReply) ReadLineData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QNetworkReply) ReadLineDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQNetworkReply_Reset
func callbackQNetworkReply_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).ResetDefault())))
}

func (ptr *QNetworkReply) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::reset", f)
	}
}

func (ptr *QNetworkReply) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::reset")
	}
}

func (ptr *QNetworkReply) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_Seek
func callbackQNetworkReply_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QNetworkReply) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::seek", f)
	}
}

func (ptr *QNetworkReply) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::seek")
	}
}

func (ptr *QNetworkReply) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QNetworkReply) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQNetworkReply_Size
func callbackQNetworkReply_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).SizeDefault())
}

func (ptr *QNetworkReply) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::size", f)
	}
}

func (ptr *QNetworkReply) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::size")
	}
}

func (ptr *QNetworkReply) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_WaitForBytesWritten
func callbackQNetworkReply_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QNetworkReply) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForBytesWritten", f)
	}
}

func (ptr *QNetworkReply) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForBytesWritten")
	}
}

func (ptr *QNetworkReply) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QNetworkReply) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQNetworkReply_WaitForReadyRead
func callbackQNetworkReply_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QNetworkReply) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForReadyRead", f)
	}
}

func (ptr *QNetworkReply) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForReadyRead")
	}
}

func (ptr *QNetworkReply) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QNetworkReply) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQNetworkReply_WriteData
func callbackQNetworkReply_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(0)
}

func (ptr *QNetworkReply) ConnectWriteData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::writeData", f)
	}
}

func (ptr *QNetworkReply) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::writeData")
	}
}

func (ptr *QNetworkReply) WriteData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQNetworkReply_TimerEvent
func callbackQNetworkReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::timerEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::timerEvent")
	}
}

func (ptr *QNetworkReply) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkReply_ChildEvent
func callbackQNetworkReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::childEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::childEvent")
	}
}

func (ptr *QNetworkReply) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkReply_ConnectNotify
func callbackQNetworkReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkReply) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::connectNotify", f)
	}
}

func (ptr *QNetworkReply) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::connectNotify")
	}
}

func (ptr *QNetworkReply) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkReply_CustomEvent
func callbackQNetworkReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::customEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::customEvent")
	}
}

func (ptr *QNetworkReply) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkReply_DeleteLater
func callbackQNetworkReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkReply) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::deleteLater", f)
	}
}

func (ptr *QNetworkReply) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::deleteLater")
	}
}

func (ptr *QNetworkReply) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkReply_DisconnectNotify
func callbackQNetworkReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkReply) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::disconnectNotify", f)
	}
}

func (ptr *QNetworkReply) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::disconnectNotify")
	}
}

func (ptr *QNetworkReply) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkReply_Event
func callbackQNetworkReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkReply) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::event", f)
	}
}

func (ptr *QNetworkReply) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::event")
	}
}

func (ptr *QNetworkReply) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkReply_EventFilter
func callbackQNetworkReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkReply) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::eventFilter", f)
	}
}

func (ptr *QNetworkReply) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::eventFilter")
	}
}

func (ptr *QNetworkReply) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkReply_MetaObject
func callbackQNetworkReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkReply) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaObject", f)
	}
}

func (ptr *QNetworkReply) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaObject")
	}
}

func (ptr *QNetworkReply) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkReply_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkReply) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkRequest struct {
	ptr unsafe.Pointer
}

type QNetworkRequest_ITF interface {
	QNetworkRequest_PTR() *QNetworkRequest
}

func (ptr *QNetworkRequest) QNetworkRequest_PTR() *QNetworkRequest {
	return ptr
}

func (ptr *QNetworkRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkRequest(ptr QNetworkRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkRequest_PTR().Pointer()
	}
	return nil
}

func NewQNetworkRequestFromPointer(ptr unsafe.Pointer) *QNetworkRequest {
	var n = new(QNetworkRequest)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkRequest__Attribute
//QNetworkRequest::Attribute
type QNetworkRequest__Attribute int64

const (
	QNetworkRequest__HttpStatusCodeAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(0)
	QNetworkRequest__HttpReasonPhraseAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(1)
	QNetworkRequest__RedirectionTargetAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(2)
	QNetworkRequest__ConnectionEncryptedAttribute          QNetworkRequest__Attribute = QNetworkRequest__Attribute(3)
	QNetworkRequest__CacheLoadControlAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(4)
	QNetworkRequest__CacheSaveControlAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(5)
	QNetworkRequest__SourceIsFromCacheAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(6)
	QNetworkRequest__DoNotBufferUploadDataAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(7)
	QNetworkRequest__HttpPipeliningAllowedAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(8)
	QNetworkRequest__HttpPipeliningWasUsedAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(9)
	QNetworkRequest__CustomVerbAttribute                   QNetworkRequest__Attribute = QNetworkRequest__Attribute(10)
	QNetworkRequest__CookieLoadControlAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(11)
	QNetworkRequest__AuthenticationReuseAttribute          QNetworkRequest__Attribute = QNetworkRequest__Attribute(12)
	QNetworkRequest__CookieSaveControlAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(13)
	QNetworkRequest__MaximumDownloadBufferSizeAttribute    QNetworkRequest__Attribute = QNetworkRequest__Attribute(14)
	QNetworkRequest__DownloadBufferAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(15)
	QNetworkRequest__SynchronousRequestAttribute           QNetworkRequest__Attribute = QNetworkRequest__Attribute(16)
	QNetworkRequest__BackgroundRequestAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(17)
	QNetworkRequest__SpdyAllowedAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(18)
	QNetworkRequest__SpdyWasUsedAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(19)
	QNetworkRequest__EmitAllUploadProgressSignalsAttribute QNetworkRequest__Attribute = QNetworkRequest__Attribute(20)
	QNetworkRequest__FollowRedirectsAttribute              QNetworkRequest__Attribute = QNetworkRequest__Attribute(21)
	QNetworkRequest__User                                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(1000)
	QNetworkRequest__UserMax                               QNetworkRequest__Attribute = QNetworkRequest__Attribute(32767)
)

//go:generate stringer -type=QNetworkRequest__CacheLoadControl
//QNetworkRequest::CacheLoadControl
type QNetworkRequest__CacheLoadControl int64

const (
	QNetworkRequest__AlwaysNetwork QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(0)
	QNetworkRequest__PreferNetwork QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(1)
	QNetworkRequest__PreferCache   QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(2)
	QNetworkRequest__AlwaysCache   QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(3)
)

//go:generate stringer -type=QNetworkRequest__KnownHeaders
//QNetworkRequest::KnownHeaders
type QNetworkRequest__KnownHeaders int64

const (
	QNetworkRequest__ContentTypeHeader        QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(0)
	QNetworkRequest__ContentLengthHeader      QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(1)
	QNetworkRequest__LocationHeader           QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(2)
	QNetworkRequest__LastModifiedHeader       QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(3)
	QNetworkRequest__CookieHeader             QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(4)
	QNetworkRequest__SetCookieHeader          QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(5)
	QNetworkRequest__ContentDispositionHeader QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(6)
	QNetworkRequest__UserAgentHeader          QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(7)
	QNetworkRequest__ServerHeader             QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(8)
)

//go:generate stringer -type=QNetworkRequest__LoadControl
//QNetworkRequest::LoadControl
type QNetworkRequest__LoadControl int64

const (
	QNetworkRequest__Automatic QNetworkRequest__LoadControl = QNetworkRequest__LoadControl(0)
	QNetworkRequest__Manual    QNetworkRequest__LoadControl = QNetworkRequest__LoadControl(1)
)

//go:generate stringer -type=QNetworkRequest__Priority
//QNetworkRequest::Priority
type QNetworkRequest__Priority int64

const (
	QNetworkRequest__HighPriority   QNetworkRequest__Priority = QNetworkRequest__Priority(1)
	QNetworkRequest__NormalPriority QNetworkRequest__Priority = QNetworkRequest__Priority(3)
	QNetworkRequest__LowPriority    QNetworkRequest__Priority = QNetworkRequest__Priority(5)
)

func NewQNetworkRequest2(other QNetworkRequest_ITF) *QNetworkRequest {
	var tmpValue = NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest2(PointerFromQNetworkRequest(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
	return tmpValue
}

func NewQNetworkRequest(url core.QUrl_ITF) *QNetworkRequest {
	var tmpValue = NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest(core.PointerFromQUrl(url)))
	runtime.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
	return tmpValue
}

func (ptr *QNetworkRequest) Attribute(code QNetworkRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkRequest_Attribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) HasRawHeader(headerName core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkRequest_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkRequest) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkRequest_Header(ptr.Pointer(), C.longlong(header)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) MaximumRedirectsAllowed() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkRequest_MaximumRedirectsAllowed(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkRequest) OriginatingObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QNetworkRequest_OriginatingObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) Priority() QNetworkRequest__Priority {
	if ptr.Pointer() != nil {
		return QNetworkRequest__Priority(C.QNetworkRequest_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkRequest) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkRequest_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) RawHeaderList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQNetworkRequestFromPointer(l.data).rawHeaderList_atList(i)
			}
			return out
		}(C.QNetworkRequest_RawHeaderList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkRequest) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkRequest) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkRequest) SetMaximumRedirectsAllowed(maxRedirectsAllowed int) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetMaximumRedirectsAllowed(ptr.Pointer(), C.int(int32(maxRedirectsAllowed)))
	}
}

func (ptr *QNetworkRequest) SetOriginatingObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetOriginatingObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QNetworkRequest) SetPriority(priority QNetworkRequest__Priority) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetPriority(ptr.Pointer(), C.longlong(priority))
	}
}

func (ptr *QNetworkRequest) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QNetworkRequest) SetSslConfiguration(config QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(config))
	}
}

func (ptr *QNetworkRequest) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkRequest) SslConfiguration() *QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslConfigurationFromPointer(C.QNetworkRequest_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) Swap(other QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_Swap(ptr.Pointer(), PointerFromQNetworkRequest(other))
	}
}

func (ptr *QNetworkRequest) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkRequest_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) DestroyQNetworkRequest() {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_DestroyQNetworkRequest(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkRequest) rawHeaderList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QNetworkRequest_rawHeaderList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

type QNetworkSession struct {
	core.QObject
}

type QNetworkSession_ITF interface {
	core.QObject_ITF
	QNetworkSession_PTR() *QNetworkSession
}

func (ptr *QNetworkSession) QNetworkSession_PTR() *QNetworkSession {
	return ptr
}

func (ptr *QNetworkSession) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkSession) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkSession(ptr QNetworkSession_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkSession_PTR().Pointer()
	}
	return nil
}

func NewQNetworkSessionFromPointer(ptr unsafe.Pointer) *QNetworkSession {
	var n = new(QNetworkSession)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNetworkSession__SessionError
//QNetworkSession::SessionError
type QNetworkSession__SessionError int64

const (
	QNetworkSession__UnknownSessionError        QNetworkSession__SessionError = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        QNetworkSession__SessionError = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               QNetworkSession__SessionError = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  QNetworkSession__SessionError = QNetworkSession__SessionError(4)
)

//go:generate stringer -type=QNetworkSession__State
//QNetworkSession::State
type QNetworkSession__State int64

const (
	QNetworkSession__Invalid      QNetworkSession__State = QNetworkSession__State(0)
	QNetworkSession__NotAvailable QNetworkSession__State = QNetworkSession__State(1)
	QNetworkSession__Connecting   QNetworkSession__State = QNetworkSession__State(2)
	QNetworkSession__Connected    QNetworkSession__State = QNetworkSession__State(3)
	QNetworkSession__Closing      QNetworkSession__State = QNetworkSession__State(4)
	QNetworkSession__Disconnected QNetworkSession__State = QNetworkSession__State(5)
	QNetworkSession__Roaming      QNetworkSession__State = QNetworkSession__State(6)
)

//go:generate stringer -type=QNetworkSession__UsagePolicy
//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(1)
)

func NewQNetworkSession(connectionConfig QNetworkConfiguration_ITF, parent core.QObject_ITF) *QNetworkSession {
	var tmpValue = NewQNetworkSessionFromPointer(C.QNetworkSession_NewQNetworkSession(PointerFromQNetworkConfiguration(connectionConfig), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNetworkSession_Accept
func callbackQNetworkSession_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::accept"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QNetworkSession) ConnectAccept(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::accept", f)
	}
}

func (ptr *QNetworkSession) DisconnectAccept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::accept")
	}
}

func (ptr *QNetworkSession) Accept() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Accept(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_AcceptDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ActiveTime() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QNetworkSession_ActiveTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) BytesReceived() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QNetworkSession_BytesReceived(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) BytesWritten() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QNetworkSession_BytesWritten(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_Close
func callbackQNetworkSession_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::close"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QNetworkSession) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::close", f)
	}
}

func (ptr *QNetworkSession) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::close")
	}
}

func (ptr *QNetworkSession) Close() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_CloseDefault(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Closed
func callbackQNetworkSession_Closed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::closed", f)
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::closed")
	}
}

func (ptr *QNetworkSession) Closed() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Closed(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Configuration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkSession_Configuration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkSession_Error2
func callbackQNetworkSession_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::error2"); signal != nil {
		signal.(func(QNetworkSession__SessionError))(QNetworkSession__SessionError(error))
	}

}

func (ptr *QNetworkSession) ConnectError2(f func(error QNetworkSession__SessionError)) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::error2", f)
	}
}

func (ptr *QNetworkSession) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::error2")
	}
}

func (ptr *QNetworkSession) Error2(error QNetworkSession__SessionError) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Error2(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkSession_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQNetworkSession_Ignore
func callbackQNetworkSession_Ignore(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::ignore"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).IgnoreDefault()
	}
}

func (ptr *QNetworkSession) ConnectIgnore(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::ignore", f)
	}
}

func (ptr *QNetworkSession) DisconnectIgnore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::ignore")
	}
}

func (ptr *QNetworkSession) Ignore() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Ignore(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) IgnoreDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_IgnoreDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Interface() *QNetworkInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkSession_Interface(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) IsOpen() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkSession_Migrate
func callbackQNetworkSession_Migrate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::migrate"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).MigrateDefault()
	}
}

func (ptr *QNetworkSession) ConnectMigrate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::migrate", f)
	}
}

func (ptr *QNetworkSession) DisconnectMigrate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::migrate")
	}
}

func (ptr *QNetworkSession) Migrate() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Migrate(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) MigrateDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_MigrateDefault(ptr.Pointer())
	}
}

//export callbackQNetworkSession_NewConfigurationActivated
func callbackQNetworkSession_NewConfigurationActivated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::newConfigurationActivated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNewConfigurationActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::newConfigurationActivated", f)
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::newConfigurationActivated")
	}
}

func (ptr *QNetworkSession) NewConfigurationActivated() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_NewConfigurationActivated(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Open
func callbackQNetworkSession_Open(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::open"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).OpenDefault()
	}
}

func (ptr *QNetworkSession) ConnectOpen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::open", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::open")
	}
}

func (ptr *QNetworkSession) Open() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Open(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) OpenDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_OpenDefault(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Opened
func callbackQNetworkSession_Opened(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::opened"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectOpened(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::opened", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::opened")
	}
}

func (ptr *QNetworkSession) Opened() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Opened(ptr.Pointer())
	}
}

//export callbackQNetworkSession_PreferredConfigurationChanged
func callbackQNetworkSession_PreferredConfigurationChanged(ptr unsafe.Pointer, config unsafe.Pointer, isSeamless C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::preferredConfigurationChanged"); signal != nil {
		signal.(func(*QNetworkConfiguration, bool))(NewQNetworkConfigurationFromPointer(config), int8(isSeamless) != 0)
	}

}

func (ptr *QNetworkSession) ConnectPreferredConfigurationChanged(f func(config *QNetworkConfiguration, isSeamless bool)) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectPreferredConfigurationChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::preferredConfigurationChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectPreferredConfigurationChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectPreferredConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::preferredConfigurationChanged")
	}
}

func (ptr *QNetworkSession) PreferredConfigurationChanged(config QNetworkConfiguration_ITF, isSeamless bool) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_PreferredConfigurationChanged(ptr.Pointer(), PointerFromQNetworkConfiguration(config), C.char(int8(qt.GoBoolToInt(isSeamless))))
	}
}

//export callbackQNetworkSession_Reject
func callbackQNetworkSession_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::reject"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QNetworkSession) ConnectReject(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::reject", f)
	}
}

func (ptr *QNetworkSession) DisconnectReject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::reject")
	}
}

func (ptr *QNetworkSession) Reject() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Reject(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) RejectDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_RejectDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) SessionProperty(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkSession_SessionProperty(ptr.Pointer(), keyC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QNetworkSession_SetSessionProperty(ptr.Pointer(), keyC, core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkSession) State() QNetworkSession__State {
	if ptr.Pointer() != nil {
		return QNetworkSession__State(C.QNetworkSession_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_StateChanged
func callbackQNetworkSession_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::stateChanged"); signal != nil {
		signal.(func(QNetworkSession__State))(QNetworkSession__State(state))
	}

}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stateChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stateChanged")
	}
}

func (ptr *QNetworkSession) StateChanged(state QNetworkSession__State) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQNetworkSession_Stop
func callbackQNetworkSession_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).StopDefault()
	}
}

func (ptr *QNetworkSession) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stop", f)
	}
}

func (ptr *QNetworkSession) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stop")
	}
}

func (ptr *QNetworkSession) Stop() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Stop(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) StopDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_StopDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	if ptr.Pointer() != nil {
		return QNetworkSession__UsagePolicy(C.QNetworkSession_UsagePolicies(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_UsagePoliciesChanged
func callbackQNetworkSession_UsagePoliciesChanged(ptr unsafe.Pointer, usagePolicies C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::usagePoliciesChanged"); signal != nil {
		signal.(func(QNetworkSession__UsagePolicy))(QNetworkSession__UsagePolicy(usagePolicies))
	}

}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectUsagePoliciesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::usagePoliciesChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::usagePoliciesChanged")
	}
}

func (ptr *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_UsagePoliciesChanged(ptr.Pointer(), C.longlong(usagePolicies))
	}
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_WaitForOpened(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQNetworkSession_DestroyQNetworkSession
func callbackQNetworkSession_DestroyQNetworkSession(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::~QNetworkSession"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).DestroyQNetworkSessionDefault()
	}
}

func (ptr *QNetworkSession) ConnectDestroyQNetworkSession(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::~QNetworkSession", f)
	}
}

func (ptr *QNetworkSession) DisconnectDestroyQNetworkSession() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::~QNetworkSession")
	}
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSession(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) DestroyQNetworkSessionDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSessionDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkSession_TimerEvent
func callbackQNetworkSession_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::timerEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::timerEvent")
	}
}

func (ptr *QNetworkSession) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkSession) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkSession_ChildEvent
func callbackQNetworkSession_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::childEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::childEvent")
	}
}

func (ptr *QNetworkSession) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkSession) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkSession_ConnectNotify
func callbackQNetworkSession_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkSessionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkSession) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::connectNotify", f)
	}
}

func (ptr *QNetworkSession) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::connectNotify")
	}
}

func (ptr *QNetworkSession) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkSession) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkSession_CustomEvent
func callbackQNetworkSession_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::customEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::customEvent")
	}
}

func (ptr *QNetworkSession) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkSession) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkSession_DeleteLater
func callbackQNetworkSession_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkSession) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::deleteLater", f)
	}
}

func (ptr *QNetworkSession) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::deleteLater")
	}
}

func (ptr *QNetworkSession) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkSession_DisconnectNotify
func callbackQNetworkSession_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkSessionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkSession) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::disconnectNotify", f)
	}
}

func (ptr *QNetworkSession) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::disconnectNotify")
	}
}

func (ptr *QNetworkSession) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkSession) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkSession_Event
func callbackQNetworkSession_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkSessionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkSession) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::event", f)
	}
}

func (ptr *QNetworkSession) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::event")
	}
}

func (ptr *QNetworkSession) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkSession) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkSession_EventFilter
func callbackQNetworkSession_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkSessionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkSession) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::eventFilter", f)
	}
}

func (ptr *QNetworkSession) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::eventFilter")
	}
}

func (ptr *QNetworkSession) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkSession) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkSession_MetaObject
func callbackQNetworkSession_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkSessionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkSession) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::metaObject", f)
	}
}

func (ptr *QNetworkSession) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::metaObject")
	}
}

func (ptr *QNetworkSession) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkSession_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkSession) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkSession_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSsl struct {
	ptr unsafe.Pointer
}

type QSsl_ITF interface {
	QSsl_PTR() *QSsl
}

func (ptr *QSsl) QSsl_PTR() *QSsl {
	return ptr
}

func (ptr *QSsl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSsl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSsl(ptr QSsl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSsl_PTR().Pointer()
	}
	return nil
}

func NewQSslFromPointer(ptr unsafe.Pointer) *QSsl {
	var n = new(QSsl)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSsl) DestroyQSsl() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QSsl__AlternativeNameEntryType
//QSsl::AlternativeNameEntryType
type QSsl__AlternativeNameEntryType int64

const (
	QSsl__EmailEntry QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(0)
	QSsl__DnsEntry   QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(1)
)

//go:generate stringer -type=QSsl__EncodingFormat
//QSsl::EncodingFormat
type QSsl__EncodingFormat int64

const (
	QSsl__Pem QSsl__EncodingFormat = QSsl__EncodingFormat(0)
	QSsl__Der QSsl__EncodingFormat = QSsl__EncodingFormat(1)
)

//go:generate stringer -type=QSsl__KeyAlgorithm
//QSsl::KeyAlgorithm
type QSsl__KeyAlgorithm int64

const (
	QSsl__Opaque QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(0)
	QSsl__Rsa    QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(1)
	QSsl__Dsa    QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(2)
	QSsl__Ec     QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(3)
)

//go:generate stringer -type=QSsl__KeyType
//QSsl::KeyType
type QSsl__KeyType int64

const (
	QSsl__PrivateKey QSsl__KeyType = QSsl__KeyType(0)
	QSsl__PublicKey  QSsl__KeyType = QSsl__KeyType(1)
)

//go:generate stringer -type=QSsl__SslOption
//QSsl::SslOption
type QSsl__SslOption int64

const (
	QSsl__SslOptionDisableEmptyFragments         QSsl__SslOption = QSsl__SslOption(0x01)
	QSsl__SslOptionDisableSessionTickets         QSsl__SslOption = QSsl__SslOption(0x02)
	QSsl__SslOptionDisableCompression            QSsl__SslOption = QSsl__SslOption(0x04)
	QSsl__SslOptionDisableServerNameIndication   QSsl__SslOption = QSsl__SslOption(0x08)
	QSsl__SslOptionDisableLegacyRenegotiation    QSsl__SslOption = QSsl__SslOption(0x10)
	QSsl__SslOptionDisableSessionSharing         QSsl__SslOption = QSsl__SslOption(0x20)
	QSsl__SslOptionDisableSessionPersistence     QSsl__SslOption = QSsl__SslOption(0x40)
	QSsl__SslOptionDisableServerCipherPreference QSsl__SslOption = QSsl__SslOption(0x80)
)

//go:generate stringer -type=QSsl__SslProtocol
//QSsl::SslProtocol
type QSsl__SslProtocol int64

var (
	QSsl__SslV3           QSsl__SslProtocol = QSsl__SslProtocol(0)
	QSsl__SslV2           QSsl__SslProtocol = QSsl__SslProtocol(1)
	QSsl__TlsV1_0         QSsl__SslProtocol = QSsl__SslProtocol(2)
	QSsl__TlsV1           QSsl__SslProtocol = QSsl__SslProtocol(QSsl__TlsV1_0)
	QSsl__TlsV1_1         QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_TlsV1_1_Type())
	QSsl__TlsV1_2         QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_TlsV1_2_Type())
	QSsl__AnyProtocol     QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_AnyProtocol_Type())
	QSsl__TlsV1SslV3      QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_TlsV1SslV3_Type())
	QSsl__SecureProtocols QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_SecureProtocols_Type())
	QSsl__TlsV1_0OrLater  QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_TlsV1_0OrLater_Type())
	QSsl__TlsV1_1OrLater  QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_TlsV1_1OrLater_Type())
	QSsl__TlsV1_2OrLater  QSsl__SslProtocol = QSsl__SslProtocol(C.QSsl_TlsV1_2OrLater_Type())
	QSsl__UnknownProtocol QSsl__SslProtocol = QSsl__SslProtocol(-1)
)

type QSslCertificate struct {
	ptr unsafe.Pointer
}

type QSslCertificate_ITF interface {
	QSslCertificate_PTR() *QSslCertificate
}

func (ptr *QSslCertificate) QSslCertificate_PTR() *QSslCertificate {
	return ptr
}

func (ptr *QSslCertificate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslCertificate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslCertificate(ptr QSslCertificate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificate_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateFromPointer(ptr unsafe.Pointer) *QSslCertificate {
	var n = new(QSslCertificate)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QSslCertificate__SubjectInfo
//QSslCertificate::SubjectInfo
type QSslCertificate__SubjectInfo int64

const (
	QSslCertificate__Organization               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(0)
	QSslCertificate__CommonName                 QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(1)
	QSslCertificate__LocalityName               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(2)
	QSslCertificate__OrganizationalUnitName     QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(3)
	QSslCertificate__CountryName                QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(4)
	QSslCertificate__StateOrProvinceName        QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(5)
	QSslCertificate__DistinguishedNameQualifier QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(6)
	QSslCertificate__SerialNumber               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(7)
	QSslCertificate__EmailAddress               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(8)
)

func NewQSslCertificate(device core.QIODevice_ITF, format QSsl__EncodingFormat) *QSslCertificate {
	var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate(core.PointerFromQIODevice(device), C.longlong(format)))
	runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func NewQSslCertificate2(data core.QByteArray_ITF, format QSsl__EncodingFormat) *QSslCertificate {
	var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate2(core.PointerFromQByteArray(data), C.longlong(format)))
	runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func NewQSslCertificate3(other QSslCertificate_ITF) *QSslCertificate {
	var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate3(PointerFromQSslCertificate(other)))
	runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func (ptr *QSslCertificate) Clear() {
	if ptr.Pointer() != nil {
		C.QSslCertificate_Clear(ptr.Pointer())
	}
}

func (ptr *QSslCertificate) Digest(algorithm core.QCryptographicHash__Algorithm) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_Digest(ptr.Pointer(), C.longlong(algorithm)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func QSslCertificate_FromData(data core.QByteArray_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslCertificateFromPointer(l.data).fromData_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromData(core.PointerFromQByteArray(data), C.longlong(format)))
}

func (ptr *QSslCertificate) FromData(data core.QByteArray_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslCertificateFromPointer(l.data).fromData_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromData(core.PointerFromQByteArray(data), C.longlong(format)))
}

func QSslCertificate_FromDevice(device core.QIODevice_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslCertificateFromPointer(l.data).fromDevice_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromDevice(core.PointerFromQIODevice(device), C.longlong(format)))
}

func (ptr *QSslCertificate) FromDevice(device core.QIODevice_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslCertificateFromPointer(l.data).fromDevice_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromDevice(core.PointerFromQIODevice(device), C.longlong(format)))
}

func QSslCertificate_FromPath(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) []*QSslCertificate {
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslCertificateFromPointer(l.data).fromPath_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromPath(pathC, C.longlong(format), C.longlong(syntax)))
}

func (ptr *QSslCertificate) FromPath(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) []*QSslCertificate {
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslCertificateFromPointer(l.data).fromPath_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromPath(pathC, C.longlong(format), C.longlong(syntax)))
}

func (ptr *QSslCertificate) IsBlacklisted() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsBlacklisted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) Swap(other QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate_Swap(ptr.Pointer(), PointerFromQSslCertificate(other))
	}
}

func (ptr *QSslCertificate) DestroyQSslCertificate() {
	if ptr.Pointer() != nil {
		C.QSslCertificate_DestroyQSslCertificate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslCertificate) EffectiveDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QSslCertificate_EffectiveDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) ExpiryDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QSslCertificate_ExpiryDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) Extensions() []*QSslCertificateExtension {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificateExtension {
			var out = make([]*QSslCertificateExtension, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslCertificateFromPointer(l.data).extensions_atList(i)
			}
			return out
		}(C.QSslCertificate_Extensions(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) IsSelfSigned() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsSelfSigned(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) IssuerInfo(subject QSslCertificate__SubjectInfo) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSslCertificate_IssuerInfo(ptr.Pointer(), C.longlong(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfo2(attribute core.QByteArray_ITF) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSslCertificate_IssuerInfo2(ptr.Pointer(), core.PointerFromQByteArray(attribute))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfoAttributes() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslCertificateFromPointer(l.data).issuerInfoAttributes_atList(i)
			}
			return out
		}(C.QSslCertificate_IssuerInfoAttributes(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) PublicKey() *QSslKey {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslCertificate_PublicKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) SerialNumber() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_SerialNumber(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) SubjectInfo(subject QSslCertificate__SubjectInfo) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSslCertificate_SubjectInfo(ptr.Pointer(), C.longlong(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfo2(attribute core.QByteArray_ITF) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSslCertificate_SubjectInfo2(ptr.Pointer(), core.PointerFromQByteArray(attribute))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfoAttributes() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslCertificateFromPointer(l.data).subjectInfoAttributes_atList(i)
			}
			return out
		}(C.QSslCertificate_SubjectInfoAttributes(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) ToDer() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_ToDer(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) ToPem() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_ToPem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) ToText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCertificate_ToText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificate) Version() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_Version(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) fromData_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_fromData_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) fromDevice_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_fromDevice_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) fromPath_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_fromPath_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) extensions_atList(i int) *QSslCertificateExtension {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateExtensionFromPointer(C.QSslCertificate_extensions_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) issuerInfoAttributes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_issuerInfoAttributes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) subjectInfoAttributes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslCertificate_subjectInfoAttributes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

type QSslCertificateExtension struct {
	ptr unsafe.Pointer
}

type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (ptr *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return ptr
}

func (ptr *QSslCertificateExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslCertificateExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtension_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return n
}
func NewQSslCertificateExtension() *QSslCertificateExtension {
	var tmpValue = NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension())
	runtime.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
	return tmpValue
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	var tmpValue = NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(PointerFromQSslCertificateExtension(other)))
	runtime.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
	return tmpValue
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsCritical(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCertificateExtension_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Oid() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCertificateExtension_Oid(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_Swap(ptr.Pointer(), PointerFromQSslCertificateExtension(other))
	}
}

func (ptr *QSslCertificateExtension) Value() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSslCertificateExtension_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslCipher struct {
	ptr unsafe.Pointer
}

type QSslCipher_ITF interface {
	QSslCipher_PTR() *QSslCipher
}

func (ptr *QSslCipher) QSslCipher_PTR() *QSslCipher {
	return ptr
}

func (ptr *QSslCipher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslCipher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslCipher(ptr QSslCipher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCipher_PTR().Pointer()
	}
	return nil
}

func NewQSslCipherFromPointer(ptr unsafe.Pointer) *QSslCipher {
	var n = new(QSslCipher)
	n.SetPointer(ptr)
	return n
}
func NewQSslCipher() *QSslCipher {
	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher())
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher4(other QSslCipher_ITF) *QSslCipher {
	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher4(PointerFromQSslCipher(other)))
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher2(name string) *QSslCipher {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher2(nameC))
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher3(name string, protocol QSsl__SslProtocol) *QSslCipher {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher3(nameC, C.longlong(protocol)))
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func (ptr *QSslCipher) AuthenticationMethod() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCipher_AuthenticationMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) EncryptionMethod() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCipher_EncryptionMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslCipher_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCipher) KeyExchangeMethod() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCipher_KeyExchangeMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCipher_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) Protocol() QSsl__SslProtocol {
	if ptr.Pointer() != nil {
		return QSsl__SslProtocol(C.QSslCipher_Protocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslCipher) ProtocolString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCipher_ProtocolString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) SupportedBits() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslCipher_SupportedBits(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslCipher) Swap(other QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCipher_Swap(ptr.Pointer(), PointerFromQSslCipher(other))
	}
}

func (ptr *QSslCipher) UsedBits() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslCipher_UsedBits(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslCipher) DestroyQSslCipher() {
	if ptr.Pointer() != nil {
		C.QSslCipher_DestroyQSslCipher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslConfiguration struct {
	ptr unsafe.Pointer
}

type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (ptr *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration {
	return ptr
}

func (ptr *QSslConfiguration) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslConfiguration) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslConfiguration(ptr QSslConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQSslConfigurationFromPointer(ptr unsafe.Pointer) *QSslConfiguration {
	var n = new(QSslConfiguration)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QSslConfiguration__NextProtocolNegotiationStatus
//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int64

const (
	QSslConfiguration__NextProtocolNegotiationNone        QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

func NewQSslConfiguration() *QSslConfiguration {
	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration())
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func NewQSslConfiguration2(other QSslConfiguration_ITF) *QSslConfiguration {
	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration2(PointerFromQSslConfiguration(other)))
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) AllowedNextProtocols() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslConfigurationFromPointer(l.data).allowedNextProtocols_atList(i)
			}
			return out
		}(C.QSslConfiguration_AllowedNextProtocols(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) CaCertificates() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			var out = make([]*QSslCertificate, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslConfigurationFromPointer(l.data).caCertificates_atList(i)
			}
			return out
		}(C.QSslConfiguration_CaCertificates(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) Ciphers() []*QSslCipher {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCipher {
			var out = make([]*QSslCipher, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslConfigurationFromPointer(l.data).ciphers_atList(i)
			}
			return out
		}(C.QSslConfiguration_Ciphers(ptr.Pointer()))
	}
	return nil
}

func QSslConfiguration_DefaultConfiguration() *QSslConfiguration {
	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_QSslConfiguration_DefaultConfiguration())
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) DefaultConfiguration() *QSslConfiguration {
	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_QSslConfiguration_DefaultConfiguration())
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) EphemeralServerKey() *QSslKey {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslConfiguration_EphemeralServerKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslConfiguration_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslConfiguration) LocalCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_LocalCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) LocalCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			var out = make([]*QSslCertificate, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslConfigurationFromPointer(l.data).localCertificateChain_atList(i)
			}
			return out
		}(C.QSslConfiguration_LocalCertificateChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) NextNegotiatedProtocol() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslConfiguration_NextNegotiatedProtocol(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) NextProtocolNegotiationStatus() QSslConfiguration__NextProtocolNegotiationStatus {
	if ptr.Pointer() != nil {
		return QSslConfiguration__NextProtocolNegotiationStatus(C.QSslConfiguration_NextProtocolNegotiationStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_PeerCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) PeerCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			var out = make([]*QSslCertificate, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslConfigurationFromPointer(l.data).peerCertificateChain_atList(i)
			}
			return out
		}(C.QSslConfiguration_PeerCertificateChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslConfiguration_PeerVerifyDepth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslConfiguration_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PrivateKey() *QSslKey {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslConfiguration_PrivateKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) Protocol() QSsl__SslProtocol {
	if ptr.Pointer() != nil {
		return QSsl__SslProtocol(C.QSslConfiguration_Protocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) SessionCipher() *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslConfiguration_SessionCipher(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) SessionProtocol() QSsl__SslProtocol {
	if ptr.Pointer() != nil {
		return QSsl__SslProtocol(C.QSslConfiguration_SessionProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) SessionTicket() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslConfiguration_SessionTicket(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) SessionTicketLifeTimeHint() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslConfiguration_SessionTicketLifeTimeHint(ptr.Pointer())))
	}
	return 0
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslConfiguration) SetProtocol(protocol QSsl__SslProtocol) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetProtocol(ptr.Pointer(), C.longlong(protocol))
	}
}

func (ptr *QSslConfiguration) SetSessionTicket(sessionTicket core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetSessionTicket(ptr.Pointer(), core.PointerFromQByteArray(sessionTicket))
	}
}

func (ptr *QSslConfiguration) SetSslOption(option QSsl__SslOption, on bool) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetSslOption(ptr.Pointer(), C.longlong(option), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func QSslConfiguration_SupportedCiphers() []*QSslCipher {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCipher {
		var out = make([]*QSslCipher, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslConfigurationFromPointer(l.data).supportedCiphers_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SupportedCiphers())
}

func (ptr *QSslConfiguration) SupportedCiphers() []*QSslCipher {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCipher {
		var out = make([]*QSslCipher, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslConfigurationFromPointer(l.data).supportedCiphers_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SupportedCiphers())
}

func (ptr *QSslConfiguration) Swap(other QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_Swap(ptr.Pointer(), PointerFromQSslConfiguration(other))
	}
}

func QSslConfiguration_SystemCaCertificates() []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslConfigurationFromPointer(l.data).systemCaCertificates_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SystemCaCertificates())
}

func (ptr *QSslConfiguration) SystemCaCertificates() []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		var out = make([]*QSslCertificate, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSslConfigurationFromPointer(l.data).systemCaCertificates_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SystemCaCertificates())
}

func (ptr *QSslConfiguration) TestSslOption(option QSsl__SslOption) bool {
	if ptr.Pointer() != nil {
		return C.QSslConfiguration_TestSslOption(ptr.Pointer(), C.longlong(option)) != 0
	}
	return false
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_DestroyQSslConfiguration(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslConfiguration) allowedNextProtocols_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslConfiguration_allowedNextProtocols_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) caCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_caCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslConfiguration_ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) ellipticCurves_atList(i int) *QSslEllipticCurve {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslConfiguration_ellipticCurves_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) localCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_localCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) peerCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_peerCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) supportedCiphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslConfiguration_supportedCiphers_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) supportedEllipticCurves_atList(i int) *QSslEllipticCurve {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslConfiguration_supportedEllipticCurves_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) systemCaCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_systemCaCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

type QSslEllipticCurve struct {
	ptr unsafe.Pointer
}

type QSslEllipticCurve_ITF interface {
	QSslEllipticCurve_PTR() *QSslEllipticCurve
}

func (ptr *QSslEllipticCurve) QSslEllipticCurve_PTR() *QSslEllipticCurve {
	return ptr
}

func (ptr *QSslEllipticCurve) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslEllipticCurve) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslEllipticCurve(ptr QSslEllipticCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslEllipticCurve_PTR().Pointer()
	}
	return nil
}

func NewQSslEllipticCurveFromPointer(ptr unsafe.Pointer) *QSslEllipticCurve {
	var n = new(QSslEllipticCurve)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslEllipticCurve) DestroyQSslEllipticCurve() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQSslEllipticCurve() *QSslEllipticCurve {
	var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_NewQSslEllipticCurve())
	runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func QSslEllipticCurve_FromLongName(name string) *QSslEllipticCurve {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromLongName(nameC))
	runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) FromLongName(name string) *QSslEllipticCurve {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromLongName(nameC))
	runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func QSslEllipticCurve_FromShortName(name string) *QSslEllipticCurve {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromShortName(nameC))
	runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) FromShortName(name string) *QSslEllipticCurve {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromShortName(nameC))
	runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {
	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsTlsNamedCurve(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) LongName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslEllipticCurve_LongName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslEllipticCurve) ShortName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslEllipticCurve_ShortName(ptr.Pointer()))
	}
	return ""
}

type QSslError struct {
	ptr unsafe.Pointer
}

type QSslError_ITF interface {
	QSslError_PTR() *QSslError
}

func (ptr *QSslError) QSslError_PTR() *QSslError {
	return ptr
}

func (ptr *QSslError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslError(ptr QSslError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslError_PTR().Pointer()
	}
	return nil
}

func NewQSslErrorFromPointer(ptr unsafe.Pointer) *QSslError {
	var n = new(QSslError)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QSslError__SslError
//QSslError::SslError
type QSslError__SslError int64

const (
	QSslError__NoError                             QSslError__SslError = QSslError__SslError(0)
	QSslError__UnableToGetIssuerCertificate        QSslError__SslError = QSslError__SslError(1)
	QSslError__UnableToDecryptCertificateSignature QSslError__SslError = QSslError__SslError(2)
	QSslError__UnableToDecodeIssuerPublicKey       QSslError__SslError = QSslError__SslError(3)
	QSslError__CertificateSignatureFailed          QSslError__SslError = QSslError__SslError(4)
	QSslError__CertificateNotYetValid              QSslError__SslError = QSslError__SslError(5)
	QSslError__CertificateExpired                  QSslError__SslError = QSslError__SslError(6)
	QSslError__InvalidNotBeforeField               QSslError__SslError = QSslError__SslError(7)
	QSslError__InvalidNotAfterField                QSslError__SslError = QSslError__SslError(8)
	QSslError__SelfSignedCertificate               QSslError__SslError = QSslError__SslError(9)
	QSslError__SelfSignedCertificateInChain        QSslError__SslError = QSslError__SslError(10)
	QSslError__UnableToGetLocalIssuerCertificate   QSslError__SslError = QSslError__SslError(11)
	QSslError__UnableToVerifyFirstCertificate      QSslError__SslError = QSslError__SslError(12)
	QSslError__CertificateRevoked                  QSslError__SslError = QSslError__SslError(13)
	QSslError__InvalidCaCertificate                QSslError__SslError = QSslError__SslError(14)
	QSslError__PathLengthExceeded                  QSslError__SslError = QSslError__SslError(15)
	QSslError__InvalidPurpose                      QSslError__SslError = QSslError__SslError(16)
	QSslError__CertificateUntrusted                QSslError__SslError = QSslError__SslError(17)
	QSslError__CertificateRejected                 QSslError__SslError = QSslError__SslError(18)
	QSslError__SubjectIssuerMismatch               QSslError__SslError = QSslError__SslError(19)
	QSslError__AuthorityIssuerSerialNumberMismatch QSslError__SslError = QSslError__SslError(20)
	QSslError__NoPeerCertificate                   QSslError__SslError = QSslError__SslError(21)
	QSslError__HostNameMismatch                    QSslError__SslError = QSslError__SslError(22)
	QSslError__NoSslSupport                        QSslError__SslError = QSslError__SslError(23)
	QSslError__CertificateBlacklisted              QSslError__SslError = QSslError__SslError(24)
	QSslError__UnspecifiedError                    QSslError__SslError = QSslError__SslError(-1)
)

func NewQSslError() *QSslError {
	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError())
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError2(error QSslError__SslError) *QSslError {
	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError2(C.longlong(error)))
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificate_ITF) *QSslError {
	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError3(C.longlong(error), PointerFromQSslCertificate(certificate)))
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError4(other QSslError_ITF) *QSslError {
	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError4(PointerFromQSslError(other)))
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func (ptr *QSslError) Certificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslError_Certificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslError) Error() QSslError__SslError {
	if ptr.Pointer() != nil {
		return QSslError__SslError(C.QSslError_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslError) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslError_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslError) Swap(other QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QSslError_Swap(ptr.Pointer(), PointerFromQSslError(other))
	}
}

func (ptr *QSslError) DestroyQSslError() {
	if ptr.Pointer() != nil {
		C.QSslError_DestroyQSslError(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslKey struct {
	ptr unsafe.Pointer
}

type QSslKey_ITF interface {
	QSslKey_PTR() *QSslKey
}

func (ptr *QSslKey) QSslKey_PTR() *QSslKey {
	return ptr
}

func (ptr *QSslKey) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslKey) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslKey(ptr QSslKey_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslKey_PTR().Pointer()
	}
	return nil
}

func NewQSslKeyFromPointer(ptr unsafe.Pointer) *QSslKey {
	var n = new(QSslKey)
	n.SetPointer(ptr)
	return n
}
func NewQSslKey() *QSslKey {
	var tmpValue = NewQSslKeyFromPointer(C.QSslKey_NewQSslKey())
	runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey3(device core.QIODevice_ITF, algorithm QSsl__KeyAlgorithm, encoding QSsl__EncodingFormat, ty QSsl__KeyType, passPhrase core.QByteArray_ITF) *QSslKey {
	var tmpValue = NewQSslKeyFromPointer(C.QSslKey_NewQSslKey3(core.PointerFromQIODevice(device), C.longlong(algorithm), C.longlong(encoding), C.longlong(ty), core.PointerFromQByteArray(passPhrase)))
	runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey2(encoded core.QByteArray_ITF, algorithm QSsl__KeyAlgorithm, encoding QSsl__EncodingFormat, ty QSsl__KeyType, passPhrase core.QByteArray_ITF) *QSslKey {
	var tmpValue = NewQSslKeyFromPointer(C.QSslKey_NewQSslKey2(core.PointerFromQByteArray(encoded), C.longlong(algorithm), C.longlong(encoding), C.longlong(ty), core.PointerFromQByteArray(passPhrase)))
	runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {
	var tmpValue = NewQSslKeyFromPointer(C.QSslKey_NewQSslKey5(PointerFromQSslKey(other)))
	runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func (ptr *QSslKey) Algorithm() QSsl__KeyAlgorithm {
	if ptr.Pointer() != nil {
		return QSsl__KeyAlgorithm(C.QSslKey_Algorithm(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslKey) Clear() {
	if ptr.Pointer() != nil {
		C.QSslKey_Clear(ptr.Pointer())
	}
}

func (ptr *QSslKey) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslKey_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslKey) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslKey_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslKey) Swap(other QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslKey_Swap(ptr.Pointer(), PointerFromQSslKey(other))
	}
}

func (ptr *QSslKey) ToDer(passPhrase core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslKey_ToDer(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslKey) ToPem(passPhrase core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslKey_ToPem(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslKey) Type() QSsl__KeyType {
	if ptr.Pointer() != nil {
		return QSsl__KeyType(C.QSslKey_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslKey) DestroyQSslKey() {
	if ptr.Pointer() != nil {
		C.QSslKey_DestroyQSslKey(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslPreSharedKeyAuthenticator struct {
	ptr unsafe.Pointer
}

type QSslPreSharedKeyAuthenticator_ITF interface {
	QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator
}

func (ptr *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator {
	return ptr
}

func (ptr *QSslPreSharedKeyAuthenticator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslPreSharedKeyAuthenticator(ptr QSslPreSharedKeyAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslPreSharedKeyAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	var n = new(QSslPreSharedKeyAuthenticator)
	n.SetPointer(ptr)
	return n
}
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	var tmpValue = NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator())
	runtime.SetFinalizer(tmpValue, (*QSslPreSharedKeyAuthenticator).DestroyQSslPreSharedKeyAuthenticator)
	return tmpValue
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {
	var tmpValue = NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))
	runtime.SetFinalizer(tmpValue, (*QSslPreSharedKeyAuthenticator).DestroyQSslPreSharedKeyAuthenticator)
	return tmpValue
}

func (ptr *QSslPreSharedKeyAuthenticator) Identity() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_Identity(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) IdentityHint() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_IdentityHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslPreSharedKeyAuthenticator_MaximumIdentityLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) PreSharedKey() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_PreSharedKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) SetIdentity(identity core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetIdentity(ptr.Pointer(), core.PointerFromQByteArray(identity))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetPreSharedKey(ptr.Pointer(), core.PointerFromQByteArray(preSharedKey))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) Swap(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_Swap(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) DestroyQSslPreSharedKeyAuthenticator() {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslSocket struct {
	QTcpSocket
}

type QSslSocket_ITF interface {
	QTcpSocket_ITF
	QSslSocket_PTR() *QSslSocket
}

func (ptr *QSslSocket) QSslSocket_PTR() *QSslSocket {
	return ptr
}

func (ptr *QSslSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QSslSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTcpSocket_PTR().SetPointer(p)
	}
}

func PointerFromQSslSocket(ptr QSslSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslSocket_PTR().Pointer()
	}
	return nil
}

func NewQSslSocketFromPointer(ptr unsafe.Pointer) *QSslSocket {
	var n = new(QSslSocket)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QSslSocket__PeerVerifyMode
//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int64

const (
	QSslSocket__VerifyNone     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(3)
)

//go:generate stringer -type=QSslSocket__SslMode
//QSslSocket::SslMode
type QSslSocket__SslMode int64

const (
	QSslSocket__UnencryptedMode QSslSocket__SslMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   QSslSocket__SslMode = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   QSslSocket__SslMode = QSslSocket__SslMode(2)
)

func NewQSslSocket(parent core.QObject_ITF) *QSslSocket {
	var tmpValue = NewQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSslSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) AddCaCertificates(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		return C.QSslSocket_AddCaCertificates(ptr.Pointer(), pathC, C.longlong(format), C.longlong(syntax)) != 0
	}
	return false
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func QSslSocket_AddDefaultCaCertificates(path string, encoding QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	return C.QSslSocket_QSslSocket_AddDefaultCaCertificates(pathC, C.longlong(encoding), C.longlong(syntax)) != 0
}

func (ptr *QSslSocket) AddDefaultCaCertificates(path string, encoding QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {
	var pathC = C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	return C.QSslSocket_QSslSocket_AddDefaultCaCertificates(pathC, C.longlong(encoding), C.longlong(syntax)) != 0
}

//export callbackQSslSocket_AtEnd
func callbackQSslSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QSslSocket) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::atEnd", f)
	}
}

func (ptr *QSslSocket) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::atEnd")
	}
}

func (ptr *QSslSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_BytesAvailable
func callbackQSslSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QSslSocket) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::bytesAvailable", f)
	}
}

func (ptr *QSslSocket) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::bytesAvailable")
	}
}

func (ptr *QSslSocket) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_BytesToWrite
func callbackQSslSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QSslSocket) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::bytesToWrite", f)
	}
}

func (ptr *QSslSocket) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::bytesToWrite")
	}
}

func (ptr *QSslSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_CanReadLine
func callbackQSslSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QSslSocket) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::canReadLine", f)
	}
}

func (ptr *QSslSocket) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::canReadLine")
	}
}

func (ptr *QSslSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_Close
func callbackQSslSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QSslSocket) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::close", f)
	}
}

func (ptr *QSslSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::close")
	}
}

func (ptr *QSslSocket) Close() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Close(ptr.Pointer())
	}
}

func (ptr *QSslSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16, mode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_ConnectToHostEncrypted(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(mode), C.longlong(protocol))
	}
}

func (ptr *QSslSocket) ConnectToHostEncrypted2(hostName string, port uint16, sslPeerName string, mode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		var sslPeerNameC = C.CString(sslPeerName)
		defer C.free(unsafe.Pointer(sslPeerNameC))
		C.QSslSocket_ConnectToHostEncrypted2(ptr.Pointer(), hostNameC, C.ushort(port), sslPeerNameC, C.longlong(mode), C.longlong(protocol))
	}
}

//export callbackQSslSocket_Encrypted
func callbackQSslSocket_Encrypted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::encrypted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encrypted", f)
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encrypted")
	}
}

func (ptr *QSslSocket) Encrypted() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Encrypted(ptr.Pointer())
	}
}

func (ptr *QSslSocket) EncryptedBytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_EncryptedBytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) EncryptedBytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_EncryptedBytesToWrite(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_EncryptedBytesWritten
func callbackQSslSocket_EncryptedBytesWritten(ptr unsafe.Pointer, written C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::encryptedBytesWritten"); signal != nil {
		signal.(func(int64))(int64(written))
	}

}

func (ptr *QSslSocket) ConnectEncryptedBytesWritten(f func(written int64)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncryptedBytesWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encryptedBytesWritten", f)
	}
}

func (ptr *QSslSocket) DisconnectEncryptedBytesWritten() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncryptedBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encryptedBytesWritten")
	}
}

func (ptr *QSslSocket) EncryptedBytesWritten(written int64) {
	if ptr.Pointer() != nil {
		C.QSslSocket_EncryptedBytesWritten(ptr.Pointer(), C.longlong(written))
	}
}

func (ptr *QSslSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_IgnoreSslErrors
func callbackQSslSocket_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::ignoreSslErrors"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QSslSocket) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::ignoreSslErrors", f)
	}
}

func (ptr *QSslSocket) DisconnectIgnoreSslErrors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::ignoreSslErrors")
	}
}

func (ptr *QSslSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QSslSocket) IgnoreSslErrorsDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrorsDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_IsEncrypted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) LocalCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_LocalCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) LocalCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			var out = make([]*QSslCertificate, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslSocketFromPointer(l.data).localCertificateChain_atList(i)
			}
			return out
		}(C.QSslSocket_LocalCertificateChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_ModeChanged
func callbackQSslSocket_ModeChanged(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::modeChanged"); signal != nil {
		signal.(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
	}

}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectModeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::modeChanged", f)
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::modeChanged")
	}
}

func (ptr *QSslSocket) ModeChanged(mode QSslSocket__SslMode) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ModeChanged(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslSocket) PeerCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_PeerCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) PeerCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			var out = make([]*QSslCertificate, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslSocketFromPointer(l.data).peerCertificateChain_atList(i)
			}
			return out
		}(C.QSslSocket_PeerCertificateChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslSocket_PeerVerifyDepth(ptr.Pointer())))
	}
	return 0
}

//export callbackQSslSocket_PeerVerifyError
func callbackQSslSocket_PeerVerifyError(ptr unsafe.Pointer, error unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::peerVerifyError"); signal != nil {
		signal.(func(*QSslError))(NewQSslErrorFromPointer(error))
	}

}

func (ptr *QSslSocket) ConnectPeerVerifyError(f func(error *QSslError)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPeerVerifyError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::peerVerifyError", f)
	}
}

func (ptr *QSslSocket) DisconnectPeerVerifyError() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPeerVerifyError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::peerVerifyError")
	}
}

func (ptr *QSslSocket) PeerVerifyError(error QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_PeerVerifyError(ptr.Pointer(), PointerFromQSslError(error))
	}
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslSocket_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslSocket_PeerVerifyName(ptr.Pointer()))
	}
	return ""
}

//export callbackQSslSocket_PreSharedKeyAuthenticationRequired
func callbackQSslSocket_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslSocket) PrivateKey() *QSslKey {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslSocket_PrivateKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) Protocol() QSsl__SslProtocol {
	if ptr.Pointer() != nil {
		return QSsl__SslProtocol(C.QSslSocket_Protocol(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_ReadData
func callbackQSslSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxlen C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxlen)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQSslSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxlen)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QSslSocket) ConnectReadData(f func(data *string, maxlen int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::readData", f)
	}
}

func (ptr *QSslSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::readData")
	}
}

func (ptr *QSslSocket) ReadData(data *string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxlen)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QSslSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxlen)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QSslSocket) ReadDataDefault(data *string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxlen)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QSslSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxlen)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQSslSocket_Resume
func callbackQSslSocket_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QSslSocket) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::resume", f)
	}
}

func (ptr *QSslSocket) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::resume")
	}
}

func (ptr *QSslSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ResumeDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) SessionCipher() *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslSocket_SessionCipher(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SessionProtocol() QSsl__SslProtocol {
	if ptr.Pointer() != nil {
		return QSsl__SslProtocol(C.QSslSocket_SessionProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) SetLocalCertificate2(path string, format QSsl__EncodingFormat) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QSslSocket_SetLocalCertificate2(ptr.Pointer(), pathC, C.longlong(format))
	}
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_SetPeerVerifyName(ptr.Pointer(), hostNameC)
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslSocket) SetPrivateKey2(fileName string, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, passPhrase core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QSslSocket_SetPrivateKey2(ptr.Pointer(), fileNameC, C.longlong(algorithm), C.longlong(format), core.PointerFromQByteArray(passPhrase))
	}
}

func (ptr *QSslSocket) SetProtocol(protocol QSsl__SslProtocol) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetProtocol(ptr.Pointer(), C.longlong(protocol))
	}
}

//export callbackQSslSocket_SetReadBufferSize
func callbackQSslSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQSslSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QSslSocket) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setReadBufferSize", f)
	}
}

func (ptr *QSslSocket) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setReadBufferSize")
	}
}

func (ptr *QSslSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QSslSocket) SetReadBufferSizeDefault(size int64) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQSslSocket_SetSocketOption
func callbackQSslSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQSslSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QSslSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setSocketOption", f)
	}
}

func (ptr *QSslSocket) DisconnectSetSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setSocketOption")
	}
}

func (ptr *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

//export callbackQSslSocket_SocketOption
func callbackQSslSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQSslSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QSslSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::socketOption", f)
	}
}

func (ptr *QSslSocket) DisconnectSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::socketOption")
	}
}

func (ptr *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSslSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSslSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SslConfiguration() *QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslConfigurationFromPointer(C.QSslSocket_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SslErrors() []*QSslError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslError {
			var out = make([]*QSslError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSslSocketFromPointer(l.data).sslErrors_atList(i)
			}
			return out
		}(C.QSslSocket_SslErrors(ptr.Pointer()))
	}
	return nil
}

func QSslSocket_SslLibraryBuildVersionNumber() int {
	return int(int32(C.QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()))
}

func (ptr *QSslSocket) SslLibraryBuildVersionNumber() int {
	return int(int32(C.QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()))
}

func QSslSocket_SslLibraryBuildVersionString() string {
	return cGoUnpackString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func (ptr *QSslSocket) SslLibraryBuildVersionString() string {
	return cGoUnpackString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func QSslSocket_SslLibraryVersionNumber() int {
	return int(int32(C.QSslSocket_QSslSocket_SslLibraryVersionNumber()))
}

func (ptr *QSslSocket) SslLibraryVersionNumber() int {
	return int(int32(C.QSslSocket_QSslSocket_SslLibraryVersionNumber()))
}

func QSslSocket_SslLibraryVersionString() string {
	return cGoUnpackString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

func (ptr *QSslSocket) SslLibraryVersionString() string {
	return cGoUnpackString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

//export callbackQSslSocket_StartClientEncryption
func callbackQSslSocket_StartClientEncryption(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::startClientEncryption"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).StartClientEncryptionDefault()
	}
}

func (ptr *QSslSocket) ConnectStartClientEncryption(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startClientEncryption", f)
	}
}

func (ptr *QSslSocket) DisconnectStartClientEncryption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startClientEncryption")
	}
}

func (ptr *QSslSocket) StartClientEncryption() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryption(ptr.Pointer())
	}
}

func (ptr *QSslSocket) StartClientEncryptionDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryptionDefault(ptr.Pointer())
	}
}

//export callbackQSslSocket_StartServerEncryption
func callbackQSslSocket_StartServerEncryption(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::startServerEncryption"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).StartServerEncryptionDefault()
	}
}

func (ptr *QSslSocket) ConnectStartServerEncryption(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startServerEncryption", f)
	}
}

func (ptr *QSslSocket) DisconnectStartServerEncryption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startServerEncryption")
	}
}

func (ptr *QSslSocket) StartServerEncryption() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryption(ptr.Pointer())
	}
}

func (ptr *QSslSocket) StartServerEncryptionDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryptionDefault(ptr.Pointer())
	}
}

func QSslSocket_SupportsSsl() bool {
	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) SupportsSsl() bool {
	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

//export callbackQSslSocket_WaitForBytesWritten
func callbackQSslSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QSslSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForBytesWritten", f)
	}
}

func (ptr *QSslSocket) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForBytesWritten")
	}
}

func (ptr *QSslSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSslSocket_WaitForConnected
func callbackQSslSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QSslSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForConnected", f)
	}
}

func (ptr *QSslSocket) DisconnectWaitForConnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForConnected")
	}
}

func (ptr *QSslSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForConnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSslSocket_WaitForDisconnected
func callbackQSslSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QSslSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForDisconnected", f)
	}
}

func (ptr *QSslSocket) DisconnectWaitForDisconnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForDisconnected")
	}
}

func (ptr *QSslSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForDisconnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForEncrypted(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSslSocket_WaitForReadyRead
func callbackQSslSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QSslSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForReadyRead", f)
	}
}

func (ptr *QSslSocket) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForReadyRead")
	}
}

func (ptr *QSslSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSslSocket_WriteData
func callbackQSslSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, len C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(len)))
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(len)))
}

func (ptr *QSslSocket) ConnectWriteData(f func(data string, len int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::writeData", f)
	}
}

func (ptr *QSslSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::writeData")
	}
}

func (ptr *QSslSocket) WriteData(data string, len int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QSslSocket_WriteData(ptr.Pointer(), dataC, C.longlong(len)))
	}
	return 0
}

func (ptr *QSslSocket) WriteDataDefault(data string, len int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QSslSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(len)))
	}
	return 0
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DestroyQSslSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslSocket) caCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_caCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslSocket_ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) defaultCaCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_defaultCaCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) defaultCiphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslSocket_defaultCiphers_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) localCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_localCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) peerCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_peerCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) sslErrors_atList(i int) *QSslError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslErrorFromPointer(C.QSslSocket_sslErrors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) supportedCiphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslSocket_supportedCiphers_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) systemCaCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_systemCaCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

//export callbackQSslSocket_ConnectToHost2
func callbackQSslSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQSslSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QSslSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost2", f)
	}
}

func (ptr *QSslSocket) DisconnectConnectToHost2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost2")
	}
}

func (ptr *QSslSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QSslSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQSslSocket_ConnectToHost
func callbackQSslSocket_ConnectToHost(ptr unsafe.Pointer, hostName C.struct_QtNetwork_PackedString, port C.ushort, openMode C.longlong, protocol C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQSslSocketFromPointer(ptr).ConnectToHostDefault(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QSslSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost", f)
	}
}

func (ptr *QSslSocket) DisconnectConnectToHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost")
	}
}

func (ptr *QSslSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QSslSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQSslSocket_DisconnectFromHost
func callbackQSslSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QSslSocket) ConnectDisconnectFromHost(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectFromHost", f)
	}
}

func (ptr *QSslSocket) DisconnectDisconnectFromHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectFromHost")
	}
}

func (ptr *QSslSocket) DisconnectFromHost() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QSslSocket) DisconnectFromHostDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQSslSocket_IsSequential
func callbackQSslSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QSslSocket) ConnectIsSequential(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::isSequential", f)
	}
}

func (ptr *QSslSocket) DisconnectIsSequential() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::isSequential")
	}
}

func (ptr *QSslSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_ReadLineData
func callbackQSslSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxlen C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxlen)))
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxlen)))
}

func (ptr *QSslSocket) ConnectReadLineData(f func(data string, maxlen int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::readLineData", f)
	}
}

func (ptr *QSslSocket) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::readLineData")
	}
}

func (ptr *QSslSocket) ReadLineData(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QSslSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QSslSocket) ReadLineDataDefault(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QSslSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

//export callbackQSslSocket_Open
func callbackQSslSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QSslSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::open", f)
	}
}

func (ptr *QSslSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::open")
	}
}

func (ptr *QSslSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QSslSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQSslSocket_Pos
func callbackQSslSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).PosDefault())
}

func (ptr *QSslSocket) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::pos", f)
	}
}

func (ptr *QSslSocket) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::pos")
	}
}

func (ptr *QSslSocket) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_Reset
func callbackQSslSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QSslSocket) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::reset", f)
	}
}

func (ptr *QSslSocket) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::reset")
	}
}

func (ptr *QSslSocket) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_Seek
func callbackQSslSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QSslSocket) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::seek", f)
	}
}

func (ptr *QSslSocket) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::seek")
	}
}

func (ptr *QSslSocket) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QSslSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQSslSocket_Size
func callbackQSslSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QSslSocket) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::size", f)
	}
}

func (ptr *QSslSocket) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::size")
	}
}

func (ptr *QSslSocket) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_TimerEvent
func callbackQSslSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::timerEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::timerEvent")
	}
}

func (ptr *QSslSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSslSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSslSocket_ChildEvent
func callbackQSslSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::childEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::childEvent")
	}
}

func (ptr *QSslSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSslSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSslSocket_ConnectNotify
func callbackQSslSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSslSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSslSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectNotify", f)
	}
}

func (ptr *QSslSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectNotify")
	}
}

func (ptr *QSslSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSslSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSslSocket_CustomEvent
func callbackQSslSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::customEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::customEvent")
	}
}

func (ptr *QSslSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSslSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSslSocket_DeleteLater
func callbackQSslSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSslSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::deleteLater", f)
	}
}

func (ptr *QSslSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::deleteLater")
	}
}

func (ptr *QSslSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSslSocket_DisconnectNotify
func callbackQSslSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSslSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSslSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectNotify", f)
	}
}

func (ptr *QSslSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectNotify")
	}
}

func (ptr *QSslSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSslSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSslSocket_Event
func callbackQSslSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSslSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::event", f)
	}
}

func (ptr *QSslSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::event")
	}
}

func (ptr *QSslSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSslSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSslSocket_EventFilter
func callbackQSslSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSslSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::eventFilter", f)
	}
}

func (ptr *QSslSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::eventFilter")
	}
}

func (ptr *QSslSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSslSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSslSocket_MetaObject
func callbackQSslSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSslSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSslSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::metaObject", f)
	}
}

func (ptr *QSslSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::metaObject")
	}
}

func (ptr *QSslSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSslSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSslSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTcpServer struct {
	core.QObject
}

type QTcpServer_ITF interface {
	core.QObject_ITF
	QTcpServer_PTR() *QTcpServer
}

func (ptr *QTcpServer) QTcpServer_PTR() *QTcpServer {
	return ptr
}

func (ptr *QTcpServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTcpServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTcpServer(ptr QTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func NewQTcpServerFromPointer(ptr unsafe.Pointer) *QTcpServer {
	var n = new(QTcpServer)
	n.SetPointer(ptr)
	return n
}
func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {
	var tmpValue = NewQTcpServerFromPointer(C.QTcpServer_NewQTcpServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTcpServer_AcceptError
func callbackQTcpServer_AcceptError(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::acceptError"); signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::acceptError", f)
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::acceptError")
	}
}

func (ptr *QTcpServer) AcceptError(socketError QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QTcpServer_AcceptError(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QTcpServer) AddPendingConnection(socket QTcpSocket_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_AddPendingConnection(ptr.Pointer(), PointerFromQTcpSocket(socket))
	}
}

func (ptr *QTcpServer) Close() {
	if ptr.Pointer() != nil {
		C.QTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTcpServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQTcpServer_HasPendingConnections
func callbackQTcpServer_HasPendingConnections(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::hasPendingConnections"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).HasPendingConnectionsDefault())))
}

func (ptr *QTcpServer) ConnectHasPendingConnections(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::hasPendingConnections", f)
	}
}

func (ptr *QTcpServer) DisconnectHasPendingConnections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::hasPendingConnections")
	}
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) HasPendingConnectionsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnectionsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) Listen(address QHostAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_Listen(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port)) != 0
	}
	return false
}

func (ptr *QTcpServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTcpServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQTcpServer_NewConnection
func callbackQTcpServer_NewConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::newConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::newConnection")
	}
}

func (ptr *QTcpServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QTcpServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQTcpServer_NextPendingConnection
func callbackQTcpServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::nextPendingConnection"); signal != nil {
		return PointerFromQTcpSocket(signal.(func() *QTcpSocket)())
	}

	return PointerFromQTcpSocket(NewQTcpServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QTcpServer) ConnectNextPendingConnection(f func() *QTcpSocket) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::nextPendingConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNextPendingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::nextPendingConnection")
	}
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) NextPendingConnectionDefault() *QTcpSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) PauseAccepting() {
	if ptr.Pointer() != nil {
		C.QTcpServer_PauseAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) Proxy() *QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QTcpServer_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QTcpServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ServerAddress() *QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QTcpServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QTcpServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QTcpServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QTcpServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut bool) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_WaitForNewConnection(ptr.Pointer(), C.int(int32(msec)), C.char(int8(qt.GoBoolToInt(timedOut)))) != 0
	}
	return false
}

//export callbackQTcpServer_DestroyQTcpServer
func callbackQTcpServer_DestroyQTcpServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::~QTcpServer"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpServerFromPointer(ptr).DestroyQTcpServerDefault()
	}
}

func (ptr *QTcpServer) ConnectDestroyQTcpServer(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::~QTcpServer", f)
	}
}

func (ptr *QTcpServer) DisconnectDestroyQTcpServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::~QTcpServer")
	}
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) DestroyQTcpServerDefault() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpServer_TimerEvent
func callbackQTcpServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::timerEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::timerEvent")
	}
}

func (ptr *QTcpServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTcpServer_ChildEvent
func callbackQTcpServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::childEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::childEvent")
	}
}

func (ptr *QTcpServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTcpServer_ConnectNotify
func callbackQTcpServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::connectNotify", f)
	}
}

func (ptr *QTcpServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::connectNotify")
	}
}

func (ptr *QTcpServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpServer_CustomEvent
func callbackQTcpServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::customEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::customEvent")
	}
}

func (ptr *QTcpServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTcpServer_DeleteLater
func callbackQTcpServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTcpServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::deleteLater", f)
	}
}

func (ptr *QTcpServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::deleteLater")
	}
}

func (ptr *QTcpServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpServer_DisconnectNotify
func callbackQTcpServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::disconnectNotify", f)
	}
}

func (ptr *QTcpServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::disconnectNotify")
	}
}

func (ptr *QTcpServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpServer_Event
func callbackQTcpServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTcpServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::event", f)
	}
}

func (ptr *QTcpServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::event")
	}
}

func (ptr *QTcpServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTcpServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTcpServer_EventFilter
func callbackQTcpServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTcpServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::eventFilter", f)
	}
}

func (ptr *QTcpServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::eventFilter")
	}
}

func (ptr *QTcpServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTcpServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTcpServer_MetaObject
func callbackQTcpServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTcpServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTcpServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::metaObject", f)
	}
}

func (ptr *QTcpServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::metaObject")
	}
}

func (ptr *QTcpServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocket_ITF interface {
	QAbstractSocket_ITF
	QTcpSocket_PTR() *QTcpSocket
}

func (ptr *QTcpSocket) QTcpSocket_PTR() *QTcpSocket {
	return ptr
}

func (ptr *QTcpSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QTcpSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSocket_PTR().SetPointer(p)
	}
}

func PointerFromQTcpSocket(ptr QTcpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = new(QTcpSocket)
	n.SetPointer(ptr)
	return n
}
func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {
	var tmpValue = NewQTcpSocketFromPointer(C.QTcpSocket_NewQTcpSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTcpSocket_DestroyQTcpSocket
func callbackQTcpSocket_DestroyQTcpSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::~QTcpSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DestroyQTcpSocketDefault()
	}
}

func (ptr *QTcpSocket) ConnectDestroyQTcpSocket(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::~QTcpSocket", f)
	}
}

func (ptr *QTcpSocket) DisconnectDestroyQTcpSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::~QTcpSocket")
	}
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) DestroyQTcpSocketDefault() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpSocket_AtEnd
func callbackQTcpSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QTcpSocket) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::atEnd", f)
	}
}

func (ptr *QTcpSocket) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::atEnd")
	}
}

func (ptr *QTcpSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTcpSocket_BytesAvailable
func callbackQTcpSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QTcpSocket) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::bytesAvailable", f)
	}
}

func (ptr *QTcpSocket) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::bytesAvailable")
	}
}

func (ptr *QTcpSocket) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQTcpSocket_BytesToWrite
func callbackQTcpSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QTcpSocket) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::bytesToWrite", f)
	}
}

func (ptr *QTcpSocket) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::bytesToWrite")
	}
}

func (ptr *QTcpSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQTcpSocket_CanReadLine
func callbackQTcpSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QTcpSocket) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::canReadLine", f)
	}
}

func (ptr *QTcpSocket) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::canReadLine")
	}
}

func (ptr *QTcpSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTcpSocket_Close
func callbackQTcpSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QTcpSocket) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::close", f)
	}
}

func (ptr *QTcpSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::close")
	}
}

func (ptr *QTcpSocket) Close() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_Close(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_CloseDefault(ptr.Pointer())
	}
}

//export callbackQTcpSocket_ConnectToHost2
func callbackQTcpSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQTcpSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QTcpSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost2", f)
	}
}

func (ptr *QTcpSocket) DisconnectConnectToHost2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost2")
	}
}

func (ptr *QTcpSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QTcpSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQTcpSocket_ConnectToHost
func callbackQTcpSocket_ConnectToHost(ptr unsafe.Pointer, hostName C.struct_QtNetwork_PackedString, port C.ushort, openMode C.longlong, protocol C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQTcpSocketFromPointer(ptr).ConnectToHostDefault(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QTcpSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost", f)
	}
}

func (ptr *QTcpSocket) DisconnectConnectToHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost")
	}
}

func (ptr *QTcpSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QTcpSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QTcpSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QTcpSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQTcpSocket_DisconnectFromHost
func callbackQTcpSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QTcpSocket) ConnectDisconnectFromHost(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectFromHost", f)
	}
}

func (ptr *QTcpSocket) DisconnectDisconnectFromHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectFromHost")
	}
}

func (ptr *QTcpSocket) DisconnectFromHost() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) DisconnectFromHostDefault() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQTcpSocket_IsSequential
func callbackQTcpSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QTcpSocket) ConnectIsSequential(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::isSequential", f)
	}
}

func (ptr *QTcpSocket) DisconnectIsSequential() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::isSequential")
	}
}

func (ptr *QTcpSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTcpSocket_ReadData
func callbackQTcpSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQTcpSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QTcpSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::readData", f)
	}
}

func (ptr *QTcpSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::readData")
	}
}

func (ptr *QTcpSocket) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QTcpSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QTcpSocket) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QTcpSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQTcpSocket_ReadLineData
func callbackQTcpSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxlen C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxlen)))
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxlen)))
}

func (ptr *QTcpSocket) ConnectReadLineData(f func(data string, maxlen int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::readLineData", f)
	}
}

func (ptr *QTcpSocket) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::readLineData")
	}
}

func (ptr *QTcpSocket) ReadLineData(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QTcpSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QTcpSocket) ReadLineDataDefault(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QTcpSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

//export callbackQTcpSocket_Resume
func callbackQTcpSocket_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QTcpSocket) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::resume", f)
	}
}

func (ptr *QTcpSocket) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::resume")
	}
}

func (ptr *QTcpSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ResumeDefault() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ResumeDefault(ptr.Pointer())
	}
}

//export callbackQTcpSocket_SetReadBufferSize
func callbackQTcpSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQTcpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QTcpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setReadBufferSize", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setReadBufferSize")
	}
}

func (ptr *QTcpSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QTcpSocket) SetReadBufferSizeDefault(size int64) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQTcpSocket_SetSocketOption
func callbackQTcpSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQTcpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QTcpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setSocketOption", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setSocketOption")
	}
}

func (ptr *QTcpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QTcpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

//export callbackQTcpSocket_SocketOption
func callbackQTcpSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQTcpSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QTcpSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::socketOption", f)
	}
}

func (ptr *QTcpSocket) DisconnectSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::socketOption")
	}
}

func (ptr *QTcpSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QTcpSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QTcpSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQTcpSocket_WaitForBytesWritten
func callbackQTcpSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QTcpSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForBytesWritten", f)
	}
}

func (ptr *QTcpSocket) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForBytesWritten")
	}
}

func (ptr *QTcpSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QTcpSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQTcpSocket_WaitForConnected
func callbackQTcpSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QTcpSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForConnected", f)
	}
}

func (ptr *QTcpSocket) DisconnectWaitForConnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForConnected")
	}
}

func (ptr *QTcpSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QTcpSocket) WaitForConnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQTcpSocket_WaitForDisconnected
func callbackQTcpSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QTcpSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForDisconnected", f)
	}
}

func (ptr *QTcpSocket) DisconnectWaitForDisconnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForDisconnected")
	}
}

func (ptr *QTcpSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QTcpSocket) WaitForDisconnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQTcpSocket_WaitForReadyRead
func callbackQTcpSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QTcpSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForReadyRead", f)
	}
}

func (ptr *QTcpSocket) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForReadyRead")
	}
}

func (ptr *QTcpSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QTcpSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQTcpSocket_WriteData
func callbackQTcpSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, size C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(size)))
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(size)))
}

func (ptr *QTcpSocket) ConnectWriteData(f func(data string, size int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::writeData", f)
	}
}

func (ptr *QTcpSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::writeData")
	}
}

func (ptr *QTcpSocket) WriteData(data string, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QTcpSocket_WriteData(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

func (ptr *QTcpSocket) WriteDataDefault(data string, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QTcpSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

//export callbackQTcpSocket_Open
func callbackQTcpSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QTcpSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::open", f)
	}
}

func (ptr *QTcpSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::open")
	}
}

func (ptr *QTcpSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QTcpSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQTcpSocket_Pos
func callbackQTcpSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).PosDefault())
}

func (ptr *QTcpSocket) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::pos", f)
	}
}

func (ptr *QTcpSocket) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::pos")
	}
}

func (ptr *QTcpSocket) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQTcpSocket_Reset
func callbackQTcpSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QTcpSocket) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::reset", f)
	}
}

func (ptr *QTcpSocket) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::reset")
	}
}

func (ptr *QTcpSocket) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTcpSocket_Seek
func callbackQTcpSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QTcpSocket) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::seek", f)
	}
}

func (ptr *QTcpSocket) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::seek")
	}
}

func (ptr *QTcpSocket) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QTcpSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQTcpSocket_Size
func callbackQTcpSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QTcpSocket) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::size", f)
	}
}

func (ptr *QTcpSocket) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::size")
	}
}

func (ptr *QTcpSocket) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQTcpSocket_TimerEvent
func callbackQTcpSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::timerEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::timerEvent")
	}
}

func (ptr *QTcpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTcpSocket_ChildEvent
func callbackQTcpSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::childEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::childEvent")
	}
}

func (ptr *QTcpSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTcpSocket_ConnectNotify
func callbackQTcpSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectNotify", f)
	}
}

func (ptr *QTcpSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectNotify")
	}
}

func (ptr *QTcpSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpSocket_CustomEvent
func callbackQTcpSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::customEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::customEvent")
	}
}

func (ptr *QTcpSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTcpSocket_DeleteLater
func callbackQTcpSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTcpSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::deleteLater", f)
	}
}

func (ptr *QTcpSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::deleteLater")
	}
}

func (ptr *QTcpSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpSocket_DisconnectNotify
func callbackQTcpSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectNotify", f)
	}
}

func (ptr *QTcpSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectNotify")
	}
}

func (ptr *QTcpSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpSocket_Event
func callbackQTcpSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTcpSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::event", f)
	}
}

func (ptr *QTcpSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::event")
	}
}

func (ptr *QTcpSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTcpSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTcpSocket_EventFilter
func callbackQTcpSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTcpSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::eventFilter", f)
	}
}

func (ptr *QTcpSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::eventFilter")
	}
}

func (ptr *QTcpSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTcpSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTcpSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTcpSocket_MetaObject
func callbackQTcpSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTcpSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTcpSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::metaObject", f)
	}
}

func (ptr *QTcpSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::metaObject")
	}
}

func (ptr *QTcpSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket {
	return ptr
}

func (ptr *QUdpSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QUdpSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSocket_PTR().SetPointer(p)
	}
}

func PointerFromQUdpSocket(ptr QUdpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocket_PTR().Pointer()
	}
	return nil
}

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = new(QUdpSocket)
	n.SetPointer(ptr)
	return n
}
func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {
	var tmpValue = NewQUdpSocketFromPointer(C.QUdpSocket_NewQUdpSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_HasPendingDatagrams(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) MulticastInterface() *QNetworkInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkInterfaceFromPointer(C.QUdpSocket_MulticastInterface(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

func (ptr *QUdpSocket) PendingDatagramSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_PendingDatagramSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) ReadDatagram(data string, maxSize int64, address QHostAddress_ITF, port uint16) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_ReadDatagram(ptr.Pointer(), dataC, C.longlong(maxSize), PointerFromQHostAddress(address), C.ushort(port)))
	}
	return 0
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(ptr.Pointer(), PointerFromQNetworkInterface(iface))
	}
}

func (ptr *QUdpSocket) WriteDatagram2(datagram core.QByteArray_ITF, host QHostAddress_ITF, port uint16) int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_WriteDatagram2(ptr.Pointer(), core.PointerFromQByteArray(datagram), PointerFromQHostAddress(host), C.ushort(port)))
	}
	return 0
}

func (ptr *QUdpSocket) WriteDatagram(data string, size int64, address QHostAddress_ITF, port uint16) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_WriteDatagram(ptr.Pointer(), dataC, C.longlong(size), PointerFromQHostAddress(address), C.ushort(port)))
	}
	return 0
}

//export callbackQUdpSocket_DestroyQUdpSocket
func callbackQUdpSocket_DestroyQUdpSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::~QUdpSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DestroyQUdpSocketDefault()
	}
}

func (ptr *QUdpSocket) ConnectDestroyQUdpSocket(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::~QUdpSocket", f)
	}
}

func (ptr *QUdpSocket) DisconnectDestroyQUdpSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::~QUdpSocket")
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocketDefault() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQUdpSocket_AtEnd
func callbackQUdpSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QUdpSocket) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::atEnd", f)
	}
}

func (ptr *QUdpSocket) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::atEnd")
	}
}

func (ptr *QUdpSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQUdpSocket_BytesAvailable
func callbackQUdpSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QUdpSocket) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::bytesAvailable", f)
	}
}

func (ptr *QUdpSocket) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::bytesAvailable")
	}
}

func (ptr *QUdpSocket) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQUdpSocket_BytesToWrite
func callbackQUdpSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QUdpSocket) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::bytesToWrite", f)
	}
}

func (ptr *QUdpSocket) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::bytesToWrite")
	}
}

func (ptr *QUdpSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQUdpSocket_CanReadLine
func callbackQUdpSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QUdpSocket) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::canReadLine", f)
	}
}

func (ptr *QUdpSocket) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::canReadLine")
	}
}

func (ptr *QUdpSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQUdpSocket_Close
func callbackQUdpSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QUdpSocket) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::close", f)
	}
}

func (ptr *QUdpSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::close")
	}
}

func (ptr *QUdpSocket) Close() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_Close(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_CloseDefault(ptr.Pointer())
	}
}

//export callbackQUdpSocket_ConnectToHost2
func callbackQUdpSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQUdpSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QUdpSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost2", f)
	}
}

func (ptr *QUdpSocket) DisconnectConnectToHost2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost2")
	}
}

func (ptr *QUdpSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QUdpSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQUdpSocket_ConnectToHost
func callbackQUdpSocket_ConnectToHost(ptr unsafe.Pointer, hostName C.struct_QtNetwork_PackedString, port C.ushort, openMode C.longlong, protocol C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQUdpSocketFromPointer(ptr).ConnectToHostDefault(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QUdpSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost", f)
	}
}

func (ptr *QUdpSocket) DisconnectConnectToHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost")
	}
}

func (ptr *QUdpSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QUdpSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QUdpSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QUdpSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQUdpSocket_DisconnectFromHost
func callbackQUdpSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QUdpSocket) ConnectDisconnectFromHost(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectFromHost", f)
	}
}

func (ptr *QUdpSocket) DisconnectDisconnectFromHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectFromHost")
	}
}

func (ptr *QUdpSocket) DisconnectFromHost() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) DisconnectFromHostDefault() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQUdpSocket_IsSequential
func callbackQUdpSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QUdpSocket) ConnectIsSequential(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::isSequential", f)
	}
}

func (ptr *QUdpSocket) DisconnectIsSequential() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::isSequential")
	}
}

func (ptr *QUdpSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQUdpSocket_ReadData
func callbackQUdpSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQUdpSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QUdpSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::readData", f)
	}
}

func (ptr *QUdpSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::readData")
	}
}

func (ptr *QUdpSocket) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QUdpSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QUdpSocket) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QUdpSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQUdpSocket_ReadLineData
func callbackQUdpSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxlen C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxlen)))
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxlen)))
}

func (ptr *QUdpSocket) ConnectReadLineData(f func(data string, maxlen int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::readLineData", f)
	}
}

func (ptr *QUdpSocket) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::readLineData")
	}
}

func (ptr *QUdpSocket) ReadLineData(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QUdpSocket) ReadLineDataDefault(data string, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

//export callbackQUdpSocket_Resume
func callbackQUdpSocket_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QUdpSocket) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::resume", f)
	}
}

func (ptr *QUdpSocket) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::resume")
	}
}

func (ptr *QUdpSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ResumeDefault() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ResumeDefault(ptr.Pointer())
	}
}

//export callbackQUdpSocket_SetReadBufferSize
func callbackQUdpSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQUdpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QUdpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setReadBufferSize", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setReadBufferSize")
	}
}

func (ptr *QUdpSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QUdpSocket) SetReadBufferSizeDefault(size int64) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQUdpSocket_SetSocketOption
func callbackQUdpSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQUdpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QUdpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setSocketOption", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setSocketOption")
	}
}

func (ptr *QUdpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QUdpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

//export callbackQUdpSocket_SocketOption
func callbackQUdpSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQUdpSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QUdpSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::socketOption", f)
	}
}

func (ptr *QUdpSocket) DisconnectSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::socketOption")
	}
}

func (ptr *QUdpSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QUdpSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QUdpSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QUdpSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQUdpSocket_WaitForBytesWritten
func callbackQUdpSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QUdpSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForBytesWritten", f)
	}
}

func (ptr *QUdpSocket) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForBytesWritten")
	}
}

func (ptr *QUdpSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QUdpSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQUdpSocket_WaitForConnected
func callbackQUdpSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QUdpSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForConnected", f)
	}
}

func (ptr *QUdpSocket) DisconnectWaitForConnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForConnected")
	}
}

func (ptr *QUdpSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QUdpSocket) WaitForConnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQUdpSocket_WaitForDisconnected
func callbackQUdpSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QUdpSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForDisconnected", f)
	}
}

func (ptr *QUdpSocket) DisconnectWaitForDisconnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForDisconnected")
	}
}

func (ptr *QUdpSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QUdpSocket) WaitForDisconnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQUdpSocket_WaitForReadyRead
func callbackQUdpSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QUdpSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForReadyRead", f)
	}
}

func (ptr *QUdpSocket) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForReadyRead")
	}
}

func (ptr *QUdpSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QUdpSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQUdpSocket_WriteData
func callbackQUdpSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, size C.longlong) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(size)))
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(size)))
}

func (ptr *QUdpSocket) ConnectWriteData(f func(data string, size int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::writeData", f)
	}
}

func (ptr *QUdpSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::writeData")
	}
}

func (ptr *QUdpSocket) WriteData(data string, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_WriteData(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

func (ptr *QUdpSocket) WriteDataDefault(data string, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

//export callbackQUdpSocket_Open
func callbackQUdpSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QUdpSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::open", f)
	}
}

func (ptr *QUdpSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::open")
	}
}

func (ptr *QUdpSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QUdpSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQUdpSocket_Pos
func callbackQUdpSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).PosDefault())
}

func (ptr *QUdpSocket) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::pos", f)
	}
}

func (ptr *QUdpSocket) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::pos")
	}
}

func (ptr *QUdpSocket) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQUdpSocket_Reset
func callbackQUdpSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QUdpSocket) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::reset", f)
	}
}

func (ptr *QUdpSocket) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::reset")
	}
}

func (ptr *QUdpSocket) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQUdpSocket_Seek
func callbackQUdpSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QUdpSocket) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::seek", f)
	}
}

func (ptr *QUdpSocket) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::seek")
	}
}

func (ptr *QUdpSocket) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QUdpSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQUdpSocket_Size
func callbackQUdpSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QUdpSocket) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::size", f)
	}
}

func (ptr *QUdpSocket) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::size")
	}
}

func (ptr *QUdpSocket) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQUdpSocket_TimerEvent
func callbackQUdpSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::timerEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::timerEvent")
	}
}

func (ptr *QUdpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUdpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQUdpSocket_ChildEvent
func callbackQUdpSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::childEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::childEvent")
	}
}

func (ptr *QUdpSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUdpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQUdpSocket_ConnectNotify
func callbackQUdpSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUdpSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUdpSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectNotify", f)
	}
}

func (ptr *QUdpSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectNotify")
	}
}

func (ptr *QUdpSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUdpSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUdpSocket_CustomEvent
func callbackQUdpSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::customEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::customEvent")
	}
}

func (ptr *QUdpSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUdpSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQUdpSocket_DeleteLater
func callbackQUdpSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QUdpSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::deleteLater", f)
	}
}

func (ptr *QUdpSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::deleteLater")
	}
}

func (ptr *QUdpSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQUdpSocket_DisconnectNotify
func callbackQUdpSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUdpSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUdpSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectNotify", f)
	}
}

func (ptr *QUdpSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectNotify")
	}
}

func (ptr *QUdpSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUdpSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUdpSocket_Event
func callbackQUdpSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QUdpSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::event", f)
	}
}

func (ptr *QUdpSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::event")
	}
}

func (ptr *QUdpSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QUdpSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQUdpSocket_EventFilter
func callbackQUdpSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QUdpSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::eventFilter", f)
	}
}

func (ptr *QUdpSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::eventFilter")
	}
}

func (ptr *QUdpSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QUdpSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQUdpSocket_MetaObject
func callbackQUdpSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQUdpSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QUdpSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::metaObject", f)
	}
}

func (ptr *QUdpSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::metaObject")
	}
}

func (ptr *QUdpSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUdpSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUdpSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUdpSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
