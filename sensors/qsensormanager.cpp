#include "qsensormanager.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorBackendFactory>
#include <QByteArray>
#include <QSensorBackend>
#include <QSensor>
#include <QSensorManager>
#include "_cgo_export.h"

class MyQSensorManager: public QSensorManager {
public:
};

QtObjectPtr QSensorManager_QSensorManager_CreateBackend(QtObjectPtr sensor){
	return QSensorManager::createBackend(static_cast<QSensor*>(sensor));
}

int QSensorManager_QSensorManager_IsBackendRegistered(QtObjectPtr ty, QtObjectPtr identifier){
	return QSensorManager::isBackendRegistered(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_RegisterBackend(QtObjectPtr ty, QtObjectPtr identifier, QtObjectPtr factory){
	QSensorManager::registerBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier), static_cast<QSensorBackendFactory*>(factory));
}

void QSensorManager_QSensorManager_SetDefaultBackend(QtObjectPtr ty, QtObjectPtr identifier){
	QSensorManager::setDefaultBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_UnregisterBackend(QtObjectPtr ty, QtObjectPtr identifier){
	QSensorManager::unregisterBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

