#include "qnearfieldsharemanager.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNearFieldShareTarget>
#include <QString>
#include <QVariant>
#include <QNearFieldShareManager>
#include "_cgo_export.h"

class MyQNearFieldShareManager: public QNearFieldShareManager {
public:
void Signal_Error(QNearFieldShareManager::ShareError error){callbackQNearFieldShareManagerError(this->objectName().toUtf8().data(), error);};
void Signal_ShareModesChanged(QNearFieldShareManager::ShareModes modes){callbackQNearFieldShareManagerShareModesChanged(this->objectName().toUtf8().data(), modes);};
void Signal_TargetDetected(QNearFieldShareTarget * shareTarget){callbackQNearFieldShareManagerTargetDetected(this->objectName().toUtf8().data(), shareTarget);};
};

void* QNearFieldShareManager_NewQNearFieldShareManager(void* parent){
	return new QNearFieldShareManager(static_cast<QObject*>(parent));
}

void QNearFieldShareManager_ConnectError(void* ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));;
}

void QNearFieldShareManager_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));;
}

void QNearFieldShareManager_SetShareModes(void* ptr, int mode){
	static_cast<QNearFieldShareManager*>(ptr)->setShareModes(static_cast<QNearFieldShareManager::ShareMode>(mode));
}

int QNearFieldShareManager_ShareError(void* ptr){
	return static_cast<QNearFieldShareManager*>(ptr)->shareError();
}

int QNearFieldShareManager_ShareModes(void* ptr){
	return static_cast<QNearFieldShareManager*>(ptr)->shareModes();
}

void QNearFieldShareManager_ConnectShareModesChanged(void* ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));;
}

void QNearFieldShareManager_DisconnectShareModesChanged(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));;
}

int QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes(){
	return QNearFieldShareManager::supportedShareModes();
}

void QNearFieldShareManager_ConnectTargetDetected(void* ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));;
}

void QNearFieldShareManager_DisconnectTargetDetected(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));;
}

void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr){
	static_cast<QNearFieldShareManager*>(ptr)->~QNearFieldShareManager();
}

