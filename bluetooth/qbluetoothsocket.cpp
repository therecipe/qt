#include "qbluetoothsocket.h"
#include <QModelIndex>
#include <QBluetoothServiceInfo>
#include <QBluetoothAddress>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QIODevice>
#include <QBluetoothUuid>
#include <QBluetoothSocket>
#include "_cgo_export.h"

class MyQBluetoothSocket: public QBluetoothSocket {
public:
void Signal_Connected(){callbackQBluetoothSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQBluetoothSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_StateChanged(QBluetoothSocket::SocketState state){callbackQBluetoothSocketStateChanged(this->objectName().toUtf8().data(), state);};
};

void QBluetoothSocket_ConnectConnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_DisconnectConnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_ConnectDisconnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_DisconnectDisconnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void QBluetoothSocket_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

QtObjectPtr QBluetoothSocket_NewQBluetoothSocket(int socketType, QtObjectPtr parent){
	return new QBluetoothSocket(static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QObject*>(parent));
}

QtObjectPtr QBluetoothSocket_NewQBluetoothSocket2(QtObjectPtr parent){
	return new QBluetoothSocket(static_cast<QObject*>(parent));
}

void QBluetoothSocket_Abort(QtObjectPtr ptr){
	static_cast<QBluetoothSocket*>(ptr)->abort();
}

int QBluetoothSocket_CanReadLine(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->canReadLine();
}

void QBluetoothSocket_Close(QtObjectPtr ptr){
	static_cast<QBluetoothSocket*>(ptr)->close();
}

void QBluetoothSocket_ConnectToService2(QtObjectPtr ptr, QtObjectPtr address, QtObjectPtr uuid, int openMode){
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothAddress*>(address), *static_cast<QBluetoothUuid*>(uuid), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_ConnectToService(QtObjectPtr ptr, QtObjectPtr service, int openMode){
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothServiceInfo*>(service), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_DisconnectFromService(QtObjectPtr ptr){
	static_cast<QBluetoothSocket*>(ptr)->disconnectFromService();
}

int QBluetoothSocket_Error(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->error();
}

char* QBluetoothSocket_ErrorString(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothSocket_IsSequential(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->isSequential();
}

char* QBluetoothSocket_LocalName(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->localName().toUtf8().data();
}

char* QBluetoothSocket_PeerName(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->peerName().toUtf8().data();
}

int QBluetoothSocket_SetSocketDescriptor(QtObjectPtr ptr, int socketDescriptor, int socketType, int socketState, int openMode){
	return static_cast<QBluetoothSocket*>(ptr)->setSocketDescriptor(socketDescriptor, static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QBluetoothSocket::SocketState>(socketState), static_cast<QIODevice::OpenModeFlag>(openMode));
}

int QBluetoothSocket_SocketDescriptor(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketDescriptor();
}

int QBluetoothSocket_SocketType(QtObjectPtr ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketType();
}

void QBluetoothSocket_DestroyQBluetoothSocket(QtObjectPtr ptr){
	static_cast<QBluetoothSocket*>(ptr)->~QBluetoothSocket();
}

