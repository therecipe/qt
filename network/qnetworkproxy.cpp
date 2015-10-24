#include "qnetworkproxy.h"
#include <QNetworkRequest>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkProxy>
#include "_cgo_export.h"

class MyQNetworkProxy: public QNetworkProxy {
public:
};

QtObjectPtr QNetworkProxy_NewQNetworkProxy(){
	return new QNetworkProxy();
}

QtObjectPtr QNetworkProxy_NewQNetworkProxy3(QtObjectPtr other){
	return new QNetworkProxy(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Capabilities(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->capabilities();
}

int QNetworkProxy_HasRawHeader(QtObjectPtr ptr, QtObjectPtr headerName){
	return static_cast<QNetworkProxy*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

char* QNetworkProxy_Header(QtObjectPtr ptr, int header){
	return static_cast<QNetworkProxy*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)).toString().toUtf8().data();
}

char* QNetworkProxy_HostName(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->hostName().toUtf8().data();
}

int QNetworkProxy_IsCachingProxy(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->isCachingProxy();
}

int QNetworkProxy_IsTransparentProxy(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->isTransparentProxy();
}

char* QNetworkProxy_Password(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->password().toUtf8().data();
}

void QNetworkProxy_QNetworkProxy_SetApplicationProxy(QtObjectPtr networkProxy){
	QNetworkProxy::setApplicationProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QNetworkProxy_SetCapabilities(QtObjectPtr ptr, int capabilities){
	static_cast<QNetworkProxy*>(ptr)->setCapabilities(static_cast<QNetworkProxy::Capability>(capabilities));
}

void QNetworkProxy_SetHeader(QtObjectPtr ptr, int header, char* value){
	static_cast<QNetworkProxy*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), QVariant(value));
}

void QNetworkProxy_SetHostName(QtObjectPtr ptr, char* hostName){
	static_cast<QNetworkProxy*>(ptr)->setHostName(QString(hostName));
}

void QNetworkProxy_SetPassword(QtObjectPtr ptr, char* password){
	static_cast<QNetworkProxy*>(ptr)->setPassword(QString(password));
}

void QNetworkProxy_SetRawHeader(QtObjectPtr ptr, QtObjectPtr headerName, QtObjectPtr headerValue){
	static_cast<QNetworkProxy*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkProxy_SetType(QtObjectPtr ptr, int ty){
	static_cast<QNetworkProxy*>(ptr)->setType(static_cast<QNetworkProxy::ProxyType>(ty));
}

void QNetworkProxy_SetUser(QtObjectPtr ptr, char* user){
	static_cast<QNetworkProxy*>(ptr)->setUser(QString(user));
}

void QNetworkProxy_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkProxy*>(ptr)->swap(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Type(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->type();
}

char* QNetworkProxy_User(QtObjectPtr ptr){
	return static_cast<QNetworkProxy*>(ptr)->user().toUtf8().data();
}

void QNetworkProxy_DestroyQNetworkProxy(QtObjectPtr ptr){
	static_cast<QNetworkProxy*>(ptr)->~QNetworkProxy();
}

