#include "script.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QLatin1String>
#include <QMetaObject>
#include <QObject>
#include <QRegExp>
#include <QScriptClass>
#include <QScriptContext>
#include <QScriptContextInfo>
#include <QScriptEngine>
#include <QScriptEngineAgent>
#include <QScriptExtensionPlugin>
#include <QScriptProgram>
#include <QScriptString>
#include <QScriptSyntaxCheckResult>
#include <QScriptValue>
#include <QScriptable>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>

class MyQScriptClass: public QScriptClass {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQScriptClass(QScriptEngine *engine) : QScriptClass(engine) {};
protected:
};

void* QScriptClass_NewQScriptClass(void* engine){
	return new MyQScriptClass(static_cast<QScriptEngine*>(engine));
}

void* QScriptClass_Engine(void* ptr){
	return static_cast<QScriptClass*>(ptr)->engine();
}

void* QScriptClass_Extension(void* ptr, int extension, void* argument){
	return new QVariant(static_cast<QScriptClass*>(ptr)->extension(static_cast<QScriptClass::Extension>(extension), *static_cast<QVariant*>(argument)));
}

char* QScriptClass_Name(void* ptr){
	return static_cast<QScriptClass*>(ptr)->name().toUtf8().data();
}

void* QScriptClass_NewIterator(void* ptr, void* object){
	return static_cast<QScriptClass*>(ptr)->newIterator(*static_cast<QScriptValue*>(object));
}

void* QScriptClass_Prototype(void* ptr){
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->prototype());
}

int QScriptClass_SupportsExtension(void* ptr, int extension){
	return static_cast<QScriptClass*>(ptr)->supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

void QScriptClass_DestroyQScriptClass(void* ptr){
	static_cast<QScriptClass*>(ptr)->~QScriptClass();
}

char* QScriptClass_ObjectNameAbs(void* ptr){
	return static_cast<MyQScriptClass*>(ptr)->objectNameAbs().toUtf8().data();
}

void QScriptClass_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQScriptClass*>(ptr)->setObjectNameAbs(QString(name));
}

void* QScriptContext_ActivationObject(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->activationObject());
}

void* QScriptContext_Argument(void* ptr, int index){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argument(index));
}

int QScriptContext_ArgumentCount(void* ptr){
	return static_cast<QScriptContext*>(ptr)->argumentCount();
}

void* QScriptContext_ArgumentsObject(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argumentsObject());
}

char* QScriptContext_Backtrace(void* ptr){
	return static_cast<QScriptContext*>(ptr)->backtrace().join(",,,").toUtf8().data();
}

void* QScriptContext_Callee(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->callee());
}

void* QScriptContext_Engine(void* ptr){
	return static_cast<QScriptContext*>(ptr)->engine();
}

int QScriptContext_IsCalledAsConstructor(void* ptr){
	return static_cast<QScriptContext*>(ptr)->isCalledAsConstructor();
}

void* QScriptContext_ParentContext(void* ptr){
	return static_cast<QScriptContext*>(ptr)->parentContext();
}

void QScriptContext_SetActivationObject(void* ptr, void* activation){
	static_cast<QScriptContext*>(ptr)->setActivationObject(*static_cast<QScriptValue*>(activation));
}

void QScriptContext_SetThisObject(void* ptr, void* thisObject){
	static_cast<QScriptContext*>(ptr)->setThisObject(*static_cast<QScriptValue*>(thisObject));
}

int QScriptContext_State(void* ptr){
	return static_cast<QScriptContext*>(ptr)->state();
}

void* QScriptContext_ThisObject(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->thisObject());
}

void* QScriptContext_ThrowError(void* ptr, int error, char* text){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(static_cast<QScriptContext::Error>(error), QString(text)));
}

void* QScriptContext_ThrowError2(void* ptr, char* text){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(QString(text)));
}

void* QScriptContext_ThrowValue(void* ptr, void* value){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwValue(*static_cast<QScriptValue*>(value)));
}

char* QScriptContext_ToString(void* ptr){
	return static_cast<QScriptContext*>(ptr)->toString().toUtf8().data();
}

void QScriptContext_DestroyQScriptContext(void* ptr){
	static_cast<QScriptContext*>(ptr)->~QScriptContext();
}

void* QScriptContextInfo_NewQScriptContextInfo3(){
	return new QScriptContextInfo();
}

void* QScriptContextInfo_NewQScriptContextInfo(void* context){
	return new QScriptContextInfo(static_cast<QScriptContext*>(context));
}

void* QScriptContextInfo_NewQScriptContextInfo2(void* other){
	return new QScriptContextInfo(*static_cast<QScriptContextInfo*>(other));
}

char* QScriptContextInfo_FileName(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->fileName().toUtf8().data();
}

int QScriptContextInfo_FunctionEndLineNumber(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionEndLineNumber();
}

int QScriptContextInfo_FunctionMetaIndex(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionMetaIndex();
}

char* QScriptContextInfo_FunctionName(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionName().toUtf8().data();
}

char* QScriptContextInfo_FunctionParameterNames(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionParameterNames().join(",,,").toUtf8().data();
}

int QScriptContextInfo_FunctionStartLineNumber(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionStartLineNumber();
}

int QScriptContextInfo_FunctionType(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionType();
}

int QScriptContextInfo_IsNull(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->isNull();
}

int QScriptContextInfo_LineNumber(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->lineNumber();
}

long long QScriptContextInfo_ScriptId(void* ptr){
	return static_cast<long long>(static_cast<QScriptContextInfo*>(ptr)->scriptId());
}

void QScriptContextInfo_DestroyQScriptContextInfo(void* ptr){
	static_cast<QScriptContextInfo*>(ptr)->~QScriptContextInfo();
}

class MyQScriptEngine: public QScriptEngine {
public:
	MyQScriptEngine() : QScriptEngine() {};
	MyQScriptEngine(QObject *parent) : QScriptEngine(parent) {};
	void Signal_SignalHandlerException(const QScriptValue & exception) { callbackQScriptEngineSignalHandlerException(this->objectName().toUtf8().data(), new QScriptValue(exception)); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQScriptEngineTimerEvent(this->objectName().toUtf8().data(), event)) { QScriptEngine::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQScriptEngineChildEvent(this->objectName().toUtf8().data(), event)) { QScriptEngine::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQScriptEngineCustomEvent(this->objectName().toUtf8().data(), event)) { QScriptEngine::customEvent(event); }; };
};

void* QScriptEngine_NewQScriptEngine(){
	return new MyQScriptEngine();
}

void* QScriptEngine_NewQScriptEngine2(void* parent){
	return new MyQScriptEngine(static_cast<QObject*>(parent));
}

void QScriptEngine_AbortEvaluation(void* ptr, void* result){
	static_cast<QScriptEngine*>(ptr)->abortEvaluation(*static_cast<QScriptValue*>(result));
}

void* QScriptEngine_Agent(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->agent();
}

char* QScriptEngine_AvailableExtensions(void* ptr){
	return static_cast<QScriptEngine*>(ptr)->availableExtensions().join(",,,").toUtf8().data();
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
	return static_cast<QScriptEngine*>(ptr)->importedExtensions().join(",,,").toUtf8().data();
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
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionBacktrace().join(",,,").toUtf8().data();
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

class MyQScriptEngineAgent: public QScriptEngineAgent {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQScriptEngineAgent(QScriptEngine *engine) : QScriptEngineAgent(engine) {};
	void contextPop() { if (!callbackQScriptEngineAgentContextPop(this->objectNameAbs().toUtf8().data())) { QScriptEngineAgent::contextPop(); }; };
	void contextPush() { if (!callbackQScriptEngineAgentContextPush(this->objectNameAbs().toUtf8().data())) { QScriptEngineAgent::contextPush(); }; };
	void exceptionCatch(qint64 scriptId, const QScriptValue & exception) { if (!callbackQScriptEngineAgentExceptionCatch(this->objectNameAbs().toUtf8().data(), static_cast<long long>(scriptId), new QScriptValue(exception))) { QScriptEngineAgent::exceptionCatch(scriptId, exception); }; };
	void exceptionThrow(qint64 scriptId, const QScriptValue & exception, bool hasHandler) { if (!callbackQScriptEngineAgentExceptionThrow(this->objectNameAbs().toUtf8().data(), static_cast<long long>(scriptId), new QScriptValue(exception), hasHandler)) { QScriptEngineAgent::exceptionThrow(scriptId, exception, hasHandler); }; };
	void functionEntry(qint64 scriptId) { if (!callbackQScriptEngineAgentFunctionEntry(this->objectNameAbs().toUtf8().data(), static_cast<long long>(scriptId))) { QScriptEngineAgent::functionEntry(scriptId); }; };
	void functionExit(qint64 scriptId, const QScriptValue & returnValue) { if (!callbackQScriptEngineAgentFunctionExit(this->objectNameAbs().toUtf8().data(), static_cast<long long>(scriptId), new QScriptValue(returnValue))) { QScriptEngineAgent::functionExit(scriptId, returnValue); }; };
	void positionChange(qint64 scriptId, int lineNumber, int columnNumber) { if (!callbackQScriptEngineAgentPositionChange(this->objectNameAbs().toUtf8().data(), static_cast<long long>(scriptId), lineNumber, columnNumber)) { QScriptEngineAgent::positionChange(scriptId, lineNumber, columnNumber); }; };
	void scriptLoad(qint64 id, const QString & program, const QString & fileName, int baseLineNumber) { if (!callbackQScriptEngineAgentScriptLoad(this->objectNameAbs().toUtf8().data(), static_cast<long long>(id), program.toUtf8().data(), fileName.toUtf8().data(), baseLineNumber)) { QScriptEngineAgent::scriptLoad(id, program, fileName, baseLineNumber); }; };
	void scriptUnload(qint64 id) { if (!callbackQScriptEngineAgentScriptUnload(this->objectNameAbs().toUtf8().data(), static_cast<long long>(id))) { QScriptEngineAgent::scriptUnload(id); }; };
protected:
};

void* QScriptEngineAgent_NewQScriptEngineAgent(void* engine){
	return new MyQScriptEngineAgent(static_cast<QScriptEngine*>(engine));
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

void QScriptEngineAgent_ExceptionCatch(void* ptr, long long scriptId, void* exception){
	static_cast<QScriptEngineAgent*>(ptr)->exceptionCatch(static_cast<long long>(scriptId), *static_cast<QScriptValue*>(exception));
}

void QScriptEngineAgent_ExceptionThrow(void* ptr, long long scriptId, void* exception, int hasHandler){
	static_cast<QScriptEngineAgent*>(ptr)->exceptionThrow(static_cast<long long>(scriptId), *static_cast<QScriptValue*>(exception), hasHandler != 0);
}

void* QScriptEngineAgent_Extension(void* ptr, int extension, void* argument){
	return new QVariant(static_cast<QScriptEngineAgent*>(ptr)->extension(static_cast<QScriptEngineAgent::Extension>(extension), *static_cast<QVariant*>(argument)));
}

void QScriptEngineAgent_FunctionEntry(void* ptr, long long scriptId){
	static_cast<QScriptEngineAgent*>(ptr)->functionEntry(static_cast<long long>(scriptId));
}

void QScriptEngineAgent_FunctionExit(void* ptr, long long scriptId, void* returnValue){
	static_cast<QScriptEngineAgent*>(ptr)->functionExit(static_cast<long long>(scriptId), *static_cast<QScriptValue*>(returnValue));
}

void QScriptEngineAgent_PositionChange(void* ptr, long long scriptId, int lineNumber, int columnNumber){
	static_cast<QScriptEngineAgent*>(ptr)->positionChange(static_cast<long long>(scriptId), lineNumber, columnNumber);
}

void QScriptEngineAgent_ScriptLoad(void* ptr, long long id, char* program, char* fileName, int baseLineNumber){
	static_cast<QScriptEngineAgent*>(ptr)->scriptLoad(static_cast<long long>(id), QString(program), QString(fileName), baseLineNumber);
}

void QScriptEngineAgent_ScriptUnload(void* ptr, long long id){
	static_cast<QScriptEngineAgent*>(ptr)->scriptUnload(static_cast<long long>(id));
}

int QScriptEngineAgent_SupportsExtension(void* ptr, int extension){
	return static_cast<QScriptEngineAgent*>(ptr)->supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

void QScriptEngineAgent_DestroyQScriptEngineAgent(void* ptr){
	static_cast<QScriptEngineAgent*>(ptr)->~QScriptEngineAgent();
}

char* QScriptEngineAgent_ObjectNameAbs(void* ptr){
	return static_cast<MyQScriptEngineAgent*>(ptr)->objectNameAbs().toUtf8().data();
}

void QScriptEngineAgent_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQScriptEngineAgent*>(ptr)->setObjectNameAbs(QString(name));
}

void QScriptExtensionPlugin_Initialize(void* ptr, char* key, void* engine){
	static_cast<QScriptExtensionPlugin*>(ptr)->initialize(QString(key), static_cast<QScriptEngine*>(engine));
}

char* QScriptExtensionPlugin_Keys(void* ptr){
	return static_cast<QScriptExtensionPlugin*>(ptr)->keys().join(",,,").toUtf8().data();
}

void* QScriptExtensionPlugin_SetupPackage(void* ptr, char* key, void* engine){
	return new QScriptValue(static_cast<QScriptExtensionPlugin*>(ptr)->setupPackage(QString(key), static_cast<QScriptEngine*>(engine)));
}

void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(void* ptr){
	static_cast<QScriptExtensionPlugin*>(ptr)->~QScriptExtensionPlugin();
}

void* QScriptProgram_NewQScriptProgram(){
	return new QScriptProgram();
}

void* QScriptProgram_NewQScriptProgram3(void* other){
	return new QScriptProgram(*static_cast<QScriptProgram*>(other));
}

void* QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber){
	return new QScriptProgram(QString(sourceCode), QString(fileName), firstLineNumber);
}

char* QScriptProgram_FileName(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->fileName().toUtf8().data();
}

int QScriptProgram_FirstLineNumber(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->firstLineNumber();
}

int QScriptProgram_IsNull(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->isNull();
}

char* QScriptProgram_SourceCode(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->sourceCode().toUtf8().data();
}

void QScriptProgram_DestroyQScriptProgram(void* ptr){
	static_cast<QScriptProgram*>(ptr)->~QScriptProgram();
}

void* QScriptString_NewQScriptString(){
	return new QScriptString();
}

void* QScriptString_NewQScriptString2(void* other){
	return new QScriptString(*static_cast<QScriptString*>(other));
}

int QScriptString_IsValid(void* ptr){
	return static_cast<QScriptString*>(ptr)->isValid();
}

char* QScriptString_ToString(void* ptr){
	return static_cast<QScriptString*>(ptr)->toString().toUtf8().data();
}

void QScriptString_DestroyQScriptString(void* ptr){
	static_cast<QScriptString*>(ptr)->~QScriptString();
}

void* QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(void* other){
	return new QScriptSyntaxCheckResult(*static_cast<QScriptSyntaxCheckResult*>(other));
}

int QScriptSyntaxCheckResult_ErrorColumnNumber(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorColumnNumber();
}

int QScriptSyntaxCheckResult_ErrorLineNumber(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorLineNumber();
}

char* QScriptSyntaxCheckResult_ErrorMessage(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorMessage().toUtf8().data();
}

int QScriptSyntaxCheckResult_State(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->state();
}

void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(void* ptr){
	static_cast<QScriptSyntaxCheckResult*>(ptr)->~QScriptSyntaxCheckResult();
}

void* QScriptValue_NewQScriptValue(){
	return new QScriptValue();
}

void* QScriptValue_NewQScriptValue10(int value){
	return new QScriptValue(static_cast<QScriptValue::SpecialValue>(value));
}

void* QScriptValue_NewQScriptValue11(int value){
	return new QScriptValue(value != 0);
}

void* QScriptValue_NewQScriptValue16(void* value){
	return new QScriptValue(*static_cast<QLatin1String*>(value));
}

void* QScriptValue_NewQScriptValue2(void* other){
	return new QScriptValue(*static_cast<QScriptValue*>(other));
}

void* QScriptValue_NewQScriptValue15(char* value){
	return new QScriptValue(QString(value));
}

void* QScriptValue_NewQScriptValue17(char* value){
	return new QScriptValue(const_cast<const char*>(value));
}

void* QScriptValue_NewQScriptValue12(int value){
	return new QScriptValue(value);
}

void* QScriptValue_Call2(void* ptr, void* thisObject, void* arguments){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->call(*static_cast<QScriptValue*>(thisObject), *static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_Construct2(void* ptr, void* arguments){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->construct(*static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_Data(void* ptr){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->data());
}

void* QScriptValue_Engine(void* ptr){
	return static_cast<QScriptValue*>(ptr)->engine();
}

int QScriptValue_Equals(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->equals(*static_cast<QScriptValue*>(other));
}

int QScriptValue_InstanceOf(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->instanceOf(*static_cast<QScriptValue*>(other));
}

int QScriptValue_IsArray(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isArray();
}

int QScriptValue_IsBool(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isBool();
}

int QScriptValue_IsDate(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isDate();
}

int QScriptValue_IsError(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isError();
}

int QScriptValue_IsFunction(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isFunction();
}

int QScriptValue_IsNull(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isNull();
}

int QScriptValue_IsNumber(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isNumber();
}

int QScriptValue_IsObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isObject();
}

int QScriptValue_IsQMetaObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isQMetaObject();
}

int QScriptValue_IsQObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isQObject();
}

int QScriptValue_IsRegExp(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isRegExp();
}

int QScriptValue_IsString(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isString();
}

int QScriptValue_IsUndefined(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isUndefined();
}

int QScriptValue_IsValid(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isValid();
}

int QScriptValue_IsVariant(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isVariant();
}

int QScriptValue_LessThan(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->lessThan(*static_cast<QScriptValue*>(other));
}

void* QScriptValue_Property2(void* ptr, void* name, int mode){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Property(void* ptr, char* name, int mode){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(QString(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

int QScriptValue_PropertyFlags2(void* ptr, void* name, int mode){
	return static_cast<QScriptValue*>(ptr)->propertyFlags(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

int QScriptValue_PropertyFlags(void* ptr, char* name, int mode){
	return static_cast<QScriptValue*>(ptr)->propertyFlags(QString(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

void* QScriptValue_Prototype(void* ptr){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->prototype());
}

void* QScriptValue_ScriptClass(void* ptr){
	return static_cast<QScriptValue*>(ptr)->scriptClass();
}

void QScriptValue_SetData(void* ptr, void* data){
	static_cast<QScriptValue*>(ptr)->setData(*static_cast<QScriptValue*>(data));
}

void QScriptValue_SetProperty2(void* ptr, void* name, void* value, int flags){
	static_cast<QScriptValue*>(ptr)->setProperty(*static_cast<QScriptString*>(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetProperty(void* ptr, char* name, void* value, int flags){
	static_cast<QScriptValue*>(ptr)->setProperty(QString(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetPrototype(void* ptr, void* prototype){
	static_cast<QScriptValue*>(ptr)->setPrototype(*static_cast<QScriptValue*>(prototype));
}

void QScriptValue_SetScriptClass(void* ptr, void* scriptClass){
	static_cast<QScriptValue*>(ptr)->setScriptClass(static_cast<QScriptClass*>(scriptClass));
}

int QScriptValue_StrictlyEquals(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->strictlyEquals(*static_cast<QScriptValue*>(other));
}

int QScriptValue_ToBool(void* ptr){
	return static_cast<QScriptValue*>(ptr)->toBool();
}

void* QScriptValue_ToDateTime(void* ptr){
	return new QDateTime(static_cast<QScriptValue*>(ptr)->toDateTime());
}

void* QScriptValue_ToQMetaObject(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QScriptValue*>(ptr)->toQMetaObject());
}

void* QScriptValue_ToQObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->toQObject();
}

void* QScriptValue_ToRegExp(void* ptr){
	return new QRegExp(static_cast<QScriptValue*>(ptr)->toRegExp());
}

char* QScriptValue_ToString(void* ptr){
	return static_cast<QScriptValue*>(ptr)->toString().toUtf8().data();
}

void* QScriptValue_ToVariant(void* ptr){
	return new QVariant(static_cast<QScriptValue*>(ptr)->toVariant());
}

void QScriptValue_DestroyQScriptValue(void* ptr){
	static_cast<QScriptValue*>(ptr)->~QScriptValue();
}

void* QScriptable_Argument(void* ptr, int index){
	return new QScriptValue(static_cast<QScriptable*>(ptr)->argument(index));
}

int QScriptable_ArgumentCount(void* ptr){
	return static_cast<QScriptable*>(ptr)->argumentCount();
}

void* QScriptable_Context(void* ptr){
	return static_cast<QScriptable*>(ptr)->context();
}

void* QScriptable_Engine(void* ptr){
	return static_cast<QScriptable*>(ptr)->engine();
}

void* QScriptable_ThisObject(void* ptr){
	return new QScriptValue(static_cast<QScriptable*>(ptr)->thisObject());
}

