#define protected public

#include "bluetooth.h"
#include "_cgo_export.h"

#include <QBluetoothAddress>
#include <QBluetoothDeviceDiscoveryAgent>
#include <QBluetoothDeviceInfo>
#include <QBluetoothHostInfo>
#include <QBluetoothLocalDevice>
#include <QBluetoothServer>
#include <QBluetoothServiceDiscoveryAgent>
#include <QBluetoothServiceInfo>
#include <QBluetoothSocket>
#include <QBluetoothTransferManager>
#include <QBluetoothTransferReply>
#include <QBluetoothTransferRequest>
#include <QBluetoothUuid>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QIODevice>
#include <QLowEnergyCharacteristic>
#include <QLowEnergyController>
#include <QLowEnergyDescriptor>
#include <QLowEnergyService>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUuid>
#include <QVariant>

void* QBluetoothAddress_NewQBluetoothAddress(){
	return new QBluetoothAddress();
}

void* QBluetoothAddress_NewQBluetoothAddress4(void* other){
	return new QBluetoothAddress(*static_cast<QBluetoothAddress*>(other));
}

void* QBluetoothAddress_NewQBluetoothAddress3(char* address){
	return new QBluetoothAddress(QString(address));
}

void QBluetoothAddress_Clear(void* ptr){
	static_cast<QBluetoothAddress*>(ptr)->clear();
}

int QBluetoothAddress_IsNull(void* ptr){
	return static_cast<QBluetoothAddress*>(ptr)->isNull();
}

char* QBluetoothAddress_ToString(void* ptr){
	return static_cast<QBluetoothAddress*>(ptr)->toString().toUtf8().data();
}

void QBluetoothAddress_DestroyQBluetoothAddress(void* ptr){
	static_cast<QBluetoothAddress*>(ptr)->~QBluetoothAddress();
}

class MyQBluetoothDeviceDiscoveryAgent: public QBluetoothDeviceDiscoveryAgent {
public:
	void Signal_Canceled() { callbackQBluetoothDeviceDiscoveryAgentCanceled(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QBluetoothDeviceDiscoveryAgent::Error error) { callbackQBluetoothDeviceDiscoveryAgentError2(this, this->objectName().toUtf8().data(), error); };
	void Signal_Finished() { callbackQBluetoothDeviceDiscoveryAgentFinished(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothDeviceDiscoveryAgentTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothDeviceDiscoveryAgentChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothDeviceDiscoveryAgentCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_Canceled(void* ptr){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->canceled();
}

void QBluetoothDeviceDiscoveryAgent_ConnectError2(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&QBluetoothDeviceDiscoveryAgent::error), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Error2));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&QBluetoothDeviceDiscoveryAgent::error), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Error2));;
}

void QBluetoothDeviceDiscoveryAgent_Error2(void* ptr, int error){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->error(static_cast<QBluetoothDeviceDiscoveryAgent::Error>(error));
}

void QBluetoothDeviceDiscoveryAgent_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothDeviceDiscoveryAgent_Finished(void* ptr){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->finished();
}

void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(void* parent){
	return new QBluetoothDeviceDiscoveryAgent(static_cast<QObject*>(parent));
}

void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(void* deviceAdapter, void* parent){
	return new QBluetoothDeviceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

int QBluetoothDeviceDiscoveryAgent_Error(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->error();
}

char* QBluetoothDeviceDiscoveryAgent_ErrorString(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothDeviceDiscoveryAgent_InquiryType(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->inquiryType();
}

int QBluetoothDeviceDiscoveryAgent_IsActive(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->isActive();
}

void QBluetoothDeviceDiscoveryAgent_SetInquiryType(void* ptr, int ty){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->setInquiryType(static_cast<QBluetoothDeviceDiscoveryAgent::InquiryType>(ty));
}

void QBluetoothDeviceDiscoveryAgent_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "start");
}

void QBluetoothDeviceDiscoveryAgent_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(void* ptr){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->~QBluetoothDeviceDiscoveryAgent();
}

void QBluetoothDeviceDiscoveryAgent_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::customEvent(static_cast<QEvent*>(event));
}

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo(){
	return new QBluetoothDeviceInfo();
}

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(void* other){
	return new QBluetoothDeviceInfo(*static_cast<QBluetoothDeviceInfo*>(other));
}

int QBluetoothDeviceInfo_CoreConfigurations(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->coreConfigurations();
}

int QBluetoothDeviceInfo_IsCached(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->isCached();
}

int QBluetoothDeviceInfo_IsValid(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->isValid();
}

int QBluetoothDeviceInfo_MajorDeviceClass(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->majorDeviceClass();
}

char* QBluetoothDeviceInfo_Name(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->name().toUtf8().data();
}

int QBluetoothDeviceInfo_ServiceClasses(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->serviceClasses();
}

int QBluetoothDeviceInfo_ServiceUuidsCompleteness(void* ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->serviceUuidsCompleteness();
}

void QBluetoothDeviceInfo_SetCached(void* ptr, int cached){
	static_cast<QBluetoothDeviceInfo*>(ptr)->setCached(cached != 0);
}

void QBluetoothDeviceInfo_SetCoreConfigurations(void* ptr, int coreConfigs){
	static_cast<QBluetoothDeviceInfo*>(ptr)->setCoreConfigurations(static_cast<QBluetoothDeviceInfo::CoreConfiguration>(coreConfigs));
}

void QBluetoothDeviceInfo_SetDeviceUuid(void* ptr, void* uuid){
	static_cast<QBluetoothDeviceInfo*>(ptr)->setDeviceUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(void* ptr){
	static_cast<QBluetoothDeviceInfo*>(ptr)->~QBluetoothDeviceInfo();
}

void* QBluetoothHostInfo_NewQBluetoothHostInfo(){
	return new QBluetoothHostInfo();
}

void* QBluetoothHostInfo_NewQBluetoothHostInfo2(void* other){
	return new QBluetoothHostInfo(*static_cast<QBluetoothHostInfo*>(other));
}

char* QBluetoothHostInfo_Name(void* ptr){
	return static_cast<QBluetoothHostInfo*>(ptr)->name().toUtf8().data();
}

void QBluetoothHostInfo_SetAddress(void* ptr, void* address){
	static_cast<QBluetoothHostInfo*>(ptr)->setAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothHostInfo_SetName(void* ptr, char* name){
	static_cast<QBluetoothHostInfo*>(ptr)->setName(QString(name));
}

void QBluetoothHostInfo_DestroyQBluetoothHostInfo(void* ptr){
	static_cast<QBluetoothHostInfo*>(ptr)->~QBluetoothHostInfo();
}

class MyQBluetoothLocalDevice: public QBluetoothLocalDevice {
public:
	MyQBluetoothLocalDevice(QObject *parent) : QBluetoothLocalDevice(parent) {};
	MyQBluetoothLocalDevice(const QBluetoothAddress &address, QObject *parent) : QBluetoothLocalDevice(address, parent) {};
	void Signal_Error(QBluetoothLocalDevice::Error error) { callbackQBluetoothLocalDeviceError(this, this->objectName().toUtf8().data(), error); };
	void Signal_HostModeStateChanged(QBluetoothLocalDevice::HostMode state) { callbackQBluetoothLocalDeviceHostModeStateChanged(this, this->objectName().toUtf8().data(), state); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothLocalDeviceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothLocalDeviceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothLocalDeviceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QBluetoothLocalDevice_ConnectError(void* ptr){
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));;
}

void QBluetoothLocalDevice_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));;
}

void QBluetoothLocalDevice_Error(void* ptr, int error){
	static_cast<QBluetoothLocalDevice*>(ptr)->error(static_cast<QBluetoothLocalDevice::Error>(error));
}

void QBluetoothLocalDevice_ConnectHostModeStateChanged(void* ptr){
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));;
}

void QBluetoothLocalDevice_DisconnectHostModeStateChanged(void* ptr){
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));;
}

void QBluetoothLocalDevice_HostModeStateChanged(void* ptr, int state){
	static_cast<QBluetoothLocalDevice*>(ptr)->hostModeStateChanged(static_cast<QBluetoothLocalDevice::HostMode>(state));
}

int QBluetoothLocalDevice_IsValid(void* ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->isValid();
}

void QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(void* ptr){
	static_cast<QBluetoothLocalDevice*>(ptr)->~QBluetoothLocalDevice();
}

void* QBluetoothLocalDevice_NewQBluetoothLocalDevice(void* parent){
	return new MyQBluetoothLocalDevice(static_cast<QObject*>(parent));
}

void* QBluetoothLocalDevice_NewQBluetoothLocalDevice2(void* address, void* parent){
	return new MyQBluetoothLocalDevice(*static_cast<QBluetoothAddress*>(address), static_cast<QObject*>(parent));
}

int QBluetoothLocalDevice_HostMode(void* ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->hostMode();
}

char* QBluetoothLocalDevice_Name(void* ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->name().toUtf8().data();
}

void QBluetoothLocalDevice_PairingConfirmation(void* ptr, int accept){
	QMetaObject::invokeMethod(static_cast<QBluetoothLocalDevice*>(ptr), "pairingConfirmation", Q_ARG(bool, accept != 0));
}

int QBluetoothLocalDevice_PairingStatus(void* ptr, void* address){
	return static_cast<QBluetoothLocalDevice*>(ptr)->pairingStatus(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothLocalDevice_PowerOn(void* ptr){
	static_cast<QBluetoothLocalDevice*>(ptr)->powerOn();
}

void QBluetoothLocalDevice_RequestPairing(void* ptr, void* address, int pairing){
	static_cast<QBluetoothLocalDevice*>(ptr)->requestPairing(*static_cast<QBluetoothAddress*>(address), static_cast<QBluetoothLocalDevice::Pairing>(pairing));
}

void QBluetoothLocalDevice_SetHostMode(void* ptr, int mode){
	static_cast<QBluetoothLocalDevice*>(ptr)->setHostMode(static_cast<QBluetoothLocalDevice::HostMode>(mode));
}

void QBluetoothLocalDevice_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothLocalDevice*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothLocalDevice_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothLocalDevice_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothLocalDevice*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothLocalDevice_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothLocalDevice_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothLocalDevice*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothLocalDevice_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::customEvent(static_cast<QEvent*>(event));
}

class MyQBluetoothServer: public QBluetoothServer {
public:
	void Signal_Error2(QBluetoothServer::Error error) { callbackQBluetoothServerError2(this, this->objectName().toUtf8().data(), error); };
	void Signal_NewConnection() { callbackQBluetoothServerNewConnection(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothServerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothServerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothServerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QBluetoothServer_NewQBluetoothServer(int serverType, void* parent){
	return new QBluetoothServer(static_cast<QBluetoothServiceInfo::Protocol>(serverType), static_cast<QObject*>(parent));
}

void QBluetoothServer_ConnectError2(void* ptr){
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)(QBluetoothServer::Error)>(&QBluetoothServer::error), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)(QBluetoothServer::Error)>(&MyQBluetoothServer::Signal_Error2));;
}

void QBluetoothServer_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)(QBluetoothServer::Error)>(&QBluetoothServer::error), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)(QBluetoothServer::Error)>(&MyQBluetoothServer::Signal_Error2));;
}

void QBluetoothServer_Error2(void* ptr, int error){
	static_cast<QBluetoothServer*>(ptr)->error(static_cast<QBluetoothServer::Error>(error));
}

void QBluetoothServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

void QBluetoothServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

void QBluetoothServer_NewConnection(void* ptr){
	static_cast<QBluetoothServer*>(ptr)->newConnection();
}

int QBluetoothServer_Error(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->error();
}

int QBluetoothServer_IsListening(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->isListening();
}

int QBluetoothServer_MaxPendingConnections(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->maxPendingConnections();
}

int QBluetoothServer_ServerType(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->serverType();
}

void QBluetoothServer_DestroyQBluetoothServer(void* ptr){
	static_cast<QBluetoothServer*>(ptr)->~QBluetoothServer();
}

void QBluetoothServer_Close(void* ptr){
	static_cast<QBluetoothServer*>(ptr)->close();
}

int QBluetoothServer_HasPendingConnections(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->hasPendingConnections();
}

void* QBluetoothServer_NextPendingConnection(void* ptr){
	return static_cast<QBluetoothServer*>(ptr)->nextPendingConnection();
}

void QBluetoothServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QBluetoothServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QBluetoothServer_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServer_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServer_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServer_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServer_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothServer_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::customEvent(static_cast<QEvent*>(event));
}

class MyQBluetoothServiceDiscoveryAgent: public QBluetoothServiceDiscoveryAgent {
public:
	void Signal_Canceled() { callbackQBluetoothServiceDiscoveryAgentCanceled(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QBluetoothServiceDiscoveryAgent::Error error) { callbackQBluetoothServiceDiscoveryAgentError2(this, this->objectName().toUtf8().data(), error); };
	void Signal_Finished() { callbackQBluetoothServiceDiscoveryAgentFinished(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothServiceDiscoveryAgentTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothServiceDiscoveryAgentChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothServiceDiscoveryAgentCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_Canceled(void* ptr){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->canceled();
}

void QBluetoothServiceDiscoveryAgent_ConnectError2(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&QBluetoothServiceDiscoveryAgent::error), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&MyQBluetoothServiceDiscoveryAgent::Signal_Error2));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&QBluetoothServiceDiscoveryAgent::error), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&MyQBluetoothServiceDiscoveryAgent::Signal_Error2));;
}

void QBluetoothServiceDiscoveryAgent_Error2(void* ptr, int error){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->error(static_cast<QBluetoothServiceDiscoveryAgent::Error>(error));
}

void QBluetoothServiceDiscoveryAgent_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothServiceDiscoveryAgent_Finished(void* ptr){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->finished();
}

void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(void* parent){
	return new QBluetoothServiceDiscoveryAgent(static_cast<QObject*>(parent));
}

void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(void* deviceAdapter, void* parent){
	return new QBluetoothServiceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

void QBluetoothServiceDiscoveryAgent_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "clear");
}

int QBluetoothServiceDiscoveryAgent_Error(void* ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->error();
}

char* QBluetoothServiceDiscoveryAgent_ErrorString(void* ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothServiceDiscoveryAgent_IsActive(void* ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->isActive();
}

int QBluetoothServiceDiscoveryAgent_SetRemoteAddress(void* ptr, void* address){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setRemoteAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothServiceDiscoveryAgent_SetUuidFilter2(void* ptr, void* uuid){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setUuidFilter(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothServiceDiscoveryAgent_Start(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "start", Q_ARG(QBluetoothServiceDiscoveryAgent::DiscoveryMode, static_cast<QBluetoothServiceDiscoveryAgent::DiscoveryMode>(mode)));
}

void QBluetoothServiceDiscoveryAgent_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(void* ptr){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->~QBluetoothServiceDiscoveryAgent();
}

void QBluetoothServiceDiscoveryAgent_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::customEvent(static_cast<QEvent*>(event));
}

int QBluetoothServiceInfo_ServiceName_Type(){
	return QBluetoothServiceInfo::ServiceName;
}

int QBluetoothServiceInfo_ServiceDescription_Type(){
	return QBluetoothServiceInfo::ServiceDescription;
}

int QBluetoothServiceInfo_ServiceProvider_Type(){
	return QBluetoothServiceInfo::ServiceProvider;
}

char* QBluetoothServiceInfo_ServiceDescription(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serviceDescription().toUtf8().data();
}

char* QBluetoothServiceInfo_ServiceName(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serviceName().toUtf8().data();
}

char* QBluetoothServiceInfo_ServiceProvider(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serviceProvider().toUtf8().data();
}

void QBluetoothServiceInfo_SetServiceDescription(void* ptr, char* description){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceDescription(QString(description));
}

void QBluetoothServiceInfo_SetServiceName(void* ptr, char* name){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceName(QString(name));
}

void QBluetoothServiceInfo_SetServiceProvider(void* ptr, char* provider){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceProvider(QString(provider));
}

void QBluetoothServiceInfo_SetServiceUuid(void* ptr, void* uuid){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void* QBluetoothServiceInfo_NewQBluetoothServiceInfo(){
	return new QBluetoothServiceInfo();
}

void* QBluetoothServiceInfo_NewQBluetoothServiceInfo2(void* other){
	return new QBluetoothServiceInfo(*static_cast<QBluetoothServiceInfo*>(other));
}

int QBluetoothServiceInfo_IsComplete(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->isComplete();
}

int QBluetoothServiceInfo_IsRegistered(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->isRegistered();
}

int QBluetoothServiceInfo_IsValid(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->isValid();
}

int QBluetoothServiceInfo_ProtocolServiceMultiplexer(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->protocolServiceMultiplexer();
}

int QBluetoothServiceInfo_RegisterService(void* ptr, void* localAdapter){
	return static_cast<QBluetoothServiceInfo*>(ptr)->registerService(*static_cast<QBluetoothAddress*>(localAdapter));
}

int QBluetoothServiceInfo_ServerChannel(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serverChannel();
}

void QBluetoothServiceInfo_SetDevice(void* ptr, void* device){
	static_cast<QBluetoothServiceInfo*>(ptr)->setDevice(*static_cast<QBluetoothDeviceInfo*>(device));
}

int QBluetoothServiceInfo_SocketProtocol(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->socketProtocol();
}

int QBluetoothServiceInfo_UnregisterService(void* ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->unregisterService();
}

void QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(void* ptr){
	static_cast<QBluetoothServiceInfo*>(ptr)->~QBluetoothServiceInfo();
}

class MyQBluetoothSocket: public QBluetoothSocket {
public:
	MyQBluetoothSocket(QBluetoothServiceInfo::Protocol socketType, QObject *parent) : QBluetoothSocket(socketType, parent) {};
	MyQBluetoothSocket(QObject *parent) : QBluetoothSocket(parent) {};
	void Signal_Connected() { callbackQBluetoothSocketConnected(this, this->objectName().toUtf8().data()); };
	void Signal_Disconnected() { callbackQBluetoothSocketDisconnected(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QBluetoothSocket::SocketError error) { callbackQBluetoothSocketError2(this, this->objectName().toUtf8().data(), error); };
	void Signal_StateChanged(QBluetoothSocket::SocketState state) { callbackQBluetoothSocketStateChanged(this, this->objectName().toUtf8().data(), state); };
	void close() { callbackQBluetoothSocketClose(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothSocketTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothSocketChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothSocketCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QBluetoothSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_Connected(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->connected();
}

void QBluetoothSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_Disconnected(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->disconnected();
}

void QBluetoothSocket_ConnectError2(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&QBluetoothSocket::error), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&MyQBluetoothSocket::Signal_Error2));;
}

void QBluetoothSocket_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&QBluetoothSocket::error), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&MyQBluetoothSocket::Signal_Error2));;
}

void QBluetoothSocket_Error2(void* ptr, int error){
	static_cast<QBluetoothSocket*>(ptr)->error(static_cast<QBluetoothSocket::SocketError>(error));
}

void QBluetoothSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void QBluetoothSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void QBluetoothSocket_StateChanged(void* ptr, int state){
	static_cast<QBluetoothSocket*>(ptr)->stateChanged(static_cast<QBluetoothSocket::SocketState>(state));
}

void* QBluetoothSocket_NewQBluetoothSocket(int socketType, void* parent){
	return new MyQBluetoothSocket(static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QObject*>(parent));
}

void* QBluetoothSocket_NewQBluetoothSocket2(void* parent){
	return new MyQBluetoothSocket(static_cast<QObject*>(parent));
}

void QBluetoothSocket_Abort(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->abort();
}

long long QBluetoothSocket_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QBluetoothSocket*>(ptr)->bytesAvailable());
}

long long QBluetoothSocket_BytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QBluetoothSocket*>(ptr)->bytesToWrite());
}

int QBluetoothSocket_CanReadLine(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->canReadLine();
}

void QBluetoothSocket_Close(void* ptr){
	static_cast<MyQBluetoothSocket*>(ptr)->close();
}

void QBluetoothSocket_CloseDefault(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::close();
}

void QBluetoothSocket_ConnectToService2(void* ptr, void* address, void* uuid, int openMode){
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothAddress*>(address), *static_cast<QBluetoothUuid*>(uuid), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_ConnectToService(void* ptr, void* service, int openMode){
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothServiceInfo*>(service), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_DisconnectFromService(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->disconnectFromService();
}

int QBluetoothSocket_Error(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->error();
}

char* QBluetoothSocket_ErrorString(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothSocket_IsSequential(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->isSequential();
}

char* QBluetoothSocket_LocalName(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->localName().toUtf8().data();
}

char* QBluetoothSocket_PeerName(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->peerName().toUtf8().data();
}

long long QBluetoothSocket_ReadData(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QBluetoothSocket*>(ptr)->readData(data, static_cast<long long>(maxSize)));
}

int QBluetoothSocket_SetSocketDescriptor(void* ptr, int socketDescriptor, int socketType, int socketState, int openMode){
	return static_cast<QBluetoothSocket*>(ptr)->setSocketDescriptor(socketDescriptor, static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QBluetoothSocket::SocketState>(socketState), static_cast<QIODevice::OpenModeFlag>(openMode));
}

int QBluetoothSocket_SocketDescriptor(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketDescriptor();
}

int QBluetoothSocket_SocketType(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketType();
}

int QBluetoothSocket_State(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->state();
}

long long QBluetoothSocket_WriteData(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QBluetoothSocket*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(maxSize)));
}

void QBluetoothSocket_DestroyQBluetoothSocket(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->~QBluetoothSocket();
}

void QBluetoothSocket_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothSocket_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothSocket_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothSocket_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothSocket_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothSocket_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::customEvent(static_cast<QEvent*>(event));
}

class MyQBluetoothTransferManager: public QBluetoothTransferManager {
public:
	void Signal_Finished(QBluetoothTransferReply * reply) { callbackQBluetoothTransferManagerFinished(this, this->objectName().toUtf8().data(), reply); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothTransferManagerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothTransferManagerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothTransferManagerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QBluetoothTransferManager_Put(void* ptr, void* request, void* data){
	return static_cast<QBluetoothTransferManager*>(ptr)->put(*static_cast<QBluetoothTransferRequest*>(request), static_cast<QIODevice*>(data));
}

void* QBluetoothTransferManager_NewQBluetoothTransferManager(void* parent){
	return new QBluetoothTransferManager(static_cast<QObject*>(parent));
}

void QBluetoothTransferManager_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));;
}

void QBluetoothTransferManager_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));;
}

void QBluetoothTransferManager_Finished(void* ptr, void* reply){
	static_cast<QBluetoothTransferManager*>(ptr)->finished(static_cast<QBluetoothTransferReply*>(reply));
}

void QBluetoothTransferManager_DestroyQBluetoothTransferManager(void* ptr){
	static_cast<QBluetoothTransferManager*>(ptr)->~QBluetoothTransferManager();
}

void QBluetoothTransferManager_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothTransferManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferManager_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferManager_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothTransferManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferManager_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferManager_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothTransferManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothTransferManager_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::customEvent(static_cast<QEvent*>(event));
}

class MyQBluetoothTransferReply: public QBluetoothTransferReply {
public:
	void Signal_Error2(QBluetoothTransferReply::TransferError errorType) { callbackQBluetoothTransferReplyError2(this, this->objectName().toUtf8().data(), errorType); };
	void Signal_Finished(QBluetoothTransferReply * reply) { callbackQBluetoothTransferReplyFinished(this, this->objectName().toUtf8().data(), reply); };
	void Signal_TransferProgress(qint64 bytesTransferred, qint64 bytesTotal) { callbackQBluetoothTransferReplyTransferProgress(this, this->objectName().toUtf8().data(), static_cast<long long>(bytesTransferred), static_cast<long long>(bytesTotal)); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothTransferReplyTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothTransferReplyChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBluetoothTransferReplyCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QBluetoothTransferReply_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferReply*>(ptr), "abort");
}

void QBluetoothTransferReply_ConnectError2(void* ptr){
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&QBluetoothTransferReply::error), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&MyQBluetoothTransferReply::Signal_Error2));;
}

void QBluetoothTransferReply_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&QBluetoothTransferReply::error), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&MyQBluetoothTransferReply::Signal_Error2));;
}

void QBluetoothTransferReply_Error2(void* ptr, int errorType){
	static_cast<QBluetoothTransferReply*>(ptr)->error(static_cast<QBluetoothTransferReply::TransferError>(errorType));
}

int QBluetoothTransferReply_Error(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->error();
}

char* QBluetoothTransferReply_ErrorString(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->errorString().toUtf8().data();
}

void QBluetoothTransferReply_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));;
}

void QBluetoothTransferReply_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));;
}

void QBluetoothTransferReply_Finished(void* ptr, void* reply){
	static_cast<QBluetoothTransferReply*>(ptr)->finished(static_cast<QBluetoothTransferReply*>(reply));
}

int QBluetoothTransferReply_IsFinished(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isFinished();
}

int QBluetoothTransferReply_IsRunning(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isRunning();
}

void* QBluetoothTransferReply_Manager(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->manager();
}

void QBluetoothTransferReply_ConnectTransferProgress(void* ptr){
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(qint64, qint64)>(&QBluetoothTransferReply::transferProgress), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(qint64, qint64)>(&MyQBluetoothTransferReply::Signal_TransferProgress));;
}

void QBluetoothTransferReply_DisconnectTransferProgress(void* ptr){
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(qint64, qint64)>(&QBluetoothTransferReply::transferProgress), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(qint64, qint64)>(&MyQBluetoothTransferReply::Signal_TransferProgress));;
}

void QBluetoothTransferReply_TransferProgress(void* ptr, long long bytesTransferred, long long bytesTotal){
	static_cast<QBluetoothTransferReply*>(ptr)->transferProgress(static_cast<long long>(bytesTransferred), static_cast<long long>(bytesTotal));
}

void QBluetoothTransferReply_DestroyQBluetoothTransferReply(void* ptr){
	static_cast<QBluetoothTransferReply*>(ptr)->~QBluetoothTransferReply();
}

void QBluetoothTransferReply_TimerEvent(void* ptr, void* event){
	static_cast<MyQBluetoothTransferReply*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferReply_TimerEventDefault(void* ptr, void* event){
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferReply_ChildEvent(void* ptr, void* event){
	static_cast<MyQBluetoothTransferReply*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferReply_ChildEventDefault(void* ptr, void* event){
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferReply_CustomEvent(void* ptr, void* event){
	static_cast<MyQBluetoothTransferReply*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothTransferReply_CustomEventDefault(void* ptr, void* event){
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::customEvent(static_cast<QEvent*>(event));
}

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest(void* address){
	return new QBluetoothTransferRequest(*static_cast<QBluetoothAddress*>(address));
}

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest2(void* other){
	return new QBluetoothTransferRequest(*static_cast<QBluetoothTransferRequest*>(other));
}

void* QBluetoothTransferRequest_Attribute(void* ptr, int code, void* defaultValue){
	return new QVariant(static_cast<QBluetoothTransferRequest*>(ptr)->attribute(static_cast<QBluetoothTransferRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

void QBluetoothTransferRequest_SetAttribute(void* ptr, int code, void* value){
	static_cast<QBluetoothTransferRequest*>(ptr)->setAttribute(static_cast<QBluetoothTransferRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(void* ptr){
	static_cast<QBluetoothTransferRequest*>(ptr)->~QBluetoothTransferRequest();
}

void* QBluetoothUuid_NewQBluetoothUuid(){
	return new QBluetoothUuid();
}

void* QBluetoothUuid_NewQBluetoothUuid4(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::CharacteristicType>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid5(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::DescriptorType>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid2(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::ProtocolUuid>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid3(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::ServiceClassUuid>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid10(void* uuid){
	return new QBluetoothUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid9(char* uuid){
	return new QBluetoothUuid(QString(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid11(void* uuid){
	return new QBluetoothUuid(*static_cast<QUuid*>(uuid));
}

char* QBluetoothUuid_QBluetoothUuid_CharacteristicToString(int uuid){
	return QBluetoothUuid::characteristicToString(static_cast<QBluetoothUuid::CharacteristicType>(uuid)).toUtf8().data();
}

char* QBluetoothUuid_QBluetoothUuid_DescriptorToString(int uuid){
	return QBluetoothUuid::descriptorToString(static_cast<QBluetoothUuid::DescriptorType>(uuid)).toUtf8().data();
}

int QBluetoothUuid_MinimumSize(void* ptr){
	return static_cast<QBluetoothUuid*>(ptr)->minimumSize();
}

char* QBluetoothUuid_QBluetoothUuid_ProtocolToString(int uuid){
	return QBluetoothUuid::protocolToString(static_cast<QBluetoothUuid::ProtocolUuid>(uuid)).toUtf8().data();
}

char* QBluetoothUuid_QBluetoothUuid_ServiceClassToString(int uuid){
	return QBluetoothUuid::serviceClassToString(static_cast<QBluetoothUuid::ServiceClassUuid>(uuid)).toUtf8().data();
}

void QBluetoothUuid_DestroyQBluetoothUuid(void* ptr){
	static_cast<QBluetoothUuid*>(ptr)->~QBluetoothUuid();
}

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic(){
	return new QLowEnergyCharacteristic();
}

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(void* other){
	return new QLowEnergyCharacteristic(*static_cast<QLowEnergyCharacteristic*>(other));
}

int QLowEnergyCharacteristic_IsValid(void* ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->isValid();
}

char* QLowEnergyCharacteristic_Name(void* ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->name().toUtf8().data();
}

int QLowEnergyCharacteristic_Properties(void* ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->properties();
}

void* QLowEnergyCharacteristic_Value(void* ptr){
	return new QByteArray(static_cast<QLowEnergyCharacteristic*>(ptr)->value());
}

void QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(void* ptr){
	static_cast<QLowEnergyCharacteristic*>(ptr)->~QLowEnergyCharacteristic();
}

class MyQLowEnergyController: public QLowEnergyController {
public:
	void Signal_Connected() { callbackQLowEnergyControllerConnected(this, this->objectName().toUtf8().data()); };
	void Signal_Disconnected() { callbackQLowEnergyControllerDisconnected(this, this->objectName().toUtf8().data()); };
	void Signal_DiscoveryFinished() { callbackQLowEnergyControllerDiscoveryFinished(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QLowEnergyController::Error newError) { callbackQLowEnergyControllerError2(this, this->objectName().toUtf8().data(), newError); };
	void Signal_StateChanged(QLowEnergyController::ControllerState state) { callbackQLowEnergyControllerStateChanged(this, this->objectName().toUtf8().data(), state); };
	void timerEvent(QTimerEvent * event) { callbackQLowEnergyControllerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLowEnergyControllerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQLowEnergyControllerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QLowEnergyController_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));;
}

void QLowEnergyController_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));;
}

void QLowEnergyController_Connected(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->connected();
}

void QLowEnergyController_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));;
}

void QLowEnergyController_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));;
}

void QLowEnergyController_Disconnected(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->disconnected();
}

void QLowEnergyController_ConnectDiscoveryFinished(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));;
}

void QLowEnergyController_DisconnectDiscoveryFinished(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));;
}

void QLowEnergyController_DiscoveryFinished(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->discoveryFinished();
}

void QLowEnergyController_ConnectError2(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::Error)>(&QLowEnergyController::error), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::Error)>(&MyQLowEnergyController::Signal_Error2));;
}

void QLowEnergyController_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::Error)>(&QLowEnergyController::error), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::Error)>(&MyQLowEnergyController::Signal_Error2));;
}

void QLowEnergyController_Error2(void* ptr, int newError){
	static_cast<QLowEnergyController*>(ptr)->error(static_cast<QLowEnergyController::Error>(newError));
}

void QLowEnergyController_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));;
}

void QLowEnergyController_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));;
}

void QLowEnergyController_StateChanged(void* ptr, int state){
	static_cast<QLowEnergyController*>(ptr)->stateChanged(static_cast<QLowEnergyController::ControllerState>(state));
}

void* QLowEnergyController_NewQLowEnergyController(void* remoteDeviceInfo, void* parent){
	return new QLowEnergyController(*static_cast<QBluetoothDeviceInfo*>(remoteDeviceInfo), static_cast<QObject*>(parent));
}

void QLowEnergyController_ConnectToDevice(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->connectToDevice();
}

void* QLowEnergyController_CreateServiceObject(void* ptr, void* serviceUuid, void* parent){
	return static_cast<QLowEnergyController*>(ptr)->createServiceObject(*static_cast<QBluetoothUuid*>(serviceUuid), static_cast<QObject*>(parent));
}

void QLowEnergyController_DisconnectFromDevice(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->disconnectFromDevice();
}

void QLowEnergyController_DiscoverServices(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->discoverServices();
}

int QLowEnergyController_Error(void* ptr){
	return static_cast<QLowEnergyController*>(ptr)->error();
}

char* QLowEnergyController_ErrorString(void* ptr){
	return static_cast<QLowEnergyController*>(ptr)->errorString().toUtf8().data();
}

int QLowEnergyController_RemoteAddressType(void* ptr){
	return static_cast<QLowEnergyController*>(ptr)->remoteAddressType();
}

char* QLowEnergyController_RemoteName(void* ptr){
	return static_cast<QLowEnergyController*>(ptr)->remoteName().toUtf8().data();
}

void QLowEnergyController_SetRemoteAddressType(void* ptr, int ty){
	static_cast<QLowEnergyController*>(ptr)->setRemoteAddressType(static_cast<QLowEnergyController::RemoteAddressType>(ty));
}

int QLowEnergyController_State(void* ptr){
	return static_cast<QLowEnergyController*>(ptr)->state();
}

void QLowEnergyController_DestroyQLowEnergyController(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->~QLowEnergyController();
}

void QLowEnergyController_TimerEvent(void* ptr, void* event){
	static_cast<MyQLowEnergyController*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyController_TimerEventDefault(void* ptr, void* event){
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyController_ChildEvent(void* ptr, void* event){
	static_cast<MyQLowEnergyController*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyController_ChildEventDefault(void* ptr, void* event){
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyController_CustomEvent(void* ptr, void* event){
	static_cast<MyQLowEnergyController*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLowEnergyController_CustomEventDefault(void* ptr, void* event){
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::customEvent(static_cast<QEvent*>(event));
}

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor(){
	return new QLowEnergyDescriptor();
}

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor2(void* other){
	return new QLowEnergyDescriptor(*static_cast<QLowEnergyDescriptor*>(other));
}

int QLowEnergyDescriptor_IsValid(void* ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->isValid();
}

char* QLowEnergyDescriptor_Name(void* ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->name().toUtf8().data();
}

int QLowEnergyDescriptor_Type(void* ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->type();
}

void* QLowEnergyDescriptor_Value(void* ptr){
	return new QByteArray(static_cast<QLowEnergyDescriptor*>(ptr)->value());
}

void QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(void* ptr){
	static_cast<QLowEnergyDescriptor*>(ptr)->~QLowEnergyDescriptor();
}

class MyQLowEnergyService: public QLowEnergyService {
public:
	void Signal_Error2(QLowEnergyService::ServiceError newError) { callbackQLowEnergyServiceError2(this, this->objectName().toUtf8().data(), newError); };
	void Signal_StateChanged(QLowEnergyService::ServiceState newState) { callbackQLowEnergyServiceStateChanged(this, this->objectName().toUtf8().data(), newState); };
	void timerEvent(QTimerEvent * event) { callbackQLowEnergyServiceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLowEnergyServiceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQLowEnergyServiceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QLowEnergyService_ConnectError2(void* ptr){
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceError)>(&QLowEnergyService::error), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceError)>(&MyQLowEnergyService::Signal_Error2));;
}

void QLowEnergyService_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceError)>(&QLowEnergyService::error), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceError)>(&MyQLowEnergyService::Signal_Error2));;
}

void QLowEnergyService_Error2(void* ptr, int newError){
	static_cast<QLowEnergyService*>(ptr)->error(static_cast<QLowEnergyService::ServiceError>(newError));
}

void QLowEnergyService_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));;
}

void QLowEnergyService_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));;
}

void QLowEnergyService_StateChanged(void* ptr, int newState){
	static_cast<QLowEnergyService*>(ptr)->stateChanged(static_cast<QLowEnergyService::ServiceState>(newState));
}

int QLowEnergyService_Contains(void* ptr, void* characteristic){
	return static_cast<QLowEnergyService*>(ptr)->contains(*static_cast<QLowEnergyCharacteristic*>(characteristic));
}

int QLowEnergyService_Contains2(void* ptr, void* descriptor){
	return static_cast<QLowEnergyService*>(ptr)->contains(*static_cast<QLowEnergyDescriptor*>(descriptor));
}

void QLowEnergyService_DiscoverDetails(void* ptr){
	static_cast<QLowEnergyService*>(ptr)->discoverDetails();
}

int QLowEnergyService_Error(void* ptr){
	return static_cast<QLowEnergyService*>(ptr)->error();
}

void QLowEnergyService_ReadCharacteristic(void* ptr, void* characteristic){
	static_cast<QLowEnergyService*>(ptr)->readCharacteristic(*static_cast<QLowEnergyCharacteristic*>(characteristic));
}

void QLowEnergyService_ReadDescriptor(void* ptr, void* descriptor){
	static_cast<QLowEnergyService*>(ptr)->readDescriptor(*static_cast<QLowEnergyDescriptor*>(descriptor));
}

char* QLowEnergyService_ServiceName(void* ptr){
	return static_cast<QLowEnergyService*>(ptr)->serviceName().toUtf8().data();
}

int QLowEnergyService_State(void* ptr){
	return static_cast<QLowEnergyService*>(ptr)->state();
}

int QLowEnergyService_Type(void* ptr){
	return static_cast<QLowEnergyService*>(ptr)->type();
}

void QLowEnergyService_WriteCharacteristic(void* ptr, void* characteristic, void* newValue, int mode){
	static_cast<QLowEnergyService*>(ptr)->writeCharacteristic(*static_cast<QLowEnergyCharacteristic*>(characteristic), *static_cast<QByteArray*>(newValue), static_cast<QLowEnergyService::WriteMode>(mode));
}

void QLowEnergyService_WriteDescriptor(void* ptr, void* descriptor, void* newValue){
	static_cast<QLowEnergyService*>(ptr)->writeDescriptor(*static_cast<QLowEnergyDescriptor*>(descriptor), *static_cast<QByteArray*>(newValue));
}

void QLowEnergyService_DestroyQLowEnergyService(void* ptr){
	static_cast<QLowEnergyService*>(ptr)->~QLowEnergyService();
}

void QLowEnergyService_TimerEvent(void* ptr, void* event){
	static_cast<MyQLowEnergyService*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyService_TimerEventDefault(void* ptr, void* event){
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyService_ChildEvent(void* ptr, void* event){
	static_cast<MyQLowEnergyService*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyService_ChildEventDefault(void* ptr, void* event){
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyService_CustomEvent(void* ptr, void* event){
	static_cast<MyQLowEnergyService*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLowEnergyService_CustomEventDefault(void* ptr, void* event){
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::customEvent(static_cast<QEvent*>(event));
}

