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

void QSqlRelationalTableModel_Clear(QtObjectPtr ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->clear();
}

char* QSqlRelationalTableModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QSqlRelationalTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

QtObjectPtr QSqlRelationalTableModel_RelationModel(QtObjectPtr ptr, int column){
	return static_cast<QSqlRelationalTableModel*>(ptr)->relationModel(column);
}

int QSqlRelationalTableModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QSqlRelationalTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_RevertRow(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revertRow", Q_ARG(int, row));
}

int QSqlRelationalTableModel_Select(QtObjectPtr ptr){
	return static_cast<QSqlRelationalTableModel*>(ptr)->select();
}

int QSqlRelationalTableModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QSqlRelationalTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

void QSqlRelationalTableModel_SetJoinMode(QtObjectPtr ptr, int joinMode){
	static_cast<QSqlRelationalTableModel*>(ptr)->setJoinMode(static_cast<QSqlRelationalTableModel::JoinMode>(joinMode));
}

void QSqlRelationalTableModel_SetRelation(QtObjectPtr ptr, int column, QtObjectPtr relation){
	static_cast<QSqlRelationalTableModel*>(ptr)->setRelation(column, *static_cast<QSqlRelation*>(relation));
}

void QSqlRelationalTableModel_SetTable(QtObjectPtr ptr, char* table){
	static_cast<QSqlRelationalTableModel*>(ptr)->setTable(QString(table));
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(QtObjectPtr ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->~QSqlRelationalTableModel();
}

