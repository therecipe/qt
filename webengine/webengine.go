package webengine

//#include "webengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"github.com/therecipe/qt/webchannel"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

//QQuickWebEngineProfile::HttpCacheType
type QQuickWebEngineProfile__HttpCacheType int64

const (
	QQuickWebEngineProfile__MemoryHttpCache = QQuickWebEngineProfile__HttpCacheType(0)
	QQuickWebEngineProfile__DiskHttpCache   = QQuickWebEngineProfile__HttpCacheType(1)
)

//QQuickWebEngineProfile::PersistentCookiesPolicy
type QQuickWebEngineProfile__PersistentCookiesPolicy int64

const (
	QQuickWebEngineProfile__NoPersistentCookies    = QQuickWebEngineProfile__PersistentCookiesPolicy(0)
	QQuickWebEngineProfile__AllowPersistentCookies = QQuickWebEngineProfile__PersistentCookiesPolicy(1)
	QQuickWebEngineProfile__ForcePersistentCookies = QQuickWebEngineProfile__PersistentCookiesPolicy(2)
)

type QQuickWebEngineProfile struct {
	core.QObject
}

type QQuickWebEngineProfile_ITF interface {
	core.QObject_ITF
	QQuickWebEngineProfile_PTR() *QQuickWebEngineProfile
}

func (p *QQuickWebEngineProfile) QQuickWebEngineProfile_PTR() *QQuickWebEngineProfile {
	return p
}

func (p *QQuickWebEngineProfile) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickWebEngineProfile) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQQuickWebEngineProfileFromPointer(ptr unsafe.Pointer) *QQuickWebEngineProfile {
	var n = NewQQuickWebEngineProfileFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickWebEngineProfile_") {
		n.SetObjectName("QQuickWebEngineProfile_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickWebEngineProfile) CachePath() string {
	defer qt.Recovering("QQuickWebEngineProfile::cachePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickWebEngineProfile_CachePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguage() string {
	defer qt.Recovering("QQuickWebEngineProfile::httpAcceptLanguage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickWebEngineProfile_HttpAcceptLanguage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSize() int {
	defer qt.Recovering("QQuickWebEngineProfile::httpCacheMaximumSize")

	if ptr.Pointer() != nil {
		return int(C.QQuickWebEngineProfile_HttpCacheMaximumSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWebEngineProfile) HttpCacheType() QQuickWebEngineProfile__HttpCacheType {
	defer qt.Recovering("QQuickWebEngineProfile::httpCacheType")

	if ptr.Pointer() != nil {
		return QQuickWebEngineProfile__HttpCacheType(C.QQuickWebEngineProfile_HttpCacheType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWebEngineProfile) HttpUserAgent() string {
	defer qt.Recovering("QQuickWebEngineProfile::httpUserAgent")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickWebEngineProfile_HttpUserAgent(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) IsOffTheRecord() bool {
	defer qt.Recovering("QQuickWebEngineProfile::isOffTheRecord")

	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_IsOffTheRecord(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicy() QQuickWebEngineProfile__PersistentCookiesPolicy {
	defer qt.Recovering("QQuickWebEngineProfile::persistentCookiesPolicy")

	if ptr.Pointer() != nil {
		return QQuickWebEngineProfile__PersistentCookiesPolicy(C.QQuickWebEngineProfile_PersistentCookiesPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePath() string {
	defer qt.Recovering("QQuickWebEngineProfile::persistentStoragePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickWebEngineProfile_PersistentStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineProfile) SetCachePath(path string) {
	defer qt.Recovering("QQuickWebEngineProfile::setCachePath")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetCachePath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	defer qt.Recovering("QQuickWebEngineProfile::setHttpAcceptLanguage")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetHttpAcceptLanguage(ptr.Pointer(), C.CString(httpAcceptLanguage))
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	defer qt.Recovering("QQuickWebEngineProfile::setHttpCacheMaximumSize")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetHttpCacheMaximumSize(ptr.Pointer(), C.int(maxSize))
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpCacheType(vqq QQuickWebEngineProfile__HttpCacheType) {
	defer qt.Recovering("QQuickWebEngineProfile::setHttpCacheType")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetHttpCacheType(ptr.Pointer(), C.int(vqq))
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpUserAgent(userAgent string) {
	defer qt.Recovering("QQuickWebEngineProfile::setHttpUserAgent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetHttpUserAgent(ptr.Pointer(), C.CString(userAgent))
	}
}

func (ptr *QQuickWebEngineProfile) SetOffTheRecord(offTheRecord bool) {
	defer qt.Recovering("QQuickWebEngineProfile::setOffTheRecord")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetOffTheRecord(ptr.Pointer(), C.int(qt.GoBoolToInt(offTheRecord)))
	}
}

func (ptr *QQuickWebEngineProfile) SetPersistentCookiesPolicy(vqq QQuickWebEngineProfile__PersistentCookiesPolicy) {
	defer qt.Recovering("QQuickWebEngineProfile::setPersistentCookiesPolicy")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetPersistentCookiesPolicy(ptr.Pointer(), C.int(vqq))
	}
}

func (ptr *QQuickWebEngineProfile) SetPersistentStoragePath(path string) {
	defer qt.Recovering("QQuickWebEngineProfile::setPersistentStoragePath")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetPersistentStoragePath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQuickWebEngineProfile) SetStorageName(name string) {
	defer qt.Recovering("QQuickWebEngineProfile::setStorageName")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetStorageName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QQuickWebEngineProfile) StorageName() string {
	defer qt.Recovering("QQuickWebEngineProfile::storageName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickWebEngineProfile_StorageName(ptr.Pointer()))
	}
	return ""
}

func NewQQuickWebEngineProfile(parent core.QObject_ITF) *QQuickWebEngineProfile {
	defer qt.Recovering("QQuickWebEngineProfile::QQuickWebEngineProfile")

	return newQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_NewQQuickWebEngineProfile(core.PointerFromQObject(parent)))
}

//export callbackQQuickWebEngineProfile_CachePathChanged
func callbackQQuickWebEngineProfile_CachePathChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::cachePathChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cachePathChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectCachePathChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::cachePathChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectCachePathChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cachePathChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectCachePathChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::cachePathChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectCachePathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cachePathChanged")
	}
}

func (ptr *QQuickWebEngineProfile) CachePathChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::cachePathChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CachePathChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) CookieStore() *QWebEngineCookieStore {
	defer qt.Recovering("QQuickWebEngineProfile::cookieStore")

	if ptr.Pointer() != nil {
		return NewQWebEngineCookieStoreFromPointer(C.QQuickWebEngineProfile_CookieStore(ptr.Pointer()))
	}
	return nil
}

func QQuickWebEngineProfile_DefaultProfile() *QQuickWebEngineProfile {
	defer qt.Recovering("QQuickWebEngineProfile::defaultProfile")

	return NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile())
}

func (ptr *QQuickWebEngineProfile) DefaultProfile() *QQuickWebEngineProfile {
	defer qt.Recovering("QQuickWebEngineProfile::defaultProfile")

	return NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile())
}

//export callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged
func callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::httpAcceptLanguageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "httpAcceptLanguageChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpAcceptLanguageChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::httpAcceptLanguageChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpAcceptLanguageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "httpAcceptLanguageChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpAcceptLanguageChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::httpAcceptLanguageChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpAcceptLanguageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "httpAcceptLanguageChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguageChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::httpAcceptLanguageChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpAcceptLanguageChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged
func callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::httpCacheMaximumSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "httpCacheMaximumSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheMaximumSizeChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::httpCacheMaximumSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpCacheMaximumSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "httpCacheMaximumSizeChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheMaximumSizeChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::httpCacheMaximumSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpCacheMaximumSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "httpCacheMaximumSizeChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSizeChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::httpCacheMaximumSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpCacheMaximumSizeChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_HttpCacheTypeChanged
func callbackQQuickWebEngineProfile_HttpCacheTypeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::httpCacheTypeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "httpCacheTypeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheTypeChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::httpCacheTypeChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpCacheTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "httpCacheTypeChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheTypeChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::httpCacheTypeChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpCacheTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "httpCacheTypeChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheTypeChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::httpCacheTypeChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpCacheTypeChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_HttpUserAgentChanged
func callbackQQuickWebEngineProfile_HttpUserAgentChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::httpUserAgentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "httpUserAgentChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpUserAgentChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::httpUserAgentChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectHttpUserAgentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "httpUserAgentChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpUserAgentChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::httpUserAgentChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpUserAgentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "httpUserAgentChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpUserAgentChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::httpUserAgentChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpUserAgentChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) InstallUrlSchemeHandler(scheme string, handler QWebEngineUrlSchemeHandler_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::installUrlSchemeHandler")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_InstallUrlSchemeHandler(ptr.Pointer(), C.CString(scheme), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

//export callbackQQuickWebEngineProfile_OffTheRecordChanged
func callbackQQuickWebEngineProfile_OffTheRecordChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::offTheRecordChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "offTheRecordChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectOffTheRecordChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::offTheRecordChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectOffTheRecordChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "offTheRecordChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectOffTheRecordChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::offTheRecordChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectOffTheRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "offTheRecordChanged")
	}
}

func (ptr *QQuickWebEngineProfile) OffTheRecordChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::offTheRecordChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_OffTheRecordChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged
func callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::persistentCookiesPolicyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "persistentCookiesPolicyChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPersistentCookiesPolicyChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::persistentCookiesPolicyChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectPersistentCookiesPolicyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "persistentCookiesPolicyChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentCookiesPolicyChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::persistentCookiesPolicyChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPersistentCookiesPolicyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "persistentCookiesPolicyChanged")
	}
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicyChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::persistentCookiesPolicyChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PersistentCookiesPolicyChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_PersistentStoragePathChanged
func callbackQQuickWebEngineProfile_PersistentStoragePathChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::persistentStoragePathChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "persistentStoragePathChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPersistentStoragePathChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::persistentStoragePathChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectPersistentStoragePathChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "persistentStoragePathChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentStoragePathChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::persistentStoragePathChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPersistentStoragePathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "persistentStoragePathChanged")
	}
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePathChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::persistentStoragePathChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PersistentStoragePathChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) RemoveAllUrlSchemeHandlers() {
	defer qt.Recovering("QQuickWebEngineProfile::removeAllUrlSchemeHandlers")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_RemoveAllUrlSchemeHandlers(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) RemoveUrlScheme(scheme string) {
	defer qt.Recovering("QQuickWebEngineProfile::removeUrlScheme")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_RemoveUrlScheme(ptr.Pointer(), C.CString(scheme))
	}
}

func (ptr *QQuickWebEngineProfile) RemoveUrlSchemeHandler(handler QWebEngineUrlSchemeHandler_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::removeUrlSchemeHandler")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_RemoveUrlSchemeHandler(ptr.Pointer(), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

func (ptr *QQuickWebEngineProfile) SetRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::setRequestInterceptor")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

//export callbackQQuickWebEngineProfile_StorageNameChanged
func callbackQQuickWebEngineProfile_StorageNameChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::storageNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "storageNameChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectStorageNameChanged(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::storageNameChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectStorageNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "storageNameChanged", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectStorageNameChanged() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::storageNameChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectStorageNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "storageNameChanged")
	}
}

func (ptr *QQuickWebEngineProfile) StorageNameChanged() {
	defer qt.Recovering("QQuickWebEngineProfile::storageNameChanged")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_StorageNameChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) UrlSchemeHandler(scheme string) *QWebEngineUrlSchemeHandler {
	defer qt.Recovering("QQuickWebEngineProfile::urlSchemeHandler")

	if ptr.Pointer() != nil {
		return NewQWebEngineUrlSchemeHandlerFromPointer(C.QQuickWebEngineProfile_UrlSchemeHandler(ptr.Pointer(), C.CString(scheme)))
	}
	return nil
}

//export callbackQQuickWebEngineProfile_TimerEvent
func callbackQQuickWebEngineProfile_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWebEngineProfile::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QQuickWebEngineProfile) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_ChildEvent
func callbackQQuickWebEngineProfile_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWebEngineProfile::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QQuickWebEngineProfile) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_ConnectNotify
func callbackQQuickWebEngineProfile_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWebEngineProfile::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickWebEngineProfile) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineProfile_CustomEvent
func callbackQQuickWebEngineProfile_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWebEngineProfile::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QQuickWebEngineProfile) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_DeleteLater
func callbackQQuickWebEngineProfile_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWebEngineProfile::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWebEngineProfile) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickWebEngineProfile) DeleteLater() {
	defer qt.Recovering("QQuickWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWebEngineProfile_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWebEngineProfile) DeleteLaterDefault() {
	defer qt.Recovering("QQuickWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWebEngineProfile_DisconnectNotify
func callbackQQuickWebEngineProfile_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWebEngineProfile::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineProfile_Event
func callbackQQuickWebEngineProfile_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickWebEngineProfile::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWebEngineProfileFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QQuickWebEngineProfile) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickWebEngineProfile::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectEvent() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QQuickWebEngineProfile) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWebEngineProfile::event")

	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWebEngineProfile::event")

	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_EventFilter
func callbackQQuickWebEngineProfile_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickWebEngineProfile::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWebEngineProfileFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickWebEngineProfile) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickWebEngineProfile) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickWebEngineProfile_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_MetaObject
func callbackQQuickWebEngineProfile_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWebEngineProfile::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWebEngineProfileFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWebEngineProfile) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickWebEngineProfile) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWebEngineProfile_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWebEngineProfile_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebEngineCertificateError::Error
type QWebEngineCertificateError__Error int64

const (
	QWebEngineCertificateError__SslPinnedKeyNotInCertificateChain  = QWebEngineCertificateError__Error(-150)
	QWebEngineCertificateError__CertificateCommonNameInvalid       = QWebEngineCertificateError__Error(-200)
	QWebEngineCertificateError__CertificateDateInvalid             = QWebEngineCertificateError__Error(-201)
	QWebEngineCertificateError__CertificateAuthorityInvalid        = QWebEngineCertificateError__Error(-202)
	QWebEngineCertificateError__CertificateContainsErrors          = QWebEngineCertificateError__Error(-203)
	QWebEngineCertificateError__CertificateNoRevocationMechanism   = QWebEngineCertificateError__Error(-204)
	QWebEngineCertificateError__CertificateUnableToCheckRevocation = QWebEngineCertificateError__Error(-205)
	QWebEngineCertificateError__CertificateRevoked                 = QWebEngineCertificateError__Error(-206)
	QWebEngineCertificateError__CertificateInvalid                 = QWebEngineCertificateError__Error(-207)
	QWebEngineCertificateError__CertificateWeakSignatureAlgorithm  = QWebEngineCertificateError__Error(-208)
	QWebEngineCertificateError__CertificateNonUniqueName           = QWebEngineCertificateError__Error(-210)
	QWebEngineCertificateError__CertificateWeakKey                 = QWebEngineCertificateError__Error(-211)
	QWebEngineCertificateError__CertificateNameConstraintViolation = QWebEngineCertificateError__Error(-212)
)

type QWebEngineCertificateError struct {
	ptr unsafe.Pointer
}

type QWebEngineCertificateError_ITF interface {
	QWebEngineCertificateError_PTR() *QWebEngineCertificateError
}

func (p *QWebEngineCertificateError) QWebEngineCertificateError_PTR() *QWebEngineCertificateError {
	return p
}

func (p *QWebEngineCertificateError) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineCertificateError) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineCertificateErrorFromPointer(ptr unsafe.Pointer) *QWebEngineCertificateError {
	var n = NewQWebEngineCertificateErrorFromPointer(ptr)
	return n
}

func (ptr *QWebEngineCertificateError) Error() QWebEngineCertificateError__Error {
	defer qt.Recovering("QWebEngineCertificateError::error")

	if ptr.Pointer() != nil {
		return QWebEngineCertificateError__Error(C.QWebEngineCertificateError_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineCertificateError) ErrorDescription() string {
	defer qt.Recovering("QWebEngineCertificateError::errorDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineCertificateError_ErrorDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineCertificateError) IsOverridable() bool {
	defer qt.Recovering("QWebEngineCertificateError::isOverridable")

	if ptr.Pointer() != nil {
		return C.QWebEngineCertificateError_IsOverridable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineCertificateError) Url() *core.QUrl {
	defer qt.Recovering("QWebEngineCertificateError::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineCertificateError_Url(ptr.Pointer()))
	}
	return nil
}

type QWebEngineCookieStore struct {
	core.QObject
}

type QWebEngineCookieStore_ITF interface {
	core.QObject_ITF
	QWebEngineCookieStore_PTR() *QWebEngineCookieStore
}

func (p *QWebEngineCookieStore) QWebEngineCookieStore_PTR() *QWebEngineCookieStore {
	return p
}

func (p *QWebEngineCookieStore) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineCookieStore) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEngineCookieStoreFromPointer(ptr unsafe.Pointer) *QWebEngineCookieStore {
	var n = NewQWebEngineCookieStoreFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineCookieStore_") {
		n.SetObjectName("QWebEngineCookieStore_" + qt.Identifier())
	}
	return n
}

//export callbackQWebEngineCookieStore_CookieAdded
func callbackQWebEngineCookieStore_CookieAdded(ptr unsafe.Pointer, ptrName *C.char, cookie unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::cookieAdded")

	if signal := qt.GetSignal(C.GoString(ptrName), "cookieAdded"); signal != nil {
		signal.(func(*network.QNetworkCookie))(network.NewQNetworkCookieFromPointer(cookie))
	}

}

func (ptr *QWebEngineCookieStore) ConnectCookieAdded(f func(cookie *network.QNetworkCookie)) {
	defer qt.Recovering("connect QWebEngineCookieStore::cookieAdded")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectCookieAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cookieAdded", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCookieAdded() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::cookieAdded")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectCookieAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cookieAdded")
	}
}

func (ptr *QWebEngineCookieStore) CookieAdded(cookie network.QNetworkCookie_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::cookieAdded")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CookieAdded(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie))
	}
}

//export callbackQWebEngineCookieStore_CookieRemoved
func callbackQWebEngineCookieStore_CookieRemoved(ptr unsafe.Pointer, ptrName *C.char, cookie unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::cookieRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "cookieRemoved"); signal != nil {
		signal.(func(*network.QNetworkCookie))(network.NewQNetworkCookieFromPointer(cookie))
	}

}

func (ptr *QWebEngineCookieStore) ConnectCookieRemoved(f func(cookie *network.QNetworkCookie)) {
	defer qt.Recovering("connect QWebEngineCookieStore::cookieRemoved")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectCookieRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cookieRemoved", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCookieRemoved() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::cookieRemoved")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectCookieRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cookieRemoved")
	}
}

func (ptr *QWebEngineCookieStore) CookieRemoved(cookie network.QNetworkCookie_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::cookieRemoved")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CookieRemoved(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie))
	}
}

func (ptr *QWebEngineCookieStore) DeleteAllCookies() {
	defer qt.Recovering("QWebEngineCookieStore::deleteAllCookies")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteAllCookies(ptr.Pointer())
	}
}

func (ptr *QWebEngineCookieStore) DeleteCookie(cookie network.QNetworkCookie_ITF, origin core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::deleteCookie")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteCookie(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(origin))
	}
}

func (ptr *QWebEngineCookieStore) DeleteSessionCookies() {
	defer qt.Recovering("QWebEngineCookieStore::deleteSessionCookies")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteSessionCookies(ptr.Pointer())
	}
}

func (ptr *QWebEngineCookieStore) LoadAllCookies() {
	defer qt.Recovering("QWebEngineCookieStore::loadAllCookies")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_LoadAllCookies(ptr.Pointer())
	}
}

func (ptr *QWebEngineCookieStore) SetCookie(cookie network.QNetworkCookie_ITF, origin core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::setCookie")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_SetCookie(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(origin))
	}
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStore() {
	defer qt.Recovering("QWebEngineCookieStore::~QWebEngineCookieStore")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineCookieStore_TimerEvent
func callbackQWebEngineCookieStore_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineCookieStore::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineCookieStore) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineCookieStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineCookieStore_ChildEvent
func callbackQWebEngineCookieStore_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineCookieStore::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineCookieStore) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineCookieStore) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineCookieStore_ConnectNotify
func callbackQWebEngineCookieStore_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineCookieStore::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineCookieStore) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineCookieStore_CustomEvent
func callbackQWebEngineCookieStore_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineCookieStore::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineCookieStore) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineCookieStore) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineCookieStore_DeleteLater
func callbackQWebEngineCookieStore_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineCookieStore::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineCookieStore) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineCookieStore::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineCookieStore) DeleteLater() {
	defer qt.Recovering("QWebEngineCookieStore::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineCookieStore_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineCookieStore) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineCookieStore::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineCookieStore_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineCookieStore_DisconnectNotify
func callbackQWebEngineCookieStore_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineCookieStore::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineCookieStore::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineCookieStore) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineCookieStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineCookieStore::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineCookieStore_Event
func callbackQWebEngineCookieStore_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineCookieStore::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineCookieStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEngineCookieStore) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineCookieStore::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineCookieStore) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineCookieStore::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineCookieStore) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineCookieStore::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineCookieStore_EventFilter
func callbackQWebEngineCookieStore_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineCookieStore::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineCookieStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineCookieStore) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineCookieStore::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineCookieStore) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineCookieStore::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineCookieStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineCookieStore::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineCookieStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineCookieStore_MetaObject
func callbackQWebEngineCookieStore_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineCookieStore::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineCookieStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineCookieStore) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineCookieStore::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineCookieStore) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineCookieStore::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineCookieStore) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineCookieStore::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineCookieStore_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineCookieStore) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineCookieStore::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineCookieStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebEngineDownloadItem::DownloadState
type QWebEngineDownloadItem__DownloadState int64

const (
	QWebEngineDownloadItem__DownloadRequested   = QWebEngineDownloadItem__DownloadState(0)
	QWebEngineDownloadItem__DownloadInProgress  = QWebEngineDownloadItem__DownloadState(1)
	QWebEngineDownloadItem__DownloadCompleted   = QWebEngineDownloadItem__DownloadState(2)
	QWebEngineDownloadItem__DownloadCancelled   = QWebEngineDownloadItem__DownloadState(3)
	QWebEngineDownloadItem__DownloadInterrupted = QWebEngineDownloadItem__DownloadState(4)
)

type QWebEngineDownloadItem struct {
	core.QObject
}

type QWebEngineDownloadItem_ITF interface {
	core.QObject_ITF
	QWebEngineDownloadItem_PTR() *QWebEngineDownloadItem
}

func (p *QWebEngineDownloadItem) QWebEngineDownloadItem_PTR() *QWebEngineDownloadItem {
	return p
}

func (p *QWebEngineDownloadItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineDownloadItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEngineDownloadItemFromPointer(ptr unsafe.Pointer) *QWebEngineDownloadItem {
	var n = NewQWebEngineDownloadItemFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineDownloadItem_") {
		n.SetObjectName("QWebEngineDownloadItem_" + qt.Identifier())
	}
	return n
}

//export callbackQWebEngineDownloadItem_Accept
func callbackQWebEngineDownloadItem_Accept(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineDownloadItem::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineDownloadItem) ConnectAccept(f func()) {
	defer qt.Recovering("connect QWebEngineDownloadItem::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectAccept() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

func (ptr *QWebEngineDownloadItem) Accept() {
	defer qt.Recovering("QWebEngineDownloadItem::accept")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_Accept(ptr.Pointer())
	}
}

//export callbackQWebEngineDownloadItem_Cancel
func callbackQWebEngineDownloadItem_Cancel(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineDownloadItem::cancel")

	if signal := qt.GetSignal(C.GoString(ptrName), "cancel"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineDownloadItem) ConnectCancel(f func()) {
	defer qt.Recovering("connect QWebEngineDownloadItem::cancel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "cancel", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectCancel() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::cancel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "cancel")
	}
}

func (ptr *QWebEngineDownloadItem) Cancel() {
	defer qt.Recovering("QWebEngineDownloadItem::cancel")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_Cancel(ptr.Pointer())
	}
}

//export callbackQWebEngineDownloadItem_DownloadProgress
func callbackQWebEngineDownloadItem_DownloadProgress(ptr unsafe.Pointer, ptrName *C.char, bytesReceived C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QWebEngineDownloadItem::downloadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "downloadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesReceived), int64(bytesTotal))
	}

}

func (ptr *QWebEngineDownloadItem) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::downloadProgress")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectDownloadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "downloadProgress", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectDownloadProgress() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::downloadProgress")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectDownloadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "downloadProgress")
	}
}

func (ptr *QWebEngineDownloadItem) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	defer qt.Recovering("QWebEngineDownloadItem::downloadProgress")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DownloadProgress(ptr.Pointer(), C.longlong(bytesReceived), C.longlong(bytesTotal))
	}
}

//export callbackQWebEngineDownloadItem_Finished
func callbackQWebEngineDownloadItem_Finished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineDownloadItem::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineDownloadItem) ConnectFinished(f func()) {
	defer qt.Recovering("connect QWebEngineDownloadItem::finished")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectFinished() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::finished")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QWebEngineDownloadItem) Finished() {
	defer qt.Recovering("QWebEngineDownloadItem::finished")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_Finished(ptr.Pointer())
	}
}

func (ptr *QWebEngineDownloadItem) IsFinished() bool {
	defer qt.Recovering("QWebEngineDownloadItem::isFinished")

	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineDownloadItem) MimeType() string {
	defer qt.Recovering("QWebEngineDownloadItem::mimeType")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineDownloadItem_MimeType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineDownloadItem) Path() string {
	defer qt.Recovering("QWebEngineDownloadItem::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineDownloadItem_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineDownloadItem) ReceivedBytes() int64 {
	defer qt.Recovering("QWebEngineDownloadItem::receivedBytes")

	if ptr.Pointer() != nil {
		return int64(C.QWebEngineDownloadItem_ReceivedBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineDownloadItem) SetPath(path string) {
	defer qt.Recovering("QWebEngineDownloadItem::setPath")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_SetPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QWebEngineDownloadItem) State() QWebEngineDownloadItem__DownloadState {
	defer qt.Recovering("QWebEngineDownloadItem::state")

	if ptr.Pointer() != nil {
		return QWebEngineDownloadItem__DownloadState(C.QWebEngineDownloadItem_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebEngineDownloadItem_StateChanged
func callbackQWebEngineDownloadItem_StateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QWebEngineDownloadItem::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QWebEngineDownloadItem__DownloadState))(QWebEngineDownloadItem__DownloadState(state))
	}

}

func (ptr *QWebEngineDownloadItem) ConnectStateChanged(f func(state QWebEngineDownloadItem__DownloadState)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QWebEngineDownloadItem) StateChanged(state QWebEngineDownloadItem__DownloadState) {
	defer qt.Recovering("QWebEngineDownloadItem::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QWebEngineDownloadItem) TotalBytes() int64 {
	defer qt.Recovering("QWebEngineDownloadItem::totalBytes")

	if ptr.Pointer() != nil {
		return int64(C.QWebEngineDownloadItem_TotalBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineDownloadItem) Url() *core.QUrl {
	defer qt.Recovering("QWebEngineDownloadItem::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineDownloadItem_Url(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineDownloadItem_TimerEvent
func callbackQWebEngineDownloadItem_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineDownloadItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineDownloadItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineDownloadItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineDownloadItem_ChildEvent
func callbackQWebEngineDownloadItem_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineDownloadItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineDownloadItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineDownloadItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineDownloadItem_ConnectNotify
func callbackQWebEngineDownloadItem_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineDownloadItem::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineDownloadItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineDownloadItem_CustomEvent
func callbackQWebEngineDownloadItem_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineDownloadItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineDownloadItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineDownloadItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineDownloadItem_DeleteLater
func callbackQWebEngineDownloadItem_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineDownloadItem::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineDownloadItem) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineDownloadItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineDownloadItem) DeleteLater() {
	defer qt.Recovering("QWebEngineDownloadItem::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineDownloadItem_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineDownloadItem) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineDownloadItem::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineDownloadItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineDownloadItem_DisconnectNotify
func callbackQWebEngineDownloadItem_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineDownloadItem::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineDownloadItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineDownloadItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineDownloadItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineDownloadItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineDownloadItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineDownloadItem_Event
func callbackQWebEngineDownloadItem_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineDownloadItem::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineDownloadItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEngineDownloadItem) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineDownloadItem::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineDownloadItem) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineDownloadItem::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineDownloadItem) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineDownloadItem::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineDownloadItem_EventFilter
func callbackQWebEngineDownloadItem_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineDownloadItem::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineDownloadItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineDownloadItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineDownloadItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineDownloadItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineDownloadItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineDownloadItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineDownloadItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineDownloadItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineDownloadItem_MetaObject
func callbackQWebEngineDownloadItem_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineDownloadItem::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineDownloadItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineDownloadItem) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineDownloadItem::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineDownloadItem) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineDownloadItem::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineDownloadItem) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineDownloadItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineDownloadItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineDownloadItem) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineDownloadItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineDownloadItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebEngineHistory struct {
	ptr unsafe.Pointer
}

type QWebEngineHistory_ITF interface {
	QWebEngineHistory_PTR() *QWebEngineHistory
}

func (p *QWebEngineHistory) QWebEngineHistory_PTR() *QWebEngineHistory {
	return p
}

func (p *QWebEngineHistory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineHistory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineHistoryFromPointer(ptr unsafe.Pointer) *QWebEngineHistory {
	var n = NewQWebEngineHistoryFromPointer(ptr)
	return n
}

func (ptr *QWebEngineHistory) Back() {
	defer qt.Recovering("QWebEngineHistory::back")

	if ptr.Pointer() != nil {
		C.QWebEngineHistory_Back(ptr.Pointer())
	}
}

func (ptr *QWebEngineHistory) BackItem() *QWebEngineHistoryItem {
	defer qt.Recovering("QWebEngineHistory::backItem")

	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_BackItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistory) CanGoBack() bool {
	defer qt.Recovering("QWebEngineHistory::canGoBack")

	if ptr.Pointer() != nil {
		return C.QWebEngineHistory_CanGoBack(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineHistory) CanGoForward() bool {
	defer qt.Recovering("QWebEngineHistory::canGoForward")

	if ptr.Pointer() != nil {
		return C.QWebEngineHistory_CanGoForward(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineHistory) Clear() {
	defer qt.Recovering("QWebEngineHistory::clear")

	if ptr.Pointer() != nil {
		C.QWebEngineHistory_Clear(ptr.Pointer())
	}
}

func (ptr *QWebEngineHistory) Count() int {
	defer qt.Recovering("QWebEngineHistory::count")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineHistory_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineHistory) CurrentItem() *QWebEngineHistoryItem {
	defer qt.Recovering("QWebEngineHistory::currentItem")

	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistory) CurrentItemIndex() int {
	defer qt.Recovering("QWebEngineHistory::currentItemIndex")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineHistory_CurrentItemIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineHistory) Forward() {
	defer qt.Recovering("QWebEngineHistory::forward")

	if ptr.Pointer() != nil {
		C.QWebEngineHistory_Forward(ptr.Pointer())
	}
}

func (ptr *QWebEngineHistory) ForwardItem() *QWebEngineHistoryItem {
	defer qt.Recovering("QWebEngineHistory::forwardItem")

	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_ForwardItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistory) GoToItem(item QWebEngineHistoryItem_ITF) {
	defer qt.Recovering("QWebEngineHistory::goToItem")

	if ptr.Pointer() != nil {
		C.QWebEngineHistory_GoToItem(ptr.Pointer(), PointerFromQWebEngineHistoryItem(item))
	}
}

func (ptr *QWebEngineHistory) ItemAt(i int) *QWebEngineHistoryItem {
	defer qt.Recovering("QWebEngineHistory::itemAt")

	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryItemFromPointer(C.QWebEngineHistory_ItemAt(ptr.Pointer(), C.int(i)))
	}
	return nil
}

type QWebEngineHistoryItem struct {
	ptr unsafe.Pointer
}

type QWebEngineHistoryItem_ITF interface {
	QWebEngineHistoryItem_PTR() *QWebEngineHistoryItem
}

func (p *QWebEngineHistoryItem) QWebEngineHistoryItem_PTR() *QWebEngineHistoryItem {
	return p
}

func (p *QWebEngineHistoryItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineHistoryItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineHistoryItemFromPointer(ptr unsafe.Pointer) *QWebEngineHistoryItem {
	var n = NewQWebEngineHistoryItemFromPointer(ptr)
	return n
}

func NewQWebEngineHistoryItem(other QWebEngineHistoryItem_ITF) *QWebEngineHistoryItem {
	defer qt.Recovering("QWebEngineHistoryItem::QWebEngineHistoryItem")

	return newQWebEngineHistoryItemFromPointer(C.QWebEngineHistoryItem_NewQWebEngineHistoryItem(PointerFromQWebEngineHistoryItem(other)))
}

func (ptr *QWebEngineHistoryItem) IsValid() bool {
	defer qt.Recovering("QWebEngineHistoryItem::isValid")

	if ptr.Pointer() != nil {
		return C.QWebEngineHistoryItem_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineHistoryItem) LastVisited() *core.QDateTime {
	defer qt.Recovering("QWebEngineHistoryItem::lastVisited")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QWebEngineHistoryItem_LastVisited(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) OriginalUrl() *core.QUrl {
	defer qt.Recovering("QWebEngineHistoryItem::originalUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineHistoryItem_OriginalUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) Title() string {
	defer qt.Recovering("QWebEngineHistoryItem::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineHistoryItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineHistoryItem) Url() *core.QUrl {
	defer qt.Recovering("QWebEngineHistoryItem::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineHistoryItem_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) DestroyQWebEngineHistoryItem() {
	defer qt.Recovering("QWebEngineHistoryItem::~QWebEngineHistoryItem")

	if ptr.Pointer() != nil {
		C.QWebEngineHistoryItem_DestroyQWebEngineHistoryItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineHistoryItem) IconUrl() *core.QUrl {
	defer qt.Recovering("QWebEngineHistoryItem::iconUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineHistoryItem_IconUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineHistoryItem) Swap(other QWebEngineHistoryItem_ITF) {
	defer qt.Recovering("QWebEngineHistoryItem::swap")

	if ptr.Pointer() != nil {
		C.QWebEngineHistoryItem_Swap(ptr.Pointer(), PointerFromQWebEngineHistoryItem(other))
	}
}

//QWebEnginePage::Feature
type QWebEnginePage__Feature int64

const (
	QWebEnginePage__Geolocation            = QWebEnginePage__Feature(1)
	QWebEnginePage__MediaAudioCapture      = QWebEnginePage__Feature(2)
	QWebEnginePage__MediaVideoCapture      = QWebEnginePage__Feature(3)
	QWebEnginePage__MediaAudioVideoCapture = QWebEnginePage__Feature(4)
	QWebEnginePage__MouseLock              = QWebEnginePage__Feature(5)
)

//QWebEnginePage::FileSelectionMode
type QWebEnginePage__FileSelectionMode int64

const (
	QWebEnginePage__FileSelectOpen         = QWebEnginePage__FileSelectionMode(0)
	QWebEnginePage__FileSelectOpenMultiple = QWebEnginePage__FileSelectionMode(1)
)

//QWebEnginePage::FindFlag
type QWebEnginePage__FindFlag int64

const (
	QWebEnginePage__FindBackward        = QWebEnginePage__FindFlag(1)
	QWebEnginePage__FindCaseSensitively = QWebEnginePage__FindFlag(2)
)

//QWebEnginePage::JavaScriptConsoleMessageLevel
type QWebEnginePage__JavaScriptConsoleMessageLevel int64

const (
	QWebEnginePage__InfoMessageLevel    = QWebEnginePage__JavaScriptConsoleMessageLevel(0)
	QWebEnginePage__WarningMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(1)
	QWebEnginePage__ErrorMessageLevel   = QWebEnginePage__JavaScriptConsoleMessageLevel(2)
)

//QWebEnginePage::NavigationType
type QWebEnginePage__NavigationType int64

const (
	QWebEnginePage__NavigationTypeLinkClicked   = QWebEnginePage__NavigationType(0)
	QWebEnginePage__NavigationTypeTyped         = QWebEnginePage__NavigationType(1)
	QWebEnginePage__NavigationTypeFormSubmitted = QWebEnginePage__NavigationType(2)
	QWebEnginePage__NavigationTypeBackForward   = QWebEnginePage__NavigationType(3)
	QWebEnginePage__NavigationTypeReload        = QWebEnginePage__NavigationType(4)
	QWebEnginePage__NavigationTypeOther         = QWebEnginePage__NavigationType(5)
)

//QWebEnginePage::PermissionPolicy
type QWebEnginePage__PermissionPolicy int64

const (
	QWebEnginePage__PermissionUnknown       = QWebEnginePage__PermissionPolicy(0)
	QWebEnginePage__PermissionGrantedByUser = QWebEnginePage__PermissionPolicy(1)
	QWebEnginePage__PermissionDeniedByUser  = QWebEnginePage__PermissionPolicy(2)
)

//QWebEnginePage::RenderProcessTerminationStatus
type QWebEnginePage__RenderProcessTerminationStatus int64

const (
	QWebEnginePage__NormalTerminationStatus   = QWebEnginePage__RenderProcessTerminationStatus(0)
	QWebEnginePage__AbnormalTerminationStatus = QWebEnginePage__RenderProcessTerminationStatus(1)
	QWebEnginePage__CrashedTerminationStatus  = QWebEnginePage__RenderProcessTerminationStatus(2)
	QWebEnginePage__KilledTerminationStatus   = QWebEnginePage__RenderProcessTerminationStatus(3)
)

//QWebEnginePage::WebAction
type QWebEnginePage__WebAction int64

const (
	QWebEnginePage__NoWebAction             = QWebEnginePage__WebAction(-1)
	QWebEnginePage__Back                    = QWebEnginePage__WebAction(0)
	QWebEnginePage__Forward                 = QWebEnginePage__WebAction(1)
	QWebEnginePage__Stop                    = QWebEnginePage__WebAction(2)
	QWebEnginePage__Reload                  = QWebEnginePage__WebAction(3)
	QWebEnginePage__Cut                     = QWebEnginePage__WebAction(4)
	QWebEnginePage__Copy                    = QWebEnginePage__WebAction(5)
	QWebEnginePage__Paste                   = QWebEnginePage__WebAction(6)
	QWebEnginePage__Undo                    = QWebEnginePage__WebAction(7)
	QWebEnginePage__Redo                    = QWebEnginePage__WebAction(8)
	QWebEnginePage__SelectAll               = QWebEnginePage__WebAction(9)
	QWebEnginePage__ReloadAndBypassCache    = QWebEnginePage__WebAction(10)
	QWebEnginePage__PasteAndMatchStyle      = QWebEnginePage__WebAction(11)
	QWebEnginePage__OpenLinkInThisWindow    = QWebEnginePage__WebAction(12)
	QWebEnginePage__OpenLinkInNewWindow     = QWebEnginePage__WebAction(13)
	QWebEnginePage__OpenLinkInNewTab        = QWebEnginePage__WebAction(14)
	QWebEnginePage__CopyLinkToClipboard     = QWebEnginePage__WebAction(15)
	QWebEnginePage__DownloadLinkToDisk      = QWebEnginePage__WebAction(16)
	QWebEnginePage__CopyImageToClipboard    = QWebEnginePage__WebAction(17)
	QWebEnginePage__CopyImageUrlToClipboard = QWebEnginePage__WebAction(18)
	QWebEnginePage__DownloadImageToDisk     = QWebEnginePage__WebAction(19)
	QWebEnginePage__CopyMediaUrlToClipboard = QWebEnginePage__WebAction(20)
	QWebEnginePage__ToggleMediaControls     = QWebEnginePage__WebAction(21)
	QWebEnginePage__ToggleMediaLoop         = QWebEnginePage__WebAction(22)
	QWebEnginePage__ToggleMediaPlayPause    = QWebEnginePage__WebAction(23)
	QWebEnginePage__ToggleMediaMute         = QWebEnginePage__WebAction(24)
	QWebEnginePage__DownloadMediaToDisk     = QWebEnginePage__WebAction(25)
	QWebEnginePage__InspectElement          = QWebEnginePage__WebAction(26)
	QWebEnginePage__ExitFullScreen          = QWebEnginePage__WebAction(27)
	QWebEnginePage__RequestClose            = QWebEnginePage__WebAction(28)
	QWebEnginePage__WebActionCount          = QWebEnginePage__WebAction(29)
)

//QWebEnginePage::WebWindowType
type QWebEnginePage__WebWindowType int64

const (
	QWebEnginePage__WebBrowserWindow = QWebEnginePage__WebWindowType(0)
	QWebEnginePage__WebBrowserTab    = QWebEnginePage__WebWindowType(1)
	QWebEnginePage__WebDialog        = QWebEnginePage__WebWindowType(2)
)

type QWebEnginePage struct {
	core.QObject
}

type QWebEnginePage_ITF interface {
	core.QObject_ITF
	QWebEnginePage_PTR() *QWebEnginePage
}

func (p *QWebEnginePage) QWebEnginePage_PTR() *QWebEnginePage {
	return p
}

func (p *QWebEnginePage) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEnginePage) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEnginePageFromPointer(ptr unsafe.Pointer) *QWebEnginePage {
	var n = NewQWebEnginePageFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEnginePage_") {
		n.SetObjectName("QWebEnginePage_" + qt.Identifier())
	}
	return n
}

func NewQWebEnginePage(parent core.QObject_ITF) *QWebEnginePage {
	defer qt.Recovering("QWebEnginePage::QWebEnginePage")

	return newQWebEnginePageFromPointer(C.QWebEnginePage_NewQWebEnginePage(core.PointerFromQObject(parent)))
}

//export callbackQWebEnginePage_AcceptNavigationRequest
func callbackQWebEnginePage_AcceptNavigationRequest(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer, ty C.int, isMainFrame C.int) C.int {
	defer qt.Recovering("callback QWebEnginePage::acceptNavigationRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "acceptNavigationRequest"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QUrl, QWebEnginePage__NavigationType, bool) bool)(core.NewQUrlFromPointer(url), QWebEnginePage__NavigationType(ty), int(isMainFrame) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).AcceptNavigationRequestDefault(core.NewQUrlFromPointer(url), QWebEnginePage__NavigationType(ty), int(isMainFrame) != 0)))
}

func (ptr *QWebEnginePage) ConnectAcceptNavigationRequest(f func(url *core.QUrl, ty QWebEnginePage__NavigationType, isMainFrame bool) bool) {
	defer qt.Recovering("connect QWebEnginePage::acceptNavigationRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "acceptNavigationRequest", f)
	}
}

func (ptr *QWebEnginePage) DisconnectAcceptNavigationRequest() {
	defer qt.Recovering("disconnect QWebEnginePage::acceptNavigationRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "acceptNavigationRequest")
	}
}

func (ptr *QWebEnginePage) AcceptNavigationRequest(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
	defer qt.Recovering("QWebEnginePage::acceptNavigationRequest")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_AcceptNavigationRequest(ptr.Pointer(), core.PointerFromQUrl(url), C.int(ty), C.int(qt.GoBoolToInt(isMainFrame))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) AcceptNavigationRequestDefault(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
	defer qt.Recovering("QWebEnginePage::acceptNavigationRequest")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_AcceptNavigationRequestDefault(ptr.Pointer(), core.PointerFromQUrl(url), C.int(ty), C.int(qt.GoBoolToInt(isMainFrame))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) Action(action QWebEnginePage__WebAction) *widgets.QAction {
	defer qt.Recovering("QWebEnginePage::action")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QWebEnginePage_Action(ptr.Pointer(), C.int(action)))
	}
	return nil
}

func (ptr *QWebEnginePage) BackgroundColor() *gui.QColor {
	defer qt.Recovering("QWebEnginePage::backgroundColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QWebEnginePage_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_ChooseFiles
func callbackQWebEnginePage_ChooseFiles(ptr unsafe.Pointer, ptrName *C.char, mode C.int, oldFiles *C.char, acceptedMimeTypes *C.char) *C.char {
	defer qt.Recovering("callback QWebEnginePage::chooseFiles")

	if signal := qt.GetSignal(C.GoString(ptrName), "chooseFiles"); signal != nil {
		return C.CString(strings.Join(signal.(func(QWebEnginePage__FileSelectionMode, []string, []string) []string)(QWebEnginePage__FileSelectionMode(mode), strings.Split(C.GoString(oldFiles), "|"), strings.Split(C.GoString(acceptedMimeTypes), "|")), "|"))
	}

	return C.CString(strings.Join(NewQWebEnginePageFromPointer(ptr).ChooseFilesDefault(QWebEnginePage__FileSelectionMode(mode), strings.Split(C.GoString(oldFiles), "|"), strings.Split(C.GoString(acceptedMimeTypes), "|")), "|"))
}

func (ptr *QWebEnginePage) ConnectChooseFiles(f func(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string) {
	defer qt.Recovering("connect QWebEnginePage::chooseFiles")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "chooseFiles", f)
	}
}

func (ptr *QWebEnginePage) DisconnectChooseFiles() {
	defer qt.Recovering("disconnect QWebEnginePage::chooseFiles")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "chooseFiles")
	}
}

func (ptr *QWebEnginePage) ChooseFiles(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
	defer qt.Recovering("QWebEnginePage::chooseFiles")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QWebEnginePage_ChooseFiles(ptr.Pointer(), C.int(mode), C.CString(strings.Join(oldFiles, "|")), C.CString(strings.Join(acceptedMimeTypes, "|")))), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebEnginePage) ChooseFilesDefault(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
	defer qt.Recovering("QWebEnginePage::chooseFiles")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QWebEnginePage_ChooseFilesDefault(ptr.Pointer(), C.int(mode), C.CString(strings.Join(oldFiles, "|")), C.CString(strings.Join(acceptedMimeTypes, "|")))), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebEnginePage) CreateStandardContextMenu() *widgets.QMenu {
	defer qt.Recovering("QWebEnginePage::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return widgets.NewQMenuFromPointer(C.QWebEnginePage_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_CreateWindow
func callbackQWebEnginePage_CreateWindow(ptr unsafe.Pointer, ptrName *C.char, ty C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebEnginePage::createWindow")

	if signal := qt.GetSignal(C.GoString(ptrName), "createWindow"); signal != nil {
		return PointerFromQWebEnginePage(signal.(func(QWebEnginePage__WebWindowType) *QWebEnginePage)(QWebEnginePage__WebWindowType(ty)))
	}

	return PointerFromQWebEnginePage(NewQWebEnginePageFromPointer(ptr).CreateWindowDefault(QWebEnginePage__WebWindowType(ty)))
}

func (ptr *QWebEnginePage) ConnectCreateWindow(f func(ty QWebEnginePage__WebWindowType) *QWebEnginePage) {
	defer qt.Recovering("connect QWebEnginePage::createWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createWindow", f)
	}
}

func (ptr *QWebEnginePage) DisconnectCreateWindow() {
	defer qt.Recovering("disconnect QWebEnginePage::createWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createWindow")
	}
}

func (ptr *QWebEnginePage) CreateWindow(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
	defer qt.Recovering("QWebEnginePage::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebEnginePageFromPointer(C.QWebEnginePage_CreateWindow(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QWebEnginePage) CreateWindowDefault(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
	defer qt.Recovering("QWebEnginePage::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebEnginePageFromPointer(C.QWebEnginePage_CreateWindowDefault(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QWebEnginePage) HasSelection() bool {
	defer qt.Recovering("QWebEnginePage::hasSelection")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEnginePage) History() *QWebEngineHistory {
	defer qt.Recovering("QWebEnginePage::history")

	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryFromPointer(C.QWebEnginePage_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) IconUrl() *core.QUrl {
	defer qt.Recovering("QWebEnginePage::iconUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEnginePage_IconUrl(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_JavaScriptAlert
func callbackQWebEnginePage_JavaScriptAlert(ptr unsafe.Pointer, ptrName *C.char, securityOrigin unsafe.Pointer, msg *C.char) {
	defer qt.Recovering("callback QWebEnginePage::javaScriptAlert")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptAlert"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(securityOrigin), C.GoString(msg))
	} else {
		NewQWebEnginePageFromPointer(ptr).JavaScriptAlertDefault(core.NewQUrlFromPointer(securityOrigin), C.GoString(msg))
	}
}

func (ptr *QWebEnginePage) ConnectJavaScriptAlert(f func(securityOrigin *core.QUrl, msg string)) {
	defer qt.Recovering("connect QWebEnginePage::javaScriptAlert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptAlert", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptAlert() {
	defer qt.Recovering("disconnect QWebEnginePage::javaScriptAlert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptAlert")
	}
}

func (ptr *QWebEnginePage) JavaScriptAlert(securityOrigin core.QUrl_ITF, msg string) {
	defer qt.Recovering("QWebEnginePage::javaScriptAlert")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_JavaScriptAlert(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.CString(msg))
	}
}

func (ptr *QWebEnginePage) JavaScriptAlertDefault(securityOrigin core.QUrl_ITF, msg string) {
	defer qt.Recovering("QWebEnginePage::javaScriptAlert")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_JavaScriptAlertDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.CString(msg))
	}
}

//export callbackQWebEnginePage_JavaScriptConfirm
func callbackQWebEnginePage_JavaScriptConfirm(ptr unsafe.Pointer, ptrName *C.char, securityOrigin unsafe.Pointer, msg *C.char) C.int {
	defer qt.Recovering("callback QWebEnginePage::javaScriptConfirm")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptConfirm"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QUrl, string) bool)(core.NewQUrlFromPointer(securityOrigin), C.GoString(msg))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).JavaScriptConfirmDefault(core.NewQUrlFromPointer(securityOrigin), C.GoString(msg))))
}

func (ptr *QWebEnginePage) ConnectJavaScriptConfirm(f func(securityOrigin *core.QUrl, msg string) bool) {
	defer qt.Recovering("connect QWebEnginePage::javaScriptConfirm")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptConfirm", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConfirm() {
	defer qt.Recovering("disconnect QWebEnginePage::javaScriptConfirm")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptConfirm")
	}
}

func (ptr *QWebEnginePage) JavaScriptConfirm(securityOrigin core.QUrl_ITF, msg string) bool {
	defer qt.Recovering("QWebEnginePage::javaScriptConfirm")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_JavaScriptConfirm(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.CString(msg)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) JavaScriptConfirmDefault(securityOrigin core.QUrl_ITF, msg string) bool {
	defer qt.Recovering("QWebEnginePage::javaScriptConfirm")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_JavaScriptConfirmDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.CString(msg)) != 0
	}
	return false
}

//export callbackQWebEnginePage_JavaScriptConsoleMessage
func callbackQWebEnginePage_JavaScriptConsoleMessage(ptr unsafe.Pointer, ptrName *C.char, level C.int, message *C.char, lineNumber C.int, sourceID *C.char) {
	defer qt.Recovering("callback QWebEnginePage::javaScriptConsoleMessage")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptConsoleMessage"); signal != nil {
		signal.(func(QWebEnginePage__JavaScriptConsoleMessageLevel, string, int, string))(QWebEnginePage__JavaScriptConsoleMessageLevel(level), C.GoString(message), int(lineNumber), C.GoString(sourceID))
	} else {
		NewQWebEnginePageFromPointer(ptr).JavaScriptConsoleMessageDefault(QWebEnginePage__JavaScriptConsoleMessageLevel(level), C.GoString(message), int(lineNumber), C.GoString(sourceID))
	}
}

func (ptr *QWebEnginePage) ConnectJavaScriptConsoleMessage(f func(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string)) {
	defer qt.Recovering("connect QWebEnginePage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptConsoleMessage", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConsoleMessage() {
	defer qt.Recovering("disconnect QWebEnginePage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptConsoleMessage")
	}
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessage(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
	defer qt.Recovering("QWebEnginePage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_JavaScriptConsoleMessage(ptr.Pointer(), C.int(level), C.CString(message), C.int(lineNumber), C.CString(sourceID))
	}
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessageDefault(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
	defer qt.Recovering("QWebEnginePage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_JavaScriptConsoleMessageDefault(ptr.Pointer(), C.int(level), C.CString(message), C.int(lineNumber), C.CString(sourceID))
	}
}

//export callbackQWebEnginePage_JavaScriptPrompt
func callbackQWebEnginePage_JavaScriptPrompt(ptr unsafe.Pointer, ptrName *C.char, securityOrigin unsafe.Pointer, msg *C.char, defaultValue *C.char, result *C.char) C.int {
	defer qt.Recovering("callback QWebEnginePage::javaScriptPrompt")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptPrompt"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QUrl, string, string, string) bool)(core.NewQUrlFromPointer(securityOrigin), C.GoString(msg), C.GoString(defaultValue), C.GoString(result))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).JavaScriptPromptDefault(core.NewQUrlFromPointer(securityOrigin), C.GoString(msg), C.GoString(defaultValue), C.GoString(result))))
}

func (ptr *QWebEnginePage) ConnectJavaScriptPrompt(f func(securityOrigin *core.QUrl, msg string, defaultValue string, result string) bool) {
	defer qt.Recovering("connect QWebEnginePage::javaScriptPrompt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptPrompt", f)
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptPrompt() {
	defer qt.Recovering("disconnect QWebEnginePage::javaScriptPrompt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptPrompt")
	}
}

func (ptr *QWebEnginePage) JavaScriptPrompt(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {
	defer qt.Recovering("QWebEnginePage::javaScriptPrompt")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_JavaScriptPrompt(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.CString(msg), C.CString(defaultValue), C.CString(result)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) JavaScriptPromptDefault(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {
	defer qt.Recovering("QWebEnginePage::javaScriptPrompt")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_JavaScriptPromptDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.CString(msg), C.CString(defaultValue), C.CString(result)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEnginePage::load")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) RequestedUrl() *core.QUrl {
	defer qt.Recovering("QWebEnginePage::requestedUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEnginePage_RequestedUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) RunJavaScript2(scriptSource string) {
	defer qt.Recovering("QWebEnginePage::runJavaScript")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_RunJavaScript2(ptr.Pointer(), C.CString(scriptSource))
	}
}

func (ptr *QWebEnginePage) SelectedText() string {
	defer qt.Recovering("QWebEnginePage::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEnginePage_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEnginePage) SetBackgroundColor(color gui.QColor_ITF) {
	defer qt.Recovering("QWebEnginePage::setBackgroundColor")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QWebEnginePage) SetContent(data string, mimeType string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebEnginePage::setContent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetContent(ptr.Pointer(), C.CString(data), C.CString(mimeType), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEnginePage) SetFeaturePermission(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature, policy QWebEnginePage__PermissionPolicy) {
	defer qt.Recovering("QWebEnginePage::setFeaturePermission")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetFeaturePermission(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.int(feature), C.int(policy))
	}
}

func (ptr *QWebEnginePage) SetHtml(html string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebEnginePage::setHtml")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetHtml(ptr.Pointer(), C.CString(html), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEnginePage) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEnginePage::setUrl")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) SetView(view widgets.QWidget_ITF) {
	defer qt.Recovering("QWebEnginePage::setView")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetView(ptr.Pointer(), widgets.PointerFromQWidget(view))
	}
}

func (ptr *QWebEnginePage) SetZoomFactor(factor float64) {
	defer qt.Recovering("QWebEnginePage::setZoomFactor")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebEnginePage) Settings() *QWebEngineSettings {
	defer qt.Recovering("QWebEnginePage::settings")

	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEnginePage_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) Title() string {
	defer qt.Recovering("QWebEnginePage::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEnginePage_Title(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebEnginePage_TriggerAction
func callbackQWebEnginePage_TriggerAction(ptr unsafe.Pointer, ptrName *C.char, action C.int, checked C.int) {
	defer qt.Recovering("callback QWebEnginePage::triggerAction")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggerAction"); signal != nil {
		signal.(func(QWebEnginePage__WebAction, bool))(QWebEnginePage__WebAction(action), int(checked) != 0)
	} else {
		NewQWebEnginePageFromPointer(ptr).TriggerActionDefault(QWebEnginePage__WebAction(action), int(checked) != 0)
	}
}

func (ptr *QWebEnginePage) ConnectTriggerAction(f func(action QWebEnginePage__WebAction, checked bool)) {
	defer qt.Recovering("connect QWebEnginePage::triggerAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "triggerAction", f)
	}
}

func (ptr *QWebEnginePage) DisconnectTriggerAction() {
	defer qt.Recovering("disconnect QWebEnginePage::triggerAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "triggerAction")
	}
}

func (ptr *QWebEnginePage) TriggerAction(action QWebEnginePage__WebAction, checked bool) {
	defer qt.Recovering("QWebEnginePage::triggerAction")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_TriggerAction(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QWebEnginePage) TriggerActionDefault(action QWebEnginePage__WebAction, checked bool) {
	defer qt.Recovering("QWebEnginePage::triggerAction")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_TriggerActionDefault(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QWebEnginePage) Url() *core.QUrl {
	defer qt.Recovering("QWebEnginePage::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEnginePage_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) View() *widgets.QWidget {
	defer qt.Recovering("QWebEnginePage::view")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QWebEnginePage_View(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) ZoomFactor() float64 {
	defer qt.Recovering("QWebEnginePage::zoomFactor")

	if ptr.Pointer() != nil {
		return float64(C.QWebEnginePage_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEnginePage) DestroyQWebEnginePage() {
	defer qt.Recovering("QWebEnginePage::~QWebEnginePage")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEnginePage_DestroyQWebEnginePage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQWebEnginePage2(profile QWebEngineProfile_ITF, parent core.QObject_ITF) *QWebEnginePage {
	defer qt.Recovering("QWebEnginePage::QWebEnginePage")

	return newQWebEnginePageFromPointer(C.QWebEnginePage_NewQWebEnginePage2(PointerFromQWebEngineProfile(profile), core.PointerFromQObject(parent)))
}

//export callbackQWebEnginePage_AuthenticationRequired
func callbackQWebEnginePage_AuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, requestUrl unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::authenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "authenticationRequired"); signal != nil {
		signal.(func(*core.QUrl, *network.QAuthenticator))(core.NewQUrlFromPointer(requestUrl), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebEnginePage) ConnectAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator)) {
	defer qt.Recovering("connect QWebEnginePage::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "authenticationRequired", f)
	}
}

func (ptr *QWebEnginePage) DisconnectAuthenticationRequired() {
	defer qt.Recovering("disconnect QWebEnginePage::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "authenticationRequired")
	}
}

func (ptr *QWebEnginePage) AuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF) {
	defer qt.Recovering("QWebEnginePage::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_AuthenticationRequired(ptr.Pointer(), core.PointerFromQUrl(requestUrl), network.PointerFromQAuthenticator(authenticator))
	}
}

//export callbackQWebEnginePage_Event
func callbackQWebEnginePage_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEnginePage::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEnginePage) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEnginePage::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEnginePage) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEnginePage::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEnginePage) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEnginePage::event")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEnginePage::event")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEnginePage_FeaturePermissionRequestCanceled
func callbackQWebEnginePage_FeaturePermissionRequestCanceled(ptr unsafe.Pointer, ptrName *C.char, securityOrigin unsafe.Pointer, feature C.int) {
	defer qt.Recovering("callback QWebEnginePage::featurePermissionRequestCanceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "featurePermissionRequestCanceled"); signal != nil {
		signal.(func(*core.QUrl, QWebEnginePage__Feature))(core.NewQUrlFromPointer(securityOrigin), QWebEnginePage__Feature(feature))
	}

}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequestCanceled(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {
	defer qt.Recovering("connect QWebEnginePage::featurePermissionRequestCanceled")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "featurePermissionRequestCanceled", f)
	}
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequestCanceled() {
	defer qt.Recovering("disconnect QWebEnginePage::featurePermissionRequestCanceled")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "featurePermissionRequestCanceled")
	}
}

func (ptr *QWebEnginePage) FeaturePermissionRequestCanceled(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {
	defer qt.Recovering("QWebEnginePage::featurePermissionRequestCanceled")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_FeaturePermissionRequestCanceled(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.int(feature))
	}
}

//export callbackQWebEnginePage_FeaturePermissionRequested
func callbackQWebEnginePage_FeaturePermissionRequested(ptr unsafe.Pointer, ptrName *C.char, securityOrigin unsafe.Pointer, feature C.int) {
	defer qt.Recovering("callback QWebEnginePage::featurePermissionRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "featurePermissionRequested"); signal != nil {
		signal.(func(*core.QUrl, QWebEnginePage__Feature))(core.NewQUrlFromPointer(securityOrigin), QWebEnginePage__Feature(feature))
	}

}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequested(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {
	defer qt.Recovering("connect QWebEnginePage::featurePermissionRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectFeaturePermissionRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "featurePermissionRequested", f)
	}
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequested() {
	defer qt.Recovering("disconnect QWebEnginePage::featurePermissionRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectFeaturePermissionRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "featurePermissionRequested")
	}
}

func (ptr *QWebEnginePage) FeaturePermissionRequested(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {
	defer qt.Recovering("QWebEnginePage::featurePermissionRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_FeaturePermissionRequested(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.int(feature))
	}
}

func (ptr *QWebEnginePage) FindText(subString string, options QWebEnginePage__FindFlag) {
	defer qt.Recovering("QWebEnginePage::findText")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_FindText(ptr.Pointer(), C.CString(subString), C.int(options))
	}
}

//export callbackQWebEnginePage_GeometryChangeRequested
func callbackQWebEnginePage_GeometryChangeRequested(ptr unsafe.Pointer, ptrName *C.char, geom unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::geometryChangeRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChangeRequested"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(geom))
	}

}

func (ptr *QWebEnginePage) ConnectGeometryChangeRequested(f func(geom *core.QRect)) {
	defer qt.Recovering("connect QWebEnginePage::geometryChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectGeometryChangeRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometryChangeRequested", f)
	}
}

func (ptr *QWebEnginePage) DisconnectGeometryChangeRequested() {
	defer qt.Recovering("disconnect QWebEnginePage::geometryChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectGeometryChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometryChangeRequested")
	}
}

func (ptr *QWebEnginePage) GeometryChangeRequested(geom core.QRect_ITF) {
	defer qt.Recovering("QWebEnginePage::geometryChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_GeometryChangeRequested(ptr.Pointer(), core.PointerFromQRect(geom))
	}
}

//export callbackQWebEnginePage_IconUrlChanged
func callbackQWebEnginePage_IconUrlChanged(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::iconUrlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconUrlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEnginePage) ConnectIconUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebEnginePage::iconUrlChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectIconUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconUrlChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectIconUrlChanged() {
	defer qt.Recovering("disconnect QWebEnginePage::iconUrlChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectIconUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconUrlChanged")
	}
}

func (ptr *QWebEnginePage) IconUrlChanged(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEnginePage::iconUrlChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_IconUrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebEnginePage_LinkHovered
func callbackQWebEnginePage_LinkHovered(ptr unsafe.Pointer, ptrName *C.char, url *C.char) {
	defer qt.Recovering("callback QWebEnginePage::linkHovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkHovered"); signal != nil {
		signal.(func(string))(C.GoString(url))
	}

}

func (ptr *QWebEnginePage) ConnectLinkHovered(f func(url string)) {
	defer qt.Recovering("connect QWebEnginePage::linkHovered")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLinkHovered() {
	defer qt.Recovering("disconnect QWebEnginePage::linkHovered")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

func (ptr *QWebEnginePage) LinkHovered(url string) {
	defer qt.Recovering("QWebEnginePage::linkHovered")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_LinkHovered(ptr.Pointer(), C.CString(url))
	}
}

//export callbackQWebEnginePage_LoadFinished
func callbackQWebEnginePage_LoadFinished(ptr unsafe.Pointer, ptrName *C.char, ok C.int) {
	defer qt.Recovering("callback QWebEnginePage::loadFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFinished"); signal != nil {
		signal.(func(bool))(int(ok) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectLoadFinished(f func(ok bool)) {
	defer qt.Recovering("connect QWebEnginePage::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFinished", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLoadFinished() {
	defer qt.Recovering("disconnect QWebEnginePage::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFinished")
	}
}

func (ptr *QWebEnginePage) LoadFinished(ok bool) {
	defer qt.Recovering("QWebEnginePage::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadFinished(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)))
	}
}

//export callbackQWebEnginePage_LoadProgress
func callbackQWebEnginePage_LoadProgress(ptr unsafe.Pointer, ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QWebEnginePage::loadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadProgress"); signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QWebEnginePage) ConnectLoadProgress(f func(progress int)) {
	defer qt.Recovering("connect QWebEnginePage::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadProgress", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLoadProgress() {
	defer qt.Recovering("disconnect QWebEnginePage::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadProgress")
	}
}

func (ptr *QWebEnginePage) LoadProgress(progress int) {
	defer qt.Recovering("QWebEnginePage::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadProgress(ptr.Pointer(), C.int(progress))
	}
}

//export callbackQWebEnginePage_LoadStarted
func callbackQWebEnginePage_LoadStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEnginePage::loadStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEnginePage) ConnectLoadStarted(f func()) {
	defer qt.Recovering("connect QWebEnginePage::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadStarted", f)
	}
}

func (ptr *QWebEnginePage) DisconnectLoadStarted() {
	defer qt.Recovering("disconnect QWebEnginePage::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadStarted")
	}
}

func (ptr *QWebEnginePage) LoadStarted() {
	defer qt.Recovering("QWebEnginePage::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadStarted(ptr.Pointer())
	}
}

func (ptr *QWebEnginePage) Profile() *QWebEngineProfile {
	defer qt.Recovering("QWebEnginePage::profile")

	if ptr.Pointer() != nil {
		return NewQWebEngineProfileFromPointer(C.QWebEnginePage_Profile(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_ProxyAuthenticationRequired
func callbackQWebEnginePage_ProxyAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, requestUrl unsafe.Pointer, authenticator unsafe.Pointer, proxyHost *C.char) {
	defer qt.Recovering("callback QWebEnginePage::proxyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "proxyAuthenticationRequired"); signal != nil {
		signal.(func(*core.QUrl, *network.QAuthenticator, string))(core.NewQUrlFromPointer(requestUrl), network.NewQAuthenticatorFromPointer(authenticator), C.GoString(proxyHost))
	}

}

func (ptr *QWebEnginePage) ConnectProxyAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator, proxyHost string)) {
	defer qt.Recovering("connect QWebEnginePage::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "proxyAuthenticationRequired", f)
	}
}

func (ptr *QWebEnginePage) DisconnectProxyAuthenticationRequired() {
	defer qt.Recovering("disconnect QWebEnginePage::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "proxyAuthenticationRequired")
	}
}

func (ptr *QWebEnginePage) ProxyAuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF, proxyHost string) {
	defer qt.Recovering("QWebEnginePage::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ProxyAuthenticationRequired(ptr.Pointer(), core.PointerFromQUrl(requestUrl), network.PointerFromQAuthenticator(authenticator), C.CString(proxyHost))
	}
}

//export callbackQWebEnginePage_RenderProcessTerminated
func callbackQWebEnginePage_RenderProcessTerminated(ptr unsafe.Pointer, ptrName *C.char, terminationStatus C.int, exitCode C.int) {
	defer qt.Recovering("callback QWebEnginePage::renderProcessTerminated")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderProcessTerminated"); signal != nil {
		signal.(func(QWebEnginePage__RenderProcessTerminationStatus, int))(QWebEnginePage__RenderProcessTerminationStatus(terminationStatus), int(exitCode))
	}

}

func (ptr *QWebEnginePage) ConnectRenderProcessTerminated(f func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int)) {
	defer qt.Recovering("connect QWebEnginePage::renderProcessTerminated")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectRenderProcessTerminated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderProcessTerminated", f)
	}
}

func (ptr *QWebEnginePage) DisconnectRenderProcessTerminated() {
	defer qt.Recovering("disconnect QWebEnginePage::renderProcessTerminated")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectRenderProcessTerminated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderProcessTerminated")
	}
}

func (ptr *QWebEnginePage) RenderProcessTerminated(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int) {
	defer qt.Recovering("QWebEnginePage::renderProcessTerminated")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_RenderProcessTerminated(ptr.Pointer(), C.int(terminationStatus), C.int(exitCode))
	}
}

//export callbackQWebEnginePage_SelectionChanged
func callbackQWebEnginePage_SelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEnginePage::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEnginePage) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QWebEnginePage::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QWebEnginePage::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

func (ptr *QWebEnginePage) SelectionChanged() {
	defer qt.Recovering("QWebEnginePage::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebEnginePage) SetWebChannel(channel webchannel.QWebChannel_ITF) {
	defer qt.Recovering("QWebEnginePage::setWebChannel")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetWebChannel(ptr.Pointer(), webchannel.PointerFromQWebChannel(channel))
	}
}

//export callbackQWebEnginePage_TitleChanged
func callbackQWebEnginePage_TitleChanged(ptr unsafe.Pointer, ptrName *C.char, title *C.char) {
	defer qt.Recovering("callback QWebEnginePage::titleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "titleChanged"); signal != nil {
		signal.(func(string))(C.GoString(title))
	}

}

func (ptr *QWebEnginePage) ConnectTitleChanged(f func(title string)) {
	defer qt.Recovering("connect QWebEnginePage::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "titleChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectTitleChanged() {
	defer qt.Recovering("disconnect QWebEnginePage::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "titleChanged")
	}
}

func (ptr *QWebEnginePage) TitleChanged(title string) {
	defer qt.Recovering("QWebEnginePage::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_TitleChanged(ptr.Pointer(), C.CString(title))
	}
}

//export callbackQWebEnginePage_UrlChanged
func callbackQWebEnginePage_UrlChanged(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::urlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEnginePage) ConnectUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebEnginePage::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "urlChanged", f)
	}
}

func (ptr *QWebEnginePage) DisconnectUrlChanged() {
	defer qt.Recovering("disconnect QWebEnginePage::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "urlChanged")
	}
}

func (ptr *QWebEnginePage) UrlChanged(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEnginePage::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) WebChannel() *webchannel.QWebChannel {
	defer qt.Recovering("QWebEnginePage::webChannel")

	if ptr.Pointer() != nil {
		return webchannel.NewQWebChannelFromPointer(C.QWebEnginePage_WebChannel(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_WindowCloseRequested
func callbackQWebEnginePage_WindowCloseRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEnginePage::windowCloseRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowCloseRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEnginePage) ConnectWindowCloseRequested(f func()) {
	defer qt.Recovering("connect QWebEnginePage::windowCloseRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectWindowCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowCloseRequested", f)
	}
}

func (ptr *QWebEnginePage) DisconnectWindowCloseRequested() {
	defer qt.Recovering("disconnect QWebEnginePage::windowCloseRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectWindowCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowCloseRequested")
	}
}

func (ptr *QWebEnginePage) WindowCloseRequested() {
	defer qt.Recovering("QWebEnginePage::windowCloseRequested")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_WindowCloseRequested(ptr.Pointer())
	}
}

//export callbackQWebEnginePage_TimerEvent
func callbackQWebEnginePage_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEnginePage::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEnginePage) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEnginePage::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEnginePage) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEnginePage::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEnginePage) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEnginePage::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEnginePage_ChildEvent
func callbackQWebEnginePage_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEnginePage::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEnginePage) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEnginePage::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEnginePage) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEnginePage::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEnginePage) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEnginePage::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEnginePage_ConnectNotify
func callbackQWebEnginePage_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEnginePageFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEnginePage) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEnginePage::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEnginePage) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEnginePage::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEnginePage) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEnginePage::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEnginePage) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEnginePage::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEnginePage_CustomEvent
func callbackQWebEnginePage_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEnginePage::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEnginePage) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEnginePage::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEnginePage) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEnginePage::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEnginePage) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEnginePage::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEnginePage_DeleteLater
func callbackQWebEnginePage_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEnginePage::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEnginePageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEnginePage) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEnginePage::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEnginePage) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEnginePage::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEnginePage) DeleteLater() {
	defer qt.Recovering("QWebEnginePage::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEnginePage_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEnginePage) DeleteLaterDefault() {
	defer qt.Recovering("QWebEnginePage::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEnginePage_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEnginePage_DisconnectNotify
func callbackQWebEnginePage_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEnginePage::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEnginePageFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEnginePage) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEnginePage::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEnginePage) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEnginePage::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEnginePage) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEnginePage::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEnginePage) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEnginePage::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEnginePage_EventFilter
func callbackQWebEnginePage_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEnginePage::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEnginePage) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEnginePage::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEnginePage) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEnginePage::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEnginePage) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEnginePage::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEnginePage) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEnginePage::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEnginePage_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEnginePage_MetaObject
func callbackQWebEnginePage_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEnginePage::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEnginePageFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEnginePage) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEnginePage::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEnginePage) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEnginePage::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEnginePage) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEnginePage::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEnginePage_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEnginePage::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEnginePage_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebEngineProfile::HttpCacheType
type QWebEngineProfile__HttpCacheType int64

const (
	QWebEngineProfile__MemoryHttpCache = QWebEngineProfile__HttpCacheType(0)
	QWebEngineProfile__DiskHttpCache   = QWebEngineProfile__HttpCacheType(1)
)

//QWebEngineProfile::PersistentCookiesPolicy
type QWebEngineProfile__PersistentCookiesPolicy int64

const (
	QWebEngineProfile__NoPersistentCookies    = QWebEngineProfile__PersistentCookiesPolicy(0)
	QWebEngineProfile__AllowPersistentCookies = QWebEngineProfile__PersistentCookiesPolicy(1)
	QWebEngineProfile__ForcePersistentCookies = QWebEngineProfile__PersistentCookiesPolicy(2)
)

type QWebEngineProfile struct {
	core.QObject
}

type QWebEngineProfile_ITF interface {
	core.QObject_ITF
	QWebEngineProfile_PTR() *QWebEngineProfile
}

func (p *QWebEngineProfile) QWebEngineProfile_PTR() *QWebEngineProfile {
	return p
}

func (p *QWebEngineProfile) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineProfile) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEngineProfileFromPointer(ptr unsafe.Pointer) *QWebEngineProfile {
	var n = NewQWebEngineProfileFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineProfile_") {
		n.SetObjectName("QWebEngineProfile_" + qt.Identifier())
	}
	return n
}

func NewQWebEngineProfile(parent core.QObject_ITF) *QWebEngineProfile {
	defer qt.Recovering("QWebEngineProfile::QWebEngineProfile")

	return newQWebEngineProfileFromPointer(C.QWebEngineProfile_NewQWebEngineProfile(core.PointerFromQObject(parent)))
}

func NewQWebEngineProfile2(storageName string, parent core.QObject_ITF) *QWebEngineProfile {
	defer qt.Recovering("QWebEngineProfile::QWebEngineProfile")

	return newQWebEngineProfileFromPointer(C.QWebEngineProfile_NewQWebEngineProfile2(C.CString(storageName), core.PointerFromQObject(parent)))
}

func (ptr *QWebEngineProfile) CachePath() string {
	defer qt.Recovering("QWebEngineProfile::cachePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineProfile_CachePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) ClearAllVisitedLinks() {
	defer qt.Recovering("QWebEngineProfile::clearAllVisitedLinks")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ClearAllVisitedLinks(ptr.Pointer())
	}
}

func (ptr *QWebEngineProfile) CookieStore() *QWebEngineCookieStore {
	defer qt.Recovering("QWebEngineProfile::cookieStore")

	if ptr.Pointer() != nil {
		return NewQWebEngineCookieStoreFromPointer(C.QWebEngineProfile_CookieStore(ptr.Pointer()))
	}
	return nil
}

func QWebEngineProfile_DefaultProfile() *QWebEngineProfile {
	defer qt.Recovering("QWebEngineProfile::defaultProfile")

	return NewQWebEngineProfileFromPointer(C.QWebEngineProfile_QWebEngineProfile_DefaultProfile())
}

func (ptr *QWebEngineProfile) DefaultProfile() *QWebEngineProfile {
	defer qt.Recovering("QWebEngineProfile::defaultProfile")

	return NewQWebEngineProfileFromPointer(C.QWebEngineProfile_QWebEngineProfile_DefaultProfile())
}

//export callbackQWebEngineProfile_DownloadRequested
func callbackQWebEngineProfile_DownloadRequested(ptr unsafe.Pointer, ptrName *C.char, download unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineProfile::downloadRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "downloadRequested"); signal != nil {
		signal.(func(*QWebEngineDownloadItem))(NewQWebEngineDownloadItemFromPointer(download))
	}

}

func (ptr *QWebEngineProfile) ConnectDownloadRequested(f func(download *QWebEngineDownloadItem)) {
	defer qt.Recovering("connect QWebEngineProfile::downloadRequested")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectDownloadRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "downloadRequested", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectDownloadRequested() {
	defer qt.Recovering("disconnect QWebEngineProfile::downloadRequested")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectDownloadRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "downloadRequested")
	}
}

func (ptr *QWebEngineProfile) DownloadRequested(download QWebEngineDownloadItem_ITF) {
	defer qt.Recovering("QWebEngineProfile::downloadRequested")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DownloadRequested(ptr.Pointer(), PointerFromQWebEngineDownloadItem(download))
	}
}

func (ptr *QWebEngineProfile) HttpAcceptLanguage() string {
	defer qt.Recovering("QWebEngineProfile::httpAcceptLanguage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineProfile_HttpAcceptLanguage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) HttpCacheMaximumSize() int {
	defer qt.Recovering("QWebEngineProfile::httpCacheMaximumSize")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineProfile_HttpCacheMaximumSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineProfile) HttpCacheType() QWebEngineProfile__HttpCacheType {
	defer qt.Recovering("QWebEngineProfile::httpCacheType")

	if ptr.Pointer() != nil {
		return QWebEngineProfile__HttpCacheType(C.QWebEngineProfile_HttpCacheType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineProfile) HttpUserAgent() string {
	defer qt.Recovering("QWebEngineProfile::httpUserAgent")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineProfile_HttpUserAgent(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) InstallUrlSchemeHandler(scheme string, handler QWebEngineUrlSchemeHandler_ITF) {
	defer qt.Recovering("QWebEngineProfile::installUrlSchemeHandler")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_InstallUrlSchemeHandler(ptr.Pointer(), C.CString(scheme), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

func (ptr *QWebEngineProfile) IsOffTheRecord() bool {
	defer qt.Recovering("QWebEngineProfile::isOffTheRecord")

	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_IsOffTheRecord(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) PersistentCookiesPolicy() QWebEngineProfile__PersistentCookiesPolicy {
	defer qt.Recovering("QWebEngineProfile::persistentCookiesPolicy")

	if ptr.Pointer() != nil {
		return QWebEngineProfile__PersistentCookiesPolicy(C.QWebEngineProfile_PersistentCookiesPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineProfile) PersistentStoragePath() string {
	defer qt.Recovering("QWebEngineProfile::persistentStoragePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineProfile_PersistentStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) RemoveAllUrlSchemeHandlers() {
	defer qt.Recovering("QWebEngineProfile::removeAllUrlSchemeHandlers")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveAllUrlSchemeHandlers(ptr.Pointer())
	}
}

func (ptr *QWebEngineProfile) RemoveUrlScheme(scheme string) {
	defer qt.Recovering("QWebEngineProfile::removeUrlScheme")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveUrlScheme(ptr.Pointer(), C.CString(scheme))
	}
}

func (ptr *QWebEngineProfile) RemoveUrlSchemeHandler(handler QWebEngineUrlSchemeHandler_ITF) {
	defer qt.Recovering("QWebEngineProfile::removeUrlSchemeHandler")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveUrlSchemeHandler(ptr.Pointer(), PointerFromQWebEngineUrlSchemeHandler(handler))
	}
}

func (ptr *QWebEngineProfile) Scripts() *QWebEngineScriptCollection {
	defer qt.Recovering("QWebEngineProfile::scripts")

	if ptr.Pointer() != nil {
		return NewQWebEngineScriptCollectionFromPointer(C.QWebEngineProfile_Scripts(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) SetCachePath(path string) {
	defer qt.Recovering("QWebEngineProfile::setCachePath")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetCachePath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	defer qt.Recovering("QWebEngineProfile::setHttpAcceptLanguage")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpAcceptLanguage(ptr.Pointer(), C.CString(httpAcceptLanguage))
	}
}

func (ptr *QWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	defer qt.Recovering("QWebEngineProfile::setHttpCacheMaximumSize")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpCacheMaximumSize(ptr.Pointer(), C.int(maxSize))
	}
}

func (ptr *QWebEngineProfile) SetHttpCacheType(httpCacheType QWebEngineProfile__HttpCacheType) {
	defer qt.Recovering("QWebEngineProfile::setHttpCacheType")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpCacheType(ptr.Pointer(), C.int(httpCacheType))
	}
}

func (ptr *QWebEngineProfile) SetHttpUserAgent(userAgent string) {
	defer qt.Recovering("QWebEngineProfile::setHttpUserAgent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpUserAgent(ptr.Pointer(), C.CString(userAgent))
	}
}

func (ptr *QWebEngineProfile) SetPersistentCookiesPolicy(newPersistentCookiesPolicy QWebEngineProfile__PersistentCookiesPolicy) {
	defer qt.Recovering("QWebEngineProfile::setPersistentCookiesPolicy")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetPersistentCookiesPolicy(ptr.Pointer(), C.int(newPersistentCookiesPolicy))
	}
}

func (ptr *QWebEngineProfile) SetPersistentStoragePath(path string) {
	defer qt.Recovering("QWebEngineProfile::setPersistentStoragePath")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetPersistentStoragePath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QWebEngineProfile) SetRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	defer qt.Recovering("QWebEngineProfile::setRequestInterceptor")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

func (ptr *QWebEngineProfile) Settings() *QWebEngineSettings {
	defer qt.Recovering("QWebEngineProfile::settings")

	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEngineProfile_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) StorageName() string {
	defer qt.Recovering("QWebEngineProfile::storageName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineProfile_StorageName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineProfile) UrlSchemeHandler(scheme string) *QWebEngineUrlSchemeHandler {
	defer qt.Recovering("QWebEngineProfile::urlSchemeHandler")

	if ptr.Pointer() != nil {
		return NewQWebEngineUrlSchemeHandlerFromPointer(C.QWebEngineProfile_UrlSchemeHandler(ptr.Pointer(), C.CString(scheme)))
	}
	return nil
}

func (ptr *QWebEngineProfile) VisitedLinksContainsUrl(url core.QUrl_ITF) bool {
	defer qt.Recovering("QWebEngineProfile::visitedLinksContainsUrl")

	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_VisitedLinksContainsUrl(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

//export callbackQWebEngineProfile_TimerEvent
func callbackQWebEngineProfile_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineProfile::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineProfile) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineProfile::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineProfile_ChildEvent
func callbackQWebEngineProfile_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineProfile::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineProfile) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineProfile::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineProfile_ConnectNotify
func callbackQWebEngineProfile_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineProfile::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineProfileFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineProfile) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineProfile::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineProfile_CustomEvent
func callbackQWebEngineProfile_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineProfile::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineProfile) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineProfile::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineProfile_DeleteLater
func callbackQWebEngineProfile_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineProfile::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineProfileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineProfile) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineProfile) DeleteLater() {
	defer qt.Recovering("QWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineProfile_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineProfile) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineProfile::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineProfile_DisconnectNotify
func callbackQWebEngineProfile_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineProfile::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineProfileFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineProfile) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineProfile::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineProfile_Event
func callbackQWebEngineProfile_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineProfile::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineProfileFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEngineProfile) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineProfile::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineProfile::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineProfile) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineProfile::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineProfile::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineProfile_EventFilter
func callbackQWebEngineProfile_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineProfile::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineProfileFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineProfile) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineProfile) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineProfile::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineProfile_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineProfile_MetaObject
func callbackQWebEngineProfile_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineProfile::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineProfileFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineProfile) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineProfile) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineProfile) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineProfile_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineProfile::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineProfile_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebEngineScript::InjectionPoint
type QWebEngineScript__InjectionPoint int64

const (
	QWebEngineScript__Deferred         = QWebEngineScript__InjectionPoint(0)
	QWebEngineScript__DocumentReady    = QWebEngineScript__InjectionPoint(1)
	QWebEngineScript__DocumentCreation = QWebEngineScript__InjectionPoint(2)
)

//QWebEngineScript::ScriptWorldId
type QWebEngineScript__ScriptWorldId int64

const (
	QWebEngineScript__MainWorld        = QWebEngineScript__ScriptWorldId(0)
	QWebEngineScript__ApplicationWorld = QWebEngineScript__ScriptWorldId(1)
	QWebEngineScript__UserWorld        = QWebEngineScript__ScriptWorldId(2)
)

type QWebEngineScript struct {
	ptr unsafe.Pointer
}

type QWebEngineScript_ITF interface {
	QWebEngineScript_PTR() *QWebEngineScript
}

func (p *QWebEngineScript) QWebEngineScript_PTR() *QWebEngineScript {
	return p
}

func (p *QWebEngineScript) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineScript) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineScriptFromPointer(ptr unsafe.Pointer) *QWebEngineScript {
	var n = NewQWebEngineScriptFromPointer(ptr)
	return n
}

func NewQWebEngineScript() *QWebEngineScript {
	defer qt.Recovering("QWebEngineScript::QWebEngineScript")

	return newQWebEngineScriptFromPointer(C.QWebEngineScript_NewQWebEngineScript())
}

func NewQWebEngineScript2(other QWebEngineScript_ITF) *QWebEngineScript {
	defer qt.Recovering("QWebEngineScript::QWebEngineScript")

	return newQWebEngineScriptFromPointer(C.QWebEngineScript_NewQWebEngineScript2(PointerFromQWebEngineScript(other)))
}

func (ptr *QWebEngineScript) InjectionPoint() QWebEngineScript__InjectionPoint {
	defer qt.Recovering("QWebEngineScript::injectionPoint")

	if ptr.Pointer() != nil {
		return QWebEngineScript__InjectionPoint(C.QWebEngineScript_InjectionPoint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineScript) IsNull() bool {
	defer qt.Recovering("QWebEngineScript::isNull")

	if ptr.Pointer() != nil {
		return C.QWebEngineScript_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineScript) Name() string {
	defer qt.Recovering("QWebEngineScript::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineScript_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineScript) RunsOnSubFrames() bool {
	defer qt.Recovering("QWebEngineScript::runsOnSubFrames")

	if ptr.Pointer() != nil {
		return C.QWebEngineScript_RunsOnSubFrames(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineScript) SetInjectionPoint(p QWebEngineScript__InjectionPoint) {
	defer qt.Recovering("QWebEngineScript::setInjectionPoint")

	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetInjectionPoint(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QWebEngineScript) SetName(scriptName string) {
	defer qt.Recovering("QWebEngineScript::setName")

	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetName(ptr.Pointer(), C.CString(scriptName))
	}
}

func (ptr *QWebEngineScript) SetRunsOnSubFrames(on bool) {
	defer qt.Recovering("QWebEngineScript::setRunsOnSubFrames")

	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetRunsOnSubFrames(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWebEngineScript) SetSourceCode(scriptSource string) {
	defer qt.Recovering("QWebEngineScript::setSourceCode")

	if ptr.Pointer() != nil {
		C.QWebEngineScript_SetSourceCode(ptr.Pointer(), C.CString(scriptSource))
	}
}

func (ptr *QWebEngineScript) SourceCode() string {
	defer qt.Recovering("QWebEngineScript::sourceCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineScript_SourceCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineScript) Swap(other QWebEngineScript_ITF) {
	defer qt.Recovering("QWebEngineScript::swap")

	if ptr.Pointer() != nil {
		C.QWebEngineScript_Swap(ptr.Pointer(), PointerFromQWebEngineScript(other))
	}
}

func (ptr *QWebEngineScript) DestroyQWebEngineScript() {
	defer qt.Recovering("QWebEngineScript::~QWebEngineScript")

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

func (p *QWebEngineScriptCollection) QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection {
	return p
}

func (p *QWebEngineScriptCollection) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineScriptCollection) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineScriptCollectionFromPointer(ptr unsafe.Pointer) *QWebEngineScriptCollection {
	var n = NewQWebEngineScriptCollectionFromPointer(ptr)
	return n
}

func (ptr *QWebEngineScriptCollection) Clear() {
	defer qt.Recovering("QWebEngineScriptCollection::clear")

	if ptr.Pointer() != nil {
		C.QWebEngineScriptCollection_Clear(ptr.Pointer())
	}
}

func (ptr *QWebEngineScriptCollection) Contains(value QWebEngineScript_ITF) bool {
	defer qt.Recovering("QWebEngineScriptCollection::contains")

	if ptr.Pointer() != nil {
		return C.QWebEngineScriptCollection_Contains(ptr.Pointer(), PointerFromQWebEngineScript(value)) != 0
	}
	return false
}

func (ptr *QWebEngineScriptCollection) Count() int {
	defer qt.Recovering("QWebEngineScriptCollection::count")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineScriptCollection_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineScriptCollection) FindScript(name string) *QWebEngineScript {
	defer qt.Recovering("QWebEngineScriptCollection::findScript")

	if ptr.Pointer() != nil {
		return NewQWebEngineScriptFromPointer(C.QWebEngineScriptCollection_FindScript(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) Insert(s QWebEngineScript_ITF) {
	defer qt.Recovering("QWebEngineScriptCollection::insert")

	if ptr.Pointer() != nil {
		C.QWebEngineScriptCollection_Insert(ptr.Pointer(), PointerFromQWebEngineScript(s))
	}
}

func (ptr *QWebEngineScriptCollection) IsEmpty() bool {
	defer qt.Recovering("QWebEngineScriptCollection::isEmpty")

	if ptr.Pointer() != nil {
		return C.QWebEngineScriptCollection_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineScriptCollection) Remove(script QWebEngineScript_ITF) bool {
	defer qt.Recovering("QWebEngineScriptCollection::remove")

	if ptr.Pointer() != nil {
		return C.QWebEngineScriptCollection_Remove(ptr.Pointer(), PointerFromQWebEngineScript(script)) != 0
	}
	return false
}

func (ptr *QWebEngineScriptCollection) Size() int {
	defer qt.Recovering("QWebEngineScriptCollection::size")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineScriptCollection_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineScriptCollection) DestroyQWebEngineScriptCollection() {
	defer qt.Recovering("QWebEngineScriptCollection::~QWebEngineScriptCollection")

	if ptr.Pointer() != nil {
		C.QWebEngineScriptCollection_DestroyQWebEngineScriptCollection(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QWebEngineSettings::FontFamily
type QWebEngineSettings__FontFamily int64

const (
	QWebEngineSettings__StandardFont  = QWebEngineSettings__FontFamily(0)
	QWebEngineSettings__FixedFont     = QWebEngineSettings__FontFamily(1)
	QWebEngineSettings__SerifFont     = QWebEngineSettings__FontFamily(2)
	QWebEngineSettings__SansSerifFont = QWebEngineSettings__FontFamily(3)
	QWebEngineSettings__CursiveFont   = QWebEngineSettings__FontFamily(4)
	QWebEngineSettings__FantasyFont   = QWebEngineSettings__FontFamily(5)
)

//QWebEngineSettings::FontSize
type QWebEngineSettings__FontSize int64

const (
	QWebEngineSettings__MinimumFontSize        = QWebEngineSettings__FontSize(0)
	QWebEngineSettings__MinimumLogicalFontSize = QWebEngineSettings__FontSize(1)
	QWebEngineSettings__DefaultFontSize        = QWebEngineSettings__FontSize(2)
	QWebEngineSettings__DefaultFixedFontSize   = QWebEngineSettings__FontSize(3)
)

//QWebEngineSettings::WebAttribute
type QWebEngineSettings__WebAttribute int64

const (
	QWebEngineSettings__AutoLoadImages                  = QWebEngineSettings__WebAttribute(0)
	QWebEngineSettings__JavascriptEnabled               = QWebEngineSettings__WebAttribute(1)
	QWebEngineSettings__JavascriptCanOpenWindows        = QWebEngineSettings__WebAttribute(2)
	QWebEngineSettings__JavascriptCanAccessClipboard    = QWebEngineSettings__WebAttribute(3)
	QWebEngineSettings__LinksIncludedInFocusChain       = QWebEngineSettings__WebAttribute(4)
	QWebEngineSettings__LocalStorageEnabled             = QWebEngineSettings__WebAttribute(5)
	QWebEngineSettings__LocalContentCanAccessRemoteUrls = QWebEngineSettings__WebAttribute(6)
	QWebEngineSettings__XSSAuditingEnabled              = QWebEngineSettings__WebAttribute(7)
	QWebEngineSettings__SpatialNavigationEnabled        = QWebEngineSettings__WebAttribute(8)
	QWebEngineSettings__LocalContentCanAccessFileUrls   = QWebEngineSettings__WebAttribute(9)
	QWebEngineSettings__HyperlinkAuditingEnabled        = QWebEngineSettings__WebAttribute(10)
	QWebEngineSettings__ScrollAnimatorEnabled           = QWebEngineSettings__WebAttribute(11)
	QWebEngineSettings__ErrorPageEnabled                = QWebEngineSettings__WebAttribute(12)
	QWebEngineSettings__PluginsEnabled                  = QWebEngineSettings__WebAttribute(13)
	QWebEngineSettings__FullScreenSupportEnabled        = QWebEngineSettings__WebAttribute(14)
)

type QWebEngineSettings struct {
	ptr unsafe.Pointer
}

type QWebEngineSettings_ITF interface {
	QWebEngineSettings_PTR() *QWebEngineSettings
}

func (p *QWebEngineSettings) QWebEngineSettings_PTR() *QWebEngineSettings {
	return p
}

func (p *QWebEngineSettings) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineSettings) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineSettingsFromPointer(ptr unsafe.Pointer) *QWebEngineSettings {
	var n = NewQWebEngineSettingsFromPointer(ptr)
	return n
}

func (ptr *QWebEngineSettings) DefaultTextEncoding() string {
	defer qt.Recovering("QWebEngineSettings::defaultTextEncoding")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineSettings_DefaultTextEncoding(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineSettings) FontFamily(which QWebEngineSettings__FontFamily) string {
	defer qt.Recovering("QWebEngineSettings::fontFamily")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineSettings_FontFamily(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWebEngineSettings) FontSize(ty QWebEngineSettings__FontSize) int {
	defer qt.Recovering("QWebEngineSettings::fontSize")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineSettings_FontSize(ptr.Pointer(), C.int(ty)))
	}
	return 0
}

func QWebEngineSettings_GlobalSettings() *QWebEngineSettings {
	defer qt.Recovering("QWebEngineSettings::globalSettings")

	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_GlobalSettings())
}

func (ptr *QWebEngineSettings) GlobalSettings() *QWebEngineSettings {
	defer qt.Recovering("QWebEngineSettings::globalSettings")

	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_GlobalSettings())
}

func (ptr *QWebEngineSettings) ResetAttribute(attribute QWebEngineSettings__WebAttribute) {
	defer qt.Recovering("QWebEngineSettings::resetAttribute")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetAttribute(ptr.Pointer(), C.int(attribute))
	}
}

func (ptr *QWebEngineSettings) ResetFontFamily(which QWebEngineSettings__FontFamily) {
	defer qt.Recovering("QWebEngineSettings::resetFontFamily")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetFontFamily(ptr.Pointer(), C.int(which))
	}
}

func (ptr *QWebEngineSettings) ResetFontSize(ty QWebEngineSettings__FontSize) {
	defer qt.Recovering("QWebEngineSettings::resetFontSize")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetFontSize(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QWebEngineSettings) SetAttribute(attribute QWebEngineSettings__WebAttribute, on bool) {
	defer qt.Recovering("QWebEngineSettings::setAttribute")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetAttribute(ptr.Pointer(), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWebEngineSettings) SetDefaultTextEncoding(encoding string) {
	defer qt.Recovering("QWebEngineSettings::setDefaultTextEncoding")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetDefaultTextEncoding(ptr.Pointer(), C.CString(encoding))
	}
}

func (ptr *QWebEngineSettings) SetFontFamily(which QWebEngineSettings__FontFamily, family string) {
	defer qt.Recovering("QWebEngineSettings::setFontFamily")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetFontFamily(ptr.Pointer(), C.int(which), C.CString(family))
	}
}

func (ptr *QWebEngineSettings) SetFontSize(ty QWebEngineSettings__FontSize, size int) {
	defer qt.Recovering("QWebEngineSettings::setFontSize")

	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetFontSize(ptr.Pointer(), C.int(ty), C.int(size))
	}
}

func (ptr *QWebEngineSettings) TestAttribute(attribute QWebEngineSettings__WebAttribute) bool {
	defer qt.Recovering("QWebEngineSettings::testAttribute")

	if ptr.Pointer() != nil {
		return C.QWebEngineSettings_TestAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func QWebEngineSettings_DefaultSettings() *QWebEngineSettings {
	defer qt.Recovering("QWebEngineSettings::defaultSettings")

	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_DefaultSettings())
}

func (ptr *QWebEngineSettings) DefaultSettings() *QWebEngineSettings {
	defer qt.Recovering("QWebEngineSettings::defaultSettings")

	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_DefaultSettings())
}

//QWebEngineUrlRequestInfo::NavigationType
type QWebEngineUrlRequestInfo__NavigationType int64

const (
	QWebEngineUrlRequestInfo__NavigationTypeLink          = QWebEngineUrlRequestInfo__NavigationType(0)
	QWebEngineUrlRequestInfo__NavigationTypeTyped         = QWebEngineUrlRequestInfo__NavigationType(1)
	QWebEngineUrlRequestInfo__NavigationTypeFormSubmitted = QWebEngineUrlRequestInfo__NavigationType(2)
	QWebEngineUrlRequestInfo__NavigationTypeBackForward   = QWebEngineUrlRequestInfo__NavigationType(3)
	QWebEngineUrlRequestInfo__NavigationTypeReload        = QWebEngineUrlRequestInfo__NavigationType(4)
	QWebEngineUrlRequestInfo__NavigationTypeOther         = QWebEngineUrlRequestInfo__NavigationType(5)
)

//QWebEngineUrlRequestInfo::ResourceType
type QWebEngineUrlRequestInfo__ResourceType int64

const (
	QWebEngineUrlRequestInfo__ResourceTypeMainFrame     = QWebEngineUrlRequestInfo__ResourceType(0)
	QWebEngineUrlRequestInfo__ResourceTypeSubFrame      = QWebEngineUrlRequestInfo__ResourceType(1)
	QWebEngineUrlRequestInfo__ResourceTypeStylesheet    = QWebEngineUrlRequestInfo__ResourceType(2)
	QWebEngineUrlRequestInfo__ResourceTypeScript        = QWebEngineUrlRequestInfo__ResourceType(3)
	QWebEngineUrlRequestInfo__ResourceTypeImage         = QWebEngineUrlRequestInfo__ResourceType(4)
	QWebEngineUrlRequestInfo__ResourceTypeFontResource  = QWebEngineUrlRequestInfo__ResourceType(5)
	QWebEngineUrlRequestInfo__ResourceTypeSubResource   = QWebEngineUrlRequestInfo__ResourceType(6)
	QWebEngineUrlRequestInfo__ResourceTypeObject        = QWebEngineUrlRequestInfo__ResourceType(7)
	QWebEngineUrlRequestInfo__ResourceTypeMedia         = QWebEngineUrlRequestInfo__ResourceType(8)
	QWebEngineUrlRequestInfo__ResourceTypeWorker        = QWebEngineUrlRequestInfo__ResourceType(9)
	QWebEngineUrlRequestInfo__ResourceTypeSharedWorker  = QWebEngineUrlRequestInfo__ResourceType(10)
	QWebEngineUrlRequestInfo__ResourceTypePrefetch      = QWebEngineUrlRequestInfo__ResourceType(11)
	QWebEngineUrlRequestInfo__ResourceTypeFavicon       = QWebEngineUrlRequestInfo__ResourceType(12)
	QWebEngineUrlRequestInfo__ResourceTypeXhr           = QWebEngineUrlRequestInfo__ResourceType(13)
	QWebEngineUrlRequestInfo__ResourceTypePing          = QWebEngineUrlRequestInfo__ResourceType(14)
	QWebEngineUrlRequestInfo__ResourceTypeServiceWorker = QWebEngineUrlRequestInfo__ResourceType(15)
	QWebEngineUrlRequestInfo__ResourceTypeUnknown       = QWebEngineUrlRequestInfo__ResourceType(16)
)

type QWebEngineUrlRequestInfo struct {
	ptr unsafe.Pointer
}

type QWebEngineUrlRequestInfo_ITF interface {
	QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo
}

func (p *QWebEngineUrlRequestInfo) QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo {
	return p
}

func (p *QWebEngineUrlRequestInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebEngineUrlRequestInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQWebEngineUrlRequestInfoFromPointer(ptr unsafe.Pointer) *QWebEngineUrlRequestInfo {
	var n = NewQWebEngineUrlRequestInfoFromPointer(ptr)
	return n
}

func (ptr *QWebEngineUrlRequestInfo) Block(shouldBlock bool) {
	defer qt.Recovering("QWebEngineUrlRequestInfo::block")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_Block(ptr.Pointer(), C.int(qt.GoBoolToInt(shouldBlock)))
	}
}

func (ptr *QWebEngineUrlRequestInfo) FirstPartyUrl() *core.QUrl {
	defer qt.Recovering("QWebEngineUrlRequestInfo::firstPartyUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_FirstPartyUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) NavigationType() QWebEngineUrlRequestInfo__NavigationType {
	defer qt.Recovering("QWebEngineUrlRequestInfo::navigationType")

	if ptr.Pointer() != nil {
		return QWebEngineUrlRequestInfo__NavigationType(C.QWebEngineUrlRequestInfo_NavigationType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineUrlRequestInfo) Redirect(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInfo::redirect")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_Redirect(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineUrlRequestInfo) RequestMethod() string {
	defer qt.Recovering("QWebEngineUrlRequestInfo::requestMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineUrlRequestInfo_RequestMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineUrlRequestInfo) RequestUrl() *core.QUrl {
	defer qt.Recovering("QWebEngineUrlRequestInfo::requestUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_RequestUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) ResourceType() QWebEngineUrlRequestInfo__ResourceType {
	defer qt.Recovering("QWebEngineUrlRequestInfo::resourceType")

	if ptr.Pointer() != nil {
		return QWebEngineUrlRequestInfo__ResourceType(C.QWebEngineUrlRequestInfo_ResourceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineUrlRequestInfo) SetHttpHeader(name string, value string) {
	defer qt.Recovering("QWebEngineUrlRequestInfo::setHttpHeader")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_SetHttpHeader(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

type QWebEngineUrlRequestInterceptor struct {
	core.QObject
}

type QWebEngineUrlRequestInterceptor_ITF interface {
	core.QObject_ITF
	QWebEngineUrlRequestInterceptor_PTR() *QWebEngineUrlRequestInterceptor
}

func (p *QWebEngineUrlRequestInterceptor) QWebEngineUrlRequestInterceptor_PTR() *QWebEngineUrlRequestInterceptor {
	return p
}

func (p *QWebEngineUrlRequestInterceptor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineUrlRequestInterceptor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEngineUrlRequestInterceptorFromPointer(ptr unsafe.Pointer) *QWebEngineUrlRequestInterceptor {
	var n = NewQWebEngineUrlRequestInterceptorFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineUrlRequestInterceptor_") {
		n.SetObjectName("QWebEngineUrlRequestInterceptor_" + qt.Identifier())
	}
	return n
}

//export callbackQWebEngineUrlRequestInterceptor_TimerEvent
func callbackQWebEngineUrlRequestInterceptor_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_ChildEvent
func callbackQWebEngineUrlRequestInterceptor_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_ConnectNotify
func callbackQWebEngineUrlRequestInterceptor_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_CustomEvent
func callbackQWebEngineUrlRequestInterceptor_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_DeleteLater
func callbackQWebEngineUrlRequestInterceptor_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DeleteLater() {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineUrlRequestInterceptor_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineUrlRequestInterceptor_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlRequestInterceptor_DisconnectNotify
func callbackQWebEngineUrlRequestInterceptor_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_Event
func callbackQWebEngineUrlRequestInterceptor_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestInterceptor) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestInterceptor_EventFilter
func callbackQWebEngineUrlRequestInterceptor_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestInterceptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestInterceptor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestInterceptor_MetaObject
func callbackQWebEngineUrlRequestInterceptor_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineUrlRequestInterceptor::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineUrlRequestInterceptor::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestInterceptor::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestInterceptor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineUrlRequestInterceptor::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestInterceptor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebEngineUrlRequestJob::Error
type QWebEngineUrlRequestJob__Error int64

const (
	QWebEngineUrlRequestJob__NoError        = QWebEngineUrlRequestJob__Error(0)
	QWebEngineUrlRequestJob__UrlNotFound    = QWebEngineUrlRequestJob__Error(1)
	QWebEngineUrlRequestJob__UrlInvalid     = QWebEngineUrlRequestJob__Error(2)
	QWebEngineUrlRequestJob__RequestAborted = QWebEngineUrlRequestJob__Error(3)
	QWebEngineUrlRequestJob__RequestDenied  = QWebEngineUrlRequestJob__Error(4)
	QWebEngineUrlRequestJob__RequestFailed  = QWebEngineUrlRequestJob__Error(5)
)

type QWebEngineUrlRequestJob struct {
	core.QObject
}

type QWebEngineUrlRequestJob_ITF interface {
	core.QObject_ITF
	QWebEngineUrlRequestJob_PTR() *QWebEngineUrlRequestJob
}

func (p *QWebEngineUrlRequestJob) QWebEngineUrlRequestJob_PTR() *QWebEngineUrlRequestJob {
	return p
}

func (p *QWebEngineUrlRequestJob) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineUrlRequestJob) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEngineUrlRequestJobFromPointer(ptr unsafe.Pointer) *QWebEngineUrlRequestJob {
	var n = NewQWebEngineUrlRequestJobFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineUrlRequestJob_") {
		n.SetObjectName("QWebEngineUrlRequestJob_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebEngineUrlRequestJob) Fail(r QWebEngineUrlRequestJob__Error) {
	defer qt.Recovering("QWebEngineUrlRequestJob::fail")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_Fail(ptr.Pointer(), C.int(r))
	}
}

func (ptr *QWebEngineUrlRequestJob) Redirect(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::redirect")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_Redirect(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineUrlRequestJob) Reply(contentType string, device core.QIODevice_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::reply")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_Reply(ptr.Pointer(), C.CString(contentType), core.PointerFromQIODevice(device))
	}
}

func (ptr *QWebEngineUrlRequestJob) RequestMethod() string {
	defer qt.Recovering("QWebEngineUrlRequestJob::requestMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineUrlRequestJob_RequestMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineUrlRequestJob) RequestUrl() *core.QUrl {
	defer qt.Recovering("QWebEngineUrlRequestJob::requestUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineUrlRequestJob_RequestUrl(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineUrlRequestJob_TimerEvent
func callbackQWebEngineUrlRequestJob_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineUrlRequestJob) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_ChildEvent
func callbackQWebEngineUrlRequestJob_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineUrlRequestJob) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_ConnectNotify
func callbackQWebEngineUrlRequestJob_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestJob_CustomEvent
func callbackQWebEngineUrlRequestJob_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineUrlRequestJob) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_DeleteLater
func callbackQWebEngineUrlRequestJob_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineUrlRequestJob) DeleteLater() {
	defer qt.Recovering("QWebEngineUrlRequestJob::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlRequestJob) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineUrlRequestJob::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlRequestJob_DisconnectNotify
func callbackQWebEngineUrlRequestJob_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlRequestJob::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestJob_Event
func callbackQWebEngineUrlRequestJob_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineUrlRequestJobFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEngineUrlRequestJob) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineUrlRequestJob) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestJob::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestJob) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestJob::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestJob_EventFilter
func callbackQWebEngineUrlRequestJob_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineUrlRequestJobFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineUrlRequestJob) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineUrlRequestJob) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestJob::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlRequestJob) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlRequestJob::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlRequestJob_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestJob_MetaObject
func callbackQWebEngineUrlRequestJob_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineUrlRequestJob::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlRequestJobFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlRequestJob) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineUrlRequestJob::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineUrlRequestJob::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineUrlRequestJob) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineUrlRequestJob::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestJob_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineUrlRequestJob::metaObject")

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

func (p *QWebEngineUrlSchemeHandler) QWebEngineUrlSchemeHandler_PTR() *QWebEngineUrlSchemeHandler {
	return p
}

func (p *QWebEngineUrlSchemeHandler) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineUrlSchemeHandler) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQWebEngineUrlSchemeHandlerFromPointer(ptr unsafe.Pointer) *QWebEngineUrlSchemeHandler {
	var n = NewQWebEngineUrlSchemeHandlerFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineUrlSchemeHandler_") {
		n.SetObjectName("QWebEngineUrlSchemeHandler_" + qt.Identifier())
	}
	return n
}

func NewQWebEngineUrlSchemeHandler(parent core.QObject_ITF) *QWebEngineUrlSchemeHandler {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::QWebEngineUrlSchemeHandler")

	return newQWebEngineUrlSchemeHandlerFromPointer(C.QWebEngineUrlSchemeHandler_NewQWebEngineUrlSchemeHandler(core.PointerFromQObject(parent)))
}

//export callbackQWebEngineUrlSchemeHandler_RequestStarted
func callbackQWebEngineUrlSchemeHandler_RequestStarted(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::requestStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestStarted"); signal != nil {
		signal.(func(*QWebEngineUrlRequestJob))(NewQWebEngineUrlRequestJobFromPointer(request))
	}

}

func (ptr *QWebEngineUrlSchemeHandler) ConnectRequestStarted(f func(request *QWebEngineUrlRequestJob)) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::requestStarted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestStarted", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectRequestStarted(request QWebEngineUrlRequestJob_ITF) {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::requestStarted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestStarted")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) RequestStarted(request QWebEngineUrlRequestJob_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::requestStarted")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_RequestStarted(ptr.Pointer(), PointerFromQWebEngineUrlRequestJob(request))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandler() {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::~QWebEngineUrlSchemeHandler")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlSchemeHandler_TimerEvent
func callbackQWebEngineUrlSchemeHandler_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_ChildEvent
func callbackQWebEngineUrlSchemeHandler_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_ConnectNotify
func callbackQWebEngineUrlSchemeHandler_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlSchemeHandler_CustomEvent
func callbackQWebEngineUrlSchemeHandler_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_DeleteLater
func callbackQWebEngineUrlSchemeHandler_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DeleteLater() {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineUrlSchemeHandler_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineUrlSchemeHandler_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineUrlSchemeHandler_DisconnectNotify
func callbackQWebEngineUrlSchemeHandler_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlSchemeHandler_Event
func callbackQWebEngineUrlSchemeHandler_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlSchemeHandler) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebEngineUrlSchemeHandler_EventFilter
func callbackQWebEngineUrlSchemeHandler_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineUrlSchemeHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineUrlSchemeHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineUrlSchemeHandler_MetaObject
func callbackQWebEngineUrlSchemeHandler_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineUrlSchemeHandler::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineUrlSchemeHandler::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineUrlSchemeHandler::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlSchemeHandler_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineUrlSchemeHandler::metaObject")

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

func (p *QWebEngineView) QWebEngineView_PTR() *QWebEngineView {
	return p
}

func (p *QWebEngineView) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QWebEngineView) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
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

func newQWebEngineViewFromPointer(ptr unsafe.Pointer) *QWebEngineView {
	var n = NewQWebEngineViewFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebEngineView_") {
		n.SetObjectName("QWebEngineView_" + qt.Identifier())
	}
	return n
}

func NewQWebEngineView(parent widgets.QWidget_ITF) *QWebEngineView {
	defer qt.Recovering("QWebEngineView::QWebEngineView")

	return newQWebEngineViewFromPointer(C.QWebEngineView_NewQWebEngineView(widgets.PointerFromQWidget(parent)))
}

//export callbackQWebEngineView_Back
func callbackQWebEngineView_Back(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::back")

	if signal := qt.GetSignal(C.GoString(ptrName), "back"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectBack(f func()) {
	defer qt.Recovering("connect QWebEngineView::back")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "back", f)
	}
}

func (ptr *QWebEngineView) DisconnectBack() {
	defer qt.Recovering("disconnect QWebEngineView::back")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "back")
	}
}

func (ptr *QWebEngineView) Back() {
	defer qt.Recovering("QWebEngineView::back")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Back(ptr.Pointer())
	}
}

//export callbackQWebEngineView_CreateWindow
func callbackQWebEngineView_CreateWindow(ptr unsafe.Pointer, ptrName *C.char, ty C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineView::createWindow")

	if signal := qt.GetSignal(C.GoString(ptrName), "createWindow"); signal != nil {
		return PointerFromQWebEngineView(signal.(func(QWebEnginePage__WebWindowType) *QWebEngineView)(QWebEnginePage__WebWindowType(ty)))
	}

	return PointerFromQWebEngineView(NewQWebEngineViewFromPointer(ptr).CreateWindowDefault(QWebEnginePage__WebWindowType(ty)))
}

func (ptr *QWebEngineView) ConnectCreateWindow(f func(ty QWebEnginePage__WebWindowType) *QWebEngineView) {
	defer qt.Recovering("connect QWebEngineView::createWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createWindow", f)
	}
}

func (ptr *QWebEngineView) DisconnectCreateWindow() {
	defer qt.Recovering("disconnect QWebEngineView::createWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createWindow")
	}
}

func (ptr *QWebEngineView) CreateWindow(ty QWebEnginePage__WebWindowType) *QWebEngineView {
	defer qt.Recovering("QWebEngineView::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebEngineViewFromPointer(C.QWebEngineView_CreateWindow(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QWebEngineView) CreateWindowDefault(ty QWebEnginePage__WebWindowType) *QWebEngineView {
	defer qt.Recovering("QWebEngineView::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebEngineViewFromPointer(C.QWebEngineView_CreateWindowDefault(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

//export callbackQWebEngineView_Forward
func callbackQWebEngineView_Forward(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::forward")

	if signal := qt.GetSignal(C.GoString(ptrName), "forward"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectForward(f func()) {
	defer qt.Recovering("connect QWebEngineView::forward")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "forward", f)
	}
}

func (ptr *QWebEngineView) DisconnectForward() {
	defer qt.Recovering("disconnect QWebEngineView::forward")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "forward")
	}
}

func (ptr *QWebEngineView) Forward() {
	defer qt.Recovering("QWebEngineView::forward")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Forward(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) HasSelection() bool {
	defer qt.Recovering("QWebEngineView::hasSelection")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineView) History() *QWebEngineHistory {
	defer qt.Recovering("QWebEngineView::history")

	if ptr.Pointer() != nil {
		return NewQWebEngineHistoryFromPointer(C.QWebEngineView_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) IconUrl() *core.QUrl {
	defer qt.Recovering("QWebEngineView::iconUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineView_IconUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineView::load")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineView) Page() *QWebEnginePage {
	defer qt.Recovering("QWebEngineView::page")

	if ptr.Pointer() != nil {
		return NewQWebEnginePageFromPointer(C.QWebEngineView_Page(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) PageAction(action QWebEnginePage__WebAction) *widgets.QAction {
	defer qt.Recovering("QWebEngineView::pageAction")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QWebEngineView_PageAction(ptr.Pointer(), C.int(action)))
	}
	return nil
}

//export callbackQWebEngineView_Reload
func callbackQWebEngineView_Reload(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::reload")

	if signal := qt.GetSignal(C.GoString(ptrName), "reload"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectReload(f func()) {
	defer qt.Recovering("connect QWebEngineView::reload")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reload", f)
	}
}

func (ptr *QWebEngineView) DisconnectReload() {
	defer qt.Recovering("disconnect QWebEngineView::reload")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reload")
	}
}

func (ptr *QWebEngineView) Reload() {
	defer qt.Recovering("QWebEngineView::reload")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Reload(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) SelectedText() string {
	defer qt.Recovering("QWebEngineView::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineView_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineView) SetContent(data string, mimeType string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineView::setContent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetContent(ptr.Pointer(), C.CString(data), C.CString(mimeType), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEngineView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineView::setHtml")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetHtml(ptr.Pointer(), C.CString(html), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEngineView) SetPage(page QWebEnginePage_ITF) {
	defer qt.Recovering("QWebEngineView::setPage")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetPage(ptr.Pointer(), PointerFromQWebEnginePage(page))
	}
}

func (ptr *QWebEngineView) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineView::setUrl")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineView) SetZoomFactor(factor float64) {
	defer qt.Recovering("QWebEngineView::setZoomFactor")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebEngineView) Settings() *QWebEngineSettings {
	defer qt.Recovering("QWebEngineView::settings")

	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEngineView_Settings(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineView_Stop
func callbackQWebEngineView_Stop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectStop(f func()) {
	defer qt.Recovering("connect QWebEngineView::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QWebEngineView) DisconnectStop() {
	defer qt.Recovering("disconnect QWebEngineView::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

func (ptr *QWebEngineView) Stop() {
	defer qt.Recovering("QWebEngineView::stop")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Stop(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) Title() string {
	defer qt.Recovering("QWebEngineView::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebEngineView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineView) TriggerPageAction(action QWebEnginePage__WebAction, checked bool) {
	defer qt.Recovering("QWebEngineView::triggerPageAction")

	if ptr.Pointer() != nil {
		C.QWebEngineView_TriggerPageAction(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QWebEngineView) Url() *core.QUrl {
	defer qt.Recovering("QWebEngineView::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebEngineView_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) ZoomFactor() float64 {
	defer qt.Recovering("QWebEngineView::zoomFactor")

	if ptr.Pointer() != nil {
		return float64(C.QWebEngineView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineView) DestroyQWebEngineView() {
	defer qt.Recovering("QWebEngineView::~QWebEngineView")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineView_DestroyQWebEngineView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineView_ContextMenuEvent
func callbackQWebEngineView_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QWebEngineView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QWebEngineView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *QWebEngineView) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWebEngineView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWebEngineView) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWebEngineView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQWebEngineView_Event
func callbackQWebEngineView_Event(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineView::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev))))
}

func (ptr *QWebEngineView) ConnectEvent(f func(ev *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineView::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebEngineView) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebEngineView::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebEngineView) Event(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineView::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebEngineView) EventDefault(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineView::event")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebEngineView) FindText(subString string, options QWebEnginePage__FindFlag) {
	defer qt.Recovering("QWebEngineView::findText")

	if ptr.Pointer() != nil {
		C.QWebEngineView_FindText(ptr.Pointer(), C.CString(subString), C.int(options))
	}
}

//export callbackQWebEngineView_HideEvent
func callbackQWebEngineView_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QWebEngineView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWebEngineView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *QWebEngineView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWebEngineView::hideEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWebEngineView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWebEngineView::hideEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWebEngineView_IconUrlChanged
func callbackQWebEngineView_IconUrlChanged(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::iconUrlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconUrlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEngineView) ConnectIconUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebEngineView::iconUrlChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectIconUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconUrlChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectIconUrlChanged() {
	defer qt.Recovering("disconnect QWebEngineView::iconUrlChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectIconUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconUrlChanged")
	}
}

func (ptr *QWebEngineView) IconUrlChanged(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineView::iconUrlChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_IconUrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebEngineView_LoadFinished
func callbackQWebEngineView_LoadFinished(ptr unsafe.Pointer, ptrName *C.char, ok C.int) {
	defer qt.Recovering("callback QWebEngineView::loadFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFinished"); signal != nil {
		signal.(func(bool))(int(ok) != 0)
	}

}

func (ptr *QWebEngineView) ConnectLoadFinished(f func(ok bool)) {
	defer qt.Recovering("connect QWebEngineView::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFinished", f)
	}
}

func (ptr *QWebEngineView) DisconnectLoadFinished() {
	defer qt.Recovering("disconnect QWebEngineView::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFinished")
	}
}

func (ptr *QWebEngineView) LoadFinished(ok bool) {
	defer qt.Recovering("QWebEngineView::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadFinished(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)))
	}
}

//export callbackQWebEngineView_LoadProgress
func callbackQWebEngineView_LoadProgress(ptr unsafe.Pointer, ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QWebEngineView::loadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadProgress"); signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QWebEngineView) ConnectLoadProgress(f func(progress int)) {
	defer qt.Recovering("connect QWebEngineView::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadProgress", f)
	}
}

func (ptr *QWebEngineView) DisconnectLoadProgress() {
	defer qt.Recovering("disconnect QWebEngineView::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadProgress")
	}
}

func (ptr *QWebEngineView) LoadProgress(progress int) {
	defer qt.Recovering("QWebEngineView::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadProgress(ptr.Pointer(), C.int(progress))
	}
}

//export callbackQWebEngineView_LoadStarted
func callbackQWebEngineView_LoadStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::loadStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectLoadStarted(f func()) {
	defer qt.Recovering("connect QWebEngineView::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadStarted", f)
	}
}

func (ptr *QWebEngineView) DisconnectLoadStarted() {
	defer qt.Recovering("disconnect QWebEngineView::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadStarted")
	}
}

func (ptr *QWebEngineView) LoadStarted() {
	defer qt.Recovering("QWebEngineView::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebEngineView_RenderProcessTerminated
func callbackQWebEngineView_RenderProcessTerminated(ptr unsafe.Pointer, ptrName *C.char, terminationStatus C.int, exitCode C.int) {
	defer qt.Recovering("callback QWebEngineView::renderProcessTerminated")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderProcessTerminated"); signal != nil {
		signal.(func(QWebEnginePage__RenderProcessTerminationStatus, int))(QWebEnginePage__RenderProcessTerminationStatus(terminationStatus), int(exitCode))
	}

}

func (ptr *QWebEngineView) ConnectRenderProcessTerminated(f func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int)) {
	defer qt.Recovering("connect QWebEngineView::renderProcessTerminated")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectRenderProcessTerminated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderProcessTerminated", f)
	}
}

func (ptr *QWebEngineView) DisconnectRenderProcessTerminated() {
	defer qt.Recovering("disconnect QWebEngineView::renderProcessTerminated")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectRenderProcessTerminated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderProcessTerminated")
	}
}

func (ptr *QWebEngineView) RenderProcessTerminated(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int) {
	defer qt.Recovering("QWebEngineView::renderProcessTerminated")

	if ptr.Pointer() != nil {
		C.QWebEngineView_RenderProcessTerminated(ptr.Pointer(), C.int(terminationStatus), C.int(exitCode))
	}
}

//export callbackQWebEngineView_SelectionChanged
func callbackQWebEngineView_SelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebEngineView) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QWebEngineView::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QWebEngineView::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

func (ptr *QWebEngineView) SelectionChanged() {
	defer qt.Recovering("QWebEngineView::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SelectionChanged(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowEvent
func callbackQWebEngineView_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QWebEngineView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWebEngineView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *QWebEngineView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWebEngineView::showEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWebEngineView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWebEngineView::showEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWebEngineView_SizeHint
func callbackQWebEngineView_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineView::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebEngineViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebEngineView) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QWebEngineView::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QWebEngineView) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QWebEngineView::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QWebEngineView) SizeHint() *core.QSize {
	defer qt.Recovering("QWebEngineView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebEngineView_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QWebEngineView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebEngineView_SizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineView_TitleChanged
func callbackQWebEngineView_TitleChanged(ptr unsafe.Pointer, ptrName *C.char, title *C.char) {
	defer qt.Recovering("callback QWebEngineView::titleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "titleChanged"); signal != nil {
		signal.(func(string))(C.GoString(title))
	}

}

func (ptr *QWebEngineView) ConnectTitleChanged(f func(title string)) {
	defer qt.Recovering("connect QWebEngineView::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "titleChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectTitleChanged() {
	defer qt.Recovering("disconnect QWebEngineView::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "titleChanged")
	}
}

func (ptr *QWebEngineView) TitleChanged(title string) {
	defer qt.Recovering("QWebEngineView::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_TitleChanged(ptr.Pointer(), C.CString(title))
	}
}

//export callbackQWebEngineView_UrlChanged
func callbackQWebEngineView_UrlChanged(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::urlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEngineView) ConnectUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebEngineView::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "urlChanged", f)
	}
}

func (ptr *QWebEngineView) DisconnectUrlChanged() {
	defer qt.Recovering("disconnect QWebEngineView::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "urlChanged")
	}
}

func (ptr *QWebEngineView) UrlChanged(url core.QUrl_ITF) {
	defer qt.Recovering("QWebEngineView::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebEngineView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebEngineView_ActionEvent
func callbackQWebEngineView_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QWebEngineView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QWebEngineView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

func (ptr *QWebEngineView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWebEngineView::actionEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWebEngineView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWebEngineView::actionEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebEngineView_DragEnterEvent
func callbackQWebEngineView_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QWebEngineView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QWebEngineView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *QWebEngineView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QWebEngineView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQWebEngineView_DragLeaveEvent
func callbackQWebEngineView_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QWebEngineView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QWebEngineView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *QWebEngineView) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QWebEngineView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQWebEngineView_DragMoveEvent
func callbackQWebEngineView_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QWebEngineView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QWebEngineView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *QWebEngineView) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QWebEngineView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQWebEngineView_DropEvent
func callbackQWebEngineView_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QWebEngineView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QWebEngineView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *QWebEngineView) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dropEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QWebEngineView) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWebEngineView::dropEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQWebEngineView_EnterEvent
func callbackQWebEngineView_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QWebEngineView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

func (ptr *QWebEngineView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::enterEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::enterEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_FocusInEvent
func callbackQWebEngineView_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWebEngineView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QWebEngineView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *QWebEngineView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebEngineView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWebEngineView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebEngineView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebEngineView_FocusOutEvent
func callbackQWebEngineView_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWebEngineView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWebEngineView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *QWebEngineView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebEngineView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWebEngineView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebEngineView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebEngineView_LeaveEvent
func callbackQWebEngineView_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QWebEngineView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

func (ptr *QWebEngineView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_MinimumSizeHint
func callbackQWebEngineView_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineView::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebEngineViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebEngineView) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QWebEngineView::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *QWebEngineView) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QWebEngineView::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *QWebEngineView) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QWebEngineView::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebEngineView_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QWebEngineView::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebEngineView_MinimumSizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineView_MoveEvent
func callbackQWebEngineView_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QWebEngineView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWebEngineView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *QWebEngineView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWebEngineView::moveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWebEngineView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWebEngineView::moveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebEngineView_PaintEvent
func callbackQWebEngineView_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QWebEngineView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QWebEngineView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

func (ptr *QWebEngineView) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWebEngineView::paintEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWebEngineView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWebEngineView::paintEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQWebEngineView_SetEnabled
func callbackQWebEngineView_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QWebEngineView::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QWebEngineView::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QWebEngineView::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *QWebEngineView) SetEnabled(vbo bool) {
	defer qt.Recovering("QWebEngineView::setEnabled")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QWebEngineView) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QWebEngineView::setEnabled")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQWebEngineView_SetStyleSheet
func callbackQWebEngineView_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback QWebEngineView::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQWebEngineViewFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QWebEngineView) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QWebEngineView::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QWebEngineView::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *QWebEngineView) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QWebEngineView::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QWebEngineView) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QWebEngineView::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetStyleSheetDefault(ptr.Pointer(), C.CString(styleSheet))
	}
}

//export callbackQWebEngineView_SetVisible
func callbackQWebEngineView_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWebEngineView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QWebEngineView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QWebEngineView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *QWebEngineView) SetVisible(visible bool) {
	defer qt.Recovering("QWebEngineView::setVisible")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWebEngineView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QWebEngineView::setVisible")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQWebEngineView_SetWindowModified
func callbackQWebEngineView_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QWebEngineView::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QWebEngineView::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QWebEngineView::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *QWebEngineView) SetWindowModified(vbo bool) {
	defer qt.Recovering("QWebEngineView::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QWebEngineView) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QWebEngineView::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQWebEngineView_SetWindowTitle
func callbackQWebEngineView_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QWebEngineView::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQWebEngineViewFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QWebEngineView) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QWebEngineView::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QWebEngineView::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *QWebEngineView) SetWindowTitle(vqs string) {
	defer qt.Recovering("QWebEngineView::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QWebEngineView) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QWebEngineView::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQWebEngineView_ChangeEvent
func callbackQWebEngineView_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QWebEngineView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *QWebEngineView) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::changeEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::changeEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_Close
func callbackQWebEngineView_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebEngineView::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).CloseDefault()))
}

func (ptr *QWebEngineView) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QWebEngineView::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QWebEngineView) DisconnectClose() {
	defer qt.Recovering("disconnect QWebEngineView::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QWebEngineView) Close() bool {
	defer qt.Recovering("QWebEngineView::close")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineView) CloseDefault() bool {
	defer qt.Recovering("QWebEngineView::close")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebEngineView_CloseEvent
func callbackQWebEngineView_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QWebEngineView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QWebEngineView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *QWebEngineView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::closeEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWebEngineView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::closeEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebEngineView_FocusNextPrevChild
func callbackQWebEngineView_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QWebEngineView::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QWebEngineView) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QWebEngineView::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QWebEngineView) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QWebEngineView::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QWebEngineView) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QWebEngineView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QWebEngineView) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QWebEngineView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQWebEngineView_HasHeightForWidth
func callbackQWebEngineView_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebEngineView::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *QWebEngineView) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QWebEngineView::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *QWebEngineView) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QWebEngineView::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *QWebEngineView) HasHeightForWidth() bool {
	defer qt.Recovering("QWebEngineView::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebEngineView) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QWebEngineView::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebEngineView_HeightForWidth
func callbackQWebEngineView_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback QWebEngineView::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewQWebEngineViewFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *QWebEngineView) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QWebEngineView::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *QWebEngineView) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QWebEngineView::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *QWebEngineView) HeightForWidth(w int) int {
	defer qt.Recovering("QWebEngineView::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineView_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWebEngineView) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QWebEngineView::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWebEngineView_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackQWebEngineView_Hide
func callbackQWebEngineView_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebEngineView) ConnectHide(f func()) {
	defer qt.Recovering("connect QWebEngineView::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QWebEngineView) DisconnectHide() {
	defer qt.Recovering("disconnect QWebEngineView::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QWebEngineView) Hide() {
	defer qt.Recovering("QWebEngineView::hide")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Hide(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) HideDefault() {
	defer qt.Recovering("QWebEngineView::hide")

	if ptr.Pointer() != nil {
		C.QWebEngineView_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_InputMethodEvent
func callbackQWebEngineView_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QWebEngineView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QWebEngineView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *QWebEngineView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWebEngineView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWebEngineView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWebEngineView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQWebEngineView_InputMethodQuery
func callbackQWebEngineView_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineView::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQWebEngineViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QWebEngineView) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QWebEngineView::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QWebEngineView) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QWebEngineView::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QWebEngineView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebEngineView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebEngineView_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QWebEngineView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebEngineView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebEngineView_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQWebEngineView_KeyPressEvent
func callbackQWebEngineView_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWebEngineView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWebEngineView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *QWebEngineView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebEngineView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWebEngineView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebEngineView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebEngineView_KeyReleaseEvent
func callbackQWebEngineView_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWebEngineView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWebEngineView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *QWebEngineView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebEngineView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWebEngineView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebEngineView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebEngineView_Lower
func callbackQWebEngineView_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebEngineView) ConnectLower(f func()) {
	defer qt.Recovering("connect QWebEngineView::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QWebEngineView) DisconnectLower() {
	defer qt.Recovering("disconnect QWebEngineView::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QWebEngineView) Lower() {
	defer qt.Recovering("QWebEngineView::lower")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Lower(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) LowerDefault() {
	defer qt.Recovering("QWebEngineView::lower")

	if ptr.Pointer() != nil {
		C.QWebEngineView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_MouseDoubleClickEvent
func callbackQWebEngineView_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebEngineView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWebEngineView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *QWebEngineView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MouseMoveEvent
func callbackQWebEngineView_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebEngineView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWebEngineView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *QWebEngineView) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MousePressEvent
func callbackQWebEngineView_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebEngineView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWebEngineView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *QWebEngineView) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MouseReleaseEvent
func callbackQWebEngineView_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebEngineView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWebEngineView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *QWebEngineView) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebEngineView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebEngineView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_NativeEvent
func callbackQWebEngineView_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QWebEngineView::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QWebEngineView) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QWebEngineView::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QWebEngineView::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QWebEngineView) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QWebEngineView::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QWebEngineView) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QWebEngineView::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQWebEngineView_Raise
func callbackQWebEngineView_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebEngineView) ConnectRaise(f func()) {
	defer qt.Recovering("connect QWebEngineView::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QWebEngineView) DisconnectRaise() {
	defer qt.Recovering("disconnect QWebEngineView::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QWebEngineView) Raise() {
	defer qt.Recovering("QWebEngineView::raise")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Raise(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) RaiseDefault() {
	defer qt.Recovering("QWebEngineView::raise")

	if ptr.Pointer() != nil {
		C.QWebEngineView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_Repaint
func callbackQWebEngineView_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebEngineView) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QWebEngineView::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *QWebEngineView) DisconnectRepaint() {
	defer qt.Recovering("disconnect QWebEngineView::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *QWebEngineView) Repaint() {
	defer qt.Recovering("QWebEngineView::repaint")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Repaint(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) RepaintDefault() {
	defer qt.Recovering("QWebEngineView::repaint")

	if ptr.Pointer() != nil {
		C.QWebEngineView_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ResizeEvent
func callbackQWebEngineView_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QWebEngineView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWebEngineView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *QWebEngineView) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWebEngineView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWebEngineView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWebEngineView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQWebEngineView_SetDisabled
func callbackQWebEngineView_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback QWebEngineView::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QWebEngineView::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QWebEngineView::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *QWebEngineView) SetDisabled(disable bool) {
	defer qt.Recovering("QWebEngineView::setDisabled")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QWebEngineView) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QWebEngineView::setDisabled")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackQWebEngineView_SetFocus2
func callbackQWebEngineView_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebEngineView) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QWebEngineView::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QWebEngineView::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *QWebEngineView) SetFocus2() {
	defer qt.Recovering("QWebEngineView::setFocus")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) SetFocus2Default() {
	defer qt.Recovering("QWebEngineView::setFocus")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebEngineView_SetHidden
func callbackQWebEngineView_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback QWebEngineView::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *QWebEngineView) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QWebEngineView::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *QWebEngineView) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QWebEngineView::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *QWebEngineView) SetHidden(hidden bool) {
	defer qt.Recovering("QWebEngineView::setHidden")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QWebEngineView) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QWebEngineView::setHidden")

	if ptr.Pointer() != nil {
		C.QWebEngineView_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackQWebEngineView_Show
func callbackQWebEngineView_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebEngineView) ConnectShow(f func()) {
	defer qt.Recovering("connect QWebEngineView::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QWebEngineView) DisconnectShow() {
	defer qt.Recovering("disconnect QWebEngineView::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QWebEngineView) Show() {
	defer qt.Recovering("QWebEngineView::show")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Show(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowDefault() {
	defer qt.Recovering("QWebEngineView::show")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowFullScreen
func callbackQWebEngineView_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QWebEngineView::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QWebEngineView::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QWebEngineView) ShowFullScreen() {
	defer qt.Recovering("QWebEngineView::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowFullScreenDefault() {
	defer qt.Recovering("QWebEngineView::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowMaximized
func callbackQWebEngineView_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QWebEngineView::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QWebEngineView::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QWebEngineView) ShowMaximized() {
	defer qt.Recovering("QWebEngineView::showMaximized")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowMaximizedDefault() {
	defer qt.Recovering("QWebEngineView::showMaximized")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowMinimized
func callbackQWebEngineView_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QWebEngineView::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QWebEngineView::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QWebEngineView) ShowMinimized() {
	defer qt.Recovering("QWebEngineView::showMinimized")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowMinimizedDefault() {
	defer qt.Recovering("QWebEngineView::showMinimized")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowNormal
func callbackQWebEngineView_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebEngineView) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QWebEngineView::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QWebEngineView) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QWebEngineView::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QWebEngineView) ShowNormal() {
	defer qt.Recovering("QWebEngineView::showNormal")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) ShowNormalDefault() {
	defer qt.Recovering("QWebEngineView::showNormal")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_TabletEvent
func callbackQWebEngineView_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QWebEngineView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWebEngineView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

func (ptr *QWebEngineView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWebEngineView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWebEngineView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWebEngineView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebEngineView_Update
func callbackQWebEngineView_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebEngineView) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QWebEngineView::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QWebEngineView) DisconnectUpdate() {
	defer qt.Recovering("disconnect QWebEngineView::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QWebEngineView) Update() {
	defer qt.Recovering("QWebEngineView::update")

	if ptr.Pointer() != nil {
		C.QWebEngineView_Update(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) UpdateDefault() {
	defer qt.Recovering("QWebEngineView::update")

	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_UpdateMicroFocus
func callbackQWebEngineView_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebEngineView) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QWebEngineView::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QWebEngineView) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QWebEngineView::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QWebEngineView) UpdateMicroFocus() {
	defer qt.Recovering("QWebEngineView::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) UpdateMicroFocusDefault() {
	defer qt.Recovering("QWebEngineView::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_WheelEvent
func callbackQWebEngineView_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QWebEngineView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWebEngineView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *QWebEngineView) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWebEngineView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWebEngineView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWebEngineView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQWebEngineView_TimerEvent
func callbackQWebEngineView_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebEngineView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebEngineView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebEngineView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineView::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebEngineView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebEngineView::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebEngineView_ChildEvent
func callbackQWebEngineView_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebEngineView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebEngineView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebEngineView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineView::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebEngineView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebEngineView::childEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineView_ConnectNotify
func callbackQWebEngineView_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineView) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineView::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebEngineView) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebEngineView::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebEngineView) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineView::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineView::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineView_CustomEvent
func callbackQWebEngineView_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebEngineView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebEngineView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebEngineView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebEngineView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebEngineView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebEngineView::customEvent")

	if ptr.Pointer() != nil {
		C.QWebEngineView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_DeleteLater
func callbackQWebEngineView_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebEngineView::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebEngineViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineView) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebEngineView::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebEngineView) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebEngineView::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebEngineView) DeleteLater() {
	defer qt.Recovering("QWebEngineView::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineView_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineView) DeleteLaterDefault() {
	defer qt.Recovering("QWebEngineView::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebEngineView_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebEngineView_DisconnectNotify
func callbackQWebEngineView_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebEngineView::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineView) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebEngineView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebEngineView) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebEngineView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebEngineView) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebEngineView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebEngineView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineView_EventFilter
func callbackQWebEngineView_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebEngineView::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebEngineView) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebEngineView::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebEngineView) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebEngineView::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebEngineView) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebEngineView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebEngineView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebEngineView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebEngineView_MetaObject
func callbackQWebEngineView_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebEngineView::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebEngineViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineView) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebEngineView::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebEngineView) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebEngineView::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebEngineView) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebEngineView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineView_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebEngineView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
