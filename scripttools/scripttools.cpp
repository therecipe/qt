// +build !minimal

#define protected public
#define private public

#include "scripttools.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QList>
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
	void Signal_EvaluationResumed() { callbackQScriptEngineDebugger_EvaluationResumed(this); };
	void Signal_EvaluationSuspended() { callbackQScriptEngineDebugger_EvaluationSuspended(this); };
	void timerEvent(QTimerEvent * event) { callbackQScriptEngineDebugger_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScriptEngineDebugger_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptEngineDebugger_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScriptEngineDebugger_CustomEvent(this, event); };
	void deleteLater() { callbackQScriptEngineDebugger_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptEngineDebugger_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScriptEngineDebugger_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptEngineDebugger_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptEngineDebugger_MetaObject(const_cast<MyQScriptEngineDebugger*>(this))); };
};

void* QScriptEngineDebugger_NewQScriptEngineDebugger(void* parent)
{
	return new MyQScriptEngineDebugger(static_cast<QObject*>(parent));
}

void* QScriptEngineDebugger_Action(void* ptr, long long action)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->action(static_cast<QScriptEngineDebugger::DebuggerAction>(action));
}

void QScriptEngineDebugger_AttachTo(void* ptr, void* engine)
{
	static_cast<QScriptEngineDebugger*>(ptr)->attachTo(static_cast<QScriptEngine*>(engine));
}

char QScriptEngineDebugger_AutoShowStandardWindow(void* ptr)
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

void QScriptEngineDebugger_SetAutoShowStandardWindow(void* ptr, char autoShow)
{
	static_cast<QScriptEngineDebugger*>(ptr)->setAutoShowStandardWindow(autoShow != 0);
}

void* QScriptEngineDebugger_StandardWindow(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->standardWindow();
}

long long QScriptEngineDebugger_State(void* ptr)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->state();
}

void* QScriptEngineDebugger_Widget(void* ptr, long long widget)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->widget(static_cast<QScriptEngineDebugger::DebuggerWidget>(widget));
}

void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr)
{
	static_cast<QScriptEngineDebugger*>(ptr)->~QScriptEngineDebugger();
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
	return new QList<QObject *>;
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
	return new QList<QObject*>;
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

char QScriptEngineDebugger_Event(void* ptr, void* e)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScriptEngineDebugger_EventDefault(void* ptr, void* e)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::event(static_cast<QEvent*>(e));
}

char QScriptEngineDebugger_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptEngineDebugger*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScriptEngineDebugger_EventFilterDefault(void* ptr, void* watched, void* event)
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

