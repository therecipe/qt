package network

//#include "network.h"
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
	for len(n.ObjectName()) < len("QNetworkAccessManager_") {
		n.SetObjectName("QNetworkAccessManager_" + qt.Identifier())
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
	defer qt.Recovering("QNetworkAccessManager::proxyFactory")

	if ptr.Pointer() != nil {
		return NewQNetworkProxyFactoryFromPointer(C.QNetworkAccessManager_ProxyFactory(ptr.Pointer()))
	}
	return nil
}

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {
	defer qt.Recovering("QNetworkAccessManager::QNetworkAccessManager")

	return NewQNetworkAccessManagerFromPointer(C.QNetworkAccessManager_NewQNetworkAccessManager(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "authenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "authenticationRequired")
	}
}

//export callbackQNetworkAccessManagerAuthenticationRequired
func callbackQNetworkAccessManagerAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::authenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "authenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF, authenticator QAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_AuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	defer qt.Recovering("QNetworkAccessManager::cache")

	if ptr.Pointer() != nil {
		return NewQAbstractNetworkCacheFromPointer(C.QNetworkAccessManager_Cache(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {
	defer qt.Recovering("QNetworkAccessManager::clearAccessCache")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearAccessCache(ptr.Pointer())
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	defer qt.Recovering("QNetworkAccessManager::cookieJar")

	if ptr.Pointer() != nil {
		return NewQNetworkCookieJarFromPointer(C.QNetworkAccessManager_CookieJar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) CreateRequest(op QNetworkAccessManager__Operation, req QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::createRequest")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequest(ptr.Pointer(), C.int(op), PointerFromQNetworkRequest(req), core.PointerFromQIODevice(outgoingData)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::deleteResource")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_DeleteResource(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {
	defer qt.Recovering("connect QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQNetworkAccessManagerEncrypted
func callbackQNetworkAccessManagerEncrypted(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::encrypted")

	if signal := qt.GetSignal(C.GoString(ptrName), "encrypted"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF) {
	defer qt.Recovering("QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Encrypted(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {
	defer qt.Recovering("connect QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	defer qt.Recovering("disconnect QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQNetworkAccessManagerFinished
func callbackQNetworkAccessManagerFinished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) Finished(reply QNetworkReply_ITF) {
	defer qt.Recovering("QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Finished(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::get")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Get(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::head")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Head(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	defer qt.Recovering("QNetworkAccessManager::networkAccessible")

	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	defer qt.Recovering("connect QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "networkAccessibleChanged", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	defer qt.Recovering("disconnect QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "networkAccessibleChanged")
	}
}

//export callbackQNetworkAccessManagerNetworkAccessibleChanged
func callbackQNetworkAccessManagerNetworkAccessibleChanged(ptr unsafe.Pointer, ptrName *C.char, accessible C.int) {
	defer qt.Recovering("callback QNetworkAccessManager::networkAccessibleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "networkAccessibleChanged"); signal != nil {
		signal.(func(QNetworkAccessManager__NetworkAccessibility))(QNetworkAccessManager__NetworkAccessibility(accessible))
	}

}

func (ptr *QNetworkAccessManager) NetworkAccessibleChanged(accessible QNetworkAccessManager__NetworkAccessibility) {
	defer qt.Recovering("QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_NetworkAccessibleChanged(ptr.Pointer(), C.int(accessible))
	}
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF, authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put2(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put3(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::sendCustomRequest")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setCache")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCache(ptr.Pointer(), PointerFromQAbstractNetworkCache(cache))
	}
}

func (ptr *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCookieJar(ptr.Pointer(), PointerFromQNetworkCookieJar(cookieJar))
	}
}

func (ptr *QNetworkAccessManager) SetNetworkAccessible(accessible QNetworkAccessManager__NetworkAccessibility) {
	defer qt.Recovering("QNetworkAccessManager::setNetworkAccessible")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetNetworkAccessible(ptr.Pointer(), C.int(accessible))
	}
}

func (ptr *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setProxy")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(proxy))
	}
}

func (ptr *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxyFactory(ptr.Pointer(), PointerFromQNetworkProxyFactory(factory))
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {
	defer qt.Recovering("QNetworkAccessManager::supportedSchemes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QNetworkAccessManager_SupportedSchemes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	defer qt.Recovering("QNetworkAccessManager::~QNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkAccessManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkAccessManagerTimerEvent
func callbackQNetworkAccessManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkAccessManagerChildEvent
func callbackQNetworkAccessManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkAccessManagerCustomEvent
func callbackQNetworkAccessManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
