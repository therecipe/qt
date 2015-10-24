#include "qpixelformat.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QPixelFormat>
#include "_cgo_export.h"

class MyQPixelFormat: public QPixelFormat {
public:
};

QtObjectPtr QPixelFormat_NewQPixelFormat(){
	return new QPixelFormat();
}

int QPixelFormat_AlphaPosition(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->alphaPosition();
}

int QPixelFormat_AlphaUsage(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->alphaUsage();
}

int QPixelFormat_ByteOrder(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->byteOrder();
}

int QPixelFormat_ColorModel(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->colorModel();
}

int QPixelFormat_Premultiplied(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->premultiplied();
}

int QPixelFormat_TypeInterpretation(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->typeInterpretation();
}

int QPixelFormat_YuvLayout(QtObjectPtr ptr){
	return static_cast<QPixelFormat*>(ptr)->yuvLayout();
}

