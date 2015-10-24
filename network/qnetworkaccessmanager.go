package network

//#include "qnetworkaccessmanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QNetworkAccessManager struct {
	core.QObject
}

type QNetworkAccessManagerITF interface {
	core.QObjectITF
	QNetworkAccessManagerPTR() *QNetworkAccessManager
}

func PointerFromQNetworkAccessManager(ptr QNetworkAccessManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAccessManagerPTR().Pointer()
	}
	return nil
}

func QNetworkAccessManagerFromPointer(ptr unsafe.Pointer) *QNetworkAccessManager {
	var n = new(QNetworkAccessManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkAccessManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkAccessManager) QNetworkAccessManagerPTR() *QNetworkAccessManager {
	return ptr
}

//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int

var (
	QNetworkAccessManager__UnknownAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           = QNetworkAccessManager__NetworkAccessibility(1)
)

//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int

var (
	QNetworkAccessManager__HeadOperation    = QNetworkAccessManager__Operation(1)
	QNetworkAccessManager__GetOperation     = QNetworkAccessManager__Operation(2)
	QNetworkAccessManager__PutOperation     = QNetworkAccessManager__Operation(3)
	QNetworkAccessManager__PostOperation    = QNetworkAccessManager__Operation(4)
	QNetworkAccessManager__DeleteOperation  = QNetworkAccessManager__Operation(5)
	QNetworkAccessManager__CustomOperation  = QNetworkAccessManager__Operation(6)
	QNetworkAccessManager__UnknownOperation = QNetworkAccessManager__Operation(0)
)

func (ptr *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {
	if ptr.Pointer() != nil {
		return QNetworkProxyFactoryFromPointer(unsafe.Pointer(C.QNetworkAccessManager_ProxyFactory(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQNetworkAccessManager(parent core.QObjectITF) *QNetworkAccessManager {
	return QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QNetworkAccessManager_NewQNetworkAccessManager(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply QNetworkReplyITF, authenticator QAuthenticatorITF)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "authenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "authenticationRequired")
	}
}

//export callbackQNetworkAccessManagerAuthenticationRequired
func callbackQNetworkAccessManagerAuthenticationRequired(ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "authenticationRequired").(func(*QNetworkReply, *QAuthenticator))(QNetworkReplyFromPointer(reply), QAuthenticatorFromPointer(authenticator))
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	if ptr.Pointer() != nil {
		return QAbstractNetworkCacheFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Cache(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearAccessCache(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	if ptr.Pointer() != nil {
		return QNetworkCookieJarFromPointer(unsafe.Pointer(C.QNetworkAccessManager_CookieJar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequestITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_DeleteResource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply QNetworkReplyITF)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectEncrypted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQNetworkAccessManagerEncrypted
func callbackQNetworkAccessManagerEncrypted(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "encrypted").(func(*QNetworkReply))(QNetworkReplyFromPointer(reply))
}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply QNetworkReplyITF)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQNetworkAccessManagerFinished
func callbackQNetworkAccessManagerFinished(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QNetworkReply))(QNetworkReplyFromPointer(reply))
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequestITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Get(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequestITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Head(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "networkAccessibleChanged", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "networkAccessibleChanged")
	}
}

//export callbackQNetworkAccessManagerNetworkAccessibleChanged
func callbackQNetworkAccessManagerNetworkAccessibleChanged(ptrName *C.char, accessible C.int) {
	qt.GetSignal(C.GoString(ptrName), "networkAccessibleChanged").(func(QNetworkAccessManager__NetworkAccessibility))(QNetworkAccessManager__NetworkAccessibility(accessible))
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequestITF, multiPart QHttpMultiPartITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Post3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(PointerFromQHttpMultiPart(multiPart)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequestITF, data core.QIODeviceITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Post(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(core.PointerFromQIODevice(data)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequestITF, data core.QByteArrayITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Post2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(core.PointerFromQByteArray(data)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply QNetworkReplyITF, authenticator QSslPreSharedKeyAuthenticatorITF)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired(ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired").(func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(QNetworkReplyFromPointer(reply), QSslPreSharedKeyAuthenticatorFromPointer(authenticator))
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequestITF, multiPart QHttpMultiPartITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Put2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(PointerFromQHttpMultiPart(multiPart)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequestITF, data core.QIODeviceITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Put(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(core.PointerFromQIODevice(data)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequestITF, data core.QByteArrayITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_Put3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(core.PointerFromQByteArray(data)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequestITF, verb core.QByteArrayITF, data core.QIODeviceITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return QNetworkReplyFromPointer(unsafe.Pointer(C.QNetworkAccessManager_SendCustomRequest(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(request)), C.QtObjectPtr(core.PointerFromQByteArray(verb)), C.QtObjectPtr(core.PointerFromQIODevice(data)))))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetCache(cache QAbstractNetworkCacheITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCache(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractNetworkCache(cache)))
	}
}

func (ptr *QNetworkAccessManager) SetConfiguration(config QNetworkConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkConfiguration(config)))
	}
}

func (ptr *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJarITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCookieJar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCookieJar(cookieJar)))
	}
}

func (ptr *QNetworkAccessManager) SetNetworkAccessible(accessible QNetworkAccessManager__NetworkAccessibility) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetNetworkAccessible(C.QtObjectPtr(ptr.Pointer()), C.int(accessible))
	}
}

func (ptr *QNetworkAccessManager) SetProxy(proxy QNetworkProxyITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkProxy(proxy)))
	}
}

func (ptr *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactoryITF) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxyFactory(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkProxyFactory(factory)))
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QNetworkAccessManager_SupportedSchemes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
