#include "qsensorgestureplugininterface.h"
#include <QSensorGesture>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorGesturePluginInterface>
#include "_cgo_export.h"

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface {
public:
};

char* QSensorGesturePluginInterface_Name(void* ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->name().toUtf8().data();
}

char* QSensorGesturePluginInterface_SupportedIds(void* ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->supportedIds().join("|").toUtf8().data();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(void* ptr){
	static_cast<QSensorGesturePluginInterface*>(ptr)->~QSensorGesturePluginInterface();
}

