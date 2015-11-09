#include "qpressuresensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QPressureSensor>
#include "_cgo_export.h"

class MyQPressureSensor: public QPressureSensor {
public:
};

void* QPressureSensor_Reading(void* ptr){
	return static_cast<QPressureSensor*>(ptr)->reading();
}

void* QPressureSensor_NewQPressureSensor(void* parent){
	return new QPressureSensor(static_cast<QObject*>(parent));
}

void QPressureSensor_DestroyQPressureSensor(void* ptr){
	static_cast<QPressureSensor*>(ptr)->~QPressureSensor();
}

