#include "qsqlrelationaltablemodel.h"
#include <QSqlRelation>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlRelationalTableModel>
#include "_cgo_export.h"

class MyQSqlRelationalTableModel: public QSqlRelationalTableModel {
public:
};

void QSqlRelationalTableModel_Clear(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->clear();
}

void* QSqlRelationalTableModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QSqlRelationalTableModel_RelationModel(void* ptr, int column){
	return static_cast<QSqlRelationalTableModel*>(ptr)->relationModel(column);
}

int QSqlRelationalTableModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlRelationalTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_RevertRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revertRow", Q_ARG(int, row));
}

int QSqlRelationalTableModel_Select(void* ptr){
	return static_cast<QSqlRelationalTableModel*>(ptr)->select();
}

int QSqlRelationalTableModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSqlRelationalTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlRelationalTableModel_SetJoinMode(void* ptr, int joinMode){
	static_cast<QSqlRelationalTableModel*>(ptr)->setJoinMode(static_cast<QSqlRelationalTableModel::JoinMode>(joinMode));
}

void QSqlRelationalTableModel_SetRelation(void* ptr, int column, void* relation){
	static_cast<QSqlRelationalTableModel*>(ptr)->setRelation(column, *static_cast<QSqlRelation*>(relation));
}

void QSqlRelationalTableModel_SetTable(void* ptr, char* table){
	static_cast<QSqlRelationalTableModel*>(ptr)->setTable(QString(table));
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->~QSqlRelationalTableModel();
}

