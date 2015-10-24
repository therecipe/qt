#include "qintvalidator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QIntValidator>
#include "_cgo_export.h"

class MyQIntValidator: public QIntValidator {
public:
};

void QIntValidator_SetBottom(QtObjectPtr ptr, int v){
	static_cast<QIntValidator*>(ptr)->setBottom(v);
}

void QIntValidator_SetTop(QtObjectPtr ptr, int v){
	static_cast<QIntValidator*>(ptr)->setTop(v);
}

QtObjectPtr QIntValidator_NewQIntValidator(QtObjectPtr parent){
	return new QIntValidator(static_cast<QObject*>(parent));
}

QtObjectPtr QIntValidator_NewQIntValidator2(int minimum, int maximum, QtObjectPtr parent){
	return new QIntValidator(minimum, maximum, static_cast<QObject*>(parent));
}

int QIntValidator_Bottom(QtObjectPtr ptr){
	return static_cast<QIntValidator*>(ptr)->bottom();
}

void QIntValidator_SetRange(QtObjectPtr ptr, int bottom, int top){
	static_cast<QIntValidator*>(ptr)->setRange(bottom, top);
}

int QIntValidator_Top(QtObjectPtr ptr){
	return static_cast<QIntValidator*>(ptr)->top();
}

void QIntValidator_DestroyQIntValidator(QtObjectPtr ptr){
	static_cast<QIntValidator*>(ptr)->~QIntValidator();
}

