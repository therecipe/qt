#include "qsqlfield.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlField>
#include "_cgo_export.h"

class MyQSqlField: public QSqlField {
public:
};

QtObjectPtr QSqlField_NewQSqlField2(QtObjectPtr other){
	return new QSqlField(*static_cast<QSqlField*>(other));
}

void QSqlField_Clear(QtObjectPtr ptr){
	static_cast<QSqlField*>(ptr)->clear();
}

char* QSqlField_DefaultValue(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->defaultValue().toString().toUtf8().data();
}

int QSqlField_IsAutoValue(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->isAutoValue();
}

int QSqlField_IsGenerated(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->isGenerated();
}

int QSqlField_IsNull(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->isNull();
}

int QSqlField_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->isReadOnly();
}

int QSqlField_IsValid(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->isValid();
}

int QSqlField_Length(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->length();
}

char* QSqlField_Name(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->name().toUtf8().data();
}

int QSqlField_Precision(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->precision();
}

int QSqlField_RequiredStatus(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->requiredStatus();
}

void QSqlField_SetAutoValue(QtObjectPtr ptr, int autoVal){
	static_cast<QSqlField*>(ptr)->setAutoValue(autoVal != 0);
}

void QSqlField_SetDefaultValue(QtObjectPtr ptr, char* value){
	static_cast<QSqlField*>(ptr)->setDefaultValue(QVariant(value));
}

void QSqlField_SetGenerated(QtObjectPtr ptr, int gen){
	static_cast<QSqlField*>(ptr)->setGenerated(gen != 0);
}

void QSqlField_SetLength(QtObjectPtr ptr, int fieldLength){
	static_cast<QSqlField*>(ptr)->setLength(fieldLength);
}

void QSqlField_SetName(QtObjectPtr ptr, char* name){
	static_cast<QSqlField*>(ptr)->setName(QString(name));
}

void QSqlField_SetPrecision(QtObjectPtr ptr, int precision){
	static_cast<QSqlField*>(ptr)->setPrecision(precision);
}

void QSqlField_SetReadOnly(QtObjectPtr ptr, int readOnly){
	static_cast<QSqlField*>(ptr)->setReadOnly(readOnly != 0);
}

void QSqlField_SetRequired(QtObjectPtr ptr, int required){
	static_cast<QSqlField*>(ptr)->setRequired(required != 0);
}

void QSqlField_SetRequiredStatus(QtObjectPtr ptr, int required){
	static_cast<QSqlField*>(ptr)->setRequiredStatus(static_cast<QSqlField::RequiredStatus>(required));
}

void QSqlField_SetValue(QtObjectPtr ptr, char* value){
	static_cast<QSqlField*>(ptr)->setValue(QVariant(value));
}

char* QSqlField_Value(QtObjectPtr ptr){
	return static_cast<QSqlField*>(ptr)->value().toString().toUtf8().data();
}

void QSqlField_DestroyQSqlField(QtObjectPtr ptr){
	static_cast<QSqlField*>(ptr)->~QSqlField();
}

