#include "qbluetoothserviceinfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothUuid>
#include <QBluetoothAddress>
#include <QBluetoothDeviceInfo>
#include <QString>
#include <QVariant>
#include <QBluetoothServiceInfo>
#include "_cgo_export.h"

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

