// +build !minimal

#define protected public
#define private public

#include "webchannel.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QJsonObject>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWebChannel>
#include <QWebChannelAbstractTransport>

class MyQWebChannel: public QWebChannel
{
public:
	MyQWebChannel(QObject *parent) : QWebChannel(parent) {};
	void Signal_BlockUpdatesChanged(bool block) { callbackQWebChannel_BlockUpdatesChanged(this, this->objectName().toUtf8().data(), block); };
	void connectTo(QWebChannelAbstractTransport * transport) { callbackQWebChannel_ConnectTo(this, this->objectName().toUtf8().data(), transport); };
	void disconnectFrom(QWebChannelAbstractTransport * transport) { callbackQWebChannel_DisconnectFrom(this, this->objectName().toUtf8().data(), transport); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebChannel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebChannel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebChannel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannel_MetaObject(const_cast<MyQWebChannel*>(this), this->objectName().toUtf8().data())); };
};

int QWebChannel_BlockUpdates(void* ptr)
{
	return static_cast<QWebChannel*>(ptr)->blockUpdates();
}

void QWebChannel_SetBlockUpdates(void* ptr, int block)
{
	static_cast<QWebChannel*>(ptr)->setBlockUpdates(block != 0);
}

void* QWebChannel_NewQWebChannel(void* parent)
{
	return new MyQWebChannel(static_cast<QObject*>(parent));
}

void QWebChannel_ConnectBlockUpdatesChanged(void* ptr)
{
	QObject::connect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));
}

void QWebChannel_DisconnectBlockUpdatesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));
}

void QWebChannel_BlockUpdatesChanged(void* ptr, int block)
{
	static_cast<QWebChannel*>(ptr)->blockUpdatesChanged(block != 0);
}

void QWebChannel_ConnectTo(void* ptr, void* transport)
{
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "connectTo", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_DeregisterObject(void* ptr, void* object)
{
	static_cast<QWebChannel*>(ptr)->deregisterObject(static_cast<QObject*>(object));
}

void QWebChannel_DisconnectFrom(void* ptr, void* transport)
{
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "disconnectFrom", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_RegisterObject(void* ptr, char* id, void* object)
{
	static_cast<QWebChannel*>(ptr)->registerObject(QString(id), static_cast<QObject*>(object));
}

void QWebChannel_DestroyQWebChannel(void* ptr)
{
	static_cast<QWebChannel*>(ptr)->~QWebChannel();
}

void QWebChannel_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebChannel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebChannel*>(ptr)->QWebChannel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannel_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebChannel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebChannel*>(ptr)->QWebChannel::childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebChannel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebChannel*>(ptr)->QWebChannel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannel_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebChannel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebChannel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebChannel*>(ptr)->QWebChannel::customEvent(static_cast<QEvent*>(event));
}

void QWebChannel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "deleteLater");
}

void QWebChannel_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebChannel*>(ptr)->QWebChannel::deleteLater();
}

void QWebChannel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebChannel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebChannel*>(ptr)->QWebChannel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebChannel_Event(void* ptr, void* e)
{
	return static_cast<QWebChannel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QWebChannel_EventDefault(void* ptr, void* e)
{
	return static_cast<QWebChannel*>(ptr)->QWebChannel::event(static_cast<QEvent*>(e));
}

int QWebChannel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebChannel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebChannel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebChannel*>(ptr)->QWebChannel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebChannel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebChannel*>(ptr)->metaObject());
}

void* QWebChannel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebChannel*>(ptr)->QWebChannel::metaObject());
}

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport
{
public:
	MyQWebChannelAbstractTransport(QObject *parent) : QWebChannelAbstractTransport(parent) {};
	void Signal_MessageReceived(const QJsonObject & message, QWebChannelAbstractTransport * transport) { callbackQWebChannelAbstractTransport_MessageReceived(this, this->objectName().toUtf8().data(), const_cast<QJsonObject*>(&message), transport); };
	void sendMessage(const QJsonObject & message) { callbackQWebChannelAbstractTransport_SendMessage(this, this->objectName().toUtf8().data(), const_cast<QJsonObject*>(&message)); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannelAbstractTransport_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebChannelAbstractTransport_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannelAbstractTransport_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannelAbstractTransport_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebChannelAbstractTransport_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannelAbstractTransport_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebChannelAbstractTransport_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannelAbstractTransport_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannelAbstractTransport_MetaObject(const_cast<MyQWebChannelAbstractTransport*>(this), this->objectName().toUtf8().data())); };
};

void* QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(void* parent)
{
	return new MyQWebChannelAbstractTransport(static_cast<QObject*>(parent));
}

void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr)
{
	QObject::connect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));
}

void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr)
{
	QObject::disconnect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));
}

void QWebChannelAbstractTransport_MessageReceived(void* ptr, void* message, void* transport)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->messageReceived(*static_cast<QJsonObject*>(message), static_cast<QWebChannelAbstractTransport*>(transport));
}

void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message)
{
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "sendMessage", Q_ARG(QJsonObject, *static_cast<QJsonObject*>(message)));
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->~QWebChannelAbstractTransport();
}

void QWebChannelAbstractTransport_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebChannelAbstractTransport_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannelAbstractTransport_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannelAbstractTransport_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannelAbstractTransport_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannelAbstractTransport_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebChannelAbstractTransport_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::customEvent(static_cast<QEvent*>(event));
}

void QWebChannelAbstractTransport_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "deleteLater");
}

void QWebChannelAbstractTransport_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::deleteLater();
}

void QWebChannelAbstractTransport_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannelAbstractTransport_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebChannelAbstractTransport_Event(void* ptr, void* e)
{
	return static_cast<QWebChannelAbstractTransport*>(ptr)->event(static_cast<QEvent*>(e));
}

int QWebChannelAbstractTransport_EventDefault(void* ptr, void* e)
{
	return static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::event(static_cast<QEvent*>(e));
}

int QWebChannelAbstractTransport_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebChannelAbstractTransport*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebChannelAbstractTransport_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebChannelAbstractTransport_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebChannelAbstractTransport*>(ptr)->metaObject());
}

void* QWebChannelAbstractTransport_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::metaObject());
}

