#define protected public

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
#include <QTcpServer>
#include <QTcpSocket>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUdpSocket>
#include <QUrl>
#include <QVariant>

class MyQAbstractNetworkCache: public QAbstractNetworkCache {
public:
	void timerEvent(QTimerEvent * event) { callbackQAbstractNetworkCacheTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractNetworkCacheChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractNetworkCacheCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

long long QAbstractNetworkCache_CacheSize(void* ptr){
	return static_cast<long long>(static_cast<QAbstractNetworkCache*>(ptr)->cacheSize());
}

void QAbstractNetworkCache_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractNetworkCache*>(ptr), "clear");
}

void* QAbstractNetworkCache_Data(void* ptr, void* url){
	return static_cast<QAbstractNetworkCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void* QAbstractNetworkCache_Prepare(void* ptr, void* metaData){
	return static_cast<QAbstractNetworkCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_UpdateMetaData(void* ptr, void* metaData){
	static_cast<QAbstractNetworkCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_DestroyQAbstractNetworkCache(void* ptr){
	static_cast<QAbstractNetworkCache*>(ptr)->~QAbstractNetworkCache();
}

void QAbstractNetworkCache_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractNetworkCache*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractNetworkCache_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractNetworkCache_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractNetworkCache*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractNetworkCache_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractNetworkCache_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractNetworkCache*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractNetworkCache_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractNetworkCache*>(ptr)->QAbstractNetworkCache::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractSocket: public QAbstractSocket {
public:
	MyQAbstractSocket(SocketType socketType, QObject *parent) : QAbstractSocket(socketType, parent) {};
	void close() { callbackQAbstractSocketClose(this, this->objectName().toUtf8().data()); };
	void Signal_Connected() { callbackQAbstractSocketConnected(this, this->objectName().toUtf8().data()); };
	void disconnectFromHost() { callbackQAbstractSocketDisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void Signal_Disconnected() { callbackQAbstractSocketDisconnected(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QAbstractSocket::SocketError socketError) { callbackQAbstractSocketError2(this, this->objectName().toUtf8().data(), socketError); };
	void Signal_HostFound() { callbackQAbstractSocketHostFound(this, this->objectName().toUtf8().data()); };
	void resume() { callbackQAbstractSocketResume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQAbstractSocketSetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQAbstractSocketSetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	void Signal_StateChanged(QAbstractSocket::SocketState socketState) { callbackQAbstractSocketStateChanged(this, this->objectName().toUtf8().data(), socketState); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSocketTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractSocketChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractSocketCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAbstractSocket_NewQAbstractSocket(int socketType, void* parent){
	return new MyQAbstractSocket(static_cast<QAbstractSocket::SocketType>(socketType), static_cast<QObject*>(parent));
}

void QAbstractSocket_Abort(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->abort();
}

int QAbstractSocket_AtEnd(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->atEnd();
}

long long QAbstractSocket_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->bytesAvailable());
}

long long QAbstractSocket_BytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->bytesToWrite());
}

int QAbstractSocket_CanReadLine(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->canReadLine();
}

void QAbstractSocket_Close(void* ptr){
	static_cast<MyQAbstractSocket*>(ptr)->close();
}

void QAbstractSocket_CloseDefault(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::close();
}

void QAbstractSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));;
}

void QAbstractSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Connected));;
}

void QAbstractSocket_Connected(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->connected();
}

void QAbstractSocket_DisconnectFromHost(void* ptr){
	static_cast<MyQAbstractSocket*>(ptr)->disconnectFromHost();
}

void QAbstractSocket_DisconnectFromHostDefault(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::disconnectFromHost();
}

void QAbstractSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));;
}

void QAbstractSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_Disconnected));;
}

void QAbstractSocket_Disconnected(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->disconnected();
}

void QAbstractSocket_ConnectError2(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::error), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketError)>(&MyQAbstractSocket::Signal_Error2));;
}

void QAbstractSocket_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::error), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketError)>(&MyQAbstractSocket::Signal_Error2));;
}

void QAbstractSocket_Error2(void* ptr, int socketError){
	static_cast<QAbstractSocket*>(ptr)->error(static_cast<QAbstractSocket::SocketError>(socketError));
}

int QAbstractSocket_Error(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->error();
}

int QAbstractSocket_Flush(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->flush();
}

void QAbstractSocket_ConnectHostFound(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));;
}

void QAbstractSocket_DisconnectHostFound(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)()>(&MyQAbstractSocket::Signal_HostFound));;
}

void QAbstractSocket_HostFound(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->hostFound();
}

int QAbstractSocket_IsSequential(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->isSequential();
}

int QAbstractSocket_IsValid(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->isValid();
}

int QAbstractSocket_PauseMode(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->pauseMode();
}

char* QAbstractSocket_PeerName(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->peerName().toUtf8().data();
}

long long QAbstractSocket_ReadBufferSize(void* ptr){
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->readBufferSize());
}

long long QAbstractSocket_ReadData(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->readData(data, static_cast<long long>(maxSize)));
}

long long QAbstractSocket_ReadLineData(void* ptr, char* data, long long maxlen){
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->readLineData(data, static_cast<long long>(maxlen)));
}

void QAbstractSocket_Resume(void* ptr){
	static_cast<MyQAbstractSocket*>(ptr)->resume();
}

void QAbstractSocket_ResumeDefault(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::resume();
}

void QAbstractSocket_SetPauseMode(void* ptr, int pauseMode){
	static_cast<QAbstractSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QAbstractSocket_SetProxy(void* ptr, void* networkProxy){
	static_cast<QAbstractSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QAbstractSocket_SetReadBufferSize(void* ptr, long long size){
	static_cast<MyQAbstractSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QAbstractSocket_SetReadBufferSizeDefault(void* ptr, long long size){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setReadBufferSize(static_cast<long long>(size));
}

void QAbstractSocket_SetSocketOption(void* ptr, int option, void* value){
	static_cast<MyQAbstractSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QAbstractSocket_SetSocketOptionDefault(void* ptr, int option, void* value){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void* QAbstractSocket_SocketOption(void* ptr, int option){
	return new QVariant(static_cast<QAbstractSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

int QAbstractSocket_SocketType(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->socketType();
}

int QAbstractSocket_State(void* ptr){
	return static_cast<QAbstractSocket*>(ptr)->state();
}

void QAbstractSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));;
}

void QAbstractSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractSocket*>(ptr), static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), static_cast<MyQAbstractSocket*>(ptr), static_cast<void (MyQAbstractSocket::*)(QAbstractSocket::SocketState)>(&MyQAbstractSocket::Signal_StateChanged));;
}

void QAbstractSocket_StateChanged(void* ptr, int socketState){
	static_cast<QAbstractSocket*>(ptr)->stateChanged(static_cast<QAbstractSocket::SocketState>(socketState));
}

int QAbstractSocket_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QAbstractSocket_WaitForConnected(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForConnected(msecs);
}

int QAbstractSocket_WaitForDisconnected(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForDisconnected(msecs);
}

int QAbstractSocket_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QAbstractSocket*>(ptr)->waitForReadyRead(msecs);
}

long long QAbstractSocket_WriteData(void* ptr, char* data, long long size){
	return static_cast<long long>(static_cast<QAbstractSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(size)));
}

void QAbstractSocket_DestroyQAbstractSocket(void* ptr){
	static_cast<QAbstractSocket*>(ptr)->~QAbstractSocket();
}

void QAbstractSocket_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractSocket_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractSocket_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractSocket_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractSocket_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractSocket_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractSocket*>(ptr)->QAbstractSocket::customEvent(static_cast<QEvent*>(event));
}

void* QAuthenticator_NewQAuthenticator(){
	return new QAuthenticator();
}

void* QAuthenticator_NewQAuthenticator2(void* other){
	return new QAuthenticator(*static_cast<QAuthenticator*>(other));
}

int QAuthenticator_IsNull(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->isNull();
}

void* QAuthenticator_Option(void* ptr, char* opt){
	return new QVariant(static_cast<QAuthenticator*>(ptr)->option(QString(opt)));
}

char* QAuthenticator_Password(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->password().toUtf8().data();
}

char* QAuthenticator_Realm(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->realm().toUtf8().data();
}

void QAuthenticator_SetOption(void* ptr, char* opt, void* value){
	static_cast<QAuthenticator*>(ptr)->setOption(QString(opt), *static_cast<QVariant*>(value));
}

void QAuthenticator_SetPassword(void* ptr, char* password){
	static_cast<QAuthenticator*>(ptr)->setPassword(QString(password));
}

void QAuthenticator_SetUser(void* ptr, char* user){
	static_cast<QAuthenticator*>(ptr)->setUser(QString(user));
}

char* QAuthenticator_User(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->user().toUtf8().data();
}

void QAuthenticator_DestroyQAuthenticator(void* ptr){
	static_cast<QAuthenticator*>(ptr)->~QAuthenticator();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord(){
	return new QDnsDomainNameRecord();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord2(void* other){
	return new QDnsDomainNameRecord(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Name(void* ptr){
	return static_cast<QDnsDomainNameRecord*>(ptr)->name().toUtf8().data();
}

void QDnsDomainNameRecord_Swap(void* ptr, void* other){
	static_cast<QDnsDomainNameRecord*>(ptr)->swap(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Value(void* ptr){
	return static_cast<QDnsDomainNameRecord*>(ptr)->value().toUtf8().data();
}

void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(void* ptr){
	static_cast<QDnsDomainNameRecord*>(ptr)->~QDnsDomainNameRecord();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord(){
	return new QDnsHostAddressRecord();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord2(void* other){
	return new QDnsHostAddressRecord(*static_cast<QDnsHostAddressRecord*>(other));
}

char* QDnsHostAddressRecord_Name(void* ptr){
	return static_cast<QDnsHostAddressRecord*>(ptr)->name().toUtf8().data();
}

void QDnsHostAddressRecord_Swap(void* ptr, void* other){
	static_cast<QDnsHostAddressRecord*>(ptr)->swap(*static_cast<QDnsHostAddressRecord*>(other));
}

void QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(void* ptr){
	static_cast<QDnsHostAddressRecord*>(ptr)->~QDnsHostAddressRecord();
}

class MyQDnsLookup: public QDnsLookup {
public:
	void Signal_Finished() { callbackQDnsLookupFinished(this, this->objectName().toUtf8().data()); };
	void Signal_NameChanged(const QString & name) { callbackQDnsLookupNameChanged(this, this->objectName().toUtf8().data(), name.toUtf8().data()); };
	void Signal_TypeChanged(QDnsLookup::Type ty) { callbackQDnsLookupTypeChanged(this, this->objectName().toUtf8().data(), ty); };
	void timerEvent(QTimerEvent * event) { callbackQDnsLookupTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDnsLookupChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDnsLookupCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QDnsLookup_NewQDnsLookup3(int ty, char* name, void* nameserver, void* parent){
	return new QDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), *static_cast<QHostAddress*>(nameserver), static_cast<QObject*>(parent));
}

int QDnsLookup_Error(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->error();
}

char* QDnsLookup_ErrorString(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->errorString().toUtf8().data();
}

char* QDnsLookup_Name(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->name().toUtf8().data();
}

void QDnsLookup_SetName(void* ptr, char* name){
	static_cast<QDnsLookup*>(ptr)->setName(QString(name));
}

void QDnsLookup_SetNameserver(void* ptr, void* nameserver){
	static_cast<QDnsLookup*>(ptr)->setNameserver(*static_cast<QHostAddress*>(nameserver));
}

void QDnsLookup_SetType(void* ptr, int v){
	static_cast<QDnsLookup*>(ptr)->setType(static_cast<QDnsLookup::Type>(v));
}

int QDnsLookup_Type(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->type();
}

void* QDnsLookup_NewQDnsLookup(void* parent){
	return new QDnsLookup(static_cast<QObject*>(parent));
}

void* QDnsLookup_NewQDnsLookup2(int ty, char* name, void* parent){
	return new QDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), static_cast<QObject*>(parent));
}

void QDnsLookup_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "abort");
}

void QDnsLookup_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));;
}

void QDnsLookup_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));;
}

void QDnsLookup_Finished(void* ptr){
	static_cast<QDnsLookup*>(ptr)->finished();
}

int QDnsLookup_IsFinished(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->isFinished();
}

void QDnsLookup_Lookup(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "lookup");
}

void QDnsLookup_ConnectNameChanged(void* ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));;
}

void QDnsLookup_DisconnectNameChanged(void* ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));;
}

void QDnsLookup_NameChanged(void* ptr, char* name){
	static_cast<QDnsLookup*>(ptr)->nameChanged(QString(name));
}

void QDnsLookup_ConnectTypeChanged(void* ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));;
}

void QDnsLookup_DisconnectTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));;
}

void QDnsLookup_TypeChanged(void* ptr, int ty){
	static_cast<QDnsLookup*>(ptr)->typeChanged(static_cast<QDnsLookup::Type>(ty));
}

void QDnsLookup_DestroyQDnsLookup(void* ptr){
	static_cast<QDnsLookup*>(ptr)->~QDnsLookup();
}

void QDnsLookup_TimerEvent(void* ptr, void* event){
	static_cast<MyQDnsLookup*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDnsLookup_TimerEventDefault(void* ptr, void* event){
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDnsLookup_ChildEvent(void* ptr, void* event){
	static_cast<MyQDnsLookup*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDnsLookup_ChildEventDefault(void* ptr, void* event){
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::childEvent(static_cast<QChildEvent*>(event));
}

void QDnsLookup_CustomEvent(void* ptr, void* event){
	static_cast<MyQDnsLookup*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDnsLookup_CustomEventDefault(void* ptr, void* event){
	static_cast<QDnsLookup*>(ptr)->QDnsLookup::customEvent(static_cast<QEvent*>(event));
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord(){
	return new QDnsMailExchangeRecord();
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(void* other){
	return new QDnsMailExchangeRecord(*static_cast<QDnsMailExchangeRecord*>(other));
}

char* QDnsMailExchangeRecord_Exchange(void* ptr){
	return static_cast<QDnsMailExchangeRecord*>(ptr)->exchange().toUtf8().data();
}

char* QDnsMailExchangeRecord_Name(void* ptr){
	return static_cast<QDnsMailExchangeRecord*>(ptr)->name().toUtf8().data();
}

void QDnsMailExchangeRecord_Swap(void* ptr, void* other){
	static_cast<QDnsMailExchangeRecord*>(ptr)->swap(*static_cast<QDnsMailExchangeRecord*>(other));
}

void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(void* ptr){
	static_cast<QDnsMailExchangeRecord*>(ptr)->~QDnsMailExchangeRecord();
}

void* QDnsServiceRecord_NewQDnsServiceRecord(){
	return new QDnsServiceRecord();
}

void* QDnsServiceRecord_NewQDnsServiceRecord2(void* other){
	return new QDnsServiceRecord(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Name(void* ptr){
	return static_cast<QDnsServiceRecord*>(ptr)->name().toUtf8().data();
}

void QDnsServiceRecord_Swap(void* ptr, void* other){
	static_cast<QDnsServiceRecord*>(ptr)->swap(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Target(void* ptr){
	return static_cast<QDnsServiceRecord*>(ptr)->target().toUtf8().data();
}

void QDnsServiceRecord_DestroyQDnsServiceRecord(void* ptr){
	static_cast<QDnsServiceRecord*>(ptr)->~QDnsServiceRecord();
}

void* QDnsTextRecord_NewQDnsTextRecord(){
	return new QDnsTextRecord();
}

void* QDnsTextRecord_NewQDnsTextRecord2(void* other){
	return new QDnsTextRecord(*static_cast<QDnsTextRecord*>(other));
}

char* QDnsTextRecord_Name(void* ptr){
	return static_cast<QDnsTextRecord*>(ptr)->name().toUtf8().data();
}

void QDnsTextRecord_Swap(void* ptr, void* other){
	static_cast<QDnsTextRecord*>(ptr)->swap(*static_cast<QDnsTextRecord*>(other));
}

void QDnsTextRecord_DestroyQDnsTextRecord(void* ptr){
	static_cast<QDnsTextRecord*>(ptr)->~QDnsTextRecord();
}

void* QHostAddress_NewQHostAddress(){
	return new QHostAddress();
}

void* QHostAddress_NewQHostAddress9(int address){
	return new QHostAddress(static_cast<QHostAddress::SpecialAddress>(address));
}

void* QHostAddress_NewQHostAddress8(void* address){
	return new QHostAddress(*static_cast<QHostAddress*>(address));
}

void* QHostAddress_NewQHostAddress7(char* address){
	return new QHostAddress(QString(address));
}

void QHostAddress_Clear(void* ptr){
	static_cast<QHostAddress*>(ptr)->clear();
}

int QHostAddress_IsInSubnet(void* ptr, void* subnet, int netmask){
	return static_cast<QHostAddress*>(ptr)->isInSubnet(*static_cast<QHostAddress*>(subnet), netmask);
}

int QHostAddress_IsLoopback(void* ptr){
	return static_cast<QHostAddress*>(ptr)->isLoopback();
}

int QHostAddress_IsNull(void* ptr){
	return static_cast<QHostAddress*>(ptr)->isNull();
}

int QHostAddress_Protocol(void* ptr){
	return static_cast<QHostAddress*>(ptr)->protocol();
}

char* QHostAddress_ScopeId(void* ptr){
	return static_cast<QHostAddress*>(ptr)->scopeId().toUtf8().data();
}

int QHostAddress_SetAddress5(void* ptr, char* address){
	return static_cast<QHostAddress*>(ptr)->setAddress(QString(address));
}

void QHostAddress_SetScopeId(void* ptr, char* id){
	static_cast<QHostAddress*>(ptr)->setScopeId(QString(id));
}

char* QHostAddress_ToString(void* ptr){
	return static_cast<QHostAddress*>(ptr)->toString().toUtf8().data();
}

void QHostAddress_DestroyQHostAddress(void* ptr){
	static_cast<QHostAddress*>(ptr)->~QHostAddress();
}

void* QHostInfo_NewQHostInfo2(void* other){
	return new QHostInfo(*static_cast<QHostInfo*>(other));
}

void* QHostInfo_NewQHostInfo(int id){
	return new QHostInfo(id);
}

void QHostInfo_QHostInfo_AbortHostLookup(int id){
	QHostInfo::abortHostLookup(id);
}

int QHostInfo_Error(void* ptr){
	return static_cast<QHostInfo*>(ptr)->error();
}

char* QHostInfo_ErrorString(void* ptr){
	return static_cast<QHostInfo*>(ptr)->errorString().toUtf8().data();
}

char* QHostInfo_HostName(void* ptr){
	return static_cast<QHostInfo*>(ptr)->hostName().toUtf8().data();
}

int QHostInfo_QHostInfo_LookupHost(char* name, void* receiver, char* member){
	return QHostInfo::lookupHost(QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QHostInfo_LookupId(void* ptr){
	return static_cast<QHostInfo*>(ptr)->lookupId();
}

void QHostInfo_SetError(void* ptr, int error){
	static_cast<QHostInfo*>(ptr)->setError(static_cast<QHostInfo::HostInfoError>(error));
}

void QHostInfo_SetErrorString(void* ptr, char* str){
	static_cast<QHostInfo*>(ptr)->setErrorString(QString(str));
}

void QHostInfo_SetHostName(void* ptr, char* hostName){
	static_cast<QHostInfo*>(ptr)->setHostName(QString(hostName));
}

void QHostInfo_SetLookupId(void* ptr, int id){
	static_cast<QHostInfo*>(ptr)->setLookupId(id);
}

void QHostInfo_DestroyQHostInfo(void* ptr){
	static_cast<QHostInfo*>(ptr)->~QHostInfo();
}

char* QHostInfo_QHostInfo_LocalHostName(){
	return QHostInfo::localHostName().toUtf8().data();
}

char* QHostInfo_QHostInfo_LocalDomainName(){
	return QHostInfo::localDomainName().toUtf8().data();
}

void* QHttpMultiPart_NewQHttpMultiPart2(int contentType, void* parent){
	return new QHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QObject*>(parent));
}

void* QHttpMultiPart_NewQHttpMultiPart(void* parent){
	return new QHttpMultiPart(static_cast<QObject*>(parent));
}

void QHttpMultiPart_Append(void* ptr, void* httpPart){
	static_cast<QHttpMultiPart*>(ptr)->append(*static_cast<QHttpPart*>(httpPart));
}

void* QHttpMultiPart_Boundary(void* ptr){
	return new QByteArray(static_cast<QHttpMultiPart*>(ptr)->boundary());
}

void QHttpMultiPart_SetBoundary(void* ptr, void* boundary){
	static_cast<QHttpMultiPart*>(ptr)->setBoundary(*static_cast<QByteArray*>(boundary));
}

void QHttpMultiPart_SetContentType(void* ptr, int contentType){
	static_cast<QHttpMultiPart*>(ptr)->setContentType(static_cast<QHttpMultiPart::ContentType>(contentType));
}

void QHttpMultiPart_DestroyQHttpMultiPart(void* ptr){
	static_cast<QHttpMultiPart*>(ptr)->~QHttpMultiPart();
}

void QHttpMultiPart_TimerEvent(void* ptr, void* event){
	static_cast<QHttpMultiPart*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHttpMultiPart_TimerEventDefault(void* ptr, void* event){
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHttpMultiPart_ChildEvent(void* ptr, void* event){
	static_cast<QHttpMultiPart*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHttpMultiPart_ChildEventDefault(void* ptr, void* event){
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::childEvent(static_cast<QChildEvent*>(event));
}

void QHttpMultiPart_CustomEvent(void* ptr, void* event){
	static_cast<QHttpMultiPart*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHttpMultiPart_CustomEventDefault(void* ptr, void* event){
	static_cast<QHttpMultiPart*>(ptr)->QHttpMultiPart::customEvent(static_cast<QEvent*>(event));
}

void* QHttpPart_NewQHttpPart(){
	return new QHttpPart();
}

void* QHttpPart_NewQHttpPart2(void* other){
	return new QHttpPart(*static_cast<QHttpPart*>(other));
}

void QHttpPart_SetBody(void* ptr, void* body){
	static_cast<QHttpPart*>(ptr)->setBody(*static_cast<QByteArray*>(body));
}

void QHttpPart_SetBodyDevice(void* ptr, void* device){
	static_cast<QHttpPart*>(ptr)->setBodyDevice(static_cast<QIODevice*>(device));
}

void QHttpPart_SetHeader(void* ptr, int header, void* value){
	static_cast<QHttpPart*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QHttpPart_SetRawHeader(void* ptr, void* headerName, void* headerValue){
	static_cast<QHttpPart*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QHttpPart_Swap(void* ptr, void* other){
	static_cast<QHttpPart*>(ptr)->swap(*static_cast<QHttpPart*>(other));
}

void QHttpPart_DestroyQHttpPart(void* ptr){
	static_cast<QHttpPart*>(ptr)->~QHttpPart();
}

class MyQLocalServer: public QLocalServer {
public:
	MyQLocalServer(QObject *parent) : QLocalServer(parent) {};
	void Signal_NewConnection() { callbackQLocalServerNewConnection(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQLocalServerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLocalServerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQLocalServerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QLocalServer_SetSocketOptions(void* ptr, int options){
	static_cast<QLocalServer*>(ptr)->setSocketOptions(static_cast<QLocalServer::SocketOption>(options));
}

void* QLocalServer_NewQLocalServer(void* parent){
	return new MyQLocalServer(static_cast<QObject*>(parent));
}

void QLocalServer_Close(void* ptr){
	static_cast<QLocalServer*>(ptr)->close();
}

char* QLocalServer_ErrorString(void* ptr){
	return static_cast<QLocalServer*>(ptr)->errorString().toUtf8().data();
}

char* QLocalServer_FullServerName(void* ptr){
	return static_cast<QLocalServer*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalServer_HasPendingConnections(void* ptr){
	return static_cast<QLocalServer*>(ptr)->hasPendingConnections();
}

int QLocalServer_IsListening(void* ptr){
	return static_cast<QLocalServer*>(ptr)->isListening();
}

int QLocalServer_Listen(void* ptr, char* name){
	return static_cast<QLocalServer*>(ptr)->listen(QString(name));
}

int QLocalServer_MaxPendingConnections(void* ptr){
	return static_cast<QLocalServer*>(ptr)->maxPendingConnections();
}

void QLocalServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));;
}

void QLocalServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QLocalServer*>(ptr), static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), static_cast<MyQLocalServer*>(ptr), static_cast<void (MyQLocalServer::*)()>(&MyQLocalServer::Signal_NewConnection));;
}

void QLocalServer_NewConnection(void* ptr){
	static_cast<QLocalServer*>(ptr)->newConnection();
}

void* QLocalServer_NextPendingConnection(void* ptr){
	return static_cast<QLocalServer*>(ptr)->nextPendingConnection();
}

int QLocalServer_QLocalServer_RemoveServer(char* name){
	return QLocalServer::removeServer(QString(name));
}

int QLocalServer_ServerError(void* ptr){
	return static_cast<QLocalServer*>(ptr)->serverError();
}

char* QLocalServer_ServerName(void* ptr){
	return static_cast<QLocalServer*>(ptr)->serverName().toUtf8().data();
}

void QLocalServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QLocalServer*>(ptr)->setMaxPendingConnections(numConnections);
}

int QLocalServer_SocketOptions(void* ptr){
	return static_cast<QLocalServer*>(ptr)->socketOptions();
}

int QLocalServer_WaitForNewConnection(void* ptr, int msec, int timedOut){
	return static_cast<QLocalServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QLocalServer_DestroyQLocalServer(void* ptr){
	static_cast<QLocalServer*>(ptr)->~QLocalServer();
}

void QLocalServer_TimerEvent(void* ptr, void* event){
	static_cast<MyQLocalServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalServer_TimerEventDefault(void* ptr, void* event){
	static_cast<QLocalServer*>(ptr)->QLocalServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalServer_ChildEvent(void* ptr, void* event){
	static_cast<MyQLocalServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLocalServer_ChildEventDefault(void* ptr, void* event){
	static_cast<QLocalServer*>(ptr)->QLocalServer::childEvent(static_cast<QChildEvent*>(event));
}

void QLocalServer_CustomEvent(void* ptr, void* event){
	static_cast<MyQLocalServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLocalServer_CustomEventDefault(void* ptr, void* event){
	static_cast<QLocalServer*>(ptr)->QLocalServer::customEvent(static_cast<QEvent*>(event));
}

class MyQLocalSocket: public QLocalSocket {
public:
	MyQLocalSocket(QObject *parent) : QLocalSocket(parent) {};
	void Signal_Connected() { callbackQLocalSocketConnected(this, this->objectName().toUtf8().data()); };
	void Signal_Disconnected() { callbackQLocalSocketDisconnected(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QLocalSocket::LocalSocketError socketError) { callbackQLocalSocketError2(this, this->objectName().toUtf8().data(), socketError); };
	void Signal_StateChanged(QLocalSocket::LocalSocketState socketState) { callbackQLocalSocketStateChanged(this, this->objectName().toUtf8().data(), socketState); };
	void close() { callbackQLocalSocketClose(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQLocalSocketTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLocalSocketChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQLocalSocketCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QLocalSocket_Open(void* ptr, int openMode){
	return static_cast<QLocalSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QLocalSocket_NewQLocalSocket(void* parent){
	return new MyQLocalSocket(static_cast<QObject*>(parent));
}

void QLocalSocket_ConnectToServer2(void* ptr, char* name, int openMode){
	static_cast<QLocalSocket*>(ptr)->connectToServer(QString(name), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Connected));;
}

void QLocalSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Connected));;
}

void QLocalSocket_Connected(void* ptr){
	static_cast<QLocalSocket*>(ptr)->connected();
}

void QLocalSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Disconnected));;
}

void QLocalSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)()>(&MyQLocalSocket::Signal_Disconnected));;
}

void QLocalSocket_Disconnected(void* ptr){
	static_cast<QLocalSocket*>(ptr)->disconnected();
}

void QLocalSocket_ConnectError2(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketError)>(&QLocalSocket::error), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketError)>(&MyQLocalSocket::Signal_Error2));;
}

void QLocalSocket_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketError)>(&QLocalSocket::error), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketError)>(&MyQLocalSocket::Signal_Error2));;
}

void QLocalSocket_Error2(void* ptr, int socketError){
	static_cast<QLocalSocket*>(ptr)->error(static_cast<QLocalSocket::LocalSocketError>(socketError));
}

char* QLocalSocket_FullServerName(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->fullServerName().toUtf8().data();
}

int QLocalSocket_IsSequential(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->isSequential();
}

char* QLocalSocket_ServerName(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->serverName().toUtf8().data();
}

void QLocalSocket_SetServerName(void* ptr, char* name){
	static_cast<QLocalSocket*>(ptr)->setServerName(QString(name));
}

int QLocalSocket_State(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->state();
}

void QLocalSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));;
}

void QLocalSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QLocalSocket*>(ptr), static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), static_cast<MyQLocalSocket*>(ptr), static_cast<void (MyQLocalSocket::*)(QLocalSocket::LocalSocketState)>(&MyQLocalSocket::Signal_StateChanged));;
}

void QLocalSocket_StateChanged(void* ptr, int socketState){
	static_cast<QLocalSocket*>(ptr)->stateChanged(static_cast<QLocalSocket::LocalSocketState>(socketState));
}

void QLocalSocket_DestroyQLocalSocket(void* ptr){
	static_cast<QLocalSocket*>(ptr)->~QLocalSocket();
}

void QLocalSocket_Abort(void* ptr){
	static_cast<QLocalSocket*>(ptr)->abort();
}

long long QLocalSocket_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->bytesAvailable());
}

long long QLocalSocket_BytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->bytesToWrite());
}

int QLocalSocket_CanReadLine(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->canReadLine();
}

void QLocalSocket_Close(void* ptr){
	static_cast<MyQLocalSocket*>(ptr)->close();
}

void QLocalSocket_CloseDefault(void* ptr){
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::close();
}

void QLocalSocket_ConnectToServer(void* ptr, int openMode){
	static_cast<QLocalSocket*>(ptr)->connectToServer(static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QLocalSocket_DisconnectFromServer(void* ptr){
	static_cast<QLocalSocket*>(ptr)->disconnectFromServer();
}

int QLocalSocket_Error(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->error();
}

int QLocalSocket_Flush(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->flush();
}

int QLocalSocket_IsValid(void* ptr){
	return static_cast<QLocalSocket*>(ptr)->isValid();
}

long long QLocalSocket_ReadBufferSize(void* ptr){
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->readBufferSize());
}

long long QLocalSocket_ReadData(void* ptr, char* data, long long c){
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->readData(data, static_cast<long long>(c)));
}

void QLocalSocket_SetReadBufferSize(void* ptr, long long size){
	static_cast<QLocalSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

int QLocalSocket_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QLocalSocket_WaitForConnected(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForConnected(msecs);
}

int QLocalSocket_WaitForDisconnected(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForDisconnected(msecs);
}

int QLocalSocket_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QLocalSocket*>(ptr)->waitForReadyRead(msecs);
}

long long QLocalSocket_WriteData(void* ptr, char* data, long long c){
	return static_cast<long long>(static_cast<QLocalSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(c)));
}

void QLocalSocket_TimerEvent(void* ptr, void* event){
	static_cast<MyQLocalSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalSocket_TimerEventDefault(void* ptr, void* event){
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLocalSocket_ChildEvent(void* ptr, void* event){
	static_cast<MyQLocalSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLocalSocket_ChildEventDefault(void* ptr, void* event){
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QLocalSocket_CustomEvent(void* ptr, void* event){
	static_cast<MyQLocalSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLocalSocket_CustomEventDefault(void* ptr, void* event){
	static_cast<QLocalSocket*>(ptr)->QLocalSocket::customEvent(static_cast<QEvent*>(event));
}

class MyQNetworkAccessManager: public QNetworkAccessManager {
public:
	MyQNetworkAccessManager(QObject *parent) : QNetworkAccessManager(parent) {};
	void Signal_AuthenticationRequired(QNetworkReply * reply, QAuthenticator * authenticator) { callbackQNetworkAccessManagerAuthenticationRequired(this, this->objectName().toUtf8().data(), reply, authenticator); };
	void Signal_Encrypted(QNetworkReply * reply) { callbackQNetworkAccessManagerEncrypted(this, this->objectName().toUtf8().data(), reply); };
	void Signal_Finished(QNetworkReply * reply) { callbackQNetworkAccessManagerFinished(this, this->objectName().toUtf8().data(), reply); };
	void Signal_NetworkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility accessible) { callbackQNetworkAccessManagerNetworkAccessibleChanged(this, this->objectName().toUtf8().data(), accessible); };
	void Signal_PreSharedKeyAuthenticationRequired(QNetworkReply * reply, QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired(this, this->objectName().toUtf8().data(), reply, authenticator); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkAccessManagerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkAccessManagerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNetworkAccessManagerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QNetworkAccessManager_ProxyFactory(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->proxyFactory();
}

void* QNetworkAccessManager_NewQNetworkAccessManager(void* parent){
	return new MyQNetworkAccessManager(static_cast<QObject*>(parent));
}

void QNetworkAccessManager_ConnectAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));;
}

void QNetworkAccessManager_DisconnectAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&QNetworkAccessManager::authenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QAuthenticator *)>(&MyQNetworkAccessManager::Signal_AuthenticationRequired));;
}

void QNetworkAccessManager_AuthenticationRequired(void* ptr, void* reply, void* authenticator){
	static_cast<QNetworkAccessManager*>(ptr)->authenticationRequired(static_cast<QNetworkReply*>(reply), static_cast<QAuthenticator*>(authenticator));
}

void* QNetworkAccessManager_Cache(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->cache();
}

void QNetworkAccessManager_ClearAccessCache(void* ptr){
	static_cast<QNetworkAccessManager*>(ptr)->clearAccessCache();
}

void* QNetworkAccessManager_CookieJar(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->cookieJar();
}

void* QNetworkAccessManager_CreateRequest(void* ptr, int op, void* req, void* outgoingData){
	return static_cast<QNetworkAccessManager*>(ptr)->createRequest(static_cast<QNetworkAccessManager::Operation>(op), *static_cast<QNetworkRequest*>(req), static_cast<QIODevice*>(outgoingData));
}

void* QNetworkAccessManager_DeleteResource(void* ptr, void* request){
	return static_cast<QNetworkAccessManager*>(ptr)->deleteResource(*static_cast<QNetworkRequest*>(request));
}

void QNetworkAccessManager_ConnectEncrypted(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));;
}

void QNetworkAccessManager_DisconnectEncrypted(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::encrypted), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Encrypted));;
}

void QNetworkAccessManager_Encrypted(void* ptr, void* reply){
	static_cast<QNetworkAccessManager*>(ptr)->encrypted(static_cast<QNetworkReply*>(reply));
}

void QNetworkAccessManager_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));;
}

void QNetworkAccessManager_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *)>(&QNetworkAccessManager::finished), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *)>(&MyQNetworkAccessManager::Signal_Finished));;
}

void QNetworkAccessManager_Finished(void* ptr, void* reply){
	static_cast<QNetworkAccessManager*>(ptr)->finished(static_cast<QNetworkReply*>(reply));
}

void* QNetworkAccessManager_Get(void* ptr, void* request){
	return static_cast<QNetworkAccessManager*>(ptr)->get(*static_cast<QNetworkRequest*>(request));
}

void* QNetworkAccessManager_Head(void* ptr, void* request){
	return static_cast<QNetworkAccessManager*>(ptr)->head(*static_cast<QNetworkRequest*>(request));
}

int QNetworkAccessManager_NetworkAccessible(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->networkAccessible();
}

void QNetworkAccessManager_ConnectNetworkAccessibleChanged(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));;
}

void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&QNetworkAccessManager::networkAccessibleChanged), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkAccessManager::NetworkAccessibility)>(&MyQNetworkAccessManager::Signal_NetworkAccessibleChanged));;
}

void QNetworkAccessManager_NetworkAccessibleChanged(void* ptr, int accessible){
	static_cast<QNetworkAccessManager*>(ptr)->networkAccessibleChanged(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void* QNetworkAccessManager_Post3(void* ptr, void* request, void* multiPart){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Post(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Post2(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->post(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QNetworkAccessManager*>(ptr), static_cast<void (QNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&QNetworkAccessManager::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkAccessManager*>(ptr), static_cast<void (MyQNetworkAccessManager::*)(QNetworkReply *, QSslPreSharedKeyAuthenticator *)>(&MyQNetworkAccessManager::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkAccessManager_PreSharedKeyAuthenticationRequired(void* ptr, void* reply, void* authenticator){
	static_cast<QNetworkAccessManager*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QNetworkReply*>(reply), static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void* QNetworkAccessManager_Put2(void* ptr, void* request, void* multiPart){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QHttpMultiPart*>(multiPart));
}

void* QNetworkAccessManager_Put(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), static_cast<QIODevice*>(data));
}

void* QNetworkAccessManager_Put3(void* ptr, void* request, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->put(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(data));
}

void* QNetworkAccessManager_SendCustomRequest(void* ptr, void* request, void* verb, void* data){
	return static_cast<QNetworkAccessManager*>(ptr)->sendCustomRequest(*static_cast<QNetworkRequest*>(request), *static_cast<QByteArray*>(verb), static_cast<QIODevice*>(data));
}

void QNetworkAccessManager_SetCache(void* ptr, void* cache){
	static_cast<QNetworkAccessManager*>(ptr)->setCache(static_cast<QAbstractNetworkCache*>(cache));
}

void QNetworkAccessManager_SetConfiguration(void* ptr, void* config){
	static_cast<QNetworkAccessManager*>(ptr)->setConfiguration(*static_cast<QNetworkConfiguration*>(config));
}

void QNetworkAccessManager_SetCookieJar(void* ptr, void* cookieJar){
	static_cast<QNetworkAccessManager*>(ptr)->setCookieJar(static_cast<QNetworkCookieJar*>(cookieJar));
}

void QNetworkAccessManager_SetNetworkAccessible(void* ptr, int accessible){
	static_cast<QNetworkAccessManager*>(ptr)->setNetworkAccessible(static_cast<QNetworkAccessManager::NetworkAccessibility>(accessible));
}

void QNetworkAccessManager_SetProxy(void* ptr, void* proxy){
	static_cast<QNetworkAccessManager*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(proxy));
}

void QNetworkAccessManager_SetProxyFactory(void* ptr, void* factory){
	static_cast<QNetworkAccessManager*>(ptr)->setProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

char* QNetworkAccessManager_SupportedSchemes(void* ptr){
	return static_cast<QNetworkAccessManager*>(ptr)->supportedSchemes().join(",,,").toUtf8().data();
}

void QNetworkAccessManager_DestroyQNetworkAccessManager(void* ptr){
	static_cast<QNetworkAccessManager*>(ptr)->~QNetworkAccessManager();
}

void QNetworkAccessManager_TimerEvent(void* ptr, void* event){
	static_cast<MyQNetworkAccessManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkAccessManager_TimerEventDefault(void* ptr, void* event){
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkAccessManager_ChildEvent(void* ptr, void* event){
	static_cast<MyQNetworkAccessManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkAccessManager_ChildEventDefault(void* ptr, void* event){
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkAccessManager_CustomEvent(void* ptr, void* event){
	static_cast<MyQNetworkAccessManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkAccessManager_CustomEventDefault(void* ptr, void* event){
	static_cast<QNetworkAccessManager*>(ptr)->QNetworkAccessManager::customEvent(static_cast<QEvent*>(event));
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry(){
	return new QNetworkAddressEntry();
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry2(void* other){
	return new QNetworkAddressEntry(*static_cast<QNetworkAddressEntry*>(other));
}

int QNetworkAddressEntry_PrefixLength(void* ptr){
	return static_cast<QNetworkAddressEntry*>(ptr)->prefixLength();
}

void QNetworkAddressEntry_SetBroadcast(void* ptr, void* newBroadcast){
	static_cast<QNetworkAddressEntry*>(ptr)->setBroadcast(*static_cast<QHostAddress*>(newBroadcast));
}

void QNetworkAddressEntry_SetIp(void* ptr, void* newIp){
	static_cast<QNetworkAddressEntry*>(ptr)->setIp(*static_cast<QHostAddress*>(newIp));
}

void QNetworkAddressEntry_SetNetmask(void* ptr, void* newNetmask){
	static_cast<QNetworkAddressEntry*>(ptr)->setNetmask(*static_cast<QHostAddress*>(newNetmask));
}

void QNetworkAddressEntry_SetPrefixLength(void* ptr, int length){
	static_cast<QNetworkAddressEntry*>(ptr)->setPrefixLength(length);
}

void QNetworkAddressEntry_Swap(void* ptr, void* other){
	static_cast<QNetworkAddressEntry*>(ptr)->swap(*static_cast<QNetworkAddressEntry*>(other));
}

void QNetworkAddressEntry_DestroyQNetworkAddressEntry(void* ptr){
	static_cast<QNetworkAddressEntry*>(ptr)->~QNetworkAddressEntry();
}

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData(){
	return new QNetworkCacheMetaData();
}

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData2(void* other){
	return new QNetworkCacheMetaData(*static_cast<QNetworkCacheMetaData*>(other));
}

void* QNetworkCacheMetaData_ExpirationDate(void* ptr){
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->expirationDate());
}

int QNetworkCacheMetaData_IsValid(void* ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->isValid();
}

void* QNetworkCacheMetaData_LastModified(void* ptr){
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->lastModified());
}

int QNetworkCacheMetaData_SaveToDisk(void* ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->saveToDisk();
}

void QNetworkCacheMetaData_SetExpirationDate(void* ptr, void* dateTime){
	static_cast<QNetworkCacheMetaData*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetLastModified(void* ptr, void* dateTime){
	static_cast<QNetworkCacheMetaData*>(ptr)->setLastModified(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetSaveToDisk(void* ptr, int allow){
	static_cast<QNetworkCacheMetaData*>(ptr)->setSaveToDisk(allow != 0);
}

void QNetworkCacheMetaData_SetUrl(void* ptr, void* url){
	static_cast<QNetworkCacheMetaData*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkCacheMetaData_Swap(void* ptr, void* other){
	static_cast<QNetworkCacheMetaData*>(ptr)->swap(*static_cast<QNetworkCacheMetaData*>(other));
}

void* QNetworkCacheMetaData_Url(void* ptr){
	return new QUrl(static_cast<QNetworkCacheMetaData*>(ptr)->url());
}

void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(void* ptr){
	static_cast<QNetworkCacheMetaData*>(ptr)->~QNetworkCacheMetaData();
}

void* QNetworkConfiguration_NewQNetworkConfiguration(){
	return new QNetworkConfiguration();
}

void* QNetworkConfiguration_NewQNetworkConfiguration2(void* other){
	return new QNetworkConfiguration(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_BearerType(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerType();
}

int QNetworkConfiguration_BearerTypeFamily(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeFamily();
}

char* QNetworkConfiguration_BearerTypeName(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeName().toUtf8().data();
}

char* QNetworkConfiguration_Identifier(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->identifier().toUtf8().data();
}

int QNetworkConfiguration_IsRoamingAvailable(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->isRoamingAvailable();
}

int QNetworkConfiguration_IsValid(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->isValid();
}

char* QNetworkConfiguration_Name(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->name().toUtf8().data();
}

int QNetworkConfiguration_Purpose(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->purpose();
}

int QNetworkConfiguration_State(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->state();
}

void QNetworkConfiguration_Swap(void* ptr, void* other){
	static_cast<QNetworkConfiguration*>(ptr)->swap(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_Type(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->type();
}

void QNetworkConfiguration_DestroyQNetworkConfiguration(void* ptr){
	static_cast<QNetworkConfiguration*>(ptr)->~QNetworkConfiguration();
}

class MyQNetworkConfigurationManager: public QNetworkConfigurationManager {
public:
	MyQNetworkConfigurationManager(QObject *parent) : QNetworkConfigurationManager(parent) {};
	void Signal_OnlineStateChanged(bool isOnline) { callbackQNetworkConfigurationManagerOnlineStateChanged(this, this->objectName().toUtf8().data(), isOnline); };
	void Signal_UpdateCompleted() { callbackQNetworkConfigurationManagerUpdateCompleted(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkConfigurationManagerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkConfigurationManagerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNetworkConfigurationManagerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QNetworkConfigurationManager_NewQNetworkConfigurationManager(void* parent){
	return new MyQNetworkConfigurationManager(static_cast<QObject*>(parent));
}

int QNetworkConfigurationManager_Capabilities(void* ptr){
	return static_cast<QNetworkConfigurationManager*>(ptr)->capabilities();
}

int QNetworkConfigurationManager_IsOnline(void* ptr){
	return static_cast<QNetworkConfigurationManager*>(ptr)->isOnline();
}

void QNetworkConfigurationManager_ConnectOnlineStateChanged(void* ptr){
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));;
}

void QNetworkConfigurationManager_DisconnectOnlineStateChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));;
}

void QNetworkConfigurationManager_OnlineStateChanged(void* ptr, int isOnline){
	static_cast<QNetworkConfigurationManager*>(ptr)->onlineStateChanged(isOnline != 0);
}

void QNetworkConfigurationManager_ConnectUpdateCompleted(void* ptr){
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));;
}

void QNetworkConfigurationManager_DisconnectUpdateCompleted(void* ptr){
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));;
}

void QNetworkConfigurationManager_UpdateCompleted(void* ptr){
	static_cast<QNetworkConfigurationManager*>(ptr)->updateCompleted();
}

void QNetworkConfigurationManager_UpdateConfigurations(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkConfigurationManager*>(ptr), "updateConfigurations");
}

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(void* ptr){
	static_cast<QNetworkConfigurationManager*>(ptr)->~QNetworkConfigurationManager();
}

void QNetworkConfigurationManager_TimerEvent(void* ptr, void* event){
	static_cast<MyQNetworkConfigurationManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkConfigurationManager_TimerEventDefault(void* ptr, void* event){
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkConfigurationManager_ChildEvent(void* ptr, void* event){
	static_cast<MyQNetworkConfigurationManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkConfigurationManager_ChildEventDefault(void* ptr, void* event){
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkConfigurationManager_CustomEvent(void* ptr, void* event){
	static_cast<MyQNetworkConfigurationManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkConfigurationManager_CustomEventDefault(void* ptr, void* event){
	static_cast<QNetworkConfigurationManager*>(ptr)->QNetworkConfigurationManager::customEvent(static_cast<QEvent*>(event));
}

void* QNetworkCookie_NewQNetworkCookie(void* name, void* value){
	return new QNetworkCookie(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

void* QNetworkCookie_NewQNetworkCookie2(void* other){
	return new QNetworkCookie(*static_cast<QNetworkCookie*>(other));
}

char* QNetworkCookie_Domain(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->domain().toUtf8().data();
}

void* QNetworkCookie_ExpirationDate(void* ptr){
	return new QDateTime(static_cast<QNetworkCookie*>(ptr)->expirationDate());
}

int QNetworkCookie_HasSameIdentifier(void* ptr, void* other){
	return static_cast<QNetworkCookie*>(ptr)->hasSameIdentifier(*static_cast<QNetworkCookie*>(other));
}

int QNetworkCookie_IsHttpOnly(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->isHttpOnly();
}

int QNetworkCookie_IsSecure(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->isSecure();
}

int QNetworkCookie_IsSessionCookie(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->isSessionCookie();
}

void* QNetworkCookie_Name(void* ptr){
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->name());
}

void QNetworkCookie_Normalize(void* ptr, void* url){
	static_cast<QNetworkCookie*>(ptr)->normalize(*static_cast<QUrl*>(url));
}

char* QNetworkCookie_Path(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->path().toUtf8().data();
}

void QNetworkCookie_SetDomain(void* ptr, char* domain){
	static_cast<QNetworkCookie*>(ptr)->setDomain(QString(domain));
}

void QNetworkCookie_SetExpirationDate(void* ptr, void* date){
	static_cast<QNetworkCookie*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(date));
}

void QNetworkCookie_SetHttpOnly(void* ptr, int enable){
	static_cast<QNetworkCookie*>(ptr)->setHttpOnly(enable != 0);
}

void QNetworkCookie_SetName(void* ptr, void* cookieName){
	static_cast<QNetworkCookie*>(ptr)->setName(*static_cast<QByteArray*>(cookieName));
}

void QNetworkCookie_SetPath(void* ptr, char* path){
	static_cast<QNetworkCookie*>(ptr)->setPath(QString(path));
}

void QNetworkCookie_SetSecure(void* ptr, int enable){
	static_cast<QNetworkCookie*>(ptr)->setSecure(enable != 0);
}

void QNetworkCookie_SetValue(void* ptr, void* value){
	static_cast<QNetworkCookie*>(ptr)->setValue(*static_cast<QByteArray*>(value));
}

void QNetworkCookie_Swap(void* ptr, void* other){
	static_cast<QNetworkCookie*>(ptr)->swap(*static_cast<QNetworkCookie*>(other));
}

void* QNetworkCookie_ToRawForm(void* ptr, int form){
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->toRawForm(static_cast<QNetworkCookie::RawForm>(form)));
}

void* QNetworkCookie_Value(void* ptr){
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->value());
}

void QNetworkCookie_DestroyQNetworkCookie(void* ptr){
	static_cast<QNetworkCookie*>(ptr)->~QNetworkCookie();
}

class MyQNetworkCookieJar: public QNetworkCookieJar {
public:
	MyQNetworkCookieJar(QObject *parent) : QNetworkCookieJar(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQNetworkCookieJarTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkCookieJarChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNetworkCookieJarCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QNetworkCookieJar_NewQNetworkCookieJar(void* parent){
	return new MyQNetworkCookieJar(static_cast<QObject*>(parent));
}

int QNetworkCookieJar_DeleteCookie(void* ptr, void* cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_InsertCookie(void* ptr, void* cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_UpdateCookie(void* ptr, void* cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_ValidateCookie(void* ptr, void* cookie, void* url){
	return static_cast<QNetworkCookieJar*>(ptr)->validateCookie(*static_cast<QNetworkCookie*>(cookie), *static_cast<QUrl*>(url));
}

void QNetworkCookieJar_DestroyQNetworkCookieJar(void* ptr){
	static_cast<QNetworkCookieJar*>(ptr)->~QNetworkCookieJar();
}

void QNetworkCookieJar_TimerEvent(void* ptr, void* event){
	static_cast<MyQNetworkCookieJar*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkCookieJar_TimerEventDefault(void* ptr, void* event){
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkCookieJar_ChildEvent(void* ptr, void* event){
	static_cast<MyQNetworkCookieJar*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkCookieJar_ChildEventDefault(void* ptr, void* event){
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkCookieJar_CustomEvent(void* ptr, void* event){
	static_cast<MyQNetworkCookieJar*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkCookieJar_CustomEventDefault(void* ptr, void* event){
	static_cast<QNetworkCookieJar*>(ptr)->QNetworkCookieJar::customEvent(static_cast<QEvent*>(event));
}

class MyQNetworkDiskCache: public QNetworkDiskCache {
public:
	MyQNetworkDiskCache(QObject *parent) : QNetworkDiskCache(parent) {};
	void clear() { if (!callbackQNetworkDiskCacheClear(this, this->objectName().toUtf8().data())) { QNetworkDiskCache::clear(); }; };
	void timerEvent(QTimerEvent * event) { callbackQNetworkDiskCacheTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkDiskCacheChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNetworkDiskCacheCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QNetworkDiskCache_NewQNetworkDiskCache(void* parent){
	return new MyQNetworkDiskCache(static_cast<QObject*>(parent));
}

char* QNetworkDiskCache_CacheDirectory(void* ptr){
	return static_cast<QNetworkDiskCache*>(ptr)->cacheDirectory().toUtf8().data();
}

long long QNetworkDiskCache_CacheSize(void* ptr){
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->cacheSize());
}

void QNetworkDiskCache_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQNetworkDiskCache*>(ptr), "clear");
}

void QNetworkDiskCache_ClearDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkDiskCache*>(ptr), "clear");
}

void* QNetworkDiskCache_Data(void* ptr, void* url){
	return static_cast<QNetworkDiskCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

long long QNetworkDiskCache_Expire(void* ptr){
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->expire());
}

long long QNetworkDiskCache_MaximumCacheSize(void* ptr){
	return static_cast<long long>(static_cast<QNetworkDiskCache*>(ptr)->maximumCacheSize());
}

void* QNetworkDiskCache_Prepare(void* ptr, void* metaData){
	return static_cast<QNetworkDiskCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_SetCacheDirectory(void* ptr, char* cacheDir){
	static_cast<QNetworkDiskCache*>(ptr)->setCacheDirectory(QString(cacheDir));
}

void QNetworkDiskCache_SetMaximumCacheSize(void* ptr, long long size){
	static_cast<QNetworkDiskCache*>(ptr)->setMaximumCacheSize(static_cast<long long>(size));
}

void QNetworkDiskCache_DestroyQNetworkDiskCache(void* ptr){
	static_cast<QNetworkDiskCache*>(ptr)->~QNetworkDiskCache();
}

void QNetworkDiskCache_TimerEvent(void* ptr, void* event){
	static_cast<MyQNetworkDiskCache*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkDiskCache_TimerEventDefault(void* ptr, void* event){
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkDiskCache_ChildEvent(void* ptr, void* event){
	static_cast<MyQNetworkDiskCache*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkDiskCache_ChildEventDefault(void* ptr, void* event){
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkDiskCache_CustomEvent(void* ptr, void* event){
	static_cast<MyQNetworkDiskCache*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkDiskCache_CustomEventDefault(void* ptr, void* event){
	static_cast<QNetworkDiskCache*>(ptr)->QNetworkDiskCache::customEvent(static_cast<QEvent*>(event));
}

void* QNetworkInterface_NewQNetworkInterface(){
	return new QNetworkInterface();
}

void* QNetworkInterface_NewQNetworkInterface2(void* other){
	return new QNetworkInterface(*static_cast<QNetworkInterface*>(other));
}

int QNetworkInterface_Flags(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->flags();
}

char* QNetworkInterface_HardwareAddress(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->hardwareAddress().toUtf8().data();
}

char* QNetworkInterface_HumanReadableName(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->humanReadableName().toUtf8().data();
}

int QNetworkInterface_Index(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->index();
}

int QNetworkInterface_IsValid(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->isValid();
}

char* QNetworkInterface_Name(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->name().toUtf8().data();
}

void QNetworkInterface_Swap(void* ptr, void* other){
	static_cast<QNetworkInterface*>(ptr)->swap(*static_cast<QNetworkInterface*>(other));
}

void QNetworkInterface_DestroyQNetworkInterface(void* ptr){
	static_cast<QNetworkInterface*>(ptr)->~QNetworkInterface();
}

void* QNetworkProxy_NewQNetworkProxy(){
	return new QNetworkProxy();
}

void* QNetworkProxy_NewQNetworkProxy3(void* other){
	return new QNetworkProxy(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Capabilities(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->capabilities();
}

int QNetworkProxy_HasRawHeader(void* ptr, void* headerName){
	return static_cast<QNetworkProxy*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkProxy_Header(void* ptr, int header){
	return new QVariant(static_cast<QNetworkProxy*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

char* QNetworkProxy_HostName(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->hostName().toUtf8().data();
}

int QNetworkProxy_IsCachingProxy(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->isCachingProxy();
}

int QNetworkProxy_IsTransparentProxy(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->isTransparentProxy();
}

char* QNetworkProxy_Password(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->password().toUtf8().data();
}

void* QNetworkProxy_RawHeader(void* ptr, void* headerName){
	return new QByteArray(static_cast<QNetworkProxy*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

void QNetworkProxy_QNetworkProxy_SetApplicationProxy(void* networkProxy){
	QNetworkProxy::setApplicationProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QNetworkProxy_SetCapabilities(void* ptr, int capabilities){
	static_cast<QNetworkProxy*>(ptr)->setCapabilities(static_cast<QNetworkProxy::Capability>(capabilities));
}

void QNetworkProxy_SetHeader(void* ptr, int header, void* value){
	static_cast<QNetworkProxy*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkProxy_SetHostName(void* ptr, char* hostName){
	static_cast<QNetworkProxy*>(ptr)->setHostName(QString(hostName));
}

void QNetworkProxy_SetPassword(void* ptr, char* password){
	static_cast<QNetworkProxy*>(ptr)->setPassword(QString(password));
}

void QNetworkProxy_SetRawHeader(void* ptr, void* headerName, void* headerValue){
	static_cast<QNetworkProxy*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkProxy_SetType(void* ptr, int ty){
	static_cast<QNetworkProxy*>(ptr)->setType(static_cast<QNetworkProxy::ProxyType>(ty));
}

void QNetworkProxy_SetUser(void* ptr, char* user){
	static_cast<QNetworkProxy*>(ptr)->setUser(QString(user));
}

void QNetworkProxy_Swap(void* ptr, void* other){
	static_cast<QNetworkProxy*>(ptr)->swap(*static_cast<QNetworkProxy*>(other));
}

int QNetworkProxy_Type(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->type();
}

char* QNetworkProxy_User(void* ptr){
	return static_cast<QNetworkProxy*>(ptr)->user().toUtf8().data();
}

void QNetworkProxy_DestroyQNetworkProxy(void* ptr){
	static_cast<QNetworkProxy*>(ptr)->~QNetworkProxy();
}

class MyQNetworkProxyFactory: public QNetworkProxyFactory {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(void* factory){
	QNetworkProxyFactory::setApplicationProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

void QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(int enable){
	QNetworkProxyFactory::setUseSystemConfiguration(enable != 0);
}

void QNetworkProxyFactory_DestroyQNetworkProxyFactory(void* ptr){
	static_cast<QNetworkProxyFactory*>(ptr)->~QNetworkProxyFactory();
}

char* QNetworkProxyFactory_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQNetworkProxyFactory*>(static_cast<QNetworkProxyFactory*>(ptr))) {
		return static_cast<MyQNetworkProxyFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QNetworkProxyFactory_BASE").toUtf8().data();
}

void QNetworkProxyFactory_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQNetworkProxyFactory*>(static_cast<QNetworkProxyFactory*>(ptr))) {
		static_cast<MyQNetworkProxyFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery(){
	return new QNetworkProxyQuery();
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery7(void* networkConfiguration, char* hostname, int port, char* protocolTag, int queryType){
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery6(void* networkConfiguration, void* requestUrl, int queryType){
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), *static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery5(void* other){
	return new QNetworkProxyQuery(*static_cast<QNetworkProxyQuery*>(other));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, int queryType){
	return new QNetworkProxyQuery(QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery2(void* requestUrl, int queryType){
	return new QNetworkProxyQuery(*static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

int QNetworkProxyQuery_LocalPort(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->localPort();
}

char* QNetworkProxyQuery_PeerHostName(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->peerHostName().toUtf8().data();
}

int QNetworkProxyQuery_PeerPort(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->peerPort();
}

char* QNetworkProxyQuery_ProtocolTag(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->protocolTag().toUtf8().data();
}

int QNetworkProxyQuery_QueryType(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->queryType();
}

void QNetworkProxyQuery_SetLocalPort(void* ptr, int port){
	static_cast<QNetworkProxyQuery*>(ptr)->setLocalPort(port);
}

void QNetworkProxyQuery_SetNetworkConfiguration(void* ptr, void* networkConfiguration){
	static_cast<QNetworkProxyQuery*>(ptr)->setNetworkConfiguration(*static_cast<QNetworkConfiguration*>(networkConfiguration));
}

void QNetworkProxyQuery_SetPeerHostName(void* ptr, char* hostname){
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerHostName(QString(hostname));
}

void QNetworkProxyQuery_SetPeerPort(void* ptr, int port){
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerPort(port);
}

void QNetworkProxyQuery_SetProtocolTag(void* ptr, char* protocolTag){
	static_cast<QNetworkProxyQuery*>(ptr)->setProtocolTag(QString(protocolTag));
}

void QNetworkProxyQuery_SetQueryType(void* ptr, int ty){
	static_cast<QNetworkProxyQuery*>(ptr)->setQueryType(static_cast<QNetworkProxyQuery::QueryType>(ty));
}

void QNetworkProxyQuery_SetUrl(void* ptr, void* url){
	static_cast<QNetworkProxyQuery*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkProxyQuery_Swap(void* ptr, void* other){
	static_cast<QNetworkProxyQuery*>(ptr)->swap(*static_cast<QNetworkProxyQuery*>(other));
}

void* QNetworkProxyQuery_Url(void* ptr){
	return new QUrl(static_cast<QNetworkProxyQuery*>(ptr)->url());
}

void QNetworkProxyQuery_DestroyQNetworkProxyQuery(void* ptr){
	static_cast<QNetworkProxyQuery*>(ptr)->~QNetworkProxyQuery();
}

class MyQNetworkReply: public QNetworkReply {
public:
	void close() { callbackQNetworkReplyClose(this, this->objectName().toUtf8().data()); };
	void Signal_DownloadProgress(qint64 bytesReceived, qint64 bytesTotal) { callbackQNetworkReplyDownloadProgress(this, this->objectName().toUtf8().data(), static_cast<long long>(bytesReceived), static_cast<long long>(bytesTotal)); };
	void Signal_Encrypted() { callbackQNetworkReplyEncrypted(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QNetworkReply::NetworkError code) { callbackQNetworkReplyError2(this, this->objectName().toUtf8().data(), code); };
	void Signal_Finished() { callbackQNetworkReplyFinished(this, this->objectName().toUtf8().data()); };
	void ignoreSslErrors() { if (!callbackQNetworkReplyIgnoreSslErrors(this, this->objectName().toUtf8().data())) { QNetworkReply::ignoreSslErrors(); }; };
	void Signal_MetaDataChanged() { callbackQNetworkReplyMetaDataChanged(this, this->objectName().toUtf8().data()); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQNetworkReplyPreSharedKeyAuthenticationRequired(this, this->objectName().toUtf8().data(), authenticator); };
	void setReadBufferSize(qint64 size) { callbackQNetworkReplySetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void Signal_UploadProgress(qint64 bytesSent, qint64 bytesTotal) { callbackQNetworkReplyUploadProgress(this, this->objectName().toUtf8().data(), static_cast<long long>(bytesSent), static_cast<long long>(bytesTotal)); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkReplyTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkReplyChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNetworkReplyCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QNetworkReply_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "abort");
}

void* QNetworkReply_Attribute(void* ptr, int code){
	return new QVariant(static_cast<QNetworkReply*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code)));
}

void QNetworkReply_Close(void* ptr){
	static_cast<MyQNetworkReply*>(ptr)->close();
}

void QNetworkReply_CloseDefault(void* ptr){
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::close();
}

void QNetworkReply_ConnectDownloadProgress(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::downloadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_DownloadProgress));;
}

void QNetworkReply_DisconnectDownloadProgress(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::downloadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_DownloadProgress));;
}

void QNetworkReply_DownloadProgress(void* ptr, long long bytesReceived, long long bytesTotal){
	static_cast<QNetworkReply*>(ptr)->downloadProgress(static_cast<long long>(bytesReceived), static_cast<long long>(bytesTotal));
}

void QNetworkReply_ConnectEncrypted(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));;
}

void QNetworkReply_DisconnectEncrypted(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::encrypted), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Encrypted));;
}

void QNetworkReply_Encrypted(void* ptr){
	static_cast<QNetworkReply*>(ptr)->encrypted();
}

void QNetworkReply_ConnectError2(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QNetworkReply::NetworkError)>(&QNetworkReply::error), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QNetworkReply::NetworkError)>(&MyQNetworkReply::Signal_Error2));;
}

void QNetworkReply_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QNetworkReply::NetworkError)>(&QNetworkReply::error), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QNetworkReply::NetworkError)>(&MyQNetworkReply::Signal_Error2));;
}

void QNetworkReply_Error2(void* ptr, int code){
	static_cast<QNetworkReply*>(ptr)->error(static_cast<QNetworkReply::NetworkError>(code));
}

int QNetworkReply_Error(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->error();
}

void QNetworkReply_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));;
}

void QNetworkReply_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::finished), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_Finished));;
}

void QNetworkReply_Finished(void* ptr){
	static_cast<QNetworkReply*>(ptr)->finished();
}

int QNetworkReply_HasRawHeader(void* ptr, void* headerName){
	return static_cast<QNetworkReply*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkReply_Header(void* ptr, int header){
	return new QVariant(static_cast<QNetworkReply*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

void QNetworkReply_IgnoreSslErrors(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQNetworkReply*>(ptr), "ignoreSslErrors");
}

void QNetworkReply_IgnoreSslErrorsDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkReply*>(ptr), "ignoreSslErrors");
}

int QNetworkReply_IsFinished(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->isFinished();
}

int QNetworkReply_IsRunning(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->isRunning();
}

void* QNetworkReply_Manager(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->manager();
}

void QNetworkReply_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));;
}

void QNetworkReply_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)()>(&QNetworkReply::metaDataChanged), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)()>(&MyQNetworkReply::Signal_MetaDataChanged));;
}

void QNetworkReply_MetaDataChanged(void* ptr){
	static_cast<QNetworkReply*>(ptr)->metaDataChanged();
}

int QNetworkReply_Operation(void* ptr){
	return static_cast<QNetworkReply*>(ptr)->operation();
}

void QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&QNetworkReply::preSharedKeyAuthenticationRequired), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(QSslPreSharedKeyAuthenticator *)>(&MyQNetworkReply::Signal_PreSharedKeyAuthenticationRequired));;
}

void QNetworkReply_PreSharedKeyAuthenticationRequired(void* ptr, void* authenticator){
	static_cast<QNetworkReply*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void* QNetworkReply_RawHeader(void* ptr, void* headerName){
	return new QByteArray(static_cast<QNetworkReply*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

long long QNetworkReply_ReadBufferSize(void* ptr){
	return static_cast<long long>(static_cast<QNetworkReply*>(ptr)->readBufferSize());
}

void QNetworkReply_SetReadBufferSize(void* ptr, long long size){
	static_cast<MyQNetworkReply*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QNetworkReply_SetReadBufferSizeDefault(void* ptr, long long size){
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::setReadBufferSize(static_cast<long long>(size));
}

void QNetworkReply_SetSslConfiguration(void* ptr, void* config){
	static_cast<QNetworkReply*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkReply_ConnectUploadProgress(void* ptr){
	QObject::connect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::uploadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_UploadProgress));;
}

void QNetworkReply_DisconnectUploadProgress(void* ptr){
	QObject::disconnect(static_cast<QNetworkReply*>(ptr), static_cast<void (QNetworkReply::*)(qint64, qint64)>(&QNetworkReply::uploadProgress), static_cast<MyQNetworkReply*>(ptr), static_cast<void (MyQNetworkReply::*)(qint64, qint64)>(&MyQNetworkReply::Signal_UploadProgress));;
}

void QNetworkReply_UploadProgress(void* ptr, long long bytesSent, long long bytesTotal){
	static_cast<QNetworkReply*>(ptr)->uploadProgress(static_cast<long long>(bytesSent), static_cast<long long>(bytesTotal));
}

void* QNetworkReply_Url(void* ptr){
	return new QUrl(static_cast<QNetworkReply*>(ptr)->url());
}

void QNetworkReply_DestroyQNetworkReply(void* ptr){
	static_cast<QNetworkReply*>(ptr)->~QNetworkReply();
}

void QNetworkReply_TimerEvent(void* ptr, void* event){
	static_cast<MyQNetworkReply*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkReply_TimerEventDefault(void* ptr, void* event){
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkReply_ChildEvent(void* ptr, void* event){
	static_cast<MyQNetworkReply*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkReply_ChildEventDefault(void* ptr, void* event){
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkReply_CustomEvent(void* ptr, void* event){
	static_cast<MyQNetworkReply*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkReply_CustomEventDefault(void* ptr, void* event){
	static_cast<QNetworkReply*>(ptr)->QNetworkReply::customEvent(static_cast<QEvent*>(event));
}

void* QNetworkRequest_NewQNetworkRequest2(void* other){
	return new QNetworkRequest(*static_cast<QNetworkRequest*>(other));
}

void* QNetworkRequest_NewQNetworkRequest(void* url){
	return new QNetworkRequest(*static_cast<QUrl*>(url));
}

void* QNetworkRequest_Attribute(void* ptr, int code, void* defaultValue){
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->attribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

int QNetworkRequest_HasRawHeader(void* ptr, void* headerName){
	return static_cast<QNetworkRequest*>(ptr)->hasRawHeader(*static_cast<QByteArray*>(headerName));
}

void* QNetworkRequest_Header(void* ptr, int header){
	return new QVariant(static_cast<QNetworkRequest*>(ptr)->header(static_cast<QNetworkRequest::KnownHeaders>(header)));
}

void* QNetworkRequest_OriginatingObject(void* ptr){
	return static_cast<QNetworkRequest*>(ptr)->originatingObject();
}

int QNetworkRequest_Priority(void* ptr){
	return static_cast<QNetworkRequest*>(ptr)->priority();
}

void* QNetworkRequest_RawHeader(void* ptr, void* headerName){
	return new QByteArray(static_cast<QNetworkRequest*>(ptr)->rawHeader(*static_cast<QByteArray*>(headerName)));
}

void QNetworkRequest_SetAttribute(void* ptr, int code, void* value){
	static_cast<QNetworkRequest*>(ptr)->setAttribute(static_cast<QNetworkRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetHeader(void* ptr, int header, void* value){
	static_cast<QNetworkRequest*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QNetworkRequest_SetOriginatingObject(void* ptr, void* object){
	static_cast<QNetworkRequest*>(ptr)->setOriginatingObject(static_cast<QObject*>(object));
}

void QNetworkRequest_SetPriority(void* ptr, int priority){
	static_cast<QNetworkRequest*>(ptr)->setPriority(static_cast<QNetworkRequest::Priority>(priority));
}

void QNetworkRequest_SetRawHeader(void* ptr, void* headerName, void* headerValue){
	static_cast<QNetworkRequest*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QNetworkRequest_SetSslConfiguration(void* ptr, void* config){
	static_cast<QNetworkRequest*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(config));
}

void QNetworkRequest_SetUrl(void* ptr, void* url){
	static_cast<QNetworkRequest*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkRequest_Swap(void* ptr, void* other){
	static_cast<QNetworkRequest*>(ptr)->swap(*static_cast<QNetworkRequest*>(other));
}

void* QNetworkRequest_Url(void* ptr){
	return new QUrl(static_cast<QNetworkRequest*>(ptr)->url());
}

void QNetworkRequest_DestroyQNetworkRequest(void* ptr){
	static_cast<QNetworkRequest*>(ptr)->~QNetworkRequest();
}

class MyQNetworkSession: public QNetworkSession {
public:
	MyQNetworkSession(const QNetworkConfiguration &connectionConfig, QObject *parent) : QNetworkSession(connectionConfig, parent) {};
	void Signal_Closed() { callbackQNetworkSessionClosed(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QNetworkSession::SessionError error) { callbackQNetworkSessionError2(this, this->objectName().toUtf8().data(), error); };
	void Signal_NewConfigurationActivated() { callbackQNetworkSessionNewConfigurationActivated(this, this->objectName().toUtf8().data()); };
	void Signal_Opened() { callbackQNetworkSessionOpened(this, this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QNetworkSession::State state) { callbackQNetworkSessionStateChanged(this, this->objectName().toUtf8().data(), state); };
	void Signal_UsagePoliciesChanged(QNetworkSession::UsagePolicies usagePolicies) { callbackQNetworkSessionUsagePoliciesChanged(this, this->objectName().toUtf8().data(), usagePolicies); };
	void timerEvent(QTimerEvent * event) { callbackQNetworkSessionTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNetworkSessionChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNetworkSessionCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QNetworkSession_NewQNetworkSession(void* connectionConfig, void* parent){
	return new MyQNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QObject*>(parent));
}

void QNetworkSession_Accept(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "accept");
}

void QNetworkSession_Close(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "close");
}

void QNetworkSession_ConnectClosed(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));;
}

void QNetworkSession_DisconnectClosed(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));;
}

void QNetworkSession_Closed(void* ptr){
	static_cast<QNetworkSession*>(ptr)->closed();
}

void QNetworkSession_ConnectError2(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::SessionError)>(&QNetworkSession::error), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::SessionError)>(&MyQNetworkSession::Signal_Error2));;
}

void QNetworkSession_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::SessionError)>(&QNetworkSession::error), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::SessionError)>(&MyQNetworkSession::Signal_Error2));;
}

void QNetworkSession_Error2(void* ptr, int error){
	static_cast<QNetworkSession*>(ptr)->error(static_cast<QNetworkSession::SessionError>(error));
}

int QNetworkSession_Error(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->error();
}

char* QNetworkSession_ErrorString(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->errorString().toUtf8().data();
}

void QNetworkSession_Ignore(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "ignore");
}

int QNetworkSession_IsOpen(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->isOpen();
}

void QNetworkSession_Migrate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "migrate");
}

void QNetworkSession_ConnectNewConfigurationActivated(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));;
}

void QNetworkSession_DisconnectNewConfigurationActivated(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));;
}

void QNetworkSession_NewConfigurationActivated(void* ptr){
	static_cast<QNetworkSession*>(ptr)->newConfigurationActivated();
}

void QNetworkSession_Open(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "open");
}

void QNetworkSession_ConnectOpened(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));;
}

void QNetworkSession_DisconnectOpened(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));;
}

void QNetworkSession_Opened(void* ptr){
	static_cast<QNetworkSession*>(ptr)->opened();
}

void QNetworkSession_Reject(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "reject");
}

void* QNetworkSession_SessionProperty(void* ptr, char* key){
	return new QVariant(static_cast<QNetworkSession*>(ptr)->sessionProperty(QString(key)));
}

void QNetworkSession_SetSessionProperty(void* ptr, char* key, void* value){
	static_cast<QNetworkSession*>(ptr)->setSessionProperty(QString(key), *static_cast<QVariant*>(value));
}

int QNetworkSession_State(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->state();
}

void QNetworkSession_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));;
}

void QNetworkSession_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));;
}

void QNetworkSession_StateChanged(void* ptr, int state){
	static_cast<QNetworkSession*>(ptr)->stateChanged(static_cast<QNetworkSession::State>(state));
}

void QNetworkSession_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "stop");
}

int QNetworkSession_UsagePolicies(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->usagePolicies();
}

void QNetworkSession_ConnectUsagePoliciesChanged(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));;
}

void QNetworkSession_DisconnectUsagePoliciesChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));;
}

void QNetworkSession_UsagePoliciesChanged(void* ptr, int usagePolicies){
	static_cast<QNetworkSession*>(ptr)->usagePoliciesChanged(static_cast<QNetworkSession::UsagePolicy>(usagePolicies));
}

int QNetworkSession_WaitForOpened(void* ptr, int msecs){
	return static_cast<QNetworkSession*>(ptr)->waitForOpened(msecs);
}

void QNetworkSession_DestroyQNetworkSession(void* ptr){
	static_cast<QNetworkSession*>(ptr)->~QNetworkSession();
}

void QNetworkSession_TimerEvent(void* ptr, void* event){
	static_cast<MyQNetworkSession*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkSession_TimerEventDefault(void* ptr, void* event){
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNetworkSession_ChildEvent(void* ptr, void* event){
	static_cast<MyQNetworkSession*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkSession_ChildEventDefault(void* ptr, void* event){
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::childEvent(static_cast<QChildEvent*>(event));
}

void QNetworkSession_CustomEvent(void* ptr, void* event){
	static_cast<MyQNetworkSession*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNetworkSession_CustomEventDefault(void* ptr, void* event){
	static_cast<QNetworkSession*>(ptr)->QNetworkSession::customEvent(static_cast<QEvent*>(event));
}

void* QSslCertificate_NewQSslCertificate3(void* other){
	return new QSslCertificate(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_Clear(void* ptr){
	static_cast<QSslCertificate*>(ptr)->clear();
}

void* QSslCertificate_Digest(void* ptr, int algorithm){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->digest(static_cast<QCryptographicHash::Algorithm>(algorithm)));
}

int QSslCertificate_IsBlacklisted(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->isBlacklisted();
}

void QSslCertificate_Swap(void* ptr, void* other){
	static_cast<QSslCertificate*>(ptr)->swap(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_DestroyQSslCertificate(void* ptr){
	static_cast<QSslCertificate*>(ptr)->~QSslCertificate();
}

void* QSslCertificate_EffectiveDate(void* ptr){
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->effectiveDate());
}

void* QSslCertificate_ExpiryDate(void* ptr){
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->expiryDate());
}

int QSslCertificate_IsNull(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->isNull();
}

int QSslCertificate_IsSelfSigned(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->isSelfSigned();
}

char* QSslCertificate_IssuerInfo(void* ptr, int subject){
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join(",,,").toUtf8().data();
}

char* QSslCertificate_IssuerInfo2(void* ptr, void* attribute){
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(*static_cast<QByteArray*>(attribute)).join(",,,").toUtf8().data();
}

void* QSslCertificate_SerialNumber(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->serialNumber());
}

char* QSslCertificate_SubjectInfo(void* ptr, int subject){
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join(",,,").toUtf8().data();
}

char* QSslCertificate_SubjectInfo2(void* ptr, void* attribute){
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(*static_cast<QByteArray*>(attribute)).join(",,,").toUtf8().data();
}

void* QSslCertificate_ToDer(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toDer());
}

void* QSslCertificate_ToPem(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toPem());
}

char* QSslCertificate_ToText(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->toText().toUtf8().data();
}

void* QSslCertificate_Version(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->version());
}

void* QSslCertificateExtension_NewQSslCertificateExtension(){
	return new QSslCertificateExtension();
}

void* QSslCertificateExtension_NewQSslCertificateExtension2(void* other){
	return new QSslCertificateExtension(*static_cast<QSslCertificateExtension*>(other));
}

int QSslCertificateExtension_IsCritical(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->isCritical();
}

int QSslCertificateExtension_IsSupported(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->isSupported();
}

char* QSslCertificateExtension_Name(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->name().toUtf8().data();
}

char* QSslCertificateExtension_Oid(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->oid().toUtf8().data();
}

void QSslCertificateExtension_Swap(void* ptr, void* other){
	static_cast<QSslCertificateExtension*>(ptr)->swap(*static_cast<QSslCertificateExtension*>(other));
}

void* QSslCertificateExtension_Value(void* ptr){
	return new QVariant(static_cast<QSslCertificateExtension*>(ptr)->value());
}

void QSslCertificateExtension_DestroyQSslCertificateExtension(void* ptr){
	static_cast<QSslCertificateExtension*>(ptr)->~QSslCertificateExtension();
}

void* QSslCipher_NewQSslCipher(){
	return new QSslCipher();
}

void* QSslCipher_NewQSslCipher4(void* other){
	return new QSslCipher(*static_cast<QSslCipher*>(other));
}

void* QSslCipher_NewQSslCipher2(char* name){
	return new QSslCipher(QString(name));
}

char* QSslCipher_AuthenticationMethod(void* ptr){
	return static_cast<QSslCipher*>(ptr)->authenticationMethod().toUtf8().data();
}

char* QSslCipher_EncryptionMethod(void* ptr){
	return static_cast<QSslCipher*>(ptr)->encryptionMethod().toUtf8().data();
}

int QSslCipher_IsNull(void* ptr){
	return static_cast<QSslCipher*>(ptr)->isNull();
}

char* QSslCipher_KeyExchangeMethod(void* ptr){
	return static_cast<QSslCipher*>(ptr)->keyExchangeMethod().toUtf8().data();
}

char* QSslCipher_Name(void* ptr){
	return static_cast<QSslCipher*>(ptr)->name().toUtf8().data();
}

char* QSslCipher_ProtocolString(void* ptr){
	return static_cast<QSslCipher*>(ptr)->protocolString().toUtf8().data();
}

int QSslCipher_SupportedBits(void* ptr){
	return static_cast<QSslCipher*>(ptr)->supportedBits();
}

void QSslCipher_Swap(void* ptr, void* other){
	static_cast<QSslCipher*>(ptr)->swap(*static_cast<QSslCipher*>(other));
}

int QSslCipher_UsedBits(void* ptr){
	return static_cast<QSslCipher*>(ptr)->usedBits();
}

void QSslCipher_DestroyQSslCipher(void* ptr){
	static_cast<QSslCipher*>(ptr)->~QSslCipher();
}

void* QSslConfiguration_NewQSslConfiguration(){
	return new QSslConfiguration();
}

void* QSslConfiguration_NewQSslConfiguration2(void* other){
	return new QSslConfiguration(*static_cast<QSslConfiguration*>(other));
}

int QSslConfiguration_IsNull(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->isNull();
}

void* QSslConfiguration_NextNegotiatedProtocol(void* ptr){
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->nextNegotiatedProtocol());
}

int QSslConfiguration_NextProtocolNegotiationStatus(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->nextProtocolNegotiationStatus();
}

int QSslConfiguration_PeerVerifyDepth(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyDepth();
}

int QSslConfiguration_PeerVerifyMode(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->peerVerifyMode();
}

void* QSslConfiguration_SessionTicket(void* ptr){
	return new QByteArray(static_cast<QSslConfiguration*>(ptr)->sessionTicket());
}

int QSslConfiguration_SessionTicketLifeTimeHint(void* ptr){
	return static_cast<QSslConfiguration*>(ptr)->sessionTicketLifeTimeHint();
}

void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(void* configuration){
	QSslConfiguration::setDefaultConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void QSslConfiguration_SetLocalCertificate(void* ptr, void* certificate){
	static_cast<QSslConfiguration*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslConfiguration_SetPeerVerifyDepth(void* ptr, int depth){
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslConfiguration_SetPeerVerifyMode(void* ptr, int mode){
	static_cast<QSslConfiguration*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslConfiguration_SetPrivateKey(void* ptr, void* key){
	static_cast<QSslConfiguration*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslConfiguration_SetSessionTicket(void* ptr, void* sessionTicket){
	static_cast<QSslConfiguration*>(ptr)->setSessionTicket(*static_cast<QByteArray*>(sessionTicket));
}

void QSslConfiguration_Swap(void* ptr, void* other){
	static_cast<QSslConfiguration*>(ptr)->swap(*static_cast<QSslConfiguration*>(other));
}

void QSslConfiguration_DestroyQSslConfiguration(void* ptr){
	static_cast<QSslConfiguration*>(ptr)->~QSslConfiguration();
}

void* QSslEllipticCurve_NewQSslEllipticCurve(){
	return new QSslEllipticCurve();
}

int QSslEllipticCurve_IsValid(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->isValid();
}

int QSslEllipticCurve_IsTlsNamedCurve(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->isTlsNamedCurve();
}

char* QSslEllipticCurve_LongName(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->longName().toUtf8().data();
}

char* QSslEllipticCurve_ShortName(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->shortName().toUtf8().data();
}

void* QSslError_NewQSslError(){
	return new QSslError();
}

void* QSslError_NewQSslError2(int error){
	return new QSslError(static_cast<QSslError::SslError>(error));
}

void* QSslError_NewQSslError3(int error, void* certificate){
	return new QSslError(static_cast<QSslError::SslError>(error), *static_cast<QSslCertificate*>(certificate));
}

void* QSslError_NewQSslError4(void* other){
	return new QSslError(*static_cast<QSslError*>(other));
}

int QSslError_Error(void* ptr){
	return static_cast<QSslError*>(ptr)->error();
}

char* QSslError_ErrorString(void* ptr){
	return static_cast<QSslError*>(ptr)->errorString().toUtf8().data();
}

void QSslError_Swap(void* ptr, void* other){
	static_cast<QSslError*>(ptr)->swap(*static_cast<QSslError*>(other));
}

void QSslError_DestroyQSslError(void* ptr){
	static_cast<QSslError*>(ptr)->~QSslError();
}

void* QSslKey_NewQSslKey(){
	return new QSslKey();
}

void* QSslKey_NewQSslKey5(void* other){
	return new QSslKey(*static_cast<QSslKey*>(other));
}

void QSslKey_Clear(void* ptr){
	static_cast<QSslKey*>(ptr)->clear();
}

int QSslKey_IsNull(void* ptr){
	return static_cast<QSslKey*>(ptr)->isNull();
}

int QSslKey_Length(void* ptr){
	return static_cast<QSslKey*>(ptr)->length();
}

void QSslKey_Swap(void* ptr, void* other){
	static_cast<QSslKey*>(ptr)->swap(*static_cast<QSslKey*>(other));
}

void* QSslKey_ToDer(void* ptr, void* passPhrase){
	return new QByteArray(static_cast<QSslKey*>(ptr)->toDer(*static_cast<QByteArray*>(passPhrase)));
}

void* QSslKey_ToPem(void* ptr, void* passPhrase){
	return new QByteArray(static_cast<QSslKey*>(ptr)->toPem(*static_cast<QByteArray*>(passPhrase)));
}

void QSslKey_DestroyQSslKey(void* ptr){
	static_cast<QSslKey*>(ptr)->~QSslKey();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator(){
	return new QSslPreSharedKeyAuthenticator();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(void* authenticator){
	return new QSslPreSharedKeyAuthenticator(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void* QSslPreSharedKeyAuthenticator_Identity(void* ptr){
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identity());
}

void* QSslPreSharedKeyAuthenticator_IdentityHint(void* ptr){
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identityHint());
}

int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(void* ptr){
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumIdentityLength();
}

int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(void* ptr){
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumPreSharedKeyLength();
}

void* QSslPreSharedKeyAuthenticator_PreSharedKey(void* ptr){
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->preSharedKey());
}

void QSslPreSharedKeyAuthenticator_SetIdentity(void* ptr, void* identity){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setIdentity(*static_cast<QByteArray*>(identity));
}

void QSslPreSharedKeyAuthenticator_SetPreSharedKey(void* ptr, void* preSharedKey){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setPreSharedKey(*static_cast<QByteArray*>(preSharedKey));
}

void QSslPreSharedKeyAuthenticator_Swap(void* ptr, void* authenticator){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->swap(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(void* ptr){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->~QSslPreSharedKeyAuthenticator();
}

class MyQSslSocket: public QSslSocket {
public:
	MyQSslSocket(QObject *parent) : QSslSocket(parent) {};
	void close() { callbackQSslSocketClose(this, this->objectName().toUtf8().data()); };
	void Signal_Encrypted() { callbackQSslSocketEncrypted(this, this->objectName().toUtf8().data()); };
	void Signal_EncryptedBytesWritten(qint64 written) { callbackQSslSocketEncryptedBytesWritten(this, this->objectName().toUtf8().data(), static_cast<long long>(written)); };
	void Signal_ModeChanged(QSslSocket::SslMode mode) { callbackQSslSocketModeChanged(this, this->objectName().toUtf8().data(), mode); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQSslSocketPreSharedKeyAuthenticationRequired(this, this->objectName().toUtf8().data(), authenticator); };
	void resume() { callbackQSslSocketResume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQSslSocketSetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QAbstractSocket::SocketOption option, const QVariant & value) { callbackQSslSocketSetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	void disconnectFromHost() { callbackQSslSocketDisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSslSocketTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSslSocketChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSslSocketCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSslSocket_NewQSslSocket(void* parent){
	return new MyQSslSocket(static_cast<QObject*>(parent));
}

void QSslSocket_Abort(void* ptr){
	static_cast<QSslSocket*>(ptr)->abort();
}

void QSslSocket_AddCaCertificate(void* ptr, void* certificate){
	static_cast<QSslSocket*>(ptr)->addCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_QSslSocket_AddDefaultCaCertificate(void* certificate){
	QSslSocket::addDefaultCaCertificate(*static_cast<QSslCertificate*>(certificate));
}

int QSslSocket_AtEnd(void* ptr){
	return static_cast<QSslSocket*>(ptr)->atEnd();
}

long long QSslSocket_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->bytesAvailable());
}

long long QSslSocket_BytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->bytesToWrite());
}

int QSslSocket_CanReadLine(void* ptr){
	return static_cast<QSslSocket*>(ptr)->canReadLine();
}

void QSslSocket_Close(void* ptr){
	static_cast<MyQSslSocket*>(ptr)->close();
}

void QSslSocket_CloseDefault(void* ptr){
	static_cast<QSslSocket*>(ptr)->QSslSocket::close();
}

void QSslSocket_ConnectEncrypted(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));;
}

void QSslSocket_DisconnectEncrypted(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)()>(&QSslSocket::encrypted), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)()>(&MyQSslSocket::Signal_Encrypted));;
}

void QSslSocket_Encrypted(void* ptr){
	static_cast<QSslSocket*>(ptr)->encrypted();
}

long long QSslSocket_EncryptedBytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->encryptedBytesAvailable());
}

long long QSslSocket_EncryptedBytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->encryptedBytesToWrite());
}

void QSslSocket_ConnectEncryptedBytesWritten(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(qint64)>(&QSslSocket::encryptedBytesWritten), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(qint64)>(&MyQSslSocket::Signal_EncryptedBytesWritten));;
}

void QSslSocket_DisconnectEncryptedBytesWritten(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(qint64)>(&QSslSocket::encryptedBytesWritten), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(qint64)>(&MyQSslSocket::Signal_EncryptedBytesWritten));;
}

void QSslSocket_EncryptedBytesWritten(void* ptr, long long written){
	static_cast<QSslSocket*>(ptr)->encryptedBytesWritten(static_cast<long long>(written));
}

int QSslSocket_Flush(void* ptr){
	return static_cast<QSslSocket*>(ptr)->flush();
}

void QSslSocket_IgnoreSslErrors(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "ignoreSslErrors");
}

int QSslSocket_IsEncrypted(void* ptr){
	return static_cast<QSslSocket*>(ptr)->isEncrypted();
}

int QSslSocket_Mode(void* ptr){
	return static_cast<QSslSocket*>(ptr)->mode();
}

void QSslSocket_ConnectModeChanged(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));;
}

void QSslSocket_DisconnectModeChanged(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslSocket::SslMode)>(&QSslSocket::modeChanged), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslSocket::SslMode)>(&MyQSslSocket::Signal_ModeChanged));;
}

void QSslSocket_ModeChanged(void* ptr, int mode){
	static_cast<QSslSocket*>(ptr)->modeChanged(static_cast<QSslSocket::SslMode>(mode));
}

int QSslSocket_PeerVerifyDepth(void* ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyDepth();
}

int QSslSocket_PeerVerifyMode(void* ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyMode();
}

char* QSslSocket_PeerVerifyName(void* ptr){
	return static_cast<QSslSocket*>(ptr)->peerVerifyName().toUtf8().data();
}

void QSslSocket_ConnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));;
}

void QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QSslSocket*>(ptr), static_cast<void (QSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QSslSocket::preSharedKeyAuthenticationRequired), static_cast<MyQSslSocket*>(ptr), static_cast<void (MyQSslSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQSslSocket::Signal_PreSharedKeyAuthenticationRequired));;
}

void QSslSocket_PreSharedKeyAuthenticationRequired(void* ptr, void* authenticator){
	static_cast<QSslSocket*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

long long QSslSocket_ReadData(void* ptr, char* data, long long maxlen){
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->readData(data, static_cast<long long>(maxlen)));
}

void QSslSocket_Resume(void* ptr){
	static_cast<MyQSslSocket*>(ptr)->resume();
}

void QSslSocket_ResumeDefault(void* ptr){
	static_cast<QSslSocket*>(ptr)->QSslSocket::resume();
}

void QSslSocket_SetLocalCertificate(void* ptr, void* certificate){
	static_cast<QSslSocket*>(ptr)->setLocalCertificate(*static_cast<QSslCertificate*>(certificate));
}

void QSslSocket_SetPeerVerifyDepth(void* ptr, int depth){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyDepth(depth);
}

void QSslSocket_SetPeerVerifyMode(void* ptr, int mode){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyMode(static_cast<QSslSocket::PeerVerifyMode>(mode));
}

void QSslSocket_SetPeerVerifyName(void* ptr, char* hostName){
	static_cast<QSslSocket*>(ptr)->setPeerVerifyName(QString(hostName));
}

void QSslSocket_SetPrivateKey(void* ptr, void* key){
	static_cast<QSslSocket*>(ptr)->setPrivateKey(*static_cast<QSslKey*>(key));
}

void QSslSocket_SetReadBufferSize(void* ptr, long long size){
	static_cast<MyQSslSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QSslSocket_SetReadBufferSizeDefault(void* ptr, long long size){
	static_cast<QSslSocket*>(ptr)->QSslSocket::setReadBufferSize(static_cast<long long>(size));
}

void QSslSocket_SetSocketOption(void* ptr, int option, void* value){
	static_cast<MyQSslSocket*>(ptr)->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSocketOptionDefault(void* ptr, int option, void* value){
	static_cast<QSslSocket*>(ptr)->QSslSocket::setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QSslSocket_SetSslConfiguration(void* ptr, void* configuration){
	static_cast<QSslSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(configuration));
}

void* QSslSocket_SocketOption(void* ptr, int option){
	return new QVariant(static_cast<QSslSocket*>(ptr)->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

char* QSslSocket_QSslSocket_SslLibraryBuildVersionString(){
	return QSslSocket::sslLibraryBuildVersionString().toUtf8().data();
}

char* QSslSocket_QSslSocket_SslLibraryVersionString(){
	return QSslSocket::sslLibraryVersionString().toUtf8().data();
}

void QSslSocket_StartClientEncryption(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startClientEncryption");
}

void QSslSocket_StartServerEncryption(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSslSocket*>(ptr), "startServerEncryption");
}

int QSslSocket_QSslSocket_SupportsSsl(){
	return QSslSocket::supportsSsl();
}

int QSslSocket_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForBytesWritten(msecs);
}

int QSslSocket_WaitForConnected(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForConnected(msecs);
}

int QSslSocket_WaitForDisconnected(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForDisconnected(msecs);
}

int QSslSocket_WaitForEncrypted(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForEncrypted(msecs);
}

int QSslSocket_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QSslSocket*>(ptr)->waitForReadyRead(msecs);
}

long long QSslSocket_WriteData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QSslSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(len)));
}

void QSslSocket_DestroyQSslSocket(void* ptr){
	static_cast<QSslSocket*>(ptr)->~QSslSocket();
}

void QSslSocket_DisconnectFromHost(void* ptr){
	static_cast<MyQSslSocket*>(ptr)->disconnectFromHost();
}

void QSslSocket_DisconnectFromHostDefault(void* ptr){
	static_cast<QSslSocket*>(ptr)->QSslSocket::disconnectFromHost();
}

void QSslSocket_TimerEvent(void* ptr, void* event){
	static_cast<MyQSslSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSslSocket_TimerEventDefault(void* ptr, void* event){
	static_cast<QSslSocket*>(ptr)->QSslSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSslSocket_ChildEvent(void* ptr, void* event){
	static_cast<MyQSslSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSslSocket_ChildEventDefault(void* ptr, void* event){
	static_cast<QSslSocket*>(ptr)->QSslSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QSslSocket_CustomEvent(void* ptr, void* event){
	static_cast<MyQSslSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSslSocket_CustomEventDefault(void* ptr, void* event){
	static_cast<QSslSocket*>(ptr)->QSslSocket::customEvent(static_cast<QEvent*>(event));
}

class MyQTcpServer: public QTcpServer {
public:
	MyQTcpServer(QObject *parent) : QTcpServer(parent) {};
	void Signal_AcceptError(QAbstractSocket::SocketError socketError) { callbackQTcpServerAcceptError(this, this->objectName().toUtf8().data(), socketError); };
	void Signal_NewConnection() { callbackQTcpServerNewConnection(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQTcpServerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTcpServerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTcpServerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QTcpServer_NewQTcpServer(void* parent){
	return new MyQTcpServer(static_cast<QObject*>(parent));
}

void QTcpServer_ConnectAcceptError(void* ptr){
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));;
}

void QTcpServer_DisconnectAcceptError(void* ptr){
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)(QAbstractSocket::SocketError)>(&QTcpServer::acceptError), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)(QAbstractSocket::SocketError)>(&MyQTcpServer::Signal_AcceptError));;
}

void QTcpServer_AcceptError(void* ptr, int socketError){
	static_cast<QTcpServer*>(ptr)->acceptError(static_cast<QAbstractSocket::SocketError>(socketError));
}

void QTcpServer_Close(void* ptr){
	static_cast<QTcpServer*>(ptr)->close();
}

char* QTcpServer_ErrorString(void* ptr){
	return static_cast<QTcpServer*>(ptr)->errorString().toUtf8().data();
}

int QTcpServer_HasPendingConnections(void* ptr){
	return static_cast<QTcpServer*>(ptr)->hasPendingConnections();
}

int QTcpServer_IsListening(void* ptr){
	return static_cast<QTcpServer*>(ptr)->isListening();
}

int QTcpServer_MaxPendingConnections(void* ptr){
	return static_cast<QTcpServer*>(ptr)->maxPendingConnections();
}

void QTcpServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));;
}

void QTcpServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QTcpServer*>(ptr), static_cast<void (QTcpServer::*)()>(&QTcpServer::newConnection), static_cast<MyQTcpServer*>(ptr), static_cast<void (MyQTcpServer::*)()>(&MyQTcpServer::Signal_NewConnection));;
}

void QTcpServer_NewConnection(void* ptr){
	static_cast<QTcpServer*>(ptr)->newConnection();
}

void* QTcpServer_NextPendingConnection(void* ptr){
	return static_cast<QTcpServer*>(ptr)->nextPendingConnection();
}

void QTcpServer_PauseAccepting(void* ptr){
	static_cast<QTcpServer*>(ptr)->pauseAccepting();
}

void QTcpServer_ResumeAccepting(void* ptr){
	static_cast<QTcpServer*>(ptr)->resumeAccepting();
}

int QTcpServer_ServerError(void* ptr){
	return static_cast<QTcpServer*>(ptr)->serverError();
}

void QTcpServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QTcpServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QTcpServer_SetProxy(void* ptr, void* networkProxy){
	static_cast<QTcpServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

int QTcpServer_WaitForNewConnection(void* ptr, int msec, int timedOut){
	return static_cast<QTcpServer*>(ptr)->waitForNewConnection(msec, NULL);
}

void QTcpServer_DestroyQTcpServer(void* ptr){
	static_cast<QTcpServer*>(ptr)->~QTcpServer();
}

void QTcpServer_TimerEvent(void* ptr, void* event){
	static_cast<MyQTcpServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpServer_TimerEventDefault(void* ptr, void* event){
	static_cast<QTcpServer*>(ptr)->QTcpServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpServer_ChildEvent(void* ptr, void* event){
	static_cast<MyQTcpServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTcpServer_ChildEventDefault(void* ptr, void* event){
	static_cast<QTcpServer*>(ptr)->QTcpServer::childEvent(static_cast<QChildEvent*>(event));
}

void QTcpServer_CustomEvent(void* ptr, void* event){
	static_cast<MyQTcpServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTcpServer_CustomEventDefault(void* ptr, void* event){
	static_cast<QTcpServer*>(ptr)->QTcpServer::customEvent(static_cast<QEvent*>(event));
}

class MyQTcpSocket: public QTcpSocket {
public:
	MyQTcpSocket(QObject *parent) : QTcpSocket(parent) {};
	void close() { callbackQTcpSocketClose(this, this->objectName().toUtf8().data()); };
	void disconnectFromHost() { callbackQTcpSocketDisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void resume() { callbackQTcpSocketResume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQTcpSocketSetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QTcpSocket::SocketOption option, const QVariant & value) { callbackQTcpSocketSetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	void timerEvent(QTimerEvent * event) { callbackQTcpSocketTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTcpSocketChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTcpSocketCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QTcpSocket_NewQTcpSocket(void* parent){
	return new MyQTcpSocket(static_cast<QObject*>(parent));
}

void QTcpSocket_DestroyQTcpSocket(void* ptr){
	static_cast<QTcpSocket*>(ptr)->~QTcpSocket();
}

void QTcpSocket_Close(void* ptr){
	static_cast<MyQTcpSocket*>(ptr)->close();
}

void QTcpSocket_CloseDefault(void* ptr){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::close();
}

void QTcpSocket_DisconnectFromHost(void* ptr){
	static_cast<MyQTcpSocket*>(ptr)->disconnectFromHost();
}

void QTcpSocket_DisconnectFromHostDefault(void* ptr){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::disconnectFromHost();
}

void QTcpSocket_Resume(void* ptr){
	static_cast<MyQTcpSocket*>(ptr)->resume();
}

void QTcpSocket_ResumeDefault(void* ptr){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::resume();
}

void QTcpSocket_SetReadBufferSize(void* ptr, long long size){
	static_cast<MyQTcpSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QTcpSocket_SetReadBufferSizeDefault(void* ptr, long long size){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::setReadBufferSize(static_cast<long long>(size));
}

void QTcpSocket_SetSocketOption(void* ptr, int option, void* value){
	static_cast<MyQTcpSocket*>(ptr)->setSocketOption(static_cast<QTcpSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QTcpSocket_SetSocketOptionDefault(void* ptr, int option, void* value){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::setSocketOption(static_cast<QTcpSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QTcpSocket_TimerEvent(void* ptr, void* event){
	static_cast<MyQTcpSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpSocket_TimerEventDefault(void* ptr, void* event){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTcpSocket_ChildEvent(void* ptr, void* event){
	static_cast<MyQTcpSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTcpSocket_ChildEventDefault(void* ptr, void* event){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QTcpSocket_CustomEvent(void* ptr, void* event){
	static_cast<MyQTcpSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTcpSocket_CustomEventDefault(void* ptr, void* event){
	static_cast<QTcpSocket*>(ptr)->QTcpSocket::customEvent(static_cast<QEvent*>(event));
}

class MyQUdpSocket: public QUdpSocket {
public:
	MyQUdpSocket(QObject *parent) : QUdpSocket(parent) {};
	void close() { callbackQUdpSocketClose(this, this->objectName().toUtf8().data()); };
	void disconnectFromHost() { callbackQUdpSocketDisconnectFromHost(this, this->objectName().toUtf8().data()); };
	void resume() { callbackQUdpSocketResume(this, this->objectName().toUtf8().data()); };
	void setReadBufferSize(qint64 size) { callbackQUdpSocketSetReadBufferSize(this, this->objectName().toUtf8().data(), static_cast<long long>(size)); };
	void setSocketOption(QUdpSocket::SocketOption option, const QVariant & value) { callbackQUdpSocketSetSocketOption(this, this->objectName().toUtf8().data(), option, new QVariant(value)); };
	void timerEvent(QTimerEvent * event) { callbackQUdpSocketTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQUdpSocketChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQUdpSocketCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QUdpSocket_NewQUdpSocket(void* parent){
	return new MyQUdpSocket(static_cast<QObject*>(parent));
}

int QUdpSocket_HasPendingDatagrams(void* ptr){
	return static_cast<QUdpSocket*>(ptr)->hasPendingDatagrams();
}

int QUdpSocket_JoinMulticastGroup(void* ptr, void* groupAddress){
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_JoinMulticastGroup2(void* ptr, void* groupAddress, void* iface){
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

int QUdpSocket_LeaveMulticastGroup(void* ptr, void* groupAddress){
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_LeaveMulticastGroup2(void* ptr, void* groupAddress, void* iface){
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

long long QUdpSocket_PendingDatagramSize(void* ptr){
	return static_cast<long long>(static_cast<QUdpSocket*>(ptr)->pendingDatagramSize());
}

void QUdpSocket_SetMulticastInterface(void* ptr, void* iface){
	static_cast<QUdpSocket*>(ptr)->setMulticastInterface(*static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_DestroyQUdpSocket(void* ptr){
	static_cast<QUdpSocket*>(ptr)->~QUdpSocket();
}

void QUdpSocket_Close(void* ptr){
	static_cast<MyQUdpSocket*>(ptr)->close();
}

void QUdpSocket_CloseDefault(void* ptr){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::close();
}

void QUdpSocket_DisconnectFromHost(void* ptr){
	static_cast<MyQUdpSocket*>(ptr)->disconnectFromHost();
}

void QUdpSocket_DisconnectFromHostDefault(void* ptr){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::disconnectFromHost();
}

void QUdpSocket_Resume(void* ptr){
	static_cast<MyQUdpSocket*>(ptr)->resume();
}

void QUdpSocket_ResumeDefault(void* ptr){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::resume();
}

void QUdpSocket_SetReadBufferSize(void* ptr, long long size){
	static_cast<MyQUdpSocket*>(ptr)->setReadBufferSize(static_cast<long long>(size));
}

void QUdpSocket_SetReadBufferSizeDefault(void* ptr, long long size){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::setReadBufferSize(static_cast<long long>(size));
}

void QUdpSocket_SetSocketOption(void* ptr, int option, void* value){
	static_cast<MyQUdpSocket*>(ptr)->setSocketOption(static_cast<QUdpSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QUdpSocket_SetSocketOptionDefault(void* ptr, int option, void* value){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::setSocketOption(static_cast<QUdpSocket::SocketOption>(option), *static_cast<QVariant*>(value));
}

void QUdpSocket_TimerEvent(void* ptr, void* event){
	static_cast<MyQUdpSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QUdpSocket_TimerEventDefault(void* ptr, void* event){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QUdpSocket_ChildEvent(void* ptr, void* event){
	static_cast<MyQUdpSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QUdpSocket_ChildEventDefault(void* ptr, void* event){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QUdpSocket_CustomEvent(void* ptr, void* event){
	static_cast<MyQUdpSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QUdpSocket_CustomEventDefault(void* ptr, void* event){
	static_cast<QUdpSocket*>(ptr)->QUdpSocket::customEvent(static_cast<QEvent*>(event));
}

