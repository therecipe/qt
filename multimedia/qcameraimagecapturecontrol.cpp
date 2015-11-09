#include "qcameraimagecapturecontrol.h"
#include <QCameraImageCapture>
#include <QObject>
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraImageCaptureControl>
#include "_cgo_export.h"

class MyQCameraImageCaptureControl: public QCameraImageCaptureControl {
public:
void Signal_Error(int id, int error, const QString & errorString){callbackQCameraImageCaptureControlError(this->objectName().toUtf8().data(), id, error, errorString.toUtf8().data());};
void Signal_ImageExposed(int requestId){callbackQCameraImageCaptureControlImageExposed(this->objectName().toUtf8().data(), requestId);};
void Signal_ImageMetadataAvailable(int id, const QString & key, const QVariant & value){callbackQCameraImageCaptureControlImageMetadataAvailable(this->objectName().toUtf8().data(), id, key.toUtf8().data(), new QVariant(value));};
void Signal_ImageSaved(int requestId, const QString & fileName){callbackQCameraImageCaptureControlImageSaved(this->objectName().toUtf8().data(), requestId, fileName.toUtf8().data());};
void Signal_ReadyForCaptureChanged(bool ready){callbackQCameraImageCaptureControlReadyForCaptureChanged(this->objectName().toUtf8().data(), ready);};
};

void QCameraImageCaptureControl_CancelCapture(void* ptr){
	static_cast<QCameraImageCaptureControl*>(ptr)->cancelCapture();
}

int QCameraImageCaptureControl_Capture(void* ptr, char* fileName){
	return static_cast<QCameraImageCaptureControl*>(ptr)->capture(QString(fileName));
}

int QCameraImageCaptureControl_DriveMode(void* ptr){
	return static_cast<QCameraImageCaptureControl*>(ptr)->driveMode();
}

void QCameraImageCaptureControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, int, const QString &)>(&QCameraImageCaptureControl::error), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, int, const QString &)>(&MyQCameraImageCaptureControl::Signal_Error));;
}

void QCameraImageCaptureControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, int, const QString &)>(&QCameraImageCaptureControl::error), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, int, const QString &)>(&MyQCameraImageCaptureControl::Signal_Error));;
}

void QCameraImageCaptureControl_ConnectImageExposed(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int)>(&QCameraImageCaptureControl::imageExposed), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int)>(&MyQCameraImageCaptureControl::Signal_ImageExposed));;
}

void QCameraImageCaptureControl_DisconnectImageExposed(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int)>(&QCameraImageCaptureControl::imageExposed), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int)>(&MyQCameraImageCaptureControl::Signal_ImageExposed));;
}

void QCameraImageCaptureControl_ConnectImageMetadataAvailable(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&QCameraImageCaptureControl::imageMetadataAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCaptureControl::Signal_ImageMetadataAvailable));;
}

void QCameraImageCaptureControl_DisconnectImageMetadataAvailable(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&QCameraImageCaptureControl::imageMetadataAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCaptureControl::Signal_ImageMetadataAvailable));;
}

void QCameraImageCaptureControl_ConnectImageSaved(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &)>(&QCameraImageCaptureControl::imageSaved), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &)>(&MyQCameraImageCaptureControl::Signal_ImageSaved));;
}

void QCameraImageCaptureControl_DisconnectImageSaved(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &)>(&QCameraImageCaptureControl::imageSaved), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &)>(&MyQCameraImageCaptureControl::Signal_ImageSaved));;
}

int QCameraImageCaptureControl_IsReadyForCapture(void* ptr){
	return static_cast<QCameraImageCaptureControl*>(ptr)->isReadyForCapture();
}

void QCameraImageCaptureControl_ConnectReadyForCaptureChanged(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(bool)>(&QCameraImageCaptureControl::readyForCaptureChanged), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(bool)>(&MyQCameraImageCaptureControl::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(bool)>(&QCameraImageCaptureControl::readyForCaptureChanged), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(bool)>(&MyQCameraImageCaptureControl::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCaptureControl_SetDriveMode(void* ptr, int mode){
	static_cast<QCameraImageCaptureControl*>(ptr)->setDriveMode(static_cast<QCameraImageCapture::DriveMode>(mode));
}

void QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(void* ptr){
	static_cast<QCameraImageCaptureControl*>(ptr)->~QCameraImageCaptureControl();
}

