// +build !minimal

#define protected public
#define private public

#include "network.h"
#include "_cgo_export.h"

#include <QAbstractNetworkCache>
#include <QAbstractSocket>
#include <QAuthenticator>
#include <QByteArray>
#include <QChildEvent>
#include <QCryptographicHash>
#include <QDate>
#include <QDateTime>
#include <QDnsDomainNameRecord>
#include <QDnsHostAddressRecord>
#include <QDnsLookup>
#include <QDnsMailExchangeRecord>
#include <QDnsServiceRecord>
#include <QDnsTextRecord>
#include <QEvent>
#include <QHostAddress>
#include <QHostInfo>
#include <QHttpMultiPart>
#include <QHttpPart>
#include <QIODevice>
#include <QLocalServer>
#include <QLocalSocket>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNetworkAccessManager>
#include <QNetworkAddressEntry>
#include <QNetworkCacheMetaData>
#include <QNetworkConfiguration>
#include <QNetworkConfigurationManager>
#include <QNetworkCookie>
#include <QNetworkCookieJar>
#include <QNetworkDiskCache>
#include <QNetworkInterface>
#include <QNetworkProxy>
#include <QNetworkProxyFactory>
#include <QNetworkProxyQuery>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QNetworkSession>
#include <QObject>
#include <QSslCertificate>
#include <QSslCertificateExtension>
#include <QSslCipher>
#include <QSslConfiguration>
#include <QSslEllipticCurve>
#include <QSslError>
#include <QSslKey>
#include <QSslPreSharedKeyAuthenticator>
#include <QSslSocket>
#include <QString>
#include <QStringList>
#include <QTcpServer>
#include <QTcpSocket>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUdpSocket>
#include <QUrl>
#include <QVariant>

class MyQAbstractNetworkCache: public QAbstractNetworkCache
{
public:
	MyQAbstractNetworkCache(QObject *parent) : QAbstractNetworkCache(parent) {};
	qint64 cacheSize() const { return static_cast<long long>(callbackQAbstractNetworkCache_CacheSize(const_cast<MyQAbstractNetworkCache*>(this), this->objectName().toUtf8().data())); };
	void clear() { callbackQAbstractNetworkCache_Clear(this, this->objectName().toUtf8().data()); };
	QIODevice * data(const QUrl & url) { return static_cast<QIODevice*>(callbackQAbstractNetworkCache_Data(this, this->objectName().toUtf8().data(), new QUrl(url))); };
	void insert(QIODevice * device) { callbackQAbstractNetworkCache_Insert(this, this->objectName().toUtf8().data(), device); };
	QNetworkCacheMetaData metaData(const QUrl & url) { return *static_cast<QNetworkCacheMetaData*>(callbackQAbstractNetworkCache_MetaData(this, this->objectName().toUtf8().data(), new QUrl(url))); };
	QIODevice * prepare(const QNetworkCacheMetaData & metaData) { return static_cast<QIODevice*>(callbackQAbstractNetworkCache_Prepare(this, this->objectName().toUtf8().data(), new QNetworkCacheMetaData(metaData))); };
	bool remove(const QUrl & url) { return callbackQAbstractNetworkCache_Remove(this, this->objectName().toUtf8().data(), new QUrl(url)) != 0; };
	void updateMetaData(const QNetworkCacheMetaData & metaData) { callbackQAbstractNetworkCache_UpdateMetaData(this, this->objectName().toUtf8().data(), new QNetworkCacheMetaData(metaData)); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractNetworkCache_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractNetworkCache_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQAbstractNetworkCache_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQAbstractNetworkCache_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQAbstractNetworkCache_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractNetworkCache_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractNetworkCache_MetaObject(const_cast<MyQAbstractNetworkCache*>(this), this->objectName().toUtf8().data())); };
};

void* QAbstractNetworkCache_NewQAbstractNetworkCache(void* parent)
{
	return new MyQAbstractNetworkCache(static_cast<QObject*>(parent));
}

long long QAbstractNetworkCache_CacheSize(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractNetworkCache*>(ptr)->cacheSize());
}

void QAbstractNetworkCache_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractNetworkCache*>(ptr), "clear");
}

void* QAbstractNetworkCache_Data(void* ptr, void* url)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void QAbstractNetworkCache_Insert(void* ptr, void* device)
{
	static_cast<QAbstractNetworkCache*>(ptr)->insert(static_cast<QIODevice*>(device));
}

void* QAbstractNetworkCache_MetaData(void* ptr, void* url)
{
	return new QNetworkCacheMetaData(static_cast<QAbstractNetworkCache*>(ptr)->metaData(*static_cast<QUrl*>(url)));
}

void* QAbstractNetworkCache_Prepare(void* ptr, void* metaData)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

int QAbstractNetworkCache_Remove(void* ptr, void* url)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->remove(*static_cast<QUrl*>(url));
}

void QAbstractNetworkCache_UpdateMetaData(void* ptr, void* metaData)
{
	static_cast<QAbstractNetworkCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_DestroyQAbstractNetworkCache(void* ptr)
{
	static_cast<QAbstractNetworkCache*>(ptr)->~QAbstractNetworkCache();
}

void QAbstractNetworkCache_TimerEvent(void* ptr, void* event)
{
	static_cast<QAbstractNetworkCache*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractNetworkCache_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractNetworkCache_ChildEvent(void* ptr, void* event)
{
	static_cast<QAbstractNetworkCache*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractNetworkCache_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractNetworkCache_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractNetworkCache*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractNetworkCache_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractNetworkCache_CustomEvent(void* ptr, void* event)
{
	static_cast<QAbstractNetworkCache*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractNetworkCache_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::customEvent(static_cast<QEvent*>(event));
}

void QAbstractNetworkCache_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractNetworkCache*>(ptr), "deleteLater");
}

void QAbstractNetworkCache_DeleteLaterDefault(void* ptr)
{
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::deleteLater();
}

void QAbstractNetworkCache_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractNetworkCache*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractNetworkCache_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAbstractNetworkCache_Event(void* ptr, void* e)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAbstractNetworkCache_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::event(static_cast<QEvent*>(e));
}

int QAbstractNetworkCache_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAbstractNetworkCache_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAbstractNetworkCache_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractNetworkCache*>(ptr)->metaObject());
}

void* QAbstractNetworkCache_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::metaObject());
}

class MyQAbstractSocket: public QAbstractSocket
{
public:
	MyQAbstractSocket(SocketType socketType, QObject *parent) : QAbstractSocket(socketType, parent) {};
	void Signal_Connected() { callbackQAbstractSocket_Connected(this, this->objectName().toUtf8().data()); };
	void disconnectFromHost() { callbackQAbstractSocket_DisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void Signal_Disconnected() { callbackQAbstractSocket_Disconnected(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocket_Error2(this, this->objectName().toUtf8().data(), socketError); };
	void Signal_HostFound() { callbackQAbstractSocket_HostFound(this, this->objectName().toUtf8().data()); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQAbstractSocket_ProxyAuthenticationRequired(this, this->objectName().toUtf8().data(), new QNetworkProxy(proxy), authenticator); };
	void resume() { callbackQAbstractSocket_Resume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocket_SetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocket_SetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQAbstractSocket_SocketOption(this, this->objectName().toUtf8().data(), option)); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocket_StateChanged(this, this->objectName().toUtf8().data(), socketState); };
	bool waitForConnected(int msecs) { return callbackQAbstractSocket_WaitForConnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQAbstractSocket_WaitForDisconnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQAbstractSocket_Open(this, this->objectName().toUtf8().data(), mode) != 0; };
	qint64 pos() const { return static_cast<long long>(callbackQAbstractSocket_Pos(const_cast<MyQAbstractSocket*>(this), this->objectName().toUtf8().data())); };
	bool reset() { return callbackQAbstractSocket_Reset(this, this->objectName().toUtf8().data()) != 0; };
	bool seek(qint64 pos) { return callbackQAbstractSocket_Seek(this, this->objectName().toUtf8().data(), static_cast<long long>(pos)) != 0; };
	qint64 size() const { return static_cast<long long>(callbackQAbstractSocket_Size(const_cast<MyQAbstractSocket*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocket_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractSocket_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSocket_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQAbstractSocket_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQAbstractSocket_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSocket_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSocket_MetaObject(const_cast<MyQAbstractSocket*>(this), this->objectName().toUtf8().data())); };
};

void* QAbstractSocket_NewQAbstractSocket(int socketType, void* parent)
{
	return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QObject*>(parent));
}

void QAbstractSocket_Abort(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->abort();
}

int QAbstractSocket_AtEnd(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->atEnd();
}

long long QAbstractSocket_BytesAvailable(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->bytesAvailable());
}

long long QAbstractSocket_BytesToWrite(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->bytesToWrite());
}

int QAbstractSocket_CanReadLine(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->canReadLine();
}

void QAbstractSocket_Close(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->close();
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
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::disconnectFromHost();
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
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::error), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketError)>(&MyQAbstractSocket::Signal_Error2));
}

void QAbstractSocket_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::error), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketError)>(&MyQAbstractSocket::Signal_Error2));
}

void QAbstractSocket_Error2(void* ptr, int socketError)
{
	static_cast<QAbstractSocket*>(ptr)->error(static_cast<QAbstractSocket::SocketError>(socketError));
}

int QAbstractSocket_Error(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->error();
}

int QAbstractSocket_Flush(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->flush();
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

int QAbstractSocket_IsSequential(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->isSequential();
}

int QAbstractSocket_IsValid(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->isValid();
}

void* QAbstractSocket_LocalAddress(void* ptr)
{
	return new QHostAddress(static_cast<QAbstractSocket*>(ptr)->localAddress());
}

int QAbstractSocket_PauseMode(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->pauseMode();
}

void* QAbstractSocket_PeerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QAbstractSocket*>(ptr)->peerAddress());
}

char* QAbstractSocket_PeerName(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->peerName().toUtf8().data();
}

void* QAbstractSocket_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QAbstractSocket*>(ptr)->proxy());
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

long long QAbstractSocket_ReadBufferSize(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->readBufferSize());
}

long long QAbstractSocket_ReadLineData(void* ptr, char* data, long long maxlen)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->readLineData(data, static_cast<long long>(maxlen)));
}

void QAbstractSocket_Resume(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->resume();
}

void QAbstractSocket_ResumeDefault(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::resume();
}

void QAbstractSocket_SetLocalAddress(void* ptr, void* address)
{
	static_cast<QAbstractSocket*>(ptr)->setLocalAddress(*static_cast<QHostAddress*>(address));
}

void QAbstractSocket_SetPauseMode(void* ptr, int pauseMode)
{
	static_cast<QAbstractSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QAbstractSocket_SetPeerAddress(void* ptr, void* address)
{
	static_cast<QAbstractSocket*>(ptr)->setPeerAddress(*static_cast<QHostAddress*>(address));
}

void QAbstractSocket_SetPeerName(void* ptr, char* name)
{
	static_cast<QAbstractSocket*>(ptr)->setPeerName(QString(name));
}

void QAbstractSocket_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QAbstractSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QAbstractSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QAbstractSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QAbstractSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setReadBufferSize(static_cast<long long>(size));
}

void QAbstractSocket_SetSocketError(void* ptr, int socketError)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketError(static_cast<QAbstractSocket::SocketError>(socketError));
}

void QAbstractSocket_SetSocketOption(void* ptr, int option, void* value)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QAbstractSocket_SetSocketOptionDefault(void* ptr, int option, void* value)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QAbstractSocket_SetSocketState(void* ptr, int state)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketState(static_cast<QAbstractSocket::SocketState>(state));
}

void* QAbstractSocket_SocketOption(void* ptr, int option)
{
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QAbstractSocket_SocketOptionDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

int QAbstractSocket_SocketType(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->socketType();
}

int QAbstractSocket_State(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->state();
}

void QAbstractSocket_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));
}

void QAbstractSocket_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));
}

void QAbstractSocket_StateChanged(void* ptr, int socketState)
{
	static_cast<QAbstractSocket*>(ptr)->stateChanged(static_cast<QAbstractSocket::SocketState>(socketState));
}

int QAbstractSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QAbstractSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForConnected(msecs);
}

int QAbstractSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForConnected(msecs);
}

int QAbstractSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForDisconnected(msecs);
}

int QAbstractSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForDisconnected(msecs);
}

int QAbstractSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForReadyRead(msecs);
}

long long QAbstractSocket_WriteData(void* ptr, char* data, long long size)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(size)));
}

void QAbstractSocket_DestroyQAbstractSocket(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->~QAbstractSocket();
}

int QAbstractSocket_Open(void* ptr, int mode)
{
	return static_cast<QAbstractSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QAbstractSocket_OpenDefault(void* ptr, int mode)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QAbstractSocket_Pos(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->pos());
}

long long QAbstractSocket_PosDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::pos());
}

int QAbstractSocket_Reset(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->reset();
}

int QAbstractSocket_ResetDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::reset();
}

int QAbstractSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QAbstractSocket*>(ptr)->seek(static_cast<long long>(pos));
}

int QAbstractSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::seek(static_cast<long long>(pos));
}

long long QAbstractSocket_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->size());
}

long long QAbstractSocket_SizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::size());
}

void QAbstractSocket_TimerEvent(void* ptr, void* event)
{
	static_cast<QAbstractSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractSocket_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractSocket_ChildEvent(void* ptr, void* event)
{
	static_cast<QAbstractSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractSocket_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractSocket_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractSocket*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractSocket_CustomEvent(void* ptr, void* event)
{
	static_cast<QAbstractSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractSocket_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::customEvent(static_cast<QEvent*>(event));
}

void QAbstractSocket_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractSocket*>(ptr), "deleteLater");
}

void QAbstractSocket_DeleteLaterDefault(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::deleteLater();
}

void QAbstractSocket_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractSocket*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAbstractSocket_Event(void* ptr, void* e)
{
	return static_cast<QAbstractSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAbstractSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::event(static_cast<QEvent*>(e));
}

int QAbstractSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAbstractSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAbstractSocket_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractSocket*>(ptr)->metaObject());
}

void* QAbstractSocket_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::metaObject());
}

void* QAuthenticator_NewQAuthenticator()
{
	return new QAuthenticator();
}

void* QAuthenticator_NewQAuthenticator2(void* other)
{
	return new QAuthenticator(*static_cast<QAuthenticator*>(other));
}

int QAuthenticator_IsNull(void* ptr)
{
	return static_cast<QAuthenticator*>(ptr)->isNull();
}

void* QAuthenticator_Option(void* ptr, char* opt)
{
	return new QVariant(static_cast<QAuthenticator*>(ptr)->option(QString(opt)));
}

char* QAuthenticator_Password(void* ptr)
{
	return static_cast<QAuthenticator*>(ptr)->password().toUtf8().data();
}

char* QAuthenticator_Realm(void* ptr)
{
	return static_cast<QAuthenticator*>(ptr)->realm().toUtf8().data();
}

void QAuthenticator_SetOption(void* ptr, char* opt, void* value)
{
	static_cast<QAuthenticator*>(ptr)->setOption(QString(opt), *static_cast<QVariant*>(value));
}

void QAuthenticator_SetPassword(void* ptr, char* password)
{
	static_cast<QAuthenticator*>(ptr)->setPassword(QString(password));
}

void QAuthenticator_SetUser(void* ptr, char* user)
{
	static_cast<QAuthenticator*>(ptr)->setUser(QString(user));
}

char* QAuthenticator_User(void* ptr)
{
	return static_cast<QAuthenticator*>(ptr)->user().toUtf8().data();
}

void QAuthenticator_DestroyQAuthenticator(void* ptr)
{
	static_cast<QAuthenticator*>(ptr)->~QAuthenticator();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord()
{
	return new QDnsDomainNameRecord();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord2(void* other)
{
	return new QDnsDomainNameRecord(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Name(void* ptr)
{
	return static_cast<QDnsDomainNameRecord*>(ptr)->name().toUtf8().data();
}

void QDnsDomainNameRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsDomainNameRecord*>(ptr)->swap(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Value(void* ptr)
{
	return static_cast<QDnsDomainNameRecord*>(ptr)->value().toUtf8().data();
}

void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(void* ptr)
{
	static_cast<QDnsDomainNameRecord*>(ptr)->~QDnsDomainNameRecord();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord()
{
	return new QDnsHostAddressRecord();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord2(void* other)
{
	return new QDnsHostAddressRecord(*static_cast<QDnsHostAddressRecord*>(other));
}

char* QDnsHostAddressRecord_Name(void* ptr)
{
	return static_cast<QDnsHostAddressRecord*>(ptr)->name().toUtf8().data();
}

void QDnsHostAddressRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsHostAddressRecord*>(ptr)->swap(*static_cast<QDnsHostAddressRecord*>(other));
}

void* QDnsHostAddressRecord_Value(void* ptr)
{
	return new QHostAddress(static_cast<QDnsHostAddressRecord*>(ptr)->value());
}

void QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(void* ptr)
{
	static_cast<QDnsHostAddressRecord*>(ptr)->~QDnsHostAddressRecord();
}

class MyQDnsLookup: public QDnsLookup
{
public:
	MyQDnsLookup(Type type, const QString &name, const QHostAddress &nameserver, QObject *parent) : QDnsLookup(type, name, nameserver, parent) {};
	MyQDnsLookup(QObject *parent) : QDnsLookup(parent) {};
	MyQDnsLookup(Type type, const QString &name, QObject *parent) : QDnsLookup(type, name, parent) {};
	void abort() { callbackQDnsLookup_Abort(this, this->objectName().toUtf8().data()); };
	void Signal_Finished() { callbackQDnsLookup_Finished(this, this->objectName().toUtf8().data()); };
	void lookup() { callbackQDnsLookup_Lookup(this, this->objectName().toUtf8().data()); };
	void Signal_NameChanged(const QString & name) { callbackQDnsLookup_NameChanged(this, this->objectName().toUtf8().data(), name.toUtf8().data()); };
	void Signal_NameserverChanged(const QHostAddress & nameserver) { callbackQDnsLookup_NameserverChanged(this, this->objectName().toUtf8().data(), new QHostAddress(nameserver)); };
	void Signal_TypeChanged(QDnsLookup::Type ty) { callbackQDnsLookup_TypeChanged(this, this->objectName().toUtf8().data(), ty); };
	void timerEvent(QTimerEvent * event) { callbackQDnsLookup_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDnsLookup_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDnsLookup_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDnsLookup_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDnsLookup_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDnsLookup_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQDnsLookup_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDnsLookup_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDnsLookup_MetaObject(const_cast<MyQDnsLookup*>(this), this->objectName().toUtf8().data())); };
};

void* QDnsLookup_NewQDnsLookup3(int ty, char* name, void* nameserver, void* parent)
{
	return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), *static_cast<QHostAddress*>(nameserver), static_cast<QObject*>(parent));
}

int QDnsLookup_Error(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->error();
}

char* QDnsLookup_ErrorString(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->errorString().toUtf8().data();
}

char* QDnsLookup_Name(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->name().toUtf8().data();
}

void* QDnsLookup_Nameserver(void* ptr)
{
	return new QHostAddress(static_cast<QDnsLookup*>(ptr)->nameserver());
}

void QDnsLookup_SetName(void* ptr, char* name)
{
	static_cast<QDnsLookup*>(ptr)->setName(QString(name));
}

void QDnsLookup_SetNameserver(void* ptr, void* nameserver)
{
	static_cast<QDnsLookup*>(ptr)->setNameserver(*static_cast<QHostAddress*>(nameserver));
}

void QDnsLookup_SetType(void* ptr, int vqd)
{
	static_cast<QDnsLookup*>(ptr)->setType(static_cast<QDnsLookup::Type>(vqd));
}

int QDnsLookup_Type(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->type();
}

void* QDnsLookup_NewQDnsLookup(void* parent)
{
	return new MyQDnsLookup(static_cast<QObject*>(parent));
}

void* QDnsLookup_NewQDnsLookup2(int ty, char* name, void* parent)
{
	return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), static_cast<QObject*>(parent));
}

void QDnsLookup_Abort(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "abort");
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

int QDnsLookup_IsFinished(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->isFinished();
}

void QDnsLookup_Lookup(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "lookup");
}

void QDnsLookup_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));
}

void QDnsLookup_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));
}

void QDnsLookup_NameChanged(void* ptr, char* name)
{
	static_cast<QDnsLookup*>(ptr)->nameChanged(QString(name));
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

void QDnsLookup_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));
}

void QDnsLookup_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));
}

void QDnsLookup_TypeChanged(void* ptr, int ty)
{
	static_cast<QDnsLookup*>(ptr)->typeChanged(static_cast<QDnsLookup::Type>(ty));
}

void QDnsLookup_DestroyQDnsLookup(void* ptr)
{
	static_cast<QDnsLookup*>(ptr)->~QDnsLookup();
}

void QDnsLookup_TimerEvent(void* ptr, void* event)
{
	static_cast<QDnsLookup*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDnsLookup_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDnsLookup_ChildEvent(void* ptr, void* event)
{
	static_cast<QDnsLookup*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDnsLookup_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::childEvent(static_cast<QChildEvent*>(event));
}

void QDnsLookup_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDnsLookup*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDnsLookup_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDnsLookup_CustomEvent(void* ptr, void* event)
{
	static_cast<QDnsLookup*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDnsLookup_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::customEvent(static_cast<QEvent*>(event));
}

void QDnsLookup_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "deleteLater");
}

void QDnsLookup_DeleteLaterDefault(void* ptr)
{
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::deleteLater();
}

void QDnsLookup_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDnsLookup*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDnsLookup_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDnsLookup_Event(void* ptr, void* e)
{
	return static_cast<QDnsLookup*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDnsLookup_EventDefault(void* ptr, void* e)
{
	return static_cast<QDnsLookup*>(ptr)->QDnsLookup::event(static_cast<QEvent*>(e));
}

int QDnsLookup_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDnsLookup*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDnsLookup_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDnsLookup*>(ptr)->QDnsLookup::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDnsLookup_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDnsLookup*>(ptr)->metaObject());
}

void* QDnsLookup_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDnsLookup*>(ptr)->QDnsLookup::metaObject());
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord()
{
	return new QDnsMailExchangeRecord();
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(void* other)
{
	return new QDnsMailExchangeRecord(*static_cast<QDnsMailExchangeRecord*>(other));
}

char* QDnsMailExchangeRecord_Exchange(void* ptr)
{
	return static_cast<QDnsMailExchangeRecord*>(ptr)->exchange().toUtf8().data();
}

char* QDnsMailExchangeRecord_Name(void* ptr)
{
	return static_cast<QDnsMailExchangeRecord*>(ptr)->name().toUtf8().data();
}

void QDnsMailExchangeRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsMailExchangeRecord*>(ptr)->swap(*static_cast<QDnsMailExchangeRecord*>(other));
}

void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(void* ptr)
{
	static_cast<QDnsMailExchangeRecord*>(ptr)->~QDnsMailExchangeRecord();
}

void* QDnsServiceRecord_NewQDnsServiceRecord()
{
	return new QDnsServiceRecord();
}

void* QDnsServiceRecord_NewQDnsServiceRecord2(void* other)
{
	return new QDnsServiceRecord(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Name(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->name().toUtf8().data();
}

void QDnsServiceRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsServiceRecord*>(ptr)->swap(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Target(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->target().toUtf8().data();
}

void QDnsServiceRecord_DestroyQDnsServiceRecord(void* ptr)
{
	static_cast<QDnsServiceRecord*>(ptr)->~QDnsServiceRecord();
}

void* QDnsTextRecord_NewQDnsTextRecord()
{
	return new QDnsTextRecord();
}

void* QDnsTextRecord_NewQDnsTextRecord2(void* other)
{
	return new QDnsTextRecord(*static_cast<QDnsTextRecord*>(other));
}

char* QDnsTextRecord_Name(void* ptr)
{
	return static_cast<QDnsTextRecord*>(ptr)->name().toUtf8().data();
}

void QDnsTextRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsTextRecord*>(ptr)->swap(*static_cast<QDnsTextRecord*>(other));
}

void QDnsTextRecord_DestroyQDnsTextRecord(void* ptr)
{
	static_cast<QDnsTextRecord*>(ptr)->~QDnsTextRecord();
}

void* QHostAddress_NewQHostAddress()
{
	return new QHostAddress();
}

void* QHostAddress_NewQHostAddress9(int address)
{
	return new QHostAddress(static_cast<QHostAddress::SpecialAddress>(address));
}

void* QHostAddress_NewQHostAddress8(void* address)
{
	return new QHostAddress(*static_cast<QHostAddress*>(address));
}

void* QHostAddress_NewQHostAddress7(char* address)
{
	return new QHostAddress(QString(address));
}

void QHostAddress_Clear(void* ptr)
{
	static_cast<QHostAddress*>(ptr)->clear();
}

int QHostAddress_IsInSubnet(void* ptr, void* subnet, int netmask)
{
	return static_cast<QHostAddress*>(ptr)->isInSubnet(*static_cast<QHostAddress*>(subnet), netmask);
}

int QHostAddress_IsLoopback(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isLoopback();
}

int QHostAddress_IsMulticast(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isMulticast();
}

int QHostAddress_IsNull(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->isNull();
}

int QHostAddress_Protocol(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->protocol();
}

char* QHostAddress_ScopeId(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->scopeId().toUtf8().data();
}

int QHostAddress_SetAddress6(void* ptr, char* address)
{
	return static_cast<QHostAddress*>(ptr)->setAddress(QString(address));
}

void QHostAddress_SetScopeId(void* ptr, char* id)
{
	static_cast<QHostAddress*>(ptr)->setScopeId(QString(id));
}

void QHostAddress_Swap(void* ptr, void* other)
{
	static_cast<QHostAddress*>(ptr)->swap(*static_cast<QHostAddress*>(other));
}

char* QHostAddress_ToString(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->toString().toUtf8().data();
}

void QHostAddress_DestroyQHostAddress(void* ptr)
{
	static_cast<QHostAddress*>(ptr)->~QHostAddress();
}

char* QHostInfo_QHostInfo_LocalHostName()
{
	return QHostInfo::localHostName().toUtf8().data();
}

void* QHostInfo_NewQHostInfo2(void* other)
{
	return new QHostInfo(*static_cast<QHostInfo*>(other));
}

void* QHostInfo_NewQHostInfo(int id)
{
	return new QHostInfo(id);
}

void QHostInfo_QHostInfo_AbortHostLookup(int id)
{
	QHostInfo::abortHostLookup(id);
}

int QHostInfo_Error(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->error();
}

char* QHostInfo_ErrorString(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->errorString().toUtf8().data();
}

void* QHostInfo_QHostInfo_FromName(char* name)
{
	return new QHostInfo(QHostInfo::fromName(QString(name)));
}

char* QHostInfo_HostName(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->hostName().toUtf8().data();
}

int QHostInfo_QHostInfo_LookupHost(char* name, void* receiver, char* member)
{
	return QHostInfo::lookupHost(QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QHostInfo_LookupId(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->lookupId();
}

void QHostInfo_SetError(void* ptr, int error)
{
	static_cast<QHostInfo*>(ptr)->setError(static_cast<QHostInfo::HostInfoError>(error));
}

void QHostInfo_SetErrorString(void* ptr, char* str)
{
	static_cast<QHostInfo*>(ptr)->setErrorString(QString(str));
}

void QHostInfo_SetHostName(void* ptr, char* hostName)
{
	static_cast<QHostInfo*>(ptr)->setHostName(QString(hostName));
}

void QHostInfo_SetLookupId(void* ptr, int id)
{
	static_cast<QHostInfo*>(ptr)->setLookupId(id);
}

void QHostInfo_DestroyQHostInfo(void* ptr)
{
	static_cast<QHostInfo*>(ptr)->~QHostInfo();
}

char* QHostInfo_QHostInfo_LocalDomainName()
{
	return QHostInfo::localDomainName().toUtf8().data();
}

void* QHttpMultiPart_NewQHttpMultiPart2(int contentType, void* parent)
{
	return new QHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QObject*>(parent));
}

void* QHttpMultiPart_NewQHttpMultiPart(void* parent)
{
	return new QHttpMultiPart(static_cast<QObject*>(parent));
}

void QHttpMultiPart_Append(void* ptr, void* httpPart)
{
	static_cast<QHttpMultiPart*>(ptr)->append(*static_cast<QHttpPart*>(httpPart));
}

char* QHttpMultiPart_Boundary(void* ptr)
{
	return QString(static_cast<QHttpMultiPart*>(ptr)->boundary()).toUtf8().data();
}

void QHttpMultiPart_SetBoundary(void* ptr, char* boundary)
{
	static_cast<QHttpMultiPart*>(ptr)->setBoundary(QByteArray(boundary));
}

void QHttpMultiPart_SetContentType(void* ptr, int contentType)
{
	static_cast<QHttpMultiPart*>(ptr)->setContentType(static_cast<QHttpMultiPart::ContentType>(contentType));
}

void QHttpMultiPart_DestroyQHttpMultiPart(void* ptr)
{
	static_cast<QHttpMultiPart*>(ptr)->~QHttpMultiPart();
}

void QHttpMultiPart_TimerEvent(void* ptr, void* event)
{
	static_cast<QHttpMultiPart*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHttpMultiPart_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHttpMultiPart_ChildEvent(void* ptr, void* event)
{
	static_cast<QHttpMultiPart*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHttpMultiPart_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::childEvent(static_cast<QChildEvent*>(event));
}

void QHttpMultiPart_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHttpMultiPart*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHttpMultiPart_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHttpMultiPart_CustomEvent(void* ptr, void* event)
{
	static_cast<QHttpMultiPart*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHttpMultiPart_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::customEvent(static_cast<QEvent*>(event));
}

void QHttpMultiPart_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHttpMultiPart*>(ptr), "deleteLater");
}

void QHttpMultiPart_DeleteLaterDefault(void* ptr)
{
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::deleteLater();
}

void QHttpMultiPart_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHttpMultiPart*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHttpMultiPart_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHttpMultiPart_Event(void* ptr, void* e)
{
	return static_cast<QHttpMultiPart*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHttpMultiPart_EventDefault(void* ptr, void* e)
{
	return static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::event(static_cast<QEvent*>(e));
}

int QHttpMultiPart_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHttpMultiPart*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHttpMultiPart_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHttpMultiPart_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHttpMultiPart*>(ptr)->metaObject());
}

void* QHttpMultiPart_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::metaObject());
}

void* QHttpPart_NewQHttpPart()
{
	return new QHttpPart();
}

void* QHttpPart_NewQHttpPart2(void* other)
{
	return new QHttpPart(*static_cast<QHttpPart*>(other));
}

void QHttpPart_SetBody(void* ptr, char* body)
{
	static_cast<QHttpPart*>(ptr)->setBody(QByteArray(body));
}

void QHttpPart_SetBodyDevice(void* ptr, void* device)
{
	static_cast<QHttpPart*>(ptr)->setBodyDevice(static_cast<QIODevice*>(device));
}

void QHttpPart_SetHeader(void* ptr, int header, void* value)
{
	static_cast<QHttpPart*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QHttpPart_SetRawHeader(void* ptr, char* headerName, char* headerValue)
{
	static_cast<QHttpPart*>(ptr)->setRawHeader(QByteArray(headerName), QByteArray(headerValue));
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
	MyQLocalServer(QObject *parent) : QLocalServer(parent) {};
	bool hasPendingConnections() const { return callbackQLocalServer_HasPendingConnections(const_cast<MyQLocalServer*>(this), this->objectName().toUtf8().data()) != 0; };
	void Signal_NewConnection() { callbackQLocalServer_NewConnection(this, this->objectName().toUtf8().data()); };
	QLocalSocket * nextPendingConnection() { return static_cast<QLocalSocket*>(callbackQLocalServer_NextPendingConnection(this, this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQLocalServer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLocalServer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLocalServer_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQLocalServer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQLocalServer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLocalServer_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQLocalServer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLocalServer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLocalServer_MetaObject(const_cast<MyQLocalServer*>(this), this->objectName().toUtf8().data())); };
};

void QLocalServer_SetSocketOptions(void* ptr, int options)
{
	static_cast<QLocalServer*>(ptr)->setSocketOptions(static_cast<QLocalServer::SocketOption>(options));
}

void* QLocalServer_NewQLocalServer(void* parent)
{
	return new MyQLocalServer(static_cast<QObject*>(parent));
}

void QLocalServer_Close(void* ptr)
{
	static_cast<QLocalServer*>(ptr)->close();
}

char* QLocalServer_ErrorString(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->errorString().toUtf8().data();
}

char* QLocalServer_FullServerName(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalServer_HasPendingConnections(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->hasPendingConnections();
}

int QLocalServer_HasPendingConnectionsDefault(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->QLocalServer::hasPendingConnections();
}

int QLocalServer_IsListening(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->isListening();
}

int QLocalServer_Listen(void* ptr, char* name)
{
	return static_cast<QLocalServer*>(ptr)->listen(QString(name));
}

int QLocalServer_MaxPendingConnections(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->maxPendingConnections();
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

void* QLocalServer_NextPendingConnection(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->nextPendingConnection();
}

void* QLocalServer_NextPendingConnectionDefault(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->QLocalServer::nextPendingConnection();
}

int QLocalServer_QLocalServer_RemoveServer(char* name)
{
	return QLocalServer::removeServer(QString(name));
}

int QLocalServer_ServerError(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->serverError();
}

char* QLocalServer_ServerName(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->serverName().toUtf8().data();
}

void QLocalServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QLocalServer*>(ptr)->setMaxPendingConnections(numConnections);
}

int QLocalServer_SocketOptions(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->socketOptions();
}

int QLocalServer_WaitForNewConnection(void* ptr, int msec, int timedOut)
{
	return static_cast<QLocalServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QLocalServer_DestroyQLocalServer(void* ptr)
{
	static_cast<QLocalServer*>(ptr)->~QLocalServer();
}

void QLocalServer_TimerEvent(void* ptr, void* event)
{
	static_cast<QLocalServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalServer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalServer_ChildEvent(void* ptr, void* event)
{
	static_cast<QLocalServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLocalServer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::childEvent(static_cast<QChildEvent*>(event));
}

void QLocalServer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QLocalServer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalServer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalServer_CustomEvent(void* ptr, void* event)
{
	static_cast<QLocalServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLocalServer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::customEvent(static_cast<QEvent*>(event));
}

void QLocalServer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLocalServer*>(ptr), "deleteLater");
}

void QLocalServer_DeleteLaterDefault(void* ptr)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::deleteLater();
}

void QLocalServer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QLocalServer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QLocalServer_Event(void* ptr, void* e)
{
	return static_cast<QLocalServer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QLocalServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QLocalServer*>(ptr)->QLocalServer::event(static_cast<QEvent*>(e));
}

int QLocalServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLocalServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QLocalServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QLocalServer*>(ptr)->QLocalServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QLocalServer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLocalServer*>(ptr)->metaObject());
}

void* QLocalServer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLocalServer*>(ptr)->QLocalServer::metaObject());
}

class MyQLocalSocket: public QLocalSocket
{
public:
	MyQLocalSocket(QObject *parent) : QLocalSocket(parent) {};
	bool open(QIODevice::OpenMode openMode) { return callbackQLocalSocket_Open(this, this->objectName().toUtf8().data(), openMode) != 0; };
	void Signal_Connected() { callbackQLocalSocket_Connected(this, this->objectName().toUtf8().data()); };
	void Signal_Disconnected() { callbackQLocalSocket_Disconnected(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QLocalSocket::LocalSocketError socketError) { callbackQLocalSocket_Error2(this, this->objectName().toUtf8().data(), socketError); };
	bool isSequential() const { return callbackQLocalSocket_IsSequential(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data()) != 0; };
	void Signal_StateChanged(QLocalSocket::LocalSocketState socketState) { callbackQLocalSocket_StateChanged(this, this->objectName().toUtf8().data(), socketState); };
	qint64 bytesAvailable() const { return static_cast<long long>(callbackQLocalSocket_BytesAvailable(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data())); };
	qint64 bytesToWrite() const { return static_cast<long long>(callbackQLocalSocket_BytesToWrite(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data())); };
	bool canReadLine() const { return callbackQLocalSocket_CanReadLine(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data()) != 0; };
	void close() { callbackQLocalSocket_Close(this, this->objectName().toUtf8().data()); };
	qint64 writeData(const char * data, qint64 c) { return static_cast<long long>(callbackQLocalSocket_WriteData(this, this->objectName().toUtf8().data(), QString(data).toUtf8().data(), static_cast<long long>(c))); };
	bool atEnd() const { return callbackQLocalSocket_AtEnd(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data()) != 0; };
	qint64 pos() const { return static_cast<long long>(callbackQLocalSocket_Pos(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data())); };
	qint64 readLineData(char * data, qint64 maxSize) { return static_cast<long long>(callbackQLocalSocket_ReadLineData(this, this->objectName().toUtf8().data(), QString(data).toUtf8().data(), static_cast<long long>(maxSize))); };
	bool reset() { return callbackQLocalSocket_Reset(this, this->objectName().toUtf8().data()) != 0; };
	bool seek(qint64 pos) { return callbackQLocalSocket_Seek(this, this->objectName().toUtf8().data(), static_cast<long long>(pos)) != 0; };
	qint64 size() const { return static_cast<long long>(callbackQLocalSocket_Size(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQLocalSocket_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLocalSocket_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLocalSocket_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQLocalSocket_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQLocalSocket_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLocalSocket_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQLocalSocket_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLocalSocket_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLocalSocket_MetaObject(const_cast<MyQLocalSocket*>(this), this->objectName().toUtf8().data())); };
};

int QLocalSocket_Open(void* ptr, int openMode)
{
	return static_cast<QLocalSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

int QLocalSocket_OpenDefault(void* ptr, int openMode)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QLocalSocket_NewQLocalSocket(void* parent)
{
	return new MyQLocalSocket(static_cast<QObject*>(parent));
}

void QLocalSocket_ConnectToServer2(void* ptr, char* name, int openMode)
{
	static_cast<QLocalSocket*>(ptr)->connectToServer(QString(name), static_cast<QIODevice::OpenModeFlag>(openMode));
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

void QLocalSocket_Error2(void* ptr, int socketError)
{
	static_cast<QLocalSocket*>(ptr)->error(static_cast<QLocalSocket::LocalSocketError>(socketError));
}

char* QLocalSocket_FullServerName(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalSocket_IsSequential(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->isSequential();
}

int QLocalSocket_IsSequentialDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::isSequential();
}

char* QLocalSocket_ServerName(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->serverName().toUtf8().data();
}

void QLocalSocket_SetServerName(void* ptr, char* name)
{
	static_cast<QLocalSocket*>(ptr)->setServerName(QString(name));
}

int QLocalSocket_State(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->state();
}

void QLocalSocket_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));
}

void QLocalSocket_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));
}

void QLocalSocket_StateChanged(void* ptr, int socketState)
{
	static_cast<QLocalSocket*>(ptr)->stateChanged(static_cast<QLocalSocket::LocalSocketState>(socketState));
}

void QLocalSocket_DestroyQLocalSocket(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->~QLocalSocket();
}

void QLocalSocket_Abort(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->abort();
}

long long QLocalSocket_BytesAvailable(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->bytesAvailable());
}

long long QLocalSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::bytesAvailable());
}

long long QLocalSocket_BytesToWrite(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->bytesToWrite());
}

long long QLocalSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::bytesToWrite());
}

int QLocalSocket_CanReadLine(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->canReadLine();
}

int QLocalSocket_CanReadLineDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::canReadLine();
}

void QLocalSocket_Close(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->close();
}

void QLocalSocket_CloseDefault(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::close();
}

void QLocalSocket_ConnectToServer(void* ptr, int openMode)
{
	static_cast<QLocalSocket*>(ptr)->connectToServer(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_DisconnectFromServer(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->disconnectFromServer();
}

int QLocalSocket_Error(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->error();
}

int QLocalSocket_Flush(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->flush();
}

int QLocalSocket_IsValid(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->isValid();
}

long long QLocalSocket_ReadBufferSize(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->readBufferSize());
}

void QLocalSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QLocalSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

int QLocalSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QLocalSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForConnected(msecs);
}

int QLocalSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForDisconnected(msecs);
}

int QLocalSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForReadyRead(msecs);
}

long long QLocalSocket_WriteData(void* ptr, char* data, long long c)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(c)));
}

long long QLocalSocket_WriteDataDefault(void* ptr, char* data, long long c)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::writeData(const_cast<const char*>(data), static_cast<long long>(c)));
}

int QLocalSocket_AtEnd(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->atEnd();
}

int QLocalSocket_AtEndDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::atEnd();
}

long long QLocalSocket_Pos(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->pos());
}

long long QLocalSocket_PosDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::pos());
}

long long QLocalSocket_ReadLineData(void* ptr, char* data, long long maxSize)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->readLineData(data, static_cast<long long>(maxSize)));
}

long long QLocalSocket_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::readLineData(data, static_cast<long long>(maxSize)));
}

int QLocalSocket_Reset(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->reset();
}

int QLocalSocket_ResetDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::reset();
}

int QLocalSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QLocalSocket*>(ptr)->seek(static_cast<long long>(pos));
}

int QLocalSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::seek(static_cast<long long>(pos));
}

long long QLocalSocket_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->size());
}

long long QLocalSocket_SizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::size());
}

void QLocalSocket_TimerEvent(void* ptr, void* event)
{
	static_cast<QLocalSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalSocket_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalSocket_ChildEvent(void* ptr, void* event)
{
	static_cast<QLocalSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLocalSocket_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QLocalSocket_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QLocalSocket*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalSocket_CustomEvent(void* ptr, void* event)
{
	static_cast<QLocalSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLocalSocket_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::customEvent(static_cast<QEvent*>(event));
}

void QLocalSocket_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLocalSocket*>(ptr), "deleteLater");
}

void QLocalSocket_DeleteLaterDefault(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::deleteLater();
}

void QLocalSocket_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QLocalSocket*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLocalSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QLocalSocket_Event(void* ptr, void* e)
{
	return static_cast<QLocalSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

int QLocalSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::event(static_cast<QEvent*>(e));
}

int QLocalSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLocalSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QLocalSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QLocalSocket_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLocalSocket*>(ptr)->metaObject());
}

void* QLocalSocket_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLocalSocket*>(ptr)->QLocalSocket::metaObject());
}

class MyQNetworkAccessManager: public QNetworkAccessManager
{
public:
	MyQNetworkAccessManager(QObject *parent) : QNetworkAccessManager(parent) {};
	void Signal_AuthenticationRequired(QNetworkReply * reply, QAuthenticator * authenticator) { callbackQNetworkAccessManager_AuthenticationRequired(this, this->objectName().toUtf8().data(), reply, authenticator); };
	QNetworkReply * createRequest(QNetworkAccessManager::Operation op, const QNetworkRequest & req, QIODevice * outgoingData) { return static_cast<QNetworkReply*>(callbackQNetworkAccessManager_CreateRequest(this, this->objectName().toUtf8().data(), op, new QNetworkRequest(req), outgoingData)); };
	void Signal_Encrypted(QNetworkReply * reply) { callbackQNetworkAccessManager_Encrypted(this, this->objectName().toUtf8().data(), reply); };
	void Signal_Finished(QNetworkReply * reply) { callbackQNetworkAccessManager_Finished(this, this->objectName().toUtf8().data(), reply); };
	void Signal_NetworkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility accessible) { callbackQNetworkAccessManager_NetworkAccessibleChanged(this, this->objectName().toUtf8().data(), accessible); };
	void Signal_PreSharedKeyAuthenticationRequired(QNetworkReply * reply, QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired(this, this->objectName().toUtf8().data(), reply, authenticator); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQNetworkAccessManager_ProxyAuthenticationRequired(this, this->objectName().toUtf8().data(), new QNetworkProxy(proxy), authenticator); };
	QStringList supportedSchemesImplementation() const { return QString(callbackQNetworkAccessManager_SupportedSchemesImplementation(const_cast<MyQNetworkAccessManager*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkAccessManager_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkAccessManager_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkAccessManager_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQNetworkAccessManager_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQNetworkAccessManager_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkAccessManager_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQNetworkAccessManager_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkAccessManager_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkAccessManager_MetaObject(const_cast<MyQNetworkAccessManager*>(this), this->objectName().toUtf8().data())); };
};

void* QNetworkAccessManager_ProxyFactory(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->proxyFactory();
}

void* QNetworkAccessManager_NewQNetworkAccessManager(void* parent)
{
	return new MyQNetworkAccessManager(static_cast<QObject*>(parent));
}

void* QNetworkAccessManager_ActiveConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkAccessManager*>(ptr)->activeConfiguration());
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

void* QNetworkAccessManager_Cache(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->cache();
}

void QNetworkAccessManager_ClearAccessCache(void* ptr)
{
	static_cast<QNetworkAccessManager*>(ptr)->clearAccessCache();
}

void* QNetworkAccessManager_Configuration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkAccessManager*>(ptr)->configuration());
}

void* QNetworkAccessManager_CookieJar(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->cookieJar();
}

void* QNetworkAccessManager_CreateRequest(void* ptr, int op, void* req, void* outgoingData)
{
	return static_cast<QNetworkAccessManager*>(ptr)->createRequest(static_cast<QNetworkAccessManager::Operation>(op), *static_cast<QNetworkRequest*>(req), static_cast<QIODevice*>(outgoingData));
}

void* QNetworkAccessManager_CreateRequestDefault(void* ptr, int op, void* req, void* outgoingData)
{
	return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::createRequest(static_cast<QNetworkAccessManager::Operation>(op), *static_cast<QNetworkRequest*>(req), static_cast<QIODevice*>(outgoingData));
}

void* QNetworkAccessManager_DeleteResource(void* ptr, void* request)
{
	return static_cast<QNetworkAccessManager*>(ptr)->deleteResource(*static_cast<QNetworkRequest*>(request));
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

void* QNetworkAccessManager_Get(void* ptr, void* request)
{
	return static_cast<QNetworkAccessManager*>(ptr)->get(*static_cast<QNetworkRequest*>(request));
}

void* QNetworkAccessManager_Head(void* ptr, void* request)
{
	return static_cast<QNetworkAccessManager*>(ptr)->head(*static_cast<QNetworkRequest*>(request));
}

int QNetworkAccessManager_NetworkAccessible(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->networkAccessible();
}

void QNetworkAccessManager_ConnectNetworkAccessibleChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));
}

void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));
}

void QNetworkAccessManager_NetworkAccessibleChanged(void* ptr, int accessible)
{
	static_cast<QNetworkAccessManager*>(ptr)->networkAccessibleChanged(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void* QNetworkAccessManager_Post3(void* ptr, void* request, void* multiPart)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Post(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Post2(void* ptr, void* request, char* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), QByteArray(data));
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

void* QNetworkAccessManager_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QNetworkAccessManager*>(ptr)->proxy());
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

void* QNetworkAccessManager_Put3(void* ptr, void* request, void* multiPart)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Put(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Put2(void* ptr, void* request, char* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), QByteArray(data));
}

void* QNetworkAccessManager_SendCustomRequest(void* ptr, void* request, char* verb, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), QByteArray(verb), static_cast<QIODevice*>(data));
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

void QNetworkAccessManager_SetNetworkAccessible(void* ptr, int accessible)
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

char* QNetworkAccessManager_SupportedSchemes(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->supportedSchemes().join("|").toUtf8().data();
}

char* QNetworkAccessManager_SupportedSchemesImplementation(void* ptr)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<QNetworkAccessManager*>(ptr), "supportedSchemesImplementation", Q_RETURN_ARG(QStringList, returnArg));
	return returnArg.join("|").toUtf8().data();
}

void QNetworkAccessManager_DestroyQNetworkAccessManager(void* ptr)
{
	static_cast<QNetworkAccessManager*>(ptr)->~QNetworkAccessManager();
}

void QNetworkAccessManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QNetworkAccessManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkAccessManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkAccessManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QNetworkAccessManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkAccessManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkAccessManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkAccessManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkAccessManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkAccessManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QNetworkAccessManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkAccessManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::customEvent(static_cast<QEvent*>(event));
}

void QNetworkAccessManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkAccessManager*>(ptr), "deleteLater");
}

void QNetworkAccessManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::deleteLater();
}

void QNetworkAccessManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkAccessManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkAccessManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QNetworkAccessManager_Event(void* ptr, void* e)
{
	return static_cast<QNetworkAccessManager*>(ptr)->event(static_cast<QEvent*>(e));
}

int QNetworkAccessManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::event(static_cast<QEvent*>(e));
}

int QNetworkAccessManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkAccessManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QNetworkAccessManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNetworkAccessManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkAccessManager*>(ptr)->metaObject());
}

void* QNetworkAccessManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::metaObject());
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry()
{
	return new QNetworkAddressEntry();
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry2(void* other)
{
	return new QNetworkAddressEntry(*static_cast<QNetworkAddressEntry*>(other));
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

int QNetworkAddressEntry_PrefixLength(void* ptr)
{
	return static_cast<QNetworkAddressEntry*>(ptr)->prefixLength();
}

void QNetworkAddressEntry_SetBroadcast(void* ptr, void* newBroadcast)
{
	static_cast<QNetworkAddressEntry*>(ptr)->setBroadcast(*static_cast<QHostAddress*>(newBroadcast));
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

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData()
{
	return new QNetworkCacheMetaData();
}

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData2(void* other)
{
	return new QNetworkCacheMetaData(*static_cast<QNetworkCacheMetaData*>(other));
}

void* QNetworkCacheMetaData_ExpirationDate(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->expirationDate());
}

int QNetworkCacheMetaData_IsValid(void* ptr)
{
	return static_cast<QNetworkCacheMetaData*>(ptr)->isValid();
}

void* QNetworkCacheMetaData_LastModified(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->lastModified());
}

int QNetworkCacheMetaData_SaveToDisk(void* ptr)
{
	return static_cast<QNetworkCacheMetaData*>(ptr)->saveToDisk();
}

void QNetworkCacheMetaData_SetExpirationDate(void* ptr, void* dateTime)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetLastModified(void* ptr, void* dateTime)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->setLastModified(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetSaveToDisk(void* ptr, int allow)
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

void* QNetworkCacheMetaData_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkCacheMetaData*>(ptr)->url());
}

void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(void* ptr)
{
	static_cast<QNetworkCacheMetaData*>(ptr)->~QNetworkCacheMetaData();
}

void* QNetworkConfiguration_NewQNetworkConfiguration()
{
	return new QNetworkConfiguration();
}

void* QNetworkConfiguration_NewQNetworkConfiguration2(void* other)
{
	return new QNetworkConfiguration(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_BearerType(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerType();
}

int QNetworkConfiguration_BearerTypeFamily(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeFamily();
}

char* QNetworkConfiguration_BearerTypeName(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeName().toUtf8().data();
}

char* QNetworkConfiguration_Identifier(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->identifier().toUtf8().data();
}

int QNetworkConfiguration_IsRoamingAvailable(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->isRoamingAvailable();
}

int QNetworkConfiguration_IsValid(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->isValid();
}

char* QNetworkConfiguration_Name(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->name().toUtf8().data();
}

int QNetworkConfiguration_Purpose(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->purpose();
}

int QNetworkConfiguration_State(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->state();
}

void QNetworkConfiguration_Swap(void* ptr, void* other)
{
	static_cast<QNetworkConfiguration*>(ptr)->swap(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_Type(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->type();
}

void QNetworkConfiguration_DestroyQNetworkConfiguration(void* ptr)
{
	static_cast<QNetworkConfiguration*>(ptr)->~QNetworkConfiguration();
}

class MyQNetworkConfigurationManager: public QNetworkConfigurationManager
{
public:
	MyQNetworkConfigurationManager(QObject *parent) : QNetworkConfigurationManager(parent) {};
	void Signal_ConfigurationAdded(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationAdded(this, this->objectName().toUtf8().data(), new QNetworkConfiguration(config)); };
	void Signal_ConfigurationChanged(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationChanged(this, this->objectName().toUtf8().data(), new QNetworkConfiguration(config)); };
	void Signal_ConfigurationRemoved(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationRemoved(this, this->objectName().toUtf8().data(), new QNetworkConfiguration(config)); };
	void Signal_OnlineStateChanged(bool isOnline) { callbackQNetworkConfigurationManager_OnlineStateChanged(this, this->objectName().toUtf8().data(), isOnline); };
	void Signal_UpdateCompleted() { callbackQNetworkConfigurationManager_UpdateCompleted(this, this->objectName().toUtf8().data()); };
	void updateConfigurations() { callbackQNetworkConfigurationManager_UpdateConfigurations(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkConfigurationManager_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkConfigurationManager_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkConfigurationManager_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQNetworkConfigurationManager_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQNetworkConfigurationManager_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkConfigurationManager_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQNetworkConfigurationManager_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkConfigurationManager_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkConfigurationManager_MetaObject(const_cast<MyQNetworkConfigurationManager*>(this), this->objectName().toUtf8().data())); };
};

void* QNetworkConfigurationManager_NewQNetworkConfigurationManager(void* parent)
{
	return new MyQNetworkConfigurationManager(static_cast<QObject*>(parent));
}

int QNetworkConfigurationManager_Capabilities(void* ptr)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->capabilities();
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

void* QNetworkConfigurationManager_ConfigurationFromIdentifier(void* ptr, char* identifier)
{
	return new QNetworkConfiguration(static_cast<QNetworkConfigurationManager*>(ptr)->configurationFromIdentifier(QString(identifier)));
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

void* QNetworkConfigurationManager_DefaultConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkConfigurationManager*>(ptr)->defaultConfiguration());
}

int QNetworkConfigurationManager_IsOnline(void* ptr)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->isOnline();
}

void QNetworkConfigurationManager_ConnectOnlineStateChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));
}

void QNetworkConfigurationManager_DisconnectOnlineStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));
}

void QNetworkConfigurationManager_OnlineStateChanged(void* ptr, int isOnline)
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

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(void* ptr)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->~QNetworkConfigurationManager();
}

void QNetworkConfigurationManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkConfigurationManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkConfigurationManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkConfigurationManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkConfigurationManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkConfigurationManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkConfigurationManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkConfigurationManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::customEvent(static_cast<QEvent*>(event));
}

void QNetworkConfigurationManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkConfigurationManager*>(ptr), "deleteLater");
}

void QNetworkConfigurationManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::deleteLater();
}

void QNetworkConfigurationManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkConfigurationManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QNetworkConfigurationManager_Event(void* ptr, void* e)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->event(static_cast<QEvent*>(e));
}

int QNetworkConfigurationManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::event(static_cast<QEvent*>(e));
}

int QNetworkConfigurationManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QNetworkConfigurationManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNetworkConfigurationManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkConfigurationManager*>(ptr)->metaObject());
}

void* QNetworkConfigurationManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::metaObject());
}

void* QNetworkCookie_NewQNetworkCookie(char* name, char* value)
{
	return new QNetworkCookie(QByteArray(name), QByteArray(value));
}

void* QNetworkCookie_NewQNetworkCookie2(void* other)
{
	return new QNetworkCookie(*static_cast<QNetworkCookie*>(other));
}

char* QNetworkCookie_Domain(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->domain().toUtf8().data();
}

void* QNetworkCookie_ExpirationDate(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCookie*>(ptr)->expirationDate());
}

int QNetworkCookie_HasSameIdentifier(void* ptr, void* other)
{
	return static_cast<QNetworkCookie*>(ptr)->hasSameIdentifier(*static_cast<QNetworkCookie*>(other));
}

int QNetworkCookie_IsHttpOnly(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->isHttpOnly();
}

int QNetworkCookie_IsSecure(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->isSecure();
}

int QNetworkCookie_IsSessionCookie(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->isSessionCookie();
}

char* QNetworkCookie_Name(void* ptr)
{
	return QString(static_cast<QNetworkCookie*>(ptr)->name()).toUtf8().data();
}

void QNetworkCookie_Normalize(void* ptr, void* url)
{
	static_cast<QNetworkCookie*>(ptr)->normalize(*static_cast<QUrl*>(url));
}

char* QNetworkCookie_Path(void* ptr)
{
	return static_cast<QNetworkCookie*>(ptr)->path().toUtf8().data();
}

void QNetworkCookie_SetDomain(void* ptr, char* domain)
{
	static_cast<QNetworkCookie*>(ptr)->setDomain(QString(domain));
}

void QNetworkCookie_SetExpirationDate(void* ptr, void* date)
{
	static_cast<QNetworkCookie*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(date));
}

void QNetworkCookie_SetHttpOnly(void* ptr, int enable)
{
	static_cast<QNetworkCookie*>(ptr)->setHttpOnly(enable != 0);
}

void QNetworkCookie_SetName(void* ptr, char* cookieName)
{
	static_cast<QNetworkCookie*>(ptr)->setName(QByteArray(cookieName));
}

void QNetworkCookie_SetPath(void* ptr, char* path)
{
	static_cast<QNetworkCookie*>(ptr)->setPath(QString(path));
}

void QNetworkCookie_SetSecure(void* ptr, int enable)
{
	static_cast<QNetworkCookie*>(ptr)->setSecure(enable != 0);
}

void QNetworkCookie_SetValue(void* ptr, char* value)
{
	static_cast<QNetworkCookie*>(ptr)->setValue(QByteArray(value));
}

void QNetworkCookie_Swap(void* ptr, void* other)
{
	static_cast<QNetworkCookie*>(ptr)->swap(*static_cast<QNetworkCookie*>(other));
}

char* QNetworkCookie_ToRawForm(void* ptr, int form)
{
	return QString(static_cast<QNetworkCookie*>(ptr)->toRawForm(static_cast<QNetworkCookie::RawForm>(form))).toUtf8().data();
}

char* QNetworkCookie_Value(void* ptr)
{
	return QString(static_cast<QNetworkCookie*>(ptr)->value()).toUtf8().data();
}

void QNetworkCookie_DestroyQNetworkCookie(void* ptr)
{
	static_cast<QNetworkCookie*>(ptr)->~QNetworkCookie();
}

class MyQNetworkCookieJar: public QNetworkCookieJar
{
public:
	MyQNetworkCookieJar(QObject *parent) : QNetworkCookieJar(parent) {};
	bool deleteCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_DeleteCookie(this, this->objectName().toUtf8().data(), new QNetworkCookie(cookie)) != 0; };
	bool insertCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_InsertCookie(this, this->objectName().toUtf8().data(), new QNetworkCookie(cookie)) != 0; };
	bool updateCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_UpdateCookie(this, this->objectName().toUtf8().data(), new QNetworkCookie(cookie)) != 0; };
	bool validateCookie(const QNetworkCookie & cookie, const QUrl & url) const { return callbackQNetworkCookieJar_ValidateCookie(const_cast<MyQNetworkCookieJar*>(this), this->objectName().toUtf8().data(), new QNetworkCookie(cookie), new QUrl(url)) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQNetworkCookieJar_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkCookieJar_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkCookieJar_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQNetworkCookieJar_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQNetworkCookieJar_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkCookieJar_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQNetworkCookieJar_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkCookieJar_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkCookieJar_MetaObject(const_cast<MyQNetworkCookieJar*>(this), this->objectName().toUtf8().data())); };
};

void* QNetworkCookieJar_NewQNetworkCookieJar(void* parent)
{
	return new MyQNetworkCookieJar(static_cast<QObject*>(parent));
}

int QNetworkCookieJar_DeleteCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_DeleteCookieDefault(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_InsertCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_InsertCookieDefault(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_UpdateCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_UpdateCookieDefault(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_ValidateCookie(void* ptr, void* cookie, void* url)
{
	return static_cast<QNetworkCookieJar*>(ptr)->validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

int QNetworkCookieJar_ValidateCookieDefault(void* ptr, void* cookie, void* url)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

void QNetworkCookieJar_DestroyQNetworkCookieJar(void* ptr)
{
	static_cast<QNetworkCookieJar*>(ptr)->~QNetworkCookieJar();
}

void QNetworkCookieJar_TimerEvent(void* ptr, void* event)
{
	static_cast<QNetworkCookieJar*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkCookieJar_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkCookieJar_ChildEvent(void* ptr, void* event)
{
	static_cast<QNetworkCookieJar*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkCookieJar_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkCookieJar_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkCookieJar*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkCookieJar_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkCookieJar_CustomEvent(void* ptr, void* event)
{
	static_cast<QNetworkCookieJar*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkCookieJar_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::customEvent(static_cast<QEvent*>(event));
}

void QNetworkCookieJar_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkCookieJar*>(ptr), "deleteLater");
}

void QNetworkCookieJar_DeleteLaterDefault(void* ptr)
{
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::deleteLater();
}

void QNetworkCookieJar_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkCookieJar*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkCookieJar_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QNetworkCookieJar_Event(void* ptr, void* e)
{
	return static_cast<QNetworkCookieJar*>(ptr)->event(static_cast<QEvent*>(e));
}

int QNetworkCookieJar_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::event(static_cast<QEvent*>(e));
}

int QNetworkCookieJar_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkCookieJar*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QNetworkCookieJar_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNetworkCookieJar_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkCookieJar*>(ptr)->metaObject());
}

void* QNetworkCookieJar_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::metaObject());
}

class MyQNetworkDiskCache: public QNetworkDiskCache
{
public:
	MyQNetworkDiskCache(QObject *parent) : QNetworkDiskCache(parent) {};
	qint64 cacheSize() const { return static_cast<long long>(callbackQNetworkDiskCache_CacheSize(const_cast<MyQNetworkDiskCache*>(this), this->objectName().toUtf8().data())); };
	void clear() { callbackQNetworkDiskCache_Clear(this, this->objectName().toUtf8().data()); };
	QIODevice * data(const QUrl & url) { return static_cast<QIODevice*>(callbackQNetworkDiskCache_Data(this, this->objectName().toUtf8().data(), new QUrl(url))); };
	qint64 expire() { return static_cast<long long>(callbackQNetworkDiskCache_Expire(this, this->objectName().toUtf8().data())); };
	void insert(QIODevice * device) { callbackQNetworkDiskCache_Insert(this, this->objectName().toUtf8().data(), device); };
	QNetworkCacheMetaData metaData(const QUrl & url) { return *static_cast<QNetworkCacheMetaData*>(callbackQNetworkDiskCache_MetaData(this, this->objectName().toUtf8().data(), new QUrl(url))); };
	QIODevice * prepare(const QNetworkCacheMetaData & metaData) { return static_cast<QIODevice*>(callbackQNetworkDiskCache_Prepare(this, this->objectName().toUtf8().data(), new QNetworkCacheMetaData(metaData))); };
	bool remove(const QUrl & url) { return callbackQNetworkDiskCache_Remove(this, this->objectName().toUtf8().data(), new QUrl(url)) != 0; };
	void updateMetaData(const QNetworkCacheMetaData & metaData) { callbackQNetworkDiskCache_UpdateMetaData(this, this->objectName().toUtf8().data(), new QNetworkCacheMetaData(metaData)); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkDiskCache_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkDiskCache_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkDiskCache_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQNetworkDiskCache_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQNetworkDiskCache_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkDiskCache_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQNetworkDiskCache_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkDiskCache_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkDiskCache_MetaObject(const_cast<MyQNetworkDiskCache*>(this), this->objectName().toUtf8().data())); };
};

void* QNetworkDiskCache_NewQNetworkDiskCache(void* parent)
{
	return new MyQNetworkDiskCache(static_cast<QObject*>(parent));
}

char* QNetworkDiskCache_CacheDirectory(void* ptr)
{
	return static_cast<QNetworkDiskCache*>(ptr)->cacheDirectory().toUtf8().data();
}

long long QNetworkDiskCache_CacheSize(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->cacheSize());
}

long long QNetworkDiskCache_CacheSizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::cacheSize());
}

void QNetworkDiskCache_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkDiskCache*>(ptr), "clear");
}

void QNetworkDiskCache_ClearDefault(void* ptr)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::clear();
}

void* QNetworkDiskCache_Data(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void* QNetworkDiskCache_DataDefault(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::data(*static_cast<QUrl*>(url));
}

long long QNetworkDiskCache_Expire(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->expire());
}

long long QNetworkDiskCache_ExpireDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::expire());
}

void* QNetworkDiskCache_FileMetaData(void* ptr, char* fileName)
{
	return new QNetworkCacheMetaData(static_cast<QNetworkDiskCache*>(ptr)->fileMetaData(QString(fileName)));
}

void QNetworkDiskCache_Insert(void* ptr, void* device)
{
	static_cast<QNetworkDiskCache*>(ptr)->insert(static_cast<QIODevice*>(device));
}

void QNetworkDiskCache_InsertDefault(void* ptr, void* device)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::insert(static_cast<QIODevice*>(device));
}

long long QNetworkDiskCache_MaximumCacheSize(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->maximumCacheSize());
}

void* QNetworkDiskCache_MetaData(void* ptr, void* url)
{
	return new QNetworkCacheMetaData(static_cast<QNetworkDiskCache*>(ptr)->metaData(*static_cast<QUrl*>(url)));
}

void* QNetworkDiskCache_MetaDataDefault(void* ptr, void* url)
{
	return new QNetworkCacheMetaData(static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::metaData(*static_cast<QUrl*>(url)));
}

void* QNetworkDiskCache_Prepare(void* ptr, void* metaData)
{
	return static_cast<QNetworkDiskCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void* QNetworkDiskCache_PrepareDefault(void* ptr, void* metaData)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

int QNetworkDiskCache_Remove(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->remove(*static_cast<QUrl*>(url));
}

int QNetworkDiskCache_RemoveDefault(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::remove(*static_cast<QUrl*>(url));
}

void QNetworkDiskCache_SetCacheDirectory(void* ptr, char* cacheDir)
{
	static_cast<QNetworkDiskCache*>(ptr)->setCacheDirectory(QString(cacheDir));
}

void QNetworkDiskCache_SetMaximumCacheSize(void* ptr, long long size)
{
	static_cast<QNetworkDiskCache*>(ptr)->setMaximumCacheSize(static_cast<long long>(size));
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

void QNetworkDiskCache_TimerEvent(void* ptr, void* event)
{
	static_cast<QNetworkDiskCache*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkDiskCache_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkDiskCache_ChildEvent(void* ptr, void* event)
{
	static_cast<QNetworkDiskCache*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkDiskCache_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkDiskCache_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkDiskCache*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkDiskCache_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkDiskCache_CustomEvent(void* ptr, void* event)
{
	static_cast<QNetworkDiskCache*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkDiskCache_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::customEvent(static_cast<QEvent*>(event));
}

void QNetworkDiskCache_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkDiskCache*>(ptr), "deleteLater");
}

void QNetworkDiskCache_DeleteLaterDefault(void* ptr)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::deleteLater();
}

void QNetworkDiskCache_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkDiskCache*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkDiskCache_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QNetworkDiskCache_Event(void* ptr, void* e)
{
	return static_cast<QNetworkDiskCache*>(ptr)->event(static_cast<QEvent*>(e));
}

int QNetworkDiskCache_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::event(static_cast<QEvent*>(e));
}

int QNetworkDiskCache_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkDiskCache*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QNetworkDiskCache_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNetworkDiskCache_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkDiskCache*>(ptr)->metaObject());
}

void* QNetworkDiskCache_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::metaObject());
}

void* QNetworkInterface_NewQNetworkInterface()
{
	return new QNetworkInterface();
}

void* QNetworkInterface_NewQNetworkInterface2(void* other)
{
	return new QNetworkInterface(*static_cast<QNetworkInterface*>(other));
}

int QNetworkInterface_Flags(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->flags();
}

char* QNetworkInterface_HardwareAddress(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->hardwareAddress().toUtf8().data();
}

char* QNetworkInterface_HumanReadableName(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->humanReadableName().toUtf8().data();
}

int QNetworkInterface_Index(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->index();
}

void* QNetworkInterface_QNetworkInterface_InterfaceFromIndex(int index)
{
	return new QNetworkInterface(QNetworkInterface::interfaceFromIndex(index));
}

void* QNetworkInterface_QNetworkInterface_InterfaceFromName(char* name)
{
	return new QNetworkInterface(QNetworkInterface::interfaceFromName(QString(name)));
}

int QNetworkInterface_IsValid(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->isValid();
}

char* QNetworkInterface_Name(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->name().toUtf8().data();
}

void QNetworkInterface_Swap(void* ptr, void* other)
{
	static_cast<QNetworkInterface*>(ptr)->swap(*static_cast<QNetworkInterface*>(other));
}

void QNetworkInterface_DestroyQNetworkInterface(void* ptr)
{
	static_cast<QNetworkInterface*>(ptr)->~QNetworkInterface();
}

void* QNetworkProxy_NewQNetworkProxy()
{
	return new QNetworkProxy();
}

void* QNetworkProxy_NewQNetworkProxy3(void* other)
{
	return new QNetworkProxy(*static_cast<QNetworkProxy*>(other));
}

void* QNetworkProxy_QNetworkProxy_ApplicationProxy()
{
	return new QNetworkProxy(QNetworkProxy::applicationProxy());
}

int QNetworkProxy_Capabilities(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->capabilities();
}

int QNetworkProxy_HasRawHeader(void* ptr, char* headerName)
{
	return static_cast<QNetworkProxy*>(ptr)->hasRawHeader(QByteArray(headerName));
}

void* QNetworkProxy_Header(void* ptr, int header)
{
	return new QVariant(static_cast<QNetworkProxy*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

char* QNetworkProxy_HostName(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->hostName().toUtf8().data();
}

int QNetworkProxy_IsCachingProxy(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->isCachingProxy();
}

int QNetworkProxy_IsTransparentProxy(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->isTransparentProxy();
}

char* QNetworkProxy_Password(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->password().toUtf8().data();
}

char* QNetworkProxy_RawHeader(void* ptr, char* headerName)
{
	return QString(static_cast<QNetworkProxy*>(ptr)->rawHeader(QByteArray(headerName))).toUtf8().data();
}

void QNetworkProxy_QNetworkProxy_SetApplicationProxy(void* networkProxy)
{
	QNetworkProxy::setApplicationProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QNetworkProxy_SetCapabilities(void* ptr, int capabilities)
{
	static_cast<QNetworkProxy*>(ptr)->setCapabilities(static_cast<QNetworkProxy::Capability>(capabilities));
}

void QNetworkProxy_SetHeader(void* ptr, int header, void* value)
{
	static_cast<QNetworkProxy*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkProxy_SetHostName(void* ptr, char* hostName)
{
	static_cast<QNetworkProxy*>(ptr)->setHostName(QString(hostName));
}

void QNetworkProxy_SetPassword(void* ptr, char* password)
{
	static_cast<QNetworkProxy*>(ptr)->setPassword(QString(password));
}

void QNetworkProxy_SetRawHeader(void* ptr, char* headerName, char* headerValue)
{
	static_cast<QNetworkProxy*>(ptr)->setRawHeader(QByteArray(headerName), QByteArray(headerValue));
}

void QNetworkProxy_SetType(void* ptr, int ty)
{
	static_cast<QNetworkProxy*>(ptr)->setType(static_cast<QNetworkProxy::ProxyType>(ty));
}

void QNetworkProxy_SetUser(void* ptr, char* user)
{
	static_cast<QNetworkProxy*>(ptr)->setUser(QString(user));
}

void QNetworkProxy_Swap(void* ptr, void* other)
{
	static_cast<QNetworkProxy*>(ptr)->swap(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Type(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->type();
}

char* QNetworkProxy_User(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->user().toUtf8().data();
}

void QNetworkProxy_DestroyQNetworkProxy(void* ptr)
{
	static_cast<QNetworkProxy*>(ptr)->~QNetworkProxy();
}

class MyQNetworkProxyFactory: public QNetworkProxyFactory
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(void* factory)
{
	QNetworkProxyFactory::setApplicationProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

void QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(int enable)
{
	QNetworkProxyFactory::setUseSystemConfiguration(enable != 0);
}

void QNetworkProxyFactory_DestroyQNetworkProxyFactory(void* ptr)
{
	static_cast<QNetworkProxyFactory*>(ptr)->~QNetworkProxyFactory();
}

char* QNetworkProxyFactory_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQNetworkProxyFactory*>(static_cast<QNetworkProxyFactory*>(ptr))) {
		return static_cast<MyQNetworkProxyFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QNetworkProxyFactory_BASE").toUtf8().data();
}

void QNetworkProxyFactory_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQNetworkProxyFactory*>(static_cast<QNetworkProxyFactory*>(ptr))) {
		static_cast<MyQNetworkProxyFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery()
{
	return new QNetworkProxyQuery();
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery6(void* networkConfiguration, char* hostname, int port, char* protocolTag, int queryType)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery5(void* networkConfiguration, void* requestUrl, int queryType)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), *static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery8(void* other)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkProxyQuery*>(other));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, int queryType)
{
	return new QNetworkProxyQuery(QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery2(void* requestUrl, int queryType)
{
	return new QNetworkProxyQuery(*static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

int QNetworkProxyQuery_LocalPort(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->localPort();
}

void* QNetworkProxyQuery_NetworkConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkProxyQuery*>(ptr)->networkConfiguration());
}

char* QNetworkProxyQuery_PeerHostName(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->peerHostName().toUtf8().data();
}

int QNetworkProxyQuery_PeerPort(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->peerPort();
}

char* QNetworkProxyQuery_ProtocolTag(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->protocolTag().toUtf8().data();
}

int QNetworkProxyQuery_QueryType(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->queryType();
}

void QNetworkProxyQuery_SetLocalPort(void* ptr, int port)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setLocalPort(port);
}

void QNetworkProxyQuery_SetNetworkConfiguration(void* ptr, void* networkConfiguration)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setNetworkConfiguration(*static_cast<QNetworkConfiguration*>(networkConfiguration));
}

void QNetworkProxyQuery_SetPeerHostName(void* ptr, char* hostname)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerHostName(QString(hostname));
}

void QNetworkProxyQuery_SetPeerPort(void* ptr, int port)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerPort(port);
}

void QNetworkProxyQuery_SetProtocolTag(void* ptr, char* protocolTag)
{
	static_cast<QNetworkProxyQuery*>(ptr)->setProtocolTag(QString(protocolTag));
}

void QNetworkProxyQuery_SetQueryType(void* ptr, int ty)
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

void* QNetworkProxyQuery_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkProxyQuery*>(ptr)->url());
}

void QNetworkProxyQuery_DestroyQNetworkProxyQuery(void* ptr)
{
	static_cast<QNetworkProxyQuery*>(ptr)->~QNetworkProxyQuery();
}

class MyQNetworkReply: public QNetworkReply
{
public:
	void setSslConfigurationImplementation(const QSslConfiguration & configuration) { callbackQNetworkReply_SetSslConfigurationImplementation(this, this->objectName().toUtf8().data(), new QSslConfiguration(configuration)); };
	void sslConfigurationImplementation(QSslConfiguration & configuration) const { callbackQNetworkReply_SslConfigurationImplementation(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data(), new QSslConfiguration(configuration)); };
	void abort() { callbackQNetworkReply_Abort(this, this->objectName().toUtf8().data()); };
	void close() { callbackQNetworkReply_Close(this, this->objectName().toUtf8().data()); };
	void Signal_DownloadProgress(qint64 bytesReceived, qint64 bytesTotal) { callbackQNetworkReply_DownloadProgress(this, this->objectName().toUtf8().data(), static_cast<long long>(bytesReceived), static_cast<long long>(bytesTotal)); };
	void Signal_Encrypted() { callbackQNetworkReply_Encrypted(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QNetworkReply::NetworkError code) { callbackQNetworkReply_Error2(this, this->objectName().toUtf8().data(), code); };
	void Signal_Finished() { callbackQNetworkReply_Finished(this, this->objectName().toUtf8().data()); };
	void ignoreSslErrors() { callbackQNetworkReply_IgnoreSslErrors(this, this->objectName().toUtf8().data()); };
	void Signal_MetaDataChanged() { callbackQNetworkReply_MetaDataChanged(this, this->objectName().toUtf8().data()); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkReply_PreSharedKeyAuthenticationRequired(this, this->objectName().toUtf8().data(), authenticator); };
	void Signal_Redirected(const QUrl & url) { callbackQNetworkReply_Redirected(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void setReadBufferSize(qint64 size) { callbackQNetworkReply_SetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void Signal_UploadProgress(qint64 bytesSent, qint64 bytesTotal) { callbackQNetworkReply_UploadProgress(this, this->objectName().toUtf8().data(), static_cast<long long>(bytesSent), static_cast<long long>(bytesTotal)); };
	bool atEnd() const { return callbackQNetworkReply_AtEnd(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data()) != 0; };
	qint64 bytesAvailable() const { return static_cast<long long>(callbackQNetworkReply_BytesAvailable(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data())); };
	qint64 bytesToWrite() const { return static_cast<long long>(callbackQNetworkReply_BytesToWrite(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data())); };
	bool canReadLine() const { return callbackQNetworkReply_CanReadLine(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data()) != 0; };
	bool isSequential() const { return callbackQNetworkReply_IsSequential(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data()) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQNetworkReply_Open(this, this->objectName().toUtf8().data(), mode) != 0; };
	qint64 pos() const { return static_cast<long long>(callbackQNetworkReply_Pos(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data())); };
	qint64 readLineData(char * data, qint64 maxSize) { return static_cast<long long>(callbackQNetworkReply_ReadLineData(this, this->objectName().toUtf8().data(), QString(data).toUtf8().data(), static_cast<long long>(maxSize))); };
	bool reset() { return callbackQNetworkReply_Reset(this, this->objectName().toUtf8().data()) != 0; };
	bool seek(qint64 pos) { return callbackQNetworkReply_Seek(this, this->objectName().toUtf8().data(), static_cast<long long>(pos)) != 0; };
	qint64 size() const { return static_cast<long long>(callbackQNetworkReply_Size(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data())); };
	bool waitForBytesWritten(int msecs) { return callbackQNetworkReply_WaitForBytesWritten(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQNetworkReply_WaitForReadyRead(this, this->objectName().toUtf8().data(), msecs) != 0; };
	qint64 writeData(const char * data, qint64 maxSize) { return static_cast<long long>(callbackQNetworkReply_WriteData(this, this->objectName().toUtf8().data(), QString(data).toUtf8().data(), static_cast<long long>(maxSize))); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkReply_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkReply_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkReply_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQNetworkReply_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQNetworkReply_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkReply_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQNetworkReply_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkReply_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkReply_MetaObject(const_cast<MyQNetworkReply*>(this), this->objectName().toUtf8().data())); };
};

void QNetworkReply_SetSslConfigurationImplementation(void* ptr, void* configuration)
{
	static_cast<QNetworkReply*>(ptr)->setSslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_SetSslConfigurationImplementationDefault(void* ptr, void* configuration)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::setSslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_SslConfigurationImplementation(void* ptr, void* configuration)
{
	static_cast<QNetworkReply*>(ptr)->sslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_SslConfigurationImplementationDefault(void* ptr, void* configuration)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::sslConfigurationImplementation(*static_cast<QSslConfiguration*>(configuration));
}

void QNetworkReply_Abort(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "abort");
}

void* QNetworkReply_Attribute(void* ptr, int code)
{
	return new QVariant(static_cast<QNetworkReply*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code)));
}

void QNetworkReply_Close(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->close();
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
	static_cast<QNetworkReply*>(ptr)->downloadProgress(static_cast<long long>(bytesReceived), static_cast<long long>(bytesTotal));
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
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QNetworkReply::NetworkError)>(&QNetworkReply::error), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QNetworkReply::NetworkError)>(&MyQNetworkReply::Signal_Error2));
}

void QNetworkReply_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QNetworkReply::NetworkError)>(&QNetworkReply::error), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QNetworkReply::NetworkError)>(&MyQNetworkReply::Signal_Error2));
}

void QNetworkReply_Error2(void* ptr, int code)
{
	static_cast<QNetworkReply*>(ptr)->error(static_cast<QNetworkReply::NetworkError>(code));
}

int QNetworkReply_Error(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->error();
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

int QNetworkReply_HasRawHeader(void* ptr, char* headerName)
{
	return static_cast<QNetworkReply*>(ptr)->hasRawHeader(QByteArray(headerName));
}

void* QNetworkReply_Header(void* ptr, int header)
{
	return new QVariant(static_cast<QNetworkReply*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

void QNetworkReply_IgnoreSslErrors(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "ignoreSslErrors");
}

void QNetworkReply_IgnoreSslErrorsDefault(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::ignoreSslErrors();
}

int QNetworkReply_IsFinished(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isFinished();
}

int QNetworkReply_IsRunning(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isRunning();
}

void* QNetworkReply_Manager(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->manager();
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

int QNetworkReply_Operation(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->operation();
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

char* QNetworkReply_RawHeader(void* ptr, char* headerName)
{
	return QString(static_cast<QNetworkReply*>(ptr)->rawHeader(QByteArray(headerName))).toUtf8().data();
}

long long QNetworkReply_ReadBufferSize(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->readBufferSize());
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

void* QNetworkReply_Request(void* ptr)
{
	return new QNetworkRequest(static_cast<QNetworkReply*>(ptr)->request());
}

void QNetworkReply_SetAttribute(void* ptr, int code, void* value)
{
	static_cast<QNetworkReply*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkReply_SetError(void* ptr, int errorCode, char* errorString)
{
	static_cast<QNetworkReply*>(ptr)->setError(static_cast<QNetworkReply::NetworkError>(errorCode), QString(errorString));
}

void QNetworkReply_SetFinished(void* ptr, int finished)
{
	static_cast<QNetworkReply*>(ptr)->setFinished(finished != 0);
}

void QNetworkReply_SetHeader(void* ptr, int header, void* value)
{
	static_cast<QNetworkReply*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkReply_SetOperation(void* ptr, int operation)
{
	static_cast<QNetworkReply*>(ptr)->setOperation(static_cast<QNetworkAccessManager::Operation>(operation));
}

void QNetworkReply_SetRawHeader(void* ptr, char* headerName, char* value)
{
	static_cast<QNetworkReply*>(ptr)->setRawHeader(QByteArray(headerName), QByteArray(value));
}

void QNetworkReply_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QNetworkReply*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QNetworkReply_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::setReadBufferSize(static_cast<long long>(size));
}

void QNetworkReply_SetRequest(void* ptr, void* request)
{
	static_cast<QNetworkReply*>(ptr)->setRequest(*static_cast<QNetworkRequest*>(request));
}

void QNetworkReply_SetSslConfiguration(void* ptr, void* config)
{
	static_cast<QNetworkReply*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkReply_SetUrl(void* ptr, void* url)
{
	static_cast<QNetworkReply*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void* QNetworkReply_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QNetworkReply*>(ptr)->sslConfiguration());
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
	static_cast<QNetworkReply*>(ptr)->uploadProgress(static_cast<long long>(bytesSent), static_cast<long long>(bytesTotal));
}

void* QNetworkReply_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkReply*>(ptr)->url());
}

void QNetworkReply_DestroyQNetworkReply(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->~QNetworkReply();
}

int QNetworkReply_AtEnd(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->atEnd();
}

int QNetworkReply_AtEndDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::atEnd();
}

long long QNetworkReply_BytesAvailable(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->bytesAvailable());
}

long long QNetworkReply_BytesAvailableDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::bytesAvailable());
}

long long QNetworkReply_BytesToWrite(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->bytesToWrite());
}

long long QNetworkReply_BytesToWriteDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::bytesToWrite());
}

int QNetworkReply_CanReadLine(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->canReadLine();
}

int QNetworkReply_CanReadLineDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::canReadLine();
}

int QNetworkReply_IsSequential(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isSequential();
}

int QNetworkReply_IsSequentialDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::isSequential();
}

int QNetworkReply_Open(void* ptr, int mode)
{
	return static_cast<QNetworkReply*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QNetworkReply_OpenDefault(void* ptr, int mode)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QNetworkReply_Pos(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->pos());
}

long long QNetworkReply_PosDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::pos());
}

long long QNetworkReply_ReadLineData(void* ptr, char* data, long long maxSize)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->readLineData(data, static_cast<long long>(maxSize)));
}

long long QNetworkReply_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::readLineData(data, static_cast<long long>(maxSize)));
}

int QNetworkReply_Reset(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->reset();
}

int QNetworkReply_ResetDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::reset();
}

int QNetworkReply_Seek(void* ptr, long long pos)
{
	return static_cast<QNetworkReply*>(ptr)->seek(static_cast<long long>(pos));
}

int QNetworkReply_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::seek(static_cast<long long>(pos));
}

long long QNetworkReply_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->size());
}

long long QNetworkReply_SizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::size());
}

int QNetworkReply_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->waitForBytesWritten(msecs);
}

int QNetworkReply_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::waitForBytesWritten(msecs);
}

int QNetworkReply_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->waitForReadyRead(msecs);
}

int QNetworkReply_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::waitForReadyRead(msecs);
}

long long QNetworkReply_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(maxSize)));
}

long long QNetworkReply_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::writeData(const_cast<const char*>(data), static_cast<long long>(maxSize)));
}

void QNetworkReply_TimerEvent(void* ptr, void* event)
{
	static_cast<QNetworkReply*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkReply_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkReply_ChildEvent(void* ptr, void* event)
{
	static_cast<QNetworkReply*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkReply_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkReply_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkReply*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkReply_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkReply_CustomEvent(void* ptr, void* event)
{
	static_cast<QNetworkReply*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkReply_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::customEvent(static_cast<QEvent*>(event));
}

void QNetworkReply_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "deleteLater");
}

void QNetworkReply_DeleteLaterDefault(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::deleteLater();
}

void QNetworkReply_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkReply*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QNetworkReply_Event(void* ptr, void* e)
{
	return static_cast<QNetworkReply*>(ptr)->event(static_cast<QEvent*>(e));
}

int QNetworkReply_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::event(static_cast<QEvent*>(e));
}

int QNetworkReply_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkReply*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QNetworkReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNetworkReply_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkReply*>(ptr)->metaObject());
}

void* QNetworkReply_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkReply*>(ptr)->QNetworkReply::metaObject());
}

void* QNetworkRequest_NewQNetworkRequest2(void* other)
{
	return new QNetworkRequest(*static_cast<QNetworkRequest*>(other));
}

void* QNetworkRequest_NewQNetworkRequest(void* url)
{
	return new QNetworkRequest(*static_cast<QUrl*>(url));
}

void* QNetworkRequest_Attribute(void* ptr, int code, void* defaultValue)
{
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

int QNetworkRequest_HasRawHeader(void* ptr, char* headerName)
{
	return static_cast<QNetworkRequest*>(ptr)->hasRawHeader(QByteArray(headerName));
}

void* QNetworkRequest_Header(void* ptr, int header)
{
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

int QNetworkRequest_MaximumRedirectsAllowed(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->maximumRedirectsAllowed();
}

void* QNetworkRequest_OriginatingObject(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->originatingObject();
}

int QNetworkRequest_Priority(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->priority();
}

char* QNetworkRequest_RawHeader(void* ptr, char* headerName)
{
	return QString(static_cast<QNetworkRequest*>(ptr)->rawHeader(QByteArray(headerName))).toUtf8().data();
}

void QNetworkRequest_SetAttribute(void* ptr, int code, void* value)
{
	static_cast<QNetworkRequest*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetHeader(void* ptr, int header, void* value)
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

void QNetworkRequest_SetPriority(void* ptr, int priority)
{
	static_cast<QNetworkRequest*>(ptr)->setPriority(static_cast<QNetworkRequest::Priority>(priority));
}

void QNetworkRequest_SetRawHeader(void* ptr, char* headerName, char* headerValue)
{
	static_cast<QNetworkRequest*>(ptr)->setRawHeader(QByteArray(headerName), QByteArray(headerValue));
}

void QNetworkRequest_SetSslConfiguration(void* ptr, void* config)
{
	static_cast<QNetworkRequest*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkRequest_SetUrl(void* ptr, void* url)
{
	static_cast<QNetworkRequest*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void* QNetworkRequest_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QNetworkRequest*>(ptr)->sslConfiguration());
}

void QNetworkRequest_Swap(void* ptr, void* other)
{
	static_cast<QNetworkRequest*>(ptr)->swap(*static_cast<QNetworkRequest*>(other));
}

void* QNetworkRequest_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkRequest*>(ptr)->url());
}

void QNetworkRequest_DestroyQNetworkRequest(void* ptr)
{
	static_cast<QNetworkRequest*>(ptr)->~QNetworkRequest();
}

class MyQNetworkSession: public QNetworkSession
{
public:
	MyQNetworkSession(const QNetworkConfiguration &connectionConfig, QObject *parent) : QNetworkSession(connectionConfig, parent) {};
	void accept() { callbackQNetworkSession_Accept(this, this->objectName().toUtf8().data()); };
	void close() { callbackQNetworkSession_Close(this, this->objectName().toUtf8().data()); };
	void Signal_Closed() { callbackQNetworkSession_Closed(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QNetworkSession::SessionError error) { callbackQNetworkSession_Error2(this, this->objectName().toUtf8().data(), error); };
	void ignore() { callbackQNetworkSession_Ignore(this, this->objectName().toUtf8().data()); };
	void migrate() { callbackQNetworkSession_Migrate(this, this->objectName().toUtf8().data()); };
	void Signal_NewConfigurationActivated() { callbackQNetworkSession_NewConfigurationActivated(this, this->objectName().toUtf8().data()); };
	void open() { callbackQNetworkSession_Open(this, this->objectName().toUtf8().data()); };
	void Signal_Opened() { callbackQNetworkSession_Opened(this, this->objectName().toUtf8().data()); };
	void Signal_PreferredConfigurationChanged(const QNetworkConfiguration & config, bool isSeamless) { callbackQNetworkSession_PreferredConfigurationChanged(this, this->objectName().toUtf8().data(), new QNetworkConfiguration(config), isSeamless); };
	void reject() { callbackQNetworkSession_Reject(this, this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QNetworkSession::State state) { callbackQNetworkSession_StateChanged(this, this->objectName().toUtf8().data(), state); };
	void stop() { callbackQNetworkSession_Stop(this, this->objectName().toUtf8().data()); };
	void Signal_UsagePoliciesChanged(QNetworkSession::UsagePolicies usagePolicies) { callbackQNetworkSession_UsagePoliciesChanged(this, this->objectName().toUtf8().data(), usagePolicies); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkSession_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkSession_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkSession_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQNetworkSession_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQNetworkSession_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkSession_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQNetworkSession_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkSession_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkSession_MetaObject(const_cast<MyQNetworkSession*>(this), this->objectName().toUtf8().data())); };
};

void* QNetworkSession_NewQNetworkSession(void* connectionConfig, void* parent)
{
	return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QObject*>(parent));
}

void QNetworkSession_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "accept");
}

void QNetworkSession_Close(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "close");
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

void* QNetworkSession_Configuration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkSession*>(ptr)->configuration());
}

void QNetworkSession_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::SessionError)>(&QNetworkSession::error), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::SessionError)>(&MyQNetworkSession::Signal_Error2));
}

void QNetworkSession_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::SessionError)>(&QNetworkSession::error), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::SessionError)>(&MyQNetworkSession::Signal_Error2));
}

void QNetworkSession_Error2(void* ptr, int error)
{
	static_cast<QNetworkSession*>(ptr)->error(static_cast<QNetworkSession::SessionError>(error));
}

int QNetworkSession_Error(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->error();
}

char* QNetworkSession_ErrorString(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->errorString().toUtf8().data();
}

void QNetworkSession_Ignore(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "ignore");
}

void* QNetworkSession_Interface(void* ptr)
{
	return new QNetworkInterface(static_cast<QNetworkSession*>(ptr)->interface());
}

int QNetworkSession_IsOpen(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->isOpen();
}

void QNetworkSession_Migrate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "migrate");
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

void QNetworkSession_PreferredConfigurationChanged(void* ptr, void* config, int isSeamless)
{
	static_cast<QNetworkSession*>(ptr)->preferredConfigurationChanged(*static_cast<QNetworkConfiguration*>(config), isSeamless != 0);
}

void QNetworkSession_Reject(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "reject");
}

void* QNetworkSession_SessionProperty(void* ptr, char* key)
{
	return new QVariant(static_cast<QNetworkSession*>(ptr)->sessionProperty(QString(key)));
}

void QNetworkSession_SetSessionProperty(void* ptr, char* key, void* value)
{
	static_cast<QNetworkSession*>(ptr)->setSessionProperty(QString(key), *static_cast<QVariant*>(value));
}

int QNetworkSession_State(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->state();
}

void QNetworkSession_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));
}

void QNetworkSession_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));
}

void QNetworkSession_StateChanged(void* ptr, int state)
{
	static_cast<QNetworkSession*>(ptr)->stateChanged(static_cast<QNetworkSession::State>(state));
}

void QNetworkSession_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "stop");
}

int QNetworkSession_UsagePolicies(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->usagePolicies();
}

void QNetworkSession_ConnectUsagePoliciesChanged(void* ptr)
{
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));
}

void QNetworkSession_DisconnectUsagePoliciesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));
}

void QNetworkSession_UsagePoliciesChanged(void* ptr, int usagePolicies)
{
	static_cast<QNetworkSession*>(ptr)->usagePoliciesChanged(static_cast<QNetworkSession::UsagePolicy>(usagePolicies));
}

int QNetworkSession_WaitForOpened(void* ptr, int msecs)
{
	return static_cast<QNetworkSession*>(ptr)->waitForOpened(msecs);
}

void QNetworkSession_DestroyQNetworkSession(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->~QNetworkSession();
}

void QNetworkSession_TimerEvent(void* ptr, void* event)
{
	static_cast<QNetworkSession*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkSession_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkSession_ChildEvent(void* ptr, void* event)
{
	static_cast<QNetworkSession*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkSession_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkSession_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkSession*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkSession_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkSession_CustomEvent(void* ptr, void* event)
{
	static_cast<QNetworkSession*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkSession_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::customEvent(static_cast<QEvent*>(event));
}

void QNetworkSession_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "deleteLater");
}

void QNetworkSession_DeleteLaterDefault(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::deleteLater();
}

void QNetworkSession_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNetworkSession*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNetworkSession_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QNetworkSession_Event(void* ptr, void* e)
{
	return static_cast<QNetworkSession*>(ptr)->event(static_cast<QEvent*>(e));
}

int QNetworkSession_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkSession*>(ptr)->QNetworkSession::event(static_cast<QEvent*>(e));
}

int QNetworkSession_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkSession*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QNetworkSession_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkSession*>(ptr)->QNetworkSession::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNetworkSession_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkSession*>(ptr)->metaObject());
}

void* QNetworkSession_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNetworkSession*>(ptr)->QNetworkSession::metaObject());
}

void* QSslCertificate_NewQSslCertificate3(void* other)
{
	return new QSslCertificate(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_Clear(void* ptr)
{
	static_cast<QSslCertificate*>(ptr)->clear();
}

char* QSslCertificate_Digest(void* ptr, int algorithm)
{
	return QString(static_cast<QSslCertificate*>(ptr)->digest(static_cast<QCryptographicHash::Algorithm>(algorithm))).toUtf8().data();
}

int QSslCertificate_IsBlacklisted(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isBlacklisted();
}

void QSslCertificate_Swap(void* ptr, void* other)
{
	static_cast<QSslCertificate*>(ptr)->swap(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_DestroyQSslCertificate(void* ptr)
{
	static_cast<QSslCertificate*>(ptr)->~QSslCertificate();
}

void* QSslCertificate_EffectiveDate(void* ptr)
{
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->effectiveDate());
}

void* QSslCertificate_ExpiryDate(void* ptr)
{
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->expiryDate());
}

int QSslCertificate_IsNull(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isNull();
}

int QSslCertificate_IsSelfSigned(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isSelfSigned();
}

char* QSslCertificate_IssuerInfo(void* ptr, int subject)
{
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8().data();
}

char* QSslCertificate_IssuerInfo2(void* ptr, char* attribute)
{
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(QByteArray(attribute)).join("|").toUtf8().data();
}

void* QSslCertificate_PublicKey(void* ptr)
{
	return new QSslKey(static_cast<QSslCertificate*>(ptr)->publicKey());
}

char* QSslCertificate_SerialNumber(void* ptr)
{
	return QString(static_cast<QSslCertificate*>(ptr)->serialNumber()).toUtf8().data();
}

char* QSslCertificate_SubjectInfo(void* ptr, int subject)
{
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8().data();
}

char* QSslCertificate_SubjectInfo2(void* ptr, char* attribute)
{
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(QByteArray(attribute)).join("|").toUtf8().data();
}

char* QSslCertificate_ToDer(void* ptr)
{
	return QString(static_cast<QSslCertificate*>(ptr)->toDer()).toUtf8().data();
}

char* QSslCertificate_ToPem(void* ptr)
{
	return QString(static_cast<QSslCertificate*>(ptr)->toPem()).toUtf8().data();
}

char* QSslCertificate_ToText(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->toText().toUtf8().data();
}

char* QSslCertificate_Version(void* ptr)
{
	return QString(static_cast<QSslCertificate*>(ptr)->version()).toUtf8().data();
}

void* QSslCertificateExtension_NewQSslCertificateExtension()
{
	return new QSslCertificateExtension();
}

void* QSslCertificateExtension_NewQSslCertificateExtension2(void* other)
{
	return new QSslCertificateExtension(*static_cast<QSslCertificateExtension*>(other));
}

int QSslCertificateExtension_IsCritical(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->isCritical();
}

int QSslCertificateExtension_IsSupported(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->isSupported();
}

char* QSslCertificateExtension_Name(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->name().toUtf8().data();
}

char* QSslCertificateExtension_Oid(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->oid().toUtf8().data();
}

void QSslCertificateExtension_Swap(void* ptr, void* other)
{
	static_cast<QSslCertificateExtension*>(ptr)->swap(*static_cast<QSslCertificateExtension*>(other));
}

void* QSslCertificateExtension_Value(void* ptr)
{
	return new QVariant(static_cast<QSslCertificateExtension*>(ptr)->value());
}

void QSslCertificateExtension_DestroyQSslCertificateExtension(void* ptr)
{
	static_cast<QSslCertificateExtension*>(ptr)->~QSslCertificateExtension();
}

void* QSslCipher_NewQSslCipher()
{
	return new QSslCipher();
}

void* QSslCipher_NewQSslCipher4(void* other)
{
	return new QSslCipher(*static_cast<QSslCipher*>(other));
}

void* QSslCipher_NewQSslCipher2(char* name)
{
	return new QSslCipher(QString(name));
}

char* QSslCipher_AuthenticationMethod(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->authenticationMethod().toUtf8().data();
}

char* QSslCipher_EncryptionMethod(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->encryptionMethod().toUtf8().data();
}

int QSslCipher_IsNull(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->isNull();
}

char* QSslCipher_KeyExchangeMethod(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->keyExchangeMethod().toUtf8().data();
}

char* QSslCipher_Name(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->name().toUtf8().data();
}

char* QSslCipher_ProtocolString(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->protocolString().toUtf8().data();
}

int QSslCipher_SupportedBits(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->supportedBits();
}

void QSslCipher_Swap(void* ptr, void* other)
{
	static_cast<QSslCipher*>(ptr)->swap(*static_cast<QSslCipher*>(other));
}

int QSslCipher_UsedBits(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->usedBits();
}

void QSslCipher_DestroyQSslCipher(void* ptr)
{
	static_cast<QSslCipher*>(ptr)->~QSslCipher();
}

void* QSslConfiguration_NewQSslConfiguration()
{
	return new QSslConfiguration();
}

void* QSslConfiguration_NewQSslConfiguration2(void* other)
{
	return new QSslConfiguration(*static_cast<QSslConfiguration*>(other));
}

void* QSslConfiguration_QSslConfiguration_DefaultConfiguration()
{
	return new QSslConfiguration(QSslConfiguration::defaultConfiguration());
}

int QSslConfiguration_IsNull(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->isNull();
}

void* QSslConfiguration_LocalCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslConfiguration*>(ptr)->localCertificate());
}

char* QSslConfiguration_NextNegotiatedProtocol(void* ptr)
{
	return QString(static_cast<QSslConfiguration*>(ptr)->nextNegotiatedProtocol()).toUtf8().data();
}

int QSslConfiguration_NextProtocolNegotiationStatus(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->nextProtocolNegotiationStatus();
}

void* QSslConfiguration_PeerCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslConfiguration*>(ptr)->peerCertificate());
}

int QSslConfiguration_PeerVerifyDepth(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyDepth();
}

int QSslConfiguration_PeerVerifyMode(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyMode();
}

void* QSslConfiguration_PrivateKey(void* ptr)
{
	return new QSslKey(static_cast<QSslConfiguration*>(ptr)->privateKey());
}

void* QSslConfiguration_SessionCipher(void* ptr)
{
	return new QSslCipher(static_cast<QSslConfiguration*>(ptr)->sessionCipher());
}

char* QSslConfiguration_SessionTicket(void* ptr)
{
	return QString(static_cast<QSslConfiguration*>(ptr)->sessionTicket()).toUtf8().data();
}

int QSslConfiguration_SessionTicketLifeTimeHint(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->sessionTicketLifeTimeHint();
}

void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(void* configuration)
{
	QSslConfiguration::setDefaultConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void QSslConfiguration_SetLocalCertificate(void* ptr, void* certificate)
{
	static_cast<QSslConfiguration*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslConfiguration_SetPeerVerifyDepth(void* ptr, int depth)
{
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslConfiguration_SetPeerVerifyMode(void* ptr, int mode)
{
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslConfiguration_SetPrivateKey(void* ptr, void* key)
{
	static_cast<QSslConfiguration*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslConfiguration_SetSessionTicket(void* ptr, char* sessionTicket)
{
	static_cast<QSslConfiguration*>(ptr)->setSessionTicket(QByteArray(sessionTicket));
}

void QSslConfiguration_Swap(void* ptr, void* other)
{
	static_cast<QSslConfiguration*>(ptr)->swap(*static_cast<QSslConfiguration*>(other));
}

void QSslConfiguration_DestroyQSslConfiguration(void* ptr)
{
	static_cast<QSslConfiguration*>(ptr)->~QSslConfiguration();
}

void* QSslEllipticCurve_NewQSslEllipticCurve()
{
	return new QSslEllipticCurve();
}

int QSslEllipticCurve_IsValid(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->isValid();
}

int QSslEllipticCurve_IsTlsNamedCurve(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->isTlsNamedCurve();
}

char* QSslEllipticCurve_LongName(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->longName().toUtf8().data();
}

char* QSslEllipticCurve_ShortName(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->shortName().toUtf8().data();
}

void* QSslError_NewQSslError()
{
	return new QSslError();
}

void* QSslError_NewQSslError2(int error)
{
	return new QSslError(static_cast<QSslError::SslError>(error));
}

void* QSslError_NewQSslError3(int error, void* certificate)
{
	return new QSslError(static_cast<QSslError::SslError>(error), *static_cast<QSslCertificate*>(certificate));
}

void* QSslError_NewQSslError4(void* other)
{
	return new QSslError(*static_cast<QSslError*>(other));
}

void* QSslError_Certificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslError*>(ptr)->certificate());
}

int QSslError_Error(void* ptr)
{
	return static_cast<QSslError*>(ptr)->error();
}

char* QSslError_ErrorString(void* ptr)
{
	return static_cast<QSslError*>(ptr)->errorString().toUtf8().data();
}

void QSslError_Swap(void* ptr, void* other)
{
	static_cast<QSslError*>(ptr)->swap(*static_cast<QSslError*>(other));
}

void QSslError_DestroyQSslError(void* ptr)
{
	static_cast<QSslError*>(ptr)->~QSslError();
}

void* QSslKey_NewQSslKey()
{
	return new QSslKey();
}

void* QSslKey_NewQSslKey5(void* other)
{
	return new QSslKey(*static_cast<QSslKey*>(other));
}

void QSslKey_Clear(void* ptr)
{
	static_cast<QSslKey*>(ptr)->clear();
}

int QSslKey_IsNull(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->isNull();
}

int QSslKey_Length(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->length();
}

void QSslKey_Swap(void* ptr, void* other)
{
	static_cast<QSslKey*>(ptr)->swap(*static_cast<QSslKey*>(other));
}

char* QSslKey_ToDer(void* ptr, char* passPhrase)
{
	return QString(static_cast<QSslKey*>(ptr)->toDer(QByteArray(passPhrase))).toUtf8().data();
}

char* QSslKey_ToPem(void* ptr, char* passPhrase)
{
	return QString(static_cast<QSslKey*>(ptr)->toPem(QByteArray(passPhrase))).toUtf8().data();
}

void QSslKey_DestroyQSslKey(void* ptr)
{
	static_cast<QSslKey*>(ptr)->~QSslKey();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator()
{
	return new QSslPreSharedKeyAuthenticator();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(void* authenticator)
{
	return new QSslPreSharedKeyAuthenticator(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

char* QSslPreSharedKeyAuthenticator_Identity(void* ptr)
{
	return QString(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identity()).toUtf8().data();
}

char* QSslPreSharedKeyAuthenticator_IdentityHint(void* ptr)
{
	return QString(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identityHint()).toUtf8().data();
}

int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(void* ptr)
{
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumIdentityLength();
}

int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(void* ptr)
{
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumPreSharedKeyLength();
}

char* QSslPreSharedKeyAuthenticator_PreSharedKey(void* ptr)
{
	return QString(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->preSharedKey()).toUtf8().data();
}

void QSslPreSharedKeyAuthenticator_SetIdentity(void* ptr, char* identity)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setIdentity(QByteArray(identity));
}

void QSslPreSharedKeyAuthenticator_SetPreSharedKey(void* ptr, char* preSharedKey)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setPreSharedKey(QByteArray(preSharedKey));
}

void QSslPreSharedKeyAuthenticator_Swap(void* ptr, void* authenticator)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->swap(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(void* ptr)
{
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->~QSslPreSharedKeyAuthenticator();
}

class MyQSslSocket: public QSslSocket
{
public:
	MyQSslSocket(QObject *parent) : QSslSocket(parent) {};
	void Signal_Encrypted() { callbackQSslSocket_Encrypted(this, this->objectName().toUtf8().data()); };
	void Signal_EncryptedBytesWritten(qint64 written) { callbackQSslSocket_EncryptedBytesWritten(this, this->objectName().toUtf8().data(), static_cast<long long>(written)); };
	void ignoreSslErrors() { callbackQSslSocket_IgnoreSslErrors(this, this->objectName().toUtf8().data()); };
	void Signal_ModeChanged(QSslSocket::SslMode mode) { callbackQSslSocket_ModeChanged(this, this->objectName().toUtf8().data(), mode); };
	void Signal_PeerVerifyError(const QSslError & error) { callbackQSslSocket_PeerVerifyError(this, this->objectName().toUtf8().data(), new QSslError(error)); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQSslSocket_PreSharedKeyAuthenticationRequired(this, this->objectName().toUtf8().data(), authenticator); };
	void resume() { callbackQSslSocket_Resume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQSslSocket_SetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQSslSocket_SetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQSslSocket_SocketOption(this, this->objectName().toUtf8().data(), option)); };
	void startClientEncryption() { callbackQSslSocket_StartClientEncryption(this, this->objectName().toUtf8().data()); };
	void startServerEncryption() { callbackQSslSocket_StartServerEncryption(this, this->objectName().toUtf8().data()); };
	bool waitForConnected(int msecs) { return callbackQSslSocket_WaitForConnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQSslSocket_WaitForDisconnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	void disconnectFromHost() { callbackQSslSocket_DisconnectFromHost(this, this->objectName().toUtf8().data()); };
	bool open(QIODevice::OpenMode mode) { return callbackQSslSocket_Open(this, this->objectName().toUtf8().data(), mode) != 0; };
	qint64 pos() const { return static_cast<long long>(callbackQSslSocket_Pos(const_cast<MyQSslSocket*>(this), this->objectName().toUtf8().data())); };
	bool reset() { return callbackQSslSocket_Reset(this, this->objectName().toUtf8().data()) != 0; };
	bool seek(qint64 pos) { return callbackQSslSocket_Seek(this, this->objectName().toUtf8().data(), static_cast<long long>(pos)) != 0; };
	qint64 size() const { return static_cast<long long>(callbackQSslSocket_Size(const_cast<MyQSslSocket*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQSslSocket_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSslSocket_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSslSocket_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSslSocket_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSslSocket_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSslSocket_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSslSocket_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSslSocket_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSslSocket_MetaObject(const_cast<MyQSslSocket*>(this), this->objectName().toUtf8().data())); };
};

void* QSslSocket_NewQSslSocket(void* parent)
{
	return new MyQSslSocket(static_cast<QObject*>(parent));
}

void QSslSocket_Abort(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->abort();
}

void QSslSocket_AddCaCertificate(void* ptr, void* certificate)
{
	static_cast<QSslSocket*>(ptr)->addCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificate(void* certificate)
{
	QSslSocket::addDefaultCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

int QSslSocket_AtEnd(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->atEnd();
}

long long QSslSocket_BytesAvailable(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->bytesAvailable());
}

long long QSslSocket_BytesToWrite(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->bytesToWrite());
}

int QSslSocket_CanReadLine(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->canReadLine();
}

void QSslSocket_Close(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->close();
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

long long QSslSocket_EncryptedBytesAvailable(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->encryptedBytesAvailable());
}

long long QSslSocket_EncryptedBytesToWrite(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->encryptedBytesToWrite());
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
	static_cast<QSslSocket*>(ptr)->encryptedBytesWritten(static_cast<long long>(written));
}

int QSslSocket_Flush(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->flush();
}

void QSslSocket_IgnoreSslErrors(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "ignoreSslErrors");
}

int QSslSocket_IsEncrypted(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->isEncrypted();
}

void* QSslSocket_LocalCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslSocket*>(ptr)->localCertificate());
}

int QSslSocket_Mode(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->mode();
}

void QSslSocket_ConnectModeChanged(void* ptr)
{
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));
}

void QSslSocket_DisconnectModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));
}

void QSslSocket_ModeChanged(void* ptr, int mode)
{
	static_cast<QSslSocket*>(ptr)->modeChanged(static_cast<QSslSocket::SslMode>(mode));
}

void* QSslSocket_PeerCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslSocket*>(ptr)->peerCertificate());
}

int QSslSocket_PeerVerifyDepth(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->peerVerifyDepth();
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

int QSslSocket_PeerVerifyMode(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->peerVerifyMode();
}

char* QSslSocket_PeerVerifyName(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->peerVerifyName().toUtf8().data();
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

void* QSslSocket_PrivateKey(void* ptr)
{
	return new QSslKey(static_cast<QSslSocket*>(ptr)->privateKey());
}

void QSslSocket_Resume(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->resume();
}

void QSslSocket_ResumeDefault(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::resume();
}

void* QSslSocket_SessionCipher(void* ptr)
{
	return new QSslCipher(static_cast<QSslSocket*>(ptr)->sessionCipher());
}

void QSslSocket_SetLocalCertificate(void* ptr, void* certificate)
{
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_SetPeerVerifyDepth(void* ptr, int depth)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslSocket_SetPeerVerifyMode(void* ptr, int mode)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslSocket_SetPeerVerifyName(void* ptr, char* hostName)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyName(QString(hostName));
}

void QSslSocket_SetPrivateKey(void* ptr, void* key)
{
	static_cast<QSslSocket*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QSslSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QSslSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::setReadBufferSize(static_cast<long long>(size));
}

void QSslSocket_SetSocketOption(void* ptr, int option, void* value)
{
	static_cast<QSslSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSocketOptionDefault(void* ptr, int option, void* value)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSslConfiguration(void* ptr, void* configuration)
{
	static_cast<QSslSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void* QSslSocket_SocketOption(void* ptr, int option)
{
	return new QVariant(static_cast<QSslSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QSslSocket_SocketOptionDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QSslSocket*>(ptr)->QSslSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QSslSocket_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QSslSocket*>(ptr)->sslConfiguration());
}

long QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()
{
	return QSslSocket::sslLibraryBuildVersionNumber();
}

char* QSslSocket_QSslSocket_SslLibraryBuildVersionString()
{
	return QSslSocket::sslLibraryBuildVersionString().toUtf8().data();
}

long QSslSocket_QSslSocket_SslLibraryVersionNumber()
{
	return QSslSocket::sslLibraryVersionNumber();
}

char* QSslSocket_QSslSocket_SslLibraryVersionString()
{
	return QSslSocket::sslLibraryVersionString().toUtf8().data();
}

void QSslSocket_StartClientEncryption(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startClientEncryption");
}

void QSslSocket_StartServerEncryption(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startServerEncryption");
}

int QSslSocket_QSslSocket_SupportsSsl()
{
	return QSslSocket::supportsSsl();
}

int QSslSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QSslSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForConnected(msecs);
}

int QSslSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForConnected(msecs);
}

int QSslSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForDisconnected(msecs);
}

int QSslSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForDisconnected(msecs);
}

int QSslSocket_WaitForEncrypted(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForEncrypted(msecs);
}

int QSslSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForReadyRead(msecs);
}

long long QSslSocket_WriteData(void* ptr, char* data, long long len)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(len)));
}

void QSslSocket_DestroyQSslSocket(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->~QSslSocket();
}

void QSslSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->disconnectFromHost();
}

void QSslSocket_DisconnectFromHostDefault(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::disconnectFromHost();
}

int QSslSocket_Open(void* ptr, int mode)
{
	return static_cast<QSslSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QSslSocket_OpenDefault(void* ptr, int mode)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QSslSocket_Pos(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->pos());
}

long long QSslSocket_PosDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->QSslSocket::pos());
}

int QSslSocket_Reset(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->reset();
}

int QSslSocket_ResetDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::reset();
}

int QSslSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QSslSocket*>(ptr)->seek(static_cast<long long>(pos));
}

int QSslSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::seek(static_cast<long long>(pos));
}

long long QSslSocket_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->size());
}

long long QSslSocket_SizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->QSslSocket::size());
}

void QSslSocket_TimerEvent(void* ptr, void* event)
{
	static_cast<QSslSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSslSocket_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSslSocket_ChildEvent(void* ptr, void* event)
{
	static_cast<QSslSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSslSocket_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QSslSocket_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSslSocket*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSslSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSslSocket_CustomEvent(void* ptr, void* event)
{
	static_cast<QSslSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSslSocket_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::customEvent(static_cast<QEvent*>(event));
}

void QSslSocket_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "deleteLater");
}

void QSslSocket_DeleteLaterDefault(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::deleteLater();
}

void QSslSocket_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSslSocket*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSslSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSslSocket_Event(void* ptr, void* e)
{
	return static_cast<QSslSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSslSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::event(static_cast<QEvent*>(e));
}

int QSslSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSslSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSslSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSslSocket_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSslSocket*>(ptr)->metaObject());
}

void* QSslSocket_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSslSocket*>(ptr)->QSslSocket::metaObject());
}

class MyQTcpServer: public QTcpServer
{
public:
	MyQTcpServer(QObject *parent) : QTcpServer(parent) {};
	void Signal_AcceptError(QAbstractSocket::SocketError socketError) { callbackQTcpServer_AcceptError(this, this->objectName().toUtf8().data(), socketError); };
	bool hasPendingConnections() const { return callbackQTcpServer_HasPendingConnections(const_cast<MyQTcpServer*>(this), this->objectName().toUtf8().data()) != 0; };
	void Signal_NewConnection() { callbackQTcpServer_NewConnection(this, this->objectName().toUtf8().data()); };
	QTcpSocket * nextPendingConnection() { return static_cast<QTcpSocket*>(callbackQTcpServer_NextPendingConnection(this, this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQTcpServer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTcpServer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTcpServer_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQTcpServer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQTcpServer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTcpServer_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQTcpServer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTcpServer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTcpServer_MetaObject(const_cast<MyQTcpServer*>(this), this->objectName().toUtf8().data())); };
};

void* QTcpServer_NewQTcpServer(void* parent)
{
	return new MyQTcpServer(static_cast<QObject*>(parent));
}

void QTcpServer_ConnectAcceptError(void* ptr)
{
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));
}

void QTcpServer_DisconnectAcceptError(void* ptr)
{
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));
}

void QTcpServer_AcceptError(void* ptr, int socketError)
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

char* QTcpServer_ErrorString(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->errorString().toUtf8().data();
}

int QTcpServer_HasPendingConnections(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->hasPendingConnections();
}

int QTcpServer_HasPendingConnectionsDefault(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->QTcpServer::hasPendingConnections();
}

int QTcpServer_IsListening(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->isListening();
}

int QTcpServer_MaxPendingConnections(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->maxPendingConnections();
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

void* QTcpServer_NextPendingConnection(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->nextPendingConnection();
}

void* QTcpServer_NextPendingConnectionDefault(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->QTcpServer::nextPendingConnection();
}

void QTcpServer_PauseAccepting(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->pauseAccepting();
}

void* QTcpServer_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QTcpServer*>(ptr)->proxy());
}

void QTcpServer_ResumeAccepting(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->resumeAccepting();
}

void* QTcpServer_ServerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QTcpServer*>(ptr)->serverAddress());
}

int QTcpServer_ServerError(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->serverError();
}

void QTcpServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QTcpServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QTcpServer_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QTcpServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

int QTcpServer_WaitForNewConnection(void* ptr, int msec, int timedOut)
{
	return static_cast<QTcpServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QTcpServer_DestroyQTcpServer(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->~QTcpServer();
}

void QTcpServer_TimerEvent(void* ptr, void* event)
{
	static_cast<QTcpServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpServer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTcpServer*>(ptr)->QTcpServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpServer_ChildEvent(void* ptr, void* event)
{
	static_cast<QTcpServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTcpServer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTcpServer*>(ptr)->QTcpServer::childEvent(static_cast<QChildEvent*>(event));
}

void QTcpServer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTcpServer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpServer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTcpServer*>(ptr)->QTcpServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpServer_CustomEvent(void* ptr, void* event)
{
	static_cast<QTcpServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTcpServer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTcpServer*>(ptr)->QTcpServer::customEvent(static_cast<QEvent*>(event));
}

void QTcpServer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTcpServer*>(ptr), "deleteLater");
}

void QTcpServer_DeleteLaterDefault(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->QTcpServer::deleteLater();
}

void QTcpServer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTcpServer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTcpServer*>(ptr)->QTcpServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTcpServer_Event(void* ptr, void* e)
{
	return static_cast<QTcpServer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTcpServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QTcpServer*>(ptr)->QTcpServer::event(static_cast<QEvent*>(e));
}

int QTcpServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTcpServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTcpServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTcpServer*>(ptr)->QTcpServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTcpServer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTcpServer*>(ptr)->metaObject());
}

void* QTcpServer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTcpServer*>(ptr)->QTcpServer::metaObject());
}

class MyQTcpSocket: public QTcpSocket
{
public:
	MyQTcpSocket(QObject *parent) : QTcpSocket(parent) {};
	void disconnectFromHost() { callbackQTcpSocket_DisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void resume() { callbackQTcpSocket_Resume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQTcpSocket_SetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQTcpSocket_SetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQTcpSocket_SocketOption(this, this->objectName().toUtf8().data(), option)); };
	bool waitForConnected(int msecs) { return callbackQTcpSocket_WaitForConnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQTcpSocket_WaitForDisconnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQTcpSocket_Open(this, this->objectName().toUtf8().data(), mode) != 0; };
	qint64 pos() const { return static_cast<long long>(callbackQTcpSocket_Pos(const_cast<MyQTcpSocket*>(this), this->objectName().toUtf8().data())); };
	bool reset() { return callbackQTcpSocket_Reset(this, this->objectName().toUtf8().data()) != 0; };
	bool seek(qint64 pos) { return callbackQTcpSocket_Seek(this, this->objectName().toUtf8().data(), static_cast<long long>(pos)) != 0; };
	qint64 size() const { return static_cast<long long>(callbackQTcpSocket_Size(const_cast<MyQTcpSocket*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQTcpSocket_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTcpSocket_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTcpSocket_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQTcpSocket_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQTcpSocket_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTcpSocket_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQTcpSocket_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTcpSocket_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTcpSocket_MetaObject(const_cast<MyQTcpSocket*>(this), this->objectName().toUtf8().data())); };
};

void* QTcpSocket_NewQTcpSocket(void* parent)
{
	return new MyQTcpSocket(static_cast<QObject*>(parent));
}

void QTcpSocket_DestroyQTcpSocket(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->~QTcpSocket();
}

void QTcpSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->disconnectFromHost();
}

void QTcpSocket_DisconnectFromHostDefault(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::disconnectFromHost();
}

void QTcpSocket_Resume(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->resume();
}

void QTcpSocket_ResumeDefault(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::resume();
}

void QTcpSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QTcpSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QTcpSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::setReadBufferSize(static_cast<long long>(size));
}

void QTcpSocket_SetSocketOption(void* ptr, int option, void* value)
{
	static_cast<QTcpSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QTcpSocket_SetSocketOptionDefault(void* ptr, int option, void* value)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void* QTcpSocket_SocketOption(void* ptr, int option)
{
	return new QVariant(static_cast<QTcpSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QTcpSocket_SocketOptionDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QTcpSocket*>(ptr)->QTcpSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

int QTcpSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->waitForConnected(msecs);
}

int QTcpSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForConnected(msecs);
}

int QTcpSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->waitForDisconnected(msecs);
}

int QTcpSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForDisconnected(msecs);
}

int QTcpSocket_Open(void* ptr, int mode)
{
	return static_cast<QTcpSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QTcpSocket_OpenDefault(void* ptr, int mode)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QTcpSocket_Pos(void* ptr)
{
	return static_cast<long long>(static_cast<QTcpSocket*>(ptr)->pos());
}

long long QTcpSocket_PosDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QTcpSocket*>(ptr)->QTcpSocket::pos());
}

int QTcpSocket_Reset(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->reset();
}

int QTcpSocket_ResetDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::reset();
}

int QTcpSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QTcpSocket*>(ptr)->seek(static_cast<long long>(pos));
}

int QTcpSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::seek(static_cast<long long>(pos));
}

long long QTcpSocket_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QTcpSocket*>(ptr)->size());
}

long long QTcpSocket_SizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QTcpSocket*>(ptr)->QTcpSocket::size());
}

void QTcpSocket_TimerEvent(void* ptr, void* event)
{
	static_cast<QTcpSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpSocket_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpSocket_ChildEvent(void* ptr, void* event)
{
	static_cast<QTcpSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTcpSocket_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QTcpSocket_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTcpSocket*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpSocket_CustomEvent(void* ptr, void* event)
{
	static_cast<QTcpSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTcpSocket_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::customEvent(static_cast<QEvent*>(event));
}

void QTcpSocket_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTcpSocket*>(ptr), "deleteLater");
}

void QTcpSocket_DeleteLaterDefault(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::deleteLater();
}

void QTcpSocket_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTcpSocket*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTcpSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTcpSocket_Event(void* ptr, void* e)
{
	return static_cast<QTcpSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTcpSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::event(static_cast<QEvent*>(e));
}

int QTcpSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTcpSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTcpSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTcpSocket_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTcpSocket*>(ptr)->metaObject());
}

void* QTcpSocket_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTcpSocket*>(ptr)->QTcpSocket::metaObject());
}

class MyQUdpSocket: public QUdpSocket
{
public:
	MyQUdpSocket(QObject *parent) : QUdpSocket(parent) {};
	void disconnectFromHost() { callbackQUdpSocket_DisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void resume() { callbackQUdpSocket_Resume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQUdpSocket_SetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQUdpSocket_SetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQUdpSocket_SocketOption(this, this->objectName().toUtf8().data(), option)); };
	bool waitForConnected(int msecs) { return callbackQUdpSocket_WaitForConnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQUdpSocket_WaitForDisconnected(this, this->objectName().toUtf8().data(), msecs) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQUdpSocket_Open(this, this->objectName().toUtf8().data(), mode) != 0; };
	qint64 pos() const { return static_cast<long long>(callbackQUdpSocket_Pos(const_cast<MyQUdpSocket*>(this), this->objectName().toUtf8().data())); };
	bool reset() { return callbackQUdpSocket_Reset(this, this->objectName().toUtf8().data()) != 0; };
	bool seek(qint64 pos) { return callbackQUdpSocket_Seek(this, this->objectName().toUtf8().data(), static_cast<long long>(pos)) != 0; };
	qint64 size() const { return static_cast<long long>(callbackQUdpSocket_Size(const_cast<MyQUdpSocket*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQUdpSocket_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQUdpSocket_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQUdpSocket_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQUdpSocket_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQUdpSocket_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQUdpSocket_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQUdpSocket_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQUdpSocket_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQUdpSocket_MetaObject(const_cast<MyQUdpSocket*>(this), this->objectName().toUtf8().data())); };
};

void* QUdpSocket_NewQUdpSocket(void* parent)
{
	return new MyQUdpSocket(static_cast<QObject*>(parent));
}

int QUdpSocket_HasPendingDatagrams(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->hasPendingDatagrams();
}

int QUdpSocket_JoinMulticastGroup(void* ptr, void* groupAddress)
{
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_JoinMulticastGroup2(void* ptr, void* groupAddress, void* iface)
{
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

int QUdpSocket_LeaveMulticastGroup(void* ptr, void* groupAddress)
{
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_LeaveMulticastGroup2(void* ptr, void* groupAddress, void* iface)
{
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

void* QUdpSocket_MulticastInterface(void* ptr)
{
	return new QNetworkInterface(static_cast<QUdpSocket*>(ptr)->multicastInterface());
}

long long QUdpSocket_PendingDatagramSize(void* ptr)
{
	return static_cast<long long>(static_cast<QUdpSocket*>(ptr)->pendingDatagramSize());
}

void QUdpSocket_SetMulticastInterface(void* ptr, void* iface)
{
	static_cast<QUdpSocket*>(ptr)->setMulticastInterface(*static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_DestroyQUdpSocket(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->~QUdpSocket();
}

void QUdpSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->disconnectFromHost();
}

void QUdpSocket_DisconnectFromHostDefault(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::disconnectFromHost();
}

void QUdpSocket_Resume(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->resume();
}

void QUdpSocket_ResumeDefault(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::resume();
}

void QUdpSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QUdpSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QUdpSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::setReadBufferSize(static_cast<long long>(size));
}

void QUdpSocket_SetSocketOption(void* ptr, int option, void* value)
{
	static_cast<QUdpSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QUdpSocket_SetSocketOptionDefault(void* ptr, int option, void* value)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void* QUdpSocket_SocketOption(void* ptr, int option)
{
	return new QVariant(static_cast<QUdpSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QUdpSocket_SocketOptionDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QUdpSocket*>(ptr)->QUdpSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

int QUdpSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->waitForConnected(msecs);
}

int QUdpSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForConnected(msecs);
}

int QUdpSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->waitForDisconnected(msecs);
}

int QUdpSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForDisconnected(msecs);
}

int QUdpSocket_Open(void* ptr, int mode)
{
	return static_cast<QUdpSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QUdpSocket_OpenDefault(void* ptr, int mode)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QUdpSocket_Pos(void* ptr)
{
	return static_cast<long long>(static_cast<QUdpSocket*>(ptr)->pos());
}

long long QUdpSocket_PosDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QUdpSocket*>(ptr)->QUdpSocket::pos());
}

int QUdpSocket_Reset(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->reset();
}

int QUdpSocket_ResetDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::reset();
}

int QUdpSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QUdpSocket*>(ptr)->seek(static_cast<long long>(pos));
}

int QUdpSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::seek(static_cast<long long>(pos));
}

long long QUdpSocket_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QUdpSocket*>(ptr)->size());
}

long long QUdpSocket_SizeDefault(void* ptr)
{
	return static_cast<long long>(static_cast<QUdpSocket*>(ptr)->QUdpSocket::size());
}

void QUdpSocket_TimerEvent(void* ptr, void* event)
{
	static_cast<QUdpSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QUdpSocket_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QUdpSocket_ChildEvent(void* ptr, void* event)
{
	static_cast<QUdpSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QUdpSocket_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QUdpSocket_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QUdpSocket*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUdpSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUdpSocket_CustomEvent(void* ptr, void* event)
{
	static_cast<QUdpSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QUdpSocket_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::customEvent(static_cast<QEvent*>(event));
}

void QUdpSocket_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QUdpSocket*>(ptr), "deleteLater");
}

void QUdpSocket_DeleteLaterDefault(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::deleteLater();
}

void QUdpSocket_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QUdpSocket*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QUdpSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QUdpSocket_Event(void* ptr, void* e)
{
	return static_cast<QUdpSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

int QUdpSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::event(static_cast<QEvent*>(e));
}

int QUdpSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QUdpSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QUdpSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QUdpSocket_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QUdpSocket*>(ptr)->metaObject());
}

void* QUdpSocket_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QUdpSocket*>(ptr)->QUdpSocket::metaObject());
}

