#include "qrotationsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QRotationSensor>
#include "_cgo_export.h"

class MyQRotationSensor: public QRotationSensor {
public:
void Signal_HasZChanged(bool hasZ){callbackQRotationSensorHasZChanged(this->objectName().toUtf8().data(), hasZ);};
};

int QRotationSensor_HasZ(void* ptr){
	return static_cast<QRotationSensor*>(ptr)->hasZ();
}

void* QRotationSensor_Reading(void* ptr){
	return static_cast<QRotationSensor*>(ptr)->reading();
}

void* QRotationSensor_NewQRotationSensor(void* parent){
	return new QRotationSensor(static_cast<QObject*>(parent));
}

void QRotationSensor_ConnectHasZChanged(void* ptr){
	QObject::connect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));;
}

void QRotationSensor_DisconnectHasZChanged(void* ptr){
	QObject::disconnect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));;
}

void QRotationSensor_SetHasZ(void* ptr, int hasZ){
	static_cast<QRotationSensor*>(ptr)->setHasZ(hasZ != 0);
}

void QRotationSensor_DestroyQRotationSensor(void* ptr){
	static_cast<QRotationSensor*>(ptr)->~QRotationSensor();
}

