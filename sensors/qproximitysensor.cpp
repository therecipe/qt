#include "qproximitysensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QProximitySensor>
#include "_cgo_export.h"

class MyQProximitySensor: public QProximitySensor {
public:
};

void* QProximitySensor_Reading(void* ptr){
	return static_cast<QProximitySensor*>(ptr)->reading();
}

void* QProximitySensor_NewQProximitySensor(void* parent){
	return new QProximitySensor(static_cast<QObject*>(parent));
}

void QProximitySensor_DestroyQProximitySensor(void* ptr){
	static_cast<QProximitySensor*>(ptr)->~QProximitySensor();
}

