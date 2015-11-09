#include "qsqlrelation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlRelation>
#include "_cgo_export.h"

class MyQSqlRelation: public QSqlRelation {
public:
};

void* QSqlRelation_NewQSqlRelation(){
	return new QSqlRelation();
}

void* QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn){
	return new QSqlRelation(QString(tableName), QString(indexColumn), QString(displayColumn));
}

char* QSqlRelation_DisplayColumn(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->displayColumn().toUtf8().data();
}

char* QSqlRelation_IndexColumn(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->indexColumn().toUtf8().data();
}

int QSqlRelation_IsValid(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->isValid();
}

char* QSqlRelation_TableName(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->tableName().toUtf8().data();
}

