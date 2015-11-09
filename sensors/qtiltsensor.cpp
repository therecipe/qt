#include "qtiltsensor.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTiltSensor>
#include "_cgo_export.h"

class MyQTiltSensor: public QTiltSensor {
public:
};

void* QTiltSensor_NewQTiltSensor(void* parent){
	return new QTiltSensor(static_cast<QObject*>(parent));
}

void* QTiltSensor_Reading(void* ptr){
	return static_cast<QTiltSensor*>(ptr)->reading();
}

void QTiltSensor_DestroyQTiltSensor(void* ptr){
	static_cast<QTiltSensor*>(ptr)->~QTiltSensor();
}

void QTiltSensor_Calibrate(void* ptr){
	static_cast<QTiltSensor*>(ptr)->calibrate();
}

