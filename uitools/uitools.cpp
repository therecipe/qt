// +build !minimal

#define protected public
#define private public

#include "uitools.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionGroup>
#include <QChildEvent>
#include <QDir>
#include <QEvent>
#include <QIODevice>
#include <QLayout>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUiLoader>
#include <QWidget>

class MyQUiLoader: public QUiLoader
{
public:
	MyQUiLoader(QObject *parent) : QUiLoader(parent) {};
	QAction * createAction(QObject * parent, const QString & name) { return static_cast<QAction*>(callbackQUiLoader_CreateAction(this, parent, const_cast<char*>(name.toUtf8().prepend("WHITESPACE").constData()+10))); };
	QActionGroup * createActionGroup(QObject * parent, const QString & name) { return static_cast<QActionGroup*>(callbackQUiLoader_CreateActionGroup(this, parent, const_cast<char*>(name.toUtf8().prepend("WHITESPACE").constData()+10))); };
	QLayout * createLayout(const QString & className, QObject * parent, const QString & name) { return static_cast<QLayout*>(callbackQUiLoader_CreateLayout(this, const_cast<char*>(className.toUtf8().prepend("WHITESPACE").constData()+10), parent, const_cast<char*>(name.toUtf8().prepend("WHITESPACE").constData()+10))); };
	QWidget * createWidget(const QString & className, QWidget * parent, const QString & name) { return static_cast<QWidget*>(callbackQUiLoader_CreateWidget(this, const_cast<char*>(className.toUtf8().prepend("WHITESPACE").constData()+10), parent, const_cast<char*>(name.toUtf8().prepend("WHITESPACE").constData()+10))); };
	 ~MyQUiLoader() { callbackQUiLoader_DestroyQUiLoader(this); };
	void timerEvent(QTimerEvent * event) { callbackQUiLoader_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQUiLoader_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQUiLoader_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQUiLoader_CustomEvent(this, event); };
	void deleteLater() { callbackQUiLoader_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQUiLoader_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQUiLoader_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQUiLoader_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQUiLoader_MetaObject(const_cast<MyQUiLoader*>(this))); };
};

void* QUiLoader_NewQUiLoader(void* parent)
{
	return new MyQUiLoader(static_cast<QObject*>(parent));
}

void QUiLoader_AddPluginPath(void* ptr, char* path)
{
	static_cast<QUiLoader*>(ptr)->addPluginPath(QString(path));
}

char* QUiLoader_AvailableLayouts(void* ptr)
{
	return const_cast<char*>(static_cast<QUiLoader*>(ptr)->availableLayouts().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

char* QUiLoader_AvailableWidgets(void* ptr)
{
	return const_cast<char*>(static_cast<QUiLoader*>(ptr)->availableWidgets().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

void QUiLoader_ClearPluginPaths(void* ptr)
{
	static_cast<QUiLoader*>(ptr)->clearPluginPaths();
}

void* QUiLoader_CreateAction(void* ptr, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->createAction(static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateActionDefault(void* ptr, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->QUiLoader::createAction(static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateActionGroup(void* ptr, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->createActionGroup(static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateActionGroupDefault(void* ptr, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->QUiLoader::createActionGroup(static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateLayout(void* ptr, char* className, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->createLayout(QString(className), static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateLayoutDefault(void* ptr, char* className, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->QUiLoader::createLayout(QString(className), static_cast<QObject*>(parent), QString(name));
}

void* QUiLoader_CreateWidget(void* ptr, char* className, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->createWidget(QString(className), static_cast<QWidget*>(parent), QString(name));
}

void* QUiLoader_CreateWidgetDefault(void* ptr, char* className, void* parent, char* name)
{
	return static_cast<QUiLoader*>(ptr)->QUiLoader::createWidget(QString(className), static_cast<QWidget*>(parent), QString(name));
}

char* QUiLoader_ErrorString(void* ptr)
{
	return const_cast<char*>(static_cast<QUiLoader*>(ptr)->errorString().toUtf8().prepend("WHITESPACE").constData()+10);
}

char QUiLoader_IsLanguageChangeEnabled(void* ptr)
{
	return static_cast<QUiLoader*>(ptr)->isLanguageChangeEnabled();
}

void* QUiLoader_Load(void* ptr, void* device, void* parentWidget)
{
	return static_cast<QUiLoader*>(ptr)->load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parentWidget));
}

char* QUiLoader_PluginPaths(void* ptr)
{
	return const_cast<char*>(static_cast<QUiLoader*>(ptr)->pluginPaths().join("|").toUtf8().prepend("WHITESPACE").constData()+10);
}

void QUiLoader_SetLanguageChangeEnabled(void* ptr, char enabled)
{
	static_cast<QUiLoader*>(ptr)->setLanguageChangeEnabled(enabled != 0);
}

void QUiLoader_SetWorkingDirectory(void* ptr, void* dir)
{
	static_cast<QUiLoader*>(ptr)->setWorkingDirectory(*static_cast<QDir*>(dir));
}

void* QUiLoader_WorkingDirectory(void* ptr)
{
	return new QDir(static_cast<QUiLoader*>(ptr)->workingDirectory());
}

void QUiLoader_DestroyQUiLoader(void* ptr)
{
	static_cast<QUiLoader*>(ptr)->~QUiLoader();
}

void QUiLoader_DestroyQUiLoaderDefault(void* ptr)
{

}

void QUiLoader_TimerEvent(void* ptr, void* event)
{
	static_cast<QUiLoader*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QUiLoader_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QUiLoader*>(ptr)->QUiLoader::timerEvent(static_cast<QTimerEvent*>(event));
}

void QUiLoader_ChildEvent(void* ptr, void* event)
{
	static_cast<QUiLoader*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QUiLoader_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QUiLoader*>(ptr)->QUiLoader::childEvent(static_cast<QChildEvent*>(event));
}

void QUiLoader_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QUiLoader*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUiLoader_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QUiLoader*>(ptr)->QUiLoader::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUiLoader_CustomEvent(void* ptr, void* event)
{
	static_cast<QUiLoader*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QUiLoader_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QUiLoader*>(ptr)->QUiLoader::customEvent(static_cast<QEvent*>(event));
}

void QUiLoader_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QUiLoader*>(ptr), "deleteLater");
}

void QUiLoader_DeleteLaterDefault(void* ptr)
{
	static_cast<QUiLoader*>(ptr)->QUiLoader::deleteLater();
}

void QUiLoader_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QUiLoader*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUiLoader_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QUiLoader*>(ptr)->QUiLoader::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QUiLoader_Event(void* ptr, void* e)
{
	return static_cast<QUiLoader*>(ptr)->event(static_cast<QEvent*>(e));
}

char QUiLoader_EventDefault(void* ptr, void* e)
{
	return static_cast<QUiLoader*>(ptr)->QUiLoader::event(static_cast<QEvent*>(e));
}

char QUiLoader_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QUiLoader*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QUiLoader_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QUiLoader*>(ptr)->QUiLoader::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QUiLoader_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QUiLoader*>(ptr)->metaObject());
}

void* QUiLoader_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QUiLoader*>(ptr)->QUiLoader::metaObject());
}

