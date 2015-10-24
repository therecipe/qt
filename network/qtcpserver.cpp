#include "qtcpserver.h"
#include <QAbstractSocket>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNetworkProxy>
#include <QTcpServer>
#include "_cgo_export.h"

class MyQTcpServer: public QTcpServer {
public:
void Signal_AcceptError(QAbstractSocket::SocketError socketError){callbackQTcpServerAcceptError(this->objectName().toUtf8().data(), socketError);};
void Signal_NewConnection(){callbackQTcpServerNewConnection(this->objectName().toUtf8().data());};
};

QtObjectPtr QTcpServer_NewQTcpServer(QtObjectPtr parent){
	return new QTcpServer(static_cast<QObject*>(parent));
}

void QTcpServer_ConnectAcceptError(QtObjectPtr ptr){
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));;
}

void QTcpServer_DisconnectAcceptError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));;
}

void QTcpServer_Close(QtObjectPtr ptr){
	static_cast<QTcpServer*>(ptr)->close();
}

char* QTcpServer_ErrorString(QtObjectPtr ptr){
	return static_cast<QTcpServer*>(ptr)->errorString().toUtf8().data();
}

int QTcpServer_HasPendingConnections(QtObjectPtr ptr){
	return static_cast<QTcpServer*>(ptr)->hasPendingConnections();
}

int QTcpServer_IsListening(QtObjectPtr ptr){
	return static_cast<QTcpServer*>(ptr)->isListening();
}

int QTcpServer_MaxPendingConnections(QtObjectPtr ptr){
	return static_cast<QTcpServer*>(ptr)->maxPendingConnections();
}

void QTcpServer_ConnectNewConnection(QtObjectPtr ptr){
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));;
}

void QTcpServer_DisconnectNewConnection(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));;
}

QtObjectPtr QTcpServer_NextPendingConnection(QtObjectPtr ptr){
	return static_cast<QTcpServer*>(ptr)->nextPendingConnection();
}

void QTcpServer_PauseAccepting(QtObjectPtr ptr){
	static_cast<QTcpServer*>(ptr)->pauseAccepting();
}

void QTcpServer_ResumeAccepting(QtObjectPtr ptr){
	static_cast<QTcpServer*>(ptr)->resumeAccepting();
}

int QTcpServer_ServerError(QtObjectPtr ptr){
	return static_cast<QTcpServer*>(ptr)->serverError();
}

void QTcpServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections){
	static_cast<QTcpServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QTcpServer_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy){
	static_cast<QTcpServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

int QTcpServer_WaitForNewConnection(QtObjectPtr ptr, int msec, int timedOut){
	return static_cast<QTcpServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QTcpServer_DestroyQTcpServer(QtObjectPtr ptr){
	static_cast<QTcpServer*>(ptr)->~QTcpServer();
}

