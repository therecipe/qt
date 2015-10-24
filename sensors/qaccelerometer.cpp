#include "qaccelerometer.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccelerometer>
#include "_cgo_export.h"

class MyQAccelerometer: public QAccelerometer {
public:
void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode){callbackQAccelerometerAccelerationModeChanged(this->objectName().toUtf8().data(), accelerationMode);};
};

int QAccelerometer_AccelerationMode(QtObjectPtr ptr){
	return static_cast<QAccelerometer*>(ptr)->accelerationMode();
}

QtObjectPtr QAccelerometer_Reading(QtObjectPtr ptr){
	return static_cast<QAccelerometer*>(ptr)->reading();
}

QtObjectPtr QAccelerometer_NewQAccelerometer(QtObjectPtr parent){
	return new QAccelerometer(static_cast<QObject*>(parent));
}

void QAccelerometer_ConnectAccelerationModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));;
}

void QAccelerometer_DisconnectAccelerationModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));;
}

void QAccelerometer_SetAccelerationMode(QtObjectPtr ptr, int accelerationMode){
	static_cast<QAccelerometer*>(ptr)->setAccelerationMode(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_DestroyQAccelerometer(QtObjectPtr ptr){
	static_cast<QAccelerometer*>(ptr)->~QAccelerometer();
}

