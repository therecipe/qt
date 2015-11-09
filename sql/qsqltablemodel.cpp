#include "qsqltablemodel.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QSqlRecord>
#include <QString>
#include <QSqlTableModel>
#include "_cgo_export.h"

class MyQSqlTableModel: public QSqlTableModel {
public:
void Signal_BeforeDelete(int row){callbackQSqlTableModelBeforeDelete(this->objectName().toUtf8().data(), row);};
};

void QSqlTableModel_ConnectBeforeDelete(void* ptr){
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_DisconnectBeforeDelete(void* ptr){
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_Clear(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->clear();
}

void* QSqlTableModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QSqlTableModel_EditStrategy(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->editStrategy();
}

int QSqlTableModel_FieldIndex(void* ptr, char* fieldName){
	return static_cast<QSqlTableModel*>(ptr)->fieldIndex(QString(fieldName));
}

char* QSqlTableModel_Filter(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->filter().toUtf8().data();
}

int QSqlTableModel_Flags(void* ptr, void* index){
	return static_cast<QSqlTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

void* QSqlTableModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

int QSqlTableModel_InsertRecord(void* ptr, int row, void* record){
	return static_cast<QSqlTableModel*>(ptr)->insertRecord(row, *static_cast<QSqlRecord*>(record));
}

int QSqlTableModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_IsDirty2(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->isDirty();
}

int QSqlTableModel_IsDirty(void* ptr, void* index){
	return static_cast<QSqlTableModel*>(ptr)->isDirty(*static_cast<QModelIndex*>(index));
}

int QSqlTableModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revert");
}

void QSqlTableModel_RevertAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revertAll");
}

void QSqlTableModel_RevertRow(void* ptr, int row){
	static_cast<QSqlTableModel*>(ptr)->revertRow(row);
}

int QSqlTableModel_RowCount(void* ptr, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_Select(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "select");
}

int QSqlTableModel_SelectRow(void* ptr, int row){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "selectRow", Q_ARG(int, row));
}

int QSqlTableModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSqlTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlTableModel_SetEditStrategy(void* ptr, int strategy){
	static_cast<QSqlTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetFilter(void* ptr, char* filter){
	static_cast<QSqlTableModel*>(ptr)->setFilter(QString(filter));
}

int QSqlTableModel_SetRecord(void* ptr, int row, void* values){
	return static_cast<QSqlTableModel*>(ptr)->setRecord(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_SetSort(void* ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetTable(void* ptr, char* tableName){
	static_cast<QSqlTableModel*>(ptr)->setTable(QString(tableName));
}

void QSqlTableModel_Sort(void* ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QSqlTableModel_Submit(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submit");
}

int QSqlTableModel_SubmitAll(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submitAll");
}

char* QSqlTableModel_TableName(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->tableName().toUtf8().data();
}

void QSqlTableModel_DestroyQSqlTableModel(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->~QSqlTableModel();
}

