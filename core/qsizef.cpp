#include "qsizef.h"
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSizeF>
#include "_cgo_export.h"

class MyQSizeF: public QSizeF {
public:
};

void* QSizeF_NewQSizeF(){
	return new QSizeF();
}

void* QSizeF_NewQSizeF2(void* size){
	return new QSizeF(*static_cast<QSize*>(size));
}

void* QSizeF_NewQSizeF3(double width, double height){
	return new QSizeF(static_cast<qreal>(width), static_cast<qreal>(height));
}

double QSizeF_Height(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->height());
}

int QSizeF_IsEmpty(void* ptr){
	return static_cast<QSizeF*>(ptr)->isEmpty();
}

int QSizeF_IsNull(void* ptr){
	return static_cast<QSizeF*>(ptr)->isNull();
}

int QSizeF_IsValid(void* ptr){
	return static_cast<QSizeF*>(ptr)->isValid();
}

double QSizeF_Rheight(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->rheight());
}

double QSizeF_Rwidth(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->rwidth());
}

void QSizeF_Scale2(void* ptr, void* size, int mode){
	static_cast<QSizeF*>(ptr)->scale(*static_cast<QSizeF*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_Scale(void* ptr, double width, double height, int mode){
	static_cast<QSizeF*>(ptr)->scale(static_cast<qreal>(width), static_cast<qreal>(height), static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_SetHeight(void* ptr, double height){
	static_cast<QSizeF*>(ptr)->setHeight(static_cast<qreal>(height));
}

void QSizeF_SetWidth(void* ptr, double width){
	static_cast<QSizeF*>(ptr)->setWidth(static_cast<qreal>(width));
}

void QSizeF_Transpose(void* ptr){
	static_cast<QSizeF*>(ptr)->transpose();
}

double QSizeF_Width(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->width());
}

