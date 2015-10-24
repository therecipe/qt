#include "qsslsocket.h"
#include <QString>
#include <QModelIndex>
#include <QSslCertificate>
#include <QSslKey>
#include <QObject>
#include <QSslConfiguration>
#include <QVariant>
#include <QUrl>
#include <QSslPreSharedKeyAuthenticator>
#include <QAbstractSocket>
#include <QMetaObject>
#include <QSslSocket>
#include "_cgo_export.h"

class MyQSslSocket: public QSslSocket {
public:
void Signal_Encrypted(){callbackQSslSocketEncrypted(this->objectName().toUtf8().data());};
void Signal_ModeChanged(QSslSocket::SslMode mode){callbackQSslSocketModeChanged(this->objectName().toUtf8().data(), mode);};
void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator){callbackQSslSocketPreSharedKeyAuthenticationRequired(this->objectName().toUtf8().data(), authenticator);};
};

QtObjectPtr QSslSocket_NewQSslSocket(QtObjectPtr parent){
	return new QSslSocket(static_cast<QObject*>(parent));
}

void QSslSocket_Abort(QtObjectPtr ptr){
	static_cast<QSslSocket*>(ptr)->abort();
}

void QSslSocket_AddCaCertificate(QtObjectPtr ptr, QtObjectPtr certificate){
	static_cast<QSslSocket*>(ptr)->addCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificate(QtObjectPtr certificate){
	QSslSocket::addDefaultCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

int QSslSocket_AtEnd(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->atEnd();
}

int QSslSocket_CanReadLine(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->canReadLine();
}

void QSslSocket_Close(QtObjectPtr ptr){
	static_cast<QSslSocket*>(ptr)->close();
}

void QSslSocket_ConnectEncrypted(QtObjectPtr ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));;
}

void QSslSocket_DisconnectEncrypted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));;
}

int QSslSocket_Flush(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->flush();
}

void QSslSocket_IgnoreSslErrors(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "ignoreSslErrors");
}

int QSslSocket_IsEncrypted(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->isEncrypted();
}

int QSslSocket_Mode(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->mode();
}

void QSslSocket_ConnectModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));;
}

void QSslSocket_DisconnectModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));;
}

int QSslSocket_PeerVerifyDepth(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyDepth();
}

int QSslSocket_PeerVerifyMode(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyMode();
}

char* QSslSocket_PeerVerifyName(QtObjectPtr ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyName().toUtf8().data();
}

void QSslSocket_ConnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));;
}

void QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));;
}

void QSslSocket_Resume(QtObjectPtr ptr){
	static_cast<QSslSocket*>(ptr)->resume();
}

void QSslSocket_SetLocalCertificate(QtObjectPtr ptr, QtObjectPtr certificate){
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_SetPeerVerifyDepth(QtObjectPtr ptr, int depth){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslSocket_SetPeerVerifyMode(QtObjectPtr ptr, int mode){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslSocket_SetPeerVerifyName(QtObjectPtr ptr, char* hostName){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyName(QString(hostName));
}

void QSslSocket_SetPrivateKey(QtObjectPtr ptr, QtObjectPtr key){
	static_cast<QSslSocket*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslSocket_SetSocketOption(QtObjectPtr ptr, int option, char* value){
	static_cast<QSslSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), QVariant(value));
}

void QSslSocket_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr configuration){
	static_cast<QSslSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

char* QSslSocket_SocketOption(QtObjectPtr ptr, int option){
	return static_cast<QSslSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)).toString().toUtf8().data();
}

char* QSslSocket_QSslSocket_SslLibraryBuildVersionString(){
	return QSslSocket::sslLibraryBuildVersionString().toUtf8().data();
}

char* QSslSocket_QSslSocket_SslLibraryVersionString(){
	return QSslSocket::sslLibraryVersionString().toUtf8().data();
}

void QSslSocket_StartClientEncryption(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startClientEncryption");
}

void QSslSocket_StartServerEncryption(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startServerEncryption");
}

int QSslSocket_QSslSocket_SupportsSsl(){
	return QSslSocket::supportsSsl();
}

int QSslSocket_WaitForBytesWritten(QtObjectPtr ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QSslSocket_WaitForConnected(QtObjectPtr ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForConnected(msecs);
}

int QSslSocket_WaitForDisconnected(QtObjectPtr ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForDisconnected(msecs);
}

int QSslSocket_WaitForEncrypted(QtObjectPtr ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForEncrypted(msecs);
}

int QSslSocket_WaitForReadyRead(QtObjectPtr ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForReadyRead(msecs);
}

void QSslSocket_DestroyQSslSocket(QtObjectPtr ptr){
	static_cast<QSslSocket*>(ptr)->~QSslSocket();
}

