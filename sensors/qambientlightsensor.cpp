#include "qambientlightsensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientLightSensor>
#include "_cgo_export.h"

class MyQAmbientLightSensor: public QAmbientLightSensor {
public:
};

QtObjectPtr QAmbientLightSensor_Reading(QtObjectPtr ptr){
	return static_cast<QAmbientLightSensor*>(ptr)->reading();
}

QtObjectPtr QAmbientLightSensor_NewQAmbientLightSensor(QtObjectPtr parent){
	return new QAmbientLightSensor(static_cast<QObject*>(parent));
}

void QAmbientLightSensor_DestroyQAmbientLightSensor(QtObjectPtr ptr){
	static_cast<QAmbientLightSensor*>(ptr)->~QAmbientLightSensor();
}

