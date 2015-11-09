#include "qcamera.h"
#include <QVariant>
#include <QUrl>
#include <QByteArray>
#include <QString>
#include <QModelIndex>
#include <QMetaObject>
#include <QCameraInfo>
#include <QObject>
#include <QAbstractVideoSurface>
#include <QCameraViewfinderSettings>
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

int QCamera_CaptureMode(void* ptr){
	return static_cast<QCamera*>(ptr)->captureMode();
}

void QCamera_SearchAndLock2(void* ptr, int locks){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_SetCaptureMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "setCaptureMode", Q_ARG(QCamera::CaptureMode, static_cast<QCamera::CaptureMode>(mode)));
}

int QCamera_Status(void* ptr){
	return static_cast<QCamera*>(ptr)->status();
}

void* QCamera_NewQCamera4(int position, void* parent){
	return new QCamera(static_cast<QCamera::Position>(position), static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera(void* parent){
	return new QCamera(static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera2(void* deviceName, void* parent){
	return new QCamera(*static_cast<QByteArray*>(deviceName), static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera3(void* cameraInfo, void* parent){
	return new QCamera(*static_cast<QCameraInfo*>(cameraInfo), static_cast<QObject*>(parent));
}

void QCamera_ConnectCaptureModeChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));;
}

void QCamera_DisconnectCaptureModeChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));;
}

int QCamera_Error(void* ptr){
	return static_cast<QCamera*>(ptr)->error();
}

char* QCamera_ErrorString(void* ptr){
	return static_cast<QCamera*>(ptr)->errorString().toUtf8().data();
}

void* QCamera_Exposure(void* ptr){
	return static_cast<QCamera*>(ptr)->exposure();
}

void* QCamera_Focus(void* ptr){
	return static_cast<QCamera*>(ptr)->focus();
}

void* QCamera_ImageProcessing(void* ptr){
	return static_cast<QCamera*>(ptr)->imageProcessing();
}

int QCamera_IsCaptureModeSupported(void* ptr, int mode){
	return static_cast<QCamera*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCamera_Load(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "load");
}

void QCamera_ConnectLockFailed(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));;
}

void QCamera_DisconnectLockFailed(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));;
}

int QCamera_LockStatus(void* ptr){
	return static_cast<QCamera*>(ptr)->lockStatus();
}

int QCamera_LockStatus2(void* ptr, int lockType){
	return static_cast<QCamera*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lockType));
}

void QCamera_ConnectLockStatusChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));;
}

void QCamera_DisconnectLockStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));;
}

void QCamera_ConnectLocked(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));;
}

void QCamera_DisconnectLocked(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));;
}

int QCamera_RequestedLocks(void* ptr){
	return static_cast<QCamera*>(ptr)->requestedLocks();
}

void QCamera_SearchAndLock(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock");
}

void QCamera_SetViewfinder3(void* ptr, void* surface){
	static_cast<QCamera*>(ptr)->setViewfinder(static_cast<QAbstractVideoSurface*>(surface));
}

void QCamera_SetViewfinderSettings(void* ptr, void* settings){
	static_cast<QCamera*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCamera_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "start");
}

void QCamera_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));;
}

void QCamera_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));;
}

void QCamera_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));;
}

void QCamera_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));;
}

void QCamera_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "stop");
}

int QCamera_SupportedLocks(void* ptr){
	return static_cast<QCamera*>(ptr)->supportedLocks();
}

void QCamera_Unload(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unload");
}

void QCamera_Unlock(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock");
}

void QCamera_Unlock2(void* ptr, int locks){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_DestroyQCamera(void* ptr){
	static_cast<QCamera*>(ptr)->~QCamera();
}

