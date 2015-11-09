#include "qbluetoothdeviceinfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothUuid>
#include <QString>
#include <QVariant>
#include <QBluetoothDeviceInfo>
#include "_cgo_export.h"

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

