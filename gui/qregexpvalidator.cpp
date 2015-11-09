#include "qregexpvalidator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QRegExp>
#include <QRegExpValidator>
#include "_cgo_export.h"

class MyQRegExpValidator: public QRegExpValidator {
public:
};

void QRegExpValidator_SetRegExp(void* ptr, void* rx){
	static_cast<QRegExpValidator*>(ptr)->setRegExp(*static_cast<QRegExp*>(rx));
}

void* QRegExpValidator_NewQRegExpValidator(void* parent){
	return new QRegExpValidator(static_cast<QObject*>(parent));
}

void* QRegExpValidator_NewQRegExpValidator2(void* rx, void* parent){
	return new QRegExpValidator(*static_cast<QRegExp*>(rx), static_cast<QObject*>(parent));
}

void* QRegExpValidator_RegExp(void* ptr){
	return new QRegExp(static_cast<QRegExpValidator*>(ptr)->regExp());
}

void QRegExpValidator_DestroyQRegExpValidator(void* ptr){
	static_cast<QRegExpValidator*>(ptr)->~QRegExpValidator();
}

