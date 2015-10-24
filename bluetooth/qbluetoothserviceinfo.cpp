#include "qbluetoothserviceinfo.h"
#include <QBluetoothDeviceInfo>
#include <QBluetoothAddress>
#include <QBluetoothUuid>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

char* QBluetoothServiceInfo_ServiceDescription(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serviceDescription().toUtf8().data();
}

char* QBluetoothServiceInfo_ServiceName(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serviceName().toUtf8().data();
}

char* QBluetoothServiceInfo_ServiceProvider(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serviceProvider().toUtf8().data();
}

void QBluetoothServiceInfo_SetServiceDescription(QtObjectPtr ptr, char* description){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceDescription(QString(description));
}

void QBluetoothServiceInfo_SetServiceName(QtObjectPtr ptr, char* name){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceName(QString(name));
}

void QBluetoothServiceInfo_SetServiceProvider(QtObjectPtr ptr, char* provider){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceProvider(QString(provider));
}

void QBluetoothServiceInfo_SetServiceUuid(QtObjectPtr ptr, QtObjectPtr uuid){
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceUuid(*static_cast<QBluetoothUuid*>(uuid));
}

QtObjectPtr QBluetoothServiceInfo_NewQBluetoothServiceInfo(){
	return new QBluetoothServiceInfo();
}

QtObjectPtr QBluetoothServiceInfo_NewQBluetoothServiceInfo2(QtObjectPtr other){
	return new QBluetoothServiceInfo(*static_cast<QBluetoothServiceInfo*>(other));
}

int QBluetoothServiceInfo_IsComplete(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->isComplete();
}

int QBluetoothServiceInfo_IsRegistered(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->isRegistered();
}

int QBluetoothServiceInfo_IsValid(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->isValid();
}

int QBluetoothServiceInfo_ProtocolServiceMultiplexer(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->protocolServiceMultiplexer();
}

int QBluetoothServiceInfo_RegisterService(QtObjectPtr ptr, QtObjectPtr localAdapter){
	return static_cast<QBluetoothServiceInfo*>(ptr)->registerService(*static_cast<QBluetoothAddress*>(localAdapter));
}

int QBluetoothServiceInfo_ServerChannel(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->serverChannel();
}

void QBluetoothServiceInfo_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QBluetoothServiceInfo*>(ptr)->setDevice(*static_cast<QBluetoothDeviceInfo*>(device));
}

int QBluetoothServiceInfo_SocketProtocol(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->socketProtocol();
}

int QBluetoothServiceInfo_UnregisterService(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceInfo*>(ptr)->unregisterService();
}

void QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(QtObjectPtr ptr){
	static_cast<QBluetoothServiceInfo*>(ptr)->~QBluetoothServiceInfo();
}

