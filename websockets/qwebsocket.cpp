#include "qwebsocket.h"
#include <QVariant>
#include <QNetworkProxy>
#include <QMetaObject>
#include <QAbstractSocket>
#include <QByteArray>
#include <QSslConfiguration>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QMaskGenerator>
#include <QObject>
#include <QWebSocket>
#include "_cgo_export.h"

class MyQWebSocket: public QWebSocket {
public:
void Signal_AboutToClose(){callbackQWebSocketAboutToClose(this->objectName().toUtf8().data());};
void Signal_BinaryFrameReceived(const QByteArray & frame, bool isLastFrame){callbackQWebSocketBinaryFrameReceived(this->objectName().toUtf8().data(), new QByteArray(frame), isLastFrame);};
void Signal_BinaryMessageReceived(const QByteArray & message){callbackQWebSocketBinaryMessageReceived(this->objectName().toUtf8().data(), new QByteArray(message));};
void Signal_Connected(){callbackQWebSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQWebSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_ReadChannelFinished(){callbackQWebSocketReadChannelFinished(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAbstractSocket::SocketState state){callbackQWebSocketStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_TextFrameReceived(const QString & frame, bool isLastFrame){callbackQWebSocketTextFrameReceived(this->objectName().toUtf8().data(), frame.toUtf8().data(), isLastFrame);};
void Signal_TextMessageReceived(const QString & message){callbackQWebSocketTextMessageReceived(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

void QWebSocket_Abort(void* ptr){
	static_cast<QWebSocket*>(ptr)->abort();
}

void QWebSocket_ConnectAboutToClose(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::aboutToClose), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_AboutToClose));;
}

void QWebSocket_DisconnectAboutToClose(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::aboutToClose), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_AboutToClose));;
}

void QWebSocket_ConnectBinaryFrameReceived(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &, bool)>(&QWebSocket::binaryFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &, bool)>(&MyQWebSocket::Signal_BinaryFrameReceived));;
}

void QWebSocket_DisconnectBinaryFrameReceived(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &, bool)>(&QWebSocket::binaryFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &, bool)>(&MyQWebSocket::Signal_BinaryFrameReceived));;
}

void QWebSocket_ConnectBinaryMessageReceived(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &)>(&QWebSocket::binaryMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &)>(&MyQWebSocket::Signal_BinaryMessageReceived));;
}

void QWebSocket_DisconnectBinaryMessageReceived(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &)>(&QWebSocket::binaryMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &)>(&MyQWebSocket::Signal_BinaryMessageReceived));;
}

char* QWebSocket_CloseReason(void* ptr){
	return static_cast<QWebSocket*>(ptr)->closeReason().toUtf8().data();
}

void QWebSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::connected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Connected));;
}

void QWebSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::connected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Connected));;
}

void QWebSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::disconnected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Disconnected));;
}

void QWebSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::disconnected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Disconnected));;
}

int QWebSocket_Error(void* ptr){
	return static_cast<QWebSocket*>(ptr)->error();
}

char* QWebSocket_ErrorString(void* ptr){
	return static_cast<QWebSocket*>(ptr)->errorString().toUtf8().data();
}

int QWebSocket_Flush(void* ptr){
	return static_cast<QWebSocket*>(ptr)->flush();
}

void QWebSocket_IgnoreSslErrors(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "ignoreSslErrors");
}

int QWebSocket_IsValid(void* ptr){
	return static_cast<QWebSocket*>(ptr)->isValid();
}

void* QWebSocket_MaskGenerator(void* ptr){
	return const_cast<QMaskGenerator*>(static_cast<QWebSocket*>(ptr)->maskGenerator());
}

void QWebSocket_Open(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "open", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

char* QWebSocket_Origin(void* ptr){
	return static_cast<QWebSocket*>(ptr)->origin().toUtf8().data();
}

int QWebSocket_PauseMode(void* ptr){
	return static_cast<QWebSocket*>(ptr)->pauseMode();
}

char* QWebSocket_PeerName(void* ptr){
	return static_cast<QWebSocket*>(ptr)->peerName().toUtf8().data();
}

void QWebSocket_Ping(void* ptr, void* payload){
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "ping", Q_ARG(QByteArray, *static_cast<QByteArray*>(payload)));
}

void QWebSocket_ConnectReadChannelFinished(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::readChannelFinished), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_ReadChannelFinished));;
}

void QWebSocket_DisconnectReadChannelFinished(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::readChannelFinished), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_ReadChannelFinished));;
}

char* QWebSocket_ResourceName(void* ptr){
	return static_cast<QWebSocket*>(ptr)->resourceName().toUtf8().data();
}

void QWebSocket_Resume(void* ptr){
	static_cast<QWebSocket*>(ptr)->resume();
}

void QWebSocket_SetMaskGenerator(void* ptr, void* maskGenerator){
	static_cast<QWebSocket*>(ptr)->setMaskGenerator(static_cast<QMaskGenerator*>(maskGenerator));
}

void QWebSocket_SetPauseMode(void* ptr, int pauseMode){
	static_cast<QWebSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QWebSocket_SetProxy(void* ptr, void* networkProxy){
	static_cast<QWebSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QWebSocket_SetSslConfiguration(void* ptr, void* sslConfiguration){
	static_cast<QWebSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(sslConfiguration));
}

void QWebSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketState)>(&QWebSocket::stateChanged), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketState)>(&MyQWebSocket::Signal_StateChanged));;
}

void QWebSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketState)>(&QWebSocket::stateChanged), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketState)>(&MyQWebSocket::Signal_StateChanged));;
}

void QWebSocket_ConnectTextFrameReceived(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &, bool)>(&QWebSocket::textFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &, bool)>(&MyQWebSocket::Signal_TextFrameReceived));;
}

void QWebSocket_DisconnectTextFrameReceived(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &, bool)>(&QWebSocket::textFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &, bool)>(&MyQWebSocket::Signal_TextFrameReceived));;
}

void QWebSocket_ConnectTextMessageReceived(void* ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &)>(&QWebSocket::textMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &)>(&MyQWebSocket::Signal_TextMessageReceived));;
}

void QWebSocket_DisconnectTextMessageReceived(void* ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &)>(&QWebSocket::textMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &)>(&MyQWebSocket::Signal_TextMessageReceived));;
}

void QWebSocket_DestroyQWebSocket(void* ptr){
	static_cast<QWebSocket*>(ptr)->~QWebSocket();
}

