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

QtObjectPtr QProximitySensor_Reading(QtObjectPtr ptr){
	return static_cast<QProximitySensor*>(ptr)->reading();
}

QtObjectPtr QProximitySensor_NewQProximitySensor(QtObjectPtr parent){
	return new QProximitySensor(static_cast<QObject*>(parent));
}

void QProximitySensor_DestroyQProximitySensor(QtObjectPtr ptr){
	static_cast<QProximitySensor*>(ptr)->~QProximitySensor();
}

