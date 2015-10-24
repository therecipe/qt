#include "qholstersensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHolsterSensor>
#include "_cgo_export.h"

class MyQHolsterSensor: public QHolsterSensor {
public:
};

QtObjectPtr QHolsterSensor_Reading(QtObjectPtr ptr){
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

QtObjectPtr QHolsterSensor_NewQHolsterSensor(QtObjectPtr parent){
	return new QHolsterSensor(static_cast<QObject*>(parent));
}

void QHolsterSensor_DestroyQHolsterSensor(QtObjectPtr ptr){
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

