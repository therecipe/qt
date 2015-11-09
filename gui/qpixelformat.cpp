#include "qpixelformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPixelFormat>
#include "_cgo_export.h"

class MyQPixelFormat: public QPixelFormat {
public:
};

void* QPixelFormat_NewQPixelFormat(){
	return new QPixelFormat();
}

int QPixelFormat_AlphaPosition(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->alphaPosition();
}

int QPixelFormat_AlphaUsage(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->alphaUsage();
}

int QPixelFormat_ByteOrder(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->byteOrder();
}

int QPixelFormat_ColorModel(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->colorModel();
}

int QPixelFormat_Premultiplied(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->premultiplied();
}

int QPixelFormat_TypeInterpretation(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->typeInterpretation();
}

int QPixelFormat_YuvLayout(void* ptr){
	return static_cast<QPixelFormat*>(ptr)->yuvLayout();
}

