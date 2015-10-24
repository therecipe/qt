#include "qcameraflashcontrol.h"
#include <QCamera>
#include <QCameraExposure>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraFlashControl>
#include "_cgo_export.h"

class MyQCameraFlashControl: public QCameraFlashControl {
public:
void Signal_FlashReady(bool ready){callbackQCameraFlashControlFlashReady(this->objectName().toUtf8().data(), ready);};
};

int QCameraFlashControl_FlashMode(QtObjectPtr ptr){
	return static_cast<QCameraFlashControl*>(ptr)->flashMode();
}

void QCameraFlashControl_ConnectFlashReady(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));;
}

void QCameraFlashControl_DisconnectFlashReady(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));;
}

int QCameraFlashControl_IsFlashModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraFlashControl*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

int QCameraFlashControl_IsFlashReady(QtObjectPtr ptr){
	return static_cast<QCameraFlashControl*>(ptr)->isFlashReady();
}

void QCameraFlashControl_SetFlashMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraFlashControl*>(ptr)->setFlashMode(static_cast<QCameraExposure::FlashMode>(mode));
}

void QCameraFlashControl_DestroyQCameraFlashControl(QtObjectPtr ptr){
	static_cast<QCameraFlashControl*>(ptr)->~QCameraFlashControl();
}

