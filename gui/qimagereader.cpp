#include "qimagereader.h"
#include <QIODevice>
#include <QString>
#include <QUrl>
#include <QRect>
#include <QColor>
#include <QSize>
#include <QImageIOHandler>
#include <QImage>
#include <QByteArray>
#include <QVariant>
#include <QModelIndex>
#include <QImageReader>
#include "_cgo_export.h"

class MyQImageReader: public QImageReader {
public:
};

QtObjectPtr QImageReader_NewQImageReader(){
	return new QImageReader();
}

QtObjectPtr QImageReader_NewQImageReader2(QtObjectPtr device, QtObjectPtr format){
	return new QImageReader(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

QtObjectPtr QImageReader_NewQImageReader3(char* fileName, QtObjectPtr format){
	return new QImageReader(QString(fileName), *static_cast<QByteArray*>(format));
}

int QImageReader_AutoDetectImageFormat(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->autoDetectImageFormat();
}

int QImageReader_AutoTransform(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->autoTransform();
}

int QImageReader_CanRead(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->canRead();
}

int QImageReader_CurrentImageNumber(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->currentImageNumber();
}

int QImageReader_DecideFormatFromContent(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->decideFormatFromContent();
}

QtObjectPtr QImageReader_Device(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->device();
}

int QImageReader_Error(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->error();
}

char* QImageReader_ErrorString(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->errorString().toUtf8().data();
}

char* QImageReader_FileName(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->fileName().toUtf8().data();
}

int QImageReader_ImageCount(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->imageCount();
}

int QImageReader_ImageFormat(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->imageFormat();
}

int QImageReader_JumpToImage(QtObjectPtr ptr, int imageNumber){
	return static_cast<QImageReader*>(ptr)->jumpToImage(imageNumber);
}

int QImageReader_JumpToNextImage(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->jumpToNextImage();
}

int QImageReader_LoopCount(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->loopCount();
}

int QImageReader_NextImageDelay(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->nextImageDelay();
}

int QImageReader_Quality(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->quality();
}

int QImageReader_Read2(QtObjectPtr ptr, QtObjectPtr image){
	return static_cast<QImageReader*>(ptr)->read(static_cast<QImage*>(image));
}

void QImageReader_SetAutoDetectImageFormat(QtObjectPtr ptr, int enabled){
	static_cast<QImageReader*>(ptr)->setAutoDetectImageFormat(enabled != 0);
}

void QImageReader_SetAutoTransform(QtObjectPtr ptr, int enabled){
	static_cast<QImageReader*>(ptr)->setAutoTransform(enabled != 0);
}

void QImageReader_SetBackgroundColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QImageReader*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QImageReader_SetClipRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QImageReader*>(ptr)->setClipRect(*static_cast<QRect*>(rect));
}

void QImageReader_SetDecideFormatFromContent(QtObjectPtr ptr, int ignored){
	static_cast<QImageReader*>(ptr)->setDecideFormatFromContent(ignored != 0);
}

void QImageReader_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QImageReader*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageReader_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QImageReader*>(ptr)->setFileName(QString(fileName));
}

void QImageReader_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QImageReader*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

void QImageReader_SetQuality(QtObjectPtr ptr, int quality){
	static_cast<QImageReader*>(ptr)->setQuality(quality);
}

void QImageReader_SetScaledClipRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QImageReader*>(ptr)->setScaledClipRect(*static_cast<QRect*>(rect));
}

void QImageReader_SetScaledSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QImageReader*>(ptr)->setScaledSize(*static_cast<QSize*>(size));
}

int QImageReader_SupportsAnimation(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->supportsAnimation();
}

int QImageReader_SupportsOption(QtObjectPtr ptr, int option){
	return static_cast<QImageReader*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

char* QImageReader_Text(QtObjectPtr ptr, char* key){
	return static_cast<QImageReader*>(ptr)->text(QString(key)).toUtf8().data();
}

char* QImageReader_TextKeys(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->textKeys().join("|").toUtf8().data();
}

int QImageReader_Transformation(QtObjectPtr ptr){
	return static_cast<QImageReader*>(ptr)->transformation();
}

void QImageReader_DestroyQImageReader(QtObjectPtr ptr){
	static_cast<QImageReader*>(ptr)->~QImageReader();
}

