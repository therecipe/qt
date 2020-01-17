// +build !minimal

#define protected public
#define private public

#include "webchannel.h"
#include "_cgo_export.h"

#include <QAudioSystemPlugin>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QJsonObject>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQmlWebChannel>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QWebChannel>
#include <QWebChannelAbstractTransport>
#include <QWidget>
#include <QWindow>

class MyQQmlWebChannel: public QQmlWebChannel
{
public:
	void Signal_BlockUpdatesChanged(bool block) { callbackQWebChannel_BlockUpdatesChanged(this, block); };
	void connectTo(QWebChannelAbstractTransport * transport) { callbackQWebChannel_ConnectTo(this, transport); };
	void disconnectFrom(QWebChannelAbstractTransport * transport) { callbackQWebChannel_DisconnectFrom(this, transport); };
	void childEvent(QChildEvent * event) { callbackQWebChannel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannel_CustomEvent(this, event); };
	void deleteLater() { callbackQWebChannel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebChannel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebChannel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWebChannel_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWebChannel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQmlWebChannel*)

int QQmlWebChannel_QQmlWebChannel_QRegisterMetaType(){qRegisterMetaType<QQmlWebChannel*>(); return qRegisterMetaType<MyQQmlWebChannel*>();}

void* QQmlWebChannel___registerObjects_objects_atList(void* ptr, struct QtWebChannel_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QQmlWebChannel___registerObjects_objects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

class MyQWebChannel: public QWebChannel
{
public:
	MyQWebChannel(QObject *parent = Q_NULLPTR) : QWebChannel(parent) {QWebChannel_QWebChannel_QRegisterMetaType();};
	void Signal_BlockUpdatesChanged(bool block) { callbackQWebChannel_BlockUpdatesChanged(this, block); };
	void connectTo(QWebChannelAbstractTransport * transport) { callbackQWebChannel_ConnectTo(this, transport); };
	void disconnectFrom(QWebChannelAbstractTransport * transport) { callbackQWebChannel_DisconnectFrom(this, transport); };
	 ~MyQWebChannel() { callbackQWebChannel_DestroyQWebChannel(this); };
	void childEvent(QChildEvent * event) { callbackQWebChannel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannel_CustomEvent(this, event); };
	void deleteLater() { callbackQWebChannel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebChannel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebChannel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWebChannel_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWebChannel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QWebChannel*)
Q_DECLARE_METATYPE(MyQWebChannel*)

int QWebChannel_QWebChannel_QRegisterMetaType(){qRegisterMetaType<QWebChannel*>(); return qRegisterMetaType<MyQWebChannel*>();}

void* QWebChannel_NewQWebChannel(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebChannel(static_cast<QObject*>(parent));
	}
}

char QWebChannel_BlockUpdates(void* ptr)
{
	return static_cast<QWebChannel*>(ptr)->blockUpdates();
}

void QWebChannel_ConnectBlockUpdatesChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged), static_cast<Qt::ConnectionType>(t));
}

void QWebChannel_DisconnectBlockUpdatesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));
}

void QWebChannel_BlockUpdatesChanged(void* ptr, char block)
{
	static_cast<QWebChannel*>(ptr)->blockUpdatesChanged(block != 0);
}

void QWebChannel_ConnectTo(void* ptr, void* transport)
{
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "connectTo", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_ConnectToDefault(void* ptr, void* transport)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::connectTo(static_cast<QWebChannelAbstractTransport*>(transport));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::connectTo(static_cast<QWebChannelAbstractTransport*>(transport));
	}
}

void QWebChannel_DeregisterObject(void* ptr, void* object)
{
	static_cast<QWebChannel*>(ptr)->deregisterObject(static_cast<QObject*>(object));
}

void QWebChannel_DisconnectFrom(void* ptr, void* transport)
{
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "disconnectFrom", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_DisconnectFromDefault(void* ptr, void* transport)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::disconnectFrom(static_cast<QWebChannelAbstractTransport*>(transport));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::disconnectFrom(static_cast<QWebChannelAbstractTransport*>(transport));
	}
}

void QWebChannel_RegisterObject(void* ptr, struct QtWebChannel_PackedString id, void* object)
{
	static_cast<QWebChannel*>(ptr)->registerObject(QString::fromUtf8(id.data, id.len), static_cast<QObject*>(object));
}

void QWebChannel_RegisterObjects(void* ptr, void* objects)
{
	static_cast<QWebChannel*>(ptr)->registerObjects(*static_cast<QHash<QString, QObject *>*>(objects));
}

struct QtWebChannel_PackedList QWebChannel_RegisteredObjects(void* ptr)
{
	return ({ QHash<QString, QObject *>* tmpValue84aaaf = new QHash<QString, QObject *>(static_cast<QWebChannel*>(ptr)->registeredObjects()); QtWebChannel_PackedList { tmpValue84aaaf, tmpValue84aaaf->size() }; });
}

void QWebChannel_SetBlockUpdates(void* ptr, char block)
{
	static_cast<QWebChannel*>(ptr)->setBlockUpdates(block != 0);
}

void QWebChannel_DestroyQWebChannel(void* ptr)
{
	static_cast<QWebChannel*>(ptr)->~QWebChannel();
}

void QWebChannel_DestroyQWebChannelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebChannel___registerObjects_objects_atList(void* ptr, struct QtWebChannel_PackedString v, int i)
{
	return ({ QObject * tmp = static_cast<QHash<QString, QObject *>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QObject *>*>(ptr)->size()-1) { static_cast<QHash<QString, QObject *>*>(ptr)->~QHash(); free(ptr); }; tmp; });
}

void QWebChannel___registerObjects_objects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i)
{
	static_cast<QHash<QString, QObject *>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), static_cast<QObject*>(i));
}

void* QWebChannel___registerObjects_objects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QObject *>();
}

struct QtWebChannel_PackedList QWebChannel___registerObjects_objects_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue6b7fc8 = new QList<QString>(static_cast<QHash<QString, QObject *>*>(ptr)->keys()); QtWebChannel_PackedList { tmpValue6b7fc8, tmpValue6b7fc8->size() }; });
}

void* QWebChannel___registeredObjects_atList(void* ptr, struct QtWebChannel_PackedString v, int i)
{
	return ({ QObject * tmp = static_cast<QHash<QString, QObject *>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QObject *>*>(ptr)->size()-1) { static_cast<QHash<QString, QObject *>*>(ptr)->~QHash(); free(ptr); }; tmp; });
}

void QWebChannel___registeredObjects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i)
{
	static_cast<QHash<QString, QObject *>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), static_cast<QObject*>(i));
}

void* QWebChannel___registeredObjects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QObject *>();
}

struct QtWebChannel_PackedList QWebChannel___registeredObjects_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue6b7fc8 = new QList<QString>(static_cast<QHash<QString, QObject *>*>(ptr)->keys()); QtWebChannel_PackedList { tmpValue6b7fc8, tmpValue6b7fc8->size() }; });
}

struct QtWebChannel_PackedString QWebChannel_____registerObjects_objects_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtWebChannel_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QWebChannel_____registerObjects_objects_keyList_setList(void* ptr, struct QtWebChannel_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebChannel_____registerObjects_objects_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtWebChannel_PackedString QWebChannel_____registeredObjects_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtWebChannel_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QWebChannel_____registeredObjects_keyList_setList(void* ptr, struct QtWebChannel_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebChannel_____registeredObjects_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QWebChannel___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebChannel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWebChannel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebChannel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebChannel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebChannel___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebChannel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebChannel___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebChannel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QWebChannel_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QWebChannel_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QWebChannel_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::customEvent(static_cast<QEvent*>(event));
	}
}

void QWebChannel_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::deleteLater();
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::deleteLater();
	}
}

void QWebChannel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QWebChannel_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QWebChannel*>(ptr)->QWebChannel::event(static_cast<QEvent*>(e));
	}
}

char QWebChannel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QWebChannel*>(ptr)->QWebChannel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QWebChannel_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QWebChannel*>(ptr)->QWebChannel::metaObject());
	}
}

void QWebChannel_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlWebChannel*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlWebChannel*>(ptr)->QQmlWebChannel::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QWebChannel*>(ptr)->QWebChannel::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport
{
public:
	MyQWebChannelAbstractTransport(QObject *parent = Q_NULLPTR) : QWebChannelAbstractTransport(parent) {QWebChannelAbstractTransport_QWebChannelAbstractTransport_QRegisterMetaType();};
	void Signal_MessageReceived(const QJsonObject & message, QWebChannelAbstractTransport * transport) { callbackQWebChannelAbstractTransport_MessageReceived(this, const_cast<QJsonObject*>(&message), transport); };
	void sendMessage(const QJsonObject & message) { callbackQWebChannelAbstractTransport_SendMessage(this, const_cast<QJsonObject*>(&message)); };
	 ~MyQWebChannelAbstractTransport() { callbackQWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(this); };
	void childEvent(QChildEvent * event) { callbackQWebChannelAbstractTransport_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannelAbstractTransport_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannelAbstractTransport_CustomEvent(this, event); };
	void deleteLater() { callbackQWebChannelAbstractTransport_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebChannelAbstractTransport_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannelAbstractTransport_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebChannelAbstractTransport_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannelAbstractTransport_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannelAbstractTransport_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWebChannel_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWebChannelAbstractTransport_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannelAbstractTransport_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QWebChannelAbstractTransport*)
Q_DECLARE_METATYPE(MyQWebChannelAbstractTransport*)

int QWebChannelAbstractTransport_QWebChannelAbstractTransport_QRegisterMetaType(){qRegisterMetaType<QWebChannelAbstractTransport*>(); return qRegisterMetaType<MyQWebChannelAbstractTransport*>();}

void* QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebChannelAbstractTransport(static_cast<QObject*>(parent));
	}
}

void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr, long long t)
{
	QObject::connect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived), static_cast<Qt::ConnectionType>(t));
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
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "sendMessage", Q_ARG(const QJsonObject, *static_cast<QJsonObject*>(message)));
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr)
{
	static_cast<QWebChannelAbstractTransport*>(ptr)->~QWebChannelAbstractTransport();
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransportDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebChannelAbstractTransport___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebChannelAbstractTransport___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWebChannelAbstractTransport___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebChannelAbstractTransport___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebChannelAbstractTransport___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebChannelAbstractTransport___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebChannelAbstractTransport___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebChannelAbstractTransport___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebChannelAbstractTransport___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QWebChannelAbstractTransport_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannelAbstractTransport_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannelAbstractTransport_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::customEvent(static_cast<QEvent*>(event));
}

void QWebChannelAbstractTransport_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::deleteLater();
}

void QWebChannelAbstractTransport_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWebChannelAbstractTransport_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::event(static_cast<QEvent*>(e));
}

char QWebChannelAbstractTransport_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebChannelAbstractTransport_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::metaObject());
}

void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::timerEvent(static_cast<QTimerEvent*>(event));
}

