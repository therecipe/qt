#include "qsqlquery.h"
#include <QModelIndex>
#include <QSqlDatabase>
#include <QSqlResult>
#include <QSqlDriver>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlQuery>
#include "_cgo_export.h"

class MyQSqlQuery: public QSqlQuery {
public:
};

QtObjectPtr QSqlQuery_NewQSqlQuery3(QtObjectPtr db){
	return new QSqlQuery(*static_cast<QSqlDatabase*>(db));
}

QtObjectPtr QSqlQuery_NewQSqlQuery(QtObjectPtr result){
	return new QSqlQuery(static_cast<QSqlResult*>(result));
}

QtObjectPtr QSqlQuery_NewQSqlQuery4(QtObjectPtr other){
	return new QSqlQuery(*static_cast<QSqlQuery*>(other));
}

QtObjectPtr QSqlQuery_NewQSqlQuery2(char* query, QtObjectPtr db){
	return new QSqlQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

int QSqlQuery_At(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->at();
}

char* QSqlQuery_BoundValue(QtObjectPtr ptr, char* placeholder){
	return static_cast<QSqlQuery*>(ptr)->boundValue(QString(placeholder)).toString().toUtf8().data();
}

char* QSqlQuery_BoundValue2(QtObjectPtr ptr, int pos){
	return static_cast<QSqlQuery*>(ptr)->boundValue(pos).toString().toUtf8().data();
}

void QSqlQuery_Clear(QtObjectPtr ptr){
	static_cast<QSqlQuery*>(ptr)->clear();
}

QtObjectPtr QSqlQuery_Driver(QtObjectPtr ptr){
	return const_cast<QSqlDriver*>(static_cast<QSqlQuery*>(ptr)->driver());
}

int QSqlQuery_Exec2(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->exec();
}

int QSqlQuery_Exec(QtObjectPtr ptr, char* query){
	return static_cast<QSqlQuery*>(ptr)->exec(QString(query));
}

int QSqlQuery_ExecBatch(QtObjectPtr ptr, int mode){
	return static_cast<QSqlQuery*>(ptr)->execBatch(static_cast<QSqlQuery::BatchExecutionMode>(mode));
}

char* QSqlQuery_ExecutedQuery(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->executedQuery().toUtf8().data();
}

void QSqlQuery_Finish(QtObjectPtr ptr){
	static_cast<QSqlQuery*>(ptr)->finish();
}

int QSqlQuery_First(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->first();
}

int QSqlQuery_IsActive(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->isActive();
}

int QSqlQuery_IsForwardOnly(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->isForwardOnly();
}

int QSqlQuery_IsNull2(QtObjectPtr ptr, char* name){
	return static_cast<QSqlQuery*>(ptr)->isNull(QString(name));
}

int QSqlQuery_IsNull(QtObjectPtr ptr, int field){
	return static_cast<QSqlQuery*>(ptr)->isNull(field);
}

int QSqlQuery_IsSelect(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->isSelect();
}

int QSqlQuery_IsValid(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->isValid();
}

int QSqlQuery_Last(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->last();
}

char* QSqlQuery_LastInsertId(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->lastInsertId().toString().toUtf8().data();
}

char* QSqlQuery_LastQuery(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->lastQuery().toUtf8().data();
}

int QSqlQuery_Next(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->next();
}

int QSqlQuery_NextResult(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->nextResult();
}

int QSqlQuery_NumRowsAffected(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->numRowsAffected();
}

int QSqlQuery_Prepare(QtObjectPtr ptr, char* query){
	return static_cast<QSqlQuery*>(ptr)->prepare(QString(query));
}

int QSqlQuery_Previous(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->previous();
}

QtObjectPtr QSqlQuery_Result(QtObjectPtr ptr){
	return const_cast<QSqlResult*>(static_cast<QSqlQuery*>(ptr)->result());
}

int QSqlQuery_Seek(QtObjectPtr ptr, int index, int relative){
	return static_cast<QSqlQuery*>(ptr)->seek(index, relative != 0);
}

void QSqlQuery_SetForwardOnly(QtObjectPtr ptr, int forward){
	static_cast<QSqlQuery*>(ptr)->setForwardOnly(forward != 0);
}

int QSqlQuery_Size(QtObjectPtr ptr){
	return static_cast<QSqlQuery*>(ptr)->size();
}

char* QSqlQuery_Value2(QtObjectPtr ptr, char* name){
	return static_cast<QSqlQuery*>(ptr)->value(QString(name)).toString().toUtf8().data();
}

char* QSqlQuery_Value(QtObjectPtr ptr, int index){
	return static_cast<QSqlQuery*>(ptr)->value(index).toString().toUtf8().data();
}

void QSqlQuery_DestroyQSqlQuery(QtObjectPtr ptr){
	static_cast<QSqlQuery*>(ptr)->~QSqlQuery();
}

