#include "qcameraimagecapture.h"
#include <QUrl>
#include <QModelIndex>
#include <QImageEncoderSettings>
#include <QImage>
#include <QCamera>
#include <QVariant>
#include <QMediaObject>
#include <QMetaObject>
#include <QVideoFrame>
#include <QObject>
#include <QString>
#include <QCameraImageCapture>
#include "_cgo_export.h"

class MyQCameraImageCapture: public QCameraImageCapture {
public:
void Signal_BufferFormatChanged(QVideoFrame::PixelFormat format){callbackQCameraImageCaptureBufferFormatChanged(this->objectName().toUtf8().data(), format);};
void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination){callbackQCameraImageCaptureCaptureDestinationChanged(this->objectName().toUtf8().data(), destination);};
void Signal_ImageExposed(int id){callbackQCameraImageCaptureImageExposed(this->objectName().toUtf8().data(), id);};
void Signal_ImageMetadataAvailable(int id, const QString & key, const QVariant & value){callbackQCameraImageCaptureImageMetadataAvailable(this->objectName().toUtf8().data(), id, key.toUtf8().data(), value.toString().toUtf8().data());};
void Signal_ImageSaved(int id, const QString & fileName){callbackQCameraImageCaptureImageSaved(this->objectName().toUtf8().data(), id, fileName.toUtf8().data());};
void Signal_ReadyForCaptureChanged(bool ready){callbackQCameraImageCaptureReadyForCaptureChanged(this->objectName().toUtf8().data(), ready);};
};

int QCameraImageCapture_IsReadyForCapture(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->isReadyForCapture();
}

QtObjectPtr QCameraImageCapture_NewQCameraImageCapture(QtObjectPtr mediaObject, QtObjectPtr parent){
	return new QCameraImageCapture(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

int QCameraImageCapture_BufferFormat(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->bufferFormat();
}

void QCameraImageCapture_ConnectBufferFormatChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&QCameraImageCapture::bufferFormatChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&MyQCameraImageCapture::Signal_BufferFormatChanged));;
}

void QCameraImageCapture_DisconnectBufferFormatChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&QCameraImageCapture::bufferFormatChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&MyQCameraImageCapture::Signal_BufferFormatChanged));;
}

void QCameraImageCapture_CancelCapture(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "cancelCapture");
}

int QCameraImageCapture_Capture(QtObjectPtr ptr, char* file){
	return QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "capture", Q_ARG(QString, QString(file)));
}

int QCameraImageCapture_CaptureDestination(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->captureDestination();
}

void QCameraImageCapture_ConnectCaptureDestinationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraImageCapture::captureDestinationChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraImageCapture::Signal_CaptureDestinationChanged));;
}

void QCameraImageCapture_DisconnectCaptureDestinationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraImageCapture::captureDestinationChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraImageCapture::Signal_CaptureDestinationChanged));;
}

int QCameraImageCapture_Error(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->error();
}

char* QCameraImageCapture_ErrorString(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->errorString().toUtf8().data();
}

char* QCameraImageCapture_ImageCodecDescription(QtObjectPtr ptr, char* codec){
	return static_cast<QCameraImageCapture*>(ptr)->imageCodecDescription(QString(codec)).toUtf8().data();
}

void QCameraImageCapture_ConnectImageExposed(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int)>(&QCameraImageCapture::imageExposed), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int)>(&MyQCameraImageCapture::Signal_ImageExposed));;
}

void QCameraImageCapture_DisconnectImageExposed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int)>(&QCameraImageCapture::imageExposed), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int)>(&MyQCameraImageCapture::Signal_ImageExposed));;
}

void QCameraImageCapture_ConnectImageMetadataAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &, const QVariant &)>(&QCameraImageCapture::imageMetadataAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCapture::Signal_ImageMetadataAvailable));;
}

void QCameraImageCapture_DisconnectImageMetadataAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &, const QVariant &)>(&QCameraImageCapture::imageMetadataAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCapture::Signal_ImageMetadataAvailable));;
}

void QCameraImageCapture_ConnectImageSaved(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &)>(&QCameraImageCapture::imageSaved), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &)>(&MyQCameraImageCapture::Signal_ImageSaved));;
}

void QCameraImageCapture_DisconnectImageSaved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &)>(&QCameraImageCapture::imageSaved), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &)>(&MyQCameraImageCapture::Signal_ImageSaved));;
}

int QCameraImageCapture_IsAvailable(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->isAvailable();
}

int QCameraImageCapture_IsCaptureDestinationSupported(QtObjectPtr ptr, int destination){
	return static_cast<QCameraImageCapture*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

QtObjectPtr QCameraImageCapture_MediaObject(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->mediaObject();
}

void QCameraImageCapture_ConnectReadyForCaptureChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(bool)>(&QCameraImageCapture::readyForCaptureChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(bool)>(&MyQCameraImageCapture::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCapture_DisconnectReadyForCaptureChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(bool)>(&QCameraImageCapture::readyForCaptureChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(bool)>(&MyQCameraImageCapture::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCapture_SetBufferFormat(QtObjectPtr ptr, int format){
	static_cast<QCameraImageCapture*>(ptr)->setBufferFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraImageCapture_SetCaptureDestination(QtObjectPtr ptr, int destination){
	static_cast<QCameraImageCapture*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraImageCapture_SetEncodingSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QCameraImageCapture*>(ptr)->setEncodingSettings(*static_cast<QImageEncoderSettings*>(settings));
}

char* QCameraImageCapture_SupportedImageCodecs(QtObjectPtr ptr){
	return static_cast<QCameraImageCapture*>(ptr)->supportedImageCodecs().join("|").toUtf8().data();
}

void QCameraImageCapture_DestroyQCameraImageCapture(QtObjectPtr ptr){
	static_cast<QCameraImageCapture*>(ptr)->~QCameraImageCapture();
}

