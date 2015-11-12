#include "qimagewriter.h"
#include <QIODevice>
#include <QImageIOHandler>
#include <QImage>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QImageWriter>
#include "_cgo_export.h"

class MyQImageWriter: public QImageWriter {
public:
};

void* QImageWriter_NewQImageWriter(){
	return new QImageWriter();
}

void* QImageWriter_NewQImageWriter2(void* device, void* format){
	return new QImageWriter(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

void* QImageWriter_NewQImageWriter3(char* fileName, void* format){
	return new QImageWriter(QString(fileName), *static_cast<QByteArray*>(format));
}

int QImageWriter_CanWrite(void* ptr){
	return static_cast<QImageWriter*>(ptr)->canWrite();
}

int QImageWriter_Compression(void* ptr){
	return static_cast<QImageWriter*>(ptr)->compression();
}

void* QImageWriter_Device(void* ptr){
	return static_cast<QImageWriter*>(ptr)->device();
}

int QImageWriter_Error(void* ptr){
	return static_cast<QImageWriter*>(ptr)->error();
}

char* QImageWriter_ErrorString(void* ptr){
	return static_cast<QImageWriter*>(ptr)->errorString().toUtf8().data();
}

char* QImageWriter_FileName(void* ptr){
	return static_cast<QImageWriter*>(ptr)->fileName().toUtf8().data();
}

void* QImageWriter_Format(void* ptr){
	return new QByteArray(static_cast<QImageWriter*>(ptr)->format());
}

int QImageWriter_OptimizedWrite(void* ptr){
	return static_cast<QImageWriter*>(ptr)->optimizedWrite();
}

int QImageWriter_ProgressiveScanWrite(void* ptr){
	return static_cast<QImageWriter*>(ptr)->progressiveScanWrite();
}

int QImageWriter_Quality(void* ptr){
	return static_cast<QImageWriter*>(ptr)->quality();
}

void QImageWriter_SetCompression(void* ptr, int compression){
	static_cast<QImageWriter*>(ptr)->setCompression(compression);
}

void QImageWriter_SetDevice(void* ptr, void* device){
	static_cast<QImageWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageWriter_SetFileName(void* ptr, char* fileName){
	static_cast<QImageWriter*>(ptr)->setFileName(QString(fileName));
}

void QImageWriter_SetFormat(void* ptr, void* format){
	static_cast<QImageWriter*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

void QImageWriter_SetOptimizedWrite(void* ptr, int optimize){
	static_cast<QImageWriter*>(ptr)->setOptimizedWrite(optimize != 0);
}

void QImageWriter_SetProgressiveScanWrite(void* ptr, int progressive){
	static_cast<QImageWriter*>(ptr)->setProgressiveScanWrite(progressive != 0);
}

void QImageWriter_SetQuality(void* ptr, int quality){
	static_cast<QImageWriter*>(ptr)->setQuality(quality);
}

void QImageWriter_SetSubType(void* ptr, void* ty){
	static_cast<QImageWriter*>(ptr)->setSubType(*static_cast<QByteArray*>(ty));
}

void QImageWriter_SetText(void* ptr, char* key, char* text){
	static_cast<QImageWriter*>(ptr)->setText(QString(key), QString(text));
}

void QImageWriter_SetTransformation(void* ptr, int transform){
	static_cast<QImageWriter*>(ptr)->setTransformation(static_cast<QImageIOHandler::Transformation>(transform));
}

void* QImageWriter_SubType(void* ptr){
	return new QByteArray(static_cast<QImageWriter*>(ptr)->subType());
}

int QImageWriter_SupportsOption(void* ptr, int option){
	return static_cast<QImageWriter*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

int QImageWriter_Transformation(void* ptr){
	return static_cast<QImageWriter*>(ptr)->transformation();
}

int QImageWriter_Write(void* ptr, void* image){
	return static_cast<QImageWriter*>(ptr)->write(*static_cast<QImage*>(image));
}

void QImageWriter_DestroyQImageWriter(void* ptr){
	static_cast<QImageWriter*>(ptr)->~QImageWriter();
}

