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

void* QDistanceSensor_Reading(void* ptr){
	return static_cast<QDistanceSensor*>(ptr)->reading();
}

void* QDistanceSensor_NewQDistanceSensor(void* parent){
	return new QDistanceSensor(static_cast<QObject*>(parent));
}

void QDistanceSensor_DestroyQDistanceSensor(void* ptr){
	static_cast<QDistanceSensor*>(ptr)->~QDistanceSensor();
}

