#include "qvalidator.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QLocale>
#include <QString>
#include <QValidator>
#include "_cgo_export.h"

class MyQValidator: public QValidator {
public:
void Signal_Changed(){callbackQValidatorChanged(this->objectName().toUtf8().data());};
};

void QValidator_ConnectChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));;
}

void QValidator_DisconnectChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));;
}

void QValidator_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QValidator*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QValidator_DestroyQValidator(QtObjectPtr ptr){
	static_cast<QValidator*>(ptr)->~QValidator();
}

