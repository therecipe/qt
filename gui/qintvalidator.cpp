#include "qintvalidator.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QIntValidator>
#include "_cgo_export.h"

class MyQIntValidator: public QIntValidator {
public:
};

void QIntValidator_SetBottom(void* ptr, int v){
	static_cast<QIntValidator*>(ptr)->setBottom(v);
}

void QIntValidator_SetTop(void* ptr, int v){
	static_cast<QIntValidator*>(ptr)->setTop(v);
}

void* QIntValidator_NewQIntValidator(void* parent){
	return new QIntValidator(static_cast<QObject*>(parent));
}

void* QIntValidator_NewQIntValidator2(int minimum, int maximum, void* parent){
	return new QIntValidator(minimum, maximum, static_cast<QObject*>(parent));
}

int QIntValidator_Bottom(void* ptr){
	return static_cast<QIntValidator*>(ptr)->bottom();
}

void QIntValidator_SetRange(void* ptr, int bottom, int top){
	static_cast<QIntValidator*>(ptr)->setRange(bottom, top);
}

int QIntValidator_Top(void* ptr){
	return static_cast<QIntValidator*>(ptr)->top();
}

void QIntValidator_DestroyQIntValidator(void* ptr){
	static_cast<QIntValidator*>(ptr)->~QIntValidator();
}

