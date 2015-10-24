#include "qtiltsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QTiltSensor>
#include "_cgo_export.h"

class MyQTiltSensor: public QTiltSensor {
public:
};

QtObjectPtr QTiltSensor_NewQTiltSensor(QtObjectPtr parent){
	return new QTiltSensor(static_cast<QObject*>(parent));
}

QtObjectPtr QTiltSensor_Reading(QtObjectPtr ptr){
	return static_cast<QTiltSensor*>(ptr)->reading();
}

void QTiltSensor_DestroyQTiltSensor(QtObjectPtr ptr){
	static_cast<QTiltSensor*>(ptr)->~QTiltSensor();
}

void QTiltSensor_Calibrate(QtObjectPtr ptr){
	static_cast<QTiltSensor*>(ptr)->calibrate();
}

