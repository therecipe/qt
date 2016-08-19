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

void* QCanBus_CreateDevice(void* ptr, char* plugin, char* interfaceName)
{
	return static_cast<QCanBus*>(ptr)->createDevice(QByteArray(plugin), QString(interfaceName));
}

void* QCanBus_QCanBus_Instance()
{
	return QCanBus::instance();
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

int QCanBus_Event(void* ptr, void* e)
{
	return static_cast<QCanBus*>(ptr)->event(static_cast<QEvent*>(e));
}

int QCanBus_EventDefault(void* ptr, void* e)
{
	return static_cast<QCanBus*>(ptr)->QCanBus::event(static_cast<QEvent*>(e));
}

int QCanBus_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCanBus*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QCanBus_EventFilterDefault(void* ptr, void* watched, void* event)
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
	void close() { callbackQCanBusDevice_Close(this, this->objectName().toUtf8().data()); };
	void Signal_ErrorOccurred(QCanBusDevice::CanBusError error) { callbackQCanBusDevice_ErrorOccurred(this, this->objectName().toUtf8().data(), error); };
	void Signal_FramesReceived() { callbackQCanBusDevice_FramesReceived(this, this->objectName().toUtf8().data()); };
	void Signal_FramesWritten(qint64 framesCount) { callbackQCanBusDevice_FramesWritten(this, this->objectName().toUtf8().data(), static_cast<long long>(framesCount)); };
	QString interpretErrorFrame(const QCanBusFrame & frame) { return QString(callbackQCanBusDevice_InterpretErrorFrame(this, this->objectName().toUtf8().data(), const_cast<QCanBusFrame*>(&frame))); };
	bool open() { return callbackQCanBusDevice_Open(this, this->objectName().toUtf8().data()) != 0; };
	void setConfigurationParameter(int key, const QVariant & value) { callbackQCanBusDevice_SetConfigurationParameter(this, this->objectName().toUtf8().data(), key, const_cast<QVariant*>(&value)); };
	void Signal_StateChanged(QCanBusDevice::CanBusDeviceState state) { callbackQCanBusDevice_StateChanged(this, this->objectName().toUtf8().data(), state); };
	bool writeFrame(const QCanBusFrame & frame) { return callbackQCanBusDevice_WriteFrame(this, this->objectName().toUtf8().data(), const_cast<QCanBusFrame*>(&frame)) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQCanBusDevice_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQCanBusDevice_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCanBusDevice_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQCanBusDevice_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCanBusDevice_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCanBusDevice_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCanBusDevice_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCanBusDevice_MetaObject(const_cast<MyQCanBusDevice*>(this), this->objectName().toUtf8().data())); };
};

long long QCanBusDevice_FramesAvailable(void* ptr)
{
	return static_cast<long long>(static_cast<QCanBusDevice*>(ptr)->framesAvailable());
}

long long QCanBusDevice_FramesToWrite(void* ptr)
{
	return static_cast<long long>(static_cast<QCanBusDevice*>(ptr)->framesToWrite());
}

void* QCanBusDevice_NewQCanBusDevice(void* parent)
{
	return new MyQCanBusDevice(static_cast<QObject*>(parent));
}

void QCanBusDevice_Close(void* ptr)
{
	static_cast<QCanBusDevice*>(ptr)->close();
}

void* QCanBusDevice_ConfigurationParameter(void* ptr, int key)
{
	return new QVariant(static_cast<QCanBusDevice*>(ptr)->configurationParameter(key));
}

int QCanBusDevice_ConnectDevice(void* ptr)
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

int QCanBusDevice_Error(void* ptr)
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

void QCanBusDevice_ErrorOccurred(void* ptr, int error)
{
	static_cast<QCanBusDevice*>(ptr)->errorOccurred(static_cast<QCanBusDevice::CanBusError>(error));
}

char* QCanBusDevice_ErrorString(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->errorString().toUtf8().data();
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
	static_cast<QCanBusDevice*>(ptr)->framesWritten(static_cast<long long>(framesCount));
}

int QCanBusDevice_HasOutgoingFrames(void* ptr)
{
	return static_cast<QCanBusDevice*>(ptr)->hasOutgoingFrames();
}

char* QCanBusDevice_InterpretErrorFrame(void* ptr, void* frame)
{
	return static_cast<QCanBusDevice*>(ptr)->interpretErrorFrame(*static_cast<QCanBusFrame*>(frame)).toUtf8().data();
}

int QCanBusDevice_Open(void* ptr)
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

void QCanBusDevice_SetError(void* ptr, char* errorText, int errorId)
{
	static_cast<QCanBusDevice*>(ptr)->setError(QString(errorText), static_cast<QCanBusDevice::CanBusError>(errorId));
}

void QCanBusDevice_SetState(void* ptr, int newState)
{
	static_cast<QCanBusDevice*>(ptr)->setState(static_cast<QCanBusDevice::CanBusDeviceState>(newState));
}

int QCanBusDevice_State(void* ptr)
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

void QCanBusDevice_StateChanged(void* ptr, int state)
{
	static_cast<QCanBusDevice*>(ptr)->stateChanged(static_cast<QCanBusDevice::CanBusDeviceState>(state));
}

int QCanBusDevice_WriteFrame(void* ptr, void* frame)
{
	return static_cast<QCanBusDevice*>(ptr)->writeFrame(*static_cast<QCanBusFrame*>(frame));
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

int QCanBusDevice_Event(void* ptr, void* e)
{
	return static_cast<QCanBusDevice*>(ptr)->event(static_cast<QEvent*>(e));
}

int QCanBusDevice_EventDefault(void* ptr, void* e)
{
	return static_cast<QCanBusDevice*>(ptr)->QCanBusDevice::event(static_cast<QEvent*>(e));
}

int QCanBusDevice_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCanBusDevice*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QCanBusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
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
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QCanBusDevice * createDevice(const QString & interfaceName) const { return static_cast<QCanBusDevice*>(callbackQCanBusFactory_CreateDevice(const_cast<MyQCanBusFactory*>(this), this->objectNameAbs().toUtf8().data(), interfaceName.toUtf8().data())); };
};

void* QCanBusFactory_CreateDevice(void* ptr, char* interfaceName)
{
	return static_cast<QCanBusFactory*>(ptr)->createDevice(QString(interfaceName));
}

char* QCanBusFactory_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQCanBusFactory*>(static_cast<QCanBusFactory*>(ptr))) {
		return static_cast<MyQCanBusFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QCanBusFactory_BASE").toUtf8().data();
}

void QCanBusFactory_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQCanBusFactory*>(static_cast<QCanBusFactory*>(ptr))) {
		static_cast<MyQCanBusFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
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

void* QCanBusFrame_NewQCanBusFrame(int ty)
{
	return new QCanBusFrame(static_cast<QCanBusFrame::FrameType>(ty));
}

int QCanBusFrame_Error(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->error();
}

int QCanBusFrame_FrameType(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->frameType();
}

int QCanBusFrame_HasExtendedFrameFormat(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->hasExtendedFrameFormat();
}

int QCanBusFrame_IsValid(void* ptr)
{
	return static_cast<QCanBusFrame*>(ptr)->isValid();
}

char* QCanBusFrame_Payload(void* ptr)
{
	return QString(static_cast<QCanBusFrame*>(ptr)->payload()).toUtf8().data();
}

void QCanBusFrame_SetError(void* ptr, int error)
{
	static_cast<QCanBusFrame*>(ptr)->setError(static_cast<QCanBusFrame::FrameError>(error));
}

void QCanBusFrame_SetExtendedFrameFormat(void* ptr, int isExtended)
{
	static_cast<QCanBusFrame*>(ptr)->setExtendedFrameFormat(isExtended != 0);
}

void QCanBusFrame_SetFrameType(void* ptr, int newType)
{
	static_cast<QCanBusFrame*>(ptr)->setFrameType(static_cast<QCanBusFrame::FrameType>(newType));
}

void QCanBusFrame_SetPayload(void* ptr, char* data)
{
	static_cast<QCanBusFrame*>(ptr)->setPayload(QByteArray(data));
}

class MyQModbusClient: public QModbusClient
{
public:
	MyQModbusClient(QObject *parent) : QModbusClient(parent) {};
	void Signal_TimeoutChanged(int newTimeout) { callbackQModbusClient_TimeoutChanged(this, this->objectName().toUtf8().data(), newTimeout); };
	bool processPrivateResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessPrivateResponse(this, this->objectName().toUtf8().data(), const_cast<QModbusResponse*>(&response), data) != 0; };
	bool processResponse(const QModbusResponse & response, QModbusDataUnit * data) { return callbackQModbusClient_ProcessResponse(this, this->objectName().toUtf8().data(), const_cast<QModbusResponse*>(&response), data) != 0; };
	void close() { callbackQModbusClient_Close(this, this->objectName().toUtf8().data()); };
	bool open() { return callbackQModbusClient_Open(this, this->objectName().toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQModbusClient_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQModbusClient_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusClient_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusClient_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQModbusClient_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusClient_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusClient_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusClient_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusClient_MetaObject(const_cast<MyQModbusClient*>(this), this->objectName().toUtf8().data())); };
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

int QModbusClient_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusClient_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusClient_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusClient_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

void QModbusClient_Close(void* ptr)
{
	static_cast<QModbusClient*>(ptr)->close();
}



int QModbusClient_Open(void* ptr)
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

int QModbusClient_Event(void* ptr, void* e)
{
	return static_cast<QModbusClient*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusClient_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusClient*>(ptr)->QModbusClient::event(static_cast<QEvent*>(e));
}

int QModbusClient_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusClient*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusClient_EventFilterDefault(void* ptr, void* watched, void* event)
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

void* QModbusDataUnit_NewQModbusDataUnit2(int ty)
{
	return new QModbusDataUnit(static_cast<QModbusDataUnit::RegisterType>(ty));
}

int QModbusDataUnit_IsValid(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->isValid();
}

int QModbusDataUnit_RegisterType(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->registerType();
}

void QModbusDataUnit_SetRegisterType(void* ptr, int ty)
{
	static_cast<QModbusDataUnit*>(ptr)->setRegisterType(static_cast<QModbusDataUnit::RegisterType>(ty));
}

void QModbusDataUnit_SetStartAddress(void* ptr, int address)
{
	static_cast<QModbusDataUnit*>(ptr)->setStartAddress(address);
}

int QModbusDataUnit_StartAddress(void* ptr)
{
	return static_cast<QModbusDataUnit*>(ptr)->startAddress();
}

class MyQModbusDevice: public QModbusDevice
{
public:
	MyQModbusDevice(QObject *parent) : QModbusDevice(parent) {};
	void close() { callbackQModbusDevice_Close(this, this->objectName().toUtf8().data()); };
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusDevice_ErrorOccurred(this, this->objectName().toUtf8().data(), error); };
	bool open() { return callbackQModbusDevice_Open(this, this->objectName().toUtf8().data()) != 0; };
	void Signal_StateChanged(QModbusDevice::State state) { callbackQModbusDevice_StateChanged(this, this->objectName().toUtf8().data(), state); };
	void timerEvent(QTimerEvent * event) { callbackQModbusDevice_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQModbusDevice_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusDevice_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusDevice_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQModbusDevice_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusDevice_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusDevice_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusDevice_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusDevice_MetaObject(const_cast<MyQModbusDevice*>(this), this->objectName().toUtf8().data())); };
};

void* QModbusDevice_NewQModbusDevice(void* parent)
{
	return new MyQModbusDevice(static_cast<QObject*>(parent));
}

void QModbusDevice_Close(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->close();
}

int QModbusDevice_ConnectDevice(void* ptr)
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

int QModbusDevice_Error(void* ptr)
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

void QModbusDevice_ErrorOccurred(void* ptr, int error)
{
	static_cast<QModbusDevice*>(ptr)->errorOccurred(static_cast<QModbusDevice::Error>(error));
}

char* QModbusDevice_ErrorString(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->errorString().toUtf8().data();
}

int QModbusDevice_Open(void* ptr)
{
	return static_cast<QModbusDevice*>(ptr)->open();
}

void QModbusDevice_SetConnectionParameter(void* ptr, int parameter, void* value)
{
	static_cast<QModbusDevice*>(ptr)->setConnectionParameter(parameter, *static_cast<QVariant*>(value));
}

void QModbusDevice_SetError(void* ptr, char* errorText, int error)
{
	static_cast<QModbusDevice*>(ptr)->setError(QString(errorText), static_cast<QModbusDevice::Error>(error));
}

void QModbusDevice_SetState(void* ptr, int newState)
{
	static_cast<QModbusDevice*>(ptr)->setState(static_cast<QModbusDevice::State>(newState));
}

int QModbusDevice_State(void* ptr)
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

void QModbusDevice_StateChanged(void* ptr, int state)
{
	static_cast<QModbusDevice*>(ptr)->stateChanged(static_cast<QModbusDevice::State>(state));
}

void QModbusDevice_DestroyQModbusDevice(void* ptr)
{
	static_cast<QModbusDevice*>(ptr)->~QModbusDevice();
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

int QModbusDevice_Event(void* ptr, void* e)
{
	return static_cast<QModbusDevice*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusDevice_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusDevice*>(ptr)->QModbusDevice::event(static_cast<QEvent*>(e));
}

int QModbusDevice_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusDevice*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusDevice_EventFilterDefault(void* ptr, void* watched, void* event)
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

int QModbusDeviceIdentification_ConformityLevel(void* ptr)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->conformityLevel();
}

int QModbusDeviceIdentification_IsValid(void* ptr)
{
	return static_cast<QModbusDeviceIdentification*>(ptr)->isValid();
}

void QModbusDeviceIdentification_SetConformityLevel(void* ptr, int level)
{
	static_cast<QModbusDeviceIdentification*>(ptr)->setConformityLevel(static_cast<QModbusDeviceIdentification::ConformityLevel>(level));
}

class MyQModbusExceptionResponse: public QModbusExceptionResponse
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQModbusExceptionResponse() : QModbusExceptionResponse() {};
	MyQModbusExceptionResponse(FunctionCode code, ExceptionCode ec) : QModbusExceptionResponse(code, ec) {};
	MyQModbusExceptionResponse(const QModbusPdu &pdu) : QModbusExceptionResponse(pdu) {};
	void setFunctionCode(QModbusPdu::FunctionCode c) { callbackQModbusExceptionResponse_SetFunctionCode(this, this->objectNameAbs().toUtf8().data(), c); };
};

void* QModbusExceptionResponse_NewQModbusExceptionResponse()
{
	return new MyQModbusExceptionResponse();
}

void* QModbusExceptionResponse_NewQModbusExceptionResponse3(int code, int ec)
{
	return new MyQModbusExceptionResponse(static_cast<QModbusPdu::FunctionCode>(code), static_cast<QModbusPdu::ExceptionCode>(ec));
}

void* QModbusExceptionResponse_NewQModbusExceptionResponse2(void* pdu)
{
	return new MyQModbusExceptionResponse(*static_cast<QModbusPdu*>(pdu));
}

void QModbusExceptionResponse_SetExceptionCode(void* ptr, int ec)
{
	static_cast<QModbusExceptionResponse*>(ptr)->setExceptionCode(static_cast<QModbusPdu::ExceptionCode>(ec));
}

void QModbusExceptionResponse_SetFunctionCode(void* ptr, int c)
{
	static_cast<QModbusExceptionResponse*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(c));
}

void QModbusExceptionResponse_SetFunctionCodeDefault(void* ptr, int c)
{
	static_cast<QModbusExceptionResponse*>(ptr)->QModbusExceptionResponse::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(c));
}

char* QModbusExceptionResponse_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQModbusExceptionResponse*>(static_cast<QModbusExceptionResponse*>(ptr))) {
		return static_cast<MyQModbusExceptionResponse*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QModbusExceptionResponse_BASE").toUtf8().data();
}

void QModbusExceptionResponse_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQModbusExceptionResponse*>(static_cast<QModbusExceptionResponse*>(ptr))) {
		static_cast<MyQModbusExceptionResponse*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQModbusPdu: public QModbusPdu
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQModbusPdu() : QModbusPdu() {};
	MyQModbusPdu(FunctionCode code, const QByteArray &data) : QModbusPdu(code, data) {};
	MyQModbusPdu(const QModbusPdu &other) : QModbusPdu(other) {};
	void setFunctionCode(QModbusPdu::FunctionCode code) { callbackQModbusPdu_SetFunctionCode(this, this->objectNameAbs().toUtf8().data(), code); };
};

void* QModbusPdu_NewQModbusPdu()
{
	return new MyQModbusPdu();
}

void* QModbusPdu_NewQModbusPdu2(int code, char* data)
{
	return new MyQModbusPdu(static_cast<QModbusPdu::FunctionCode>(code), QByteArray(data));
}

void* QModbusPdu_NewQModbusPdu3(void* other)
{
	return new MyQModbusPdu(*static_cast<QModbusPdu*>(other));
}

char* QModbusPdu_Data(void* ptr)
{
	return QString(static_cast<QModbusPdu*>(ptr)->data()).toUtf8().data();
}

int QModbusPdu_ExceptionCode(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->exceptionCode();
}

int QModbusPdu_FunctionCode(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->functionCode();
}

int QModbusPdu_IsException(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->isException();
}

int QModbusPdu_IsValid(void* ptr)
{
	return static_cast<QModbusPdu*>(ptr)->isValid();
}

void QModbusPdu_SetData(void* ptr, char* data)
{
	static_cast<QModbusPdu*>(ptr)->setData(QByteArray(data));
}

void QModbusPdu_SetFunctionCode(void* ptr, int code)
{
	static_cast<QModbusPdu*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusPdu_SetFunctionCodeDefault(void* ptr, int code)
{
	static_cast<QModbusPdu*>(ptr)->QModbusPdu::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusPdu_DestroyQModbusPdu(void* ptr)
{
	static_cast<QModbusPdu*>(ptr)->~QModbusPdu();
}

char* QModbusPdu_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQModbusPdu*>(static_cast<QModbusPdu*>(ptr))) {
		return static_cast<MyQModbusPdu*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QModbusPdu_BASE").toUtf8().data();
}

void QModbusPdu_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQModbusPdu*>(static_cast<QModbusPdu*>(ptr))) {
		static_cast<MyQModbusPdu*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQModbusReply: public QModbusReply
{
public:
	MyQModbusReply(ReplyType type, int serverAddress, QObject *parent) : QModbusReply(type, serverAddress, parent) {};
	void Signal_ErrorOccurred(QModbusDevice::Error error) { callbackQModbusReply_ErrorOccurred(this, this->objectName().toUtf8().data(), error); };
	void Signal_Finished() { callbackQModbusReply_Finished(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQModbusReply_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQModbusReply_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusReply_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusReply_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQModbusReply_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusReply_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusReply_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusReply_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusReply_MetaObject(const_cast<MyQModbusReply*>(this), this->objectName().toUtf8().data())); };
};

void* QModbusReply_NewQModbusReply(int ty, int serverAddress, void* parent)
{
	return new MyQModbusReply(static_cast<QModbusReply::ReplyType>(ty), serverAddress, static_cast<QObject*>(parent));
}

int QModbusReply_Error(void* ptr)
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

void QModbusReply_ErrorOccurred(void* ptr, int error)
{
	static_cast<QModbusReply*>(ptr)->errorOccurred(static_cast<QModbusDevice::Error>(error));
}

char* QModbusReply_ErrorString(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->errorString().toUtf8().data();
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

int QModbusReply_IsFinished(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->isFinished();
}

int QModbusReply_ServerAddress(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->serverAddress();
}

int QModbusReply_Type(void* ptr)
{
	return static_cast<QModbusReply*>(ptr)->type();
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

int QModbusReply_Event(void* ptr, void* e)
{
	return static_cast<QModbusReply*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusReply_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusReply*>(ptr)->QModbusReply::event(static_cast<QEvent*>(e));
}

int QModbusReply_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusReply*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusReply_EventFilterDefault(void* ptr, void* watched, void* event)
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

void* QModbusRequest_NewQModbusRequest3(int code, char* data)
{
	return new QModbusRequest(static_cast<QModbusPdu::FunctionCode>(code), QByteArray(data));
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

void QModbusRequest_SetFunctionCode(void* ptr, int code)
{
	static_cast<QModbusRequest*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusRequest_SetFunctionCodeDefault(void* ptr, int code)
{
	static_cast<QModbusRequest*>(ptr)->QModbusRequest::setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void* QModbusResponse_NewQModbusResponse()
{
	return new QModbusResponse();
}

void* QModbusResponse_NewQModbusResponse3(int code, char* data)
{
	return new QModbusResponse(static_cast<QModbusPdu::FunctionCode>(code), QByteArray(data));
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

void QModbusResponse_SetFunctionCode(void* ptr, int code)
{
	static_cast<QModbusResponse*>(ptr)->setFunctionCode(static_cast<QModbusPdu::FunctionCode>(code));
}

void QModbusResponse_SetFunctionCodeDefault(void* ptr, int code)
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

int QModbusRtuSerialMaster_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusRtuSerialMaster_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusRtuSerialMaster_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusRtuSerialMaster_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

void QModbusRtuSerialMaster_Close(void* ptr)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->close();
}

void QModbusRtuSerialMaster_CloseDefault(void* ptr)
{
	static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::close();
}

int QModbusRtuSerialMaster_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->open();
}

int QModbusRtuSerialMaster_OpenDefault(void* ptr)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::open();
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

int QModbusRtuSerialMaster_Event(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusRtuSerialMaster_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->QModbusRtuSerialMaster::event(static_cast<QEvent*>(e));
}

int QModbusRtuSerialMaster_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusRtuSerialMaster*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusRtuSerialMaster_EventFilterDefault(void* ptr, void* watched, void* event)
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









int QModbusRtuSerialSlave_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->processesBroadcast();
}

int QModbusRtuSerialSlave_ProcessesBroadcastDefault(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::processesBroadcast();
}

int QModbusRtuSerialSlave_SetValue(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->setValue(option, *static_cast<QVariant*>(newValue));
}

int QModbusRtuSerialSlave_SetValueDefault(void* ptr, int option, void* newValue)
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

int QModbusRtuSerialSlave_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

int QModbusRtuSerialSlave_WriteDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::writeData(*static_cast<QModbusDataUnit*>(newData));
}

void QModbusRtuSerialSlave_Close(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->close();
}

void QModbusRtuSerialSlave_CloseDefault(void* ptr)
{
	static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::close();
}

int QModbusRtuSerialSlave_Open(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->open();
}

int QModbusRtuSerialSlave_OpenDefault(void* ptr)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::open();
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

int QModbusRtuSerialSlave_Event(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusRtuSerialSlave_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->QModbusRtuSerialSlave::event(static_cast<QEvent*>(e));
}

int QModbusRtuSerialSlave_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusRtuSerialSlave*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusRtuSerialSlave_EventFilterDefault(void* ptr, void* watched, void* event)
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
	void Signal_DataWritten(QModbusDataUnit::RegisterType regist, int address, int size) { callbackQModbusServer_DataWritten(this, this->objectName().toUtf8().data(), regist, address, size); };
	QModbusResponse processPrivateRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessPrivateRequest(this, this->objectName().toUtf8().data(), const_cast<QModbusPdu*>(&request))); };
	QModbusResponse processRequest(const QModbusPdu & request) { return *static_cast<QModbusResponse*>(callbackQModbusServer_ProcessRequest(this, this->objectName().toUtf8().data(), const_cast<QModbusPdu*>(&request))); };
	bool processesBroadcast() const { return callbackQModbusServer_ProcessesBroadcast(const_cast<MyQModbusServer*>(this), this->objectName().toUtf8().data()) != 0; };
	bool setValue(int option, const QVariant & newValue) { return callbackQModbusServer_SetValue(this, this->objectName().toUtf8().data(), option, const_cast<QVariant*>(&newValue)) != 0; };
	QVariant value(int option) const { return *static_cast<QVariant*>(callbackQModbusServer_Value(const_cast<MyQModbusServer*>(this), this->objectName().toUtf8().data(), option)); };
	bool writeData(const QModbusDataUnit & newData) { return callbackQModbusServer_WriteData(this, this->objectName().toUtf8().data(), const_cast<QModbusDataUnit*>(&newData)) != 0; };
	void close() { callbackQModbusServer_Close(this, this->objectName().toUtf8().data()); };
	bool open() { return callbackQModbusServer_Open(this, this->objectName().toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQModbusServer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQModbusServer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQModbusServer_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQModbusServer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQModbusServer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQModbusServer_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQModbusServer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQModbusServer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQModbusServer_MetaObject(const_cast<MyQModbusServer*>(this), this->objectName().toUtf8().data())); };
};

void* QModbusServer_NewQModbusServer(void* parent)
{
	return new MyQModbusServer(static_cast<QObject*>(parent));
}

int QModbusServer_Data(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->data(static_cast<QModbusDataUnit*>(newData));
}

void QModbusServer_ConnectDataWritten(void* ptr)
{
	QObject::connect(static_cast<QModbusServer*>(ptr), static_cast<void (QModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&QModbusServer::dataWritten), static_cast<MyQModbusServer*>(ptr), static_cast<void (MyQModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&MyQModbusServer::Signal_DataWritten));
}

void QModbusServer_DisconnectDataWritten(void* ptr)
{
	QObject::disconnect(static_cast<QModbusServer*>(ptr), static_cast<void (QModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&QModbusServer::dataWritten), static_cast<MyQModbusServer*>(ptr), static_cast<void (MyQModbusServer::*)(QModbusDataUnit::RegisterType, int, int)>(&MyQModbusServer::Signal_DataWritten));
}

void QModbusServer_DataWritten(void* ptr, int regist, int address, int size)
{
	static_cast<QModbusServer*>(ptr)->dataWritten(static_cast<QModbusDataUnit::RegisterType>(regist), address, size);
}









int QModbusServer_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->processesBroadcast();
}

int QModbusServer_ProcessesBroadcastDefault(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::processesBroadcast();
}

int QModbusServer_ServerAddress(void* ptr)
{
	return static_cast<QModbusServer*>(ptr)->serverAddress();
}

int QModbusServer_SetData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->setData(*static_cast<QModbusDataUnit*>(newData));
}

void QModbusServer_SetServerAddress(void* ptr, int serverAddress)
{
	static_cast<QModbusServer*>(ptr)->setServerAddress(serverAddress);
}

int QModbusServer_SetValue(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusServer*>(ptr)->setValue(option, *static_cast<QVariant*>(newValue));
}

int QModbusServer_SetValueDefault(void* ptr, int option, void* newValue)
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

int QModbusServer_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

int QModbusServer_WriteDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::writeData(*static_cast<QModbusDataUnit*>(newData));
}

void QModbusServer_Close(void* ptr)
{
	static_cast<QModbusServer*>(ptr)->close();
}



int QModbusServer_Open(void* ptr)
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

int QModbusServer_Event(void* ptr, void* e)
{
	return static_cast<QModbusServer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusServer*>(ptr)->QModbusServer::event(static_cast<QEvent*>(e));
}

int QModbusServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusServer_EventFilterDefault(void* ptr, void* watched, void* event)
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

int QModbusTcpClient_ProcessPrivateResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusTcpClient_ProcessPrivateResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::processPrivateResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusTcpClient_ProcessResponse(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

int QModbusTcpClient_ProcessResponseDefault(void* ptr, void* response, void* data)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::processResponse(*static_cast<QModbusResponse*>(response), static_cast<QModbusDataUnit*>(data));
}

void QModbusTcpClient_Close(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->close();
}

void QModbusTcpClient_CloseDefault(void* ptr)
{
	static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::close();
}

int QModbusTcpClient_Open(void* ptr)
{
	return static_cast<QModbusTcpClient*>(ptr)->open();
}

int QModbusTcpClient_OpenDefault(void* ptr)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::open();
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

int QModbusTcpClient_Event(void* ptr, void* e)
{
	return static_cast<QModbusTcpClient*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusTcpClient_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusTcpClient*>(ptr)->QModbusTcpClient::event(static_cast<QEvent*>(e));
}

int QModbusTcpClient_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusTcpClient*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusTcpClient_EventFilterDefault(void* ptr, void* watched, void* event)
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









int QModbusTcpServer_ProcessesBroadcast(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->processesBroadcast();
}

int QModbusTcpServer_ProcessesBroadcastDefault(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::processesBroadcast();
}

int QModbusTcpServer_SetValue(void* ptr, int option, void* newValue)
{
	return static_cast<QModbusTcpServer*>(ptr)->setValue(option, *static_cast<QVariant*>(newValue));
}

int QModbusTcpServer_SetValueDefault(void* ptr, int option, void* newValue)
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

int QModbusTcpServer_WriteData(void* ptr, void* newData)
{
	return static_cast<QModbusTcpServer*>(ptr)->writeData(*static_cast<QModbusDataUnit*>(newData));
}

int QModbusTcpServer_WriteDataDefault(void* ptr, void* newData)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::writeData(*static_cast<QModbusDataUnit*>(newData));
}

void QModbusTcpServer_Close(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->close();
}

void QModbusTcpServer_CloseDefault(void* ptr)
{
	static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::close();
}

int QModbusTcpServer_Open(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->open();
}

int QModbusTcpServer_OpenDefault(void* ptr)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::open();
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

int QModbusTcpServer_Event(void* ptr, void* e)
{
	return static_cast<QModbusTcpServer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QModbusTcpServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QModbusTcpServer*>(ptr)->QModbusTcpServer::event(static_cast<QEvent*>(e));
}

int QModbusTcpServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QModbusTcpServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QModbusTcpServer_EventFilterDefault(void* ptr, void* watched, void* event)
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

