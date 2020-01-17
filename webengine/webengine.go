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
	"github.com/therecipe/qt/printsupport"
	"github.com/therecipe/qt/webchannel"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtWebEngine_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWebEngine_PackedString) []byte {
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
			C.QQuickWebEngineProfile_ConnectCachePathChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cachePathChanged")))
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
			C.QQuickWebEngineProfile_ConnectDownloadPathChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "downloadPathChanged")))
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
			C.QQuickWebEngineProfile_ConnectHttpAcceptLanguageChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "httpAcceptLanguageChanged")))
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
			C.QQuickWebEngineProfile_ConnectHttpCacheMaximumSizeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "httpCacheMaximumSizeChanged")))
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
			C.QQuickWebEngineProfile_ConnectHttpCacheTypeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "httpCacheTypeChanged")))
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
			C.QQuickWebEngineProfile_ConnectHttpUserAgentChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "httpUserAgentChanged")))
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
			C.QQuickWebEngineProfile_ConnectOffTheRecordChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "offTheRecordChanged")))
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
			C.QQuickWebEngineProfile_ConnectPersistentCookiesPolicyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "persistentCookiesPolicyChanged")))
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
			C.QQuickWebEngineProfile_ConnectPersistentStoragePathChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "persistentStoragePathChanged")))
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
			C.QQuickWebEngineProfile_ConnectPresentNotification(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "presentNotification")))
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
			C.QQuickWebEngineProfile_ConnectSpellCheckEnabledChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "spellCheckEnabledChanged")))
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
			C.QQuickWebEngineProfile_ConnectSpellCheckLanguagesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "spellCheckLanguagesChanged")))
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
			C.QQuickWebEngineProfile_ConnectStorageNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "storageNameChanged")))
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
			C.QQuickWebEngineProfile_ConnectUseForGlobalCertificateVerificationChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "useForGlobalCertificateVerificationChanged")))
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QQuickWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQQuickWebEngineProfile_MetaObject
func callbackQQuickWebEngineProfile_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQQuickWebEngineProfileFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWebEngineProfile) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWebEngineProfile_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
			C.QQuickWebEngineScript_ConnectInjectionPointChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "injectionPointChanged")))
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
			C.QQuickWebEngineScript_ConnectNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nameChanged")))
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
			C.QQuickWebEngineScript_ConnectRunOnSubframesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "runOnSubframesChanged")))
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
			C.QQuickWebEngineScript_ConnectSourceCodeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sourceCodeChanged")))
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
			C.QQuickWebEngineScript_ConnectSourceUrlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sourceUrlChanged")))
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
			C.QQuickWebEngineScript_ConnectWorldIdChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "worldIdChanged")))
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QQuickWebEngineScript_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQQuickWebEngineScript_MetaObject
func callbackQQuickWebEngineScript_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQQuickWebEngineScriptFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWebEngineScript) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWebEngineScript_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

type QWebEngineClientCertificateSelection struct {
	ptr unsafe.Pointer
}

type QWebEngineClientCertificateSelection_ITF interface {
	QWebEngineClientCertificateSelection_PTR() *QWebEngineClientCertificateSelection
}

func (ptr *QWebEngineClientCertificateSelection) QWebEngineClientCertificateSelection_PTR() *QWebEngineClientCertificateSelection {
	return ptr
}

func (ptr *QWebEngineClientCertificateSelection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebEngineClientCertificateSelection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebEngineClientCertificateSelection(ptr QWebEngineClientCertificateSelection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineClientCertificateSelection_PTR().Pointer()
	}
	return nil
}

func NewQWebEngineClientCertificateSelectionFromPointer(ptr unsafe.Pointer) (n *QWebEngineClientCertificateSelection) {
	n = new(QWebEngineClientCertificateSelection)
	n.SetPointer(ptr)
	return
}
func (ptr *QWebEngineClientCertificateSelection) DestroyQWebEngineClientCertificateSelection() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQWebEngineClientCertificateSelection(vqw QWebEngineClientCertificateSelection_ITF) *QWebEngineClientCertificateSelection {
	tmpValue := NewQWebEngineClientCertificateSelectionFromPointer(C.QWebEngineClientCertificateSelection_NewQWebEngineClientCertificateSelection(PointerFromQWebEngineClientCertificateSelection(vqw)))
	qt.SetFinalizer(tmpValue, (*QWebEngineClientCertificateSelection).DestroyQWebEngineClientCertificateSelection)
	return tmpValue
}

func (ptr *QWebEngineClientCertificateSelection) Certificates() []*network.QSslCertificate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebEngine_PackedList) []*network.QSslCertificate {
			out := make([]*network.QSslCertificate, int(l.len))
			tmpList := NewQWebEngineClientCertificateSelectionFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__certificates_atList(i)
			}
			return out
		}(C.QWebEngineClientCertificateSelection_Certificates(ptr.Pointer()))
	}
	return make([]*network.QSslCertificate, 0)
}

func (ptr *QWebEngineClientCertificateSelection) Host() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineClientCertificateSelection_Host(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineClientCertificateSelection) Select(certificate network.QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateSelection_Select(ptr.Pointer(), network.PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QWebEngineClientCertificateSelection) SelectNone() {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateSelection_SelectNone(ptr.Pointer())
	}
}

func (ptr *QWebEngineClientCertificateSelection) __certificates_atList(i int) *network.QSslCertificate {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslCertificateFromPointer(C.QWebEngineClientCertificateSelection___certificates_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*network.QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineClientCertificateSelection) __certificates_setList(i network.QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineClientCertificateSelection___certificates_setList(ptr.Pointer(), network.PointerFromQSslCertificate(i))
	}
}

func (ptr *QWebEngineClientCertificateSelection) __certificates_newList() unsafe.Pointer {
	return C.QWebEngineClientCertificateSelection___certificates_newList(ptr.Pointer())
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*network.QSslCertificate).DestroyQSslCertificate)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
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
			C.QWebEngineCookieStore_ConnectCookieAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cookieAdded")))
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
			C.QWebEngineCookieStore_ConnectCookieRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cookieRemoved")))
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStoreDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineCookieStore_DestroyQWebEngineCookieStoreDefault(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineCookieStore_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQWebEngineCookieStore_MetaObject
func callbackQWebEngineCookieStore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineCookieStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineCookieStore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineCookieStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
	qt.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
	return tmpValue
}

func NewQWebEngineHttpRequest2(other QWebEngineHttpRequest_ITF) *QWebEngineHttpRequest {
	tmpValue := NewQWebEngineHttpRequestFromPointer(C.QWebEngineHttpRequest_NewQWebEngineHttpRequest2(PointerFromQWebEngineHttpRequest(other)))
	qt.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	qt.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
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
	qt.SetFinalizer(tmpValue, (*QWebEngineHttpRequest).DestroyQWebEngineHttpRequest)
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineHttpRequest) DestroyQWebEngineHttpRequest() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineHttpRequest_DestroyQWebEngineHttpRequest(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineHttpRequest) __headers_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineHttpRequest___headers_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
			C.QWebEngineNotification_ConnectClosed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "closed")))
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
		qt.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineNotification_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQWebEngineNotification_MetaObject
func callbackQWebEngineNotification_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineNotificationFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineNotification) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineNotification_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQWebEnginePage(parent core.QObject_ITF) *QWebEnginePage {
	tmpValue := NewQWebEnginePageFromPointer(C.QWebEnginePage_NewQWebEnginePage(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQWebEnginePage2(profile QWebEngineProfile_ITF, parent core.QObject_ITF) *QWebEnginePage {
	tmpValue := NewQWebEnginePageFromPointer(C.QWebEnginePage_NewQWebEnginePage2(PointerFromQWebEngineProfile(profile), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEnginePage_AcceptNavigationRequest
func callbackQWebEnginePage_AcceptNavigationRequest(ptr unsafe.Pointer, url unsafe.Pointer, ty C.longlong, isMainFrame C.char) C.char {
	if signal := qt.GetSignal(ptr, "acceptNavigationRequest"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl, QWebEnginePage__NavigationType, bool) bool)(signal))(core.NewQUrlFromPointer(url), QWebEnginePage__NavigationType(ty), int8(isMainFrame) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).AcceptNavigationRequestDefault(core.NewQUrlFromPointer(url), QWebEnginePage__NavigationType(ty), int8(isMainFrame) != 0))))
}

func (ptr *QWebEnginePage) ConnectAcceptNavigationRequest(f func(url *core.QUrl, ty QWebEnginePage__NavigationType, isMainFrame bool) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "acceptNavigationRequest"); signal != nil {
			f := func(url *core.QUrl, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
				(*(*func(*core.QUrl, QWebEnginePage__NavigationType, bool) bool)(signal))(url, ty, isMainFrame)
				return f(url, ty, isMainFrame)
			}
			qt.ConnectSignal(ptr.Pointer(), "acceptNavigationRequest", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "acceptNavigationRequest", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectAcceptNavigationRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "acceptNavigationRequest")
	}
}

func (ptr *QWebEnginePage) AcceptNavigationRequest(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_AcceptNavigationRequest(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty), C.char(int8(qt.GoBoolToInt(isMainFrame))))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) AcceptNavigationRequestDefault(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_AcceptNavigationRequestDefault(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty), C.char(int8(qt.GoBoolToInt(isMainFrame))))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) Action(action QWebEnginePage__WebAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QWebEnginePage_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_AudioMutedChanged
func callbackQWebEnginePage_AudioMutedChanged(ptr unsafe.Pointer, muted C.char) {
	if signal := qt.GetSignal(ptr, "audioMutedChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(muted) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectAudioMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "audioMutedChanged") {
			C.QWebEnginePage_ConnectAudioMutedChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "audioMutedChanged")))
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

//export callbackQWebEnginePage_AuthenticationRequired
func callbackQWebEnginePage_AuthenticationRequired(ptr unsafe.Pointer, requestUrl unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "authenticationRequired"); signal != nil {
		(*(*func(*core.QUrl, *network.QAuthenticator))(signal))(core.NewQUrlFromPointer(requestUrl), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebEnginePage) ConnectAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "authenticationRequired") {
			C.QWebEnginePage_ConnectAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "authenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "authenticationRequired"); signal != nil {
			f := func(requestUrl *core.QUrl, authenticator *network.QAuthenticator) {
				(*(*func(*core.QUrl, *network.QAuthenticator))(signal))(requestUrl, authenticator)
				f(requestUrl, authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "authenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "authenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "authenticationRequired")
	}
}

func (ptr *QWebEnginePage) AuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_AuthenticationRequired(ptr.Pointer(), core.PointerFromQUrl(requestUrl), network.PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QWebEnginePage) BackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QWebEnginePage_BackgroundColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_CertificateError
func callbackQWebEnginePage_CertificateError(ptr unsafe.Pointer, certificateError unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "certificateError"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QWebEngineCertificateError) bool)(signal))(NewQWebEngineCertificateErrorFromPointer(certificateError)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).CertificateErrorDefault(NewQWebEngineCertificateErrorFromPointer(certificateError)))))
}

func (ptr *QWebEnginePage) ConnectCertificateError(f func(certificateError *QWebEngineCertificateError) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "certificateError"); signal != nil {
			f := func(certificateError *QWebEngineCertificateError) bool {
				(*(*func(*QWebEngineCertificateError) bool)(signal))(certificateError)
				return f(certificateError)
			}
			qt.ConnectSignal(ptr.Pointer(), "certificateError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "certificateError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectCertificateError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "certificateError")
	}
}

func (ptr *QWebEnginePage) CertificateError(certificateError QWebEngineCertificateError_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_CertificateError(ptr.Pointer(), PointerFromQWebEngineCertificateError(certificateError))) != 0
	}
	return false
}

func (ptr *QWebEnginePage) CertificateErrorDefault(certificateError QWebEngineCertificateError_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_CertificateErrorDefault(ptr.Pointer(), PointerFromQWebEngineCertificateError(certificateError))) != 0
	}
	return false
}

//export callbackQWebEnginePage_ChooseFiles
func callbackQWebEnginePage_ChooseFiles(ptr unsafe.Pointer, mode C.longlong, oldFiles C.struct_QtWebEngine_PackedString, acceptedMimeTypes C.struct_QtWebEngine_PackedString) C.struct_QtWebEngine_PackedString {
	if signal := qt.GetSignal(ptr, "chooseFiles"); signal != nil {
		tempVal := (*(*func(QWebEnginePage__FileSelectionMode, []string, []string) []string)(signal))(QWebEnginePage__FileSelectionMode(mode), unpackStringList(cGoUnpackString(oldFiles)), unpackStringList(cGoUnpackString(acceptedMimeTypes)))
		return C.struct_QtWebEngine_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewQWebEnginePageFromPointer(ptr).ChooseFilesDefault(QWebEnginePage__FileSelectionMode(mode), unpackStringList(cGoUnpackString(oldFiles)), unpackStringList(cGoUnpackString(acceptedMimeTypes)))
	return C.struct_QtWebEngine_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QWebEnginePage) ConnectChooseFiles(f func(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "chooseFiles"); signal != nil {
			f := func(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
				(*(*func(QWebEnginePage__FileSelectionMode, []string, []string) []string)(signal))(mode, oldFiles, acceptedMimeTypes)
				return f(mode, oldFiles, acceptedMimeTypes)
			}
			qt.ConnectSignal(ptr.Pointer(), "chooseFiles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "chooseFiles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectChooseFiles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "chooseFiles")
	}
}

func (ptr *QWebEnginePage) ChooseFiles(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
	if ptr.Pointer() != nil {
		oldFilesC := C.CString(strings.Join(oldFiles, "¡¦!"))
		defer C.free(unsafe.Pointer(oldFilesC))
		acceptedMimeTypesC := C.CString(strings.Join(acceptedMimeTypes, "¡¦!"))
		defer C.free(unsafe.Pointer(acceptedMimeTypesC))
		return unpackStringList(cGoUnpackString(C.QWebEnginePage_ChooseFiles(ptr.Pointer(), C.longlong(mode), C.struct_QtWebEngine_PackedString{data: oldFilesC, len: C.longlong(len(strings.Join(oldFiles, "¡¦!")))}, C.struct_QtWebEngine_PackedString{data: acceptedMimeTypesC, len: C.longlong(len(strings.Join(acceptedMimeTypes, "¡¦!")))})))
	}
	return make([]string, 0)
}

func (ptr *QWebEnginePage) ChooseFilesDefault(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {
	if ptr.Pointer() != nil {
		oldFilesC := C.CString(strings.Join(oldFiles, "¡¦!"))
		defer C.free(unsafe.Pointer(oldFilesC))
		acceptedMimeTypesC := C.CString(strings.Join(acceptedMimeTypes, "¡¦!"))
		defer C.free(unsafe.Pointer(acceptedMimeTypesC))
		return unpackStringList(cGoUnpackString(C.QWebEnginePage_ChooseFilesDefault(ptr.Pointer(), C.longlong(mode), C.struct_QtWebEngine_PackedString{data: oldFilesC, len: C.longlong(len(strings.Join(oldFiles, "¡¦!")))}, C.struct_QtWebEngine_PackedString{data: acceptedMimeTypesC, len: C.longlong(len(strings.Join(acceptedMimeTypes, "¡¦!")))})))
	}
	return make([]string, 0)
}

func (ptr *QWebEnginePage) ContentsSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QWebEnginePage_ContentsSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
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
			C.QWebEnginePage_ConnectContentsSizeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "contentsSizeChanged")))
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

func (ptr *QWebEnginePage) ContextMenuData() *QWebEngineContextMenuData {
	if ptr.Pointer() != nil {
		return NewQWebEngineContextMenuDataFromPointer(C.QWebEnginePage_ContextMenuData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEnginePage) CreateStandardContextMenu() *widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQMenuFromPointer(C.QWebEnginePage_CreateStandardContextMenu(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_CreateWindow
func callbackQWebEnginePage_CreateWindow(ptr unsafe.Pointer, ty C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createWindow"); signal != nil {
		return PointerFromQWebEnginePage((*(*func(QWebEnginePage__WebWindowType) *QWebEnginePage)(signal))(QWebEnginePage__WebWindowType(ty)))
	}

	return PointerFromQWebEnginePage(NewQWebEnginePageFromPointer(ptr).CreateWindowDefault(QWebEnginePage__WebWindowType(ty)))
}

func (ptr *QWebEnginePage) ConnectCreateWindow(f func(ty QWebEnginePage__WebWindowType) *QWebEnginePage) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createWindow"); signal != nil {
			f := func(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
				(*(*func(QWebEnginePage__WebWindowType) *QWebEnginePage)(signal))(ty)
				return f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "createWindow", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createWindow", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectCreateWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createWindow")
	}
}

func (ptr *QWebEnginePage) CreateWindow(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEnginePageFromPointer(C.QWebEnginePage_CreateWindow(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) CreateWindowDefault(ty QWebEnginePage__WebWindowType) *QWebEnginePage {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEnginePageFromPointer(C.QWebEnginePage_CreateWindowDefault(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) DevToolsPage() *QWebEnginePage {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEnginePageFromPointer(C.QWebEnginePage_DevToolsPage(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) Download(url core.QUrl_ITF, filename string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		C.QWebEnginePage_Download(ptr.Pointer(), core.PointerFromQUrl(url), C.struct_QtWebEngine_PackedString{data: filenameC, len: C.longlong(len(filename))})
	}
}

//export callbackQWebEnginePage_Event
func callbackQWebEnginePage_Event(ptr unsafe.Pointer, vqe unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(vqe)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).EventDefault(core.NewQEventFromPointer(vqe)))))
}

func (ptr *QWebEnginePage) EventDefault(vqe core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_EventDefault(ptr.Pointer(), core.PointerFromQEvent(vqe))) != 0
	}
	return false
}

//export callbackQWebEnginePage_FeaturePermissionRequestCanceled
func callbackQWebEnginePage_FeaturePermissionRequestCanceled(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, feature C.longlong) {
	if signal := qt.GetSignal(ptr, "featurePermissionRequestCanceled"); signal != nil {
		(*(*func(*core.QUrl, QWebEnginePage__Feature))(signal))(core.NewQUrlFromPointer(securityOrigin), QWebEnginePage__Feature(feature))
	}

}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequestCanceled(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "featurePermissionRequestCanceled") {
			C.QWebEnginePage_ConnectFeaturePermissionRequestCanceled(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "featurePermissionRequestCanceled")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "featurePermissionRequestCanceled"); signal != nil {
			f := func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature) {
				(*(*func(*core.QUrl, QWebEnginePage__Feature))(signal))(securityOrigin, feature)
				f(securityOrigin, feature)
			}
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequestCanceled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequestCanceled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequestCanceled() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "featurePermissionRequestCanceled")
	}
}

func (ptr *QWebEnginePage) FeaturePermissionRequestCanceled(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_FeaturePermissionRequestCanceled(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.longlong(feature))
	}
}

//export callbackQWebEnginePage_FeaturePermissionRequested
func callbackQWebEnginePage_FeaturePermissionRequested(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, feature C.longlong) {
	if signal := qt.GetSignal(ptr, "featurePermissionRequested"); signal != nil {
		(*(*func(*core.QUrl, QWebEnginePage__Feature))(signal))(core.NewQUrlFromPointer(securityOrigin), QWebEnginePage__Feature(feature))
	}

}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequested(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "featurePermissionRequested") {
			C.QWebEnginePage_ConnectFeaturePermissionRequested(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "featurePermissionRequested")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "featurePermissionRequested"); signal != nil {
			f := func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature) {
				(*(*func(*core.QUrl, QWebEnginePage__Feature))(signal))(securityOrigin, feature)
				f(securityOrigin, feature)
			}
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequested", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequested", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectFeaturePermissionRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "featurePermissionRequested")
	}
}

func (ptr *QWebEnginePage) FeaturePermissionRequested(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_FeaturePermissionRequested(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.longlong(feature))
	}
}

//export callbackQWebEnginePage_GeometryChangeRequested
func callbackQWebEnginePage_GeometryChangeRequested(ptr unsafe.Pointer, geom unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChangeRequested"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(geom))
	}

}

func (ptr *QWebEnginePage) ConnectGeometryChangeRequested(f func(geom *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "geometryChangeRequested") {
			C.QWebEnginePage_ConnectGeometryChangeRequested(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "geometryChangeRequested")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "geometryChangeRequested"); signal != nil {
			f := func(geom *core.QRect) {
				(*(*func(*core.QRect))(signal))(geom)
				f(geom)
			}
			qt.ConnectSignal(ptr.Pointer(), "geometryChangeRequested", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometryChangeRequested", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectGeometryChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectGeometryChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "geometryChangeRequested")
	}
}

func (ptr *QWebEnginePage) GeometryChangeRequested(geom core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_GeometryChangeRequested(ptr.Pointer(), core.PointerFromQRect(geom))
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
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
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
			C.QWebEnginePage_ConnectIconChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "iconChanged")))
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
			C.QWebEnginePage_ConnectIconUrlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "iconUrlChanged")))
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

func (ptr *QWebEnginePage) InspectedPage() *QWebEnginePage {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEnginePageFromPointer(C.QWebEnginePage_InspectedPage(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) IsAudioMuted() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEnginePage_IsAudioMuted(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebEnginePage_JavaScriptAlert
func callbackQWebEnginePage_JavaScriptAlert(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, msg C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "javaScriptAlert"); signal != nil {
		(*(*func(*core.QUrl, string))(signal))(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg))
	} else {
		NewQWebEnginePageFromPointer(ptr).JavaScriptAlertDefault(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg))
	}
}

func (ptr *QWebEnginePage) ConnectJavaScriptAlert(f func(securityOrigin *core.QUrl, msg string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptAlert"); signal != nil {
			f := func(securityOrigin *core.QUrl, msg string) {
				(*(*func(*core.QUrl, string))(signal))(securityOrigin, msg)
				f(securityOrigin, msg)
			}
			qt.ConnectSignal(ptr.Pointer(), "javaScriptAlert", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptAlert", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptAlert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptAlert")
	}
}

func (ptr *QWebEnginePage) JavaScriptAlert(securityOrigin core.QUrl_ITF, msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QWebEnginePage_JavaScriptAlert(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.struct_QtWebEngine_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

func (ptr *QWebEnginePage) JavaScriptAlertDefault(securityOrigin core.QUrl_ITF, msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QWebEnginePage_JavaScriptAlertDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.struct_QtWebEngine_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackQWebEnginePage_JavaScriptConfirm
func callbackQWebEnginePage_JavaScriptConfirm(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, msg C.struct_QtWebEngine_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "javaScriptConfirm"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl, string) bool)(signal))(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).JavaScriptConfirmDefault(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg)))))
}

func (ptr *QWebEnginePage) ConnectJavaScriptConfirm(f func(securityOrigin *core.QUrl, msg string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptConfirm"); signal != nil {
			f := func(securityOrigin *core.QUrl, msg string) bool {
				(*(*func(*core.QUrl, string) bool)(signal))(securityOrigin, msg)
				return f(securityOrigin, msg)
			}
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConfirm", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConfirm", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConfirm() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptConfirm")
	}
}

func (ptr *QWebEnginePage) JavaScriptConfirm(securityOrigin core.QUrl_ITF, msg string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		return int8(C.QWebEnginePage_JavaScriptConfirm(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.struct_QtWebEngine_PackedString{data: msgC, len: C.longlong(len(msg))})) != 0
	}
	return false
}

func (ptr *QWebEnginePage) JavaScriptConfirmDefault(securityOrigin core.QUrl_ITF, msg string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		return int8(C.QWebEnginePage_JavaScriptConfirmDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.struct_QtWebEngine_PackedString{data: msgC, len: C.longlong(len(msg))})) != 0
	}
	return false
}

//export callbackQWebEnginePage_JavaScriptConsoleMessage
func callbackQWebEnginePage_JavaScriptConsoleMessage(ptr unsafe.Pointer, level C.longlong, message C.struct_QtWebEngine_PackedString, lineNumber C.int, sourceID C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "javaScriptConsoleMessage"); signal != nil {
		(*(*func(QWebEnginePage__JavaScriptConsoleMessageLevel, string, int, string))(signal))(QWebEnginePage__JavaScriptConsoleMessageLevel(level), cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	} else {
		NewQWebEnginePageFromPointer(ptr).JavaScriptConsoleMessageDefault(QWebEnginePage__JavaScriptConsoleMessageLevel(level), cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	}
}

func (ptr *QWebEnginePage) ConnectJavaScriptConsoleMessage(f func(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptConsoleMessage"); signal != nil {
			f := func(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
				(*(*func(QWebEnginePage__JavaScriptConsoleMessageLevel, string, int, string))(signal))(level, message, lineNumber, sourceID)
				f(level, message, lineNumber, sourceID)
			}
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConsoleMessage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConsoleMessage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConsoleMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptConsoleMessage")
	}
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessage(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		var sourceIDC *C.char
		if sourceID != "" {
			sourceIDC = C.CString(sourceID)
			defer C.free(unsafe.Pointer(sourceIDC))
		}
		C.QWebEnginePage_JavaScriptConsoleMessage(ptr.Pointer(), C.longlong(level), C.struct_QtWebEngine_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(lineNumber)), C.struct_QtWebEngine_PackedString{data: sourceIDC, len: C.longlong(len(sourceID))})
	}
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessageDefault(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		var sourceIDC *C.char
		if sourceID != "" {
			sourceIDC = C.CString(sourceID)
			defer C.free(unsafe.Pointer(sourceIDC))
		}
		C.QWebEnginePage_JavaScriptConsoleMessageDefault(ptr.Pointer(), C.longlong(level), C.struct_QtWebEngine_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(lineNumber)), C.struct_QtWebEngine_PackedString{data: sourceIDC, len: C.longlong(len(sourceID))})
	}
}

//export callbackQWebEnginePage_JavaScriptPrompt
func callbackQWebEnginePage_JavaScriptPrompt(ptr unsafe.Pointer, securityOrigin unsafe.Pointer, msg C.struct_QtWebEngine_PackedString, defaultValue C.struct_QtWebEngine_PackedString, result C.struct_QtWebEngine_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "javaScriptPrompt"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl, string, string, string) bool)(signal))(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg), cGoUnpackString(defaultValue), cGoUnpackString(result)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEnginePageFromPointer(ptr).JavaScriptPromptDefault(core.NewQUrlFromPointer(securityOrigin), cGoUnpackString(msg), cGoUnpackString(defaultValue), cGoUnpackString(result)))))
}

func (ptr *QWebEnginePage) ConnectJavaScriptPrompt(f func(securityOrigin *core.QUrl, msg string, defaultValue string, result string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptPrompt"); signal != nil {
			f := func(securityOrigin *core.QUrl, msg string, defaultValue string, result string) bool {
				(*(*func(*core.QUrl, string, string, string) bool)(signal))(securityOrigin, msg, defaultValue, result)
				return f(securityOrigin, msg, defaultValue, result)
			}
			qt.ConnectSignal(ptr.Pointer(), "javaScriptPrompt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptPrompt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectJavaScriptPrompt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptPrompt")
	}
}

func (ptr *QWebEnginePage) JavaScriptPrompt(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		var defaultValueC *C.char
		if defaultValue != "" {
			defaultValueC = C.CString(defaultValue)
			defer C.free(unsafe.Pointer(defaultValueC))
		}
		var resultC *C.char
		if result != "" {
			resultC = C.CString(result)
			defer C.free(unsafe.Pointer(resultC))
		}
		return int8(C.QWebEnginePage_JavaScriptPrompt(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.struct_QtWebEngine_PackedString{data: msgC, len: C.longlong(len(msg))}, C.struct_QtWebEngine_PackedString{data: defaultValueC, len: C.longlong(len(defaultValue))}, C.struct_QtWebEngine_PackedString{data: resultC, len: C.longlong(len(result))})) != 0
	}
	return false
}

func (ptr *QWebEnginePage) JavaScriptPromptDefault(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		var defaultValueC *C.char
		if defaultValue != "" {
			defaultValueC = C.CString(defaultValue)
			defer C.free(unsafe.Pointer(defaultValueC))
		}
		var resultC *C.char
		if result != "" {
			resultC = C.CString(result)
			defer C.free(unsafe.Pointer(resultC))
		}
		return int8(C.QWebEnginePage_JavaScriptPromptDefault(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.struct_QtWebEngine_PackedString{data: msgC, len: C.longlong(len(msg))}, C.struct_QtWebEngine_PackedString{data: defaultValueC, len: C.longlong(len(defaultValue))}, C.struct_QtWebEngine_PackedString{data: resultC, len: C.longlong(len(result))})) != 0
	}
	return false
}

//export callbackQWebEnginePage_LinkHovered
func callbackQWebEnginePage_LinkHovered(ptr unsafe.Pointer, url C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "linkHovered"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(url))
	}

}

func (ptr *QWebEnginePage) ConnectLinkHovered(f func(url string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkHovered") {
			C.QWebEnginePage_ConnectLinkHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "linkHovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkHovered"); signal != nil {
			f := func(url string) {
				(*(*func(string))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "linkHovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkHovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkHovered")
	}
}

func (ptr *QWebEnginePage) LinkHovered(url string) {
	if ptr.Pointer() != nil {
		var urlC *C.char
		if url != "" {
			urlC = C.CString(url)
			defer C.free(unsafe.Pointer(urlC))
		}
		C.QWebEnginePage_LinkHovered(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: urlC, len: C.longlong(len(url))})
	}
}

func (ptr *QWebEnginePage) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) Load2(request QWebEngineHttpRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_Load2(ptr.Pointer(), PointerFromQWebEngineHttpRequest(request))
	}
}

//export callbackQWebEnginePage_LoadFinished
func callbackQWebEnginePage_LoadFinished(ptr unsafe.Pointer, ok C.char) {
	if signal := qt.GetSignal(ptr, "loadFinished"); signal != nil {
		(*(*func(bool))(signal))(int8(ok) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadFinished") {
			C.QWebEnginePage_ConnectLoadFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "loadFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadFinished"); signal != nil {
			f := func(ok bool) {
				(*(*func(bool))(signal))(ok)
				f(ok)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadFinished")
	}
}

func (ptr *QWebEnginePage) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQWebEnginePage_LoadProgress
func callbackQWebEnginePage_LoadProgress(ptr unsafe.Pointer, progress C.int) {
	if signal := qt.GetSignal(ptr, "loadProgress"); signal != nil {
		(*(*func(int))(signal))(int(int32(progress)))
	}

}

func (ptr *QWebEnginePage) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadProgress") {
			C.QWebEnginePage_ConnectLoadProgress(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "loadProgress")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadProgress"); signal != nil {
			f := func(progress int) {
				(*(*func(int))(signal))(progress)
				f(progress)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadProgress")
	}
}

func (ptr *QWebEnginePage) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQWebEnginePage_LoadStarted
func callbackQWebEnginePage_LoadStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadStarted"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEnginePage) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadStarted") {
			C.QWebEnginePage_ConnectLoadStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "loadStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadStarted")
	}
}

func (ptr *QWebEnginePage) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebEnginePage_MetaObject
func callbackQWebEnginePage_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEnginePageFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEnginePage) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEnginePage_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebEnginePage_PdfPrintingFinished
func callbackQWebEnginePage_PdfPrintingFinished(ptr unsafe.Pointer, filePath C.struct_QtWebEngine_PackedString, success C.char) {
	if signal := qt.GetSignal(ptr, "pdfPrintingFinished"); signal != nil {
		(*(*func(string, bool))(signal))(cGoUnpackString(filePath), int8(success) != 0)
	}

}

func (ptr *QWebEnginePage) ConnectPdfPrintingFinished(f func(filePath string, success bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pdfPrintingFinished") {
			C.QWebEnginePage_ConnectPdfPrintingFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pdfPrintingFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pdfPrintingFinished"); signal != nil {
			f := func(filePath string, success bool) {
				(*(*func(string, bool))(signal))(filePath, success)
				f(filePath, success)
			}
			qt.ConnectSignal(ptr.Pointer(), "pdfPrintingFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pdfPrintingFinished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectPdfPrintingFinished() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectPdfPrintingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pdfPrintingFinished")
	}
}

func (ptr *QWebEnginePage) PdfPrintingFinished(filePath string, success bool) {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		C.QWebEnginePage_PdfPrintingFinished(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: filePathC, len: C.longlong(len(filePath))}, C.char(int8(qt.GoBoolToInt(success))))
	}
}

func (ptr *QWebEnginePage) Print(printer printsupport.QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_Print(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

//export callbackQWebEnginePage_PrintRequested
func callbackQWebEnginePage_PrintRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "printRequested"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEnginePage) ConnectPrintRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "printRequested") {
			C.QWebEnginePage_ConnectPrintRequested(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "printRequested")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "printRequested"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "printRequested", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "printRequested", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectPrintRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectPrintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "printRequested")
	}
}

func (ptr *QWebEnginePage) PrintRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_PrintRequested(ptr.Pointer())
	}
}

func (ptr *QWebEnginePage) PrintToPdf(filePath string, layout gui.QPageLayout_ITF) {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		C.QWebEnginePage_PrintToPdf(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: filePathC, len: C.longlong(len(filePath))}, gui.PointerFromQPageLayout(layout))
	}
}

func (ptr *QWebEnginePage) Profile() *QWebEngineProfile {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEngineProfileFromPointer(C.QWebEnginePage_Profile(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_ProxyAuthenticationRequired
func callbackQWebEnginePage_ProxyAuthenticationRequired(ptr unsafe.Pointer, requestUrl unsafe.Pointer, authenticator unsafe.Pointer, proxyHost C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "proxyAuthenticationRequired"); signal != nil {
		(*(*func(*core.QUrl, *network.QAuthenticator, string))(signal))(core.NewQUrlFromPointer(requestUrl), network.NewQAuthenticatorFromPointer(authenticator), cGoUnpackString(proxyHost))
	}

}

func (ptr *QWebEnginePage) ConnectProxyAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator, proxyHost string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "proxyAuthenticationRequired") {
			C.QWebEnginePage_ConnectProxyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "proxyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "proxyAuthenticationRequired"); signal != nil {
			f := func(requestUrl *core.QUrl, authenticator *network.QAuthenticator, proxyHost string) {
				(*(*func(*core.QUrl, *network.QAuthenticator, string))(signal))(requestUrl, authenticator, proxyHost)
				f(requestUrl, authenticator, proxyHost)
			}
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "proxyAuthenticationRequired")
	}
}

func (ptr *QWebEnginePage) ProxyAuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF, proxyHost string) {
	if ptr.Pointer() != nil {
		var proxyHostC *C.char
		if proxyHost != "" {
			proxyHostC = C.CString(proxyHost)
			defer C.free(unsafe.Pointer(proxyHostC))
		}
		C.QWebEnginePage_ProxyAuthenticationRequired(ptr.Pointer(), core.PointerFromQUrl(requestUrl), network.PointerFromQAuthenticator(authenticator), C.struct_QtWebEngine_PackedString{data: proxyHostC, len: C.longlong(len(proxyHost))})
	}
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
			C.QWebEnginePage_ConnectRecentlyAudibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "recentlyAudibleChanged")))
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

func (ptr *QWebEnginePage) ReplaceMisspelledWord(replacement string) {
	if ptr.Pointer() != nil {
		var replacementC *C.char
		if replacement != "" {
			replacementC = C.CString(replacement)
			defer C.free(unsafe.Pointer(replacementC))
		}
		C.QWebEnginePage_ReplaceMisspelledWord(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: replacementC, len: C.longlong(len(replacement))})
	}
}

func (ptr *QWebEnginePage) RequestedUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEnginePage_RequestedUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) RunJavaScript(scriptSource string) {
	if ptr.Pointer() != nil {
		var scriptSourceC *C.char
		if scriptSource != "" {
			scriptSourceC = C.CString(scriptSource)
			defer C.free(unsafe.Pointer(scriptSourceC))
		}
		C.QWebEnginePage_RunJavaScript(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: scriptSourceC, len: C.longlong(len(scriptSource))})
	}
}

func (ptr *QWebEnginePage) RunJavaScript2(scriptSource string, worldId uint) {
	if ptr.Pointer() != nil {
		var scriptSourceC *C.char
		if scriptSource != "" {
			scriptSourceC = C.CString(scriptSource)
			defer C.free(unsafe.Pointer(scriptSourceC))
		}
		C.QWebEnginePage_RunJavaScript2(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: scriptSourceC, len: C.longlong(len(scriptSource))}, C.uint(uint32(worldId)))
	}
}

func (ptr *QWebEnginePage) ScrollPosition() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QWebEnginePage_ScrollPosition(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
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
			C.QWebEnginePage_ConnectScrollPositionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "scrollPositionChanged")))
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

//export callbackQWebEnginePage_SelectClientCertificate
func callbackQWebEnginePage_SelectClientCertificate(ptr unsafe.Pointer, clientCertSelection unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectClientCertificate"); signal != nil {
		(*(*func(*QWebEngineClientCertificateSelection))(signal))(func() *QWebEngineClientCertificateSelection {
			tmpValue := NewQWebEngineClientCertificateSelectionFromPointer(clientCertSelection)
			qt.SetFinalizer(tmpValue, (*QWebEngineClientCertificateSelection).DestroyQWebEngineClientCertificateSelection)
			return tmpValue
		}())
	}

}

func (ptr *QWebEnginePage) ConnectSelectClientCertificate(f func(clientCertSelection *QWebEngineClientCertificateSelection)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectClientCertificate") {
			C.QWebEnginePage_ConnectSelectClientCertificate(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectClientCertificate")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectClientCertificate"); signal != nil {
			f := func(clientCertSelection *QWebEngineClientCertificateSelection) {
				(*(*func(*QWebEngineClientCertificateSelection))(signal))(clientCertSelection)
				f(clientCertSelection)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectClientCertificate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectClientCertificate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectSelectClientCertificate() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectSelectClientCertificate(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectClientCertificate")
	}
}

func (ptr *QWebEnginePage) SelectClientCertificate(clientCertSelection QWebEngineClientCertificateSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SelectClientCertificate(ptr.Pointer(), PointerFromQWebEngineClientCertificateSelection(clientCertSelection))
	}
}

func (ptr *QWebEnginePage) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEnginePage_SelectedText(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebEnginePage_SelectionChanged
func callbackQWebEnginePage_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEnginePage) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QWebEnginePage_ConnectSelectionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QWebEnginePage) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SelectionChanged(ptr.Pointer())
	}
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
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		C.QWebEnginePage_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.struct_QtWebEngine_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEnginePage) SetDevToolsPage(page QWebEnginePage_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetDevToolsPage(ptr.Pointer(), PointerFromQWebEnginePage(page))
	}
}

func (ptr *QWebEnginePage) SetFeaturePermission(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature, policy QWebEnginePage__PermissionPolicy) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetFeaturePermission(ptr.Pointer(), core.PointerFromQUrl(securityOrigin), C.longlong(feature), C.longlong(policy))
	}
}

func (ptr *QWebEnginePage) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.QWebEnginePage_SetHtml(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: htmlC, len: C.longlong(len(html))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEnginePage) SetInspectedPage(page QWebEnginePage_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetInspectedPage(ptr.Pointer(), PointerFromQWebEnginePage(page))
	}
}

func (ptr *QWebEnginePage) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) SetUrlRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetUrlRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

func (ptr *QWebEnginePage) SetView(view widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetView(ptr.Pointer(), widgets.PointerFromQWidget(view))
	}
}

func (ptr *QWebEnginePage) SetWebChannel(vqw webchannel.QWebChannel_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetWebChannel(ptr.Pointer(), webchannel.PointerFromQWebChannel(vqw))
	}
}

func (ptr *QWebEnginePage) SetWebChannel2(vqw webchannel.QWebChannel_ITF, worldId uint) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_SetWebChannel2(ptr.Pointer(), webchannel.PointerFromQWebChannel(vqw), C.uint(uint32(worldId)))
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

//export callbackQWebEnginePage_TitleChanged
func callbackQWebEnginePage_TitleChanged(ptr unsafe.Pointer, title C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

func (ptr *QWebEnginePage) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.QWebEnginePage_ConnectTitleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "titleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			f := func(title string) {
				(*(*func(string))(signal))(title)
				f(title)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *QWebEnginePage) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QWebEnginePage_TitleChanged(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func QWebEnginePage_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QWebEnginePage_QWebEnginePage_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebEnginePage) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QWebEnginePage_QWebEnginePage_Tr(sC, cC, C.int(int32(n))))
}

//export callbackQWebEnginePage_TriggerAction
func callbackQWebEnginePage_TriggerAction(ptr unsafe.Pointer, action C.longlong, checked C.char) {
	if signal := qt.GetSignal(ptr, "triggerAction"); signal != nil {
		(*(*func(QWebEnginePage__WebAction, bool))(signal))(QWebEnginePage__WebAction(action), int8(checked) != 0)
	} else {
		NewQWebEnginePageFromPointer(ptr).TriggerActionDefault(QWebEnginePage__WebAction(action), int8(checked) != 0)
	}
}

func (ptr *QWebEnginePage) ConnectTriggerAction(f func(action QWebEnginePage__WebAction, checked bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "triggerAction"); signal != nil {
			f := func(action QWebEnginePage__WebAction, checked bool) {
				(*(*func(QWebEnginePage__WebAction, bool))(signal))(action, checked)
				f(action, checked)
			}
			qt.ConnectSignal(ptr.Pointer(), "triggerAction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggerAction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectTriggerAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "triggerAction")
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
		tmpValue := core.NewQUrlFromPointer(C.QWebEnginePage_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_UrlChanged
func callbackQWebEnginePage_UrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "urlChanged"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebEnginePage) ConnectUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "urlChanged") {
			C.QWebEnginePage_ConnectUrlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "urlChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "urlChanged"); signal != nil {
			f := func(url *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "urlChanged")
	}
}

func (ptr *QWebEnginePage) UrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEnginePage) View() *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QWebEnginePage_View(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEnginePage) WebChannel() *webchannel.QWebChannel {
	if ptr.Pointer() != nil {
		tmpValue := webchannel.NewQWebChannelFromPointer(C.QWebEnginePage_WebChannel(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEnginePage_WindowCloseRequested
func callbackQWebEnginePage_WindowCloseRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowCloseRequested"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEnginePage) ConnectWindowCloseRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "windowCloseRequested") {
			C.QWebEnginePage_ConnectWindowCloseRequested(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "windowCloseRequested")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "windowCloseRequested"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "windowCloseRequested", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "windowCloseRequested", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectWindowCloseRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_DisconnectWindowCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "windowCloseRequested")
	}
}

func (ptr *QWebEnginePage) WindowCloseRequested() {
	if ptr.Pointer() != nil {
		C.QWebEnginePage_WindowCloseRequested(ptr.Pointer())
	}
}

func (ptr *QWebEnginePage) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebEnginePage_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebEnginePage_DestroyQWebEnginePage
func callbackQWebEnginePage_DestroyQWebEnginePage(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebEnginePage"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEnginePageFromPointer(ptr).DestroyQWebEnginePageDefault()
	}
}

func (ptr *QWebEnginePage) ConnectDestroyQWebEnginePage(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebEnginePage"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebEnginePage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebEnginePage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEnginePage) DisconnectDestroyQWebEnginePage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebEnginePage")
	}
}

func (ptr *QWebEnginePage) DestroyQWebEnginePage() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEnginePage_DestroyQWebEnginePage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEnginePage) DestroyQWebEnginePageDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEnginePage_DestroyQWebEnginePageDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEnginePage_DeleteLaterDefault(ptr.Pointer())
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

func NewQWebEngineProfile(parent core.QObject_ITF) *QWebEngineProfile {
	tmpValue := NewQWebEngineProfileFromPointer(C.QWebEngineProfile_NewQWebEngineProfile(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQWebEngineProfile2(name string, parent core.QObject_ITF) *QWebEngineProfile {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQWebEngineProfileFromPointer(C.QWebEngineProfile_NewQWebEngineProfile2(C.struct_QtWebEngine_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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

func (ptr *QWebEngineProfile) ClearVisitedLinks(urls []*core.QUrl) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_ClearVisitedLinks(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQWebEngineProfileFromPointer(NewQWebEngineProfileFromPointer(nil).__clearVisitedLinks_urls_newList())
			for _, v := range urls {
				tmpList.__clearVisitedLinks_urls_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QWebEngineProfile) ClientCertificateStore() *QWebEngineClientCertificateStore {
	if ptr.Pointer() != nil {
		return NewQWebEngineClientCertificateStoreFromPointer(C.QWebEngineProfile_ClientCertificateStore(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) CookieStore() *QWebEngineCookieStore {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEngineCookieStoreFromPointer(C.QWebEngineProfile_CookieStore(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QWebEngineProfile_DefaultProfile() *QWebEngineProfile {
	tmpValue := NewQWebEngineProfileFromPointer(C.QWebEngineProfile_QWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebEngineProfile) DefaultProfile() *QWebEngineProfile {
	tmpValue := NewQWebEngineProfileFromPointer(C.QWebEngineProfile_QWebEngineProfile_DefaultProfile())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebEngineProfile) DownloadPath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_DownloadPath(ptr.Pointer()))
	}
	return ""
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

func (ptr *QWebEngineProfile) InstallUrlSchemeHandler(scheme core.QByteArray_ITF, vqw QWebEngineUrlSchemeHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_InstallUrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(scheme), PointerFromQWebEngineUrlSchemeHandler(vqw))
	}
}

func (ptr *QWebEngineProfile) IsOffTheRecord() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineProfile_IsOffTheRecord(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) IsSpellCheckEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineProfile_IsSpellCheckEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebEngineProfile) IsUsedForGlobalCertificateVerification() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineProfile_IsUsedForGlobalCertificateVerification(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebEngineProfile_MetaObject
func callbackQWebEngineProfile_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineProfileFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineProfile) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineProfile_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func (ptr *QWebEngineProfile) RemoveUrlSchemeHandler(vqw QWebEngineUrlSchemeHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_RemoveUrlSchemeHandler(ptr.Pointer(), PointerFromQWebEngineUrlSchemeHandler(vqw))
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
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QWebEngineProfile_SetCachePath(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QWebEngineProfile) SetDownloadPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QWebEngineProfile_SetDownloadPath(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	if ptr.Pointer() != nil {
		var httpAcceptLanguageC *C.char
		if httpAcceptLanguage != "" {
			httpAcceptLanguageC = C.CString(httpAcceptLanguage)
			defer C.free(unsafe.Pointer(httpAcceptLanguageC))
		}
		C.QWebEngineProfile_SetHttpAcceptLanguage(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: httpAcceptLanguageC, len: C.longlong(len(httpAcceptLanguage))})
	}
}

func (ptr *QWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpCacheMaximumSize(ptr.Pointer(), C.int(int32(maxSize)))
	}
}

func (ptr *QWebEngineProfile) SetHttpCacheType(vqw QWebEngineProfile__HttpCacheType) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetHttpCacheType(ptr.Pointer(), C.longlong(vqw))
	}
}

func (ptr *QWebEngineProfile) SetHttpUserAgent(userAgent string) {
	if ptr.Pointer() != nil {
		var userAgentC *C.char
		if userAgent != "" {
			userAgentC = C.CString(userAgent)
			defer C.free(unsafe.Pointer(userAgentC))
		}
		C.QWebEngineProfile_SetHttpUserAgent(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: userAgentC, len: C.longlong(len(userAgent))})
	}
}

func (ptr *QWebEngineProfile) SetPersistentCookiesPolicy(vqw QWebEngineProfile__PersistentCookiesPolicy) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetPersistentCookiesPolicy(ptr.Pointer(), C.longlong(vqw))
	}
}

func (ptr *QWebEngineProfile) SetPersistentStoragePath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QWebEngineProfile_SetPersistentStoragePath(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QWebEngineProfile) SetRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

func (ptr *QWebEngineProfile) SetSpellCheckEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetSpellCheckEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QWebEngineProfile) SetSpellCheckLanguages(languages []string) {
	if ptr.Pointer() != nil {
		languagesC := C.CString(strings.Join(languages, "¡¦!"))
		defer C.free(unsafe.Pointer(languagesC))
		C.QWebEngineProfile_SetSpellCheckLanguages(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: languagesC, len: C.longlong(len(strings.Join(languages, "¡¦!")))})
	}
}

func (ptr *QWebEngineProfile) SetUrlRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetUrlRequestInterceptor(ptr.Pointer(), PointerFromQWebEngineUrlRequestInterceptor(interceptor))
	}
}

func (ptr *QWebEngineProfile) SetUseForGlobalCertificateVerification(enabled bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineProfile_SetUseForGlobalCertificateVerification(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QWebEngineProfile) Settings() *QWebEngineSettings {
	if ptr.Pointer() != nil {
		return NewQWebEngineSettingsFromPointer(C.QWebEngineProfile_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineProfile) SpellCheckLanguages() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QWebEngineProfile_SpellCheckLanguages(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QWebEngineProfile) StorageName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebEngineProfile_StorageName(ptr.Pointer()))
	}
	return ""
}

func QWebEngineProfile_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QWebEngineProfile_QWebEngineProfile_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebEngineProfile) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QWebEngineProfile_QWebEngineProfile_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebEngineProfile) UrlSchemeHandler(vqb core.QByteArray_ITF) *QWebEngineUrlSchemeHandler {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEngineUrlSchemeHandlerFromPointer(C.QWebEngineProfile_UrlSchemeHandler(ptr.Pointer(), core.PointerFromQByteArray(vqb)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineProfile) VisitedLinksContainsUrl(url core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineProfile_VisitedLinksContainsUrl(ptr.Pointer(), core.PointerFromQUrl(url))) != 0
	}
	return false
}

//export callbackQWebEngineProfile_DestroyQWebEngineProfile
func callbackQWebEngineProfile_DestroyQWebEngineProfile(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebEngineProfile"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineProfileFromPointer(ptr).DestroyQWebEngineProfileDefault()
	}
}

func (ptr *QWebEngineProfile) ConnectDestroyQWebEngineProfile(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebEngineProfile"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineProfile", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineProfile", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineProfile) DisconnectDestroyQWebEngineProfile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebEngineProfile")
	}
}

func (ptr *QWebEngineProfile) DestroyQWebEngineProfile() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineProfile_DestroyQWebEngineProfile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineProfile) DestroyQWebEngineProfileDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineProfile_DestroyQWebEngineProfileDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineProfile___clearVisitedLinks_urls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineProfile_DeleteLaterDefault(ptr.Pointer())
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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

func NewQWebEngineSettings2(parentSettings QWebEngineSettings_ITF) *QWebEngineSettings {
	tmpValue := NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_NewQWebEngineSettings2(PointerFromQWebEngineSettings(parentSettings)))
	qt.SetFinalizer(tmpValue, (*QWebEngineSettings).DestroyQWebEngineSettings)
	return tmpValue
}

func QWebEngineSettings_DefaultSettings() *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_DefaultSettings())
}

func (ptr *QWebEngineSettings) DefaultSettings() *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(C.QWebEngineSettings_QWebEngineSettings_DefaultSettings())
}

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

func (ptr *QWebEngineSettings) ResetAttribute(attr QWebEngineSettings__WebAttribute) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetAttribute(ptr.Pointer(), C.longlong(attr))
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

func (ptr *QWebEngineSettings) ResetUnknownUrlSchemePolicy() {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_ResetUnknownUrlSchemePolicy(ptr.Pointer())
	}
}

func (ptr *QWebEngineSettings) SetAttribute(attr QWebEngineSettings__WebAttribute, on bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetAttribute(ptr.Pointer(), C.longlong(attr), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QWebEngineSettings) SetDefaultTextEncoding(encoding string) {
	if ptr.Pointer() != nil {
		var encodingC *C.char
		if encoding != "" {
			encodingC = C.CString(encoding)
			defer C.free(unsafe.Pointer(encodingC))
		}
		C.QWebEngineSettings_SetDefaultTextEncoding(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: encodingC, len: C.longlong(len(encoding))})
	}
}

func (ptr *QWebEngineSettings) SetFontFamily(which QWebEngineSettings__FontFamily, family string) {
	if ptr.Pointer() != nil {
		var familyC *C.char
		if family != "" {
			familyC = C.CString(family)
			defer C.free(unsafe.Pointer(familyC))
		}
		C.QWebEngineSettings_SetFontFamily(ptr.Pointer(), C.longlong(which), C.struct_QtWebEngine_PackedString{data: familyC, len: C.longlong(len(family))})
	}
}

func (ptr *QWebEngineSettings) SetFontSize(ty QWebEngineSettings__FontSize, size int) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetFontSize(ptr.Pointer(), C.longlong(ty), C.int(int32(size)))
	}
}

func (ptr *QWebEngineSettings) SetUnknownUrlSchemePolicy(policy QWebEngineSettings__UnknownUrlSchemePolicy) {
	if ptr.Pointer() != nil {
		C.QWebEngineSettings_SetUnknownUrlSchemePolicy(ptr.Pointer(), C.longlong(policy))
	}
}

func (ptr *QWebEngineSettings) TestAttribute(attr QWebEngineSettings__WebAttribute) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineSettings_TestAttribute(ptr.Pointer(), C.longlong(attr))) != 0
	}
	return false
}

func (ptr *QWebEngineSettings) UnknownUrlSchemePolicy() QWebEngineSettings__UnknownUrlSchemePolicy {
	if ptr.Pointer() != nil {
		return QWebEngineSettings__UnknownUrlSchemePolicy(C.QWebEngineSettings_UnknownUrlSchemePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebEngineSettings) DestroyQWebEngineSettings() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineSettings_DestroyQWebEngineSettings(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQWebEngineUrlRequestInfoFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestInfo) {
	n = new(QWebEngineUrlRequestInfo)
	n.SetPointer(ptr)
	return
}
func (ptr *QWebEngineUrlRequestInfo) DestroyQWebEngineUrlRequestInfo() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineUrlRequestInfo_RequestUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineUrlRequestInterceptor_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQWebEngineUrlRequestInterceptor_MetaObject
func callbackQWebEngineUrlRequestInterceptor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlRequestInterceptorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlRequestInterceptor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestInterceptor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineUrlRequestJob_RequestUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_atList(v core.QByteArray_ITF, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebEngineUrlRequestJob___requestHeaders_atList(ptr.Pointer(), core.PointerFromQByteArray(v), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineUrlRequestJob_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQWebEngineUrlRequestJob_MetaObject
func callbackQWebEngineUrlRequestJob_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlRequestJobFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlRequestJob) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlRequestJob_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
	qt.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func NewQWebEngineUrlScheme2(name core.QByteArray_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme2(core.PointerFromQByteArray(name)))
	qt.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func NewQWebEngineUrlScheme3(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme3(PointerFromQWebEngineUrlScheme(that)))
	qt.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func NewQWebEngineUrlScheme4(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_NewQWebEngineUrlScheme4(PointerFromQWebEngineUrlScheme(that)))
	qt.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	qt.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
	return tmpValue
}

func (ptr *QWebEngineUrlScheme) SchemeByName(name core.QByteArray_ITF) *QWebEngineUrlScheme {
	tmpValue := NewQWebEngineUrlSchemeFromPointer(C.QWebEngineUrlScheme_QWebEngineUrlScheme_SchemeByName(core.PointerFromQByteArray(name)))
	qt.SetFinalizer(tmpValue, (*QWebEngineUrlScheme).DestroyQWebEngineUrlScheme)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineUrlScheme_DestroyQWebEngineUrlScheme(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandlerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineUrlSchemeHandler_DeleteLaterDefault(ptr.Pointer())
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

//export callbackQWebEngineUrlSchemeHandler_MetaObject
func callbackQWebEngineUrlSchemeHandler_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineUrlSchemeHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineUrlSchemeHandler) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineUrlSchemeHandler_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
func NewQWebEngineView(parent widgets.QWidget_ITF) *QWebEngineView {
	tmpValue := NewQWebEngineViewFromPointer(C.QWebEngineView_NewQWebEngineView(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebEngineView_Back
func callbackQWebEngineView_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "back"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).BackDefault()
	}
}

func (ptr *QWebEngineView) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "back"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "back", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "back", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectBack() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "back")
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

//export callbackQWebEngineView_ContextMenuEvent
func callbackQWebEngineView_ContextMenuEvent(ptr unsafe.Pointer, vqc unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(vqc))
	} else {
		NewQWebEngineViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(vqc))
	}
}

func (ptr *QWebEngineView) ContextMenuEventDefault(vqc gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(vqc))
	}
}

//export callbackQWebEngineView_DragEnterEvent
func callbackQWebEngineView_DragEnterEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

//export callbackQWebEngineView_DragLeaveEvent
func callbackQWebEngineView_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQWebEngineView_DragMoveEvent
func callbackQWebEngineView_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQWebEngineView_DropEvent
func callbackQWebEngineView_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQWebEngineViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QWebEngineView) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQWebEngineView_Event
func callbackQWebEngineView_Event(ptr unsafe.Pointer, vqe unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(vqe)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebEngineViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(vqe)))))
}

func (ptr *QWebEngineView) EventDefault(vqe core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebEngineView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(vqe))) != 0
	}
	return false
}

//export callbackQWebEngineView_Forward
func callbackQWebEngineView_Forward(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "forward"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ForwardDefault()
	}
}

func (ptr *QWebEngineView) ConnectForward(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "forward"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "forward", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "forward", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectForward() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "forward")
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
		return int8(C.QWebEngineView_HasSelection(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebEngineView_HideEvent
func callbackQWebEngineView_HideEvent(ptr unsafe.Pointer, vqh unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQWebEngineViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QWebEngineView) HideEventDefault(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QWebEngineView) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWebEngineView_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
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
			C.QWebEngineView_ConnectIconChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "iconChanged")))
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
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
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
			C.QWebEngineView_ConnectIconUrlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "iconUrlChanged")))
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

func (ptr *QWebEngineView) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebEngineView) Load2(request QWebEngineHttpRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_Load2(ptr.Pointer(), PointerFromQWebEngineHttpRequest(request))
	}
}

//export callbackQWebEngineView_LoadFinished
func callbackQWebEngineView_LoadFinished(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "loadFinished"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

func (ptr *QWebEngineView) ConnectLoadFinished(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadFinished") {
			C.QWebEngineView_ConnectLoadFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "loadFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadFinished"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadFinished")
	}
}

func (ptr *QWebEngineView) LoadFinished(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebEngineView_LoadProgress
func callbackQWebEngineView_LoadProgress(ptr unsafe.Pointer, progress C.int) {
	if signal := qt.GetSignal(ptr, "loadProgress"); signal != nil {
		(*(*func(int))(signal))(int(int32(progress)))
	}

}

func (ptr *QWebEngineView) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadProgress") {
			C.QWebEngineView_ConnectLoadProgress(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "loadProgress")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadProgress"); signal != nil {
			f := func(progress int) {
				(*(*func(int))(signal))(progress)
				f(progress)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadProgress")
	}
}

func (ptr *QWebEngineView) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQWebEngineView_LoadStarted
func callbackQWebEngineView_LoadStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadStarted"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEngineView) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadStarted") {
			C.QWebEngineView_ConnectLoadStarted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "loadStarted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadStarted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadStarted")
	}
}

func (ptr *QWebEngineView) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebEngineView_MetaObject
func callbackQWebEngineView_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebEngineViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebEngineView) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebEngineView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebEngineView) Page() *QWebEnginePage {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebEnginePageFromPointer(C.QWebEngineView_Page(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_Reload
func callbackQWebEngineView_Reload(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reload"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).ReloadDefault()
	}
}

func (ptr *QWebEngineView) ConnectReload(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reload"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "reload", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reload", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectReload() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reload")
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

//export callbackQWebEngineView_SelectionChanged
func callbackQWebEngineView_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebEngineView) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QWebEngineView_ConnectSelectionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QWebEngineView) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebEngineView) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		C.QWebEngineView_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.struct_QtWebEngine_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebEngineView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.QWebEngineView_SetHtml(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: htmlC, len: C.longlong(len(html))}, core.PointerFromQUrl(baseUrl))
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

//export callbackQWebEngineView_ShowEvent
func callbackQWebEngineView_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQWebEngineViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QWebEngineView) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
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
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_Stop
func callbackQWebEngineView_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).StopDefault()
	}
}

func (ptr *QWebEngineView) ConnectStop(f func()) {
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

func (ptr *QWebEngineView) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
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

//export callbackQWebEngineView_TitleChanged
func callbackQWebEngineView_TitleChanged(ptr unsafe.Pointer, title C.struct_QtWebEngine_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

func (ptr *QWebEngineView) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.QWebEngineView_ConnectTitleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "titleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			f := func(title string) {
				(*(*func(string))(signal))(title)
				f(title)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *QWebEngineView) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QWebEngineView_TitleChanged(ptr.Pointer(), C.struct_QtWebEngine_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func QWebEngineView_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QWebEngineView_QWebEngineView_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebEngineView) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QWebEngineView_QWebEngineView_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebEngineView) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebEngineView_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebEngineView_UrlChanged
func callbackQWebEngineView_UrlChanged(ptr unsafe.Pointer, vqu unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "urlChanged"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(vqu))
	}

}

func (ptr *QWebEngineView) ConnectUrlChanged(f func(vqu *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "urlChanged") {
			C.QWebEngineView_ConnectUrlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "urlChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "urlChanged"); signal != nil {
			f := func(vqu *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(vqu)
				f(vqu)
			}
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebEngineView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "urlChanged")
	}
}

func (ptr *QWebEngineView) UrlChanged(vqu core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebEngineView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(vqu))
	}
}

func (ptr *QWebEngineView) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebEngineView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebEngineView_DestroyQWebEngineView
func callbackQWebEngineView_DestroyQWebEngineView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebEngineView"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebEngineViewFromPointer(ptr).DestroyQWebEngineViewDefault()
	}
}

func (ptr *QWebEngineView) ConnectDestroyQWebEngineView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebEngineView"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineView", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebEngineView", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebEngineView) DisconnectDestroyQWebEngineView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebEngineView")
	}
}

func (ptr *QWebEngineView) DestroyQWebEngineView() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineView_DestroyQWebEngineView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebEngineView) DestroyQWebEngineViewDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineView_DestroyQWebEngineViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

//export callbackQWebEngineView_CustomContextMenuRequested
func callbackQWebEngineView_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
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
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
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

		qt.SetFinalizer(ptr, nil)
		C.QWebEngineView_DeleteLaterDefault(ptr.Pointer())
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type UserScriptData struct {
	ptr unsafe.Pointer
}

type UserScriptData_ITF interface {
	UserScriptData_PTR() *UserScriptData
}

func (ptr *UserScriptData) UserScriptData_PTR() *UserScriptData {
	return ptr
}

func (ptr *UserScriptData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *UserScriptData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromUserScriptData(ptr UserScriptData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.UserScriptData_PTR().Pointer()
	}
	return nil
}

func NewUserScriptDataFromPointer(ptr unsafe.Pointer) (n *UserScriptData) {
	n = new(UserScriptData)
	n.SetPointer(ptr)
	return
}
func (ptr *UserScriptData) DestroyUserScriptData() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func init() {
	qt.ItfMap["webengine.QQuickWebEngineProfile_ITF"] = QQuickWebEngineProfile{}
	qt.FuncMap["webengine.NewQQuickWebEngineProfile"] = NewQQuickWebEngineProfile
	qt.FuncMap["webengine.QQuickWebEngineProfile_DefaultProfile"] = QQuickWebEngineProfile_DefaultProfile
	qt.EnumMap["webengine.QQuickWebEngineProfile__MemoryHttpCache"] = int64(QQuickWebEngineProfile__MemoryHttpCache)
	qt.EnumMap["webengine.QQuickWebEngineProfile__DiskHttpCache"] = int64(QQuickWebEngineProfile__DiskHttpCache)
	qt.EnumMap["webengine.QQuickWebEngineProfile__NoCache"] = int64(QQuickWebEngineProfile__NoCache)
	qt.EnumMap["webengine.QQuickWebEngineProfile__NoPersistentCookies"] = int64(QQuickWebEngineProfile__NoPersistentCookies)
	qt.EnumMap["webengine.QQuickWebEngineProfile__AllowPersistentCookies"] = int64(QQuickWebEngineProfile__AllowPersistentCookies)
	qt.EnumMap["webengine.QQuickWebEngineProfile__ForcePersistentCookies"] = int64(QQuickWebEngineProfile__ForcePersistentCookies)
	qt.ItfMap["webengine.QQuickWebEngineScript_ITF"] = QQuickWebEngineScript{}
	qt.FuncMap["webengine.NewQQuickWebEngineScript"] = NewQQuickWebEngineScript
	qt.EnumMap["webengine.QQuickWebEngineScript__Deferred"] = int64(QQuickWebEngineScript__Deferred)
	qt.EnumMap["webengine.QQuickWebEngineScript__DocumentReady"] = int64(QQuickWebEngineScript__DocumentReady)
	qt.EnumMap["webengine.QQuickWebEngineScript__DocumentCreation"] = int64(QQuickWebEngineScript__DocumentCreation)
	qt.EnumMap["webengine.QQuickWebEngineScript__MainWorld"] = int64(QQuickWebEngineScript__MainWorld)
	qt.EnumMap["webengine.QQuickWebEngineScript__ApplicationWorld"] = int64(QQuickWebEngineScript__ApplicationWorld)
	qt.EnumMap["webengine.QQuickWebEngineScript__UserWorld"] = int64(QQuickWebEngineScript__UserWorld)
	qt.ItfMap["webengine.QWebEngineCallback_ITF"] = QWebEngineCallback{}
	qt.ItfMap["webengine.QWebEngineCertificateError_ITF"] = QWebEngineCertificateError{}
	qt.EnumMap["webengine.QWebEngineCertificateError__SslPinnedKeyNotInCertificateChain"] = int64(QWebEngineCertificateError__SslPinnedKeyNotInCertificateChain)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateCommonNameInvalid"] = int64(QWebEngineCertificateError__CertificateCommonNameInvalid)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateDateInvalid"] = int64(QWebEngineCertificateError__CertificateDateInvalid)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateAuthorityInvalid"] = int64(QWebEngineCertificateError__CertificateAuthorityInvalid)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateContainsErrors"] = int64(QWebEngineCertificateError__CertificateContainsErrors)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateNoRevocationMechanism"] = int64(QWebEngineCertificateError__CertificateNoRevocationMechanism)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateUnableToCheckRevocation"] = int64(QWebEngineCertificateError__CertificateUnableToCheckRevocation)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateRevoked"] = int64(QWebEngineCertificateError__CertificateRevoked)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateInvalid"] = int64(QWebEngineCertificateError__CertificateInvalid)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateWeakSignatureAlgorithm"] = int64(QWebEngineCertificateError__CertificateWeakSignatureAlgorithm)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateNonUniqueName"] = int64(QWebEngineCertificateError__CertificateNonUniqueName)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateWeakKey"] = int64(QWebEngineCertificateError__CertificateWeakKey)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateNameConstraintViolation"] = int64(QWebEngineCertificateError__CertificateNameConstraintViolation)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateValidityTooLong"] = int64(QWebEngineCertificateError__CertificateValidityTooLong)
	qt.EnumMap["webengine.QWebEngineCertificateError__CertificateTransparencyRequired"] = int64(QWebEngineCertificateError__CertificateTransparencyRequired)
	qt.ItfMap["webengine.QWebEngineClientCertificateSelection_ITF"] = QWebEngineClientCertificateSelection{}
	qt.FuncMap["webengine.NewQWebEngineClientCertificateSelection"] = NewQWebEngineClientCertificateSelection
	qt.ItfMap["webengine.QWebEngineClientCertificateStore_ITF"] = QWebEngineClientCertificateStore{}
	qt.ItfMap["webengine.QWebEngineContextMenuData_ITF"] = QWebEngineContextMenuData{}
	qt.ItfMap["webengine.QWebEngineCookieStore_ITF"] = QWebEngineCookieStore{}
	qt.ItfMap["webengine.QWebEngineHttpRequest_ITF"] = QWebEngineHttpRequest{}
	qt.FuncMap["webengine.NewQWebEngineHttpRequest"] = NewQWebEngineHttpRequest
	qt.FuncMap["webengine.NewQWebEngineHttpRequest2"] = NewQWebEngineHttpRequest2
	qt.FuncMap["webengine.QWebEngineHttpRequest_PostRequest"] = QWebEngineHttpRequest_PostRequest
	qt.EnumMap["webengine.QWebEngineHttpRequest__Get"] = int64(QWebEngineHttpRequest__Get)
	qt.EnumMap["webengine.QWebEngineHttpRequest__Post"] = int64(QWebEngineHttpRequest__Post)
	qt.ItfMap["webengine.QWebEngineNotification_ITF"] = QWebEngineNotification{}
	qt.ItfMap["webengine.QWebEnginePage_ITF"] = QWebEnginePage{}
	qt.FuncMap["webengine.NewQWebEnginePage"] = NewQWebEnginePage
	qt.FuncMap["webengine.NewQWebEnginePage2"] = NewQWebEnginePage2
	qt.FuncMap["webengine.QWebEnginePage_Tr"] = QWebEnginePage_Tr
	qt.EnumMap["webengine.QWebEnginePage__FileSelectOpen"] = int64(QWebEnginePage__FileSelectOpen)
	qt.EnumMap["webengine.QWebEnginePage__FileSelectOpenMultiple"] = int64(QWebEnginePage__FileSelectOpenMultiple)
	qt.EnumMap["webengine.QWebEnginePage__InfoMessageLevel"] = int64(QWebEnginePage__InfoMessageLevel)
	qt.EnumMap["webengine.QWebEnginePage__WarningMessageLevel"] = int64(QWebEnginePage__WarningMessageLevel)
	qt.EnumMap["webengine.QWebEnginePage__ErrorMessageLevel"] = int64(QWebEnginePage__ErrorMessageLevel)
	qt.EnumMap["webengine.QWebEnginePage__Back"] = int64(QWebEnginePage__Back)
	qt.EnumMap["webengine.QWebEnginePage__Forward"] = int64(QWebEnginePage__Forward)
	qt.EnumMap["webengine.QWebEnginePage__Stop"] = int64(QWebEnginePage__Stop)
	qt.EnumMap["webengine.QWebEnginePage__Reload"] = int64(QWebEnginePage__Reload)
	qt.EnumMap["webengine.QWebEnginePage__Cut"] = int64(QWebEnginePage__Cut)
	qt.EnumMap["webengine.QWebEnginePage__Copy"] = int64(QWebEnginePage__Copy)
	qt.EnumMap["webengine.QWebEnginePage__Paste"] = int64(QWebEnginePage__Paste)
	qt.EnumMap["webengine.QWebEnginePage__Undo"] = int64(QWebEnginePage__Undo)
	qt.EnumMap["webengine.QWebEnginePage__Redo"] = int64(QWebEnginePage__Redo)
	qt.EnumMap["webengine.QWebEnginePage__SelectAll"] = int64(QWebEnginePage__SelectAll)
	qt.EnumMap["webengine.QWebEnginePage__ReloadAndBypassCache"] = int64(QWebEnginePage__ReloadAndBypassCache)
	qt.EnumMap["webengine.QWebEnginePage__PasteAndMatchStyle"] = int64(QWebEnginePage__PasteAndMatchStyle)
	qt.EnumMap["webengine.QWebEnginePage__OpenLinkInThisWindow"] = int64(QWebEnginePage__OpenLinkInThisWindow)
	qt.EnumMap["webengine.QWebEnginePage__OpenLinkInNewWindow"] = int64(QWebEnginePage__OpenLinkInNewWindow)
	qt.EnumMap["webengine.QWebEnginePage__OpenLinkInNewTab"] = int64(QWebEnginePage__OpenLinkInNewTab)
	qt.EnumMap["webengine.QWebEnginePage__CopyLinkToClipboard"] = int64(QWebEnginePage__CopyLinkToClipboard)
	qt.EnumMap["webengine.QWebEnginePage__DownloadLinkToDisk"] = int64(QWebEnginePage__DownloadLinkToDisk)
	qt.EnumMap["webengine.QWebEnginePage__CopyImageToClipboard"] = int64(QWebEnginePage__CopyImageToClipboard)
	qt.EnumMap["webengine.QWebEnginePage__CopyImageUrlToClipboard"] = int64(QWebEnginePage__CopyImageUrlToClipboard)
	qt.EnumMap["webengine.QWebEnginePage__DownloadImageToDisk"] = int64(QWebEnginePage__DownloadImageToDisk)
	qt.EnumMap["webengine.QWebEnginePage__CopyMediaUrlToClipboard"] = int64(QWebEnginePage__CopyMediaUrlToClipboard)
	qt.EnumMap["webengine.QWebEnginePage__ToggleMediaControls"] = int64(QWebEnginePage__ToggleMediaControls)
	qt.EnumMap["webengine.QWebEnginePage__ToggleMediaLoop"] = int64(QWebEnginePage__ToggleMediaLoop)
	qt.EnumMap["webengine.QWebEnginePage__ToggleMediaPlayPause"] = int64(QWebEnginePage__ToggleMediaPlayPause)
	qt.EnumMap["webengine.QWebEnginePage__ToggleMediaMute"] = int64(QWebEnginePage__ToggleMediaMute)
	qt.EnumMap["webengine.QWebEnginePage__DownloadMediaToDisk"] = int64(QWebEnginePage__DownloadMediaToDisk)
	qt.EnumMap["webengine.QWebEnginePage__InspectElement"] = int64(QWebEnginePage__InspectElement)
	qt.EnumMap["webengine.QWebEnginePage__ExitFullScreen"] = int64(QWebEnginePage__ExitFullScreen)
	qt.EnumMap["webengine.QWebEnginePage__RequestClose"] = int64(QWebEnginePage__RequestClose)
	qt.EnumMap["webengine.QWebEnginePage__Unselect"] = int64(QWebEnginePage__Unselect)
	qt.EnumMap["webengine.QWebEnginePage__SavePage"] = int64(QWebEnginePage__SavePage)
	qt.EnumMap["webengine.QWebEnginePage__OpenLinkInNewBackgroundTab"] = int64(QWebEnginePage__OpenLinkInNewBackgroundTab)
	qt.EnumMap["webengine.QWebEnginePage__ViewSource"] = int64(QWebEnginePage__ViewSource)
	qt.EnumMap["webengine.QWebEnginePage__ToggleBold"] = int64(QWebEnginePage__ToggleBold)
	qt.EnumMap["webengine.QWebEnginePage__ToggleItalic"] = int64(QWebEnginePage__ToggleItalic)
	qt.EnumMap["webengine.QWebEnginePage__ToggleUnderline"] = int64(QWebEnginePage__ToggleUnderline)
	qt.EnumMap["webengine.QWebEnginePage__ToggleStrikethrough"] = int64(QWebEnginePage__ToggleStrikethrough)
	qt.EnumMap["webengine.QWebEnginePage__AlignLeft"] = int64(QWebEnginePage__AlignLeft)
	qt.EnumMap["webengine.QWebEnginePage__AlignCenter"] = int64(QWebEnginePage__AlignCenter)
	qt.EnumMap["webengine.QWebEnginePage__AlignRight"] = int64(QWebEnginePage__AlignRight)
	qt.EnumMap["webengine.QWebEnginePage__AlignJustified"] = int64(QWebEnginePage__AlignJustified)
	qt.EnumMap["webengine.QWebEnginePage__Indent"] = int64(QWebEnginePage__Indent)
	qt.EnumMap["webengine.QWebEnginePage__Outdent"] = int64(QWebEnginePage__Outdent)
	qt.EnumMap["webengine.QWebEnginePage__InsertOrderedList"] = int64(QWebEnginePage__InsertOrderedList)
	qt.EnumMap["webengine.QWebEnginePage__InsertUnorderedList"] = int64(QWebEnginePage__InsertUnorderedList)
	qt.EnumMap["webengine.QWebEnginePage__WebActionCount"] = int64(QWebEnginePage__WebActionCount)
	qt.EnumMap["webengine.QWebEnginePage__WebBrowserWindow"] = int64(QWebEnginePage__WebBrowserWindow)
	qt.EnumMap["webengine.QWebEnginePage__WebBrowserTab"] = int64(QWebEnginePage__WebBrowserTab)
	qt.EnumMap["webengine.QWebEnginePage__WebDialog"] = int64(QWebEnginePage__WebDialog)
	qt.EnumMap["webengine.QWebEnginePage__WebBrowserBackgroundTab"] = int64(QWebEnginePage__WebBrowserBackgroundTab)
	qt.EnumMap["webengine.QWebEnginePage__PermissionUnknown"] = int64(QWebEnginePage__PermissionUnknown)
	qt.EnumMap["webengine.QWebEnginePage__PermissionGrantedByUser"] = int64(QWebEnginePage__PermissionGrantedByUser)
	qt.EnumMap["webengine.QWebEnginePage__PermissionDeniedByUser"] = int64(QWebEnginePage__PermissionDeniedByUser)
	qt.EnumMap["webengine.QWebEnginePage__NavigationTypeLinkClicked"] = int64(QWebEnginePage__NavigationTypeLinkClicked)
	qt.EnumMap["webengine.QWebEnginePage__NavigationTypeTyped"] = int64(QWebEnginePage__NavigationTypeTyped)
	qt.EnumMap["webengine.QWebEnginePage__NavigationTypeFormSubmitted"] = int64(QWebEnginePage__NavigationTypeFormSubmitted)
	qt.EnumMap["webengine.QWebEnginePage__NavigationTypeBackForward"] = int64(QWebEnginePage__NavigationTypeBackForward)
	qt.EnumMap["webengine.QWebEnginePage__NavigationTypeReload"] = int64(QWebEnginePage__NavigationTypeReload)
	qt.EnumMap["webengine.QWebEnginePage__NavigationTypeOther"] = int64(QWebEnginePage__NavigationTypeOther)
	qt.EnumMap["webengine.QWebEnginePage__FindBackward"] = int64(QWebEnginePage__FindBackward)
	qt.EnumMap["webengine.QWebEnginePage__FindCaseSensitively"] = int64(QWebEnginePage__FindCaseSensitively)
	qt.EnumMap["webengine.QWebEnginePage__Notifications"] = int64(QWebEnginePage__Notifications)
	qt.EnumMap["webengine.QWebEnginePage__Geolocation"] = int64(QWebEnginePage__Geolocation)
	qt.EnumMap["webengine.QWebEnginePage__MediaAudioCapture"] = int64(QWebEnginePage__MediaAudioCapture)
	qt.EnumMap["webengine.QWebEnginePage__MediaVideoCapture"] = int64(QWebEnginePage__MediaVideoCapture)
	qt.EnumMap["webengine.QWebEnginePage__MediaAudioVideoCapture"] = int64(QWebEnginePage__MediaAudioVideoCapture)
	qt.EnumMap["webengine.QWebEnginePage__MouseLock"] = int64(QWebEnginePage__MouseLock)
	qt.EnumMap["webengine.QWebEnginePage__DesktopVideoCapture"] = int64(QWebEnginePage__DesktopVideoCapture)
	qt.EnumMap["webengine.QWebEnginePage__DesktopAudioVideoCapture"] = int64(QWebEnginePage__DesktopAudioVideoCapture)
	qt.ItfMap["webengine.QWebEngineProfile_ITF"] = QWebEngineProfile{}
	qt.FuncMap["webengine.NewQWebEngineProfile"] = NewQWebEngineProfile
	qt.FuncMap["webengine.NewQWebEngineProfile2"] = NewQWebEngineProfile2
	qt.FuncMap["webengine.QWebEngineProfile_DefaultProfile"] = QWebEngineProfile_DefaultProfile
	qt.FuncMap["webengine.QWebEngineProfile_Tr"] = QWebEngineProfile_Tr
	qt.EnumMap["webengine.QWebEngineProfile__MemoryHttpCache"] = int64(QWebEngineProfile__MemoryHttpCache)
	qt.EnumMap["webengine.QWebEngineProfile__DiskHttpCache"] = int64(QWebEngineProfile__DiskHttpCache)
	qt.EnumMap["webengine.QWebEngineProfile__NoCache"] = int64(QWebEngineProfile__NoCache)
	qt.EnumMap["webengine.QWebEngineProfile__NoPersistentCookies"] = int64(QWebEngineProfile__NoPersistentCookies)
	qt.EnumMap["webengine.QWebEngineProfile__AllowPersistentCookies"] = int64(QWebEngineProfile__AllowPersistentCookies)
	qt.EnumMap["webengine.QWebEngineProfile__ForcePersistentCookies"] = int64(QWebEngineProfile__ForcePersistentCookies)
	qt.ItfMap["webengine.QWebEngineQuotaRequest_ITF"] = QWebEngineQuotaRequest{}
	qt.ItfMap["webengine.QWebEngineRegisterProtocolHandlerRequest_ITF"] = QWebEngineRegisterProtocolHandlerRequest{}
	qt.ItfMap["webengine.QWebEngineScript_ITF"] = QWebEngineScript{}
	qt.EnumMap["webengine.QWebEngineScript__Deferred"] = int64(QWebEngineScript__Deferred)
	qt.EnumMap["webengine.QWebEngineScript__DocumentReady"] = int64(QWebEngineScript__DocumentReady)
	qt.EnumMap["webengine.QWebEngineScript__DocumentCreation"] = int64(QWebEngineScript__DocumentCreation)
	qt.EnumMap["webengine.QWebEngineScript__MainWorld"] = int64(QWebEngineScript__MainWorld)
	qt.EnumMap["webengine.QWebEngineScript__ApplicationWorld"] = int64(QWebEngineScript__ApplicationWorld)
	qt.EnumMap["webengine.QWebEngineScript__UserWorld"] = int64(QWebEngineScript__UserWorld)
	qt.ItfMap["webengine.QWebEngineScriptCollection_ITF"] = QWebEngineScriptCollection{}
	qt.ItfMap["webengine.QWebEngineSettings_ITF"] = QWebEngineSettings{}
	qt.FuncMap["webengine.NewQWebEngineSettings2"] = NewQWebEngineSettings2
	qt.FuncMap["webengine.QWebEngineSettings_DefaultSettings"] = QWebEngineSettings_DefaultSettings
	qt.FuncMap["webengine.QWebEngineSettings_GlobalSettings"] = QWebEngineSettings_GlobalSettings
	qt.EnumMap["webengine.QWebEngineSettings__StandardFont"] = int64(QWebEngineSettings__StandardFont)
	qt.EnumMap["webengine.QWebEngineSettings__FixedFont"] = int64(QWebEngineSettings__FixedFont)
	qt.EnumMap["webengine.QWebEngineSettings__SerifFont"] = int64(QWebEngineSettings__SerifFont)
	qt.EnumMap["webengine.QWebEngineSettings__SansSerifFont"] = int64(QWebEngineSettings__SansSerifFont)
	qt.EnumMap["webengine.QWebEngineSettings__CursiveFont"] = int64(QWebEngineSettings__CursiveFont)
	qt.EnumMap["webengine.QWebEngineSettings__FantasyFont"] = int64(QWebEngineSettings__FantasyFont)
	qt.EnumMap["webengine.QWebEngineSettings__PictographFont"] = int64(QWebEngineSettings__PictographFont)
	qt.EnumMap["webengine.QWebEngineSettings__AutoLoadImages"] = int64(QWebEngineSettings__AutoLoadImages)
	qt.EnumMap["webengine.QWebEngineSettings__JavascriptEnabled"] = int64(QWebEngineSettings__JavascriptEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__JavascriptCanOpenWindows"] = int64(QWebEngineSettings__JavascriptCanOpenWindows)
	qt.EnumMap["webengine.QWebEngineSettings__JavascriptCanAccessClipboard"] = int64(QWebEngineSettings__JavascriptCanAccessClipboard)
	qt.EnumMap["webengine.QWebEngineSettings__LinksIncludedInFocusChain"] = int64(QWebEngineSettings__LinksIncludedInFocusChain)
	qt.EnumMap["webengine.QWebEngineSettings__LocalStorageEnabled"] = int64(QWebEngineSettings__LocalStorageEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__LocalContentCanAccessRemoteUrls"] = int64(QWebEngineSettings__LocalContentCanAccessRemoteUrls)
	qt.EnumMap["webengine.QWebEngineSettings__XSSAuditingEnabled"] = int64(QWebEngineSettings__XSSAuditingEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__SpatialNavigationEnabled"] = int64(QWebEngineSettings__SpatialNavigationEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__LocalContentCanAccessFileUrls"] = int64(QWebEngineSettings__LocalContentCanAccessFileUrls)
	qt.EnumMap["webengine.QWebEngineSettings__HyperlinkAuditingEnabled"] = int64(QWebEngineSettings__HyperlinkAuditingEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__ScrollAnimatorEnabled"] = int64(QWebEngineSettings__ScrollAnimatorEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__ErrorPageEnabled"] = int64(QWebEngineSettings__ErrorPageEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__PluginsEnabled"] = int64(QWebEngineSettings__PluginsEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__FullScreenSupportEnabled"] = int64(QWebEngineSettings__FullScreenSupportEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__ScreenCaptureEnabled"] = int64(QWebEngineSettings__ScreenCaptureEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__WebGLEnabled"] = int64(QWebEngineSettings__WebGLEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__Accelerated2dCanvasEnabled"] = int64(QWebEngineSettings__Accelerated2dCanvasEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__AutoLoadIconsForPage"] = int64(QWebEngineSettings__AutoLoadIconsForPage)
	qt.EnumMap["webengine.QWebEngineSettings__TouchIconsEnabled"] = int64(QWebEngineSettings__TouchIconsEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__FocusOnNavigationEnabled"] = int64(QWebEngineSettings__FocusOnNavigationEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__PrintElementBackgrounds"] = int64(QWebEngineSettings__PrintElementBackgrounds)
	qt.EnumMap["webengine.QWebEngineSettings__AllowRunningInsecureContent"] = int64(QWebEngineSettings__AllowRunningInsecureContent)
	qt.EnumMap["webengine.QWebEngineSettings__AllowGeolocationOnInsecureOrigins"] = int64(QWebEngineSettings__AllowGeolocationOnInsecureOrigins)
	qt.EnumMap["webengine.QWebEngineSettings__AllowWindowActivationFromJavaScript"] = int64(QWebEngineSettings__AllowWindowActivationFromJavaScript)
	qt.EnumMap["webengine.QWebEngineSettings__ShowScrollBars"] = int64(QWebEngineSettings__ShowScrollBars)
	qt.EnumMap["webengine.QWebEngineSettings__PlaybackRequiresUserGesture"] = int64(QWebEngineSettings__PlaybackRequiresUserGesture)
	qt.EnumMap["webengine.QWebEngineSettings__WebRTCPublicInterfacesOnly"] = int64(QWebEngineSettings__WebRTCPublicInterfacesOnly)
	qt.EnumMap["webengine.QWebEngineSettings__JavascriptCanPaste"] = int64(QWebEngineSettings__JavascriptCanPaste)
	qt.EnumMap["webengine.QWebEngineSettings__DnsPrefetchEnabled"] = int64(QWebEngineSettings__DnsPrefetchEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__PdfViewerEnabled"] = int64(QWebEngineSettings__PdfViewerEnabled)
	qt.EnumMap["webengine.QWebEngineSettings__MinimumFontSize"] = int64(QWebEngineSettings__MinimumFontSize)
	qt.EnumMap["webengine.QWebEngineSettings__MinimumLogicalFontSize"] = int64(QWebEngineSettings__MinimumLogicalFontSize)
	qt.EnumMap["webengine.QWebEngineSettings__DefaultFontSize"] = int64(QWebEngineSettings__DefaultFontSize)
	qt.EnumMap["webengine.QWebEngineSettings__DefaultFixedFontSize"] = int64(QWebEngineSettings__DefaultFixedFontSize)
	qt.EnumMap["webengine.QWebEngineSettings__DisallowUnknownUrlSchemes"] = int64(QWebEngineSettings__DisallowUnknownUrlSchemes)
	qt.EnumMap["webengine.QWebEngineSettings__AllowUnknownUrlSchemesFromUserInteraction"] = int64(QWebEngineSettings__AllowUnknownUrlSchemesFromUserInteraction)
	qt.EnumMap["webengine.QWebEngineSettings__AllowAllUnknownUrlSchemes"] = int64(QWebEngineSettings__AllowAllUnknownUrlSchemes)
	qt.ItfMap["webengine.QWebEngineUrlRequestInfo_ITF"] = QWebEngineUrlRequestInfo{}
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeMainFrame"] = int64(QWebEngineUrlRequestInfo__ResourceTypeMainFrame)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeSubFrame"] = int64(QWebEngineUrlRequestInfo__ResourceTypeSubFrame)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeStylesheet"] = int64(QWebEngineUrlRequestInfo__ResourceTypeStylesheet)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeScript"] = int64(QWebEngineUrlRequestInfo__ResourceTypeScript)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeImage"] = int64(QWebEngineUrlRequestInfo__ResourceTypeImage)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeFontResource"] = int64(QWebEngineUrlRequestInfo__ResourceTypeFontResource)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeSubResource"] = int64(QWebEngineUrlRequestInfo__ResourceTypeSubResource)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeObject"] = int64(QWebEngineUrlRequestInfo__ResourceTypeObject)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeMedia"] = int64(QWebEngineUrlRequestInfo__ResourceTypeMedia)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeWorker"] = int64(QWebEngineUrlRequestInfo__ResourceTypeWorker)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeSharedWorker"] = int64(QWebEngineUrlRequestInfo__ResourceTypeSharedWorker)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypePrefetch"] = int64(QWebEngineUrlRequestInfo__ResourceTypePrefetch)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeFavicon"] = int64(QWebEngineUrlRequestInfo__ResourceTypeFavicon)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeXhr"] = int64(QWebEngineUrlRequestInfo__ResourceTypeXhr)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypePing"] = int64(QWebEngineUrlRequestInfo__ResourceTypePing)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeServiceWorker"] = int64(QWebEngineUrlRequestInfo__ResourceTypeServiceWorker)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeCspReport"] = int64(QWebEngineUrlRequestInfo__ResourceTypeCspReport)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypePluginResource"] = int64(QWebEngineUrlRequestInfo__ResourceTypePluginResource)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__ResourceTypeUnknown"] = int64(QWebEngineUrlRequestInfo__ResourceTypeUnknown)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__NavigationTypeLink"] = int64(QWebEngineUrlRequestInfo__NavigationTypeLink)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__NavigationTypeTyped"] = int64(QWebEngineUrlRequestInfo__NavigationTypeTyped)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__NavigationTypeFormSubmitted"] = int64(QWebEngineUrlRequestInfo__NavigationTypeFormSubmitted)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__NavigationTypeBackForward"] = int64(QWebEngineUrlRequestInfo__NavigationTypeBackForward)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__NavigationTypeReload"] = int64(QWebEngineUrlRequestInfo__NavigationTypeReload)
	qt.EnumMap["webengine.QWebEngineUrlRequestInfo__NavigationTypeOther"] = int64(QWebEngineUrlRequestInfo__NavigationTypeOther)
	qt.ItfMap["webengine.QWebEngineUrlRequestInterceptor_ITF"] = QWebEngineUrlRequestInterceptor{}
	qt.FuncMap["webengine.NewQWebEngineUrlRequestInterceptor2"] = NewQWebEngineUrlRequestInterceptor2
	qt.ItfMap["webengine.QWebEngineUrlRequestJob_ITF"] = QWebEngineUrlRequestJob{}
	qt.EnumMap["webengine.QWebEngineUrlRequestJob__NoError"] = int64(QWebEngineUrlRequestJob__NoError)
	qt.EnumMap["webengine.QWebEngineUrlRequestJob__UrlNotFound"] = int64(QWebEngineUrlRequestJob__UrlNotFound)
	qt.EnumMap["webengine.QWebEngineUrlRequestJob__UrlInvalid"] = int64(QWebEngineUrlRequestJob__UrlInvalid)
	qt.EnumMap["webengine.QWebEngineUrlRequestJob__RequestAborted"] = int64(QWebEngineUrlRequestJob__RequestAborted)
	qt.EnumMap["webengine.QWebEngineUrlRequestJob__RequestDenied"] = int64(QWebEngineUrlRequestJob__RequestDenied)
	qt.EnumMap["webengine.QWebEngineUrlRequestJob__RequestFailed"] = int64(QWebEngineUrlRequestJob__RequestFailed)
	qt.ItfMap["webengine.QWebEngineUrlScheme_ITF"] = QWebEngineUrlScheme{}
	qt.FuncMap["webengine.NewQWebEngineUrlScheme"] = NewQWebEngineUrlScheme
	qt.FuncMap["webengine.NewQWebEngineUrlScheme2"] = NewQWebEngineUrlScheme2
	qt.FuncMap["webengine.NewQWebEngineUrlScheme3"] = NewQWebEngineUrlScheme3
	qt.FuncMap["webengine.NewQWebEngineUrlScheme4"] = NewQWebEngineUrlScheme4
	qt.FuncMap["webengine.QWebEngineUrlScheme_RegisterScheme"] = QWebEngineUrlScheme_RegisterScheme
	qt.FuncMap["webengine.QWebEngineUrlScheme_SchemeByName"] = QWebEngineUrlScheme_SchemeByName
	qt.EnumMap["webengine.QWebEngineUrlScheme__HostPortAndUserInformation"] = int64(QWebEngineUrlScheme__HostPortAndUserInformation)
	qt.EnumMap["webengine.QWebEngineUrlScheme__HostAndPort"] = int64(QWebEngineUrlScheme__HostAndPort)
	qt.EnumMap["webengine.QWebEngineUrlScheme__Host"] = int64(QWebEngineUrlScheme__Host)
	qt.EnumMap["webengine.QWebEngineUrlScheme__Path"] = int64(QWebEngineUrlScheme__Path)
	qt.EnumMap["webengine.QWebEngineUrlScheme__PortUnspecified"] = int64(QWebEngineUrlScheme__PortUnspecified)
	qt.EnumMap["webengine.QWebEngineUrlScheme__SecureScheme"] = int64(QWebEngineUrlScheme__SecureScheme)
	qt.EnumMap["webengine.QWebEngineUrlScheme__LocalScheme"] = int64(QWebEngineUrlScheme__LocalScheme)
	qt.EnumMap["webengine.QWebEngineUrlScheme__LocalAccessAllowed"] = int64(QWebEngineUrlScheme__LocalAccessAllowed)
	qt.EnumMap["webengine.QWebEngineUrlScheme__NoAccessAllowed"] = int64(QWebEngineUrlScheme__NoAccessAllowed)
	qt.EnumMap["webengine.QWebEngineUrlScheme__ServiceWorkersAllowed"] = int64(QWebEngineUrlScheme__ServiceWorkersAllowed)
	qt.EnumMap["webengine.QWebEngineUrlScheme__ViewSourceAllowed"] = int64(QWebEngineUrlScheme__ViewSourceAllowed)
	qt.EnumMap["webengine.QWebEngineUrlScheme__ContentSecurityPolicyIgnored"] = int64(QWebEngineUrlScheme__ContentSecurityPolicyIgnored)
	qt.ItfMap["webengine.QWebEngineUrlSchemeHandler_ITF"] = QWebEngineUrlSchemeHandler{}
	qt.FuncMap["webengine.NewQWebEngineUrlSchemeHandler"] = NewQWebEngineUrlSchemeHandler
	qt.ItfMap["webengine.QWebEngineView_ITF"] = QWebEngineView{}
	qt.FuncMap["webengine.NewQWebEngineView"] = NewQWebEngineView
	qt.FuncMap["webengine.QWebEngineView_Tr"] = QWebEngineView_Tr
	qt.ItfMap["webengine.QtWebEngine_ITF"] = QtWebEngine{}
	qt.FuncMap["webengine.QtWebEngine_Initialize"] = QtWebEngine_Initialize
}
