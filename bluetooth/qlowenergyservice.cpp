#include "qlowenergyservice.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLowEnergyCharacteristic>
#include <QByteArray>
#include <QLowEnergyDescriptor>
#include <QLowEnergyService>
#include "_cgo_export.h"

class MyQLowEnergyService: public QLowEnergyService {
public:
void Signal_StateChanged(QLowEnergyService::ServiceState newState){callbackQLowEnergyServiceStateChanged(this->objectName().toUtf8().data(), newState);};
};

void QLowEnergyService_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));;
}

void QLowEnergyService_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));;
}

int QLowEnergyService_Contains(QtObjectPtr ptr, QtObjectPtr characteristic){
	return static_cast<QLowEnergyService*>(ptr)->contains(*static_cast<QLowEnergyCharacteristic*>(characteristic));
}

int QLowEnergyService_Contains2(QtObjectPtr ptr, QtObjectPtr descriptor){
	return static_cast<QLowEnergyService*>(ptr)->contains(*static_cast<QLowEnergyDescriptor*>(descriptor));
}

void QLowEnergyService_DiscoverDetails(QtObjectPtr ptr){
	static_cast<QLowEnergyService*>(ptr)->discoverDetails();
}

int QLowEnergyService_Error(QtObjectPtr ptr){
	return static_cast<QLowEnergyService*>(ptr)->error();
}

void QLowEnergyService_ReadCharacteristic(QtObjectPtr ptr, QtObjectPtr characteristic){
	static_cast<QLowEnergyService*>(ptr)->readCharacteristic(*static_cast<QLowEnergyCharacteristic*>(characteristic));
}

void QLowEnergyService_ReadDescriptor(QtObjectPtr ptr, QtObjectPtr descriptor){
	static_cast<QLowEnergyService*>(ptr)->readDescriptor(*static_cast<QLowEnergyDescriptor*>(descriptor));
}

char* QLowEnergyService_ServiceName(QtObjectPtr ptr){
	return static_cast<QLowEnergyService*>(ptr)->serviceName().toUtf8().data();
}

int QLowEnergyService_Type(QtObjectPtr ptr){
	return static_cast<QLowEnergyService*>(ptr)->type();
}

void QLowEnergyService_WriteCharacteristic(QtObjectPtr ptr, QtObjectPtr characteristic, QtObjectPtr newValue, int mode){
	static_cast<QLowEnergyService*>(ptr)->writeCharacteristic(*static_cast<QLowEnergyCharacteristic*>(characteristic), *static_cast<QByteArray*>(newValue), static_cast<QLowEnergyService::WriteMode>(mode));
}

void QLowEnergyService_WriteDescriptor(QtObjectPtr ptr, QtObjectPtr descriptor, QtObjectPtr newValue){
	static_cast<QLowEnergyService*>(ptr)->writeDescriptor(*static_cast<QLowEnergyDescriptor*>(descriptor), *static_cast<QByteArray*>(newValue));
}

void QLowEnergyService_DestroyQLowEnergyService(QtObjectPtr ptr){
	static_cast<QLowEnergyService*>(ptr)->~QLowEnergyService();
}

