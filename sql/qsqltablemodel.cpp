#include "qsqltablemodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlRecord>
#include <QObject>
#include <QMetaObject>
#include <QSqlTableModel>
#include "_cgo_export.h"

class MyQSqlTableModel: public QSqlTableModel {
public:
void Signal_BeforeDelete(int row){callbackQSqlTableModelBeforeDelete(this->objectName().toUtf8().data(), row);};
};

void QSqlTableModel_ConnectBeforeDelete(QtObjectPtr ptr){
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_DisconnectBeforeDelete(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_Clear(QtObjectPtr ptr){
	static_cast<QSqlTableModel*>(ptr)->clear();
}

char* QSqlTableModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QSqlTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

int QSqlTableModel_EditStrategy(QtObjectPtr ptr){
	return static_cast<QSqlTableModel*>(ptr)->editStrategy();
}

int QSqlTableModel_FieldIndex(QtObjectPtr ptr, char* fieldName){
	return static_cast<QSqlTableModel*>(ptr)->fieldIndex(QString(fieldName));
}

char* QSqlTableModel_Filter(QtObjectPtr ptr){
	return static_cast<QSqlTableModel*>(ptr)->filter().toUtf8().data();
}

int QSqlTableModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QSqlTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

char* QSqlTableModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QSqlTableModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

int QSqlTableModel_InsertRecord(QtObjectPtr ptr, int row, QtObjectPtr record){
	return static_cast<QSqlTableModel*>(ptr)->insertRecord(row, *static_cast<QSqlRecord*>(record));
}

int QSqlTableModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QSqlTableModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_IsDirty2(QtObjectPtr ptr){
	return static_cast<QSqlTableModel*>(ptr)->isDirty();
}

int QSqlTableModel_IsDirty(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QSqlTableModel*>(ptr)->isDirty(*static_cast<QModelIndex*>(index));
}

int QSqlTableModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QSqlTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QSqlTableModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_Revert(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revert");
}

void QSqlTableModel_RevertAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revertAll");
}

void QSqlTableModel_RevertRow(QtObjectPtr ptr, int row){
	static_cast<QSqlTableModel*>(ptr)->revertRow(row);
}

int QSqlTableModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSqlTableModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_Select(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "select");
}

int QSqlTableModel_SelectRow(QtObjectPtr ptr, int row){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "selectRow", Q_ARG(int, row));
}

int QSqlTableModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QSqlTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

void QSqlTableModel_SetEditStrategy(QtObjectPtr ptr, int strategy){
	static_cast<QSqlTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetFilter(QtObjectPtr ptr, char* filter){
	static_cast<QSqlTableModel*>(ptr)->setFilter(QString(filter));
}

int QSqlTableModel_SetRecord(QtObjectPtr ptr, int row, QtObjectPtr values){
	return static_cast<QSqlTableModel*>(ptr)->setRecord(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_SetSort(QtObjectPtr ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetTable(QtObjectPtr ptr, char* tableName){
	static_cast<QSqlTableModel*>(ptr)->setTable(QString(tableName));
}

void QSqlTableModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QSqlTableModel_Submit(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submit");
}

int QSqlTableModel_SubmitAll(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submitAll");
}

char* QSqlTableModel_TableName(QtObjectPtr ptr){
	return static_cast<QSqlTableModel*>(ptr)->tableName().toUtf8().data();
}

void QSqlTableModel_DestroyQSqlTableModel(QtObjectPtr ptr){
	static_cast<QSqlTableModel*>(ptr)->~QSqlTableModel();
}

