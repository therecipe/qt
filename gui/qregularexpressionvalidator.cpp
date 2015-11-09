#include "qregularexpressionvalidator.h"
#include <QMetaObject>
#include <QRegularExpression>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegularExpressionValidator>
#include "_cgo_export.h"

class MyQRegularExpressionValidator: public QRegularExpressionValidator {
public:
void Signal_RegularExpressionChanged(const QRegularExpression & re){callbackQRegularExpressionValidatorRegularExpressionChanged(this->objectName().toUtf8().data(), new QRegularExpression(re));};
};

void* QRegularExpressionValidator_RegularExpression(void* ptr){
	return new QRegularExpression(static_cast<QRegularExpressionValidator*>(ptr)->regularExpression());
}

void QRegularExpressionValidator_SetRegularExpression(void* ptr, void* re){
	QMetaObject::invokeMethod(static_cast<QRegularExpressionValidator*>(ptr), "setRegularExpression", Q_ARG(QRegularExpression, *static_cast<QRegularExpression*>(re)));
}

void* QRegularExpressionValidator_NewQRegularExpressionValidator(void* parent){
	return new QRegularExpressionValidator(static_cast<QObject*>(parent));
}

void* QRegularExpressionValidator_NewQRegularExpressionValidator2(void* re, void* parent){
	return new QRegularExpressionValidator(*static_cast<QRegularExpression*>(re), static_cast<QObject*>(parent));
}

void QRegularExpressionValidator_ConnectRegularExpressionChanged(void* ptr){
	QObject::connect(static_cast<QRegularExpressionValidator*>(ptr), static_cast<void (QRegularExpressionValidator::*)(const QRegularExpression &)>(&QRegularExpressionValidator::regularExpressionChanged), static_cast<MyQRegularExpressionValidator*>(ptr), static_cast<void (MyQRegularExpressionValidator::*)(const QRegularExpression &)>(&MyQRegularExpressionValidator::Signal_RegularExpressionChanged));;
}

void QRegularExpressionValidator_DisconnectRegularExpressionChanged(void* ptr){
	QObject::disconnect(static_cast<QRegularExpressionValidator*>(ptr), static_cast<void (QRegularExpressionValidator::*)(const QRegularExpression &)>(&QRegularExpressionValidator::regularExpressionChanged), static_cast<MyQRegularExpressionValidator*>(ptr), static_cast<void (MyQRegularExpressionValidator::*)(const QRegularExpression &)>(&MyQRegularExpressionValidator::Signal_RegularExpressionChanged));;
}

void QRegularExpressionValidator_DestroyQRegularExpressionValidator(void* ptr){
	static_cast<QRegularExpressionValidator*>(ptr)->~QRegularExpressionValidator();
}

