// +build !minimal

#define protected public
#define private public

#include "script.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLatin1String>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
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
#include <QSignalSpy>
#include <QString>
#include <QStringList>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

class MyQScriptClass: public QScriptClass
{
public:
	MyQScriptClass(QScriptEngine *engine) : QScriptClass(engine) {};
	QScriptClassPropertyIterator * newIterator(const QScriptValue & object) { return static_cast<QScriptClassPropertyIterator*>(callbackQScriptClass_NewIterator(this, const_cast<QScriptValue*>(&object))); };
	QScriptValue property(const QScriptValue & object, const QScriptString & name, uint id) { return *static_cast<QScriptValue*>(callbackQScriptClass_Property(this, const_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), id)); };
	QScriptValue::PropertyFlags propertyFlags(const QScriptValue & object, const QScriptString & name, uint id) { return static_cast<QScriptValue::PropertyFlag>(callbackQScriptClass_PropertyFlags(this, const_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), id)); };
	QVariant extension(QScriptClass::Extension extension, const QVariant & argument) { return *static_cast<QVariant*>(callbackQScriptClass_Extension(this, extension, const_cast<QVariant*>(&argument))); };
	QueryFlags queryProperty(const QScriptValue & object, const QScriptString & name, QScriptClass::QueryFlags flags, uint * id) { return static_cast<QScriptClass::QueryFlag>(callbackQScriptClass_QueryProperty(this, const_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), flags, *id)); };
	void setProperty(QScriptValue & object, const QScriptString & name, uint id, const QScriptValue & value) { callbackQScriptClass_SetProperty(this, static_cast<QScriptValue*>(&object), const_cast<QScriptString*>(&name), id, const_cast<QScriptValue*>(&value)); };
	 ~MyQScriptClass() { callbackQScriptClass_DestroyQScriptClass(this); };
	QScriptValue prototype() const { return *static_cast<QScriptValue*>(callbackQScriptClass_Prototype(const_cast<void*>(static_cast<const void*>(this)))); };
	QString name() const { return ({ QtScript_PackedString tempVal = callbackQScriptClass_Name(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool supportsExtension(QScriptClass::Extension extension) const { return callbackQScriptClass_SupportsExtension(const_cast<void*>(static_cast<const void*>(this)), extension) != 0; };
};

void* QScriptClass_NewQScriptClass(void* engine)
{
	return new MyQScriptClass(static_cast<QScriptEngine*>(engine));
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

void* QScriptClass_Extension(void* ptr, long long extension, void* argument)
{
	return new QVariant(static_cast<QScriptClass*>(ptr)->extension(static_cast<QScriptClass::Extension>(extension), *static_cast<QVariant*>(argument)));
}

void* QScriptClass_ExtensionDefault(void* ptr, long long extension, void* argument)
{
		return new QVariant(static_cast<QScriptClass*>(ptr)->QScriptClass::extension(static_cast<QScriptClass::Extension>(extension), *static_cast<QVariant*>(argument)));
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

void QScriptClass_DestroyQScriptClass(void* ptr)
{
	static_cast<QScriptClass*>(ptr)->~QScriptClass();
}

void QScriptClass_DestroyQScriptClassDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScriptClass_Engine(void* ptr)
{
	return static_cast<QScriptClass*>(ptr)->engine();
}

void* QScriptClass_Prototype(void* ptr)
{
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->prototype());
}

void* QScriptClass_PrototypeDefault(void* ptr)
{
		return new QScriptValue(static_cast<QScriptClass*>(ptr)->QScriptClass::prototype());
}

struct QtScript_PackedString QScriptClass_Name(void* ptr)
{
	return ({ QByteArray t130656 = static_cast<QScriptClass*>(ptr)->name().toUtf8(); QtScript_PackedString { const_cast<char*>(t130656.prepend("WHITESPACE").constData()+10), t130656.size()-10 }; });
}

struct QtScript_PackedString QScriptClass_NameDefault(void* ptr)
{
		return ({ QByteArray t9c2133 = static_cast<QScriptClass*>(ptr)->QScriptClass::name().toUtf8(); QtScript_PackedString { const_cast<char*>(t9c2133.prepend("WHITESPACE").constData()+10), t9c2133.size()-10 }; });
}

char QScriptClass_SupportsExtension(void* ptr, long long extension)
{
	return static_cast<QScriptClass*>(ptr)->supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

char QScriptClass_SupportsExtensionDefault(void* ptr, long long extension)
{
		return static_cast<QScriptClass*>(ptr)->QScriptClass::supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

void* QScriptContext_ThrowError(void* ptr, long long error, struct QtScript_PackedString text)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(static_cast<QScriptContext::Error>(error), QString::fromUtf8(text.data, text.len)));
}

void* QScriptContext_ThrowError2(void* ptr, struct QtScript_PackedString text)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(QString::fromUtf8(text.data, text.len)));
}

void* QScriptContext_ThrowValue(void* ptr, void* value)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwValue(*static_cast<QScriptValue*>(value)));
}

void QScriptContext_SetActivationObject(void* ptr, void* activation)
{
	static_cast<QScriptContext*>(ptr)->setActivationObject(*static_cast<QScriptValue*>(activation));
}

void QScriptContext_SetThisObject(void* ptr, void* thisObject)
{
	static_cast<QScriptContext*>(ptr)->setThisObject(*static_cast<QScriptValue*>(thisObject));
}

void QScriptContext_DestroyQScriptContext(void* ptr)
{
	static_cast<QScriptContext*>(ptr)->~QScriptContext();
}

long long QScriptContext_State(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->state();
}

void* QScriptContext_ParentContext(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->parentContext();
}

void* QScriptContext_Engine(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->engine();
}

void* QScriptContext_ActivationObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->activationObject());
}

void* QScriptContext_Argument(void* ptr, int index)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argument(index));
}

void* QScriptContext_ArgumentsObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argumentsObject());
}

void* QScriptContext_Callee(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->callee());
}

void* QScriptContext_ThisObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->thisObject());
}

struct QtScript_PackedString QScriptContext_ToString(void* ptr)
{
	return ({ QByteArray t23b7ea = static_cast<QScriptContext*>(ptr)->toString().toUtf8(); QtScript_PackedString { const_cast<char*>(t23b7ea.prepend("WHITESPACE").constData()+10), t23b7ea.size()-10 }; });
}

struct QtScript_PackedString QScriptContext_Backtrace(void* ptr)
{
	return ({ QByteArray t5c5475 = static_cast<QScriptContext*>(ptr)->backtrace().join("|").toUtf8(); QtScript_PackedString { const_cast<char*>(t5c5475.prepend("WHITESPACE").constData()+10), t5c5475.size()-10 }; });
}

char QScriptContext_IsCalledAsConstructor(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->isCalledAsConstructor();
}

int QScriptContext_ArgumentCount(void* ptr)
{
	return static_cast<QScriptContext*>(ptr)->argumentCount();
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

void QScriptContextInfo_DestroyQScriptContextInfo(void* ptr)
{
	static_cast<QScriptContextInfo*>(ptr)->~QScriptContextInfo();
}

long long QScriptContextInfo_FunctionType(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionType();
}

struct QtScript_PackedString QScriptContextInfo_FileName(void* ptr)
{
	return ({ QByteArray tb20c89 = static_cast<QScriptContextInfo*>(ptr)->fileName().toUtf8(); QtScript_PackedString { const_cast<char*>(tb20c89.prepend("WHITESPACE").constData()+10), tb20c89.size()-10 }; });
}

struct QtScript_PackedString QScriptContextInfo_FunctionName(void* ptr)
{
	return ({ QByteArray tacfaa1 = static_cast<QScriptContextInfo*>(ptr)->functionName().toUtf8(); QtScript_PackedString { const_cast<char*>(tacfaa1.prepend("WHITESPACE").constData()+10), tacfaa1.size()-10 }; });
}

struct QtScript_PackedString QScriptContextInfo_FunctionParameterNames(void* ptr)
{
	return ({ QByteArray te54f54 = static_cast<QScriptContextInfo*>(ptr)->functionParameterNames().join("|").toUtf8(); QtScript_PackedString { const_cast<char*>(te54f54.prepend("WHITESPACE").constData()+10), te54f54.size()-10 }; });
}

char QScriptContextInfo_IsNull(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->isNull();
}

int QScriptContextInfo_FunctionEndLineNumber(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionEndLineNumber();
}

int QScriptContextInfo_FunctionMetaIndex(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionMetaIndex();
}

int QScriptContextInfo_FunctionStartLineNumber(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->functionStartLineNumber();
}

int QScriptContextInfo_LineNumber(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->lineNumber();
}

long long QScriptContextInfo_ScriptId(void* ptr)
{
	return static_cast<QScriptContextInfo*>(ptr)->scriptId();
}

class MyQScriptEngine: public QScriptEngine
{
public:
	MyQScriptEngine() : QScriptEngine() {QScriptEngine_QScriptEngine_QRegisterMetaType();};
	MyQScriptEngine(QObject *parent) : QScriptEngine(parent) {QScriptEngine_QScriptEngine_QRegisterMetaType();};
	void Signal_SignalHandlerException(const QScriptValue & exception) { callbackQScriptEngine_SignalHandlerException(this, const_cast<QScriptValue*>(&exception)); };
	 ~MyQScriptEngine() { callbackQScriptEngine_DestroyQScriptEngine(this); };
	bool event(QEvent * e) { return callbackQScriptEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptEngine_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScriptEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScriptEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQScriptEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScriptEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScript_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScriptEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScriptEngine_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScriptEngine*)

int QScriptEngine_QScriptEngine_QRegisterMetaType(){qRegisterMetaType<QScriptEngine*>(); return qRegisterMetaType<MyQScriptEngine*>();}

void* QScriptEngine_PushContext(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->pushContext();
}

void* QScriptEngine_NewQScriptEngine()
{
	return new MyQScriptEngine();
}

void* QScriptEngine_NewQScriptEngine2(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScriptEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQScriptEngine(static_cast<QObject*>(parent));
	}
}

void* QScriptEngine_ToStringHandle(void* ptr, struct QtScript_PackedString str)
{
	return new QScriptString(static_cast<QScriptEngine*>(ptr)->toStringHandle(QString::fromUtf8(str.data, str.len)));
}

void* QScriptEngine_QScriptEngine_CheckSyntax(struct QtScript_PackedString program)
{
	return new QScriptSyntaxCheckResult(QScriptEngine::checkSyntax(QString::fromUtf8(program.data, program.len)));
}

void* QScriptEngine_Evaluate2(void* ptr, void* program)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->evaluate(*static_cast<QScriptProgram*>(program)));
}

void* QScriptEngine_Evaluate(void* ptr, struct QtScript_PackedString program, struct QtScript_PackedString fileName, int lineNumber)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->evaluate(QString::fromUtf8(program.data, program.len), QString::fromUtf8(fileName.data, fileName.len), lineNumber));
}

void* QScriptEngine_ImportExtension(void* ptr, struct QtScript_PackedString extension)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->importExtension(QString::fromUtf8(extension.data, extension.len)));
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

void* QScriptEngine_NewRegExp2(void* ptr, struct QtScript_PackedString pattern, struct QtScript_PackedString flags)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->newRegExp(QString::fromUtf8(pattern.data, pattern.len), QString::fromUtf8(flags.data, flags.len)));
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

void* QScriptEngine_ToObject(void* ptr, void* value)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->toObject(*static_cast<QScriptValue*>(value)));
}

void* QScriptEngine_UndefinedValue(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->undefinedValue());
}

void QScriptEngine_AbortEvaluation(void* ptr, void* result)
{
	static_cast<QScriptEngine*>(ptr)->abortEvaluation(*static_cast<QScriptValue*>(result));
}

void QScriptEngine_ClearExceptions(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->clearExceptions();
}

void QScriptEngine_CollectGarbage(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->collectGarbage();
}

void QScriptEngine_InstallTranslatorFunctions(void* ptr, void* object)
{
	static_cast<QScriptEngine*>(ptr)->installTranslatorFunctions(*static_cast<QScriptValue*>(object));
}

void QScriptEngine_PopContext(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->popContext();
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

void QScriptEngine_DestroyQScriptEngine(void* ptr)
{
	static_cast<QScriptEngine*>(ptr)->~QScriptEngine();
}

void QScriptEngine_DestroyQScriptEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScriptEngine_CurrentContext(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->currentContext();
}

void* QScriptEngine_Agent(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->agent();
}

void* QScriptEngine_DefaultPrototype(void* ptr, int metaTypeId)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->defaultPrototype(metaTypeId));
}

void* QScriptEngine_GlobalObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->globalObject());
}

void* QScriptEngine_UncaughtException(void* ptr)
{
	return new QScriptValue(static_cast<QScriptEngine*>(ptr)->uncaughtException());
}

struct QtScript_PackedString QScriptEngine_AvailableExtensions(void* ptr)
{
	return ({ QByteArray tf41e6e = static_cast<QScriptEngine*>(ptr)->availableExtensions().join("|").toUtf8(); QtScript_PackedString { const_cast<char*>(tf41e6e.prepend("WHITESPACE").constData()+10), tf41e6e.size()-10 }; });
}

struct QtScript_PackedString QScriptEngine_ImportedExtensions(void* ptr)
{
	return ({ QByteArray td77625 = static_cast<QScriptEngine*>(ptr)->importedExtensions().join("|").toUtf8(); QtScript_PackedString { const_cast<char*>(td77625.prepend("WHITESPACE").constData()+10), td77625.size()-10 }; });
}

struct QtScript_PackedString QScriptEngine_UncaughtExceptionBacktrace(void* ptr)
{
	return ({ QByteArray t6a16bc = static_cast<QScriptEngine*>(ptr)->uncaughtExceptionBacktrace().join("|").toUtf8(); QtScript_PackedString { const_cast<char*>(t6a16bc.prepend("WHITESPACE").constData()+10), t6a16bc.size()-10 }; });
}

char QScriptEngine_HasUncaughtException(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->hasUncaughtException();
}

char QScriptEngine_IsEvaluating(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->isEvaluating();
}

int QScriptEngine_ProcessEventsInterval(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->processEventsInterval();
}

int QScriptEngine_UncaughtExceptionLineNumber(void* ptr)
{
	return static_cast<QScriptEngine*>(ptr)->uncaughtExceptionLineNumber();
}

void* QScriptEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScriptEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScriptEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QScriptEngine___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptEngine___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngine___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptEngine___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptEngine___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptEngine___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScriptEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScriptEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QScriptEngine*>(ptr)->QScriptEngine::event(static_cast<QEvent*>(e));
}

char QScriptEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QScriptEngine*>(ptr)->QScriptEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QScriptEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScriptEngine*>(ptr)->QScriptEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScriptEngine*>(ptr)->QScriptEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScriptEngine*>(ptr)->QScriptEngine::customEvent(static_cast<QEvent*>(event));
}

void QScriptEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QScriptEngine*>(ptr)->QScriptEngine::deleteLater();
}

void QScriptEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScriptEngine*>(ptr)->QScriptEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScriptEngine*>(ptr)->QScriptEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QScriptEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QScriptEngine*>(ptr)->QScriptEngine::metaObject());
}

class MyQScriptEngineAgent: public QScriptEngineAgent
{
public:
	MyQScriptEngineAgent(QScriptEngine *engine) : QScriptEngineAgent(engine) {};
	QVariant extension(QScriptEngineAgent::Extension extension, const QVariant & argument) { return *static_cast<QVariant*>(callbackQScriptEngineAgent_Extension(this, extension, const_cast<QVariant*>(&argument))); };
	void contextPop() { callbackQScriptEngineAgent_ContextPop(this); };
	void contextPush() { callbackQScriptEngineAgent_ContextPush(this); };
	void exceptionCatch(qint64 scriptId, const QScriptValue & exception) { callbackQScriptEngineAgent_ExceptionCatch(this, scriptId, const_cast<QScriptValue*>(&exception)); };
	void exceptionThrow(qint64 scriptId, const QScriptValue & exception, bool hasHandler) { callbackQScriptEngineAgent_ExceptionThrow(this, scriptId, const_cast<QScriptValue*>(&exception), hasHandler); };
	void functionEntry(qint64 scriptId) { callbackQScriptEngineAgent_FunctionEntry(this, scriptId); };
	void functionExit(qint64 scriptId, const QScriptValue & returnValue) { callbackQScriptEngineAgent_FunctionExit(this, scriptId, const_cast<QScriptValue*>(&returnValue)); };
	void positionChange(qint64 scriptId, int lineNumber, int columnNumber) { callbackQScriptEngineAgent_PositionChange(this, scriptId, lineNumber, columnNumber); };
	void scriptLoad(qint64 id, const QString & program, const QString & fileName, int baseLineNumber) { QByteArray t81d9ae = program.toUtf8(); QtScript_PackedString programPacked = { const_cast<char*>(t81d9ae.prepend("WHITESPACE").constData()+10), t81d9ae.size()-10 };QByteArray td83e09 = fileName.toUtf8(); QtScript_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQScriptEngineAgent_ScriptLoad(this, id, programPacked, fileNamePacked, baseLineNumber); };
	void scriptUnload(qint64 id) { callbackQScriptEngineAgent_ScriptUnload(this, id); };
	 ~MyQScriptEngineAgent() { callbackQScriptEngineAgent_DestroyQScriptEngineAgent(this); };
	bool supportsExtension(QScriptEngineAgent::Extension extension) const { return callbackQScriptEngineAgent_SupportsExtension(const_cast<void*>(static_cast<const void*>(this)), extension) != 0; };
};

void* QScriptEngineAgent_NewQScriptEngineAgent(void* engine)
{
	return new MyQScriptEngineAgent(static_cast<QScriptEngine*>(engine));
}

void* QScriptEngineAgent_Extension(void* ptr, long long extension, void* argument)
{
	return new QVariant(static_cast<QScriptEngineAgent*>(ptr)->extension(static_cast<QScriptEngineAgent::Extension>(extension), *static_cast<QVariant*>(argument)));
}

void* QScriptEngineAgent_ExtensionDefault(void* ptr, long long extension, void* argument)
{
		return new QVariant(static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::extension(static_cast<QScriptEngineAgent::Extension>(extension), *static_cast<QVariant*>(argument)));
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

void QScriptEngineAgent_ScriptLoad(void* ptr, long long id, struct QtScript_PackedString program, struct QtScript_PackedString fileName, int baseLineNumber)
{
	static_cast<QScriptEngineAgent*>(ptr)->scriptLoad(id, QString::fromUtf8(program.data, program.len), QString::fromUtf8(fileName.data, fileName.len), baseLineNumber);
}

void QScriptEngineAgent_ScriptLoadDefault(void* ptr, long long id, struct QtScript_PackedString program, struct QtScript_PackedString fileName, int baseLineNumber)
{
		static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::scriptLoad(id, QString::fromUtf8(program.data, program.len), QString::fromUtf8(fileName.data, fileName.len), baseLineNumber);
}

void QScriptEngineAgent_ScriptUnload(void* ptr, long long id)
{
	static_cast<QScriptEngineAgent*>(ptr)->scriptUnload(id);
}

void QScriptEngineAgent_ScriptUnloadDefault(void* ptr, long long id)
{
		static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::scriptUnload(id);
}

void QScriptEngineAgent_DestroyQScriptEngineAgent(void* ptr)
{
	static_cast<QScriptEngineAgent*>(ptr)->~QScriptEngineAgent();
}

void QScriptEngineAgent_DestroyQScriptEngineAgentDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScriptEngineAgent_Engine(void* ptr)
{
	return static_cast<QScriptEngineAgent*>(ptr)->engine();
}

char QScriptEngineAgent_SupportsExtension(void* ptr, long long extension)
{
	return static_cast<QScriptEngineAgent*>(ptr)->supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

char QScriptEngineAgent_SupportsExtensionDefault(void* ptr, long long extension)
{
		return static_cast<QScriptEngineAgent*>(ptr)->QScriptEngineAgent::supportsExtension(static_cast<QScriptEngineAgent::Extension>(extension));
}

class MyQScriptExtensionPlugin: public QScriptExtensionPlugin
{
public:
	MyQScriptExtensionPlugin(QObject *parent = Q_NULLPTR) : QScriptExtensionPlugin(parent) {QScriptExtensionPlugin_QScriptExtensionPlugin_QRegisterMetaType();};
	void initialize(const QString & key, QScriptEngine * engine) { QByteArray ta62f22 = key.toUtf8(); QtScript_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQScriptExtensionPlugin_Initialize(this, keyPacked, engine); };
	QStringList keys() const { return ({ QtScript_PackedString tempVal = callbackQScriptExtensionPlugin_Keys(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	bool event(QEvent * e) { return callbackQScriptExtensionPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScriptExtensionPlugin_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScriptExtensionPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScriptExtensionPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScriptExtensionPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQScriptExtensionPlugin_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScriptExtensionPlugin_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScriptExtensionPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScript_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScriptExtensionPlugin_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScriptExtensionPlugin_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScriptExtensionPlugin_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScriptExtensionPlugin*)

int QScriptExtensionPlugin_QScriptExtensionPlugin_QRegisterMetaType(){qRegisterMetaType<QScriptExtensionPlugin*>(); return qRegisterMetaType<MyQScriptExtensionPlugin*>();}

void* QScriptExtensionPlugin_NewQScriptExtensionPlugin(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScriptExtensionPlugin(static_cast<QWindow*>(parent));
	} else {
		return new MyQScriptExtensionPlugin(static_cast<QObject*>(parent));
	}
}

void QScriptExtensionPlugin_Initialize(void* ptr, struct QtScript_PackedString key, void* engine)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->initialize(QString::fromUtf8(key.data, key.len), static_cast<QScriptEngine*>(engine));
}

void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(void* ptr)
{
	static_cast<QScriptExtensionPlugin*>(ptr)->~QScriptExtensionPlugin();
}

void* QScriptExtensionPlugin_SetupPackage(void* ptr, struct QtScript_PackedString key, void* engine)
{
	return new QScriptValue(static_cast<QScriptExtensionPlugin*>(ptr)->setupPackage(QString::fromUtf8(key.data, key.len), static_cast<QScriptEngine*>(engine)));
}

struct QtScript_PackedString QScriptExtensionPlugin_Keys(void* ptr)
{
	return ({ QByteArray tcd9b88 = static_cast<QScriptExtensionPlugin*>(ptr)->keys().join("|").toUtf8(); QtScript_PackedString { const_cast<char*>(tcd9b88.prepend("WHITESPACE").constData()+10), tcd9b88.size()-10 }; });
}

void* QScriptExtensionPlugin___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScriptExtensionPlugin___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScriptExtensionPlugin___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QScriptExtensionPlugin___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptExtensionPlugin___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptExtensionPlugin___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptExtensionPlugin___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptExtensionPlugin___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptExtensionPlugin___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptExtensionPlugin___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScriptExtensionPlugin___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptExtensionPlugin___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScriptExtensionPlugin___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScriptExtensionPlugin___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScriptExtensionPlugin___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScriptExtensionPlugin_EventDefault(void* ptr, void* e)
{
		return static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::event(static_cast<QEvent*>(e));
}

char QScriptExtensionPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QScriptExtensionPlugin_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QScriptExtensionPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptExtensionPlugin_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::customEvent(static_cast<QEvent*>(event));
}

void QScriptExtensionPlugin_DeleteLaterDefault(void* ptr)
{
		static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::deleteLater();
}

void QScriptExtensionPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScriptExtensionPlugin_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScriptExtensionPlugin*>(ptr)->QScriptExtensionPlugin::timerEvent(static_cast<QTimerEvent*>(event));
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

void* QScriptProgram_NewQScriptProgram2(struct QtScript_PackedString sourceCode, struct QtScript_PackedString fileName, int firstLineNumber)
{
	return new QScriptProgram(QString::fromUtf8(sourceCode.data, sourceCode.len), QString::fromUtf8(fileName.data, fileName.len), firstLineNumber);
}

void QScriptProgram_DestroyQScriptProgram(void* ptr)
{
	static_cast<QScriptProgram*>(ptr)->~QScriptProgram();
}

struct QtScript_PackedString QScriptProgram_FileName(void* ptr)
{
	return ({ QByteArray t3d5784 = static_cast<QScriptProgram*>(ptr)->fileName().toUtf8(); QtScript_PackedString { const_cast<char*>(t3d5784.prepend("WHITESPACE").constData()+10), t3d5784.size()-10 }; });
}

struct QtScript_PackedString QScriptProgram_SourceCode(void* ptr)
{
	return ({ QByteArray t385173 = static_cast<QScriptProgram*>(ptr)->sourceCode().toUtf8(); QtScript_PackedString { const_cast<char*>(t385173.prepend("WHITESPACE").constData()+10), t385173.size()-10 }; });
}

char QScriptProgram_IsNull(void* ptr)
{
	return static_cast<QScriptProgram*>(ptr)->isNull();
}

int QScriptProgram_FirstLineNumber(void* ptr)
{
	return static_cast<QScriptProgram*>(ptr)->firstLineNumber();
}

void* QScriptString_NewQScriptString()
{
	return new QScriptString();
}

void* QScriptString_NewQScriptString2(void* other)
{
	return new QScriptString(*static_cast<QScriptString*>(other));
}

void QScriptString_DestroyQScriptString(void* ptr)
{
	static_cast<QScriptString*>(ptr)->~QScriptString();
}

struct QtScript_PackedString QScriptString_ToString(void* ptr)
{
	return ({ QByteArray t12c6d8 = static_cast<QScriptString*>(ptr)->toString().toUtf8(); QtScript_PackedString { const_cast<char*>(t12c6d8.prepend("WHITESPACE").constData()+10), t12c6d8.size()-10 }; });
}

char QScriptString_IsValid(void* ptr)
{
	return static_cast<QScriptString*>(ptr)->isValid();
}

unsigned int QScriptString_ToArrayIndex(void* ptr, char ok)
{
	Q_UNUSED(ok);
	return static_cast<QScriptString*>(ptr)->toArrayIndex(NULL);
}

void* QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(void* other)
{
	return new QScriptSyntaxCheckResult(*static_cast<QScriptSyntaxCheckResult*>(other));
}

void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(void* ptr)
{
	static_cast<QScriptSyntaxCheckResult*>(ptr)->~QScriptSyntaxCheckResult();
}

struct QtScript_PackedString QScriptSyntaxCheckResult_ErrorMessage(void* ptr)
{
	return ({ QByteArray t5bd6f3 = static_cast<QScriptSyntaxCheckResult*>(ptr)->errorMessage().toUtf8(); QtScript_PackedString { const_cast<char*>(t5bd6f3.prepend("WHITESPACE").constData()+10), t5bd6f3.size()-10 }; });
}

long long QScriptSyntaxCheckResult_State(void* ptr)
{
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->state();
}

int QScriptSyntaxCheckResult_ErrorColumnNumber(void* ptr)
{
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorColumnNumber();
}

int QScriptSyntaxCheckResult_ErrorLineNumber(void* ptr)
{
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorLineNumber();
}

void* QScriptValue_Construct2(void* ptr, void* arguments)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->construct(*static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_NewQScriptValue()
{
	return new QScriptValue();
}

void* QScriptValue_Call2(void* ptr, void* thisObject, void* arguments)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->call(*static_cast<QScriptValue*>(thisObject), *static_cast<QScriptValue*>(arguments)));
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

void* QScriptValue_NewQScriptValue15(struct QtScript_PackedString value)
{
	return new QScriptValue(QString::fromUtf8(value.data, value.len));
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

void QScriptValue_SetData(void* ptr, void* data)
{
	static_cast<QScriptValue*>(ptr)->setData(*static_cast<QScriptValue*>(data));
}

void QScriptValue_SetProperty3(void* ptr, void* name, void* value, long long flags)
{
	static_cast<QScriptValue*>(ptr)->setProperty(*static_cast<QScriptString*>(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetProperty(void* ptr, struct QtScript_PackedString name, void* value, long long flags)
{
	static_cast<QScriptValue*>(ptr)->setProperty(QString::fromUtf8(name.data, name.len), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
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

void QScriptValue_DestroyQScriptValue(void* ptr)
{
	static_cast<QScriptValue*>(ptr)->~QScriptValue();
}

void* QScriptValue_ToDateTime(void* ptr)
{
	return new QDateTime(static_cast<QScriptValue*>(ptr)->toDateTime());
}

void* QScriptValue_ToQObject(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toQObject();
}

void* QScriptValue_ToRegExp(void* ptr)
{
	return new QRegExp(static_cast<QScriptValue*>(ptr)->toRegExp());
}

void* QScriptValue_ScriptClass(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->scriptClass();
}

void* QScriptValue_Engine(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->engine();
}

void* QScriptValue_Data(void* ptr)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->data());
}

void* QScriptValue_Property3(void* ptr, void* name, long long mode)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Property(void* ptr, struct QtScript_PackedString name, long long mode)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(QString::fromUtf8(name.data, name.len), static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Property2(void* ptr, unsigned int arrayIndex, long long mode)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(arrayIndex, static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Prototype(void* ptr)
{
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->prototype());
}

long long QScriptValue_PropertyFlags2(void* ptr, void* name, long long mode)
{
	return static_cast<QScriptValue*>(ptr)->propertyFlags(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

long long QScriptValue_PropertyFlags(void* ptr, struct QtScript_PackedString name, long long mode)
{
	return static_cast<QScriptValue*>(ptr)->propertyFlags(QString::fromUtf8(name.data, name.len), static_cast<QScriptValue::ResolveFlag>(mode));
}

struct QtScript_PackedString QScriptValue_ToString(void* ptr)
{
	return ({ QByteArray t3ccbbf = static_cast<QScriptValue*>(ptr)->toString().toUtf8(); QtScript_PackedString { const_cast<char*>(t3ccbbf.prepend("WHITESPACE").constData()+10), t3ccbbf.size()-10 }; });
}

void* QScriptValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QScriptValue*>(ptr)->toVariant());
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

char QScriptValue_StrictlyEquals(void* ptr, void* other)
{
	return static_cast<QScriptValue*>(ptr)->strictlyEquals(*static_cast<QScriptValue*>(other));
}

char QScriptValue_ToBool(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toBool();
}

void* QScriptValue_ToQMetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScriptValue*>(ptr)->toQMetaObject());
}

int QScriptValue_ToInt32(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toInt32();
}

unsigned short QScriptValue_ToUInt16(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toUInt16();
}

unsigned int QScriptValue_ToUInt32(void* ptr)
{
	return static_cast<QScriptValue*>(ptr)->toUInt32();
}

void* QScriptable_Context(void* ptr)
{
	return static_cast<QScriptable*>(ptr)->context();
}

void* QScriptable_Engine(void* ptr)
{
	return static_cast<QScriptable*>(ptr)->engine();
}

void* QScriptable_Argument(void* ptr, int index)
{
	return new QScriptValue(static_cast<QScriptable*>(ptr)->argument(index));
}

void* QScriptable_ThisObject(void* ptr)
{
	return new QScriptValue(static_cast<QScriptable*>(ptr)->thisObject());
}

int QScriptable_ArgumentCount(void* ptr)
{
	return static_cast<QScriptable*>(ptr)->argumentCount();
}

