#include "qsvggenerator.h"
#include <QIODevice>
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSvgGenerator>
#include "_cgo_export.h"

class MyQSvgGenerator: public QSvgGenerator {
public:
};

char* QSvgGenerator_Description(QtObjectPtr ptr){
	return static_cast<QSvgGenerator*>(ptr)->description().toUtf8().data();
}

char* QSvgGenerator_FileName(QtObjectPtr ptr){
	return static_cast<QSvgGenerator*>(ptr)->fileName().toUtf8().data();
}

QtObjectPtr QSvgGenerator_OutputDevice(QtObjectPtr ptr){
	return static_cast<QSvgGenerator*>(ptr)->outputDevice();
}

int QSvgGenerator_Resolution(QtObjectPtr ptr){
	return static_cast<QSvgGenerator*>(ptr)->resolution();
}

void QSvgGenerator_SetDescription(QtObjectPtr ptr, char* description){
	static_cast<QSvgGenerator*>(ptr)->setDescription(QString(description));
}

void QSvgGenerator_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QSvgGenerator*>(ptr)->setFileName(QString(fileName));
}

void QSvgGenerator_SetOutputDevice(QtObjectPtr ptr, QtObjectPtr outputDevice){
	static_cast<QSvgGenerator*>(ptr)->setOutputDevice(static_cast<QIODevice*>(outputDevice));
}

void QSvgGenerator_SetResolution(QtObjectPtr ptr, int dpi){
	static_cast<QSvgGenerator*>(ptr)->setResolution(dpi);
}

void QSvgGenerator_SetSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QSvgGenerator*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QSvgGenerator_SetTitle(QtObjectPtr ptr, char* title){
	static_cast<QSvgGenerator*>(ptr)->setTitle(QString(title));
}

void QSvgGenerator_SetViewBox(QtObjectPtr ptr, QtObjectPtr viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRect*>(viewBox));
}

void QSvgGenerator_SetViewBox2(QtObjectPtr ptr, QtObjectPtr viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRectF*>(viewBox));
}

char* QSvgGenerator_Title(QtObjectPtr ptr){
	return static_cast<QSvgGenerator*>(ptr)->title().toUtf8().data();
}

QtObjectPtr QSvgGenerator_NewQSvgGenerator(){
	return new QSvgGenerator();
}

void QSvgGenerator_DestroyQSvgGenerator(QtObjectPtr ptr){
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

