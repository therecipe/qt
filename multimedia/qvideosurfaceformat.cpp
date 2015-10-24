#include "qvideosurfaceformat.h"
#include <QSize>
#include <QVideoFrame>
#include <QRect>
#include <QAbstractVideoBuffer>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoSurfaceFormat>
#include "_cgo_export.h"

class MyQVideoSurfaceFormat: public QVideoSurfaceFormat {
public:
};

QtObjectPtr QVideoSurfaceFormat_NewQVideoSurfaceFormat(){
	return new QVideoSurfaceFormat();
}

QtObjectPtr QVideoSurfaceFormat_NewQVideoSurfaceFormat2(QtObjectPtr size, int format, int ty){
	return new QVideoSurfaceFormat(*static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format), static_cast<QAbstractVideoBuffer::HandleType>(ty));
}

QtObjectPtr QVideoSurfaceFormat_NewQVideoSurfaceFormat3(QtObjectPtr other){
	return new QVideoSurfaceFormat(*static_cast<QVideoSurfaceFormat*>(other));
}

int QVideoSurfaceFormat_FrameHeight(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameHeight();
}

int QVideoSurfaceFormat_FrameWidth(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameWidth();
}

int QVideoSurfaceFormat_HandleType(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->handleType();
}

int QVideoSurfaceFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->isValid();
}

int QVideoSurfaceFormat_PixelFormat(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->pixelFormat();
}

char* QVideoSurfaceFormat_Property(QtObjectPtr ptr, char* name){
	return static_cast<QVideoSurfaceFormat*>(ptr)->property(const_cast<const char*>(name)).toString().toUtf8().data();
}

int QVideoSurfaceFormat_ScanLineDirection(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->scanLineDirection();
}

void QVideoSurfaceFormat_SetFrameSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(*static_cast<QSize*>(size));
}

void QVideoSurfaceFormat_SetFrameSize2(QtObjectPtr ptr, int width, int height){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(width, height);
}

void QVideoSurfaceFormat_SetPixelAspectRatio(QtObjectPtr ptr, QtObjectPtr ratio){
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QVideoSurfaceFormat_SetPixelAspectRatio2(QtObjectPtr ptr, int horizontal, int vertical){
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QVideoSurfaceFormat_SetProperty(QtObjectPtr ptr, char* name, char* value){
	static_cast<QVideoSurfaceFormat*>(ptr)->setProperty(const_cast<const char*>(name), QVariant(value));
}

void QVideoSurfaceFormat_SetScanLineDirection(QtObjectPtr ptr, int direction){
	static_cast<QVideoSurfaceFormat*>(ptr)->setScanLineDirection(static_cast<QVideoSurfaceFormat::Direction>(direction));
}

void QVideoSurfaceFormat_SetViewport(QtObjectPtr ptr, QtObjectPtr viewport){
	static_cast<QVideoSurfaceFormat*>(ptr)->setViewport(*static_cast<QRect*>(viewport));
}

void QVideoSurfaceFormat_SetYCbCrColorSpace(QtObjectPtr ptr, int space){
	static_cast<QVideoSurfaceFormat*>(ptr)->setYCbCrColorSpace(static_cast<QVideoSurfaceFormat::YCbCrColorSpace>(space));
}

int QVideoSurfaceFormat_YCbCrColorSpace(QtObjectPtr ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->yCbCrColorSpace();
}

void QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(QtObjectPtr ptr){
	static_cast<QVideoSurfaceFormat*>(ptr)->~QVideoSurfaceFormat();
}

