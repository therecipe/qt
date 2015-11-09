#include "qnetworkrequest.h"
#include <QObject>
#include <QSslConfiguration>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkRequest>
#include "_cgo_export.h"

class MyQNetworkRequest: public QNetworkRequest {
public:
};

void* QNetworkRequest_NewQNetworkRequest2(void* other){
	return new QNetworkRequest(*static_cast<QNetworkRequest*>(other));
}

void* QNetworkRequest_NewQNetworkRequest(void* url){
	return new QNetworkRequest(*static_cast<QUrl*>(url));
}

void* QNetworkRequest_Attribute(void* ptr, int code, void* defaultValue){
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

int QNetworkRequest_HasRawHeader(void* ptr, void* headerName){
	return static_cast<QNetworkRequest*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkRequest_Header(void* ptr, int header){
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

void* QNetworkRequest_OriginatingObject(void* ptr){
	return static_cast<QNetworkRequest*>(ptr)->originatingObject();
}

int QNetworkRequest_Priority(void* ptr){
	return static_cast<QNetworkRequest*>(ptr)->priority();
}

void* QNetworkRequest_RawHeader(void* ptr, void* headerName){
	return new QByteArray(static_cast<QNetworkRequest*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

void QNetworkRequest_SetAttribute(void* ptr, int code, void* value){
	static_cast<QNetworkRequest*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetHeader(void* ptr, int header, void* value){
	static_cast<QNetworkRequest*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetOriginatingObject(void* ptr, void* object){
	static_cast<QNetworkRequest*>(ptr)->setOriginatingObject(static_cast<QObject*>(object));
}

void QNetworkRequest_SetPriority(void* ptr, int priority){
	static_cast<QNetworkRequest*>(ptr)->setPriority(static_cast<QNetworkRequest::Priority>(priority));
}

void QNetworkRequest_SetRawHeader(void* ptr, void* headerName, void* headerValue){
	static_cast<QNetworkRequest*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkRequest_SetSslConfiguration(void* ptr, void* config){
	static_cast<QNetworkRequest*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkRequest_SetUrl(void* ptr, void* url){
	static_cast<QNetworkRequest*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkRequest_Swap(void* ptr, void* other){
	static_cast<QNetworkRequest*>(ptr)->swap(*static_cast<QNetworkRequest*>(other));
}

void QNetworkRequest_DestroyQNetworkRequest(void* ptr){
	static_cast<QNetworkRequest*>(ptr)->~QNetworkRequest();
}

