#include "qdoublevalidator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDoubleValidator>
#include "_cgo_export.h"

class MyQDoubleValidator: public QDoubleValidator {
public:
};

int QDoubleValidator_Notation(QtObjectPtr ptr){
	return static_cast<QDoubleValidator*>(ptr)->notation();
}

void QDoubleValidator_SetDecimals(QtObjectPtr ptr, int v){
	static_cast<QDoubleValidator*>(ptr)->setDecimals(v);
}

void QDoubleValidator_SetNotation(QtObjectPtr ptr, int v){
	static_cast<QDoubleValidator*>(ptr)->setNotation(static_cast<QDoubleValidator::Notation>(v));
}

QtObjectPtr QDoubleValidator_NewQDoubleValidator(QtObjectPtr parent){
	return new QDoubleValidator(static_cast<QObject*>(parent));
}

int QDoubleValidator_Decimals(QtObjectPtr ptr){
	return static_cast<QDoubleValidator*>(ptr)->decimals();
}

void QDoubleValidator_DestroyQDoubleValidator(QtObjectPtr ptr){
	static_cast<QDoubleValidator*>(ptr)->~QDoubleValidator();
}

