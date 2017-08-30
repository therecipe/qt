// +build !minimal

#define protected public
#define private public

#include "serialbus.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QCanBus>
#include <QCanBusDevice>
#include <QCanBusDeviceInfo>
#include <QCanBusFactoryV2>
#include <QCanBusFrame>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
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
#include <QModbusTcpServer>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQCanBus: public QCanBus
{
public:
	bool event(QEvent * e) { return callbackQCanBus_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCanBus_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQCanBus_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCanBus_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCanBus_CustomEvent(this, event); };
	void deleteLater() { callbackQCanBus_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCanBus_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCanBus_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQCanBus_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCanBus_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCanBus_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQCanBus*)

int QCanBus_QCanBus_QRegisterMetaType(){qRegisterMetaType<QCanBus*>(); return qRegisterMetaType<MyQCanBus*>();}

void* QCanBus_QCanBus_Instance()
{
	return QCanBus::instance();
}

void* QCanBus_CreateDevice(void* ptr, struct QtSerialBus_PackedString plugin, struct QtSerialBus_PackedString interfaceName, struct QtSerialBus_PackedString errorMessage)
{
	return static_cast<QCanBus*>(ptr)->createDevice(QString::fromUtf8(plugin.data, plugin.len), QString::fromUtf8(interfaceName.data, interfaceName.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)));
}

struct QtSerialBus_PackedList QCanBus_AvailableDevices(void* ptr, struct QtSerialBus_PackedString plugin, struct QtSerialBus_PackedString errorMessage)
{
	return ({ QList<QCanBusDeviceInfo>* tmpValue = new QList<QCanBusDeviceInfo>(static_cast<QCanBus*>(ptr)->availableDevices(QString::fromUtf8(plugin.data, plugin.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)))); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSerialBus_PackedString QCanBus_Plugins(void* ptr)
{
	return ({ QByteArray t48ad25 = static_cast<QCanBus*>(ptr)->plugins().join("|").toUtf8(); QtSerialBus_PackedString { const_cast<char*>(t48ad25.prepend("WHITESPACE").constData()+10), t48ad25.size()-10 }; });
}

void* QCanBus___availableDevices_atList(void* ptr, int i)
{
	return new QCanBusDeviceInfo(static_cast<QList<QCanBusDeviceInfo>*>(ptr)->at(i));
}

void QCanBus___availableDevices_setList(void* ptr, void* i)
{
	static_cast<QList<QCanBusDeviceInfo>*>(ptr)->append(*static_cast<QCanBusDeviceInfo*>(i));
}

void* QCanBus___availableDevices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCanBusDeviceInfo>;
}

void* QCanBus___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QCanBus___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCanBus___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QCanBus___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QCanBus___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QCanBus___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QCanBus___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QCanBus___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QCanBus___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QCanBus___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QCanBus___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBus___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QCanBus_EventDefault(void* ptr, void* e)
{
		return static_cast<QCanBus*>(ptr)->QCanBus::event(static_cast<QEvent*>(e));
}

char QCanBus_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QCanBus*>(ptr)->QCanBus::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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

void QCanBus_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QCanBus*>(ptr)->QCanBus::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QCanBus_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QCanBus*>(ptr)->QCanBus::metaObject());
}

class MyQCanBusDevice: public QCanBusDevice
{
public:
	MyQCanBusDevice(QObject *parent = nullptr) : QCanBusDevice(parent) {QCanBusDevice_QCanBusDevice_QRegisterMetaType();};
	QString interpretErrorFrame(const QCanBusFrame & frame) { return ({ QtSerialBus_PackedString tempVal = callbackQCanBusDevice_InterpretErrorFrame(this, const_cast<QCanBusFrame*>(&frame)); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool open() { return callbackQCanBusDevice_Open(this) != 0; };
	bool waitForFramesReceived(int msecs) { return callbackQCanBusDevice_WaitForFramesReceived(this, msecs) != 0; };
	bool waitForFramesWritten(int msecs) { return callbackQCanBusDevice_WaitForFramesWritten(this, msecs) != 0; };
	bool writeFrame(const QCanBusFrame & frame) { return callbackQCanBusDevice_WriteFrame(this, const_cast<QCanBusFrame*>(&frame)) != 0; };
	void close() { callbackQCanBusDevice_Close(this); };
	void Signal_ErrorOccurred(QCanBusDevice::CanBusError error) { callbackQCanBusDevice_ErrorOccurred(this, error); };
	void Signal_FramesReceived() { callbackQCanBusDevice_FramesReceived(this); };
	void Signal_FramesWritten(qint64 framesCount) { callbackQCanBusDevice_FramesWritten(this, framesCount); };
	void setConfigurationParameter(int key, const QVariant & value) { callbackQCanBusDevice_SetConfigurationParameter(this, key, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QCanBusDevice::CanBusDeviceState state) { callbackQCanBusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQCanBusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCanBusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQCanBusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCanBusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQCanBusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCanBusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQCanBusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCanBusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCanBusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQCanBusDevice*)

int QCanBusDevice_QCanBusDevice_QRegisterMetaType(){qRegisterMetaType<QCanBusDevice*>(); return qRegisterMetaType<MyQCanBusDevice*>();}

void* QCanBusDevice_NewQCanBusDevice(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCanBusDevice(static_cast<QWindow*>(parent));
	} else {
		return new MyQCanBusDevice(static_cast<QObject*>(parent));
	}
}

struct QtSerialBus_PackedString QCanBusDevice_InterpretErrorFrame(void* ptr, void* frame)
{
	return ({ QByteArray t5dc8ff = static_cast<QCanBusDevice*>(ptr)->interpretErrorFrame(*static_cast<QCanBusFrame*>(frame)).toUtf8(); QtSerialBus_PackedString { const_cast<char*>(t5dc8ff.prepend("WHITESPACE").constData()+10), t5dc8ff.size()-10 }; });
}

char QCanBusDevice_ConnectDevice(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->connectDevice();
}

char QCanBusDevice_Open(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->open();
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

void QCanBusDevice_Close(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->close();
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

void QCanBusDevice_ConnectErrorOccurred(void* ptr)
{
	qRegisterMetaType<QCanBusDevice::CanBusError>();
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusError)>(&QCanBusDevice::errorOccurred), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusError)>(&MyQCanBusDevice::Signal_ErrorOccurred));
}

void QCanBusDevice_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusError)>(&QCanBusDevice::errorOccurred), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusError)>(&MyQCanBusDevice::Signal_ErrorOccurred));
}

void QCanBusDevice_ErrorOccurred(void* ptr, long long error)
{
	static_cast<QCanBusDevice*>(ptr)->errorOccurred(static_cast<QCanBusDevice::CanBusError>(error));
}

void QCanBusDevice_ConnectFramesReceived(void* ptr)
{
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)()>(&QCanBusDevice::framesReceived), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)()>(&MyQCanBusDevice::Signal_FramesReceived));
}

void QCanBusDevice_DisconnectFramesReceived(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)()>(&QCanBusDevice::framesReceived), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)()>(&MyQCanBusDevice::Signal_FramesReceived));
}

void QCanBusDevice_FramesReceived(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->framesReceived();
}

void QCanBusDevice_ConnectFramesWritten(void* ptr)
{
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(qint64)>(&QCanBusDevice::framesWritten), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(qint64)>(&MyQCanBusDevice::Signal_FramesWritten));
}

void QCanBusDevice_DisconnectFramesWritten(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(qint64)>(&QCanBusDevice::framesWritten), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(qint64)>(&MyQCanBusDevice::Signal_FramesWritten));
}

void QCanBusDevice_FramesWritten(void* ptr, long long framesCount)
{
	static_cast<QCanBusDevice*>(ptr)->framesWritten(framesCount);
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

void QCanBusDevice_ConnectStateChanged(void* ptr)
{
	qRegisterMetaType<QCanBusDevice::CanBusDeviceState>();
	QObject::connect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&QCanBusDevice::stateChanged), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&MyQCanBusDevice::Signal_StateChanged));
}

void QCanBusDevice_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCanBusDevice*>(ptr), static_cast<void (QCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&QCanBusDevice::stateChanged), static_cast<MyQCanBusDevice*>(ptr), static_cast<void (MyQCanBusDevice::*)(QCanBusDevice::CanBusDeviceState)>(&MyQCanBusDevice::Signal_StateChanged));
}

void QCanBusDevice_StateChanged(void* ptr, long long state)
{
	static_cast<QCanBusDevice*>(ptr)->stateChanged(static_cast<QCanBusDevice::CanBusDeviceState>(state));
}

long long QCanBusDevice_State(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->state();
}

long long QCanBusDevice_Error(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->error();
}

struct QtSerialBus_PackedString QCanBusDevice_ErrorString(void* ptr)
{
	return ({ QByteArray t646153 = static_cast<QCanBusDevice*>(ptr)->errorString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(t646153.prepend("WHITESPACE").constData()+10), t646153.size()-10 }; });
}

void* QCanBusDevice_ConfigurationParameter(void* ptr, int key)
{
	return new QVariant(static_cast<QCanBusDevice*>(ptr)->configurationParameter(key));
}

struct QtSerialBus_PackedList QCanBusDevice_ConfigurationKeys(void* ptr)
{
	return ({ QVector<int>* tmpValue = new QVector<int>(static_cast<QCanBusDevice*>(ptr)->configurationKeys()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
}

char QCanBusDevice_HasOutgoingFrames(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->hasOutgoingFrames();
}

long long QCanBusDevice_FramesAvailable(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->framesAvailable();
}

long long QCanBusDevice_FramesToWrite(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->framesToWrite();
}

void QCanBusDevice___enqueueReceivedFrames_newFrames_setList(void* ptr, void* i)
{
	static_cast<QVector<QCanBusFrame>*>(ptr)->append(*static_cast<QCanBusFrame*>(i));
}

void* QCanBusDevice___enqueueReceivedFrames_newFrames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QCanBusFrame>;
}

int QCanBusDevice___configurationKeys_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QCanBusDevice___configurationKeys_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QCanBusDevice___configurationKeys_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* QCanBusDevice___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QCanBusDevice___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCanBusDevice___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QCanBusDevice___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QCanBusDevice___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QCanBusDevice___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QCanBusDevice___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QCanBusDevice___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QCanBusDevice___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QCanBusDevice___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QCanBusDevice___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCanBusDevice___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QCanBusDevice_EventDefault(void* ptr, void* e)
{
		return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::event(static_cast<QEvent*>(e));
}

char QCanBusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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

void QCanBusDevice_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QCanBusDevice_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::metaObject());
}

void QCanBusDeviceInfo_Swap(void* ptr, void* other)
{
	static_cast<QCanBusDeviceInfo*>(ptr)->swap(*static_cast<QCanBusDeviceInfo*>(other));
}

void QCanBusDeviceInfo_DestroyQCanBusDeviceInfo(void* ptr)
{
	static_cast<QCanBusDeviceInfo*>(ptr)->~QCanBusDeviceInfo();
}

struct QtSerialBus_PackedString QCanBusDeviceInfo_Name(void* ptr)
{
	return ({ QByteArray tdc7ae5 = static_cast<QCanBusDeviceInfo*>(ptr)->name().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(tdc7ae5.prepend("WHITESPACE").constData()+10), tdc7ae5.size()-10 }; });
}

char QCanBusDeviceInfo_HasFlexibleDataRate(void* ptr)
{
	return static_cast<QCanBusDeviceInfo*>(ptr)->hasFlexibleDataRate();
}

char QCanBusDeviceInfo_IsVirtual(void* ptr)
{
	return static_cast<QCanBusDeviceInfo*>(ptr)->isVirtual();
}

class MyQCanBusFactoryV2: public QCanBusFactoryV2
{
public:
	QCanBusDevice * createDevice(const QString & interfaceName, QString * errorMessage) const { QByteArray tf83f00 = interfaceName.toUtf8(); QtSerialBus_PackedString interfaceNamePacked = { const_cast<char*>(tf83f00.prepend("WHITESPACE").constData()+10), tf83f00.size()-10 };QByteArray t3f2abc = errorMessage->toUtf8(); QtSerialBus_PackedString errorMessagePacked = { const_cast<char*>(t3f2abc.prepend("WHITESPACE").constData()+10), t3f2abc.size()-10 };return static_cast<QCanBusDevice*>(callbackQCanBusFactoryV2_CreateDevice(const_cast<void*>(static_cast<const void*>(this)), interfaceNamePacked, errorMessagePacked)); };
	QList<QCanBusDeviceInfo> availableDevices(QString * errorMessage) const { QByteArray t3f2abc = errorMessage->toUtf8(); QtSerialBus_PackedString errorMessagePacked = { const_cast<char*>(t3f2abc.prepend("WHITESPACE").constData()+10), t3f2abc.size()-10 };return *static_cast<QList<QCanBusDeviceInfo>*>(callbackQCanBusFactoryV2_AvailableDevices(const_cast<void*>(static_cast<const void*>(this)), errorMessagePacked)); };
};

void* QCanBusFactoryV2_CreateDevice(void* ptr, struct QtSerialBus_PackedString interfaceName, struct QtSerialBus_PackedString errorMessage)
{
	return static_cast<QCanBusFactoryV2*>(ptr)->createDevice(QString::fromUtf8(interfaceName.data, interfaceName.len), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)));
}

struct QtSerialBus_PackedList QCanBusFactoryV2_AvailableDevices(void* ptr, struct QtSerialBus_PackedString errorMessage)
{
	return ({ QList<QCanBusDeviceInfo>* tmpValue = new QList<QCanBusDeviceInfo>(static_cast<QCanBusFactoryV2*>(ptr)->availableDevices(new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)))); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
}

void* QCanBusFactoryV2___availableDevices_atList(void* ptr, int i)
{
	return new QCanBusDeviceInfo(static_cast<QList<QCanBusDeviceInfo>*>(ptr)->at(i));
}

void QCanBusFactoryV2___availableDevices_setList(void* ptr, void* i)
{
	static_cast<QList<QCanBusDeviceInfo>*>(ptr)->append(*static_cast<QCanBusDeviceInfo*>(i));
}

void* QCanBusFactoryV2___availableDevices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCanBusDeviceInfo>;
}

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

void QCanBusFrame_SetPayload(void* ptr, void* data)
{
	static_cast<QCanBusFrame*>(ptr)->setPayload(*static_cast<QByteArray*>(data));
}

long long QCanBusFrame_Error(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->error();
}

long long QCanBusFrame_FrameType(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->frameType();
}

void* QCanBusFrame_Payload(void* ptr)
{
	return new QByteArray(static_cast<QCanBusFrame*>(ptr)->payload());
}

struct QtSerialBus_PackedString QCanBusFrame_ToString(void* ptr)
{
	return ({ QByteArray tda41b1 = static_cast<QCanBusFrame*>(ptr)->toString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(tda41b1.prepend("WHITESPACE").constData()+10), tda41b1.size()-10 }; });
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

char QCanBusFrame_IsValid(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->isValid();
}

unsigned int QCanBusFrame_FrameId(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->frameId();
}

class MyQModbusClient: public QModbusClient
{
public:
	MyQModbusClient(QObject *parent = nullptr) : QModbusClient(parent) {QModbusClient_QModbusClient_QRegisterMetaType();};
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	bool open() { return callbackQModbusClient_Open(this) != 0; };
	void close() { callbackQModbusClient_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusClient*)

int QModbusClient_QModbusClient_QRegisterMetaType(){qRegisterMetaType<QModbusClient*>(); return qRegisterMetaType<MyQModbusClient*>();}

void* QModbusClient_NewQModbusClient(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusClient(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusClient(static_cast<QObject*>(parent));
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

void QModbusClient_SetNumberOfRetries(void* ptr, int number)
{
	static_cast<QModbusClient*>(ptr)->setNumberOfRetries(number);
}

void QModbusClient_SetTimeout(void* ptr, int newTimeout)
{
	static_cast<QModbusClient*>(ptr)->setTimeout(newTimeout);
}

void QModbusClient_ConnectTimeoutChanged(void* ptr)
{
	QObject::connect(static_cast<QModbusClient*>(ptr), static_cast<void (QModbusClient::*)(int)>(&QModbusClient::timeoutChanged), static_cast<MyQModbusClient*>(ptr), static_cast<void (MyQModbusClient::*)(int)>(&MyQModbusClient::Signal_TimeoutChanged));
}

void QModbusClient_DisconnectTimeoutChanged(void* ptr)
{
	QObject::disconnect(static_cast<QModbusClient*>(ptr), static_cast<void (QModbusClient::*)(int)>(&QModbusClient::timeoutChanged), static_cast<MyQModbusClient*>(ptr), static_cast<void (MyQModbusClient::*)(int)>(&MyQModbusClient::Signal_TimeoutChanged));
}

void QModbusClient_TimeoutChanged(void* ptr, int newTimeout)
{
	static_cast<QModbusClient*>(ptr)->timeoutChanged(newTimeout);
}

int QModbusClient_NumberOfRetries(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->numberOfRetries();
}

int QModbusClient_Timeout(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->timeout();
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

void* QModbusDataUnit_NewQModbusDataUnit()
{
	return new QModbusDataUnit();
}

void* QModbusDataUnit_NewQModbusDataUnit2(long long ty)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty));
}

void* QModbusDataUnit_NewQModbusDataUnit4(long long ty, int address, void* data)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty), address, *static_cast<QVector<quint16>*>(data));
}

void* QModbusDataUnit_NewQModbusDataUnit3(long long ty, int address, unsigned short size)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty), address, size);
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

struct QtSerialBus_PackedList QModbusDataUnit_Values(void* ptr)
{
	return ({ QVector<quint16>* tmpValue = new QVector<quint16>(static_cast<QModbusDataUnit*>(ptr)->values()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
}

long long QModbusDataUnit_RegisterType(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->registerType();
}

char QModbusDataUnit_IsValid(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->isValid();
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

unsigned short QModbusDataUnit___QModbusDataUnit_data_atList4(void* ptr, int i)
{
	return static_cast<QVector<quint16>*>(ptr)->at(i);
}

void QModbusDataUnit___QModbusDataUnit_data_setList4(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___QModbusDataUnit_data_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>;
}

unsigned short QModbusDataUnit___setValues_values_atList(void* ptr, int i)
{
	return static_cast<QVector<quint16>*>(ptr)->at(i);
}

void QModbusDataUnit___setValues_values_setList(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___setValues_values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>;
}

unsigned short QModbusDataUnit___values_atList(void* ptr, int i)
{
	return static_cast<QVector<quint16>*>(ptr)->at(i);
}

void QModbusDataUnit___values_setList(void* ptr, unsigned short i)
{
	static_cast<QVector<quint16>*>(ptr)->append(i);
}

void* QModbusDataUnit___values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint16>;
}

class MyQModbusDevice: public QModbusDevice
{
public:
	MyQModbusDevice(QObject *parent = nullptr) : QModbusDevice(parent) {QModbusDevice_QModbusDevice_QRegisterMetaType();};
	bool open() { return callbackQModbusDevice_Open(this) != 0; };
	void close() { callbackQModbusDevice_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusDevice*)

int QModbusDevice_QModbusDevice_QRegisterMetaType(){qRegisterMetaType<QModbusDevice*>(); return qRegisterMetaType<MyQModbusDevice*>();}

void* QModbusDevice_NewQModbusDevice(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusDevice(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusDevice(static_cast<QObject*>(parent));
	}
}

char QModbusDevice_ConnectDevice(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->connectDevice();
}

char QModbusDevice_Open(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->open();
}

void QModbusDevice_Close(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->close();
}

void QModbusDevice_DisconnectDevice(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->disconnectDevice();
}

void QModbusDevice_ConnectErrorOccurred(void* ptr)
{
	qRegisterMetaType<QModbusDevice::Error>();
	QObject::connect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::Error)>(&QModbusDevice::errorOccurred), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::Error)>(&MyQModbusDevice::Signal_ErrorOccurred));
}

void QModbusDevice_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::Error)>(&QModbusDevice::errorOccurred), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::Error)>(&MyQModbusDevice::Signal_ErrorOccurred));
}

void QModbusDevice_ErrorOccurred(void* ptr, long long error)
{
	static_cast<QModbusDevice*>(ptr)->errorOccurred(static_cast<QModbusDevice::Error>(error));
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

void QModbusDevice_ConnectStateChanged(void* ptr)
{
	qRegisterMetaType<QModbusDevice::State>();
	QObject::connect(static_cast<QModbusDevice*>(ptr), static_cast<void (QModbusDevice::*)(QModbusDevice::State)>(&QModbusDevice::stateChanged), static_cast<MyQModbusDevice*>(ptr), static_cast<void (MyQModbusDevice::*)(QModbusDevice::State)>(&MyQModbusDevice::Signal_StateChanged));
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

long long QModbusDevice_Error(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->error();
}

struct QtSerialBus_PackedString QModbusDevice_ErrorString(void* ptr)
{
	return ({ QByteArray tb8b824 = static_cast<QModbusDevice*>(ptr)->errorString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(tb8b824.prepend("WHITESPACE").constData()+10), tb8b824.size()-10 }; });
}

void* QModbusDevice_ConnectionParameter(void* ptr, int parameter)
{
	return new QVariant(static_cast<QModbusDevice*>(ptr)->connectionParameter(parameter));
}

long long QModbusDevice_State(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->state();
}

void* QModbusDevice___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusDevice___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusDevice___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QModbusDevice___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusDevice___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QModbusDevice___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusDevice___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QModbusDevice___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusDevice___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QModbusDevice___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusDevice___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusDevice___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
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

void* QModbusDeviceIdentification_QModbusDeviceIdentification_FromByteArray(void* ba)
{
	return new QModbusDeviceIdentification(QModbusDeviceIdentification::fromByteArray(*static_cast<QByteArray*>(ba)));
}

void* QModbusDeviceIdentification_NewQModbusDeviceIdentification()
{
	return new QModbusDeviceIdentification();
}

char QModbusDeviceIdentification_Insert(void* ptr, unsigned int objectId, void* value)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->insert(objectId, *static_cast<QByteArray*>(value));
}

void QModbusDeviceIdentification_Remove(void* ptr, unsigned int objectId)
{
	static_cast<QModbusDeviceIdentification*>(ptr)->remove(objectId);
}

void QModbusDeviceIdentification_SetConformityLevel(void* ptr, long long level)
{
	static_cast<QModbusDeviceIdentification*>(ptr)->setConformityLevel(static_cast<QModbusDeviceIdentification::ConformityLevel>(level));
}

long long QModbusDeviceIdentification_ConformityLevel(void* ptr)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->conformityLevel();
}

void* QModbusDeviceIdentification_Value(void* ptr, unsigned int objectId)
{
	return new QByteArray(static_cast<QModbusDeviceIdentification*>(ptr)->value(objectId));
}

struct QtSerialBus_PackedList QModbusDeviceIdentification_ObjectIds(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QModbusDeviceIdentification*>(ptr)->objectIds()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
}

char QModbusDeviceIdentification_Contains(void* ptr, unsigned int objectId)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->contains(objectId);
}

char QModbusDeviceIdentification_IsValid(void* ptr)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->isValid();
}

int QModbusDeviceIdentification___objectIds_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QModbusDeviceIdentification___objectIds_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QModbusDeviceIdentification___objectIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

class MyQModbusExceptionResponse: public QModbusExceptionResponse
{
public:
	MyQModbusExceptionResponse() : QModbusExceptionResponse() {};
	MyQModbusExceptionResponse(FunctionCode code, ExceptionCode ec) : QModbusExceptionResponse(code, ec) {};
	MyQModbusExceptionResponse(const QModbusPdu &pdu) : QModbusExceptionResponse(pdu) {};
	void setFunctionCode(QModbusPdu::FunctionCode c) { callbackQModbusPdu_SetFunctionCode(this, c); };
};

void* QModbusExceptionResponse_NewQModbusExceptionResponse()
{
	return new MyQModbusExceptionResponse();
}

void* QModbusExceptionResponse_NewQModbusExceptionResponse3(long long code, long long ec)
{
	return new MyQModbusExceptionResponse(static_cast<QModbusPdu::FunctionCode>(code), static_cast<QModbusPdu::ExceptionCode>(ec));
}

void* QModbusExceptionResponse_NewQModbusExceptionResponse2(void* pdu)
{
	return new MyQModbusExceptionResponse(*static_cast<QModbusPdu*>(pdu));
}

void QModbusExceptionResponse_SetExceptionCode(void* ptr, long long ec)
{
	static_cast<QModbusExceptionResponse*>(ptr)->setExceptionCode(static_cast<QModbusPdu::ExceptionCode>(ec));
}

class MyQModbusPdu: public QModbusPdu
{
public:
	MyQModbusPdu() : QModbusPdu() {};
	MyQModbusPdu(FunctionCode code, const QByteArray &data) : QModbusPdu(code, data) {};
	MyQModbusPdu(const QModbusPdu &other) : QModbusPdu(other) {};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, code); };
	 ~MyQModbusPdu() { callbackQModbusPdu_DestroyQModbusPdu(this); };
};

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
	if (dynamic_cast<QModbusExceptionResponse*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<QModbusExceptionResponse*>(ptr)->QModbusExceptionResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	} else if (dynamic_cast<QModbusResponse*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<QModbusResponse*>(ptr)->QModbusResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	} else if (dynamic_cast<QModbusRequest*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<QModbusRequest*>(ptr)->QModbusRequest::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	} else {
		static_cast<QModbusPdu*>(ptr)->QModbusPdu::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
	}
}

void QModbusPdu_DestroyQModbusPdu(void* ptr)
{
	static_cast<QModbusPdu*>(ptr)->~QModbusPdu();
}

void QModbusPdu_DestroyQModbusPduDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QModbusPdu_ExceptionCode(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->exceptionCode();
}

long long QModbusPdu_FunctionCode(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->functionCode();
}

void* QModbusPdu_Data(void* ptr)
{
	return new QByteArray(static_cast<QModbusPdu*>(ptr)->data());
}

char QModbusPdu_IsException(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->isException();
}

char QModbusPdu_IsValid(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->isValid();
}

short QModbusPdu_DataSize(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->dataSize();
}

short QModbusPdu_Size(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->size();
}

void* QModbusPdu___encode_vector_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusPdu___encode_vector_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusPdu___encode_vector_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

class MyQModbusReply: public QModbusReply
{
public:
	MyQModbusReply(ReplyType ty, int serverAddress, QObject *parent = nullptr) : QModbusReply(ty, serverAddress, parent) {QModbusReply_QModbusReply_QRegisterMetaType();};
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusReply_ErrorOccurred(this, error); };
	void Signal_Finished() { callbackQModbusReply_Finished(this); };
	bool event(QEvent * e) { return callbackQModbusReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusReply_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusReply_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusReply_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusReply_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusReply_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusReply_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusReply_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusReply*)

int QModbusReply_QModbusReply_QRegisterMetaType(){qRegisterMetaType<QModbusReply*>(); return qRegisterMetaType<MyQModbusReply*>();}

void* QModbusReply_NewQModbusReply(long long ty, int serverAddress, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QObject*>(parent));
	}
}

void QModbusReply_ConnectErrorOccurred(void* ptr)
{
	qRegisterMetaType<QModbusDevice::Error>();
	QObject::connect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)(QModbusDevice::Error)>(&QModbusReply::errorOccurred), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)(QModbusDevice::Error)>(&MyQModbusReply::Signal_ErrorOccurred));
}

void QModbusReply_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)(QModbusDevice::Error)>(&QModbusReply::errorOccurred), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)(QModbusDevice::Error)>(&MyQModbusReply::Signal_ErrorOccurred));
}

void QModbusReply_ErrorOccurred(void* ptr, long long error)
{
	static_cast<QModbusReply*>(ptr)->errorOccurred(static_cast<QModbusDevice::Error>(error));
}

void QModbusReply_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)()>(&QModbusReply::finished), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)()>(&MyQModbusReply::Signal_Finished));
}

void QModbusReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QModbusReply*>(ptr), static_cast<void (QModbusReply::*)()>(&QModbusReply::finished), static_cast<MyQModbusReply*>(ptr), static_cast<void (MyQModbusReply::*)()>(&MyQModbusReply::Signal_Finished));
}

void QModbusReply_Finished(void* ptr)
{
	static_cast<QModbusReply*>(ptr)->finished();
}

void* QModbusReply_Result(void* ptr)
{
	return new QModbusDataUnit(static_cast<QModbusReply*>(ptr)->result());
}

long long QModbusReply_Error(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->error();
}

void* QModbusReply_RawResult(void* ptr)
{
	return new QModbusResponse(static_cast<QModbusReply*>(ptr)->rawResult());
}

struct QtSerialBus_PackedString QModbusReply_ErrorString(void* ptr)
{
	return ({ QByteArray teb6a0e = static_cast<QModbusReply*>(ptr)->errorString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(teb6a0e.prepend("WHITESPACE").constData()+10), teb6a0e.size()-10 }; });
}

long long QModbusReply_Type(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->type();
}

char QModbusReply_IsFinished(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->isFinished();
}

int QModbusReply_ServerAddress(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->serverAddress();
}

void* QModbusReply___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusReply___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusReply___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QModbusReply___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusReply___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QModbusReply___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusReply___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QModbusReply___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusReply___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QModbusReply___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusReply___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QModbusReply_EventDefault(void* ptr, void* e)
{
		return static_cast<QModbusReply*>(ptr)->QModbusReply::event(static_cast<QEvent*>(e));
}

char QModbusReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QModbusReply*>(ptr)->QModbusReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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

void QModbusReply_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QModbusReply*>(ptr)->QModbusReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QModbusReply_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QModbusReply*>(ptr)->QModbusReply::metaObject());
}

class MyQModbusRequest: public QModbusRequest
{
public:
	MyQModbusRequest() : QModbusRequest() {};
	MyQModbusRequest(FunctionCode code, const QByteArray &data = QByteArray()) : QModbusRequest(code, data) {};
	MyQModbusRequest(const QModbusPdu &pdu) : QModbusRequest(pdu) {};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, code); };
};

void* QModbusRequest_NewQModbusRequest()
{
	return new MyQModbusRequest();
}

void* QModbusRequest_NewQModbusRequest3(long long code, void* data)
{
	return new MyQModbusRequest(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

void* QModbusRequest_NewQModbusRequest2(void* pdu)
{
	return new MyQModbusRequest(*static_cast<QModbusPdu*>(pdu));
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
	MyQModbusResponse() : QModbusResponse() {};
	MyQModbusResponse(FunctionCode code, const QByteArray &data = QByteArray()) : QModbusResponse(code, data) {};
	MyQModbusResponse(const QModbusPdu &pdu) : QModbusResponse(pdu) {};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, code); };
};

void* QModbusResponse_NewQModbusResponse()
{
	return new MyQModbusResponse();
}

void* QModbusResponse_NewQModbusResponse3(long long code, void* data)
{
	return new MyQModbusResponse(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

void* QModbusResponse_NewQModbusResponse2(void* pdu)
{
	return new MyQModbusResponse(*static_cast<QModbusPdu*>(pdu));
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
	MyQModbusRtuSerialMaster(QObject *parent = nullptr) : QModbusRtuSerialMaster(parent) {QModbusRtuSerialMaster_QModbusRtuSerialMaster_QRegisterMetaType();};
	bool open() { return callbackQModbusRtuSerialMaster_Open(this) != 0; };
	void close() { callbackQModbusRtuSerialMaster_Close(this); };
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusRtuSerialMaster*)

int QModbusRtuSerialMaster_QModbusRtuSerialMaster_QRegisterMetaType(){qRegisterMetaType<QModbusRtuSerialMaster*>(); return qRegisterMetaType<MyQModbusRtuSerialMaster*>();}

void* QModbusRtuSerialMaster_NewQModbusRtuSerialMaster(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialMaster(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusRtuSerialMaster(static_cast<QObject*>(parent));
	}
}

char QModbusRtuSerialMaster_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->open();
}

char QModbusRtuSerialMaster_OpenDefault(void* ptr)
{
		return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::open();
}

void QModbusRtuSerialMaster_Close(void* ptr)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->close();
}

void QModbusRtuSerialMaster_CloseDefault(void* ptr)
{
		static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::close();
}

void QModbusRtuSerialMaster_SetInterFrameDelay(void* ptr, int microseconds)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->setInterFrameDelay(microseconds);
}

int QModbusRtuSerialMaster_InterFrameDelay(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->interFrameDelay();
}

class MyQModbusRtuSerialSlave: public QModbusRtuSerialSlave
{
public:
	MyQModbusRtuSerialSlave(QObject *parent = nullptr) : QModbusRtuSerialSlave(parent) {QModbusRtuSerialSlave_QModbusRtuSerialSlave_QRegisterMetaType();};
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool open() { return callbackQModbusRtuSerialSlave_Open(this) != 0; };
	void close() { callbackQModbusRtuSerialSlave_Close(this); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void Signal_DataWritten(QModbusDataUnit::RegisterType regist, int address, int size) { callbackQModbusServer_DataWritten(this, regist, address, size); };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<void*>(static_cast<const void*>(this)), option)); };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<void*>(static_cast<const void*>(this)), newData) != 0; };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusRtuSerialSlave*)

int QModbusRtuSerialSlave_QModbusRtuSerialSlave_QRegisterMetaType(){qRegisterMetaType<QModbusRtuSerialSlave*>(); return qRegisterMetaType<MyQModbusRtuSerialSlave*>();}

void* QModbusRtuSerialSlave_NewQModbusRtuSerialSlave(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusRtuSerialSlave(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusRtuSerialSlave(static_cast<QObject*>(parent));
	}
}

char QModbusRtuSerialSlave_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->open();
}

char QModbusRtuSerialSlave_OpenDefault(void* ptr)
{
		return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::open();
}

void QModbusRtuSerialSlave_Close(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->close();
}

void QModbusRtuSerialSlave_CloseDefault(void* ptr)
{
		static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::close();
}

void QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->~QModbusRtuSerialSlave();
}

class MyQModbusServer: public QModbusServer
{
public:
	MyQModbusServer(QObject *parent = nullptr) : QModbusServer(parent) {QModbusServer_QModbusServer_QRegisterMetaType();};
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void Signal_DataWritten(QModbusDataUnit::RegisterType regist, int address, int size) { callbackQModbusServer_DataWritten(this, regist, address, size); };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<void*>(static_cast<const void*>(this)), option)); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<void*>(static_cast<const void*>(this)), newData) != 0; };
	bool open() { return callbackQModbusServer_Open(this) != 0; };
	void close() { callbackQModbusServer_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusServer*)

int QModbusServer_QModbusServer_QRegisterMetaType(){qRegisterMetaType<QModbusServer*>(); return qRegisterMetaType<MyQModbusServer*>();}

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

void* QModbusServer_NewQModbusServer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusServer(static_cast<QObject*>(parent));
	}
}

char QModbusServer_SetData2(void* ptr, long long table, unsigned short address, unsigned short data)
{
	return static_cast<QModbusServer*>(ptr)->setData(static_cast<QModbusDataUnit::RegisterType>(table), address, data);
}

char QModbusServer_SetData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->setData(*static_cast<QModbusDataUnit*>(newData));
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

void QModbusServer_ConnectDataWritten(void* ptr)
{
	qRegisterMetaType<QModbusDataUnit::RegisterType>();
	QObject::connect(static_cast<QModbusServer*>(ptr), static_cast<void (QModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&QModbusServer::dataWritten), static_cast<MyQModbusServer*>(ptr), static_cast<void (MyQModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&MyQModbusServer::Signal_DataWritten));
}

void QModbusServer_DisconnectDataWritten(void* ptr)
{
	QObject::disconnect(static_cast<QModbusServer*>(ptr), static_cast<void (QModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&QModbusServer::dataWritten), static_cast<MyQModbusServer*>(ptr), static_cast<void (MyQModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&MyQModbusServer::Signal_DataWritten));
}

void QModbusServer_DataWritten(void* ptr, long long regist, int address, int size)
{
	static_cast<QModbusServer*>(ptr)->dataWritten(static_cast<QModbusDataUnit::RegisterType>(regist), address, size);
}

void QModbusServer_SetServerAddress(void* ptr, int serverAddress)
{
	static_cast<QModbusServer*>(ptr)->setServerAddress(serverAddress);
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

char QModbusServer_Data(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_Data2(void* ptr, long long table, unsigned short address, unsigned short data)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit::RegisterType>(table), address, &data);
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

class MyQModbusTcpClient: public QModbusTcpClient
{
public:
	MyQModbusTcpClient(QObject *parent = nullptr) : QModbusTcpClient(parent) {QModbusTcpClient_QModbusTcpClient_QRegisterMetaType();};
	bool open() { return callbackQModbusTcpClient_Open(this) != 0; };
	void close() { callbackQModbusTcpClient_Close(this); };
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusTcpClient*)

int QModbusTcpClient_QModbusTcpClient_QRegisterMetaType(){qRegisterMetaType<QModbusTcpClient*>(); return qRegisterMetaType<MyQModbusTcpClient*>();}

void* QModbusTcpClient_NewQModbusTcpClient(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpClient(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusTcpClient(static_cast<QObject*>(parent));
	}
}

char QModbusTcpClient_Open(void* ptr)
{
	return static_cast<QModbusTcpClient*>(ptr)->open();
}

char QModbusTcpClient_OpenDefault(void* ptr)
{
		return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::open();
}

void QModbusTcpClient_Close(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->close();
}

void QModbusTcpClient_CloseDefault(void* ptr)
{
		static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::close();
}

void QModbusTcpClient_DestroyQModbusTcpClient(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->~QModbusTcpClient();
}

class MyQModbusTcpServer: public QModbusTcpServer
{
public:
	MyQModbusTcpServer(QObject *parent = nullptr) : QModbusTcpServer(parent) {QModbusTcpServer_QModbusTcpServer_QRegisterMetaType();};
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool open() { return callbackQModbusTcpServer_Open(this) != 0; };
	void close() { callbackQModbusTcpServer_Close(this); };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void Signal_DataWritten(QModbusDataUnit::RegisterType regist, int address, int size) { callbackQModbusServer_DataWritten(this, regist, address, size); };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<void*>(static_cast<const void*>(this)), option)); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<void*>(static_cast<const void*>(this)), newData) != 0; };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQModbusDevice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSerialBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQModbusDevice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQModbusTcpServer*)

int QModbusTcpServer_QModbusTcpServer_QRegisterMetaType(){qRegisterMetaType<QModbusTcpServer*>(); return qRegisterMetaType<MyQModbusTcpServer*>();}

void* QModbusTcpServer_NewQModbusTcpServer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQModbusTcpServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQModbusTcpServer(static_cast<QObject*>(parent));
	}
}

char QModbusTcpServer_Open(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->open();
}

char QModbusTcpServer_OpenDefault(void* ptr)
{
		return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::open();
}

void QModbusTcpServer_Close(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->close();
}

void QModbusTcpServer_CloseDefault(void* ptr)
{
		static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::close();
}

void QModbusTcpServer_DestroyQModbusTcpServer(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->~QModbusTcpServer();
}

