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

QtObjectPtr QScriptEngineAgent_NewQScriptEngineAgent(QtObjectPtr engine){
	return new QScriptEngineAgent(static_cast<QScriptEngine*>(engine));
}

void QScriptEngineAgent_ContextPop(QtObjectPtr ptr){
	static_cast<QScriptEngineAgent*>(ptr)->contextPop();
}

void QScriptEngineAgent_ContextPush(QtObjectPtr ptr){
	static_cast<QScriptEngineAgent*>(ptr)->contextPush();
}

QtObjectPtr QScriptEngineAgent_Engine(QtObjectPtr ptr){
	return static_cast<QScriptEngineAgent*>(ptr)->engine();
}

char* QScriptEngineAgent_Extension(QtObjectPtr ptr, int extension, char* argument){
	return static_cast<QScriptEngineAgent*>(ptr)->extension(static_cast<QScriptEngineAgent::Extension>(extension), QVariant(argument)).toString().toUtf8().data();
}

int QScriptEngineAgent_SupportsExtension(QtObjectPtr ptr, int extension){
	return static_cast<QScriptEngineAgent*>(ptr)->supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

void QScriptEngineAgent_DestroyQScriptEngineAgent(QtObjectPtr ptr){
	static_cast<QScriptEngineAgent*>(ptr)->~QScriptEngineAgent();
}

