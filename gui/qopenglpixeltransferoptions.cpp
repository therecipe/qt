#include "qopenglpixeltransferoptions.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLPixelTransferOptions>
#include "_cgo_export.h"

class MyQOpenGLPixelTransferOptions: public QOpenGLPixelTransferOptions {
public:
};

QtObjectPtr QOpenGLPixelTransferOptions_NewQOpenGLPixelTransferOptions(){
	return new QOpenGLPixelTransferOptions();
}

int QOpenGLPixelTransferOptions_Alignment(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->alignment();
}

int QOpenGLPixelTransferOptions_ImageHeight(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->imageHeight();
}

int QOpenGLPixelTransferOptions_IsLeastSignificantBitFirst(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->isLeastSignificantBitFirst();
}

int QOpenGLPixelTransferOptions_IsSwapBytesEnabled(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->isSwapBytesEnabled();
}

void QOpenGLPixelTransferOptions_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setAlignment(alignment);
}

void QOpenGLPixelTransferOptions_SetImageHeight(QtObjectPtr ptr, int imageHeight){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setImageHeight(imageHeight);
}

void QOpenGLPixelTransferOptions_SetSkipImages(QtObjectPtr ptr, int skipImages){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setSkipImages(skipImages);
}

void QOpenGLPixelTransferOptions_SetSkipPixels(QtObjectPtr ptr, int skipPixels){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setSkipPixels(skipPixels);
}

void QOpenGLPixelTransferOptions_SetSkipRows(QtObjectPtr ptr, int skipRows){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setSkipRows(skipRows);
}

int QOpenGLPixelTransferOptions_SkipImages(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->skipImages();
}

int QOpenGLPixelTransferOptions_SkipPixels(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->skipPixels();
}

int QOpenGLPixelTransferOptions_SkipRows(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->skipRows();
}

int QOpenGLPixelTransferOptions_RowLength(QtObjectPtr ptr){
	return static_cast<QOpenGLPixelTransferOptions*>(ptr)->rowLength();
}

void QOpenGLPixelTransferOptions_SetLeastSignificantByteFirst(QtObjectPtr ptr, int lsbFirst){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setLeastSignificantByteFirst(lsbFirst != 0);
}

void QOpenGLPixelTransferOptions_SetRowLength(QtObjectPtr ptr, int rowLength){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setRowLength(rowLength);
}

void QOpenGLPixelTransferOptions_SetSwapBytesEnabled(QtObjectPtr ptr, int swapBytes){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->setSwapBytesEnabled(swapBytes != 0);
}

void QOpenGLPixelTransferOptions_DestroyQOpenGLPixelTransferOptions(QtObjectPtr ptr){
	static_cast<QOpenGLPixelTransferOptions*>(ptr)->~QOpenGLPixelTransferOptions();
}

