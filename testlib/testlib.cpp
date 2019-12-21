// +build !minimal

#define protected public
#define private public

#include "testlib.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractItemModelTester>
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
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPoint>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QSignalSpy>
#include <QString>
#include <QTestEventList>
#include <QTimerEvent>
#include <QVector>
#include <QWidget>
#include <QWindow>

Q_DECLARE_METATYPE(QAbstractItemModelTester*)
void* QAbstractItemModelTester_NewQAbstractItemModelTester(void* model, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QWindow*>(parent));
	} else {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
	}
}

void* QAbstractItemModelTester_NewQAbstractItemModelTester2(void* model, long long mode, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QWindow*>(parent));
	} else {
		return new QAbstractItemModelTester(static_cast<QAbstractItemModel*>(model), static_cast<QAbstractItemModelTester::FailureReportingMode>(mode), static_cast<QObject*>(parent));
	}
}

void* QAbstractItemModelTester_Model(void* ptr)
{
	return static_cast<QAbstractItemModelTester*>(ptr)->model();
}

class MyQSignalSpy: public QSignalSpy
{
public:
	MyQSignalSpy(const QObject *object, const char *sign) : QSignalSpy(object, sign) {QSignalSpy_QSignalSpy_QRegisterMetaType();};
	void childEvent(QChildEvent * event) { callbackQSignalSpy_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSignalSpy_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSignalSpy_CustomEvent(this, event); };
	void deleteLater() { callbackQSignalSpy_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSignalSpy_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSignalSpy_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSignalSpy_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSignalSpy_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSignalSpy_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtTestLib_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSignalSpy_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSignalSpy_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSignalSpy*)
Q_DECLARE_METATYPE(MyQSignalSpy*)

int QSignalSpy_QSignalSpy_QRegisterMetaType(){qRegisterMetaType<QSignalSpy*>(); return qRegisterMetaType<MyQSignalSpy*>();}

void* QSignalSpy_NewQSignalSpy(void* object, char* sign)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QAudioSystemPlugin*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QCameraImageCapture*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QDBusPendingCallWatcher*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QExtensionFactory*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QExtensionManager*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QGraphicsObject*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QGraphicsWidget*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QLayout*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QMediaPlaylist*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QMediaRecorder*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QMediaServiceProviderPlugin*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QOffscreenSurface*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QPaintDeviceWindow*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QPdfWriter*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QQuickItem*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QRadioData*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QRemoteObjectPendingCallWatcher*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QScriptExtensionPlugin*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QWidget*>(object), const_cast<const char*>(sign));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(object))) {
		return new MyQSignalSpy(static_cast<QWindow*>(object), const_cast<const char*>(sign));
	} else {
		return new MyQSignalSpy(static_cast<QObject*>(object), const_cast<const char*>(sign));
	}
}

char QSignalSpy_IsValid(void* ptr)
{
	return static_cast<QSignalSpy*>(ptr)->isValid();
}

void* QSignalSpy_Signal(void* ptr)
{
	return new QByteArray(static_cast<QSignalSpy*>(ptr)->signal());
}

char QSignalSpy_Wait(void* ptr, int timeout)
{
	return static_cast<QSignalSpy*>(ptr)->wait(timeout);
}

int QSignalSpy___args_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QSignalSpy___args_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QSignalSpy___args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

int QSignalSpy___setArgs__atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QSignalSpy___setArgs__setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QSignalSpy___setArgs__newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QSignalSpy___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSignalSpy___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSignalSpy___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QSignalSpy___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSignalSpy___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSignalSpy___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSignalSpy___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSignalSpy___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSignalSpy___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QSignalSpy___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSignalSpy___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSignalSpy___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QSignalSpy_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::childEvent(static_cast<QChildEvent*>(event));
}

void QSignalSpy_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSignalSpy_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::customEvent(static_cast<QEvent*>(event));
}

void QSignalSpy_DeleteLaterDefault(void* ptr)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::deleteLater();
}

void QSignalSpy_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSignalSpy_EventDefault(void* ptr, void* e)
{
		return static_cast<QSignalSpy*>(ptr)->QSignalSpy::event(static_cast<QEvent*>(e));
}

char QSignalSpy_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSignalSpy*>(ptr)->QSignalSpy::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSignalSpy_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSignalSpy*>(ptr)->QSignalSpy::metaObject());
}

void QSignalSpy_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QTestEventList*)
void* QTestEventList_NewQTestEventList()
{
	return new QTestEventList();
}

void* QTestEventList_NewQTestEventList2(void* other)
{
	return new QTestEventList(*static_cast<QTestEventList*>(other));
}

void QTestEventList_AddDelay(void* ptr, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addDelay(msecs);
}

void QTestEventList_AddKeyClick(void* ptr, long long qtKey, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyClick(static_cast<Qt::Key>(qtKey), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyClick2(void* ptr, char* ascii, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyClick(*ascii, static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyClicks(void* ptr, struct QtTestLib_PackedString keys, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyClicks(QString::fromUtf8(keys.data, keys.len), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyPress(void* ptr, long long qtKey, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyPress(static_cast<Qt::Key>(qtKey), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyPress2(void* ptr, char* ascii, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyPress(*ascii, static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyRelease(void* ptr, long long qtKey, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyRelease(static_cast<Qt::Key>(qtKey), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyRelease2(void* ptr, char* ascii, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyRelease(*ascii, static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddMouseClick(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseClick(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMouseDClick(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseDClick(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMouseMove(void* ptr, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseMove(*static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMousePress(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMousePress(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMouseRelease(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseRelease(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_Clear(void* ptr)
{
	static_cast<QTestEventList*>(ptr)->clear();
}

void QTestEventList_Simulate(void* ptr, void* w)
{
	static_cast<QTestEventList*>(ptr)->simulate(static_cast<QWidget*>(w));
}

void QTestEventList_DestroyQTestEventList(void* ptr)
{
	static_cast<QTestEventList*>(ptr)->~QTestEventList();
}

