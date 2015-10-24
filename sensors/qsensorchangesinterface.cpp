#include "qsensorchangesinterface.h"
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorChangesInterface>
#include "_cgo_export.h"

class MyQSensorChangesInterface: public QSensorChangesInterface {
public:
};

void QSensorChangesInterface_SensorsChanged(QtObjectPtr ptr){
	static_cast<QSensorChangesInterface*>(ptr)->sensorsChanged();
}

