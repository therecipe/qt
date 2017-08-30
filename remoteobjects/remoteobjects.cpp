// +build !minimal

#define protected public
#define private public

#include "remoteobjects.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
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
#include <QItemSelection>
#include <QItemSelectionModel>
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
#include <QRemoteObjectDynamicReplica>
#include <QRemoteObjectHost>
#include <QRemoteObjectHostBase>
#include <QRemoteObjectNode>
#include <QRemoteObjectPersistedStore>
#include <QRemoteObjectRegistry>
#include <QRemoteObjectRegistryHost>
#include <QRemoteObjectReplica>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQRemoteObjectDynamicReplica: public QRemoteObjectDynamicReplica
{
public:
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectDynamicReplica*)

int QRemoteObjectDynamicReplica_QRemoteObjectDynamicReplica_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectDynamicReplica*>(); return qRegisterMetaType<MyQRemoteObjectDynamicReplica*>();}

void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(void* ptr)
{
	static_cast<QRemoteObjectDynamicReplica*>(ptr)->~QRemoteObjectDynamicReplica();
}

class MyQRemoteObjectHost: public QRemoteObjectHost
{
public:
	MyQRemoteObjectHost(QObject *parent = nullptr) : QRemoteObjectHost(parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	MyQRemoteObjectHost(const QUrl &address, QObject *parent) : QRemoteObjectHost(address, parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	MyQRemoteObjectHost(const QUrl &address, const QUrl &registryAddress = QUrl(), QObject *parent = nullptr) : QRemoteObjectHost(address, registryAddress, parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	bool setHostUrl(const QUrl & hostAddress) { return callbackQRemoteObjectHost_SetHostUrl(this, const_cast<QUrl*>(&hostAddress)) != 0; };
	 ~MyQRemoteObjectHost() { callbackQRemoteObjectHost_DestroyQRemoteObjectHost(this); };
	QUrl hostUrl() const { return *static_cast<QUrl*>(callbackQRemoteObjectHost_HostUrl(const_cast<void*>(static_cast<const void*>(this)))); };
	void setName(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void Signal_Error(QRemoteObjectNode::ErrorCode errorCode) { callbackQRemoteObjectNode_Error(this, errorCode); };
	
	
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectHost*)

int QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectHost*>(); return qRegisterMetaType<MyQRemoteObjectHost*>();}

void* QRemoteObjectHost_NewQRemoteObjectHost(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_NewQRemoteObjectHost3(void* address, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_NewQRemoteObjectHost2(void* address, void* registryAddress, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

char QRemoteObjectHost_SetHostUrl(void* ptr, void* hostAddress)
{
	return static_cast<QRemoteObjectHost*>(ptr)->setHostUrl(*static_cast<QUrl*>(hostAddress));
}

char QRemoteObjectHost_SetHostUrlDefault(void* ptr, void* hostAddress)
{
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setHostUrl(*static_cast<QUrl*>(hostAddress));
}

void QRemoteObjectHost_DestroyQRemoteObjectHost(void* ptr)
{
	static_cast<QRemoteObjectHost*>(ptr)->~QRemoteObjectHost();
}

void QRemoteObjectHost_DestroyQRemoteObjectHostDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QRemoteObjectHost_HostUrl(void* ptr)
{
	return new QUrl(static_cast<QRemoteObjectHost*>(ptr)->hostUrl());
}

void* QRemoteObjectHost_HostUrlDefault(void* ptr)
{
		return new QUrl(static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::hostUrl());
}

class MyQRemoteObjectHostBase: public QRemoteObjectHostBase
{
public:
	void setName(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void Signal_Error(QRemoteObjectNode::ErrorCode errorCode) { callbackQRemoteObjectNode_Error(this, errorCode); };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectHostBase*)

int QRemoteObjectHostBase_QRemoteObjectHostBase_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectHostBase*>(); return qRegisterMetaType<MyQRemoteObjectHostBase*>();}

char QRemoteObjectHostBase_DisableRemoting(void* ptr, void* remoteObject)
{
	return static_cast<QRemoteObjectHostBase*>(ptr)->disableRemoting(static_cast<QObject*>(remoteObject));
}

char QRemoteObjectHostBase_EnableRemoting3(void* ptr, void* model, struct QtRemoteObjects_PackedString name, void* roles, void* selectionModel)
{
	return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QAbstractItemModel*>(model), QString::fromUtf8(name.data, name.len), *static_cast<QVector<int>*>(roles), static_cast<QItemSelectionModel*>(selectionModel));
}

char QRemoteObjectHostBase_EnableRemoting2(void* ptr, void* object, struct QtRemoteObjects_PackedString name)
{
	return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len));
}

int QRemoteObjectHostBase___enableRemoting_roles_atList3(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QRemoteObjectHostBase___enableRemoting_roles_setList3(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QRemoteObjectHostBase___enableRemoting_roles_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

class MyQRemoteObjectNode: public QRemoteObjectNode
{
public:
	MyQRemoteObjectNode(QObject *parent = nullptr) : QRemoteObjectNode(parent) {QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();};
	MyQRemoteObjectNode(const QUrl &registryAddress, QObject *parent = nullptr) : QRemoteObjectNode(registryAddress, parent) {QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();};
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void Signal_Error(QRemoteObjectNode::ErrorCode errorCode) { callbackQRemoteObjectNode_Error(this, errorCode); };
	void setName(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	 ~MyQRemoteObjectNode() { callbackQRemoteObjectNode_DestroyQRemoteObjectNode(this); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectNode*)

int QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectNode*>(); return qRegisterMetaType<MyQRemoteObjectNode*>();}

void* QRemoteObjectNode_AcquireDynamic(void* ptr, struct QtRemoteObjects_PackedString name)
{
	return static_cast<QRemoteObjectNode*>(ptr)->acquireDynamic(QString::fromUtf8(name.data, name.len));
}

void* QRemoteObjectNode_NewQRemoteObjectNode(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectNode(static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectNode_NewQRemoteObjectNode2(void* registryAddress, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

char QRemoteObjectNode_ConnectToNode(void* ptr, void* address)
{
	return static_cast<QRemoteObjectNode*>(ptr)->connectToNode(*static_cast<QUrl*>(address));
}

char QRemoteObjectNode_SetRegistryUrl(void* ptr, void* registryAddress)
{
	return static_cast<QRemoteObjectNode*>(ptr)->setRegistryUrl(*static_cast<QUrl*>(registryAddress));
}

char QRemoteObjectNode_SetRegistryUrlDefault(void* ptr, void* registryAddress)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	}
}

char QRemoteObjectNode_WaitForRegistry(void* ptr, int timeout)
{
	return static_cast<QRemoteObjectNode*>(ptr)->waitForRegistry(timeout);
}

void QRemoteObjectNode_ConnectError(void* ptr)
{
	qRegisterMetaType<QRemoteObjectNode::ErrorCode>();
	QObject::connect(static_cast<QRemoteObjectNode*>(ptr), static_cast<void (QRemoteObjectNode::*)(QRemoteObjectNode::ErrorCode)>(&QRemoteObjectNode::error), static_cast<MyQRemoteObjectNode*>(ptr), static_cast<void (MyQRemoteObjectNode::*)(QRemoteObjectNode::ErrorCode)>(&MyQRemoteObjectNode::Signal_Error));
}

void QRemoteObjectNode_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectNode*>(ptr), static_cast<void (QRemoteObjectNode::*)(QRemoteObjectNode::ErrorCode)>(&QRemoteObjectNode::error), static_cast<MyQRemoteObjectNode*>(ptr), static_cast<void (MyQRemoteObjectNode::*)(QRemoteObjectNode::ErrorCode)>(&MyQRemoteObjectNode::Signal_Error));
}

void QRemoteObjectNode_Error(void* ptr, long long errorCode)
{
	static_cast<QRemoteObjectNode*>(ptr)->error(static_cast<QRemoteObjectNode::ErrorCode>(errorCode));
}

void QRemoteObjectNode_SetName(void* ptr, struct QtRemoteObjects_PackedString name)
{
	static_cast<QRemoteObjectNode*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QRemoteObjectNode_SetNameDefault(void* ptr, struct QtRemoteObjects_PackedString name)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::setName(QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setName(QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::setName(QString::fromUtf8(name.data, name.len));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::setName(QString::fromUtf8(name.data, name.len));
	}
}

void QRemoteObjectNode_SetPersistedStore(void* ptr, void* store, long long ownership)
{
	static_cast<QRemoteObjectNode*>(ptr)->setPersistedStore(static_cast<QRemoteObjectPersistedStore*>(store), static_cast<QRemoteObjectNode::StorageOwnership>(ownership));
}

void QRemoteObjectNode_TimerEventDefault(void* ptr, void* vqt)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::timerEvent(static_cast<QTimerEvent*>(vqt));
	}
}

void QRemoteObjectNode_DestroyQRemoteObjectNode(void* ptr)
{
	static_cast<QRemoteObjectNode*>(ptr)->~QRemoteObjectNode();
}

void QRemoteObjectNode_DestroyQRemoteObjectNodeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QRemoteObjectNode_LastError(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->lastError();
}

struct QtRemoteObjects_PackedString QRemoteObjectNode_Instances2(void* ptr, struct QtRemoteObjects_PackedString typeName)
{
	return ({ QByteArray t1bf5b2 = static_cast<QRemoteObjectNode*>(ptr)->instances(QString::fromUtf8(typeName.data, typeName.len)).join("|").toUtf8(); QtRemoteObjects_PackedString { const_cast<char*>(t1bf5b2.prepend("WHITESPACE").constData()+10), t1bf5b2.size()-10 }; });
}

void* QRemoteObjectNode_RegistryUrl(void* ptr)
{
	return new QUrl(static_cast<QRemoteObjectNode*>(ptr)->registryUrl());
}

void* QRemoteObjectNode_Registry(void* ptr)
{
	return const_cast<QRemoteObjectRegistry*>(static_cast<QRemoteObjectNode*>(ptr)->registry());
}

void* QRemoteObjectNode___retrieveProperties_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectNode___retrieveProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectNode___retrieveProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectNode___persistProperties_props_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectNode___persistProperties_props_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectNode___persistProperties_props_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectNode___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QRemoteObjectNode___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectNode___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QRemoteObjectNode___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRemoteObjectNode___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectNode___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRemoteObjectNode___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRemoteObjectNode___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectNode___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRemoteObjectNode___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRemoteObjectNode___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectNode___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRemoteObjectNode___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QRemoteObjectNode___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectNode___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QRemoteObjectNode_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectNode_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QRemoteObjectNode_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectNode_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectNode_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectNode_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::deleteLater();
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::deleteLater();
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::deleteLater();
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::deleteLater();
	}
}

void QRemoteObjectNode_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void* QRemoteObjectNode_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::metaObject());
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::metaObject());
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::metaObject());
	}
}

class MyQRemoteObjectPersistedStore: public QRemoteObjectPersistedStore
{
public:
	QList<QVariant> restoreProperties(const QString & repName, const QByteArray & repSig) { QByteArray t7884ba = repName.toUtf8(); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba.prepend("WHITESPACE").constData()+10), t7884ba.size()-10 };return *static_cast<QList<QVariant>*>(callbackQRemoteObjectPersistedStore_RestoreProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig))); };
	void saveProperties(const QString & repName, const QByteArray & repSig, const QVariantList & values) { QByteArray t7884ba = repName.toUtf8(); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba.prepend("WHITESPACE").constData()+10), t7884ba.size()-10 };callbackQRemoteObjectPersistedStore_SaveProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig), ({ QList<QVariant>* tmpValue = new QList<QVariant>(values); QtRemoteObjects_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~MyQRemoteObjectPersistedStore() { callbackQRemoteObjectPersistedStore_DestroyQRemoteObjectPersistedStore(this); };
};

struct QtRemoteObjects_PackedList QRemoteObjectPersistedStore_RestoreProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig)
{
	return ({ QList<QVariant>* tmpValue = new QList<QVariant>(static_cast<QRemoteObjectPersistedStore*>(ptr)->restoreProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig))); QtRemoteObjects_PackedList { tmpValue, tmpValue->size() }; });
}

void QRemoteObjectPersistedStore_SaveProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig, void* values)
{
	static_cast<QRemoteObjectPersistedStore*>(ptr)->saveProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig), *static_cast<QList<QVariant>*>(values));
}

void QRemoteObjectPersistedStore_DestroyQRemoteObjectPersistedStore(void* ptr)
{
	static_cast<QRemoteObjectPersistedStore*>(ptr)->~QRemoteObjectPersistedStore();
}

void QRemoteObjectPersistedStore_DestroyQRemoteObjectPersistedStoreDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QRemoteObjectPersistedStore___restoreProperties_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectPersistedStore___restoreProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectPersistedStore___restoreProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectPersistedStore___saveProperties_values_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectPersistedStore___saveProperties_values_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectPersistedStore___saveProperties_values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

class MyQRemoteObjectRegistry: public QRemoteObjectRegistry
{
public:
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectRegistry*)

int QRemoteObjectRegistry_QRemoteObjectRegistry_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectRegistry*>(); return qRegisterMetaType<MyQRemoteObjectRegistry*>();}

void QRemoteObjectRegistry_QRemoteObjectRegistry_RegisterMetatypes()
{
	QRemoteObjectRegistry::registerMetatypes();
}

void QRemoteObjectRegistry_DestroyQRemoteObjectRegistry(void* ptr)
{
	static_cast<QRemoteObjectRegistry*>(ptr)->~QRemoteObjectRegistry();
}

class MyQRemoteObjectRegistryHost: public QRemoteObjectRegistryHost
{
public:
	MyQRemoteObjectRegistryHost(const QUrl &registryAddress = QUrl(), QObject *parent = nullptr) : QRemoteObjectRegistryHost(registryAddress, parent) {QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType();};
	bool setRegistryUrl(const QUrl & registryUrl) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryUrl)) != 0; };
	 ~MyQRemoteObjectRegistryHost() { callbackQRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHost(this); };
	void setName(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_Error(QRemoteObjectNode::ErrorCode errorCode) { callbackQRemoteObjectNode_Error(this, errorCode); };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectRegistryHost*)

int QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectRegistryHost*>(); return qRegisterMetaType<MyQRemoteObjectRegistryHost*>();}

void* QRemoteObjectRegistryHost_NewQRemoteObjectRegistryHost(void* registryAddress, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

void QRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHost(void* ptr)
{
	static_cast<QRemoteObjectRegistryHost*>(ptr)->~QRemoteObjectRegistryHost();
}

void QRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHostDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQRemoteObjectReplica: public QRemoteObjectReplica
{
public:
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRemoteObjectReplica*)

int QRemoteObjectReplica_QRemoteObjectReplica_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectReplica*>(); return qRegisterMetaType<MyQRemoteObjectReplica*>();}

void* QRemoteObjectReplica_Node(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->node();
}

void QRemoteObjectReplica_ConnectInitialized(void* ptr)
{
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::initialized), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Initialized));
}

void QRemoteObjectReplica_DisconnectInitialized(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::initialized), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Initialized));
}

void QRemoteObjectReplica_Initialized(void* ptr)
{
	static_cast<QRemoteObjectReplica*>(ptr)->initialized();
}

void QRemoteObjectReplica_SetNode(void* ptr, void* node)
{
	static_cast<QRemoteObjectReplica*>(ptr)->setNode(static_cast<QRemoteObjectNode*>(node));
}

char QRemoteObjectReplica_WaitForSource(void* ptr, int timeout)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->waitForSource(timeout);
}

void QRemoteObjectReplica_ConnectStateChanged(void* ptr)
{
	qRegisterMetaType<QRemoteObjectReplica::State>();
	qRegisterMetaType<QRemoteObjectReplica::State>();
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&QRemoteObjectReplica::stateChanged), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&MyQRemoteObjectReplica::Signal_StateChanged));
}

void QRemoteObjectReplica_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&QRemoteObjectReplica::stateChanged), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&MyQRemoteObjectReplica::Signal_StateChanged));
}

void QRemoteObjectReplica_StateChanged(void* ptr, long long state, long long oldState)
{
	static_cast<QRemoteObjectReplica*>(ptr)->stateChanged(static_cast<QRemoteObjectReplica::State>(state), static_cast<QRemoteObjectReplica::State>(oldState));
}

long long QRemoteObjectReplica_State(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->state();
}

char QRemoteObjectReplica_IsInitialized(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->isInitialized();
}

char QRemoteObjectReplica_IsReplicaValid(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->isReplicaValid();
}

void* QRemoteObjectReplica___sendWithReply_args_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectReplica___sendWithReply_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___sendWithReply_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectReplica___send_args_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectReplica___send_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___send_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectReplica___setProperties_properties_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectReplica___setProperties_properties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___setProperties_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectReplica___retrieveProperties_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectReplica___retrieveProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___retrieveProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectReplica___persistProperties_props_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QRemoteObjectReplica___persistProperties_props_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___persistProperties_props_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>;
}

void* QRemoteObjectReplica___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QRemoteObjectReplica___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectReplica___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QRemoteObjectReplica___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRemoteObjectReplica___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectReplica___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRemoteObjectReplica___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRemoteObjectReplica___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectReplica___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRemoteObjectReplica___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRemoteObjectReplica___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectReplica___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRemoteObjectReplica___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QRemoteObjectReplica___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRemoteObjectReplica___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QRemoteObjectReplica_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectReplica_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QRemoteObjectReplica_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectReplica_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectReplica_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectReplica_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::deleteLater();
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::deleteLater();
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::deleteLater();
	}
}

void QRemoteObjectReplica_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectReplica_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QRemoteObjectReplica_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::metaObject());
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::metaObject());
	}
}

struct QtRemoteObjects_PackedString QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Local()
{
	return ({ QByteArray t4c3a40 = QRemoteObjectStringLiterals::local().toUtf8(); QtRemoteObjects_PackedString { const_cast<char*>(t4c3a40.prepend("WHITESPACE").constData()+10), t4c3a40.size()-10 }; });
}

struct QtRemoteObjects_PackedString QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Tcp()
{
	return ({ QByteArray t44baf1 = QRemoteObjectStringLiterals::tcp().toUtf8(); QtRemoteObjects_PackedString { const_cast<char*>(t44baf1.prepend("WHITESPACE").constData()+10), t44baf1.size()-10 }; });
}

