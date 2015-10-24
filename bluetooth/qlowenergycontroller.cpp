#include "qlowenergycontroller.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothUuid>
#include <QObject>
#include <QBluetoothDeviceInfo>
#include <QString>
#include <QLowEnergyController>
#include "_cgo_export.h"

class MyQLowEnergyController: public QLowEnergyController {
public:
void Signal_Connected(){callbackQLowEnergyControllerConnected(this->objectName().toUtf8().data());};
void Signal_Disconnected(){callbackQLowEnergyControllerDisconnected(this->objectName().toUtf8().data());};
void Signal_DiscoveryFinished(){callbackQLowEnergyControllerDiscoveryFinished(this->objectName().toUtf8().data());};
void Signal_StateChanged(QLowEnergyController::ControllerState state){callbackQLowEnergyControllerStateChanged(this->objectName().toUtf8().data(), state);};
};

void QLowEnergyController_ConnectConnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));;
}

void QLowEnergyController_DisconnectConnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));;
}

void QLowEnergyController_ConnectDisconnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));;
}

void QLowEnergyController_DisconnectDisconnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));;
}

void QLowEnergyController_ConnectDiscoveryFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));;
}

void QLowEnergyController_DisconnectDiscoveryFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));;
}

void QLowEnergyController_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));;
}

void QLowEnergyController_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));;
}

QtObjectPtr QLowEnergyController_NewQLowEnergyController(QtObjectPtr remoteDeviceInfo, QtObjectPtr parent){
	return new QLowEnergyController(*static_cast<QBluetoothDeviceInfo*>(remoteDeviceInfo), static_cast<QObject*>(parent));
}

void QLowEnergyController_ConnectToDevice(QtObjectPtr ptr){
	static_cast<QLowEnergyController*>(ptr)->connectToDevice();
}

QtObjectPtr QLowEnergyController_CreateServiceObject(QtObjectPtr ptr, QtObjectPtr serviceUuid, QtObjectPtr parent){
	return static_cast<QLowEnergyController*>(ptr)->createServiceObject(*static_cast<QBluetoothUuid*>(serviceUuid), static_cast<QObject*>(parent));
}

void QLowEnergyController_DisconnectFromDevice(QtObjectPtr ptr){
	static_cast<QLowEnergyController*>(ptr)->disconnectFromDevice();
}

void QLowEnergyController_DiscoverServices(QtObjectPtr ptr){
	static_cast<QLowEnergyController*>(ptr)->discoverServices();
}

int QLowEnergyController_Error(QtObjectPtr ptr){
	return static_cast<QLowEnergyController*>(ptr)->error();
}

char* QLowEnergyController_ErrorString(QtObjectPtr ptr){
	return static_cast<QLowEnergyController*>(ptr)->errorString().toUtf8().data();
}

int QLowEnergyController_RemoteAddressType(QtObjectPtr ptr){
	return static_cast<QLowEnergyController*>(ptr)->remoteAddressType();
}

char* QLowEnergyController_RemoteName(QtObjectPtr ptr){
	return static_cast<QLowEnergyController*>(ptr)->remoteName().toUtf8().data();
}

void QLowEnergyController_SetRemoteAddressType(QtObjectPtr ptr, int ty){
	static_cast<QLowEnergyController*>(ptr)->setRemoteAddressType(static_cast<QLowEnergyController::RemoteAddressType>(ty));
}

void QLowEnergyController_DestroyQLowEnergyController(QtObjectPtr ptr){
	static_cast<QLowEnergyController*>(ptr)->~QLowEnergyController();
}

