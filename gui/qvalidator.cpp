#include "qvalidator.h"
#include <QLocale>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QValidator>
#include "_cgo_export.h"

class MyQValidator: public QValidator {
public:
void Signal_Changed(){callbackQValidatorChanged(this->objectName().toUtf8().data());};
};

void QValidator_ConnectChanged(void* ptr){
	QObject::connect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));;
}

void QValidator_DisconnectChanged(void* ptr){
	QObject::disconnect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));;
}

void QValidator_SetLocale(void* ptr, void* locale){
	static_cast<QValidator*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QValidator_DestroyQValidator(void* ptr){
	static_cast<QValidator*>(ptr)->~QValidator();
}

