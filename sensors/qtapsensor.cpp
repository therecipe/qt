#include "qtapsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QTapSensor>
#include "_cgo_export.h"

class MyQTapSensor: public QTapSensor {
public:
void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents){callbackQTapSensorReturnDoubleTapEventsChanged(this->objectName().toUtf8().data(), returnDoubleTapEvents);};
};

QtObjectPtr QTapSensor_Reading(QtObjectPtr ptr){
	return static_cast<QTapSensor*>(ptr)->reading();
}

int QTapSensor_ReturnDoubleTapEvents(QtObjectPtr ptr){
	return static_cast<QTapSensor*>(ptr)->returnDoubleTapEvents();
}

void QTapSensor_SetReturnDoubleTapEvents(QtObjectPtr ptr, int returnDoubleTapEvents){
	static_cast<QTapSensor*>(ptr)->setReturnDoubleTapEvents(returnDoubleTapEvents != 0);
}

QtObjectPtr QTapSensor_NewQTapSensor(QtObjectPtr parent){
	return new QTapSensor(static_cast<QObject*>(parent));
}

void QTapSensor_ConnectReturnDoubleTapEventsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));;
}

void QTapSensor_DisconnectReturnDoubleTapEventsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));;
}

void QTapSensor_DestroyQTapSensor(QtObjectPtr ptr){
	static_cast<QTapSensor*>(ptr)->~QTapSensor();
}

