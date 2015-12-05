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
#include <QIODevice>
#include <QLowEnergyCharacteristic>
#include <QLowEnergyController>
#include <QLowEnergyDescriptor>
#include <QLowEnergyService>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QUuid>
#include <QVariant>

class MyQBluetoothAddress: public QBluetoothAddress {
public:
};

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
void Signal_Canceled(){callbackQBluetoothDeviceDiscoveryAgentCanceled(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQBluetoothDeviceDiscoveryAgentFinished(this->objectName().toUtf8().data());};
};

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
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

class MyQBluetoothDeviceInfo: public QBluetoothDeviceInfo {
public:
};

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

class MyQBluetoothHostInfo: public QBluetoothHostInfo {
public:
};

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
void Signal_Error(QBluetoothLocalDevice::Error error){callbackQBluetoothLocalDeviceError(this->objectName().toUtf8().data(), error);};
void Signal_HostModeStateChanged(QBluetoothLocalDevice::HostMode state){callbackQBluetoothLocalDeviceHostModeStateChanged(this->objectName().toUtf8().data(), state);};
};

void QBluetoothLocalDevice_ConnectError(void* ptr){
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));;
}

void QBluetoothLocalDevice_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));;
}

void QBluetoothLocalDevice_ConnectHostModeStateChanged(void* ptr){
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));;
}

void QBluetoothLocalDevice_DisconnectHostModeStateChanged(void* ptr){
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));;
}

int QBluetoothLocalDevice_IsValid(void* ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->isValid();
}

void QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(void* ptr){
	static_cast<QBluetoothLocalDevice*>(ptr)->~QBluetoothLocalDevice();
}

void* QBluetoothLocalDevice_NewQBluetoothLocalDevice(void* parent){
	return new QBluetoothLocalDevice(static_cast<QObject*>(parent));
}

void* QBluetoothLocalDevice_NewQBluetoothLocalDevice2(void* address, void* parent){
	return new QBluetoothLocalDevice(*static_cast<QBluetoothAddress*>(address), static_cast<QObject*>(parent));
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

class MyQBluetoothServer: public QBluetoothServer {
public:
void Signal_NewConnection(){callbackQBluetoothServerNewConnection(this->objectName().toUtf8().data());};
};

void* QBluetoothServer_NewQBluetoothServer(int serverType, void* parent){
	return new QBluetoothServer(static_cast<QBluetoothServiceInfo::Protocol>(serverType), static_cast<QObject*>(parent));
}

void QBluetoothServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
}

void QBluetoothServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));;
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

class MyQBluetoothServiceDiscoveryAgent: public QBluetoothServiceDiscoveryAgent {
public:
void Signal_Canceled(){callbackQBluetoothServiceDiscoveryAgentCanceled(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQBluetoothServiceDiscoveryAgentFinished(this->objectName().toUtf8().data());};
};

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
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

class MyQBluetoothServiceInfo: public QBluetoothServiceInfo {
public:
};

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
void Signal_Connected(){callbackQBluetoothSocketConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQBluetoothSocketDisconnected(this->objectName().toUtf8().data());};
void Signal_StateChanged(QBluetoothSocket::SocketState state){callbackQBluetoothSocketStateChanged(this->objectName().toUtf8().data(), state);};
};

void QBluetoothSocket_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));;
}

void QBluetoothSocket_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));;
}

void QBluetoothSocket_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void QBluetoothSocket_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));;
}

void* QBluetoothSocket_NewQBluetoothSocket(int socketType, void* parent){
	return new QBluetoothSocket(static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QObject*>(parent));
}

void* QBluetoothSocket_NewQBluetoothSocket2(void* parent){
	return new QBluetoothSocket(static_cast<QObject*>(parent));
}

void QBluetoothSocket_Abort(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->abort();
}

int QBluetoothSocket_CanReadLine(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->canReadLine();
}

void QBluetoothSocket_Close(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->close();
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

int QBluetoothSocket_SetSocketDescriptor(void* ptr, int socketDescriptor, int socketType, int socketState, int openMode){
	return static_cast<QBluetoothSocket*>(ptr)->setSocketDescriptor(socketDescriptor, static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QBluetoothSocket::SocketState>(socketState), static_cast<QIODevice::OpenModeFlag>(openMode));
}

int QBluetoothSocket_SocketDescriptor(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketDescriptor();
}

int QBluetoothSocket_SocketType(void* ptr){
	return static_cast<QBluetoothSocket*>(ptr)->socketType();
}

void QBluetoothSocket_DestroyQBluetoothSocket(void* ptr){
	static_cast<QBluetoothSocket*>(ptr)->~QBluetoothSocket();
}

class MyQBluetoothTransferManager: public QBluetoothTransferManager {
public:
void Signal_Finished(QBluetoothTransferReply * reply){callbackQBluetoothTransferManagerFinished(this->objectName().toUtf8().data(), reply);};
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

void QBluetoothTransferManager_DestroyQBluetoothTransferManager(void* ptr){
	static_cast<QBluetoothTransferManager*>(ptr)->~QBluetoothTransferManager();
}

class MyQBluetoothTransferReply: public QBluetoothTransferReply {
public:
void Signal_Finished(QBluetoothTransferReply * reply){callbackQBluetoothTransferReplyFinished(this->objectName().toUtf8().data(), reply);};
};

void QBluetoothTransferReply_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferReply*>(ptr), "abort");
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

int QBluetoothTransferReply_IsFinished(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isFinished();
}

int QBluetoothTransferReply_IsRunning(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isRunning();
}

void* QBluetoothTransferReply_Manager(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->manager();
}

void QBluetoothTransferReply_DestroyQBluetoothTransferReply(void* ptr){
	static_cast<QBluetoothTransferReply*>(ptr)->~QBluetoothTransferReply();
}

class MyQBluetoothTransferRequest: public QBluetoothTransferRequest {
public:
};

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

class MyQBluetoothUuid: public QBluetoothUuid {
public:
};

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

class MyQLowEnergyCharacteristic: public QLowEnergyCharacteristic {
public:
};

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
void Signal_Connected(){callbackQLowEnergyControllerConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQLowEnergyControllerDisconnected(this->objectName().toUtf8().data());};
void Signal_DiscoveryFinished(){callbackQLowEnergyControllerDiscoveryFinished(this->objectName().toUtf8().data());};
void Signal_StateChanged(QLowEnergyController::ControllerState state){callbackQLowEnergyControllerStateChanged(this->objectName().toUtf8().data(), state);};
};

void QLowEnergyController_ConnectConnected(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));;
}

void QLowEnergyController_DisconnectConnected(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));;
}

void QLowEnergyController_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));;
}

void QLowEnergyController_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));;
}

void QLowEnergyController_ConnectDiscoveryFinished(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));;
}

void QLowEnergyController_DisconnectDiscoveryFinished(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));;
}

void QLowEnergyController_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));;
}

void QLowEnergyController_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));;
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

void QLowEnergyController_DestroyQLowEnergyController(void* ptr){
	static_cast<QLowEnergyController*>(ptr)->~QLowEnergyController();
}

class MyQLowEnergyDescriptor: public QLowEnergyDescriptor {
public:
};

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
void Signal_StateChanged(QLowEnergyService::ServiceState newState){callbackQLowEnergyServiceStateChanged(this->objectName().toUtf8().data(), newState);};
};

void QLowEnergyService_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));;
}

void QLowEnergyService_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));;
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

