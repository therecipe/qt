#include "qlocalserver.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QLocalServer>
#include "_cgo_export.h"

class MyQLocalServer: public QLocalServer {
public:
void Signal_NewConnection(){callbackQLocalServerNewConnection(this->objectName().toUtf8().data());};
};

void QLocalServer_SetSocketOptions(void* ptr, int options){
	static_cast<QLocalServer*>(ptr)->setSocketOptions(static_cast<QLocalServer::SocketOption>(options));
}

void* QLocalServer_NewQLocalServer(void* parent){
	return new QLocalServer(static_cast<QObject*>(parent));
}

void QLocalServer_Close(void* ptr){
	static_cast<QLocalServer*>(ptr)->close();
}

char* QLocalServer_ErrorString(void* ptr){
	return static_cast<QLocalServer*>(ptr)->errorString().toUtf8().data();
}

char* QLocalServer_FullServerName(void* ptr){
	return static_cast<QLocalServer*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalServer_HasPendingConnections(void* ptr){
	return static_cast<QLocalServer*>(ptr)->hasPendingConnections();
}

int QLocalServer_IsListening(void* ptr){
	return static_cast<QLocalServer*>(ptr)->isListening();
}

int QLocalServer_Listen(void* ptr, char* name){
	return static_cast<QLocalServer*>(ptr)->listen(QString(name));
}

int QLocalServer_MaxPendingConnections(void* ptr){
	return static_cast<QLocalServer*>(ptr)->maxPendingConnections();
}

void QLocalServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));;
}

void QLocalServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));;
}

void* QLocalServer_NextPendingConnection(void* ptr){
	return static_cast<QLocalServer*>(ptr)->nextPendingConnection();
}

int QLocalServer_QLocalServer_RemoveServer(char* name){
	return QLocalServer::removeServer(QString(name));
}

int QLocalServer_ServerError(void* ptr){
	return static_cast<QLocalServer*>(ptr)->serverError();
}

char* QLocalServer_ServerName(void* ptr){
	return static_cast<QLocalServer*>(ptr)->serverName().toUtf8().data();
}

void QLocalServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QLocalServer*>(ptr)->setMaxPendingConnections(numConnections);
}

int QLocalServer_SocketOptions(void* ptr){
	return static_cast<QLocalServer*>(ptr)->socketOptions();
}

int QLocalServer_WaitForNewConnection(void* ptr, int msec, int timedOut){
	return static_cast<QLocalServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QLocalServer_DestroyQLocalServer(void* ptr){
	static_cast<QLocalServer*>(ptr)->~QLocalServer();
}

