#include "qambientlightsensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QAmbientLightSensor>
#include "_cgo_export.h"

class MyQAmbientLightSensor: public QAmbientLightSensor {
public:
};

void* QAmbientLightSensor_Reading(void* ptr){
	return static_cast<QAmbientLightSensor*>(ptr)->reading();
}

void* QAmbientLightSensor_NewQAmbientLightSensor(void* parent){
	return new QAmbientLightSensor(static_cast<QObject*>(parent));
}

void QAmbientLightSensor_DestroyQAmbientLightSensor(void* ptr){
	static_cast<QAmbientLightSensor*>(ptr)->~QAmbientLightSensor();
}

