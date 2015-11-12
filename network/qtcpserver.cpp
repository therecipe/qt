#include "qtcpserver.h"
#include <QModelIndex>
#include <QNetworkProxy>
#include <QAbstractSocket>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTcpServer>
#include "_cgo_export.h"

class MyQTcpServer: public QTcpServer {
public:
void Signal_AcceptError(QAbstractSocket::SocketError socketError){callbackQTcpServerAcceptError(this->objectName().toUtf8().data(), socketError);};
void Signal_NewConnection(){callbackQTcpServerNewConnection(this->objectName().toUtf8().data());};
};

void* QTcpServer_NewQTcpServer(void* parent){
	return new QTcpServer(static_cast<QObject*>(parent));
}

void QTcpServer_ConnectAcceptError(void* ptr){
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));;
}

void QTcpServer_DisconnectAcceptError(void* ptr){
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));;
}

void QTcpServer_Close(void* ptr){
	static_cast<QTcpServer*>(ptr)->close();
}

char* QTcpServer_ErrorString(void* ptr){
	return static_cast<QTcpServer*>(ptr)->errorString().toUtf8().data();
}

int QTcpServer_HasPendingConnections(void* ptr){
	return static_cast<QTcpServer*>(ptr)->hasPendingConnections();
}

int QTcpServer_IsListening(void* ptr){
	return static_cast<QTcpServer*>(ptr)->isListening();
}

int QTcpServer_MaxPendingConnections(void* ptr){
	return static_cast<QTcpServer*>(ptr)->maxPendingConnections();
}

void QTcpServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));;
}

void QTcpServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));;
}

void* QTcpServer_NextPendingConnection(void* ptr){
	return static_cast<QTcpServer*>(ptr)->nextPendingConnection();
}

void QTcpServer_PauseAccepting(void* ptr){
	static_cast<QTcpServer*>(ptr)->pauseAccepting();
}

void QTcpServer_ResumeAccepting(void* ptr){
	static_cast<QTcpServer*>(ptr)->resumeAccepting();
}

int QTcpServer_ServerError(void* ptr){
	return static_cast<QTcpServer*>(ptr)->serverError();
}

void QTcpServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QTcpServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QTcpServer_SetProxy(void* ptr, void* networkProxy){
	static_cast<QTcpServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

int QTcpServer_WaitForNewConnection(void* ptr, int msec, int timedOut){
	return static_cast<QTcpServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QTcpServer_DestroyQTcpServer(void* ptr){
	static_cast<QTcpServer*>(ptr)->~QTcpServer();
}

