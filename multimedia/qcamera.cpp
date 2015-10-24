#include "qcamera.h"
#include <QByteArray>
#include <QCameraInfo>
#include <QMetaObject>
#include <QAbstractVideoSurface>
#include <QCameraViewfinderSettings>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QCamera>
#include "_cgo_export.h"

class MyQCamera: public QCamera {
public:
void Signal_CaptureModeChanged(QCamera::CaptureModes mode){callbackQCameraCaptureModeChanged(this->objectName().toUtf8().data(), mode);};
void Signal_LockFailed(){callbackQCameraLockFailed(this->objectName().toUtf8().data());};
void Signal_LockStatusChanged(QCamera::LockStatus status, QCamera::LockChangeReason reason){callbackQCameraLockStatusChanged(this->objectName().toUtf8().data(), status, reason);};
void Signal_Locked(){callbackQCameraLocked(this->objectName().toUtf8().data());};
void Signal_StateChanged(QCamera::State state){callbackQCameraStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StatusChanged(QCamera::Status status){callbackQCameraStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QCamera_CaptureMode(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->captureMode();
}

void QCamera_SearchAndLock2(QtObjectPtr ptr, int locks){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_SetCaptureMode(QtObjectPtr ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "setCaptureMode", Q_ARG(QCamera::CaptureMode, static_cast<QCamera::CaptureMode>(mode)));
}

int QCamera_Status(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->status();
}

QtObjectPtr QCamera_NewQCamera4(int position, QtObjectPtr parent){
	return new QCamera(static_cast<QCamera::Position>(position), static_cast<QObject*>(parent));
}

QtObjectPtr QCamera_NewQCamera(QtObjectPtr parent){
	return new QCamera(static_cast<QObject*>(parent));
}

QtObjectPtr QCamera_NewQCamera2(QtObjectPtr deviceName, QtObjectPtr parent){
	return new QCamera(*static_cast<QByteArray*>(deviceName), static_cast<QObject*>(parent));
}

QtObjectPtr QCamera_NewQCamera3(QtObjectPtr cameraInfo, QtObjectPtr parent){
	return new QCamera(*static_cast<QCameraInfo*>(cameraInfo), static_cast<QObject*>(parent));
}

void QCamera_ConnectCaptureModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));;
}

void QCamera_DisconnectCaptureModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));;
}

int QCamera_Error(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->error();
}

char* QCamera_ErrorString(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->errorString().toUtf8().data();
}

QtObjectPtr QCamera_Exposure(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->exposure();
}

QtObjectPtr QCamera_Focus(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->focus();
}

QtObjectPtr QCamera_ImageProcessing(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->imageProcessing();
}

int QCamera_IsCaptureModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCamera*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCamera_Load(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "load");
}

void QCamera_ConnectLockFailed(QtObjectPtr ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));;
}

void QCamera_DisconnectLockFailed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));;
}

int QCamera_LockStatus(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->lockStatus();
}

int QCamera_LockStatus2(QtObjectPtr ptr, int lockType){
	return static_cast<QCamera*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lockType));
}

void QCamera_ConnectLockStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));;
}

void QCamera_DisconnectLockStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));;
}

void QCamera_ConnectLocked(QtObjectPtr ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));;
}

void QCamera_DisconnectLocked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));;
}

int QCamera_RequestedLocks(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->requestedLocks();
}

void QCamera_SearchAndLock(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock");
}

void QCamera_SetViewfinder3(QtObjectPtr ptr, QtObjectPtr surface){
	static_cast<QCamera*>(ptr)->setViewfinder(static_cast<QAbstractVideoSurface*>(surface));
}

void QCamera_SetViewfinderSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QCamera*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCamera_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "start");
}

void QCamera_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));;
}

void QCamera_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));;
}

void QCamera_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));;
}

void QCamera_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));;
}

void QCamera_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "stop");
}

int QCamera_SupportedLocks(QtObjectPtr ptr){
	return static_cast<QCamera*>(ptr)->supportedLocks();
}

void QCamera_Unload(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unload");
}

void QCamera_Unlock(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock");
}

void QCamera_Unlock2(QtObjectPtr ptr, int locks){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_DestroyQCamera(QtObjectPtr ptr){
	static_cast<QCamera*>(ptr)->~QCamera();
}

