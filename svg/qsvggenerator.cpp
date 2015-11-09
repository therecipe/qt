#include "qsvggenerator.h"
#include <QIODevice>
#include <QRectF>
#include <QSize>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSvgGenerator>
#include "_cgo_export.h"

class MyQSvgGenerator: public QSvgGenerator {
public:
};

char* QSvgGenerator_Description(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->description().toUtf8().data();
}

char* QSvgGenerator_FileName(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->fileName().toUtf8().data();
}

void* QSvgGenerator_OutputDevice(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->outputDevice();
}

int QSvgGenerator_Resolution(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->resolution();
}

void QSvgGenerator_SetDescription(void* ptr, char* description){
	static_cast<QSvgGenerator*>(ptr)->setDescription(QString(description));
}

void QSvgGenerator_SetFileName(void* ptr, char* fileName){
	static_cast<QSvgGenerator*>(ptr)->setFileName(QString(fileName));
}

void QSvgGenerator_SetOutputDevice(void* ptr, void* outputDevice){
	static_cast<QSvgGenerator*>(ptr)->setOutputDevice(static_cast<QIODevice*>(outputDevice));
}

void QSvgGenerator_SetResolution(void* ptr, int dpi){
	static_cast<QSvgGenerator*>(ptr)->setResolution(dpi);
}

void QSvgGenerator_SetSize(void* ptr, void* size){
	static_cast<QSvgGenerator*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QSvgGenerator_SetTitle(void* ptr, char* title){
	static_cast<QSvgGenerator*>(ptr)->setTitle(QString(title));
}

void QSvgGenerator_SetViewBox(void* ptr, void* viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRect*>(viewBox));
}

void QSvgGenerator_SetViewBox2(void* ptr, void* viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRectF*>(viewBox));
}

char* QSvgGenerator_Title(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->title().toUtf8().data();
}

void* QSvgGenerator_NewQSvgGenerator(){
	return new QSvgGenerator();
}

void QSvgGenerator_DestroyQSvgGenerator(void* ptr){
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

