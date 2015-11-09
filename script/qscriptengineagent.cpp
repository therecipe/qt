#include "qscriptengineagent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptEngine>
#include <QScriptEngineAgent>
#include "_cgo_export.h"

class MyQScriptEngineAgent: public QScriptEngineAgent {
public:
};

void* QScriptEngineAgent_NewQScriptEngineAgent(void* engine){
	return new QScriptEngineAgent(static_cast<QScriptEngine*>(engine));
}

void QScriptEngineAgent_ContextPop(void* ptr){
	static_cast<QScriptEngineAgent*>(ptr)->contextPop();
}

void QScriptEngineAgent_ContextPush(void* ptr){
	static_cast<QScriptEngineAgent*>(ptr)->contextPush();
}

void* QScriptEngineAgent_Engine(void* ptr){
	return static_cast<QScriptEngineAgent*>(ptr)->engine();
}

void* QScriptEngineAgent_Extension(void* ptr, int extension, void* argument){
	return new QVariant(static_cast<QScriptEngineAgent*>(ptr)->extension(static_cast<QScriptEngineAgent::Extension>(extension), *static_cast<QVariant*>(argument)));
}

int QScriptEngineAgent_SupportsExtension(void* ptr, int extension){
	return static_cast<QScriptEngineAgent*>(ptr)->supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

void QScriptEngineAgent_DestroyQScriptEngineAgent(void* ptr){
	static_cast<QScriptEngineAgent*>(ptr)->~QScriptEngineAgent();
}

