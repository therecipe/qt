#include "qirproximitysensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QIRProximitySensor>
#include "_cgo_export.h"

class MyQIRProximitySensor: public QIRProximitySensor {
public:
};

void* QIRProximitySensor_Reading(void* ptr){
	return static_cast<QIRProximitySensor*>(ptr)->reading();
}

void* QIRProximitySensor_NewQIRProximitySensor(void* parent){
	return new QIRProximitySensor(static_cast<QObject*>(parent));
}

void QIRProximitySensor_DestroyQIRProximitySensor(void* ptr){
	static_cast<QIRProximitySensor*>(ptr)->~QIRProximitySensor();
}

