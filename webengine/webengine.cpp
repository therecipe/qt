// +build !minimal

#define protected public
#define private public

#include "webengine.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
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
#include <QImage>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QNetworkCookie>
#include <QObject>
#include <QOffscreenSurface>
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
#include <QRemoteObjectPendingCallWatcher>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QSslCertificate>
#include <QSslKey>
#include <QString>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWebEngineCertificateError>
#include <QWebEngineClientCertificateStore>
#include <QWebEngineCookieStore>
#include <QWebEngineHttpRequest>
#include <QWebEngineNotification>
#include <QWebEnginePage>
#include <QWebEngineProfile>
#include <QWebEngineQuotaRequest>
#include <QWebEngineRegisterProtocolHandlerRequest>
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
	void Signal_DownloadPathChanged() { callbackQQuickWebEngineProfile_DownloadPathChanged(this); };
	void Signal_HttpAcceptLanguageChanged() { callbackQQuickWebEngineProfile_HttpAcceptLanguageChanged(this); };
	void Signal_HttpCacheMaximumSizeChanged() { callbackQQuickWebEngineProfile_HttpCacheMaximumSizeChanged(this); };
	void Signal_HttpCacheTypeChanged() { callbackQQuickWebEngineProfile_HttpCacheTypeChanged(this); };
	void Signal_HttpUserAgentChanged() { callbackQQuickWebEngineProfile_HttpUserAgentChanged(this); };
	void Signal_OffTheRecordChanged() { callbackQQuickWebEngineProfile_OffTheRecordChanged(this); };
	void Signal_PersistentCookiesPolicyChanged() { callbackQQuickWebEngineProfile_PersistentCookiesPolicyChanged(this); };
	void Signal_PersistentStoragePathChanged() { callbackQQuickWebEngineProfile_PersistentStoragePathChanged(this); };
	void Signal_PresentNotification(QWebEngineNotification * notification) { callbackQQuickWebEngineProfile_PresentNotification(this, notification); };
	void Signal_SpellCheckEnabledChanged() { callbackQQuickWebEngineProfile_SpellCheckEnabledChanged(this); };
	void Signal_SpellCheckLanguagesChanged() { callbackQQuickWebEngineProfile_SpellCheckLanguagesChanged(this); };
	void Signal_StorageNameChanged() { callbackQQuickWebEngineProfile_StorageNameChanged(this); };
	void Signal_UseForGlobalCertificateVerificationChanged() { callbackQQuickWebEngineProfile_UseForGlobalCertificateVerificationChanged(this); };
	void childEvent(QChildEvent * event) { callbackQQuickWebEngineProfile_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineProfile_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWebEngineProfile_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWebEngineProfile_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWebEngineProfile_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineProfile_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQuickWebEngineProfile_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWebEngineProfile_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickWebEngineProfile_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWebEngineProfile_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQuickWebEngineProfile*)

int QQuickWebEngineProfile_QQuickWebEngineProfile_QRegisterMetaType(){qRegisterMetaType<QQuickWebEngineProfile*>(); return qRegisterMetaType<MyQQuickWebEngineProfile*>();}

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
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineProfile(static_cast<QWindow*>(parent));
	} else {
		return new MyQQuickWebEngineProfile(static_cast<QObject*>(parent));
	}
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_CachePath(void* ptr)
{
	return ({ QByteArray t8a236f = static_cast<QQuickWebEngineProfile*>(ptr)->cachePath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t8a236f.prepend("WHITESPACE").constData()+10), t8a236f.size()-10 }; });
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

void* QQuickWebEngineProfile_ClientCertificateStore(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->clientCertificateStore();
}

void* QQuickWebEngineProfile_CookieStore(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->cookieStore();
}

void* QQuickWebEngineProfile_QQuickWebEngineProfile_DefaultProfile()
{
	return QQuickWebEngineProfile::defaultProfile();
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_DownloadPath(void* ptr)
{
	return ({ QByteArray td44cdc = static_cast<QQuickWebEngineProfile*>(ptr)->downloadPath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(td44cdc.prepend("WHITESPACE").constData()+10), td44cdc.size()-10 }; });
}

void QQuickWebEngineProfile_ConnectDownloadPathChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::downloadPathChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_DownloadPathChanged));
}

void QQuickWebEngineProfile_DisconnectDownloadPathChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::downloadPathChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_DownloadPathChanged));
}

void QQuickWebEngineProfile_DownloadPathChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->downloadPathChanged();
}

struct QtWebEngine_PackedString QQuickWebEngineProfile_HttpAcceptLanguage(void* ptr)
{
	return ({ QByteArray taf6551 = static_cast<QQuickWebEngineProfile*>(ptr)->httpAcceptLanguage().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(taf6551.prepend("WHITESPACE").constData()+10), taf6551.size()-10 }; });
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

int QQuickWebEngineProfile_HttpCacheMaximumSize(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->httpCacheMaximumSize();
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

long long QQuickWebEngineProfile_HttpCacheType(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->httpCacheType();
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

struct QtWebEngine_PackedString QQuickWebEngineProfile_HttpUserAgent(void* ptr)
{
	return ({ QByteArray t7ee5e9 = static_cast<QQuickWebEngineProfile*>(ptr)->httpUserAgent().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t7ee5e9.prepend("WHITESPACE").constData()+10), t7ee5e9.size()-10 }; });
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

char QQuickWebEngineProfile_IsOffTheRecord(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->isOffTheRecord();
}

char QQuickWebEngineProfile_IsSpellCheckEnabled(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->isSpellCheckEnabled();
}

char QQuickWebEngineProfile_IsUsedForGlobalCertificateVerification(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->isUsedForGlobalCertificateVerification();
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

long long QQuickWebEngineProfile_PersistentCookiesPolicy(void* ptr)
{
	return static_cast<QQuickWebEngineProfile*>(ptr)->persistentCookiesPolicy();
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

struct QtWebEngine_PackedString QQuickWebEngineProfile_PersistentStoragePath(void* ptr)
{
	return ({ QByteArray t5eccf6 = static_cast<QQuickWebEngineProfile*>(ptr)->persistentStoragePath().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t5eccf6.prepend("WHITESPACE").constData()+10), t5eccf6.size()-10 }; });
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

void QQuickWebEngineProfile_ConnectPresentNotification(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)(QWebEngineNotification *)>(&QQuickWebEngineProfile::presentNotification), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)(QWebEngineNotification *)>(&MyQQuickWebEngineProfile::Signal_PresentNotification));
}

void QQuickWebEngineProfile_DisconnectPresentNotification(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)(QWebEngineNotification *)>(&QQuickWebEngineProfile::presentNotification), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)(QWebEngineNotification *)>(&MyQQuickWebEngineProfile::Signal_PresentNotification));
}

void QQuickWebEngineProfile_PresentNotification(void* ptr, void* notification)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->presentNotification(static_cast<QWebEngineNotification*>(notification));
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

void QQuickWebEngineProfile_SetDownloadPath(void* ptr, struct QtWebEngine_PackedString path)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setDownloadPath(QString::fromUtf8(path.data, path.len));
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

void QQuickWebEngineProfile_SetSpellCheckEnabled(void* ptr, char enabled)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setSpellCheckEnabled(enabled != 0);
}

void QQuickWebEngineProfile_SetSpellCheckLanguages(void* ptr, struct QtWebEngine_PackedString languages)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setSpellCheckLanguages(QString::fromUtf8(languages.data, languages.len).split("¡¦!", QString::SkipEmptyParts));
}

void QQuickWebEngineProfile_SetStorageName(void* ptr, struct QtWebEngine_PackedString name)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setStorageName(QString::fromUtf8(name.data, name.len));
}

void QQuickWebEngineProfile_SetUrlRequestInterceptor(void* ptr, void* interceptor)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setUrlRequestInterceptor(static_cast<QWebEngineUrlRequestInterceptor*>(interceptor));
}

void QQuickWebEngineProfile_SetUseForGlobalCertificateVerification(void* ptr, char b)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->setUseForGlobalCertificateVerification(b != 0);
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

struct QtWebEngine_PackedString QQuickWebEngineProfile_SpellCheckLanguages(void* ptr)
{
	return ({ QByteArray td376a8 = static_cast<QQuickWebEngineProfile*>(ptr)->spellCheckLanguages().join("¡¦!").toUtf8(); QtWebEngine_PackedString { const_cast<char*>(td376a8.prepend("WHITESPACE").constData()+10), td376a8.size()-10 }; });
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

struct QtWebEngine_PackedString QQuickWebEngineProfile_StorageName(void* ptr)
{
	return ({ QByteArray t6af56e = static_cast<QQuickWebEngineProfile*>(ptr)->storageName().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t6af56e.prepend("WHITESPACE").constData()+10), t6af56e.size()-10 }; });
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

void* QQuickWebEngineProfile_UrlSchemeHandler(void* ptr, void* scheme)
{
	return const_cast<QWebEngineUrlSchemeHandler*>(static_cast<QQuickWebEngineProfile*>(ptr)->urlSchemeHandler(*static_cast<QByteArray*>(scheme)));
}

void QQuickWebEngineProfile_ConnectUseForGlobalCertificateVerificationChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::useForGlobalCertificateVerificationChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_UseForGlobalCertificateVerificationChanged));
}

void QQuickWebEngineProfile_DisconnectUseForGlobalCertificateVerificationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWebEngineProfile*>(ptr), static_cast<void (QQuickWebEngineProfile::*)()>(&QQuickWebEngineProfile::useForGlobalCertificateVerificationChanged), static_cast<MyQQuickWebEngineProfile*>(ptr), static_cast<void (MyQQuickWebEngineProfile::*)()>(&MyQQuickWebEngineProfile::Signal_UseForGlobalCertificateVerificationChanged));
}

void QQuickWebEngineProfile_UseForGlobalCertificateVerificationChanged(void* ptr)
{
	static_cast<QQuickWebEngineProfile*>(ptr)->useForGlobalCertificateVerificationChanged();
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

void* QQuickWebEngineProfile___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineProfile___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineProfile___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QQuickWebEngineProfile_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::event(static_cast<QEvent*>(e));
}

char QQuickWebEngineProfile_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQuickWebEngineProfile*>(ptr)->QQuickWebEngineProfile::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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
	void childEvent(QChildEvent * event) { callbackQQuickWebEngineScript_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineScript_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWebEngineScript_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWebEngineScript_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWebEngineScript_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWebEngineScript_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQuickWebEngineScript_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWebEngineScript_EventFilter(this, watched, event) != 0; };
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
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWebEngineScript(static_cast<QWindow*>(parent));
	} else {
		return new MyQQuickWebEngineScript(static_cast<QObject*>(parent));
	}
}

long long QQuickWebEngineScript_InjectionPoint(void* ptr)
{
	return static_cast<QQuickWebEngineScript*>(ptr)->injectionPoint();
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

struct QtWebEngine_PackedString QQuickWebEngineScript_Name(void* ptr)
{
	return ({ QByteArray te34dea = static_cast<QQuickWebEngineScript*>(ptr)->name().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(te34dea.prepend("WHITESPACE").constData()+10), te34dea.size()-10 }; });
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

char QQuickWebEngineScript_RunOnSubframes(void* ptr)
{
	return static_cast<QQuickWebEngineScript*>(ptr)->runOnSubframes();
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

struct QtWebEngine_PackedString QQuickWebEngineScript_SourceCode(void* ptr)
{
	return ({ QByteArray ta033c9 = static_cast<QQuickWebEngineScript*>(ptr)->sourceCode().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(ta033c9.prepend("WHITESPACE").constData()+10), ta033c9.size()-10 }; });
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

void* QQuickWebEngineScript_SourceUrl(void* ptr)
{
	return new QUrl(static_cast<QQuickWebEngineScript*>(ptr)->sourceUrl());
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

struct QtWebEngine_PackedString QQuickWebEngineScript_ToString(void* ptr)
{
	return ({ QByteArray t0f4281 = static_cast<QQuickWebEngineScript*>(ptr)->toString().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t0f4281.prepend("WHITESPACE").constData()+10), t0f4281.size()-10 }; });
}

long long QQuickWebEngineScript_WorldId(void* ptr)
{
	return static_cast<QQuickWebEngineScript*>(ptr)->worldId();
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

void* QQuickWebEngineScript___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQuickWebEngineScript___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQuickWebEngineScript___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QQuickWebEngineScript_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::event(static_cast<QEvent*>(e));
}

char QQuickWebEngineScript_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQuickWebEngineScript*>(ptr)->QQuickWebEngineScript::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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

void QWebEngineClientCertificateStore_Add(void* ptr, void* certificate, void* privateKey)
{
	static_cast<QWebEngineClientCertificateStore*>(ptr)->add(*static_cast<QSslCertificate*>(certificate), *static_cast<QSslKey*>(privateKey));
}

struct QtWebEngine_PackedList QWebEngineClientCertificateStore_Certificates(void* ptr)
{
	return ({ QVector<QSslCertificate>* tmpValue = new QVector<QSslCertificate>(static_cast<QWebEngineClientCertificateStore*>(ptr)->certificates()); QtWebEngine_PackedList { tmpValue, tmpValue->size() }; });
}

void QWebEngineClientCertificateStore_Clear(void* ptr)
{
	static_cast<QWebEngineClientCertificateStore*>(ptr)->clear();
}

void QWebEngineClientCertificateStore_Remove(void* ptr, void* certificate)
{
	static_cast<QWebEngineClientCertificateStore*>(ptr)->remove(*static_cast<QSslCertificate*>(certificate));
}

void* QWebEngineClientCertificateStore___certificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QVector<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QVector<QSslCertificate>*>(ptr)->size()-1) { static_cast<QVector<QSslCertificate>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QWebEngineClientCertificateStore___certificates_setList(void* ptr, void* i)
{
	static_cast<QVector<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QWebEngineClientCertificateStore___certificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QSslCertificate>();
}

class MyQWebEngineCookieStore: public QWebEngineCookieStore
{
public:
	void Signal_CookieAdded(const QNetworkCookie & cookie) { callbackQWebEngineCookieStore_CookieAdded(this, const_cast<QNetworkCookie*>(&cookie)); };
	void Signal_CookieRemoved(const QNetworkCookie & cookie) { callbackQWebEngineCookieStore_CookieRemoved(this, const_cast<QNetworkCookie*>(&cookie)); };
	 ~MyQWebEngineCookieStore() { callbackQWebEngineCookieStore_DestroyQWebEngineCookieStore(this); };
	void childEvent(QChildEvent * event) { callbackQWebEngineCookieStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineCookieStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineCookieStore_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineCookieStore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineCookieStore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineCookieStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEngineCookieStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineCookieStore_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineCookieStore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineCookieStore_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineCookieStore*)

int QWebEngineCookieStore_QWebEngineCookieStore_QRegisterMetaType(){qRegisterMetaType<QWebEngineCookieStore*>(); return qRegisterMetaType<MyQWebEngineCookieStore*>();}

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

void* QWebEngineCookieStore___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineCookieStore___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineCookieStore___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QWebEngineCookieStore_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::event(static_cast<QEvent*>(e));
}

char QWebEngineCookieStore_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineCookieStore_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineCookieStore*>(ptr)->QWebEngineCookieStore::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebEngineHttpRequest_NewQWebEngineHttpRequest(void* url, long long method)
{
	return new QWebEngineHttpRequest(*static_cast<QUrl*>(url), static_cast<QWebEngineHttpRequest::Method>(method));
}

void* QWebEngineHttpRequest_NewQWebEngineHttpRequest2(void* other)
{
	return new QWebEngineHttpRequest(*static_cast<QWebEngineHttpRequest*>(other));
}

char QWebEngineHttpRequest_HasHeader(void* ptr, void* headerName)
{
	return static_cast<QWebEngineHttpRequest*>(ptr)->hasHeader(*static_cast<QByteArray*>(headerName));
}

void* QWebEngineHttpRequest_Header(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QWebEngineHttpRequest*>(ptr)->header(*static_cast<QByteArray*>(headerName)));
}

struct QtWebEngine_PackedList QWebEngineHttpRequest_Headers(void* ptr)
{
	return ({ QVector<QByteArray>* tmpValue = new QVector<QByteArray>(static_cast<QWebEngineHttpRequest*>(ptr)->headers()); QtWebEngine_PackedList { tmpValue, tmpValue->size() }; });
}

long long QWebEngineHttpRequest_Method(void* ptr)
{
	return static_cast<QWebEngineHttpRequest*>(ptr)->method();
}

void* QWebEngineHttpRequest_PostData(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineHttpRequest*>(ptr)->postData());
}

void* QWebEngineHttpRequest_QWebEngineHttpRequest_PostRequest(void* url, void* postData)
{
	return new QWebEngineHttpRequest(QWebEngineHttpRequest::postRequest(*static_cast<QUrl*>(url), *static_cast<QMap<QString, QString>*>(postData)));
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

void* QWebEngineHttpRequest_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEngineHttpRequest*>(ptr)->url());
}

void QWebEngineHttpRequest_DestroyQWebEngineHttpRequest(void* ptr)
{
	static_cast<QWebEngineHttpRequest*>(ptr)->~QWebEngineHttpRequest();
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

class MyQWebEngineNotification: public QWebEngineNotification
{
public:
	void click() const { callbackQWebEngineNotification_Click(const_cast<void*>(static_cast<const void*>(this))); };
	void close() const { callbackQWebEngineNotification_Close(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_Closed() { callbackQWebEngineNotification_Closed(this); };
	void show() const { callbackQWebEngineNotification_Show(const_cast<void*>(static_cast<const void*>(this))); };
	void childEvent(QChildEvent * event) { callbackQWebEngineNotification_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineNotification_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineNotification_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineNotification_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineNotification_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineNotification_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEngineNotification_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineNotification_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineNotification_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineNotification_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineNotification*)

int QWebEngineNotification_QWebEngineNotification_QRegisterMetaType(){qRegisterMetaType<QWebEngineNotification*>(); return qRegisterMetaType<MyQWebEngineNotification*>();}

void QWebEngineNotification_Click(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineNotification*>(ptr), "click");
}

void QWebEngineNotification_ClickDefault(void* ptr)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::click();
}

void QWebEngineNotification_Close(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineNotification*>(ptr), "close");
}

void QWebEngineNotification_CloseDefault(void* ptr)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::close();
}

void QWebEngineNotification_ConnectClosed(void* ptr)
{
	QObject::connect(static_cast<QWebEngineNotification*>(ptr), static_cast<void (QWebEngineNotification::*)()>(&QWebEngineNotification::closed), static_cast<MyQWebEngineNotification*>(ptr), static_cast<void (MyQWebEngineNotification::*)()>(&MyQWebEngineNotification::Signal_Closed));
}

void QWebEngineNotification_DisconnectClosed(void* ptr)
{
	QObject::disconnect(static_cast<QWebEngineNotification*>(ptr), static_cast<void (QWebEngineNotification::*)()>(&QWebEngineNotification::closed), static_cast<MyQWebEngineNotification*>(ptr), static_cast<void (MyQWebEngineNotification::*)()>(&MyQWebEngineNotification::Signal_Closed));
}

void QWebEngineNotification_Closed(void* ptr)
{
	static_cast<QWebEngineNotification*>(ptr)->closed();
}

long long QWebEngineNotification_Direction(void* ptr)
{
	return static_cast<QWebEngineNotification*>(ptr)->direction();
}

void* QWebEngineNotification_Icon(void* ptr)
{
	return new QImage(static_cast<QWebEngineNotification*>(ptr)->icon());
}

struct QtWebEngine_PackedString QWebEngineNotification_Language(void* ptr)
{
	return ({ QByteArray t03088c = static_cast<QWebEngineNotification*>(ptr)->language().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t03088c.prepend("WHITESPACE").constData()+10), t03088c.size()-10 }; });
}

char QWebEngineNotification_Matches(void* ptr, void* other)
{
	return static_cast<QWebEngineNotification*>(ptr)->matches(static_cast<QWebEngineNotification*>(other));
}

struct QtWebEngine_PackedString QWebEngineNotification_Message(void* ptr)
{
	return ({ QByteArray t198a65 = static_cast<QWebEngineNotification*>(ptr)->message().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t198a65.prepend("WHITESPACE").constData()+10), t198a65.size()-10 }; });
}

void* QWebEngineNotification_Origin(void* ptr)
{
	return new QUrl(static_cast<QWebEngineNotification*>(ptr)->origin());
}

void QWebEngineNotification_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebEngineNotification*>(ptr), "show");
}

void QWebEngineNotification_ShowDefault(void* ptr)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::show();
}

struct QtWebEngine_PackedString QWebEngineNotification_Tag(void* ptr)
{
	return ({ QByteArray t352fbd = static_cast<QWebEngineNotification*>(ptr)->tag().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t352fbd.prepend("WHITESPACE").constData()+10), t352fbd.size()-10 }; });
}

struct QtWebEngine_PackedString QWebEngineNotification_Title(void* ptr)
{
	return ({ QByteArray tfe021d = static_cast<QWebEngineNotification*>(ptr)->title().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(tfe021d.prepend("WHITESPACE").constData()+10), tfe021d.size()-10 }; });
}

void* QWebEngineNotification___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineNotification___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineNotification___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWebEngineNotification___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineNotification___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineNotification___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebEngineNotification___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineNotification___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineNotification___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineNotification___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineNotification___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineNotification___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebEngineNotification___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineNotification___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineNotification___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QWebEngineNotification_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::childEvent(static_cast<QChildEvent*>(event));
}

void QWebEngineNotification_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebEngineNotification_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::customEvent(static_cast<QEvent*>(event));
}

void QWebEngineNotification_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::deleteLater();
}

void QWebEngineNotification_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWebEngineNotification_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::event(static_cast<QEvent*>(e));
}

char QWebEngineNotification_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineNotification_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineNotification*>(ptr)->QWebEngineNotification::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEnginePage: public QWebEnginePage
{
public:
	void Signal_AudioMutedChanged(bool muted) { callbackQWebEnginePage_AudioMutedChanged(this, muted); };
	void Signal_ContentsSizeChanged(const QSizeF & size) { callbackQWebEnginePage_ContentsSizeChanged(this, const_cast<QSizeF*>(&size)); };
	void Signal_IconChanged(const QIcon & icon) { callbackQWebEnginePage_IconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_IconUrlChanged(const QUrl & url) { callbackQWebEnginePage_IconUrlChanged(this, const_cast<QUrl*>(&url)); };
	void Signal_RecentlyAudibleChanged(bool recentlyAudible) { callbackQWebEnginePage_RecentlyAudibleChanged(this, recentlyAudible); };
	void Signal_ScrollPositionChanged(const QPointF & position) { callbackQWebEnginePage_ScrollPositionChanged(this, const_cast<QPointF*>(&position)); };
	void childEvent(QChildEvent * event) { callbackQWebEnginePage_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEnginePage_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEnginePage_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEnginePage_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEnginePage_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEnginePage_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEnginePage_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEnginePage_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEnginePage_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEnginePage_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEnginePage*)

int QWebEnginePage_QWebEnginePage_QRegisterMetaType(){qRegisterMetaType<QWebEnginePage*>(); return qRegisterMetaType<MyQWebEnginePage*>();}

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

void* QWebEnginePage_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QWebEnginePage*>(ptr)->backgroundColor());
}

void* QWebEnginePage_ContentsSize(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QWebEnginePage*>(ptr)->contentsSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
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

char QWebEnginePage_HasSelection(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->hasSelection();
}

void* QWebEnginePage_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebEnginePage*>(ptr)->icon());
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

void* QWebEnginePage_IconUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEnginePage*>(ptr)->iconUrl());
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

char QWebEnginePage_IsAudioMuted(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->isAudioMuted();
}

char QWebEnginePage_RecentlyAudible(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->recentlyAudible();
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

void* QWebEnginePage_RequestedUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEnginePage*>(ptr)->requestedUrl());
}

void* QWebEnginePage_ScrollPosition(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QWebEnginePage*>(ptr)->scrollPosition(); new QPointF(tmpValue.x(), tmpValue.y()); });
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

struct QtWebEngine_PackedString QWebEnginePage_SelectedText(void* ptr)
{
	return ({ QByteArray t729dfc = static_cast<QWebEnginePage*>(ptr)->selectedText().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t729dfc.prepend("WHITESPACE").constData()+10), t729dfc.size()-10 }; });
}

void QWebEnginePage_SetAudioMuted(void* ptr, char muted)
{
	static_cast<QWebEnginePage*>(ptr)->setAudioMuted(muted != 0);
}

void QWebEnginePage_SetBackgroundColor(void* ptr, void* color)
{
	static_cast<QWebEnginePage*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QWebEnginePage_SetUrl(void* ptr, void* url)
{
	static_cast<QWebEnginePage*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebEnginePage_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebEnginePage*>(ptr)->setZoomFactor(factor);
}

struct QtWebEngine_PackedString QWebEnginePage_Title(void* ptr)
{
	return ({ QByteArray t311622 = static_cast<QWebEnginePage*>(ptr)->title().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t311622.prepend("WHITESPACE").constData()+10), t311622.size()-10 }; });
}

void* QWebEnginePage_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEnginePage*>(ptr)->url());
}

double QWebEnginePage_ZoomFactor(void* ptr)
{
	return static_cast<QWebEnginePage*>(ptr)->zoomFactor();
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

void* QWebEnginePage___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEnginePage___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEnginePage___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QWebEnginePage_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::event(static_cast<QEvent*>(e));
}

char QWebEnginePage_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEnginePage_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEnginePage*>(ptr)->QWebEnginePage::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEngineProfile: public QWebEngineProfile
{
public:
	void childEvent(QChildEvent * event) { callbackQWebEngineProfile_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineProfile_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineProfile_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineProfile_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineProfile_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineProfile_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEngineProfile_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineProfile_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineProfile_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineProfile_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineProfile*)

int QWebEngineProfile_QWebEngineProfile_QRegisterMetaType(){qRegisterMetaType<QWebEngineProfile*>(); return qRegisterMetaType<MyQWebEngineProfile*>();}

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

void* QWebEngineProfile___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineProfile___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineProfile___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QWebEngineProfile_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::event(static_cast<QEvent*>(e));
}

char QWebEngineProfile_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineProfile_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineProfile*>(ptr)->QWebEngineProfile::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebEngineQuotaRequest_Accept(void* ptr)
{
	static_cast<QWebEngineQuotaRequest*>(ptr)->accept();
}

void* QWebEngineQuotaRequest_Origin(void* ptr)
{
	return new QUrl(static_cast<QWebEngineQuotaRequest*>(ptr)->origin());
}

void QWebEngineQuotaRequest_Reject(void* ptr)
{
	static_cast<QWebEngineQuotaRequest*>(ptr)->reject();
}

long long QWebEngineQuotaRequest_RequestedSize(void* ptr)
{
	return static_cast<QWebEngineQuotaRequest*>(ptr)->requestedSize();
}

void QWebEngineRegisterProtocolHandlerRequest_Accept(void* ptr)
{
	static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->accept();
}

void* QWebEngineRegisterProtocolHandlerRequest_Origin(void* ptr)
{
	return new QUrl(static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->origin());
}

void QWebEngineRegisterProtocolHandlerRequest_Reject(void* ptr)
{
	static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->reject();
}

struct QtWebEngine_PackedString QWebEngineRegisterProtocolHandlerRequest_Scheme(void* ptr)
{
	return ({ QByteArray t21b2e9 = static_cast<QWebEngineRegisterProtocolHandlerRequest*>(ptr)->scheme().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t21b2e9.prepend("WHITESPACE").constData()+10), t21b2e9.size()-10 }; });
}

void QWebEngineUrlRequestInfo_Block(void* ptr, char shouldBlock)
{
	static_cast<QWebEngineUrlRequestInfo*>(ptr)->block(shouldBlock != 0);
}

void* QWebEngineUrlRequestInfo_FirstPartyUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestInfo*>(ptr)->firstPartyUrl());
}

long long QWebEngineUrlRequestInfo_NavigationType(void* ptr)
{
	return static_cast<QWebEngineUrlRequestInfo*>(ptr)->navigationType();
}

void QWebEngineUrlRequestInfo_Redirect(void* ptr, void* url)
{
	static_cast<QWebEngineUrlRequestInfo*>(ptr)->redirect(*static_cast<QUrl*>(url));
}

void* QWebEngineUrlRequestInfo_RequestMethod(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineUrlRequestInfo*>(ptr)->requestMethod());
}

void* QWebEngineUrlRequestInfo_RequestUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestInfo*>(ptr)->requestUrl());
}

long long QWebEngineUrlRequestInfo_ResourceType(void* ptr)
{
	return static_cast<QWebEngineUrlRequestInfo*>(ptr)->resourceType();
}

void QWebEngineUrlRequestInfo_SetHttpHeader(void* ptr, void* name, void* value)
{
	static_cast<QWebEngineUrlRequestInfo*>(ptr)->setHttpHeader(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

class MyQWebEngineUrlRequestInterceptor: public QWebEngineUrlRequestInterceptor
{
public:
	MyQWebEngineUrlRequestInterceptor(QObject *p = Q_NULLPTR) : QWebEngineUrlRequestInterceptor(p) {QWebEngineUrlRequestInterceptor_QWebEngineUrlRequestInterceptor_QRegisterMetaType();};
	void interceptRequest(QWebEngineUrlRequestInfo & info) { callbackQWebEngineUrlRequestInterceptor_InterceptRequest(this, static_cast<QWebEngineUrlRequestInfo*>(&info)); };
	void childEvent(QChildEvent * event) { callbackQWebEngineUrlRequestInterceptor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestInterceptor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineUrlRequestInterceptor_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineUrlRequestInterceptor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineUrlRequestInterceptor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestInterceptor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEngineUrlRequestInterceptor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineUrlRequestInterceptor_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineUrlRequestInterceptor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineUrlRequestInterceptor_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineUrlRequestInterceptor*)

int QWebEngineUrlRequestInterceptor_QWebEngineUrlRequestInterceptor_QRegisterMetaType(){qRegisterMetaType<QWebEngineUrlRequestInterceptor*>(); return qRegisterMetaType<MyQWebEngineUrlRequestInterceptor*>();}

void* QWebEngineUrlRequestInterceptor_NewQWebEngineUrlRequestInterceptor2(void* p)
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
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(p))) {
		return new MyQWebEngineUrlRequestInterceptor(static_cast<QRemoteObjectPendingCallWatcher*>(p));
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

void* QWebEngineUrlRequestInterceptor___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestInterceptor___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestInterceptor___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QWebEngineUrlRequestInterceptor_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::event(static_cast<QEvent*>(e));
}

char QWebEngineUrlRequestInterceptor_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineUrlRequestInterceptor_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestInterceptor*>(ptr)->QWebEngineUrlRequestInterceptor::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEngineUrlRequestJob: public QWebEngineUrlRequestJob
{
public:
	void childEvent(QChildEvent * event) { callbackQWebEngineUrlRequestJob_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestJob_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineUrlRequestJob_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineUrlRequestJob_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineUrlRequestJob_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlRequestJob_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEngineUrlRequestJob_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineUrlRequestJob_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineUrlRequestJob_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineUrlRequestJob_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineUrlRequestJob*)

int QWebEngineUrlRequestJob_QWebEngineUrlRequestJob_QRegisterMetaType(){qRegisterMetaType<QWebEngineUrlRequestJob*>(); return qRegisterMetaType<MyQWebEngineUrlRequestJob*>();}

void QWebEngineUrlRequestJob_Fail(void* ptr, long long r)
{
	static_cast<QWebEngineUrlRequestJob*>(ptr)->fail(static_cast<QWebEngineUrlRequestJob::Error>(r));
}

void* QWebEngineUrlRequestJob_Initiator(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestJob*>(ptr)->initiator());
}

void QWebEngineUrlRequestJob_Redirect(void* ptr, void* url)
{
	static_cast<QWebEngineUrlRequestJob*>(ptr)->redirect(*static_cast<QUrl*>(url));
}

void QWebEngineUrlRequestJob_Reply(void* ptr, void* contentType, void* device)
{
	static_cast<QWebEngineUrlRequestJob*>(ptr)->reply(*static_cast<QByteArray*>(contentType), static_cast<QIODevice*>(device));
}

struct QtWebEngine_PackedList QWebEngineUrlRequestJob_RequestHeaders(void* ptr)
{
	return ({ QMap<QByteArray, QByteArray>* tmpValue = new QMap<QByteArray, QByteArray>(static_cast<QWebEngineUrlRequestJob*>(ptr)->requestHeaders()); QtWebEngine_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebEngineUrlRequestJob_RequestMethod(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineUrlRequestJob*>(ptr)->requestMethod());
}

void* QWebEngineUrlRequestJob_RequestUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineUrlRequestJob*>(ptr)->requestUrl());
}

void* QWebEngineUrlRequestJob___requestHeaders_atList(void* ptr, void* v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<QByteArray, QByteArray>*>(ptr)->value(*static_cast<QByteArray*>(v)); if (i == static_cast<QMap<QByteArray, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<QByteArray, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QWebEngineUrlRequestJob___requestHeaders_setList(void* ptr, void* key, void* i)
{
	static_cast<QMap<QByteArray, QByteArray>*>(ptr)->insert(*static_cast<QByteArray*>(key), *static_cast<QByteArray*>(i));
}

void* QWebEngineUrlRequestJob___requestHeaders_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QByteArray, QByteArray>();
}

struct QtWebEngine_PackedList QWebEngineUrlRequestJob___requestHeaders_keyList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QMap<QByteArray, QByteArray>*>(ptr)->keys()); QtWebEngine_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebEngineUrlRequestJob_____requestHeaders_keyList_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebEngineUrlRequestJob_____requestHeaders_keyList_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebEngineUrlRequestJob_____requestHeaders_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
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

void* QWebEngineUrlRequestJob___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlRequestJob___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlRequestJob___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QWebEngineUrlRequestJob_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::event(static_cast<QEvent*>(e));
}

char QWebEngineUrlRequestJob_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineUrlRequestJob_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlRequestJob*>(ptr)->QWebEngineUrlRequestJob::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme()
{
	return new QWebEngineUrlScheme();
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme2(void* name)
{
	return new QWebEngineUrlScheme(*static_cast<QByteArray*>(name));
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme3(void* that)
{
	return new QWebEngineUrlScheme(*static_cast<QWebEngineUrlScheme*>(that));
}

void* QWebEngineUrlScheme_NewQWebEngineUrlScheme4(void* that)
{
	return new QWebEngineUrlScheme(*static_cast<QWebEngineUrlScheme*>(that));
}

int QWebEngineUrlScheme_DefaultPort(void* ptr)
{
	return static_cast<QWebEngineUrlScheme*>(ptr)->defaultPort();
}

long long QWebEngineUrlScheme_Flags(void* ptr)
{
	return static_cast<QWebEngineUrlScheme*>(ptr)->flags();
}

void* QWebEngineUrlScheme_Name(void* ptr)
{
	return new QByteArray(static_cast<QWebEngineUrlScheme*>(ptr)->name());
}

void QWebEngineUrlScheme_QWebEngineUrlScheme_RegisterScheme(void* scheme)
{
	QWebEngineUrlScheme::registerScheme(*static_cast<QWebEngineUrlScheme*>(scheme));
}

void* QWebEngineUrlScheme_QWebEngineUrlScheme_SchemeByName(void* name)
{
	return new QWebEngineUrlScheme(QWebEngineUrlScheme::schemeByName(*static_cast<QByteArray*>(name)));
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

class MyQWebEngineUrlSchemeHandler: public QWebEngineUrlSchemeHandler
{
public:
	MyQWebEngineUrlSchemeHandler(QObject *parent = Q_NULLPTR) : QWebEngineUrlSchemeHandler(parent) {QWebEngineUrlSchemeHandler_QWebEngineUrlSchemeHandler_QRegisterMetaType();};
	void requestStarted(QWebEngineUrlRequestJob * request) { callbackQWebEngineUrlSchemeHandler_RequestStarted(this, request); };
	 ~MyQWebEngineUrlSchemeHandler() { callbackQWebEngineUrlSchemeHandler_DestroyQWebEngineUrlSchemeHandler(this); };
	void childEvent(QChildEvent * event) { callbackQWebEngineUrlSchemeHandler_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlSchemeHandler_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineUrlSchemeHandler_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineUrlSchemeHandler_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineUrlSchemeHandler_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineUrlSchemeHandler_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWebEngineUrlSchemeHandler_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineUrlSchemeHandler_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineUrlSchemeHandler_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineUrlSchemeHandler_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineUrlSchemeHandler*)

int QWebEngineUrlSchemeHandler_QWebEngineUrlSchemeHandler_QRegisterMetaType(){qRegisterMetaType<QWebEngineUrlSchemeHandler*>(); return qRegisterMetaType<MyQWebEngineUrlSchemeHandler*>();}

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
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebEngineUrlSchemeHandler(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
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

void* QWebEngineUrlSchemeHandler___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineUrlSchemeHandler___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineUrlSchemeHandler___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QWebEngineUrlSchemeHandler_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::event(static_cast<QEvent*>(e));
}

char QWebEngineUrlSchemeHandler_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineUrlSchemeHandler_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineUrlSchemeHandler*>(ptr)->QWebEngineUrlSchemeHandler::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQWebEngineView: public QWebEngineView
{
public:
	void Signal_IconChanged(const QIcon & vqi) { callbackQWebEngineView_IconChanged(this, const_cast<QIcon*>(&vqi)); };
	void Signal_IconUrlChanged(const QUrl & vqu) { callbackQWebEngineView_IconUrlChanged(this, const_cast<QUrl*>(&vqu)); };
	void actionEvent(QActionEvent * event) { callbackQWebEngineView_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQWebEngineView_ChangeEvent(this, event); };
	bool close() { return callbackQWebEngineView_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWebEngineView_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWebEngineView_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQWebEngineView_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQWebEngineView_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQWebEngineView_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQWebEngineView_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQWebEngineView_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQWebEngineView_EnterEvent(this, event); };
	bool event(QEvent * event) { return callbackQWebEngineView_Event(this, event) != 0; };
	void focusInEvent(QFocusEvent * event) { callbackQWebEngineView_FocusInEvent(this, event); };
	bool focusNextPrevChild(bool next) { return callbackQWebEngineView_FocusNextPrevChild(this, next) != 0; };
	void focusOutEvent(QFocusEvent * event) { callbackQWebEngineView_FocusOutEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWebEngineView_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWebEngineView_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWebEngineView_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWebEngineView_HideEvent(this, event); };
	void initPainter(QPainter * painter) const { callbackQWebEngineView_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQWebEngineView_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQWebEngineView_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQWebEngineView_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWebEngineView_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQWebEngineView_LeaveEvent(this, event); };
	void lower() { callbackQWebEngineView_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWebEngineView_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQWebEngineView_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWebEngineView_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQWebEngineView_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWebEngineView_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWebEngineView_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQWebEngineView_MoveEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQWebEngineView_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWebEngineView_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
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
	void showEvent(QShowEvent * event) { callbackQWebEngineView_ShowEvent(this, event); };
	void showFullScreen() { callbackQWebEngineView_ShowFullScreen(this); };
	void showMaximized() { callbackQWebEngineView_ShowMaximized(this); };
	void showMinimized() { callbackQWebEngineView_ShowMinimized(this); };
	void showNormal() { callbackQWebEngineView_ShowNormal(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWebEngineView_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void tabletEvent(QTabletEvent * event) { callbackQWebEngineView_TabletEvent(this, event); };
	void update() { callbackQWebEngineView_Update(this); };
	void updateMicroFocus() { callbackQWebEngineView_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQWebEngineView_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQWebEngineView_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebEngine_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebEngineView_WindowTitleChanged(this, titlePacked); };
	void childEvent(QChildEvent * event) { callbackQWebEngineView_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebEngineView_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebEngineView_CustomEvent(this, event); };
	void deleteLater() { callbackQWebEngineView_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebEngineView_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebEngineView_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebEngineView_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebEngine_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebEngineView_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebEngineView_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebEngineView*)

int QWebEngineView_QWebEngineView_QRegisterMetaType(){qRegisterMetaType<QWebEngineView*>(); return qRegisterMetaType<MyQWebEngineView*>();}

char QWebEngineView_HasSelection(void* ptr)
{
	return static_cast<QWebEngineView*>(ptr)->hasSelection();
}

void* QWebEngineView_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebEngineView*>(ptr)->icon());
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

void* QWebEngineView_IconUrl(void* ptr)
{
	return new QUrl(static_cast<QWebEngineView*>(ptr)->iconUrl());
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

struct QtWebEngine_PackedString QWebEngineView_SelectedText(void* ptr)
{
	return ({ QByteArray t0a8997 = static_cast<QWebEngineView*>(ptr)->selectedText().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t0a8997.prepend("WHITESPACE").constData()+10), t0a8997.size()-10 }; });
}

void QWebEngineView_SetUrl(void* ptr, void* url)
{
	static_cast<QWebEngineView*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebEngineView_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebEngineView*>(ptr)->setZoomFactor(factor);
}

struct QtWebEngine_PackedString QWebEngineView_Title(void* ptr)
{
	return ({ QByteArray t838611 = static_cast<QWebEngineView*>(ptr)->title().toUtf8(); QtWebEngine_PackedString { const_cast<char*>(t838611.prepend("WHITESPACE").constData()+10), t838611.size()-10 }; });
}

void* QWebEngineView_Url(void* ptr)
{
	return new QUrl(static_cast<QWebEngineView*>(ptr)->url());
}

double QWebEngineView_ZoomFactor(void* ptr)
{
	return static_cast<QWebEngineView*>(ptr)->zoomFactor();
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

void* QWebEngineView___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebEngineView___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebEngineView___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QWebEngineView_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::actionEvent(static_cast<QActionEvent*>(event));
}

void QWebEngineView_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::changeEvent(static_cast<QEvent*>(event));
}

char QWebEngineView_CloseDefault(void* ptr)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::close();
}

void QWebEngineView_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::closeEvent(static_cast<QCloseEvent*>(event));
}

void QWebEngineView_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QWebEngineView_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QWebEngineView_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QWebEngineView_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QWebEngineView_DropEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::dropEvent(static_cast<QDropEvent*>(event));
}

void QWebEngineView_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::enterEvent(static_cast<QEvent*>(event));
}

char QWebEngineView_EventDefault(void* ptr, void* event)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::event(static_cast<QEvent*>(event));
}

void QWebEngineView_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::focusInEvent(static_cast<QFocusEvent*>(event));
}

char QWebEngineView_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::focusNextPrevChild(next != 0);
}

void QWebEngineView_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::focusOutEvent(static_cast<QFocusEvent*>(event));
}

char QWebEngineView_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::hasHeightForWidth();
}

int QWebEngineView_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::heightForWidth(w);
}

void QWebEngineView_HideDefault(void* ptr)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::hide();
}

void QWebEngineView_HideEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::hideEvent(static_cast<QHideEvent*>(event));
}

void QWebEngineView_InitPainterDefault(void* ptr, void* painter)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::initPainter(static_cast<QPainter*>(painter));
}

void QWebEngineView_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QWebEngineView_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QWebEngineView*>(ptr)->QWebEngineView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
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

int QWebEngineView_MetricDefault(void* ptr, long long m)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QWebEngineView_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebEngineView*>(ptr)->QWebEngineView::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
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

char QWebEngineView_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void* QWebEngineView_PaintEngineDefault(void* ptr)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::paintEngine();
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

void QWebEngineView_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::showEvent(static_cast<QShowEvent*>(event));
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

void* QWebEngineView_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebEngineView*>(ptr)->QWebEngineView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
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

char QWebEngineView_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebEngineView*>(ptr)->QWebEngineView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebEngineView_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebEngineView*>(ptr)->QWebEngineView::timerEvent(static_cast<QTimerEvent*>(event));
}

void QtWebEngine_QtWebEngine_Initialize()
{
	QtWebEngine::initialize();
}

