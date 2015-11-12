#include "qcameracapturedestinationcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QCameraImageCapture>
#include <QCamera>
#include <QCameraCaptureDestinationControl>
#include "_cgo_export.h"

class MyQCameraCaptureDestinationControl: public QCameraCaptureDestinationControl {
public:
void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination){callbackQCameraCaptureDestinationControlCaptureDestinationChanged(this->objectName().toUtf8().data(), destination);};
};

int QCameraCaptureDestinationControl_CaptureDestination(void* ptr){
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->captureDestination();
}

void QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(void* ptr){
	QObject::connect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));;
}

void QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));;
}

int QCameraCaptureDestinationControl_IsCaptureDestinationSupported(void* ptr, int destination){
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_SetCaptureDestination(void* ptr, int destination){
	static_cast<QCameraCaptureDestinationControl*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(void* ptr){
	static_cast<QCameraCaptureDestinationControl*>(ptr)->~QCameraCaptureDestinationControl();
}

