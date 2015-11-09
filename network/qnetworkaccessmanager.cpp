#include "qnetworkaccessmanager.h"
#include <QString>
#include <QModelIndex>
#include <QByteArray>
#include <QNetworkProxy>
#include <QAuthenticator>
#include <QNetworkReply>
#include <QNetworkCookieJar>
#include <QVariant>
#include <QSslPreSharedKeyAuthenticator>
#include <QObject>
#include <QHttpMultiPart>
#include <QAbstractNetworkCache>
#include <QUrl>
#include <QNetworkProxyFactory>
#include <QIODevice>
#include <QNetworkCookie>
#include <QNetworkConfiguration>
#include <QNetworkRequest>
#include <QNetworkAccessManager>
#include "_cgo_export.h"

class MyQNetworkAccessManager: public QNetworkAccessManager {
public:
void Signal_AuthenticationRequired(QNetworkReply * reply, QAuthenticator * authenticator){callbackQNetworkAccessManagerAuthenticationRequired(this->objectName().toUtf8().data(), reply, authenticator);};
void Signal_Encrypted(QNetworkReply * reply){callbackQNetworkAccessManagerEncrypted(this->objectName().toUtf8().data(), reply);};
void Signal_Finished(QNetworkReply * reply){callbackQNetworkAccessManagerFinished(this->objectName().toUtf8().data(), reply);};
void Signal_NetworkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility accessible){callbackQNetworkAccessManagerNetworkAccessibleChanged(this->objectName().toUtf8().data(), accessible);};
void Signal_PreSharedKeyAuthenticationRequired(QNetworkReply * reply, QSslPreSharedKeyAuthenticator * authenticator){callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired(this->objectName().toUtf8().data(), reply, authenticator);};
};

void* QNetworkAccessManager_ProxyFactory(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->proxyFactory();
}

void* QNetworkAccessManager_NewQNetworkAccessManager(void* parent){
	return new QNetworkAccessManager(static_cast<QObject*>(parent));
}

void QNetworkAccessManager_ConnectAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));;
}

void QNetworkAccessManager_DisconnectAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));;
}

void* QNetworkAccessManager_Cache(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->cache();
}

void QNetworkAccessManager_ClearAccessCache(void* ptr){
	static_cast<QNetworkAccessManager*>(ptr)->clearAccessCache();
}

void* QNetworkAccessManager_CookieJar(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->cookieJar();
}

void* QNetworkAccessManager_DeleteResource(void* ptr, void* request){
	return static_cast<QNetworkAccessManager*>(ptr)->deleteResource(*static_cast<QNetworkRequest*>(request));
}

void QNetworkAccessManager_ConnectEncrypted(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));;
}

void QNetworkAccessManager_DisconnectEncrypted(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));;
}

void QNetworkAccessManager_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));;
}

void QNetworkAccessManager_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));;
}

void* QNetworkAccessManager_Get(void* ptr, void* request){
	return static_cast<QNetworkAccessManager*>(ptr)->get(*static_cast<QNetworkRequest*>(request));
}

void* QNetworkAccessManager_Head(void* ptr, void* request){
	return static_cast<QNetworkAccessManager*>(ptr)->head(*static_cast<QNetworkRequest*>(request));
}

int QNetworkAccessManager_NetworkAccessible(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->networkAccessible();
}

void QNetworkAccessManager_ConnectNetworkAccessibleChanged(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));;
}

void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));;
}

void* QNetworkAccessManager_Post3(void* ptr, void* request, void* multiPart){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Post(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Post2(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));;
}

void* QNetworkAccessManager_Put2(void* ptr, void* request, void* multiPart){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Put(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Put3(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void* QNetworkAccessManager_SendCustomRequest(void* ptr, void* request, void* verb, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), static_cast<QIODevice*>(data));
}

void QNetworkAccessManager_SetCache(void* ptr, void* cache){
	static_cast<QNetworkAccessManager*>(ptr)->setCache(static_cast<QAbstractNetworkCache*>(cache));
}

void QNetworkAccessManager_SetConfiguration(void* ptr, void* config){
	static_cast<QNetworkAccessManager*>(ptr)->setConfiguration(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkAccessManager_SetCookieJar(void* ptr, void* cookieJar){
	static_cast<QNetworkAccessManager*>(ptr)->setCookieJar(static_cast<QNetworkCookieJar*>(cookieJar));
}

void QNetworkAccessManager_SetNetworkAccessible(void* ptr, int accessible){
	static_cast<QNetworkAccessManager*>(ptr)->setNetworkAccessible(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void QNetworkAccessManager_SetProxy(void* ptr, void* proxy){
	static_cast<QNetworkAccessManager*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(proxy));
}

void QNetworkAccessManager_SetProxyFactory(void* ptr, void* factory){
	static_cast<QNetworkAccessManager*>(ptr)->setProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

char* QNetworkAccessManager_SupportedSchemes(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->supportedSchemes().join("|").toUtf8().data();
}

void QNetworkAccessManager_DestroyQNetworkAccessManager(void* ptr){
	static_cast<QNetworkAccessManager*>(ptr)->~QNetworkAccessManager();
}

