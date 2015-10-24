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

QtObjectPtr QMagnetometer_Reading(QtObjectPtr ptr){
	return static_cast<QMagnetometer*>(ptr)->reading();
}

int QMagnetometer_ReturnGeoValues(QtObjectPtr ptr){
	return static_cast<QMagnetometer*>(ptr)->returnGeoValues();
}

void QMagnetometer_SetReturnGeoValues(QtObjectPtr ptr, int returnGeoValues){
	static_cast<QMagnetometer*>(ptr)->setReturnGeoValues(returnGeoValues != 0);
}

QtObjectPtr QMagnetometer_NewQMagnetometer(QtObjectPtr parent){
	return new QMagnetometer(static_cast<QObject*>(parent));
}

void QMagnetometer_ConnectReturnGeoValuesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));;
}

void QMagnetometer_DisconnectReturnGeoValuesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));;
}

void QMagnetometer_DestroyQMagnetometer(QtObjectPtr ptr){
	static_cast<QMagnetometer*>(ptr)->~QMagnetometer();
}

