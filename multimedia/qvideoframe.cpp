#include "qvideoframe.h"
#include <QAbstractVideoBuffer>
#include <QImage>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoFrame>
#include "_cgo_export.h"

class MyQVideoFrame: public QVideoFrame {
public:
};

QtObjectPtr QVideoFrame_NewQVideoFrame(){
	return new QVideoFrame();
}

QtObjectPtr QVideoFrame_NewQVideoFrame2(QtObjectPtr buffer, QtObjectPtr size, int format){
	return new QVideoFrame(static_cast<QAbstractVideoBuffer*>(buffer), *static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format));
}

QtObjectPtr QVideoFrame_NewQVideoFrame4(QtObjectPtr image){
	return new QVideoFrame(*static_cast<QImage*>(image));
}

QtObjectPtr QVideoFrame_NewQVideoFrame5(QtObjectPtr other){
	return new QVideoFrame(*static_cast<QVideoFrame*>(other));
}

QtObjectPtr QVideoFrame_NewQVideoFrame3(int bytes, QtObjectPtr size, int bytesPerLine, int format){
	return new QVideoFrame(bytes, *static_cast<QSize*>(size), bytesPerLine, static_cast<QVideoFrame::PixelFormat>(format));
}

int QVideoFrame_BytesPerLine(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine();
}

int QVideoFrame_BytesPerLine2(QtObjectPtr ptr, int plane){
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine(plane);
}

int QVideoFrame_FieldType(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->fieldType();
}

char* QVideoFrame_Handle(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->handle().toString().toUtf8().data();
}

int QVideoFrame_HandleType(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->handleType();
}

int QVideoFrame_Height(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->height();
}

int QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(int format){
	return QVideoFrame::imageFormatFromPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

int QVideoFrame_IsMapped(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->isMapped();
}

int QVideoFrame_IsReadable(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->isReadable();
}

int QVideoFrame_IsValid(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->isValid();
}

int QVideoFrame_IsWritable(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->isWritable();
}

int QVideoFrame_Map(QtObjectPtr ptr, int mode){
	return static_cast<QVideoFrame*>(ptr)->map(static_cast<QAbstractVideoBuffer::MapMode>(mode));
}

int QVideoFrame_MapMode(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->mapMode();
}

int QVideoFrame_MappedBytes(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->mappedBytes();
}

char* QVideoFrame_MetaData(QtObjectPtr ptr, char* key){
	return static_cast<QVideoFrame*>(ptr)->metaData(QString(key)).toString().toUtf8().data();
}

int QVideoFrame_PixelFormat(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->pixelFormat();
}

int QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(int format){
	return QVideoFrame::pixelFormatFromImageFormat(static_cast<QImage::Format>(format));
}

int QVideoFrame_PlaneCount(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->planeCount();
}

void QVideoFrame_SetFieldType(QtObjectPtr ptr, int field){
	static_cast<QVideoFrame*>(ptr)->setFieldType(static_cast<QVideoFrame::FieldType>(field));
}

void QVideoFrame_SetMetaData(QtObjectPtr ptr, char* key, char* value){
	static_cast<QVideoFrame*>(ptr)->setMetaData(QString(key), QVariant(value));
}

void QVideoFrame_Unmap(QtObjectPtr ptr){
	static_cast<QVideoFrame*>(ptr)->unmap();
}

int QVideoFrame_Width(QtObjectPtr ptr){
	return static_cast<QVideoFrame*>(ptr)->width();
}

void QVideoFrame_DestroyQVideoFrame(QtObjectPtr ptr){
	static_cast<QVideoFrame*>(ptr)->~QVideoFrame();
}

