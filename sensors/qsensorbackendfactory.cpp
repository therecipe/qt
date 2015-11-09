#include "qsensorbackendfactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorBackend>
#include <QSensorBackendFactory>
#include "_cgo_export.h"

class MyQSensorBackendFactory: public QSensorBackendFactory {
public:
};

void* QSensorBackendFactory_CreateBackend(void* ptr, void* sensor){
	return static_cast<QSensorBackendFactory*>(ptr)->createBackend(static_cast<QSensor*>(sensor));
}

