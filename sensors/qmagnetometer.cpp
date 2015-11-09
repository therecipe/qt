#include "qmagnetometer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMagnetometer>
#include "_cgo_export.h"

class MyQMagnetometer: public QMagnetometer {
public:
void Signal_ReturnGeoValuesChanged(bool returnGeoValues){callbackQMagnetometerReturnGeoValuesChanged(this->objectName().toUtf8().data(), returnGeoValues);};
};

void* QMagnetometer_Reading(void* ptr){
	return static_cast<QMagnetometer*>(ptr)->reading();
}

int QMagnetometer_ReturnGeoValues(void* ptr){
	return static_cast<QMagnetometer*>(ptr)->returnGeoValues();
}

void QMagnetometer_SetReturnGeoValues(void* ptr, int returnGeoValues){
	static_cast<QMagnetometer*>(ptr)->setReturnGeoValues(returnGeoValues != 0);
}

void* QMagnetometer_NewQMagnetometer(void* parent){
	return new QMagnetometer(static_cast<QObject*>(parent));
}

void QMagnetometer_ConnectReturnGeoValuesChanged(void* ptr){
	QObject::connect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));;
}

void QMagnetometer_DisconnectReturnGeoValuesChanged(void* ptr){
	QObject::disconnect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));;
}

void QMagnetometer_DestroyQMagnetometer(void* ptr){
	static_cast<QMagnetometer*>(ptr)->~QMagnetometer();
}

