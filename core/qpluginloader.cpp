#include "qpluginloader.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QJsonObject>
#include <QLibrary>
#include <QPluginLoader>
#include "_cgo_export.h"

class MyQPluginLoader: public QPluginLoader {
public:
};

char* QPluginLoader_FileName(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->fileName().toUtf8().data();
}

int QPluginLoader_LoadHints(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->loadHints();
}

void QPluginLoader_SetFileName(void* ptr, char* fileName){
	static_cast<QPluginLoader*>(ptr)->setFileName(QString(fileName));
}

void QPluginLoader_SetLoadHints(void* ptr, int loadHints){
	static_cast<QPluginLoader*>(ptr)->setLoadHints(static_cast<QLibrary::LoadHint>(loadHints));
}

void* QPluginLoader_NewQPluginLoader(void* parent){
	return new QPluginLoader(static_cast<QObject*>(parent));
}

void* QPluginLoader_NewQPluginLoader2(char* fileName, void* parent){
	return new QPluginLoader(QString(fileName), static_cast<QObject*>(parent));
}

char* QPluginLoader_ErrorString(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->errorString().toUtf8().data();
}

void* QPluginLoader_Instance(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->instance();
}

int QPluginLoader_IsLoaded(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->isLoaded();
}

int QPluginLoader_Load(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->load();
}

void* QPluginLoader_MetaData(void* ptr){
	return new QJsonObject(static_cast<QPluginLoader*>(ptr)->metaData());
}

int QPluginLoader_Unload(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->unload();
}

void QPluginLoader_DestroyQPluginLoader(void* ptr){
	static_cast<QPluginLoader*>(ptr)->~QPluginLoader();
}

