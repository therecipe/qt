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
#include <QList>
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
#include <QRegExp>
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
#include <QVector>

class MyQAbstractNetworkCache: public QAbstractNetworkCache
{
public:
	MyQAbstractNetworkCache(QObject *parent) : QAbstractNetworkCache(parent) {};
	qint64 cacheSize() const { return callbackQAbstractNetworkCache_CacheSize(const_cast<MyQAbstractNetworkCache*>(this)); };
	void clear() { callbackQAbstractNetworkCache_Clear(this); };
	QIODevice * data(const QUrl & url) { return static_cast<QIODevice*>(callbackQAbstractNetworkCache_Data(this, const_cast<QUrl*>(&url))); };
	void insert(QIODevice * device) { callbackQAbstractNetworkCache_Insert(this, device); };
	QNetworkCacheMetaData metaData(const QUrl & url) { return *static_cast<QNetworkCacheMetaData*>(callbackQAbstractNetworkCache_MetaData(this, const_cast<QUrl*>(&url))); };
	QIODevice * prepare(const QNetworkCacheMetaData & metaData) { return static_cast<QIODevice*>(callbackQAbstractNetworkCache_Prepare(this, const_cast<QNetworkCacheMetaData*>(&metaData))); };
	bool remove(const QUrl & url) { return callbackQAbstractNetworkCache_Remove(this, const_cast<QUrl*>(&url)) != 0; };
	void updateMetaData(const QNetworkCacheMetaData & metaData) { callbackQAbstractNetworkCache_UpdateMetaData(this, const_cast<QNetworkCacheMetaData*>(&metaData)); };
	 ~MyQAbstractNetworkCache() { callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache(this); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractNetworkCache_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAbstractNetworkCache_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractNetworkCache_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractNetworkCache_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractNetworkCache_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAbstractNetworkCache_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractNetworkCache_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractNetworkCache_MetaObject(const_cast<MyQAbstractNetworkCache*>(this))); };
};

void* QAbstractNetworkCache_NewQAbstractNetworkCache(void* parent)
{
	return new MyQAbstractNetworkCache(static_cast<QObject*>(parent));
}

long long QAbstractNetworkCache_CacheSize(void* ptr)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->cacheSize();
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

char QAbstractNetworkCache_Remove(void* ptr, void* url)
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

void QAbstractNetworkCache_DestroyQAbstractNetworkCacheDefault(void* ptr)
{

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

char QAbstractNetworkCache_Event(void* ptr, void* e)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAbstractNetworkCache_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::event(static_cast<QEvent*>(e));
}

char QAbstractNetworkCache_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractNetworkCache*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAbstractNetworkCache_EventFilterDefault(void* ptr, void* watched, void* event)
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
	bool atEnd() const { return callbackQAbstractSocket_AtEnd(const_cast<MyQAbstractSocket*>(this)) != 0; };
	qint64 bytesAvailable() const { return callbackQAbstractSocket_BytesAvailable(const_cast<MyQAbstractSocket*>(this)); };
	qint64 bytesToWrite() const { return callbackQAbstractSocket_BytesToWrite(const_cast<MyQAbstractSocket*>(this)); };
	bool canReadLine() const { return callbackQAbstractSocket_CanReadLine(const_cast<MyQAbstractSocket*>(this)) != 0; };
	void close() { callbackQAbstractSocket_Close(this); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQAbstractSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQAbstractSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void Signal_Connected() { callbackQAbstractSocket_Connected(this); };
	void disconnectFromHost() { callbackQAbstractSocket_DisconnectFromHost(this); };
	void Signal_Disconnected() { callbackQAbstractSocket_Disconnected(this); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocket_Error2(this, socketError); };
	void Signal_HostFound() { callbackQAbstractSocket_HostFound(this); };
	bool isSequential() const { return callbackQAbstractSocket_IsSequential(const_cast<MyQAbstractSocket*>(this)) != 0; };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQAbstractSocket_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQAbstractSocket_ReadLineData(this, dataPacked, maxlen); };
	void resume() { callbackQAbstractSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocket_SetReadBufferSize(this, size); };
	
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQAbstractSocket_SocketOption(this, option)); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocket_StateChanged(this, socketState); };
	bool waitForBytesWritten(int msecs) { return callbackQAbstractSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQAbstractSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQAbstractSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQAbstractSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 size) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQAbstractSocket_WriteData(this, dataPacked, size); };
	 ~MyQAbstractSocket() { callbackQAbstractSocket_DestroyQAbstractSocket(this); };
	bool open(QIODevice::OpenMode mode) { return callbackQAbstractSocket_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQAbstractSocket_Pos(const_cast<MyQAbstractSocket*>(this)); };
	bool reset() { return callbackQAbstractSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQAbstractSocket_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQAbstractSocket_Size(const_cast<MyQAbstractSocket*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocket_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAbstractSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSocket_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAbstractSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSocket_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSocket_MetaObject(const_cast<MyQAbstractSocket*>(this))); };
};

void* QAbstractSocket_NewQAbstractSocket(long long socketType, void* parent)
{
	return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QObject*>(parent));
}

void QAbstractSocket_Abort(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->abort();
}

char QAbstractSocket_AtEnd(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->atEnd();
}

char QAbstractSocket_AtEndDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::atEnd();
}

char QAbstractSocket_Bind(void* ptr, void* address, unsigned short port, long long mode)
{
	return static_cast<QAbstractSocket*>(ptr)->bind(*static_cast<QHostAddress*>(address), port, static_cast<QAbstractSocket::BindFlag>(mode));
}

char QAbstractSocket_Bind2(void* ptr, unsigned short port, long long mode)
{
	return static_cast<QAbstractSocket*>(ptr)->bind(port, static_cast<QAbstractSocket::BindFlag>(mode));
}

long long QAbstractSocket_BytesAvailable(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->bytesAvailable();
}

long long QAbstractSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::bytesAvailable();
}

long long QAbstractSocket_BytesToWrite(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->bytesToWrite();
}

long long QAbstractSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::bytesToWrite();
}

char QAbstractSocket_CanReadLine(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->canReadLine();
}

char QAbstractSocket_CanReadLineDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::canReadLine();
}

void QAbstractSocket_Close(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->close();
}

void QAbstractSocket_CloseDefault(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::close();
}

void QAbstractSocket_ConnectToHost2(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QAbstractSocket*>(ptr)->connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QAbstractSocket_ConnectToHost2Default(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QAbstractSocket_ConnectToHost(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QAbstractSocket*>(ptr)->connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QAbstractSocket_ConnectToHostDefault(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
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

void QAbstractSocket_Error2(void* ptr, long long socketError)
{
	static_cast<QAbstractSocket*>(ptr)->error(static_cast<QAbstractSocket::SocketError>(socketError));
}

long long QAbstractSocket_Error(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->error();
}

char QAbstractSocket_Flush(void* ptr)
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

char QAbstractSocket_IsSequential(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->isSequential();
}

char QAbstractSocket_IsSequentialDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::isSequential();
}

char QAbstractSocket_IsValid(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->isValid();
}

void* QAbstractSocket_LocalAddress(void* ptr)
{
	return new QHostAddress(static_cast<QAbstractSocket*>(ptr)->localAddress());
}

unsigned short QAbstractSocket_LocalPort(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->localPort();
}

long long QAbstractSocket_PauseMode(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->pauseMode();
}

void* QAbstractSocket_PeerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QAbstractSocket*>(ptr)->peerAddress());
}

struct QtNetwork_PackedString QAbstractSocket_PeerName(void* ptr)
{
	return ({ QByteArray tf7fe06 = static_cast<QAbstractSocket*>(ptr)->peerName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tf7fe06.prepend("WHITESPACE").constData()+10), tf7fe06.size()-10 }; });
}

unsigned short QAbstractSocket_PeerPort(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->peerPort();
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
	return static_cast<QAbstractSocket*>(ptr)->readBufferSize();
}

long long QAbstractSocket_ReadLineData(void* ptr, char* data, long long maxlen)
{
	return static_cast<QAbstractSocket*>(ptr)->readLineData(data, maxlen);
}

long long QAbstractSocket_ReadLineDataDefault(void* ptr, char* data, long long maxlen)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::readLineData(data, maxlen);
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

void QAbstractSocket_SetPeerName(void* ptr, char* name)
{
	static_cast<QAbstractSocket*>(ptr)->setPeerName(QString(name));
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
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setReadBufferSize(size);
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
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QAbstractSocket_SetSocketState(void* ptr, long long state)
{
	static_cast<QAbstractSocket*>(ptr)->setSocketState(static_cast<QAbstractSocket::SocketState>(state));
}





void* QAbstractSocket_SocketOption(void* ptr, long long option)
{
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QAbstractSocket_SocketOptionDefault(void* ptr, long long option)
{
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

long long QAbstractSocket_SocketType(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->socketType();
}

long long QAbstractSocket_State(void* ptr)
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

void QAbstractSocket_StateChanged(void* ptr, long long socketState)
{
	static_cast<QAbstractSocket*>(ptr)->stateChanged(static_cast<QAbstractSocket::SocketState>(socketState));
}

char QAbstractSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForBytesWritten(msecs);
}

char QAbstractSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForBytesWritten(msecs);
}

char QAbstractSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForConnected(msecs);
}

char QAbstractSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForConnected(msecs);
}

char QAbstractSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForDisconnected(msecs);
}

char QAbstractSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForDisconnected(msecs);
}

char QAbstractSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->waitForReadyRead(msecs);
}

char QAbstractSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::waitForReadyRead(msecs);
}

long long QAbstractSocket_WriteData(void* ptr, char* data, long long size)
{
	return static_cast<QAbstractSocket*>(ptr)->writeData(const_cast<const char*>(data), size);
}

long long QAbstractSocket_WriteDataDefault(void* ptr, char* data, long long size)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::writeData(const_cast<const char*>(data), size);
}

void QAbstractSocket_DestroyQAbstractSocket(void* ptr)
{
	static_cast<QAbstractSocket*>(ptr)->~QAbstractSocket();
}

void QAbstractSocket_DestroyQAbstractSocketDefault(void* ptr)
{

}

char QAbstractSocket_Open(void* ptr, long long mode)
{
	return static_cast<QAbstractSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QAbstractSocket_OpenDefault(void* ptr, long long mode)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QAbstractSocket_Pos(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->pos();
}

long long QAbstractSocket_PosDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::pos();
}

char QAbstractSocket_Reset(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->reset();
}

char QAbstractSocket_ResetDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::reset();
}

char QAbstractSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QAbstractSocket*>(ptr)->seek(pos);
}

char QAbstractSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::seek(pos);
}

long long QAbstractSocket_Size(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->size();
}

long long QAbstractSocket_SizeDefault(void* ptr)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::size();
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

char QAbstractSocket_Event(void* ptr, void* e)
{
	return static_cast<QAbstractSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAbstractSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::event(static_cast<QEvent*>(e));
}

char QAbstractSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAbstractSocket_EventFilterDefault(void* ptr, void* watched, void* event)
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

char QAuthenticator_IsNull(void* ptr)
{
	return static_cast<QAuthenticator*>(ptr)->isNull();
}

void* QAuthenticator_Option(void* ptr, char* opt)
{
	return new QVariant(static_cast<QAuthenticator*>(ptr)->option(QString(opt)));
}

struct QtNetwork_PackedString QAuthenticator_Password(void* ptr)
{
	return ({ QByteArray t31072f = static_cast<QAuthenticator*>(ptr)->password().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t31072f.prepend("WHITESPACE").constData()+10), t31072f.size()-10 }; });
}

struct QtNetwork_PackedString QAuthenticator_Realm(void* ptr)
{
	return ({ QByteArray tcc1e3d = static_cast<QAuthenticator*>(ptr)->realm().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tcc1e3d.prepend("WHITESPACE").constData()+10), tcc1e3d.size()-10 }; });
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

struct QtNetwork_PackedString QAuthenticator_User(void* ptr)
{
	return ({ QByteArray ta76119 = static_cast<QAuthenticator*>(ptr)->user().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta76119.prepend("WHITESPACE").constData()+10), ta76119.size()-10 }; });
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

struct QtNetwork_PackedString QDnsDomainNameRecord_Name(void* ptr)
{
	return ({ QByteArray t22074d = static_cast<QDnsDomainNameRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t22074d.prepend("WHITESPACE").constData()+10), t22074d.size()-10 }; });
}

void QDnsDomainNameRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsDomainNameRecord*>(ptr)->swap(*static_cast<QDnsDomainNameRecord*>(other));
}

unsigned int QDnsDomainNameRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsDomainNameRecord*>(ptr)->timeToLive();
}

struct QtNetwork_PackedString QDnsDomainNameRecord_Value(void* ptr)
{
	return ({ QByteArray tb334d2 = static_cast<QDnsDomainNameRecord*>(ptr)->value().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb334d2.prepend("WHITESPACE").constData()+10), tb334d2.size()-10 }; });
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

struct QtNetwork_PackedString QDnsHostAddressRecord_Name(void* ptr)
{
	return ({ QByteArray tb52211 = static_cast<QDnsHostAddressRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb52211.prepend("WHITESPACE").constData()+10), tb52211.size()-10 }; });
}

void QDnsHostAddressRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsHostAddressRecord*>(ptr)->swap(*static_cast<QDnsHostAddressRecord*>(other));
}

unsigned int QDnsHostAddressRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsHostAddressRecord*>(ptr)->timeToLive();
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
	void abort() { callbackQDnsLookup_Abort(this); };
	void Signal_Finished() { callbackQDnsLookup_Finished(this); };
	void lookup() { callbackQDnsLookup_Lookup(this); };
	void Signal_NameChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtNetwork_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDnsLookup_NameChanged(this, namePacked); };
	void Signal_NameserverChanged(const QHostAddress & nameserver) { callbackQDnsLookup_NameserverChanged(this, const_cast<QHostAddress*>(&nameserver)); };
	void Signal_TypeChanged(QDnsLookup::Type ty) { callbackQDnsLookup_TypeChanged(this, ty); };
	void timerEvent(QTimerEvent * event) { callbackQDnsLookup_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDnsLookup_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDnsLookup_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDnsLookup_CustomEvent(this, event); };
	void deleteLater() { callbackQDnsLookup_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDnsLookup_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDnsLookup_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDnsLookup_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDnsLookup_MetaObject(const_cast<MyQDnsLookup*>(this))); };
};

void* QDnsLookup_NewQDnsLookup3(long long ty, char* name, void* nameserver, void* parent)
{
	return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), *static_cast<QHostAddress*>(nameserver), static_cast<QObject*>(parent));
}

long long QDnsLookup_Error(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->error();
}

struct QtNetwork_PackedString QDnsLookup_ErrorString(void* ptr)
{
	return ({ QByteArray ta68e7f = static_cast<QDnsLookup*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta68e7f.prepend("WHITESPACE").constData()+10), ta68e7f.size()-10 }; });
}

struct QtNetwork_PackedString QDnsLookup_Name(void* ptr)
{
	return ({ QByteArray td8dec7 = static_cast<QDnsLookup*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(td8dec7.prepend("WHITESPACE").constData()+10), td8dec7.size()-10 }; });
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

void QDnsLookup_SetType(void* ptr, long long vqd)
{
	static_cast<QDnsLookup*>(ptr)->setType(static_cast<QDnsLookup::Type>(vqd));
}

long long QDnsLookup_Type(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->type();
}

void* QDnsLookup_NewQDnsLookup(void* parent)
{
	return new MyQDnsLookup(static_cast<QObject*>(parent));
}

void* QDnsLookup_NewQDnsLookup2(long long ty, char* name, void* parent)
{
	return new MyQDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), static_cast<QObject*>(parent));
}

void QDnsLookup_Abort(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "abort");
}

struct QtNetwork_PackedList QDnsLookup_CanonicalNameRecords(void* ptr)
{
	return ({ QList<QDnsDomainNameRecord>* tmpValue = new QList<QDnsDomainNameRecord>(static_cast<QDnsLookup*>(ptr)->canonicalNameRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

struct QtNetwork_PackedList QDnsLookup_HostAddressRecords(void* ptr)
{
	return ({ QList<QDnsHostAddressRecord>* tmpValue = new QList<QDnsHostAddressRecord>(static_cast<QDnsLookup*>(ptr)->hostAddressRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QDnsLookup_IsFinished(void* ptr)
{
	return static_cast<QDnsLookup*>(ptr)->isFinished();
}

void QDnsLookup_Lookup(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "lookup");
}

struct QtNetwork_PackedList QDnsLookup_MailExchangeRecords(void* ptr)
{
	return ({ QList<QDnsMailExchangeRecord>* tmpValue = new QList<QDnsMailExchangeRecord>(static_cast<QDnsLookup*>(ptr)->mailExchangeRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

struct QtNetwork_PackedList QDnsLookup_NameServerRecords(void* ptr)
{
	return ({ QList<QDnsDomainNameRecord>* tmpValue = new QList<QDnsDomainNameRecord>(static_cast<QDnsLookup*>(ptr)->nameServerRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

struct QtNetwork_PackedList QDnsLookup_PointerRecords(void* ptr)
{
	return ({ QList<QDnsDomainNameRecord>* tmpValue = new QList<QDnsDomainNameRecord>(static_cast<QDnsLookup*>(ptr)->pointerRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_ServiceRecords(void* ptr)
{
	return ({ QList<QDnsServiceRecord>* tmpValue = new QList<QDnsServiceRecord>(static_cast<QDnsLookup*>(ptr)->serviceRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QDnsLookup_TextRecords(void* ptr)
{
	return ({ QList<QDnsTextRecord>* tmpValue = new QList<QDnsTextRecord>(static_cast<QDnsLookup*>(ptr)->textRecords()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void QDnsLookup_ConnectTypeChanged(void* ptr)
{
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

void* QDnsLookup_canonicalNameRecords_atList(void* ptr, int i)
{
	return new QDnsDomainNameRecord(static_cast<QList<QDnsDomainNameRecord>*>(ptr)->at(i));
}

void* QDnsLookup_hostAddressRecords_atList(void* ptr, int i)
{
	return new QDnsHostAddressRecord(static_cast<QList<QDnsHostAddressRecord>*>(ptr)->at(i));
}

void* QDnsLookup_mailExchangeRecords_atList(void* ptr, int i)
{
	return new QDnsMailExchangeRecord(static_cast<QList<QDnsMailExchangeRecord>*>(ptr)->at(i));
}

void* QDnsLookup_nameServerRecords_atList(void* ptr, int i)
{
	return new QDnsDomainNameRecord(static_cast<QList<QDnsDomainNameRecord>*>(ptr)->at(i));
}

void* QDnsLookup_pointerRecords_atList(void* ptr, int i)
{
	return new QDnsDomainNameRecord(static_cast<QList<QDnsDomainNameRecord>*>(ptr)->at(i));
}

void* QDnsLookup_serviceRecords_atList(void* ptr, int i)
{
	return new QDnsServiceRecord(static_cast<QList<QDnsServiceRecord>*>(ptr)->at(i));
}

void* QDnsLookup_textRecords_atList(void* ptr, int i)
{
	return new QDnsTextRecord(static_cast<QList<QDnsTextRecord>*>(ptr)->at(i));
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

char QDnsLookup_Event(void* ptr, void* e)
{
	return static_cast<QDnsLookup*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDnsLookup_EventDefault(void* ptr, void* e)
{
	return static_cast<QDnsLookup*>(ptr)->QDnsLookup::event(static_cast<QEvent*>(e));
}

char QDnsLookup_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDnsLookup*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDnsLookup_EventFilterDefault(void* ptr, void* watched, void* event)
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

void QDnsMailExchangeRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsMailExchangeRecord*>(ptr)->swap(*static_cast<QDnsMailExchangeRecord*>(other));
}

unsigned int QDnsMailExchangeRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsMailExchangeRecord*>(ptr)->timeToLive();
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

struct QtNetwork_PackedString QDnsServiceRecord_Name(void* ptr)
{
	return ({ QByteArray t9ceb61 = static_cast<QDnsServiceRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t9ceb61.prepend("WHITESPACE").constData()+10), t9ceb61.size()-10 }; });
}

unsigned short QDnsServiceRecord_Port(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->port();
}

unsigned short QDnsServiceRecord_Priority(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->priority();
}

void QDnsServiceRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsServiceRecord*>(ptr)->swap(*static_cast<QDnsServiceRecord*>(other));
}

struct QtNetwork_PackedString QDnsServiceRecord_Target(void* ptr)
{
	return ({ QByteArray te2873b = static_cast<QDnsServiceRecord*>(ptr)->target().toUtf8(); QtNetwork_PackedString { const_cast<char*>(te2873b.prepend("WHITESPACE").constData()+10), te2873b.size()-10 }; });
}

unsigned int QDnsServiceRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->timeToLive();
}

unsigned short QDnsServiceRecord_Weight(void* ptr)
{
	return static_cast<QDnsServiceRecord*>(ptr)->weight();
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

struct QtNetwork_PackedString QDnsTextRecord_Name(void* ptr)
{
	return ({ QByteArray tb66dc5 = static_cast<QDnsTextRecord*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb66dc5.prepend("WHITESPACE").constData()+10), tb66dc5.size()-10 }; });
}

void QDnsTextRecord_Swap(void* ptr, void* other)
{
	static_cast<QDnsTextRecord*>(ptr)->swap(*static_cast<QDnsTextRecord*>(other));
}

unsigned int QDnsTextRecord_TimeToLive(void* ptr)
{
	return static_cast<QDnsTextRecord*>(ptr)->timeToLive();
}

struct QtNetwork_PackedList QDnsTextRecord_Values(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QDnsTextRecord*>(ptr)->values()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void QDnsTextRecord_DestroyQDnsTextRecord(void* ptr)
{
	static_cast<QDnsTextRecord*>(ptr)->~QDnsTextRecord();
}

void* QDnsTextRecord_values_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
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

void* QHostAddress_NewQHostAddress7(char* address)
{
	return new QHostAddress(QString(address));
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

void QHostAddress_Clear(void* ptr)
{
	static_cast<QHostAddress*>(ptr)->clear();
}

char QHostAddress_IsInSubnet(void* ptr, void* subnet, int netmask)
{
	return static_cast<QHostAddress*>(ptr)->isInSubnet(*static_cast<QHostAddress*>(subnet), netmask);
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

long long QHostAddress_Protocol(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->protocol();
}

struct QtNetwork_PackedString QHostAddress_ScopeId(void* ptr)
{
	return ({ QByteArray t9c6602 = static_cast<QHostAddress*>(ptr)->scopeId().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t9c6602.prepend("WHITESPACE").constData()+10), t9c6602.size()-10 }; });
}

char QHostAddress_SetAddress6(void* ptr, char* address)
{
	return static_cast<QHostAddress*>(ptr)->setAddress(QString(address));
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

void QHostAddress_SetScopeId(void* ptr, char* id)
{
	static_cast<QHostAddress*>(ptr)->setScopeId(QString(id));
}

void QHostAddress_Swap(void* ptr, void* other)
{
	static_cast<QHostAddress*>(ptr)->swap(*static_cast<QHostAddress*>(other));
}

unsigned int QHostAddress_ToIPv4Address(void* ptr)
{
	return static_cast<QHostAddress*>(ptr)->toIPv4Address();
}

unsigned int QHostAddress_ToIPv4Address2(void* ptr, char ok)
{
	return static_cast<QHostAddress*>(ptr)->toIPv4Address(NULL);
}

struct QtNetwork_PackedString QHostAddress_ToString(void* ptr)
{
	return ({ QByteArray tc5ceab = static_cast<QHostAddress*>(ptr)->toString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc5ceab.prepend("WHITESPACE").constData()+10), tc5ceab.size()-10 }; });
}

void QHostAddress_DestroyQHostAddress(void* ptr)
{
	static_cast<QHostAddress*>(ptr)->~QHostAddress();
}

struct QtNetwork_PackedString QHostInfo_QHostInfo_LocalHostName()
{
	return ({ QByteArray t63826c = QHostInfo::localHostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t63826c.prepend("WHITESPACE").constData()+10), t63826c.size()-10 }; });
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

struct QtNetwork_PackedList QHostInfo_Addresses(void* ptr)
{
	return ({ QList<QHostAddress>* tmpValue = new QList<QHostAddress>(static_cast<QHostInfo*>(ptr)->addresses()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QHostInfo_Error(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->error();
}

struct QtNetwork_PackedString QHostInfo_ErrorString(void* ptr)
{
	return ({ QByteArray taf4307 = static_cast<QHostInfo*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(taf4307.prepend("WHITESPACE").constData()+10), taf4307.size()-10 }; });
}

void* QHostInfo_QHostInfo_FromName(char* name)
{
	return new QHostInfo(QHostInfo::fromName(QString(name)));
}

struct QtNetwork_PackedString QHostInfo_HostName(void* ptr)
{
	return ({ QByteArray t7a1d02 = static_cast<QHostInfo*>(ptr)->hostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t7a1d02.prepend("WHITESPACE").constData()+10), t7a1d02.size()-10 }; });
}

int QHostInfo_QHostInfo_LookupHost(char* name, void* receiver, char* member)
{
	return QHostInfo::lookupHost(QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QHostInfo_LookupId(void* ptr)
{
	return static_cast<QHostInfo*>(ptr)->lookupId();
}

void QHostInfo_SetError(void* ptr, long long error)
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

struct QtNetwork_PackedString QHostInfo_QHostInfo_LocalDomainName()
{
	return ({ QByteArray t5517d9 = QHostInfo::localDomainName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t5517d9.prepend("WHITESPACE").constData()+10), t5517d9.size()-10 }; });
}

void* QHostInfo_addresses_atList(void* ptr, int i)
{
	return new QHostAddress(static_cast<QList<QHostAddress>*>(ptr)->at(i));
}

void* QHttpMultiPart_NewQHttpMultiPart2(long long contentType, void* parent)
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

void* QHttpMultiPart_Boundary(void* ptr)
{
	return new QByteArray(static_cast<QHttpMultiPart*>(ptr)->boundary());
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

char QHttpMultiPart_Event(void* ptr, void* e)
{
	return static_cast<QHttpMultiPart*>(ptr)->event(static_cast<QEvent*>(e));
}

char QHttpMultiPart_EventDefault(void* ptr, void* e)
{
	return static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::event(static_cast<QEvent*>(e));
}

char QHttpMultiPart_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHttpMultiPart*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QHttpMultiPart_EventFilterDefault(void* ptr, void* watched, void* event)
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
	MyQLocalServer(QObject *parent) : QLocalServer(parent) {};
	bool hasPendingConnections() const { return callbackQLocalServer_HasPendingConnections(const_cast<MyQLocalServer*>(this)) != 0; };
	void incomingConnection(quintptr socketDescriptor) { callbackQLocalServer_IncomingConnection(this, socketDescriptor); };
	void Signal_NewConnection() { callbackQLocalServer_NewConnection(this); };
	QLocalSocket * nextPendingConnection() { return static_cast<QLocalSocket*>(callbackQLocalServer_NextPendingConnection(this)); };
	void timerEvent(QTimerEvent * event) { callbackQLocalServer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQLocalServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLocalServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLocalServer_CustomEvent(this, event); };
	void deleteLater() { callbackQLocalServer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLocalServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLocalServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLocalServer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLocalServer_MetaObject(const_cast<MyQLocalServer*>(this))); };
};

void QLocalServer_SetSocketOptions(void* ptr, long long options)
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

struct QtNetwork_PackedString QLocalServer_ErrorString(void* ptr)
{
	return ({ QByteArray tf5dac0 = static_cast<QLocalServer*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tf5dac0.prepend("WHITESPACE").constData()+10), tf5dac0.size()-10 }; });
}

struct QtNetwork_PackedString QLocalServer_FullServerName(void* ptr)
{
	return ({ QByteArray tb91f50 = static_cast<QLocalServer*>(ptr)->fullServerName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb91f50.prepend("WHITESPACE").constData()+10), tb91f50.size()-10 }; });
}

char QLocalServer_HasPendingConnections(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->hasPendingConnections();
}

char QLocalServer_HasPendingConnectionsDefault(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->QLocalServer::hasPendingConnections();
}

void QLocalServer_IncomingConnection(void* ptr, uintptr_t socketDescriptor)
{
	static_cast<QLocalServer*>(ptr)->incomingConnection(socketDescriptor);
}

void QLocalServer_IncomingConnectionDefault(void* ptr, uintptr_t socketDescriptor)
{
	static_cast<QLocalServer*>(ptr)->QLocalServer::incomingConnection(socketDescriptor);
}

char QLocalServer_IsListening(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->isListening();
}

char QLocalServer_Listen(void* ptr, char* name)
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

char QLocalServer_QLocalServer_RemoveServer(char* name)
{
	return QLocalServer::removeServer(QString(name));
}

long long QLocalServer_ServerError(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->serverError();
}

struct QtNetwork_PackedString QLocalServer_ServerName(void* ptr)
{
	return ({ QByteArray t054e78 = static_cast<QLocalServer*>(ptr)->serverName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t054e78.prepend("WHITESPACE").constData()+10), t054e78.size()-10 }; });
}

void QLocalServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QLocalServer*>(ptr)->setMaxPendingConnections(numConnections);
}

long long QLocalServer_SocketOptions(void* ptr)
{
	return static_cast<QLocalServer*>(ptr)->socketOptions();
}

char QLocalServer_WaitForNewConnection(void* ptr, int msec, char timedOut)
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

char QLocalServer_Event(void* ptr, void* e)
{
	return static_cast<QLocalServer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QLocalServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QLocalServer*>(ptr)->QLocalServer::event(static_cast<QEvent*>(e));
}

char QLocalServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLocalServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QLocalServer_EventFilterDefault(void* ptr, void* watched, void* event)
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
	bool open(QIODevice::OpenMode openMode) { return callbackQLocalSocket_Open(this, openMode) != 0; };
	void Signal_Connected() { callbackQLocalSocket_Connected(this); };
	void Signal_Disconnected() { callbackQLocalSocket_Disconnected(this); };
	void Signal_Error2(QLocalSocket::LocalSocketError socketError) { callbackQLocalSocket_Error2(this, socketError); };
	bool isSequential() const { return callbackQLocalSocket_IsSequential(const_cast<MyQLocalSocket*>(this)) != 0; };
	void Signal_StateChanged(QLocalSocket::LocalSocketState socketState) { callbackQLocalSocket_StateChanged(this, socketState); };
	qint64 bytesAvailable() const { return callbackQLocalSocket_BytesAvailable(const_cast<MyQLocalSocket*>(this)); };
	qint64 bytesToWrite() const { return callbackQLocalSocket_BytesToWrite(const_cast<MyQLocalSocket*>(this)); };
	bool canReadLine() const { return callbackQLocalSocket_CanReadLine(const_cast<MyQLocalSocket*>(this)) != 0; };
	void close() { callbackQLocalSocket_Close(this); };
	bool waitForBytesWritten(int msecs) { return callbackQLocalSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQLocalSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 c) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), c };return callbackQLocalSocket_WriteData(this, dataPacked, c); };
	bool atEnd() const { return callbackQLocalSocket_AtEnd(const_cast<MyQLocalSocket*>(this)) != 0; };
	qint64 pos() const { return callbackQLocalSocket_Pos(const_cast<MyQLocalSocket*>(this)); };
	qint64 readLineData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQLocalSocket_ReadLineData(this, dataPacked, maxSize); };
	bool reset() { return callbackQLocalSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQLocalSocket_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQLocalSocket_Size(const_cast<MyQLocalSocket*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQLocalSocket_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQLocalSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLocalSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLocalSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQLocalSocket_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLocalSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLocalSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLocalSocket_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLocalSocket_MetaObject(const_cast<MyQLocalSocket*>(this))); };
};

char QLocalSocket_Open(void* ptr, long long openMode)
{
	return static_cast<QLocalSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

char QLocalSocket_OpenDefault(void* ptr, long long openMode)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QLocalSocket_NewQLocalSocket(void* parent)
{
	return new MyQLocalSocket(static_cast<QObject*>(parent));
}

void QLocalSocket_ConnectToServer2(void* ptr, char* name, long long openMode)
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

void QLocalSocket_Error2(void* ptr, long long socketError)
{
	static_cast<QLocalSocket*>(ptr)->error(static_cast<QLocalSocket::LocalSocketError>(socketError));
}

struct QtNetwork_PackedString QLocalSocket_FullServerName(void* ptr)
{
	return ({ QByteArray ta11089 = static_cast<QLocalSocket*>(ptr)->fullServerName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta11089.prepend("WHITESPACE").constData()+10), ta11089.size()-10 }; });
}

char QLocalSocket_IsSequential(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->isSequential();
}

char QLocalSocket_IsSequentialDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::isSequential();
}

struct QtNetwork_PackedString QLocalSocket_ServerName(void* ptr)
{
	return ({ QByteArray t348d56 = static_cast<QLocalSocket*>(ptr)->serverName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t348d56.prepend("WHITESPACE").constData()+10), t348d56.size()-10 }; });
}

void QLocalSocket_SetServerName(void* ptr, char* name)
{
	static_cast<QLocalSocket*>(ptr)->setServerName(QString(name));
}

long long QLocalSocket_State(void* ptr)
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

void QLocalSocket_StateChanged(void* ptr, long long socketState)
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
	return static_cast<QLocalSocket*>(ptr)->bytesAvailable();
}

long long QLocalSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::bytesAvailable();
}

long long QLocalSocket_BytesToWrite(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->bytesToWrite();
}

long long QLocalSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::bytesToWrite();
}

char QLocalSocket_CanReadLine(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->canReadLine();
}

char QLocalSocket_CanReadLineDefault(void* ptr)
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

void QLocalSocket_ConnectToServer(void* ptr, long long openMode)
{
	static_cast<QLocalSocket*>(ptr)->connectToServer(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_DisconnectFromServer(void* ptr)
{
	static_cast<QLocalSocket*>(ptr)->disconnectFromServer();
}

long long QLocalSocket_Error(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->error();
}

char QLocalSocket_Flush(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->flush();
}

char QLocalSocket_IsValid(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->isValid();
}

long long QLocalSocket_ReadBufferSize(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->readBufferSize();
}

void QLocalSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QLocalSocket*>(ptr)->setReadBufferSize(size);
}

char QLocalSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForBytesWritten(msecs);
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

char QLocalSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->waitForReadyRead(msecs);
}

char QLocalSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::waitForReadyRead(msecs);
}

long long QLocalSocket_WriteData(void* ptr, char* data, long long c)
{
	return static_cast<QLocalSocket*>(ptr)->writeData(const_cast<const char*>(data), c);
}

long long QLocalSocket_WriteDataDefault(void* ptr, char* data, long long c)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::writeData(const_cast<const char*>(data), c);
}

char QLocalSocket_AtEnd(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->atEnd();
}

char QLocalSocket_AtEndDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::atEnd();
}

long long QLocalSocket_Pos(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->pos();
}

long long QLocalSocket_PosDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::pos();
}

long long QLocalSocket_ReadLineData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QLocalSocket*>(ptr)->readLineData(data, maxSize);
}

long long QLocalSocket_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::readLineData(data, maxSize);
}

char QLocalSocket_Reset(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->reset();
}

char QLocalSocket_ResetDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::reset();
}

char QLocalSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QLocalSocket*>(ptr)->seek(pos);
}

char QLocalSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::seek(pos);
}

long long QLocalSocket_Size(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->size();
}

long long QLocalSocket_SizeDefault(void* ptr)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::size();
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

char QLocalSocket_Event(void* ptr, void* e)
{
	return static_cast<QLocalSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

char QLocalSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QLocalSocket*>(ptr)->QLocalSocket::event(static_cast<QEvent*>(e));
}

char QLocalSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLocalSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QLocalSocket_EventFilterDefault(void* ptr, void* watched, void* event)
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
	void Signal_AuthenticationRequired(QNetworkReply * reply, QAuthenticator * authenticator) { callbackQNetworkAccessManager_AuthenticationRequired(this, reply, authenticator); };
	QNetworkReply * createRequest(QNetworkAccessManager::Operation op, const QNetworkRequest & req, QIODevice * outgoingData) { return static_cast<QNetworkReply*>(callbackQNetworkAccessManager_CreateRequest(this, op, const_cast<QNetworkRequest*>(&req), outgoingData)); };
	void Signal_Encrypted(QNetworkReply * reply) { callbackQNetworkAccessManager_Encrypted(this, reply); };
	void Signal_Finished(QNetworkReply * reply) { callbackQNetworkAccessManager_Finished(this, reply); };
	void Signal_NetworkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility accessible) { callbackQNetworkAccessManager_NetworkAccessibleChanged(this, accessible); };
	void Signal_PreSharedKeyAuthenticationRequired(QNetworkReply * reply, QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired(this, reply, authenticator); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQNetworkAccessManager_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	QStringList supportedSchemesImplementation() const { return QString(callbackQNetworkAccessManager_SupportedSchemesImplementation(const_cast<MyQNetworkAccessManager*>(this))).split("|", QString::SkipEmptyParts); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkAccessManager_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNetworkAccessManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkAccessManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkAccessManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkAccessManager_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkAccessManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNetworkAccessManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkAccessManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkAccessManager_MetaObject(const_cast<MyQNetworkAccessManager*>(this))); };
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

void QNetworkAccessManager_ConnectToHost(void* ptr, char* hostName, unsigned short port)
{
	static_cast<QNetworkAccessManager*>(ptr)->connectToHost(QString(hostName), port);
}

void QNetworkAccessManager_ConnectToHostEncrypted(void* ptr, char* hostName, unsigned short port, void* sslConfiguration)
{
	static_cast<QNetworkAccessManager*>(ptr)->connectToHostEncrypted(QString(hostName), port, *static_cast<QSslConfiguration*>(sslConfiguration));
}

void* QNetworkAccessManager_CookieJar(void* ptr)
{
	return static_cast<QNetworkAccessManager*>(ptr)->cookieJar();
}

void* QNetworkAccessManager_CreateRequest(void* ptr, long long op, void* req, void* outgoingData)
{
	return static_cast<QNetworkAccessManager*>(ptr)->createRequest(static_cast<QNetworkAccessManager::Operation>(op), *static_cast<QNetworkRequest*>(req), static_cast<QIODevice*>(outgoingData));
}

void* QNetworkAccessManager_CreateRequestDefault(void* ptr, long long op, void* req, void* outgoingData)
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

long long QNetworkAccessManager_NetworkAccessible(void* ptr)
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

void QNetworkAccessManager_NetworkAccessibleChanged(void* ptr, long long accessible)
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

void* QNetworkAccessManager_Post2(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
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

void* QNetworkAccessManager_Put2(void* ptr, void* request, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void* QNetworkAccessManager_SendCustomRequest(void* ptr, void* request, void* verb, void* data)
{
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), static_cast<QIODevice*>(data));
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

char QNetworkAccessManager_Event(void* ptr, void* e)
{
	return static_cast<QNetworkAccessManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNetworkAccessManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::event(static_cast<QEvent*>(e));
}

char QNetworkAccessManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkAccessManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNetworkAccessManager_EventFilterDefault(void* ptr, void* watched, void* event)
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

char QNetworkCacheMetaData_IsValid(void* ptr)
{
	return static_cast<QNetworkCacheMetaData*>(ptr)->isValid();
}

void* QNetworkCacheMetaData_LastModified(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->lastModified());
}

char QNetworkCacheMetaData_SaveToDisk(void* ptr)
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

long long QNetworkConfiguration_BearerType(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerType();
}

long long QNetworkConfiguration_BearerTypeFamily(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeFamily();
}

struct QtNetwork_PackedString QNetworkConfiguration_BearerTypeName(void* ptr)
{
	return ({ QByteArray t89aad8 = static_cast<QNetworkConfiguration*>(ptr)->bearerTypeName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t89aad8.prepend("WHITESPACE").constData()+10), t89aad8.size()-10 }; });
}

struct QtNetwork_PackedList QNetworkConfiguration_Children(void* ptr)
{
	return ({ QList<QNetworkConfiguration>* tmpValue = new QList<QNetworkConfiguration>(static_cast<QNetworkConfiguration*>(ptr)->children()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedString QNetworkConfiguration_Identifier(void* ptr)
{
	return ({ QByteArray tae5c30 = static_cast<QNetworkConfiguration*>(ptr)->identifier().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tae5c30.prepend("WHITESPACE").constData()+10), tae5c30.size()-10 }; });
}

char QNetworkConfiguration_IsRoamingAvailable(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->isRoamingAvailable();
}

char QNetworkConfiguration_IsValid(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->isValid();
}

struct QtNetwork_PackedString QNetworkConfiguration_Name(void* ptr)
{
	return ({ QByteArray t38ed5d = static_cast<QNetworkConfiguration*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t38ed5d.prepend("WHITESPACE").constData()+10), t38ed5d.size()-10 }; });
}

long long QNetworkConfiguration_Purpose(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->purpose();
}

long long QNetworkConfiguration_State(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->state();
}

void QNetworkConfiguration_Swap(void* ptr, void* other)
{
	static_cast<QNetworkConfiguration*>(ptr)->swap(*static_cast<QNetworkConfiguration*>(other));
}

long long QNetworkConfiguration_Type(void* ptr)
{
	return static_cast<QNetworkConfiguration*>(ptr)->type();
}

void QNetworkConfiguration_DestroyQNetworkConfiguration(void* ptr)
{
	static_cast<QNetworkConfiguration*>(ptr)->~QNetworkConfiguration();
}

void* QNetworkConfiguration_children_atList(void* ptr, int i)
{
	return new QNetworkConfiguration(static_cast<QList<QNetworkConfiguration>*>(ptr)->at(i));
}

class MyQNetworkConfigurationManager: public QNetworkConfigurationManager
{
public:
	MyQNetworkConfigurationManager(QObject *parent) : QNetworkConfigurationManager(parent) {};
	void Signal_ConfigurationAdded(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationAdded(this, const_cast<QNetworkConfiguration*>(&config)); };
	void Signal_ConfigurationChanged(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationChanged(this, const_cast<QNetworkConfiguration*>(&config)); };
	void Signal_ConfigurationRemoved(const QNetworkConfiguration & config) { callbackQNetworkConfigurationManager_ConfigurationRemoved(this, const_cast<QNetworkConfiguration*>(&config)); };
	void Signal_OnlineStateChanged(bool isOnline) { callbackQNetworkConfigurationManager_OnlineStateChanged(this, isOnline); };
	void Signal_UpdateCompleted() { callbackQNetworkConfigurationManager_UpdateCompleted(this); };
	void updateConfigurations() { callbackQNetworkConfigurationManager_UpdateConfigurations(this); };
	 ~MyQNetworkConfigurationManager() { callbackQNetworkConfigurationManager_DestroyQNetworkConfigurationManager(this); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkConfigurationManager_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNetworkConfigurationManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkConfigurationManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkConfigurationManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkConfigurationManager_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkConfigurationManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNetworkConfigurationManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkConfigurationManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkConfigurationManager_MetaObject(const_cast<MyQNetworkConfigurationManager*>(this))); };
};

void* QNetworkConfigurationManager_NewQNetworkConfigurationManager(void* parent)
{
	return new MyQNetworkConfigurationManager(static_cast<QObject*>(parent));
}

struct QtNetwork_PackedList QNetworkConfigurationManager_AllConfigurations(void* ptr, long long filter)
{
	return ({ QList<QNetworkConfiguration>* tmpValue = new QList<QNetworkConfiguration>(static_cast<QNetworkConfigurationManager*>(ptr)->allConfigurations(static_cast<QNetworkConfiguration::StateFlag>(filter))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkConfigurationManager_Capabilities(void* ptr)
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

char QNetworkConfigurationManager_IsOnline(void* ptr)
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

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(void* ptr)
{
	static_cast<QNetworkConfigurationManager*>(ptr)->~QNetworkConfigurationManager();
}

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManagerDefault(void* ptr)
{

}

void* QNetworkConfigurationManager_allConfigurations_atList(void* ptr, int i)
{
	return new QNetworkConfiguration(static_cast<QList<QNetworkConfiguration>*>(ptr)->at(i));
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

char QNetworkConfigurationManager_Event(void* ptr, void* e)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNetworkConfigurationManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::event(static_cast<QEvent*>(e));
}

char QNetworkConfigurationManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkConfigurationManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNetworkConfigurationManager_EventFilterDefault(void* ptr, void* watched, void* event)
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

void* QNetworkCookie_NewQNetworkCookie(void* name, void* value)
{
	return new QNetworkCookie(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

void* QNetworkCookie_NewQNetworkCookie2(void* other)
{
	return new QNetworkCookie(*static_cast<QNetworkCookie*>(other));
}

struct QtNetwork_PackedString QNetworkCookie_Domain(void* ptr)
{
	return ({ QByteArray tb84845 = static_cast<QNetworkCookie*>(ptr)->domain().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb84845.prepend("WHITESPACE").constData()+10), tb84845.size()-10 }; });
}

void* QNetworkCookie_ExpirationDate(void* ptr)
{
	return new QDateTime(static_cast<QNetworkCookie*>(ptr)->expirationDate());
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

void* QNetworkCookie_Name(void* ptr)
{
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->name());
}

void QNetworkCookie_Normalize(void* ptr, void* url)
{
	static_cast<QNetworkCookie*>(ptr)->normalize(*static_cast<QUrl*>(url));
}

struct QtNetwork_PackedList QNetworkCookie_QNetworkCookie_ParseCookies(void* cookieString)
{
	return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(QNetworkCookie::parseCookies(*static_cast<QByteArray*>(cookieString))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedString QNetworkCookie_Path(void* ptr)
{
	return ({ QByteArray tc58c07 = static_cast<QNetworkCookie*>(ptr)->path().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc58c07.prepend("WHITESPACE").constData()+10), tc58c07.size()-10 }; });
}

void QNetworkCookie_SetDomain(void* ptr, char* domain)
{
	static_cast<QNetworkCookie*>(ptr)->setDomain(QString(domain));
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

void QNetworkCookie_SetPath(void* ptr, char* path)
{
	static_cast<QNetworkCookie*>(ptr)->setPath(QString(path));
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

void* QNetworkCookie_ToRawForm(void* ptr, long long form)
{
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->toRawForm(static_cast<QNetworkCookie::RawForm>(form)));
}

void* QNetworkCookie_Value(void* ptr)
{
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->value());
}

void QNetworkCookie_DestroyQNetworkCookie(void* ptr)
{
	static_cast<QNetworkCookie*>(ptr)->~QNetworkCookie();
}

void* QNetworkCookie_parseCookies_atList(void* ptr, int i)
{
	return new QNetworkCookie(static_cast<QList<QNetworkCookie>*>(ptr)->at(i));
}

class MyQNetworkCookieJar: public QNetworkCookieJar
{
public:
	MyQNetworkCookieJar(QObject *parent) : QNetworkCookieJar(parent) {};
	bool deleteCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_DeleteCookie(this, const_cast<QNetworkCookie*>(&cookie)) != 0; };
	bool insertCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_InsertCookie(this, const_cast<QNetworkCookie*>(&cookie)) != 0; };
	bool updateCookie(const QNetworkCookie & cookie) { return callbackQNetworkCookieJar_UpdateCookie(this, const_cast<QNetworkCookie*>(&cookie)) != 0; };
	bool validateCookie(const QNetworkCookie & cookie, const QUrl & url) const { return callbackQNetworkCookieJar_ValidateCookie(const_cast<MyQNetworkCookieJar*>(this), const_cast<QNetworkCookie*>(&cookie), const_cast<QUrl*>(&url)) != 0; };
	 ~MyQNetworkCookieJar() { callbackQNetworkCookieJar_DestroyQNetworkCookieJar(this); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkCookieJar_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNetworkCookieJar_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkCookieJar_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkCookieJar_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkCookieJar_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkCookieJar_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNetworkCookieJar_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkCookieJar_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkCookieJar_MetaObject(const_cast<MyQNetworkCookieJar*>(this))); };
};

void* QNetworkCookieJar_NewQNetworkCookieJar(void* parent)
{
	return new MyQNetworkCookieJar(static_cast<QObject*>(parent));
}

struct QtNetwork_PackedList QNetworkCookieJar_AllCookies(void* ptr)
{
	return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(static_cast<QNetworkCookieJar*>(ptr)->allCookies()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkCookieJar_CookiesForUrl(void* ptr, void* url)
{
	return ({ QList<QNetworkCookie>* tmpValue = new QList<QNetworkCookie>(static_cast<QNetworkCookieJar*>(ptr)->cookiesForUrl(*static_cast<QUrl*>(url))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

char QNetworkCookieJar_UpdateCookie(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_UpdateCookieDefault(void* ptr, void* cookie)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

char QNetworkCookieJar_ValidateCookie(void* ptr, void* cookie, void* url)
{
	return static_cast<QNetworkCookieJar*>(ptr)->validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

char QNetworkCookieJar_ValidateCookieDefault(void* ptr, void* cookie, void* url)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

void QNetworkCookieJar_DestroyQNetworkCookieJar(void* ptr)
{
	static_cast<QNetworkCookieJar*>(ptr)->~QNetworkCookieJar();
}

void QNetworkCookieJar_DestroyQNetworkCookieJarDefault(void* ptr)
{

}

void* QNetworkCookieJar_allCookies_atList(void* ptr, int i)
{
	return new QNetworkCookie(static_cast<QList<QNetworkCookie>*>(ptr)->at(i));
}

void* QNetworkCookieJar_cookiesForUrl_atList(void* ptr, int i)
{
	return new QNetworkCookie(static_cast<QList<QNetworkCookie>*>(ptr)->at(i));
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

char QNetworkCookieJar_Event(void* ptr, void* e)
{
	return static_cast<QNetworkCookieJar*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNetworkCookieJar_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::event(static_cast<QEvent*>(e));
}

char QNetworkCookieJar_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkCookieJar*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNetworkCookieJar_EventFilterDefault(void* ptr, void* watched, void* event)
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
	qint64 cacheSize() const { return callbackQNetworkDiskCache_CacheSize(const_cast<MyQNetworkDiskCache*>(this)); };
	void clear() { callbackQNetworkDiskCache_Clear(this); };
	QIODevice * data(const QUrl & url) { return static_cast<QIODevice*>(callbackQNetworkDiskCache_Data(this, const_cast<QUrl*>(&url))); };
	qint64 expire() { return callbackQNetworkDiskCache_Expire(this); };
	void insert(QIODevice * device) { callbackQNetworkDiskCache_Insert(this, device); };
	QNetworkCacheMetaData metaData(const QUrl & url) { return *static_cast<QNetworkCacheMetaData*>(callbackQNetworkDiskCache_MetaData(this, const_cast<QUrl*>(&url))); };
	QIODevice * prepare(const QNetworkCacheMetaData & metaData) { return static_cast<QIODevice*>(callbackQNetworkDiskCache_Prepare(this, const_cast<QNetworkCacheMetaData*>(&metaData))); };
	bool remove(const QUrl & url) { return callbackQNetworkDiskCache_Remove(this, const_cast<QUrl*>(&url)) != 0; };
	void updateMetaData(const QNetworkCacheMetaData & metaData) { callbackQNetworkDiskCache_UpdateMetaData(this, const_cast<QNetworkCacheMetaData*>(&metaData)); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkDiskCache_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNetworkDiskCache_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkDiskCache_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkDiskCache_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkDiskCache_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkDiskCache_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNetworkDiskCache_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkDiskCache_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkDiskCache_MetaObject(const_cast<MyQNetworkDiskCache*>(this))); };
};

void* QNetworkDiskCache_NewQNetworkDiskCache(void* parent)
{
	return new MyQNetworkDiskCache(static_cast<QObject*>(parent));
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
	return static_cast<QNetworkDiskCache*>(ptr)->expire();
}

long long QNetworkDiskCache_ExpireDefault(void* ptr)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::expire();
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
	return static_cast<QNetworkDiskCache*>(ptr)->maximumCacheSize();
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

char QNetworkDiskCache_Remove(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->remove(*static_cast<QUrl*>(url));
}

char QNetworkDiskCache_RemoveDefault(void* ptr, void* url)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::remove(*static_cast<QUrl*>(url));
}

void QNetworkDiskCache_SetCacheDirectory(void* ptr, char* cacheDir)
{
	static_cast<QNetworkDiskCache*>(ptr)->setCacheDirectory(QString(cacheDir));
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

char QNetworkDiskCache_Event(void* ptr, void* e)
{
	return static_cast<QNetworkDiskCache*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNetworkDiskCache_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::event(static_cast<QEvent*>(e));
}

char QNetworkDiskCache_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkDiskCache*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNetworkDiskCache_EventFilterDefault(void* ptr, void* watched, void* event)
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

struct QtNetwork_PackedList QNetworkInterface_AddressEntries(void* ptr)
{
	return ({ QList<QNetworkAddressEntry>* tmpValue = new QList<QNetworkAddressEntry>(static_cast<QNetworkInterface*>(ptr)->addressEntries()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkInterface_QNetworkInterface_AllAddresses()
{
	return ({ QList<QHostAddress>* tmpValue = new QList<QHostAddress>(QNetworkInterface::allAddresses()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QNetworkInterface_QNetworkInterface_AllInterfaces()
{
	return ({ QList<QNetworkInterface>* tmpValue = new QList<QNetworkInterface>(QNetworkInterface::allInterfaces()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkInterface_Flags(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->flags();
}

struct QtNetwork_PackedString QNetworkInterface_HardwareAddress(void* ptr)
{
	return ({ QByteArray t25386c = static_cast<QNetworkInterface*>(ptr)->hardwareAddress().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t25386c.prepend("WHITESPACE").constData()+10), t25386c.size()-10 }; });
}

struct QtNetwork_PackedString QNetworkInterface_HumanReadableName(void* ptr)
{
	return ({ QByteArray tebd539 = static_cast<QNetworkInterface*>(ptr)->humanReadableName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tebd539.prepend("WHITESPACE").constData()+10), tebd539.size()-10 }; });
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

int QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(char* name)
{
	return QNetworkInterface::interfaceIndexFromName(QString(name));
}

struct QtNetwork_PackedString QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(int index)
{
	return ({ QByteArray ta95340 = QNetworkInterface::interfaceNameFromIndex(index).toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta95340.prepend("WHITESPACE").constData()+10), ta95340.size()-10 }; });
}

char QNetworkInterface_IsValid(void* ptr)
{
	return static_cast<QNetworkInterface*>(ptr)->isValid();
}

struct QtNetwork_PackedString QNetworkInterface_Name(void* ptr)
{
	return ({ QByteArray tb9dead = static_cast<QNetworkInterface*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb9dead.prepend("WHITESPACE").constData()+10), tb9dead.size()-10 }; });
}

void QNetworkInterface_Swap(void* ptr, void* other)
{
	static_cast<QNetworkInterface*>(ptr)->swap(*static_cast<QNetworkInterface*>(other));
}

void QNetworkInterface_DestroyQNetworkInterface(void* ptr)
{
	static_cast<QNetworkInterface*>(ptr)->~QNetworkInterface();
}

void* QNetworkInterface_addressEntries_atList(void* ptr, int i)
{
	return new QNetworkAddressEntry(static_cast<QList<QNetworkAddressEntry>*>(ptr)->at(i));
}

void* QNetworkInterface_allAddresses_atList(void* ptr, int i)
{
	return new QHostAddress(static_cast<QList<QHostAddress>*>(ptr)->at(i));
}

void* QNetworkInterface_allInterfaces_atList(void* ptr, int i)
{
	return new QNetworkInterface(static_cast<QList<QNetworkInterface>*>(ptr)->at(i));
}

void* QNetworkProxy_NewQNetworkProxy()
{
	return new QNetworkProxy();
}

void* QNetworkProxy_NewQNetworkProxy2(long long ty, char* hostName, unsigned short port, char* user, char* password)
{
	return new QNetworkProxy(static_cast<QNetworkProxy::ProxyType>(ty), QString(hostName), port, QString(user), QString(password));
}

void* QNetworkProxy_NewQNetworkProxy3(void* other)
{
	return new QNetworkProxy(*static_cast<QNetworkProxy*>(other));
}

void* QNetworkProxy_QNetworkProxy_ApplicationProxy()
{
	return new QNetworkProxy(QNetworkProxy::applicationProxy());
}

long long QNetworkProxy_Capabilities(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->capabilities();
}

char QNetworkProxy_HasRawHeader(void* ptr, void* headerName)
{
	return static_cast<QNetworkProxy*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkProxy_Header(void* ptr, long long header)
{
	return new QVariant(static_cast<QNetworkProxy*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

struct QtNetwork_PackedString QNetworkProxy_HostName(void* ptr)
{
	return ({ QByteArray t422f46 = static_cast<QNetworkProxy*>(ptr)->hostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t422f46.prepend("WHITESPACE").constData()+10), t422f46.size()-10 }; });
}

char QNetworkProxy_IsCachingProxy(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->isCachingProxy();
}

char QNetworkProxy_IsTransparentProxy(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->isTransparentProxy();
}

struct QtNetwork_PackedString QNetworkProxy_Password(void* ptr)
{
	return ({ QByteArray t6e003a = static_cast<QNetworkProxy*>(ptr)->password().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t6e003a.prepend("WHITESPACE").constData()+10), t6e003a.size()-10 }; });
}

unsigned short QNetworkProxy_Port(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->port();
}

void* QNetworkProxy_RawHeader(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QNetworkProxy*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

struct QtNetwork_PackedList QNetworkProxy_RawHeaderList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QNetworkProxy*>(ptr)->rawHeaderList()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

void QNetworkProxy_SetHostName(void* ptr, char* hostName)
{
	static_cast<QNetworkProxy*>(ptr)->setHostName(QString(hostName));
}

void QNetworkProxy_SetPassword(void* ptr, char* password)
{
	static_cast<QNetworkProxy*>(ptr)->setPassword(QString(password));
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

void QNetworkProxy_SetUser(void* ptr, char* user)
{
	static_cast<QNetworkProxy*>(ptr)->setUser(QString(user));
}

void QNetworkProxy_Swap(void* ptr, void* other)
{
	static_cast<QNetworkProxy*>(ptr)->swap(*static_cast<QNetworkProxy*>(other));
}

long long QNetworkProxy_Type(void* ptr)
{
	return static_cast<QNetworkProxy*>(ptr)->type();
}

struct QtNetwork_PackedString QNetworkProxy_User(void* ptr)
{
	return ({ QByteArray tcd9fd7 = static_cast<QNetworkProxy*>(ptr)->user().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tcd9fd7.prepend("WHITESPACE").constData()+10), tcd9fd7.size()-10 }; });
}

void QNetworkProxy_DestroyQNetworkProxy(void* ptr)
{
	static_cast<QNetworkProxy*>(ptr)->~QNetworkProxy();
}

void* QNetworkProxy_rawHeaderList_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

class MyQNetworkProxyFactory: public QNetworkProxyFactory
{
public:
	
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

}

struct QtNetwork_PackedList QNetworkProxyFactory_QNetworkProxyFactory_SystemProxyForQuery(void* query)
{
	return ({ QList<QNetworkProxy>* tmpValue = new QList<QNetworkProxy>(QNetworkProxyFactory::systemProxyForQuery(*static_cast<QNetworkProxyQuery*>(query))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QNetworkProxyFactory_proxyForQuery_atList(void* ptr, int i)
{
	return new QNetworkProxy(static_cast<QList<QNetworkProxy>*>(ptr)->at(i));
}

void* QNetworkProxyFactory_queryProxy_atList(void* ptr, int i)
{
	return new QNetworkProxy(static_cast<QList<QNetworkProxy>*>(ptr)->at(i));
}

void* QNetworkProxyFactory_systemProxyForQuery_atList(void* ptr, int i)
{
	return new QNetworkProxy(static_cast<QList<QNetworkProxy>*>(ptr)->at(i));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery()
{
	return new QNetworkProxyQuery();
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery6(void* networkConfiguration, char* hostname, int port, char* protocolTag, long long queryType)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery5(void* networkConfiguration, void* requestUrl, long long queryType)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), *static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery7(void* networkConfiguration, unsigned short bindPort, char* protocolTag, long long queryType)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), bindPort, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery8(void* other)
{
	return new QNetworkProxyQuery(*static_cast<QNetworkProxyQuery*>(other));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, long long queryType)
{
	return new QNetworkProxyQuery(QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery2(void* requestUrl, long long queryType)
{
	return new QNetworkProxyQuery(*static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery4(unsigned short bindPort, char* protocolTag, long long queryType)
{
	return new QNetworkProxyQuery(bindPort, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

int QNetworkProxyQuery_LocalPort(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->localPort();
}

void* QNetworkProxyQuery_NetworkConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QNetworkProxyQuery*>(ptr)->networkConfiguration());
}

struct QtNetwork_PackedString QNetworkProxyQuery_PeerHostName(void* ptr)
{
	return ({ QByteArray t878927 = static_cast<QNetworkProxyQuery*>(ptr)->peerHostName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t878927.prepend("WHITESPACE").constData()+10), t878927.size()-10 }; });
}

int QNetworkProxyQuery_PeerPort(void* ptr)
{
	return static_cast<QNetworkProxyQuery*>(ptr)->peerPort();
}

struct QtNetwork_PackedString QNetworkProxyQuery_ProtocolTag(void* ptr)
{
	return ({ QByteArray teab311 = static_cast<QNetworkProxyQuery*>(ptr)->protocolTag().toUtf8(); QtNetwork_PackedString { const_cast<char*>(teab311.prepend("WHITESPACE").constData()+10), teab311.size()-10 }; });
}

long long QNetworkProxyQuery_QueryType(void* ptr)
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
	void setSslConfigurationImplementation(const QSslConfiguration & configuration) { callbackQNetworkReply_SetSslConfigurationImplementation(this, const_cast<QSslConfiguration*>(&configuration)); };
	void sslConfigurationImplementation(QSslConfiguration & configuration) const { callbackQNetworkReply_SslConfigurationImplementation(const_cast<MyQNetworkReply*>(this), static_cast<QSslConfiguration*>(&configuration)); };
	void abort() { callbackQNetworkReply_Abort(this); };
	void close() { callbackQNetworkReply_Close(this); };
	void Signal_DownloadProgress(qint64 bytesReceived, qint64 bytesTotal) { callbackQNetworkReply_DownloadProgress(this, bytesReceived, bytesTotal); };
	void Signal_Encrypted() { callbackQNetworkReply_Encrypted(this); };
	void Signal_Error2(QNetworkReply::NetworkError code) { callbackQNetworkReply_Error2(this, code); };
	void Signal_Finished() { callbackQNetworkReply_Finished(this); };
	void ignoreSslErrors() { callbackQNetworkReply_IgnoreSslErrors(this); };
	void Signal_MetaDataChanged() { callbackQNetworkReply_MetaDataChanged(this); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkReply_PreSharedKeyAuthenticationRequired(this, authenticator); };
	void Signal_Redirected(const QUrl & url) { callbackQNetworkReply_Redirected(this, const_cast<QUrl*>(&url)); };
	void setReadBufferSize(qint64 size) { callbackQNetworkReply_SetReadBufferSize(this, size); };
	void Signal_UploadProgress(qint64 bytesSent, qint64 bytesTotal) { callbackQNetworkReply_UploadProgress(this, bytesSent, bytesTotal); };
	bool atEnd() const { return callbackQNetworkReply_AtEnd(const_cast<MyQNetworkReply*>(this)) != 0; };
	qint64 bytesAvailable() const { return callbackQNetworkReply_BytesAvailable(const_cast<MyQNetworkReply*>(this)); };
	qint64 bytesToWrite() const { return callbackQNetworkReply_BytesToWrite(const_cast<MyQNetworkReply*>(this)); };
	bool canReadLine() const { return callbackQNetworkReply_CanReadLine(const_cast<MyQNetworkReply*>(this)) != 0; };
	bool isSequential() const { return callbackQNetworkReply_IsSequential(const_cast<MyQNetworkReply*>(this)) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQNetworkReply_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQNetworkReply_Pos(const_cast<MyQNetworkReply*>(this)); };
	qint64 readLineData(char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { data, maxSize };return callbackQNetworkReply_ReadLineData(this, dataPacked, maxSize); };
	bool reset() { return callbackQNetworkReply_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQNetworkReply_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQNetworkReply_Size(const_cast<MyQNetworkReply*>(this)); };
	bool waitForBytesWritten(int msecs) { return callbackQNetworkReply_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQNetworkReply_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 maxSize) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), maxSize };return callbackQNetworkReply_WriteData(this, dataPacked, maxSize); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkReply_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNetworkReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkReply_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkReply_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNetworkReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkReply_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkReply_MetaObject(const_cast<MyQNetworkReply*>(this))); };
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

void* QNetworkReply_Attribute(void* ptr, long long code)
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

long long QNetworkReply_Error(void* ptr)
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

char QNetworkReply_HasRawHeader(void* ptr, void* headerName)
{
	return static_cast<QNetworkReply*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkReply_Header(void* ptr, long long header)
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

char QNetworkReply_IsFinished(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isFinished();
}

char QNetworkReply_IsRunning(void* ptr)
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

long long QNetworkReply_Operation(void* ptr)
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

void* QNetworkReply_RawHeader(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QNetworkReply*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

struct QtNetwork_PackedList QNetworkReply_RawHeaderList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QNetworkReply*>(ptr)->rawHeaderList()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QNetworkReply_ReadBufferSize(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->readBufferSize();
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

void QNetworkReply_SetAttribute(void* ptr, long long code, void* value)
{
	static_cast<QNetworkReply*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkReply_SetError(void* ptr, long long errorCode, char* errorString)
{
	static_cast<QNetworkReply*>(ptr)->setError(static_cast<QNetworkReply::NetworkError>(errorCode), QString(errorString));
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
	static_cast<QNetworkReply*>(ptr)->uploadProgress(bytesSent, bytesTotal);
}

void* QNetworkReply_Url(void* ptr)
{
	return new QUrl(static_cast<QNetworkReply*>(ptr)->url());
}

void QNetworkReply_DestroyQNetworkReply(void* ptr)
{
	static_cast<QNetworkReply*>(ptr)->~QNetworkReply();
}

void* QNetworkReply_rawHeaderList_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

char QNetworkReply_AtEnd(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->atEnd();
}

char QNetworkReply_AtEndDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::atEnd();
}

long long QNetworkReply_BytesAvailable(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->bytesAvailable();
}

long long QNetworkReply_BytesAvailableDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::bytesAvailable();
}

long long QNetworkReply_BytesToWrite(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->bytesToWrite();
}

long long QNetworkReply_BytesToWriteDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::bytesToWrite();
}

char QNetworkReply_CanReadLine(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->canReadLine();
}

char QNetworkReply_CanReadLineDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::canReadLine();
}

char QNetworkReply_IsSequential(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->isSequential();
}

char QNetworkReply_IsSequentialDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::isSequential();
}

char QNetworkReply_Open(void* ptr, long long mode)
{
	return static_cast<QNetworkReply*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QNetworkReply_OpenDefault(void* ptr, long long mode)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QNetworkReply_Pos(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->pos();
}

long long QNetworkReply_PosDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::pos();
}

long long QNetworkReply_ReadLineData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QNetworkReply*>(ptr)->readLineData(data, maxSize);
}

long long QNetworkReply_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::readLineData(data, maxSize);
}

char QNetworkReply_Reset(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->reset();
}

char QNetworkReply_ResetDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::reset();
}

char QNetworkReply_Seek(void* ptr, long long pos)
{
	return static_cast<QNetworkReply*>(ptr)->seek(pos);
}

char QNetworkReply_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::seek(pos);
}

long long QNetworkReply_Size(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->size();
}

long long QNetworkReply_SizeDefault(void* ptr)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::size();
}

char QNetworkReply_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->waitForBytesWritten(msecs);
}

char QNetworkReply_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::waitForBytesWritten(msecs);
}

char QNetworkReply_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->waitForReadyRead(msecs);
}

char QNetworkReply_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::waitForReadyRead(msecs);
}

long long QNetworkReply_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QNetworkReply*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QNetworkReply_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::writeData(const_cast<const char*>(data), maxSize);
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

char QNetworkReply_Event(void* ptr, void* e)
{
	return static_cast<QNetworkReply*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNetworkReply_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkReply*>(ptr)->QNetworkReply::event(static_cast<QEvent*>(e));
}

char QNetworkReply_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkReply*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNetworkReply_EventFilterDefault(void* ptr, void* watched, void* event)
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

void* QNetworkRequest_Attribute(void* ptr, long long code, void* defaultValue)
{
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

char QNetworkRequest_HasRawHeader(void* ptr, void* headerName)
{
	return static_cast<QNetworkRequest*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkRequest_Header(void* ptr, long long header)
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

long long QNetworkRequest_Priority(void* ptr)
{
	return static_cast<QNetworkRequest*>(ptr)->priority();
}

void* QNetworkRequest_RawHeader(void* ptr, void* headerName)
{
	return new QByteArray(static_cast<QNetworkRequest*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

struct QtNetwork_PackedList QNetworkRequest_RawHeaderList(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QNetworkRequest*>(ptr)->rawHeaderList()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

void* QNetworkRequest_rawHeaderList_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

class MyQNetworkSession: public QNetworkSession
{
public:
	MyQNetworkSession(const QNetworkConfiguration &connectionConfig, QObject *parent) : QNetworkSession(connectionConfig, parent) {};
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
	void timerEvent(QTimerEvent * event) { callbackQNetworkSession_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNetworkSession_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNetworkSession_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNetworkSession_CustomEvent(this, event); };
	void deleteLater() { callbackQNetworkSession_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNetworkSession_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNetworkSession_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNetworkSession_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNetworkSession_MetaObject(const_cast<MyQNetworkSession*>(this))); };
};

void* QNetworkSession_NewQNetworkSession(void* connectionConfig, void* parent)
{
	return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QObject*>(parent));
}

void QNetworkSession_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "accept");
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

void QNetworkSession_Error2(void* ptr, long long error)
{
	static_cast<QNetworkSession*>(ptr)->error(static_cast<QNetworkSession::SessionError>(error));
}

long long QNetworkSession_Error(void* ptr)
{
	return static_cast<QNetworkSession*>(ptr)->error();
}

struct QtNetwork_PackedString QNetworkSession_ErrorString(void* ptr)
{
	return ({ QByteArray t57e370 = static_cast<QNetworkSession*>(ptr)->errorString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t57e370.prepend("WHITESPACE").constData()+10), t57e370.size()-10 }; });
}

void QNetworkSession_Ignore(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "ignore");
}

void* QNetworkSession_Interface(void* ptr)
{
	return new QNetworkInterface(static_cast<QNetworkSession*>(ptr)->interface());
}

char QNetworkSession_IsOpen(void* ptr)
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

void QNetworkSession_PreferredConfigurationChanged(void* ptr, void* config, char isSeamless)
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

long long QNetworkSession_State(void* ptr)
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

void QNetworkSession_StateChanged(void* ptr, long long state)
{
	static_cast<QNetworkSession*>(ptr)->stateChanged(static_cast<QNetworkSession::State>(state));
}

void QNetworkSession_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "stop");
}

long long QNetworkSession_UsagePolicies(void* ptr)
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

void QNetworkSession_UsagePoliciesChanged(void* ptr, long long usagePolicies)
{
	static_cast<QNetworkSession*>(ptr)->usagePoliciesChanged(static_cast<QNetworkSession::UsagePolicy>(usagePolicies));
}

char QNetworkSession_WaitForOpened(void* ptr, int msecs)
{
	return static_cast<QNetworkSession*>(ptr)->waitForOpened(msecs);
}

void QNetworkSession_DestroyQNetworkSession(void* ptr)
{
	static_cast<QNetworkSession*>(ptr)->~QNetworkSession();
}

void QNetworkSession_DestroyQNetworkSessionDefault(void* ptr)
{

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

char QNetworkSession_Event(void* ptr, void* e)
{
	return static_cast<QNetworkSession*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNetworkSession_EventDefault(void* ptr, void* e)
{
	return static_cast<QNetworkSession*>(ptr)->QNetworkSession::event(static_cast<QEvent*>(e));
}

char QNetworkSession_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNetworkSession*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNetworkSession_EventFilterDefault(void* ptr, void* watched, void* event)
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

int QSsl_TlsV1_1_Type()
{
	return QSsl::TlsV1_1;
}

int QSsl_TlsV1_2_Type()
{
	return QSsl::TlsV1_2;
}

int QSsl_AnyProtocol_Type()
{
	return QSsl::AnyProtocol;
}

int QSsl_TlsV1SslV3_Type()
{
	return QSsl::TlsV1SslV3;
}

int QSsl_SecureProtocols_Type()
{
	return QSsl::SecureProtocols;
}

int QSsl_TlsV1_0OrLater_Type()
{
	return QSsl::TlsV1_0OrLater;
}

int QSsl_TlsV1_1OrLater_Type()
{
	return QSsl::TlsV1_1OrLater;
}

int QSsl_TlsV1_2OrLater_Type()
{
	return QSsl::TlsV1_2OrLater;
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

void QSslCertificate_Clear(void* ptr)
{
	static_cast<QSslCertificate*>(ptr)->clear();
}

void* QSslCertificate_Digest(void* ptr, long long algorithm)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->digest(static_cast<QCryptographicHash::Algorithm>(algorithm)));
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_FromData(void* data, long long format)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslCertificate::fromData(*static_cast<QByteArray*>(data), static_cast<QSsl::EncodingFormat>(format))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_FromDevice(void* device, long long format)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslCertificate::fromDevice(static_cast<QIODevice*>(device), static_cast<QSsl::EncodingFormat>(format))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslCertificate_QSslCertificate_FromPath(char* path, long long format, long long syntax)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslCertificate::fromPath(QString(path), static_cast<QSsl::EncodingFormat>(format), static_cast<QRegExp::PatternSyntax>(syntax))); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QSslCertificate_IsBlacklisted(void* ptr)
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

struct QtNetwork_PackedList QSslCertificate_Extensions(void* ptr)
{
	return ({ QList<QSslCertificateExtension>* tmpValue = new QList<QSslCertificateExtension>(static_cast<QSslCertificate*>(ptr)->extensions()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QSslCertificate_IsNull(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isNull();
}

char QSslCertificate_IsSelfSigned(void* ptr)
{
	return static_cast<QSslCertificate*>(ptr)->isSelfSigned();
}

struct QtNetwork_PackedString QSslCertificate_IssuerInfo(void* ptr, long long subject)
{
	return ({ QByteArray t768c47 = static_cast<QSslCertificate*>(ptr)->issuerInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(t768c47.prepend("WHITESPACE").constData()+10), t768c47.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_IssuerInfo2(void* ptr, void* attribute)
{
	return ({ QByteArray tc820f1 = static_cast<QSslCertificate*>(ptr)->issuerInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc820f1.prepend("WHITESPACE").constData()+10), tc820f1.size()-10 }; });
}

struct QtNetwork_PackedList QSslCertificate_IssuerInfoAttributes(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QSslCertificate*>(ptr)->issuerInfoAttributes()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslCertificate_PublicKey(void* ptr)
{
	return new QSslKey(static_cast<QSslCertificate*>(ptr)->publicKey());
}

void* QSslCertificate_SerialNumber(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->serialNumber());
}

struct QtNetwork_PackedString QSslCertificate_SubjectInfo(void* ptr, long long subject)
{
	return ({ QByteArray tee2197 = static_cast<QSslCertificate*>(ptr)->subjectInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tee2197.prepend("WHITESPACE").constData()+10), tee2197.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificate_SubjectInfo2(void* ptr, void* attribute)
{
	return ({ QByteArray tc13a73 = static_cast<QSslCertificate*>(ptr)->subjectInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8(); QtNetwork_PackedString { const_cast<char*>(tc13a73.prepend("WHITESPACE").constData()+10), tc13a73.size()-10 }; });
}

struct QtNetwork_PackedList QSslCertificate_SubjectInfoAttributes(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QSslCertificate*>(ptr)->subjectInfoAttributes()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslCertificate_ToDer(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toDer());
}

void* QSslCertificate_ToPem(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toPem());
}

struct QtNetwork_PackedString QSslCertificate_ToText(void* ptr)
{
	return ({ QByteArray t52ef8a = static_cast<QSslCertificate*>(ptr)->toText().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t52ef8a.prepend("WHITESPACE").constData()+10), t52ef8a.size()-10 }; });
}

void* QSslCertificate_Version(void* ptr)
{
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->version());
}

void* QSslCertificate_fromData_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslCertificate_fromDevice_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslCertificate_fromPath_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslCertificate_extensions_atList(void* ptr, int i)
{
	return new QSslCertificateExtension(static_cast<QList<QSslCertificateExtension>*>(ptr)->at(i));
}

void* QSslCertificate_issuerInfoAttributes_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void* QSslCertificate_subjectInfoAttributes_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void* QSslCertificateExtension_NewQSslCertificateExtension()
{
	return new QSslCertificateExtension();
}

void* QSslCertificateExtension_NewQSslCertificateExtension2(void* other)
{
	return new QSslCertificateExtension(*static_cast<QSslCertificateExtension*>(other));
}

char QSslCertificateExtension_IsCritical(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->isCritical();
}

char QSslCertificateExtension_IsSupported(void* ptr)
{
	return static_cast<QSslCertificateExtension*>(ptr)->isSupported();
}

struct QtNetwork_PackedString QSslCertificateExtension_Name(void* ptr)
{
	return ({ QByteArray t994389 = static_cast<QSslCertificateExtension*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t994389.prepend("WHITESPACE").constData()+10), t994389.size()-10 }; });
}

struct QtNetwork_PackedString QSslCertificateExtension_Oid(void* ptr)
{
	return ({ QByteArray t615506 = static_cast<QSslCertificateExtension*>(ptr)->oid().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t615506.prepend("WHITESPACE").constData()+10), t615506.size()-10 }; });
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

void* QSslCipher_NewQSslCipher3(char* name, long long protocol)
{
	return new QSslCipher(QString(name), static_cast<QSsl::SslProtocol>(protocol));
}

struct QtNetwork_PackedString QSslCipher_AuthenticationMethod(void* ptr)
{
	return ({ QByteArray tfc1f5a = static_cast<QSslCipher*>(ptr)->authenticationMethod().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfc1f5a.prepend("WHITESPACE").constData()+10), tfc1f5a.size()-10 }; });
}

struct QtNetwork_PackedString QSslCipher_EncryptionMethod(void* ptr)
{
	return ({ QByteArray ta39f95 = static_cast<QSslCipher*>(ptr)->encryptionMethod().toUtf8(); QtNetwork_PackedString { const_cast<char*>(ta39f95.prepend("WHITESPACE").constData()+10), ta39f95.size()-10 }; });
}

char QSslCipher_IsNull(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->isNull();
}

struct QtNetwork_PackedString QSslCipher_KeyExchangeMethod(void* ptr)
{
	return ({ QByteArray tfbfb25 = static_cast<QSslCipher*>(ptr)->keyExchangeMethod().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tfbfb25.prepend("WHITESPACE").constData()+10), tfbfb25.size()-10 }; });
}

struct QtNetwork_PackedString QSslCipher_Name(void* ptr)
{
	return ({ QByteArray t9ef3a9 = static_cast<QSslCipher*>(ptr)->name().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t9ef3a9.prepend("WHITESPACE").constData()+10), t9ef3a9.size()-10 }; });
}

long long QSslCipher_Protocol(void* ptr)
{
	return static_cast<QSslCipher*>(ptr)->protocol();
}

struct QtNetwork_PackedString QSslCipher_ProtocolString(void* ptr)
{
	return ({ QByteArray t99c307 = static_cast<QSslCipher*>(ptr)->protocolString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t99c307.prepend("WHITESPACE").constData()+10), t99c307.size()-10 }; });
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

struct QtNetwork_PackedList QSslConfiguration_AllowedNextProtocols(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QSslConfiguration*>(ptr)->allowedNextProtocols()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_CaCertificates(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslConfiguration*>(ptr)->caCertificates()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtNetwork_PackedList QSslConfiguration_Ciphers(void* ptr)
{
	return ({ QList<QSslCipher>* tmpValue = new QList<QSslCipher>(static_cast<QSslConfiguration*>(ptr)->ciphers()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslConfiguration_QSslConfiguration_DefaultConfiguration()
{
	return new QSslConfiguration(QSslConfiguration::defaultConfiguration());
}

void* QSslConfiguration_EphemeralServerKey(void* ptr)
{
	return new QSslKey(static_cast<QSslConfiguration*>(ptr)->ephemeralServerKey());
}

char QSslConfiguration_IsNull(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->isNull();
}

void* QSslConfiguration_LocalCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslConfiguration*>(ptr)->localCertificate());
}

struct QtNetwork_PackedList QSslConfiguration_LocalCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslConfiguration*>(ptr)->localCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSslConfiguration_NextNegotiatedProtocol(void* ptr)
{
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->nextNegotiatedProtocol());
}

long long QSslConfiguration_NextProtocolNegotiationStatus(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->nextProtocolNegotiationStatus();
}

void* QSslConfiguration_PeerCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslConfiguration*>(ptr)->peerCertificate());
}

struct QtNetwork_PackedList QSslConfiguration_PeerCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslConfiguration*>(ptr)->peerCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

int QSslConfiguration_PeerVerifyDepth(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyDepth();
}

long long QSslConfiguration_PeerVerifyMode(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyMode();
}

void* QSslConfiguration_PrivateKey(void* ptr)
{
	return new QSslKey(static_cast<QSslConfiguration*>(ptr)->privateKey());
}

long long QSslConfiguration_Protocol(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->protocol();
}

void* QSslConfiguration_SessionCipher(void* ptr)
{
	return new QSslCipher(static_cast<QSslConfiguration*>(ptr)->sessionCipher());
}

long long QSslConfiguration_SessionProtocol(void* ptr)
{
	return static_cast<QSslConfiguration*>(ptr)->sessionProtocol();
}

void* QSslConfiguration_SessionTicket(void* ptr)
{
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->sessionTicket());
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

void QSslConfiguration_SetPeerVerifyMode(void* ptr, long long mode)
{
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
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

struct QtNetwork_PackedList QSslConfiguration_QSslConfiguration_SupportedCiphers()
{
	return ({ QList<QSslCipher>* tmpValue = new QList<QSslCipher>(QSslConfiguration::supportedCiphers()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

void QSslConfiguration_Swap(void* ptr, void* other)
{
	static_cast<QSslConfiguration*>(ptr)->swap(*static_cast<QSslConfiguration*>(other));
}

struct QtNetwork_PackedList QSslConfiguration_QSslConfiguration_SystemCaCertificates()
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(QSslConfiguration::systemCaCertificates()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

char QSslConfiguration_TestSslOption(void* ptr, long long option)
{
	return static_cast<QSslConfiguration*>(ptr)->testSslOption(static_cast<QSsl::SslOption>(option));
}

void QSslConfiguration_DestroyQSslConfiguration(void* ptr)
{
	static_cast<QSslConfiguration*>(ptr)->~QSslConfiguration();
}

void* QSslConfiguration_allowedNextProtocols_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void* QSslConfiguration_caCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslConfiguration_ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(static_cast<QList<QSslCipher>*>(ptr)->at(i));
}

void* QSslConfiguration_ellipticCurves_atList(void* ptr, int i)
{
	return new QSslEllipticCurve(static_cast<QVector<QSslEllipticCurve>*>(ptr)->at(i));
}

void* QSslConfiguration_localCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslConfiguration_peerCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslConfiguration_supportedCiphers_atList(void* ptr, int i)
{
	return new QSslCipher(static_cast<QList<QSslCipher>*>(ptr)->at(i));
}

void* QSslConfiguration_supportedEllipticCurves_atList(void* ptr, int i)
{
	return new QSslEllipticCurve(static_cast<QVector<QSslEllipticCurve>*>(ptr)->at(i));
}

void* QSslConfiguration_systemCaCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslEllipticCurve_NewQSslEllipticCurve()
{
	return new QSslEllipticCurve();
}

char QSslEllipticCurve_IsValid(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->isValid();
}

void* QSslEllipticCurve_QSslEllipticCurve_FromLongName(char* name)
{
	return new QSslEllipticCurve(QSslEllipticCurve::fromLongName(QString(name)));
}

void* QSslEllipticCurve_QSslEllipticCurve_FromShortName(char* name)
{
	return new QSslEllipticCurve(QSslEllipticCurve::fromShortName(QString(name)));
}

char QSslEllipticCurve_IsTlsNamedCurve(void* ptr)
{
	return static_cast<QSslEllipticCurve*>(ptr)->isTlsNamedCurve();
}

struct QtNetwork_PackedString QSslEllipticCurve_LongName(void* ptr)
{
	return ({ QByteArray t85b564 = static_cast<QSslEllipticCurve*>(ptr)->longName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t85b564.prepend("WHITESPACE").constData()+10), t85b564.size()-10 }; });
}

struct QtNetwork_PackedString QSslEllipticCurve_ShortName(void* ptr)
{
	return ({ QByteArray tb6e6fc = static_cast<QSslEllipticCurve*>(ptr)->shortName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tb6e6fc.prepend("WHITESPACE").constData()+10), tb6e6fc.size()-10 }; });
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

long long QSslKey_Algorithm(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->algorithm();
}

void QSslKey_Clear(void* ptr)
{
	static_cast<QSslKey*>(ptr)->clear();
}

char QSslKey_IsNull(void* ptr)
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

void* QSslKey_ToDer(void* ptr, void* passPhrase)
{
	return new QByteArray(static_cast<QSslKey*>(ptr)->toDer(*static_cast<QByteArray*>(passPhrase)));
}

void* QSslKey_ToPem(void* ptr, void* passPhrase)
{
	return new QByteArray(static_cast<QSslKey*>(ptr)->toPem(*static_cast<QByteArray*>(passPhrase)));
}

long long QSslKey_Type(void* ptr)
{
	return static_cast<QSslKey*>(ptr)->type();
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

void* QSslPreSharedKeyAuthenticator_Identity(void* ptr)
{
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identity());
}

void* QSslPreSharedKeyAuthenticator_IdentityHint(void* ptr)
{
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identityHint());
}

int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(void* ptr)
{
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumIdentityLength();
}

int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(void* ptr)
{
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumPreSharedKeyLength();
}

void* QSslPreSharedKeyAuthenticator_PreSharedKey(void* ptr)
{
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->preSharedKey());
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

class MyQSslSocket: public QSslSocket
{
public:
	MyQSslSocket(QObject *parent) : QSslSocket(parent) {};
	bool atEnd() const { return callbackQSslSocket_AtEnd(const_cast<MyQSslSocket*>(this)) != 0; };
	qint64 bytesAvailable() const { return callbackQSslSocket_BytesAvailable(const_cast<MyQSslSocket*>(this)); };
	qint64 bytesToWrite() const { return callbackQSslSocket_BytesToWrite(const_cast<MyQSslSocket*>(this)); };
	bool canReadLine() const { return callbackQSslSocket_CanReadLine(const_cast<MyQSslSocket*>(this)) != 0; };
	void close() { callbackQSslSocket_Close(this); };
	void Signal_Encrypted() { callbackQSslSocket_Encrypted(this); };
	void Signal_EncryptedBytesWritten(qint64 written) { callbackQSslSocket_EncryptedBytesWritten(this, written); };
	void ignoreSslErrors() { callbackQSslSocket_IgnoreSslErrors(this); };
	void Signal_ModeChanged(QSslSocket::SslMode mode) { callbackQSslSocket_ModeChanged(this, mode); };
	void Signal_PeerVerifyError(const QSslError & error) { callbackQSslSocket_PeerVerifyError(this, const_cast<QSslError*>(&error)); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQSslSocket_PreSharedKeyAuthenticationRequired(this, authenticator); };
	void resume() { callbackQSslSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQSslSocket_SetReadBufferSize(this, size); };
	
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQSslSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQSslSocket_SocketOption(this, option)); };
	void startClientEncryption() { callbackQSslSocket_StartClientEncryption(this); };
	void startServerEncryption() { callbackQSslSocket_StartServerEncryption(this); };
	bool waitForBytesWritten(int msecs) { return callbackQSslSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQSslSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQSslSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQSslSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 len) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), len };return callbackQSslSocket_WriteData(this, dataPacked, len); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQSslSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQSslSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void disconnectFromHost() { callbackQSslSocket_DisconnectFromHost(this); };
	bool isSequential() const { return callbackQSslSocket_IsSequential(const_cast<MyQSslSocket*>(this)) != 0; };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQSslSocket_ReadLineData(this, dataPacked, maxlen); };
	
	bool open(QIODevice::OpenMode mode) { return callbackQSslSocket_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQSslSocket_Pos(const_cast<MyQSslSocket*>(this)); };
	bool reset() { return callbackQSslSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQSslSocket_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQSslSocket_Size(const_cast<MyQSslSocket*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQSslSocket_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSslSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSslSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSslSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQSslSocket_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSslSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSslSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSslSocket_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSslSocket_MetaObject(const_cast<MyQSslSocket*>(this))); };
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

char QSslSocket_AddCaCertificates(void* ptr, char* path, long long format, long long syntax)
{
	return static_cast<QSslSocket*>(ptr)->addCaCertificates(QString(path), static_cast<QSsl::EncodingFormat>(format), static_cast<QRegExp::PatternSyntax>(syntax));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificate(void* certificate)
{
	QSslSocket::addDefaultCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

char QSslSocket_QSslSocket_AddDefaultCaCertificates(char* path, long long encoding, long long syntax)
{
	return QSslSocket::addDefaultCaCertificates(QString(path), static_cast<QSsl::EncodingFormat>(encoding), static_cast<QRegExp::PatternSyntax>(syntax));
}

char QSslSocket_AtEnd(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->atEnd();
}

char QSslSocket_AtEndDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::atEnd();
}

long long QSslSocket_BytesAvailable(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->bytesAvailable();
}

long long QSslSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::bytesAvailable();
}

long long QSslSocket_BytesToWrite(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->bytesToWrite();
}

long long QSslSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::bytesToWrite();
}

char QSslSocket_CanReadLine(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->canReadLine();
}

char QSslSocket_CanReadLineDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::canReadLine();
}

void QSslSocket_Close(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->close();
}

void QSslSocket_CloseDefault(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::close();
}

void QSslSocket_ConnectToHostEncrypted(void* ptr, char* hostName, unsigned short port, long long mode, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->connectToHostEncrypted(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QSslSocket_ConnectToHostEncrypted2(void* ptr, char* hostName, unsigned short port, char* sslPeerName, long long mode, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->connectToHostEncrypted(QString(hostName), port, QString(sslPeerName), static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
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
	return static_cast<QSslSocket*>(ptr)->encryptedBytesAvailable();
}

long long QSslSocket_EncryptedBytesToWrite(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->encryptedBytesToWrite();
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

char QSslSocket_Flush(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->flush();
}

void QSslSocket_IgnoreSslErrors(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "ignoreSslErrors");
}

char QSslSocket_IsEncrypted(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->isEncrypted();
}

void* QSslSocket_LocalCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslSocket*>(ptr)->localCertificate());
}

struct QtNetwork_PackedList QSslSocket_LocalCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslSocket*>(ptr)->localCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long long QSslSocket_Mode(void* ptr)
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

void QSslSocket_ModeChanged(void* ptr, long long mode)
{
	static_cast<QSslSocket*>(ptr)->modeChanged(static_cast<QSslSocket::SslMode>(mode));
}

void* QSslSocket_PeerCertificate(void* ptr)
{
	return new QSslCertificate(static_cast<QSslSocket*>(ptr)->peerCertificate());
}

struct QtNetwork_PackedList QSslSocket_PeerCertificateChain(void* ptr)
{
	return ({ QList<QSslCertificate>* tmpValue = new QList<QSslCertificate>(static_cast<QSslSocket*>(ptr)->peerCertificateChain()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
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

long long QSslSocket_PeerVerifyMode(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->peerVerifyMode();
}

struct QtNetwork_PackedString QSslSocket_PeerVerifyName(void* ptr)
{
	return ({ QByteArray tefa5dd = static_cast<QSslSocket*>(ptr)->peerVerifyName().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tefa5dd.prepend("WHITESPACE").constData()+10), tefa5dd.size()-10 }; });
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

long long QSslSocket_Protocol(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->protocol();
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

long long QSslSocket_SessionProtocol(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->sessionProtocol();
}

void QSslSocket_SetLocalCertificate(void* ptr, void* certificate)
{
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_SetLocalCertificate2(void* ptr, char* path, long long format)
{
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(QString(path), static_cast<QSsl::EncodingFormat>(format));
}

void QSslSocket_SetPeerVerifyDepth(void* ptr, int depth)
{
	static_cast<QSslSocket*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslSocket_SetPeerVerifyMode(void* ptr, long long mode)
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

void QSslSocket_SetPrivateKey2(void* ptr, char* fileName, long long algorithm, long long format, void* passPhrase)
{
	static_cast<QSslSocket*>(ptr)->setPrivateKey(QString(fileName), static_cast<QSsl::KeyAlgorithm>(algorithm), static_cast<QSsl::EncodingFormat>(format), *static_cast<QByteArray*>(passPhrase));
}

void QSslSocket_SetProtocol(void* ptr, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->setProtocol(static_cast<QSsl::SslProtocol>(protocol));
}

void QSslSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QSslSocket*>(ptr)->setReadBufferSize(size);
}

void QSslSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::setReadBufferSize(size);
}





void QSslSocket_SetSocketOption(void* ptr, long long option, void* value)
{
	static_cast<QSslSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSocketOptionDefault(void* ptr, long long option, void* value)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSslConfiguration(void* ptr, void* configuration)
{
	static_cast<QSslSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void* QSslSocket_SocketOption(void* ptr, long long option)
{
	return new QVariant(static_cast<QSslSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QSslSocket_SocketOptionDefault(void* ptr, long long option)
{
	return new QVariant(static_cast<QSslSocket*>(ptr)->QSslSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QSslSocket_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QSslSocket*>(ptr)->sslConfiguration());
}

struct QtNetwork_PackedList QSslSocket_SslErrors(void* ptr)
{
	return ({ QList<QSslError>* tmpValue = new QList<QSslError>(static_cast<QSslSocket*>(ptr)->sslErrors()); QtNetwork_PackedList { tmpValue, tmpValue->size() }; });
}

long QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()
{
	return QSslSocket::sslLibraryBuildVersionNumber();
}

struct QtNetwork_PackedString QSslSocket_QSslSocket_SslLibraryBuildVersionString()
{
	return ({ QByteArray t55b90e = QSslSocket::sslLibraryBuildVersionString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(t55b90e.prepend("WHITESPACE").constData()+10), t55b90e.size()-10 }; });
}

long QSslSocket_QSslSocket_SslLibraryVersionNumber()
{
	return QSslSocket::sslLibraryVersionNumber();
}

struct QtNetwork_PackedString QSslSocket_QSslSocket_SslLibraryVersionString()
{
	return ({ QByteArray tdd64b4 = QSslSocket::sslLibraryVersionString().toUtf8(); QtNetwork_PackedString { const_cast<char*>(tdd64b4.prepend("WHITESPACE").constData()+10), tdd64b4.size()-10 }; });
}

void QSslSocket_StartClientEncryption(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startClientEncryption");
}

void QSslSocket_StartServerEncryption(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startServerEncryption");
}

char QSslSocket_QSslSocket_SupportsSsl()
{
	return QSslSocket::supportsSsl();
}

char QSslSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForBytesWritten(msecs);
}

char QSslSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForBytesWritten(msecs);
}

char QSslSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForConnected(msecs);
}

char QSslSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForConnected(msecs);
}

char QSslSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForDisconnected(msecs);
}

char QSslSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForDisconnected(msecs);
}

char QSslSocket_WaitForEncrypted(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForEncrypted(msecs);
}

char QSslSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->waitForReadyRead(msecs);
}

char QSslSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::waitForReadyRead(msecs);
}

long long QSslSocket_WriteData(void* ptr, char* data, long long len)
{
	return static_cast<QSslSocket*>(ptr)->writeData(const_cast<const char*>(data), len);
}

long long QSslSocket_WriteDataDefault(void* ptr, char* data, long long len)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::writeData(const_cast<const char*>(data), len);
}

void QSslSocket_DestroyQSslSocket(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->~QSslSocket();
}

void* QSslSocket_caCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslSocket_ciphers_atList(void* ptr, int i)
{
	return new QSslCipher(static_cast<QList<QSslCipher>*>(ptr)->at(i));
}

void* QSslSocket_defaultCaCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslSocket_defaultCiphers_atList(void* ptr, int i)
{
	return new QSslCipher(static_cast<QList<QSslCipher>*>(ptr)->at(i));
}

void* QSslSocket_localCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslSocket_peerCertificateChain_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void* QSslSocket_sslErrors_atList(void* ptr, int i)
{
	return new QSslError(static_cast<QList<QSslError>*>(ptr)->at(i));
}

void* QSslSocket_supportedCiphers_atList(void* ptr, int i)
{
	return new QSslCipher(static_cast<QList<QSslCipher>*>(ptr)->at(i));
}

void* QSslSocket_systemCaCertificates_atList(void* ptr, int i)
{
	return new QSslCertificate(static_cast<QList<QSslCertificate>*>(ptr)->at(i));
}

void QSslSocket_ConnectToHost2(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QSslSocket*>(ptr)->connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QSslSocket_ConnectToHost2Default(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QSslSocket_ConnectToHost(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QSslSocket_ConnectToHostDefault(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QSslSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->disconnectFromHost();
}

void QSslSocket_DisconnectFromHostDefault(void* ptr)
{
	static_cast<QSslSocket*>(ptr)->QSslSocket::disconnectFromHost();
}

char QSslSocket_IsSequential(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->isSequential();
}

char QSslSocket_IsSequentialDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::isSequential();
}

long long QSslSocket_ReadLineData(void* ptr, char* data, long long maxlen)
{
	return static_cast<QSslSocket*>(ptr)->readLineData(data, maxlen);
}

long long QSslSocket_ReadLineDataDefault(void* ptr, char* data, long long maxlen)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::readLineData(data, maxlen);
}





char QSslSocket_Open(void* ptr, long long mode)
{
	return static_cast<QSslSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QSslSocket_OpenDefault(void* ptr, long long mode)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QSslSocket_Pos(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->pos();
}

long long QSslSocket_PosDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::pos();
}

char QSslSocket_Reset(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->reset();
}

char QSslSocket_ResetDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::reset();
}

char QSslSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QSslSocket*>(ptr)->seek(pos);
}

char QSslSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::seek(pos);
}

long long QSslSocket_Size(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->size();
}

long long QSslSocket_SizeDefault(void* ptr)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::size();
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

char QSslSocket_Event(void* ptr, void* e)
{
	return static_cast<QSslSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSslSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QSslSocket*>(ptr)->QSslSocket::event(static_cast<QEvent*>(e));
}

char QSslSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSslSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSslSocket_EventFilterDefault(void* ptr, void* watched, void* event)
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
	void Signal_AcceptError(QAbstractSocket::SocketError socketError) { callbackQTcpServer_AcceptError(this, socketError); };
	bool hasPendingConnections() const { return callbackQTcpServer_HasPendingConnections(const_cast<MyQTcpServer*>(this)) != 0; };
	
	void Signal_NewConnection() { callbackQTcpServer_NewConnection(this); };
	QTcpSocket * nextPendingConnection() { return static_cast<QTcpSocket*>(callbackQTcpServer_NextPendingConnection(this)); };
	 ~MyQTcpServer() { callbackQTcpServer_DestroyQTcpServer(this); };
	void timerEvent(QTimerEvent * event) { callbackQTcpServer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQTcpServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTcpServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTcpServer_CustomEvent(this, event); };
	void deleteLater() { callbackQTcpServer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTcpServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTcpServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTcpServer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTcpServer_MetaObject(const_cast<MyQTcpServer*>(this))); };
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

char QTcpServer_Listen(void* ptr, void* address, unsigned short port)
{
	return static_cast<QTcpServer*>(ptr)->listen(*static_cast<QHostAddress*>(address), port);
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

long long QTcpServer_ServerError(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->serverError();
}

unsigned short QTcpServer_ServerPort(void* ptr)
{
	return static_cast<QTcpServer*>(ptr)->serverPort();
}

void QTcpServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QTcpServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QTcpServer_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QTcpServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

char QTcpServer_WaitForNewConnection(void* ptr, int msec, char timedOut)
{
	return static_cast<QTcpServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QTcpServer_DestroyQTcpServer(void* ptr)
{
	static_cast<QTcpServer*>(ptr)->~QTcpServer();
}

void QTcpServer_DestroyQTcpServerDefault(void* ptr)
{

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

char QTcpServer_Event(void* ptr, void* e)
{
	return static_cast<QTcpServer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QTcpServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QTcpServer*>(ptr)->QTcpServer::event(static_cast<QEvent*>(e));
}

char QTcpServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTcpServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QTcpServer_EventFilterDefault(void* ptr, void* watched, void* event)
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
	 ~MyQTcpSocket() { callbackQTcpSocket_DestroyQTcpSocket(this); };
	bool atEnd() const { return callbackQTcpSocket_AtEnd(const_cast<MyQTcpSocket*>(this)) != 0; };
	qint64 bytesAvailable() const { return callbackQTcpSocket_BytesAvailable(const_cast<MyQTcpSocket*>(this)); };
	qint64 bytesToWrite() const { return callbackQTcpSocket_BytesToWrite(const_cast<MyQTcpSocket*>(this)); };
	bool canReadLine() const { return callbackQTcpSocket_CanReadLine(const_cast<MyQTcpSocket*>(this)) != 0; };
	void close() { callbackQTcpSocket_Close(this); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQTcpSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQTcpSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void disconnectFromHost() { callbackQTcpSocket_DisconnectFromHost(this); };
	bool isSequential() const { return callbackQTcpSocket_IsSequential(const_cast<MyQTcpSocket*>(this)) != 0; };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQTcpSocket_ReadLineData(this, dataPacked, maxlen); };
	void resume() { callbackQTcpSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQTcpSocket_SetReadBufferSize(this, size); };
	
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQTcpSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQTcpSocket_SocketOption(this, option)); };
	bool waitForBytesWritten(int msecs) { return callbackQTcpSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQTcpSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQTcpSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQTcpSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 size) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQTcpSocket_WriteData(this, dataPacked, size); };
	bool open(QIODevice::OpenMode mode) { return callbackQTcpSocket_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQTcpSocket_Pos(const_cast<MyQTcpSocket*>(this)); };
	bool reset() { return callbackQTcpSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQTcpSocket_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQTcpSocket_Size(const_cast<MyQTcpSocket*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQTcpSocket_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQTcpSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTcpSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTcpSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQTcpSocket_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTcpSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTcpSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTcpSocket_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTcpSocket_MetaObject(const_cast<MyQTcpSocket*>(this))); };
};

void* QTcpSocket_NewQTcpSocket(void* parent)
{
	return new MyQTcpSocket(static_cast<QObject*>(parent));
}

void QTcpSocket_DestroyQTcpSocket(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->~QTcpSocket();
}

void QTcpSocket_DestroyQTcpSocketDefault(void* ptr)
{

}

char QTcpSocket_AtEnd(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->atEnd();
}

char QTcpSocket_AtEndDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::atEnd();
}

long long QTcpSocket_BytesAvailable(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->bytesAvailable();
}

long long QTcpSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::bytesAvailable();
}

long long QTcpSocket_BytesToWrite(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->bytesToWrite();
}

long long QTcpSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::bytesToWrite();
}

char QTcpSocket_CanReadLine(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->canReadLine();
}

char QTcpSocket_CanReadLineDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::canReadLine();
}

void QTcpSocket_Close(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->close();
}

void QTcpSocket_CloseDefault(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::close();
}

void QTcpSocket_ConnectToHost2(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QTcpSocket*>(ptr)->connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QTcpSocket_ConnectToHost2Default(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QTcpSocket_ConnectToHost(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QTcpSocket*>(ptr)->connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QTcpSocket_ConnectToHostDefault(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QTcpSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->disconnectFromHost();
}

void QTcpSocket_DisconnectFromHostDefault(void* ptr)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::disconnectFromHost();
}

char QTcpSocket_IsSequential(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->isSequential();
}

char QTcpSocket_IsSequentialDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::isSequential();
}

long long QTcpSocket_ReadLineData(void* ptr, char* data, long long maxlen)
{
	return static_cast<QTcpSocket*>(ptr)->readLineData(data, maxlen);
}

long long QTcpSocket_ReadLineDataDefault(void* ptr, char* data, long long maxlen)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::readLineData(data, maxlen);
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
	static_cast<QTcpSocket*>(ptr)->setReadBufferSize(size);
}

void QTcpSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::setReadBufferSize(size);
}





void QTcpSocket_SetSocketOption(void* ptr, long long option, void* value)
{
	static_cast<QTcpSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QTcpSocket_SetSocketOptionDefault(void* ptr, long long option, void* value)
{
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}





void* QTcpSocket_SocketOption(void* ptr, long long option)
{
	return new QVariant(static_cast<QTcpSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QTcpSocket_SocketOptionDefault(void* ptr, long long option)
{
	return new QVariant(static_cast<QTcpSocket*>(ptr)->QTcpSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

char QTcpSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->waitForBytesWritten(msecs);
}

char QTcpSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForBytesWritten(msecs);
}

char QTcpSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->waitForConnected(msecs);
}

char QTcpSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForConnected(msecs);
}

char QTcpSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->waitForDisconnected(msecs);
}

char QTcpSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForDisconnected(msecs);
}

char QTcpSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->waitForReadyRead(msecs);
}

char QTcpSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::waitForReadyRead(msecs);
}

long long QTcpSocket_WriteData(void* ptr, char* data, long long size)
{
	return static_cast<QTcpSocket*>(ptr)->writeData(const_cast<const char*>(data), size);
}

long long QTcpSocket_WriteDataDefault(void* ptr, char* data, long long size)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::writeData(const_cast<const char*>(data), size);
}

char QTcpSocket_Open(void* ptr, long long mode)
{
	return static_cast<QTcpSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QTcpSocket_OpenDefault(void* ptr, long long mode)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QTcpSocket_Pos(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->pos();
}

long long QTcpSocket_PosDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::pos();
}

char QTcpSocket_Reset(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->reset();
}

char QTcpSocket_ResetDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::reset();
}

char QTcpSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QTcpSocket*>(ptr)->seek(pos);
}

char QTcpSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::seek(pos);
}

long long QTcpSocket_Size(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->size();
}

long long QTcpSocket_SizeDefault(void* ptr)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::size();
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

char QTcpSocket_Event(void* ptr, void* e)
{
	return static_cast<QTcpSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

char QTcpSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QTcpSocket*>(ptr)->QTcpSocket::event(static_cast<QEvent*>(e));
}

char QTcpSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTcpSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QTcpSocket_EventFilterDefault(void* ptr, void* watched, void* event)
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
	 ~MyQUdpSocket() { callbackQUdpSocket_DestroyQUdpSocket(this); };
	bool atEnd() const { return callbackQUdpSocket_AtEnd(const_cast<MyQUdpSocket*>(this)) != 0; };
	qint64 bytesAvailable() const { return callbackQUdpSocket_BytesAvailable(const_cast<MyQUdpSocket*>(this)); };
	qint64 bytesToWrite() const { return callbackQUdpSocket_BytesToWrite(const_cast<MyQUdpSocket*>(this)); };
	bool canReadLine() const { return callbackQUdpSocket_CanReadLine(const_cast<MyQUdpSocket*>(this)) != 0; };
	void close() { callbackQUdpSocket_Close(this); };
	void connectToHost(const QHostAddress & address, quint16 port, QIODevice::OpenMode openMode) { callbackQUdpSocket_ConnectToHost2(this, const_cast<QHostAddress*>(&address), port, openMode); };
	void connectToHost(const QString & hostName, quint16 port, QIODevice::OpenMode openMode, QAbstractSocket::NetworkLayerProtocol protocol) { QByteArray tcf2288 = hostName.toUtf8(); QtNetwork_PackedString hostNamePacked = { const_cast<char*>(tcf2288.prepend("WHITESPACE").constData()+10), tcf2288.size()-10 };callbackQUdpSocket_ConnectToHost(this, hostNamePacked, port, openMode, protocol); };
	void disconnectFromHost() { callbackQUdpSocket_DisconnectFromHost(this); };
	bool isSequential() const { return callbackQUdpSocket_IsSequential(const_cast<MyQUdpSocket*>(this)) != 0; };
	qint64 readLineData(char * data, qint64 maxlen) { QtNetwork_PackedString dataPacked = { data, maxlen };return callbackQUdpSocket_ReadLineData(this, dataPacked, maxlen); };
	void resume() { callbackQUdpSocket_Resume(this); };
	void setReadBufferSize(qint64 size) { callbackQUdpSocket_SetReadBufferSize(this, size); };
	
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQUdpSocket_SetSocketOption(this, option, const_cast<QVariant*>(&value)); };
	
	QVariant socketOption(QAbstractSocket::SocketOption option) { return *static_cast<QVariant*>(callbackQUdpSocket_SocketOption(this, option)); };
	bool waitForBytesWritten(int msecs) { return callbackQUdpSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForConnected(int msecs) { return callbackQUdpSocket_WaitForConnected(this, msecs) != 0; };
	bool waitForDisconnected(int msecs) { return callbackQUdpSocket_WaitForDisconnected(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQUdpSocket_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 size) { QtNetwork_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQUdpSocket_WriteData(this, dataPacked, size); };
	bool open(QIODevice::OpenMode mode) { return callbackQUdpSocket_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQUdpSocket_Pos(const_cast<MyQUdpSocket*>(this)); };
	bool reset() { return callbackQUdpSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQUdpSocket_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQUdpSocket_Size(const_cast<MyQUdpSocket*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQUdpSocket_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQUdpSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQUdpSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQUdpSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQUdpSocket_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQUdpSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQUdpSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQUdpSocket_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQUdpSocket_MetaObject(const_cast<MyQUdpSocket*>(this))); };
};

void* QUdpSocket_NewQUdpSocket(void* parent)
{
	return new MyQUdpSocket(static_cast<QObject*>(parent));
}

char QUdpSocket_HasPendingDatagrams(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->hasPendingDatagrams();
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

void* QUdpSocket_MulticastInterface(void* ptr)
{
	return new QNetworkInterface(static_cast<QUdpSocket*>(ptr)->multicastInterface());
}

long long QUdpSocket_PendingDatagramSize(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->pendingDatagramSize();
}

long long QUdpSocket_ReadDatagram(void* ptr, char* data, long long maxSize, void* address, unsigned short port)
{
	return static_cast<QUdpSocket*>(ptr)->readDatagram(data, maxSize, static_cast<QHostAddress*>(address), &port);
}

void QUdpSocket_SetMulticastInterface(void* ptr, void* iface)
{
	static_cast<QUdpSocket*>(ptr)->setMulticastInterface(*static_cast<QNetworkInterface*>(iface));
}

long long QUdpSocket_WriteDatagram2(void* ptr, void* datagram, void* host, unsigned short port)
{
	return static_cast<QUdpSocket*>(ptr)->writeDatagram(*static_cast<QByteArray*>(datagram), *static_cast<QHostAddress*>(host), port);
}

long long QUdpSocket_WriteDatagram(void* ptr, char* data, long long size, void* address, unsigned short port)
{
	return static_cast<QUdpSocket*>(ptr)->writeDatagram(const_cast<const char*>(data), size, *static_cast<QHostAddress*>(address), port);
}

void QUdpSocket_DestroyQUdpSocket(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->~QUdpSocket();
}

void QUdpSocket_DestroyQUdpSocketDefault(void* ptr)
{

}

char QUdpSocket_AtEnd(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->atEnd();
}

char QUdpSocket_AtEndDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::atEnd();
}

long long QUdpSocket_BytesAvailable(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->bytesAvailable();
}

long long QUdpSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::bytesAvailable();
}

long long QUdpSocket_BytesToWrite(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->bytesToWrite();
}

long long QUdpSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::bytesToWrite();
}

char QUdpSocket_CanReadLine(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->canReadLine();
}

char QUdpSocket_CanReadLineDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::canReadLine();
}

void QUdpSocket_Close(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->close();
}

void QUdpSocket_CloseDefault(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::close();
}

void QUdpSocket_ConnectToHost2(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QUdpSocket*>(ptr)->connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QUdpSocket_ConnectToHost2Default(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::connectToHost(*static_cast<QHostAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QUdpSocket_ConnectToHost(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QUdpSocket*>(ptr)->connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QUdpSocket_ConnectToHostDefault(void* ptr, char* hostName, unsigned short port, long long openMode, long long protocol)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::connectToHost(QString(hostName), port, static_cast<QIODevice::OpenModeFlag>(openMode), static_cast<QAbstractSocket::NetworkLayerProtocol>(protocol));
}

void QUdpSocket_DisconnectFromHost(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->disconnectFromHost();
}

void QUdpSocket_DisconnectFromHostDefault(void* ptr)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::disconnectFromHost();
}

char QUdpSocket_IsSequential(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->isSequential();
}

char QUdpSocket_IsSequentialDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::isSequential();
}

long long QUdpSocket_ReadLineData(void* ptr, char* data, long long maxlen)
{
	return static_cast<QUdpSocket*>(ptr)->readLineData(data, maxlen);
}

long long QUdpSocket_ReadLineDataDefault(void* ptr, char* data, long long maxlen)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::readLineData(data, maxlen);
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
	static_cast<QUdpSocket*>(ptr)->setReadBufferSize(size);
}

void QUdpSocket_SetReadBufferSizeDefault(void* ptr, long long size)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::setReadBufferSize(size);
}





void QUdpSocket_SetSocketOption(void* ptr, long long option, void* value)
{
	static_cast<QUdpSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QUdpSocket_SetSocketOptionDefault(void* ptr, long long option, void* value)
{
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}





void* QUdpSocket_SocketOption(void* ptr, long long option)
{
	return new QVariant(static_cast<QUdpSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

void* QUdpSocket_SocketOptionDefault(void* ptr, long long option)
{
	return new QVariant(static_cast<QUdpSocket*>(ptr)->QUdpSocket::socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

char QUdpSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->waitForBytesWritten(msecs);
}

char QUdpSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForBytesWritten(msecs);
}

char QUdpSocket_WaitForConnected(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->waitForConnected(msecs);
}

char QUdpSocket_WaitForConnectedDefault(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForConnected(msecs);
}

char QUdpSocket_WaitForDisconnected(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->waitForDisconnected(msecs);
}

char QUdpSocket_WaitForDisconnectedDefault(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForDisconnected(msecs);
}

char QUdpSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->waitForReadyRead(msecs);
}

char QUdpSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::waitForReadyRead(msecs);
}

long long QUdpSocket_WriteData(void* ptr, char* data, long long size)
{
	return static_cast<QUdpSocket*>(ptr)->writeData(const_cast<const char*>(data), size);
}

long long QUdpSocket_WriteDataDefault(void* ptr, char* data, long long size)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::writeData(const_cast<const char*>(data), size);
}

char QUdpSocket_Open(void* ptr, long long mode)
{
	return static_cast<QUdpSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QUdpSocket_OpenDefault(void* ptr, long long mode)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QUdpSocket_Pos(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->pos();
}

long long QUdpSocket_PosDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::pos();
}

char QUdpSocket_Reset(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->reset();
}

char QUdpSocket_ResetDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::reset();
}

char QUdpSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QUdpSocket*>(ptr)->seek(pos);
}

char QUdpSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::seek(pos);
}

long long QUdpSocket_Size(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->size();
}

long long QUdpSocket_SizeDefault(void* ptr)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::size();
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

char QUdpSocket_Event(void* ptr, void* e)
{
	return static_cast<QUdpSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

char QUdpSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QUdpSocket*>(ptr)->QUdpSocket::event(static_cast<QEvent*>(e));
}

char QUdpSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QUdpSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QUdpSocket_EventFilterDefault(void* ptr, void* watched, void* event)
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

