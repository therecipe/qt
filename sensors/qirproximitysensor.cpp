#include "qirproximitysensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIRProximitySensor>
#include "_cgo_export.h"

class MyQIRProximitySensor: public QIRProximitySensor {
public:
};

QtObjectPtr QIRProximitySensor_Reading(QtObjectPtr ptr){
	return static_cast<QIRProximitySensor*>(ptr)->reading();
}

QtObjectPtr QIRProximitySensor_NewQIRProximitySensor(QtObjectPtr parent){
	return new QIRProximitySensor(static_cast<QObject*>(parent));
}

void QIRProximitySensor_DestroyQIRProximitySensor(QtObjectPtr ptr){
	static_cast<QIRProximitySensor*>(ptr)->~QIRProximitySensor();
}

