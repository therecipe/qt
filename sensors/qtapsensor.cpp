#include "qtapsensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTapSensor>
#include "_cgo_export.h"

class MyQTapSensor: public QTapSensor {
public:
void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents){callbackQTapSensorReturnDoubleTapEventsChanged(this->objectName().toUtf8().data(), returnDoubleTapEvents);};
};

void* QTapSensor_Reading(void* ptr){
	return static_cast<QTapSensor*>(ptr)->reading();
}

int QTapSensor_ReturnDoubleTapEvents(void* ptr){
	return static_cast<QTapSensor*>(ptr)->returnDoubleTapEvents();
}

void QTapSensor_SetReturnDoubleTapEvents(void* ptr, int returnDoubleTapEvents){
	static_cast<QTapSensor*>(ptr)->setReturnDoubleTapEvents(returnDoubleTapEvents != 0);
}

void* QTapSensor_NewQTapSensor(void* parent){
	return new QTapSensor(static_cast<QObject*>(parent));
}

void QTapSensor_ConnectReturnDoubleTapEventsChanged(void* ptr){
	QObject::connect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));;
}

void QTapSensor_DisconnectReturnDoubleTapEventsChanged(void* ptr){
	QObject::disconnect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));;
}

void QTapSensor_DestroyQTapSensor(void* ptr){
	static_cast<QTapSensor*>(ptr)->~QTapSensor();
}

