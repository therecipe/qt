#include "qmarginsf.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMargins>
#include <QString>
#include <QMarginsF>
#include "_cgo_export.h"

class MyQMarginsF: public QMarginsF {
public:
};

void* QMarginsF_NewQMarginsF(){
	return new QMarginsF();
}

void* QMarginsF_NewQMarginsF3(void* margins){
	return new QMarginsF(*static_cast<QMargins*>(margins));
}

void* QMarginsF_NewQMarginsF2(double left, double top, double right, double bottom){
	return new QMarginsF(static_cast<qreal>(left), static_cast<qreal>(top), static_cast<qreal>(right), static_cast<qreal>(bottom));
}

double QMarginsF_Bottom(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->bottom());
}

int QMarginsF_IsNull(void* ptr){
	return static_cast<QMarginsF*>(ptr)->isNull();
}

double QMarginsF_Left(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->left());
}

double QMarginsF_Right(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->right());
}

void QMarginsF_SetBottom(void* ptr, double bottom){
	static_cast<QMarginsF*>(ptr)->setBottom(static_cast<qreal>(bottom));
}

void QMarginsF_SetLeft(void* ptr, double left){
	static_cast<QMarginsF*>(ptr)->setLeft(static_cast<qreal>(left));
}

void QMarginsF_SetRight(void* ptr, double right){
	static_cast<QMarginsF*>(ptr)->setRight(static_cast<qreal>(right));
}

void QMarginsF_SetTop(void* ptr, double Top){
	static_cast<QMarginsF*>(ptr)->setTop(static_cast<qreal>(Top));
}

double QMarginsF_Top(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->top());
}

