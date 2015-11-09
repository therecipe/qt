#include "qvideoframe.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QAbstractVideoBuffer>
#include <QImage>
#include <QVideoFrame>
#include "_cgo_export.h"

class MyQVideoFrame: public QVideoFrame {
public:
};

void* QVideoFrame_NewQVideoFrame(){
	return new QVideoFrame();
}

void* QVideoFrame_NewQVideoFrame2(void* buffer, void* size, int format){
	return new QVideoFrame(static_cast<QAbstractVideoBuffer*>(buffer), *static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format));
}

void* QVideoFrame_NewQVideoFrame4(void* image){
	return new QVideoFrame(*static_cast<QImage*>(image));
}

void* QVideoFrame_NewQVideoFrame5(void* other){
	return new QVideoFrame(*static_cast<QVideoFrame*>(other));
}

void* QVideoFrame_NewQVideoFrame3(int bytes, void* size, int bytesPerLine, int format){
	return new QVideoFrame(bytes, *static_cast<QSize*>(size), bytesPerLine, static_cast<QVideoFrame::PixelFormat>(format));
}

int QVideoFrame_BytesPerLine(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine();
}

int QVideoFrame_BytesPerLine2(void* ptr, int plane){
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine(plane);
}

int QVideoFrame_FieldType(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->fieldType();
}

void* QVideoFrame_Handle(void* ptr){
	return new QVariant(static_cast<QVideoFrame*>(ptr)->handle());
}

int QVideoFrame_HandleType(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->handleType();
}

int QVideoFrame_Height(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->height();
}

int QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(int format){
	return QVideoFrame::imageFormatFromPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

int QVideoFrame_IsMapped(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isMapped();
}

int QVideoFrame_IsReadable(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isReadable();
}

int QVideoFrame_IsValid(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isValid();
}

int QVideoFrame_IsWritable(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isWritable();
}

int QVideoFrame_Map(void* ptr, int mode){
	return static_cast<QVideoFrame*>(ptr)->map(static_cast<QAbstractVideoBuffer::MapMode>(mode));
}

int QVideoFrame_MapMode(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->mapMode();
}

int QVideoFrame_MappedBytes(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->mappedBytes();
}

void* QVideoFrame_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QVideoFrame*>(ptr)->metaData(QString(key)));
}

int QVideoFrame_PixelFormat(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->pixelFormat();
}

int QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(int format){
	return QVideoFrame::pixelFormatFromImageFormat(static_cast<QImage::Format>(format));
}

int QVideoFrame_PlaneCount(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->planeCount();
}

void QVideoFrame_SetFieldType(void* ptr, int field){
	static_cast<QVideoFrame*>(ptr)->setFieldType(static_cast<QVideoFrame::FieldType>(field));
}

void QVideoFrame_SetMetaData(void* ptr, char* key, void* value){
	static_cast<QVideoFrame*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QVideoFrame_Unmap(void* ptr){
	static_cast<QVideoFrame*>(ptr)->unmap();
}

int QVideoFrame_Width(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->width();
}

void QVideoFrame_DestroyQVideoFrame(void* ptr){
	static_cast<QVideoFrame*>(ptr)->~QVideoFrame();
}

