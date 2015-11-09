#include "qnetworkreply.h"
#include <QByteArray>
#include <QSslPreSharedKeyAuthenticator>
#include <QObject>
#include <QSslConfiguration>
#include <QNetworkRequest>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkReply>
#include "_cgo_export.h"

class MyQNetworkReply: public QNetworkReply {
public:
void Signal_Encrypted(){callbackQNetworkReplyEncrypted(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQNetworkReplyFinished(this->objectName().toUtf8().data());};
void Signal_MetaDataChanged(){callbackQNetworkReplyMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator){callbackQNetworkReplyPreSharedKeyAuthenticationRequired(this->objectName().toUtf8().data(), authenticator);};
};

void QNetworkReply_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "abort");
}

void* QNetworkReply_Attribute(void* ptr, int code){
	return new QVariant(static_cast<QNetworkReply*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code)));
}

void QNetworkReply_Close(void* ptr){
	static_cast<QNetworkReply*>(ptr)->close();
}

void QNetworkReply_ConnectEncrypted(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));;
}

void QNetworkReply_DisconnectEncrypted(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));;
}

int QNetworkReply_Error(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->error();
}

void QNetworkReply_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));;
}

void QNetworkReply_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));;
}

int QNetworkReply_HasRawHeader(void* ptr, void* headerName){
	return static_cast<QNetworkReply*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkReply_Header(void* ptr, int header){
	return new QVariant(static_cast<QNetworkReply*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

void QNetworkReply_IgnoreSslErrors(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "ignoreSslErrors");
}

int QNetworkReply_IsFinished(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->isFinished();
}

int QNetworkReply_IsRunning(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->isRunning();
}

void* QNetworkReply_Manager(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->manager();
}

void QNetworkReply_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));;
}

void QNetworkReply_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));;
}

int QNetworkReply_Operation(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->operation();
}

void QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));;
}

void* QNetworkReply_RawHeader(void* ptr, void* headerName){
	return new QByteArray(static_cast<QNetworkReply*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

void QNetworkReply_SetSslConfiguration(void* ptr, void* config){
	static_cast<QNetworkReply*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkReply_DestroyQNetworkReply(void* ptr){
	static_cast<QNetworkReply*>(ptr)->~QNetworkReply();
}

