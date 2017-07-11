// +build !minimal

#define protected public
#define private public

#include "scripttools.h"
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
#include <QScriptEngine>
#include <QScriptEngineDebugger>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

class MyQScriptEngineDebugger: public QScriptEngineDebugger
{
public:
	MyQScriptEngineDebugger(QObject *parent = Q_NULLPTR) : QScriptEngineDebugger(parent) {QScriptEngineDebugger_QScriptEngineDebugger_QRegisterMetaType();};
	void Signal_EvaluationResumed() { callbackQScriptEngineDebugger_EvaluationResumed(this); };
	void Signal_EvaluationSuspended() { callbackQScriptEngineDebugger_EvaluationSuspended(this); };
	bool event(QEvent * e) { return callbackQScriptEngineDebugger_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptEngineDebugger_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScriptEngineDebugger_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptEngineDebugger_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScriptEngineDebugger_CustomEvent(this, event); };
	void deleteLater() { callbackQScriptEngineDebugger_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScriptEngineDebugger_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptEngineDebugger_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScriptTools_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScriptEngineDebugger_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScriptEngineDebugger_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptEngineDebugger_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScriptEngineDebugger*)

int QScriptEngineDebugger_QScriptEngineDebugger_QRegisterMetaType(){qRegisterMetaType<QScriptEngineDebugger*>(); return qRegisterMetaType<MyQScriptEngineDebugger*>();}

void* QScriptEngineDebugger_CreateStandardMenu(void* ptr, void* parent)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardMenu(static_cast<QWidget*>(parent));
}

void* QScriptEngineDebugger_NewQScriptEngineDebugger(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngineDebugger(static_cast<QWindow*>(parent));
	} else {
		return new MyQScriptEngineDebugger(static_cast<QObject*>(parent));
	}
}

void* QScriptEngineDebugger_CreateStandardToolBar(void* ptr, void* parent)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardToolBar(static_cast<QWidget*>(parent));
}

void QScriptEngineDebugger_AttachTo(void* ptr, void* engine)
{
	static_cast<QScriptEngineDebugger*>(ptr)->attachTo(static_cast<QScriptEngine*>(engine));
}

void QScriptEngineDebugger_Detach(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->detach();
}

void QScriptEngineDebugger_ConnectEvaluationResumed(void* ptr)
{
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationResumed), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationResumed));
}

void QScriptEngineDebugger_DisconnectEvaluationResumed(void* ptr)
{
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationResumed), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationResumed));
}

void QScriptEngineDebugger_EvaluationResumed(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->evaluationResumed();
}

void QScriptEngineDebugger_ConnectEvaluationSuspended(void* ptr)
{
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));
}

void QScriptEngineDebugger_DisconnectEvaluationSuspended(void* ptr)
{
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));
}

void QScriptEngineDebugger_EvaluationSuspended(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->evaluationSuspended();
}

void QScriptEngineDebugger_SetAutoShowStandardWindow(void* ptr, char autoShow)
{
	static_cast<QScriptEngineDebugger*>(ptr)->setAutoShowStandardWindow(autoShow != 0);
}

void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->~QScriptEngineDebugger();
}

long long QScriptEngineDebugger_State(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->state();
}

void* QScriptEngineDebugger_Action(void* ptr, long long action)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->action(static_cast<QScriptEngineDebugger::DebuggerAction>(action));
}

void* QScriptEngineDebugger_StandardWindow(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->standardWindow();
}

void* QScriptEngineDebugger_Widget(void* ptr, long long widget)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->widget(static_cast<QScriptEngineDebugger::DebuggerWidget>(widget));
}

char QScriptEngineDebugger_AutoShowStandardWindow(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->autoShowStandardWindow();
}

void* QScriptEngineDebugger___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScriptEngineDebugger___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScriptEngineDebugger___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QScriptEngineDebugger___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptEngineDebugger___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngineDebugger___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptEngineDebugger___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptEngineDebugger___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngineDebugger___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptEngineDebugger___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptEngineDebugger___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngineDebugger___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptEngineDebugger___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScriptEngineDebugger___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngineDebugger___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScriptEngineDebugger_EventDefault(void* ptr, void* e)
{
		return static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::event(static_cast<QEvent*>(e));
}

char QScriptEngineDebugger_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QScriptEngineDebugger_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngineDebugger_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngineDebugger_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::customEvent(static_cast<QEvent*>(event));
}

void QScriptEngineDebugger_DeleteLaterDefault(void* ptr)
{
		static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::deleteLater();
}

void QScriptEngineDebugger_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngineDebugger_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QScriptEngineDebugger_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::metaObject());
}

