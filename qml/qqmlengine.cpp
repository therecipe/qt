#include "qqmlengine.h"
#include <QObject>
#include <QQmlImageProviderBase>
#include <QQmlIncubationController>
#include <QQmlNetworkAccessManagerFactory>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlContext>
#include <QQmlEngine>
#include "_cgo_export.h"

class MyQQmlEngine: public QQmlEngine {
public:
void Signal_Quit(){callbackQQmlEngineQuit(this->objectName().toUtf8().data());};
};

char* QQmlEngine_OfflineStoragePath(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->offlineStoragePath().toUtf8().data();
}

void QQmlEngine_SetOfflineStoragePath(QtObjectPtr ptr, char* dir){
	static_cast<QQmlEngine*>(ptr)->setOfflineStoragePath(QString(dir));
}

QtObjectPtr QQmlEngine_NewQQmlEngine(QtObjectPtr parent){
	return new QQmlEngine(static_cast<QObject*>(parent));
}

void QQmlEngine_AddImageProvider(QtObjectPtr ptr, char* providerId, QtObjectPtr provider){
	static_cast<QQmlEngine*>(ptr)->addImageProvider(QString(providerId), static_cast<QQmlImageProviderBase*>(provider));
}

void QQmlEngine_AddImportPath(QtObjectPtr ptr, char* path){
	static_cast<QQmlEngine*>(ptr)->addImportPath(QString(path));
}

void QQmlEngine_AddPluginPath(QtObjectPtr ptr, char* path){
	static_cast<QQmlEngine*>(ptr)->addPluginPath(QString(path));
}

char* QQmlEngine_BaseUrl(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->baseUrl().toString().toUtf8().data();
}

void QQmlEngine_ClearComponentCache(QtObjectPtr ptr){
	static_cast<QQmlEngine*>(ptr)->clearComponentCache();
}

QtObjectPtr QQmlEngine_QQmlEngine_ContextForObject(QtObjectPtr object){
	return QQmlEngine::contextForObject(static_cast<QObject*>(object));
}

QtObjectPtr QQmlEngine_ImageProvider(QtObjectPtr ptr, char* providerId){
	return static_cast<QQmlEngine*>(ptr)->imageProvider(QString(providerId));
}

char* QQmlEngine_ImportPathList(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->importPathList().join("|").toUtf8().data();
}

QtObjectPtr QQmlEngine_IncubationController(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->incubationController();
}

QtObjectPtr QQmlEngine_NetworkAccessManager(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->networkAccessManager();
}

QtObjectPtr QQmlEngine_NetworkAccessManagerFactory(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->networkAccessManagerFactory();
}

int QQmlEngine_QQmlEngine_ObjectOwnership(QtObjectPtr object){
	return QQmlEngine::objectOwnership(static_cast<QObject*>(object));
}

int QQmlEngine_OutputWarningsToStandardError(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->outputWarningsToStandardError();
}

char* QQmlEngine_PluginPathList(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->pluginPathList().join("|").toUtf8().data();
}

void QQmlEngine_ConnectQuit(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));;
}

void QQmlEngine_DisconnectQuit(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));;
}

void QQmlEngine_RemoveImageProvider(QtObjectPtr ptr, char* providerId){
	static_cast<QQmlEngine*>(ptr)->removeImageProvider(QString(providerId));
}

QtObjectPtr QQmlEngine_RootContext(QtObjectPtr ptr){
	return static_cast<QQmlEngine*>(ptr)->rootContext();
}

void QQmlEngine_SetBaseUrl(QtObjectPtr ptr, char* url){
	static_cast<QQmlEngine*>(ptr)->setBaseUrl(QUrl(QString(url)));
}

void QQmlEngine_QQmlEngine_SetContextForObject(QtObjectPtr object, QtObjectPtr context){
	QQmlEngine::setContextForObject(static_cast<QObject*>(object), static_cast<QQmlContext*>(context));
}

void QQmlEngine_SetImportPathList(QtObjectPtr ptr, char* paths){
	static_cast<QQmlEngine*>(ptr)->setImportPathList(QString(paths).split("|", QString::SkipEmptyParts));
}

void QQmlEngine_SetIncubationController(QtObjectPtr ptr, QtObjectPtr controller){
	static_cast<QQmlEngine*>(ptr)->setIncubationController(static_cast<QQmlIncubationController*>(controller));
}

void QQmlEngine_SetNetworkAccessManagerFactory(QtObjectPtr ptr, QtObjectPtr factory){
	static_cast<QQmlEngine*>(ptr)->setNetworkAccessManagerFactory(static_cast<QQmlNetworkAccessManagerFactory*>(factory));
}

void QQmlEngine_QQmlEngine_SetObjectOwnership(QtObjectPtr object, int ownership){
	QQmlEngine::setObjectOwnership(static_cast<QObject*>(object), static_cast<QQmlEngine::ObjectOwnership>(ownership));
}

void QQmlEngine_SetOutputWarningsToStandardError(QtObjectPtr ptr, int enabled){
	static_cast<QQmlEngine*>(ptr)->setOutputWarningsToStandardError(enabled != 0);
}

void QQmlEngine_SetPluginPathList(QtObjectPtr ptr, char* paths){
	static_cast<QQmlEngine*>(ptr)->setPluginPathList(QString(paths).split("|", QString::SkipEmptyParts));
}

void QQmlEngine_TrimComponentCache(QtObjectPtr ptr){
	static_cast<QQmlEngine*>(ptr)->trimComponentCache();
}

void QQmlEngine_DestroyQQmlEngine(QtObjectPtr ptr){
	static_cast<QQmlEngine*>(ptr)->~QQmlEngine();
}

