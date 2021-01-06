// +build !minimal

package webengine

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/gui"
	"github.com/dev-drprasad/qt/internal"
	"github.com/dev-drprasad/qt/network"
	"github.com/dev-drprasad/qt/printsupport"
	"github.com/dev-drprasad/qt/webchannel"
	"github.com/dev-drprasad/qt/widgets"
	"unsafe"
)

type CertificateErrorController struct {
	internal.Internal
}

type CertificateErrorController_ITF interface {
	CertificateErrorController_PTR() *CertificateErrorController
}

func (ptr *CertificateErrorController) CertificateErrorController_PTR() *CertificateErrorController {
	return ptr
}

func (ptr *CertificateErrorController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *CertificateErrorController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromCertificateErrorController(ptr CertificateErrorController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CertificateErrorController_PTR().Pointer()
	}
	return nil
}

func (n *CertificateErrorController) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewCertificateErrorControllerFromPointer(ptr unsafe.Pointer) (n *CertificateErrorController) {
	n = new(CertificateErrorController)
	n.InitFromInternal(uintptr(ptr), "webengine.CertificateErrorController")
	return
}

func (ptr *CertificateErrorController) DestroyCertificateErrorController() {
}

type ClientCertSelectController struct {
	internal.Internal
}

type ClientCertSelectController_ITF interface {
	ClientCertSelectController_PTR() *ClientCertSelectController
}

func (ptr *ClientCertSelectController) ClientCertSelectController_PTR() *ClientCertSelectController {
	return ptr
}

func (ptr *ClientCertSelectController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *ClientCertSelectController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromClientCertSelectController(ptr ClientCertSelectController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ClientCertSelectController_PTR().Pointer()
	}
	return nil
}

func (n *ClientCertSelectController) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewClientCertSelectControllerFromPointer(ptr unsafe.Pointer) (n *ClientCertSelectController) {
	n = new(ClientCertSelectController)
	n.InitFromInternal(uintptr(ptr), "webengine.ClientCertSelectController")
	return
}

func (ptr *ClientCertSelectController) DestroyClientCertSelectController() {
}

type CommandLinePrefStoreQt struct {
	internal.Internal
}

type CommandLinePrefStoreQt_ITF interface {
	CommandLinePrefStoreQt_PTR() *CommandLinePrefStoreQt
}

func (ptr *CommandLinePrefStoreQt) CommandLinePrefStoreQt_PTR() *CommandLinePrefStoreQt {
	return ptr
}

func (ptr *CommandLinePrefStoreQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *CommandLinePrefStoreQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromCommandLinePrefStoreQt(ptr CommandLinePrefStoreQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CommandLinePrefStoreQt_PTR().Pointer()
	}
	return nil
}

func (n *CommandLinePrefStoreQt) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewCommandLinePrefStoreQtFromPointer(ptr unsafe.Pointer) (n *CommandLinePrefStoreQt) {
	n = new(CommandLinePrefStoreQt)
	n.InitFromInternal(uintptr(ptr), "webengine.CommandLinePrefStoreQt")
	return
}

func (ptr *CommandLinePrefStoreQt) DestroyCommandLinePrefStoreQt() {
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

func (n *GLContextHelper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *GLContextHelper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewGLContextHelperFromPointer(ptr unsafe.Pointer) (n *GLContextHelper) {
	n = new(GLContextHelper)
	n.InitFromInternal(uintptr(ptr), "webengine.GLContextHelper")
	return
}

func (ptr *GLContextHelper) DestroyGLContextHelper() {
}

type ProxyConfigServiceQt struct {
	internal.Internal
}

type ProxyConfigServiceQt_ITF interface {
	ProxyConfigServiceQt_PTR() *ProxyConfigServiceQt
}

func (ptr *ProxyConfigServiceQt) ProxyConfigServiceQt_PTR() *ProxyConfigServiceQt {
	return ptr
}

func (ptr *ProxyConfigServiceQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *ProxyConfigServiceQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromProxyConfigServiceQt(ptr ProxyConfigServiceQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ProxyConfigServiceQt_PTR().Pointer()
	}
	return nil
}

func (n *ProxyConfigServiceQt) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewProxyConfigServiceQtFromPointer(ptr unsafe.Pointer) (n *ProxyConfigServiceQt) {
	n = new(ProxyConfigServiceQt)
	n.InitFromInternal(uintptr(ptr), "webengine.ProxyConfigServiceQt")
	return
}

func (ptr *ProxyConfigServiceQt) DestroyProxyConfigServiceQt() {
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

func (n *QQuickWebEngineProfile) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickWebEngineProfile) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickWebEngineProfileFromPointer(ptr unsafe.Pointer) (n *QQuickWebEngineProfile) {
	n = new(QQuickWebEngineProfile)
	n.InitFromInternal(uintptr(ptr), "webengine.QQuickWebEngineProfile")
	return
}

func (ptr *QQuickWebEngineProfile) DestroyQQuickWebEngineProfile() {
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQQuickWebEngineProfile", "", parent}).(*QQuickWebEngineProfile)
}

func (ptr *QQuickWebEngineProfile) CachePath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CachePath"}).(string)
}

func (ptr *QQuickWebEngineProfile) ConnectCachePathChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCachePathChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectCachePathChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCachePathChanged"})
}

func (ptr *QQuickWebEngineProfile) CachePathChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CachePathChanged"})
}

func (ptr *QQuickWebEngineProfile) ClearHttpCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearHttpCache"})
}

func (ptr *QQuickWebEngineProfile) ClientCertificateStore() *QWebEngineClientCertificateStore {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClientCertificateStore"}).(*QWebEngineClientCertificateStore)
}

func (ptr *QQuickWebEngineProfile) CookieStore() *QWebEngineCookieStore {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookieStore"}).(*QWebEngineCookieStore)
}

func QQuickWebEngineProfile_DefaultProfile() *QQuickWebEngineProfile {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QQuickWebEngineProfile_DefaultProfile", ""}).(*QQuickWebEngineProfile)
}

func (ptr *QQuickWebEngineProfile) DefaultProfile() *QQuickWebEngineProfile {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QQuickWebEngineProfile_DefaultProfile", ""}).(*QQuickWebEngineProfile)
}

func (ptr *QQuickWebEngineProfile) DownloadPath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DownloadPath"}).(string)
}

func (ptr *QQuickWebEngineProfile) ConnectDownloadPathChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDownloadPathChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectDownloadPathChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDownloadPathChanged"})
}

func (ptr *QQuickWebEngineProfile) DownloadPathChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DownloadPathChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguage() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpAcceptLanguage"}).(string)
}

func (ptr *QQuickWebEngineProfile) ConnectHttpAcceptLanguageChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHttpAcceptLanguageChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpAcceptLanguageChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHttpAcceptLanguageChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpAcceptLanguageChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpAcceptLanguageChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSize() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpCacheMaximumSize"}).(float64))
}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheMaximumSizeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHttpCacheMaximumSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheMaximumSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHttpCacheMaximumSizeChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpCacheMaximumSizeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpCacheMaximumSizeChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpCacheType() QQuickWebEngineProfile__HttpCacheType {

	return QQuickWebEngineProfile__HttpCacheType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpCacheType"}).(float64))
}

func (ptr *QQuickWebEngineProfile) ConnectHttpCacheTypeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHttpCacheTypeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpCacheTypeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHttpCacheTypeChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpCacheTypeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpCacheTypeChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpUserAgent() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpUserAgent"}).(string)
}

func (ptr *QQuickWebEngineProfile) ConnectHttpUserAgentChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHttpUserAgentChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectHttpUserAgentChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHttpUserAgentChanged"})
}

func (ptr *QQuickWebEngineProfile) HttpUserAgentChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpUserAgentChanged"})
}

func (ptr *QQuickWebEngineProfile) InstallUrlSchemeHandler(scheme core.QByteArray_ITF, handler QWebEngineUrlSchemeHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstallUrlSchemeHandler", scheme, handler})
}

func (ptr *QQuickWebEngineProfile) IsOffTheRecord() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOffTheRecord"}).(bool)
}

func (ptr *QQuickWebEngineProfile) IsSpellCheckEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSpellCheckEnabled"}).(bool)
}

func (ptr *QQuickWebEngineProfile) IsUsedForGlobalCertificateVerification() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUsedForGlobalCertificateVerification"}).(bool)
}

func (ptr *QQuickWebEngineProfile) ConnectOffTheRecordChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOffTheRecordChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectOffTheRecordChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOffTheRecordChanged"})
}

func (ptr *QQuickWebEngineProfile) OffTheRecordChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OffTheRecordChanged"})
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicy() QQuickWebEngineProfile__PersistentCookiesPolicy {

	return QQuickWebEngineProfile__PersistentCookiesPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistentCookiesPolicy"}).(float64))
}

func (ptr *QQuickWebEngineProfile) ConnectPersistentCookiesPolicyChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPersistentCookiesPolicyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentCookiesPolicyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPersistentCookiesPolicyChanged"})
}

func (ptr *QQuickWebEngineProfile) PersistentCookiesPolicyChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistentCookiesPolicyChanged"})
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistentStoragePath"}).(string)
}

func (ptr *QQuickWebEngineProfile) ConnectPersistentStoragePathChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPersistentStoragePathChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectPersistentStoragePathChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPersistentStoragePathChanged"})
}

func (ptr *QQuickWebEngineProfile) PersistentStoragePathChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistentStoragePathChanged"})
}

func (ptr *QQuickWebEngineProfile) ConnectPresentNotification(f func(notification *QWebEngineNotification)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPresentNotification", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectPresentNotification() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPresentNotification"})
}

func (ptr *QQuickWebEngineProfile) PresentNotification(notification QWebEngineNotification_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PresentNotification", notification})
}

func (ptr *QQuickWebEngineProfile) RemoveAllUrlSchemeHandlers() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAllUrlSchemeHandlers"})
}

func (ptr *QQuickWebEngineProfile) RemoveUrlScheme(scheme core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveUrlScheme", scheme})
}

func (ptr *QQuickWebEngineProfile) RemoveUrlSchemeHandler(handler QWebEngineUrlSchemeHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveUrlSchemeHandler", handler})
}

func (ptr *QQuickWebEngineProfile) SetCachePath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCachePath", path})
}

func (ptr *QQuickWebEngineProfile) SetDownloadPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDownloadPath", path})
}

func (ptr *QQuickWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpAcceptLanguage", httpAcceptLanguage})
}

func (ptr *QQuickWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpCacheMaximumSize", maxSize})
}

func (ptr *QQuickWebEngineProfile) SetHttpCacheType(vqq QQuickWebEngineProfile__HttpCacheType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpCacheType", vqq})
}

func (ptr *QQuickWebEngineProfile) SetHttpUserAgent(userAgent string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpUserAgent", userAgent})
}

func (ptr *QQuickWebEngineProfile) SetOffTheRecord(offTheRecord bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOffTheRecord", offTheRecord})
}

func (ptr *QQuickWebEngineProfile) SetPersistentCookiesPolicy(vqq QQuickWebEngineProfile__PersistentCookiesPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistentCookiesPolicy", vqq})
}

func (ptr *QQuickWebEngineProfile) SetPersistentStoragePath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistentStoragePath", path})
}

func (ptr *QQuickWebEngineProfile) SetSpellCheckEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSpellCheckEnabled", enabled})
}

func (ptr *QQuickWebEngineProfile) SetSpellCheckLanguages(languages []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSpellCheckLanguages", languages})
}

func (ptr *QQuickWebEngineProfile) SetStorageName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStorageName", name})
}

func (ptr *QQuickWebEngineProfile) SetUrlRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrlRequestInterceptor", interceptor})
}

func (ptr *QQuickWebEngineProfile) SetUseForGlobalCertificateVerification(b bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUseForGlobalCertificateVerification", b})
}

func (ptr *QQuickWebEngineProfile) ConnectSpellCheckEnabledChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSpellCheckEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectSpellCheckEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSpellCheckEnabledChanged"})
}

func (ptr *QQuickWebEngineProfile) SpellCheckEnabledChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpellCheckEnabledChanged"})
}

func (ptr *QQuickWebEngineProfile) SpellCheckLanguages() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpellCheckLanguages"}).([]string)
}

func (ptr *QQuickWebEngineProfile) ConnectSpellCheckLanguagesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSpellCheckLanguagesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectSpellCheckLanguagesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSpellCheckLanguagesChanged"})
}

func (ptr *QQuickWebEngineProfile) SpellCheckLanguagesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpellCheckLanguagesChanged"})
}

func (ptr *QQuickWebEngineProfile) StorageName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StorageName"}).(string)
}

func (ptr *QQuickWebEngineProfile) ConnectStorageNameChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStorageNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectStorageNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStorageNameChanged"})
}

func (ptr *QQuickWebEngineProfile) StorageNameChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StorageNameChanged"})
}

func (ptr *QQuickWebEngineProfile) UrlSchemeHandler(scheme core.QByteArray_ITF) *QWebEngineUrlSchemeHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UrlSchemeHandler", scheme}).(*QWebEngineUrlSchemeHandler)
}

func (ptr *QQuickWebEngineProfile) ConnectUseForGlobalCertificateVerificationChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUseForGlobalCertificateVerificationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineProfile) DisconnectUseForGlobalCertificateVerificationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUseForGlobalCertificateVerificationChanged"})
}

func (ptr *QQuickWebEngineProfile) UseForGlobalCertificateVerificationChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseForGlobalCertificateVerificationChanged"})
}

func (ptr *QQuickWebEngineProfile) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickWebEngineProfile) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickWebEngineProfile) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineProfile) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickWebEngineProfile) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickWebEngineProfile) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineProfile) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickWebEngineProfile) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickWebEngineProfile) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineProfile) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickWebEngineProfile) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickWebEngineProfile) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickWebEngineProfile) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickWebEngineProfile) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QQuickWebEngineScript) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQuickWebEngineScript) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQuickWebEngineScriptFromPointer(ptr unsafe.Pointer) (n *QQuickWebEngineScript) {
	n = new(QQuickWebEngineScript)
	n.InitFromInternal(uintptr(ptr), "webengine.QQuickWebEngineScript")
	return
}

func (ptr *QQuickWebEngineScript) DestroyQQuickWebEngineScript() {
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQQuickWebEngineScript", "", parent}).(*QQuickWebEngineScript)
}

func (ptr *QQuickWebEngineScript) InjectionPoint() QQuickWebEngineScript__InjectionPoint {

	return QQuickWebEngineScript__InjectionPoint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InjectionPoint"}).(float64))
}

func (ptr *QQuickWebEngineScript) ConnectInjectionPointChanged(f func(injectionPoint QQuickWebEngineScript__InjectionPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInjectionPointChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineScript) DisconnectInjectionPointChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInjectionPointChanged"})
}

func (ptr *QQuickWebEngineScript) InjectionPointChanged(injectionPoint QQuickWebEngineScript__InjectionPoint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InjectionPointChanged", injectionPoint})
}

func (ptr *QQuickWebEngineScript) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QQuickWebEngineScript) ConnectNameChanged(f func(name string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineScript) DisconnectNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNameChanged"})
}

func (ptr *QQuickWebEngineScript) NameChanged(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameChanged", name})
}

func (ptr *QQuickWebEngineScript) RunOnSubframes() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RunOnSubframes"}).(bool)
}

func (ptr *QQuickWebEngineScript) ConnectRunOnSubframesChanged(f func(on bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRunOnSubframesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineScript) DisconnectRunOnSubframesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRunOnSubframesChanged"})
}

func (ptr *QQuickWebEngineScript) RunOnSubframesChanged(on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RunOnSubframesChanged", on})
}

func (ptr *QQuickWebEngineScript) SetInjectionPoint(injectionPoint QQuickWebEngineScript__InjectionPoint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInjectionPoint", injectionPoint})
}

func (ptr *QQuickWebEngineScript) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QQuickWebEngineScript) SetRunOnSubframes(on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRunOnSubframes", on})
}

func (ptr *QQuickWebEngineScript) SetSourceCode(code string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceCode", code})
}

func (ptr *QQuickWebEngineScript) SetSourceUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSourceUrl", url})
}

func (ptr *QQuickWebEngineScript) SetWorldId(scriptWorldId QQuickWebEngineScript__ScriptWorldId) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWorldId", scriptWorldId})
}

func (ptr *QQuickWebEngineScript) SourceCode() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceCode"}).(string)
}

func (ptr *QQuickWebEngineScript) ConnectSourceCodeChanged(f func(code string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSourceCodeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineScript) DisconnectSourceCodeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSourceCodeChanged"})
}

func (ptr *QQuickWebEngineScript) SourceCodeChanged(code string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceCodeChanged", code})
}

func (ptr *QQuickWebEngineScript) SourceUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceUrl"}).(*core.QUrl)
}

func (ptr *QQuickWebEngineScript) ConnectSourceUrlChanged(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSourceUrlChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineScript) DisconnectSourceUrlChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSourceUrlChanged"})
}

func (ptr *QQuickWebEngineScript) SourceUrlChanged(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceUrlChanged", url})
}

func (ptr *QQuickWebEngineScript) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QQuickWebEngineScript) WorldId() QQuickWebEngineScript__ScriptWorldId {

	return QQuickWebEngineScript__ScriptWorldId(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WorldId"}).(float64))
}

func (ptr *QQuickWebEngineScript) ConnectWorldIdChanged(f func(scriptWorldId QQuickWebEngineScript__ScriptWorldId)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWorldIdChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QQuickWebEngineScript) DisconnectWorldIdChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWorldIdChanged"})
}

func (ptr *QQuickWebEngineScript) WorldIdChanged(scriptWorldId QQuickWebEngineScript__ScriptWorldId) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WorldIdChanged", scriptWorldId})
}

func (ptr *QQuickWebEngineScript) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QQuickWebEngineScript) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QQuickWebEngineScript) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineScript) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QQuickWebEngineScript) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QQuickWebEngineScript) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineScript) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QQuickWebEngineScript) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QQuickWebEngineScript) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineScript) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QQuickWebEngineScript) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QQuickWebEngineScript) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QQuickWebEngineScript) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QQuickWebEngineScript) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QQuickWebEngineScript) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QQuickWebEngineScript) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QQuickWebEngineScript) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QQuickWebEngineScript) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QQuickWebEngineScript) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QQuickWebEngineScript) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QQuickWebEngineScript) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QWebEngineCallback struct {
	internal.Internal
}

type QWebEngineCallback_ITF interface {
	QWebEngineCallback_PTR() *QWebEngineCallback
}

func (ptr *QWebEngineCallback) QWebEngineCallback_PTR() *QWebEngineCallback {
	return ptr
}

func (ptr *QWebEngineCallback) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineCallback) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineCallback(ptr QWebEngineCallback_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineCallback_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineCallback) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineCallbackFromPointer(ptr unsafe.Pointer) (n *QWebEngineCallback) {
	n = new(QWebEngineCallback)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineCallback")
	return
}

func (ptr *QWebEngineCallback) DestroyQWebEngineCallback() {
}

type QWebEngineCertificateError struct {
	internal.Internal
}

type QWebEngineCertificateError_ITF interface {
	QWebEngineCertificateError_PTR() *QWebEngineCertificateError
}

func (ptr *QWebEngineCertificateError) QWebEngineCertificateError_PTR() *QWebEngineCertificateError {
	return ptr
}

func (ptr *QWebEngineCertificateError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineCertificateError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineCertificateError(ptr QWebEngineCertificateError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineCertificateError_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineCertificateError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineCertificateErrorFromPointer(ptr unsafe.Pointer) (n *QWebEngineCertificateError) {
	n = new(QWebEngineCertificateError)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineCertificateError")
	return
}

func (ptr *QWebEngineCertificateError) DestroyQWebEngineCertificateError() {
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

	return QWebEngineCertificateError__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QWebEngineCertificateError) ErrorDescription() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorDescription"}).(string)
}

func (ptr *QWebEngineCertificateError) IsOverridable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOverridable"}).(bool)
}

func (ptr *QWebEngineCertificateError) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

type QWebEngineClientCertificateSelection struct {
	internal.Internal
}

type QWebEngineClientCertificateSelection_ITF interface {
	QWebEngineClientCertificateSelection_PTR() *QWebEngineClientCertificateSelection
}

func (ptr *QWebEngineClientCertificateSelection) QWebEngineClientCertificateSelection_PTR() *QWebEngineClientCertificateSelection {
	return ptr
}

func (ptr *QWebEngineClientCertificateSelection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineClientCertificateSelection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineClientCertificateSelection(ptr QWebEngineClientCertificateSelection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineClientCertificateSelection_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineClientCertificateSelection) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineClientCertificateSelectionFromPointer(ptr unsafe.Pointer) (n *QWebEngineClientCertificateSelection) {
	n = new(QWebEngineClientCertificateSelection)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineClientCertificateSelection")
	return
}

func (ptr *QWebEngineClientCertificateSelection) DestroyQWebEngineClientCertificateSelection() {
}

func NewQWebEngineClientCertificateSelection(vqw QWebEngineClientCertificateSelection_ITF) *QWebEngineClientCertificateSelection {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineClientCertificateSelection", "", vqw}).(*QWebEngineClientCertificateSelection)
}

func (ptr *QWebEngineClientCertificateSelection) Certificates() []*network.QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Certificates"}).([]*network.QSslCertificate)
}

func (ptr *QWebEngineClientCertificateSelection) Host() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Host"}).(*core.QUrl)
}

func (ptr *QWebEngineClientCertificateSelection) Select(certificate network.QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Select", certificate})
}

func (ptr *QWebEngineClientCertificateSelection) SelectNone() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectNone"})
}

func (ptr *QWebEngineClientCertificateSelection) __certificates_atList(i int) *network.QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__certificates_atList", i}).(*network.QSslCertificate)
}

func (ptr *QWebEngineClientCertificateSelection) __certificates_setList(i network.QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__certificates_setList", i})
}

func (ptr *QWebEngineClientCertificateSelection) __certificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__certificates_newList"}).(unsafe.Pointer)
}

type QWebEngineClientCertificateStore struct {
	internal.Internal
}

type QWebEngineClientCertificateStore_ITF interface {
	QWebEngineClientCertificateStore_PTR() *QWebEngineClientCertificateStore
}

func (ptr *QWebEngineClientCertificateStore) QWebEngineClientCertificateStore_PTR() *QWebEngineClientCertificateStore {
	return ptr
}

func (ptr *QWebEngineClientCertificateStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineClientCertificateStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineClientCertificateStore(ptr QWebEngineClientCertificateStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineClientCertificateStore_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineClientCertificateStore) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineClientCertificateStoreFromPointer(ptr unsafe.Pointer) (n *QWebEngineClientCertificateStore) {
	n = new(QWebEngineClientCertificateStore)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineClientCertificateStore")
	return
}

func (ptr *QWebEngineClientCertificateStore) DestroyQWebEngineClientCertificateStore() {
}

func (ptr *QWebEngineClientCertificateStore) Add(certificate network.QSslCertificate_ITF, privateKey network.QSslKey_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Add", certificate, privateKey})
}

func (ptr *QWebEngineClientCertificateStore) Certificates() []*network.QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Certificates"}).([]*network.QSslCertificate)
}

func (ptr *QWebEngineClientCertificateStore) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QWebEngineClientCertificateStore) Remove(certificate network.QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", certificate})
}

func (ptr *QWebEngineClientCertificateStore) __certificates_atList(i int) *network.QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__certificates_atList", i}).(*network.QSslCertificate)
}

func (ptr *QWebEngineClientCertificateStore) __certificates_setList(i network.QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__certificates_setList", i})
}

func (ptr *QWebEngineClientCertificateStore) __certificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__certificates_newList"}).(unsafe.Pointer)
}

type QWebEngineContextMenuData struct {
	internal.Internal
}

type QWebEngineContextMenuData_ITF interface {
	QWebEngineContextMenuData_PTR() *QWebEngineContextMenuData
}

func (ptr *QWebEngineContextMenuData) QWebEngineContextMenuData_PTR() *QWebEngineContextMenuData {
	return ptr
}

func (ptr *QWebEngineContextMenuData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineContextMenuData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineContextMenuData(ptr QWebEngineContextMenuData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineContextMenuData_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineContextMenuData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineContextMenuDataFromPointer(ptr unsafe.Pointer) (n *QWebEngineContextMenuData) {
	n = new(QWebEngineContextMenuData)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineContextMenuData")
	return
}

func (ptr *QWebEngineContextMenuData) DestroyQWebEngineContextMenuData() {
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

func (n *QWebEngineCookieStore) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineCookieStore) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEngineCookieStoreFromPointer(ptr unsafe.Pointer) (n *QWebEngineCookieStore) {
	n = new(QWebEngineCookieStore)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineCookieStore")
	return
}
func (ptr *QWebEngineCookieStore) ConnectCookieAdded(f func(cookie *network.QNetworkCookie)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCookieAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineCookieStore) DisconnectCookieAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCookieAdded"})
}

func (ptr *QWebEngineCookieStore) CookieAdded(cookie network.QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookieAdded", cookie})
}

func (ptr *QWebEngineCookieStore) ConnectCookieRemoved(f func(cookie *network.QNetworkCookie)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCookieRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineCookieStore) DisconnectCookieRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCookieRemoved"})
}

func (ptr *QWebEngineCookieStore) CookieRemoved(cookie network.QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookieRemoved", cookie})
}

func (ptr *QWebEngineCookieStore) DeleteAllCookies() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteAllCookies"})
}

func (ptr *QWebEngineCookieStore) DeleteCookie(cookie network.QNetworkCookie_ITF, origin core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteCookie", cookie, origin})
}

func (ptr *QWebEngineCookieStore) DeleteSessionCookies() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteSessionCookies"})
}

func (ptr *QWebEngineCookieStore) LoadAllCookies() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadAllCookies"})
}

func (ptr *QWebEngineCookieStore) SetCookie(cookie network.QNetworkCookie_ITF, origin core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCookie", cookie, origin})
}

func (ptr *QWebEngineCookieStore) ConnectDestroyQWebEngineCookieStore(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebEngineCookieStore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineCookieStore) DisconnectDestroyQWebEngineCookieStore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebEngineCookieStore"})
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStore() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineCookieStore"})
}

func (ptr *QWebEngineCookieStore) DestroyQWebEngineCookieStoreDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineCookieStoreDefault"})
}

func (ptr *QWebEngineCookieStore) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineCookieStore) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineCookieStore) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineCookieStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineCookieStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineCookieStore) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineCookieStore) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineCookieStore) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineCookieStore) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineCookieStore) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineCookieStore) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineCookieStore) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineCookieStore) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineCookieStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineCookieStore) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineCookieStore) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineCookieStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineCookieStore) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebEngineCookieStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineCookieStore) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineCookieStore) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QWebEngineHttpRequest struct {
	internal.Internal
}

type QWebEngineHttpRequest_ITF interface {
	QWebEngineHttpRequest_PTR() *QWebEngineHttpRequest
}

func (ptr *QWebEngineHttpRequest) QWebEngineHttpRequest_PTR() *QWebEngineHttpRequest {
	return ptr
}

func (ptr *QWebEngineHttpRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineHttpRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineHttpRequest(ptr QWebEngineHttpRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineHttpRequest_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineHttpRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineHttpRequestFromPointer(ptr unsafe.Pointer) (n *QWebEngineHttpRequest) {
	n = new(QWebEngineHttpRequest)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineHttpRequest")
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineHttpRequest", "", url, method}).(*QWebEngineHttpRequest)
}

func NewQWebEngineHttpRequest2(other QWebEngineHttpRequest_ITF) *QWebEngineHttpRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineHttpRequest2", "", other}).(*QWebEngineHttpRequest)
}

func (ptr *QWebEngineHttpRequest) HasHeader(headerName core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeader", headerName}).(bool)
}

func (ptr *QWebEngineHttpRequest) Header(headerName core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Header", headerName}).(*core.QByteArray)
}

func (ptr *QWebEngineHttpRequest) Headers() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Headers"}).([]*core.QByteArray)
}

func (ptr *QWebEngineHttpRequest) Method() QWebEngineHttpRequest__Method {

	return QWebEngineHttpRequest__Method(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Method"}).(float64))
}

func (ptr *QWebEngineHttpRequest) PostData() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PostData"}).(*core.QByteArray)
}

func QWebEngineHttpRequest_PostRequest(url core.QUrl_ITF, postData map[string]string) *QWebEngineHttpRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineHttpRequest_PostRequest", "", url, postData}).(*QWebEngineHttpRequest)
}

func (ptr *QWebEngineHttpRequest) PostRequest(url core.QUrl_ITF, postData map[string]string) *QWebEngineHttpRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineHttpRequest_PostRequest", "", url, postData}).(*QWebEngineHttpRequest)
}

func (ptr *QWebEngineHttpRequest) SetHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeader", headerName, headerValue})
}

func (ptr *QWebEngineHttpRequest) SetMethod(method QWebEngineHttpRequest__Method) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMethod", method})
}

func (ptr *QWebEngineHttpRequest) SetPostData(postData core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPostData", postData})
}

func (ptr *QWebEngineHttpRequest) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QWebEngineHttpRequest) Swap(other QWebEngineHttpRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QWebEngineHttpRequest) UnsetHeader(key core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnsetHeader", key})
}

func (ptr *QWebEngineHttpRequest) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QWebEngineHttpRequest) DestroyQWebEngineHttpRequest() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineHttpRequest"})
}

func (ptr *QWebEngineHttpRequest) __headers_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__headers_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineHttpRequest) __headers_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__headers_setList", i})
}

func (ptr *QWebEngineHttpRequest) __headers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__headers_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_atList(v string, i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__postRequest_postData_atList", v, i}).(string)
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_setList(key string, i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__postRequest_postData_setList", key, i})
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__postRequest_postData_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineHttpRequest) __postRequest_postData_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__postRequest_postData_keyList"}).([]string)
}

func (ptr *QWebEngineHttpRequest) ____postRequest_postData_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____postRequest_postData_keyList_atList", i}).(string)
}

func (ptr *QWebEngineHttpRequest) ____postRequest_postData_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____postRequest_postData_keyList_setList", i})
}

func (ptr *QWebEngineHttpRequest) ____postRequest_postData_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____postRequest_postData_keyList_newList"}).(unsafe.Pointer)
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

func (n *QWebEngineNotification) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineNotification) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEngineNotificationFromPointer(ptr unsafe.Pointer) (n *QWebEngineNotification) {
	n = new(QWebEngineNotification)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineNotification")
	return
}

func (ptr *QWebEngineNotification) DestroyQWebEngineNotification() {
}

func (ptr *QWebEngineNotification) ConnectClick(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClick", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineNotification) DisconnectClick() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClick"})
}

func (ptr *QWebEngineNotification) Click() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Click"})
}

func (ptr *QWebEngineNotification) ClickDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClickDefault"})
}

func (ptr *QWebEngineNotification) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineNotification) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QWebEngineNotification) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QWebEngineNotification) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QWebEngineNotification) ConnectClosed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClosed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineNotification) DisconnectClosed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClosed"})
}

func (ptr *QWebEngineNotification) Closed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Closed"})
}

func (ptr *QWebEngineNotification) Direction() core.Qt__LayoutDirection {

	return core.Qt__LayoutDirection(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Direction"}).(float64))
}

func (ptr *QWebEngineNotification) Icon() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Icon"}).(*gui.QImage)
}

func (ptr *QWebEngineNotification) Language() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Language"}).(string)
}

func (ptr *QWebEngineNotification) Matches(other QWebEngineNotification_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Matches", other}).(bool)
}

func (ptr *QWebEngineNotification) Message() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Message"}).(string)
}

func (ptr *QWebEngineNotification) Origin() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Origin"}).(*core.QUrl)
}

func (ptr *QWebEngineNotification) ConnectShow(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineNotification) DisconnectShow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShow"})
}

func (ptr *QWebEngineNotification) Show() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Show"})
}

func (ptr *QWebEngineNotification) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QWebEngineNotification) Tag() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Tag"}).(string)
}

func (ptr *QWebEngineNotification) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QWebEngineNotification) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineNotification) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineNotification) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineNotification) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineNotification) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineNotification) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineNotification) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineNotification) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineNotification) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineNotification) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineNotification) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineNotification) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineNotification) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineNotification) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineNotification) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineNotification) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineNotification) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineNotification) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebEngineNotification) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineNotification) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineNotification) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QWebEnginePage) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEnginePage) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEnginePageFromPointer(ptr unsafe.Pointer) (n *QWebEnginePage) {
	n = new(QWebEnginePage)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEnginePage")
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEnginePage", "", parent}).(*QWebEnginePage)
}

func NewQWebEnginePage2(profile QWebEngineProfile_ITF, parent core.QObject_ITF) *QWebEnginePage {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEnginePage2", "", profile, parent}).(*QWebEnginePage)
}

func (ptr *QWebEnginePage) ConnectAcceptNavigationRequest(f func(url *core.QUrl, ty QWebEnginePage__NavigationType, isMainFrame bool) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAcceptNavigationRequest", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectAcceptNavigationRequest() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAcceptNavigationRequest"})
}

func (ptr *QWebEnginePage) AcceptNavigationRequest(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptNavigationRequest", url, ty, isMainFrame}).(bool)
}

func (ptr *QWebEnginePage) AcceptNavigationRequestDefault(url core.QUrl_ITF, ty QWebEnginePage__NavigationType, isMainFrame bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptNavigationRequestDefault", url, ty, isMainFrame}).(bool)
}

func (ptr *QWebEnginePage) Action(action QWebEnginePage__WebAction) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Action", action}).(*widgets.QAction)
}

func (ptr *QWebEnginePage) ConnectAudioMutedChanged(f func(muted bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAudioMutedChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectAudioMutedChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAudioMutedChanged"})
}

func (ptr *QWebEnginePage) AudioMutedChanged(muted bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AudioMutedChanged", muted})
}

func (ptr *QWebEnginePage) ConnectAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAuthenticationRequired"})
}

func (ptr *QWebEnginePage) AuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AuthenticationRequired", requestUrl, authenticator})
}

func (ptr *QWebEnginePage) BackgroundColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundColor"}).(*gui.QColor)
}

func (ptr *QWebEnginePage) ConnectCertificateError(f func(certificateError *QWebEngineCertificateError) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCertificateError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectCertificateError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCertificateError"})
}

func (ptr *QWebEnginePage) CertificateError(certificateError QWebEngineCertificateError_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CertificateError", certificateError}).(bool)
}

func (ptr *QWebEnginePage) CertificateErrorDefault(certificateError QWebEngineCertificateError_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CertificateErrorDefault", certificateError}).(bool)
}

func (ptr *QWebEnginePage) ConnectChooseFiles(f func(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectChooseFiles", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectChooseFiles() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectChooseFiles"})
}

func (ptr *QWebEnginePage) ChooseFiles(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChooseFiles", mode, oldFiles, acceptedMimeTypes}).([]string)
}

func (ptr *QWebEnginePage) ChooseFilesDefault(mode QWebEnginePage__FileSelectionMode, oldFiles []string, acceptedMimeTypes []string) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChooseFilesDefault", mode, oldFiles, acceptedMimeTypes}).([]string)
}

func (ptr *QWebEnginePage) ContentsSize() *core.QSizeF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsSize"}).(*core.QSizeF)
}

func (ptr *QWebEnginePage) ConnectContentsSizeChanged(f func(size *core.QSizeF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentsSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectContentsSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentsSizeChanged"})
}

func (ptr *QWebEnginePage) ContentsSizeChanged(size core.QSizeF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentsSizeChanged", size})
}

func (ptr *QWebEnginePage) ContextMenuData() *QWebEngineContextMenuData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuData"}).(*QWebEngineContextMenuData)
}

func (ptr *QWebEnginePage) CreateStandardContextMenu() *widgets.QMenu {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateStandardContextMenu"}).(*widgets.QMenu)
}

func (ptr *QWebEnginePage) ConnectCreateWindow(f func(ty QWebEnginePage__WebWindowType) *QWebEnginePage) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectCreateWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateWindow"})
}

func (ptr *QWebEnginePage) CreateWindow(ty QWebEnginePage__WebWindowType) *QWebEnginePage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateWindow", ty}).(*QWebEnginePage)
}

func (ptr *QWebEnginePage) CreateWindowDefault(ty QWebEnginePage__WebWindowType) *QWebEnginePage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateWindowDefault", ty}).(*QWebEnginePage)
}

func (ptr *QWebEnginePage) DevToolsPage() *QWebEnginePage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DevToolsPage"}).(*QWebEnginePage)
}

func (ptr *QWebEnginePage) Download(url core.QUrl_ITF, filename string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Download", url, filename})
}

func (ptr *QWebEnginePage) EventDefault(vqe core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", vqe}).(bool)
}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequestCanceled(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFeaturePermissionRequestCanceled", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequestCanceled() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFeaturePermissionRequestCanceled"})
}

func (ptr *QWebEnginePage) FeaturePermissionRequestCanceled(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FeaturePermissionRequestCanceled", securityOrigin, feature})
}

func (ptr *QWebEnginePage) ConnectFeaturePermissionRequested(f func(securityOrigin *core.QUrl, feature QWebEnginePage__Feature)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFeaturePermissionRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectFeaturePermissionRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFeaturePermissionRequested"})
}

func (ptr *QWebEnginePage) FeaturePermissionRequested(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FeaturePermissionRequested", securityOrigin, feature})
}

func (ptr *QWebEnginePage) ConnectGeometryChangeRequested(f func(geom *core.QRect)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGeometryChangeRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectGeometryChangeRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGeometryChangeRequested"})
}

func (ptr *QWebEnginePage) GeometryChangeRequested(geom core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GeometryChangeRequested", geom})
}

func (ptr *QWebEnginePage) HasSelection() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasSelection"}).(bool)
}

func (ptr *QWebEnginePage) Icon() *gui.QIcon {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Icon"}).(*gui.QIcon)
}

func (ptr *QWebEnginePage) ConnectIconChanged(f func(icon *gui.QIcon)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIconChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectIconChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIconChanged"})
}

func (ptr *QWebEnginePage) IconChanged(icon gui.QIcon_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconChanged", icon})
}

func (ptr *QWebEnginePage) IconUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconUrl"}).(*core.QUrl)
}

func (ptr *QWebEnginePage) ConnectIconUrlChanged(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIconUrlChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectIconUrlChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIconUrlChanged"})
}

func (ptr *QWebEnginePage) IconUrlChanged(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconUrlChanged", url})
}

func (ptr *QWebEnginePage) InspectedPage() *QWebEnginePage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InspectedPage"}).(*QWebEnginePage)
}

func (ptr *QWebEnginePage) IsAudioMuted() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAudioMuted"}).(bool)
}

func (ptr *QWebEnginePage) ConnectJavaScriptAlert(f func(securityOrigin *core.QUrl, msg string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectJavaScriptAlert", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectJavaScriptAlert() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectJavaScriptAlert"})
}

func (ptr *QWebEnginePage) JavaScriptAlert(securityOrigin core.QUrl_ITF, msg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptAlert", securityOrigin, msg})
}

func (ptr *QWebEnginePage) JavaScriptAlertDefault(securityOrigin core.QUrl_ITF, msg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptAlertDefault", securityOrigin, msg})
}

func (ptr *QWebEnginePage) ConnectJavaScriptConfirm(f func(securityOrigin *core.QUrl, msg string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectJavaScriptConfirm", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConfirm() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectJavaScriptConfirm"})
}

func (ptr *QWebEnginePage) JavaScriptConfirm(securityOrigin core.QUrl_ITF, msg string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptConfirm", securityOrigin, msg}).(bool)
}

func (ptr *QWebEnginePage) JavaScriptConfirmDefault(securityOrigin core.QUrl_ITF, msg string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptConfirmDefault", securityOrigin, msg}).(bool)
}

func (ptr *QWebEnginePage) ConnectJavaScriptConsoleMessage(f func(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectJavaScriptConsoleMessage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectJavaScriptConsoleMessage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectJavaScriptConsoleMessage"})
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessage(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptConsoleMessage", level, message, lineNumber, sourceID})
}

func (ptr *QWebEnginePage) JavaScriptConsoleMessageDefault(level QWebEnginePage__JavaScriptConsoleMessageLevel, message string, lineNumber int, sourceID string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptConsoleMessageDefault", level, message, lineNumber, sourceID})
}

func (ptr *QWebEnginePage) ConnectJavaScriptPrompt(f func(securityOrigin *core.QUrl, msg string, defaultValue string, result string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectJavaScriptPrompt", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectJavaScriptPrompt() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectJavaScriptPrompt"})
}

func (ptr *QWebEnginePage) JavaScriptPrompt(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptPrompt", securityOrigin, msg, defaultValue, result}).(bool)
}

func (ptr *QWebEnginePage) JavaScriptPromptDefault(securityOrigin core.QUrl_ITF, msg string, defaultValue string, result string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JavaScriptPromptDefault", securityOrigin, msg, defaultValue, result}).(bool)
}

func (ptr *QWebEnginePage) ConnectLinkHovered(f func(url string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLinkHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectLinkHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLinkHovered"})
}

func (ptr *QWebEnginePage) LinkHovered(url string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinkHovered", url})
}

func (ptr *QWebEnginePage) Load(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", url})
}

func (ptr *QWebEnginePage) Load2(request QWebEngineHttpRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2", request})
}

func (ptr *QWebEnginePage) ConnectLoadFinished(f func(ok bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectLoadFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadFinished"})
}

func (ptr *QWebEnginePage) LoadFinished(ok bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadFinished", ok})
}

func (ptr *QWebEnginePage) ConnectLoadProgress(f func(progress int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadProgress", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectLoadProgress() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadProgress"})
}

func (ptr *QWebEnginePage) LoadProgress(progress int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadProgress", progress})
}

func (ptr *QWebEnginePage) ConnectLoadStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectLoadStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadStarted"})
}

func (ptr *QWebEnginePage) LoadStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadStarted"})
}

func (ptr *QWebEnginePage) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEnginePage) ConnectPdfPrintingFinished(f func(filePath string, success bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPdfPrintingFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectPdfPrintingFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPdfPrintingFinished"})
}

func (ptr *QWebEnginePage) PdfPrintingFinished(filePath string, success bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PdfPrintingFinished", filePath, success})
}

func (ptr *QWebEnginePage) Print(printer printsupport.QPrinter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Print", printer})
}

func (ptr *QWebEnginePage) ConnectPrintRequested(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrintRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectPrintRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrintRequested"})
}

func (ptr *QWebEnginePage) PrintRequested() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintRequested"})
}

func (ptr *QWebEnginePage) PrintToPdf(filePath string, layout gui.QPageLayout_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintToPdf", filePath, layout})
}

func (ptr *QWebEnginePage) Profile() *QWebEngineProfile {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Profile"}).(*QWebEngineProfile)
}

func (ptr *QWebEnginePage) ConnectProxyAuthenticationRequired(f func(requestUrl *core.QUrl, authenticator *network.QAuthenticator, proxyHost string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProxyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectProxyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProxyAuthenticationRequired"})
}

func (ptr *QWebEnginePage) ProxyAuthenticationRequired(requestUrl core.QUrl_ITF, authenticator network.QAuthenticator_ITF, proxyHost string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProxyAuthenticationRequired", requestUrl, authenticator, proxyHost})
}

func (ptr *QWebEnginePage) RecentlyAudible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecentlyAudible"}).(bool)
}

func (ptr *QWebEnginePage) ConnectRecentlyAudibleChanged(f func(recentlyAudible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRecentlyAudibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectRecentlyAudibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRecentlyAudibleChanged"})
}

func (ptr *QWebEnginePage) RecentlyAudibleChanged(recentlyAudible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecentlyAudibleChanged", recentlyAudible})
}

func (ptr *QWebEnginePage) DisconnectRenderProcessTerminated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRenderProcessTerminated"})
}

func (ptr *QWebEnginePage) ReplaceMisspelledWord(replacement string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReplaceMisspelledWord", replacement})
}

func (ptr *QWebEnginePage) RequestedUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestedUrl"}).(*core.QUrl)
}

func (ptr *QWebEnginePage) RunJavaScript(scriptSource string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RunJavaScript", scriptSource})
}

func (ptr *QWebEnginePage) RunJavaScript2(scriptSource string, worldId uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RunJavaScript2", scriptSource, worldId})
}

func (ptr *QWebEnginePage) ScrollPosition() *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollPosition"}).(*core.QPointF)
}

func (ptr *QWebEnginePage) ConnectScrollPositionChanged(f func(position *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScrollPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectScrollPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScrollPositionChanged"})
}

func (ptr *QWebEnginePage) ScrollPositionChanged(position core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollPositionChanged", position})
}

func (ptr *QWebEnginePage) ConnectSelectClientCertificate(f func(clientCertSelection *QWebEngineClientCertificateSelection)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectClientCertificate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectSelectClientCertificate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectClientCertificate"})
}

func (ptr *QWebEnginePage) SelectClientCertificate(clientCertSelection QWebEngineClientCertificateSelection_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectClientCertificate", clientCertSelection})
}

func (ptr *QWebEnginePage) SelectedText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedText"}).(string)
}

func (ptr *QWebEnginePage) ConnectSelectionChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectSelectionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionChanged"})
}

func (ptr *QWebEnginePage) SelectionChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionChanged"})
}

func (ptr *QWebEnginePage) SetAudioMuted(muted bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAudioMuted", muted})
}

func (ptr *QWebEnginePage) SetBackgroundColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundColor", color})
}

func (ptr *QWebEnginePage) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent", data, mimeType, baseUrl})
}

func (ptr *QWebEnginePage) SetDevToolsPage(page QWebEnginePage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDevToolsPage", page})
}

func (ptr *QWebEnginePage) SetFeaturePermission(securityOrigin core.QUrl_ITF, feature QWebEnginePage__Feature, policy QWebEnginePage__PermissionPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFeaturePermission", securityOrigin, feature, policy})
}

func (ptr *QWebEnginePage) SetHtml(html string, baseUrl core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHtml", html, baseUrl})
}

func (ptr *QWebEnginePage) SetInspectedPage(page QWebEnginePage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInspectedPage", page})
}

func (ptr *QWebEnginePage) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QWebEnginePage) SetUrlRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrlRequestInterceptor", interceptor})
}

func (ptr *QWebEnginePage) SetView(view widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetView", view})
}

func (ptr *QWebEnginePage) SetWebChannel(vqw webchannel.QWebChannel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWebChannel", vqw})
}

func (ptr *QWebEnginePage) SetWebChannel2(vqw webchannel.QWebChannel_ITF, worldId uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWebChannel2", vqw, worldId})
}

func (ptr *QWebEnginePage) SetZoomFactor(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomFactor", factor})
}

func (ptr *QWebEnginePage) Settings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Settings"}).(*QWebEngineSettings)
}

func (ptr *QWebEnginePage) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QWebEnginePage) ConnectTitleChanged(f func(title string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectTitleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleChanged"})
}

func (ptr *QWebEnginePage) TitleChanged(title string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleChanged", title})
}

func QWebEnginePage_Tr(s string, c string, n int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEnginePage_Tr", "", s, c, n}).(string)
}

func (ptr *QWebEnginePage) Tr(s string, c string, n int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEnginePage_Tr", "", s, c, n}).(string)
}

func (ptr *QWebEnginePage) ConnectTriggerAction(f func(action QWebEnginePage__WebAction, checked bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTriggerAction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectTriggerAction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTriggerAction"})
}

func (ptr *QWebEnginePage) TriggerAction(action QWebEnginePage__WebAction, checked bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TriggerAction", action, checked})
}

func (ptr *QWebEnginePage) TriggerActionDefault(action QWebEnginePage__WebAction, checked bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TriggerActionDefault", action, checked})
}

func (ptr *QWebEnginePage) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QWebEnginePage) ConnectUrlChanged(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUrlChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectUrlChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUrlChanged"})
}

func (ptr *QWebEnginePage) UrlChanged(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UrlChanged", url})
}

func (ptr *QWebEnginePage) View() *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "View"}).(*widgets.QWidget)
}

func (ptr *QWebEnginePage) WebChannel() *webchannel.QWebChannel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WebChannel"}).(*webchannel.QWebChannel)
}

func (ptr *QWebEnginePage) ConnectWindowCloseRequested(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWindowCloseRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectWindowCloseRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWindowCloseRequested"})
}

func (ptr *QWebEnginePage) WindowCloseRequested() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowCloseRequested"})
}

func (ptr *QWebEnginePage) ZoomFactor() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomFactor"}).(float64)
}

func (ptr *QWebEnginePage) ConnectDestroyQWebEnginePage(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebEnginePage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEnginePage) DisconnectDestroyQWebEnginePage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebEnginePage"})
}

func (ptr *QWebEnginePage) DestroyQWebEnginePage() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEnginePage"})
}

func (ptr *QWebEnginePage) DestroyQWebEnginePageDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEnginePageDefault"})
}

func (ptr *QWebEnginePage) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEnginePage) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEnginePage) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEnginePage) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEnginePage) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEnginePage) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEnginePage) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEnginePage) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEnginePage) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEnginePage) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEnginePage) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEnginePage) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEnginePage) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEnginePage) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEnginePage) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEnginePage) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEnginePage) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEnginePage) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEnginePage) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QWebEngineProfile) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineProfile) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEngineProfileFromPointer(ptr unsafe.Pointer) (n *QWebEngineProfile) {
	n = new(QWebEngineProfile)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineProfile")
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineProfile", "", parent}).(*QWebEngineProfile)
}

func NewQWebEngineProfile2(name string, parent core.QObject_ITF) *QWebEngineProfile {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineProfile2", "", name, parent}).(*QWebEngineProfile)
}

func (ptr *QWebEngineProfile) CachePath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CachePath"}).(string)
}

func (ptr *QWebEngineProfile) ClearAllVisitedLinks() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearAllVisitedLinks"})
}

func (ptr *QWebEngineProfile) ClearHttpCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearHttpCache"})
}

func (ptr *QWebEngineProfile) ClearVisitedLinks(urls []*core.QUrl) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearVisitedLinks", urls})
}

func (ptr *QWebEngineProfile) ClientCertificateStore() *QWebEngineClientCertificateStore {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClientCertificateStore"}).(*QWebEngineClientCertificateStore)
}

func (ptr *QWebEngineProfile) CookieStore() *QWebEngineCookieStore {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookieStore"}).(*QWebEngineCookieStore)
}

func QWebEngineProfile_DefaultProfile() *QWebEngineProfile {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineProfile_DefaultProfile", ""}).(*QWebEngineProfile)
}

func (ptr *QWebEngineProfile) DefaultProfile() *QWebEngineProfile {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineProfile_DefaultProfile", ""}).(*QWebEngineProfile)
}

func (ptr *QWebEngineProfile) DownloadPath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DownloadPath"}).(string)
}

func (ptr *QWebEngineProfile) DisconnectDownloadRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDownloadRequested"})
}

func (ptr *QWebEngineProfile) HttpAcceptLanguage() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpAcceptLanguage"}).(string)
}

func (ptr *QWebEngineProfile) HttpCacheMaximumSize() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpCacheMaximumSize"}).(float64))
}

func (ptr *QWebEngineProfile) HttpCacheType() QWebEngineProfile__HttpCacheType {

	return QWebEngineProfile__HttpCacheType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpCacheType"}).(float64))
}

func (ptr *QWebEngineProfile) HttpUserAgent() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HttpUserAgent"}).(string)
}

func (ptr *QWebEngineProfile) InstallUrlSchemeHandler(scheme core.QByteArray_ITF, vqw QWebEngineUrlSchemeHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstallUrlSchemeHandler", scheme, vqw})
}

func (ptr *QWebEngineProfile) IsOffTheRecord() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOffTheRecord"}).(bool)
}

func (ptr *QWebEngineProfile) IsSpellCheckEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSpellCheckEnabled"}).(bool)
}

func (ptr *QWebEngineProfile) IsUsedForGlobalCertificateVerification() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUsedForGlobalCertificateVerification"}).(bool)
}

func (ptr *QWebEngineProfile) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineProfile) PersistentCookiesPolicy() QWebEngineProfile__PersistentCookiesPolicy {

	return QWebEngineProfile__PersistentCookiesPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistentCookiesPolicy"}).(float64))
}

func (ptr *QWebEngineProfile) PersistentStoragePath() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PersistentStoragePath"}).(string)
}

func (ptr *QWebEngineProfile) RemoveAllUrlSchemeHandlers() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAllUrlSchemeHandlers"})
}

func (ptr *QWebEngineProfile) RemoveUrlScheme(scheme core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveUrlScheme", scheme})
}

func (ptr *QWebEngineProfile) RemoveUrlSchemeHandler(vqw QWebEngineUrlSchemeHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveUrlSchemeHandler", vqw})
}

func (ptr *QWebEngineProfile) Scripts() *QWebEngineScriptCollection {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Scripts"}).(*QWebEngineScriptCollection)
}

func (ptr *QWebEngineProfile) SetCachePath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCachePath", path})
}

func (ptr *QWebEngineProfile) SetDownloadPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDownloadPath", path})
}

func (ptr *QWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpAcceptLanguage", httpAcceptLanguage})
}

func (ptr *QWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpCacheMaximumSize", maxSize})
}

func (ptr *QWebEngineProfile) SetHttpCacheType(vqw QWebEngineProfile__HttpCacheType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpCacheType", vqw})
}

func (ptr *QWebEngineProfile) SetHttpUserAgent(userAgent string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpUserAgent", userAgent})
}

func (ptr *QWebEngineProfile) SetPersistentCookiesPolicy(vqw QWebEngineProfile__PersistentCookiesPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistentCookiesPolicy", vqw})
}

func (ptr *QWebEngineProfile) SetPersistentStoragePath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistentStoragePath", path})
}

func (ptr *QWebEngineProfile) SetRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequestInterceptor", interceptor})
}

func (ptr *QWebEngineProfile) SetSpellCheckEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSpellCheckEnabled", enabled})
}

func (ptr *QWebEngineProfile) SetSpellCheckLanguages(languages []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSpellCheckLanguages", languages})
}

func (ptr *QWebEngineProfile) SetUrlRequestInterceptor(interceptor QWebEngineUrlRequestInterceptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrlRequestInterceptor", interceptor})
}

func (ptr *QWebEngineProfile) SetUseForGlobalCertificateVerification(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUseForGlobalCertificateVerification", enabled})
}

func (ptr *QWebEngineProfile) Settings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Settings"}).(*QWebEngineSettings)
}

func (ptr *QWebEngineProfile) SpellCheckLanguages() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpellCheckLanguages"}).([]string)
}

func (ptr *QWebEngineProfile) StorageName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StorageName"}).(string)
}

func QWebEngineProfile_Tr(s string, c string, n int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineProfile_Tr", "", s, c, n}).(string)
}

func (ptr *QWebEngineProfile) Tr(s string, c string, n int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineProfile_Tr", "", s, c, n}).(string)
}

func (ptr *QWebEngineProfile) UrlSchemeHandler(vqb core.QByteArray_ITF) *QWebEngineUrlSchemeHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UrlSchemeHandler", vqb}).(*QWebEngineUrlSchemeHandler)
}

func (ptr *QWebEngineProfile) VisitedLinksContainsUrl(url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisitedLinksContainsUrl", url}).(bool)
}

func (ptr *QWebEngineProfile) ConnectDestroyQWebEngineProfile(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebEngineProfile", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineProfile) DisconnectDestroyQWebEngineProfile() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebEngineProfile"})
}

func (ptr *QWebEngineProfile) DestroyQWebEngineProfile() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineProfile"})
}

func (ptr *QWebEngineProfile) DestroyQWebEngineProfileDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineProfileDefault"})
}

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_atList(i int) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__clearVisitedLinks_urls_atList", i}).(*core.QUrl)
}

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_setList(i core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__clearVisitedLinks_urls_setList", i})
}

func (ptr *QWebEngineProfile) __clearVisitedLinks_urls_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__clearVisitedLinks_urls_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineProfile) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineProfile) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineProfile) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineProfile) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineProfile) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineProfile) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineProfile) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineProfile) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineProfile) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineProfile) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineProfile) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineProfile) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineProfile) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineProfile) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineProfile) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineProfile) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineProfile) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineProfile) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebEngineProfile) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineProfile) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QWebEngineQuotaRequest struct {
	internal.Internal
}

type QWebEngineQuotaRequest_ITF interface {
	QWebEngineQuotaRequest_PTR() *QWebEngineQuotaRequest
}

func (ptr *QWebEngineQuotaRequest) QWebEngineQuotaRequest_PTR() *QWebEngineQuotaRequest {
	return ptr
}

func (ptr *QWebEngineQuotaRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineQuotaRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineQuotaRequest(ptr QWebEngineQuotaRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineQuotaRequest_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineQuotaRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineQuotaRequestFromPointer(ptr unsafe.Pointer) (n *QWebEngineQuotaRequest) {
	n = new(QWebEngineQuotaRequest)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineQuotaRequest")
	return
}

func (ptr *QWebEngineQuotaRequest) DestroyQWebEngineQuotaRequest() {
}

func (ptr *QWebEngineQuotaRequest) Accept() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Accept"})
}

func (ptr *QWebEngineQuotaRequest) Origin() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Origin"}).(*core.QUrl)
}

func (ptr *QWebEngineQuotaRequest) Reject() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reject"})
}

func (ptr *QWebEngineQuotaRequest) RequestedSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestedSize"}).(float64))
}

type QWebEngineRegisterProtocolHandlerRequest struct {
	internal.Internal
}

type QWebEngineRegisterProtocolHandlerRequest_ITF interface {
	QWebEngineRegisterProtocolHandlerRequest_PTR() *QWebEngineRegisterProtocolHandlerRequest
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) QWebEngineRegisterProtocolHandlerRequest_PTR() *QWebEngineRegisterProtocolHandlerRequest {
	return ptr
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineRegisterProtocolHandlerRequest(ptr QWebEngineRegisterProtocolHandlerRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineRegisterProtocolHandlerRequest_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineRegisterProtocolHandlerRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineRegisterProtocolHandlerRequestFromPointer(ptr unsafe.Pointer) (n *QWebEngineRegisterProtocolHandlerRequest) {
	n = new(QWebEngineRegisterProtocolHandlerRequest)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineRegisterProtocolHandlerRequest")
	return
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) DestroyQWebEngineRegisterProtocolHandlerRequest() {
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Accept() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Accept"})
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Origin() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Origin"}).(*core.QUrl)
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Reject() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reject"})
}

func (ptr *QWebEngineRegisterProtocolHandlerRequest) Scheme() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Scheme"}).(string)
}

type QWebEngineScript struct {
	internal.Internal
}

type QWebEngineScript_ITF interface {
	QWebEngineScript_PTR() *QWebEngineScript
}

func (ptr *QWebEngineScript) QWebEngineScript_PTR() *QWebEngineScript {
	return ptr
}

func (ptr *QWebEngineScript) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineScript) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineScript(ptr QWebEngineScript_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineScript_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineScript) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineScriptFromPointer(ptr unsafe.Pointer) (n *QWebEngineScript) {
	n = new(QWebEngineScript)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineScript")
	return
}

func (ptr *QWebEngineScript) DestroyQWebEngineScript() {
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
	internal.Internal
}

type QWebEngineScriptCollection_ITF interface {
	QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection
}

func (ptr *QWebEngineScriptCollection) QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection {
	return ptr
}

func (ptr *QWebEngineScriptCollection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineScriptCollection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineScriptCollection(ptr QWebEngineScriptCollection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineScriptCollection_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineScriptCollection) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineScriptCollectionFromPointer(ptr unsafe.Pointer) (n *QWebEngineScriptCollection) {
	n = new(QWebEngineScriptCollection)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineScriptCollection")
	return
}

func (ptr *QWebEngineScriptCollection) DestroyQWebEngineScriptCollection() {
}

type QWebEngineSettings struct {
	internal.Internal
}

type QWebEngineSettings_ITF interface {
	QWebEngineSettings_PTR() *QWebEngineSettings
}

func (ptr *QWebEngineSettings) QWebEngineSettings_PTR() *QWebEngineSettings {
	return ptr
}

func (ptr *QWebEngineSettings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineSettings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineSettings(ptr QWebEngineSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineSettings_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineSettings) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineSettingsFromPointer(ptr unsafe.Pointer) (n *QWebEngineSettings) {
	n = new(QWebEngineSettings)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineSettings")
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineSettings2", "", parentSettings}).(*QWebEngineSettings)
}

func QWebEngineSettings_DefaultSettings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineSettings_DefaultSettings", ""}).(*QWebEngineSettings)
}

func (ptr *QWebEngineSettings) DefaultSettings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineSettings_DefaultSettings", ""}).(*QWebEngineSettings)
}

func (ptr *QWebEngineSettings) DefaultTextEncoding() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultTextEncoding"}).(string)
}

func (ptr *QWebEngineSettings) FontFamily(which QWebEngineSettings__FontFamily) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontFamily", which}).(string)
}

func (ptr *QWebEngineSettings) FontSize(ty QWebEngineSettings__FontSize) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontSize", ty}).(float64))
}

func QWebEngineSettings_GlobalSettings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineSettings_GlobalSettings", ""}).(*QWebEngineSettings)
}

func (ptr *QWebEngineSettings) GlobalSettings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineSettings_GlobalSettings", ""}).(*QWebEngineSettings)
}

func (ptr *QWebEngineSettings) ResetAttribute(attr QWebEngineSettings__WebAttribute) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetAttribute", attr})
}

func (ptr *QWebEngineSettings) ResetFontFamily(which QWebEngineSettings__FontFamily) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetFontFamily", which})
}

func (ptr *QWebEngineSettings) ResetFontSize(ty QWebEngineSettings__FontSize) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetFontSize", ty})
}

func (ptr *QWebEngineSettings) ResetUnknownUrlSchemePolicy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetUnknownUrlSchemePolicy"})
}

func (ptr *QWebEngineSettings) SetAttribute(attr QWebEngineSettings__WebAttribute, on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", attr, on})
}

func (ptr *QWebEngineSettings) SetDefaultTextEncoding(encoding string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDefaultTextEncoding", encoding})
}

func (ptr *QWebEngineSettings) SetFontFamily(which QWebEngineSettings__FontFamily, family string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFontFamily", which, family})
}

func (ptr *QWebEngineSettings) SetFontSize(ty QWebEngineSettings__FontSize, size int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFontSize", ty, size})
}

func (ptr *QWebEngineSettings) SetUnknownUrlSchemePolicy(policy QWebEngineSettings__UnknownUrlSchemePolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUnknownUrlSchemePolicy", policy})
}

func (ptr *QWebEngineSettings) TestAttribute(attr QWebEngineSettings__WebAttribute) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TestAttribute", attr}).(bool)
}

func (ptr *QWebEngineSettings) UnknownUrlSchemePolicy() QWebEngineSettings__UnknownUrlSchemePolicy {

	return QWebEngineSettings__UnknownUrlSchemePolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnknownUrlSchemePolicy"}).(float64))
}

func (ptr *QWebEngineSettings) DestroyQWebEngineSettings() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineSettings"})
}

type QWebEngineUrlRequestInfo struct {
	internal.Internal
}

type QWebEngineUrlRequestInfo_ITF interface {
	QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo
}

func (ptr *QWebEngineUrlRequestInfo) QWebEngineUrlRequestInfo_PTR() *QWebEngineUrlRequestInfo {
	return ptr
}

func (ptr *QWebEngineUrlRequestInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineUrlRequestInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineUrlRequestInfo(ptr QWebEngineUrlRequestInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlRequestInfo_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineUrlRequestInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineUrlRequestInfoFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestInfo) {
	n = new(QWebEngineUrlRequestInfo)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineUrlRequestInfo")
	return
}

func (ptr *QWebEngineUrlRequestInfo) DestroyQWebEngineUrlRequestInfo() {
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

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Block", shouldBlock})
}

func (ptr *QWebEngineUrlRequestInfo) FirstPartyUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstPartyUrl"}).(*core.QUrl)
}

func (ptr *QWebEngineUrlRequestInfo) NavigationType() QWebEngineUrlRequestInfo__NavigationType {

	return QWebEngineUrlRequestInfo__NavigationType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NavigationType"}).(float64))
}

func (ptr *QWebEngineUrlRequestInfo) Redirect(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Redirect", url})
}

func (ptr *QWebEngineUrlRequestInfo) RequestMethod() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestMethod"}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestInfo) RequestUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUrl"}).(*core.QUrl)
}

func (ptr *QWebEngineUrlRequestInfo) ResourceType() QWebEngineUrlRequestInfo__ResourceType {

	return QWebEngineUrlRequestInfo__ResourceType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResourceType"}).(float64))
}

func (ptr *QWebEngineUrlRequestInfo) SetHttpHeader(name core.QByteArray_ITF, value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpHeader", name, value})
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

func (n *QWebEngineUrlRequestInterceptor) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineUrlRequestInterceptor) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEngineUrlRequestInterceptorFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestInterceptor) {
	n = new(QWebEngineUrlRequestInterceptor)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineUrlRequestInterceptor")
	return
}

func (ptr *QWebEngineUrlRequestInterceptor) DestroyQWebEngineUrlRequestInterceptor() {
}

func NewQWebEngineUrlRequestInterceptor2(p core.QObject_ITF) *QWebEngineUrlRequestInterceptor {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineUrlRequestInterceptor2", "", p}).(*QWebEngineUrlRequestInterceptor)
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectInterceptRequest(f func(info *QWebEngineUrlRequestInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInterceptRequest", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectInterceptRequest() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInterceptRequest"})
}

func (ptr *QWebEngineUrlRequestInterceptor) InterceptRequest(info QWebEngineUrlRequestInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InterceptRequest", info})
}

func (ptr *QWebEngineUrlRequestInterceptor) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlRequestInterceptor) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineUrlRequestInterceptor) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestInterceptor) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestInterceptor) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineUrlRequestInterceptor) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineUrlRequestInterceptor) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestInterceptor) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineUrlRequestInterceptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineUrlRequestInterceptor) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineUrlRequestInterceptor) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineUrlRequestInterceptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineUrlRequestInterceptor) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebEngineUrlRequestInterceptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineUrlRequestInterceptor) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineUrlRequestInterceptor) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QWebEngineUrlRequestJob) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineUrlRequestJob) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEngineUrlRequestJobFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlRequestJob) {
	n = new(QWebEngineUrlRequestJob)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineUrlRequestJob")
	return
}

func (ptr *QWebEngineUrlRequestJob) DestroyQWebEngineUrlRequestJob() {
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

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Fail", r})
}

func (ptr *QWebEngineUrlRequestJob) Initiator() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Initiator"}).(*core.QUrl)
}

func (ptr *QWebEngineUrlRequestJob) Redirect(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Redirect", url})
}

func (ptr *QWebEngineUrlRequestJob) Reply(contentType core.QByteArray_ITF, device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reply", contentType, device})
}

func (ptr *QWebEngineUrlRequestJob) RequestHeaders() map[*core.QByteArray]*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestHeaders"}).(map[*core.QByteArray]*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestJob) RequestMethod() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestMethod"}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestJob) RequestUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUrl"}).(*core.QUrl)
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_atList(v core.QByteArray_ITF, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__requestHeaders_atList", v, i}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_setList(key core.QByteArray_ITF, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__requestHeaders_setList", key, i})
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__requestHeaders_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestJob) __requestHeaders_keyList() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__requestHeaders_keyList"}).([]*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestJob) ____requestHeaders_keyList_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____requestHeaders_keyList_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestJob) ____requestHeaders_keyList_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____requestHeaders_keyList_setList", i})
}

func (ptr *QWebEngineUrlRequestJob) ____requestHeaders_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____requestHeaders_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestJob) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlRequestJob) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineUrlRequestJob) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestJob) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlRequestJob) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineUrlRequestJob) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineUrlRequestJob) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlRequestJob) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineUrlRequestJob) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineUrlRequestJob) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineUrlRequestJob) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineUrlRequestJob) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineUrlRequestJob) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebEngineUrlRequestJob) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineUrlRequestJob) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineUrlRequestJob) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QWebEngineUrlScheme struct {
	internal.Internal
}

type QWebEngineUrlScheme_ITF interface {
	QWebEngineUrlScheme_PTR() *QWebEngineUrlScheme
}

func (ptr *QWebEngineUrlScheme) QWebEngineUrlScheme_PTR() *QWebEngineUrlScheme {
	return ptr
}

func (ptr *QWebEngineUrlScheme) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebEngineUrlScheme) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebEngineUrlScheme(ptr QWebEngineUrlScheme_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebEngineUrlScheme_PTR().Pointer()
	}
	return nil
}

func (n *QWebEngineUrlScheme) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebEngineUrlSchemeFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlScheme) {
	n = new(QWebEngineUrlScheme)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineUrlScheme")
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

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineUrlScheme", ""}).(*QWebEngineUrlScheme)
}

func NewQWebEngineUrlScheme2(name core.QByteArray_ITF) *QWebEngineUrlScheme {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineUrlScheme2", "", name}).(*QWebEngineUrlScheme)
}

func NewQWebEngineUrlScheme3(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineUrlScheme3", "", that}).(*QWebEngineUrlScheme)
}

func NewQWebEngineUrlScheme4(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineUrlScheme4", "", that}).(*QWebEngineUrlScheme)
}

func (ptr *QWebEngineUrlScheme) DefaultPort() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultPort"}).(float64))
}

func (ptr *QWebEngineUrlScheme) Flags() QWebEngineUrlScheme__Flag {

	return QWebEngineUrlScheme__Flag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QWebEngineUrlScheme) Name() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(*core.QByteArray)
}

func QWebEngineUrlScheme_RegisterScheme(scheme QWebEngineUrlScheme_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineUrlScheme_RegisterScheme", "", scheme})
}

func (ptr *QWebEngineUrlScheme) RegisterScheme(scheme QWebEngineUrlScheme_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineUrlScheme_RegisterScheme", "", scheme})
}

func QWebEngineUrlScheme_SchemeByName(name core.QByteArray_ITF) *QWebEngineUrlScheme {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineUrlScheme_SchemeByName", "", name}).(*QWebEngineUrlScheme)
}

func (ptr *QWebEngineUrlScheme) SchemeByName(name core.QByteArray_ITF) *QWebEngineUrlScheme {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineUrlScheme_SchemeByName", "", name}).(*QWebEngineUrlScheme)
}

func (ptr *QWebEngineUrlScheme) SetDefaultPort(newValue int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDefaultPort", newValue})
}

func (ptr *QWebEngineUrlScheme) SetFlags(newValue QWebEngineUrlScheme__Flag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlags", newValue})
}

func (ptr *QWebEngineUrlScheme) SetName(newValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", newValue})
}

func (ptr *QWebEngineUrlScheme) SetSyntax(newValue QWebEngineUrlScheme__Syntax) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSyntax", newValue})
}

func (ptr *QWebEngineUrlScheme) DestroyQWebEngineUrlScheme() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineUrlScheme"})
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

func (n *QWebEngineUrlSchemeHandler) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineUrlSchemeHandler) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebEngineUrlSchemeHandlerFromPointer(ptr unsafe.Pointer) (n *QWebEngineUrlSchemeHandler) {
	n = new(QWebEngineUrlSchemeHandler)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineUrlSchemeHandler")
	return
}
func NewQWebEngineUrlSchemeHandler(parent core.QObject_ITF) *QWebEngineUrlSchemeHandler {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineUrlSchemeHandler", "", parent}).(*QWebEngineUrlSchemeHandler)
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectRequestStarted(f func(request *QWebEngineUrlRequestJob)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectRequestStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestStarted"})
}

func (ptr *QWebEngineUrlSchemeHandler) RequestStarted(request QWebEngineUrlRequestJob_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestStarted", request})
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectDestroyQWebEngineUrlSchemeHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebEngineUrlSchemeHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectDestroyQWebEngineUrlSchemeHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebEngineUrlSchemeHandler"})
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineUrlSchemeHandler"})
}

func (ptr *QWebEngineUrlSchemeHandler) DestroyQWebEngineUrlSchemeHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineUrlSchemeHandlerDefault"})
}

func (ptr *QWebEngineUrlSchemeHandler) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlSchemeHandler) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineUrlSchemeHandler) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlSchemeHandler) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineUrlSchemeHandler) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineUrlSchemeHandler) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineUrlSchemeHandler) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineUrlSchemeHandler) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineUrlSchemeHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineUrlSchemeHandler) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineUrlSchemeHandler) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineUrlSchemeHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineUrlSchemeHandler) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebEngineUrlSchemeHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineUrlSchemeHandler) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineUrlSchemeHandler) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QWebEngineView) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebEngineView) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQWebEngineViewFromPointer(ptr unsafe.Pointer) (n *QWebEngineView) {
	n = new(QWebEngineView)
	n.InitFromInternal(uintptr(ptr), "webengine.QWebEngineView")
	return
}
func NewQWebEngineView(parent widgets.QWidget_ITF) *QWebEngineView {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.NewQWebEngineView", "", parent}).(*QWebEngineView)
}

func (ptr *QWebEngineView) ConnectBack(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBack", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectBack() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBack"})
}

func (ptr *QWebEngineView) Back() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Back"})
}

func (ptr *QWebEngineView) BackDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackDefault"})
}

func (ptr *QWebEngineView) ContextMenuEventDefault(vqc gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", vqc})
}

func (ptr *QWebEngineView) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", e})
}

func (ptr *QWebEngineView) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", e})
}

func (ptr *QWebEngineView) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", e})
}

func (ptr *QWebEngineView) DropEventDefault(e gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", e})
}

func (ptr *QWebEngineView) EventDefault(vqe core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", vqe}).(bool)
}

func (ptr *QWebEngineView) ConnectForward(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectForward", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectForward() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectForward"})
}

func (ptr *QWebEngineView) Forward() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Forward"})
}

func (ptr *QWebEngineView) ForwardDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ForwardDefault"})
}

func (ptr *QWebEngineView) HasSelection() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasSelection"}).(bool)
}

func (ptr *QWebEngineView) HideEventDefault(vqh gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", vqh})
}

func (ptr *QWebEngineView) Icon() *gui.QIcon {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Icon"}).(*gui.QIcon)
}

func (ptr *QWebEngineView) ConnectIconChanged(f func(vqi *gui.QIcon)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIconChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectIconChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIconChanged"})
}

func (ptr *QWebEngineView) IconChanged(vqi gui.QIcon_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconChanged", vqi})
}

func (ptr *QWebEngineView) IconUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconUrl"}).(*core.QUrl)
}

func (ptr *QWebEngineView) ConnectIconUrlChanged(f func(vqu *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIconUrlChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectIconUrlChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIconUrlChanged"})
}

func (ptr *QWebEngineView) IconUrlChanged(vqu core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IconUrlChanged", vqu})
}

func (ptr *QWebEngineView) Load(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", url})
}

func (ptr *QWebEngineView) Load2(request QWebEngineHttpRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2", request})
}

func (ptr *QWebEngineView) ConnectLoadFinished(f func(vbo bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectLoadFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadFinished"})
}

func (ptr *QWebEngineView) LoadFinished(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadFinished", vbo})
}

func (ptr *QWebEngineView) ConnectLoadProgress(f func(progress int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadProgress", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectLoadProgress() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadProgress"})
}

func (ptr *QWebEngineView) LoadProgress(progress int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadProgress", progress})
}

func (ptr *QWebEngineView) ConnectLoadStarted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoadStarted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectLoadStarted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoadStarted"})
}

func (ptr *QWebEngineView) LoadStarted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadStarted"})
}

func (ptr *QWebEngineView) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebEngineView) Page() *QWebEnginePage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Page"}).(*QWebEnginePage)
}

func (ptr *QWebEngineView) ConnectReload(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReload", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectReload() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReload"})
}

func (ptr *QWebEngineView) Reload() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reload"})
}

func (ptr *QWebEngineView) ReloadDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReloadDefault"})
}

func (ptr *QWebEngineView) SelectedText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedText"}).(string)
}

func (ptr *QWebEngineView) ConnectSelectionChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectSelectionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionChanged"})
}

func (ptr *QWebEngineView) SelectionChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionChanged"})
}

func (ptr *QWebEngineView) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent", data, mimeType, baseUrl})
}

func (ptr *QWebEngineView) SetHtml(html string, baseUrl core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHtml", html, baseUrl})
}

func (ptr *QWebEngineView) SetPage(page QWebEnginePage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPage", page})
}

func (ptr *QWebEngineView) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QWebEngineView) SetZoomFactor(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomFactor", factor})
}

func (ptr *QWebEngineView) Settings() *QWebEngineSettings {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Settings"}).(*QWebEngineSettings)
}

func (ptr *QWebEngineView) ShowEventDefault(vqs gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", vqs})
}

func (ptr *QWebEngineView) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QWebEngineView) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QWebEngineView) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QWebEngineView) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func (ptr *QWebEngineView) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QWebEngineView) ConnectTitleChanged(f func(title string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectTitleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleChanged"})
}

func (ptr *QWebEngineView) TitleChanged(title string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleChanged", title})
}

func QWebEngineView_Tr(s string, c string, n int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineView_Tr", "", s, c, n}).(string)
}

func (ptr *QWebEngineView) Tr(s string, c string, n int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "webengine.QWebEngineView_Tr", "", s, c, n}).(string)
}

func (ptr *QWebEngineView) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QWebEngineView) ConnectUrlChanged(f func(vqu *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUrlChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectUrlChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUrlChanged"})
}

func (ptr *QWebEngineView) UrlChanged(vqu core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UrlChanged", vqu})
}

func (ptr *QWebEngineView) ZoomFactor() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomFactor"}).(float64)
}

func (ptr *QWebEngineView) ConnectDestroyQWebEngineView(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebEngineView", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebEngineView) DisconnectDestroyQWebEngineView() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebEngineView"})
}

func (ptr *QWebEngineView) DestroyQWebEngineView() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineView"})
}

func (ptr *QWebEngineView) DestroyQWebEngineViewDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebEngineViewDefault"})
}

func (ptr *QWebEngineView) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QWebEngineView) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QWebEngineView) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QWebEngineView) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QWebEngineView) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QWebEngineView) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QWebEngineView) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineView) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebEngineView) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebEngineView) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebEngineView) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebEngineView) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebEngineView) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebEngineView) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebEngineView) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebEngineView) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QWebEngineView) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QWebEngineView) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QWebEngineView) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QWebEngineView) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QWebEngineView) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QWebEngineView) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QWebEngineView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QWebEngineView) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QWebEngineView) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QWebEngineView) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QWebEngineView) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QWebEngineView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QWebEngineView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QWebEngineView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QWebEngineView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QWebEngineView) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QWebEngineView) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QWebEngineView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QWebEngineView) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QWebEngineView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QWebEngineView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QWebEngineView) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QWebEngineView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QWebEngineView) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QWebEngineView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QWebEngineView) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QWebEngineView) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QWebEngineView) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QWebEngineView) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QWebEngineView) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QWebEngineView) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QWebEngineView) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QWebEngineView) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QWebEngineView) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QWebEngineView) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QWebEngineView) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QWebEngineView) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QWebEngineView) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QWebEngineView) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QWebEngineView) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QWebEngineView) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QWebEngineView) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QWebEngineView) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QWebEngineView) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QWebEngineView) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QWebEngineView) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QWebEngineView) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QWebEngineView) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebEngineView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebEngineView) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebEngineView) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebEngineView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebEngineView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebEngineView) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QtWebEngine struct {
	internal.Internal
}

type QtWebEngine_ITF interface {
	QtWebEngine_PTR() *QtWebEngine
}

func (ptr *QtWebEngine) QtWebEngine_PTR() *QtWebEngine {
	return ptr
}

func (ptr *QtWebEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QtWebEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQtWebEngine(ptr QtWebEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWebEngine_PTR().Pointer()
	}
	return nil
}

func (n *QtWebEngine) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQtWebEngineFromPointer(ptr unsafe.Pointer) (n *QtWebEngine) {
	n = new(QtWebEngine)
	n.InitFromInternal(uintptr(ptr), "webengine.QtWebEngine")
	return
}

func (ptr *QtWebEngine) DestroyQtWebEngine() {
}

func QtWebEngine_Initialize() {

	internal.CallLocalFunction([]interface{}{"", "", "webengine.QtWebEngine_Initialize", ""})
}

func (ptr *QtWebEngine) Initialize() {

	internal.CallLocalFunction([]interface{}{"", "", "webengine.QtWebEngine_Initialize", ""})
}

type RenderViewObserverQt struct {
	internal.Internal
}

type RenderViewObserverQt_ITF interface {
	RenderViewObserverQt_PTR() *RenderViewObserverQt
}

func (ptr *RenderViewObserverQt) RenderViewObserverQt_PTR() *RenderViewObserverQt {
	return ptr
}

func (ptr *RenderViewObserverQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *RenderViewObserverQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromRenderViewObserverQt(ptr RenderViewObserverQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.RenderViewObserverQt_PTR().Pointer()
	}
	return nil
}

func (n *RenderViewObserverQt) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewRenderViewObserverQtFromPointer(ptr unsafe.Pointer) (n *RenderViewObserverQt) {
	n = new(RenderViewObserverQt)
	n.InitFromInternal(uintptr(ptr), "webengine.RenderViewObserverQt")
	return
}

func (ptr *RenderViewObserverQt) DestroyRenderViewObserverQt() {
}

type ServiceQt struct {
	internal.Internal
}

type ServiceQt_ITF interface {
	ServiceQt_PTR() *ServiceQt
}

func (ptr *ServiceQt) ServiceQt_PTR() *ServiceQt {
	return ptr
}

func (ptr *ServiceQt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *ServiceQt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromServiceQt(ptr ServiceQt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ServiceQt_PTR().Pointer()
	}
	return nil
}

func (n *ServiceQt) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewServiceQtFromPointer(ptr unsafe.Pointer) (n *ServiceQt) {
	n = new(ServiceQt)
	n.InitFromInternal(uintptr(ptr), "webengine.ServiceQt")
	return
}

func (ptr *ServiceQt) DestroyServiceQt() {
}

type UserResourceController struct {
	internal.Internal
}

type UserResourceController_ITF interface {
	UserResourceController_PTR() *UserResourceController
}

func (ptr *UserResourceController) UserResourceController_PTR() *UserResourceController {
	return ptr
}

func (ptr *UserResourceController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *UserResourceController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromUserResourceController(ptr UserResourceController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.UserResourceController_PTR().Pointer()
	}
	return nil
}

func (n *UserResourceController) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewUserResourceControllerFromPointer(ptr unsafe.Pointer) (n *UserResourceController) {
	n = new(UserResourceController)
	n.InitFromInternal(uintptr(ptr), "webengine.UserResourceController")
	return
}

func (ptr *UserResourceController) DestroyUserResourceController() {
}

type UserScriptData struct {
	internal.Internal
}

type UserScriptData_ITF interface {
	UserScriptData_PTR() *UserScriptData
}

func (ptr *UserScriptData) UserScriptData_PTR() *UserScriptData {
	return ptr
}

func (ptr *UserScriptData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *UserScriptData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromUserScriptData(ptr UserScriptData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.UserScriptData_PTR().Pointer()
	}
	return nil
}

func (n *UserScriptData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewUserScriptDataFromPointer(ptr unsafe.Pointer) (n *UserScriptData) {
	n = new(UserScriptData)
	n.InitFromInternal(uintptr(ptr), "webengine.UserScriptData")
	return
}

func (ptr *UserScriptData) DestroyUserScriptData() {
}

type WebEngineError struct {
	internal.Internal
}

type WebEngineError_ITF interface {
	WebEngineError_PTR() *WebEngineError
}

func (ptr *WebEngineError) WebEngineError_PTR() *WebEngineError {
	return ptr
}

func (ptr *WebEngineError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *WebEngineError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromWebEngineError(ptr WebEngineError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WebEngineError_PTR().Pointer()
	}
	return nil
}

func (n *WebEngineError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewWebEngineErrorFromPointer(ptr unsafe.Pointer) (n *WebEngineError) {
	n = new(WebEngineError)
	n.InitFromInternal(uintptr(ptr), "webengine.WebEngineError")
	return
}

func (ptr *WebEngineError) DestroyWebEngineError() {
}

type WebEngineLibraryInfo struct {
	internal.Internal
}

type WebEngineLibraryInfo_ITF interface {
	WebEngineLibraryInfo_PTR() *WebEngineLibraryInfo
}

func (ptr *WebEngineLibraryInfo) WebEngineLibraryInfo_PTR() *WebEngineLibraryInfo {
	return ptr
}

func (ptr *WebEngineLibraryInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *WebEngineLibraryInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromWebEngineLibraryInfo(ptr WebEngineLibraryInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WebEngineLibraryInfo_PTR().Pointer()
	}
	return nil
}

func (n *WebEngineLibraryInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewWebEngineLibraryInfoFromPointer(ptr unsafe.Pointer) (n *WebEngineLibraryInfo) {
	n = new(WebEngineLibraryInfo)
	n.InitFromInternal(uintptr(ptr), "webengine.WebEngineLibraryInfo")
	return
}

func (ptr *WebEngineLibraryInfo) DestroyWebEngineLibraryInfo() {
}

type WebEventFactory struct {
	internal.Internal
}

type WebEventFactory_ITF interface {
	WebEventFactory_PTR() *WebEventFactory
}

func (ptr *WebEventFactory) WebEventFactory_PTR() *WebEventFactory {
	return ptr
}

func (ptr *WebEventFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *WebEventFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromWebEventFactory(ptr WebEventFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WebEventFactory_PTR().Pointer()
	}
	return nil
}

func (n *WebEventFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewWebEventFactoryFromPointer(ptr unsafe.Pointer) (n *WebEventFactory) {
	n = new(WebEventFactory)
	n.InitFromInternal(uintptr(ptr), "webengine.WebEventFactory")
	return
}

func (ptr *WebEventFactory) DestroyWebEventFactory() {
}

func init() {
	internal.ConstructorTable["webengine.QQuickWebEngineProfile"] = NewQQuickWebEngineProfileFromPointer
	internal.ConstructorTable["webengine.QQuickWebEngineScript"] = NewQQuickWebEngineScriptFromPointer
	internal.ConstructorTable["webengine.QWebEngineCallback"] = NewQWebEngineCallbackFromPointer
	internal.ConstructorTable["webengine.QWebEngineCertificateError"] = NewQWebEngineCertificateErrorFromPointer
	internal.ConstructorTable["webengine.QWebEngineClientCertificateSelection"] = NewQWebEngineClientCertificateSelectionFromPointer
	internal.ConstructorTable["webengine.QWebEngineClientCertificateStore"] = NewQWebEngineClientCertificateStoreFromPointer
	internal.ConstructorTable["webengine.QWebEngineContextMenuData"] = NewQWebEngineContextMenuDataFromPointer
	internal.ConstructorTable["webengine.QWebEngineCookieStore"] = NewQWebEngineCookieStoreFromPointer
	internal.ConstructorTable["webengine.QWebEngineHttpRequest"] = NewQWebEngineHttpRequestFromPointer
	internal.ConstructorTable["webengine.QWebEngineNotification"] = NewQWebEngineNotificationFromPointer
	internal.ConstructorTable["webengine.QWebEnginePage"] = NewQWebEnginePageFromPointer
	internal.ConstructorTable["webengine.QWebEngineProfile"] = NewQWebEngineProfileFromPointer
	internal.ConstructorTable["webengine.QWebEngineQuotaRequest"] = NewQWebEngineQuotaRequestFromPointer
	internal.ConstructorTable["webengine.QWebEngineRegisterProtocolHandlerRequest"] = NewQWebEngineRegisterProtocolHandlerRequestFromPointer
	internal.ConstructorTable["webengine.QWebEngineScript"] = NewQWebEngineScriptFromPointer
	internal.ConstructorTable["webengine.QWebEngineScriptCollection"] = NewQWebEngineScriptCollectionFromPointer
	internal.ConstructorTable["webengine.QWebEngineSettings"] = NewQWebEngineSettingsFromPointer
	internal.ConstructorTable["webengine.QWebEngineUrlRequestInfo"] = NewQWebEngineUrlRequestInfoFromPointer
	internal.ConstructorTable["webengine.QWebEngineUrlRequestInterceptor"] = NewQWebEngineUrlRequestInterceptorFromPointer
	internal.ConstructorTable["webengine.QWebEngineUrlRequestJob"] = NewQWebEngineUrlRequestJobFromPointer
	internal.ConstructorTable["webengine.QWebEngineUrlScheme"] = NewQWebEngineUrlSchemeFromPointer
	internal.ConstructorTable["webengine.QWebEngineUrlSchemeHandler"] = NewQWebEngineUrlSchemeHandlerFromPointer
	internal.ConstructorTable["webengine.QWebEngineView"] = NewQWebEngineViewFromPointer
	internal.ConstructorTable["webengine.QtWebEngine"] = NewQtWebEngineFromPointer
}
