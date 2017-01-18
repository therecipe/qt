// +build !minimal

package webengine

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "webengine.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"github.com/therecipe/qt/webchannel"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtWebEngine_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QQuickWebEngineProfile struct {
	core.QObject
}

type QQuickWebEngineProfile_ITF interface {
	core.QObject_ITF
	QQuickWebEngineProfile_PTR() *QQuickWebEngineProfile
}

func (ptr *QQuickWebEngineProfile) QQuickWebEngineProfile_PTR() *QQuickWebEngineProfile {
	return ptr
}

func (ptr *QQuickWebEngineProfile) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickWebEngineProfile(ptr QQuickWebEngineProfile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWebEngineProfile_PTR().Pointer()
	}
	return nil
}

func NewQQuickWebEngineProfileFromPointer(ptr unsafe.Pointer) *QQuickWebEngineProfile {
	var n = new(QQuickWebEngineProfile)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuickWebEngineProfile) DestroyQQuickWebEngineProfile() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QQuickWebEngineProfile__HttpCacheType
//QQuickWebEngineProfile::HttpCacheType
type QQuickWebEngineProfile__HttpCacheType int64

const (
	QQuickWebEngineProfile__MemoryHttpCache QQuickWebEngineProfile__HttpCacheType = QQuickWebEngineProfile__HttpCacheType(0)
	QQuickWebEngineProfile__DiskHttpCache   QQuickWebEngineProfile__HttpCacheType = QQuickWebEngineProfile__HttpCacheType(1)
	QQuickWebEngineProfile__NoCache         QQuickWebEngineProfile__HttpCacheType = QQuickWebEngineProfile__HttpCacheType(2)
)

//go:generate stringer -type=QQuickWebEngineProfile__PersistentCookiesPolicy
//QQuickWebEngineProfile::PersistentCookiesPolicy
type QQuickWebEngineProfile__PersistentCookiesPolicy int64

const (
	QQuickWebEngineProfile__NoPersistentCookies    QQuickWebEngineProfile__PersistentCookiesPolicy = QQuickWebEngineProfile__PersistentCookiesPolicy(0)
	QQuickWebEngineProfile__AllowPersistentCookies QQuickWebEngineProfile__PersistentCookiesPolicy = QQuickWebEngineProfile__PersistentCookiesPolicy(1)
	QQuickWebEngineProfile__ForcePersistentCookies QQuickWebEngineProfile__PersistentCookiesPolicy = QQuickWebEngineProfile__PersistentCookiesPolicy(2)
)

func (ptr *QQuickWebEngineProfile) CachePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_CachePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_HttpAcceptLanguage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWebEngineProfile_HttpCacheMaximumSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickWebEngineProfile) HttpCacheType() QQuickWebEngineProfile__HttpCacheType {
	if ptr.Pointer() != nil {
		return QQuickWebEngineProfile__HttpCacheType(C.QQuickWebEngineProfile_HttpCacheType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWebEngineProfile) HttpUserAgent() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_HttpUserAgent(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) IsOffTheRecord() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_IsOffTheRecord(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicy() QQuickWebEngineProfile__PersistentCookiesPolicy {
	if ptr.Pointer() != nil {
		return QQuickWebEngineProfile__PersistentCookiesPolicy(C.QQuickWebEngineProfile_PersistentCookiesPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_PersistentStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) SetCachePath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QQuickWebEngineProfile_SetCachePath(ptr.Pointer(), pathC)
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	if ptr.Pointer() != nil {
		var httpAcceptLanguageC = C.CString(httpAcceptLanguage)
		defer C.free(unsafe.Pointer(httpAcceptLanguageC))
		C.QQuickWebEngineProfile_SetHttpAcceptLanguage(ptr.Pointer(), httpAcceptLanguageC)
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetHttpCacheMaximumSize(ptr.Pointer(), C.int(int32(maxSize)))
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpCacheType(vqq QQuickWebEngineProfile__HttpCacheType) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetHttpCacheType(ptr.Pointer(), C.longlong(vqq))
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpUserAgent(userAgent string) {
	if ptr.Pointer() != nil {
		var userAgentC = C.CString(userAgent)
		defer C.free(unsafe.Pointer(userAgentC))
		C.QQuickWebEngineProfile_SetHttpUserAgent(ptr.Pointer(), userAgentC)
	}
}

func (ptr *QQuickWebEngineProfile) SetOffTheRecord(offTheRecord bool) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetOffTheRecord(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(offTheRecord))))
	}
}

func (ptr *QQuickWebEngineProfile) SetPersistentCookiesPolicy(vqq QQuickWebEngineProfile__PersistentCookiesPolicy) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetPersistentCookiesPolicy(ptr.Pointer(), C.longlong(vqq))
	}
}

func (ptr *QQuickWebEngineProfile) SetPersistentStoragePath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QQuickWebEngineProfile_SetPersistentStoragePath(ptr.Pointer(), pathC)
	}
}

func (ptr *QQuickWebEngineProfile) SetStorageName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QQuickWebEngineProfile_SetStorageName(ptr.Pointer(), nameC)
	}
}

func (ptr *QQuickWebEngineProfile) StorageName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_StorageName(ptr.Pointer()))
	}
	return ""
}

func NewQQuickWebEngineProfile(parent core.QObject_ITF) *QQuickWebEngineProfile {
	var tmpValue = NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_NewQQuickWebEngineProfile(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickWebEngineProfile_CachePathChanged
func callbackQQuickWebEngineProfile_CachePathChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::cachePathChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectCachePathChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectCachePathChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::cachePathChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectCachePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectCachePathChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::cachePathChanged")
	}
}

func (ptr *QQuickWebEngineProfile) CachePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CachePathChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) ClearHttpCache() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ClearHttpCache(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) CookieStore() *QWebEngineCookieStore {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineCookieStoreFromPointer(C.QQuickWebEngineProfile_CookieStore(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QQuickWebEngineProfile_DefaultProfile() *QQuickWebEngineProfile {
	var tmpValue = NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickWebEngineProfile) DefaultProfile() *QQuickWebEngineProfile {
	var tmpValue = NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged
func callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::httpAcceptLanguageChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpAcceptLanguageChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpAcceptLanguageChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpAcceptLanguageChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpAcceptLanguageChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpAcceptLanguageChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpAcceptLanguageChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguageChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpAcceptLanguageChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged
func callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::httpCacheMaximumSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheMaximumSizeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpCacheMaximumSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpCacheMaximumSizeChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheMaximumSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpCacheMaximumSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpCacheMaximumSizeChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpCacheMaximumSizeChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_HttpCacheTypeChanged
func callbackQQuickWebEngineProfile_HttpCacheTypeChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::httpCacheTypeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheTypeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpCacheTypeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpCacheTypeChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpCacheTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpCacheTypeChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpCacheTypeChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_HttpUserAgentChanged
func callbackQQuickWebEngineProfile_HttpUserAgentChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::httpUserAgentChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpUserAgentChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpUserAgentChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpUserAgentChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpUserAgentChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpUserAgentChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::httpUserAgentChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpUserAgentChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpUserAgentChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) InstallUrlSchemeHandler(scheme core.QByteArray_ITF, handler QWebEngineUrlSchemeHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_InstallUrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(scheme), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

//export callbackQQuickWebEngineProfile_OffTheRecordChanged
func callbackQQuickWebEngineProfile_OffTheRecordChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::offTheRecordChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectOffTheRecordChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectOffTheRecordChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::offTheRecordChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectOffTheRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectOffTheRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::offTheRecordChanged")
	}
}

func (ptr *QQuickWebEngineProfile) OffTheRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_OffTheRecordChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged
func callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::persistentCookiesPolicyChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPersistentCookiesPolicyChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectPersistentCookiesPolicyChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::persistentCookiesPolicyChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentCookiesPolicyChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPersistentCookiesPolicyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::persistentCookiesPolicyChanged")
	}
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicyChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PersistentCookiesPolicyChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_PersistentStoragePathChanged
func callbackQQuickWebEngineProfile_PersistentStoragePathChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::persistentStoragePathChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPersistentStoragePathChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectPersistentStoragePathChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::persistentStoragePathChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentStoragePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPersistentStoragePathChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::persistentStoragePathChanged")
	}
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PersistentStoragePathChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) RemoveAllUrlSchemeHandlers() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_RemoveAllUrlSchemeHandlers(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) RemoveUrlScheme(scheme core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_RemoveUrlScheme(ptr.Pointer(), core.PointerFromQByteArray(scheme))
	}
}

func (ptr *QQuickWebEngineProfile) RemoveUrlSchemeHandler(handler QWebEngineUrlSchemeHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_RemoveUrlSchemeHandler(ptr.Pointer(), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

func (ptr *QQuickWebEngineProfile) SetRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

//export callbackQQuickWebEngineProfile_StorageNameChanged
func callbackQQuickWebEngineProfile_StorageNameChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::storageNameChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectStorageNameChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectStorageNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::storageNameChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectStorageNameChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectStorageNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::storageNameChanged")
	}
}

func (ptr *QQuickWebEngineProfile) StorageNameChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_StorageNameChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) UrlSchemeHandler(scheme core.QByteArray_ITF) *QWebEngineUrlSchemeHandler {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineUrlSchemeHandlerFromPointer(C.QQuickWebEngineProfile_UrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(scheme)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWebEngineProfile_TimerEvent
func callbackQQuickWebEngineProfile_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::timerEvent", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::timerEvent")
	}
}

func (ptr *QQuickWebEngineProfile) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_ChildEvent
func callbackQQuickWebEngineProfile_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::childEvent", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::childEvent")
	}
}

func (ptr *QQuickWebEngineProfile) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_ConnectNotify
func callbackQQuickWebEngineProfile_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::connectNotify", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::connectNotify")
	}
}

func (ptr *QQuickWebEngineProfile) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineProfile_CustomEvent
func callbackQQuickWebEngineProfile_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::customEvent", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::customEvent")
	}
}

func (ptr *QQuickWebEngineProfile) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_DeleteLater
func callbackQQuickWebEngineProfile_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWebEngineProfile) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::deleteLater", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::deleteLater")
	}
}

func (ptr *QQuickWebEngineProfile) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWebEngineProfile) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWebEngineProfile_DisconnectNotify
func callbackQQuickWebEngineProfile_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::disconnectNotify", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::disconnectNotify")
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineProfile_Event
func callbackQQuickWebEngineProfile_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWebEngineProfileFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWebEngineProfile) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::event", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::event")
	}
}

func (ptr *QQuickWebEngineProfile) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_EventFilter
func callbackQQuickWebEngineProfile_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWebEngineProfileFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWebEngineProfile) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::eventFilter", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::eventFilter")
	}
}

func (ptr *QQuickWebEngineProfile) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_MetaObject
func callbackQQuickWebEngineProfile_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWebEngineProfile::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWebEngineProfileFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWebEngineProfile) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::metaObject", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWebEngineProfile::metaObject")
	}
}

func (ptr *QQuickWebEngineProfile) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWebEngineProfile_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWebEngineProfile_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineCertificateError struct {
	ptr unsafe.Pointer
}

type QWebEngineCertificateError_ITF interface {
	QWebEngineCertificateError_PTR() *QWebEngineCertificateError
}

func (ptr *QWebEngineCertificateError) QWebEngineCertificateError_PTR() *QWebEngineCertificateError {
	return ptr
}

func (ptr *QWebEngineCertificateError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineCertificateError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineCertificateError(ptr QWebEngineCertificateError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineCertificateError_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineCertificateErrorFromPointer(ptr unsafe.Pointer) *QWebEngineCertificateError {
	var n = new(QWebEngineCertificateError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineCertificateError) DestroyQWebEngineCertificateError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebEngineCertificateError__Error
//QWebEngineCertificateError::Error
type QWebEngineCertificateError__Error int64

const (
	QWebEngineCertificateError__SslPinnedKeyNotInCertificateChain  QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-150)
	QWebEngineCertificateError__CertificateCommonNameInvalid       QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-200)
	QWebEngineCertificateError__CertificateDateInvalid             QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-201)
	QWebEngineCertificateError__CertificateAuthorityInvalid        QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-202)
	QWebEngineCertificateError__CertificateContainsErrors          QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-203)
	QWebEngineCertificateError__CertificateNoRevocationMechanism   QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-204)
	QWebEngineCertificateError__CertificateUnableToCheckRevocation QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-205)
	QWebEngineCertificateError__CertificateRevoked                 QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-206)
	QWebEngineCertificateError__CertificateInvalid                 QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-207)
	QWebEngineCertificateError__CertificateWeakSignatureAlgorithm  QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-208)
	QWebEngineCertificateError__CertificateNonUniqueName           QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-210)
	QWebEngineCertificateError__CertificateWeakKey                 QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-211)
	QWebEngineCertificateError__CertificateNameConstraintViolation QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-212)
)

func (ptr *QWebEngineCertificateError) Error() QWebEngineCertificateError__Error {
	if ptr.Pointer() != nil {
		return QWebEngineCertificateError__Error(C.QWebEngineCertificateError_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineCertificateError) ErrorDescription() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineCertificateError_ErrorDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineCertificateError) IsOverridable() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineCertificateError_IsOverridable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineCertificateError) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineCertificateError_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

type QWebEngineContextMenuData struct {
	ptr unsafe.Pointer
}

type QWebEngineContextMenuData_ITF interface {
	QWebEngineContextMenuData_PTR() *QWebEngineContextMenuData
}

func (ptr *QWebEngineContextMenuData) QWebEngineContextMenuData_PTR() *QWebEngineContextMenuData {
	return ptr
}

func (ptr *QWebEngineContextMenuData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineContextMenuData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineContextMenuData(ptr QWebEngineContextMenuData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineContextMenuData_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineContextMenuDataFromPointer(ptr unsafe.Pointer) *QWebEngineContextMenuData {
	var n = new(QWebEngineContextMenuData)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebEngineContextMenuData__MediaType
//QWebEngineContextMenuData::MediaType
type QWebEngineContextMenuData__MediaType int64

const (
	QWebEngineContextMenuData__MediaTypeNone   QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(0)
	QWebEngineContextMenuData__MediaTypeImage  QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(1)
	QWebEngineContextMenuData__MediaTypeVideo  QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(2)
	QWebEngineContextMenuData__MediaTypeAudio  QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(3)
	QWebEngineContextMenuData__MediaTypeCanvas QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(4)
	QWebEngineContextMenuData__MediaTypeFile   QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(5)
	QWebEngineContextMenuData__MediaTypePlugin QWebEngineContextMenuData__MediaType = QWebEngineContextMenuData__MediaType(6)
)

func NewQWebEngineContextMenuData() *QWebEngineContextMenuData {
	var tmpValue = NewQWebEngineContextMenuDataFromPointer(C.QWebEngineContextMenuData_NewQWebEngineContextMenuData())
	runtime.SetFinalizer(tmpValue, (*QWebEngineContextMenuData).DestroyQWebEngineContextMenuData)
	return tmpValue
}

func NewQWebEngineContextMenuData2(other QWebEngineContextMenuData_ITF) *QWebEngineContextMenuData {
	var tmpValue = NewQWebEngineContextMenuDataFromPointer(C.QWebEngineContextMenuData_NewQWebEngineContextMenuData2(PointerFromQWebEngineContextMenuData(other)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineContextMenuData).DestroyQWebEngineContextMenuData)
	return tmpValue
}

func (ptr *QWebEngineContextMenuData) IsContentEditable() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineContextMenuData_IsContentEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineContextMenuData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineContextMenuData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineContextMenuData) LinkText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineContextMenuData_LinkText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineContextMenuData) LinkUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineContextMenuData_LinkUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineContextMenuData) MediaType() QWebEngineContextMenuData__MediaType {
	if ptr.Pointer() != nil {
		return QWebEngineContextMenuData__MediaType(C.QWebEngineContextMenuData_MediaType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineContextMenuData) MediaUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineContextMenuData_MediaUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineContextMenuData) Position() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QWebEngineContextMenuData_Position(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineContextMenuData) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineContextMenuData_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineContextMenuData) DestroyQWebEngineContextMenuData() {
	if ptr.Pointer() != nil {
		C.QWebEngineContextMenuData_DestroyQWebEngineContextMenuData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWebEngineCookieStore struct {
	core.QObject
}

type QWebEngineCookieStore_ITF interface {
	core.QObject_ITF
	QWebEngineCookieStore_PTR() *QWebEngineCookieStore
}

func (ptr *QWebEngineCookieStore) QWebEngineCookieStore_PTR() *QWebEngineCookieStore {
	return ptr
}

func (ptr *QWebEngineCookieStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineCookieStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineCookieStore(ptr QWebEngineCookieStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineCookieStore_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineCookieStoreFromPointer(ptr unsafe.Pointer) *QWebEngineCookieStore {
	var n = new(QWebEngineCookieStore)
	n.SetPointer(ptr)
	return n
}

//export callbackQWebEngineCookieStore_CookieAdded
func callbackQWebEngineCookieStore_CookieAdded(ptr unsafe.Pointer, cookie unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::cookieAdded"); signal != nil {
		signal.(func(*network.QNetworkCookie))(network.NewQNetworkCookieFromPointer(cookie))
	}

}

func (ptr *QWebEngineCookieStore) ConnectCookieAdded(f func(cookie *network.QNetworkCookie)) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectCookieAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::cookieAdded", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCookieAdded() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectCookieAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::cookieAdded")
	}
}

func (ptr *QWebEngineCookieStore) CookieAdded(cookie network.QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CookieAdded(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie))
	}
}

//export callbackQWebEngineCookieStore_CookieRemoved
func callbackQWebEngineCookieStore_CookieRemoved(ptr unsafe.Pointer, cookie unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::cookieRemoved"); signal != nil {
		signal.(func(*network.QNetworkCookie))(network.NewQNetworkCookieFromPointer(cookie))
	}

}

func (ptr *QWebEngineCookieStore) ConnectCookieRemoved(f func(cookie *network.QNetworkCookie)) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectCookieRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::cookieRemoved", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCookieRemoved() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectCookieRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::cookieRemoved")
	}
}

func (ptr *QWebEngineCookieStore) CookieRemoved(cookie network.QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CookieRemoved(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie))
	}
}

func (ptr *QWebEngineCookieStore) DeleteAllCookies() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteAllCookies(ptr.Pointer())
	}
}

func (ptr *QWebEngineCookieStore) DeleteCookie(cookie network.QNetworkCookie_ITF, origin core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteCookie(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(origin))
	}
}

func (ptr *QWebEngineCookieStore) DeleteSessionCookies() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteSessionCookies(ptr.Pointer())
	}
}

func (ptr *QWebEngineCookieStore) LoadAllCookies() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_LoadAllCookies(ptr.Pointer())
	}
}

func (ptr *QWebEngineCookieStore) SetCookie(cookie network.QNetworkCookie_ITF, origin core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_SetCookie(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(origin))
	}
}

//export callbackQWebEngineCookieStore_DestroyQWebEngineCookieStore
func callbackQWebEngineCookieStore_DestroyQWebEngineCookieStore(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::~QWebEngineCookieStore"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DestroyQWebEngineCookieStoreDefault()
	}
}

func (ptr *QWebEngineCookieStore) ConnectDestroyQWebEngineCookieStore(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::~QWebEngineCookieStore", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectDestroyQWebEngineCookieStore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::~QWebEngineCookieStore")
	}
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStore() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStore(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStoreDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStoreDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineCookieStore_TimerEvent
func callbackQWebEngineCookieStore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::timerEvent", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::timerEvent")
	}
}

func (ptr *QWebEngineCookieStore) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineCookieStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineCookieStore_ChildEvent
func callbackQWebEngineCookieStore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::childEvent", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::childEvent")
	}
}

func (ptr *QWebEngineCookieStore) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineCookieStore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineCookieStore_ConnectNotify
func callbackQWebEngineCookieStore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::connectNotify", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::connectNotify")
	}
}

func (ptr *QWebEngineCookieStore) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineCookieStore_CustomEvent
func callbackQWebEngineCookieStore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::customEvent", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::customEvent")
	}
}

func (ptr *QWebEngineCookieStore) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineCookieStore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineCookieStore_DeleteLater
func callbackQWebEngineCookieStore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineCookieStore) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::deleteLater", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::deleteLater")
	}
}

func (ptr *QWebEngineCookieStore) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineCookieStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineCookieStore_DisconnectNotify
func callbackQWebEngineCookieStore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::disconnectNotify", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::disconnectNotify")
	}
}

func (ptr *QWebEngineCookieStore) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineCookieStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineCookieStore_Event
func callbackQWebEngineCookieStore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineCookieStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineCookieStore) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::event", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::event")
	}
}

func (ptr *QWebEngineCookieStore) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineCookieStore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineCookieStore_EventFilter
func callbackQWebEngineCookieStore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineCookieStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineCookieStore) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::eventFilter", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::eventFilter")
	}
}

func (ptr *QWebEngineCookieStore) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineCookieStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineCookieStore_MetaObject
func callbackQWebEngineCookieStore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineCookieStore::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineCookieStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineCookieStore) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::metaObject", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineCookieStore::metaObject")
	}
}

func (ptr *QWebEngineCookieStore) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineCookieStore_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineCookieStore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineCookieStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineDownloadItem struct {
	core.QObject
}

type QWebEngineDownloadItem_ITF interface {
	core.QObject_ITF
	QWebEngineDownloadItem_PTR() *QWebEngineDownloadItem
}

func (ptr *QWebEngineDownloadItem) QWebEngineDownloadItem_PTR() *QWebEngineDownloadItem {
	return ptr
}

func (ptr *QWebEngineDownloadItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineDownloadItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineDownloadItem(ptr QWebEngineDownloadItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineDownloadItem_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineDownloadItemFromPointer(ptr unsafe.Pointer) *QWebEngineDownloadItem {
	var n = new(QWebEngineDownloadItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineDownloadItem) DestroyQWebEngineDownloadItem() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebEngineDownloadItem__DownloadState
//QWebEngineDownloadItem::DownloadState
type QWebEngineDownloadItem__DownloadState int64

const (
	QWebEngineDownloadItem__DownloadRequested   QWebEngineDownloadItem__DownloadState = QWebEngineDownloadItem__DownloadState(0)
	QWebEngineDownloadItem__DownloadInProgress  QWebEngineDownloadItem__DownloadState = QWebEngineDownloadItem__DownloadState(1)
	QWebEngineDownloadItem__DownloadCompleted   QWebEngineDownloadItem__DownloadState = QWebEngineDownloadItem__DownloadState(2)
	QWebEngineDownloadItem__DownloadCancelled   QWebEngineDownloadItem__DownloadState = QWebEngineDownloadItem__DownloadState(3)
	QWebEngineDownloadItem__DownloadInterrupted QWebEngineDownloadItem__DownloadState = QWebEngineDownloadItem__DownloadState(4)
)

//go:generate stringer -type=QWebEngineDownloadItem__SavePageFormat
//QWebEngineDownloadItem::SavePageFormat
type QWebEngineDownloadItem__SavePageFormat int64

const (
	QWebEngineDownloadItem__UnknownSaveFormat      QWebEngineDownloadItem__SavePageFormat = QWebEngineDownloadItem__SavePageFormat(-1)
	QWebEngineDownloadItem__SingleHtmlSaveFormat   QWebEngineDownloadItem__SavePageFormat = QWebEngineDownloadItem__SavePageFormat(0)
	QWebEngineDownloadItem__CompleteHtmlSaveFormat QWebEngineDownloadItem__SavePageFormat = QWebEngineDownloadItem__SavePageFormat(1)
	QWebEngineDownloadItem__MimeHtmlSaveFormat     QWebEngineDownloadItem__SavePageFormat = QWebEngineDownloadItem__SavePageFormat(2)
)

//export callbackQWebEngineDownloadItem_Accept
func callbackQWebEngineDownloadItem_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::accept"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QWebEngineDownloadItem) ConnectAccept(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::accept", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectAccept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::accept")
	}
}

func (ptr *QWebEngineDownloadItem) Accept() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_Accept(ptr.Pointer())
	}
}

func (ptr *QWebEngineDownloadItem) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineDownloadItem_Cancel
func callbackQWebEngineDownloadItem_Cancel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::cancel"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).CancelDefault()
	}
}

func (ptr *QWebEngineDownloadItem) ConnectCancel(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::cancel", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectCancel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::cancel")
	}
}

func (ptr *QWebEngineDownloadItem) Cancel() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_Cancel(ptr.Pointer())
	}
}

func (ptr *QWebEngineDownloadItem) CancelDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_CancelDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineDownloadItem_DownloadProgress
func callbackQWebEngineDownloadItem_DownloadProgress(ptr unsafe.Pointer, bytesReceived C.longlong, bytesTotal C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::downloadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesReceived), int64(bytesTotal))
	}

}

func (ptr *QWebEngineDownloadItem) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectDownloadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::downloadProgress", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectDownloadProgress() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectDownloadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::downloadProgress")
	}
}

func (ptr *QWebEngineDownloadItem) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DownloadProgress(ptr.Pointer(), C.longlong(bytesReceived), C.longlong(bytesTotal))
	}
}

//export callbackQWebEngineDownloadItem_Finished
func callbackQWebEngineDownloadItem_Finished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineDownloadItem) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::finished", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::finished")
	}
}

func (ptr *QWebEngineDownloadItem) Finished() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_Finished(ptr.Pointer())
	}
}

func (ptr *QWebEngineDownloadItem) Id() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QWebEngineDownloadItem_Id(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineDownloadItem) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineDownloadItem) MimeType() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineDownloadItem_MimeType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineDownloadItem) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineDownloadItem_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineDownloadItem) ReceivedBytes() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebEngineDownloadItem_ReceivedBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineDownloadItem) SavePageFormat() QWebEngineDownloadItem__SavePageFormat {
	if ptr.Pointer() != nil {
		return QWebEngineDownloadItem__SavePageFormat(C.QWebEngineDownloadItem_SavePageFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineDownloadItem) SetPath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QWebEngineDownloadItem_SetPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QWebEngineDownloadItem) SetSavePageFormat(format QWebEngineDownloadItem__SavePageFormat) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_SetSavePageFormat(ptr.Pointer(), C.longlong(format))
	}
}

func (ptr *QWebEngineDownloadItem) State() QWebEngineDownloadItem__DownloadState {
	if ptr.Pointer() != nil {
		return QWebEngineDownloadItem__DownloadState(C.QWebEngineDownloadItem_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebEngineDownloadItem_StateChanged
func callbackQWebEngineDownloadItem_StateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::stateChanged"); signal != nil {
		signal.(func(QWebEngineDownloadItem__DownloadState))(QWebEngineDownloadItem__DownloadState(state))
	}

}

func (ptr *QWebEngineDownloadItem) ConnectStateChanged(f func(state QWebEngineDownloadItem__DownloadState)) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::stateChanged", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::stateChanged")
	}
}

func (ptr *QWebEngineDownloadItem) StateChanged(state QWebEngineDownloadItem__DownloadState) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func (ptr *QWebEngineDownloadItem) TotalBytes() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebEngineDownloadItem_TotalBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineDownloadItem) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineDownloadItem_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineDownloadItem_TimerEvent
func callbackQWebEngineDownloadItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::timerEvent", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::timerEvent")
	}
}

func (ptr *QWebEngineDownloadItem) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineDownloadItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineDownloadItem_ChildEvent
func callbackQWebEngineDownloadItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::childEvent", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::childEvent")
	}
}

func (ptr *QWebEngineDownloadItem) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineDownloadItem) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineDownloadItem_ConnectNotify
func callbackQWebEngineDownloadItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::connectNotify", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::connectNotify")
	}
}

func (ptr *QWebEngineDownloadItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineDownloadItem_CustomEvent
func callbackQWebEngineDownloadItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::customEvent", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::customEvent")
	}
}

func (ptr *QWebEngineDownloadItem) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineDownloadItem) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineDownloadItem_DeleteLater
func callbackQWebEngineDownloadItem_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineDownloadItem) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::deleteLater", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::deleteLater")
	}
}

func (ptr *QWebEngineDownloadItem) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineDownloadItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineDownloadItem_DisconnectNotify
func callbackQWebEngineDownloadItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::disconnectNotify", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::disconnectNotify")
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineDownloadItem_Event
func callbackQWebEngineDownloadItem_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineDownloadItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineDownloadItem) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::event", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::event")
	}
}

func (ptr *QWebEngineDownloadItem) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineDownloadItem) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineDownloadItem_EventFilter
func callbackQWebEngineDownloadItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineDownloadItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineDownloadItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::eventFilter", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::eventFilter")
	}
}

func (ptr *QWebEngineDownloadItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineDownloadItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineDownloadItem_MetaObject
func callbackQWebEngineDownloadItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineDownloadItem::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineDownloadItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineDownloadItem) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::metaObject", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineDownloadItem::metaObject")
	}
}

func (ptr *QWebEngineDownloadItem) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineDownloadItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineDownloadItem) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineDownloadItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineFullScreenRequest struct {
	ptr unsafe.Pointer
}

type QWebEngineFullScreenRequest_ITF interface {
	QWebEngineFullScreenRequest_PTR() *QWebEngineFullScreenRequest
}

func (ptr *QWebEngineFullScreenRequest) QWebEngineFullScreenRequest_PTR() *QWebEngineFullScreenRequest {
	return ptr
}

func (ptr *QWebEngineFullScreenRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineFullScreenRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineFullScreenRequest(ptr QWebEngineFullScreenRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineFullScreenRequest_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineFullScreenRequestFromPointer(ptr unsafe.Pointer) *QWebEngineFullScreenRequest {
	var n = new(QWebEngineFullScreenRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineFullScreenRequest) DestroyQWebEngineFullScreenRequest() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineFullScreenRequest) Accept() {
	if ptr.Pointer() != nil {
		C.QWebEngineFullScreenRequest_Accept(ptr.Pointer())
	}
}

func (ptr *QWebEngineFullScreenRequest) Origin() *core.QUrl {
	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineFullScreenRequest_Origin(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineFullScreenRequest) Reject() {
	if ptr.Pointer() != nil {
		C.QWebEngineFullScreenRequest_Reject(ptr.Pointer())
	}
}

func (ptr *QWebEngineFullScreenRequest) ToggleOn() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineFullScreenRequest_ToggleOn(ptr.Pointer()) != 0
	}
	return false
}

type QWebEngineHistory struct {
	ptr unsafe.Pointer
}

type QWebEngineHistory_ITF interface {
	QWebEngineHistory_PTR() *QWebEngineHistory
}

func (ptr *QWebEngineHistory) QWebEngineHistory_PTR() *QWebEngineHistory {
	return ptr
}

func (ptr *QWebEngineHistory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineHistory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineHistory(ptr QWebEngineHistory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineHistory_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineHistoryFromPointer(ptr unsafe.Pointer) *QWebEngineHistory {
	var n = new(QWebEngineHistory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineHistory) DestroyQWebEngineHistory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineHistory) Back() {
	if ptr.Pointer() != nil {
		C.QWebEngineHistory_Back(ptr.Pointer())
	}
}

func (ptr *QWebEngineHistory) BackItem() *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_BackItem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistory) BackItems(maxItems int) []*QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*QWebEngineHistoryItem {
			var out = make([]*QWebEngineHistoryItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebEngineHistoryFromPointer(l.data).backItems_atList(i)
			}
			return out
		}(C.QWebEngineHistory_BackItems(ptr.Pointer(), C.int(int32(maxItems))))
	}
	return nil
}

func (ptr *QWebEngineHistory) CanGoBack() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineHistory_CanGoBack(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineHistory) CanGoForward() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineHistory_CanGoForward(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineHistory) Clear() {
	if ptr.Pointer() != nil {
		C.QWebEngineHistory_Clear(ptr.Pointer())
	}
}

func (ptr *QWebEngineHistory) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineHistory_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineHistory) CurrentItem() *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_CurrentItem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistory) CurrentItemIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineHistory_CurrentItemIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineHistory) Forward() {
	if ptr.Pointer() != nil {
		C.QWebEngineHistory_Forward(ptr.Pointer())
	}
}

func (ptr *QWebEngineHistory) ForwardItem() *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_ForwardItem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistory) ForwardItems(maxItems int) []*QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*QWebEngineHistoryItem {
			var out = make([]*QWebEngineHistoryItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebEngineHistoryFromPointer(l.data).forwardItems_atList(i)
			}
			return out
		}(C.QWebEngineHistory_ForwardItems(ptr.Pointer(), C.int(int32(maxItems))))
	}
	return nil
}

func (ptr *QWebEngineHistory) GoToItem(item QWebEngineHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHistory_GoToItem(ptr.Pointer(), PointerFromQWebEngineHistoryItem(item))
	}
}

func (ptr *QWebEngineHistory) ItemAt(i int) *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_ItemAt(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistory) Items() []*QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*QWebEngineHistoryItem {
			var out = make([]*QWebEngineHistoryItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebEngineHistoryFromPointer(l.data).items_atList(i)
			}
			return out
		}(C.QWebEngineHistory_Items(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistory) backItems_atList(i int) *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_backItems_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistory) forwardItems_atList(i int) *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_forwardItems_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistory) items_atList(i int) *QWebEngineHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_items_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
		return tmpValue
	}
	return nil
}

type QWebEngineHistoryItem struct {
	ptr unsafe.Pointer
}

type QWebEngineHistoryItem_ITF interface {
	QWebEngineHistoryItem_PTR() *QWebEngineHistoryItem
}

func (ptr *QWebEngineHistoryItem) QWebEngineHistoryItem_PTR() *QWebEngineHistoryItem {
	return ptr
}

func (ptr *QWebEngineHistoryItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineHistoryItem(ptr QWebEngineHistoryItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineHistoryItem_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineHistoryItemFromPointer(ptr unsafe.Pointer) *QWebEngineHistoryItem {
	var n = new(QWebEngineHistoryItem)
	n.SetPointer(ptr)
	return n
}
func NewQWebEngineHistoryItem(other QWebEngineHistoryItem_ITF) *QWebEngineHistoryItem {
	var tmpValue = NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistoryItem_NewQWebEngineHistoryItem(PointerFromQWebEngineHistoryItem(other)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineHistoryItem).DestroyQWebEngineHistoryItem)
	return tmpValue
}

func (ptr *QWebEngineHistoryItem) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineHistoryItem_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineHistoryItem) LastVisited() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QWebEngineHistoryItem_LastVisited(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) OriginalUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineHistoryItem_OriginalUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineHistoryItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineHistoryItem) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineHistoryItem_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) DestroyQWebEngineHistoryItem() {
	if ptr.Pointer() != nil {
		C.QWebEngineHistoryItem_DestroyQWebEngineHistoryItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineHistoryItem) IconUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineHistoryItem_IconUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) Swap(other QWebEngineHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHistoryItem_Swap(ptr.Pointer(), PointerFromQWebEngineHistoryItem(other))
	}
}

type QWebEnginePage struct {
	core.QObject
}

type QWebEnginePage_ITF interface {
	core.QObject_ITF
	QWebEnginePage_PTR() *QWebEnginePage
}

func (ptr *QWebEnginePage) QWebEnginePage_PTR() *QWebEnginePage {
	return ptr
}

func (ptr *QWebEnginePage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEnginePage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEnginePage(ptr QWebEnginePage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEnginePage_PTR().Pointer()
	}
	return nil
}

func NewQWebEnginePageFromPointer(ptr unsafe.Pointer) *QWebEnginePage {
	var n = new(QWebEnginePage)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebEnginePage__Feature
//QWebEnginePage::Feature
type QWebEnginePage__Feature int64

const (
	QWebEnginePage__Geolocation            QWebEnginePage__Feature = QWebEnginePage__Feature(1)
	QWebEnginePage__MediaAudioCapture      QWebEnginePage__Feature = QWebEnginePage__Feature(2)
	QWebEnginePage__MediaVideoCapture      QWebEnginePage__Feature = QWebEnginePage__Feature(3)
	QWebEnginePage__MediaAudioVideoCapture QWebEnginePage__Feature = QWebEnginePage__Feature(4)
	QWebEnginePage__MouseLock              QWebEnginePage__Feature = QWebEnginePage__Feature(5)
)

//go:generate stringer -type=QWebEnginePage__FileSelectionMode
//QWebEnginePage::FileSelectionMode
type QWebEnginePage__FileSelectionMode int64

const (
	QWebEnginePage__FileSelectOpen         QWebEnginePage__FileSelectionMode = QWebEnginePage__FileSelectionMode(0)
	QWebEnginePage__FileSelectOpenMultiple QWebEnginePage__FileSelectionMode = QWebEnginePage__FileSelectionMode(1)
)

//go:generate stringer -type=QWebEnginePage__FindFlag
//QWebEnginePage::FindFlag
type QWebEnginePage__FindFlag int64

const (
	QWebEnginePage__FindBackward        QWebEnginePage__FindFlag = QWebEnginePage__FindFlag(1)
	QWebEnginePage__FindCaseSensitively QWebEnginePage__FindFlag = QWebEnginePage__FindFlag(2)
)

//go:generate stringer -type=QWebEnginePage__JavaScriptConsoleMessageLevel
//QWebEnginePage::JavaScriptConsoleMessageLevel
type QWebEnginePage__JavaScriptConsoleMessageLevel int64

const (
	QWebEnginePage__InfoMessageLevel    QWebEnginePage__JavaScriptConsoleMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(0)
	QWebEnginePage__WarningMessageLevel QWebEnginePage__JavaScriptConsoleMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(1)
	QWebEnginePage__ErrorMessageLevel   QWebEnginePage__JavaScriptConsoleMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(2)
)

//go:generate stringer -type=QWebEnginePage__NavigationType
//QWebEnginePage::NavigationType
type QWebEnginePage__NavigationType int64

const (
	QWebEnginePage__NavigationTypeLinkClicked   QWebEnginePage__NavigationType = QWebEnginePage__NavigationType(0)
	QWebEnginePage__NavigationTypeTyped         QWebEnginePage__NavigationType = QWebEnginePage__NavigationType(1)
	QWebEnginePage__NavigationTypeFormSubmitted QWebEnginePage__NavigationType = QWebEnginePage__NavigationType(2)
	QWebEnginePage__NavigationTypeBackForward   QWebEnginePage__NavigationType = QWebEnginePage__NavigationType(3)
	QWebEnginePage__NavigationTypeReload        QWebEnginePage__NavigationType = QWebEnginePage__NavigationType(4)
	QWebEnginePage__NavigationTypeOther         QWebEnginePage__NavigationType = QWebEnginePage__NavigationType(5)
)

//go:generate stringer -type=QWebEnginePage__PermissionPolicy
//QWebEnginePage::PermissionPolicy
type QWebEnginePage__PermissionPolicy int64

const (
	QWebEnginePage__PermissionUnknown       QWebEnginePage__PermissionPolicy = QWebEnginePage__PermissionPolicy(0)
	QWebEnginePage__PermissionGrantedByUser QWebEnginePage__PermissionPolicy = QWebEnginePage__PermissionPolicy(1)
	QWebEnginePage__PermissionDeniedByUser  QWebEnginePage__PermissionPolicy = QWebEnginePage__PermissionPolicy(2)
)

//go:generate stringer -type=QWebEnginePage__RenderProcessTerminationStatus
//QWebEnginePage::RenderProcessTerminationStatus
type QWebEnginePage__RenderProcessTerminationStatus int64

const (
	QWebEnginePage__NormalTerminationStatus   QWebEnginePage__RenderProcessTerminationStatus = QWebEnginePage__RenderProcessTerminationStatus(0)
	QWebEnginePage__AbnormalTerminationStatus QWebEnginePage__RenderProcessTerminationStatus = QWebEnginePage__RenderProcessTerminationStatus(1)
	QWebEnginePage__CrashedTerminationStatus  QWebEnginePage__RenderProcessTerminationStatus = QWebEnginePage__RenderProcessTerminationStatus(2)
	QWebEnginePage__KilledTerminationStatus   QWebEnginePage__RenderProcessTerminationStatus = QWebEnginePage__RenderProcessTerminationStatus(3)
)

//go:generate stringer -type=QWebEnginePage__WebAction
//QWebEnginePage::WebAction
type QWebEnginePage__WebAction int64

const (
	QWebEnginePage__NoWebAction                QWebEnginePage__WebAction = QWebEnginePage__WebAction(-1)
	QWebEnginePage__Back                       QWebEnginePage__WebAction = QWebEnginePage__WebAction(0)
	QWebEnginePage__Forward                    QWebEnginePage__WebAction = QWebEnginePage__WebAction(1)
	QWebEnginePage__Stop                       QWebEnginePage__WebAction = QWebEnginePage__WebAction(2)
	QWebEnginePage__Reload                     QWebEnginePage__WebAction = QWebEnginePage__WebAction(3)
	QWebEnginePage__Cut                        QWebEnginePage__WebAction = QWebEnginePage__WebAction(4)
	QWebEnginePage__Copy                       QWebEnginePage__WebAction = QWebEnginePage__WebAction(5)
	QWebEnginePage__Paste                      QWebEnginePage__WebAction = QWebEnginePage__WebAction(6)
	QWebEnginePage__Undo                       QWebEnginePage__WebAction = QWebEnginePage__WebAction(7)
	QWebEnginePage__Redo                       QWebEnginePage__WebAction = QWebEnginePage__WebAction(8)
	QWebEnginePage__SelectAll                  QWebEnginePage__WebAction = QWebEnginePage__WebAction(9)
	QWebEnginePage__ReloadAndBypassCache       QWebEnginePage__WebAction = QWebEnginePage__WebAction(10)
	QWebEnginePage__PasteAndMatchStyle         QWebEnginePage__WebAction = QWebEnginePage__WebAction(11)
	QWebEnginePage__OpenLinkInThisWindow       QWebEnginePage__WebAction = QWebEnginePage__WebAction(12)
	QWebEnginePage__OpenLinkInNewWindow        QWebEnginePage__WebAction = QWebEnginePage__WebAction(13)
	QWebEnginePage__OpenLinkInNewTab           QWebEnginePage__WebAction = QWebEnginePage__WebAction(14)
	QWebEnginePage__CopyLinkToClipboard        QWebEnginePage__WebAction = QWebEnginePage__WebAction(15)
	QWebEnginePage__DownloadLinkToDisk         QWebEnginePage__WebAction = QWebEnginePage__WebAction(16)
	QWebEnginePage__CopyImageToClipboard       QWebEnginePage__WebAction = QWebEnginePage__WebAction(17)
	QWebEnginePage__CopyImageUrlToClipboard    QWebEnginePage__WebAction = QWebEnginePage__WebAction(18)
	QWebEnginePage__DownloadImageToDisk        QWebEnginePage__WebAction = QWebEnginePage__WebAction(19)
	QWebEnginePage__CopyMediaUrlToClipboard    QWebEnginePage__WebAction = QWebEnginePage__WebAction(20)
	QWebEnginePage__ToggleMediaControls        QWebEnginePage__WebAction = QWebEnginePage__WebAction(21)
	QWebEnginePage__ToggleMediaLoop            QWebEnginePage__WebAction = QWebEnginePage__WebAction(22)
	QWebEnginePage__ToggleMediaPlayPause       QWebEnginePage__WebAction = QWebEnginePage__WebAction(23)
	QWebEnginePage__ToggleMediaMute            QWebEnginePage__WebAction = QWebEnginePage__WebAction(24)
	QWebEnginePage__DownloadMediaToDisk        QWebEnginePage__WebAction = QWebEnginePage__WebAction(25)
	QWebEnginePage__InspectElement             QWebEnginePage__WebAction = QWebEnginePage__WebAction(26)
	QWebEnginePage__ExitFullScreen             QWebEnginePage__WebAction = QWebEnginePage__WebAction(27)
	QWebEnginePage__RequestClose               QWebEnginePage__WebAction = QWebEnginePage__WebAction(28)
	QWebEnginePage__Unselect                   QWebEnginePage__WebAction = QWebEnginePage__WebAction(29)
	QWebEnginePage__SavePage                   QWebEnginePage__WebAction = QWebEnginePage__WebAction(30)
	QWebEnginePage__OpenLinkInNewBackgroundTab QWebEnginePage__WebAction = QWebEnginePage__WebAction(31)
	QWebEnginePage__WebActionCount             QWebEnginePage__WebAction = QWebEnginePage__WebAction(32)
)

//go:generate stringer -type=QWebEnginePage__WebWindowType
//QWebEnginePage::WebWindowType
type QWebEnginePage__WebWindowType int64

const (
	QWebEnginePage__WebBrowserWindow        QWebEnginePage__WebWindowType = QWebEnginePage__WebWindowType(0)
	QWebEnginePage__WebBrowserTab           QWebEnginePage__WebWindowType = QWebEnginePage__WebWindowType(1)
	QWebEnginePage__WebDialog               QWebEnginePage__WebWindowType = QWebEnginePage__WebWindowType(2)
	QWebEnginePage__WebBrowserBackgroundTab QWebEnginePage__WebWindowType = QWebEnginePage__WebWindowType(3)
)

func NewQWebEnginePage(parent core.QObject_ITF) *QWebEnginePage {
	var tmpValue = NewQWebEnginePageFromPointer(C.QWebEnginePage_NewQWebEnginePage(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEnginePage_AcceptNavigationRequest
func callbackQWebEnginePage_AcceptNavigationRequest(ptr unsafe.Pointer, url unsafe.Pointer, ty C.longlong, isMainFrame C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::acceptNavigationRequest"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl, QWebEnginePage__NavigationType, bool) bool)(core.NewQUrlFromPointer(url), QWebEnginePage__NavigationType(ty), int8(isMainFrame) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).AcceptNavigationRequestDefault(core.NewQUrlFromPointer(url), QWebEnginePage__NavigationType(ty), int8(isMainFrame) != 0))))
}

func (ptr *QWebEnginePage) ConnectAcceptNavigationRequest(f func(url *core.QUrl, ty QWebEnginePage__NavigationType, isMainFrame bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::acceptNavigationRequest", f)
	}
}

func (ptr *QWebEnginePage) DisconnectAcceptNavigationRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::acceptNavigationRequest")
	}
}

func (ptr *QWebEnginePage) AcceptNavigationRequest(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_AcceptNavigationRequest(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty), C.char(int8(qt.GoBoolToInt(isMainFrame)))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) AcceptNavigationRequestDefault(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_AcceptNavigationRequestDefault(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty), C.char(int8(qt.GoBoolToInt(isMainFrame)))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) Action(action QWebEnginePage__WebAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebEnginePage_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) BackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QWebEnginePage_BackgroundColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_CertificateError
func callbackQWebEnginePage_CertificateError(ptr unsafe.Pointer, certificateError unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::certificateError"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QWebEngineCertificateError) bool)(NewQWebEngineCertificateErrorFromPointer(certificateError)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).CertificateErrorDefault(NewQWebEngineCertificateErrorFromPointer(certificateError)))))
}

func (ptr *QWebEnginePage) ConnectCertificateError(f func(certificateError *QWebEngineCertificateError) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::certificateError", f)
	}
}

func (ptr *QWebEnginePage) DisconnectCertificateError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::certificateError")
	}
}

func (ptr *QWebEnginePage) CertificateError(certificateError QWebEngineCertificateError_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_CertificateError(ptr.Pointer(), PointerFromQWebEngineCertificateError(certificateError)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) CertificateErrorDefault(certificateError QWebEngineCertificateError_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_CertificateErrorDefault(ptr.Pointer(), PointerFromQWebEngineCertificateError(certificateError)) != 0
	}
	return false
}

//export callbackQWebEnginePage_ChooseFiles
func callbackQWebEnginePage_ChooseFiles(ptr unsafe.Pointer, mode C.longlong, oldFiles C.struct_QtWebEngine_PackedString, acceptedMimeTypes C.struct_QtWebEngine_PackedString) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::chooseFiles"); signal != nil {
		return C.CString(strings.Join(signal.(func(QWebEnginePage__FileSelectionMode, []string, []string) []string)(QWebEnginePage__FileSelectionMode(mode), strings.Split(cGoUnpackString(oldFiles), "|"), strings.Split(cGoUnpackString(acceptedMimeTypes), "|")), "|"))
	}

	return C.CString(strings.Join(NewQWebEnginePageFromPointer(ptr).ChooseFilesDefault(QWebEnginePage__FileSelectionMode(mode), strings.Split(cGoUnpackString(oldFiles), "|"), strings.Split(cGoUnpackString(acceptedMimeTypes), "|")), "|"))
}

func (ptr *QWebEnginePage) ConnectChooseFiles(f func(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::chooseFiles", f)
	}
}

func (ptr *QWebEnginePage) DisconnectChooseFiles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::chooseFiles")
	}
}

func (ptr *QWebEnginePage) ChooseFiles(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
	if ptr.Pointer() != nil {
		var oldFilesC = C.CString(strings.Join(oldFiles, "|"))
		defer C.free(unsafe.Pointer(oldFilesC))
		var acceptedMimeTypesC = C.CString(strings.Join(acceptedMimeTypes, "|"))
		defer C.free(unsafe.Pointer(acceptedMimeTypesC))
		return strings.Split(cGoUnpackString(C.QWebEnginePage_ChooseFiles(ptr.Pointer(), C.longlong(mode), oldFilesC, acceptedMimeTypesC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebEnginePage) ChooseFilesDefault(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
	if ptr.Pointer() != nil {
		var oldFilesC = C.CString(strings.Join(oldFiles, "|"))
		defer C.free(unsafe.Pointer(oldFilesC))
		var acceptedMimeTypesC = C.CString(strings.Join(acceptedMimeTypes, "|"))
		defer C.free(unsafe.Pointer(acceptedMimeTypesC))
		return strings.Split(cGoUnpackString(C.QWebEnginePage_ChooseFilesDefault(ptr.Pointer(), C.longlong(mode), oldFilesC, acceptedMimeTypesC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebEnginePage) ContentsSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFFromPointer(C.QWebEnginePage_ContentsSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) CreateStandardContextMenu() *widgets.QMenu {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQMenuFromPointer(C.QWebEnginePage_CreateStandardContextMenu(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_CreateWindow
func callbackQWebEnginePage_CreateWindow(ptr unsafe.Pointer, ty C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::createWindow"); signal != nil {
		return PointerFromQWebEnginePage(signal.(func(QWebEnginePage__WebWindowType) *QWebEnginePage)(QWebEnginePage__WebWindowType(ty)))
	}

	return PointerFromQWebEnginePage(NewQWebEnginePageFromPointer(ptr).CreateWindowDefault(QWebEnginePage__WebWindowType(ty)))
}

func (ptr *QWebEnginePage) ConnectCreateWindow(f func(ty QWebEnginePage__WebWindowType) *QWebEnginePage) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::createWindow", f)
	}
}

func (ptr *QWebEnginePage) DisconnectCreateWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::createWindow")
	}
}

func (ptr *QWebEnginePage) CreateWindow(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEnginePageFromPointer(C.QWebEnginePage_CreateWindow(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) CreateWindowDefault(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEnginePageFromPointer(C.QWebEnginePage_CreateWindowDefault(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEnginePage) History() *QWebEngineHistory {
	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryFromPointer(C.QWebEnginePage_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QWebEnginePage_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) IconUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEnginePage_IconUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) IsAudioMuted() bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_IsAudioMuted(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebEnginePage_JavaScriptAlert
func callbackQWebEnginePage_JavaScriptAlert(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, msg C.struct_QtWebEngine_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::javaScriptAlert"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg))
	} else {
		NewQWebEnginePageFromPointer(ptr).JavaScriptAlertDefault(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg))
	}
}

func (ptr *QWebEnginePage) ConnectJavaScriptAlert(f func(securityOrigin *core.QUrl, msg string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptAlert", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptAlert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptAlert")
	}
}

func (ptr *QWebEnginePage) JavaScriptAlert(securityOrigin core.QUrl_ITF, msg string) {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QWebEnginePage_JavaScriptAlert(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), msgC)
	}
}

func (ptr *QWebEnginePage) JavaScriptAlertDefault(securityOrigin core.QUrl_ITF, msg string) {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QWebEnginePage_JavaScriptAlertDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), msgC)
	}
}

//export callbackQWebEnginePage_JavaScriptConfirm
func callbackQWebEnginePage_JavaScriptConfirm(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, msg C.struct_QtWebEngine_PackedString) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::javaScriptConfirm"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl, string) bool)(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).JavaScriptConfirmDefault(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg)))))
}

func (ptr *QWebEnginePage) ConnectJavaScriptConfirm(f func(securityOrigin *core.QUrl, msg string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptConfirm", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConfirm() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptConfirm")
	}
}

func (ptr *QWebEnginePage) JavaScriptConfirm(securityOrigin core.QUrl_ITF, msg string) bool {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		return C.QWebEnginePage_JavaScriptConfirm(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), msgC) != 0
	}
	return false
}

func (ptr *QWebEnginePage) JavaScriptConfirmDefault(securityOrigin core.QUrl_ITF, msg string) bool {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		return C.QWebEnginePage_JavaScriptConfirmDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), msgC) != 0
	}
	return false
}

//export callbackQWebEnginePage_JavaScriptConsoleMessage
func callbackQWebEnginePage_JavaScriptConsoleMessage(ptr unsafe.Pointer, level C.longlong, message C.struct_QtWebEngine_PackedString, lineNumber C.int, sourceID C.struct_QtWebEngine_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::javaScriptConsoleMessage"); signal != nil {
		signal.(func(QWebEnginePage__JavaScriptConsoleMessageLevel, string, int, string))(QWebEnginePage__JavaScriptConsoleMessageLevel(level), cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	} else {
		NewQWebEnginePageFromPointer(ptr).JavaScriptConsoleMessageDefault(QWebEnginePage__JavaScriptConsoleMessageLevel(level), cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	}
}

func (ptr *QWebEnginePage) ConnectJavaScriptConsoleMessage(f func(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptConsoleMessage", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConsoleMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptConsoleMessage")
	}
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessage(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		var sourceIDC = C.CString(sourceID)
		defer C.free(unsafe.Pointer(sourceIDC))
		C.QWebEnginePage_JavaScriptConsoleMessage(ptr.Pointer(), C.longlong(level), messageC, C.int(int32(lineNumber)), sourceIDC)
	}
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessageDefault(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		var sourceIDC = C.CString(sourceID)
		defer C.free(unsafe.Pointer(sourceIDC))
		C.QWebEnginePage_JavaScriptConsoleMessageDefault(ptr.Pointer(), C.longlong(level), messageC, C.int(int32(lineNumber)), sourceIDC)
	}
}

//export callbackQWebEnginePage_JavaScriptPrompt
func callbackQWebEnginePage_JavaScriptPrompt(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, msg C.struct_QtWebEngine_PackedString, defaultValue C.struct_QtWebEngine_PackedString, result C.struct_QtWebEngine_PackedString) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::javaScriptPrompt"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl, string, string, string) bool)(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg), cGoUnpackString(defaultValue), cGoUnpackString(result)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).JavaScriptPromptDefault(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg), cGoUnpackString(defaultValue), cGoUnpackString(result)))))
}

func (ptr *QWebEnginePage) ConnectJavaScriptPrompt(f func(securityOrigin *core.QUrl, msg string, defaultValue string, result string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptPrompt", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptPrompt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::javaScriptPrompt")
	}
}

func (ptr *QWebEnginePage) JavaScriptPrompt(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		var defaultValueC = C.CString(defaultValue)
		defer C.free(unsafe.Pointer(defaultValueC))
		var resultC = C.CString(result)
		defer C.free(unsafe.Pointer(resultC))
		return C.QWebEnginePage_JavaScriptPrompt(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), msgC, defaultValueC, resultC) != 0
	}
	return false
}

func (ptr *QWebEnginePage) JavaScriptPromptDefault(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		var defaultValueC = C.CString(defaultValue)
		defer C.free(unsafe.Pointer(defaultValueC))
		var resultC = C.CString(result)
		defer C.free(unsafe.Pointer(resultC))
		return C.QWebEnginePage_JavaScriptPromptDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), msgC, defaultValueC, resultC) != 0
	}
	return false
}

func (ptr *QWebEnginePage) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) RecentlyAudible() bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_RecentlyAudible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEnginePage) RequestedUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEnginePage_RequestedUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) RunJavaScript4(scriptSource string) {
	if ptr.Pointer() != nil {
		var scriptSourceC = C.CString(scriptSource)
		defer C.free(unsafe.Pointer(scriptSourceC))
		C.QWebEnginePage_RunJavaScript4(ptr.Pointer(), scriptSourceC)
	}
}

func (ptr *QWebEnginePage) RunJavaScript2(scriptSource string, worldId uint) {
	if ptr.Pointer() != nil {
		var scriptSourceC = C.CString(scriptSource)
		defer C.free(unsafe.Pointer(scriptSourceC))
		C.QWebEnginePage_RunJavaScript2(ptr.Pointer(), scriptSourceC, C.uint(uint32(worldId)))
	}
}

func (ptr *QWebEnginePage) ScrollPosition() *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QWebEnginePage_ScrollPosition(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEnginePage_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEnginePage) SetAudioMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetAudioMuted(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(muted))))
	}
}

func (ptr *QWebEnginePage) SetBackgroundColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QWebEnginePage) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var mimeTypeC = C.CString(mimeType)
		defer C.free(unsafe.Pointer(mimeTypeC))
		C.QWebEnginePage_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), mimeTypeC, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEnginePage) SetFeaturePermission(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature, policy QWebEnginePage__PermissionPolicy) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetFeaturePermission(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.longlong(feature), C.longlong(policy))
	}
}

func (ptr *QWebEnginePage) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC = C.CString(html)
		defer C.free(unsafe.Pointer(htmlC))
		C.QWebEnginePage_SetHtml(ptr.Pointer(), htmlC, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEnginePage) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) SetView(view widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetView(ptr.Pointer(), widgets.PointerFromQWidget(view))
	}
}

func (ptr *QWebEnginePage) SetZoomFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebEnginePage) Settings() *QWebEngineSettings {
	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEnginePage_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEnginePage_Title(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebEnginePage_TriggerAction
func callbackQWebEnginePage_TriggerAction(ptr unsafe.Pointer, action C.longlong, checked C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::triggerAction"); signal != nil {
		signal.(func(QWebEnginePage__WebAction, bool))(QWebEnginePage__WebAction(action), int8(checked) != 0)
	} else {
		NewQWebEnginePageFromPointer(ptr).TriggerActionDefault(QWebEnginePage__WebAction(action), int8(checked) != 0)
	}
}

func (ptr *QWebEnginePage) ConnectTriggerAction(f func(action QWebEnginePage__WebAction, checked bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::triggerAction", f)
	}
}

func (ptr *QWebEnginePage) DisconnectTriggerAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::triggerAction")
	}
}

func (ptr *QWebEnginePage) TriggerAction(action QWebEnginePage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_TriggerAction(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

func (ptr *QWebEnginePage) TriggerActionDefault(action QWebEnginePage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_TriggerActionDefault(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

func (ptr *QWebEnginePage) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEnginePage_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) View() *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QWebEnginePage_View(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebEnginePage_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEnginePage) DestroyQWebEnginePage() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DestroyQWebEnginePage(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQWebEnginePage2(profile QWebEngineProfile_ITF, parent core.QObject_ITF) *QWebEnginePage {
	var tmpValue = NewQWebEnginePageFromPointer(C.QWebEnginePage_NewQWebEnginePage2(PointerFromQWebEngineProfile(profile), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEnginePage_AudioMutedChanged
func callbackQWebEnginePage_AudioMutedChanged(ptr unsafe.Pointer, muted C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::audioMutedChanged"); signal != nil {
		signal.(func(bool))(int8(muted) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectAudioMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectAudioMutedChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::audioMutedChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectAudioMutedChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectAudioMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::audioMutedChanged")
	}
}

func (ptr *QWebEnginePage) AudioMutedChanged(muted bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_AudioMutedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(muted))))
	}
}

//export callbackQWebEnginePage_AuthenticationRequired
func callbackQWebEnginePage_AuthenticationRequired(ptr unsafe.Pointer, requestUrl unsafe.Pointer, authenticator unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::authenticationRequired"); signal != nil {
		signal.(func(*core.QUrl, *network.QAuthenticator))(core.NewQUrlFromPointer(requestUrl), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebEnginePage) ConnectAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::authenticationRequired", f)
	}
}

func (ptr *QWebEnginePage) DisconnectAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::authenticationRequired")
	}
}

func (ptr *QWebEnginePage) AuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_AuthenticationRequired(ptr.Pointer(), core.PointerFromQUrl(requestUrl), network.PointerFromQAuthenticator(authenticator))
	}
}

//export callbackQWebEnginePage_ContentsSizeChanged
func callbackQWebEnginePage_ContentsSizeChanged(ptr unsafe.Pointer, size unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::contentsSizeChanged"); signal != nil {
		signal.(func(*core.QSizeF))(core.NewQSizeFFromPointer(size))
	}

}

func (ptr *QWebEnginePage) ConnectContentsSizeChanged(f func(size *core.QSizeF)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectContentsSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::contentsSizeChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::contentsSizeChanged")
	}
}

func (ptr *QWebEnginePage) ContentsSizeChanged(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ContentsSizeChanged(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QWebEnginePage) ContextMenuData() *QWebEngineContextMenuData {
	if ptr.Pointer() != nil {
		return NewQWebEngineContextMenuDataFromPointer(C.QWebEnginePage_ContextMenuData(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_Event
func callbackQWebEnginePage_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEnginePage) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::event", f)
	}
}

func (ptr *QWebEnginePage) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::event")
	}
}

func (ptr *QWebEnginePage) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEnginePage_FeaturePermissionRequestCanceled
func callbackQWebEnginePage_FeaturePermissionRequestCanceled(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, feature C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::featurePermissionRequestCanceled"); signal != nil {
		signal.(func(*core.QUrl, QWebEnginePage__Feature))(core.NewQUrlFromPointer(securityOrigin), QWebEnginePage__Feature(feature))
	}

}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequestCanceled(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::featurePermissionRequestCanceled", f)
	}
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequestCanceled() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::featurePermissionRequestCanceled")
	}
}

func (ptr *QWebEnginePage) FeaturePermissionRequestCanceled(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_FeaturePermissionRequestCanceled(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.longlong(feature))
	}
}

//export callbackQWebEnginePage_FeaturePermissionRequested
func callbackQWebEnginePage_FeaturePermissionRequested(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, feature C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::featurePermissionRequested"); signal != nil {
		signal.(func(*core.QUrl, QWebEnginePage__Feature))(core.NewQUrlFromPointer(securityOrigin), QWebEnginePage__Feature(feature))
	}

}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequested(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectFeaturePermissionRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::featurePermissionRequested", f)
	}
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectFeaturePermissionRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::featurePermissionRequested")
	}
}

func (ptr *QWebEnginePage) FeaturePermissionRequested(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_FeaturePermissionRequested(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.longlong(feature))
	}
}

func (ptr *QWebEnginePage) FindText(subString string, options QWebEnginePage__FindFlag) {
	if ptr.Pointer() != nil {
		var subStringC = C.CString(subString)
		defer C.free(unsafe.Pointer(subStringC))
		C.QWebEnginePage_FindText(ptr.Pointer(), subStringC, C.longlong(options))
	}
}

//export callbackQWebEnginePage_GeometryChangeRequested
func callbackQWebEnginePage_GeometryChangeRequested(ptr unsafe.Pointer, geom unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::geometryChangeRequested"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(geom))
	}

}

func (ptr *QWebEnginePage) ConnectGeometryChangeRequested(f func(geom *core.QRect)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectGeometryChangeRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::geometryChangeRequested", f)
	}
}

func (ptr *QWebEnginePage) DisconnectGeometryChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectGeometryChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::geometryChangeRequested")
	}
}

func (ptr *QWebEnginePage) GeometryChangeRequested(geom core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_GeometryChangeRequested(ptr.Pointer(), core.PointerFromQRect(geom))
	}
}

//export callbackQWebEnginePage_IconChanged
func callbackQWebEnginePage_IconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::iconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

func (ptr *QWebEnginePage) ConnectIconChanged(f func(icon *gui.QIcon)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectIconChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::iconChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::iconChanged")
	}
}

func (ptr *QWebEnginePage) IconChanged(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_IconChanged(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

//export callbackQWebEnginePage_IconUrlChanged
func callbackQWebEnginePage_IconUrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::iconUrlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEnginePage) ConnectIconUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectIconUrlChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::iconUrlChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectIconUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectIconUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::iconUrlChanged")
	}
}

func (ptr *QWebEnginePage) IconUrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_IconUrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebEnginePage_LinkHovered
func callbackQWebEnginePage_LinkHovered(ptr unsafe.Pointer, url C.struct_QtWebEngine_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::linkHovered"); signal != nil {
		signal.(func(string))(cGoUnpackString(url))
	}

}

func (ptr *QWebEnginePage) ConnectLinkHovered(f func(url string)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::linkHovered", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::linkHovered")
	}
}

func (ptr *QWebEnginePage) LinkHovered(url string) {
	if ptr.Pointer() != nil {
		var urlC = C.CString(url)
		defer C.free(unsafe.Pointer(urlC))
		C.QWebEnginePage_LinkHovered(ptr.Pointer(), urlC)
	}
}

//export callbackQWebEnginePage_LoadFinished
func callbackQWebEnginePage_LoadFinished(ptr unsafe.Pointer, ok C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::loadFinished"); signal != nil {
		signal.(func(bool))(int8(ok) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::loadFinished", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::loadFinished")
	}
}

func (ptr *QWebEnginePage) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQWebEnginePage_LoadProgress
func callbackQWebEnginePage_LoadProgress(ptr unsafe.Pointer, progress C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::loadProgress"); signal != nil {
		signal.(func(int))(int(int32(progress)))
	}

}

func (ptr *QWebEnginePage) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::loadProgress", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::loadProgress")
	}
}

func (ptr *QWebEnginePage) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQWebEnginePage_LoadStarted
func callbackQWebEnginePage_LoadStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEnginePage) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::loadStarted", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::loadStarted")
	}
}

func (ptr *QWebEnginePage) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadStarted(ptr.Pointer())
	}
}

func (ptr *QWebEnginePage) PrintToPdf(filePath string, pageLayout gui.QPageLayout_ITF) {
	if ptr.Pointer() != nil {
		var filePathC = C.CString(filePath)
		defer C.free(unsafe.Pointer(filePathC))
		C.QWebEnginePage_PrintToPdf(ptr.Pointer(), filePathC, gui.PointerFromQPageLayout(pageLayout))
	}
}

func (ptr *QWebEnginePage) Profile() *QWebEngineProfile {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineProfileFromPointer(C.QWebEnginePage_Profile(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_ProxyAuthenticationRequired
func callbackQWebEnginePage_ProxyAuthenticationRequired(ptr unsafe.Pointer, requestUrl unsafe.Pointer, authenticator unsafe.Pointer, proxyHost C.struct_QtWebEngine_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::proxyAuthenticationRequired"); signal != nil {
		signal.(func(*core.QUrl, *network.QAuthenticator, string))(core.NewQUrlFromPointer(requestUrl), network.NewQAuthenticatorFromPointer(authenticator), cGoUnpackString(proxyHost))
	}

}

func (ptr *QWebEnginePage) ConnectProxyAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator, proxyHost string)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::proxyAuthenticationRequired", f)
	}
}

func (ptr *QWebEnginePage) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::proxyAuthenticationRequired")
	}
}

func (ptr *QWebEnginePage) ProxyAuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF, proxyHost string) {
	if ptr.Pointer() != nil {
		var proxyHostC = C.CString(proxyHost)
		defer C.free(unsafe.Pointer(proxyHostC))
		C.QWebEnginePage_ProxyAuthenticationRequired(ptr.Pointer(), core.PointerFromQUrl(requestUrl), network.PointerFromQAuthenticator(authenticator), proxyHostC)
	}
}

//export callbackQWebEnginePage_RecentlyAudibleChanged
func callbackQWebEnginePage_RecentlyAudibleChanged(ptr unsafe.Pointer, recentlyAudible C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::recentlyAudibleChanged"); signal != nil {
		signal.(func(bool))(int8(recentlyAudible) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectRecentlyAudibleChanged(f func(recentlyAudible bool)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectRecentlyAudibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::recentlyAudibleChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectRecentlyAudibleChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectRecentlyAudibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::recentlyAudibleChanged")
	}
}

func (ptr *QWebEnginePage) RecentlyAudibleChanged(recentlyAudible bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_RecentlyAudibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(recentlyAudible))))
	}
}

//export callbackQWebEnginePage_RenderProcessTerminated
func callbackQWebEnginePage_RenderProcessTerminated(ptr unsafe.Pointer, terminationStatus C.longlong, exitCode C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::renderProcessTerminated"); signal != nil {
		signal.(func(QWebEnginePage__RenderProcessTerminationStatus, int))(QWebEnginePage__RenderProcessTerminationStatus(terminationStatus), int(int32(exitCode)))
	}

}

func (ptr *QWebEnginePage) ConnectRenderProcessTerminated(f func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectRenderProcessTerminated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::renderProcessTerminated", f)
	}
}

func (ptr *QWebEnginePage) DisconnectRenderProcessTerminated() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectRenderProcessTerminated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::renderProcessTerminated")
	}
}

func (ptr *QWebEnginePage) RenderProcessTerminated(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_RenderProcessTerminated(ptr.Pointer(), C.longlong(terminationStatus), C.int(int32(exitCode)))
	}
}

//export callbackQWebEnginePage_ScrollPositionChanged
func callbackQWebEnginePage_ScrollPositionChanged(ptr unsafe.Pointer, position unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::scrollPositionChanged"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(position))
	}

}

func (ptr *QWebEnginePage) ConnectScrollPositionChanged(f func(position *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectScrollPositionChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::scrollPositionChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectScrollPositionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectScrollPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::scrollPositionChanged")
	}
}

func (ptr *QWebEnginePage) ScrollPositionChanged(position core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ScrollPositionChanged(ptr.Pointer(), core.PointerFromQPointF(position))
	}
}

//export callbackQWebEnginePage_SelectionChanged
func callbackQWebEnginePage_SelectionChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEnginePage) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::selectionChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::selectionChanged")
	}
}

func (ptr *QWebEnginePage) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebEnginePage) SetWebChannel2(channel webchannel.QWebChannel_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetWebChannel2(ptr.Pointer(), webchannel.PointerFromQWebChannel(channel))
	}
}

func (ptr *QWebEnginePage) SetWebChannel(channel webchannel.QWebChannel_ITF, worldId uint) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetWebChannel(ptr.Pointer(), webchannel.PointerFromQWebChannel(channel), C.uint(uint32(worldId)))
	}
}

//export callbackQWebEnginePage_TitleChanged
func callbackQWebEnginePage_TitleChanged(ptr unsafe.Pointer, title C.struct_QtWebEngine_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::titleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

func (ptr *QWebEnginePage) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::titleChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::titleChanged")
	}
}

func (ptr *QWebEnginePage) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
		C.QWebEnginePage_TitleChanged(ptr.Pointer(), titleC)
	}
}

//export callbackQWebEnginePage_UrlChanged
func callbackQWebEnginePage_UrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEnginePage) ConnectUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::urlChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::urlChanged")
	}
}

func (ptr *QWebEnginePage) UrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) WebChannel() *webchannel.QWebChannel {
	if ptr.Pointer() != nil {
		var tmpValue = webchannel.NewQWebChannelFromPointer(C.QWebEnginePage_WebChannel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_WindowCloseRequested
func callbackQWebEnginePage_WindowCloseRequested(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::windowCloseRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEnginePage) ConnectWindowCloseRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectWindowCloseRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::windowCloseRequested", f)
	}
}

func (ptr *QWebEnginePage) DisconnectWindowCloseRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectWindowCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::windowCloseRequested")
	}
}

func (ptr *QWebEnginePage) WindowCloseRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_WindowCloseRequested(ptr.Pointer())
	}
}

//export callbackQWebEnginePage_TimerEvent
func callbackQWebEnginePage_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::timerEvent", f)
	}
}

func (ptr *QWebEnginePage) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::timerEvent")
	}
}

func (ptr *QWebEnginePage) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEnginePage) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEnginePage_ChildEvent
func callbackQWebEnginePage_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::childEvent", f)
	}
}

func (ptr *QWebEnginePage) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::childEvent")
	}
}

func (ptr *QWebEnginePage) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEnginePage) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEnginePage_ConnectNotify
func callbackQWebEnginePage_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEnginePageFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEnginePage) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::connectNotify", f)
	}
}

func (ptr *QWebEnginePage) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::connectNotify")
	}
}

func (ptr *QWebEnginePage) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEnginePage) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEnginePage_CustomEvent
func callbackQWebEnginePage_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::customEvent", f)
	}
}

func (ptr *QWebEnginePage) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::customEvent")
	}
}

func (ptr *QWebEnginePage) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEnginePage) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEnginePage_DeleteLater
func callbackQWebEnginePage_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEnginePageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEnginePage) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::deleteLater", f)
	}
}

func (ptr *QWebEnginePage) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::deleteLater")
	}
}

func (ptr *QWebEnginePage) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEnginePage) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEnginePage_DisconnectNotify
func callbackQWebEnginePage_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEnginePageFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEnginePage) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::disconnectNotify", f)
	}
}

func (ptr *QWebEnginePage) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::disconnectNotify")
	}
}

func (ptr *QWebEnginePage) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEnginePage) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEnginePage_EventFilter
func callbackQWebEnginePage_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEnginePage) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::eventFilter", f)
	}
}

func (ptr *QWebEnginePage) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::eventFilter")
	}
}

func (ptr *QWebEnginePage) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEnginePage_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEnginePage_MetaObject
func callbackQWebEnginePage_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEnginePage::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEnginePageFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEnginePage) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::metaObject", f)
	}
}

func (ptr *QWebEnginePage) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEnginePage::metaObject")
	}
}

func (ptr *QWebEnginePage) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEnginePage_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEnginePage_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineProfile struct {
	core.QObject
}

type QWebEngineProfile_ITF interface {
	core.QObject_ITF
	QWebEngineProfile_PTR() *QWebEngineProfile
}

func (ptr *QWebEngineProfile) QWebEngineProfile_PTR() *QWebEngineProfile {
	return ptr
}

func (ptr *QWebEngineProfile) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineProfile) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineProfile(ptr QWebEngineProfile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineProfile_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineProfileFromPointer(ptr unsafe.Pointer) *QWebEngineProfile {
	var n = new(QWebEngineProfile)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineProfile) DestroyQWebEngineProfile() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebEngineProfile__HttpCacheType
//QWebEngineProfile::HttpCacheType
type QWebEngineProfile__HttpCacheType int64

const (
	QWebEngineProfile__MemoryHttpCache QWebEngineProfile__HttpCacheType = QWebEngineProfile__HttpCacheType(0)
	QWebEngineProfile__DiskHttpCache   QWebEngineProfile__HttpCacheType = QWebEngineProfile__HttpCacheType(1)
	QWebEngineProfile__NoCache         QWebEngineProfile__HttpCacheType = QWebEngineProfile__HttpCacheType(2)
)

//go:generate stringer -type=QWebEngineProfile__PersistentCookiesPolicy
//QWebEngineProfile::PersistentCookiesPolicy
type QWebEngineProfile__PersistentCookiesPolicy int64

const (
	QWebEngineProfile__NoPersistentCookies    QWebEngineProfile__PersistentCookiesPolicy = QWebEngineProfile__PersistentCookiesPolicy(0)
	QWebEngineProfile__AllowPersistentCookies QWebEngineProfile__PersistentCookiesPolicy = QWebEngineProfile__PersistentCookiesPolicy(1)
	QWebEngineProfile__ForcePersistentCookies QWebEngineProfile__PersistentCookiesPolicy = QWebEngineProfile__PersistentCookiesPolicy(2)
)

func NewQWebEngineProfile(parent core.QObject_ITF) *QWebEngineProfile {
	var tmpValue = NewQWebEngineProfileFromPointer(C.QWebEngineProfile_NewQWebEngineProfile(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQWebEngineProfile2(storageName string, parent core.QObject_ITF) *QWebEngineProfile {
	var storageNameC = C.CString(storageName)
	defer C.free(unsafe.Pointer(storageNameC))
	var tmpValue = NewQWebEngineProfileFromPointer(C.QWebEngineProfile_NewQWebEngineProfile2(storageNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebEngineProfile) CachePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_CachePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) ClearAllVisitedLinks() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ClearAllVisitedLinks(ptr.Pointer())
	}
}

func (ptr *QWebEngineProfile) ClearHttpCache() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ClearHttpCache(ptr.Pointer())
	}
}

func (ptr *QWebEngineProfile) CookieStore() *QWebEngineCookieStore {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineCookieStoreFromPointer(C.QWebEngineProfile_CookieStore(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QWebEngineProfile_DefaultProfile() *QWebEngineProfile {
	var tmpValue = NewQWebEngineProfileFromPointer(C.QWebEngineProfile_QWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebEngineProfile) DefaultProfile() *QWebEngineProfile {
	var tmpValue = NewQWebEngineProfileFromPointer(C.QWebEngineProfile_QWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineProfile_DownloadRequested
func callbackQWebEngineProfile_DownloadRequested(ptr unsafe.Pointer, download unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::downloadRequested"); signal != nil {
		signal.(func(*QWebEngineDownloadItem))(NewQWebEngineDownloadItemFromPointer(download))
	}

}

func (ptr *QWebEngineProfile) ConnectDownloadRequested(f func(download *QWebEngineDownloadItem)) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectDownloadRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::downloadRequested", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectDownloadRequested() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectDownloadRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::downloadRequested")
	}
}

func (ptr *QWebEngineProfile) DownloadRequested(download QWebEngineDownloadItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DownloadRequested(ptr.Pointer(), PointerFromQWebEngineDownloadItem(download))
	}
}

func (ptr *QWebEngineProfile) HttpAcceptLanguage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_HttpAcceptLanguage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) HttpCacheMaximumSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineProfile_HttpCacheMaximumSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineProfile) HttpCacheType() QWebEngineProfile__HttpCacheType {
	if ptr.Pointer() != nil {
		return QWebEngineProfile__HttpCacheType(C.QWebEngineProfile_HttpCacheType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineProfile) HttpUserAgent() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_HttpUserAgent(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) InstallUrlSchemeHandler(scheme core.QByteArray_ITF, handler QWebEngineUrlSchemeHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_InstallUrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(scheme), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

func (ptr *QWebEngineProfile) IsOffTheRecord() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_IsOffTheRecord(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) PersistentCookiesPolicy() QWebEngineProfile__PersistentCookiesPolicy {
	if ptr.Pointer() != nil {
		return QWebEngineProfile__PersistentCookiesPolicy(C.QWebEngineProfile_PersistentCookiesPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineProfile) PersistentStoragePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_PersistentStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) RemoveAllUrlSchemeHandlers() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveAllUrlSchemeHandlers(ptr.Pointer())
	}
}

func (ptr *QWebEngineProfile) RemoveUrlScheme(scheme core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveUrlScheme(ptr.Pointer(), core.PointerFromQByteArray(scheme))
	}
}

func (ptr *QWebEngineProfile) RemoveUrlSchemeHandler(handler QWebEngineUrlSchemeHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveUrlSchemeHandler(ptr.Pointer(), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

func (ptr *QWebEngineProfile) Scripts() *QWebEngineScriptCollection {
	if ptr.Pointer() != nil {
		return NewQWebEngineScriptCollectionFromPointer(C.QWebEngineProfile_Scripts(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) SetCachePath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QWebEngineProfile_SetCachePath(ptr.Pointer(), pathC)
	}
}

func (ptr *QWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	if ptr.Pointer() != nil {
		var httpAcceptLanguageC = C.CString(httpAcceptLanguage)
		defer C.free(unsafe.Pointer(httpAcceptLanguageC))
		C.QWebEngineProfile_SetHttpAcceptLanguage(ptr.Pointer(), httpAcceptLanguageC)
	}
}

func (ptr *QWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpCacheMaximumSize(ptr.Pointer(), C.int(int32(maxSize)))
	}
}

func (ptr *QWebEngineProfile) SetHttpCacheType(httpCacheType QWebEngineProfile__HttpCacheType) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpCacheType(ptr.Pointer(), C.longlong(httpCacheType))
	}
}

func (ptr *QWebEngineProfile) SetHttpUserAgent(userAgent string) {
	if ptr.Pointer() != nil {
		var userAgentC = C.CString(userAgent)
		defer C.free(unsafe.Pointer(userAgentC))
		C.QWebEngineProfile_SetHttpUserAgent(ptr.Pointer(), userAgentC)
	}
}

func (ptr *QWebEngineProfile) SetPersistentCookiesPolicy(newPersistentCookiesPolicy QWebEngineProfile__PersistentCookiesPolicy) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetPersistentCookiesPolicy(ptr.Pointer(), C.longlong(newPersistentCookiesPolicy))
	}
}

func (ptr *QWebEngineProfile) SetPersistentStoragePath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QWebEngineProfile_SetPersistentStoragePath(ptr.Pointer(), pathC)
	}
}

func (ptr *QWebEngineProfile) SetRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

func (ptr *QWebEngineProfile) Settings() *QWebEngineSettings {
	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEngineProfile_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) StorageName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_StorageName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) UrlSchemeHandler(scheme core.QByteArray_ITF) *QWebEngineUrlSchemeHandler {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineUrlSchemeHandlerFromPointer(C.QWebEngineProfile_UrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(scheme)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) VisitedLinksContainsUrl(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_VisitedLinksContainsUrl(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

//export callbackQWebEngineProfile_TimerEvent
func callbackQWebEngineProfile_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::timerEvent", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::timerEvent")
	}
}

func (ptr *QWebEngineProfile) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineProfile_ChildEvent
func callbackQWebEngineProfile_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::childEvent", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::childEvent")
	}
}

func (ptr *QWebEngineProfile) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineProfile_ConnectNotify
func callbackQWebEngineProfile_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineProfileFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::connectNotify", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::connectNotify")
	}
}

func (ptr *QWebEngineProfile) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineProfile_CustomEvent
func callbackQWebEngineProfile_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::customEvent", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::customEvent")
	}
}

func (ptr *QWebEngineProfile) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineProfile_DeleteLater
func callbackQWebEngineProfile_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineProfileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineProfile) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::deleteLater", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::deleteLater")
	}
}

func (ptr *QWebEngineProfile) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineProfile) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineProfile_DisconnectNotify
func callbackQWebEngineProfile_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineProfileFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::disconnectNotify", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::disconnectNotify")
	}
}

func (ptr *QWebEngineProfile) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineProfile_Event
func callbackQWebEngineProfile_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineProfileFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineProfile) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::event", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::event")
	}
}

func (ptr *QWebEngineProfile) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineProfile_EventFilter
func callbackQWebEngineProfile_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineProfileFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineProfile) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::eventFilter", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::eventFilter")
	}
}

func (ptr *QWebEngineProfile) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineProfile_MetaObject
func callbackQWebEngineProfile_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineProfile::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineProfileFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineProfile) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::metaObject", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineProfile::metaObject")
	}
}

func (ptr *QWebEngineProfile) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineProfile_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineProfile_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineScript struct {
	ptr unsafe.Pointer
}

type QWebEngineScript_ITF interface {
	QWebEngineScript_PTR() *QWebEngineScript
}

func (ptr *QWebEngineScript) QWebEngineScript_PTR() *QWebEngineScript {
	return ptr
}

func (ptr *QWebEngineScript) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineScript) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineScript(ptr QWebEngineScript_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineScript_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineScriptFromPointer(ptr unsafe.Pointer) *QWebEngineScript {
	var n = new(QWebEngineScript)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebEngineScript__InjectionPoint
//QWebEngineScript::InjectionPoint
type QWebEngineScript__InjectionPoint int64

const (
	QWebEngineScript__Deferred         QWebEngineScript__InjectionPoint = QWebEngineScript__InjectionPoint(0)
	QWebEngineScript__DocumentReady    QWebEngineScript__InjectionPoint = QWebEngineScript__InjectionPoint(1)
	QWebEngineScript__DocumentCreation QWebEngineScript__InjectionPoint = QWebEngineScript__InjectionPoint(2)
)

//go:generate stringer -type=QWebEngineScript__ScriptWorldId
//QWebEngineScript::ScriptWorldId
type QWebEngineScript__ScriptWorldId int64

const (
	QWebEngineScript__MainWorld        QWebEngineScript__ScriptWorldId = QWebEngineScript__ScriptWorldId(0)
	QWebEngineScript__ApplicationWorld QWebEngineScript__ScriptWorldId = QWebEngineScript__ScriptWorldId(1)
	QWebEngineScript__UserWorld        QWebEngineScript__ScriptWorldId = QWebEngineScript__ScriptWorldId(2)
)

func NewQWebEngineScript() *QWebEngineScript {
	var tmpValue = NewQWebEngineScriptFromPointer(C.QWebEngineScript_NewQWebEngineScript())
	runtime.SetFinalizer(tmpValue, (*QWebEngineScript).DestroyQWebEngineScript)
	return tmpValue
}

func NewQWebEngineScript2(other QWebEngineScript_ITF) *QWebEngineScript {
	var tmpValue = NewQWebEngineScriptFromPointer(C.QWebEngineScript_NewQWebEngineScript2(PointerFromQWebEngineScript(other)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineScript).DestroyQWebEngineScript)
	return tmpValue
}

func (ptr *QWebEngineScript) InjectionPoint() QWebEngineScript__InjectionPoint {
	if ptr.Pointer() != nil {
		return QWebEngineScript__InjectionPoint(C.QWebEngineScript_InjectionPoint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineScript) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineScript_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineScript) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineScript_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineScript) RunsOnSubFrames() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineScript_RunsOnSubFrames(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineScript) SetInjectionPoint(p QWebEngineScript__InjectionPoint) {
	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetInjectionPoint(ptr.Pointer(), C.longlong(p))
	}
}

func (ptr *QWebEngineScript) SetName(scriptName string) {
	if ptr.Pointer() != nil {
		var scriptNameC = C.CString(scriptName)
		defer C.free(unsafe.Pointer(scriptNameC))
		C.QWebEngineScript_SetName(ptr.Pointer(), scriptNameC)
	}
}

func (ptr *QWebEngineScript) SetRunsOnSubFrames(on bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetRunsOnSubFrames(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QWebEngineScript) SetSourceCode(scriptSource string) {
	if ptr.Pointer() != nil {
		var scriptSourceC = C.CString(scriptSource)
		defer C.free(unsafe.Pointer(scriptSourceC))
		C.QWebEngineScript_SetSourceCode(ptr.Pointer(), scriptSourceC)
	}
}

func (ptr *QWebEngineScript) SetWorldId(id uint) {
	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetWorldId(ptr.Pointer(), C.uint(uint32(id)))
	}
}

func (ptr *QWebEngineScript) SourceCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineScript_SourceCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineScript) Swap(other QWebEngineScript_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineScript_Swap(ptr.Pointer(), PointerFromQWebEngineScript(other))
	}
}

func (ptr *QWebEngineScript) WorldId() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QWebEngineScript_WorldId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineScript) DestroyQWebEngineScript() {
	if ptr.Pointer() != nil {
		C.QWebEngineScript_DestroyQWebEngineScript(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWebEngineScriptCollection struct {
	ptr unsafe.Pointer
}

type QWebEngineScriptCollection_ITF interface {
	QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection
}

func (ptr *QWebEngineScriptCollection) QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection {
	return ptr
}

func (ptr *QWebEngineScriptCollection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineScriptCollection(ptr QWebEngineScriptCollection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineScriptCollection_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineScriptCollectionFromPointer(ptr unsafe.Pointer) *QWebEngineScriptCollection {
	var n = new(QWebEngineScriptCollection)
	n.SetPointer(ptr)
	return n
}
func (ptr *QWebEngineScriptCollection) Clear() {
	if ptr.Pointer() != nil {
		C.QWebEngineScriptCollection_Clear(ptr.Pointer())
	}
}

func (ptr *QWebEngineScriptCollection) Contains(value QWebEngineScript_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineScriptCollection_Contains(ptr.Pointer(), PointerFromQWebEngineScript(value)) != 0
	}
	return false
}

func (ptr *QWebEngineScriptCollection) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineScriptCollection_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineScriptCollection) FindScript(name string) *QWebEngineScript {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = NewQWebEngineScriptFromPointer(C.QWebEngineScriptCollection_FindScript(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*QWebEngineScript).DestroyQWebEngineScript)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) FindScripts(name string) []*QWebEngineScript {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return func(l C.struct_QtWebEngine_PackedList) []*QWebEngineScript {
			var out = make([]*QWebEngineScript, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebEngineScriptCollectionFromPointer(l.data).findScripts_atList(i)
			}
			return out
		}(C.QWebEngineScriptCollection_FindScripts(ptr.Pointer(), nameC))
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) Insert(s QWebEngineScript_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineScriptCollection_Insert(ptr.Pointer(), PointerFromQWebEngineScript(s))
	}
}

func (ptr *QWebEngineScriptCollection) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineScriptCollection_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineScriptCollection) Remove(script QWebEngineScript_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineScriptCollection_Remove(ptr.Pointer(), PointerFromQWebEngineScript(script)) != 0
	}
	return false
}

func (ptr *QWebEngineScriptCollection) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineScriptCollection_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineScriptCollection) ToList() []*QWebEngineScript {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*QWebEngineScript {
			var out = make([]*QWebEngineScript, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebEngineScriptCollectionFromPointer(l.data).toList_atList(i)
			}
			return out
		}(C.QWebEngineScriptCollection_ToList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) DestroyQWebEngineScriptCollection() {
	if ptr.Pointer() != nil {
		C.QWebEngineScriptCollection_DestroyQWebEngineScriptCollection(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineScriptCollection) findScripts_atList(i int) *QWebEngineScript {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineScriptFromPointer(C.QWebEngineScriptCollection_findScripts_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebEngineScript).DestroyQWebEngineScript)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) toList_atList(i int) *QWebEngineScript {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineScriptFromPointer(C.QWebEngineScriptCollection_toList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebEngineScript).DestroyQWebEngineScript)
		return tmpValue
	}
	return nil
}

type QWebEngineSettings struct {
	ptr unsafe.Pointer
}

type QWebEngineSettings_ITF interface {
	QWebEngineSettings_PTR() *QWebEngineSettings
}

func (ptr *QWebEngineSettings) QWebEngineSettings_PTR() *QWebEngineSettings {
	return ptr
}

func (ptr *QWebEngineSettings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineSettings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineSettings(ptr QWebEngineSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineSettings_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineSettingsFromPointer(ptr unsafe.Pointer) *QWebEngineSettings {
	var n = new(QWebEngineSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineSettings) DestroyQWebEngineSettings() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebEngineSettings__FontFamily
//QWebEngineSettings::FontFamily
type QWebEngineSettings__FontFamily int64

const (
	QWebEngineSettings__StandardFont   QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(0)
	QWebEngineSettings__FixedFont      QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(1)
	QWebEngineSettings__SerifFont      QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(2)
	QWebEngineSettings__SansSerifFont  QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(3)
	QWebEngineSettings__CursiveFont    QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(4)
	QWebEngineSettings__FantasyFont    QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(5)
	QWebEngineSettings__PictographFont QWebEngineSettings__FontFamily = QWebEngineSettings__FontFamily(6)
)

//go:generate stringer -type=QWebEngineSettings__FontSize
//QWebEngineSettings::FontSize
type QWebEngineSettings__FontSize int64

const (
	QWebEngineSettings__MinimumFontSize        QWebEngineSettings__FontSize = QWebEngineSettings__FontSize(0)
	QWebEngineSettings__MinimumLogicalFontSize QWebEngineSettings__FontSize = QWebEngineSettings__FontSize(1)
	QWebEngineSettings__DefaultFontSize        QWebEngineSettings__FontSize = QWebEngineSettings__FontSize(2)
	QWebEngineSettings__DefaultFixedFontSize   QWebEngineSettings__FontSize = QWebEngineSettings__FontSize(3)
)

//go:generate stringer -type=QWebEngineSettings__WebAttribute
//QWebEngineSettings::WebAttribute
type QWebEngineSettings__WebAttribute int64

const (
	QWebEngineSettings__AutoLoadImages                  QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(0)
	QWebEngineSettings__JavascriptEnabled               QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(1)
	QWebEngineSettings__JavascriptCanOpenWindows        QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(2)
	QWebEngineSettings__JavascriptCanAccessClipboard    QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(3)
	QWebEngineSettings__LinksIncludedInFocusChain       QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(4)
	QWebEngineSettings__LocalStorageEnabled             QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(5)
	QWebEngineSettings__LocalContentCanAccessRemoteUrls QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(6)
	QWebEngineSettings__XSSAuditingEnabled              QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(7)
	QWebEngineSettings__SpatialNavigationEnabled        QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(8)
	QWebEngineSettings__LocalContentCanAccessFileUrls   QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(9)
	QWebEngineSettings__HyperlinkAuditingEnabled        QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(10)
	QWebEngineSettings__ScrollAnimatorEnabled           QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(11)
	QWebEngineSettings__ErrorPageEnabled                QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(12)
	QWebEngineSettings__PluginsEnabled                  QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(13)
	QWebEngineSettings__FullScreenSupportEnabled        QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(14)
	QWebEngineSettings__ScreenCaptureEnabled            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(15)
	QWebEngineSettings__WebGLEnabled                    QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(16)
	QWebEngineSettings__Accelerated2dCanvasEnabled      QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(17)
	QWebEngineSettings__AutoLoadIconsForPage            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(18)
	QWebEngineSettings__TouchIconsEnabled               QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(19)
)

func (ptr *QWebEngineSettings) DefaultTextEncoding() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineSettings_DefaultTextEncoding(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineSettings) FontFamily(which QWebEngineSettings__FontFamily) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineSettings_FontFamily(ptr.Pointer(), C.longlong(which)))
	}
	return ""
}

func (ptr *QWebEngineSettings) FontSize(ty QWebEngineSettings__FontSize) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineSettings_FontSize(ptr.Pointer(), C.longlong(ty))))
	}
	return 0
}

func QWebEngineSettings_GlobalSettings() *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_GlobalSettings())
}

func (ptr *QWebEngineSettings) GlobalSettings() *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_GlobalSettings())
}

func (ptr *QWebEngineSettings) ResetAttribute(attribute QWebEngineSettings__WebAttribute) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetAttribute(ptr.Pointer(), C.longlong(attribute))
	}
}

func (ptr *QWebEngineSettings) ResetFontFamily(which QWebEngineSettings__FontFamily) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetFontFamily(ptr.Pointer(), C.longlong(which))
	}
}

func (ptr *QWebEngineSettings) ResetFontSize(ty QWebEngineSettings__FontSize) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetFontSize(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QWebEngineSettings) SetAttribute(attribute QWebEngineSettings__WebAttribute, on bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetAttribute(ptr.Pointer(), C.longlong(attribute), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QWebEngineSettings) SetDefaultTextEncoding(encoding string) {
	if ptr.Pointer() != nil {
		var encodingC = C.CString(encoding)
		defer C.free(unsafe.Pointer(encodingC))
		C.QWebEngineSettings_SetDefaultTextEncoding(ptr.Pointer(), encodingC)
	}
}

func (ptr *QWebEngineSettings) SetFontFamily(which QWebEngineSettings__FontFamily, family string) {
	if ptr.Pointer() != nil {
		var familyC = C.CString(family)
		defer C.free(unsafe.Pointer(familyC))
		C.QWebEngineSettings_SetFontFamily(ptr.Pointer(), C.longlong(which), familyC)
	}
}

func (ptr *QWebEngineSettings) SetFontSize(ty QWebEngineSettings__FontSize, size int) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetFontSize(ptr.Pointer(), C.longlong(ty), C.int(int32(size)))
	}
}

func (ptr *QWebEngineSettings) TestAttribute(attribute QWebEngineSettings__WebAttribute) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineSettings_TestAttribute(ptr.Pointer(), C.longlong(attribute)) != 0
	}
	return false
}

func QWebEngineSettings_DefaultSettings() *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_DefaultSettings())
}

func (ptr *QWebEngineSettings) DefaultSettings() *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_DefaultSettings())
}

type QWebEngineUrlRequestInfo struct {
	ptr unsafe.Pointer
}

type QWebEngineUrlRequestInfo_ITF interface {
	QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo
}

func (ptr *QWebEngineUrlRequestInfo) QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo {
	return ptr
}

func (ptr *QWebEngineUrlRequestInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineUrlRequestInfo(ptr QWebEngineUrlRequestInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlRequestInfo_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineUrlRequestInfoFromPointer(ptr unsafe.Pointer) *QWebEngineUrlRequestInfo {
	var n = new(QWebEngineUrlRequestInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineUrlRequestInfo) DestroyQWebEngineUrlRequestInfo() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebEngineUrlRequestInfo__NavigationType
//QWebEngineUrlRequestInfo::NavigationType
type QWebEngineUrlRequestInfo__NavigationType int64

const (
	QWebEngineUrlRequestInfo__NavigationTypeLink          QWebEngineUrlRequestInfo__NavigationType = QWebEngineUrlRequestInfo__NavigationType(0)
	QWebEngineUrlRequestInfo__NavigationTypeTyped         QWebEngineUrlRequestInfo__NavigationType = QWebEngineUrlRequestInfo__NavigationType(1)
	QWebEngineUrlRequestInfo__NavigationTypeFormSubmitted QWebEngineUrlRequestInfo__NavigationType = QWebEngineUrlRequestInfo__NavigationType(2)
	QWebEngineUrlRequestInfo__NavigationTypeBackForward   QWebEngineUrlRequestInfo__NavigationType = QWebEngineUrlRequestInfo__NavigationType(3)
	QWebEngineUrlRequestInfo__NavigationTypeReload        QWebEngineUrlRequestInfo__NavigationType = QWebEngineUrlRequestInfo__NavigationType(4)
	QWebEngineUrlRequestInfo__NavigationTypeOther         QWebEngineUrlRequestInfo__NavigationType = QWebEngineUrlRequestInfo__NavigationType(5)
)

//go:generate stringer -type=QWebEngineUrlRequestInfo__ResourceType
//QWebEngineUrlRequestInfo::ResourceType
type QWebEngineUrlRequestInfo__ResourceType int64

const (
	QWebEngineUrlRequestInfo__ResourceTypeMainFrame      QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(0)
	QWebEngineUrlRequestInfo__ResourceTypeSubFrame       QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(1)
	QWebEngineUrlRequestInfo__ResourceTypeStylesheet     QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(2)
	QWebEngineUrlRequestInfo__ResourceTypeScript         QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(3)
	QWebEngineUrlRequestInfo__ResourceTypeImage          QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(4)
	QWebEngineUrlRequestInfo__ResourceTypeFontResource   QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(5)
	QWebEngineUrlRequestInfo__ResourceTypeSubResource    QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(6)
	QWebEngineUrlRequestInfo__ResourceTypeObject         QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(7)
	QWebEngineUrlRequestInfo__ResourceTypeMedia          QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(8)
	QWebEngineUrlRequestInfo__ResourceTypeWorker         QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(9)
	QWebEngineUrlRequestInfo__ResourceTypeSharedWorker   QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(10)
	QWebEngineUrlRequestInfo__ResourceTypePrefetch       QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(11)
	QWebEngineUrlRequestInfo__ResourceTypeFavicon        QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(12)
	QWebEngineUrlRequestInfo__ResourceTypeXhr            QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(13)
	QWebEngineUrlRequestInfo__ResourceTypePing           QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(14)
	QWebEngineUrlRequestInfo__ResourceTypeServiceWorker  QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(15)
	QWebEngineUrlRequestInfo__ResourceTypeCspReport      QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(16)
	QWebEngineUrlRequestInfo__ResourceTypePluginResource QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(17)
	QWebEngineUrlRequestInfo__ResourceTypeUnknown        QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(18)
)

func (ptr *QWebEngineUrlRequestInfo) Block(shouldBlock bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_Block(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(shouldBlock))))
	}
}

func (ptr *QWebEngineUrlRequestInfo) FirstPartyUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_FirstPartyUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) NavigationType() QWebEngineUrlRequestInfo__NavigationType {
	if ptr.Pointer() != nil {
		return QWebEngineUrlRequestInfo__NavigationType(C.QWebEngineUrlRequestInfo_NavigationType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineUrlRequestInfo) Redirect(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_Redirect(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineUrlRequestInfo) RequestMethod() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestInfo_RequestMethod(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_RequestUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) ResourceType() QWebEngineUrlRequestInfo__ResourceType {
	if ptr.Pointer() != nil {
		return QWebEngineUrlRequestInfo__ResourceType(C.QWebEngineUrlRequestInfo_ResourceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineUrlRequestInfo) SetHttpHeader(name core.QByteArray_ITF, value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_SetHttpHeader(ptr.Pointer(), core.PointerFromQByteArray(name), core.PointerFromQByteArray(value))
	}
}

type QWebEngineUrlRequestInterceptor struct {
	core.QObject
}

type QWebEngineUrlRequestInterceptor_ITF interface {
	core.QObject_ITF
	QWebEngineUrlRequestInterceptor_PTR() *QWebEngineUrlRequestInterceptor
}

func (ptr *QWebEngineUrlRequestInterceptor) QWebEngineUrlRequestInterceptor_PTR() *QWebEngineUrlRequestInterceptor {
	return ptr
}

func (ptr *QWebEngineUrlRequestInterceptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineUrlRequestInterceptor(ptr QWebEngineUrlRequestInterceptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlRequestInterceptor_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineUrlRequestInterceptorFromPointer(ptr unsafe.Pointer) *QWebEngineUrlRequestInterceptor {
	var n = new(QWebEngineUrlRequestInterceptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineUrlRequestInterceptor) DestroyQWebEngineUrlRequestInterceptor() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQWebEngineUrlRequestInterceptor(p core.QObject_ITF) *QWebEngineUrlRequestInterceptor {
	var tmpValue = NewQWebEngineUrlRequestInterceptorFromPointer(C.QWebEngineUrlRequestInterceptor_NewQWebEngineUrlRequestInterceptor(core.PointerFromQObject(p)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineUrlRequestInterceptor_InterceptRequest
func callbackQWebEngineUrlRequestInterceptor_InterceptRequest(ptr unsafe.Pointer, info unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::interceptRequest"); signal != nil {
		signal.(func(*QWebEngineUrlRequestInfo))(NewQWebEngineUrlRequestInfoFromPointer(info))
	}

}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectInterceptRequest(f func(info *QWebEngineUrlRequestInfo)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::interceptRequest", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectInterceptRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::interceptRequest")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) InterceptRequest(info QWebEngineUrlRequestInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_InterceptRequest(ptr.Pointer(), PointerFromQWebEngineUrlRequestInfo(info))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_TimerEvent
func callbackQWebEngineUrlRequestInterceptor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::timerEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::timerEvent")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_ChildEvent
func callbackQWebEngineUrlRequestInterceptor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::childEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::childEvent")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_ConnectNotify
func callbackQWebEngineUrlRequestInterceptor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::connectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::connectNotify")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_CustomEvent
func callbackQWebEngineUrlRequestInterceptor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::customEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::customEvent")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_DeleteLater
func callbackQWebEngineUrlRequestInterceptor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::deleteLater", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::deleteLater")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlRequestInterceptor_DisconnectNotify
func callbackQWebEngineUrlRequestInterceptor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::disconnectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::disconnectNotify")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_Event
func callbackQWebEngineUrlRequestInterceptor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::event", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::event")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestInterceptor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestInterceptor_EventFilter
func callbackQWebEngineUrlRequestInterceptor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::eventFilter", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::eventFilter")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestInterceptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestInterceptor_MetaObject
func callbackQWebEngineUrlRequestInterceptor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestInterceptor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::metaObject", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestInterceptor::metaObject")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestInterceptor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestInterceptor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineUrlRequestJob struct {
	core.QObject
}

type QWebEngineUrlRequestJob_ITF interface {
	core.QObject_ITF
	QWebEngineUrlRequestJob_PTR() *QWebEngineUrlRequestJob
}

func (ptr *QWebEngineUrlRequestJob) QWebEngineUrlRequestJob_PTR() *QWebEngineUrlRequestJob {
	return ptr
}

func (ptr *QWebEngineUrlRequestJob) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineUrlRequestJob(ptr QWebEngineUrlRequestJob_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlRequestJob_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineUrlRequestJobFromPointer(ptr unsafe.Pointer) *QWebEngineUrlRequestJob {
	var n = new(QWebEngineUrlRequestJob)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebEngineUrlRequestJob) DestroyQWebEngineUrlRequestJob() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebEngineUrlRequestJob__Error
//QWebEngineUrlRequestJob::Error
type QWebEngineUrlRequestJob__Error int64

const (
	QWebEngineUrlRequestJob__NoError        QWebEngineUrlRequestJob__Error = QWebEngineUrlRequestJob__Error(0)
	QWebEngineUrlRequestJob__UrlNotFound    QWebEngineUrlRequestJob__Error = QWebEngineUrlRequestJob__Error(1)
	QWebEngineUrlRequestJob__UrlInvalid     QWebEngineUrlRequestJob__Error = QWebEngineUrlRequestJob__Error(2)
	QWebEngineUrlRequestJob__RequestAborted QWebEngineUrlRequestJob__Error = QWebEngineUrlRequestJob__Error(3)
	QWebEngineUrlRequestJob__RequestDenied  QWebEngineUrlRequestJob__Error = QWebEngineUrlRequestJob__Error(4)
	QWebEngineUrlRequestJob__RequestFailed  QWebEngineUrlRequestJob__Error = QWebEngineUrlRequestJob__Error(5)
)

func (ptr *QWebEngineUrlRequestJob) Fail(r QWebEngineUrlRequestJob__Error) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_Fail(ptr.Pointer(), C.longlong(r))
	}
}

func (ptr *QWebEngineUrlRequestJob) Redirect(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_Redirect(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineUrlRequestJob) Reply(contentType core.QByteArray_ITF, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_Reply(ptr.Pointer(), core.PointerFromQByteArray(contentType), core.PointerFromQIODevice(device))
	}
}

func (ptr *QWebEngineUrlRequestJob) RequestMethod() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestJob_RequestMethod(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineUrlRequestJob_RequestUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineUrlRequestJob_TimerEvent
func callbackQWebEngineUrlRequestJob_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::timerEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::timerEvent")
	}
}

func (ptr *QWebEngineUrlRequestJob) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_ChildEvent
func callbackQWebEngineUrlRequestJob_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::childEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::childEvent")
	}
}

func (ptr *QWebEngineUrlRequestJob) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_ConnectNotify
func callbackQWebEngineUrlRequestJob_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::connectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::connectNotify")
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestJob_CustomEvent
func callbackQWebEngineUrlRequestJob_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::customEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::customEvent")
	}
}

func (ptr *QWebEngineUrlRequestJob) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_DeleteLater
func callbackQWebEngineUrlRequestJob_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::deleteLater", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::deleteLater")
	}
}

func (ptr *QWebEngineUrlRequestJob) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlRequestJob) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlRequestJob_DisconnectNotify
func callbackQWebEngineUrlRequestJob_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::disconnectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::disconnectNotify")
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestJob_Event
func callbackQWebEngineUrlRequestJob_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestJobFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineUrlRequestJob) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::event", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::event")
	}
}

func (ptr *QWebEngineUrlRequestJob) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestJob) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestJob_EventFilter
func callbackQWebEngineUrlRequestJob_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestJobFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineUrlRequestJob) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::eventFilter", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::eventFilter")
	}
}

func (ptr *QWebEngineUrlRequestJob) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestJob) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestJob_MetaObject
func callbackQWebEngineUrlRequestJob_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlRequestJob::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlRequestJobFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlRequestJob) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::metaObject", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlRequestJob::metaObject")
	}
}

func (ptr *QWebEngineUrlRequestJob) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestJob_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestJob_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineUrlSchemeHandler struct {
	core.QObject
}

type QWebEngineUrlSchemeHandler_ITF interface {
	core.QObject_ITF
	QWebEngineUrlSchemeHandler_PTR() *QWebEngineUrlSchemeHandler
}

func (ptr *QWebEngineUrlSchemeHandler) QWebEngineUrlSchemeHandler_PTR() *QWebEngineUrlSchemeHandler {
	return ptr
}

func (ptr *QWebEngineUrlSchemeHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineUrlSchemeHandler(ptr QWebEngineUrlSchemeHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlSchemeHandler_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineUrlSchemeHandlerFromPointer(ptr unsafe.Pointer) *QWebEngineUrlSchemeHandler {
	var n = new(QWebEngineUrlSchemeHandler)
	n.SetPointer(ptr)
	return n
}
func NewQWebEngineUrlSchemeHandler(parent core.QObject_ITF) *QWebEngineUrlSchemeHandler {
	var tmpValue = NewQWebEngineUrlSchemeHandlerFromPointer(C.QWebEngineUrlSchemeHandler_NewQWebEngineUrlSchemeHandler(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineUrlSchemeHandler_RequestStarted
func callbackQWebEngineUrlSchemeHandler_RequestStarted(ptr unsafe.Pointer, request unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::requestStarted"); signal != nil {
		signal.(func(*QWebEngineUrlRequestJob))(NewQWebEngineUrlRequestJobFromPointer(request))
	}

}

func (ptr *QWebEngineUrlSchemeHandler) ConnectRequestStarted(f func(request *QWebEngineUrlRequestJob)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::requestStarted", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectRequestStarted() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::requestStarted")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) RequestStarted(request QWebEngineUrlRequestJob_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_RequestStarted(ptr.Pointer(), PointerFromQWebEngineUrlRequestJob(request))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandler() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlSchemeHandler_TimerEvent
func callbackQWebEngineUrlSchemeHandler_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::timerEvent", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::timerEvent")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_ChildEvent
func callbackQWebEngineUrlSchemeHandler_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::childEvent", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::childEvent")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_ConnectNotify
func callbackQWebEngineUrlSchemeHandler_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::connectNotify", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::connectNotify")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlSchemeHandler_CustomEvent
func callbackQWebEngineUrlSchemeHandler_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::customEvent", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::customEvent")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_DeleteLater
func callbackQWebEngineUrlSchemeHandler_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::deleteLater", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::deleteLater")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlSchemeHandler_DisconnectNotify
func callbackQWebEngineUrlSchemeHandler_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::disconnectNotify", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::disconnectNotify")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlSchemeHandler_Event
func callbackQWebEngineUrlSchemeHandler_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::event", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::event")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlSchemeHandler) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineUrlSchemeHandler_EventFilter
func callbackQWebEngineUrlSchemeHandler_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::eventFilter", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::eventFilter")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlSchemeHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineUrlSchemeHandler_MetaObject
func callbackQWebEngineUrlSchemeHandler_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineUrlSchemeHandler::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::metaObject", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineUrlSchemeHandler::metaObject")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlSchemeHandler_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlSchemeHandler_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineView struct {
	widgets.QWidget
}

type QWebEngineView_ITF interface {
	widgets.QWidget_ITF
	QWebEngineView_PTR() *QWebEngineView
}

func (ptr *QWebEngineView) QWebEngineView_PTR() *QWebEngineView {
	return ptr
}

func (ptr *QWebEngineView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineView(ptr QWebEngineView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineView_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineViewFromPointer(ptr unsafe.Pointer) *QWebEngineView {
	var n = new(QWebEngineView)
	n.SetPointer(ptr)
	return n
}
func NewQWebEngineView(parent widgets.QWidget_ITF) *QWebEngineView {
	var tmpValue = NewQWebEngineViewFromPointer(C.QWebEngineView_NewQWebEngineView(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineView_Back
func callbackQWebEngineView_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::back"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).BackDefault()
	}
}

func (ptr *QWebEngineView) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::back", f)
	}
}

func (ptr *QWebEngineView) DisconnectBack() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::back")
	}
}

func (ptr *QWebEngineView) Back() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Back(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) BackDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_BackDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_CreateWindow
func callbackQWebEngineView_CreateWindow(ptr unsafe.Pointer, ty C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::createWindow"); signal != nil {
		return PointerFromQWebEngineView(signal.(func(QWebEnginePage__WebWindowType) *QWebEngineView)(QWebEnginePage__WebWindowType(ty)))
	}

	return PointerFromQWebEngineView(NewQWebEngineViewFromPointer(ptr).CreateWindowDefault(QWebEnginePage__WebWindowType(ty)))
}

func (ptr *QWebEngineView) ConnectCreateWindow(f func(ty QWebEnginePage__WebWindowType) *QWebEngineView) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::createWindow", f)
	}
}

func (ptr *QWebEngineView) DisconnectCreateWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::createWindow")
	}
}

func (ptr *QWebEngineView) CreateWindow(ty QWebEnginePage__WebWindowType) *QWebEngineView {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineViewFromPointer(C.QWebEngineView_CreateWindow(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) CreateWindowDefault(ty QWebEnginePage__WebWindowType) *QWebEngineView {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEngineViewFromPointer(C.QWebEngineView_CreateWindowDefault(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_Forward
func callbackQWebEngineView_Forward(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::forward"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ForwardDefault()
	}
}

func (ptr *QWebEngineView) ConnectForward(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::forward", f)
	}
}

func (ptr *QWebEngineView) DisconnectForward() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::forward")
	}
}

func (ptr *QWebEngineView) Forward() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Forward(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ForwardDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ForwardDefault(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineView) History() *QWebEngineHistory {
	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryFromPointer(C.QWebEngineView_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QWebEngineView_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) IconUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineView_IconUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineView) Page() *QWebEnginePage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebEnginePageFromPointer(C.QWebEngineView_Page(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) PageAction(action QWebEnginePage__WebAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebEngineView_PageAction(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_Reload
func callbackQWebEngineView_Reload(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::reload"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ReloadDefault()
	}
}

func (ptr *QWebEngineView) ConnectReload(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::reload", f)
	}
}

func (ptr *QWebEngineView) DisconnectReload() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::reload")
	}
}

func (ptr *QWebEngineView) Reload() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Reload(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ReloadDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ReloadDefault(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineView_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineView) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var mimeTypeC = C.CString(mimeType)
		defer C.free(unsafe.Pointer(mimeTypeC))
		C.QWebEngineView_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), mimeTypeC, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEngineView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC = C.CString(html)
		defer C.free(unsafe.Pointer(htmlC))
		C.QWebEngineView_SetHtml(ptr.Pointer(), htmlC, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEngineView) SetPage(page QWebEnginePage_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetPage(ptr.Pointer(), PointerFromQWebEnginePage(page))
	}
}

func (ptr *QWebEngineView) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineView) SetZoomFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebEngineView) Settings() *QWebEngineSettings {
	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEngineView_Settings(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineView_Stop
func callbackQWebEngineView_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).StopDefault()
	}
}

func (ptr *QWebEngineView) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::stop", f)
	}
}

func (ptr *QWebEngineView) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::stop")
	}
}

func (ptr *QWebEngineView) Stop() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Stop(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) StopDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_StopDefault(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineView) TriggerPageAction(action QWebEnginePage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TriggerPageAction(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

func (ptr *QWebEngineView) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebEngineView_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebEngineView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebEngineView_DestroyQWebEngineView
func callbackQWebEngineView_DestroyQWebEngineView(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::~QWebEngineView"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).DestroyQWebEngineViewDefault()
	}
}

func (ptr *QWebEngineView) ConnectDestroyQWebEngineView(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::~QWebEngineView", f)
	}
}

func (ptr *QWebEngineView) DisconnectDestroyQWebEngineView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::~QWebEngineView")
	}
}

func (ptr *QWebEngineView) DestroyQWebEngineView() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DestroyQWebEngineView(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineView) DestroyQWebEngineViewDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DestroyQWebEngineViewDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineView_ContextMenuEvent
func callbackQWebEngineView_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::contextMenuEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::contextMenuEvent")
	}
}

func (ptr *QWebEngineView) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWebEngineView) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQWebEngineView_DragEnterEvent
func callbackQWebEngineView_DragEnterEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dragEnterEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dragEnterEvent")
	}
}

func (ptr *QWebEngineView) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QWebEngineView) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

//export callbackQWebEngineView_DragLeaveEvent
func callbackQWebEngineView_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dragLeaveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dragLeaveEvent")
	}
}

func (ptr *QWebEngineView) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QWebEngineView) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQWebEngineView_DragMoveEvent
func callbackQWebEngineView_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dragMoveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dragMoveEvent")
	}
}

func (ptr *QWebEngineView) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QWebEngineView) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQWebEngineView_DropEvent
func callbackQWebEngineView_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dropEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::dropEvent")
	}
}

func (ptr *QWebEngineView) DropEvent(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QWebEngineView) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQWebEngineView_Event
func callbackQWebEngineView_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QWebEngineView) ConnectEvent(f func(ev *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::event", f)
	}
}

func (ptr *QWebEngineView) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::event")
	}
}

func (ptr *QWebEngineView) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebEngineView) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebEngineView) FindText(subString string, options QWebEnginePage__FindFlag) {
	if ptr.Pointer() != nil {
		var subStringC = C.CString(subString)
		defer C.free(unsafe.Pointer(subStringC))
		C.QWebEngineView_FindText(ptr.Pointer(), subStringC, C.longlong(options))
	}
}

//export callbackQWebEngineView_HideEvent
func callbackQWebEngineView_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::hideEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::hideEvent")
	}
}

func (ptr *QWebEngineView) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWebEngineView) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWebEngineView_IconChanged
func callbackQWebEngineView_IconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::iconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

func (ptr *QWebEngineView) ConnectIconChanged(f func(icon *gui.QIcon)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectIconChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::iconChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::iconChanged")
	}
}

func (ptr *QWebEngineView) IconChanged(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_IconChanged(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

//export callbackQWebEngineView_IconUrlChanged
func callbackQWebEngineView_IconUrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::iconUrlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEngineView) ConnectIconUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectIconUrlChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::iconUrlChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectIconUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectIconUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::iconUrlChanged")
	}
}

func (ptr *QWebEngineView) IconUrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_IconUrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebEngineView_LoadFinished
func callbackQWebEngineView_LoadFinished(ptr unsafe.Pointer, ok C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::loadFinished"); signal != nil {
		signal.(func(bool))(int8(ok) != 0)
	}

}

func (ptr *QWebEngineView) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::loadFinished", f)
	}
}

func (ptr *QWebEngineView) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::loadFinished")
	}
}

func (ptr *QWebEngineView) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQWebEngineView_LoadProgress
func callbackQWebEngineView_LoadProgress(ptr unsafe.Pointer, progress C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::loadProgress"); signal != nil {
		signal.(func(int))(int(int32(progress)))
	}

}

func (ptr *QWebEngineView) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::loadProgress", f)
	}
}

func (ptr *QWebEngineView) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::loadProgress")
	}
}

func (ptr *QWebEngineView) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQWebEngineView_LoadStarted
func callbackQWebEngineView_LoadStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::loadStarted", f)
	}
}

func (ptr *QWebEngineView) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::loadStarted")
	}
}

func (ptr *QWebEngineView) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebEngineView_RenderProcessTerminated
func callbackQWebEngineView_RenderProcessTerminated(ptr unsafe.Pointer, terminationStatus C.longlong, exitCode C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::renderProcessTerminated"); signal != nil {
		signal.(func(QWebEnginePage__RenderProcessTerminationStatus, int))(QWebEnginePage__RenderProcessTerminationStatus(terminationStatus), int(int32(exitCode)))
	}

}

func (ptr *QWebEngineView) ConnectRenderProcessTerminated(f func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectRenderProcessTerminated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::renderProcessTerminated", f)
	}
}

func (ptr *QWebEngineView) DisconnectRenderProcessTerminated() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectRenderProcessTerminated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::renderProcessTerminated")
	}
}

func (ptr *QWebEngineView) RenderProcessTerminated(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_RenderProcessTerminated(ptr.Pointer(), C.longlong(terminationStatus), C.int(int32(exitCode)))
	}
}

//export callbackQWebEngineView_SelectionChanged
func callbackQWebEngineView_SelectionChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::selectionChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::selectionChanged")
	}
}

func (ptr *QWebEngineView) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SelectionChanged(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowEvent
func callbackQWebEngineView_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showEvent")
	}
}

func (ptr *QWebEngineView) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWebEngineView) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWebEngineView_SizeHint
func callbackQWebEngineView_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebEngineViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebEngineView) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::sizeHint", f)
	}
}

func (ptr *QWebEngineView) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::sizeHint")
	}
}

func (ptr *QWebEngineView) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebEngineView_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebEngineView_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_TitleChanged
func callbackQWebEngineView_TitleChanged(ptr unsafe.Pointer, title C.struct_QtWebEngine_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::titleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

func (ptr *QWebEngineView) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::titleChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::titleChanged")
	}
}

func (ptr *QWebEngineView) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
		C.QWebEngineView_TitleChanged(ptr.Pointer(), titleC)
	}
}

//export callbackQWebEngineView_UrlChanged
func callbackQWebEngineView_UrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEngineView) ConnectUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::urlChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::urlChanged")
	}
}

func (ptr *QWebEngineView) UrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebEngineView_ActionEvent
func callbackQWebEngineView_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::actionEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectActionEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::actionEvent")
	}
}

func (ptr *QWebEngineView) ActionEvent(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWebEngineView) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebEngineView_EnterEvent
func callbackQWebEngineView_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectEnterEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::enterEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::enterEvent")
	}
}

func (ptr *QWebEngineView) EnterEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_FocusInEvent
func callbackQWebEngineView_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::focusInEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::focusInEvent")
	}
}

func (ptr *QWebEngineView) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWebEngineView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebEngineView_FocusOutEvent
func callbackQWebEngineView_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::focusOutEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::focusOutEvent")
	}
}

func (ptr *QWebEngineView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWebEngineView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebEngineView_LeaveEvent
func callbackQWebEngineView_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::leaveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::leaveEvent")
	}
}

func (ptr *QWebEngineView) LeaveEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_Metric
func callbackQWebEngineView_Metric(ptr unsafe.Pointer, m C.longlong) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQWebEngineViewFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QWebEngineView) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::metric", f)
	}
}

func (ptr *QWebEngineView) DisconnectMetric() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::metric")
	}
}

func (ptr *QWebEngineView) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineView_Metric(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

func (ptr *QWebEngineView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineView_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQWebEngineView_MinimumSizeHint
func callbackQWebEngineView_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebEngineViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebEngineView) ConnectMinimumSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::minimumSizeHint", f)
	}
}

func (ptr *QWebEngineView) DisconnectMinimumSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::minimumSizeHint")
	}
}

func (ptr *QWebEngineView) MinimumSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebEngineView_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebEngineView_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_MoveEvent
func callbackQWebEngineView_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::moveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::moveEvent")
	}
}

func (ptr *QWebEngineView) MoveEvent(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWebEngineView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebEngineView_PaintEngine
func callbackQWebEngineView_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQWebEngineViewFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWebEngineView) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::paintEngine", f)
	}
}

func (ptr *QWebEngineView) DisconnectPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::paintEngine")
	}
}

func (ptr *QWebEngineView) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebEngineView_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebEngineView_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineView_PaintEvent
func callbackQWebEngineView_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::paintEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::paintEvent")
	}
}

func (ptr *QWebEngineView) PaintEvent(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWebEngineView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQWebEngineView_SetEnabled
func callbackQWebEngineView_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setEnabled", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setEnabled")
	}
}

func (ptr *QWebEngineView) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QWebEngineView) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebEngineView_SetStyleSheet
func callbackQWebEngineView_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQWebEngineViewFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QWebEngineView) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setStyleSheet", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setStyleSheet")
	}
}

func (ptr *QWebEngineView) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QWebEngineView_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QWebEngineView) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QWebEngineView_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQWebEngineView_SetVisible
func callbackQWebEngineView_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setVisible", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setVisible")
	}
}

func (ptr *QWebEngineView) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QWebEngineView) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebEngineView_SetWindowModified
func callbackQWebEngineView_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetWindowModified(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setWindowModified", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetWindowModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setWindowModified")
	}
}

func (ptr *QWebEngineView) SetWindowModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QWebEngineView) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebEngineView_SetWindowTitle
func callbackQWebEngineView_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQWebEngineViewFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QWebEngineView) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setWindowTitle", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setWindowTitle")
	}
}

func (ptr *QWebEngineView) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QWebEngineView_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QWebEngineView) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QWebEngineView_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQWebEngineView_ChangeEvent
func callbackQWebEngineView_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectChangeEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::changeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectChangeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::changeEvent")
	}
}

func (ptr *QWebEngineView) ChangeEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_Close
func callbackQWebEngineView_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).CloseDefault())))
}

func (ptr *QWebEngineView) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::close", f)
	}
}

func (ptr *QWebEngineView) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::close")
	}
}

func (ptr *QWebEngineView) Close() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebEngineView_CloseEvent
func callbackQWebEngineView_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::closeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::closeEvent")
	}
}

func (ptr *QWebEngineView) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWebEngineView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebEngineView_FocusNextPrevChild
func callbackQWebEngineView_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QWebEngineView) ConnectFocusNextPrevChild(f func(next bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::focusNextPrevChild", f)
	}
}

func (ptr *QWebEngineView) DisconnectFocusNextPrevChild() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::focusNextPrevChild")
	}
}

func (ptr *QWebEngineView) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QWebEngineView) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQWebEngineView_HasHeightForWidth
func callbackQWebEngineView_HasHeightForWidth(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QWebEngineView) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::hasHeightForWidth", f)
	}
}

func (ptr *QWebEngineView) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::hasHeightForWidth")
	}
}

func (ptr *QWebEngineView) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineView) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebEngineView_HeightForWidth
func callbackQWebEngineView_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQWebEngineViewFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QWebEngineView) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::heightForWidth", f)
	}
}

func (ptr *QWebEngineView) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::heightForWidth")
	}
}

func (ptr *QWebEngineView) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineView_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QWebEngineView) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineView_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQWebEngineView_Hide
func callbackQWebEngineView_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebEngineView) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::hide", f)
	}
}

func (ptr *QWebEngineView) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::hide")
	}
}

func (ptr *QWebEngineView) Hide() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Hide(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_InputMethodEvent
func callbackQWebEngineView_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::inputMethodEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::inputMethodEvent")
	}
}

func (ptr *QWebEngineView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWebEngineView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQWebEngineView_InputMethodQuery
func callbackQWebEngineView_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQWebEngineViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QWebEngineView) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::inputMethodQuery", f)
	}
}

func (ptr *QWebEngineView) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::inputMethodQuery")
	}
}

func (ptr *QWebEngineView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QWebEngineView_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QWebEngineView_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_KeyPressEvent
func callbackQWebEngineView_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::keyPressEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::keyPressEvent")
	}
}

func (ptr *QWebEngineView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWebEngineView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebEngineView_KeyReleaseEvent
func callbackQWebEngineView_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::keyReleaseEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::keyReleaseEvent")
	}
}

func (ptr *QWebEngineView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWebEngineView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebEngineView_Lower
func callbackQWebEngineView_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebEngineView) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::lower", f)
	}
}

func (ptr *QWebEngineView) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::lower")
	}
}

func (ptr *QWebEngineView) Lower() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Lower(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_MouseDoubleClickEvent
func callbackQWebEngineView_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mouseDoubleClickEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mouseDoubleClickEvent")
	}
}

func (ptr *QWebEngineView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MouseMoveEvent
func callbackQWebEngineView_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mouseMoveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mouseMoveEvent")
	}
}

func (ptr *QWebEngineView) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MousePressEvent
func callbackQWebEngineView_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mousePressEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mousePressEvent")
	}
}

func (ptr *QWebEngineView) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MouseReleaseEvent
func callbackQWebEngineView_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mouseReleaseEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::mouseReleaseEvent")
	}
}

func (ptr *QWebEngineView) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_NativeEvent
func callbackQWebEngineView_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QWebEngineView) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::nativeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::nativeEvent")
	}
}

func (ptr *QWebEngineView) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QWebEngineView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQWebEngineView_Raise
func callbackQWebEngineView_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebEngineView) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::raise", f)
	}
}

func (ptr *QWebEngineView) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::raise")
	}
}

func (ptr *QWebEngineView) Raise() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Raise(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_Repaint
func callbackQWebEngineView_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebEngineView) ConnectRepaint(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::repaint", f)
	}
}

func (ptr *QWebEngineView) DisconnectRepaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::repaint")
	}
}

func (ptr *QWebEngineView) Repaint() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Repaint(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ResizeEvent
func callbackQWebEngineView_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::resizeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::resizeEvent")
	}
}

func (ptr *QWebEngineView) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWebEngineView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQWebEngineView_SetDisabled
func callbackQWebEngineView_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setDisabled", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setDisabled")
	}
}

func (ptr *QWebEngineView) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QWebEngineView) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQWebEngineView_SetFocus2
func callbackQWebEngineView_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebEngineView) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setFocus2", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setFocus2")
	}
}

func (ptr *QWebEngineView) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebEngineView_SetHidden
func callbackQWebEngineView_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetHidden(f func(hidden bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setHidden", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::setHidden")
	}
}

func (ptr *QWebEngineView) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QWebEngineView) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQWebEngineView_Show
func callbackQWebEngineView_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::show"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebEngineView) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::show", f)
	}
}

func (ptr *QWebEngineView) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::show")
	}
}

func (ptr *QWebEngineView) Show() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Show(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowFullScreen
func callbackQWebEngineView_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showFullScreen", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showFullScreen")
	}
}

func (ptr *QWebEngineView) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowMaximized
func callbackQWebEngineView_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showMaximized", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showMaximized")
	}
}

func (ptr *QWebEngineView) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowMinimized
func callbackQWebEngineView_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showMinimized", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showMinimized")
	}
}

func (ptr *QWebEngineView) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowNormal
func callbackQWebEngineView_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showNormal", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::showNormal")
	}
}

func (ptr *QWebEngineView) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_TabletEvent
func callbackQWebEngineView_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::tabletEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::tabletEvent")
	}
}

func (ptr *QWebEngineView) TabletEvent(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWebEngineView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebEngineView_Update
func callbackQWebEngineView_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::update"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebEngineView) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::update", f)
	}
}

func (ptr *QWebEngineView) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::update")
	}
}

func (ptr *QWebEngineView) Update() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Update(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_UpdateMicroFocus
func callbackQWebEngineView_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebEngineView) ConnectUpdateMicroFocus(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::updateMicroFocus", f)
	}
}

func (ptr *QWebEngineView) DisconnectUpdateMicroFocus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::updateMicroFocus")
	}
}

func (ptr *QWebEngineView) UpdateMicroFocus() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_WheelEvent
func callbackQWebEngineView_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::wheelEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::wheelEvent")
	}
}

func (ptr *QWebEngineView) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWebEngineView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQWebEngineView_TimerEvent
func callbackQWebEngineView_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::timerEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::timerEvent")
	}
}

func (ptr *QWebEngineView) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineView) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineView_ChildEvent
func callbackQWebEngineView_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::childEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::childEvent")
	}
}

func (ptr *QWebEngineView) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineView) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineView_ConnectNotify
func callbackQWebEngineView_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineView) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::connectNotify", f)
	}
}

func (ptr *QWebEngineView) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::connectNotify")
	}
}

func (ptr *QWebEngineView) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineView_CustomEvent
func callbackQWebEngineView_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::customEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::customEvent")
	}
}

func (ptr *QWebEngineView) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_DeleteLater
func callbackQWebEngineView_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineView) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::deleteLater", f)
	}
}

func (ptr *QWebEngineView) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::deleteLater")
	}
}

func (ptr *QWebEngineView) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineView_DisconnectNotify
func callbackQWebEngineView_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineView) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::disconnectNotify", f)
	}
}

func (ptr *QWebEngineView) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::disconnectNotify")
	}
}

func (ptr *QWebEngineView) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineView_EventFilter
func callbackQWebEngineView_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineView) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::eventFilter", f)
	}
}

func (ptr *QWebEngineView) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::eventFilter")
	}
}

func (ptr *QWebEngineView) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebEngineView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineView_MetaObject
func callbackQWebEngineView_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebEngineView::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineView) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::metaObject", f)
	}
}

func (ptr *QWebEngineView) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebEngineView::metaObject")
	}
}

func (ptr *QWebEngineView) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineView_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QtWebEngineCore struct {
	ptr unsafe.Pointer
}

type QtWebEngineCore_ITF interface {
	QtWebEngineCore_PTR() *QtWebEngineCore
}

func (ptr *QtWebEngineCore) QtWebEngineCore_PTR() *QtWebEngineCore {
	return ptr
}

func (ptr *QtWebEngineCore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtWebEngineCore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtWebEngineCore(ptr QtWebEngineCore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWebEngineCore_PTR().Pointer()
	}
	return nil
}

func NewQtWebEngineCoreFromPointer(ptr unsafe.Pointer) *QtWebEngineCore {
	var n = new(QtWebEngineCore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QtWebEngineCore) DestroyQtWebEngineCore() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QtWebEngineCore__TextureTarget
//QtWebEngineCore::TextureTarget
type QtWebEngineCore__TextureTarget int64

const (
	QtWebEngineCore__ExternalTarget  QtWebEngineCore__TextureTarget = QtWebEngineCore__TextureTarget(0)
	QtWebEngineCore__RectangleTarget QtWebEngineCore__TextureTarget = QtWebEngineCore__TextureTarget(1)
)
