#include "qwebsocket.h"
#include <QString>
#include <QUrl>
#include <QNetworkProxy>
#include <QAbstractSocket>
#include <QSslConfiguration>
#include <QMaskGenerator>
#include <QVariant>
#include <QModelIndex>
#include <QByteArray>
#include <QMetaObject>
#include <QObject>
#include <QWebSocket>
#include "_cgo_export.h"

class MyQWebSocket: public QWebSocket {
public:
void Signal_AboutToClose(){callbackQWebSocketAboutToClose(this->objectName().toUtf8().data());};
void Signal_Connected(){callbackQWebSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQWebSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_ReadChannelFinished(){callbackQWebSocketReadChannelFinished(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAbstractSocket::SocketState state){callbackQWebSocketStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_TextFrameReceived(const QString & frame, bool isLastFrame){callbackQWebSocketTextFrameReceived(this->objectName().toUtf8().data(), frame.toUtf8().data(), isLastFrame);};
void Signal_TextMessageReceived(const QString & message){callbackQWebSocketTextMessageReceived(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

void QWebSocket_Abort(QtObjectPtr ptr){
	static_cast<QWebSocket*>(ptr)->abort();
}

void QWebSocket_ConnectAboutToClose(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::aboutToClose), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_AboutToClose));;
}

void QWebSocket_DisconnectAboutToClose(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::aboutToClose), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_AboutToClose));;
}

char* QWebSocket_CloseReason(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->closeReason().toUtf8().data();
}

void QWebSocket_ConnectConnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::connected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Connected));;
}

void QWebSocket_DisconnectConnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::connected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Connected));;
}

void QWebSocket_ConnectDisconnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::disconnected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Disconnected));;
}

void QWebSocket_DisconnectDisconnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::disconnected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Disconnected));;
}

int QWebSocket_Error(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->error();
}

char* QWebSocket_ErrorString(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->errorString().toUtf8().data();
}

int QWebSocket_Flush(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->flush();
}

void QWebSocket_IgnoreSslErrors(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "ignoreSslErrors");
}

int QWebSocket_IsValid(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->isValid();
}

QtObjectPtr QWebSocket_MaskGenerator(QtObjectPtr ptr){
	return const_cast<QMaskGenerator*>(static_cast<QWebSocket*>(ptr)->maskGenerator());
}

void QWebSocket_Open(QtObjectPtr ptr, char* url){
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "open", Q_ARG(QUrl, QUrl(QString(url))));
}

char* QWebSocket_Origin(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->origin().toUtf8().data();
}

int QWebSocket_PauseMode(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->pauseMode();
}

char* QWebSocket_PeerName(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->peerName().toUtf8().data();
}

void QWebSocket_Ping(QtObjectPtr ptr, QtObjectPtr payload){
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "ping", Q_ARG(QByteArray, *static_cast<QByteArray*>(payload)));
}

void QWebSocket_ConnectReadChannelFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::readChannelFinished), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_ReadChannelFinished));;
}

void QWebSocket_DisconnectReadChannelFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::readChannelFinished), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_ReadChannelFinished));;
}

char* QWebSocket_RequestUrl(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->requestUrl().toString().toUtf8().data();
}

char* QWebSocket_ResourceName(QtObjectPtr ptr){
	return static_cast<QWebSocket*>(ptr)->resourceName().toUtf8().data();
}

void QWebSocket_Resume(QtObjectPtr ptr){
	static_cast<QWebSocket*>(ptr)->resume();
}

void QWebSocket_SetMaskGenerator(QtObjectPtr ptr, QtObjectPtr maskGenerator){
	static_cast<QWebSocket*>(ptr)->setMaskGenerator(static_cast<QMaskGenerator*>(maskGenerator));
}

void QWebSocket_SetPauseMode(QtObjectPtr ptr, int pauseMode){
	static_cast<QWebSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QWebSocket_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy){
	static_cast<QWebSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QWebSocket_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr sslConfiguration){
	static_cast<QWebSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(sslConfiguration));
}

void QWebSocket_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketState)>(&QWebSocket::stateChanged), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketState)>(&MyQWebSocket::Signal_StateChanged));;
}

void QWebSocket_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketState)>(&QWebSocket::stateChanged), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketState)>(&MyQWebSocket::Signal_StateChanged));;
}

void QWebSocket_ConnectTextFrameReceived(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &, bool)>(&QWebSocket::textFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &, bool)>(&MyQWebSocket::Signal_TextFrameReceived));;
}

void QWebSocket_DisconnectTextFrameReceived(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &, bool)>(&QWebSocket::textFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &, bool)>(&MyQWebSocket::Signal_TextFrameReceived));;
}

void QWebSocket_ConnectTextMessageReceived(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &)>(&QWebSocket::textMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &)>(&MyQWebSocket::Signal_TextMessageReceived));;
}

void QWebSocket_DisconnectTextMessageReceived(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &)>(&QWebSocket::textMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &)>(&MyQWebSocket::Signal_TextMessageReceived));;
}

void QWebSocket_DestroyQWebSocket(QtObjectPtr ptr){
	static_cast<QWebSocket*>(ptr)->~QWebSocket();
}

