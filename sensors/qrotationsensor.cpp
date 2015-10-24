#include "qrotationsensor.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QRotationSensor>
#include "_cgo_export.h"

class MyQRotationSensor: public QRotationSensor {
public:
void Signal_HasZChanged(bool hasZ){callbackQRotationSensorHasZChanged(this->objectName().toUtf8().data(), hasZ);};
};

int QRotationSensor_HasZ(QtObjectPtr ptr){
	return static_cast<QRotationSensor*>(ptr)->hasZ();
}

QtObjectPtr QRotationSensor_Reading(QtObjectPtr ptr){
	return static_cast<QRotationSensor*>(ptr)->reading();
}

QtObjectPtr QRotationSensor_NewQRotationSensor(QtObjectPtr parent){
	return new QRotationSensor(static_cast<QObject*>(parent));
}

void QRotationSensor_ConnectHasZChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));;
}

void QRotationSensor_DisconnectHasZChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));;
}

void QRotationSensor_SetHasZ(QtObjectPtr ptr, int hasZ){
	static_cast<QRotationSensor*>(ptr)->setHasZ(hasZ != 0);
}

void QRotationSensor_DestroyQRotationSensor(QtObjectPtr ptr){
	static_cast<QRotationSensor*>(ptr)->~QRotationSensor();
}

