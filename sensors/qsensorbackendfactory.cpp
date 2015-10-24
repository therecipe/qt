#include "qsensorbackendfactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorBackend>
#include <QSensor>
#include <QSensorBackendFactory>
#include "_cgo_export.h"

class MyQSensorBackendFactory: public QSensorBackendFactory {
public:
};

QtObjectPtr QSensorBackendFactory_CreateBackend(QtObjectPtr ptr, QtObjectPtr sensor){
	return static_cast<QSensorBackendFactory*>(ptr)->createBackend(static_cast<QSensor*>(sensor));
}

