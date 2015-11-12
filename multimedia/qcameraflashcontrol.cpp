#include "qcameraflashcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraExposure>
#include <QObject>
#include <QCamera>
#include <QCameraFlashControl>
#include "_cgo_export.h"

class MyQCameraFlashControl: public QCameraFlashControl {
public:
void Signal_FlashReady(bool ready){callbackQCameraFlashControlFlashReady(this->objectName().toUtf8().data(), ready);};
};

int QCameraFlashControl_FlashMode(void* ptr){
	return static_cast<QCameraFlashControl*>(ptr)->flashMode();
}

void QCameraFlashControl_ConnectFlashReady(void* ptr){
	QObject::connect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));;
}

void QCameraFlashControl_DisconnectFlashReady(void* ptr){
	QObject::disconnect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));;
}

int QCameraFlashControl_IsFlashModeSupported(void* ptr, int mode){
	return static_cast<QCameraFlashControl*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

int QCameraFlashControl_IsFlashReady(void* ptr){
	return static_cast<QCameraFlashControl*>(ptr)->isFlashReady();
}

void QCameraFlashControl_SetFlashMode(void* ptr, int mode){
	static_cast<QCameraFlashControl*>(ptr)->setFlashMode(static_cast<QCameraExposure::FlashMode>(mode));
}

void QCameraFlashControl_DestroyQCameraFlashControl(void* ptr){
	static_cast<QCameraFlashControl*>(ptr)->~QCameraFlashControl();
}

