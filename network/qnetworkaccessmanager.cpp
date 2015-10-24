#include "qnetworkaccessmanager.h"
#include <QNetworkCookieJar>
#include <QNetworkProxy>
#include <QString>
#include <QAbstractNetworkCache>
#include <QAuthenticator>
#include <QNetworkRequest>
#include <QByteArray>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkReply>
#include <QObject>
#include <QVariant>
#include <QIODevice>
#include <QNetworkCookie>
#include <QNetworkConfiguration>
#include <QSslPreSharedKeyAuthenticator>
#include <QHttpMultiPart>
#include <QNetworkProxyFactory>
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

QtObjectPtr QNetworkAccessManager_ProxyFactory(QtObjectPtr ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->proxyFactory();
}

QtObjectPtr QNetworkAccessManager_NewQNetworkAccessManager(QtObjectPtr parent){
	return new QNetworkAccessManager(static_cast<QObject*>(parent));
}

void QNetworkAccessManager_ConnectAuthenticationRequired(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));;
}

void QNetworkAccessManager_DisconnectAuthenticationRequired(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));;
}

QtObjectPtr QNetworkAccessManager_Cache(QtObjectPtr ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->cache();
}

void QNetworkAccessManager_ClearAccessCache(QtObjectPtr ptr){
	static_cast<QNetworkAccessManager*>(ptr)->clearAccessCache();
}

QtObjectPtr QNetworkAccessManager_CookieJar(QtObjectPtr ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->cookieJar();
}

QtObjectPtr QNetworkAccessManager_DeleteResource(QtObjectPtr ptr, QtObjectPtr request){
	return static_cast<QNetworkAccessManager*>(ptr)->deleteResource(*static_cast<QNetworkRequest*>(request));
}

void QNetworkAccessManager_ConnectEncrypted(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));;
}

void QNetworkAccessManager_DisconnectEncrypted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));;
}

void QNetworkAccessManager_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));;
}

void QNetworkAccessManager_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));;
}

QtObjectPtr QNetworkAccessManager_Get(QtObjectPtr ptr, QtObjectPtr request){
	return static_cast<QNetworkAccessManager*>(ptr)->get(*static_cast<QNetworkRequest*>(request));
}

QtObjectPtr QNetworkAccessManager_Head(QtObjectPtr ptr, QtObjectPtr request){
	return static_cast<QNetworkAccessManager*>(ptr)->head(*static_cast<QNetworkRequest*>(request));
}

int QNetworkAccessManager_NetworkAccessible(QtObjectPtr ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->networkAccessible();
}

void QNetworkAccessManager_ConnectNetworkAccessibleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));;
}

void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));;
}

QtObjectPtr QNetworkAccessManager_Post3(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr multiPart){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

QtObjectPtr QNetworkAccessManager_Post(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

QtObjectPtr QNetworkAccessManager_Post2(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));;
}

QtObjectPtr QNetworkAccessManager_Put2(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr multiPart){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

QtObjectPtr QNetworkAccessManager_Put(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

QtObjectPtr QNetworkAccessManager_Put3(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

QtObjectPtr QNetworkAccessManager_SendCustomRequest(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr verb, QtObjectPtr data){
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), static_cast<QIODevice*>(data));
}

void QNetworkAccessManager_SetCache(QtObjectPtr ptr, QtObjectPtr cache){
	static_cast<QNetworkAccessManager*>(ptr)->setCache(static_cast<QAbstractNetworkCache*>(cache));
}

void QNetworkAccessManager_SetConfiguration(QtObjectPtr ptr, QtObjectPtr config){
	static_cast<QNetworkAccessManager*>(ptr)->setConfiguration(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkAccessManager_SetCookieJar(QtObjectPtr ptr, QtObjectPtr cookieJar){
	static_cast<QNetworkAccessManager*>(ptr)->setCookieJar(static_cast<QNetworkCookieJar*>(cookieJar));
}

void QNetworkAccessManager_SetNetworkAccessible(QtObjectPtr ptr, int accessible){
	static_cast<QNetworkAccessManager*>(ptr)->setNetworkAccessible(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void QNetworkAccessManager_SetProxy(QtObjectPtr ptr, QtObjectPtr proxy){
	static_cast<QNetworkAccessManager*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(proxy));
}

void QNetworkAccessManager_SetProxyFactory(QtObjectPtr ptr, QtObjectPtr factory){
	static_cast<QNetworkAccessManager*>(ptr)->setProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

char* QNetworkAccessManager_SupportedSchemes(QtObjectPtr ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->supportedSchemes().join("|").toUtf8().data();
}

void QNetworkAccessManager_DestroyQNetworkAccessManager(QtObjectPtr ptr){
	static_cast<QNetworkAccessManager*>(ptr)->~QNetworkAccessManager();
}

