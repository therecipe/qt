// +build !minimal

#define protected public
#define private public

#include "serialbus.h"
#include "_cgo_export.h"

#include <QAudioSystemPlugin>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QCanBus>
#include <QCanBusDevice>
#include <QCanBusDeviceInfo>
#include <QCanBusFactory>
#include <QCanBusFactoryV2>
#include <QCanBusFrame>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QModbusClient>
#include <QModbusDataUnit>
#include <QModbusDevice>
#include <QModbusDeviceIdentification>
#include <QModbusExceptionResponse>
#include <QModbusPdu>
#include <QModbusReply>
#include <QModbusRequest>
#include <QModbusResponse>
#include <QModbusRtuSerialMaster>
#include <QModbusRtuSerialSlave>
#include <QModbusServer>
#include <QModbusTcpClient>
#include <QModbusTcpConnectionObserver>
#include <QModbusTcpServer>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QString>
#include <QTcpSocket>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQCanBus: public QCanBus
{
public:
	void childEvent(QChildEvent * event) { callbackQCanBus_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCanBus_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCanBus_CustomEvent(this, event); };
	void deleteLater() { callbackQCanBus_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCanBus_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCanBus_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCanBus_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCanBus_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCanBus_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQCanBus_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCanBus_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QCanBus*)
Q_DECLARE_METATYPE(MyQCanBus*)

int QCanBus_QCanBus_QRegisterMetaType(){qRegisterMetaType<QCanBus*>(); return qRegisterMetaType<MyQCanBus*>();}

struct QtSerialBus_PackedList QCanBus_AvailableDevices(void* ptr, struct QtSerialBus_PackedString plugin, struct QtSerialBus_PackedString errorMessage)
{
	return ({ QList<QCanBusDeviceInfo>* tmpValue899674 = new QList<QCanBusDeviceInfo>(static_cast<QCanBus*>(ptr)->availableDevices(QString::fromUtf8(plugin.data, plugin.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)))); QtSerialBus_PackedList { tmpValue899674, tmpValue899674->size() }; });
}

void* QCanBus_CreateDevice(void* ptr, struct QtSerialBus_PackedString plugin, struct QtSerialBus_PackedString interfaceName, struct QtSerialBus_PackedString errorMessage)
{
	return static_cast<QCanBus*>(ptr)->createDevice(QString::fromUtf8(plugin.data, plugin.len), QString::fromUtf8(interfaceName.data, interfaceName.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)));
}

void* QCanBus_QCanBus_Instance()
{
	return QCanBus::instance();
}

struct QtSerialBus_PackedString QCanBus_Plugins(void* ptr)
{
	return ({ QByteArray* t48ad25 = new QByteArray(static_cast<QCanBus*>(ptr)->plugins().join("¡¦!").toUtf8()); QtSerialBus_PackedString { const_cast<char*>(t48ad25->prepend("WHITESPACE").constData()+10), t48ad25->size()-10, t48ad25 }; });
}

void* QCanBus___availableDevices_atList(void* ptr, int i)
{
	return new QCanBusDeviceInfo(({QCanBusDeviceInfo tmp = static_cast<QList<QCanBusDeviceInfo>*>(ptr)->at(i); if (i == static_cast<QList<QCanBusDeviceInfo>*>(ptr)->size()-1) { static_cast<QList<QCanBusDeviceInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QCanBus___availableDevices_setList(void* ptr, void* i)
{
	static_cast<QList<QCanBusDeviceInfo>*>(ptr)->append(*static_cast<QCanBusDeviceInfo*>(i));
}

void* QCanBus___availableDevices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCanBusDeviceInfo>();
}

void* QCanBus___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCanBus___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QCanBus___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QCanBus___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCanBus___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QCanBus___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCanBus___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCanBus___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCanBus___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QCanBus_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QCanBus*>(ptr)->QCanBus::childEvent(static_cast<QChildEvent*>(event));
}

void QCanBus_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QCanBus*>(ptr)->QCanBus::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBus_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QCanBus*>(ptr)->QCanBus::customEvent(static_cast<QEvent*>(event));
}

void QCanBus_DeleteLaterDefault(void* ptr)
{
		static_cast<QCanBus*>(ptr)->QCanBus::deleteLater();
}

void QCanBus_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QCanBus*>(ptr)->QCanBus::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCanBus_EventDefault(void* ptr, void* e)
{
		return static_cast<QCanBus*>(ptr)->QCanBus::event(static_cast<QEvent*>(e));
}

char QCanBus_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QCanBus*>(ptr)->QCanBus::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCanBus_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QCanBus*>(ptr)->QCanBus::metaObject());
}

void QCanBus_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QCanBus*>(ptr)->QCanBus::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQCanBusDevice: public QCanBusDevice
{
public:
	MyQCanBusDevice(QObject *parent = Q_NULLPTR) : QCanBusDevice(parent) {QCanBusDevice_QCanBusDevice_QRegisterMetaType();};
	void close() { callbackQCanBusDevice_Close(this); };
	void Signal_ErrorOccurred(QCanBusDevice::CanBusError vqc) { callbackQCanBusDevice_ErrorOccurred(this, vqc); };
	void Signal_FramesReceived() { callbackQCanBusDevice_FramesReceived(this); };
	void Signal_FramesWritten(qint64 framesCount) { callbackQCanBusDevice_FramesWritten(this, framesCount); };
	QString interpretErrorFrame(const QCanBusFrame & frame) { return ({ QtSerialBus_PackedString tempVal = callbackQCanBusDevice_InterpretErrorFrame(this, const_cast<QCanBusFrame*>(&frame)); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool open() { return callbackQCanBusDevice_Open(this) != 0; };
	void setConfigurationParameter(int key, const QVariant & value) { callbackQCanBusDevice_SetConfigurationParameter(this, key, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QCanBusDevice::CanBusDeviceState state) { callbackQCanBusDevice_StateChanged(this, state); };
	bool waitForFramesReceived(int msecs) { return callbackQCanBusDevice_WaitForFramesReceived(this, msecs) != 0; };
	bool waitForFramesWritten(int msecs) { return callbackQCanBusDevice_WaitForFramesWritten(this, msecs) != 0; };
	bool writeFrame(const QCanBusFrame & frame) { return callbackQCanBusDevice_WriteFrame(this, const_cast<QCanBusFrame*>(&frame)) != 0; };
	void childEvent(QChildEvent * event) { callbackQCanBusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCanBusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQCanBusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCanBusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCanBusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCanBusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCanBusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQCanBusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCanBusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QCanBusDevice*)
Q_DECLARE_METATYPE(MyQCanBusDevice*)

int QCanBusDevice_QCanBusDevice_QRegisterMetaType(){qRegisterMetaType<QCanBusDevice*>(); return qRegisterMetaType<MyQCanBusDevice*>();}

void* QCanBusDevice_NewQCanBusDevice2(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QWindow*>(parent));
	} else {
		return new MyQCanBusDevice(static_cast<QObject*>(parent));
	}
}

void QCanBusDevice_Clear(void* ptr, long long direction)
{
	static_cast<QCanBusDevice*>(ptr)->clear(static_cast<QCanBusDevice::Direction>(direction));
}

void QCanBusDevice_Close(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->close();
}

struct QtSerialBus_PackedList QCanBusDevice_ConfigurationKeys(void* ptr)
{
	return ({ QVector<int>* tmpValue49545d = new QVector<int>(static_cast<QCanBusDevice*>(ptr)->configurationKeys()); QtSerialBus_PackedList { tmpValue49545d, tmpValue49545d->size() }; });
}

void* QCanBusDevice_ConfigurationParameter(void* ptr, int key)
{
	return new QVariant(static_cast<QCanBusDevice*>(ptr)->configurationParameter(key));
}

char QCanBusDevice_ConnectDevice(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->connectDevice();
}

void QCanBusDevice_DisconnectDevice(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->disconnectDevice();
}

void QCanBusDevice_EnqueueOutgoingFrame(void* ptr, void* newFrame)
{
	static_cast<QCanBusDevice*>(ptr)->enqueueOutgoingFrame(*static_cast<QCanBusFrame*>(newFrame));
}

void QCanBusDevice_EnqueueReceivedFrames(void* ptr, void* newFrames)
{
	static_cast<QCanBusDevice*>(ptr)->enqueueReceivedFrames(*static_cast<QVector<QCanBusFrame>*>(newFrames));
}

long long QCanBusDevice_Error(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->error();
}

void QCanBusDevice_ConnectErrorOccurred(void* ptr, long long t)
{
	qRegisterMetaType<QCanBusDevice::CanBusError>();
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusError)>(&QCanBusDevice::errorOccurred), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusError)>(&MyQCanBusDevice::Signal_ErrorOccurred), static_cast<Qt::ConnectionType>(t));
}

void QCanBusDevice_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusError)>(&QCanBusDevice::errorOccurred), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusError)>(&MyQCanBusDevice::Signal_ErrorOccurred));
}

void QCanBusDevice_ErrorOccurred(void* ptr, long long vqc)
{
	static_cast<QCanBusDevice*>(ptr)->errorOccurred(static_cast<QCanBusDevice::CanBusError>(vqc));
}

struct QtSerialBus_PackedString QCanBusDevice_ErrorString(void* ptr)
{
	return ({ QByteArray* t646153 = new QByteArray(static_cast<QCanBusDevice*>(ptr)->errorString().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(t646153->prepend("WHITESPACE").constData()+10), t646153->size()-10, t646153 }; });
}

long long QCanBusDevice_FramesAvailable(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->framesAvailable();
}

void QCanBusDevice_ConnectFramesReceived(void* ptr, long long t)
{
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)()>(&QCanBusDevice::framesReceived), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)()>(&MyQCanBusDevice::Signal_FramesReceived), static_cast<Qt::ConnectionType>(t));
}

void QCanBusDevice_DisconnectFramesReceived(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)()>(&QCanBusDevice::framesReceived), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)()>(&MyQCanBusDevice::Signal_FramesReceived));
}

void QCanBusDevice_FramesReceived(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->framesReceived();
}

long long QCanBusDevice_FramesToWrite(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->framesToWrite();
}

void QCanBusDevice_ConnectFramesWritten(void* ptr, long long t)
{
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(qint64)>(&QCanBusDevice::framesWritten), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(qint64)>(&MyQCanBusDevice::Signal_FramesWritten), static_cast<Qt::ConnectionType>(t));
}

void QCanBusDevice_DisconnectFramesWritten(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(qint64)>(&QCanBusDevice::framesWritten), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(qint64)>(&MyQCanBusDevice::Signal_FramesWritten));
}

void QCanBusDevice_FramesWritten(void* ptr, long long framesCount)
{
	static_cast<QCanBusDevice*>(ptr)->framesWritten(framesCount);
}

char QCanBusDevice_HasOutgoingFrames(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->hasOutgoingFrames();
}

struct QtSerialBus_PackedString QCanBusDevice_InterpretErrorFrame(void* ptr, void* frame)
{
	return ({ QByteArray* t5dc8ff = new QByteArray(static_cast<QCanBusDevice*>(ptr)->interpretErrorFrame(*static_cast<QCanBusFrame*>(frame)).toUtf8()); QtSerialBus_PackedString { const_cast<char*>(t5dc8ff->prepend("WHITESPACE").constData()+10), t5dc8ff->size()-10, t5dc8ff }; });
}

char QCanBusDevice_Open(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->open();
}

void QCanBusDevice_SetConfigurationParameter(void* ptr, int key, void* value)
{
	static_cast<QCanBusDevice*>(ptr)->setConfigurationParameter(key, *static_cast<QVariant*>(value));
}

void QCanBusDevice_SetConfigurationParameterDefault(void* ptr, int key, void* value)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::setConfigurationParameter(key, *static_cast<QVariant*>(value));
}

void QCanBusDevice_SetError(void* ptr, struct QtSerialBus_PackedString errorText, long long errorId)
{
	static_cast<QCanBusDevice*>(ptr)->setError(QString::fromUtf8(errorText.data, errorText.len), static_cast<QCanBusDevice::CanBusError>(errorId));
}

void QCanBusDevice_SetState(void* ptr, long long newState)
{
	static_cast<QCanBusDevice*>(ptr)->setState(static_cast<QCanBusDevice::CanBusDeviceState>(newState));
}

long long QCanBusDevice_State(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->state();
}

void QCanBusDevice_ConnectStateChanged(void* ptr, long long t)
{
	qRegisterMetaType<QCanBusDevice::CanBusDeviceState>();
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&QCanBusDevice::stateChanged), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&MyQCanBusDevice::Signal_StateChanged), static_cast<Qt::ConnectionType>(t));
}

void QCanBusDevice_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&QCanBusDevice::stateChanged), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&MyQCanBusDevice::Signal_StateChanged));
}

void QCanBusDevice_StateChanged(void* ptr, long long state)
{
	static_cast<QCanBusDevice*>(ptr)->stateChanged(static_cast<QCanBusDevice::CanBusDeviceState>(state));
}

char QCanBusDevice_WaitForFramesReceived(void* ptr, int msecs)
{
	return static_cast<QCanBusDevice*>(ptr)->waitForFramesReceived(msecs);
}

char QCanBusDevice_WaitForFramesReceivedDefault(void* ptr, int msecs)
{
		return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::waitForFramesReceived(msecs);
}

char QCanBusDevice_WaitForFramesWritten(void* ptr, int msecs)
{
	return static_cast<QCanBusDevice*>(ptr)->waitForFramesWritten(msecs);
}

char QCanBusDevice_WaitForFramesWrittenDefault(void* ptr, int msecs)
{
		return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::waitForFramesWritten(msecs);
}

char QCanBusDevice_WriteFrame(void* ptr, void* frame)
{
	return static_cast<QCanBusDevice*>(ptr)->writeFrame(*static_cast<QCanBusFrame*>(frame));
}

int QCanBusDevice___configurationKeys_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QCanBusDevice___configurationKeys_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QCanBusDevice___configurationKeys_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void QCanBusDevice___enqueueReceivedFrames_newFrames_setList(void* ptr, void* i)
{
	static_cast<QVector<QCanBusFrame>*>(ptr)->append(*static_cast<QCanBusFrame*>(i));
}

void* QCanBusDevice___enqueueReceivedFrames_newFrames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QCanBusFrame>();
}

void* QCanBusDevice___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCanBusDevice___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QCanBusDevice___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QCanBusDevice___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCanBusDevice___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QCanBusDevice___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCanBusDevice___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCanBusDevice___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCanBusDevice___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QCanBusDevice_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::childEvent(static_cast<QChildEvent*>(event));
}

void QCanBusDevice_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBusDevice_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::customEvent(static_cast<QEvent*>(event));
}

void QCanBusDevice_DeleteLaterDefault(void* ptr)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::deleteLater();
}

void QCanBusDevice_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCanBusDevice_EventDefault(void* ptr, void* e)
{
		return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::event(static_cast<QEvent*>(e));
}

char QCanBusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCanBusDevice_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::metaObject());
}

void QCanBusDevice_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

int QCanBusDeviceInfo_Channel(void* ptr)
{
	return static_cast<QCanBusDeviceInfo*>(ptr)->channel();
}

struct QtSerialBus_PackedString QCanBusDeviceInfo_Description(void* ptr)
{
	return ({ QByteArray* t574a7a = new QByteArray(static_cast<QCanBusDeviceInfo*>(ptr)->description().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(t574a7a->prepend("WHITESPACE").constData()+10), t574a7a->size()-10, t574a7a }; });
}

char QCanBusDeviceInfo_HasFlexibleDataRate(void* ptr)
{
	return static_cast<QCanBusDeviceInfo*>(ptr)->hasFlexibleDataRate();
}

char QCanBusDeviceInfo_IsVirtual(void* ptr)
{
	return static_cast<QCanBusDeviceInfo*>(ptr)->isVirtual();
}

struct QtSerialBus_PackedString QCanBusDeviceInfo_Name(void* ptr)
{
	return ({ QByteArray* tdc7ae5 = new QByteArray(static_cast<QCanBusDeviceInfo*>(ptr)->name().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(tdc7ae5->prepend("WHITESPACE").constData()+10), tdc7ae5->size()-10, tdc7ae5 }; });
}

struct QtSerialBus_PackedString QCanBusDeviceInfo_SerialNumber(void* ptr)
{
	return ({ QByteArray* t8317b9 = new QByteArray(static_cast<QCanBusDeviceInfo*>(ptr)->serialNumber().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(t8317b9->prepend("WHITESPACE").constData()+10), t8317b9->size()-10, t8317b9 }; });
}

void QCanBusDeviceInfo_Swap(void* ptr, void* other)
{
	static_cast<QCanBusDeviceInfo*>(ptr)->swap(*static_cast<QCanBusDeviceInfo*>(other));
}

void QCanBusDeviceInfo_DestroyQCanBusDeviceInfo(void* ptr)
{
	static_cast<QCanBusDeviceInfo*>(ptr)->~QCanBusDeviceInfo();
}

class MyQCanBusFactory: public QCanBusFactory
{
public:
	QCanBusDevice * createDevice(const QString & interfaceName, QString * errorMessage) const { QByteArray* tf83f00 = new QByteArray(interfaceName.toUtf8()); QtSerialBus_PackedString interfaceNamePacked = { const_cast<char*>(tf83f00->prepend("WHITESPACE").constData()+10), tf83f00->size()-10, tf83f00 };QByteArray* t3f2abc = new QByteArray(errorMessage->toUtf8()); QtSerialBus_PackedString errorMessagePacked = { const_cast<char*>(t3f2abc->prepend("WHITESPACE").constData()+10), t3f2abc->size()-10, t3f2abc };return static_cast<QCanBusDevice*>(callbackQCanBusFactory_CreateDevice(const_cast<void*>(static_cast<const void*>(this)), interfaceNamePacked, errorMessagePacked)); };
};

Q_DECLARE_METATYPE(QCanBusFactory*)
Q_DECLARE_METATYPE(MyQCanBusFactory*)

int QCanBusFactory_QCanBusFactory_QRegisterMetaType(){qRegisterMetaType<QCanBusFactory*>(); return qRegisterMetaType<MyQCanBusFactory*>();}

void* QCanBusFactory_CreateDevice(void* ptr, struct QtSerialBus_PackedString interfaceName, struct QtSerialBus_PackedString errorMessage)
{
	return static_cast<QCanBusFactory*>(ptr)->createDevice(QString::fromUtf8(interfaceName.data, interfaceName.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)));
}

class MyQCanBusFactoryV2: public QCanBusFactoryV2
{
public:
	QList<QCanBusDeviceInfo> availableDevices(QString * errorMessage) const { QByteArray* t3f2abc = new QByteArray(errorMessage->toUtf8()); QtSerialBus_PackedString errorMessagePacked = { const_cast<char*>(t3f2abc->prepend("WHITESPACE").constData()+10), t3f2abc->size()-10, t3f2abc };return ({ QList<QCanBusDeviceInfo>* tmpP = static_cast<QList<QCanBusDeviceInfo>*>(callbackQCanBusFactoryV2_AvailableDevices(const_cast<void*>(static_cast<const void*>(this)), errorMessagePacked)); QList<QCanBusDeviceInfo> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QCanBusDevice * createDevice(const QString & interfaceName, QString * errorMessage) const { QByteArray* tf83f00 = new QByteArray(interfaceName.toUtf8()); QtSerialBus_PackedString interfaceNamePacked = { const_cast<char*>(tf83f00->prepend("WHITESPACE").constData()+10), tf83f00->size()-10, tf83f00 };QByteArray* t3f2abc = new QByteArray(errorMessage->toUtf8()); QtSerialBus_PackedString errorMessagePacked = { const_cast<char*>(t3f2abc->prepend("WHITESPACE").constData()+10), t3f2abc->size()-10, t3f2abc };return static_cast<QCanBusDevice*>(callbackQCanBusFactoryV2_CreateDevice(const_cast<void*>(static_cast<const void*>(this)), interfaceNamePacked, errorMessagePacked)); };
};

Q_DECLARE_METATYPE(QCanBusFactoryV2*)
Q_DECLARE_METATYPE(MyQCanBusFactoryV2*)

int QCanBusFactoryV2_QCanBusFactoryV2_QRegisterMetaType(){qRegisterMetaType<QCanBusFactoryV2*>(); return qRegisterMetaType<MyQCanBusFactoryV2*>();}

struct QtSerialBus_PackedList QCanBusFactoryV2_AvailableDevices(void* ptr, struct QtSerialBus_PackedString errorMessage)
{
	return ({ QList<QCanBusDeviceInfo>* tmpValue4723b9 = new QList<QCanBusDeviceInfo>(static_cast<QCanBusFactoryV2*>(ptr)->availableDevices(new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)))); QtSerialBus_PackedList { tmpValue4723b9, tmpValue4723b9->size() }; });
}

void* QCanBusFactoryV2_CreateDevice(void* ptr, struct QtSerialBus_PackedString interfaceName, struct QtSerialBus_PackedString errorMessage)
{
	return static_cast<QCanBusFactoryV2*>(ptr)->createDevice(QString::fromUtf8(interfaceName.data, interfaceName.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)));
}

void* QCanBusFactoryV2___availableDevices_atList(void* ptr, int i)
{
	return new QCanBusDeviceInfo(({QCanBusDeviceInfo tmp = static_cast<QList<QCanBusDeviceInfo>*>(ptr)->at(i); if (i == static_cast<QList<QCanBusDeviceInfo>*>(ptr)->size()-1) { static_cast<QList<QCanBusDeviceInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QCanBusFactoryV2___availableDevices_setList(void* ptr, void* i)
{
	static_cast<QList<QCanBusDeviceInfo>*>(ptr)->append(*static_cast<QCanBusDeviceInfo*>(i));
}

void* QCanBusFactoryV2___availableDevices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCanBusDeviceInfo>();
}

Q_DECLARE_METATYPE(QCanBusFrame*)
int QCanBusFrame_TransmissionTimeoutError_Type()
{
	return QCanBusFrame::TransmissionTimeoutError;
}

int QCanBusFrame_LostArbitrationError_Type()
{
	return QCanBusFrame::LostArbitrationError;
}

int QCanBusFrame_ControllerError_Type()
{
	return QCanBusFrame::ControllerError;
}

int QCanBusFrame_ProtocolViolationError_Type()
{
	return QCanBusFrame::ProtocolViolationError;
}

int QCanBusFrame_TransceiverError_Type()
{
	return QCanBusFrame::TransceiverError;
}

int QCanBusFrame_MissingAcknowledgmentError_Type()
{
	return QCanBusFrame::MissingAcknowledgmentError;
}

int QCanBusFrame_BusOffError_Type()
{
	return QCanBusFrame::BusOffError;
}

int QCanBusFrame_BusError_Type()
{
	return QCanBusFrame::BusError;
}

int QCanBusFrame_ControllerRestartError_Type()
{
	return QCanBusFrame::ControllerRestartError;
}

int QCanBusFrame_UnknownError_Type()
{
	return QCanBusFrame::UnknownError;
}

int QCanBusFrame_AnyError_Type()
{
	return QCanBusFrame::AnyError;
}

void* QCanBusFrame_NewQCanBusFrame(long long ty)
{
	return new QCanBusFrame(static_cast<QCanBusFrame::FrameType>(ty));
}

void* QCanBusFrame_NewQCanBusFrame2(unsigned int identifier, void* data)
{
	return new QCanBusFrame(identifier, *static_cast<QByteArray*>(data));
}

long long QCanBusFrame_Error(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->error();
}

unsigned int QCanBusFrame_FrameId(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->frameId();
}

long long QCanBusFrame_FrameType(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->frameType();
}

char QCanBusFrame_HasBitrateSwitch(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasBitrateSwitch();
}

char QCanBusFrame_HasErrorStateIndicator(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasErrorStateIndicator();
}

char QCanBusFrame_HasExtendedFrameFormat(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasExtendedFrameFormat();
}

char QCanBusFrame_HasFlexibleDataRateFormat(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasFlexibleDataRateFormat();
}

char QCanBusFrame_HasLocalEcho(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasLocalEcho();
}

char QCanBusFrame_IsValid(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->isValid();
}

void* QCanBusFrame_Payload(void* ptr)
{
	return new QByteArray(static_cast<QCanBusFrame*>(ptr)->payload());
}

void QCanBusFrame_SetBitrateSwitch(void* ptr, char bitrateSwitch)
{
	static_cast<QCanBusFrame*>(ptr)->setBitrateSwitch(bitrateSwitch != 0);
}

void QCanBusFrame_SetError(void* ptr, long long error)
{
	static_cast<QCanBusFrame*>(ptr)->setError(static_cast<QCanBusFrame::FrameError>(error));
}

void QCanBusFrame_SetErrorStateIndicator(void* ptr, char errorStateIndicator)
{
	static_cast<QCanBusFrame*>(ptr)->setErrorStateIndicator(errorStateIndicator != 0);
}

void QCanBusFrame_SetExtendedFrameFormat(void* ptr, char isExtended)
{
	static_cast<QCanBusFrame*>(ptr)->setExtendedFrameFormat(isExtended != 0);
}

void QCanBusFrame_SetFlexibleDataRateFormat(void* ptr, char isFlexibleData)
{
	static_cast<QCanBusFrame*>(ptr)->setFlexibleDataRateFormat(isFlexibleData != 0);
}

void QCanBusFrame_SetFrameId(void* ptr, unsigned int newFrameId)
{
	static_cast<QCanBusFrame*>(ptr)->setFrameId(newFrameId);
}

void QCanBusFrame_SetFrameType(void* ptr, long long newType)
{
	static_cast<QCanBusFrame*>(ptr)->setFrameType(static_cast<QCanBusFrame::FrameType>(newType));
}

void QCanBusFrame_SetLocalEcho(void* ptr, char echo)
{
	static_cast<QCanBusFrame*>(ptr)->setLocalEcho(echo != 0);
}

void QCanBusFrame_SetPayload(void* ptr, void* data)
{
	static_cast<QCanBusFrame*>(ptr)->setPayload(*static_cast<QByteArray*>(data));
}

struct QtSerialBus_PackedString QCanBusFrame_ToString(void* ptr)
{
	return ({ QByteArray* tda41b1 = new QByteArray(static_cast<QCanBusFrame*>(ptr)->toString().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(tda41b1->prepend("WHITESPACE").constData()+10), tda41b1->size()-10, tda41b1 }; });
}

class MyQModbusClient: public QModbusClient
{
public:
	MyQModbusClient(QObject *parent = Q_NULLPTR) : QModbusClient(parent) {QModbusClient_QModbusClient_QRegisterMetaType();};
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	void close() { callbackQModbusClient_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	bool open() { return callbackQModbusClient_Open(this) != 0; };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusClient*)
Q_DECLARE_METATYPE(MyQModbusClient*)

int QModbusClient_QModbusClient_QRegisterMetaType(){qRegisterMetaType<QModbusClient*>(); return qRegisterMetaType<MyQModbusClient*>();}

void* QModbusClient_NewQModbusClient(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusClient(static_cast<QObject*>(parent));
	}
}

int QModbusClient_NumberOfRetries(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->numberOfRetries();
}

char QModbusClient_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusClient_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
	} else {
		return static_cast<QModbusClient*>(ptr)->QModbusClient::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
	}
}

char QModbusClient_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusClient_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
	} else {
		return static_cast<QModbusClient*>(ptr)->QModbusClient::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
	}
}

void* QModbusClient_SendRawRequest(void* ptr, void* request, int serverAddress)
{
	return static_cast<QModbusClient*>(ptr)->sendRawRequest(*static_cast<QModbusRequest*>(request), serverAddress);
}

void* QModbusClient_SendReadRequest(void* ptr, void* read, int serverAddress)
{
	return static_cast<QModbusClient*>(ptr)->sendReadRequest(*static_cast<QModbusDataUnit*>(read), serverAddress);
}

void* QModbusClient_SendReadWriteRequest(void* ptr, void* read, void* write, int serverAddress)
{
	return static_cast<QModbusClient*>(ptr)->sendReadWriteRequest(*static_cast<QModbusDataUnit*>(read), *static_cast<QModbusDataUnit*>(write), serverAddress);
}

void* QModbusClient_SendWriteRequest(void* ptr, void* write, int serverAddress)
{
	return static_cast<QModbusClient*>(ptr)->sendWriteRequest(*static_cast<QModbusDataUnit*>(write), serverAddress);
}

void QModbusClient_SetNumberOfRetries(void* ptr, int number)
{
	static_cast<QModbusClient*>(ptr)->setNumberOfRetries(number);
}

void QModbusClient_SetTimeout(void* ptr, int newTimeout)
{
	static_cast<QModbusClient*>(ptr)->setTimeout(newTimeout);
}

int QModbusClient_Timeout(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->timeout();
}

void QModbusClient_ConnectTimeoutChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QModbusClient*>(ptr), static_cast<void (QModbusClient::*)(int)>(&QModbusClient::timeoutChanged), static_cast<MyQModbusClient*>(ptr), static_cast<void (MyQModbusClient::*)(int)>(&MyQModbusClient::Signal_TimeoutChanged), static_cast<Qt::ConnectionType>(t));
}

void QModbusClient_DisconnectTimeoutChanged(void* ptr)
{
	QObject::disconnect(static_cast<QModbusClient*>(ptr), static_cast<void (QModbusClient::*)(int)>(&QModbusClient::timeoutChanged), static_cast<MyQModbusClient*>(ptr), static_cast<void (MyQModbusClient::*)(int)>(&MyQModbusClient::Signal_TimeoutChanged));
}

void QModbusClient_TimeoutChanged(void* ptr, int newTimeout)
{
	static_cast<QModbusClient*>(ptr)->timeoutChanged(newTimeout);
}

void QModbusClient_Close(void* ptr)
{
	static_cast<QModbusClient*>(ptr)->close();
}

void QModbusClient_CloseDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::close();
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::close();
	} else {
	
	}
}

char QModbusClient_Open(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->open();
}

char QModbusClient_OpenDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::open();
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::open();
	} else {
	
	}
}

Q_DECLARE_METATYPE(QModbusDataUnit)
Q_DECLARE_METATYPE(QModbusDataUnit*)
void* QModbusDataUnit_NewQModbusDataUnit()
{
	return new QModbusDataUnit();
}

void* QModbusDataUnit_NewQModbusDataUnit2(long long ty)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty));
}

void* QModbusDataUnit_NewQModbusDataUnit3(long long ty, int address, unsigned short size)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty), address, size);
}

void* QModbusDataUnit_NewQModbusDataUnit4(long long ty, int address, void* data)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty), address, *static_cast<QVector<quint16>*>(data));
}

char QModbusDataUnit_IsValid(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->isValid();
}

long long QModbusDataUnit_RegisterType(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->registerType();
}

void QModbusDataUnit_SetRegisterType(void* ptr, long long ty)
{
	static_cast<QModbusDataUnit*>(ptr)->setRegisterType(static_cast<QModbusDataUnit::RegisterType>(ty));
}

void QModbusDataUnit_SetStartAddress(void* ptr, int address)
{
	static_cast<QModbusDataUnit*>(ptr)->setStartAddress(address);
}

void QModbusDataUnit_SetValue(void* ptr, int index, unsigned short value)
{
	static_cast<QModbusDataUnit*>(ptr)->setValue(index, value);
}

void QModbusDataUnit_SetValueCount(void* ptr, unsigned int newCount)
{
	static_cast<QModbusDataUnit*>(ptr)->setValueCount(newCount);
}

void QModbusDataUnit_SetValues(void* ptr, void* values)
{
	static_cast<QModbusDataUnit*>(ptr)->setValues(*static_cast<QVector<quint16>*>(values));
}

int QModbusDataUnit_StartAddress(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->startAddress();
}

unsigned short QModbusDataUnit_Value(void* ptr, int index)
{
	return static_cast<QModbusDataUnit*>(ptr)->value(index);
}

unsigned int QModbusDataUnit_ValueCount(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->valueCount();
}

struct QtSerialBus_PackedList QModbusDataUnit_Values(void* ptr)
{
	return ({ QVector<quint16>* tmpValue31b435 = new QVector<quint16>(static_cast<QModbusDataUnit*>(ptr)->values()); QtSerialBus_PackedList { tmpValue31b435, tmpValue31b435->size() }; });
}

unsigned short QModbusDataUnit___QModbusDataUnit_data_atList4(void* ptr, int i)
{
	return ({quint16 tmp = static_cast<QVector<quint16>*>(ptr)->at(i); if (i == static_cast<QVector<quint16>*>(ptr)->size()-1) { static_cast<QVector<quint16>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QModbusDataUnit___QModbusDataUnit_data_setList4(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___QModbusDataUnit_data_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>();
}

unsigned short QModbusDataUnit___setValues_values_atList(void* ptr, int i)
{
	return ({quint16 tmp = static_cast<QVector<quint16>*>(ptr)->at(i); if (i == static_cast<QVector<quint16>*>(ptr)->size()-1) { static_cast<QVector<quint16>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QModbusDataUnit___setValues_values_setList(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___setValues_values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>();
}

unsigned short QModbusDataUnit___values_atList(void* ptr, int i)
{
	return ({quint16 tmp = static_cast<QVector<quint16>*>(ptr)->at(i); if (i == static_cast<QVector<quint16>*>(ptr)->size()-1) { static_cast<QVector<quint16>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QModbusDataUnit___values_setList(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>();
}

unsigned short QModbusDataUnit___m_values_atList(void* ptr, int i)
{
	return ({quint16 tmp = static_cast<QVector<quint16>*>(ptr)->at(i); if (i == static_cast<QVector<quint16>*>(ptr)->size()-1) { static_cast<QVector<quint16>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QModbusDataUnit___m_values_setList(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___m_values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>();
}

unsigned short QModbusDataUnit___setM_values__atList(void* ptr, int i)
{
	return ({quint16 tmp = static_cast<QVector<quint16>*>(ptr)->at(i); if (i == static_cast<QVector<quint16>*>(ptr)->size()-1) { static_cast<QVector<quint16>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QModbusDataUnit___setM_values__setList(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___setM_values__newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>();
}

class MyQModbusDevice: public QModbusDevice
{
public:
	MyQModbusDevice(QObject *parent = Q_NULLPTR) : QModbusDevice(parent) {QModbusDevice_QModbusDevice_QRegisterMetaType();};
	void close() { callbackQModbusDevice_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	bool open() { return callbackQModbusDevice_Open(this) != 0; };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	 ~MyQModbusDevice() { callbackQModbusDevice_DestroyQModbusDevice(this); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusDevice*)
Q_DECLARE_METATYPE(MyQModbusDevice*)

int QModbusDevice_QModbusDevice_QRegisterMetaType(){qRegisterMetaType<QModbusDevice*>(); return qRegisterMetaType<MyQModbusDevice*>();}

void* QModbusDevice_NewQModbusDevice(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusDevice(static_cast<QObject*>(parent));
	}
}

void QModbusDevice_Close(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->close();
}

char QModbusDevice_ConnectDevice(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->connectDevice();
}

void* QModbusDevice_ConnectionParameter(void* ptr, int parameter)
{
	return new QVariant(static_cast<QModbusDevice*>(ptr)->connectionParameter(parameter));
}

void QModbusDevice_DisconnectDevice(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->disconnectDevice();
}

long long QModbusDevice_Error(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->error();
}

void QModbusDevice_ConnectErrorOccurred(void* ptr, long long t)
{
	qRegisterMetaType<QModbusDevice::Error>();
	QObject::connect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::Error)>(&QModbusDevice::errorOccurred), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::Error)>(&MyQModbusDevice::Signal_ErrorOccurred), static_cast<Qt::ConnectionType>(t));
}

void QModbusDevice_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::Error)>(&QModbusDevice::errorOccurred), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::Error)>(&MyQModbusDevice::Signal_ErrorOccurred));
}

void QModbusDevice_ErrorOccurred(void* ptr, long long error)
{
	static_cast<QModbusDevice*>(ptr)->errorOccurred(static_cast<QModbusDevice::Error>(error));
}

struct QtSerialBus_PackedString QModbusDevice_ErrorString(void* ptr)
{
	return ({ QByteArray* tb8b824 = new QByteArray(static_cast<QModbusDevice*>(ptr)->errorString().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(tb8b824->prepend("WHITESPACE").constData()+10), tb8b824->size()-10, tb8b824 }; });
}

char QModbusDevice_Open(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->open();
}

void QModbusDevice_SetConnectionParameter(void* ptr, int parameter, void* value)
{
	static_cast<QModbusDevice*>(ptr)->setConnectionParameter(parameter, *static_cast<QVariant*>(value));
}

void QModbusDevice_SetError(void* ptr, struct QtSerialBus_PackedString errorText, long long error)
{
	static_cast<QModbusDevice*>(ptr)->setError(QString::fromUtf8(errorText.data, errorText.len), static_cast<QModbusDevice::Error>(error));
}

void QModbusDevice_SetState(void* ptr, long long newState)
{
	static_cast<QModbusDevice*>(ptr)->setState(static_cast<QModbusDevice::State>(newState));
}

long long QModbusDevice_State(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->state();
}

void QModbusDevice_ConnectStateChanged(void* ptr, long long t)
{
	qRegisterMetaType<QModbusDevice::State>();
	QObject::connect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::State)>(&QModbusDevice::stateChanged), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::State)>(&MyQModbusDevice::Signal_StateChanged), static_cast<Qt::ConnectionType>(t));
}

void QModbusDevice_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::State)>(&QModbusDevice::stateChanged), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::State)>(&MyQModbusDevice::Signal_StateChanged));
}

void QModbusDevice_StateChanged(void* ptr, long long state)
{
	static_cast<QModbusDevice*>(ptr)->stateChanged(static_cast<QModbusDevice::State>(state));
}

void QModbusDevice_DestroyQModbusDevice(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->~QModbusDevice();
}

void QModbusDevice_DestroyQModbusDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QModbusDevice___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusDevice___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QModbusDevice___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QModbusDevice___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusDevice___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QModbusDevice___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusDevice___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QModbusDevice___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusDevice___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QModbusDevice_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusServer*>(ptr)->QModbusServer::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusClient*>(ptr)->QModbusClient::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QModbusDevice*>(ptr)->QModbusDevice::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QModbusDevice_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusServer*>(ptr)->QModbusServer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusClient*>(ptr)->QModbusClient::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QModbusDevice*>(ptr)->QModbusDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QModbusDevice_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusServer*>(ptr)->QModbusServer::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusClient*>(ptr)->QModbusClient::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QModbusDevice*>(ptr)->QModbusDevice::customEvent(static_cast<QEvent*>(event));
	}
}

void QModbusDevice_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::deleteLater();
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::deleteLater();
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusServer*>(ptr)->QModbusServer::deleteLater();
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::deleteLater();
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::deleteLater();
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusClient*>(ptr)->QModbusClient::deleteLater();
	} else {
		static_cast<QModbusDevice*>(ptr)->QModbusDevice::deleteLater();
	}
}

void QModbusDevice_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusServer*>(ptr)->QModbusServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusClient*>(ptr)->QModbusClient::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QModbusDevice*>(ptr)->QModbusDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QModbusDevice_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusServer*>(ptr)->QModbusServer::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusClient*>(ptr)->QModbusClient::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QModbusDevice*>(ptr)->QModbusDevice::event(static_cast<QEvent*>(e));
	}
}

char QModbusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusServer*>(ptr)->QModbusServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusClient*>(ptr)->QModbusClient::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QModbusDevice*>(ptr)->QModbusDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QModbusDevice_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::metaObject());
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::metaObject());
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QModbusServer*>(ptr)->QModbusServer::metaObject());
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::metaObject());
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::metaObject());
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QModbusClient*>(ptr)->QModbusClient::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QModbusDevice*>(ptr)->QModbusDevice::metaObject());
	}
}

void QModbusDevice_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QModbusServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusServer*>(ptr)->QModbusServer::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QModbusTcpClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QModbusRtuSerialMaster*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QModbusClient*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusClient*>(ptr)->QModbusClient::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QModbusDevice*>(ptr)->QModbusDevice::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

Q_DECLARE_METATYPE(QModbusDeviceIdentification*)
void* QModbusDeviceIdentification_NewQModbusDeviceIdentification()
{
	return new QModbusDeviceIdentification();
}

long long QModbusDeviceIdentification_ConformityLevel(void* ptr)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->conformityLevel();
}

char QModbusDeviceIdentification_Contains(void* ptr, unsigned int objectId)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->contains(objectId);
}

void* QModbusDeviceIdentification_QModbusDeviceIdentification_FromByteArray(void* ba)
{
	return new QModbusDeviceIdentification(QModbusDeviceIdentification::fromByteArray(*static_cast<QByteArray*>(ba)));
}

char QModbusDeviceIdentification_Insert(void* ptr, unsigned int objectId, void* value)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->insert(objectId, *static_cast<QByteArray*>(value));
}

char QModbusDeviceIdentification_IsValid(void* ptr)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->isValid();
}

struct QtSerialBus_PackedList QModbusDeviceIdentification_ObjectIds(void* ptr)
{
	return ({ QList<int>* tmpValue17b29a = new QList<int>(static_cast<QModbusDeviceIdentification*>(ptr)->objectIds()); QtSerialBus_PackedList { tmpValue17b29a, tmpValue17b29a->size() }; });
}

void QModbusDeviceIdentification_Remove(void* ptr, unsigned int objectId)
{
	static_cast<QModbusDeviceIdentification*>(ptr)->remove(objectId);
}

void QModbusDeviceIdentification_SetConformityLevel(void* ptr, long long level)
{
	static_cast<QModbusDeviceIdentification*>(ptr)->setConformityLevel(static_cast<QModbusDeviceIdentification::ConformityLevel>(level));
}

void* QModbusDeviceIdentification_Value(void* ptr, unsigned int objectId)
{
	return new QByteArray(static_cast<QModbusDeviceIdentification*>(ptr)->value(objectId));
}

int QModbusDeviceIdentification___objectIds_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusDeviceIdentification___objectIds_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QModbusDeviceIdentification___objectIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QModbusDeviceIdentification___m_objects_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<int, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<int, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QModbusDeviceIdentification___m_objects_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QModbusDeviceIdentification___m_objects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QByteArray>();
}

struct QtSerialBus_PackedList QModbusDeviceIdentification___m_objects_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7f22a8 = new QList<int>(static_cast<QMap<int, QByteArray>*>(ptr)->keys()); QtSerialBus_PackedList { tmpValue7f22a8, tmpValue7f22a8->size() }; });
}

void* QModbusDeviceIdentification___setM_objects__atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<int, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<int, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QModbusDeviceIdentification___setM_objects__setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QModbusDeviceIdentification___setM_objects__newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QByteArray>();
}

struct QtSerialBus_PackedList QModbusDeviceIdentification___setM_objects__keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7f22a8 = new QList<int>(static_cast<QMap<int, QByteArray>*>(ptr)->keys()); QtSerialBus_PackedList { tmpValue7f22a8, tmpValue7f22a8->size() }; });
}

int QModbusDeviceIdentification_____m_objects_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusDeviceIdentification_____m_objects_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QModbusDeviceIdentification_____m_objects_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QModbusDeviceIdentification_____setM_objects__keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusDeviceIdentification_____setM_objects__keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QModbusDeviceIdentification_____setM_objects__keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

class MyQModbusExceptionResponse: public QModbusExceptionResponse
{
public:
	MyQModbusExceptionResponse() : QModbusExceptionResponse() {QModbusExceptionResponse_QModbusExceptionResponse_QRegisterMetaType();};
	MyQModbusExceptionResponse(const QModbusPdu &pdu) : QModbusExceptionResponse(pdu) {QModbusExceptionResponse_QModbusExceptionResponse_QRegisterMetaType();};
	MyQModbusExceptionResponse(QModbusPdu::FunctionCode code, QModbusPdu::ExceptionCode ec) : QModbusExceptionResponse(code, ec) {QModbusExceptionResponse_QModbusExceptionResponse_QRegisterMetaType();};
	void setFunctionCode(QModbusPdu::FunctionCode c) { callbackQModbusPdu_SetFunctionCode(this, c); };
};

Q_DECLARE_METATYPE(QModbusExceptionResponse*)
Q_DECLARE_METATYPE(MyQModbusExceptionResponse*)

int QModbusExceptionResponse_QModbusExceptionResponse_QRegisterMetaType(){qRegisterMetaType<QModbusExceptionResponse*>(); return qRegisterMetaType<MyQModbusExceptionResponse*>();}

void* QModbusExceptionResponse_NewQModbusExceptionResponse()
{
	return new MyQModbusExceptionResponse();
}

void* QModbusExceptionResponse_NewQModbusExceptionResponse2(void* pdu)
{
	return new MyQModbusExceptionResponse(*static_cast<QModbusPdu*>(pdu));
}

void* QModbusExceptionResponse_NewQModbusExceptionResponse3(long long code, long long ec)
{
	return new MyQModbusExceptionResponse(static_cast<QModbusPdu::FunctionCode>(code), static_cast<QModbusPdu::ExceptionCode>(ec));
}

void QModbusExceptionResponse_SetExceptionCode(void* ptr, long long ec)
{
	static_cast<QModbusExceptionResponse*>(ptr)->setExceptionCode(static_cast<QModbusPdu::ExceptionCode>(ec));
}

class MyQModbusPdu: public QModbusPdu
{
public:
	MyQModbusPdu() : QModbusPdu() {QModbusPdu_QModbusPdu_QRegisterMetaType();};
	MyQModbusPdu(QModbusPdu::FunctionCode code, const QByteArray &data) : QModbusPdu(code, data) {QModbusPdu_QModbusPdu_QRegisterMetaType();};
	MyQModbusPdu(const QModbusPdu &other) : QModbusPdu(other) {QModbusPdu_QModbusPdu_QRegisterMetaType();};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, code); };
	 ~MyQModbusPdu() { callbackQModbusPdu_DestroyQModbusPdu(this); };
};

Q_DECLARE_METATYPE(QModbusPdu*)
Q_DECLARE_METATYPE(MyQModbusPdu*)

int QModbusPdu_QModbusPdu_QRegisterMetaType(){qRegisterMetaType<QModbusPdu*>(); return qRegisterMetaType<MyQModbusPdu*>();}

void* QModbusPdu_NewQModbusPdu()
{
	return new MyQModbusPdu();
}

void* QModbusPdu_NewQModbusPdu2(long long code, void* data)
{
	return new MyQModbusPdu(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

void* QModbusPdu_NewQModbusPdu3(void* other)
{
	return new MyQModbusPdu(*static_cast<QModbusPdu*>(other));
}

void* QModbusPdu_Data(void* ptr)
{
	return new QByteArray(static_cast<QModbusPdu*>(ptr)->data());
}

short QModbusPdu_DataSize(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->dataSize();
}

long long QModbusPdu_ExceptionCode(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->exceptionCode();
}

long long QModbusPdu_FunctionCode(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->functionCode();
}

char QModbusPdu_IsException(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->isException();
}

char QModbusPdu_IsValid(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->isValid();
}

void QModbusPdu_SetData(void* ptr, void* data)
{
	static_cast<QModbusPdu*>(ptr)->setData(*static_cast<QByteArray*>(data));
}

void QModbusPdu_SetFunctionCode(void* ptr, long long code)
{
	static_cast<QModbusPdu*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusPdu_SetFunctionCodeDefault(void* ptr, long long code)
{
	if (dynamic_cast<QModbusRequest*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<QModbusRequest*>(ptr)->QModbusRequest::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	} else if (dynamic_cast<QModbusExceptionResponse*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<QModbusExceptionResponse*>(ptr)->QModbusExceptionResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	} else if (dynamic_cast<QModbusResponse*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<QModbusResponse*>(ptr)->QModbusResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	} else {
		static_cast<QModbusPdu*>(ptr)->QModbusPdu::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	}
}

short QModbusPdu_Size(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->size();
}

void QModbusPdu_DestroyQModbusPdu(void* ptr)
{
	static_cast<QModbusPdu*>(ptr)->~QModbusPdu();
}

void QModbusPdu_DestroyQModbusPduDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtSerialBus_PackedString QModbusPdu_QModbusPdu_ExceptionByte()
{
	return ({ const quint8 pret9044a4 = QModbusPdu::ExceptionByte; char* t9044a4 = static_cast<char*>(static_cast<void*>(const_cast<quint8*>(&pret9044a4))); QtSerialBus_PackedString { t9044a4, -1, NULL }; });
}

void* QModbusPdu___encode_vector_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusPdu___encode_vector_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusPdu___encode_vector_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

class MyQModbusReply: public QModbusReply
{
public:
	MyQModbusReply(QModbusReply::ReplyType ty, int serverAddress, QObject *parent = Q_NULLPTR) : QModbusReply(ty, serverAddress, parent) {QModbusReply_QModbusReply_QRegisterMetaType();};
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusReply_ErrorOccurred(this, error); };
	void Signal_Finished() { callbackQModbusReply_Finished(this); };
	void childEvent(QChildEvent * event) { callbackQModbusReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusReply_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusReply_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusReply_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusReply_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusReply_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusReply_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusReply_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusReply*)
Q_DECLARE_METATYPE(MyQModbusReply*)

int QModbusReply_QModbusReply_QRegisterMetaType(){qRegisterMetaType<QModbusReply*>(); return qRegisterMetaType<MyQModbusReply*>();}

void* QModbusReply_NewQModbusReply(long long ty, int serverAddress, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QObject*>(parent));
	}
}

long long QModbusReply_Error(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->error();
}

void QModbusReply_ConnectErrorOccurred(void* ptr, long long t)
{
	qRegisterMetaType<QModbusDevice::Error>();
	QObject::connect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)(QModbusDevice::Error)>(&QModbusReply::errorOccurred), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)(QModbusDevice::Error)>(&MyQModbusReply::Signal_ErrorOccurred), static_cast<Qt::ConnectionType>(t));
}

void QModbusReply_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)(QModbusDevice::Error)>(&QModbusReply::errorOccurred), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)(QModbusDevice::Error)>(&MyQModbusReply::Signal_ErrorOccurred));
}

void QModbusReply_ErrorOccurred(void* ptr, long long error)
{
	static_cast<QModbusReply*>(ptr)->errorOccurred(static_cast<QModbusDevice::Error>(error));
}

struct QtSerialBus_PackedString QModbusReply_ErrorString(void* ptr)
{
	return ({ QByteArray* teb6a0e = new QByteArray(static_cast<QModbusReply*>(ptr)->errorString().toUtf8()); QtSerialBus_PackedString { const_cast<char*>(teb6a0e->prepend("WHITESPACE").constData()+10), teb6a0e->size()-10, teb6a0e }; });
}

void QModbusReply_ConnectFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)()>(&QModbusReply::finished), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)()>(&MyQModbusReply::Signal_Finished), static_cast<Qt::ConnectionType>(t));
}

void QModbusReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)()>(&QModbusReply::finished), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)()>(&MyQModbusReply::Signal_Finished));
}

void QModbusReply_Finished(void* ptr)
{
	static_cast<QModbusReply*>(ptr)->finished();
}

char QModbusReply_IsFinished(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->isFinished();
}

void* QModbusReply_RawResult(void* ptr)
{
	return new QModbusResponse(static_cast<QModbusReply*>(ptr)->rawResult());
}

void* QModbusReply_Result(void* ptr)
{
	return new QModbusDataUnit(static_cast<QModbusReply*>(ptr)->result());
}

int QModbusReply_ServerAddress(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->serverAddress();
}

long long QModbusReply_Type(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->type();
}

void* QModbusReply___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusReply___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QModbusReply___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QModbusReply___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusReply___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QModbusReply___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusReply___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QModbusReply___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QModbusReply___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QModbusReply_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusReply_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusReply_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::customEvent(static_cast<QEvent*>(event));
}

void QModbusReply_DeleteLaterDefault(void* ptr)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::deleteLater();
}

void QModbusReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusReply_EventDefault(void* ptr, void* e)
{
		return static_cast<QModbusReply*>(ptr)->QModbusReply::event(static_cast<QEvent*>(e));
}

char QModbusReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QModbusReply*>(ptr)->QModbusReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusReply_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QModbusReply*>(ptr)->QModbusReply::metaObject());
}

void QModbusReply_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQModbusRequest: public QModbusRequest
{
public:
	MyQModbusRequest() : QModbusRequest() {QModbusRequest_QModbusRequest_QRegisterMetaType();};
	MyQModbusRequest(const QModbusPdu &pdu) : QModbusRequest(pdu) {QModbusRequest_QModbusRequest_QRegisterMetaType();};
	MyQModbusRequest(QModbusPdu::FunctionCode code, const QByteArray &data = QByteArray()) : QModbusRequest(code, data) {QModbusRequest_QModbusRequest_QRegisterMetaType();};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, code); };
};

Q_DECLARE_METATYPE(QModbusRequest*)
Q_DECLARE_METATYPE(MyQModbusRequest*)

int QModbusRequest_QModbusRequest_QRegisterMetaType(){qRegisterMetaType<QModbusRequest*>(); return qRegisterMetaType<MyQModbusRequest*>();}

void* QModbusRequest_NewQModbusRequest()
{
	return new MyQModbusRequest();
}

void* QModbusRequest_NewQModbusRequest2(void* pdu)
{
	return new MyQModbusRequest(*static_cast<QModbusPdu*>(pdu));
}

void* QModbusRequest_NewQModbusRequest3(long long code, void* data)
{
	return new MyQModbusRequest(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

int QModbusRequest_QModbusRequest_CalculateDataSize(void* request)
{
	return QModbusRequest::calculateDataSize(*static_cast<QModbusRequest*>(request));
}

int QModbusRequest_QModbusRequest_MinimumDataSize(void* request)
{
	return QModbusRequest::minimumDataSize(*static_cast<QModbusRequest*>(request));
}

class MyQModbusResponse: public QModbusResponse
{
public:
	MyQModbusResponse() : QModbusResponse() {QModbusResponse_QModbusResponse_QRegisterMetaType();};
	MyQModbusResponse(const QModbusPdu &pdu) : QModbusResponse(pdu) {QModbusResponse_QModbusResponse_QRegisterMetaType();};
	MyQModbusResponse(QModbusPdu::FunctionCode code, const QByteArray &data = QByteArray()) : QModbusResponse(code, data) {QModbusResponse_QModbusResponse_QRegisterMetaType();};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, code); };
};

Q_DECLARE_METATYPE(QModbusResponse*)
Q_DECLARE_METATYPE(MyQModbusResponse*)

int QModbusResponse_QModbusResponse_QRegisterMetaType(){qRegisterMetaType<QModbusResponse*>(); return qRegisterMetaType<MyQModbusResponse*>();}

void* QModbusResponse_NewQModbusResponse()
{
	return new MyQModbusResponse();
}

void* QModbusResponse_NewQModbusResponse2(void* pdu)
{
	return new MyQModbusResponse(*static_cast<QModbusPdu*>(pdu));
}

void* QModbusResponse_NewQModbusResponse3(long long code, void* data)
{
	return new MyQModbusResponse(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

int QModbusResponse_QModbusResponse_CalculateDataSize(void* response)
{
	return QModbusResponse::calculateDataSize(*static_cast<QModbusResponse*>(response));
}

int QModbusResponse_QModbusResponse_MinimumDataSize(void* response)
{
	return QModbusResponse::minimumDataSize(*static_cast<QModbusResponse*>(response));
}

class MyQModbusRtuSerialMaster: public QModbusRtuSerialMaster
{
public:
	MyQModbusRtuSerialMaster(QObject *parent = Q_NULLPTR) : QModbusRtuSerialMaster(parent) {QModbusRtuSerialMaster_QModbusRtuSerialMaster_QRegisterMetaType();};
	void close() { callbackQModbusRtuSerialMaster_Close(this); };
	bool open() { return callbackQModbusRtuSerialMaster_Open(this) != 0; };
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusRtuSerialMaster*)
Q_DECLARE_METATYPE(MyQModbusRtuSerialMaster*)

int QModbusRtuSerialMaster_QModbusRtuSerialMaster_QRegisterMetaType(){qRegisterMetaType<QModbusRtuSerialMaster*>(); return qRegisterMetaType<MyQModbusRtuSerialMaster*>();}

void* QModbusRtuSerialMaster_NewQModbusRtuSerialMaster(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusRtuSerialMaster(static_cast<QObject*>(parent));
	}
}

void QModbusRtuSerialMaster_Close(void* ptr)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->close();
}

void QModbusRtuSerialMaster_CloseDefault(void* ptr)
{
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::close();
}

int QModbusRtuSerialMaster_InterFrameDelay(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->interFrameDelay();
}

char QModbusRtuSerialMaster_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->open();
}

char QModbusRtuSerialMaster_OpenDefault(void* ptr)
{
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::open();
}

void QModbusRtuSerialMaster_SetInterFrameDelay(void* ptr, int microseconds)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->setInterFrameDelay(microseconds);
}

void QModbusRtuSerialMaster_SetTurnaroundDelay(void* ptr, int turnaroundDelay)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->setTurnaroundDelay(turnaroundDelay);
}

int QModbusRtuSerialMaster_TurnaroundDelay(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->turnaroundDelay();
}

class MyQModbusRtuSerialSlave: public QModbusRtuSerialSlave
{
public:
	MyQModbusRtuSerialSlave(QObject *parent = Q_NULLPTR) : QModbusRtuSerialSlave(parent) {QModbusRtuSerialSlave_QModbusRtuSerialSlave_QRegisterMetaType();};
	void close() { callbackQModbusRtuSerialSlave_Close(this); };
	bool open() { return callbackQModbusRtuSerialSlave_Open(this) != 0; };
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	 ~MyQModbusRtuSerialSlave() { callbackQModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(this); };
	void Signal_DataWritten(QModbusDataUnit::RegisterType table, int address, int size) { callbackQModbusServer_DataWritten(this, table, address, size); };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<void*>(static_cast<const void*>(this)), newData) != 0; };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<void*>(static_cast<const void*>(this)), option)); };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusRtuSerialSlave*)
Q_DECLARE_METATYPE(MyQModbusRtuSerialSlave*)

int QModbusRtuSerialSlave_QModbusRtuSerialSlave_QRegisterMetaType(){qRegisterMetaType<QModbusRtuSerialSlave*>(); return qRegisterMetaType<MyQModbusRtuSerialSlave*>();}

void* QModbusRtuSerialSlave_NewQModbusRtuSerialSlave(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusRtuSerialSlave(static_cast<QObject*>(parent));
	}
}

void QModbusRtuSerialSlave_Close(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->close();
}

void QModbusRtuSerialSlave_CloseDefault(void* ptr)
{
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::close();
}

char QModbusRtuSerialSlave_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->open();
}

char QModbusRtuSerialSlave_OpenDefault(void* ptr)
{
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::open();
}

void QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->~QModbusRtuSerialSlave();
}

void QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlaveDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQModbusServer: public QModbusServer
{
public:
	MyQModbusServer(QObject *parent = Q_NULLPTR) : QModbusServer(parent) {QModbusServer_QModbusServer_QRegisterMetaType();};
	void Signal_DataWritten(QModbusDataUnit::RegisterType table, int address, int size) { callbackQModbusServer_DataWritten(this, table, address, size); };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<void*>(static_cast<const void*>(this)), newData) != 0; };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<void*>(static_cast<const void*>(this)), option)); };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void close() { callbackQModbusServer_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	bool open() { return callbackQModbusServer_Open(this) != 0; };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusServer*)
Q_DECLARE_METATYPE(MyQModbusServer*)

int QModbusServer_QModbusServer_QRegisterMetaType(){qRegisterMetaType<QModbusServer*>(); return qRegisterMetaType<MyQModbusServer*>();}

void* QModbusServer_NewQModbusServer(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusServer(static_cast<QObject*>(parent));
	}
}

char QModbusServer_Data(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_Data2(void* ptr, long long table, unsigned short address, unsigned short data)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit::RegisterType>(table), address, &data);
}

void QModbusServer_ConnectDataWritten(void* ptr, long long t)
{
	qRegisterMetaType<QModbusDataUnit::RegisterType>();
	QObject::connect(static_cast<QModbusServer*>(ptr), static_cast<void (QModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&QModbusServer::dataWritten), static_cast<MyQModbusServer*>(ptr), static_cast<void (MyQModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&MyQModbusServer::Signal_DataWritten), static_cast<Qt::ConnectionType>(t));
}

void QModbusServer_DisconnectDataWritten(void* ptr)
{
	QObject::disconnect(static_cast<QModbusServer*>(ptr), static_cast<void (QModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&QModbusServer::dataWritten), static_cast<MyQModbusServer*>(ptr), static_cast<void (MyQModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&MyQModbusServer::Signal_DataWritten));
}

void QModbusServer_DataWritten(void* ptr, long long table, int address, int size)
{
	static_cast<QModbusServer*>(ptr)->dataWritten(static_cast<QModbusDataUnit::RegisterType>(table), address, size);
}

void* QModbusServer_ProcessPrivateRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusServer*>(ptr)->processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusServer_ProcessPrivateRequestDefault(void* ptr, void* request)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return new QModbusResponse(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processPrivateRequest(*static_cast<QModbusPdu*>(request)));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return new QModbusResponse(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processPrivateRequest(*static_cast<QModbusPdu*>(request)));
	} else {
		return new QModbusResponse(static_cast<QModbusServer*>(ptr)->QModbusServer::processPrivateRequest(*static_cast<QModbusPdu*>(request)));
	}
}

void* QModbusServer_ProcessRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusServer*>(ptr)->processRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusServer_ProcessRequestDefault(void* ptr, void* request)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return new QModbusResponse(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processRequest(*static_cast<QModbusPdu*>(request)));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return new QModbusResponse(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processRequest(*static_cast<QModbusPdu*>(request)));
	} else {
		return new QModbusResponse(static_cast<QModbusServer*>(ptr)->QModbusServer::processRequest(*static_cast<QModbusPdu*>(request)));
	}
}

char QModbusServer_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->processesBroadcast();
}

char QModbusServer_ProcessesBroadcastDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processesBroadcast();
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processesBroadcast();
	} else {
		return static_cast<QModbusServer*>(ptr)->QModbusServer::processesBroadcast();
	}
}

char QModbusServer_ReadData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->readData(static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_ReadDataDefault(void* ptr, void* newData)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::readData(static_cast<QModbusDataUnit*>(newData));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::readData(static_cast<QModbusDataUnit*>(newData));
	} else {
		return static_cast<QModbusServer*>(ptr)->QModbusServer::readData(static_cast<QModbusDataUnit*>(newData));
	}
}

int QModbusServer_ServerAddress(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->serverAddress();
}

char QModbusServer_SetData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->setData(*static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_SetData2(void* ptr, long long table, unsigned short address, unsigned short data)
{
	return static_cast<QModbusServer*>(ptr)->setData(static_cast<QModbusDataUnit::RegisterType>(table), address, data);
}

void QModbusServer_SetServerAddress(void* ptr, int serverAddress)
{
	static_cast<QModbusServer*>(ptr)->setServerAddress(serverAddress);
}

char QModbusServer_SetValue(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusServer*>(ptr)->setValue(option, *static_cast<QVariant*>(newValue));
}

char QModbusServer_SetValueDefault(void* ptr, int option, void* newValue)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::setValue(option, *static_cast<QVariant*>(newValue));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::setValue(option, *static_cast<QVariant*>(newValue));
	} else {
		return static_cast<QModbusServer*>(ptr)->QModbusServer::setValue(option, *static_cast<QVariant*>(newValue));
	}
}

void* QModbusServer_Value(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusServer*>(ptr)->value(option));
}

void* QModbusServer_ValueDefault(void* ptr, int option)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::value(option));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::value(option));
	} else {
		return new QVariant(static_cast<QModbusServer*>(ptr)->QModbusServer::value(option));
	}
}

char QModbusServer_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_WriteDataDefault(void* ptr, void* newData)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::writeData(*static_cast<QModbusDataUnit*>(newData));
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::writeData(*static_cast<QModbusDataUnit*>(newData));
	} else {
		return static_cast<QModbusServer*>(ptr)->QModbusServer::writeData(*static_cast<QModbusDataUnit*>(newData));
	}
}

void QModbusServer_Close(void* ptr)
{
	static_cast<QModbusServer*>(ptr)->close();
}

void QModbusServer_CloseDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::close();
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::close();
	} else {
	
	}
}

char QModbusServer_Open(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->open();
}

char QModbusServer_OpenDefault(void* ptr)
{
	if (dynamic_cast<QModbusTcpServer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::open();
	} else if (dynamic_cast<QModbusRtuSerialSlave*>(static_cast<QObject*>(ptr))) {
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::open();
	} else {
	
	}
}

class MyQModbusTcpClient: public QModbusTcpClient
{
public:
	MyQModbusTcpClient(QObject *parent = Q_NULLPTR) : QModbusTcpClient(parent) {QModbusTcpClient_QModbusTcpClient_QRegisterMetaType();};
	void close() { callbackQModbusTcpClient_Close(this); };
	bool open() { return callbackQModbusTcpClient_Open(this) != 0; };
	 ~MyQModbusTcpClient() { callbackQModbusTcpClient_DestroyQModbusTcpClient(this); };
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusTcpClient*)
Q_DECLARE_METATYPE(MyQModbusTcpClient*)

int QModbusTcpClient_QModbusTcpClient_QRegisterMetaType(){qRegisterMetaType<QModbusTcpClient*>(); return qRegisterMetaType<MyQModbusTcpClient*>();}

void* QModbusTcpClient_NewQModbusTcpClient(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusTcpClient(static_cast<QObject*>(parent));
	}
}

void QModbusTcpClient_Close(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->close();
}

void QModbusTcpClient_CloseDefault(void* ptr)
{
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::close();
}

char QModbusTcpClient_Open(void* ptr)
{
	return static_cast<QModbusTcpClient*>(ptr)->open();
}

char QModbusTcpClient_OpenDefault(void* ptr)
{
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::open();
}

void QModbusTcpClient_DestroyQModbusTcpClient(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->~QModbusTcpClient();
}

void QModbusTcpClient_DestroyQModbusTcpClientDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQModbusTcpConnectionObserver: public QModbusTcpConnectionObserver
{
public:
	bool acceptNewConnection(QTcpSocket * newClient) { return callbackQModbusTcpConnectionObserver_AcceptNewConnection(this, newClient) != 0; };
};

Q_DECLARE_METATYPE(QModbusTcpConnectionObserver*)
Q_DECLARE_METATYPE(MyQModbusTcpConnectionObserver*)

int QModbusTcpConnectionObserver_QModbusTcpConnectionObserver_QRegisterMetaType(){qRegisterMetaType<QModbusTcpConnectionObserver*>(); return qRegisterMetaType<MyQModbusTcpConnectionObserver*>();}

char QModbusTcpConnectionObserver_AcceptNewConnection(void* ptr, void* newClient)
{
	return static_cast<QModbusTcpConnectionObserver*>(ptr)->acceptNewConnection(static_cast<QTcpSocket*>(newClient));
}

class MyQModbusTcpServer: public QModbusTcpServer
{
public:
	MyQModbusTcpServer(QObject *parent = Q_NULLPTR) : QModbusTcpServer(parent) {QModbusTcpServer_QModbusTcpServer_QRegisterMetaType();};
	void close() { callbackQModbusTcpServer_Close(this); };
	void Signal_ModbusClientDisconnected(QTcpSocket * modbusClient) { callbackQModbusTcpServer_ModbusClientDisconnected(this, modbusClient); };
	bool open() { return callbackQModbusTcpServer_Open(this) != 0; };
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	 ~MyQModbusTcpServer() { callbackQModbusTcpServer_DestroyQModbusTcpServer(this); };
	void Signal_DataWritten(QModbusDataUnit::RegisterType table, int address, int size) { callbackQModbusServer_DataWritten(this, table, address, size); };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<void*>(static_cast<const void*>(this)), newData) != 0; };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<void*>(static_cast<const void*>(this)), option)); };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QModbusTcpServer*)
Q_DECLARE_METATYPE(MyQModbusTcpServer*)

int QModbusTcpServer_QModbusTcpServer_QRegisterMetaType(){qRegisterMetaType<QModbusTcpServer*>(); return qRegisterMetaType<MyQModbusTcpServer*>();}

void* QModbusTcpServer_NewQModbusTcpServer(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusTcpServer(static_cast<QObject*>(parent));
	}
}

void QModbusTcpServer_Close(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->close();
}

void QModbusTcpServer_CloseDefault(void* ptr)
{
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::close();
}

void QModbusTcpServer_InstallConnectionObserver(void* ptr, void* observer)
{
	static_cast<QModbusTcpServer*>(ptr)->installConnectionObserver(static_cast<QModbusTcpConnectionObserver*>(observer));
}

void QModbusTcpServer_ConnectModbusClientDisconnected(void* ptr, long long t)
{
	QObject::connect(static_cast<QModbusTcpServer*>(ptr), static_cast<void (QModbusTcpServer::*)(QTcpSocket *)>(&QModbusTcpServer::modbusClientDisconnected), static_cast<MyQModbusTcpServer*>(ptr), static_cast<void (MyQModbusTcpServer::*)(QTcpSocket *)>(&MyQModbusTcpServer::Signal_ModbusClientDisconnected), static_cast<Qt::ConnectionType>(t));
}

void QModbusTcpServer_DisconnectModbusClientDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QModbusTcpServer*>(ptr), static_cast<void (QModbusTcpServer::*)(QTcpSocket *)>(&QModbusTcpServer::modbusClientDisconnected), static_cast<MyQModbusTcpServer*>(ptr), static_cast<void (MyQModbusTcpServer::*)(QTcpSocket *)>(&MyQModbusTcpServer::Signal_ModbusClientDisconnected));
}

void QModbusTcpServer_ModbusClientDisconnected(void* ptr, void* modbusClient)
{
	static_cast<QModbusTcpServer*>(ptr)->modbusClientDisconnected(static_cast<QTcpSocket*>(modbusClient));
}

char QModbusTcpServer_Open(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->open();
}

char QModbusTcpServer_OpenDefault(void* ptr)
{
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::open();
}

void QModbusTcpServer_DestroyQModbusTcpServer(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->~QModbusTcpServer();
}

void QModbusTcpServer_DestroyQModbusTcpServerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

