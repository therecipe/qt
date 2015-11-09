#include "qcameralockscontrol.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraLocksControl>
#include "_cgo_export.h"

class MyQCameraLocksControl: public QCameraLocksControl {
public:
void Signal_LockStatusChanged(QCamera::LockType lock, QCamera::LockStatus status, QCamera::LockChangeReason reason){callbackQCameraLocksControlLockStatusChanged(this->objectName().toUtf8().data(), lock, status, reason);};
};

int QCameraLocksControl_LockStatus(void* ptr, int lock){
	return static_cast<QCameraLocksControl*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lock));
}

void QCameraLocksControl_ConnectLockStatusChanged(void* ptr){
	QObject::connect(static_cast<QCameraLocksControl*>(ptr), static_cast<void (QCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCameraLocksControl::lockStatusChanged), static_cast<MyQCameraLocksControl*>(ptr), static_cast<void (MyQCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCameraLocksControl::Signal_LockStatusChanged));;
}

void QCameraLocksControl_DisconnectLockStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraLocksControl*>(ptr), static_cast<void (QCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCameraLocksControl::lockStatusChanged), static_cast<MyQCameraLocksControl*>(ptr), static_cast<void (MyQCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCameraLocksControl::Signal_LockStatusChanged));;
}

void QCameraLocksControl_SearchAndLock(void* ptr, int locks){
	static_cast<QCameraLocksControl*>(ptr)->searchAndLock(static_cast<QCamera::LockType>(locks));
}

int QCameraLocksControl_SupportedLocks(void* ptr){
	return static_cast<QCameraLocksControl*>(ptr)->supportedLocks();
}

void QCameraLocksControl_Unlock(void* ptr, int locks){
	static_cast<QCameraLocksControl*>(ptr)->unlock(static_cast<QCamera::LockType>(locks));
}

void QCameraLocksControl_DestroyQCameraLocksControl(void* ptr){
	static_cast<QCameraLocksControl*>(ptr)->~QCameraLocksControl();
}

