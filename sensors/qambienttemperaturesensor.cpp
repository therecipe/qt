#include "qambienttemperaturesensor.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QAmbientTemperatureSensor>
#include "_cgo_export.h"

class MyQAmbientTemperatureSensor: public QAmbientTemperatureSensor {
public:
};

QtObjectPtr QAmbientTemperatureSensor_Reading(QtObjectPtr ptr){
	return static_cast<QAmbientTemperatureSensor*>(ptr)->reading();
}

QtObjectPtr QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(QtObjectPtr parent){
	return new QAmbientTemperatureSensor(static_cast<QObject*>(parent));
}

void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(QtObjectPtr ptr){
	static_cast<QAmbientTemperatureSensor*>(ptr)->~QAmbientTemperatureSensor();
}

