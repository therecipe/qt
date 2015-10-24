#include "qimagewriter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QImageIOHandler>
#include <QImage>
#include <QByteArray>
#include <QImageWriter>
#include "_cgo_export.h"

class MyQImageWriter: public QImageWriter {
public:
};

QtObjectPtr QImageWriter_NewQImageWriter(){
	return new QImageWriter();
}

QtObjectPtr QImageWriter_NewQImageWriter2(QtObjectPtr device, QtObjectPtr format){
	return new QImageWriter(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

QtObjectPtr QImageWriter_NewQImageWriter3(char* fileName, QtObjectPtr format){
	return new QImageWriter(QString(fileName), *static_cast<QByteArray*>(format));
}

int QImageWriter_CanWrite(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->canWrite();
}

int QImageWriter_Compression(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->compression();
}

QtObjectPtr QImageWriter_Device(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->device();
}

int QImageWriter_Error(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->error();
}

char* QImageWriter_ErrorString(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->errorString().toUtf8().data();
}

char* QImageWriter_FileName(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->fileName().toUtf8().data();
}

int QImageWriter_OptimizedWrite(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->optimizedWrite();
}

int QImageWriter_ProgressiveScanWrite(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->progressiveScanWrite();
}

int QImageWriter_Quality(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->quality();
}

void QImageWriter_SetCompression(QtObjectPtr ptr, int compression){
	static_cast<QImageWriter*>(ptr)->setCompression(compression);
}

void QImageWriter_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QImageWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageWriter_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QImageWriter*>(ptr)->setFileName(QString(fileName));
}

void QImageWriter_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QImageWriter*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

void QImageWriter_SetOptimizedWrite(QtObjectPtr ptr, int optimize){
	static_cast<QImageWriter*>(ptr)->setOptimizedWrite(optimize != 0);
}

void QImageWriter_SetProgressiveScanWrite(QtObjectPtr ptr, int progressive){
	static_cast<QImageWriter*>(ptr)->setProgressiveScanWrite(progressive != 0);
}

void QImageWriter_SetQuality(QtObjectPtr ptr, int quality){
	static_cast<QImageWriter*>(ptr)->setQuality(quality);
}

void QImageWriter_SetSubType(QtObjectPtr ptr, QtObjectPtr ty){
	static_cast<QImageWriter*>(ptr)->setSubType(*static_cast<QByteArray*>(ty));
}

void QImageWriter_SetText(QtObjectPtr ptr, char* key, char* text){
	static_cast<QImageWriter*>(ptr)->setText(QString(key), QString(text));
}

void QImageWriter_SetTransformation(QtObjectPtr ptr, int transform){
	static_cast<QImageWriter*>(ptr)->setTransformation(static_cast<QImageIOHandler::Transformation>(transform));
}

int QImageWriter_SupportsOption(QtObjectPtr ptr, int option){
	return static_cast<QImageWriter*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

int QImageWriter_Transformation(QtObjectPtr ptr){
	return static_cast<QImageWriter*>(ptr)->transformation();
}

int QImageWriter_Write(QtObjectPtr ptr, QtObjectPtr image){
	return static_cast<QImageWriter*>(ptr)->write(*static_cast<QImage*>(image));
}

void QImageWriter_DestroyQImageWriter(QtObjectPtr ptr){
	static_cast<QImageWriter*>(ptr)->~QImageWriter();
}

