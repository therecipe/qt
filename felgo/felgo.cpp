// +build !minimal

#define protected public
#define private public

#include "felgo.h"
#include "_cgo_export.h"

#include <FelgoApplication>
#include <FelgoLiveClient>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFileSelector>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQmlEngine>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWidget>
#include <QWindow>

class MyFelgoApplication: public FelgoApplication
{
public:
	MyFelgoApplication(QObject *parent = Q_NULLPTR) : FelgoApplication(parent) {FelgoApplication_FelgoApplication_QRegisterMetaType();};
	void objectCreatedHandler(QObject * object, const QUrl & url) { callbackFelgoApplication_ObjectCreatedHandler(this, object, const_cast<QUrl*>(&url)); };
	bool event(QEvent * e) { return callbackFelgoApplication_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackFelgoApplication_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackFelgoApplication_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackFelgoApplication_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackFelgoApplication_CustomEvent(this, event); };
	void deleteLater() { callbackFelgoApplication_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackFelgoApplication_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackFelgoApplication_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtFelgo_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackFelgoApplication_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackFelgoApplication_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackFelgoApplication_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyFelgoApplication*)

int FelgoApplication_FelgoApplication_QRegisterMetaType(){qRegisterMetaType<FelgoApplication*>(); return qRegisterMetaType<MyFelgoApplication*>();}

void FelgoApplication_ObjectCreatedHandler(void* ptr, void* object, void* url)
{
	QMetaObject::invokeMethod(static_cast<FelgoApplication*>(ptr), "objectCreatedHandler", Q_ARG(QObject*, static_cast<QObject*>(object)), Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void FelgoApplication_ObjectCreatedHandlerDefault(void* ptr, void* object, void* url)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::objectCreatedHandler(static_cast<QObject*>(object), *static_cast<QUrl*>(url));
}

void* FelgoApplication_NewFelgoApplication(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyFelgoApplication(static_cast<QWindow*>(parent));
	} else {
		return new MyFelgoApplication(static_cast<QObject*>(parent));
	}
}

void FelgoApplication_AddCustomImportPaths(void* ptr)
{
	static_cast<FelgoApplication*>(ptr)->addCustomImportPaths();
}

void FelgoApplication_AddImportPath(void* ptr, struct QtFelgo_PackedString path)
{
	static_cast<FelgoApplication*>(ptr)->addImportPath(QString::fromUtf8(path.data, path.len));
}

struct QtFelgo_PackedString FelgoApplication_AdjustImportPathToCanonical(void* ptr, struct QtFelgo_PackedString vqs)
{
	return ({ QByteArray t58f38b = static_cast<FelgoApplication*>(ptr)->adjustImportPathToCanonical(QString::fromUtf8(vqs.data, vqs.len)).toUtf8(); QtFelgo_PackedString { const_cast<char*>(t58f38b.prepend("WHITESPACE").constData()+10), t58f38b.size()-10 }; });
}

double FelgoApplication_ContentScale(void* ptr)
{
	return static_cast<FelgoApplication*>(ptr)->contentScale();
}

struct QtFelgo_PackedList FelgoApplication_ContentScaleThresholds(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValued38115 = new QMap<QString, QVariant>(static_cast<FelgoApplication*>(ptr)->contentScaleThresholds()); QtFelgo_PackedList { tmpValued38115, tmpValued38115->size() }; });
}

struct QtFelgo_PackedString FelgoApplication_FileSelectorList(void* ptr)
{
	return ({ QByteArray t3dc72b = static_cast<FelgoApplication*>(ptr)->fileSelectorList().join("¡¦!").toUtf8(); QtFelgo_PackedString { const_cast<char*>(t3dc72b.prepend("WHITESPACE").constData()+10), t3dc72b.size()-10 }; });
}

void FelgoApplication_Initialize(void* ptr, void* engine)
{
	static_cast<FelgoApplication*>(ptr)->initialize(static_cast<QQmlEngine*>(engine));
}

void FelgoApplication_InitializeEngine(void* ptr, void* engine, char* uri)
{
	static_cast<FelgoApplication*>(ptr)->initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

struct QtFelgo_PackedString FelgoApplication_MainQmlFileName(void* ptr)
{
	return ({ QByteArray t188539 = static_cast<FelgoApplication*>(ptr)->mainQmlFileName().toUtf8(); QtFelgo_PackedString { const_cast<char*>(t188539.prepend("WHITESPACE").constData()+10), t188539.size()-10 }; });
}

void* FelgoApplication_QmlEngine(void* ptr)
{
	return static_cast<FelgoApplication*>(ptr)->qmlEngine();
}

void FelgoApplication_RegisterTypes(void* ptr, char* uri)
{
	static_cast<FelgoApplication*>(ptr)->registerTypes(const_cast<const char*>(uri));
}

double FelgoApplication_ScaleFactorForImages(void* ptr)
{
	return static_cast<FelgoApplication*>(ptr)->scaleFactorForImages();
}

void FelgoApplication_SetContentScaleAndFileSelectors(void* ptr, double contentScale)
{
	static_cast<FelgoApplication*>(ptr)->setContentScaleAndFileSelectors(contentScale);
}

void FelgoApplication_SetContentScaleThresholds(void* ptr, void* contentScaleThresholds)
{
	static_cast<FelgoApplication*>(ptr)->setContentScaleThresholds(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(contentScaleThresholds); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void FelgoApplication_SetDefaultContentScalingAndFileSelectors(void* ptr)
{
	static_cast<FelgoApplication*>(ptr)->setDefaultContentScalingAndFileSelectors();
}

void FelgoApplication_SetFileSelectorList(void* ptr, struct QtFelgo_PackedString fileSelectorList)
{
	static_cast<FelgoApplication*>(ptr)->setFileSelectorList(QString::fromUtf8(fileSelectorList.data, fileSelectorList.len).split("¡¦!", QString::SkipEmptyParts));
}

void FelgoApplication_SetGameWindowIsItem(void* ptr)
{
	static_cast<FelgoApplication*>(ptr)->setGameWindowIsItem();
}

void FelgoApplication_SetLicenseKey(void* ptr, struct QtFelgo_PackedString licenseKey)
{
	static_cast<FelgoApplication*>(ptr)->setLicenseKey(QString::fromUtf8(licenseKey.data, licenseKey.len));
}

void FelgoApplication_SetMainQmlFileName(void* ptr, struct QtFelgo_PackedString qmlFileName)
{
	static_cast<FelgoApplication*>(ptr)->setMainQmlFileName(QString::fromUtf8(qmlFileName.data, qmlFileName.len));
}

void FelgoApplication_SetPreservePlatformFonts(void* ptr, char preserve)
{
	static_cast<FelgoApplication*>(ptr)->setPreservePlatformFonts(preserve != 0);
}

void* FelgoApplication_VplayFileSelector(void* ptr)
{
	return static_cast<QFileSelector*>(&static_cast<FelgoApplication*>(ptr)->vplayFileSelector());
}

void FelgoApplication_DestroyFelgoApplication(void* ptr)
{
	static_cast<FelgoApplication*>(ptr)->~FelgoApplication();
}

void* FelgoApplication___contentScaleThresholds_atList(void* ptr, struct QtFelgo_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void FelgoApplication___contentScaleThresholds_setList(void* ptr, struct QtFelgo_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* FelgoApplication___contentScaleThresholds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtFelgo_PackedList FelgoApplication___contentScaleThresholds_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtFelgo_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* FelgoApplication___setContentScaleThresholds_contentScaleThresholds_atList(void* ptr, struct QtFelgo_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void FelgoApplication___setContentScaleThresholds_contentScaleThresholds_setList(void* ptr, struct QtFelgo_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* FelgoApplication___setContentScaleThresholds_contentScaleThresholds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtFelgo_PackedList FelgoApplication___setContentScaleThresholds_contentScaleThresholds_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtFelgo_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtFelgo_PackedString FelgoApplication_____contentScaleThresholds_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtFelgo_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void FelgoApplication_____contentScaleThresholds_keyList_setList(void* ptr, struct QtFelgo_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* FelgoApplication_____contentScaleThresholds_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtFelgo_PackedString FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtFelgo_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_setList(void* ptr, struct QtFelgo_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* FelgoApplication___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void FelgoApplication___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* FelgoApplication___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* FelgoApplication___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoApplication___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoApplication___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* FelgoApplication___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoApplication___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoApplication___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* FelgoApplication___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoApplication___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoApplication___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* FelgoApplication___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoApplication___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoApplication___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char FelgoApplication_EventDefault(void* ptr, void* e)
{
		return static_cast<FelgoApplication*>(ptr)->FelgoApplication::event(static_cast<QEvent*>(e));
}

char FelgoApplication_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<FelgoApplication*>(ptr)->FelgoApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void FelgoApplication_ChildEventDefault(void* ptr, void* event)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::childEvent(static_cast<QChildEvent*>(event));
}

void FelgoApplication_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void FelgoApplication_CustomEventDefault(void* ptr, void* event)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::customEvent(static_cast<QEvent*>(event));
}

void FelgoApplication_DeleteLaterDefault(void* ptr)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::deleteLater();
}

void FelgoApplication_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void FelgoApplication_TimerEventDefault(void* ptr, void* event)
{
		static_cast<FelgoApplication*>(ptr)->FelgoApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

void* FelgoApplication_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<FelgoApplication*>(ptr)->FelgoApplication::metaObject());
}

class MyFelgoLiveClient: public FelgoLiveClient
{
public:
	MyFelgoLiveClient(QQmlEngine *engine = Q_NULLPTR, QObject *parent = Q_NULLPTR) : FelgoLiveClient(engine, parent) {FelgoLiveClient_FelgoLiveClient_QRegisterMetaType();};
	void clearProject(const QString & projectName) { QByteArray t112427 = projectName.toUtf8(); QtFelgo_PackedString projectNamePacked = { const_cast<char*>(t112427.prepend("WHITESPACE").constData()+10), t112427.size()-10 };callbackFelgoLiveClient_ClearProject(this, projectNamePacked); };
	void clearProjects() { callbackFelgoLiveClient_ClearProjects(this); };
	void Signal_ClientNameChanged(const QString & clientName) { QByteArray tf85f74 = clientName.toUtf8(); QtFelgo_PackedString clientNamePacked = { const_cast<char*>(tf85f74.prepend("WHITESPACE").constData()+10), tf85f74.size()-10 };callbackFelgoLiveClient_ClientNameChanged(this, clientNamePacked); };
	void connectAsLocalDesktopClient(const QString & serverUrl) { QByteArray td73dcf = serverUrl.toUtf8(); QtFelgo_PackedString serverUrlPacked = { const_cast<char*>(td73dcf.prepend("WHITESPACE").constData()+10), td73dcf.size()-10 };callbackFelgoLiveClient_ConnectAsLocalDesktopClient(this, serverUrlPacked); };
	void connectToLocalServer(const QString & serverUrl, const QString & userId, const QString & deviceId) { QByteArray td73dcf = serverUrl.toUtf8(); QtFelgo_PackedString serverUrlPacked = { const_cast<char*>(td73dcf.prepend("WHITESPACE").constData()+10), td73dcf.size()-10 };QByteArray tdb3666 = userId.toUtf8(); QtFelgo_PackedString userIdPacked = { const_cast<char*>(tdb3666.prepend("WHITESPACE").constData()+10), tdb3666.size()-10 };QByteArray t3f3fee = deviceId.toUtf8(); QtFelgo_PackedString deviceIdPacked = { const_cast<char*>(t3f3fee.prepend("WHITESPACE").constData()+10), t3f3fee.size()-10 };callbackFelgoLiveClient_ConnectToLocalServer(this, serverUrlPacked, userIdPacked, deviceIdPacked); };
	void disconnectWebReceiver() { callbackFelgoLiveClient_DisconnectWebReceiver(this); };
	void modifyLoadedDocument() { callbackFelgoLiveClient_ModifyLoadedDocument(this); };
	void onDocumentUpdated(const QString & document, const QString & content) { QByteArray t4f8278 = document.toUtf8(); QtFelgo_PackedString documentPacked = { const_cast<char*>(t4f8278.prepend("WHITESPACE").constData()+10), t4f8278.size()-10 };QByteArray t040f06 = content.toUtf8(); QtFelgo_PackedString contentPacked = { const_cast<char*>(t040f06.prepend("WHITESPACE").constData()+10), t040f06.size()-10 };callbackFelgoLiveClient_OnDocumentUpdated(this, documentPacked, contentPacked); };
	void onProjectChanged(const QString & projectName, const QString & projectMainFile) { QByteArray t112427 = projectName.toUtf8(); QtFelgo_PackedString projectNamePacked = { const_cast<char*>(t112427.prepend("WHITESPACE").constData()+10), t112427.size()-10 };QByteArray t1e043c = projectMainFile.toUtf8(); QtFelgo_PackedString projectMainFilePacked = { const_cast<char*>(t1e043c.prepend("WHITESPACE").constData()+10), t1e043c.size()-10 };callbackFelgoLiveClient_OnProjectChanged(this, projectNamePacked, projectMainFilePacked); };
	void Signal_PendingProject(const QString & projectName, const QString & projectMainFile) { QByteArray t112427 = projectName.toUtf8(); QtFelgo_PackedString projectNamePacked = { const_cast<char*>(t112427.prepend("WHITESPACE").constData()+10), t112427.size()-10 };QByteArray t1e043c = projectMainFile.toUtf8(); QtFelgo_PackedString projectMainFilePacked = { const_cast<char*>(t1e043c.prepend("WHITESPACE").constData()+10), t1e043c.size()-10 };callbackFelgoLiveClient_PendingProject(this, projectNamePacked, projectMainFilePacked); };
	void prepareLoadingDocument() { callbackFelgoLiveClient_PrepareLoadingDocument(this); };
	void Signal_WebReceiverConnectionRefused(const QString & reason) { QByteArray t7947e9 = reason.toUtf8(); QtFelgo_PackedString reasonPacked = { const_cast<char*>(t7947e9.prepend("WHITESPACE").constData()+10), t7947e9.size()-10 };callbackFelgoLiveClient_WebReceiverConnectionRefused(this, reasonPacked); };
	void Signal_ReceivedMatchId(const QString & matchId) { QByteArray t0b9296 = matchId.toUtf8(); QtFelgo_PackedString matchIdPacked = { const_cast<char*>(t0b9296.prepend("WHITESPACE").constData()+10), t0b9296.size()-10 };callbackFelgoLiveClient_ReceivedMatchId(this, matchIdPacked); };
	void shakeDetected() { callbackFelgoLiveClient_ShakeDetected(this); };
	void Signal_WebReceiverConnected() { callbackFelgoLiveClient_WebReceiverConnected(this); };
	bool event(QEvent * e) { return callbackFelgoLiveClient_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackFelgoLiveClient_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackFelgoLiveClient_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackFelgoLiveClient_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackFelgoLiveClient_CustomEvent(this, event); };
	void deleteLater() { callbackFelgoLiveClient_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackFelgoLiveClient_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackFelgoLiveClient_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtFelgo_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackFelgoLiveClient_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackFelgoLiveClient_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackFelgoLiveClient_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyFelgoLiveClient*)

int FelgoLiveClient_FelgoLiveClient_QRegisterMetaType(){qRegisterMetaType<FelgoLiveClient*>(); return qRegisterMetaType<MyFelgoLiveClient*>();}

void* FelgoLiveClient_NewFelgoLiveClient(void* engine, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
	} else {
		return new MyFelgoLiveClient(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
	}
}

void FelgoLiveClient_AddShakeDetection(void* ptr)
{
	static_cast<FelgoLiveClient*>(ptr)->addShakeDetection();
}

void FelgoLiveClient_ClearProject(void* ptr, struct QtFelgo_PackedString projectName)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "clearProject", Q_ARG(const QString, QString::fromUtf8(projectName.data, projectName.len)));
}

void FelgoLiveClient_ClearProjectDefault(void* ptr, struct QtFelgo_PackedString projectName)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::clearProject(QString::fromUtf8(projectName.data, projectName.len));
}

void FelgoLiveClient_ClearProjects(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "clearProjects");
}

void FelgoLiveClient_ClearProjectsDefault(void* ptr)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::clearProjects();
}

struct QtFelgo_PackedString FelgoLiveClient_ClientName(void* ptr)
{
	return ({ QByteArray t334812 = static_cast<FelgoLiveClient*>(ptr)->clientName().toUtf8(); QtFelgo_PackedString { const_cast<char*>(t334812.prepend("WHITESPACE").constData()+10), t334812.size()-10 }; });
}

void FelgoLiveClient_ConnectClientNameChanged(void* ptr)
{
	QObject::connect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &)>(&FelgoLiveClient::clientNameChanged), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &)>(&MyFelgoLiveClient::Signal_ClientNameChanged));
}

void FelgoLiveClient_DisconnectClientNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &)>(&FelgoLiveClient::clientNameChanged), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &)>(&MyFelgoLiveClient::Signal_ClientNameChanged));
}

void FelgoLiveClient_ClientNameChanged(void* ptr, struct QtFelgo_PackedString clientName)
{
	static_cast<FelgoLiveClient*>(ptr)->clientNameChanged(QString::fromUtf8(clientName.data, clientName.len));
}

void FelgoLiveClient_ConnectAsLocalDesktopClient(void* ptr, struct QtFelgo_PackedString serverUrl)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "connectAsLocalDesktopClient", Q_ARG(const QString, QString::fromUtf8(serverUrl.data, serverUrl.len)));
}

void FelgoLiveClient_ConnectAsLocalDesktopClientDefault(void* ptr, struct QtFelgo_PackedString serverUrl)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::connectAsLocalDesktopClient(QString::fromUtf8(serverUrl.data, serverUrl.len));
}

void FelgoLiveClient_ConnectToLocalServer(void* ptr, struct QtFelgo_PackedString serverUrl, struct QtFelgo_PackedString userId, struct QtFelgo_PackedString deviceId)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "connectToLocalServer", Q_ARG(const QString, QString::fromUtf8(serverUrl.data, serverUrl.len)), Q_ARG(const QString, QString::fromUtf8(userId.data, userId.len)), Q_ARG(const QString, QString::fromUtf8(deviceId.data, deviceId.len)));
}

void FelgoLiveClient_ConnectToLocalServerDefault(void* ptr, struct QtFelgo_PackedString serverUrl, struct QtFelgo_PackedString userId, struct QtFelgo_PackedString deviceId)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::connectToLocalServer(QString::fromUtf8(serverUrl.data, serverUrl.len), QString::fromUtf8(userId.data, userId.len), QString::fromUtf8(deviceId.data, deviceId.len));
}

void FelgoLiveClient_ConnectToWebServer(void* ptr, struct QtFelgo_PackedString userId, struct QtFelgo_PackedString deviceId)
{
	static_cast<FelgoLiveClient*>(ptr)->connectToWebServer(QString::fromUtf8(userId.data, userId.len), QString::fromUtf8(deviceId.data, deviceId.len));
}

void FelgoLiveClient_LoadProject(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile)
{
	static_cast<FelgoLiveClient*>(ptr)->loadProject(QString::fromUtf8(projectName.data, projectName.len), QString::fromUtf8(projectMainFile.data, projectMainFile.len));
}

void FelgoLiveClient_DisableDisplaySleep(void* ptr)
{
	static_cast<FelgoLiveClient*>(ptr)->disableDisplaySleep();
}

void FelgoLiveClient_DisconnectWebReceiver(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "disconnectWebReceiver");
}

void FelgoLiveClient_DisconnectWebReceiverDefault(void* ptr)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::disconnectWebReceiver();
}

void FelgoLiveClient_ModifyLoadedDocument(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "modifyLoadedDocument");
}

void FelgoLiveClient_ModifyLoadedDocumentDefault(void* ptr)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::modifyLoadedDocument();
}

void FelgoLiveClient_OnDocumentUpdated(void* ptr, struct QtFelgo_PackedString document, struct QtFelgo_PackedString content)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "onDocumentUpdated", Q_ARG(const QString, QString::fromUtf8(document.data, document.len)), Q_ARG(const QString, QString::fromUtf8(content.data, content.len)));
}

void FelgoLiveClient_OnDocumentUpdatedDefault(void* ptr, struct QtFelgo_PackedString document, struct QtFelgo_PackedString content)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::onDocumentUpdated(QString::fromUtf8(document.data, document.len), QString::fromUtf8(content.data, content.len));
}

void FelgoLiveClient_OnProjectChanged(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "onProjectChanged", Q_ARG(const QString, QString::fromUtf8(projectName.data, projectName.len)), Q_ARG(const QString, QString::fromUtf8(projectMainFile.data, projectMainFile.len)));
}

void FelgoLiveClient_OnProjectChangedDefault(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::onProjectChanged(QString::fromUtf8(projectName.data, projectName.len), QString::fromUtf8(projectMainFile.data, projectMainFile.len));
}

void FelgoLiveClient_ConnectPendingProject(void* ptr)
{
	QObject::connect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &, const QString &)>(&FelgoLiveClient::pendingProject), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &, const QString &)>(&MyFelgoLiveClient::Signal_PendingProject));
}

void FelgoLiveClient_DisconnectPendingProject(void* ptr)
{
	QObject::disconnect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &, const QString &)>(&FelgoLiveClient::pendingProject), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &, const QString &)>(&MyFelgoLiveClient::Signal_PendingProject));
}

void FelgoLiveClient_PendingProject(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile)
{
	static_cast<FelgoLiveClient*>(ptr)->pendingProject(QString::fromUtf8(projectName.data, projectName.len), QString::fromUtf8(projectMainFile.data, projectMainFile.len));
}

void FelgoLiveClient_PrepareLoadingDocument(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "prepareLoadingDocument");
}

void FelgoLiveClient_PrepareLoadingDocumentDefault(void* ptr)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::prepareLoadingDocument();
}

void FelgoLiveClient_ConnectWebReceiverConnectionRefused(void* ptr)
{
	QObject::connect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &)>(&FelgoLiveClient::webReceiverConnectionRefused), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &)>(&MyFelgoLiveClient::Signal_WebReceiverConnectionRefused));
}

void FelgoLiveClient_DisconnectWebReceiverConnectionRefused(void* ptr)
{
	QObject::disconnect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &)>(&FelgoLiveClient::webReceiverConnectionRefused), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &)>(&MyFelgoLiveClient::Signal_WebReceiverConnectionRefused));
}

void FelgoLiveClient_WebReceiverConnectionRefused(void* ptr, struct QtFelgo_PackedString reason)
{
	static_cast<FelgoLiveClient*>(ptr)->webReceiverConnectionRefused(QString::fromUtf8(reason.data, reason.len));
}

void FelgoLiveClient_ConnectReceivedMatchId(void* ptr)
{
	QObject::connect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &)>(&FelgoLiveClient::receivedMatchId), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &)>(&MyFelgoLiveClient::Signal_ReceivedMatchId));
}

void FelgoLiveClient_DisconnectReceivedMatchId(void* ptr)
{
	QObject::disconnect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)(const QString &)>(&FelgoLiveClient::receivedMatchId), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)(const QString &)>(&MyFelgoLiveClient::Signal_ReceivedMatchId));
}

void FelgoLiveClient_ReceivedMatchId(void* ptr, struct QtFelgo_PackedString matchId)
{
	static_cast<FelgoLiveClient*>(ptr)->receivedMatchId(QString::fromUtf8(matchId.data, matchId.len));
}

void FelgoLiveClient_SetClientName(void* ptr, struct QtFelgo_PackedString clientName)
{
	static_cast<FelgoLiveClient*>(ptr)->setClientName(QString::fromUtf8(clientName.data, clientName.len));
}

void FelgoLiveClient_SetQmlEngine(void* ptr, void* engine)
{
	static_cast<FelgoLiveClient*>(ptr)->setQmlEngine(static_cast<QQmlEngine*>(engine));
}

void FelgoLiveClient_SetWelcomeQmlFile(void* ptr, struct QtFelgo_PackedString file)
{
	static_cast<FelgoLiveClient*>(ptr)->setWelcomeQmlFile(QString::fromUtf8(file.data, file.len));
}

void FelgoLiveClient_ShakeDetected(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<FelgoLiveClient*>(ptr), "shakeDetected");
}

void FelgoLiveClient_ShakeDetectedDefault(void* ptr)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::shakeDetected();
}

void FelgoLiveClient_ConnectWebReceiverConnected(void* ptr)
{
	QObject::connect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)()>(&FelgoLiveClient::webReceiverConnected), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)()>(&MyFelgoLiveClient::Signal_WebReceiverConnected));
}

void FelgoLiveClient_DisconnectWebReceiverConnected(void* ptr)
{
	QObject::disconnect(static_cast<FelgoLiveClient*>(ptr), static_cast<void (FelgoLiveClient::*)()>(&FelgoLiveClient::webReceiverConnected), static_cast<MyFelgoLiveClient*>(ptr), static_cast<void (MyFelgoLiveClient::*)()>(&MyFelgoLiveClient::Signal_WebReceiverConnected));
}

void FelgoLiveClient_WebReceiverConnected(void* ptr)
{
	static_cast<FelgoLiveClient*>(ptr)->webReceiverConnected();
}

void* FelgoLiveClient___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void FelgoLiveClient___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* FelgoLiveClient___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* FelgoLiveClient___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoLiveClient___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoLiveClient___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* FelgoLiveClient___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoLiveClient___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoLiveClient___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* FelgoLiveClient___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoLiveClient___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoLiveClient___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* FelgoLiveClient___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void FelgoLiveClient___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* FelgoLiveClient___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char FelgoLiveClient_EventDefault(void* ptr, void* e)
{
		return static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::event(static_cast<QEvent*>(e));
}

char FelgoLiveClient_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void FelgoLiveClient_ChildEventDefault(void* ptr, void* event)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::childEvent(static_cast<QChildEvent*>(event));
}

void FelgoLiveClient_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void FelgoLiveClient_CustomEventDefault(void* ptr, void* event)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::customEvent(static_cast<QEvent*>(event));
}

void FelgoLiveClient_DeleteLaterDefault(void* ptr)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::deleteLater();
}

void FelgoLiveClient_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void FelgoLiveClient_TimerEventDefault(void* ptr, void* event)
{
		static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::timerEvent(static_cast<QTimerEvent*>(event));
}

void* FelgoLiveClient_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<FelgoLiveClient*>(ptr)->FelgoLiveClient::metaObject());
}

