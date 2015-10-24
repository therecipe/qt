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

void QRegExpValidator_SetRegExp(QtObjectPtr ptr, QtObjectPtr rx){
	static_cast<QRegExpValidator*>(ptr)->setRegExp(*static_cast<QRegExp*>(rx));
}

QtObjectPtr QRegExpValidator_NewQRegExpValidator(QtObjectPtr parent){
	return new QRegExpValidator(static_cast<QObject*>(parent));
}

QtObjectPtr QRegExpValidator_NewQRegExpValidator2(QtObjectPtr rx, QtObjectPtr parent){
	return new QRegExpValidator(*static_cast<QRegExp*>(rx), static_cast<QObject*>(parent));
}

void QRegExpValidator_DestroyQRegExpValidator(QtObjectPtr ptr){
	static_cast<QRegExpValidator*>(ptr)->~QRegExpValidator();
}

