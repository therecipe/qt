#include "qcameracapturebufferformatcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QObject>
#include <QVideoFrame>
#include <QCameraCaptureBufferFormatControl>
#include "_cgo_export.h"

class MyQCameraCaptureBufferFormatControl: public QCameraCaptureBufferFormatControl {
public:
void Signal_BufferFormatChanged(QVideoFrame::PixelFormat format){callbackQCameraCaptureBufferFormatControlBufferFormatChanged(this->objectName().toUtf8().data(), format);};
};

int QCameraCaptureBufferFormatControl_BufferFormat(void* ptr){
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->bufferFormat();
}

void QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(void* ptr){
	QObject::connect(static_cast<QCameraCaptureBufferFormatControl*>(ptr), static_cast<void (QCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&QCameraCaptureBufferFormatControl::bufferFormatChanged), static_cast<MyQCameraCaptureBufferFormatControl*>(ptr), static_cast<void (MyQCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&MyQCameraCaptureBufferFormatControl::Signal_BufferFormatChanged));;
}

void QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraCaptureBufferFormatControl*>(ptr), static_cast<void (QCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&QCameraCaptureBufferFormatControl::bufferFormatChanged), static_cast<MyQCameraCaptureBufferFormatControl*>(ptr), static_cast<void (MyQCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&MyQCameraCaptureBufferFormatControl::Signal_BufferFormatChanged));;
}

void QCameraCaptureBufferFormatControl_SetBufferFormat(void* ptr, int format){
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->setBufferFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(void* ptr){
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->~QCameraCaptureBufferFormatControl();
}

