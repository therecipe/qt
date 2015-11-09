#include "qmargins.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QMargins>
#include "_cgo_export.h"

class MyQMargins: public QMargins {
public:
};

void* QMargins_NewQMargins(){
	return new QMargins();
}

void* QMargins_NewQMargins2(int left, int top, int right, int bottom){
	return new QMargins(left, top, right, bottom);
}

int QMargins_Bottom(void* ptr){
	return static_cast<QMargins*>(ptr)->bottom();
}

int QMargins_IsNull(void* ptr){
	return static_cast<QMargins*>(ptr)->isNull();
}

int QMargins_Left(void* ptr){
	return static_cast<QMargins*>(ptr)->left();
}

int QMargins_Right(void* ptr){
	return static_cast<QMargins*>(ptr)->right();
}

void QMargins_SetBottom(void* ptr, int bottom){
	static_cast<QMargins*>(ptr)->setBottom(bottom);
}

void QMargins_SetLeft(void* ptr, int left){
	static_cast<QMargins*>(ptr)->setLeft(left);
}

void QMargins_SetRight(void* ptr, int right){
	static_cast<QMargins*>(ptr)->setRight(right);
}

void QMargins_SetTop(void* ptr, int Top){
	static_cast<QMargins*>(ptr)->setTop(Top);
}

int QMargins_Top(void* ptr){
	return static_cast<QMargins*>(ptr)->top();
}

