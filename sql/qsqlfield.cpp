#include "qsqlfield.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlField>
#include "_cgo_export.h"

class MyQSqlField: public QSqlField {
public:
};

void* QSqlField_NewQSqlField2(void* other){
	return new QSqlField(*static_cast<QSqlField*>(other));
}

void QSqlField_Clear(void* ptr){
	static_cast<QSqlField*>(ptr)->clear();
}

void* QSqlField_DefaultValue(void* ptr){
	return new QVariant(static_cast<QSqlField*>(ptr)->defaultValue());
}

int QSqlField_IsAutoValue(void* ptr){
	return static_cast<QSqlField*>(ptr)->isAutoValue();
}

int QSqlField_IsGenerated(void* ptr){
	return static_cast<QSqlField*>(ptr)->isGenerated();
}

int QSqlField_IsNull(void* ptr){
	return static_cast<QSqlField*>(ptr)->isNull();
}

int QSqlField_IsReadOnly(void* ptr){
	return static_cast<QSqlField*>(ptr)->isReadOnly();
}

int QSqlField_IsValid(void* ptr){
	return static_cast<QSqlField*>(ptr)->isValid();
}

int QSqlField_Length(void* ptr){
	return static_cast<QSqlField*>(ptr)->length();
}

char* QSqlField_Name(void* ptr){
	return static_cast<QSqlField*>(ptr)->name().toUtf8().data();
}

int QSqlField_Precision(void* ptr){
	return static_cast<QSqlField*>(ptr)->precision();
}

int QSqlField_RequiredStatus(void* ptr){
	return static_cast<QSqlField*>(ptr)->requiredStatus();
}

void QSqlField_SetAutoValue(void* ptr, int autoVal){
	static_cast<QSqlField*>(ptr)->setAutoValue(autoVal != 0);
}

void QSqlField_SetDefaultValue(void* ptr, void* value){
	static_cast<QSqlField*>(ptr)->setDefaultValue(*static_cast<QVariant*>(value));
}

void QSqlField_SetGenerated(void* ptr, int gen){
	static_cast<QSqlField*>(ptr)->setGenerated(gen != 0);
}

void QSqlField_SetLength(void* ptr, int fieldLength){
	static_cast<QSqlField*>(ptr)->setLength(fieldLength);
}

void QSqlField_SetName(void* ptr, char* name){
	static_cast<QSqlField*>(ptr)->setName(QString(name));
}

void QSqlField_SetPrecision(void* ptr, int precision){
	static_cast<QSqlField*>(ptr)->setPrecision(precision);
}

void QSqlField_SetReadOnly(void* ptr, int readOnly){
	static_cast<QSqlField*>(ptr)->setReadOnly(readOnly != 0);
}

void QSqlField_SetRequired(void* ptr, int required){
	static_cast<QSqlField*>(ptr)->setRequired(required != 0);
}

void QSqlField_SetRequiredStatus(void* ptr, int required){
	static_cast<QSqlField*>(ptr)->setRequiredStatus(static_cast<QSqlField::RequiredStatus>(required));
}

void QSqlField_SetValue(void* ptr, void* value){
	static_cast<QSqlField*>(ptr)->setValue(*static_cast<QVariant*>(value));
}

void* QSqlField_Value(void* ptr){
	return new QVariant(static_cast<QSqlField*>(ptr)->value());
}

void QSqlField_DestroyQSqlField(void* ptr){
	static_cast<QSqlField*>(ptr)->~QSqlField();
}

