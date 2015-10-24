#include "qlightsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QLightSensor>
#include "_cgo_export.h"

class MyQLightSensor: public QLightSensor {
public:
};

QtObjectPtr QLightSensor_Reading(QtObjectPtr ptr){
	return static_cast<QLightSensor*>(ptr)->reading();
}

QtObjectPtr QLightSensor_NewQLightSensor(QtObjectPtr parent){
	return new QLightSensor(static_cast<QObject*>(parent));
}

void QLightSensor_DestroyQLightSensor(QtObjectPtr ptr){
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

