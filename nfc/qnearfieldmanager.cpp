#include "qnearfieldmanager.h"
#include <QObject>
#include <QNdefRecord>
#include <QNdefFilter>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QByteArray>
#include <QModelIndex>
#include <QNearFieldTarget>
#include <QNearFieldManager>
#include "_cgo_export.h"

class MyQNearFieldManager: public QNearFieldManager {
public:
void Signal_TargetDetected(QNearFieldTarget * target){callbackQNearFieldManagerTargetDetected(this->objectName().toUtf8().data(), target);};
void Signal_TargetLost(QNearFieldTarget * target){callbackQNearFieldManagerTargetLost(this->objectName().toUtf8().data(), target);};
};

int QNearFieldManager_RegisterNdefMessageHandler(void* ptr, void* object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_StartTargetDetection(void* ptr){
	return static_cast<QNearFieldManager*>(ptr)->startTargetDetection();
}

void* QNearFieldManager_NewQNearFieldManager(void* parent){
	return new QNearFieldManager(static_cast<QObject*>(parent));
}

int QNearFieldManager_IsAvailable(void* ptr){
	return static_cast<QNearFieldManager*>(ptr)->isAvailable();
}

int QNearFieldManager_RegisterNdefMessageHandler2(void* ptr, int typeNameFormat, void* ty, void* object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_RegisterNdefMessageHandler3(void* ptr, void* filter, void* object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(*static_cast<QNdefFilter*>(filter), static_cast<QObject*>(object), const_cast<const char*>(method));
}

void QNearFieldManager_SetTargetAccessModes(void* ptr, int accessModes){
	static_cast<QNearFieldManager*>(ptr)->setTargetAccessModes(static_cast<QNearFieldManager::TargetAccessMode>(accessModes));
}

void QNearFieldManager_StopTargetDetection(void* ptr){
	static_cast<QNearFieldManager*>(ptr)->stopTargetDetection();
}

int QNearFieldManager_TargetAccessModes(void* ptr){
	return static_cast<QNearFieldManager*>(ptr)->targetAccessModes();
}

void QNearFieldManager_ConnectTargetDetected(void* ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));;
}

void QNearFieldManager_DisconnectTargetDetected(void* ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));;
}

void QNearFieldManager_ConnectTargetLost(void* ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

void QNearFieldManager_DisconnectTargetLost(void* ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

int QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId){
	return static_cast<QNearFieldManager*>(ptr)->unregisterNdefMessageHandler(handlerId);
}

void QNearFieldManager_DestroyQNearFieldManager(void* ptr){
	static_cast<QNearFieldManager*>(ptr)->~QNearFieldManager();
}

