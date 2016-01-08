#define protected public

#include "uitools.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QDir>
#include <QEvent>
#include <QIODevice>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUiLoader>
#include <QWidget>

class MyQUiLoader: public QUiLoader {
public:
	MyQUiLoader(QObject *parent) : QUiLoader(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQUiLoaderTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQUiLoaderChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQUiLoaderCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QUiLoader_NewQUiLoader(void* parent){
	return new MyQUiLoader(static_cast<QObject*>(parent));
}

void QUiLoader_AddPluginPath(void* ptr, char* path){
	static_cast<QUiLoader*>(ptr)->addPluginPath(QString(path));
}

char* QUiLoader_AvailableLayouts(void* ptr){
	return static_cast<QUiLoader*>(ptr)->availableLayouts().join("|").toUtf8().data();
}

char* QUiLoader_AvailableWidgets(void* ptr){
	return static_cast<QUiLoader*>(ptr)->availableWidgets().join("|").toUtf8().data();
}

void QUiLoader_ClearPluginPaths(void* ptr){
	static_cast<QUiLoader*>(ptr)->clearPluginPaths();
}

void* QUiLoader_CreateAction(void* ptr, void* parent, char* name){
	return static_cast<QUiLoader*>(ptr)->createAction(static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateActionGroup(void* ptr, void* parent, char* name){
	return static_cast<QUiLoader*>(ptr)->createActionGroup(static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateLayout(void* ptr, char* className, void* parent, char* name){
	return static_cast<QUiLoader*>(ptr)->createLayout(QString(className), static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateWidget(void* ptr, char* className, void* parent, char* name){
	return static_cast<QUiLoader*>(ptr)->createWidget(QString(className), static_cast<QWidget*>(parent), QString(name));
}

char* QUiLoader_ErrorString(void* ptr){
	return static_cast<QUiLoader*>(ptr)->errorString().toUtf8().data();
}

int QUiLoader_IsLanguageChangeEnabled(void* ptr){
	return static_cast<QUiLoader*>(ptr)->isLanguageChangeEnabled();
}

void* QUiLoader_Load(void* ptr, void* device, void* parentWidget){
	return static_cast<QUiLoader*>(ptr)->load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parentWidget));
}

char* QUiLoader_PluginPaths(void* ptr){
	return static_cast<QUiLoader*>(ptr)->pluginPaths().join("|").toUtf8().data();
}

void QUiLoader_SetLanguageChangeEnabled(void* ptr, int enabled){
	static_cast<QUiLoader*>(ptr)->setLanguageChangeEnabled(enabled != 0);
}

void QUiLoader_SetWorkingDirectory(void* ptr, void* dir){
	static_cast<QUiLoader*>(ptr)->setWorkingDirectory(*static_cast<QDir*>(dir));
}

void* QUiLoader_WorkingDirectory(void* ptr){
	return new QDir(static_cast<QUiLoader*>(ptr)->workingDirectory());
}

void QUiLoader_DestroyQUiLoader(void* ptr){
	static_cast<QUiLoader*>(ptr)->~QUiLoader();
}

void QUiLoader_TimerEvent(void* ptr, void* event){
	static_cast<MyQUiLoader*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QUiLoader_TimerEventDefault(void* ptr, void* event){
	static_cast<QUiLoader*>(ptr)->QUiLoader::timerEvent(static_cast<QTimerEvent*>(event));
}

void QUiLoader_ChildEvent(void* ptr, void* event){
	static_cast<MyQUiLoader*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QUiLoader_ChildEventDefault(void* ptr, void* event){
	static_cast<QUiLoader*>(ptr)->QUiLoader::childEvent(static_cast<QChildEvent*>(event));
}

void QUiLoader_CustomEvent(void* ptr, void* event){
	static_cast<MyQUiLoader*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QUiLoader_CustomEventDefault(void* ptr, void* event){
	static_cast<QUiLoader*>(ptr)->QUiLoader::customEvent(static_cast<QEvent*>(event));
}

