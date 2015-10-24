#include "qsqlrecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlField>
#include <QSqlRecord>
#include "_cgo_export.h"

class MyQSqlRecord: public QSqlRecord {
public:
};

QtObjectPtr QSqlRecord_NewQSqlRecord(){
	return new QSqlRecord();
}

QtObjectPtr QSqlRecord_NewQSqlRecord2(QtObjectPtr other){
	return new QSqlRecord(*static_cast<QSqlRecord*>(other));
}

void QSqlRecord_Append(QtObjectPtr ptr, QtObjectPtr field){
	static_cast<QSqlRecord*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlRecord_Clear(QtObjectPtr ptr){
	static_cast<QSqlRecord*>(ptr)->clear();
}

void QSqlRecord_ClearValues(QtObjectPtr ptr){
	static_cast<QSqlRecord*>(ptr)->clearValues();
}

int QSqlRecord_Contains(QtObjectPtr ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->contains(QString(name));
}

int QSqlRecord_Count(QtObjectPtr ptr){
	return static_cast<QSqlRecord*>(ptr)->count();
}

char* QSqlRecord_FieldName(QtObjectPtr ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->fieldName(index).toUtf8().data();
}

int QSqlRecord_IndexOf(QtObjectPtr ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->indexOf(QString(name));
}

int QSqlRecord_IsEmpty(QtObjectPtr ptr){
	return static_cast<QSqlRecord*>(ptr)->isEmpty();
}

int QSqlRecord_IsGenerated(QtObjectPtr ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->isGenerated(QString(name));
}

int QSqlRecord_IsGenerated2(QtObjectPtr ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->isGenerated(index);
}

int QSqlRecord_IsNull(QtObjectPtr ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->isNull(QString(name));
}

int QSqlRecord_IsNull2(QtObjectPtr ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->isNull(index);
}

void QSqlRecord_SetGenerated(QtObjectPtr ptr, char* name, int generated){
	static_cast<QSqlRecord*>(ptr)->setGenerated(QString(name), generated != 0);
}

void QSqlRecord_SetGenerated2(QtObjectPtr ptr, int index, int generated){
	static_cast<QSqlRecord*>(ptr)->setGenerated(index, generated != 0);
}

void QSqlRecord_SetNull2(QtObjectPtr ptr, char* name){
	static_cast<QSqlRecord*>(ptr)->setNull(QString(name));
}

void QSqlRecord_SetNull(QtObjectPtr ptr, int index){
	static_cast<QSqlRecord*>(ptr)->setNull(index);
}

void QSqlRecord_SetValue2(QtObjectPtr ptr, char* name, char* val){
	static_cast<QSqlRecord*>(ptr)->setValue(QString(name), QVariant(val));
}

void QSqlRecord_SetValue(QtObjectPtr ptr, int index, char* val){
	static_cast<QSqlRecord*>(ptr)->setValue(index, QVariant(val));
}

char* QSqlRecord_Value2(QtObjectPtr ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->value(QString(name)).toString().toUtf8().data();
}

char* QSqlRecord_Value(QtObjectPtr ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->value(index).toString().toUtf8().data();
}

void QSqlRecord_DestroyQSqlRecord(QtObjectPtr ptr){
	static_cast<QSqlRecord*>(ptr)->~QSqlRecord();
}

