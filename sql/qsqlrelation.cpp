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

QtObjectPtr QSqlRelation_NewQSqlRelation(){
	return new QSqlRelation();
}

QtObjectPtr QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn){
	return new QSqlRelation(QString(tableName), QString(indexColumn), QString(displayColumn));
}

char* QSqlRelation_DisplayColumn(QtObjectPtr ptr){
	return static_cast<QSqlRelation*>(ptr)->displayColumn().toUtf8().data();
}

char* QSqlRelation_IndexColumn(QtObjectPtr ptr){
	return static_cast<QSqlRelation*>(ptr)->indexColumn().toUtf8().data();
}

int QSqlRelation_IsValid(QtObjectPtr ptr){
	return static_cast<QSqlRelation*>(ptr)->isValid();
}

char* QSqlRelation_TableName(QtObjectPtr ptr){
	return static_cast<QSqlRelation*>(ptr)->tableName().toUtf8().data();
}

