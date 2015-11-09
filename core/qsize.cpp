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

void* QSize_NewQSize(){
	return new QSize();
}

void* QSize_NewQSize2(int width, int height){
	return new QSize(width, height);
}

int QSize_Height(void* ptr){
	return static_cast<QSize*>(ptr)->height();
}

int QSize_IsEmpty(void* ptr){
	return static_cast<QSize*>(ptr)->isEmpty();
}

int QSize_IsNull(void* ptr){
	return static_cast<QSize*>(ptr)->isNull();
}

int QSize_IsValid(void* ptr){
	return static_cast<QSize*>(ptr)->isValid();
}

int QSize_Rheight(void* ptr){
	return static_cast<QSize*>(ptr)->rheight();
}

int QSize_Rwidth(void* ptr){
	return static_cast<QSize*>(ptr)->rwidth();
}

void QSize_Scale2(void* ptr, void* size, int mode){
	static_cast<QSize*>(ptr)->scale(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_Scale(void* ptr, int width, int height, int mode){
	static_cast<QSize*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_SetHeight(void* ptr, int height){
	static_cast<QSize*>(ptr)->setHeight(height);
}

void QSize_SetWidth(void* ptr, int width){
	static_cast<QSize*>(ptr)->setWidth(width);
}

void QSize_Transpose(void* ptr){
	static_cast<QSize*>(ptr)->transpose();
}

int QSize_Width(void* ptr){
	return static_cast<QSize*>(ptr)->width();
}

