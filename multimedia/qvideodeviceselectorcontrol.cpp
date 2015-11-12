#include "qvideodeviceselectorcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QVideoDeviceSelectorControl>
#include "_cgo_export.h"

class MyQVideoDeviceSelectorControl: public QVideoDeviceSelectorControl {
public:
void Signal_DevicesChanged(){callbackQVideoDeviceSelectorControlDevicesChanged(this->objectName().toUtf8().data());};
void Signal_SelectedDeviceChanged(int index){callbackQVideoDeviceSelectorControlSelectedDeviceChanged(this->objectName().toUtf8().data(), index);};
};

int QVideoDeviceSelectorControl_DefaultDevice(void* ptr){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->defaultDevice();
}

int QVideoDeviceSelectorControl_DeviceCount(void* ptr){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceCount();
}

char* QVideoDeviceSelectorControl_DeviceDescription(void* ptr, int index){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceDescription(index).toUtf8().data();
}

char* QVideoDeviceSelectorControl_DeviceName(void* ptr, int index){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceName(index).toUtf8().data();
}

void QVideoDeviceSelectorControl_ConnectDevicesChanged(void* ptr){
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)()>(&QVideoDeviceSelectorControl::devicesChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)()>(&MyQVideoDeviceSelectorControl::Signal_DevicesChanged));;
}

void QVideoDeviceSelectorControl_DisconnectDevicesChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)()>(&QVideoDeviceSelectorControl::devicesChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)()>(&MyQVideoDeviceSelectorControl::Signal_DevicesChanged));;
}

int QVideoDeviceSelectorControl_SelectedDevice(void* ptr){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->selectedDevice();
}

void QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(void* ptr){
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(int)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(int)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged));;
}

void QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(int)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(int)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged));;
}

void QVideoDeviceSelectorControl_SetSelectedDevice(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QVideoDeviceSelectorControl*>(ptr), "setSelectedDevice", Q_ARG(int, index));
}

void QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(void* ptr){
	static_cast<QVideoDeviceSelectorControl*>(ptr)->~QVideoDeviceSelectorControl();
}

