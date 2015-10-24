#include "qsizef.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSizeF>
#include "_cgo_export.h"

class MyQSizeF: public QSizeF {
public:
};

QtObjectPtr QSizeF_NewQSizeF(){
	return new QSizeF();
}

QtObjectPtr QSizeF_NewQSizeF2(QtObjectPtr size){
	return new QSizeF(*static_cast<QSize*>(size));
}

int QSizeF_IsEmpty(QtObjectPtr ptr){
	return static_cast<QSizeF*>(ptr)->isEmpty();
}

int QSizeF_IsNull(QtObjectPtr ptr){
	return static_cast<QSizeF*>(ptr)->isNull();
}

int QSizeF_IsValid(QtObjectPtr ptr){
	return static_cast<QSizeF*>(ptr)->isValid();
}

void QSizeF_Scale2(QtObjectPtr ptr, QtObjectPtr size, int mode){
	static_cast<QSizeF*>(ptr)->scale(*static_cast<QSizeF*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_Transpose(QtObjectPtr ptr){
	static_cast<QSizeF*>(ptr)->transpose();
}

