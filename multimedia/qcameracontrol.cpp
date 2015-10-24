#include "qcameracontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QCameraControl>
#include "_cgo_export.h"

class MyQCameraControl: public QCameraControl {
public:
void Signal_CaptureModeChanged(QCamera::CaptureModes mode){callbackQCameraControlCaptureModeChanged(this->objectName().toUtf8().data(), mode);};
void Signal_Error(int error, const QString & errorString){callbackQCameraControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data());};
void Signal_StateChanged(QCamera::State state){callbackQCameraControlStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StatusChanged(QCamera::Status status){callbackQCameraControlStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QCameraControl_CanChangeProperty(QtObjectPtr ptr, int changeType, int status){
	return static_cast<QCameraControl*>(ptr)->canChangeProperty(static_cast<QCameraControl::PropertyChangeType>(changeType), static_cast<QCamera::Status>(status));
}

int QCameraControl_CaptureMode(QtObjectPtr ptr){
	return static_cast<QCameraControl*>(ptr)->captureMode();
}

void QCameraControl_ConnectCaptureModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::CaptureModes)>(&QCameraControl::captureModeChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::CaptureModes)>(&MyQCameraControl::Signal_CaptureModeChanged));;
}

void QCameraControl_DisconnectCaptureModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::CaptureModes)>(&QCameraControl::captureModeChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::CaptureModes)>(&MyQCameraControl::Signal_CaptureModeChanged));;
}

void QCameraControl_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(int, const QString &)>(&QCameraControl::error), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(int, const QString &)>(&MyQCameraControl::Signal_Error));;
}

void QCameraControl_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(int, const QString &)>(&QCameraControl::error), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(int, const QString &)>(&MyQCameraControl::Signal_Error));;
}

int QCameraControl_IsCaptureModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraControl*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_SetCaptureMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraControl*>(ptr)->setCaptureMode(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_SetState(QtObjectPtr ptr, int state){
	static_cast<QCameraControl*>(ptr)->setState(static_cast<QCamera::State>(state));
}

void QCameraControl_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::State)>(&QCameraControl::stateChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::State)>(&MyQCameraControl::Signal_StateChanged));;
}

void QCameraControl_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::State)>(&QCameraControl::stateChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::State)>(&MyQCameraControl::Signal_StateChanged));;
}

int QCameraControl_Status(QtObjectPtr ptr){
	return static_cast<QCameraControl*>(ptr)->status();
}

void QCameraControl_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::Status)>(&QCameraControl::statusChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::Status)>(&MyQCameraControl::Signal_StatusChanged));;
}

void QCameraControl_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::Status)>(&QCameraControl::statusChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::Status)>(&MyQCameraControl::Signal_StatusChanged));;
}

void QCameraControl_DestroyQCameraControl(QtObjectPtr ptr){
	static_cast<QCameraControl*>(ptr)->~QCameraControl();
}

