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

type QNetworkAccessManager_ITF interface {
	core.QObject_ITF
	QNetworkAccessManager_PTR() *QNetworkAccessManager
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
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkAccessManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkAccessManager) QNetworkAccessManager_PTR() *QNetworkAccessManager {
	return ptr
}

//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           = QNetworkAccessManager__NetworkAccessibility(1)
)

//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int64

const (
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
		return NewQNetworkProxyFactoryFromPointer(C.QNetworkAccessManager_ProxyFactory(ptr.Pointer()))
	}
	return nil
}

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {
	return NewQNetworkAccessManagerFromPointer(C.QNetworkAccessManager_NewQNetworkAccessManager(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "authenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "authenticationRequired")
	}
}

//export callbackQNetworkAccessManagerAuthenticationRequired
func callbackQNetworkAccessManagerAuthenticationRequired(ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "authenticationRequired").(func(*QNetworkReply, *QAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQAuthenticatorFromPointer(authenticator))
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	if ptr.Pointer() != nil {
		return NewQAbstractNetworkCacheFromPointer(C.QNetworkAccessManager_Cache(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearAccessCache(ptr.Pointer())
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	if ptr.Pointer() != nil {
		return NewQNetworkCookieJarFromPointer(C.QNetworkAccessManager_CookieJar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_DeleteResource(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQNetworkAccessManagerEncrypted
func callbackQNetworkAccessManagerEncrypted(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "encrypted").(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQNetworkAccessManagerFinished
func callbackQNetworkAccessManagerFinished(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Get(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Head(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "networkAccessibleChanged", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "networkAccessibleChanged")
	}
}

//export callbackQNetworkAccessManagerNetworkAccessibleChanged
func callbackQNetworkAccessManagerNetworkAccessibleChanged(ptrName *C.char, accessible C.int) {
	qt.GetSignal(C.GoString(ptrName), "networkAccessibleChanged").(func(QNetworkAccessManager__NetworkAccessibility))(QNetworkAccessManager__NetworkAccessibility(accessible))
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired(ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired").(func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put2(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put3(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QIODevice_ITF) *QNetworkReply {
	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), core.PointerFromQIODevice(data)))
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
		C.QNetworkAccessManager_SetNetworkAccessible(ptr.Pointer(), C.int(accessible))
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
		return strings.Split(C.GoString(C.QNetworkAccessManager_SupportedSchemes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
