#include "qlowenergyservice.h"
#include <QByteArray>
#include <QLowEnergyDescriptor>
#include <QLowEnergyCharacteristic>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QLowEnergyService>
#include "_cgo_export.h"

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

