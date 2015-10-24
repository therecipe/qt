#include "qsqlquerymodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlQuery>
#include <QSqlDatabase>
#include <QSqlQueryModel>
#include "_cgo_export.h"

class MyQSqlQueryModel: public QSqlQueryModel {
public:
};

int QSqlQueryModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSqlQueryModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

char* QSqlQueryModel_Data(QtObjectPtr ptr, QtObjectPtr item, int role){
	return static_cast<QSqlQueryModel*>(ptr)->data(*static_cast<QModelIndex*>(item), role).toString().toUtf8().data();
}

int QSqlQueryModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSqlQueryModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_Clear(QtObjectPtr ptr){
	static_cast<QSqlQueryModel*>(ptr)->clear();
}

int QSqlQueryModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QSqlQueryModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlQueryModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QSqlQueryModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

char* QSqlQueryModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QSqlQueryModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

int QSqlQueryModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QSqlQueryModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QSqlQueryModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role){
	return static_cast<QSqlQueryModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), QVariant(value), role);
}

void QSqlQueryModel_SetQuery(QtObjectPtr ptr, QtObjectPtr query){
	static_cast<QSqlQueryModel*>(ptr)->setQuery(*static_cast<QSqlQuery*>(query));
}

void QSqlQueryModel_SetQuery2(QtObjectPtr ptr, char* query, QtObjectPtr db){
	static_cast<QSqlQueryModel*>(ptr)->setQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

void QSqlQueryModel_DestroyQSqlQueryModel(QtObjectPtr ptr){
	static_cast<QSqlQueryModel*>(ptr)->~QSqlQueryModel();
}

