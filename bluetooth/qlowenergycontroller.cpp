#include "qlowenergycontroller.h"
#include <QBluetoothUuid>
#include <QObject>
#include <QBluetoothDeviceInfo>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLowEnergyController>
#include "_cgo_export.h"

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

