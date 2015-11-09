#include "qlocalsocket.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QLocalSocket>
#include "_cgo_export.h"

class MyQLocalSocket: public QLocalSocket {
public:
void Signal_Connected(){callbackQLocalSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQLocalSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_StateChanged(QLocalSocket::LocalSocketState socketState){callbackQLocalSocketStateChanged(this->objectName().toUtf8().data(), socketState);};
};

int QLocalSocket_Open(void* ptr, int openMode){
	return static_cast<QLocalSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QLocalSocket_NewQLocalSocket(void* parent){
	return new QLocalSocket(static_cast<QObject*>(parent));
}

void QLocalSocket_ConnectToServer2(void* ptr, char* name, int openMode){
	static_cast<QLocalSocket*>(ptr)->connectToServer(QString(name), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Connected));;
}

void QLocalSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Connected));;
}

void QLocalSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Disconnected));;
}

void QLocalSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Disconnected));;
}

char* QLocalSocket_FullServerName(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalSocket_IsSequential(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->isSequential();
}

char* QLocalSocket_ServerName(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->serverName().toUtf8().data();
}

void QLocalSocket_SetServerName(void* ptr, char* name){
	static_cast<QLocalSocket*>(ptr)->setServerName(QString(name));
}

void QLocalSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));;
}

void QLocalSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));;
}

void QLocalSocket_DestroyQLocalSocket(void* ptr){
	static_cast<QLocalSocket*>(ptr)->~QLocalSocket();
}

void QLocalSocket_Abort(void* ptr){
	static_cast<QLocalSocket*>(ptr)->abort();
}

int QLocalSocket_CanReadLine(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->canReadLine();
}

void QLocalSocket_Close(void* ptr){
	static_cast<QLocalSocket*>(ptr)->close();
}

void QLocalSocket_ConnectToServer(void* ptr, int openMode){
	static_cast<QLocalSocket*>(ptr)->connectToServer(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_DisconnectFromServer(void* ptr){
	static_cast<QLocalSocket*>(ptr)->disconnectFromServer();
}

int QLocalSocket_Error(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->error();
}

int QLocalSocket_Flush(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->flush();
}

int QLocalSocket_IsValid(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->isValid();
}

int QLocalSocket_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QLocalSocket_WaitForConnected(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForConnected(msecs);
}

int QLocalSocket_WaitForDisconnected(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForDisconnected(msecs);
}

int QLocalSocket_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForReadyRead(msecs);
}

