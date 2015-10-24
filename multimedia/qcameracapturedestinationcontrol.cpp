#include "qcameracapturedestinationcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraImageCapture>
#include <QObject>
#include <QCameraCaptureDestinationControl>
#include "_cgo_export.h"

class MyQCameraCaptureDestinationControl: public QCameraCaptureDestinationControl {
public:
void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination){callbackQCameraCaptureDestinationControlCaptureDestinationChanged(this->objectName().toUtf8().data(), destination);};
};

int QCameraCaptureDestinationControl_CaptureDestination(QtObjectPtr ptr){
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->captureDestination();
}

void QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));;
}

void QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));;
}

int QCameraCaptureDestinationControl_IsCaptureDestinationSupported(QtObjectPtr ptr, int destination){
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_SetCaptureDestination(QtObjectPtr ptr, int destination){
	static_cast<QCameraCaptureDestinationControl*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(QtObjectPtr ptr){
	static_cast<QCameraCaptureDestinationControl*>(ptr)->~QCameraCaptureDestinationControl();
}

