#include "qnearfieldsharetarget.h"
#include <QUrl>
#include <QModelIndex>
#include <QNdefMessage>
#include <QObject>
#include <QNearFieldShareManager>
#include <QString>
#include <QVariant>
#include <QNearFieldShareTarget>
#include "_cgo_export.h"

class MyQNearFieldShareTarget: public QNearFieldShareTarget {
public:
void Signal_Error(QNearFieldShareManager::ShareError error){callbackQNearFieldShareTargetError(this->objectName().toUtf8().data(), error);};
void Signal_ShareFinished(){callbackQNearFieldShareTargetShareFinished(this->objectName().toUtf8().data());};
};

void QNearFieldShareTarget_Cancel(QtObjectPtr ptr){
	static_cast<QNearFieldShareTarget*>(ptr)->cancel();
}

void QNearFieldShareTarget_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));;
}

void QNearFieldShareTarget_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));;
}

int QNearFieldShareTarget_IsShareInProgress(QtObjectPtr ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->isShareInProgress();
}

int QNearFieldShareTarget_Share(QtObjectPtr ptr, QtObjectPtr message){
	return static_cast<QNearFieldShareTarget*>(ptr)->share(*static_cast<QNdefMessage*>(message));
}

int QNearFieldShareTarget_ShareError(QtObjectPtr ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->shareError();
}

void QNearFieldShareTarget_ConnectShareFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));;
}

void QNearFieldShareTarget_DisconnectShareFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));;
}

int QNearFieldShareTarget_ShareModes(QtObjectPtr ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->shareModes();
}

void QNearFieldShareTarget_DestroyQNearFieldShareTarget(QtObjectPtr ptr){
	static_cast<QNearFieldShareTarget*>(ptr)->~QNearFieldShareTarget();
}

