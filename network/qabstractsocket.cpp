#include "qabstractsocket.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNetworkProxy>
#include <QString>
#include <QAbstractSocket>
#include "_cgo_export.h"

class MyQAbstractSocket: public QAbstractSocket {
public:
void Signal_Connected(){callbackQAbstractSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQAbstractSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_HostFound(){callbackQAbstractSocketHostFound(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAbstractSocket::SocketState socketState){callbackQAbstractSocketStateChanged(this->objectName().toUtf8().data(), socketState);};
};

void* QAbstractSocket_NewQAbstractSocket(int socketType, void* parent){
	return new QAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QObject*>(parent));
}

void QAbstractSocket_Abort(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->abort();
}

int QAbstractSocket_AtEnd(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->atEnd();
}

int QAbstractSocket_CanReadLine(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->canReadLine();
}

void QAbstractSocket_Close(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->close();
}

void QAbstractSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));;
}

void QAbstractSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));;
}

void QAbstractSocket_DisconnectFromHost(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->disconnectFromHost();
}

void QAbstractSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));;
}

void QAbstractSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));;
}

int QAbstractSocket_Error(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->error();
}

int QAbstractSocket_Flush(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->flush();
}

void QAbstractSocket_ConnectHostFound(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));;
}

void QAbstractSocket_DisconnectHostFound(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));;
}

int QAbstractSocket_IsSequential(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->isSequential();
}

int QAbstractSocket_IsValid(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->isValid();
}

int QAbstractSocket_PauseMode(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->pauseMode();
}

char* QAbstractSocket_PeerName(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->peerName().toUtf8().data();
}

void QAbstractSocket_Resume(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->resume();
}

void QAbstractSocket_SetPauseMode(void* ptr, int pauseMode){
	static_cast<QAbstractSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QAbstractSocket_SetProxy(void* ptr, void* networkProxy){
	static_cast<QAbstractSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QAbstractSocket_SetSocketOption(void* ptr, int option, void* value){
	static_cast<QAbstractSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void* QAbstractSocket_SocketOption(void* ptr, int option){
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

int QAbstractSocket_SocketType(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->socketType();
}

void QAbstractSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));;
}

void QAbstractSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));;
}

int QAbstractSocket_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QAbstractSocket_WaitForConnected(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForConnected(msecs);
}

int QAbstractSocket_WaitForDisconnected(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForDisconnected(msecs);
}

int QAbstractSocket_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForReadyRead(msecs);
}

void QAbstractSocket_DestroyQAbstractSocket(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->~QAbstractSocket();
}

