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
	void Signal_EvaluationResumed() { callbackQScriptEngineDebuggerEvaluationResumed(this->objectName().toUtf8().data()); };
	void Signal_EvaluationSuspended() { callbackQScriptEngineDebuggerEvaluationSuspended(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQScriptEngineDebuggerTimerEvent(this->objectName().toUtf8().data(), event)) { QScriptEngineDebugger::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQScriptEngineDebuggerChildEvent(this->objectName().toUtf8().data(), event)) { QScriptEngineDebugger::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQScriptEngineDebuggerCustomEvent(this->objectName().toUtf8().data(), event)) { QScriptEngineDebugger::customEvent(event); }; };
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

void QScriptEngineDebugger_ConnectEvaluationSuspended(void* ptr){
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));;
}

void QScriptEngineDebugger_DisconnectEvaluationSuspended(void* ptr){
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));;
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

