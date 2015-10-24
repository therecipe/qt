#include "qabstractsocket.h"
#include <QNetworkProxy>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractSocket>
#include "_cgo_export.h"

class MyQAbstractSocket: public QAbstractSocket {
public:
void Signal_Connected(){callbackQAbstractSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQAbstractSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_HostFound(){callbackQAbstractSocketHostFound(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAbstractSocket::SocketState socketState){callbackQAbstractSocketStateChanged(this->objectName().toUtf8().data(), socketState);};
};

QtObjectPtr QAbstractSocket_NewQAbstractSocket(int socketType, QtObjectPtr parent){
	return new QAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QObject*>(parent));
}

void QAbstractSocket_Abort(QtObjectPtr ptr){
	static_cast<QAbstractSocket*>(ptr)->abort();
}

int QAbstractSocket_AtEnd(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->atEnd();
}

int QAbstractSocket_CanReadLine(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->canReadLine();
}

void QAbstractSocket_Close(QtObjectPtr ptr){
	static_cast<QAbstractSocket*>(ptr)->close();
}

void QAbstractSocket_ConnectConnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));;
}

void QAbstractSocket_DisconnectConnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));;
}

void QAbstractSocket_DisconnectFromHost(QtObjectPtr ptr){
	static_cast<QAbstractSocket*>(ptr)->disconnectFromHost();
}

void QAbstractSocket_ConnectDisconnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));;
}

void QAbstractSocket_DisconnectDisconnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));;
}

int QAbstractSocket_Error(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->error();
}

int QAbstractSocket_Flush(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->flush();
}

void QAbstractSocket_ConnectHostFound(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));;
}

void QAbstractSocket_DisconnectHostFound(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));;
}

int QAbstractSocket_IsSequential(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->isSequential();
}

int QAbstractSocket_IsValid(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->isValid();
}

int QAbstractSocket_PauseMode(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->pauseMode();
}

char* QAbstractSocket_PeerName(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->peerName().toUtf8().data();
}

void QAbstractSocket_Resume(QtObjectPtr ptr){
	static_cast<QAbstractSocket*>(ptr)->resume();
}

void QAbstractSocket_SetPauseMode(QtObjectPtr ptr, int pauseMode){
	static_cast<QAbstractSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QAbstractSocket_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy){
	static_cast<QAbstractSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QAbstractSocket_SetSocketOption(QtObjectPtr ptr, int option, char* value){
	static_cast<QAbstractSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), QVariant(value));
}

char* QAbstractSocket_SocketOption(QtObjectPtr ptr, int option){
	return static_cast<QAbstractSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)).toString().toUtf8().data();
}

int QAbstractSocket_SocketType(QtObjectPtr ptr){
	return static_cast<QAbstractSocket*>(ptr)->socketType();
}

void QAbstractSocket_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));;
}

void QAbstractSocket_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));;
}

int QAbstractSocket_WaitForBytesWritten(QtObjectPtr ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QAbstractSocket_WaitForConnected(QtObjectPtr ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForConnected(msecs);
}

int QAbstractSocket_WaitForDisconnected(QtObjectPtr ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForDisconnected(msecs);
}

int QAbstractSocket_WaitForReadyRead(QtObjectPtr ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForReadyRead(msecs);
}

void QAbstractSocket_DestroyQAbstractSocket(QtObjectPtr ptr){
	static_cast<QAbstractSocket*>(ptr)->~QAbstractSocket();
}

