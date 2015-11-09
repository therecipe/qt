#include "qsensorplugininterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorPluginInterface>
#include "_cgo_export.h"

class MyQSensorPluginInterface: public QSensorPluginInterface {
public:
};

void QSensorPluginInterface_RegisterSensors(void* ptr){
	static_cast<QSensorPluginInterface*>(ptr)->registerSensors();
}

