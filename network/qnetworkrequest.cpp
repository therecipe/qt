#include "qnetworkrequest.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslConfiguration>
#include <QObject>
#include <QByteArray>
#include <QString>
#include <QNetworkRequest>
#include "_cgo_export.h"

class MyQNetworkRequest: public QNetworkRequest {
public:
};

QtObjectPtr QNetworkRequest_NewQNetworkRequest2(QtObjectPtr other){
	return new QNetworkRequest(*static_cast<QNetworkRequest*>(other));
}

QtObjectPtr QNetworkRequest_NewQNetworkRequest(char* url){
	return new QNetworkRequest(QUrl(QString(url)));
}

char* QNetworkRequest_Attribute(QtObjectPtr ptr, int code, char* defaultValue){
	return static_cast<QNetworkRequest*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code), QVariant(defaultValue)).toString().toUtf8().data();
}

int QNetworkRequest_HasRawHeader(QtObjectPtr ptr, QtObjectPtr headerName){
	return static_cast<QNetworkRequest*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

char* QNetworkRequest_Header(QtObjectPtr ptr, int header){
	return static_cast<QNetworkRequest*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)).toString().toUtf8().data();
}

QtObjectPtr QNetworkRequest_OriginatingObject(QtObjectPtr ptr){
	return static_cast<QNetworkRequest*>(ptr)->originatingObject();
}

int QNetworkRequest_Priority(QtObjectPtr ptr){
	return static_cast<QNetworkRequest*>(ptr)->priority();
}

void QNetworkRequest_SetAttribute(QtObjectPtr ptr, int code, char* value){
	static_cast<QNetworkRequest*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), QVariant(value));
}

void QNetworkRequest_SetHeader(QtObjectPtr ptr, int header, char* value){
	static_cast<QNetworkRequest*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), QVariant(value));
}

void QNetworkRequest_SetOriginatingObject(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QNetworkRequest*>(ptr)->setOriginatingObject(static_cast<QObject*>(object));
}

void QNetworkRequest_SetPriority(QtObjectPtr ptr, int priority){
	static_cast<QNetworkRequest*>(ptr)->setPriority(static_cast<QNetworkRequest::Priority>(priority));
}

void QNetworkRequest_SetRawHeader(QtObjectPtr ptr, QtObjectPtr headerName, QtObjectPtr headerValue){
	static_cast<QNetworkRequest*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkRequest_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr config){
	static_cast<QNetworkRequest*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkRequest_SetUrl(QtObjectPtr ptr, char* url){
	static_cast<QNetworkRequest*>(ptr)->setUrl(QUrl(QString(url)));
}

void QNetworkRequest_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkRequest*>(ptr)->swap(*static_cast<QNetworkRequest*>(other));
}

char* QNetworkRequest_Url(QtObjectPtr ptr){
	return static_cast<QNetworkRequest*>(ptr)->url().toString().toUtf8().data();
}

void QNetworkRequest_DestroyQNetworkRequest(QtObjectPtr ptr){
	static_cast<QNetworkRequest*>(ptr)->~QNetworkRequest();
}

