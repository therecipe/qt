#include "qnearfieldmanager.h"
#include <QString>
#include <QModelIndex>
#include <QByteArray>
#include <QNearFieldTarget>
#include <QNdefFilter>
#include <QVariant>
#include <QUrl>
#include <QNdefRecord>
#include <QObject>
#include <QNearFieldManager>
#include "_cgo_export.h"

class MyQNearFieldManager: public QNearFieldManager {
public:
void Signal_TargetDetected(QNearFieldTarget * target){callbackQNearFieldManagerTargetDetected(this->objectName().toUtf8().data(), target);};
void Signal_TargetLost(QNearFieldTarget * target){callbackQNearFieldManagerTargetLost(this->objectName().toUtf8().data(), target);};
};

int QNearFieldManager_RegisterNdefMessageHandler(QtObjectPtr ptr, QtObjectPtr object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_StartTargetDetection(QtObjectPtr ptr){
	return static_cast<QNearFieldManager*>(ptr)->startTargetDetection();
}

QtObjectPtr QNearFieldManager_NewQNearFieldManager(QtObjectPtr parent){
	return new QNearFieldManager(static_cast<QObject*>(parent));
}

int QNearFieldManager_IsAvailable(QtObjectPtr ptr){
	return static_cast<QNearFieldManager*>(ptr)->isAvailable();
}

int QNearFieldManager_RegisterNdefMessageHandler2(QtObjectPtr ptr, int typeNameFormat, QtObjectPtr ty, QtObjectPtr object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_RegisterNdefMessageHandler3(QtObjectPtr ptr, QtObjectPtr filter, QtObjectPtr object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(*static_cast<QNdefFilter*>(filter), static_cast<QObject*>(object), const_cast<const char*>(method));
}

void QNearFieldManager_SetTargetAccessModes(QtObjectPtr ptr, int accessModes){
	static_cast<QNearFieldManager*>(ptr)->setTargetAccessModes(static_cast<QNearFieldManager::TargetAccessMode>(accessModes));
}

void QNearFieldManager_StopTargetDetection(QtObjectPtr ptr){
	static_cast<QNearFieldManager*>(ptr)->stopTargetDetection();
}

int QNearFieldManager_TargetAccessModes(QtObjectPtr ptr){
	return static_cast<QNearFieldManager*>(ptr)->targetAccessModes();
}

void QNearFieldManager_ConnectTargetDetected(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));;
}

void QNearFieldManager_DisconnectTargetDetected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));;
}

void QNearFieldManager_ConnectTargetLost(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

void QNearFieldManager_DisconnectTargetLost(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

int QNearFieldManager_UnregisterNdefMessageHandler(QtObjectPtr ptr, int handlerId){
	return static_cast<QNearFieldManager*>(ptr)->unregisterNdefMessageHandler(handlerId);
}

void QNearFieldManager_DestroyQNearFieldManager(QtObjectPtr ptr){
	static_cast<QNearFieldManager*>(ptr)->~QNearFieldManager();
}

