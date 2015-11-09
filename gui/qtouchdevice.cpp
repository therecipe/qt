#include "qtouchdevice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTouchDevice>
#include "_cgo_export.h"

class MyQTouchDevice: public QTouchDevice {
public:
};

void* QTouchDevice_NewQTouchDevice(){
	return new QTouchDevice();
}

int QTouchDevice_Capabilities(void* ptr){
	return static_cast<QTouchDevice*>(ptr)->capabilities();
}

int QTouchDevice_MaximumTouchPoints(void* ptr){
	return static_cast<QTouchDevice*>(ptr)->maximumTouchPoints();
}

char* QTouchDevice_Name(void* ptr){
	return static_cast<QTouchDevice*>(ptr)->name().toUtf8().data();
}

void QTouchDevice_SetCapabilities(void* ptr, int caps){
	static_cast<QTouchDevice*>(ptr)->setCapabilities(static_cast<QTouchDevice::CapabilityFlag>(caps));
}

void QTouchDevice_SetMaximumTouchPoints(void* ptr, int max){
	static_cast<QTouchDevice*>(ptr)->setMaximumTouchPoints(max);
}

void QTouchDevice_SetName(void* ptr, char* name){
	static_cast<QTouchDevice*>(ptr)->setName(QString(name));
}

void QTouchDevice_SetType(void* ptr, int devType){
	static_cast<QTouchDevice*>(ptr)->setType(static_cast<QTouchDevice::DeviceType>(devType));
}

int QTouchDevice_Type(void* ptr){
	return static_cast<QTouchDevice*>(ptr)->type();
}

void QTouchDevice_DestroyQTouchDevice(void* ptr){
	static_cast<QTouchDevice*>(ptr)->~QTouchDevice();
}

