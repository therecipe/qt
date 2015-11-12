#include "qqmlengine.h"
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QQmlNetworkAccessManagerFactory>
#include <QVariant>
#include <QQmlIncubationController>
#include <QQmlImageProviderBase>
#include <QQmlContext>
#include <QQmlEngine>
#include "_cgo_export.h"

class MyQQmlEngine: public QQmlEngine {
public:
void Signal_Quit(){callbackQQmlEngineQuit(this->objectName().toUtf8().data());};
};

char* QQmlEngine_OfflineStoragePath(void* ptr){
	return static_cast<QQmlEngine*>(ptr)->offlineStoragePath().toUtf8().data();
}

void QQmlEngine_SetOfflineStoragePath(void* ptr, char* dir){
	static_cast<QQmlEngine*>(ptr)->setOfflineStoragePath(QString(dir));
}

void* QQmlEngine_NewQQmlEngine(void* parent){
	return new QQmlEngine(static_cast<QObject*>(parent));
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
	return static_cast<QQmlEngine*>(ptr)->importPathList().join("|").toUtf8().data();
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
	return static_cast<QQmlEngine*>(ptr)->pluginPathList().join("|").toUtf8().data();
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
	static_cast<QQmlEngine*>(ptr)->setImportPathList(QString(paths).split("|", QString::SkipEmptyParts));
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
	static_cast<QQmlEngine*>(ptr)->setPluginPathList(QString(paths).split("|", QString::SkipEmptyParts));
}

void QQmlEngine_TrimComponentCache(void* ptr){
	static_cast<QQmlEngine*>(ptr)->trimComponentCache();
}

void QQmlEngine_DestroyQQmlEngine(void* ptr){
	static_cast<QQmlEngine*>(ptr)->~QQmlEngine();
}

