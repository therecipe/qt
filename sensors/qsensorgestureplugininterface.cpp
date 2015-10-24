#include "qsensorgestureplugininterface.h"
#include <QModelIndex>
#include <QSensor>
#include <QSensorGesture>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensorGesturePluginInterface>
#include "_cgo_export.h"

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface {
public:
};

char* QSensorGesturePluginInterface_Name(QtObjectPtr ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->name().toUtf8().data();
}

char* QSensorGesturePluginInterface_SupportedIds(QtObjectPtr ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->supportedIds().join("|").toUtf8().data();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(QtObjectPtr ptr){
	static_cast<QSensorGesturePluginInterface*>(ptr)->~QSensorGesturePluginInterface();
}

