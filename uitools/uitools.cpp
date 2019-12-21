// +build !minimal

#define protected public
#define private public

#include "uitools.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionGroup>
#include <QAudioSystemPlugin>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QDir>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QIODevice>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QString>
#include <QTimerEvent>
#include <QUiLoader>
#include <QWidget>
#include <QWindow>

class MyQUiLoader: public QUiLoader
{
public:
	MyQUiLoader(QObject *parent = Q_NULLPTR) : QUiLoader(parent) {QUiLoader_QUiLoader_QRegisterMetaType();};
	QAction * createAction(QObject * parent, const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtUiTools_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return static_cast<QAction*>(callbackQUiLoader_CreateAction(this, parent, namePacked)); };
	QActionGroup * createActionGroup(QObject * parent, const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtUiTools_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return static_cast<QActionGroup*>(callbackQUiLoader_CreateActionGroup(this, parent, namePacked)); };
	QLayout * createLayout(const QString & className, QObject * parent, const QString & name) { QByteArray* td80a05 = new QByteArray(className.toUtf8()); QtUiTools_PackedString classNamePacked = { const_cast<char*>(td80a05->prepend("WHITESPACE").constData()+10), td80a05->size()-10, td80a05 };QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtUiTools_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return static_cast<QLayout*>(callbackQUiLoader_CreateLayout(this, classNamePacked, parent, namePacked)); };
	QWidget * createWidget(const QString & className, QWidget * parent, const QString & name) { QByteArray* td80a05 = new QByteArray(className.toUtf8()); QtUiTools_PackedString classNamePacked = { const_cast<char*>(td80a05->prepend("WHITESPACE").constData()+10), td80a05->size()-10, td80a05 };QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtUiTools_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return static_cast<QWidget*>(callbackQUiLoader_CreateWidget(this, classNamePacked, parent, namePacked)); };
	 ~MyQUiLoader() { callbackQUiLoader_DestroyQUiLoader(this); };
	void childEvent(QChildEvent * event) { callbackQUiLoader_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQUiLoader_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQUiLoader_CustomEvent(this, event); };
	void deleteLater() { callbackQUiLoader_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQUiLoader_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQUiLoader_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQUiLoader_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQUiLoader_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQUiLoader_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtUiTools_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQUiLoader_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQUiLoader_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QUiLoader*)
Q_DECLARE_METATYPE(MyQUiLoader*)

int QUiLoader_QUiLoader_QRegisterMetaType(){qRegisterMetaType<QUiLoader*>(); return qRegisterMetaType<MyQUiLoader*>();}

void* QUiLoader_NewQUiLoader(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQUiLoader(static_cast<QWindow*>(parent));
	} else {
		return new MyQUiLoader(static_cast<QObject*>(parent));
	}
}

void QUiLoader_AddPluginPath(void* ptr, struct QtUiTools_PackedString path)
{
	static_cast<QUiLoader*>(ptr)->addPluginPath(QString::fromUtf8(path.data, path.len));
}

struct QtUiTools_PackedString QUiLoader_AvailableLayouts(void* ptr)
{
	return ({ QByteArray* te9f98b = new QByteArray(static_cast<QUiLoader*>(ptr)->availableLayouts().join("¡¦!").toUtf8()); QtUiTools_PackedString { const_cast<char*>(te9f98b->prepend("WHITESPACE").constData()+10), te9f98b->size()-10, te9f98b }; });
}

struct QtUiTools_PackedString QUiLoader_AvailableWidgets(void* ptr)
{
	return ({ QByteArray* tf979e0 = new QByteArray(static_cast<QUiLoader*>(ptr)->availableWidgets().join("¡¦!").toUtf8()); QtUiTools_PackedString { const_cast<char*>(tf979e0->prepend("WHITESPACE").constData()+10), tf979e0->size()-10, tf979e0 }; });
}

void QUiLoader_ClearPluginPaths(void* ptr)
{
	static_cast<QUiLoader*>(ptr)->clearPluginPaths();
}

void* QUiLoader_CreateAction(void* ptr, void* parent, struct QtUiTools_PackedString name)
{
	return static_cast<QUiLoader*>(ptr)->createAction(static_cast<QObject*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateActionDefault(void* ptr, void* parent, struct QtUiTools_PackedString name)
{
		return static_cast<QUiLoader*>(ptr)->QUiLoader::createAction(static_cast<QObject*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateActionGroup(void* ptr, void* parent, struct QtUiTools_PackedString name)
{
	return static_cast<QUiLoader*>(ptr)->createActionGroup(static_cast<QObject*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateActionGroupDefault(void* ptr, void* parent, struct QtUiTools_PackedString name)
{
		return static_cast<QUiLoader*>(ptr)->QUiLoader::createActionGroup(static_cast<QObject*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateLayout(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name)
{
	return static_cast<QUiLoader*>(ptr)->createLayout(QString::fromUtf8(className.data, className.len), static_cast<QObject*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateLayoutDefault(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name)
{
		return static_cast<QUiLoader*>(ptr)->QUiLoader::createLayout(QString::fromUtf8(className.data, className.len), static_cast<QObject*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateWidget(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name)
{
	return static_cast<QUiLoader*>(ptr)->createWidget(QString::fromUtf8(className.data, className.len), static_cast<QWidget*>(parent), QString::fromUtf8(name.data, name.len));
}

void* QUiLoader_CreateWidgetDefault(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name)
{
		return static_cast<QUiLoader*>(ptr)->QUiLoader::createWidget(QString::fromUtf8(className.data, className.len), static_cast<QWidget*>(parent), QString::fromUtf8(name.data, name.len));
}

struct QtUiTools_PackedString QUiLoader_ErrorString(void* ptr)
{
	return ({ QByteArray* tadc93d = new QByteArray(static_cast<QUiLoader*>(ptr)->errorString().toUtf8()); QtUiTools_PackedString { const_cast<char*>(tadc93d->prepend("WHITESPACE").constData()+10), tadc93d->size()-10, tadc93d }; });
}

char QUiLoader_IsLanguageChangeEnabled(void* ptr)
{
	return static_cast<QUiLoader*>(ptr)->isLanguageChangeEnabled();
}

void* QUiLoader_Load(void* ptr, void* device, void* parentWidget)
{
	return static_cast<QUiLoader*>(ptr)->load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parentWidget));
}

struct QtUiTools_PackedString QUiLoader_PluginPaths(void* ptr)
{
	return ({ QByteArray* t84d911 = new QByteArray(static_cast<QUiLoader*>(ptr)->pluginPaths().join("¡¦!").toUtf8()); QtUiTools_PackedString { const_cast<char*>(t84d911->prepend("WHITESPACE").constData()+10), t84d911->size()-10, t84d911 }; });
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
	Q_UNUSED(ptr);

}

void* QUiLoader___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QUiLoader___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QUiLoader___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QUiLoader___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUiLoader___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QUiLoader___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QUiLoader___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QUiLoader___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QUiLoader___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QUiLoader___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QUiLoader___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QUiLoader___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QUiLoader_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QUiLoader*>(ptr)->QUiLoader::childEvent(static_cast<QChildEvent*>(event));
}

void QUiLoader_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QUiLoader*>(ptr)->QUiLoader::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUiLoader_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QUiLoader*>(ptr)->QUiLoader::customEvent(static_cast<QEvent*>(event));
}

void QUiLoader_DeleteLaterDefault(void* ptr)
{
		static_cast<QUiLoader*>(ptr)->QUiLoader::deleteLater();
}

void QUiLoader_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QUiLoader*>(ptr)->QUiLoader::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QUiLoader_EventDefault(void* ptr, void* e)
{
		return static_cast<QUiLoader*>(ptr)->QUiLoader::event(static_cast<QEvent*>(e));
}

char QUiLoader_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QUiLoader*>(ptr)->QUiLoader::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QUiLoader_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QUiLoader*>(ptr)->QUiLoader::metaObject());
}

void QUiLoader_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QUiLoader*>(ptr)->QUiLoader::timerEvent(static_cast<QTimerEvent*>(event));
}

