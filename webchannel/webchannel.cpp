// +build !minimal

#define protected public
#define private public

#include "webchannel.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QJsonObject>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWebChannel>
#include <QWebChannelAbstractTransport>
#include <QWidget>
#include <QWindow>

class MyQWebChannel: public QWebChannel
{
public:
	MyQWebChannel(QObject *parent = Q_NULLPTR) : QWebChannel(parent) {QWebChannel_QWebChannel_QRegisterMetaType();};
	void Signal_BlockUpdatesChanged(bool block) { callbackQWebChannel_BlockUpdatesChanged(this, block); };
	void connectTo(QWebChannelAbstractTransport * transport) { callbackQWebChannel_ConnectTo(this, transport); };
	void disconnectFrom(QWebChannelAbstractTransport * transport) { callbackQWebChannel_DisconnectFrom(this, transport); };
	bool event(QEvent * e) { return callbackQWebChannel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebChannel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannel_CustomEvent(this, event); };
	void deleteLater() { callbackQWebChannel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebChannel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebChannel_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebChannel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebChannel*)

int QWebChannel_QWebChannel_QRegisterMetaType(){qRegisterMetaType<QWebChannel*>(); return qRegisterMetaType<MyQWebChannel*>();}

void* QWebChannel_NewQWebChannel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannel(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebChannel(static_cast<QObject*>(parent));
	}
}

void QWebChannel_ConnectBlockUpdatesChanged(void* ptr)
{
	QObject::connect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));
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
		static_cast<QWebChannel*>(ptr)->QWebChannel::connectTo(static_cast<QWebChannelAbstractTransport*>(transport));
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
		static_cast<QWebChannel*>(ptr)->QWebChannel::disconnectFrom(static_cast<QWebChannelAbstractTransport*>(transport));
}

void QWebChannel_RegisterObject(void* ptr, struct QtWebChannel_PackedString id, void* object)
{
	static_cast<QWebChannel*>(ptr)->registerObject(QString::fromUtf8(id.data, id.len), static_cast<QObject*>(object));
}

void QWebChannel_RegisterObjects(void* ptr, void* objects)
{
	static_cast<QWebChannel*>(ptr)->registerObjects(*static_cast<QHash<QString, QObject *>*>(objects));
}

void QWebChannel_SetBlockUpdates(void* ptr, char block)
{
	static_cast<QWebChannel*>(ptr)->setBlockUpdates(block != 0);
}

void QWebChannel_DestroyQWebChannel(void* ptr)
{
	static_cast<QWebChannel*>(ptr)->~QWebChannel();
}

struct QtWebChannel_PackedList QWebChannel_RegisteredObjects(void* ptr)
{
	return ({ QHash<QString, QObject *>* tmpValue = new QHash<QString, QObject *>(static_cast<QWebChannel*>(ptr)->registeredObjects()); QtWebChannel_PackedList { tmpValue, tmpValue->size() }; });
}

char QWebChannel_BlockUpdates(void* ptr)
{
	return static_cast<QWebChannel*>(ptr)->blockUpdates();
}

void* QWebChannel___registerObjects_objects_atList(void* ptr, struct QtWebChannel_PackedString i)
{
	return const_cast<QObject*>(static_cast<QHash<QString, QObject *>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QWebChannel___registerObjects_objects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i)
{
	static_cast<QHash<QString, QObject *>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), static_cast<QObject*>(i));
}

void* QWebChannel___registerObjects_objects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QObject *>;
}

struct QtWebChannel_PackedList QWebChannel___registerObjects_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QHash<QString, QObject *>*>(ptr)->keys()); QtWebChannel_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebChannel___registeredObjects_atList(void* ptr, struct QtWebChannel_PackedString i)
{
	return const_cast<QObject*>(static_cast<QHash<QString, QObject *>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QWebChannel___registeredObjects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i)
{
	static_cast<QHash<QString, QObject *>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), static_cast<QObject*>(i));
}

void* QWebChannel___registeredObjects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QObject *>;
}

struct QtWebChannel_PackedList QWebChannel___registeredObjects_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QHash<QString, QObject *>*>(ptr)->keys()); QtWebChannel_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebChannel_PackedString QWebChannel_____registerObjects_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebChannel_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebChannel_____registerObjects_keyList_setList(void* ptr, struct QtWebChannel_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebChannel_____registerObjects_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtWebChannel_PackedString QWebChannel_____registeredObjects_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebChannel_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebChannel_____registeredObjects_keyList_setList(void* ptr, struct QtWebChannel_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebChannel_____registeredObjects_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

void* QWebChannel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebChannel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebChannel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebChannel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebChannel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebChannel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebChannel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebChannel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebChannel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebChannel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebChannel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebChannel_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebChannel*>(ptr)->QWebChannel::event(static_cast<QEvent*>(e));
}

char QWebChannel_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebChannel*>(ptr)->QWebChannel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebChannel_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannel*>(ptr)->QWebChannel::childEvent(static_cast<QChildEvent*>(event));
}

void QWebChannel_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebChannel*>(ptr)->QWebChannel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannel_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannel*>(ptr)->QWebChannel::customEvent(static_cast<QEvent*>(event));
}

void QWebChannel_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebChannel*>(ptr)->QWebChannel::deleteLater();
}

void QWebChannel_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebChannel*>(ptr)->QWebChannel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebChannel_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannel*>(ptr)->QWebChannel::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebChannel_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebChannel*>(ptr)->QWebChannel::metaObject());
}

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport
{
public:
	MyQWebChannelAbstractTransport(QObject *parent = Q_NULLPTR) : QWebChannelAbstractTransport(parent) {QWebChannelAbstractTransport_QWebChannelAbstractTransport_QRegisterMetaType();};
	void Signal_MessageReceived(const QJsonObject & message, QWebChannelAbstractTransport * transport) { callbackQWebChannelAbstractTransport_MessageReceived(this, const_cast<QJsonObject*>(&message), transport); };
	void sendMessage(const QJsonObject & message) { callbackQWebChannelAbstractTransport_SendMessage(this, const_cast<QJsonObject*>(&message)); };
	 ~MyQWebChannelAbstractTransport() { callbackQWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(this); };
	bool event(QEvent * e) { return callbackQWebChannelAbstractTransport_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebChannelAbstractTransport_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebChannelAbstractTransport_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebChannelAbstractTransport_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebChannelAbstractTransport_CustomEvent(this, event); };
	void deleteLater() { callbackQWebChannelAbstractTransport_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebChannelAbstractTransport_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebChannelAbstractTransport_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebChannel_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebChannelAbstractTransport_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebChannelAbstractTransport_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebChannelAbstractTransport_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebChannelAbstractTransport*)

int QWebChannelAbstractTransport_QWebChannelAbstractTransport_QRegisterMetaType(){qRegisterMetaType<QWebChannelAbstractTransport*>(); return qRegisterMetaType<MyQWebChannelAbstractTransport*>();}

void* QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebChannelAbstractTransport(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebChannelAbstractTransport(static_cast<QObject*>(parent));
	}
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

void* QWebChannelAbstractTransport___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebChannelAbstractTransport___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebChannelAbstractTransport___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebChannelAbstractTransport___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebChannelAbstractTransport___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebChannelAbstractTransport___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebChannelAbstractTransport___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebChannelAbstractTransport___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebChannelAbstractTransport___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebChannelAbstractTransport___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebChannelAbstractTransport___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebChannelAbstractTransport___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebChannelAbstractTransport_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::event(static_cast<QEvent*>(e));
}

char QWebChannelAbstractTransport_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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

void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebChannelAbstractTransport_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebChannelAbstractTransport*>(ptr)->QWebChannelAbstractTransport::metaObject());
}

