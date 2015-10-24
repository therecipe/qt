#include "qsqlerror.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlError>
#include "_cgo_export.h"

class MyQSqlError: public QSqlError {
public:
};

QtObjectPtr QSqlError_NewQSqlError3(QtObjectPtr other){
	return new QSqlError(*static_cast<QSqlError*>(other));
}

QtObjectPtr QSqlError_NewQSqlError(char* driverText, char* databaseText, int ty, char* code){
	return new QSqlError(QString(driverText), QString(databaseText), static_cast<QSqlError::ErrorType>(ty), QString(code));
}

char* QSqlError_DatabaseText(QtObjectPtr ptr){
	return static_cast<QSqlError*>(ptr)->databaseText().toUtf8().data();
}

char* QSqlError_DriverText(QtObjectPtr ptr){
	return static_cast<QSqlError*>(ptr)->driverText().toUtf8().data();
}

int QSqlError_IsValid(QtObjectPtr ptr){
	return static_cast<QSqlError*>(ptr)->isValid();
}

char* QSqlError_NativeErrorCode(QtObjectPtr ptr){
	return static_cast<QSqlError*>(ptr)->nativeErrorCode().toUtf8().data();
}

char* QSqlError_Text(QtObjectPtr ptr){
	return static_cast<QSqlError*>(ptr)->text().toUtf8().data();
}

int QSqlError_Type(QtObjectPtr ptr){
	return static_cast<QSqlError*>(ptr)->type();
}

void QSqlError_DestroyQSqlError(QtObjectPtr ptr){
	static_cast<QSqlError*>(ptr)->~QSqlError();
}

