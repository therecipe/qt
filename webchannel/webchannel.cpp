#include "qwebchannel.h"
#include <QModelIndex>
#include <QWebChannelAbstractTransport>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWebChannel>
#include "_cgo_export.h"

class MyQWebChannel: public QWebChannel {
public:
void Signal_BlockUpdatesChanged(bool block){callbackQWebChannelBlockUpdatesChanged(this->objectName().toUtf8().data(), block);};
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

#include "qwebchannelabstracttransport.h"
#include <QJsonObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QWebChannel>
#include <QWebChannelAbstractTransport>
#include "_cgo_export.h"

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport {
public:
void Signal_MessageReceived(const QJsonObject & message, QWebChannelAbstractTransport * transport){callbackQWebChannelAbstractTransportMessageReceived(this->objectName().toUtf8().data(), new QJsonObject(message), transport);};
};

void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr){
	QObject::connect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));;
}

void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr){
	QObject::disconnect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));;
}

void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message){
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "sendMessage", Q_ARG(QJsonObject, *static_cast<QJsonObject*>(message)));
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr){
	static_cast<QWebChannelAbstractTransport*>(ptr)->~QWebChannelAbstractTransport();
}

