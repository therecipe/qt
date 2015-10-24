#include "qmargins.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMargins>
#include "_cgo_export.h"

class MyQMargins: public QMargins {
public:
};

QtObjectPtr QMargins_NewQMargins(){
	return new QMargins();
}

QtObjectPtr QMargins_NewQMargins2(int left, int top, int right, int bottom){
	return new QMargins(left, top, right, bottom);
}

int QMargins_Bottom(QtObjectPtr ptr){
	return static_cast<QMargins*>(ptr)->bottom();
}

int QMargins_IsNull(QtObjectPtr ptr){
	return static_cast<QMargins*>(ptr)->isNull();
}

int QMargins_Left(QtObjectPtr ptr){
	return static_cast<QMargins*>(ptr)->left();
}

int QMargins_Right(QtObjectPtr ptr){
	return static_cast<QMargins*>(ptr)->right();
}

void QMargins_SetBottom(QtObjectPtr ptr, int bottom){
	static_cast<QMargins*>(ptr)->setBottom(bottom);
}

void QMargins_SetLeft(QtObjectPtr ptr, int left){
	static_cast<QMargins*>(ptr)->setLeft(left);
}

void QMargins_SetRight(QtObjectPtr ptr, int right){
	static_cast<QMargins*>(ptr)->setRight(right);
}

void QMargins_SetTop(QtObjectPtr ptr, int Top){
	static_cast<QMargins*>(ptr)->setTop(Top);
}

int QMargins_Top(QtObjectPtr ptr){
	return static_cast<QMargins*>(ptr)->top();
}

