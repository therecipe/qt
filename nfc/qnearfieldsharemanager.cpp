#include "qnearfieldsharemanager.h"
#include <QNearFieldShareTarget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNearFieldShareManager>
#include "_cgo_export.h"

class MyQNearFieldShareManager: public QNearFieldShareManager {
public:
void Signal_Error(QNearFieldShareManager::ShareError error){callbackQNearFieldShareManagerError(this->objectName().toUtf8().data(), error);};
void Signal_ShareModesChanged(QNearFieldShareManager::ShareModes modes){callbackQNearFieldShareManagerShareModesChanged(this->objectName().toUtf8().data(), modes);};
void Signal_TargetDetected(QNearFieldShareTarget * shareTarget){callbackQNearFieldShareManagerTargetDetected(this->objectName().toUtf8().data(), shareTarget);};
};

QtObjectPtr QNearFieldShareManager_NewQNearFieldShareManager(QtObjectPtr parent){
	return new QNearFieldShareManager(static_cast<QObject*>(parent));
}

void QNearFieldShareManager_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));;
}

void QNearFieldShareManager_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));;
}

void QNearFieldShareManager_SetShareModes(QtObjectPtr ptr, int mode){
	static_cast<QNearFieldShareManager*>(ptr)->setShareModes(static_cast<QNearFieldShareManager::ShareMode>(mode));
}

int QNearFieldShareManager_ShareError(QtObjectPtr ptr){
	return static_cast<QNearFieldShareManager*>(ptr)->shareError();
}

int QNearFieldShareManager_ShareModes(QtObjectPtr ptr){
	return static_cast<QNearFieldShareManager*>(ptr)->shareModes();
}

void QNearFieldShareManager_ConnectShareModesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));;
}

void QNearFieldShareManager_DisconnectShareModesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));;
}

int QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes(){
	return QNearFieldShareManager::supportedShareModes();
}

void QNearFieldShareManager_ConnectTargetDetected(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));;
}

void QNearFieldShareManager_DisconnectTargetDetected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));;
}

void QNearFieldShareManager_DestroyQNearFieldShareManager(QtObjectPtr ptr){
	static_cast<QNearFieldShareManager*>(ptr)->~QNearFieldShareManager();
}

