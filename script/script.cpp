// +build !minimal

#define protected public
#define private public

#include "script.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QLatin1String>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QRegExp>
#include <QScriptClass>
#include <QScriptClassPropertyIterator>
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
#include <QStringList>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>

class MyQScriptClass: public QScriptClass
{
public:
	MyQScriptClass(QScriptEngine *engine) : QScriptClass(engine) {};
	QVariant extension(QScriptClass::Extension extension, const QVariant & argument) { return *static_cast<QVariant*>(callbackQScriptClass_Extension(this, extension, const_cast<QVariant*>(&argument))); };
	QString name() const { return QString(callbackQScriptClass_Name(const_cast<MyQScriptClass*>(this))); };
	QScriptClassPropertyIterator * newIterator(const QScriptValue & object) { return static_cast<QScriptClassPropertyIterator*>(callbackQScriptClass_NewIterator(this, const_cast<QScriptValue*>(&object))); };
	QScriptValue property(const QScriptValue & object, const QScriptString & name, uint id) { return *static_cast<QScriptValue*>(callbackQScriptClass_Property(this, const_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), id)); };
	QScriptValue::PropertyFlags propertyFlags(const QScriptValue & object, const QScriptString & name, uint id) { return static_cast<QScriptValue::PropertyFlag>(callbackQScriptClass_PropertyFlags(this, const_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), id)); };
	QScriptValue prototype() const { return *static_cast<QScriptValue*>(callbackQScriptClass_Prototype(const_cast<MyQScriptClass*>(this))); };
	QueryFlags queryProperty(const QScriptValue & object, const QScriptString & name, QScriptClass::QueryFlags flags, uint * id) { return static_cast<QScriptClass::QueryFlag>(callbackQScriptClass_QueryProperty(this, const_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), flags, *id)); };
	void setProperty(QScriptValue & object, const QScriptString & name, uint id, const QScriptValue & value) { callbackQScriptClass_SetProperty(this, new QScriptValue(object), const_cast<QScriptString*>(&name), id, const_cast<QScriptValue*>(&value)); };
	bool supportsExtension(QScriptClass::Extension extension) const { return callbackQScriptClass_SupportsExtension(const_cast<MyQScriptClass*>(this), extension) != 0; };
	 ~MyQScriptClass() { callbackQScriptClass_DestroyQScriptClass(this); };
};

void* QScriptClass_NewQScriptClass(void* engine)
{
	return new MyQScriptClass(static_cast<QScriptEngine*>(engine));
}

void* QScriptClass_Engine(void* ptr)
{
	return static_cast<QScriptClass*>(ptr)->engine();
}

void* QScriptClass_Extension(void* ptr, long long extension, void* argument)
{
	return new QVariant(static_cast<QScriptClass*>(ptr)->extension(static_cast<QScriptClass::Extension>(extension), *static_cast<QVariant*>(argument)));
}

void* QScriptClass_ExtensionDefault(void* ptr, long long extension, void* argument)
{
	return new QVariant(static_cast<QScriptClass*>(ptr)->QScriptClass::extension(static_cast<QScriptClass::Extension>(extension), *static_cast<QVariant*>(argument)));
}

char* QScriptClass_Name(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptClass*>(ptr)->name().toUtf8().prepend("WHITESPACE").constData()+10);
}

char* QScriptClass_NameDefault(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptClass*>(ptr)->QScriptClass::name().toUtf8().prepend("WHITESPACE").constData()+10);
}

void* QScriptClass_NewIterator(void* ptr, void* object)
{
	return static_cast<QScriptClass*>(ptr)->newIterator(*static_cast<QScriptValue*>(object));
}

void* QScriptClass_NewIteratorDefault(void* ptr, void* object)
{
	return static_cast<QScriptClass*>(ptr)->QScriptClass::newIterator(*static_cast<QScriptValue*>(object));
}

void* QScriptClass_Property(void* ptr, void* object, void* name, unsigned int id)
{
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->property(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), id));
}

void* QScriptClass_PropertyDefault(void* ptr, void* object, void* name, unsigned int id)
{
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->QScriptClass::property(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), id));
}

long long QScriptClass_PropertyFlags(void* ptr, void* object, void* name, unsigned int id)
{
	return static_cast<QScriptClass*>(ptr)->propertyFlags(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), id);
}

long long QScriptClass_PropertyFlagsDefault(void* ptr, void* object, void* name, unsigned int id)
{
	return static_cast<QScriptClass*>(ptr)->QScriptClass::propertyFlags(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), id);
}

void* QScriptClass_Prototype(void* ptr)
{
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->prototype());
}

void* QScriptClass_PrototypeDefault(void* ptr)
{
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->QScriptClass::prototype());
}

long long QScriptClass_QueryProperty(void* ptr, void* object, void* name, long long flags, unsigned int id)
{
	return static_cast<QScriptClass*>(ptr)->queryProperty(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), static_cast<QScriptClass::QueryFlag>(flags), &id);
}

long long QScriptClass_QueryPropertyDefault(void* ptr, void* object, void* name, long long flags, unsigned int id)
{
	return static_cast<QScriptClass*>(ptr)->QScriptClass::queryProperty(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), static_cast<QScriptClass::QueryFlag>(flags), &id);
}

void QScriptClass_SetProperty(void* ptr, void* object, void* name, unsigned int id, void* value)
{
	static_cast<QScriptClass*>(ptr)->setProperty(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), id, *static_cast<QScriptValue*>(value));
}

void QScriptClass_SetPropertyDefault(void* ptr, void* object, void* name, unsigned int id, void* value)
{
	static_cast<QScriptClass*>(ptr)->QScriptClass::setProperty(*static_cast<QScriptValue*>(object), *static_cast<QScriptString*>(name), id, *static_cast<QScriptValue*>(value));
}

char QScriptClass_SupportsExtension(void* ptr, long long extension)
{
	return static_cast<QScriptClass*>(ptr)->supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

char QScriptClass_SupportsExtensionDefault(void* ptr, long long extension)
{
	return static_cast<QScriptClass*>(ptr)->QScriptClass::supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

void QScriptClass_DestroyQScriptClass(void* ptr)
{
	static_cast<QScriptClass*>(ptr)->~QScriptClass();
}

void QScriptClass_DestroyQScriptClassDefault(void* ptr)
{

}

void* QScriptContext_ActivationObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->activationObject());
}

void* QScriptContext_Argument(void* ptr, int index)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argument(index));
}

int QScriptContext_ArgumentCount(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->argumentCount();
}

void* QScriptContext_ArgumentsObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argumentsObject());
}

char* QScriptContext_Backtrace(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptContext*>(ptr)->backtrace().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

void* QScriptContext_Callee(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->callee());
}

void* QScriptContext_Engine(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->engine();
}

char QScriptContext_IsCalledAsConstructor(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->isCalledAsConstructor();
}

void* QScriptContext_ParentContext(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->parentContext();
}

void QScriptContext_SetActivationObject(void* ptr, void* activation)
{
	static_cast<QScriptContext*>(ptr)->setActivationObject(*static_cast<QScriptValue*>(activation));
}

void QScriptContext_SetThisObject(void* ptr, void* thisObject)
{
	static_cast<QScriptContext*>(ptr)->setThisObject(*static_cast<QScriptValue*>(thisObject));
}

long long QScriptContext_State(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->state();
}

void* QScriptContext_ThisObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->thisObject());
}

void* QScriptContext_ThrowError(void* ptr, long long error, char* text)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(static_cast<QScriptContext::Error>(error), QString(text)));
}

void* QScriptContext_ThrowError2(void* ptr, char* text)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(QString(text)));
}

void* QScriptContext_ThrowValue(void* ptr, void* value)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwValue(*static_cast<QScriptValue*>(value)));
}

char* QScriptContext_ToString(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptContext*>(ptr)->toString().toUtf8().prepend("WHITESPACE").constData()+10);
}

void QScriptContext_DestroyQScriptContext(void* ptr)
{
	static_cast<QScriptContext*>(ptr)->~QScriptContext();
}

void* QScriptContextInfo_NewQScriptContextInfo3()
{
	return new QScriptContextInfo();
}

void* QScriptContextInfo_NewQScriptContextInfo(void* context)
{
	return new QScriptContextInfo(static_cast<QScriptContext*>(context));
}

void* QScriptContextInfo_NewQScriptContextInfo2(void* other)
{
	return new QScriptContextInfo(*static_cast<QScriptContextInfo*>(other));
}

char* QScriptContextInfo_FileName(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptContextInfo*>(ptr)->fileName().toUtf8().prepend("WHITESPACE").constData()+10);
}

int QScriptContextInfo_FunctionEndLineNumber(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionEndLineNumber();
}

int QScriptContextInfo_FunctionMetaIndex(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionMetaIndex();
}

char* QScriptContextInfo_FunctionName(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptContextInfo*>(ptr)->functionName().toUtf8().prepend("WHITESPACE").constData()+10);
}

char* QScriptContextInfo_FunctionParameterNames(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptContextInfo*>(ptr)->functionParameterNames().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

int QScriptContextInfo_FunctionStartLineNumber(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionStartLineNumber();
}

long long QScriptContextInfo_FunctionType(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionType();
}

char QScriptContextInfo_IsNull(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->isNull();
}

int QScriptContextInfo_LineNumber(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->lineNumber();
}

long long QScriptContextInfo_ScriptId(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->scriptId();
}

void QScriptContextInfo_DestroyQScriptContextInfo(void* ptr)
{
	static_cast<QScriptContextInfo*>(ptr)->~QScriptContextInfo();
}

class MyQScriptEngine: public QScriptEngine
{
public:
	MyQScriptEngine() : QScriptEngine() {};
	MyQScriptEngine(QObject *parent) : QScriptEngine(parent) {};
	void Signal_SignalHandlerException(const QScriptValue & exception) { callbackQScriptEngine_SignalHandlerException(this, const_cast<QScriptValue*>(&exception)); };
	 ~MyQScriptEngine() { callbackQScriptEngine_DestroyQScriptEngine(this); };
	void timerEvent(QTimerEvent * event) { callbackQScriptEngine_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScriptEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScriptEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQScriptEngine_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScriptEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptEngine_MetaObject(const_cast<MyQScriptEngine*>(this))); };
};

void* QScriptEngine_NewQScriptEngine()
{
	return new MyQScriptEngine();
}

void* QScriptEngine_NewQScriptEngine2(void* parent)
{
	return new MyQScriptEngine(static_cast<QObject*>(parent));
}

void QScriptEngine_AbortEvaluation(void* ptr, void* result)
{
	static_cast<QScriptEngine*>(ptr)->abortEvaluation(*static_cast<QScriptValue*>(result));
}

void* QScriptEngine_Agent(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->agent();
}

char* QScriptEngine_AvailableExtensions(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptEngine*>(ptr)->availableExtensions().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

void* QScriptEngine_QScriptEngine_CheckSyntax(char* program)
{
	return new QScriptSyntaxCheckResult(QScriptEngine::checkSyntax(QString(program)));
}

void QScriptEngine_ClearExceptions(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->clearExceptions();
}

void QScriptEngine_CollectGarbage(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->collectGarbage();
}

void* QScriptEngine_CurrentContext(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->currentContext();
}

void* QScriptEngine_DefaultPrototype(void* ptr, int metaTypeId)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->defaultPrototype(metaTypeId));
}

void* QScriptEngine_Evaluate2(void* ptr, void* program)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->evaluate(*static_cast<QScriptProgram*>(program)));
}

void* QScriptEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->evaluate(QString(program), QString(fileName), lineNumber));
}

void* QScriptEngine_GlobalObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->globalObject());
}

char QScriptEngine_HasUncaughtException(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->hasUncaughtException();
}

void* QScriptEngine_ImportExtension(void* ptr, char* extension)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->importExtension(QString(extension)));
}

char* QScriptEngine_ImportedExtensions(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptEngine*>(ptr)->importedExtensions().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

void QScriptEngine_InstallTranslatorFunctions(void* ptr, void* object)
{
	static_cast<QScriptEngine*>(ptr)->installTranslatorFunctions(*static_cast<QScriptValue*>(object));
}

char QScriptEngine_IsEvaluating(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->isEvaluating();
}

void* QScriptEngine_NewArray(void* ptr, unsigned int length)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newArray(length));
}

void* QScriptEngine_NewDate2(void* ptr, void* value)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newDate(*static_cast<QDateTime*>(value)));
}

void* QScriptEngine_NewObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newObject());
}

void* QScriptEngine_NewObject2(void* ptr, void* scriptClass, void* data)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newObject(static_cast<QScriptClass*>(scriptClass), *static_cast<QScriptValue*>(data)));
}

void* QScriptEngine_NewQMetaObject(void* ptr, void* metaObject, void* ctor)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newQMetaObject(static_cast<QMetaObject*>(metaObject), *static_cast<QScriptValue*>(ctor)));
}

void* QScriptEngine_NewQObject(void* ptr, void* object, long long ownership, long long options)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newQObject(static_cast<QObject*>(object), static_cast<QScriptEngine::ValueOwnership>(ownership), static_cast<QScriptEngine::QObjectWrapOption>(options)));
}

void* QScriptEngine_NewQObject2(void* ptr, void* scriptObject, void* qtObject, long long ownership, long long options)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newQObject(*static_cast<QScriptValue*>(scriptObject), static_cast<QObject*>(qtObject), static_cast<QScriptEngine::ValueOwnership>(ownership), static_cast<QScriptEngine::QObjectWrapOption>(options)));
}

void* QScriptEngine_NewRegExp(void* ptr, void* regexp)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newRegExp(*static_cast<QRegExp*>(regexp)));
}

void* QScriptEngine_NewRegExp2(void* ptr, char* pattern, char* flags)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newRegExp(QString(pattern), QString(flags)));
}

void* QScriptEngine_NewVariant2(void* ptr, void* object, void* value)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newVariant(*static_cast<QScriptValue*>(object), *static_cast<QVariant*>(value)));
}

void* QScriptEngine_NewVariant(void* ptr, void* value)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newVariant(*static_cast<QVariant*>(value)));
}

void* QScriptEngine_NullValue(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->nullValue());
}

void QScriptEngine_PopContext(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->popContext();
}

int QScriptEngine_ProcessEventsInterval(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->processEventsInterval();
}

void* QScriptEngine_PushContext(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->pushContext();
}

void QScriptEngine_ReportAdditionalMemoryCost(void* ptr, int size)
{
	static_cast<QScriptEngine*>(ptr)->reportAdditionalMemoryCost(size);
}

void QScriptEngine_SetAgent(void* ptr, void* agent)
{
	static_cast<QScriptEngine*>(ptr)->setAgent(static_cast<QScriptEngineAgent*>(agent));
}

void QScriptEngine_SetDefaultPrototype(void* ptr, int metaTypeId, void* prototype)
{
	static_cast<QScriptEngine*>(ptr)->setDefaultPrototype(metaTypeId, *static_cast<QScriptValue*>(prototype));
}

void QScriptEngine_SetGlobalObject(void* ptr, void* object)
{
	static_cast<QScriptEngine*>(ptr)->setGlobalObject(*static_cast<QScriptValue*>(object));
}

void QScriptEngine_SetProcessEventsInterval(void* ptr, int interval)
{
	static_cast<QScriptEngine*>(ptr)->setProcessEventsInterval(interval);
}

void QScriptEngine_ConnectSignalHandlerException(void* ptr)
{
	QObject::connect(static_cast<QScriptEngine*>(ptr), static_cast<void (QScriptEngine::*)(const QScriptValue &)>(&QScriptEngine::signalHandlerException), static_cast<MyQScriptEngine*>(ptr), static_cast<void (MyQScriptEngine::*)(const QScriptValue &)>(&MyQScriptEngine::Signal_SignalHandlerException));
}

void QScriptEngine_DisconnectSignalHandlerException(void* ptr)
{
	QObject::disconnect(static_cast<QScriptEngine*>(ptr), static_cast<void (QScriptEngine::*)(const QScriptValue &)>(&QScriptEngine::signalHandlerException), static_cast<MyQScriptEngine*>(ptr), static_cast<void (MyQScriptEngine::*)(const QScriptValue &)>(&MyQScriptEngine::Signal_SignalHandlerException));
}

void QScriptEngine_SignalHandlerException(void* ptr, void* exception)
{
	static_cast<QScriptEngine*>(ptr)->signalHandlerException(*static_cast<QScriptValue*>(exception));
}

void* QScriptEngine_ToObject(void* ptr, void* value)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->toObject(*static_cast<QScriptValue*>(value)));
}

void* QScriptEngine_ToStringHandle(void* ptr, char* str)
{
	return new QScriptString(static_cast<QScriptEngine*>(ptr)->toStringHandle(QString(str)));
}

void* QScriptEngine_UncaughtException(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->uncaughtException());
}

char* QScriptEngine_UncaughtExceptionBacktrace(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptEngine*>(ptr)->uncaughtExceptionBacktrace().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

int QScriptEngine_UncaughtExceptionLineNumber(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionLineNumber();
}

void* QScriptEngine_UndefinedValue(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->undefinedValue());
}

void QScriptEngine_DestroyQScriptEngine(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->~QScriptEngine();
}

void QScriptEngine_DestroyQScriptEngineDefault(void* ptr)
{

}

void QScriptEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QScriptEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScriptEngine*>(ptr)->QScriptEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QScriptEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScriptEngine*>(ptr)->QScriptEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScriptEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScriptEngine*>(ptr)->QScriptEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QScriptEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScriptEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScriptEngine*>(ptr)->QScriptEngine::customEvent(static_cast<QEvent*>(event));
}

void QScriptEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScriptEngine*>(ptr), "deleteLater");
}

void QScriptEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->QScriptEngine::deleteLater();
}

void QScriptEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScriptEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScriptEngine*>(ptr)->QScriptEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScriptEngine_Event(void* ptr, void* e)
{
	return static_cast<QScriptEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScriptEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QScriptEngine*>(ptr)->QScriptEngine::event(static_cast<QEvent*>(e));
}

char QScriptEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScriptEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptEngine*>(ptr)->QScriptEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScriptEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptEngine*>(ptr)->metaObject());
}

void* QScriptEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptEngine*>(ptr)->QScriptEngine::metaObject());
}

class MyQScriptEngineAgent: public QScriptEngineAgent
{
public:
	MyQScriptEngineAgent(QScriptEngine *engine) : QScriptEngineAgent(engine) {};
	void contextPop() { callbackQScriptEngineAgent_ContextPop(this); };
	void contextPush() { callbackQScriptEngineAgent_ContextPush(this); };
	void exceptionCatch(qint64 scriptId, const QScriptValue & exception) { callbackQScriptEngineAgent_ExceptionCatch(this, scriptId, const_cast<QScriptValue*>(&exception)); };
	void exceptionThrow(qint64 scriptId, const QScriptValue & exception, bool hasHandler) { callbackQScriptEngineAgent_ExceptionThrow(this, scriptId, const_cast<QScriptValue*>(&exception), hasHandler); };
	QVariant extension(QScriptEngineAgent::Extension extension, const QVariant & argument) { return *static_cast<QVariant*>(callbackQScriptEngineAgent_Extension(this, extension, const_cast<QVariant*>(&argument))); };
	void functionEntry(qint64 scriptId) { callbackQScriptEngineAgent_FunctionEntry(this, scriptId); };
	void functionExit(qint64 scriptId, const QScriptValue & returnValue) { callbackQScriptEngineAgent_FunctionExit(this, scriptId, const_cast<QScriptValue*>(&returnValue)); };
	void positionChange(qint64 scriptId, int lineNumber, int columnNumber) { callbackQScriptEngineAgent_PositionChange(this, scriptId, lineNumber, columnNumber); };
	void scriptLoad(qint64 id, const QString & program, const QString & fileName, int baseLineNumber) { callbackQScriptEngineAgent_ScriptLoad(this, id, const_cast<char*>(program.toUtf8().prepend("WHITESPACE").constData()+10), const_cast<char*>(fileName.toUtf8().prepend("WHITESPACE").constData()+10), baseLineNumber); };
	void scriptUnload(qint64 id) { callbackQScriptEngineAgent_ScriptUnload(this, id); };
	bool supportsExtension(QScriptEngineAgent::Extension extension) const { return callbackQScriptEngineAgent_SupportsExtension(const_cast<MyQScriptEngineAgent*>(this), extension) != 0; };
	 ~MyQScriptEngineAgent() { callbackQScriptEngineAgent_DestroyQScriptEngineAgent(this); };
};

void* QScriptEngineAgent_NewQScriptEngineAgent(void* engine)
{
	return new MyQScriptEngineAgent(static_cast<QScriptEngine*>(engine));
}

void QScriptEngineAgent_ContextPop(void* ptr)
{
	static_cast<QScriptEngineAgent*>(ptr)->contextPop();
}

void QScriptEngineAgent_ContextPopDefault(void* ptr)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::contextPop();
}

void QScriptEngineAgent_ContextPush(void* ptr)
{
	static_cast<QScriptEngineAgent*>(ptr)->contextPush();
}

void QScriptEngineAgent_ContextPushDefault(void* ptr)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::contextPush();
}

void* QScriptEngineAgent_Engine(void* ptr)
{
	return static_cast<QScriptEngineAgent*>(ptr)->engine();
}

void QScriptEngineAgent_ExceptionCatch(void* ptr, long long scriptId, void* exception)
{
	static_cast<QScriptEngineAgent*>(ptr)->exceptionCatch(scriptId, *static_cast<QScriptValue*>(exception));
}

void QScriptEngineAgent_ExceptionCatchDefault(void* ptr, long long scriptId, void* exception)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::exceptionCatch(scriptId, *static_cast<QScriptValue*>(exception));
}

void QScriptEngineAgent_ExceptionThrow(void* ptr, long long scriptId, void* exception, char hasHandler)
{
	static_cast<QScriptEngineAgent*>(ptr)->exceptionThrow(scriptId, *static_cast<QScriptValue*>(exception), hasHandler != 0);
}

void QScriptEngineAgent_ExceptionThrowDefault(void* ptr, long long scriptId, void* exception, char hasHandler)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::exceptionThrow(scriptId, *static_cast<QScriptValue*>(exception), hasHandler != 0);
}

void* QScriptEngineAgent_Extension(void* ptr, long long extension, void* argument)
{
	return new QVariant(static_cast<QScriptEngineAgent*>(ptr)->extension(static_cast<QScriptEngineAgent::Extension>(extension), *static_cast<QVariant*>(argument)));
}

void* QScriptEngineAgent_ExtensionDefault(void* ptr, long long extension, void* argument)
{
	return new QVariant(static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::extension(static_cast<QScriptEngineAgent::Extension>(extension), *static_cast<QVariant*>(argument)));
}

void QScriptEngineAgent_FunctionEntry(void* ptr, long long scriptId)
{
	static_cast<QScriptEngineAgent*>(ptr)->functionEntry(scriptId);
}

void QScriptEngineAgent_FunctionEntryDefault(void* ptr, long long scriptId)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::functionEntry(scriptId);
}

void QScriptEngineAgent_FunctionExit(void* ptr, long long scriptId, void* returnValue)
{
	static_cast<QScriptEngineAgent*>(ptr)->functionExit(scriptId, *static_cast<QScriptValue*>(returnValue));
}

void QScriptEngineAgent_FunctionExitDefault(void* ptr, long long scriptId, void* returnValue)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::functionExit(scriptId, *static_cast<QScriptValue*>(returnValue));
}

void QScriptEngineAgent_PositionChange(void* ptr, long long scriptId, int lineNumber, int columnNumber)
{
	static_cast<QScriptEngineAgent*>(ptr)->positionChange(scriptId, lineNumber, columnNumber);
}

void QScriptEngineAgent_PositionChangeDefault(void* ptr, long long scriptId, int lineNumber, int columnNumber)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::positionChange(scriptId, lineNumber, columnNumber);
}

void QScriptEngineAgent_ScriptLoad(void* ptr, long long id, char* program, char* fileName, int baseLineNumber)
{
	static_cast<QScriptEngineAgent*>(ptr)->scriptLoad(id, QString(program), QString(fileName), baseLineNumber);
}

void QScriptEngineAgent_ScriptLoadDefault(void* ptr, long long id, char* program, char* fileName, int baseLineNumber)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::scriptLoad(id, QString(program), QString(fileName), baseLineNumber);
}

void QScriptEngineAgent_ScriptUnload(void* ptr, long long id)
{
	static_cast<QScriptEngineAgent*>(ptr)->scriptUnload(id);
}

void QScriptEngineAgent_ScriptUnloadDefault(void* ptr, long long id)
{
	static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::scriptUnload(id);
}

char QScriptEngineAgent_SupportsExtension(void* ptr, long long extension)
{
	return static_cast<QScriptEngineAgent*>(ptr)->supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

char QScriptEngineAgent_SupportsExtensionDefault(void* ptr, long long extension)
{
	return static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

void QScriptEngineAgent_DestroyQScriptEngineAgent(void* ptr)
{
	static_cast<QScriptEngineAgent*>(ptr)->~QScriptEngineAgent();
}

void QScriptEngineAgent_DestroyQScriptEngineAgentDefault(void* ptr)
{

}

class MyQScriptExtensionPlugin: public QScriptExtensionPlugin
{
public:
	MyQScriptExtensionPlugin(QObject *parent) : QScriptExtensionPlugin(parent) {};
	void initialize(const QString & key, QScriptEngine * engine) { callbackQScriptExtensionPlugin_Initialize(this, const_cast<char*>(key.toUtf8().prepend("WHITESPACE").constData()+10), engine); };
	QStringList keys() const { return QString(callbackQScriptExtensionPlugin_Keys(const_cast<MyQScriptExtensionPlugin*>(this))).split("|", QString::SkipEmptyParts); };
	void timerEvent(QTimerEvent * event) { callbackQScriptExtensionPlugin_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScriptExtensionPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptExtensionPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScriptExtensionPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQScriptExtensionPlugin_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptExtensionPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScriptExtensionPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptExtensionPlugin_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptExtensionPlugin_MetaObject(const_cast<MyQScriptExtensionPlugin*>(this))); };
};

void* QScriptExtensionPlugin_NewQScriptExtensionPlugin(void* parent)
{
	return new MyQScriptExtensionPlugin(static_cast<QObject*>(parent));
}

void QScriptExtensionPlugin_Initialize(void* ptr, char* key, void* engine)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->initialize(QString(key), static_cast<QScriptEngine*>(engine));
}

char* QScriptExtensionPlugin_Keys(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptExtensionPlugin*>(ptr)->keys().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

void* QScriptExtensionPlugin_SetupPackage(void* ptr, char* key, void* engine)
{
	return new QScriptValue(static_cast<QScriptExtensionPlugin*>(ptr)->setupPackage(QString(key), static_cast<QScriptEngine*>(engine)));
}

void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(void* ptr)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->~QScriptExtensionPlugin();
}

void QScriptExtensionPlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptExtensionPlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScriptExtensionPlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScriptExtensionPlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptExtensionPlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptExtensionPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptExtensionPlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScriptExtensionPlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::customEvent(static_cast<QEvent*>(event));
}

void QScriptExtensionPlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScriptExtensionPlugin*>(ptr), "deleteLater");
}

void QScriptExtensionPlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::deleteLater();
}

void QScriptExtensionPlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptExtensionPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScriptExtensionPlugin_Event(void* ptr, void* e)
{
	return static_cast<QScriptExtensionPlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScriptExtensionPlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::event(static_cast<QEvent*>(e));
}

char QScriptExtensionPlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptExtensionPlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScriptExtensionPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScriptExtensionPlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptExtensionPlugin*>(ptr)->metaObject());
}

void* QScriptExtensionPlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::metaObject());
}

void* QScriptProgram_NewQScriptProgram()
{
	return new QScriptProgram();
}

void* QScriptProgram_NewQScriptProgram3(void* other)
{
	return new QScriptProgram(*static_cast<QScriptProgram*>(other));
}

void* QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber)
{
	return new QScriptProgram(QString(sourceCode), QString(fileName), firstLineNumber);
}

char* QScriptProgram_FileName(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptProgram*>(ptr)->fileName().toUtf8().prepend("WHITESPACE").constData()+10);
}

int QScriptProgram_FirstLineNumber(void* ptr)
{
	return static_cast<QScriptProgram*>(ptr)->firstLineNumber();
}

char QScriptProgram_IsNull(void* ptr)
{
	return static_cast<QScriptProgram*>(ptr)->isNull();
}

char* QScriptProgram_SourceCode(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptProgram*>(ptr)->sourceCode().toUtf8().prepend("WHITESPACE").constData()+10);
}

void QScriptProgram_DestroyQScriptProgram(void* ptr)
{
	static_cast<QScriptProgram*>(ptr)->~QScriptProgram();
}

void* QScriptString_NewQScriptString()
{
	return new QScriptString();
}

void* QScriptString_NewQScriptString2(void* other)
{
	return new QScriptString(*static_cast<QScriptString*>(other));
}

char QScriptString_IsValid(void* ptr)
{
	return static_cast<QScriptString*>(ptr)->isValid();
}

unsigned int QScriptString_ToArrayIndex(void* ptr, char ok)
{
	return static_cast<QScriptString*>(ptr)->toArrayIndex(NULL);
}

char* QScriptString_ToString(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptString*>(ptr)->toString().toUtf8().prepend("WHITESPACE").constData()+10);
}

void QScriptString_DestroyQScriptString(void* ptr)
{
	static_cast<QScriptString*>(ptr)->~QScriptString();
}

void* QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(void* other)
{
	return new QScriptSyntaxCheckResult(*static_cast<QScriptSyntaxCheckResult*>(other));
}

int QScriptSyntaxCheckResult_ErrorColumnNumber(void* ptr)
{
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorColumnNumber();
}

int QScriptSyntaxCheckResult_ErrorLineNumber(void* ptr)
{
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorLineNumber();
}

char* QScriptSyntaxCheckResult_ErrorMessage(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptSyntaxCheckResult*>(ptr)->errorMessage().toUtf8().prepend("WHITESPACE").constData()+10);
}

long long QScriptSyntaxCheckResult_State(void* ptr)
{
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->state();
}

void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(void* ptr)
{
	static_cast<QScriptSyntaxCheckResult*>(ptr)->~QScriptSyntaxCheckResult();
}

void* QScriptValue_NewQScriptValue()
{
	return new QScriptValue();
}

void* QScriptValue_NewQScriptValue10(long long value)
{
	return new QScriptValue(static_cast<QScriptValue::SpecialValue>(value));
}

void* QScriptValue_NewQScriptValue11(char value)
{
	return new QScriptValue(value != 0);
}

void* QScriptValue_NewQScriptValue16(void* value)
{
	return new QScriptValue(*static_cast<QLatin1String*>(value));
}

void* QScriptValue_NewQScriptValue2(void* other)
{
	return new QScriptValue(*static_cast<QScriptValue*>(other));
}

void* QScriptValue_NewQScriptValue15(char* value)
{
	return new QScriptValue(QString(value));
}

void* QScriptValue_NewQScriptValue17(char* value)
{
	return new QScriptValue(const_cast<const char*>(value));
}

void* QScriptValue_NewQScriptValue12(int value)
{
	return new QScriptValue(value);
}

void* QScriptValue_NewQScriptValue13(unsigned int value)
{
	return new QScriptValue(value);
}

void* QScriptValue_Call2(void* ptr, void* thisObject, void* arguments)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->call(*static_cast<QScriptValue*>(thisObject), *static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_Construct2(void* ptr, void* arguments)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->construct(*static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_Data(void* ptr)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->data());
}

void* QScriptValue_Engine(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->engine();
}

char QScriptValue_Equals(void* ptr, void* other)
{
	return static_cast<QScriptValue*>(ptr)->equals(*static_cast<QScriptValue*>(other));
}

char QScriptValue_InstanceOf(void* ptr, void* other)
{
	return static_cast<QScriptValue*>(ptr)->instanceOf(*static_cast<QScriptValue*>(other));
}

char QScriptValue_IsArray(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isArray();
}

char QScriptValue_IsBool(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isBool();
}

char QScriptValue_IsDate(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isDate();
}

char QScriptValue_IsError(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isError();
}

char QScriptValue_IsFunction(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isFunction();
}

char QScriptValue_IsNull(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isNull();
}

char QScriptValue_IsNumber(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isNumber();
}

char QScriptValue_IsObject(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isObject();
}

char QScriptValue_IsQMetaObject(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isQMetaObject();
}

char QScriptValue_IsQObject(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isQObject();
}

char QScriptValue_IsRegExp(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isRegExp();
}

char QScriptValue_IsString(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isString();
}

char QScriptValue_IsUndefined(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isUndefined();
}

char QScriptValue_IsValid(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isValid();
}

char QScriptValue_IsVariant(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->isVariant();
}

char QScriptValue_LessThan(void* ptr, void* other)
{
	return static_cast<QScriptValue*>(ptr)->lessThan(*static_cast<QScriptValue*>(other));
}

void* QScriptValue_Property3(void* ptr, void* name, long long mode)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Property(void* ptr, char* name, long long mode)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(QString(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Property2(void* ptr, unsigned int arrayIndex, long long mode)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(arrayIndex, static_cast<QScriptValue::ResolveFlag>(mode)));
}

long long QScriptValue_PropertyFlags2(void* ptr, void* name, long long mode)
{
	return static_cast<QScriptValue*>(ptr)->propertyFlags(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

long long QScriptValue_PropertyFlags(void* ptr, char* name, long long mode)
{
	return static_cast<QScriptValue*>(ptr)->propertyFlags(QString(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

void* QScriptValue_Prototype(void* ptr)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->prototype());
}

void* QScriptValue_ScriptClass(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->scriptClass();
}

void QScriptValue_SetData(void* ptr, void* data)
{
	static_cast<QScriptValue*>(ptr)->setData(*static_cast<QScriptValue*>(data));
}

void QScriptValue_SetProperty3(void* ptr, void* name, void* value, long long flags)
{
	static_cast<QScriptValue*>(ptr)->setProperty(*static_cast<QScriptString*>(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetProperty(void* ptr, char* name, void* value, long long flags)
{
	static_cast<QScriptValue*>(ptr)->setProperty(QString(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetProperty2(void* ptr, unsigned int arrayIndex, void* value, long long flags)
{
	static_cast<QScriptValue*>(ptr)->setProperty(arrayIndex, *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetPrototype(void* ptr, void* prototype)
{
	static_cast<QScriptValue*>(ptr)->setPrototype(*static_cast<QScriptValue*>(prototype));
}

void QScriptValue_SetScriptClass(void* ptr, void* scriptClass)
{
	static_cast<QScriptValue*>(ptr)->setScriptClass(static_cast<QScriptClass*>(scriptClass));
}

char QScriptValue_StrictlyEquals(void* ptr, void* other)
{
	return static_cast<QScriptValue*>(ptr)->strictlyEquals(*static_cast<QScriptValue*>(other));
}

char QScriptValue_ToBool(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toBool();
}

void* QScriptValue_ToDateTime(void* ptr)
{
	return new QDateTime(static_cast<QScriptValue*>(ptr)->toDateTime());
}

int QScriptValue_ToInt32(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toInt32();
}

void* QScriptValue_ToQMetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptValue*>(ptr)->toQMetaObject());
}

void* QScriptValue_ToQObject(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toQObject();
}

void* QScriptValue_ToRegExp(void* ptr)
{
	return new QRegExp(static_cast<QScriptValue*>(ptr)->toRegExp());
}

char* QScriptValue_ToString(void* ptr)
{
	return const_cast<char*>(static_cast<QScriptValue*>(ptr)->toString().toUtf8().prepend("WHITESPACE").constData()+10);
}

unsigned short QScriptValue_ToUInt16(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toUInt16();
}

unsigned int QScriptValue_ToUInt32(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toUInt32();
}

void* QScriptValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QScriptValue*>(ptr)->toVariant());
}

void QScriptValue_DestroyQScriptValue(void* ptr)
{
	static_cast<QScriptValue*>(ptr)->~QScriptValue();
}

void* QScriptable_Argument(void* ptr, int index)
{
	return new QScriptValue(static_cast<QScriptable*>(ptr)->argument(index));
}

int QScriptable_ArgumentCount(void* ptr)
{
	return static_cast<QScriptable*>(ptr)->argumentCount();
}

void* QScriptable_Context(void* ptr)
{
	return static_cast<QScriptable*>(ptr)->context();
}

void* QScriptable_Engine(void* ptr)
{
	return static_cast<QScriptable*>(ptr)->engine();
}

void* QScriptable_ThisObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptable*>(ptr)->thisObject());
}

