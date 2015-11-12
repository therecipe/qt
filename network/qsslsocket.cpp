#include "qsslsocket.h"
#include <QSslCertificate>
#include <QSslKey>
#include <QObject>
#include <QVariant>
#include <QSslPreSharedKeyAuthenticator>
#include <QMetaObject>
#include <QAbstractSocket>
#include <QSslConfiguration>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QSslSocket>
#include "_cgo_export.h"

class MyQSslSocket: public QSslSocket {
public:
void Signal_Encrypted(){callbackQSslSocketEncrypted(this->objectName().toUtf8().data());};
void Signal_ModeChanged(QSslSocket::SslMode mode){callbackQSslSocketModeChanged(this->objectName().toUtf8().data(), mode);};
void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator){callbackQSslSocketPreSharedKeyAuthenticationRequired(this->objectName().toUtf8().data(), authenticator);};
};

void* QSslSocket_NewQSslSocket(void* parent){
	return new QSslSocket(static_cast<QObject*>(parent));
}

void QSslSocket_Abort(void* ptr){
	static_cast<QSslSocket*>(ptr)->abort();
}

void QSslSocket_AddCaCertificate(void* ptr, void* certificate){
	static_cast<QSslSocket*>(ptr)->addCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificate(void* certificate){
	QSslSocket::addDefaultCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

int QSslSocket_AtEnd(void* ptr){
	return static_cast<QSslSocket*>(ptr)->atEnd();
}

int QSslSocket_CanReadLine(void* ptr){
	return static_cast<QSslSocket*>(ptr)->canReadLine();
}

void QSslSocket_Close(void* ptr){
	static_cast<QSslSocket*>(ptr)->close();
}

void QSslSocket_ConnectEncrypted(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));;
}

void QSslSocket_DisconnectEncrypted(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));;
}

int QSslSocket_Flush(void* ptr){
	return static_cast<QSslSocket*>(ptr)->flush();
}

void QSslSocket_IgnoreSslErrors(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "ignoreSslErrors");
}

int QSslSocket_IsEncrypted(void* ptr){
	return static_cast<QSslSocket*>(ptr)->isEncrypted();
}

int QSslSocket_Mode(void* ptr){
	return static_cast<QSslSocket*>(ptr)->mode();
}

void QSslSocket_ConnectModeChanged(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));;
}

void QSslSocket_DisconnectModeChanged(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));;
}

int QSslSocket_PeerVerifyDepth(void* ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyDepth();
}

int QSslSocket_PeerVerifyMode(void* ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyMode();
}

char* QSslSocket_PeerVerifyName(void* ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyName().toUtf8().data();
}

void QSslSocket_ConnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));;
}

void QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));;
}

void QSslSocket_Resume(void* ptr){
	static_cast<QSslSocket*>(ptr)->resume();
}

void QSslSocket_SetLocalCertificate(void* ptr, void* certificate){
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_SetPeerVerifyDepth(void* ptr, int depth){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslSocket_SetPeerVerifyMode(void* ptr, int mode){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslSocket_SetPeerVerifyName(void* ptr, char* hostName){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyName(QString(hostName));
}

void QSslSocket_SetPrivateKey(void* ptr, void* key){
	static_cast<QSslSocket*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslSocket_SetSocketOption(void* ptr, int option, void* value){
	static_cast<QSslSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSslConfiguration(void* ptr, void* configuration){
	static_cast<QSslSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void* QSslSocket_SocketOption(void* ptr, int option){
	return new QVariant(static_cast<QSslSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

char* QSslSocket_QSslSocket_SslLibraryBuildVersionString(){
	return QSslSocket::sslLibraryBuildVersionString().toUtf8().data();
}

char* QSslSocket_QSslSocket_SslLibraryVersionString(){
	return QSslSocket::sslLibraryVersionString().toUtf8().data();
}

void QSslSocket_StartClientEncryption(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startClientEncryption");
}

void QSslSocket_StartServerEncryption(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startServerEncryption");
}

int QSslSocket_QSslSocket_SupportsSsl(){
	return QSslSocket::supportsSsl();
}

int QSslSocket_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QSslSocket_WaitForConnected(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForConnected(msecs);
}

int QSslSocket_WaitForDisconnected(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForDisconnected(msecs);
}

int QSslSocket_WaitForEncrypted(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForEncrypted(msecs);
}

int QSslSocket_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForReadyRead(msecs);
}

void QSslSocket_DestroyQSslSocket(void* ptr){
	static_cast<QSslSocket*>(ptr)->~QSslSocket();
}

