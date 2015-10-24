#include "qregularexpressionvalidator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegularExpression>
#include <QMetaObject>
#include <QObject>
#include <QRegularExpressionValidator>
#include "_cgo_export.h"

class MyQRegularExpressionValidator: public QRegularExpressionValidator {
public:
};

void QRegularExpressionValidator_SetRegularExpression(QtObjectPtr ptr, QtObjectPtr re){
	QMetaObject::invokeMethod(static_cast<QRegularExpressionValidator*>(ptr), "setRegularExpression", Q_ARG(QRegularExpression, *static_cast<QRegularExpression*>(re)));
}

QtObjectPtr QRegularExpressionValidator_NewQRegularExpressionValidator(QtObjectPtr parent){
	return new QRegularExpressionValidator(static_cast<QObject*>(parent));
}

QtObjectPtr QRegularExpressionValidator_NewQRegularExpressionValidator2(QtObjectPtr re, QtObjectPtr parent){
	return new QRegularExpressionValidator(*static_cast<QRegularExpression*>(re), static_cast<QObject*>(parent));
}

void QRegularExpressionValidator_DestroyQRegularExpressionValidator(QtObjectPtr ptr){
	static_cast<QRegularExpressionValidator*>(ptr)->~QRegularExpressionValidator();
}

