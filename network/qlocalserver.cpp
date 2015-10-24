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

void QLocalServer_SetSocketOptions(QtObjectPtr ptr, int options){
	static_cast<QLocalServer*>(ptr)->setSocketOptions(static_cast<QLocalServer::SocketOption>(options));
}

QtObjectPtr QLocalServer_NewQLocalServer(QtObjectPtr parent){
	return new QLocalServer(static_cast<QObject*>(parent));
}

void QLocalServer_Close(QtObjectPtr ptr){
	static_cast<QLocalServer*>(ptr)->close();
}

char* QLocalServer_ErrorString(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->errorString().toUtf8().data();
}

char* QLocalServer_FullServerName(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalServer_HasPendingConnections(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->hasPendingConnections();
}

int QLocalServer_IsListening(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->isListening();
}

int QLocalServer_Listen(QtObjectPtr ptr, char* name){
	return static_cast<QLocalServer*>(ptr)->listen(QString(name));
}

int QLocalServer_MaxPendingConnections(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->maxPendingConnections();
}

void QLocalServer_ConnectNewConnection(QtObjectPtr ptr){
	QObject::connect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));;
}

void QLocalServer_DisconnectNewConnection(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));;
}

QtObjectPtr QLocalServer_NextPendingConnection(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->nextPendingConnection();
}

int QLocalServer_QLocalServer_RemoveServer(char* name){
	return QLocalServer::removeServer(QString(name));
}

int QLocalServer_ServerError(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->serverError();
}

char* QLocalServer_ServerName(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->serverName().toUtf8().data();
}

void QLocalServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections){
	static_cast<QLocalServer*>(ptr)->setMaxPendingConnections(numConnections);
}

int QLocalServer_SocketOptions(QtObjectPtr ptr){
	return static_cast<QLocalServer*>(ptr)->socketOptions();
}

int QLocalServer_WaitForNewConnection(QtObjectPtr ptr, int msec, int timedOut){
	return static_cast<QLocalServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QLocalServer_DestroyQLocalServer(QtObjectPtr ptr){
	static_cast<QLocalServer*>(ptr)->~QLocalServer();
}

