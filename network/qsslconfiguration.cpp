#include "qsslconfiguration.h"
#include <QSslKey>
#include <QByteArray>
#include <QSslSocket>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslCertificate>
#include <QSslConfiguration>
#include "_cgo_export.h"

class MyQSslConfiguration: public QSslConfiguration {
public:
};

void* QSslConfiguration_NewQSslConfiguration(){
	return new QSslConfiguration();
}

void* QSslConfiguration_NewQSslConfiguration2(void* other){
	return new QSslConfiguration(*static_cast<QSslConfiguration*>(other));
}

int QSslConfiguration_IsNull(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->isNull();
}

void* QSslConfiguration_NextNegotiatedProtocol(void* ptr){
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->nextNegotiatedProtocol());
}

int QSslConfiguration_NextProtocolNegotiationStatus(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->nextProtocolNegotiationStatus();
}

int QSslConfiguration_PeerVerifyDepth(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyDepth();
}

int QSslConfiguration_PeerVerifyMode(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyMode();
}

void* QSslConfiguration_SessionTicket(void* ptr){
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->sessionTicket());
}

int QSslConfiguration_SessionTicketLifeTimeHint(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->sessionTicketLifeTimeHint();
}

void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(void* configuration){
	QSslConfiguration::setDefaultConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void QSslConfiguration_SetLocalCertificate(void* ptr, void* certificate){
	static_cast<QSslConfiguration*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslConfiguration_SetPeerVerifyDepth(void* ptr, int depth){
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslConfiguration_SetPeerVerifyMode(void* ptr, int mode){
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslConfiguration_SetPrivateKey(void* ptr, void* key){
	static_cast<QSslConfiguration*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslConfiguration_SetSessionTicket(void* ptr, void* sessionTicket){
	static_cast<QSslConfiguration*>(ptr)->setSessionTicket(*static_cast<QByteArray*>(sessionTicket));
}

void QSslConfiguration_Swap(void* ptr, void* other){
	static_cast<QSslConfiguration*>(ptr)->swap(*static_cast<QSslConfiguration*>(other));
}

void QSslConfiguration_DestroyQSslConfiguration(void* ptr){
	static_cast<QSslConfiguration*>(ptr)->~QSslConfiguration();
}

