#include "qdoublevalidator.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QDoubleValidator>
#include "_cgo_export.h"

class MyQDoubleValidator: public QDoubleValidator {
public:
};

int QDoubleValidator_Notation(void* ptr){
	return static_cast<QDoubleValidator*>(ptr)->notation();
}

void QDoubleValidator_SetDecimals(void* ptr, int v){
	static_cast<QDoubleValidator*>(ptr)->setDecimals(v);
}

void QDoubleValidator_SetNotation(void* ptr, int v){
	static_cast<QDoubleValidator*>(ptr)->setNotation(static_cast<QDoubleValidator::Notation>(v));
}

void* QDoubleValidator_NewQDoubleValidator(void* parent){
	return new QDoubleValidator(static_cast<QObject*>(parent));
}

int QDoubleValidator_Decimals(void* ptr){
	return static_cast<QDoubleValidator*>(ptr)->decimals();
}

void QDoubleValidator_DestroyQDoubleValidator(void* ptr){
	static_cast<QDoubleValidator*>(ptr)->~QDoubleValidator();
}

