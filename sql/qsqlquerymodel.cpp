#include "qsqlquerymodel.h"
#include <QModelIndex>
#include <QSqlQuery>
#include <QSqlDatabase>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlQueryModel>
#include "_cgo_export.h"

class MyQSqlQueryModel: public QSqlQueryModel {
public:
};

int QSqlQueryModel_RowCount(void* ptr, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_Data(void* ptr, void* item, int role){
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->data(*static_cast<QModelIndex*>(item), role));
}

int QSqlQueryModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_Clear(void* ptr){
	static_cast<QSqlQueryModel*>(ptr)->clear();
}

int QSqlQueryModel_ColumnCount(void* ptr, void* index){
	return static_cast<QSqlQueryModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlQueryModel_FetchMore(void* ptr, void* parent){
	static_cast<QSqlQueryModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

int QSqlQueryModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QSqlQueryModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QSqlQueryModel_SetQuery(void* ptr, void* query){
	static_cast<QSqlQueryModel*>(ptr)->setQuery(*static_cast<QSqlQuery*>(query));
}

void QSqlQueryModel_SetQuery2(void* ptr, char* query, void* db){
	static_cast<QSqlQueryModel*>(ptr)->setQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

void QSqlQueryModel_DestroyQSqlQueryModel(void* ptr){
	static_cast<QSqlQueryModel*>(ptr)->~QSqlQueryModel();
}

