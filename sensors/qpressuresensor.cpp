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

QtObjectPtr QPressureSensor_Reading(QtObjectPtr ptr){
	return static_cast<QPressureSensor*>(ptr)->reading();
}

QtObjectPtr QPressureSensor_NewQPressureSensor(QtObjectPtr parent){
	return new QPressureSensor(static_cast<QObject*>(parent));
}

void QPressureSensor_DestroyQPressureSensor(QtObjectPtr ptr){
	static_cast<QPressureSensor*>(ptr)->~QPressureSensor();
}

