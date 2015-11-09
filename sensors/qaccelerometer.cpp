#include "qaccelerometer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAccelerometer>
#include "_cgo_export.h"

class MyQAccelerometer: public QAccelerometer {
public:
void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode){callbackQAccelerometerAccelerationModeChanged(this->objectName().toUtf8().data(), accelerationMode);};
};

int QAccelerometer_AccelerationMode(void* ptr){
	return static_cast<QAccelerometer*>(ptr)->accelerationMode();
}

void* QAccelerometer_Reading(void* ptr){
	return static_cast<QAccelerometer*>(ptr)->reading();
}

void* QAccelerometer_NewQAccelerometer(void* parent){
	return new QAccelerometer(static_cast<QObject*>(parent));
}

void QAccelerometer_ConnectAccelerationModeChanged(void* ptr){
	QObject::connect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));;
}

void QAccelerometer_DisconnectAccelerationModeChanged(void* ptr){
	QObject::disconnect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));;
}

void QAccelerometer_SetAccelerationMode(void* ptr, int accelerationMode){
	static_cast<QAccelerometer*>(ptr)->setAccelerationMode(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_DestroyQAccelerometer(void* ptr){
	static_cast<QAccelerometer*>(ptr)->~QAccelerometer();
}

