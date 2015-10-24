#include "qnetworkreply.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QSslConfiguration>
#include <QNetworkRequest>
#include <QSslPreSharedKeyAuthenticator>
#include <QString>
#include <QObject>
#include <QByteArray>
#include <QNetworkReply>
#include "_cgo_export.h"

class MyQNetworkReply: public QNetworkReply {
public:
void Signal_Encrypted(){callbackQNetworkReplyEncrypted(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQNetworkReplyFinished(this->objectName().toUtf8().data());};
void Signal_MetaDataChanged(){callbackQNetworkReplyMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator){callbackQNetworkReplyPreSharedKeyAuthenticationRequired(this->objectName().toUtf8().data(), authenticator);};
};

void QNetworkReply_Abort(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "abort");
}

char* QNetworkReply_Attribute(QtObjectPtr ptr, int code){
	return static_cast<QNetworkReply*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code)).toString().toUtf8().data();
}

void QNetworkReply_Close(QtObjectPtr ptr){
	static_cast<QNetworkReply*>(ptr)->close();
}

void QNetworkReply_ConnectEncrypted(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));;
}

void QNetworkReply_DisconnectEncrypted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));;
}

int QNetworkReply_Error(QtObjectPtr ptr){
	return static_cast<QNetworkReply*>(ptr)->error();
}

void QNetworkReply_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));;
}

void QNetworkReply_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));;
}

int QNetworkReply_HasRawHeader(QtObjectPtr ptr, QtObjectPtr headerName){
	return static_cast<QNetworkReply*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

char* QNetworkReply_Header(QtObjectPtr ptr, int header){
	return static_cast<QNetworkReply*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)).toString().toUtf8().data();
}

void QNetworkReply_IgnoreSslErrors(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "ignoreSslErrors");
}

int QNetworkReply_IsFinished(QtObjectPtr ptr){
	return static_cast<QNetworkReply*>(ptr)->isFinished();
}

int QNetworkReply_IsRunning(QtObjectPtr ptr){
	return static_cast<QNetworkReply*>(ptr)->isRunning();
}

QtObjectPtr QNetworkReply_Manager(QtObjectPtr ptr){
	return static_cast<QNetworkReply*>(ptr)->manager();
}

void QNetworkReply_ConnectMetaDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));;
}

void QNetworkReply_DisconnectMetaDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));;
}

int QNetworkReply_Operation(QtObjectPtr ptr){
	return static_cast<QNetworkReply*>(ptr)->operation();
}

void QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkReply_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr config){
	static_cast<QNetworkReply*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

char* QNetworkReply_Url(QtObjectPtr ptr){
	return static_cast<QNetworkReply*>(ptr)->url().toString().toUtf8().data();
}

void QNetworkReply_DestroyQNetworkReply(QtObjectPtr ptr){
	static_cast<QNetworkReply*>(ptr)->~QNetworkReply();
}

