// +build !minimal

#define protected public
#define private public

#include "network.h"
#include "_cgo_export.h"

#include <QAbstractNetworkCache>
#include <QAbstractSocket>
#include <QAuthenticator>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCryptographicHash>
#include <QDBusPendingCallWatcher>
#include <QDateTime>
#include <QDeadlineTimer>
#include <QDnsDomainNameRecord>
#include <QDnsHostAddressRecord>
#include <QDnsLookup>
#include <QDnsMailExchangeRecord>
#include <QDnsServiceRecord>
#include <QDnsTextRecord>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QHostAddress>
#include <QHostInfo>
#include <QHstsPolicy>
#include <QHttpMultiPart>
#include <QHttpPart>
#include <QIODevice>
#include <QLayout>
#include <QLocalServer>
#include <QLocalSocket>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMultiMap>
#include <QNetworkAccessManager>
#include <QNetworkAddressEntry>
#include <QNetworkCacheMetaData>
#include <QNetworkConfiguration>
#include <QNetworkConfigurationManager>
#include <QNetworkCookie>
#include <QNetworkCookieJar>
#include <QNetworkDatagram>
#include <QNetworkDiskCache>
#include <QNetworkInterface>
#include <QNetworkProxy>
#include <QNetworkProxyFactory>
#include <QNetworkProxyQuery>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QNetworkSession>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QRegExp>
#include <QSslCertificate>
#include <QSslCertificateExtension>
#include <QSslCipher>
#include <QSslConfiguration>
#include <QSslDiffieHellmanParameters>
#include <QSslEllipticCurve>
#include <QSslError>
#include <QSslKey>
#include <QSslPreSharedKeyAuthenticator>
#include <QSslSocket>
#include <QString>
#include <QTcpServer>
#include <QTcpSocket>
#include <QTimerEvent>
#include <QUdpSocket>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQAbstractNetworkCache: public QAbstractNetworkCache
{
public:
	MyQAbstractNetworkCache(QObject *parent = Q_NULLPTR) : QAbstractNetworkCache(parent) {QAbstractNetworkCache_QAbstractNetworkCache_QRegisterMetaType();};
	QIODevice * data(const QUrl & url) { return static_cast<QIODevice*>(callbackQAbstractNetworkCache_Data(this, const_cast<QUrl*>(&url))); };
	QIODevice * prepare(const QNetworkCacheMetaData & metaData) { return static_cast<QIODevice*>(callbackQAbstractNetworkCache_Prepare(this, const_cast<QNetworkCacheMetaData*>(&metaData))); };
	QNetworkCacheMetaData metaData(const QUrl & url) { return *static_cast<QNetworkCacheMetaData*>(callbackQAbstractNetworkCache_MetaData(this, const_cast<QUrl*>(&url))); };
	bool remove(const QUrl & url) { return callbackQAbstractNetworkCache_Remove(this, const_cast<QUrl*>(&url)) != 0; };
	void clear() { callbackQAbstractNetworkCache_Clear(this); };
	void insert(QIODevice * device) { callbackQAbstractNetworkCache_Insert(this, device); };
	void updateMetaData(const QNetworkCacheMetaData & metaData) { callbackQAbstractNetworkCache_UpdateMetaData(this, const_cast<QNetworkCacheMetaData*>(&metaData)); };
	 ~MyQAbstractNetworkCache() { callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractNetworkCache_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	qint64 cacheSize() const { return callbackQAbstractNetworkCache_CacheSize(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQAbstractNetworkCache_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractNetworkCache_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractNetworkCache_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractNetworkCache_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractNetworkCache_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractNetworkCache_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractNetworkCache_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractNetworkCache_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractNetworkCache*)

int QAbstractNetworkCache_QAbstractNetworkCache_QRegisterMetaType(){qRegisterMetaType<QAbstractNetworkCache*>(); return qRegisterMetaType<MyQAbstractNetworkCache*>();}

void* QAbstractNetworkCache_NewQAbstractNetworkCache(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractNetworkCache(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractNetworkCache(static_cast<QObject*>(parent));
	}
}

void* QAbstractNetworkCache_Data(void* ptr, void* url)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void* QAbstractNetworkCache_Prepare(void* ptr, void* metaData)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void* QAbstractNetworkCache_MetaData(void* ptr, void* url)
{
	return new QNetworkCacheMetaData(static_cast<QAbstractNetworkCache*>(ptr)->metaData(*static_cast<QUrl*>(url)));
}

struct QtNetwork_PackedString QAbstractNetworkCache_QAbstractNetworkCache_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t3472a0 = QAbstractNetworkCache::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t3472a0.prepend("WHITESPACE").constData()+10), t3472a0.size()-10 }; });
}

struct QtNetwork_PackedString QAbstractNetworkCache_QAbstractNetworkCache_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray ta19293 = QAbstractNetworkCache::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta19293.prepend("WHITESPACE").constData()+10), ta19293.size()-10 }; });
}

char QAbstractNetworkCache_Remove(void* ptr, void* url)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->remove(*static_cast<QUrl*>(url));
}

void QAbstractNetworkCache_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractNetworkCache*>(ptr), "clear");
}

void QAbstractNetworkCache_Insert(void* ptr, void* device)
{
	static_cast<QAbstractNetworkCache*>(ptr)->insert(static_cast<QIODevice*>(device));
}

void QAbstractNetworkCache_UpdateMetaData(void* ptr, void* metaData)
{
	static_cast<QAbstractNetworkCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_DestroyQAbstractNetworkCache(void* ptr)
{
	static_cast<QAbstractNetworkCache*>(ptr)->~QAbstractNetworkCache();
}

void QAbstractNetworkCache_DestroyQAbstractNetworkCacheDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractNetworkCache_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::metaObject());
	}
}

long long QAbstractNetworkCache_CacheSize(void* ptr)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->cacheSize();
}

void* QAbstractNetworkCache___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractNetworkCache___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractNetworkCache___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAbstractNetworkCache___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractNetworkCache___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractNetworkCache___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractNetworkCache___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractNetworkCache___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractNetworkCache___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractNetworkCache___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractNetworkCache___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractNetworkCache___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractNetworkCache___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractNetworkCache___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractNetworkCache___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QAbstractNetworkCache_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::event(static_cast<QEvent*>(e));
	}
}

char QAbstractNetworkCache_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QAbstractNetworkCache_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QAbstractNetworkCache_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractNetworkCache_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::customEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractNetworkCache_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::deleteLater();
	} else {
		static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::deleteLater();
	}
}

void QAbstractNetworkCache_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractNetworkCache_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QNetworkDiskCache*>(static_cast<QObject*>(ptr))) {
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQAbstractSocket: public QAbstractSocket
{
public:
	MyQAbstractSocket(QAbstractSocket::SocketType socketType, QObject *parent) : QAbstractSocket(socketType, parent) {QAbstractSocket_QAbstractSocket_QRegisterMetaType();};
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQAbstractSocket_SocketOption(this, option)); };
	bool waitForBytesWritten(int msecs) { return callbackQAbstractSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQAbstractSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQAbstractSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQAbstractSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQAbstractSocket_ReadData(this, dataPacked, maxSize); };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQAbstractSocket_ReadLineData(this, dataPacked, maxlen); };
	qint64 writeData(const char * data, qint64 size) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQAbstractSocket_WriteData(this, dataPacked, size); };
	void close() { callbackQAbstractSocket_Close(this); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQAbstractSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQAbstractSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void Signal_Connected() { callbackQAbstractSocket_Connected(this); };
	void disconnectFromHost() { callbackQAbstractSocket_DisconnectFromHost(this); };
	void Signal_Disconnected() { callbackQAbstractSocket_Disconnected(this); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocket_Error2(this, socketError); };
	void Signal_HostFound() { callbackQAbstractSocket_HostFound(this); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQAbstractSocket_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	void resume() { callbackQAbstractSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocket_SetReadBufferSize(this, size); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocket_StateChanged(this, socketState); };
	 ~MyQAbstractSocket() { callbackQAbstractSocket_DestroyQAbstractSocket(this); };
	bool atEnd() const { return callbackQAbstractSocket_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool canReadLine() const { return callbackQAbstractSocket_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isSequential() const { return callbackQAbstractSocket_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSocket_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	qint64 bytesAvailable() const { return callbackQAbstractSocket_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQAbstractSocket_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	bool open(QIODevice::OpenMode mode) { return callbackQAbstractSocket_Open(this, mode) != 0; };
	bool reset() { return callbackQAbstractSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQAbstractSocket_Seek(this, pos) != 0; };
	void Signal_AboutToClose() { callbackQAbstractSocket_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQAbstractSocket_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQAbstractSocket_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQAbstractSocket_ChannelReadyRead(this, channel); };
	void Signal_ReadChannelFinished() { callbackQAbstractSocket_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQAbstractSocket_ReadyRead(this); };
	qint64 pos() const { return callbackQAbstractSocket_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQAbstractSocket_Size(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQAbstractSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSocket_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSocket_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSocket_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSocket_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocket_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractSocket*)

int QAbstractSocket_QAbstractSocket_QRegisterMetaType(){qRegisterMetaType<QAbstractSocket*>(); return qRegisterMetaType<MyQAbstractSocket*>();}

void* QAbstractSocket_NewQAbstractSocket(long long socketType, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QAbstractSocket_QAbstractSocket_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tc30490 = QAbstractSocket::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc30490.prepend("WHITESPACE").constData()+10), tc30490.size()-10 }; });
}

struct QtNetwork_PackedString QAbstractSocket_QAbstractSocket_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t211326 = QAbstractSocket::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t211326.prepend("WHITESPACE").constData()+10), t211326.size()-10 }; });
}

void* QAbstractSocket_SocketOption(void* ptr, long long option)
{
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QAbstractSocket_SocketOptionDefault(void* ptr, long long option)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QUdpSocket*>(ptr)->QUdpSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSslSocket*>(ptr)->QSslSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QTcpSocket*>(ptr)->QTcpSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
	} else {
		return new QVariant(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
	}
}

char QAbstractSocket_Bind(void* ptr, void* address, unsigned short port, long long mode)
{
	return static_cast<QAbstractSocket*>(ptr)->bind(*static_cast<QHostAddress*>(address), port, static_cast<QAbstractSocket::BindFlag>(mode));
}

char QAbstractSocket_Bind2(void* ptr, unsigned short port, long long mode)
{
	return static_cast<QAbstractSocket*>(ptr)->bind(port, static_cast<QAbstractSocket::BindFlag>(mode));
}

char QAbstractSocket_Flush(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->flush();
}

char QAbstractSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForBytesWritten(msecs);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForBytesWritten(msecs);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForBytesWritten(msecs);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForBytesWritten(msecs);
	}
}

char QAbstractSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForConnected(msecs);
}

char QAbstractSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForConnected(msecs);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForConnected(msecs);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForConnected(msecs);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForConnected(msecs);
	}
}

char QAbstractSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForDisconnected(msecs);
}

char QAbstractSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForDisconnected(msecs);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForDisconnected(msecs);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForDisconnected(msecs);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForDisconnected(msecs);
	}
}

char QAbstractSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForReadyRead(msecs);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForReadyRead(msecs);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForReadyRead(msecs);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForReadyRead(msecs);
	}
}

long long QAbstractSocket_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QAbstractSocket*>(ptr)->readData(data, maxSize);
}

long long QAbstractSocket_ReadDataDefault(void* ptr, char* data, long long maxSize)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::readData(data, maxSize);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::readData(data, maxSize);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::readData(data, maxSize);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::readData(data, maxSize);
	}
}

long long QAbstractSocket_ReadLineDataDefault(void* ptr, char* data, long long maxlen)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::readLineData(data, maxlen);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::readLineData(data, maxlen);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::readLineData(data, maxlen);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::readLineData(data, maxlen);
	}
}

long long QAbstractSocket_WriteData(void* ptr, char* data, long long size)
{
	return static_cast<QAbstractSocket*>(ptr)->writeData(const_cast<const char*>(data), size);
}

long long QAbstractSocket_WriteDataDefault(void* ptr, char* data, long long size)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::writeData(const_cast<const char*>(data), size);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::writeData(const_cast<const char*>(data), size);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::writeData(const_cast<const char*>(data), size);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::writeData(const_cast<const char*>(data), size);
	}
}

void QAbstractSocket_Abort(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->abort();
}

void QAbstractSocket_CloseDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::close();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::close();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::close();
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::close();
	}
}

void QAbstractSocket_ConnectToHost2(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QAbstractSocket*>(ptr)->connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QAbstractSocket_ConnectToHost2Default(void* ptr, void* address, unsigned short port, long long openMode)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
	}
}

void QAbstractSocket_ConnectToHost(void* ptr, struct QtNetwork_PackedString hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QAbstractSocket*>(ptr)->connectToHost(QString::fromUtf8(hostName.data, hostName.len), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QAbstractSocket_ConnectToHostDefault(void* ptr, struct QtNetwork_PackedString hostName, unsigned short port, long long openMode, long long protocol)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::connectToHost(QString::fromUtf8(hostName.data, hostName.len), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::connectToHost(QString::fromUtf8(hostName.data, hostName.len), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::connectToHost(QString::fromUtf8(hostName.data, hostName.len), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::connectToHost(QString::fromUtf8(hostName.data, hostName.len), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
	}
}

void QAbstractSocket_ConnectConnected(void* ptr)
{
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));
}

void QAbstractSocket_DisconnectConnected(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));
}

void QAbstractSocket_Connected(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->connected();
}

void QAbstractSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->disconnectFromHost();
}

void QAbstractSocket_DisconnectFromHostDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::disconnectFromHost();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::disconnectFromHost();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::disconnectFromHost();
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::disconnectFromHost();
	}
}

void QAbstractSocket_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));
}

void QAbstractSocket_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));
}

void QAbstractSocket_Disconnected(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->disconnected();
}

void QAbstractSocket_ConnectError2(void* ptr)
{
	qRegisterMetaType<QAbstractSocket::SocketError>();
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::error), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketError)>(&MyQAbstractSocket::Signal_Error2));
}

void QAbstractSocket_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::error), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketError)>(&MyQAbstractSocket::Signal_Error2));
}

void QAbstractSocket_Error2(void* ptr, long long socketError)
{
	static_cast<QAbstractSocket*>(ptr)->error(static_cast<QAbstractSocket::SocketError>(socketError));
}

void QAbstractSocket_ConnectHostFound(void* ptr)
{
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));
}

void QAbstractSocket_DisconnectHostFound(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));
}

void QAbstractSocket_HostFound(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->hostFound();
}

void QAbstractSocket_ConnectProxyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&QAbstractSocket::proxyAuthenticationRequired), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&MyQAbstractSocket::Signal_ProxyAuthenticationRequired));
}

void QAbstractSocket_DisconnectProxyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&QAbstractSocket::proxyAuthenticationRequired), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&MyQAbstractSocket::Signal_ProxyAuthenticationRequired));
}

void QAbstractSocket_ProxyAuthenticationRequired(void* ptr, void* proxy, void* authenticator)
{
	static_cast<QAbstractSocket*>(ptr)->proxyAuthenticationRequired(*static_cast<QNetworkProxy*>(proxy), static_cast<QAuthenticator*>(authenticator));
}

void QAbstractSocket_Resume(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->resume();
}

void QAbstractSocket_ResumeDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::resume();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::resume();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::resume();
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::resume();
	}
}

void QAbstractSocket_SetLocalAddress(void* ptr, void* address)
{
	static_cast<QAbstractSocket*>(ptr)->setLocalAddress(*static_cast<QHostAddress*>(address));
}

void QAbstractSocket_SetLocalPort(void* ptr, unsigned short port)
{
	static_cast<QAbstractSocket*>(ptr)->setLocalPort(port);
}

void QAbstractSocket_SetPauseMode(void* ptr, long long pauseMode)
{
	static_cast<QAbstractSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QAbstractSocket_SetPeerAddress(void* ptr, void* address)
{
	static_cast<QAbstractSocket*>(ptr)->setPeerAddress(*static_cast<QHostAddress*>(address));
}

void QAbstractSocket_SetPeerName(void* ptr, struct QtNetwork_PackedString name)
{
	static_cast<QAbstractSocket*>(ptr)->setPeerName(QString::fromUtf8(name.data, name.len));
}

void QAbstractSocket_SetPeerPort(void* ptr, unsigned short port)
{
	static_cast<QAbstractSocket*>(ptr)->setPeerPort(port);
}

void QAbstractSocket_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QAbstractSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QAbstractSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QAbstractSocket*>(ptr)->setReadBufferSize(size);
}

void QAbstractSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::setReadBufferSize(size);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::setReadBufferSize(size);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::setReadBufferSize(size);
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setReadBufferSize(size);
	}
}

void QAbstractSocket_SetSocketError(void* ptr, long long socketError)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketError(static_cast<QAbstractSocket::SocketError>(socketError));
}

void QAbstractSocket_SetSocketOption(void* ptr, long long option, void* value)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QAbstractSocket_SetSocketOptionDefault(void* ptr, long long option, void* value)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
	}
}

void QAbstractSocket_SetSocketState(void* ptr, long long state)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketState(static_cast<QAbstractSocket::SocketState>(state));
}

void QAbstractSocket_ConnectStateChanged(void* ptr)
{
	qRegisterMetaType<QAbstractSocket::SocketState>();
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));
}

void QAbstractSocket_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));
}

void QAbstractSocket_StateChanged(void* ptr, long long socketState)
{
	static_cast<QAbstractSocket*>(ptr)->stateChanged(static_cast<QAbstractSocket::SocketState>(socketState));
}

void QAbstractSocket_DestroyQAbstractSocket(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->~QAbstractSocket();
}

void QAbstractSocket_DestroyQAbstractSocketDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstractSocket_PauseMode(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->pauseMode();
}

long long QAbstractSocket_Error(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->error();
}

long long QAbstractSocket_State(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->state();
}

long long QAbstractSocket_SocketType(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->socketType();
}

void* QAbstractSocket_LocalAddress(void* ptr)
{
	return new QHostAddress(static_cast<QAbstractSocket*>(ptr)->localAddress());
}

void* QAbstractSocket_PeerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QAbstractSocket*>(ptr)->peerAddress());
}

void* QAbstractSocket_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QAbstractSocket*>(ptr)->proxy());
}

struct QtNetwork_PackedString QAbstractSocket_PeerName(void* ptr)
{
	return ({ QByteArray tf7fe06 = static_cast<QAbstractSocket*>(ptr)->peerName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tf7fe06.prepend("WHITESPACE").constData()+10), tf7fe06.size()-10 }; });
}

char QAbstractSocket_AtEndDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::atEnd();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::atEnd();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::atEnd();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::atEnd();
	}
}

char QAbstractSocket_CanReadLineDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::canReadLine();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::canReadLine();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::canReadLine();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::canReadLine();
	}
}

char QAbstractSocket_IsSequentialDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::isSequential();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::isSequential();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::isSequential();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::isSequential();
	}
}

char QAbstractSocket_IsValid(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->isValid();
}

void* QAbstractSocket_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QUdpSocket*>(ptr)->QUdpSocket::metaObject());
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QSslSocket*>(ptr)->QSslSocket::metaObject());
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QTcpSocket*>(ptr)->QTcpSocket::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::metaObject());
	}
}

long long QAbstractSocket_BytesAvailableDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::bytesAvailable();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::bytesAvailable();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::bytesAvailable();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::bytesAvailable();
	}
}

long long QAbstractSocket_BytesToWriteDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::bytesToWrite();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::bytesToWrite();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::bytesToWrite();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::bytesToWrite();
	}
}

long long QAbstractSocket_ReadBufferSize(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->readBufferSize();
}

unsigned short QAbstractSocket_LocalPort(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->localPort();
}

unsigned short QAbstractSocket_PeerPort(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->peerPort();
}

void* QAbstractSocket___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractSocket___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractSocket___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAbstractSocket___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSocket___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSocket___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractSocket___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSocket___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSocket___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractSocket___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSocket___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSocket___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractSocket___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSocket___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSocket___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QAbstractSocket_OpenDefault(void* ptr, long long mode)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
	}
}

char QAbstractSocket_ResetDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::reset();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::reset();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::reset();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::reset();
	}
}

char QAbstractSocket_SeekDefault(void* ptr, long long pos)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::seek(pos);
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::seek(pos);
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::seek(pos);
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::seek(pos);
	}
}

long long QAbstractSocket_PosDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::pos();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::pos();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::pos();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::pos();
	}
}

long long QAbstractSocket_SizeDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::size();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::size();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::size();
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::size();
	}
}

char QAbstractSocket_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::event(static_cast<QEvent*>(e));
	}
}

char QAbstractSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QUdpSocket*>(ptr)->QUdpSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSslSocket*>(ptr)->QSslSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTcpSocket*>(ptr)->QTcpSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QAbstractSocket_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QAbstractSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractSocket_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::customEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractSocket_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::deleteLater();
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::deleteLater();
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::deleteLater();
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::deleteLater();
	}
}

void QAbstractSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractSocket_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QUdpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QUdpSocket*>(ptr)->QUdpSocket::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QSslSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QSslSocket*>(ptr)->QSslSocket::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTcpSocket*>(static_cast<QObject*>(ptr))) {
		static_cast<QTcpSocket*>(ptr)->QTcpSocket::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QAuthenticator_NewQAuthenticator()
{
	return new QAuthenticator();
}

void* QAuthenticator_NewQAuthenticator2(void* other)
{
	return new QAuthenticator(*static_cast<QAuthenticator*>(other));
}

void QAuthenticator_SetOption(void* ptr, struct QtNetwork_PackedString opt, void* value)
{
	static_cast<QAuthenticator*>(ptr)->setOption(QString::fromUtf8(opt.data, opt.len), *static_cast<QVariant*>(value));
}

void QAuthenticator_SetPassword(void* ptr, struct QtNetwork_PackedString password)
{
	static_cast<QAuthenticator*>(ptr)->setPassword(QString::fromUtf8(password.data, password.len));
}

void QAuthenticator_SetUser(void* ptr, struct QtNetwork_PackedString user)
{
	static_cast<QAuthenticator*>(ptr)->setUser(QString::fromUtf8(user.data, user.len));
}

void QAuthenticator_DestroyQAuthenticator(void* ptr)
{
	static_cast<QAuthenticator*>(ptr)->~QAuthenticator();
}

struct QtNetwork_PackedString QAuthenticator_Password(void* ptr)
{
	return ({ QByteArray t31072f = static_cast<QAuthenticator*>(ptr)->password().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t31072f.prepend("WHITESPACE").constData()+10), t31072f.size()-10 }; });
}

struct QtNetwork_PackedString QAuthenticator_Realm(void* ptr)
{
	return ({ QByteArray tcc1e3d = static_cast<QAuthenticator*>(ptr)->realm().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tcc1e3d.prepend("WHITESPACE").constData()+10), tcc1e3d.size()-10 }; });
}

struct QtNetwork_PackedString QAuthenticator_User(void* ptr)
{
	return ({ QByteArray ta76119 = static_cast<QAuthenticator*>(ptr)->user().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta76119.prepend("WHITESPACE").constData()+10), ta76119.size()-10 }; });
}

void* QAuthenticator_Option(void* ptr, struct QtNetwork_PackedString opt)
{
	return new QVariant(static_cast<QAuthenticator*>(ptr)->option(QString::fromUtf8(opt.data, opt.len)));
}

struct QtNetwork_PackedList QAuthenticator_Options(void* ptr)
{
	return ({ QHash<QString, QVariant>* tmpValue = new QHash<QString, QVariant>(static_cast<QAuthenticator*>(ptr)->options()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QAuthenticator_IsNull(void* ptr)
{
	return static_cast<QAuthenticator*>(ptr)->isNull();
}

void* QAuthenticator___options_atList(void* ptr, struct QtNetwork_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAuthenticator___options_setList(void* ptr, struct QtNetwork_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QAuthenticator___options_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtNetwork_PackedList QAuthenticator___options_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedString QAuthenticator_____options_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QAuthenticator_____options_keyList_setList(void* ptr, struct QtNetwork_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QAuthenticator_____options_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord()
{
	return new QDnsDomainNameRecord();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord2(void* other)
{
	return new QDnsDomainNameRecord(*static_cast<QDnsDomainNameRecord*>(other));
}

void QDnsDomainNameRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsDomainNameRecord*>(ptr)->swap(*static_cast<QDnsDomainNameRecord*>(other));
}

void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(void* ptr)
{
	static_cast<QDnsDomainNameRecord*>(ptr)->~QDnsDomainNameRecord();
}

struct QtNetwork_PackedString QDnsDomainNameRecord_Name(void* ptr)
{
	return ({ QByteArray t22074d = static_cast<QDnsDomainNameRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t22074d.prepend("WHITESPACE").constData()+10), t22074d.size()-10 }; });
}

struct QtNetwork_PackedString QDnsDomainNameRecord_Value(void* ptr)
{
	return ({ QByteArray tb334d2 = static_cast<QDnsDomainNameRecord*>(ptr)->value().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb334d2.prepend("WHITESPACE").constData()+10), tb334d2.size()-10 }; });
}

unsigned int QDnsDomainNameRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsDomainNameRecord*>(ptr)->timeToLive();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord()
{
	return new QDnsHostAddressRecord();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord2(void* other)
{
	return new QDnsHostAddressRecord(*static_cast<QDnsHostAddressRecord*>(other));
}

void QDnsHostAddressRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsHostAddressRecord*>(ptr)->swap(*static_cast<QDnsHostAddressRecord*>(other));
}

void QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(void* ptr)
{
	static_cast<QDnsHostAddressRecord*>(ptr)->~QDnsHostAddressRecord();
}

void* QDnsHostAddressRecord_Value(void* ptr)
{
	return new QHostAddress(static_cast<QDnsHostAddressRecord*>(ptr)->value());
}

struct QtNetwork_PackedString QDnsHostAddressRecord_Name(void* ptr)
{
	return ({ QByteArray tb52211 = static_cast<QDnsHostAddressRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb52211.prepend("WHITESPACE").constData()+10), tb52211.size()-10 }; });
}

unsigned int QDnsHostAddressRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsHostAddressRecord*>(ptr)->timeToLive();
}

class MyQDnsLookup: public QDnsLookup
{
public:
	MyQDnsLookup(QDnsLookup::Type ty, const QString &name, QObject *parent = Q_NULLPTR) : QDnsLookup(ty, name, parent) {QDnsLookup_QDnsLookup_QRegisterMetaType();};
	MyQDnsLookup(QDnsLookup::Type ty, const QString &name, const QHostAddress &nameserver, QObject *parent = Q_NULLPTR) : QDnsLookup(ty, name, nameserver, parent) {QDnsLookup_QDnsLookup_QRegisterMetaType();};
	MyQDnsLookup(QObject *parent = Q_NULLPTR) : QDnsLookup(parent) {QDnsLookup_QDnsLookup_QRegisterMetaType();};
	void abort() { callbackQDnsLookup_Abort(this); };
	void Signal_Finished() { callbackQDnsLookup_Finished(this); };
	void lookup() { callbackQDnsLookup_Lookup(this); };
	void Signal_NameChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtNetwork_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDnsLookup_NameChanged(this, namePacked); };
	void Signal_NameserverChanged(const QHostAddress & nameserver) { callbackQDnsLookup_NameserverChanged(this, const_cast<QHostAddress*>(&nameserver)); };
	void Signal_TypeChanged(QDnsLookup::Type ty) { callbackQDnsLookup_TypeChanged(this, ty); };
	 ~MyQDnsLookup() { callbackQDnsLookup_DestroyQDnsLookup(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDnsLookup_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQDnsLookup_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDnsLookup_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDnsLookup_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDnsLookup_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDnsLookup_CustomEvent(this, event); };
	void deleteLater() { callbackQDnsLookup_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDnsLookup_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDnsLookup_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDnsLookup_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDnsLookup_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQDnsLookup*)

int QDnsLookup_QDnsLookup_QRegisterMetaType(){qRegisterMetaType<QDnsLookup*>(); return qRegisterMetaType<MyQDnsLookup*>();}

void* QDnsLookup_NewQDnsLookup2(long long ty, struct QtNetwork_PackedString name, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), static_cast<QObject*>(parent));
	}
}

void* QDnsLookup_NewQDnsLookup3(long long ty, struct QtNetwork_PackedString name, void* nameserver, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QWindow*>(parent));
	} else {
		return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString::fromUtf8(name.data, name.len), *static_cast<QHostAddress*>(nameserver), static_cast<QObject*>(parent));
	}
}

void* QDnsLookup_NewQDnsLookup(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDnsLookup(static_cast<QWindow*>(parent));
	} else {
		return new MyQDnsLookup(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QDnsLookup_QDnsLookup_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t935ed6 = QDnsLookup::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t935ed6.prepend("WHITESPACE").constData()+10), t935ed6.size()-10 }; });
}

struct QtNetwork_PackedString QDnsLookup_QDnsLookup_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tceb8c4 = QDnsLookup::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(tceb8c4.prepend("WHITESPACE").constData()+10), tceb8c4.size()-10 }; });
}

void QDnsLookup_Abort(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "abort");
}

void QDnsLookup_AbortDefault(void* ptr)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::abort();
}

void QDnsLookup_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));
}

void QDnsLookup_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));
}

void QDnsLookup_Finished(void* ptr)
{
	static_cast<QDnsLookup*>(ptr)->finished();
}

void QDnsLookup_Lookup(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "lookup");
}

void QDnsLookup_LookupDefault(void* ptr)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::lookup();
}

void QDnsLookup_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));
}

void QDnsLookup_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));
}

void QDnsLookup_NameChanged(void* ptr, struct QtNetwork_PackedString name)
{
	static_cast<QDnsLookup*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

void QDnsLookup_ConnectNameserverChanged(void* ptr)
{
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QHostAddress &)>(&QDnsLookup::nameserverChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QHostAddress &)>(&MyQDnsLookup::Signal_NameserverChanged));
}

void QDnsLookup_DisconnectNameserverChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QHostAddress &)>(&QDnsLookup::nameserverChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QHostAddress &)>(&MyQDnsLookup::Signal_NameserverChanged));
}

void QDnsLookup_NameserverChanged(void* ptr, void* nameserver)
{
	static_cast<QDnsLookup*>(ptr)->nameserverChanged(*static_cast<QHostAddress*>(nameserver));
}

void QDnsLookup_SetName(void* ptr, struct QtNetwork_PackedString name)
{
	static_cast<QDnsLookup*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QDnsLookup_SetNameserver(void* ptr, void* nameserver)
{
	static_cast<QDnsLookup*>(ptr)->setNameserver(*static_cast<QHostAddress*>(nameserver));
}

void QDnsLookup_SetType(void* ptr, long long vqd)
{
	static_cast<QDnsLookup*>(ptr)->setType(static_cast<QDnsLookup::Type>(vqd));
}

void QDnsLookup_ConnectTypeChanged(void* ptr)
{
	qRegisterMetaType<QDnsLookup::Type>();
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));
}

void QDnsLookup_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));
}

void QDnsLookup_TypeChanged(void* ptr, long long ty)
{
	static_cast<QDnsLookup*>(ptr)->typeChanged(static_cast<QDnsLookup::Type>(ty));
}

void QDnsLookup_DestroyQDnsLookup(void* ptr)
{
	static_cast<QDnsLookup*>(ptr)->~QDnsLookup();
}

void QDnsLookup_DestroyQDnsLookupDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QDnsLookup_Error(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->error();
}

long long QDnsLookup_Type(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->type();
}

void* QDnsLookup_Nameserver(void* ptr)
{
	return new QHostAddress(static_cast<QDnsLookup*>(ptr)->nameserver());
}

struct QtNetwork_PackedList QDnsLookup_CanonicalNameRecords(void* ptr)
{
	return ({ QList<QDnsDomainNameRecord>* tmpValue = new QList<QDnsDomainNameRecord>(static_cast<QDnsLookup*>(ptr)->canonicalNameRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_NameServerRecords(void* ptr)
{
	return ({ QList<QDnsDomainNameRecord>* tmpValue = new QList<QDnsDomainNameRecord>(static_cast<QDnsLookup*>(ptr)->nameServerRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_PointerRecords(void* ptr)
{
	return ({ QList<QDnsDomainNameRecord>* tmpValue = new QList<QDnsDomainNameRecord>(static_cast<QDnsLookup*>(ptr)->pointerRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_HostAddressRecords(void* ptr)
{
	return ({ QList<QDnsHostAddressRecord>* tmpValue = new QList<QDnsHostAddressRecord>(static_cast<QDnsLookup*>(ptr)->hostAddressRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_MailExchangeRecords(void* ptr)
{
	return ({ QList<QDnsMailExchangeRecord>* tmpValue = new QList<QDnsMailExchangeRecord>(static_cast<QDnsLookup*>(ptr)->mailExchangeRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_ServiceRecords(void* ptr)
{
	return ({ QList<QDnsServiceRecord>* tmpValue = new QList<QDnsServiceRecord>(static_cast<QDnsLookup*>(ptr)->serviceRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_TextRecords(void* ptr)
{
	return ({ QList<QDnsTextRecord>* tmpValue = new QList<QDnsTextRecord>(static_cast<QDnsLookup*>(ptr)->textRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedString QDnsLookup_ErrorString(void* ptr)
{
	return ({ QByteArray ta68e7f = static_cast<QDnsLookup*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta68e7f.prepend("WHITESPACE").constData()+10), ta68e7f.size()-10 }; });
}

struct QtNetwork_PackedString QDnsLookup_Name(void* ptr)
{
	return ({ QByteArray td8dec7 = static_cast<QDnsLookup*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(td8dec7.prepend("WHITESPACE").constData()+10), td8dec7.size()-10 }; });
}

char QDnsLookup_IsFinished(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->isFinished();
}

void* QDnsLookup_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDnsLookup*>(ptr)->QDnsLookup::metaObject());
}

void* QDnsLookup___canonicalNameRecords_atList(void* ptr, int i)
{
	return new QDnsDomainNameRecord(({QDnsDomainNameRecord tmp = static_cast<QList<QDnsDomainNameRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsDomainNameRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsDomainNameRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___canonicalNameRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsDomainNameRecord>*>(ptr)->append(*static_cast<QDnsDomainNameRecord*>(i));
}

void* QDnsLookup___canonicalNameRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsDomainNameRecord>();
}

void* QDnsLookup___nameServerRecords_atList(void* ptr, int i)
{
	return new QDnsDomainNameRecord(({QDnsDomainNameRecord tmp = static_cast<QList<QDnsDomainNameRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsDomainNameRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsDomainNameRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___nameServerRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsDomainNameRecord>*>(ptr)->append(*static_cast<QDnsDomainNameRecord*>(i));
}

void* QDnsLookup___nameServerRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsDomainNameRecord>();
}

void* QDnsLookup___pointerRecords_atList(void* ptr, int i)
{
	return new QDnsDomainNameRecord(({QDnsDomainNameRecord tmp = static_cast<QList<QDnsDomainNameRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsDomainNameRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsDomainNameRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___pointerRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsDomainNameRecord>*>(ptr)->append(*static_cast<QDnsDomainNameRecord*>(i));
}

void* QDnsLookup___pointerRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsDomainNameRecord>();
}

void* QDnsLookup___hostAddressRecords_atList(void* ptr, int i)
{
	return new QDnsHostAddressRecord(({QDnsHostAddressRecord tmp = static_cast<QList<QDnsHostAddressRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsHostAddressRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsHostAddressRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___hostAddressRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsHostAddressRecord>*>(ptr)->append(*static_cast<QDnsHostAddressRecord*>(i));
}

void* QDnsLookup___hostAddressRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsHostAddressRecord>();
}

void* QDnsLookup___mailExchangeRecords_atList(void* ptr, int i)
{
	return new QDnsMailExchangeRecord(({QDnsMailExchangeRecord tmp = static_cast<QList<QDnsMailExchangeRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsMailExchangeRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsMailExchangeRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___mailExchangeRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsMailExchangeRecord>*>(ptr)->append(*static_cast<QDnsMailExchangeRecord*>(i));
}

void* QDnsLookup___mailExchangeRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsMailExchangeRecord>();
}

void* QDnsLookup___serviceRecords_atList(void* ptr, int i)
{
	return new QDnsServiceRecord(({QDnsServiceRecord tmp = static_cast<QList<QDnsServiceRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsServiceRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsServiceRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___serviceRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsServiceRecord>*>(ptr)->append(*static_cast<QDnsServiceRecord*>(i));
}

void* QDnsLookup___serviceRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsServiceRecord>();
}

void* QDnsLookup___textRecords_atList(void* ptr, int i)
{
	return new QDnsTextRecord(({QDnsTextRecord tmp = static_cast<QList<QDnsTextRecord>*>(ptr)->at(i); if (i == static_cast<QList<QDnsTextRecord>*>(ptr)->size()-1) { static_cast<QList<QDnsTextRecord>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___textRecords_setList(void* ptr, void* i)
{
	static_cast<QList<QDnsTextRecord>*>(ptr)->append(*static_cast<QDnsTextRecord*>(i));
}

void* QDnsLookup___textRecords_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDnsTextRecord>();
}

void* QDnsLookup___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsLookup___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDnsLookup___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QDnsLookup___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDnsLookup___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QDnsLookup___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QDnsLookup___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDnsLookup___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QDnsLookup___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QDnsLookup___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDnsLookup___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QDnsLookup___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QDnsLookup___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDnsLookup___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QDnsLookup___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QDnsLookup_EventDefault(void* ptr, void* e)
{
		return static_cast<QDnsLookup*>(ptr)->QDnsLookup::event(static_cast<QEvent*>(e));
}

char QDnsLookup_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QDnsLookup*>(ptr)->QDnsLookup::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QDnsLookup_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::childEvent(static_cast<QChildEvent*>(event));
}

void QDnsLookup_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDnsLookup_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::customEvent(static_cast<QEvent*>(event));
}

void QDnsLookup_DeleteLaterDefault(void* ptr)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::deleteLater();
}

void QDnsLookup_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDnsLookup_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDnsLookup*>(ptr)->QDnsLookup::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord()
{
	return new QDnsMailExchangeRecord();
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(void* other)
{
	return new QDnsMailExchangeRecord(*static_cast<QDnsMailExchangeRecord*>(other));
}

void QDnsMailExchangeRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsMailExchangeRecord*>(ptr)->swap(*static_cast<QDnsMailExchangeRecord*>(other));
}

void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(void* ptr)
{
	static_cast<QDnsMailExchangeRecord*>(ptr)->~QDnsMailExchangeRecord();
}

struct QtNetwork_PackedString QDnsMailExchangeRecord_Exchange(void* ptr)
{
	return ({ QByteArray t48794c = static_cast<QDnsMailExchangeRecord*>(ptr)->exchange().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t48794c.prepend("WHITESPACE").constData()+10), t48794c.size()-10 }; });
}

struct QtNetwork_PackedString QDnsMailExchangeRecord_Name(void* ptr)
{
	return ({ QByteArray t60d75b = static_cast<QDnsMailExchangeRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t60d75b.prepend("WHITESPACE").constData()+10), t60d75b.size()-10 }; });
}

unsigned short QDnsMailExchangeRecord_Preference(void* ptr)
{
	return static_cast<QDnsMailExchangeRecord*>(ptr)->preference();
}

unsigned int QDnsMailExchangeRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsMailExchangeRecord*>(ptr)->timeToLive();
}

void* QDnsServiceRecord_NewQDnsServiceRecord()
{
	return new QDnsServiceRecord();
}

void* QDnsServiceRecord_NewQDnsServiceRecord2(void* other)
{
	return new QDnsServiceRecord(*static_cast<QDnsServiceRecord*>(other));
}

void QDnsServiceRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsServiceRecord*>(ptr)->swap(*static_cast<QDnsServiceRecord*>(other));
}

void QDnsServiceRecord_DestroyQDnsServiceRecord(void* ptr)
{
	static_cast<QDnsServiceRecord*>(ptr)->~QDnsServiceRecord();
}

struct QtNetwork_PackedString QDnsServiceRecord_Name(void* ptr)
{
	return ({ QByteArray t9ceb61 = static_cast<QDnsServiceRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t9ceb61.prepend("WHITESPACE").constData()+10), t9ceb61.size()-10 }; });
}

struct QtNetwork_PackedString QDnsServiceRecord_Target(void* ptr)
{
	return ({ QByteArray te2873b = static_cast<QDnsServiceRecord*>(ptr)->target().toUtf8(); QtNetwork_PackedString { const_cast<char*>(te2873b.prepend("WHITESPACE").constData()+10), te2873b.size()-10 }; });
}

unsigned short QDnsServiceRecord_Port(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->port();
}

unsigned short QDnsServiceRecord_Priority(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->priority();
}

unsigned short QDnsServiceRecord_Weight(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->weight();
}

unsigned int QDnsServiceRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->timeToLive();
}

void* QDnsTextRecord_NewQDnsTextRecord()
{
	return new QDnsTextRecord();
}

void* QDnsTextRecord_NewQDnsTextRecord2(void* other)
{
	return new QDnsTextRecord(*static_cast<QDnsTextRecord*>(other));
}

void QDnsTextRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsTextRecord*>(ptr)->swap(*static_cast<QDnsTextRecord*>(other));
}

void QDnsTextRecord_DestroyQDnsTextRecord(void* ptr)
{
	static_cast<QDnsTextRecord*>(ptr)->~QDnsTextRecord();
}

struct QtNetwork_PackedList QDnsTextRecord_Values(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QDnsTextRecord*>(ptr)->values()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedString QDnsTextRecord_Name(void* ptr)
{
	return ({ QByteArray tb66dc5 = static_cast<QDnsTextRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb66dc5.prepend("WHITESPACE").constData()+10), tb66dc5.size()-10 }; });
}

unsigned int QDnsTextRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsTextRecord*>(ptr)->timeToLive();
}

void* QDnsTextRecord___values_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDnsTextRecord___values_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDnsTextRecord___values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QHostAddress_NewQHostAddress()
{
	return new QHostAddress();
}

void* QHostAddress_NewQHostAddress9(long long address)
{
	return new QHostAddress(static_cast<QHostAddress::SpecialAddress>(address));
}

void* QHostAddress_NewQHostAddress8(void* address)
{
	return new QHostAddress(*static_cast<QHostAddress*>(address));
}

void* QHostAddress_NewQHostAddress7(struct QtNetwork_PackedString address)
{
	return new QHostAddress(QString::fromUtf8(address.data, address.len));
}

void* QHostAddress_NewQHostAddress4(char* ip6Addr)
{
	return new QHostAddress(const_cast<const quint8*>(static_cast<quint8*>(static_cast<void*>(ip6Addr))));
}

void* QHostAddress_NewQHostAddress2(unsigned int ip4Addr)
{
	return new QHostAddress(ip4Addr);
}

void* QHostAddress_NewQHostAddress3(char* ip6Addr)
{
	return new QHostAddress(static_cast<quint8*>(static_cast<void*>(ip6Addr)));
}

char QHostAddress_SetAddress6(void* ptr, struct QtNetwork_PackedString address)
{
	return static_cast<QHostAddress*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void QHostAddress_Clear(void* ptr)
{
	static_cast<QHostAddress*>(ptr)->clear();
}

void QHostAddress_SetAddress7(void* ptr, long long address)
{
	static_cast<QHostAddress*>(ptr)->setAddress(static_cast<QHostAddress::SpecialAddress>(address));
}

void QHostAddress_SetAddress3(void* ptr, char* ip6Addr)
{
	static_cast<QHostAddress*>(ptr)->setAddress(const_cast<const quint8*>(static_cast<quint8*>(static_cast<void*>(ip6Addr))));
}

void QHostAddress_SetAddress(void* ptr, unsigned int ip4Addr)
{
	static_cast<QHostAddress*>(ptr)->setAddress(ip4Addr);
}

void QHostAddress_SetAddress2(void* ptr, char* ip6Addr)
{
	static_cast<QHostAddress*>(ptr)->setAddress(static_cast<quint8*>(static_cast<void*>(ip6Addr)));
}

void QHostAddress_SetScopeId(void* ptr, struct QtNetwork_PackedString id)
{
	static_cast<QHostAddress*>(ptr)->setScopeId(QString::fromUtf8(id.data, id.len));
}

void QHostAddress_Swap(void* ptr, void* other)
{
	static_cast<QHostAddress*>(ptr)->swap(*static_cast<QHostAddress*>(other));
}

void QHostAddress_DestroyQHostAddress(void* ptr)
{
	static_cast<QHostAddress*>(ptr)->~QHostAddress();
}

long long QHostAddress_Protocol(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->protocol();
}

struct QtNetwork_PackedString QHostAddress_ScopeId(void* ptr)
{
	return ({ QByteArray t9c6602 = static_cast<QHostAddress*>(ptr)->scopeId().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t9c6602.prepend("WHITESPACE").constData()+10), t9c6602.size()-10 }; });
}

struct QtNetwork_PackedString QHostAddress_ToString(void* ptr)
{
	return ({ QByteArray tc5ceab = static_cast<QHostAddress*>(ptr)->toString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc5ceab.prepend("WHITESPACE").constData()+10), tc5ceab.size()-10 }; });
}

char QHostAddress_IsBroadcast(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isBroadcast();
}

char QHostAddress_IsEqual(void* ptr, void* other, long long mode)
{
	return static_cast<QHostAddress*>(ptr)->isEqual(*static_cast<QHostAddress*>(other), static_cast<QHostAddress::ConversionModeFlag>(mode));
}

char QHostAddress_IsGlobal(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isGlobal();
}

char QHostAddress_IsInSubnet(void* ptr, void* subnet, int netmask)
{
	return static_cast<QHostAddress*>(ptr)->isInSubnet(*static_cast<QHostAddress*>(subnet), netmask);
}

char QHostAddress_IsLinkLocal(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isLinkLocal();
}

char QHostAddress_IsLoopback(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isLoopback();
}

char QHostAddress_IsMulticast(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isMulticast();
}

char QHostAddress_IsNull(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isNull();
}

char QHostAddress_IsSiteLocal(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isSiteLocal();
}

char QHostAddress_IsUniqueLocalUnicast(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isUniqueLocalUnicast();
}

unsigned int QHostAddress_ToIPv4Address(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->toIPv4Address();
}

unsigned int QHostAddress_ToIPv4Address2(void* ptr, char* ok)
{
	return static_cast<QHostAddress*>(ptr)->toIPv4Address(reinterpret_cast<bool*>(ok));
}

void* QHostInfo_QHostInfo_FromName(struct QtNetwork_PackedString name)
{
	return new QHostInfo(QHostInfo::fromName(QString::fromUtf8(name.data, name.len)));
}

void* QHostInfo_NewQHostInfo2(void* other)
{
	return new QHostInfo(*static_cast<QHostInfo*>(other));
}

void* QHostInfo_NewQHostInfo(int id)
{
	return new QHostInfo(id);
}

struct QtNetwork_PackedString QHostInfo_QHostInfo_LocalDomainName()
{
	return ({ QByteArray t5517d9 = QHostInfo::localDomainName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t5517d9.prepend("WHITESPACE").constData()+10), t5517d9.size()-10 }; });
}

struct QtNetwork_PackedString QHostInfo_QHostInfo_LocalHostName()
{
	return ({ QByteArray t63826c = QHostInfo::localHostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t63826c.prepend("WHITESPACE").constData()+10), t63826c.size()-10 }; });
}

int QHostInfo_QHostInfo_LookupHost(struct QtNetwork_PackedString name, void* receiver, char* member)
{
	return QHostInfo::lookupHost(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QHostInfo_QHostInfo_AbortHostLookup(int id)
{
	QHostInfo::abortHostLookup(id);
}

void QHostInfo_SetAddresses(void* ptr, void* addresses)
{
	static_cast<QHostInfo*>(ptr)->setAddresses(*static_cast<QList<QHostAddress>*>(addresses));
}

void QHostInfo_SetError(void* ptr, long long error)
{
	static_cast<QHostInfo*>(ptr)->setError(static_cast<QHostInfo::HostInfoError>(error));
}

void QHostInfo_SetErrorString(void* ptr, struct QtNetwork_PackedString str)
{
	static_cast<QHostInfo*>(ptr)->setErrorString(QString::fromUtf8(str.data, str.len));
}

void QHostInfo_SetHostName(void* ptr, struct QtNetwork_PackedString hostName)
{
	static_cast<QHostInfo*>(ptr)->setHostName(QString::fromUtf8(hostName.data, hostName.len));
}

void QHostInfo_SetLookupId(void* ptr, int id)
{
	static_cast<QHostInfo*>(ptr)->setLookupId(id);
}

void QHostInfo_Swap(void* ptr, void* other)
{
	static_cast<QHostInfo*>(ptr)->swap(*static_cast<QHostInfo*>(other));
}

void QHostInfo_DestroyQHostInfo(void* ptr)
{
	static_cast<QHostInfo*>(ptr)->~QHostInfo();
}

long long QHostInfo_Error(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->error();
}

struct QtNetwork_PackedList QHostInfo_Addresses(void* ptr)
{
	return ({ QList<QHostAddress>* tmpValue = new QList<QHostAddress>(static_cast<QHostInfo*>(ptr)->addresses()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedString QHostInfo_ErrorString(void* ptr)
{
	return ({ QByteArray taf4307 = static_cast<QHostInfo*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(taf4307.prepend("WHITESPACE").constData()+10), taf4307.size()-10 }; });
}

struct QtNetwork_PackedString QHostInfo_HostName(void* ptr)
{
	return ({ QByteArray t7a1d02 = static_cast<QHostInfo*>(ptr)->hostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t7a1d02.prepend("WHITESPACE").constData()+10), t7a1d02.size()-10 }; });
}

int QHostInfo_LookupId(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->lookupId();
}

void* QHostInfo___setAddresses_addresses_atList(void* ptr, int i)
{
	return new QHostAddress(({QHostAddress tmp = static_cast<QList<QHostAddress>*>(ptr)->at(i); if (i == static_cast<QList<QHostAddress>*>(ptr)->size()-1) { static_cast<QList<QHostAddress>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QHostInfo___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QHostAddress>*>(ptr)->append(*static_cast<QHostAddress*>(i));
}

void* QHostInfo___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHostAddress>();
}

void* QHostInfo___addresses_atList(void* ptr, int i)
{
	return new QHostAddress(({QHostAddress tmp = static_cast<QList<QHostAddress>*>(ptr)->at(i); if (i == static_cast<QList<QHostAddress>*>(ptr)->size()-1) { static_cast<QList<QHostAddress>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QHostInfo___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QHostAddress>*>(ptr)->append(*static_cast<QHostAddress*>(i));
}

void* QHostInfo___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHostAddress>();
}

void* QHstsPolicy_NewQHstsPolicy()
{
	return new QHstsPolicy();
}

void* QHstsPolicy_NewQHstsPolicy2(void* expiry, long long flags, struct QtNetwork_PackedString host, long long mode)
{
	return new QHstsPolicy(*static_cast<QDateTime*>(expiry), static_cast<QHstsPolicy::PolicyFlag>(flags), QString::fromUtf8(host.data, host.len), static_cast<QUrl::ParsingMode>(mode));
}

void* QHstsPolicy_NewQHstsPolicy3(void* other)
{
	return new QHstsPolicy(*static_cast<QHstsPolicy*>(other));
}

void QHstsPolicy_SetExpiry(void* ptr, void* expiry)
{
	static_cast<QHstsPolicy*>(ptr)->setExpiry(*static_cast<QDateTime*>(expiry));
}

void QHstsPolicy_SetHost(void* ptr, struct QtNetwork_PackedString host, long long mode)
{
	static_cast<QHstsPolicy*>(ptr)->setHost(QString::fromUtf8(host.data, host.len), static_cast<QUrl::ParsingMode>(mode));
}

void QHstsPolicy_SetIncludesSubDomains(void* ptr, char include)
{
	static_cast<QHstsPolicy*>(ptr)->setIncludesSubDomains(include != 0);
}

void QHstsPolicy_Swap(void* ptr, void* other)
{
	static_cast<QHstsPolicy*>(ptr)->swap(*static_cast<QHstsPolicy*>(other));
}

void QHstsPolicy_DestroyQHstsPolicy(void* ptr)
{
	static_cast<QHstsPolicy*>(ptr)->~QHstsPolicy();
}

void* QHstsPolicy_Expiry(void* ptr)
{
	return new QDateTime(static_cast<QHstsPolicy*>(ptr)->expiry());
}

struct QtNetwork_PackedString QHstsPolicy_Host(void* ptr, long long options)
{
	return ({ QByteArray tfcde23 = static_cast<QHstsPolicy*>(ptr)->host(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfcde23.prepend("WHITESPACE").constData()+10), tfcde23.size()-10 }; });
}

char QHstsPolicy_IncludesSubDomains(void* ptr)
{
	return static_cast<QHstsPolicy*>(ptr)->includesSubDomains();
}

char QHstsPolicy_IsExpired(void* ptr)
{
	return static_cast<QHstsPolicy*>(ptr)->isExpired();
}

class MyQHttpMultiPart: public QHttpMultiPart
{
public:
	MyQHttpMultiPart(QHttpMultiPart::ContentType contentType, QObject *parent = Q_NULLPTR) : QHttpMultiPart(contentType, parent) {QHttpMultiPart_QHttpMultiPart_QRegisterMetaType();};
	MyQHttpMultiPart(QObject *parent = Q_NULLPTR) : QHttpMultiPart(parent) {QHttpMultiPart_QHttpMultiPart_QRegisterMetaType();};
	 ~MyQHttpMultiPart() { callbackQHttpMultiPart_DestroyQHttpMultiPart(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHttpMultiPart_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQHttpMultiPart_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHttpMultiPart_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHttpMultiPart_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHttpMultiPart_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHttpMultiPart_CustomEvent(this, event); };
	void deleteLater() { callbackQHttpMultiPart_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHttpMultiPart_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHttpMultiPart_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHttpMultiPart_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHttpMultiPart_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHttpMultiPart*)

int QHttpMultiPart_QHttpMultiPart_QRegisterMetaType(){qRegisterMetaType<QHttpMultiPart*>(); return qRegisterMetaType<MyQHttpMultiPart*>();}

void* QHttpMultiPart_NewQHttpMultiPart2(long long contentType, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QWindow*>(parent));
	} else {
		return new MyQHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QObject*>(parent));
	}
}

void* QHttpMultiPart_NewQHttpMultiPart(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHttpMultiPart(static_cast<QWindow*>(parent));
	} else {
		return new MyQHttpMultiPart(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QHttpMultiPart_QHttpMultiPart_Tr(char* s, char* c, int n)
{
	return ({ QByteArray te425de = QHttpMultiPart::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(te425de.prepend("WHITESPACE").constData()+10), te425de.size()-10 }; });
}

struct QtNetwork_PackedString QHttpMultiPart_QHttpMultiPart_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray ta85372 = QHttpMultiPart::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta85372.prepend("WHITESPACE").constData()+10), ta85372.size()-10 }; });
}

void QHttpMultiPart_Append(void* ptr, void* httpPart)
{
	static_cast<QHttpMultiPart*>(ptr)->append(*static_cast<QHttpPart*>(httpPart));
}

void QHttpMultiPart_SetBoundary(void* ptr, void* boundary)
{
	static_cast<QHttpMultiPart*>(ptr)->setBoundary(*static_cast<QByteArray*>(boundary));
}

void QHttpMultiPart_SetContentType(void* ptr, long long contentType)
{
	static_cast<QHttpMultiPart*>(ptr)->setContentType(static_cast<QHttpMultiPart::ContentType>(contentType));
}

void QHttpMultiPart_DestroyQHttpMultiPart(void* ptr)
{
	static_cast<QHttpMultiPart*>(ptr)->~QHttpMultiPart();
}

void QHttpMultiPart_DestroyQHttpMultiPartDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QHttpMultiPart_Boundary(void* ptr)
{
	return new QByteArray(static_cast<QHttpMultiPart*>(ptr)->boundary());
}

void* QHttpMultiPart_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::metaObject());
}

void* QHttpMultiPart___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QHttpMultiPart___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHttpMultiPart___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QHttpMultiPart___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHttpMultiPart___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHttpMultiPart___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHttpMultiPart___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHttpMultiPart___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHttpMultiPart___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHttpMultiPart___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHttpMultiPart___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHttpMultiPart___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHttpMultiPart___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHttpMultiPart___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHttpMultiPart___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QHttpMultiPart_EventDefault(void* ptr, void* e)
{
		return static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::event(static_cast<QEvent*>(e));
}

char QHttpMultiPart_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHttpMultiPart_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::childEvent(static_cast<QChildEvent*>(event));
}

void QHttpMultiPart_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHttpMultiPart_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::customEvent(static_cast<QEvent*>(event));
}

void QHttpMultiPart_DeleteLaterDefault(void* ptr)
{
		static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::deleteLater();
}

void QHttpMultiPart_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHttpMultiPart_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QHttpPart_NewQHttpPart()
{
	return new QHttpPart();
}

void* QHttpPart_NewQHttpPart2(void* other)
{
	return new QHttpPart(*static_cast<QHttpPart*>(other));
}

void QHttpPart_SetBody(void* ptr, void* body)
{
	static_cast<QHttpPart*>(ptr)->setBody(*static_cast<QByteArray*>(body));
}

void QHttpPart_SetBodyDevice(void* ptr, void* device)
{
	static_cast<QHttpPart*>(ptr)->setBodyDevice(static_cast<QIODevice*>(device));
}

void QHttpPart_SetHeader(void* ptr, long long header, void* value)
{
	static_cast<QHttpPart*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QHttpPart_SetRawHeader(void* ptr, void* headerName, void* headerValue)
{
	static_cast<QHttpPart*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QHttpPart_Swap(void* ptr, void* other)
{
	static_cast<QHttpPart*>(ptr)->swap(*static_cast<QHttpPart*>(other));
}

void QHttpPart_DestroyQHttpPart(void* ptr)
{
	static_cast<QHttpPart*>(ptr)->~QHttpPart();
}

class MyQLocalServer: public QLocalServer
{
public:
	MyQLocalServer(QObject *parent = Q_NULLPTR) : QLocalServer(parent) {QLocalServer_QLocalServer_QRegisterMetaType();};
	QLocalSocket * nextPendingConnection() { return static_cast<QLocalSocket*>(callbackQLocalServer_NextPendingConnection(this)); };
	void incomingConnection(quintptr socketDescriptor) { callbackQLocalServer_IncomingConnection(this, socketDescriptor); };
	void Signal_NewConnection() { callbackQLocalServer_NewConnection(this); };
	 ~MyQLocalServer() { callbackQLocalServer_DestroyQLocalServer(this); };
	bool hasPendingConnections() const { return callbackQLocalServer_HasPendingConnections(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLocalServer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQLocalServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLocalServer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLocalServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLocalServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLocalServer_CustomEvent(this, event); };
	void deleteLater() { callbackQLocalServer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLocalServer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLocalServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLocalServer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLocalServer_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQLocalServer*)

int QLocalServer_QLocalServer_QRegisterMetaType(){qRegisterMetaType<QLocalServer*>(); return qRegisterMetaType<MyQLocalServer*>();}

void* QLocalServer_NewQLocalServer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLocalServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQLocalServer(static_cast<QObject*>(parent));
	}
}

void* QLocalServer_NextPendingConnection(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->nextPendingConnection();
}

void* QLocalServer_NextPendingConnectionDefault(void* ptr)
{
		return static_cast<QLocalServer*>(ptr)->QLocalServer::nextPendingConnection();
}

struct QtNetwork_PackedString QLocalServer_QLocalServer_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t77806e = QLocalServer::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t77806e.prepend("WHITESPACE").constData()+10), t77806e.size()-10 }; });
}

struct QtNetwork_PackedString QLocalServer_QLocalServer_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t37bb1f = QLocalServer::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t37bb1f.prepend("WHITESPACE").constData()+10), t37bb1f.size()-10 }; });
}

char QLocalServer_Listen(void* ptr, struct QtNetwork_PackedString name)
{
	return static_cast<QLocalServer*>(ptr)->listen(QString::fromUtf8(name.data, name.len));
}

char QLocalServer_QLocalServer_RemoveServer(struct QtNetwork_PackedString name)
{
	return QLocalServer::removeServer(QString::fromUtf8(name.data, name.len));
}

char QLocalServer_WaitForNewConnection(void* ptr, int msec, char* timedOut)
{
	return static_cast<QLocalServer*>(ptr)->waitForNewConnection(msec, reinterpret_cast<bool*>(timedOut));
}

void QLocalServer_Close(void* ptr)
{
	static_cast<QLocalServer*>(ptr)->close();
}

void QLocalServer_IncomingConnection(void* ptr, uintptr_t socketDescriptor)
{
	static_cast<QLocalServer*>(ptr)->incomingConnection(socketDescriptor);
}

void QLocalServer_IncomingConnectionDefault(void* ptr, uintptr_t socketDescriptor)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::incomingConnection(socketDescriptor);
}

void QLocalServer_ConnectNewConnection(void* ptr)
{
	QObject::connect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));
}

void QLocalServer_DisconnectNewConnection(void* ptr)
{
	QObject::disconnect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));
}

void QLocalServer_NewConnection(void* ptr)
{
	static_cast<QLocalServer*>(ptr)->newConnection();
}

void QLocalServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QLocalServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QLocalServer_SetSocketOptions(void* ptr, long long options)
{
	static_cast<QLocalServer*>(ptr)->setSocketOptions(static_cast<QLocalServer::SocketOption>(options));
}

void QLocalServer_DestroyQLocalServer(void* ptr)
{
	static_cast<QLocalServer*>(ptr)->~QLocalServer();
}

void QLocalServer_DestroyQLocalServerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QLocalServer_ServerError(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->serverError();
}

long long QLocalServer_SocketOptions(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->socketOptions();
}

struct QtNetwork_PackedString QLocalServer_ErrorString(void* ptr)
{
	return ({ QByteArray tf5dac0 = static_cast<QLocalServer*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tf5dac0.prepend("WHITESPACE").constData()+10), tf5dac0.size()-10 }; });
}

struct QtNetwork_PackedString QLocalServer_FullServerName(void* ptr)
{
	return ({ QByteArray tb91f50 = static_cast<QLocalServer*>(ptr)->fullServerName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb91f50.prepend("WHITESPACE").constData()+10), tb91f50.size()-10 }; });
}

struct QtNetwork_PackedString QLocalServer_ServerName(void* ptr)
{
	return ({ QByteArray t054e78 = static_cast<QLocalServer*>(ptr)->serverName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t054e78.prepend("WHITESPACE").constData()+10), t054e78.size()-10 }; });
}

char QLocalServer_HasPendingConnections(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->hasPendingConnections();
}

char QLocalServer_HasPendingConnectionsDefault(void* ptr)
{
		return static_cast<QLocalServer*>(ptr)->QLocalServer::hasPendingConnections();
}

char QLocalServer_IsListening(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->isListening();
}

void* QLocalServer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QLocalServer*>(ptr)->QLocalServer::metaObject());
}

int QLocalServer_MaxPendingConnections(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->maxPendingConnections();
}

void* QLocalServer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLocalServer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QLocalServer___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QLocalServer___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalServer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalServer___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLocalServer___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalServer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalServer___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLocalServer___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalServer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalServer___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLocalServer___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalServer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalServer___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QLocalServer_EventDefault(void* ptr, void* e)
{
		return static_cast<QLocalServer*>(ptr)->QLocalServer::event(static_cast<QEvent*>(e));
}

char QLocalServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QLocalServer*>(ptr)->QLocalServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QLocalServer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::childEvent(static_cast<QChildEvent*>(event));
}

void QLocalServer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalServer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::customEvent(static_cast<QEvent*>(event));
}

void QLocalServer_DeleteLaterDefault(void* ptr)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::deleteLater();
}

void QLocalServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalServer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QLocalServer*>(ptr)->QLocalServer::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQLocalSocket: public QLocalSocket
{
public:
	MyQLocalSocket(QObject *parent = Q_NULLPTR) : QLocalSocket(parent) {QLocalSocket_QLocalSocket_QRegisterMetaType();};
	bool open(QIODevice::OpenMode openMode) { return callbackQLocalSocket_Open(this, openMode) != 0; };
	bool waitForBytesWritten(int msecs) { return callbackQLocalSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQLocalSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 readData(char * data, qint64 c) { QtNetwork_PackedString dataPacked = { data, c };return callbackQLocalSocket_ReadData(this, dataPacked, c); };
	qint64 writeData(const char * data, qint64 c) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), c };return callbackQLocalSocket_WriteData(this, dataPacked, c); };
	void close() { callbackQLocalSocket_Close(this); };
	void Signal_Connected() { callbackQLocalSocket_Connected(this); };
	void Signal_Disconnected() { callbackQLocalSocket_Disconnected(this); };
	void Signal_Error2(QLocalSocket::LocalSocketError socketError) { callbackQLocalSocket_Error2(this, socketError); };
	void Signal_StateChanged(QLocalSocket::LocalSocketState socketState) { callbackQLocalSocket_StateChanged(this, socketState); };
	 ~MyQLocalSocket() { callbackQLocalSocket_DestroyQLocalSocket(this); };
	bool canReadLine() const { return callbackQLocalSocket_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isSequential() const { return callbackQLocalSocket_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLocalSocket_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	qint64 bytesAvailable() const { return callbackQLocalSocket_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQLocalSocket_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	bool reset() { return callbackQLocalSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQLocalSocket_Seek(this, pos) != 0; };
	qint64 readLineData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQLocalSocket_ReadLineData(this, dataPacked, maxSize); };
	void Signal_AboutToClose() { callbackQLocalSocket_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQLocalSocket_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQLocalSocket_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQLocalSocket_ChannelReadyRead(this, channel); };
	void Signal_ReadChannelFinished() { callbackQLocalSocket_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQLocalSocket_ReadyRead(this); };
	bool atEnd() const { return callbackQLocalSocket_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	qint64 pos() const { return callbackQLocalSocket_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQLocalSocket_Size(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQLocalSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLocalSocket_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLocalSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLocalSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLocalSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQLocalSocket_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLocalSocket_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLocalSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLocalSocket_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLocalSocket_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQLocalSocket*)

int QLocalSocket_QLocalSocket_QRegisterMetaType(){qRegisterMetaType<QLocalSocket*>(); return qRegisterMetaType<MyQLocalSocket*>();}

void* QLocalSocket_NewQLocalSocket(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLocalSocket(static_cast<QWindow*>(parent));
	} else {
		return new MyQLocalSocket(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QLocalSocket_QLocalSocket_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t69a8c0 = QLocalSocket::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t69a8c0.prepend("WHITESPACE").constData()+10), t69a8c0.size()-10 }; });
}

struct QtNetwork_PackedString QLocalSocket_QLocalSocket_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t7fb76e = QLocalSocket::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t7fb76e.prepend("WHITESPACE").constData()+10), t7fb76e.size()-10 }; });
}

char QLocalSocket_Flush(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->flush();
}

char QLocalSocket_OpenDefault(void* ptr, long long openMode)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

char QLocalSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::waitForBytesWritten(msecs);
}

char QLocalSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForConnected(msecs);
}

char QLocalSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForDisconnected(msecs);
}

char QLocalSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::waitForReadyRead(msecs);
}

long long QLocalSocket_ReadData(void* ptr, char* data, long long c)
{
	return static_cast<QLocalSocket*>(ptr)->readData(data, c);
}

long long QLocalSocket_ReadDataDefault(void* ptr, char* data, long long c)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::readData(data, c);
}

long long QLocalSocket_WriteData(void* ptr, char* data, long long c)
{
	return static_cast<QLocalSocket*>(ptr)->writeData(const_cast<const char*>(data), c);
}

long long QLocalSocket_WriteDataDefault(void* ptr, char* data, long long c)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::writeData(const_cast<const char*>(data), c);
}

void QLocalSocket_Abort(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->abort();
}

void QLocalSocket_CloseDefault(void* ptr)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::close();
}

void QLocalSocket_ConnectToServer(void* ptr, long long openMode)
{
	static_cast<QLocalSocket*>(ptr)->connectToServer(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_ConnectToServer2(void* ptr, struct QtNetwork_PackedString name, long long openMode)
{
	static_cast<QLocalSocket*>(ptr)->connectToServer(QString::fromUtf8(name.data, name.len), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_ConnectConnected(void* ptr)
{
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Connected));
}

void QLocalSocket_DisconnectConnected(void* ptr)
{
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Connected));
}

void QLocalSocket_Connected(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->connected();
}

void QLocalSocket_DisconnectFromServer(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->disconnectFromServer();
}

void QLocalSocket_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Disconnected));
}

void QLocalSocket_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Disconnected));
}

void QLocalSocket_Disconnected(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->disconnected();
}

void QLocalSocket_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketError)>(&QLocalSocket::error), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketError)>(&MyQLocalSocket::Signal_Error2));
}

void QLocalSocket_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketError)>(&QLocalSocket::error), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketError)>(&MyQLocalSocket::Signal_Error2));
}

void QLocalSocket_Error2(void* ptr, long long socketError)
{
	static_cast<QLocalSocket*>(ptr)->error(static_cast<QLocalSocket::LocalSocketError>(socketError));
}

void QLocalSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QLocalSocket*>(ptr)->setReadBufferSize(size);
}

void QLocalSocket_SetServerName(void* ptr, struct QtNetwork_PackedString name)
{
	static_cast<QLocalSocket*>(ptr)->setServerName(QString::fromUtf8(name.data, name.len));
}

void QLocalSocket_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));
}

void QLocalSocket_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));
}

void QLocalSocket_StateChanged(void* ptr, long long socketState)
{
	static_cast<QLocalSocket*>(ptr)->stateChanged(static_cast<QLocalSocket::LocalSocketState>(socketState));
}

void QLocalSocket_DestroyQLocalSocket(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->~QLocalSocket();
}

void QLocalSocket_DestroyQLocalSocketDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QLocalSocket_Error(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->error();
}

long long QLocalSocket_State(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->state();
}

struct QtNetwork_PackedString QLocalSocket_FullServerName(void* ptr)
{
	return ({ QByteArray ta11089 = static_cast<QLocalSocket*>(ptr)->fullServerName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta11089.prepend("WHITESPACE").constData()+10), ta11089.size()-10 }; });
}

struct QtNetwork_PackedString QLocalSocket_ServerName(void* ptr)
{
	return ({ QByteArray t348d56 = static_cast<QLocalSocket*>(ptr)->serverName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t348d56.prepend("WHITESPACE").constData()+10), t348d56.size()-10 }; });
}

char QLocalSocket_CanReadLineDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::canReadLine();
}

char QLocalSocket_IsSequentialDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::isSequential();
}

char QLocalSocket_IsValid(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->isValid();
}

void* QLocalSocket_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::metaObject());
}

long long QLocalSocket_BytesAvailableDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::bytesAvailable();
}

long long QLocalSocket_BytesToWriteDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::bytesToWrite();
}

long long QLocalSocket_ReadBufferSize(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->readBufferSize();
}

void* QLocalSocket___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLocalSocket___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QLocalSocket___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QLocalSocket___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalSocket___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalSocket___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLocalSocket___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalSocket___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalSocket___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLocalSocket___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalSocket___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalSocket___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLocalSocket___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocalSocket___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLocalSocket___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QLocalSocket_ResetDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::reset();
}

char QLocalSocket_SeekDefault(void* ptr, long long pos)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::seek(pos);
}

long long QLocalSocket_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::readLineData(data, maxSize);
}

char QLocalSocket_AtEndDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::atEnd();
}

long long QLocalSocket_PosDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::pos();
}

long long QLocalSocket_SizeDefault(void* ptr)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::size();
}

char QLocalSocket_EventDefault(void* ptr, void* e)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::event(static_cast<QEvent*>(e));
}

char QLocalSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QLocalSocket*>(ptr)->QLocalSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QLocalSocket_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QLocalSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalSocket_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::customEvent(static_cast<QEvent*>(event));
}

void QLocalSocket_DeleteLaterDefault(void* ptr)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::deleteLater();
}

void QLocalSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalSocket_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QLocalSocket*>(ptr)->QLocalSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQNetworkAccessManager: public QNetworkAccessManager
{
public:
	MyQNetworkAccessManager(QObject *parent = Q_NULLPTR) : QNetworkAccessManager(parent) {QNetworkAccessManager_QNetworkAccessManager_QRegisterMetaType();};
	QNetworkReply * createRequest(QNetworkAccessManager::Operation op, const QNetworkRequest & originalReq, QIODevice * outgoingData) { return static_cast<QNetworkReply*>(callbackQNetworkAccessManager_CreateRequest(this, op, const_cast<QNetworkRequest*>(&originalReq), outgoingData)); };
	void Signal_AuthenticationRequired(QNetworkReply * reply, QAuthenticator * authenticator) { callbackQNetworkAccessManager_AuthenticationRequired(this, reply, authenticator); };
	void Signal_Encrypted(QNetworkReply * reply) { callbackQNetworkAccessManager_Encrypted(this, reply); };
	void Signal_Finished(QNetworkReply * reply) { callbackQNetworkAccessManager_Finished(this, reply); };
	void Signal_NetworkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility accessible) { callbackQNetworkAccessManager_NetworkAccessibleChanged(this, accessible); };
	void Signal_PreSharedKeyAuthenticationRequired(QNetworkReply * reply, QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired(this, reply, authenticator); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQNetworkAccessManager_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	void Signal_SslErrors(QNetworkReply * reply, const QList<QSslError> & errors) { callbackQNetworkAccessManager_SslErrors(this, reply, ({ QList<QSslError>* tmpValue = const_cast<QList<QSslError>*>(&errors); QtNetwork_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~MyQNetworkAccessManager() { callbackQNetworkAccessManager_DestroyQNetworkAccessManager(this); };
	QStringList supportedSchemesImplementation() const { return ({ QtNetwork_PackedString tempVal = callbackQNetworkAccessManager_SupportedSchemesImplementation(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkAccessManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNetworkAccessManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkAccessManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNetworkAccessManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkAccessManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkAccessManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkAccessManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNetworkAccessManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkAccessManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNetworkAccessManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkAccessManager_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNetworkAccessManager*)

int QNetworkAccessManager_QNetworkAccessManager_QRegisterMetaType(){qRegisterMetaType<QNetworkAccessManager*>(); return qRegisterMetaType<MyQNetworkAccessManager*>();}

void* QNetworkAccessManager_NewQNetworkAccessManager(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkAccessManager(static_cast<QWindow*>(parent));
	} else {
		return new MyQNetworkAccessManager(static_cast<QObject*>(parent));
	}
}

void* QNetworkAccessManager_CreateRequest(void* ptr, long long op, void* originalReq, void* outgoingData)
{
	return static_cast<QNetworkAccessManager*>(ptr)->createRequest(static_cast<QNetworkAccessManager::Operation>(op), *static_cast<QNetworkRequest*>(originalReq), static_cast<QIODevice*>(outgoingData));
}

void* QNetworkAccessManager_CreateRequestDefault(void* ptr, long long op, void* originalReq, void* outgoingData)
{
		return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::createRequest(static_cast<QNetworkAccessManager::Operation>(op), *static_cast<QNetworkRequest*>(originalReq), static_cast<QIODevice*>(outgoingData));
}

void* QNetworkAccessManager_DeleteResource(void* ptr, void* request)
{
	return static_cast<QNetworkAccessManager*>(ptr)->deleteResource(*static_cast<QNetworkRequest*>(request));
}

void* QNetworkAccessManager_Get(void* ptr, void* request)
{
	return static_cast<QNetworkAccessManager*>(ptr)->get(*static_cast<QNetworkRequest*>(request));
}

void* QNetworkAccessManager_Head(void* ptr, void* request)
{
	return static_cast<QNetworkAccessManager*>(ptr)->head(*static_cast<QNetworkRequest*>(request));
}

void* QNetworkAccessManager_Post3(void* ptr, void* request, void* multiPart)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Post(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Post2(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void* QNetworkAccessManager_Put3(void* ptr, void* request, void* multiPart)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Put(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Put2(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void* QNetworkAccessManager_SendCustomRequest3(void* ptr, void* request, void* verb, void* multiPart)
{
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_SendCustomRequest(void* ptr, void* request, void* verb, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_SendCustomRequest2(void* ptr, void* request, void* verb, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), *static_cast<QByteArray*>(data));
}

struct QtNetwork_PackedString QNetworkAccessManager_QNetworkAccessManager_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t02c015 = QNetworkAccessManager::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t02c015.prepend("WHITESPACE").constData()+10), t02c015.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkAccessManager_QNetworkAccessManager_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t0482aa = QNetworkAccessManager::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t0482aa.prepend("WHITESPACE").constData()+10), t0482aa.size()-10 }; });
}

void QNetworkAccessManager_AddStrictTransportSecurityHosts(void* ptr, void* knownHosts)
{
	static_cast<QNetworkAccessManager*>(ptr)->addStrictTransportSecurityHosts(*static_cast<QVector<QHstsPolicy>*>(knownHosts));
}

void QNetworkAccessManager_ConnectAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));
}

void QNetworkAccessManager_DisconnectAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));
}

void QNetworkAccessManager_AuthenticationRequired(void* ptr, void* reply, void* authenticator)
{
	static_cast<QNetworkAccessManager*>(ptr)->authenticationRequired(static_cast<QNetworkReply*>(reply), static_cast<QAuthenticator*>(authenticator));
}

void QNetworkAccessManager_ClearAccessCache(void* ptr)
{
	static_cast<QNetworkAccessManager*>(ptr)->clearAccessCache();
}

void QNetworkAccessManager_ClearConnectionCache(void* ptr)
{
	static_cast<QNetworkAccessManager*>(ptr)->clearConnectionCache();
}

void QNetworkAccessManager_ConnectToHost(void* ptr, struct QtNetwork_PackedString hostName, unsigned short port)
{
	static_cast<QNetworkAccessManager*>(ptr)->connectToHost(QString::fromUtf8(hostName.data, hostName.len), port);
}

void QNetworkAccessManager_ConnectToHostEncrypted(void* ptr, struct QtNetwork_PackedString hostName, unsigned short port, void* sslConfiguration)
{
	static_cast<QNetworkAccessManager*>(ptr)->connectToHostEncrypted(QString::fromUtf8(hostName.data, hostName.len), port, *static_cast<QSslConfiguration*>(sslConfiguration));
}

void QNetworkAccessManager_EnableStrictTransportSecurityStore(void* ptr, char enabled, struct QtNetwork_PackedString storeDir)
{
	static_cast<QNetworkAccessManager*>(ptr)->enableStrictTransportSecurityStore(enabled != 0, QString::fromUtf8(storeDir.data, storeDir.len));
}

void QNetworkAccessManager_ConnectEncrypted(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));
}

void QNetworkAccessManager_DisconnectEncrypted(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));
}

void QNetworkAccessManager_Encrypted(void* ptr, void* reply)
{
	static_cast<QNetworkAccessManager*>(ptr)->encrypted(static_cast<QNetworkReply*>(reply));
}

void QNetworkAccessManager_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));
}

void QNetworkAccessManager_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));
}

void QNetworkAccessManager_Finished(void* ptr, void* reply)
{
	static_cast<QNetworkAccessManager*>(ptr)->finished(static_cast<QNetworkReply*>(reply));
}

void QNetworkAccessManager_ConnectNetworkAccessibleChanged(void* ptr)
{
	qRegisterMetaType<QNetworkAccessManager::NetworkAccessibility>();
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));
}

void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));
}

void QNetworkAccessManager_NetworkAccessibleChanged(void* ptr, long long accessible)
{
	static_cast<QNetworkAccessManager*>(ptr)->networkAccessibleChanged(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));
}

void QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));
}

void QNetworkAccessManager_PreSharedKeyAuthenticationRequired(void* ptr, void* reply, void* authenticator)
{
	static_cast<QNetworkAccessManager*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QNetworkReply*>(reply), static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QNetworkAccessManager_ConnectProxyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(const QNetworkProxy &, QAuthenticator *)>(&QNetworkAccessManager::proxyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(const QNetworkProxy &, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_ProxyAuthenticationRequired));
}

void QNetworkAccessManager_DisconnectProxyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(const QNetworkProxy &, QAuthenticator *)>(&QNetworkAccessManager::proxyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(const QNetworkProxy &, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_ProxyAuthenticationRequired));
}

void QNetworkAccessManager_ProxyAuthenticationRequired(void* ptr, void* proxy, void* authenticator)
{
	static_cast<QNetworkAccessManager*>(ptr)->proxyAuthenticationRequired(*static_cast<QNetworkProxy*>(proxy), static_cast<QAuthenticator*>(authenticator));
}

void QNetworkAccessManager_SetCache(void* ptr, void* cache)
{
	static_cast<QNetworkAccessManager*>(ptr)->setCache(static_cast<QAbstractNetworkCache*>(cache));
}

void QNetworkAccessManager_SetConfiguration(void* ptr, void* config)
{
	static_cast<QNetworkAccessManager*>(ptr)->setConfiguration(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkAccessManager_SetCookieJar(void* ptr, void* cookieJar)
{
	static_cast<QNetworkAccessManager*>(ptr)->setCookieJar(static_cast<QNetworkCookieJar*>(cookieJar));
}

void QNetworkAccessManager_SetNetworkAccessible(void* ptr, long long accessible)
{
	static_cast<QNetworkAccessManager*>(ptr)->setNetworkAccessible(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void QNetworkAccessManager_SetProxy(void* ptr, void* proxy)
{
	static_cast<QNetworkAccessManager*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(proxy));
}

void QNetworkAccessManager_SetProxyFactory(void* ptr, void* factory)
{
	static_cast<QNetworkAccessManager*>(ptr)->setProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

void QNetworkAccessManager_SetRedirectPolicy(void* ptr, long long policy)
{
	static_cast<QNetworkAccessManager*>(ptr)->setRedirectPolicy(static_cast<QNetworkRequest::RedirectPolicy>(policy));
}

void QNetworkAccessManager_SetStrictTransportSecurityEnabled(void* ptr, char enabled)
{
	static_cast<QNetworkAccessManager*>(ptr)->setStrictTransportSecurityEnabled(enabled != 0);
}

void QNetworkAccessManager_ConnectSslErrors(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, const QList<QSslError> &)>(&QNetworkAccessManager::sslErrors), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, const QList<QSslError> &)>(&MyQNetworkAccessManager::Signal_SslErrors));
}

void QNetworkAccessManager_DisconnectSslErrors(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, const QList<QSslError> &)>(&QNetworkAccessManager::sslErrors), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, const QList<QSslError> &)>(&MyQNetworkAccessManager::Signal_SslErrors));
}

void QNetworkAccessManager_SslErrors(void* ptr, void* reply, void* errors)
{
	static_cast<QNetworkAccessManager*>(ptr)->sslErrors(static_cast<QNetworkReply*>(reply), *static_cast<QList<QSslError>*>(errors));
}

void QNetworkAccessManager_DestroyQNetworkAccessManager(void* ptr)
{
	static_cast<QNetworkAccessManager*>(ptr)->~QNetworkAccessManager();
}

void QNetworkAccessManager_DestroyQNetworkAccessManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNetworkAccessManager_Cache(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->cache();
}

long long QNetworkAccessManager_NetworkAccessible(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->networkAccessible();
}

void* QNetworkAccessManager_ActiveConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkAccessManager*>(ptr)->activeConfiguration());
}

void* QNetworkAccessManager_Configuration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkAccessManager*>(ptr)->configuration());
}

void* QNetworkAccessManager_CookieJar(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->cookieJar();
}

void* QNetworkAccessManager_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QNetworkAccessManager*>(ptr)->proxy());
}

void* QNetworkAccessManager_ProxyFactory(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->proxyFactory();
}

long long QNetworkAccessManager_RedirectPolicy(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->redirectPolicy();
}

struct QtNetwork_PackedString QNetworkAccessManager_SupportedSchemes(void* ptr)
{
	return ({ QByteArray tda9ca8 = static_cast<QNetworkAccessManager*>(ptr)->supportedSchemes().join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tda9ca8.prepend("WHITESPACE").constData()+10), tda9ca8.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkAccessManager_SupportedSchemesImplementation(void* ptr)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<QNetworkAccessManager*>(ptr), "supportedSchemesImplementation", Q_RETURN_ARG(QStringList, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkAccessManager_SupportedSchemesImplementationDefault(void* ptr)
{
		return ({ QByteArray t5e649f = static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::supportedSchemesImplementation().join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(t5e649f.prepend("WHITESPACE").constData()+10), t5e649f.size()-10 }; });
}

struct QtNetwork_PackedList QNetworkAccessManager_StrictTransportSecurityHosts(void* ptr)
{
	return ({ QVector<QHstsPolicy>* tmpValue = new QVector<QHstsPolicy>(static_cast<QNetworkAccessManager*>(ptr)->strictTransportSecurityHosts()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QNetworkAccessManager_IsStrictTransportSecurityEnabled(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->isStrictTransportSecurityEnabled();
}

char QNetworkAccessManager_IsStrictTransportSecurityStoreEnabled(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->isStrictTransportSecurityStoreEnabled();
}

void* QNetworkAccessManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::metaObject());
}

void* QNetworkAccessManager___addStrictTransportSecurityHosts_knownHosts_atList(void* ptr, int i)
{
	return new QHstsPolicy(({QHstsPolicy tmp = static_cast<QVector<QHstsPolicy>*>(ptr)->at(i); if (i == static_cast<QVector<QHstsPolicy>*>(ptr)->size()-1) { static_cast<QVector<QHstsPolicy>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QNetworkAccessManager___addStrictTransportSecurityHosts_knownHosts_setList(void* ptr, void* i)
{
	static_cast<QVector<QHstsPolicy>*>(ptr)->append(*static_cast<QHstsPolicy*>(i));
}

void* QNetworkAccessManager___addStrictTransportSecurityHosts_knownHosts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QHstsPolicy>();
}

void* QNetworkAccessManager___sslErrors_errors_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkAccessManager___sslErrors_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QNetworkAccessManager___sslErrors_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QNetworkAccessManager___strictTransportSecurityHosts_atList(void* ptr, int i)
{
	return new QHstsPolicy(({QHstsPolicy tmp = static_cast<QVector<QHstsPolicy>*>(ptr)->at(i); if (i == static_cast<QVector<QHstsPolicy>*>(ptr)->size()-1) { static_cast<QVector<QHstsPolicy>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QNetworkAccessManager___strictTransportSecurityHosts_setList(void* ptr, void* i)
{
	static_cast<QVector<QHstsPolicy>*>(ptr)->append(*static_cast<QHstsPolicy*>(i));
}

void* QNetworkAccessManager___strictTransportSecurityHosts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QHstsPolicy>();
}

void* QNetworkAccessManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkAccessManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkAccessManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNetworkAccessManager___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkAccessManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkAccessManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkAccessManager___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkAccessManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkAccessManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkAccessManager___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkAccessManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkAccessManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkAccessManager___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkAccessManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkAccessManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNetworkAccessManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::event(static_cast<QEvent*>(e));
}

char QNetworkAccessManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNetworkAccessManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkAccessManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkAccessManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::customEvent(static_cast<QEvent*>(event));
}

void QNetworkAccessManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::deleteLater();
}

void QNetworkAccessManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkAccessManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry()
{
	return new QNetworkAddressEntry();
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry2(void* other)
{
	return new QNetworkAddressEntry(*static_cast<QNetworkAddressEntry*>(other));
}

void QNetworkAddressEntry_ClearAddressLifetime(void* ptr)
{
	static_cast<QNetworkAddressEntry*>(ptr)->clearAddressLifetime();
}

void QNetworkAddressEntry_SetAddressLifetime(void* ptr, void* preferred, void* validity)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setAddressLifetime(*static_cast<QDeadlineTimer*>(preferred), *static_cast<QDeadlineTimer*>(validity));
}

void QNetworkAddressEntry_SetBroadcast(void* ptr, void* newBroadcast)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setBroadcast(*static_cast<QHostAddress*>(newBroadcast));
}

void QNetworkAddressEntry_SetDnsEligibility(void* ptr, long long status)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setDnsEligibility(static_cast<QNetworkAddressEntry::DnsEligibilityStatus>(status));
}

void QNetworkAddressEntry_SetIp(void* ptr, void* newIp)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setIp(*static_cast<QHostAddress*>(newIp));
}

void QNetworkAddressEntry_SetNetmask(void* ptr, void* newNetmask)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setNetmask(*static_cast<QHostAddress*>(newNetmask));
}

void QNetworkAddressEntry_SetPrefixLength(void* ptr, int length)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setPrefixLength(length);
}

void QNetworkAddressEntry_Swap(void* ptr, void* other)
{
	static_cast<QNetworkAddressEntry*>(ptr)->swap(*static_cast<QNetworkAddressEntry*>(other));
}

void QNetworkAddressEntry_DestroyQNetworkAddressEntry(void* ptr)
{
	static_cast<QNetworkAddressEntry*>(ptr)->~QNetworkAddressEntry();
}

void* QNetworkAddressEntry_Broadcast(void* ptr)
{
	return new QHostAddress(static_cast<QNetworkAddressEntry*>(ptr)->broadcast());
}

void* QNetworkAddressEntry_Ip(void* ptr)
{
	return new QHostAddress(static_cast<QNetworkAddressEntry*>(ptr)->ip());
}

void* QNetworkAddressEntry_Netmask(void* ptr)
{
	return new QHostAddress(static_cast<QNetworkAddressEntry*>(ptr)->netmask());
}

long long QNetworkAddressEntry_DnsEligibility(void* ptr)
{
	return static_cast<QNetworkAddressEntry*>(ptr)->dnsEligibility();
}

char QNetworkAddressEntry_IsLifetimeKnown(void* ptr)
{
	return static_cast<QNetworkAddressEntry*>(ptr)->isLifetimeKnown();
}

char QNetworkAddressEntry_IsPermanent(void* ptr)
{
	return static_cast<QNetworkAddressEntry*>(ptr)->isPermanent();
}

char QNetworkAddressEntry_IsTemporary(void* ptr)
{
	return static_cast<QNetworkAddressEntry*>(ptr)->isTemporary();
}

int QNetworkAddressEntry_PrefixLength(void* ptr)
{
	return static_cast<QNetworkAddressEntry*>(ptr)->prefixLength();
}

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData()
{
	return new QNetworkCacheMetaData();
}

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData2(void* other)
{
	return new QNetworkCacheMetaData(*static_cast<QNetworkCacheMetaData*>(other));
}

void QNetworkCacheMetaData_SetExpirationDate(void* ptr, void* dateTime)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetLastModified(void* ptr, void* dateTime)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->setLastModified(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetSaveToDisk(void* ptr, char allow)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->setSaveToDisk(allow != 0);
}

void QNetworkCacheMetaData_SetUrl(void* ptr, void* url)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkCacheMetaData_Swap(void* ptr, void* other)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->swap(*static_cast<QNetworkCacheMetaData*>(other));
}

void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(void* ptr)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->~QNetworkCacheMetaData();
}

void* QNetworkCacheMetaData_ExpirationDate(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->expirationDate());
}

void* QNetworkCacheMetaData_LastModified(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->lastModified());
}

void* QNetworkCacheMetaData_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkCacheMetaData*>(ptr)->url());
}

char QNetworkCacheMetaData_IsValid(void* ptr)
{
	return static_cast<QNetworkCacheMetaData*>(ptr)->isValid();
}

char QNetworkCacheMetaData_SaveToDisk(void* ptr)
{
	return static_cast<QNetworkCacheMetaData*>(ptr)->saveToDisk();
}

void* QNetworkConfiguration_NewQNetworkConfiguration()
{
	return new QNetworkConfiguration();
}

void* QNetworkConfiguration_NewQNetworkConfiguration2(void* other)
{
	return new QNetworkConfiguration(*static_cast<QNetworkConfiguration*>(other));
}

char QNetworkConfiguration_SetConnectTimeout(void* ptr, int timeout)
{
	return static_cast<QNetworkConfiguration*>(ptr)->setConnectTimeout(timeout);
}

void QNetworkConfiguration_Swap(void* ptr, void* other)
{
	static_cast<QNetworkConfiguration*>(ptr)->swap(*static_cast<QNetworkConfiguration*>(other));
}

void QNetworkConfiguration_DestroyQNetworkConfiguration(void* ptr)
{
	static_cast<QNetworkConfiguration*>(ptr)->~QNetworkConfiguration();
}

struct QtNetwork_PackedList QNetworkConfiguration_Children(void* ptr)
{
	return ({ QList<QNetworkConfiguration>* tmpValue = new QList<QNetworkConfiguration>(static_cast<QNetworkConfiguration*>(ptr)->children()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkConfiguration_BearerType(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerType();
}

long long QNetworkConfiguration_BearerTypeFamily(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeFamily();
}

long long QNetworkConfiguration_Purpose(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->purpose();
}

long long QNetworkConfiguration_State(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->state();
}

long long QNetworkConfiguration_Type(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->type();
}

struct QtNetwork_PackedString QNetworkConfiguration_BearerTypeName(void* ptr)
{
	return ({ QByteArray t89aad8 = static_cast<QNetworkConfiguration*>(ptr)->bearerTypeName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t89aad8.prepend("WHITESPACE").constData()+10), t89aad8.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkConfiguration_Identifier(void* ptr)
{
	return ({ QByteArray tae5c30 = static_cast<QNetworkConfiguration*>(ptr)->identifier().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tae5c30.prepend("WHITESPACE").constData()+10), tae5c30.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkConfiguration_Name(void* ptr)
{
	return ({ QByteArray t38ed5d = static_cast<QNetworkConfiguration*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t38ed5d.prepend("WHITESPACE").constData()+10), t38ed5d.size()-10 }; });
}

char QNetworkConfiguration_IsRoamingAvailable(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->isRoamingAvailable();
}

char QNetworkConfiguration_IsValid(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->isValid();
}

int QNetworkConfiguration_ConnectTimeout(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->connectTimeout();
}

void* QNetworkConfiguration___children_atList(void* ptr, int i)
{
	return new QNetworkConfiguration(({QNetworkConfiguration tmp = static_cast<QList<QNetworkConfiguration>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkConfiguration>*>(ptr)->size()-1) { static_cast<QList<QNetworkConfiguration>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkConfiguration___children_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkConfiguration>*>(ptr)->append(*static_cast<QNetworkConfiguration*>(i));
}

void* QNetworkConfiguration___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkConfiguration>();
}

class MyQNetworkConfigurationManager: public QNetworkConfigurationManager
{
public:
	MyQNetworkConfigurationManager(QObject *parent = Q_NULLPTR) : QNetworkConfigurationManager(parent) {QNetworkConfigurationManager_QNetworkConfigurationManager_QRegisterMetaType();};
	void Signal_ConfigurationAdded(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationAdded(this, const_cast<QNetworkConfiguration*>(&config)); };
	void Signal_ConfigurationChanged(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationChanged(this, const_cast<QNetworkConfiguration*>(&config)); };
	void Signal_ConfigurationRemoved(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationRemoved(this, const_cast<QNetworkConfiguration*>(&config)); };
	void Signal_OnlineStateChanged(bool isOnline) { callbackQNetworkConfigurationManager_OnlineStateChanged(this, isOnline); };
	void Signal_UpdateCompleted() { callbackQNetworkConfigurationManager_UpdateCompleted(this); };
	void updateConfigurations() { callbackQNetworkConfigurationManager_UpdateConfigurations(this); };
	 ~MyQNetworkConfigurationManager() { callbackQNetworkConfigurationManager_DestroyQNetworkConfigurationManager(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkConfigurationManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNetworkConfigurationManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkConfigurationManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNetworkConfigurationManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkConfigurationManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkConfigurationManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkConfigurationManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNetworkConfigurationManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkConfigurationManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNetworkConfigurationManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkConfigurationManager_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNetworkConfigurationManager*)

int QNetworkConfigurationManager_QNetworkConfigurationManager_QRegisterMetaType(){qRegisterMetaType<QNetworkConfigurationManager*>(); return qRegisterMetaType<MyQNetworkConfigurationManager*>();}

void* QNetworkConfigurationManager_NewQNetworkConfigurationManager(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkConfigurationManager(static_cast<QWindow*>(parent));
	} else {
		return new MyQNetworkConfigurationManager(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QNetworkConfigurationManager_QNetworkConfigurationManager_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t784a85 = QNetworkConfigurationManager::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t784a85.prepend("WHITESPACE").constData()+10), t784a85.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkConfigurationManager_QNetworkConfigurationManager_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t7f863f = QNetworkConfigurationManager::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t7f863f.prepend("WHITESPACE").constData()+10), t7f863f.size()-10 }; });
}

void QNetworkConfigurationManager_ConnectConfigurationAdded(void* ptr)
{
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&QNetworkConfigurationManager::configurationAdded), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&MyQNetworkConfigurationManager::Signal_ConfigurationAdded));
}

void QNetworkConfigurationManager_DisconnectConfigurationAdded(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&QNetworkConfigurationManager::configurationAdded), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&MyQNetworkConfigurationManager::Signal_ConfigurationAdded));
}

void QNetworkConfigurationManager_ConfigurationAdded(void* ptr, void* config)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->configurationAdded(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkConfigurationManager_ConnectConfigurationChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&QNetworkConfigurationManager::configurationChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&MyQNetworkConfigurationManager::Signal_ConfigurationChanged));
}

void QNetworkConfigurationManager_DisconnectConfigurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&QNetworkConfigurationManager::configurationChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&MyQNetworkConfigurationManager::Signal_ConfigurationChanged));
}

void QNetworkConfigurationManager_ConfigurationChanged(void* ptr, void* config)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->configurationChanged(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkConfigurationManager_ConnectConfigurationRemoved(void* ptr)
{
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&QNetworkConfigurationManager::configurationRemoved), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&MyQNetworkConfigurationManager::Signal_ConfigurationRemoved));
}

void QNetworkConfigurationManager_DisconnectConfigurationRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&QNetworkConfigurationManager::configurationRemoved), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(const QNetworkConfiguration &)>(&MyQNetworkConfigurationManager::Signal_ConfigurationRemoved));
}

void QNetworkConfigurationManager_ConfigurationRemoved(void* ptr, void* config)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->configurationRemoved(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkConfigurationManager_ConnectOnlineStateChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));
}

void QNetworkConfigurationManager_DisconnectOnlineStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));
}

void QNetworkConfigurationManager_OnlineStateChanged(void* ptr, char isOnline)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->onlineStateChanged(isOnline != 0);
}

void QNetworkConfigurationManager_ConnectUpdateCompleted(void* ptr)
{
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));
}

void QNetworkConfigurationManager_DisconnectUpdateCompleted(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));
}

void QNetworkConfigurationManager_UpdateCompleted(void* ptr)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->updateCompleted();
}

void QNetworkConfigurationManager_UpdateConfigurations(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkConfigurationManager*>(ptr), "updateConfigurations");
}

void QNetworkConfigurationManager_UpdateConfigurationsDefault(void* ptr)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::updateConfigurations();
}

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(void* ptr)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->~QNetworkConfigurationManager();
}

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtNetwork_PackedList QNetworkConfigurationManager_AllConfigurations(void* ptr, long long filter)
{
	return ({ QList<QNetworkConfiguration>* tmpValue = new QList<QNetworkConfiguration>(static_cast<QNetworkConfigurationManager*>(ptr)->allConfigurations(static_cast<QNetworkConfiguration::StateFlag>(filter))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QNetworkConfigurationManager_ConfigurationFromIdentifier(void* ptr, struct QtNetwork_PackedString identifier)
{
	return new QNetworkConfiguration(static_cast<QNetworkConfigurationManager*>(ptr)->configurationFromIdentifier(QString::fromUtf8(identifier.data, identifier.len)));
}

void* QNetworkConfigurationManager_DefaultConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkConfigurationManager*>(ptr)->defaultConfiguration());
}

long long QNetworkConfigurationManager_Capabilities(void* ptr)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->capabilities();
}

char QNetworkConfigurationManager_IsOnline(void* ptr)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->isOnline();
}

void* QNetworkConfigurationManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::metaObject());
}

void* QNetworkConfigurationManager___allConfigurations_atList(void* ptr, int i)
{
	return new QNetworkConfiguration(({QNetworkConfiguration tmp = static_cast<QList<QNetworkConfiguration>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkConfiguration>*>(ptr)->size()-1) { static_cast<QList<QNetworkConfiguration>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkConfigurationManager___allConfigurations_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkConfiguration>*>(ptr)->append(*static_cast<QNetworkConfiguration*>(i));
}

void* QNetworkConfigurationManager___allConfigurations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkConfiguration>();
}

void* QNetworkConfigurationManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkConfigurationManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkConfigurationManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNetworkConfigurationManager___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkConfigurationManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkConfigurationManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkConfigurationManager___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkConfigurationManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkConfigurationManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkConfigurationManager___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkConfigurationManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkConfigurationManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkConfigurationManager___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkConfigurationManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkConfigurationManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNetworkConfigurationManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::event(static_cast<QEvent*>(e));
}

char QNetworkConfigurationManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNetworkConfigurationManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkConfigurationManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkConfigurationManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::customEvent(static_cast<QEvent*>(event));
}

void QNetworkConfigurationManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::deleteLater();
}

void QNetworkConfigurationManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkConfigurationManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::timerEvent(static_cast<QTimerEvent*>(event));
}

struct QtNetwork_PackedList QNetworkCookie_QNetworkCookie_ParseCookies(void* cookieString)
{
	return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(QNetworkCookie::parseCookies(*static_cast<QByteArray*>(cookieString))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QNetworkCookie_NewQNetworkCookie(void* name, void* value)
{
	return new QNetworkCookie(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

void* QNetworkCookie_NewQNetworkCookie2(void* other)
{
	return new QNetworkCookie(*static_cast<QNetworkCookie*>(other));
}

void QNetworkCookie_Normalize(void* ptr, void* url)
{
	static_cast<QNetworkCookie*>(ptr)->normalize(*static_cast<QUrl*>(url));
}

void QNetworkCookie_SetDomain(void* ptr, struct QtNetwork_PackedString domain)
{
	static_cast<QNetworkCookie*>(ptr)->setDomain(QString::fromUtf8(domain.data, domain.len));
}

void QNetworkCookie_SetExpirationDate(void* ptr, void* date)
{
	static_cast<QNetworkCookie*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(date));
}

void QNetworkCookie_SetHttpOnly(void* ptr, char enable)
{
	static_cast<QNetworkCookie*>(ptr)->setHttpOnly(enable != 0);
}

void QNetworkCookie_SetName(void* ptr, void* cookieName)
{
	static_cast<QNetworkCookie*>(ptr)->setName(*static_cast<QByteArray*>(cookieName));
}

void QNetworkCookie_SetPath(void* ptr, struct QtNetwork_PackedString path)
{
	static_cast<QNetworkCookie*>(ptr)->setPath(QString::fromUtf8(path.data, path.len));
}

void QNetworkCookie_SetSecure(void* ptr, char enable)
{
	static_cast<QNetworkCookie*>(ptr)->setSecure(enable != 0);
}

void QNetworkCookie_SetValue(void* ptr, void* value)
{
	static_cast<QNetworkCookie*>(ptr)->setValue(*static_cast<QByteArray*>(value));
}

void QNetworkCookie_Swap(void* ptr, void* other)
{
	static_cast<QNetworkCookie*>(ptr)->swap(*static_cast<QNetworkCookie*>(other));
}

void QNetworkCookie_DestroyQNetworkCookie(void* ptr)
{
	static_cast<QNetworkCookie*>(ptr)->~QNetworkCookie();
}

void* QNetworkCookie_Name(void* ptr)
{
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->name());
}

void* QNetworkCookie_ToRawForm(void* ptr, long long form)
{
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->toRawForm(static_cast<QNetworkCookie::RawForm>(form)));
}

void* QNetworkCookie_Value(void* ptr)
{
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->value());
}

void* QNetworkCookie_ExpirationDate(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCookie*>(ptr)->expirationDate());
}

struct QtNetwork_PackedString QNetworkCookie_Domain(void* ptr)
{
	return ({ QByteArray tb84845 = static_cast<QNetworkCookie*>(ptr)->domain().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb84845.prepend("WHITESPACE").constData()+10), tb84845.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkCookie_Path(void* ptr)
{
	return ({ QByteArray tc58c07 = static_cast<QNetworkCookie*>(ptr)->path().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc58c07.prepend("WHITESPACE").constData()+10), tc58c07.size()-10 }; });
}

char QNetworkCookie_HasSameIdentifier(void* ptr, void* other)
{
	return static_cast<QNetworkCookie*>(ptr)->hasSameIdentifier(*static_cast<QNetworkCookie*>(other));
}

char QNetworkCookie_IsHttpOnly(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->isHttpOnly();
}

char QNetworkCookie_IsSecure(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->isSecure();
}

char QNetworkCookie_IsSessionCookie(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->isSessionCookie();
}

void* QNetworkCookie___parseCookies_atList(void* ptr, int i)
{
	return new QNetworkCookie(({QNetworkCookie tmp = static_cast<QList<QNetworkCookie>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkCookie>*>(ptr)->size()-1) { static_cast<QList<QNetworkCookie>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkCookie___parseCookies_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkCookie>*>(ptr)->append(*static_cast<QNetworkCookie*>(i));
}

void* QNetworkCookie___parseCookies_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkCookie>();
}

class MyQNetworkCookieJar: public QNetworkCookieJar
{
public:
	MyQNetworkCookieJar(QObject *parent = Q_NULLPTR) : QNetworkCookieJar(parent) {QNetworkCookieJar_QNetworkCookieJar_QRegisterMetaType();};
	bool deleteCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_DeleteCookie(this, const_cast<QNetworkCookie*>(&cookie)) != 0; };
	bool insertCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_InsertCookie(this, const_cast<QNetworkCookie*>(&cookie)) != 0; };
	bool setCookiesFromUrl(const QList<QNetworkCookie> & cookieList, const QUrl & url) { return callbackQNetworkCookieJar_SetCookiesFromUrl(this, ({ QList<QNetworkCookie>* tmpValue = const_cast<QList<QNetworkCookie>*>(&cookieList); QtNetwork_PackedList { tmpValue, tmpValue->size() }; }), const_cast<QUrl*>(&url)) != 0; };
	bool updateCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_UpdateCookie(this, const_cast<QNetworkCookie*>(&cookie)) != 0; };
	 ~MyQNetworkCookieJar() { callbackQNetworkCookieJar_DestroyQNetworkCookieJar(this); };
	QList<QNetworkCookie> cookiesForUrl(const QUrl & url) const { return ({ QList<QNetworkCookie>* tmpP = static_cast<QList<QNetworkCookie>*>(callbackQNetworkCookieJar_CookiesForUrl(const_cast<void*>(static_cast<const void*>(this)), const_cast<QUrl*>(&url))); QList<QNetworkCookie> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	bool validateCookie(const QNetworkCookie & cookie, const QUrl & url) const { return callbackQNetworkCookieJar_ValidateCookie(const_cast<void*>(static_cast<const void*>(this)), const_cast<QNetworkCookie*>(&cookie), const_cast<QUrl*>(&url)) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkCookieJar_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNetworkCookieJar_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkCookieJar_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNetworkCookieJar_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkCookieJar_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkCookieJar_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkCookieJar_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNetworkCookieJar_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkCookieJar_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNetworkCookieJar_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkCookieJar_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNetworkCookieJar*)

int QNetworkCookieJar_QNetworkCookieJar_QRegisterMetaType(){qRegisterMetaType<QNetworkCookieJar*>(); return qRegisterMetaType<MyQNetworkCookieJar*>();}

void* QNetworkCookieJar_NewQNetworkCookieJar(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkCookieJar(static_cast<QWindow*>(parent));
	} else {
		return new MyQNetworkCookieJar(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QNetworkCookieJar_QNetworkCookieJar_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t8553ed = QNetworkCookieJar::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t8553ed.prepend("WHITESPACE").constData()+10), t8553ed.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkCookieJar_QNetworkCookieJar_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t634e2c = QNetworkCookieJar::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t634e2c.prepend("WHITESPACE").constData()+10), t634e2c.size()-10 }; });
}

char QNetworkCookieJar_DeleteCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_DeleteCookieDefault(void* ptr, void* cookie)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_InsertCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_InsertCookieDefault(void* ptr, void* cookie)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_SetCookiesFromUrl(void* ptr, void* cookieList, void* url)
{
	return static_cast<QNetworkCookieJar*>(ptr)->setCookiesFromUrl(*static_cast<QList<QNetworkCookie>*>(cookieList), *static_cast<QUrl*>(url));
}

char QNetworkCookieJar_SetCookiesFromUrlDefault(void* ptr, void* cookieList, void* url)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::setCookiesFromUrl(*static_cast<QList<QNetworkCookie>*>(cookieList), *static_cast<QUrl*>(url));
}

char QNetworkCookieJar_UpdateCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_UpdateCookieDefault(void* ptr, void* cookie)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

void QNetworkCookieJar_SetAllCookies(void* ptr, void* cookieList)
{
	static_cast<QNetworkCookieJar*>(ptr)->setAllCookies(*static_cast<QList<QNetworkCookie>*>(cookieList));
}

void QNetworkCookieJar_DestroyQNetworkCookieJar(void* ptr)
{
	static_cast<QNetworkCookieJar*>(ptr)->~QNetworkCookieJar();
}

void QNetworkCookieJar_DestroyQNetworkCookieJarDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtNetwork_PackedList QNetworkCookieJar_AllCookies(void* ptr)
{
	return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(static_cast<QNetworkCookieJar*>(ptr)->allCookies()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkCookieJar_CookiesForUrl(void* ptr, void* url)
{
	return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(static_cast<QNetworkCookieJar*>(ptr)->cookiesForUrl(*static_cast<QUrl*>(url))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkCookieJar_CookiesForUrlDefault(void* ptr, void* url)
{
		return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::cookiesForUrl(*static_cast<QUrl*>(url))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QNetworkCookieJar_ValidateCookie(void* ptr, void* cookie, void* url)
{
	return static_cast<QNetworkCookieJar*>(ptr)->validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

char QNetworkCookieJar_ValidateCookieDefault(void* ptr, void* cookie, void* url)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

void* QNetworkCookieJar_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::metaObject());
}

void* QNetworkCookieJar___setCookiesFromUrl_cookieList_atList(void* ptr, int i)
{
	return new QNetworkCookie(({QNetworkCookie tmp = static_cast<QList<QNetworkCookie>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkCookie>*>(ptr)->size()-1) { static_cast<QList<QNetworkCookie>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkCookieJar___setCookiesFromUrl_cookieList_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkCookie>*>(ptr)->append(*static_cast<QNetworkCookie*>(i));
}

void* QNetworkCookieJar___setCookiesFromUrl_cookieList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkCookie>();
}

void* QNetworkCookieJar___setAllCookies_cookieList_atList(void* ptr, int i)
{
	return new QNetworkCookie(({QNetworkCookie tmp = static_cast<QList<QNetworkCookie>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkCookie>*>(ptr)->size()-1) { static_cast<QList<QNetworkCookie>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkCookieJar___setAllCookies_cookieList_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkCookie>*>(ptr)->append(*static_cast<QNetworkCookie*>(i));
}

void* QNetworkCookieJar___setAllCookies_cookieList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkCookie>();
}

void* QNetworkCookieJar___allCookies_atList(void* ptr, int i)
{
	return new QNetworkCookie(({QNetworkCookie tmp = static_cast<QList<QNetworkCookie>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkCookie>*>(ptr)->size()-1) { static_cast<QList<QNetworkCookie>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkCookieJar___allCookies_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkCookie>*>(ptr)->append(*static_cast<QNetworkCookie*>(i));
}

void* QNetworkCookieJar___allCookies_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkCookie>();
}

void* QNetworkCookieJar___cookiesForUrl_atList(void* ptr, int i)
{
	return new QNetworkCookie(({QNetworkCookie tmp = static_cast<QList<QNetworkCookie>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkCookie>*>(ptr)->size()-1) { static_cast<QList<QNetworkCookie>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkCookieJar___cookiesForUrl_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkCookie>*>(ptr)->append(*static_cast<QNetworkCookie*>(i));
}

void* QNetworkCookieJar___cookiesForUrl_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkCookie>();
}

void* QNetworkCookieJar___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkCookieJar___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkCookieJar___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNetworkCookieJar___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkCookieJar___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkCookieJar___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkCookieJar___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkCookieJar___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkCookieJar___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkCookieJar___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkCookieJar___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkCookieJar___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkCookieJar___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkCookieJar___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkCookieJar___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNetworkCookieJar_EventDefault(void* ptr, void* e)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::event(static_cast<QEvent*>(e));
}

char QNetworkCookieJar_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNetworkCookieJar_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkCookieJar_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkCookieJar_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::customEvent(static_cast<QEvent*>(event));
}

void QNetworkCookieJar_DeleteLaterDefault(void* ptr)
{
		static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::deleteLater();
}

void QNetworkCookieJar_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkCookieJar_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QNetworkDatagram_MakeReply2(void* ptr, void* payload)
{
	return new QNetworkDatagram(static_cast<QNetworkDatagram*>(ptr)->makeReply(*static_cast<QByteArray*>(payload)));
}

void* QNetworkDatagram_NewQNetworkDatagram()
{
	return new QNetworkDatagram();
}

void* QNetworkDatagram_NewQNetworkDatagram4(void* other)
{
	return new QNetworkDatagram(*static_cast<QNetworkDatagram*>(other));
}

void* QNetworkDatagram_NewQNetworkDatagram2(void* data, void* destinationAddress, unsigned short port)
{
	return new QNetworkDatagram(*static_cast<QByteArray*>(data), *static_cast<QHostAddress*>(destinationAddress), port);
}

void* QNetworkDatagram_NewQNetworkDatagram3(void* other)
{
	return new QNetworkDatagram(*static_cast<QNetworkDatagram*>(other));
}

void QNetworkDatagram_Clear(void* ptr)
{
	static_cast<QNetworkDatagram*>(ptr)->clear();
}

void QNetworkDatagram_SetData(void* ptr, void* data)
{
	static_cast<QNetworkDatagram*>(ptr)->setData(*static_cast<QByteArray*>(data));
}

void QNetworkDatagram_SetDestination(void* ptr, void* address, unsigned short port)
{
	static_cast<QNetworkDatagram*>(ptr)->setDestination(*static_cast<QHostAddress*>(address), port);
}

void QNetworkDatagram_SetHopLimit(void* ptr, int count)
{
	static_cast<QNetworkDatagram*>(ptr)->setHopLimit(count);
}

void QNetworkDatagram_SetInterfaceIndex(void* ptr, unsigned int index)
{
	static_cast<QNetworkDatagram*>(ptr)->setInterfaceIndex(index);
}

void QNetworkDatagram_SetSender(void* ptr, void* address, unsigned short port)
{
	static_cast<QNetworkDatagram*>(ptr)->setSender(*static_cast<QHostAddress*>(address), port);
}

void QNetworkDatagram_Swap(void* ptr, void* other)
{
	static_cast<QNetworkDatagram*>(ptr)->swap(*static_cast<QNetworkDatagram*>(other));
}

void QNetworkDatagram_DestroyQNetworkDatagram(void* ptr)
{
	static_cast<QNetworkDatagram*>(ptr)->~QNetworkDatagram();
}

void* QNetworkDatagram_Data(void* ptr)
{
	return new QByteArray(static_cast<QNetworkDatagram*>(ptr)->data());
}

void* QNetworkDatagram_DestinationAddress(void* ptr)
{
	return new QHostAddress(static_cast<QNetworkDatagram*>(ptr)->destinationAddress());
}

void* QNetworkDatagram_SenderAddress(void* ptr)
{
	return new QHostAddress(static_cast<QNetworkDatagram*>(ptr)->senderAddress());
}

void* QNetworkDatagram_MakeReply(void* ptr, void* payload)
{
	return new QNetworkDatagram(static_cast<QNetworkDatagram*>(ptr)->makeReply(*static_cast<QByteArray*>(payload)));
}

char QNetworkDatagram_IsNull(void* ptr)
{
	return static_cast<QNetworkDatagram*>(ptr)->isNull();
}

char QNetworkDatagram_IsValid(void* ptr)
{
	return static_cast<QNetworkDatagram*>(ptr)->isValid();
}

int QNetworkDatagram_DestinationPort(void* ptr)
{
	return static_cast<QNetworkDatagram*>(ptr)->destinationPort();
}

int QNetworkDatagram_HopLimit(void* ptr)
{
	return static_cast<QNetworkDatagram*>(ptr)->hopLimit();
}

int QNetworkDatagram_SenderPort(void* ptr)
{
	return static_cast<QNetworkDatagram*>(ptr)->senderPort();
}

unsigned int QNetworkDatagram_InterfaceIndex(void* ptr)
{
	return static_cast<QNetworkDatagram*>(ptr)->interfaceIndex();
}

class MyQNetworkDiskCache: public QNetworkDiskCache
{
public:
	MyQNetworkDiskCache(QObject *parent = Q_NULLPTR) : QNetworkDiskCache(parent) {QNetworkDiskCache_QNetworkDiskCache_QRegisterMetaType();};
	QIODevice * data(const QUrl & url) { return static_cast<QIODevice*>(callbackQNetworkDiskCache_Data(this, const_cast<QUrl*>(&url))); };
	QIODevice * prepare(const QNetworkCacheMetaData & metaData) { return static_cast<QIODevice*>(callbackQNetworkDiskCache_Prepare(this, const_cast<QNetworkCacheMetaData*>(&metaData))); };
	QNetworkCacheMetaData metaData(const QUrl & url) { return *static_cast<QNetworkCacheMetaData*>(callbackQNetworkDiskCache_MetaData(this, const_cast<QUrl*>(&url))); };
	bool remove(const QUrl & url) { return callbackQNetworkDiskCache_Remove(this, const_cast<QUrl*>(&url)) != 0; };
	qint64 expire() { return callbackQNetworkDiskCache_Expire(this); };
	void clear() { callbackQNetworkDiskCache_Clear(this); };
	void insert(QIODevice * device) { callbackQNetworkDiskCache_Insert(this, device); };
	void updateMetaData(const QNetworkCacheMetaData & metaData) { callbackQNetworkDiskCache_UpdateMetaData(this, const_cast<QNetworkCacheMetaData*>(&metaData)); };
	 ~MyQNetworkDiskCache() { callbackQNetworkDiskCache_DestroyQNetworkDiskCache(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractNetworkCache_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	qint64 cacheSize() const { return callbackQNetworkDiskCache_CacheSize(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQAbstractNetworkCache_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractNetworkCache_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractNetworkCache_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractNetworkCache_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractNetworkCache_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractNetworkCache_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractNetworkCache_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractNetworkCache_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNetworkDiskCache*)

int QNetworkDiskCache_QNetworkDiskCache_QRegisterMetaType(){qRegisterMetaType<QNetworkDiskCache*>(); return qRegisterMetaType<MyQNetworkDiskCache*>();}

void* QNetworkDiskCache_Data(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void* QNetworkDiskCache_DataDefault(void* ptr, void* url)
{
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::data(*static_cast<QUrl*>(url));
}

void* QNetworkDiskCache_Prepare(void* ptr, void* metaData)
{
	return static_cast<QNetworkDiskCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void* QNetworkDiskCache_PrepareDefault(void* ptr, void* metaData)
{
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void* QNetworkDiskCache_MetaData(void* ptr, void* url)
{
	return new QNetworkCacheMetaData(static_cast<QNetworkDiskCache*>(ptr)->metaData(*static_cast<QUrl*>(url)));
}

void* QNetworkDiskCache_MetaDataDefault(void* ptr, void* url)
{
		return new QNetworkCacheMetaData(static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::metaData(*static_cast<QUrl*>(url)));
}

void* QNetworkDiskCache_NewQNetworkDiskCache(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkDiskCache(static_cast<QWindow*>(parent));
	} else {
		return new MyQNetworkDiskCache(static_cast<QObject*>(parent));
	}
}

char QNetworkDiskCache_Remove(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->remove(*static_cast<QUrl*>(url));
}

char QNetworkDiskCache_RemoveDefault(void* ptr, void* url)
{
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::remove(*static_cast<QUrl*>(url));
}

long long QNetworkDiskCache_Expire(void* ptr)
{
	return static_cast<QNetworkDiskCache*>(ptr)->expire();
}

long long QNetworkDiskCache_ExpireDefault(void* ptr)
{
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::expire();
}

void QNetworkDiskCache_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkDiskCache*>(ptr), "clear");
}

void QNetworkDiskCache_ClearDefault(void* ptr)
{
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::clear();
}

void QNetworkDiskCache_Insert(void* ptr, void* device)
{
	static_cast<QNetworkDiskCache*>(ptr)->insert(static_cast<QIODevice*>(device));
}

void QNetworkDiskCache_InsertDefault(void* ptr, void* device)
{
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::insert(static_cast<QIODevice*>(device));
}

void QNetworkDiskCache_SetCacheDirectory(void* ptr, struct QtNetwork_PackedString cacheDir)
{
	static_cast<QNetworkDiskCache*>(ptr)->setCacheDirectory(QString::fromUtf8(cacheDir.data, cacheDir.len));
}

void QNetworkDiskCache_SetMaximumCacheSize(void* ptr, long long size)
{
	static_cast<QNetworkDiskCache*>(ptr)->setMaximumCacheSize(size);
}

void QNetworkDiskCache_UpdateMetaData(void* ptr, void* metaData)
{
	static_cast<QNetworkDiskCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_UpdateMetaDataDefault(void* ptr, void* metaData)
{
		static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_DestroyQNetworkDiskCache(void* ptr)
{
	static_cast<QNetworkDiskCache*>(ptr)->~QNetworkDiskCache();
}

void QNetworkDiskCache_DestroyQNetworkDiskCacheDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNetworkDiskCache_FileMetaData(void* ptr, struct QtNetwork_PackedString fileName)
{
	return new QNetworkCacheMetaData(static_cast<QNetworkDiskCache*>(ptr)->fileMetaData(QString::fromUtf8(fileName.data, fileName.len)));
}

struct QtNetwork_PackedString QNetworkDiskCache_CacheDirectory(void* ptr)
{
	return ({ QByteArray t85cfc0 = static_cast<QNetworkDiskCache*>(ptr)->cacheDirectory().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t85cfc0.prepend("WHITESPACE").constData()+10), t85cfc0.size()-10 }; });
}

long long QNetworkDiskCache_CacheSize(void* ptr)
{
	return static_cast<QNetworkDiskCache*>(ptr)->cacheSize();
}

long long QNetworkDiskCache_CacheSizeDefault(void* ptr)
{
		return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::cacheSize();
}

long long QNetworkDiskCache_MaximumCacheSize(void* ptr)
{
	return static_cast<QNetworkDiskCache*>(ptr)->maximumCacheSize();
}

struct QtNetwork_PackedList QNetworkInterface_QNetworkInterface_AllAddresses()
{
	return ({ QList<QHostAddress>* tmpValue = new QList<QHostAddress>(QNetworkInterface::allAddresses()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkInterface_QNetworkInterface_AllInterfaces()
{
	return ({ QList<QNetworkInterface>* tmpValue = new QList<QNetworkInterface>(QNetworkInterface::allInterfaces()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QNetworkInterface_QNetworkInterface_InterfaceFromIndex(int index)
{
	return new QNetworkInterface(QNetworkInterface::interfaceFromIndex(index));
}

void* QNetworkInterface_QNetworkInterface_InterfaceFromName(struct QtNetwork_PackedString name)
{
	return new QNetworkInterface(QNetworkInterface::interfaceFromName(QString::fromUtf8(name.data, name.len)));
}

void* QNetworkInterface_NewQNetworkInterface()
{
	return new QNetworkInterface();
}

void* QNetworkInterface_NewQNetworkInterface2(void* other)
{
	return new QNetworkInterface(*static_cast<QNetworkInterface*>(other));
}

struct QtNetwork_PackedString QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(int index)
{
	return ({ QByteArray ta95340 = QNetworkInterface::interfaceNameFromIndex(index).toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta95340.prepend("WHITESPACE").constData()+10), ta95340.size()-10 }; });
}

int QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(struct QtNetwork_PackedString name)
{
	return QNetworkInterface::interfaceIndexFromName(QString::fromUtf8(name.data, name.len));
}

void QNetworkInterface_Swap(void* ptr, void* other)
{
	static_cast<QNetworkInterface*>(ptr)->swap(*static_cast<QNetworkInterface*>(other));
}

void QNetworkInterface_DestroyQNetworkInterface(void* ptr)
{
	static_cast<QNetworkInterface*>(ptr)->~QNetworkInterface();
}

struct QtNetwork_PackedList QNetworkInterface_AddressEntries(void* ptr)
{
	return ({ QList<QNetworkAddressEntry>* tmpValue = new QList<QNetworkAddressEntry>(static_cast<QNetworkInterface*>(ptr)->addressEntries()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkInterface_Flags(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->flags();
}

long long QNetworkInterface_Type(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->type();
}

struct QtNetwork_PackedString QNetworkInterface_HardwareAddress(void* ptr)
{
	return ({ QByteArray t25386c = static_cast<QNetworkInterface*>(ptr)->hardwareAddress().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t25386c.prepend("WHITESPACE").constData()+10), t25386c.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkInterface_HumanReadableName(void* ptr)
{
	return ({ QByteArray tebd539 = static_cast<QNetworkInterface*>(ptr)->humanReadableName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tebd539.prepend("WHITESPACE").constData()+10), tebd539.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkInterface_Name(void* ptr)
{
	return ({ QByteArray tb9dead = static_cast<QNetworkInterface*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb9dead.prepend("WHITESPACE").constData()+10), tb9dead.size()-10 }; });
}

char QNetworkInterface_IsValid(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->isValid();
}

int QNetworkInterface_Index(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->index();
}

int QNetworkInterface_MaximumTransmissionUnit(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->maximumTransmissionUnit();
}

void* QNetworkInterface___allAddresses_atList(void* ptr, int i)
{
	return new QHostAddress(({QHostAddress tmp = static_cast<QList<QHostAddress>*>(ptr)->at(i); if (i == static_cast<QList<QHostAddress>*>(ptr)->size()-1) { static_cast<QList<QHostAddress>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkInterface___allAddresses_setList(void* ptr, void* i)
{
	static_cast<QList<QHostAddress>*>(ptr)->append(*static_cast<QHostAddress*>(i));
}

void* QNetworkInterface___allAddresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHostAddress>();
}

void* QNetworkInterface___allInterfaces_atList(void* ptr, int i)
{
	return new QNetworkInterface(({QNetworkInterface tmp = static_cast<QList<QNetworkInterface>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkInterface>*>(ptr)->size()-1) { static_cast<QList<QNetworkInterface>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkInterface___allInterfaces_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkInterface>*>(ptr)->append(*static_cast<QNetworkInterface*>(i));
}

void* QNetworkInterface___allInterfaces_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkInterface>();
}

void* QNetworkInterface___addressEntries_atList(void* ptr, int i)
{
	return new QNetworkAddressEntry(({QNetworkAddressEntry tmp = static_cast<QList<QNetworkAddressEntry>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkAddressEntry>*>(ptr)->size()-1) { static_cast<QList<QNetworkAddressEntry>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkInterface___addressEntries_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkAddressEntry>*>(ptr)->append(*static_cast<QNetworkAddressEntry*>(i));
}

void* QNetworkInterface___addressEntries_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkAddressEntry>();
}

void* QNetworkProxy_QNetworkProxy_ApplicationProxy()
{
	return new QNetworkProxy(QNetworkProxy::applicationProxy());
}

void* QNetworkProxy_NewQNetworkProxy()
{
	return new QNetworkProxy();
}

void* QNetworkProxy_NewQNetworkProxy2(long long ty, struct QtNetwork_PackedString hostName, unsigned short port, struct QtNetwork_PackedString user, struct QtNetwork_PackedString password)
{
	return new QNetworkProxy(static_cast<QNetworkProxy::ProxyType>(ty), QString::fromUtf8(hostName.data, hostName.len), port, QString::fromUtf8(user.data, user.len), QString::fromUtf8(password.data, password.len));
}

void* QNetworkProxy_NewQNetworkProxy3(void* other)
{
	return new QNetworkProxy(*static_cast<QNetworkProxy*>(other));
}

void QNetworkProxy_QNetworkProxy_SetApplicationProxy(void* networkProxy)
{
	QNetworkProxy::setApplicationProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QNetworkProxy_SetCapabilities(void* ptr, long long capabilities)
{
	static_cast<QNetworkProxy*>(ptr)->setCapabilities(static_cast<QNetworkProxy::Capability>(capabilities));
}

void QNetworkProxy_SetHeader(void* ptr, long long header, void* value)
{
	static_cast<QNetworkProxy*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkProxy_SetHostName(void* ptr, struct QtNetwork_PackedString hostName)
{
	static_cast<QNetworkProxy*>(ptr)->setHostName(QString::fromUtf8(hostName.data, hostName.len));
}

void QNetworkProxy_SetPassword(void* ptr, struct QtNetwork_PackedString password)
{
	static_cast<QNetworkProxy*>(ptr)->setPassword(QString::fromUtf8(password.data, password.len));
}

void QNetworkProxy_SetPort(void* ptr, unsigned short port)
{
	static_cast<QNetworkProxy*>(ptr)->setPort(port);
}

void QNetworkProxy_SetRawHeader(void* ptr, void* headerName, void* headerValue)
{
	static_cast<QNetworkProxy*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkProxy_SetType(void* ptr, long long ty)
{
	static_cast<QNetworkProxy*>(ptr)->setType(static_cast<QNetworkProxy::ProxyType>(ty));
}

void QNetworkProxy_SetUser(void* ptr, struct QtNetwork_PackedString user)
{
	static_cast<QNetworkProxy*>(ptr)->setUser(QString::fromUtf8(user.data, user.len));
}

void QNetworkProxy_Swap(void* ptr, void* other)
{
	static_cast<QNetworkProxy*>(ptr)->swap(*static_cast<QNetworkProxy*>(other));
}

void QNetworkProxy_DestroyQNetworkProxy(void* ptr)
{
	static_cast<QNetworkProxy*>(ptr)->~QNetworkProxy();
}

void* QNetworkProxy_RawHeader(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QNetworkProxy*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

struct QtNetwork_PackedList QNetworkProxy_RawHeaderList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QNetworkProxy*>(ptr)->rawHeaderList()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkProxy_Capabilities(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->capabilities();
}

long long QNetworkProxy_Type(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->type();
}

struct QtNetwork_PackedString QNetworkProxy_HostName(void* ptr)
{
	return ({ QByteArray t422f46 = static_cast<QNetworkProxy*>(ptr)->hostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t422f46.prepend("WHITESPACE").constData()+10), t422f46.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkProxy_Password(void* ptr)
{
	return ({ QByteArray t6e003a = static_cast<QNetworkProxy*>(ptr)->password().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t6e003a.prepend("WHITESPACE").constData()+10), t6e003a.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkProxy_User(void* ptr)
{
	return ({ QByteArray tcd9fd7 = static_cast<QNetworkProxy*>(ptr)->user().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tcd9fd7.prepend("WHITESPACE").constData()+10), tcd9fd7.size()-10 }; });
}

void* QNetworkProxy_Header(void* ptr, long long header)
{
	return new QVariant(static_cast<QNetworkProxy*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

char QNetworkProxy_HasRawHeader(void* ptr, void* headerName)
{
	return static_cast<QNetworkProxy*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

char QNetworkProxy_IsCachingProxy(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->isCachingProxy();
}

char QNetworkProxy_IsTransparentProxy(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->isTransparentProxy();
}

unsigned short QNetworkProxy_Port(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->port();
}

void* QNetworkProxy___rawHeaderList_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkProxy___rawHeaderList_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkProxy___rawHeaderList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

class MyQNetworkProxyFactory: public QNetworkProxyFactory
{
public:
	MyQNetworkProxyFactory() : QNetworkProxyFactory() {};
	QList<QNetworkProxy> queryProxy(const QNetworkProxyQuery & query) { return ({ QList<QNetworkProxy>* tmpP = static_cast<QList<QNetworkProxy>*>(callbackQNetworkProxyFactory_QueryProxy(this, const_cast<QNetworkProxyQuery*>(&query))); QList<QNetworkProxy> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	 ~MyQNetworkProxyFactory() { callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory(this); };
};

struct QtNetwork_PackedList QNetworkProxyFactory_QNetworkProxyFactory_ProxyForQuery(void* query)
{
	return ({ QList<QNetworkProxy>* tmpValue = new QList<QNetworkProxy>(QNetworkProxyFactory::proxyForQuery(*static_cast<QNetworkProxyQuery*>(query))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkProxyFactory_QueryProxy(void* ptr, void* query)
{
	return ({ QList<QNetworkProxy>* tmpValue = new QList<QNetworkProxy>(static_cast<QNetworkProxyFactory*>(ptr)->queryProxy(*static_cast<QNetworkProxyQuery*>(query))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkProxyFactory_QNetworkProxyFactory_SystemProxyForQuery(void* query)
{
	return ({ QList<QNetworkProxy>* tmpValue = new QList<QNetworkProxy>(QNetworkProxyFactory::systemProxyForQuery(*static_cast<QNetworkProxyQuery*>(query))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QNetworkProxyFactory_NewQNetworkProxyFactory()
{
	return new MyQNetworkProxyFactory();
}

char QNetworkProxyFactory_QNetworkProxyFactory_UsesSystemConfiguration()
{
	return QNetworkProxyFactory::usesSystemConfiguration();
}

void QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(void* factory)
{
	QNetworkProxyFactory::setApplicationProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

void QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(char enable)
{
	QNetworkProxyFactory::setUseSystemConfiguration(enable != 0);
}

void QNetworkProxyFactory_DestroyQNetworkProxyFactory(void* ptr)
{
	static_cast<QNetworkProxyFactory*>(ptr)->~QNetworkProxyFactory();
}

void QNetworkProxyFactory_DestroyQNetworkProxyFactoryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNetworkProxyFactory___proxyForQuery_atList(void* ptr, int i)
{
	return new QNetworkProxy(({QNetworkProxy tmp = static_cast<QList<QNetworkProxy>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkProxy>*>(ptr)->size()-1) { static_cast<QList<QNetworkProxy>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkProxyFactory___proxyForQuery_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkProxy>*>(ptr)->append(*static_cast<QNetworkProxy*>(i));
}

void* QNetworkProxyFactory___proxyForQuery_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkProxy>();
}

void* QNetworkProxyFactory___queryProxy_atList(void* ptr, int i)
{
	return new QNetworkProxy(({QNetworkProxy tmp = static_cast<QList<QNetworkProxy>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkProxy>*>(ptr)->size()-1) { static_cast<QList<QNetworkProxy>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkProxyFactory___queryProxy_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkProxy>*>(ptr)->append(*static_cast<QNetworkProxy*>(i));
}

void* QNetworkProxyFactory___queryProxy_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkProxy>();
}

void* QNetworkProxyFactory___systemProxyForQuery_atList(void* ptr, int i)
{
	return new QNetworkProxy(({QNetworkProxy tmp = static_cast<QList<QNetworkProxy>*>(ptr)->at(i); if (i == static_cast<QList<QNetworkProxy>*>(ptr)->size()-1) { static_cast<QList<QNetworkProxy>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkProxyFactory___systemProxyForQuery_setList(void* ptr, void* i)
{
	static_cast<QList<QNetworkProxy>*>(ptr)->append(*static_cast<QNetworkProxy*>(i));
}

void* QNetworkProxyFactory___systemProxyForQuery_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QNetworkProxy>();
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery()
{
	return new QNetworkProxyQuery();
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery8(void* other)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkProxyQuery*>(other));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery3(struct QtNetwork_PackedString hostname, int port, struct QtNetwork_PackedString protocolTag, long long queryType)
{
	return new QNetworkProxyQuery(QString::fromUtf8(hostname.data, hostname.len), port, QString::fromUtf8(protocolTag.data, protocolTag.len), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery2(void* requestUrl, long long queryType)
{
	return new QNetworkProxyQuery(*static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery4(unsigned short bindPort, struct QtNetwork_PackedString protocolTag, long long queryType)
{
	return new QNetworkProxyQuery(bindPort, QString::fromUtf8(protocolTag.data, protocolTag.len), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void QNetworkProxyQuery_SetLocalPort(void* ptr, int port)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setLocalPort(port);
}

void QNetworkProxyQuery_SetPeerHostName(void* ptr, struct QtNetwork_PackedString hostname)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerHostName(QString::fromUtf8(hostname.data, hostname.len));
}

void QNetworkProxyQuery_SetPeerPort(void* ptr, int port)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerPort(port);
}

void QNetworkProxyQuery_SetProtocolTag(void* ptr, struct QtNetwork_PackedString protocolTag)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setProtocolTag(QString::fromUtf8(protocolTag.data, protocolTag.len));
}

void QNetworkProxyQuery_SetQueryType(void* ptr, long long ty)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setQueryType(static_cast<QNetworkProxyQuery::QueryType>(ty));
}

void QNetworkProxyQuery_SetUrl(void* ptr, void* url)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkProxyQuery_Swap(void* ptr, void* other)
{
	static_cast<QNetworkProxyQuery*>(ptr)->swap(*static_cast<QNetworkProxyQuery*>(other));
}

void QNetworkProxyQuery_DestroyQNetworkProxyQuery(void* ptr)
{
	static_cast<QNetworkProxyQuery*>(ptr)->~QNetworkProxyQuery();
}

long long QNetworkProxyQuery_QueryType(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->queryType();
}

struct QtNetwork_PackedString QNetworkProxyQuery_PeerHostName(void* ptr)
{
	return ({ QByteArray t878927 = static_cast<QNetworkProxyQuery*>(ptr)->peerHostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t878927.prepend("WHITESPACE").constData()+10), t878927.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkProxyQuery_ProtocolTag(void* ptr)
{
	return ({ QByteArray teab311 = static_cast<QNetworkProxyQuery*>(ptr)->protocolTag().toUtf8(); QtNetwork_PackedString { const_cast<char*>(teab311.prepend("WHITESPACE").constData()+10), teab311.size()-10 }; });
}

void* QNetworkProxyQuery_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkProxyQuery*>(ptr)->url());
}

int QNetworkProxyQuery_LocalPort(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->localPort();
}

int QNetworkProxyQuery_PeerPort(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->peerPort();
}

class MyQNetworkReply: public QNetworkReply
{
public:
	MyQNetworkReply(QObject *parent = Q_NULLPTR) : QNetworkReply(parent) {QNetworkReply_QNetworkReply_QRegisterMetaType();};
	void abort() { callbackQNetworkReply_Abort(this); };
	void close() { callbackQNetworkReply_Close(this); };
	void Signal_DownloadProgress(qint64 bytesReceived, qint64 bytesTotal) { callbackQNetworkReply_DownloadProgress(this, bytesReceived, bytesTotal); };
	void Signal_Encrypted() { callbackQNetworkReply_Encrypted(this); };
	void Signal_Error2(QNetworkReply::NetworkError code) { callbackQNetworkReply_Error2(this, code); };
	void Signal_Finished() { callbackQNetworkReply_Finished(this); };
	void ignoreSslErrors() { callbackQNetworkReply_IgnoreSslErrors(this); };
	void ignoreSslErrorsImplementation(const QList<QSslError> & errors) { callbackQNetworkReply_IgnoreSslErrorsImplementation(this, ({ QList<QSslError>* tmpValue = const_cast<QList<QSslError>*>(&errors); QtNetwork_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_MetaDataChanged() { callbackQNetworkReply_MetaDataChanged(this); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkReply_PreSharedKeyAuthenticationRequired(this, authenticator); };
	void Signal_RedirectAllowed() { callbackQNetworkReply_RedirectAllowed(this); };
	void Signal_Redirected(const QUrl & url) { callbackQNetworkReply_Redirected(this, const_cast<QUrl*>(&url)); };
	void setReadBufferSize(qint64 size) { callbackQNetworkReply_SetReadBufferSize(this, size); };
	void setSslConfigurationImplementation(const QSslConfiguration & configuration) { callbackQNetworkReply_SetSslConfigurationImplementation(this, const_cast<QSslConfiguration*>(&configuration)); };
	void Signal_SslErrors(const QList<QSslError> & errors) { callbackQNetworkReply_SslErrors(this, ({ QList<QSslError>* tmpValue = const_cast<QList<QSslError>*>(&errors); QtNetwork_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_UploadProgress(qint64 bytesSent, qint64 bytesTotal) { callbackQNetworkReply_UploadProgress(this, bytesSent, bytesTotal); };
	 ~MyQNetworkReply() { callbackQNetworkReply_DestroyQNetworkReply(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkReply_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void sslConfigurationImplementation(QSslConfiguration & configuration) const { callbackQNetworkReply_SslConfigurationImplementation(const_cast<void*>(static_cast<const void*>(this)), static_cast<QSslConfiguration*>(&configuration)); };
	bool open(QIODevice::OpenMode mode) { return callbackQNetworkReply_Open(this, mode) != 0; };
	bool reset() { return callbackQNetworkReply_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQNetworkReply_Seek(this, pos) != 0; };
	bool waitForBytesWritten(int msecs) { return callbackQNetworkReply_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQNetworkReply_WaitForReadyRead(this, msecs) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQNetworkReply_ReadData(this, dataPacked, maxSize); };
	qint64 readLineData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQNetworkReply_ReadLineData(this, dataPacked, maxSize); };
	qint64 writeData(const char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), maxSize };return callbackQNetworkReply_WriteData(this, dataPacked, maxSize); };
	void Signal_AboutToClose() { callbackQNetworkReply_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQNetworkReply_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQNetworkReply_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQNetworkReply_ChannelReadyRead(this, channel); };
	void Signal_ReadChannelFinished() { callbackQNetworkReply_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQNetworkReply_ReadyRead(this); };
	bool atEnd() const { return callbackQNetworkReply_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool canReadLine() const { return callbackQNetworkReply_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isSequential() const { return callbackQNetworkReply_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	qint64 bytesAvailable() const { return callbackQNetworkReply_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQNetworkReply_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 pos() const { return callbackQNetworkReply_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQNetworkReply_Size(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQNetworkReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkReply_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNetworkReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkReply_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkReply_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNetworkReply_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNetworkReply_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkReply_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNetworkReply*)

int QNetworkReply_QNetworkReply_QRegisterMetaType(){qRegisterMetaType<QNetworkReply*>(); return qRegisterMetaType<MyQNetworkReply*>();}

void* QNetworkReply_NewQNetworkReply(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkReply(static_cast<QWindow*>(parent));
	} else {
		return new MyQNetworkReply(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QNetworkReply_QNetworkReply_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t809fbb = QNetworkReply::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t809fbb.prepend("WHITESPACE").constData()+10), t809fbb.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkReply_QNetworkReply_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tf07360 = QNetworkReply::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(tf07360.prepend("WHITESPACE").constData()+10), tf07360.size()-10 }; });
}

void QNetworkReply_Abort(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "abort");
}

void QNetworkReply_CloseDefault(void* ptr)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::close();
}

void QNetworkReply_ConnectDownloadProgress(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::downloadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_DownloadProgress));
}

void QNetworkReply_DisconnectDownloadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::downloadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_DownloadProgress));
}

void QNetworkReply_DownloadProgress(void* ptr, long long bytesReceived, long long bytesTotal)
{
	static_cast<QNetworkReply*>(ptr)->downloadProgress(bytesReceived, bytesTotal);
}

void QNetworkReply_ConnectEncrypted(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));
}

void QNetworkReply_DisconnectEncrypted(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));
}

void QNetworkReply_Encrypted(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->encrypted();
}

void QNetworkReply_ConnectError2(void* ptr)
{
	qRegisterMetaType<QNetworkReply::NetworkError>();
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QNetworkReply::NetworkError)>(&QNetworkReply::error), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QNetworkReply::NetworkError)>(&MyQNetworkReply::Signal_Error2));
}

void QNetworkReply_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QNetworkReply::NetworkError)>(&QNetworkReply::error), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QNetworkReply::NetworkError)>(&MyQNetworkReply::Signal_Error2));
}

void QNetworkReply_Error2(void* ptr, long long code)
{
	static_cast<QNetworkReply*>(ptr)->error(static_cast<QNetworkReply::NetworkError>(code));
}

void QNetworkReply_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));
}

void QNetworkReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));
}

void QNetworkReply_Finished(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->finished();
}

void QNetworkReply_IgnoreSslErrors(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "ignoreSslErrors");
}

void QNetworkReply_IgnoreSslErrorsDefault(void* ptr)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::ignoreSslErrors();
}

void QNetworkReply_IgnoreSslErrors2(void* ptr, void* errors)
{
	static_cast<QNetworkReply*>(ptr)->ignoreSslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QNetworkReply_IgnoreSslErrorsImplementation(void* ptr, void* errors)
{
	static_cast<QNetworkReply*>(ptr)->ignoreSslErrorsImplementation(*static_cast<QList<QSslError>*>(errors));
}

void QNetworkReply_IgnoreSslErrorsImplementationDefault(void* ptr, void* errors)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::ignoreSslErrorsImplementation(*static_cast<QList<QSslError>*>(errors));
}

void QNetworkReply_ConnectMetaDataChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));
}

void QNetworkReply_DisconnectMetaDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));
}

void QNetworkReply_MetaDataChanged(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->metaDataChanged();
}

void QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));
}

void QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));
}

void QNetworkReply_PreSharedKeyAuthenticationRequired(void* ptr, void* authenticator)
{
	static_cast<QNetworkReply*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QNetworkReply_ConnectRedirectAllowed(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::redirectAllowed), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_RedirectAllowed));
}

void QNetworkReply_DisconnectRedirectAllowed(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::redirectAllowed), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_RedirectAllowed));
}

void QNetworkReply_RedirectAllowed(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->redirectAllowed();
}

void QNetworkReply_ConnectRedirected(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(const QUrl &)>(&QNetworkReply::redirected), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(const QUrl &)>(&MyQNetworkReply::Signal_Redirected));
}

void QNetworkReply_DisconnectRedirected(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(const QUrl &)>(&QNetworkReply::redirected), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(const QUrl &)>(&MyQNetworkReply::Signal_Redirected));
}

void QNetworkReply_Redirected(void* ptr, void* url)
{
	static_cast<QNetworkReply*>(ptr)->redirected(*static_cast<QUrl*>(url));
}

void QNetworkReply_SetAttribute(void* ptr, long long code, void* value)
{
	static_cast<QNetworkReply*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkReply_SetError(void* ptr, long long errorCode, struct QtNetwork_PackedString errorString)
{
	static_cast<QNetworkReply*>(ptr)->setError(static_cast<QNetworkReply::NetworkError>(errorCode), QString::fromUtf8(errorString.data, errorString.len));
}

void QNetworkReply_SetFinished(void* ptr, char finished)
{
	static_cast<QNetworkReply*>(ptr)->setFinished(finished != 0);
}

void QNetworkReply_SetHeader(void* ptr, long long header, void* value)
{
	static_cast<QNetworkReply*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkReply_SetOperation(void* ptr, long long operation)
{
	static_cast<QNetworkReply*>(ptr)->setOperation(static_cast<QNetworkAccessManager::Operation>(operation));
}

void QNetworkReply_SetRawHeader(void* ptr, void* headerName, void* value)
{
	static_cast<QNetworkReply*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(value));
}

void QNetworkReply_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QNetworkReply*>(ptr)->setReadBufferSize(size);
}

void QNetworkReply_SetReadBufferSizeDefault(void* ptr, long long size)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::setReadBufferSize(size);
}

void QNetworkReply_SetRequest(void* ptr, void* request)
{
	static_cast<QNetworkReply*>(ptr)->setRequest(*static_cast<QNetworkRequest*>(request));
}

void QNetworkReply_SetSslConfiguration(void* ptr, void* config)
{
	static_cast<QNetworkReply*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkReply_SetSslConfigurationImplementation(void* ptr, void* configuration)
{
	static_cast<QNetworkReply*>(ptr)->setSslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_SetSslConfigurationImplementationDefault(void* ptr, void* configuration)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::setSslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_SetUrl(void* ptr, void* url)
{
	static_cast<QNetworkReply*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkReply_ConnectSslErrors(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(const QList<QSslError> &)>(&QNetworkReply::sslErrors), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(const QList<QSslError> &)>(&MyQNetworkReply::Signal_SslErrors));
}

void QNetworkReply_DisconnectSslErrors(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(const QList<QSslError> &)>(&QNetworkReply::sslErrors), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(const QList<QSslError> &)>(&MyQNetworkReply::Signal_SslErrors));
}

void QNetworkReply_SslErrors(void* ptr, void* errors)
{
	static_cast<QNetworkReply*>(ptr)->sslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QNetworkReply_ConnectUploadProgress(void* ptr)
{
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::uploadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_UploadProgress));
}

void QNetworkReply_DisconnectUploadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::uploadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_UploadProgress));
}

void QNetworkReply_UploadProgress(void* ptr, long long bytesSent, long long bytesTotal)
{
	static_cast<QNetworkReply*>(ptr)->uploadProgress(bytesSent, bytesTotal);
}

void QNetworkReply_DestroyQNetworkReply(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->~QNetworkReply();
}

void QNetworkReply_DestroyQNetworkReplyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNetworkReply_RawHeader(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QNetworkReply*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

struct QtNetwork_PackedList QNetworkReply_RawHeaderList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QNetworkReply*>(ptr)->rawHeaderList()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QNetworkReply_Manager(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->manager();
}

long long QNetworkReply_Operation(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->operation();
}

long long QNetworkReply_Error(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->error();
}

void* QNetworkReply_Request(void* ptr)
{
	return new QNetworkRequest(static_cast<QNetworkReply*>(ptr)->request());
}

void* QNetworkReply_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QNetworkReply*>(ptr)->sslConfiguration());
}

void* QNetworkReply_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkReply*>(ptr)->url());
}

void* QNetworkReply_Attribute(void* ptr, long long code)
{
	return new QVariant(static_cast<QNetworkReply*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code)));
}

void* QNetworkReply_Header(void* ptr, long long header)
{
	return new QVariant(static_cast<QNetworkReply*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

char QNetworkReply_HasRawHeader(void* ptr, void* headerName)
{
	return static_cast<QNetworkReply*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

char QNetworkReply_IsFinished(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isFinished();
}

char QNetworkReply_IsRunning(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isRunning();
}

void* QNetworkReply_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::metaObject());
}

long long QNetworkReply_ReadBufferSize(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->readBufferSize();
}

void QNetworkReply_SslConfigurationImplementation(void* ptr, void* configuration)
{
	static_cast<QNetworkReply*>(ptr)->sslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_SslConfigurationImplementationDefault(void* ptr, void* configuration)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::sslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void* QNetworkReply___ignoreSslErrors_errors_atList2(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkReply___ignoreSslErrors_errors_setList2(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QNetworkReply___ignoreSslErrors_errors_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QNetworkReply___ignoreSslErrorsImplementation_errors_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkReply___ignoreSslErrorsImplementation_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QNetworkReply___ignoreSslErrorsImplementation_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QNetworkReply___sslErrors_errors_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkReply___sslErrors_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QNetworkReply___sslErrors_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QNetworkReply___rawHeaderList_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkReply___rawHeaderList_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkReply___rawHeaderList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNetworkReply___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkReply___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkReply___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNetworkReply___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkReply___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkReply___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkReply___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkReply___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkReply___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkReply___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkReply___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkReply___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkReply___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkReply___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkReply___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNetworkReply_OpenDefault(void* ptr, long long mode)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QNetworkReply_ResetDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::reset();
}

char QNetworkReply_SeekDefault(void* ptr, long long pos)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::seek(pos);
}

char QNetworkReply_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::waitForBytesWritten(msecs);
}

char QNetworkReply_WaitForReadyReadDefault(void* ptr, int msecs)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::waitForReadyRead(msecs);
}

long long QNetworkReply_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QNetworkReply*>(ptr)->readData(data, maxSize);
}

long long QNetworkReply_ReadDataDefault(void* ptr, char* data, long long maxSize)
{
	Q_UNUSED(ptr);
	Q_UNUSED(data);
	Q_UNUSED(maxSize);
	
}

long long QNetworkReply_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::readLineData(data, maxSize);
}

long long QNetworkReply_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QNetworkReply*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QNetworkReply_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::writeData(const_cast<const char*>(data), maxSize);
}

char QNetworkReply_AtEndDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::atEnd();
}

char QNetworkReply_CanReadLineDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::canReadLine();
}

char QNetworkReply_IsSequentialDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::isSequential();
}

long long QNetworkReply_BytesAvailableDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::bytesAvailable();
}

long long QNetworkReply_BytesToWriteDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::bytesToWrite();
}

long long QNetworkReply_PosDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::pos();
}

long long QNetworkReply_SizeDefault(void* ptr)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::size();
}

char QNetworkReply_EventDefault(void* ptr, void* e)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::event(static_cast<QEvent*>(e));
}

char QNetworkReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNetworkReply*>(ptr)->QNetworkReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNetworkReply_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkReply_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkReply_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::customEvent(static_cast<QEvent*>(event));
}

void QNetworkReply_DeleteLaterDefault(void* ptr)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::deleteLater();
}

void QNetworkReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkReply_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkReply*>(ptr)->QNetworkReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QNetworkRequest_NewQNetworkRequest2(void* other)
{
	return new QNetworkRequest(*static_cast<QNetworkRequest*>(other));
}

void* QNetworkRequest_NewQNetworkRequest(void* url)
{
	return new QNetworkRequest(*static_cast<QUrl*>(url));
}

void QNetworkRequest_SetAttribute(void* ptr, long long code, void* value)
{
	static_cast<QNetworkRequest*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetHeader(void* ptr, long long header, void* value)
{
	static_cast<QNetworkRequest*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetMaximumRedirectsAllowed(void* ptr, int maxRedirectsAllowed)
{
	static_cast<QNetworkRequest*>(ptr)->setMaximumRedirectsAllowed(maxRedirectsAllowed);
}

void QNetworkRequest_SetOriginatingObject(void* ptr, void* object)
{
	static_cast<QNetworkRequest*>(ptr)->setOriginatingObject(static_cast<QObject*>(object));
}

void QNetworkRequest_SetPriority(void* ptr, long long priority)
{
	static_cast<QNetworkRequest*>(ptr)->setPriority(static_cast<QNetworkRequest::Priority>(priority));
}

void QNetworkRequest_SetRawHeader(void* ptr, void* headerName, void* headerValue)
{
	static_cast<QNetworkRequest*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkRequest_SetSslConfiguration(void* ptr, void* config)
{
	static_cast<QNetworkRequest*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkRequest_SetUrl(void* ptr, void* url)
{
	static_cast<QNetworkRequest*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkRequest_Swap(void* ptr, void* other)
{
	static_cast<QNetworkRequest*>(ptr)->swap(*static_cast<QNetworkRequest*>(other));
}

void QNetworkRequest_DestroyQNetworkRequest(void* ptr)
{
	static_cast<QNetworkRequest*>(ptr)->~QNetworkRequest();
}

void* QNetworkRequest_RawHeader(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QNetworkRequest*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

struct QtNetwork_PackedList QNetworkRequest_RawHeaderList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QNetworkRequest*>(ptr)->rawHeaderList()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkRequest_Priority(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->priority();
}

void* QNetworkRequest_OriginatingObject(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->originatingObject();
}

void* QNetworkRequest_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QNetworkRequest*>(ptr)->sslConfiguration());
}

void* QNetworkRequest_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkRequest*>(ptr)->url());
}

void* QNetworkRequest_Attribute(void* ptr, long long code, void* defaultValue)
{
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

void* QNetworkRequest_Header(void* ptr, long long header)
{
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

char QNetworkRequest_HasRawHeader(void* ptr, void* headerName)
{
	return static_cast<QNetworkRequest*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

int QNetworkRequest_MaximumRedirectsAllowed(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->maximumRedirectsAllowed();
}

void* QNetworkRequest___rawHeaderList_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkRequest___rawHeaderList_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkRequest___rawHeaderList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

class MyQNetworkSession: public QNetworkSession
{
public:
	MyQNetworkSession(const QNetworkConfiguration &connectionConfig, QObject *parent = Q_NULLPTR) : QNetworkSession(connectionConfig, parent) {QNetworkSession_QNetworkSession_QRegisterMetaType();};
	void accept() { callbackQNetworkSession_Accept(this); };
	void close() { callbackQNetworkSession_Close(this); };
	void Signal_Closed() { callbackQNetworkSession_Closed(this); };
	void Signal_Error2(QNetworkSession::SessionError error) { callbackQNetworkSession_Error2(this, error); };
	void ignore() { callbackQNetworkSession_Ignore(this); };
	void migrate() { callbackQNetworkSession_Migrate(this); };
	void Signal_NewConfigurationActivated() { callbackQNetworkSession_NewConfigurationActivated(this); };
	void open() { callbackQNetworkSession_Open(this); };
	void Signal_Opened() { callbackQNetworkSession_Opened(this); };
	void Signal_PreferredConfigurationChanged(const QNetworkConfiguration & config, bool isSeamless) { callbackQNetworkSession_PreferredConfigurationChanged(this, const_cast<QNetworkConfiguration*>(&config), isSeamless); };
	void reject() { callbackQNetworkSession_Reject(this); };
	void Signal_StateChanged(QNetworkSession::State state) { callbackQNetworkSession_StateChanged(this, state); };
	void stop() { callbackQNetworkSession_Stop(this); };
	void Signal_UsagePoliciesChanged(QNetworkSession::UsagePolicies usagePolicies) { callbackQNetworkSession_UsagePoliciesChanged(this, usagePolicies); };
	 ~MyQNetworkSession() { callbackQNetworkSession_DestroyQNetworkSession(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkSession_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQNetworkSession_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkSession_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNetworkSession_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkSession_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkSession_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkSession_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQNetworkSession_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkSession_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQNetworkSession_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkSession_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNetworkSession*)

int QNetworkSession_QNetworkSession_QRegisterMetaType(){qRegisterMetaType<QNetworkSession*>(); return qRegisterMetaType<MyQNetworkSession*>();}

void* QNetworkSession_NewQNetworkSession(void* connectionConfig, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QWindow*>(parent));
	} else {
		return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QNetworkSession_QNetworkSession_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t4a26a4 = QNetworkSession::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t4a26a4.prepend("WHITESPACE").constData()+10), t4a26a4.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkSession_QNetworkSession_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t3b0f4c = QNetworkSession::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t3b0f4c.prepend("WHITESPACE").constData()+10), t3b0f4c.size()-10 }; });
}

char QNetworkSession_WaitForOpened(void* ptr, int msecs)
{
	return static_cast<QNetworkSession*>(ptr)->waitForOpened(msecs);
}

void QNetworkSession_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "accept");
}

void QNetworkSession_AcceptDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::accept();
}

void QNetworkSession_Close(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "close");
}

void QNetworkSession_CloseDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::close();
}

void QNetworkSession_ConnectClosed(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));
}

void QNetworkSession_DisconnectClosed(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));
}

void QNetworkSession_Closed(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->closed();
}

void QNetworkSession_ConnectError2(void* ptr)
{
	qRegisterMetaType<QNetworkSession::SessionError>();
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::SessionError)>(&QNetworkSession::error), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::SessionError)>(&MyQNetworkSession::Signal_Error2));
}

void QNetworkSession_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::SessionError)>(&QNetworkSession::error), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::SessionError)>(&MyQNetworkSession::Signal_Error2));
}

void QNetworkSession_Error2(void* ptr, long long error)
{
	static_cast<QNetworkSession*>(ptr)->error(static_cast<QNetworkSession::SessionError>(error));
}

void QNetworkSession_Ignore(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "ignore");
}

void QNetworkSession_IgnoreDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::ignore();
}

void QNetworkSession_Migrate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "migrate");
}

void QNetworkSession_MigrateDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::migrate();
}

void QNetworkSession_ConnectNewConfigurationActivated(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));
}

void QNetworkSession_DisconnectNewConfigurationActivated(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));
}

void QNetworkSession_NewConfigurationActivated(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->newConfigurationActivated();
}

void QNetworkSession_Open(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "open");
}

void QNetworkSession_OpenDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::open();
}

void QNetworkSession_ConnectOpened(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));
}

void QNetworkSession_DisconnectOpened(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));
}

void QNetworkSession_Opened(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->opened();
}

void QNetworkSession_ConnectPreferredConfigurationChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(const QNetworkConfiguration &, bool)>(&QNetworkSession::preferredConfigurationChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(const QNetworkConfiguration &, bool)>(&MyQNetworkSession::Signal_PreferredConfigurationChanged));
}

void QNetworkSession_DisconnectPreferredConfigurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(const QNetworkConfiguration &, bool)>(&QNetworkSession::preferredConfigurationChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(const QNetworkConfiguration &, bool)>(&MyQNetworkSession::Signal_PreferredConfigurationChanged));
}

void QNetworkSession_PreferredConfigurationChanged(void* ptr, void* config, char isSeamless)
{
	static_cast<QNetworkSession*>(ptr)->preferredConfigurationChanged(*static_cast<QNetworkConfiguration*>(config), isSeamless != 0);
}

void QNetworkSession_Reject(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "reject");
}

void QNetworkSession_RejectDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::reject();
}

void QNetworkSession_SetSessionProperty(void* ptr, struct QtNetwork_PackedString key, void* value)
{
	static_cast<QNetworkSession*>(ptr)->setSessionProperty(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(value));
}

void QNetworkSession_ConnectStateChanged(void* ptr)
{
	qRegisterMetaType<QNetworkSession::State>();
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));
}

void QNetworkSession_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));
}

void QNetworkSession_StateChanged(void* ptr, long long state)
{
	static_cast<QNetworkSession*>(ptr)->stateChanged(static_cast<QNetworkSession::State>(state));
}

void QNetworkSession_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "stop");
}

void QNetworkSession_StopDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::stop();
}

void QNetworkSession_ConnectUsagePoliciesChanged(void* ptr)
{
	qRegisterMetaType<QNetworkSession::UsagePolicies>();
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));
}

void QNetworkSession_DisconnectUsagePoliciesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));
}

void QNetworkSession_UsagePoliciesChanged(void* ptr, long long usagePolicies)
{
	static_cast<QNetworkSession*>(ptr)->usagePoliciesChanged(static_cast<QNetworkSession::UsagePolicy>(usagePolicies));
}

void QNetworkSession_DestroyQNetworkSession(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->~QNetworkSession();
}

void QNetworkSession_DestroyQNetworkSessionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNetworkSession_Configuration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkSession*>(ptr)->configuration());
}

void* QNetworkSession_Interface(void* ptr)
{
	return new QNetworkInterface(static_cast<QNetworkSession*>(ptr)->interface());
}

long long QNetworkSession_Error(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->error();
}

long long QNetworkSession_State(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->state();
}

long long QNetworkSession_UsagePolicies(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->usagePolicies();
}

struct QtNetwork_PackedString QNetworkSession_ErrorString(void* ptr)
{
	return ({ QByteArray t57e370 = static_cast<QNetworkSession*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t57e370.prepend("WHITESPACE").constData()+10), t57e370.size()-10 }; });
}

void* QNetworkSession_SessionProperty(void* ptr, struct QtNetwork_PackedString key)
{
	return new QVariant(static_cast<QNetworkSession*>(ptr)->sessionProperty(QString::fromUtf8(key.data, key.len)));
}

char QNetworkSession_IsOpen(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->isOpen();
}

void* QNetworkSession_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QNetworkSession*>(ptr)->QNetworkSession::metaObject());
}

unsigned long long QNetworkSession_ActiveTime(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->activeTime();
}

unsigned long long QNetworkSession_BytesReceived(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->bytesReceived();
}

unsigned long long QNetworkSession_BytesWritten(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->bytesWritten();
}

void* QNetworkSession___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QNetworkSession___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNetworkSession___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QNetworkSession___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkSession___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkSession___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkSession___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkSession___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkSession___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkSession___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkSession___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkSession___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QNetworkSession___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QNetworkSession___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNetworkSession___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QNetworkSession_EventDefault(void* ptr, void* e)
{
		return static_cast<QNetworkSession*>(ptr)->QNetworkSession::event(static_cast<QEvent*>(e));
}

char QNetworkSession_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QNetworkSession*>(ptr)->QNetworkSession::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QNetworkSession_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkSession_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkSession_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::customEvent(static_cast<QEvent*>(event));
}

void QNetworkSession_DeleteLaterDefault(void* ptr)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::deleteLater();
}

void QNetworkSession_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkSession_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QNetworkSession*>(ptr)->QNetworkSession::timerEvent(static_cast<QTimerEvent*>(event));
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_FromData(void* data, long long format)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslCertificate::fromData(*static_cast<QByteArray*>(data), static_cast<QSsl::EncodingFormat>(format))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_FromDevice(void* device, long long format)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslCertificate::fromDevice(static_cast<QIODevice*>(device), static_cast<QSsl::EncodingFormat>(format))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_FromPath(struct QtNetwork_PackedString path, long long format, long long syntax)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslCertificate::fromPath(QString::fromUtf8(path.data, path.len), static_cast<QSsl::EncodingFormat>(format), static_cast<QRegExp::PatternSyntax>(syntax))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_Verify(void* certificateChain, struct QtNetwork_PackedString hostName)
{
	return ({ QList<QSslError>* tmpValue = new QList<QSslError>(QSslCertificate::verify(({ QList<QSslCertificate>* tmpP = static_cast<QList<QSslCertificate>*>(certificateChain); QList<QSslCertificate> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }), QString::fromUtf8(hostName.data, hostName.len))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslCertificate_NewQSslCertificate(void* device, long long format)
{
	return new QSslCertificate(static_cast<QIODevice*>(device), static_cast<QSsl::EncodingFormat>(format));
}

void* QSslCertificate_NewQSslCertificate2(void* data, long long format)
{
	return new QSslCertificate(*static_cast<QByteArray*>(data), static_cast<QSsl::EncodingFormat>(format));
}

void* QSslCertificate_NewQSslCertificate3(void* other)
{
	return new QSslCertificate(*static_cast<QSslCertificate*>(other));
}

char QSslCertificate_QSslCertificate_ImportPkcs12(void* device, void* key, void* certificate, void* caCertificates, void* passPhrase)
{
	return QSslCertificate::importPkcs12(static_cast<QIODevice*>(device), static_cast<QSslKey*>(key), static_cast<QSslCertificate*>(certificate), static_cast<QList<QSslCertificate>*>(caCertificates), *static_cast<QByteArray*>(passPhrase));
}

void QSslCertificate_Clear(void* ptr)
{
	static_cast<QSslCertificate*>(ptr)->clear();
}

void QSslCertificate_Swap(void* ptr, void* other)
{
	static_cast<QSslCertificate*>(ptr)->swap(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_DestroyQSslCertificate(void* ptr)
{
	static_cast<QSslCertificate*>(ptr)->~QSslCertificate();
}

void* QSslCertificate_Digest(void* ptr, long long algorithm)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->digest(static_cast<QCryptographicHash::Algorithm>(algorithm)));
}

void* QSslCertificate_SerialNumber(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->serialNumber());
}

void* QSslCertificate_ToDer(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toDer());
}

void* QSslCertificate_ToPem(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toPem());
}

void* QSslCertificate_Version(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->version());
}

void* QSslCertificate_EffectiveDate(void* ptr)
{
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->effectiveDate());
}

void* QSslCertificate_ExpiryDate(void* ptr)
{
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->expiryDate());
}

struct QtNetwork_PackedList QSslCertificate_IssuerInfoAttributes(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QSslCertificate*>(ptr)->issuerInfoAttributes()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_SubjectInfoAttributes(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QSslCertificate*>(ptr)->subjectInfoAttributes()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_Extensions(void* ptr)
{
	return ({ QList<QSslCertificateExtension>* tmpValue = new QList<QSslCertificateExtension>(static_cast<QSslCertificate*>(ptr)->extensions()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_SubjectAlternativeNames(void* ptr)
{
	return ({ QMultiMap<QSsl::AlternativeNameEntryType, QString>* tmpValue = new QMultiMap<QSsl::AlternativeNameEntryType, QString>(static_cast<QSslCertificate*>(ptr)->subjectAlternativeNames()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslCertificate_PublicKey(void* ptr)
{
	return new QSslKey(static_cast<QSslCertificate*>(ptr)->publicKey());
}

struct QtNetwork_PackedString QSslCertificate_IssuerDisplayName(void* ptr)
{
	return ({ QByteArray tad77ca = static_cast<QSslCertificate*>(ptr)->issuerDisplayName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tad77ca.prepend("WHITESPACE").constData()+10), tad77ca.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_SubjectDisplayName(void* ptr)
{
	return ({ QByteArray t65a96d = static_cast<QSslCertificate*>(ptr)->subjectDisplayName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t65a96d.prepend("WHITESPACE").constData()+10), t65a96d.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_ToText(void* ptr)
{
	return ({ QByteArray t52ef8a = static_cast<QSslCertificate*>(ptr)->toText().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t52ef8a.prepend("WHITESPACE").constData()+10), t52ef8a.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_IssuerInfo(void* ptr, long long subject)
{
	return ({ QByteArray t768c47 = static_cast<QSslCertificate*>(ptr)->issuerInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(t768c47.prepend("WHITESPACE").constData()+10), t768c47.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_IssuerInfo2(void* ptr, void* attribute)
{
	return ({ QByteArray tc820f1 = static_cast<QSslCertificate*>(ptr)->issuerInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc820f1.prepend("WHITESPACE").constData()+10), tc820f1.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_SubjectInfo(void* ptr, long long subject)
{
	return ({ QByteArray tee2197 = static_cast<QSslCertificate*>(ptr)->subjectInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tee2197.prepend("WHITESPACE").constData()+10), tee2197.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_SubjectInfo2(void* ptr, void* attribute)
{
	return ({ QByteArray tc13a73 = static_cast<QSslCertificate*>(ptr)->subjectInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc13a73.prepend("WHITESPACE").constData()+10), tc13a73.size()-10 }; });
}

char QSslCertificate_IsBlacklisted(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isBlacklisted();
}

char QSslCertificate_IsNull(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isNull();
}

char QSslCertificate_IsSelfSigned(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isSelfSigned();
}

void* QSslCertificate___fromData_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___fromData_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslCertificate___fromData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslCertificate___fromDevice_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___fromDevice_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslCertificate___fromDevice_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslCertificate___fromPath_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___fromPath_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslCertificate___fromPath_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslCertificate___verify_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___verify_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QSslCertificate___verify_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QSslCertificate___verify_certificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___verify_certificateChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslCertificate___verify_certificateChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslCertificate___importPkcs12_caCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___importPkcs12_caCertificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslCertificate___importPkcs12_caCertificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslCertificate___issuerInfoAttributes_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___issuerInfoAttributes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSslCertificate___issuerInfoAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSslCertificate___subjectInfoAttributes_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___subjectInfoAttributes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSslCertificate___subjectInfoAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSslCertificate___extensions_atList(void* ptr, int i)
{
	return new QSslCertificateExtension(({QSslCertificateExtension tmp = static_cast<QList<QSslCertificateExtension>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificateExtension>*>(ptr)->size()-1) { static_cast<QList<QSslCertificateExtension>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslCertificate___extensions_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificateExtension>*>(ptr)->append(*static_cast<QSslCertificateExtension*>(i));
}

void* QSslCertificate___extensions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificateExtension>();
}

struct QtNetwork_PackedString QSslCertificate___subjectAlternativeNames_atList(void* ptr, long long v, int i)
{
	return ({ QByteArray tde657c = ({ QString tmp = static_cast<QMultiMap<QSsl::AlternativeNameEntryType, QString>*>(ptr)->value(static_cast<QSsl::AlternativeNameEntryType>(v)); if (i == static_cast<QMultiMap<QSsl::AlternativeNameEntryType, QString>*>(ptr)->size()-1) { static_cast<QMultiMap<QSsl::AlternativeNameEntryType, QString>*>(ptr)->~QMultiMap(); free(ptr); }; tmp; }).toUtf8(); QtNetwork_PackedString { const_cast<char*>(tde657c.prepend("WHITESPACE").constData()+10), tde657c.size()-10 }; });
}

void QSslCertificate___subjectAlternativeNames_setList(void* ptr, long long key, struct QtNetwork_PackedString i)
{
	static_cast<QMultiMap<QSsl::AlternativeNameEntryType, QString>*>(ptr)->insert(static_cast<QSsl::AlternativeNameEntryType>(key), QString::fromUtf8(i.data, i.len));
}

void* QSslCertificate___subjectAlternativeNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMultiMap<QSsl::AlternativeNameEntryType, QString>();
}

struct QtNetwork_PackedList QSslCertificate___subjectAlternativeNames_keyList(void* ptr)
{
	return ({ QList<QSsl::AlternativeNameEntryType>* tmpValue = new QList<QSsl::AlternativeNameEntryType>(static_cast<QMultiMap<QSsl::AlternativeNameEntryType, QString>*>(ptr)->keys()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QSslCertificate_____subjectAlternativeNames_keyList_atList(void* ptr, int i)
{
	return ({QSsl::AlternativeNameEntryType tmp = static_cast<QList<QSsl::AlternativeNameEntryType>*>(ptr)->at(i); if (i == static_cast<QList<QSsl::AlternativeNameEntryType>*>(ptr)->size()-1) { static_cast<QList<QSsl::AlternativeNameEntryType>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSslCertificate_____subjectAlternativeNames_keyList_setList(void* ptr, long long i)
{
	static_cast<QList<QSsl::AlternativeNameEntryType>*>(ptr)->append(static_cast<QSsl::AlternativeNameEntryType>(i));
}

void* QSslCertificate_____subjectAlternativeNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSsl::AlternativeNameEntryType>();
}

void* QSslCertificateExtension_NewQSslCertificateExtension()
{
	return new QSslCertificateExtension();
}

void* QSslCertificateExtension_NewQSslCertificateExtension2(void* other)
{
	return new QSslCertificateExtension(*static_cast<QSslCertificateExtension*>(other));
}

void QSslCertificateExtension_Swap(void* ptr, void* other)
{
	static_cast<QSslCertificateExtension*>(ptr)->swap(*static_cast<QSslCertificateExtension*>(other));
}

void QSslCertificateExtension_DestroyQSslCertificateExtension(void* ptr)
{
	static_cast<QSslCertificateExtension*>(ptr)->~QSslCertificateExtension();
}

struct QtNetwork_PackedString QSslCertificateExtension_Name(void* ptr)
{
	return ({ QByteArray t994389 = static_cast<QSslCertificateExtension*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t994389.prepend("WHITESPACE").constData()+10), t994389.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificateExtension_Oid(void* ptr)
{
	return ({ QByteArray t615506 = static_cast<QSslCertificateExtension*>(ptr)->oid().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t615506.prepend("WHITESPACE").constData()+10), t615506.size()-10 }; });
}

void* QSslCertificateExtension_Value(void* ptr)
{
	return new QVariant(static_cast<QSslCertificateExtension*>(ptr)->value());
}

char QSslCertificateExtension_IsCritical(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->isCritical();
}

char QSslCertificateExtension_IsSupported(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->isSupported();
}

void* QSslCipher_NewQSslCipher()
{
	return new QSslCipher();
}

void* QSslCipher_NewQSslCipher4(void* other)
{
	return new QSslCipher(*static_cast<QSslCipher*>(other));
}

void* QSslCipher_NewQSslCipher2(struct QtNetwork_PackedString name)
{
	return new QSslCipher(QString::fromUtf8(name.data, name.len));
}

void* QSslCipher_NewQSslCipher3(struct QtNetwork_PackedString name, long long protocol)
{
	return new QSslCipher(QString::fromUtf8(name.data, name.len), static_cast<QSsl::SslProtocol>(protocol));
}

void QSslCipher_Swap(void* ptr, void* other)
{
	static_cast<QSslCipher*>(ptr)->swap(*static_cast<QSslCipher*>(other));
}

void QSslCipher_DestroyQSslCipher(void* ptr)
{
	static_cast<QSslCipher*>(ptr)->~QSslCipher();
}

long long QSslCipher_Protocol(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->protocol();
}

struct QtNetwork_PackedString QSslCipher_AuthenticationMethod(void* ptr)
{
	return ({ QByteArray tfc1f5a = static_cast<QSslCipher*>(ptr)->authenticationMethod().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfc1f5a.prepend("WHITESPACE").constData()+10), tfc1f5a.size()-10 }; });
}

struct QtNetwork_PackedString QSslCipher_EncryptionMethod(void* ptr)
{
	return ({ QByteArray ta39f95 = static_cast<QSslCipher*>(ptr)->encryptionMethod().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta39f95.prepend("WHITESPACE").constData()+10), ta39f95.size()-10 }; });
}

struct QtNetwork_PackedString QSslCipher_KeyExchangeMethod(void* ptr)
{
	return ({ QByteArray tfbfb25 = static_cast<QSslCipher*>(ptr)->keyExchangeMethod().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfbfb25.prepend("WHITESPACE").constData()+10), tfbfb25.size()-10 }; });
}

struct QtNetwork_PackedString QSslCipher_Name(void* ptr)
{
	return ({ QByteArray t9ef3a9 = static_cast<QSslCipher*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t9ef3a9.prepend("WHITESPACE").constData()+10), t9ef3a9.size()-10 }; });
}

struct QtNetwork_PackedString QSslCipher_ProtocolString(void* ptr)
{
	return ({ QByteArray t99c307 = static_cast<QSslCipher*>(ptr)->protocolString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t99c307.prepend("WHITESPACE").constData()+10), t99c307.size()-10 }; });
}

char QSslCipher_IsNull(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->isNull();
}

int QSslCipher_SupportedBits(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->supportedBits();
}

int QSslCipher_UsedBits(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->usedBits();
}

struct QtNetwork_PackedList QSslConfiguration_QSslConfiguration_SystemCaCertificates()
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslConfiguration::systemCaCertificates()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_QSslConfiguration_SupportedCiphers()
{
	return ({ QList<QSslCipher>* tmpValue = new QList<QSslCipher>(QSslConfiguration::supportedCiphers()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslConfiguration_QSslConfiguration_DefaultConfiguration()
{
	return new QSslConfiguration(QSslConfiguration::defaultConfiguration());
}

void* QSslConfiguration_NewQSslConfiguration()
{
	return new QSslConfiguration();
}

void* QSslConfiguration_NewQSslConfiguration2(void* other)
{
	return new QSslConfiguration(*static_cast<QSslConfiguration*>(other));
}

void QSslConfiguration_SetAllowedNextProtocols(void* ptr, void* protocols)
{
	static_cast<QSslConfiguration*>(ptr)->setAllowedNextProtocols(({ QList<QByteArray>* tmpP = static_cast<QList<QByteArray>*>(protocols); QList<QByteArray> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QSslConfiguration_SetBackendConfiguration(void* ptr, void* backendConfiguration)
{
	static_cast<QSslConfiguration*>(ptr)->setBackendConfiguration(*static_cast<QMap<QByteArray, QVariant>*>(backendConfiguration));
}

void QSslConfiguration_SetBackendConfigurationOption(void* ptr, void* name, void* value)
{
	static_cast<QSslConfiguration*>(ptr)->setBackendConfigurationOption(*static_cast<QByteArray*>(name), *static_cast<QVariant*>(value));
}

void QSslConfiguration_SetCaCertificates(void* ptr, void* certificates)
{
	static_cast<QSslConfiguration*>(ptr)->setCaCertificates(*static_cast<QList<QSslCertificate>*>(certificates));
}

void QSslConfiguration_SetCiphers(void* ptr, void* ciphers)
{
	static_cast<QSslConfiguration*>(ptr)->setCiphers(*static_cast<QList<QSslCipher>*>(ciphers));
}

void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(void* configuration)
{
	QSslConfiguration::setDefaultConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void QSslConfiguration_SetDiffieHellmanParameters(void* ptr, void* dhparams)
{
	static_cast<QSslConfiguration*>(ptr)->setDiffieHellmanParameters(*static_cast<QSslDiffieHellmanParameters*>(dhparams));
}

void QSslConfiguration_SetEllipticCurves(void* ptr, void* curves)
{
	static_cast<QSslConfiguration*>(ptr)->setEllipticCurves(*static_cast<QVector<QSslEllipticCurve>*>(curves));
}

void QSslConfiguration_SetLocalCertificate(void* ptr, void* certificate)
{
	static_cast<QSslConfiguration*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslConfiguration_SetLocalCertificateChain(void* ptr, void* localChain)
{
	static_cast<QSslConfiguration*>(ptr)->setLocalCertificateChain(*static_cast<QList<QSslCertificate>*>(localChain));
}

void QSslConfiguration_SetPeerVerifyDepth(void* ptr, int depth)
{
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslConfiguration_SetPeerVerifyMode(void* ptr, long long mode)
{
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslConfiguration_SetPreSharedKeyIdentityHint(void* ptr, void* hint)
{
	static_cast<QSslConfiguration*>(ptr)->setPreSharedKeyIdentityHint(*static_cast<QByteArray*>(hint));
}

void QSslConfiguration_SetPrivateKey(void* ptr, void* key)
{
	static_cast<QSslConfiguration*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslConfiguration_SetProtocol(void* ptr, long long protocol)
{
	static_cast<QSslConfiguration*>(ptr)->setProtocol(static_cast<QSsl::SslProtocol>(protocol));
}

void QSslConfiguration_SetSessionTicket(void* ptr, void* sessionTicket)
{
	static_cast<QSslConfiguration*>(ptr)->setSessionTicket(*static_cast<QByteArray*>(sessionTicket));
}

void QSslConfiguration_SetSslOption(void* ptr, long long option, char on)
{
	static_cast<QSslConfiguration*>(ptr)->setSslOption(static_cast<QSsl::SslOption>(option), on != 0);
}

void QSslConfiguration_Swap(void* ptr, void* other)
{
	static_cast<QSslConfiguration*>(ptr)->swap(*static_cast<QSslConfiguration*>(other));
}

void QSslConfiguration_DestroyQSslConfiguration(void* ptr)
{
	static_cast<QSslConfiguration*>(ptr)->~QSslConfiguration();
}

void* QSslConfiguration_NextNegotiatedProtocol(void* ptr)
{
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->nextNegotiatedProtocol());
}

void* QSslConfiguration_PreSharedKeyIdentityHint(void* ptr)
{
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->preSharedKeyIdentityHint());
}

void* QSslConfiguration_SessionTicket(void* ptr)
{
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->sessionTicket());
}

struct QtNetwork_PackedList QSslConfiguration_AllowedNextProtocols(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QSslConfiguration*>(ptr)->allowedNextProtocols()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_CaCertificates(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslConfiguration*>(ptr)->caCertificates()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_LocalCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslConfiguration*>(ptr)->localCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_PeerCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslConfiguration*>(ptr)->peerCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_Ciphers(void* ptr)
{
	return ({ QList<QSslCipher>* tmpValue = new QList<QSslCipher>(static_cast<QSslConfiguration*>(ptr)->ciphers()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_BackendConfiguration(void* ptr)
{
	return ({ QMap<QByteArray, QVariant>* tmpValue = new QMap<QByteArray, QVariant>(static_cast<QSslConfiguration*>(ptr)->backendConfiguration()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QSslConfiguration_Protocol(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->protocol();
}

long long QSslConfiguration_SessionProtocol(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->sessionProtocol();
}

void* QSslConfiguration_LocalCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslConfiguration*>(ptr)->localCertificate());
}

void* QSslConfiguration_PeerCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslConfiguration*>(ptr)->peerCertificate());
}

void* QSslConfiguration_SessionCipher(void* ptr)
{
	return new QSslCipher(static_cast<QSslConfiguration*>(ptr)->sessionCipher());
}

long long QSslConfiguration_NextProtocolNegotiationStatus(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->nextProtocolNegotiationStatus();
}

void* QSslConfiguration_DiffieHellmanParameters(void* ptr)
{
	return new QSslDiffieHellmanParameters(static_cast<QSslConfiguration*>(ptr)->diffieHellmanParameters());
}

void* QSslConfiguration_EphemeralServerKey(void* ptr)
{
	return new QSslKey(static_cast<QSslConfiguration*>(ptr)->ephemeralServerKey());
}

void* QSslConfiguration_PrivateKey(void* ptr)
{
	return new QSslKey(static_cast<QSslConfiguration*>(ptr)->privateKey());
}

long long QSslConfiguration_PeerVerifyMode(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyMode();
}

char QSslConfiguration_IsNull(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->isNull();
}

char QSslConfiguration_TestSslOption(void* ptr, long long option)
{
	return static_cast<QSslConfiguration*>(ptr)->testSslOption(static_cast<QSsl::SslOption>(option));
}

int QSslConfiguration_PeerVerifyDepth(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyDepth();
}

int QSslConfiguration_SessionTicketLifeTimeHint(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->sessionTicketLifeTimeHint();
}

void* QSslConfiguration___systemCaCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___systemCaCertificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslConfiguration___systemCaCertificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslConfiguration___supportedCiphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___supportedCiphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslConfiguration___supportedCiphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslConfiguration___supportedEllipticCurves_atList(void* ptr, int i)
{
	return new QSslEllipticCurve(({QSslEllipticCurve tmp = static_cast<QVector<QSslEllipticCurve>*>(ptr)->at(i); if (i == static_cast<QVector<QSslEllipticCurve>*>(ptr)->size()-1) { static_cast<QVector<QSslEllipticCurve>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QSslConfiguration___supportedEllipticCurves_setList(void* ptr, void* i)
{
	static_cast<QVector<QSslEllipticCurve>*>(ptr)->append(*static_cast<QSslEllipticCurve*>(i));
}

void* QSslConfiguration___supportedEllipticCurves_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QSslEllipticCurve>();
}

void* QSslConfiguration___setAllowedNextProtocols_protocols_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___setAllowedNextProtocols_protocols_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSslConfiguration___setAllowedNextProtocols_protocols_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSslConfiguration___setBackendConfiguration_backendConfiguration_atList(void* ptr, void* v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QByteArray, QVariant>*>(ptr)->value(*static_cast<QByteArray*>(v)); if (i == static_cast<QMap<QByteArray, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QByteArray, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QSslConfiguration___setBackendConfiguration_backendConfiguration_setList(void* ptr, void* key, void* i)
{
	static_cast<QMap<QByteArray, QVariant>*>(ptr)->insert(*static_cast<QByteArray*>(key), *static_cast<QVariant*>(i));
}

void* QSslConfiguration___setBackendConfiguration_backendConfiguration_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QByteArray, QVariant>();
}

struct QtNetwork_PackedList QSslConfiguration___setBackendConfiguration_backendConfiguration_keyList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QMap<QByteArray, QVariant>*>(ptr)->keys()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslConfiguration___setCaCertificates_certificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___setCaCertificates_certificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslConfiguration___setCaCertificates_certificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslConfiguration___setCiphers_ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___setCiphers_ciphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslConfiguration___setCiphers_ciphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslConfiguration___setEllipticCurves_curves_atList(void* ptr, int i)
{
	return new QSslEllipticCurve(({QSslEllipticCurve tmp = static_cast<QVector<QSslEllipticCurve>*>(ptr)->at(i); if (i == static_cast<QVector<QSslEllipticCurve>*>(ptr)->size()-1) { static_cast<QVector<QSslEllipticCurve>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QSslConfiguration___setEllipticCurves_curves_setList(void* ptr, void* i)
{
	static_cast<QVector<QSslEllipticCurve>*>(ptr)->append(*static_cast<QSslEllipticCurve*>(i));
}

void* QSslConfiguration___setEllipticCurves_curves_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QSslEllipticCurve>();
}

void* QSslConfiguration___setLocalCertificateChain_localChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___setLocalCertificateChain_localChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslConfiguration___setLocalCertificateChain_localChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslConfiguration___allowedNextProtocols_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___allowedNextProtocols_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSslConfiguration___allowedNextProtocols_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSslConfiguration___caCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___caCertificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslConfiguration___caCertificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslConfiguration___localCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___localCertificateChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslConfiguration___localCertificateChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslConfiguration___peerCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___peerCertificateChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslConfiguration___peerCertificateChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslConfiguration___ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration___ciphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslConfiguration___ciphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslConfiguration___backendConfiguration_atList(void* ptr, void* v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QByteArray, QVariant>*>(ptr)->value(*static_cast<QByteArray*>(v)); if (i == static_cast<QMap<QByteArray, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QByteArray, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QSslConfiguration___backendConfiguration_setList(void* ptr, void* key, void* i)
{
	static_cast<QMap<QByteArray, QVariant>*>(ptr)->insert(*static_cast<QByteArray*>(key), *static_cast<QVariant*>(i));
}

void* QSslConfiguration___backendConfiguration_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QByteArray, QVariant>();
}

struct QtNetwork_PackedList QSslConfiguration___backendConfiguration_keyList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QMap<QByteArray, QVariant>*>(ptr)->keys()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslConfiguration___ellipticCurves_atList(void* ptr, int i)
{
	return new QSslEllipticCurve(({QSslEllipticCurve tmp = static_cast<QVector<QSslEllipticCurve>*>(ptr)->at(i); if (i == static_cast<QVector<QSslEllipticCurve>*>(ptr)->size()-1) { static_cast<QVector<QSslEllipticCurve>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QSslConfiguration___ellipticCurves_setList(void* ptr, void* i)
{
	static_cast<QVector<QSslEllipticCurve>*>(ptr)->append(*static_cast<QSslEllipticCurve*>(i));
}

void* QSslConfiguration___ellipticCurves_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QSslEllipticCurve>();
}

void* QSslConfiguration_____setBackendConfiguration_backendConfiguration_keyList_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration_____setBackendConfiguration_backendConfiguration_keyList_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSslConfiguration_____setBackendConfiguration_backendConfiguration_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSslConfiguration_____backendConfiguration_keyList_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslConfiguration_____backendConfiguration_keyList_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSslConfiguration_____backendConfiguration_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_DefaultParameters()
{
	return new QSslDiffieHellmanParameters(QSslDiffieHellmanParameters::defaultParameters());
}

void* QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_FromEncoded2(void* device, long long encoding)
{
	return new QSslDiffieHellmanParameters(QSslDiffieHellmanParameters::fromEncoded(static_cast<QIODevice*>(device), static_cast<QSsl::EncodingFormat>(encoding)));
}

void* QSslDiffieHellmanParameters_QSslDiffieHellmanParameters_FromEncoded(void* encoded, long long encoding)
{
	return new QSslDiffieHellmanParameters(QSslDiffieHellmanParameters::fromEncoded(*static_cast<QByteArray*>(encoded), static_cast<QSsl::EncodingFormat>(encoding)));
}

void* QSslDiffieHellmanParameters_NewQSslDiffieHellmanParameters()
{
	return new QSslDiffieHellmanParameters();
}

void* QSslDiffieHellmanParameters_NewQSslDiffieHellmanParameters3(void* other)
{
	return new QSslDiffieHellmanParameters(*static_cast<QSslDiffieHellmanParameters*>(other));
}

void* QSslDiffieHellmanParameters_NewQSslDiffieHellmanParameters2(void* other)
{
	return new QSslDiffieHellmanParameters(*static_cast<QSslDiffieHellmanParameters*>(other));
}

void QSslDiffieHellmanParameters_Swap(void* ptr, void* other)
{
	static_cast<QSslDiffieHellmanParameters*>(ptr)->swap(*static_cast<QSslDiffieHellmanParameters*>(other));
}

void QSslDiffieHellmanParameters_DestroyQSslDiffieHellmanParameters(void* ptr)
{
	static_cast<QSslDiffieHellmanParameters*>(ptr)->~QSslDiffieHellmanParameters();
}

long long QSslDiffieHellmanParameters_Error(void* ptr)
{
	return static_cast<QSslDiffieHellmanParameters*>(ptr)->error();
}

struct QtNetwork_PackedString QSslDiffieHellmanParameters_ErrorString(void* ptr)
{
	return ({ QByteArray tfc2570 = static_cast<QSslDiffieHellmanParameters*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfc2570.prepend("WHITESPACE").constData()+10), tfc2570.size()-10 }; });
}

char QSslDiffieHellmanParameters_IsEmpty(void* ptr)
{
	return static_cast<QSslDiffieHellmanParameters*>(ptr)->isEmpty();
}

char QSslDiffieHellmanParameters_IsValid(void* ptr)
{
	return static_cast<QSslDiffieHellmanParameters*>(ptr)->isValid();
}

void* QSslEllipticCurve_QSslEllipticCurve_FromLongName(struct QtNetwork_PackedString name)
{
	return new QSslEllipticCurve(QSslEllipticCurve::fromLongName(QString::fromUtf8(name.data, name.len)));
}

void* QSslEllipticCurve_QSslEllipticCurve_FromShortName(struct QtNetwork_PackedString name)
{
	return new QSslEllipticCurve(QSslEllipticCurve::fromShortName(QString::fromUtf8(name.data, name.len)));
}

void* QSslEllipticCurve_NewQSslEllipticCurve()
{
	return new QSslEllipticCurve();
}

struct QtNetwork_PackedString QSslEllipticCurve_LongName(void* ptr)
{
	return ({ QByteArray t85b564 = static_cast<QSslEllipticCurve*>(ptr)->longName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t85b564.prepend("WHITESPACE").constData()+10), t85b564.size()-10 }; });
}

struct QtNetwork_PackedString QSslEllipticCurve_ShortName(void* ptr)
{
	return ({ QByteArray tb6e6fc = static_cast<QSslEllipticCurve*>(ptr)->shortName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb6e6fc.prepend("WHITESPACE").constData()+10), tb6e6fc.size()-10 }; });
}

char QSslEllipticCurve_IsTlsNamedCurve(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->isTlsNamedCurve();
}

char QSslEllipticCurve_IsValid(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->isValid();
}

void* QSslError_NewQSslError()
{
	return new QSslError();
}

void* QSslError_NewQSslError2(long long error)
{
	return new QSslError(static_cast<QSslError::SslError>(error));
}

void* QSslError_NewQSslError3(long long error, void* certificate)
{
	return new QSslError(static_cast<QSslError::SslError>(error), *static_cast<QSslCertificate*>(certificate));
}

void* QSslError_NewQSslError4(void* other)
{
	return new QSslError(*static_cast<QSslError*>(other));
}

void QSslError_Swap(void* ptr, void* other)
{
	static_cast<QSslError*>(ptr)->swap(*static_cast<QSslError*>(other));
}

void QSslError_DestroyQSslError(void* ptr)
{
	static_cast<QSslError*>(ptr)->~QSslError();
}

void* QSslError_Certificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslError*>(ptr)->certificate());
}

long long QSslError_Error(void* ptr)
{
	return static_cast<QSslError*>(ptr)->error();
}

struct QtNetwork_PackedString QSslError_ErrorString(void* ptr)
{
	return ({ QByteArray t759e3e = static_cast<QSslError*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t759e3e.prepend("WHITESPACE").constData()+10), t759e3e.size()-10 }; });
}

void* QSslKey_NewQSslKey()
{
	return new QSslKey();
}

void* QSslKey_NewQSslKey3(void* device, long long algorithm, long long encoding, long long ty, void* passPhrase)
{
	return new QSslKey(static_cast<QIODevice*>(device), static_cast<QSsl::KeyAlgorithm>(algorithm), static_cast<QSsl::EncodingFormat>(encoding), static_cast<QSsl::KeyType>(ty), *static_cast<QByteArray*>(passPhrase));
}

void* QSslKey_NewQSslKey2(void* encoded, long long algorithm, long long encoding, long long ty, void* passPhrase)
{
	return new QSslKey(*static_cast<QByteArray*>(encoded), static_cast<QSsl::KeyAlgorithm>(algorithm), static_cast<QSsl::EncodingFormat>(encoding), static_cast<QSsl::KeyType>(ty), *static_cast<QByteArray*>(passPhrase));
}

void* QSslKey_NewQSslKey5(void* other)
{
	return new QSslKey(*static_cast<QSslKey*>(other));
}

void QSslKey_Clear(void* ptr)
{
	static_cast<QSslKey*>(ptr)->clear();
}

void QSslKey_Swap(void* ptr, void* other)
{
	static_cast<QSslKey*>(ptr)->swap(*static_cast<QSslKey*>(other));
}

void QSslKey_DestroyQSslKey(void* ptr)
{
	static_cast<QSslKey*>(ptr)->~QSslKey();
}

void* QSslKey_ToDer(void* ptr, void* passPhrase)
{
	return new QByteArray(static_cast<QSslKey*>(ptr)->toDer(*static_cast<QByteArray*>(passPhrase)));
}

void* QSslKey_ToPem(void* ptr, void* passPhrase)
{
	return new QByteArray(static_cast<QSslKey*>(ptr)->toPem(*static_cast<QByteArray*>(passPhrase)));
}

long long QSslKey_Algorithm(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->algorithm();
}

long long QSslKey_Type(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->type();
}

char QSslKey_IsNull(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->isNull();
}

int QSslKey_Length(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->length();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator()
{
	return new QSslPreSharedKeyAuthenticator();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(void* authenticator)
{
	return new QSslPreSharedKeyAuthenticator(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslPreSharedKeyAuthenticator_SetIdentity(void* ptr, void* identity)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setIdentity(*static_cast<QByteArray*>(identity));
}

void QSslPreSharedKeyAuthenticator_SetPreSharedKey(void* ptr, void* preSharedKey)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setPreSharedKey(*static_cast<QByteArray*>(preSharedKey));
}

void QSslPreSharedKeyAuthenticator_Swap(void* ptr, void* authenticator)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->swap(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(void* ptr)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->~QSslPreSharedKeyAuthenticator();
}

void* QSslPreSharedKeyAuthenticator_Identity(void* ptr)
{
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identity());
}

void* QSslPreSharedKeyAuthenticator_IdentityHint(void* ptr)
{
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identityHint());
}

void* QSslPreSharedKeyAuthenticator_PreSharedKey(void* ptr)
{
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->preSharedKey());
}

int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(void* ptr)
{
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumIdentityLength();
}

int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(void* ptr)
{
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumPreSharedKeyLength();
}

class MyQSslSocket: public QSslSocket
{
public:
	MyQSslSocket(QObject *parent = Q_NULLPTR) : QSslSocket(parent) {QSslSocket_QSslSocket_QRegisterMetaType();};
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQAbstractSocket_SocketOption(this, option)); };
	bool waitForBytesWritten(int msecs) { return callbackQAbstractSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQAbstractSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQAbstractSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQAbstractSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 readData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQAbstractSocket_ReadData(this, dataPacked, maxlen); };
	qint64 writeData(const char * data, qint64 l) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), l };return callbackQAbstractSocket_WriteData(this, dataPacked, l); };
	void close() { callbackQAbstractSocket_Close(this); };
	void Signal_Encrypted() { callbackQSslSocket_Encrypted(this); };
	void Signal_EncryptedBytesWritten(qint64 written) { callbackQSslSocket_EncryptedBytesWritten(this, written); };
	void ignoreSslErrors() { callbackQSslSocket_IgnoreSslErrors(this); };
	void Signal_ModeChanged(QSslSocket::SslMode mode) { callbackQSslSocket_ModeChanged(this, mode); };
	void Signal_PeerVerifyError(const QSslError & error) { callbackQSslSocket_PeerVerifyError(this, const_cast<QSslError*>(&error)); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQSslSocket_PreSharedKeyAuthenticationRequired(this, authenticator); };
	void resume() { callbackQAbstractSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocket_SetReadBufferSize(this, size); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	void Signal_SslErrors2(const QList<QSslError> & errors) { callbackQSslSocket_SslErrors2(this, ({ QList<QSslError>* tmpValue = const_cast<QList<QSslError>*>(&errors); QtNetwork_PackedList { tmpValue, tmpValue->size() }; })); };
	void startClientEncryption() { callbackQSslSocket_StartClientEncryption(this); };
	void startServerEncryption() { callbackQSslSocket_StartServerEncryption(this); };
	 ~MyQSslSocket() { callbackQSslSocket_DestroyQSslSocket(this); };
	bool atEnd() const { return callbackQAbstractSocket_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool canReadLine() const { return callbackQAbstractSocket_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSocket_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	qint64 bytesAvailable() const { return callbackQAbstractSocket_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQAbstractSocket_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQAbstractSocket_ReadLineData(this, dataPacked, maxlen); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQAbstractSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQAbstractSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void Signal_Connected() { callbackQAbstractSocket_Connected(this); };
	void disconnectFromHost() { callbackQAbstractSocket_DisconnectFromHost(this); };
	void Signal_Disconnected() { callbackQAbstractSocket_Disconnected(this); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocket_Error2(this, socketError); };
	void Signal_HostFound() { callbackQAbstractSocket_HostFound(this); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQAbstractSocket_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocket_StateChanged(this, socketState); };
	bool isSequential() const { return callbackQAbstractSocket_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQAbstractSocket_Open(this, mode) != 0; };
	bool reset() { return callbackQAbstractSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQAbstractSocket_Seek(this, pos) != 0; };
	void Signal_AboutToClose() { callbackQAbstractSocket_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQAbstractSocket_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQAbstractSocket_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQAbstractSocket_ChannelReadyRead(this, channel); };
	void Signal_ReadChannelFinished() { callbackQAbstractSocket_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQAbstractSocket_ReadyRead(this); };
	qint64 pos() const { return callbackQAbstractSocket_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQAbstractSocket_Size(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQAbstractSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSocket_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSocket_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSocket_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSocket_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocket_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQSslSocket*)

int QSslSocket_QSslSocket_QRegisterMetaType(){qRegisterMetaType<QSslSocket*>(); return qRegisterMetaType<MyQSslSocket*>();}

void* QSslSocket_NewQSslSocket(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSslSocket(static_cast<QWindow*>(parent));
	} else {
		return new MyQSslSocket(static_cast<QObject*>(parent));
	}
}

struct QtNetwork_PackedString QSslSocket_QSslSocket_SslLibraryBuildVersionString()
{
	return ({ QByteArray t55b90e = QSslSocket::sslLibraryBuildVersionString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t55b90e.prepend("WHITESPACE").constData()+10), t55b90e.size()-10 }; });
}

struct QtNetwork_PackedString QSslSocket_QSslSocket_SslLibraryVersionString()
{
	return ({ QByteArray tdd64b4 = QSslSocket::sslLibraryVersionString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tdd64b4.prepend("WHITESPACE").constData()+10), tdd64b4.size()-10 }; });
}

char QSslSocket_AddCaCertificates(void* ptr, struct QtNetwork_PackedString path, long long format, long long syntax)
{
	return static_cast<QSslSocket*>(ptr)->addCaCertificates(QString::fromUtf8(path.data, path.len), static_cast<QSsl::EncodingFormat>(format), static_cast<QRegExp::PatternSyntax>(syntax));
}

char QSslSocket_QSslSocket_AddDefaultCaCertificates(struct QtNetwork_PackedString path, long long encoding, long long syntax)
{
	return QSslSocket::addDefaultCaCertificates(QString::fromUtf8(path.data, path.len), static_cast<QSsl::EncodingFormat>(encoding), static_cast<QRegExp::PatternSyntax>(syntax));
}

char QSslSocket_QSslSocket_SupportsSsl()
{
	return QSslSocket::supportsSsl();
}

char QSslSocket_WaitForEncrypted(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForEncrypted(msecs);
}

long QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()
{
	return QSslSocket::sslLibraryBuildVersionNumber();
}

long QSslSocket_QSslSocket_SslLibraryVersionNumber()
{
	return QSslSocket::sslLibraryVersionNumber();
}

void QSslSocket_AddCaCertificate(void* ptr, void* certificate)
{
	static_cast<QSslSocket*>(ptr)->addCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_AddCaCertificates2(void* ptr, void* certificates)
{
	static_cast<QSslSocket*>(ptr)->addCaCertificates(*static_cast<QList<QSslCertificate>*>(certificates));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificate(void* certificate)
{
	QSslSocket::addDefaultCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificates2(void* certificates)
{
	QSslSocket::addDefaultCaCertificates(*static_cast<QList<QSslCertificate>*>(certificates));
}

void QSslSocket_ConnectToHostEncrypted(void* ptr, struct QtNetwork_PackedString hostName, unsigned short port, long long mode, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->connectToHostEncrypted(QString::fromUtf8(hostName.data, hostName.len), port, static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QSslSocket_ConnectToHostEncrypted2(void* ptr, struct QtNetwork_PackedString hostName, unsigned short port, struct QtNetwork_PackedString sslPeerName, long long mode, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->connectToHostEncrypted(QString::fromUtf8(hostName.data, hostName.len), port, QString::fromUtf8(sslPeerName.data, sslPeerName.len), static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QSslSocket_ConnectEncrypted(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));
}

void QSslSocket_DisconnectEncrypted(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));
}

void QSslSocket_Encrypted(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->encrypted();
}

void QSslSocket_ConnectEncryptedBytesWritten(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(qint64)>(&QSslSocket::encryptedBytesWritten), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(qint64)>(&MyQSslSocket::Signal_EncryptedBytesWritten));
}

void QSslSocket_DisconnectEncryptedBytesWritten(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(qint64)>(&QSslSocket::encryptedBytesWritten), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(qint64)>(&MyQSslSocket::Signal_EncryptedBytesWritten));
}

void QSslSocket_EncryptedBytesWritten(void* ptr, long long written)
{
	static_cast<QSslSocket*>(ptr)->encryptedBytesWritten(written);
}

void QSslSocket_IgnoreSslErrors(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "ignoreSslErrors");
}

void QSslSocket_IgnoreSslErrorsDefault(void* ptr)
{
		static_cast<QSslSocket*>(ptr)->QSslSocket::ignoreSslErrors();
}

void QSslSocket_IgnoreSslErrors2(void* ptr, void* errors)
{
	static_cast<QSslSocket*>(ptr)->ignoreSslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QSslSocket_ConnectModeChanged(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));
}

void QSslSocket_DisconnectModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));
}

void QSslSocket_ModeChanged(void* ptr, long long mode)
{
	static_cast<QSslSocket*>(ptr)->modeChanged(static_cast<QSslSocket::SslMode>(mode));
}

void QSslSocket_ConnectPeerVerifyError(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(const QSslError &)>(&QSslSocket::peerVerifyError), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(const QSslError &)>(&MyQSslSocket::Signal_PeerVerifyError));
}

void QSslSocket_DisconnectPeerVerifyError(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(const QSslError &)>(&QSslSocket::peerVerifyError), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(const QSslError &)>(&MyQSslSocket::Signal_PeerVerifyError));
}

void QSslSocket_PeerVerifyError(void* ptr, void* error)
{
	static_cast<QSslSocket*>(ptr)->peerVerifyError(*static_cast<QSslError*>(error));
}

void QSslSocket_ConnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));
}

void QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));
}

void QSslSocket_PreSharedKeyAuthenticationRequired(void* ptr, void* authenticator)
{
	static_cast<QSslSocket*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslSocket_SetLocalCertificate(void* ptr, void* certificate)
{
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_SetLocalCertificate2(void* ptr, struct QtNetwork_PackedString path, long long format)
{
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(QString::fromUtf8(path.data, path.len), static_cast<QSsl::EncodingFormat>(format));
}

void QSslSocket_SetLocalCertificateChain(void* ptr, void* localChain)
{
	static_cast<QSslSocket*>(ptr)->setLocalCertificateChain(*static_cast<QList<QSslCertificate>*>(localChain));
}

void QSslSocket_SetPeerVerifyDepth(void* ptr, int depth)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslSocket_SetPeerVerifyMode(void* ptr, long long mode)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslSocket_SetPeerVerifyName(void* ptr, struct QtNetwork_PackedString hostName)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyName(QString::fromUtf8(hostName.data, hostName.len));
}

void QSslSocket_SetPrivateKey(void* ptr, void* key)
{
	static_cast<QSslSocket*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslSocket_SetPrivateKey2(void* ptr, struct QtNetwork_PackedString fileName, long long algorithm, long long format, void* passPhrase)
{
	static_cast<QSslSocket*>(ptr)->setPrivateKey(QString::fromUtf8(fileName.data, fileName.len), static_cast<QSsl::KeyAlgorithm>(algorithm), static_cast<QSsl::EncodingFormat>(format), *static_cast<QByteArray*>(passPhrase));
}

void QSslSocket_SetProtocol(void* ptr, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->setProtocol(static_cast<QSsl::SslProtocol>(protocol));
}

void QSslSocket_SetSslConfiguration(void* ptr, void* configuration)
{
	static_cast<QSslSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void QSslSocket_ConnectSslErrors2(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(const QList<QSslError> &)>(&QSslSocket::sslErrors), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(const QList<QSslError> &)>(&MyQSslSocket::Signal_SslErrors2));
}

void QSslSocket_DisconnectSslErrors2(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(const QList<QSslError> &)>(&QSslSocket::sslErrors), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(const QList<QSslError> &)>(&MyQSslSocket::Signal_SslErrors2));
}

void QSslSocket_SslErrors2(void* ptr, void* errors)
{
	static_cast<QSslSocket*>(ptr)->sslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QSslSocket_StartClientEncryption(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startClientEncryption");
}

void QSslSocket_StartClientEncryptionDefault(void* ptr)
{
		static_cast<QSslSocket*>(ptr)->QSslSocket::startClientEncryption();
}

void QSslSocket_StartServerEncryption(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startServerEncryption");
}

void QSslSocket_StartServerEncryptionDefault(void* ptr)
{
		static_cast<QSslSocket*>(ptr)->QSslSocket::startServerEncryption();
}

void QSslSocket_DestroyQSslSocket(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->~QSslSocket();
}

void QSslSocket_DestroyQSslSocketDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtNetwork_PackedList QSslSocket_LocalCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslSocket*>(ptr)->localCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslSocket_PeerCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslSocket*>(ptr)->peerCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslSocket_SslErrors(void* ptr)
{
	return ({ QList<QSslError>* tmpValue = new QList<QSslError>(static_cast<QSslSocket*>(ptr)->sslErrors()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QSslSocket_Protocol(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->protocol();
}

long long QSslSocket_SessionProtocol(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->sessionProtocol();
}

void* QSslSocket_LocalCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslSocket*>(ptr)->localCertificate());
}

void* QSslSocket_PeerCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslSocket*>(ptr)->peerCertificate());
}

void* QSslSocket_SessionCipher(void* ptr)
{
	return new QSslCipher(static_cast<QSslSocket*>(ptr)->sessionCipher());
}

void* QSslSocket_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QSslSocket*>(ptr)->sslConfiguration());
}

void* QSslSocket_PrivateKey(void* ptr)
{
	return new QSslKey(static_cast<QSslSocket*>(ptr)->privateKey());
}

long long QSslSocket_PeerVerifyMode(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->peerVerifyMode();
}

long long QSslSocket_Mode(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->mode();
}

struct QtNetwork_PackedString QSslSocket_PeerVerifyName(void* ptr)
{
	return ({ QByteArray tefa5dd = static_cast<QSslSocket*>(ptr)->peerVerifyName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tefa5dd.prepend("WHITESPACE").constData()+10), tefa5dd.size()-10 }; });
}

char QSslSocket_IsEncrypted(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->isEncrypted();
}

int QSslSocket_PeerVerifyDepth(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->peerVerifyDepth();
}

long long QSslSocket_EncryptedBytesAvailable(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->encryptedBytesAvailable();
}

long long QSslSocket_EncryptedBytesToWrite(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->encryptedBytesToWrite();
}

void* QSslSocket___defaultCaCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___defaultCaCertificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___defaultCaCertificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___systemCaCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___systemCaCertificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___systemCaCertificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___defaultCiphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___defaultCiphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslSocket___defaultCiphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslSocket___supportedCiphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___supportedCiphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslSocket___supportedCiphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslSocket___addCaCertificates_certificates_atList2(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___addCaCertificates_certificates_setList2(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___addCaCertificates_certificates_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___addDefaultCaCertificates_certificates_atList2(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___addDefaultCaCertificates_certificates_setList2(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___addDefaultCaCertificates_certificates_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___ignoreSslErrors_errors_atList2(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___ignoreSslErrors_errors_setList2(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QSslSocket___ignoreSslErrors_errors_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QSslSocket___setCaCertificates_certificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___setCaCertificates_certificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___setCaCertificates_certificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___setCiphers_ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___setCiphers_ciphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslSocket___setCiphers_ciphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslSocket___setDefaultCaCertificates_certificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___setDefaultCaCertificates_certificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___setDefaultCaCertificates_certificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___setDefaultCiphers_ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___setDefaultCiphers_ciphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslSocket___setDefaultCiphers_ciphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslSocket___setLocalCertificateChain_localChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___setLocalCertificateChain_localChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___setLocalCertificateChain_localChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___sslErrors_errors_atList2(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___sslErrors_errors_setList2(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QSslSocket___sslErrors_errors_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QSslSocket___caCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___caCertificates_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___caCertificates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___localCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___localCertificateChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___localCertificateChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___peerCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(({QSslCertificate tmp = static_cast<QList<QSslCertificate>*>(ptr)->at(i); if (i == static_cast<QList<QSslCertificate>*>(ptr)->size()-1) { static_cast<QList<QSslCertificate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___peerCertificateChain_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCertificate>*>(ptr)->append(*static_cast<QSslCertificate*>(i));
}

void* QSslSocket___peerCertificateChain_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCertificate>();
}

void* QSslSocket___ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(({QSslCipher tmp = static_cast<QList<QSslCipher>*>(ptr)->at(i); if (i == static_cast<QList<QSslCipher>*>(ptr)->size()-1) { static_cast<QList<QSslCipher>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___ciphers_setList(void* ptr, void* i)
{
	static_cast<QList<QSslCipher>*>(ptr)->append(*static_cast<QSslCipher*>(i));
}

void* QSslSocket___ciphers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslCipher>();
}

void* QSslSocket___sslErrors_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSslSocket___sslErrors_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QSslSocket___sslErrors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

class MyQTcpServer: public QTcpServer
{
public:
	MyQTcpServer(QObject *parent = Q_NULLPTR) : QTcpServer(parent) {QTcpServer_QTcpServer_QRegisterMetaType();};
	QTcpSocket * nextPendingConnection() { return static_cast<QTcpSocket*>(callbackQTcpServer_NextPendingConnection(this)); };
	void Signal_AcceptError(QAbstractSocket::SocketError socketError) { callbackQTcpServer_AcceptError(this, socketError); };
	void Signal_NewConnection() { callbackQTcpServer_NewConnection(this); };
	 ~MyQTcpServer() { callbackQTcpServer_DestroyQTcpServer(this); };
	bool hasPendingConnections() const { return callbackQTcpServer_HasPendingConnections(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTcpServer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQTcpServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTcpServer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQTcpServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTcpServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTcpServer_CustomEvent(this, event); };
	void deleteLater() { callbackQTcpServer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTcpServer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTcpServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQTcpServer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTcpServer_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQTcpServer*)

int QTcpServer_QTcpServer_QRegisterMetaType(){qRegisterMetaType<QTcpServer*>(); return qRegisterMetaType<MyQTcpServer*>();}

struct QtNetwork_PackedString QTcpServer_QTcpServer_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tfa2729 = QTcpServer::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfa2729.prepend("WHITESPACE").constData()+10), tfa2729.size()-10 }; });
}

struct QtNetwork_PackedString QTcpServer_QTcpServer_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t8e2170 = QTcpServer::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtNetwork_PackedString { const_cast<char*>(t8e2170.prepend("WHITESPACE").constData()+10), t8e2170.size()-10 }; });
}

void* QTcpServer_NewQTcpServer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTcpServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQTcpServer(static_cast<QObject*>(parent));
	}
}

void* QTcpServer_NextPendingConnection(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->nextPendingConnection();
}

void* QTcpServer_NextPendingConnectionDefault(void* ptr)
{
		return static_cast<QTcpServer*>(ptr)->QTcpServer::nextPendingConnection();
}

char QTcpServer_Listen(void* ptr, void* address, unsigned short port)
{
	return static_cast<QTcpServer*>(ptr)->listen(*static_cast<QHostAddress*>(address), port);
}

char QTcpServer_WaitForNewConnection(void* ptr, int msec, char* timedOut)
{
	return static_cast<QTcpServer*>(ptr)->waitForNewConnection(msec, reinterpret_cast<bool*>(timedOut));
}

void QTcpServer_ConnectAcceptError(void* ptr)
{
	qRegisterMetaType<QAbstractSocket::SocketError>();
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));
}

void QTcpServer_DisconnectAcceptError(void* ptr)
{
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));
}

void QTcpServer_AcceptError(void* ptr, long long socketError)
{
	static_cast<QTcpServer*>(ptr)->acceptError(static_cast<QAbstractSocket::SocketError>(socketError));
}

void QTcpServer_AddPendingConnection(void* ptr, void* socket)
{
	static_cast<QTcpServer*>(ptr)->addPendingConnection(static_cast<QTcpSocket*>(socket));
}

void QTcpServer_Close(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->close();
}

void QTcpServer_ConnectNewConnection(void* ptr)
{
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));
}

void QTcpServer_DisconnectNewConnection(void* ptr)
{
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));
}

void QTcpServer_NewConnection(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->newConnection();
}

void QTcpServer_PauseAccepting(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->pauseAccepting();
}

void QTcpServer_ResumeAccepting(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->resumeAccepting();
}

void QTcpServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QTcpServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QTcpServer_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QTcpServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QTcpServer_DestroyQTcpServer(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->~QTcpServer();
}

void QTcpServer_DestroyQTcpServerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QTcpServer_ServerError(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->serverError();
}

void* QTcpServer_ServerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QTcpServer*>(ptr)->serverAddress());
}

void* QTcpServer_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QTcpServer*>(ptr)->proxy());
}

struct QtNetwork_PackedString QTcpServer_ErrorString(void* ptr)
{
	return ({ QByteArray taa8c34 = static_cast<QTcpServer*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(taa8c34.prepend("WHITESPACE").constData()+10), taa8c34.size()-10 }; });
}

char QTcpServer_HasPendingConnections(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->hasPendingConnections();
}

char QTcpServer_HasPendingConnectionsDefault(void* ptr)
{
		return static_cast<QTcpServer*>(ptr)->QTcpServer::hasPendingConnections();
}

char QTcpServer_IsListening(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->isListening();
}

void* QTcpServer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QTcpServer*>(ptr)->QTcpServer::metaObject());
}

int QTcpServer_MaxPendingConnections(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->maxPendingConnections();
}

unsigned short QTcpServer_ServerPort(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->serverPort();
}

void* QTcpServer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTcpServer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTcpServer___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTcpServer___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTcpServer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTcpServer___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTcpServer___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTcpServer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTcpServer___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTcpServer___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTcpServer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTcpServer___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTcpServer___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTcpServer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTcpServer___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QTcpServer_EventDefault(void* ptr, void* e)
{
		return static_cast<QTcpServer*>(ptr)->QTcpServer::event(static_cast<QEvent*>(e));
}

char QTcpServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QTcpServer*>(ptr)->QTcpServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QTcpServer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QTcpServer*>(ptr)->QTcpServer::childEvent(static_cast<QChildEvent*>(event));
}

void QTcpServer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTcpServer*>(ptr)->QTcpServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpServer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QTcpServer*>(ptr)->QTcpServer::customEvent(static_cast<QEvent*>(event));
}

void QTcpServer_DeleteLaterDefault(void* ptr)
{
		static_cast<QTcpServer*>(ptr)->QTcpServer::deleteLater();
}

void QTcpServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTcpServer*>(ptr)->QTcpServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpServer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QTcpServer*>(ptr)->QTcpServer::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQTcpSocket: public QTcpSocket
{
public:
	MyQTcpSocket(QObject *parent = Q_NULLPTR) : QTcpSocket(parent) {QTcpSocket_QTcpSocket_QRegisterMetaType();};
	 ~MyQTcpSocket() { callbackQTcpSocket_DestroyQTcpSocket(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSocket_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQAbstractSocket_SocketOption(this, option)); };
	bool waitForBytesWritten(int msecs) { return callbackQAbstractSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQAbstractSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQAbstractSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQAbstractSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQAbstractSocket_ReadData(this, dataPacked, maxSize); };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQAbstractSocket_ReadLineData(this, dataPacked, maxlen); };
	qint64 writeData(const char * data, qint64 size) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQAbstractSocket_WriteData(this, dataPacked, size); };
	void close() { callbackQAbstractSocket_Close(this); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQAbstractSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQAbstractSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void Signal_Connected() { callbackQAbstractSocket_Connected(this); };
	void disconnectFromHost() { callbackQAbstractSocket_DisconnectFromHost(this); };
	void Signal_Disconnected() { callbackQAbstractSocket_Disconnected(this); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocket_Error2(this, socketError); };
	void Signal_HostFound() { callbackQAbstractSocket_HostFound(this); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQAbstractSocket_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	void resume() { callbackQAbstractSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocket_SetReadBufferSize(this, size); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocket_StateChanged(this, socketState); };
	bool atEnd() const { return callbackQAbstractSocket_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool canReadLine() const { return callbackQAbstractSocket_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isSequential() const { return callbackQAbstractSocket_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	qint64 bytesAvailable() const { return callbackQAbstractSocket_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQAbstractSocket_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	bool open(QIODevice::OpenMode mode) { return callbackQAbstractSocket_Open(this, mode) != 0; };
	bool reset() { return callbackQAbstractSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQAbstractSocket_Seek(this, pos) != 0; };
	void Signal_AboutToClose() { callbackQAbstractSocket_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQAbstractSocket_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQAbstractSocket_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQAbstractSocket_ChannelReadyRead(this, channel); };
	void Signal_ReadChannelFinished() { callbackQAbstractSocket_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQAbstractSocket_ReadyRead(this); };
	qint64 pos() const { return callbackQAbstractSocket_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQAbstractSocket_Size(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQAbstractSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSocket_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSocket_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSocket_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSocket_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocket_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQTcpSocket*)

int QTcpSocket_QTcpSocket_QRegisterMetaType(){qRegisterMetaType<QTcpSocket*>(); return qRegisterMetaType<MyQTcpSocket*>();}

void* QTcpSocket_NewQTcpSocket(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTcpSocket(static_cast<QWindow*>(parent));
	} else {
		return new MyQTcpSocket(static_cast<QObject*>(parent));
	}
}

void QTcpSocket_DestroyQTcpSocket(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->~QTcpSocket();
}

void QTcpSocket_DestroyQTcpSocketDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQUdpSocket: public QUdpSocket
{
public:
	MyQUdpSocket(QObject *parent = Q_NULLPTR) : QUdpSocket(parent) {QUdpSocket_QUdpSocket_QRegisterMetaType();};
	 ~MyQUdpSocket() { callbackQUdpSocket_DestroyQUdpSocket(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSocket_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQAbstractSocket_SocketOption(this, option)); };
	bool waitForBytesWritten(int msecs) { return callbackQAbstractSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQAbstractSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQAbstractSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQAbstractSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQAbstractSocket_ReadData(this, dataPacked, maxSize); };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQAbstractSocket_ReadLineData(this, dataPacked, maxlen); };
	qint64 writeData(const char * data, qint64 size) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQAbstractSocket_WriteData(this, dataPacked, size); };
	void close() { callbackQAbstractSocket_Close(this); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQAbstractSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQAbstractSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void Signal_Connected() { callbackQAbstractSocket_Connected(this); };
	void disconnectFromHost() { callbackQAbstractSocket_DisconnectFromHost(this); };
	void Signal_Disconnected() { callbackQAbstractSocket_Disconnected(this); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocket_Error2(this, socketError); };
	void Signal_HostFound() { callbackQAbstractSocket_HostFound(this); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQAbstractSocket_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	void resume() { callbackQAbstractSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocket_SetReadBufferSize(this, size); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocket_StateChanged(this, socketState); };
	bool atEnd() const { return callbackQAbstractSocket_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool canReadLine() const { return callbackQAbstractSocket_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isSequential() const { return callbackQAbstractSocket_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	qint64 bytesAvailable() const { return callbackQAbstractSocket_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQAbstractSocket_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	bool open(QIODevice::OpenMode mode) { return callbackQAbstractSocket_Open(this, mode) != 0; };
	bool reset() { return callbackQAbstractSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQAbstractSocket_Seek(this, pos) != 0; };
	void Signal_AboutToClose() { callbackQAbstractSocket_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQAbstractSocket_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQAbstractSocket_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQAbstractSocket_ChannelReadyRead(this, channel); };
	void Signal_ReadChannelFinished() { callbackQAbstractSocket_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQAbstractSocket_ReadyRead(this); };
	qint64 pos() const { return callbackQAbstractSocket_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQAbstractSocket_Size(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQAbstractSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSocket_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSocket_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSocket_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtNetwork_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSocket_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocket_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQUdpSocket*)

int QUdpSocket_QUdpSocket_QRegisterMetaType(){qRegisterMetaType<QUdpSocket*>(); return qRegisterMetaType<MyQUdpSocket*>();}

void* QUdpSocket_ReceiveDatagram(void* ptr, long long maxSize)
{
	return new QNetworkDatagram(static_cast<QUdpSocket*>(ptr)->receiveDatagram(maxSize));
}

void* QUdpSocket_NewQUdpSocket(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQUdpSocket(static_cast<QWindow*>(parent));
	} else {
		return new MyQUdpSocket(static_cast<QObject*>(parent));
	}
}

char QUdpSocket_JoinMulticastGroup(void* ptr, void* groupAddress)
{
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

char QUdpSocket_JoinMulticastGroup2(void* ptr, void* groupAddress, void* iface)
{
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

char QUdpSocket_LeaveMulticastGroup(void* ptr, void* groupAddress)
{
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

char QUdpSocket_LeaveMulticastGroup2(void* ptr, void* groupAddress, void* iface)
{
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

long long QUdpSocket_ReadDatagram(void* ptr, char* data, long long maxSize, void* address, unsigned short port)
{
	return static_cast<QUdpSocket*>(ptr)->readDatagram(data, maxSize, static_cast<QHostAddress*>(address), &port);
}

long long QUdpSocket_WriteDatagram3(void* ptr, void* datagram, void* host, unsigned short port)
{
	return static_cast<QUdpSocket*>(ptr)->writeDatagram(*static_cast<QByteArray*>(datagram), *static_cast<QHostAddress*>(host), port);
}

long long QUdpSocket_WriteDatagram2(void* ptr, void* datagram)
{
	return static_cast<QUdpSocket*>(ptr)->writeDatagram(*static_cast<QNetworkDatagram*>(datagram));
}

long long QUdpSocket_WriteDatagram(void* ptr, char* data, long long size, void* address, unsigned short port)
{
	return static_cast<QUdpSocket*>(ptr)->writeDatagram(const_cast<const char*>(data), size, *static_cast<QHostAddress*>(address), port);
}

void QUdpSocket_SetMulticastInterface(void* ptr, void* iface)
{
	static_cast<QUdpSocket*>(ptr)->setMulticastInterface(*static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_DestroyQUdpSocket(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->~QUdpSocket();
}

void QUdpSocket_DestroyQUdpSocketDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QUdpSocket_MulticastInterface(void* ptr)
{
	return new QNetworkInterface(static_cast<QUdpSocket*>(ptr)->multicastInterface());
}

char QUdpSocket_HasPendingDatagrams(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->hasPendingDatagrams();
}

long long QUdpSocket_PendingDatagramSize(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->pendingDatagramSize();
}

