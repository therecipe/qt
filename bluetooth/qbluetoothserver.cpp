#include "qbluetoothserver.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothServiceInfo>
#include <QBluetoothServer>
#include "_cgo_export.h"

class MyQBluetoothServer: public QBluetoothServer {
public:
void Signal_NewConnection(){callbackQBluetoothServerNewConnection(this->objectName().toUtf8().data());};
};

QtObjectPtr QBluetoothServer_NewQBluetoothServer(int serverType, QtObjectPtr parent){
	return new QBluetoothServer(static_cast<QBluetoothServiceInfo::Protocol>(serverType), static_cast<QObject*>(parent));
}

void QBluetoothServer_ConnectNewConnection(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

void QBluetoothServer_DisconnectNewConnection(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

int QBluetoothServer_Error(QtObjectPtr ptr){
	return static_cast<QBluetoothServer*>(ptr)->error();
}

int QBluetoothServer_IsListening(QtObjectPtr ptr){
	return static_cast<QBluetoothServer*>(ptr)->isListening();
}

int QBluetoothServer_MaxPendingConnections(QtObjectPtr ptr){
	return static_cast<QBluetoothServer*>(ptr)->maxPendingConnections();
}

int QBluetoothServer_ServerType(QtObjectPtr ptr){
	return static_cast<QBluetoothServer*>(ptr)->serverType();
}

void QBluetoothServer_DestroyQBluetoothServer(QtObjectPtr ptr){
	static_cast<QBluetoothServer*>(ptr)->~QBluetoothServer();
}

void QBluetoothServer_Close(QtObjectPtr ptr){
	static_cast<QBluetoothServer*>(ptr)->close();
}

int QBluetoothServer_HasPendingConnections(QtObjectPtr ptr){
	return static_cast<QBluetoothServer*>(ptr)->hasPendingConnections();
}

QtObjectPtr QBluetoothServer_NextPendingConnection(QtObjectPtr ptr){
	return static_cast<QBluetoothServer*>(ptr)->nextPendingConnection();
}

void QBluetoothServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections){
	static_cast<QBluetoothServer*>(ptr)->setMaxPendingConnections(numConnections);
}

