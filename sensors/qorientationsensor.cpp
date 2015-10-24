#include "qorientationsensor.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOrientationSensor>
#include "_cgo_export.h"

class MyQOrientationSensor: public QOrientationSensor {
public:
};

QtObjectPtr QOrientationSensor_Reading(QtObjectPtr ptr){
	return static_cast<QOrientationSensor*>(ptr)->reading();
}

QtObjectPtr QOrientationSensor_NewQOrientationSensor(QtObjectPtr parent){
	return new QOrientationSensor(static_cast<QObject*>(parent));
}

void QOrientationSensor_DestroyQOrientationSensor(QtObjectPtr ptr){
	static_cast<QOrientationSensor*>(ptr)->~QOrientationSensor();
}

