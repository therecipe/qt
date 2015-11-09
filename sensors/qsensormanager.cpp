#include "qsensormanager.h"
#include <QSensorBackendFactory>
#include <QByteArray>
#include <QSensorBackend>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorManager>
#include "_cgo_export.h"

class MyQSensorManager: public QSensorManager {
public:
};

void* QSensorManager_QSensorManager_CreateBackend(void* sensor){
	return QSensorManager::createBackend(static_cast<QSensor*>(sensor));
}

int QSensorManager_QSensorManager_IsBackendRegistered(void* ty, void* identifier){
	return QSensorManager::isBackendRegistered(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_RegisterBackend(void* ty, void* identifier, void* factory){
	QSensorManager::registerBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier), static_cast<QSensorBackendFactory*>(factory));
}

void QSensorManager_QSensorManager_SetDefaultBackend(void* ty, void* identifier){
	QSensorManager::setDefaultBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_UnregisterBackend(void* ty, void* identifier){
	QSensorManager::unregisterBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

