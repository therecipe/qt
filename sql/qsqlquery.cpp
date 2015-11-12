#include "qsqlquery.h"
#include <QUrl>
#include <QModelIndex>
#include <QSqlDriver>
#include <QSqlDatabase>
#include <QSqlResult>
#include <QString>
#include <QVariant>
#include <QSqlQuery>
#include "_cgo_export.h"

class MyQSqlQuery: public QSqlQuery {
public:
};

void* QSqlQuery_NewQSqlQuery3(void* db){
	return new QSqlQuery(*static_cast<QSqlDatabase*>(db));
}

void* QSqlQuery_NewQSqlQuery(void* result){
	return new QSqlQuery(static_cast<QSqlResult*>(result));
}

void* QSqlQuery_NewQSqlQuery4(void* other){
	return new QSqlQuery(*static_cast<QSqlQuery*>(other));
}

void* QSqlQuery_NewQSqlQuery2(char* query, void* db){
	return new QSqlQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

int QSqlQuery_At(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->at();
}

void* QSqlQuery_BoundValue(void* ptr, char* placeholder){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(QString(placeholder)));
}

void* QSqlQuery_BoundValue2(void* ptr, int pos){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(pos));
}

void QSqlQuery_Clear(void* ptr){
	static_cast<QSqlQuery*>(ptr)->clear();
}

void* QSqlQuery_Driver(void* ptr){
	return const_cast<QSqlDriver*>(static_cast<QSqlQuery*>(ptr)->driver());
}

int QSqlQuery_Exec2(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->exec();
}

int QSqlQuery_Exec(void* ptr, char* query){
	return static_cast<QSqlQuery*>(ptr)->exec(QString(query));
}

int QSqlQuery_ExecBatch(void* ptr, int mode){
	return static_cast<QSqlQuery*>(ptr)->execBatch(static_cast<QSqlQuery::BatchExecutionMode>(mode));
}

char* QSqlQuery_ExecutedQuery(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->executedQuery().toUtf8().data();
}

void QSqlQuery_Finish(void* ptr){
	static_cast<QSqlQuery*>(ptr)->finish();
}

int QSqlQuery_First(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->first();
}

int QSqlQuery_IsActive(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isActive();
}

int QSqlQuery_IsForwardOnly(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isForwardOnly();
}

int QSqlQuery_IsNull2(void* ptr, char* name){
	return static_cast<QSqlQuery*>(ptr)->isNull(QString(name));
}

int QSqlQuery_IsNull(void* ptr, int field){
	return static_cast<QSqlQuery*>(ptr)->isNull(field);
}

int QSqlQuery_IsSelect(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isSelect();
}

int QSqlQuery_IsValid(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isValid();
}

int QSqlQuery_Last(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->last();
}

void* QSqlQuery_LastInsertId(void* ptr){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->lastInsertId());
}

char* QSqlQuery_LastQuery(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->lastQuery().toUtf8().data();
}

int QSqlQuery_Next(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->next();
}

int QSqlQuery_NextResult(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->nextResult();
}

int QSqlQuery_NumRowsAffected(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->numRowsAffected();
}

int QSqlQuery_Prepare(void* ptr, char* query){
	return static_cast<QSqlQuery*>(ptr)->prepare(QString(query));
}

int QSqlQuery_Previous(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->previous();
}

void* QSqlQuery_Result(void* ptr){
	return const_cast<QSqlResult*>(static_cast<QSqlQuery*>(ptr)->result());
}

int QSqlQuery_Seek(void* ptr, int index, int relative){
	return static_cast<QSqlQuery*>(ptr)->seek(index, relative != 0);
}

void QSqlQuery_SetForwardOnly(void* ptr, int forward){
	static_cast<QSqlQuery*>(ptr)->setForwardOnly(forward != 0);
}

int QSqlQuery_Size(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->size();
}

void* QSqlQuery_Value2(void* ptr, char* name){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(QString(name)));
}

void* QSqlQuery_Value(void* ptr, int index){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(index));
}

void QSqlQuery_DestroyQSqlQuery(void* ptr){
	static_cast<QSqlQuery*>(ptr)->~QSqlQuery();
}

