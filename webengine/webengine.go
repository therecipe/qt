// +build !minimal

package webengine

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "webengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtWebEngine_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWebEngine_PackedString) []byte {
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

type CertificateErrorController struct {
	ptr unsafe.Pointer
}

type CertificateErrorController_ITF interface {
	CertificateErrorController_PTR() *CertificateErrorController
}

func (ptr *CertificateErrorController) CertificateErrorController_PTR() *CertificateErrorController {
	return ptr
}

func (ptr *CertificateErrorController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *CertificateErrorController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromCertificateErrorController(ptr CertificateErrorController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CertificateErrorController_PTR().Pointer()
	}
	return nil
}

func NewCertificateErrorControllerFromPointer(ptr unsafe.Pointer) (n *CertificateErrorController) {
	n = new(CertificateErrorController)
	n.SetPointer(ptr)
	return
}

func (ptr *CertificateErrorController) DestroyCertificateErrorController() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type ClientCertSelectController struct {
	ptr unsafe.Pointer
}

type ClientCertSelectController_ITF interface {
	ClientCertSelectController_PTR() *ClientCertSelectController
}

func (ptr *ClientCertSelectController) ClientCertSelectController_PTR() *ClientCertSelectController {
	return ptr
}

func (ptr *ClientCertSelectController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *ClientCertSelectController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromClientCertSelectController(ptr ClientCertSelectController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ClientCertSelectController_PTR().Pointer()
	}
	return nil
}

func NewClientCertSelectControllerFromPointer(ptr unsafe.Pointer) (n *ClientCertSelectController) {
	n = new(ClientCertSelectController)
	n.SetPointer(ptr)
	return
}

func (ptr *ClientCertSelectController) DestroyClientCertSelectController() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type CommandLinePrefStoreQt struct {
	ptr unsafe.Pointer
}

type CommandLinePrefStoreQt_ITF interface {
	CommandLinePrefStoreQt_PTR() *CommandLinePrefStoreQt
}

func (ptr *CommandLinePrefStoreQt) CommandLinePrefStoreQt_PTR() *CommandLinePrefStoreQt {
	return ptr
}

func (ptr *CommandLinePrefStoreQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *CommandLinePrefStoreQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromCommandLinePrefStoreQt(ptr CommandLinePrefStoreQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CommandLinePrefStoreQt_PTR().Pointer()
	}
	return nil
}

func NewCommandLinePrefStoreQtFromPointer(ptr unsafe.Pointer) (n *CommandLinePrefStoreQt) {
	n = new(CommandLinePrefStoreQt)
	n.SetPointer(ptr)
	return
}

func (ptr *CommandLinePrefStoreQt) DestroyCommandLinePrefStoreQt() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type GLContextHelper struct {
	core.QObject
}

type GLContextHelper_ITF interface {
	core.QObject_ITF
	GLContextHelper_PTR() *GLContextHelper
}

func (ptr *GLContextHelper) GLContextHelper_PTR() *GLContextHelper {
	return ptr
}

func (ptr *GLContextHelper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *GLContextHelper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromGLContextHelper(ptr GLContextHelper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.GLContextHelper_PTR().Pointer()
	}
	return nil
}

func NewGLContextHelperFromPointer(ptr unsafe.Pointer) (n *GLContextHelper) {
	n = new(GLContextHelper)
	n.SetPointer(ptr)
	return
}

type ProxyConfigServiceQt struct {
	ptr unsafe.Pointer
}

type ProxyConfigServiceQt_ITF interface {
	ProxyConfigServiceQt_PTR() *ProxyConfigServiceQt
}

func (ptr *ProxyConfigServiceQt) ProxyConfigServiceQt_PTR() *ProxyConfigServiceQt {
	return ptr
}

func (ptr *ProxyConfigServiceQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *ProxyConfigServiceQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromProxyConfigServiceQt(ptr ProxyConfigServiceQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ProxyConfigServiceQt_PTR().Pointer()
	}
	return nil
}

func NewProxyConfigServiceQtFromPointer(ptr unsafe.Pointer) (n *ProxyConfigServiceQt) {
	n = new(ProxyConfigServiceQt)
	n.SetPointer(ptr)
	return
}

func (ptr *ProxyConfigServiceQt) DestroyProxyConfigServiceQt() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQQuickWebEngineProfileFromPointer(ptr unsafe.Pointer) (n *QQuickWebEngineProfile) {
	n = new(QQuickWebEngineProfile)
	n.SetPointer(ptr)
	return
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

func NewQQuickWebEngineProfile(parent core.QObject_ITF) *QQuickWebEngineProfile {
	tmpValue := NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_NewQQuickWebEngineProfile(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickWebEngineProfile) CachePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_CachePath(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineProfile_CachePathChanged
func callbackQQuickWebEngineProfile_CachePathChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cachePathChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectCachePathChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cachePathChanged") {
			C.QQuickWebEngineProfile_ConnectCachePathChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cachePathChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cachePathChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cachePathChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectCachePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectCachePathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cachePathChanged")
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

func (ptr *QQuickWebEngineProfile) ClientCertificateStore() *QWebEngineClientCertificateStore {
	if ptr.Pointer() != nil {
		return NewQWebEngineClientCertificateStoreFromPointer(C.QQuickWebEngineProfile_ClientCertificateStore(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) CookieStore() *QWebEngineCookieStore {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEngineCookieStoreFromPointer(C.QQuickWebEngineProfile_CookieStore(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QQuickWebEngineProfile_DefaultProfile() *QQuickWebEngineProfile {
	tmpValue := NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickWebEngineProfile) DefaultProfile() *QQuickWebEngineProfile {
	tmpValue := NewQQuickWebEngineProfileFromPointer(C.QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickWebEngineProfile) DownloadPath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_DownloadPath(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineProfile_DownloadPathChanged
func callbackQQuickWebEngineProfile_DownloadPathChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "downloadPathChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectDownloadPathChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "downloadPathChanged") {
			C.QQuickWebEngineProfile_ConnectDownloadPathChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "downloadPathChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "downloadPathChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "downloadPathChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectDownloadPathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectDownloadPathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "downloadPathChanged")
	}
}

func (ptr *QQuickWebEngineProfile) DownloadPathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DownloadPathChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_HttpAcceptLanguage(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged
func callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "httpAcceptLanguageChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpAcceptLanguageChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "httpAcceptLanguageChanged") {
			C.QQuickWebEngineProfile_ConnectHttpAcceptLanguageChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "httpAcceptLanguageChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "httpAcceptLanguageChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "httpAcceptLanguageChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpAcceptLanguageChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpAcceptLanguageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "httpAcceptLanguageChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguageChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpAcceptLanguageChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWebEngineProfile_HttpCacheMaximumSize(ptr.Pointer())))
	}
	return 0
}

//export callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged
func callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "httpCacheMaximumSizeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheMaximumSizeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "httpCacheMaximumSizeChanged") {
			C.QQuickWebEngineProfile_ConnectHttpCacheMaximumSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "httpCacheMaximumSizeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "httpCacheMaximumSizeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "httpCacheMaximumSizeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheMaximumSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpCacheMaximumSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "httpCacheMaximumSizeChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpCacheMaximumSizeChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheType() QQuickWebEngineProfile__HttpCacheType {
	if ptr.Pointer() != nil {
		return QQuickWebEngineProfile__HttpCacheType(C.QQuickWebEngineProfile_HttpCacheType(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWebEngineProfile_HttpCacheTypeChanged
func callbackQQuickWebEngineProfile_HttpCacheTypeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "httpCacheTypeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheTypeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "httpCacheTypeChanged") {
			C.QQuickWebEngineProfile_ConnectHttpCacheTypeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "httpCacheTypeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "httpCacheTypeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "httpCacheTypeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpCacheTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "httpCacheTypeChanged")
	}
}

func (ptr *QQuickWebEngineProfile) HttpCacheTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_HttpCacheTypeChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) HttpUserAgent() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_HttpUserAgent(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineProfile_HttpUserAgentChanged
func callbackQQuickWebEngineProfile_HttpUserAgentChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "httpUserAgentChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectHttpUserAgentChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "httpUserAgentChanged") {
			C.QQuickWebEngineProfile_ConnectHttpUserAgentChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "httpUserAgentChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "httpUserAgentChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "httpUserAgentChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpUserAgentChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectHttpUserAgentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "httpUserAgentChanged")
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

func (ptr *QQuickWebEngineProfile) IsOffTheRecord() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineProfile_IsOffTheRecord(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) IsSpellCheckEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineProfile_IsSpellCheckEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWebEngineProfile) IsUsedForGlobalCertificateVerification() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineProfile_IsUsedForGlobalCertificateVerification(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_OffTheRecordChanged
func callbackQQuickWebEngineProfile_OffTheRecordChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "offTheRecordChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectOffTheRecordChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "offTheRecordChanged") {
			C.QQuickWebEngineProfile_ConnectOffTheRecordChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "offTheRecordChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "offTheRecordChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "offTheRecordChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectOffTheRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectOffTheRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "offTheRecordChanged")
	}
}

func (ptr *QQuickWebEngineProfile) OffTheRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_OffTheRecordChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicy() QQuickWebEngineProfile__PersistentCookiesPolicy {
	if ptr.Pointer() != nil {
		return QQuickWebEngineProfile__PersistentCookiesPolicy(C.QQuickWebEngineProfile_PersistentCookiesPolicy(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged
func callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "persistentCookiesPolicyChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPersistentCookiesPolicyChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "persistentCookiesPolicyChanged") {
			C.QQuickWebEngineProfile_ConnectPersistentCookiesPolicyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "persistentCookiesPolicyChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "persistentCookiesPolicyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "persistentCookiesPolicyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentCookiesPolicyChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPersistentCookiesPolicyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "persistentCookiesPolicyChanged")
	}
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicyChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PersistentCookiesPolicyChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_PersistentStoragePath(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineProfile_PersistentStoragePathChanged
func callbackQQuickWebEngineProfile_PersistentStoragePathChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "persistentStoragePathChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPersistentStoragePathChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "persistentStoragePathChanged") {
			C.QQuickWebEngineProfile_ConnectPersistentStoragePathChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "persistentStoragePathChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "persistentStoragePathChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "persistentStoragePathChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentStoragePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPersistentStoragePathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "persistentStoragePathChanged")
	}
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePathChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PersistentStoragePathChanged(ptr.Pointer())
	}
}

//export callbackQQuickWebEngineProfile_PresentNotification
func callbackQQuickWebEngineProfile_PresentNotification(ptr unsafe.Pointer, notification unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "presentNotification"); signal != nil {
		(*(*func(*QWebEngineNotification))(signal))(NewQWebEngineNotificationFromPointer(notification))
	}

}

func (ptr *QQuickWebEngineProfile) ConnectPresentNotification(f func(notification *QWebEngineNotification)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "presentNotification") {
			C.QQuickWebEngineProfile_ConnectPresentNotification(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "presentNotification"); signal != nil {
			f := func(notification *QWebEngineNotification) {
				(*(*func(*QWebEngineNotification))(signal))(notification)
				f(notification)
			}
			qt.ConnectSignal(ptr.Pointer(), "presentNotification", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "presentNotification", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectPresentNotification() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectPresentNotification(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "presentNotification")
	}
}

func (ptr *QQuickWebEngineProfile) PresentNotification(notification QWebEngineNotification_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_PresentNotification(ptr.Pointer(), PointerFromQWebEngineNotification(notification))
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

func (ptr *QQuickWebEngineProfile) SetCachePath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QQuickWebEngineProfile_SetCachePath(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QQuickWebEngineProfile) SetDownloadPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QQuickWebEngineProfile_SetDownloadPath(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QQuickWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	if ptr.Pointer() != nil {
		var httpAcceptLanguageC *C.char
		if httpAcceptLanguage != "" {
			httpAcceptLanguageC = C.CString(httpAcceptLanguage)
			defer C.free(unsafe.Pointer(httpAcceptLanguageC))
		}
		C.QQuickWebEngineProfile_SetHttpAcceptLanguage(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: httpAcceptLanguageC, len: C.longlong(len(httpAcceptLanguage))})
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
		var userAgentC *C.char
		if userAgent != "" {
			userAgentC = C.CString(userAgent)
			defer C.free(unsafe.Pointer(userAgentC))
		}
		C.QQuickWebEngineProfile_SetHttpUserAgent(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: userAgentC, len: C.longlong(len(userAgent))})
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
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QQuickWebEngineProfile_SetPersistentStoragePath(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QQuickWebEngineProfile) SetSpellCheckEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetSpellCheckEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickWebEngineProfile) SetSpellCheckLanguages(languages []string) {
	if ptr.Pointer() != nil {
		languagesC := C.CString(strings.Join(languages, "¡¦!"))
		defer C.free(unsafe.Pointer(languagesC))
		C.QQuickWebEngineProfile_SetSpellCheckLanguages(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: languagesC, len: C.longlong(len(strings.Join(languages, "¡¦!")))})
	}
}

func (ptr *QQuickWebEngineProfile) SetStorageName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QQuickWebEngineProfile_SetStorageName(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QQuickWebEngineProfile) SetUrlRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetUrlRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

func (ptr *QQuickWebEngineProfile) SetUseForGlobalCertificateVerification(b bool) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SetUseForGlobalCertificateVerification(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(b))))
	}
}

//export callbackQQuickWebEngineProfile_SpellCheckEnabledChanged
func callbackQQuickWebEngineProfile_SpellCheckEnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "spellCheckEnabledChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectSpellCheckEnabledChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "spellCheckEnabledChanged") {
			C.QQuickWebEngineProfile_ConnectSpellCheckEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "spellCheckEnabledChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "spellCheckEnabledChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "spellCheckEnabledChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectSpellCheckEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectSpellCheckEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "spellCheckEnabledChanged")
	}
}

func (ptr *QQuickWebEngineProfile) SpellCheckEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SpellCheckEnabledChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) SpellCheckLanguages() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QQuickWebEngineProfile_SpellCheckLanguages(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQQuickWebEngineProfile_SpellCheckLanguagesChanged
func callbackQQuickWebEngineProfile_SpellCheckLanguagesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "spellCheckLanguagesChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectSpellCheckLanguagesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "spellCheckLanguagesChanged") {
			C.QQuickWebEngineProfile_ConnectSpellCheckLanguagesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "spellCheckLanguagesChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "spellCheckLanguagesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "spellCheckLanguagesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectSpellCheckLanguagesChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectSpellCheckLanguagesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "spellCheckLanguagesChanged")
	}
}

func (ptr *QQuickWebEngineProfile) SpellCheckLanguagesChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_SpellCheckLanguagesChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) StorageName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineProfile_StorageName(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineProfile_StorageNameChanged
func callbackQQuickWebEngineProfile_StorageNameChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "storageNameChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectStorageNameChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "storageNameChanged") {
			C.QQuickWebEngineProfile_ConnectStorageNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "storageNameChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "storageNameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "storageNameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectStorageNameChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectStorageNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "storageNameChanged")
	}
}

func (ptr *QQuickWebEngineProfile) StorageNameChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_StorageNameChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) UrlSchemeHandler(scheme core.QByteArray_ITF) *QWebEngineUrlSchemeHandler {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEngineUrlSchemeHandlerFromPointer(C.QQuickWebEngineProfile_UrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(scheme)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWebEngineProfile_UseForGlobalCertificateVerificationChanged
func callbackQQuickWebEngineProfile_UseForGlobalCertificateVerificationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "useForGlobalCertificateVerificationChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQuickWebEngineProfile) ConnectUseForGlobalCertificateVerificationChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "useForGlobalCertificateVerificationChanged") {
			C.QQuickWebEngineProfile_ConnectUseForGlobalCertificateVerificationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "useForGlobalCertificateVerificationChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "useForGlobalCertificateVerificationChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "useForGlobalCertificateVerificationChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectUseForGlobalCertificateVerificationChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectUseForGlobalCertificateVerificationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "useForGlobalCertificateVerificationChanged")
	}
}

func (ptr *QQuickWebEngineProfile) UseForGlobalCertificateVerificationChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_UseForGlobalCertificateVerificationChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWebEngineProfile) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineProfile___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineProfile) __children_newList() unsafe.Pointer {
	return C.QQuickWebEngineProfile___children_newList(ptr.Pointer())
}

func (ptr *QQuickWebEngineProfile) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickWebEngineProfile___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickWebEngineProfile) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickWebEngineProfile___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickWebEngineProfile) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineProfile___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineProfile) __findChildren_newList() unsafe.Pointer {
	return C.QQuickWebEngineProfile___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickWebEngineProfile) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineProfile___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineProfile) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickWebEngineProfile___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickWebEngineProfile) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineProfile___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineProfile) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineProfile) __qFindChildren_newList2() unsafe.Pointer {
	return C.QQuickWebEngineProfile___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQQuickWebEngineProfile_ChildEvent
func callbackQQuickWebEngineProfile_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_ConnectNotify
func callbackQQuickWebEngineProfile_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineProfile_CustomEvent
func callbackQQuickWebEngineProfile_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWebEngineProfile_DeleteLater
func callbackQQuickWebEngineProfile_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWebEngineProfile) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickWebEngineProfile_Destroyed
func callbackQQuickWebEngineProfile_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickWebEngineProfile_DisconnectNotify
func callbackQQuickWebEngineProfile_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineProfile_Event
func callbackQQuickWebEngineProfile_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWebEngineProfileFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineProfile_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_EventFilter
func callbackQQuickWebEngineProfile_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWebEngineProfileFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineProfile_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickWebEngineProfile_ObjectNameChanged
func callbackQQuickWebEngineProfile_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickWebEngineProfile_TimerEvent
func callbackQQuickWebEngineProfile_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWebEngineProfileFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineProfile_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickWebEngineScript struct {
	core.QObject
}

type QQuickWebEngineScript_ITF interface {
	core.QObject_ITF
	QQuickWebEngineScript_PTR() *QQuickWebEngineScript
}

func (ptr *QQuickWebEngineScript) QQuickWebEngineScript_PTR() *QQuickWebEngineScript {
	return ptr
}

func (ptr *QQuickWebEngineScript) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickWebEngineScript) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickWebEngineScript(ptr QQuickWebEngineScript_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWebEngineScript_PTR().Pointer()
	}
	return nil
}

func NewQQuickWebEngineScriptFromPointer(ptr unsafe.Pointer) (n *QQuickWebEngineScript) {
	n = new(QQuickWebEngineScript)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQuickWebEngineScript__InjectionPoint
//QQuickWebEngineScript::InjectionPoint
type QQuickWebEngineScript__InjectionPoint int64

const (
	QQuickWebEngineScript__Deferred         QQuickWebEngineScript__InjectionPoint = QQuickWebEngineScript__InjectionPoint(0)
	QQuickWebEngineScript__DocumentReady    QQuickWebEngineScript__InjectionPoint = QQuickWebEngineScript__InjectionPoint(1)
	QQuickWebEngineScript__DocumentCreation QQuickWebEngineScript__InjectionPoint = QQuickWebEngineScript__InjectionPoint(2)
)

//go:generate stringer -type=QQuickWebEngineScript__ScriptWorldId
//QQuickWebEngineScript::ScriptWorldId
type QQuickWebEngineScript__ScriptWorldId int64

const (
	QQuickWebEngineScript__MainWorld        QQuickWebEngineScript__ScriptWorldId = QQuickWebEngineScript__ScriptWorldId(0)
	QQuickWebEngineScript__ApplicationWorld QQuickWebEngineScript__ScriptWorldId = QQuickWebEngineScript__ScriptWorldId(1)
	QQuickWebEngineScript__UserWorld        QQuickWebEngineScript__ScriptWorldId = QQuickWebEngineScript__ScriptWorldId(2)
)

func NewQQuickWebEngineScript(parent core.QObject_ITF) *QQuickWebEngineScript {
	tmpValue := NewQQuickWebEngineScriptFromPointer(C.QQuickWebEngineScript_NewQQuickWebEngineScript(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickWebEngineScript) InjectionPoint() QQuickWebEngineScript__InjectionPoint {
	if ptr.Pointer() != nil {
		return QQuickWebEngineScript__InjectionPoint(C.QQuickWebEngineScript_InjectionPoint(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWebEngineScript_InjectionPointChanged
func callbackQQuickWebEngineScript_InjectionPointChanged(ptr unsafe.Pointer, injectionPoint C.longlong) {
	if signal := qt.GetSignal(ptr, "injectionPointChanged"); signal != nil {
		(*(*func(QQuickWebEngineScript__InjectionPoint))(signal))(QQuickWebEngineScript__InjectionPoint(injectionPoint))
	}

}

func (ptr *QQuickWebEngineScript) ConnectInjectionPointChanged(f func(injectionPoint QQuickWebEngineScript__InjectionPoint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "injectionPointChanged") {
			C.QQuickWebEngineScript_ConnectInjectionPointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "injectionPointChanged"); signal != nil {
			f := func(injectionPoint QQuickWebEngineScript__InjectionPoint) {
				(*(*func(QQuickWebEngineScript__InjectionPoint))(signal))(injectionPoint)
				f(injectionPoint)
			}
			qt.ConnectSignal(ptr.Pointer(), "injectionPointChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "injectionPointChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineScript) DisconnectInjectionPointChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectInjectionPointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "injectionPointChanged")
	}
}

func (ptr *QQuickWebEngineScript) InjectionPointChanged(injectionPoint QQuickWebEngineScript__InjectionPoint) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_InjectionPointChanged(ptr.Pointer(), C.longlong(injectionPoint))
	}
}

func (ptr *QQuickWebEngineScript) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineScript_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineScript_NameChanged
func callbackQQuickWebEngineScript_NameChanged(ptr unsafe.Pointer, name C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "nameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(name))
	}

}

func (ptr *QQuickWebEngineScript) ConnectNameChanged(f func(name string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameChanged") {
			C.QQuickWebEngineScript_ConnectNameChanged(ptr.Pointer())
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

func (ptr *QQuickWebEngineScript) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nameChanged")
	}
}

func (ptr *QQuickWebEngineScript) NameChanged(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QQuickWebEngineScript_NameChanged(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QQuickWebEngineScript) RunOnSubframes() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineScript_RunOnSubframes(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickWebEngineScript_RunOnSubframesChanged
func callbackQQuickWebEngineScript_RunOnSubframesChanged(ptr unsafe.Pointer, on C.char) {
	if signal := qt.GetSignal(ptr, "runOnSubframesChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(on) != 0)
	}

}

func (ptr *QQuickWebEngineScript) ConnectRunOnSubframesChanged(f func(on bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "runOnSubframesChanged") {
			C.QQuickWebEngineScript_ConnectRunOnSubframesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "runOnSubframesChanged"); signal != nil {
			f := func(on bool) {
				(*(*func(bool))(signal))(on)
				f(on)
			}
			qt.ConnectSignal(ptr.Pointer(), "runOnSubframesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "runOnSubframesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineScript) DisconnectRunOnSubframesChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectRunOnSubframesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "runOnSubframesChanged")
	}
}

func (ptr *QQuickWebEngineScript) RunOnSubframesChanged(on bool) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_RunOnSubframesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QQuickWebEngineScript) SetInjectionPoint(injectionPoint QQuickWebEngineScript__InjectionPoint) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_SetInjectionPoint(ptr.Pointer(), C.longlong(injectionPoint))
	}
}

func (ptr *QQuickWebEngineScript) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QQuickWebEngineScript_SetName(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QQuickWebEngineScript) SetRunOnSubframes(on bool) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_SetRunOnSubframes(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QQuickWebEngineScript) SetSourceCode(code string) {
	if ptr.Pointer() != nil {
		var codeC *C.char
		if code != "" {
			codeC = C.CString(code)
			defer C.free(unsafe.Pointer(codeC))
		}
		C.QQuickWebEngineScript_SetSourceCode(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: codeC, len: C.longlong(len(code))})
	}
}

func (ptr *QQuickWebEngineScript) SetSourceUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_SetSourceUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickWebEngineScript) SetWorldId(scriptWorldId QQuickWebEngineScript__ScriptWorldId) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_SetWorldId(ptr.Pointer(), C.longlong(scriptWorldId))
	}
}

func (ptr *QQuickWebEngineScript) SourceCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineScript_SourceCode(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickWebEngineScript_SourceCodeChanged
func callbackQQuickWebEngineScript_SourceCodeChanged(ptr unsafe.Pointer, code C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "sourceCodeChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(code))
	}

}

func (ptr *QQuickWebEngineScript) ConnectSourceCodeChanged(f func(code string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sourceCodeChanged") {
			C.QQuickWebEngineScript_ConnectSourceCodeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sourceCodeChanged"); signal != nil {
			f := func(code string) {
				(*(*func(string))(signal))(code)
				f(code)
			}
			qt.ConnectSignal(ptr.Pointer(), "sourceCodeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sourceCodeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineScript) DisconnectSourceCodeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectSourceCodeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sourceCodeChanged")
	}
}

func (ptr *QQuickWebEngineScript) SourceCodeChanged(code string) {
	if ptr.Pointer() != nil {
		var codeC *C.char
		if code != "" {
			codeC = C.CString(code)
			defer C.free(unsafe.Pointer(codeC))
		}
		C.QQuickWebEngineScript_SourceCodeChanged(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: codeC, len: C.longlong(len(code))})
	}
}

func (ptr *QQuickWebEngineScript) SourceUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQuickWebEngineScript_SourceUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWebEngineScript_SourceUrlChanged
func callbackQQuickWebEngineScript_SourceUrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sourceUrlChanged"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQuickWebEngineScript) ConnectSourceUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sourceUrlChanged") {
			C.QQuickWebEngineScript_ConnectSourceUrlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sourceUrlChanged"); signal != nil {
			f := func(url *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "sourceUrlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sourceUrlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineScript) DisconnectSourceUrlChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectSourceUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sourceUrlChanged")
	}
}

func (ptr *QQuickWebEngineScript) SourceUrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_SourceUrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickWebEngineScript) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickWebEngineScript_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickWebEngineScript) WorldId() QQuickWebEngineScript__ScriptWorldId {
	if ptr.Pointer() != nil {
		return QQuickWebEngineScript__ScriptWorldId(C.QQuickWebEngineScript_WorldId(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWebEngineScript_WorldIdChanged
func callbackQQuickWebEngineScript_WorldIdChanged(ptr unsafe.Pointer, scriptWorldId C.longlong) {
	if signal := qt.GetSignal(ptr, "worldIdChanged"); signal != nil {
		(*(*func(QQuickWebEngineScript__ScriptWorldId))(signal))(QQuickWebEngineScript__ScriptWorldId(scriptWorldId))
	}

}

func (ptr *QQuickWebEngineScript) ConnectWorldIdChanged(f func(scriptWorldId QQuickWebEngineScript__ScriptWorldId)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "worldIdChanged") {
			C.QQuickWebEngineScript_ConnectWorldIdChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "worldIdChanged"); signal != nil {
			f := func(scriptWorldId QQuickWebEngineScript__ScriptWorldId) {
				(*(*func(QQuickWebEngineScript__ScriptWorldId))(signal))(scriptWorldId)
				f(scriptWorldId)
			}
			qt.ConnectSignal(ptr.Pointer(), "worldIdChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "worldIdChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQuickWebEngineScript) DisconnectWorldIdChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectWorldIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "worldIdChanged")
	}
}

func (ptr *QQuickWebEngineScript) WorldIdChanged(scriptWorldId QQuickWebEngineScript__ScriptWorldId) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_WorldIdChanged(ptr.Pointer(), C.longlong(scriptWorldId))
	}
}

func (ptr *QQuickWebEngineScript) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineScript___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineScript) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineScript) __children_newList() unsafe.Pointer {
	return C.QQuickWebEngineScript___children_newList(ptr.Pointer())
}

func (ptr *QQuickWebEngineScript) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickWebEngineScript___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineScript) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickWebEngineScript) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickWebEngineScript___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickWebEngineScript) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineScript___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineScript) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineScript) __findChildren_newList() unsafe.Pointer {
	return C.QQuickWebEngineScript___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickWebEngineScript) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineScript___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineScript) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineScript) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickWebEngineScript___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickWebEngineScript) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWebEngineScript___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWebEngineScript) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWebEngineScript) __qFindChildren_newList2() unsafe.Pointer {
	return C.QQuickWebEngineScript___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQQuickWebEngineScript_ChildEvent
func callbackQQuickWebEngineScript_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWebEngineScriptFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineScript) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWebEngineScript_ConnectNotify
func callbackQQuickWebEngineScript_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineScriptFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineScript) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineScript_CustomEvent
func callbackQQuickWebEngineScript_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWebEngineScriptFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineScript) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWebEngineScript_DeleteLater
func callbackQQuickWebEngineScript_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQQuickWebEngineScriptFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWebEngineScript) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickWebEngineScript_Destroyed
func callbackQQuickWebEngineScript_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickWebEngineScript_DisconnectNotify
func callbackQQuickWebEngineScript_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWebEngineScriptFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWebEngineScript) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWebEngineScript_Event
func callbackQQuickWebEngineScript_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWebEngineScriptFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWebEngineScript) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineScript_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickWebEngineScript_EventFilter
func callbackQQuickWebEngineScript_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWebEngineScriptFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWebEngineScript) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWebEngineScript_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickWebEngineScript_ObjectNameChanged
func callbackQQuickWebEngineScript_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickWebEngineScript_TimerEvent
func callbackQQuickWebEngineScript_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWebEngineScriptFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWebEngineScript) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWebEngineScript_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWebEngineCallback struct {
	ptr unsafe.Pointer
}

type QWebEngineCallback_ITF interface {
	QWebEngineCallback_PTR() *QWebEngineCallback
}

func (ptr *QWebEngineCallback) QWebEngineCallback_PTR() *QWebEngineCallback {
	return ptr
}

func (ptr *QWebEngineCallback) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineCallback) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineCallback(ptr QWebEngineCallback_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineCallback_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineCallbackFromPointer(ptr unsafe.Pointer) (n *QWebEngineCallback) {
	n = new(QWebEngineCallback)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineCallback) DestroyQWebEngineCallback() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQWebEngineCertificateErrorFromPointer(ptr unsafe.Pointer) (n *QWebEngineCertificateError) {
	n = new(QWebEngineCertificateError)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineCertificateError) DestroyQWebEngineCertificateError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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
	QWebEngineCertificateError__CertificateValidityTooLong         QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-213)
	QWebEngineCertificateError__CertificateTransparencyRequired    QWebEngineCertificateError__Error = QWebEngineCertificateError__Error(-214)
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
		return int8(C.QWebEngineCertificateError_IsOverridable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebEngineCertificateError) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineCertificateError_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

type QWebEngineClientCertificateStore struct {
	ptr unsafe.Pointer
}

type QWebEngineClientCertificateStore_ITF interface {
	QWebEngineClientCertificateStore_PTR() *QWebEngineClientCertificateStore
}

func (ptr *QWebEngineClientCertificateStore) QWebEngineClientCertificateStore_PTR() *QWebEngineClientCertificateStore {
	return ptr
}

func (ptr *QWebEngineClientCertificateStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineClientCertificateStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineClientCertificateStore(ptr QWebEngineClientCertificateStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineClientCertificateStore_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineClientCertificateStoreFromPointer(ptr unsafe.Pointer) (n *QWebEngineClientCertificateStore) {
	n = new(QWebEngineClientCertificateStore)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineClientCertificateStore) DestroyQWebEngineClientCertificateStore() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineClientCertificateStore) Add(certificate network.QSslCertificate_ITF, privateKey network.QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateStore_Add(ptr.Pointer(), network.PointerFromQSslCertificate(certificate), network.PointerFromQSslKey(privateKey))
	}
}

func (ptr *QWebEngineClientCertificateStore) Certificates() []*network.QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*network.QSslCertificate {
			out := make([]*network.QSslCertificate, int(l.len))
			tmpList := NewQWebEngineClientCertificateStoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__certificates_atList(i)
			}
			return out
		}(C.QWebEngineClientCertificateStore_Certificates(ptr.Pointer()))
	}
	return make([]*network.QSslCertificate, 0)
}

func (ptr *QWebEngineClientCertificateStore) Clear() {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateStore_Clear(ptr.Pointer())
	}
}

func (ptr *QWebEngineClientCertificateStore) Remove(certificate network.QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateStore_Remove(ptr.Pointer(), network.PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QWebEngineClientCertificateStore) __certificates_atList(i int) *network.QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslCertificateFromPointer(C.QWebEngineClientCertificateStore___certificates_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*network.QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineClientCertificateStore) __certificates_setList(i network.QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateStore___certificates_setList(ptr.Pointer(), network.PointerFromQSslCertificate(i))
	}
}

func (ptr *QWebEngineClientCertificateStore) __certificates_newList() unsafe.Pointer {
	return C.QWebEngineClientCertificateStore___certificates_newList(ptr.Pointer())
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

func NewQWebEngineContextMenuDataFromPointer(ptr unsafe.Pointer) (n *QWebEngineContextMenuData) {
	n = new(QWebEngineContextMenuData)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineContextMenuData) DestroyQWebEngineContextMenuData() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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

func NewQWebEngineCookieStoreFromPointer(ptr unsafe.Pointer) (n *QWebEngineCookieStore) {
	n = new(QWebEngineCookieStore)
	n.SetPointer(ptr)
	return
}

//export callbackQWebEngineCookieStore_CookieAdded
func callbackQWebEngineCookieStore_CookieAdded(ptr unsafe.Pointer, cookie unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cookieAdded"); signal != nil {
		(*(*func(*network.QNetworkCookie))(signal))(network.NewQNetworkCookieFromPointer(cookie))
	}

}

func (ptr *QWebEngineCookieStore) ConnectCookieAdded(f func(cookie *network.QNetworkCookie)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cookieAdded") {
			C.QWebEngineCookieStore_ConnectCookieAdded(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cookieAdded"); signal != nil {
			f := func(cookie *network.QNetworkCookie) {
				(*(*func(*network.QNetworkCookie))(signal))(cookie)
				f(cookie)
			}
			qt.ConnectSignal(ptr.Pointer(), "cookieAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cookieAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCookieAdded() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectCookieAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cookieAdded")
	}
}

func (ptr *QWebEngineCookieStore) CookieAdded(cookie network.QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CookieAdded(ptr.Pointer(), network.PointerFromQNetworkCookie(cookie))
	}
}

//export callbackQWebEngineCookieStore_CookieRemoved
func callbackQWebEngineCookieStore_CookieRemoved(ptr unsafe.Pointer, cookie unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cookieRemoved"); signal != nil {
		(*(*func(*network.QNetworkCookie))(signal))(network.NewQNetworkCookieFromPointer(cookie))
	}

}

func (ptr *QWebEngineCookieStore) ConnectCookieRemoved(f func(cookie *network.QNetworkCookie)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cookieRemoved") {
			C.QWebEngineCookieStore_ConnectCookieRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cookieRemoved"); signal != nil {
			f := func(cookie *network.QNetworkCookie) {
				(*(*func(*network.QNetworkCookie))(signal))(cookie)
				f(cookie)
			}
			qt.ConnectSignal(ptr.Pointer(), "cookieRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cookieRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineCookieStore) DisconnectCookieRemoved() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectCookieRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cookieRemoved")
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
	if signal := qt.GetSignal(ptr, "~QWebEngineCookieStore"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DestroyQWebEngineCookieStoreDefault()
	}
}

func (ptr *QWebEngineCookieStore) ConnectDestroyQWebEngineCookieStore(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebEngineCookieStore"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineCookieStore", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineCookieStore", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineCookieStore) DisconnectDestroyQWebEngineCookieStore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebEngineCookieStore")
	}
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStore() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStore(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStoreDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStoreDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineCookieStore) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineCookieStore___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineCookieStore) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineCookieStore) __children_newList() unsafe.Pointer {
	return C.QWebEngineCookieStore___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineCookieStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineCookieStore___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineCookieStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineCookieStore) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineCookieStore___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineCookieStore) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineCookieStore___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineCookieStore) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineCookieStore) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineCookieStore___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineCookieStore) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineCookieStore___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineCookieStore) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineCookieStore) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineCookieStore___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineCookieStore) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineCookieStore___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineCookieStore) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineCookieStore) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineCookieStore___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineCookieStore_ChildEvent
func callbackQWebEngineCookieStore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineCookieStore_ConnectNotify
func callbackQWebEngineCookieStore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineCookieStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineCookieStore_CustomEvent
func callbackQWebEngineCookieStore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineCookieStore_DeleteLater
func callbackQWebEngineCookieStore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineCookieStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineCookieStore_Destroyed
func callbackQWebEngineCookieStore_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineCookieStore_DisconnectNotify
func callbackQWebEngineCookieStore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineCookieStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineCookieStore_Event
func callbackQWebEngineCookieStore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineCookieStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineCookieStore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineCookieStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEngineCookieStore_EventFilter
func callbackQWebEngineCookieStore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineCookieStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineCookieStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineCookieStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineCookieStore_ObjectNameChanged
func callbackQWebEngineCookieStore_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineCookieStore_TimerEvent
func callbackQWebEngineCookieStore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineCookieStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineCookieStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineCookieStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWebEngineHttpRequest struct {
	ptr unsafe.Pointer
}

type QWebEngineHttpRequest_ITF interface {
	QWebEngineHttpRequest_PTR() *QWebEngineHttpRequest
}

func (ptr *QWebEngineHttpRequest) QWebEngineHttpRequest_PTR() *QWebEngineHttpRequest {
	return ptr
}

func (ptr *QWebEngineHttpRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineHttpRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineHttpRequest(ptr QWebEngineHttpRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineHttpRequest_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineHttpRequestFromPointer(ptr unsafe.Pointer) (n *QWebEngineHttpRequest) {
	n = new(QWebEngineHttpRequest)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QWebEngineHttpRequest__Method
//QWebEngineHttpRequest::Method
type QWebEngineHttpRequest__Method int64

const (
	QWebEngineHttpRequest__Get  QWebEngineHttpRequest__Method = QWebEngineHttpRequest__Method(0)
	QWebEngineHttpRequest__Post QWebEngineHttpRequest__Method = QWebEngineHttpRequest__Method(1)
)

func NewQWebEngineHttpRequest(url core.QUrl_ITF, method QWebEngineHttpRequest__Method) *QWebEngineHttpRequest {
	tmpValue := NewQWebEngineHttpRequestFromPointer(C.QWebEngineHttpRequest_NewQWebEngineHttpRequest(core.PointerFromQUrl(url), C.longlong(method)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
	return tmpValue
}

func NewQWebEngineHttpRequest2(other QWebEngineHttpRequest_ITF) *QWebEngineHttpRequest {
	tmpValue := NewQWebEngineHttpRequestFromPointer(C.QWebEngineHttpRequest_NewQWebEngineHttpRequest2(PointerFromQWebEngineHttpRequest(other)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
	return tmpValue
}

func (ptr *QWebEngineHttpRequest) HasHeader(headerName core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineHttpRequest_HasHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QWebEngineHttpRequest) Header(headerName core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineHttpRequest_Header(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHttpRequest) Headers() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQWebEngineHttpRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__headers_atList(i)
			}
			return out
		}(C.QWebEngineHttpRequest_Headers(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QWebEngineHttpRequest) Method() QWebEngineHttpRequest__Method {
	if ptr.Pointer() != nil {
		return QWebEngineHttpRequest__Method(C.QWebEngineHttpRequest_Method(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineHttpRequest) PostData() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineHttpRequest_PostData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func QWebEngineHttpRequest_PostRequest(url core.QUrl_ITF, postData map[string]string) *QWebEngineHttpRequest {
	tmpValue := NewQWebEngineHttpRequestFromPointer(C.QWebEngineHttpRequest_QWebEngineHttpRequest_PostRequest(core.PointerFromQUrl(url), func() unsafe.Pointer {
		tmpList := NewQWebEngineHttpRequestFromPointer(NewQWebEngineHttpRequestFromPointer(nil).__postRequest_postData_newList())
		for k, v := range postData {
			tmpList.__postRequest_postData_setList(k, v)
		}
		return tmpList.Pointer()
	}()))
	runtime.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
	return tmpValue
}

func (ptr *QWebEngineHttpRequest) PostRequest(url core.QUrl_ITF, postData map[string]string) *QWebEngineHttpRequest {
	tmpValue := NewQWebEngineHttpRequestFromPointer(C.QWebEngineHttpRequest_QWebEngineHttpRequest_PostRequest(core.PointerFromQUrl(url), func() unsafe.Pointer {
		tmpList := NewQWebEngineHttpRequestFromPointer(NewQWebEngineHttpRequestFromPointer(nil).__postRequest_postData_newList())
		for k, v := range postData {
			tmpList.__postRequest_postData_setList(k, v)
		}
		return tmpList.Pointer()
	}()))
	runtime.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
	return tmpValue
}

func (ptr *QWebEngineHttpRequest) SetHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_SetHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QWebEngineHttpRequest) SetMethod(method QWebEngineHttpRequest__Method) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_SetMethod(ptr.Pointer(), C.longlong(method))
	}
}

func (ptr *QWebEngineHttpRequest) SetPostData(postData core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_SetPostData(ptr.Pointer(), core.PointerFromQByteArray(postData))
	}
}

func (ptr *QWebEngineHttpRequest) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineHttpRequest) Swap(other QWebEngineHttpRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_Swap(ptr.Pointer(), PointerFromQWebEngineHttpRequest(other))
	}
}

func (ptr *QWebEngineHttpRequest) UnsetHeader(key core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_UnsetHeader(ptr.Pointer(), core.PointerFromQByteArray(key))
	}
}

func (ptr *QWebEngineHttpRequest) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineHttpRequest_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHttpRequest) DestroyQWebEngineHttpRequest() {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest_DestroyQWebEngineHttpRequest(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineHttpRequest) __headers_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineHttpRequest___headers_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHttpRequest) __headers_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineHttpRequest___headers_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineHttpRequest) __headers_newList() unsafe.Pointer {
	return C.QWebEngineHttpRequest___headers_newList(ptr.Pointer())
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_atList(v string, i int) string {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		return cGoUnpackString(C.QWebEngineHttpRequest___postRequest_postData_atList(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_setList(key string, i string) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebEngineHttpRequest___postRequest_postData_setList(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: keyC, len: C.longlong(len(key))}, C.struct_QtWebEngine_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_newList() unsafe.Pointer {
	return C.QWebEngineHttpRequest___postRequest_postData_newList(ptr.Pointer())
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQWebEngineHttpRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____postRequest_postData_keyList_atList(i)
			}
			return out
		}(C.QWebEngineHttpRequest___postRequest_postData_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebEngineHttpRequest) ____postRequest_postData_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineHttpRequest_____postRequest_postData_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebEngineHttpRequest) ____postRequest_postData_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebEngineHttpRequest_____postRequest_postData_keyList_setList(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebEngineHttpRequest) ____postRequest_postData_keyList_newList() unsafe.Pointer {
	return C.QWebEngineHttpRequest_____postRequest_postData_keyList_newList(ptr.Pointer())
}

type QWebEngineNotification struct {
	core.QObject
}

type QWebEngineNotification_ITF interface {
	core.QObject_ITF
	QWebEngineNotification_PTR() *QWebEngineNotification
}

func (ptr *QWebEngineNotification) QWebEngineNotification_PTR() *QWebEngineNotification {
	return ptr
}

func (ptr *QWebEngineNotification) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebEngineNotification) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebEngineNotification(ptr QWebEngineNotification_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineNotification_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineNotificationFromPointer(ptr unsafe.Pointer) (n *QWebEngineNotification) {
	n = new(QWebEngineNotification)
	n.SetPointer(ptr)
	return
}

//export callbackQWebEngineNotification_Click
func callbackQWebEngineNotification_Click(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "click"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineNotificationFromPointer(ptr).ClickDefault()
	}
}

func (ptr *QWebEngineNotification) ConnectClick(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "click"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineNotification) DisconnectClick() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "click")
	}
}

func (ptr *QWebEngineNotification) Click() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_Click(ptr.Pointer())
	}
}

func (ptr *QWebEngineNotification) ClickDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_ClickDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineNotification_Close
func callbackQWebEngineNotification_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineNotificationFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QWebEngineNotification) ConnectClose(f func()) {
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

func (ptr *QWebEngineNotification) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QWebEngineNotification) Close() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_Close(ptr.Pointer())
	}
}

func (ptr *QWebEngineNotification) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_CloseDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineNotification_Closed
func callbackQWebEngineNotification_Closed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEngineNotification) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closed") {
			C.QWebEngineNotification_ConnectClosed(ptr.Pointer())
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

func (ptr *QWebEngineNotification) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "closed")
	}
}

func (ptr *QWebEngineNotification) Closed() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_Closed(ptr.Pointer())
	}
}

func (ptr *QWebEngineNotification) Direction() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QWebEngineNotification_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineNotification) Icon() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QWebEngineNotification_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineNotification) Language() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineNotification_Language(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineNotification) Matches(other QWebEngineNotification_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineNotification_Matches(ptr.Pointer(), PointerFromQWebEngineNotification(other))) != 0
	}
	return false
}

func (ptr *QWebEngineNotification) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineNotification_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineNotification) Origin() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineNotification_Origin(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineNotification_Show
func callbackQWebEngineNotification_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineNotificationFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebEngineNotification) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "show"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineNotification) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "show")
	}
}

func (ptr *QWebEngineNotification) Show() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_Show(ptr.Pointer())
	}
}

func (ptr *QWebEngineNotification) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_ShowDefault(ptr.Pointer())
	}
}

func (ptr *QWebEngineNotification) Tag() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineNotification_Tag(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineNotification) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineNotification_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineNotification) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineNotification___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineNotification) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineNotification) __children_newList() unsafe.Pointer {
	return C.QWebEngineNotification___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineNotification) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineNotification___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineNotification) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineNotification) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineNotification___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineNotification) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineNotification___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineNotification) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineNotification) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineNotification___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineNotification) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineNotification___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineNotification) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineNotification) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineNotification___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineNotification) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineNotification___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineNotification) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineNotification) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineNotification___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineNotification_ChildEvent
func callbackQWebEngineNotification_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineNotificationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineNotification) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineNotification_ConnectNotify
func callbackQWebEngineNotification_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineNotificationFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineNotification) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineNotification_CustomEvent
func callbackQWebEngineNotification_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineNotificationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineNotification) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineNotification_DeleteLater
func callbackQWebEngineNotification_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineNotificationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineNotification) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineNotification_Destroyed
func callbackQWebEngineNotification_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineNotification_DisconnectNotify
func callbackQWebEngineNotification_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineNotificationFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineNotification) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineNotification_Event
func callbackQWebEngineNotification_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineNotificationFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineNotification) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineNotification_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEngineNotification_EventFilter
func callbackQWebEngineNotification_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineNotificationFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineNotification) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineNotification_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineNotification_ObjectNameChanged
func callbackQWebEngineNotification_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineNotification_TimerEvent
func callbackQWebEngineNotification_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineNotificationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineNotification) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineNotification_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
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

func NewQWebEnginePageFromPointer(ptr unsafe.Pointer) (n *QWebEnginePage) {
	n = new(QWebEnginePage)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QWebEnginePage__FileSelectionMode
//QWebEnginePage::FileSelectionMode
type QWebEnginePage__FileSelectionMode int64

const (
	QWebEnginePage__FileSelectOpen         QWebEnginePage__FileSelectionMode = QWebEnginePage__FileSelectionMode(0)
	QWebEnginePage__FileSelectOpenMultiple QWebEnginePage__FileSelectionMode = QWebEnginePage__FileSelectionMode(1)
)

//go:generate stringer -type=QWebEnginePage__JavaScriptConsoleMessageLevel
//QWebEnginePage::JavaScriptConsoleMessageLevel
type QWebEnginePage__JavaScriptConsoleMessageLevel int64

const (
	QWebEnginePage__InfoMessageLevel    QWebEnginePage__JavaScriptConsoleMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(0)
	QWebEnginePage__WarningMessageLevel QWebEnginePage__JavaScriptConsoleMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(1)
	QWebEnginePage__ErrorMessageLevel   QWebEnginePage__JavaScriptConsoleMessageLevel = QWebEnginePage__JavaScriptConsoleMessageLevel(2)
)

//go:generate stringer -type=QWebEnginePage__WebAction
//QWebEnginePage::WebAction
type QWebEnginePage__WebAction int64

const (
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
	QWebEnginePage__ViewSource                 QWebEnginePage__WebAction = QWebEnginePage__WebAction(32)
	QWebEnginePage__ToggleBold                 QWebEnginePage__WebAction = QWebEnginePage__WebAction(33)
	QWebEnginePage__ToggleItalic               QWebEnginePage__WebAction = QWebEnginePage__WebAction(34)
	QWebEnginePage__ToggleUnderline            QWebEnginePage__WebAction = QWebEnginePage__WebAction(35)
	QWebEnginePage__ToggleStrikethrough        QWebEnginePage__WebAction = QWebEnginePage__WebAction(36)
	QWebEnginePage__AlignLeft                  QWebEnginePage__WebAction = QWebEnginePage__WebAction(37)
	QWebEnginePage__AlignCenter                QWebEnginePage__WebAction = QWebEnginePage__WebAction(38)
	QWebEnginePage__AlignRight                 QWebEnginePage__WebAction = QWebEnginePage__WebAction(39)
	QWebEnginePage__AlignJustified             QWebEnginePage__WebAction = QWebEnginePage__WebAction(40)
	QWebEnginePage__Indent                     QWebEnginePage__WebAction = QWebEnginePage__WebAction(41)
	QWebEnginePage__Outdent                    QWebEnginePage__WebAction = QWebEnginePage__WebAction(42)
	QWebEnginePage__InsertOrderedList          QWebEnginePage__WebAction = QWebEnginePage__WebAction(43)
	QWebEnginePage__InsertUnorderedList        QWebEnginePage__WebAction = QWebEnginePage__WebAction(44)
	QWebEnginePage__WebActionCount             QWebEnginePage__WebAction = QWebEnginePage__WebAction(45)
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

//go:generate stringer -type=QWebEnginePage__PermissionPolicy
//QWebEnginePage::PermissionPolicy
type QWebEnginePage__PermissionPolicy int64

const (
	QWebEnginePage__PermissionUnknown       QWebEnginePage__PermissionPolicy = QWebEnginePage__PermissionPolicy(0)
	QWebEnginePage__PermissionGrantedByUser QWebEnginePage__PermissionPolicy = QWebEnginePage__PermissionPolicy(1)
	QWebEnginePage__PermissionDeniedByUser  QWebEnginePage__PermissionPolicy = QWebEnginePage__PermissionPolicy(2)
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

//go:generate stringer -type=QWebEnginePage__FindFlag
//QWebEnginePage::FindFlag
type QWebEnginePage__FindFlag int64

const (
	QWebEnginePage__FindBackward        QWebEnginePage__FindFlag = QWebEnginePage__FindFlag(1)
	QWebEnginePage__FindCaseSensitively QWebEnginePage__FindFlag = QWebEnginePage__FindFlag(2)
)

//go:generate stringer -type=QWebEnginePage__Feature
//QWebEnginePage::Feature
type QWebEnginePage__Feature int64

const (
	QWebEnginePage__Notifications            QWebEnginePage__Feature = QWebEnginePage__Feature(0)
	QWebEnginePage__Geolocation              QWebEnginePage__Feature = QWebEnginePage__Feature(1)
	QWebEnginePage__MediaAudioCapture        QWebEnginePage__Feature = QWebEnginePage__Feature(2)
	QWebEnginePage__MediaVideoCapture        QWebEnginePage__Feature = QWebEnginePage__Feature(3)
	QWebEnginePage__MediaAudioVideoCapture   QWebEnginePage__Feature = QWebEnginePage__Feature(4)
	QWebEnginePage__MouseLock                QWebEnginePage__Feature = QWebEnginePage__Feature(5)
	QWebEnginePage__DesktopVideoCapture      QWebEnginePage__Feature = QWebEnginePage__Feature(6)
	QWebEnginePage__DesktopAudioVideoCapture QWebEnginePage__Feature = QWebEnginePage__Feature(7)
)

//export callbackQWebEnginePage_AudioMutedChanged
func callbackQWebEnginePage_AudioMutedChanged(ptr unsafe.Pointer, muted C.char) {
	if signal := qt.GetSignal(ptr, "audioMutedChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(muted) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectAudioMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "audioMutedChanged") {
			C.QWebEnginePage_ConnectAudioMutedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "audioMutedChanged"); signal != nil {
			f := func(muted bool) {
				(*(*func(bool))(signal))(muted)
				f(muted)
			}
			qt.ConnectSignal(ptr.Pointer(), "audioMutedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "audioMutedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectAudioMutedChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectAudioMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "audioMutedChanged")
	}
}

func (ptr *QWebEnginePage) AudioMutedChanged(muted bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_AudioMutedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(muted))))
	}
}

func (ptr *QWebEnginePage) BackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QWebEnginePage_BackgroundColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) ContentsSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QWebEnginePage_ContentsSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_ContentsSizeChanged
func callbackQWebEnginePage_ContentsSizeChanged(ptr unsafe.Pointer, size unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsSizeChanged"); signal != nil {
		(*(*func(*core.QSizeF))(signal))(core.NewQSizeFFromPointer(size))
	}

}

func (ptr *QWebEnginePage) ConnectContentsSizeChanged(f func(size *core.QSizeF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsSizeChanged") {
			C.QWebEnginePage_ConnectContentsSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsSizeChanged"); signal != nil {
			f := func(size *core.QSizeF) {
				(*(*func(*core.QSizeF))(signal))(size)
				f(size)
			}
			qt.ConnectSignal(ptr.Pointer(), "contentsSizeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsSizeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsSizeChanged")
	}
}

func (ptr *QWebEnginePage) ContentsSizeChanged(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ContentsSizeChanged(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QWebEnginePage) HasSelection() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_HasSelection(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebEnginePage) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWebEnginePage_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_IconChanged
func callbackQWebEnginePage_IconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconChanged"); signal != nil {
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

func (ptr *QWebEnginePage) ConnectIconChanged(f func(icon *gui.QIcon)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconChanged") {
			C.QWebEnginePage_ConnectIconChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconChanged"); signal != nil {
			f := func(icon *gui.QIcon) {
				(*(*func(*gui.QIcon))(signal))(icon)
				f(icon)
			}
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconChanged")
	}
}

func (ptr *QWebEnginePage) IconChanged(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_IconChanged(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWebEnginePage) IconUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEnginePage_IconUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_IconUrlChanged
func callbackQWebEnginePage_IconUrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconUrlChanged"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEnginePage) ConnectIconUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconUrlChanged") {
			C.QWebEnginePage_ConnectIconUrlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconUrlChanged"); signal != nil {
			f := func(url *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "iconUrlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconUrlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectIconUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectIconUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconUrlChanged")
	}
}

func (ptr *QWebEnginePage) IconUrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_IconUrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) IsAudioMuted() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_IsAudioMuted(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebEnginePage) RecentlyAudible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_RecentlyAudible(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebEnginePage_RecentlyAudibleChanged
func callbackQWebEnginePage_RecentlyAudibleChanged(ptr unsafe.Pointer, recentlyAudible C.char) {
	if signal := qt.GetSignal(ptr, "recentlyAudibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(recentlyAudible) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectRecentlyAudibleChanged(f func(recentlyAudible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "recentlyAudibleChanged") {
			C.QWebEnginePage_ConnectRecentlyAudibleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "recentlyAudibleChanged"); signal != nil {
			f := func(recentlyAudible bool) {
				(*(*func(bool))(signal))(recentlyAudible)
				f(recentlyAudible)
			}
			qt.ConnectSignal(ptr.Pointer(), "recentlyAudibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "recentlyAudibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectRecentlyAudibleChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectRecentlyAudibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "recentlyAudibleChanged")
	}
}

func (ptr *QWebEnginePage) RecentlyAudibleChanged(recentlyAudible bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_RecentlyAudibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(recentlyAudible))))
	}
}

func (ptr *QWebEnginePage) RequestedUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEnginePage_RequestedUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) ScrollPosition() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QWebEnginePage_ScrollPosition(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_ScrollPositionChanged
func callbackQWebEnginePage_ScrollPositionChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollPositionChanged"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(position))
	}

}

func (ptr *QWebEnginePage) ConnectScrollPositionChanged(f func(position *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "scrollPositionChanged") {
			C.QWebEnginePage_ConnectScrollPositionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "scrollPositionChanged"); signal != nil {
			f := func(position *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(position)
				f(position)
			}
			qt.ConnectSignal(ptr.Pointer(), "scrollPositionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scrollPositionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectScrollPositionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectScrollPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "scrollPositionChanged")
	}
}

func (ptr *QWebEnginePage) ScrollPositionChanged(position core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ScrollPositionChanged(ptr.Pointer(), core.PointerFromQPointF(position))
	}
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

func (ptr *QWebEnginePage) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) SetZoomFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebEnginePage) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEnginePage_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEnginePage) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEnginePage_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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

func (ptr *QWebEnginePage) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEnginePage___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEnginePage) __children_newList() unsafe.Pointer {
	return C.QWebEnginePage___children_newList(ptr.Pointer())
}

func (ptr *QWebEnginePage) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEnginePage___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEnginePage) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEnginePage___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEnginePage) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEnginePage___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEnginePage) __findChildren_newList() unsafe.Pointer {
	return C.QWebEnginePage___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEnginePage) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEnginePage___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEnginePage) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEnginePage___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEnginePage) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEnginePage___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEnginePage) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEnginePage___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEnginePage_ChildEvent
func callbackQWebEnginePage_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEnginePage_ConnectNotify
func callbackQWebEnginePage_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEnginePageFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEnginePage) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEnginePage_CustomEvent
func callbackQWebEnginePage_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEnginePage_DeleteLater
func callbackQWebEnginePage_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEnginePageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEnginePage) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEnginePage_Destroyed
func callbackQWebEnginePage_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEnginePage_DisconnectNotify
func callbackQWebEnginePage_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEnginePageFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEnginePage) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEnginePage_Event
func callbackQWebEnginePage_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEnginePage) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEnginePage_EventFilter
func callbackQWebEnginePage_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEnginePage) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEnginePage_ObjectNameChanged
func callbackQWebEnginePage_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEnginePage_TimerEvent
func callbackQWebEnginePage_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEnginePageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEnginePage) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQWebEngineProfileFromPointer(ptr unsafe.Pointer) (n *QWebEngineProfile) {
	n = new(QWebEngineProfile)
	n.SetPointer(ptr)
	return
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

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineProfile___clearVisitedLinks_urls_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile___clearVisitedLinks_urls_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_newList() unsafe.Pointer {
	return C.QWebEngineProfile___clearVisitedLinks_urls_newList(ptr.Pointer())
}

func (ptr *QWebEngineProfile) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineProfile___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineProfile) __children_newList() unsafe.Pointer {
	return C.QWebEngineProfile___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineProfile) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineProfile___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineProfile) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineProfile___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineProfile) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineProfile___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineProfile) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineProfile___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineProfile) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineProfile___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineProfile) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineProfile___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineProfile) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineProfile___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineProfile) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineProfile___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineProfile_ChildEvent
func callbackQWebEngineProfile_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineProfile_ConnectNotify
func callbackQWebEngineProfile_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineProfileFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineProfile_CustomEvent
func callbackQWebEngineProfile_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineProfile_DeleteLater
func callbackQWebEngineProfile_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineProfileFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineProfile) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineProfile_Destroyed
func callbackQWebEngineProfile_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineProfile_DisconnectNotify
func callbackQWebEngineProfile_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineProfileFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineProfile_Event
func callbackQWebEngineProfile_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineProfileFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineProfile_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEngineProfile_EventFilter
func callbackQWebEngineProfile_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineProfileFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineProfile_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineProfile_ObjectNameChanged
func callbackQWebEngineProfile_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineProfile_TimerEvent
func callbackQWebEngineProfile_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineProfileFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWebEngineQuotaRequest struct {
	ptr unsafe.Pointer
}

type QWebEngineQuotaRequest_ITF interface {
	QWebEngineQuotaRequest_PTR() *QWebEngineQuotaRequest
}

func (ptr *QWebEngineQuotaRequest) QWebEngineQuotaRequest_PTR() *QWebEngineQuotaRequest {
	return ptr
}

func (ptr *QWebEngineQuotaRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineQuotaRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineQuotaRequest(ptr QWebEngineQuotaRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineQuotaRequest_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineQuotaRequestFromPointer(ptr unsafe.Pointer) (n *QWebEngineQuotaRequest) {
	n = new(QWebEngineQuotaRequest)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineQuotaRequest) DestroyQWebEngineQuotaRequest() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineQuotaRequest) Accept() {
	if ptr.Pointer() != nil {
		C.QWebEngineQuotaRequest_Accept(ptr.Pointer())
	}
}

func (ptr *QWebEngineQuotaRequest) Origin() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineQuotaRequest_Origin(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineQuotaRequest) Reject() {
	if ptr.Pointer() != nil {
		C.QWebEngineQuotaRequest_Reject(ptr.Pointer())
	}
}

func (ptr *QWebEngineQuotaRequest) RequestedSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebEngineQuotaRequest_RequestedSize(ptr.Pointer()))
	}
	return 0
}

type QWebEngineRegisterProtocolHandlerRequest struct {
	ptr unsafe.Pointer
}

type QWebEngineRegisterProtocolHandlerRequest_ITF interface {
	QWebEngineRegisterProtocolHandlerRequest_PTR() *QWebEngineRegisterProtocolHandlerRequest
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) QWebEngineRegisterProtocolHandlerRequest_PTR() *QWebEngineRegisterProtocolHandlerRequest {
	return ptr
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineRegisterProtocolHandlerRequest(ptr QWebEngineRegisterProtocolHandlerRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineRegisterProtocolHandlerRequest_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineRegisterProtocolHandlerRequestFromPointer(ptr unsafe.Pointer) (n *QWebEngineRegisterProtocolHandlerRequest) {
	n = new(QWebEngineRegisterProtocolHandlerRequest)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) DestroyQWebEngineRegisterProtocolHandlerRequest() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Accept() {
	if ptr.Pointer() != nil {
		C.QWebEngineRegisterProtocolHandlerRequest_Accept(ptr.Pointer())
	}
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Origin() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineRegisterProtocolHandlerRequest_Origin(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Reject() {
	if ptr.Pointer() != nil {
		C.QWebEngineRegisterProtocolHandlerRequest_Reject(ptr.Pointer())
	}
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Scheme() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineRegisterProtocolHandlerRequest_Scheme(ptr.Pointer()))
	}
	return ""
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

func NewQWebEngineScriptFromPointer(ptr unsafe.Pointer) (n *QWebEngineScript) {
	n = new(QWebEngineScript)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineScript) DestroyQWebEngineScript() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQWebEngineScriptCollectionFromPointer(ptr unsafe.Pointer) (n *QWebEngineScriptCollection) {
	n = new(QWebEngineScriptCollection)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineScriptCollection) DestroyQWebEngineScriptCollection() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQWebEngineSettingsFromPointer(ptr unsafe.Pointer) (n *QWebEngineSettings) {
	n = new(QWebEngineSettings)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineSettings) DestroyQWebEngineSettings() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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

//go:generate stringer -type=QWebEngineSettings__WebAttribute
//QWebEngineSettings::WebAttribute
type QWebEngineSettings__WebAttribute int64

const (
	QWebEngineSettings__AutoLoadImages                      QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(0)
	QWebEngineSettings__JavascriptEnabled                   QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(1)
	QWebEngineSettings__JavascriptCanOpenWindows            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(2)
	QWebEngineSettings__JavascriptCanAccessClipboard        QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(3)
	QWebEngineSettings__LinksIncludedInFocusChain           QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(4)
	QWebEngineSettings__LocalStorageEnabled                 QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(5)
	QWebEngineSettings__LocalContentCanAccessRemoteUrls     QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(6)
	QWebEngineSettings__XSSAuditingEnabled                  QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(7)
	QWebEngineSettings__SpatialNavigationEnabled            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(8)
	QWebEngineSettings__LocalContentCanAccessFileUrls       QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(9)
	QWebEngineSettings__HyperlinkAuditingEnabled            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(10)
	QWebEngineSettings__ScrollAnimatorEnabled               QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(11)
	QWebEngineSettings__ErrorPageEnabled                    QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(12)
	QWebEngineSettings__PluginsEnabled                      QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(13)
	QWebEngineSettings__FullScreenSupportEnabled            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(14)
	QWebEngineSettings__ScreenCaptureEnabled                QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(15)
	QWebEngineSettings__WebGLEnabled                        QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(16)
	QWebEngineSettings__Accelerated2dCanvasEnabled          QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(17)
	QWebEngineSettings__AutoLoadIconsForPage                QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(18)
	QWebEngineSettings__TouchIconsEnabled                   QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(19)
	QWebEngineSettings__FocusOnNavigationEnabled            QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(20)
	QWebEngineSettings__PrintElementBackgrounds             QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(21)
	QWebEngineSettings__AllowRunningInsecureContent         QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(22)
	QWebEngineSettings__AllowGeolocationOnInsecureOrigins   QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(23)
	QWebEngineSettings__AllowWindowActivationFromJavaScript QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(24)
	QWebEngineSettings__ShowScrollBars                      QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(25)
	QWebEngineSettings__PlaybackRequiresUserGesture         QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(26)
	QWebEngineSettings__WebRTCPublicInterfacesOnly          QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(27)
	QWebEngineSettings__JavascriptCanPaste                  QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(28)
	QWebEngineSettings__DnsPrefetchEnabled                  QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(29)
	QWebEngineSettings__PdfViewerEnabled                    QWebEngineSettings__WebAttribute = QWebEngineSettings__WebAttribute(30)
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

//go:generate stringer -type=QWebEngineSettings__UnknownUrlSchemePolicy
//QWebEngineSettings::UnknownUrlSchemePolicy
type QWebEngineSettings__UnknownUrlSchemePolicy int64

const (
	QWebEngineSettings__DisallowUnknownUrlSchemes                 QWebEngineSettings__UnknownUrlSchemePolicy = QWebEngineSettings__UnknownUrlSchemePolicy(1)
	QWebEngineSettings__AllowUnknownUrlSchemesFromUserInteraction QWebEngineSettings__UnknownUrlSchemePolicy = QWebEngineSettings__UnknownUrlSchemePolicy(2)
	QWebEngineSettings__AllowAllUnknownUrlSchemes                 QWebEngineSettings__UnknownUrlSchemePolicy = QWebEngineSettings__UnknownUrlSchemePolicy(3)
)

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

func NewQWebEngineUrlRequestInfoFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestInfo) {
	n = new(QWebEngineUrlRequestInfo)
	n.SetPointer(ptr)
	return
}

func (ptr *QWebEngineUrlRequestInfo) DestroyQWebEngineUrlRequestInfo() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

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
	QWebEngineUrlRequestInfo__ResourceTypeUnknown        QWebEngineUrlRequestInfo__ResourceType = QWebEngineUrlRequestInfo__ResourceType(255)
)

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

func (ptr *QWebEngineUrlRequestInfo) Block(shouldBlock bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInfo_Block(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(shouldBlock))))
	}
}

func (ptr *QWebEngineUrlRequestInfo) FirstPartyUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_FirstPartyUrl(ptr.Pointer()))
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
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestInfo_RequestMethod(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_RequestUrl(ptr.Pointer()))
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

func NewQWebEngineUrlRequestInterceptorFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestInterceptor) {
	n = new(QWebEngineUrlRequestInterceptor)
	n.SetPointer(ptr)
	return
}
func NewQWebEngineUrlRequestInterceptor2(p core.QObject_ITF) *QWebEngineUrlRequestInterceptor {
	tmpValue := NewQWebEngineUrlRequestInterceptorFromPointer(C.QWebEngineUrlRequestInterceptor_NewQWebEngineUrlRequestInterceptor2(core.PointerFromQObject(p)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineUrlRequestInterceptor_InterceptRequest
func callbackQWebEngineUrlRequestInterceptor_InterceptRequest(ptr unsafe.Pointer, info unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "interceptRequest"); signal != nil {
		(*(*func(*QWebEngineUrlRequestInfo))(signal))(NewQWebEngineUrlRequestInfoFromPointer(info))
	}

}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectInterceptRequest(f func(info *QWebEngineUrlRequestInfo)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "interceptRequest"); signal != nil {
			f := func(info *QWebEngineUrlRequestInfo) {
				(*(*func(*QWebEngineUrlRequestInfo))(signal))(info)
				f(info)
			}
			qt.ConnectSignal(ptr.Pointer(), "interceptRequest", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "interceptRequest", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectInterceptRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "interceptRequest")
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) InterceptRequest(info QWebEngineUrlRequestInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_InterceptRequest(ptr.Pointer(), PointerFromQWebEngineUrlRequestInfo(info))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestInterceptor___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) __children_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestInterceptor___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestInterceptor) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestInterceptor___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestInterceptor___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestInterceptor___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestInterceptor___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestInterceptor___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineUrlRequestInterceptor___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestInterceptor) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestInterceptor___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInterceptor) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineUrlRequestInterceptor___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineUrlRequestInterceptor_ChildEvent
func callbackQWebEngineUrlRequestInterceptor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_ConnectNotify
func callbackQWebEngineUrlRequestInterceptor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_CustomEvent
func callbackQWebEngineUrlRequestInterceptor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_DeleteLater
func callbackQWebEngineUrlRequestInterceptor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineUrlRequestInterceptor_Destroyed
func callbackQWebEngineUrlRequestInterceptor_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineUrlRequestInterceptor_DisconnectNotify
func callbackQWebEngineUrlRequestInterceptor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestInterceptor_Event
func callbackQWebEngineUrlRequestInterceptor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineUrlRequestInterceptor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineUrlRequestInterceptor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestInterceptor_EventFilter
func callbackQWebEngineUrlRequestInterceptor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineUrlRequestInterceptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineUrlRequestInterceptor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestInterceptor_ObjectNameChanged
func callbackQWebEngineUrlRequestInterceptor_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineUrlRequestInterceptor_TimerEvent
func callbackQWebEngineUrlRequestInterceptor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestInterceptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestInterceptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestInterceptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQWebEngineUrlRequestJobFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestJob) {
	n = new(QWebEngineUrlRequestJob)
	n.SetPointer(ptr)
	return
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

func (ptr *QWebEngineUrlRequestJob) Initiator() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineUrlRequestJob_Initiator(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
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

func (ptr *QWebEngineUrlRequestJob) RequestHeaders() map[*core.QByteArray]*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) map[*core.QByteArray]*core.QByteArray {
			out := make(map[*core.QByteArray]*core.QByteArray, int(l.len))
			tmpList := NewQWebEngineUrlRequestJobFromPointer(l.data)
			for i, v := range tmpList.__requestHeaders_keyList() {
				out[v] = tmpList.__requestHeaders_atList(v, i)
			}
			return out
		}(C.QWebEngineUrlRequestJob_RequestHeaders(ptr.Pointer()))
	}
	return make(map[*core.QByteArray]*core.QByteArray, 0)
}

func (ptr *QWebEngineUrlRequestJob) RequestMethod() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestJob_RequestMethod(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineUrlRequestJob_RequestUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_atList(v core.QByteArray_ITF, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestJob___requestHeaders_atList(ptr.Pointer(), core.PointerFromQByteArray(v), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_setList(key core.QByteArray_ITF, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob___requestHeaders_setList(ptr.Pointer(), core.PointerFromQByteArray(key), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob___requestHeaders_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_keyList() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQWebEngineUrlRequestJobFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____requestHeaders_keyList_atList(i)
			}
			return out
		}(C.QWebEngineUrlRequestJob___requestHeaders_keyList(ptr.Pointer()))
	}
	return make([]*core.QByteArray, 0)
}

func (ptr *QWebEngineUrlRequestJob) ____requestHeaders_keyList_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestJob_____requestHeaders_keyList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) ____requestHeaders_keyList_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_____requestHeaders_keyList_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) ____requestHeaders_keyList_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob_____requestHeaders_keyList_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestJob) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestJob___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) __children_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestJob) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestJob___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestJob___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestJob___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineUrlRequestJob) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlRequestJob___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlRequestJob) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineUrlRequestJob___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineUrlRequestJob_ChildEvent
func callbackQWebEngineUrlRequestJob_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_ConnectNotify
func callbackQWebEngineUrlRequestJob_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestJob_CustomEvent
func callbackQWebEngineUrlRequestJob_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlRequestJob_DeleteLater
func callbackQWebEngineUrlRequestJob_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlRequestJob) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineUrlRequestJob_Destroyed
func callbackQWebEngineUrlRequestJob_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineUrlRequestJob_DisconnectNotify
func callbackQWebEngineUrlRequestJob_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlRequestJob) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlRequestJob_Event
func callbackQWebEngineUrlRequestJob_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestJobFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineUrlRequestJob) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineUrlRequestJob_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestJob_EventFilter
func callbackQWebEngineUrlRequestJob_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlRequestJobFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineUrlRequestJob) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineUrlRequestJob_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineUrlRequestJob_ObjectNameChanged
func callbackQWebEngineUrlRequestJob_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineUrlRequestJob_TimerEvent
func callbackQWebEngineUrlRequestJob_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlRequestJobFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlRequestJob) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlRequestJob_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWebEngineUrlScheme struct {
	ptr unsafe.Pointer
}

type QWebEngineUrlScheme_ITF interface {
	QWebEngineUrlScheme_PTR() *QWebEngineUrlScheme
}

func (ptr *QWebEngineUrlScheme) QWebEngineUrlScheme_PTR() *QWebEngineUrlScheme {
	return ptr
}

func (ptr *QWebEngineUrlScheme) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineUrlScheme) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineUrlScheme(ptr QWebEngineUrlScheme_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlScheme_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineUrlSchemeFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlScheme) {
	n = new(QWebEngineUrlScheme)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QWebEngineUrlScheme__Syntax
//QWebEngineUrlScheme::Syntax
type QWebEngineUrlScheme__Syntax int64

const (
	QWebEngineUrlScheme__HostPortAndUserInformation QWebEngineUrlScheme__Syntax = QWebEngineUrlScheme__Syntax(0)
	QWebEngineUrlScheme__HostAndPort                QWebEngineUrlScheme__Syntax = QWebEngineUrlScheme__Syntax(1)
	QWebEngineUrlScheme__Host                       QWebEngineUrlScheme__Syntax = QWebEngineUrlScheme__Syntax(2)
	QWebEngineUrlScheme__Path                       QWebEngineUrlScheme__Syntax = QWebEngineUrlScheme__Syntax(3)
)

//go:generate stringer -type=QWebEngineUrlScheme__SpecialPort
//QWebEngineUrlScheme::SpecialPort
type QWebEngineUrlScheme__SpecialPort int64

const (
	QWebEngineUrlScheme__PortUnspecified QWebEngineUrlScheme__SpecialPort = QWebEngineUrlScheme__SpecialPort(-1)
)

//go:generate stringer -type=QWebEngineUrlScheme__Flag
//QWebEngineUrlScheme::Flag
type QWebEngineUrlScheme__Flag int64

const (
	QWebEngineUrlScheme__SecureScheme                 QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x1)
	QWebEngineUrlScheme__LocalScheme                  QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x2)
	QWebEngineUrlScheme__LocalAccessAllowed           QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x4)
	QWebEngineUrlScheme__NoAccessAllowed              QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x8)
	QWebEngineUrlScheme__ServiceWorkersAllowed        QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x10)
	QWebEngineUrlScheme__ViewSourceAllowed            QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x20)
	QWebEngineUrlScheme__ContentSecurityPolicyIgnored QWebEngineUrlScheme__Flag = QWebEngineUrlScheme__Flag(0x40)
)

func NewQWebEngineUrlScheme() *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme())
	runtime.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func NewQWebEngineUrlScheme2(name core.QByteArray_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme2(core.PointerFromQByteArray(name)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func NewQWebEngineUrlScheme3(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme3(PointerFromQWebEngineUrlScheme(that)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func NewQWebEngineUrlScheme4(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme4(PointerFromQWebEngineUrlScheme(that)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func (ptr *QWebEngineUrlScheme) DefaultPort() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineUrlScheme_DefaultPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebEngineUrlScheme) Flags() QWebEngineUrlScheme__Flag {
	if ptr.Pointer() != nil {
		return QWebEngineUrlScheme__Flag(C.QWebEngineUrlScheme_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineUrlScheme) Name() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlScheme_Name(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func QWebEngineUrlScheme_RegisterScheme(scheme QWebEngineUrlScheme_ITF) {
	C.QWebEngineUrlScheme_QWebEngineUrlScheme_RegisterScheme(PointerFromQWebEngineUrlScheme(scheme))
}

func (ptr *QWebEngineUrlScheme) RegisterScheme(scheme QWebEngineUrlScheme_ITF) {
	C.QWebEngineUrlScheme_QWebEngineUrlScheme_RegisterScheme(PointerFromQWebEngineUrlScheme(scheme))
}

func QWebEngineUrlScheme_SchemeByName(name core.QByteArray_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_QWebEngineUrlScheme_SchemeByName(core.PointerFromQByteArray(name)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func (ptr *QWebEngineUrlScheme) SchemeByName(name core.QByteArray_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_QWebEngineUrlScheme_SchemeByName(core.PointerFromQByteArray(name)))
	runtime.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func (ptr *QWebEngineUrlScheme) SetDefaultPort(newValue int) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlScheme_SetDefaultPort(ptr.Pointer(), C.int(int32(newValue)))
	}
}

func (ptr *QWebEngineUrlScheme) SetFlags(newValue QWebEngineUrlScheme__Flag) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlScheme_SetFlags(ptr.Pointer(), C.longlong(newValue))
	}
}

func (ptr *QWebEngineUrlScheme) SetName(newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlScheme_SetName(ptr.Pointer(), core.PointerFromQByteArray(newValue))
	}
}

func (ptr *QWebEngineUrlScheme) SetSyntax(newValue QWebEngineUrlScheme__Syntax) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlScheme_SetSyntax(ptr.Pointer(), C.longlong(newValue))
	}
}

func (ptr *QWebEngineUrlScheme) DestroyQWebEngineUrlScheme() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlScheme_DestroyQWebEngineUrlScheme(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQWebEngineUrlSchemeHandlerFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlSchemeHandler) {
	n = new(QWebEngineUrlSchemeHandler)
	n.SetPointer(ptr)
	return
}
func NewQWebEngineUrlSchemeHandler(parent core.QObject_ITF) *QWebEngineUrlSchemeHandler {
	tmpValue := NewQWebEngineUrlSchemeHandlerFromPointer(C.QWebEngineUrlSchemeHandler_NewQWebEngineUrlSchemeHandler(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineUrlSchemeHandler_RequestStarted
func callbackQWebEngineUrlSchemeHandler_RequestStarted(ptr unsafe.Pointer, request unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestStarted"); signal != nil {
		(*(*func(*QWebEngineUrlRequestJob))(signal))(NewQWebEngineUrlRequestJobFromPointer(request))
	}

}

func (ptr *QWebEngineUrlSchemeHandler) ConnectRequestStarted(f func(request *QWebEngineUrlRequestJob)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestStarted"); signal != nil {
			f := func(request *QWebEngineUrlRequestJob) {
				(*(*func(*QWebEngineUrlRequestJob))(signal))(request)
				f(request)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestStarted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectRequestStarted() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestStarted")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) RequestStarted(request QWebEngineUrlRequestJob_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_RequestStarted(ptr.Pointer(), PointerFromQWebEngineUrlRequestJob(request))
	}
}

//export callbackQWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler
func callbackQWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebEngineUrlSchemeHandler"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DestroyQWebEngineUrlSchemeHandlerDefault()
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectDestroyQWebEngineUrlSchemeHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebEngineUrlSchemeHandler"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineUrlSchemeHandler", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineUrlSchemeHandler", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectDestroyQWebEngineUrlSchemeHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebEngineUrlSchemeHandler")
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandler() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlSchemeHandler___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) __children_newList() unsafe.Pointer {
	return C.QWebEngineUrlSchemeHandler___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlSchemeHandler) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlSchemeHandler___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineUrlSchemeHandler___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlSchemeHandler___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineUrlSchemeHandler___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlSchemeHandler___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineUrlSchemeHandler___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineUrlSchemeHandler) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineUrlSchemeHandler___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlSchemeHandler) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineUrlSchemeHandler___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineUrlSchemeHandler_ChildEvent
func callbackQWebEngineUrlSchemeHandler_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_ConnectNotify
func callbackQWebEngineUrlSchemeHandler_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlSchemeHandler_CustomEvent
func callbackQWebEngineUrlSchemeHandler_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineUrlSchemeHandler_DeleteLater
func callbackQWebEngineUrlSchemeHandler_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineUrlSchemeHandler_Destroyed
func callbackQWebEngineUrlSchemeHandler_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineUrlSchemeHandler_DisconnectNotify
func callbackQWebEngineUrlSchemeHandler_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineUrlSchemeHandler_Event
func callbackQWebEngineUrlSchemeHandler_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebEngineUrlSchemeHandler) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineUrlSchemeHandler_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebEngineUrlSchemeHandler_EventFilter
func callbackQWebEngineUrlSchemeHandler_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineUrlSchemeHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineUrlSchemeHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineUrlSchemeHandler_ObjectNameChanged
func callbackQWebEngineUrlSchemeHandler_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineUrlSchemeHandler_TimerEvent
func callbackQWebEngineUrlSchemeHandler_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineUrlSchemeHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineUrlSchemeHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineUrlSchemeHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQWebEngineViewFromPointer(ptr unsafe.Pointer) (n *QWebEngineView) {
	n = new(QWebEngineView)
	n.SetPointer(ptr)
	return
}
func (ptr *QWebEngineView) HasSelection() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_HasSelection(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebEngineView) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWebEngineView_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_IconChanged
func callbackQWebEngineView_IconChanged(ptr unsafe.Pointer, vqi unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconChanged"); signal != nil {
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(vqi))
	}

}

func (ptr *QWebEngineView) ConnectIconChanged(f func(vqi *gui.QIcon)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconChanged") {
			C.QWebEngineView_ConnectIconChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconChanged"); signal != nil {
			f := func(vqi *gui.QIcon) {
				(*(*func(*gui.QIcon))(signal))(vqi)
				f(vqi)
			}
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconChanged")
	}
}

func (ptr *QWebEngineView) IconChanged(vqi gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_IconChanged(ptr.Pointer(), gui.PointerFromQIcon(vqi))
	}
}

func (ptr *QWebEngineView) IconUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineView_IconUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_IconUrlChanged
func callbackQWebEngineView_IconUrlChanged(ptr unsafe.Pointer, vqu unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconUrlChanged"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(vqu))
	}

}

func (ptr *QWebEngineView) ConnectIconUrlChanged(f func(vqu *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconUrlChanged") {
			C.QWebEngineView_ConnectIconUrlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconUrlChanged"); signal != nil {
			f := func(vqu *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(vqu)
				f(vqu)
			}
			qt.ConnectSignal(ptr.Pointer(), "iconUrlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconUrlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectIconUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectIconUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconUrlChanged")
	}
}

func (ptr *QWebEngineView) IconUrlChanged(vqu core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_IconUrlChanged(ptr.Pointer(), core.PointerFromQUrl(vqu))
	}
}

func (ptr *QWebEngineView) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineView_SelectedText(ptr.Pointer()))
	}
	return ""
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

func (ptr *QWebEngineView) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebEngineView) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineView_Url(ptr.Pointer()))
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

func (ptr *QWebEngineView) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QWebEngineView___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebEngineView) __actions_newList() unsafe.Pointer {
	return C.QWebEngineView___actions_newList(ptr.Pointer())
}

func (ptr *QWebEngineView) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QWebEngineView___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebEngineView) __addActions_actions_newList() unsafe.Pointer {
	return C.QWebEngineView___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QWebEngineView) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QWebEngineView___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebEngineView) __insertActions_actions_newList() unsafe.Pointer {
	return C.QWebEngineView___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QWebEngineView) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineView___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineView) __children_newList() unsafe.Pointer {
	return C.QWebEngineView___children_newList(ptr.Pointer())
}

func (ptr *QWebEngineView) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineView___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebEngineView) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebEngineView___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebEngineView) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineView___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineView) __findChildren_newList() unsafe.Pointer {
	return C.QWebEngineView___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebEngineView) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineView___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineView) __findChildren_newList3() unsafe.Pointer {
	return C.QWebEngineView___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebEngineView) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebEngineView___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineView) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebEngineView) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWebEngineView___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQWebEngineView_ActionEvent
func callbackQWebEngineView_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*gui.QActionEvent))(signal))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebEngineView_ChangeEvent
func callbackQWebEngineView_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_Close
func callbackQWebEngineView_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).CloseDefault())))
}

func (ptr *QWebEngineView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebEngineView_CloseEvent
func callbackQWebEngineView_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebEngineView_ContextMenuEvent
func callbackQWebEngineView_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQWebEngineView_CustomContextMenuRequested
func callbackQWebEngineView_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQWebEngineView_DragEnterEvent
func callbackQWebEngineView_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQWebEngineView_DragLeaveEvent
func callbackQWebEngineView_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQWebEngineView_DragMoveEvent
func callbackQWebEngineView_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQWebEngineView_DropEvent
func callbackQWebEngineView_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQWebEngineView_EnterEvent
func callbackQWebEngineView_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_Event
func callbackQWebEngineView_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineView) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineView_FocusInEvent
func callbackQWebEngineView_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebEngineView_FocusNextPrevChild
func callbackQWebEngineView_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QWebEngineView) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQWebEngineView_FocusOutEvent
func callbackQWebEngineView_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebEngineView_HasHeightForWidth
func callbackQWebEngineView_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QWebEngineView) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebEngineView_HeightForWidth
func callbackQWebEngineView_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQWebEngineViewFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QWebEngineView) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineView_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQWebEngineView_Hide
func callbackQWebEngineView_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebEngineView) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_HideEvent
func callbackQWebEngineView_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWebEngineView_InitPainter
func callbackQWebEngineView_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*gui.QPainter))(signal))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQWebEngineViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QWebEngineView) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQWebEngineView_InputMethodEvent
func callbackQWebEngineView_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQWebEngineView_InputMethodQuery
func callbackQWebEngineView_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQWebEngineViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QWebEngineView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QWebEngineView_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_KeyPressEvent
func callbackQWebEngineView_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebEngineView_KeyReleaseEvent
func callbackQWebEngineView_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebEngineView_LeaveEvent
func callbackQWebEngineView_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_Lower
func callbackQWebEngineView_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebEngineView) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_Metric
func callbackQWebEngineView_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQWebEngineViewFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QWebEngineView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebEngineView_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQWebEngineView_MinimumSizeHint
func callbackQWebEngineView_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQWebEngineViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebEngineView) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWebEngineView_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_MouseDoubleClickEvent
func callbackQWebEngineView_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MouseMoveEvent
func callbackQWebEngineView_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MousePressEvent
func callbackQWebEngineView_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MouseReleaseEvent
func callbackQWebEngineView_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebEngineView_MoveEvent
func callbackQWebEngineView_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*gui.QMoveEvent))(signal))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebEngineView_NativeEvent
func callbackQWebEngineView_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QByteArray, unsafe.Pointer, *int) bool)(signal))(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *QWebEngineView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.QWebEngineView_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackQWebEngineView_PaintEngine
func callbackQWebEngineView_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQWebEngineViewFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWebEngineView) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebEngineView_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEngineView_PaintEvent
func callbackQWebEngineView_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQWebEngineView_Raise
func callbackQWebEngineView_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebEngineView) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_Repaint
func callbackQWebEngineView_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebEngineView) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ResizeEvent
func callbackQWebEngineView_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQWebEngineView_SetDisabled
func callbackQWebEngineView_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QWebEngineView) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQWebEngineView_SetEnabled
func callbackQWebEngineView_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebEngineView) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebEngineView_SetFocus2
func callbackQWebEngineView_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebEngineView) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebEngineView_SetHidden
func callbackQWebEngineView_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QWebEngineView) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQWebEngineView_SetStyleSheet
func callbackQWebEngineView_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewQWebEngineViewFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QWebEngineView) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QWebEngineView_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQWebEngineView_SetVisible
func callbackQWebEngineView_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QWebEngineView) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebEngineView_SetWindowModified
func callbackQWebEngineView_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQWebEngineViewFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebEngineView) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebEngineView_SetWindowTitle
func callbackQWebEngineView_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQWebEngineViewFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QWebEngineView) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWebEngineView_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQWebEngineView_Show
func callbackQWebEngineView_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebEngineView) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowEvent
func callbackQWebEngineView_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWebEngineView_ShowFullScreen
func callbackQWebEngineView_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebEngineView) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowMaximized
func callbackQWebEngineView_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebEngineView) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowMinimized
func callbackQWebEngineView_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebEngineView) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_ShowNormal
func callbackQWebEngineView_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebEngineView) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_SizeHint
func callbackQWebEngineView_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQWebEngineViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebEngineView) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWebEngineView_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_TabletEvent
func callbackQWebEngineView_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*gui.QTabletEvent))(signal))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebEngineView_Update
func callbackQWebEngineView_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebEngineView) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_UpdateMicroFocus
func callbackQWebEngineView_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebEngineView) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebEngineView_WheelEvent
func callbackQWebEngineView_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*gui.QWheelEvent))(signal))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQWebEngineView_WindowIconChanged
func callbackQWebEngineView_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQWebEngineView_WindowTitleChanged
func callbackQWebEngineView_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackQWebEngineView_ChildEvent
func callbackQWebEngineView_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebEngineView_ConnectNotify
func callbackQWebEngineView_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineView_CustomEvent
func callbackQWebEngineView_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebEngineView_DeleteLater
func callbackQWebEngineView_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebEngineView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebEngineView_Destroyed
func callbackQWebEngineView_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebEngineView_DisconnectNotify
func callbackQWebEngineView_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebEngineViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebEngineView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebEngineView_EventFilter
func callbackQWebEngineView_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebEngineView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebEngineView_ObjectNameChanged
func callbackQWebEngineView_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebEngineView_TimerEvent
func callbackQWebEngineView_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebEngineViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebEngineView) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QtWebEngine struct {
	ptr unsafe.Pointer
}

type QtWebEngine_ITF interface {
	QtWebEngine_PTR() *QtWebEngine
}

func (ptr *QtWebEngine) QtWebEngine_PTR() *QtWebEngine {
	return ptr
}

func (ptr *QtWebEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtWebEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtWebEngine(ptr QtWebEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWebEngine_PTR().Pointer()
	}
	return nil
}

func NewQtWebEngineFromPointer(ptr unsafe.Pointer) (n *QtWebEngine) {
	n = new(QtWebEngine)
	n.SetPointer(ptr)
	return
}

func (ptr *QtWebEngine) DestroyQtWebEngine() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func QtWebEngine_Initialize() {
	C.QtWebEngine_QtWebEngine_Initialize()
}

func (ptr *QtWebEngine) Initialize() {
	C.QtWebEngine_QtWebEngine_Initialize()
}

type RenderViewObserverQt struct {
	ptr unsafe.Pointer
}

type RenderViewObserverQt_ITF interface {
	RenderViewObserverQt_PTR() *RenderViewObserverQt
}

func (ptr *RenderViewObserverQt) RenderViewObserverQt_PTR() *RenderViewObserverQt {
	return ptr
}

func (ptr *RenderViewObserverQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *RenderViewObserverQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromRenderViewObserverQt(ptr RenderViewObserverQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.RenderViewObserverQt_PTR().Pointer()
	}
	return nil
}

func NewRenderViewObserverQtFromPointer(ptr unsafe.Pointer) (n *RenderViewObserverQt) {
	n = new(RenderViewObserverQt)
	n.SetPointer(ptr)
	return
}

func (ptr *RenderViewObserverQt) DestroyRenderViewObserverQt() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type ServiceQt struct {
	ptr unsafe.Pointer
}

type ServiceQt_ITF interface {
	ServiceQt_PTR() *ServiceQt
}

func (ptr *ServiceQt) ServiceQt_PTR() *ServiceQt {
	return ptr
}

func (ptr *ServiceQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *ServiceQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromServiceQt(ptr ServiceQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ServiceQt_PTR().Pointer()
	}
	return nil
}

func NewServiceQtFromPointer(ptr unsafe.Pointer) (n *ServiceQt) {
	n = new(ServiceQt)
	n.SetPointer(ptr)
	return
}

func (ptr *ServiceQt) DestroyServiceQt() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type UserResourceController struct {
	ptr unsafe.Pointer
}

type UserResourceController_ITF interface {
	UserResourceController_PTR() *UserResourceController
}

func (ptr *UserResourceController) UserResourceController_PTR() *UserResourceController {
	return ptr
}

func (ptr *UserResourceController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *UserResourceController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromUserResourceController(ptr UserResourceController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.UserResourceController_PTR().Pointer()
	}
	return nil
}

func NewUserResourceControllerFromPointer(ptr unsafe.Pointer) (n *UserResourceController) {
	n = new(UserResourceController)
	n.SetPointer(ptr)
	return
}

func (ptr *UserResourceController) DestroyUserResourceController() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type WebEngineError struct {
	ptr unsafe.Pointer
}

type WebEngineError_ITF interface {
	WebEngineError_PTR() *WebEngineError
}

func (ptr *WebEngineError) WebEngineError_PTR() *WebEngineError {
	return ptr
}

func (ptr *WebEngineError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *WebEngineError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromWebEngineError(ptr WebEngineError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WebEngineError_PTR().Pointer()
	}
	return nil
}

func NewWebEngineErrorFromPointer(ptr unsafe.Pointer) (n *WebEngineError) {
	n = new(WebEngineError)
	n.SetPointer(ptr)
	return
}

func (ptr *WebEngineError) DestroyWebEngineError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type WebEngineLibraryInfo struct {
	ptr unsafe.Pointer
}

type WebEngineLibraryInfo_ITF interface {
	WebEngineLibraryInfo_PTR() *WebEngineLibraryInfo
}

func (ptr *WebEngineLibraryInfo) WebEngineLibraryInfo_PTR() *WebEngineLibraryInfo {
	return ptr
}

func (ptr *WebEngineLibraryInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *WebEngineLibraryInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromWebEngineLibraryInfo(ptr WebEngineLibraryInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WebEngineLibraryInfo_PTR().Pointer()
	}
	return nil
}

func NewWebEngineLibraryInfoFromPointer(ptr unsafe.Pointer) (n *WebEngineLibraryInfo) {
	n = new(WebEngineLibraryInfo)
	n.SetPointer(ptr)
	return
}

func (ptr *WebEngineLibraryInfo) DestroyWebEngineLibraryInfo() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type WebEventFactory struct {
	ptr unsafe.Pointer
}

type WebEventFactory_ITF interface {
	WebEventFactory_PTR() *WebEventFactory
}

func (ptr *WebEventFactory) WebEventFactory_PTR() *WebEventFactory {
	return ptr
}

func (ptr *WebEventFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *WebEventFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromWebEventFactory(ptr WebEventFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WebEventFactory_PTR().Pointer()
	}
	return nil
}

func NewWebEventFactoryFromPointer(ptr unsafe.Pointer) (n *WebEventFactory) {
	n = new(WebEventFactory)
	n.SetPointer(ptr)
	return
}

func (ptr *WebEventFactory) DestroyWebEventFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}
