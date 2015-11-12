#include "qscriptenginedebugger.h"
#include <QObject>
#include <QScriptEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QScriptEngineDebugger>
#include "_cgo_export.h"

class MyQScriptEngineDebugger: public QScriptEngineDebugger {
public:
void Signal_EvaluationResumed(){callbackQScriptEngineDebuggerEvaluationResumed(this->objectName().toUtf8().data());};
void Signal_EvaluationSuspended(){callbackQScriptEngineDebuggerEvaluationSuspended(this->objectName().toUtf8().data());};
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

void* QScriptEngineDebugger_Widget(void* ptr, int widget){
	return static_cast<QScriptEngineDebugger*>(ptr)->widget(static_cast<QScriptEngineDebugger::DebuggerWidget>(widget));
}

void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->~QScriptEngineDebugger();
}

