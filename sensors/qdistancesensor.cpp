#include "qdistancesensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDistanceSensor>
#include "_cgo_export.h"

class MyQDistanceSensor: public QDistanceSensor {
public:
};

QtObjectPtr QDistanceSensor_Reading(QtObjectPtr ptr){
	return static_cast<QDistanceSensor*>(ptr)->reading();
}

QtObjectPtr QDistanceSensor_NewQDistanceSensor(QtObjectPtr parent){
	return new QDistanceSensor(static_cast<QObject*>(parent));
}

void QDistanceSensor_DestroyQDistanceSensor(QtObjectPtr ptr){
	static_cast<QDistanceSensor*>(ptr)->~QDistanceSensor();
}

