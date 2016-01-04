#define protected public

#include "webchannel.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QJsonObject>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWebChannel>
#include <QWebChannelAbstractTransport>

class MyQWebChannel: public QWebChannel {
public:
	void Signal_BlockUpdatesChanged(bool block) { callbackQWebChannelBlockUpdatesChanged(this, this->objectName().toUtf8().data(), block); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebChannelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQWebChannelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QWebChannel_BlockUpdates(void* ptr){
	return static_cast<QWebChannel*>(ptr)->blockUpdates();
}

void QWebChannel_SetBlockUpdates(void* ptr, int block){
	static_cast<QWebChannel*>(ptr)->setBlockUpdates(block != 0);
}

void* QWebChannel_NewQWebChannel(void* parent){
	return new QWebChannel(static_cast<QObject*>(parent));
}

void QWebChannel_ConnectBlockUpdatesChanged(void* ptr){
	QObject::connect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));;
}

void QWebChannel_DisconnectBlockUpdatesChanged(void* ptr){
	QObject::disconnect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));;
}

void QWebChannel_BlockUpdatesChanged(void* ptr, int block){
	static_cast<QWebChannel*>(ptr)->blockUpdatesChanged(block != 0);
}

void QWebChannel_ConnectTo(void* ptr, void* transport){
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "connectTo", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_DeregisterObject(void* ptr, void* object){
	static_cast<QWebChannel*>(ptr)->deregisterObject(static_cast<QObject*>(object));
}

void QWebChannel_DisconnectFrom(void* ptr, void* transport){
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "disconnectFrom", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_RegisterObject(void* ptr, char* id, void* object){
	static_cast<QWebChannel*>(ptr)->registerObject(QString(id), static_cast<QObject*>(object));
}

void QWebChannel_DestroyQWebChannel(void* ptr){
	static_cast<QWebChannel*>(ptr)->~QWebChannel();
}

void QWebChannel_TimerEvent(void* ptr, void* event){
	static_cast<MyQWebChannel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannel_TimerEventDefault(void* ptr, void* event){
	static_cast<QWebChannel*>(ptr)->QWebChannel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannel_ChildEvent(void* ptr, void* event){
	static_cast<MyQWebChannel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannel_ChildEventDefault(void* ptr, void* event){
	static_cast<QWebChannel*>(ptr)->QWebChannel::childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannel_CustomEvent(void* ptr, void* event){
	static_cast<MyQWebChannel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebChannel_CustomEventDefault(void* ptr, void* event){
	static_cast<QWebChannel*>(ptr)->QWebChannel::customEvent(static_cast<QEvent*>(event));
}

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport {
public:
	void Signal_MessageReceived(const QJsonObject & message, QWebChannelAbstractTransport * transport) { callbackQWebChannelAbstractTransportMessageReceived(this, this->objectName().toUtf8().data(), new QJsonObject(message), transport); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannelAbstractTransportTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebChannelAbstractTransportChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQWebChannelAbstractTransportCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr){
	QObject::connect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));;
}

void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr){
	QObject::disconnect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));;
}

void QWebChannelAbstractTransport_MessageReceived(void* ptr, void* message, void* transport){
	static_cast<QWebChannelAbstractTransport*>(ptr)->messageReceived(*static_cast<QJsonObject*>(message), static_cast<QWebChannelAbstractTransport*>(transport));
}

void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message){
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "sendMessage", Q_ARG(QJsonObject, *static_cast<QJsonObject*>(message)));
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr){
	static_cast<QWebChannelAbstractTransport*>(ptr)->~QWebChannelAbstractTransport();
}

void QWebChannelAbstractTransport_TimerEvent(void* ptr, void* event){
	static_cast<MyQWebChannelAbstractTransport*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event){
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannelAbstractTransport_ChildEvent(void* ptr, void* event){
	static_cast<MyQWebChannelAbstractTransport*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannelAbstractTransport_ChildEventDefault(void* ptr, void* event){
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannelAbstractTransport_CustomEvent(void* ptr, void* event){
	static_cast<MyQWebChannelAbstractTransport*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebChannelAbstractTransport_CustomEventDefault(void* ptr, void* event){
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::customEvent(static_cast<QEvent*>(event));
}

