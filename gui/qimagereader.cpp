#include "qimagereader.h"
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QByteArray>
#include <QImageIOHandler>
#include <QUrl>
#include <QImage>
#include <QRect>
#include <QSize>
#include <QIODevice>
#include <QColor>
#include <QImageReader>
#include "_cgo_export.h"

class MyQImageReader: public QImageReader {
public:
};

void* QImageReader_NewQImageReader(){
	return new QImageReader();
}

void* QImageReader_NewQImageReader2(void* device, void* format){
	return new QImageReader(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

void* QImageReader_NewQImageReader3(char* fileName, void* format){
	return new QImageReader(QString(fileName), *static_cast<QByteArray*>(format));
}

int QImageReader_AutoDetectImageFormat(void* ptr){
	return static_cast<QImageReader*>(ptr)->autoDetectImageFormat();
}

int QImageReader_AutoTransform(void* ptr){
	return static_cast<QImageReader*>(ptr)->autoTransform();
}

void* QImageReader_BackgroundColor(void* ptr){
	return new QColor(static_cast<QImageReader*>(ptr)->backgroundColor());
}

int QImageReader_CanRead(void* ptr){
	return static_cast<QImageReader*>(ptr)->canRead();
}

int QImageReader_CurrentImageNumber(void* ptr){
	return static_cast<QImageReader*>(ptr)->currentImageNumber();
}

int QImageReader_DecideFormatFromContent(void* ptr){
	return static_cast<QImageReader*>(ptr)->decideFormatFromContent();
}

void* QImageReader_Device(void* ptr){
	return static_cast<QImageReader*>(ptr)->device();
}

int QImageReader_Error(void* ptr){
	return static_cast<QImageReader*>(ptr)->error();
}

char* QImageReader_ErrorString(void* ptr){
	return static_cast<QImageReader*>(ptr)->errorString().toUtf8().data();
}

char* QImageReader_FileName(void* ptr){
	return static_cast<QImageReader*>(ptr)->fileName().toUtf8().data();
}

void* QImageReader_Format(void* ptr){
	return new QByteArray(static_cast<QImageReader*>(ptr)->format());
}

int QImageReader_ImageCount(void* ptr){
	return static_cast<QImageReader*>(ptr)->imageCount();
}

void* QImageReader_QImageReader_ImageFormat3(void* device){
	return new QByteArray(QImageReader::imageFormat(static_cast<QIODevice*>(device)));
}

void* QImageReader_QImageReader_ImageFormat2(char* fileName){
	return new QByteArray(QImageReader::imageFormat(QString(fileName)));
}

int QImageReader_ImageFormat(void* ptr){
	return static_cast<QImageReader*>(ptr)->imageFormat();
}

int QImageReader_JumpToImage(void* ptr, int imageNumber){
	return static_cast<QImageReader*>(ptr)->jumpToImage(imageNumber);
}

int QImageReader_JumpToNextImage(void* ptr){
	return static_cast<QImageReader*>(ptr)->jumpToNextImage();
}

int QImageReader_LoopCount(void* ptr){
	return static_cast<QImageReader*>(ptr)->loopCount();
}

int QImageReader_NextImageDelay(void* ptr){
	return static_cast<QImageReader*>(ptr)->nextImageDelay();
}

int QImageReader_Quality(void* ptr){
	return static_cast<QImageReader*>(ptr)->quality();
}

int QImageReader_Read2(void* ptr, void* image){
	return static_cast<QImageReader*>(ptr)->read(static_cast<QImage*>(image));
}

void QImageReader_SetAutoDetectImageFormat(void* ptr, int enabled){
	static_cast<QImageReader*>(ptr)->setAutoDetectImageFormat(enabled != 0);
}

void QImageReader_SetAutoTransform(void* ptr, int enabled){
	static_cast<QImageReader*>(ptr)->setAutoTransform(enabled != 0);
}

void QImageReader_SetBackgroundColor(void* ptr, void* color){
	static_cast<QImageReader*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QImageReader_SetClipRect(void* ptr, void* rect){
	static_cast<QImageReader*>(ptr)->setClipRect(*static_cast<QRect*>(rect));
}

void QImageReader_SetDecideFormatFromContent(void* ptr, int ignored){
	static_cast<QImageReader*>(ptr)->setDecideFormatFromContent(ignored != 0);
}

void QImageReader_SetDevice(void* ptr, void* device){
	static_cast<QImageReader*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageReader_SetFileName(void* ptr, char* fileName){
	static_cast<QImageReader*>(ptr)->setFileName(QString(fileName));
}

void QImageReader_SetFormat(void* ptr, void* format){
	static_cast<QImageReader*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

void QImageReader_SetQuality(void* ptr, int quality){
	static_cast<QImageReader*>(ptr)->setQuality(quality);
}

void QImageReader_SetScaledClipRect(void* ptr, void* rect){
	static_cast<QImageReader*>(ptr)->setScaledClipRect(*static_cast<QRect*>(rect));
}

void QImageReader_SetScaledSize(void* ptr, void* size){
	static_cast<QImageReader*>(ptr)->setScaledSize(*static_cast<QSize*>(size));
}

void* QImageReader_SubType(void* ptr){
	return new QByteArray(static_cast<QImageReader*>(ptr)->subType());
}

int QImageReader_SupportsAnimation(void* ptr){
	return static_cast<QImageReader*>(ptr)->supportsAnimation();
}

int QImageReader_SupportsOption(void* ptr, int option){
	return static_cast<QImageReader*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

char* QImageReader_Text(void* ptr, char* key){
	return static_cast<QImageReader*>(ptr)->text(QString(key)).toUtf8().data();
}

char* QImageReader_TextKeys(void* ptr){
	return static_cast<QImageReader*>(ptr)->textKeys().join("|").toUtf8().data();
}

int QImageReader_Transformation(void* ptr){
	return static_cast<QImageReader*>(ptr)->transformation();
}

void QImageReader_DestroyQImageReader(void* ptr){
	static_cast<QImageReader*>(ptr)->~QImageReader();
}

