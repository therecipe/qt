#define protected public
#define private public

#include "qml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QFile>
#include <QFileSelector>
#include <QJSEngine>
#include <QJSValue>
#include <QLatin1String>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNetworkAccessManager>
#include <QObject>
#include <QQmlAbstractUrlInterceptor>
#include <QQmlApplicationEngine>
#include <QQmlComponent>
#include <QQmlContext>
#include <QQmlEngine>
#include <QQmlError>
#include <QQmlExpression>
#include <QQmlExtensionPlugin>
#include <QQmlFileSelector>
#include <QQmlImageProviderBase>
#include <QQmlIncubationController>
#include <QQmlIncubator>
#include <QQmlListReference>
#include <QQmlNetworkAccessManagerFactory>
#include <QQmlParserStatus>
#include <QQmlProperty>
#include <QQmlPropertyMap>
#include <QQmlPropertyValueSource>
#include <QQmlScriptString>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>

class MyQJSEngine: public QJSEngine
{
public:
	MyQJSEngine() : QJSEngine() {};
	MyQJSEngine(QObject *parent) : QJSEngine(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQJSEngine_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQJSEngine_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQJSEngine_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQJSEngine_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQJSEngine_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQJSEngine_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQJSEngine_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQJSEngine_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQJSEngine_MetaObject(const_cast<MyQJSEngine*>(this), this->objectName().toUtf8().data())); };
};

void* QJSEngine_NewQJSEngine()
{
	return new MyQJSEngine();
}

void* QJSEngine_NewQJSEngine2(void* parent)
{
	return new MyQJSEngine(static_cast<QObject*>(parent));
}

void QJSEngine_CollectGarbage(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->collectGarbage();
}

void* QJSEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->evaluate(QString(program), QString(fileName), lineNumber));
}

void* QJSEngine_GlobalObject(void* ptr)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->globalObject());
}

void QJSEngine_InstallExtensions(void* ptr, int extensions, void* object)
{
	static_cast<QJSEngine*>(ptr)->installExtensions(static_cast<QJSEngine::Extension>(extensions), *static_cast<QJSValue*>(object));
}

void* QJSEngine_NewObject(void* ptr)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newObject());
}

void* QJSEngine_NewQObject(void* ptr, void* object)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newQObject(static_cast<QObject*>(object)));
}

void QJSEngine_DestroyQJSEngine(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->~QJSEngine();
}

void QJSEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QJSEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QJSEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QJSEngine*>(ptr)->QJSEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QJSEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QJSEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QJSEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QJSEngine*>(ptr)->QJSEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QJSEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QJSEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QJSEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QJSEngine*>(ptr)->QJSEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QJSEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QJSEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QJSEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QJSEngine*>(ptr)->QJSEngine::customEvent(static_cast<QEvent*>(event));
}

void QJSEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QJSEngine*>(ptr), "deleteLater");
}

void QJSEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->QJSEngine::deleteLater();
}

void QJSEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QJSEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QJSEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QJSEngine*>(ptr)->QJSEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QJSEngine_Event(void* ptr, void* e)
{
	return static_cast<QJSEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QJSEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QJSEngine*>(ptr)->QJSEngine::event(static_cast<QEvent*>(e));
}

int QJSEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QJSEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QJSEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QJSEngine*>(ptr)->QJSEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QJSEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QJSEngine*>(ptr)->metaObject());
}

void* QJSEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QJSEngine*>(ptr)->QJSEngine::metaObject());
}

void* QJSValue_NewQJSValue3(void* other)
{
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue(int value)
{
	return new QJSValue(static_cast<QJSValue::SpecialValue>(value));
}

void* QJSValue_NewQJSValue4(int value)
{
	return new QJSValue(value != 0);
}

void* QJSValue_NewQJSValue2(void* other)
{
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue9(void* value)
{
	return new QJSValue(*static_cast<QLatin1String*>(value));
}

void* QJSValue_NewQJSValue8(char* value)
{
	return new QJSValue(QString(value));
}

void* QJSValue_NewQJSValue10(char* value)
{
	return new QJSValue(const_cast<const char*>(value));
}

void* QJSValue_NewQJSValue5(int value)
{
	return new QJSValue(value);
}

int QJSValue_DeleteProperty(void* ptr, char* name)
{
	return static_cast<QJSValue*>(ptr)->deleteProperty(QString(name));
}

int QJSValue_Equals(void* ptr, void* other)
{
	return static_cast<QJSValue*>(ptr)->equals(*static_cast<QJSValue*>(other));
}

int QJSValue_HasOwnProperty(void* ptr, char* name)
{
	return static_cast<QJSValue*>(ptr)->hasOwnProperty(QString(name));
}

int QJSValue_HasProperty(void* ptr, char* name)
{
	return static_cast<QJSValue*>(ptr)->hasProperty(QString(name));
}

int QJSValue_IsArray(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isArray();
}

int QJSValue_IsBool(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isBool();
}

int QJSValue_IsCallable(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isCallable();
}

int QJSValue_IsDate(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isDate();
}

int QJSValue_IsError(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isError();
}

int QJSValue_IsNull(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isNull();
}

int QJSValue_IsNumber(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isNumber();
}

int QJSValue_IsObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isObject();
}

int QJSValue_IsQObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isQObject();
}

int QJSValue_IsRegExp(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isRegExp();
}

int QJSValue_IsString(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isString();
}

int QJSValue_IsUndefined(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isUndefined();
}

int QJSValue_IsVariant(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isVariant();
}

void* QJSValue_Property(void* ptr, char* name)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(QString(name)));
}

void* QJSValue_Prototype(void* ptr)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->prototype());
}

void QJSValue_SetProperty(void* ptr, char* name, void* value)
{
	static_cast<QJSValue*>(ptr)->setProperty(QString(name), *static_cast<QJSValue*>(value));
}

void QJSValue_SetPrototype(void* ptr, void* prototype)
{
	static_cast<QJSValue*>(ptr)->setPrototype(*static_cast<QJSValue*>(prototype));
}

int QJSValue_StrictlyEquals(void* ptr, void* other)
{
	return static_cast<QJSValue*>(ptr)->strictlyEquals(*static_cast<QJSValue*>(other));
}

int QJSValue_ToBool(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toBool();
}

void* QJSValue_ToDateTime(void* ptr)
{
	return new QDateTime(static_cast<QJSValue*>(ptr)->toDateTime());
}

void* QJSValue_ToQObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toQObject();
}

char* QJSValue_ToString(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toString().toUtf8().data();
}

void* QJSValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJSValue*>(ptr)->toVariant());
}

void QJSValue_DestroyQJSValue(void* ptr)
{
	static_cast<QJSValue*>(ptr)->~QJSValue();
}

class MyQQmlAbstractUrlInterceptor: public QQmlAbstractUrlInterceptor
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQmlAbstractUrlInterceptor() : QQmlAbstractUrlInterceptor() {};
	QUrl intercept(const QUrl & url, QQmlAbstractUrlInterceptor::DataType ty) { return *static_cast<QUrl*>(callbackQQmlAbstractUrlInterceptor_Intercept(this, this->objectNameAbs().toUtf8().data(), new QUrl(url), ty)); };
};

void* QQmlAbstractUrlInterceptor_NewQQmlAbstractUrlInterceptor()
{
	return new MyQQmlAbstractUrlInterceptor();
}

void* QQmlAbstractUrlInterceptor_Intercept(void* ptr, void* url, int ty)
{
	return new QUrl(static_cast<QQmlAbstractUrlInterceptor*>(ptr)->intercept(*static_cast<QUrl*>(url), static_cast<QQmlAbstractUrlInterceptor::DataType>(ty)));
}

void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(void* ptr)
{
	static_cast<QQmlAbstractUrlInterceptor*>(ptr)->~QQmlAbstractUrlInterceptor();
}

char* QQmlAbstractUrlInterceptor_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlAbstractUrlInterceptor*>(static_cast<QQmlAbstractUrlInterceptor*>(ptr))) {
		return static_cast<MyQQmlAbstractUrlInterceptor*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlAbstractUrlInterceptor_BASE").toUtf8().data();
}

void QQmlAbstractUrlInterceptor_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlAbstractUrlInterceptor*>(static_cast<QQmlAbstractUrlInterceptor*>(ptr))) {
		static_cast<MyQQmlAbstractUrlInterceptor*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQmlApplicationEngine: public QQmlApplicationEngine
{
public:
	MyQQmlApplicationEngine(QObject *parent) : QQmlApplicationEngine(parent) {};
	MyQQmlApplicationEngine(const QString &filePath, QObject *parent) : QQmlApplicationEngine(filePath, parent) {};
	MyQQmlApplicationEngine(const QUrl &url, QObject *parent) : QQmlApplicationEngine(url, parent) {};
	void load(const QString & filePath) { callbackQQmlApplicationEngine_Load2(this, this->objectName().toUtf8().data(), filePath.toUtf8().data()); };
	void load(const QUrl & url) { callbackQQmlApplicationEngine_Load(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void loadData(const QByteArray & data, const QUrl & url) { callbackQQmlApplicationEngine_LoadData(this, this->objectName().toUtf8().data(), QString(data).toUtf8().data(), new QUrl(url)); };
	void Signal_ObjectCreated(QObject * object, const QUrl & url) { callbackQQmlApplicationEngine_ObjectCreated(this, this->objectName().toUtf8().data(), object, new QUrl(url)); };
	bool event(QEvent * e) { return callbackQQmlApplicationEngine_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQQmlApplicationEngine_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlApplicationEngine_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlApplicationEngine_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlApplicationEngine_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlApplicationEngine_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlApplicationEngine_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlApplicationEngine_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlApplicationEngine_MetaObject(const_cast<MyQQmlApplicationEngine*>(this), this->objectName().toUtf8().data())); };
};

void* QQmlApplicationEngine_NewQQmlApplicationEngine(void* parent)
{
	return new MyQQmlApplicationEngine(static_cast<QObject*>(parent));
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine3(char* filePath, void* parent)
{
	return new MyQQmlApplicationEngine(QString(filePath), static_cast<QObject*>(parent));
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine2(void* url, void* parent)
{
	return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QObject*>(parent));
}

void QQmlApplicationEngine_Load2(void* ptr, char* filePath)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QString, QString(filePath)));
}

void QQmlApplicationEngine_Load(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_LoadData(void* ptr, char* data, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "loadData", Q_ARG(QByteArray, QByteArray(data)), Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_ConnectObjectCreated(void* ptr)
{
	QObject::connect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));
}

void QQmlApplicationEngine_DisconnectObjectCreated(void* ptr)
{
	QObject::disconnect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));
}

void QQmlApplicationEngine_ObjectCreated(void* ptr, void* object, void* url)
{
	static_cast<QQmlApplicationEngine*>(ptr)->objectCreated(static_cast<QObject*>(object), *static_cast<QUrl*>(url));
}

void QQmlApplicationEngine_DestroyQQmlApplicationEngine(void* ptr)
{
	static_cast<QQmlApplicationEngine*>(ptr)->~QQmlApplicationEngine();
}

int QQmlApplicationEngine_Event(void* ptr, void* e)
{
	return static_cast<QQmlApplicationEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlApplicationEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::event(static_cast<QEvent*>(e));
}

void QQmlApplicationEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlApplicationEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlApplicationEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlApplicationEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlApplicationEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlApplicationEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlApplicationEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlApplicationEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlApplicationEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlApplicationEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlApplicationEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlApplicationEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::customEvent(static_cast<QEvent*>(event));
}

void QQmlApplicationEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "deleteLater");
}

void QQmlApplicationEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::deleteLater();
}

void QQmlApplicationEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlApplicationEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlApplicationEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlApplicationEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlApplicationEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlApplicationEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlApplicationEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlApplicationEngine*>(ptr)->metaObject());
}

void* QQmlApplicationEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::metaObject());
}

class MyQQmlComponent: public QQmlComponent
{
public:
	MyQQmlComponent(QQmlEngine *engine, QObject *parent) : QQmlComponent(engine, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QString &fileName, CompilationMode mode, QObject *parent) : QQmlComponent(engine, fileName, mode, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QString &fileName, QObject *parent) : QQmlComponent(engine, fileName, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QUrl &url, CompilationMode mode, QObject *parent) : QQmlComponent(engine, url, mode, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QUrl &url, QObject *parent) : QQmlComponent(engine, url, parent) {};
	QObject * beginCreate(QQmlContext * publicContext) { return static_cast<QObject*>(callbackQQmlComponent_BeginCreate(this, this->objectName().toUtf8().data(), publicContext)); };
	void completeCreate() { callbackQQmlComponent_CompleteCreate(this, this->objectName().toUtf8().data()); };
	QObject * create(QQmlContext * context) { return static_cast<QObject*>(callbackQQmlComponent_Create(this, this->objectName().toUtf8().data(), context)); };
	void loadUrl(const QUrl & url) { callbackQQmlComponent_LoadUrl(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void loadUrl(const QUrl & url, QQmlComponent::CompilationMode mode) { callbackQQmlComponent_LoadUrl2(this, this->objectName().toUtf8().data(), new QUrl(url), mode); };
	void Signal_ProgressChanged(qreal progress) { callbackQQmlComponent_ProgressChanged(this, this->objectName().toUtf8().data(), static_cast<double>(progress)); };
	void setData(const QByteArray & data, const QUrl & url) { callbackQQmlComponent_SetData(this, this->objectName().toUtf8().data(), QString(data).toUtf8().data(), new QUrl(url)); };
	void Signal_StatusChanged(QQmlComponent::Status status) { callbackQQmlComponent_StatusChanged(this, this->objectName().toUtf8().data(), status); };
	void timerEvent(QTimerEvent * event) { callbackQQmlComponent_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlComponent_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlComponent_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlComponent_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlComponent_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlComponent_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQmlComponent_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlComponent_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlComponent_MetaObject(const_cast<MyQQmlComponent*>(this), this->objectName().toUtf8().data())); };
};

double QQmlComponent_Progress(void* ptr)
{
	return static_cast<double>(static_cast<QQmlComponent*>(ptr)->progress());
}

int QQmlComponent_Status(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->status();
}

void* QQmlComponent_Url(void* ptr)
{
	return new QUrl(static_cast<QQmlComponent*>(ptr)->url());
}

void* QQmlComponent_NewQQmlComponent(void* engine, void* parent)
{
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent4(void* engine, char* fileName, int mode, void* parent)
{
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent3(void* engine, char* fileName, void* parent)
{
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent6(void* engine, void* url, int mode, void* parent)
{
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent5(void* engine, void* url, void* parent)
{
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QObject*>(parent));
}

void* QQmlComponent_BeginCreate(void* ptr, void* publicContext)
{
	return static_cast<QQmlComponent*>(ptr)->beginCreate(static_cast<QQmlContext*>(publicContext));
}

void* QQmlComponent_BeginCreateDefault(void* ptr, void* publicContext)
{
	return static_cast<QQmlComponent*>(ptr)->QQmlComponent::beginCreate(static_cast<QQmlContext*>(publicContext));
}

void QQmlComponent_CompleteCreate(void* ptr)
{
	static_cast<QQmlComponent*>(ptr)->completeCreate();
}

void QQmlComponent_CompleteCreateDefault(void* ptr)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::completeCreate();
}

void* QQmlComponent_Create(void* ptr, void* context)
{
	return static_cast<QQmlComponent*>(ptr)->create(static_cast<QQmlContext*>(context));
}

void* QQmlComponent_CreateDefault(void* ptr, void* context)
{
	return static_cast<QQmlComponent*>(ptr)->QQmlComponent::create(static_cast<QQmlContext*>(context));
}

void QQmlComponent_Create2(void* ptr, void* incubator, void* context, void* forContext)
{
	static_cast<QQmlComponent*>(ptr)->create(*static_cast<QQmlIncubator*>(incubator), static_cast<QQmlContext*>(context), static_cast<QQmlContext*>(forContext));
}

void* QQmlComponent_CreationContext(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->creationContext();
}

int QQmlComponent_IsError(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isError();
}

int QQmlComponent_IsLoading(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isLoading();
}

int QQmlComponent_IsNull(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isNull();
}

int QQmlComponent_IsReady(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isReady();
}

void QQmlComponent_LoadUrl(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_LoadUrl2(void* ptr, void* url, int mode)
{
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, *static_cast<QUrl*>(url)), Q_ARG(QQmlComponent::CompilationMode, static_cast<QQmlComponent::CompilationMode>(mode)));
}

void QQmlComponent_ConnectProgressChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(qreal)>(&QQmlComponent::progressChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(qreal)>(&MyQQmlComponent::Signal_ProgressChanged));
}

void QQmlComponent_DisconnectProgressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(qreal)>(&QQmlComponent::progressChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(qreal)>(&MyQQmlComponent::Signal_ProgressChanged));
}

void QQmlComponent_ProgressChanged(void* ptr, double progress)
{
	static_cast<QQmlComponent*>(ptr)->progressChanged(static_cast<double>(progress));
}

void QQmlComponent_SetData(void* ptr, char* data, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "setData", Q_ARG(QByteArray, QByteArray(data)), Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));
}

void QQmlComponent_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));
}

void QQmlComponent_StatusChanged(void* ptr, int status)
{
	static_cast<QQmlComponent*>(ptr)->statusChanged(static_cast<QQmlComponent::Status>(status));
}

void QQmlComponent_DestroyQQmlComponent(void* ptr)
{
	static_cast<QQmlComponent*>(ptr)->~QQmlComponent();
}

void QQmlComponent_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlComponent*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlComponent_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlComponent_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlComponent*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlComponent_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlComponent_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlComponent*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlComponent_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlComponent_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlComponent*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlComponent_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::customEvent(static_cast<QEvent*>(event));
}

void QQmlComponent_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "deleteLater");
}

void QQmlComponent_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::deleteLater();
}

void QQmlComponent_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlComponent*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlComponent_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlComponent*>(ptr)->QQmlComponent::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlComponent_Event(void* ptr, void* e)
{
	return static_cast<QQmlComponent*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlComponent_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlComponent*>(ptr)->QQmlComponent::event(static_cast<QEvent*>(e));
}

int QQmlComponent_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlComponent*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlComponent_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlComponent*>(ptr)->QQmlComponent::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlComponent_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlComponent*>(ptr)->metaObject());
}

void* QQmlComponent_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlComponent*>(ptr)->QQmlComponent::metaObject());
}

class MyQQmlContext: public QQmlContext
{
public:
	MyQQmlContext(QQmlContext *parentContext, QObject *parent) : QQmlContext(parentContext, parent) {};
	MyQQmlContext(QQmlEngine *engine, QObject *parent) : QQmlContext(engine, parent) {};
	void timerEvent(QTimerEvent * event) { callbackQQmlContext_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlContext_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlContext_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlContext_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlContext_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlContext_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQmlContext_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlContext_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlContext_MetaObject(const_cast<MyQQmlContext*>(this), this->objectName().toUtf8().data())); };
};

void* QQmlContext_NewQQmlContext2(void* parentContext, void* parent)
{
	return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QObject*>(parent));
}

void* QQmlContext_NewQQmlContext(void* engine, void* parent)
{
	return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlContext_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QQmlContext*>(ptr)->baseUrl());
}

void* QQmlContext_ContextObject(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->contextObject();
}

void* QQmlContext_ContextProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QQmlContext*>(ptr)->contextProperty(QString(name)));
}

void* QQmlContext_Engine(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->engine();
}

int QQmlContext_IsValid(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->isValid();
}

char* QQmlContext_NameForObject(void* ptr, void* object)
{
	return static_cast<QQmlContext*>(ptr)->nameForObject(static_cast<QObject*>(object)).toUtf8().data();
}

void* QQmlContext_ParentContext(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->parentContext();
}

void* QQmlContext_ResolvedUrl(void* ptr, void* src)
{
	return new QUrl(static_cast<QQmlContext*>(ptr)->resolvedUrl(*static_cast<QUrl*>(src)));
}

void QQmlContext_SetBaseUrl(void* ptr, void* baseUrl)
{
	static_cast<QQmlContext*>(ptr)->setBaseUrl(*static_cast<QUrl*>(baseUrl));
}

void QQmlContext_SetContextObject(void* ptr, void* object)
{
	static_cast<QQmlContext*>(ptr)->setContextObject(static_cast<QObject*>(object));
}

void QQmlContext_SetContextProperty(void* ptr, char* name, void* value)
{
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), static_cast<QObject*>(value));
}

void QQmlContext_SetContextProperty2(void* ptr, char* name, void* value)
{
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), *static_cast<QVariant*>(value));
}

void QQmlContext_DestroyQQmlContext(void* ptr)
{
	static_cast<QQmlContext*>(ptr)->~QQmlContext();
}

void QQmlContext_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlContext*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlContext_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlContext*>(ptr)->QQmlContext::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlContext_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlContext*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlContext_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlContext*>(ptr)->QQmlContext::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlContext_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlContext*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlContext_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlContext*>(ptr)->QQmlContext::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlContext_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlContext*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlContext_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlContext*>(ptr)->QQmlContext::customEvent(static_cast<QEvent*>(event));
}

void QQmlContext_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlContext*>(ptr), "deleteLater");
}

void QQmlContext_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlContext*>(ptr)->QQmlContext::deleteLater();
}

void QQmlContext_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlContext*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlContext_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlContext*>(ptr)->QQmlContext::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlContext_Event(void* ptr, void* e)
{
	return static_cast<QQmlContext*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlContext_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlContext*>(ptr)->QQmlContext::event(static_cast<QEvent*>(e));
}

int QQmlContext_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlContext*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlContext_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlContext*>(ptr)->QQmlContext::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlContext_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlContext*>(ptr)->metaObject());
}

void* QQmlContext_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlContext*>(ptr)->QQmlContext::metaObject());
}

class MyQQmlEngine: public QQmlEngine
{
public:
	MyQQmlEngine(QObject *parent) : QQmlEngine(parent) {};
	bool event(QEvent * e) { return callbackQQmlEngine_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	void Signal_Quit() { callbackQQmlEngine_Quit(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQmlEngine_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlEngine_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlEngine_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlEngine_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlEngine_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlEngine_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlEngine_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlEngine_MetaObject(const_cast<MyQQmlEngine*>(this), this->objectName().toUtf8().data())); };
};

char* QQmlEngine_OfflineStoragePath(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->offlineStoragePath().toUtf8().data();
}

void QQmlEngine_SetOfflineStoragePath(void* ptr, char* dir)
{
	static_cast<QQmlEngine*>(ptr)->setOfflineStoragePath(QString(dir));
}

void* QQmlEngine_NewQQmlEngine(void* parent)
{
	return new MyQQmlEngine(static_cast<QObject*>(parent));
}

void QQmlEngine_AddImageProvider(void* ptr, char* providerId, void* provider)
{
	static_cast<QQmlEngine*>(ptr)->addImageProvider(QString(providerId), static_cast<QQmlImageProviderBase*>(provider));
}

void QQmlEngine_AddImportPath(void* ptr, char* path)
{
	static_cast<QQmlEngine*>(ptr)->addImportPath(QString(path));
}

void QQmlEngine_AddPluginPath(void* ptr, char* path)
{
	static_cast<QQmlEngine*>(ptr)->addPluginPath(QString(path));
}

void* QQmlEngine_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QQmlEngine*>(ptr)->baseUrl());
}

void QQmlEngine_ClearComponentCache(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->clearComponentCache();
}

void* QQmlEngine_QQmlEngine_ContextForObject(void* object)
{
	return QQmlEngine::contextForObject(static_cast<QObject*>(object));
}

int QQmlEngine_Event(void* ptr, void* e)
{
	return static_cast<QQmlEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlEngine*>(ptr)->QQmlEngine::event(static_cast<QEvent*>(e));
}

void* QQmlEngine_ImageProvider(void* ptr, char* providerId)
{
	return static_cast<QQmlEngine*>(ptr)->imageProvider(QString(providerId));
}

char* QQmlEngine_ImportPathList(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->importPathList().join("|").toUtf8().data();
}

void* QQmlEngine_IncubationController(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->incubationController();
}

void* QQmlEngine_NetworkAccessManager(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->networkAccessManager();
}

void* QQmlEngine_NetworkAccessManagerFactory(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->networkAccessManagerFactory();
}

int QQmlEngine_QQmlEngine_ObjectOwnership(void* object)
{
	return QQmlEngine::objectOwnership(static_cast<QObject*>(object));
}

int QQmlEngine_OutputWarningsToStandardError(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->outputWarningsToStandardError();
}

char* QQmlEngine_PluginPathList(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->pluginPathList().join("|").toUtf8().data();
}

void QQmlEngine_ConnectQuit(void* ptr)
{
	QObject::connect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));
}

void QQmlEngine_DisconnectQuit(void* ptr)
{
	QObject::disconnect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));
}

void QQmlEngine_Quit(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->quit();
}

void QQmlEngine_RemoveImageProvider(void* ptr, char* providerId)
{
	static_cast<QQmlEngine*>(ptr)->removeImageProvider(QString(providerId));
}

void* QQmlEngine_RootContext(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->rootContext();
}

void QQmlEngine_SetBaseUrl(void* ptr, void* url)
{
	static_cast<QQmlEngine*>(ptr)->setBaseUrl(*static_cast<QUrl*>(url));
}

void QQmlEngine_QQmlEngine_SetContextForObject(void* object, void* context)
{
	QQmlEngine::setContextForObject(static_cast<QObject*>(object), static_cast<QQmlContext*>(context));
}

void QQmlEngine_SetImportPathList(void* ptr, char* paths)
{
	static_cast<QQmlEngine*>(ptr)->setImportPathList(QString(paths).split("|", QString::SkipEmptyParts));
}

void QQmlEngine_SetIncubationController(void* ptr, void* controller)
{
	static_cast<QQmlEngine*>(ptr)->setIncubationController(static_cast<QQmlIncubationController*>(controller));
}

void QQmlEngine_SetNetworkAccessManagerFactory(void* ptr, void* factory)
{
	static_cast<QQmlEngine*>(ptr)->setNetworkAccessManagerFactory(static_cast<QQmlNetworkAccessManagerFactory*>(factory));
}

void QQmlEngine_QQmlEngine_SetObjectOwnership(void* object, int ownership)
{
	QQmlEngine::setObjectOwnership(static_cast<QObject*>(object), static_cast<QQmlEngine::ObjectOwnership>(ownership));
}

void QQmlEngine_SetOutputWarningsToStandardError(void* ptr, int enabled)
{
	static_cast<QQmlEngine*>(ptr)->setOutputWarningsToStandardError(enabled != 0);
}

void QQmlEngine_SetPluginPathList(void* ptr, char* paths)
{
	static_cast<QQmlEngine*>(ptr)->setPluginPathList(QString(paths).split("|", QString::SkipEmptyParts));
}

void QQmlEngine_TrimComponentCache(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->trimComponentCache();
}

void QQmlEngine_DestroyQQmlEngine(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->~QQmlEngine();
}

void QQmlEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlEngine*>(ptr)->QQmlEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlEngine*>(ptr)->QQmlEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlEngine*>(ptr)->QQmlEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlEngine*>(ptr)->QQmlEngine::customEvent(static_cast<QEvent*>(event));
}

void QQmlEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlEngine*>(ptr), "deleteLater");
}

void QQmlEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->QQmlEngine::deleteLater();
}

void QQmlEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlEngine*>(ptr)->QQmlEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlEngine*>(ptr)->QQmlEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlEngine*>(ptr)->metaObject());
}

void* QQmlEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlEngine*>(ptr)->QQmlEngine::metaObject());
}

void* QQmlError_NewQQmlError()
{
	return new QQmlError();
}

void* QQmlError_NewQQmlError2(void* other)
{
	return new QQmlError(*static_cast<QQmlError*>(other));
}

int QQmlError_Column(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->column();
}

char* QQmlError_Description(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->description().toUtf8().data();
}

int QQmlError_IsValid(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->isValid();
}

int QQmlError_Line(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->line();
}

void* QQmlError_Object(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->object();
}

void QQmlError_SetColumn(void* ptr, int column)
{
	static_cast<QQmlError*>(ptr)->setColumn(column);
}

void QQmlError_SetDescription(void* ptr, char* description)
{
	static_cast<QQmlError*>(ptr)->setDescription(QString(description));
}

void QQmlError_SetLine(void* ptr, int line)
{
	static_cast<QQmlError*>(ptr)->setLine(line);
}

void QQmlError_SetObject(void* ptr, void* object)
{
	static_cast<QQmlError*>(ptr)->setObject(static_cast<QObject*>(object));
}

void QQmlError_SetUrl(void* ptr, void* url)
{
	static_cast<QQmlError*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

char* QQmlError_ToString(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->toString().toUtf8().data();
}

void* QQmlError_Url(void* ptr)
{
	return new QUrl(static_cast<QQmlError*>(ptr)->url());
}

class MyQQmlExpression: public QQmlExpression
{
public:
	MyQQmlExpression() : QQmlExpression() {};
	MyQQmlExpression(QQmlContext *ctxt, QObject *scope, const QString &expression, QObject *parent) : QQmlExpression(ctxt, scope, expression, parent) {};
	MyQQmlExpression(const QQmlScriptString &script, QQmlContext *ctxt, QObject *scope, QObject *parent) : QQmlExpression(script, ctxt, scope, parent) {};
	void Signal_ValueChanged() { callbackQQmlExpression_ValueChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQmlExpression_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlExpression_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlExpression_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlExpression_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlExpression_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlExpression_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQmlExpression_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlExpression_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlExpression_MetaObject(const_cast<MyQQmlExpression*>(this), this->objectName().toUtf8().data())); };
};

void* QQmlExpression_NewQQmlExpression()
{
	return new MyQQmlExpression();
}

void* QQmlExpression_NewQQmlExpression2(void* ctxt, void* scope, char* expression, void* parent)
{
	return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), QString(expression), static_cast<QObject*>(parent));
}

void* QQmlExpression_NewQQmlExpression3(void* script, void* ctxt, void* scope, void* parent)
{
	return new MyQQmlExpression(*static_cast<QQmlScriptString*>(script), static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), static_cast<QObject*>(parent));
}

void QQmlExpression_ClearError(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->clearError();
}

int QQmlExpression_ColumnNumber(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->columnNumber();
}

void* QQmlExpression_Context(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->context();
}

void* QQmlExpression_Engine(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->engine();
}

void* QQmlExpression_Error(void* ptr)
{
	return new QQmlError(static_cast<QQmlExpression*>(ptr)->error());
}

void* QQmlExpression_Evaluate(void* ptr, int valueIsUndefined)
{
	return new QVariant(static_cast<QQmlExpression*>(ptr)->evaluate(NULL));
}

char* QQmlExpression_Expression(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->expression().toUtf8().data();
}

int QQmlExpression_HasError(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->hasError();
}

int QQmlExpression_LineNumber(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->lineNumber();
}

int QQmlExpression_NotifyOnValueChanged(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->notifyOnValueChanged();
}

void* QQmlExpression_ScopeObject(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->scopeObject();
}

void QQmlExpression_SetExpression(void* ptr, char* expression)
{
	static_cast<QQmlExpression*>(ptr)->setExpression(QString(expression));
}

void QQmlExpression_SetNotifyOnValueChanged(void* ptr, int notifyOnChange)
{
	static_cast<QQmlExpression*>(ptr)->setNotifyOnValueChanged(notifyOnChange != 0);
}

void QQmlExpression_SetSourceLocation(void* ptr, char* url, int line, int column)
{
	static_cast<QQmlExpression*>(ptr)->setSourceLocation(QString(url), line, column);
}

char* QQmlExpression_SourceFile(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->sourceFile().toUtf8().data();
}

void QQmlExpression_ConnectValueChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));
}

void QQmlExpression_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));
}

void QQmlExpression_ValueChanged(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->valueChanged();
}

void QQmlExpression_DestroyQQmlExpression(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->~QQmlExpression();
}

void QQmlExpression_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlExpression*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlExpression_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlExpression*>(ptr)->QQmlExpression::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlExpression_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlExpression*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlExpression_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlExpression*>(ptr)->QQmlExpression::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlExpression_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlExpression*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExpression_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlExpression*>(ptr)->QQmlExpression::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExpression_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlExpression*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlExpression_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlExpression*>(ptr)->QQmlExpression::customEvent(static_cast<QEvent*>(event));
}

void QQmlExpression_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlExpression*>(ptr), "deleteLater");
}

void QQmlExpression_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->QQmlExpression::deleteLater();
}

void QQmlExpression_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlExpression*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExpression_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlExpression*>(ptr)->QQmlExpression::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlExpression_Event(void* ptr, void* e)
{
	return static_cast<QQmlExpression*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlExpression_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlExpression*>(ptr)->QQmlExpression::event(static_cast<QEvent*>(e));
}

int QQmlExpression_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlExpression*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlExpression_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlExpression*>(ptr)->QQmlExpression::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlExpression_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlExpression*>(ptr)->metaObject());
}

void* QQmlExpression_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlExpression*>(ptr)->QQmlExpression::metaObject());
}

class MyQQmlExtensionPlugin: public QQmlExtensionPlugin
{
public:
	MyQQmlExtensionPlugin(QObject *parent) : QQmlExtensionPlugin(parent) {};
	void initializeEngine(QQmlEngine * engine, const char * uri) { callbackQQmlExtensionPlugin_InitializeEngine(this, this->objectName().toUtf8().data(), engine, QString(uri).toUtf8().data()); };
	void registerTypes(const char * uri) { callbackQQmlExtensionPlugin_RegisterTypes(this, this->objectName().toUtf8().data(), QString(uri).toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQmlExtensionPlugin_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlExtensionPlugin_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlExtensionPlugin_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlExtensionPlugin_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlExtensionPlugin_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlExtensionPlugin_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQmlExtensionPlugin_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlExtensionPlugin_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlExtensionPlugin_MetaObject(const_cast<MyQQmlExtensionPlugin*>(this), this->objectName().toUtf8().data())); };
};

void QQmlExtensionPlugin_InitializeEngine(void* ptr, void* engine, char* uri)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

void QQmlExtensionPlugin_InitializeEngineDefault(void* ptr, void* engine, char* uri)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

void* QQmlExtensionPlugin_NewQQmlExtensionPlugin(void* parent)
{
	return new MyQQmlExtensionPlugin(static_cast<QObject*>(parent));
}

void* QQmlExtensionPlugin_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QQmlExtensionPlugin*>(ptr)->baseUrl());
}

void QQmlExtensionPlugin_RegisterTypes(void* ptr, char* uri)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->registerTypes(const_cast<const char*>(uri));
}

void QQmlExtensionPlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlExtensionPlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlExtensionPlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlExtensionPlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlExtensionPlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExtensionPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExtensionPlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlExtensionPlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::customEvent(static_cast<QEvent*>(event));
}

void QQmlExtensionPlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlExtensionPlugin*>(ptr), "deleteLater");
}

void QQmlExtensionPlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::deleteLater();
}

void QQmlExtensionPlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExtensionPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlExtensionPlugin_Event(void* ptr, void* e)
{
	return static_cast<QQmlExtensionPlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlExtensionPlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::event(static_cast<QEvent*>(e));
}

int QQmlExtensionPlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlExtensionPlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlExtensionPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlExtensionPlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlExtensionPlugin*>(ptr)->metaObject());
}

void* QQmlExtensionPlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlExtensionPlugin*>(ptr)->QQmlExtensionPlugin::metaObject());
}

void* QQmlFileSelector_NewQQmlFileSelector(void* engine, void* parent)
{
	return new QQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlFileSelector_QQmlFileSelector_Get(void* engine)
{
	return QQmlFileSelector::get(static_cast<QQmlEngine*>(engine));
}

void QQmlFileSelector_SetExtraSelectors(void* ptr, char* strin)
{
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split("|", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetExtraSelectors2(void* ptr, char* strin)
{
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split("|", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetSelector(void* ptr, void* selector)
{
	static_cast<QQmlFileSelector*>(ptr)->setSelector(static_cast<QFileSelector*>(selector));
}

void QQmlFileSelector_DestroyQQmlFileSelector(void* ptr)
{
	static_cast<QQmlFileSelector*>(ptr)->~QQmlFileSelector();
}

void QQmlFileSelector_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlFileSelector*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlFileSelector_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlFileSelector_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlFileSelector*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlFileSelector_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlFileSelector_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlFileSelector*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlFileSelector_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlFileSelector_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlFileSelector*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlFileSelector_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::customEvent(static_cast<QEvent*>(event));
}

void QQmlFileSelector_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlFileSelector*>(ptr), "deleteLater");
}

void QQmlFileSelector_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::deleteLater();
}

void QQmlFileSelector_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlFileSelector*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlFileSelector_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlFileSelector_Event(void* ptr, void* e)
{
	return static_cast<QQmlFileSelector*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlFileSelector_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::event(static_cast<QEvent*>(e));
}

int QQmlFileSelector_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlFileSelector*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlFileSelector_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlFileSelector_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlFileSelector*>(ptr)->metaObject());
}

void* QQmlFileSelector_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::metaObject());
}

class MyQQmlImageProviderBase: public QQmlImageProviderBase
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	Flags flags() const { return static_cast<QQmlImageProviderBase::Flag>(callbackQQmlImageProviderBase_Flags(const_cast<MyQQmlImageProviderBase*>(this), this->objectNameAbs().toUtf8().data())); };
	ImageType imageType() const { return static_cast<QQmlImageProviderBase::ImageType>(callbackQQmlImageProviderBase_ImageType(const_cast<MyQQmlImageProviderBase*>(this), this->objectNameAbs().toUtf8().data())); };
};

int QQmlImageProviderBase_Flags(void* ptr)
{
	return static_cast<QQmlImageProviderBase*>(ptr)->flags();
}

int QQmlImageProviderBase_ImageType(void* ptr)
{
	return static_cast<QQmlImageProviderBase*>(ptr)->imageType();
}

char* QQmlImageProviderBase_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlImageProviderBase*>(static_cast<QQmlImageProviderBase*>(ptr))) {
		return static_cast<MyQQmlImageProviderBase*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlImageProviderBase_BASE").toUtf8().data();
}

void QQmlImageProviderBase_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlImageProviderBase*>(static_cast<QQmlImageProviderBase*>(ptr))) {
		static_cast<MyQQmlImageProviderBase*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQmlIncubationController: public QQmlIncubationController
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQmlIncubationController() : QQmlIncubationController() {};
	void incubatingObjectCountChanged(int incubatingObjectCount) { callbackQQmlIncubationController_IncubatingObjectCountChanged(this, this->objectNameAbs().toUtf8().data(), incubatingObjectCount); };
};

void* QQmlIncubationController_NewQQmlIncubationController()
{
	return new MyQQmlIncubationController();
}

void* QQmlIncubationController_Engine(void* ptr)
{
	return static_cast<QQmlIncubationController*>(ptr)->engine();
}

void QQmlIncubationController_IncubateFor(void* ptr, int msecs)
{
	static_cast<QQmlIncubationController*>(ptr)->incubateFor(msecs);
}

int QQmlIncubationController_IncubatingObjectCount(void* ptr)
{
	return static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCount();
}

void QQmlIncubationController_IncubatingObjectCountChanged(void* ptr, int incubatingObjectCount)
{
	static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCountChanged(incubatingObjectCount);
}

void QQmlIncubationController_IncubatingObjectCountChangedDefault(void* ptr, int incubatingObjectCount)
{
	static_cast<QQmlIncubationController*>(ptr)->QQmlIncubationController::incubatingObjectCountChanged(incubatingObjectCount);
}

char* QQmlIncubationController_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlIncubationController*>(static_cast<QQmlIncubationController*>(ptr))) {
		return static_cast<MyQQmlIncubationController*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlIncubationController_BASE").toUtf8().data();
}

void QQmlIncubationController_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlIncubationController*>(static_cast<QQmlIncubationController*>(ptr))) {
		static_cast<MyQQmlIncubationController*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQmlIncubator: public QQmlIncubator
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQmlIncubator(IncubationMode mode) : QQmlIncubator(mode) {};
	void setInitialState(QObject * object) { callbackQQmlIncubator_SetInitialState(this, this->objectNameAbs().toUtf8().data(), object); };
	void statusChanged(QQmlIncubator::Status status) { callbackQQmlIncubator_StatusChanged(this, this->objectNameAbs().toUtf8().data(), status); };
};

void* QQmlIncubator_NewQQmlIncubator(int mode)
{
	return new MyQQmlIncubator(static_cast<QQmlIncubator::IncubationMode>(mode));
}

void QQmlIncubator_Clear(void* ptr)
{
	static_cast<QQmlIncubator*>(ptr)->clear();
}

void QQmlIncubator_ForceCompletion(void* ptr)
{
	static_cast<QQmlIncubator*>(ptr)->forceCompletion();
}

int QQmlIncubator_IncubationMode(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->incubationMode();
}

int QQmlIncubator_IsError(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isError();
}

int QQmlIncubator_IsLoading(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isLoading();
}

int QQmlIncubator_IsNull(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isNull();
}

int QQmlIncubator_IsReady(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isReady();
}

void* QQmlIncubator_Object(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->object();
}

void QQmlIncubator_SetInitialState(void* ptr, void* object)
{
	static_cast<QQmlIncubator*>(ptr)->setInitialState(static_cast<QObject*>(object));
}

void QQmlIncubator_SetInitialStateDefault(void* ptr, void* object)
{
	static_cast<QQmlIncubator*>(ptr)->QQmlIncubator::setInitialState(static_cast<QObject*>(object));
}

int QQmlIncubator_Status(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->status();
}

void QQmlIncubator_StatusChanged(void* ptr, int status)
{
	static_cast<QQmlIncubator*>(ptr)->statusChanged(static_cast<QQmlIncubator::Status>(status));
}

void QQmlIncubator_StatusChangedDefault(void* ptr, int status)
{
	static_cast<QQmlIncubator*>(ptr)->QQmlIncubator::statusChanged(static_cast<QQmlIncubator::Status>(status));
}

char* QQmlIncubator_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlIncubator*>(static_cast<QQmlIncubator*>(ptr))) {
		return static_cast<MyQQmlIncubator*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlIncubator_BASE").toUtf8().data();
}

void QQmlIncubator_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlIncubator*>(static_cast<QQmlIncubator*>(ptr))) {
		static_cast<MyQQmlIncubator*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QQmlListReference_NewQQmlListReference()
{
	return new QQmlListReference();
}

void* QQmlListReference_NewQQmlListReference2(void* object, char* property, void* engine)
{
	return new QQmlListReference(static_cast<QObject*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
}

int QQmlListReference_Append(void* ptr, void* object)
{
	return static_cast<QQmlListReference*>(ptr)->append(static_cast<QObject*>(object));
}

void* QQmlListReference_At(void* ptr, int index)
{
	return static_cast<QQmlListReference*>(ptr)->at(index);
}

int QQmlListReference_CanAppend(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canAppend();
}

int QQmlListReference_CanAt(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canAt();
}

int QQmlListReference_CanClear(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canClear();
}

int QQmlListReference_CanCount(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canCount();
}

int QQmlListReference_Clear(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->clear();
}

int QQmlListReference_Count(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->count();
}

int QQmlListReference_IsManipulable(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->isManipulable();
}

int QQmlListReference_IsReadable(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->isReadable();
}

int QQmlListReference_IsValid(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->isValid();
}

void* QQmlListReference_ListElementType(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlListReference*>(ptr)->listElementType());
}

void* QQmlListReference_Object(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->object();
}

class MyQQmlNetworkAccessManagerFactory: public QQmlNetworkAccessManagerFactory
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QNetworkAccessManager * create(QObject * parent) { return static_cast<QNetworkAccessManager*>(callbackQQmlNetworkAccessManagerFactory_Create(this, this->objectNameAbs().toUtf8().data(), parent)); };
};

void* QQmlNetworkAccessManagerFactory_Create(void* ptr, void* parent)
{
	return static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->create(static_cast<QObject*>(parent));
}

void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(void* ptr)
{
	static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->~QQmlNetworkAccessManagerFactory();
}

char* QQmlNetworkAccessManagerFactory_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlNetworkAccessManagerFactory*>(static_cast<QQmlNetworkAccessManagerFactory*>(ptr))) {
		return static_cast<MyQQmlNetworkAccessManagerFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlNetworkAccessManagerFactory_BASE").toUtf8().data();
}

void QQmlNetworkAccessManagerFactory_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlNetworkAccessManagerFactory*>(static_cast<QQmlNetworkAccessManagerFactory*>(ptr))) {
		static_cast<MyQQmlNetworkAccessManagerFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQmlParserStatus: public QQmlParserStatus
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void classBegin() { callbackQQmlParserStatus_ClassBegin(this, this->objectNameAbs().toUtf8().data()); };
	void componentComplete() { callbackQQmlParserStatus_ComponentComplete(this, this->objectNameAbs().toUtf8().data()); };
};

void QQmlParserStatus_ClassBegin(void* ptr)
{
	static_cast<QQmlParserStatus*>(ptr)->classBegin();
}

void QQmlParserStatus_ComponentComplete(void* ptr)
{
	static_cast<QQmlParserStatus*>(ptr)->componentComplete();
}

char* QQmlParserStatus_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlParserStatus*>(static_cast<QQmlParserStatus*>(ptr))) {
		return static_cast<MyQQmlParserStatus*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlParserStatus_BASE").toUtf8().data();
}

void QQmlParserStatus_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlParserStatus*>(static_cast<QQmlParserStatus*>(ptr))) {
		static_cast<MyQQmlParserStatus*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QQmlProperty_NewQQmlProperty()
{
	return new QQmlProperty();
}

void* QQmlProperty_NewQQmlProperty2(void* obj)
{
	return new QQmlProperty(static_cast<QObject*>(obj));
}

void* QQmlProperty_NewQQmlProperty3(void* obj, void* ctxt)
{
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlContext*>(ctxt));
}

void* QQmlProperty_NewQQmlProperty4(void* obj, void* engine)
{
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlEngine*>(engine));
}

void* QQmlProperty_NewQQmlProperty5(void* obj, char* name)
{
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name));
}

void* QQmlProperty_NewQQmlProperty6(void* obj, char* name, void* ctxt)
{
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlContext*>(ctxt));
}

void* QQmlProperty_NewQQmlProperty7(void* obj, char* name, void* engine)
{
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlEngine*>(engine));
}

void* QQmlProperty_NewQQmlProperty8(void* other)
{
	return new QQmlProperty(*static_cast<QQmlProperty*>(other));
}

int QQmlProperty_ConnectNotifySignal(void* ptr, void* dest, char* slot)
{
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), const_cast<const char*>(slot));
}

int QQmlProperty_ConnectNotifySignal2(void* ptr, void* dest, int method)
{
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), method);
}

int QQmlProperty_HasNotifySignal(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->hasNotifySignal();
}

int QQmlProperty_Index(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->index();
}

int QQmlProperty_IsDesignable(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isDesignable();
}

int QQmlProperty_IsProperty(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isProperty();
}

int QQmlProperty_IsResettable(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isResettable();
}

int QQmlProperty_IsSignalProperty(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isSignalProperty();
}

int QQmlProperty_IsValid(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isValid();
}

int QQmlProperty_IsWritable(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isWritable();
}

void* QQmlProperty_Method(void* ptr)
{
	return new QMetaMethod(static_cast<QQmlProperty*>(ptr)->method());
}

char* QQmlProperty_Name(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->name().toUtf8().data();
}

int QQmlProperty_NeedsNotifySignal(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->needsNotifySignal();
}

void* QQmlProperty_Object(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->object();
}

int QQmlProperty_PropertyType(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->propertyType();
}

int QQmlProperty_PropertyTypeCategory(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->propertyTypeCategory();
}

char* QQmlProperty_PropertyTypeName(void* ptr)
{
	return QString(static_cast<QQmlProperty*>(ptr)->propertyTypeName()).toUtf8().data();
}

void* QQmlProperty_QQmlProperty_Read2(void* object, char* name)
{
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name)));
}

void* QQmlProperty_QQmlProperty_Read3(void* object, char* name, void* ctxt)
{
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlContext*>(ctxt)));
}

void* QQmlProperty_QQmlProperty_Read4(void* object, char* name, void* engine)
{
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlEngine*>(engine)));
}

void* QQmlProperty_Read(void* ptr)
{
	return new QVariant(static_cast<QQmlProperty*>(ptr)->read());
}

int QQmlProperty_Reset(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->reset();
}

int QQmlProperty_Type(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->type();
}

int QQmlProperty_QQmlProperty_Write2(void* object, char* name, void* value)
{
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value));
}

int QQmlProperty_QQmlProperty_Write3(void* object, char* name, void* value, void* ctxt)
{
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value), static_cast<QQmlContext*>(ctxt));
}

int QQmlProperty_QQmlProperty_Write4(void* object, char* name, void* value, void* engine)
{
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value), static_cast<QQmlEngine*>(engine));
}

int QQmlProperty_Write(void* ptr, void* value)
{
	return static_cast<QQmlProperty*>(ptr)->write(*static_cast<QVariant*>(value));
}

class MyQQmlPropertyMap: public QQmlPropertyMap
{
public:
	MyQQmlPropertyMap(QObject *parent) : QQmlPropertyMap(parent) {};
	QVariant updateValue(const QString & key, const QVariant & input) { return *static_cast<QVariant*>(callbackQQmlPropertyMap_UpdateValue(this, this->objectName().toUtf8().data(), key.toUtf8().data(), new QVariant(input))); };
	void Signal_ValueChanged(const QString & key, const QVariant & value) { callbackQQmlPropertyMap_ValueChanged(this, this->objectName().toUtf8().data(), key.toUtf8().data(), new QVariant(value)); };
	void timerEvent(QTimerEvent * event) { callbackQQmlPropertyMap_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlPropertyMap_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlPropertyMap_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQmlPropertyMap_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQmlPropertyMap_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlPropertyMap_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQmlPropertyMap_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlPropertyMap_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlPropertyMap_MetaObject(const_cast<MyQQmlPropertyMap*>(this), this->objectName().toUtf8().data())); };
};

void* QQmlPropertyMap_NewQQmlPropertyMap(void* parent)
{
	return new MyQQmlPropertyMap(static_cast<QObject*>(parent));
}

void QQmlPropertyMap_Clear(void* ptr, char* key)
{
	static_cast<QQmlPropertyMap*>(ptr)->clear(QString(key));
}

int QQmlPropertyMap_Contains(void* ptr, char* key)
{
	return static_cast<QQmlPropertyMap*>(ptr)->contains(QString(key));
}

int QQmlPropertyMap_Count(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->count();
}

void QQmlPropertyMap_Insert(void* ptr, char* key, void* value)
{
	static_cast<QQmlPropertyMap*>(ptr)->insert(QString(key), *static_cast<QVariant*>(value));
}

int QQmlPropertyMap_IsEmpty(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->isEmpty();
}

char* QQmlPropertyMap_Keys(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->keys().join("|").toUtf8().data();
}

int QQmlPropertyMap_Size(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->size();
}

void* QQmlPropertyMap_UpdateValue(void* ptr, char* key, void* input)
{
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->updateValue(QString(key), *static_cast<QVariant*>(input)));
}

void* QQmlPropertyMap_UpdateValueDefault(void* ptr, char* key, void* input)
{
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::updateValue(QString(key), *static_cast<QVariant*>(input)));
}

void* QQmlPropertyMap_Value(void* ptr, char* key)
{
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->value(QString(key)));
}

void QQmlPropertyMap_ConnectValueChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));
}

void QQmlPropertyMap_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));
}

void QQmlPropertyMap_ValueChanged(void* ptr, char* key, void* value)
{
	static_cast<QQmlPropertyMap*>(ptr)->valueChanged(QString(key), *static_cast<QVariant*>(value));
}

void QQmlPropertyMap_DestroyQQmlPropertyMap(void* ptr)
{
	static_cast<QQmlPropertyMap*>(ptr)->~QQmlPropertyMap();
}

void QQmlPropertyMap_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlPropertyMap*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlPropertyMap_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlPropertyMap_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlPropertyMap*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlPropertyMap_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlPropertyMap_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlPropertyMap*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlPropertyMap_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlPropertyMap_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlPropertyMap*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlPropertyMap_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::customEvent(static_cast<QEvent*>(event));
}

void QQmlPropertyMap_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlPropertyMap*>(ptr), "deleteLater");
}

void QQmlPropertyMap_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::deleteLater();
}

void QQmlPropertyMap_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlPropertyMap*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlPropertyMap_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQmlPropertyMap_Event(void* ptr, void* e)
{
	return static_cast<QQmlPropertyMap*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQmlPropertyMap_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::event(static_cast<QEvent*>(e));
}

int QQmlPropertyMap_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlPropertyMap*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQmlPropertyMap_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlPropertyMap_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlPropertyMap*>(ptr)->metaObject());
}

void* QQmlPropertyMap_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::metaObject());
}

class MyQQmlPropertyValueSource: public QQmlPropertyValueSource
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQmlPropertyValueSource() : QQmlPropertyValueSource() {};
	void setTarget(const QQmlProperty & property) { callbackQQmlPropertyValueSource_SetTarget(this, this->objectNameAbs().toUtf8().data(), new QQmlProperty(property)); };
};

void* QQmlPropertyValueSource_NewQQmlPropertyValueSource()
{
	return new MyQQmlPropertyValueSource();
}

void QQmlPropertyValueSource_SetTarget(void* ptr, void* property)
{
	static_cast<QQmlPropertyValueSource*>(ptr)->setTarget(*static_cast<QQmlProperty*>(property));
}

void QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(void* ptr)
{
	static_cast<QQmlPropertyValueSource*>(ptr)->~QQmlPropertyValueSource();
}

char* QQmlPropertyValueSource_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQmlPropertyValueSource*>(static_cast<QQmlPropertyValueSource*>(ptr))) {
		return static_cast<MyQQmlPropertyValueSource*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQmlPropertyValueSource_BASE").toUtf8().data();
}

void QQmlPropertyValueSource_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQmlPropertyValueSource*>(static_cast<QQmlPropertyValueSource*>(ptr))) {
		static_cast<MyQQmlPropertyValueSource*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QQmlScriptString_NewQQmlScriptString()
{
	return new QQmlScriptString();
}

void* QQmlScriptString_NewQQmlScriptString2(void* other)
{
	return new QQmlScriptString(*static_cast<QQmlScriptString*>(other));
}

int QQmlScriptString_BooleanLiteral(void* ptr, int ok)
{
	return static_cast<QQmlScriptString*>(ptr)->booleanLiteral(NULL);
}

int QQmlScriptString_IsEmpty(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->isEmpty();
}

int QQmlScriptString_IsNullLiteral(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->isNullLiteral();
}

int QQmlScriptString_IsUndefinedLiteral(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->isUndefinedLiteral();
}

double QQmlScriptString_NumberLiteral(void* ptr, int ok)
{
	return static_cast<double>(static_cast<QQmlScriptString*>(ptr)->numberLiteral(NULL));
}

char* QQmlScriptString_StringLiteral(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->stringLiteral().toUtf8().data();
}

