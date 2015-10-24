#include "qbluetoothlocaldevice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothAddress>
#include <QMetaObject>
#include <QObject>
#include <QBluetoothLocalDevice>
#include "_cgo_export.h"

class MyQBluetoothLocalDevice: public QBluetoothLocalDevice {
public:
void Signal_Error(QBluetoothLocalDevice::Error error){callbackQBluetoothLocalDeviceError(this->objectName().toUtf8().data(), error);};
void Signal_HostModeStateChanged(QBluetoothLocalDevice::HostMode state){callbackQBluetoothLocalDeviceHostModeStateChanged(this->objectName().toUtf8().data(), state);};
};

void QBluetoothLocalDevice_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));;
}

void QBluetoothLocalDevice_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));;
}

void QBluetoothLocalDevice_ConnectHostModeStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));;
}

void QBluetoothLocalDevice_DisconnectHostModeStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));;
}

int QBluetoothLocalDevice_IsValid(QtObjectPtr ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->isValid();
}

void QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(QtObjectPtr ptr){
	static_cast<QBluetoothLocalDevice*>(ptr)->~QBluetoothLocalDevice();
}

QtObjectPtr QBluetoothLocalDevice_NewQBluetoothLocalDevice(QtObjectPtr parent){
	return new QBluetoothLocalDevice(static_cast<QObject*>(parent));
}

QtObjectPtr QBluetoothLocalDevice_NewQBluetoothLocalDevice2(QtObjectPtr address, QtObjectPtr parent){
	return new QBluetoothLocalDevice(*static_cast<QBluetoothAddress*>(address), static_cast<QObject*>(parent));
}

int QBluetoothLocalDevice_HostMode(QtObjectPtr ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->hostMode();
}

char* QBluetoothLocalDevice_Name(QtObjectPtr ptr){
	return static_cast<QBluetoothLocalDevice*>(ptr)->name().toUtf8().data();
}

void QBluetoothLocalDevice_PairingConfirmation(QtObjectPtr ptr, int accept){
	QMetaObject::invokeMethod(static_cast<QBluetoothLocalDevice*>(ptr), "pairingConfirmation", Q_ARG(bool, accept != 0));
}

int QBluetoothLocalDevice_PairingStatus(QtObjectPtr ptr, QtObjectPtr address){
	return static_cast<QBluetoothLocalDevice*>(ptr)->pairingStatus(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothLocalDevice_PowerOn(QtObjectPtr ptr){
	static_cast<QBluetoothLocalDevice*>(ptr)->powerOn();
}

void QBluetoothLocalDevice_RequestPairing(QtObjectPtr ptr, QtObjectPtr address, int pairing){
	static_cast<QBluetoothLocalDevice*>(ptr)->requestPairing(*static_cast<QBluetoothAddress*>(address), static_cast<QBluetoothLocalDevice::Pairing>(pairing));
}

void QBluetoothLocalDevice_SetHostMode(QtObjectPtr ptr, int mode){
	static_cast<QBluetoothLocalDevice*>(ptr)->setHostMode(static_cast<QBluetoothLocalDevice::HostMode>(mode));
}

