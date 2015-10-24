#include "qsensorbackend.h"
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorBackend>
#include "_cgo_export.h"

class MyQSensorBackend: public QSensorBackend {
public:
};

int QSensorBackend_IsFeatureSupported(QtObjectPtr ptr, int feature){
	return static_cast<QSensorBackend*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensorBackend_SensorBusy(QtObjectPtr ptr){
	static_cast<QSensorBackend*>(ptr)->sensorBusy();
}

void QSensorBackend_SensorError(QtObjectPtr ptr, int error){
	static_cast<QSensorBackend*>(ptr)->sensorError(error);
}

void QSensorBackend_NewReadingAvailable(QtObjectPtr ptr){
	static_cast<QSensorBackend*>(ptr)->newReadingAvailable();
}

QtObjectPtr QSensorBackend_Reading(QtObjectPtr ptr){
	return static_cast<QSensorBackend*>(ptr)->reading();
}

QtObjectPtr QSensorBackend_Sensor(QtObjectPtr ptr){
	return static_cast<QSensorBackend*>(ptr)->sensor();
}

void QSensorBackend_SensorStopped(QtObjectPtr ptr){
	static_cast<QSensorBackend*>(ptr)->sensorStopped();
}

void QSensorBackend_SetDataRates(QtObjectPtr ptr, QtObjectPtr otherSensor){
	static_cast<QSensorBackend*>(ptr)->setDataRates(static_cast<QSensor*>(otherSensor));
}

void QSensorBackend_SetDescription(QtObjectPtr ptr, char* description){
	static_cast<QSensorBackend*>(ptr)->setDescription(QString(description));
}

void QSensorBackend_Start(QtObjectPtr ptr){
	static_cast<QSensorBackend*>(ptr)->start();
}

void QSensorBackend_Stop(QtObjectPtr ptr){
	static_cast<QSensorBackend*>(ptr)->stop();
}

