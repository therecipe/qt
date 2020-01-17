// +build !minimal

package network

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"reflect"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtNetwork_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtNetwork_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
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

func NewQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) (n *QAbstractNetworkCache) {
	n = new(QAbstractNetworkCache)
	n.SetPointer(ptr)
	return
}
func NewQAbstractNetworkCache(parent core.QObject_ITF) *QAbstractNetworkCache {
	tmpValue := NewQAbstractNetworkCacheFromPointer(C.QAbstractNetworkCache_NewQAbstractNetworkCache(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractNetworkCache_CacheSize
func callbackQAbstractNetworkCache_CacheSize(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "cacheSize"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QAbstractNetworkCache) ConnectCacheSize(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cacheSize"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cacheSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cacheSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCacheSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cacheSize")
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
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractNetworkCache) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QAbstractNetworkCache) Clear() {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(ptr.Pointer())
	}
}

//export callbackQAbstractNetworkCache_Data
func callbackQAbstractNetworkCache_Data(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQIODevice((*(*func(*core.QUrl) *core.QIODevice)(signal))(core.NewQUrlFromPointer(url)))
	}

	return core.PointerFromQIODevice(core.NewQIODevice())
}

func (ptr *QAbstractNetworkCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			f := func(url *core.QUrl) *core.QIODevice {
				(*(*func(*core.QUrl) *core.QIODevice)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Insert
func callbackQAbstractNetworkCache_Insert(ptr unsafe.Pointer, device unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "insert"); signal != nil {
		(*(*func(*core.QIODevice))(signal))(core.NewQIODeviceFromPointer(device))
	}

}

func (ptr *QAbstractNetworkCache) ConnectInsert(f func(device *core.QIODevice)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insert"); signal != nil {
			f := func(device *core.QIODevice) {
				(*(*func(*core.QIODevice))(signal))(device)
				f(device)
			}
			qt.ConnectSignal(ptr.Pointer(), "insert", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insert", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectInsert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insert")
	}
}

func (ptr *QAbstractNetworkCache) Insert(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Insert(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

//export callbackQAbstractNetworkCache_MetaData
func callbackQAbstractNetworkCache_MetaData(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaData"); signal != nil {
		return PointerFromQNetworkCacheMetaData((*(*func(*core.QUrl) *QNetworkCacheMetaData)(signal))(core.NewQUrlFromPointer(url)))
	}

	return PointerFromQNetworkCacheMetaData(NewQNetworkCacheMetaData())
}

func (ptr *QAbstractNetworkCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metaData"); signal != nil {
			f := func(url *core.QUrl) *QNetworkCacheMetaData {
				(*(*func(*core.QUrl) *QNetworkCacheMetaData)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "metaData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metaData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metaData")
	}
}

func (ptr *QAbstractNetworkCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCacheMetaDataFromPointer(C.QAbstractNetworkCache_MetaData(ptr.Pointer(), core.PointerFromQUrl(url)))
		qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Prepare
func callbackQAbstractNetworkCache_Prepare(ptr unsafe.Pointer, metaData unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "prepare"); signal != nil {
		return core.PointerFromQIODevice((*(*func(*QNetworkCacheMetaData) *core.QIODevice)(signal))(NewQNetworkCacheMetaDataFromPointer(metaData)))
	}

	return core.PointerFromQIODevice(core.NewQIODevice())
}

func (ptr *QAbstractNetworkCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "prepare"); signal != nil {
			f := func(metaData *QNetworkCacheMetaData) *core.QIODevice {
				(*(*func(*QNetworkCacheMetaData) *core.QIODevice)(signal))(metaData)
				return f(metaData)
			}
			qt.ConnectSignal(ptr.Pointer(), "prepare", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "prepare", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectPrepare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "prepare")
	}
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Remove
func callbackQAbstractNetworkCache_Remove(ptr unsafe.Pointer, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl) bool)(signal))(core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAbstractNetworkCache) ConnectRemove(f func(url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			f := func(url *core.QUrl) bool {
				(*(*func(*core.QUrl) bool)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "remove", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectRemove() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *QAbstractNetworkCache) Remove(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractNetworkCache_Remove(ptr.Pointer(), core.PointerFromQUrl(url))) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_UpdateMetaData
func callbackQAbstractNetworkCache_UpdateMetaData(ptr unsafe.Pointer, metaData unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMetaData"); signal != nil {
		(*(*func(*QNetworkCacheMetaData))(signal))(NewQNetworkCacheMetaDataFromPointer(metaData))
	}

}

func (ptr *QAbstractNetworkCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateMetaData"); signal != nil {
			f := func(metaData *QNetworkCacheMetaData) {
				(*(*func(*QNetworkCacheMetaData))(signal))(metaData)
				f(metaData)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateMetaData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateMetaData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectUpdateMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateMetaData")
	}
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

//export callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache
func callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractNetworkCache"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DestroyQAbstractNetworkCacheDefault()
	}
}

func (ptr *QAbstractNetworkCache) ConnectDestroyQAbstractNetworkCache(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractNetworkCache"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractNetworkCache", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractNetworkCache", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDestroyQAbstractNetworkCache() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractNetworkCache")
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCacheDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCacheDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractNetworkCache___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractNetworkCache) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractNetworkCache) __children_newList() unsafe.Pointer {
	return C.QAbstractNetworkCache___children_newList(ptr.Pointer())
}

func (ptr *QAbstractNetworkCache) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractNetworkCache___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractNetworkCache) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractNetworkCache) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractNetworkCache___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractNetworkCache) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractNetworkCache___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractNetworkCache) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractNetworkCache) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractNetworkCache___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractNetworkCache) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractNetworkCache___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractNetworkCache) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractNetworkCache) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractNetworkCache___findChildren_newList3(ptr.Pointer())
}

//export callbackQAbstractNetworkCache_ChildEvent
func callbackQAbstractNetworkCache_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractNetworkCache_ConnectNotify
func callbackQAbstractNetworkCache_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractNetworkCache_CustomEvent
func callbackQAbstractNetworkCache_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractNetworkCache_DeleteLater
func callbackQAbstractNetworkCache_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractNetworkCache) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractNetworkCache_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAbstractNetworkCache_Destroyed
func callbackQAbstractNetworkCache_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractNetworkCache_DisconnectNotify
func callbackQAbstractNetworkCache_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractNetworkCache) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractNetworkCache_Event
func callbackQAbstractNetworkCache_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractNetworkCacheFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractNetworkCache) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractNetworkCache_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_EventFilter
func callbackQAbstractNetworkCache_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractNetworkCacheFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractNetworkCache) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractNetworkCache_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_MetaObject
func callbackQAbstractNetworkCache_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQAbstractNetworkCacheFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractNetworkCache) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractNetworkCache_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractNetworkCache_ObjectNameChanged
func callbackQAbstractNetworkCache_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractNetworkCache_TimerEvent
func callbackQAbstractNetworkCache_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQAbstractSocketFromPointer(ptr unsafe.Pointer) (n *QAbstractSocket) {
	n = new(QAbstractSocket)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractSocket__SocketType
//QAbstractSocket::SocketType
type QAbstractSocket__SocketType int64

const (
	QAbstractSocket__TcpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(0)
	QAbstractSocket__UdpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(1)
	QAbstractSocket__SctpSocket        QAbstractSocket__SocketType = QAbstractSocket__SocketType(2)
	QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = QAbstractSocket__SocketType(-1)
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
	QAbstractSocket__PathMtuSocketOption           QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(7)
)

//go:generate stringer -type=QAbstractSocket__BindFlag
//QAbstractSocket::BindFlag
type QAbstractSocket__BindFlag int64

const (
	QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x0)
	QAbstractSocket__ShareAddress       QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x1)
	QAbstractSocket__DontShareAddress   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x2)
	QAbstractSocket__ReuseAddressHint   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x4)
)

//go:generate stringer -type=QAbstractSocket__PauseMode
//QAbstractSocket::PauseMode
type QAbstractSocket__PauseMode int64

const (
	QAbstractSocket__PauseNever       QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x0)
	QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x1)
)

func NewQAbstractSocket(socketType QAbstractSocket__SocketType, parent core.QObject_ITF) *QAbstractSocket {
	tmpValue := NewQAbstractSocketFromPointer(C.QAbstractSocket_NewQAbstractSocket(C.longlong(socketType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QAbstractSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_AtEndDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Bind(address QHostAddress_ITF, port uint16, mode QAbstractSocket__BindFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_Bind(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(mode))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Bind2(port uint16, mode QAbstractSocket__BindFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_Bind2(ptr.Pointer(), C.ushort(port), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQAbstractSocket_BytesAvailable
func callbackQAbstractSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QAbstractSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_BytesToWrite
func callbackQAbstractSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesToWrite"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QAbstractSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_CanReadLine
func callbackQAbstractSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QAbstractSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_CanReadLineDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractSocket_Close
func callbackQAbstractSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QAbstractSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_CloseDefault(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_ConnectToHost
func callbackQAbstractSocket_ConnectToHost(ptr unsafe.Pointer, hostName C.struct_QtNetwork_PackedString, port C.ushort, openMode C.longlong, protocol C.longlong) {
	if signal := qt.GetSignal(ptr, "connectToHost"); signal != nil {
		(*(*func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(signal))(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectToHostDefault(cGoUnpackString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QAbstractSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectToHost"); signal != nil {
			f := func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
				(*(*func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(signal))(hostName, port, openMode, protocol)
				f(hostName, port, openMode, protocol)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectToHost", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectToHost", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectConnectToHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectToHost")
	}
}

func (ptr *QAbstractSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QAbstractSocket_ConnectToHost(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QAbstractSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QAbstractSocket_ConnectToHostDefault(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQAbstractSocket_ConnectToHost2
func callbackQAbstractSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	if signal := qt.GetSignal(ptr, "connectToHost2"); signal != nil {
		(*(*func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(signal))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QAbstractSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectToHost2"); signal != nil {
			f := func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag) {
				(*(*func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(signal))(address, port, openMode)
				f(address, port, openMode)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectToHost2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectToHost2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectConnectToHost2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectToHost2")
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

//export callbackQAbstractSocket_Connected
func callbackQAbstractSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QAbstractSocket_ConnectConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connected")
	}
}

func (ptr *QAbstractSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Connected(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_DisconnectFromHost
func callbackQAbstractSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectFromHost"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDisconnectFromHost(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "disconnectFromHost"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnectFromHost", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnectFromHost", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnectFromHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "disconnectFromHost")
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
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QAbstractSocket_ConnectDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "disconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "disconnected")
	}
}

func (ptr *QAbstractSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Disconnected(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) Error() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QAbstractSocket_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_Error2
func callbackQAbstractSocket_Error2(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QAbstractSocket__SocketError))(signal))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QAbstractSocket) ConnectError2(f func(socketError QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QAbstractSocket_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(socketError QAbstractSocket__SocketError) {
				(*(*func(QAbstractSocket__SocketError))(signal))(socketError)
				f(socketError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QAbstractSocket) Error2(socketError QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Error2(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QAbstractSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_Flush(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractSocket_HostFound
func callbackQAbstractSocket_HostFound(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hostFound"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSocket) ConnectHostFound(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hostFound") {
			C.QAbstractSocket_ConnectHostFound(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hostFound")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hostFound"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hostFound", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostFound", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectHostFound() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectHostFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hostFound")
	}
}

func (ptr *QAbstractSocket) HostFound() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_HostFound(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_IsSequential
func callbackQAbstractSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QAbstractSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_IsSequentialDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSocket) LocalAddress() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QAbstractSocket_LocalAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
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
		tmpValue := NewQHostAddressFromPointer(C.QAbstractSocket_PeerAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
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

func (ptr *QAbstractSocket) ProtocolTag() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractSocket_ProtocolTag(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSocket) Proxy() *QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkProxyFromPointer(C.QAbstractSocket_Proxy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractSocket_ProxyAuthenticationRequired
func callbackQAbstractSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "proxyAuthenticationRequired"); signal != nil {
		(*(*func(*QNetworkProxy, *QAuthenticator))(signal))(NewQNetworkProxyFromPointer(proxy), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QAbstractSocket) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "proxyAuthenticationRequired") {
			C.QAbstractSocket_ConnectProxyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "proxyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "proxyAuthenticationRequired"); signal != nil {
			f := func(proxy *QNetworkProxy, authenticator *QAuthenticator) {
				(*(*func(*QNetworkProxy, *QAuthenticator))(signal))(proxy, authenticator)
				f(proxy, authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "proxyAuthenticationRequired")
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
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQAbstractSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QAbstractSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			f := func(data *string, maxSize int64) int64 {
				(*(*func(*string, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "readData")
	}
}

func (ptr *QAbstractSocket) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QAbstractSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QAbstractSocket) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QAbstractSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQAbstractSocket_ReadLineData
func callbackQAbstractSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxlen C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readLineData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxlen)))
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackBytes(data), int64(maxlen)))
}

func (ptr *QAbstractSocket) ReadLineDataDefault(data []byte, maxlen int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QAbstractSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

//export callbackQAbstractSocket_Resume
func callbackQAbstractSocket_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resume"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QAbstractSocket) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resume"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "resume", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resume", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resume")
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
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QAbstractSocket_SetPeerName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QAbstractSocket) SetPeerPort(port uint16) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPeerPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QAbstractSocket) SetProtocolTag(tag string) {
	if ptr.Pointer() != nil {
		var tagC *C.char
		if tag != "" {
			tagC = C.CString(tag)
			defer C.free(unsafe.Pointer(tagC))
		}
		C.QAbstractSocket_SetProtocolTag(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: tagC, len: C.longlong(len(tag))})
	}
}

func (ptr *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

//export callbackQAbstractSocket_SetReadBufferSize
func callbackQAbstractSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	if signal := qt.GetSignal(ptr, "setReadBufferSize"); signal != nil {
		(*(*func(int64))(signal))(int64(size))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QAbstractSocket) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setReadBufferSize"); signal != nil {
			f := func(size int64) {
				(*(*func(int64))(signal))(size)
				f(size)
			}
			qt.ConnectSignal(ptr.Pointer(), "setReadBufferSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setReadBufferSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setReadBufferSize")
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
	if signal := qt.GetSignal(ptr, "setSocketOption"); signal != nil {
		(*(*func(QAbstractSocket__SocketOption, *core.QVariant))(signal))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QAbstractSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSocketOption"); signal != nil {
			f := func(option QAbstractSocket__SocketOption, value *core.QVariant) {
				(*(*func(QAbstractSocket__SocketOption, *core.QVariant))(signal))(option, value)
				f(option, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSocketOption", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSocketOption", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectSetSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSocketOption")
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
	if signal := qt.GetSignal(ptr, "socketOption"); signal != nil {
		return core.PointerFromQVariant((*(*func(QAbstractSocket__SocketOption) *core.QVariant)(signal))(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQAbstractSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QAbstractSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "socketOption"); signal != nil {
			f := func(option QAbstractSocket__SocketOption) *core.QVariant {
				(*(*func(QAbstractSocket__SocketOption) *core.QVariant)(signal))(option)
				return f(option)
			}
			qt.ConnectSignal(ptr.Pointer(), "socketOption", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "socketOption", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectSocketOption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "socketOption")
	}
}

func (ptr *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAbstractSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAbstractSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QAbstractSocket__SocketState))(signal))(QAbstractSocket__SocketState(socketState))
	}

}

func (ptr *QAbstractSocket) ConnectStateChanged(f func(socketState QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QAbstractSocket_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(socketState QAbstractSocket__SocketState) {
				(*(*func(QAbstractSocket__SocketState))(signal))(socketState)
				f(socketState)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QAbstractSocket) StateChanged(socketState QAbstractSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_StateChanged(ptr.Pointer(), C.longlong(socketState))
	}
}

//export callbackQAbstractSocket_WaitForBytesWritten
func callbackQAbstractSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForConnected
func callbackQAbstractSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "waitForConnected"); signal != nil {
			f := func(msecs int) bool {
				(*(*func(int) bool)(signal))(msecs)
				return f(msecs)
			}
			qt.ConnectSignal(ptr.Pointer(), "waitForConnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "waitForConnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForConnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "waitForConnected")
	}
}

func (ptr *QAbstractSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForConnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForDisconnected
func callbackQAbstractSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "waitForDisconnected"); signal != nil {
			f := func(msecs int) bool {
				(*(*func(int) bool)(signal))(msecs)
				return f(msecs)
			}
			qt.ConnectSignal(ptr.Pointer(), "waitForDisconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "waitForDisconnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForDisconnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "waitForDisconnected")
	}
}

func (ptr *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForDisconnectedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForReadyRead
func callbackQAbstractSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WriteData
func callbackQAbstractSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, size C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(size)))
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(size)))
}

func (ptr *QAbstractSocket) ConnectWriteData(f func(data []byte, size int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			f := func(data []byte, size int64) int64 {
				(*(*func([]byte, int64) int64)(signal))(data, size)
				return f(data, size)
			}
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QAbstractSocket) WriteData(data []byte, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QAbstractSocket_WriteData(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

func (ptr *QAbstractSocket) WriteDataDefault(data []byte, size int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QAbstractSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

//export callbackQAbstractSocket_DestroyQAbstractSocket
func callbackQAbstractSocket_DestroyQAbstractSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSocketFromPointer(ptr).DestroyQAbstractSocketDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDestroyQAbstractSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractSocket", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSocket) DisconnectDestroyQAbstractSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractSocket")
	}
}

func (ptr *QAbstractSocket) DestroyQAbstractSocket() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSocket_DestroyQAbstractSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) DestroyQAbstractSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSocket_DestroyQAbstractSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractSocket___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractSocket) __children_newList() unsafe.Pointer {
	return C.QAbstractSocket___children_newList(ptr.Pointer())
}

func (ptr *QAbstractSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractSocket___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractSocket) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractSocket___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractSocket) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractSocket___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractSocket) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractSocket___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractSocket) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractSocket___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractSocket) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractSocket___findChildren_newList3(ptr.Pointer())
}

//export callbackQAbstractSocket_AboutToClose
func callbackQAbstractSocket_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQAbstractSocket_BytesWritten
func callbackQAbstractSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(bytes))
	}

}

//export callbackQAbstractSocket_ChannelBytesWritten
func callbackQAbstractSocket_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		(*(*func(int, int64))(signal))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQAbstractSocket_ChannelReadyRead
func callbackQAbstractSocket_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		(*(*func(int))(signal))(int(int32(channel)))
	}

}

//export callbackQAbstractSocket_Open
func callbackQAbstractSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(core.QIODevice__OpenModeFlag) bool)(signal))(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QAbstractSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_OpenDefault(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQAbstractSocket_Pos
func callbackQAbstractSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).PosDefault())
}

func (ptr *QAbstractSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_ReadChannelFinished
func callbackQAbstractSocket_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQAbstractSocket_ReadyRead
func callbackQAbstractSocket_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQAbstractSocket_Reset
func callbackQAbstractSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QAbstractSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_ResetDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractSocket_Seek
func callbackQAbstractSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int64) bool)(signal))(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QAbstractSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_SeekDefault(ptr.Pointer(), C.longlong(pos))) != 0
	}
	return false
}

//export callbackQAbstractSocket_Size
func callbackQAbstractSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QAbstractSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_ChildEvent
func callbackQAbstractSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractSocket_ConnectNotify
func callbackQAbstractSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSocket_CustomEvent
func callbackQAbstractSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractSocket_DeleteLater
func callbackQAbstractSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSocket_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_Destroyed
func callbackQAbstractSocket_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractSocket_DisconnectNotify
func callbackQAbstractSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSocket_Event
func callbackQAbstractSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractSocket_EventFilter
func callbackQAbstractSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractSocket_MetaObject
func callbackQAbstractSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQAbstractSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractSocket_ObjectNameChanged
func callbackQAbstractSocket_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractSocket_TimerEvent
func callbackQAbstractSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) (n *QAuthenticator) {
	n = new(QAuthenticator)
	n.SetPointer(ptr)
	return
}
func NewQAuthenticator() *QAuthenticator {
	tmpValue := NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator())
	qt.SetFinalizer(tmpValue, (*QAuthenticator).DestroyQAuthenticator)
	return tmpValue
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {
	tmpValue := NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator2(PointerFromQAuthenticator(other)))
	qt.SetFinalizer(tmpValue, (*QAuthenticator).DestroyQAuthenticator)
	return tmpValue
}

func (ptr *QAuthenticator) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAuthenticator_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {
	if ptr.Pointer() != nil {
		var optC *C.char
		if opt != "" {
			optC = C.CString(opt)
			defer C.free(unsafe.Pointer(optC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QAuthenticator_Option(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: optC, len: C.longlong(len(opt))}))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAuthenticator) Options() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQAuthenticatorFromPointer(l.data)
			for i, v := range tmpList.__options_keyList() {
				out[v] = tmpList.__options_atList(v, i)
			}
			return out
		}(C.QAuthenticator_Options(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
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
		var optC *C.char
		if opt != "" {
			optC = C.CString(opt)
			defer C.free(unsafe.Pointer(optC))
		}
		C.QAuthenticator_SetOption(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: optC, len: C.longlong(len(opt))}, core.PointerFromQVariant(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.QAuthenticator_SetPassword(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	if ptr.Pointer() != nil {
		var userC *C.char
		if user != "" {
			userC = C.CString(user)
			defer C.free(unsafe.Pointer(userC))
		}
		C.QAuthenticator_SetUser(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: userC, len: C.longlong(len(user))})
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

		qt.SetFinalizer(ptr, nil)
		C.QAuthenticator_DestroyQAuthenticator(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAuthenticator) __options_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QAuthenticator___options_atList(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAuthenticator) __options_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QAuthenticator___options_setList(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QAuthenticator) __options_newList() unsafe.Pointer {
	return C.QAuthenticator___options_newList(ptr.Pointer())
}

func (ptr *QAuthenticator) __options_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQAuthenticatorFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____options_keyList_atList(i)
			}
			return out
		}(C.QAuthenticator___options_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QAuthenticator) ____options_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAuthenticator_____options_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QAuthenticator) ____options_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QAuthenticator_____options_keyList_setList(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QAuthenticator) ____options_keyList_newList() unsafe.Pointer {
	return C.QAuthenticator_____options_keyList_newList(ptr.Pointer())
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

func NewQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) (n *QDnsDomainNameRecord) {
	n = new(QDnsDomainNameRecord)
	n.SetPointer(ptr)
	return
}
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	tmpValue := NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord())
	qt.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
	return tmpValue
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecord_ITF) *QDnsDomainNameRecord {
	tmpValue := NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord2(PointerFromQDnsDomainNameRecord(other)))
	qt.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
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

		qt.SetFinalizer(ptr, nil)
		C.QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) (n *QDnsHostAddressRecord) {
	n = new(QDnsHostAddressRecord)
	n.SetPointer(ptr)
	return
}
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	tmpValue := NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord())
	qt.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
	return tmpValue
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecord_ITF) *QDnsHostAddressRecord {
	tmpValue := NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord2(PointerFromQDnsHostAddressRecord(other)))
	qt.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
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
		tmpValue := NewQHostAddressFromPointer(C.QDnsHostAddressRecord_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) (n *QDnsLookup) {
	n = new(QDnsLookup)
	n.SetPointer(ptr)
	return
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

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {
	tmpValue := NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup2(C.longlong(ty), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup3(C.longlong(ty), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQHostAddress(nameserver), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDnsLookup_Abort
func callbackQDnsLookup_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "abort"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDnsLookupFromPointer(ptr).AbortDefault()
	}
}

func (ptr *QDnsLookup) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "abort"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "abort")
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
			out := make([]*QDnsDomainNameRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__canonicalNameRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_CanonicalNameRecords(ptr.Pointer()))
	}
	return make([]*QDnsDomainNameRecord, 0)
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

//export callbackQDnsLookup_Finished
func callbackQDnsLookup_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QDnsLookup_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
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
			out := make([]*QDnsHostAddressRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__hostAddressRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_HostAddressRecords(ptr.Pointer()))
	}
	return make([]*QDnsHostAddressRecord, 0)
}

func (ptr *QDnsLookup) IsFinished() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDnsLookup_IsFinished(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQDnsLookup_Lookup
func callbackQDnsLookup_Lookup(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lookup"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDnsLookupFromPointer(ptr).LookupDefault()
	}
}

func (ptr *QDnsLookup) ConnectLookup(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lookup"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lookup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lookup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectLookup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lookup")
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
			out := make([]*QDnsMailExchangeRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mailExchangeRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_MailExchangeRecords(ptr.Pointer()))
	}
	return make([]*QDnsMailExchangeRecord, 0)
}

func (ptr *QDnsLookup) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDnsLookup_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQDnsLookup_NameChanged
func callbackQDnsLookup_NameChanged(ptr unsafe.Pointer, name C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "nameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(name))
	}

}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameChanged") {
			C.QDnsLookup_ConnectNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nameChanged"); signal != nil {
			f := func(name string) {
				(*(*func(string))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nameChanged")
	}
}

func (ptr *QDnsLookup) NameChanged(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDnsLookup_NameChanged(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QDnsLookup) NameServerRecords() []*QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsDomainNameRecord {
			out := make([]*QDnsDomainNameRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__nameServerRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_NameServerRecords(ptr.Pointer()))
	}
	return make([]*QDnsDomainNameRecord, 0)
}

func (ptr *QDnsLookup) Nameserver() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QDnsLookup_Nameserver(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

//export callbackQDnsLookup_NameserverChanged
func callbackQDnsLookup_NameserverChanged(ptr unsafe.Pointer, nameserver unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "nameserverChanged"); signal != nil {
		(*(*func(*QHostAddress))(signal))(NewQHostAddressFromPointer(nameserver))
	}

}

func (ptr *QDnsLookup) ConnectNameserverChanged(f func(nameserver *QHostAddress)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameserverChanged") {
			C.QDnsLookup_ConnectNameserverChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nameserverChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nameserverChanged"); signal != nil {
			f := func(nameserver *QHostAddress) {
				(*(*func(*QHostAddress))(signal))(nameserver)
				f(nameserver)
			}
			qt.ConnectSignal(ptr.Pointer(), "nameserverChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nameserverChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectNameserverChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameserverChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nameserverChanged")
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
			out := make([]*QDnsDomainNameRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__pointerRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_PointerRecords(ptr.Pointer()))
	}
	return make([]*QDnsDomainNameRecord, 0)
}

func (ptr *QDnsLookup) ServiceRecords() []*QDnsServiceRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsServiceRecord {
			out := make([]*QDnsServiceRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__serviceRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_ServiceRecords(ptr.Pointer()))
	}
	return make([]*QDnsServiceRecord, 0)
}

func (ptr *QDnsLookup) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDnsLookup_SetName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})
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

func (ptr *QDnsLookup) TextRecords() []*QDnsTextRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QDnsTextRecord {
			out := make([]*QDnsTextRecord, int(l.len))
			tmpList := NewQDnsLookupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__textRecords_atList(i)
			}
			return out
		}(C.QDnsLookup_TextRecords(ptr.Pointer()))
	}
	return make([]*QDnsTextRecord, 0)
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQDnsLookup_TypeChanged
func callbackQDnsLookup_TypeChanged(ptr unsafe.Pointer, ty C.longlong) {
	if signal := qt.GetSignal(ptr, "typeChanged"); signal != nil {
		(*(*func(QDnsLookup__Type))(signal))(QDnsLookup__Type(ty))
	}

}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeChanged") {
			C.QDnsLookup_ConnectTypeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "typeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "typeChanged"); signal != nil {
			f := func(ty QDnsLookup__Type) {
				(*(*func(QDnsLookup__Type))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeChanged")
	}
}

func (ptr *QDnsLookup) TypeChanged(ty QDnsLookup__Type) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_TypeChanged(ptr.Pointer(), C.longlong(ty))
	}
}

//export callbackQDnsLookup_DestroyQDnsLookup
func callbackQDnsLookup_DestroyQDnsLookup(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDnsLookup"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDnsLookupFromPointer(ptr).DestroyQDnsLookupDefault()
	}
}

func (ptr *QDnsLookup) ConnectDestroyQDnsLookup(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDnsLookup"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QDnsLookup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDnsLookup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDnsLookup) DisconnectDestroyQDnsLookup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDnsLookup")
	}
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDnsLookup_DestroyQDnsLookup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) DestroyQDnsLookupDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDnsLookup_DestroyQDnsLookupDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) __canonicalNameRecords_atList(i int) *QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsDomainNameRecordFromPointer(C.QDnsLookup___canonicalNameRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __canonicalNameRecords_setList(i QDnsDomainNameRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___canonicalNameRecords_setList(ptr.Pointer(), PointerFromQDnsDomainNameRecord(i))
	}
}

func (ptr *QDnsLookup) __canonicalNameRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___canonicalNameRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __hostAddressRecords_atList(i int) *QDnsHostAddressRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsHostAddressRecordFromPointer(C.QDnsLookup___hostAddressRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __hostAddressRecords_setList(i QDnsHostAddressRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___hostAddressRecords_setList(ptr.Pointer(), PointerFromQDnsHostAddressRecord(i))
	}
}

func (ptr *QDnsLookup) __hostAddressRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___hostAddressRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __mailExchangeRecords_atList(i int) *QDnsMailExchangeRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsMailExchangeRecordFromPointer(C.QDnsLookup___mailExchangeRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __mailExchangeRecords_setList(i QDnsMailExchangeRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___mailExchangeRecords_setList(ptr.Pointer(), PointerFromQDnsMailExchangeRecord(i))
	}
}

func (ptr *QDnsLookup) __mailExchangeRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___mailExchangeRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __nameServerRecords_atList(i int) *QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsDomainNameRecordFromPointer(C.QDnsLookup___nameServerRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __nameServerRecords_setList(i QDnsDomainNameRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___nameServerRecords_setList(ptr.Pointer(), PointerFromQDnsDomainNameRecord(i))
	}
}

func (ptr *QDnsLookup) __nameServerRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___nameServerRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __pointerRecords_atList(i int) *QDnsDomainNameRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsDomainNameRecordFromPointer(C.QDnsLookup___pointerRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __pointerRecords_setList(i QDnsDomainNameRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___pointerRecords_setList(ptr.Pointer(), PointerFromQDnsDomainNameRecord(i))
	}
}

func (ptr *QDnsLookup) __pointerRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___pointerRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __serviceRecords_atList(i int) *QDnsServiceRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsServiceRecordFromPointer(C.QDnsLookup___serviceRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __serviceRecords_setList(i QDnsServiceRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___serviceRecords_setList(ptr.Pointer(), PointerFromQDnsServiceRecord(i))
	}
}

func (ptr *QDnsLookup) __serviceRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___serviceRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __textRecords_atList(i int) *QDnsTextRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQDnsTextRecordFromPointer(C.QDnsLookup___textRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __textRecords_setList(i QDnsTextRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___textRecords_setList(ptr.Pointer(), PointerFromQDnsTextRecord(i))
	}
}

func (ptr *QDnsLookup) __textRecords_newList() unsafe.Pointer {
	return C.QDnsLookup___textRecords_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDnsLookup___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDnsLookup) __children_newList() unsafe.Pointer {
	return C.QDnsLookup___children_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDnsLookup___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDnsLookup) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDnsLookup___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDnsLookup___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDnsLookup) __findChildren_newList() unsafe.Pointer {
	return C.QDnsLookup___findChildren_newList(ptr.Pointer())
}

func (ptr *QDnsLookup) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDnsLookup___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDnsLookup) __findChildren_newList3() unsafe.Pointer {
	return C.QDnsLookup___findChildren_newList3(ptr.Pointer())
}

//export callbackQDnsLookup_ChildEvent
func callbackQDnsLookup_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDnsLookup_ConnectNotify
func callbackQDnsLookup_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDnsLookupFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDnsLookup) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDnsLookup_CustomEvent
func callbackQDnsLookup_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDnsLookup_DeleteLater
func callbackQDnsLookup_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDnsLookupFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDnsLookup) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDnsLookup_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQDnsLookup_Destroyed
func callbackQDnsLookup_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDnsLookup_DisconnectNotify
func callbackQDnsLookup_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDnsLookupFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDnsLookup) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDnsLookup_Event
func callbackQDnsLookup_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDnsLookupFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDnsLookup) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDnsLookup_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDnsLookup_EventFilter
func callbackQDnsLookup_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDnsLookupFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDnsLookup) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDnsLookup_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDnsLookup_MetaObject
func callbackQDnsLookup_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQDnsLookupFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDnsLookup) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDnsLookup_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQDnsLookup_ObjectNameChanged
func callbackQDnsLookup_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQDnsLookup_TimerEvent
func callbackQDnsLookup_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) (n *QDnsMailExchangeRecord) {
	n = new(QDnsMailExchangeRecord)
	n.SetPointer(ptr)
	return
}
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	tmpValue := NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord())
	qt.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
	return tmpValue
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecord_ITF) *QDnsMailExchangeRecord {
	tmpValue := NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(PointerFromQDnsMailExchangeRecord(other)))
	qt.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
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

		qt.SetFinalizer(ptr, nil)
		C.QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQDnsServiceRecordFromPointer(ptr unsafe.Pointer) (n *QDnsServiceRecord) {
	n = new(QDnsServiceRecord)
	n.SetPointer(ptr)
	return
}
func NewQDnsServiceRecord() *QDnsServiceRecord {
	tmpValue := NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord())
	qt.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
	return tmpValue
}

func NewQDnsServiceRecord2(other QDnsServiceRecord_ITF) *QDnsServiceRecord {
	tmpValue := NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord2(PointerFromQDnsServiceRecord(other)))
	qt.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
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

		qt.SetFinalizer(ptr, nil)
		C.QDnsServiceRecord_DestroyQDnsServiceRecord(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQDnsTextRecordFromPointer(ptr unsafe.Pointer) (n *QDnsTextRecord) {
	n = new(QDnsTextRecord)
	n.SetPointer(ptr)
	return
}
func NewQDnsTextRecord() *QDnsTextRecord {
	tmpValue := NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord())
	qt.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
	return tmpValue
}

func NewQDnsTextRecord2(other QDnsTextRecord_ITF) *QDnsTextRecord {
	tmpValue := NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord2(PointerFromQDnsTextRecord(other)))
	qt.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
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
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQDnsTextRecordFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__values_atList(i)
			}
			return out
		}(C.QDnsTextRecord_Values(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDnsTextRecord_DestroyQDnsTextRecord(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsTextRecord) __values_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDnsTextRecord___values_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsTextRecord) __values_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsTextRecord___values_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDnsTextRecord) __values_newList() unsafe.Pointer {
	return C.QDnsTextRecord___values_newList(ptr.Pointer())
}

type QDtls struct {
	core.QObject
}

type QDtls_ITF interface {
	core.QObject_ITF
	QDtls_PTR() *QDtls
}

func (ptr *QDtls) QDtls_PTR() *QDtls {
	return ptr
}

func (ptr *QDtls) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDtls) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDtls(ptr QDtls_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDtls_PTR().Pointer()
	}
	return nil
}

func NewQDtlsFromPointer(ptr unsafe.Pointer) (n *QDtls) {
	n = new(QDtls)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDtls__QDtlsError
//QDtls::QDtlsError
type QDtls__QDtlsError int64

const (
	QDtls__NoError                     QDtls__QDtlsError = QDtls__QDtlsError(0)
	QDtls__InvalidInputParameters      QDtls__QDtlsError = QDtls__QDtlsError(1)
	QDtls__InvalidOperation            QDtls__QDtlsError = QDtls__QDtlsError(2)
	QDtls__UnderlyingSocketError       QDtls__QDtlsError = QDtls__QDtlsError(3)
	QDtls__RemoteClosedConnectionError QDtls__QDtlsError = QDtls__QDtlsError(4)
	QDtls__PeerVerificationError       QDtls__QDtlsError = QDtls__QDtlsError(5)
	QDtls__TlsInitializationError      QDtls__QDtlsError = QDtls__QDtlsError(6)
	QDtls__TlsFatalError               QDtls__QDtlsError = QDtls__QDtlsError(7)
	QDtls__TlsNonFatalError            QDtls__QDtlsError = QDtls__QDtlsError(8)
)

//go:generate stringer -type=QDtls__HandshakeState
//QDtls::HandshakeState
type QDtls__HandshakeState int64

const (
	QDtls__HandshakeNotStarted    QDtls__HandshakeState = QDtls__HandshakeState(0)
	QDtls__HandshakeInProgress    QDtls__HandshakeState = QDtls__HandshakeState(1)
	QDtls__PeerVerificationFailed QDtls__HandshakeState = QDtls__HandshakeState(2)
	QDtls__HandshakeComplete      QDtls__HandshakeState = QDtls__HandshakeState(3)
)

type QDtlsClientVerifier struct {
	core.QObject
}

type QDtlsClientVerifier_ITF interface {
	core.QObject_ITF
	QDtlsClientVerifier_PTR() *QDtlsClientVerifier
}

func (ptr *QDtlsClientVerifier) QDtlsClientVerifier_PTR() *QDtlsClientVerifier {
	return ptr
}

func (ptr *QDtlsClientVerifier) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDtlsClientVerifier) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDtlsClientVerifier(ptr QDtlsClientVerifier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDtlsClientVerifier_PTR().Pointer()
	}
	return nil
}

func NewQDtlsClientVerifierFromPointer(ptr unsafe.Pointer) (n *QDtlsClientVerifier) {
	n = new(QDtlsClientVerifier)
	n.SetPointer(ptr)
	return
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

func NewQHostAddressFromPointer(ptr unsafe.Pointer) (n *QHostAddress) {
	n = new(QHostAddress)
	n.SetPointer(ptr)
	return
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

//go:generate stringer -type=QHostAddress__ConversionModeFlag
//QHostAddress::ConversionModeFlag
type QHostAddress__ConversionModeFlag int64

const (
	QHostAddress__ConvertV4MappedToIPv4     QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(1)
	QHostAddress__ConvertV4CompatToIPv4     QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(2)
	QHostAddress__ConvertUnspecifiedAddress QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(4)
	QHostAddress__ConvertLocalHost          QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(8)
	QHostAddress__TolerantConversion        QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(0xff)
	QHostAddress__StrictConversion          QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(0)
)

func NewQHostAddress() *QHostAddress {
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress())
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress2(ip4Addr uint) *QHostAddress {
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress2(C.uint(uint32(ip4Addr))))
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress3(ip6Addr string) *QHostAddress {
	var ip6AddrC *C.char
	if ip6Addr != "" {
		ip6AddrC = C.CString(ip6Addr)
		defer C.free(unsafe.Pointer(ip6AddrC))
	}
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress3(ip6AddrC))
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress4(ip6Addr string) *QHostAddress {
	var ip6AddrC *C.char
	if ip6Addr != "" {
		ip6AddrC = C.CString(ip6Addr)
		defer C.free(unsafe.Pointer(ip6AddrC))
	}
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress4(ip6AddrC))
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress7(address string) *QHostAddress {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress7(C.struct_QtNetwork_PackedString{data: addressC, len: C.longlong(len(address))}))
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress8(address QHostAddress_ITF) *QHostAddress {
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress8(PointerFromQHostAddress(address)))
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress9(address QHostAddress__SpecialAddress) *QHostAddress {
	tmpValue := NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress9(C.longlong(address)))
	qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func (ptr *QHostAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QHostAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QHostAddress) IsBroadcast() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsBroadcast(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsEqual(other QHostAddress_ITF, mode QHostAddress__ConversionModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsEqual(ptr.Pointer(), PointerFromQHostAddress(other), C.longlong(mode))) != 0
	}
	return false
}

func (ptr *QHostAddress) IsGlobal() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsGlobal(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsInSubnet(ptr.Pointer(), PointerFromQHostAddress(subnet), C.int(int32(netmask)))) != 0
	}
	return false
}

func (ptr *QHostAddress) IsLinkLocal() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsLinkLocal(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsLoopback() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsLoopback(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsMulticast() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsMulticast(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsSiteLocal() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsSiteLocal(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHostAddress) IsUniqueLocalUnicast() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHostAddress_IsUniqueLocalUnicast(ptr.Pointer())) != 0
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

func (ptr *QHostAddress) SetAddress(ip4Addr uint) {
	if ptr.Pointer() != nil {
		C.QHostAddress_SetAddress(ptr.Pointer(), C.uint(uint32(ip4Addr)))
	}
}

func (ptr *QHostAddress) SetAddress2(ip6Addr string) {
	if ptr.Pointer() != nil {
		var ip6AddrC *C.char
		if ip6Addr != "" {
			ip6AddrC = C.CString(ip6Addr)
			defer C.free(unsafe.Pointer(ip6AddrC))
		}
		C.QHostAddress_SetAddress2(ptr.Pointer(), ip6AddrC)
	}
}

func (ptr *QHostAddress) SetAddress3(ip6Addr string) {
	if ptr.Pointer() != nil {
		var ip6AddrC *C.char
		if ip6Addr != "" {
			ip6AddrC = C.CString(ip6Addr)
			defer C.free(unsafe.Pointer(ip6AddrC))
		}
		C.QHostAddress_SetAddress3(ptr.Pointer(), ip6AddrC)
	}
}

func (ptr *QHostAddress) SetAddress6(address string) bool {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		return int8(C.QHostAddress_SetAddress6(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: addressC, len: C.longlong(len(address))})) != 0
	}
	return false
}

func (ptr *QHostAddress) SetAddress7(address QHostAddress__SpecialAddress) {
	if ptr.Pointer() != nil {
		C.QHostAddress_SetAddress7(ptr.Pointer(), C.longlong(address))
	}
}

func (ptr *QHostAddress) SetScopeId(id string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.QHostAddress_SetScopeId(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: idC, len: C.longlong(len(id))})
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

func (ptr *QHostAddress) ToIPv4Address2(ok *bool) uint {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return uint(uint32(C.QHostAddress_ToIPv4Address2(ptr.Pointer(), &okC)))
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

		qt.SetFinalizer(ptr, nil)
		C.QHostAddress_DestroyQHostAddress(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQHostInfoFromPointer(ptr unsafe.Pointer) (n *QHostInfo) {
	n = new(QHostInfo)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QHostInfo__HostInfoError
//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      QHostInfo__HostInfoError = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound QHostInfo__HostInfoError = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError QHostInfo__HostInfoError = QHostInfo__HostInfoError(2)
)

func NewQHostInfo(id int) *QHostInfo {
	tmpValue := NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo(C.int(int32(id))))
	qt.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func NewQHostInfo2(other QHostInfo_ITF) *QHostInfo {
	tmpValue := NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo2(PointerFromQHostInfo(other)))
	qt.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
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
			out := make([]*QHostAddress, int(l.len))
			tmpList := NewQHostInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addresses_atList(i)
			}
			return out
		}(C.QHostInfo_Addresses(ptr.Pointer()))
	}
	return make([]*QHostAddress, 0)
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
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQHostInfoFromPointer(C.QHostInfo_QHostInfo_FromName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func (ptr *QHostInfo) FromName(name string) *QHostInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQHostInfoFromPointer(C.QHostInfo_QHostInfo_FromName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func (ptr *QHostInfo) HostName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHostInfo_HostName(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_LocalDomainName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalDomainName())
}

func (ptr *QHostInfo) LocalDomainName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalDomainName())
}

func QHostInfo_LocalHostName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalHostName())
}

func (ptr *QHostInfo) LocalHostName() string {
	return cGoUnpackString(C.QHostInfo_QHostInfo_LocalHostName())
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var memberC *C.char
	if member != "" {
		memberC = C.CString(member)
		defer C.free(unsafe.Pointer(memberC))
	}
	return int(int32(C.QHostInfo_QHostInfo_LookupHost(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(receiver), memberC)))
}

func (ptr *QHostInfo) LookupHost(name string, receiver core.QObject_ITF, member string) int {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var memberC *C.char
	if member != "" {
		memberC = C.CString(member)
		defer C.free(unsafe.Pointer(memberC))
	}
	return int(int32(C.QHostInfo_QHostInfo_LookupHost(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(receiver), memberC)))
}

func (ptr *QHostInfo) LookupId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHostInfo_LookupId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostInfo) SetAddresses(addresses []*QHostAddress) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetAddresses(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQHostInfoFromPointer(NewQHostInfoFromPointer(nil).__setAddresses_addresses_newList())
			for _, v := range addresses {
				tmpList.__setAddresses_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		C.QHostInfo_SetErrorString(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: strC, len: C.longlong(len(str))})
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QHostInfo_SetHostName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))})
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QHostInfo) Swap(other QHostInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QHostInfo_Swap(ptr.Pointer(), PointerFromQHostInfo(other))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHostInfo_DestroyQHostInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHostInfo) __addresses_atList(i int) *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QHostInfo___addresses_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QHostInfo) __addresses_setList(i QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QHostInfo___addresses_setList(ptr.Pointer(), PointerFromQHostAddress(i))
	}
}

func (ptr *QHostInfo) __addresses_newList() unsafe.Pointer {
	return C.QHostInfo___addresses_newList(ptr.Pointer())
}

func (ptr *QHostInfo) __setAddresses_addresses_atList(i int) *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QHostInfo___setAddresses_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QHostInfo) __setAddresses_addresses_setList(i QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QHostInfo___setAddresses_addresses_setList(ptr.Pointer(), PointerFromQHostAddress(i))
	}
}

func (ptr *QHostInfo) __setAddresses_addresses_newList() unsafe.Pointer {
	return C.QHostInfo___setAddresses_addresses_newList(ptr.Pointer())
}

type QHstsPolicy struct {
	ptr unsafe.Pointer
}

type QHstsPolicy_ITF interface {
	QHstsPolicy_PTR() *QHstsPolicy
}

func (ptr *QHstsPolicy) QHstsPolicy_PTR() *QHstsPolicy {
	return ptr
}

func (ptr *QHstsPolicy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHstsPolicy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHstsPolicy(ptr QHstsPolicy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHstsPolicy_PTR().Pointer()
	}
	return nil
}

func NewQHstsPolicyFromPointer(ptr unsafe.Pointer) (n *QHstsPolicy) {
	n = new(QHstsPolicy)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QHstsPolicy__PolicyFlag
//QHstsPolicy::PolicyFlag
type QHstsPolicy__PolicyFlag int64

const (
	QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = QHstsPolicy__PolicyFlag(1)
)

func NewQHstsPolicy() *QHstsPolicy {
	tmpValue := NewQHstsPolicyFromPointer(C.QHstsPolicy_NewQHstsPolicy())
	qt.SetFinalizer(tmpValue, (*QHstsPolicy).DestroyQHstsPolicy)
	return tmpValue
}

func NewQHstsPolicy2(expiry core.QDateTime_ITF, flags QHstsPolicy__PolicyFlag, host string, mode core.QUrl__ParsingMode) *QHstsPolicy {
	var hostC *C.char
	if host != "" {
		hostC = C.CString(host)
		defer C.free(unsafe.Pointer(hostC))
	}
	tmpValue := NewQHstsPolicyFromPointer(C.QHstsPolicy_NewQHstsPolicy2(core.PointerFromQDateTime(expiry), C.longlong(flags), C.struct_QtNetwork_PackedString{data: hostC, len: C.longlong(len(host))}, C.longlong(mode)))
	qt.SetFinalizer(tmpValue, (*QHstsPolicy).DestroyQHstsPolicy)
	return tmpValue
}

func NewQHstsPolicy3(other QHstsPolicy_ITF) *QHstsPolicy {
	tmpValue := NewQHstsPolicyFromPointer(C.QHstsPolicy_NewQHstsPolicy3(PointerFromQHstsPolicy(other)))
	qt.SetFinalizer(tmpValue, (*QHstsPolicy).DestroyQHstsPolicy)
	return tmpValue
}

func (ptr *QHstsPolicy) Expiry() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QHstsPolicy_Expiry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QHstsPolicy) Host(options core.QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHstsPolicy_Host(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QHstsPolicy) IncludesSubDomains() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHstsPolicy_IncludesSubDomains(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHstsPolicy) IsExpired() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHstsPolicy_IsExpired(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHstsPolicy) SetExpiry(expiry core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QHstsPolicy_SetExpiry(ptr.Pointer(), core.PointerFromQDateTime(expiry))
	}
}

func (ptr *QHstsPolicy) SetHost(host string, mode core.QUrl__ParsingMode) {
	if ptr.Pointer() != nil {
		var hostC *C.char
		if host != "" {
			hostC = C.CString(host)
			defer C.free(unsafe.Pointer(hostC))
		}
		C.QHstsPolicy_SetHost(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostC, len: C.longlong(len(host))}, C.longlong(mode))
	}
}

func (ptr *QHstsPolicy) SetIncludesSubDomains(include bool) {
	if ptr.Pointer() != nil {
		C.QHstsPolicy_SetIncludesSubDomains(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(include))))
	}
}

func (ptr *QHstsPolicy) Swap(other QHstsPolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QHstsPolicy_Swap(ptr.Pointer(), PointerFromQHstsPolicy(other))
	}
}

func (ptr *QHstsPolicy) DestroyQHstsPolicy() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHstsPolicy_DestroyQHstsPolicy(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) (n *QHttpMultiPart) {
	n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	return
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

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	tmpValue := NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	tmpValue := NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.longlong(contentType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
		tmpValue := core.NewQByteArrayFromPointer(C.QHttpMultiPart_Boundary(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQHttpMultiPart_DestroyQHttpMultiPart
func callbackQHttpMultiPart_DestroyQHttpMultiPart(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHttpMultiPart"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHttpMultiPartFromPointer(ptr).DestroyQHttpMultiPartDefault()
	}
}

func (ptr *QHttpMultiPart) ConnectDestroyQHttpMultiPart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHttpMultiPart"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHttpMultiPart", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHttpMultiPart", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHttpMultiPart) DisconnectDestroyQHttpMultiPart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHttpMultiPart")
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPartDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHttpMultiPart_DestroyQHttpMultiPartDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHttpMultiPart) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHttpMultiPart___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHttpMultiPart) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHttpMultiPart) __children_newList() unsafe.Pointer {
	return C.QHttpMultiPart___children_newList(ptr.Pointer())
}

func (ptr *QHttpMultiPart) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHttpMultiPart___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHttpMultiPart) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHttpMultiPart) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHttpMultiPart___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHttpMultiPart) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHttpMultiPart___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHttpMultiPart) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHttpMultiPart) __findChildren_newList() unsafe.Pointer {
	return C.QHttpMultiPart___findChildren_newList(ptr.Pointer())
}

func (ptr *QHttpMultiPart) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHttpMultiPart___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHttpMultiPart) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHttpMultiPart) __findChildren_newList3() unsafe.Pointer {
	return C.QHttpMultiPart___findChildren_newList3(ptr.Pointer())
}

//export callbackQHttpMultiPart_ChildEvent
func callbackQHttpMultiPart_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHttpMultiPart_ConnectNotify
func callbackQHttpMultiPart_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHttpMultiPart_CustomEvent
func callbackQHttpMultiPart_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHttpMultiPart_DeleteLater
func callbackQHttpMultiPart_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHttpMultiPartFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHttpMultiPart) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHttpMultiPart_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHttpMultiPart_Destroyed
func callbackQHttpMultiPart_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHttpMultiPart_DisconnectNotify
func callbackQHttpMultiPart_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHttpMultiPartFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHttpMultiPart) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHttpMultiPart_Event
func callbackQHttpMultiPart_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHttpMultiPartFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHttpMultiPart) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHttpMultiPart_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHttpMultiPart_EventFilter
func callbackQHttpMultiPart_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHttpMultiPartFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHttpMultiPart) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHttpMultiPart_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHttpMultiPart_MetaObject
func callbackQHttpMultiPart_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHttpMultiPartFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHttpMultiPart) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHttpMultiPart_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHttpMultiPart_ObjectNameChanged
func callbackQHttpMultiPart_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHttpMultiPart_TimerEvent
func callbackQHttpMultiPart_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQHttpPartFromPointer(ptr unsafe.Pointer) (n *QHttpPart) {
	n = new(QHttpPart)
	n.SetPointer(ptr)
	return
}
func NewQHttpPart() *QHttpPart {
	tmpValue := NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart())
	qt.SetFinalizer(tmpValue, (*QHttpPart).DestroyQHttpPart)
	return tmpValue
}

func NewQHttpPart2(other QHttpPart_ITF) *QHttpPart {
	tmpValue := NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart2(PointerFromQHttpPart(other)))
	qt.SetFinalizer(tmpValue, (*QHttpPart).DestroyQHttpPart)
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

		qt.SetFinalizer(ptr, nil)
		C.QHttpPart_DestroyQHttpPart(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QIPv6Address struct {
	ptr unsafe.Pointer
}

type QIPv6Address_ITF interface {
	QIPv6Address_PTR() *QIPv6Address
}

func (ptr *QIPv6Address) QIPv6Address_PTR() *QIPv6Address {
	return ptr
}

func (ptr *QIPv6Address) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QIPv6Address) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQIPv6Address(ptr QIPv6Address_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIPv6Address_PTR().Pointer()
	}
	return nil
}

func NewQIPv6AddressFromPointer(ptr unsafe.Pointer) (n *QIPv6Address) {
	n = new(QIPv6Address)
	n.SetPointer(ptr)
	return
}
func (ptr *QIPv6Address) DestroyQIPv6Address() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
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

func NewQLocalServerFromPointer(ptr unsafe.Pointer) (n *QLocalServer) {
	n = new(QLocalServer)
	n.SetPointer(ptr)
	return
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

func NewQLocalServer(parent core.QObject_ITF) *QLocalServer {
	tmpValue := NewQLocalServerFromPointer(C.QLocalServer_NewQLocalServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	if signal := qt.GetSignal(ptr, "hasPendingConnections"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).HasPendingConnectionsDefault())))
}

func (ptr *QLocalServer) ConnectHasPendingConnections(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasPendingConnections"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hasPendingConnections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasPendingConnections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalServer) DisconnectHasPendingConnections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasPendingConnections")
	}
}

func (ptr *QLocalServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalServer_HasPendingConnections(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalServer) HasPendingConnectionsDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalServer_HasPendingConnectionsDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLocalServer_IncomingConnection
func callbackQLocalServer_IncomingConnection(ptr unsafe.Pointer, socketDescriptor C.uintptr_t) {
	if signal := qt.GetSignal(ptr, "incomingConnection"); signal != nil {
		(*(*func(uintptr))(signal))(uintptr(socketDescriptor))
	} else {
		NewQLocalServerFromPointer(ptr).IncomingConnectionDefault(uintptr(socketDescriptor))
	}
}

func (ptr *QLocalServer) ConnectIncomingConnection(f func(socketDescriptor uintptr)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "incomingConnection"); signal != nil {
			f := func(socketDescriptor uintptr) {
				(*(*func(uintptr))(signal))(socketDescriptor)
				f(socketDescriptor)
			}
			qt.ConnectSignal(ptr.Pointer(), "incomingConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "incomingConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalServer) DisconnectIncomingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "incomingConnection")
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
		return int8(C.QLocalServer_IsListening(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalServer) Listen(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QLocalServer_Listen(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
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
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QLocalServer_ConnectNewConnection(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "newConnection")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newConnection")
	}
}

func (ptr *QLocalServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QLocalServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQLocalServer_NextPendingConnection
func callbackQLocalServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextPendingConnection"); signal != nil {
		return PointerFromQLocalSocket((*(*func() *QLocalSocket)(signal))())
	}

	return PointerFromQLocalSocket(NewQLocalServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QLocalServer) ConnectNextPendingConnection(f func() *QLocalSocket) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextPendingConnection"); signal != nil {
			f := func() *QLocalSocket {
				(*(*func() *QLocalSocket)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalServer) DisconnectNextPendingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nextPendingConnection")
	}
}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) NextPendingConnectionDefault() *QLocalSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QLocalServer_RemoveServer(name string) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QLocalServer_QLocalServer_RemoveServer(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
}

func (ptr *QLocalServer) RemoveServer(name string) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QLocalServer_QLocalServer_RemoveServer(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
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

func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {
	if ptr.Pointer() != nil {
		C.QLocalServer_SetSocketOptions(ptr.Pointer(), C.longlong(options))
	}
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {
	if ptr.Pointer() != nil {
		return QLocalServer__SocketOption(C.QLocalServer_SocketOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut *bool) bool {
	if ptr.Pointer() != nil {
		var timedOutC C.char
		if timedOut != nil {
			timedOutC = C.char(int8(qt.GoBoolToInt(*timedOut)))
			defer func() { *timedOut = int8(timedOutC) != 0 }()
		}
		return int8(C.QLocalServer_WaitForNewConnection(ptr.Pointer(), C.int(int32(msec)), &timedOutC)) != 0
	}
	return false
}

//export callbackQLocalServer_DestroyQLocalServer
func callbackQLocalServer_DestroyQLocalServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLocalServer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLocalServerFromPointer(ptr).DestroyQLocalServerDefault()
	}
}

func (ptr *QLocalServer) ConnectDestroyQLocalServer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLocalServer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLocalServer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLocalServer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalServer) DisconnectDestroyQLocalServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLocalServer")
	}
}

func (ptr *QLocalServer) DestroyQLocalServer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocalServer_DestroyQLocalServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalServer) DestroyQLocalServerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocalServer_DestroyQLocalServerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLocalServer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLocalServer) __children_newList() unsafe.Pointer {
	return C.QLocalServer___children_newList(ptr.Pointer())
}

func (ptr *QLocalServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLocalServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLocalServer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QLocalServer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLocalServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLocalServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLocalServer) __findChildren_newList() unsafe.Pointer {
	return C.QLocalServer___findChildren_newList(ptr.Pointer())
}

func (ptr *QLocalServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLocalServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLocalServer) __findChildren_newList3() unsafe.Pointer {
	return C.QLocalServer___findChildren_newList3(ptr.Pointer())
}

//export callbackQLocalServer_ChildEvent
func callbackQLocalServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLocalServer_ConnectNotify
func callbackQLocalServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalServer_CustomEvent
func callbackQLocalServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLocalServer_DeleteLater
func callbackQLocalServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLocalServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLocalServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocalServer_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLocalServer_Destroyed
func callbackQLocalServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLocalServer_DisconnectNotify
func callbackQLocalServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalServer_Event
func callbackQLocalServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLocalServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQLocalServer_EventFilter
func callbackQLocalServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLocalServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLocalServer_MetaObject
func callbackQLocalServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQLocalServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLocalServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLocalServer_ObjectNameChanged
func callbackQLocalServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLocalServer_TimerEvent
func callbackQLocalServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQLocalSocketFromPointer(ptr unsafe.Pointer) (n *QLocalSocket) {
	n = new(QLocalSocket)
	n.SetPointer(ptr)
	return
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

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	tmpValue := NewQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLocalSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

//export callbackQLocalSocket_BytesAvailable
func callbackQLocalSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QLocalSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_BytesToWrite
func callbackQLocalSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesToWrite"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QLocalSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_CanReadLine
func callbackQLocalSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QLocalSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_CanReadLineDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLocalSocket_Close
func callbackQLocalSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLocalSocketFromPointer(ptr).CloseDefault()
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

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(openMode))
	}
}

//export callbackQLocalSocket_Connected
func callbackQLocalSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QLocalSocket_ConnectConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connected")
	}
}

func (ptr *QLocalSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Connected(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

//export callbackQLocalSocket_Disconnected
func callbackQLocalSocket_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QLocalSocket_ConnectDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "disconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "disconnected")
	}
}

func (ptr *QLocalSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Disconnected(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_Error2
func callbackQLocalSocket_Error2(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QLocalSocket__LocalSocketError))(signal))(QLocalSocket__LocalSocketError(socketError))
	}

}

func (ptr *QLocalSocket) ConnectError2(f func(socketError QLocalSocket__LocalSocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QLocalSocket_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(socketError QLocalSocket__LocalSocketError) {
				(*(*func(QLocalSocket__LocalSocketError))(signal))(socketError)
				f(socketError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QLocalSocket) Error2(socketError QLocalSocket__LocalSocketError) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Error2(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QLocalSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_Flush(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalSocket) FullServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

//export callbackQLocalSocket_IsSequential
func callbackQLocalSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QLocalSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_IsSequentialDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_IsValid(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLocalSocket_Open
func callbackQLocalSocket_Open(ptr unsafe.Pointer, openMode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(core.QIODevice__OpenModeFlag) bool)(signal))(core.QIODevice__OpenModeFlag(openMode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(openMode)))))
}

func (ptr *QLocalSocket) OpenDefault(openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_OpenDefault(ptr.Pointer(), C.longlong(openMode))) != 0
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
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(c)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQLocalSocketFromPointer(ptr).ReadDataDefault(&retS, int64(c)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QLocalSocket) ConnectReadData(f func(data *string, c int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			f := func(data *string, c int64) int64 {
				(*(*func(*string, int64) int64)(signal))(data, c)
				return f(data, c)
			}
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "readData")
	}
}

func (ptr *QLocalSocket) ReadData(data *string, c int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(c)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QLocalSocket_ReadData(ptr.Pointer(), dataC, C.longlong(c)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QLocalSocket) ReadDataDefault(data *string, c int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(c)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QLocalSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(c)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QLocalSocket) ServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QLocalSocket) SetServerName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QLocalSocket_SetServerName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})
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
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QLocalSocket__LocalSocketState))(signal))(QLocalSocket__LocalSocketState(socketState))
	}

}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QLocalSocket_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(socketState QLocalSocket__LocalSocketState) {
				(*(*func(QLocalSocket__LocalSocketState))(signal))(socketState)
				f(socketState)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QLocalSocket) StateChanged(socketState QLocalSocket__LocalSocketState) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_StateChanged(ptr.Pointer(), C.longlong(socketState))
	}
}

//export callbackQLocalSocket_WaitForBytesWritten
func callbackQLocalSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QLocalSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQLocalSocket_WaitForReadyRead
func callbackQLocalSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QLocalSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQLocalSocket_WriteData
func callbackQLocalSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, c C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(c)))
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(c)))
}

func (ptr *QLocalSocket) ConnectWriteData(f func(data []byte, c int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			f := func(data []byte, c int64) int64 {
				(*(*func([]byte, int64) int64)(signal))(data, c)
				return f(data, c)
			}
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QLocalSocket) WriteData(data []byte, c int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QLocalSocket_WriteData(ptr.Pointer(), dataC, C.longlong(c)))
	}
	return 0
}

func (ptr *QLocalSocket) WriteDataDefault(data []byte, c int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QLocalSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(c)))
	}
	return 0
}

//export callbackQLocalSocket_DestroyQLocalSocket
func callbackQLocalSocket_DestroyQLocalSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLocalSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLocalSocketFromPointer(ptr).DestroyQLocalSocketDefault()
	}
}

func (ptr *QLocalSocket) ConnectDestroyQLocalSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLocalSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLocalSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLocalSocket", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLocalSocket) DisconnectDestroyQLocalSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLocalSocket")
	}
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) DestroyQLocalSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocalSocket_DestroyQLocalSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLocalSocket___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalSocket) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLocalSocket) __children_newList() unsafe.Pointer {
	return C.QLocalSocket___children_newList(ptr.Pointer())
}

func (ptr *QLocalSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLocalSocket___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLocalSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLocalSocket) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QLocalSocket___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLocalSocket) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLocalSocket___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalSocket) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLocalSocket) __findChildren_newList() unsafe.Pointer {
	return C.QLocalSocket___findChildren_newList(ptr.Pointer())
}

func (ptr *QLocalSocket) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLocalSocket___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalSocket) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLocalSocket) __findChildren_newList3() unsafe.Pointer {
	return C.QLocalSocket___findChildren_newList3(ptr.Pointer())
}

//export callbackQLocalSocket_AboutToClose
func callbackQLocalSocket_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLocalSocket_AtEnd
func callbackQLocalSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QLocalSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_AtEndDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLocalSocket_BytesWritten
func callbackQLocalSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(bytes))
	}

}

//export callbackQLocalSocket_ChannelBytesWritten
func callbackQLocalSocket_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		(*(*func(int, int64))(signal))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQLocalSocket_ChannelReadyRead
func callbackQLocalSocket_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		(*(*func(int))(signal))(int(int32(channel)))
	}

}

//export callbackQLocalSocket_Pos
func callbackQLocalSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).PosDefault())
}

func (ptr *QLocalSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_ReadChannelFinished
func callbackQLocalSocket_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLocalSocket_ReadLineData
func callbackQLocalSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readLineData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QLocalSocket) ReadLineDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QLocalSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQLocalSocket_ReadyRead
func callbackQLocalSocket_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLocalSocket_Reset
func callbackQLocalSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QLocalSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_ResetDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLocalSocket_Seek
func callbackQLocalSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int64) bool)(signal))(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QLocalSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_SeekDefault(ptr.Pointer(), C.longlong(pos))) != 0
	}
	return false
}

//export callbackQLocalSocket_Size
func callbackQLocalSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QLocalSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_ChildEvent
func callbackQLocalSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLocalSocket_ConnectNotify
func callbackQLocalSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalSocket_CustomEvent
func callbackQLocalSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLocalSocket_DeleteLater
func callbackQLocalSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLocalSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLocalSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocalSocket_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLocalSocket_Destroyed
func callbackQLocalSocket_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLocalSocket_DisconnectNotify
func callbackQLocalSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalSocket_Event
func callbackQLocalSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLocalSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQLocalSocket_EventFilter
func callbackQLocalSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLocalSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLocalSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLocalSocket_MetaObject
func callbackQLocalSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQLocalSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLocalSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLocalSocket_ObjectNameChanged
func callbackQLocalSocket_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLocalSocket_TimerEvent
func callbackQLocalSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQNetworkAccessManagerFromPointer(ptr unsafe.Pointer) (n *QNetworkAccessManager) {
	n = new(QNetworkAccessManager)
	n.SetPointer(ptr)
	return
}

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

//go:generate stringer -type=QNetworkAccessManager__NetworkAccessibility
//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(1)
)

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {
	tmpValue := NewQNetworkAccessManagerFromPointer(C.QNetworkAccessManager_NewQNetworkAccessManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkAccessManager) ActiveConfiguration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkAccessManager_ActiveConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) AddStrictTransportSecurityHosts(knownHosts []*QHstsPolicy) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_AddStrictTransportSecurityHosts(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkAccessManagerFromPointer(NewQNetworkAccessManagerFromPointer(nil).__addStrictTransportSecurityHosts_knownHosts_newList())
			for _, v := range knownHosts {
				tmpList.__addStrictTransportSecurityHosts_knownHosts_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQNetworkAccessManager_AuthenticationRequired
func callbackQNetworkAccessManager_AuthenticationRequired(ptr unsafe.Pointer, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "authenticationRequired"); signal != nil {
		(*(*func(*QNetworkReply, *QAuthenticator))(signal))(NewQNetworkReplyFromPointer(reply), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "authenticationRequired") {
			C.QNetworkAccessManager_ConnectAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "authenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "authenticationRequired"); signal != nil {
			f := func(reply *QNetworkReply, authenticator *QAuthenticator) {
				(*(*func(*QNetworkReply, *QAuthenticator))(signal))(reply, authenticator)
				f(reply, authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "authenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "authenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "authenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF, authenticator QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_AuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractNetworkCacheFromPointer(C.QNetworkAccessManager_Cache(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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

func (ptr *QNetworkAccessManager) ClearConnectionCache() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearConnectionCache(ptr.Pointer())
	}
}

func (ptr *QNetworkAccessManager) Configuration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkAccessManager_Configuration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectToHost(hostName string, port uint16) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QNetworkAccessManager_ConnectToHost(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port))
	}
}

func (ptr *QNetworkAccessManager) ConnectToHostEncrypted(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QNetworkAccessManager_ConnectToHostEncrypted(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QNetworkAccessManager) ConnectToHostEncrypted2(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF, peerName string) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		var peerNameC *C.char
		if peerName != "" {
			peerNameC = C.CString(peerName)
			defer C.free(unsafe.Pointer(peerNameC))
		}
		C.QNetworkAccessManager_ConnectToHostEncrypted2(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), PointerFromQSslConfiguration(sslConfiguration), C.struct_QtNetwork_PackedString{data: peerNameC, len: C.longlong(len(peerName))})
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCookieJarFromPointer(C.QNetworkAccessManager_CookieJar(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_CreateRequest
func callbackQNetworkAccessManager_CreateRequest(ptr unsafe.Pointer, op C.longlong, originalReq unsafe.Pointer, outgoingData unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createRequest"); signal != nil {
		return PointerFromQNetworkReply((*(*func(QNetworkAccessManager__Operation, *QNetworkRequest, *core.QIODevice) *QNetworkReply)(signal))(QNetworkAccessManager__Operation(op), NewQNetworkRequestFromPointer(originalReq), core.NewQIODeviceFromPointer(outgoingData)))
	}

	return PointerFromQNetworkReply(NewQNetworkAccessManagerFromPointer(ptr).CreateRequestDefault(QNetworkAccessManager__Operation(op), NewQNetworkRequestFromPointer(originalReq), core.NewQIODeviceFromPointer(outgoingData)))
}

func (ptr *QNetworkAccessManager) ConnectCreateRequest(f func(op QNetworkAccessManager__Operation, originalReq *QNetworkRequest, outgoingData *core.QIODevice) *QNetworkReply) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createRequest"); signal != nil {
			f := func(op QNetworkAccessManager__Operation, originalReq *QNetworkRequest, outgoingData *core.QIODevice) *QNetworkReply {
				(*(*func(QNetworkAccessManager__Operation, *QNetworkRequest, *core.QIODevice) *QNetworkReply)(signal))(op, originalReq, outgoingData)
				return f(op, originalReq, outgoingData)
			}
			qt.ConnectSignal(ptr.Pointer(), "createRequest", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createRequest", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectCreateRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createRequest")
	}
}

func (ptr *QNetworkAccessManager) CreateRequest(op QNetworkAccessManager__Operation, originalReq QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequest(ptr.Pointer(), C.longlong(op), PointerFromQNetworkRequest(originalReq), core.PointerFromQIODevice(outgoingData)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) CreateRequestDefault(op QNetworkAccessManager__Operation, originalReq QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequestDefault(ptr.Pointer(), C.longlong(op), PointerFromQNetworkRequest(originalReq), core.PointerFromQIODevice(outgoingData)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_DeleteResource(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) EnableStrictTransportSecurityStore(enabled bool, storeDir string) {
	if ptr.Pointer() != nil {
		var storeDirC *C.char
		if storeDir != "" {
			storeDirC = C.CString(storeDir)
			defer C.free(unsafe.Pointer(storeDirC))
		}
		C.QNetworkAccessManager_EnableStrictTransportSecurityStore(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))), C.struct_QtNetwork_PackedString{data: storeDirC, len: C.longlong(len(storeDir))})
	}
}

//export callbackQNetworkAccessManager_Encrypted
func callbackQNetworkAccessManager_Encrypted(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "encrypted"); signal != nil {
		(*(*func(*QNetworkReply))(signal))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "encrypted") {
			C.QNetworkAccessManager_ConnectEncrypted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "encrypted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "encrypted"); signal != nil {
			f := func(reply *QNetworkReply) {
				(*(*func(*QNetworkReply))(signal))(reply)
				f(reply)
			}
			qt.ConnectSignal(ptr.Pointer(), "encrypted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encrypted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "encrypted")
	}
}

func (ptr *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Encrypted(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

//export callbackQNetworkAccessManager_Finished
func callbackQNetworkAccessManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(*QNetworkReply))(signal))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QNetworkAccessManager_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(reply *QNetworkReply) {
				(*(*func(*QNetworkReply))(signal))(reply)
				f(reply)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QNetworkAccessManager) Finished(reply QNetworkReply_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Finished(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Get(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Head(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) IsStrictTransportSecurityEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAccessManager_IsStrictTransportSecurityEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkAccessManager) IsStrictTransportSecurityStoreEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAccessManager_IsStrictTransportSecurityStoreEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkAccessManager_NetworkAccessibleChanged
func callbackQNetworkAccessManager_NetworkAccessibleChanged(ptr unsafe.Pointer, accessible C.longlong) {
	if signal := qt.GetSignal(ptr, "networkAccessibleChanged"); signal != nil {
		(*(*func(QNetworkAccessManager__NetworkAccessibility))(signal))(QNetworkAccessManager__NetworkAccessibility(accessible))
	}

}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "networkAccessibleChanged") {
			C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "networkAccessibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "networkAccessibleChanged"); signal != nil {
			f := func(accessible QNetworkAccessManager__NetworkAccessibility) {
				(*(*func(QNetworkAccessManager__NetworkAccessibility))(signal))(accessible)
				f(accessible)
			}
			qt.ConnectSignal(ptr.Pointer(), "networkAccessibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "networkAccessibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "networkAccessibleChanged")
	}
}

func (ptr *QNetworkAccessManager) NetworkAccessibleChanged(accessible QNetworkAccessManager__NetworkAccessibility) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_NetworkAccessibleChanged(ptr.Pointer(), C.longlong(accessible))
	}
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "preSharedKeyAuthenticationRequired"); signal != nil {
		(*(*func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(signal))(NewQNetworkReplyFromPointer(reply), NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preSharedKeyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			f := func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator) {
				(*(*func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(signal))(reply, authenticator)
				f(reply, authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF, authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Proxy() *QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkProxyFromPointer(C.QNetworkAccessManager_Proxy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_ProxyAuthenticationRequired
func callbackQNetworkAccessManager_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "proxyAuthenticationRequired"); signal != nil {
		(*(*func(*QNetworkProxy, *QAuthenticator))(signal))(NewQNetworkProxyFromPointer(proxy), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "proxyAuthenticationRequired") {
			C.QNetworkAccessManager_ConnectProxyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "proxyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "proxyAuthenticationRequired"); signal != nil {
			f := func(proxy *QNetworkProxy, authenticator *QAuthenticator) {
				(*(*func(*QNetworkProxy, *QAuthenticator))(signal))(proxy, authenticator)
				f(proxy, authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "proxyAuthenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ProxyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkProxy(proxy), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {
	if ptr.Pointer() != nil {
		return NewQNetworkProxyFactoryFromPointer(C.QNetworkAccessManager_ProxyFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) RedirectPolicy() QNetworkRequest__RedirectPolicy {
	if ptr.Pointer() != nil {
		return QNetworkRequest__RedirectPolicy(C.QNetworkAccessManager_RedirectPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest2(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), core.PointerFromQByteArray(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest3(request QNetworkRequest_ITF, verb core.QByteArray_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest3(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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

func (ptr *QNetworkAccessManager) SetRedirectPolicy(policy QNetworkRequest__RedirectPolicy) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetRedirectPolicy(ptr.Pointer(), C.longlong(policy))
	}
}

func (ptr *QNetworkAccessManager) SetStrictTransportSecurityEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetStrictTransportSecurityEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQNetworkAccessManager_SslErrors
func callbackQNetworkAccessManager_SslErrors(ptr unsafe.Pointer, reply unsafe.Pointer, errors C.struct_QtNetwork_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors"); signal != nil {
		(*(*func(*QNetworkReply, []*QSslError))(signal))(NewQNetworkReplyFromPointer(reply), func(l C.struct_QtNetwork_PackedList) []*QSslError {
			out := make([]*QSslError, int(l.len))
			tmpList := NewQNetworkAccessManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sslErrors_errors_atList(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QNetworkAccessManager) ConnectSslErrors(f func(reply *QNetworkReply, errors []*QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors") {
			C.QNetworkAccessManager_ConnectSslErrors(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sslErrors")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors"); signal != nil {
			f := func(reply *QNetworkReply, errors []*QSslError) {
				(*(*func(*QNetworkReply, []*QSslError))(signal))(reply, errors)
				f(reply, errors)
			}
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectSslErrors() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectSslErrors(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sslErrors")
	}
}

func (ptr *QNetworkAccessManager) SslErrors(reply QNetworkReply_ITF, errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SslErrors(ptr.Pointer(), PointerFromQNetworkReply(reply), func() unsafe.Pointer {
			tmpList := NewQNetworkAccessManagerFromPointer(NewQNetworkAccessManagerFromPointer(nil).__sslErrors_errors_newList())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QNetworkAccessManager) StrictTransportSecurityHosts() []*QHstsPolicy {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QHstsPolicy {
			out := make([]*QHstsPolicy, int(l.len))
			tmpList := NewQNetworkAccessManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__strictTransportSecurityHosts_atList(i)
			}
			return out
		}(C.QNetworkAccessManager_StrictTransportSecurityHosts(ptr.Pointer()))
	}
	return make([]*QHstsPolicy, 0)
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QNetworkAccessManager_SupportedSchemes(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQNetworkAccessManager_SupportedSchemesImplementation
func callbackQNetworkAccessManager_SupportedSchemesImplementation(ptr unsafe.Pointer) C.struct_QtNetwork_PackedString {
	if signal := qt.GetSignal(ptr, "supportedSchemesImplementation"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_QtNetwork_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewQNetworkAccessManagerFromPointer(ptr).SupportedSchemesImplementationDefault()
	return C.struct_QtNetwork_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QNetworkAccessManager) ConnectSupportedSchemesImplementation(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedSchemesImplementation"); signal != nil {
			f := func() []string {
				(*(*func() []string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "supportedSchemesImplementation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedSchemesImplementation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectSupportedSchemesImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportedSchemesImplementation")
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementation() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QNetworkAccessManager_SupportedSchemesImplementation(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementationDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QNetworkAccessManager_SupportedSchemesImplementationDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQNetworkAccessManager_DestroyQNetworkAccessManager
func callbackQNetworkAccessManager_DestroyQNetworkAccessManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNetworkAccessManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DestroyQNetworkAccessManagerDefault()
	}
}

func (ptr *QNetworkAccessManager) ConnectDestroyQNetworkAccessManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkAccessManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkAccessManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkAccessManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkAccessManager) DisconnectDestroyQNetworkAccessManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkAccessManager")
	}
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkAccessManager_DestroyQNetworkAccessManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkAccessManager) __addStrictTransportSecurityHosts_knownHosts_atList(i int) *QHstsPolicy {
	if ptr.Pointer() != nil {
		tmpValue := NewQHstsPolicyFromPointer(C.QNetworkAccessManager___addStrictTransportSecurityHosts_knownHosts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHstsPolicy).DestroyQHstsPolicy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __addStrictTransportSecurityHosts_knownHosts_setList(i QHstsPolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___addStrictTransportSecurityHosts_knownHosts_setList(ptr.Pointer(), PointerFromQHstsPolicy(i))
	}
}

func (ptr *QNetworkAccessManager) __addStrictTransportSecurityHosts_knownHosts_newList() unsafe.Pointer {
	return C.QNetworkAccessManager___addStrictTransportSecurityHosts_knownHosts_newList(ptr.Pointer())
}

func (ptr *QNetworkAccessManager) __sslErrors_errors_atList(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QNetworkAccessManager___sslErrors_errors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __sslErrors_errors_setList(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___sslErrors_errors_setList(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QNetworkAccessManager) __sslErrors_errors_newList() unsafe.Pointer {
	return C.QNetworkAccessManager___sslErrors_errors_newList(ptr.Pointer())
}

func (ptr *QNetworkAccessManager) __strictTransportSecurityHosts_atList(i int) *QHstsPolicy {
	if ptr.Pointer() != nil {
		tmpValue := NewQHstsPolicyFromPointer(C.QNetworkAccessManager___strictTransportSecurityHosts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHstsPolicy).DestroyQHstsPolicy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __strictTransportSecurityHosts_setList(i QHstsPolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___strictTransportSecurityHosts_setList(ptr.Pointer(), PointerFromQHstsPolicy(i))
	}
}

func (ptr *QNetworkAccessManager) __strictTransportSecurityHosts_newList() unsafe.Pointer {
	return C.QNetworkAccessManager___strictTransportSecurityHosts_newList(ptr.Pointer())
}

func (ptr *QNetworkAccessManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkAccessManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkAccessManager) __children_newList() unsafe.Pointer {
	return C.QNetworkAccessManager___children_newList(ptr.Pointer())
}

func (ptr *QNetworkAccessManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkAccessManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkAccessManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNetworkAccessManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNetworkAccessManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkAccessManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkAccessManager) __findChildren_newList() unsafe.Pointer {
	return C.QNetworkAccessManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QNetworkAccessManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkAccessManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkAccessManager) __findChildren_newList3() unsafe.Pointer {
	return C.QNetworkAccessManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQNetworkAccessManager_ChildEvent
func callbackQNetworkAccessManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkAccessManager_ConnectNotify
func callbackQNetworkAccessManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkAccessManager_CustomEvent
func callbackQNetworkAccessManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkAccessManager_DeleteLater
func callbackQNetworkAccessManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkAccessManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkAccessManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNetworkAccessManager_Destroyed
func callbackQNetworkAccessManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNetworkAccessManager_DisconnectNotify
func callbackQNetworkAccessManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkAccessManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkAccessManager_Event
func callbackQNetworkAccessManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkAccessManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkAccessManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAccessManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNetworkAccessManager_EventFilter
func callbackQNetworkAccessManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkAccessManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkAccessManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAccessManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNetworkAccessManager_MetaObject
func callbackQNetworkAccessManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNetworkAccessManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkAccessManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkAccessManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNetworkAccessManager_ObjectNameChanged
func callbackQNetworkAccessManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNetworkAccessManager_TimerEvent
func callbackQNetworkAccessManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) (n *QNetworkAddressEntry) {
	n = new(QNetworkAddressEntry)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNetworkAddressEntry__DnsEligibilityStatus
//QNetworkAddressEntry::DnsEligibilityStatus
type QNetworkAddressEntry__DnsEligibilityStatus int64

const (
	QNetworkAddressEntry__DnsEligibilityUnknown QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(-1)
	QNetworkAddressEntry__DnsIneligible         QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(0)
	QNetworkAddressEntry__DnsEligible           QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(1)
)

func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	tmpValue := NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry())
	qt.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
	return tmpValue
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {
	tmpValue := NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry2(PointerFromQNetworkAddressEntry(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
	return tmpValue
}

func (ptr *QNetworkAddressEntry) Broadcast() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QNetworkAddressEntry_Broadcast(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) ClearAddressLifetime() {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_ClearAddressLifetime(ptr.Pointer())
	}
}

func (ptr *QNetworkAddressEntry) DnsEligibility() QNetworkAddressEntry__DnsEligibilityStatus {
	if ptr.Pointer() != nil {
		return QNetworkAddressEntry__DnsEligibilityStatus(C.QNetworkAddressEntry_DnsEligibility(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAddressEntry) Ip() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QNetworkAddressEntry_Ip(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) IsLifetimeKnown() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAddressEntry_IsLifetimeKnown(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkAddressEntry) IsPermanent() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAddressEntry_IsPermanent(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkAddressEntry) IsTemporary() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkAddressEntry_IsTemporary(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkAddressEntry) Netmask() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QNetworkAddressEntry_Netmask(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
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

func (ptr *QNetworkAddressEntry) SetAddressLifetime(preferred core.QDeadlineTimer_ITF, validity core.QDeadlineTimer_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetAddressLifetime(ptr.Pointer(), core.PointerFromQDeadlineTimer(preferred), core.PointerFromQDeadlineTimer(validity))
	}
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetBroadcast(ptr.Pointer(), PointerFromQHostAddress(newBroadcast))
	}
}

func (ptr *QNetworkAddressEntry) SetDnsEligibility(status QNetworkAddressEntry__DnsEligibilityStatus) {
	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetDnsEligibility(ptr.Pointer(), C.longlong(status))
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

		qt.SetFinalizer(ptr, nil)
		C.QNetworkAddressEntry_DestroyQNetworkAddressEntry(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) (n *QNetworkCacheMetaData) {
	n = new(QNetworkCacheMetaData)
	n.SetPointer(ptr)
	return
}
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	tmpValue := NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData())
	qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
	return tmpValue
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {
	tmpValue := NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(PointerFromQNetworkCacheMetaData(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
	return tmpValue
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_ExpirationDate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCacheMetaData_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_LastModified(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) RawHeaders() []*QNetworkCacheMetaData_RawHeader {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkCacheMetaData_RawHeader {
			out := make([]*QNetworkCacheMetaData_RawHeader, int(l.len))
			tmpList := NewQNetworkCacheMetaDataFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__rawHeaders_atList(i)
			}
			return out
		}(C.QNetworkCacheMetaData_RawHeaders(ptr.Pointer()))
	}
	return make([]*QNetworkCacheMetaData_RawHeader, 0)
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCacheMetaData_SaveToDisk(ptr.Pointer())) != 0
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

func (ptr *QNetworkCacheMetaData) SetRawHeaders(list []*QNetworkCacheMetaData_RawHeader) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetRawHeaders(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkCacheMetaDataFromPointer(NewQNetworkCacheMetaDataFromPointer(nil).__setRawHeaders_list_newList())
			for _, v := range list {
				tmpList.__setRawHeaders_list_setList(v)
			}
			return tmpList.Pointer()
		}())
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
		tmpValue := core.NewQUrlFromPointer(C.QNetworkCacheMetaData_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCacheMetaData) __rawHeaders_atList(i int) *QNetworkCacheMetaData_RawHeader {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCacheMetaData_RawHeaderFromPointer(C.QNetworkCacheMetaData___rawHeaders_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData_RawHeader).DestroyQNetworkCacheMetaData_RawHeader)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) __rawHeaders_setList(i QNetworkCacheMetaData_RawHeader_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData___rawHeaders_setList(ptr.Pointer(), PointerFromQNetworkCacheMetaData_RawHeader(i))
	}
}

func (ptr *QNetworkCacheMetaData) __rawHeaders_newList() unsafe.Pointer {
	return C.QNetworkCacheMetaData___rawHeaders_newList(ptr.Pointer())
}

func (ptr *QNetworkCacheMetaData) __setRawHeaders_list_atList(i int) *QNetworkCacheMetaData_RawHeader {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCacheMetaData_RawHeaderFromPointer(C.QNetworkCacheMetaData___setRawHeaders_list_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData_RawHeader).DestroyQNetworkCacheMetaData_RawHeader)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) __setRawHeaders_list_setList(i QNetworkCacheMetaData_RawHeader_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData___setRawHeaders_list_setList(ptr.Pointer(), PointerFromQNetworkCacheMetaData_RawHeader(i))
	}
}

func (ptr *QNetworkCacheMetaData) __setRawHeaders_list_newList() unsafe.Pointer {
	return C.QNetworkCacheMetaData___setRawHeaders_list_newList(ptr.Pointer())
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

func NewQNetworkConfigurationFromPointer(ptr unsafe.Pointer) (n *QNetworkConfiguration) {
	n = new(QNetworkConfiguration)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNetworkConfiguration__Type
//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint QNetworkConfiguration__Type = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      QNetworkConfiguration__Type = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          QNetworkConfiguration__Type = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             QNetworkConfiguration__Type = QNetworkConfiguration__Type(3)
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

func NewQNetworkConfiguration() *QNetworkConfiguration {
	tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration())
	qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
	return tmpValue
}

func NewQNetworkConfiguration2(other QNetworkConfiguration_ITF) *QNetworkConfiguration {
	tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration2(PointerFromQNetworkConfiguration(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
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
			out := make([]*QNetworkConfiguration, int(l.len))
			tmpList := NewQNetworkConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__children_atList(i)
			}
			return out
		}(C.QNetworkConfiguration_Children(ptr.Pointer()))
	}
	return make([]*QNetworkConfiguration, 0)
}

func (ptr *QNetworkConfiguration) ConnectTimeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkConfiguration_ConnectTimeout(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkConfiguration) Identifier() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkConfiguration_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkConfiguration_IsRoamingAvailable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkConfiguration_IsValid(ptr.Pointer())) != 0
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

func (ptr *QNetworkConfiguration) SetConnectTimeout(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkConfiguration_SetConnectTimeout(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
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

		qt.SetFinalizer(ptr, nil)
		C.QNetworkConfiguration_DestroyQNetworkConfiguration(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfiguration) __children_atList(i int) *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration___children_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfiguration) __children_setList(i QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfiguration___children_setList(ptr.Pointer(), PointerFromQNetworkConfiguration(i))
	}
}

func (ptr *QNetworkConfiguration) __children_newList() unsafe.Pointer {
	return C.QNetworkConfiguration___children_newList(ptr.Pointer())
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

func NewQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) (n *QNetworkConfigurationManager) {
	n = new(QNetworkConfigurationManager)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQNetworkConfigurationManagerFromPointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkConfigurationManager) AllConfigurations(filter QNetworkConfiguration__StateFlag) []*QNetworkConfiguration {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkConfiguration {
			out := make([]*QNetworkConfiguration, int(l.len))
			tmpList := NewQNetworkConfigurationManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__allConfigurations_atList(i)
			}
			return out
		}(C.QNetworkConfigurationManager_AllConfigurations(ptr.Pointer(), C.longlong(filter)))
	}
	return make([]*QNetworkConfiguration, 0)
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkConfigurationManager_ConfigurationAdded
func callbackQNetworkConfigurationManager_ConfigurationAdded(ptr unsafe.Pointer, config unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "configurationAdded"); signal != nil {
		(*(*func(*QNetworkConfiguration))(signal))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationAdded(f func(config *QNetworkConfiguration)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "configurationAdded") {
			C.QNetworkConfigurationManager_ConnectConfigurationAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "configurationAdded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "configurationAdded"); signal != nil {
			f := func(config *QNetworkConfiguration) {
				(*(*func(*QNetworkConfiguration))(signal))(config)
				f(config)
			}
			qt.ConnectSignal(ptr.Pointer(), "configurationAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "configurationAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationAdded() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "configurationAdded")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationAdded(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationAdded(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

//export callbackQNetworkConfigurationManager_ConfigurationChanged
func callbackQNetworkConfigurationManager_ConfigurationChanged(ptr unsafe.Pointer, config unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "configurationChanged"); signal != nil {
		(*(*func(*QNetworkConfiguration))(signal))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationChanged(f func(config *QNetworkConfiguration)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "configurationChanged") {
			C.QNetworkConfigurationManager_ConnectConfigurationChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "configurationChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "configurationChanged"); signal != nil {
			f := func(config *QNetworkConfiguration) {
				(*(*func(*QNetworkConfiguration))(signal))(config)
				f(config)
			}
			qt.ConnectSignal(ptr.Pointer(), "configurationChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "configurationChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "configurationChanged")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationChanged(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationChanged(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationFromIdentifier(identifier string) *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_ConfigurationFromIdentifier(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: identifierC, len: C.longlong(len(identifier))}))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkConfigurationManager_ConfigurationRemoved
func callbackQNetworkConfigurationManager_ConfigurationRemoved(ptr unsafe.Pointer, config unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "configurationRemoved"); signal != nil {
		(*(*func(*QNetworkConfiguration))(signal))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationRemoved(f func(config *QNetworkConfiguration)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "configurationRemoved") {
			C.QNetworkConfigurationManager_ConnectConfigurationRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "configurationRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "configurationRemoved"); signal != nil {
			f := func(config *QNetworkConfiguration) {
				(*(*func(*QNetworkConfiguration))(signal))(config)
				f(config)
			}
			qt.ConnectSignal(ptr.Pointer(), "configurationRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "configurationRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationRemoved() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "configurationRemoved")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationRemoved(config QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationRemoved(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkConfigurationManager) DefaultConfiguration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_DefaultConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkConfigurationManager_IsOnline(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_OnlineStateChanged
func callbackQNetworkConfigurationManager_OnlineStateChanged(ptr unsafe.Pointer, isOnline C.char) {
	if signal := qt.GetSignal(ptr, "onlineStateChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(isOnline) != 0)
	}

}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onlineStateChanged") {
			C.QNetworkConfigurationManager_ConnectOnlineStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "onlineStateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onlineStateChanged"); signal != nil {
			f := func(isOnline bool) {
				(*(*func(bool))(signal))(isOnline)
				f(isOnline)
			}
			qt.ConnectSignal(ptr.Pointer(), "onlineStateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onlineStateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onlineStateChanged")
	}
}

func (ptr *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_OnlineStateChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isOnline))))
	}
}

//export callbackQNetworkConfigurationManager_UpdateCompleted
func callbackQNetworkConfigurationManager_UpdateCompleted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateCompleted"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateCompleted") {
			C.QNetworkConfigurationManager_ConnectUpdateCompleted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "updateCompleted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateCompleted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "updateCompleted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateCompleted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateCompleted")
	}
}

func (ptr *QNetworkConfigurationManager) UpdateCompleted() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateCompleted(ptr.Pointer())
	}
}

//export callbackQNetworkConfigurationManager_UpdateConfigurations
func callbackQNetworkConfigurationManager_UpdateConfigurations(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateConfigurations"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).UpdateConfigurationsDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateConfigurations(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateConfigurations"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "updateConfigurations", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateConfigurations", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateConfigurations() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateConfigurations")
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
	if signal := qt.GetSignal(ptr, "~QNetworkConfigurationManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DestroyQNetworkConfigurationManagerDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDestroyQNetworkConfigurationManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkConfigurationManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkConfigurationManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkConfigurationManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDestroyQNetworkConfigurationManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkConfigurationManager")
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) __allConfigurations_atList(i int) *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager___allConfigurations_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) __allConfigurations_setList(i QNetworkConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager___allConfigurations_setList(ptr.Pointer(), PointerFromQNetworkConfiguration(i))
	}
}

func (ptr *QNetworkConfigurationManager) __allConfigurations_newList() unsafe.Pointer {
	return C.QNetworkConfigurationManager___allConfigurations_newList(ptr.Pointer())
}

func (ptr *QNetworkConfigurationManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkConfigurationManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkConfigurationManager) __children_newList() unsafe.Pointer {
	return C.QNetworkConfigurationManager___children_newList(ptr.Pointer())
}

func (ptr *QNetworkConfigurationManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkConfigurationManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkConfigurationManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNetworkConfigurationManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNetworkConfigurationManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkConfigurationManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkConfigurationManager) __findChildren_newList() unsafe.Pointer {
	return C.QNetworkConfigurationManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QNetworkConfigurationManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkConfigurationManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkConfigurationManager) __findChildren_newList3() unsafe.Pointer {
	return C.QNetworkConfigurationManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQNetworkConfigurationManager_ChildEvent
func callbackQNetworkConfigurationManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_ConnectNotify
func callbackQNetworkConfigurationManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkConfigurationManager_CustomEvent
func callbackQNetworkConfigurationManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_DeleteLater
func callbackQNetworkConfigurationManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkConfigurationManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkConfigurationManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNetworkConfigurationManager_Destroyed
func callbackQNetworkConfigurationManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNetworkConfigurationManager_DisconnectNotify
func callbackQNetworkConfigurationManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkConfigurationManager_Event
func callbackQNetworkConfigurationManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkConfigurationManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkConfigurationManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkConfigurationManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_EventFilter
func callbackQNetworkConfigurationManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkConfigurationManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkConfigurationManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkConfigurationManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_MetaObject
func callbackQNetworkConfigurationManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNetworkConfigurationManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkConfigurationManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkConfigurationManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNetworkConfigurationManager_ObjectNameChanged
func callbackQNetworkConfigurationManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNetworkConfigurationManager_TimerEvent
func callbackQNetworkConfigurationManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQNetworkCookieFromPointer(ptr unsafe.Pointer) (n *QNetworkCookie) {
	n = new(QNetworkCookie)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNetworkCookie__RawForm
//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             QNetworkCookie__RawForm = QNetworkCookie__RawForm(1)
)

func NewQNetworkCookie(name core.QByteArray_ITF, value core.QByteArray_ITF) *QNetworkCookie {
	tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie(core.PointerFromQByteArray(name), core.PointerFromQByteArray(value)))
	qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
	return tmpValue
}

func NewQNetworkCookie2(other QNetworkCookie_ITF) *QNetworkCookie {
	tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie2(PointerFromQNetworkCookie(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
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
		tmpValue := core.NewQDateTimeFromPointer(C.QNetworkCookie_ExpirationDate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookie_HasSameIdentifier(ptr.Pointer(), PointerFromQNetworkCookie(other))) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookie_IsHttpOnly(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSecure() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookie_IsSecure(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookie_IsSessionCookie(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCookie) Name() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkCookie_Name(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
		out := make([]*QNetworkCookie, int(l.len))
		tmpList := NewQNetworkCookieFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__parseCookies_atList(i)
		}
		return out
	}(C.QNetworkCookie_QNetworkCookie_ParseCookies(core.PointerFromQByteArray(cookieString)))
}

func (ptr *QNetworkCookie) ParseCookies(cookieString core.QByteArray_ITF) []*QNetworkCookie {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
		out := make([]*QNetworkCookie, int(l.len))
		tmpList := NewQNetworkCookieFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__parseCookies_atList(i)
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
		var domainC *C.char
		if domain != "" {
			domainC = C.CString(domain)
			defer C.free(unsafe.Pointer(domainC))
		}
		C.QNetworkCookie_SetDomain(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: domainC, len: C.longlong(len(domain))})
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
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QNetworkCookie_SetPath(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))})
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
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkCookie_ToRawForm(ptr.Pointer(), C.longlong(form)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkCookie_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkCookie_DestroyQNetworkCookie(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookie) __parseCookies_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookie___parseCookies_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) __parseCookies_setList(i QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie___parseCookies_setList(ptr.Pointer(), PointerFromQNetworkCookie(i))
	}
}

func (ptr *QNetworkCookie) __parseCookies_newList() unsafe.Pointer {
	return C.QNetworkCookie___parseCookies_newList(ptr.Pointer())
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

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) (n *QNetworkCookieJar) {
	n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	return
}
func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {
	tmpValue := NewQNetworkCookieJarFromPointer(C.QNetworkCookieJar_NewQNetworkCookieJar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkCookieJar) AllCookies() []*QNetworkCookie {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
			out := make([]*QNetworkCookie, int(l.len))
			tmpList := NewQNetworkCookieJarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__allCookies_atList(i)
			}
			return out
		}(C.QNetworkCookieJar_AllCookies(ptr.Pointer()))
	}
	return make([]*QNetworkCookie, 0)
}

//export callbackQNetworkCookieJar_CookiesForUrl
func callbackQNetworkCookieJar_CookiesForUrl(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "cookiesForUrl"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQNetworkCookieJarFromPointer(NewQNetworkCookieJarFromPointer(nil).__cookiesForUrl_newList())
			for _, v := range (*(*func(*core.QUrl) []*QNetworkCookie)(signal))(core.NewQUrlFromPointer(url)) {
				tmpList.__cookiesForUrl_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQNetworkCookieJarFromPointer(NewQNetworkCookieJarFromPointer(nil).__cookiesForUrl_newList())
		for _, v := range NewQNetworkCookieJarFromPointer(ptr).CookiesForUrlDefault(core.NewQUrlFromPointer(url)) {
			tmpList.__cookiesForUrl_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QNetworkCookieJar) ConnectCookiesForUrl(f func(url *core.QUrl) []*QNetworkCookie) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cookiesForUrl"); signal != nil {
			f := func(url *core.QUrl) []*QNetworkCookie {
				(*(*func(*core.QUrl) []*QNetworkCookie)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "cookiesForUrl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cookiesForUrl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectCookiesForUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cookiesForUrl")
	}
}

func (ptr *QNetworkCookieJar) CookiesForUrl(url core.QUrl_ITF) []*QNetworkCookie {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
			out := make([]*QNetworkCookie, int(l.len))
			tmpList := NewQNetworkCookieJarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__cookiesForUrl_atList(i)
			}
			return out
		}(C.QNetworkCookieJar_CookiesForUrl(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return make([]*QNetworkCookie, 0)
}

func (ptr *QNetworkCookieJar) CookiesForUrlDefault(url core.QUrl_ITF) []*QNetworkCookie {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
			out := make([]*QNetworkCookie, int(l.len))
			tmpList := NewQNetworkCookieJarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__cookiesForUrl_atList(i)
			}
			return out
		}(C.QNetworkCookieJar_CookiesForUrlDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return make([]*QNetworkCookie, 0)
}

//export callbackQNetworkCookieJar_DeleteCookie
func callbackQNetworkCookieJar_DeleteCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "deleteCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QNetworkCookie) bool)(signal))(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).DeleteCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectDeleteCookie(f func(cookie *QNetworkCookie) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deleteCookie"); signal != nil {
			f := func(cookie *QNetworkCookie) bool {
				(*(*func(*QNetworkCookie) bool)(signal))(cookie)
				return f(cookie)
			}
			qt.ConnectSignal(ptr.Pointer(), "deleteCookie", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deleteCookie", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectDeleteCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deleteCookie")
	}
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_DeleteCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DeleteCookieDefault(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_DeleteCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_InsertCookie
func callbackQNetworkCookieJar_InsertCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QNetworkCookie) bool)(signal))(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).InsertCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectInsertCookie(f func(cookie *QNetworkCookie) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insertCookie"); signal != nil {
			f := func(cookie *QNetworkCookie) bool {
				(*(*func(*QNetworkCookie) bool)(signal))(cookie)
				return f(cookie)
			}
			qt.ConnectSignal(ptr.Pointer(), "insertCookie", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insertCookie", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectInsertCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insertCookie")
	}
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_InsertCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookieDefault(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_InsertCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) SetAllCookies(cookieList []*QNetworkCookie) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_SetAllCookies(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkCookieJarFromPointer(NewQNetworkCookieJarFromPointer(nil).__setAllCookies_cookieList_newList())
			for _, v := range cookieList {
				tmpList.__setAllCookies_cookieList_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQNetworkCookieJar_SetCookiesFromUrl
func callbackQNetworkCookieJar_SetCookiesFromUrl(ptr unsafe.Pointer, cookieList C.struct_QtNetwork_PackedList, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "setCookiesFromUrl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func([]*QNetworkCookie, *core.QUrl) bool)(signal))(func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
			out := make([]*QNetworkCookie, int(l.len))
			tmpList := NewQNetworkCookieJarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setCookiesFromUrl_cookieList_atList(i)
			}
			return out
		}(cookieList), core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).SetCookiesFromUrlDefault(func(l C.struct_QtNetwork_PackedList) []*QNetworkCookie {
		out := make([]*QNetworkCookie, int(l.len))
		tmpList := NewQNetworkCookieJarFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__setCookiesFromUrl_cookieList_atList(i)
		}
		return out
	}(cookieList), core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkCookieJar) ConnectSetCookiesFromUrl(f func(cookieList []*QNetworkCookie, url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCookiesFromUrl"); signal != nil {
			f := func(cookieList []*QNetworkCookie, url *core.QUrl) bool {
				(*(*func([]*QNetworkCookie, *core.QUrl) bool)(signal))(cookieList, url)
				return f(cookieList, url)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCookiesFromUrl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCookiesFromUrl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectSetCookiesFromUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCookiesFromUrl")
	}
}

func (ptr *QNetworkCookieJar) SetCookiesFromUrl(cookieList []*QNetworkCookie, url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_SetCookiesFromUrl(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkCookieJarFromPointer(NewQNetworkCookieJarFromPointer(nil).__setCookiesFromUrl_cookieList_newList())
			for _, v := range cookieList {
				tmpList.__setCookiesFromUrl_cookieList_setList(v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQUrl(url))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) SetCookiesFromUrlDefault(cookieList []*QNetworkCookie, url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_SetCookiesFromUrlDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkCookieJarFromPointer(NewQNetworkCookieJarFromPointer(nil).__setCookiesFromUrl_cookieList_newList())
			for _, v := range cookieList {
				tmpList.__setCookiesFromUrl_cookieList_setList(v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQUrl(url))) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_UpdateCookie
func callbackQNetworkCookieJar_UpdateCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "updateCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QNetworkCookie) bool)(signal))(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).UpdateCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectUpdateCookie(f func(cookie *QNetworkCookie) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateCookie"); signal != nil {
			f := func(cookie *QNetworkCookie) bool {
				(*(*func(*QNetworkCookie) bool)(signal))(cookie)
				return f(cookie)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateCookie", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateCookie", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectUpdateCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateCookie")
	}
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_UpdateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookieDefault(cookie QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_UpdateCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_ValidateCookie
func callbackQNetworkCookieJar_ValidateCookie(ptr unsafe.Pointer, cookie unsafe.Pointer, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "validateCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QNetworkCookie, *core.QUrl) bool)(signal))(NewQNetworkCookieFromPointer(cookie), core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).ValidateCookieDefault(NewQNetworkCookieFromPointer(cookie), core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkCookieJar) ConnectValidateCookie(f func(cookie *QNetworkCookie, url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "validateCookie"); signal != nil {
			f := func(cookie *QNetworkCookie, url *core.QUrl) bool {
				(*(*func(*QNetworkCookie, *core.QUrl) bool)(signal))(cookie, url)
				return f(cookie, url)
			}
			qt.ConnectSignal(ptr.Pointer(), "validateCookie", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "validateCookie", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectValidateCookie() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "validateCookie")
	}
}

func (ptr *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_ValidateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) ValidateCookieDefault(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_ValidateCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url))) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_DestroyQNetworkCookieJar
func callbackQNetworkCookieJar_DestroyQNetworkCookieJar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNetworkCookieJar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DestroyQNetworkCookieJarDefault()
	}
}

func (ptr *QNetworkCookieJar) ConnectDestroyQNetworkCookieJar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkCookieJar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkCookieJar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkCookieJar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkCookieJar) DisconnectDestroyQNetworkCookieJar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkCookieJar")
	}
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJarDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkCookieJar_DestroyQNetworkCookieJarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) __allCookies_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookieJar___allCookies_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __allCookies_setList(i QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___allCookies_setList(ptr.Pointer(), PointerFromQNetworkCookie(i))
	}
}

func (ptr *QNetworkCookieJar) __allCookies_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___allCookies_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __cookiesForUrl_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookieJar___cookiesForUrl_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __cookiesForUrl_setList(i QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___cookiesForUrl_setList(ptr.Pointer(), PointerFromQNetworkCookie(i))
	}
}

func (ptr *QNetworkCookieJar) __cookiesForUrl_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___cookiesForUrl_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __setAllCookies_cookieList_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookieJar___setAllCookies_cookieList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __setAllCookies_cookieList_setList(i QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___setAllCookies_cookieList_setList(ptr.Pointer(), PointerFromQNetworkCookie(i))
	}
}

func (ptr *QNetworkCookieJar) __setAllCookies_cookieList_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___setAllCookies_cookieList_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __setCookiesFromUrl_cookieList_atList(i int) *QNetworkCookie {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCookieFromPointer(C.QNetworkCookieJar___setCookiesFromUrl_cookieList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __setCookiesFromUrl_cookieList_setList(i QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___setCookiesFromUrl_cookieList_setList(ptr.Pointer(), PointerFromQNetworkCookie(i))
	}
}

func (ptr *QNetworkCookieJar) __setCookiesFromUrl_cookieList_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___setCookiesFromUrl_cookieList_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkCookieJar___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkCookieJar) __children_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___children_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkCookieJar___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkCookieJar) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkCookieJar___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkCookieJar) __findChildren_newList() unsafe.Pointer {
	return C.QNetworkCookieJar___findChildren_newList(ptr.Pointer())
}

func (ptr *QNetworkCookieJar) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkCookieJar___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookieJar) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkCookieJar) __findChildren_newList3() unsafe.Pointer {
	return C.QNetworkCookieJar___findChildren_newList3(ptr.Pointer())
}

//export callbackQNetworkCookieJar_ChildEvent
func callbackQNetworkCookieJar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkCookieJar_ConnectNotify
func callbackQNetworkCookieJar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkCookieJar_CustomEvent
func callbackQNetworkCookieJar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkCookieJar_DeleteLater
func callbackQNetworkCookieJar_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkCookieJar) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkCookieJar_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNetworkCookieJar_Destroyed
func callbackQNetworkCookieJar_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNetworkCookieJar_DisconnectNotify
func callbackQNetworkCookieJar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkCookieJar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkCookieJar_Event
func callbackQNetworkCookieJar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkCookieJar) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_EventFilter
func callbackQNetworkCookieJar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkCookieJar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkCookieJar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_MetaObject
func callbackQNetworkCookieJar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNetworkCookieJarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkCookieJar) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkCookieJar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNetworkCookieJar_ObjectNameChanged
func callbackQNetworkCookieJar_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNetworkCookieJar_TimerEvent
func callbackQNetworkCookieJar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QNetworkDatagram struct {
	ptr unsafe.Pointer
}

type QNetworkDatagram_ITF interface {
	QNetworkDatagram_PTR() *QNetworkDatagram
}

func (ptr *QNetworkDatagram) QNetworkDatagram_PTR() *QNetworkDatagram {
	return ptr
}

func (ptr *QNetworkDatagram) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkDatagram) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkDatagram(ptr QNetworkDatagram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDatagram_PTR().Pointer()
	}
	return nil
}

func NewQNetworkDatagramFromPointer(ptr unsafe.Pointer) (n *QNetworkDatagram) {
	n = new(QNetworkDatagram)
	n.SetPointer(ptr)
	return
}
func (ptr *QNetworkDatagram) DestroyQNetworkDatagram() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQNetworkDatagram() *QNetworkDatagram {
	tmpValue := NewQNetworkDatagramFromPointer(C.QNetworkDatagram_NewQNetworkDatagram())
	qt.SetFinalizer(tmpValue, (*QNetworkDatagram).DestroyQNetworkDatagram)
	return tmpValue
}

func NewQNetworkDatagram2(data core.QByteArray_ITF, destinationAddress QHostAddress_ITF, port uint16) *QNetworkDatagram {
	tmpValue := NewQNetworkDatagramFromPointer(C.QNetworkDatagram_NewQNetworkDatagram2(core.PointerFromQByteArray(data), PointerFromQHostAddress(destinationAddress), C.ushort(port)))
	qt.SetFinalizer(tmpValue, (*QNetworkDatagram).DestroyQNetworkDatagram)
	return tmpValue
}

func NewQNetworkDatagram3(other QNetworkDatagram_ITF) *QNetworkDatagram {
	tmpValue := NewQNetworkDatagramFromPointer(C.QNetworkDatagram_NewQNetworkDatagram3(PointerFromQNetworkDatagram(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkDatagram).DestroyQNetworkDatagram)
	return tmpValue
}

func (ptr *QNetworkDatagram) Clear() {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDatagram) Data() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkDatagram_Data(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDatagram) DestinationAddress() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QNetworkDatagram_DestinationAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDatagram) DestinationPort() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkDatagram_DestinationPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkDatagram) HopLimit() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkDatagram_HopLimit(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkDatagram) InterfaceIndex() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QNetworkDatagram_InterfaceIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkDatagram) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkDatagram_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkDatagram) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkDatagram_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkDatagram) MakeReply(payload core.QByteArray_ITF) *QNetworkDatagram {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkDatagramFromPointer(C.QNetworkDatagram_MakeReply(ptr.Pointer(), core.PointerFromQByteArray(payload)))
		qt.SetFinalizer(tmpValue, (*QNetworkDatagram).DestroyQNetworkDatagram)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDatagram) MakeReply2(payload core.QByteArray_ITF) *QNetworkDatagram {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkDatagramFromPointer(C.QNetworkDatagram_MakeReply2(ptr.Pointer(), core.PointerFromQByteArray(payload)))
		qt.SetFinalizer(tmpValue, (*QNetworkDatagram).DestroyQNetworkDatagram)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDatagram) SenderAddress() *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QNetworkDatagram_SenderAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDatagram) SenderPort() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkDatagram_SenderPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkDatagram) SetData(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_SetData(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

func (ptr *QNetworkDatagram) SetDestination(address QHostAddress_ITF, port uint16) {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_SetDestination(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port))
	}
}

func (ptr *QNetworkDatagram) SetHopLimit(count int) {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_SetHopLimit(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QNetworkDatagram) SetInterfaceIndex(index uint) {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_SetInterfaceIndex(ptr.Pointer(), C.uint(uint32(index)))
	}
}

func (ptr *QNetworkDatagram) SetSender(address QHostAddress_ITF, port uint16) {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_SetSender(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port))
	}
}

func (ptr *QNetworkDatagram) Swap(other QNetworkDatagram_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkDatagram_Swap(ptr.Pointer(), PointerFromQNetworkDatagram(other))
	}
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

func NewQNetworkDiskCacheFromPointer(ptr unsafe.Pointer) (n *QNetworkDiskCache) {
	n = new(QNetworkDiskCache)
	n.SetPointer(ptr)
	return
}
func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	tmpValue := NewQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	if signal := qt.GetSignal(ptr, "cacheSize"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQNetworkDiskCacheFromPointer(ptr).CacheSizeDefault())
}

func (ptr *QNetworkDiskCache) ConnectCacheSize(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cacheSize"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cacheSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cacheSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectCacheSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cacheSize")
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
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QNetworkDiskCache) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
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
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQIODevice((*(*func(*core.QUrl) *core.QIODevice)(signal))(core.NewQUrlFromPointer(url)))
	}

	return core.PointerFromQIODevice(NewQNetworkDiskCacheFromPointer(ptr).DataDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QNetworkDiskCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			f := func(url *core.QUrl) *core.QIODevice {
				(*(*func(*core.QUrl) *core.QIODevice)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) DataDefault(url core.QUrl_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QNetworkDiskCache_DataDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Expire
func callbackQNetworkDiskCache_Expire(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expire"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQNetworkDiskCacheFromPointer(ptr).ExpireDefault())
}

func (ptr *QNetworkDiskCache) ConnectExpire(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "expire"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "expire", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expire", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectExpire() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "expire")
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
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		tmpValue := NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_FileMetaData(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
		qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Insert
func callbackQNetworkDiskCache_Insert(ptr unsafe.Pointer, device unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "insert"); signal != nil {
		(*(*func(*core.QIODevice))(signal))(core.NewQIODeviceFromPointer(device))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).InsertDefault(core.NewQIODeviceFromPointer(device))
	}
}

func (ptr *QNetworkDiskCache) ConnectInsert(f func(device *core.QIODevice)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insert"); signal != nil {
			f := func(device *core.QIODevice) {
				(*(*func(*core.QIODevice))(signal))(device)
				f(device)
			}
			qt.ConnectSignal(ptr.Pointer(), "insert", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insert", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectInsert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insert")
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
	if signal := qt.GetSignal(ptr, "metaData"); signal != nil {
		return PointerFromQNetworkCacheMetaData((*(*func(*core.QUrl) *QNetworkCacheMetaData)(signal))(core.NewQUrlFromPointer(url)))
	}

	return PointerFromQNetworkCacheMetaData(NewQNetworkDiskCacheFromPointer(ptr).MetaDataDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QNetworkDiskCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metaData"); signal != nil {
			f := func(url *core.QUrl) *QNetworkCacheMetaData {
				(*(*func(*core.QUrl) *QNetworkCacheMetaData)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "metaData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metaData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metaData")
	}
}

func (ptr *QNetworkDiskCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_MetaData(ptr.Pointer(), core.PointerFromQUrl(url)))
		qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) MetaDataDefault(url core.QUrl_ITF) *QNetworkCacheMetaData {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_MetaDataDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
		qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Prepare
func callbackQNetworkDiskCache_Prepare(ptr unsafe.Pointer, metaData unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "prepare"); signal != nil {
		return core.PointerFromQIODevice((*(*func(*QNetworkCacheMetaData) *core.QIODevice)(signal))(NewQNetworkCacheMetaDataFromPointer(metaData)))
	}

	return core.PointerFromQIODevice(NewQNetworkDiskCacheFromPointer(ptr).PrepareDefault(NewQNetworkCacheMetaDataFromPointer(metaData)))
}

func (ptr *QNetworkDiskCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "prepare"); signal != nil {
			f := func(metaData *QNetworkCacheMetaData) *core.QIODevice {
				(*(*func(*QNetworkCacheMetaData) *core.QIODevice)(signal))(metaData)
				return f(metaData)
			}
			qt.ConnectSignal(ptr.Pointer(), "prepare", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "prepare", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectPrepare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "prepare")
	}
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) PrepareDefault(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QNetworkDiskCache_PrepareDefault(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Remove
func callbackQNetworkDiskCache_Remove(ptr unsafe.Pointer, url unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl) bool)(signal))(core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).RemoveDefault(core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkDiskCache) ConnectRemove(f func(url *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			f := func(url *core.QUrl) bool {
				(*(*func(*core.QUrl) bool)(signal))(url)
				return f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "remove", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectRemove() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *QNetworkDiskCache) Remove(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkDiskCache_Remove(ptr.Pointer(), core.PointerFromQUrl(url))) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) RemoveDefault(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkDiskCache_RemoveDefault(ptr.Pointer(), core.PointerFromQUrl(url))) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	if ptr.Pointer() != nil {
		var cacheDirC *C.char
		if cacheDir != "" {
			cacheDirC = C.CString(cacheDir)
			defer C.free(unsafe.Pointer(cacheDirC))
		}
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: cacheDirC, len: C.longlong(len(cacheDir))})
	}
}

func (ptr *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetMaximumCacheSize(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQNetworkDiskCache_UpdateMetaData
func callbackQNetworkDiskCache_UpdateMetaData(ptr unsafe.Pointer, metaData unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMetaData"); signal != nil {
		(*(*func(*QNetworkCacheMetaData))(signal))(NewQNetworkCacheMetaDataFromPointer(metaData))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).UpdateMetaDataDefault(NewQNetworkCacheMetaDataFromPointer(metaData))
	}
}

func (ptr *QNetworkDiskCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateMetaData"); signal != nil {
			f := func(metaData *QNetworkCacheMetaData) {
				(*(*func(*QNetworkCacheMetaData))(signal))(metaData)
				f(metaData)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateMetaData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateMetaData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectUpdateMetaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateMetaData")
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

//export callbackQNetworkDiskCache_DestroyQNetworkDiskCache
func callbackQNetworkDiskCache_DestroyQNetworkDiskCache(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNetworkDiskCache"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).DestroyQNetworkDiskCacheDefault()
	}
}

func (ptr *QNetworkDiskCache) ConnectDestroyQNetworkDiskCache(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkDiskCache"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkDiskCache", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkDiskCache", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkDiskCache) DisconnectDestroyQNetworkDiskCache() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkDiskCache")
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCacheDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkDiskCache_DestroyQNetworkDiskCacheDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQNetworkInterfaceFromPointer(ptr unsafe.Pointer) (n *QNetworkInterface) {
	n = new(QNetworkInterface)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNetworkInterface__InterfaceType
//QNetworkInterface::InterfaceType
type QNetworkInterface__InterfaceType int64

const (
	QNetworkInterface__Loopback   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(1)
	QNetworkInterface__Virtual    QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(2)
	QNetworkInterface__Ethernet   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(3)
	QNetworkInterface__Slip       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(4)
	QNetworkInterface__CanBus     QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(5)
	QNetworkInterface__Ppp        QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(6)
	QNetworkInterface__Fddi       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(7)
	QNetworkInterface__Wifi       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(8)
	QNetworkInterface__Ieee80211  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(QNetworkInterface__Wifi)
	QNetworkInterface__Phonet     QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(9)
	QNetworkInterface__Ieee802154 QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(10)
	QNetworkInterface__SixLoWPAN  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(11)
	QNetworkInterface__Ieee80216  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(12)
	QNetworkInterface__Ieee1394   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(13)
	QNetworkInterface__Unknown    QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(0)
)

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
	tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface())
	qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {
	tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface2(PointerFromQNetworkInterface(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) AddressEntries() []*QNetworkAddressEntry {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkAddressEntry {
			out := make([]*QNetworkAddressEntry, int(l.len))
			tmpList := NewQNetworkInterfaceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addressEntries_atList(i)
			}
			return out
		}(C.QNetworkInterface_AddressEntries(ptr.Pointer()))
	}
	return make([]*QNetworkAddressEntry, 0)
}

func QNetworkInterface_AllAddresses() []*QHostAddress {
	return func(l C.struct_QtNetwork_PackedList) []*QHostAddress {
		out := make([]*QHostAddress, int(l.len))
		tmpList := NewQNetworkInterfaceFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__allAddresses_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllAddresses())
}

func (ptr *QNetworkInterface) AllAddresses() []*QHostAddress {
	return func(l C.struct_QtNetwork_PackedList) []*QHostAddress {
		out := make([]*QHostAddress, int(l.len))
		tmpList := NewQNetworkInterfaceFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__allAddresses_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllAddresses())
}

func QNetworkInterface_AllInterfaces() []*QNetworkInterface {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkInterface {
		out := make([]*QNetworkInterface, int(l.len))
		tmpList := NewQNetworkInterfaceFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__allInterfaces_atList(i)
		}
		return out
	}(C.QNetworkInterface_QNetworkInterface_AllInterfaces())
}

func (ptr *QNetworkInterface) AllInterfaces() []*QNetworkInterface {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkInterface {
		out := make([]*QNetworkInterface, int(l.len))
		tmpList := NewQNetworkInterfaceFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__allInterfaces_atList(i)
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
	tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromIndex(C.int(int32(index))))
	qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) InterfaceFromIndex(index int) *QNetworkInterface {
	tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromIndex(C.int(int32(index))))
	qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func QNetworkInterface_InterfaceFromName(name string) *QNetworkInterface {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) InterfaceFromName(name string) *QNetworkInterface {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func QNetworkInterface_InterfaceIndexFromName(name string) int {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int(int32(C.QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})))
}

func (ptr *QNetworkInterface) InterfaceIndexFromName(name string) int {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int(int32(C.QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))})))
}

func QNetworkInterface_InterfaceNameFromIndex(index int) string {
	return cGoUnpackString(C.QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(C.int(int32(index))))
}

func (ptr *QNetworkInterface) InterfaceNameFromIndex(index int) string {
	return cGoUnpackString(C.QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(C.int(int32(index))))
}

func (ptr *QNetworkInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkInterface_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkInterface) MaximumTransmissionUnit() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkInterface_MaximumTransmissionUnit(ptr.Pointer())))
	}
	return 0
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

func (ptr *QNetworkInterface) Type() QNetworkInterface__InterfaceType {
	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceType(C.QNetworkInterface_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkInterface_DestroyQNetworkInterface(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkInterface) __addressEntries_atList(i int) *QNetworkAddressEntry {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkAddressEntryFromPointer(C.QNetworkInterface___addressEntries_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkInterface) __addressEntries_setList(i QNetworkAddressEntry_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkInterface___addressEntries_setList(ptr.Pointer(), PointerFromQNetworkAddressEntry(i))
	}
}

func (ptr *QNetworkInterface) __addressEntries_newList() unsafe.Pointer {
	return C.QNetworkInterface___addressEntries_newList(ptr.Pointer())
}

func (ptr *QNetworkInterface) __allAddresses_atList(i int) *QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQHostAddressFromPointer(C.QNetworkInterface___allAddresses_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkInterface) __allAddresses_setList(i QHostAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkInterface___allAddresses_setList(ptr.Pointer(), PointerFromQHostAddress(i))
	}
}

func (ptr *QNetworkInterface) __allAddresses_newList() unsafe.Pointer {
	return C.QNetworkInterface___allAddresses_newList(ptr.Pointer())
}

func (ptr *QNetworkInterface) __allInterfaces_atList(i int) *QNetworkInterface {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkInterface___allInterfaces_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkInterface) __allInterfaces_setList(i QNetworkInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkInterface___allInterfaces_setList(ptr.Pointer(), PointerFromQNetworkInterface(i))
	}
}

func (ptr *QNetworkInterface) __allInterfaces_newList() unsafe.Pointer {
	return C.QNetworkInterface___allInterfaces_newList(ptr.Pointer())
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

func NewQNetworkProxyFromPointer(ptr unsafe.Pointer) (n *QNetworkProxy) {
	n = new(QNetworkProxy)
	n.SetPointer(ptr)
	return
}

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

//go:generate stringer -type=QNetworkProxy__Capability
//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   QNetworkProxy__Capability = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        QNetworkProxy__Capability = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = QNetworkProxy__Capability(0x0010)
	QNetworkProxy__SctpTunnelingCapability  QNetworkProxy__Capability = QNetworkProxy__Capability(0x00020)
	QNetworkProxy__SctpListeningCapability  QNetworkProxy__Capability = QNetworkProxy__Capability(0x00040)
)

func NewQNetworkProxy() *QNetworkProxy {
	tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy())
	qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func NewQNetworkProxy2(ty QNetworkProxy__ProxyType, hostName string, port uint16, user string, password string) *QNetworkProxy {
	var hostNameC *C.char
	if hostName != "" {
		hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
	}
	var userC *C.char
	if user != "" {
		userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
	}
	var passwordC *C.char
	if password != "" {
		passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
	}
	tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy2(C.longlong(ty), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), C.struct_QtNetwork_PackedString{data: userC, len: C.longlong(len(user))}, C.struct_QtNetwork_PackedString{data: passwordC, len: C.longlong(len(password))}))
	qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func NewQNetworkProxy3(other QNetworkProxy_ITF) *QNetworkProxy {
	tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy3(PointerFromQNetworkProxy(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func QNetworkProxy_ApplicationProxy() *QNetworkProxy {
	tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxy_QNetworkProxy_ApplicationProxy())
	qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func (ptr *QNetworkProxy) ApplicationProxy() *QNetworkProxy {
	tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxy_QNetworkProxy_ApplicationProxy())
	qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
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
		return int8(C.QNetworkProxy_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QNetworkProxy_Header(ptr.Pointer(), C.longlong(header)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
		return int8(C.QNetworkProxy_IsCachingProxy(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkProxy_IsTransparentProxy(ptr.Pointer())) != 0
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

func (ptr *QNetworkProxy) QNetworkCacheMetaData_RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkProxy_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxy) RawHeaderList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQNetworkProxyFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__rawHeaderList_atList(i)
			}
			return out
		}(C.QNetworkProxy_RawHeaderList(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
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
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QNetworkProxy_SetHostName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))})
	}
}

func (ptr *QNetworkProxy) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.QNetworkProxy_SetPassword(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: passwordC, len: C.longlong(len(password))})
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
		var userC *C.char
		if user != "" {
			userC = C.CString(user)
			defer C.free(unsafe.Pointer(userC))
		}
		C.QNetworkProxy_SetUser(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: userC, len: C.longlong(len(user))})
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

		qt.SetFinalizer(ptr, nil)
		C.QNetworkProxy_DestroyQNetworkProxy(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkProxy) __rawHeaderList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkProxy___rawHeaderList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxy) __rawHeaderList_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy___rawHeaderList_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkProxy) __rawHeaderList_newList() unsafe.Pointer {
	return C.QNetworkProxy___rawHeaderList_newList(ptr.Pointer())
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

func NewQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) (n *QNetworkProxyFactory) {
	n = new(QNetworkProxyFactory)
	n.SetPointer(ptr)
	return
}
func NewQNetworkProxyFactory() *QNetworkProxyFactory {
	return NewQNetworkProxyFactoryFromPointer(C.QNetworkProxyFactory_NewQNetworkProxyFactory())
}

func QNetworkProxyFactory_ProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		out := make([]*QNetworkProxy, int(l.len))
		tmpList := NewQNetworkProxyFactoryFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__proxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_ProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func (ptr *QNetworkProxyFactory) ProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		out := make([]*QNetworkProxy, int(l.len))
		tmpList := NewQNetworkProxyFactoryFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__proxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_ProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

//export callbackQNetworkProxyFactory_QueryProxy
func callbackQNetworkProxyFactory_QueryProxy(ptr unsafe.Pointer, query unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "queryProxy"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQNetworkProxyFactoryFromPointer(NewQNetworkProxyFactoryFromPointer(nil).__queryProxy_newList())
			for _, v := range (*(*func(*QNetworkProxyQuery) []*QNetworkProxy)(signal))(NewQNetworkProxyQueryFromPointer(query)) {
				tmpList.__queryProxy_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQNetworkProxyFactoryFromPointer(NewQNetworkProxyFactoryFromPointer(nil).__queryProxy_newList())
		for _, v := range make([]*QNetworkProxy, 0) {
			tmpList.__queryProxy_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QNetworkProxyFactory) ConnectQueryProxy(f func(query *QNetworkProxyQuery) []*QNetworkProxy) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "queryProxy"); signal != nil {
			f := func(query *QNetworkProxyQuery) []*QNetworkProxy {
				(*(*func(*QNetworkProxyQuery) []*QNetworkProxy)(signal))(query)
				return f(query)
			}
			qt.ConnectSignal(ptr.Pointer(), "queryProxy", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "queryProxy", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkProxyFactory) DisconnectQueryProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "queryProxy")
	}
}

func (ptr *QNetworkProxyFactory) QueryProxy(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
			out := make([]*QNetworkProxy, int(l.len))
			tmpList := NewQNetworkProxyFactoryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__queryProxy_atList(i)
			}
			return out
		}(C.QNetworkProxyFactory_QueryProxy(ptr.Pointer(), PointerFromQNetworkProxyQuery(query)))
	}
	return make([]*QNetworkProxy, 0)
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

func QNetworkProxyFactory_SystemProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		out := make([]*QNetworkProxy, int(l.len))
		tmpList := NewQNetworkProxyFactoryFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__systemProxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_SystemProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func (ptr *QNetworkProxyFactory) SystemProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {
	return func(l C.struct_QtNetwork_PackedList) []*QNetworkProxy {
		out := make([]*QNetworkProxy, int(l.len))
		tmpList := NewQNetworkProxyFactoryFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__systemProxyForQuery_atList(i)
		}
		return out
	}(C.QNetworkProxyFactory_QNetworkProxyFactory_SystemProxyForQuery(PointerFromQNetworkProxyQuery(query)))
}

func QNetworkProxyFactory_UsesSystemConfiguration() bool {
	return int8(C.QNetworkProxyFactory_QNetworkProxyFactory_UsesSystemConfiguration()) != 0
}

func (ptr *QNetworkProxyFactory) UsesSystemConfiguration() bool {
	return int8(C.QNetworkProxyFactory_QNetworkProxyFactory_UsesSystemConfiguration()) != 0
}

//export callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory
func callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNetworkProxyFactory"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkProxyFactoryFromPointer(ptr).DestroyQNetworkProxyFactoryDefault()
	}
}

func (ptr *QNetworkProxyFactory) ConnectDestroyQNetworkProxyFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkProxyFactory"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkProxyFactory", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkProxyFactory", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkProxyFactory) DisconnectDestroyQNetworkProxyFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkProxyFactory")
	}
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactoryDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactoryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkProxyFactory) __proxyForQuery_atList(i int) *QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxyFactory___proxyForQuery_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyFactory) __proxyForQuery_setList(i QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory___proxyForQuery_setList(ptr.Pointer(), PointerFromQNetworkProxy(i))
	}
}

func (ptr *QNetworkProxyFactory) __proxyForQuery_newList() unsafe.Pointer {
	return C.QNetworkProxyFactory___proxyForQuery_newList(ptr.Pointer())
}

func (ptr *QNetworkProxyFactory) __queryProxy_atList(i int) *QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxyFactory___queryProxy_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyFactory) __queryProxy_setList(i QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory___queryProxy_setList(ptr.Pointer(), PointerFromQNetworkProxy(i))
	}
}

func (ptr *QNetworkProxyFactory) __queryProxy_newList() unsafe.Pointer {
	return C.QNetworkProxyFactory___queryProxy_newList(ptr.Pointer())
}

func (ptr *QNetworkProxyFactory) __systemProxyForQuery_atList(i int) *QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkProxyFromPointer(C.QNetworkProxyFactory___systemProxyForQuery_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyFactory) __systemProxyForQuery_setList(i QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory___systemProxyForQuery_setList(ptr.Pointer(), PointerFromQNetworkProxy(i))
	}
}

func (ptr *QNetworkProxyFactory) __systemProxyForQuery_newList() unsafe.Pointer {
	return C.QNetworkProxyFactory___systemProxyForQuery_newList(ptr.Pointer())
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

func NewQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) (n *QNetworkProxyQuery) {
	n = new(QNetworkProxyQuery)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNetworkProxyQuery__QueryType
//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__SctpSocket QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(2)
	QNetworkProxyQuery__TcpServer  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(101)
	QNetworkProxyQuery__SctpServer QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(102)
)

func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	tmpValue := NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery())
	qt.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery2(requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	tmpValue := NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery2(core.PointerFromQUrl(requestUrl), C.longlong(queryType)))
	qt.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var hostnameC *C.char
	if hostname != "" {
		hostnameC = C.CString(hostname)
		defer C.free(unsafe.Pointer(hostnameC))
	}
	var protocolTagC *C.char
	if protocolTag != "" {
		protocolTagC = C.CString(protocolTag)
		defer C.free(unsafe.Pointer(protocolTagC))
	}
	tmpValue := NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery3(C.struct_QtNetwork_PackedString{data: hostnameC, len: C.longlong(len(hostname))}, C.int(int32(port)), C.struct_QtNetwork_PackedString{data: protocolTagC, len: C.longlong(len(protocolTag))}, C.longlong(queryType)))
	qt.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery4(bindPort uint16, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	var protocolTagC *C.char
	if protocolTag != "" {
		protocolTagC = C.CString(protocolTag)
		defer C.free(unsafe.Pointer(protocolTagC))
	}
	tmpValue := NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery4(C.ushort(bindPort), C.struct_QtNetwork_PackedString{data: protocolTagC, len: C.longlong(len(protocolTag))}, C.longlong(queryType)))
	qt.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery8(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {
	tmpValue := NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery8(PointerFromQNetworkProxyQuery(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func (ptr *QNetworkProxyQuery) LocalPort() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkProxyQuery_LocalPort(ptr.Pointer())))
	}
	return 0
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

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	if ptr.Pointer() != nil {
		var hostnameC *C.char
		if hostname != "" {
			hostnameC = C.CString(hostname)
			defer C.free(unsafe.Pointer(hostnameC))
		}
		C.QNetworkProxyQuery_SetPeerHostName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostnameC, len: C.longlong(len(hostname))})
	}
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	if ptr.Pointer() != nil {
		var protocolTagC *C.char
		if protocolTag != "" {
			protocolTagC = C.CString(protocolTag)
			defer C.free(unsafe.Pointer(protocolTagC))
		}
		C.QNetworkProxyQuery_SetProtocolTag(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: protocolTagC, len: C.longlong(len(protocolTag))})
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
		tmpValue := core.NewQUrlFromPointer(C.QNetworkProxyQuery_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkProxyQuery_DestroyQNetworkProxyQuery(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQNetworkReplyFromPointer(ptr unsafe.Pointer) (n *QNetworkReply) {
	n = new(QNetworkReply)
	n.SetPointer(ptr)
	return
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

func NewQNetworkReply(parent core.QObject_ITF) *QNetworkReply {
	tmpValue := NewQNetworkReplyFromPointer(C.QNetworkReply_NewQNetworkReply(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNetworkReply_Abort
func callbackQNetworkReply_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "abort"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkReply) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "abort"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "abort")
	}
}

func (ptr *QNetworkReply) Abort() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Abort(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Attribute(code QNetworkRequest__Attribute) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QNetworkReply_Attribute(ptr.Pointer(), C.longlong(code)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_Close
func callbackQNetworkReply_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkReplyFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QNetworkReply) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_CloseDefault(ptr.Pointer())
	}
}

//export callbackQNetworkReply_DownloadProgress
func callbackQNetworkReply_DownloadProgress(ptr unsafe.Pointer, bytesReceived C.longlong, bytesTotal C.longlong) {
	if signal := qt.GetSignal(ptr, "downloadProgress"); signal != nil {
		(*(*func(int64, int64))(signal))(int64(bytesReceived), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "downloadProgress") {
			C.QNetworkReply_ConnectDownloadProgress(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "downloadProgress")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "downloadProgress"); signal != nil {
			f := func(bytesReceived int64, bytesTotal int64) {
				(*(*func(int64, int64))(signal))(bytesReceived, bytesTotal)
				f(bytesReceived, bytesTotal)
			}
			qt.ConnectSignal(ptr.Pointer(), "downloadProgress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "downloadProgress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectDownloadProgress() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectDownloadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "downloadProgress")
	}
}

func (ptr *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DownloadProgress(ptr.Pointer(), C.longlong(bytesReceived), C.longlong(bytesTotal))
	}
}

//export callbackQNetworkReply_Encrypted
func callbackQNetworkReply_Encrypted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "encrypted"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkReply) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "encrypted") {
			C.QNetworkReply_ConnectEncrypted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "encrypted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "encrypted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "encrypted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encrypted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "encrypted")
	}
}

func (ptr *QNetworkReply) Encrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Encrypted(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Error() QNetworkReply__NetworkError {
	if ptr.Pointer() != nil {
		return QNetworkReply__NetworkError(C.QNetworkReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_Error2
func callbackQNetworkReply_Error2(ptr unsafe.Pointer, code C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QNetworkReply__NetworkError))(signal))(QNetworkReply__NetworkError(code))
	}

}

func (ptr *QNetworkReply) ConnectError2(f func(code QNetworkReply__NetworkError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QNetworkReply_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(code QNetworkReply__NetworkError) {
				(*(*func(QNetworkReply__NetworkError))(signal))(code)
				f(code)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QNetworkReply) Error2(code QNetworkReply__NetworkError) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Error2(ptr.Pointer(), C.longlong(code))
	}
}

//export callbackQNetworkReply_Finished
func callbackQNetworkReply_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QNetworkReply_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QNetworkReply) Finished() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Finished(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) HasRawHeader(headerName core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QNetworkReply_Header(ptr.Pointer(), C.longlong(header)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_IgnoreSslErrors
func callbackQNetworkReply_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ignoreSslErrors"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkReplyFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignoreSslErrors"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignoreSslErrors")
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

func (ptr *QNetworkReply) IgnoreSslErrors2(errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrors2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkReplyFromPointer(NewQNetworkReplyFromPointer(nil).__ignoreSslErrors_errors_newList2())
			for _, v := range errors {
				tmpList.__ignoreSslErrors_errors_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQNetworkReply_IgnoreSslErrorsImplementation
func callbackQNetworkReply_IgnoreSslErrorsImplementation(ptr unsafe.Pointer, errors C.struct_QtNetwork_PackedList) {
	if signal := qt.GetSignal(ptr, "ignoreSslErrorsImplementation"); signal != nil {
		(*(*func([]*QSslError))(signal))(func(l C.struct_QtNetwork_PackedList) []*QSslError {
			out := make([]*QSslError, int(l.len))
			tmpList := NewQNetworkReplyFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__ignoreSslErrorsImplementation_errors_atList(i)
			}
			return out
		}(errors))
	} else {
		NewQNetworkReplyFromPointer(ptr).IgnoreSslErrorsImplementationDefault(func(l C.struct_QtNetwork_PackedList) []*QSslError {
			out := make([]*QSslError, int(l.len))
			tmpList := NewQNetworkReplyFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__ignoreSslErrorsImplementation_errors_atList(i)
			}
			return out
		}(errors))
	}
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrorsImplementation(f func(errors []*QSslError)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignoreSslErrorsImplementation"); signal != nil {
			f := func(errors []*QSslError) {
				(*(*func([]*QSslError))(signal))(errors)
				f(errors)
			}
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrorsImplementation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrorsImplementation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrorsImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignoreSslErrorsImplementation")
	}
}

func (ptr *QNetworkReply) IgnoreSslErrorsImplementation(errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrorsImplementation(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkReplyFromPointer(NewQNetworkReplyFromPointer(nil).__ignoreSslErrorsImplementation_errors_newList())
			for _, v := range errors {
				tmpList.__ignoreSslErrorsImplementation_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QNetworkReply) IgnoreSslErrorsImplementationDefault(errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrorsImplementationDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkReplyFromPointer(NewQNetworkReplyFromPointer(nil).__ignoreSslErrorsImplementation_errors_newList())
			for _, v := range errors {
				tmpList.__ignoreSslErrorsImplementation_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QNetworkReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_IsFinished(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsRunning() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_IsRunning(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkReply) Manager() *QNetworkAccessManager {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkAccessManagerFromPointer(C.QNetworkReply_Manager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_MetaDataChanged
func callbackQNetworkReply_MetaDataChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "metaDataChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkReply) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "metaDataChanged") {
			C.QNetworkReply_ConnectMetaDataChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "metaDataChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "metaDataChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "metaDataChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metaDataChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "metaDataChanged")
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
	if signal := qt.GetSignal(ptr, "preSharedKeyAuthenticationRequired"); signal != nil {
		(*(*func(*QSslPreSharedKeyAuthenticator))(signal))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkReply) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preSharedKeyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			f := func(authenticator *QSslPreSharedKeyAuthenticator) {
				(*(*func(*QSslPreSharedKeyAuthenticator))(signal))(authenticator)
				f(authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkReply) QNetworkCacheMetaData_RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkReply_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) RawHeaderList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQNetworkReplyFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__rawHeaderList_atList(i)
			}
			return out
		}(C.QNetworkReply_RawHeaderList(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QNetworkReply) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_RedirectAllowed
func callbackQNetworkReply_RedirectAllowed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "redirectAllowed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkReply) ConnectRedirectAllowed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "redirectAllowed") {
			C.QNetworkReply_ConnectRedirectAllowed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "redirectAllowed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "redirectAllowed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "redirectAllowed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "redirectAllowed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectRedirectAllowed() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectRedirectAllowed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "redirectAllowed")
	}
}

func (ptr *QNetworkReply) RedirectAllowed() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_RedirectAllowed(ptr.Pointer())
	}
}

//export callbackQNetworkReply_Redirected
func callbackQNetworkReply_Redirected(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "redirected"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QNetworkReply) ConnectRedirected(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "redirected") {
			C.QNetworkReply_ConnectRedirected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "redirected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "redirected"); signal != nil {
			f := func(url *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "redirected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "redirected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectRedirected() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectRedirected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "redirected")
	}
}

func (ptr *QNetworkReply) Redirected(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Redirected(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkReply) Request() *QNetworkRequest {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkRequestFromPointer(C.QNetworkReply_Request(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
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
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		C.QNetworkReply_SetError(ptr.Pointer(), C.longlong(errorCode), C.struct_QtNetwork_PackedString{data: errorStringC, len: C.longlong(len(errorString))})
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
	if signal := qt.GetSignal(ptr, "setReadBufferSize"); signal != nil {
		(*(*func(int64))(signal))(int64(size))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QNetworkReply) ConnectSetReadBufferSize(f func(size int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setReadBufferSize"); signal != nil {
			f := func(size int64) {
				(*(*func(int64))(signal))(size)
				f(size)
			}
			qt.ConnectSignal(ptr.Pointer(), "setReadBufferSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setReadBufferSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectSetReadBufferSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setReadBufferSize")
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

//export callbackQNetworkReply_SetSslConfigurationImplementation
func callbackQNetworkReply_SetSslConfigurationImplementation(ptr unsafe.Pointer, configuration unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSslConfigurationImplementation"); signal != nil {
		(*(*func(*QSslConfiguration))(signal))(NewQSslConfigurationFromPointer(configuration))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetSslConfigurationImplementationDefault(NewQSslConfigurationFromPointer(configuration))
	}
}

func (ptr *QNetworkReply) ConnectSetSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSslConfigurationImplementation"); signal != nil {
			f := func(configuration *QSslConfiguration) {
				(*(*func(*QSslConfiguration))(signal))(configuration)
				f(configuration)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSslConfigurationImplementation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSslConfigurationImplementation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectSetSslConfigurationImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSslConfigurationImplementation")
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

func (ptr *QNetworkReply) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkReply) SslConfiguration() *QSslConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslConfigurationFromPointer(C.QNetworkReply_SslConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_SslConfigurationImplementation
func callbackQNetworkReply_SslConfigurationImplementation(ptr unsafe.Pointer, configuration unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sslConfigurationImplementation"); signal != nil {
		(*(*func(*QSslConfiguration))(signal))(NewQSslConfigurationFromPointer(configuration))
	} else {
		NewQNetworkReplyFromPointer(ptr).SslConfigurationImplementationDefault(NewQSslConfigurationFromPointer(configuration))
	}
}

func (ptr *QNetworkReply) ConnectSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sslConfigurationImplementation"); signal != nil {
			f := func(configuration *QSslConfiguration) {
				(*(*func(*QSslConfiguration))(signal))(configuration)
				f(configuration)
			}
			qt.ConnectSignal(ptr.Pointer(), "sslConfigurationImplementation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslConfigurationImplementation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectSslConfigurationImplementation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sslConfigurationImplementation")
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

//export callbackQNetworkReply_SslErrors
func callbackQNetworkReply_SslErrors(ptr unsafe.Pointer, errors C.struct_QtNetwork_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors"); signal != nil {
		(*(*func([]*QSslError))(signal))(func(l C.struct_QtNetwork_PackedList) []*QSslError {
			out := make([]*QSslError, int(l.len))
			tmpList := NewQNetworkReplyFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sslErrors_errors_atList(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QNetworkReply) ConnectSslErrors(f func(errors []*QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors") {
			C.QNetworkReply_ConnectSslErrors(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sslErrors")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors"); signal != nil {
			f := func(errors []*QSslError) {
				(*(*func([]*QSslError))(signal))(errors)
				f(errors)
			}
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectSslErrors() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectSslErrors(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sslErrors")
	}
}

func (ptr *QNetworkReply) SslErrors(errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SslErrors(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNetworkReplyFromPointer(NewQNetworkReplyFromPointer(nil).__sslErrors_errors_newList())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQNetworkReply_UploadProgress
func callbackQNetworkReply_UploadProgress(ptr unsafe.Pointer, bytesSent C.longlong, bytesTotal C.longlong) {
	if signal := qt.GetSignal(ptr, "uploadProgress"); signal != nil {
		(*(*func(int64, int64))(signal))(int64(bytesSent), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) ConnectUploadProgress(f func(bytesSent int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "uploadProgress") {
			C.QNetworkReply_ConnectUploadProgress(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "uploadProgress")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "uploadProgress"); signal != nil {
			f := func(bytesSent int64, bytesTotal int64) {
				(*(*func(int64, int64))(signal))(bytesSent, bytesTotal)
				f(bytesSent, bytesTotal)
			}
			qt.ConnectSignal(ptr.Pointer(), "uploadProgress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "uploadProgress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectUploadProgress() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectUploadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "uploadProgress")
	}
}

func (ptr *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_UploadProgress(ptr.Pointer(), C.longlong(bytesSent), C.longlong(bytesTotal))
	}
}

func (ptr *QNetworkReply) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QNetworkReply_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_DestroyQNetworkReply
func callbackQNetworkReply_DestroyQNetworkReply(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNetworkReply"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkReplyFromPointer(ptr).DestroyQNetworkReplyDefault()
	}
}

func (ptr *QNetworkReply) ConnectDestroyQNetworkReply(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkReply"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkReply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkReply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkReply) DisconnectDestroyQNetworkReply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkReply")
	}
}

func (ptr *QNetworkReply) DestroyQNetworkReply() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkReply_DestroyQNetworkReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkReply) DestroyQNetworkReplyDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkReply_DestroyQNetworkReplyDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkReply) __ignoreSslErrors_errors_atList2(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QNetworkReply___ignoreSslErrors_errors_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __ignoreSslErrors_errors_setList2(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___ignoreSslErrors_errors_setList2(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QNetworkReply) __ignoreSslErrors_errors_newList2() unsafe.Pointer {
	return C.QNetworkReply___ignoreSslErrors_errors_newList2(ptr.Pointer())
}

func (ptr *QNetworkReply) __ignoreSslErrorsImplementation_errors_atList(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QNetworkReply___ignoreSslErrorsImplementation_errors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __ignoreSslErrorsImplementation_errors_setList(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___ignoreSslErrorsImplementation_errors_setList(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QNetworkReply) __ignoreSslErrorsImplementation_errors_newList() unsafe.Pointer {
	return C.QNetworkReply___ignoreSslErrorsImplementation_errors_newList(ptr.Pointer())
}

func (ptr *QNetworkReply) __rawHeaderList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkReply___rawHeaderList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __rawHeaderList_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___rawHeaderList_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkReply) __rawHeaderList_newList() unsafe.Pointer {
	return C.QNetworkReply___rawHeaderList_newList(ptr.Pointer())
}

func (ptr *QNetworkReply) __sslErrors_errors_atList(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QNetworkReply___sslErrors_errors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __sslErrors_errors_setList(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___sslErrors_errors_setList(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QNetworkReply) __sslErrors_errors_newList() unsafe.Pointer {
	return C.QNetworkReply___sslErrors_errors_newList(ptr.Pointer())
}

func (ptr *QNetworkReply) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkReply___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkReply) __children_newList() unsafe.Pointer {
	return C.QNetworkReply___children_newList(ptr.Pointer())
}

func (ptr *QNetworkReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkReply___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkReply) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNetworkReply___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNetworkReply) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkReply___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkReply) __findChildren_newList() unsafe.Pointer {
	return C.QNetworkReply___findChildren_newList(ptr.Pointer())
}

func (ptr *QNetworkReply) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkReply___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkReply) __findChildren_newList3() unsafe.Pointer {
	return C.QNetworkReply___findChildren_newList3(ptr.Pointer())
}

//export callbackQNetworkReply_AboutToClose
func callbackQNetworkReply_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQNetworkReply_AtEnd
func callbackQNetworkReply_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).AtEndDefault())))
}

func (ptr *QNetworkReply) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_AtEndDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNetworkReply_BytesAvailable
func callbackQNetworkReply_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QNetworkReply) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_BytesToWrite
func callbackQNetworkReply_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesToWrite"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QNetworkReply) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_BytesWritten
func callbackQNetworkReply_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(bytes))
	}

}

//export callbackQNetworkReply_CanReadLine
func callbackQNetworkReply_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QNetworkReply) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_CanReadLineDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNetworkReply_ChannelBytesWritten
func callbackQNetworkReply_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		(*(*func(int, int64))(signal))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQNetworkReply_ChannelReadyRead
func callbackQNetworkReply_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		(*(*func(int))(signal))(int(int32(channel)))
	}

}

//export callbackQNetworkReply_IsSequential
func callbackQNetworkReply_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QNetworkReply) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_IsSequentialDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNetworkReply_Open
func callbackQNetworkReply_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(core.QIODevice__OpenModeFlag) bool)(signal))(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QNetworkReply) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_OpenDefault(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQNetworkReply_Pos
func callbackQNetworkReply_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).PosDefault())
}

func (ptr *QNetworkReply) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_ReadChannelFinished
func callbackQNetworkReply_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQNetworkReply_ReadData
func callbackQNetworkReply_ReadData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQNetworkReplyFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QNetworkReply) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QNetworkReply_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QNetworkReply) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QNetworkReply_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQNetworkReply_ReadLineData
func callbackQNetworkReply_ReadLineData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readLineData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).ReadLineDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QNetworkReply) ReadLineDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QNetworkReply_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQNetworkReply_ReadyRead
func callbackQNetworkReply_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQNetworkReply_Reset
func callbackQNetworkReply_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).ResetDefault())))
}

func (ptr *QNetworkReply) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_ResetDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNetworkReply_Seek
func callbackQNetworkReply_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int64) bool)(signal))(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QNetworkReply) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_SeekDefault(ptr.Pointer(), C.longlong(pos))) != 0
	}
	return false
}

//export callbackQNetworkReply_Size
func callbackQNetworkReply_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).SizeDefault())
}

func (ptr *QNetworkReply) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_WaitForBytesWritten
func callbackQNetworkReply_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QNetworkReply) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQNetworkReply_WaitForReadyRead
func callbackQNetworkReply_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QNetworkReply) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQNetworkReply_WriteData
func callbackQNetworkReply_WriteData(ptr unsafe.Pointer, data C.struct_QtNetwork_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QNetworkReply) WriteData(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QNetworkReply_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QNetworkReply) WriteDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QNetworkReply_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQNetworkReply_ChildEvent
func callbackQNetworkReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkReply_ConnectNotify
func callbackQNetworkReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkReply_CustomEvent
func callbackQNetworkReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkReply_DeleteLater
func callbackQNetworkReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkReply_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNetworkReply_Destroyed
func callbackQNetworkReply_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNetworkReply_DisconnectNotify
func callbackQNetworkReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkReply_Event
func callbackQNetworkReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNetworkReply_EventFilter
func callbackQNetworkReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNetworkReply_MetaObject
func callbackQNetworkReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNetworkReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkReply) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNetworkReply_ObjectNameChanged
func callbackQNetworkReply_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNetworkReply_TimerEvent
func callbackQNetworkReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQNetworkRequestFromPointer(ptr unsafe.Pointer) (n *QNetworkRequest) {
	n = new(QNetworkRequest)
	n.SetPointer(ptr)
	return
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
	QNetworkRequest__HTTP2AllowedAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(22)
	QNetworkRequest__HTTP2WasUsedAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(23)
	QNetworkRequest__OriginalContentLengthAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(24)
	QNetworkRequest__RedirectPolicyAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(25)
	QNetworkRequest__Http2DirectAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(26)
	QNetworkRequest__ResourceTypeAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(27)
	QNetworkRequest__User                                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(1000)
	QNetworkRequest__UserMax                               QNetworkRequest__Attribute = QNetworkRequest__Attribute(32767)
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
	QNetworkRequest__IfModifiedSinceHeader    QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(9)
	QNetworkRequest__ETagHeader               QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(10)
	QNetworkRequest__IfMatchHeader            QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(11)
	QNetworkRequest__IfNoneMatchHeader        QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(12)
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

//go:generate stringer -type=QNetworkRequest__RedirectPolicy
//QNetworkRequest::RedirectPolicy
type QNetworkRequest__RedirectPolicy int64

const (
	QNetworkRequest__ManualRedirectPolicy       QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(0)
	QNetworkRequest__NoLessSafeRedirectPolicy   QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(1)
	QNetworkRequest__SameOriginRedirectPolicy   QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(2)
	QNetworkRequest__UserVerifiedRedirectPolicy QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(3)
)

func NewQNetworkRequest(url core.QUrl_ITF) *QNetworkRequest {
	tmpValue := NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest(core.PointerFromQUrl(url)))
	qt.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
	return tmpValue
}

func NewQNetworkRequest2(other QNetworkRequest_ITF) *QNetworkRequest {
	tmpValue := NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest2(PointerFromQNetworkRequest(other)))
	qt.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
	return tmpValue
}

func (ptr *QNetworkRequest) Attribute(code QNetworkRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QNetworkRequest_Attribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(defaultValue)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) HasRawHeader(headerName core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkRequest_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QNetworkRequest) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QNetworkRequest_Header(ptr.Pointer(), C.longlong(header)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
		tmpValue := core.NewQObjectFromPointer(C.QNetworkRequest_OriginatingObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) PeerVerifyName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkRequest_PeerVerifyName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkRequest) Priority() QNetworkRequest__Priority {
	if ptr.Pointer() != nil {
		return QNetworkRequest__Priority(C.QNetworkRequest_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkRequest) QNetworkCacheMetaData_RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkRequest_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) RawHeaderList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQNetworkRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__rawHeaderList_atList(i)
			}
			return out
		}(C.QNetworkRequest_RawHeaderList(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
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

func (ptr *QNetworkRequest) SetPeerVerifyName(peerName string) {
	if ptr.Pointer() != nil {
		var peerNameC *C.char
		if peerName != "" {
			peerNameC = C.CString(peerName)
			defer C.free(unsafe.Pointer(peerNameC))
		}
		C.QNetworkRequest_SetPeerVerifyName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: peerNameC, len: C.longlong(len(peerName))})
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
		tmpValue := NewQSslConfigurationFromPointer(C.QNetworkRequest_SslConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
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
		tmpValue := core.NewQUrlFromPointer(C.QNetworkRequest_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) DestroyQNetworkRequest() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkRequest_DestroyQNetworkRequest(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkRequest) __rawHeaderList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkRequest___rawHeaderList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) __rawHeaderList_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest___rawHeaderList_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkRequest) __rawHeaderList_newList() unsafe.Pointer {
	return C.QNetworkRequest___rawHeaderList_newList(ptr.Pointer())
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

func NewQNetworkSessionFromPointer(ptr unsafe.Pointer) (n *QNetworkSession) {
	n = new(QNetworkSession)
	n.SetPointer(ptr)
	return
}

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

//go:generate stringer -type=QNetworkSession__UsagePolicy
//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(1)
)

func NewQNetworkSession(connectionConfig QNetworkConfiguration_ITF, parent core.QObject_ITF) *QNetworkSession {
	tmpValue := NewQNetworkSessionFromPointer(C.QNetworkSession_NewQNetworkSession(PointerFromQNetworkConfiguration(connectionConfig), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNetworkSession_Accept
func callbackQNetworkSession_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QNetworkSession) ConnectAccept(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accept"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "accept", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accept", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectAccept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accept")
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
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QNetworkSession) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
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
	if signal := qt.GetSignal(ptr, "closed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closed") {
			C.QNetworkSession_ConnectClosed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "closed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "closed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "closed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "closed")
	}
}

func (ptr *QNetworkSession) Closed() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Closed(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Configuration() *QNetworkConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkConfigurationFromPointer(C.QNetworkSession_Configuration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_Error2
func callbackQNetworkSession_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QNetworkSession__SessionError))(signal))(QNetworkSession__SessionError(error))
	}

}

func (ptr *QNetworkSession) ConnectError2(f func(error QNetworkSession__SessionError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QNetworkSession_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error QNetworkSession__SessionError) {
				(*(*func(QNetworkSession__SessionError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QNetworkSession) Error2(error QNetworkSession__SessionError) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Error2(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNetworkSession) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNetworkSession_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQNetworkSession_Ignore
func callbackQNetworkSession_Ignore(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ignore"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).IgnoreDefault()
	}
}

func (ptr *QNetworkSession) ConnectIgnore(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignore"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "ignore", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignore", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectIgnore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignore")
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
		tmpValue := NewQNetworkInterfaceFromPointer(C.QNetworkSession_Interface(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) IsOpen() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkSession_IsOpen(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNetworkSession_Migrate
func callbackQNetworkSession_Migrate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "migrate"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).MigrateDefault()
	}
}

func (ptr *QNetworkSession) ConnectMigrate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "migrate"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "migrate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "migrate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectMigrate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "migrate")
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
	if signal := qt.GetSignal(ptr, "newConfigurationActivated"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConfigurationActivated") {
			C.QNetworkSession_ConnectNewConfigurationActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "newConfigurationActivated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConfigurationActivated"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newConfigurationActivated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConfigurationActivated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newConfigurationActivated")
	}
}

func (ptr *QNetworkSession) NewConfigurationActivated() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_NewConfigurationActivated(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Open
func callbackQNetworkSession_Open(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).OpenDefault()
	}
}

func (ptr *QNetworkSession) ConnectOpen(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
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
	if signal := qt.GetSignal(ptr, "opened"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "opened") {
			C.QNetworkSession_ConnectOpened(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "opened")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "opened"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "opened", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "opened", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "opened")
	}
}

func (ptr *QNetworkSession) Opened() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Opened(ptr.Pointer())
	}
}

//export callbackQNetworkSession_PreferredConfigurationChanged
func callbackQNetworkSession_PreferredConfigurationChanged(ptr unsafe.Pointer, config unsafe.Pointer, isSeamless C.char) {
	if signal := qt.GetSignal(ptr, "preferredConfigurationChanged"); signal != nil {
		(*(*func(*QNetworkConfiguration, bool))(signal))(NewQNetworkConfigurationFromPointer(config), int8(isSeamless) != 0)
	}

}

func (ptr *QNetworkSession) ConnectPreferredConfigurationChanged(f func(config *QNetworkConfiguration, isSeamless bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preferredConfigurationChanged") {
			C.QNetworkSession_ConnectPreferredConfigurationChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preferredConfigurationChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preferredConfigurationChanged"); signal != nil {
			f := func(config *QNetworkConfiguration, isSeamless bool) {
				(*(*func(*QNetworkConfiguration, bool))(signal))(config, isSeamless)
				f(config, isSeamless)
			}
			qt.ConnectSignal(ptr.Pointer(), "preferredConfigurationChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preferredConfigurationChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectPreferredConfigurationChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectPreferredConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preferredConfigurationChanged")
	}
}

func (ptr *QNetworkSession) PreferredConfigurationChanged(config QNetworkConfiguration_ITF, isSeamless bool) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_PreferredConfigurationChanged(ptr.Pointer(), PointerFromQNetworkConfiguration(config), C.char(int8(qt.GoBoolToInt(isSeamless))))
	}
}

//export callbackQNetworkSession_Reject
func callbackQNetworkSession_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).RejectDefault()
	}
}

func (ptr *QNetworkSession) ConnectReject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "reject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectReject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reject")
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
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QNetworkSession_SessionProperty(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: keyC, len: C.longlong(len(key))}))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QNetworkSession_SetSessionProperty(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(value))
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
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QNetworkSession__State))(signal))(QNetworkSession__State(state))
	}

}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QNetworkSession_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(state QNetworkSession__State) {
				(*(*func(QNetworkSession__State))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QNetworkSession) StateChanged(state QNetworkSession__State) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQNetworkSession_Stop
func callbackQNetworkSession_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).StopDefault()
	}
}

func (ptr *QNetworkSession) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
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
	if signal := qt.GetSignal(ptr, "usagePoliciesChanged"); signal != nil {
		(*(*func(QNetworkSession__UsagePolicy))(signal))(QNetworkSession__UsagePolicy(usagePolicies))
	}

}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "usagePoliciesChanged") {
			C.QNetworkSession_ConnectUsagePoliciesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "usagePoliciesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "usagePoliciesChanged"); signal != nil {
			f := func(usagePolicies QNetworkSession__UsagePolicy) {
				(*(*func(QNetworkSession__UsagePolicy))(signal))(usagePolicies)
				f(usagePolicies)
			}
			qt.ConnectSignal(ptr.Pointer(), "usagePoliciesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "usagePoliciesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "usagePoliciesChanged")
	}
}

func (ptr *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_UsagePoliciesChanged(ptr.Pointer(), C.longlong(usagePolicies))
	}
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkSession_WaitForOpened(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQNetworkSession_DestroyQNetworkSession
func callbackQNetworkSession_DestroyQNetworkSession(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNetworkSession"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).DestroyQNetworkSessionDefault()
	}
}

func (ptr *QNetworkSession) ConnectDestroyQNetworkSession(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNetworkSession"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkSession", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNetworkSession", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNetworkSession) DisconnectDestroyQNetworkSession() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNetworkSession")
	}
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkSession_DestroyQNetworkSession(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) DestroyQNetworkSessionDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkSession_DestroyQNetworkSessionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkSession___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkSession) __children_newList() unsafe.Pointer {
	return C.QNetworkSession___children_newList(ptr.Pointer())
}

func (ptr *QNetworkSession) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNetworkSession___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNetworkSession) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNetworkSession___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNetworkSession) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkSession___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkSession) __findChildren_newList() unsafe.Pointer {
	return C.QNetworkSession___findChildren_newList(ptr.Pointer())
}

func (ptr *QNetworkSession) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNetworkSession___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNetworkSession) __findChildren_newList3() unsafe.Pointer {
	return C.QNetworkSession___findChildren_newList3(ptr.Pointer())
}

//export callbackQNetworkSession_ChildEvent
func callbackQNetworkSession_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkSession_ConnectNotify
func callbackQNetworkSession_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkSessionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkSession) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkSession_CustomEvent
func callbackQNetworkSession_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkSession_DeleteLater
func callbackQNetworkSession_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNetworkSessionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkSession) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNetworkSession_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Destroyed
func callbackQNetworkSession_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNetworkSession_DisconnectNotify
func callbackQNetworkSession_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkSessionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkSession) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkSession_Event
func callbackQNetworkSession_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkSessionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkSession) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkSession_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNetworkSession_EventFilter
func callbackQNetworkSession_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkSessionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkSession) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNetworkSession_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNetworkSession_MetaObject
func callbackQNetworkSession_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNetworkSessionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkSession) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkSession_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNetworkSession_ObjectNameChanged
func callbackQNetworkSession_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNetworkSession_TimerEvent
func callbackQNetworkSession_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QOcspResponse struct {
	ptr unsafe.Pointer
}

type QOcspResponse_ITF interface {
	QOcspResponse_PTR() *QOcspResponse
}

func (ptr *QOcspResponse) QOcspResponse_PTR() *QOcspResponse {
	return ptr
}

func (ptr *QOcspResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QOcspResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQOcspResponse(ptr QOcspResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOcspResponse_PTR().Pointer()
	}
	return nil
}

func NewQOcspResponseFromPointer(ptr unsafe.Pointer) (n *QOcspResponse) {
	n = new(QOcspResponse)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QOcspResponse__QOcspCertificateStatus
//QOcspResponse::QOcspCertificateStatus
type QOcspResponse__QOcspCertificateStatus int64

const (
	QOcspResponse__Good    QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(0)
	QOcspResponse__Revoked QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(1)
	QOcspResponse__Unknown QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(2)
)

//go:generate stringer -type=QOcspResponse__QOcspRevocationReason
//QOcspResponse::QOcspRevocationReason
type QOcspResponse__QOcspRevocationReason int64

const (
	QOcspResponse__None                 QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(-1)
	QOcspResponse__Unspecified          QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(0)
	QOcspResponse__KeyCompromise        QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(1)
	QOcspResponse__CACompromise         QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(2)
	QOcspResponse__AffiliationChanged   QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(3)
	QOcspResponse__Superseded           QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(4)
	QOcspResponse__CessationOfOperation QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(5)
	QOcspResponse__CertificateHold      QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(6)
	QOcspResponse__RemoveFromCRL        QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(7)
)

func NewQOcspResponse() *QOcspResponse {
	tmpValue := NewQOcspResponseFromPointer(C.QOcspResponse_NewQOcspResponse())
	qt.SetFinalizer(tmpValue, (*QOcspResponse).DestroyQOcspResponse)
	return tmpValue
}

func NewQOcspResponse2(other QOcspResponse_ITF) *QOcspResponse {
	tmpValue := NewQOcspResponseFromPointer(C.QOcspResponse_NewQOcspResponse2(PointerFromQOcspResponse(other)))
	qt.SetFinalizer(tmpValue, (*QOcspResponse).DestroyQOcspResponse)
	return tmpValue
}

func NewQOcspResponse3(other QOcspResponse_ITF) *QOcspResponse {
	tmpValue := NewQOcspResponseFromPointer(C.QOcspResponse_NewQOcspResponse3(PointerFromQOcspResponse(other)))
	qt.SetFinalizer(tmpValue, (*QOcspResponse).DestroyQOcspResponse)
	return tmpValue
}

func (ptr *QOcspResponse) Subject() *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QOcspResponse_Subject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QOcspResponse) Swap(other QOcspResponse_ITF) {
	if ptr.Pointer() != nil {
		C.QOcspResponse_Swap(ptr.Pointer(), PointerFromQOcspResponse(other))
	}
}

func (ptr *QOcspResponse) DestroyQOcspResponse() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QOcspResponse_DestroyQOcspResponse(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPasswordDigestor struct {
	ptr unsafe.Pointer
}

type QPasswordDigestor_ITF interface {
	QPasswordDigestor_PTR() *QPasswordDigestor
}

func (ptr *QPasswordDigestor) QPasswordDigestor_PTR() *QPasswordDigestor {
	return ptr
}

func (ptr *QPasswordDigestor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPasswordDigestor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPasswordDigestor(ptr QPasswordDigestor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPasswordDigestor_PTR().Pointer()
	}
	return nil
}

func NewQPasswordDigestorFromPointer(ptr unsafe.Pointer) (n *QPasswordDigestor) {
	n = new(QPasswordDigestor)
	n.SetPointer(ptr)
	return
}
func (ptr *QPasswordDigestor) DestroyQPasswordDigestor() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSctpServer struct {
	QTcpServer
}

type QSctpServer_ITF interface {
	QTcpServer_ITF
	QSctpServer_PTR() *QSctpServer
}

func (ptr *QSctpServer) QSctpServer_PTR() *QSctpServer {
	return ptr
}

func (ptr *QSctpServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func (ptr *QSctpServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTcpServer_PTR().SetPointer(p)
	}
}

func PointerFromQSctpServer(ptr QSctpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSctpServer_PTR().Pointer()
	}
	return nil
}

func NewQSctpServerFromPointer(ptr unsafe.Pointer) (n *QSctpServer) {
	n = new(QSctpServer)
	n.SetPointer(ptr)
	return
}

type QSctpSocket struct {
	QTcpSocket
}

type QSctpSocket_ITF interface {
	QTcpSocket_ITF
	QSctpSocket_PTR() *QSctpSocket
}

func (ptr *QSctpSocket) QSctpSocket_PTR() *QSctpSocket {
	return ptr
}

func (ptr *QSctpSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QSctpSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTcpSocket_PTR().SetPointer(p)
	}
}

func PointerFromQSctpSocket(ptr QSctpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSctpSocket_PTR().Pointer()
	}
	return nil
}

func NewQSctpSocketFromPointer(ptr unsafe.Pointer) (n *QSctpSocket) {
	n = new(QSctpSocket)
	n.SetPointer(ptr)
	return
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

func NewQSslFromPointer(ptr unsafe.Pointer) (n *QSsl) {
	n = new(QSsl)
	n.SetPointer(ptr)
	return
}
func (ptr *QSsl) DestroyQSsl() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QSsl__KeyType
//QSsl::KeyType
type QSsl__KeyType int64

const (
	QSsl__PrivateKey QSsl__KeyType = QSsl__KeyType(0)
	QSsl__PublicKey  QSsl__KeyType = QSsl__KeyType(1)
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
	QSsl__Dh     QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(4)
)

//go:generate stringer -type=QSsl__AlternativeNameEntryType
//QSsl::AlternativeNameEntryType
type QSsl__AlternativeNameEntryType int64

const (
	QSsl__EmailEntry     QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(0)
	QSsl__DnsEntry       QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(1)
	QSsl__IpAddressEntry QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(2)
)

//go:generate stringer -type=QSsl__SslProtocol
//QSsl::SslProtocol
type QSsl__SslProtocol int64

const (
	QSsl__SslV3           QSsl__SslProtocol = QSsl__SslProtocol(0)
	QSsl__SslV2           QSsl__SslProtocol = QSsl__SslProtocol(1)
	QSsl__TlsV1_0         QSsl__SslProtocol = QSsl__SslProtocol(2)
	QSsl__TlsV1           QSsl__SslProtocol = QSsl__SslProtocol(QSsl__TlsV1_0)
	QSsl__TlsV1_1         QSsl__SslProtocol = QSsl__SslProtocol(3)
	QSsl__TlsV1_2         QSsl__SslProtocol = QSsl__SslProtocol(4)
	QSsl__AnyProtocol     QSsl__SslProtocol = QSsl__SslProtocol(5)
	QSsl__TlsV1SslV3      QSsl__SslProtocol = QSsl__SslProtocol(6)
	QSsl__SecureProtocols QSsl__SslProtocol = QSsl__SslProtocol(7)
	QSsl__TlsV1_0OrLater  QSsl__SslProtocol = QSsl__SslProtocol(8)
	QSsl__TlsV1_1OrLater  QSsl__SslProtocol = QSsl__SslProtocol(9)
	QSsl__TlsV1_2OrLater  QSsl__SslProtocol = QSsl__SslProtocol(10)
	QSsl__DtlsV1_0        QSsl__SslProtocol = QSsl__SslProtocol(11)
	QSsl__DtlsV1_0OrLater QSsl__SslProtocol = QSsl__SslProtocol(12)
	QSsl__DtlsV1_2        QSsl__SslProtocol = QSsl__SslProtocol(13)
	QSsl__DtlsV1_2OrLater QSsl__SslProtocol = QSsl__SslProtocol(14)
	QSsl__TlsV1_3         QSsl__SslProtocol = QSsl__SslProtocol(15)
	QSsl__TlsV1_3OrLater  QSsl__SslProtocol = QSsl__SslProtocol(16)
	QSsl__UnknownProtocol QSsl__SslProtocol = QSsl__SslProtocol(-1)
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

func NewQSslCertificateFromPointer(ptr unsafe.Pointer) (n *QSslCertificate) {
	n = new(QSslCertificate)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate(core.PointerFromQIODevice(device), C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func NewQSslCertificate2(data core.QByteArray_ITF, format QSsl__EncodingFormat) *QSslCertificate {
	tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate2(core.PointerFromQByteArray(data), C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func NewQSslCertificate3(other QSslCertificate_ITF) *QSslCertificate {
	tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate3(PointerFromQSslCertificate(other)))
	qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func (ptr *QSslCertificate) Clear() {
	if ptr.Pointer() != nil {
		C.QSslCertificate_Clear(ptr.Pointer())
	}
}

func (ptr *QSslCertificate) Digest(algorithm core.QCryptographicHash__Algorithm) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate_Digest(ptr.Pointer(), C.longlong(algorithm)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) EffectiveDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QSslCertificate_EffectiveDate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) ExpiryDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QSslCertificate_ExpiryDate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) Extensions() []*QSslCertificateExtension {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificateExtension {
			out := make([]*QSslCertificateExtension, int(l.len))
			tmpList := NewQSslCertificateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__extensions_atList(i)
			}
			return out
		}(C.QSslCertificate_Extensions(ptr.Pointer()))
	}
	return make([]*QSslCertificateExtension, 0)
}

func QSslCertificate_FromData(data core.QByteArray_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__fromData_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromData(core.PointerFromQByteArray(data), C.longlong(format)))
}

func (ptr *QSslCertificate) FromData(data core.QByteArray_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__fromData_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromData(core.PointerFromQByteArray(data), C.longlong(format)))
}

func QSslCertificate_FromDevice(device core.QIODevice_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__fromDevice_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromDevice(core.PointerFromQIODevice(device), C.longlong(format)))
}

func (ptr *QSslCertificate) FromDevice(device core.QIODevice_ITF, format QSsl__EncodingFormat) []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__fromDevice_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromDevice(core.PointerFromQIODevice(device), C.longlong(format)))
}

func QSslCertificate_FromPath(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) []*QSslCertificate {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__fromPath_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromPath(C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(format), C.longlong(syntax)))
}

func (ptr *QSslCertificate) FromPath(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) []*QSslCertificate {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__fromPath_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_FromPath(C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(format), C.longlong(syntax)))
}

func QSslCertificate_ImportPkcs12(device core.QIODevice_ITF, key QSslKey_ITF, certificate QSslCertificate_ITF, caCertificates []*QSslCertificate, passPhrase core.QByteArray_ITF) bool {
	return int8(C.QSslCertificate_QSslCertificate_ImportPkcs12(core.PointerFromQIODevice(device), PointerFromQSslKey(key), PointerFromQSslCertificate(certificate), func() unsafe.Pointer {
		tmpList := NewQSslCertificateFromPointer(NewQSslCertificateFromPointer(nil).__importPkcs12_caCertificates_newList())
		for _, v := range caCertificates {
			tmpList.__importPkcs12_caCertificates_setList(v)
		}
		return tmpList.Pointer()
	}(), core.PointerFromQByteArray(passPhrase))) != 0
}

func (ptr *QSslCertificate) ImportPkcs12(device core.QIODevice_ITF, key QSslKey_ITF, certificate QSslCertificate_ITF, caCertificates []*QSslCertificate, passPhrase core.QByteArray_ITF) bool {
	return int8(C.QSslCertificate_QSslCertificate_ImportPkcs12(core.PointerFromQIODevice(device), PointerFromQSslKey(key), PointerFromQSslCertificate(certificate), func() unsafe.Pointer {
		tmpList := NewQSslCertificateFromPointer(NewQSslCertificateFromPointer(nil).__importPkcs12_caCertificates_newList())
		for _, v := range caCertificates {
			tmpList.__importPkcs12_caCertificates_setList(v)
		}
		return tmpList.Pointer()
	}(), core.PointerFromQByteArray(passPhrase))) != 0
}

func (ptr *QSslCertificate) IsBlacklisted() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslCertificate_IsBlacklisted(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificate) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslCertificate_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificate) IsSelfSigned() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslCertificate_IsSelfSigned(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificate) IssuerDisplayName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCertificate_IssuerDisplayName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificate) IssuerInfo(subject QSslCertificate__SubjectInfo) []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSslCertificate_IssuerInfo(ptr.Pointer(), C.longlong(subject))))
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfo2(attribute core.QByteArray_ITF) []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSslCertificate_IssuerInfo2(ptr.Pointer(), core.PointerFromQByteArray(attribute))))
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfoAttributes() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQSslCertificateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__issuerInfoAttributes_atList(i)
			}
			return out
		}(C.QSslCertificate_IssuerInfoAttributes(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QSslCertificate) PublicKey() *QSslKey {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslKeyFromPointer(C.QSslCertificate_PublicKey(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) SerialNumber() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate_SerialNumber(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) SubjectAlternativeNames() map[QSsl__AlternativeNameEntryType]string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) map[QSsl__AlternativeNameEntryType]string {
			out := make(map[QSsl__AlternativeNameEntryType]string, int(l.len))
			tmpList := NewQSslCertificateFromPointer(l.data)
			for i, v := range tmpList.__subjectAlternativeNames_keyList() {
				out[v] = tmpList.__subjectAlternativeNames_atList(v, i)
			}
			return out
		}(C.QSslCertificate_SubjectAlternativeNames(ptr.Pointer()))
	}
	return make(map[QSsl__AlternativeNameEntryType]string, 0)
}

func (ptr *QSslCertificate) SubjectDisplayName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCertificate_SubjectDisplayName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificate) SubjectInfo(subject QSslCertificate__SubjectInfo) []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSslCertificate_SubjectInfo(ptr.Pointer(), C.longlong(subject))))
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfo2(attribute core.QByteArray_ITF) []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSslCertificate_SubjectInfo2(ptr.Pointer(), core.PointerFromQByteArray(attribute))))
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfoAttributes() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQSslCertificateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__subjectInfoAttributes_atList(i)
			}
			return out
		}(C.QSslCertificate_SubjectInfoAttributes(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QSslCertificate) Swap(other QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate_Swap(ptr.Pointer(), PointerFromQSslCertificate(other))
	}
}

func (ptr *QSslCertificate) ToDer() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate_ToDer(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) ToPem() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate_ToPem(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func QSslCertificate_Verify(certificateChain []*QSslCertificate, hostName string) []*QSslError {
	var hostNameC *C.char
	if hostName != "" {
		hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
	}
	return func(l C.struct_QtNetwork_PackedList) []*QSslError {
		out := make([]*QSslError, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__verify_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_Verify(func() unsafe.Pointer {
		tmpList := NewQSslCertificateFromPointer(NewQSslCertificateFromPointer(nil).__verify_certificateChain_newList())
		for _, v := range certificateChain {
			tmpList.__verify_certificateChain_setList(v)
		}
		return tmpList.Pointer()
	}(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}))
}

func (ptr *QSslCertificate) Verify(certificateChain []*QSslCertificate, hostName string) []*QSslError {
	var hostNameC *C.char
	if hostName != "" {
		hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
	}
	return func(l C.struct_QtNetwork_PackedList) []*QSslError {
		out := make([]*QSslError, int(l.len))
		tmpList := NewQSslCertificateFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__verify_atList(i)
		}
		return out
	}(C.QSslCertificate_QSslCertificate_Verify(func() unsafe.Pointer {
		tmpList := NewQSslCertificateFromPointer(NewQSslCertificateFromPointer(nil).__verify_certificateChain_newList())
		for _, v := range certificateChain {
			tmpList.__verify_certificateChain_setList(v)
		}
		return tmpList.Pointer()
	}(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}))
}

func (ptr *QSslCertificate) Version() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate_Version(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) DestroyQSslCertificate() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSslCertificate_DestroyQSslCertificate(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslCertificate) __extensions_atList(i int) *QSslCertificateExtension {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateExtensionFromPointer(C.QSslCertificate___extensions_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __extensions_setList(i QSslCertificateExtension_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___extensions_setList(ptr.Pointer(), PointerFromQSslCertificateExtension(i))
	}
}

func (ptr *QSslCertificate) __extensions_newList() unsafe.Pointer {
	return C.QSslCertificate___extensions_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __fromData_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate___fromData_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __fromData_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___fromData_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslCertificate) __fromData_newList() unsafe.Pointer {
	return C.QSslCertificate___fromData_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __fromDevice_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate___fromDevice_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __fromDevice_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___fromDevice_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslCertificate) __fromDevice_newList() unsafe.Pointer {
	return C.QSslCertificate___fromDevice_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __fromPath_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate___fromPath_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __fromPath_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___fromPath_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslCertificate) __fromPath_newList() unsafe.Pointer {
	return C.QSslCertificate___fromPath_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __importPkcs12_caCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate___importPkcs12_caCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __importPkcs12_caCertificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___importPkcs12_caCertificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslCertificate) __importPkcs12_caCertificates_newList() unsafe.Pointer {
	return C.QSslCertificate___importPkcs12_caCertificates_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __issuerInfoAttributes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate___issuerInfoAttributes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __issuerInfoAttributes_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___issuerInfoAttributes_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSslCertificate) __issuerInfoAttributes_newList() unsafe.Pointer {
	return C.QSslCertificate___issuerInfoAttributes_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __subjectAlternativeNames_atList(v QSsl__AlternativeNameEntryType, i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslCertificate___subjectAlternativeNames_atList(ptr.Pointer(), C.longlong(v), C.int(int32(i))))
	}
	return ""
}

func (ptr *QSslCertificate) __subjectAlternativeNames_setList(key QSsl__AlternativeNameEntryType, i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QSslCertificate___subjectAlternativeNames_setList(ptr.Pointer(), C.longlong(key), C.struct_QtNetwork_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QSslCertificate) __subjectAlternativeNames_newList() unsafe.Pointer {
	return C.QSslCertificate___subjectAlternativeNames_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __subjectAlternativeNames_keyList() []QSsl__AlternativeNameEntryType {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []QSsl__AlternativeNameEntryType {
			out := make([]QSsl__AlternativeNameEntryType, int(l.len))
			tmpList := NewQSslCertificateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____subjectAlternativeNames_keyList_atList(i)
			}
			return out
		}(C.QSslCertificate___subjectAlternativeNames_keyList(ptr.Pointer()))
	}
	return make([]QSsl__AlternativeNameEntryType, 0)
}

func (ptr *QSslCertificate) __subjectInfoAttributes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslCertificate___subjectInfoAttributes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __subjectInfoAttributes_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___subjectInfoAttributes_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSslCertificate) __subjectInfoAttributes_newList() unsafe.Pointer {
	return C.QSslCertificate___subjectInfoAttributes_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __verify_atList(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QSslCertificate___verify_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __verify_setList(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___verify_setList(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QSslCertificate) __verify_newList() unsafe.Pointer {
	return C.QSslCertificate___verify_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) __verify_certificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslCertificate___verify_certificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) __verify_certificateChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate___verify_certificateChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslCertificate) __verify_certificateChain_newList() unsafe.Pointer {
	return C.QSslCertificate___verify_certificateChain_newList(ptr.Pointer())
}

func (ptr *QSslCertificate) ____subjectAlternativeNames_keyList_atList(i int) QSsl__AlternativeNameEntryType {
	if ptr.Pointer() != nil {
		return QSsl__AlternativeNameEntryType(C.QSslCertificate_____subjectAlternativeNames_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QSslCertificate) ____subjectAlternativeNames_keyList_setList(i QSsl__AlternativeNameEntryType) {
	if ptr.Pointer() != nil {
		C.QSslCertificate_____subjectAlternativeNames_keyList_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QSslCertificate) ____subjectAlternativeNames_keyList_newList() unsafe.Pointer {
	return C.QSslCertificate_____subjectAlternativeNames_keyList_newList(ptr.Pointer())
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

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) (n *QSslCertificateExtension) {
	n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return
}
func NewQSslCertificateExtension() *QSslCertificateExtension {
	tmpValue := NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension())
	qt.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
	return tmpValue
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	tmpValue := NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(PointerFromQSslCertificateExtension(other)))
	qt.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
	return tmpValue
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslCertificateExtension_IsCritical(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslCertificateExtension_IsSupported(ptr.Pointer())) != 0
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
		tmpValue := core.NewQVariantFromPointer(C.QSslCertificateExtension_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQSslCipherFromPointer(ptr unsafe.Pointer) (n *QSslCipher) {
	n = new(QSslCipher)
	n.SetPointer(ptr)
	return
}
func NewQSslCipher() *QSslCipher {
	tmpValue := NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher())
	qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher2(name string) *QSslCipher {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher2(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher3(name string, protocol QSsl__SslProtocol) *QSslCipher {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher3(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(protocol)))
	qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher4(other QSslCipher_ITF) *QSslCipher {
	tmpValue := NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher4(PointerFromQSslCipher(other)))
	qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
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
		return int8(C.QSslCipher_IsNull(ptr.Pointer())) != 0
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

		qt.SetFinalizer(ptr, nil)
		C.QSslCipher_DestroyQSslCipher(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQSslConfigurationFromPointer(ptr unsafe.Pointer) (n *QSslConfiguration) {
	n = new(QSslConfiguration)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration())
	qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func NewQSslConfiguration2(other QSslConfiguration_ITF) *QSslConfiguration {
	tmpValue := NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration2(PointerFromQSslConfiguration(other)))
	qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) AllowedNextProtocols() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__allowedNextProtocols_atList(i)
			}
			return out
		}(C.QSslConfiguration_AllowedNextProtocols(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QSslConfiguration) BackendConfiguration() map[*core.QByteArray]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) map[*core.QByteArray]*core.QVariant {
			out := make(map[*core.QByteArray]*core.QVariant, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i, v := range tmpList.__backendConfiguration_keyList() {
				out[v] = tmpList.__backendConfiguration_atList(v, i)
			}
			return out
		}(C.QSslConfiguration_BackendConfiguration(ptr.Pointer()))
	}
	return make(map[*core.QByteArray]*core.QVariant, 0)
}

func (ptr *QSslConfiguration) CaCertificates() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			out := make([]*QSslCertificate, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__caCertificates_atList(i)
			}
			return out
		}(C.QSslConfiguration_CaCertificates(ptr.Pointer()))
	}
	return make([]*QSslCertificate, 0)
}

func (ptr *QSslConfiguration) Ciphers() []*QSslCipher {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCipher {
			out := make([]*QSslCipher, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__ciphers_atList(i)
			}
			return out
		}(C.QSslConfiguration_Ciphers(ptr.Pointer()))
	}
	return make([]*QSslCipher, 0)
}

func QSslConfiguration_DefaultConfiguration() *QSslConfiguration {
	tmpValue := NewQSslConfigurationFromPointer(C.QSslConfiguration_QSslConfiguration_DefaultConfiguration())
	qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) DefaultConfiguration() *QSslConfiguration {
	tmpValue := NewQSslConfigurationFromPointer(C.QSslConfiguration_QSslConfiguration_DefaultConfiguration())
	qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) DiffieHellmanParameters() *QSslDiffieHellmanParameters {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslConfiguration_DiffieHellmanParameters(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) EphemeralServerKey() *QSslKey {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslKeyFromPointer(C.QSslConfiguration_EphemeralServerKey(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslConfiguration_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslConfiguration) LocalCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration_LocalCertificate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) LocalCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			out := make([]*QSslCertificate, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__localCertificateChain_atList(i)
			}
			return out
		}(C.QSslConfiguration_LocalCertificateChain(ptr.Pointer()))
	}
	return make([]*QSslCertificate, 0)
}

func (ptr *QSslConfiguration) NextNegotiatedProtocol() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration_NextNegotiatedProtocol(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QSslConfiguration) OcspStaplingEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslConfiguration_OcspStaplingEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslConfiguration) PeerCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration_PeerCertificate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) PeerCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			out := make([]*QSslCertificate, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__peerCertificateChain_atList(i)
			}
			return out
		}(C.QSslConfiguration_PeerCertificateChain(ptr.Pointer()))
	}
	return make([]*QSslCertificate, 0)
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

func (ptr *QSslConfiguration) PreSharedKeyIdentityHint() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration_PreSharedKeyIdentityHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) PrivateKey() *QSslKey {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslKeyFromPointer(C.QSslConfiguration_PrivateKey(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
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
		tmpValue := NewQSslCipherFromPointer(C.QSslConfiguration_SessionCipher(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
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
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration_SessionTicket(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QSslConfiguration) SetAllowedNextProtocols(protocols []*core.QByteArray) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetAllowedNextProtocols(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslConfigurationFromPointer(NewQSslConfigurationFromPointer(nil).__setAllowedNextProtocols_protocols_newList())
			for _, v := range protocols {
				tmpList.__setAllowedNextProtocols_protocols_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QSslConfiguration) SetBackendConfiguration(backendConfiguration map[*core.QByteArray]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetBackendConfiguration(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslConfigurationFromPointer(NewQSslConfigurationFromPointer(nil).__setBackendConfiguration_backendConfiguration_newList())
			for k, v := range backendConfiguration {
				tmpList.__setBackendConfiguration_backendConfiguration_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QSslConfiguration) SetBackendConfigurationOption(name core.QByteArray_ITF, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetBackendConfigurationOption(ptr.Pointer(), core.PointerFromQByteArray(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslConfiguration) SetCaCertificates(certificates []*QSslCertificate) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetCaCertificates(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslConfigurationFromPointer(NewQSslConfigurationFromPointer(nil).__setCaCertificates_certificates_newList())
			for _, v := range certificates {
				tmpList.__setCaCertificates_certificates_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QSslConfiguration) SetCiphers(ciphers []*QSslCipher) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetCiphers(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslConfigurationFromPointer(NewQSslConfigurationFromPointer(nil).__setCiphers_ciphers_newList())
			for _, v := range ciphers {
				tmpList.__setCiphers_ciphers_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetDiffieHellmanParameters(dhparams QSslDiffieHellmanParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetDiffieHellmanParameters(ptr.Pointer(), PointerFromQSslDiffieHellmanParameters(dhparams))
	}
}

func (ptr *QSslConfiguration) SetEllipticCurves(curves []*QSslEllipticCurve) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetEllipticCurves(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslConfigurationFromPointer(NewQSslConfigurationFromPointer(nil).__setEllipticCurves_curves_newList())
			for _, v := range curves {
				tmpList.__setEllipticCurves_curves_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslConfiguration) SetLocalCertificateChain(localChain []*QSslCertificate) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificateChain(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslConfigurationFromPointer(NewQSslConfigurationFromPointer(nil).__setLocalCertificateChain_localChain_newList())
			for _, v := range localChain {
				tmpList.__setLocalCertificateChain_localChain_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QSslConfiguration) SetOcspStaplingEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetOcspStaplingEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
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

func (ptr *QSslConfiguration) SetPreSharedKeyIdentityHint(hint core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPreSharedKeyIdentityHint(ptr.Pointer(), core.PointerFromQByteArray(hint))
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
		out := make([]*QSslCipher, int(l.len))
		tmpList := NewQSslConfigurationFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__supportedCiphers_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SupportedCiphers())
}

func (ptr *QSslConfiguration) SupportedCiphers() []*QSslCipher {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCipher {
		out := make([]*QSslCipher, int(l.len))
		tmpList := NewQSslConfigurationFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__supportedCiphers_atList(i)
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
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslConfigurationFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__systemCaCertificates_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SystemCaCertificates())
}

func (ptr *QSslConfiguration) SystemCaCertificates() []*QSslCertificate {
	return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
		out := make([]*QSslCertificate, int(l.len))
		tmpList := NewQSslConfigurationFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__systemCaCertificates_atList(i)
		}
		return out
	}(C.QSslConfiguration_QSslConfiguration_SystemCaCertificates())
}

func (ptr *QSslConfiguration) TestSslOption(option QSsl__SslOption) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslConfiguration_TestSslOption(ptr.Pointer(), C.longlong(option))) != 0
	}
	return false
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSslConfiguration_DestroyQSslConfiguration(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslConfiguration) __allowedNextProtocols_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration___allowedNextProtocols_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __allowedNextProtocols_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___allowedNextProtocols_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSslConfiguration) __allowedNextProtocols_newList() unsafe.Pointer {
	return C.QSslConfiguration___allowedNextProtocols_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __backendConfiguration_atList(v core.QByteArray_ITF, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSslConfiguration___backendConfiguration_atList(ptr.Pointer(), core.PointerFromQByteArray(v), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __backendConfiguration_setList(key core.QByteArray_ITF, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___backendConfiguration_setList(ptr.Pointer(), core.PointerFromQByteArray(key), core.PointerFromQVariant(i))
	}
}

func (ptr *QSslConfiguration) __backendConfiguration_newList() unsafe.Pointer {
	return C.QSslConfiguration___backendConfiguration_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __backendConfiguration_keyList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____backendConfiguration_keyList_atList(i)
			}
			return out
		}(C.QSslConfiguration___backendConfiguration_keyList(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QSslConfiguration) __caCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration___caCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __caCertificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___caCertificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslConfiguration) __caCertificates_newList() unsafe.Pointer {
	return C.QSslConfiguration___caCertificates_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslConfiguration___ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __ciphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___ciphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslConfiguration) __ciphers_newList() unsafe.Pointer {
	return C.QSslConfiguration___ciphers_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __ellipticCurves_atList(i int) *QSslEllipticCurve {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslEllipticCurveFromPointer(C.QSslConfiguration___ellipticCurves_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __ellipticCurves_setList(i QSslEllipticCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___ellipticCurves_setList(ptr.Pointer(), PointerFromQSslEllipticCurve(i))
	}
}

func (ptr *QSslConfiguration) __ellipticCurves_newList() unsafe.Pointer {
	return C.QSslConfiguration___ellipticCurves_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __localCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration___localCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __localCertificateChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___localCertificateChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslConfiguration) __localCertificateChain_newList() unsafe.Pointer {
	return C.QSslConfiguration___localCertificateChain_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __peerCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration___peerCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __peerCertificateChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___peerCertificateChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslConfiguration) __peerCertificateChain_newList() unsafe.Pointer {
	return C.QSslConfiguration___peerCertificateChain_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __setAllowedNextProtocols_protocols_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration___setAllowedNextProtocols_protocols_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __setAllowedNextProtocols_protocols_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___setAllowedNextProtocols_protocols_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSslConfiguration) __setAllowedNextProtocols_protocols_newList() unsafe.Pointer {
	return C.QSslConfiguration___setAllowedNextProtocols_protocols_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_atList(v core.QByteArray_ITF, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSslConfiguration___setBackendConfiguration_backendConfiguration_atList(ptr.Pointer(), core.PointerFromQByteArray(v), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_setList(key core.QByteArray_ITF, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___setBackendConfiguration_backendConfiguration_setList(ptr.Pointer(), core.PointerFromQByteArray(key), core.PointerFromQVariant(i))
	}
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_newList() unsafe.Pointer {
	return C.QSslConfiguration___setBackendConfiguration_backendConfiguration_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_keyList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQSslConfigurationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setBackendConfiguration_backendConfiguration_keyList_atList(i)
			}
			return out
		}(C.QSslConfiguration___setBackendConfiguration_backendConfiguration_keyList(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QSslConfiguration) __setCaCertificates_certificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration___setCaCertificates_certificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __setCaCertificates_certificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___setCaCertificates_certificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslConfiguration) __setCaCertificates_certificates_newList() unsafe.Pointer {
	return C.QSslConfiguration___setCaCertificates_certificates_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __setCiphers_ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslConfiguration___setCiphers_ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __setCiphers_ciphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___setCiphers_ciphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslConfiguration) __setCiphers_ciphers_newList() unsafe.Pointer {
	return C.QSslConfiguration___setCiphers_ciphers_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __setEllipticCurves_curves_atList(i int) *QSslEllipticCurve {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslEllipticCurveFromPointer(C.QSslConfiguration___setEllipticCurves_curves_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __setEllipticCurves_curves_setList(i QSslEllipticCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___setEllipticCurves_curves_setList(ptr.Pointer(), PointerFromQSslEllipticCurve(i))
	}
}

func (ptr *QSslConfiguration) __setEllipticCurves_curves_newList() unsafe.Pointer {
	return C.QSslConfiguration___setEllipticCurves_curves_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __setLocalCertificateChain_localChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration___setLocalCertificateChain_localChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __setLocalCertificateChain_localChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___setLocalCertificateChain_localChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslConfiguration) __setLocalCertificateChain_localChain_newList() unsafe.Pointer {
	return C.QSslConfiguration___setLocalCertificateChain_localChain_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __supportedCiphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslConfiguration___supportedCiphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __supportedCiphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___supportedCiphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslConfiguration) __supportedCiphers_newList() unsafe.Pointer {
	return C.QSslConfiguration___supportedCiphers_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __supportedEllipticCurves_atList(i int) *QSslEllipticCurve {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslEllipticCurveFromPointer(C.QSslConfiguration___supportedEllipticCurves_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __supportedEllipticCurves_setList(i QSslEllipticCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___supportedEllipticCurves_setList(ptr.Pointer(), PointerFromQSslEllipticCurve(i))
	}
}

func (ptr *QSslConfiguration) __supportedEllipticCurves_newList() unsafe.Pointer {
	return C.QSslConfiguration___supportedEllipticCurves_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) __systemCaCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslConfiguration___systemCaCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) __systemCaCertificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration___systemCaCertificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslConfiguration) __systemCaCertificates_newList() unsafe.Pointer {
	return C.QSslConfiguration___systemCaCertificates_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) ____backendConfiguration_keyList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration_____backendConfiguration_keyList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) ____backendConfiguration_keyList_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_____backendConfiguration_keyList_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSslConfiguration) ____backendConfiguration_keyList_newList() unsafe.Pointer {
	return C.QSslConfiguration_____backendConfiguration_keyList_newList(ptr.Pointer())
}

func (ptr *QSslConfiguration) ____setBackendConfiguration_backendConfiguration_keyList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslConfiguration_____setBackendConfiguration_backendConfiguration_keyList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) ____setBackendConfiguration_backendConfiguration_keyList_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_____setBackendConfiguration_backendConfiguration_keyList_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSslConfiguration) ____setBackendConfiguration_backendConfiguration_keyList_newList() unsafe.Pointer {
	return C.QSslConfiguration_____setBackendConfiguration_backendConfiguration_keyList_newList(ptr.Pointer())
}

type QSslDiffieHellmanParameters struct {
	ptr unsafe.Pointer
}

type QSslDiffieHellmanParameters_ITF interface {
	QSslDiffieHellmanParameters_PTR() *QSslDiffieHellmanParameters
}

func (ptr *QSslDiffieHellmanParameters) QSslDiffieHellmanParameters_PTR() *QSslDiffieHellmanParameters {
	return ptr
}

func (ptr *QSslDiffieHellmanParameters) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSslDiffieHellmanParameters) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSslDiffieHellmanParameters(ptr QSslDiffieHellmanParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslDiffieHellmanParameters_PTR().Pointer()
	}
	return nil
}

func NewQSslDiffieHellmanParametersFromPointer(ptr unsafe.Pointer) (n *QSslDiffieHellmanParameters) {
	n = new(QSslDiffieHellmanParameters)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSslDiffieHellmanParameters__Error
//QSslDiffieHellmanParameters::Error
type QSslDiffieHellmanParameters__Error int64

const (
	QSslDiffieHellmanParameters__NoError               QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(0)
	QSslDiffieHellmanParameters__InvalidInputDataError QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(1)
	QSslDiffieHellmanParameters__UnsafeParametersError QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(2)
)

func NewQSslDiffieHellmanParameters() *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_NewQSslDiffieHellmanParameters())
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func NewQSslDiffieHellmanParameters2(other QSslDiffieHellmanParameters_ITF) *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_NewQSslDiffieHellmanParameters2(PointerFromQSslDiffieHellmanParameters(other)))
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func NewQSslDiffieHellmanParameters3(other QSslDiffieHellmanParameters_ITF) *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_NewQSslDiffieHellmanParameters3(PointerFromQSslDiffieHellmanParameters(other)))
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func QSslDiffieHellmanParameters_DefaultParameters() *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_DefaultParameters())
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func (ptr *QSslDiffieHellmanParameters) DefaultParameters() *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_DefaultParameters())
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func (ptr *QSslDiffieHellmanParameters) Error() QSslDiffieHellmanParameters__Error {
	if ptr.Pointer() != nil {
		return QSslDiffieHellmanParameters__Error(C.QSslDiffieHellmanParameters_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslDiffieHellmanParameters) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSslDiffieHellmanParameters_ErrorString(ptr.Pointer()))
	}
	return ""
}

func QSslDiffieHellmanParameters_FromEncoded(encoded core.QByteArray_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_FromEncoded(core.PointerFromQByteArray(encoded), C.longlong(encoding)))
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func (ptr *QSslDiffieHellmanParameters) FromEncoded(encoded core.QByteArray_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_FromEncoded(core.PointerFromQByteArray(encoded), C.longlong(encoding)))
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func QSslDiffieHellmanParameters_FromEncoded2(device core.QIODevice_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_FromEncoded2(core.PointerFromQIODevice(device), C.longlong(encoding)))
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func (ptr *QSslDiffieHellmanParameters) FromEncoded2(device core.QIODevice_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {
	tmpValue := NewQSslDiffieHellmanParametersFromPointer(C.QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_FromEncoded2(core.PointerFromQIODevice(device), C.longlong(encoding)))
	qt.SetFinalizer(tmpValue, (*QSslDiffieHellmanParameters).DestroyQSslDiffieHellmanParameters)
	return tmpValue
}

func (ptr *QSslDiffieHellmanParameters) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslDiffieHellmanParameters_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslDiffieHellmanParameters) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslDiffieHellmanParameters_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslDiffieHellmanParameters) Swap(other QSslDiffieHellmanParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QSslDiffieHellmanParameters_Swap(ptr.Pointer(), PointerFromQSslDiffieHellmanParameters(other))
	}
}

func (ptr *QSslDiffieHellmanParameters) DestroyQSslDiffieHellmanParameters() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSslDiffieHellmanParameters_DestroyQSslDiffieHellmanParameters(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQSslEllipticCurveFromPointer(ptr unsafe.Pointer) (n *QSslEllipticCurve) {
	n = new(QSslEllipticCurve)
	n.SetPointer(ptr)
	return
}
func (ptr *QSslEllipticCurve) DestroyQSslEllipticCurve() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQSslEllipticCurve() *QSslEllipticCurve {
	tmpValue := NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_NewQSslEllipticCurve())
	qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func QSslEllipticCurve_FromLongName(name string) *QSslEllipticCurve {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromLongName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) FromLongName(name string) *QSslEllipticCurve {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromLongName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func QSslEllipticCurve_FromShortName(name string) *QSslEllipticCurve {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromShortName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) FromShortName(name string) *QSslEllipticCurve {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_QSslEllipticCurve_FromShortName(C.struct_QtNetwork_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslEllipticCurve_IsTlsNamedCurve(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslEllipticCurve_IsValid(ptr.Pointer())) != 0
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

func NewQSslErrorFromPointer(ptr unsafe.Pointer) (n *QSslError) {
	n = new(QSslError)
	n.SetPointer(ptr)
	return
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
	QSslError__CertificateStatusUnknown            QSslError__SslError = QSslError__SslError(25)
	QSslError__OcspNoResponseFound                 QSslError__SslError = QSslError__SslError(26)
	QSslError__OcspMalformedRequest                QSslError__SslError = QSslError__SslError(27)
	QSslError__OcspMalformedResponse               QSslError__SslError = QSslError__SslError(28)
	QSslError__OcspInternalError                   QSslError__SslError = QSslError__SslError(29)
	QSslError__OcspTryLater                        QSslError__SslError = QSslError__SslError(30)
	QSslError__OcspSigRequred                      QSslError__SslError = QSslError__SslError(31)
	QSslError__OcspUnauthorized                    QSslError__SslError = QSslError__SslError(32)
	QSslError__OcspResponseCannotBeTrusted         QSslError__SslError = QSslError__SslError(33)
	QSslError__OcspResponseCertIdUnknown           QSslError__SslError = QSslError__SslError(34)
	QSslError__OcspResponseExpired                 QSslError__SslError = QSslError__SslError(35)
	QSslError__OcspStatusUnknown                   QSslError__SslError = QSslError__SslError(36)
	QSslError__UnspecifiedError                    QSslError__SslError = QSslError__SslError(-1)
)

func NewQSslError() *QSslError {
	tmpValue := NewQSslErrorFromPointer(C.QSslError_NewQSslError())
	qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError2(error QSslError__SslError) *QSslError {
	tmpValue := NewQSslErrorFromPointer(C.QSslError_NewQSslError2(C.longlong(error)))
	qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificate_ITF) *QSslError {
	tmpValue := NewQSslErrorFromPointer(C.QSslError_NewQSslError3(C.longlong(error), PointerFromQSslCertificate(certificate)))
	qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError4(other QSslError_ITF) *QSslError {
	tmpValue := NewQSslErrorFromPointer(C.QSslError_NewQSslError4(PointerFromQSslError(other)))
	qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func (ptr *QSslError) Certificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslError_Certificate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
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

		qt.SetFinalizer(ptr, nil)
		C.QSslError_DestroyQSslError(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQSslKeyFromPointer(ptr unsafe.Pointer) (n *QSslKey) {
	n = new(QSslKey)
	n.SetPointer(ptr)
	return
}
func NewQSslKey() *QSslKey {
	tmpValue := NewQSslKeyFromPointer(C.QSslKey_NewQSslKey())
	qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey2(encoded core.QByteArray_ITF, algorithm QSsl__KeyAlgorithm, encoding QSsl__EncodingFormat, ty QSsl__KeyType, passPhrase core.QByteArray_ITF) *QSslKey {
	tmpValue := NewQSslKeyFromPointer(C.QSslKey_NewQSslKey2(core.PointerFromQByteArray(encoded), C.longlong(algorithm), C.longlong(encoding), C.longlong(ty), core.PointerFromQByteArray(passPhrase)))
	qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey3(device core.QIODevice_ITF, algorithm QSsl__KeyAlgorithm, encoding QSsl__EncodingFormat, ty QSsl__KeyType, passPhrase core.QByteArray_ITF) *QSslKey {
	tmpValue := NewQSslKeyFromPointer(C.QSslKey_NewQSslKey3(core.PointerFromQIODevice(device), C.longlong(algorithm), C.longlong(encoding), C.longlong(ty), core.PointerFromQByteArray(passPhrase)))
	qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {
	tmpValue := NewQSslKeyFromPointer(C.QSslKey_NewQSslKey5(PointerFromQSslKey(other)))
	qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
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
		return int8(C.QSslKey_IsNull(ptr.Pointer())) != 0
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
		tmpValue := core.NewQByteArrayFromPointer(C.QSslKey_ToDer(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslKey) ToPem(passPhrase core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslKey_ToPem(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QSslKey_DestroyQSslKey(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) (n *QSslPreSharedKeyAuthenticator) {
	n = new(QSslPreSharedKeyAuthenticator)
	n.SetPointer(ptr)
	return
}
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	tmpValue := NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator())
	qt.SetFinalizer(tmpValue, (*QSslPreSharedKeyAuthenticator).DestroyQSslPreSharedKeyAuthenticator)
	return tmpValue
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {
	tmpValue := NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))
	qt.SetFinalizer(tmpValue, (*QSslPreSharedKeyAuthenticator).DestroyQSslPreSharedKeyAuthenticator)
	return tmpValue
}

func (ptr *QSslPreSharedKeyAuthenticator) Identity() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_Identity(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) IdentityHint() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_IdentityHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
		tmpValue := core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_PreSharedKey(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(ptr.Pointer())
		C.free(ptr.Pointer())
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

func NewQSslSocketFromPointer(ptr unsafe.Pointer) (n *QSslSocket) {
	n = new(QSslSocket)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSslSocket__SslMode
//QSslSocket::SslMode
type QSslSocket__SslMode int64

const (
	QSslSocket__UnencryptedMode QSslSocket__SslMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   QSslSocket__SslMode = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   QSslSocket__SslMode = QSslSocket__SslMode(2)
)

//go:generate stringer -type=QSslSocket__PeerVerifyMode
//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int64

const (
	QSslSocket__VerifyNone     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(3)
)

func NewQSslSocket(parent core.QObject_ITF) *QSslSocket {
	tmpValue := NewQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) AddCaCertificates(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		return int8(C.QSslSocket_AddCaCertificates(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(format), C.longlong(syntax))) != 0
	}
	return false
}

func (ptr *QSslSocket) AddCaCertificates2(certificates []*QSslCertificate) {
	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificates2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslSocketFromPointer(NewQSslSocketFromPointer(nil).__addCaCertificates_certificates_newList2())
			for _, v := range certificates {
				tmpList.__addCaCertificates_certificates_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func QSslSocket_AddDefaultCaCertificates(path string, encoding QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	return int8(C.QSslSocket_QSslSocket_AddDefaultCaCertificates(C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(encoding), C.longlong(syntax))) != 0
}

func (ptr *QSslSocket) AddDefaultCaCertificates(path string, encoding QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	return int8(C.QSslSocket_QSslSocket_AddDefaultCaCertificates(C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(encoding), C.longlong(syntax))) != 0
}

func QSslSocket_AddDefaultCaCertificates2(certificates []*QSslCertificate) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificates2(func() unsafe.Pointer {
		tmpList := NewQSslSocketFromPointer(NewQSslSocketFromPointer(nil).__addDefaultCaCertificates_certificates_newList2())
		for _, v := range certificates {
			tmpList.__addDefaultCaCertificates_certificates_setList2(v)
		}
		return tmpList.Pointer()
	}())
}

func (ptr *QSslSocket) AddDefaultCaCertificates2(certificates []*QSslCertificate) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificates2(func() unsafe.Pointer {
		tmpList := NewQSslSocketFromPointer(NewQSslSocketFromPointer(nil).__addDefaultCaCertificates_certificates_newList2())
		for _, v := range certificates {
			tmpList.__addDefaultCaCertificates_certificates_setList2(v)
		}
		return tmpList.Pointer()
	}())
}

func (ptr *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16, mode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QSslSocket_ConnectToHostEncrypted(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), C.longlong(mode), C.longlong(protocol))
	}
}

func (ptr *QSslSocket) ConnectToHostEncrypted2(hostName string, port uint16, sslPeerName string, mode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	if ptr.Pointer() != nil {
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		var sslPeerNameC *C.char
		if sslPeerName != "" {
			sslPeerNameC = C.CString(sslPeerName)
			defer C.free(unsafe.Pointer(sslPeerNameC))
		}
		C.QSslSocket_ConnectToHostEncrypted2(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))}, C.ushort(port), C.struct_QtNetwork_PackedString{data: sslPeerNameC, len: C.longlong(len(sslPeerName))}, C.longlong(mode), C.longlong(protocol))
	}
}

//export callbackQSslSocket_Encrypted
func callbackQSslSocket_Encrypted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "encrypted"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "encrypted") {
			C.QSslSocket_ConnectEncrypted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "encrypted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "encrypted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "encrypted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encrypted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "encrypted")
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
	if signal := qt.GetSignal(ptr, "encryptedBytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(written))
	}

}

func (ptr *QSslSocket) ConnectEncryptedBytesWritten(f func(written int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "encryptedBytesWritten") {
			C.QSslSocket_ConnectEncryptedBytesWritten(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "encryptedBytesWritten")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "encryptedBytesWritten"); signal != nil {
			f := func(written int64) {
				(*(*func(int64))(signal))(written)
				f(written)
			}
			qt.ConnectSignal(ptr.Pointer(), "encryptedBytesWritten", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encryptedBytesWritten", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectEncryptedBytesWritten() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncryptedBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "encryptedBytesWritten")
	}
}

func (ptr *QSslSocket) EncryptedBytesWritten(written int64) {
	if ptr.Pointer() != nil {
		C.QSslSocket_EncryptedBytesWritten(ptr.Pointer(), C.longlong(written))
	}
}

//export callbackQSslSocket_IgnoreSslErrors
func callbackQSslSocket_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ignoreSslErrors"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSslSocketFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QSslSocket) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignoreSslErrors"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectIgnoreSslErrors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignoreSslErrors")
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

func (ptr *QSslSocket) IgnoreSslErrors2(errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslSocketFromPointer(NewQSslSocketFromPointer(nil).__ignoreSslErrors_errors_newList2())
			for _, v := range errors {
				tmpList.__ignoreSslErrors_errors_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslSocket_IsEncrypted(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslSocket) LocalCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket_LocalCertificate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) LocalCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			out := make([]*QSslCertificate, int(l.len))
			tmpList := NewQSslSocketFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__localCertificateChain_atList(i)
			}
			return out
		}(C.QSslSocket_LocalCertificateChain(ptr.Pointer()))
	}
	return make([]*QSslCertificate, 0)
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_ModeChanged
func callbackQSslSocket_ModeChanged(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "modeChanged"); signal != nil {
		(*(*func(QSslSocket__SslMode))(signal))(QSslSocket__SslMode(mode))
	}

}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modeChanged") {
			C.QSslSocket_ConnectModeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modeChanged"); signal != nil {
			f := func(mode QSslSocket__SslMode) {
				(*(*func(QSslSocket__SslMode))(signal))(mode)
				f(mode)
			}
			qt.ConnectSignal(ptr.Pointer(), "modeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modeChanged")
	}
}

func (ptr *QSslSocket) ModeChanged(mode QSslSocket__SslMode) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ModeChanged(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslSocket) OcspResponses() []*QOcspResponse {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QOcspResponse {
			out := make([]*QOcspResponse, int(l.len))
			tmpList := NewQSslSocketFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__ocspResponses_atList(i)
			}
			return out
		}(C.QSslSocket_OcspResponses(ptr.Pointer()))
	}
	return make([]*QOcspResponse, 0)
}

func (ptr *QSslSocket) PeerCertificate() *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket_PeerCertificate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) PeerCertificateChain() []*QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslCertificate {
			out := make([]*QSslCertificate, int(l.len))
			tmpList := NewQSslSocketFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__peerCertificateChain_atList(i)
			}
			return out
		}(C.QSslSocket_PeerCertificateChain(ptr.Pointer()))
	}
	return make([]*QSslCertificate, 0)
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSslSocket_PeerVerifyDepth(ptr.Pointer())))
	}
	return 0
}

//export callbackQSslSocket_PeerVerifyError
func callbackQSslSocket_PeerVerifyError(ptr unsafe.Pointer, error unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "peerVerifyError"); signal != nil {
		(*(*func(*QSslError))(signal))(NewQSslErrorFromPointer(error))
	}

}

func (ptr *QSslSocket) ConnectPeerVerifyError(f func(error *QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "peerVerifyError") {
			C.QSslSocket_ConnectPeerVerifyError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "peerVerifyError")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "peerVerifyError"); signal != nil {
			f := func(error *QSslError) {
				(*(*func(*QSslError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "peerVerifyError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peerVerifyError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectPeerVerifyError() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPeerVerifyError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "peerVerifyError")
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
	if signal := qt.GetSignal(ptr, "preSharedKeyAuthenticationRequired"); signal != nil {
		(*(*func(*QSslPreSharedKeyAuthenticator))(signal))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preSharedKeyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			f := func(authenticator *QSslPreSharedKeyAuthenticator) {
				(*(*func(*QSslPreSharedKeyAuthenticator))(signal))(authenticator)
				f(authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslSocket) PrivateKey() *QSslKey {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslKeyFromPointer(C.QSslSocket_PrivateKey(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
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

func (ptr *QSslSocket) SessionCipher() *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslSocket_SessionCipher(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
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
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QSslSocket_SetLocalCertificate2(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(format))
	}
}

func (ptr *QSslSocket) SetLocalCertificateChain(localChain []*QSslCertificate) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificateChain(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslSocketFromPointer(NewQSslSocketFromPointer(nil).__setLocalCertificateChain_localChain_newList())
			for _, v := range localChain {
				tmpList.__setLocalCertificateChain_localChain_setList(v)
			}
			return tmpList.Pointer()
		}())
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
		var hostNameC *C.char
		if hostName != "" {
			hostNameC = C.CString(hostName)
			defer C.free(unsafe.Pointer(hostNameC))
		}
		C.QSslSocket_SetPeerVerifyName(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: hostNameC, len: C.longlong(len(hostName))})
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslSocket) SetPrivateKey2(fileName string, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, passPhrase core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QSslSocket_SetPrivateKey2(ptr.Pointer(), C.struct_QtNetwork_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.longlong(algorithm), C.longlong(format), core.PointerFromQByteArray(passPhrase))
	}
}

func (ptr *QSslSocket) SetProtocol(protocol QSsl__SslProtocol) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetProtocol(ptr.Pointer(), C.longlong(protocol))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QSslSocket) SslConfiguration() *QSslConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslConfigurationFromPointer(C.QSslSocket_SslConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SslErrors() []*QSslError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNetwork_PackedList) []*QSslError {
			out := make([]*QSslError, int(l.len))
			tmpList := NewQSslSocketFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sslErrors_atList(i)
			}
			return out
		}(C.QSslSocket_SslErrors(ptr.Pointer()))
	}
	return make([]*QSslError, 0)
}

//export callbackQSslSocket_SslErrors2
func callbackQSslSocket_SslErrors2(ptr unsafe.Pointer, errors C.struct_QtNetwork_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors2"); signal != nil {
		(*(*func([]*QSslError))(signal))(func(l C.struct_QtNetwork_PackedList) []*QSslError {
			out := make([]*QSslError, int(l.len))
			tmpList := NewQSslSocketFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sslErrors_errors_atList2(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QSslSocket) ConnectSslErrors2(f func(errors []*QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors2") {
			C.QSslSocket_ConnectSslErrors2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sslErrors")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors2"); signal != nil {
			f := func(errors []*QSslError) {
				(*(*func([]*QSslError))(signal))(errors)
				f(errors)
			}
			qt.ConnectSignal(ptr.Pointer(), "sslErrors2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectSslErrors2() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectSslErrors2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sslErrors2")
	}
}

func (ptr *QSslSocket) SslErrors2(errors []*QSslError) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SslErrors2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSslSocketFromPointer(NewQSslSocketFromPointer(nil).__sslErrors_errors_newList2())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
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
	if signal := qt.GetSignal(ptr, "startClientEncryption"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSslSocketFromPointer(ptr).StartClientEncryptionDefault()
	}
}

func (ptr *QSslSocket) ConnectStartClientEncryption(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startClientEncryption"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startClientEncryption", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startClientEncryption", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectStartClientEncryption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startClientEncryption")
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
	if signal := qt.GetSignal(ptr, "startServerEncryption"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSslSocketFromPointer(ptr).StartServerEncryptionDefault()
	}
}

func (ptr *QSslSocket) ConnectStartServerEncryption(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startServerEncryption"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startServerEncryption", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startServerEncryption", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectStartServerEncryption() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startServerEncryption")
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
	return int8(C.QSslSocket_QSslSocket_SupportsSsl()) != 0
}

func (ptr *QSslSocket) SupportsSsl() bool {
	return int8(C.QSslSocket_QSslSocket_SupportsSsl()) != 0
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSslSocket_WaitForEncrypted(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQSslSocket_DestroyQSslSocket
func callbackQSslSocket_DestroyQSslSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSslSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSslSocketFromPointer(ptr).DestroyQSslSocketDefault()
	}
}

func (ptr *QSslSocket) ConnectDestroyQSslSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSslSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSslSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSslSocket", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSslSocket) DisconnectDestroyQSslSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSslSocket")
	}
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSslSocket_DestroyQSslSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslSocket) DestroyQSslSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSslSocket_DestroyQSslSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslSocket) __addCaCertificates_certificates_atList2(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___addCaCertificates_certificates_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __addCaCertificates_certificates_setList2(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___addCaCertificates_certificates_setList2(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __addCaCertificates_certificates_newList2() unsafe.Pointer {
	return C.QSslSocket___addCaCertificates_certificates_newList2(ptr.Pointer())
}

func (ptr *QSslSocket) __addDefaultCaCertificates_certificates_atList2(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___addDefaultCaCertificates_certificates_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __addDefaultCaCertificates_certificates_setList2(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___addDefaultCaCertificates_certificates_setList2(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __addDefaultCaCertificates_certificates_newList2() unsafe.Pointer {
	return C.QSslSocket___addDefaultCaCertificates_certificates_newList2(ptr.Pointer())
}

func (ptr *QSslSocket) __caCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___caCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __caCertificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___caCertificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __caCertificates_newList() unsafe.Pointer {
	return C.QSslSocket___caCertificates_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslSocket___ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __ciphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___ciphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslSocket) __ciphers_newList() unsafe.Pointer {
	return C.QSslSocket___ciphers_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __defaultCaCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___defaultCaCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __defaultCaCertificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___defaultCaCertificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __defaultCaCertificates_newList() unsafe.Pointer {
	return C.QSslSocket___defaultCaCertificates_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __defaultCiphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslSocket___defaultCiphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __defaultCiphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___defaultCiphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslSocket) __defaultCiphers_newList() unsafe.Pointer {
	return C.QSslSocket___defaultCiphers_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __ignoreSslErrors_errors_atList2(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QSslSocket___ignoreSslErrors_errors_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __ignoreSslErrors_errors_setList2(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___ignoreSslErrors_errors_setList2(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QSslSocket) __ignoreSslErrors_errors_newList2() unsafe.Pointer {
	return C.QSslSocket___ignoreSslErrors_errors_newList2(ptr.Pointer())
}

func (ptr *QSslSocket) __localCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___localCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __localCertificateChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___localCertificateChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __localCertificateChain_newList() unsafe.Pointer {
	return C.QSslSocket___localCertificateChain_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __ocspResponses_atList(i int) *QOcspResponse {
	if ptr.Pointer() != nil {
		tmpValue := NewQOcspResponseFromPointer(C.QSslSocket___ocspResponses_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QOcspResponse).DestroyQOcspResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __ocspResponses_setList(i QOcspResponse_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___ocspResponses_setList(ptr.Pointer(), PointerFromQOcspResponse(i))
	}
}

func (ptr *QSslSocket) __ocspResponses_newList() unsafe.Pointer {
	return C.QSslSocket___ocspResponses_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __peerCertificateChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___peerCertificateChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __peerCertificateChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___peerCertificateChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __peerCertificateChain_newList() unsafe.Pointer {
	return C.QSslSocket___peerCertificateChain_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __setCaCertificates_certificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___setCaCertificates_certificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __setCaCertificates_certificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___setCaCertificates_certificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __setCaCertificates_certificates_newList() unsafe.Pointer {
	return C.QSslSocket___setCaCertificates_certificates_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __setCiphers_ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslSocket___setCiphers_ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __setCiphers_ciphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___setCiphers_ciphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslSocket) __setCiphers_ciphers_newList() unsafe.Pointer {
	return C.QSslSocket___setCiphers_ciphers_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __setDefaultCaCertificates_certificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___setDefaultCaCertificates_certificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __setDefaultCaCertificates_certificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___setDefaultCaCertificates_certificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __setDefaultCaCertificates_certificates_newList() unsafe.Pointer {
	return C.QSslSocket___setDefaultCaCertificates_certificates_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __setDefaultCiphers_ciphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslSocket___setDefaultCiphers_ciphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __setDefaultCiphers_ciphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___setDefaultCiphers_ciphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslSocket) __setDefaultCiphers_ciphers_newList() unsafe.Pointer {
	return C.QSslSocket___setDefaultCiphers_ciphers_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __setLocalCertificateChain_localChain_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___setLocalCertificateChain_localChain_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __setLocalCertificateChain_localChain_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___setLocalCertificateChain_localChain_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __setLocalCertificateChain_localChain_newList() unsafe.Pointer {
	return C.QSslSocket___setLocalCertificateChain_localChain_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __sslErrors_atList(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QSslSocket___sslErrors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __sslErrors_setList(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___sslErrors_setList(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QSslSocket) __sslErrors_newList() unsafe.Pointer {
	return C.QSslSocket___sslErrors_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __sslErrors_errors_atList2(i int) *QSslError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslErrorFromPointer(C.QSslSocket___sslErrors_errors_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __sslErrors_errors_setList2(i QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___sslErrors_errors_setList2(ptr.Pointer(), PointerFromQSslError(i))
	}
}

func (ptr *QSslSocket) __sslErrors_errors_newList2() unsafe.Pointer {
	return C.QSslSocket___sslErrors_errors_newList2(ptr.Pointer())
}

func (ptr *QSslSocket) __supportedCiphers_atList(i int) *QSslCipher {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCipherFromPointer(C.QSslSocket___supportedCiphers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __supportedCiphers_setList(i QSslCipher_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___supportedCiphers_setList(ptr.Pointer(), PointerFromQSslCipher(i))
	}
}

func (ptr *QSslSocket) __supportedCiphers_newList() unsafe.Pointer {
	return C.QSslSocket___supportedCiphers_newList(ptr.Pointer())
}

func (ptr *QSslSocket) __systemCaCertificates_atList(i int) *QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := NewQSslCertificateFromPointer(C.QSslSocket___systemCaCertificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) __systemCaCertificates_setList(i QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket___systemCaCertificates_setList(ptr.Pointer(), PointerFromQSslCertificate(i))
	}
}

func (ptr *QSslSocket) __systemCaCertificates_newList() unsafe.Pointer {
	return C.QSslSocket___systemCaCertificates_newList(ptr.Pointer())
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

func NewQTcpServerFromPointer(ptr unsafe.Pointer) (n *QTcpServer) {
	n = new(QTcpServer)
	n.SetPointer(ptr)
	return
}
func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {
	tmpValue := NewQTcpServerFromPointer(C.QTcpServer_NewQTcpServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTcpServer_AcceptError
func callbackQTcpServer_AcceptError(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(ptr, "acceptError"); signal != nil {
		(*(*func(QAbstractSocket__SocketError))(signal))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "acceptError") {
			C.QTcpServer_ConnectAcceptError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "acceptError")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "acceptError"); signal != nil {
			f := func(socketError QAbstractSocket__SocketError) {
				(*(*func(QAbstractSocket__SocketError))(signal))(socketError)
				f(socketError)
			}
			qt.ConnectSignal(ptr.Pointer(), "acceptError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "acceptError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "acceptError")
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
	if signal := qt.GetSignal(ptr, "hasPendingConnections"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).HasPendingConnectionsDefault())))
}

func (ptr *QTcpServer) ConnectHasPendingConnections(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasPendingConnections"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hasPendingConnections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasPendingConnections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTcpServer) DisconnectHasPendingConnections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasPendingConnections")
	}
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTcpServer_HasPendingConnections(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTcpServer) HasPendingConnectionsDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTcpServer_HasPendingConnectionsDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTcpServer_IsListening(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTcpServer) Listen(address QHostAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTcpServer_Listen(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port))) != 0
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
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QTcpServer_ConnectNewConnection(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "newConnection")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newConnection")
	}
}

func (ptr *QTcpServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QTcpServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQTcpServer_NextPendingConnection
func callbackQTcpServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextPendingConnection"); signal != nil {
		return PointerFromQTcpSocket((*(*func() *QTcpSocket)(signal))())
	}

	return PointerFromQTcpSocket(NewQTcpServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QTcpServer) ConnectNextPendingConnection(f func() *QTcpSocket) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextPendingConnection"); signal != nil {
			f := func() *QTcpSocket {
				(*(*func() *QTcpSocket)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTcpServer) DisconnectNextPendingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nextPendingConnection")
	}
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) NextPendingConnectionDefault() *QTcpSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
		tmpValue := NewQNetworkProxyFromPointer(C.QTcpServer_Proxy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
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
		tmpValue := NewQHostAddressFromPointer(C.QTcpServer_ServerAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
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

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut *bool) bool {
	if ptr.Pointer() != nil {
		var timedOutC C.char
		if timedOut != nil {
			timedOutC = C.char(int8(qt.GoBoolToInt(*timedOut)))
			defer func() { *timedOut = int8(timedOutC) != 0 }()
		}
		return int8(C.QTcpServer_WaitForNewConnection(ptr.Pointer(), C.int(int32(msec)), &timedOutC)) != 0
	}
	return false
}

//export callbackQTcpServer_DestroyQTcpServer
func callbackQTcpServer_DestroyQTcpServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTcpServer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTcpServerFromPointer(ptr).DestroyQTcpServerDefault()
	}
}

func (ptr *QTcpServer) ConnectDestroyQTcpServer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTcpServer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTcpServer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTcpServer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTcpServer) DisconnectDestroyQTcpServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTcpServer")
	}
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTcpServer_DestroyQTcpServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) DestroyQTcpServerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTcpServer_DestroyQTcpServerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTcpServer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTcpServer) __children_newList() unsafe.Pointer {
	return C.QTcpServer___children_newList(ptr.Pointer())
}

func (ptr *QTcpServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTcpServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QTcpServer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QTcpServer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QTcpServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTcpServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTcpServer) __findChildren_newList() unsafe.Pointer {
	return C.QTcpServer___findChildren_newList(ptr.Pointer())
}

func (ptr *QTcpServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTcpServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTcpServer) __findChildren_newList3() unsafe.Pointer {
	return C.QTcpServer___findChildren_newList3(ptr.Pointer())
}

//export callbackQTcpServer_ChildEvent
func callbackQTcpServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTcpServer_ConnectNotify
func callbackQTcpServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpServer_CustomEvent
func callbackQTcpServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTcpServer_DeleteLater
func callbackQTcpServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTcpServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTcpServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTcpServer_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQTcpServer_Destroyed
func callbackQTcpServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQTcpServer_DisconnectNotify
func callbackQTcpServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpServer_Event
func callbackQTcpServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTcpServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTcpServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQTcpServer_EventFilter
func callbackQTcpServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTcpServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTcpServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQTcpServer_MetaObject
func callbackQTcpServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQTcpServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTcpServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQTcpServer_ObjectNameChanged
func callbackQTcpServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNetwork_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQTcpServer_TimerEvent
func callbackQTcpServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) (n *QTcpSocket) {
	n = new(QTcpSocket)
	n.SetPointer(ptr)
	return
}
func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {
	tmpValue := NewQTcpSocketFromPointer(C.QTcpSocket_NewQTcpSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTcpSocket_DestroyQTcpSocket
func callbackQTcpSocket_DestroyQTcpSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTcpSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTcpSocketFromPointer(ptr).DestroyQTcpSocketDefault()
	}
}

func (ptr *QTcpSocket) ConnectDestroyQTcpSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTcpSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTcpSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTcpSocket", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTcpSocket) DisconnectDestroyQTcpSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTcpSocket")
	}
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTcpSocket_DestroyQTcpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) DestroyQTcpSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTcpSocket_DestroyQTcpSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) (n *QUdpSocket) {
	n = new(QUdpSocket)
	n.SetPointer(ptr)
	return
}
func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {
	tmpValue := NewQUdpSocketFromPointer(C.QUdpSocket_NewQUdpSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	if ptr.Pointer() != nil {
		return int8(C.QUdpSocket_HasPendingDatagrams(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QUdpSocket_JoinMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress))) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QUdpSocket_JoinMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface))) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QUdpSocket_LeaveMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress))) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QUdpSocket_LeaveMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface))) != 0
	}
	return false
}

func (ptr *QUdpSocket) MulticastInterface() *QNetworkInterface {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkInterfaceFromPointer(C.QUdpSocket_MulticastInterface(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
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

func (ptr *QUdpSocket) ReadDatagram(data []byte, maxSize int64, address QHostAddress_ITF, port uint16) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QUdpSocket_ReadDatagram(ptr.Pointer(), dataC, C.longlong(maxSize), PointerFromQHostAddress(address), C.ushort(port)))
	}
	return 0
}

func (ptr *QUdpSocket) ReceiveDatagram(maxSize int64) *QNetworkDatagram {
	if ptr.Pointer() != nil {
		tmpValue := NewQNetworkDatagramFromPointer(C.QUdpSocket_ReceiveDatagram(ptr.Pointer(), C.longlong(maxSize)))
		qt.SetFinalizer(tmpValue, (*QNetworkDatagram).DestroyQNetworkDatagram)
		return tmpValue
	}
	return nil
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(ptr.Pointer(), PointerFromQNetworkInterface(iface))
	}
}

func (ptr *QUdpSocket) WriteDatagram(data []byte, size int64, address QHostAddress_ITF, port uint16) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QUdpSocket_WriteDatagram(ptr.Pointer(), dataC, C.longlong(size), PointerFromQHostAddress(address), C.ushort(port)))
	}
	return 0
}

func (ptr *QUdpSocket) WriteDatagram2(datagram QNetworkDatagram_ITF) int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_WriteDatagram2(ptr.Pointer(), PointerFromQNetworkDatagram(datagram)))
	}
	return 0
}

func (ptr *QUdpSocket) WriteDatagram3(datagram core.QByteArray_ITF, host QHostAddress_ITF, port uint16) int64 {
	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_WriteDatagram3(ptr.Pointer(), core.PointerFromQByteArray(datagram), PointerFromQHostAddress(host), C.ushort(port)))
	}
	return 0
}

//export callbackQUdpSocket_DestroyQUdpSocket
func callbackQUdpSocket_DestroyQUdpSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QUdpSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQUdpSocketFromPointer(ptr).DestroyQUdpSocketDefault()
	}
}

func (ptr *QUdpSocket) ConnectDestroyQUdpSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QUdpSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QUdpSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QUdpSocket", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QUdpSocket) DisconnectDestroyQUdpSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QUdpSocket")
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QUdpSocket_DestroyQUdpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QUdpSocket_DestroyQUdpSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkCacheMetaData_RawHeader struct {
	ptr unsafe.Pointer
}

type QNetworkCacheMetaData_RawHeader_ITF interface {
	QNetworkCacheMetaData_RawHeader_PTR() *QNetworkCacheMetaData_RawHeader
}

func (ptr *QNetworkCacheMetaData_RawHeader) QNetworkCacheMetaData_RawHeader_PTR() *QNetworkCacheMetaData_RawHeader {
	return ptr
}

func (ptr *QNetworkCacheMetaData_RawHeader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNetworkCacheMetaData_RawHeader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNetworkCacheMetaData_RawHeader(ptr QNetworkCacheMetaData_RawHeader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_RawHeader_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCacheMetaData_RawHeaderFromPointer(ptr unsafe.Pointer) (n *QNetworkCacheMetaData_RawHeader) {
	n = new(QNetworkCacheMetaData_RawHeader)
	n.SetPointer(ptr)
	return
}
func (ptr *QNetworkCacheMetaData_RawHeader) DestroyQNetworkCacheMetaData_RawHeader() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQNetworkCacheMetaData_RawHeader() *QNetworkCacheMetaData_RawHeader {
	tmpValue := NewQNetworkCacheMetaData_RawHeaderFromPointer(C.RawHeader_NewRawHeader())
	qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData_RawHeader).DestroyQNetworkCacheMetaData_RawHeader)
	return tmpValue
}

func NewQNetworkCacheMetaData_RawHeader2(first core.QByteArray_ITF, second core.QByteArray_ITF) *QNetworkCacheMetaData_RawHeader {
	tmpValue := NewQNetworkCacheMetaData_RawHeaderFromPointer(C.RawHeader_NewRawHeader2(core.PointerFromQByteArray(first), core.PointerFromQByteArray(second)))
	qt.SetFinalizer(tmpValue, (*QNetworkCacheMetaData_RawHeader).DestroyQNetworkCacheMetaData_RawHeader)
	return tmpValue
}

func (ptr *QNetworkCacheMetaData_RawHeader) First() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.RawHeader_First(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData_RawHeader) SetFirst(vqb core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.RawHeader_SetFirst(ptr.Pointer(), core.PointerFromQByteArray(vqb))
	}
}

func (ptr *QNetworkCacheMetaData_RawHeader) Second() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.RawHeader_Second(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData_RawHeader) SetSecond(vqb core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.RawHeader_SetSecond(ptr.Pointer(), core.PointerFromQByteArray(vqb))
	}
}

func init() {
	qt.ItfMap["network.QAbstractNetworkCache_ITF"] = QAbstractNetworkCache{}
	qt.FuncMap["network.NewQAbstractNetworkCache"] = NewQAbstractNetworkCache
	qt.ItfMap["network.QAbstractSocket_ITF"] = QAbstractSocket{}
	qt.FuncMap["network.NewQAbstractSocket"] = NewQAbstractSocket
	qt.EnumMap["network.QAbstractSocket__TcpSocket"] = int64(QAbstractSocket__TcpSocket)
	qt.EnumMap["network.QAbstractSocket__UdpSocket"] = int64(QAbstractSocket__UdpSocket)
	qt.EnumMap["network.QAbstractSocket__SctpSocket"] = int64(QAbstractSocket__SctpSocket)
	qt.EnumMap["network.QAbstractSocket__UnknownSocketType"] = int64(QAbstractSocket__UnknownSocketType)
	qt.EnumMap["network.QAbstractSocket__IPv4Protocol"] = int64(QAbstractSocket__IPv4Protocol)
	qt.EnumMap["network.QAbstractSocket__IPv6Protocol"] = int64(QAbstractSocket__IPv6Protocol)
	qt.EnumMap["network.QAbstractSocket__AnyIPProtocol"] = int64(QAbstractSocket__AnyIPProtocol)
	qt.EnumMap["network.QAbstractSocket__UnknownNetworkLayerProtocol"] = int64(QAbstractSocket__UnknownNetworkLayerProtocol)
	qt.EnumMap["network.QAbstractSocket__ConnectionRefusedError"] = int64(QAbstractSocket__ConnectionRefusedError)
	qt.EnumMap["network.QAbstractSocket__RemoteHostClosedError"] = int64(QAbstractSocket__RemoteHostClosedError)
	qt.EnumMap["network.QAbstractSocket__HostNotFoundError"] = int64(QAbstractSocket__HostNotFoundError)
	qt.EnumMap["network.QAbstractSocket__SocketAccessError"] = int64(QAbstractSocket__SocketAccessError)
	qt.EnumMap["network.QAbstractSocket__SocketResourceError"] = int64(QAbstractSocket__SocketResourceError)
	qt.EnumMap["network.QAbstractSocket__SocketTimeoutError"] = int64(QAbstractSocket__SocketTimeoutError)
	qt.EnumMap["network.QAbstractSocket__DatagramTooLargeError"] = int64(QAbstractSocket__DatagramTooLargeError)
	qt.EnumMap["network.QAbstractSocket__NetworkError"] = int64(QAbstractSocket__NetworkError)
	qt.EnumMap["network.QAbstractSocket__AddressInUseError"] = int64(QAbstractSocket__AddressInUseError)
	qt.EnumMap["network.QAbstractSocket__SocketAddressNotAvailableError"] = int64(QAbstractSocket__SocketAddressNotAvailableError)
	qt.EnumMap["network.QAbstractSocket__UnsupportedSocketOperationError"] = int64(QAbstractSocket__UnsupportedSocketOperationError)
	qt.EnumMap["network.QAbstractSocket__UnfinishedSocketOperationError"] = int64(QAbstractSocket__UnfinishedSocketOperationError)
	qt.EnumMap["network.QAbstractSocket__ProxyAuthenticationRequiredError"] = int64(QAbstractSocket__ProxyAuthenticationRequiredError)
	qt.EnumMap["network.QAbstractSocket__SslHandshakeFailedError"] = int64(QAbstractSocket__SslHandshakeFailedError)
	qt.EnumMap["network.QAbstractSocket__ProxyConnectionRefusedError"] = int64(QAbstractSocket__ProxyConnectionRefusedError)
	qt.EnumMap["network.QAbstractSocket__ProxyConnectionClosedError"] = int64(QAbstractSocket__ProxyConnectionClosedError)
	qt.EnumMap["network.QAbstractSocket__ProxyConnectionTimeoutError"] = int64(QAbstractSocket__ProxyConnectionTimeoutError)
	qt.EnumMap["network.QAbstractSocket__ProxyNotFoundError"] = int64(QAbstractSocket__ProxyNotFoundError)
	qt.EnumMap["network.QAbstractSocket__ProxyProtocolError"] = int64(QAbstractSocket__ProxyProtocolError)
	qt.EnumMap["network.QAbstractSocket__OperationError"] = int64(QAbstractSocket__OperationError)
	qt.EnumMap["network.QAbstractSocket__SslInternalError"] = int64(QAbstractSocket__SslInternalError)
	qt.EnumMap["network.QAbstractSocket__SslInvalidUserDataError"] = int64(QAbstractSocket__SslInvalidUserDataError)
	qt.EnumMap["network.QAbstractSocket__TemporaryError"] = int64(QAbstractSocket__TemporaryError)
	qt.EnumMap["network.QAbstractSocket__UnknownSocketError"] = int64(QAbstractSocket__UnknownSocketError)
	qt.EnumMap["network.QAbstractSocket__UnconnectedState"] = int64(QAbstractSocket__UnconnectedState)
	qt.EnumMap["network.QAbstractSocket__HostLookupState"] = int64(QAbstractSocket__HostLookupState)
	qt.EnumMap["network.QAbstractSocket__ConnectingState"] = int64(QAbstractSocket__ConnectingState)
	qt.EnumMap["network.QAbstractSocket__ConnectedState"] = int64(QAbstractSocket__ConnectedState)
	qt.EnumMap["network.QAbstractSocket__BoundState"] = int64(QAbstractSocket__BoundState)
	qt.EnumMap["network.QAbstractSocket__ListeningState"] = int64(QAbstractSocket__ListeningState)
	qt.EnumMap["network.QAbstractSocket__ClosingState"] = int64(QAbstractSocket__ClosingState)
	qt.EnumMap["network.QAbstractSocket__LowDelayOption"] = int64(QAbstractSocket__LowDelayOption)
	qt.EnumMap["network.QAbstractSocket__KeepAliveOption"] = int64(QAbstractSocket__KeepAliveOption)
	qt.EnumMap["network.QAbstractSocket__MulticastTtlOption"] = int64(QAbstractSocket__MulticastTtlOption)
	qt.EnumMap["network.QAbstractSocket__MulticastLoopbackOption"] = int64(QAbstractSocket__MulticastLoopbackOption)
	qt.EnumMap["network.QAbstractSocket__TypeOfServiceOption"] = int64(QAbstractSocket__TypeOfServiceOption)
	qt.EnumMap["network.QAbstractSocket__SendBufferSizeSocketOption"] = int64(QAbstractSocket__SendBufferSizeSocketOption)
	qt.EnumMap["network.QAbstractSocket__ReceiveBufferSizeSocketOption"] = int64(QAbstractSocket__ReceiveBufferSizeSocketOption)
	qt.EnumMap["network.QAbstractSocket__PathMtuSocketOption"] = int64(QAbstractSocket__PathMtuSocketOption)
	qt.EnumMap["network.QAbstractSocket__DefaultForPlatform"] = int64(QAbstractSocket__DefaultForPlatform)
	qt.EnumMap["network.QAbstractSocket__ShareAddress"] = int64(QAbstractSocket__ShareAddress)
	qt.EnumMap["network.QAbstractSocket__DontShareAddress"] = int64(QAbstractSocket__DontShareAddress)
	qt.EnumMap["network.QAbstractSocket__ReuseAddressHint"] = int64(QAbstractSocket__ReuseAddressHint)
	qt.EnumMap["network.QAbstractSocket__PauseNever"] = int64(QAbstractSocket__PauseNever)
	qt.EnumMap["network.QAbstractSocket__PauseOnSslErrors"] = int64(QAbstractSocket__PauseOnSslErrors)
	qt.ItfMap["network.QAuthenticator_ITF"] = QAuthenticator{}
	qt.FuncMap["network.NewQAuthenticator"] = NewQAuthenticator
	qt.FuncMap["network.NewQAuthenticator2"] = NewQAuthenticator2
	qt.ItfMap["network.QDnsDomainNameRecord_ITF"] = QDnsDomainNameRecord{}
	qt.FuncMap["network.NewQDnsDomainNameRecord"] = NewQDnsDomainNameRecord
	qt.FuncMap["network.NewQDnsDomainNameRecord2"] = NewQDnsDomainNameRecord2
	qt.ItfMap["network.QDnsHostAddressRecord_ITF"] = QDnsHostAddressRecord{}
	qt.FuncMap["network.NewQDnsHostAddressRecord"] = NewQDnsHostAddressRecord
	qt.FuncMap["network.NewQDnsHostAddressRecord2"] = NewQDnsHostAddressRecord2
	qt.ItfMap["network.QDnsLookup_ITF"] = QDnsLookup{}
	qt.FuncMap["network.NewQDnsLookup"] = NewQDnsLookup
	qt.FuncMap["network.NewQDnsLookup2"] = NewQDnsLookup2
	qt.FuncMap["network.NewQDnsLookup3"] = NewQDnsLookup3
	qt.EnumMap["network.QDnsLookup__NoError"] = int64(QDnsLookup__NoError)
	qt.EnumMap["network.QDnsLookup__ResolverError"] = int64(QDnsLookup__ResolverError)
	qt.EnumMap["network.QDnsLookup__OperationCancelledError"] = int64(QDnsLookup__OperationCancelledError)
	qt.EnumMap["network.QDnsLookup__InvalidRequestError"] = int64(QDnsLookup__InvalidRequestError)
	qt.EnumMap["network.QDnsLookup__InvalidReplyError"] = int64(QDnsLookup__InvalidReplyError)
	qt.EnumMap["network.QDnsLookup__ServerFailureError"] = int64(QDnsLookup__ServerFailureError)
	qt.EnumMap["network.QDnsLookup__ServerRefusedError"] = int64(QDnsLookup__ServerRefusedError)
	qt.EnumMap["network.QDnsLookup__NotFoundError"] = int64(QDnsLookup__NotFoundError)
	qt.EnumMap["network.QDnsLookup__A"] = int64(QDnsLookup__A)
	qt.EnumMap["network.QDnsLookup__AAAA"] = int64(QDnsLookup__AAAA)
	qt.EnumMap["network.QDnsLookup__ANY"] = int64(QDnsLookup__ANY)
	qt.EnumMap["network.QDnsLookup__CNAME"] = int64(QDnsLookup__CNAME)
	qt.EnumMap["network.QDnsLookup__MX"] = int64(QDnsLookup__MX)
	qt.EnumMap["network.QDnsLookup__NS"] = int64(QDnsLookup__NS)
	qt.EnumMap["network.QDnsLookup__PTR"] = int64(QDnsLookup__PTR)
	qt.EnumMap["network.QDnsLookup__SRV"] = int64(QDnsLookup__SRV)
	qt.EnumMap["network.QDnsLookup__TXT"] = int64(QDnsLookup__TXT)
	qt.ItfMap["network.QDnsMailExchangeRecord_ITF"] = QDnsMailExchangeRecord{}
	qt.FuncMap["network.NewQDnsMailExchangeRecord"] = NewQDnsMailExchangeRecord
	qt.FuncMap["network.NewQDnsMailExchangeRecord2"] = NewQDnsMailExchangeRecord2
	qt.ItfMap["network.QDnsServiceRecord_ITF"] = QDnsServiceRecord{}
	qt.FuncMap["network.NewQDnsServiceRecord"] = NewQDnsServiceRecord
	qt.FuncMap["network.NewQDnsServiceRecord2"] = NewQDnsServiceRecord2
	qt.ItfMap["network.QDnsTextRecord_ITF"] = QDnsTextRecord{}
	qt.FuncMap["network.NewQDnsTextRecord"] = NewQDnsTextRecord
	qt.FuncMap["network.NewQDnsTextRecord2"] = NewQDnsTextRecord2
	qt.EnumMap["network.QDtls__NoError"] = int64(QDtls__NoError)
	qt.EnumMap["network.QDtls__InvalidInputParameters"] = int64(QDtls__InvalidInputParameters)
	qt.EnumMap["network.QDtls__InvalidOperation"] = int64(QDtls__InvalidOperation)
	qt.EnumMap["network.QDtls__UnderlyingSocketError"] = int64(QDtls__UnderlyingSocketError)
	qt.EnumMap["network.QDtls__RemoteClosedConnectionError"] = int64(QDtls__RemoteClosedConnectionError)
	qt.EnumMap["network.QDtls__PeerVerificationError"] = int64(QDtls__PeerVerificationError)
	qt.EnumMap["network.QDtls__TlsInitializationError"] = int64(QDtls__TlsInitializationError)
	qt.EnumMap["network.QDtls__TlsFatalError"] = int64(QDtls__TlsFatalError)
	qt.EnumMap["network.QDtls__TlsNonFatalError"] = int64(QDtls__TlsNonFatalError)
	qt.EnumMap["network.QDtls__HandshakeNotStarted"] = int64(QDtls__HandshakeNotStarted)
	qt.EnumMap["network.QDtls__HandshakeInProgress"] = int64(QDtls__HandshakeInProgress)
	qt.EnumMap["network.QDtls__PeerVerificationFailed"] = int64(QDtls__PeerVerificationFailed)
	qt.EnumMap["network.QDtls__HandshakeComplete"] = int64(QDtls__HandshakeComplete)
	qt.ItfMap["network.QHostAddress_ITF"] = QHostAddress{}
	qt.FuncMap["network.NewQHostAddress"] = NewQHostAddress
	qt.FuncMap["network.NewQHostAddress2"] = NewQHostAddress2
	qt.FuncMap["network.NewQHostAddress3"] = NewQHostAddress3
	qt.FuncMap["network.NewQHostAddress4"] = NewQHostAddress4
	qt.FuncMap["network.NewQHostAddress7"] = NewQHostAddress7
	qt.FuncMap["network.NewQHostAddress8"] = NewQHostAddress8
	qt.FuncMap["network.NewQHostAddress9"] = NewQHostAddress9
	qt.EnumMap["network.QHostAddress__Null"] = int64(QHostAddress__Null)
	qt.EnumMap["network.QHostAddress__Broadcast"] = int64(QHostAddress__Broadcast)
	qt.EnumMap["network.QHostAddress__LocalHost"] = int64(QHostAddress__LocalHost)
	qt.EnumMap["network.QHostAddress__LocalHostIPv6"] = int64(QHostAddress__LocalHostIPv6)
	qt.EnumMap["network.QHostAddress__Any"] = int64(QHostAddress__Any)
	qt.EnumMap["network.QHostAddress__AnyIPv6"] = int64(QHostAddress__AnyIPv6)
	qt.EnumMap["network.QHostAddress__AnyIPv4"] = int64(QHostAddress__AnyIPv4)
	qt.EnumMap["network.QHostAddress__ConvertV4MappedToIPv4"] = int64(QHostAddress__ConvertV4MappedToIPv4)
	qt.EnumMap["network.QHostAddress__ConvertV4CompatToIPv4"] = int64(QHostAddress__ConvertV4CompatToIPv4)
	qt.EnumMap["network.QHostAddress__ConvertUnspecifiedAddress"] = int64(QHostAddress__ConvertUnspecifiedAddress)
	qt.EnumMap["network.QHostAddress__ConvertLocalHost"] = int64(QHostAddress__ConvertLocalHost)
	qt.EnumMap["network.QHostAddress__TolerantConversion"] = int64(QHostAddress__TolerantConversion)
	qt.EnumMap["network.QHostAddress__StrictConversion"] = int64(QHostAddress__StrictConversion)
	qt.ItfMap["network.QHostInfo_ITF"] = QHostInfo{}
	qt.FuncMap["network.NewQHostInfo"] = NewQHostInfo
	qt.FuncMap["network.NewQHostInfo2"] = NewQHostInfo2
	qt.FuncMap["network.QHostInfo_AbortHostLookup"] = QHostInfo_AbortHostLookup
	qt.FuncMap["network.QHostInfo_FromName"] = QHostInfo_FromName
	qt.FuncMap["network.QHostInfo_LocalDomainName"] = QHostInfo_LocalDomainName
	qt.FuncMap["network.QHostInfo_LocalHostName"] = QHostInfo_LocalHostName
	qt.FuncMap["network.QHostInfo_LookupHost"] = QHostInfo_LookupHost
	qt.EnumMap["network.QHostInfo__NoError"] = int64(QHostInfo__NoError)
	qt.EnumMap["network.QHostInfo__HostNotFound"] = int64(QHostInfo__HostNotFound)
	qt.EnumMap["network.QHostInfo__UnknownError"] = int64(QHostInfo__UnknownError)
	qt.ItfMap["network.QHstsPolicy_ITF"] = QHstsPolicy{}
	qt.FuncMap["network.NewQHstsPolicy"] = NewQHstsPolicy
	qt.FuncMap["network.NewQHstsPolicy2"] = NewQHstsPolicy2
	qt.FuncMap["network.NewQHstsPolicy3"] = NewQHstsPolicy3
	qt.EnumMap["network.QHstsPolicy__IncludeSubDomains"] = int64(QHstsPolicy__IncludeSubDomains)
	qt.ItfMap["network.QHttpMultiPart_ITF"] = QHttpMultiPart{}
	qt.FuncMap["network.NewQHttpMultiPart"] = NewQHttpMultiPart
	qt.FuncMap["network.NewQHttpMultiPart2"] = NewQHttpMultiPart2
	qt.EnumMap["network.QHttpMultiPart__MixedType"] = int64(QHttpMultiPart__MixedType)
	qt.EnumMap["network.QHttpMultiPart__RelatedType"] = int64(QHttpMultiPart__RelatedType)
	qt.EnumMap["network.QHttpMultiPart__FormDataType"] = int64(QHttpMultiPart__FormDataType)
	qt.EnumMap["network.QHttpMultiPart__AlternativeType"] = int64(QHttpMultiPart__AlternativeType)
	qt.ItfMap["network.QHttpPart_ITF"] = QHttpPart{}
	qt.FuncMap["network.NewQHttpPart"] = NewQHttpPart
	qt.FuncMap["network.NewQHttpPart2"] = NewQHttpPart2
	qt.ItfMap["network.QIPv6Address_ITF"] = QIPv6Address{}
	qt.ItfMap["network.QLocalServer_ITF"] = QLocalServer{}
	qt.FuncMap["network.NewQLocalServer"] = NewQLocalServer
	qt.FuncMap["network.QLocalServer_RemoveServer"] = QLocalServer_RemoveServer
	qt.EnumMap["network.QLocalServer__NoOptions"] = int64(QLocalServer__NoOptions)
	qt.EnumMap["network.QLocalServer__UserAccessOption"] = int64(QLocalServer__UserAccessOption)
	qt.EnumMap["network.QLocalServer__GroupAccessOption"] = int64(QLocalServer__GroupAccessOption)
	qt.EnumMap["network.QLocalServer__OtherAccessOption"] = int64(QLocalServer__OtherAccessOption)
	qt.EnumMap["network.QLocalServer__WorldAccessOption"] = int64(QLocalServer__WorldAccessOption)
	qt.ItfMap["network.QLocalSocket_ITF"] = QLocalSocket{}
	qt.FuncMap["network.NewQLocalSocket"] = NewQLocalSocket
	qt.EnumMap["network.QLocalSocket__ConnectionRefusedError"] = int64(QLocalSocket__ConnectionRefusedError)
	qt.EnumMap["network.QLocalSocket__PeerClosedError"] = int64(QLocalSocket__PeerClosedError)
	qt.EnumMap["network.QLocalSocket__ServerNotFoundError"] = int64(QLocalSocket__ServerNotFoundError)
	qt.EnumMap["network.QLocalSocket__SocketAccessError"] = int64(QLocalSocket__SocketAccessError)
	qt.EnumMap["network.QLocalSocket__SocketResourceError"] = int64(QLocalSocket__SocketResourceError)
	qt.EnumMap["network.QLocalSocket__SocketTimeoutError"] = int64(QLocalSocket__SocketTimeoutError)
	qt.EnumMap["network.QLocalSocket__DatagramTooLargeError"] = int64(QLocalSocket__DatagramTooLargeError)
	qt.EnumMap["network.QLocalSocket__ConnectionError"] = int64(QLocalSocket__ConnectionError)
	qt.EnumMap["network.QLocalSocket__UnsupportedSocketOperationError"] = int64(QLocalSocket__UnsupportedSocketOperationError)
	qt.EnumMap["network.QLocalSocket__UnknownSocketError"] = int64(QLocalSocket__UnknownSocketError)
	qt.EnumMap["network.QLocalSocket__OperationError"] = int64(QLocalSocket__OperationError)
	qt.EnumMap["network.QLocalSocket__UnconnectedState"] = int64(QLocalSocket__UnconnectedState)
	qt.EnumMap["network.QLocalSocket__ConnectingState"] = int64(QLocalSocket__ConnectingState)
	qt.EnumMap["network.QLocalSocket__ConnectedState"] = int64(QLocalSocket__ConnectedState)
	qt.EnumMap["network.QLocalSocket__ClosingState"] = int64(QLocalSocket__ClosingState)
	qt.ItfMap["network.QNetworkAccessManager_ITF"] = QNetworkAccessManager{}
	qt.FuncMap["network.NewQNetworkAccessManager"] = NewQNetworkAccessManager
	qt.EnumMap["network.QNetworkAccessManager__HeadOperation"] = int64(QNetworkAccessManager__HeadOperation)
	qt.EnumMap["network.QNetworkAccessManager__GetOperation"] = int64(QNetworkAccessManager__GetOperation)
	qt.EnumMap["network.QNetworkAccessManager__PutOperation"] = int64(QNetworkAccessManager__PutOperation)
	qt.EnumMap["network.QNetworkAccessManager__PostOperation"] = int64(QNetworkAccessManager__PostOperation)
	qt.EnumMap["network.QNetworkAccessManager__DeleteOperation"] = int64(QNetworkAccessManager__DeleteOperation)
	qt.EnumMap["network.QNetworkAccessManager__CustomOperation"] = int64(QNetworkAccessManager__CustomOperation)
	qt.EnumMap["network.QNetworkAccessManager__UnknownOperation"] = int64(QNetworkAccessManager__UnknownOperation)
	qt.EnumMap["network.QNetworkAccessManager__UnknownAccessibility"] = int64(QNetworkAccessManager__UnknownAccessibility)
	qt.EnumMap["network.QNetworkAccessManager__NotAccessible"] = int64(QNetworkAccessManager__NotAccessible)
	qt.EnumMap["network.QNetworkAccessManager__Accessible"] = int64(QNetworkAccessManager__Accessible)
	qt.ItfMap["network.QNetworkAddressEntry_ITF"] = QNetworkAddressEntry{}
	qt.FuncMap["network.NewQNetworkAddressEntry"] = NewQNetworkAddressEntry
	qt.FuncMap["network.NewQNetworkAddressEntry2"] = NewQNetworkAddressEntry2
	qt.EnumMap["network.QNetworkAddressEntry__DnsEligibilityUnknown"] = int64(QNetworkAddressEntry__DnsEligibilityUnknown)
	qt.EnumMap["network.QNetworkAddressEntry__DnsIneligible"] = int64(QNetworkAddressEntry__DnsIneligible)
	qt.EnumMap["network.QNetworkAddressEntry__DnsEligible"] = int64(QNetworkAddressEntry__DnsEligible)
	qt.ItfMap["network.QNetworkCacheMetaData_ITF"] = QNetworkCacheMetaData{}
	qt.FuncMap["network.NewQNetworkCacheMetaData"] = NewQNetworkCacheMetaData
	qt.FuncMap["network.NewQNetworkCacheMetaData2"] = NewQNetworkCacheMetaData2
	qt.ItfMap["network.QNetworkConfiguration_ITF"] = QNetworkConfiguration{}
	qt.FuncMap["network.NewQNetworkConfiguration"] = NewQNetworkConfiguration
	qt.FuncMap["network.NewQNetworkConfiguration2"] = NewQNetworkConfiguration2
	qt.EnumMap["network.QNetworkConfiguration__InternetAccessPoint"] = int64(QNetworkConfiguration__InternetAccessPoint)
	qt.EnumMap["network.QNetworkConfiguration__ServiceNetwork"] = int64(QNetworkConfiguration__ServiceNetwork)
	qt.EnumMap["network.QNetworkConfiguration__UserChoice"] = int64(QNetworkConfiguration__UserChoice)
	qt.EnumMap["network.QNetworkConfiguration__Invalid"] = int64(QNetworkConfiguration__Invalid)
	qt.EnumMap["network.QNetworkConfiguration__UnknownPurpose"] = int64(QNetworkConfiguration__UnknownPurpose)
	qt.EnumMap["network.QNetworkConfiguration__PublicPurpose"] = int64(QNetworkConfiguration__PublicPurpose)
	qt.EnumMap["network.QNetworkConfiguration__PrivatePurpose"] = int64(QNetworkConfiguration__PrivatePurpose)
	qt.EnumMap["network.QNetworkConfiguration__ServiceSpecificPurpose"] = int64(QNetworkConfiguration__ServiceSpecificPurpose)
	qt.EnumMap["network.QNetworkConfiguration__Undefined"] = int64(QNetworkConfiguration__Undefined)
	qt.EnumMap["network.QNetworkConfiguration__Defined"] = int64(QNetworkConfiguration__Defined)
	qt.EnumMap["network.QNetworkConfiguration__Discovered"] = int64(QNetworkConfiguration__Discovered)
	qt.EnumMap["network.QNetworkConfiguration__Active"] = int64(QNetworkConfiguration__Active)
	qt.EnumMap["network.QNetworkConfiguration__BearerUnknown"] = int64(QNetworkConfiguration__BearerUnknown)
	qt.EnumMap["network.QNetworkConfiguration__BearerEthernet"] = int64(QNetworkConfiguration__BearerEthernet)
	qt.EnumMap["network.QNetworkConfiguration__BearerWLAN"] = int64(QNetworkConfiguration__BearerWLAN)
	qt.EnumMap["network.QNetworkConfiguration__Bearer2G"] = int64(QNetworkConfiguration__Bearer2G)
	qt.EnumMap["network.QNetworkConfiguration__BearerCDMA2000"] = int64(QNetworkConfiguration__BearerCDMA2000)
	qt.EnumMap["network.QNetworkConfiguration__BearerWCDMA"] = int64(QNetworkConfiguration__BearerWCDMA)
	qt.EnumMap["network.QNetworkConfiguration__BearerHSPA"] = int64(QNetworkConfiguration__BearerHSPA)
	qt.EnumMap["network.QNetworkConfiguration__BearerBluetooth"] = int64(QNetworkConfiguration__BearerBluetooth)
	qt.EnumMap["network.QNetworkConfiguration__BearerWiMAX"] = int64(QNetworkConfiguration__BearerWiMAX)
	qt.EnumMap["network.QNetworkConfiguration__BearerEVDO"] = int64(QNetworkConfiguration__BearerEVDO)
	qt.EnumMap["network.QNetworkConfiguration__BearerLTE"] = int64(QNetworkConfiguration__BearerLTE)
	qt.EnumMap["network.QNetworkConfiguration__Bearer3G"] = int64(QNetworkConfiguration__Bearer3G)
	qt.EnumMap["network.QNetworkConfiguration__Bearer4G"] = int64(QNetworkConfiguration__Bearer4G)
	qt.ItfMap["network.QNetworkConfigurationManager_ITF"] = QNetworkConfigurationManager{}
	qt.FuncMap["network.NewQNetworkConfigurationManager"] = NewQNetworkConfigurationManager
	qt.EnumMap["network.QNetworkConfigurationManager__CanStartAndStopInterfaces"] = int64(QNetworkConfigurationManager__CanStartAndStopInterfaces)
	qt.EnumMap["network.QNetworkConfigurationManager__DirectConnectionRouting"] = int64(QNetworkConfigurationManager__DirectConnectionRouting)
	qt.EnumMap["network.QNetworkConfigurationManager__SystemSessionSupport"] = int64(QNetworkConfigurationManager__SystemSessionSupport)
	qt.EnumMap["network.QNetworkConfigurationManager__ApplicationLevelRoaming"] = int64(QNetworkConfigurationManager__ApplicationLevelRoaming)
	qt.EnumMap["network.QNetworkConfigurationManager__ForcedRoaming"] = int64(QNetworkConfigurationManager__ForcedRoaming)
	qt.EnumMap["network.QNetworkConfigurationManager__DataStatistics"] = int64(QNetworkConfigurationManager__DataStatistics)
	qt.EnumMap["network.QNetworkConfigurationManager__NetworkSessionRequired"] = int64(QNetworkConfigurationManager__NetworkSessionRequired)
	qt.ItfMap["network.QNetworkCookie_ITF"] = QNetworkCookie{}
	qt.FuncMap["network.NewQNetworkCookie"] = NewQNetworkCookie
	qt.FuncMap["network.NewQNetworkCookie2"] = NewQNetworkCookie2
	qt.FuncMap["network.QNetworkCookie_ParseCookies"] = QNetworkCookie_ParseCookies
	qt.EnumMap["network.QNetworkCookie__NameAndValueOnly"] = int64(QNetworkCookie__NameAndValueOnly)
	qt.EnumMap["network.QNetworkCookie__Full"] = int64(QNetworkCookie__Full)
	qt.ItfMap["network.QNetworkCookieJar_ITF"] = QNetworkCookieJar{}
	qt.FuncMap["network.NewQNetworkCookieJar"] = NewQNetworkCookieJar
	qt.ItfMap["network.QNetworkDatagram_ITF"] = QNetworkDatagram{}
	qt.FuncMap["network.NewQNetworkDatagram"] = NewQNetworkDatagram
	qt.FuncMap["network.NewQNetworkDatagram2"] = NewQNetworkDatagram2
	qt.FuncMap["network.NewQNetworkDatagram3"] = NewQNetworkDatagram3
	qt.ItfMap["network.QNetworkDiskCache_ITF"] = QNetworkDiskCache{}
	qt.FuncMap["network.NewQNetworkDiskCache"] = NewQNetworkDiskCache
	qt.ItfMap["network.QNetworkInterface_ITF"] = QNetworkInterface{}
	qt.FuncMap["network.NewQNetworkInterface"] = NewQNetworkInterface
	qt.FuncMap["network.NewQNetworkInterface2"] = NewQNetworkInterface2
	qt.FuncMap["network.QNetworkInterface_AllAddresses"] = QNetworkInterface_AllAddresses
	qt.FuncMap["network.QNetworkInterface_AllInterfaces"] = QNetworkInterface_AllInterfaces
	qt.FuncMap["network.QNetworkInterface_InterfaceFromIndex"] = QNetworkInterface_InterfaceFromIndex
	qt.FuncMap["network.QNetworkInterface_InterfaceFromName"] = QNetworkInterface_InterfaceFromName
	qt.FuncMap["network.QNetworkInterface_InterfaceIndexFromName"] = QNetworkInterface_InterfaceIndexFromName
	qt.FuncMap["network.QNetworkInterface_InterfaceNameFromIndex"] = QNetworkInterface_InterfaceNameFromIndex
	qt.EnumMap["network.QNetworkInterface__Loopback"] = int64(QNetworkInterface__Loopback)
	qt.EnumMap["network.QNetworkInterface__Virtual"] = int64(QNetworkInterface__Virtual)
	qt.EnumMap["network.QNetworkInterface__Ethernet"] = int64(QNetworkInterface__Ethernet)
	qt.EnumMap["network.QNetworkInterface__Slip"] = int64(QNetworkInterface__Slip)
	qt.EnumMap["network.QNetworkInterface__CanBus"] = int64(QNetworkInterface__CanBus)
	qt.EnumMap["network.QNetworkInterface__Ppp"] = int64(QNetworkInterface__Ppp)
	qt.EnumMap["network.QNetworkInterface__Fddi"] = int64(QNetworkInterface__Fddi)
	qt.EnumMap["network.QNetworkInterface__Wifi"] = int64(QNetworkInterface__Wifi)
	qt.EnumMap["network.QNetworkInterface__Ieee80211"] = int64(QNetworkInterface__Ieee80211)
	qt.EnumMap["network.QNetworkInterface__Phonet"] = int64(QNetworkInterface__Phonet)
	qt.EnumMap["network.QNetworkInterface__Ieee802154"] = int64(QNetworkInterface__Ieee802154)
	qt.EnumMap["network.QNetworkInterface__SixLoWPAN"] = int64(QNetworkInterface__SixLoWPAN)
	qt.EnumMap["network.QNetworkInterface__Ieee80216"] = int64(QNetworkInterface__Ieee80216)
	qt.EnumMap["network.QNetworkInterface__Ieee1394"] = int64(QNetworkInterface__Ieee1394)
	qt.EnumMap["network.QNetworkInterface__Unknown"] = int64(QNetworkInterface__Unknown)
	qt.EnumMap["network.QNetworkInterface__IsUp"] = int64(QNetworkInterface__IsUp)
	qt.EnumMap["network.QNetworkInterface__IsRunning"] = int64(QNetworkInterface__IsRunning)
	qt.EnumMap["network.QNetworkInterface__CanBroadcast"] = int64(QNetworkInterface__CanBroadcast)
	qt.EnumMap["network.QNetworkInterface__IsLoopBack"] = int64(QNetworkInterface__IsLoopBack)
	qt.EnumMap["network.QNetworkInterface__IsPointToPoint"] = int64(QNetworkInterface__IsPointToPoint)
	qt.EnumMap["network.QNetworkInterface__CanMulticast"] = int64(QNetworkInterface__CanMulticast)
	qt.ItfMap["network.QNetworkProxy_ITF"] = QNetworkProxy{}
	qt.FuncMap["network.NewQNetworkProxy"] = NewQNetworkProxy
	qt.FuncMap["network.NewQNetworkProxy2"] = NewQNetworkProxy2
	qt.FuncMap["network.NewQNetworkProxy3"] = NewQNetworkProxy3
	qt.FuncMap["network.QNetworkProxy_ApplicationProxy"] = QNetworkProxy_ApplicationProxy
	qt.FuncMap["network.QNetworkProxy_SetApplicationProxy"] = QNetworkProxy_SetApplicationProxy
	qt.EnumMap["network.QNetworkProxy__DefaultProxy"] = int64(QNetworkProxy__DefaultProxy)
	qt.EnumMap["network.QNetworkProxy__Socks5Proxy"] = int64(QNetworkProxy__Socks5Proxy)
	qt.EnumMap["network.QNetworkProxy__NoProxy"] = int64(QNetworkProxy__NoProxy)
	qt.EnumMap["network.QNetworkProxy__HttpProxy"] = int64(QNetworkProxy__HttpProxy)
	qt.EnumMap["network.QNetworkProxy__HttpCachingProxy"] = int64(QNetworkProxy__HttpCachingProxy)
	qt.EnumMap["network.QNetworkProxy__FtpCachingProxy"] = int64(QNetworkProxy__FtpCachingProxy)
	qt.EnumMap["network.QNetworkProxy__TunnelingCapability"] = int64(QNetworkProxy__TunnelingCapability)
	qt.EnumMap["network.QNetworkProxy__ListeningCapability"] = int64(QNetworkProxy__ListeningCapability)
	qt.EnumMap["network.QNetworkProxy__UdpTunnelingCapability"] = int64(QNetworkProxy__UdpTunnelingCapability)
	qt.EnumMap["network.QNetworkProxy__CachingCapability"] = int64(QNetworkProxy__CachingCapability)
	qt.EnumMap["network.QNetworkProxy__HostNameLookupCapability"] = int64(QNetworkProxy__HostNameLookupCapability)
	qt.EnumMap["network.QNetworkProxy__SctpTunnelingCapability"] = int64(QNetworkProxy__SctpTunnelingCapability)
	qt.EnumMap["network.QNetworkProxy__SctpListeningCapability"] = int64(QNetworkProxy__SctpListeningCapability)
	qt.ItfMap["network.QNetworkProxyFactory_ITF"] = QNetworkProxyFactory{}
	qt.FuncMap["network.NewQNetworkProxyFactory"] = NewQNetworkProxyFactory
	qt.FuncMap["network.QNetworkProxyFactory_ProxyForQuery"] = QNetworkProxyFactory_ProxyForQuery
	qt.FuncMap["network.QNetworkProxyFactory_SetApplicationProxyFactory"] = QNetworkProxyFactory_SetApplicationProxyFactory
	qt.FuncMap["network.QNetworkProxyFactory_SetUseSystemConfiguration"] = QNetworkProxyFactory_SetUseSystemConfiguration
	qt.FuncMap["network.QNetworkProxyFactory_SystemProxyForQuery"] = QNetworkProxyFactory_SystemProxyForQuery
	qt.FuncMap["network.QNetworkProxyFactory_UsesSystemConfiguration"] = QNetworkProxyFactory_UsesSystemConfiguration
	qt.ItfMap["network.QNetworkProxyQuery_ITF"] = QNetworkProxyQuery{}
	qt.FuncMap["network.NewQNetworkProxyQuery"] = NewQNetworkProxyQuery
	qt.FuncMap["network.NewQNetworkProxyQuery2"] = NewQNetworkProxyQuery2
	qt.FuncMap["network.NewQNetworkProxyQuery3"] = NewQNetworkProxyQuery3
	qt.FuncMap["network.NewQNetworkProxyQuery4"] = NewQNetworkProxyQuery4
	qt.FuncMap["network.NewQNetworkProxyQuery8"] = NewQNetworkProxyQuery8
	qt.EnumMap["network.QNetworkProxyQuery__TcpSocket"] = int64(QNetworkProxyQuery__TcpSocket)
	qt.EnumMap["network.QNetworkProxyQuery__UdpSocket"] = int64(QNetworkProxyQuery__UdpSocket)
	qt.EnumMap["network.QNetworkProxyQuery__SctpSocket"] = int64(QNetworkProxyQuery__SctpSocket)
	qt.EnumMap["network.QNetworkProxyQuery__TcpServer"] = int64(QNetworkProxyQuery__TcpServer)
	qt.EnumMap["network.QNetworkProxyQuery__UrlRequest"] = int64(QNetworkProxyQuery__UrlRequest)
	qt.EnumMap["network.QNetworkProxyQuery__SctpServer"] = int64(QNetworkProxyQuery__SctpServer)
	qt.ItfMap["network.QNetworkReply_ITF"] = QNetworkReply{}
	qt.FuncMap["network.NewQNetworkReply"] = NewQNetworkReply
	qt.EnumMap["network.QNetworkReply__NoError"] = int64(QNetworkReply__NoError)
	qt.EnumMap["network.QNetworkReply__ConnectionRefusedError"] = int64(QNetworkReply__ConnectionRefusedError)
	qt.EnumMap["network.QNetworkReply__RemoteHostClosedError"] = int64(QNetworkReply__RemoteHostClosedError)
	qt.EnumMap["network.QNetworkReply__HostNotFoundError"] = int64(QNetworkReply__HostNotFoundError)
	qt.EnumMap["network.QNetworkReply__TimeoutError"] = int64(QNetworkReply__TimeoutError)
	qt.EnumMap["network.QNetworkReply__OperationCanceledError"] = int64(QNetworkReply__OperationCanceledError)
	qt.EnumMap["network.QNetworkReply__SslHandshakeFailedError"] = int64(QNetworkReply__SslHandshakeFailedError)
	qt.EnumMap["network.QNetworkReply__TemporaryNetworkFailureError"] = int64(QNetworkReply__TemporaryNetworkFailureError)
	qt.EnumMap["network.QNetworkReply__NetworkSessionFailedError"] = int64(QNetworkReply__NetworkSessionFailedError)
	qt.EnumMap["network.QNetworkReply__BackgroundRequestNotAllowedError"] = int64(QNetworkReply__BackgroundRequestNotAllowedError)
	qt.EnumMap["network.QNetworkReply__TooManyRedirectsError"] = int64(QNetworkReply__TooManyRedirectsError)
	qt.EnumMap["network.QNetworkReply__InsecureRedirectError"] = int64(QNetworkReply__InsecureRedirectError)
	qt.EnumMap["network.QNetworkReply__UnknownNetworkError"] = int64(QNetworkReply__UnknownNetworkError)
	qt.EnumMap["network.QNetworkReply__ProxyConnectionRefusedError"] = int64(QNetworkReply__ProxyConnectionRefusedError)
	qt.EnumMap["network.QNetworkReply__ProxyConnectionClosedError"] = int64(QNetworkReply__ProxyConnectionClosedError)
	qt.EnumMap["network.QNetworkReply__ProxyNotFoundError"] = int64(QNetworkReply__ProxyNotFoundError)
	qt.EnumMap["network.QNetworkReply__ProxyTimeoutError"] = int64(QNetworkReply__ProxyTimeoutError)
	qt.EnumMap["network.QNetworkReply__ProxyAuthenticationRequiredError"] = int64(QNetworkReply__ProxyAuthenticationRequiredError)
	qt.EnumMap["network.QNetworkReply__UnknownProxyError"] = int64(QNetworkReply__UnknownProxyError)
	qt.EnumMap["network.QNetworkReply__ContentAccessDenied"] = int64(QNetworkReply__ContentAccessDenied)
	qt.EnumMap["network.QNetworkReply__ContentOperationNotPermittedError"] = int64(QNetworkReply__ContentOperationNotPermittedError)
	qt.EnumMap["network.QNetworkReply__ContentNotFoundError"] = int64(QNetworkReply__ContentNotFoundError)
	qt.EnumMap["network.QNetworkReply__AuthenticationRequiredError"] = int64(QNetworkReply__AuthenticationRequiredError)
	qt.EnumMap["network.QNetworkReply__ContentReSendError"] = int64(QNetworkReply__ContentReSendError)
	qt.EnumMap["network.QNetworkReply__ContentConflictError"] = int64(QNetworkReply__ContentConflictError)
	qt.EnumMap["network.QNetworkReply__ContentGoneError"] = int64(QNetworkReply__ContentGoneError)
	qt.EnumMap["network.QNetworkReply__UnknownContentError"] = int64(QNetworkReply__UnknownContentError)
	qt.EnumMap["network.QNetworkReply__ProtocolUnknownError"] = int64(QNetworkReply__ProtocolUnknownError)
	qt.EnumMap["network.QNetworkReply__ProtocolInvalidOperationError"] = int64(QNetworkReply__ProtocolInvalidOperationError)
	qt.EnumMap["network.QNetworkReply__ProtocolFailure"] = int64(QNetworkReply__ProtocolFailure)
	qt.EnumMap["network.QNetworkReply__InternalServerError"] = int64(QNetworkReply__InternalServerError)
	qt.EnumMap["network.QNetworkReply__OperationNotImplementedError"] = int64(QNetworkReply__OperationNotImplementedError)
	qt.EnumMap["network.QNetworkReply__ServiceUnavailableError"] = int64(QNetworkReply__ServiceUnavailableError)
	qt.EnumMap["network.QNetworkReply__UnknownServerError"] = int64(QNetworkReply__UnknownServerError)
	qt.ItfMap["network.QNetworkRequest_ITF"] = QNetworkRequest{}
	qt.FuncMap["network.NewQNetworkRequest"] = NewQNetworkRequest
	qt.FuncMap["network.NewQNetworkRequest2"] = NewQNetworkRequest2
	qt.EnumMap["network.QNetworkRequest__HttpStatusCodeAttribute"] = int64(QNetworkRequest__HttpStatusCodeAttribute)
	qt.EnumMap["network.QNetworkRequest__HttpReasonPhraseAttribute"] = int64(QNetworkRequest__HttpReasonPhraseAttribute)
	qt.EnumMap["network.QNetworkRequest__RedirectionTargetAttribute"] = int64(QNetworkRequest__RedirectionTargetAttribute)
	qt.EnumMap["network.QNetworkRequest__ConnectionEncryptedAttribute"] = int64(QNetworkRequest__ConnectionEncryptedAttribute)
	qt.EnumMap["network.QNetworkRequest__CacheLoadControlAttribute"] = int64(QNetworkRequest__CacheLoadControlAttribute)
	qt.EnumMap["network.QNetworkRequest__CacheSaveControlAttribute"] = int64(QNetworkRequest__CacheSaveControlAttribute)
	qt.EnumMap["network.QNetworkRequest__SourceIsFromCacheAttribute"] = int64(QNetworkRequest__SourceIsFromCacheAttribute)
	qt.EnumMap["network.QNetworkRequest__DoNotBufferUploadDataAttribute"] = int64(QNetworkRequest__DoNotBufferUploadDataAttribute)
	qt.EnumMap["network.QNetworkRequest__HttpPipeliningAllowedAttribute"] = int64(QNetworkRequest__HttpPipeliningAllowedAttribute)
	qt.EnumMap["network.QNetworkRequest__HttpPipeliningWasUsedAttribute"] = int64(QNetworkRequest__HttpPipeliningWasUsedAttribute)
	qt.EnumMap["network.QNetworkRequest__CustomVerbAttribute"] = int64(QNetworkRequest__CustomVerbAttribute)
	qt.EnumMap["network.QNetworkRequest__CookieLoadControlAttribute"] = int64(QNetworkRequest__CookieLoadControlAttribute)
	qt.EnumMap["network.QNetworkRequest__AuthenticationReuseAttribute"] = int64(QNetworkRequest__AuthenticationReuseAttribute)
	qt.EnumMap["network.QNetworkRequest__CookieSaveControlAttribute"] = int64(QNetworkRequest__CookieSaveControlAttribute)
	qt.EnumMap["network.QNetworkRequest__MaximumDownloadBufferSizeAttribute"] = int64(QNetworkRequest__MaximumDownloadBufferSizeAttribute)
	qt.EnumMap["network.QNetworkRequest__DownloadBufferAttribute"] = int64(QNetworkRequest__DownloadBufferAttribute)
	qt.EnumMap["network.QNetworkRequest__SynchronousRequestAttribute"] = int64(QNetworkRequest__SynchronousRequestAttribute)
	qt.EnumMap["network.QNetworkRequest__BackgroundRequestAttribute"] = int64(QNetworkRequest__BackgroundRequestAttribute)
	qt.EnumMap["network.QNetworkRequest__SpdyAllowedAttribute"] = int64(QNetworkRequest__SpdyAllowedAttribute)
	qt.EnumMap["network.QNetworkRequest__SpdyWasUsedAttribute"] = int64(QNetworkRequest__SpdyWasUsedAttribute)
	qt.EnumMap["network.QNetworkRequest__EmitAllUploadProgressSignalsAttribute"] = int64(QNetworkRequest__EmitAllUploadProgressSignalsAttribute)
	qt.EnumMap["network.QNetworkRequest__FollowRedirectsAttribute"] = int64(QNetworkRequest__FollowRedirectsAttribute)
	qt.EnumMap["network.QNetworkRequest__HTTP2AllowedAttribute"] = int64(QNetworkRequest__HTTP2AllowedAttribute)
	qt.EnumMap["network.QNetworkRequest__HTTP2WasUsedAttribute"] = int64(QNetworkRequest__HTTP2WasUsedAttribute)
	qt.EnumMap["network.QNetworkRequest__OriginalContentLengthAttribute"] = int64(QNetworkRequest__OriginalContentLengthAttribute)
	qt.EnumMap["network.QNetworkRequest__RedirectPolicyAttribute"] = int64(QNetworkRequest__RedirectPolicyAttribute)
	qt.EnumMap["network.QNetworkRequest__Http2DirectAttribute"] = int64(QNetworkRequest__Http2DirectAttribute)
	qt.EnumMap["network.QNetworkRequest__ResourceTypeAttribute"] = int64(QNetworkRequest__ResourceTypeAttribute)
	qt.EnumMap["network.QNetworkRequest__User"] = int64(QNetworkRequest__User)
	qt.EnumMap["network.QNetworkRequest__UserMax"] = int64(QNetworkRequest__UserMax)
	qt.EnumMap["network.QNetworkRequest__ContentTypeHeader"] = int64(QNetworkRequest__ContentTypeHeader)
	qt.EnumMap["network.QNetworkRequest__ContentLengthHeader"] = int64(QNetworkRequest__ContentLengthHeader)
	qt.EnumMap["network.QNetworkRequest__LocationHeader"] = int64(QNetworkRequest__LocationHeader)
	qt.EnumMap["network.QNetworkRequest__LastModifiedHeader"] = int64(QNetworkRequest__LastModifiedHeader)
	qt.EnumMap["network.QNetworkRequest__CookieHeader"] = int64(QNetworkRequest__CookieHeader)
	qt.EnumMap["network.QNetworkRequest__SetCookieHeader"] = int64(QNetworkRequest__SetCookieHeader)
	qt.EnumMap["network.QNetworkRequest__ContentDispositionHeader"] = int64(QNetworkRequest__ContentDispositionHeader)
	qt.EnumMap["network.QNetworkRequest__UserAgentHeader"] = int64(QNetworkRequest__UserAgentHeader)
	qt.EnumMap["network.QNetworkRequest__ServerHeader"] = int64(QNetworkRequest__ServerHeader)
	qt.EnumMap["network.QNetworkRequest__IfModifiedSinceHeader"] = int64(QNetworkRequest__IfModifiedSinceHeader)
	qt.EnumMap["network.QNetworkRequest__ETagHeader"] = int64(QNetworkRequest__ETagHeader)
	qt.EnumMap["network.QNetworkRequest__IfMatchHeader"] = int64(QNetworkRequest__IfMatchHeader)
	qt.EnumMap["network.QNetworkRequest__IfNoneMatchHeader"] = int64(QNetworkRequest__IfNoneMatchHeader)
	qt.EnumMap["network.QNetworkRequest__AlwaysNetwork"] = int64(QNetworkRequest__AlwaysNetwork)
	qt.EnumMap["network.QNetworkRequest__PreferNetwork"] = int64(QNetworkRequest__PreferNetwork)
	qt.EnumMap["network.QNetworkRequest__PreferCache"] = int64(QNetworkRequest__PreferCache)
	qt.EnumMap["network.QNetworkRequest__AlwaysCache"] = int64(QNetworkRequest__AlwaysCache)
	qt.EnumMap["network.QNetworkRequest__Automatic"] = int64(QNetworkRequest__Automatic)
	qt.EnumMap["network.QNetworkRequest__Manual"] = int64(QNetworkRequest__Manual)
	qt.EnumMap["network.QNetworkRequest__HighPriority"] = int64(QNetworkRequest__HighPriority)
	qt.EnumMap["network.QNetworkRequest__NormalPriority"] = int64(QNetworkRequest__NormalPriority)
	qt.EnumMap["network.QNetworkRequest__LowPriority"] = int64(QNetworkRequest__LowPriority)
	qt.EnumMap["network.QNetworkRequest__ManualRedirectPolicy"] = int64(QNetworkRequest__ManualRedirectPolicy)
	qt.EnumMap["network.QNetworkRequest__NoLessSafeRedirectPolicy"] = int64(QNetworkRequest__NoLessSafeRedirectPolicy)
	qt.EnumMap["network.QNetworkRequest__SameOriginRedirectPolicy"] = int64(QNetworkRequest__SameOriginRedirectPolicy)
	qt.EnumMap["network.QNetworkRequest__UserVerifiedRedirectPolicy"] = int64(QNetworkRequest__UserVerifiedRedirectPolicy)
	qt.ItfMap["network.QNetworkSession_ITF"] = QNetworkSession{}
	qt.FuncMap["network.NewQNetworkSession"] = NewQNetworkSession
	qt.EnumMap["network.QNetworkSession__Invalid"] = int64(QNetworkSession__Invalid)
	qt.EnumMap["network.QNetworkSession__NotAvailable"] = int64(QNetworkSession__NotAvailable)
	qt.EnumMap["network.QNetworkSession__Connecting"] = int64(QNetworkSession__Connecting)
	qt.EnumMap["network.QNetworkSession__Connected"] = int64(QNetworkSession__Connected)
	qt.EnumMap["network.QNetworkSession__Closing"] = int64(QNetworkSession__Closing)
	qt.EnumMap["network.QNetworkSession__Disconnected"] = int64(QNetworkSession__Disconnected)
	qt.EnumMap["network.QNetworkSession__Roaming"] = int64(QNetworkSession__Roaming)
	qt.EnumMap["network.QNetworkSession__UnknownSessionError"] = int64(QNetworkSession__UnknownSessionError)
	qt.EnumMap["network.QNetworkSession__SessionAbortedError"] = int64(QNetworkSession__SessionAbortedError)
	qt.EnumMap["network.QNetworkSession__RoamingError"] = int64(QNetworkSession__RoamingError)
	qt.EnumMap["network.QNetworkSession__OperationNotSupportedError"] = int64(QNetworkSession__OperationNotSupportedError)
	qt.EnumMap["network.QNetworkSession__InvalidConfigurationError"] = int64(QNetworkSession__InvalidConfigurationError)
	qt.EnumMap["network.QNetworkSession__NoPolicy"] = int64(QNetworkSession__NoPolicy)
	qt.EnumMap["network.QNetworkSession__NoBackgroundTrafficPolicy"] = int64(QNetworkSession__NoBackgroundTrafficPolicy)
	qt.ItfMap["network.QOcspResponse_ITF"] = QOcspResponse{}
	qt.FuncMap["network.NewQOcspResponse"] = NewQOcspResponse
	qt.FuncMap["network.NewQOcspResponse2"] = NewQOcspResponse2
	qt.FuncMap["network.NewQOcspResponse3"] = NewQOcspResponse3
	qt.EnumMap["network.QOcspResponse__Good"] = int64(QOcspResponse__Good)
	qt.EnumMap["network.QOcspResponse__Revoked"] = int64(QOcspResponse__Revoked)
	qt.EnumMap["network.QOcspResponse__Unknown"] = int64(QOcspResponse__Unknown)
	qt.EnumMap["network.QOcspResponse__None"] = int64(QOcspResponse__None)
	qt.EnumMap["network.QOcspResponse__Unspecified"] = int64(QOcspResponse__Unspecified)
	qt.EnumMap["network.QOcspResponse__KeyCompromise"] = int64(QOcspResponse__KeyCompromise)
	qt.EnumMap["network.QOcspResponse__CACompromise"] = int64(QOcspResponse__CACompromise)
	qt.EnumMap["network.QOcspResponse__AffiliationChanged"] = int64(QOcspResponse__AffiliationChanged)
	qt.EnumMap["network.QOcspResponse__Superseded"] = int64(QOcspResponse__Superseded)
	qt.EnumMap["network.QOcspResponse__CessationOfOperation"] = int64(QOcspResponse__CessationOfOperation)
	qt.EnumMap["network.QOcspResponse__CertificateHold"] = int64(QOcspResponse__CertificateHold)
	qt.EnumMap["network.QOcspResponse__RemoveFromCRL"] = int64(QOcspResponse__RemoveFromCRL)
	qt.ItfMap["network.QSsl_ITF"] = QSsl{}
	qt.EnumMap["network.QSsl__PrivateKey"] = int64(QSsl__PrivateKey)
	qt.EnumMap["network.QSsl__PublicKey"] = int64(QSsl__PublicKey)
	qt.EnumMap["network.QSsl__Pem"] = int64(QSsl__Pem)
	qt.EnumMap["network.QSsl__Der"] = int64(QSsl__Der)
	qt.EnumMap["network.QSsl__Opaque"] = int64(QSsl__Opaque)
	qt.EnumMap["network.QSsl__Rsa"] = int64(QSsl__Rsa)
	qt.EnumMap["network.QSsl__Dsa"] = int64(QSsl__Dsa)
	qt.EnumMap["network.QSsl__Ec"] = int64(QSsl__Ec)
	qt.EnumMap["network.QSsl__Dh"] = int64(QSsl__Dh)
	qt.EnumMap["network.QSsl__EmailEntry"] = int64(QSsl__EmailEntry)
	qt.EnumMap["network.QSsl__DnsEntry"] = int64(QSsl__DnsEntry)
	qt.EnumMap["network.QSsl__IpAddressEntry"] = int64(QSsl__IpAddressEntry)
	qt.EnumMap["network.QSsl__SslV3"] = int64(QSsl__SslV3)
	qt.EnumMap["network.QSsl__SslV2"] = int64(QSsl__SslV2)
	qt.EnumMap["network.QSsl__TlsV1_0"] = int64(QSsl__TlsV1_0)
	qt.EnumMap["network.QSsl__TlsV1"] = int64(QSsl__TlsV1)
	qt.EnumMap["network.QSsl__TlsV1_1"] = int64(QSsl__TlsV1_1)
	qt.EnumMap["network.QSsl__TlsV1_2"] = int64(QSsl__TlsV1_2)
	qt.EnumMap["network.QSsl__AnyProtocol"] = int64(QSsl__AnyProtocol)
	qt.EnumMap["network.QSsl__TlsV1SslV3"] = int64(QSsl__TlsV1SslV3)
	qt.EnumMap["network.QSsl__SecureProtocols"] = int64(QSsl__SecureProtocols)
	qt.EnumMap["network.QSsl__TlsV1_0OrLater"] = int64(QSsl__TlsV1_0OrLater)
	qt.EnumMap["network.QSsl__TlsV1_1OrLater"] = int64(QSsl__TlsV1_1OrLater)
	qt.EnumMap["network.QSsl__TlsV1_2OrLater"] = int64(QSsl__TlsV1_2OrLater)
	qt.EnumMap["network.QSsl__DtlsV1_0"] = int64(QSsl__DtlsV1_0)
	qt.EnumMap["network.QSsl__DtlsV1_0OrLater"] = int64(QSsl__DtlsV1_0OrLater)
	qt.EnumMap["network.QSsl__DtlsV1_2"] = int64(QSsl__DtlsV1_2)
	qt.EnumMap["network.QSsl__DtlsV1_2OrLater"] = int64(QSsl__DtlsV1_2OrLater)
	qt.EnumMap["network.QSsl__TlsV1_3"] = int64(QSsl__TlsV1_3)
	qt.EnumMap["network.QSsl__TlsV1_3OrLater"] = int64(QSsl__TlsV1_3OrLater)
	qt.EnumMap["network.QSsl__UnknownProtocol"] = int64(QSsl__UnknownProtocol)
	qt.EnumMap["network.QSsl__SslOptionDisableEmptyFragments"] = int64(QSsl__SslOptionDisableEmptyFragments)
	qt.EnumMap["network.QSsl__SslOptionDisableSessionTickets"] = int64(QSsl__SslOptionDisableSessionTickets)
	qt.EnumMap["network.QSsl__SslOptionDisableCompression"] = int64(QSsl__SslOptionDisableCompression)
	qt.EnumMap["network.QSsl__SslOptionDisableServerNameIndication"] = int64(QSsl__SslOptionDisableServerNameIndication)
	qt.EnumMap["network.QSsl__SslOptionDisableLegacyRenegotiation"] = int64(QSsl__SslOptionDisableLegacyRenegotiation)
	qt.EnumMap["network.QSsl__SslOptionDisableSessionSharing"] = int64(QSsl__SslOptionDisableSessionSharing)
	qt.EnumMap["network.QSsl__SslOptionDisableSessionPersistence"] = int64(QSsl__SslOptionDisableSessionPersistence)
	qt.EnumMap["network.QSsl__SslOptionDisableServerCipherPreference"] = int64(QSsl__SslOptionDisableServerCipherPreference)
	qt.ItfMap["network.QSslCertificate_ITF"] = QSslCertificate{}
	qt.FuncMap["network.NewQSslCertificate"] = NewQSslCertificate
	qt.FuncMap["network.NewQSslCertificate2"] = NewQSslCertificate2
	qt.FuncMap["network.NewQSslCertificate3"] = NewQSslCertificate3
	qt.FuncMap["network.QSslCertificate_FromData"] = QSslCertificate_FromData
	qt.FuncMap["network.QSslCertificate_FromDevice"] = QSslCertificate_FromDevice
	qt.FuncMap["network.QSslCertificate_FromPath"] = QSslCertificate_FromPath
	qt.FuncMap["network.QSslCertificate_ImportPkcs12"] = QSslCertificate_ImportPkcs12
	qt.FuncMap["network.QSslCertificate_Verify"] = QSslCertificate_Verify
	qt.EnumMap["network.QSslCertificate__Organization"] = int64(QSslCertificate__Organization)
	qt.EnumMap["network.QSslCertificate__CommonName"] = int64(QSslCertificate__CommonName)
	qt.EnumMap["network.QSslCertificate__LocalityName"] = int64(QSslCertificate__LocalityName)
	qt.EnumMap["network.QSslCertificate__OrganizationalUnitName"] = int64(QSslCertificate__OrganizationalUnitName)
	qt.EnumMap["network.QSslCertificate__CountryName"] = int64(QSslCertificate__CountryName)
	qt.EnumMap["network.QSslCertificate__StateOrProvinceName"] = int64(QSslCertificate__StateOrProvinceName)
	qt.EnumMap["network.QSslCertificate__DistinguishedNameQualifier"] = int64(QSslCertificate__DistinguishedNameQualifier)
	qt.EnumMap["network.QSslCertificate__SerialNumber"] = int64(QSslCertificate__SerialNumber)
	qt.EnumMap["network.QSslCertificate__EmailAddress"] = int64(QSslCertificate__EmailAddress)
	qt.ItfMap["network.QSslCertificateExtension_ITF"] = QSslCertificateExtension{}
	qt.FuncMap["network.NewQSslCertificateExtension"] = NewQSslCertificateExtension
	qt.FuncMap["network.NewQSslCertificateExtension2"] = NewQSslCertificateExtension2
	qt.ItfMap["network.QSslCipher_ITF"] = QSslCipher{}
	qt.FuncMap["network.NewQSslCipher"] = NewQSslCipher
	qt.FuncMap["network.NewQSslCipher2"] = NewQSslCipher2
	qt.FuncMap["network.NewQSslCipher3"] = NewQSslCipher3
	qt.FuncMap["network.NewQSslCipher4"] = NewQSslCipher4
	qt.ItfMap["network.QSslConfiguration_ITF"] = QSslConfiguration{}
	qt.FuncMap["network.NewQSslConfiguration"] = NewQSslConfiguration
	qt.FuncMap["network.NewQSslConfiguration2"] = NewQSslConfiguration2
	qt.FuncMap["network.QSslConfiguration_DefaultConfiguration"] = QSslConfiguration_DefaultConfiguration
	qt.FuncMap["network.QSslConfiguration_SetDefaultConfiguration"] = QSslConfiguration_SetDefaultConfiguration
	qt.FuncMap["network.QSslConfiguration_SupportedCiphers"] = QSslConfiguration_SupportedCiphers
	qt.FuncMap["network.QSslConfiguration_SystemCaCertificates"] = QSslConfiguration_SystemCaCertificates
	qt.EnumMap["network.QSslConfiguration__NextProtocolNegotiationNone"] = int64(QSslConfiguration__NextProtocolNegotiationNone)
	qt.EnumMap["network.QSslConfiguration__NextProtocolNegotiationNegotiated"] = int64(QSslConfiguration__NextProtocolNegotiationNegotiated)
	qt.EnumMap["network.QSslConfiguration__NextProtocolNegotiationUnsupported"] = int64(QSslConfiguration__NextProtocolNegotiationUnsupported)
	qt.ItfMap["network.QSslDiffieHellmanParameters_ITF"] = QSslDiffieHellmanParameters{}
	qt.FuncMap["network.NewQSslDiffieHellmanParameters"] = NewQSslDiffieHellmanParameters
	qt.FuncMap["network.NewQSslDiffieHellmanParameters2"] = NewQSslDiffieHellmanParameters2
	qt.FuncMap["network.NewQSslDiffieHellmanParameters3"] = NewQSslDiffieHellmanParameters3
	qt.FuncMap["network.QSslDiffieHellmanParameters_DefaultParameters"] = QSslDiffieHellmanParameters_DefaultParameters
	qt.FuncMap["network.QSslDiffieHellmanParameters_FromEncoded"] = QSslDiffieHellmanParameters_FromEncoded
	qt.FuncMap["network.QSslDiffieHellmanParameters_FromEncoded2"] = QSslDiffieHellmanParameters_FromEncoded2
	qt.EnumMap["network.QSslDiffieHellmanParameters__NoError"] = int64(QSslDiffieHellmanParameters__NoError)
	qt.EnumMap["network.QSslDiffieHellmanParameters__InvalidInputDataError"] = int64(QSslDiffieHellmanParameters__InvalidInputDataError)
	qt.EnumMap["network.QSslDiffieHellmanParameters__UnsafeParametersError"] = int64(QSslDiffieHellmanParameters__UnsafeParametersError)
	qt.ItfMap["network.QSslEllipticCurve_ITF"] = QSslEllipticCurve{}
	qt.FuncMap["network.NewQSslEllipticCurve"] = NewQSslEllipticCurve
	qt.FuncMap["network.QSslEllipticCurve_FromLongName"] = QSslEllipticCurve_FromLongName
	qt.FuncMap["network.QSslEllipticCurve_FromShortName"] = QSslEllipticCurve_FromShortName
	qt.ItfMap["network.QSslError_ITF"] = QSslError{}
	qt.FuncMap["network.NewQSslError"] = NewQSslError
	qt.FuncMap["network.NewQSslError2"] = NewQSslError2
	qt.FuncMap["network.NewQSslError3"] = NewQSslError3
	qt.FuncMap["network.NewQSslError4"] = NewQSslError4
	qt.EnumMap["network.QSslError__NoError"] = int64(QSslError__NoError)
	qt.EnumMap["network.QSslError__UnableToGetIssuerCertificate"] = int64(QSslError__UnableToGetIssuerCertificate)
	qt.EnumMap["network.QSslError__UnableToDecryptCertificateSignature"] = int64(QSslError__UnableToDecryptCertificateSignature)
	qt.EnumMap["network.QSslError__UnableToDecodeIssuerPublicKey"] = int64(QSslError__UnableToDecodeIssuerPublicKey)
	qt.EnumMap["network.QSslError__CertificateSignatureFailed"] = int64(QSslError__CertificateSignatureFailed)
	qt.EnumMap["network.QSslError__CertificateNotYetValid"] = int64(QSslError__CertificateNotYetValid)
	qt.EnumMap["network.QSslError__CertificateExpired"] = int64(QSslError__CertificateExpired)
	qt.EnumMap["network.QSslError__InvalidNotBeforeField"] = int64(QSslError__InvalidNotBeforeField)
	qt.EnumMap["network.QSslError__InvalidNotAfterField"] = int64(QSslError__InvalidNotAfterField)
	qt.EnumMap["network.QSslError__SelfSignedCertificate"] = int64(QSslError__SelfSignedCertificate)
	qt.EnumMap["network.QSslError__SelfSignedCertificateInChain"] = int64(QSslError__SelfSignedCertificateInChain)
	qt.EnumMap["network.QSslError__UnableToGetLocalIssuerCertificate"] = int64(QSslError__UnableToGetLocalIssuerCertificate)
	qt.EnumMap["network.QSslError__UnableToVerifyFirstCertificate"] = int64(QSslError__UnableToVerifyFirstCertificate)
	qt.EnumMap["network.QSslError__CertificateRevoked"] = int64(QSslError__CertificateRevoked)
	qt.EnumMap["network.QSslError__InvalidCaCertificate"] = int64(QSslError__InvalidCaCertificate)
	qt.EnumMap["network.QSslError__PathLengthExceeded"] = int64(QSslError__PathLengthExceeded)
	qt.EnumMap["network.QSslError__InvalidPurpose"] = int64(QSslError__InvalidPurpose)
	qt.EnumMap["network.QSslError__CertificateUntrusted"] = int64(QSslError__CertificateUntrusted)
	qt.EnumMap["network.QSslError__CertificateRejected"] = int64(QSslError__CertificateRejected)
	qt.EnumMap["network.QSslError__SubjectIssuerMismatch"] = int64(QSslError__SubjectIssuerMismatch)
	qt.EnumMap["network.QSslError__AuthorityIssuerSerialNumberMismatch"] = int64(QSslError__AuthorityIssuerSerialNumberMismatch)
	qt.EnumMap["network.QSslError__NoPeerCertificate"] = int64(QSslError__NoPeerCertificate)
	qt.EnumMap["network.QSslError__HostNameMismatch"] = int64(QSslError__HostNameMismatch)
	qt.EnumMap["network.QSslError__NoSslSupport"] = int64(QSslError__NoSslSupport)
	qt.EnumMap["network.QSslError__CertificateBlacklisted"] = int64(QSslError__CertificateBlacklisted)
	qt.EnumMap["network.QSslError__CertificateStatusUnknown"] = int64(QSslError__CertificateStatusUnknown)
	qt.EnumMap["network.QSslError__OcspNoResponseFound"] = int64(QSslError__OcspNoResponseFound)
	qt.EnumMap["network.QSslError__OcspMalformedRequest"] = int64(QSslError__OcspMalformedRequest)
	qt.EnumMap["network.QSslError__OcspMalformedResponse"] = int64(QSslError__OcspMalformedResponse)
	qt.EnumMap["network.QSslError__OcspInternalError"] = int64(QSslError__OcspInternalError)
	qt.EnumMap["network.QSslError__OcspTryLater"] = int64(QSslError__OcspTryLater)
	qt.EnumMap["network.QSslError__OcspSigRequred"] = int64(QSslError__OcspSigRequred)
	qt.EnumMap["network.QSslError__OcspUnauthorized"] = int64(QSslError__OcspUnauthorized)
	qt.EnumMap["network.QSslError__OcspResponseCannotBeTrusted"] = int64(QSslError__OcspResponseCannotBeTrusted)
	qt.EnumMap["network.QSslError__OcspResponseCertIdUnknown"] = int64(QSslError__OcspResponseCertIdUnknown)
	qt.EnumMap["network.QSslError__OcspResponseExpired"] = int64(QSslError__OcspResponseExpired)
	qt.EnumMap["network.QSslError__OcspStatusUnknown"] = int64(QSslError__OcspStatusUnknown)
	qt.EnumMap["network.QSslError__UnspecifiedError"] = int64(QSslError__UnspecifiedError)
	qt.ItfMap["network.QSslKey_ITF"] = QSslKey{}
	qt.FuncMap["network.NewQSslKey"] = NewQSslKey
	qt.FuncMap["network.NewQSslKey2"] = NewQSslKey2
	qt.FuncMap["network.NewQSslKey3"] = NewQSslKey3
	qt.FuncMap["network.NewQSslKey5"] = NewQSslKey5
	qt.ItfMap["network.QSslPreSharedKeyAuthenticator_ITF"] = QSslPreSharedKeyAuthenticator{}
	qt.FuncMap["network.NewQSslPreSharedKeyAuthenticator"] = NewQSslPreSharedKeyAuthenticator
	qt.FuncMap["network.NewQSslPreSharedKeyAuthenticator2"] = NewQSslPreSharedKeyAuthenticator2
	qt.ItfMap["network.QSslSocket_ITF"] = QSslSocket{}
	qt.FuncMap["network.NewQSslSocket"] = NewQSslSocket
	qt.FuncMap["network.QSslSocket_AddDefaultCaCertificate"] = QSslSocket_AddDefaultCaCertificate
	qt.FuncMap["network.QSslSocket_AddDefaultCaCertificates"] = QSslSocket_AddDefaultCaCertificates
	qt.FuncMap["network.QSslSocket_AddDefaultCaCertificates2"] = QSslSocket_AddDefaultCaCertificates2
	qt.FuncMap["network.QSslSocket_SslLibraryBuildVersionNumber"] = QSslSocket_SslLibraryBuildVersionNumber
	qt.FuncMap["network.QSslSocket_SslLibraryBuildVersionString"] = QSslSocket_SslLibraryBuildVersionString
	qt.FuncMap["network.QSslSocket_SslLibraryVersionNumber"] = QSslSocket_SslLibraryVersionNumber
	qt.FuncMap["network.QSslSocket_SslLibraryVersionString"] = QSslSocket_SslLibraryVersionString
	qt.FuncMap["network.QSslSocket_SupportsSsl"] = QSslSocket_SupportsSsl
	qt.EnumMap["network.QSslSocket__UnencryptedMode"] = int64(QSslSocket__UnencryptedMode)
	qt.EnumMap["network.QSslSocket__SslClientMode"] = int64(QSslSocket__SslClientMode)
	qt.EnumMap["network.QSslSocket__SslServerMode"] = int64(QSslSocket__SslServerMode)
	qt.EnumMap["network.QSslSocket__VerifyNone"] = int64(QSslSocket__VerifyNone)
	qt.EnumMap["network.QSslSocket__QueryPeer"] = int64(QSslSocket__QueryPeer)
	qt.EnumMap["network.QSslSocket__VerifyPeer"] = int64(QSslSocket__VerifyPeer)
	qt.EnumMap["network.QSslSocket__AutoVerifyPeer"] = int64(QSslSocket__AutoVerifyPeer)
	qt.ItfMap["network.QTcpServer_ITF"] = QTcpServer{}
	qt.FuncMap["network.NewQTcpServer"] = NewQTcpServer
	qt.ItfMap["network.QTcpSocket_ITF"] = QTcpSocket{}
	qt.FuncMap["network.NewQTcpSocket"] = NewQTcpSocket
	qt.ItfMap["network.QUdpSocket_ITF"] = QUdpSocket{}
	qt.FuncMap["network.NewQUdpSocket"] = NewQUdpSocket
	qt.ItfMap["network.QNetworkCacheMetaData_RawHeader_ITF"] = QNetworkCacheMetaData_RawHeader{}
	qt.FuncMap["network.NewQNetworkCacheMetaData_RawHeader"] = NewQNetworkCacheMetaData_RawHeader
	qt.FuncMap["network.NewQNetworkCacheMetaData_RawHeader2"] = NewQNetworkCacheMetaData_RawHeader2
}
