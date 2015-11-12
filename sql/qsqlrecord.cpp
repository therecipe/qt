#include "qsqlrecord.h"
#include <QModelIndex>
#include <QSqlField>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlRecord>
#include "_cgo_export.h"

class MyQSqlRecord: public QSqlRecord {
public:
};

void* QSqlRecord_NewQSqlRecord(){
	return new QSqlRecord();
}

void* QSqlRecord_NewQSqlRecord2(void* other){
	return new QSqlRecord(*static_cast<QSqlRecord*>(other));
}

void QSqlRecord_Append(void* ptr, void* field){
	static_cast<QSqlRecord*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlRecord_Clear(void* ptr){
	static_cast<QSqlRecord*>(ptr)->clear();
}

void QSqlRecord_ClearValues(void* ptr){
	static_cast<QSqlRecord*>(ptr)->clearValues();
}

int QSqlRecord_Contains(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->contains(QString(name));
}

int QSqlRecord_Count(void* ptr){
	return static_cast<QSqlRecord*>(ptr)->count();
}

char* QSqlRecord_FieldName(void* ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->fieldName(index).toUtf8().data();
}

int QSqlRecord_IndexOf(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->indexOf(QString(name));
}

int QSqlRecord_IsEmpty(void* ptr){
	return static_cast<QSqlRecord*>(ptr)->isEmpty();
}

int QSqlRecord_IsGenerated(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->isGenerated(QString(name));
}

int QSqlRecord_IsGenerated2(void* ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->isGenerated(index);
}

int QSqlRecord_IsNull(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->isNull(QString(name));
}

int QSqlRecord_IsNull2(void* ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->isNull(index);
}

void QSqlRecord_SetGenerated(void* ptr, char* name, int generated){
	static_cast<QSqlRecord*>(ptr)->setGenerated(QString(name), generated != 0);
}

void QSqlRecord_SetGenerated2(void* ptr, int index, int generated){
	static_cast<QSqlRecord*>(ptr)->setGenerated(index, generated != 0);
}

void QSqlRecord_SetNull2(void* ptr, char* name){
	static_cast<QSqlRecord*>(ptr)->setNull(QString(name));
}

void QSqlRecord_SetNull(void* ptr, int index){
	static_cast<QSqlRecord*>(ptr)->setNull(index);
}

void QSqlRecord_SetValue2(void* ptr, char* name, void* val){
	static_cast<QSqlRecord*>(ptr)->setValue(QString(name), *static_cast<QVariant*>(val));
}

void QSqlRecord_SetValue(void* ptr, int index, void* val){
	static_cast<QSqlRecord*>(ptr)->setValue(index, *static_cast<QVariant*>(val));
}

void* QSqlRecord_Value2(void* ptr, char* name){
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(QString(name)));
}

void* QSqlRecord_Value(void* ptr, int index){
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(index));
}

void QSqlRecord_DestroyQSqlRecord(void* ptr){
	static_cast<QSqlRecord*>(ptr)->~QSqlRecord();
}

