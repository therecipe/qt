// +build !minimal

#define protected public
#define private public

#include "serialbus.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCanBus>
#include <QCanBusDevice>
#include <QCanBusFactory>
#include <QCanBusFrame>
#include <QChildEvent>
#include <QEvent>
#include <QList>
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
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>

void* QCanBus_CreateDevice(void* ptr, void* plugin, char* interfaceName)
{
	return static_cast<QCanBus*>(ptr)->createDevice(*static_cast<QByteArray*>(plugin), QString(interfaceName));
}

void* QCanBus_QCanBus_Instance()
{
	return QCanBus::instance();
}

struct QtSerialBus_PackedList QCanBus_Plugins(void* ptr)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QCanBus*>(ptr)->plugins()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
}

void* QCanBus___plugins_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QCanBus___plugins_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCanBus___plugins_newList(void* ptr)
{
	return new QList<QByteArray>;
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
	return new QList<QObject *>;
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
	return new QList<QObject*>;
}

void QCanBus_TimerEvent(void* ptr, void* event)
{
	static_cast<QCanBus*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCanBus_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCanBus*>(ptr)->QCanBus::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCanBus_ChildEvent(void* ptr, void* event)
{
	static_cast<QCanBus*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCanBus_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCanBus*>(ptr)->QCanBus::childEvent(static_cast<QChildEvent*>(event));
}

void QCanBus_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCanBus*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBus_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCanBus*>(ptr)->QCanBus::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBus_CustomEvent(void* ptr, void* event)
{
	static_cast<QCanBus*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCanBus_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCanBus*>(ptr)->QCanBus::customEvent(static_cast<QEvent*>(event));
}

void QCanBus_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCanBus*>(ptr), "deleteLater");
}

void QCanBus_DeleteLaterDefault(void* ptr)
{
	static_cast<QCanBus*>(ptr)->QCanBus::deleteLater();
}

void QCanBus_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCanBus*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBus_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCanBus*>(ptr)->QCanBus::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCanBus_Event(void* ptr, void* e)
{
	return static_cast<QCanBus*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCanBus_EventDefault(void* ptr, void* e)
{
	return static_cast<QCanBus*>(ptr)->QCanBus::event(static_cast<QEvent*>(e));
}

char QCanBus_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCanBus*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCanBus_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCanBus*>(ptr)->QCanBus::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCanBus_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCanBus*>(ptr)->metaObject());
}

void* QCanBus_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCanBus*>(ptr)->QCanBus::metaObject());
}

class MyQCanBusDevice: public QCanBusDevice
{
public:
	MyQCanBusDevice(QObject *parent) : QCanBusDevice(parent) {};
	void close() { callbackQCanBusDevice_Close(this); };
	void Signal_ErrorOccurred(QCanBusDevice::CanBusError error) { callbackQCanBusDevice_ErrorOccurred(this, error); };
	void Signal_FramesReceived() { callbackQCanBusDevice_FramesReceived(this); };
	void Signal_FramesWritten(qint64 framesCount) { callbackQCanBusDevice_FramesWritten(this, framesCount); };
	QString interpretErrorFrame(const QCanBusFrame & frame) { return QString(callbackQCanBusDevice_InterpretErrorFrame(this, const_cast<QCanBusFrame*>(&frame))); };
	bool open() { return callbackQCanBusDevice_Open(this) != 0; };
	void setConfigurationParameter(int key, const QVariant & value) { callbackQCanBusDevice_SetConfigurationParameter(this, key, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QCanBusDevice::CanBusDeviceState state) { callbackQCanBusDevice_StateChanged(this, state); };
	bool writeFrame(const QCanBusFrame & frame) { return callbackQCanBusDevice_WriteFrame(this, const_cast<QCanBusFrame*>(&frame)) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQCanBusDevice_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCanBusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCanBusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQCanBusDevice_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCanBusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCanBusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCanBusDevice_MetaObject(const_cast<MyQCanBusDevice*>(this))); };
};

long long QCanBusDevice_FramesAvailable(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->framesAvailable();
}

long long QCanBusDevice_FramesToWrite(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->framesToWrite();
}

void* QCanBusDevice_NewQCanBusDevice(void* parent)
{
	return new MyQCanBusDevice(static_cast<QObject*>(parent));
}

void QCanBusDevice_Close(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->close();
}

struct QtSerialBus_PackedList QCanBusDevice_ConfigurationKeys(void* ptr)
{
	return ({ QVector<int>* tmpValue = new QVector<int>(static_cast<QCanBusDevice*>(ptr)->configurationKeys()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
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

void QCanBusDevice_ConnectErrorOccurred(void* ptr)
{
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

struct QtSerialBus_PackedString QCanBusDevice_ErrorString(void* ptr)
{
	return ({ QByteArray t646153 = static_cast<QCanBusDevice*>(ptr)->errorString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(t646153.prepend("WHITESPACE").constData()+10), t646153.size()-10 }; });
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

char QCanBusDevice_HasOutgoingFrames(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->hasOutgoingFrames();
}

struct QtSerialBus_PackedString QCanBusDevice_InterpretErrorFrame(void* ptr, void* frame)
{
	return ({ QByteArray t5dc8ff = static_cast<QCanBusDevice*>(ptr)->interpretErrorFrame(*static_cast<QCanBusFrame*>(frame)).toUtf8(); QtSerialBus_PackedString { const_cast<char*>(t5dc8ff.prepend("WHITESPACE").constData()+10), t5dc8ff.size()-10 }; });
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

void QCanBusDevice_SetError(void* ptr, char* errorText, long long errorId)
{
	static_cast<QCanBusDevice*>(ptr)->setError(QString(errorText), static_cast<QCanBusDevice::CanBusError>(errorId));
}

void QCanBusDevice_SetState(void* ptr, long long newState)
{
	static_cast<QCanBusDevice*>(ptr)->setState(static_cast<QCanBusDevice::CanBusDeviceState>(newState));
}

long long QCanBusDevice_State(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->state();
}

void QCanBusDevice_ConnectStateChanged(void* ptr)
{
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

char QCanBusDevice_WriteFrame(void* ptr, void* frame)
{
	return static_cast<QCanBusDevice*>(ptr)->writeFrame(*static_cast<QCanBusFrame*>(frame));
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
	return new QVector<int>;
}

void QCanBusDevice___enqueueReceivedFrames_newFrames_setList(void* ptr, void* i)
{
	static_cast<QVector<QCanBusFrame>*>(ptr)->append(*static_cast<QCanBusFrame*>(i));
}

void* QCanBusDevice___enqueueReceivedFrames_newFrames_newList(void* ptr)
{
	return new QVector<QCanBusFrame>;
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
	return new QList<QObject *>;
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
	return new QList<QObject*>;
}

void QCanBusDevice_TimerEvent(void* ptr, void* event)
{
	static_cast<QCanBusDevice*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCanBusDevice_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCanBusDevice_ChildEvent(void* ptr, void* event)
{
	static_cast<QCanBusDevice*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCanBusDevice_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::childEvent(static_cast<QChildEvent*>(event));
}

void QCanBusDevice_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCanBusDevice*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBusDevice_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBusDevice_CustomEvent(void* ptr, void* event)
{
	static_cast<QCanBusDevice*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCanBusDevice_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::customEvent(static_cast<QEvent*>(event));
}

void QCanBusDevice_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCanBusDevice*>(ptr), "deleteLater");
}

void QCanBusDevice_DeleteLaterDefault(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::deleteLater();
}

void QCanBusDevice_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCanBusDevice*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCanBusDevice_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCanBusDevice_Event(void* ptr, void* e)
{
	return static_cast<QCanBusDevice*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCanBusDevice_EventDefault(void* ptr, void* e)
{
	return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::event(static_cast<QEvent*>(e));
}

char QCanBusDevice_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCanBusDevice*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCanBusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCanBusDevice_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCanBusDevice*>(ptr)->metaObject());
}

void* QCanBusDevice_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::metaObject());
}

class MyQCanBusFactory: public QCanBusFactory
{
public:
	QCanBusDevice * createDevice(const QString & interfaceName) const { QByteArray tf83f00 = interfaceName.toUtf8(); QtSerialBus_PackedString interfaceNamePacked = { const_cast<char*>(tf83f00.prepend("WHITESPACE").constData()+10), tf83f00.size()-10 };return static_cast<QCanBusDevice*>(callbackQCanBusFactory_CreateDevice(const_cast<MyQCanBusFactory*>(this), interfaceNamePacked)); };
};

void* QCanBusFactory_CreateDevice(void* ptr, char* interfaceName)
{
	return static_cast<QCanBusFactory*>(ptr)->createDevice(QString(interfaceName));
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

char QCanBusFrame_HasExtendedFrameFormat(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasExtendedFrameFormat();
}

char QCanBusFrame_IsValid(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->isValid();
}

void* QCanBusFrame_Payload(void* ptr)
{
	return new QByteArray(static_cast<QCanBusFrame*>(ptr)->payload());
}

void QCanBusFrame_SetError(void* ptr, long long error)
{
	static_cast<QCanBusFrame*>(ptr)->setError(static_cast<QCanBusFrame::FrameError>(error));
}

void QCanBusFrame_SetExtendedFrameFormat(void* ptr, char isExtended)
{
	static_cast<QCanBusFrame*>(ptr)->setExtendedFrameFormat(isExtended != 0);
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

class MyQModbusClient: public QModbusClient
{
public:
	MyQModbusClient(QObject *parent) : QModbusClient(parent) {};
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, newTimeout); };
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, const_cast<QModbusResponse*>(&response), data) != 0; };
	void close() { callbackQModbusClient_Close(this); };
	bool open() { return callbackQModbusClient_Open(this) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQModbusClient_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQModbusClient_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusClient_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusClient_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusClient_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusClient_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusClient_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusClient_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusClient_MetaObject(const_cast<MyQModbusClient*>(this))); };
};

int QModbusClient_Timeout(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->timeout();
}

void* QModbusClient_NewQModbusClient(void* parent)
{
	return new MyQModbusClient(static_cast<QObject*>(parent));
}

int QModbusClient_NumberOfRetries(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->numberOfRetries();
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

char QModbusClient_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusClient_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusClient_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusClient_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

void QModbusClient_Close(void* ptr)
{
	static_cast<QModbusClient*>(ptr)->close();
}

char QModbusClient_Open(void* ptr)
{
	return static_cast<QModbusClient*>(ptr)->open();
}

void QModbusClient_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusClient*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusClient_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusClient*>(ptr)->QModbusClient::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusClient_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusClient*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusClient_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusClient*>(ptr)->QModbusClient::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusClient_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusClient*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusClient_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusClient*>(ptr)->QModbusClient::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusClient_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusClient*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusClient_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusClient*>(ptr)->QModbusClient::customEvent(static_cast<QEvent*>(event));
}

void QModbusClient_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusClient*>(ptr), "deleteLater");
}

void QModbusClient_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusClient*>(ptr)->QModbusClient::deleteLater();
}

void QModbusClient_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusClient*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusClient_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusClient*>(ptr)->QModbusClient::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusClient_Event(void* ptr, void* e)
{
	return static_cast<QModbusClient*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusClient_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::event(static_cast<QEvent*>(e));
}

char QModbusClient_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusClient*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusClient_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusClient_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusClient*>(ptr)->metaObject());
}

void* QModbusClient_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusClient*>(ptr)->QModbusClient::metaObject());
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
	return ({ QVector<quint16>* tmpValue = new QVector<quint16>(static_cast<QModbusDataUnit*>(ptr)->values()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
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
	return new QVector<quint16>;
}

class MyQModbusDevice: public QModbusDevice
{
public:
	MyQModbusDevice(QObject *parent) : QModbusDevice(parent) {};
	void close() { callbackQModbusDevice_Close(this); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, error); };
	bool open() { return callbackQModbusDevice_Open(this) != 0; };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, state); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<MyQModbusDevice*>(this))); };
};

void* QModbusDevice_NewQModbusDevice(void* parent)
{
	return new MyQModbusDevice(static_cast<QObject*>(parent));
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

void QModbusDevice_ConnectErrorOccurred(void* ptr)
{
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

struct QtSerialBus_PackedString QModbusDevice_ErrorString(void* ptr)
{
	return ({ QByteArray tb8b824 = static_cast<QModbusDevice*>(ptr)->errorString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(tb8b824.prepend("WHITESPACE").constData()+10), tb8b824.size()-10 }; });
}

char QModbusDevice_Open(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->open();
}

void QModbusDevice_SetConnectionParameter(void* ptr, int parameter, void* value)
{
	static_cast<QModbusDevice*>(ptr)->setConnectionParameter(parameter, *static_cast<QVariant*>(value));
}

void QModbusDevice_SetError(void* ptr, char* errorText, long long error)
{
	static_cast<QModbusDevice*>(ptr)->setError(QString(errorText), static_cast<QModbusDevice::Error>(error));
}

void QModbusDevice_SetState(void* ptr, long long newState)
{
	static_cast<QModbusDevice*>(ptr)->setState(static_cast<QModbusDevice::State>(newState));
}

long long QModbusDevice_State(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->state();
}

void QModbusDevice_ConnectStateChanged(void* ptr)
{
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
	return new QList<QObject *>;
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
	return new QList<QObject*>;
}

void QModbusDevice_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusDevice*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusDevice_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusDevice*>(ptr)->QModbusDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusDevice_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusDevice*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusDevice_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusDevice*>(ptr)->QModbusDevice::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusDevice_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusDevice*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusDevice_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusDevice*>(ptr)->QModbusDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusDevice_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusDevice*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusDevice_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusDevice*>(ptr)->QModbusDevice::customEvent(static_cast<QEvent*>(event));
}

void QModbusDevice_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusDevice*>(ptr), "deleteLater");
}

void QModbusDevice_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->QModbusDevice::deleteLater();
}

void QModbusDevice_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusDevice*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusDevice_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusDevice*>(ptr)->QModbusDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusDevice_Event(void* ptr, void* e)
{
	return static_cast<QModbusDevice*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusDevice_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusDevice*>(ptr)->QModbusDevice::event(static_cast<QEvent*>(e));
}

char QModbusDevice_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusDevice*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusDevice*>(ptr)->QModbusDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusDevice_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusDevice*>(ptr)->metaObject());
}

void* QModbusDevice_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusDevice*>(ptr)->QModbusDevice::metaObject());
}

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
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QModbusDeviceIdentification*>(ptr)->objectIds()); QtSerialBus_PackedList { tmpValue, tmpValue->size() }; });
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
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QModbusDeviceIdentification___objectIds_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QModbusDeviceIdentification___objectIds_newList(void* ptr)
{
	return new QList<int>;
}

class MyQModbusExceptionResponse: public QModbusExceptionResponse
{
public:
	MyQModbusExceptionResponse() : QModbusExceptionResponse() {};
	MyQModbusExceptionResponse(FunctionCode code, ExceptionCode ec) : QModbusExceptionResponse(code, ec) {};
	MyQModbusExceptionResponse(const QModbusPdu &pdu) : QModbusExceptionResponse(pdu) {};
	void setFunctionCode(QModbusPdu::FunctionCode c) { callbackQModbusExceptionResponse_SetFunctionCode(this, c); };
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

void QModbusExceptionResponse_SetFunctionCode(void* ptr, long long c)
{
	static_cast<QModbusExceptionResponse*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(c));
}

void QModbusExceptionResponse_SetFunctionCodeDefault(void* ptr, long long c)
{
	static_cast<QModbusExceptionResponse*>(ptr)->QModbusExceptionResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(c));
}

void* QModbusExceptionResponse___encode_vector_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusExceptionResponse___encode_vector_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusExceptionResponse___encode_vector_newList(void* ptr)
{
	return new QList<QObject*>;
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

short QModbusPdu_DataSize(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->dataSize();
}

void* QModbusPdu_Data(void* ptr)
{
	return new QByteArray(static_cast<QModbusPdu*>(ptr)->data());
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
	static_cast<QModbusPdu*>(ptr)->QModbusPdu::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
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
	return new QList<QObject*>;
}

class MyQModbusReply: public QModbusReply
{
public:
	MyQModbusReply(ReplyType type, int serverAddress, QObject *parent) : QModbusReply(type, serverAddress, parent) {};
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusReply_ErrorOccurred(this, error); };
	void Signal_Finished() { callbackQModbusReply_Finished(this); };
	void timerEvent(QTimerEvent * event) { callbackQModbusReply_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQModbusReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusReply_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusReply_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusReply_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusReply_MetaObject(const_cast<MyQModbusReply*>(this))); };
};

void* QModbusReply_NewQModbusReply(long long ty, int serverAddress, void* parent)
{
	return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QObject*>(parent));
}

long long QModbusReply_Error(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->error();
}

void QModbusReply_ConnectErrorOccurred(void* ptr)
{
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

struct QtSerialBus_PackedString QModbusReply_ErrorString(void* ptr)
{
	return ({ QByteArray teb6a0e = static_cast<QModbusReply*>(ptr)->errorString().toUtf8(); QtSerialBus_PackedString { const_cast<char*>(teb6a0e.prepend("WHITESPACE").constData()+10), teb6a0e.size()-10 }; });
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
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusReply___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusReply___children_newList(void* ptr)
{
	return new QList<QObject *>;
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
	return new QList<QObject*>;
}

void QModbusReply_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusReply*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusReply_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusReply*>(ptr)->QModbusReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusReply_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusReply*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusReply_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusReply*>(ptr)->QModbusReply::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusReply_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusReply*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusReply_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusReply*>(ptr)->QModbusReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusReply_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusReply*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusReply_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusReply*>(ptr)->QModbusReply::customEvent(static_cast<QEvent*>(event));
}

void QModbusReply_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusReply*>(ptr), "deleteLater");
}

void QModbusReply_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusReply*>(ptr)->QModbusReply::deleteLater();
}

void QModbusReply_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusReply*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusReply*>(ptr)->QModbusReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusReply_Event(void* ptr, void* e)
{
	return static_cast<QModbusReply*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusReply_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusReply*>(ptr)->QModbusReply::event(static_cast<QEvent*>(e));
}

char QModbusReply_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusReply*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusReply*>(ptr)->QModbusReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusReply_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusReply*>(ptr)->metaObject());
}

void* QModbusReply_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusReply*>(ptr)->QModbusReply::metaObject());
}

void* QModbusRequest_NewQModbusRequest()
{
	return new QModbusRequest();
}

void* QModbusRequest_NewQModbusRequest3(long long code, void* data)
{
	return new QModbusRequest(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

void* QModbusRequest_NewQModbusRequest2(void* pdu)
{
	return new QModbusRequest(*static_cast<QModbusPdu*>(pdu));
}

int QModbusRequest_QModbusRequest_CalculateDataSize(void* request)
{
	return QModbusRequest::calculateDataSize(*static_cast<QModbusRequest*>(request));
}

int QModbusRequest_QModbusRequest_MinimumDataSize(void* request)
{
	return QModbusRequest::minimumDataSize(*static_cast<QModbusRequest*>(request));
}

void QModbusRequest_SetFunctionCode(void* ptr, long long code)
{
	static_cast<QModbusRequest*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusRequest_SetFunctionCodeDefault(void* ptr, long long code)
{
	static_cast<QModbusRequest*>(ptr)->QModbusRequest::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void* QModbusResponse_NewQModbusResponse()
{
	return new QModbusResponse();
}

void* QModbusResponse_NewQModbusResponse3(long long code, void* data)
{
	return new QModbusResponse(static_cast<QModbusPdu::FunctionCode>(code), *static_cast<QByteArray*>(data));
}

void* QModbusResponse_NewQModbusResponse2(void* pdu)
{
	return new QModbusResponse(*static_cast<QModbusPdu*>(pdu));
}

int QModbusResponse_QModbusResponse_CalculateDataSize(void* response)
{
	return QModbusResponse::calculateDataSize(*static_cast<QModbusResponse*>(response));
}

int QModbusResponse_QModbusResponse_MinimumDataSize(void* response)
{
	return QModbusResponse::minimumDataSize(*static_cast<QModbusResponse*>(response));
}

void QModbusResponse_SetFunctionCode(void* ptr, long long code)
{
	static_cast<QModbusResponse*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusResponse_SetFunctionCodeDefault(void* ptr, long long code)
{
	static_cast<QModbusResponse*>(ptr)->QModbusResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void* QModbusRtuSerialMaster_NewQModbusRtuSerialMaster(void* parent)
{
	return new QModbusRtuSerialMaster(static_cast<QObject*>(parent));
}

int QModbusRtuSerialMaster_InterFrameDelay(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->interFrameDelay();
}

void QModbusRtuSerialMaster_SetInterFrameDelay(void* ptr, int microseconds)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->setInterFrameDelay(microseconds);
}

void* QModbusRtuSerialMaster___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusRtuSerialMaster___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialMaster___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QModbusRtuSerialMaster___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusRtuSerialMaster___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusRtuSerialMaster___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QModbusRtuSerialMaster___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusRtuSerialMaster___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialMaster___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusRtuSerialMaster___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusRtuSerialMaster___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialMaster___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusRtuSerialMaster___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusRtuSerialMaster___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialMaster___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

char QModbusRtuSerialMaster_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusRtuSerialMaster_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusRtuSerialMaster_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusRtuSerialMaster_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

void QModbusRtuSerialMaster_Close(void* ptr)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->close();
}

char QModbusRtuSerialMaster_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->open();
}

void QModbusRtuSerialMaster_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusRtuSerialMaster_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusRtuSerialMaster_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusRtuSerialMaster_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusRtuSerialMaster_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusRtuSerialMaster_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusRtuSerialMaster_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusRtuSerialMaster_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::customEvent(static_cast<QEvent*>(event));
}

void QModbusRtuSerialMaster_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusRtuSerialMaster*>(ptr), "deleteLater");
}

void QModbusRtuSerialMaster_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::deleteLater();
}

void QModbusRtuSerialMaster_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusRtuSerialMaster_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusRtuSerialMaster_Event(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusRtuSerialMaster_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::event(static_cast<QEvent*>(e));
}

char QModbusRtuSerialMaster_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusRtuSerialMaster_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusRtuSerialMaster_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusRtuSerialMaster*>(ptr)->metaObject());
}

void* QModbusRtuSerialMaster_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::metaObject());
}

void* QModbusRtuSerialSlave_NewQModbusRtuSerialSlave(void* parent)
{
	return new QModbusRtuSerialSlave(static_cast<QObject*>(parent));
}

void QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->~QModbusRtuSerialSlave();
}

void* QModbusRtuSerialSlave___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusRtuSerialSlave___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialSlave___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QModbusRtuSerialSlave___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusRtuSerialSlave___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusRtuSerialSlave___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QModbusRtuSerialSlave___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusRtuSerialSlave___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialSlave___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusRtuSerialSlave___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusRtuSerialSlave___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialSlave___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusRtuSerialSlave___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusRtuSerialSlave___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusRtuSerialSlave___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusRtuSerialSlave_ProcessPrivateRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusRtuSerialSlave*>(ptr)->processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusRtuSerialSlave_ProcessPrivateRequestDefault(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusRtuSerialSlave_ProcessRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusRtuSerialSlave*>(ptr)->processRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusRtuSerialSlave_ProcessRequestDefault(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processRequest(*static_cast<QModbusPdu*>(request)));
}

char QModbusRtuSerialSlave_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->processesBroadcast();
}

char QModbusRtuSerialSlave_ProcessesBroadcastDefault(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processesBroadcast();
}

char QModbusRtuSerialSlave_ReadData(void* ptr, void* newData)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->readData(static_cast<QModbusDataUnit*>(newData));
}

char QModbusRtuSerialSlave_ReadDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::readData(static_cast<QModbusDataUnit*>(newData));
}

char QModbusRtuSerialSlave_SetValue(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->setValue(option, *static_cast<QVariant*>(newValue));
}

char QModbusRtuSerialSlave_SetValueDefault(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::setValue(option, *static_cast<QVariant*>(newValue));
}

void* QModbusRtuSerialSlave_Value(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusRtuSerialSlave*>(ptr)->value(option));
}

void* QModbusRtuSerialSlave_ValueDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::value(option));
}

char QModbusRtuSerialSlave_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

char QModbusRtuSerialSlave_WriteDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::writeData(*static_cast<QModbusDataUnit*>(newData));
}

void QModbusRtuSerialSlave_Close(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->close();
}

char QModbusRtuSerialSlave_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->open();
}

void QModbusRtuSerialSlave_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusRtuSerialSlave_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusRtuSerialSlave_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusRtuSerialSlave_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusRtuSerialSlave_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusRtuSerialSlave_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusRtuSerialSlave_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusRtuSerialSlave_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::customEvent(static_cast<QEvent*>(event));
}

void QModbusRtuSerialSlave_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusRtuSerialSlave*>(ptr), "deleteLater");
}

void QModbusRtuSerialSlave_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::deleteLater();
}

void QModbusRtuSerialSlave_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusRtuSerialSlave_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusRtuSerialSlave_Event(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusRtuSerialSlave_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::event(static_cast<QEvent*>(e));
}

char QModbusRtuSerialSlave_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusRtuSerialSlave_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusRtuSerialSlave_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusRtuSerialSlave*>(ptr)->metaObject());
}

void* QModbusRtuSerialSlave_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::metaObject());
}

class MyQModbusServer: public QModbusServer
{
public:
	MyQModbusServer(QObject *parent) : QModbusServer(parent) {};
	void Signal_DataWritten(QModbusDataUnit::RegisterType regist, int address, int size) { callbackQModbusServer_DataWritten(this, regist, address, size); };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, const_cast<QModbusPdu*>(&request))); };
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, const_cast<QModbusPdu*>(&request))); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<MyQModbusServer*>(this)) != 0; };
	bool readData(QModbusDataUnit * newData) const { return callbackQModbusServer_ReadData(const_cast<MyQModbusServer*>(this), newData) != 0; };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, option, const_cast<QVariant*>(&newValue)) != 0; };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<MyQModbusServer*>(this), option)); };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void close() { callbackQModbusServer_Close(this); };
	bool open() { return callbackQModbusServer_Open(this) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQModbusServer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQModbusServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusServer_CustomEvent(this, event); };
	void deleteLater() { callbackQModbusServer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusServer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusServer_MetaObject(const_cast<MyQModbusServer*>(this))); };
};

void* QModbusServer_NewQModbusServer(void* parent)
{
	return new MyQModbusServer(static_cast<QObject*>(parent));
}

char QModbusServer_Data(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_Data2(void* ptr, long long table, unsigned short address, unsigned short data)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit::RegisterType>(table), address, &data);
}

void QModbusServer_ConnectDataWritten(void* ptr)
{
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

void* QModbusServer_ProcessPrivateRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusServer*>(ptr)->processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusServer_ProcessPrivateRequestDefault(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusServer*>(ptr)->QModbusServer::processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusServer_ProcessRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusServer*>(ptr)->processRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusServer_ProcessRequestDefault(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusServer*>(ptr)->QModbusServer::processRequest(*static_cast<QModbusPdu*>(request)));
}

char QModbusServer_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->processesBroadcast();
}

char QModbusServer_ProcessesBroadcastDefault(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::processesBroadcast();
}

char QModbusServer_ReadData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->readData(static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_ReadDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::readData(static_cast<QModbusDataUnit*>(newData));
}

int QModbusServer_ServerAddress(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->serverAddress();
}

char QModbusServer_SetData2(void* ptr, long long table, unsigned short address, unsigned short data)
{
	return static_cast<QModbusServer*>(ptr)->setData(static_cast<QModbusDataUnit::RegisterType>(table), address, data);
}

char QModbusServer_SetData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->setData(*static_cast<QModbusDataUnit*>(newData));
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
	return static_cast<QModbusServer*>(ptr)->QModbusServer::setValue(option, *static_cast<QVariant*>(newValue));
}

void* QModbusServer_Value(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusServer*>(ptr)->value(option));
}

void* QModbusServer_ValueDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusServer*>(ptr)->QModbusServer::value(option));
}

char QModbusServer_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

char QModbusServer_WriteDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::writeData(*static_cast<QModbusDataUnit*>(newData));
}

void* QModbusServer___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusServer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusServer___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QModbusServer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusServer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusServer___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QModbusServer___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusServer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusServer___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusServer___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusServer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusServer___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusServer___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusServer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusServer___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QModbusServer_Close(void* ptr)
{
	static_cast<QModbusServer*>(ptr)->close();
}

char QModbusServer_Open(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->open();
}

void QModbusServer_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusServer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusServer*>(ptr)->QModbusServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusServer_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusServer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusServer*>(ptr)->QModbusServer::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusServer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusServer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusServer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusServer*>(ptr)->QModbusServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusServer_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusServer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusServer*>(ptr)->QModbusServer::customEvent(static_cast<QEvent*>(event));
}

void QModbusServer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusServer*>(ptr), "deleteLater");
}

void QModbusServer_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusServer*>(ptr)->QModbusServer::deleteLater();
}

void QModbusServer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusServer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusServer*>(ptr)->QModbusServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusServer_Event(void* ptr, void* e)
{
	return static_cast<QModbusServer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::event(static_cast<QEvent*>(e));
}

char QModbusServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusServer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusServer*>(ptr)->metaObject());
}

void* QModbusServer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusServer*>(ptr)->QModbusServer::metaObject());
}

void* QModbusTcpClient_NewQModbusTcpClient(void* parent)
{
	return new QModbusTcpClient(static_cast<QObject*>(parent));
}

void QModbusTcpClient_DestroyQModbusTcpClient(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->~QModbusTcpClient();
}

void* QModbusTcpClient___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusTcpClient___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpClient___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QModbusTcpClient___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusTcpClient___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusTcpClient___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QModbusTcpClient___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusTcpClient___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpClient___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusTcpClient___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusTcpClient___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpClient___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusTcpClient___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusTcpClient___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpClient___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

char QModbusTcpClient_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusTcpClient_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusTcpClient_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

char QModbusTcpClient_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

void QModbusTcpClient_Close(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->close();
}

char QModbusTcpClient_Open(void* ptr)
{
	return static_cast<QModbusTcpClient*>(ptr)->open();
}

void QModbusTcpClient_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusTcpClient*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusTcpClient_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusTcpClient_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusTcpClient*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusTcpClient_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusTcpClient_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusTcpClient*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusTcpClient_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusTcpClient_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusTcpClient*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusTcpClient_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::customEvent(static_cast<QEvent*>(event));
}

void QModbusTcpClient_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusTcpClient*>(ptr), "deleteLater");
}

void QModbusTcpClient_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::deleteLater();
}

void QModbusTcpClient_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusTcpClient*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusTcpClient_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusTcpClient_Event(void* ptr, void* e)
{
	return static_cast<QModbusTcpClient*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusTcpClient_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::event(static_cast<QEvent*>(e));
}

char QModbusTcpClient_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusTcpClient*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusTcpClient_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusTcpClient_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusTcpClient*>(ptr)->metaObject());
}

void* QModbusTcpClient_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::metaObject());
}

void* QModbusTcpServer_NewQModbusTcpServer(void* parent)
{
	return new QModbusTcpServer(static_cast<QObject*>(parent));
}

void QModbusTcpServer_DestroyQModbusTcpServer(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->~QModbusTcpServer();
}

void* QModbusTcpServer___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QModbusTcpServer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpServer___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QModbusTcpServer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QModbusTcpServer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QModbusTcpServer___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QModbusTcpServer___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusTcpServer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpServer___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusTcpServer___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusTcpServer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpServer___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusTcpServer___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QModbusTcpServer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QModbusTcpServer___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QModbusTcpServer_ProcessPrivateRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusTcpServer*>(ptr)->processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusTcpServer_ProcessPrivateRequestDefault(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processPrivateRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusTcpServer_ProcessRequest(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusTcpServer*>(ptr)->processRequest(*static_cast<QModbusPdu*>(request)));
}

void* QModbusTcpServer_ProcessRequestDefault(void* ptr, void* request)
{
	return new QModbusResponse(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processRequest(*static_cast<QModbusPdu*>(request)));
}

char QModbusTcpServer_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->processesBroadcast();
}

char QModbusTcpServer_ProcessesBroadcastDefault(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processesBroadcast();
}

char QModbusTcpServer_ReadData(void* ptr, void* newData)
{
	return static_cast<QModbusTcpServer*>(ptr)->readData(static_cast<QModbusDataUnit*>(newData));
}

char QModbusTcpServer_ReadDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::readData(static_cast<QModbusDataUnit*>(newData));
}

char QModbusTcpServer_SetValue(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusTcpServer*>(ptr)->setValue(option, *static_cast<QVariant*>(newValue));
}

char QModbusTcpServer_SetValueDefault(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::setValue(option, *static_cast<QVariant*>(newValue));
}

void* QModbusTcpServer_Value(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusTcpServer*>(ptr)->value(option));
}

void* QModbusTcpServer_ValueDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::value(option));
}

char QModbusTcpServer_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusTcpServer*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

char QModbusTcpServer_WriteDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::writeData(*static_cast<QModbusDataUnit*>(newData));
}

void QModbusTcpServer_Close(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->close();
}

char QModbusTcpServer_Open(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->open();
}

void QModbusTcpServer_TimerEvent(void* ptr, void* event)
{
	static_cast<QModbusTcpServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusTcpServer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QModbusTcpServer_ChildEvent(void* ptr, void* event)
{
	static_cast<QModbusTcpServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QModbusTcpServer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::childEvent(static_cast<QChildEvent*>(event));
}

void QModbusTcpServer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusTcpServer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusTcpServer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusTcpServer_CustomEvent(void* ptr, void* event)
{
	static_cast<QModbusTcpServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QModbusTcpServer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::customEvent(static_cast<QEvent*>(event));
}

void QModbusTcpServer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QModbusTcpServer*>(ptr), "deleteLater");
}

void QModbusTcpServer_DeleteLaterDefault(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::deleteLater();
}

void QModbusTcpServer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QModbusTcpServer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QModbusTcpServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QModbusTcpServer_Event(void* ptr, void* e)
{
	return static_cast<QModbusTcpServer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QModbusTcpServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::event(static_cast<QEvent*>(e));
}

char QModbusTcpServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusTcpServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QModbusTcpServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QModbusTcpServer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusTcpServer*>(ptr)->metaObject());
}

void* QModbusTcpServer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::metaObject());
}

