#include "qpluginloader.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QLibrary>
#include <QPluginLoader>
#include "_cgo_export.h"

class MyQPluginLoader: public QPluginLoader {
public:
};

char* QPluginLoader_FileName(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->fileName().toUtf8().data();
}

int QPluginLoader_LoadHints(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->loadHints();
}

void QPluginLoader_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QPluginLoader*>(ptr)->setFileName(QString(fileName));
}

void QPluginLoader_SetLoadHints(QtObjectPtr ptr, int loadHints){
	static_cast<QPluginLoader*>(ptr)->setLoadHints(static_cast<QLibrary::LoadHint>(loadHints));
}

QtObjectPtr QPluginLoader_NewQPluginLoader(QtObjectPtr parent){
	return new QPluginLoader(static_cast<QObject*>(parent));
}

QtObjectPtr QPluginLoader_NewQPluginLoader2(char* fileName, QtObjectPtr parent){
	return new QPluginLoader(QString(fileName), static_cast<QObject*>(parent));
}

char* QPluginLoader_ErrorString(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->errorString().toUtf8().data();
}

QtObjectPtr QPluginLoader_Instance(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->instance();
}

int QPluginLoader_IsLoaded(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->isLoaded();
}

int QPluginLoader_Load(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->load();
}

int QPluginLoader_Unload(QtObjectPtr ptr){
	return static_cast<QPluginLoader*>(ptr)->unload();
}

void QPluginLoader_DestroyQPluginLoader(QtObjectPtr ptr){
	static_cast<QPluginLoader*>(ptr)->~QPluginLoader();
}

