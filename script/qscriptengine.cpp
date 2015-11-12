#include "qscriptengine.h"
#include <QScriptValue>
#include <QMetaObject>
#include <QScriptProgram>
#include <QDateTime>
#include <QRegExp>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptEngineAgent>
#include <QDate>
#include <QObject>
#include <QScriptClass>
#include <QString>
#include <QScriptEngine>
#include "_cgo_export.h"

class MyQScriptEngine: public QScriptEngine {
public:
void Signal_SignalHandlerException(const QScriptValue & exception){callbackQScriptEngineSignalHandlerException(this->objectName().toUtf8().data(), new QScriptValue(exception));};
};

void* QScriptEngine_NewQScriptEngine(){
	return new QScriptEngine();
}

void* QScriptEngine_NewQScriptEngine2(void* parent){
	return new QScriptEngine(static_cast<QObject*>(parent));
}

void QScriptEngine_AbortEvaluation(void* ptr, void* result){
	static_cast<QScriptEngine*>(ptr)->abortEvaluation(*static_cast<QScriptValue*>(result));
}

void* QScriptEngine_Agent(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->agent();
}

char* QScriptEngine_AvailableExtensions(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->availableExtensions().join("|").toUtf8().data();
}

void QScriptEngine_ClearExceptions(void* ptr){
	static_cast<QScriptEngine*>(ptr)->clearExceptions();
}

void QScriptEngine_CollectGarbage(void* ptr){
	static_cast<QScriptEngine*>(ptr)->collectGarbage();
}

void* QScriptEngine_CurrentContext(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->currentContext();
}

void* QScriptEngine_DefaultPrototype(void* ptr, int metaTypeId){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->defaultPrototype(metaTypeId));
}

void* QScriptEngine_Evaluate2(void* ptr, void* program){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->evaluate(*static_cast<QScriptProgram*>(program)));
}

void* QScriptEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->evaluate(QString(program), QString(fileName), lineNumber));
}

void* QScriptEngine_GlobalObject(void* ptr){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->globalObject());
}

int QScriptEngine_HasUncaughtException(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->hasUncaughtException();
}

void* QScriptEngine_ImportExtension(void* ptr, char* extension){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->importExtension(QString(extension)));
}

char* QScriptEngine_ImportedExtensions(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->importedExtensions().join("|").toUtf8().data();
}

void QScriptEngine_InstallTranslatorFunctions(void* ptr, void* object){
	static_cast<QScriptEngine*>(ptr)->installTranslatorFunctions(*static_cast<QScriptValue*>(object));
}

int QScriptEngine_IsEvaluating(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->isEvaluating();
}

void* QScriptEngine_NewDate2(void* ptr, void* value){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newDate(*static_cast<QDateTime*>(value)));
}

void* QScriptEngine_NewObject(void* ptr){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newObject());
}

void* QScriptEngine_NewObject2(void* ptr, void* scriptClass, void* data){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newObject(static_cast<QScriptClass*>(scriptClass), *static_cast<QScriptValue*>(data)));
}

void* QScriptEngine_NewQMetaObject(void* ptr, void* metaObject, void* ctor){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newQMetaObject(static_cast<QMetaObject*>(metaObject), *static_cast<QScriptValue*>(ctor)));
}

void* QScriptEngine_NewQObject(void* ptr, void* object, int ownership, int options){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newQObject(static_cast<QObject*>(object), static_cast<QScriptEngine::ValueOwnership>(ownership), static_cast<QScriptEngine::QObjectWrapOption>(options)));
}

void* QScriptEngine_NewQObject2(void* ptr, void* scriptObject, void* qtObject, int ownership, int options){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newQObject(*static_cast<QScriptValue*>(scriptObject), static_cast<QObject*>(qtObject), static_cast<QScriptEngine::ValueOwnership>(ownership), static_cast<QScriptEngine::QObjectWrapOption>(options)));
}

void* QScriptEngine_NewRegExp(void* ptr, void* regexp){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newRegExp(*static_cast<QRegExp*>(regexp)));
}

void* QScriptEngine_NewRegExp2(void* ptr, char* pattern, char* flags){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newRegExp(QString(pattern), QString(flags)));
}

void* QScriptEngine_NewVariant2(void* ptr, void* object, void* value){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newVariant(*static_cast<QScriptValue*>(object), *static_cast<QVariant*>(value)));
}

void* QScriptEngine_NewVariant(void* ptr, void* value){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newVariant(*static_cast<QVariant*>(value)));
}

void* QScriptEngine_NullValue(void* ptr){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->nullValue());
}

void QScriptEngine_PopContext(void* ptr){
	static_cast<QScriptEngine*>(ptr)->popContext();
}

int QScriptEngine_ProcessEventsInterval(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->processEventsInterval();
}

void* QScriptEngine_PushContext(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->pushContext();
}

void QScriptEngine_ReportAdditionalMemoryCost(void* ptr, int size){
	static_cast<QScriptEngine*>(ptr)->reportAdditionalMemoryCost(size);
}

void QScriptEngine_SetAgent(void* ptr, void* agent){
	static_cast<QScriptEngine*>(ptr)->setAgent(static_cast<QScriptEngineAgent*>(agent));
}

void QScriptEngine_SetDefaultPrototype(void* ptr, int metaTypeId, void* prototype){
	static_cast<QScriptEngine*>(ptr)->setDefaultPrototype(metaTypeId, *static_cast<QScriptValue*>(prototype));
}

void QScriptEngine_SetGlobalObject(void* ptr, void* object){
	static_cast<QScriptEngine*>(ptr)->setGlobalObject(*static_cast<QScriptValue*>(object));
}

void QScriptEngine_SetProcessEventsInterval(void* ptr, int interval){
	static_cast<QScriptEngine*>(ptr)->setProcessEventsInterval(interval);
}

void QScriptEngine_ConnectSignalHandlerException(void* ptr){
	QObject::connect(static_cast<QScriptEngine*>(ptr), static_cast<void (QScriptEngine::*)(const QScriptValue &)>(&QScriptEngine::signalHandlerException), static_cast<MyQScriptEngine*>(ptr), static_cast<void (MyQScriptEngine::*)(const QScriptValue &)>(&MyQScriptEngine::Signal_SignalHandlerException));;
}

void QScriptEngine_DisconnectSignalHandlerException(void* ptr){
	QObject::disconnect(static_cast<QScriptEngine*>(ptr), static_cast<void (QScriptEngine::*)(const QScriptValue &)>(&QScriptEngine::signalHandlerException), static_cast<MyQScriptEngine*>(ptr), static_cast<void (MyQScriptEngine::*)(const QScriptValue &)>(&MyQScriptEngine::Signal_SignalHandlerException));;
}

void* QScriptEngine_ToObject(void* ptr, void* value){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->toObject(*static_cast<QScriptValue*>(value)));
}

void* QScriptEngine_UncaughtException(void* ptr){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->uncaughtException());
}

char* QScriptEngine_UncaughtExceptionBacktrace(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionBacktrace().join("|").toUtf8().data();
}

int QScriptEngine_UncaughtExceptionLineNumber(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionLineNumber();
}

void* QScriptEngine_UndefinedValue(void* ptr){
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->undefinedValue());
}

void QScriptEngine_DestroyQScriptEngine(void* ptr){
	static_cast<QScriptEngine*>(ptr)->~QScriptEngine();
}

