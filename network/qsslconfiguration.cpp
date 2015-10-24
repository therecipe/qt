#include "qsslconfiguration.h"
#include <QSslCertificate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslKey>
#include <QSslSocket>
#include <QByteArray>
#include <QSslConfiguration>
#include "_cgo_export.h"

class MyQSslConfiguration: public QSslConfiguration {
public:
};

QtObjectPtr QSslConfiguration_NewQSslConfiguration(){
	return new QSslConfiguration();
}

QtObjectPtr QSslConfiguration_NewQSslConfiguration2(QtObjectPtr other){
	return new QSslConfiguration(*static_cast<QSslConfiguration*>(other));
}

int QSslConfiguration_IsNull(QtObjectPtr ptr){
	return static_cast<QSslConfiguration*>(ptr)->isNull();
}

int QSslConfiguration_NextProtocolNegotiationStatus(QtObjectPtr ptr){
	return static_cast<QSslConfiguration*>(ptr)->nextProtocolNegotiationStatus();
}

int QSslConfiguration_PeerVerifyDepth(QtObjectPtr ptr){
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyDepth();
}

int QSslConfiguration_PeerVerifyMode(QtObjectPtr ptr){
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyMode();
}

int QSslConfiguration_SessionTicketLifeTimeHint(QtObjectPtr ptr){
	return static_cast<QSslConfiguration*>(ptr)->sessionTicketLifeTimeHint();
}

void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(QtObjectPtr configuration){
	QSslConfiguration::setDefaultConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void QSslConfiguration_SetLocalCertificate(QtObjectPtr ptr, QtObjectPtr certificate){
	static_cast<QSslConfiguration*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslConfiguration_SetPeerVerifyDepth(QtObjectPtr ptr, int depth){
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslConfiguration_SetPeerVerifyMode(QtObjectPtr ptr, int mode){
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslConfiguration_SetPrivateKey(QtObjectPtr ptr, QtObjectPtr key){
	static_cast<QSslConfiguration*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslConfiguration_SetSessionTicket(QtObjectPtr ptr, QtObjectPtr sessionTicket){
	static_cast<QSslConfiguration*>(ptr)->setSessionTicket(*static_cast<QByteArray*>(sessionTicket));
}

void QSslConfiguration_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QSslConfiguration*>(ptr)->swap(*static_cast<QSslConfiguration*>(other));
}

void QSslConfiguration_DestroyQSslConfiguration(QtObjectPtr ptr){
	static_cast<QSslConfiguration*>(ptr)->~QSslConfiguration();
}

