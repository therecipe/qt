// +build !minimal

#define protected public
#define private public

#include "webengine.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QAuthenticator>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QColor>
#include <QContextMenuEvent>
#include <QDBusPendingCallWatcher>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QIODevice>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QNetworkCookie>
#include <QObject>
#include <QOffscreenSurface>
#include <QPageLayout>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPdfWriter>
#include <QPoint>
#include <QPointF>
#include <QQuickItem>
#include <QQuickWebEngineProfile>
#include <QQuickWebEngineScript>
#include <QRadioData>
#include <QRect>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWebChannel>
#include <QWebEngineCertificateError>
#include <QWebEngineCookieStore>
#include <QWebEngineHttpRequest>
#include <QWebEnginePage>
#include <QWebEngineProfile>
#include <QWebEngineQuotaRequest>
#include <QWebEngineRegisterProtocolHandlerRequest>
#include <QWebEngineScript>
#include <QWebEngineScriptCollection>
#include <QWebEngineSettings>
#include <QWebEngineUrlRequestInfo>
#include <QWebEngineUrlRequestInterceptor>
#include <QWebEngineUrlRequestJob>
#include <QWebEngineUrlScheme>
#include <QWebEngineUrlSchemeHandler>
#include <QWebEngineView>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>
#include <QtWebEngine>

class MyQQuickWebEngineProfile: public QQuickWebEngineProfile
{
public:
	MyQQuickWebEngineProfile(QObject *parent = Q_NULLPTR) : QQuickWebEngineProfile(parent) {QQuickWebEngineProfile_QQuickWebEngineProfile_QRegisterMetaType();};
	void Signal_CachePathChanged() { callbackQQuickWebEngineProfile_CachePathChanged(this); };
	void Signal_HttpAcceptLanguageChanged() { callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged(this); };
	void Signal_HttpCacheMaximumSizeChanged() { callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged(this); };
	void Signal_HttpCacheTypeChanged() { callbackQQuickWebEngineProfile_HttpCacheTypeChanged(this); };
	void Signal_HttpUserAgentChanged() { callbackQQuickWebEngineProfile_HttpUserAgentChanged(this); };
	void Signal_OffTheRecordChanged() { callbackQQuickWebEngineProfile_OffTheRecordChanged(this); };
	void Signal_PersistentCookiesPolicyChanged() { callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged(this); };
	void Signal_PersistentStoragePathChanged() { callbackQQuickWebEngineProfile_PersistentStoragePathChanged(this); };
	void Signal_SpellCheckEnabledChanged() { callbackQQuickWebEngineProfile_SpellCheckEnabledChanged(this); };
	void Signal_SpellCheckLanguagesChanged() { callbackQQuickWebEngineProfile_SpellCheckLanguagesChanged(this); };
	void Signal_StorageNameChanged() { callbackQQuickWebEngineProfile_StorageNameChanged(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWebEngineProfile_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQQuickWebEngineProfile_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWebEngineProfile_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickWebEngineProfile_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineProfile_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWebEngineProfile_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWebEngineProfile_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWebEngineProfile_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineProfile_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickWebEngineProfile_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWebEngineProfile_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQuickWebEngineProfile*)

int QQuickWebEngineProfile_QQuickWebEngineProfile_QRegisterMetaType(){qRegisterMetaType<QQuickWebEngineProfile*>(); return qRegisterMetaType<MyQQuickWebEngineProfile*>();}

void* QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile()
{
	return QQuickWebEngineProfile::defaultProfile();
}

void* QQuickWebEngineProfile_NewQQuickWebEngineProfile(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QWindow*>(parent));
	} else {
		return new MyQQuickWebEngineProfile(static_cast<QObject*>(parent));
	}
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_QQuickWebEngineProfile_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t45943f = QQuickWebEngineProfile::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t45943f.prepend("WHITESPACE").constData()+10), t45943f.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_QQuickWebEngineProfile_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t4de620 = QQuickWebEngineProfile::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t4de620.prepend("WHITESPACE").constData()+10), t4de620.size()-10 }; });
}

void QQuickWebEngineProfile_ConnectCachePathChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::cachePathChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_CachePathChanged));
}

void QQuickWebEngineProfile_DisconnectCachePathChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::cachePathChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_CachePathChanged));
}

void QQuickWebEngineProfile_CachePathChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->cachePathChanged();
}

void QQuickWebEngineProfile_ClearHttpCache(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->clearHttpCache();
}

void QQuickWebEngineProfile_ConnectHttpAcceptLanguageChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpAcceptLanguageChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpAcceptLanguageChanged));
}

void QQuickWebEngineProfile_DisconnectHttpAcceptLanguageChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpAcceptLanguageChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpAcceptLanguageChanged));
}

void QQuickWebEngineProfile_HttpAcceptLanguageChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->httpAcceptLanguageChanged();
}

void QQuickWebEngineProfile_ConnectHttpCacheMaximumSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpCacheMaximumSizeChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpCacheMaximumSizeChanged));
}

void QQuickWebEngineProfile_DisconnectHttpCacheMaximumSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpCacheMaximumSizeChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpCacheMaximumSizeChanged));
}

void QQuickWebEngineProfile_HttpCacheMaximumSizeChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->httpCacheMaximumSizeChanged();
}

void QQuickWebEngineProfile_ConnectHttpCacheTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpCacheTypeChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpCacheTypeChanged));
}

void QQuickWebEngineProfile_DisconnectHttpCacheTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpCacheTypeChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpCacheTypeChanged));
}

void QQuickWebEngineProfile_HttpCacheTypeChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->httpCacheTypeChanged();
}

void QQuickWebEngineProfile_ConnectHttpUserAgentChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpUserAgentChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpUserAgentChanged));
}

void QQuickWebEngineProfile_DisconnectHttpUserAgentChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::httpUserAgentChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_HttpUserAgentChanged));
}

void QQuickWebEngineProfile_HttpUserAgentChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->httpUserAgentChanged();
}

void QQuickWebEngineProfile_InstallUrlSchemeHandler(void* ptr, void* scheme, void* handler)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->installUrlSchemeHandler(*static_cast<QByteArray*>(scheme), static_cast<QWebEngineUrlSchemeHandler*>(handler));
}

void QQuickWebEngineProfile_ConnectOffTheRecordChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::offTheRecordChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_OffTheRecordChanged));
}

void QQuickWebEngineProfile_DisconnectOffTheRecordChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::offTheRecordChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_OffTheRecordChanged));
}

void QQuickWebEngineProfile_OffTheRecordChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->offTheRecordChanged();
}

void QQuickWebEngineProfile_ConnectPersistentCookiesPolicyChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::persistentCookiesPolicyChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_PersistentCookiesPolicyChanged));
}

void QQuickWebEngineProfile_DisconnectPersistentCookiesPolicyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::persistentCookiesPolicyChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_PersistentCookiesPolicyChanged));
}

void QQuickWebEngineProfile_PersistentCookiesPolicyChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->persistentCookiesPolicyChanged();
}

void QQuickWebEngineProfile_ConnectPersistentStoragePathChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::persistentStoragePathChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_PersistentStoragePathChanged));
}

void QQuickWebEngineProfile_DisconnectPersistentStoragePathChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::persistentStoragePathChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_PersistentStoragePathChanged));
}

void QQuickWebEngineProfile_PersistentStoragePathChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->persistentStoragePathChanged();
}

void QQuickWebEngineProfile_RemoveAllUrlSchemeHandlers(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->removeAllUrlSchemeHandlers();
}

void QQuickWebEngineProfile_RemoveUrlScheme(void* ptr, void* scheme)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->removeUrlScheme(*static_cast<QByteArray*>(scheme));
}

void QQuickWebEngineProfile_RemoveUrlSchemeHandler(void* ptr, void* handler)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->removeUrlSchemeHandler(static_cast<QWebEngineUrlSchemeHandler*>(handler));
}

void QQuickWebEngineProfile_SetCachePath(void* ptr, struct QtWebEngine_PackedString path)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setCachePath(QString::fromUtf8(path.data, path.len));
}

void QQuickWebEngineProfile_SetHttpAcceptLanguage(void* ptr, struct QtWebEngine_PackedString httpAcceptLanguage)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setHttpAcceptLanguage(QString::fromUtf8(httpAcceptLanguage.data, httpAcceptLanguage.len));
}

void QQuickWebEngineProfile_SetHttpCacheMaximumSize(void* ptr, int maxSize)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setHttpCacheMaximumSize(maxSize);
}

void QQuickWebEngineProfile_SetHttpCacheType(void* ptr, long long vqq)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setHttpCacheType(static_cast<QQuickWebEngineProfile::HttpCacheType>(vqq));
}

void QQuickWebEngineProfile_SetHttpUserAgent(void* ptr, struct QtWebEngine_PackedString userAgent)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setHttpUserAgent(QString::fromUtf8(userAgent.data, userAgent.len));
}

void QQuickWebEngineProfile_SetOffTheRecord(void* ptr, char offTheRecord)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setOffTheRecord(offTheRecord != 0);
}

void QQuickWebEngineProfile_SetPersistentCookiesPolicy(void* ptr, long long vqq)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setPersistentCookiesPolicy(static_cast<QQuickWebEngineProfile::PersistentCookiesPolicy>(vqq));
}

void QQuickWebEngineProfile_SetPersistentStoragePath(void* ptr, struct QtWebEngine_PackedString path)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setPersistentStoragePath(QString::fromUtf8(path.data, path.len));
}

void QQuickWebEngineProfile_SetRequestInterceptor(void* ptr, void* interceptor)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setRequestInterceptor(static_cast<QWebEngineUrlRequestInterceptor*>(interceptor));
}

void QQuickWebEngineProfile_SetSpellCheckEnabled(void* ptr, char enabled)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setSpellCheckEnabled(enabled != 0);
}

void QQuickWebEngineProfile_SetSpellCheckLanguages(void* ptr, struct QtWebEngine_PackedString languages)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setSpellCheckLanguages(QString::fromUtf8(languages.data, languages.len).split("|", QString::SkipEmptyParts));
}

void QQuickWebEngineProfile_SetStorageName(void* ptr, struct QtWebEngine_PackedString name)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setStorageName(QString::fromUtf8(name.data, name.len));
}

void QQuickWebEngineProfile_ConnectSpellCheckEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::spellCheckEnabledChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_SpellCheckEnabledChanged));
}

void QQuickWebEngineProfile_DisconnectSpellCheckEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::spellCheckEnabledChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_SpellCheckEnabledChanged));
}

void QQuickWebEngineProfile_SpellCheckEnabledChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->spellCheckEnabledChanged();
}

void QQuickWebEngineProfile_ConnectSpellCheckLanguagesChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::spellCheckLanguagesChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_SpellCheckLanguagesChanged));
}

void QQuickWebEngineProfile_DisconnectSpellCheckLanguagesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::spellCheckLanguagesChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_SpellCheckLanguagesChanged));
}

void QQuickWebEngineProfile_SpellCheckLanguagesChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->spellCheckLanguagesChanged();
}

void QQuickWebEngineProfile_ConnectStorageNameChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::storageNameChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_StorageNameChanged));
}

void QQuickWebEngineProfile_DisconnectStorageNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::storageNameChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_StorageNameChanged));
}

void QQuickWebEngineProfile_StorageNameChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->storageNameChanged();
}

long long QQuickWebEngineProfile_HttpCacheType(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->httpCacheType();
}

long long QQuickWebEngineProfile_PersistentCookiesPolicy(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->persistentCookiesPolicy();
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_CachePath(void* ptr)
{
	return ({ QByteArray t8a236f = static_cast<QQuickWebEngineProfile*>(ptr)->cachePath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t8a236f.prepend("WHITESPACE").constData()+10), t8a236f.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_HttpAcceptLanguage(void* ptr)
{
	return ({ QByteArray taf6551 = static_cast<QQuickWebEngineProfile*>(ptr)->httpAcceptLanguage().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(taf6551.prepend("WHITESPACE").constData()+10), taf6551.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_HttpUserAgent(void* ptr)
{
	return ({ QByteArray t7ee5e9 = static_cast<QQuickWebEngineProfile*>(ptr)->httpUserAgent().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t7ee5e9.prepend("WHITESPACE").constData()+10), t7ee5e9.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_PersistentStoragePath(void* ptr)
{
	return ({ QByteArray t5eccf6 = static_cast<QQuickWebEngineProfile*>(ptr)->persistentStoragePath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t5eccf6.prepend("WHITESPACE").constData()+10), t5eccf6.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_StorageName(void* ptr)
{
	return ({ QByteArray t6af56e = static_cast<QQuickWebEngineProfile*>(ptr)->storageName().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t6af56e.prepend("WHITESPACE").constData()+10), t6af56e.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_SpellCheckLanguages(void* ptr)
{
	return ({ QByteArray td376a8 = static_cast<QQuickWebEngineProfile*>(ptr)->spellCheckLanguages().join("|").toUtf8(); QtWebEngine_PackedString { const_cast<char*>(td376a8.prepend("WHITESPACE").constData()+10), td376a8.size()-10 }; });
}

void* QQuickWebEngineProfile_CookieStore(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->cookieStore();
}

char QQuickWebEngineProfile_IsOffTheRecord(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->isOffTheRecord();
}

char QQuickWebEngineProfile_IsSpellCheckEnabled(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->isSpellCheckEnabled();
}

void* QQuickWebEngineProfile_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::metaObject());
}

void* QQuickWebEngineProfile_UrlSchemeHandler(void* ptr, void* scheme)
{
	return const_cast<QWebEngineUrlSchemeHandler*>(static_cast<QQuickWebEngineProfile*>(ptr)->urlSchemeHandler(*static_cast<QByteArray*>(scheme)));
}

int QQuickWebEngineProfile_HttpCacheMaximumSize(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->httpCacheMaximumSize();
}

void* QQuickWebEngineProfile___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQuickWebEngineProfile___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickWebEngineProfile___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQuickWebEngineProfile___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineProfile___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineProfile___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQuickWebEngineProfile___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineProfile___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineProfile___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQuickWebEngineProfile___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineProfile___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineProfile___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQuickWebEngineProfile___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineProfile___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineProfile___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QQuickWebEngineProfile_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::event(static_cast<QEvent*>(e));
}

char QQuickWebEngineProfile_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QQuickWebEngineProfile_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWebEngineProfile_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWebEngineProfile_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::customEvent(static_cast<QEvent*>(event));
}

void QQuickWebEngineProfile_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::deleteLater();
}

void QQuickWebEngineProfile_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWebEngineProfile_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQuickWebEngineScript: public QQuickWebEngineScript
{
public:
	MyQQuickWebEngineScript(QObject *parent = Q_NULLPTR) : QQuickWebEngineScript(parent) {QQuickWebEngineScript_QQuickWebEngineScript_QRegisterMetaType();};
	void Signal_InjectionPointChanged(QQuickWebEngineScript::InjectionPoint injectionPoint) { callbackQQuickWebEngineScript_InjectionPointChanged(this, injectionPoint); };
	void Signal_NameChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtWebEngine_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQQuickWebEngineScript_NameChanged(this, namePacked); };
	void Signal_RunOnSubframesChanged(bool on) { callbackQQuickWebEngineScript_RunOnSubframesChanged(this, on); };
	void Signal_SourceCodeChanged(const QString & code) { QByteArray te6fb06 = code.toUtf8(); QtWebEngine_PackedString codePacked = { const_cast<char*>(te6fb06.prepend("WHITESPACE").constData()+10), te6fb06.size()-10 };callbackQQuickWebEngineScript_SourceCodeChanged(this, codePacked); };
	void Signal_SourceUrlChanged(const QUrl & url) { callbackQQuickWebEngineScript_SourceUrlChanged(this, const_cast<QUrl*>(&url)); };
	void Signal_WorldIdChanged(QQuickWebEngineScript::ScriptWorldId scriptWorldId) { callbackQQuickWebEngineScript_WorldIdChanged(this, scriptWorldId); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWebEngineScript_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQQuickWebEngineScript_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWebEngineScript_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickWebEngineScript_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineScript_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWebEngineScript_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWebEngineScript_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWebEngineScript_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineScript_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickWebEngineScript_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWebEngineScript_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQuickWebEngineScript*)

int QQuickWebEngineScript_QQuickWebEngineScript_QRegisterMetaType(){qRegisterMetaType<QQuickWebEngineScript*>(); return qRegisterMetaType<MyQQuickWebEngineScript*>();}

void* QQuickWebEngineScript_NewQQuickWebEngineScript(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QWindow*>(parent));
	} else {
		return new MyQQuickWebEngineScript(static_cast<QObject*>(parent));
	}
}

struct QtWebEngine_PackedString QQuickWebEngineScript_QQuickWebEngineScript_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t19a788 = QQuickWebEngineScript::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t19a788.prepend("WHITESPACE").constData()+10), t19a788.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineScript_QQuickWebEngineScript_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t34c795 = QQuickWebEngineScript::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t34c795.prepend("WHITESPACE").constData()+10), t34c795.size()-10 }; });
}

void QQuickWebEngineScript_ConnectInjectionPointChanged(void* ptr)
{
	qRegisterMetaType<QQuickWebEngineScript::InjectionPoint>();
	QObject::connect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(QQuickWebEngineScript::InjectionPoint)>(&QQuickWebEngineScript::injectionPointChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(QQuickWebEngineScript::InjectionPoint)>(&MyQQuickWebEngineScript::Signal_InjectionPointChanged));
}

void QQuickWebEngineScript_DisconnectInjectionPointChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(QQuickWebEngineScript::InjectionPoint)>(&QQuickWebEngineScript::injectionPointChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(QQuickWebEngineScript::InjectionPoint)>(&MyQQuickWebEngineScript::Signal_InjectionPointChanged));
}

void QQuickWebEngineScript_InjectionPointChanged(void* ptr, long long injectionPoint)
{
	static_cast<QQuickWebEngineScript*>(ptr)->injectionPointChanged(static_cast<QQuickWebEngineScript::InjectionPoint>(injectionPoint));
}

void QQuickWebEngineScript_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(const QString &)>(&QQuickWebEngineScript::nameChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(const QString &)>(&MyQQuickWebEngineScript::Signal_NameChanged));
}

void QQuickWebEngineScript_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(const QString &)>(&QQuickWebEngineScript::nameChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(const QString &)>(&MyQQuickWebEngineScript::Signal_NameChanged));
}

void QQuickWebEngineScript_NameChanged(void* ptr, struct QtWebEngine_PackedString name)
{
	static_cast<QQuickWebEngineScript*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

void QQuickWebEngineScript_ConnectRunOnSubframesChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(bool)>(&QQuickWebEngineScript::runOnSubframesChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(bool)>(&MyQQuickWebEngineScript::Signal_RunOnSubframesChanged));
}

void QQuickWebEngineScript_DisconnectRunOnSubframesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(bool)>(&QQuickWebEngineScript::runOnSubframesChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(bool)>(&MyQQuickWebEngineScript::Signal_RunOnSubframesChanged));
}

void QQuickWebEngineScript_RunOnSubframesChanged(void* ptr, char on)
{
	static_cast<QQuickWebEngineScript*>(ptr)->runOnSubframesChanged(on != 0);
}

void QQuickWebEngineScript_SetInjectionPoint(void* ptr, long long injectionPoint)
{
	static_cast<QQuickWebEngineScript*>(ptr)->setInjectionPoint(static_cast<QQuickWebEngineScript::InjectionPoint>(injectionPoint));
}

void QQuickWebEngineScript_SetName(void* ptr, struct QtWebEngine_PackedString name)
{
	static_cast<QQuickWebEngineScript*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QQuickWebEngineScript_SetRunOnSubframes(void* ptr, char on)
{
	static_cast<QQuickWebEngineScript*>(ptr)->setRunOnSubframes(on != 0);
}

void QQuickWebEngineScript_SetSourceCode(void* ptr, struct QtWebEngine_PackedString code)
{
	static_cast<QQuickWebEngineScript*>(ptr)->setSourceCode(QString::fromUtf8(code.data, code.len));
}

void QQuickWebEngineScript_SetSourceUrl(void* ptr, void* url)
{
	static_cast<QQuickWebEngineScript*>(ptr)->setSourceUrl(*static_cast<QUrl*>(url));
}

void QQuickWebEngineScript_SetWorldId(void* ptr, long long scriptWorldId)
{
	static_cast<QQuickWebEngineScript*>(ptr)->setWorldId(static_cast<QQuickWebEngineScript::ScriptWorldId>(scriptWorldId));
}

void QQuickWebEngineScript_ConnectSourceCodeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(const QString &)>(&QQuickWebEngineScript::sourceCodeChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(const QString &)>(&MyQQuickWebEngineScript::Signal_SourceCodeChanged));
}

void QQuickWebEngineScript_DisconnectSourceCodeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(const QString &)>(&QQuickWebEngineScript::sourceCodeChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(const QString &)>(&MyQQuickWebEngineScript::Signal_SourceCodeChanged));
}

void QQuickWebEngineScript_SourceCodeChanged(void* ptr, struct QtWebEngine_PackedString code)
{
	static_cast<QQuickWebEngineScript*>(ptr)->sourceCodeChanged(QString::fromUtf8(code.data, code.len));
}

void QQuickWebEngineScript_ConnectSourceUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(const QUrl &)>(&QQuickWebEngineScript::sourceUrlChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(const QUrl &)>(&MyQQuickWebEngineScript::Signal_SourceUrlChanged));
}

void QQuickWebEngineScript_DisconnectSourceUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(const QUrl &)>(&QQuickWebEngineScript::sourceUrlChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(const QUrl &)>(&MyQQuickWebEngineScript::Signal_SourceUrlChanged));
}

void QQuickWebEngineScript_SourceUrlChanged(void* ptr, void* url)
{
	static_cast<QQuickWebEngineScript*>(ptr)->sourceUrlChanged(*static_cast<QUrl*>(url));
}

void QQuickWebEngineScript_ConnectWorldIdChanged(void* ptr)
{
	qRegisterMetaType<QQuickWebEngineScript::ScriptWorldId>();
	QObject::connect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(QQuickWebEngineScript::ScriptWorldId)>(&QQuickWebEngineScript::worldIdChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(QQuickWebEngineScript::ScriptWorldId)>(&MyQQuickWebEngineScript::Signal_WorldIdChanged));
}

void QQuickWebEngineScript_DisconnectWorldIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineScript*>(ptr), static_cast<void (QQuickWebEngineScript::*)(QQuickWebEngineScript::ScriptWorldId)>(&QQuickWebEngineScript::worldIdChanged), static_cast<MyQQuickWebEngineScript*>(ptr), static_cast<void (MyQQuickWebEngineScript::*)(QQuickWebEngineScript::ScriptWorldId)>(&MyQQuickWebEngineScript::Signal_WorldIdChanged));
}

void QQuickWebEngineScript_WorldIdChanged(void* ptr, long long scriptWorldId)
{
	static_cast<QQuickWebEngineScript*>(ptr)->worldIdChanged(static_cast<QQuickWebEngineScript::ScriptWorldId>(scriptWorldId));
}

long long QQuickWebEngineScript_InjectionPoint(void* ptr)
{
	return static_cast<QQuickWebEngineScript*>(ptr)->injectionPoint();
}

long long QQuickWebEngineScript_WorldId(void* ptr)
{
	return static_cast<QQuickWebEngineScript*>(ptr)->worldId();
}

struct QtWebEngine_PackedString QQuickWebEngineScript_Name(void* ptr)
{
	return ({ QByteArray te34dea = static_cast<QQuickWebEngineScript*>(ptr)->name().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(te34dea.prepend("WHITESPACE").constData()+10), te34dea.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineScript_SourceCode(void* ptr)
{
	return ({ QByteArray ta033c9 = static_cast<QQuickWebEngineScript*>(ptr)->sourceCode().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(ta033c9.prepend("WHITESPACE").constData()+10), ta033c9.size()-10 }; });
}

struct QtWebEngine_PackedString QQuickWebEngineScript_ToString(void* ptr)
{
	return ({ QByteArray t0f4281 = static_cast<QQuickWebEngineScript*>(ptr)->toString().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t0f4281.prepend("WHITESPACE").constData()+10), t0f4281.size()-10 }; });
}

void* QQuickWebEngineScript_SourceUrl(void* ptr)
{
	return new QUrl(static_cast<QQuickWebEngineScript*>(ptr)->sourceUrl());
}

char QQuickWebEngineScript_RunOnSubframes(void* ptr)
{
	return static_cast<QQuickWebEngineScript*>(ptr)->runOnSubframes();
}

void* QQuickWebEngineScript_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::metaObject());
}

void* QQuickWebEngineScript___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQuickWebEngineScript___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickWebEngineScript___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQuickWebEngineScript___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineScript___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineScript___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQuickWebEngineScript___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineScript___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineScript___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQuickWebEngineScript___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineScript___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineScript___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQuickWebEngineScript___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineScript___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineScript___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QQuickWebEngineScript_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::event(static_cast<QEvent*>(e));
}

char QQuickWebEngineScript_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QQuickWebEngineScript_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWebEngineScript_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWebEngineScript_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::customEvent(static_cast<QEvent*>(event));
}

void QQuickWebEngineScript_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::deleteLater();
}

void QQuickWebEngineScript_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWebEngineScript_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::timerEvent(static_cast<QTimerEvent*>(event));
}

long long QWebEngineCertificateError_Error(void* ptr)
{
	return static_cast<QWebEngineCertificateError*>(ptr)->error();
}

struct QtWebEngine_PackedString QWebEngineCertificateError_ErrorDescription(void* ptr)
{
	return ({ QByteArray t3bb4ec = static_cast<QWebEngineCertificateError*>(ptr)->errorDescription().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t3bb4ec.prepend("WHITESPACE").constData()+10), t3bb4ec.size()-10 }; });
}

char QWebEngineCertificateError_IsOverridable(void* ptr)
{
	return static_cast<QWebEngineCertificateError*>(ptr)->isOverridable();
}

void* QWebEngineCertificateError_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEngineCertificateError*>(ptr)->url());
}

class MyQWebEngineCookieStore: public QWebEngineCookieStore
{
public:
	void Signal_CookieAdded(const QNetworkCookie & cookie) { callbackQWebEngineCookieStore_CookieAdded(this, const_cast<QNetworkCookie*>(&cookie)); };
	void Signal_CookieRemoved(const QNetworkCookie & cookie) { callbackQWebEngineCookieStore_CookieRemoved(this, const_cast<QNetworkCookie*>(&cookie)); };
	 ~MyQWebEngineCookieStore() { callbackQWebEngineCookieStore_DestroyQWebEngineCookieStore(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEngineCookieStore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebEngineCookieStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineCookieStore_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEngineCookieStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineCookieStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineCookieStore_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineCookieStore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineCookieStore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineCookieStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineCookieStore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineCookieStore_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineCookieStore*)

int QWebEngineCookieStore_QWebEngineCookieStore_QRegisterMetaType(){qRegisterMetaType<QWebEngineCookieStore*>(); return qRegisterMetaType<MyQWebEngineCookieStore*>();}

struct QtWebEngine_PackedString QWebEngineCookieStore_QWebEngineCookieStore_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tb12676 = QWebEngineCookieStore::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tb12676.prepend("WHITESPACE").constData()+10), tb12676.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineCookieStore_QWebEngineCookieStore_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t9ed19c = QWebEngineCookieStore::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t9ed19c.prepend("WHITESPACE").constData()+10), t9ed19c.size()-10 }; });
}

void QWebEngineCookieStore_ConnectCookieAdded(void* ptr)
{
	QObject::connect(static_cast<QWebEngineCookieStore*>(ptr), static_cast<void (QWebEngineCookieStore::*)(const QNetworkCookie &)>(&QWebEngineCookieStore::cookieAdded), static_cast<MyQWebEngineCookieStore*>(ptr), static_cast<void (MyQWebEngineCookieStore::*)(const QNetworkCookie &)>(&MyQWebEngineCookieStore::Signal_CookieAdded));
}

void QWebEngineCookieStore_DisconnectCookieAdded(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineCookieStore*>(ptr), static_cast<void (QWebEngineCookieStore::*)(const QNetworkCookie &)>(&QWebEngineCookieStore::cookieAdded), static_cast<MyQWebEngineCookieStore*>(ptr), static_cast<void (MyQWebEngineCookieStore::*)(const QNetworkCookie &)>(&MyQWebEngineCookieStore::Signal_CookieAdded));
}

void QWebEngineCookieStore_CookieAdded(void* ptr, void* cookie)
{
	static_cast<QWebEngineCookieStore*>(ptr)->cookieAdded(*static_cast<QNetworkCookie*>(cookie));
}

void QWebEngineCookieStore_ConnectCookieRemoved(void* ptr)
{
	QObject::connect(static_cast<QWebEngineCookieStore*>(ptr), static_cast<void (QWebEngineCookieStore::*)(const QNetworkCookie &)>(&QWebEngineCookieStore::cookieRemoved), static_cast<MyQWebEngineCookieStore*>(ptr), static_cast<void (MyQWebEngineCookieStore::*)(const QNetworkCookie &)>(&MyQWebEngineCookieStore::Signal_CookieRemoved));
}

void QWebEngineCookieStore_DisconnectCookieRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineCookieStore*>(ptr), static_cast<void (QWebEngineCookieStore::*)(const QNetworkCookie &)>(&QWebEngineCookieStore::cookieRemoved), static_cast<MyQWebEngineCookieStore*>(ptr), static_cast<void (MyQWebEngineCookieStore::*)(const QNetworkCookie &)>(&MyQWebEngineCookieStore::Signal_CookieRemoved));
}

void QWebEngineCookieStore_CookieRemoved(void* ptr, void* cookie)
{
	static_cast<QWebEngineCookieStore*>(ptr)->cookieRemoved(*static_cast<QNetworkCookie*>(cookie));
}

void QWebEngineCookieStore_DeleteAllCookies(void* ptr)
{
	static_cast<QWebEngineCookieStore*>(ptr)->deleteAllCookies();
}

void QWebEngineCookieStore_DeleteCookie(void* ptr, void* cookie, void* origin)
{
	static_cast<QWebEngineCookieStore*>(ptr)->deleteCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(origin));
}

void QWebEngineCookieStore_DeleteSessionCookies(void* ptr)
{
	static_cast<QWebEngineCookieStore*>(ptr)->deleteSessionCookies();
}

void QWebEngineCookieStore_LoadAllCookies(void* ptr)
{
	static_cast<QWebEngineCookieStore*>(ptr)->loadAllCookies();
}

void QWebEngineCookieStore_SetCookie(void* ptr, void* cookie, void* origin)
{
	static_cast<QWebEngineCookieStore*>(ptr)->setCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(origin));
}

void QWebEngineCookieStore_DestroyQWebEngineCookieStore(void* ptr)
{
	static_cast<QWebEngineCookieStore*>(ptr)->~QWebEngineCookieStore();
}

void QWebEngineCookieStore_DestroyQWebEngineCookieStoreDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebEngineCookieStore_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::metaObject());
}

void* QWebEngineCookieStore___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineCookieStore___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineCookieStore___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineCookieStore___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineCookieStore___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineCookieStore___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineCookieStore___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineCookieStore___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineCookieStore___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineCookieStore___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineCookieStore___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineCookieStore___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineCookieStore___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineCookieStore___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineCookieStore___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEngineCookieStore_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::event(static_cast<QEvent*>(e));
}

char QWebEngineCookieStore_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineCookieStore_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineCookieStore_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineCookieStore_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineCookieStore_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::deleteLater();
}

void QWebEngineCookieStore_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineCookieStore_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebEngineHttpRequest_QWebEngineHttpRequest_PostRequest(void* url, void* postData)
{
	return new QWebEngineHttpRequest(QWebEngineHttpRequest::postRequest(*static_cast<QUrl*>(url), *static_cast<QMap<QString, QString>*>(postData)));
}

void* QWebEngineHttpRequest_NewQWebEngineHttpRequest(void* url, long long method)
{
	return new QWebEngineHttpRequest(*static_cast<QUrl*>(url), static_cast<QWebEngineHttpRequest::Method>(method));
}

void* QWebEngineHttpRequest_NewQWebEngineHttpRequest2(void* other)
{
	return new QWebEngineHttpRequest(*static_cast<QWebEngineHttpRequest*>(other));
}

void QWebEngineHttpRequest_SetHeader(void* ptr, void* headerName, void* headerValue)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->setHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QWebEngineHttpRequest_SetMethod(void* ptr, long long method)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->setMethod(static_cast<QWebEngineHttpRequest::Method>(method));
}

void QWebEngineHttpRequest_SetPostData(void* ptr, void* postData)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->setPostData(*static_cast<QByteArray*>(postData));
}

void QWebEngineHttpRequest_SetUrl(void* ptr, void* url)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebEngineHttpRequest_Swap(void* ptr, void* other)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->swap(*static_cast<QWebEngineHttpRequest*>(other));
}

void QWebEngineHttpRequest_UnsetHeader(void* ptr, void* key)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->unsetHeader(*static_cast<QByteArray*>(key));
}

void QWebEngineHttpRequest_DestroyQWebEngineHttpRequest(void* ptr)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->~QWebEngineHttpRequest();
}

void* QWebEngineHttpRequest_Header(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QWebEngineHttpRequest*>(ptr)->header(*static_cast<QByteArray*>(headerName)));
}

void* QWebEngineHttpRequest_PostData(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineHttpRequest*>(ptr)->postData());
}

void* QWebEngineHttpRequest_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEngineHttpRequest*>(ptr)->url());
}

struct QtWebEngine_PackedList QWebEngineHttpRequest_Headers(void* ptr)
{
	return ({ QVector<QByteArray>* tmpValue = new QVector<QByteArray>(static_cast<QWebEngineHttpRequest*>(ptr)->headers()); QtWebEngine_PackedList { tmpValue, tmpValue->size() }; });
}

long long QWebEngineHttpRequest_Method(void* ptr)
{
	return static_cast<QWebEngineHttpRequest*>(ptr)->method();
}

char QWebEngineHttpRequest_HasHeader(void* ptr, void* headerName)
{
	return static_cast<QWebEngineHttpRequest*>(ptr)->hasHeader(*static_cast<QByteArray*>(headerName));
}

struct QtWebEngine_PackedString QWebEngineHttpRequest___postRequest_postData_atList(void* ptr, struct QtWebEngine_PackedString v, int i)
{
	return ({ QByteArray t1a0c54 = ({ QString tmp = static_cast<QMap<QString, QString>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QString>*>(ptr)->size()-1) { static_cast<QMap<QString, QString>*>(ptr)->~QMap(); free(ptr); }; tmp; }).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t1a0c54.prepend("WHITESPACE").constData()+10), t1a0c54.size()-10 }; });
}

void QWebEngineHttpRequest___postRequest_postData_setList(void* ptr, struct QtWebEngine_PackedString key, struct QtWebEngine_PackedString i)
{
	static_cast<QMap<QString, QString>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), QString::fromUtf8(i.data, i.len));
}

void* QWebEngineHttpRequest___postRequest_postData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QString>();
}

struct QtWebEngine_PackedList QWebEngineHttpRequest___postRequest_postData_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QString>*>(ptr)->keys()); QtWebEngine_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebEngineHttpRequest___headers_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QVector<QByteArray>*>(ptr)->at(i); if (i == static_cast<QVector<QByteArray>*>(ptr)->size()-1) { static_cast<QVector<QByteArray>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QWebEngineHttpRequest___headers_setList(void* ptr, void* i)
{
	static_cast<QVector<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineHttpRequest___headers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QByteArray>();
}

struct QtWebEngine_PackedString QWebEngineHttpRequest_____postRequest_postData_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QWebEngineHttpRequest_____postRequest_postData_keyList_setList(void* ptr, struct QtWebEngine_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebEngineHttpRequest_____postRequest_postData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

class MyQWebEnginePage: public QWebEnginePage
{
public:
	MyQWebEnginePage(QObject *parent = Q_NULLPTR) : QWebEnginePage(parent) {QWebEnginePage_QWebEnginePage_QRegisterMetaType();};
	MyQWebEnginePage(QWebEngineProfile *profile, QObject *parent = Q_NULLPTR) : QWebEnginePage(profile, parent) {QWebEnginePage_QWebEnginePage_QRegisterMetaType();};
	QStringList chooseFiles(QWebEnginePage::FileSelectionMode mode, const QStringList & oldFiles, const QStringList & acceptedMimeTypes) { QByteArray t76015f = oldFiles.join("|").toUtf8(); QtWebEngine_PackedString oldFilesPacked = { const_cast<char*>(t76015f.prepend("WHITESPACE").constData()+10), t76015f.size()-10 };QByteArray t541092 = acceptedMimeTypes.join("|").toUtf8(); QtWebEngine_PackedString acceptedMimeTypesPacked = { const_cast<char*>(t541092.prepend("WHITESPACE").constData()+10), t541092.size()-10 };return ({ QtWebEngine_PackedString tempVal = callbackQWebEnginePage_ChooseFiles(this, mode, oldFilesPacked, acceptedMimeTypesPacked); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QWebEnginePage * createWindow(QWebEnginePage::WebWindowType ty) { return static_cast<QWebEnginePage*>(callbackQWebEnginePage_CreateWindow(this, ty)); };
	bool acceptNavigationRequest(const QUrl & url, QWebEnginePage::NavigationType ty, bool isMainFrame) { return callbackQWebEnginePage_AcceptNavigationRequest(this, const_cast<QUrl*>(&url), ty, isMainFrame) != 0; };
	bool certificateError(const QWebEngineCertificateError & certificateError) { return callbackQWebEnginePage_CertificateError(this, const_cast<QWebEngineCertificateError*>(&certificateError)) != 0; };
	bool event(QEvent * vqe) { return callbackQWebEnginePage_Event(this, vqe) != 0; };
	bool javaScriptConfirm(const QUrl & securityOrigin, const QString & msg) { QByteArray t19f34e = msg.toUtf8(); QtWebEngine_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };return callbackQWebEnginePage_JavaScriptConfirm(this, const_cast<QUrl*>(&securityOrigin), msgPacked) != 0; };
	bool javaScriptPrompt(const QUrl & securityOrigin, const QString & msg, const QString & defaultValue, QString * result) { QByteArray t19f34e = msg.toUtf8(); QtWebEngine_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };QByteArray te940d2 = defaultValue.toUtf8(); QtWebEngine_PackedString defaultValuePacked = { const_cast<char*>(te940d2.prepend("WHITESPACE").constData()+10), te940d2.size()-10 };QByteArray t37a530 = result->toUtf8(); QtWebEngine_PackedString resultPacked = { const_cast<char*>(t37a530.prepend("WHITESPACE").constData()+10), t37a530.size()-10 };return callbackQWebEnginePage_JavaScriptPrompt(this, const_cast<QUrl*>(&securityOrigin), msgPacked, defaultValuePacked, resultPacked) != 0; };
	void Signal_AudioMutedChanged(bool muted) { callbackQWebEnginePage_AudioMutedChanged(this, muted); };
	void Signal_AuthenticationRequired(const QUrl & requestUrl, QAuthenticator * authenticator) { callbackQWebEnginePage_AuthenticationRequired(this, const_cast<QUrl*>(&requestUrl), authenticator); };
	void Signal_ContentsSizeChanged(const QSizeF & size) { callbackQWebEnginePage_ContentsSizeChanged(this, const_cast<QSizeF*>(&size)); };
	void Signal_FeaturePermissionRequestCanceled(const QUrl & securityOrigin, QWebEnginePage::Feature feature) { callbackQWebEnginePage_FeaturePermissionRequestCanceled(this, const_cast<QUrl*>(&securityOrigin), feature); };
	void Signal_FeaturePermissionRequested(const QUrl & securityOrigin, QWebEnginePage::Feature feature) { callbackQWebEnginePage_FeaturePermissionRequested(this, const_cast<QUrl*>(&securityOrigin), feature); };
	void Signal_GeometryChangeRequested(const QRect & geom) { callbackQWebEnginePage_GeometryChangeRequested(this, const_cast<QRect*>(&geom)); };
	void Signal_IconChanged(const QIcon & icon) { callbackQWebEnginePage_IconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_IconUrlChanged(const QUrl & url) { callbackQWebEnginePage_IconUrlChanged(this, const_cast<QUrl*>(&url)); };
	void javaScriptAlert(const QUrl & securityOrigin, const QString & msg) { QByteArray t19f34e = msg.toUtf8(); QtWebEngine_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQWebEnginePage_JavaScriptAlert(this, const_cast<QUrl*>(&securityOrigin), msgPacked); };
	void javaScriptConsoleMessage(QWebEnginePage::JavaScriptConsoleMessageLevel level, const QString & message, int lineNumber, const QString & sourceID) { QByteArray t6f9b9a = message.toUtf8(); QtWebEngine_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };QByteArray tf767e3 = sourceID.toUtf8(); QtWebEngine_PackedString sourceIDPacked = { const_cast<char*>(tf767e3.prepend("WHITESPACE").constData()+10), tf767e3.size()-10 };callbackQWebEnginePage_JavaScriptConsoleMessage(this, level, messagePacked, lineNumber, sourceIDPacked); };
	void Signal_LinkHovered(const QString & url) { QByteArray t817363 = url.toUtf8(); QtWebEngine_PackedString urlPacked = { const_cast<char*>(t817363.prepend("WHITESPACE").constData()+10), t817363.size()-10 };callbackQWebEnginePage_LinkHovered(this, urlPacked); };
	void Signal_LoadFinished(bool ok) { callbackQWebEnginePage_LoadFinished(this, ok); };
	void Signal_LoadProgress(int progress) { callbackQWebEnginePage_LoadProgress(this, progress); };
	void Signal_LoadStarted() { callbackQWebEnginePage_LoadStarted(this); };
	void Signal_PdfPrintingFinished(const QString & filePath, bool success) { QByteArray t7df503 = filePath.toUtf8(); QtWebEngine_PackedString filePathPacked = { const_cast<char*>(t7df503.prepend("WHITESPACE").constData()+10), t7df503.size()-10 };callbackQWebEnginePage_PdfPrintingFinished(this, filePathPacked, success); };
	void Signal_PrintRequested() { callbackQWebEnginePage_PrintRequested(this); };
	void Signal_ProxyAuthenticationRequired(const QUrl & requestUrl, QAuthenticator * authenticator, const QString & proxyHost) { QByteArray teddfac = proxyHost.toUtf8(); QtWebEngine_PackedString proxyHostPacked = { const_cast<char*>(teddfac.prepend("WHITESPACE").constData()+10), teddfac.size()-10 };callbackQWebEnginePage_ProxyAuthenticationRequired(this, const_cast<QUrl*>(&requestUrl), authenticator, proxyHostPacked); };
	void Signal_RecentlyAudibleChanged(bool recentlyAudible) { callbackQWebEnginePage_RecentlyAudibleChanged(this, recentlyAudible); };
	void Signal_RenderProcessTerminated(QWebEnginePage::RenderProcessTerminationStatus terminationStatus, int exitCode) { callbackQWebEnginePage_RenderProcessTerminated(this, terminationStatus, exitCode); };
	void Signal_ScrollPositionChanged(const QPointF & position) { callbackQWebEnginePage_ScrollPositionChanged(this, const_cast<QPointF*>(&position)); };
	void Signal_SelectionChanged() { callbackQWebEnginePage_SelectionChanged(this); };
	void Signal_TitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebEngine_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebEnginePage_TitleChanged(this, titlePacked); };
	void triggerAction(QWebEnginePage::WebAction action, bool checked) { callbackQWebEnginePage_TriggerAction(this, action, checked); };
	void Signal_UrlChanged(const QUrl & url) { callbackQWebEnginePage_UrlChanged(this, const_cast<QUrl*>(&url)); };
	void Signal_WindowCloseRequested() { callbackQWebEnginePage_WindowCloseRequested(this); };
	 ~MyQWebEnginePage() { callbackQWebEnginePage_DestroyQWebEnginePage(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEnginePage_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEnginePage_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEnginePage_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEnginePage_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEnginePage_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEnginePage_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEnginePage_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEnginePage_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEnginePage_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEnginePage_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEnginePage*)

int QWebEnginePage_QWebEnginePage_QRegisterMetaType(){qRegisterMetaType<QWebEnginePage*>(); return qRegisterMetaType<MyQWebEnginePage*>();}

void* QWebEnginePage_CreateStandardContextMenu(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->createStandardContextMenu();
}

struct QtWebEngine_PackedString QWebEnginePage_QWebEnginePage_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t3acb8e = QWebEnginePage::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t3acb8e.prepend("WHITESPACE").constData()+10), t3acb8e.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEnginePage_QWebEnginePage_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray ta11142 = QWebEnginePage::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(ta11142.prepend("WHITESPACE").constData()+10), ta11142.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEnginePage_ChooseFiles(void* ptr, long long mode, struct QtWebEngine_PackedString oldFiles, struct QtWebEngine_PackedString acceptedMimeTypes)
{
	return ({ QByteArray t5f4ad5 = static_cast<QWebEnginePage*>(ptr)->chooseFiles(static_cast<QWebEnginePage::FileSelectionMode>(mode), QString::fromUtf8(oldFiles.data, oldFiles.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(acceptedMimeTypes.data, acceptedMimeTypes.len).split("|", QString::SkipEmptyParts)).join("|").toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t5f4ad5.prepend("WHITESPACE").constData()+10), t5f4ad5.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEnginePage_ChooseFilesDefault(void* ptr, long long mode, struct QtWebEngine_PackedString oldFiles, struct QtWebEngine_PackedString acceptedMimeTypes)
{
		return ({ QByteArray t194716 = static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::chooseFiles(static_cast<QWebEnginePage::FileSelectionMode>(mode), QString::fromUtf8(oldFiles.data, oldFiles.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(acceptedMimeTypes.data, acceptedMimeTypes.len).split("|", QString::SkipEmptyParts)).join("|").toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t194716.prepend("WHITESPACE").constData()+10), t194716.size()-10 }; });
}

void* QWebEnginePage_CreateWindow(void* ptr, long long ty)
{
	return static_cast<QWebEnginePage*>(ptr)->createWindow(static_cast<QWebEnginePage::WebWindowType>(ty));
}

void* QWebEnginePage_CreateWindowDefault(void* ptr, long long ty)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::createWindow(static_cast<QWebEnginePage::WebWindowType>(ty));
}

void* QWebEnginePage_NewQWebEnginePage(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebEnginePage(static_cast<QObject*>(parent));
	}
}

void* QWebEnginePage_NewQWebEnginePage2(void* profile, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QWindow*>(parent));
	} else {
		return new MyQWebEnginePage(static_cast<QWebEngineProfile*>(profile), static_cast<QObject*>(parent));
	}
}

char QWebEnginePage_AcceptNavigationRequest(void* ptr, void* url, long long ty, char isMainFrame)
{
	return static_cast<QWebEnginePage*>(ptr)->acceptNavigationRequest(*static_cast<QUrl*>(url), static_cast<QWebEnginePage::NavigationType>(ty), isMainFrame != 0);
}

char QWebEnginePage_AcceptNavigationRequestDefault(void* ptr, void* url, long long ty, char isMainFrame)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::acceptNavigationRequest(*static_cast<QUrl*>(url), static_cast<QWebEnginePage::NavigationType>(ty), isMainFrame != 0);
}

char QWebEnginePage_CertificateError(void* ptr, void* certificateError)
{
	return static_cast<QWebEnginePage*>(ptr)->certificateError(*static_cast<QWebEngineCertificateError*>(certificateError));
}

char QWebEnginePage_CertificateErrorDefault(void* ptr, void* certificateError)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::certificateError(*static_cast<QWebEngineCertificateError*>(certificateError));
}

char QWebEnginePage_EventDefault(void* ptr, void* vqe)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::event(static_cast<QEvent*>(vqe));
}

char QWebEnginePage_JavaScriptConfirm(void* ptr, void* securityOrigin, struct QtWebEngine_PackedString msg)
{
	return static_cast<QWebEnginePage*>(ptr)->javaScriptConfirm(*static_cast<QUrl*>(securityOrigin), QString::fromUtf8(msg.data, msg.len));
}

char QWebEnginePage_JavaScriptConfirmDefault(void* ptr, void* securityOrigin, struct QtWebEngine_PackedString msg)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::javaScriptConfirm(*static_cast<QUrl*>(securityOrigin), QString::fromUtf8(msg.data, msg.len));
}

char QWebEnginePage_JavaScriptPrompt(void* ptr, void* securityOrigin, struct QtWebEngine_PackedString msg, struct QtWebEngine_PackedString defaultValue, struct QtWebEngine_PackedString result)
{
	return static_cast<QWebEnginePage*>(ptr)->javaScriptPrompt(*static_cast<QUrl*>(securityOrigin), QString::fromUtf8(msg.data, msg.len), QString::fromUtf8(defaultValue.data, defaultValue.len), new QString(QString::fromUtf8(result.data, result.len)));
}

char QWebEnginePage_JavaScriptPromptDefault(void* ptr, void* securityOrigin, struct QtWebEngine_PackedString msg, struct QtWebEngine_PackedString defaultValue, struct QtWebEngine_PackedString result)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::javaScriptPrompt(*static_cast<QUrl*>(securityOrigin), QString::fromUtf8(msg.data, msg.len), QString::fromUtf8(defaultValue.data, defaultValue.len), new QString(QString::fromUtf8(result.data, result.len)));
}

void QWebEnginePage_ConnectAudioMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(bool)>(&QWebEnginePage::audioMutedChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(bool)>(&MyQWebEnginePage::Signal_AudioMutedChanged));
}

void QWebEnginePage_DisconnectAudioMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(bool)>(&QWebEnginePage::audioMutedChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(bool)>(&MyQWebEnginePage::Signal_AudioMutedChanged));
}

void QWebEnginePage_AudioMutedChanged(void* ptr, char muted)
{
	static_cast<QWebEnginePage*>(ptr)->audioMutedChanged(muted != 0);
}

void QWebEnginePage_ConnectAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QAuthenticator *)>(&QWebEnginePage::authenticationRequired), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QAuthenticator *)>(&MyQWebEnginePage::Signal_AuthenticationRequired));
}

void QWebEnginePage_DisconnectAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QAuthenticator *)>(&QWebEnginePage::authenticationRequired), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QAuthenticator *)>(&MyQWebEnginePage::Signal_AuthenticationRequired));
}

void QWebEnginePage_AuthenticationRequired(void* ptr, void* requestUrl, void* authenticator)
{
	static_cast<QWebEnginePage*>(ptr)->authenticationRequired(*static_cast<QUrl*>(requestUrl), static_cast<QAuthenticator*>(authenticator));
}

void QWebEnginePage_ConnectContentsSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QSizeF &)>(&QWebEnginePage::contentsSizeChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QSizeF &)>(&MyQWebEnginePage::Signal_ContentsSizeChanged));
}

void QWebEnginePage_DisconnectContentsSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QSizeF &)>(&QWebEnginePage::contentsSizeChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QSizeF &)>(&MyQWebEnginePage::Signal_ContentsSizeChanged));
}

void QWebEnginePage_ContentsSizeChanged(void* ptr, void* size)
{
	static_cast<QWebEnginePage*>(ptr)->contentsSizeChanged(*static_cast<QSizeF*>(size));
}

void QWebEnginePage_Download(void* ptr, void* url, struct QtWebEngine_PackedString filename)
{
	static_cast<QWebEnginePage*>(ptr)->download(*static_cast<QUrl*>(url), QString::fromUtf8(filename.data, filename.len));
}

void QWebEnginePage_ConnectFeaturePermissionRequestCanceled(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&QWebEnginePage::featurePermissionRequestCanceled), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&MyQWebEnginePage::Signal_FeaturePermissionRequestCanceled));
}

void QWebEnginePage_DisconnectFeaturePermissionRequestCanceled(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&QWebEnginePage::featurePermissionRequestCanceled), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&MyQWebEnginePage::Signal_FeaturePermissionRequestCanceled));
}

void QWebEnginePage_FeaturePermissionRequestCanceled(void* ptr, void* securityOrigin, long long feature)
{
	static_cast<QWebEnginePage*>(ptr)->featurePermissionRequestCanceled(*static_cast<QUrl*>(securityOrigin), static_cast<QWebEnginePage::Feature>(feature));
}

void QWebEnginePage_ConnectFeaturePermissionRequested(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&QWebEnginePage::featurePermissionRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&MyQWebEnginePage::Signal_FeaturePermissionRequested));
}

void QWebEnginePage_DisconnectFeaturePermissionRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&QWebEnginePage::featurePermissionRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QWebEnginePage::Feature)>(&MyQWebEnginePage::Signal_FeaturePermissionRequested));
}

void QWebEnginePage_FeaturePermissionRequested(void* ptr, void* securityOrigin, long long feature)
{
	static_cast<QWebEnginePage*>(ptr)->featurePermissionRequested(*static_cast<QUrl*>(securityOrigin), static_cast<QWebEnginePage::Feature>(feature));
}

void QWebEnginePage_ConnectGeometryChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QRect &)>(&QWebEnginePage::geometryChangeRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QRect &)>(&MyQWebEnginePage::Signal_GeometryChangeRequested));
}

void QWebEnginePage_DisconnectGeometryChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QRect &)>(&QWebEnginePage::geometryChangeRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QRect &)>(&MyQWebEnginePage::Signal_GeometryChangeRequested));
}

void QWebEnginePage_GeometryChangeRequested(void* ptr, void* geom)
{
	static_cast<QWebEnginePage*>(ptr)->geometryChangeRequested(*static_cast<QRect*>(geom));
}

void QWebEnginePage_ConnectIconChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QIcon &)>(&QWebEnginePage::iconChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QIcon &)>(&MyQWebEnginePage::Signal_IconChanged));
}

void QWebEnginePage_DisconnectIconChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QIcon &)>(&QWebEnginePage::iconChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QIcon &)>(&MyQWebEnginePage::Signal_IconChanged));
}

void QWebEnginePage_IconChanged(void* ptr, void* icon)
{
	static_cast<QWebEnginePage*>(ptr)->iconChanged(*static_cast<QIcon*>(icon));
}

void QWebEnginePage_ConnectIconUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &)>(&QWebEnginePage::iconUrlChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &)>(&MyQWebEnginePage::Signal_IconUrlChanged));
}

void QWebEnginePage_DisconnectIconUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &)>(&QWebEnginePage::iconUrlChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &)>(&MyQWebEnginePage::Signal_IconUrlChanged));
}

void QWebEnginePage_IconUrlChanged(void* ptr, void* url)
{
	static_cast<QWebEnginePage*>(ptr)->iconUrlChanged(*static_cast<QUrl*>(url));
}

void QWebEnginePage_JavaScriptAlert(void* ptr, void* securityOrigin, struct QtWebEngine_PackedString msg)
{
	static_cast<QWebEnginePage*>(ptr)->javaScriptAlert(*static_cast<QUrl*>(securityOrigin), QString::fromUtf8(msg.data, msg.len));
}

void QWebEnginePage_JavaScriptAlertDefault(void* ptr, void* securityOrigin, struct QtWebEngine_PackedString msg)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::javaScriptAlert(*static_cast<QUrl*>(securityOrigin), QString::fromUtf8(msg.data, msg.len));
}

void QWebEnginePage_JavaScriptConsoleMessage(void* ptr, long long level, struct QtWebEngine_PackedString message, int lineNumber, struct QtWebEngine_PackedString sourceID)
{
	static_cast<QWebEnginePage*>(ptr)->javaScriptConsoleMessage(static_cast<QWebEnginePage::JavaScriptConsoleMessageLevel>(level), QString::fromUtf8(message.data, message.len), lineNumber, QString::fromUtf8(sourceID.data, sourceID.len));
}

void QWebEnginePage_JavaScriptConsoleMessageDefault(void* ptr, long long level, struct QtWebEngine_PackedString message, int lineNumber, struct QtWebEngine_PackedString sourceID)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::javaScriptConsoleMessage(static_cast<QWebEnginePage::JavaScriptConsoleMessageLevel>(level), QString::fromUtf8(message.data, message.len), lineNumber, QString::fromUtf8(sourceID.data, sourceID.len));
}

void QWebEnginePage_ConnectLinkHovered(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QString &)>(&QWebEnginePage::linkHovered), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QString &)>(&MyQWebEnginePage::Signal_LinkHovered));
}

void QWebEnginePage_DisconnectLinkHovered(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QString &)>(&QWebEnginePage::linkHovered), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QString &)>(&MyQWebEnginePage::Signal_LinkHovered));
}

void QWebEnginePage_LinkHovered(void* ptr, struct QtWebEngine_PackedString url)
{
	static_cast<QWebEnginePage*>(ptr)->linkHovered(QString::fromUtf8(url.data, url.len));
}

void QWebEnginePage_Load(void* ptr, void* url)
{
	static_cast<QWebEnginePage*>(ptr)->load(*static_cast<QUrl*>(url));
}

void QWebEnginePage_Load2(void* ptr, void* request)
{
	static_cast<QWebEnginePage*>(ptr)->load(*static_cast<QWebEngineHttpRequest*>(request));
}

void QWebEnginePage_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(bool)>(&QWebEnginePage::loadFinished), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(bool)>(&MyQWebEnginePage::Signal_LoadFinished));
}

void QWebEnginePage_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(bool)>(&QWebEnginePage::loadFinished), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(bool)>(&MyQWebEnginePage::Signal_LoadFinished));
}

void QWebEnginePage_LoadFinished(void* ptr, char ok)
{
	static_cast<QWebEnginePage*>(ptr)->loadFinished(ok != 0);
}

void QWebEnginePage_ConnectLoadProgress(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(int)>(&QWebEnginePage::loadProgress), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(int)>(&MyQWebEnginePage::Signal_LoadProgress));
}

void QWebEnginePage_DisconnectLoadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(int)>(&QWebEnginePage::loadProgress), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(int)>(&MyQWebEnginePage::Signal_LoadProgress));
}

void QWebEnginePage_LoadProgress(void* ptr, int progress)
{
	static_cast<QWebEnginePage*>(ptr)->loadProgress(progress);
}

void QWebEnginePage_ConnectLoadStarted(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::loadStarted), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_LoadStarted));
}

void QWebEnginePage_DisconnectLoadStarted(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::loadStarted), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_LoadStarted));
}

void QWebEnginePage_LoadStarted(void* ptr)
{
	static_cast<QWebEnginePage*>(ptr)->loadStarted();
}

void QWebEnginePage_ConnectPdfPrintingFinished(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QString &, bool)>(&QWebEnginePage::pdfPrintingFinished), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QString &, bool)>(&MyQWebEnginePage::Signal_PdfPrintingFinished));
}

void QWebEnginePage_DisconnectPdfPrintingFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QString &, bool)>(&QWebEnginePage::pdfPrintingFinished), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QString &, bool)>(&MyQWebEnginePage::Signal_PdfPrintingFinished));
}

void QWebEnginePage_PdfPrintingFinished(void* ptr, struct QtWebEngine_PackedString filePath, char success)
{
	static_cast<QWebEnginePage*>(ptr)->pdfPrintingFinished(QString::fromUtf8(filePath.data, filePath.len), success != 0);
}

void QWebEnginePage_ConnectPrintRequested(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::printRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_PrintRequested));
}

void QWebEnginePage_DisconnectPrintRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::printRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_PrintRequested));
}

void QWebEnginePage_PrintRequested(void* ptr)
{
	static_cast<QWebEnginePage*>(ptr)->printRequested();
}

void QWebEnginePage_PrintToPdf(void* ptr, struct QtWebEngine_PackedString filePath, void* layout)
{
	static_cast<QWebEnginePage*>(ptr)->printToPdf(QString::fromUtf8(filePath.data, filePath.len), *static_cast<QPageLayout*>(layout));
}

void QWebEnginePage_ConnectProxyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QAuthenticator *, const QString &)>(&QWebEnginePage::proxyAuthenticationRequired), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QAuthenticator *, const QString &)>(&MyQWebEnginePage::Signal_ProxyAuthenticationRequired));
}

void QWebEnginePage_DisconnectProxyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &, QAuthenticator *, const QString &)>(&QWebEnginePage::proxyAuthenticationRequired), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &, QAuthenticator *, const QString &)>(&MyQWebEnginePage::Signal_ProxyAuthenticationRequired));
}

void QWebEnginePage_ProxyAuthenticationRequired(void* ptr, void* requestUrl, void* authenticator, struct QtWebEngine_PackedString proxyHost)
{
	static_cast<QWebEnginePage*>(ptr)->proxyAuthenticationRequired(*static_cast<QUrl*>(requestUrl), static_cast<QAuthenticator*>(authenticator), QString::fromUtf8(proxyHost.data, proxyHost.len));
}

void QWebEnginePage_ConnectRecentlyAudibleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(bool)>(&QWebEnginePage::recentlyAudibleChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(bool)>(&MyQWebEnginePage::Signal_RecentlyAudibleChanged));
}

void QWebEnginePage_DisconnectRecentlyAudibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(bool)>(&QWebEnginePage::recentlyAudibleChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(bool)>(&MyQWebEnginePage::Signal_RecentlyAudibleChanged));
}

void QWebEnginePage_RecentlyAudibleChanged(void* ptr, char recentlyAudible)
{
	static_cast<QWebEnginePage*>(ptr)->recentlyAudibleChanged(recentlyAudible != 0);
}

void QWebEnginePage_ConnectRenderProcessTerminated(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(QWebEnginePage::RenderProcessTerminationStatus, int)>(&QWebEnginePage::renderProcessTerminated), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(QWebEnginePage::RenderProcessTerminationStatus, int)>(&MyQWebEnginePage::Signal_RenderProcessTerminated));
}

void QWebEnginePage_DisconnectRenderProcessTerminated(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(QWebEnginePage::RenderProcessTerminationStatus, int)>(&QWebEnginePage::renderProcessTerminated), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(QWebEnginePage::RenderProcessTerminationStatus, int)>(&MyQWebEnginePage::Signal_RenderProcessTerminated));
}

void QWebEnginePage_RenderProcessTerminated(void* ptr, long long terminationStatus, int exitCode)
{
	static_cast<QWebEnginePage*>(ptr)->renderProcessTerminated(static_cast<QWebEnginePage::RenderProcessTerminationStatus>(terminationStatus), exitCode);
}

void QWebEnginePage_ReplaceMisspelledWord(void* ptr, struct QtWebEngine_PackedString replacement)
{
	static_cast<QWebEnginePage*>(ptr)->replaceMisspelledWord(QString::fromUtf8(replacement.data, replacement.len));
}

void QWebEnginePage_RunJavaScript(void* ptr, struct QtWebEngine_PackedString scriptSource)
{
	static_cast<QWebEnginePage*>(ptr)->runJavaScript(QString::fromUtf8(scriptSource.data, scriptSource.len));
}

void QWebEnginePage_RunJavaScript2(void* ptr, struct QtWebEngine_PackedString scriptSource, unsigned int worldId)
{
	static_cast<QWebEnginePage*>(ptr)->runJavaScript(QString::fromUtf8(scriptSource.data, scriptSource.len), worldId);
}

void QWebEnginePage_ConnectScrollPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QPointF &)>(&QWebEnginePage::scrollPositionChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QPointF &)>(&MyQWebEnginePage::Signal_ScrollPositionChanged));
}

void QWebEnginePage_DisconnectScrollPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QPointF &)>(&QWebEnginePage::scrollPositionChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QPointF &)>(&MyQWebEnginePage::Signal_ScrollPositionChanged));
}

void QWebEnginePage_ScrollPositionChanged(void* ptr, void* position)
{
	static_cast<QWebEnginePage*>(ptr)->scrollPositionChanged(*static_cast<QPointF*>(position));
}

void QWebEnginePage_ConnectSelectionChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::selectionChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_SelectionChanged));
}

void QWebEnginePage_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::selectionChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_SelectionChanged));
}

void QWebEnginePage_SelectionChanged(void* ptr)
{
	static_cast<QWebEnginePage*>(ptr)->selectionChanged();
}

void QWebEnginePage_SetAudioMuted(void* ptr, char muted)
{
	static_cast<QWebEnginePage*>(ptr)->setAudioMuted(muted != 0);
}

void QWebEnginePage_SetBackgroundColor(void* ptr, void* color)
{
	static_cast<QWebEnginePage*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QWebEnginePage_SetContent(void* ptr, void* data, struct QtWebEngine_PackedString mimeType, void* baseUrl)
{
	static_cast<QWebEnginePage*>(ptr)->setContent(*static_cast<QByteArray*>(data), QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QUrl*>(baseUrl));
}

void QWebEnginePage_SetDevToolsPage(void* ptr, void* page)
{
	static_cast<QWebEnginePage*>(ptr)->setDevToolsPage(static_cast<QWebEnginePage*>(page));
}

void QWebEnginePage_SetFeaturePermission(void* ptr, void* securityOrigin, long long feature, long long policy)
{
	static_cast<QWebEnginePage*>(ptr)->setFeaturePermission(*static_cast<QUrl*>(securityOrigin), static_cast<QWebEnginePage::Feature>(feature), static_cast<QWebEnginePage::PermissionPolicy>(policy));
}

void QWebEnginePage_SetHtml(void* ptr, struct QtWebEngine_PackedString html, void* baseUrl)
{
	static_cast<QWebEnginePage*>(ptr)->setHtml(QString::fromUtf8(html.data, html.len), *static_cast<QUrl*>(baseUrl));
}

void QWebEnginePage_SetInspectedPage(void* ptr, void* page)
{
	static_cast<QWebEnginePage*>(ptr)->setInspectedPage(static_cast<QWebEnginePage*>(page));
}

void QWebEnginePage_SetUrl(void* ptr, void* url)
{
	static_cast<QWebEnginePage*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebEnginePage_SetView(void* ptr, void* view)
{
	static_cast<QWebEnginePage*>(ptr)->setView(static_cast<QWidget*>(view));
}

void QWebEnginePage_SetWebChannel(void* ptr, void* vqw)
{
	static_cast<QWebEnginePage*>(ptr)->setWebChannel(static_cast<QWebChannel*>(vqw));
}

void QWebEnginePage_SetWebChannel2(void* ptr, void* vqw, unsigned int worldId)
{
	static_cast<QWebEnginePage*>(ptr)->setWebChannel(static_cast<QWebChannel*>(vqw), worldId);
}

void QWebEnginePage_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebEnginePage*>(ptr)->setZoomFactor(factor);
}

void QWebEnginePage_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QString &)>(&QWebEnginePage::titleChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QString &)>(&MyQWebEnginePage::Signal_TitleChanged));
}

void QWebEnginePage_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QString &)>(&QWebEnginePage::titleChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QString &)>(&MyQWebEnginePage::Signal_TitleChanged));
}

void QWebEnginePage_TitleChanged(void* ptr, struct QtWebEngine_PackedString title)
{
	static_cast<QWebEnginePage*>(ptr)->titleChanged(QString::fromUtf8(title.data, title.len));
}

void QWebEnginePage_TriggerAction(void* ptr, long long action, char checked)
{
	static_cast<QWebEnginePage*>(ptr)->triggerAction(static_cast<QWebEnginePage::WebAction>(action), checked != 0);
}

void QWebEnginePage_TriggerActionDefault(void* ptr, long long action, char checked)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::triggerAction(static_cast<QWebEnginePage::WebAction>(action), checked != 0);
}

void QWebEnginePage_ConnectUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &)>(&QWebEnginePage::urlChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &)>(&MyQWebEnginePage::Signal_UrlChanged));
}

void QWebEnginePage_DisconnectUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)(const QUrl &)>(&QWebEnginePage::urlChanged), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)(const QUrl &)>(&MyQWebEnginePage::Signal_UrlChanged));
}

void QWebEnginePage_UrlChanged(void* ptr, void* url)
{
	static_cast<QWebEnginePage*>(ptr)->urlChanged(*static_cast<QUrl*>(url));
}

void QWebEnginePage_ConnectWindowCloseRequested(void* ptr)
{
	QObject::connect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::windowCloseRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_WindowCloseRequested));
}

void QWebEnginePage_DisconnectWindowCloseRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebEnginePage*>(ptr), static_cast<void (QWebEnginePage::*)()>(&QWebEnginePage::windowCloseRequested), static_cast<MyQWebEnginePage*>(ptr), static_cast<void (MyQWebEnginePage::*)()>(&MyQWebEnginePage::Signal_WindowCloseRequested));
}

void QWebEnginePage_WindowCloseRequested(void* ptr)
{
	static_cast<QWebEnginePage*>(ptr)->windowCloseRequested();
}

void QWebEnginePage_DestroyQWebEnginePage(void* ptr)
{
	static_cast<QWebEnginePage*>(ptr)->~QWebEnginePage();
}

void QWebEnginePage_DestroyQWebEnginePageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebEnginePage_Action(void* ptr, long long action)
{
	return static_cast<QWebEnginePage*>(ptr)->action(static_cast<QWebEnginePage::WebAction>(action));
}

void* QWebEnginePage_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QWebEnginePage*>(ptr)->backgroundColor());
}

void* QWebEnginePage_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebEnginePage*>(ptr)->icon());
}

void* QWebEnginePage_ScrollPosition(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QWebEnginePage*>(ptr)->scrollPosition(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QWebEnginePage_ContentsSize(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QWebEnginePage*>(ptr)->contentsSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

struct QtWebEngine_PackedString QWebEnginePage_SelectedText(void* ptr)
{
	return ({ QByteArray t729dfc = static_cast<QWebEnginePage*>(ptr)->selectedText().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t729dfc.prepend("WHITESPACE").constData()+10), t729dfc.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEnginePage_Title(void* ptr)
{
	return ({ QByteArray t311622 = static_cast<QWebEnginePage*>(ptr)->title().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t311622.prepend("WHITESPACE").constData()+10), t311622.size()-10 }; });
}

void* QWebEnginePage_IconUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEnginePage*>(ptr)->iconUrl());
}

void* QWebEnginePage_RequestedUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEnginePage*>(ptr)->requestedUrl());
}

void* QWebEnginePage_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEnginePage*>(ptr)->url());
}

void* QWebEnginePage_WebChannel(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->webChannel();
}

void* QWebEnginePage_DevToolsPage(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->devToolsPage();
}

void* QWebEnginePage_InspectedPage(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->inspectedPage();
}

void* QWebEnginePage_Profile(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->profile();
}

void* QWebEnginePage_Settings(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->settings();
}

void* QWebEnginePage_View(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->view();
}

char QWebEnginePage_HasSelection(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->hasSelection();
}

char QWebEnginePage_IsAudioMuted(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->isAudioMuted();
}

char QWebEnginePage_RecentlyAudible(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->recentlyAudible();
}

void* QWebEnginePage_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::metaObject());
}

double QWebEnginePage_ZoomFactor(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->zoomFactor();
}

void* QWebEnginePage___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEnginePage___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEnginePage___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEnginePage___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEnginePage___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEnginePage___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEnginePage___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEnginePage___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEnginePage___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEnginePage___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEnginePage___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEnginePage___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEnginePage___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEnginePage___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEnginePage___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEnginePage_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEnginePage_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEnginePage_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEnginePage_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::customEvent(static_cast<QEvent*>(event));
}

void QWebEnginePage_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::deleteLater();
}

void QWebEnginePage_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEnginePage_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEngineProfile: public QWebEngineProfile
{
public:
	MyQWebEngineProfile(QObject *parent = Q_NULLPTR) : QWebEngineProfile(parent) {QWebEngineProfile_QWebEngineProfile_QRegisterMetaType();};
	MyQWebEngineProfile(const QString &name, QObject *parent = Q_NULLPTR) : QWebEngineProfile(name, parent) {QWebEngineProfile_QWebEngineProfile_QRegisterMetaType();};
	 ~MyQWebEngineProfile() { callbackQWebEngineProfile_DestroyQWebEngineProfile(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEngineProfile_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebEngineProfile_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineProfile_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEngineProfile_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineProfile_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineProfile_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineProfile_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineProfile_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineProfile_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineProfile_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineProfile_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineProfile*)

int QWebEngineProfile_QWebEngineProfile_QRegisterMetaType(){qRegisterMetaType<QWebEngineProfile*>(); return qRegisterMetaType<MyQWebEngineProfile*>();}

struct QtWebEngine_PackedString QWebEngineProfile_QWebEngineProfile_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tf8cda5 = QWebEngineProfile::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tf8cda5.prepend("WHITESPACE").constData()+10), tf8cda5.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineProfile_QWebEngineProfile_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tb63b40 = QWebEngineProfile::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tb63b40.prepend("WHITESPACE").constData()+10), tb63b40.size()-10 }; });
}

void* QWebEngineProfile_CookieStore(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->cookieStore();
}

void* QWebEngineProfile_QWebEngineProfile_DefaultProfile()
{
	return QWebEngineProfile::defaultProfile();
}

void* QWebEngineProfile_NewQWebEngineProfile(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebEngineProfile(static_cast<QObject*>(parent));
	}
}

void* QWebEngineProfile_NewQWebEngineProfile2(struct QtWebEngine_PackedString name, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQWebEngineProfile(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(parent));
	}
}

void QWebEngineProfile_ClearAllVisitedLinks(void* ptr)
{
	static_cast<QWebEngineProfile*>(ptr)->clearAllVisitedLinks();
}

void QWebEngineProfile_ClearHttpCache(void* ptr)
{
	static_cast<QWebEngineProfile*>(ptr)->clearHttpCache();
}

void QWebEngineProfile_ClearVisitedLinks(void* ptr, void* urls)
{
	static_cast<QWebEngineProfile*>(ptr)->clearVisitedLinks(*static_cast<QList<QUrl>*>(urls));
}

void QWebEngineProfile_InstallUrlSchemeHandler(void* ptr, void* scheme, void* vqw)
{
	static_cast<QWebEngineProfile*>(ptr)->installUrlSchemeHandler(*static_cast<QByteArray*>(scheme), static_cast<QWebEngineUrlSchemeHandler*>(vqw));
}

void QWebEngineProfile_RemoveAllUrlSchemeHandlers(void* ptr)
{
	static_cast<QWebEngineProfile*>(ptr)->removeAllUrlSchemeHandlers();
}

void QWebEngineProfile_RemoveUrlScheme(void* ptr, void* scheme)
{
	static_cast<QWebEngineProfile*>(ptr)->removeUrlScheme(*static_cast<QByteArray*>(scheme));
}

void QWebEngineProfile_RemoveUrlSchemeHandler(void* ptr, void* vqw)
{
	static_cast<QWebEngineProfile*>(ptr)->removeUrlSchemeHandler(static_cast<QWebEngineUrlSchemeHandler*>(vqw));
}

void QWebEngineProfile_SetCachePath(void* ptr, struct QtWebEngine_PackedString path)
{
	static_cast<QWebEngineProfile*>(ptr)->setCachePath(QString::fromUtf8(path.data, path.len));
}

void QWebEngineProfile_SetHttpAcceptLanguage(void* ptr, struct QtWebEngine_PackedString httpAcceptLanguage)
{
	static_cast<QWebEngineProfile*>(ptr)->setHttpAcceptLanguage(QString::fromUtf8(httpAcceptLanguage.data, httpAcceptLanguage.len));
}

void QWebEngineProfile_SetHttpCacheMaximumSize(void* ptr, int maxSize)
{
	static_cast<QWebEngineProfile*>(ptr)->setHttpCacheMaximumSize(maxSize);
}

void QWebEngineProfile_SetHttpCacheType(void* ptr, long long vqw)
{
	static_cast<QWebEngineProfile*>(ptr)->setHttpCacheType(static_cast<QWebEngineProfile::HttpCacheType>(vqw));
}

void QWebEngineProfile_SetHttpUserAgent(void* ptr, struct QtWebEngine_PackedString userAgent)
{
	static_cast<QWebEngineProfile*>(ptr)->setHttpUserAgent(QString::fromUtf8(userAgent.data, userAgent.len));
}

void QWebEngineProfile_SetPersistentCookiesPolicy(void* ptr, long long vqw)
{
	static_cast<QWebEngineProfile*>(ptr)->setPersistentCookiesPolicy(static_cast<QWebEngineProfile::PersistentCookiesPolicy>(vqw));
}

void QWebEngineProfile_SetPersistentStoragePath(void* ptr, struct QtWebEngine_PackedString path)
{
	static_cast<QWebEngineProfile*>(ptr)->setPersistentStoragePath(QString::fromUtf8(path.data, path.len));
}

void QWebEngineProfile_SetRequestInterceptor(void* ptr, void* interceptor)
{
	static_cast<QWebEngineProfile*>(ptr)->setRequestInterceptor(static_cast<QWebEngineUrlRequestInterceptor*>(interceptor));
}

void QWebEngineProfile_SetSpellCheckEnabled(void* ptr, char enabled)
{
	static_cast<QWebEngineProfile*>(ptr)->setSpellCheckEnabled(enabled != 0);
}

void QWebEngineProfile_SetSpellCheckLanguages(void* ptr, struct QtWebEngine_PackedString languages)
{
	static_cast<QWebEngineProfile*>(ptr)->setSpellCheckLanguages(QString::fromUtf8(languages.data, languages.len).split("|", QString::SkipEmptyParts));
}

void QWebEngineProfile_DestroyQWebEngineProfile(void* ptr)
{
	static_cast<QWebEngineProfile*>(ptr)->~QWebEngineProfile();
}

void QWebEngineProfile_DestroyQWebEngineProfileDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtWebEngine_PackedString QWebEngineProfile_CachePath(void* ptr)
{
	return ({ QByteArray td354e7 = static_cast<QWebEngineProfile*>(ptr)->cachePath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(td354e7.prepend("WHITESPACE").constData()+10), td354e7.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineProfile_HttpAcceptLanguage(void* ptr)
{
	return ({ QByteArray t46ad8d = static_cast<QWebEngineProfile*>(ptr)->httpAcceptLanguage().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t46ad8d.prepend("WHITESPACE").constData()+10), t46ad8d.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineProfile_HttpUserAgent(void* ptr)
{
	return ({ QByteArray t5c173a = static_cast<QWebEngineProfile*>(ptr)->httpUserAgent().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t5c173a.prepend("WHITESPACE").constData()+10), t5c173a.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineProfile_PersistentStoragePath(void* ptr)
{
	return ({ QByteArray t1c9837 = static_cast<QWebEngineProfile*>(ptr)->persistentStoragePath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t1c9837.prepend("WHITESPACE").constData()+10), t1c9837.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineProfile_StorageName(void* ptr)
{
	return ({ QByteArray tcbe040 = static_cast<QWebEngineProfile*>(ptr)->storageName().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tcbe040.prepend("WHITESPACE").constData()+10), tcbe040.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineProfile_SpellCheckLanguages(void* ptr)
{
	return ({ QByteArray tba4fdb = static_cast<QWebEngineProfile*>(ptr)->spellCheckLanguages().join("|").toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tba4fdb.prepend("WHITESPACE").constData()+10), tba4fdb.size()-10 }; });
}

long long QWebEngineProfile_HttpCacheType(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->httpCacheType();
}

long long QWebEngineProfile_PersistentCookiesPolicy(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->persistentCookiesPolicy();
}

void* QWebEngineProfile_Scripts(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->scripts();
}

void* QWebEngineProfile_Settings(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->settings();
}

char QWebEngineProfile_IsOffTheRecord(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->isOffTheRecord();
}

char QWebEngineProfile_IsSpellCheckEnabled(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->isSpellCheckEnabled();
}

char QWebEngineProfile_VisitedLinksContainsUrl(void* ptr, void* url)
{
	return static_cast<QWebEngineProfile*>(ptr)->visitedLinksContainsUrl(*static_cast<QUrl*>(url));
}

void* QWebEngineProfile_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::metaObject());
}

void* QWebEngineProfile_UrlSchemeHandler(void* ptr, void* vqb)
{
	return const_cast<QWebEngineUrlSchemeHandler*>(static_cast<QWebEngineProfile*>(ptr)->urlSchemeHandler(*static_cast<QByteArray*>(vqb)));
}

int QWebEngineProfile_HttpCacheMaximumSize(void* ptr)
{
	return static_cast<QWebEngineProfile*>(ptr)->httpCacheMaximumSize();
}

void* QWebEngineProfile___clearVisitedLinks_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineProfile___clearVisitedLinks_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QWebEngineProfile___clearVisitedLinks_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QWebEngineProfile___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineProfile___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineProfile___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineProfile___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineProfile___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineProfile___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineProfile___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineProfile___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineProfile___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineProfile___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineProfile___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineProfile___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineProfile___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineProfile___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineProfile___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEngineProfile_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::event(static_cast<QEvent*>(e));
}

char QWebEngineProfile_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineProfile_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineProfile_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineProfile_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineProfile_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::deleteLater();
}

void QWebEngineProfile_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineProfile_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebEngineQuotaRequest_Accept(void* ptr)
{
	static_cast<QWebEngineQuotaRequest*>(ptr)->accept();
}

void QWebEngineQuotaRequest_Reject(void* ptr)
{
	static_cast<QWebEngineQuotaRequest*>(ptr)->reject();
}

void* QWebEngineQuotaRequest_Origin(void* ptr)
{
	return new QUrl(static_cast<QWebEngineQuotaRequest*>(ptr)->origin());
}

long long QWebEngineQuotaRequest_RequestedSize(void* ptr)
{
	return static_cast<QWebEngineQuotaRequest*>(ptr)->requestedSize();
}

void QWebEngineRegisterProtocolHandlerRequest_Accept(void* ptr)
{
	static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->accept();
}

void QWebEngineRegisterProtocolHandlerRequest_Reject(void* ptr)
{
	static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->reject();
}

struct QtWebEngine_PackedString QWebEngineRegisterProtocolHandlerRequest_Scheme(void* ptr)
{
	return ({ QByteArray t21b2e9 = static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->scheme().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t21b2e9.prepend("WHITESPACE").constData()+10), t21b2e9.size()-10 }; });
}

void* QWebEngineRegisterProtocolHandlerRequest_Origin(void* ptr)
{
	return new QUrl(static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->origin());
}

void* QWebEngineScript_NewQWebEngineScript()
{
	return new QWebEngineScript();
}

void* QWebEngineScript_NewQWebEngineScript2(void* other)
{
	return new QWebEngineScript(*static_cast<QWebEngineScript*>(other));
}

void QWebEngineScript_SetInjectionPoint(void* ptr, long long vqw)
{
	static_cast<QWebEngineScript*>(ptr)->setInjectionPoint(static_cast<QWebEngineScript::InjectionPoint>(vqw));
}

void QWebEngineScript_SetName(void* ptr, struct QtWebEngine_PackedString vqs)
{
	static_cast<QWebEngineScript*>(ptr)->setName(QString::fromUtf8(vqs.data, vqs.len));
}

void QWebEngineScript_SetRunsOnSubFrames(void* ptr, char on)
{
	static_cast<QWebEngineScript*>(ptr)->setRunsOnSubFrames(on != 0);
}

void QWebEngineScript_SetSourceCode(void* ptr, struct QtWebEngine_PackedString vqs)
{
	static_cast<QWebEngineScript*>(ptr)->setSourceCode(QString::fromUtf8(vqs.data, vqs.len));
}

void QWebEngineScript_SetWorldId(void* ptr, unsigned int vqu)
{
	static_cast<QWebEngineScript*>(ptr)->setWorldId(vqu);
}

void QWebEngineScript_Swap(void* ptr, void* other)
{
	static_cast<QWebEngineScript*>(ptr)->swap(*static_cast<QWebEngineScript*>(other));
}

void QWebEngineScript_DestroyQWebEngineScript(void* ptr)
{
	static_cast<QWebEngineScript*>(ptr)->~QWebEngineScript();
}

struct QtWebEngine_PackedString QWebEngineScript_Name(void* ptr)
{
	return ({ QByteArray te185e9 = static_cast<QWebEngineScript*>(ptr)->name().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(te185e9.prepend("WHITESPACE").constData()+10), te185e9.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineScript_SourceCode(void* ptr)
{
	return ({ QByteArray t7bcb57 = static_cast<QWebEngineScript*>(ptr)->sourceCode().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t7bcb57.prepend("WHITESPACE").constData()+10), t7bcb57.size()-10 }; });
}

long long QWebEngineScript_InjectionPoint(void* ptr)
{
	return static_cast<QWebEngineScript*>(ptr)->injectionPoint();
}

char QWebEngineScript_IsNull(void* ptr)
{
	return static_cast<QWebEngineScript*>(ptr)->isNull();
}

char QWebEngineScript_RunsOnSubFrames(void* ptr)
{
	return static_cast<QWebEngineScript*>(ptr)->runsOnSubFrames();
}

unsigned int QWebEngineScript_WorldId(void* ptr)
{
	return static_cast<QWebEngineScript*>(ptr)->worldId();
}

void QWebEngineScriptCollection_Clear(void* ptr)
{
	static_cast<QWebEngineScriptCollection*>(ptr)->clear();
}

void QWebEngineScriptCollection_DestroyQWebEngineScriptCollection(void* ptr)
{
	static_cast<QWebEngineScriptCollection*>(ptr)->~QWebEngineScriptCollection();
}

char QWebEngineScriptCollection_IsEmpty(void* ptr)
{
	return static_cast<QWebEngineScriptCollection*>(ptr)->isEmpty();
}

int QWebEngineScriptCollection_Count(void* ptr)
{
	return static_cast<QWebEngineScriptCollection*>(ptr)->count();
}

int QWebEngineScriptCollection_Size(void* ptr)
{
	return static_cast<QWebEngineScriptCollection*>(ptr)->size();
}

void* QWebEngineSettings_QWebEngineSettings_DefaultSettings()
{
	return QWebEngineSettings::defaultSettings();
}

void* QWebEngineSettings_QWebEngineSettings_GlobalSettings()
{
	return QWebEngineSettings::globalSettings();
}

void QWebEngineSettings_ResetAttribute(void* ptr, long long attr)
{
	static_cast<QWebEngineSettings*>(ptr)->resetAttribute(static_cast<QWebEngineSettings::WebAttribute>(attr));
}

void QWebEngineSettings_ResetFontFamily(void* ptr, long long which)
{
	static_cast<QWebEngineSettings*>(ptr)->resetFontFamily(static_cast<QWebEngineSettings::FontFamily>(which));
}

void QWebEngineSettings_ResetFontSize(void* ptr, long long ty)
{
	static_cast<QWebEngineSettings*>(ptr)->resetFontSize(static_cast<QWebEngineSettings::FontSize>(ty));
}

void QWebEngineSettings_ResetUnknownUrlSchemePolicy(void* ptr)
{
	static_cast<QWebEngineSettings*>(ptr)->resetUnknownUrlSchemePolicy();
}

void QWebEngineSettings_SetAttribute(void* ptr, long long attr, char on)
{
	static_cast<QWebEngineSettings*>(ptr)->setAttribute(static_cast<QWebEngineSettings::WebAttribute>(attr), on != 0);
}

void QWebEngineSettings_SetDefaultTextEncoding(void* ptr, struct QtWebEngine_PackedString encoding)
{
	static_cast<QWebEngineSettings*>(ptr)->setDefaultTextEncoding(QString::fromUtf8(encoding.data, encoding.len));
}

void QWebEngineSettings_SetFontFamily(void* ptr, long long which, struct QtWebEngine_PackedString family)
{
	static_cast<QWebEngineSettings*>(ptr)->setFontFamily(static_cast<QWebEngineSettings::FontFamily>(which), QString::fromUtf8(family.data, family.len));
}

void QWebEngineSettings_SetFontSize(void* ptr, long long ty, int size)
{
	static_cast<QWebEngineSettings*>(ptr)->setFontSize(static_cast<QWebEngineSettings::FontSize>(ty), size);
}

void QWebEngineSettings_SetUnknownUrlSchemePolicy(void* ptr, long long policy)
{
	static_cast<QWebEngineSettings*>(ptr)->setUnknownUrlSchemePolicy(static_cast<QWebEngineSettings::UnknownUrlSchemePolicy>(policy));
}

struct QtWebEngine_PackedString QWebEngineSettings_DefaultTextEncoding(void* ptr)
{
	return ({ QByteArray t01cf63 = static_cast<QWebEngineSettings*>(ptr)->defaultTextEncoding().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t01cf63.prepend("WHITESPACE").constData()+10), t01cf63.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineSettings_FontFamily(void* ptr, long long which)
{
	return ({ QByteArray t118959 = static_cast<QWebEngineSettings*>(ptr)->fontFamily(static_cast<QWebEngineSettings::FontFamily>(which)).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t118959.prepend("WHITESPACE").constData()+10), t118959.size()-10 }; });
}

long long QWebEngineSettings_UnknownUrlSchemePolicy(void* ptr)
{
	return static_cast<QWebEngineSettings*>(ptr)->unknownUrlSchemePolicy();
}

char QWebEngineSettings_TestAttribute(void* ptr, long long attr)
{
	return static_cast<QWebEngineSettings*>(ptr)->testAttribute(static_cast<QWebEngineSettings::WebAttribute>(attr));
}

int QWebEngineSettings_FontSize(void* ptr, long long ty)
{
	return static_cast<QWebEngineSettings*>(ptr)->fontSize(static_cast<QWebEngineSettings::FontSize>(ty));
}

void QWebEngineUrlRequestInfo_Block(void* ptr, char shouldBlock)
{
	static_cast<QWebEngineUrlRequestInfo*>(ptr)->block(shouldBlock != 0);
}

void QWebEngineUrlRequestInfo_Redirect(void* ptr, void* url)
{
	static_cast<QWebEngineUrlRequestInfo*>(ptr)->redirect(*static_cast<QUrl*>(url));
}

void QWebEngineUrlRequestInfo_SetHttpHeader(void* ptr, void* name, void* value)
{
	static_cast<QWebEngineUrlRequestInfo*>(ptr)->setHttpHeader(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

void* QWebEngineUrlRequestInfo_RequestMethod(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineUrlRequestInfo*>(ptr)->requestMethod());
}

void* QWebEngineUrlRequestInfo_FirstPartyUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestInfo*>(ptr)->firstPartyUrl());
}

void* QWebEngineUrlRequestInfo_RequestUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestInfo*>(ptr)->requestUrl());
}

long long QWebEngineUrlRequestInfo_NavigationType(void* ptr)
{
	return static_cast<QWebEngineUrlRequestInfo*>(ptr)->navigationType();
}

long long QWebEngineUrlRequestInfo_ResourceType(void* ptr)
{
	return static_cast<QWebEngineUrlRequestInfo*>(ptr)->resourceType();
}

class MyQWebEngineUrlRequestInterceptor: public QWebEngineUrlRequestInterceptor
{
public:
	MyQWebEngineUrlRequestInterceptor(QObject *p = Q_NULLPTR) : QWebEngineUrlRequestInterceptor(p) {QWebEngineUrlRequestInterceptor_QWebEngineUrlRequestInterceptor_QRegisterMetaType();};
	void interceptRequest(QWebEngineUrlRequestInfo & info) { callbackQWebEngineUrlRequestInterceptor_InterceptRequest(this, static_cast<QWebEngineUrlRequestInfo*>(&info)); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEngineUrlRequestInterceptor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebEngineUrlRequestInterceptor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineUrlRequestInterceptor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEngineUrlRequestInterceptor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestInterceptor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineUrlRequestInterceptor_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineUrlRequestInterceptor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineUrlRequestInterceptor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestInterceptor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineUrlRequestInterceptor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineUrlRequestInterceptor_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineUrlRequestInterceptor*)

int QWebEngineUrlRequestInterceptor_QWebEngineUrlRequestInterceptor_QRegisterMetaType(){qRegisterMetaType<QWebEngineUrlRequestInterceptor*>(); return qRegisterMetaType<MyQWebEngineUrlRequestInterceptor*>();}

struct QtWebEngine_PackedString QWebEngineUrlRequestInterceptor_QWebEngineUrlRequestInterceptor_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t612836 = QWebEngineUrlRequestInterceptor::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t612836.prepend("WHITESPACE").constData()+10), t612836.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineUrlRequestInterceptor_QWebEngineUrlRequestInterceptor_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tbb9bf3 = QWebEngineUrlRequestInterceptor::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tbb9bf3.prepend("WHITESPACE").constData()+10), tbb9bf3.size()-10 }; });
}

void* QWebEngineUrlRequestInterceptor_NewQWebEngineUrlRequestInterceptor(void* p)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QCameraImageCapture*>(p));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QDBusPendingCallWatcher*>(p));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QExtensionFactory*>(p));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QExtensionManager*>(p));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QGraphicsObject*>(p));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QGraphicsWidget*>(p));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QLayout*>(p));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QMediaPlaylist*>(p));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QMediaRecorder*>(p));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QOffscreenSurface*>(p));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QPaintDeviceWindow*>(p));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QPdfWriter*>(p));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QQuickItem*>(p));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QRadioData*>(p));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QWidget*>(p));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QWindow*>(p));
	} else {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QObject*>(p));
	}
}

void QWebEngineUrlRequestInterceptor_InterceptRequest(void* ptr, void* info)
{
	static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->interceptRequest(*static_cast<QWebEngineUrlRequestInfo*>(info));
}

void* QWebEngineUrlRequestInterceptor_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::metaObject());
}

void* QWebEngineUrlRequestInterceptor___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineUrlRequestInterceptor___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineUrlRequestInterceptor___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineUrlRequestInterceptor___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestInterceptor___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestInterceptor___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlRequestInterceptor___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestInterceptor___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestInterceptor___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlRequestInterceptor___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestInterceptor___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestInterceptor___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlRequestInterceptor___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestInterceptor___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestInterceptor___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEngineUrlRequestInterceptor_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::event(static_cast<QEvent*>(e));
}

char QWebEngineUrlRequestInterceptor_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineUrlRequestInterceptor_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineUrlRequestInterceptor_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineUrlRequestInterceptor_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineUrlRequestInterceptor_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::deleteLater();
}

void QWebEngineUrlRequestInterceptor_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineUrlRequestInterceptor_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEngineUrlRequestJob: public QWebEngineUrlRequestJob
{
public:
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEngineUrlRequestJob_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebEngineUrlRequestJob_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineUrlRequestJob_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEngineUrlRequestJob_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestJob_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineUrlRequestJob_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineUrlRequestJob_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineUrlRequestJob_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestJob_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineUrlRequestJob_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineUrlRequestJob_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineUrlRequestJob*)

int QWebEngineUrlRequestJob_QWebEngineUrlRequestJob_QRegisterMetaType(){qRegisterMetaType<QWebEngineUrlRequestJob*>(); return qRegisterMetaType<MyQWebEngineUrlRequestJob*>();}

struct QtWebEngine_PackedString QWebEngineUrlRequestJob_QWebEngineUrlRequestJob_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tf9e2e0 = QWebEngineUrlRequestJob::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tf9e2e0.prepend("WHITESPACE").constData()+10), tf9e2e0.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineUrlRequestJob_QWebEngineUrlRequestJob_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t4194d7 = QWebEngineUrlRequestJob::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t4194d7.prepend("WHITESPACE").constData()+10), t4194d7.size()-10 }; });
}

void QWebEngineUrlRequestJob_Fail(void* ptr, long long r)
{
	static_cast<QWebEngineUrlRequestJob*>(ptr)->fail(static_cast<QWebEngineUrlRequestJob::Error>(r));
}

void QWebEngineUrlRequestJob_Redirect(void* ptr, void* url)
{
	static_cast<QWebEngineUrlRequestJob*>(ptr)->redirect(*static_cast<QUrl*>(url));
}

void QWebEngineUrlRequestJob_Reply(void* ptr, void* contentType, void* device)
{
	static_cast<QWebEngineUrlRequestJob*>(ptr)->reply(*static_cast<QByteArray*>(contentType), static_cast<QIODevice*>(device));
}

void* QWebEngineUrlRequestJob_RequestMethod(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineUrlRequestJob*>(ptr)->requestMethod());
}

void* QWebEngineUrlRequestJob_Initiator(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestJob*>(ptr)->initiator());
}

void* QWebEngineUrlRequestJob_RequestUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestJob*>(ptr)->requestUrl());
}

void* QWebEngineUrlRequestJob_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::metaObject());
}

void* QWebEngineUrlRequestJob___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineUrlRequestJob___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineUrlRequestJob___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineUrlRequestJob___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestJob___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestJob___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlRequestJob___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestJob___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestJob___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlRequestJob___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestJob___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestJob___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlRequestJob___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestJob___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestJob___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEngineUrlRequestJob_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::event(static_cast<QEvent*>(e));
}

char QWebEngineUrlRequestJob_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineUrlRequestJob_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineUrlRequestJob_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineUrlRequestJob_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineUrlRequestJob_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::deleteLater();
}

void QWebEngineUrlRequestJob_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineUrlRequestJob_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebEngineUrlScheme_QWebEngineUrlScheme_SchemeByName(void* name)
{
	return new QWebEngineUrlScheme(QWebEngineUrlScheme::schemeByName(*static_cast<QByteArray*>(name)));
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme()
{
	return new QWebEngineUrlScheme();
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme4(void* that)
{
	return new QWebEngineUrlScheme(*static_cast<QWebEngineUrlScheme*>(that));
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme2(void* name)
{
	return new QWebEngineUrlScheme(*static_cast<QByteArray*>(name));
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme3(void* that)
{
	return new QWebEngineUrlScheme(*static_cast<QWebEngineUrlScheme*>(that));
}

void QWebEngineUrlScheme_QWebEngineUrlScheme_RegisterScheme(void* scheme)
{
	QWebEngineUrlScheme::registerScheme(*static_cast<QWebEngineUrlScheme*>(scheme));
}

void QWebEngineUrlScheme_SetDefaultPort(void* ptr, int newValue)
{
	static_cast<QWebEngineUrlScheme*>(ptr)->setDefaultPort(newValue);
}

void QWebEngineUrlScheme_SetFlags(void* ptr, long long newValue)
{
	static_cast<QWebEngineUrlScheme*>(ptr)->setFlags(static_cast<QWebEngineUrlScheme::Flag>(newValue));
}

void QWebEngineUrlScheme_SetName(void* ptr, void* newValue)
{
	static_cast<QWebEngineUrlScheme*>(ptr)->setName(*static_cast<QByteArray*>(newValue));
}

void QWebEngineUrlScheme_SetSyntax(void* ptr, long long newValue)
{
	static_cast<QWebEngineUrlScheme*>(ptr)->setSyntax(static_cast<QWebEngineUrlScheme::Syntax>(newValue));
}

void QWebEngineUrlScheme_DestroyQWebEngineUrlScheme(void* ptr)
{
	static_cast<QWebEngineUrlScheme*>(ptr)->~QWebEngineUrlScheme();
}

void* QWebEngineUrlScheme_Name(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineUrlScheme*>(ptr)->name());
}

long long QWebEngineUrlScheme_Flags(void* ptr)
{
	return static_cast<QWebEngineUrlScheme*>(ptr)->flags();
}

int QWebEngineUrlScheme_DefaultPort(void* ptr)
{
	return static_cast<QWebEngineUrlScheme*>(ptr)->defaultPort();
}

class MyQWebEngineUrlSchemeHandler: public QWebEngineUrlSchemeHandler
{
public:
	MyQWebEngineUrlSchemeHandler(QObject *parent = Q_NULLPTR) : QWebEngineUrlSchemeHandler(parent) {QWebEngineUrlSchemeHandler_QWebEngineUrlSchemeHandler_QRegisterMetaType();};
	void requestStarted(QWebEngineUrlRequestJob * request) { callbackQWebEngineUrlSchemeHandler_RequestStarted(this, request); };
	 ~MyQWebEngineUrlSchemeHandler() { callbackQWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEngineUrlSchemeHandler_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebEngineUrlSchemeHandler_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineUrlSchemeHandler_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEngineUrlSchemeHandler_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlSchemeHandler_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineUrlSchemeHandler_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineUrlSchemeHandler_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineUrlSchemeHandler_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlSchemeHandler_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineUrlSchemeHandler_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineUrlSchemeHandler_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineUrlSchemeHandler*)

int QWebEngineUrlSchemeHandler_QWebEngineUrlSchemeHandler_QRegisterMetaType(){qRegisterMetaType<QWebEngineUrlSchemeHandler*>(); return qRegisterMetaType<MyQWebEngineUrlSchemeHandler*>();}

struct QtWebEngine_PackedString QWebEngineUrlSchemeHandler_QWebEngineUrlSchemeHandler_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tb887be = QWebEngineUrlSchemeHandler::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tb887be.prepend("WHITESPACE").constData()+10), tb887be.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineUrlSchemeHandler_QWebEngineUrlSchemeHandler_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t7ab116 = QWebEngineUrlSchemeHandler::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t7ab116.prepend("WHITESPACE").constData()+10), t7ab116.size()-10 }; });
}

void* QWebEngineUrlSchemeHandler_NewQWebEngineUrlSchemeHandler(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QObject*>(parent));
	}
}

void QWebEngineUrlSchemeHandler_RequestStarted(void* ptr, void* request)
{
	static_cast<QWebEngineUrlSchemeHandler*>(ptr)->requestStarted(static_cast<QWebEngineUrlRequestJob*>(request));
}

void QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(void* ptr)
{
	static_cast<QWebEngineUrlSchemeHandler*>(ptr)->~QWebEngineUrlSchemeHandler();
}

void QWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebEngineUrlSchemeHandler_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::metaObject());
}

void* QWebEngineUrlSchemeHandler___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineUrlSchemeHandler___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineUrlSchemeHandler___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineUrlSchemeHandler___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlSchemeHandler___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlSchemeHandler___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlSchemeHandler___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlSchemeHandler___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlSchemeHandler___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlSchemeHandler___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlSchemeHandler___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlSchemeHandler___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineUrlSchemeHandler___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlSchemeHandler___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlSchemeHandler___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEngineUrlSchemeHandler_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::event(static_cast<QEvent*>(e));
}

char QWebEngineUrlSchemeHandler_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineUrlSchemeHandler_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineUrlSchemeHandler_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineUrlSchemeHandler_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineUrlSchemeHandler_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::deleteLater();
}

void QWebEngineUrlSchemeHandler_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineUrlSchemeHandler_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEngineView: public QWebEngineView
{
public:
	MyQWebEngineView(QWidget *parent = Q_NULLPTR) : QWebEngineView(parent) {QWebEngineView_QWebEngineView_QRegisterMetaType();};
	bool event(QEvent * vqe) { return callbackQWebEngineView_Event(this, vqe) != 0; };
	void back() { callbackQWebEngineView_Back(this); };
	void contextMenuEvent(QContextMenuEvent * vqc) { callbackQWebEngineView_ContextMenuEvent(this, vqc); };
	void dragEnterEvent(QDragEnterEvent * e) { callbackQWebEngineView_DragEnterEvent(this, e); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQWebEngineView_DragLeaveEvent(this, e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQWebEngineView_DragMoveEvent(this, e); };
	void dropEvent(QDropEvent * e) { callbackQWebEngineView_DropEvent(this, e); };
	void forward() { callbackQWebEngineView_Forward(this); };
	void hideEvent(QHideEvent * vqh) { callbackQWebEngineView_HideEvent(this, vqh); };
	void Signal_IconChanged(const QIcon & vqi) { callbackQWebEngineView_IconChanged(this, const_cast<QIcon*>(&vqi)); };
	void Signal_IconUrlChanged(const QUrl & vqu) { callbackQWebEngineView_IconUrlChanged(this, const_cast<QUrl*>(&vqu)); };
	void Signal_LoadFinished(bool vbo) { callbackQWebEngineView_LoadFinished(this, vbo); };
	void Signal_LoadProgress(int progress) { callbackQWebEngineView_LoadProgress(this, progress); };
	void Signal_LoadStarted() { callbackQWebEngineView_LoadStarted(this); };
	void reload() { callbackQWebEngineView_Reload(this); };
	void Signal_SelectionChanged() { callbackQWebEngineView_SelectionChanged(this); };
	void showEvent(QShowEvent * vqs) { callbackQWebEngineView_ShowEvent(this, vqs); };
	void stop() { callbackQWebEngineView_Stop(this); };
	void Signal_TitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebEngine_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebEngineView_TitleChanged(this, titlePacked); };
	void Signal_UrlChanged(const QUrl & vqu) { callbackQWebEngineView_UrlChanged(this, const_cast<QUrl*>(&vqu)); };
	 ~MyQWebEngineView() { callbackQWebEngineView_DestroyQWebEngineView(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWebEngineView_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebEngineView_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWebEngineView_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQWebEngineView_FocusNextPrevChild(this, next) != 0; };
	void actionEvent(QActionEvent * event) { callbackQWebEngineView_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQWebEngineView_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQWebEngineView_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQWebEngineView_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackQWebEngineView_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQWebEngineView_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQWebEngineView_FocusOutEvent(this, event); };
	void hide() { callbackQWebEngineView_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQWebEngineView_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWebEngineView_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWebEngineView_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQWebEngineView_LeaveEvent(this, event); };
	void lower() { callbackQWebEngineView_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWebEngineView_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQWebEngineView_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWebEngineView_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWebEngineView_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQWebEngineView_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQWebEngineView_PaintEvent(this, event); };
	void raise() { callbackQWebEngineView_Raise(this); };
	void repaint() { callbackQWebEngineView_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQWebEngineView_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWebEngineView_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWebEngineView_SetEnabled(this, vbo); };
	void setFocus() { callbackQWebEngineView_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQWebEngineView_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtWebEngine_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQWebEngineView_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQWebEngineView_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQWebEngineView_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtWebEngine_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQWebEngineView_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWebEngineView_Show(this); };
	void showFullScreen() { callbackQWebEngineView_ShowFullScreen(this); };
	void showMaximized() { callbackQWebEngineView_ShowMaximized(this); };
	void showMinimized() { callbackQWebEngineView_ShowMinimized(this); };
	void showNormal() { callbackQWebEngineView_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQWebEngineView_TabletEvent(this, event); };
	void update() { callbackQWebEngineView_Update(this); };
	void updateMicroFocus() { callbackQWebEngineView_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQWebEngineView_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQWebEngineView_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebEngine_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebEngineView_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWebEngineView_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQWebEngineView_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQWebEngineView_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQWebEngineView_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWebEngineView_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWebEngineView_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackQWebEngineView_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineView_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebEngineView_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineView_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineView_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineView_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineView_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineView_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineView_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineView_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineView*)

int QWebEngineView_QWebEngineView_QRegisterMetaType(){qRegisterMetaType<QWebEngineView*>(); return qRegisterMetaType<MyQWebEngineView*>();}

struct QtWebEngine_PackedString QWebEngineView_QWebEngineView_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t7fa76b = QWebEngineView::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t7fa76b.prepend("WHITESPACE").constData()+10), t7fa76b.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineView_QWebEngineView_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tec1013 = QWebEngineView::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tec1013.prepend("WHITESPACE").constData()+10), tec1013.size()-10 }; });
}

void* QWebEngineView_NewQWebEngineView(void* parent)
{
		return new MyQWebEngineView(static_cast<QWidget*>(parent));
}

char QWebEngineView_EventDefault(void* ptr, void* vqe)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::event(static_cast<QEvent*>(vqe));
}

void QWebEngineView_Back(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineView*>(ptr), "back");
}

void QWebEngineView_BackDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::back();
}

void QWebEngineView_ContextMenuEventDefault(void* ptr, void* vqc)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::contextMenuEvent(static_cast<QContextMenuEvent*>(vqc));
}

void QWebEngineView_DragEnterEventDefault(void* ptr, void* e)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void QWebEngineView_DragLeaveEventDefault(void* ptr, void* e)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QWebEngineView_DragMoveEventDefault(void* ptr, void* e)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QWebEngineView_DropEventDefault(void* ptr, void* e)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dropEvent(static_cast<QDropEvent*>(e));
}

void QWebEngineView_Forward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineView*>(ptr), "forward");
}

void QWebEngineView_ForwardDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::forward();
}

void QWebEngineView_HideEventDefault(void* ptr, void* vqh)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::hideEvent(static_cast<QHideEvent*>(vqh));
}

void QWebEngineView_ConnectIconChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QIcon &)>(&QWebEngineView::iconChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QIcon &)>(&MyQWebEngineView::Signal_IconChanged));
}

void QWebEngineView_DisconnectIconChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QIcon &)>(&QWebEngineView::iconChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QIcon &)>(&MyQWebEngineView::Signal_IconChanged));
}

void QWebEngineView_IconChanged(void* ptr, void* vqi)
{
	static_cast<QWebEngineView*>(ptr)->iconChanged(*static_cast<QIcon*>(vqi));
}

void QWebEngineView_ConnectIconUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QUrl &)>(&QWebEngineView::iconUrlChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QUrl &)>(&MyQWebEngineView::Signal_IconUrlChanged));
}

void QWebEngineView_DisconnectIconUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QUrl &)>(&QWebEngineView::iconUrlChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QUrl &)>(&MyQWebEngineView::Signal_IconUrlChanged));
}

void QWebEngineView_IconUrlChanged(void* ptr, void* vqu)
{
	static_cast<QWebEngineView*>(ptr)->iconUrlChanged(*static_cast<QUrl*>(vqu));
}

void QWebEngineView_Load(void* ptr, void* url)
{
	static_cast<QWebEngineView*>(ptr)->load(*static_cast<QUrl*>(url));
}

void QWebEngineView_Load2(void* ptr, void* request)
{
	static_cast<QWebEngineView*>(ptr)->load(*static_cast<QWebEngineHttpRequest*>(request));
}

void QWebEngineView_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(bool)>(&QWebEngineView::loadFinished), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(bool)>(&MyQWebEngineView::Signal_LoadFinished));
}

void QWebEngineView_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(bool)>(&QWebEngineView::loadFinished), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(bool)>(&MyQWebEngineView::Signal_LoadFinished));
}

void QWebEngineView_LoadFinished(void* ptr, char vbo)
{
	static_cast<QWebEngineView*>(ptr)->loadFinished(vbo != 0);
}

void QWebEngineView_ConnectLoadProgress(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(int)>(&QWebEngineView::loadProgress), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(int)>(&MyQWebEngineView::Signal_LoadProgress));
}

void QWebEngineView_DisconnectLoadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(int)>(&QWebEngineView::loadProgress), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(int)>(&MyQWebEngineView::Signal_LoadProgress));
}

void QWebEngineView_LoadProgress(void* ptr, int progress)
{
	static_cast<QWebEngineView*>(ptr)->loadProgress(progress);
}

void QWebEngineView_ConnectLoadStarted(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)()>(&QWebEngineView::loadStarted), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)()>(&MyQWebEngineView::Signal_LoadStarted));
}

void QWebEngineView_DisconnectLoadStarted(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)()>(&QWebEngineView::loadStarted), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)()>(&MyQWebEngineView::Signal_LoadStarted));
}

void QWebEngineView_LoadStarted(void* ptr)
{
	static_cast<QWebEngineView*>(ptr)->loadStarted();
}

void QWebEngineView_Reload(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineView*>(ptr), "reload");
}

void QWebEngineView_ReloadDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::reload();
}

void QWebEngineView_ConnectSelectionChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)()>(&QWebEngineView::selectionChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)()>(&MyQWebEngineView::Signal_SelectionChanged));
}

void QWebEngineView_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)()>(&QWebEngineView::selectionChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)()>(&MyQWebEngineView::Signal_SelectionChanged));
}

void QWebEngineView_SelectionChanged(void* ptr)
{
	static_cast<QWebEngineView*>(ptr)->selectionChanged();
}

void QWebEngineView_SetContent(void* ptr, void* data, struct QtWebEngine_PackedString mimeType, void* baseUrl)
{
	static_cast<QWebEngineView*>(ptr)->setContent(*static_cast<QByteArray*>(data), QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QUrl*>(baseUrl));
}

void QWebEngineView_SetHtml(void* ptr, struct QtWebEngine_PackedString html, void* baseUrl)
{
	static_cast<QWebEngineView*>(ptr)->setHtml(QString::fromUtf8(html.data, html.len), *static_cast<QUrl*>(baseUrl));
}

void QWebEngineView_SetPage(void* ptr, void* page)
{
	static_cast<QWebEngineView*>(ptr)->setPage(static_cast<QWebEnginePage*>(page));
}

void QWebEngineView_SetUrl(void* ptr, void* url)
{
	static_cast<QWebEngineView*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebEngineView_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebEngineView*>(ptr)->setZoomFactor(factor);
}

void QWebEngineView_ShowEventDefault(void* ptr, void* vqs)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::showEvent(static_cast<QShowEvent*>(vqs));
}

void QWebEngineView_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineView*>(ptr), "stop");
}

void QWebEngineView_StopDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::stop();
}

void QWebEngineView_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QString &)>(&QWebEngineView::titleChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QString &)>(&MyQWebEngineView::Signal_TitleChanged));
}

void QWebEngineView_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QString &)>(&QWebEngineView::titleChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QString &)>(&MyQWebEngineView::Signal_TitleChanged));
}

void QWebEngineView_TitleChanged(void* ptr, struct QtWebEngine_PackedString title)
{
	static_cast<QWebEngineView*>(ptr)->titleChanged(QString::fromUtf8(title.data, title.len));
}

void QWebEngineView_ConnectUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QUrl &)>(&QWebEngineView::urlChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QUrl &)>(&MyQWebEngineView::Signal_UrlChanged));
}

void QWebEngineView_DisconnectUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineView*>(ptr), static_cast<void (QWebEngineView::*)(const QUrl &)>(&QWebEngineView::urlChanged), static_cast<MyQWebEngineView*>(ptr), static_cast<void (MyQWebEngineView::*)(const QUrl &)>(&MyQWebEngineView::Signal_UrlChanged));
}

void QWebEngineView_UrlChanged(void* ptr, void* vqu)
{
	static_cast<QWebEngineView*>(ptr)->urlChanged(*static_cast<QUrl*>(vqu));
}

void QWebEngineView_DestroyQWebEngineView(void* ptr)
{
	static_cast<QWebEngineView*>(ptr)->~QWebEngineView();
}

void QWebEngineView_DestroyQWebEngineViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebEngineView_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebEngineView*>(ptr)->icon());
}

void* QWebEngineView_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebEngineView*>(ptr)->QWebEngineView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtWebEngine_PackedString QWebEngineView_SelectedText(void* ptr)
{
	return ({ QByteArray t0a8997 = static_cast<QWebEngineView*>(ptr)->selectedText().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t0a8997.prepend("WHITESPACE").constData()+10), t0a8997.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineView_Title(void* ptr)
{
	return ({ QByteArray t838611 = static_cast<QWebEngineView*>(ptr)->title().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t838611.prepend("WHITESPACE").constData()+10), t838611.size()-10 }; });
}

void* QWebEngineView_IconUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineView*>(ptr)->iconUrl());
}

void* QWebEngineView_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEngineView*>(ptr)->url());
}

void* QWebEngineView_Page(void* ptr)
{
	return static_cast<QWebEngineView*>(ptr)->page();
}

void* QWebEngineView_Settings(void* ptr)
{
	return static_cast<QWebEngineView*>(ptr)->settings();
}

char QWebEngineView_HasSelection(void* ptr)
{
	return static_cast<QWebEngineView*>(ptr)->hasSelection();
}

void* QWebEngineView_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebEngineView*>(ptr)->QWebEngineView::metaObject());
}

double QWebEngineView_ZoomFactor(void* ptr)
{
	return static_cast<QWebEngineView*>(ptr)->zoomFactor();
}

void* QWebEngineView___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebEngineView___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QWebEngineView___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebEngineView___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QWebEngineView___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebEngineView___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QWebEngineView___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineView___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineView___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineView___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineView___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineView___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineView___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineView___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineView___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineView___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineView___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebEngineView_CloseDefault(void* ptr)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::close();
}

char QWebEngineView_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::focusNextPrevChild(next != 0);
}

void QWebEngineView_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::actionEvent(static_cast<QActionEvent*>(event));
}

void QWebEngineView_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::changeEvent(static_cast<QEvent*>(event));
}

void QWebEngineView_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::closeEvent(static_cast<QCloseEvent*>(event));
}

void QWebEngineView_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::enterEvent(static_cast<QEvent*>(event));
}

void QWebEngineView_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QWebEngineView_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QWebEngineView_HideDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::hide();
}

void QWebEngineView_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QWebEngineView_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QWebEngineView_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QWebEngineView_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::leaveEvent(static_cast<QEvent*>(event));
}

void QWebEngineView_LowerDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::lower();
}

void QWebEngineView_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QWebEngineView_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QWebEngineView_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QWebEngineView_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QWebEngineView_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::moveEvent(static_cast<QMoveEvent*>(event));
}

void QWebEngineView_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::paintEvent(static_cast<QPaintEvent*>(event));
}

void QWebEngineView_RaiseDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::raise();
}

void QWebEngineView_RepaintDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::repaint();
}

void QWebEngineView_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QWebEngineView_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setDisabled(disable != 0);
}

void QWebEngineView_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setEnabled(vbo != 0);
}

void QWebEngineView_SetFocus2Default(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setFocus();
}

void QWebEngineView_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setHidden(hidden != 0);
}

void QWebEngineView_SetStyleSheetDefault(void* ptr, struct QtWebEngine_PackedString styleSheet)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QWebEngineView_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setVisible(visible != 0);
}

void QWebEngineView_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setWindowModified(vbo != 0);
}

void QWebEngineView_SetWindowTitleDefault(void* ptr, struct QtWebEngine_PackedString vqs)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QWebEngineView_ShowDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::show();
}

void QWebEngineView_ShowFullScreenDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::showFullScreen();
}

void QWebEngineView_ShowMaximizedDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::showMaximized();
}

void QWebEngineView_ShowMinimizedDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::showMinimized();
}

void QWebEngineView_ShowNormalDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::showNormal();
}

void QWebEngineView_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebEngineView_UpdateDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::update();
}

void QWebEngineView_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::updateMicroFocus();
}

void QWebEngineView_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QWebEngineView_PaintEngineDefault(void* ptr)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::paintEngine();
}

void* QWebEngineView_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebEngineView*>(ptr)->QWebEngineView::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWebEngineView_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QWebEngineView*>(ptr)->QWebEngineView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QWebEngineView_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::hasHeightForWidth();
}

int QWebEngineView_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::heightForWidth(w);
}

int QWebEngineView_MetricDefault(void* ptr, long long m)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QWebEngineView_InitPainterDefault(void* ptr, void* painter)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::initPainter(static_cast<QPainter*>(painter));
}

char QWebEngineView_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineView_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineView_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineView_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineView_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::deleteLater();
}

void QWebEngineView_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineView_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::timerEvent(static_cast<QTimerEvent*>(event));
}

void QtWebEngine_QtWebEngine_Initialize()
{
	QtWebEngine::initialize();
}

