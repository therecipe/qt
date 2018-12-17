// +build !minimal

#define protected public
#define private public

#include "websockets.h"
#include "_cgo_export.h"

#include <QAbstractSocket>
#include <QAuthenticator>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHostAddress>
#include <QLayout>
#include <QMaskGenerator>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNetworkProxy>
#include <QNetworkRequest>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSslConfiguration>
#include <QSslError>
#include <QSslPreSharedKeyAuthenticator>
#include <QString>
#include <QTcpSocket>
#include <QTimerEvent>
#include <QUrl>
#include <QWebSocket>
#include <QWebSocketCorsAuthenticator>
#include <QWebSocketServer>
#include <QWidget>
#include <QWindow>

class MyQMaskGenerator: public QMaskGenerator
{
public:
	MyQMaskGenerator(QObject *parent = Q_NULLPTR) : QMaskGenerator(parent) {QMaskGenerator_QMaskGenerator_QRegisterMetaType();};
	bool seed() { return callbackQMaskGenerator_Seed(this) != 0; };
	quint32 nextMask() { return callbackQMaskGenerator_NextMask(this); };
	 ~MyQMaskGenerator() { callbackQMaskGenerator_DestroyQMaskGenerator(this); };
	bool event(QEvent * e) { return callbackQMaskGenerator_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMaskGenerator_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQMaskGenerator_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMaskGenerator_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMaskGenerator_CustomEvent(this, event); };
	void deleteLater() { callbackQMaskGenerator_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQMaskGenerator_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMaskGenerator_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebSockets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQMaskGenerator_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQMaskGenerator_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMaskGenerator_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQMaskGenerator*)

int QMaskGenerator_QMaskGenerator_QRegisterMetaType(){qRegisterMetaType<QMaskGenerator*>(); return qRegisterMetaType<MyQMaskGenerator*>();}

void* QMaskGenerator_NewQMaskGenerator(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMaskGenerator(static_cast<QWindow*>(parent));
	} else {
		return new MyQMaskGenerator(static_cast<QObject*>(parent));
	}
}

char QMaskGenerator_Seed(void* ptr)
{
	return static_cast<QMaskGenerator*>(ptr)->seed();
}

unsigned int QMaskGenerator_NextMask(void* ptr)
{
	return static_cast<QMaskGenerator*>(ptr)->nextMask();
}

void QMaskGenerator_DestroyQMaskGenerator(void* ptr)
{
	static_cast<QMaskGenerator*>(ptr)->~QMaskGenerator();
}

void QMaskGenerator_DestroyQMaskGeneratorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMaskGenerator___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMaskGenerator___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMaskGenerator___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMaskGenerator___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMaskGenerator___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMaskGenerator___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QMaskGenerator___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMaskGenerator___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMaskGenerator___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QMaskGenerator___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMaskGenerator___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMaskGenerator___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QMaskGenerator___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMaskGenerator___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QMaskGenerator___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QMaskGenerator_EventDefault(void* ptr, void* e)
{
		return static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::event(static_cast<QEvent*>(e));
}

char QMaskGenerator_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QMaskGenerator_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::childEvent(static_cast<QChildEvent*>(event));
}

void QMaskGenerator_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMaskGenerator_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::customEvent(static_cast<QEvent*>(event));
}

void QMaskGenerator_DeleteLaterDefault(void* ptr)
{
		static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::deleteLater();
}

void QMaskGenerator_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMaskGenerator_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QMaskGenerator_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QMaskGenerator*>(ptr)->QMaskGenerator::metaObject());
}

class MyQWebSocket: public QWebSocket
{
public:
	MyQWebSocket(const QString &origin = QString(), QWebSocketProtocol::Version version = QWebSocketProtocol::VersionLatest, QObject *parent = Q_NULLPTR) : QWebSocket(origin, version, parent) {QWebSocket_QWebSocket_QRegisterMetaType();};
	void Signal_AboutToClose() { callbackQWebSocket_AboutToClose(this); };
	void Signal_BinaryFrameReceived(const QByteArray & frame, bool isLastFrame) { callbackQWebSocket_BinaryFrameReceived(this, const_cast<QByteArray*>(&frame), isLastFrame); };
	void Signal_BinaryMessageReceived(const QByteArray & message) { callbackQWebSocket_BinaryMessageReceived(this, const_cast<QByteArray*>(&message)); };
	void Signal_BytesWritten(qint64 bytes) { callbackQWebSocket_BytesWritten(this, bytes); };
	void close(QWebSocketProtocol::CloseCode closeCode, const QString & reason) { QByteArray t7947e9 = reason.toUtf8(); QtWebSockets_PackedString reasonPacked = { const_cast<char*>(t7947e9.prepend("WHITESPACE").constData()+10), t7947e9.size()-10 };callbackQWebSocket_Close(this, closeCode, reasonPacked); };
	void Signal_Connected() { callbackQWebSocket_Connected(this); };
	void Signal_Disconnected() { callbackQWebSocket_Disconnected(this); };
	void Signal_Error2(QAbstractSocket::SocketError error) { callbackQWebSocket_Error2(this, error); };
	void ignoreSslErrors() { callbackQWebSocket_IgnoreSslErrors(this); };
	void open(const QNetworkRequest & request) { callbackQWebSocket_Open2(this, const_cast<QNetworkRequest*>(&request)); };
	void open(const QUrl & url) { callbackQWebSocket_Open(this, const_cast<QUrl*>(&url)); };
	void ping(const QByteArray & payload) { callbackQWebSocket_Ping(this, const_cast<QByteArray*>(&payload)); };
	void Signal_Pong(quint64 elapsedTime, const QByteArray & payload) { callbackQWebSocket_Pong(this, elapsedTime, const_cast<QByteArray*>(&payload)); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQWebSocket_PreSharedKeyAuthenticationRequired(this, authenticator); };
	void Signal_ProxyAuthenticationRequired(const QNetworkProxy & proxy, QAuthenticator * authenticator) { callbackQWebSocket_ProxyAuthenticationRequired(this, const_cast<QNetworkProxy*>(&proxy), authenticator); };
	void Signal_ReadChannelFinished() { callbackQWebSocket_ReadChannelFinished(this); };
	void Signal_SslErrors(const QList<QSslError> & errors) { callbackQWebSocket_SslErrors(this, ({ QList<QSslError>* tmpValue = const_cast<QList<QSslError>*>(&errors); QtWebSockets_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_StateChanged(QAbstractSocket::SocketState state) { callbackQWebSocket_StateChanged(this, state); };
	void Signal_TextFrameReceived(const QString & frame, bool isLastFrame) { QByteArray t39d88b = frame.toUtf8(); QtWebSockets_PackedString framePacked = { const_cast<char*>(t39d88b.prepend("WHITESPACE").constData()+10), t39d88b.size()-10 };callbackQWebSocket_TextFrameReceived(this, framePacked, isLastFrame); };
	void Signal_TextMessageReceived(const QString & message) { QByteArray t6f9b9a = message.toUtf8(); QtWebSockets_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackQWebSocket_TextMessageReceived(this, messagePacked); };
	 ~MyQWebSocket() { callbackQWebSocket_DestroyQWebSocket(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebSocket_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebSocket_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQWebSocket_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebSocket_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebSockets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebSocket_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebSocket_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebSocket*)

int QWebSocket_QWebSocket_QRegisterMetaType(){qRegisterMetaType<QWebSocket*>(); return qRegisterMetaType<MyQWebSocket*>();}

struct QtWebSockets_PackedString QWebSocket_QWebSocket_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t2e2041 = QWebSocket::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t2e2041.prepend("WHITESPACE").constData()+10), t2e2041.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocket_QWebSocket_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t41d947 = QWebSocket::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t41d947.prepend("WHITESPACE").constData()+10), t41d947.size()-10 }; });
}

void* QWebSocket_NewQWebSocket(struct QtWebSockets_PackedString origin, long long version, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QWindow*>(parent));
	} else {
		return new MyQWebSocket(QString::fromUtf8(origin.data, origin.len), static_cast<QWebSocketProtocol::Version>(version), static_cast<QObject*>(parent));
	}
}

char QWebSocket_Flush(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->flush();
}

long long QWebSocket_SendBinaryMessage(void* ptr, void* data)
{
	return static_cast<QWebSocket*>(ptr)->sendBinaryMessage(*static_cast<QByteArray*>(data));
}

long long QWebSocket_SendTextMessage(void* ptr, struct QtWebSockets_PackedString message)
{
	return static_cast<QWebSocket*>(ptr)->sendTextMessage(QString::fromUtf8(message.data, message.len));
}

void QWebSocket_Abort(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->abort();
}

void QWebSocket_ConnectAboutToClose(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::aboutToClose), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_AboutToClose));
}

void QWebSocket_DisconnectAboutToClose(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::aboutToClose), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_AboutToClose));
}

void QWebSocket_AboutToClose(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->aboutToClose();
}

void QWebSocket_ConnectBinaryFrameReceived(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &, bool)>(&QWebSocket::binaryFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &, bool)>(&MyQWebSocket::Signal_BinaryFrameReceived));
}

void QWebSocket_DisconnectBinaryFrameReceived(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &, bool)>(&QWebSocket::binaryFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &, bool)>(&MyQWebSocket::Signal_BinaryFrameReceived));
}

void QWebSocket_BinaryFrameReceived(void* ptr, void* frame, char isLastFrame)
{
	static_cast<QWebSocket*>(ptr)->binaryFrameReceived(*static_cast<QByteArray*>(frame), isLastFrame != 0);
}

void QWebSocket_ConnectBinaryMessageReceived(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &)>(&QWebSocket::binaryMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &)>(&MyQWebSocket::Signal_BinaryMessageReceived));
}

void QWebSocket_DisconnectBinaryMessageReceived(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QByteArray &)>(&QWebSocket::binaryMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QByteArray &)>(&MyQWebSocket::Signal_BinaryMessageReceived));
}

void QWebSocket_BinaryMessageReceived(void* ptr, void* message)
{
	static_cast<QWebSocket*>(ptr)->binaryMessageReceived(*static_cast<QByteArray*>(message));
}

void QWebSocket_ConnectBytesWritten(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(qint64)>(&QWebSocket::bytesWritten), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(qint64)>(&MyQWebSocket::Signal_BytesWritten));
}

void QWebSocket_DisconnectBytesWritten(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(qint64)>(&QWebSocket::bytesWritten), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(qint64)>(&MyQWebSocket::Signal_BytesWritten));
}

void QWebSocket_BytesWritten(void* ptr, long long bytes)
{
	static_cast<QWebSocket*>(ptr)->bytesWritten(bytes);
}

void QWebSocket_Close(void* ptr, long long closeCode, struct QtWebSockets_PackedString reason)
{
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "close", Q_ARG(QWebSocketProtocol::CloseCode, static_cast<QWebSocketProtocol::CloseCode>(closeCode)), Q_ARG(const QString, QString::fromUtf8(reason.data, reason.len)));
}

void QWebSocket_CloseDefault(void* ptr, long long closeCode, struct QtWebSockets_PackedString reason)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::close(static_cast<QWebSocketProtocol::CloseCode>(closeCode), QString::fromUtf8(reason.data, reason.len));
}

void QWebSocket_ConnectConnected(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::connected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Connected));
}

void QWebSocket_DisconnectConnected(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::connected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Connected));
}

void QWebSocket_Connected(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->connected();
}

void QWebSocket_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::disconnected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Disconnected));
}

void QWebSocket_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::disconnected), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_Disconnected));
}

void QWebSocket_Disconnected(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->disconnected();
}

void QWebSocket_ConnectError2(void* ptr)
{
	qRegisterMetaType<QAbstractSocket::SocketError>();
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketError)>(&QWebSocket::error), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketError)>(&MyQWebSocket::Signal_Error2));
}

void QWebSocket_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketError)>(&QWebSocket::error), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketError)>(&MyQWebSocket::Signal_Error2));
}

void QWebSocket_Error2(void* ptr, long long error)
{
	static_cast<QWebSocket*>(ptr)->error(static_cast<QAbstractSocket::SocketError>(error));
}

void QWebSocket_IgnoreSslErrors(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "ignoreSslErrors");
}

void QWebSocket_IgnoreSslErrorsDefault(void* ptr)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::ignoreSslErrors();
}

void QWebSocket_IgnoreSslErrors2(void* ptr, void* errors)
{
	static_cast<QWebSocket*>(ptr)->ignoreSslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QWebSocket_Open2(void* ptr, void* request)
{
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "open", Q_ARG(const QNetworkRequest, *static_cast<QNetworkRequest*>(request)));
}

void QWebSocket_Open2Default(void* ptr, void* request)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::open(*static_cast<QNetworkRequest*>(request));
}

void QWebSocket_Open(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "open", Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QWebSocket_OpenDefault(void* ptr, void* url)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::open(*static_cast<QUrl*>(url));
}

void QWebSocket_Ping(void* ptr, void* payload)
{
	QMetaObject::invokeMethod(static_cast<QWebSocket*>(ptr), "ping", Q_ARG(const QByteArray, *static_cast<QByteArray*>(payload)));
}

void QWebSocket_PingDefault(void* ptr, void* payload)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::ping(*static_cast<QByteArray*>(payload));
}

void QWebSocket_ConnectPong(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(quint64, const QByteArray &)>(&QWebSocket::pong), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(quint64, const QByteArray &)>(&MyQWebSocket::Signal_Pong));
}

void QWebSocket_DisconnectPong(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(quint64, const QByteArray &)>(&QWebSocket::pong), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(quint64, const QByteArray &)>(&MyQWebSocket::Signal_Pong));
}

void QWebSocket_Pong(void* ptr, unsigned long long elapsedTime, void* payload)
{
	static_cast<QWebSocket*>(ptr)->pong(elapsedTime, *static_cast<QByteArray*>(payload));
}

void QWebSocket_ConnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QWebSocket::preSharedKeyAuthenticationRequired), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQWebSocket::Signal_PreSharedKeyAuthenticationRequired));
}

void QWebSocket_DisconnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QSslPreSharedKeyAuthenticator *)>(&QWebSocket::preSharedKeyAuthenticationRequired), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QSslPreSharedKeyAuthenticator *)>(&MyQWebSocket::Signal_PreSharedKeyAuthenticationRequired));
}

void QWebSocket_PreSharedKeyAuthenticationRequired(void* ptr, void* authenticator)
{
	static_cast<QWebSocket*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QWebSocket_ConnectProxyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&QWebSocket::proxyAuthenticationRequired), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&MyQWebSocket::Signal_ProxyAuthenticationRequired));
}

void QWebSocket_DisconnectProxyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&QWebSocket::proxyAuthenticationRequired), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QNetworkProxy &, QAuthenticator *)>(&MyQWebSocket::Signal_ProxyAuthenticationRequired));
}

void QWebSocket_ProxyAuthenticationRequired(void* ptr, void* proxy, void* authenticator)
{
	static_cast<QWebSocket*>(ptr)->proxyAuthenticationRequired(*static_cast<QNetworkProxy*>(proxy), static_cast<QAuthenticator*>(authenticator));
}

void QWebSocket_ConnectReadChannelFinished(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::readChannelFinished), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_ReadChannelFinished));
}

void QWebSocket_DisconnectReadChannelFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)()>(&QWebSocket::readChannelFinished), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)()>(&MyQWebSocket::Signal_ReadChannelFinished));
}

void QWebSocket_ReadChannelFinished(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->readChannelFinished();
}

void QWebSocket_Resume(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->resume();
}

void QWebSocket_SetMaskGenerator(void* ptr, void* maskGenerator)
{
	static_cast<QWebSocket*>(ptr)->setMaskGenerator(static_cast<QMaskGenerator*>(maskGenerator));
}

void QWebSocket_SetPauseMode(void* ptr, long long pauseMode)
{
	static_cast<QWebSocket*>(ptr)->setPauseMode(static_cast<QAbstractSocket::PauseMode>(pauseMode));
}

void QWebSocket_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QWebSocket*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QWebSocket_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QWebSocket*>(ptr)->setReadBufferSize(size);
}

void QWebSocket_SetSslConfiguration(void* ptr, void* sslConfiguration)
{
	static_cast<QWebSocket*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(sslConfiguration));
}

void QWebSocket_ConnectSslErrors(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QList<QSslError> &)>(&QWebSocket::sslErrors), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QList<QSslError> &)>(&MyQWebSocket::Signal_SslErrors));
}

void QWebSocket_DisconnectSslErrors(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QList<QSslError> &)>(&QWebSocket::sslErrors), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QList<QSslError> &)>(&MyQWebSocket::Signal_SslErrors));
}

void QWebSocket_SslErrors(void* ptr, void* errors)
{
	static_cast<QWebSocket*>(ptr)->sslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QWebSocket_ConnectStateChanged(void* ptr)
{
	qRegisterMetaType<QAbstractSocket::SocketState>();
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketState)>(&QWebSocket::stateChanged), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketState)>(&MyQWebSocket::Signal_StateChanged));
}

void QWebSocket_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(QAbstractSocket::SocketState)>(&QWebSocket::stateChanged), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(QAbstractSocket::SocketState)>(&MyQWebSocket::Signal_StateChanged));
}

void QWebSocket_StateChanged(void* ptr, long long state)
{
	static_cast<QWebSocket*>(ptr)->stateChanged(static_cast<QAbstractSocket::SocketState>(state));
}

void QWebSocket_ConnectTextFrameReceived(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &, bool)>(&QWebSocket::textFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &, bool)>(&MyQWebSocket::Signal_TextFrameReceived));
}

void QWebSocket_DisconnectTextFrameReceived(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &, bool)>(&QWebSocket::textFrameReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &, bool)>(&MyQWebSocket::Signal_TextFrameReceived));
}

void QWebSocket_TextFrameReceived(void* ptr, struct QtWebSockets_PackedString frame, char isLastFrame)
{
	static_cast<QWebSocket*>(ptr)->textFrameReceived(QString::fromUtf8(frame.data, frame.len), isLastFrame != 0);
}

void QWebSocket_ConnectTextMessageReceived(void* ptr)
{
	QObject::connect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &)>(&QWebSocket::textMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &)>(&MyQWebSocket::Signal_TextMessageReceived));
}

void QWebSocket_DisconnectTextMessageReceived(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocket*>(ptr), static_cast<void (QWebSocket::*)(const QString &)>(&QWebSocket::textMessageReceived), static_cast<MyQWebSocket*>(ptr), static_cast<void (MyQWebSocket::*)(const QString &)>(&MyQWebSocket::Signal_TextMessageReceived));
}

void QWebSocket_TextMessageReceived(void* ptr, struct QtWebSockets_PackedString message)
{
	static_cast<QWebSocket*>(ptr)->textMessageReceived(QString::fromUtf8(message.data, message.len));
}

void QWebSocket_DestroyQWebSocket(void* ptr)
{
	static_cast<QWebSocket*>(ptr)->~QWebSocket();
}

void QWebSocket_DestroyQWebSocketDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QWebSocket_PauseMode(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->pauseMode();
}

long long QWebSocket_Error(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->error();
}

long long QWebSocket_State(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->state();
}

void* QWebSocket_LocalAddress(void* ptr)
{
	return new QHostAddress(static_cast<QWebSocket*>(ptr)->localAddress());
}

void* QWebSocket_PeerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QWebSocket*>(ptr)->peerAddress());
}

void* QWebSocket_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QWebSocket*>(ptr)->proxy());
}

void* QWebSocket_Request(void* ptr)
{
	return new QNetworkRequest(static_cast<QWebSocket*>(ptr)->request());
}

void* QWebSocket_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QWebSocket*>(ptr)->sslConfiguration());
}

struct QtWebSockets_PackedString QWebSocket_CloseReason(void* ptr)
{
	return ({ QByteArray t003333 = static_cast<QWebSocket*>(ptr)->closeReason().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t003333.prepend("WHITESPACE").constData()+10), t003333.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocket_ErrorString(void* ptr)
{
	return ({ QByteArray t67d61f = static_cast<QWebSocket*>(ptr)->errorString().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t67d61f.prepend("WHITESPACE").constData()+10), t67d61f.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocket_Origin(void* ptr)
{
	return ({ QByteArray t74f689 = static_cast<QWebSocket*>(ptr)->origin().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t74f689.prepend("WHITESPACE").constData()+10), t74f689.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocket_PeerName(void* ptr)
{
	return ({ QByteArray tc206d7 = static_cast<QWebSocket*>(ptr)->peerName().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(tc206d7.prepend("WHITESPACE").constData()+10), tc206d7.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocket_ResourceName(void* ptr)
{
	return ({ QByteArray t1e5662 = static_cast<QWebSocket*>(ptr)->resourceName().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t1e5662.prepend("WHITESPACE").constData()+10), t1e5662.size()-10 }; });
}

void* QWebSocket_RequestUrl(void* ptr)
{
	return new QUrl(static_cast<QWebSocket*>(ptr)->requestUrl());
}

long long QWebSocket_CloseCode(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->closeCode();
}

long long QWebSocket_Version(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->version();
}

char QWebSocket_IsValid(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->isValid();
}

void* QWebSocket_MaskGenerator(void* ptr)
{
	return const_cast<QMaskGenerator*>(static_cast<QWebSocket*>(ptr)->maskGenerator());
}

void* QWebSocket_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebSocket*>(ptr)->QWebSocket::metaObject());
}

long long QWebSocket_BytesToWrite(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->bytesToWrite();
}

long long QWebSocket_ReadBufferSize(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->readBufferSize();
}

unsigned short QWebSocket_LocalPort(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->localPort();
}

unsigned short QWebSocket_PeerPort(void* ptr)
{
	return static_cast<QWebSocket*>(ptr)->peerPort();
}

void* QWebSocket___ignoreSslErrors_errors_atList2(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebSocket___ignoreSslErrors_errors_setList2(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QWebSocket___ignoreSslErrors_errors_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QWebSocket___sslErrors_errors_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebSocket___sslErrors_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QWebSocket___sslErrors_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

void* QWebSocket___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebSocket___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebSocket___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebSocket___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocket___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocket___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebSocket___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocket___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocket___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebSocket___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocket___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocket___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebSocket___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocket___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocket___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebSocket_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebSocket*>(ptr)->QWebSocket::event(static_cast<QEvent*>(e));
}

char QWebSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebSocket*>(ptr)->QWebSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebSocket_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QWebSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebSocket_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::customEvent(static_cast<QEvent*>(event));
}

void QWebSocket_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::deleteLater();
}

void QWebSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebSocket_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebSocket*>(ptr)->QWebSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(void* other)
{
	return new QWebSocketCorsAuthenticator(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(struct QtWebSockets_PackedString origin)
{
	return new QWebSocketCorsAuthenticator(QString::fromUtf8(origin.data, origin.len));
}

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(void* other)
{
	return new QWebSocketCorsAuthenticator(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

void QWebSocketCorsAuthenticator_SetAllowed(void* ptr, char allowed)
{
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->setAllowed(allowed != 0);
}

void QWebSocketCorsAuthenticator_Swap(void* ptr, void* other)
{
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->swap(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

void QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(void* ptr)
{
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->~QWebSocketCorsAuthenticator();
}

struct QtWebSockets_PackedString QWebSocketCorsAuthenticator_Origin(void* ptr)
{
	return ({ QByteArray tde65a2 = static_cast<QWebSocketCorsAuthenticator*>(ptr)->origin().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(tde65a2.prepend("WHITESPACE").constData()+10), tde65a2.size()-10 }; });
}

char QWebSocketCorsAuthenticator_Allowed(void* ptr)
{
	return static_cast<QWebSocketCorsAuthenticator*>(ptr)->allowed();
}

class MyQWebSocketServer: public QWebSocketServer
{
public:
	MyQWebSocketServer(const QString &serverName, QWebSocketServer::SslMode secureMode, QObject *parent = Q_NULLPTR) : QWebSocketServer(serverName, secureMode, parent) {QWebSocketServer_QWebSocketServer_QRegisterMetaType();};
	QWebSocket * nextPendingConnection() { return static_cast<QWebSocket*>(callbackQWebSocketServer_NextPendingConnection(this)); };
	void Signal_AcceptError(QAbstractSocket::SocketError socketError) { callbackQWebSocketServer_AcceptError(this, socketError); };
	void Signal_Closed() { callbackQWebSocketServer_Closed(this); };
	void Signal_NewConnection() { callbackQWebSocketServer_NewConnection(this); };
	void Signal_OriginAuthenticationRequired(QWebSocketCorsAuthenticator * authenticator) { callbackQWebSocketServer_OriginAuthenticationRequired(this, authenticator); };
	void Signal_PeerVerifyError(const QSslError & error) { callbackQWebSocketServer_PeerVerifyError(this, const_cast<QSslError*>(&error)); };
	void Signal_PreSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator * authenticator) { callbackQWebSocketServer_PreSharedKeyAuthenticationRequired(this, authenticator); };
	void Signal_ServerError(QWebSocketProtocol::CloseCode closeCode) { callbackQWebSocketServer_ServerError(this, closeCode); };
	void Signal_SslErrors(const QList<QSslError> & errors) { callbackQWebSocketServer_SslErrors(this, ({ QList<QSslError>* tmpValue = const_cast<QList<QSslError>*>(&errors); QtWebSockets_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~MyQWebSocketServer() { callbackQWebSocketServer_DestroyQWebSocketServer(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebSocketServer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQWebSocketServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebSocketServer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebSocketServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebSocketServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebSocketServer_CustomEvent(this, event); };
	void deleteLater() { callbackQWebSocketServer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebSocketServer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebSocketServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebSockets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebSocketServer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebSocketServer_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQWebSocketServer*)

int QWebSocketServer_QWebSocketServer_QRegisterMetaType(){qRegisterMetaType<QWebSocketServer*>(); return qRegisterMetaType<MyQWebSocketServer*>();}

struct QtWebSockets_PackedString QWebSocketServer_QWebSocketServer_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tda0dc0 = QWebSocketServer::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebSockets_PackedString { const_cast<char*>(tda0dc0.prepend("WHITESPACE").constData()+10), tda0dc0.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocketServer_QWebSocketServer_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t9813c0 = QWebSocketServer::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t9813c0.prepend("WHITESPACE").constData()+10), t9813c0.size()-10 }; });
}

void* QWebSocketServer_NextPendingConnection(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->nextPendingConnection();
}

void* QWebSocketServer_NextPendingConnectionDefault(void* ptr)
{
		return static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::nextPendingConnection();
}

void* QWebSocketServer_NewQWebSocketServer(struct QtWebSockets_PackedString serverName, long long secureMode, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QWindow*>(parent));
	} else {
		return new MyQWebSocketServer(QString::fromUtf8(serverName.data, serverName.len), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QObject*>(parent));
	}
}

char QWebSocketServer_Listen(void* ptr, void* address, unsigned short port)
{
	return static_cast<QWebSocketServer*>(ptr)->listen(*static_cast<QHostAddress*>(address), port);
}

void QWebSocketServer_ConnectAcceptError(void* ptr)
{
	qRegisterMetaType<QAbstractSocket::SocketError>();
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QAbstractSocket::SocketError)>(&QWebSocketServer::acceptError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QAbstractSocket::SocketError)>(&MyQWebSocketServer::Signal_AcceptError));
}

void QWebSocketServer_DisconnectAcceptError(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QAbstractSocket::SocketError)>(&QWebSocketServer::acceptError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QAbstractSocket::SocketError)>(&MyQWebSocketServer::Signal_AcceptError));
}

void QWebSocketServer_AcceptError(void* ptr, long long socketError)
{
	static_cast<QWebSocketServer*>(ptr)->acceptError(static_cast<QAbstractSocket::SocketError>(socketError));
}

void QWebSocketServer_Close(void* ptr)
{
	static_cast<QWebSocketServer*>(ptr)->close();
}

void QWebSocketServer_ConnectClosed(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::closed), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_Closed));
}

void QWebSocketServer_DisconnectClosed(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::closed), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_Closed));
}

void QWebSocketServer_Closed(void* ptr)
{
	static_cast<QWebSocketServer*>(ptr)->closed();
}

void QWebSocketServer_ConnectNewConnection(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::newConnection), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_NewConnection));
}

void QWebSocketServer_DisconnectNewConnection(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::newConnection), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_NewConnection));
}

void QWebSocketServer_NewConnection(void* ptr)
{
	static_cast<QWebSocketServer*>(ptr)->newConnection();
}

void QWebSocketServer_ConnectOriginAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&QWebSocketServer::originAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&MyQWebSocketServer::Signal_OriginAuthenticationRequired));
}

void QWebSocketServer_DisconnectOriginAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&QWebSocketServer::originAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&MyQWebSocketServer::Signal_OriginAuthenticationRequired));
}

void QWebSocketServer_OriginAuthenticationRequired(void* ptr, void* authenticator)
{
	static_cast<QWebSocketServer*>(ptr)->originAuthenticationRequired(static_cast<QWebSocketCorsAuthenticator*>(authenticator));
}

void QWebSocketServer_PauseAccepting(void* ptr)
{
	static_cast<QWebSocketServer*>(ptr)->pauseAccepting();
}

void QWebSocketServer_ConnectPeerVerifyError(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(const QSslError &)>(&QWebSocketServer::peerVerifyError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(const QSslError &)>(&MyQWebSocketServer::Signal_PeerVerifyError));
}

void QWebSocketServer_DisconnectPeerVerifyError(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(const QSslError &)>(&QWebSocketServer::peerVerifyError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(const QSslError &)>(&MyQWebSocketServer::Signal_PeerVerifyError));
}

void QWebSocketServer_PeerVerifyError(void* ptr, void* error)
{
	static_cast<QWebSocketServer*>(ptr)->peerVerifyError(*static_cast<QSslError*>(error));
}

void QWebSocketServer_ConnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QSslPreSharedKeyAuthenticator *)>(&QWebSocketServer::preSharedKeyAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QSslPreSharedKeyAuthenticator *)>(&MyQWebSocketServer::Signal_PreSharedKeyAuthenticationRequired));
}

void QWebSocketServer_DisconnectPreSharedKeyAuthenticationRequired(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QSslPreSharedKeyAuthenticator *)>(&QWebSocketServer::preSharedKeyAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QSslPreSharedKeyAuthenticator *)>(&MyQWebSocketServer::Signal_PreSharedKeyAuthenticationRequired));
}

void QWebSocketServer_PreSharedKeyAuthenticationRequired(void* ptr, void* authenticator)
{
	static_cast<QWebSocketServer*>(ptr)->preSharedKeyAuthenticationRequired(static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QWebSocketServer_ResumeAccepting(void* ptr)
{
	static_cast<QWebSocketServer*>(ptr)->resumeAccepting();
}

void QWebSocketServer_ConnectServerError(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketProtocol::CloseCode)>(&QWebSocketServer::serverError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketProtocol::CloseCode)>(&MyQWebSocketServer::Signal_ServerError));
}

void QWebSocketServer_DisconnectServerError(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketProtocol::CloseCode)>(&QWebSocketServer::serverError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketProtocol::CloseCode)>(&MyQWebSocketServer::Signal_ServerError));
}

void QWebSocketServer_ServerError(void* ptr, long long closeCode)
{
	static_cast<QWebSocketServer*>(ptr)->serverError(static_cast<QWebSocketProtocol::CloseCode>(closeCode));
}

void QWebSocketServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QWebSocketServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QWebSocketServer_SetProxy(void* ptr, void* networkProxy)
{
	static_cast<QWebSocketServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QWebSocketServer_SetServerName(void* ptr, struct QtWebSockets_PackedString serverName)
{
	static_cast<QWebSocketServer*>(ptr)->setServerName(QString::fromUtf8(serverName.data, serverName.len));
}

void QWebSocketServer_SetSslConfiguration(void* ptr, void* sslConfiguration)
{
	static_cast<QWebSocketServer*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(sslConfiguration));
}

void QWebSocketServer_ConnectSslErrors(void* ptr)
{
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(const QList<QSslError> &)>(&QWebSocketServer::sslErrors), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(const QList<QSslError> &)>(&MyQWebSocketServer::Signal_SslErrors));
}

void QWebSocketServer_DisconnectSslErrors(void* ptr)
{
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(const QList<QSslError> &)>(&QWebSocketServer::sslErrors), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(const QList<QSslError> &)>(&MyQWebSocketServer::Signal_SslErrors));
}

void QWebSocketServer_SslErrors(void* ptr, void* errors)
{
	static_cast<QWebSocketServer*>(ptr)->sslErrors(*static_cast<QList<QSslError>*>(errors));
}

void QWebSocketServer_DestroyQWebSocketServer(void* ptr)
{
	static_cast<QWebSocketServer*>(ptr)->~QWebSocketServer();
}

void QWebSocketServer_DestroyQWebSocketServerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebSocketServer_ServerAddress(void* ptr)
{
	return new QHostAddress(static_cast<QWebSocketServer*>(ptr)->serverAddress());
}

struct QtWebSockets_PackedList QWebSocketServer_SupportedVersions(void* ptr)
{
	return ({ QList<QWebSocketProtocol::Version>* tmpValue = new QList<QWebSocketProtocol::Version>(static_cast<QWebSocketServer*>(ptr)->supportedVersions()); QtWebSockets_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebSocketServer_Proxy(void* ptr)
{
	return new QNetworkProxy(static_cast<QWebSocketServer*>(ptr)->proxy());
}

void* QWebSocketServer_SslConfiguration(void* ptr)
{
	return new QSslConfiguration(static_cast<QWebSocketServer*>(ptr)->sslConfiguration());
}

struct QtWebSockets_PackedString QWebSocketServer_ErrorString(void* ptr)
{
	return ({ QByteArray t92e239 = static_cast<QWebSocketServer*>(ptr)->errorString().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(t92e239.prepend("WHITESPACE").constData()+10), t92e239.size()-10 }; });
}

struct QtWebSockets_PackedString QWebSocketServer_ServerName(void* ptr)
{
	return ({ QByteArray tf9a582 = static_cast<QWebSocketServer*>(ptr)->serverName().toUtf8(); QtWebSockets_PackedString { const_cast<char*>(tf9a582.prepend("WHITESPACE").constData()+10), tf9a582.size()-10 }; });
}

void* QWebSocketServer_ServerUrl(void* ptr)
{
	return new QUrl(static_cast<QWebSocketServer*>(ptr)->serverUrl());
}

long long QWebSocketServer_Error(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->error();
}

long long QWebSocketServer_SecureMode(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->secureMode();
}

char QWebSocketServer_HasPendingConnections(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->hasPendingConnections();
}

char QWebSocketServer_IsListening(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->isListening();
}

void* QWebSocketServer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::metaObject());
}

int QWebSocketServer_MaxPendingConnections(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->maxPendingConnections();
}

unsigned short QWebSocketServer_ServerPort(void* ptr)
{
	return static_cast<QWebSocketServer*>(ptr)->serverPort();
}

void QWebSocketServer_HandleConnection(void* ptr, void* socket)
{
	static_cast<QWebSocketServer*>(ptr)->handleConnection(static_cast<QTcpSocket*>(socket));
}

void* QWebSocketServer___sslErrors_errors_atList(void* ptr, int i)
{
	return new QSslError(({QSslError tmp = static_cast<QList<QSslError>*>(ptr)->at(i); if (i == static_cast<QList<QSslError>*>(ptr)->size()-1) { static_cast<QList<QSslError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebSocketServer___sslErrors_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QSslError>*>(ptr)->append(*static_cast<QSslError*>(i));
}

void* QWebSocketServer___sslErrors_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSslError>();
}

long long QWebSocketServer___supportedVersions_atList(void* ptr, int i)
{
	return ({QWebSocketProtocol::Version tmp = static_cast<QList<QWebSocketProtocol::Version>*>(ptr)->at(i); if (i == static_cast<QList<QWebSocketProtocol::Version>*>(ptr)->size()-1) { static_cast<QList<QWebSocketProtocol::Version>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocketServer___supportedVersions_setList(void* ptr, long long i)
{
	static_cast<QList<QWebSocketProtocol::Version>*>(ptr)->append(static_cast<QWebSocketProtocol::Version>(i));
}

void* QWebSocketServer___supportedVersions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebSocketProtocol::Version>();
}

void* QWebSocketServer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWebSocketServer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebSocketServer___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWebSocketServer___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocketServer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocketServer___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebSocketServer___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocketServer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocketServer___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebSocketServer___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocketServer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocketServer___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWebSocketServer___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWebSocketServer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebSocketServer___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QWebSocketServer_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::event(static_cast<QEvent*>(e));
}

char QWebSocketServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebSocketServer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::childEvent(static_cast<QChildEvent*>(event));
}

void QWebSocketServer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebSocketServer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::customEvent(static_cast<QEvent*>(event));
}

void QWebSocketServer_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::deleteLater();
}

void QWebSocketServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebSocketServer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebSocketServer*>(ptr)->QWebSocketServer::timerEvent(static_cast<QTimerEvent*>(event));
}

