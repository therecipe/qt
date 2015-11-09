#include "qnetworkproxy.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkRequest>
#include <QByteArray>
#include <QNetworkProxy>
#include "_cgo_export.h"

class MyQNetworkProxy: public QNetworkProxy {
public:
};

void* QNetworkProxy_NewQNetworkProxy(){
	return new QNetworkProxy();
}

void* QNetworkProxy_NewQNetworkProxy3(void* other){
	return new QNetworkProxy(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Capabilities(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->capabilities();
}

int QNetworkProxy_HasRawHeader(void* ptr, void* headerName){
	return static_cast<QNetworkProxy*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkProxy_Header(void* ptr, int header){
	return new QVariant(static_cast<QNetworkProxy*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

char* QNetworkProxy_HostName(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->hostName().toUtf8().data();
}

int QNetworkProxy_IsCachingProxy(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->isCachingProxy();
}

int QNetworkProxy_IsTransparentProxy(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->isTransparentProxy();
}

char* QNetworkProxy_Password(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->password().toUtf8().data();
}

void* QNetworkProxy_RawHeader(void* ptr, void* headerName){
	return new QByteArray(static_cast<QNetworkProxy*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

void QNetworkProxy_QNetworkProxy_SetApplicationProxy(void* networkProxy){
	QNetworkProxy::setApplicationProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QNetworkProxy_SetCapabilities(void* ptr, int capabilities){
	static_cast<QNetworkProxy*>(ptr)->setCapabilities(static_cast<QNetworkProxy::Capability>(capabilities));
}

void QNetworkProxy_SetHeader(void* ptr, int header, void* value){
	static_cast<QNetworkProxy*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkProxy_SetHostName(void* ptr, char* hostName){
	static_cast<QNetworkProxy*>(ptr)->setHostName(QString(hostName));
}

void QNetworkProxy_SetPassword(void* ptr, char* password){
	static_cast<QNetworkProxy*>(ptr)->setPassword(QString(password));
}

void QNetworkProxy_SetRawHeader(void* ptr, void* headerName, void* headerValue){
	static_cast<QNetworkProxy*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkProxy_SetType(void* ptr, int ty){
	static_cast<QNetworkProxy*>(ptr)->setType(static_cast<QNetworkProxy::ProxyType>(ty));
}

void QNetworkProxy_SetUser(void* ptr, char* user){
	static_cast<QNetworkProxy*>(ptr)->setUser(QString(user));
}

void QNetworkProxy_Swap(void* ptr, void* other){
	static_cast<QNetworkProxy*>(ptr)->swap(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Type(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->type();
}

char* QNetworkProxy_User(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->user().toUtf8().data();
}

void QNetworkProxy_DestroyQNetworkProxy(void* ptr){
	static_cast<QNetworkProxy*>(ptr)->~QNetworkProxy();
}

