#define protected public

#include "scripttools.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QObject>
#include <QScriptEngine>
#include <QScriptEngineDebugger>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>

class MyQScriptEngineDebugger: public QScriptEngineDebugger {
public:
	void Signal_EvaluationResumed() { callbackQScriptEngineDebuggerEvaluationResumed(this, this->objectName().toUtf8().data()); };
	void Signal_EvaluationSuspended() { callbackQScriptEngineDebuggerEvaluationSuspended(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQScriptEngineDebuggerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScriptEngineDebuggerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQScriptEngineDebuggerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QScriptEngineDebugger_NewQScriptEngineDebugger(void* parent){
	return new QScriptEngineDebugger(static_cast<QObject*>(parent));
}

void* QScriptEngineDebugger_Action(void* ptr, int action){
	return static_cast<QScriptEngineDebugger*>(ptr)->action(static_cast<QScriptEngineDebugger::DebuggerAction>(action));
}

void QScriptEngineDebugger_AttachTo(void* ptr, void* engine){
	static_cast<QScriptEngineDebugger*>(ptr)->attachTo(static_cast<QScriptEngine*>(engine));
}

int QScriptEngineDebugger_AutoShowStandardWindow(void* ptr){
	return static_cast<QScriptEngineDebugger*>(ptr)->autoShowStandardWindow();
}

void* QScriptEngineDebugger_CreateStandardMenu(void* ptr, void* parent){
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardMenu(static_cast<QWidget*>(parent));
}

void* QScriptEngineDebugger_CreateStandardToolBar(void* ptr, void* parent){
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardToolBar(static_cast<QWidget*>(parent));
}

void QScriptEngineDebugger_Detach(void* ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->detach();
}

void QScriptEngineDebugger_ConnectEvaluationResumed(void* ptr){
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationResumed), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationResumed));;
}

void QScriptEngineDebugger_DisconnectEvaluationResumed(void* ptr){
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationResumed), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationResumed));;
}

void QScriptEngineDebugger_EvaluationResumed(void* ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->evaluationResumed();
}

void QScriptEngineDebugger_ConnectEvaluationSuspended(void* ptr){
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));;
}

void QScriptEngineDebugger_DisconnectEvaluationSuspended(void* ptr){
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));;
}

void QScriptEngineDebugger_EvaluationSuspended(void* ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->evaluationSuspended();
}

void QScriptEngineDebugger_SetAutoShowStandardWindow(void* ptr, int autoShow){
	static_cast<QScriptEngineDebugger*>(ptr)->setAutoShowStandardWindow(autoShow != 0);
}

void* QScriptEngineDebugger_StandardWindow(void* ptr){
	return static_cast<QScriptEngineDebugger*>(ptr)->standardWindow();
}

int QScriptEngineDebugger_State(void* ptr){
	return static_cast<QScriptEngineDebugger*>(ptr)->state();
}

void* QScriptEngineDebugger_Widget(void* ptr, int widget){
	return static_cast<QScriptEngineDebugger*>(ptr)->widget(static_cast<QScriptEngineDebugger::DebuggerWidget>(widget));
}

void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->~QScriptEngineDebugger();
}

void QScriptEngineDebugger_TimerEvent(void* ptr, void* event){
	static_cast<MyQScriptEngineDebugger*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptEngineDebugger_TimerEventDefault(void* ptr, void* event){
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptEngineDebugger_ChildEvent(void* ptr, void* event){
	static_cast<MyQScriptEngineDebugger*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngineDebugger_ChildEventDefault(void* ptr, void* event){
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngineDebugger_CustomEvent(void* ptr, void* event){
	static_cast<MyQScriptEngineDebugger*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScriptEngineDebugger_CustomEventDefault(void* ptr, void* event){
	static_cast<QScriptEngineDebugger*>(ptr)->QScriptEngineDebugger::customEvent(static_cast<QEvent*>(event));
}

