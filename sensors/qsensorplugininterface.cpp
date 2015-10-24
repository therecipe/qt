#include "qsensorplugininterface.h"
#include <QModelIndex>
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensorPluginInterface>
#include "_cgo_export.h"

class MyQSensorPluginInterface: public QSensorPluginInterface {
public:
};

void QSensorPluginInterface_RegisterSensors(QtObjectPtr ptr){
	static_cast<QSensorPluginInterface*>(ptr)->registerSensors();
}

