#include "qsize.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include "_cgo_export.h"

class MyQSize: public QSize {
public:
};

QtObjectPtr QSize_NewQSize(){
	return new QSize();
}

QtObjectPtr QSize_NewQSize2(int width, int height){
	return new QSize(width, height);
}

int QSize_Height(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->height();
}

int QSize_IsEmpty(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->isEmpty();
}

int QSize_IsNull(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->isNull();
}

int QSize_IsValid(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->isValid();
}

int QSize_Rheight(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->rheight();
}

int QSize_Rwidth(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->rwidth();
}

void QSize_Scale2(QtObjectPtr ptr, QtObjectPtr size, int mode){
	static_cast<QSize*>(ptr)->scale(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_Scale(QtObjectPtr ptr, int width, int height, int mode){
	static_cast<QSize*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_SetHeight(QtObjectPtr ptr, int height){
	static_cast<QSize*>(ptr)->setHeight(height);
}

void QSize_SetWidth(QtObjectPtr ptr, int width){
	static_cast<QSize*>(ptr)->setWidth(width);
}

void QSize_Transpose(QtObjectPtr ptr){
	static_cast<QSize*>(ptr)->transpose();
}

int QSize_Width(QtObjectPtr ptr){
	return static_cast<QSize*>(ptr)->width();
}

