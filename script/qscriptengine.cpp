#include "qscriptengine.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QScriptEngineAgent>
#include <QScriptValue>
#include <QScriptEngine>
#include "_cgo_export.h"

class MyQScriptEngine: public QScriptEngine {
public:
};

QtObjectPtr QScriptEngine_NewQScriptEngine(){
	return new QScriptEngine();
}

QtObjectPtr QScriptEngine_NewQScriptEngine2(QtObjectPtr parent){
	return new QScriptEngine(static_cast<QObject*>(parent));
}

void QScriptEngine_AbortEvaluation(QtObjectPtr ptr, QtObjectPtr result){
	static_cast<QScriptEngine*>(ptr)->abortEvaluation(*static_cast<QScriptValue*>(result));
}

QtObjectPtr QScriptEngine_Agent(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->agent();
}

char* QScriptEngine_AvailableExtensions(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->availableExtensions().join("|").toUtf8().data();
}

void QScriptEngine_ClearExceptions(QtObjectPtr ptr){
	static_cast<QScriptEngine*>(ptr)->clearExceptions();
}

void QScriptEngine_CollectGarbage(QtObjectPtr ptr){
	static_cast<QScriptEngine*>(ptr)->collectGarbage();
}

QtObjectPtr QScriptEngine_CurrentContext(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->currentContext();
}

int QScriptEngine_HasUncaughtException(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->hasUncaughtException();
}

char* QScriptEngine_ImportedExtensions(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->importedExtensions().join("|").toUtf8().data();
}

void QScriptEngine_InstallTranslatorFunctions(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QScriptEngine*>(ptr)->installTranslatorFunctions(*static_cast<QScriptValue*>(object));
}

int QScriptEngine_IsEvaluating(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->isEvaluating();
}

void QScriptEngine_PopContext(QtObjectPtr ptr){
	static_cast<QScriptEngine*>(ptr)->popContext();
}

int QScriptEngine_ProcessEventsInterval(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->processEventsInterval();
}

QtObjectPtr QScriptEngine_PushContext(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->pushContext();
}

void QScriptEngine_ReportAdditionalMemoryCost(QtObjectPtr ptr, int size){
	static_cast<QScriptEngine*>(ptr)->reportAdditionalMemoryCost(size);
}

void QScriptEngine_SetAgent(QtObjectPtr ptr, QtObjectPtr agent){
	static_cast<QScriptEngine*>(ptr)->setAgent(static_cast<QScriptEngineAgent*>(agent));
}

void QScriptEngine_SetDefaultPrototype(QtObjectPtr ptr, int metaTypeId, QtObjectPtr prototype){
	static_cast<QScriptEngine*>(ptr)->setDefaultPrototype(metaTypeId, *static_cast<QScriptValue*>(prototype));
}

void QScriptEngine_SetGlobalObject(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QScriptEngine*>(ptr)->setGlobalObject(*static_cast<QScriptValue*>(object));
}

void QScriptEngine_SetProcessEventsInterval(QtObjectPtr ptr, int interval){
	static_cast<QScriptEngine*>(ptr)->setProcessEventsInterval(interval);
}

char* QScriptEngine_UncaughtExceptionBacktrace(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionBacktrace().join("|").toUtf8().data();
}

int QScriptEngine_UncaughtExceptionLineNumber(QtObjectPtr ptr){
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionLineNumber();
}

void QScriptEngine_DestroyQScriptEngine(QtObjectPtr ptr){
	static_cast<QScriptEngine*>(ptr)->~QScriptEngine();
}

