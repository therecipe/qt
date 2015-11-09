#include "qbluetoothserver.h"
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothServiceInfo>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QBluetoothServer>
#include "_cgo_export.h"

class MyQBluetoothServer: public QBluetoothServer {
public:
void Signal_NewConnection(){callbackQBluetoothServerNewConnection(this->objectName().toUtf8().data());};
};

void* QBluetoothServer_NewQBluetoothServer(int serverType, void* parent){
	return new QBluetoothServer(static_cast<QBluetoothServiceInfo::Protocol>(serverType), static_cast<QObject*>(parent));
}

void QBluetoothServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

void QBluetoothServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

int QBluetoothServer_Error(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->error();
}

int QBluetoothServer_IsListening(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->isListening();
}

int QBluetoothServer_MaxPendingConnections(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->maxPendingConnections();
}

int QBluetoothServer_ServerType(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->serverType();
}

void QBluetoothServer_DestroyQBluetoothServer(void* ptr){
	static_cast<QBluetoothServer*>(ptr)->~QBluetoothServer();
}

void QBluetoothServer_Close(void* ptr){
	static_cast<QBluetoothServer*>(ptr)->close();
}

int QBluetoothServer_HasPendingConnections(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->hasPendingConnections();
}

void* QBluetoothServer_NextPendingConnection(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->nextPendingConnection();
}

void QBluetoothServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QBluetoothServer*>(ptr)->setMaxPendingConnections(numConnections);
}

