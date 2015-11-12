#include "qambienttemperaturesensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QAmbientTemperatureSensor>
#include "_cgo_export.h"

class MyQAmbientTemperatureSensor: public QAmbientTemperatureSensor {
public:
};

void* QAmbientTemperatureSensor_Reading(void* ptr){
	return static_cast<QAmbientTemperatureSensor*>(ptr)->reading();
}

void* QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(void* parent){
	return new QAmbientTemperatureSensor(static_cast<QObject*>(parent));
}

void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(void* ptr){
	static_cast<QAmbientTemperatureSensor*>(ptr)->~QAmbientTemperatureSensor();
}

