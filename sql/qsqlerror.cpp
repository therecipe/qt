#include "qsqlerror.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSqlError>
#include "_cgo_export.h"

class MyQSqlError: public QSqlError {
public:
};

void* QSqlError_NewQSqlError3(void* other){
	return new QSqlError(*static_cast<QSqlError*>(other));
}

void* QSqlError_NewQSqlError(char* driverText, char* databaseText, int ty, char* code){
	return new QSqlError(QString(driverText), QString(databaseText), static_cast<QSqlError::ErrorType>(ty), QString(code));
}

char* QSqlError_DatabaseText(void* ptr){
	return static_cast<QSqlError*>(ptr)->databaseText().toUtf8().data();
}

char* QSqlError_DriverText(void* ptr){
	return static_cast<QSqlError*>(ptr)->driverText().toUtf8().data();
}

int QSqlError_IsValid(void* ptr){
	return static_cast<QSqlError*>(ptr)->isValid();
}

char* QSqlError_NativeErrorCode(void* ptr){
	return static_cast<QSqlError*>(ptr)->nativeErrorCode().toUtf8().data();
}

char* QSqlError_Text(void* ptr){
	return static_cast<QSqlError*>(ptr)->text().toUtf8().data();
}

int QSqlError_Type(void* ptr){
	return static_cast<QSqlError*>(ptr)->type();
}

void QSqlError_DestroyQSqlError(void* ptr){
	static_cast<QSqlError*>(ptr)->~QSqlError();
}

