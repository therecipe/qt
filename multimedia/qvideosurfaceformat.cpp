#include "qvideosurfaceformat.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoFrame>
#include <QAbstractVideoBuffer>
#include <QSize>
#include <QRect>
#include <QString>
#include <QVideoSurfaceFormat>
#include "_cgo_export.h"

class MyQVideoSurfaceFormat: public QVideoSurfaceFormat {
public:
};

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat(){
	return new QVideoSurfaceFormat();
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat2(void* size, int format, int ty){
	return new QVideoSurfaceFormat(*static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format), static_cast<QAbstractVideoBuffer::HandleType>(ty));
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat3(void* other){
	return new QVideoSurfaceFormat(*static_cast<QVideoSurfaceFormat*>(other));
}

int QVideoSurfaceFormat_FrameHeight(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameHeight();
}

double QVideoSurfaceFormat_FrameRate(void* ptr){
	return static_cast<double>(static_cast<QVideoSurfaceFormat*>(ptr)->frameRate());
}

int QVideoSurfaceFormat_FrameWidth(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameWidth();
}

int QVideoSurfaceFormat_HandleType(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->handleType();
}

int QVideoSurfaceFormat_IsValid(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->isValid();
}

int QVideoSurfaceFormat_PixelFormat(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->pixelFormat();
}

void* QVideoSurfaceFormat_Property(void* ptr, char* name){
	return new QVariant(static_cast<QVideoSurfaceFormat*>(ptr)->property(const_cast<const char*>(name)));
}

int QVideoSurfaceFormat_ScanLineDirection(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->scanLineDirection();
}

void QVideoSurfaceFormat_SetFrameRate(void* ptr, double rate){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameRate(static_cast<qreal>(rate));
}

void QVideoSurfaceFormat_SetFrameSize(void* ptr, void* size){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(*static_cast<QSize*>(size));
}

void QVideoSurfaceFormat_SetFrameSize2(void* ptr, int width, int height){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(width, height);
}

void QVideoSurfaceFormat_SetPixelAspectRatio(void* ptr, void* ratio){
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QVideoSurfaceFormat_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical){
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QVideoSurfaceFormat_SetProperty(void* ptr, char* name, void* value){
	static_cast<QVideoSurfaceFormat*>(ptr)->setProperty(const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

void QVideoSurfaceFormat_SetScanLineDirection(void* ptr, int direction){
	static_cast<QVideoSurfaceFormat*>(ptr)->setScanLineDirection(static_cast<QVideoSurfaceFormat::Direction>(direction));
}

void QVideoSurfaceFormat_SetViewport(void* ptr, void* viewport){
	static_cast<QVideoSurfaceFormat*>(ptr)->setViewport(*static_cast<QRect*>(viewport));
}

void QVideoSurfaceFormat_SetYCbCrColorSpace(void* ptr, int space){
	static_cast<QVideoSurfaceFormat*>(ptr)->setYCbCrColorSpace(static_cast<QVideoSurfaceFormat::YCbCrColorSpace>(space));
}

int QVideoSurfaceFormat_YCbCrColorSpace(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->yCbCrColorSpace();
}

void QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(void* ptr){
	static_cast<QVideoSurfaceFormat*>(ptr)->~QVideoSurfaceFormat();
}

