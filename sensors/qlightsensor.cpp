#include "qlightsensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QLightSensor>
#include "_cgo_export.h"

class MyQLightSensor: public QLightSensor {
public:
};

double QLightSensor_FieldOfView(void* ptr){
	return static_cast<double>(static_cast<QLightSensor*>(ptr)->fieldOfView());
}

void* QLightSensor_Reading(void* ptr){
	return static_cast<QLightSensor*>(ptr)->reading();
}

void* QLightSensor_NewQLightSensor(void* parent){
	return new QLightSensor(static_cast<QObject*>(parent));
}

void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView){
	static_cast<QLightSensor*>(ptr)->setFieldOfView(static_cast<qreal>(fieldOfView));
}

void QLightSensor_DestroyQLightSensor(void* ptr){
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

