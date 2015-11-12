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

void QSensorBackend_AddDataRate(void* ptr, double min, double max){
	static_cast<QSensorBackend*>(ptr)->addDataRate(static_cast<qreal>(min), static_cast<qreal>(max));
}

int QSensorBackend_IsFeatureSupported(void* ptr, int feature){
	return static_cast<QSensorBackend*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensorBackend_SensorBusy(void* ptr){
	static_cast<QSensorBackend*>(ptr)->sensorBusy();
}

void QSensorBackend_SensorError(void* ptr, int error){
	static_cast<QSensorBackend*>(ptr)->sensorError(error);
}

void QSensorBackend_AddOutputRange(void* ptr, double min, double max, double accuracy){
	static_cast<QSensorBackend*>(ptr)->addOutputRange(static_cast<qreal>(min), static_cast<qreal>(max), static_cast<qreal>(accuracy));
}

void QSensorBackend_NewReadingAvailable(void* ptr){
	static_cast<QSensorBackend*>(ptr)->newReadingAvailable();
}

void* QSensorBackend_Reading(void* ptr){
	return static_cast<QSensorBackend*>(ptr)->reading();
}

void* QSensorBackend_Sensor(void* ptr){
	return static_cast<QSensorBackend*>(ptr)->sensor();
}

void QSensorBackend_SensorStopped(void* ptr){
	static_cast<QSensorBackend*>(ptr)->sensorStopped();
}

void QSensorBackend_SetDataRates(void* ptr, void* otherSensor){
	static_cast<QSensorBackend*>(ptr)->setDataRates(static_cast<QSensor*>(otherSensor));
}

void QSensorBackend_SetDescription(void* ptr, char* description){
	static_cast<QSensorBackend*>(ptr)->setDescription(QString(description));
}

void QSensorBackend_Start(void* ptr){
	static_cast<QSensorBackend*>(ptr)->start();
}

void QSensorBackend_Stop(void* ptr){
	static_cast<QSensorBackend*>(ptr)->stop();
}

