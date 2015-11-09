#include "qholstersensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHolsterSensor>
#include "_cgo_export.h"

class MyQHolsterSensor: public QHolsterSensor {
public:
};

void* QHolsterSensor_Reading(void* ptr){
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

void* QHolsterSensor_NewQHolsterSensor(void* parent){
	return new QHolsterSensor(static_cast<QObject*>(parent));
}

void QHolsterSensor_DestroyQHolsterSensor(void* ptr){
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

