#include "qorientationsensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QOrientationSensor>
#include "_cgo_export.h"

class MyQOrientationSensor: public QOrientationSensor {
public:
};

void* QOrientationSensor_Reading(void* ptr){
	return static_cast<QOrientationSensor*>(ptr)->reading();
}

void* QOrientationSensor_NewQOrientationSensor(void* parent){
	return new QOrientationSensor(static_cast<QObject*>(parent));
}

void QOrientationSensor_DestroyQOrientationSensor(void* ptr){
	static_cast<QOrientationSensor*>(ptr)->~QOrientationSensor();
}

