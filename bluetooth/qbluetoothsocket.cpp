#include "qbluetoothsocket.h"
#include <QObject>
#include <QBluetoothServiceInfo>
#include <QBluetoothAddress>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothUuid>
#include <QVariant>
#include <QIODevice>
#include <QBluetoothSocket>
#include "_cgo_export.h"

class MyQBluetoothSocket: public QBluetoothSocket {
public:
void Signal_Connected(){callbackQBluetoothSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQBluetoothSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_StateChanged(QBluetoothSocket::SocketState state){callbackQBluetoothSocketStateChanged(this->objectName().toUtf8().data(), state);};
};

void QBluetoothSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void QBluetoothSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void* QBluetoothSocket_NewQBluetoothSocket(int socketType, void* parent){
	return new QBluetoothSocket(static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QObject*>(parent));
}

void* QBluetoothSocket_NewQBluetoothSocket2(void* parent){
	return new QBluetoothSocket(static_cast<QObject*>(parent));
}

void QBluetoothSocket_Abort(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->abort();
}

int QBluetoothSocket_CanReadLine(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->canReadLine();
}

void QBluetoothSocket_Close(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->close();
}

void QBluetoothSocket_ConnectToService2(void* ptr, void* address, void* uuid, int openMode){
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothAddress*>(address), *static_cast<QBluetoothUuid*>(uuid), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_ConnectToService(void* ptr, void* service, int openMode){
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothServiceInfo*>(service), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_DisconnectFromService(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->disconnectFromService();
}

int QBluetoothSocket_Error(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->error();
}

char* QBluetoothSocket_ErrorString(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothSocket_IsSequential(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->isSequential();
}

char* QBluetoothSocket_LocalName(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->localName().toUtf8().data();
}

char* QBluetoothSocket_PeerName(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->peerName().toUtf8().data();
}

int QBluetoothSocket_SetSocketDescriptor(void* ptr, int socketDescriptor, int socketType, int socketState, int openMode){
	return static_cast<QBluetoothSocket*>(ptr)->setSocketDescriptor(socketDescriptor, static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QBluetoothSocket::SocketState>(socketState), static_cast<QIODevice::OpenModeFlag>(openMode));
}

int QBluetoothSocket_SocketDescriptor(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketDescriptor();
}

int QBluetoothSocket_SocketType(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketType();
}

void QBluetoothSocket_DestroyQBluetoothSocket(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->~QBluetoothSocket();
}

