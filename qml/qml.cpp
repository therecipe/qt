#include "qml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QDate>
#include <QDateTime>
#include <QFile>
#include <QFileSelector>
#include <QJSEngine>
#include <QJSValue>
#include <QLatin1String>
#include <QMetaObject>
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
#include <QUrl>
#include <QVariant>

class MyQJSEngine: public QJSEngine {
public:
	MyQJSEngine() : QJSEngine() {};
	MyQJSEngine(QObject *parent) : QJSEngine(parent) {};
protected:
};

void* QJSEngine_NewQJSEngine(){
	return new MyQJSEngine();
}

void* QJSEngine_NewQJSEngine2(void* parent){
	return new MyQJSEngine(static_cast<QObject*>(parent));
}

void QJSEngine_CollectGarbage(void* ptr){
	static_cast<QJSEngine*>(ptr)->collectGarbage();
}

void* QJSEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->evaluate(QString(program), QString(fileName), lineNumber));
}

void* QJSEngine_GlobalObject(void* ptr){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->globalObject());
}

void QJSEngine_InstallTranslatorFunctions(void* ptr, void* object){
	static_cast<QJSEngine*>(ptr)->installTranslatorFunctions(*static_cast<QJSValue*>(object));
}

void* QJSEngine_NewObject(void* ptr){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newObject());
}

void* QJSEngine_NewQObject(void* ptr, void* object){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newQObject(static_cast<QObject*>(object)));
}

void QJSEngine_DestroyQJSEngine(void* ptr){
	static_cast<QJSEngine*>(ptr)->~QJSEngine();
}

void* QJSValue_NewQJSValue3(void* other){
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue(int value){
	return new QJSValue(static_cast<QJSValue::SpecialValue>(value));
}

void* QJSValue_NewQJSValue4(int value){
	return new QJSValue(value != 0);
}

void* QJSValue_NewQJSValue2(void* other){
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue9(void* value){
	return new QJSValue(*static_cast<QLatin1String*>(value));
}

void* QJSValue_NewQJSValue8(char* value){
	return new QJSValue(QString(value));
}

void* QJSValue_NewQJSValue10(char* value){
	return new QJSValue(const_cast<const char*>(value));
}

void* QJSValue_NewQJSValue5(int value){
	return new QJSValue(value);
}

int QJSValue_DeleteProperty(void* ptr, char* name){
	return static_cast<QJSValue*>(ptr)->deleteProperty(QString(name));
}

int QJSValue_Equals(void* ptr, void* other){
	return static_cast<QJSValue*>(ptr)->equals(*static_cast<QJSValue*>(other));
}

int QJSValue_HasOwnProperty(void* ptr, char* name){
	return static_cast<QJSValue*>(ptr)->hasOwnProperty(QString(name));
}

int QJSValue_HasProperty(void* ptr, char* name){
	return static_cast<QJSValue*>(ptr)->hasProperty(QString(name));
}

int QJSValue_IsArray(void* ptr){
	return static_cast<QJSValue*>(ptr)->isArray();
}

int QJSValue_IsBool(void* ptr){
	return static_cast<QJSValue*>(ptr)->isBool();
}

int QJSValue_IsCallable(void* ptr){
	return static_cast<QJSValue*>(ptr)->isCallable();
}

int QJSValue_IsDate(void* ptr){
	return static_cast<QJSValue*>(ptr)->isDate();
}

int QJSValue_IsError(void* ptr){
	return static_cast<QJSValue*>(ptr)->isError();
}

int QJSValue_IsNull(void* ptr){
	return static_cast<QJSValue*>(ptr)->isNull();
}

int QJSValue_IsNumber(void* ptr){
	return static_cast<QJSValue*>(ptr)->isNumber();
}

int QJSValue_IsObject(void* ptr){
	return static_cast<QJSValue*>(ptr)->isObject();
}

int QJSValue_IsQObject(void* ptr){
	return static_cast<QJSValue*>(ptr)->isQObject();
}

int QJSValue_IsRegExp(void* ptr){
	return static_cast<QJSValue*>(ptr)->isRegExp();
}

int QJSValue_IsString(void* ptr){
	return static_cast<QJSValue*>(ptr)->isString();
}

int QJSValue_IsUndefined(void* ptr){
	return static_cast<QJSValue*>(ptr)->isUndefined();
}

int QJSValue_IsVariant(void* ptr){
	return static_cast<QJSValue*>(ptr)->isVariant();
}

void* QJSValue_Property(void* ptr, char* name){
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(QString(name)));
}

void* QJSValue_Prototype(void* ptr){
	return new QJSValue(static_cast<QJSValue*>(ptr)->prototype());
}

void QJSValue_SetProperty(void* ptr, char* name, void* value){
	static_cast<QJSValue*>(ptr)->setProperty(QString(name), *static_cast<QJSValue*>(value));
}

void QJSValue_SetPrototype(void* ptr, void* prototype){
	static_cast<QJSValue*>(ptr)->setPrototype(*static_cast<QJSValue*>(prototype));
}

int QJSValue_StrictlyEquals(void* ptr, void* other){
	return static_cast<QJSValue*>(ptr)->strictlyEquals(*static_cast<QJSValue*>(other));
}

int QJSValue_ToBool(void* ptr){
	return static_cast<QJSValue*>(ptr)->toBool();
}

void* QJSValue_ToDateTime(void* ptr){
	return new QDateTime(static_cast<QJSValue*>(ptr)->toDateTime());
}

void* QJSValue_ToQObject(void* ptr){
	return static_cast<QJSValue*>(ptr)->toQObject();
}

char* QJSValue_ToString(void* ptr){
	return static_cast<QJSValue*>(ptr)->toString().toUtf8().data();
}

void* QJSValue_ToVariant(void* ptr){
	return new QVariant(static_cast<QJSValue*>(ptr)->toVariant());
}

void QJSValue_DestroyQJSValue(void* ptr){
	static_cast<QJSValue*>(ptr)->~QJSValue();
}

class MyQQmlAbstractUrlInterceptor: public QQmlAbstractUrlInterceptor {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void* QQmlAbstractUrlInterceptor_Intercept(void* ptr, void* url, int ty){
	return new QUrl(static_cast<QQmlAbstractUrlInterceptor*>(ptr)->intercept(*static_cast<QUrl*>(url), static_cast<QQmlAbstractUrlInterceptor::DataType>(ty)));
}

void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(void* ptr){
	static_cast<QQmlAbstractUrlInterceptor*>(ptr)->~QQmlAbstractUrlInterceptor();
}

char* QQmlAbstractUrlInterceptor_ObjectNameAbs(void* ptr){
	return static_cast<MyQQmlAbstractUrlInterceptor*>(ptr)->objectNameAbs().toUtf8().data();
}

void QQmlAbstractUrlInterceptor_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQQmlAbstractUrlInterceptor*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQQmlApplicationEngine: public QQmlApplicationEngine {
public:
	void Signal_ObjectCreated(QObject * object, const QUrl & url) { callbackQQmlApplicationEngineObjectCreated(this->objectName().toUtf8().data(), object, new QUrl(url)); };
protected:
};

void* QQmlApplicationEngine_NewQQmlApplicationEngine(void* parent){
	return new QQmlApplicationEngine(static_cast<QObject*>(parent));
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine3(char* filePath, void* parent){
	return new QQmlApplicationEngine(QString(filePath), static_cast<QObject*>(parent));
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine2(void* url, void* parent){
	return new QQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QObject*>(parent));
}

void QQmlApplicationEngine_Load2(void* ptr, char* filePath){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QString, QString(filePath)));
}

void QQmlApplicationEngine_Load(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_LoadData(void* ptr, void* data, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "loadData", Q_ARG(QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_ConnectObjectCreated(void* ptr){
	QObject::connect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));;
}

void QQmlApplicationEngine_DisconnectObjectCreated(void* ptr){
	QObject::disconnect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));;
}

void QQmlApplicationEngine_DestroyQQmlApplicationEngine(void* ptr){
	static_cast<QQmlApplicationEngine*>(ptr)->~QQmlApplicationEngine();
}

class MyQQmlComponent: public QQmlComponent {
public:
	MyQQmlComponent(QQmlEngine *engine, QObject *parent) : QQmlComponent(engine, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QString &fileName, CompilationMode mode, QObject *parent) : QQmlComponent(engine, fileName, mode, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QString &fileName, QObject *parent) : QQmlComponent(engine, fileName, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QUrl &url, CompilationMode mode, QObject *parent) : QQmlComponent(engine, url, mode, parent) {};
	MyQQmlComponent(QQmlEngine *engine, const QUrl &url, QObject *parent) : QQmlComponent(engine, url, parent) {};
	void completeCreate() { if (!callbackQQmlComponentCompleteCreate(this->objectName().toUtf8().data())) { QQmlComponent::completeCreate(); }; };
	void Signal_ProgressChanged(qreal progress) { callbackQQmlComponentProgressChanged(this->objectName().toUtf8().data(), static_cast<double>(progress)); };
	void Signal_StatusChanged(QQmlComponent::Status status) { callbackQQmlComponentStatusChanged(this->objectName().toUtf8().data(), status); };
protected:
};

double QQmlComponent_Progress(void* ptr){
	return static_cast<double>(static_cast<QQmlComponent*>(ptr)->progress());
}

int QQmlComponent_Status(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->status();
}

void* QQmlComponent_Url(void* ptr){
	return new QUrl(static_cast<QQmlComponent*>(ptr)->url());
}

void* QQmlComponent_NewQQmlComponent(void* engine, void* parent){
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent4(void* engine, char* fileName, int mode, void* parent){
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent3(void* engine, char* fileName, void* parent){
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent6(void* engine, void* url, int mode, void* parent){
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent5(void* engine, void* url, void* parent){
	return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QObject*>(parent));
}

void* QQmlComponent_BeginCreate(void* ptr, void* publicContext){
	return static_cast<QQmlComponent*>(ptr)->beginCreate(static_cast<QQmlContext*>(publicContext));
}

void QQmlComponent_CompleteCreate(void* ptr){
	static_cast<QQmlComponent*>(ptr)->completeCreate();
}

void* QQmlComponent_Create(void* ptr, void* context){
	return static_cast<QQmlComponent*>(ptr)->create(static_cast<QQmlContext*>(context));
}

void QQmlComponent_Create2(void* ptr, void* incubator, void* context, void* forContext){
	static_cast<QQmlComponent*>(ptr)->create(*static_cast<QQmlIncubator*>(incubator), static_cast<QQmlContext*>(context), static_cast<QQmlContext*>(forContext));
}

void* QQmlComponent_CreationContext(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->creationContext();
}

int QQmlComponent_IsError(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isError();
}

int QQmlComponent_IsLoading(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isLoading();
}

int QQmlComponent_IsNull(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isNull();
}

int QQmlComponent_IsReady(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isReady();
}

void QQmlComponent_LoadUrl(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_LoadUrl2(void* ptr, void* url, int mode){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, *static_cast<QUrl*>(url)), Q_ARG(QQmlComponent::CompilationMode, static_cast<QQmlComponent::CompilationMode>(mode)));
}

void QQmlComponent_ConnectProgressChanged(void* ptr){
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(qreal)>(&QQmlComponent::progressChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(qreal)>(&MyQQmlComponent::Signal_ProgressChanged));;
}

void QQmlComponent_DisconnectProgressChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(qreal)>(&QQmlComponent::progressChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(qreal)>(&MyQQmlComponent::Signal_ProgressChanged));;
}

void QQmlComponent_SetData(void* ptr, void* data, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "setData", Q_ARG(QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));;
}

void QQmlComponent_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));;
}

void QQmlComponent_DestroyQQmlComponent(void* ptr){
	static_cast<QQmlComponent*>(ptr)->~QQmlComponent();
}

class MyQQmlContext: public QQmlContext {
public:
	MyQQmlContext(QQmlContext *parentContext, QObject *parent) : QQmlContext(parentContext, parent) {};
	MyQQmlContext(QQmlEngine *engine, QObject *parent) : QQmlContext(engine, parent) {};
protected:
};

void* QQmlContext_NewQQmlContext2(void* parentContext, void* parent){
	return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QObject*>(parent));
}

void* QQmlContext_NewQQmlContext(void* engine, void* parent){
	return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlContext_BaseUrl(void* ptr){
	return new QUrl(static_cast<QQmlContext*>(ptr)->baseUrl());
}

void* QQmlContext_ContextObject(void* ptr){
	return static_cast<QQmlContext*>(ptr)->contextObject();
}

void* QQmlContext_ContextProperty(void* ptr, char* name){
	return new QVariant(static_cast<QQmlContext*>(ptr)->contextProperty(QString(name)));
}

void* QQmlContext_Engine(void* ptr){
	return static_cast<QQmlContext*>(ptr)->engine();
}

int QQmlContext_IsValid(void* ptr){
	return static_cast<QQmlContext*>(ptr)->isValid();
}

char* QQmlContext_NameForObject(void* ptr, void* object){
	return static_cast<QQmlContext*>(ptr)->nameForObject(static_cast<QObject*>(object)).toUtf8().data();
}

void* QQmlContext_ParentContext(void* ptr){
	return static_cast<QQmlContext*>(ptr)->parentContext();
}

void* QQmlContext_ResolvedUrl(void* ptr, void* src){
	return new QUrl(static_cast<QQmlContext*>(ptr)->resolvedUrl(*static_cast<QUrl*>(src)));
}

void QQmlContext_SetBaseUrl(void* ptr, void* baseUrl){
	static_cast<QQmlContext*>(ptr)->setBaseUrl(*static_cast<QUrl*>(baseUrl));
}

void QQmlContext_SetContextObject(void* ptr, void* object){
	static_cast<QQmlContext*>(ptr)->setContextObject(static_cast<QObject*>(object));
}

void QQmlContext_SetContextProperty(void* ptr, char* name, void* value){
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), static_cast<QObject*>(value));
}

void QQmlContext_SetContextProperty2(void* ptr, char* name, void* value){
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), *static_cast<QVariant*>(value));
}

void QQmlContext_DestroyQQmlContext(void* ptr){
	static_cast<QQmlContext*>(ptr)->~QQmlContext();
}

class MyQQmlEngine: public QQmlEngine {
public:
	MyQQmlEngine(QObject *parent) : QQmlEngine(parent) {};
	void Signal_Quit() { callbackQQmlEngineQuit(this->objectName().toUtf8().data()); };
protected:
};

char* QQmlEngine_OfflineStoragePath(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->offlineStoragePath().toUtf8().data();
}

void QQmlEngine_SetOfflineStoragePath(void* ptr, char* dir){
	static_cast<QQmlEngine*>(ptr)->setOfflineStoragePath(QString(dir));
}

void* QQmlEngine_NewQQmlEngine(void* parent){
	return new MyQQmlEngine(static_cast<QObject*>(parent));
}

void QQmlEngine_AddImageProvider(void* ptr, char* providerId, void* provider){
	static_cast<QQmlEngine*>(ptr)->addImageProvider(QString(providerId), static_cast<QQmlImageProviderBase*>(provider));
}

void QQmlEngine_AddImportPath(void* ptr, char* path){
	static_cast<QQmlEngine*>(ptr)->addImportPath(QString(path));
}

void QQmlEngine_AddPluginPath(void* ptr, char* path){
	static_cast<QQmlEngine*>(ptr)->addPluginPath(QString(path));
}

void* QQmlEngine_BaseUrl(void* ptr){
	return new QUrl(static_cast<QQmlEngine*>(ptr)->baseUrl());
}

void QQmlEngine_ClearComponentCache(void* ptr){
	static_cast<QQmlEngine*>(ptr)->clearComponentCache();
}

void* QQmlEngine_QQmlEngine_ContextForObject(void* object){
	return QQmlEngine::contextForObject(static_cast<QObject*>(object));
}

void* QQmlEngine_ImageProvider(void* ptr, char* providerId){
	return static_cast<QQmlEngine*>(ptr)->imageProvider(QString(providerId));
}

char* QQmlEngine_ImportPathList(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->importPathList().join(",,,").toUtf8().data();
}

void* QQmlEngine_IncubationController(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->incubationController();
}

void* QQmlEngine_NetworkAccessManager(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->networkAccessManager();
}

void* QQmlEngine_NetworkAccessManagerFactory(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->networkAccessManagerFactory();
}

int QQmlEngine_QQmlEngine_ObjectOwnership(void* object){
	return QQmlEngine::objectOwnership(static_cast<QObject*>(object));
}

int QQmlEngine_OutputWarningsToStandardError(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->outputWarningsToStandardError();
}

char* QQmlEngine_PluginPathList(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->pluginPathList().join(",,,").toUtf8().data();
}

void QQmlEngine_ConnectQuit(void* ptr){
	QObject::connect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));;
}

void QQmlEngine_DisconnectQuit(void* ptr){
	QObject::disconnect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));;
}

void QQmlEngine_RemoveImageProvider(void* ptr, char* providerId){
	static_cast<QQmlEngine*>(ptr)->removeImageProvider(QString(providerId));
}

void* QQmlEngine_RootContext(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->rootContext();
}

void QQmlEngine_SetBaseUrl(void* ptr, void* url){
	static_cast<QQmlEngine*>(ptr)->setBaseUrl(*static_cast<QUrl*>(url));
}

void QQmlEngine_QQmlEngine_SetContextForObject(void* object, void* context){
	QQmlEngine::setContextForObject(static_cast<QObject*>(object), static_cast<QQmlContext*>(context));
}

void QQmlEngine_SetImportPathList(void* ptr, char* paths){
	static_cast<QQmlEngine*>(ptr)->setImportPathList(QString(paths).split(",,,", QString::SkipEmptyParts));
}

void QQmlEngine_SetIncubationController(void* ptr, void* controller){
	static_cast<QQmlEngine*>(ptr)->setIncubationController(static_cast<QQmlIncubationController*>(controller));
}

void QQmlEngine_SetNetworkAccessManagerFactory(void* ptr, void* factory){
	static_cast<QQmlEngine*>(ptr)->setNetworkAccessManagerFactory(static_cast<QQmlNetworkAccessManagerFactory*>(factory));
}

void QQmlEngine_QQmlEngine_SetObjectOwnership(void* object, int ownership){
	QQmlEngine::setObjectOwnership(static_cast<QObject*>(object), static_cast<QQmlEngine::ObjectOwnership>(ownership));
}

void QQmlEngine_SetOutputWarningsToStandardError(void* ptr, int enabled){
	static_cast<QQmlEngine*>(ptr)->setOutputWarningsToStandardError(enabled != 0);
}

void QQmlEngine_SetPluginPathList(void* ptr, char* paths){
	static_cast<QQmlEngine*>(ptr)->setPluginPathList(QString(paths).split(",,,", QString::SkipEmptyParts));
}

void QQmlEngine_TrimComponentCache(void* ptr){
	static_cast<QQmlEngine*>(ptr)->trimComponentCache();
}

void QQmlEngine_DestroyQQmlEngine(void* ptr){
	static_cast<QQmlEngine*>(ptr)->~QQmlEngine();
}

void* QQmlError_NewQQmlError(){
	return new QQmlError();
}

void* QQmlError_NewQQmlError2(void* other){
	return new QQmlError(*static_cast<QQmlError*>(other));
}

int QQmlError_Column(void* ptr){
	return static_cast<QQmlError*>(ptr)->column();
}

char* QQmlError_Description(void* ptr){
	return static_cast<QQmlError*>(ptr)->description().toUtf8().data();
}

int QQmlError_IsValid(void* ptr){
	return static_cast<QQmlError*>(ptr)->isValid();
}

int QQmlError_Line(void* ptr){
	return static_cast<QQmlError*>(ptr)->line();
}

void* QQmlError_Object(void* ptr){
	return static_cast<QQmlError*>(ptr)->object();
}

void QQmlError_SetColumn(void* ptr, int column){
	static_cast<QQmlError*>(ptr)->setColumn(column);
}

void QQmlError_SetDescription(void* ptr, char* description){
	static_cast<QQmlError*>(ptr)->setDescription(QString(description));
}

void QQmlError_SetLine(void* ptr, int line){
	static_cast<QQmlError*>(ptr)->setLine(line);
}

void QQmlError_SetObject(void* ptr, void* object){
	static_cast<QQmlError*>(ptr)->setObject(static_cast<QObject*>(object));
}

void QQmlError_SetUrl(void* ptr, void* url){
	static_cast<QQmlError*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

char* QQmlError_ToString(void* ptr){
	return static_cast<QQmlError*>(ptr)->toString().toUtf8().data();
}

void* QQmlError_Url(void* ptr){
	return new QUrl(static_cast<QQmlError*>(ptr)->url());
}

class MyQQmlExpression: public QQmlExpression {
public:
	MyQQmlExpression() : QQmlExpression() {};
	MyQQmlExpression(QQmlContext *ctxt, QObject *scope, const QString &expression, QObject *parent) : QQmlExpression(ctxt, scope, expression, parent) {};
	MyQQmlExpression(const QQmlScriptString &script, QQmlContext *ctxt, QObject *scope, QObject *parent) : QQmlExpression(script, ctxt, scope, parent) {};
	void Signal_ValueChanged() { callbackQQmlExpressionValueChanged(this->objectName().toUtf8().data()); };
protected:
};

void* QQmlExpression_NewQQmlExpression(){
	return new MyQQmlExpression();
}

void* QQmlExpression_NewQQmlExpression2(void* ctxt, void* scope, char* expression, void* parent){
	return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), QString(expression), static_cast<QObject*>(parent));
}

void* QQmlExpression_NewQQmlExpression3(void* script, void* ctxt, void* scope, void* parent){
	return new MyQQmlExpression(*static_cast<QQmlScriptString*>(script), static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), static_cast<QObject*>(parent));
}

void QQmlExpression_ClearError(void* ptr){
	static_cast<QQmlExpression*>(ptr)->clearError();
}

int QQmlExpression_ColumnNumber(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->columnNumber();
}

void* QQmlExpression_Context(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->context();
}

void* QQmlExpression_Engine(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->engine();
}

void* QQmlExpression_Evaluate(void* ptr, int valueIsUndefined){
	return new QVariant(static_cast<QQmlExpression*>(ptr)->evaluate(NULL));
}

char* QQmlExpression_Expression(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->expression().toUtf8().data();
}

int QQmlExpression_HasError(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->hasError();
}

int QQmlExpression_LineNumber(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->lineNumber();
}

int QQmlExpression_NotifyOnValueChanged(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->notifyOnValueChanged();
}

void* QQmlExpression_ScopeObject(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->scopeObject();
}

void QQmlExpression_SetExpression(void* ptr, char* expression){
	static_cast<QQmlExpression*>(ptr)->setExpression(QString(expression));
}

void QQmlExpression_SetNotifyOnValueChanged(void* ptr, int notifyOnChange){
	static_cast<QQmlExpression*>(ptr)->setNotifyOnValueChanged(notifyOnChange != 0);
}

void QQmlExpression_SetSourceLocation(void* ptr, char* url, int line, int column){
	static_cast<QQmlExpression*>(ptr)->setSourceLocation(QString(url), line, column);
}

char* QQmlExpression_SourceFile(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->sourceFile().toUtf8().data();
}

void QQmlExpression_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));;
}

void QQmlExpression_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));;
}

void QQmlExpression_DestroyQQmlExpression(void* ptr){
	static_cast<QQmlExpression*>(ptr)->~QQmlExpression();
}

class MyQQmlExtensionPlugin: public QQmlExtensionPlugin {
public:
protected:
};

void* QQmlExtensionPlugin_BaseUrl(void* ptr){
	return new QUrl(static_cast<QQmlExtensionPlugin*>(ptr)->baseUrl());
}

void QQmlExtensionPlugin_RegisterTypes(void* ptr, char* uri){
	static_cast<QQmlExtensionPlugin*>(ptr)->registerTypes(const_cast<const char*>(uri));
}

void* QQmlFileSelector_NewQQmlFileSelector(void* engine, void* parent){
	return new QQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlFileSelector_QQmlFileSelector_Get(void* engine){
	return QQmlFileSelector::get(static_cast<QQmlEngine*>(engine));
}

void QQmlFileSelector_SetExtraSelectors(void* ptr, char* strin){
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split(",,,", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetExtraSelectors2(void* ptr, char* strin){
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split(",,,", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetSelector(void* ptr, void* selector){
	static_cast<QQmlFileSelector*>(ptr)->setSelector(static_cast<QFileSelector*>(selector));
}

void QQmlFileSelector_DestroyQQmlFileSelector(void* ptr){
	static_cast<QQmlFileSelector*>(ptr)->~QQmlFileSelector();
}

int QQmlImageProviderBase_Flags(void* ptr){
	return static_cast<QQmlImageProviderBase*>(ptr)->flags();
}

int QQmlImageProviderBase_ImageType(void* ptr){
	return static_cast<QQmlImageProviderBase*>(ptr)->imageType();
}

class MyQQmlIncubationController: public QQmlIncubationController {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQmlIncubationController() : QQmlIncubationController() {};
protected:
	void incubatingObjectCountChanged(int incubatingObjectCount) { if (!callbackQQmlIncubationControllerIncubatingObjectCountChanged(this->objectNameAbs().toUtf8().data(), incubatingObjectCount)) { QQmlIncubationController::incubatingObjectCountChanged(incubatingObjectCount); }; };
};

void* QQmlIncubationController_NewQQmlIncubationController(){
	return new MyQQmlIncubationController();
}

void* QQmlIncubationController_Engine(void* ptr){
	return static_cast<QQmlIncubationController*>(ptr)->engine();
}

void QQmlIncubationController_IncubateFor(void* ptr, int msecs){
	static_cast<QQmlIncubationController*>(ptr)->incubateFor(msecs);
}

int QQmlIncubationController_IncubatingObjectCount(void* ptr){
	return static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCount();
}

char* QQmlIncubationController_ObjectNameAbs(void* ptr){
	return static_cast<MyQQmlIncubationController*>(ptr)->objectNameAbs().toUtf8().data();
}

void QQmlIncubationController_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQQmlIncubationController*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQQmlIncubator: public QQmlIncubator {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQmlIncubator(IncubationMode mode) : QQmlIncubator(mode) {};
protected:
	void setInitialState(QObject * object) { if (!callbackQQmlIncubatorSetInitialState(this->objectNameAbs().toUtf8().data(), object)) { QQmlIncubator::setInitialState(object); }; };
	void statusChanged(QQmlIncubator::Status status) { if (!callbackQQmlIncubatorStatusChanged(this->objectNameAbs().toUtf8().data(), status)) { QQmlIncubator::statusChanged(status); }; };
};

void* QQmlIncubator_NewQQmlIncubator(int mode){
	return new MyQQmlIncubator(static_cast<QQmlIncubator::IncubationMode>(mode));
}

void QQmlIncubator_Clear(void* ptr){
	static_cast<QQmlIncubator*>(ptr)->clear();
}

void QQmlIncubator_ForceCompletion(void* ptr){
	static_cast<QQmlIncubator*>(ptr)->forceCompletion();
}

int QQmlIncubator_IncubationMode(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->incubationMode();
}

int QQmlIncubator_IsError(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isError();
}

int QQmlIncubator_IsLoading(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isLoading();
}

int QQmlIncubator_IsNull(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isNull();
}

int QQmlIncubator_IsReady(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->isReady();
}

void* QQmlIncubator_Object(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->object();
}

int QQmlIncubator_Status(void* ptr){
	return static_cast<QQmlIncubator*>(ptr)->status();
}

char* QQmlIncubator_ObjectNameAbs(void* ptr){
	return static_cast<MyQQmlIncubator*>(ptr)->objectNameAbs().toUtf8().data();
}

void QQmlIncubator_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQQmlIncubator*>(ptr)->setObjectNameAbs(QString(name));
}

void* QQmlListReference_NewQQmlListReference(){
	return new QQmlListReference();
}

void* QQmlListReference_NewQQmlListReference2(void* object, char* property, void* engine){
	return new QQmlListReference(static_cast<QObject*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
}

int QQmlListReference_Append(void* ptr, void* object){
	return static_cast<QQmlListReference*>(ptr)->append(static_cast<QObject*>(object));
}

void* QQmlListReference_At(void* ptr, int index){
	return static_cast<QQmlListReference*>(ptr)->at(index);
}

int QQmlListReference_CanAppend(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canAppend();
}

int QQmlListReference_CanAt(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canAt();
}

int QQmlListReference_CanClear(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canClear();
}

int QQmlListReference_CanCount(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->canCount();
}

int QQmlListReference_Clear(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->clear();
}

int QQmlListReference_Count(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->count();
}

int QQmlListReference_IsManipulable(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->isManipulable();
}

int QQmlListReference_IsReadable(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->isReadable();
}

int QQmlListReference_IsValid(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->isValid();
}

void* QQmlListReference_ListElementType(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QQmlListReference*>(ptr)->listElementType());
}

void* QQmlListReference_Object(void* ptr){
	return static_cast<QQmlListReference*>(ptr)->object();
}

class MyQQmlNetworkAccessManagerFactory: public QQmlNetworkAccessManagerFactory {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void* QQmlNetworkAccessManagerFactory_Create(void* ptr, void* parent){
	return static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->create(static_cast<QObject*>(parent));
}

void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(void* ptr){
	static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->~QQmlNetworkAccessManagerFactory();
}

char* QQmlNetworkAccessManagerFactory_ObjectNameAbs(void* ptr){
	return static_cast<MyQQmlNetworkAccessManagerFactory*>(ptr)->objectNameAbs().toUtf8().data();
}

void QQmlNetworkAccessManagerFactory_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQQmlNetworkAccessManagerFactory*>(ptr)->setObjectNameAbs(QString(name));
}

void QQmlParserStatus_ClassBegin(void* ptr){
	static_cast<QQmlParserStatus*>(ptr)->classBegin();
}

void QQmlParserStatus_ComponentComplete(void* ptr){
	static_cast<QQmlParserStatus*>(ptr)->componentComplete();
}

void* QQmlProperty_NewQQmlProperty(){
	return new QQmlProperty();
}

void* QQmlProperty_NewQQmlProperty2(void* obj){
	return new QQmlProperty(static_cast<QObject*>(obj));
}

void* QQmlProperty_NewQQmlProperty3(void* obj, void* ctxt){
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlContext*>(ctxt));
}

void* QQmlProperty_NewQQmlProperty4(void* obj, void* engine){
	return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlEngine*>(engine));
}

void* QQmlProperty_NewQQmlProperty5(void* obj, char* name){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name));
}

void* QQmlProperty_NewQQmlProperty6(void* obj, char* name, void* ctxt){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlContext*>(ctxt));
}

void* QQmlProperty_NewQQmlProperty7(void* obj, char* name, void* engine){
	return new QQmlProperty(static_cast<QObject*>(obj), QString(name), static_cast<QQmlEngine*>(engine));
}

void* QQmlProperty_NewQQmlProperty8(void* other){
	return new QQmlProperty(*static_cast<QQmlProperty*>(other));
}

int QQmlProperty_ConnectNotifySignal(void* ptr, void* dest, char* slot){
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), const_cast<const char*>(slot));
}

int QQmlProperty_ConnectNotifySignal2(void* ptr, void* dest, int method){
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), method);
}

int QQmlProperty_HasNotifySignal(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->hasNotifySignal();
}

int QQmlProperty_Index(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->index();
}

int QQmlProperty_IsDesignable(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isDesignable();
}

int QQmlProperty_IsProperty(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isProperty();
}

int QQmlProperty_IsResettable(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isResettable();
}

int QQmlProperty_IsSignalProperty(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isSignalProperty();
}

int QQmlProperty_IsValid(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isValid();
}

int QQmlProperty_IsWritable(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->isWritable();
}

char* QQmlProperty_Name(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->name().toUtf8().data();
}

int QQmlProperty_NeedsNotifySignal(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->needsNotifySignal();
}

void* QQmlProperty_Object(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->object();
}

int QQmlProperty_PropertyType(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->propertyType();
}

int QQmlProperty_PropertyTypeCategory(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->propertyTypeCategory();
}

void* QQmlProperty_QQmlProperty_Read2(void* object, char* name){
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name)));
}

void* QQmlProperty_QQmlProperty_Read3(void* object, char* name, void* ctxt){
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlContext*>(ctxt)));
}

void* QQmlProperty_QQmlProperty_Read4(void* object, char* name, void* engine){
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString(name), static_cast<QQmlEngine*>(engine)));
}

void* QQmlProperty_Read(void* ptr){
	return new QVariant(static_cast<QQmlProperty*>(ptr)->read());
}

int QQmlProperty_Reset(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->reset();
}

int QQmlProperty_Type(void* ptr){
	return static_cast<QQmlProperty*>(ptr)->type();
}

int QQmlProperty_QQmlProperty_Write2(void* object, char* name, void* value){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value));
}

int QQmlProperty_QQmlProperty_Write3(void* object, char* name, void* value, void* ctxt){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value), static_cast<QQmlContext*>(ctxt));
}

int QQmlProperty_QQmlProperty_Write4(void* object, char* name, void* value, void* engine){
	return QQmlProperty::write(static_cast<QObject*>(object), QString(name), *static_cast<QVariant*>(value), static_cast<QQmlEngine*>(engine));
}

int QQmlProperty_Write(void* ptr, void* value){
	return static_cast<QQmlProperty*>(ptr)->write(*static_cast<QVariant*>(value));
}

class MyQQmlPropertyMap: public QQmlPropertyMap {
public:
	MyQQmlPropertyMap(QObject *parent) : QQmlPropertyMap(parent) {};
	void Signal_ValueChanged(const QString & key, const QVariant & value) { callbackQQmlPropertyMapValueChanged(this->objectName().toUtf8().data(), key.toUtf8().data(), new QVariant(value)); };
protected:
};

void* QQmlPropertyMap_NewQQmlPropertyMap(void* parent){
	return new MyQQmlPropertyMap(static_cast<QObject*>(parent));
}

void QQmlPropertyMap_Clear(void* ptr, char* key){
	static_cast<QQmlPropertyMap*>(ptr)->clear(QString(key));
}

int QQmlPropertyMap_Contains(void* ptr, char* key){
	return static_cast<QQmlPropertyMap*>(ptr)->contains(QString(key));
}

int QQmlPropertyMap_Count(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->count();
}

int QQmlPropertyMap_IsEmpty(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->isEmpty();
}

char* QQmlPropertyMap_Keys(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->keys().join(",,,").toUtf8().data();
}

int QQmlPropertyMap_Size(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->size();
}

void* QQmlPropertyMap_Value(void* ptr, char* key){
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->value(QString(key)));
}

void QQmlPropertyMap_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));;
}

void QQmlPropertyMap_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));;
}

void QQmlPropertyMap_DestroyQQmlPropertyMap(void* ptr){
	static_cast<QQmlPropertyMap*>(ptr)->~QQmlPropertyMap();
}

class MyQQmlPropertyValueSource: public QQmlPropertyValueSource {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void QQmlPropertyValueSource_SetTarget(void* ptr, void* property){
	static_cast<QQmlPropertyValueSource*>(ptr)->setTarget(*static_cast<QQmlProperty*>(property));
}

void QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(void* ptr){
	static_cast<QQmlPropertyValueSource*>(ptr)->~QQmlPropertyValueSource();
}

char* QQmlPropertyValueSource_ObjectNameAbs(void* ptr){
	return static_cast<MyQQmlPropertyValueSource*>(ptr)->objectNameAbs().toUtf8().data();
}

void QQmlPropertyValueSource_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQQmlPropertyValueSource*>(ptr)->setObjectNameAbs(QString(name));
}

void* QQmlScriptString_NewQQmlScriptString(){
	return new QQmlScriptString();
}

void* QQmlScriptString_NewQQmlScriptString2(void* other){
	return new QQmlScriptString(*static_cast<QQmlScriptString*>(other));
}

int QQmlScriptString_BooleanLiteral(void* ptr, int ok){
	return static_cast<QQmlScriptString*>(ptr)->booleanLiteral(NULL);
}

int QQmlScriptString_IsEmpty(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->isEmpty();
}

int QQmlScriptString_IsNullLiteral(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->isNullLiteral();
}

int QQmlScriptString_IsUndefinedLiteral(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->isUndefinedLiteral();
}

double QQmlScriptString_NumberLiteral(void* ptr, int ok){
	return static_cast<double>(static_cast<QQmlScriptString*>(ptr)->numberLiteral(NULL));
}

char* QQmlScriptString_StringLiteral(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->stringLiteral().toUtf8().data();
}

