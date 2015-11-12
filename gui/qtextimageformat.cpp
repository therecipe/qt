#include "qtextimageformat.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QTextImageFormat>
#include "_cgo_export.h"

class MyQTextImageFormat: public QTextImageFormat {
public:
};

void* QTextImageFormat_NewQTextImageFormat(){
	return new QTextImageFormat();
}

double QTextImageFormat_Height(void* ptr){
	return static_cast<double>(static_cast<QTextImageFormat*>(ptr)->height());
}

int QTextImageFormat_IsValid(void* ptr){
	return static_cast<QTextImageFormat*>(ptr)->isValid();
}

char* QTextImageFormat_Name(void* ptr){
	return static_cast<QTextImageFormat*>(ptr)->name().toUtf8().data();
}

void QTextImageFormat_SetHeight(void* ptr, double height){
	static_cast<QTextImageFormat*>(ptr)->setHeight(static_cast<qreal>(height));
}

void QTextImageFormat_SetName(void* ptr, char* name){
	static_cast<QTextImageFormat*>(ptr)->setName(QString(name));
}

void QTextImageFormat_SetWidth(void* ptr, double width){
	static_cast<QTextImageFormat*>(ptr)->setWidth(static_cast<qreal>(width));
}

double QTextImageFormat_Width(void* ptr){
	return static_cast<double>(static_cast<QTextImageFormat*>(ptr)->width());
}

