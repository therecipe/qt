#define protected public
#define private public

#include "scripttools.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QScriptEngine>
#include <QScriptEngineDebugger>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>

class MyQScriptEngineDebugger: public QScriptEngineDebugger
{
public:
	MyQScriptEngineDebugger(QObject *parent) : QScriptEngineDebugger(parent) {};
	void Signal_EvaluationResumed() { callbackQScriptEngineDebugger_EvaluationResumed(this, this->objectName().toUtf8().data()); };
	void Signal_EvaluationSuspended() { callbackQScriptEngineDebugger_EvaluationSuspended(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQScriptEngineDebugger_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScriptEngineDebugger_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptEngineDebugger_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQScriptEngineDebugger_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScriptEngineDebugger_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptEngineDebugger_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQScriptEngineDebugger_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptEngineDebugger_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptEngineDebugger_MetaObject(const_cast<MyQScriptEngineDebugger*>(this), this->objectName().toUtf8().data())); };
};

void* QScriptEngineDebugger_NewQScriptEngineDebugger(void* parent)
{
	return new MyQScriptEngineDebugger(static_cast<QObject*>(parent));
}

void* QScriptEngineDebugger_Action(void* ptr, int action)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->action(static_cast<QScriptEngineDebugger::DebuggerAction>(action));
}

void QScriptEngineDebugger_AttachTo(void* ptr, void* engine)
{
	static_cast<QScriptEngineDebugger*>(ptr)->attachTo(static_cast<QScriptEngine*>(engine));
}

int QScriptEngineDebugger_AutoShowStandardWindow(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->autoShowStandardWindow();
}

void* QScriptEngineDebugger_CreateStandardMenu(void* ptr, void* parent)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardMenu(static_cast<QWidget*>(parent));
}

void* QScriptEngineDebugger_CreateStandardToolBar(void* ptr, void* parent)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardToolBar(static_cast<QWidget*>(parent));
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

void QScriptEngineDebugger_SetAutoShowStandardWindow(void* ptr, int autoShow)
{
	static_cast<QScriptEngineDebugger*>(ptr)->setAutoShowStandardWindow(autoShow != 0);
}

void* QScriptEngineDebugger_StandardWindow(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->standardWindow();
}

int QScriptEngineDebugger_State(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->state();
}

void* QScriptEngineDebugger_Widget(void* ptr, int widget)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->widget(static_cast<QScriptEngineDebugger::DebuggerWidget>(widget));
}

void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->~QScriptEngineDebugger();
}

void QScriptEngineDebugger_TimerEvent(void* ptr, void* event)
{
	static_cast<QScriptEngineDebugger*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptEngineDebugger_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptEngineDebugger_ChildEvent(void* ptr, void* event)
{
	static_cast<QScriptEngineDebugger*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngineDebugger_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngineDebugger_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScriptEngineDebugger*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngineDebugger_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngineDebugger_CustomEvent(void* ptr, void* event)
{
	static_cast<QScriptEngineDebugger*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScriptEngineDebugger_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::customEvent(static_cast<QEvent*>(event));
}

void QScriptEngineDebugger_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScriptEngineDebugger*>(ptr), "deleteLater");
}

void QScriptEngineDebugger_DeleteLaterDefault(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::deleteLater();
}

void QScriptEngineDebugger_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScriptEngineDebugger*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngineDebugger_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScriptEngineDebugger_Event(void* ptr, void* e)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScriptEngineDebugger_EventDefault(void* ptr, void* e)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::event(static_cast<QEvent*>(e));
}

int QScriptEngineDebugger_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScriptEngineDebugger_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScriptEngineDebugger_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptEngineDebugger*>(ptr)->metaObject());
}

void* QScriptEngineDebugger_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::metaObject());
}

