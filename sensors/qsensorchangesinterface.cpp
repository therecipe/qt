#include "qsensorchangesinterface.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QString>
#include <QSensorChangesInterface>
#include "_cgo_export.h"

class MyQSensorChangesInterface: public QSensorChangesInterface {
public:
};

void QSensorChangesInterface_SensorsChanged(void* ptr){
	static_cast<QSensorChangesInterface*>(ptr)->sensorsChanged();
}

